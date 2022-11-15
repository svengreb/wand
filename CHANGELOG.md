<p align="center"><img src="https://raw.githubusercontent.com/svengreb/wand/main/assets/images/repository-hero.svg?sanitize=true"/></p>

<p align="center">Changelog of <em>wand</em>, a simple and powerful toolkit for <a href="https://magefile.org" target="_blank" rel="noreferrer">Mage</a>.</p>

<!--lint disable no-duplicate-headings no-duplicate-headings-in-section-->

# 0.8.0

![Release Date: 2022-11-15](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2022-11-15&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.8.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/users/svengreb/projects/6/views/3) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.6.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/9)

⇅ [Show all commits][187]

## Improvements

<details>
<summary><strong>Improve <code>runtime/debug</code> Go 1.18 incompatibility via stable <code>go.mod</code> file parsing</strong> — #129 ⇄ #130 (⊶ df291299)</summary>

↠ [As of Go 1.18][192] the [debug.ReadBuildInfo][193] function does not work for Mage executables anymore because the way how module information is stored changed. Therefore the fields of the returned [debug.Module][194] type only has zero values, including the module path. The [debug.Module.Version][195] field has a [default value][196] (`(devel)`) which is not Semver compatible and causes the parsing to fail. [The change in Go 1.18][200] also [came with the new `debug/buildinfo` package][201] which allows to read the information from compiled binaries while the `runtime/debug.ReadBuildInfo` function returns information from within the running binary. Both are not suitable anymore which is also described in the Go 1.18 `version` command release notes:

> The underlying data format of the embedded build information can change with new `go` releases, so an older version of `go` may not handle the build information produced with a newer version of `go`. To read the version information from a binary built with `go` 1.18, use the `go` version command and the `debug/buildinfo` package from `go` 1.18+.

To get the required module information that was previously provided by the [runtime/debug][197] package the official [golang.org/x/mod/modfile][198] package is now used instead that provides the implementation for a parser and formatter for [`go.mod` files][199] [^1]. This allows to safely get the module path without the need to depend on runtime/dynamic logic that might change in future Go versions.

Note that **this change also increased the minimum Go version from `1.17` to `1.19`!**

</details>

## Bug Fixes

<details>
<summary><strong>Update to <code>tmpl-go</code> template repository version <code>0.11.0</code> and <code>0.12.0</code></strong> — #112, #127 ⇄ #113, #128 (⊶ a4e2a38f, c4fe6cfc)</summary>

↠ Updated to [`tmpl-go` version `0.11.0`][203] and [`0.12.0`][204] which…

1. [fixed `golangci-lint` running errors due to `revive`s unknown `time-equal` rule][189].
2. [disabled the `revive` linter rule `package-comments`][190].
3. [updated to the `tmpl` template repository version `0.11.0`][202].

See the [full `tmpl-go` version `0.11.0`][203] and [`0.12.0`][204] and changelogs for all details.

</details>

# 0.7.0

![Release Date: 2021-11-21](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2021-11-21&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.7.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/11) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.6.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/8)

⇅ [Show all commits][157]

## Improvements

<details>
<summary><strong>Update to <code>tmpl-go</code> template repository version <code>0.9.0</code></strong> — #104 ⇄ #105 (⊶ 9caf10f9)</summary>

↠ Updated to [`tmpl-go` version `0.9.0`][164] which…

1. [updated to `golangci-lint` version `1.43.0`][165] — new linters are introduced and configurations of already supported ones are improved or added.
2. [updated the Go module to Go `1.17`][166].
3. [optimized the GitHub action workflows for Go and Node][167] — the `ci` workflow has been optimized by splitting it into new `ci-go` and `ci-node` workflows.
4. [updated to the `tmpl` template repository version `0.10.0`][168].

See the [full `tmpl-go` version `0.9.0` changelog][164] for all details.

</details>

<details>
<summary><strong>Upgrade default <code>GoModule</code> task versions</strong> — #106 ⇄ #107 (⊶ cabd635c)</summary>

↠ Most of the [`GoModule` tasks][169] used an outdated default Go module version so the following tasks have been updated and adjusted to the currently latest versions:

1. [mvdan.cc/gofumpt][170] — The [`github.com/svengreb/wand/pkg/task/gofumpt` task][171] used version `v0.1.1` and has been updated to [version `0.2.0`][170] by…
   1.1 removing the `-r` flag which has been removed in favor of `gofmt -r`.
   1.2 removing the `-s` flag ([`WithSimplify` option][173]) as it is always enabled.
2. [golang.org/x/tools/cmd/goimports][58] — The [`github.com/svengreb/wand/pkg/task/goimports` task][175] used version `v0.1.0` and has been updated to [version `0.1.7`][176].
3. [github.com/golangci/golangci-lint/cmd/golangci-lint][177] — The [`github.com/svengreb/wand/pkg/task/golangcilint` task][178] used version `v1.39.0` and has been updated to [version `1.43.0`][179]. The configuration has already been updated in #104.

</details>

<details>
<summary><strong>Update to <code>tmpl-go</code> template repository version <code>0.10.0</code></strong> — #110 ⇄ #111 (⊶ ee52f086)</summary>

↠ Updated to [`tmpl-go` version `0.10.0`][184] which…

1. [disables `golangci-lint`'s default excluded issues][185] — this prevents that explicitly enabled rules are not ignored due to the default set of excluded issues.
2. [caches Go dependencies and build outputs in `ci-go` workflow][186] — this improves the workflow execution time.

See the [full `tmpl-go` version `0.10.0` changelog][184] for all details.

</details>

## Bug Fixes

<details>
<summary><strong>Insufficient repository fetch-depth for action workflows</strong> — #108 ⇄ #109 (⊶ c39b2c42)</summary>

↠ The [GitHub action workflows][180] uses the [`actions/checkout` action][181] to fetch the repository that triggered the workflow. However, by default only the history of the latest commit was fetched which resulted in errors when _wand_ tried to extract repository metadata information like the amount of commits ahead of the latest commit. As an example this can be seen when [running the `bootstrap` command in the `test` job of the `ci-go` workflow][182] which failed with an `object not found` error because the history only contained a single commit.

To fix this problem `action/checkout` provides an option to [fetch all history for all tags and branches][183] which is now used to prevent errors like this in the pipeline.

</details>

## Tasks

<details>
<summary><strong>Go module dependency & GitHub action version updates</strong> — #97, #98, #102, #103</summary>

↠ Bumped outdated Go module dependencies and GitHub actions to their latest versions:

- #97 (⊶ 03ab1043) [`github.com/fatih/color`][158] from [v1.10.0 to v1.11.0][159]
- #98 (⊶ 7b2ac860) [`github.com/fatih/color`][158] from [v1.11.0 to v1.12.0][160]
- #102 (⊶ eecbc520) [`github.com/fatih/color`][158] from [v1.12.0 to v1.13.0][161]
- #103 (⊶ b8b906c4) [`actions/setup-node`][115] from [v2.1.5 to v2.4.1][163]

</details>

# 0.6.0

![Release Date: 2021-04-29](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2021-04-29&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.6.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/10) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.6.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/7)

⇅ [Show all commits][133]

## Features

<details>
<summary><strong>Expose task name via <code>Task</code> interface</strong> — #79, #87 ⇄ #80, #88 (⊶ bd158245, 8b30110e)</summary>

↠ Most tasks provided a `TaskName` package constant that contained the name of the task, but this was not an idiomatic and consistent way. To make sure that this information is part of the API, the new `Name() string` method has been added to the [`Task` interface][18].

</details>

<details>
<summary><strong>Task for Go toolchain <code>env</code> command</strong> — #81 ⇄ #82 (⊶ 5e3764a3)</summary>

↠ To support the [`go env` command of the Go toolchain][9], a new [`Task`][18] has been implemented in the new [`env`][144] package that can be used through a [Go toolchain `Runner`][14].
The task is customizable through the following functions:

- `WithEnv(env map[string]string) env.Option` — sets the task specific environment.
- `WithEnvVars(envVars ...string) env.Option` — sets the names of the target environment variables.
- `WithExtraArgs(extraArgs ...string) env.Option` — sets additional arguments to pass to the command.

