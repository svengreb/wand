// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the license file.

// Package exec provides utilities to run executables.
package exec

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	glFS "github.com/svengreb/golib/pkg/io/fs"

	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

// GetGoExecPath looks up the executable search path(s) of the current environment for the Go executable with the given
// name. It looks up the paths defined in the system "PATH" environment variable and continues with the Go specific
// "GOBIN" path.
// See https://pkg.go.dev/cmd/go#hdr-Environment_variables for more details about Go specific environment variables.
func GetGoExecPath(name string) (string, error) {
	// Look up the system executable search path(s)...
	execPath, pathErr := exec.LookPath(name)
	os.Environ()

	// ...and continue with the Go specific executable search path.
	if pathErr != nil {
		var execDir string

		if execDir = os.Getenv(taskGo.DefaultEnvVarGOBIN); execDir == "" {
			if execDir = os.Getenv(taskGo.DefaultEnvVarGOPATH); execDir != "" {
				execDir = filepath.Join(execDir, taskGo.DefaultGOBINSubDirName)
			}
		}

		execPath = filepath.Join(execDir, name)
		execExits, fsErr := glFS.RegularFileExists(execPath)
		if fsErr != nil {
			return "", fmt.Errorf("check if %q exists: %w", execPath, fsErr)
		}

		if !execExits {
			return "", fmt.Errorf("%q not found in executable search path(s): %v", name, append(os.Environ(), execDir))
		}
	}

	return execPath, nil
}
