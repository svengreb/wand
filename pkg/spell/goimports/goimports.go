// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package goimports provides a spell incantation for the "golang.org/x/tools/cmd/goimports" Go module command.
// See https://pkg.go.dev/golang.org/x/tools/cmd/goimports for more details about "goimports".
// The source code of "goimports" is available at https://github.com/golang/tools/tree/master/cmd/goimports.
package goimports

import (
	"fmt"
	"strings"

	"github.com/svengreb/wand"
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/spell"
)

// Spell is a spell incantation for the "golang.org/x/tools/cmd/goimports" Go module command.
type Spell struct {
	ac   app.Config
	opts *Options
}

// Formula returns the spell incantation formula.
func (s *Spell) Formula() []string {
	var args []string

	// List files whose formatting are non-compliant to the style guide.
	if s.opts.listNonCompliantFiles {
		args = append(args, "-l")
	}

	// A comma-separated list of prefixes for local package imports to be put after 3rd-party packages.
	if len(s.opts.localPkgs) > 0 {
		args = append(args, "-local", fmt.Sprintf("'%s'", strings.Join(s.opts.localPkgs, ",")))
	}

	// Report all errors and not just the first 10 on different lines.
	if s.opts.reportAllErrors {
		args = append(args, "-e")
	}

	// Write result to source files instead of stdout.
	if s.opts.persistChanges {
		args = append(args, "-w")
	}

	// Enable verbose output.
	if s.opts.verbose {
		args = append(args, "-v")
	}

	// Include additionally configured arguments.
	args = append(args, s.opts.extraArgs...)

	// Only search in specified paths for Go source files...
	if len(s.opts.paths) > 0 {
		args = append(args, s.opts.paths...)
	} else {
		// ...or otherwise search recursively starting from the current working directory.
		args = append(args, ".")
	}

	return args
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
	opt, optErr := newOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}
	return &Spell{ac: ac, opts: opt}, nil
}
