<p align="center"><img src="https://github.com/svengreb/wand/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

<p align="center"><a href="https://github.com/svengreb/wand/releases/latest"><img src="https://img.shields.io/github/release/svengreb/wand.svg?style=flat-square&label=Release&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0"/></a></p>

<p align="center">Changelog of <em>wand</em>, a simple and powerful toolkit for <a href="https://magefile.org" target="_blank">Mage</a>.</p>

<!--lint disable no-duplicate-headings no-duplicate-headings-in-section-->

# 0.2.0

![Release Date: 2020-12-07](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-12-07&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.2.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/5) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.2.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/2)

⇅ [Show all commits][gh-compare-tag-v0.1.0_v0.2.0]

This release version comes with a large API breaking change to introduce the new "task" + "runner" based API that uses a “normalized“ naming scheme.

## Features

<details>
<summary><strong>“Task“ API: Simplified usage and “normalized“ naming scheme</strong> — #49 ⇄ #51 (⊶ f51a4bfa)</summary>

↠ With #14 the “abstract“ _wand_ API was introduced with a naming scheme is inspired by the fantasy novel [“Harry Potter“][wikip-hp] that was used to to define interfaces.
The main motivation was to create a matching naming to the overall “magic“ topic and the actual target project [Mage][], but in retrospect this is way too abstract and confusing.

The goal of this change was to…

- rewrite the API to **make it way easier to use**.
- use a **“normal“ naming scheme**.
- improve all **documentations to be more user-scoped** and provide **guides and examples**.

#### New API Concept

The basic mindset of the API will remain partially the same, but it will be designed around the concept of **tasks** and the ways to **run** them.

##### Command Runner

