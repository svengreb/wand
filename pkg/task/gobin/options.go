// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package gobin

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

const (
	// DefaultGoModulePath is the default Go module import path of the runner command.
	DefaultGoModulePath = "github.com/myitcv/gobin"

	// DefaultGoModuleVersion is the default Go module version of the runner command.
	DefaultGoModuleVersion = "v0.0.14"

	// DefaultRunnerExec is the default path to the runner command executable.
	DefaultRunnerExec = "gobin"

	// RunnerName is the name of the runner.
	RunnerName = "gobin"
)

// RunnerOption is a runner option.
type RunnerOption func(*RunnerOptions)

// RunnerOptions are runner options.
type RunnerOptions struct {
	// Env is the runner specific environment.
	Env map[string]string

	// Exec is the name or path of the runner command executable.
	Exec string

	// goModule is the Go module identifier.
	goModule *project.GoModuleID

	// Quiet indicates whether the runner output should be minimal.
	Quiet bool
}

// NewRunnerOptions creates new runner options.
func NewRunnerOptions(opts ...RunnerOption) (*RunnerOptions, error) {
	version, versionErr := semver.NewVersion(DefaultGoModuleVersion)
	if versionErr != nil {
		return nil, &task.ErrRunner{
			Err:  fmt.Errorf("parsing default module version %q: %w", DefaultGoModulePath, versionErr),
			Kind: task.ErrInvalidRunnerOpts,
		}
	}

	opt := &RunnerOptions{
		Env:  make(map[string]string),
		Exec: DefaultRunnerExec,
		goModule: &project.GoModuleID{
			Path:    DefaultGoModulePath,
			Version: version,
		},
	}
	for _, o := range opts {
		o(opt)
	}

	return opt, nil
}

// WithEnv sets the runner specific environment.
func WithEnv(env map[string]string) RunnerOption {
	return func(o *RunnerOptions) {
		o.Env = env
	}
}

// WithExec sets the name or path of the runner command executable.
// Defaults to DefaultExecFileName.
func WithExec(nameOrPath string) RunnerOption {
	return func(o *RunnerOptions) {
		if nameOrPath != "" {
			o.Exec = nameOrPath
		}
	}
}

// WithModulePath sets the Go module import path of the runner command.
// Defaults to DefaultGoModulePath.
func WithModulePath(path string) RunnerOption {
	return func(o *RunnerOptions) {
		if path != "" {
			o.goModule.Path = path
		}
	}
}

// WithModuleVersion sets the Go module version of the runner command.
// Defaults to DefaultGoModuleVersion.
func WithModuleVersion(version *semver.Version) RunnerOption {
	return func(o *RunnerOptions) {
		if version != nil {
			o.goModule.Version = version
		}
	}
}

// WithQuiet indicates whether the runner output should be minimal.
func WithQuiet(quiet bool) RunnerOption {
	return func(o *RunnerOptions) {
		o.Quiet = quiet
	}
}
