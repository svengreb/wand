<p align="center"><img src="https://raw.githubusercontent.com/svengreb/wand/main/assets/images/repository-hero.svg?sanitize=true"/></p>

<p align="center"><a href="https://github.com/svengreb/wand/releases/latest" target="_blank" rel="noreferrer"><img src="https://img.shields.io/github/release/svengreb/wand.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/svengreb/wand/blob/main/changelog.md" target="_blank" rel="noreferrer"><img src="https://img.shields.io/github/release/svengreb/wand.svg?style=flat-square&label=Changelog&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://pkg.go.dev/github.com/svengreb/wand" target="_blank" rel="noreferrer"><img src="https://img.shields.io/github/release/svengreb/wand.svg?style=flat-square&label=GoDoc&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center"><a href="https://github.com/svengreb/wand/actions?query=workflow%3Aci-go" target="_blank" rel="noreferrer"><img src="https://img.shields.io/github/workflow/status/svengreb/wand/ci-go.svg?style=flat-square&label=CI%20Go&logo=github&logoColor=eceff4&colorA=4c566a"/></a> <a href="https://github.com/svengreb/wand/actions?query=workflow%3Aci-node" target="_blank" rel="noreferrer"><img src="https://img.shields.io/github/workflow/status/svengreb/wand/ci-node.svg?style=flat-square&label=CI%20Node&logo=github&logoColor=eceff4&colorA=4c566a"/></a> <a href="https://codecov.io/gh/svengreb/wand" target="_blank" rel="noreferrer"><img src="https://img.shields.io/codecov/c/github/svengreb/wand/main.svg?style=flat-square&label=Coverage&logo=codecov&logoColor=eceff4&colorA=4c566a"/></a></p>

<p align="center"><a href="https://golang.org/doc/effective_go.html#formatting" target="_blank" rel="noreferrer"><img src="https://img.shields.io/static/v1?style=flat-square&label=Go%20Style%20Guide&message=gofmt&logo=go&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a> <a href="https://github.com/arcticicestudio/styleguide-markdown/releases/latest" target="_blank" rel="noreferrer"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-markdown.svg?style=flat-square&label=Markdown%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIzOSIgaGVpZ2h0PSIzOSIgdmlld0JveD0iMCAwIDM5IDM5Ij48cGF0aCBmaWxsPSJub25lIiBzdHJva2U9IiNEOERFRTkiIHN0cm9rZS13aWR0aD0iMyIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBkPSJNMS41IDEuNWgzNnYzNmgtMzZ6Ii8%2BPHBhdGggZmlsbD0iI0Q4REVFOSIgZD0iTTIwLjY4MyAyNS42NTVsNS44NzItMTMuNDhoLjU2Nmw1Ljg3MyAxMy40OGgtMS45OTZsLTQuMTU5LTEwLjA1Ni00LjE2MSAxMC4wNTZoLTEuOTk1em0tMi42OTYgMGwtMTMuNDgtNS44NzJ2LS41NjZsMTMuNDgtNS44NzJ2MS45OTVMNy45MzEgMTkuNWwxMC4wNTYgNC4xNnoiLz48L3N2Zz4%3D"/></a> <a href="https://github.com/arcticicestudio/styleguide-git/releases/latest" target="_blank" rel="noreferrer"><img src="https://img.shields.io/github/release/arcticicestudio/styleguide-git.svg?style=flat-square&label=Git%20Style%20Guide&logoColor=eceff4&colorA=4c566a&colorB=88c0d0&logo=git"/></a></p>

<blockquote cite="https://en.wikipedia.org/wiki/Harry_Potter_and_the_Philosopher%27s_Stone_(film)">
  <p>“The wand chooses the mage, remember.“</p>
  <footer>— Garrick Ollivander, <cite><em>Harry Potter and the Sorcerer’s Stone</em></cite></footer>
</blockquote>

<p align="center">A simple and powerful toolkit for <a href="https://magefile.org" target="_blank" rel="noreferrer">Mage</a>.</p>

## Features

_wand_ is a toolkit for common and often recurring project processes for the task automation tool [Mage][1].
The provided [API packages][74] allow users to compose their own, reusable set of tasks and helpers or built up on the [reference implementation][56].

