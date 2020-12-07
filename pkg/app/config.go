// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package app

// Config holds information and metadata of an application.
type Config struct {
	// BaseOutputDir is the base output directory for an application.
	BaseOutputDir string

	// DisplayName is the display name of an application.
	DisplayName string

	// Name is the name of an application.
	Name string

	// PathRel is the relative path to an application root directory.
	PathRel string

	// PkgImportPath is the import path of an application package.
	PkgImportPath string
}
