// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package vcs

import (
	"fmt"
	"strings"
)

const (
	// KindNameGit is the Kind name for Git repositories.
	KindNameGit = "git"

	// KindNameNone is the Kind name for repositories that are not managed by any VCS.
	KindNameNone = "none"

	// KindNameUnknown is the name for a unknown repository Kind.
	KindNameUnknown = "unknown"
)

const (
	// KindNone is the Kind for repositories that are not managed by any VCS.
	KindNone Kind = iota

	// KindGit is the Kind for Git repositories.
	//
	// See https://git-scm.com for more details.
	KindGit
)

// Kind defines the kind of a vcs.Repository.
type Kind uint32

// MarshalText returns the textual representation of itself.
func (k Kind) MarshalText() ([]byte, error) {
	switch k {
	case KindGit:
		return []byte(KindNameGit), nil
	case KindNone:
		return []byte(KindNameNone), nil
	}

	return nil, fmt.Errorf("not a valid kind %d", k)
}

func (k Kind) String() string {
	if b, err := k.MarshalText(); err == nil {
		return string(b)
	}
	return KindNameUnknown
}

// UnmarshalText implements encoding.TextUnmarshaler to unmarshal a textual representation of itself.
func (k *Kind) UnmarshalText(text []byte) error {
	parsed, err := ParseKind(string(text))
	if err != nil {
		return err
	}

	*k = parsed
	return nil
}

// ParseKind takes a Kind name and returns the Kind constant.
func ParseKind(name string) (Kind, error) {
	switch strings.ToLower(name) {
	case KindNameGit:
		return KindGit, nil
	case KindNameNone:
		return KindNone, nil
	}

	var k Kind
	return k, fmt.Errorf("not a valid kind: %q", name)
}
