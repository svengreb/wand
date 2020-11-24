// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package gobin provides a caster to install and run Go module executables using the "github.com/myitcv/gobin" module
// command.
// See https://pkg.go.dev/github.com/myitcv/gobin for more details about "gobin".
// The source code of the "gobin" is available at https://github.com/myitcv/gobin.
//
// Go Executable Installation
//
// When installing a Go executable from within a Go module (1) directory using the "go install" command (2), it is
// installed into the Go executable search path that is defined through the "GOBIN" environment variable (3) and can
// also be shown and modified using the "go env" command (4).
// Even though the executable gets installed globally, the "go.mod" file (5) will be updated to include the installed
// packages since this is the default behavior of the "go get" command (6) when running in "module" mode (7).
//
// Next to this problem, the installed executable will also overwrite any executable of the same module/package that was
// installed already, but maybe from a different version. Therefore only one version of a executable can be installed at
// a time which makes it impossible to work on different projects that use the same tool but with different versions.
//
// History and Future
//
// The local installation of executables built from Go modules/packages has always been a somewhat controversial point
// which unfortunately, partly for historical reasons, does not offer an optimal and user-friendly solution up to now.
// The "go" command (8) is a fantastic toolchain that provides many great features one would expect to be provided
// out-of-the-box from a modern and well designed programming language without the requirement to use a third-party
// solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging
// utilities and many more.
// Unfortunately the way the "go install" command (9) of Go versions less or equal to 1.15 handles the installation of
// an Go module/package executable is still not optimal.
//
// The general problem of tool dependencies is a long-time known issue/weak point of the current Go toolchain and is a
// highly rated change request from the Go community with discussions like https://github.com/golang/go/issues/30515,
// https://github.com/golang/go/issues/25922 and https://github.com/golang/go/issues/27653 to improve this essential
// feature, but they've been around for quite a long time without a solution that works without introducing breaking
// changes and most users and the Go team agree on.
// Luckily, this topic was finally picked up for the next upcoming Go release version 1.16 (9) and
// https://github.com/golang/go/issues/40276 introduces a way to install executables in module mode outside a module.
// The release note preview also already includes details about this change (10) and how installation of executables
// from Go modules will be handled in the future.
//
// The Workaround
//
// Beside the great news and anticipation about an official solution for the problem the usage of a workaround is almost
// inevitable until Go 1.16 is finally released.
//
// The official Go wiki (11) provides a section on "How can I track tool dependencies for a module?" (12) that describes
// a workaround that tracks tool dependencies. It allows to use the Go module logic by using a file like "tools.go" with
// a dedicated "tools" build tag that prevents the included module dependencies to be picked up included for normal
// executable builds. This approach works fine for non-main packages, but CLI tools that are only implemented in the
// "main" package can not be imported in such a file.
//
// In order to tackle this problem, a user from the community created "gobin" (13), an experimental, module-aware
// command to install/run main packages.
// It allows to install or run main-package commands without "polluting" the "go.mod" file by default. It downloads
// modules in version-aware mode into a binary cache path within the systems cache directory (14).
// It prevents problems due to already globally installed executables by placing each version in its own directory.
// The decision to use a cache directory instead of sub-directories within the "GOBIN" path keeps the system clean.
//
// "gobin" is still in an early development state, but has already received a lot of positive feedback and is used in
// many projects. There are also members of the core Go team that have contributed to the project and the chance is high
// that the changes for Go 1.16 were influenced or partially ported from it.
// It is currently the best workaround to...
//   1. prevent the Go toolchain to pick up the "GOMOD" environment variable (15) (see "go env GOMOD" (15)) that is
//      initialized automatically with the path to the "go.mod" file in the current working directory.
//   2. install module/package executables globally without "polluting" the "go.mod" file.
//   3. install module/package executables globally without overriding already installed executables of different
//      versions.
//
// See gobin's FAQ page (16) in the repository wiki for more details about the project.
//
// The Go Module Caster
//
// To allow to manage the tool dependency problem, this caster uses "gobin" through to prevent the "pollution" of the
// project "go.mod" file and allows to...
//   1. install "gobin" itself into "GOBIN" (`go env GOBIN` (15)).
//   2. cast any spell incantation (17) of kind "KindGoModule" (18) by installing the executable globally into the
//      dedicated "gobin" cache.
//
// References
//
//   (1) https://golang.org/ref/mod
//   (2) https://golang.org/cmd/go#hdr-Compile_and_install_packages_and_dependencies
//   (3) https://golang.org/cmd/go/#hdr-Environment_variables
//   (4) https://golang.org/cmd/go/#hdr-Print_Go_environment_information
//   (5) https://golang.org/ref/mod#go-mod-file
//   (6) https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them
//   (7) https://golang.org/ref/mod#mod-commands
//   (8) https://golang.org/cmd/go
//   (9) https://github.com/golang/go/milestone/145
//   (10) https://tip.golang.org/doc/go1.16#modules
//   (11) https://github.com/golang/go/wiki
//   (12) https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
//   (13) https://github.com/myitcv/gobin
//   (14) https://golang.org/pkg/os/#UserCacheDir
//   (15) https://golang.org/cmd/go/#hdr-Print_Go_environment_information
//   (16) https://github.com/myitcv/gobin/wiki/FAQ
//   (17) https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Incantation
//   (18) https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#KindGoModule
package gobin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/sh"
	glFS "github.com/svengreb/golib/pkg/io/fs"

	osSupport "github.com/svengreb/wand/internal/support/os"
	"github.com/svengreb/wand/pkg/cast"
	castGoToolchain "github.com/svengreb/wand/pkg/cast/golang/toolchain"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/spell"
)

