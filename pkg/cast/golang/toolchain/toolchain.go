// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package toolchain

import (
	"fmt"
	"os/exec"

	"github.com/magefile/mage/sh"
	glFS "github.com/svengreb/golib/pkg/io/fs"

	"github.com/svengreb/wand/pkg/cast"
	"github.com/svengreb/wand/pkg/spell"
)

// Caster is a Go toolchain command caster.
type Caster struct {
	opts *Options
}

// GetExec returns the path to the binary executable.
func (c *Caster) GetExec() string {
	return c.opts.Exec
}

// Cast casts a spell incantation.
// It returns an error of type *cast.ErrCast when the spell is not a spell.KindBinary and any other error that occurs
// during the command execution.
func (c *Caster) Cast(si spell.Incantation) error {
	if si.Kind() != spell.KindBinary {
		return &cast.ErrCast{
			Err:  fmt.Errorf("%q", si.Kind()),
			Kind: cast.ErrCasterSpellIncantationKindUnsupported,
		}
	}

	s, ok := si.(spell.Binary)
	if !ok {
		return &cast.ErrCast{
			Err:  fmt.Errorf("expected %q but got %q", s.Kind(), si.Kind()),
			Kind: cast.ErrCasterSpellIncantationKindUnsupported,
		}
	}

	args := si.Formula()
	for k, v := range s.Env() {
		c.opts.Env[k] = v
	}

	return sh.RunWithV(c.opts.Env, c.opts.Exec, args...)
}

// Handles returns the supported spell.Kind.
func (c *Caster) Handles() spell.Kind {
	return spell.KindBinary
}

// Validate validates the Go toolchain command caster.
// It returns an error of type *cast.ErrCast when the binary executable does not exists at the configured path and when
// it is also not available in the executable search paths of the current environment.
func (c *Caster) Validate() error {
	// Check if the Go executable exists,...
	execExits, fsErr := glFS.FileExists(c.opts.Exec)
	if fsErr != nil {
		return &cast.ErrCast{
			Err:  fmt.Errorf("caster %q: %w", CasterName, fsErr),
			Kind: cast.ErrCasterValidation,
		}
	}
	// ...otherwise try to look up the system-wide executable paths.
	if !execExits {
		path, pathErr := exec.LookPath(c.opts.Exec)
		if pathErr != nil {
			return &cast.ErrCast{
				Err:  fmt.Errorf("caster %q: %q not found or does not exist: %w", CasterName, c.opts.Exec, pathErr),
				Kind: cast.ErrCasterValidation,
			}
		}
		c.opts.Exec = path
	}

	return nil
}

// NewCaster creates a new Go toolchain command caster.
func NewCaster(opts ...Option) *Caster {
	return &Caster{opts: newOptions(opts...)}
}
