// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package os provides utilities for operating system related operations and interactions.
package os

import (
	"fmt"
	"strings"
)

// EnvMapToSlice transforms a map of environment variable key/value pairs into a slice separated by an equal sign.
func EnvMapToSlice(src map[string]string) []string {
	dst := make([]string, len(src))
	for k, v := range src {
		dst = append(dst, fmt.Sprintf("%s=%s", k, v))
	}
	return dst
}

// EnvSliceToMap transforms a slice of environment variables separated by an equal sign into a map.
func EnvSliceToMap(src []string, dst map[string]string) {
	for _, envVar := range src {
		kv := strings.Split(envVar, "=")
		if len(kv) == 1 {
			dst[kv[0]] = ""
		} else {
			dst[kv[0]] = kv[1]
		}
	}
}
