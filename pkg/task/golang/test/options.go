// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package test

import (
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

const (
	// DefaultIntegrationTestTag is the default tag name for integration tests.
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

	// TaskName is the name of the task.
	TaskName = "go/test"
)

// Option is a task option.
type Option func(*Options)

// Options are parameter build options.
type Options struct {
	*taskGo.Options

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

	// EnableMemoryProfile indicates whether the tests should be run with memory profiling.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	EnableMemoryProfile bool

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

	// taskGoOpts are shared Go toolchain task options.
	taskGoOpts []taskGo.Option

	// TraceProfileOutputFileName is the file name for the execution trace profile file.
	//
	// See `go help test` and the `go` command documentations for more details:
	//   - https://golang.org/cmd/go/#hdr-Testing_flags
	TraceProfileOutputFileName string
}

// NewOptions creates new task options.
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

	opt.Options = taskGo.NewOptions(opt.taskGoOpts...)

	return opt
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

// WithBlockProfileOutputFileName sets the file name for the Goroutine blocking profile file.
// Defaults to DefaultBlockProfileOutputFileName
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithBlockProfileOutputFileName(blockProfileOutputFileName string) Option {
	return func(o *Options) {
		o.BlockProfileOutputFileName = blockProfileOutputFileName
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

// WithCoverageProfileOutputFileName sets the file name for the test coverage profile file.
// Defaults to DefaultCoverageOutputFileName.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithCoverageProfileOutputFileName(coverageProfileOutputFileName string) Option {
	return func(o *Options) {
		o.CoverageProfileOutputFileName = coverageProfileOutputFileName
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

// WithCPUProfileOutputFileName sets the file name for the CPU profile file.
// Defaults to DefaultCPUProfileOutputFileName.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithCPUProfileOutputFileName(cpuProfileOutputFileName string) Option {
	return func(o *Options) {
		o.CPUProfileOutputFileName = cpuProfileOutputFileName
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

// WithGoOptions sets shared Go toolchain task options.
func WithGoOptions(goOpts ...taskGo.Option) Option {
	return func(o *Options) {
		o.taskGoOpts = append(o.taskGoOpts, goOpts...)
	}
}

// WithMemoryProfile indicates whether the tests should be run with memory profiling.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithMemoryProfile(withMemoryProfile bool) Option {
	return func(o *Options) {
		o.EnableMemoryProfile = withMemoryProfile
	}
}

// WithMemoryProfileOutputFileName sets the file name for the memory profile file.
// Defaults to DefaultMemoryProfileOutputFileName.
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
// Defaults to DefaultMutexProfileOutputFileName.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithMutexProfileOutputFileName(mutexProfileOutputFileName string) Option {
	return func(o *Options) {
		o.MutexProfileOutputFileName = mutexProfileOutputFileName
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

// WithOutputDir sets the output directory, relative to the project root, for reports like coverage or benchmark
// profiles.
// Defaults to DefaultOutputDirName.
//
// See `go help test` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Testing_flags
func WithOutputDir(outputDir string) Option {
	return func(o *Options) {
		o.OutputDir = outputDir
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
// Defaults to DefaultTraceProfileOutputFileName.
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
