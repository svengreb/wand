// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package task

import (
	"fmt"
	"strings"
)

const (
	// KindNameBase is the kind name for base tasks.
	KindNameBase = "base"
	// KindNameExec is the kind name for executable file tasks.
	KindNameExec = "executable"
	// KindNameGoModule is the kind name for Go module tasks.
	KindNameGoModule = "go.module"
	// KindNameUnknown is the name for a unknown task kind.
	KindNameUnknown = "unknown"
)

const (
	// KindBase is the kind for base tasks.
	KindBase Kind = iota
	// KindExec is the kind for executable file tasks.
	KindExec
	// KindGoModule is the kind for Go module tasks.
	KindGoModule
)

// Kind defines the kind of tasks.
type Kind uint32

// MarshalText returns the textual representation of itself.
func (k Kind) MarshalText() ([]byte, error) {
	switch k {
	case KindBase:
		return []byte(KindNameBase), nil
	case KindExec:
		return []byte(KindNameExec), nil
	case KindGoModule:
		return []byte(KindNameGoModule), nil
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

// ParseKind takes a kind name and returns the Kind constant.
func ParseKind(name string) (Kind, error) {
	switch strings.ToLower(name) {
	case KindNameBase:
		return KindBase, nil
	case KindNameExec:
		return KindExec, nil
	case KindNameGoModule:
		return KindGoModule, nil
	}

	var k Kind
	return k, fmt.Errorf("not a valid kind: %q", name)
}
