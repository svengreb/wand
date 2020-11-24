// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package golangcilint

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/cast"
	"github.com/svengreb/wand/pkg/project"
)

const (
	// DefaultGoModulePath is the default "goimports" module command import path.
	DefaultGoModulePath = "github.com/golangci/golangci-lint/cmd/golangci-lint"

	// DefaultGoModuleVersion is the default "goimports" module version.
	DefaultGoModuleVersion = "v1.32.0"
)

// DefaultArgs are the default arguments to pass to the "golangci-lint" command.
var DefaultArgs = []string{"run"}

// Options are spell incantation options for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module
// command.
type Options struct {
	// args are arguments to pass to the "golangci-lint" command.
	args []string

	// env are spell incantation specific environment variables.
	env map[string]string

	// goModule are partial Go module identifier information.
	goModule *project.GoModuleID

	// verbose indicates whether the output should be verbose.
	verbose bool
}

// Option is a spell incantation option for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module command.
type Option func(*Options)

// WithArgs sets additional arguments to pass to the "golangci-lint" module command.
// By default DefaultArgs are passed.
func WithArgs(args ...string) Option {
	return func(o *Options) {
		o.args = append(o.args, args...)
	}
}

// WithEnv sets the spell incantation specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.env = env
	}
}

// WithModulePath sets the "golangci-lint" module command import path.
// Defaults to DefaultGoModulePath.
func WithModulePath(path string) Option {
	return func(o *Options) {
		if path != "" {
			o.goModule.Path = path
		}
	}
}

// WithModuleVersion sets the "golangci-lint" module version.
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

// newOptions creates new spell incantation options for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go
// module command.
func newOptions(opts ...Option) (*Options, error) {
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
			Path:    DefaultGoModulePath,
			Version: version,
		},
	}
	for _, o := range opts {
		o(opt)
	}

	if len(opt.args) == 0 {
		opt.args = append(opt.args, DefaultArgs...)
	}

	return opt, nil
}
