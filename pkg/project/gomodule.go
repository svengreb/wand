// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package project

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
)

const (
	// GoModuleVersionLatest is the "version query suffix" for the latest version of a Go module.
	// See https://golang.org/ref/mod#version-queries for more details.
	GoModuleVersionLatest = "latest"

	// GoModuleVersionSuffixSeparator is the character that separates the Go module version from a import path.
	GoModuleVersionSuffixSeparator = "@"
)

// GoModuleID stores partial information to identify a Go module.
//
// See https://golang.org/ref/mod#modules-overview for more details.
type GoModuleID struct {
	// IsLatest indicates whether the Go module version uses GoModuleVersionLatest as "version query suffix".
	IsLatest bool

	// Path is the canonical name for a module, declared with the module directive in the module's go.mod file.
	//
	// References
	//
	//   (1) https://golang.org/ref/mod#module-path
	//   (2) https://golang.org/ref/mod#go-mod-file-module
	//   (3) https://golang.org/ref/mod#glos-go-mod-file
	Path string

	// Version identifies an immutable snapshot of a module starting with the letter "v", followed by a semantic
	// version.
	// Note that a nil value is resolved using GoModuleVersionLatest s "version query suffix".
	//
	// References
	//
	//   (1) https://golang.org/ref/mod#versions
	//   (2) https://golang.org/ref/mod#version-queries
	//   (3) https://semver.org/spec/v2.0.0.html
	//   (4) https://golang.org/cmd/go/#hdr-Pseudo_versions
	//   (5) https://blog.golang.org/publishing-go-modules
	Version *semver.Version
}

// ExecName returns the name of the compiled executable when the Go module Path is a "main" package.
func (gm GoModuleID) ExecName() string {
	return filepath.Base(gm.Path)
}

func (gm GoModuleID) String() string {
	if gm.Version != nil && !gm.IsLatest {
		return fmt.Sprintf("%s@%s", gm.Path, gm.Version.Original())
	}
	return fmt.Sprintf("%s@%s", gm.Path, GoModuleVersionLatest)
}

// GoModuleFromImportPath creates a GoModuleID from the given import path.
// The path must be a valid Go module import path, that can optionally include the version suffix, in the "pkg@version"
// format.
func GoModuleFromImportPath(importPath string) (*GoModuleID, error) {
	pathElements := strings.Split(importPath, GoModuleVersionSuffixSeparator)
	if len(pathElements) == 0 {
		return nil, fmt.Errorf("invalid import path: %q", importPath)
	}

	gm := &GoModuleID{Path: pathElements[0]}
	// Handle as latest Go module version when the import path has no separator or the suffix equals "latest".
	if len(pathElements) == 1 || pathElements[len(pathElements)-1] == GoModuleVersionLatest {
		gm.IsLatest = true
		return gm, nil
	}

	version, semVerErr := semver.NewVersion(pathElements[1])
	if semVerErr != nil {
		return nil, &ErrProject{
			Err:  fmt.Errorf("parse version from import path %q: %w", importPath, semVerErr),
			Kind: ErrDetermineGoModuleInformation,
		}
	}
	gm.Version = version
	return gm, nil
}
