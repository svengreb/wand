// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

//go:build go1.17

// Package elder is a wand reference implementation that provides common Mage tasks and stores application
// configurations and metadata of a project.
// Note that the elder package requires at least Go 1.17 due to the usage of specific Go command features and behaviors
// that are required for tasks like [github.com/svengreb/wand/pkg/task/golang/run]!
//
// The naming is inspired by the "Elder Wand", an extremely powerful wand made of elder wood, from the fantasy novel
// "Harry Potter". See https://en.wikipedia.org/wiki/Magical_objects_in_Harry_Potter#Elder_Wand for more details.
package elder

import (
	"fmt"
	"os"
	"path/filepath"

	glFilePath "github.com/svengreb/golib/pkg/io/fs/filepath"
	"github.com/svengreb/nib"

	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
	taskFSClean "github.com/svengreb/wand/pkg/task/fs/clean"
	taskGofumpt "github.com/svengreb/wand/pkg/task/gofumpt"
	taskGoimports "github.com/svengreb/wand/pkg/task/goimports"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoBuild "github.com/svengreb/wand/pkg/task/golang/build"
	taskGoTest "github.com/svengreb/wand/pkg/task/golang/test"
	taskGolangCILint "github.com/svengreb/wand/pkg/task/golangcilint"
	taskGoModUpgrade "github.com/svengreb/wand/pkg/task/gomodupgrade"
	taskGoTool "github.com/svengreb/wand/pkg/task/gotool"
	taskGox "github.com/svengreb/wand/pkg/task/gox"
)

// Elder is a wand.Wand reference implementation that provides common Mage tasks and stores configurations and metadata
// for applications of a project.
type Elder struct {
	nib.Nib
	as           app.Store
	goRunner     *taskGo.Runner
	goToolRunner *taskGoTool.Runner
	opts         *Options
	project      *project.Metadata
}

// Bootstrap runs initialization tasks to ensure the wand is operational.
//
// NOTE(Go 1.17): As of version 0.9.0 Bootstrap is a no-op!
// Go 1.17 finally added support to [run commands in module-aware mode] which  makes it obsolete to install the
// executables locally but allows to [run them on-the-fly]. Support for this feature is provided by the
// [the "run" Go task] which can be used as is in simple Mage task functions.
// It returns a slice of errors with type [*task.ErrRunner] containing any error that occurs during the initialization.
//
// Deprecated: As of version 0.9.0 Bootstrap is a no-op and will be removed in version 0.10.0! To install executables
// anyway use the [*Elder.CacheExecutables] method instead. To ensure that the wand is properly initialized and operational
// use the [*Elder.Validate] method.
//
// [run commands in module-aware mode]: https://go.dev/doc/go1.17#go%20run
// [run them on-the-fly]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
// [the "run" Go task]: https://pkg.go.dev/github.com/svengreb/wand@v0.9.0/pkg/task/golang/run
func (e *Elder) Bootstrap(goModuleImportPaths ...string) []error {
	e.Warnf(`As of version 0.9.0 the "Bootstrap"" method has been deprecated and is a no-op! It will be removed in version 0.10.0!
	To install executables anyway use the "CacheExecutables" method instead.
	To ensure that the wand is properly initialized and operational use the "Validate"" method.
	`)
	return []error{}
}

// CacheExecutables installs and caches executables from Go module-based "main" packages into a local cache within the
// working directory. Note that this only works when the [taskGoTool.WithCache] option was set to `true`!
// The given paths must be valid Go module import paths, that can optionally include the version suffix in the
// "pkg@version" format. See [the documentation about the "gotool" task] for more details about the installation
// runner.
// It returns any error that occurs during the execution.
// This method is a kind of workaround for the deprecated [*Elder.Bootstrap] method and allows to still cache command
// executables locally.
//
// [the documentation about the "gotool" task]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gotool
func (e *Elder) CacheExecutables(goModuleImportPaths ...string) error {
	for _, path := range goModuleImportPaths {
		gm, gmErr := project.GoModuleFromImportPath(path)
		if gmErr != nil {
			return gmErr
		}
		if installErr := e.goToolRunner.Install(gm); installErr != nil {
			return installErr
		}
	}
	return nil
}