- **Adapts to any “normal“ or [“mono“][105] repository layout** — handle as many module _commands_ as you want. _wand_ uses an abstraction by managing every `main` package as _application_ so that tasks can be processed for all or just individual _commands_.
- **Runs any `main` package of a [Go module][39] without the requirement for the user to install it beforehand** — Run any command of a Go module using the [module-aware][119] `pkg@version` syntax, or optionally cache executables in a local directory within the project root, using the [`gotool` runner][93]. See the [“Command Runners“ sections below](#command-runners) for details.
- **Comes with support for basic [Go toolchain][48] commands and popular modules from the Go ecosystem** — run common commands like `go build`, `go install` and `go test` or great tools like [gofumpt][62], [golangci-lint][60] and [gox][61] in no time.

See the [API](#api) and [“Elder Wand“](#elder-wand) sections for more details. The [user guides](#user-guides) for more information about how to build your own tasks and runners and the [examples](#examples) for different repositories layouts (single or [“monorepo“][105]) and use cases.

## Motivation

<!--lint disable no-heading-punctuation-->

### Why Mage?

Every project involves processes that are often recurring. These can mostly be done with the tools supplied with the respective programming language, which in turn, in many cases, involve more time and the memorizing of longer commands with their flags and parameters.
In order to significantly reduce this effort or to avoid it completely, project task automation tools are used which often establish a defined standard to enable the widest possible use and unify tasks. They offer a user-friendly and comfortable interface to handle the processes consistently with time savings and without the need for developers to remember many and/or complex commands.
But these tools come with a cost: the introduction of standards and the restriction to predefined ways how to handle tasks is also usually the biggest disadvantage when it comes to adaptability for use cases that are individual for a single project, tasks that deviate from the standard or not covered by it at all.

[Mage][1] is a project task automation tool which gives the user complete freedom by **not specifying how tasks are solved, but only how they are started and connected with each other**. This is an absolute advantage over tools that force how how a task has to be solved while leaving out individual and project specific preferences.

If you would now ask me “But why not just use [Make][9]?“, my answer would be “Why use a tool that is not native to the programming language it is intended for?“.
_Make_ has somehow become a popular choice as task automation tool for Go projects and up to today I don‘t get it. Don‘t get me wrong: this is no bad talking against _Make_ but a clarification that it is not intended for Go but rather for _C_ projects, e.g. [the Linux kernel][3], since [_Make_ is also written in _C_][38]. Even [Go itself is built using shell and Windows DOS scripts][36] instead of _Make_.
If you take a closer look, _Make_ is nothing more than a [DSL][106] for [shell commands][37] so using shell/Windows DOS scripts directly instead is a way more flexible option. Therefore _Make_ can not fullfil an important criteria: full cross-platform compatibility. The command(s) that each task runs must be available on the system, e.g. other tools must be installed globally and available in the [executable search path][109], as well as requiring the syntax to be compatible with the underlying shell which makes it hard to use [shell builtin][110] commands like `cd`.

In my opinion, **a task automation tool for a project should always be written in the same programming language that it is intended for**. This concept has already been proven for many other languages, e.g. official tools like [cargo][104] for _Rust_ and [NPM][6] for _Node.js_‘s or community projects like [Gradle][7] or [Maven][8] for _Java_. All of them can natively understand and interact with their target programming language to provide the widest range of features, good stability and often also ways to simply extend their functionality through plugin systems.

This is where _Mage_ comes in:

- Written in [pure Go without any external dependencies][22] for fully native compatibility and easy expansion.
- [No installation][102] required.
- Allows to [declare dependencies between targets][97] in a _makefile_-style tree and optionally [runs them in parallel][98].
- Targets can be defined in shared packages and [imported][100] in any [_Magefile_][99]. No mechanics like plugins or extensions required, just use any Go module and the whole Go ecosystem.

### Why _wand_?

While _Mage_ is often already sufficient on its own, I‘ve noticed that I had to implement almost identical tasks over and over again when starting a new project or migrating an existing one to _Mage_. Even though the actual [target functions][101] could be moved into their own Go module to allow to simply [import them][100] in different projects, it was often required to copy & paste code across projects that can not be exported that easily. That was the moment where I decided to create a way that simplifies the integration and usage of _Mage_ without loosing any of its flexibility and dynamic structure.

Please note that this package has mainly been created for my personal use in mind to avoid copying source code between my projects. The default configurations or reference implementation might not fit your needs, but the [API](#api) packages have been designed so that they can be flexibly adapted to different use cases and environments or used to create and compose your own [`wand.Wand`][72].

See the [API](#api) and [“Elder Wand“](#elder-wand) sections to learn how to adapt or extend _wand_ for your project.

<!--lint enable no-heading-punctuation-->

## Wording

Since _wand_ is a toolkit for [Mage][1], is partly makes use of an abstract naming scheme that matches the fantasy of magic which in case of _wand_ has been derived from the fantasy novel [“Harry Potter“][108]. This is mainly limited to the [main “Wand“ interface][72] and the [“Elder Wand“](#elder-wand) reference implementation.
The basic mindset of the API is designed around the concept of **tasks** and the ways to **run** them.

- **Runner** — Components that run a command with parameters in a specific environment, in most cases a (binary) [executable][107] of external commands or Go module `main` packages.
- **Tasks** — Components that are scoped for Mage [“target“][101] usage in order to run an action.

## API

The public _wand_ API is located in the [`pkg`][74] package while the main interface [`wand.Wand`][72], that manages a project and its applications and stores their metadata, is defined in the [`wand`][95] package.

Please see the individual documentations of each package for more details.

### Application Configurations

The [`app`][47] package provides the functionality for application configurations. A [`Config`][78] holds information and metadata of an application that is stored by types that implement the [`Store` interface][65]. The [`NewStore() app.Store`][59] function returns a reference implementation of this interface.

### Command Runners

The [`task`][82] package defines the API for runner of commands. [`Runner`][69] is the base interface while [`RunnerExec` interface][70] is a specialized for (binary) executables of a command.

The package already provides runners for the [Go toolchain][48] and [gotool][93] to handle Go module-based executables:

- **Go Toolchain** — to interact with the [Go toolchain][48], also known as the `go` executable, the [`golang.Runner`][80] can be used.
- **`gotool` Go module-based executables** — to install and run [Go module-based][39] `main` packages, the [`gotool.Runner`][81] makes use of the Go 1.16 `go install` command features.
  1. **(Optional) Go Executable Installation & Caching** — [Go 1.16 introduced `go install` command support for the `pkg@version` module syntax][40] which allows to install commands without “polluting“ a projects `go.mod` file. The resulting executables are placed in the Go executable search path that is defined by the [`GOBIN` environment variable][50] (see the [`go env` command][53] to show or modify the Go toolchain environment).
     The problem is that installed executables will overwrite any previously installed executable of the same module/package regardless of the version. Therefore only one version of an executable can be installed at a time which makes it impossible to work on different projects that make use of the same executable but require different versions.
  2. **UX Before Go 1.16** — The installation concept for `main` package executables was always a somewhat controversial point which unfortunately, partly for historical reasons, did not offer an optimal and user-friendly solution until Go 1.16.
     The [`go` command][48] is a fantastic toolchain that provides many great features one would expect to be provided out-of-the-box from a modern and well designed programming language without the requirement to use a third-party solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging utilities and many more.
     This did not apply for the [`go install` command][52] of Go versions less than 1.16.
     The general problem of tool dependencies was a long-time known issue/weak point of the Go toolchain and was a highly rated change request from the Go community with discussions like [golang/go#30515][25], [golang/go#25922][23] and [golang/go#27653][24] to improve this essential feature. They have been around for quite a long time without a solution that worked without introducing breaking changes and most users and the Go team agree on.
     Luckily, this topic was [finally resolved in the Go release version 1.16][40] and and [golang/go#40276][26] introduced a way to install executables in module mode outside a module.
  3. **UX As Of Go 1.17** — With the [introduction in Go 1.17 of running commands in module-aware mode][112] the (local) installation (and caching) of Go module executables has been made kind of obsolete since `go run` can now be used [to run Go commands][115] in module-aware by passing the package and version suffix as argument, without affecting the `main` module and not "pollute" the `go.mod` file 🎉
     The [`pkg/task/golang/run` package][116] package provides a ready-to-use [task implementation][117]. The runner is therefore halfway obsolete, but there are still some drawbacks that are documented below.
     As of [_wand_ version `0.9.0`][114] the default behavior is to not use a local cache directory anymore to store Gomodule-based command executable but make use of the module-aware `go run pkg@version` support!
     To opt-in to the previous behavior set the [`WithCache` option][113] to `true` when initializing a new runner.
  4. **The Leftover Drawback** — Even though the `go install` command works totally fine to globally install executables, the problem that only a single version can be installed at a time is still left. The executable is placed in the path defined by `go env GOBIN` so the previously installed executable will be overridden. It is not possible to install multiple versions of the same package and `go install` still messes up the local user environment.
  5. **The Workaround** — To work around the leftover drawback, the [`gotool` package][93] provides a runner that uses `go install` under the hood, but allows to place the compiled executable in a custom cache directory instead of `go env GOBIN`. It checks if the executable already exists, installs it if not so, and executes it afterwards.
     The concept of storing dependencies locally on a per-project basis is well-known from the [`node_modules` directory][103] of the [Node][2] package manager [npm][5]. Storing executables in a cache directory within the repository (not tracked by Git) allows to use `go install` mechanisms while not affect the global user environment and executables stored in `go env GOBIN`.
     The runner achieves this by temporarily changing the `GOBIN` environment variable to the custom cache directory during the execution of `go install`.
     The only known disadvantage is the increased usage of storage disk space, but since most Go executables are small in size anyway, this is perfectly acceptable compared to the clearly outweighing advantages. Note that the runner dynamically runs executables based on the given task so the `Validate` method is a _NOOP_.
     This is currently the best workaround to…
     1. install `main` package executables locally for the current user without “polluting“ the `go.mod` file.
     2. install `main` package executables locally for the current user without overriding already installed executables of different versions.
  6. **Future Changes** — The provided runner is still not a clean solution that uses the Go toolchain without any special logic so as soon as the following changes are made to the Go toolchain (Go 1.17 or later), the runner can be made opt-in or removed at all:
  - [golang/go#44469][96] — tracks the process of making `go build` module-aware as well as adding support to `go install` for the `-o` flag like for the `go build` command. The second feature, [mentioned in a comment][118], would make the "install" feature of this runner in (or the whole runner at all) obsolete since commands of Go modules could be run and installed using pure Go toolchain functionality.

### Project Metadata

The [`project`][76] package defines the API for metadata and [VCS][111] information of a project. The [`New(opts ...project.Option) (*project.Metadata, error)`][58] function can be used to create a new [project metadata][79].

The package also already provides a [VCS `Repository` interface reference implementation][66] for [Git][4]:

- **VCS “Git“** — the [`git`][77] package provides VCS utility functions to interact with [Git][4] repositories.

### Tasks

The [`task`][82] package defines the API for tasks. [`Task`][71] is the base interface while [`Exec`][67] and [`GoModule`][68] are a specialized to represent the (binary) executable of either an “external“ or Go module-based command.

The package also already provides tasks for basic [Go toolchain][48] commands and popular modules from the Go ecosystem:

- **`go-mod-upgrade`** — the [`gomodupgrade`][92] package provides a task for the [`github.com/oligot/go-mod-upgrade`][63] Go module command. `go-mod-upgrade` allows to update outdated Go module dependencies interactively. The source code of `go-mod-upgrade` is [available in the GitHub repository][30].
- **`gofumpt`** — the [`gofumpt`][84] package provides a task for the [`mvdan.cc/gofumpt`][73] Go module command. `gofumpt` enforces a stricter format than [`gofmt`][55] and provides additional rules, while being backwards compatible. It is a modified fork of `gofmt` so it can be used as a drop-in replacement.
- **`goimports`** — the [`goimports`][85] package provides a task for the [`golang.org/x/tools/cmd/goimports`][64] Go module command. `goimports` allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in the same style as [`gofmt`][55] so it can be used as a replacement. The source code of `goimports` is [available in the GitHub repository][27].
- **Go** — The [`golang`][86] package provides tasks for [Go toolchain][48] commands.
  - **`build`** — to run the [`build` command of the Go toolchain][49] the task of the [`build`][87] package can be used.
  - **`env`** — to run the [`env` command of the Go toolchain][51] the task of the [`env`][88] package can be used.
  - **`install`** — to run the [`install` command of the Go toolchain][52] the task of the [`install`][89] package can be used.
  - **`run`** — to run the [`run` command of the Go toolchain][54] the task of the [`test`][90] package can be used.
  - **`test`** — to run the [`test` command of the Go toolchain][115] the task of the [`run`][116] package can be used.
- **`golangci-lint`** — the [`golangcilint`][91] package provides a task for the [`github.com/golangci/golangci-lint/cmd/golangci-lint`][60] Go module command. `golangci-lint` is a fast, parallel runner for dozens of Go linters that uses caching, supports YAML configurations and has integrations with all major IDEs. The source code of `golangci-lint` is [available in the GitHub repository][28].
- **`gox`** — the [`gox`][94] package provides a task for the [`github.com/mitchellh/gox`][61] Go module command. `gox` is a dead simple, no frills Go cross compile tool that behaves a lot like the standard [Go toolchain `build` command][49]. The source code of `gox` is [available in the GitHub repository][29].

There are also tasks that don‘t need to implement the task API but make use of some “loose“ features like information about a project application are shared as well as the dynamic option system. They can be used without a `task.Runner`, just like a “normal“ package, and provide Go functions/methods that can be called directly:

- **Filesystem Cleaning** — The [`clean`][83] package provides a task to remove directories in a filesystem.

## Usage Guides

In the following sections you can learn how to use the _wand_ reference implementation [“elder wand“](#elder-wand), compose/extend it or simply implement your own tasks and runners.

### Elder Wand

The [`elder`][56] package is the reference implementation of the main [`wand.Wand`][72] interface that provides common Mage tasks and stores configurations and metadata for applications of a project. Next to task methods for the Go toolchain and Go module commands, it comes with additional methods like `Validate` to ensure that the _wand_ is initialized properly and operational.

Create your [_Magefile_][99], e.g `magefile.go`, and use the [`New`][57] function to initialize a new wand and register any amount of applications.
Create a global variable of type `*elder.Elder` and assign the created “elder wand“ to make it available to all functions in your _Magefile_. Even though global variables are a bad practice and should be avoid at all, it‘s totally fine for your task automation since it is non-production code.

Note that the _Mage_ specific **`// +build mage` [build constraint][75], also known as a build tag, is important** in order to mark the file as _Magefile_. See the [official _Mage_ documentation][99] for more details.

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

Now you can create [_Mage_ target functions][101] using the task methods of the “elder wand“.

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

_wand_ comes with tasks and runners for common [Go toolchain][48] commands, [gotool][93] to handle Go module-based executables and other popular modules from the Go ecosystem, but the chance is high that you want to build your own for your specific use cases.

#### Custom Tasks

To create your own task that is compatible with the _wand_ API, implement the [`Task`][71] base interface or any of its specialized interfaces. The `Kind() task.Kind` method must return [`KindBase`][43] while `Options() task.Options` can return anything since [`task.Options`][46] is just an alias for `interface{}`.

1. If your task is **intended for an executable command** you need to implement the [`Exec`][67] interface where…
   - the `Kind() task.Kind` method must return [`KindExec`][44].
   - the `BuildParams() []string` method must return all the parameters that should be passed to the executable.
2. If your task is **intended for the `main` package of a Go module**, so basically also an executable command, you need to implement the [`GoModule`][68] interface where…
   - the `Kind() task.Kind` method must return [`KindGoModule`][45].
   - the `BuildParams() []string` method must return all the parameters that should be passed to the executable that was compiled from the `main` package of the Go module.
   - the returned type of the `ID() *project.GoModuleID` method must provide the import path and module version of the target `main` package.

For sample code of a custom task please see the [examples](#examples) section.
Based on your task kind, you can also either use one of the [already provided command runners](#command-runners), like for the [Go toolchain][86] and [gotool][93], or [implement your own custom runner](#custom-runners).

#### Custom Runners

To create your own command runner that is compatible with the _wand_ API, implement the [`Runner`][69] base interface or any of its specialized interfaces. The `Handles() Kind` method must return the [`Kind`][42] that can be handled while the actual business logic of `Validate() errors` is not bound to any constraint, but like the method names states, should ensure that the runner is configured properly and is operational. The `Run(task.Task) error` method represents the main functionality of the interface and is responsible for running the given [`task.Task`][71] by passing all task parameters, obtained through the `BuildParams() []string` method, and finally execute the configured file. Optionally you can also inspect and use its [`task.Options`][46] by casting the type returned from the `Options() task.Options` method.

1. If your runner is **intended for an executable command** you need to implement the [`RunnerExec`][70] interface where…
   - the `Handles() Kind` method can return kinds like [`KindExec`][44] or [`KindGoModule`][45].
   - the `Run(task.Task) error` method runs the given [`task.Task`][71] by passing all task parameters, obtained through the `BuildParams() []string` method, and finally execute the configured file.
   - it is recommended that the `Validate() error` method tests if the executable file of the command exists at the configured path in the target filesystem or maybe also check other (default) paths if this is not the case. It is also often a good preventative measure to prevent problems to check that the current process actually has permissions to read and execute the file.

For a sample code of a custom command runner please see the [examples](#examples) section.
Based on the kind your command runner can handle, you can also either use one of the [already provided tasks](#tasks) or [implement your own custom task](#custom-task).

## Examples

To learn how to use the _wand_ API and its packages, the [`examples` repository directory][31] contains code samples for multiple use cases:

- **Create your own command runner** — The [`custom_runner`][32] directory contains code samples to demonstrate [how to create a custom command runner](#custom-runners). The `FruitMixerRunner` struct implements the [`RunnerExec`][70] interface for the imaginary `fruitctl` application.
- **Create your own task** — The [`custom_task`][33] directory contains code samples to demonstrate [how to create a custom task](#custom-tasks). The `MixTask` struct implements the [`Exec`][67] interface for the imaginary `fruitctl` application.
- **Usage in a [monorepo][105] layout** — The [`monorepo`][34] directory contains code samples to demonstrate the usage in a _monorepo_ layout for three example applications `cli`, `daemon` and `promexp`. The _Magefile_ provides a `build` target to build all applications. Each application also has a dedicated `:build` target using the [`mg.Namespace`][41] to only build it individually.
- **Usage with a simple, single command repository layout** — The [`simple`][35] directory contains code samples to demonstrate the usage in a “simple“ repository that only provides a single command. The _Magefile_ provides a `build` target to build the `fruitctl` application.

## Contributing

_wand_ is an open source project and contributions are always welcome!

There are many ways to contribute, from [writing- and improving documentation and tutorials][13], [reporting bugs][12], [submitting enhancement suggestions][14] that can be added to _wand_ by [submitting pull requests][18].

Please take a moment to read the [contributing guide][21] to learn about the development process, the [styleguides][19] to which this project adheres as well as the [branch organization][11] and [versioning][20] model.

The guide also includes information about [minimal, complete, and verifiable examples][17] and other ways to contribute to the project like [improving existing issues][16] and [giving feedback on issues and pull requests][15].

<p align="center">Copyright &copy; 2019-present <a href="https://www.svengreb.de" target="_blank" rel="noreferrer">Sven Greb</a></p>

<p align="center"><a href="https://github.com/svengreb/wand/blob/main/license" target="_blank" rel="noreferrer"><img src="https://img.shields.io/static/v1.svg?style=flat-square&label=License&message=MIT&logoColor=eceff4&logo=github&colorA=4c566a&colorB=88c0d0"/></a></p>

[1]: https://magefile.org
[2]: https://nodejs.org
[3]: https://www.kernel.org
[4]: https://git-scm.com
[5]: https://www.npmjs.com
[6]: https://npm.community
[7]: https://gradle.org
[8]: https://maven.apache.org
[9]: https://www.gnu.org/software/make
[11]: https://github.com/svengreb/wand/blob/main/contributing.md#branch-organization
[12]: https://github.com/svengreb/wand/blob/main/contributing.md#bug-reports
[13]: https://github.com/svengreb/wand/blob/main/contributing.md#documentations
[14]: https://github.com/svengreb/wand/blob/main/contributing.md#enhancement-suggestions
[15]: https://github.com/svengreb/wand/blob/main/contributing.md#give-feedback-on-issues-and-pull-requests
[16]: https://github.com/svengreb/wand/blob/main/contributing.md#improve-issues
[17]: https://github.com/svengreb/wand/blob/main/contributing.md#mcve
[18]: https://github.com/svengreb/wand/blob/main/contributing.md#pull-requests
[19]: https://github.com/svengreb/wand/blob/main/contributing.md#styleguides
[20]: https://github.com/svengreb/wand/blob/main/contributing.md#versioning
[21]: https://github.com/svengreb/wand/blob/main/contributing.md
[22]: https://github.com/magefile/mage/blob/d30a2cfe/go.mod
[23]: https://github.com/golang/go/issues/25922
[24]: https://github.com/golang/go/issues/27653
[25]: https://github.com/golang/go/issues/30515
[26]: https://github.com/golang/go/issues/40276
[27]: https://github.com/golang/tools/tree/master/cmd/goimports
[28]: https://github.com/golangci/golangci-lint/tree/master/cmd/golangci-lint
[29]: https://github.com/mitchellh/gox
[30]: https://github.com/oligot/go-mod-upgrade
[31]: https://github.com/svengreb/wand/tree/main/examples
[32]: https://github.com/svengreb/wand/tree/main/examples/custom_runner
[33]: https://github.com/svengreb/wand/tree/main/examples/custom_task
[34]: https://github.com/svengreb/wand/tree/main/examples/monorepo
[35]: https://github.com/svengreb/wand/tree/main/examples/simple
[36]: https://github.com/golang/go/tree/926994fd/src
[37]: https://www.gnu.org/software/make/manual/html_node/Choosing-the-Shell.html
[38]: https://savannah.gnu.org/git/?group=make
[39]: https://golang.org/ref/mod
[40]: https://golang.org/doc/go1.16#modules
[41]: https://pkg.go.dev/github.com/magefile/mage/mg#Namespace
[42]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Kind
[43]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#KindBase
[44]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#KindExec
[45]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#KindGoModule
[46]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Options
[47]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app
[48]: https://pkg.go.dev/cmd/go
[49]: https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies
[50]: https://pkg.go.dev/cmd/go/#hdr-Environment_variables
[51]: https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information
[52]: https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
[53]: https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
[54]: https://pkg.go.dev/cmd/go#hdr-Test_packages
[55]: https://pkg.go.dev/cmd/gofmt
[56]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder
[57]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder#New
[58]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#New
[59]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#NewStore
[60]: https://pkg.go.dev/github.com/golangci/golangci-lint/cmd/golangci-lint
[61]: https://pkg.go.dev/github.com/mitchellh/gox
[62]: https://github.com/mvdan/gofumpt
[63]: https://pkg.go.dev/github.com/oligot/go-mod-upgrade
[64]: https://pkg.go.dev/golang.org/x/tools/cmd/goimports
[65]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#Store
[66]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project/vcs#Repository
[67]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Exec
[68]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#GoModule
[69]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Runner
[70]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#RunnerExec
[71]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Task
[72]: https://pkg.go.dev/github.com/svengreb/wand#Wand
[73]: https://pkg.go.dev/mvdan.cc/gofumpt
[74]: https://pkg.go.dev/github.com/svengreb/wand/pkg
[75]: https://pkg.go.dev/pkg/go/build/#hdr-Build_Constraints
[76]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project
[77]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project/vcs/git
[78]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#Config
[79]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#Metadata
[80]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang#Runner
[81]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gotool#Runner
[82]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task
[83]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/fs/clean
[84]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gofumpt
[85]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/goimports
[86]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang
[87]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/build
[88]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/env
[89]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/install
[90]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/test
[91]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golangcilint
[92]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gomodupgrade
[93]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gotool
[94]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gox
[95]: https://pkg.go.dev/github.com/svengreb/wand
[96]: https:github.com/golang/go/issues/44469#issuecomment-784534876
[97]: https://magefile.org/dependencies
[98]: https://magefile.org/dependencies/#parallelism
[99]: https://magefile.org/magefiles
[100]: https://magefile.org/importing
[101]: https://magefile.org/targets
[102]: https://magefile.org/zeroinstall
[103]: https://docs.npmjs.com/cli/v7/configuring-npm/folders#node-modules
[104]: https://doc.rust-lang.org/stable/cargo
[105]: https://trunkbaseddevelopment.com/monorepos
[106]: https://en.wikipedia.org/wiki/Domain-specific_language
[107]: https://en.wikipedia.org/wiki/Executable
[108]: https://en.wikipedia.org/wiki/Harry_Potter
[109]: https://en.wikipedia.org/wiki/PATH_(variable)
[110]: https://en.wikipedia.org/wiki/Shell_builtin
[111]: https://en.wikipedia.org/wiki/Version_control
[112]: https://go.dev/doc/go1.17#go%20run
[113]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/run#WithCache
[114]: https://pkg.go.dev/github.com/svengreb/wand@v0.9.0
[115]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
[116]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/run
[117]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/run#Task
[118]: https://github.com/golang/go/issues/44469#issuecomment-784534876
[119]: https://go.dev/ref/mod#mod-commands
