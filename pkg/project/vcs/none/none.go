// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

// Package none provides a nonexistent repository.
package none

import (
	"github.com/svengreb/wand/pkg/project/vcs"
)

// None represents a nonexistent repository.
type None struct {
	opts *Options
}

// DeriveVersion derives the repository version.
// Note that this is always nil for a nonexistent repository.
func (n *None) DeriveVersion() error {
	return nil
}

// Kind returns the repository Kind.
func (n *None) Kind() vcs.Kind {
	return vcs.KindNone
}

// Version returns the repository version as type string.
// Note that this is always the configured default version.
func (n *None) Version() interface{} {
	return n.opts.defaultVersion
}

// New creates a new repository.
func New(opts ...Option) *None {
	return &None{opts: newOptions(opts...)}
}
