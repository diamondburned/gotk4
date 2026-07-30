# Changelog

All notable changes to `gotk4` will be documented in this file.

## [4.22.4] - 2026-07-30

### Version Upgrades

- **GTK4 Package Version**: Bumped from **4.14.4** to **4.22.4**.
- **Go Toolchain**: Updated to **1.26.5** (`nixos-26.05`).
- **Nixpkgs Channel**: Synchronized with `nixos-26.05`
  (`21ea275a7c46aef9d4d6ddc962e6d562e9d94183`).
- **`gotk4-nix` Flake Input**: Pointed directly to
  `github:diamondburned/gotk4-nix` (`ae8ad99936067cb00f7900601a59cb428f8f239e`).

### GIR Code Generation & Generator Fixes

1. **`Gdk-4.MemoryFormat` Enum Member Renaming**
   (`gir/cmd/gir-generate/gendata/gendata.go`)
   - Added `RenameEnumMembers` preprocessor rules for 3-plane formats matching
     `MEMORY_G(\d+)_(B\d+)_(R\d+)_(.*)` and `MEMORY_G(\d+)_(R\d+)_(B\d+)_(.*)`
     to append `_3PLANE` (yielding identifiers like `MemoryG8B8R84203Plane`) to
     resolve identifier collisions with existing 2-plane formats.

2. **C-to-Go Fixed-Size Array Slicing** (`gir/girgen/types/typeconv/c-go.go`)
   - Added type prefix check (`strings.HasPrefix(value.In.Type, "*")`). For
     pointer types, `unsafe.Slice(src, len)` is emitted; for array types
     (`[N]C.T`), `unsafe.Slice(&(src)[0], len)` is emitted.

3. **`gotk4-nix` Hook Alignment** (`build-shell.nix`, `build-package.nix`)
   - Updated hook references to
     `(pkgs.wrapGAppsHook4 or pkgs.wrapGAppsHook3 or pkgs.wrapGAppsHook)` for
     compatibility with Nixpkgs hook renames.

### New APIs Overview

#### GTK 4 (`pkg/gtk/v4`)

- **New Widgets & Expressions**:
  - `PopoverBin`: Bin container for popovers and menu models with child input
    handling (`NewPopoverBin`).
  - `SVG`: Dedicated SVG rendering widget and paintable with support for loading
    from bytes or resources (`NewSVG`, `NewSVGFromBytes`, `NewSVGFromResource`).
  - `TryExpression`: Expression type (`GtkTryExpression`) supporting fallback
    evaluation chains (`NewTryExpression`).
- **Accessibility Improvements**:
  - `AccessibleHyperlink` & `AccessibleHypertext`: Hyperlinked text
    accessibility interfaces and objects (`NewAccessibleHyperlink`).
  - Added `Accessible.AccessibleID()` and `Accessible.UpdatePlatformState()`.
- **System Preferences & Display Enums**:
  - Added system settings enums: `FontRendering`, `InterfaceColorScheme`,
    `InterfaceContrast`, `ReducedMotion`, and `WindowGravity`.
- **Dialogs & Text Buffer Updates**:
  - `FileDialog`: Helper methods for async text file operations with encoding
    and line-ending detection (`OpenTextFile`, `OpenMultipleTextFiles`,
    `SaveTextFile`).
  - `TextBufferCommitNotify` and `TextBufferNotifyFlags` for tracking text
    buffer commit events.

#### GDK 4 (`pkg/gdk/v4`)

- **Color Management & Color States**:
  - `ColorState`: First-class support for color spaces including sRGB, Rec.2100,
    Oklab, and Oklch (`ColorStateGetSrgb`, `ColorStateGetSrgbLinear`,
    `ColorStateGetRec2100Linear`, `ColorStateGetRec2100Pq`,
    `ColorStateGetOklab`, `ColorStateGetOklch`).
  - `CICPParams`: Handling Coding-Independent Code Points (CICP) color
    properties (`NewCICPParams`, `BuildColorState`).
- **Texture Builders**:
  - `MemoryTextureBuilder`: Builder API for constructing memory textures with
    color state, multi-plane strides, and offsets (`NewMemoryTextureBuilder`).
  - Integrated `ColorState` into `DmabufTextureBuilder`, `GLTextureBuilder`, and
    `TextureDownloader`.
- **Display & Windowing**:
  - `Monitor`: Monitor query methods (`Connector`, `Description`, `Display`,
    `Geometry`, `RefreshRate`) and `ConnectInvalidate` signal handle.
  - Enums/Flags: Added `ScrollRelativeDirection`, `ToplevelCapabilities`,
    `CICPRange`, and `ColorChannel`.

#### GSK 4 (`pkg/gsk/v4`)

- **New Render Nodes**:
  - `ComponentTransferNode`: Render node applying discrete, gamma, linear, or
    table component transfer functions to color channels
    (`NewComponentTransferNode`, `ComponentTransfer`).
  - `CompositeNode`: Render node blending child and mask nodes using Porter-Duff
    blend modes (`NewCompositeNode`, `PorterDuff`).
  - `CopyNode` & `PasteNode`: Inter-layer rendering nodes for copying/pasting
    render operations (`NewCopyNode`, `NewPasteNode`).
  - `IsolationNode`: Group isolation render node (`NewIsolationNode`,
    `Isolation`).
- **Path Operations & RenderNode Introspection**:
  - `Path`: Added `ForEachIntersection`, `TightBounds`, and `Equal` comparison
    helpers.
  - `RenderNode`: Added `Children()` and `OpaqueRect()` introspection methods.

#### Gio (`pkg/gio/v2`)

- **Socket Control Messages**:
  - `IPToSMessage`: Socket control message for IP Type of Service (ToS) DSCP and
    ECN headers (`NewIPToSMessage`).
  - `IPv6TClassMessage`: IPv6 Traffic Class socket control message
    (`NewIPv6TClassMessage`).
- **Enums & Flags**:
  - Added `ECNCodePoint` for Explicit Congestion Notification codepoints.
  - Added `IOModuleScopeFlags`, `TLSCertificateRequestFlags`, and
    `TLSDatabaseLookupFlags`.

### Affected Packages

- `pkg/gdk/v4`: Regenerated with GTK 4.22 additions and unique 3-plane
  `MemoryFormat` constants.
- `pkg/gsk/v4`: Regenerated with GTK 4.22 render node updates, fixing
  `BorderNode.Colors()` array slicing.
- `pkg/gtk/v4`: Updated with GTK 4.22.4 API changes.
- `pkg/gio/v2`, `pkg/glib/v2`, `pkg/pango`, `pkg/pangocairo`, `pkg/atk`,
  `pkg/graphene`: Regenerated against updated Nixpkgs GIR definitions.
