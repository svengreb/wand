// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"

	appCLI "github.com/svengreb/wand/examples/monorepo/pkg/cmd/cli"
	appDaemon "github.com/svengreb/wand/examples/monorepo/pkg/cmd/daemon"
	appPromExp "github.com/svengreb/wand/examples/monorepo/pkg/cmd/promexp"
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
			wandProj.WithModulePath("github.com/svengreb/wand/examples/monorepo"),
		),
		// Use "github.com/svengreb/nib/inkpen" module as line printer for human-facing messages.
		elder.WithNib(inkpen.New()),
	)
	if ewErr != nil {
		fmt.Printf("Failed to initialize elder wand: %v\n", ewErr)
		os.Exit(1)
	}

	// Register any amount of project applications (monorepo layout).
	apps := []struct {
		name, displayName, pathRel string
	}{
		{appCLI.Name, appCLI.DisplayName, "apps/cli"},
		{appDaemon.Name, appDaemon.DisplayName, "apps/daemon"},
		{appPromExp.Name, appPromExp.DisplayName, "apps/promexp"},
	}
	for _, app := range apps {
		if regErr := ew.RegisterApp(app.name, app.displayName, app.pathRel); regErr != nil {
			ew.ExitPrintf(1, nib.ErrorVerbosity, "Failed to register application %q: %v", app.name, regErr)
		}
	}

	elderWand = ew
}

func baseGoBuild(appName string) {
	buildErr := elderWand.GoBuild(
		appName,
		taskGoBuild.WithBinaryArtifactName(appName),
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

func Build(mageCtx context.Context) {
	mg.SerialDeps(
		CLI.Build,
		Daemon.Build,
		PrometheusExporter.Build,
	)
}

type CLI mg.Namespace

func (CLI) Build(mageCtx context.Context) { baseGoBuild(appCLI.Name) }

type Daemon mg.Namespace

func (Daemon) Build(mageCtx context.Context) { baseGoBuild(appDaemon.Name) }

type PrometheusExporter mg.Namespace

func (PrometheusExporter) Build(mageCtx context.Context) { baseGoBuild(appPromExp.Name) }
