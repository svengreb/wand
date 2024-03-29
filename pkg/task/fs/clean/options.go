// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package clean

const (
	// taskName is the name of the task.
	taskName = "fs/clean"
)

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	// limitToAppOutputDir indicates whether only paths within the configured application output directory should be
	// allowed.
	limitToAppOutputDir bool

	// name is the task name.
	name string

	// paths are paths to remove.
	// Note that only paths within the configured application output directory are allowed when limitToAppOutputDir is
	// enabled
	paths []string
}

// NewOptions creates new task options.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		name: taskName,
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// WithLimitToAppOutputDir indicates whether only paths within the configured application output directory should be
// allowed.
func WithLimitToAppOutputDir(limitToAppOutputDir bool) Option {
	return func(o *Options) {
		o.limitToAppOutputDir = limitToAppOutputDir
	}
}

// WithPaths sets the paths to remove.
// Note that only paths within the configured application output directory are allowed when WithLimitToAppOutputDir is
// enabled.
func WithPaths(paths ...string) Option {
	return func(o *Options) {
		o.paths = append(o.paths, paths...)
	}
}
