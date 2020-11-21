// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package toolchain

import (
	"github.com/magefile/mage/mg"
)

const (
	// CasterName is the name of the Go toolchain command caster.
	CasterName = "golang"

	// DefaultEnvVarGO111MODULE is the default environment variable name to toggle the Go 1.11 module mode.
	DefaultEnvVarGO111MODULE = "GO111MODULE"

	// DefaultEnvVarGOBIN is the default environment variable name for the Go binary executable path.
	DefaultEnvVarGOBIN = "GOBIN"

	// DefaultEnvVarGOFLAGS is the default environment variable name for Go tool flags.
	DefaultEnvVarGOFLAGS = "GOFLAGS"

	// DefaultEnvVarGOPATH is the default environment variable name for the Go path.
	DefaultEnvVarGOPATH = "GOPATH"
)

// DefaultExec is the default path to the Go executable.
var DefaultExec = mg.GoCmd()

// Options stores Go toolchain command caster options.
type Options struct {
	Env  map[string]string
	Exec string
}

// Option is a Go toolchain command caster option.
type Option func(*Options)

// WithEnv sets the caster environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.Env = env
	}
}

// WithExec sets the path to the Go executable.
func WithExec(nameOrPath string) Option {
	return func(o *Options) {
		o.Exec = nameOrPath
	}
}

// newOptions creates new Go toolchain command caster options.
func newOptions(opts ...Option) *Options {
	opt := &Options{
		Env:  make(map[string]string),
		Exec: DefaultExec,
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}