</details>

<details>
<summary><strong><code>RunOut</code> method for <code>Runner</code> interface</strong> — #83 ⇄ #84 (⊶ d8180656)</summary>

↠ The `Run` method of the [`Runner` interface][153] allows to run a command, but did not return its output. This was blocking when running commands like `go env GOBIN` to [get the path to the `GOBIN` environment variable][141].
To support such uses cases, the new `RunOut(Task) (string, error)` method has been added to the `Runner` interface that runs a command and returns its output.

</details>

<details>
<summary><strong>Replace deprecated <code>gobin</code> with custom <code>go install</code> based task runner</strong> — #89 ⇄ #90 (⊶ 9c510a7c)</summary>

↠ This feature supersedes #78 which documents how the [official deprecation][136] of [`gobin`][8] in favor of the new Go 1.16 [`go install pkg@version`][9] syntax feature should have been handled for this project. The idea was to replace the [`gobin` task runner][148] with a one that leverages [bingo][132], a project similar to `gobin`, that comes with many great features and also allows to manage development tools on a per-module basis. The problem is that `bingo` uses some non-default and nontransparent mechanisms under the hood and automatically generates files in the repository without the option to disable this behavior. It does not make use of the `go install` command but relies on custom dependency resolution mechanisms, making it prone to future changes in the Go toolchain and therefore not a good choice for the maintainability of projects.

### `go install` is still not perfect

Support for the new `go install` features, which allow to install commands without affecting the `main` module, have already been added in #71 as an alternative to `gobin`, but one significant problem was still not addressed: install module/package executables globally without overriding already installed executables of different versions.
Since `go install` will always place compiled binaries in the path defined by `go env GOBIN`, any already existing executable with the same name will be replaced. It is not possible to install a module command with two different versions since `go install` still messes up the local user environment.

### The Workaround: Hybrid `go install` task runner

The solution was to implement a custom [`Runner`][17] that uses `go install` under the hood, but places the compiled executable in a custom cache directory instead of `go env GOBIN`. The runner checks if the executable already exists, installs it if not so, and executes it afterwards.

The concept of storing dependencies locally on a per-project basis is well-known from the [`node_modules` directory][155] of the [Node][4] package manager [npm][3]. Storing executables in a cache directory within the repository (not tracked by Git) allows to use `go install` mechanisms while not affect the global user environment and executables stored in `go env GOBIN`. The runner achieves this by changing the `GOBIN` environment variable to the custom cache directory during the execution of `go install`. This way it bypasses the need for “dirty hacks“ while using a custom output path.

The only known disadvantage is the increased usage of storage disk space, but since most Go executables are small in size anyway, this is perfectly acceptable compared to the clearly outweighing advantages.

Note that the runner dynamically runs executables based on the given task so `Validate() error` is a _NOOP_.

### Upcoming Changes

The solution described above works totally fine, but is still not a clean solution that uses the Go toolchain without any special logic so as soon as the following changes are made to the Go toolchain (Go 1.17 or later), the custom runner will be removed again:

- [golang/go/issues#42088][134] — tracks the process of adding support for the Go module syntax to the `go run` command. This will allow to let the Go toolchain handle the way how compiled executable are stored, located and executed.
- [golang/go#44469][135] — tracks the process of making `go install` aware of the `-o` flag like the `go build` command which is the only reason why the custom runner has been implemented.

### Further Adjustments

Because the new custom task runner dynamically runs executables based on the given task, the [`Bootstrap` method][147] of the [`Wand`][19] reference implementation [`Elder`][143] now additionally allows to pass Go module import paths, optionally including a version suffix (`pkg@version`), to install executables from Go module-based `main` packages into the local cache directory. This way the local development environment can be set up, for e.g. by running it as [startup task][154] in _JetBrains_ IDEs.
The method also ensures that the local cache directory exists and will create a `.gitignore` file that includes ignore pattern for the cache directory.

</details>

<details>
<summary><strong>Task for <code>go-mod-upgrade</code> Go module command</strong> — #95 ⇄ #96 (⊶ c944173f)</summary>

↠ The [github.com/oligot/go-mod-upgrade][137] Go module provides the `go-mod-upgrade` command, a tool that to update outdated Go module dependencies interactively.

To configure and run the `go-mod-upgrade` command, a new [`task.GoModule`][15] has been implemented in the new [`gomodupgrade`][145] package. It can be be run using a [command runner][17] that handles tasks of kind [`KindGoModule`][16].

The task is customizable through the following functions:

- `WithEnv(map[string]string) gomodupgrade.Option` — sets the task specific environment.
- `WithExtraArgs(...string) gomodupgrade.Option` — sets additional arguments to pass to the command.
- `WithModulePath(string) gomodupgrade.Option` — sets the module import path.
- `WithModuleVersion(*semver.Version) gomodupgrade.Option` — sets the module version.

The [`Elder`][11] reference implementation will provide a new [`GoModUpgrade` method][142].

</details>

## Improvements

<details>
<summary><strong>Remove unnecessary <code>Wand</code> parameter in <code>Task</code> creation functions</strong> — #76 ⇄ #77 (⊶ 536556b6)</summary>

↠ Most `Task` creation functions [<sup>1</sup>][149] [<sup>2</sup>][150] [<sup>3</sup>][151] [<sup>4</sup>][152] required a `Wand` as parameter which was not used but blocked the internal usage for task runners. Therefore these parameters have been removed. When necessary, it can be added individually later on or can be reintroduced through a dedicated function with extended parameters to cover different use cases.

</details>

<details>
<summary><strong>Remove unnecessary <code>app.Config</code> parameter from <code>Task</code> creation functions</strong> — #85 ⇄ #86 (⊶ 72dd6a1a)</summary>

↠ Some functions that create a [`Task`][18] required an [`app.Config` struct][146], but most tasks did not use the data in any way. To improve the code quality and simplify the internal usage of tasks these parameters have been removed as well as the field from the structs that implement the `Task` interfaces.

</details>

<details>
<summary><strong>Update to <code>tmpl-go</code> template repository version <code>0.8.0</code></strong> — #91 ⇄ #92 (⊶ 3e189171)</summary>

↠ Updated to [`tmpl-go` version `0.8.0`][138] which [updates `golangci-lint` to version `1.39.0`][139] and [the `tmpl` repository version `0.9.0`][140].

</details>

<details>
<summary><strong>Dogfooding: Introduce Mage with wand toolkit</strong> — #93 ⇄ #94 (⊶ 85c466d7)</summary>

↠ The project only used _GitHub Action_ workflows for CI but not _Mage_ to automate tasks for itself though.
Following the [“dogfooding“ concept][156] _Mage_ has finally been added to the repository, using wand as toolkit through the [`Elder` wand reference][143] implementation.

</details>

# 0.5.0

![Release Date: 2021-04-22](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2021-04-22&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.5.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/9) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.5.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/6)

⇅ [Show all commits][123]

This release comes with support for Go 1.16 features like the new `install` command behavior and removes the now unnecessary `pkger` task runner in favor of the new `embed` package and `//go:embed` directive.

## Features

<details>
<summary><strong>Task for Go toolchain <code>install</code> command</strong> — #70 ⇄ #71 (⊶ c36e8f31)</summary>

↠ As of Go version 1.16 [`go install $pkg@$version`][125] allows to install commands without affecting the `main` module. Additionally commands like `go build` and `go test` no longer modify `go.mod` and `go.sum` files by default but report an error if a module requirement or checksum needs to be added or updated (as if the `-mod=readonly` flag were used).
This can be used as alternative to the already existing [`gobin` runner][129].

To support the [`go install` command of the Go toolchain][9], a new [`Task`][18] has been implemented in the new [`install`][131] package that can be used through a [Go toolchain `Runner`][14].
The task is customizable through the following functions:

- `WithEnv(env map[string]string) install.Option` — sets the task specific environment.
- `WithModulePath(path string) install.Option` — sets the module import path.
- `WithModuleVersion(version *semver.Version) install.Option` — sets the module version.

</details>

## Tasks

<details>
<summary><strong>Updated to "tmpl-go" template repository version <code>0.7.0</code></strong> — #72 ⇄ #73 (⊶ 53fd75ec)</summary>

↠ Updated to ["tmpl-go" version 0.7.0][124] which comes with updates to GitHub Actions and Node development dependencies.

</details>

<details>
<summary><strong>Removed <code>pkger</code> task in favor of Go 1.16 <code>embed</code> package</strong> — #74 ⇄ #75 (⊶ 1fc1f253)</summary>

↠ In #52 a task for the [github.com/markbates/pkger][12] Go module was added, a tool for embedding static files into Go binaries.
The issue also includes the “Official Static Assets Embedding“ section which mentions that the task might be removed later on again as soon as [Go 1.16][126] will be released as it comes with [toolchain support for embedding static assets (files)][127] through the [`embed` package][128]. Also see [markbates/pkger#114][7] for more details about the project future of `pkger`.

The [`pkger` package][130] has been removed and the `//go:embed` directive should be used instead.

</details>

# 0.4.1

![Release Date: 2021-01-04](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2021-01-04&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.4.1&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/8) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.4.1&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/5)

⇅ [Show all commits][116]

This release version fixes a bug that could occur when running the `Install` method of the `gobin` task runner in minimal environments like containers.

## Bug Fixes

<details>
<summary><strong>Fix missing environment variables in <code>Install</code> method of <code>gobin</code> task</strong> — #63 ⇄ #62 (⊶ ff54e917)</summary>

↠ Fixed possible errors like

```raw
build cache is required, but could not be located: GOCACHE is not defined and neither $XDG_CACHE_HOME nor $HOME are defined
```

when running the method in minimal environments like containers by ensuring that the inherited OS environment is prepended before applying custom environment variables.

Before the [`Install` method of the `gobin` task runner][121] has set the environment of the command that gets executed initially to [`os.Environ()`][118], but has overwritten it later on with custom variables configured through the [`WithEnv(map[string]string)` option][120].

This change also improves the debugging process by including the combined output (`stdout` + `stderr`) in the error when the command execution fails.

</details>

## Tasks

<details>
<summary><strong>Go module dependency & GitHub action version updates</strong> — #60, #61</summary>

↠ Bumped outdated Go module dependencies and GitHub actions to their latest versions:

- #60 (⊶ 3fd3f8b4) [`actions/setup-node`][115] from [v2.1.3 to v2.1.4][114]
- #61 (⊶ 6dd713e5) [`github.com/magefile/mage`][119] from [v1.10.0 to v1.11.0][117] - This release finally introduces a long-time requested feature: [Target functions with arguments][122]!
  This allows to pass parameters to targets from the CLI to make functions even more dynamic.

</details>

# 0.4.0

![Release Date: 2020-12-11](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-12-11&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.4.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/7) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.4.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/4)

