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

package app

import (
	"os"
	"os/signal"

	"github.com/olive-io/olive/meta"
	"github.com/olive-io/olive/pkg/signalutil"
	"github.com/olive-io/olive/pkg/version"
	"github.com/spf13/cobra"
)

func NewMetaCommand() *cobra.Command {
	app := &cobra.Command{
		Use:           "olive-meta",
		Short:         "a component of olive",
		Version:       version.Version,
		RunE:          runMeta,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	app.ResetFlags()
	flags := app.PersistentFlags()
	meta.AddFlagSet(flags)

	return app
}

func runMeta(cmd *cobra.Command, args []string) error {
	flags := cmd.PersistentFlags()

	cfg, err := meta.ConfigFromFlagSet(flags)
	if err != nil {
		return err
	}
	if err = cfg.Validate(); err != nil {
		return err
	}

	if err = setupMetaServer(cfg); err != nil {
		return err
	}

	return nil
}

func setupMetaServer(cfg meta.Config) error {
	ms, err := meta.NewServer(cfg)
	if err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signalutil.Shutdown()...)

	if err = ms.Start(); err != nil {
		return err
	}

	select {
	// wait on kill signal
	case <-ch:
	}

	ms.Stop()

	return nil
}
