// Copyright 2023 The olive Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runner

import (
	"context"
	"encoding/binary"
	"net/url"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/gofrs/flock"
	"github.com/olive-io/bpmn/tracing"
	pb "github.com/olive-io/olive/api/olivepb"
	"github.com/olive-io/olive/client"
	dsy "github.com/olive-io/olive/pkg/discovery"
	"github.com/olive-io/olive/pkg/runtime"
	genericserver "github.com/olive-io/olive/pkg/server"
	"github.com/olive-io/olive/runner/backend"
	"github.com/olive-io/olive/runner/buckets"
	"github.com/olive-io/olive/runner/raft"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type Runner struct {
	genericserver.Inner

	ctx    context.Context
	cancel context.CancelFunc

	cfg Config

	gLock *flock.Flock

	oct *client.Client

	be backend.IBackend

	controller *raft.Controller
	traces     <-chan tracing.ITrace
	pr         *pb.Runner
}

func NewRunner(cfg Config) (*Runner, error) {
	lg := cfg.Logger

	gLock, err := cfg.LockDataDir()
	if err != nil {
		return nil, err
	}
	lg.Debug("protected directory: " + cfg.DataDir)

	oct, err := client.New(cfg.Config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	be := newBackend(&cfg)

	inner := genericserver.NewInnerServer(lg)
	runner := &Runner{
		Inner:  inner,
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
		gLock:  gLock,

		oct: oct,
		be:  be,
	}

	return runner, nil
}

func (r *Runner) Logger() *zap.Logger {
	return r.cfg.Logger
}

func (r *Runner) Start(stopc <-chan struct{}) error {
	if err := r.start(); err != nil {
		return err
	}

	r.Destroy(r.destroy)
	r.GoAttach(r.process)
	r.GoAttach(r.watching)

	<-stopc

	return r.stop()
}

func (r *Runner) start() error {
	var err error

	defer r.be.ForceCommit()
	tx := r.be.BatchTx()
	tx.Lock()
	tx.UnsafeCreateBucket(buckets.Meta)
	tx.UnsafeCreateBucket(buckets.Region)
	tx.UnsafeCreateBucket(buckets.Key)
	tx.Unlock()
	if err = tx.Commit(); err != nil {
		r.Logger().Error("pebble commit", zap.Error(err))
	}

	r.pr, err = r.register()
	if err != nil {
		return err
	}

	r.controller, r.traces, err = r.startRaftController()
	if err != nil {
		return err
	}

	return nil
}

func (r *Runner) stop() error {
	r.Inner.Shutdown()
	return nil
}

func (r *Runner) startRaftController() (*raft.Controller, <-chan tracing.ITrace, error) {
	cfg := r.cfg
	lg := r.Logger()

	dopts := []dsy.Option{
		dsy.Prefix(runtime.DefaultRunnerDiscoveryNode),
		dsy.SetLogger(lg),
	}
	discovery, err := dsy.NewDiscovery(r.oct.Client, dopts...)
	if err != nil {
		return nil, nil, err
	}

	listenPeerURL, err := url.Parse(cfg.ListenPeerURL)
	if err != nil {
		return nil, nil, err
	}
	raftAddr := listenPeerURL.Host

	cc := raft.NewConfig()
	cc.DataDir = cfg.RegionDir()
	cc.RaftAddress = raftAddr
	cc.HeartbeatMs = cfg.HeartbeatMs
	cc.RaftRTTMillisecond = cfg.RaftRTTMillisecond
	cc.Logger = lg

	controller, err := raft.NewController(r.ctx, cc, r.be, discovery, r.pr)
	if err != nil {
		return nil, nil, err
	}

	lg.Info("start raft container")
	if err = controller.Start(r.StoppingNotify()); err != nil {
		return nil, nil, err
	}

	ctx := r.ctx
	prefix := runtime.DefaultRunnerRegion
	rev := r.getRev()

	options := []clientv3.OpOption{
		clientv3.WithPrefix(),
		clientv3.WithSerializable(),
		clientv3.WithRev(rev),
	}

	rsp, err := r.oct.Get(ctx, prefix, append(options, clientv3.WithMinModRev(rev+1))...)
	if err != nil {
		return nil, nil, err
	}

	regions := make([]*pb.Region, 0)
	definitions := make([]*pb.Definition, 0)
	processes := make([]*pb.ProcessInstance, 0)
	for _, kv := range rsp.Kvs {
		rev = kv.ModRevision
		key := string(kv.Value)
		switch {
		case strings.HasPrefix(key, runtime.DefaultRunnerRegion):
			region, match, err := parseRegionKV(kv, r.pr.Id)
			if err != nil || !match {
				continue
			}
			regions = append(regions, region)
		case strings.HasPrefix(key, runtime.DefaultRunnerDefinitions):
			definition, match, err := parseDefinitionKV(kv)
			if err != nil || !match {
				continue
			}
			definitions = append(definitions, definition)
		case strings.HasPrefix(key, runtime.DefaultRunnerProcessInstance):
			proc, match, err := parseProcessInstanceKV(kv)
			if err != nil || !match {
				continue
			}
			processes = append(processes, proc)
		}
	}
	for _, region := range regions {
		if err = controller.SyncRegion(ctx, region); err != nil {
			return nil, nil, errors.Wrap(err, "sync region")
		}
	}
	for _, definition := range definitions {
		if err = controller.DeployDefinition(ctx, definition); err != nil {
			lg.Error("deploy definition", zap.Error(err))
		}
	}
	for _, process := range processes {
		if err = controller.ExecuteDefinition(ctx, process); err != nil {
			lg.Error("execute definition", zap.Error(err))
		} else {
			commitProcessInstance(ctx, lg, r.oct, process)
		}
	}

	r.setRev(rev)
	return controller, controller.SubscribeTrace(), nil
}

func (r *Runner) watching() {
	ctx, cancel := context.WithCancel(r.ctx)
	defer cancel()

	rev := r.getRev()
	wopts := []clientv3.OpOption{
		clientv3.WithPrefix(),
		clientv3.WithPrevKV(),
		clientv3.WithRev(rev + 1),
	}
	wch := r.oct.Watch(ctx, runtime.DefaultRunnerPrefix, wopts...)
	for {
		select {
		case <-r.StoppingNotify():
			return
		case resp := <-wch:
			for _, event := range resp.Events {
				if event.Kv.ModRevision > rev {
					rev = event.Kv.ModRevision
					r.setRev(rev)
				}
				r.processEvent(r.ctx, event)
			}
		}
	}
}

func (r *Runner) getRev() int64 {
	key := []byte("cur_rev")
	tx := r.be.ReadTx()
	tx.RLock()
	value, err := tx.UnsafeGet(buckets.Meta, key)
	tx.RUnlock()
	if err != nil || len(value) == 0 {
		return 0
	}
	value = value[:8]
	return int64(binary.LittleEndian.Uint64(value))
}

func (r *Runner) setRev(rev int64) {
	key := []byte("cur_rev")
	value := make([]byte, 8)
	binary.LittleEndian.PutUint64(value, uint64(rev))
	tx := r.be.BatchTx()
	tx.Lock()
	tx.UnsafePut(buckets.Meta, key, value)
	tx.Unlock()
	tx.Commit()
}

func (r *Runner) destroy() {
	r.cancel()
	if err := r.gLock.Unlock(); err != nil {
		r.Logger().Error("released "+r.cfg.DataDir, zap.Error(err))
	}
}

func (r *Runner) processEvent(ctx context.Context, event *clientv3.Event) {
	kv := event.Kv
	key := string(kv.Key)
	//if event.Type == clientv3.EventTypeDelete {
	//	rev := event.Kv.CreateRevision
	//	options := []clientv3.OpOption{clientv3.WithRev(rev)}
	//	rsp, err := r.oct.Get(ctx, string(event.Kv.Key), options...)
	//	if err != nil {
	//		r.Logger().Error("get key-value")
	//	}
	//	if len(rsp.Kvs) > 0 {
	//		event.Kv = rsp.Kvs[0]
	//	}
	//}
	switch {
	case strings.HasPrefix(key, runtime.DefaultRunnerRegion):
		r.processRegion(ctx, event)
	case strings.HasPrefix(key, runtime.DefaultRunnerDefinitions):
		r.processBpmnDefinition(ctx, event)
	case strings.HasPrefix(key, runtime.DefaultRunnerProcessInstance):
		r.processBpmnProcess(ctx, event)
	}
}

func (r *Runner) processRegion(ctx context.Context, event *clientv3.Event) {
	lg := r.Logger()
	kv := event.Kv
	if event.Type == clientv3.EventTypeDelete {
		return
	}

	region, match, err := parseRegionKV(kv, r.pr.Id)
	if err != nil {
		lg.Error("parse region data", zap.Error(err))
		return
	}
	if !match {
		return
	}

	if event.IsCreate() {
		lg.Info("create region", zap.Stringer("body", region))
		if err = r.controller.CreateRegion(ctx, region); err != nil {
			lg.Error("create region", zap.Error(err))
		}
		return
	}

	if event.IsModify() {
		lg.Info("sync region", zap.Stringer("body", region))
		if err = r.controller.SyncRegion(ctx, region); err != nil {
			lg.Error("sync region", zap.Error(err))
		}
		return
	}
}

func (r *Runner) processBpmnDefinition(ctx context.Context, event *clientv3.Event) {
	lg := r.Logger()
	kv := event.Kv
	if event.Type == clientv3.EventTypeDelete {
		return
	}

	definition, match, err := parseDefinitionKV(kv)
	if err != nil {
		lg.Error("parse definition data", zap.Error(err))
		return
	}
	if !match {
		return
	}

	if err = r.controller.DeployDefinition(ctx, definition); err != nil {
		lg.Error("definition deploy",
			zap.String("id", definition.Id),
			zap.Uint64("version", definition.Version),
			zap.Error(err))
		return
	}
}

func (r *Runner) processBpmnProcess(ctx context.Context, event *clientv3.Event) {
	lg := r.Logger()
	kv := event.Kv
	if event.Type == clientv3.EventTypeDelete {
		return
	}

	process, match, err := parseProcessInstanceKV(kv)
	if err != nil {
		lg.Error("parse process data", zap.Error(err))
		return
	}
	if !match {
		return
	}

	if err = r.controller.ExecuteDefinition(ctx, process); err != nil {
		lg.Error("execute definition",
			zap.String("id", process.DefinitionId),
			zap.Uint64("version", process.DefinitionVersion),
			zap.Error(err))
		return
	}

	commitProcessInstance(ctx, lg, r.oct, process)
}
