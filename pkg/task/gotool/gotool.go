// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// +build go1.16

// Package gotool provides a runner to install and run compiled executables of Go module-based "main" packages.
//
// Go Executable Installation
//
// As of Go 1.16 `go install` supports the `pkg@version` syntax [1] which allows to install commands without "polluting"
// a projects `go.mod` file. The resulting executables are placed in the Go executable search path that is defined by
// the "GOBIN" environment variable [2] (see the "go env" command [3] to show or modify the Go toolchain environment).
// The problem is that installed executables will overwrite any previously installed executable of the same
// module/package regardless of the version. Therefore only one version of an executable can be installed at a time
// which makes it impossible to work on different projects that make use of the same executable but require different
// versions.
//
// UX Before Go 1.16
//
// The installation concept for "main" package executables was always a somewhat controversial point which
// unfortunately, partly for historical reasons, did not offer an optimal and user-friendly solution until Go 1.16.
// The "go" command [4] is a fantastic toolchain that provides many great features one would expect to be provided
// out-of-the-box from a modern and well designed programming language without the requirement to use a third-party
// solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging
// utilities and many more.
// This did not apply for the "go install" command [1] of Go versions less than 1.16.
//
// The general problem of tool dependencies was a long-time known issue/weak point of the Go toolchain and was a highly
// rated change request from the Go community with discussions like https://github.com/golang/go/issues/30515,
// https://github.com/golang/go/issues/25922 and https://github.com/golang/go/issues/27653 to improve this essential
// feature. They have been around for quite a long time without a solution that worked without introducing breaking
// changes and most users and the Go team agree on.
// Luckily, this topic was finally resolved in the Go release version 1.16 [5] and
// https://github.com/golang/go/issues/40276 introduced a way to install executables in module mode outside a module.
//
// The Leftover Drawback
//
// Even though the "go install" command works totally fine to globally install executables, the problem that only a
// single version can be installed at a time is still left. The executable is placed in the path defined by
// "go env GOBIN" so the previously installed executable will be overridden. It is not possible to install multiple
// versions of the same package and "go install" still messes up the local user environment.
//
// The Workaround
//
// To work around the leftover drawback, this package provides a runner that uses "go install" under the
// hood, but allows to place the compiled executable in a custom cache directory instead of `go env GOBIN`. It checks if the
// executable already exists, installs it if not so, and executes it afterwards.
//
// The concept of storing dependencies locally on a per-project basis is well-known from the "node_modules"
// directory [6] of the Node [7] package manager npm [8]. Storing executables in a cache directory within the
// repository (not tracked by Git) allows to use "go install" mechanisms while not affect the global user environment
// and executables stored in "go env GOBIN".
// The runner achieves this by temporarily changing the "GOBIN" environment variable to the custom cache directory
// during the execution of "go install".
//
// The only known disadvantage is the increased usage of storage disk space, but since most Go executables are small in
// size anyway, this is perfectly acceptable compared to the clearly outweighing advantages.
//
// Note that the runner dynamically runs executables based on the given task so the "Validate" method is a NOOP.
//
// Future Changes
//
// The provided runner is still not a clean solution that uses the Go toolchain without any special logic so as soon as
// the following changes are made to the Go toolchain (Go 1.17 or later), the runner will be removed again:
//
// - https://github.com/golang/go/issues/42088 tracks the process of adding support for the Go module syntax to the
//   "go run" command. This will allow to let the Go toolchain handle the way how compiled executable are stored,
//   located and executed.
// - https://github.com/golang/go/issues/44469#issuecomment-784534876 tracks the process of making "go install" aware of
//   the "-o" flag like the "go build" command which is the only reason why the provided runner exists.
//
// References
//
//   [1]: https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
//   [2]: https://pkg.go.dev/cmd/go/#hdr-Environment_variables
//   [3]: https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
//   [4]: https://pkg.go.dev/cmd/go
//   [5]: https://golang.org/doc/go1.16#modules
//   [6]: https://docs.npmjs.com/cli/v7/configuring-npm/folders#node-modules
//   [7]: https://nodejs.org
//   [8]: https://www.npmjs.com
package gotool

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	glFS "github.com/svengreb/golib/pkg/io/fs"

	osSupport "github.com/svengreb/wand/internal/support/os"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoInstall "github.com/svengreb/wand/pkg/task/golang/install"
)

