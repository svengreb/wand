// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package build provides a spell incantation for the "build" command of the Go toolchain.
package build

import (
	"path/filepath"

	"github.com/svengreb/wand/pkg/app"

	"github.com/svengreb/wand"
	"github.com/svengreb/wand/pkg/spell"
	spellGo "github.com/svengreb/wand/pkg/spell/golang"
)

// Spell is a spell incantation for the "build" command of the Go toolchain.
type Spell struct {
	ac   app.Config
	opts *Options
}

// Formula returns the spell incantation formula.
// Note that configured flags are applied after the "GOFLAGS" environment variable and could overwrite already defined
// flags.
// See `go help environment`, `go help env` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Environment_variables
func (s *Spell) Formula() []string {
	args := []string{"build"}

	args = append(args, spellGo.CompileFormula(s.opts.spellGoOpts...)...)

	if len(s.opts.Flags) > 0 {
		args = append(args, s.opts.Flags...)
	}

	args = append(
		args,
		"-o",
		filepath.Join(s.opts.OutputDir, s.opts.BinaryArtifactName),
		s.ac.PkgPath,
	)

	return args
}

// Kind returns the spell incantation kind.
func (s *Spell) Kind() spell.Kind {
	return spell.KindBinary
}

// Options returns the spell incantation options.
func (s *Spell) Options() interface{} {
	return *s.opts
}

// Env returns spell incantation specific environment variables.
func (s *Spell) Env() map[string]string {
	return s.opts.Env
}

// New creates a new spell incantation for the "build" command of the Go toolchain.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(wand wand.Wand, ac app.Config, opts ...Option) *Spell {
	opt := NewOptions(opts...)

	if opt.BinaryArtifactName == "" {
		opt.BinaryArtifactName = ac.Name
	}

	// Store build artifacts in the application specific subdirectory.
	if opt.OutputDir == "" {
		opt.OutputDir = ac.BaseOutputDir
	}

	return &Spell{ac: ac, opts: opt}
}
