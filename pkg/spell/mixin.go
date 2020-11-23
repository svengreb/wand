// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package spell

// Options is a generic representation for spell incantation options.
type Options interface{}

// Mixin allows to compose functions that process Options of spell incantations.
type Mixin interface {
	// Apply applies generic Options to spell incantation options.
	Apply(Options) (Options, error)
}
