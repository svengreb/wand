// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package project

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

const (
	// GoModuleVersionLatest is the "version query suffix" for the latest version of a Go module.
	//
	// See https://golang.org/ref/mod#version-queries for more details.
	GoModuleVersionLatest = "latest"
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

func (gm GoModuleID) String() string {
	if gm.Version != nil && !gm.IsLatest {
		return fmt.Sprintf("%s@%s", gm.Path, gm.Version.Original())
	}
	return fmt.Sprintf("%s@%s", gm.Path, GoModuleVersionLatest)
}
