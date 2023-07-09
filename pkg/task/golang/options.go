// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

//go:build go1.17

package golang

import (
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/magefile/mage/mg"

	"github.com/svengreb/wand/pkg/task"
)

const (
	// DefaultEnvVarGO111MODULE is the default environment variable name to toggle the Go 1.11 module mode.
	DefaultEnvVarGO111MODULE = "GO111MODULE"

	// DefaultEnvVarGOBIN is the default environment variable name for the Go binary executable search path.
	DefaultEnvVarGOBIN = "GOBIN"

	// DefaultEnvVarGOFLAGS is the default environment variable name for Go tool flags.
	DefaultEnvVarGOFLAGS = "GOFLAGS"

	// DefaultEnvVarGOPATH is the default environment variable name for the Go path.
	DefaultEnvVarGOPATH = "GOPATH"

	// DefaultGOBINSubDirName is the default name of the subdirectory for the Go executables within DefaultEnvVarGOBIN.
	DefaultGOBINSubDirName = "bin"

	// RunnerName is the name of the runner.
	RunnerName = "golang"
)

// DefaultRunnerExec is the default path to the runner executable.
var DefaultRunnerExec = mg.GoCmd()

// Option is a shared Go toolchain task option.
type Option func(*Options)

// Options are shared Go toolchain task options.
//
// References
//
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
type Options struct {
	// AsmFlags are arguments for the `-asmflags` flag that are passed to each `go tool asm` invocation.
	//
	// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	AsmFlags []string

	// EnableRaceDetector indicates whether the race detector should be enabled.
	//
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableRaceDetector bool

	// EnableTrimPath indicates whether all file system paths should be removed from the resulting executable.
	// This is done by adding compiler and linker flags to remove the absolute path to the project root directory from
	// binary artifacts.
	//
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	//   - https://golang.org/doc/go1.13#go-command
	EnableTrimPath bool

	// Env is the Go toolchain specific environment.
	Env map[string]string

	// Flags are additional flags passed to the `go` command.
	//
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	Flags []string

	// FlagsPrefixAll indicates whether values of `-asmflags` and `-gcflags` should be prefixed with the `all=`
	// pattern in order to apply to all packages.
	// As of Go 1.10 (https://golang.org/doc/go1.10#build), the value specified to `-asmflags` and `-gcflags` are only
	// applied to the current package, therefore the `all=` pattern is used to apply the flag to all packages.
	//
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	//   - https://golang.org/doc/go1.10#build
	FlagsPrefixAll bool

	// GcFlags are arguments for the `-gcflags` flag that are passed to each `go tool compile` invocation.
	//
	// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Build_modes
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	GcFlags []string

	// LdFlags are arguments for the `-ldflags` flag that are passed to each `go tool link` invocation.
	//
	// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Build_modes
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	LdFlags []string

	// mixins are parameter mixins that can be applied by option consumers.
	mixins []task.Mixin

	// Tags are Go tags.
	//
	// See `go help build` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	Tags []string
}

// RunnerOption is a runner option.
type RunnerOption func(*RunnerOptions)

// RunnerOptions are runner options.
type RunnerOptions struct {
	// Env is the runner specific environment.
	Env map[string]string

	// Exec is the name or path of the runner command executable.
	Exec string

	// Quiet indicates whether the runner output should be minimal.
	Quiet bool
}

// BuildGoOptions builds shared Go toolchain options.
func BuildGoOptions(opts ...Option) []string {
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

// NewOptions creates new shared Go toolchain options.
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

// NewRunnerOptions creates new runner options.
func NewRunnerOptions(opts ...RunnerOption) *RunnerOptions {
	opt := &RunnerOptions{
		Env:  make(map[string]string),
		Exec: DefaultRunnerExec,
	}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// WithAsmFlags sets flags to pass on each `go tool asm` invocation.
//
// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Build_modes
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithAsmFlags(asmFlags ...string) Option {
	return func(o *Options) {
		o.AsmFlags = append(o.AsmFlags, asmFlags...)
	}
}

// WithEnv sets the runner specific environment.
func WithEnv(env map[string]string) Option {
	return func(o *Options) {
		for k, v := range env {
			o.Env[k] = v
		}
	}
}

// WithFlags sets additional Go toolchain flags.
//
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithFlags(flags ...string) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}

// WithFlagsPrefixAll indicates whether the values of `-asmflags` and `-gcflags` should be prefixed with the `all=`
// pattern in order to apply to all packages.
// As of Go 1.10 (https://golang.org/doc/go1.10#build), the value specified to `-asmflags` and `-gcflags` are only
// applied to the current package, therefore the `all=` pattern is used to apply the flag to all packages.
//
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/doc/go1.10#build
func WithFlagsPrefixAll(flagsPrefixAll bool) Option {
	return func(o *Options) {
		o.FlagsPrefixAll = flagsPrefixAll
	}
}

// WithGcFlags sets flags to pass on each `go tool compile` invocation.
//
// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Build_modes
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithGcFlags(gcFlags ...string) Option {
	return func(o *Options) {
		o.GcFlags = append(o.GcFlags, gcFlags...)
	}
}

// WithLdFlags sets flags to pass on each `go tool link` invocation.
//
// See `go help buildmode`, `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Build_modes
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithLdFlags(ldFlags ...string) Option {
	return func(o *Options) {
		o.LdFlags = append(o.LdFlags, ldFlags...)
	}
}

// WithMixins sets parameter mixins that can be applied by option consumers.
func WithMixins(mixins ...task.Mixin) Option {
	return func(o *Options) {
		o.mixins = append(o.mixins, mixins...)
	}
}

// WithRaceDetector indicates whether the race detector should be enabled.
//
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithRaceDetector(enableRaceDetector bool) Option {
	return func(o *Options) {
		o.EnableRaceDetector = enableRaceDetector
	}
}

// WithRunnerEnv sets the runner specific environment.
func WithRunnerEnv(env map[string]string) RunnerOption {
	return func(o *RunnerOptions) {
		o.Env = env
	}
}

// WithRunnerExec sets the name or path of the runner command executable.
// Defaults to DefaultRunnerExec.
func WithRunnerExec(nameOrPath string) RunnerOption {
	return func(o *RunnerOptions) {
		o.Exec = nameOrPath
	}
}

// WithRunnerQuiet indicates whether the runner output should be minimal.
func WithRunnerQuiet(quiet bool) RunnerOption {
	return func(o *RunnerOptions) {
		o.Quiet = quiet
	}
}

// WithTags sets Go toolchain tags.
//
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithTags(tags ...string) Option {
	return func(o *Options) {
		o.Tags = append(o.Tags, tags...)
	}
}

// WithTrimmedPath indicates whether all file system paths should be removed from the resulting executable.
// This is done by adding compiler and linker flags to remove the absolute path to the project root directory from
// binary artifacts.
//
// See `go help build` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/doc/go1.13#go-command
func WithTrimmedPath(enableTrimPath bool) Option {
	return func(o *Options) {
		o.EnableTrimPath = enableTrimPath
	}
}
