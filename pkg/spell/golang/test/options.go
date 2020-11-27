// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package test

import (
	spellGo "github.com/svengreb/wand/pkg/spell/golang"
)

const (
	// DefaultIntegrationTestTag is the default name of tag for integration tests.
	DefaultIntegrationTestTag = "integration"

	// DefaultBlockProfileOutputFileName is the default file name for the Goroutine blocking profile file.
	DefaultBlockProfileOutputFileName = "block_profile.out"

	// DefaultCoverageOutputFileName is the default file name for the test coverage profile file.
	DefaultCoverageOutputFileName = "cover_profile.out"

	// DefaultCPUProfileOutputFileName is the default file name for the CPU profile file.
	DefaultCPUProfileOutputFileName = "cpu_profile.out"

	// DefaultMemoryProfileOutputFileName is the default file name for the memory profile file.
	DefaultMemoryProfileOutputFileName = "mem_profile.out"

	// DefaultMutexProfileOutputFileName is the default file name for the mutex profile file.
	DefaultMutexProfileOutputFileName = "mutex_profile.out"

	// DefaultOutputDirName is the default output directory name for test artifacts like profiles and reports.
	DefaultOutputDirName = "test"

	// DefaultTraceProfileOutputFileName is the default file name for the execution trace profile file.
	DefaultTraceProfileOutputFileName = "trace_profile.out"
)

// Options are spell incantation options for the Go toolchain "test" command.
type Options struct {
	*spellGo.Options

	// BlockProfileOutputFileName is the file name for the Goroutine blocking profile file.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	BlockProfileOutputFileName string

	// CoverageProfileOutputFileName is the file name for the test coverage profile file.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	CoverageProfileOutputFileName string

	// CPUProfileOutputFileName is the file name for the CPU profile file.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	CPUProfileOutputFileName string

	// DisableCache indicates whether the tests should be run without test caching that is enabled by Go by default.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	DisableCache bool

	// EnableBlockProfile indicates whether the tests should be run with a Goroutine blocking profiling.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableBlockProfile bool

	// EnableCoverageProfile indicates whether the tests should be run with coverage profiling.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableCoverageProfile bool

	// EnableCPUProfile indicates whether the tests should be run with CPU profiling.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableCPUProfile bool

	// EnableMemProfile indicates whether the tests should be run with memory profiling.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableMemProfile bool

	// EnableMutexProfile indicates whether the tests should be run with mutex profiling.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableMutexProfile bool

	// EnableTraceProfile indicates whether the tests should be run with trace profiling.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableTraceProfile bool

	// EnableVerboseOutput indicates whether the test output should be verbose.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableVerboseOutput bool

	// Flags are additional flags that are passed to the Go `test` command along with the base Go flags.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
	Flags []string

	// MemoryProfileOutputFileName is the file name for the memory profile file.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	MemoryProfileOutputFileName string

	// MutexProfileOutputFileName is the file name for the mutex profile file.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	MutexProfileOutputFileName string

	// OutputDir is the output directory, relative to the project root, for reports like
	// coverage or benchmark profiles.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	OutputDir string

	// Pkgs is a list of packages to test.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	Pkgs []string

	// spellGoOpts are shared Go toolchain command options.
	spellGoOpts []spellGo.Option

	// TraceProfileOutputFileName is the file name for the execution trace profile file.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	TraceProfileOutputFileName string
}

// Option is a spell incantation option for the Go toolchain "test" command.
type Option func(*Options)

// WithBlockProfileOutputFileName sets the file name for the Goroutine blocking profile file.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithBlockProfileOutputFileName(blockProfileOutputFileName string) Option {
	return func(o *Options) {
		o.BlockProfileOutputFileName = blockProfileOutputFileName
	}
}

// WithCoverageProfileOutputFileName sets the file name for the test coverage profile file.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithCoverageProfileOutputFileName(coverageProfileOutputFileName string) Option {
	return func(o *Options) {
		o.CoverageProfileOutputFileName = coverageProfileOutputFileName
	}
}

// WithCPUProfileOutputFileName sets the file name for the CPU profile file.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithCPUProfileOutputFileName(cpuProfileOutputFileName string) Option {
	return func(o *Options) {
		o.CPUProfileOutputFileName = cpuProfileOutputFileName
	}
}

