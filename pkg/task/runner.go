// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package task

// Runner runs a command with parameters in a specific environment.
type Runner interface {
	// Handles returns the supported task kind.
	Handles() Kind

	// Run runs a command.
	Run(Task) error

	// Validate validates the runner.
	Validate() error
}

// RunnerExec is a runner for a (binary) command executable.
type RunnerExec interface {
	Runner

	// FilePath returns the path to the (binary) command executable.
	FilePath() string
}
