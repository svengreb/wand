// Copyright (c) 2019-present Sven Greb <development@svengreb.de>
// This source code is licensed under the MIT license found in the LICENSE file.

// +build mage

// wand - a simple and powerful toolkit for Mage.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/nib/pencil"

	osSupport "github.com/svengreb/wand/internal/support/os"
	projectSupport "github.com/svengreb/wand/internal/support/project"
	"github.com/svengreb/wand/pkg/elder"
	wandProj "github.com/svengreb/wand/pkg/project"
	wandProjVCS "github.com/svengreb/wand/pkg/project/vcs"
	taskFSClean "github.com/svengreb/wand/pkg/task/fs/clean"
	taskGofumpt "github.com/svengreb/wand/pkg/task/gofumpt"
	taskGoimports "github.com/svengreb/wand/pkg/task/goimports"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoTest "github.com/svengreb/wand/pkg/task/golang/test"
	taskGolangCI "github.com/svengreb/wand/pkg/task/golangcilint"
)

const (
	// defaultIntegrationTestTag is the default tag name for integration tests.
	defaultIntegrationTestTag = taskGoTest.DefaultIntegrationTestTag

	// defaultTestOutputDirName is the default output directory name for test artifacts like profiles and reports.
	defaultTestOutputDirName = taskGoTest.DefaultOutputDirName

	// goModulePathToolGofumpt is the import path for the "gofumpt" Go module command.
	goModulePathToolGofumpt = taskGofumpt.DefaultGoModulePath

	// goModulePathToolGoimports is the version for the "goimports" Go module command.
	goModulePathToolGoimports = taskGoimports.DefaultGoModulePath

	// goModulePathToolGolangCI is the version for the "golangci-lint" Go module command.
	goModulePathToolGolangCI = taskGolangCI.DefaultGoModulePath

	// goModuleVersionToolGofumpt is the version for the "gofumpt" Go module command.
	goModuleVersionToolGofumpt = taskGofumpt.DefaultGoModuleVersion

	// goModuleVersionToolGoimports is the version for the "goimports" Go module command.
	goModuleVersionToolGoimports = taskGoimports.DefaultGoModuleVersion

	// goModuleVersionToolGolangCI is the version for the "golangci-lint" Go module command.
	goModuleVersionToolGolangCI = taskGolangCI.DefaultGoModuleVersion
)

const (
	// wandVerbosity is the name of the environment variable for the wand verbosity level.
	wandVerbosity env = iota
)

var (
	// envPrefix is the full uppercase project name used as prefix for all project-specific environment variables.
	envPrefix = strings.ToUpper(projectSupport.Name)

	// ew is the project's Mage wand.Wand.
	ew *elder.Elder

	// optsTaskTest returns composed options for test task.
	optsTaskTest = func(extraOpts ...taskGoTest.Option) []taskGoTest.Option {
		opts := []taskGoTest.Option{
			// Add the "./..." shortcut to recursively include tests of all sub-packages.
			taskGoTest.WithPkgs(fmt.Sprintf("%s/...", ew.GetProjectMetadata().Options().GoModule.Path)),
			taskGoTest.WithOutputDir(testOutputDir(ew.GetProjectMetadata().Options().BaseOutputDir)),
		}
		return append(opts, extraOpts...)
	}

	// testOutputDir returns the path to the subdirectory within the output directory that is used to store test profiles
	// and reports.
	testOutputDir = func(baseOutputDir string) string {
		return filepath.Join(baseOutputDir, defaultTestOutputDirName)
	}
)

// env is a project-specific environment variable.
type env int

