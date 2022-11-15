// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

// Package git provides VCS utility functions to interact with Git repositories.
//
// See https://git-scm.com for more details about Git.
package git

import (
	"fmt"

	glGit "github.com/svengreb/golib/pkg/vcs/git"

	"github.com/svengreb/wand/pkg/project/vcs"
)

// Git represents a Git repository.
//
// See https://git-scm.com for more details.
type Git struct {
	opts *Options
}

// DeriveVersion derives the repository version based on Git metadata.
//
// References
//
//   (1) https://git-scm.com/docs/git-tag
//   (2) https://git-scm.com/book/en/v2/Git-Internals-Git-Objects
//   (3) https://git-scm.com/book/en/v2/Git-Internals-Git-References
func (g *Git) DeriveVersion() error {
	v, err := glGit.DeriveVersion(g.opts.defaultVersion, g.opts.path)
	if err != nil {
		return fmt.Errorf("failed to derive repository version from Git metadata: %w", err)
	}
	g.opts.version = v

	return nil
}

// Kind returns the repository Kind.
func (g *Git) Kind() vcs.Kind {
	return vcs.KindGit
}

// Version returns the repository version as type *Version.
func (g *Git) Version() interface{} {
	return g.opts.version
}

// New creates a new repository.
func New(opts ...Option) *Git {
	return &Git{opts: newOptions(opts...)}
}
