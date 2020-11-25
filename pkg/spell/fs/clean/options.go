// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package clean

// Options are clean.Spell options.
type Options struct {
	// limitToAppOutputDir indicates whether only paths within the configured application output directory should be
	// allowed.
	limitToAppOutputDir bool

	// paths are paths to remove.
	// Note that only paths within the configured application output directory are allowed when limitToAppOutputDir is
	// enabled
	paths []string
}

// Option is a clean.Spell option.
type Option func(*Options)

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

// NewOptions creates new clean.Spell options.
func NewOptions(opts ...Option) (*Options, error) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}

	return opt, nil
}
