// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package golangcilint

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

const (
	// DefaultGoModulePath is the default module import path.
	DefaultGoModulePath = "github.com/golangci/golangci-lint/cmd/golangci-lint"

	// DefaultGoModuleVersion is the default module version.
	DefaultGoModuleVersion = "v1.32.0"

	// taskName is the name of the task.
	taskName = "golangcilint"
)

// DefaultArgs are default arguments passed to the command.
var DefaultArgs = []string{"run"}

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	// args are arguments passed to the command.
	args []string

	// env is the task specific environment.
	env map[string]string

	// goModule is the Go module identifier.
	goModule *project.GoModuleID

	// name is the task name.
	name string

	// verbose indicates whether the output should be verbose.
	verbose bool
}

// NewOptions creates new task options.
func NewOptions(opts ...Option) (*Options, error) {
	version, versionErr := semver.NewVersion(DefaultGoModuleVersion)
	if versionErr != nil {
		return nil, &task.ErrTask{
			Err:  fmt.Errorf("parsing default module version %q: %w", DefaultGoModulePath, versionErr),
			Kind: task.ErrInvalidTaskOpts,
		}
	}

	opt := &Options{
		env: make(map[string]string),
		goModule: &project.GoModuleID{
			Path:    DefaultGoModulePath,
			Version: version,
		},
		name: taskName,
	}
	for _, o := range opts {
		o(opt)
	}

	if len(opt.args) == 0 {
		opt.args = append(opt.args, DefaultArgs...)
	}

	return opt, nil
}

// WithArgs sets additional arguments to pass to the command.
// Defaults to DefaultArgs.
func WithArgs(args ...string) Option {
	return func(o *Options) {
		o.args = append(o.args, args...)
	}
}

// WithEnv sets the task specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.env = env
	}
}

// WithModulePath sets the module import path.
// Defaults to DefaultGoModulePath.
func WithModulePath(path string) Option {
	return func(o *Options) {
		if path != "" {
			o.goModule.Path = path
		}
	}
}

// WithModuleVersion sets the module version.
// Defaults to DefaultGoModuleVersion.
func WithModuleVersion(version *semver.Version) Option {
	return func(o *Options) {
		if version != nil {
			o.goModule.Version = version
		}
	}
}

// WithVerboseOutput indicates whether the output should be verbose.
func WithVerboseOutput(verbose bool) Option {
	return func(o *Options) {
		o.verbose = verbose
	}
}
