// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

//go:build go1.17

package run

import (
	"github.com/Masterminds/semver/v3"
	"github.com/svengreb/wand/pkg/project"
)

const (
	// taskName is the name of the task.
	taskName = "go/run"
)

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
}

// NewOptions creates new task options.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		goModule: &project.GoModuleID{},
		name:     taskName,
	}
	for _, o := range opts {
		o(opt)
	}
	if opt.goModule.Version == nil {
		opt.goModule.IsLatest = true
	}
	return opt
}

// WithArgs sets additional arguments to pass to the command.
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
		o.goModule.Version = version
	}
}
