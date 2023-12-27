# Contributing to gotk4

## Prerequisites

This project is being developed with the use of [Nix][nix]. This ensures that
developers have a pure and constant environment where builds and generation
outputs are easily reproducible to the exact line of code.

[nix]: https://nixos.org/guides/how-nix-works.html

However, this also means that contributors are strongly recommended to have Nix
or Docker on their machine. This helps ensure that pull requests to the
repository have consistent versioning and outputs, and that nothing has been
changed where they don't have to be.

### Using Nix for development

If you haven't installed [Nix][nix] yet, see the [Install Nix][install-nix]
section of the Nix guides. The prerequisite is "any Linux distribution, macOS
[or] Windows (via WSL)."

[install-nix]: https://nixos.org/guides/install-nix.html

Afterwards, to use Nix for developing, simply run

```sh
nix-shell
```

This will drop the user into a new shell with Go, GLib libraries and more, and
all those dependencies will be versioned in a constant fashion as pinned in the
shell.nix file.

### Using Docker for development

A small Dockerfile is provided for developers who don't have Nix installed, but
do have Docker running. The Docker environment will contain everything that the
Nix environment does, including reproducible, pinned inputs.

To build the Dockerfile and regenerate the repository using it, do:

```sh
docker build -t gotk4 .
docker run --rm --volume "$PWD:/gotk4/" -u "$(id -u):$(id -g)" gotk4 generate
```

To build the generated packages, run:

```sh
docker run --rm --volume "$PWD:/gotk4/" -u "$(id -u):$(id -g)" gotk4 build
```

If neither `generate` nor `build` is provided, then the `-it` flag could be
added for Docker to successfully drop the user into a Bash shell with the right
environments.

## Generating Bindings

To re-generate the gotk4 bindings, simply run:

```sh
go generate
```

If additional debug information is needed, then the `GIR_VERBOSE` environment
variable can be set to `1`:

```sh
GIR_VERBOSE=1 go generate
```

This will print out much more information about the generation process,
including the reasons why certain types and functions are skipped.

Assuming you're in the Nix environment, the generated code should always be
idempotent. This means that the generated code should always be the same
no matter how many times you run `go generate`. If this is not the case, then
please file an issue.

If you need to generate gotk4 for a newer version of GTK, see the [Updating
Nixpkgs](#updating-nixpkgs) section below.

## Updating Nixpkgs

In case a new release of something is added to Nixpkgs Unstable, for example, a
new GTK+ version, then the `sourceNixpkgs.rev` and `sourceNixpkgs.sha256`
variables inside the [`.nix/default.nix`][] file should be updated
appropriately.  Some packages, such as Go, do not come directly from Nixpkgs,
but from [gotk4-nix.git][]; to update the gotk4-nix, then very similarly the
`gotk4-nix.rev` and `gotk4-nix.sha256` variables inside [`.nix/default.nix`][]
should be updated.

> [!NOTE]
> Some repositories such as [gotk4-adwaita][gotk4-adwaita] may choose to
> contain all this information in its single `shell.nix` file for convenience.
> In this case, the `sourceNixpkgs` and `gotk4-nix` variables should be updated
> there instead.

[`.nix/default.nix`]: ./.nix/default.nix
[gotk4-nix.git]: https://github.com/diamondburned/gotk4-nix
[gotk4-adwaita]: https://github.com/diamondburned/gotk4-adwaita

Care should be taken when updating Nixpkgs, as the output could change
drastically, and Nixpkgs versions should **never** be downgraded to ensure the
newly generated code doesn't break existing code.

## Project Layout

This project is divided into several parts:

- [`gir/`][] contains the library that provides data structures and low-level
  wrappers over parsing, reading and storing GIR files.
    - [`gir/cmd/`][] contains executable tools to generate either the whole
      repository or provide helper tools to ease development.
      - [`cmd/gir-generate/`][] contains the main generator program.
    - [`gir/girgen/`][] contains the generator code that consumes stored GIR
      files and generate Go bindings appropriately.
- [`pkg/`][] is the directory for all generated code.
- [`gotk4.go`][] contains the `//go:generate` boilerplate.
- [`shell.nix`][] and [`.nix/`][] contain the pinned Nixpkgs that ensures
  constant development environments.


[`gir/`]:                  ./gir/
[`gir/cmd/`]:              ./gir/cmd/
[`gir/cmd/gir-generate/`]: ./gir/cmd/gir-generate/
[`gir/girgen/`]:           ./gir/girgen/
[`pkg/`]:                  ./pkg/
[`gotk4.go`]:              ./gotk4.go
[`shell.nix`]:             ./shell.nix
[`.nix/`]:                 ./.nix/

## Project Guidelines

Below is a list of guidelines to keep in mind (in no particular order) when
contributing to this project.

0. Using Nix is recommended for reasons mentioned above.
1. Contributors should always use `goimports` over what their IDE provides
   (including GoLand). This ensures that no sections of the manually-written
   code are changed unnecessarily. It would be a given that all contributions
   should also be formatted using `goimports` (or `go fmt` if no imports were
   changed). As a sidenote, `goimports` is included in the Nix shell.
2. `.gitignore` files should be avoided to not clutter the repository
   up.  Developers using editors that produce scrap files should add
   them to their own `~/.config/git/ignore` so that they are not
   tempted to clutter up the `.gitignore` files in every repository
   that they contribute to.
3. Project-wide refactors are large and will consume a lot of time to review, so
   they should be avoided.
