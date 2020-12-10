<p align="center"><img src="https://raw.githubusercontent.com/svengreb/wand/main/assets/images/repository-hero.svg?sanitize=true"/></p>

<p align="center"><a href="https://github.com/svengreb/wand/releases/latest"><img src="https://img.shields.io/github/release/svengreb/wand.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/svengreb/wand/blob/main/CHANGELOG.md"><img src="https://img.shields.io/github/release/svengreb/wand.svg?style=flat-square&label=Changelog&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://pkg.go.dev/github.com/svengreb/wand"><img src="https://img.shields.io/github/release/svengreb/wand.svg?style=flat-square&label=GoDoc&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center"><a href="https://github.com/svengreb/wand/actions?query=workflow%3Aci" target="_blank"><img src="https://img.shields.io/github/workflow/status/svengreb/wand/ci.svg?style=flat-square&label=CI&logo=github&logoColor=eceff4&colorA=4c566a"/></a> <a href="https://codecov.io/gh/svengreb/wand" target="_blank"><img src="https://img.shields.io/codecov/c/github/svengreb/wand/main.svg?style=flat-square&label=Coverage&logo=codecov&logoColor=eceff4&colorA=4c566a"/></a></p>

<p align="center"><a href="https://golang.org/doc/effective_go.html#formatting" target="_blank"><img src="https://img.shields.io/static/v1?style=flat-square&label=Go%20Style%20Guide&message=gofmt&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/arcticicestudio/styleguide-markdown/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-markdown.svg?style=flat-square&label=Markdown%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIzOSIgaGVpZ2h0PSIzOSIgdmlld0JveD0iMCAwIDM5IDM5Ij48cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNEOERFRTkiIHN0cm9rZS13aWR0aD0iMyIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBkPSJNMS41IDEuNWgzNnYzNmgtMzZ6Ii8%2BPHBhdGggZmlsbD0iI0Q4REVFOSIgZD0iTTIwLjY4MyAyNS42NTVsNS44NzItMTMuNDhoLjU2Nmw1Ljg3MyAxMy40OGgtMS45OTZsLTQuMTU5LTEwLjA1Ni00LjE2MSAxMC4wNTZoLTEuOTk1em0tMi42OTYgMGwtMTMuNDgtNS44NzJ2LS41NjZsMTMuNDgtNS44NzJ2MS45OTVMNy45MzEgMTkuNWwxMC4wNTYgNC4xNnoiLz48L3N2Zz4%3D"/></a> <a href="https://github.com/arcticicestudio/styleguide-git/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-git.svg?style=flat-square&label=Git%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=git"/></a></p>

<blockquote cite="https://en.wikipedia.org/wiki/Harry_Potter_and_the_Philosopher%27s_Stone_(film)">
  <p>“The wand chooses the mage, remember.“</p>
  <footer>— Garrick Ollivander, <cite><em>Harry Potter and the Sorcerer’s Stone</em></cite></footer>
</blockquote>

<p align="center">A simple and powerful toolkit for <a href="https://magefile.org" target="_blank">Mage</a>.</p>

## Features

_wand_ is a toolkit for common and often recurring project processes for the task automation tool [Mage][].
The provided [API packages][go-pkg-pkg] allow users to compose their own, reusable set of tasks and helpers or built up on the [reference implementation][go-pkg-elder].

- **Adapts to any “normal“ or [“mono“][trunkbasedev-monorepos] repository layout** — handle as many module _commands_ as you want. _wand_ uses an abstraction by managing every `main` package as _application_ so that tasks can be processed for all or just individual _commands_.
- **Runs any `main` package of a [Go module][go-docs-ref-mod] without the requirement for the user to install it beforehand** — thanks to the awesome [gobin][] project, there is no need for the user to `go get` the `main` package of a Go module in order to run its compiled executable.
- **Comes with support for basic [Go toolchain][go-pkg-cmd/go] commands and popular modules from the Go ecosystem** — run common commands like `go build` and `go test` or great tools like [goimports][go-pkg-golang.org/x/tools/cmd/goimports], [golangci-lint][go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint] and [gox][go-pkg-github.com/mitchellh/gox] in no time.

