// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

//go:build go1.17

// Package run provides a task for the Go toolchain "run" command.
// It requires at least Go version 1.17 which comes with support to [run commands in module-aware mode], by passing
// version suffixes to `go run` arguments, without affecting the `main` module and will not "pollute" the `go.mod`
// file.
// See the documentation about [how to compile and run Go programs] for more details.
//
// [run commands in module-aware mode]: https://go.dev/doc/go1.17#go%20run
// [how to compile and run Go programs]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
package run

import "github.com/svengreb/wand/pkg/task"

// Task is a task for the Go toolchain "run" command.
// It requires at least Go version 1.17 which comes with support to [run commands in module-aware mode], by passing
// version suffixes to `go run` arguments, without affecting the `main` module and will not "pollute" the `go.mod`
// file anymore.
// See the documentation about [how to compile and run Go programs] for more details.
//
// [run commands in module-aware mode]: https://go.dev/doc/go1.17#go%20run
// [how to compile and run Go programs]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
type Task struct {
	opts *Options
}

// BuildParams builds the parameters.
// Note that configured flags are applied after the "GOFLAGS" environment variable and could overwrite already defined
// flags.
//
// See the [Go command documentation about environment variables] and the builtin helps for more details:
//
//	go help environment
//	go help env
//
// [Go command documentation about environment variables]: https://golang.org/cmd/go/#hdr-Environment_variables
func (t *Task) BuildParams() []string {
	params := []string{"run"}
	params = append(params, t.opts.goModule.String())
	params = append(params, t.opts.args...)
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

// New creates a new task for the Go toolchain "run" command.
func New(opts ...Option) *Task {
	return &Task{opts: NewOptions(opts...)}
}
