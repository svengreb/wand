// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package gomodupgrade provides a task for the "github.com/oligot/go-mod-upgrade" Go module command.
// "go-mod-upgrade" allows to update outdated Go module dependencies interactively.
//
// See https://pkg.go.dev/github.com/oligot/go-mod-upgrade for more details about "go-mod-upgrade".
// The source code of "go-mod-upgrade" is available at https://github.com/oligot/go-mod-upgrade.
package gomodupgrade

import (
	"fmt"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task for the "github.com/oligot/go-mod-upgrade" Go module command.
// "go-mod-upgrade" allows to update outdated Go module dependencies interactively.
//
// See https://pkg.go.dev/github.com/oligot/go-mod-upgrade for more details about "go-mod-upgrade".
// The source code of "go-mod-upgrade" is available at https://github.com/oligot/go-mod-upgrade.
type Task struct {
	opts *Options
}

// BuildParams builds the parameters.
func (t *Task) BuildParams() []string {
	return t.opts.extraArgs
}

// Env returns the task specific environment.
func (t *Task) Env() map[string]string {
	return t.opts.env
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

// New creates a new task for the "github.com/oligot/go-mod-upgrade" Go module command.
func New(opts ...Option) (*Task, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, fmt.Errorf("create %q task options: %w", taskName, optErr)
	}
	return &Task{opts: opt}, nil
}
