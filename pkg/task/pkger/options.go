// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package pkger

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

const (
	// DefaultGoModulePath is the default module import path.
	DefaultGoModulePath = "github.com/markbates/pkger/cmd/pkger"

	// DefaultGoModuleVersion is the default module version.
	DefaultGoModuleVersion = "v0.17.1"

	// MonorepoWorkaroundDummyFileContent is the file content for MonorepoWorkaroundDummyFileName.
	MonorepoWorkaroundDummyFileContent = `// Code generated by wand. DELETE THIS FILE IF NOT REMOVED AUTOMATICALLY.

package main`

	// MonorepoWorkaroundDummyFileName is the filename for the "dummy" workaround file.
	MonorepoWorkaroundDummyFileName = "wand_task_pkger_dummy_workaround"

	// TaskName is the name of the task.
	TaskName = "pkger"
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

	// includePathsRel are the relative paths of files and directories that should be included.
	// By default the paths will be detected by pkger itself when used within any of the packages of the target Go module.
	includePathsRel []string

	// outputDir is the relative path to the output directory of the generated "pkger.go" file.
	outputDirRel string
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

// WithIncludes adds the relative paths of files and directories that should be included.
// By default the paths will be detected by pkger itself when used within any of the packages of the target Go module.
func WithIncludes(includes ...string) Option {
	return func(o *Options) {
		o.includePathsRel = append(o.includePathsRel, includes...)
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