// init initializes and populates the project and CLI message printer configurations.
func init() {
	// Initialize and configure the project-wide CLI message printer.
	verbLvl := nib.InfoVerbosity
	wandVerbLvl, isVerbLvlSet := os.LookupEnv(wandVerbosity.String())
	if isVerbLvlSet {
		lvl, parseErr := nib.ParseVerbosity(wandVerbLvl)
		if parseErr != nil {
			color.Yellow("Invalid wand verbosity level %q, using default %q level", wandVerbLvl, verbLvl)
		}
		verbLvl = lvl
	}
	ink := inkpen.New(inkpen.WithPencilOptions(pencil.WithVerbosity(verbLvl)))

	elderWand, ewErr := elder.New(
		elder.WithProjectOptions(
			wandProj.WithName(projectSupport.Name),
			wandProj.WithDisplayName(projectSupport.DisplayName),
			wandProj.WithVCSKind(wandProjVCS.KindGit),
		),
		elder.WithGoRunnerOptions(
			taskGo.WithRunnerEnv(osSupport.EnvSliceToMap(os.Environ())),
		),
		elder.WithNib(ink),
	)

	if ewErr != nil {
		color.Red("Initialization failed: %v\n", ewErr)
		os.Exit(1)
	}
	ew = elderWand
}

func (e env) String() string {
	envVars := []string{
		"VERBOSITY",
	}
	return fmt.Sprintf("%s_%s", envPrefix, envVars[e])
}

// Bootstrap runs initialization tasks and sets up the local development environment by installing required tools and
// build dependencies.
func Bootstrap() {
	importPath := func(path, version string) string { return fmt.Sprintf("%s@%s", path, version) }
	goTools := []string{
		importPath(goModulePathToolGofumpt, goModuleVersionToolGofumpt),
		importPath(goModulePathToolGoimports, goModuleVersionToolGoimports),
		importPath(goModulePathToolGolangCI, goModuleVersionToolGolangCI),
	}

	ew.Infof("Installing development tools:\n")
	for _, path := range goTools {
		printRawf("  ↳ %s\n", path)
	}
	errs := ew.Bootstrap(goTools...)
	if len(errs) > 0 {
		for _, err := range errs {
			ew.Errorf(err.Error())
		}
		ew.ExitPrintf(1, nib.FatalVerbosity, "Boostrap incomplete")
	}

	ew.Successf("Bootstrap completed")
}

// Clean removes artifacts from previous task executions.
func Clean() {
	ew.Infof("Removing previous test artifacts")
	cleaned, err := ew.Clean(
		ew.GetProjectMetadata().Options().Name,
		// TODO: Required?
		taskFSClean.WithLimitToAppOutputDir(true),
		taskFSClean.WithPaths(testOutputDir(ew.GetProjectMetadata().Options().BaseOutputDir)),
	)
	if err != nil {
		ew.ExitPrintf(1, nib.ErrorVerbosity, "%s\n  ↳ %v", color.RedString("Clean failed:"), err)
	}
	ew.Successf("Clean completed")
	for _, cp := range cleaned {
		printRawf("  ↳ %s\n", cp)
	}
}

// Format formats all Go source files with import optimizations and additional rules according to the Go standard code
// style.
func Format() {
	mg.SerialDeps(
		func() {
			ew.Infof("Formatting all Go source files with import optimizations according to the Go standard code style")
			opts := []taskGoimports.Option{
				taskGoimports.WithListNonCompliantFiles(true),
				taskGoimports.WithLocalPkgs(ew.GetProjectMetadata().Options().GoModule.Path),
				taskGoimports.WithPersistedChanges(true),
				taskGoimports.WithReportAllErrors(true),
			}
			if err := ew.Goimports(opts...); err != nil {
				ew.ExitPrintf(1, nib.ErrorVerbosity, "Formatting failed:\n  ↳ %v", err)
			}
		},
		func() {
			ew.Infof("Formatting all Go source files with additional rules according to the Go standard code style")
			opts := []taskGofumpt.Option{
				taskGofumpt.WithExtraRules(true),
				taskGofumpt.WithListNonCompliantFiles(true),
				taskGofumpt.WithPersistedChanges(true),
				taskGofumpt.WithReportAllErrors(true),
				taskGofumpt.WithSimplify(true),
			}
			if err := ew.Gofumpt(opts...); err != nil {
				ew.ExitPrintf(1, nib.ErrorVerbosity, "Formatting failed:\n  ↳ %v", err)
			}
		},
	)
	ew.Successf("Formatting completed")
}

