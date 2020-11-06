// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package gobin provides a runner for the "github.com/myitcv/gobin" Go module command.
// "gobin" is an experimental, module-aware command to install and run "main" packages.
//
// See https://pkg.go.dev/github.com/myitcv/gobin for more details about "gobin".
// The source code of the "gobin" is available at https://github.com/myitcv/gobin.
//
// Go Executable Installation
//
// Using the "go install" (2) or "go get" (6) command for a Go module (1) "main" package, the resulting executables are
// placed in the Go executable search path that is defined by the "GOBIN" environment variable (3) (see the "go env"
// command (4) to show or modify the Go toolchain environment).
// Even though executables are installed "globally" for the current user, any "go.mod" file (5) in the current working
// directory will be updated to include the Go module. This is the default behavior of the "go get" command (6) when
// running in "module mode" (7) (see "GO111MODULE" environment variable).
//
// Next to this problem, installed executables will also overwrite any previously installed executable of the same
// module/package regardless of the version. Therefore only one version of a executable can be installed at a time which
// makes it impossible to work on different projects that make use of the same executable but require different
// versions.
//
// History and Future
//
// The installation concept for "main" package executables has always been a somewhat controversial point which
// unfortunately, partly for historical reasons, does not offer an optimal and user-friendly solution up to now.
// The "go" command (8) is a fantastic toolchain that provides many great features one would expect to be provided
// out-of-the-box from a modern and well designed programming language without the requirement to use a third-party
// solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging
// utilities and many more.
// Unfortunately this doesn't apply for the "go install" command (2) of Go versions less or equal to 1.15.
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
// a dedicated "tools" build tag that prevents the included module dependencies to be picked up for "normal" executable
// builds. This approach works fine for non-main packages, but CLI tools that are only implemented in the
// "main" package can not be imported in such a file.
//
// In order to tackle this problem, a well-known user from the community created "gobin" (13), an experimental,
// module-aware command to install and run "main" packages.
// It allows to install or run "main" package commands without "polluting" the "go.mod" file. Modules are downloaded in
// version-aware mode into a cache path within the users local cache directory (14). This way it prevents problems due
// to already installed executables by placing each version of an executable in its own directory.
// The decision to use a cache directory instead of sub-directories within the "GOBIN" path doesn't require to mess with
// the users setup and keep the Go toolchain specific paths clean and unchanged.
//
// "gobin" is still in an early development state, but has already received a lot of positive feedback and is used in
// many projects. There are also members of the core Go team that have contributed to the project and the chance is high
// that the changes for Go 1.16 were influenced or partially ported from it.
// It is currently the best workaround to...
//   1. prevent the Go toolchain to pick up the "GOMOD" ("go env GOMOD") environment variable (4) that is initialized
//      automatically with the path to the "go.mod" file (5) in the current working directory.
//   2. install "main" package executables locally for the current user without "polluting" the "go.mod" file.
//   3. install "main" package executables locally for the current user without overriding already installed executables
//      of different versions.
//
// See gobin's FAQ page (15) in the repository wiki for more details about the project.
//
// The Go Module Command Runner
//
// To allow to manage the tool dependency problem, this package provides a command runner that uses "gobin" in order to
// prevent the problems described in the sections above like the "pollution" of the "go.mod" file and allows to...
//   1. install "gobin" itself into "GOBIN" (`go env GOBIN` (4)).
//   2. run any Go module command by installing "main" package executables locally for the current user into the
//      dedicated "gobin" cache.
//
// References
//
//   (1) https://golang.org/ref/mod
//   (2) https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
//   (3) https://pkg.go.dev/cmd/go/#hdr-Environment_variables
//   (4) https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
//   (5) https://golang.org/ref/mod#go-mod-file
//   (6) https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
//   (7) https://golang.org/ref/mod#mod-commands
//   (8) https://golang.org/cmd/go
//   (9) https://github.com/golang/go/milestone/145
//   (10) https://tip.golang.org/doc/go1.16#modules
//   (11) https://github.com/golang/go/wiki
//   (12) https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
//   (13) https://github.com/myitcv/gobin
//   (14) https://pkg.go.dev/os/#UserCacheDir
//   (15) https://github.com/myitcv/gobin/wiki/FAQ
package gobin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/sh"
	glFS "github.com/svengreb/golib/pkg/io/fs"

	osSupport "github.com/svengreb/wand/internal/support/os"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

