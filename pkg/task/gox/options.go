// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package gox

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoBuild "github.com/svengreb/wand/pkg/task/golang/build"
)

const (
	// DefaultGoModulePath is the default module import path.
	DefaultGoModulePath = "github.com/mitchellh/gox"

	// DefaultGoModuleVersion is the default module version.
	DefaultGoModuleVersion = "v1.0.1"

	// taskName is the name of the task.
	taskName = "gox"
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

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	*taskGoBuild.Options

	// env is the task specific environment.
	env map[string]string

	// goCmd is the path to the Go toolchain executable.
	goCmd string

	// goModule is the Go module identifier.
	goModule *project.GoModuleID

	// name is the task name.
	name string

	// outputTemplate is the name template for cross-compile platform targets.
	outputTemplate string

	// taskGoBuildOpts are Go toolchain "build" command task options.
	taskGoBuildOpts []taskGoBuild.Option

	// taskGoOpts are shared Go toolchain task options.
	taskGoOpts []taskGo.Option

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

	goBuildOpts := append(
		[]taskGoBuild.Option{taskGoBuild.WithGoOptions(opt.taskGoOpts...)},
		opt.taskGoBuildOpts...,
	)
	opt.Options = taskGoBuild.NewOptions(goBuildOpts...)

	if opt.outputTemplate == "" && opt.BinaryArtifactName != "" {
		opt.outputTemplate = DefaultCrossCompileBinaryNameTemplate(opt.BinaryArtifactName)
	}

	if len(opt.CrossCompileTargetPlatforms) == 0 {
		opt.CrossCompileTargetPlatforms = DefaultCrossCompileTargetPlatforms
	}

	return opt, nil
}

// WithEnv sets the task specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.env = env
	}
}

// WithGoBuildOptions sets Go toolchain "build" command task options.
func WithGoBuildOptions(goBuildOpts ...taskGoBuild.Option) Option {
	return func(o *Options) {
		o.taskGoBuildOpts = append(o.taskGoBuildOpts, goBuildOpts...)
	}
}

// WithGoCmd sets the path to the Go toolchain executable.
func WithGoCmd(goCmd string) Option {
	return func(o *Options) {
		o.goCmd = goCmd
	}
}

// WithGoOptions sets shared Go toolchain task options.
func WithGoOptions(goOpts ...taskGo.Option) Option {
	return func(o *Options) {
		o.taskGoOpts = append(o.taskGoOpts, goOpts...)
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

// WithOutputTemplate sets the name template for cross-compile platform targets.
// Defaults to DefaultCrossCompileBinaryNameTemplate.
func WithOutputTemplate(outputTemplate string) Option {
	return func(o *Options) {
		o.outputTemplate = outputTemplate
	}
}

// WithVerboseOutput indicates whether the output should be verbose.
func WithVerboseOutput(verbose bool) Option {
	return func(o *Options) {
		o.verbose = verbose
	}
}