// Caster is a "github.com/myitcv/gobin" module caster.
type Caster struct {
	opts *Options
}

// GoModule returns partial Go module identifier information for the "github.com/myitcv/gobin" module.
func (c *Caster) GoModule() project.GoModuleID {
	return *c.opts.goModule
}

// GetExec returns the path to the installed executable of the "github.com/myitcv/gobin" module.
func (c *Caster) GetExec() string {
	return c.opts.Exec
}

// Cast casts a spell incantation.
// It returns an error of type *cast.ErrCast when the spell is not a spell.KindGoModule and any other error that occurs
// during the command execution.
func (c *Caster) Cast(si spell.Incantation) error {
	if si.Kind() != spell.KindGoModule {
		return &cast.ErrCast{
			Err:  fmt.Errorf("%q", si.Kind()),
			Kind: cast.ErrCasterSpellIncantationKindUnsupported,
		}
	}

	s, ok := si.(spell.GoModule)
	if !ok {
		return &cast.ErrCast{
			Err:  fmt.Errorf("expected %q but got %q", s.Kind(), si.Kind()),
			Kind: cast.ErrCasterSpellIncantationKindUnsupported,
		}
	}

	args := append([]string{"-run", s.GoModuleID().String()}, si.Formula()...)
	for k, v := range s.Env() {
		c.opts.Env[k] = v
	}

	return sh.RunWithV(c.opts.Env, c.opts.Exec, args...)
}

// Install installs the executable of the "github.com/myitcv/gobin" module.
// It does not "pollute" the "go.mod" file of the project the installation outside of the project root directory but
// using a the systems temporary directory instead.
// See the package documentation for details: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/gobin
func (c *Caster) Install(goCaster *castGoToolchain.Caster) error {
	goToolchainExec := goCaster.GetExec()
	cmd := exec.Command(goToolchainExec, "get", "-v", c.opts.goModule.String())
	cmd.Dir = os.TempDir()
	cmd.Env = os.Environ()

	// Explicitly enable "module" mode to install a pinned "github.com/myitcv/gobin" module version.
	c.opts.Env[castGoToolchain.DefaultEnvVarGO111MODULE] = "on"
	cmd.Env = osSupport.EnvMapToSlice(c.opts.Env)

	if err := cmd.Run(); err != nil {
		return &cast.ErrCast{
			Err:  err,
			Kind: cast.ErrCasterCasting,
		}
	}
	return nil
}

// Handles returns the supported spell.Kind.
func (c *Caster) Handles() spell.Kind {
	return spell.KindGoModule
}

// Validate validates the "github.com/myitcv/gobin" module caster.
// It returns an error of type *cast.ErrCast when the binary executable does not exists at the configured path and when
// it is also not available in the executable search paths of the current environment.
func (c *Caster) Validate() error {
	// Check if the "gobin" executable exists at the configured path,...
	execExits, fsErr := glFS.RegularFileExists(c.opts.Exec)
	if fsErr != nil {
		return &cast.ErrCast{
			Err:  fmt.Errorf("caster %q: %w", CasterName, fsErr),
			Kind: cast.ErrCasterValidation,
		}
	}

	// ...otherwise try to look up the system-wide executable search paths of the current environment...
	if !execExits {
		execPath, pathErr := exec.LookPath(c.opts.Exec)

		// ...and the local Go binary installation path.
		if pathErr != nil {
			var execDirGoEnv string

			if execDirGoEnv = os.Getenv(castGoToolchain.DefaultEnvVarGOBIN); execDirGoEnv == "" {
				if execDirGoEnv = os.Getenv(castGoToolchain.DefaultEnvVarGOPATH); execDirGoEnv != "" {
					execDirGoEnv = filepath.Join(execDirGoEnv, castGoToolchain.DefaultGOBINSubDirName)
				}
			}

			execPath = filepath.Join(execDirGoEnv, c.opts.Exec)
			execExits, fsErr = glFS.RegularFileExists(execPath)
			if fsErr != nil {
				return &cast.ErrCast{
					Err:  fmt.Errorf("caster %q: %w", CasterName, fsErr),
					Kind: cast.ErrCasterValidation,
				}
			}

			if !execExits {
				return &cast.ErrCast{
					Err:  fmt.Errorf("caster %q: %w", CasterName, fsErr),
					Kind: cast.ErrCasterValidation,
				}
			}
		}

		c.opts.Exec = execPath
	}

	return nil
}

// NewCaster creates a new "github.com/myitcv/gobin" module caster.
func NewCaster(opts ...Option) (*Caster, error) {
	opt, optErr := newOptions(opts...)
	if optErr != nil {
		return nil, &cast.ErrCast{
			Err:  optErr,
			Kind: cast.ErrCasterInvalidOpts,
		}
	}

	return &Caster{opts: opt}, nil
}
