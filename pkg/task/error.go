// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package task

import (
	"errors"
	"fmt"

	wErr "github.com/svengreb/wand/pkg/error"
)

const (
	// ErrIncompatibleRunner indicates that a command runner is not compatible for a task.
	ErrIncompatibleRunner = wErr.ErrString("incompatible command runner")

	// ErrInvalidRunnerOpts indicates invalid command runner options.
	ErrInvalidRunnerOpts = wErr.ErrString("invalid options")

	// ErrInvalidTaskOpts indicates invalid task options.
	ErrInvalidTaskOpts = wErr.ErrString("invalid options")

	// ErrRun indicates that a runner failed to run.
	ErrRun = wErr.ErrString("failed to run")

	// ErrRunnerValidation indicates that a command runner validation failed.
	ErrRunnerValidation = wErr.ErrString("validation failed")

	// ErrUnsupportedTaskKind indicates that a task kind is not supported.
	ErrUnsupportedTaskKind = wErr.ErrString("unsupported task kind")

	// ErrUnsupportedTaskOptions indicates that the task options are not supported.
	ErrUnsupportedTaskOptions = wErr.ErrString("unsupported task kind")
)

// ErrRunner represents a runner error.
type ErrRunner struct {
	// Err is a wrapped error.
	Err error
	// Kind is the error kind.
	Kind error
}

func (e *ErrRunner) Error() string {
	msg := "runner error"
	if e.Kind != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Kind)
	}
	if e.Err != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Err)
	}

	return msg
}

// Is enables usage of errors.Is() to determine the kind of error that occurred.
func (e *ErrRunner) Is(err error) bool {
	return errors.Is(err, e.Kind)
}

// Unwrap returns the underlying error for usage with errors.Unwrap().
func (e *ErrRunner) Unwrap() error { return e.Err }

// ErrTask represents a task error.
type ErrTask struct {
	// Err is a wrapped error.
	Err error
	// Kind is the error kind.
	Kind error
}

func (e *ErrTask) Error() string {
	msg := "task error"
	if e.Kind != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Kind)
	}
	if e.Err != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Err)
	}

	return msg
}

// Is enables usage of errors.Is() to determine the kind of error that occurred.
func (e *ErrTask) Is(err error) bool {
	return errors.Is(err, e.Kind)
}

// Unwrap returns the underlying error for usage with errors.Unwrap().
func (e *ErrTask) Unwrap() error { return e.Err }
