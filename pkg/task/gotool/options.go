// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package gotool

import (
	"fmt"
	"path/filepath"
)

const (
	// RunnerName is the name of the runner.
	RunnerName = "gotool"
)

// RunnerOption is a runner option.
type RunnerOption func(*RunnerOptions)

// RunnerOptions are runner options.
type RunnerOptions struct {
	// Env is the runner specific environment.
	Env map[string]string

	// toolsBinDir is the path to the directory where compiled executables of Go module-based "main" packages are placed.
	toolsBinDir string

	// Quiet indicates whether the runner output should be minimal.
	Quiet bool
}

// NewRunnerOptions creates new runner options.
func NewRunnerOptions(opts ...RunnerOption) (*RunnerOptions, error) {
	opt := &RunnerOptions{
		Env: make(map[string]string),
	}
	for _, o := range opts {
		o(opt)
	}

	if !filepath.IsAbs(opt.toolsBinDir) {
		return nil, fmt.Errorf("expect an absolute path for tool binaries directory, but got %q", opt.toolsBinDir)
	}

	return opt, nil
}

// WithEnv sets the runner specific environment.
func WithEnv(env map[string]string) RunnerOption {
	return func(o *RunnerOptions) {
		o.Env = env
	}
}

// WithToolsBinDir sets the path to the directory where compiled binaries of Go module-based tools are placed.
// Defaults to DefaultToolsBinDir.
func WithToolsBinDir(toolsBinDir string) RunnerOption {
	return func(o *RunnerOptions) {
		o.toolsBinDir = toolsBinDir
	}
}

// WithQuiet indicates whether the runner output should be minimal.
func WithQuiet(quiet bool) RunnerOption {
	return func(o *RunnerOptions) {
		o.Quiet = quiet
	}
}
