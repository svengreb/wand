// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package task provides tasks and runner for Mage.
package task

import (
	"github.com/svengreb/wand/pkg/project"
)

// Exec is a task for a executable command.
type Exec interface {
	Task

	// BuildParams builds the parameters for a command runner.
	// Parameters consist of options, flags and arguments.
	// The separation of parameters from commands enables a flexible usage, e.g. when parameters can be reused for
	// different tasks.
	//
	// References
	//
	//   (1) https://en.wikipedia.org/wiki/Command-line_interface#Anatomy_of_a_shell_CLI
	BuildParams() []string

	// Env returns the task specific environment.
	Env() map[string]string
}

// GoModule is a task for a Go module command.
//
// See https://golang.org/ref/mod for more details about Go modules.
type GoModule interface {
	Exec

	// ID returns the identifier of a Go module.
	ID() *project.GoModuleID
}

// Task is a wand task for Mage.
//
// See https://magefile.org/targets for more details about Mage targets.
type Task interface {
	// Kind returns the task kind.
	Kind() Kind

	// Options returns the task options.
	Options() Options
}