// Clean is a task to remove filesystem paths, e.g. output data like artifacts and reports from previous development,
// test, production and distribution builds.
// It returns paths that have been cleaned along with an error when the task execution fails.
//
// See the "github.com/svengreb/wand/pkg/task/fs/clean" package for all available options.
func (e *Elder) Clean(appName string, opts ...taskFSClean.Option) ([]string, error) {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return []string{}, fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	t := taskFSClean.New(e.GetProjectMetadata(), ac, opts...)
	return t.Clean()
}

// ExitPrintf simplifies the logging for process exits with a suitable verbosity.
//
// References
//
//   - https://unix.stackexchange.com/questions/418784/what-is-the-min-and-max-values-of-exit-codes-in-linux
func (e *Elder) ExitPrintf(code int, verb nib.Verbosity, format string, args ...interface{}) {
	if code < 0 {
		code = 1
	}
	switch verb {
	case nib.DebugVerbosity:
		e.Debugf(format, args...)
	case nib.ErrorVerbosity:
		e.Errorf(format, args...)
	case nib.FatalVerbosity:
		e.Fatalf(format, args...)
	case nib.InfoVerbosity:
		e.Infof(format, args...)
	case nib.SuccessVerbosity:
		e.Successf(format, args...)
	case nib.WarnVerbosity:
		e.Warnf(format, args...)
	case nib.SuppressVerbosity:
		// noop
	}

	//nolint:revive // Exiting is an explicit logic of this method.
	os.Exit(code)
}

// GetAppConfig returns an application configuration.
// An empty application configuration is returned along with an error of type *app.ErrApp when there is no configuration
// in the store for the given name.
func (e *Elder) GetAppConfig(name string) (app.Config, error) {
	ac, acErr := e.as.Get(name)
	if acErr != nil {
		return app.Config{}, fmt.Errorf("get %q application configuration: %w", name, acErr)
	}

	return *ac, nil
}

// GetProjectMetadata returns metadata of the project.
func (e *Elder) GetProjectMetadata() project.Metadata {
	return *e.project
}

// GoBuild is a task for the Go toolchain "build" command.
// When any error occurs it will be of type *app.ErrApp or *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/golang/build" package for all available options.
func (e *Elder) GoBuild(appName string, opts ...taskGoBuild.Option) error {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	return e.goRunner.Run(taskGoBuild.New(ac, opts...))
}

// Gofumpt is a task for the "mvdan.cc/gofumpt" Go module command.
// "gofumpt" enforce a stricter format than "https://pkg.go.dev/cmd/gofmt", while being backwards compatible,
// and provides additional rules.
// It is a modified fork of "https://pkg.go.dev/cmd/gofmt" so it can be used as a drop-in replacement.
// When any error occurs it will be of type *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/gofumpt" package for all available options.
// See https://github.com/mvdan/gofumpt#added-rules for more details about available rules.
//
// See https://pkg.go.dev/mvdan.cc/gofumpt for more details about "gofumpt".
// The source code of "gofumpt" is available at https://github.com/mvdan/gofumpt.
func (e *Elder) Gofumpt(opts ...taskGofumpt.Option) error {
	t, tErr := taskGofumpt.New(opts...)
	if tErr != nil {
		return fmt.Errorf(`create "gofumpt" task: %w`, tErr)
	}

	return e.goToolRunner.Run(t)
}

// Goimports is a task for the "golang.org/x/tools/cmd/goimports" Go module command.
// "goimports" allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in
// the same style as "https://pkg.go.dev/cmd/gofmt" so it can be used as a replacement.
// When any error occurs it will be of type *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/goimports" package for all available options.
//
// See https://pkg.go.dev/golang.org/x/tools/cmd/goimports for more details about "goimports".
// The source code of "goimports" is available at https://github.com/golang/tools/tree/master/cmd/goimports.
func (e *Elder) Goimports(opts ...taskGoimports.Option) error {
	t, tErr := taskGoimports.New(opts...)
	if tErr != nil {
		return fmt.Errorf(`create "goimports" task: %w`, tErr)
	}

	return e.goToolRunner.Run(t)
}

