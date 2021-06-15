# gotk4

[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

A GTK4 bindings generator for Go.

Progress tracker: https://github.com/diamondburned/gotk4/issues/2

All generated packages are in `pkg/`. The generation code is in `gir/girgen/`.
At the moment, the repository depends on gotk3's GLib. This may change in the
future.

## Contributing to gotk4

For contributing guidelines, see [CONTRIBUTING.md](./CONTRIBUTING.md).

## Wishes

- I wish I generated models first before generating functions
- I wish I used a registry of converted model types with their respective Go
  names to ease translation
- I wish I used common struct types that share names and such.
