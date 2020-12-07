// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package golang

import (
	"errors"
	"fmt"

	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

// MixinImproveDebugging is a task.Mixin for golang.Options to add linker flags to improve the debugging of binary
// artifacts.
// This includes the disabling of inlining and all compiler optimizations to improve the compatibility for debuggers.
//
// Note that this mixin adds the "all" prefix for "-gcflags" parameters to make sure all packages are affected.
// If you disabled the "all" prefix on purpose you need to handle this conflict on your own, e.g. by creating more than
// one binary artifact each with different build options.
//
// See `go help build` and `go tool compile -help` for the documentation of supported flags.
//
// References
//
//   - https://golang.org/cmd/link
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/doc/go1.10#build
type MixinImproveDebugging struct{}

// Apply applies the mixin to the task options.
func (m MixinImproveDebugging) Apply(so task.Options) (task.Options, error) {
	goOpts, ok := so.(*Options)
	if !ok {
		return nil, &task.ErrTask{Kind: task.ErrUnsupportedTaskOptions}
	}

	// Make sure the flags are applied to all packages.
	goOpts.FlagsPrefixAll = true
	goOpts.GcFlags = append(goOpts.GcFlags,
		// Disable all compiler optimizations.
		"-N",
		// Disable inlining.
		"-l",
	)

	return goOpts, nil
}

// MixinImproveEscapeAnalysis is a task.Mixin for golang.Options to add linker flags to improve the escape analysis of
// binary artifacts.
// Enables 2/4 level for reporting verbosity, higher levels are too noisy and rarely necessary.
//
// Note that this mixin removes the "all" prefix for "-gcflags" parameters to make sure only the target package is
// affected, otherwise reports for (traverse) dependencies would be included as well. If you enabled the "all" prefix
// on purpose you need to handle this conflict on your own, e.g. by creating more than one binary artifact each with
// different build options.
//
// See `go help build` and `go tool compile -help` for the documentation of supported flags.
//
// References
//
//   - https://golang.org/cmd/link
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/doc/go1.10#build
type MixinImproveEscapeAnalysis struct{}

// Apply applies the mixin to the task options.
func (m MixinImproveEscapeAnalysis) Apply(so task.Options) (task.Options, error) {
	goOpts, ok := so.(*Options)
	if !ok {
		return nil, &task.ErrTask{Kind: task.ErrUnsupportedTaskOptions}
	}

	// Only enable for the target package, otherwise includes reports for (traverse) dependencies as well.
	goOpts.FlagsPrefixAll = false
	goOpts.GcFlags = append(goOpts.GcFlags,
		// Enable 2/4 levels of escape analysis reporting, higher levels are too noisy and rarely necessary.
		"-m",
		"-m",
	)

	return goOpts, nil
}

// MixinInjectBuildTimeVariableValues is a task.Mixin for golang.Options to inject build-time values through the `-X`
// linker flags to populate e.g. application metadata variables.
//
// See `go help build`, `go tool compile -help` and the `go` command documentations for more details:
//   - https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
//   - https://blog.cloudflare.com/setting-go-variables-at-compile-time
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
type MixinInjectBuildTimeVariableValues struct {
	// Data is the map of key/value pairs to inject to variables at build-time.
	// The key must be the path to the variable in form of "<IMPORT_PATH>.<VARIABLE_NAME>",
	// e.g. "pkg/internal/support/app.version".
	// The value is the actual value that will be assigned to the variable, e.g. the application version.
	Data map[string]string

	// GoModule is the identifier of the target Go module to inject the given key/value pairs into.
	GoModule *project.GoModuleID
}

// Apply applies the mixin to the task options.
func (m MixinInjectBuildTimeVariableValues) Apply(so task.Options) (task.Options, error) {
	if m.GoModule == nil {
		return nil, &task.ErrTask{
			Err:  errors.New("module path is required"),
			Kind: task.ErrInvalidTaskOpts,
		}
	}
	goOpts, ok := so.(*Options)
	if !ok {
		return nil, &task.ErrTask{Kind: task.ErrUnsupportedTaskOptions}
	}

	if m.Data != nil {
		for k, v := range m.Data {
			goOpts.LdFlags = append(goOpts.LdFlags, fmt.Sprintf("-X %s/%s=%s", m.GoModule.Path, k, v))
		}
	}

	return goOpts, nil
}

// MixinStripDebugMetadata is a task.Mixin for golang.Options to add linker flags to strip debug information from
// binary artifacts.
// This includes DWARF tables needed for debuggers, but keeps annotations needed for stack traces so panics are still
// readable. It also shrinks the file size and memory overhead as well as reducing the chance for possible security
// related problems due to enabled development features or debug information leaks.
//
// See `go tool compile -help` and `go doc cmd/link` for the documentation of supported flags.
//
// See the official Go documentations and other resources for more details:
//   - https://golang.org/cmd/link
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://rakyll.org/go-tool-flags
//   - https://github.com/golang/go/wiki/CompilerOptimizations
//   - https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick
//
// Note that this mixin adds the "all" prefix for "-gcflags" parameters to make sure all packages are affected
// If you disabled the "all" prefix on purpose you need to handle this conflict on your own, e.g. by creating more than
// one binary artifact each with different build options.
//
// See `go help build`, `go tool compile -help` and the `go` command documentations for more details:
//   - https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies
//   - https://golang.org/doc/go1.10#build
//
// A subsequent optimization could be the usage of UPX (https://github.com/upx/upx) to compress the binary artifact
// afterwards.
type MixinStripDebugMetadata struct{}

// Apply applies the mixin to the task options.
func (m MixinStripDebugMetadata) Apply(so task.Options) (task.Options, error) {
	goOpts, ok := so.(*Options)
	if !ok {
		return nil, &task.ErrTask{Kind: task.ErrUnsupportedTaskOptions}
	}

	// Make sure the flags are applied to all packages.
	goOpts.FlagsPrefixAll = true
	goOpts.LdFlags = append(goOpts.LdFlags,
		// Omit the symbol table and debug information.
		"-s",
		// Omit the DWARF symbol table.
		"-w",
	)

	return goOpts, nil
}
