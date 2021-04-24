// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package golangcilint provides a task for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module command.
// "golangci-lint" a fast, parallel runner for dozens of Go linters that uses caching, supports YAML configurations and
// has integrations with all major IDEs.
//
// See https://pkg.go.dev/github.com/golangci/golangci-lint for more details about "golangci-lint".
// The source code of "golangci-lint" is available at
// https://github.com/golangci/golangci-lint/tree/master/cmd/golangci-lint.
package golangcilint

import (
	"fmt"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module command.
type Task struct {
	opts *Options
}

// BuildParams builds the parameters.
func (t *Task) BuildParams() []string {
	return t.opts.args
}

// Env returns the task specific environment.
func (t *Task) Env() map[string]string {
	return t.opts.env
}

// ExecName returns the executable name.
func (t *Task) ExecName() string {
	return t.opts.execName
}

// ID returns the identifier of the Go module.
func (t *Task) ID() *project.GoModuleID {
	return t.opts.goModule
}

// Kind returns the task kind.
func (t *Task) Kind() task.Kind {
	return task.KindGoModule
}

// Name returns the task name.
func (t *Task) Name() string {
	return t.opts.name
}

// Options returns the task options.
func (t *Task) Options() task.Options {
	return *t.opts
}

// New creates a new task for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module command.
// If no extra arguments are configured, DefaultArgs are passed to the command.
func New(opts ...Option) (*Task, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, fmt.Errorf("create %q task options: %w", taskName, optErr)
	}
	return &Task{opts: opt}, nil
}