// Lint runs all configured "golangci-lint" linters.
// See the ".golangci.yml" configuration file and official GolangCI documentations at https://golangci-lint.run
// and https://github.com/golangci/golangci-lint for more details.
func Lint() {
	mg.SerialDeps(
		func() {
			ew.Infof(`Running configured "golangci-lint" linters`)
			err := ew.GolangCILint(taskGolangCI.WithVerboseOutput(true))
			if err != nil {
				ew.ExitPrintf(1, nib.ErrorVerbosity, "Linting failed:\n  ↳ %v", err)
			}
		},
	)
	ew.Successf("Linting completed")
}

// Test runs all unit tests.
func Test() {
	Clean()
	ew.Infof("Running %s tests", color.CyanString(ew.GetProjectMetadata().Options().GoModule.Path))
	runTest()
	ew.Successf("Test(s) completed:\n  ↳ %s", color.GreenString(ew.GetProjectMetadata().Options().GoModule.Path))
}

// TestCover runs all unit tests with coverage reports.
func TestCover() {
	Clean()
	ew.Infof("Running %s tests with coverage", color.CyanString(ew.GetProjectMetadata().Options().GoModule.Path))
	runTest(taskGoTest.WithCoverageProfile(true))
	ew.Successf("Test(s) completed:\n  ↳ %s", color.GreenString(ew.GetProjectMetadata().Options().GoModule.Path))
}

// TestIntegration runs all integration tests.
func TestIntegration() {
	Clean()
	ew.Infof("Running %s integration tests", color.CyanString(ew.GetProjectMetadata().Options().GoModule.Path))
	runTest(taskGoTest.WithGoOptions(
		taskGo.WithTags(defaultIntegrationTestTag),
	))
}

// TestRace runs all unit tests with enabled race detection.
// Please note that race detection will fail when using a Go executable that has been build in PIE mode!
// See "go help buildmode" and https://github.com/golang/go/issues/33514 for more details and make sure to run this
// task with a Go executable that was build without "-buildmode=pie" flag.
// Also see https://golang.org/doc/articles/race_detector.html for more details about the race detector.
func TestRace() {
	Clean()
	ew.Infof(
		"Running %s tests with enabled race detection",
		color.CyanString(ew.GetProjectMetadata().Options().GoModule.Path),
	)
	runTest(taskGoTest.WithGoOptions(
		taskGo.WithRaceDetector(true),
	))
}

// UpgradeMods updates outdated Go module dependencies interactively.
func UpgradeMods() {
	ew.Infof("Updating outdated Go dependencies")
	if err := ew.GoModUpgrade(); err != nil {
		ew.ExitPrintf(1, nib.ErrorVerbosity, "%s\n  ↳ %v", color.RedString("Updating outdated dependencies failed:"), err)
	}
}

// printRawf writes a message to the underlying io.Writer of ew without any specific formatting.
// If an error occurs while writing to the underlying io.Writer the message is printed to os.Stdout instead.
// When this also returns an error the error is written to os.Stderr instead.
func printRawf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if _, err := fmt.Fprint(ew.Writer(), msg); err != nil {
		if _, err = fmt.Fprint(os.Stdout, msg); err != nil {
			_, _ = fmt.Fprint(os.Stderr, err)
		}
	}
}

func runTest(opts ...taskGoTest.Option) {
	if err := ew.GoTest(ew.GetProjectMetadata().Options().Name, optsTaskTest(opts...)...); err != nil {
		ew.Warnf(color.YellowString(
			"Please note that race detection will fail when using a Go executable that has been build in PIE mode!",
		))
		ew.ExitPrintf(1, nib.ErrorVerbosity, "%s\n  ↳ %v", color.RedString("Test(s) failed:"), err)
	}
}
