// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package test provide a spell incantation for the "test" command of the Go toolchain.
package test

import (
	"fmt"
	"path/filepath"

	"github.com/svengreb/wand"
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/spell"
	spellGo "github.com/svengreb/wand/pkg/spell/golang"
)

// Spell is a spell incantation for the "test" command of the Go toolchain.
type Spell struct {
	ac   app.Config
	opts *Options
}

// Formula returns the spell incantation formula.
// Note that configured flags are applied after the "GOFLAGS" environment variable and could overwrite already defined
// flags. In addition, the output directory for test artifacts like profiles and reports must exist or must be be
// created before, otherwise the "test" Go toolchain command will fail to run.
// See `go help environment`, `go help env` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Environment_variables
func (s *Spell) Formula() []string {
	args := []string{"test"}

	args = append(args, spellGo.CompileFormula(s.opts.spellGoOpts...)...)

	if s.opts.EnableVerboseOutput {
		args = append(args, "-v")
	}

	if s.opts.DisableCache {
		args = append(args, "-count=1")
	}

	if s.opts.EnableBlockProfile {
		args = append(args,
			fmt.Sprintf(
				"-blockprofile=%s",
				filepath.Join(s.opts.OutputDir, s.opts.BlockProfileOutputFileName),
			),
		)
	}

	if s.opts.EnableCoverageProfile {
		args = append(args,
			fmt.Sprintf(
				"-coverprofile=%s",
				filepath.Join(s.opts.OutputDir, s.opts.CoverageProfileOutputFileName),
			),
		)
	}

	if s.opts.EnableCPUProfile {
		args = append(args,
			fmt.Sprintf("-cpuprofile=%s",
				filepath.Join(s.opts.OutputDir, s.opts.CPUProfileOutputFileName),
			),
		)
	}

	if s.opts.EnableMemProfile {
		args = append(args,
			fmt.Sprintf("-memprofile=%s",
				filepath.Join(s.opts.OutputDir, s.opts.MemoryProfileOutputFileName),
			),
		)
	}

	if s.opts.EnableMutexProfile {
		args = append(args,
			fmt.Sprintf("-mutexprofile=%s",
				filepath.Join(s.opts.OutputDir, s.opts.MutexProfileOutputFileName),
			),
		)
	}

	if s.opts.EnableTraceProfile {
		args = append(args,
			fmt.Sprintf("-trace=%s",
				filepath.Join(s.opts.OutputDir, s.opts.TraceProfileOutputFileName),
			),
		)
	}

	if len(s.opts.Flags) > 0 {
		args = append(args, s.opts.Flags...)
	}

	args = append(args, s.opts.Pkgs...)

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

	// Store test profiles and reports within the application specific subdirectory.
	if opt.OutputDir == "" {
		opt.OutputDir = filepath.Join(ac.BaseOutputDir, DefaultOutputDirName)
	}

	return &Spell{ac: ac, opts: opt}
}
