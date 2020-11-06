// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package gox provides a task for the github.com/mitchellh/gox Go module command.
// "gox" is a dead simple, no frills Go cross compile tool that behaves a lot like the standard Go toolchain "build"
// command.
//
// See https://pkg.go.dev/github.com/mitchellh/gox for more details about "gox".
// The source code of the "gox" is available at https://github.com/mitchellh/gox.
package gox

import (
	"fmt"
	"strings"

	"github.com/svengreb/wand"
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

// Task is a task for the "github.com/mitchellh/gox" Go module command.
type Task struct {
	ac   app.Config
	opts *Options
}

// BuildParams builds the parameters.
func (t *Task) BuildParams() []string {
	params := taskGo.BuildGoOptions(t.opts.taskGoOpts...)

	// Workaround to allow the usage of the "-trimpath" flag that has been introduced in Go 1.13.0.
	// The currently latest version of "gox" does not support the flag yet.
	//
	// See https://github.com/mitchellh/gox/pull/138 for more details.
	for idx, arg := range params {
		if arg == "-trimpath" {
			params = append(params[:idx], params[idx+1:]...)
			// Set the flag via the GOFLAGS environment variable instead.
			t.opts.env[taskGo.DefaultEnvVarGOFLAGS] = fmt.Sprintf(
				"%s %s -trimpath",
				t.opts.Env[taskGo.DefaultEnvVarGOFLAGS],
				t.opts.env[taskGo.DefaultEnvVarGOFLAGS],
			)
		}
	}

	if t.opts.verbose {
		params = append(params, "-verbose")
	}

	if t.opts.goCmd != "" {
		params = append(params, fmt.Sprintf("-gocmd=%s", t.opts.goCmd))
	}

	if len(t.opts.CrossCompileTargetPlatforms) > 0 {
		params = append(params, fmt.Sprintf("-osarch=%s", strings.Join(t.opts.CrossCompileTargetPlatforms, " ")))
	}

	params = append(params, fmt.Sprintf("--output=%s/%s", t.opts.OutputDir, t.opts.outputTemplate))

	if len(t.opts.Flags) > 0 {
		params = append(params, t.opts.Flags...)
	}

	return append(params, t.ac.PkgImportPath)
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

// New creates a new task for the "github.com/mitchellh/gox" Go module command.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(wand wand.Wand, ac app.Config, opts ...Option) (*Task, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}

	if opt.BinaryArtifactName == "" {
		opt.BinaryArtifactName = ac.Name
	}

	// Store build artifacts in the application specific subdirectory.
	if opt.OutputDir == "" {
		opt.OutputDir = ac.BaseOutputDir
	}

	if opt.outputTemplate == "" {
		opt.outputTemplate = DefaultCrossCompileBinaryNameTemplate(opt.BinaryArtifactName)
	}

	return &Task{ac: ac, opts: opt}, nil
}
