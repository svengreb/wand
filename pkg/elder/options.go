// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package elder

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"

	castGobin "github.com/svengreb/wand/pkg/cast/gobin"
	castGoToolchain "github.com/svengreb/wand/pkg/cast/golang/toolchain"
	"github.com/svengreb/wand/pkg/project"
)

// Options are options for the wand.Wand reference implementation "elder".
type Options struct {
	// gobinCasterOpts are "gobin" caster options.
	gobinCasterOpts []castGobin.Option

	// goToolchainCasterOpts are Go toolchain caster options.
	goToolchainCasterOpts []castGoToolchain.Option

	// nib is the log-level based line printer for human-facing messages.
	nib nib.Nib

	// goToolchainCasterOpts are project options.
	projectOpts []project.Option
}

// Option is a option for the wand.Wand reference implementation "elder".
type Option func(*Options)

// WithGobinCasterOptions sets "gobin" caster options.
func WithGobinCasterOptions(opts ...castGobin.Option) Option {
	return func(o *Options) {
		o.gobinCasterOpts = append(o.gobinCasterOpts, opts...)
	}
}

// WithGoToolchainCasterOptions sets Go toolchain caster options.
func WithGoToolchainCasterOptions(opts ...castGoToolchain.Option) Option {
	return func(o *Options) {
		o.goToolchainCasterOpts = append(o.goToolchainCasterOpts, opts...)
	}
}

// WithNib sets the log-level based line printer for human-facing messages.
func WithNib(n nib.Nib) Option {
	return func(o *Options) {
		if n != nil {
			o.nib = n
		}
	}
}

// WithProjectOptions sets project options.
func WithProjectOptions(opts ...project.Option) Option {
	return func(o *Options) {
		o.projectOpts = append(o.projectOpts, opts...)
	}
}

// NewOptions creates new options for the wand.Wand reference implementation "elder".
func NewOptions(opts ...Option) *Options {
	opt := &Options{nib: inkpen.New()}
	for _, o := range opts {
		o(opt)
	}

	return opt
}
