// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
	"golang.org/x/mod/modfile"

	glFS "github.com/svengreb/golib/pkg/io/fs"
)

const (
	// GoModuleDefaultBuildInfoVersion is the default version for the build info of a Go module when it is not set or no
	// when no version was detected by the [runtime/debug.ReadBuildInfo] function. [As of Go 1.18] this defaults to [this
	// value] instead of an empty string and [the `go version -m` documentation] also describes the meaning of the value.
	//
	// [runtime/debug.ReadBuildInfo]: https://pkg.go.dev/runtime/debug@go1.18#ReadBuildInfo
	// [As of Go 1.18]: https://github.com/golang/go/commit/9cec77ac11b012283e654b423cf85cf9976bedd9#diff-abdadaf0d85a2e6c8e45da716909b2697d830b0c75149b9e35accda9c38622bdR2234
	// [this value]: https://github.com/golang/go/blob/122a22e0e9eba7fe712030d429fc4bcf6f447f5e/src/cmd/go/internal/load/pkg.go#L2288
	// [the `go version -m` documentation]: https://go.dev/ref/mod#modules-overview#go-version-m
	GoModuleDefaultBuildInfoVersion = "(devel)"

	// GoModuleDefaultFileName is the default name for a Go module file.
	GoModuleDefaultFileName = "go.mod"

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

// GoModuleFromFile parses a Go module file ([GoModuleDefaultBuildInfoVersion]).
// This is required because [as of Go 1.18] the [debug.ReadBuildInfo] function does not work for Mage executables
// anymore because the way how module information is stored changed. Therefore the fields of the returned
// [debug.Module] type only has zero values, including the module path. The [debug.Module.Version] field has a
// [default value] ([GoModuleDefaultBuildInfoVersion]) which is not Semver compatible and causes the parsing to fail.
//
// To get the required module information that was previously provided by the [runtime/debug] package the official
// [golang.org/x/mod/modfile] package is used instead.
//
// [debug.ReadBuildInfo]: https://pkg.go.dev/runtime/debug@go1.18#ReadBuildInfo
// [as of Go 1.18]: https://github.com/golang/go/commit/9cec77ac11b012283e654b423cf85cf9976bedd9#diff-abdadaf0d85a2e6c8e45da716909b2697d830b0c75149b9e35accda9c38622bdR2234
// [default value]: https://github.com/golang/go/blob/122a22e0e9eba7fe712030d429fc4bcf6f447f5e/src/cmd/go/internal/load/pkg.go#L2288
func GoModuleFromFile(dirAbs string) (*GoModuleID, error) {
	goModFilePath := filepath.Join(dirAbs, GoModuleDefaultFileName)
	hasModFile, fsErr := glFS.RegularFileExists(goModFilePath)
	if fsErr != nil {
		return nil, fsErr
	}
	if !hasModFile {
		return nil, fmt.Errorf("no %q file in project root directory path %q", GoModuleDefaultFileName, dirAbs)
	}
	goModFileData, osErr := os.ReadFile(goModFilePath)
	if osErr != nil {
		return nil, fmt.Errorf("read Go module file %q: %w", goModFilePath, osErr)
	}
	goModFile, goModFileErr := modfile.Parse(goModFilePath, goModFileData, nil)
	if goModFileErr != nil {
		return nil, fmt.Errorf("parse Go module file %q: %w", goModFilePath, goModFileErr)
	}

	gmfv := goModFile.Module.Mod.Version
	gm := &GoModuleID{Path: goModFile.Module.Mod.Path}
	// Handle as latest Go module version when the module version equals the default value or is not set/detected and set
	// a valid Semver version string as fallback.
	if gmfv == "" || goModFile.Module.Mod.Version == GoModuleDefaultBuildInfoVersion {
		gm.IsLatest = true
		gmfv = DefaultVersion
	}

	version, semVerErr := semver.NewVersion(gmfv)
	if semVerErr != nil {
		return nil, fmt.Errorf("parse version %q from Go module %q: %w", gmfv, gm.Path, semVerErr)
	}
	gm.Version = version
	return gm, nil
}
