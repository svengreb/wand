// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package cast

import (
	"errors"
	"fmt"

	wErr "github.com/svengreb/wand/pkg/error"
)

const (
	// ErrCasterCasting indicates that a caster failed to cast.
	ErrCasterCasting = wErr.ErrString("failed to cast")

	// ErrCasterInvalidOpts indicates invalid caster options.
	ErrCasterInvalidOpts = wErr.ErrString("invalid caster options")

	// ErrCasterValidation indicates that a caster validation failed.
	ErrCasterValidation = wErr.ErrString("caster validation failed")

	// ErrCasterSpellIncantationKindUnsupported indicates that a spell incantation kind is not supported by a caster.
	ErrCasterSpellIncantationKindUnsupported = wErr.ErrString("unsupported spell incantation kind")
)

// ErrCast represents a cast error.
type ErrCast struct {
	// Err is a wrapped error.
	Err error
	// Kind is the error kind.
	Kind error
}

func (e *ErrCast) Error() string {
	msg := "cast error"
	if e.Kind != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Kind)
	}
	if e.Err != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.Err)
	}

	return msg
}

// Is enables usage of errors.Is() to determine the kind of error that occurred.
func (e *ErrCast) Is(err error) bool {
	return errors.Is(err, e.Kind)
}

// Unwrap returns the underlying error for usage with errors.Unwrap().
func (e *ErrCast) Unwrap() error { return e.Err }
