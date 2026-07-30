# Updating gotk4

This guide details the procedure for updating `gotk4` to newer GTK and Nixpkgs
versions.

## Updating `gotk4-nix`

1. **Update Flake Inputs**
   - In `gotk4-nix`, run `nix flake update`.
   - Ensure shell and package derivation hooks remain compatible with the latest
     Nixpkgs (e.g., handling renamed hooks like `wrapGAppsHook` ->
     `wrapGAppsHook4`).

2. **Submit PR & Merge**
   - If any changes are required in `gotk4-nix`, create a pull request.
   - Wait for the PR to be reviewed and merged into `gotk4-nix` main branch
     before proceeding to update `gotk4`.

## Updating `gotk4`

_Assumes `gotk4-nix` is up-to-date and merged, and that you're inside
`nix develop` already._

1. **Update `gotk4` Flake Inputs**
   - Update `gotk4/flake.nix` to point `nixpkgs-gotk4` to `nixos-26.05` (or
     desired target channel/release branch) to prevent GTK versions from
     advancing too far.
   - Update `gotk4-nix` input in `flake.nix` / `flake.lock` to the latest merged
     version of `gotk4-nix`.
   - Ensure the Go package reference in `flake.nix` matches available attributes
     in Nixpkgs (`pkgs.go`).
   - Run `nix flake update` in `gotk4`.

2. **Run Code Generation**
   - Run `go generate` inside the Nix development shell:
     ```sh
     go generate
     ```
   - Resolve any codegen errors or enum member collisions by adding
     preprocessors or filters in `gir/cmd/gir-generate/gendata/gendata.go`.

3. **Validate Package Compilation**
   - Build all generated packages in `pkg/`:
     ```sh
     (cd pkg && go build -v ./...)
     ```
   - If CGo compilation errors occur due to generated code, update generator
     logic under `gir/girgen/` or preprocessors/filters in `gendata.go`, then
     re-run `go generate` and `go build`.

4. **Run Tests**
   - Validate tests pass across the repository:
     ```sh
     go test ./...
     cd pkg && go test ./...
     ```

5. **Document Changes in `CHANGELOG.md`**
   - Add a new release section in `CHANGELOG.md` corresponding to the target
     version (e.g., `## [4.22.4] - YYYY-MM-DD`).
   - Include the following sections in each changelog entry:
      - **Required:** Version updates. Record GTK4 version bump, Go toolchain
       version, Nixpkgs channel/commit, and `gotk4-nix` flake input commit.
     - **Required:** GIR Code Generation & Generator Fixes.
     - Optional: Summarize new widgets, render nodes, color management APIs, and
       network/socket control messages across `GTK 4`, `GDK 4`, `GSK 4`, and
       `Gio`. Pay special attention to exact Go casing for functions, types, and
       constants.

6. **Branch Management**
   - Identify the GTK minor version from `pkg-config --modversion gtk4` (e.g.,
     `4.22.4` -> minor version `22`).
   - Keep generated code on branch `4` and update/create branch `4.<minor>`
     (e.g., `4.22`).

## Pitfalls

> [!NOTE]
>
> This section is being expanded with each version bump. If you encounter any
> GIR issues that are worth documenting, please call them out here.

- **Enum Member Collisions**: When upstream GTK introduces new enum values whose
  C identifiers normalize to the same camelCase Go identifier (such as 2-plane
  vs 3-plane formats in `GdkMemoryFormat`), `RenameEnumMembers` preprocessors
  must be added.

- **Go Regexp Replacement Syntax**: In Go `regexp.ReplaceAllString`, group
  capture replacements followed by underscores or alphanumeric characters must
  use braces (e.g., `${1}_${2}_3PLANE`), because `$1_` attempts to look up a
  named group `$1_`.

- **CGo Fixed-Size Array Slicing**: `unsafe.Slice(ptr, len)` requires a pointer
  (`*T`). When CGo generates a pointer type (`*C.T`), pass `ptr` directly; when
  CGo generates an array type (`[N]C.T`), pass `&(arr)[0]`.

- **Nixpkgs Hook Renames**: Nixpkgs periodically renames hooks (e.g.,
  `wrapGAppsHook` -> `wrapGAppsHook4`). Fallback expressions like
  `(pkgs.wrapGAppsHook4 or pkgs.wrapGAppsHook3 or pkgs.wrapGAppsHook)` preserve
  cross-channel compatibility.
