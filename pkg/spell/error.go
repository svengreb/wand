// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package spell

import (
	"errors"
	"fmt"

	wErr "github.com/svengreb/wand/pkg/error"
)

const (
	// ErrExec indicates that a GoCode spell incantation returned an error during the code execution.
	ErrExec = wErr.ErrString("error returned during execution")
)

// ErrGoCode represents a GoCode error.
type ErrGoCode struct {
	// Err is a wrapped error.
	Err error
	// Kind is the error kind.
	Kind error
}

func (e *ErrGoCode) Error() string {
	msg := "Go code"
	if e.Kind != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Kind)
	}
	if e.Err != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Err)
	}

	return msg
}

// Is enables usage of errors.Is() to determine the kind of error that occurred.
func (e *ErrGoCode) Is(err error) bool {
	return errors.Is(err, e.Kind)
}

// Unwrap returns the underlying error for usage with errors.Unwrap().
func (e *ErrGoCode) Unwrap() error { return e.Err }
