// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package clean provides a task to remove filesystem paths, e.g. output data like artifacts and reports from previous
// development, test, production and distribution builds.
package clean

import (
	"fmt"
	"os"
	"path/filepath"

	glFS "github.com/svengreb/golib/pkg/io/fs"
	glFilePath "github.com/svengreb/golib/pkg/io/fs/filepath"

	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/project"
	"github.com/svengreb/wand/pkg/task"
)

// Task is a task to remove filesystem paths, e.g. output data like artifacts and reports from previous development,
// test, production and distribution builds.
type Task struct {
	ac   app.Config
	opts *Options
	proj project.Metadata
}

// Clean removes the configured paths.
// It returns an error of type *param.ErrGoCode for any error that occurs during the execution of the Go code.
func (t *Task) Clean() ([]string, error) {
	var cleaned []string

	for _, p := range t.opts.paths {
		pAbs := filepath.Join(t.proj.Options().RootDirPathAbs, p)

		if t.opts.limitToAppOutputDir {
			appDir := filepath.Join(t.proj.Options().RootDirPathAbs, t.ac.BaseOutputDir)
			pAbs = filepath.Join(t.proj.Options().RootDirPathAbs, p)

			isSubDir, subDirErr := glFilePath.IsSubDir(appDir, pAbs, false)
			if subDirErr != nil {
				return cleaned, &task.ErrTask{
					Err:  fmt.Errorf("check if %q is a subdirectory of %q: %w", pAbs, appDir, subDirErr),
					Kind: task.ErrInvalidTaskOpts,
				}
			}
			if !isSubDir {
				return cleaned, &task.ErrTask{
					Err:  fmt.Errorf("%q is not a subdirectory of %q", pAbs, appDir),
					Kind: task.ErrInvalidTaskOpts,
				}
			}
		}

		nodeExists, fsErr := glFS.FileExists(pAbs)
		if fsErr != nil {
			return cleaned, &task.ErrTask{
				Err:  fmt.Errorf("check if %q exists: %w", pAbs, fsErr),
				Kind: task.ErrInvalidTaskOpts,
			}
		}
		if nodeExists {
			if err := os.RemoveAll(pAbs); err != nil {
				return cleaned, &task.ErrRunner{
					Err:  fmt.Errorf("remove path %q: %w", pAbs, err),
					Kind: task.ErrRun,
				}
			}
			cleaned = append(cleaned, p)
		}
	}

	return cleaned, nil
}

// Options returns the task options.
func (t *Task) Options() task.Options {
	return *t.opts
}

// New creates a new task.
//nolint:gocritic // The app.Config struct is passed as value by design to ensure immutability.
func New(proj project.Metadata, ac app.Config, opts ...Option) (*Task, error) {
	opt, optErr := NewOptions(opts...)
	if optErr != nil {
		return nil, optErr
	}

	return &Task{ac: ac, proj: proj, opts: opt}, nil
}
