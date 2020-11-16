// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package project

import (
	"errors"
	"fmt"
)

const (
	// ErrDeriveVCSInformation indicates that the derivation of VCS version information failed.
	ErrDeriveVCSInformation = ErrString("failed to derive VCS version information")

	// ErrDetectProjectRootDirPath indicates that the detection of a project root directory path failed.
	ErrDetectProjectRootDirPath = ErrString("failed to detect project root directory path")

	// ErrDetermineGoModuleInformation indicates that a determination of Go module information failed.
	ErrDetermineGoModuleInformation = ErrString("failed to determine Go module information")

	// ErrPathNotRelative indicates that a path is not releative.
	ErrPathNotRelative = ErrString("path is not relative")
)

// ErrString is a string type for implementing constant errors.
type ErrString string

func (e ErrString) Error() string {
	return string(e)
}

// ErrProject represents a project error.
type ErrProject struct {
	// Err is a wrapped error.
	Err error
	// Kind is the error kind.
	Kind error
}

func (e *ErrProject) Error() string {
	msg := "project error"
	if e.Kind != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Kind)
	}
	if e.Err != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Err)
	}

	return msg
}

// Is enables usage of errors.Is() to determine the kind of error that occurred.
func (e *ErrProject) Is(err error) bool {
	return errors.Is(err, e.Kind)
}

// Unwrap returns the underlying error for usage with errors.Unwrap().
func (e *ErrProject) Unwrap() error { return e.Err }
