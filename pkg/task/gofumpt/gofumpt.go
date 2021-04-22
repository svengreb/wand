// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package gofumpt provides a task for the "mvdan.cc/gofumpt" Go module command.
// "gofumpt" enforce a stricter format than "https://pkg.go.dev/cmd/gofmt" and provides additional rules, while being
// backwards compatible.
// It is a modified fork of "https://pkg.go.dev/cmd/gofmt" so it can be used as a drop-in replacement.
//
// See https://pkg.go.dev/mvdan.cc/gofumpt for more details about "gofumpt".
// The source code of "gofumpt" is available at https://github.com/mvdan/gofumpt.
// See https://github.com/mvdan/gofumpt#added-rules for more details about available rules.
package gofumpt

import (
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task for the "mvdan.cc/gofumpt" Go module command.
// "gofumpt" enforce a stricter format than "https://pkg.go.dev/cmd/gofmt", while being backwards compatible,
// and provides additional rules.
// It is a modified fork of "https://pkg.go.dev/cmd/gofmt" so it can be used as a drop-in replacement.
//
// See https://pkg.go.dev/mvdan.cc/gofumpt for more details about "gofumpt".
// The source code of "gofumpt" is available at https://github.com/mvdan/gofumpt.
// See https://github.com/mvdan/gofumpt#added-rules for more details about available rules.
type Task struct {
	ac   app.Config
	opts *Options
}

// BuildParams builds the parameters.
func (t *Task) BuildParams() []string {
	var params []string

	// Enable gofumpt's extra rules.
	if t.opts.extraRules {
		params = append(params, "-extra")
	}

	// List files whose formatting are non-compliant to gofumpt's styles.
	if t.opts.listNonCompliantFiles {
		params = append(params, "-l")
	}

	// Write result to source files instead of stdout.
	if t.opts.persistChanges {
		params = append(params, "-w")
	}

	// Report all errors and not just the first 10 on different lines.
	if t.opts.reportAllErrors {
		params = append(params, "-e")
	}

	if t.opts.simplify {
		params = append(params, "-s")
	}

	// Include additionally configured arguments.
	params = append(params, t.opts.extraArgs...)

	// Only search in specified paths for Go source files...
	if len(t.opts.paths) > 0 {
		params = append(params, t.opts.paths...)
	} else {
		// ...or otherwise search recursively starting from the working directory of the current process.
		params = append(params, ".")
	}

	return params
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

// Options returns the task options.
func (t *Task) Options() task.Options {
	return *t.opts
}

// New creates a new task for the "mvdan.cc/gofumpt" Go module command.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(ac app.Config, opts ...Option) *Task {
	return &Task{ac: ac, opts: NewOptions(opts...)}
}
