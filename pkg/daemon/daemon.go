/*
Copyright 2023 The olive Authors

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

package daemon

import (
	"net/http"
	"strings"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type IDaemon interface {
	// StopNotify returns a channel that receives an empty struct
	// when the server is stopped.
	StopNotify() <-chan struct{}
	// StoppingNotify returns a channel that receives an empty struct
	// when the server is being stopped.
	StoppingNotify() <-chan struct{}
	// GoAttach creates a goroutine on a given function and tracks it using the waitgroup.
	// The passed function should interrupt on s.StoppingNotify().
	GoAttach(fn func())
	// OnDestroy registry destroy function that running when the server stop
	OnDestroy(fn func())
	// Shutdown sends signal to stop channel and all goroutines stop
	Shutdown()
}

type embedDaemon struct {
	lg *zap.Logger

	stopping chan struct{}
	done     chan struct{}
	stop     chan struct{}

	wgMu sync.RWMutex
	wg   sync.WaitGroup
}

func NewEmbedDaemon(lg *zap.Logger) IDaemon {
	s := &embedDaemon{
		lg:       lg,
		stopping: make(chan struct{}, 1),
		done:     make(chan struct{}, 1),
		stop:     make(chan struct{}, 1),
		wgMu:     sync.RWMutex{},
		wg:       sync.WaitGroup{},
	}

	return s
}

func (s *embedDaemon) StopNotify() <-chan struct{} { return s.done }

func (s *embedDaemon) StoppingNotify() <-chan struct{} { return s.stopping }

func (s *embedDaemon) GoAttach(fn func()) {
	s.wgMu.RLock() // this blocks with ongoing close(s.stopping)
	defer s.wgMu.RUnlock()
	select {
	case <-s.stopping:
		s.lg.Warn("server has stopped; skipping GoAttach")
		return
	default:
	}

	// now safe to add since waitgroup wait has not started yet
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		fn()
	}()
}

func (s *embedDaemon) OnDestroy(fn func()) {
	go s.doDestroy(fn)
}

func (s *embedDaemon) doDestroy(fn func()) {
	defer func() {
		s.wgMu.Lock() // block concurrent waitgroup adds in GoAttach while stopping
		close(s.stopping)
		s.wgMu.Unlock()

		s.wg.Wait()

		// clean something
		s.lg.Debug("server has stopped, running destroy operations")
		fn()

		close(s.done)
	}()

	<-s.stop
}

func (s *embedDaemon) Shutdown() {
	select {
	case s.stop <- struct{}{}:
	case <-s.done:
		return
	}
	<-s.done
}

// HybridHandler returns a http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Given in gRPC docs.
func HybridHandler(gh *grpc.Server, hh http.Handler) http.Handler {
	h2s := &http2.Server{}
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			gh.ServeHTTP(w, r)
		} else {
			hh.ServeHTTP(w, r)
		}
	}), h2s)
}
