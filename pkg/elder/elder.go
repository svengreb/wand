// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package elder is a wand reference implementation that provides common Mage tasks and stores application
// configurations and metadata of a project.
//
// The naming is inspired by the "Elder Wand", an extremely powerful wand made of elder wood, from the fantasy novel
// "Harry Potter". See https://en.wikipedia.org/wiki/Magical_objects_in_Harry_Potter#Elder_Wand for more details.
package elder

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	glFilePath "github.com/svengreb/golib/pkg/io/fs/filepath"
	"github.com/svengreb/nib"

	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
	taskFSClean "github.com/svengreb/wand/pkg/task/fs/clean"
	taskGobin "github.com/svengreb/wand/pkg/task/gobin"
	taskGofumpt "github.com/svengreb/wand/pkg/task/gofumpt"
	taskGoimports "github.com/svengreb/wand/pkg/task/goimports"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoBuild "github.com/svengreb/wand/pkg/task/golang/build"
	taskGoTest "github.com/svengreb/wand/pkg/task/golang/test"
	taskGolangCILint "github.com/svengreb/wand/pkg/task/golangcilint"
	taskGox "github.com/svengreb/wand/pkg/task/gox"
	taskPkger "github.com/svengreb/wand/pkg/task/pkger"
)

// Elder is a wand.Wand reference implementation that provides common Mage tasks and stores configurations and metadata
// for applications of a project.
type Elder struct {
	nib.Nib

	as          app.Store
	gobinRunner *taskGobin.Runner
	goRunner    *taskGo.Runner
	opts        *Options
	project     *project.Metadata
}

