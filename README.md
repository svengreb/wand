<p align="center"><img src="https://github.com/svengreb/wand/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

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

Please see the [“Design & Usage“](#design--usage) section for more details about the [API](#api) packages and [“Elder Wand“](#elder-wand) reference implementation.

## Motivation

<!--lint disable no-heading-punctuation-->

### Why Mage?

Every project involves processes that are often recurring. These can mostly be done with the tools supplied with the respective programming language, which in turn, in many cases, involve more time and the memorizing of longer commands with their flags and parameters.
In order to reduce this effort or to avoid it completely, project task automation tools are used which often establish a defined standard to enable the widest possible use and unify tasks. They offer a user-friendly and comfortable interface to handle the processes consistently with time savings and without the need for developers to remember many and/or complex commands.
But these tools come with a cost: the introduction of standards and the restriction to predefined ways how to handle tasks is also usually the biggest disadvantage when it comes to adaptability for use cases that are individual for a single project, tasks that deviate from the standard or not covered by it at all.

[Mage][] is a project task automation tool which gives the user complete freedom by **not specifying how tasks are solved, but only how they are started and connected with each other**. This is an absolute advantage over tools that force how how a task has to be solved while leaving out individual and project specific preferences.

If you would now ask me “But why not just use [Make][]?“, my answer would be “Why use a tool that is not native to the programming language it is intended for?“.
_Make_ has somehow become a popular choice as task automation tool for Go projects and up to today I don‘t get it. Don‘t get me wrong: this is no bad talking against _Make_ but a clarification that it is not intended for Go but rather for _C_ projects, e.g. [the Linux kernel][linux], since [_Make_ is also written in _C_][gnu-make-repo]. Even [Go itself is built using shell and Windows DOS scripts][gh-tree-golang/go/src] instead of _Make_.
If you take a closer look, _Make_ is nothing more than a [DSL][wikip-dsl] for [shell commands][gnu-make-docs-shell] so using shell/Windows DOS scripts directly instead is a way more flexible option. Therefore _Make_ can not fullfil an important criteria: full cross-platform compatibility. The command(s) that each task runs must be available on the system, e.g. other tools must be installed globally and available in the [executable search path][wikip-path_var], as well as requiring the syntax to be compatible with the underlying shell which makes it hard to use [shell builtin][wikip-shell_builtin] commands like `cd`.

In my opinion, **a task automation tool for a project should always be written in the same programming language that it is intended for**. This concept has already been proven for many other languages, e.g. official tools like [cargo][rust-docs-cargo] for _Rust_ and [NPM][npm-com] for _Node.js_‘s or community projects like [Gradle][] or [Maven][] for _Java_. All of them can natively understand and interact with their target programming language to provide the widest range of features, good stability and often also ways to simply extend their functionality through plugin systems.

This is where _Mage_ comes in:

- Written in [pure Go without any external dependencies][gh-blob-magefile/mage/go.mod] for fully native compatibility and easy expansion.
- [No installation][mage-zero_install] required.
- Allows to [declare dependencies between tasks][mage-deps] in a _makefile_-style tree and optionally [runs them in parallel][mage-deps#paral].
- Tasks can be defined in shared packages and [imported in any _Magefile_][mage-importing]. No mechanics like plugins or extensions required, just use any Go module and the whole Go ecosystem.

### Why _wand_?

While _Mage_ is often already sufficient on its own, I‘ve noticed that I had to implement almost identical tasks over and over again when starting a new project or migrating an existing one to _Mage_. Even though the actual [task functions][mage-files] could be moved into their own Go module to allow to simply [import them][mage-importing] in different projects, it was often required to copy & paste code across projects that can not be exported that easily. That was the moment where I decided to create a way that simplifies the integration and usage of _Mage_ without loosing any of its flexibility and dynamic structure.

significantly reduced effort.

Please note that this package has mainly been created for my personal use in mind to avoid copying source code between my projects. The default configurations or reference implementation might not fit your needs, but the [API](#api) packages have been designed so that they can be flexibly adapted to different use cases and environments or used to create and compose your own [`wand.Wand`][go-pkg-if-wand#wand].

Please see the [“Design & Usage“](#design--usage) section below to learn about the API and how to adapt or extend _wand_ for your project.

<!--lint enable no-heading-punctuation-->

## Design & Usage

_wand_ has been designed as a toolkit for [Mage][] and tries to follow its goal to provide a good developer experience through simple and small interfaces.

### Wording

Since _wand_ is a toolkit for [Mage][], the API has been designed with an abstract naming scheme in mind that matches the fantasy of magic which in case of _wand_ have been derived from the fantasy novel [“Harry Potter“][wikip-hp].
As this might be a bit confusing for some users this section provides a mapping to the actual functionality and code logic of the different API components:

- **Spell Incantation** — An abstract representation of flags and parameters for a command or action, in most cases a (binary) [executable][wikip-exec].
  The naming is inspired by the fact that it is almost only possible to [cast a magic spell][wikip-hp_magic#cast] through a [incantation][wikip-inc]. The parameters can be seen as the _formula_.
- **Caster** — An abstract representation for a command or action, in most cases a (binary) [executable][wikip-exec], that provides corresponding information like the path to the executable.
  The naming is inspired by the fact that a caster can [cast a magic spell][wikip-hp_magic#cast] through a [incantation][wikip-inc]. Command can be seen as the [magicians][wikip-magic#magicians] that cast a magic spell.

### API

The public _wand_ API is located in the [`pkg`][go-pkg-pkg] package while the main interface [`wand.Wand`][go-pkg-if-wand#wand], that manages a project and its applications and stores their metadata, is defined in the [`wand`][go-pkg-wand] package.

Please also see the individual documentations of each package for more details.

#### Application Configurations

The [`pkg/app`][go-pkg-app] package provides the functionality for application configurations. The [`Config` struct type][go-pkg-stc-app#config] holds information and metadata of an application which are stored in defined by [the `Store` interface][go-pkg-if-app#store]. The [`NewStore() app.Store`][go-pkg-func-app#newstore] function returns a reference implementation of the `Store` interface.

#### Spell Incantation Casters

The [`pkg/cast`][go-pkg-cast] package provides caster for spell incantations. The [`BinaryCaster` interface][go-pkg-if-cast#binarycaster] is a specialized [`Caster`][go-pkg-if-cast#caster] to run commands using a (binary) executable.

##### Go Toolchain Caster

The [`pkg/cast/golang/toolchain`][go-pkg-cast/golang/toolchain] package provides a caster to interact with the [Go toolchain][go-pkg-cmd/go], in most cases the `go` executable.

##### "gobin" Go Module Caster

The [`pkg/cast/gobin`][go-pkg-cast/gobin] package provides a caster to install and run [Go module][go-docs-ref-mod] executables using the [`github.com/myitcv/gobin`][go-pkg-github.com/myitcv/gobin] module command.

1. **Go Executable Installation** — When installing a Go executable from within a [Go module][go-docs-ref-mod] directory using the [`go install` command][go-pkg-cmd/go#install], it is installed into the Go executable search path that is defined through the [`GOBIN` environment variable][go-pkg-cmd/go#env_vars] and can also be shown and modified using the [`go env` command][go-pkg-cmd/go#print_env].
   Even though the executable gets installed globally, the [`go.mod` file][go-docs-cmd/go#go.mod] will be updated to include the installed packages since this is the default behavior of the [`go get` command][go-pkg-cmd/go#add_deps] when running in [“module“ mode][go-docs-cmd/go#mod_cmds].
   Next to this problem, the installed executable will also overwrite any executable of the same module/package that was installed already, but maybe from a different version. Therefore only one version of a executable can be installed at a time which makes it impossible to work on different projects that use the same tool but with different versions.
2. **History and Future** — The local installation of executables built from [Go modules][go-docs-ref-mod]/`main` packages has always been a somewhat controversial point which unfortunately, partly for historical reasons, does not offer an optimal and user-friendly solution up to now. The [`go` command][go-pkg-cmd/go] is a fantastic toolchain that provides many great features one would expect to be provided out-of-the-box from a modern and well designed programming language without the requirement to use a third-party solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging utilities and many more. Unfortunately the way the [`go install` command][go-pkg-cmd/go#install] of Go versions less or equal to 1.15 handles the installation is still not optimal.
   The general problem of tool dependencies is a long-time known issue/weak point of the current Go toolchain and is a highly rated change request from the Go community with discussions like [golang/go#30515][gh-golang/go#30515], [golang/go#25922][gh-golang/go#25922] and [golang/go#27653][gh-golang/go#27653] to improve this essential feature, but they‘ve been around for quite a long time without a solution that works without introducing breaking changes and most users and the Go team agree on.
   Luckily, this topic was finally picked up for the next [upcoming Go release version 1.16][gh-golang/go-ms-145] and [golang/go#40276][gh-golang/go#40276] introduces a way to install executables in module mode outside a module. The [release note preview also already includes details about this change][go-docs-tip-rln-1.16#mods] and how installation of executables from Go modules will be handled in the future.
3. **The Workaround** — Beside the great news and anticipation about an official solution for the problem the usage of a workaround is almost inevitable until Go 1.16 is finally released.
   The [official Go wiki][gh-golang/go-wiki] provides a section on [“How can I track tool dependencies for a module?“][gh-golang/go-wiki-mods#tool_deps] that describes a workaround that tracks tool dependencies. It allows to use the Go module logic by using a file like `tools.go` with a dedicated `tools` build tag that prevents the included module dependencies to be picked up included for normal executable builds. This approach works fine for non-main packages, but CLI tools that are only implemented in the `main` package can not be imported in such a file.
   In order to tackle this problem, a user from the community created [gobin][], an experimental, module-aware command to install/run main packages. It allows to install or run main-package commands without “polluting“ the `go.mod` file by default. It downloads modules in version-aware mode into a binary cache path within the [systems cache directory][go-pkg-func-os#usercachedir]. It prevents problems due to already globally installed executables by placing each version in its own directory. The decision to use a cache directory instead of sub-directories within the `GOBIN` path keeps the system clean.
   `gobin` is still in an early development state, but has already received a lot of positive feedback and is used in many projects. There are also members of the core Go team that have contributed to the project and the chance is high that the changes for Go 1.16 were influenced or partially ported from it. See [gobin‘s FAQ page in the repository wiki][gobin-wiki-faq] for more details about the project.
   It is currently the best workaround to…
   1. prevent the Go toolchain to pick up the [`GOMOD` environment variable][go-pkg-cmd/go#add_deps] (see [`go env GOMOD`][go-pkg-cmd/go#add_deps]) that is initialized automatically with the path to the [`go.mod` file][go-docs-cmd/go#go.mod] in the current working directory.
   2. install module/package executables globally without “polluting“ the [`go.mod` file][go-docs-cmd/go#go.mod].
   3. install module/package executables globally without overriding already installed executables of different versions.
4. **The Go Module `Caster`** — To allow to manage the tool dependency problem, this caster uses `gobin` through to prevent the “pollution“ of the project [`go.mod` file][go-docs-cmd/go#go.mod] and allows to...
   1. install `gobin` itself into `GOBIN` (see [`go env GOBIN`][go-pkg-cmd/go#print_env]).
   2. cast any [spell incantation][go-pkg-if-spell#incantation] of kind [`KindGoModule`][go-pkg-const-spell#kindgomodule] by installing the executable globally into the dedicated `gobin` cache.

#### Project Metadata

The [`pkg/project`][go-pkg-project] package provides metadata and [VCS][wikip-vcs] information of a project.

##### VCS "Git"

The [`pkg/project/vcs/git`][go-pkg-project/vcs/git] package provides [VCS][wikip-vcs] utility functions to interact with [Git][] repositories.

#### Spell Incantations

The [`pkg/spell`][go-pkg-spell] package provides spell incantations for different kinds.

##### Filesystem Cleaning Spell Incantation

The [`pkg/spell/fs/clean`][go-pkg-spell/fs/clean] package provides a spell incantation to remove directories in a filesystem. It implements [`spell.GoCode`][go-pkg-if-spell#gocode] and can be used without a [`cast.Caster`][go-pkg-if-cast#caster].

##### "goimports" Go Module Spell Incantation

The [`pkg/spell/goimports`][go-pkg-spell/goimports] package provides a spell incantation for the [`golang.org/x/tools/cmd/goimports`][go-pkg-golang.org/x/tools/cmd/goimports] Go module command that allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in the same style as [`gofmt`][go-pkg-cmd/gofmt] so it can be used as a replacement. The source code of `goimports` is [available in the GitHub repository][gh-golang/tools-tree-cmd/goimports].

##### Go Toolchain Spell Incantations

The [`pkg/spell/golang`][go-pkg-spell/golang] package provides spell incantations for [Go toolchain][go-pkg-cmd/go] commands.

###### "build" Go Toolchain Spell Incantation

The [`pkg/spell/golang`][go-pkg-spell/golang/build] package provides a spell incantation for the [`build` command of the Go toolchain][go-pkg-cmd/go#build].

###### "test" Go Toolchain Spell Incantation

The [`pkg/spell/golang/test`][go-pkg-spell/golang/test] package provides a spell incantation for the [`test` command of the Go toolchain][go-pkg-cmd/go#test].

##### "golangci-lint" Go Module Spell Incantation

The [`pkg/spell/golangcilint`][go-pkg-spell/golangcilint] package provides a spell incantation for the [`github.com/golangci/golangci-lint/cmd/golangci-lint`][go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint] Go module command, a fast, parallel runner for dozens of Go linters that uses caching, supports YAML configurations and has integrations with all major IDEs. The source code of `golangci-lint` is [available in the GitHub repository][gh-golangci/golangci-lint].

##### "gox" Go Module Spell Incantation

The [`pkg/spell/gox`][go-pkg-spell/gox] package provides a spell incantation for the [`github.com/mitchellh/gox`][go-pkg-github.com/mitchellh/gox] Go module command, a dead simple, no frills Go cross compile tool that behaves a lot like the standard [Go toolchain `build` command][go-pkg-cmd/go#build]. The source code of `golangci-lint` is [available in the GitHub repository][gh-mitchellh/gox].

### Elder Wand

The [`elder`][go-pkg-elder] package contains a reference implementation of the main [`wand.Wand`][go-pkg-if-wand#wand] interface that provides common Mage tasks and stores configurations and metadata for applications of a project. Next to task methods for the Go toolchain and Go module commands, it comes with additional methods like `Bootstrap`, that runs initialization tasks to ensure the _wand_ is operational, or `Validate`, that ensures that all casters are properly initialized and available.

<!--lint disable no-tabs-->

```go
// +build mage

package main

import (
  "fmt"
	"os"

	"github.com/svengreb/nib/inkpen"
	"github.com/svengreb/wand/pkg/elder"
	wandProj "github.com/svengreb/wand/pkg/project"
)
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
		{"cli", "Fruit Mixer CLI", "apps/cli"},
		{"daemon", "Fruit Mixer Daemon", "apps/daemon"},
		{"prometheus-exporter", "Fruit Mixer Prometheus Exporter", "apps/promexp"},
	}
	for _, app := range apps {
		if regErr := ew.RegisterApp(app.name, app.displayName, app.pathRel); regErr != nil {
			ew.ExitPrintf(1, nib.ErrorVerbosity, "Failed to register application %q: %v", app.name, regErr)
		}
	}
}
```

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
[gh-tree-golang/go/src]: https://github.com/golang/go/tree/926994fd/src
[git]: https://git-scm.com
[gnu-make-docs-shell]: https://www.gnu.org/software/make/manual/html_node/Choosing-the-Shell.html
[gnu-make-repo]: https://savannah.gnu.org/git/?group=make
[go-docs-cmd/go#go.mod]: https://golang.org/ref/mod#go-mod-file
[go-docs-cmd/go#mod_cmds]: https://golang.org/ref/mod#mod-commands
[go-docs-ref-mod]: https://golang.org/ref/mod
[go-docs-tip-rln-1.16#mods]: https://tip.golang.org/doc/go1.16#modules
[go-pkg-app]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app
[go-pkg-cast]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast
[go-pkg-cast/gobin]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/gobin
[go-pkg-cast/golang/toolchain]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/golang/toolchain
[go-pkg-cmd/go]: https://pkg.go.dev/cmd/go
[go-pkg-cmd/go#add_deps]: https://pkg.go.dev/cmd/go/#hdr-PrintGoenvironment_information
[go-pkg-cmd/go#build]: https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies
[go-pkg-cmd/go#env_vars]: https://pkg.go.dev/cmd/go/#hdr-Environment_variables
[go-pkg-cmd/go#install]: https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
[go-pkg-cmd/go#print_env]: https://pkg.go.dev/cmd/go#hdr-PrintGoenvironment_information
[go-pkg-cmd/go#test]: https://pkg.go.dev/cmd/go#hdr-Test_packages
[go-pkg-cmd/gofmt]: https://pkg.go.dev/cmd/gofmt
[go-pkg-const-spell#kindgomodule]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#KindGoModule
[go-pkg-elder]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder
[go-pkg-func-app#newstore]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#NewStore
[go-pkg-func-os#usercachedir]: https://pkg.go.dev/os/#UserCacheDir
[go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint]: https://pkg.go.dev/github.com/golangci/golangci-lint/cmd/golangci-lint
[go-pkg-github.com/mitchellh/gox]: https://pkg.go.dev/github.com/mitchellh/gox
[go-pkg-github.com/myitcv/gobin]: https://pkg.go.dev/github.com/myitcv/gobin
[go-pkg-golang.org/x/tools/cmd/goimports]: https://pkg.go.dev/golang.org/x/tools/cmd/goimports
[go-pkg-if-app#store]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#Store
[go-pkg-if-cast#binarycaster]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#BinaryCaster
[go-pkg-if-cast#caster]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#Caster
[go-pkg-if-spell#gocode]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#GoCode
[go-pkg-if-spell#incantation]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Incantation
[go-pkg-if-wand#wand]: https://pkg.go.dev/github.com/svengreb/wand#Wand
[go-pkg-pkg]: https://pkg.go.dev/github.com/svengreb/wand/pkg
[go-pkg-project]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project
[go-pkg-project/vcs/git]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project/vcs/git
[go-pkg-spell]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell
[go-pkg-spell/fs/clean]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/fs/clean
[go-pkg-spell/goimports]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/goimports
[go-pkg-spell/golang]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang
[go-pkg-spell/golang/build]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang/build
[go-pkg-spell/golang/test]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang/test
[go-pkg-spell/golangcilint]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golangcilint
[go-pkg-spell/gox]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/gox
[go-pkg-stc-app#config]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#Config
[go-pkg-wand]: https://pkg.go.dev/github.com/svengreb/wand
[gobin-wiki-faq]: https://github.com/myitcv/gobin/wiki/FAQ
[gobin]: https://github.com/myitcv/gobin
[gradle]: https://gradle.org
[linux]: https://www.kernel.org
[mage-deps]: https://magefile.org/dependencies
[mage-deps#paral]: https://magefile.org/dependencies/#parallelism
[mage-files]: https://magefile.org/magefiles
[mage-importing]: https://magefile.org/importing
[mage-zero_install]: https://magefile.org/zeroinstall
[mage]: https://magefile.org
[make]: https://www.gnu.org/software/make
[maven]: https://maven.apache.org
[npm-com]: https://npm.community
[rust-docs-cargo]: https://doc.rust-lang.org/stable/cargo
[trunkbasedev-monorepos]: https://trunkbaseddevelopment.com/monorepos
[wikip-dsl]: https://en.wikipedia.org/wiki/Domain-specific_language
[wikip-exec]: https://en.wikipedia.org/wiki/Executable
[wikip-hp_magic#cast]: https://en.wikipedia.org/wiki/Magic_in_Harry_Potter#Spellcasting
[wikip-hp]: https://en.wikipedia.org/wiki/Harry_Potter
[wikip-inc]: https://en.wikipedia.org/wiki/Incantation
[wikip-magic#magicians]: https://en.wikipedia.org/wiki/Magic_(supernatural)#Magicians
[wikip-path_var]: https://en.wikipedia.org/wiki/PATH_(variable)
[wikip-shell_builtin]: https://en.wikipedia.org/wiki/Shell_builtin
[wikip-vcs]: https://en.wikipedia.org/wiki/Version_control
