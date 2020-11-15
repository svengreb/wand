// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package git

// Options stores repository options.
type Options struct {
	// defaultVersion is the default repository version.
	defaultVersion string

	// path is the absolute repository path.
	path string

	// version is the repository version derived from Git metadata.
	version *Version
}

// Option is a repository option.
type Option func(*Options)

// WithDefaultVersion sets the default version.
func WithDefaultVersion(defaultVersion string) Option {
	return func(o *Options) {
		o.defaultVersion = defaultVersion
	}
}

// WithPath sets the repository path.
func WithPath(path string) Option {
	return func(o *Options) {
		o.path = path
	}
}

// newOptions creates new repository options.
func newOptions(opts ...Option) *Options {
	opt := &Options{
		version: &Version{},
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}
