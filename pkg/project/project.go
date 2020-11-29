// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package project provides metadata and VCS information of a project.
package project

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/svengreb/wand/pkg/project/vcs"
	vcsGit "github.com/svengreb/wand/pkg/project/vcs/git"
)

// Metadata represents information about a project.
type Metadata struct {
	opts *Options
}

// Options returns the project Options.
func (m Metadata) Options() Options {
	return *m.opts
}

// New creates new project metadata.
//
// The absolute path to the root directory is automatically set based on the current working directory while the Go
// module name is determined using the runtime/debug package.
//
// The project version is derived from the vcs.Repository if not of type vcs.KindNone. The currently only supported
// vcs.Kind is vcs.KindGit. To set the vcs.Kind the WithVCSKind() project Option can be used.
//
// If any error occurs nil is returned along with an error of type *ErrProject.
func New(opts ...Option) (*Metadata, error) {
	opt, optErr := newOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}

	if filepath.IsAbs(opt.BaseOutputDir) {
		return nil, &ErrProject{
			Err:  fmt.Errorf("failed to validate base output directory %q", opt.BaseOutputDir),
			Kind: ErrPathNotRelative,
		}
	}

	var versionErr error
	switch k := opt.VCSKind; k {
	case vcs.KindGit:
		opt.Repository = vcsGit.New(
			vcsGit.WithDefaultVersion(opt.DefaultVersion),
			vcsGit.WithPath(opt.RootDirPathAbs),
		)
		versionErr = opt.Repository.DeriveVersion()
	case vcs.KindNone:
	}

	if versionErr != nil {
		return nil, &ErrProject{
			Err:  errors.New(versionErr.Error()),
			Kind: ErrDeriveVCSInformation,
		}
	}

	return &Metadata{opts: opt}, nil
}
