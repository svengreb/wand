// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package goimports

import (
	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project"
)

const (
	// DefaultGoModulePath is the default module import path.
	DefaultGoModulePath = "golang.org/x/tools/cmd/goimports"
)

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	// env is the task specific environment.
	env map[string]string

	// extraArgs are additional arguments passed to the command.
	extraArgs []string

	// goModule is the Go module identifier.
	goModule *project.GoModuleID

	// listNonCompliantFiles indicates whether files, whose formatting are not conform to the style guide, should be
	// listed.
	listNonCompliantFiles bool

	// localPkgs are local packages whose imports will be placed after 3rd-party packages.
	localPkgs []string

	// paths are the paths to search for Go source files.
	// By default all directories are scanned recursively starting from the current working directory.
	paths []string

	// persistChanges indicates whether results are written to the source files instead of standard output.
	persistChanges bool

	// reportAllErrors indicates whether all errors should be printed instead of only the first 10 on different lines.
	reportAllErrors bool

	// verbose indicates whether the output should be verbose.
	verbose bool
}

// NewOptions creates new task options.
func NewOptions(opts ...Option) (*Options, error) {
	opt := &Options{
		env: make(map[string]string),
		goModule: &project.GoModuleID{
			Path:     DefaultGoModulePath,
			IsLatest: true,
		},
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

// WithExtraArgs sets additional arguments to pass to the command.
func WithExtraArgs(extraArgs ...string) Option {
	return func(o *Options) {
		o.extraArgs = append(o.extraArgs, extraArgs...)
	}
}

// WithListNonCompliantFiles indicates whether files, whose formatting are not conform to the style guide, are listed.
func WithListNonCompliantFiles(listNonCompliantFiles bool) Option {
	return func(o *Options) {
		o.listNonCompliantFiles = listNonCompliantFiles
	}
}

// WithLocalPkgs sets local packages whose imports will be placed after 3rd-party packages.
func WithLocalPkgs(localPkgs ...string) Option {
	return func(o *Options) {
		o.localPkgs = append(o.localPkgs, localPkgs...)
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

// WithPaths sets the paths to search for Go source files.
// By default all directories are scanned recursively starting from the current working directory.
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

// WithVerboseOutput indicates whether the output should be verbose.
func WithVerboseOutput(verbose bool) Option {
	return func(o *Options) {
		o.verbose = verbose
	}
}
