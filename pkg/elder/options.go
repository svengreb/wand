// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package elder

import (
	//nolint:golint // Idiomatic for the Go 1.16 "embed" package.
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	glFS "github.com/svengreb/golib/pkg/io/fs"
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"

	"github.com/svengreb/wand/pkg/project"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoTool "github.com/svengreb/wand/pkg/task/gotool"
)

var (
	// DefaultGoToolsBinDir is the default directory for compiled executables of Go module-based "main" packages.
	DefaultGoToolsBinDir = filepath.Join(project.DefaultWandCacheDataDir, "tools", "bin")

	// wandDataGitIgnoreFileName is the name for the written wandDataGitIgnoreTmpl file.
	wandDataGitIgnoreFileName = ".gitignore"

	//go:embed gitignore.tmpl
	wandDataGitIgnoreTmpl []byte
)

// Option is a wand option.
type Option func(*Options)

// Options are wand options.
type Options struct {
	// disableAutoGenWandDataDir indicates whether the auto-generation of the directory for wand specific data should be
	// disabled.
	disableAutoGenWandDataDir bool

	// goRunnerOpts are Go toolchain runner options.
	goRunnerOpts []taskGo.RunnerOption

	// goToolRunnerOpts are Go module-based tool runner options.
	goToolRunnerOpts []taskGoTool.RunnerOption

	// nib is the log-level based line printer for human-facing messages.
	nib nib.Nib

	// projectOpts are project options.
	projectOpts []project.Option
}

// NewOptions creates new wand options.
func NewOptions(opts ...Option) *Options {
	opt := &Options{nib: inkpen.New()}
	for _, o := range opts {
		o(opt)
	}

	return opt
}

// WithDisableAutoGenWandDataDir indicates whether the auto-generation of the directory for wand specific data should be
// disabled.
func WithDisableAutoGenWandDataDir(disableAutoGenWandDataDir bool) Option {
	return func(o *Options) {
		o.disableAutoGenWandDataDir = disableAutoGenWandDataDir
	}
}

// WithGoRunnerOptions sets Go toolchain runner options.
func WithGoRunnerOptions(opts ...taskGo.RunnerOption) Option {
	return func(o *Options) {
		o.goRunnerOpts = append(o.goRunnerOpts, opts...)
	}
}

// WithGoToolRunnerOptions sets Go module-based tool runner options.
func WithGoToolRunnerOptions(opts ...taskGoTool.RunnerOption) Option {
	return func(o *Options) {
		o.goToolRunnerOpts = append(o.goToolRunnerOpts, opts...)
	}
}

// WithNib sets the log-level based line printer for human-facing messages.
func WithNib(n nib.Nib) Option {
	return func(o *Options) {
		if n != nil {
			o.nib = n
		}
	}
}

// WithProjectOptions sets project options.
func WithProjectOptions(opts ...project.Option) Option {
	return func(o *Options) {
		o.projectOpts = append(o.projectOpts, opts...)
	}
}

// generateWandDataDir generates the wand specific data directory structure and files.
func generateWandDataDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("make %q directory structure: %w", path, err)
	}

	gitIgnoreFilePath := filepath.Join(path, wandDataGitIgnoreFileName)
	gitIgnoreExists, fsErr := glFS.RegularFileExists(gitIgnoreFilePath)
	if fsErr != nil {
		return fmt.Errorf("check regular file %q: %w", gitIgnoreFilePath, fsErr)
	}
	if !gitIgnoreExists {
		if err := os.WriteFile(gitIgnoreFilePath, wandDataGitIgnoreTmpl, os.ModePerm); err != nil {
			return fmt.Errorf("write %q: %w", gitIgnoreFilePath, err)
		}
	}

	return nil
}