// GolangCILint is a task to run the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module
// command.
// "golangci-lint" is a fast, parallel runner for dozens of Go linters Go that uses caching, supports YAML
// configurations and has integrations with all major IDEs.
// When any error occurs it will be of type *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/golangcilint" package for all available options.
//
// See https://pkg.go.dev/github.com/golangci/golangci-lint and the official website at https://golangci-lint.run for
// more details about "golangci-lint".
// The source code of "golangci-lint" is available at https://github.com/golangci/golangci-lint.
func (e *Elder) GolangCILint(opts ...taskGolangCILint.Option) error {
	t, tErr := taskGolangCILint.New(opts...)
	if tErr != nil {
		return fmt.Errorf(`create "golangci-lint" task: %w`, tErr)
	}

	return e.goToolRunner.Run(t)
}

// GoModUpgrade is a task for the "github.com/oligot/go-mod-upgrade" Go module command.
// "go-mod-upgrade" allows to update outdated Go module dependencies interactively.
// When any error occurs it will be of type *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/gomodupgrade" package for all available options.
//
// See https://pkg.go.dev/github.com/oligot/go-mod-upgrade for more details about "go-mod-upgrade".
// The source code of "go-mod-upgrade" is available at https://github.com/oligot/go-mod-upgrade.
func (e *Elder) GoModUpgrade(opts ...taskGoModUpgrade.Option) error {
	t, tErr := taskGoModUpgrade.New(opts...)
	if tErr != nil {
		return fmt.Errorf(`create "gomodupgrade" task: %w`, tErr)
	}

	return e.goToolRunner.Run(t)
}

// GoTest is a task to run the Go toolchain "test" command.
// The configured output directory for reports like coverage or benchmark profiles will be created recursively when it
// does not exist yet.
// When any error occurs it will be of type *app.ErrApp, *task.ErrRunner or os.PathError.
//
// See the "github.com/svengreb/wand/pkg/task/param/golang/test" package for all available options.
func (e *Elder) GoTest(appName string, opts ...taskGoTest.Option) error {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	t := taskGoTest.New(ac, opts...)
	tOpts, ok := t.Options().(taskGoTest.Options)
	if !ok {
		return fmt.Errorf(`convert task options to "%T"`, taskGoTest.Options{})
	}

	if err := os.MkdirAll(tOpts.OutputDir, os.ModePerm); err != nil {
		return fmt.Errorf("create output directory %q: %w", tOpts.OutputDir, err)
	}

	return e.goRunner.Run(t)
}

// Gox is a task to run the "github.com/mitchellh/gox" Go module command.
// "gox" is a dead simple, no frills Go cross compile tool that behaves a lot like the standard Go toolchain "build"
// command.
// When any error occurs it will be of type *app.ErrApp or *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/gox" package for all available options.
//
// See https://pkg.go.dev/github.com/mitchellh/gox for more details about "gox".
// The source code of the "gox" is available at https://github.com/mitchellh/gox.
func (e *Elder) Gox(appName string, opts ...taskGox.Option) error {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	t, tErr := taskGox.New(ac, opts...)
	if tErr != nil {
		return fmt.Errorf(`create "gox" task: %w`, tErr)
	}

	return e.goToolRunner.Run(t)
}

