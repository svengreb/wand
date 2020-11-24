// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package gox provides a spell incantation for the "github.com/mitchellh/gox" Go module command, a dead simple,
// no frills Go cross compile tool that behaves a lot like the standard Go toolchain "build" command.
// See https://pkg.go.dev/github.com/mitchellh/gox for more details about "gox".
// The source code of the "gox" is available at https://github.com/mitchellh/gox.
package gox

import (
	"fmt"
	"strings"

	"github.com/svengreb/wand"
	"github.com/svengreb/wand/pkg/app"
	castGoToolchain "github.com/svengreb/wand/pkg/cast/golang/toolchain"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/spell"
	spellGo "github.com/svengreb/wand/pkg/spell/golang"
)

// Spell is a spell incantation for the "github.com/mitchellh/gox" Go module command.
type Spell struct {
	ac   app.Config
	opts *Options
}

// Formula returns the spell incantation formula.
func (s *Spell) Formula() []string {
	args := spellGo.CompileFormula(s.opts.spellGoOpts...)

	// Workaround to allow the usage of the "-trimpath" flag that has been introduced in Go 1.13.0.
	// The currently latest version of "gox" does not support the flag yet.
	// See https://github.com/mitchellh/gox/pull/138 for more details.
	for idx, arg := range args {
		if arg == "-trimpath" {
			args = append(args[:idx], args[idx+1:]...)
			// Set the flag via the GOFLAGS environment variable instead.
			s.opts.env[castGoToolchain.DefaultEnvVarGOFLAGS] = fmt.Sprintf(
				"%s %s -trimpath",
				s.opts.Env[castGoToolchain.DefaultEnvVarGOFLAGS],
				s.opts.env[castGoToolchain.DefaultEnvVarGOFLAGS],
			)
		}
	}

	if s.opts.verbose {
		args = append(args, "-verbose")
	}

	if s.opts.goCmd != "" {
		args = append(args, fmt.Sprintf("-gocmd=%s", s.opts.goCmd))
	}

	if len(s.opts.CrossCompileTargetPlatforms) > 0 {
		args = append(args, fmt.Sprintf("-osarch=%s", strings.Join(s.opts.CrossCompileTargetPlatforms, " ")))
	}

	args = append(args, fmt.Sprintf("--output=%s/%s", s.opts.OutputDir, s.opts.outputTemplate))

	if len(s.opts.Flags) > 0 {
		args = append(args, s.opts.Flags...)
	}

	return append(args, s.ac.PkgPath)
}

// Kind returns the spell incantation kind.
func (s *Spell) Kind() spell.Kind {
	return spell.KindGoModule
}

// Options returns the spell incantation options.
func (s *Spell) Options() interface{} {
	return *s.opts
}

// GoModuleID returns partial Go module identifier information.
func (s *Spell) GoModuleID() *project.GoModuleID {
	return s.opts.goModule
}

// Env returns spell incantation specific environment variables.
func (s *Spell) Env() map[string]string {
	return s.opts.env
}

// New creates a new spell incantation for the "build" command of the Go toolchain.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(wand wand.Wand, ac app.Config, opts ...Option) (*Spell, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}

	if opt.BinaryArtifactName == "" {
		opt.BinaryArtifactName = ac.Name
	}

	// Store builds artifacts in the application specific sub-folder.
	if opt.OutputDir == "" {
		opt.OutputDir = ac.BaseOutputDir
	}

	if opt.outputTemplate == "" {
		opt.outputTemplate = DefaultCrossCompileBinaryNameTemplate(opt.BinaryArtifactName)
	}

	return &Spell{ac: ac, opts: opt}, nil
}
