// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package cast

import (
	"github.com/svengreb/wand/pkg/spell"
)

// Caster casts a spell.Incantation using a command for a specific spell.Kind.
//
// The abstract view and naming is inspired by the fantasy novel "Harry Potter" in which a caster can cast a magic spell
// through a incantation.
//
// See
//
//   (1) https://en.wikipedia.org/wiki/Magic_in_Harry_Potter#Spellcasting
//   (2) https://en.wikipedia.org/wiki/Incantation
type Caster interface {
	// Cast casts a spell incantation.
	Cast(spell.Incantation) error

	// Handles returns the spell kind that can be casted.
	Handles() spell.Kind

	// Validate validates the caster command.
	Validate() error
}

// BinaryCaster is a Caster to run commands using a binary executable.
type BinaryCaster interface {
	Caster

	// GetExec returns the path to the binary executable of the command.
	GetExec() string
}
