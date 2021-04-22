// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package test provides a task for the Go toolchain "test" command.
package test

import (
	"fmt"
	"path/filepath"

	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/task"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

// Task is a task for the Go toolchain "test" command.
type Task struct {
	ac   app.Config
	opts *Options
}

// BuildParams builds the parameters.
// Note that configured flags are applied after the "GOFLAGS" environment variable and could overwrite already defined
// flags. In addition, the output directory for test artifacts like profiles and reports must exist or must be be
// created before, otherwise the "test" Go toolchain command will fail to run.
//
// See `go help environment`, `go help env` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Environment_variables
func (t *Task) BuildParams() []string {
	params := []string{"test"}

	params = append(params, taskGo.BuildGoOptions(t.opts.taskGoOpts...)...)

	if t.opts.EnableVerboseOutput {
		params = append(params, "-v")
	}

	if t.opts.DisableCache {
		params = append(params, "-count=1")
	}

	if t.opts.EnableBlockProfile {
		params = append(params,
			fmt.Sprintf(
				"-blockprofile=%s",
				filepath.Join(t.opts.OutputDir, t.opts.BlockProfileOutputFileName),
			),
		)
	}

	if t.opts.EnableCoverageProfile {
		params = append(params,
			fmt.Sprintf(
				"-coverprofile=%s",
				filepath.Join(t.opts.OutputDir, t.opts.CoverageProfileOutputFileName),
			),
		)
	}

	if t.opts.EnableCPUProfile {
		params = append(params,
			fmt.Sprintf("-cpuprofile=%s",
				filepath.Join(t.opts.OutputDir, t.opts.CPUProfileOutputFileName),
			),
		)
	}

	if t.opts.EnableMemoryProfile {
		params = append(params,
			fmt.Sprintf("-memprofile=%s",
				filepath.Join(t.opts.OutputDir, t.opts.MemoryProfileOutputFileName),
			),
		)
	}

	if t.opts.EnableMutexProfile {
		params = append(params,
			fmt.Sprintf("-mutexprofile=%s",
				filepath.Join(t.opts.OutputDir, t.opts.MutexProfileOutputFileName),
			),
		)
	}

	if t.opts.EnableTraceProfile {
		params = append(params,
			fmt.Sprintf("-trace=%s",
				filepath.Join(t.opts.OutputDir, t.opts.TraceProfileOutputFileName),
			),
		)
	}

	if len(t.opts.Flags) > 0 {
		params = append(params, t.opts.Flags...)
	}

	params = append(params, t.opts.Pkgs...)

	return params
}

// Env returns the task specific environment.
func (t *Task) Env() map[string]string {
	return t.opts.Env
}

// Kind returns the task kind.
func (t *Task) Kind() task.Kind {
	return task.KindExec
}

// Options returns the task options.
func (t *Task) Options() task.Options {
	return *t.opts
}

// New creates a new task for the Go toolchain "test" command.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(ac app.Config, opts ...Option) *Task {
	opt := NewOptions(opts...)

	// Store test profiles and reports within the application specific subdirectory.
	if opt.OutputDir == "" {
		opt.OutputDir = filepath.Join(ac.BaseOutputDir, DefaultOutputDirName)
	}

	return &Task{ac: ac, opts: opt}
}
