// +build mage

// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"

	"github.com/svengreb/wand/examples/simple/pkg/cmd/cli"
	"github.com/svengreb/wand/pkg/elder"
	wandProj "github.com/svengreb/wand/pkg/project"
	wandProjVCS "github.com/svengreb/wand/pkg/project/vcs"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoBuild "github.com/svengreb/wand/pkg/task/golang/build"
)

const (
	projectDisplayName = "Fruit Mixer"
	projectName        = "fruit-mixer"
)

var elderWand *elder.Elder

func init() {
	// Create a new "elder wand".
	ew, ewErr := elder.New(
		// Provide information about the project.
		elder.WithProjectOptions(
			wandProj.WithName(projectName),
			wandProj.WithDisplayName(projectDisplayName),
			wandProj.WithVCSKind(wandProjVCS.KindNone),
			wandProj.WithModulePath("github.com/svengreb/wand/examples/single_cmd"),
		),
		// Use "github.com/svengreb/nib/inkpen" module as line printer for human-facing messages.
		elder.WithNib(inkpen.New()),
	)
	if ewErr != nil {
		fmt.Printf("Failed to initialize elder wand: %v\n", ewErr)
		os.Exit(1)
	}

	// Register the CLI application.
	if err := ew.RegisterApp(cli.Name, cli.DisplayName, "."); err != nil {
		ew.ExitPrintf(1, nib.ErrorVerbosity, "Failed to register application: %v", err)
	}

	elderWand = ew
}

func Build(mageCtx context.Context) {
	buildErr := elderWand.GoBuild(
		cli.Name,
		taskGoBuild.WithBinaryArtifactName(cli.Name),
		taskGoBuild.WithOutputDir("out"),
		taskGoBuild.WithGoOptions(
			taskGo.WithTrimmedPath(true),
		),
	)
	if buildErr != nil {
		fmt.Printf("Build incomplete: %v\n", buildErr)
	}
	elderWand.Successf("Build completed")
}