⇅ [Show all commits][109]

This release version introduces a new task for the “mvdan.cc/gofumpt“ Go module command.

## Features

<details>
<summary><strong>Task for “mvdan.cc/gofumpt“ Go module command</strong> — #56 ⇄ #57 (⊶ 3273e91f)</summary>

↠ The [mvdan.cc/gofumpt][112] Go module provides the `gofumpt` command, a tool that enforces a stricter format than [`gofmt`][10] and [provides additional rules][110], while being backwards compatible. It is a modified fork of `gofmt` so it can be used as a drop-in replacement.

To configure and run the `gofumpt` command, a new [`task.GoModule`][15] has been implemented in the new [gofumpt][113] package that can be run using the [gobin command runner][13] or any other [command runner][17] that handles tasks of kind [`KindGoModule`][16].

The task is customizable through the following functions:

- `WithEnv(map[string]string) gofumpt.Option` — sets the task specific environment.
- `WithExtraArgs(...string) gofumpt.Option` — sets additional arguments to pass to the command.
- `WithExtraRules(bool) gofumpt.Option` — indicates whether `gofumpt`‘s extra rules should be enabled. See the [repository documentation for a listing of available rules][110].
- `WithListNonCompliantFiles(bool) gofumpt.Option` — indicates whether files, whose formatting are not conform to the style guide, are listed.
- `WithModulePath(string) gofumpt.Option` — sets the module import path.
- `WithModuleVersion(*semver.Version) gofumpt.Option` — sets the module version.
- `WithPaths(...string) gofumpt.Option` — sets the paths to search for Go source files. By default all directories are scanned recursively starting from the current working directory.
- `WithReportAllErrors(bool) gofumpt.Option` — indicates whether all errors should be printed instead of only the first 10 on different lines.
- `WithSimplify(bool) gofumpt.Option` — indicates whether code should be simplified.

The “elder“ reference implementation provides the new [`Gofumpt` method][111].

</details>

# 0.3.0

![Release Date: 2020-12-10](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-12-10&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.3.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/6) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.3.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/3)

⇅ [Show all commits][97]

This release version introduces a new task for the “github.com/markbates/pkger“ Go module command and updates for outdated dependencies.

## Features

<details>
<summary><strong>Task for “github.com/markbates/pkger“ Go module command</strong> — #52 ⇄ #53 (⊶ 660601dd)</summary>

↠ The [github.com/markbates/pkger][12] Go module provides the `pkger` command, a tool for embedding static files into Go binaries.

To configure and run the `pkger` command, a new [`task.GoModule`][15] has been implemented in a the [pkger][107] package that can be run using the [gobin command runner][13] or any other [command runner][17] that handles tasks of kind [`KindGoModule`][16].

The task is customizable through the following functions:

- `WithEnv(env map[string]string) pkger.Option` — sets the task specific environment.
- `WithExtraArgs(extraArgs ...string) pkger.Option` — sets additional arguments to pass to the command.
- `WithIncludes(includes ...string) pkger.Option` — adds the relative paths of files and directories that should be included.
  By default the paths will be detected by `pkger` itself when used within any of the packages of the target Go module.
- `WithModulePath(path string) pkger.Option` — sets the module import path.
- `WithModuleVersion(version *semver.Version) pkger.Option` — sets the module version.

