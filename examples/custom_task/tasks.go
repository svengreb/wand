// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

//go:build mage
// +build mage

package main

import (
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/task"
)

// MixOption is a mix task option.
type MixOption func(*MixOptions)

// MixOptions are mix task options.
type MixOptions struct {
	// env is the task specific environment.
	env map[string]string

	// extraArgs are additional arguments passed to the command.
	extraArgs []string

	// fruits are the fruits.
	fruits []string

	// verbose indicates whether the output should be verbose.
	verbose bool
}

// MixTask is a mix task for the fruit CLI.
type MixTask struct {
	ac   app.Config
	opts *MixOptions
}

// BuildParams builds the parameters.
func (t *MixTask) BuildParams() []string {
	var params []string

	// Toggle verbose output.
	if t.opts.verbose {
		params = append(params, "-v")
	}

	// Include additionally configured arguments.
	params = append(params, t.opts.extraArgs...)

	// Append all fruits.
	params = append(params, t.opts.fruits...)

	return params
}

// Env returns the task specific environment.
func (t *MixTask) Env() map[string]string {
	return t.opts.env
}

// Kind returns the task kind.
func (t *MixTask) Kind() task.Kind {
	return task.KindExec
}

// Options returns the task options.
func (t *MixTask) Options() task.Options {
	return *t.opts
}

// NewMixTask creates a new mix task for the fruit CLI.
func NewMixTask(ac app.Config, opts ...MixOption) (*MixTask, error) {
	return &MixTask{ac: ac, opts: NewMixOptions(opts...)}, nil
}

// NewMixOptions creates new mix task options.
func NewMixOptions(opts ...MixOption) *MixOptions {
	opt := &MixOptions{
		env: make(map[string]string),
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// WithEnv sets the mix task specific environment.
func WithEnv(env map[string]string) MixOption {
	return func(o *MixOptions) {
		o.env = env
	}
}

// WithFruits adds fruits.
func WithFruits(fruits ...string) MixOption {
	return func(o *MixOptions) {
		o.fruits = append(o.fruits, fruits...)
	}
}

// WithVerboseOutput indicates whether the output should be verbose.
func WithVerboseOutput(verbose bool) MixOption {
	return func(o *MixOptions) {
		o.verbose = verbose
	}
}
