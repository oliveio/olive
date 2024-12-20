/*
Copyright 2024 The olive Authors

This program is offered under a commercial and under the AGPL license.
For AGPL licensing, see below.

AGPL licensing:
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package gather

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	corev1 "github.com/olive-io/olive/apis/core/v1"
	"github.com/olive-io/olive/apis/version"
	"github.com/olive-io/olive/runner/metrics"
	"github.com/olive-io/olive/runner/storage/backend"
)

const runnerPrefix = "/runner"

type Config struct {
	Name        string
	ListenURL   string
	HeartbeatMs int64
}

// Gather gets the information of the runner.
type Gather struct {
	ctx context.Context
	cfg *Config
	be  backend.IBackend

	runner *corev1.Runner
}

func NewGather(ctx context.Context, cfg *Config, be backend.IBackend) (*Gather, error) {
	runner := new(corev1.Runner)
	resp, err := be.Get(ctx, runnerPrefix)
	if err != nil {
		if !errors.Is(err, backend.ErrNotFound) {
			return nil, err
		}

		runner = &corev1.Runner{}
		runner.SetName(cfg.Name)

	} else {
		if err = resp.Unmarshal(runner); err != nil {
			return nil, err
		}
	}

	cpuTotal := uint64(0)
	cpus, err := cpu.Counts(false)
	if err != nil {
		return nil, fmt.Errorf("read system cpu: %w", err)
	}
	cpuInfos, _ := cpu.Info()
	if len(cpuInfos) > 0 {
		cpuTotal = uint64(cpus) * uint64(cpuInfos[0].Mhz)
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("read system memory: %w", err)
	}

	runner.Spec.ListenURL = cfg.ListenURL
	runner.Spec.HeartbeatMs = cfg.HeartbeatMs
	runner.Spec.Hostname, _ = os.Hostname()
	runner.Spec.Features = metrics.GetFeatures()
	runner.Spec.Version = version.Version
	runner.Status.CpuTotal = float64(cpuTotal)
	runner.Status.MemoryTotal = float64(vm.Total)

	if err = be.Put(ctx, runnerPrefix, runner); err != nil {
		return nil, err
	}

	delegate := &Gather{
		ctx: ctx,
		cfg: cfg,
		be:  be,

		runner: runner,
	}

	return delegate, nil
}

func (d *Gather) GetRunner() *corev1.Runner {
	return d.runner.DeepCopy()
}

func (d *Gather) GetStat() (*corev1.RunnerStatistics, error) {
	rs := &corev1.RunnerStatistics{
		Name: d.runner.Name,
		BpmnStat: &corev1.BpmnStatistics{
			Processes: int64(metrics.ProcessCounter.Get()),
			Events:    int64(metrics.EventCounter.Get()),
			Tasks:     int64(metrics.TaskCounter.Get()),
		},
	}
	interval := time.Millisecond * 300
	percents, err := cpu.Percent(interval, false)
	if err != nil {
		return nil, fmt.Errorf("current cpu percent: %w", err)
	}
	if len(percents) > 0 {
		rs.CpuUsed = percents[0] * d.runner.Status.CpuTotal
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("current memory percent: %w", err)
	}
	if vm != nil {
		rs.MemoryUsed = vm.UsedPercent * d.runner.Status.MemoryTotal
	}

	return rs, nil
}
