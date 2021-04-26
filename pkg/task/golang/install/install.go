// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// +build go1.16

// Package install provides a task for the Go toolchain "install" command.
// It requires at least Go version 1.16 which comes with support to install commands via `go install` (1) without
// affecting the `main` module and will not "pollute" the `go.mod` file (2) anymore.
// See https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies for more details about the
// `go install` command.
//
// References
//
//   (1) https://blog.golang.org/go116-module-changes#TOC_4.
//   (2) https://blog.golang.org/go116-module-changes#TOC_3.
package install

import (
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task for the Go toolchain "install" command.
// It requires at least Go version 1.16 which comes with support to install commands via `go install` (1) without
// affecting the `main` module and will not "pollute" the `go.mod` file (2) anymore.
// See https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies for more details about the
// `go install` command.
//
// References
//
//   (1) https://blog.golang.org/go116-module-changes#TOC_4.
//   (2) https://blog.golang.org/go116-module-changes#TOC_3.
type Task struct {
	opts *Options
}

// BuildParams builds the parameters.
// Note that configured flags are applied after the "GOFLAGS" environment variable and could overwrite already defined
// flags.
//
// See `go help environment`, `go help env` and the `go` command documentations for more details:
// https://golang.org/cmd/go/#hdr-Environment_variables
func (t *Task) BuildParams() []string {
	params := []string{"install"}

	params = append(params, t.opts.goModule.String())

	return params
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

// New creates a new task for the Go toolchain "install" command.
func New(opts ...Option) *Task {
	return &Task{opts: NewOptions(opts...)}
}
