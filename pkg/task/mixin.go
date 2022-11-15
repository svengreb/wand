// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package task

// Mixin allows to compose functions that process task options.
type Mixin interface {
	// Apply applies the mixin to task options.
	Apply(Options) (Options, error)
}

// Options is a generic representation for task options.
type Options interface{}