[🅸 `task.Runner`][go-pkg-if-task#runner] is a new base interface that runs a command with parameters in a specific environment. It can be compared to the previous [🅸 `cast.Caster`][go-pkg-if-cast#caster] interface, but provides a cleaner method set accepting the new [🅸 `task.Task`][go-pkg-if-task#task] interface.

- 🅼 `Handles() task.Kind` — returns the supported [task kind][go-pkg-al-task#kind].
- 🅼 `Run(task.Task) error` — runs a command.
- 🅼 `Validate() error` — validates the runner.

The new [🅸 `task.RunnerExec`][go-pkg-if-task#runnerexec] interface is a specialized `task.Runner` and serves as an abstract representation for a command or action, in most cases a (binary) [executable][wikip-exec] of external commands or Go module `main` packages, that provides corresponding information like the path to the executable. It can be compared to the previous [`BinaryCaster`][go-pkg-if-cast#binarycaster] interface, but also comes with a cleaner method set and a more appropriate name.

- 🅼 `FilePath() string` — returns the path to the (binary) command executable.

##### Tasks

[🅸 `task.Task`][go-pkg-if-task#task] is the new interface that is scoped for Mage [“target“][mage-docs-targets] usage. It can be compared to the previous [🅸 `spell.Incantation`][go-pkg-if-spell#incantation] interface, but provides a smaller method set without `Formula() []string`.

- 🅼 `Kind() task.Kind` — returns the [task kind][go-pkg-al-task#kind].
- 🅼 `Options() task.Options` — returns the [task options][go-pkg-if-task#options].

The new [🅸 `task.Exec`][go-pkg-if-task#exec] interface is a specialized `task.Task` and serves as an abstract task for an executable command. It can be compared to the previous [`Binary`][go-pkg-if-spell#binary] interface, but also comes with the new `BuildParams() []string` method that enables a more flexible usage by exposing the parameters for command runner like `task.RunnerExec` and also allows to compose with other tasks. See the Wikipedia page about [the anatomy of a shell CLI][wikip-cli#anaton] for more details about parameters.

- 🅼 `BuildParams() []string` — builds the parameters for a command runner where parameters can consist of options, flags and arguments.
- 🅼 `Env() map[string]string` — returns the task specific environment.

The new [🅸 `task.GoModule`][go-pkg-if-task#gomodule] interface is a specialized `task.Exec` for a executable Go module command. It can be compared to the previous [`spell.GoModule`][go-pkg-if-spell#gomodule] interface and the method set has not changed except a renaming of the `GoModuleID() *project.GoModuleID` to the more appropriate name `ID() *project.GoModuleID`. See the official [Go module reference documentation][go-ref-mod] for more details about Go modules.

- 🅼 `ID() *project.GoModuleID` — returns the identifier of a Go module.

#### New API Naming Scheme

The following listing shows the new name concept and how the previous API components can be mapped to the changes:

1. **Runner** — A component that runs a command with parameters in a specific environment, in most cases a (binary) [executable][wikip-exec] of external commands or Go module `main` packages. The current API component that can be compared to runners is [🅸 `cast.Caster`][go-pkg-if-cast#caster] and its specialized interfaces.
2. **Tasks** — A component that is scoped for Mage [“target“][mage-docs-targets] usage in order to run a action. The current API component that can be compared to tasks is [🅸 `spell.Incantation`][go-pkg-if-spell#incantation] and its specialized interfaces.

#### API Usage

Even though the API has been changed quite heavily, the basic usage almost did not change.

→ **A `task.Task` can only be run through a `task.Runner`!**

Before a `spell.Incantation` was passed to a `cast.Caster` in order to run it, in most cases a (binary) executable of a command that uses the `Formula() []string` method of `spell.Incantation` to pass the result as parameters.
The new API works the same: A `task.Task` is passed to a `task.Runner` that calls the `BuildParams() []string` method when the runner is specialized for (binary) executable of commands.

#### Improved Documentations

Before the documentation was mainly scoped on technical details, but lacked more user-friendly sections about topics like the way how to implement own API components, how to compose the [“elder“ reference implementation][go-pkg-elder] or usage examples for single or [monorepo][trunkbasedev-monorepos] project layouts.

##### User Guide

Most of the current sections have been rewritten or removed entirely while new sections now provide more user-friendly guides about how to…

- use or compose the [“elder“ reference implementation][go-pkg-elder].
- build own tasks and runners using the new API.
- structure repositories independent of the layout, single or “monorepo“.

##### Usage Examples

Some examples have been added, that are linked and documented in the user guides described above, to show how to…

- use or compose the [“elder“ reference implementation][go-pkg-elder].
- build own tasks and runners using the new API.
- structure repositories independent of the layout, single or “monorepo“.

</details>

# 0.1.0

![Release Date: 2020-11-29](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-11-29&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/4) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/1)

⇅ [Show all commits][gh-compare-tag-init_v0.1.0]

This is the initial release version of _wand_.
The basic project setup, structure and development workflow has been bootstrapped by [the _tmpl-go_ template repository][gh-svengreb/tmpl-go].
The following sections of this version changelog summarize used technologies, explain design decisions and provide an overview of the API and “elder“ reference implementation.

## Features

<details>
<summary><strong>Bootstrap based on “tmpl-go“ template repository</strong> — #1, #2, #4, #12 ⇄ #3, #5, #13 (⊶ dbf11bc0, f1eee4a1, f778fd97, 5d417258)</summary>

<p align="center"><img src="https://github.com/svengreb/tmpl-go/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

↠ Bootstrapped the basic project setup, structure and development workflow [from version 0.3.0][gh-svengreb/tmpl-go-rl-v0.3.0] of the [“tmpl-go“ template repository][gh-svengreb/tmpl-go].
Project specific files like the repository hero image, documentations and GitHub issue/PR templates have been adjusted.

</details>

<details>
<summary><strong>Application configuration store</strong> — #8 ⇄ #9 (⊶ a233575d)</summary>

↠ Like described in [the `/apps` directory documentation][gh-svengreb/tmpl-go-tree-apps] of the _tmpl-go_ template repository, _wand_ also aims to support the [monorepo][trunkbasedev-monorepos] layout.
In order to manage multiple applications, their information and metadata is recorded in a configuration store where each entry is identified by a unique ID, usually the name of the application. The `pkg/app` package provides two interfaces and an unexported struct that implements it that can be used through the exported `NewStore() Store` function.

- 🆃 `pkg/app.Config` — A `struct` type that holds information and metadata of an application.
- 🅸 `pkg/app.Store` — A storage that provides methods to record application configurations:
  - `Add(*Config)` — Adds a application configuration.
  - `Get(string) (*Config, error)` — Returns the application configuration for the given name or nil along with an error when not stored.
- 🆃 `appStore` — A storage for application configurations.
- 🅵 `NewStore() Store` — Creates a new store for application configurations.

</details>

<details>
<summary><strong>Project and VCS repository</strong> — #10, #18 ⇄ #11, #19 (⊶ 3e8add21, 3fa84e35)</summary>

↠ In [GH-9][gh-svengreb/wand#9] the store and configuration for applications has been implemented. _wand_ applications are not standalone but part of a project which in turn is stored in a repository of [a VCS like Git][git-book-intro-vcs]. In case of _wand_ this can also be a [monorepo][trunkbasedev-monorepos] to manage multiple applications, but there is always only a single project which all these applications are part of.
To store project and VCS repository information, some of the newly implemented packages provide the following types:

- 🆃 `pkg/project.Metadata` — A `struct` type that stores information and metadata of a project.
- 🆃 `pkg/project.GoModuleID` — A `struct` type that stores partial information to identify a [Go module][go-ref-mod].
- 🆃 `pkg/vcs.Kind` — A `struct` type that defines the kind of a `pkg/vcs.Repository`.
- 🅸 `pkg/vcs.Repository` — A `interface` type to represents a VCS repository that provides methods to receive repository information:
  - `Kind() Kind` — returns the repository `pkg/vcs.Kind`.
  - `DeriveVersion() error` — derives the repository version based on the `pkg/vcs.Kind`.
  - `Version() interface{}` — returns the repository version.
- 🆃 `pkg/vcs/git.Git` — A `struct` type that implements `pkg/vcs.Repository` to represent a [Git][] repository.
- 🆃 `pkg/vcs/git.Version` — A `struct` type that stores version information and metadata derived from a [Git][] repository.
- 🆃 `pkg/vcs/none.None` — A `struct` type that implements `pkg/vcs.Repository` to represent a nonexistent repository.

</details>

<details>
<summary><strong>Abstract “task“ API: _spell incantation_, _kind_ and _caster_</strong> — #14 ⇄ #15 (⊶ 2b13b840)</summary>

↠ The _wand_ API is inspired by the fantasy novel [“Harry Potter“][wikip-hp] and uses an abstract view to define interfaces. The main motivation to create a matching naming to the overall “magic“ topic and the actual target project [Mage][]. This might be too abstract for some, but is kept understandable insofar as it should allow everyone to use the “task“ API and to derive their own tasks from it.

- 🅸 `cast.Caster` — A `interface` type that casts a `spell.Incantation` using a command for a specific `spell.Kind`:
  - `Cast(spell.Incantation) error` — casts a spell incantation.
  - `Handles() spell.Kind` — returns the spell kind that can be casted.
  - `Validate() error` — validates the caster command.
- 🅸 `cast.BinaryCaster` — A `interface` type that composes `cast.Caster` to run commands using a binary executable:
  - `GetExec() string` — returns the path to the binary executable of the command.
- 🅸 `spell.Incantation` — A `interface` type that is the abstract representation of parameters for a command or action:
  - `Formula() []string` — returns all parameters of a spell.
  - `Kind() Kind` — returns the Kind of a spell.
  - `Options() interface{}` — return the options of a spell.
- 🅸 `cast.Binary` — A `interface` type that composes `cast.Caster` for commands which are using a binary executable:
  - `Env() map[string]string` — returns additional environment variables.
- 🅸 `cast.GoCode` — A `interface` type that composes `cast.Caster` for actions that can be casted without a `cast.Caster`:
  - `Cast() (interface{}, error)` — casts itself.
- 🅸 `cast.GoModule` — A `interface` type that composes `cast.Binary` for commands that are compiled from a [Go module][go-ref-mod]
  - `GoModuleID() *project.GoModuleID` — returns the identifier of a Go module.
- 🆃 `spell.Kind` — A `struct` type that defines the kind of a spell.

The API components can be roughly translated to their purpose:

- `cast.Caster` → an executable command
  It validates the command and defines which `spell.Kind` can be handled by this caster. It could be executed without parameters (`spell.Incantation`), but in most cases needs at least one parameter.
  - `cast.BinaryCaster` → a composed `cast.Caster` to run commands using a binary executable.
    It ensures that the executable file exists and stores information like the path. It could also be executed without parameters (`spell.Incantation`), but would not have any effect im many cases.
- `spell.Incantation` → the parameters of a executable command
  It assemble all parameters based on the given options and ensures the they are correctly formatted for the execution in a shell environment. Except for special incantations like `spell.GoCode` a incantation cannot be used alone but must be passed to a `cast.Caster` that is able to handle the `spell.Kind` of this incantation.
  - `spell.Binary` → a composed `spell.Incantation` to run commands that are using binary executable.
    It can inject or override environment variables in the shell environment in which the the command will be run.
  - `spell.GoCode` → a composed `spell.Incantation` for pure Go code instead of a (binary) executable command.
    It can “cast itself“, e.g. to simply delete a directory using packages like `os` from the Go standard library. It has been designed this way to also allow such tasks to be handled by the incantation API.
  - `spell.GoModule` → a composed `spell.Binary` to run binary commands managed by a [Go module][go-ref-mod], in other words executables installed in `GOBIN` or received via `go get`.
    It requires the module identifier (`path@version`) in order to download and run the executable.

</details>

<details>
<summary><strong>Basic “wand“ API</strong> — #16 ⇄ #17 (⊶ cc9f7c4b)</summary>

↠ In [GH-15][gh-svengreb/wand#15] some parts of the _wand_ API have been implemented in form of spell _incantations_, _kinds_ and _casters_, inspired by the fantasy novel [“Harry Potter“][wikip-hp] as an abstract view to define interfaces. In [GH-9][gh-svengreb/wand#9] and [GH-11][gh-svengreb/wand#11] the API implementations for an application configuration store as well as project and VCS repository metadata were introduced.
These implementations are usable in a combined form via the main _wand_ API that consists of the following types:

- 🅸 `wand.Wand` — A `interface` type that manages a project and its applications and stores their metadata. Applications are registered using a unique name and the stored metadata can be received based on this name:
  - `GetAppConfig(appName string) (app.Config, error)` — returns an application configuration.
  - `GetProjectMetadata() project.Metadata` — returns the project metadata.
  - `RegisterApp(name, displayName, pathRel string) error` — registers a new application.
- 🆃 `wand.ctxKey` — A `struct` type that serves as context key used to wrap a `wand.Wand`.
- 🅵 `wand.GetCtxKey() interface{}` — A `func` type that returns the key used to wrap a `wand.Wand`.
- 🅵 `wand.WrapCtx(parentCtx context.Context, wand Wand) context.Context` — A `func` type that wraps the given `wand.Wand` into the parent context. Use `wand.GetCtxKey() interface{}` to receive the key used to wrap the `wand.Wand`.

</details>

<details>
<summary><strong>Go toolchain “caster“</strong> — #20 ⇄ #21 (⊶ 55e8eb46)</summary>

↠ To use the Go toolchain, also known as [the `go` command][go-pkg-cmd/go], a new [caster][go-pkg-if-cast#caster] (introduced in #14) has been implemented.
The new [`ErrCast`][go-pkg-stc-cast#errcast] `struct` type unifies the handling of errors in the [cast][go-pkg-cast] package.

The [`Validate` function][go-pkg-fn-cast#validate] of the new caster returns an error of type `*cast.ErrCast` when the `go` binary executable does not exist at the configured path or when it is also not available in the [executable search paths][wikip-path_var] of the current environment.

</details>

<details>
<summary><strong>“gobin“ Go module caster</strong> — #22 ⇄ #23 (⊶ 95c22a00)</summary>

##### Go Executable Installation

When installing a Go executable from within a [Go module][go-ref-mod] directory using the [`go install` command][go-pkg-cmd/go#install], it is installed into the Go executable search path that is defined through [the `GOBIN` environment variable][go-pkg-cmd/go#env_vars] and can also be shown and modified using the [`go env` command][go-pkg-cmd/go#print_env]. Even though the executable gets installed globally, the [`go.mod` file][go-ref-mod#file] will be updated to include the installed packages since this is the default behavior of [the `go get` command][go-pkg-cmd/go#get] when running in [_module_ mode][go-docs-cmd-go#mod_aware_cmds].

Next to this problem, the installed executable will also overwrite any executable of the same module/package that was installed already, but maybe from a different version. Therefore only one version of a executable can be installed at a time which makes it impossible to work on different projects that use the same tool but with different versions.

##### History & Future

The local installation of executables built from Go modules/packages has always been a somewhat controversial point which unfortunately, partly for historical reasons, does not offer an optimal and user-friendly solution up to now. The [`go` command][go-pkg-cmd/go] is a fantastic toolchain that provides many great features one would expect to be provided out-of-the-box from a modern and well designed programming language without the requirement to use a third-party solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging utilities and many more.
Unfortunately the way the [`go install` command][go-pkg-cmd/go#install] of Go versions less or equal to 1.15 handles the installation of an Go module/package executable is still not optimal.

The general problem of tool dependencies is a long-time known issue/weak point of the current Go toolchain and is a highly rated change request from the Go community with discussions like [golang/go#30515][gh-golang/go#30515], [golang/go#25922][gh-golang/go#25922] and [golang/go#27653][gh-golang/go#27653] to improve this essential feature, but they‘ve been around for quite a long time without a solution that works without introducing breaking changes and most users and the Go team agree on.
Luckily, this topic was finally picked up for [the next upcoming Go release version 1.16][gh-ms-golang/go#145] and [gh-golang/go#40276][] introduces a way to install executables in module mode outside a module. The [release note preview also already includes details about this change][go-docs-tip-rln-1.16#mod] and how installation of executables from Go modules will be handled in the future.

##### The Workaround

Beside the great news and anticipation about an official solution for the problem the usage of a workaround is almost inevitable until Go 1.16 is finally released.

The [official Go wiki][gh-golang/go-wiki] provides a section on [“How can I track tool dependencies for a module?”][go-wiki-tool_dep] that describes a workaround that tracks tool dependencies. It allows to use the Go module logic by using a file like `tools.go` with a dedicated `tools` build tag that prevents the included module dependencies to be picked up included for normal executable builds. This approach works fine for non-main packages, but CLI tools that are only implemented in the `main` package can not be imported in such a file.

In order to tackle this problem, a user from the community created [gobin][gh-myitcv/gobin], _an experimental, module-aware command to install/run main packages_.
It allows to install or run main-package commands without “polluting“ the `go.mod` file by default. It downloads modules in version-aware mode into a binary cache path within [the systems cache directory][go-pkg-os#cachedir].
It prevents problems due to already globally installed executables by placing each version in its own directory. The decision to use a cache directory instead of sub-directories within the `GOBIN` path keeps the system clean.

_gobin_ is still in an early development state, but has already received a lot of positive feedback and is used in many projects. There are also members of the core Go team that have contributed to the project and the chance is high that the changes for Go 1.16 were influenced or partially ported from it.
It is currently the best workaround to…

1. …prevent the Go toolchain to pick up the [`GOMOD` environment variable][go-pkg-cmd/go#print_env] (see [`go env GOMOD`][go-pkg-cmd/go#print_env]) that is initialized automatically with the path to the `go.mod` file in the current working directory.
2. …install module/package executables globally without “polluting“ the `go.mod` file.
3. …install module/package executables globally without overriding already installed executables of different versions.

See [gobin‘s FAQ page][gh-myitcv/gobin-wiki-faq] in the repository wiki for more details about the project.

#### The Go Module Caster

To allow to manage the tool dependency problem, _wand_ uses `gobin` through [a new caster][go-pkg-stc-cast/gobin#caster] that prevents the “pollution“ of the project `go.mod` file and allows to…

1. …install `gobin` itself into `GOBIN` ([`go env GOBIN`][go-pkg-cmd/go#print_env]).
2. …cast any [spell incantation][go-pkg-spell#incantation] of kind [`KindGoModule`][go-pkg-stc-spell#kindgomodule] by installing the executable globally into the dedicated `gobin` cache.

</details>

<details>
<summary><strong>Spell incantation options “mixin“</strong> — #25 ⇄ #26 (⊶ 9ae4f892)</summary>

↠ To allow to compose, manipulate and read spell incantation options after the initial creation, two new types have been added for the [spell][go-pkg-spell] package:

- 🅸 `spell.Options` — A `interface` type as a generic representation for `spell.Incantation` options.
- 🅸 `spell.Mixin` — A `interface` type that allows to compose functions that process `spell.Options` of `spell.Incantation`s.
  - `Apply(Options) (Options, error)` — applies generic `spell.Options` to `spell.Incantation` options.

</details>

<details>
<summary><strong>Spell incantation for Go toolchain <code>build</code> command</strong> — #27 ⇄ #28 (⊶ 060b3328)</summary>

↠ To run the `go build` command of the Go toolchain, a new [`spell.Incantation`][go-pkg-if-spell#incantation] has been implemented in the new [build][go-pkg-spell/golang/build] package that can be used through a [Go toolchain caster][go-pkg-stc-cast/golang#caster].
The spell incantation is configurable through the following functions:

- `WithBinaryArtifactName(name string) build.Option` — sets the name for the binary build artifact.
- `WithCrossCompileTargetPlatforms(platforms ...string) build.Option` — sets the names of cross-compile platform targets.
- `WithFlags(flags ...string) build.Option` — sets additional flags to pass to the Go `build` command along with the base Go flags.
- `WithGoOptions(goOpts ...spellGo.Option) build.Option` — sets shared Go toolchain commands options.
- `WithOutputDir(dir string) build.Option` — sets the output directory, relative to the project root, for compilation artifacts.

To unify further implementations for the Go toolchain, a new `struct` type is available in the [golang][go-pkg-spell/golang] package to store global/shared Go toolchain options that are shared between multiple Go toolchain commands:

- `WithAsmFlags(asmFlags ...string) golang.Option` — sets flags to pass on each `go tool asm` invocation.
- `WithRaceDetector(enableRaceDetector bool) golang.Option` — indicates if the race detector should be enabled.
- `WithTrimmedPath(enableTrimPath bool) golang.Option` — indicates if all file system paths should be removed from the resulting executable.
- `WithEnv(env map[string]string) golang.Option` — adds or overrides Go toolchain command specific environment variables.
- `WithFlags(flags ...string) golang.Option` — sets additional Go toolchain command flags.
- `WithFlagsPrefixAll(flagsPrefixAll bool) golang.Option` — indicates if the values of `-asmflags` and `-gcflags` should be prefixed with the `all=` pattern in order to apply to all packages.
- `WithGcFlags(gcFlags ...string) golang.Option` — sets flags to pass on each `go tool compile` invocation.
- `WithLdFlags(ldFlags ...string) golang.Option` — sets flags to pass on each `go tool link` invocation.
- `WithMixins(mixins ...spell.Mixin) golang.Option` — sets `spell.Mixin`s that can be applied by option consumers.
- `WithTags(tags ...string) golang.Option` — sets Go toolchain tags.

The new [`CompileFormula(opts ...Option) []string` function][go-pkg-fn-spell/golang#compileformula] can be used to compile the formula for these options.

</details>

<details>
<summary><strong>Spell incantation for Go toolchain <code>test</code> command</strong> — #29 ⇄ #30 (⊶ 166a2dc0)</summary>

↠ To run the `go test` command of the Go toolchain, a new [`spell.Incantation`][go-pkg-if-spell#incantation] is available in the new [test][go-pkg-spell/golang/test] package that can be used through a [Go toolchain caster][go-pkg-stc-cast/golang#caster].
The spell incantation is customizable through the following functions:

- `WithBlockProfileOutputFileName(blockProfileOutputFileName string) test.Option` — sets the file name for the Goroutine blocking profile file.
- `WithCoverageProfileOutputFileName(coverageProfileOutputFileName string) test.Option` — sets the file name for the test coverage profile file.
- `WithCPUProfileOutputFileName(cpuProfileOutputFileName string) test.Option` — sets the file name for the CPU profile file.
- `WithBlockProfile(withBlockProfile bool) test.Option` — indicates if the tests should be run with a Goroutine blocking profiling.
- `WithCoverageProfile(withCoverageProfile bool) test.Option` — indicates if the tests should be run with coverage profiling.
- `WithCPUProfile(withCPUProfile bool) test.Option` — indicates if the tests should be run with CPU profiling.
- `WithFlags(flags ...string) test.Option` — sets additional flags that are passed to the Go "test" command along with the shared Go flags.
- `WithGoOptions(goOpts ...spellGo.Option) test.Option` — sets shared Go toolchain command options.
- `WithMemProfile(withMemProfile bool) test.Option` — indicates if the tests should be run with memory profiling.
- `WithMemoryProfileOutputFileName(memoryProfileOutputFileName string) test.Option` — sets the file name for the memory profile file.
- `WithMutexProfile(withMutexProfile bool) test.Option` — indicates if the tests should be run with mutex profiling.
- `WithMutexProfileOutputFileName(mutexProfileOutputFileName string) test.Option` — sets the file name for the mutex profile file.
- `WithOutputDir(outputDir string) test.Option` — sets the output directory, relative to the project root, for reports like coverage or benchmark profiles.
- `WithoutCache(withoutCache bool) test.Option` — indicates if the tests should be run without test caching that is enabled by Go by default.
- `WithPkgs(pkgs ...string) test.Option` — sets the list of packages to test.
- `WithTraceProfile(withTraceProfile bool) test.Option` — indicates if the tests should be run with trace profiling.
- `WithTraceProfileOutputFileName(traceProfileOutputFileName string) test.Option` — sets the file name for the execution trace profile file.
- `WithVerboseOutput(withVerboseOutput bool) test.Option` — indicates if the test output should be verbose.

</details>

<details>
<summary><strong>Spell incantation for <code>golang.org/x/tools/cmd/goimports</code> Go module</strong> — #31 ⇄ #32 (⊶ 8c9b450c)</summary>

↠ The [golang.org/x/tools/cmd/goimports][go-pkg-golang.org/x/tools/cmd/goimports] Go module allows to update Go import lines, adding missing ones and removing unreferenced ones. It also formats code in the same style as [gofmt][go-pkg-cmd/gofmt] so it can be used as a replacement. The source code for the `goimports` command can be found in the [golang/tools][gh-golang/tools-tree-cmd/goimports] repository.

To configure and run the `goimports` command, a new [`spell.Incantation`][go-pkg-if-spell#incantation] is available in the new [goimports][go-pkg-spell/goimports] package that can be casted using the [gobin caster][go-pkg-stc-cast/gobin#caster] or any other [spell caster][go-pkg-if-cast#caster] that handles [spell incantations][go-pkg-if-spell#incantation] of kind [`KindGoModule`][go-pkg-const-spell#kindgomodule].

The spell incantation is customizable through the following functions:

- `WithEnv(env map[string]string) goimports.Option` — sets the spell incantation specific environment.
- `WithExtraArgs(extraArgs ...string) goimports.Option` — sets additional arguments to pass to the `goimports` command.
- `WithListNonCompliantFiles(listNonCompliantFiles bool) goimports.Option` — indicates whether files, whose formatting are not conform to the style guide, are listed.
- `WithLocalPkgs(localPkgs ...string) goimports.Option` — sets local packages whose imports will be placed after 3rd-party packages.
- `WithModulePath(path string) goimports.Option` — sets the `goimports` module import path. Defaults to `goimports.DefaultGoModulePath`.
- `WithModuleVersion(version *semver.Version) goimports.Option` — sets the `goimports` module version. Defaults to `goimports.DefaultGoModuleVersion`.
- `WithPaths(paths ...string) goimports.Option` — sets the paths to search for Go source files. By default all directories are scanned recursively starting from the current working directory.
- `WithPersistedChanges(persistChanges bool) goimports.Option` — indicates whether results are written to the source files instead of standard output.
- `WithReportAllErrors(reportAllErrors bool) goimports.Option` — indicates whether all errors should be printed instead of only the first 10 on different lines.
- `WithVerboseOutput(verbose bool) goimports.Option` — indicates whether the output should be verbose.

</details>

<details>
<summary><strong>Spell incantation for <code>github.com/golangci/golangci-lint</code> Go module</strong> — #33 ⇄ #34 (⊶ 11c9f627)</summary>

↠ The [github.com/golangci/golangci-lint][go-pkg-github.com/golangci/golangci-lint] Go module provides the `golangci-lint` command, a fast, parallel runner for dozens of Go linters Go that uses caching, supports YAML configurations and has integrations with all major IDEs. The source code for the `golangci-lint` command can be found in the [golangci/golangci-lint][gh-golangci/golangci-lint-tree-cmd/golangci-lint] repository.

To configure and run the `golangci-lint` command, a new [`spell.Incantation`][go-pkg-if-spell#incantation] is available in the new [golangcilint][go-pkg-spell/golangcilint] package that can be casted using the [gobin caster][go-pkg-stc-cast/gobin#caster] or any other [spell caster][go-pkg-if-cast#caster] that handles [spell incantations][go-pkg-if-spell#incantation] of kind [`KindGoModule`][go-pkg-const-spell#kindgomodule].

The spell incantation is customizable through the following functions:

- `WithArgs(args ...string) golangcilint.Option` — sets additional arguments to pass to the `golangci-lint` module command.
- `WithEnv(env map[string]string) golangcilint.Option` — sets the spell incantation specific environment.
- `WithModulePath(path string) golangcilint.Option` — sets the `golangci-lint` module command import path. Defaults to `golangcilint.DefaultGoModulePath`.
- `WithModuleVersion(version *semver.Version) golangcilint.Option` — sets the `golangci-lint` module version. Defaults to `golangcilint.DefaultGoModuleVersion`.
- `WithVerboseOutput(verbose bool) golangcilint.Option` — indicates whether the output should be verbose.

</details>

<details>
<summary><strong>Spell incantation for the <code>github.com/mitchellh/gox</code> Go module</strong> — #35 ⇄ #36 (⊶ 4b285060)</summary>

↠ The [github.com/mitchellh/gox][go-pkg-github.com/mitchellh/gox] Go module provides the `gox` command, a dead simple, no frills Go cross compile tool that behaves a lot like the standard Go toolchain `build` command.

To configure and run the `gox` command, a new [`spell.Incantation`][go-pkg-if-spell#incantation] is available in the new [gox][go-pkg-spell/gox] package that can be casted using the [gobin caster][go-pkg-stc-cast/gobin#caster] or any other [spell caster][go-pkg-if-cast#caster] that handles [spell incantations][go-pkg-if-spell#incantation] of kind [`KindGoModule`][go-pkg-const-spell#kindgomodule].

The spell incantation is customizable through the following functions:

- `WithEnv(env map[string]string) gox.Option` — sets the spell incantation specific environment.
- `WithGoCmd(goCmd string) gox.Option` — sets the path to the Go toolchain executable.
- `WithOutputTemplate(outputTemplate string) gox.Option` — sets the name template for cross-compile platform targets. Defaults to `gox.DefaultCrossCompileBinaryNameTemplate`.
- `WithGoOptions(goOpts ...spellGo.Option) gox.Option` — sets shared Go toolchain command options.
- `WithGoBuildOptions(goBuildOpts ...spellGoBuild.Option) gox.Option` — sets options for the Go toolchain `build` command.
- `WithModulePath(path string) gox.Option` — sets the `gox` module command import path. Defaults to `gox.DefaultGoModulePath`.
- `WithModuleVersion(version *semver.Version) gox.Option` — sets the `gox` module version. Defaults to `gox.DefaultGoModuleVersion`.
- `WithVerboseOutput(verbose bool) gox.Option` — indicates whether the output should be verbose.

</details>

<details>
<summary><strong>Spell mixins for Go toolchain options</strong> — #37 ⇄ #38 (⊶ d5a189be)</summary>

↠ To support common use cases for debugging and production optimization, some [spell mixins][go-pkg-if-spell#mixin] have been implemented in the [golang][go-pkg-spell/golang] package:

- 🆂 `MixinImproveDebugging` — A `struct` type that adds linker flags to improve the debugging of binary artifacts. This includes the disabling of inlining and all compiler optimizations tp improve the compatibility for debuggers.
  Note that this mixin will add the `all` prefix for `—gcflags` parameters to make sure all packages are affected. If you disabled the `all` prefix on purpose you need to handle this conflict on your own, e.g. by creating more than one binary artifact each with different build options.
- 🆂 `MixinImproveEscapeAnalysis` — A `struct` type that will add linker flags to improve the escape analysis of binary artifacts.
  Note that this mixin removes the `all` prefix for `—gcflags` parameters to make sure only the target package is affected, otherwise reports for (traverse) dependencies would be included as well. If you enabled the `all` prefix on purpose you need to handle this conflict on your own, e.g. by creating more than one binary artifact each with different build options.
- 🆂 `MixinStripDebugMetadata` — A `struct` type that will add linker flags to strip debug information from binary artifacts. This will include _DWARF_ tables needed for debuggers, but keeps annotations needed for stack traces so panics are still readable. It will also shrink the file size and memory overhead as well as reducing the chance for possible security related problems due to enabled development features or debug information leaks.
  Note that this mixin will add the `all` prefix for `—gcflags` parameters to make sure all packages are affected. If you disabled the `all` prefix on purpose you need to handle this conflict on your own, e.g. by creating more than one binary artifact each with different build options.
- 🆂 `MixinInjectBuildTimeVariableValues` — A `struct` type that will inject build—time values through the `—X` linker flags to populate e.g. application metadata variables.
  It will store a `map[string]string` of key/value pairs to inject to variables at build—time. The key must be the path to the variable in form of `<IMPORT_PATH>.<VARIABLE_NAME>`, e.g. `pkg/internal/support/app.version`. The value is the actual value that will be assigned to the variable, e.g. the application version.
  A field of type [`*project.GoModuleID`][go-pkg-stc-project#gomoduleid] will store partial information about the target Go module to inject the key/value pairs from the data map into.

</details>

<details>
<summary><strong>Go code spell for filesystem cleaning</strong> — #39 ⇄ #40 (⊶ 04a3aeb9)</summary>

↠ To clean paths in a filesystem, like application specific output directories, a new [`GoCode` spell incantation][go-pkg-if-spell#gocode] is available in the new [clean][go-pkg-spell/fs/clean] package that can be used without a [caster][go-pkg-if-cast#caster].

The spell incantation provides the following methods:

- `Clean() ([]string, error)` — removes the configured paths. It returns an error of type `*spell.ErrGoCode` for any error that occurs during the execution of the Go code.

The spell incantation is customizable through the following functions:

- `WithLimitToAppOutputDir(limitToAppOutputDir bool) clean.Option` — indicates whether only paths within the configured application output directory should be allowed.
- `WithPaths(paths ...string) clean.Option` — sets the paths to remove. Note that only paths within the configured application output directory are allowed when `WithLimitToAppOutputDir` is enabled.

</details>

<details>
<summary><strong>Wand reference implementation “elder“</strong> — #41 ⇄ #42 (⊶ 6397641b)</summary>

↠ The default way to use the [_wand_ API][go-pkg-if#wand], with its [casters][go-pkg-cast] and [spells][go-pkg-spell], is the reference implementation [“elder“][go-pkg-elder].
It provides a way to use all _wand_ spells and additionally comes with helper methods to bootstrap a project, validate all _casters_ and simplify logging for process exits:

- `Bootstrap() error` — runs initialization tasks to ensure the wand is operational. This includes the installation of configured caster like [`cast.BinaryCaster`][go-pkg-if-cast#binarycaster] that can handle spell incantations of kind [`spell.KindGoModule`][go-pkg-const-spell#kindgomodule].
- `Clean(appName string, opts ...clean.Option) ([]string, error)` — a [`spell.GoCode`][go-pkg-if-spell#gocode] to remove configured filesystem paths, e.g. output data like artifacts and reports from previous development, test, production and distribution builds. It returns paths that have been cleaned along with an error of type [`*spell.ErrGoCode`][go-pkg-stc-spell#errgocode] when an error occurred during the execution of the Go code. When any error occurs it will be of type [`*app.ErrApp`][go-pkg-stc-app#errapp] or [`*cast.ErrCast`][go-pkg-stc-cast#errcast]. See the [clean][go-pkg-spell/fs/clean] package for all available options.
- `ExitPrintf(code int, verb nib.Verbosity, format string, args ...interface{})` — simplifies the logging for process exits with a suitable [`nib.Verbosity`][go-pkg-stc-github.com/svengreb/nib#verbosity].
- `GetAppConfig(name string) (app.Config, error)` — returns an application configuration. An empty application configuration is returned along with an error of type [`*app.ErrApp`][go-pkg-stc-app#errapp] when there is no configuration in the store for the given name.
- `GetProjectMetadata() project.Metadata` — returns metadata of the project.
- `GoBuild(appName string, opts ...build.Option)` — casts the spell incantation for the [`build`][go-pkg-cmd/go#build] command of the [Go toolchain][go-pkg-cmd/go]. When any error occurs it will be of type [`*app.ErrApp`][go-pkg-stc-app#errapp] or [`*cast.ErrCast`][go-pkg-stc-cast#errcast]. See the [build][go-pkg-spell/golang/build] package for all available options.
- `Goimports(appName string, opts ...goimports.Option) error` — casts the spell incantation for the [golang.org/x/tools/cmd/goimports][go-pkg-golang.org/x/tools/cmd/goimports] Go module command that allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in the same style as [`gofmt` command][go-pkg-cmd/gofmt] so it can be used as a replacement. When any error occurs it will be of type [`*app.ErrApp`][go-pkg-stc-app#errapp] or [`*cast.ErrCast`][go-pkg-stc-cast#errcast].
  See the [goimports][go-pkg-spell/goimports] package for all available options. For more details about `goimports` see [the module documentation][go-pkg-golang.org/x/tools/cmd/goimports]. The source code of `goimports` is [available in the GitHub repository][gh-golang/tools-tree-cmd/goimports].
- `GolangCILint(appName string, opts ...golangcilint.Option) error` — casts the spell incantation for the [github.com/golangci/golangci-lint/cmd/golangci-lint][go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint] Go module command, a fast, parallel runner for dozens of Go linters Go that uses caching, supports YAML configurations and has integrations with all major IDEs. When any error occurs it will be of type [`*app.ErrApp`][go-pkg-stc-app#errapp] or [`*cast.ErrCast`][go-pkg-stc-cast#errcast]. See the [golangcilint][go-pkg-spell/golangcilint] package for all available options.
  For more details about `golangci-lint` see [the module documentation][go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint] and the [official website][golangci-lint]. The source code of `golangci-lint` is [available in the GitHub repository][gh-golangci/golangci-lint].
- `GoTest(appName string, opts ...spellGoTest.Option) error` — casts the spell incantation for the [`test`][go-pkg-cmd/go#test] command of the [Go toolchain][go-pkg-cmd/go]. When any error occurs it will be of type [`*app.ErrApp`][go-pkg-stc-app#errapp] or [`*cast.ErrCast`][go-pkg-stc-cast#errcast]. See the [test][go-pkg-spell/golang/test] package for all available options.
- `Gox(appName string, opts ...spellGox.Option) error` — casts the spell incantation for the [github.com/mitchellh/gox][go-pkg-github.com/mitchellh/gox] Go module command, a dead simple, no frills Go cross compile tool that behaves a lot like the standard Go toolchain [`build`][go-pkg-cmd/go#build] command. When any error occurs it will be of type [`*app.ErrApp`][go-pkg-stc-app#errapp] or [`*cast.ErrCast`][go-pkg-stc-cast#errcast]. See the [gox][go-pkg-spell/gox] package for all available options.
  For more details about `gox` see [the module documentation][go-pkg-github.com/mitchellh/gox]. The source code of `gox` is [available in the GitHub repository][gh-mitchellh/gox].
- `RegisterApp(name, displayName, pathRel string) error` — creates and stores a new application configuration. Note that the package path must be relative to the project root directory!
  It returns an error of type [\*app.ErrApp][go-pkg-stc-app#errapp] when the application path is not relative to the project root directory, when it is not a subdirectory of it or when any other error occurs.
- `Validate() error` — ensures that all casters are properly initialized and available. It returns an error of type [\*cast.ErrCast][go-pkg-stc-cast#errcast] when the validation of any of the supported casters fails.
- `New(opts ...Option) (*Elder, error)` — creates a new elder wand.
  The module name is determined automatically using the [`runtime/debug`][go-pkg-runtime/debug] package. The absolute path to the root directory is automatically set based on the current working directory. Note that the working directory must be set manually when the “magefile“ is not placed in the root directory by pointing Mage to it:
  - `-d <PATH>` option to set the directory from which “magefiles“ are read (defaults to `.`).
  - `-w <PATH>` option to set the working directory where “magefiles“ will run (defaults to value of `-d` flag).
    If any error occurs it will be of type [\*cast.ErrCast][go-pkg-stc-cast#errcast] or [\*project.ErrProject][go-pkg-stc-project#errproject].

It is customizable through the following functions:

- `WithGobinCasterOptions(opts ...castGobin.Option) elder.Option` — sets [“gobin“ caster][go-pkg-cast/gobin] options.
- `WithGoToolchainCasterOptions(opts ...castGoToolchain.Option) elder.Option` — sets [Go toolchain caster][go-pkg-cast/golang/toolchain] options.
- `WithNib(n nib.Nib) elder.Option` — sets the [log-level based line printer for human-facing messages][go-pkg-github.com/svengreb/nib].
- `WithProjectOptions(opts ...project.Option) elder.Option` — sets [project][go-pkg-project] options.

</details>

<details>
<summary><strong>Initial project documentation</strong> — #43 ⇄ #44 (⊶ c953c4b1)</summary>

↠ The initial project documentation includes…

1. …an overview of the project features.
2. …information about the project motivation:
   1. “Why should I use [Mage][]…“
   2. “…and why _wand_?“
3. …the project design decisions and how to use it:
   1. The overall wording and inspiration.
   2. A basic overview of the API.
   3. An introduction to the “elder“ reference implementation.
4. …information about how to contribute to this project.

</details>

<p align="center">Copyright &copy; 2019-present <a href="https://www.svengreb.de" target="_blank">Sven Greb</a></p>

<p align="center"><a href="https://github.com/svengreb/wand/blob/main/LICENSE"><img src="https://img.shields.io/static/v1.svg?style=flat-square&label=License&message=MIT&logoColor=eceff4&logo=github&colorA=4c566a&colorB=88c0d0"/></a></p>

<!--
+------------------+
+ Formatting Notes +
+------------------+

The `<summary />` tag must be separated with a blank line from the actual item content paragraph,
otherwise Markdown elements are not parsed and rendered!

+------------------+
+ Symbol Reference +
+------------------+
↠ (U+21A0): Start of a log section description
— (U+2014): Separator between a log section title and the metadata
⇄ (U+21C4): Separator between a issue ID and pull request ID in a log metadata
⊶ (U+22B6): Icon prefix for the short commit SHA checksum in a log metadata
⇅ (U+21C5): Icon prefix for the link of the Git commit history comparison on GitHub
-->

<!--lint disable final-definition-->

<!-- Base Links -->

[go-ref-mod]: https://golang.org/ref/mod
[mage]: https://magefile.org
[trunkbasedev-monorepos]: https://trunkbaseddevelopment.com/monorepos
[wikip-hp]: https://en.wikipedia.org/wiki/Harry_Potter

<!-- v0.1.0 -->

[gh-compare-tag-init_v0.1.0]: https://github.com/svengreb/wand/compare/dbf11bc0...v0.1.0
[gh-golang/go-wiki]: https://github.com/golang/go/wiki
[gh-golang/go#25922]: https://github.com/golang/go/issues/25922
[gh-golang/go#27653]: https://github.com/golang/go/issues/27653
[gh-golang/go#30515]: https://github.com/golang/go/issues/30515
[gh-golang/go#40276]: https://github.com/golang/go/issues/40276
[gh-golang/tools-tree-cmd/goimports]: https://github.com/golang/tools/tree/master/cmd/goimports
[gh-golangci/golangci-lint-tree-cmd/golangci-lint]: https://github.com/golangci/golangci-lint/tree/master/cmd/golangci-lint
[gh-golangci/golangci-lint]: https://github.com/golangci/golangci-lint
[gh-mitchellh/gox]: https://github.com/mitchellh/gox
[gh-ms-golang/go#145]: https://github.com/golang/go/milestone/145
[gh-myitcv/gobin-wiki-faq]: https://github.com/myitcv/gobin/wiki/FAQ
[gh-myitcv/gobin]: https://github.com/myitcv/gobin
[gh-svengreb/tmpl-go-rl-v0.3.0]: https://github.com/svengreb/tmpl-go/releases/tag/v0.3.0
[gh-svengreb/tmpl-go-tree-apps]: https://github.com/svengreb/tmpl-go/tree/main/apps
[gh-svengreb/tmpl-go]: https://github.com/svengreb/tmpl-go
[gh-svengreb/wand#11]: https://github.com/svengreb/wand/issues/11
[gh-svengreb/wand#15]: https://github.com/svengreb/wand/issues/15
[gh-svengreb/wand#9]: https://github.com/svengreb/wand/issues/9
[git-book-intro-vcs]: https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control
[git]: https://git-scm.com
[go-docs-cmd-go#mod_aware_cmds]: https://golang.org/ref/mod#mod-commands
[go-docs-tip-rln-1.16#mod]: https://tip.golang.org/doc/go1.16#modules
[go-pkg-cast]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast
[go-pkg-cast/gobin]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/gobin
[go-pkg-cast/golang/toolchain]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/golang/toolchain
[go-pkg-cmd/go]: https://pkg.go.dev/cmd/go
[go-pkg-cmd/go#build]: https://pkg.go.dev/cmd/go/#hdr-Compile_packages_and_dependencies
[go-pkg-cmd/go#env_vars]: https://pkg.go.dev/cmd/go/#hdr-Environment_variables
[go-pkg-cmd/go#get]: https://pkg.go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them
[go-pkg-cmd/go#install]: https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
[go-pkg-cmd/go#print_env]: https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
[go-pkg-cmd/go#test]: https://pkg.go.dev/cmd/go/#hdr-Test_packages
[go-pkg-cmd/gofmt]: https://pkg.go.dev/cmd/gofmt
[go-pkg-const-spell#kindgomodule]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#KindGoModule
[go-pkg-elder]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder
[go-pkg-fn-cast#validate]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#Validate
[go-pkg-fn-spell/golang#compileformula]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang#CompileFormula
[go-pkg-github.com/golangci/golangci-lint]: https://pkg.go.dev/github.com/golangci/golangci-lint
[go-pkg-github.com/golangci/golangci-lint/cmd/golangci-lint]: https://pkg.go.dev/github.com/golangci/golangci-lint/cmd/golangci-lint
[go-pkg-github.com/mitchellh/gox]: https://pkg.go.dev/github.com/mitchellh/gox
[go-pkg-github.com/svengreb/nib]: https://pkg.go.dev/github.com/svengreb/nib
[go-pkg-golang.org/x/tools/cmd/goimports]: https://pkg.go.dev/golang.org/x/tools/cmd/goimports
[go-pkg-if-cast#binarycaster]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#BinaryCaster
[go-pkg-if-cast#caster]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#Caster
[go-pkg-if-spell#gocode]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#GoCode
[go-pkg-if-spell#incantation]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Incantation
[go-pkg-if-spell#mixin]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Mixin
[go-pkg-if#wand]: https://pkg.go.dev/github.com/svengreb/wand#Wand
[go-pkg-os#cachedir]: https://pkg.go.dev/os/#UserCacheDir
[go-pkg-project]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project
[go-pkg-runtime/debug]: https://pkg.go.dev/runtime/debug
[go-pkg-spell]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell
[go-pkg-spell/fs/clean]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/fs/clean
[go-pkg-spell/goimports]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/goimports
[go-pkg-spell/golang]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang
[go-pkg-spell/golang/build]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang/build
[go-pkg-spell/golang/test]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang/test
[go-pkg-spell/golangcilint]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golangcilint
[go-pkg-spell/gox]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/gox
[go-pkg-spell#incantation]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Incantation
[go-pkg-stc-app#errapp]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#ErrApp
[go-pkg-stc-cast/gobin#caster]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/gobin#Caster
[go-pkg-stc-cast/golang#caster]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/golang#Caster
[go-pkg-stc-cast#errcast]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#ErrCast
[go-pkg-stc-github.com/svengreb/nib#verbosity]: https://pkg.go.dev/github.com/svengreb/nib#Verbosity
[go-pkg-stc-project#errproject]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#ErrProject
[go-pkg-stc-project#gomoduleid]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#GoModuleID
[go-pkg-stc-spell#errgocode]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#ErrGoCode
[go-pkg-stc-spell#kindgomodule]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#KindGoModule
[go-ref-mod#file]: https://golang.org/ref/mod#go-mod-file
[go-wiki-tool_dep]: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
[golangci-lint]: https://golangci-lint.run
[wikip-path_var]: https://en.wikipedia.org/wiki/PATH_(variable)

<!-- v0.2.0 -->

[gh-compare-tag-v0.1.0_v0.2.0]: https://github.com/svengreb/wand/compare/v0.1.0...v0.2.0
[go-pkg-al-task#kind]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Kind
[go-pkg-if-spell#binary]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Binary
[go-pkg-if-spell#gomodule]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#GoModule
[go-pkg-if-task#exec]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Exec
[go-pkg-if-task#gomodule]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#GoModule
[go-pkg-if-task#options]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Options
[go-pkg-if-task#runner]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Runner
[go-pkg-if-task#runnerexec]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#RunnerExec
[go-pkg-if-task#task]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Task
[mage-docs-targets]: https://magefile.org/targets
[wikip-cli#anaton]: https://en.wikipedia.org/wiki/Command-line_interface#Anatomy_of_a_shell_CLI
[wikip-exec]: https://en.wikipedia.org/wiki/Executable