See the [API](#api) and [“Elder Wand“](#elder-wand) sections for more details. The [user guides](#user-guides) for more information about how to build your own tasks and runners and the [examples](#examples) for different repositories layouts (single or [“monorepo“][trunkbasedev-monorepos]) and use cases.

## Motivation

<!--lint disable no-heading-punctuation-->

### Why Mage?

Every project involves processes that are often recurring. These can mostly be done with the tools supplied with the respective programming language, which in turn, in many cases, involve more time and the memorizing of longer commands with their flags and parameters.
In order to significantly reduce this effort or to avoid it completely, project task automation tools are used which often establish a defined standard to enable the widest possible use and unify tasks. They offer a user-friendly and comfortable interface to handle the processes consistently with time savings and without the need for developers to remember many and/or complex commands.
But these tools come with a cost: the introduction of standards and the restriction to predefined ways how to handle tasks is also usually the biggest disadvantage when it comes to adaptability for use cases that are individual for a single project, tasks that deviate from the standard or not covered by it at all.

[Mage][] is a project task automation tool which gives the user complete freedom by **not specifying how tasks are solved, but only how they are started and connected with each other**. This is an absolute advantage over tools that force how how a task has to be solved while leaving out individual and project specific preferences.

If you would now ask me “But why not just use [Make][]?“, my answer would be “Why use a tool that is not native to the programming language it is intended for?“.
_Make_ has somehow become a popular choice as task automation tool for Go projects and up to today I don‘t get it. Don‘t get me wrong: this is no bad talking against _Make_ but a clarification that it is not intended for Go but rather for _C_ projects, e.g. [the Linux kernel][linux], since [_Make_ is also written in _C_][gnu-make-repo]. Even [Go itself is built using shell and Windows DOS scripts][gh-tree-golang/go/src] instead of _Make_.
If you take a closer look, _Make_ is nothing more than a [DSL][wikip-dsl] for [shell commands][gnu-make-docs-shell] so using shell/Windows DOS scripts directly instead is a way more flexible option. Therefore _Make_ can not fullfil an important criteria: full cross-platform compatibility. The command(s) that each task runs must be available on the system, e.g. other tools must be installed globally and available in the [executable search path][wikip-path_var], as well as requiring the syntax to be compatible with the underlying shell which makes it hard to use [shell builtin][wikip-shell_builtin] commands like `cd`.

In my opinion, **a task automation tool for a project should always be written in the same programming language that it is intended for**. This concept has already been proven for many other languages, e.g. official tools like [cargo][rust-docs-cargo] for _Rust_ and [NPM][npm-com] for _Node.js_‘s or community projects like [Gradle][] or [Maven][] for _Java_. All of them can natively understand and interact with their target programming language to provide the widest range of features, good stability and often also ways to simply extend their functionality through plugin systems.

This is where _Mage_ comes in:

- Written in [pure Go without any external dependencies][gh-blob-magefile/mage/go.mod] for fully native compatibility and easy expansion.
- [No installation][mage-zero_install] required.
- Allows to [declare dependencies between targets][mage-deps] in a _makefile_-style tree and optionally [runs them in parallel][mage-deps#paral].
- Targets can be defined in shared packages and [imported][mage-importing] in any [_Magefile_][mage-files]. No mechanics like plugins or extensions required, just use any Go module and the whole Go ecosystem.

### Why _wand_?

While _Mage_ is often already sufficient on its own, I‘ve noticed that I had to implement almost identical tasks over and over again when starting a new project or migrating an existing one to _Mage_. Even though the actual [target functions][mage-targets] could be moved into their own Go module to allow to simply [import them][mage-importing] in different projects, it was often required to copy & paste code across projects that can not be exported that easily. That was the moment where I decided to create a way that simplifies the integration and usage of _Mage_ without loosing any of its flexibility and dynamic structure.

Please note that this package has mainly been created for my personal use in mind to avoid copying source code between my projects. The default configurations or reference implementation might not fit your needs, but the [API](#api) packages have been designed so that they can be flexibly adapted to different use cases and environments or used to create and compose your own [`wand.Wand`][go-pkg-if-wand#wand].

See the [API](#api) and [“Elder Wand“](#elder-wand) sections to learn how to adapt or extend _wand_ for your project.

<!--lint enable no-heading-punctuation-->

## Wording

Since _wand_ is a toolkit for [Mage][], is partly makes use of an abstract naming scheme that matches the fantasy of magic which in case of _wand_ has been derived from the fantasy novel [“Harry Potter“][wikip-hp]. This is mainly limited to the [main “Wand“ interface][go-pkg-if-wand#wand] and the [“Elder Wand“](#elder-wand) reference implementation.
The basic mindset of the API is designed around the concept of **tasks** and the ways to **run** them.

- **Runner** — Components that run a command with parameters in a specific environment, in most cases a (binary) [executable][wikip-exec] of external commands or Go module `main` packages.
- **Tasks** — Components that are scoped for Mage [“target“][mage-targets] usage in order to run an action.

## API

The public _wand_ API is located in the [`pkg`][go-pkg-pkg] package while the main interface [`wand.Wand`][go-pkg-if-wand#wand], that manages a project and its applications and stores their metadata, is defined in the [`wand`][go-pkg-wand] package.

Please see the individual documentations of each package for more details.

### Application Configurations

The [`app`][go-pkg-app] package provides the functionality for application configurations. A [`Config`][go-pkg-stc-app#config] holds information and metadata of an application that is stored by types that implement the [`Store` interface][go-pkg-if-app#store]. The [`NewStore() app.Store`][go-pkg-func-app#newstore] function returns a reference implementation of this interface.

### Command Runners

The [`task`][go-pkg-task] package defines the API for runner of commands. [`Runner`][go-pkg-if-task#runner] is the base interface while [`RunnerExec` interface][go-pkg-if-task#runnerexec] is a specialized for (binary) executables of a command.

The package already provides runners for the [Go toolchain][go-pkg-cmd/go] and the [gobin][] Go module:

- **Go Toolchain** — to interact with the [Go toolchain][go-pkg-cmd/go], also known as the `go` executable, the [`golang.Runner`][go-pkg-stc-task/golang#runner] can be used.
- **`gobin` Go Module** — to install and run [Go module][go-docs-ref-mod] `main` the [`Runner`][go-pkg-stc-task/gobin#runner] makes use of the [`github.com/myitcv/gobin`][go-pkg-github.com/myitcv/gobin] command.
  1. **Go Executable Installation** — Using the [`go install`][go-pkg-cmd/go#install] or [`go get`][go-pkg-cmd/go#print_env] command for a [Go module][go-ref-mod] `main` package, the resulting executables are placed in the Go executable search path that is defined by the [`GOBIN` environment variable][go-pkg-cmd/go#env_vars] (see the [`go env` command][go-pkg-cmd/go#print_env] to show or modify the Go toolchain environment).
     Even though executables are installed “globally“ for the current user, any [`go.mod` file][go-ref-mod#go.mod] in the current working directory will be updated to include the Go module. This is the default behavior of the [`go get` command][go-pkg-cmd/go#print_env] when running in [“module mode“][go-pkg-cmd/go#mod_cmds] (see [`GO111MODULE` environment variable).
     Next to this problem, installed executables will also overwrite any previously installed executable of the same module/package regardless of the version. Therefore only one version of a executable can be installed at a time which makes it impossible to work on different projects that make use of the same executable but require different versions.
  2. **History and Future** — The installation concept for `main` package executables has always been a somewhat controversial point which unfortunately, partly for historical reasons, does not offer an optimal and user-friendly solution up to now.
     The [`go` command][go-pkg-cmd/go] is a fantastic toolchain that provides many great features one would expect to be provided out-of-the-box from a modern and well designed programming language without the requirement to use a third-party solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging utilities and many more.
     Unfortunately this does not apply for the [`go install` command][go-pkg-cmd/go#install] of Go versions less or equal to 1.15.
     The general problem of tool dependencies is a long-time known issue/weak point of the current Go toolchain and is a highly rated change request from the Go community with discussions like [golang/go#30515][gh-golang/go#30515], [golang/go#25922][gh-golang/go#25922] and [golang/go#27653][gh-golang/go#27653] to improve this essential feature, but they've been around for quite a long time without a solution that works without introducing breaking changes and most users and the Go team agree on.
     Luckily, this topic was finally picked up for the next [upcoming Go release version 1.16][gh-golang/go-ms-145] and [golang/go#40276][gh-golang/go#40276] introduces a way to install executables in module mode outside a module.
     The [release note preview also already includes details about this change][go-docs-tip-rln-1.16#mods] and how installation of executables from Go modules will be handled in the future.
  3. **The Workaround** — Beside the great news and anticipation about an official solution for the problem the usage of a workaround is almost inevitable until Go 1.16 is finally released.
     The [official Go wiki][gh-golang/go-wiki] provides a section on [“How can I track tool dependencies for a module?“][gh-golang/go-wiki-mods#tool_deps] that describes a workaround that tracks tool dependencies. It allows to use the Go module logic by using a file like `tools.go` with a dedicated `tools` build tag that prevents the included module dependencies to be picked up for “normal“ executable builds. This approach works fine for non-main packages, but CLI tools that are only implemented in the `main` package can not be imported in such a file.
     In order to tackle this problem, a well-known user from the community created [gobin][], an experimental, module-aware command to install and run `main` packages.
     It allows to install or run `main` package commands without “polluting“ the `go.mod` file. Modules are downloaded in version-aware mode into a cache path within the [users local cache directory][go-pkg-func-os#usercachedir]. This way it prevents problems due to already installed executables by placing each version of an executable in its own directory.
     The decision to use a cache directory instead of sub-directories within the `GOBIN` path doesn't require to mess with the users setup and keep the Go toolchain specific paths clean and unchanged.
     `gobin` is still in an early development state, but has already received a lot of positive feedback and is used in many projects. There are also members of the core Go team that have contributed to the project and the chance is high that the changes for Go 1.16 were influenced or partially ported from it. See [gobin‘s FAQ page][gobin-wiki-faq] in the repository wiki for more details about the project.
     It is currently the best workaround to…
     1. prevent the Go toolchain to pick up the [`GOMOD` (`go env GOMOD`) environment variable][go-pkg-cmd/go#print_env] that is initialized automatically with the path to the [`go.mod` file][go-ref-mod#go.mod] in the current working directory.
     2. install `main` package executables locally for the current user without “polluting“ the `go.mod` file.
     3. install `main` package executables locally for the current user without overriding already installed executables of different versions.
  4. **The Go Module `Runner`** — To allow to manage the tool dependency problem, this package provides a command runner that uses `gobin` in order to prevent the problems described in the sections above like the “pollution“ of the "go.mod" file and allows to…
     1. install `gobin` itself into `GOBIN` ([`go env GOBIN`][go-pkg-cmd/go#print_env]).
     2. run any Go module command by installing `main` package executables locally for the current user into the dedicated `gobin` cache.

### Project Metadata

The [`project`][go-pkg-project] package defines the API for metadata and [VCS][wikip-vcs] information of a project. The [`New(opts ...project.Option) (*project.Metadata, error)`][go-pkg-fn-project#new] function can be used to create a new [project metadata][go-pkg-stc-project#metadata].

The package also already provides a [VCS `Repository` interface reference implementation][go-pkg-if-project/vcs#repository] for [Git][]:

- **VCS “Git“** — the [`git`][go-pkg-project/vcs/git] package provides VCS utility functions to interact with [Git][] repositories.

### Tasks

The [`task`][go-pkg-task] package defines the API for tasks. [`Task`][go-pkg-if-task#task] is the base interface while [`Exec`][go-pkg-if-task#exec] and [`GoModule`][go-pkg-if-task#gomodule] are a specialized to represent the (binary) executable of either an “external“ or Go module based command.

The package also already provides tasks for basic [Go toolchain][go-pkg-cmd/go] commands and popular modules from the Go ecosystem:

- **`goimports`** — the [`goimports`][go-pkg-task/goimports] package provides a task for the [`golang.org/x/tools/cmd/goimports`][go-pkg-golang.org/x/tools/cmd/goimports] Go module command. `goimports` allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in the same style as [`gofmt`][go-pkg-cmd/gofmt] so it can be used as a replacement. The source code of `goimports` is [available in the GitHub repository][gh-golang/tools-tree-cmd/goimports].
- **Go** — The [`golang`][go-pkg-task/golang] package provides tasks for [Go toolchain][go-pkg-cmd/go] commands.
  - **`build`** — to run the [`build` command of the Go toolchain][go-pkg-cmd/go#build] the task of the [`build`][go-pkg-task/golang/build] package can be used.
  - **`test`** — to run the [`test` command of the Go toolchain][go-pkg-cmd/go#test] the task of the [`test`][go-pkg-task/golang/test] package can be used.
- **`golangci-lint`** — the [`golangcilint`][go-pkg-task/golangcilint] package provides a task for the [`github.com/golangci/golangci-lint/cmd/golangci-lint`][go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint] Go module command. `golangci-lint` is a fast, parallel runner for dozens of Go linters that uses caching, supports YAML configurations and has integrations with all major IDEs. The source code of `golangci-lint` is [available in the GitHub repository][gh-golangci/golangci-lint].
- **`gox`** — the [`gox`][go-pkg-task/gox] package provides a task for the [`github.com/mitchellh/gox`][go-pkg-github.com/mitchellh/gox] Go module command. `gox` is a dead simple, no frills Go cross compile tool that behaves a lot like the standard [Go toolchain `build` command][go-pkg-cmd/go#build]. The source code of `gox` is [available in the GitHub repository][gh-mitchellh/gox].
- **`pkger`** — the [`pkger`][go-pkg-task/pkger] package provides a task for the [`github.com/markbates/pkger`][go-pkg-github.com/markbates/pkger] Go module command. `pkger` is a tool for embedding static files into Go binaries.

There are also tasks that don‘t need to implement the task API but make use of some “loose“ features like information about a project application are shared as well as the dynamic option system. They can be used without a `task.Runner`, just like a “normal“ package, and provide Go functions/methods that can be called directly:

- **Filesystem Cleaning** — The [`clean`][go-pkg-task/fs/clean] package provides a task to remove directories in a filesystem.

## Usage Guides

In the following sections you can learn how to use the _wand_ reference implementation [“elder wand“](#elder-wand), compose/extend it or simply implement your own tasks and runners.

### Elder Wand

The [`elder`][go-pkg-elder] package is the reference implementation of the main [`wand.Wand`][go-pkg-if-wand#wand] interface that provides common Mage tasks and stores configurations and metadata for applications of a project. Next to task methods for the Go toolchain and Go module commands, it comes with additional methods like `Bootstrap` to run initialization actions or `Validate` to ensure that the _wand_ is initialized properly.

Create your [_Magefile_][mage-files], e.g `magefile.go`, and use the [`New`][go-pkg-fn-elder#new] function to initialize a new wand and register any amount of applications.
Create a global variable of type `*elder.Elder` and assign the created “elder wand“ to make it available to all functions in your _Magefile_. Even though global variables are a bad practice and should be avoid at all, it‘s totally fine for your task automation since it is non-production code.

Note that the _Mage_ specific **`// +build mage` [build constraint][go-pkg-pkg/go/build#constraints], also known as a build tag, is important** in order to mark the file as _Magefile_. See the [official _Mage_ documentation][mage-files] for more details.

<!--lint disable no-tabs-->

```go
// +build mage

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/svengreb/nib"
	"github.com/svengreb/nib/inkpen"

	"github.com/svengreb/wand/pkg/elder"
	wandProj "github.com/svengreb/wand/pkg/project"
	wandProjVCS "github.com/svengreb/wand/pkg/project/vcs"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
	taskGoBuild "github.com/svengreb/wand/pkg/task/golang/build"
)

var elderWand *elder.Elder

func init() {
	// Create a new "elder wand".
	ew, ewErr := elder.New(
		// Provide information about the project.
		elder.WithProjectOptions(
			wandProj.WithName("fruit-mixer"),
			wandProj.WithDisplayName("Fruit Mixer"),
			wandProj.WithVCSKind(wandProjVCS.KindGit),
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
		{"fruitctl", "Fruit CLI", "apps/cli"},
		{"fruitd", "Fruit Daemon", "apps/daemon"},
		{"fruitpromexp", "Fruit Prometheus Exporter", "apps/promexp"},
	}
	for _, app := range apps {
		if regErr := ew.RegisterApp(app.name, app.displayName, app.pathRel); regErr != nil {
			ew.ExitPrintf(1, nib.ErrorVerbosity, "Failed to register application %q: %v", app.name, regErr)
		}
	}

	elderWand = ew
}
```

Now you can create [_Mage_ target functions][mage-targets] using the task methods of the “elder wand“.

```go
func Build(mageCtx context.Context) {
	buildErr := elderWand.GoBuild(
		cliAppName,
		taskGoBuild.WithBinaryArtifactName(cliAppName),
		taskGoBuild.WithGoOptions(
			taskGo.WithTrimmedPath(true),
		),
	)
	if buildErr != nil {
		fmt.Printf("Build incomplete: %v\n", buildErr)
	}
}
```

<!--lint enable no-tabs-->

See the [examples](#examples) to learn about more uses cases and way how to structure your _Mage_ setup.

### Build It Yourself

_wand_ comes with tasks and runners for common [Go toolchain][go-pkg-cmd/go] commands, the [gobin][] and other popular modules from the Go ecosystem, but the chance is high that you want to build your own for your specific use cases.

#### Custom Tasks

To create your own task that is compatible with the _wand_ API, implement the [`Task`][go-pkg-if-task#task] base interface or any of its specialized interfaces. The `Kind() task.Kind` method must return [`KindBase`][go-pkg-al-task#kindbase] while `Options() task.Options` can return anything since [`task.Options`][go-pkg-al-task#options] is just an alias for `interface{}`.

1. If your task is **intended for an executable command** you need to implement the [`Exec`][go-pkg-if-task#exec] interface where…
   - the `Kind() task.Kind` method must return [`KindExec`][go-pkg-al-task#kindexec].
   - the `BuildParams() []string` method must return all the parameters that should be passed to the executable.
2. If your task is **intended for the `main` package of a Go module**, so basically also an executable command, you need to implement the [`GoModule`][go-pkg-if-task#gomodule] interface where…
   - the `Kind() task.Kind` method must return [`KindGoModule`][go-pkg-al-task#kindgomodule].
   - the `BuildParams() []string` method must return all the parameters that should be passed to the executable that was compiled from the `main` package of the Go module.
   - the returned type of the `ID() *project.GoModuleID` method must provide the import path and module version of the target `main` package.

For sample code of a custom task please see the [examples](#examples) section.
Based on your task kind, you can also either use one of the [already provided command runners](#command-runners), like for the [Go toolchain][go-pkg-task/golang] and [gobin][], or [implement your own custom runner](#custom-runners).

#### Custom Runners

To create your own command runner that is compatible with the _wand_ API, implement the [`Runner`][go-pkg-if-task#runner] base interface or any of its specialized interfaces. The `Handles() Kind` method must return the [`Kind`][go-pkg-al-task#kind] that can be handled while the actual business logic of `Validate() errors` is not bound to any constraint, but like the method names states, should ensure that the runner is configured properly and is operational. The `Run(task.Task) error` method represents the main functionality of the interface and is responsible for running the given [`task.Task`][go-pkg-if-task#task] by passing all task parameters, obtained through the `BuildParams() []string` method, and finally execute the configured file. Optionally you can also inspect and use its [`task.Options`][go-pkg-al-task#options] by casting the type returned from the `Options() task.Options` method.

1. If your runner is **intended for an executable command** you need to implement the [`RunnerExec`][go-pkg-if-task#runnerexec] interface where…
   - the `Handles() Kind` method can return kinds like [`KindExec`][go-pkg-al-task#kindexec] or [`KindGoModule`][go-pkg-al-task#kindgomodule].
   - the `Run(task.Task) error` method runs the given [`task.Task`][go-pkg-if-task#task] by passing all task parameters, obtained through the `BuildParams() []string` method, and finally execute the configured file.
   - it is recommended that the `Validate() error` method tests if the executable file of the command exists at the configured path in the target filesystem or maybe also check other (default) paths if this is not the case. It is also often a good preventative measure to prevent problems to check that the current process actually has permissions to read and execute the file.

For a sample code of a custom command runner please see the [examples](#examples) section.
Based on the kind your command runner can handle, you can also either use one of the [already provided tasks](#tasks) or [implement your own custom task](#custom-task).

## Examples

To learn how to use the _wand_ API and its packages, the [`examples` repository directory][gh-tree-examples] contains code samples for multiple use cases:

- **Create your own command runner** — The [`custom_runner`][gh-tree-examples/custom_runner] directory contains code samples to demonstrate [how to create a custom command runner](#custom-runners). The `FruitMixerRunner` struct implements the [`RunnerExec`][go-pkg-if-task#runnerexec] interface for the imaginary `fruitctl` application.
- **Create your own task** — The [`custom_task`][gh-tree-examples/custom_task] directory contains code samples to demonstrate [how to create a custom task](#custom-tasks). The `MixTask` struct implements the [`Exec`][go-pkg-if-task#exec] interface for the imaginary `fruitctl` application.
- **Usage in a [monorepo][trunkbasedev-monorepos] layout** — The [`monorepo`][gh-tree-examples/monorepo] directory contains code samples to demonstrate the usage in a _monorepo_ layout for three example applications `cli`, `daemon` and `promexp`. The _Magefile_ provides a `build` target to build all applications. Each application also has a dedicated `:build` target using the [`mg.Namespace`][go-pkg-al-github.com/magefile/mage/mg#namespace] to only build it individually.
- **Usage with a simple, single command repository layout** — The [`simple`][gh-tree-examples/simple] directory contains code samples to demonstrate the usage in a “simple“ repository that only provides a single command. The _Magefile_ provides a `build` target to build the `fruitctl` application.

## Contributing

_wand_ is an open source project and contributions are always welcome!

There are many ways to contribute, from [writing- and improving documentation and tutorials][contrib-guide-docs], [reporting bugs][contrib-guide-bugs], [submitting enhancement suggestions][contrib-guide-enhance] that can be added to _wand_ by [submitting pull requests][contrib-guide-pr].

Please take a moment to read the [contributing guide][contrib-guide] to learn about the development process, the [styleguides][contrib-guide-styles] to which this project adheres as well as the [branch organization][contrib-guide-branching] and [versioning][contrib-guide-versioning] model.

The guide also includes information about [minimal, complete, and verifiable examples][contrib-guide-mcve] and other ways to contribute to the project like [improving existing issues][contrib-guide-impr-issues] and [giving feedback on issues and pull requests][contrib-guide-feedback].

<p align="center">Copyright &copy; 2019-present <a href="https://www.svengreb.de" target="_blank">Sven Greb</a></p>

<p align="center"><a href="https://github.com/svengreb/wand/blob/main/LICENSE"><img src="https://img.shields.io/static/v1.svg?style=flat-square&label=License&message=MIT&logoColor=eceff4&logo=github&colorA=4c566a&colorB=88c0d0"/></a></p>

[contrib-guide-branching]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#branch-organization
[contrib-guide-bugs]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#bug-reports
[contrib-guide-docs]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#documentations
[contrib-guide-enhance]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#enhancement-suggestions
[contrib-guide-feedback]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#give-feedback-on-issues-and-pull-requests
[contrib-guide-impr-issues]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#improve-issues
[contrib-guide-mcve]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#mcve
[contrib-guide-pr]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#pull-requests
[contrib-guide-styles]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#styleguides
[contrib-guide-versioning]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md#versioning
[contrib-guide]: https://github.com/svengreb/wand/blob/main/CONTRIBUTING.md
[gh-blob-magefile/mage/go.mod]: https://github.com/magefile/mage/blob/d30a2cfe/go.mod
[gh-golang/go-ms-145]: https://github.com/golang/go/milestone/145
[gh-golang/go-wiki-mods#tool_deps]: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
[gh-golang/go-wiki]: https://github.com/golang/go/wiki
[gh-golang/go#25922]: https://github.com/golang/go/issues/25922
[gh-golang/go#27653]: https://github.com/golang/go/issues/27653
[gh-golang/go#30515]: https://github.com/golang/go/issues/30515
[gh-golang/go#40276]: https://github.com/golang/go/issues/40276
[gh-golang/tools-tree-cmd/goimports]: https://github.com/golang/tools/tree/master/cmd/goimports
[gh-golangci/golangci-lint]: https://github.com/golangci/golangci-lint/tree/master/cmd/golangci-lint
[gh-mitchellh/gox]: https://github.com/mitchellh/gox
[gh-tree-examples]: https://github.com/svengreb/wand/tree/main/examples
[gh-tree-examples/custom_runner]: https://github.com/svengreb/wand/tree/main/examples/custom_runner
[gh-tree-examples/custom_task]: https://github.com/svengreb/wand/tree/main/examples/custom_task
[gh-tree-examples/monorepo]: https://github.com/svengreb/wand/tree/main/examples/monorepo
[gh-tree-examples/simple]: https://github.com/svengreb/wand/tree/main/examples/simple
[gh-tree-golang/go/src]: https://github.com/golang/go/tree/926994fd/src
[git]: https://git-scm.com
[gnu-make-docs-shell]: https://www.gnu.org/software/make/manual/html_node/Choosing-the-Shell.html
[gnu-make-repo]: https://savannah.gnu.org/git/?group=make
[go-docs-ref-mod]: https://golang.org/ref/mod
[go-docs-tip-rln-1.16#mods]: https://tip.golang.org/doc/go1.16#modules
[go-pkg-al-github.com/magefile/mage/mg#namespace]: https://pkg.go.dev/github.com/magefile/mage/mg#Namespace
[go-pkg-al-task#kind]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Kind
[go-pkg-al-task#kindbase]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#KindBase
[go-pkg-al-task#kindexec]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#KindExec
[go-pkg-al-task#kindgomodule]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#KindGoModule
[go-pkg-al-task#options]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Options
[go-pkg-app]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app
[go-pkg-cmd/go]: https://pkg.go.dev/cmd/go
[go-pkg-cmd/go#build]: https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies
[go-pkg-cmd/go#env_vars]: https://pkg.go.dev/cmd/go/#hdr-Environment_variables
[go-pkg-cmd/go#install]: https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
[go-pkg-cmd/go#mod_cmds]: https://golang.org/ref/mod#mod-commands
[go-pkg-cmd/go#print_env]: https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
[go-pkg-cmd/go#test]: https://pkg.go.dev/cmd/go#hdr-Test_packages
[go-pkg-cmd/gofmt]: https://pkg.go.dev/cmd/gofmt
[go-pkg-elder]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder
[go-pkg-fn-elder#new]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder#New
[go-pkg-fn-project#new]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#New
[go-pkg-func-app#newstore]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#NewStore
[go-pkg-func-os#usercachedir]: https://pkg.go.dev/os/#UserCacheDir
[go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint]: https://pkg.go.dev/github.com/golangci/golangci-lint/cmd/golangci-lint
[go-pkg-github.com/markbates/pkger]: https://pkg.go.dev/github.com/markbates/pkger
[go-pkg-github.com/mitchellh/gox]: https://pkg.go.dev/github.com/mitchellh/gox
[go-pkg-github.com/myitcv/gobin]: https://pkg.go.dev/github.com/myitcv/gobin
[go-pkg-golang.org/x/tools/cmd/goimports]: https://pkg.go.dev/golang.org/x/tools/cmd/goimports
[go-pkg-if-app#store]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#Store
[go-pkg-if-project/vcs#repository]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project/vcs#Repository
[go-pkg-if-task#exec]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Exec
[go-pkg-if-task#gomodule]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#GoModule
[go-pkg-if-task#runner]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Runner
[go-pkg-if-task#runnerexec]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#RunnerExec
[go-pkg-if-task#task]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Task
[go-pkg-if-wand#wand]: https://pkg.go.dev/github.com/svengreb/wand#Wand
[go-pkg-pkg]: https://pkg.go.dev/github.com/svengreb/wand/pkg
[go-pkg-pkg/go/build#constraints]: https://pkg.go.dev/pkg/go/build/#hdr-Build_Constraints
[go-pkg-project]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project
[go-pkg-project/vcs/git]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project/vcs/git
[go-pkg-stc-app#config]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#Config
[go-pkg-stc-project#metadata]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#Metadata
[go-pkg-stc-task/gobin#runner]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gobin#Runner
[go-pkg-stc-task/golang#runner]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang#Runner
[go-pkg-task]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task
[go-pkg-task/fs/clean]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/fs/clean
[go-pkg-task/goimports]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/goimports
[go-pkg-task/golang]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang
[go-pkg-task/golang/build]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/build
[go-pkg-task/golang/test]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/test
[go-pkg-task/golangcilint]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golangcilint
[go-pkg-task/gox]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gox
[go-pkg-task/pkger]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/pkger
[go-pkg-wand]: https://pkg.go.dev/github.com/svengreb/wand
[go-ref-mod]: https://golang.org/ref/mod
[go-ref-mod#go.mod]: https://golang.org/ref/mod#go-mod-file
[gobin-wiki-faq]: https://github.com/myitcv/gobin/wiki/FAQ
[gobin]: https://github.com/myitcv/gobin
[gradle]: https://gradle.org
[linux]: https://www.kernel.org
[mage-deps]: https://magefile.org/dependencies
[mage-deps#paral]: https://magefile.org/dependencies/#parallelism
[mage-files]: https://magefile.org/magefiles
[mage-importing]: https://magefile.org/importing
[mage-targets]: https://magefile.org/targets
[mage-zero_install]: https://magefile.org/zeroinstall
[mage]: https://magefile.org
[make]: https://www.gnu.org/software/make
[maven]: https://maven.apache.org
[npm-com]: https://npm.community
[rust-docs-cargo]: https://doc.rust-lang.org/stable/cargo
[trunkbasedev-monorepos]: https://trunkbaseddevelopment.com/monorepos
[wikip-dsl]: https://en.wikipedia.org/wiki/Domain-specific_language
[wikip-exec]: https://en.wikipedia.org/wiki/Executable
[wikip-hp]: https://en.wikipedia.org/wiki/Harry_Potter
[wikip-path_var]: https://en.wikipedia.org/wiki/PATH_(variable)
[wikip-shell_builtin]: https://en.wikipedia.org/wiki/Shell_builtin
[wikip-vcs]: https://en.wikipedia.org/wiki/Version_control
