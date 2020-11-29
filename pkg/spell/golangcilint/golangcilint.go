// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package golangcilint provides a spell incantation for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go
// module command, a fast, parallel runner for dozens of Go linters that uses caching, supports YAML configurations
// and has integrations with all major IDEs.
//
// See https://pkg.go.dev/github.com/golangci/golangci-lint for more details about "golangci-lint".
// The source code of "golangci-lint" is available at
// https://github.com/golangci/golangci-lint/tree/master/cmd/golangci-lint.
package golangcilint

import (
	"github.com/svengreb/wand"
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/spell"
)

// Spell is a spell incantation for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module command.
// If not extra arguments are configured, DefaultArgs are passed to the executable.
type Spell struct {
	ac   app.Config
	opts *Options
}

// Formula returns the spell incantation formula.
func (s *Spell) Formula() []string {
	return s.opts.args
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

// New creates a new spell incantation for the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module command.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(wand wand.Wand, ac app.Config, opts ...Option) (*Spell, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}
	return &Spell{ac: ac, opts: opt}, nil
}
