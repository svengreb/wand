// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package goimports

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/cast"
	"github.com/svengreb/wand/pkg/project"
)

const (
	// DefaultGoModulePath is the default "goimports" module command import path.
	DefaultGoModulePath = "golang.org/x/tools/cmd/goimports"

	// DefaultGoModuleVersion is the default "goimports" module version.
	DefaultGoModuleVersion = "latest"
)

// Options are spell incantation options for the "golang.org/x/tools/cmd/goimports" Go module command.
type Options struct {
	// env are spell incantation specific environment variables.
	env map[string]string

	// extraArgs are additional arguments to pass to the "goimports" command.
	extraArgs []string

	// goModule are partial Go module identifier information.
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

// Option is a spell incantation option for the "golang.org/x/tools/cmd/goimports" Go module command.
type Option func(*Options)

// WithEnv sets the spell incantation specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.env = env
	}
}

// WithExtraArgs sets additional arguments to pass to the "goimports" module command.
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

// WithModulePath sets the "goimports" module import path.
// Defaults to DefaultGoModulePath.
func WithModulePath(path string) Option {
	return func(o *Options) {
		if path != "" {
			o.goModule.Path = path
		}
	}
}

// WithModuleVersion sets the "goimports" module version.
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

// NewOptions creates new spell incantation options for the "golang.org/x/tools/cmd/goimports" Go module command.
func NewOptions(opts ...Option) (*Options, error) {
	version, versionErr := semver.NewVersion(DefaultGoModuleVersion)
	if versionErr != nil {
		return nil, &cast.ErrCast{
			Err:  fmt.Errorf("parsing default module version %q: %w", DefaultGoModulePath, versionErr),
			Kind: cast.ErrCasterInvalidOpts,
		}
	}
	opt := &Options{
		env: make(map[string]string),
		goModule: &project.GoModuleID{
			Path:     DefaultGoModulePath,
			Version:  version,
			IsLatest: true,
		},
	}
	for _, o := range opts {
		o(opt)
	}

	return opt, nil
}
