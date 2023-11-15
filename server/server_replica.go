package server

import (
	"context"
	"fmt"
	"math"
	"time"

	json "github.com/json-iterator/go"
	"github.com/lni/dragonboat/v4"
	"github.com/olive-io/olive/api/raftpb"
	"github.com/olive-io/olive/server/config"
	"github.com/olive-io/olive/server/membership"
	"go.uber.org/zap"
)

type replicaRequest struct {
	shardID        uint64
	nodeID         uint64
	staleRead      *replicaStaleRead
	syncRead       *replicaSyncRead
	syncPropose    *replicaSyncPropose
	syncConfChange *replicaSyncConfChange
}

type replicaStaleRead struct {
	query any
	rc    chan any
	ec    chan error
}

type replicaSyncRead struct {
	ctx   context.Context
	query any
	rc    chan any
	ec    chan error
}

type replicaSyncPropose struct {
	ctx  context.Context
	data []byte
	rc   chan any
	ec   chan error
}

type replicaSyncConfChange struct {
	ctx context.Context
	cc  raftpb.ConfChange
	rc  chan any
}

func (s *OliveServer) StartReplica(cfg config.ShardConfig) error {
	ra, members, join, rc, err := s.NewReplica(cfg)
	if err != nil {
		return err
	}

	start := time.Now()
	err = s.nh.StartOnDiskReplica(members, join, ra.NewDiskStateMachine, rc)
	if err != nil {
		return err
	}

	electionTimeout := cfg.ElectionTimeout
	if electionTimeout == 0 {
		electionTimeout = time.Duration(math.MaxInt64)
	}
	after := time.NewTimer(electionTimeout)
	defer after.Stop()
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()
	for {
		select {
		case <-s.stop:
			return ErrStopped
		case <-after.C:
			return fmt.Errorf("wait shard ready: %w", ErrTimeout)
		case <-ticker.C:
		}

		leaderID, term, ok, e1 := s.nh.GetLeaderID(cfg.ShardID)
		if ok {
			s.lg.Info("start new shard",
				zap.Uint64("leader", leaderID),
				zap.Uint64("term", term),
				zap.Duration("duration", time.Now().Sub(start)))

			break
		}
		if e1 != nil {
			return fmt.Errorf("get leader %v", e1)
		}
	}

	return nil
}

func (s *OliveServer) processReplicaEvent() {
	for {
		select {
		case <-s.stopping:
			return
		case ch := <-s.raftEventCh:
			ra, ok := s.GetReplica(ch.ShardID)
			if ok {
				ra.setLead(ch.LeaderID)
				ra.setTerm(ch.Term)
				ra.ChangeNotify()
			}
		}
	}
}

func (s *OliveServer) processReplicaRequest() {
	for {
		select {
		case <-s.stopping:
			return
		case req := <-s.replicaRequestC:
			if rr := req.staleRead; rr != nil {
				s.requestStaleRead(req.shardID, rr)
			}
			if rr := req.syncRead; rr != nil {
				s.requestSyncRead(req.shardID, rr)
			}
			if rr := req.syncPropose; rr != nil {
				s.requestSyncPropose(req.shardID, rr)
			}
			if rr := req.syncConfChange; rr != nil {
				s.requestConfChange(req.shardID, rr)
			}
		}
	}
}

func (s *OliveServer) requestStaleRead(shardID uint64, r *replicaStaleRead) {
	result, err := s.nh.StaleRead(shardID, r.query)
	if err != nil {
		r.ec <- err
		return
	}
	r.rc <- result
}

func (s *OliveServer) requestSyncRead(shardID uint64, r *replicaSyncRead) {
	result, err := s.nh.SyncRead(r.ctx, shardID, r.query)
	if err != nil {
		r.ec <- err
		return
	}
	r.rc <- result
}

func (s *OliveServer) requestSyncPropose(shardID uint64, r *replicaSyncPropose) {
	session := s.nh.GetNoOPSession(shardID)
	result, err := s.nh.SyncPropose(r.ctx, session, r.data)
	if err != nil {
		r.ec <- err
		return
	}
	r.rc <- result
}

func (s *OliveServer) requestConfChange(shardID uint64, r *replicaSyncConfChange) {
	ctx := r.ctx
	cc := r.cc

	resp := &confChangeResponse{}
	defer func() { r.rc <- resp }()

	ms, err := s.nh.SyncGetShardMembership(ctx, shardID)
	if err != nil {
		resp.err = err
		return
	}

	var memb *membership.Member
	_ = json.Unmarshal(cc.Context, &memb)

	switch cc.Type {
	case raftpb.ConfChangeType_ConfChangeAddNode:
		resp.err = s.nh.SyncRequestAddReplica(ctx, shardID, memb.ID, memb.PickPeerURL(), ms.ConfigChangeID)
	case raftpb.ConfChangeType_ConfChangeUpdateNode:

	case raftpb.ConfChangeType_ConfChangeRemoveNode:
		resp.err = s.nh.SyncRequestDeleteReplica(ctx, shardID, memb.ID, ms.ConfigChangeID)
	case raftpb.ConfChangeType_ConfChangeAddLearnerNode:
		resp.err = s.nh.SyncRequestAddNonVoting(ctx, shardID, memb.ID, memb.PickPeerURL(), ms.ConfigChangeID)
	}
}

func (s *OliveServer) GetReplica(shardID uint64) (*Replica, bool) {
	s.rmu.RLock()
	defer s.rmu.RUnlock()
	ssm, ok := s.replicas[shardID]
	return ssm, ok
}

// TransferLeadership transfers the leader to the chosen transferee.
func (s *OliveServer) TransferLeadership() error {
	// TODO: TransferLeadership
	return nil
}

func (s *OliveServer) ShardView(shard uint64) {
	lg := s.Logger()
	opt := dragonboat.DefaultNodeHostInfoOption
	opt.SkipLogInfo = true
	nhi := s.nh.GetNodeHostInfo(opt)
	lg.Sugar().Infof("%+v", nhi)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	ms, e := s.nh.SyncGetShardMembership(ctx, shard)
	if e != nil {

	}
	lg.Sugar().Infof("%+v", ms)
}
