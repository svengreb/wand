// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package env

const (
	// taskName is the name of the task.
	taskName = "go/env"
)

// Option is a task option.
type Option func(*Options)

// Options are task options.
type Options struct {
	// EnableJSONOutput indicates whether the output should be in JSON format.
	EnableJSONOutput bool

	// env is the task specific environment.
	env map[string]string

	// EnvVars are the names of the target environment variables.
	EnvVars []string

	// extraArgs are additional arguments passed to the command.
	extraArgs []string

	// name is the task name.
	name string
}

// NewOptions creates new task options.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		name: taskName,
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// WithEnv sets the task specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		o.env = env
	}
}

// WithEnvVars sets the names of the target environment variables.
func WithEnvVars(envVars ...string) Option {
	return func(o *Options) {
		o.EnvVars = append(o.EnvVars, envVars...)
	}
}

// WithExtraArgs sets additional arguments to pass to the command.
func WithExtraArgs(extraArgs ...string) Option {
	return func(o *Options) {
		o.extraArgs = append(o.extraArgs, extraArgs...)
	}
}
