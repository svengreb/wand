// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// Package project provides global project information and metadata, such as the project, module and display name and
// release version.
package project

const (
	// DisplayName is the display name of the project.
	DisplayName = "wand"

	// ModuleName is the name of the ZUP CLI Go module.
	ModuleName = "github.com/svengreb/" + Name

	// Name is the name of the project.
	Name = "wand"

	// ReleaseVersion is the current project release version.
	ReleaseVersion = "v0.7.0"
)