The “elder“ reference implementation provides the new [`Pkger` method][104] including the handling of the [“monorepo“ workaround](#monorepo-workaround).

### Official “Static Assets Embedding“

Please note that the _pkger_ project might be superseded and discontinued due to the official Go toolchain [support for embedding static assets (files)][98] that will most probably be released with [Go version 1.16][6].

Please see the official [draft document][108] and [markbates/pkger#114][7] for more details.

### “Monorepo“ Workaround

_pkger_ tries to mimic the Go standard library and the way how the Go toolchain handles modules, but is therefore also affected by its problems and edge cases.
When the `pkger` command is used from the root of a Go module repository, the directory where the `go.mod` file is located, and there is no valid Go source file, the command will fail because it internally uses the same logic like the [`list` command of the Go toolchain][103] (`go list`).
Therefore a “dummy“ Go source file may need to be created as a workaround. This is mostly only required for repositories that use a [“monorepo“ layout][21] where one or more `main` packages are placed in a subdirectory relative to the root directory, e.g. `apps` or `cmd`. For repositories where the root directory already has a Go package, that does not contain any build constraints/tags, or uses a “library“ layout, a “dummy“ file is probably not needed.
Please see [markbates/pkger#109][100] and [markbates/pkger#121][101] for more details.

The new [`Pkger` method][104] of the [“elder“ reference implementation][11] handles the creation of a temporary “dummy“ file that gets deleted automatically when the tasks finishes in order to avoid the need for the user to add such a file to the repository and commit it into the VCS.

</details>

<details>
<summary><strong>Update outdated dependencies</strong> — #47, #48</summary>

↠ Bumped outdated Go module dependencies to their latest versions:

- #47 (⊶ 41e11b94) [`github.com/Masterminds/semver/v3`][106] from [3.1.0 to 3.1.1][102] — Fixes an issue with generated regular expression operations.
- #48 (⊶ 41e11b94) [`github.com/imdario/mergo`][105] from [0.3.9 to 0.3.11][99] — Includes a bunch of bug fixes that were pending, removes unused test code, reverts a faulty PR and announces a code freeze in preparation for a “cleanroom“ implementation with a new API in order to allow the codebase to be maintainable and clear again.

</details>

# 0.2.0

![Release Date: 2020-12-07](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-12-07&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.2.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/5) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.2.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/2)

⇅ [Show all commits][87]

This release version comes with a large API breaking change to introduce the new “task“ + “runner“ based API that uses a “normalized“ naming scheme.

## Features

<details>
<summary><strong>“Task“ API: Simplified usage and “normalized“ naming scheme</strong> — #49 ⇄ #51 (⊶ f51a4bfa)</summary>

↠ With #14 the “abstract“ _wand_ API was introduced with a naming scheme is inspired by the fantasy novel [“Harry Potter“][22] that was used to to define interfaces.
The main motivation was to create a matching naming to the overall “magic“ topic and the actual target project [Mage][1], but in retrospect this is way too abstract and confusing.

The goal of this change was to…

- rewrite the API to **make it way easier to use**.
- use a **“normal“ naming scheme**.
- improve all **documentations to be more user-scoped** and provide **guides and examples**.

#### New API Concept

The basic mindset of the API will remain partially the same, but it will be designed around the concept of **tasks** and the ways to **run** them.

##### Command Runner

[🅸 `task.Runner`][17] is a new base interface that runs a command with parameters in a specific environment. It can be compared to the previous [🅸 `cast.Caster`][60] interface, but provides a cleaner method set accepting the new [🅸 `task.Task`][18] interface.

- 🅼 `Handles() task.Kind` — returns the supported [task kind][88].
- 🅼 `Run(task.Task) error` — runs a command.
- 🅼 `Validate() error` — validates the runner.

The new [🅸 `task.RunnerExec`][93] interface is a specialized `task.Runner` and serves as an abstract representation for a command or action, in most cases a (binary) [executable][96] of external commands or Go module `main` packages, that provides corresponding information like the path to the executable. It can be compared to the previous [`BinaryCaster`][59] interface, but also comes with a cleaner method set and a more appropriate name.

- 🅼 `FilePath() string` — returns the path to the (binary) command executable.

##### Tasks

[🅸 `task.Task`][18] is the new interface that is scoped for Mage [“target“][94] usage. It can be compared to the previous [🅸 `spell.Incantation`][62] interface, but provides a smaller method set without `Formula() []string`.

- 🅼 `Kind() task.Kind` — returns the [task kind][88].
- 🅼 `Options() task.Options` — returns the [task options][92].

The new [🅸 `task.Exec`][91] interface is a specialized `task.Task` and serves as an abstract task for an executable command. It can be compared to the previous [`Binary`][89] interface, but also comes with the new `BuildParams() []string` method that enables a more flexible usage by exposing the parameters for command runner like `task.RunnerExec` and also allows to compose with other tasks. See the Wikipedia page about [the anatomy of a shell CLI][95] for more details about parameters.

- 🅼 `BuildParams() []string` — builds the parameters for a command runner where parameters can consist of options, flags and arguments.
- 🅼 `Env() map[string]string` — returns the task specific environment.

The new [🅸 `task.GoModule`][15] interface is a specialized `task.Exec` for a executable Go module command. It can be compared to the previous [`spell.GoModule`][90] interface and the method set has not changed except a renaming of the `GoModuleID() *project.GoModuleID` to the more appropriate name `ID() *project.GoModuleID`. See the official [Go module reference documentation][20] for more details about Go modules.

- 🅼 `ID() *project.GoModuleID` — returns the identifier of a Go module.

#### New API Naming Scheme

The following listing shows the new name concept and how the previous API components can be mapped to the changes:

1. **Runner** — A component that runs a command with parameters in a specific environment, in most cases a (binary) [executable][96] of external commands or Go module `main` packages. The current API component that can be compared to runners is [🅸 `cast.Caster`][60] and its specialized interfaces.
2. **Tasks** — A component that is scoped for Mage [“target“][94] usage in order to run a action. The current API component that can be compared to tasks is [🅸 `spell.Incantation`][62] and its specialized interfaces.

#### API Usage

Even though the API has been changed quite heavily, the basic usage almost did not change.

→ **A `task.Task` can only be run through a `task.Runner`!**

Before a `spell.Incantation` was passed to a `cast.Caster` in order to run it, in most cases a (binary) executable of a command that uses the `Formula() []string` method of `spell.Incantation` to pass the result as parameters.
The new API works the same: A `task.Task` is passed to a `task.Runner` that calls the `BuildParams() []string` method when the runner is specialized for (binary) executable of commands.

#### Improved Documentations

Before the documentation was mainly scoped on technical details, but lacked more user-friendly sections about topics like the way how to implement own API components, how to compose the [“elder“ reference implementation][11] or usage examples for single or [monorepo][21] project layouts.

##### User Guide

Most of the current sections have been rewritten or removed entirely while new sections now provide more user-friendly guides about how to…

- use or compose the [“elder“ reference implementation][11].
- build own tasks and runners using the new API.
- structure repositories independent of the layout, single or “monorepo“.

##### Usage Examples

Some examples have been added, that are linked and documented in the user guides described above, to show how to…

- use or compose the [“elder“ reference implementation][11].
- build own tasks and runners using the new API.
- structure repositories independent of the layout, single or “monorepo“.

</details>

# 0.1.0

![Release Date: 2020-11-29](https://img.shields.io/static/v1?style=flat-square&label=Release%20Date&message=2020-11-29&colorA=4c566a&colorB=88c0d0) [![Project Board](https://img.shields.io/static/v1?style=flat-square&label=Project%20Board&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/projects/4) [![Milestone](https://img.shields.io/static/v1?style=flat-square&label=Milestone&message=0.1.0&logo=github&logoColor=eceff4&colorA=4c566a&colorB=88c0d0)](https://github.com/svengreb/wand/milestone/1)

⇅ [Show all commits][23]

This is the initial release version of _wand_.
The basic project setup, structure and development workflow has been bootstrapped by [the _tmpl-go_ template repository][35].
The following sections of this version changelog summarize used technologies, explain design decisions and provide an overview of the API and “elder“ reference implementation.

## Features

<details>
<summary><strong>Bootstrap based on “tmpl-go“ template repository</strong> — #1, #2, #4, #12 ⇄ #3, #5, #13 (⊶ dbf11bc0, f1eee4a1, f778fd97, 5d417258)</summary>

<p align="center"><img src="https://github.com/svengreb/tmpl-go/blob/main/assets/images/repository-hero.svg?raw=true"/></p>

↠ Bootstrapped the basic project setup, structure and development workflow [from version 0.3.0][33] of the [“tmpl-go“ template repository][35].
Project specific files like the repository hero image, documentations and GitHub issue/PR templates have been adjusted.

</details>

<details>
<summary><strong>Application configuration store</strong> — #8 ⇄ #9 (⊶ a233575d)</summary>

↠ Like described in [the `/apps` directory documentation][34] of the _tmpl-go_ template repository, _wand_ also aims to support the [monorepo][21] layout.
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

↠ In [GH-9][38] the store and configuration for applications has been implemented. _wand_ applications are not standalone but part of a project which in turn is stored in a repository of [a VCS like Git][39]. In case of _wand_ this can also be a [monorepo][21] to manage multiple applications, but there is always only a single project which all these applications are part of.
To store project and VCS repository information, some of the newly implemented packages provide the following types:

- 🆃 `pkg/project.Metadata` — A `struct` type that stores information and metadata of a project.
- 🆃 `pkg/project.GoModuleID` — A `struct` type that stores partial information to identify a [Go module][20].
- 🆃 `pkg/vcs.Kind` — A `struct` type that defines the kind of a `pkg/vcs.Repository`.
- 🅸 `pkg/vcs.Repository` — A `interface` type to represents a VCS repository that provides methods to receive repository information:
  - `Kind() Kind` — returns the repository `pkg/vcs.Kind`.
  - `DeriveVersion() error` — derives the repository version based on the `pkg/vcs.Kind`.
  - `Version() interface{}` — returns the repository version.
- 🆃 `pkg/vcs/git.Git` — A `struct` type that implements `pkg/vcs.Repository` to represent a [Git][2] repository.
- 🆃 `pkg/vcs/git.Version` — A `struct` type that stores version information and metadata derived from a [Git][2] repository.
- 🆃 `pkg/vcs/none.None` — A `struct` type that implements `pkg/vcs.Repository` to represent a nonexistent repository.

</details>

<details>
<summary><strong>Abstract “task“ API: _spell incantation_, _kind_ and _caster_</strong> — #14 ⇄ #15 (⊶ 2b13b840)</summary>

↠ The _wand_ API is inspired by the fantasy novel [“Harry Potter“][22] and uses an abstract view to define interfaces. The main motivation to create a matching naming to the overall “magic“ topic and the actual target project [Mage][1]. This might be too abstract for some, but is kept understandable insofar as it should allow everyone to use the “task“ API and to derive their own tasks from it.

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
- 🅸 `cast.GoModule` — A `interface` type that composes `cast.Binary` for commands that are compiled from a [Go module][20]
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
  - `spell.GoModule` → a composed `spell.Binary` to run binary commands managed by a [Go module][20], in other words executables installed in `GOBIN` or received via `go get`.
    It requires the module identifier (`path@version`) in order to download and run the executable.

</details>

<details>
<summary><strong>Basic “wand“ API</strong> — #16 ⇄ #17 (⊶ cc9f7c4b)</summary>

↠ In [GH-15][37] some parts of the _wand_ API have been implemented in form of spell _incantations_, _kinds_ and _casters_, inspired by the fantasy novel [“Harry Potter“][22] as an abstract view to define interfaces. In [GH-9][38] and [GH-11][36] the API implementations for an application configuration store as well as project and VCS repository metadata were introduced.
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

↠ To use the Go toolchain, also known as [the `go` command][45], a new [caster][60] (introduced in #14) has been implemented.
The new [`ErrCast`][78] `struct` type unifies the handling of errors in the [cast][42] package.

The [`Validate` function][52] of the new caster returns an error of type `*cast.ErrCast` when the `go` binary executable does not exist at the configured path or when it is also not available in the [executable search paths][86] of the current environment.

</details>

<details>
<summary><strong>“gobin“ Go module caster</strong> — #22 ⇄ #23 (⊶ 95c22a00)</summary>

##### Go Executable Installation

When installing a Go executable from within a [Go module][20] directory using the [`go install` command][9], it is installed into the Go executable search path that is defined through [the `GOBIN` environment variable][47] and can also be shown and modified using the [`go env` command][49]. Even though the executable gets installed globally, the [`go.mod` file][83] will be updated to include the installed packages since this is the default behavior of [the `go get` command][48] when running in [_module_ mode][40].

Next to this problem, the installed executable will also overwrite any executable of the same module/package that was installed already, but maybe from a different version. Therefore only one version of a executable can be installed at a time which makes it impossible to work on different projects that use the same tool but with different versions.

##### History & Future

The local installation of executables built from Go modules/packages has always been a somewhat controversial point which unfortunately, partly for historical reasons, does not offer an optimal and user-friendly solution up to now. The [`go` command][45] is a fantastic toolchain that provides many great features one would expect to be provided out-of-the-box from a modern and well designed programming language without the requirement to use a third-party solution: from compiling code, running unit/integration/benchmark tests, quality and error analysis, debugging utilities and many more.
Unfortunately the way the [`go install` command][9] of Go versions less or equal to 1.15 handles the installation of an Go module/package executable is still not optimal.

The general problem of tool dependencies is a long-time known issue/weak point of the current Go toolchain and is a highly rated change request from the Go community with discussions like [golang/go#30515][27], [golang/go#25922][25] and [golang/go#27653][26] to improve this essential feature, but they‘ve been around for quite a long time without a solution that works without introducing breaking changes and most users and the Go team agree on.
Luckily, this topic was finally picked up for [the next upcoming Go release version 1.16][6] and [gh-golang/go#40276][5] introduces a way to install executables in module mode outside a module. The [release note preview also already includes details about this change][41] and how installation of executables from Go modules will be handled in the future.

##### The Workaround

Beside the great news and anticipation about an official solution for the problem the usage of a workaround is almost inevitable until Go 1.16 is finally released.

The [official Go wiki][24] provides a section on [“How can I track tool dependencies for a module?”][84] that describes a workaround that tracks tool dependencies. It allows to use the Go module logic by using a file like `tools.go` with a dedicated `tools` build tag that prevents the included module dependencies to be picked up included for normal executable builds. This approach works fine for non-main packages, but CLI tools that are only implemented in the `main` package can not be imported in such a file.

In order to tackle this problem, a user from the community created [gobin][8], _an experimental, module-aware command to install/run main packages_.
It allows to install or run main-package commands without “polluting“ the `go.mod` file by default. It downloads modules in version-aware mode into a binary cache path within [the systems cache directory][64].
It prevents problems due to already globally installed executables by placing each version in its own directory. The decision to use a cache directory instead of sub-directories within the `GOBIN` path keeps the system clean.

_gobin_ is still in an early development state, but has already received a lot of positive feedback and is used in many projects. There are also members of the core Go team that have contributed to the project and the chance is high that the changes for Go 1.16 were influenced or partially ported from it.
It is currently the best workaround to…

1. …prevent the Go toolchain to pick up the [`GOMOD` environment variable][49] (see [`go env GOMOD`][49]) that is initialized automatically with the path to the `go.mod` file in the current working directory.
2. …install module/package executables globally without “polluting“ the `go.mod` file.
3. …install module/package executables globally without overriding already installed executables of different versions.

See [gobin‘s FAQ page][32] in the repository wiki for more details about the project.

#### The Go Module Caster

To allow to manage the tool dependency problem, _wand_ uses `gobin` through [a new caster][76] that prevents the “pollution“ of the project `go.mod` file and allows to…

1. …install `gobin` itself into `GOBIN` ([`go env GOBIN`][49]).
2. …cast any [spell incantation][62] of kind [`KindGoModule`][51] by installing the executable globally into the dedicated `gobin` cache.

</details>

<details>
<summary><strong>Spell incantation options “mixin“</strong> — #25 ⇄ #26 (⊶ 9ae4f892)</summary>

↠ To allow to compose, manipulate and read spell incantation options after the initial creation, two new types have been added for the [spell][67] package:

- 🅸 `spell.Options` — A `interface` type as a generic representation for `spell.Incantation` options.
- 🅸 `spell.Mixin` — A `interface` type that allows to compose functions that process `spell.Options` of `spell.Incantation`s.
  - `Apply(Options) (Options, error)` — applies generic `spell.Options` to `spell.Incantation` options.

</details>

<details>
<summary><strong>Spell incantation for Go toolchain <code>build</code> command</strong> — #27 ⇄ #28 (⊶ 060b3328)</summary>

↠ To run the `go build` command of the Go toolchain, a new [`spell.Incantation`][62] has been implemented in the new [build][71] package that can be used through a [Go toolchain caster][77].
The spell incantation is configurable through the following functions:

- `WithBinaryArtifactName(name string) build.Option` — sets the name for the binary build artifact.
- `WithCrossCompileTargetPlatforms(platforms ...string) build.Option` — sets the names of cross-compile platform targets.
- `WithFlags(flags ...string) build.Option` — sets additional flags to pass to the Go `build` command along with the base Go flags.
- `WithGoOptions(goOpts ...spellGo.Option) build.Option` — sets shared Go toolchain commands options.
- `WithOutputDir(dir string) build.Option` — sets the output directory, relative to the project root, for compilation artifacts.

To unify further implementations for the Go toolchain, a new `struct` type is available in the [golang][70] package to store global/shared Go toolchain options that are shared between multiple Go toolchain commands:

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

The new [`CompileFormula(opts ...Option) []string` function][53] can be used to compile the formula for these options.

</details>

<details>
<summary><strong>Spell incantation for Go toolchain <code>test</code> command</strong> — #29 ⇄ #30 (⊶ 166a2dc0)</summary>

↠ To run the `go test` command of the Go toolchain, a new [`spell.Incantation`][62] is available in the new [test][72] package that can be used through a [Go toolchain caster][77].
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

↠ The [golang.org/x/tools/cmd/goimports][58] Go module allows to update Go import lines, adding missing ones and removing unreferenced ones. It also formats code in the same style as [gofmt][10] so it can be used as a replacement. The source code for the `goimports` command can be found in the [golang/tools][28] repository.

To configure and run the `goimports` command, a new [`spell.Incantation`][62] is available in the new [goimports][69] package that can be casted using the [gobin caster][76] or any other [spell caster][60] that handles [spell incantations][62] of kind [`KindGoModule`][51].

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

↠ The [github.com/golangci/golangci-lint][54] Go module provides the `golangci-lint` command, a fast, parallel runner for dozens of Go linters Go that uses caching, supports YAML configurations and has integrations with all major IDEs. The source code for the `golangci-lint` command can be found in the [golangci/golangci-lint][29] repository.

To configure and run the `golangci-lint` command, a new [`spell.Incantation`][62] is available in the new [golangcilint][73] package that can be casted using the [gobin caster][76] or any other [spell caster][60] that handles [spell incantations][62] of kind [`KindGoModule`][51].

The spell incantation is customizable through the following functions:

- `WithArgs(args ...string) golangcilint.Option` — sets additional arguments to pass to the `golangci-lint` module command.
- `WithEnv(env map[string]string) golangcilint.Option` — sets the spell incantation specific environment.
- `WithModulePath(path string) golangcilint.Option` — sets the `golangci-lint` module command import path. Defaults to `golangcilint.DefaultGoModulePath`.
- `WithModuleVersion(version *semver.Version) golangcilint.Option` — sets the `golangci-lint` module version. Defaults to `golangcilint.DefaultGoModuleVersion`.
- `WithVerboseOutput(verbose bool) golangcilint.Option` — indicates whether the output should be verbose.

</details>

<details>
<summary><strong>Spell incantation for the <code>github.com/mitchellh/gox</code> Go module</strong> — #35 ⇄ #36 (⊶ 4b285060)</summary>

↠ The [github.com/mitchellh/gox][56] Go module provides the `gox` command, a dead simple, no frills Go cross compile tool that behaves a lot like the standard Go toolchain `build` command.

To configure and run the `gox` command, a new [`spell.Incantation`][62] is available in the new [gox][74] package that can be casted using the [gobin caster][76] or any other [spell caster][60] that handles [spell incantations][62] of kind [`KindGoModule`][51].

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

↠ To support common use cases for debugging and production optimization, some [spell mixins][63] have been implemented in the [golang][70] package:

- 🆂 `MixinImproveDebugging` — A `struct` type that adds linker flags to improve the debugging of binary artifacts. This includes the disabling of inlining and all compiler optimizations tp improve the compatibility for debuggers.
  Note that this mixin will add the `all` prefix for `—gcflags` parameters to make sure all packages are affected. If you disabled the `all` prefix on purpose you need to handle this conflict on your own, e.g. by creating more than one binary artifact each with different build options.
- 🆂 `MixinImproveEscapeAnalysis` — A `struct` type that will add linker flags to improve the escape analysis of binary artifacts.
  Note that this mixin removes the `all` prefix for `—gcflags` parameters to make sure only the target package is affected, otherwise reports for (traverse) dependencies would be included as well. If you enabled the `all` prefix on purpose you need to handle this conflict on your own, e.g. by creating more than one binary artifact each with different build options.
- 🆂 `MixinStripDebugMetadata` — A `struct` type that will add linker flags to strip debug information from binary artifacts. This will include _DWARF_ tables needed for debuggers, but keeps annotations needed for stack traces so panics are still readable. It will also shrink the file size and memory overhead as well as reducing the chance for possible security related problems due to enabled development features or debug information leaks.
  Note that this mixin will add the `all` prefix for `—gcflags` parameters to make sure all packages are affected. If you disabled the `all` prefix on purpose you need to handle this conflict on your own, e.g. by creating more than one binary artifact each with different build options.
- 🆂 `MixinInjectBuildTimeVariableValues` — A `struct` type that will inject build—time values through the `—X` linker flags to populate e.g. application metadata variables.
  It will store a `map[string]string` of key/value pairs to inject to variables at build—time. The key must be the path to the variable in form of `<IMPORT_PATH>.<VARIABLE_NAME>`, e.g. `pkg/internal/support/app.version`. The value is the actual value that will be assigned to the variable, e.g. the application version.
  A field of type [`*project.GoModuleID`][81] will store partial information about the target Go module to inject the key/value pairs from the data map into.

</details>

<details>
<summary><strong>Go code spell for filesystem cleaning</strong> — #39 ⇄ #40 (⊶ 04a3aeb9)</summary>

↠ To clean paths in a filesystem, like application specific output directories, a new [`GoCode` spell incantation][61] is available in the new [clean][68] package that can be used without a [caster][60].

The spell incantation provides the following methods:

- `Clean() ([]string, error)` — removes the configured paths. It returns an error of type `*spell.ErrGoCode` for any error that occurs during the execution of the Go code.

The spell incantation is customizable through the following functions:

- `WithLimitToAppOutputDir(limitToAppOutputDir bool) clean.Option` — indicates whether only paths within the configured application output directory should be allowed.
- `WithPaths(paths ...string) clean.Option` — sets the paths to remove. Note that only paths within the configured application output directory are allowed when `WithLimitToAppOutputDir` is enabled.

</details>

<details>
<summary><strong>Wand reference implementation “elder“</strong> — #41 ⇄ #42 (⊶ 6397641b)</summary>

↠ The default way to use the [_wand_ API][19], with its [casters][42] and [spells][67], is the reference implementation [“elder“][11].
It provides a way to use all _wand_ spells and additionally comes with helper methods to bootstrap a project, validate all _casters_ and simplify logging for process exits:

- `Bootstrap() error` — runs initialization tasks to ensure the wand is operational. This includes the installation of configured caster like [`cast.BinaryCaster`][59] that can handle spell incantations of kind [`spell.KindGoModule`][51].
- `Clean(appName string, opts ...clean.Option) ([]string, error)` — a [`spell.GoCode`][61] to remove configured filesystem paths, e.g. output data like artifacts and reports from previous development, test, production and distribution builds. It returns paths that have been cleaned along with an error of type [`*spell.ErrGoCode`][82] when an error occurred during the execution of the Go code. When any error occurs it will be of type [`*app.ErrApp`][75] or [`*cast.ErrCast`][78]. See the [clean][68] package for all available options.
- `ExitPrintf(code int, verb nib.Verbosity, format string, args ...interface{})` — simplifies the logging for process exits with a suitable [`nib.Verbosity`][79].
- `GetAppConfig(name string) (app.Config, error)` — returns an application configuration. An empty application configuration is returned along with an error of type [`*app.ErrApp`][75] when there is no configuration in the store for the given name.
- `GetProjectMetadata() project.Metadata` — returns metadata of the project.
- `GoBuild(appName string, opts ...build.Option)` — casts the spell incantation for the [`build`][46] command of the [Go toolchain][45]. When any error occurs it will be of type [`*app.ErrApp`][75] or [`*cast.ErrCast`][78]. See the [build][71] package for all available options.
- `Goimports(appName string, opts ...goimports.Option) error` — casts the spell incantation for the [golang.org/x/tools/cmd/goimports][58] Go module command that allows to update Go import lines, add missing ones and remove unreferenced ones. It also formats code in the same style as [`gofmt` command][10] so it can be used as a replacement. When any error occurs it will be of type [`*app.ErrApp`][75] or [`*cast.ErrCast`][78].
  See the [goimports][69] package for all available options. For more details about `goimports` see [the module documentation][58]. The source code of `goimports` is [available in the GitHub repository][28].
- `GolangCILint(appName string, opts ...golangcilint.Option) error` — casts the spell incantation for the [github.com/golangci/golangci-lint/cmd/golangci-lint][55] Go module command, a fast, parallel runner for dozens of Go linters Go that uses caching, supports YAML configurations and has integrations with all major IDEs. When any error occurs it will be of type [`*app.ErrApp`][75] or [`*cast.ErrCast`][78]. See the [golangcilint][73] package for all available options.
  For more details about `golangci-lint` see [the module documentation][55] and the [official website][85]. The source code of `golangci-lint` is [available in the GitHub repository][30].
- `GoTest(appName string, opts ...spellGoTest.Option) error` — casts the spell incantation for the [`test`][50] command of the [Go toolchain][45]. When any error occurs it will be of type [`*app.ErrApp`][75] or [`*cast.ErrCast`][78]. See the [test][72] package for all available options.
- `Gox(appName string, opts ...spellGox.Option) error` — casts the spell incantation for the [github.com/mitchellh/gox][56] Go module command, a dead simple, no frills Go cross compile tool that behaves a lot like the standard Go toolchain [`build`][46] command. When any error occurs it will be of type [`*app.ErrApp`][75] or [`*cast.ErrCast`][78]. See the [gox][74] package for all available options.
  For more details about `gox` see [the module documentation][56]. The source code of `gox` is [available in the GitHub repository][31].
- `RegisterApp(name, displayName, pathRel string) error` — creates and stores a new application configuration. Note that the package path must be relative to the project root directory!
  It returns an error of type [\*app.ErrApp][75] when the application path is not relative to the project root directory, when it is not a subdirectory of it or when any other error occurs.
- `Validate() error` — ensures that all casters are properly initialized and available. It returns an error of type [\*cast.ErrCast][78] when the validation of any of the supported casters fails.
- `New(opts ...Option) (*Elder, error)` — creates a new elder wand.
  The module name is determined automatically using the [`runtime/debug`][66] package. The absolute path to the root directory is automatically set based on the current working directory. Note that the working directory must be set manually when the “magefile“ is not placed in the root directory by pointing Mage to it:
  - `-d <PATH>` option to set the directory from which “magefiles“ are read (defaults to `.`).
  - `-w <PATH>` option to set the working directory where “magefiles“ will run (defaults to value of `-d` flag).
    If any error occurs it will be of type [\*cast.ErrCast][78] or [\*project.ErrProject][80].

It is customizable through the following functions:

- `WithGobinCasterOptions(opts ...castGobin.Option) elder.Option` — sets [“gobin“ caster][43] options.
- `WithGoToolchainCasterOptions(opts ...castGoToolchain.Option) elder.Option` — sets [Go toolchain caster][44] options.
- `WithNib(n nib.Nib) elder.Option` — sets the [log-level based line printer for human-facing messages][57].
- `WithProjectOptions(opts ...project.Option) elder.Option` — sets [project][65] options.

</details>

<details>
<summary><strong>Initial project documentation</strong> — #43 ⇄ #44 (⊶ c953c4b1)</summary>

↠ The initial project documentation includes…

1. …an overview of the project features.
2. …information about the project motivation:
   1. “Why should I use [Mage][1]…“
   2. “…and why _wand_?“
3. …the project design decisions and how to use it:
   1. The overall wording and inspiration.
   2. A basic overview of the API.
   3. An introduction to the “elder“ reference implementation.
4. …information about how to contribute to this project.

</details>

<p align="center">Copyright &copy; 2019-present <a href="https://www.svengreb.de" target="_blank" rel="noreferrer">Sven Greb</a></p>

<p align="center"><a href="https://github.com/svengreb/wand/blob/main/license" target="_blank" rel="noreferrer"><img src="https://img.shields.io/static/v1.svg?style=flat-square&label=License&message=MIT&logoColor=eceff4&logo=github&colorA=4c566a&colorB=88c0d0"/></a></p>

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

<!-- Base -->

[1]: https://magefile.org
[2]: https://git-scm.com
[3]: https://www.npmjs.com
[4]: https://nodejs.org

<!-- Shared -->

[5]: https://github.com/golang/go/issues/40276
[6]: https://github.com/golang/go/milestone/145
[7]: https://github.com/markbates/pkger/issues/114
[8]: https://github.com/myitcv/gobin
[9]: https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
[10]: https://pkg.go.dev/cmd/gofmt
[11]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder
[12]: https://pkg.go.dev/github.com/markbates/pkger
[13]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gobin#Runner
[14]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang#Runner
[15]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#GoModule
[16]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#KindGoModule
[17]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Runner
[18]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Task
[19]: https://pkg.go.dev/github.com/svengreb/wand#Wand
[20]: https://golang.org/ref/mod
[21]: https://trunkbaseddevelopment.com/monorepos
[22]: https://en.wikipedia.org/wiki/Harry_Potter
[58]: https://pkg.go.dev/golang.org/x/tools/cmd/goimports
[115]: https://github.com/actions/setup-node

<!-- v0.1.0 -->

[23]: https://github.com/svengreb/wand/compare/dbf11bc0...v0.1.0
[24]: https://github.com/golang/go/wiki
[25]: https://github.com/golang/go/issues/25922
[26]: https://github.com/golang/go/issues/27653
[27]: https://github.com/golang/go/issues/30515
[28]: https://github.com/golang/tools/tree/master/cmd/goimports
[29]: https://github.com/golangci/golangci-lint/tree/master/cmd/golangci-lint
[30]: https://github.com/golangci/golangci-lint
[31]: https://github.com/mitchellh/gox
[32]: https://github.com/myitcv/gobin/wiki/FAQ
[33]: https://github.com/svengreb/tmpl-go/releases/tag/v0.3.0
[34]: https://github.com/svengreb/tmpl-go/tree/main/apps
[35]: https://github.com/svengreb/tmpl-go
[36]: https://github.com/svengreb/wand/issues/11
[37]: https://github.com/svengreb/wand/issues/15
[38]: https://github.com/svengreb/wand/issues/9
[39]: https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control
[40]: https://golang.org/ref/mod#mod-commands
[41]: https://tip.golang.org/doc/go1.16#modules
[42]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast
[43]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/gobin
[44]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/golang/toolchain
[45]: https://pkg.go.dev/cmd/go
[46]: https://pkg.go.dev/cmd/go/#hdr-Compile_packages_and_dependencies
[47]: https://pkg.go.dev/cmd/go/#hdr-Environment_variables
[48]: https://pkg.go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them
[49]: https://pkg.go.dev/cmd/go/#hdr-Print_Go_environment_information
[50]: https://pkg.go.dev/cmd/go/#hdr-Test_packages
[51]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#KindGoModule
[52]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#Validate
[53]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang#CompileFormula
[54]: https://pkg.go.dev/github.com/golangci/golangci-lint
[55]: https://pkg.go.dev/github.com/golangci/golangci-lint/cmd/golangci-lint
[56]: https://pkg.go.dev/github.com/mitchellh/gox
[57]: https://pkg.go.dev/github.com/svengreb/nib
[59]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#BinaryCaster
[60]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#Caster
[61]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#GoCode
[62]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Incantation
[63]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Mixin
[64]: https://pkg.go.dev/os/#UserCacheDir
[65]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project
[66]: https://pkg.go.dev/runtime/debug
[67]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell
[68]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/fs/clean
[69]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/goimports
[70]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang
[71]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang/build
[72]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golang/test
[73]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/golangcilint
[74]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell/gox
[75]: https://pkg.go.dev/github.com/svengreb/wand/pkg/app#ErrApp
[76]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/gobin#Caster
[77]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast/golang#Caster
[78]: https://pkg.go.dev/github.com/svengreb/wand/pkg/cast#ErrCast
[79]: https://pkg.go.dev/github.com/svengreb/nib#Verbosity
[80]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#ErrProject
[81]: https://pkg.go.dev/github.com/svengreb/wand/pkg/project#GoModuleID
[82]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#ErrGoCode
[83]: https://golang.org/ref/mod#go-mod-file
[84]: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
[85]: https://golangci-lint.run
[86]: https://en.wikipedia.org/wiki/PATH_(variable)

<!-- v0.2.0 -->

[87]: https://github.com/svengreb/wand/compare/v0.1.0...v0.2.0
[88]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Kind
[89]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#Binary
[90]: https://pkg.go.dev/github.com/svengreb/wand/pkg/spell#GoModule
[91]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Exec
[92]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#Options
[93]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task#RunnerExec
[94]: https://magefile.org/targets
[95]: https://en.wikipedia.org/wiki/Command-line_interface#Anatomy_of_a_shell_CLI
[96]: https://en.wikipedia.org/wiki/Executable

<!-- v0.3.0 -->

[97]: https://github.com/svengreb/wand/compare/v0.2.0...v0.3.0
[98]: https://github.com/golang/go/issues/41191
[99]: https://github.com/imdario/mergo/compare/v0.3.9...v0.3.11
[100]: https://github.com/markbates/pkger/issues/109
[101]: https://github.com/markbates/pkger/issues/121
[102]: https://github.com/Masterminds/semver/compare/v3.1.0...v3.1.1
[103]: https://pkg.go.dev/cmd/go/#hdr-List_packages_or_modules
[104]: https://pkg.go.dev/github.com/svengreb/wand/elder#Elder.Pkger
[105]: https://pkg.go.dev/github.com/imdario/mergo
[106]: https://pkg.go.dev/github.com/Masterminds/semver/v3
[107]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/pkger
[108]: https://go.googlesource.com/proposal/+/master/design/draft-embed.md

<!-- v0.4.0 -->

[109]: https://github.com/svengreb/wand/compare/v0.3.0...v0.4.0
[110]: https://github.com/mvdan/gofumpt#added-rules
[111]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder#Elder.Gofumpt
[112]: https://pkg.go.dev/mvdan.cc/gofumpt
[113]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gofumpt

<!-- v0.4.1 -->

[114]: https://github.com/actions/setup-node/compare/v2.1.3...c46424ee
[116]: https://github.com/svengreb/wand/compare/v0.4.0...v0.4.1
[117]: https://github.com/magefile/mage/compare/v1.10.0...v1.11.0
[118]: https://pkg.go.dev/os/#Environ
[119]: https://pkg.go.dev/github.com/magefile/mage
[120]: https://pkg.go.dev/github.com/svengreb/wand@v0.4.0/pkg/task/gobin#WithEnv
[121]: https://pkg.go.dev/github.com/svengreb/wand@v0.4.0/pkg/task/gobin#Runner.Install
[122]: https://magefile.org/targets/#arguments

<!-- v0.5.0 -->

[123]: https://github.com/svengreb/wand/compare/v0.4.1...v0.5.0
[124]: https://github.com/svengreb/tmpl-go/releases/tag/v0.7.0
[125]: https://blog.golang.org/go116-module-changes#TOC_4.
[126]: https://blog.golang.org/go1.16
[127]: https://golang.org/doc/go1.16#library-embed
[128]: https://pkg.go.dev/embed
[129]: https://pkg.go.dev/github.com/svengreb/wand@v0.4.1/pkg/task/gobin
[130]: https://pkg.go.dev/github.com/svengreb/wand@v0.4.1/pkg/task/pkger
[131]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/install

<!-- v0.6.0 -->

[132]: https://github.com/bwplotka/bingo
[133]: https://github.com/svengreb/wand/compare/v0.5.0...v0.6.0
[134]: https://github.com/golang/go/issues/42088
[135]: https://github.com/golang/go/issues/44469#issuecomment-784534876
[136]: https://github.com/myitcv/gobin/issues/103
[137]: https://github.com/oligot/go-mod-upgrade
[138]: https://github.com/svengreb/tmpl-go/releases/tag/v0.8.0
[139]: https://github.com/svengreb/tmpl-go/issues/56
[140]: https://github.com/svengreb/tmpl-go/issues/58
[141]: https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information
[142]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder#Elder.GoModUpgrade
[143]: https://pkg.go.dev/github.com/svengreb/wand/pkg/elder#Elder
[144]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/golang/env
[145]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task/gomodupgrade
[146]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/app#Config
[147]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/elder#Elder.Bootstrap
[148]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/task/gobin#Runner
[149]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/task/gofumpt#New
[150]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/task/goimports#New
[151]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/task/golang/build#New
[152]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/task/golang/install#New
[153]: https://pkg.go.dev/github.com/svengreb/wand@v0.5.0/pkg/task#Runner
[154]: https://www.jetbrains.com/help/idea/settings-tools-startup-tasks.html
[155]: https://docs.npmjs.com/cli/v7/configuring-npm/folders#node-modules
[156]: https://en.wikipedia.org/wiki/Eating_your_own_dog_food

<!-- v0.7.0 -->

[157]: https://github.com/svengreb/wand/compare/v0.6.0...v0.7.0
[158]: https://github.com/fatih/color
[159]: https://github.com/fatih/color/compare/v1.10.0...v1.11.0
[160]: https://github.com/fatih/color/compare/v1.11.0...v1.12.0
[161]: https://github.com/fatih/color/compare/v1.12.0...v1.13.0
[163]: https://github.com/actions/setup-node/compare/v2.1.5...v2.4.1
[164]: https://github.com/svengreb/tmpl-go/releases/tag/v0.9.0
[165]: https://github.com/svengreb/tmpl-go/issues/64
[166]: https://github.com/svengreb/tmpl-go/issues/66
[167]: https://github.com/svengreb/tmpl-go/issues/68
[168]: https://github.com/svengreb/tmpl-go/issues/70
[169]: https://pkg.go.dev/github.com/svengreb/wand/pkg/task@v0.6.0#GoModule
[170]: https://github.com/mvdan/gofumpt/releases/tag/v0.2.0
[171]: https://pkg.go.dev/github.com/svengreb/wand@v0.6.0/pkg/task/gofumpt
[173]: https://pkg.go.dev/github.com/svengreb/wand@v0.6.0/pkg/task/gofumpt#WithSimplify
[175]: https://pkg.go.dev/github.com/svengreb/wand@v0.6.0/pkg/task/goimports
[176]: https://pkg.go.dev/golang.org/x/tools@v0.1.7/cmd/goimports
[177]: https://github.com/golangci/golangci-lint/cmd/golangci-lint
[178]: https://pkg.go.dev/github.com/svengreb/wand@v0.6.0/pkg/task/golangcilint
[179]: https://github.com/golangci/golangci-lint/releases/tag/v1.43.0
[180]: https://github.com/svengreb/wand/tree/9caf10f9d3b0c97e1f6c18b29c175e71764b0ece/.github/workflows
[181]: https://github.com/actions/checkout
[182]: https://github.com/svengreb/wand/runs/4275275079?check_suite_focus=true
[183]: https://github.com/actions/checkout#Fetch-all-history-for-all-tags-and-branches
[184]: https://github.com/svengreb/tmpl-go/releases/tag/v0.10.0
[185]: https://github.com/svengreb/tmpl-go/issues/72
[186]: https://github.com/svengreb/tmpl-go/issues/74

<!-- v0.8.0 -->

[187]: https://github.com/svengreb/wand/compare/v0.7.0...v0.7.1
[189]: https://github.com/svengreb/tmpl-go/issues/76
[190]: https://github.com/svengreb/tmpl-go/issues/78
[192]: https://github.com/golang/go/commit/9cec77ac#diff-abdadaf0d85a2e6c8e45da716909b2697d830b0c75149b9e35accda9c38622bdR2234
[193]: https://pkg.go.dev/runtime/debug@go1.18#ReadBuildInfo
[194]: https://pkg.go.dev/runtime/debug#Module
[195]: https://github.com/golang/go/blob/9cec77ac/src/runtime/debug/mod.go#L52
[196]: https://github.com/golang/go/blob/122a22e0e9eba7fe712030d429fc4bcf6f447f5e/src/cmd/go/internal/load/pkg.go#L2288
[197]: https://pkg.go.dev/runtime/debug@go1.18.8
[198]: https://pkg.go.dev/golang.org/x/mod/modfile
[199]: https://pkg.go.dev/cmd/go#hdr-The_go_mod_file
[200]: https://tip.golang.org/doc/go1.18#go-version
[201]: https://tip.golang.org/doc/go1.18#debug/buildinfo
[202]: https://github.com/svengreb/tmpl-go/issues/91
[203]: https://github.com/svengreb/tmpl-go/releases/tag/v0.11.0
[204]: https://github.com/svengreb/tmpl-go/releases/tag/v0.12.0

<!--lint disable no-duplicate-definitions-->

[^1]: https://go.dev/ref/mod#go-mod-file
