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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "github.com/olive-io/olive/apis/core/v1"
)

// RunnerStatusApplyConfiguration represents a declarative configuration of the RunnerStatus type for use
// with apply.
type RunnerStatusApplyConfiguration struct {
	Phase       *corev1.RunnerPhase                 `json:"phase,omitempty"`
	Message     *string                             `json:"message,omitempty"`
	CpuTotal    *float64                            `json:"cpuTotal,omitempty"`
	MemoryTotal *float64                            `json:"memoryTotal,omitempty"`
	DiskSize    *int64                              `json:"diskSize,omitempty"`
	Stat        *RunnerStatisticsApplyConfiguration `json:"stat,omitempty"`
}

// RunnerStatusApplyConfiguration constructs a declarative configuration of the RunnerStatus type for use with
// apply.
func RunnerStatus() *RunnerStatusApplyConfiguration {
	return &RunnerStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *RunnerStatusApplyConfiguration) WithPhase(value corev1.RunnerPhase) *RunnerStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *RunnerStatusApplyConfiguration) WithMessage(value string) *RunnerStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithCpuTotal sets the CpuTotal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CpuTotal field is set to the value of the last call.
func (b *RunnerStatusApplyConfiguration) WithCpuTotal(value float64) *RunnerStatusApplyConfiguration {
	b.CpuTotal = &value
	return b
}

// WithMemoryTotal sets the MemoryTotal field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MemoryTotal field is set to the value of the last call.
func (b *RunnerStatusApplyConfiguration) WithMemoryTotal(value float64) *RunnerStatusApplyConfiguration {
	b.MemoryTotal = &value
	return b
}

// WithDiskSize sets the DiskSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DiskSize field is set to the value of the last call.
func (b *RunnerStatusApplyConfiguration) WithDiskSize(value int64) *RunnerStatusApplyConfiguration {
	b.DiskSize = &value
	return b
}

// WithStat sets the Stat field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Stat field is set to the value of the last call.
func (b *RunnerStatusApplyConfiguration) WithStat(value *RunnerStatisticsApplyConfiguration) *RunnerStatusApplyConfiguration {
	b.Stat = value
	return b
}