// Runner is a runner for the "github.com/myitcv/gobin" Go module command.
// "gobin" is an experimental, module-aware command to install and run "main" packages.
//
// See https://pkg.go.dev/github.com/myitcv/gobin for more details about "gobin".
// The source code of the "gobin" is available at https://github.com/myitcv/gobin.
//
// Go Executable Installation
//
// Using the "go install" (2) or "go get" (6) command for a Go module (1) "main" package, the resulting executables are
// placed in the Go executable search path that is defined by the "GOBIN" environment variable (3) (see the "go env"
// command (4) to show or modify the Go toolchain environment).
// Even though executables are installed "globally" for the current user, any "go.mod" file (5) in the current working
// directory will be updated to include the Go module. This is the default behavior of the "go get" command (6) when
// running in "module mode" (7) (see "GO111MODULE" environment variable).
//
// Next to this problem, installed executables will also overwrite any previously installed executable of the same
// module/package regardless of the version. Therefore only one version of a executable can be installed at a time which
// makes it impossible to work on different projects that make use of the same executable but require different
// versions.
//
// History and Future
//
// The installation concept for "main" package executables has always been a somewhat controversial point which
// unfortunately, partly for historical reasons, does not offer an optimal and user-friendly solution up to now.
// The "go" command (8) is a fantastic toolchain that provides many great features one would expect to be provided
// out-of-the-box from a modern and well designed programming language without the requirement to use a third-party
// solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging
// utilities and many more.
// Unfortunately this doesn't apply for the "go install" command (2) of Go versions less or equal to 1.15.
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
// a dedicated "tools" build tag that prevents the included module dependencies to be picked up for "normal" executable
// builds. This approach works fine for non-main packages, but CLI tools that are only implemented in the
// "main" package can not be imported in such a file.
//
// In order to tackle this problem, a well-known user from the community created "gobin" (13), an experimental,
// module-aware command to install and run "main" packages.
// It allows to install or run "main" package commands without "polluting" the "go.mod" file. Modules are downloaded in
// version-aware mode into a cache path within the users local cache directory (14). This way it prevents problems due
// to already installed executables by placing each version of an executable in its own directory.
// The decision to use a cache directory instead of sub-directories within the "GOBIN" path doesn't require to mess with
// the users setup and keep the Go toolchain specific paths clean and unchanged.
//
// "gobin" is still in an early development state, but has already received a lot of positive feedback and is used in
// many projects. There are also members of the core Go team that have contributed to the project and the chance is high
// that the changes for Go 1.16 were influenced or partially ported from it.
// It is currently the best workaround to...
//   1. prevent the Go toolchain to pick up the "GOMOD" ("go env GOMOD") environment variable (4) that is initialized
//      automatically with the path to the "go.mod" file (5) in the current working directory.
//   2. install "main" package executables locally for the current user without "polluting" the "go.mod" file.
//   3. install "main" package executables locally for the current user without overriding already installed executables
//      of different versions.
//
// See gobin's FAQ page (15) in the repository wiki for more details about the project.
//
// The Go Module Command Runner
//
// To allow to manage the tool dependency problem, this package provides a command runner that uses "gobin" in order to
// prevent the problems described in the sections above like the "pollution" of the "go.mod" file and allows to...
//   1. install "gobin" itself into "GOBIN" (`go env GOBIN` (4)).
//   2. run any Go module command by installing "main" package executables locally for the current user into the
//      dedicated "gobin" cache.
//
// References
//
//   (1) https://golang.org/ref/mod
//   (2) https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
//   (3) https://pkg.go.dev/cmd/go/#hdr-Environment_variables
//   (4) https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
//   (5) https://golang.org/ref/mod#go-mod-file
//   (6) https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
//   (7) https://golang.org/ref/mod#mod-commands
//   (8) https://golang.org/cmd/go
//   (9) https://github.com/golang/go/milestone/145
//   (10) https://tip.golang.org/doc/go1.16#modules
//   (11) https://github.com/golang/go/wiki
//   (12) https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
//   (13) https://github.com/myitcv/gobin
//   (14) https://pkg.go.dev/os/#UserCacheDir
//   (15) https://github.com/myitcv/gobin/wiki/FAQ
type Runner struct {
	opts *RunnerOptions
}

