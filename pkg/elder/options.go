// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package elder

import (
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"

	"github.com/svengreb/wand/pkg/project"
	taskGobin "github.com/svengreb/wand/pkg/task/gobin"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

// Option is a wand option.
type Option func(*Options)

// Options are wand options.
type Options struct {
	// gobinRunnerOpts are "gobin" runner options.
	gobinRunnerOpts []taskGobin.RunnerOption

	// goRunnerOpts are Go toolchain runner options.
	goRunnerOpts []taskGo.RunnerOption

	// nib is the log-level based line printer for human-facing messages.
	nib nib.Nib

	// projectOpts are project options.
	projectOpts []project.Option
}

// NewOptions creates new wand options.
func NewOptions(opts ...Option) *Options {
	opt := &Options{nib: inkpen.New()}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// WithGobinRunnerOptions sets "gobin" runner options.
func WithGobinRunnerOptions(opts ...taskGobin.RunnerOption) Option {
	return func(o *Options) {
		o.gobinRunnerOpts = append(o.gobinRunnerOpts, opts...)
	}
}

// WithGoRunnerOptions sets Go toolchain runner options.
func WithGoRunnerOptions(opts ...taskGo.RunnerOption) Option {
	return func(o *Options) {
		o.goRunnerOpts = append(o.goRunnerOpts, opts...)
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
