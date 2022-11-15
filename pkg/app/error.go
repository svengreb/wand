// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package app

import (
	"errors"
	"fmt"

	wErr "github.com/svengreb/wand/pkg/error"
)

const (
	// ErrNoSuchConfig indicates that an application configuration was not found in the store.
	ErrNoSuchConfig = wErr.ErrString("no such configuration")

	// ErrEmptyName indicates that an application name is empty.
	ErrEmptyName = wErr.ErrString("application name is empty")

	// ErrPathNotRelative indicates that an application path is not relative.
	ErrPathNotRelative = wErr.ErrString("path is not relative")

	// ErrNonProjectRootSubDir indicates that an application path is not a subdirectory of the project root directory.
	ErrNonProjectRootSubDir = wErr.ErrString("path is not a subdirectory of the project root directory")
)

// ErrApp represents a application error.
type ErrApp struct {
	// Err is a wrapped error.
	Err error
	// Kind is the error kind.
	Kind error
}

func (e *ErrApp) Error() string {
	msg := "application error"
	if e.Kind != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Kind)
	}
	if e.Err != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Err)
	}

	return msg
}

// Is enables usage of errors.Is() to determine the kind of error that occurred.
func (e *ErrApp) Is(err error) bool {
	return errors.Is(err, e.Kind)
}

// Unwrap returns the underlying error for usage with errors.Unwrap().
func (e *ErrApp) Unwrap() error { return e.Err }
