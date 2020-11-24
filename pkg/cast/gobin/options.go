// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package gobin

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/svengreb/wand/pkg/cast"
	"github.com/svengreb/wand/pkg/project"
)

const (
	// CasterName is the name of the Go toolchain command caster.
	CasterName = "gobin"

	// DefaultExec is the default name of the "github.com/myitcv/gobin" module executable.
	DefaultExec = "gobin"

	// DefaultGoModulePath is the default "gobin" module import path.
	DefaultGoModulePath = "github.com/myitcv/gobin"

	// DefaultGoModuleVersion is the default "gobin" module version.
	DefaultGoModuleVersion = "v0.0.14"
)

// Options stores "github.com/myitcv/gobin" module caster options.
type Options struct {
	// Env are caster specific environment variables.
	Env map[string]string

	// Exec ist the name or path of the "gobin" module executable.
	Exec string

	goModule *project.GoModuleID
}

// Option is a "github.com/myitcv/gobin" module caster option.
type Option func(*Options)

// WithExec sets the name or path to the "github.com/myitcv/gobin" module executable.
// Defaults to DefaultExec.
func WithExec(nameOrPath string) Option {
	return func(o *Options) {
		if nameOrPath != "" {
			o.Exec = nameOrPath
		}
	}
}

// WithModulePath sets the "gobin" module import path.
// Defaults to DefaultGoModulePath.
func WithModulePath(path string) Option {
	return func(o *Options) {
		if path != "" {
			o.goModule.Path = path
		}
	}
}

// WithModuleVersion sets the "gobin" module version.
// Defaults to DefaultGoModuleVersion.
func WithModuleVersion(version *semver.Version) Option {
	return func(o *Options) {
		if version != nil {
			o.goModule.Version = version
		}
	}
}

// newOptions creates new "github.com/myitcv/gobin" module caster options.
func newOptions(opts ...Option) (*Options, error) {
	version, versionErr := semver.NewVersion(DefaultGoModuleVersion)
	if versionErr != nil {
		return nil, &cast.ErrCast{
			Err:  fmt.Errorf("parsing default module version %q: %w", DefaultGoModulePath, versionErr),
			Kind: cast.ErrCasterInvalidOpts,
		}
	}
	opt := &Options{
		Env:  make(map[string]string),
		Exec: DefaultExec,
		goModule: &project.GoModuleID{
			Path:    DefaultGoModulePath,
			Version: version,
		},
	}
	for _, o := range opts {
		o(opt)
	}

	return opt, nil
}
