// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package golang

import (
	"github.com/imdario/mergo"

	"github.com/svengreb/wand/pkg/spell"
)

// Options are shared Go toolchain commands options.
// See:
// - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
type Options struct {
	// AsmFlags are the arguments for the `-asmflags` flag that are passed to each `go tool asm` invocation.
	// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	AsmFlags []string

	// EnableRaceDetector indicates whether the race detector should be enabled.
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableRaceDetector bool

	// EnableTrimPath indicates whether all file system paths should be removed from the resulting executable.
	// This is done by adding compiler and linker flags to remove the absolute path to the project root directory from
	// binary artifacts.
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	//   - https://golang.org/doc/go1.13#go-command
	EnableTrimPath bool

	// Env are Go toolchain command specific environment variables.
	Env map[string]string

	// Flags are additional flags passed to the `go` command.
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	Flags []string

	// FlagsPrefixAll indicates whether the values of `-asmflags` and `-gcflags` should be prefixed with the `all=`
	// pattern in order to apply to all packages.
	// As of Go 1.10 (https://golang.org/doc/go1.10#build), the value specified to `-asmflags` and `-gcflags` are only
	// applied to the current package, therefore the `all=` pattern is used to apply the flag to all packages.
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	//   - https://golang.org/doc/go1.10#build
	FlagsPrefixAll bool

	// GcFlags are the arguments for the `-gcflags` flag that are passed to each `go tool compile` invocation.
	// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Build_modes
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	GcFlags []string

	// LdFlags are the arguments for the `-ldflags` flag that are passed to each `go tool link` invocation.
	// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Build_modes
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	LdFlags []string

	// mixins are spell mixins that can be applied by option consumers.
	mixins []spell.Mixin

	// Tags are the Go tags.
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	Tags []string
}

// Option is a pencil option.
type Option func(*Options)

// WithAsmFlags sets flags to pass on each `go tool asm` invocation.
// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Build_modes
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithAsmFlags(asmFlags ...string) Option {
	return func(o *Options) {
		o.AsmFlags = append(o.AsmFlags, asmFlags...)
	}
}

// WithRaceDetector indicates whether the race detector should be enabled.
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithRaceDetector(enableRaceDetector bool) Option {
	return func(o *Options) {
		o.EnableRaceDetector = enableRaceDetector
	}
}

// WithTrimmedPath indicates whether all file system paths should be removed from the resulting executable.
// This is done by adding compiler and linker flags to remove the absolute path to the project root directory from
// binary artifacts.
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/doc/go1.13#go-command
func WithTrimmedPath(enableTrimPath bool) Option {
	return func(o *Options) {
		o.EnableTrimPath = enableTrimPath
	}
}

// WithEnv adds or overrides Go toolchain command specific environment variables.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		for k, v := range env {
			o.Env[k] = v
		}
	}
}

// WithFlags sets additional Go toolchain command flags.
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithFlags(flags ...string) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}

// WithFlagsPrefixAll indicates whether the values of `-asmflags` and `-gcflags` should be prefixed with the `all=` pattern
// in order to apply to all packages.
// As of Go 1.10 (https://golang.org/doc/go1.10#build), the value specified to `-asmflags` and `-gcflags` are only
// applied to the current package, therefore the `all=` pattern is used to apply the flag to all packages.
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/doc/go1.10#build
func WithFlagsPrefixAll(flagsPrefixAll bool) Option {
	return func(o *Options) {
		o.FlagsPrefixAll = flagsPrefixAll
	}
}

// WithGcFlags sets flags to pass on each `go tool compile` invocation.
// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Build_modes
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithGcFlags(gcFlags ...string) Option {
	return func(o *Options) {
		o.GcFlags = append(o.GcFlags, gcFlags...)
	}
}

// WithLdFlags sets flags to pass on each `go tool link` invocation.
// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Build_modes
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithLdFlags(ldFlags ...string) Option {
	return func(o *Options) {
		o.LdFlags = append(o.LdFlags, ldFlags...)
	}
}

// WithMixins sets spell mixins that can be applied by option consumers.
func WithMixins(mixins ...spell.Mixin) Option {
	return func(o *Options) {
		o.mixins = append(o.mixins, mixins...)
	}
}

// WithTags sets Go toolchain tags.
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithTags(tags ...string) Option {
	return func(o *Options) {
		o.Tags = append(o.Tags, tags...)
	}
}

// NewOptions creates new shared Go toolchain commands options.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		Env: make(map[string]string),
	}
	for _, o := range opts {
		o(opt)
	}

	for _, m := range opt.mixins {
		mixedOpt, mixErr := m.Apply(opt)
		if mixErr != nil {
			continue
		}
		_ = mergo.Merge(opt, mixedOpt)
	}

	return opt
}
