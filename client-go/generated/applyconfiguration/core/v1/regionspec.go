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

// RegionSpecApplyConfiguration represents a declarative configuration of the RegionSpec type for use
// with apply.
type RegionSpecApplyConfiguration struct {
	Id               *int64                            `json:"id,omitempty"`
	DeploymentId     *int64                            `json:"deploymentId,omitempty"`
	InitialReplicas  []RegionReplicaApplyConfiguration `json:"initialReplicas,omitempty"`
	ElectionRTT      *int64                            `json:"electionRTT,omitempty"`
	HeartbeatRTT     *int64                            `json:"heartbeatRTT,omitempty"`
	Leader           *int64                            `json:"leader,omitempty"`
	DefinitionsLimit *int64                            `json:"definitionsLimit,omitempty"`
}

// RegionSpecApplyConfiguration constructs a declarative configuration of the RegionSpec type for use with
// apply.
func RegionSpec() *RegionSpecApplyConfiguration {
	return &RegionSpecApplyConfiguration{}
}

// WithId sets the Id field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Id field is set to the value of the last call.
func (b *RegionSpecApplyConfiguration) WithId(value int64) *RegionSpecApplyConfiguration {
	b.Id = &value
	return b
}

// WithDeploymentId sets the DeploymentId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeploymentId field is set to the value of the last call.
func (b *RegionSpecApplyConfiguration) WithDeploymentId(value int64) *RegionSpecApplyConfiguration {
	b.DeploymentId = &value
	return b
}

// WithInitialReplicas adds the given value to the InitialReplicas field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InitialReplicas field.
func (b *RegionSpecApplyConfiguration) WithInitialReplicas(values ...*RegionReplicaApplyConfiguration) *RegionSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithInitialReplicas")
		}
		b.InitialReplicas = append(b.InitialReplicas, *values[i])
	}
	return b
}

// WithElectionRTT sets the ElectionRTT field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ElectionRTT field is set to the value of the last call.
func (b *RegionSpecApplyConfiguration) WithElectionRTT(value int64) *RegionSpecApplyConfiguration {
	b.ElectionRTT = &value
	return b
}

// WithHeartbeatRTT sets the HeartbeatRTT field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeartbeatRTT field is set to the value of the last call.
func (b *RegionSpecApplyConfiguration) WithHeartbeatRTT(value int64) *RegionSpecApplyConfiguration {
	b.HeartbeatRTT = &value
	return b
}

// WithLeader sets the Leader field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Leader field is set to the value of the last call.
func (b *RegionSpecApplyConfiguration) WithLeader(value int64) *RegionSpecApplyConfiguration {
	b.Leader = &value
	return b
}

// WithDefinitionsLimit sets the DefinitionsLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefinitionsLimit field is set to the value of the last call.
func (b *RegionSpecApplyConfiguration) WithDefinitionsLimit(value int64) *RegionSpecApplyConfiguration {
	b.DefinitionsLimit = &value
	return b
}
