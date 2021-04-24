// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package env provides a task for the Go toolchain `env` command.
// See `go help environment`, `go help env` and the `go` command documentations at
// https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information and
// https://pkg.go.dev/cmd/go#hdr-Environment_variables for more details.
//
// References
//
//   (1) https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information
//   (2) https://pkg.go.dev/cmd/go#hdr-Environment_variables
//   (3) https://pkg.go.dev/cmd/go/internal/envcmd
package env

import (
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task for the Go toolchain `env` command.
// See `go help environment`, `go help env` and the `go` command documentations at
// https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information and
// https://pkg.go.dev/cmd/go#hdr-Environment_variables for more details.
//
// References
//
//   (1) https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information
//   (2) https://pkg.go.dev/cmd/go#hdr-Environment_variables
//   (3) https://pkg.go.dev/cmd/go/internal/envcmd
type Task struct {
	opts *Options
}

// BuildParams builds the parameters.
func (t *Task) BuildParams() []string {
	params := []string{"env"}

	// Enable JSON output format.
	if t.opts.EnableJSONOutput {
		params = append(params, "-json")
	}

	// Include additionally configured arguments.
	params = append(params, t.opts.extraArgs...)

	return append(params, t.opts.EnvVars...)
}

// Env returns the task specific environment.
func (t *Task) Env() map[string]string {
	return t.opts.env
}

// Kind returns the task kind.
func (t *Task) Kind() task.Kind {
	return task.KindExec
}

// Name returns the task name.
func (t *Task) Name() string {
	return t.opts.name
}

// Options returns the task options.
func (t *Task) Options() task.Options {
	return *t.opts
}

// New creates a new task for the Go toolchain `env` command.
func New(opts ...Option) *Task {
	return &Task{opts: NewOptions(opts...)}
}
