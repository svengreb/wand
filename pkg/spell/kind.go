// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package spell

import (
	"fmt"
	"strings"
)

const (
	// KindNameBinary is the Kind name for binary spells.
	KindNameBinary = "binary"
	// KindNameGoCode is the Kind name for Go code spells.
	KindNameGoCode = "go.code"
	// KindNameGoModule is the Kind name for Go module spells.
	KindNameGoModule = "go.module"
	// KindNameUnknown is the name for a unknown spell Kind.
	KindNameUnknown = "unknown"
)

const (
	// KindBinary is the Kind for binary spells.
	KindBinary Kind = iota
	// KindGoCode is the Kind for Go code spells.
	KindGoCode
	// KindGoModule is the Kind for Go module spells.
	KindGoModule
)

// Kind defines the kind of a spell.
type Kind uint32

// MarshalText returns the textual representation of itself.
func (k Kind) MarshalText() ([]byte, error) {
	switch k {
	case KindBinary:
		return []byte(KindNameBinary), nil
	case KindGoCode:
		return []byte(KindNameGoCode), nil
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
	case KindNameBinary:
		return KindBinary, nil
	case KindNameGoCode:
		return KindGoCode, nil
	case KindNameGoModule:
		return KindGoModule, nil
	}

	var k Kind
	return k, fmt.Errorf("not a valid kind: %q", name)
}
