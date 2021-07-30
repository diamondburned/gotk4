# Contributing to gotk4

## Prerequisites

This project is being developed with the use of [Nix][nix]. This ensures that
developers have a pure and constant environment where builds and generation
outputs are easily reproducible to the exact line of code.

[nix]: https://nixos.org/guides/how-nix-works.html

However, this also means that contributors are strongly recommended to have Nix
on their machine, similar to how other projects use Docker for development,
except Nix is a lot more lightweight. This helps ensure that pull requests to
the repository have consistent versioning and outputs, and that nothing has been
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

### Updating Nixpkgs

In case a new release of something is added to Nixpkgs Unstable, for example, a
new Go version, then the `rev` and `sha256` of the `unstable` variable inside
the [.nix/overlay.nix][shell.nix] file should be updated appropriately. The structure
of the [shell.nix][shell.nix] file is commented inside the file itself.

Care should be taken when updating Nixpkgs, as the output could change
drastically, and Nixpkgs versions should **never** be downgraded to ensure the
newly generated code doesn't break existing code.

## Project Layout

This project is divided into several parts:

- cmd/ contains executable tools to generate either the whole repository or
  provide helper tools to ease development.
  - cmd/gir_generate/ contains the main generator program.
    - cmd/gir_generate/config.go contains tuning variables such as what is
	  generated and ignored.
- gir/ contains the library that provides data structures and low-level wrappers
  over parsing, reading and storing GIR files.
    - gir/girgen/ contains the generator code that consumes stored GIR files and
	  generate Go bindings appropriately.
- internal/ contains internal libraries that the generated code can use for
  convenience.
- pkg/ is the directory for all generated code.
- gotk4.go contains the `//go:generate` boilerplate.
- [shell.nix][shell.nix] and [`.nix`][dot-nix] contain the pinned Nixpkgs that
  ensures constant development environments.

[shell.nix]: https://github.com/diamondburned/gotk4/blob/4/shell.nix
[dot-nix]: https://github.com/diamondburned/gotk4/tree/4/.nix/

## Project Guidelines

Below is a list of guidelines to keep in mind (in no particular order) when
contributing to this project.

0. Using Nix is recommended for reasons mentioned above.
1. Contributors should always use `goimports` over what their IDE provides
   (including GoLand). This ensures that no sections of the manually-written
   code are changed unnecessarily. It would be a given that all contributions
   should also be formatted using `goimports` (or `go fmt` if no imports were
   changed). As a sidenote, `goimports` is included in the Nix shell.
2. `.gitignore` files should be avoided to not clutter the repository up.
   Developers using editors that produce scrap files should add them to
   their own `.git/info/exclude`.
3. Project-wide refactors are large and will consume a lot of time to review, so
   they should be avoided.