// Runner is runner to install and run compiled executables of Go module-based "main" packages.
//
// Go Executable Installation
//
// As of Go 1.16 `go install` supports the `pkg@version` syntax [1] which allows to install commands without "polluting"
// a projects `go.mod` file. The resulting executables are placed in the Go executable search path that is defined by
// the "GOBIN" environment variable [2] (see the "go env" command [3] to show or modify the Go toolchain environment).
// The problem is that installed executables will overwrite any previously installed executable of the same
// module/package regardless of the version. Therefore only one version of an executable can be installed at a time
// which makes it impossible to work on different projects that make use of the same executable but require different
// versions.
//
// UX Before Go 1.16
//
// The installation concept for "main" package executables was always a somewhat controversial point which
// unfortunately, partly for historical reasons, did not offer an optimal and user-friendly solution until Go 1.16.
// The "go" command [4] is a fantastic toolchain that provides many great features one would expect to be provided
// out-of-the-box from a modern and well designed programming language without the requirement to use a third-party
// solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging
// utilities and many more.
// This did not apply for the "go install" command [1] of Go versions less than 1.16.
//
// The general problem of tool dependencies was a long-time known issue/weak point of the Go toolchain and was a highly
// rated change request from the Go community with discussions like https://github.com/golang/go/issues/30515,
// https://github.com/golang/go/issues/25922 and https://github.com/golang/go/issues/27653 to improve this essential
// feature. They have been around for quite a long time without a solution that worked without introducing breaking
// changes and most users and the Go team agree on.
// Luckily, this topic was finally resolved in the Go release version 1.16 [5] and
// https://github.com/golang/go/issues/40276 introduced a way to install executables in module mode outside a module.
//
// The Leftover Drawback
//
// Even though the "go install" command works totally fine to globally install executables, the problem that only a
// single version can be installed at a time is still left. The executable is placed in the path defined by
// "go env GOBIN" so the previously installed executable will be overridden. It is not possible to install multiple
// versions of the same package and "go install" still messes up the local user environment.
//
// The Workaround
//
// To work around the leftover drawback, this runner uses "go install" under the hood, but allows to place the compiled
// executable in a custom cache directory instead of `go env GOBIN`. It checks if the executable already exists,
// installs it if not so, and executes it afterwards.
//
// The concept of storing dependencies locally on a per-project basis is well-known from the "node_modules"
// directory [6] of the Node [7] package manager npm [8]. Storing executables in a cache directory within the
// repository (not tracked by Git) allows to use "go install" mechanisms while not affect the global user environment
// and executables stored in "go env GOBIN".
// The runner achieves this by temporarily changing the "GOBIN" environment variable to the custom cache directory
// during the execution of "go install".
//
// The only known disadvantage is the increased usage of storage disk space, but since most Go executables are small in
// size anyway, this is perfectly acceptable compared to the clearly outweighing advantages.
//
// Note that the runner dynamically runs executables based on the given task so the "Validate" method is a NOOP.
//
// Future Changes
//
// This runner is still not a clean solution that uses the Go toolchain without any special logic so as soon as the
// following changes are made to the Go toolchain (Go 1.17 or later), the runner will be removed again:
//
// - https://github.com/golang/go/issues/42088 tracks the process of adding support for the Go module syntax to the
//   "go run" command. This will allow to let the Go toolchain handle the way how compiled executable are stored,
//   located and executed.
// - https://github.com/golang/go/issues/44469#issuecomment-784534876 tracks the process of making "go install" aware of
//   the "-o" flag like the "go build" command which is the only reason why this runner exists.
//
// References
//
//   [1]: https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
//   [2]: https://pkg.go.dev/cmd/go/#hdr-Environment_variables
//   [3]: https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
//   [4]: https://pkg.go.dev/cmd/go
//   [5]: https://golang.org/doc/go1.16#modules
//   [6]: https://docs.npmjs.com/cli/v7/configuring-npm/folders#node-modules
//   [7]: https://nodejs.org
//   [8]: https://www.npmjs.com
type Runner struct {
	goRunner *taskGo.Runner
	opts     *RunnerOptions
}

// Handles returns the supported task kind.
func (r *Runner) Handles() task.Kind {
	return task.KindGoModule
}

// Install installs the executable of the given Go module.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
//
// // See https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies for more details.
func (r *Runner) Install(goModule *project.GoModuleID) error {
	_, err := r.prepareExec(goModule)
	if err != nil {
		return &task.ErrRunner{
			Err:  fmt.Errorf("runner %q: %w", RunnerName, err),
			Kind: task.ErrRun,
		}
	}

	return nil
}

// Run runs the command.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
func (r *Runner) Run(t task.Task) error {
	tGM, tErr := r.prepareTask(t)
	if tErr != nil {
		return fmt.Errorf("runner %q: %w", RunnerName, tErr)
	}

	execPath, preExecErr := r.prepareExec(tGM.ID())
	if preExecErr != nil {
		return &task.ErrRunner{
			Err:  fmt.Errorf("runner %q: %w", RunnerName, preExecErr),
			Kind: task.ErrRunnerValidation,
		}
	}

	if r.opts.Quiet {
		if err := sh.RunWith(r.opts.Env, execPath, tGM.BuildParams()...); err != nil {
			return &task.ErrRunner{
				Err:  fmt.Errorf("run task %q: %w", t.Name(), err),
				Kind: task.ErrRun,
			}
		}
	}
	if err := sh.RunWithV(r.opts.Env, execPath, tGM.BuildParams()...); err != nil {
		return &task.ErrRunner{
			Err:  fmt.Errorf("run task %q: %w", t.Name(), err),
			Kind: task.ErrRun,
		}
	}
	return nil
}