// Bootstrap runs initialization tasks to ensure the wand is operational.
// If an error occurs it will be of type *task.ErrRunner.
func (e *Elder) Bootstrap() error {
	if valErr := e.gobinRunner.Validate(); valErr != nil {
		e.Infof("Installing %q", e.gobinRunner.GoMod())
		if installErr := e.gobinRunner.Install(e.goRunner); installErr != nil {
			e.Errorf("Failed to install %q: %v", e.gobinRunner.GoMod(), installErr)
			return fmt.Errorf("install %q: %w", e.gobinRunner.GoMod(), installErr)
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
	t, tErr := taskFSClean.New(e.GetProjectMetadata(), ac, opts...)
	if tErr != nil {
		return []string{}, fmt.Errorf("create %q task: %w", taskFSClean.TaskName, tErr)
	}

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

	return e.goRunner.Run(taskGoBuild.New(e, ac, opts...))
}

// Gofumpt is a task for the "mvdan.cc/gofumpt" Go module command.
// "gofumpt" enforce a stricter format than "https://pkg.go.dev/cmd/gofmt", while being backwards compatible,
// and provides additional rules.
// It is a modified fork of "https://pkg.go.dev/cmd/gofmt" so it can be used as a drop-in replacement.
//
// See the "github.com/svengreb/wand/pkg/task/gofumpt" package for all available options.
// See https://github.com/mvdan/gofumpt#added-rules for more details about available rules.
//
// See https://pkg.go.dev/mvdan.cc/gofumpt for more details about "gofumpt".
// The source code of "gofumpt" is available at https://github.com/mvdan/gofumpt.
func (e *Elder) Gofumpt(appName string, opts ...taskGofumpt.Option) error {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	return e.gobinRunner.Run(taskGofumpt.New(e, ac, opts...))
}

// Goimports is a task for the "golang.org/x/tools/cmd/goimports" Go module command.
// "goimports" allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in
// the same style as "https://pkg.go.dev/cmd/gofmt" so it can be used as a replacement.
//
// See the "github.com/svengreb/wand/pkg/task/goimports" package for all available options.
//
// See https://pkg.go.dev/golang.org/x/tools/cmd/goimports for more details about "goimports".
// The source code of "goimports" is available at https://github.com/golang/tools/tree/master/cmd/goimports.
func (e *Elder) Goimports(appName string, opts ...taskGoimports.Option) error {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	t, tErr := taskGoimports.New(e, ac, opts...)
	if tErr != nil {
		return fmt.Errorf("create %q task: %w", taskGoimports.TaskName, tErr)
	}

	return e.gobinRunner.Run(t)
}

// GolangCILint is a task to run the "github.com/golangci/golangci-lint/cmd/golangci-lint" Go module
// command.
// "golangci-lint" is a fast, parallel runner for dozens of Go linters Go that uses caching, supports YAML
// configurations and has integrations with all major IDEs.
// When any error occurs it will be of type *app.ErrApp or *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/golangcilint" package for all available options.
//
// See https://pkg.go.dev/github.com/golangci/golangci-lint and the official website at https://golangci-lint.run for
// more details about "golangci-lint".
// The source code of "golangci-lint" is available at https://github.com/golangci/golangci-lint.
func (e *Elder) GolangCILint(appName string, opts ...taskGolangCILint.Option) error {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	t, tErr := taskGolangCILint.New(e, ac, opts...)
	if tErr != nil {
		return fmt.Errorf("create %q task: %w", taskGolangCILint.TaskName, tErr)
	}

	return e.gobinRunner.Run(t)
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

	t := taskGoTest.New(e, ac, opts...)
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

	t, tErr := taskGox.New(e, ac, opts...)
	if tErr != nil {
		return fmt.Errorf("create %q task: %w", taskGox.TaskName, tErr)
	}

	return e.gobinRunner.Run(t)
}

// Pkger is a task to run the "github.com/markbates/pkger/cmd/pkger" Go module command.
// "pkger" is a tool for embedding static files into Go binaries.
// When any error occurs it will be of type *app.ErrApp or *task.ErrRunner.
//
// See the "github.com/svengreb/wand/pkg/task/pkger" package for all available options.
//
// See https://pkg.go.dev/github.com/markbates/pkger for more details about "pkger".
// The source code of the "pkger" is available at github.com/markbates/pkger.
//
// Official "Static Assets Embedding"
//
// Please note that the "pkger" project might be superseded and discontinued due to the official Go toolchain support
// for embedding static assets (files) that will most probably be released with Go version 1.16.
//
// Please see https://go.googlesource.com/proposal/+/master/design/draft-embed.md and
// https://github.com/markbates/pkger/issues/114 for more details.
//
// "Monorepo" Workaround
//
// "pkger" tries to mimic the Go standard library and the way how the Go toolchain handles modules, but is therefore
// also affected by its problems and edge cases.
// When the "pkger" command is used from the root of a Go module repository, the directory where the "go.mod" file is
// located, and there is no valid Go source file, the command will fail because it internally uses the same logic like
// the "list" command of the Go toolchain ("go list").
// Therefore a "dummy" Go source file may need to be created as a workaround. This is mostly only required for
// repositories that use a "monorepo" layout where one or more "main" packages are placed in a subdirectory relative to
// the root directory, e.g. "apps" or "cmd". For repositories where the root directory already has a Go package,
// that does not contain any build constraints/tags, or uses a "library" layout, a "dummy" file is probably not needed.
//
// Please see https://github.com/markbates/pkger/issues/109 and https://github.com/markbates/pkger/issues/121 for more
// details.
func (e *Elder) Pkger(appName string, opts ...taskPkger.Option) error {
	ac, acErr := e.GetAppConfig(appName)
	if acErr != nil {
		return fmt.Errorf("get %q application configuration: %w", appName, acErr)
	}

	t, tErr := taskPkger.New(e, ac, opts...)
	if tErr != nil {
		return fmt.Errorf("create %q task: %w", taskPkger.TaskName, tErr)
	}

	dummyWorkaroundFilePath := filepath.Join(
		e.GetProjectMetadata().Options().RootDirPathAbs,
		fmt.Sprintf("%s.go", taskPkger.MonorepoWorkaroundDummyFileName),
	)

	cleanDummyFile := func(path string) {
		if osErr := os.Remove(path); osErr != nil && !errors.Is(osErr, os.ErrNotExist) {
			e.Warnf("Failed to delete \"pkger\" dummy workaround file %q: %w", path, osErr)
			e.Warnf("Please remove %q manually", path)
			return
		}
		e.Debugf("Removed \"pkger\" dummy workaround file %q", path)
	}

	cleanDummyFile(dummyWorkaroundFilePath)
	defer cleanDummyFile(dummyWorkaroundFilePath)

	wErr := ioutil.WriteFile(dummyWorkaroundFilePath, []byte(taskPkger.MonorepoWorkaroundDummyFileContent), os.ModePerm)
	if wErr != nil {
		return fmt.Errorf("write \"pkger\" dummy workaround file %q: %w", dummyWorkaroundFilePath, wErr)
	}

	return e.gobinRunner.Run(t)
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

// Validate ensures that all tasks are properly initialized and operational.
// It returns an error of type *task.ErrRunner when the validation of any of the supported task fails.
func (e *Elder) Validate() error {
	for _, t := range []task.Runner{e.goRunner, e.gobinRunner} {
		if err := t.Validate(); err != nil {
			return fmt.Errorf("failed to validate runner: %w", err)
		}
	}

	return nil
}

// New creates a new elder wand.
//
// The module name is determined automatically using the "runtime/debug" package.
// The absolute path to the root directory is automatically set based on the current working directory.
// Note that the working directory must be set manually when the "magefile" is not placed in the root directory by
// pointing Mage to it:
//   - "-d <PATH>" option to set the directory from which "magefiles" are read (defaults to ".").
//   - "-w <PATH>" option to set the working directory where "magefiles" will run (defaults to value of "-d" flag).
//
// If any error occurs it will be of type *cmd.ErrCmd or *project.ErrProject.
//
// References
//
//   (1) https://magefile.org/#usage
//   (2) https://golang.org/pkg/os/#Getwd
//   (3) https://golang.org/pkg/runtime/debug/#ReadBuildInfo
//   (4) https://pkg.go.dev/runtime/debug
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

	gobinRunner, gobinRunnerErr := taskGobin.NewRunner(e.opts.gobinRunnerOpts...)
	if gobinRunnerErr != nil {
		return nil, fmt.Errorf("failed to create %q runner: %w", "gobin", gobinRunnerErr)
	}
	e.gobinRunner = gobinRunner

	if err := e.RegisterApp(e.project.Options().Name, e.project.Options().DisplayName, project.AppRelPath); err != nil {
		e.ExitPrintf(1, nib.ErrorVerbosity, "registering application %q: %v", e.project.Options().Name, err)
	}

	return e, nil
}
