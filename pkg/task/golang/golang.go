// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

//go:build go1.17

// Package golang provides Go toolchain tasks and runner.
// Note that this package requires at least Go 1.17 due to the usage of specific Go command features and behaviors that
// are required for tasks like [github.com/svengreb/wand/pkg/task/golang/run]!
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
	tExec, tErr := r.prepareTask(t)
	if tErr != nil {
		return fmt.Errorf("runner %q: %w", RunnerName, tErr)
	}

	if r.opts.Quiet {
		if err := sh.RunWith(r.opts.Env, r.opts.Exec, tExec.BuildParams()...); err != nil {
			return &task.ErrRunner{
				Err:  fmt.Errorf("run task %q: %w", t.Name(), err),
				Kind: task.ErrRun,
			}
		}
		return nil
	}
	if err := sh.RunWithV(r.opts.Env, r.opts.Exec, tExec.BuildParams()...); err != nil {
		return &task.ErrRunner{
			Err:  fmt.Errorf("run task %q: %w", t.Name(), err),
			Kind: task.ErrRun,
		}
	}
	return nil
}

// RunOut runs the command and returns its output.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
func (r *Runner) RunOut(t task.Task) (string, error) {
	tExec, tErr := r.prepareTask(t)
	if tErr != nil {
		return "", fmt.Errorf("runner %q: %w", RunnerName, tErr)
	}

	out, runErr := sh.OutputWith(r.opts.Env, r.opts.Exec, tExec.BuildParams()...)
	if runErr != nil {
		return "", &task.ErrRunner{
			Err:  fmt.Errorf("run task %q: %w", t.Name(), runErr),
			Kind: task.ErrRun,
		}
	}
	return out, nil
}

// Validate validates the command executable.
// It returns an error of type *task.ErrRunner when the executable does not exist and when it is also not available in
// the executable search path(s) of the current environment.
func (r *Runner) Validate() error {
	// Check if the executable exists,...
	execExits, fsErr := glFS.RegularFileExists(r.opts.Exec)
	if fsErr != nil {
		return &task.ErrRunner{
			Err:  fmt.Errorf("runner %q: %w", RunnerName, fsErr),
			Kind: task.ErrRunnerValidation,
		}
	}
	// ...otherwise try to look up the executable search path(s).
	if !execExits {
		path, pathErr := exec.LookPath(r.opts.Exec)
		if pathErr != nil {
			return &task.ErrRunner{
				Err:  fmt.Errorf("runner %q: %q not found in PATH: %w", RunnerName, r.opts.Exec, pathErr),
				Kind: task.ErrRunnerValidation,
			}
		}
		r.opts.Exec = path
	}

	return nil
}

// prepareTask checks if the given task is of type task.Exec and prepares the task specific environment.
// It returns an error of type *task.ErrRunner when any error occurs during the execution.
func (r *Runner) prepareTask(t task.Task) (task.Exec, error) {
	tExec, ok := t.(task.Exec)
	if t.Kind() != task.KindExec || !ok {
		return nil, &task.ErrRunner{
			Err:  fmt.Errorf("expected %q but got %q", r.Handles(), t.Kind()),
			Kind: task.ErrUnsupportedTaskKind,
		}
	}

	for k, v := range tExec.Env() {
		r.opts.Env[k] = v
	}

	return tExec, nil
}

// NewRunner creates a new Go toolchain command runner.
func NewRunner(opts ...RunnerOption) *Runner {
	return &Runner{opts: NewRunnerOptions(opts...)}
}
