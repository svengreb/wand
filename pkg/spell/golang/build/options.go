// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package build

import (
	spellGo "github.com/svengreb/wand/pkg/spell/golang"
)

const (
	// DefaultDistOutputDirName is the default directory name for production and distribution builds.
	DefaultDistOutputDirName = "dist"
)

// Options are spell incantation options for the Go toolchain "build" command.
type Options struct {
	*spellGo.Options

	// BinaryArtifactName is the name for the binary build artifact.
	BinaryArtifactName string

	// CrossCompileTargetPlatforms are the names of cross-compile platform targets.
	// See `go tool dist list` and the `go` command documentations for more details:
	// - https://github.com/golang/go/blob/master/src/cmd/dist/build.go
	CrossCompileTargetPlatforms []string

	// Flags are additional flags to pass to the Go `build` command along with the base Go flags.
	// See `go help build` and the Go command documentation for more details:
	// - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	Flags []string

	// OutputDir is the output directory, relative to the project root, for compilation artifacts.
	OutputDir string

	// spellGoOpts are shared Go toolchain commands options.
	spellGoOpts []spellGo.Option
}

// Option is a spell incantation option for the Go toolchain "build" command.
type Option func(*Options)

// WithBinaryArtifactName sets the name for the binary build artifact.
func WithBinaryArtifactName(name string) Option {
	return func(o *Options) {
		o.BinaryArtifactName = name
	}
}

// WithCrossCompileTargetPlatforms sets the names of cross-compile platform targets.
func WithCrossCompileTargetPlatforms(platforms ...string) Option {
	return func(o *Options) {
		o.CrossCompileTargetPlatforms = append(o.CrossCompileTargetPlatforms, platforms...)
	}
}

// WithFlags sets additional flags to pass to the Go `build` command along with the base Go flags.
func WithFlags(flags ...string) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}

// WithGoOptions sets shared Go toolchain commands options.
func WithGoOptions(goOpts ...spellGo.Option) Option {
	return func(o *Options) {
		o.spellGoOpts = append(o.spellGoOpts, goOpts...)
	}
}

// WithOutputDir sets the output directory, relative to the project root, for compilation artifacts.
func WithOutputDir(dir string) Option {
	return func(o *Options) {
		o.OutputDir = dir
	}
}

// NewOptions creates new spell incantation options for the Go toolchain "build" command.
func NewOptions(opts ...Option) *Options {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}

	opt.Options = spellGo.NewOptions(opt.spellGoOpts...)

	return opt
}