// RunOut runs the command and returns its output.
// It returns an error of type *task.ErrRunner when any error occurs during the command execution.
func (r *Runner) RunOut(t task.Task) (string, error) {
	tGM, tErr := r.prepareTask(t)
	if tErr != nil {
		return "", fmt.Errorf("runner %q: %w", RunnerName, tErr)
	}

	execPath, preExecErr := r.prepareExec(tGM.ID())
	if preExecErr != nil {
		return "", &task.ErrRunner{
			Err:  fmt.Errorf("runner %q: %w", RunnerName, preExecErr),
			Kind: task.ErrRunnerValidation,
		}
	}

	out, runErr := sh.OutputWith(r.opts.Env, execPath, tGM.BuildParams()...)
	if runErr != nil {
		return "", &task.ErrRunner{
			Err:  fmt.Errorf("run task %q: %w", t.Name(), runErr),
			Kind: task.ErrRun,
		}
	}
	return out, nil
}

// Validate validates the runner.
// This runner uses dynamic executables based on the given task so this method is a NOOP.
func (r *Runner) Validate() error {
	return nil
}

// buildExecDir builds and returns the path to the directory for the executable.
func (r *Runner) buildExecDir(goModule *project.GoModuleID) string {
	path := filepath.Join(r.opts.toolsBinDir, goModule.ExecName())

	if goModule.Version != nil && !goModule.IsLatest {
		return filepath.Join(path, goModule.Version.String())
	}

	return filepath.Join(path, project.GoModuleVersionLatest)
}

// install installs the compiled executable of a Go module-based "main" package.
// It returns an error of type *task.ErrRunner when any error occurs during the installation.
func (r *Runner) install(execDir string, goModule *project.GoModuleID) error {
	env := osSupport.EnvSliceToMap(os.Environ())
	for k, v := range r.opts.Env {
		env[k] = v
	}
	// Override the "GOBIN" environment variable to use the given path for the compiled executable.
	env[taskGo.DefaultEnvVarGOBIN] = execDir

	t := taskGoInstall.New(
		taskGoInstall.WithModulePath(goModule.Path),
		taskGoInstall.WithModuleVersion(goModule.Version),
		taskGoInstall.WithEnv(env),
	)

	if err := r.goRunner.Run(t); err != nil {
		return fmt.Errorf("run %q: %w", t.Name(), err)
	}

	return nil
}

// prepareExec prepares the ensure that the executable exists and returns the path.
func (r *Runner) prepareExec(goModule *project.GoModuleID) (string, error) {
	execDir := r.buildExecDir(goModule)
	execPath := filepath.Join(execDir, goModule.ExecName())

	exists, fsErr := glFS.RegularFileExists(execPath)
	if fsErr != nil {
		return "", fmt.Errorf("check executable %q: %w", execPath, fsErr)
	}

	if !exists {
		if err := os.MkdirAll(execDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("create directory structure %q for execuable: %w", execDir, err)
		}

		if err := r.install(execDir, goModule); err != nil {
			return "", fmt.Errorf("install executable %q: %w", execPath, err)
		}
	}

	return execPath, nil
}

// prepareTask checks if the given task is of type task.GoModule and prepares the task specific environment.
// It returns an error of type *task.ErrRunner when any error occurs during the execution.
func (r *Runner) prepareTask(t task.Task) (task.GoModule, error) {
	tGM, ok := t.(task.GoModule)
	if t.Kind() != task.KindGoModule || !ok {
		return nil, &task.ErrRunner{
			Err:  fmt.Errorf("expected %q but got %q", r.Handles(), t.Kind()),
			Kind: task.ErrUnsupportedTaskKind,
		}
	}

	for k, v := range tGM.Env() {
		r.opts.Env[k] = v
	}

	return tGM, nil
}

// NewRunner creates a new command runner for Go module-based tools.
// It returns an error of type *task.ErrRunner when any error occurs during the creation.
func NewRunner(goRunner *taskGo.Runner, opts ...RunnerOption) (*Runner, error) {
	opt, optErr := NewRunnerOptions(opts...)
	if optErr != nil {
		return nil, &task.ErrRunner{
			Err:  fmt.Errorf("create %q runner options: %w", RunnerName, optErr),
			Kind: task.ErrInvalidRunnerOpts,
		}
	}

	return &Runner{goRunner: goRunner, opts: opt}, nil
}
