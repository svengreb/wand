// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package vcs

// Repository is a VCS repository.
type Repository interface {
	// Kind returns the repository Kind.
	Kind() Kind
	// DeriveVersion derives the repository version based on the Kind.
	DeriveVersion() error
	// Version returns the repository version.
	Version() interface{}
}