// FilePath returns the path to the runner executable.
func (r *Runner) FilePath() string {
	return r.opts.Exec
}

// GoMod returns the Go module identifier.
func (r *Runner) GoMod() *project.GoModuleID {
	return r.opts.goModule
}

// Handles returns the supported task kind.
func (r *Runner) Handles() task.Kind {
	return task.KindGoModule
}

// Install installs the runner executable.
// It does not "pollute" the "go.mod" file of the project by running the installation outside of the project root
// directory using a temporary path instead.
//
// See the package documentation for details: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gobin
func (r *Runner) Install(goRunner *taskGo.Runner) error {
	goRunnerExec := goRunner.FilePath()
	executor := exec.Command(goRunnerExec, "get", "-v", r.opts.goModule.String())
	executor.Dir = os.TempDir()
	executor.Env = os.Environ()

	// Explicitly enable "module" mode to install a pinned version.
	r.opts.Env[taskGo.DefaultEnvVarGO111MODULE] = "on"
	executor.Env = osSupport.EnvMapToSlice(r.opts.Env)

	if err := executor.Run(); err != nil {
		return &task.ErrRunner{
			Err:  err,
			Kind: task.ErrRun,
		}
	}
	return nil
}

// Run runs the command.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
func (r *Runner) Run(t task.Task) error {
	tGM, ok := t.(task.GoModule)
	if t.Kind() != task.KindGoModule || !ok {
		return &task.ErrRunner{
			Err:  fmt.Errorf("expected %q but got %q", r.Handles(), t.Kind()),
			Kind: task.ErrUnsupportedTaskKind,
		}
	}

	runFn := sh.RunWithV
	if r.opts.Quiet {
		runFn = sh.RunWith
	}

	params := append([]string{"-run", tGM.ID().String()}, tGM.BuildParams()...)

	for k, v := range tGM.Env() {
		r.opts.Env[k] = v
	}

	return runFn(r.opts.Env, r.opts.Exec, params...)
}

// Validate validates the command executable.
// It returns an error of type *task.ErrRunner when the executable does not exist and when it is also not available in
// the executable search path(s) of the current environment.
func (r *Runner) Validate() error {
	// Check if the executable exists,...
	execExits, fsErr := glFS.RegularFileExists(r.opts.Exec)
	if fsErr != nil {
		return &task.ErrRunner{
			Err:  fmt.Errorf("runner %q: %w", RunnerName, fsErr),
			Kind: task.ErrRunnerValidation,
		}
	}

	// ...otherwise try to look up the executable search path(s)...
	if !execExits {
		execPath, pathErr := exec.LookPath(r.opts.Exec)

		// ...and the Go specific executable search path.
		if pathErr != nil {
			var execDirGoEnv string

			if execDirGoEnv = os.Getenv(taskGo.DefaultEnvVarGOBIN); execDirGoEnv == "" {
				if execDirGoEnv = os.Getenv(taskGo.DefaultEnvVarGOPATH); execDirGoEnv != "" {
					execDirGoEnv = filepath.Join(execDirGoEnv, taskGo.DefaultGOBINSubDirName)
				}
			}

			execPath = filepath.Join(execDirGoEnv, r.opts.Exec)
			execExits, fsErr = glFS.RegularFileExists(execPath)
			if fsErr != nil {
				return &task.ErrRunner{
					Err:  fmt.Errorf("runner %q: %w", RunnerName, fsErr),
					Kind: task.ErrRunnerValidation,
				}
			}

			if !execExits {
				return &task.ErrRunner{
					Err:  fmt.Errorf("runner %q: %q not found or does not exist", RunnerName, execPath),
					Kind: task.ErrRunnerValidation,
				}
			}
		}

		r.opts.Exec = execPath
	}

	return nil
}

// NewRunner creates a new command runner for the "github.com/myitcv/gobin" Go module.
func NewRunner(opts ...RunnerOption) (*Runner, error) {
	opt, optErr := NewRunnerOptions(opts...)
	if optErr != nil {
		return nil, &task.ErrRunner{
			Err:  optErr,
			Kind: task.ErrInvalidRunnerOpts,
		}
	}

	return &Runner{opts: opt}, nil
}
