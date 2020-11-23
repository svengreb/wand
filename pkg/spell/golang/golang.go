// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package golang provides spell incantations for Go toolchain commands.
package golang

import (
	"fmt"
	"strings"
)

// CompileFormula compiles the formula for shared Go toolchain command options.
func CompileFormula(opts ...Option) []string {
	opt := NewOptions(opts...)
	var args []string

	if len(opt.Tags) > 0 {
		args = append(args, fmt.Sprintf("-tags='%s'", strings.Join(opt.Tags, " ")))
	}

	if opt.EnableRaceDetector {
		args = append(args, "-race")
	}

	if opt.EnableTrimPath {
		args = append(args, "-trimpath")
	}

	if len(opt.AsmFlags) > 0 {
		flag := "-asmflags"
		if opt.FlagsPrefixAll {
			flag = fmt.Sprintf("%s=all", flag)
		}
		args = append(args, fmt.Sprintf("%s=%s", flag, strings.Join(opt.AsmFlags, " ")))
	}

	if len(opt.GcFlags) > 0 {
		flag := "-gcflags"
		if opt.FlagsPrefixAll {
			flag = fmt.Sprintf("%s=all", flag)
		}
		args = append(args, fmt.Sprintf("%s=%s", flag, strings.Join(opt.GcFlags, " ")))
	}

	if len(opt.LdFlags) > 0 {
		flag := "-ldflags"
		if opt.FlagsPrefixAll {
			flag = fmt.Sprintf("%s=all", flag)
		}
		args = append(args, fmt.Sprintf("%s=%s", flag, strings.Join(opt.LdFlags, " ")))
	}

	if len(opt.Flags) > 0 {
		args = append(args, opt.Flags...)
	}

	return args
}
