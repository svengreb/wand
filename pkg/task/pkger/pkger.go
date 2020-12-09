// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package pkger provides a task for the "github.com/markbates/pkger/cmd/pkger" Go module command.
// "pkger" is a tool for embedding static files into Go binaries.
//
// See https://pkg.go.dev/github.com/markbates/pkger for more details about "pkger".
// The source code of "goimports" is available at https://github.com/markbates/pkger.
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
package pkger

import (
	"fmt"
	"path/filepath"

	glFilePath "github.com/svengreb/golib/pkg/io/fs/filepath"

	"github.com/svengreb/wand"
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task for the "github.com/markbates/pkger/cmd/pkger" Go module command.
// "pkger" is a tool for embedding static files into Go binaries.
//
// See https://pkg.go.dev/github.com/markbates/pkger for more details about "pkger".
// The source code of "goimports" is available at https://github.com/markbates/pkger.
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
type Task struct {
	ac   app.Config
	opts *Options
}

// BuildParams builds the parameters.
func (t *Task) BuildParams() []string {
	var params []string

	// Adds all manually configured paths of files and directories that should be included.
	// By default the paths will be detected by "pkger" itself when used within any of the packages of the target module.
	for _, p := range t.opts.includePathsRel {
		params = append(params, "-include", p)
	}

	if t.opts.outputDirRel != "" {
		params = append(params, "-o", t.opts.outputDirRel)
	}

	return params
}

// Env returns the task specific environment.
func (t *Task) Env() map[string]string {
	return t.opts.env
}

// ID returns the identifier of the Go module.
func (t *Task) ID() *project.GoModuleID {
	return t.opts.goModule
}

// Kind returns the task kind.
func (t *Task) Kind() task.Kind {
	return task.KindGoModule
}

// Options returns the task options.
func (t *Task) Options() task.Options {
	return *t.opts
}

// New creates a new task for the "github.com/markbates/pkger/cmd/pkger" Go module command.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(w wand.Wand, ac app.Config, opts ...Option) (*Task, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}

	// Ensure to place the generated "pkged.go" file in the "main" package directory when the application is not the root
	// directory of the Go module.
	if ac.PathRel != project.AppRelPath {
		opt.outputDirRel = ac.PathRel
	}

	for idx, p := range opt.includePathsRel {
		if filepath.IsAbs(p) {
			return nil, &task.ErrTask{
				Err:  fmt.Errorf("\"include\" path is not relative: %s", p),
				Kind: task.ErrInvalidTaskOpts,
			}
		}

		rootDir := w.GetProjectMetadata().Options().RootDirPathAbs
		isSubDir, fsErr := glFilePath.IsSubDir(rootDir, p, false)
		if fsErr != nil {
			return nil, &task.ErrTask{
				Err:  fmt.Errorf("check if %q is a subdirectory of %q", p, rootDir),
				Kind: task.ErrTaskValidation,
			}
		}
		if isSubDir {
			opt.includePathsRel[idx] = fmt.Sprintf("/%s", p)
		}
	}

	return &Task{ac: ac, opts: opt}, nil
}
