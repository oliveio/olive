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

package raft

import (
	"context"
	"fmt"

	pb "github.com/olive-io/olive/api/olivepb"
	"github.com/olive-io/olive/pkg/bytesutil"
	"go.uber.org/zap"
)

var (
	definitionPrefix = []byte("definitions")
	processPrefix    = []byte("processes")
)

func (r *Region) deployDefinition(ctx context.Context, definition *pb.Definition) error {
	isLeader, err := r.readyLeader(ctx)
	if err != nil {
		return err
	}
	if !isLeader {
		return nil
	}

	prefix := bytesutil.PathJoin(definitionPrefix, []byte(definition.Id))
	key := bytesutil.PathJoin(prefix, []byte(fmt.Sprintf("%d", definition.Version)))

	if kvs, _ := r.getRange(prefix, nil, 0); len(kvs) == 0 {
		r.metric.definition.Add(1)
	}

	data, _ := definition.Marshal()
	r.put(key, data, true)

	return nil
}

func (r *Region) executeDefinition(ctx context.Context, process *pb.ProcessInstance) error {
	isLeader, err := r.readyLeader(ctx)
	if err != nil {
		return err
	}
	if !isLeader {
		return nil
	}

	prefix := bytesutil.PathJoin(processPrefix,
		[]byte(process.DefinitionId),
		[]byte(fmt.Sprintf("%d", process.DefinitionVersion)))
	key := bytesutil.PathJoin(prefix, []byte(fmt.Sprintf("%d", process.Id)))

	v, _ := r.get(key)
	if v != nil {
		r.lg.Warn("process-instance be executed",
			zap.String("definition", process.DefinitionId),
			zap.Uint64("version", process.DefinitionVersion),
			zap.Uint64("id", process.Id))
		return nil
	}

	process.Status = pb.ProcessInstance_Prepare
	data, _ := process.Marshal()
	r.put(key, data, true)

	return nil
}
