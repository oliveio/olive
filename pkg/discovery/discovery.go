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

package discovery

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	pb "github.com/olive-io/olive/api/discoverypb"
)

var (
	DefaultRegistryTimeout = time.Second * 10
	// DefaultNamespace the default value of namespace
	DefaultNamespace = "default"

	ErrNoDiscovery = errors.New("no discovery")
	// ErrNotFound not found error when GetService is called
	ErrNotFound = errors.New("service not found")
	// ErrWatcherStopped watcher stopped error when watcher is stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

// IDiscovery the registry provides an interface for service discovery
// and an abstraction over varying implementations
type IDiscovery interface {
	Options() Options
	Register(context.Context, *pb.Service, ...RegisterOption) error
	Deregister(context.Context, *pb.Service, ...DeregisterOption) error
	GetService(context.Context, string, ...GetOption) ([]*pb.Service, error)
	ListServices(context.Context, ...ListOption) ([]*pb.Service, error)
	Watch(context.Context, ...WatchOption) (Watcher, error)
}
