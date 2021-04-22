// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package build provides a task for the Go toolchain "build" command.
package build

import (
	"path/filepath"

	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/task"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

// Task is a task for the Go toolchain "build" command.
type Task struct {
	ac   app.Config
	opts *Options
}

// BuildParams builds the parameters.
// Note that configured flags are applied after the "GOFLAGS" environment variable and could overwrite already defined
// flags.
//
// See `go help environment`, `go help env` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Environment_variables
func (t *Task) BuildParams() []string {
	params := []string{"build"}

	params = append(params, taskGo.BuildGoOptions(t.opts.taskGoOpts...)...)

	if len(t.opts.Flags) > 0 {
		params = append(params, t.opts.Flags...)
	}

	params = append(
		params,
		"-o",
		filepath.Join(t.opts.OutputDir, t.opts.BinaryArtifactName),
		t.ac.PkgImportPath,
	)

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

// New creates a new task for the Go toolchain "build" command.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(ac app.Config, opts ...Option) *Task {
	opt := NewOptions(opts...)

	if opt.BinaryArtifactName == "" {
		opt.BinaryArtifactName = ac.Name
	}

	// Store build artifacts in the application specific subdirectory.
	if opt.OutputDir == "" {
		opt.OutputDir = ac.BaseOutputDir
	}

	return &Task{ac: ac, opts: opt}
}
