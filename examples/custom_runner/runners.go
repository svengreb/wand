// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"

	"github.com/magefile/mage/sh"
	glFS "github.com/svengreb/golib/pkg/io/fs"

	"github.com/svengreb/wand/pkg/task"
)

const (
	// DefaultRunnerExec is the default name of the runner executable.
	DefaultRunnerExec = "fruitctl"

	// RunnerName is the name of the runner.
	RunnerName = "fruit_mixer"
)

// FruitMixerOption is a fruit mixer runner option.
type FruitMixerOption func(*FruitMixerOptions)

// FruitMixerOptions are fruit mixer runner options.
type FruitMixerOptions struct {
	// Env is the runner specific environment.
	Env map[string]string

	// Exec is the name or path of the runner command executable.
	Exec string

	// Quiet indicates whether the runner output should be minimal.
	Quiet bool
}

// FruitMixerRunner is a task runner for the fruit mixer.
type FruitMixerRunner struct {
	opts *FruitMixerOptions
}

// FilePath returns the path to the runner executable.
func (r *FruitMixerRunner) FilePath() string {
	return r.opts.Exec
}

// Handles returns the supported task kind.
func (r *FruitMixerRunner) Handles() task.Kind {
	return task.KindExec
}

// Run runs the command.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
func (r *FruitMixerRunner) Run(t task.Task) error {
	tExec, tErr := r.prepareTask(t)
	if tErr != nil {
		return tErr
	}

	if r.opts.Quiet {
		return sh.RunWith(r.opts.Env, r.opts.Exec, tExec.BuildParams()...)
	}
	return sh.RunWithV(r.opts.Env, r.opts.Exec, tExec.BuildParams()...)
}

// RunOut runs the command and returns its output.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
func (r *FruitMixerRunner) RunOut(t task.Task) (string, error) {
	tExec, tErr := r.prepareTask(t)
	if tErr != nil {
		return "", tErr
	}

	return sh.OutputWith(r.opts.Env, r.opts.Exec, tExec.BuildParams()...)
}

// Validate validates the command executable.
// It returns an error of type *task.ErrRunner when the executable does not exist and when it is also not available in
// the executable search path(s) of the current environment.
func (r *FruitMixerRunner) Validate() error {
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

// prepareTask checks if the given task is of type task.Exec and prepares the task specific environment.
func (r *FruitMixerRunner) prepareTask(t task.Task) (task.Exec, error) {
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

// NewFruitMixerRunner creates a new fruit mixer command runner.
func NewFruitMixerRunner(opts ...FruitMixerOption) *FruitMixerRunner {
	return &FruitMixerRunner{opts: NewFruitMixerRunnerOptions(opts...)}
}

// NewFruitMixerRunnerOptions creates new fruit mixer runner options.
func NewFruitMixerRunnerOptions(opts ...FruitMixerOption) *FruitMixerOptions {
	opt := &FruitMixerOptions{
		Env:  make(map[string]string),
		Exec: DefaultRunnerExec,
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// WithEnv sets the runner specific environment.
func WithEnv(env map[string]string) FruitMixerOption {
	return func(o *FruitMixerOptions) {
		o.Env = env
	}
}

// WithExec sets the name or path of the runner command executable.
// Defaults to DefaultRunnerExec.
func WithExec(nameOrPath string) FruitMixerOption {
	return func(o *FruitMixerOptions) {
		o.Exec = nameOrPath
	}
}

// WithQuiet indicates whether the runner output should be minimal.
func WithQuiet(quiet bool) FruitMixerOption {
	return func(o *FruitMixerOptions) {
		o.Quiet = quiet
	}
}
