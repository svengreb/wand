// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package golang provides Go toolchain tasks and runner.
//
// See https://golang.org/cmd/go for more details.
package golang

import (
	"fmt"
	"os/exec"

	"github.com/magefile/mage/sh"
	glFS "github.com/svengreb/golib/pkg/io/fs"

	"github.com/svengreb/wand/pkg/task"
)

// Runner is a task runner for the Go toolchain.
type Runner struct {
	opts *RunnerOptions
}

// FilePath returns the path to the runner executable.
func (r *Runner) FilePath() string {
	return r.opts.Exec
}

// Handles returns the supported task kind.
func (r *Runner) Handles() task.Kind {
	return task.KindExec
}

// Run runs the command.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
func (r *Runner) Run(t task.Task) error {
	tExec, ok := t.(task.Exec)
	if t.Kind() != task.KindExec || !ok {
		return &task.ErrRunner{
			Err:  fmt.Errorf("expected %q but got %q", r.Handles(), t.Kind()),
			Kind: task.ErrUnsupportedTaskKind,
		}
	}

	runFn := sh.RunWithV
	if r.opts.Quiet {
		runFn = sh.RunWith
	}

	for k, v := range tExec.Env() {
		r.opts.Env[k] = v
	}

	return runFn(r.opts.Env, r.opts.Exec, tExec.BuildParams()...)
}

// Validate validates the command executable.
// It returns an error of type *task.ErrRunner when the executable does not exist and when it is also not available in
// the executable search path(s) of the current environment.
func (r *Runner) Validate() error {
	// Check if the executable exists,...
	execExits, fsErr := glFS.RegularFileExists(r.opts.Exec)
	if fsErr != nil {
		return &task.ErrRunner{
			Err:  fmt.Errorf("command runner %q: %w", RunnerName, fsErr),
			Kind: task.ErrRunnerValidation,
		}
	}
	// ...otherwise try to look up the executable search path(s).
	if !execExits {
		path, pathErr := exec.LookPath(r.opts.Exec)
		if pathErr != nil {
			return &task.ErrRunner{
				Err:  fmt.Errorf("command runner %q: %q not found in PATH: %w", RunnerName, r.opts.Exec, pathErr),
				Kind: task.ErrRunnerValidation,
			}
		}
		r.opts.Exec = path
	}

	return nil
}

// NewRunner creates a new Go toolchain command runner.
func NewRunner(opts ...RunnerOption) *Runner {
	return &Runner{opts: NewRunnerOptions(opts...)}
}
