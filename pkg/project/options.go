// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package project

import (
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/project/vcs"
	vcsNone "github.com/svengreb/wand/pkg/project/vcs/none"
)

const (
	// DefaultBaseOutputDir is the default base output directory relative to Options.RootDirPathAbs for compile, test
	// and production artifacts as well as distribution bundles, static web files or metric/statistic reports.
	DefaultBaseOutputDir = "out"

	// DefaultVersion is the default version for a project vcs.Repository.
	DefaultVersion = "v0.0.0"
)

// Options stores project options.
type Options struct {
	// BaseOutputDir is the base project output directory, relative to RootDirPathAbs, for compile, test and production
	// artifacts as well as distribution bundles, static web files or metric/statistic reports.
	BaseOutputDir string

	// DefaultVersion is the default project version.
	DefaultVersion string

	// DisplayName is the project display name.
	DisplayName string

	// GoModule is the project Go module.
	GoModule *GoModuleID

	// Name is the project name.
	Name string

	// Repository is the project repository.
	Repository vcs.Repository

	// RootDirPathAbs is the absolute path to the project root directory.
	RootDirPathAbs string

	// VCSKind is the VCS kind of the project Repository.
	VCSKind vcs.Kind
}

// Option is a project option.
type Option func(*Options)

// WithBaseOutputDir sets the base output directory.
func WithBaseOutputDir(dir string) Option {
	return func(o *Options) {
		o.BaseOutputDir = dir
	}
}

// WithDefaultVersion set the project default verion.
func WithDefaultVersion(defaultVersion string) Option {
	return func(o *Options) {
		o.DefaultVersion = defaultVersion
	}
}

// WithDisplayName sets the project display name.
func WithDisplayName(name string) Option {
	return func(o *Options) {
		o.DisplayName = name
	}
}

// WithGoModule sets the project Go module.
func WithGoModule(module *GoModuleID) Option {
	return func(o *Options) {
		o.GoModule = module
	}
}

// WithName sets the project name.
func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

// WithVCSKind sets the vcs.Kind of the project vcs.Repository.
func WithVCSKind(kind vcs.Kind) Option {
	return func(o *Options) {
		o.VCSKind = kind
	}
}

// newOptions creates new project options.
// The absolute path to the root directory is automatically set based on the current working directory and the Go module
// name is automatically determined using the runtime/debug package.
func newOptions(opts ...Option) (*Options, error) {
	rootDirPath, pwdErr := os.Getwd()
	if pwdErr != nil {
		return nil, &ErrProject{
			Err:  pwdErr,
			Kind: ErrDetectProjectRootDirPath,
		}
	}

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, &ErrProject{Kind: ErrDetermineGoModuleInformation}
	}

	goModuleVersion, goModuleVersionErr := semver.NewVersion(buildInfo.Main.Version)
	if goModuleVersionErr == nil {
		return nil, &ErrProject{Err: goModuleVersionErr, Kind: ErrDetermineGoModuleInformation}
	}

	opt := &Options{
		BaseOutputDir:  DefaultBaseOutputDir,
		DefaultVersion: DefaultVersion,
		DisplayName:    filepath.Base(rootDirPath),
		GoModule: &GoModuleID{
			Path:    buildInfo.Main.Path,
			Version: goModuleVersion,
		},
		Name:           filepath.Base(rootDirPath),
		Repository:     vcsNone.New(),
		RootDirPathAbs: rootDirPath,
		VCSKind:        vcs.KindNone,
	}
	for _, o := range opts {
		o(opt)
	}

	return opt, nil
}
