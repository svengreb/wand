// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package build

import (
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

const (
	// DefaultDistOutputDirName is the default directory name for production and distribution builds.
	DefaultDistOutputDirName = "dist"

	// taskName is the name of the task.
	taskName = "go/build"
)

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	*taskGo.Options

	// BinaryArtifactName is the name for the binary build artifact.
	BinaryArtifactName string

	// CrossCompileTargetPlatforms are the names of cross-compile platform targets.
	//
	// See `go tool dist list` and the `go` command documentations for more details:
	//   - https://github.com/golang/go/blob/master/src/cmd/dist/build.go
	CrossCompileTargetPlatforms []string

	// Flags are additional flags to pass to the Go `build` command along with the base Go flags.
	//
	// See `go help build` and the Go command documentation for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	Flags []string

	// name is the task name.
	name string

	// OutputDir is the output directory, relative to the project root, for compilation artifacts.
	OutputDir string

	// taskGoOpts are shared Go toolchain task options.
	taskGoOpts []taskGo.Option
}

// NewOptions creates new task options.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		name: taskName,
	}
	for _, o := range opts {
		o(opt)
	}

	opt.Options = taskGo.NewOptions(opt.taskGoOpts...)

	return opt
}

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

// WithGoOptions sets shared Go toolchain task options.
func WithGoOptions(goOpts ...taskGo.Option) Option {
	return func(o *Options) {
		o.taskGoOpts = append(o.taskGoOpts, goOpts...)
	}
}

// WithOutputDir sets the output directory, relative to the project root, for compilation artifacts.
func WithOutputDir(dir string) Option {
	return func(o *Options) {
		o.OutputDir = dir
	}
}
