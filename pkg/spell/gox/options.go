// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package gox

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/cast"
	"github.com/svengreb/wand/pkg/project"
	spellGo "github.com/svengreb/wand/pkg/spell/golang"
	spellGoBuild "github.com/svengreb/wand/pkg/spell/golang/build"
)

const (
	// DefaultGoModulePath is the default "gox" module command import path.
	DefaultGoModulePath = "github.com/mitchellh/gox"

	// DefaultGoModuleVersion is the default "gox" module version.
	DefaultGoModuleVersion = "v1.0.1"
)

var (
	// DefaultCrossCompileBinaryNameTemplate is the default name template for cross-compilation binary artifacts.
	DefaultCrossCompileBinaryNameTemplate = func(name string) string {
		return name + "-{{.OS}}-{{.Arch}}"
	}

	// DefaultCrossCompileTargetPlatforms are the names of default cross-compile platform targets.
	//
	// See `go tool dist list` and https://github.com/golang/go/blob/master/src/cmd/dist/build.go
	// for more details and a list of supported platforms.
	DefaultCrossCompileTargetPlatforms = []string{
		"darwin/amd64",
		"linux/amd64",
		"windows/amd64",
	}
)

// Options are spell incantation options for the "github.com/mitchellh/gox" Go module command.
type Options struct {
	*spellGoBuild.Options

	// env are spell incantation specific environment variables.
	env map[string]string

	// goCmd is the path to the Go toolchain executable.
	goCmd string

	// goModule are partial Go module identifier information.
	goModule *project.GoModuleID

	// outputTemplate is the name template for cross-compile platform targets.
	outputTemplate string

	// spellGoOpts are shared Go toolchain command options.
	spellGoOpts []spellGo.Option

	// spellGoOpts are options for the Go toolchain "build" command.
	spellGoBuildOpts []spellGoBuild.Option

	// verbose indicates whether the output should be verbose.
	verbose bool
}

// Option is a spell incantation option for the "github.com/mitchellh/gox" Go module command.
type Option func(*Options)

// WithEnv sets the spell incantation specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.env = env
	}
}

// WithGoCmd sets the path to the Go toolchain executable.
func WithGoCmd(goCmd string) Option {
	return func(o *Options) {
		o.goCmd = goCmd
	}
}

// WithOutputTemplate sets the name template for cross-compile platform targets.
// Defaults to DefaultCrossCompileBinaryNameTemplate.
func WithOutputTemplate(outputTemplate string) Option {
	return func(o *Options) {
		o.outputTemplate = outputTemplate
	}
}

// WithGoOptions sets shared Go toolchain command options.
func WithGoOptions(goOpts ...spellGo.Option) Option {
	return func(o *Options) {
		o.spellGoOpts = append(o.spellGoOpts, goOpts...)
	}
}

// WithGoBuildOptions sets options for the Go toolchain "build" command.
func WithGoBuildOptions(goBuildOpts ...spellGoBuild.Option) Option {
	return func(o *Options) {
		o.spellGoBuildOpts = append(o.spellGoBuildOpts, goBuildOpts...)
	}
}

// WithModulePath sets the "gox" module command import path.
// Defaults to DefaultGoModulePath.
func WithModulePath(path string) Option {
	return func(o *Options) {
		if path != "" {
			o.goModule.Path = path
		}
	}
}

// WithModuleVersion sets the "gox" module version.
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

// NewOptions creates new spell incantation options for the "github.com/mitchellh/gox" Go module command.
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
			Path:    DefaultGoModulePath,
			Version: version,
		},
	}
	for _, o := range opts {
		o(opt)
	}

	goBuildOpts := append(
		[]spellGoBuild.Option{spellGoBuild.WithGoOptions(opt.spellGoOpts...)},
		opt.spellGoBuildOpts...,
	)
	opt.Options = spellGoBuild.NewOptions(goBuildOpts...)

	if opt.outputTemplate == "" && opt.BinaryArtifactName != "" {
		opt.outputTemplate = DefaultCrossCompileBinaryNameTemplate(opt.BinaryArtifactName)
	}

	if len(opt.CrossCompileTargetPlatforms) == 0 {
		opt.CrossCompileTargetPlatforms = DefaultCrossCompileTargetPlatforms
	}

	return opt, nil
}
