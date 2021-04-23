// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package gofumpt

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

const (
	// DefaultExecName is the default name of the module executable.
	DefaultExecName = "gofumpt"

	// DefaultGoModulePath is the default module import path.
	DefaultGoModulePath = "mvdan.cc/gofumpt"

	// DefaultGoModuleVersion is the default Go module version of the runner command.
	DefaultGoModuleVersion = "v0.1.1"

	// taskName is the name of the task.
	taskName = "gofumpt"
)

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	// env is the task specific environment.
	env map[string]string

	// execName is the unique executable name.
	execName string

	// extraArgs are additional arguments passed to the command.
	extraArgs []string

	// extraRules indicates whether gofumpt's extra rules should be enabled.
	// See https://github.com/mvdan/gofumpt#added-rules for more details about available rules.
	extraRules bool

	// goModule is the Go module identifier.
	goModule *project.GoModuleID

	// listNonCompliantFiles indicates whether files, whose formatting are not conform to the style guide, should be
	// listed.
	listNonCompliantFiles bool

	// name is the task name.
	name string

	// paths are the paths to search for Go source files.
	// By default all directories are scanned recursively starting from the working directory of the current process.
	paths []string

	// persistChanges indicates whether results are written to the source files instead of standard output.
	persistChanges bool

	// reportAllErrors indicates whether all errors should be printed instead of only the first 10 on different lines.
	reportAllErrors bool

	// simplify indicates whether code should be simplified.
	simplify bool
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
		execName: DefaultExecName,
		goModule: &project.GoModuleID{
			Path:     DefaultGoModulePath,
			Version: version,
		},
		name: taskName,
	}
	for _, o := range opts {
		o(opt)
	}

	return opt, nil
}

// WithEnv sets the task specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.env = env
	}
}

// WithExecName sets the name of the executable.
func WithExecName(execName string) Option {
	return func(o *Options) {
		o.execName = execName
	}
}

// WithExtraArgs sets additional arguments to pass to the command.
func WithExtraArgs(extraArgs ...string) Option {
	return func(o *Options) {
		o.extraArgs = append(o.extraArgs, extraArgs...)
	}
}

// WithExtraRules indicates whether gofumpt's extra rules should be enabled.
// See https://github.com/mvdan/gofumpt#added-rules for more details about available rules.
func WithExtraRules(extraRules bool) Option {
	return func(o *Options) {
		o.extraRules = extraRules
	}
}

// WithListNonCompliantFiles indicates whether files, whose formatting are not conform to the style guide, are listed.
func WithListNonCompliantFiles(listNonCompliantFiles bool) Option {
	return func(o *Options) {
		o.listNonCompliantFiles = listNonCompliantFiles
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
func WithModuleVersion(version *semver.Version) Option {
	return func(o *Options) {
		if version != nil {
			o.goModule.Version = version
		}
	}
}

// WithPaths sets the paths to search for Go source files.
// By default all directories are scanned recursively starting from the working directory of the current process.
func WithPaths(paths ...string) Option {
	return func(o *Options) {
		o.paths = append(o.paths, paths...)
	}
}

// WithPersistedChanges indicates whether results are written to the source files instead of standard output.
func WithPersistedChanges(persistChanges bool) Option {
	return func(o *Options) {
		o.persistChanges = persistChanges
	}
}

// WithReportAllErrors indicates whether all errors should be printed instead of only the first 10 on different lines.
func WithReportAllErrors(reportAllErrors bool) Option {
	return func(o *Options) {
		o.reportAllErrors = reportAllErrors
	}
}

// WithSimplify indicates whether code should be simplified.
func WithSimplify(simplify bool) Option {
	return func(o *Options) {
		o.simplify = simplify
	}
}