// RegisterApp creates and stores a new application configuration.
// Note that the package path must be relative to the project root directory!
//
// It returns an error of type *app.ErrApp when the application path is not relative to the project root directory,
// when it is not a subdirectory of it or when any other error occurs.
func (e *Elder) RegisterApp(name, displayName, pathRel string) error {
	// Ensure the application name is valid...
	if name == "" {
		return &app.ErrApp{Kind: app.ErrEmptyName}
	}
	// ...and use it as fallback when the display name has not been set explicitly.
	if displayName == "" {
		displayName = name
	}

	if filepath.IsAbs(pathRel) {
		return &app.ErrApp{
			Err:  fmt.Errorf("check application path %q", pathRel),
			Kind: app.ErrPathNotRelative,
		}
	}

	isSubDir, subDirErr := glFilePath.IsSubDir(e.project.Options().RootDirPathAbs, pathRel, true)
	if subDirErr != nil {
		return &app.ErrApp{
			Err: fmt.Errorf(
				"check if %q is a subdirectory of %q: %w",
				pathRel, e.project.Options().RootDirPathAbs, subDirErr,
			),
		}
	}
	if !isSubDir {
		return &app.ErrApp{
			Err:  fmt.Errorf("%q is not a subdirectory of %q", pathRel, e.project.Options().RootDirPathAbs),
			Kind: app.ErrNonProjectRootSubDir,
		}
	}

	ac := &app.Config{
		BaseOutputDir: filepath.Join(e.project.Options().BaseOutputDir, pathRel),
		DisplayName:   displayName,
		Name:          name,
		PathRel:       pathRel,
		PkgImportPath: filepath.Clean(fmt.Sprintf("%s/%s", e.project.Options().GoModule.Path, pathRel)),
	}

	e.as.Add(ac)
	return nil
}

// Validate ensures that the wand is properly initialized and operational.
// Optionally pass the [task.Runner] that should be validated or nothing to validate all currently supported runners.
// It returns a slice of errors that occurred during the execution.
func (e *Elder) Validate(runners ...task.Runner) []error {
	if len(runners) == 0 || runners == nil {
		runners = []task.Runner{e.goRunner, e.goToolRunner}
	}
	var errs []error
	for _, r := range runners {
		if err := r.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("failed to validate runner: %w", err))
		}
	}
	return errs
}

// New creates a new elder wand.
//
// The module name is determined automatically using the "runtime/debug" package.
// The absolute path to the root directory is automatically set based on the current working directory.
// When the WithDisableAutoGenWandDataDir option is set to `false` the auto-generation of the directory for wand
// specific data will be disabled.
// Note that the working directory must be set manually when the "magefile" is not placed in the root directory by
// pointing Mage to it:
//   - "-d <PATH>" option to set the directory from which "magefiles" are read (defaults to ".").
//   - "-w <PATH>" option to set the working directory where "magefiles" will run (defaults to value of "-d" flag).
//
// If any error occurs it will be of type *cmd.ErrCmd or *project.ErrProject.
//
// References:
//
//  1. https://magefile.org/#usage
//  2. https://golang.org/pkg/os/#Getwd
//  3. https://golang.org/pkg/runtime/debug/#ReadBuildInfo
//  4. https://pkg.go.dev/runtime/debug
func New(opts ...Option) (*Elder, error) {
	opt := NewOptions(opts...)

	e := &Elder{
		as:   app.NewStore(),
		opts: opt,
	}
	e.Nib = e.opts.nib

	proj, projErr := project.New(e.opts.projectOpts...)
	if projErr != nil {
		return nil, fmt.Errorf("failed to create project metadata: %w", projErr)
	}
	e.project = proj

	e.goRunner = taskGo.NewRunner(e.opts.goRunnerOpts...)

	goToolRunnerOpts := append(
		[]taskGoTool.RunnerOption{
			taskGoTool.WithToolsBinDir(filepath.Join(e.project.Options().WandDataDir, taskGoTool.DefaultGoToolsBinDir)),
			taskGoTool.WithQuiet(true),
		},
		e.opts.goToolRunnerOpts...,
	)
	goToolRunner, goToolRunnerErr := taskGoTool.NewRunner(e.goRunner, goToolRunnerOpts...)
	if goToolRunnerErr != nil {
		return nil, fmt.Errorf("create %q runner: %w", taskGoTool.RunnerName, goToolRunnerErr)
	}
	e.goToolRunner = goToolRunner

	if !e.opts.disableAutoGenWandDataDir {
		if err := generateWandDataDir(e.project.Options().WandDataDir); err != nil {
			return nil, fmt.Errorf("generate wand specific data directory %q: %w", e.project.Options().WandDataDir, err)
		}
	}

	if err := e.RegisterApp(e.project.Options().Name, e.project.Options().DisplayName, project.AppRelPath); err != nil {
		e.ExitPrintf(1, nib.ErrorVerbosity, "registering application %q: %v", e.project.Options().Name, err)
	}

	return e, nil
}
