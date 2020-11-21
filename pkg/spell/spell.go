// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package spell provides incantations for different kinds.
package spell

import "github.com/svengreb/wand/pkg/project"

// Incantation is the abstract representation of parameters for a command or action.
// It is mainly handled by a cast.Caster that provides the corresponding information about the command like the path
// to the executable.
//
// The separation of parameters from commands enables a flexible usage, e.g. when the parameters can be reused for a
// different command.
//
// The abstract view and naming is inspired by the fantasy novel "Harry Potter" in which it is almost only possible to
// cast a magic spell through a incantation.
//
// See
//
//   (1) https://en.wikipedia.org/wiki/Incantation
//   (2) https://en.wikipedia.org/wiki/Magic_in_Harry_Potter#Spellcasting
//   (3) https://scifi.stackexchange.com/a/33234
//   (4) https://harrypotter.fandom.com/wiki/Spell
//   (5) https://diffsense.com/diff/incantation/spell
type Incantation interface {
	// Formula returns all parameters of a spell.
	Formula() []string

	// Kind returns the Kind of a spell.
	Kind() Kind

	// Options return the options of a spell.
	Options() interface{}
}

// Binary is a Incantation for commands which are using a binary executable.
type Binary interface {
	Incantation

	// Env returns additional environment variables.
	Env() map[string]string
}

// GoCode is a Incantation for actions that can be casted without a cast.Caster.
// It is a special incantations in that it allows to use Go code as spell while still being compatible to the
// incantation API.
// Note that the Incantation.Formula of a GoCode must always return an empty slice,
// otherwise it is a "normal" Incantation that requires a cast.Caster.
//
// Seen from the abstract "Harry Potter" view this is equal to a "non-verbal" spell that is a special technique that can
// be used for spells that have been specially designed to be used non-verbally.
//
// See
//
//   (1) https://en.wikipedia.org/wiki/Magic_in_Harry_Potter#Spellcasting
//   (2) https://www.reddit.com/r/harrypotter/comments/4z9rwl/what_is_the_difference_between_a_spell_charm
type GoCode interface {
	Incantation

	// Cast casts itself.
	Cast() (interface{}, error)
}

// GoModule is a Binary for binary command executables managed by a Go module.
//
// See https://golang.org/ref/mod for more details.
type GoModule interface {
	Binary

	// GoModuleID returns the identifier of a Go module.
	GoModuleID() *project.GoModuleID
}
