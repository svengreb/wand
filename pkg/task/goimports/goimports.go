// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

// Package goimports provides a task for the "golang.org/x/tools/cmd/goimports" Go module command.
// "goimports" allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in
// the same style as "https://pkg.go.dev/cmd/gofmt" so it can be used as a replacement.
//
// See https://pkg.go.dev/golang.org/x/tools/cmd/goimports for more details about "goimports".
// The source code of "goimports" is available at https://github.com/golang/tools/tree/master/cmd/goimports.
package goimports

import (
	"fmt"
	"strings"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task for the "golang.org/x/tools/cmd/goimports" Go module command.
// "goimports" allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in
// the same style as "https://pkg.go.dev/cmd/gofmt" so it can be used as a replacement.
//
// See https://pkg.go.dev/golang.org/x/tools/cmd/goimports for more details about "goimports".
// The source code of "goimports" is available at https://github.com/golang/tools/tree/master/cmd/goimports.
type Task struct {
	opts *Options
}

// BuildParams builds the parameters.
func (t *Task) BuildParams() []string {
	var params []string

	// List files whose formatting are non-compliant to the style guide.
	if t.opts.listNonCompliantFiles {
		params = append(params, "-l")
	}

	// A comma-separated list of prefixes for local package imports to be put after 3rd-party packages.
	if len(t.opts.localPkgs) > 0 {
		params = append(params, "-local", fmt.Sprintf("'%s'", strings.Join(t.opts.localPkgs, ",")))
	}

	// Report all errors and not just the first 10 on different lines.
	if t.opts.reportAllErrors {
		params = append(params, "-e")
	}

	// Write result to source files instead of stdout.
	if t.opts.persistChanges {
		params = append(params, "-w")
	}

	// Toggle verbose output.
	if t.opts.verbose {
		params = append(params, "-v")
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

// Name returns the task name.
func (t *Task) Name() string {
	return t.opts.name
}

// Options returns the task options.
func (t *Task) Options() task.Options {
	return *t.opts
}

// New creates a new task for the "golang.org/x/tools/cmd/goimports" Go module command.
func New(opts ...Option) (*Task, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, fmt.Errorf("create %q task options: %w", taskName, optErr)
	}
	return &Task{opts: opt}, nil
}