// WithBlockProfile indicates whether the tests should be run with a Goroutine blocking profiling.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithBlockProfile(withBlockProfile bool) Option {
	return func(o *Options) {
		o.EnableBlockProfile = withBlockProfile
	}
}

// WithCoverageProfile indicates whether the tests should be run with coverage profiling.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithCoverageProfile(withCoverageProfile bool) Option {
	return func(o *Options) {
		o.EnableCoverageProfile = withCoverageProfile
	}
}

// WithCPUProfile indicates whether the tests should be run with CPU profiling.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithCPUProfile(withCPUProfile bool) Option {
	return func(o *Options) {
		o.EnableCPUProfile = withCPUProfile
	}
}

// WithFlags sets additional flags that are passed to the Go "test" command along with the shared Go flags.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
func WithFlags(flags ...string) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}

// WithGoOptions sets shared Go toolchain command options.
func WithGoOptions(goOpts ...spellGo.Option) Option {
	return func(o *Options) {
		o.spellGoOpts = append(o.spellGoOpts, goOpts...)
	}
}

// WithMemProfile indicates whether the tests should be run with memory profiling.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithMemProfile(withMemProfile bool) Option {
	return func(o *Options) {
		o.EnableMemProfile = withMemProfile
	}
}

// WithMemoryProfileOutputFileName sets the file name for the memory profile file.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithMemoryProfileOutputFileName(memoryProfileOutputFileName string) Option {
	return func(o *Options) {
		o.MemoryProfileOutputFileName = memoryProfileOutputFileName
	}
}

// WithMutexProfile indicates whether the tests should be run with mutex profiling.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithMutexProfile(withMutexProfile bool) Option {
	return func(o *Options) {
		o.EnableMutexProfile = withMutexProfile
	}
}

// WithMutexProfileOutputFileName sets the file name for the mutex profile file.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithMutexProfileOutputFileName(mutexProfileOutputFileName string) Option {
	return func(o *Options) {
		o.MutexProfileOutputFileName = mutexProfileOutputFileName
	}
}

// WithOutputDir sets the output directory, relative to the project root, for reports like coverage or benchmark
// profiles.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithOutputDir(outputDir string) Option {
	return func(o *Options) {
		o.OutputDir = outputDir
	}
}

// WithoutCache indicates whether the tests should be run without test caching that is enabled by Go by default.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithoutCache(withoutCache bool) Option {
	return func(o *Options) {
		o.DisableCache = withoutCache
	}
}

// WithPkgs sets the list of packages to test.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithPkgs(pkgs ...string) Option {
	return func(o *Options) {
		o.Pkgs = append(o.Pkgs, pkgs...)
	}
}

// WithTraceProfile indicates whether the tests should be run with trace profiling.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithTraceProfile(withTraceProfile bool) Option {
	return func(o *Options) {
		o.EnableTraceProfile = withTraceProfile
	}
}

// WithTraceProfileOutputFileName sets the file name for the execution trace profile file.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithTraceProfileOutputFileName(traceProfileOutputFileName string) Option {
	return func(o *Options) {
		o.TraceProfileOutputFileName = traceProfileOutputFileName
	}
}

// WithVerboseOutput indicates whether the test output should be verbose.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithVerboseOutput(withVerboseOutput bool) Option {
	return func(o *Options) {
		o.EnableVerboseOutput = withVerboseOutput
	}
}

// NewOptions creates new spell incantation options for the Go toolchain "test" command.
func NewOptions(opts ...Option) *Options {
	opt := &Options{
		BlockProfileOutputFileName:    DefaultBlockProfileOutputFileName,
		CoverageProfileOutputFileName: DefaultCoverageOutputFileName,
		CPUProfileOutputFileName:      DefaultCPUProfileOutputFileName,
		MemoryProfileOutputFileName:   DefaultMemoryProfileOutputFileName,
		MutexProfileOutputFileName:    DefaultMutexProfileOutputFileName,
		TraceProfileOutputFileName:    DefaultTraceProfileOutputFileName,
	}
	for _, o := range opts {
		o(opt)
	}

	opt.Options = spellGo.NewOptions(opt.spellGoOpts...)

	return opt
}
