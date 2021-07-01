// Package file provides per-file state helpers, such as for tracking imports.
package file

import (
	"fmt"

	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// CoreImportPath is the path to the core import path.
const CoreImportPath = "github.com/diamondburned/gotk4/core"

// Header describes the side effects of the conversion, such as importing new
// things or modifying the CGo preamble. A zero-value instance is a valid
// instance.
type Header struct {
	Imports        map[string]string
	CIncludes      map[string]struct{}
	Packages       map[string]struct{} // for pkg-config
	Callbacks      map[string]struct{}
	CallbackDelete bool
}

// ImportCore adds a core import.
func (h *Header) ImportCore(core string) {
	h.Import(CoreImportPath + "/" + core)
}

func (h *Header) Import(path string) {
	h.ImportAlias(path, "")
}

func (h *Header) ImportAlias(path, alias string) {
	if h.Imports == nil {
		h.Imports = map[string]string{}
	}

	h.Imports[path] = alias
}

// needsExternGLib adds the external gotk3/glib import.
func (h *Header) NeedsExternGLib() {
	h.ImportAlias("github.com/gotk3/gotk3/glib", "externglib")
}

func (h *Header) ImportPubl(resolved *types.Resolved) {
	if resolved == nil {
		return
	}

	h.ImportResolvedType(resolved.PublImport)

	if resolved.Extern != nil {
		callback, ok := resolved.Extern.Type.(*gir.Callback)
		if ok {
			h.AddCallback(callback)
		}
	}
}

func (h *Header) ImportImpl(resolved *types.Resolved) {
	if resolved == nil {
		return
	}

	h.ImportResolvedType(resolved.ImplImport)

	if resolved.Extern != nil {
		callback, ok := resolved.Extern.Type.(*gir.Callback)
		if ok {
			h.AddCallback(callback)
		}
	}
}

func (h *Header) ImportResolvedType(imports types.ResolvedImport) {
	if imports.Path != "" {
		h.ImportAlias(imports.Path, imports.Package)
	}
}

func (h *Header) AddCallback(callback *gir.Callback) {
	h.AddCallbackHeader(CallbackCHeader(callback))
}

// CallbackPrefix is the prefix to prepend to a C callback that bridges CGo.
// Generators should use this prefix when generating.
const CallbackPrefix = "gotk4_"

// CallbackCHeader renders the C function signature.
func CallbackCHeader(cb *gir.Callback) string {
	var ctail pen.Joints
	if cb.Parameters != nil {
		ctail = pen.NewJoints(", ", len(cb.Parameters.Parameters))

		for _, param := range cb.Parameters.Parameters {
			ctail.Add(types.AnyTypeC(param.AnyType))
		}
	}

	cReturn := "void"
	if cb.ReturnValue != nil {
		cReturn = types.AnyTypeC(cb.ReturnValue.AnyType)
	}

	goName := strcases.PascalToGo(cb.Name)
	return fmt.Sprintf("%s %s(%s);", cReturn, CallbackPrefix+goName, ctail.Join())
}

func (h *Header) AddCallbackHeader(header string) {
	if h.Callbacks == nil {
		h.Callbacks = map[string]struct{}{}
	}

	h.Callbacks[header] = struct{}{}
}

// AddPackage adds a pkg-config package.
func (h *Header) AddPackage(pkg string) {
	if h.Packages == nil {
		h.Packages = map[string]struct{}{}
	}

	h.Packages[pkg] = struct{}{}
}

// IncludeC adds a C header file into the cgo preamble.
func (h *Header) IncludeC(include string) {
	if h.CIncludes == nil {
		h.CIncludes = map[string]struct{}{}
	}

	h.CIncludes[include] = struct{}{}
}

// needsCbool adds the C stdbool.h include.
func (h *Header) needsCbool() {
	h.IncludeC("stdbool.h")
}

// NeedsGLibObject adds the glib-object.h include and the glib-2.0 package.
func (h *Header) NeedsGLibObject() {
	// Need this for g_value_get_boxed.
	h.IncludeC("glib-object.h")
	// Need this for the above header.
	h.AddPackage("glib-2.0")
}

// ApplyHeader applies the side effects of the conversion. The caller is
// responsible for calling this.
func (h *Header) ApplyHeader(dst *Header) {
	if h.CallbackDelete {
		dst.CallbackDelete = true
	}
	for path, alias := range h.Imports {
		dst.ImportAlias(path, alias)
	}
	for callback := range h.Callbacks {
		dst.AddCallbackHeader(callback)
	}
	for cIncl := range h.CIncludes {
		dst.IncludeC(cIncl)
	}
	for pkg := range h.Packages {
		dst.AddPackage(pkg)
	}
}
