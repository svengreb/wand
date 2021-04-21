// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package install

import (
	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project"
)

const (
	// TaskName is the name of the task.
	TaskName = "go/install"
)

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	// env is the task specific environment.
	env map[string]string

	// goModule is the Go module identifier.
	goModule *project.GoModuleID
}

// NewOptions creates new task options.
func NewOptions(opts ...Option) (*Options, error) {
	opt := &Options{
		goModule: &project.GoModuleID{},
	}
	for _, o := range opts {
		o(opt)
	}

	if opt.goModule.Version == nil {
		opt.goModule.IsLatest = true
	}

	return opt, nil
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
