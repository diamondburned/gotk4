// Package file provides per-file state helpers, such as for tracking imports.
package file

import (
	"fmt"
	"sort"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// CoreImportPath is the path to the core import path.
const CoreImportPath = "github.com/diamondburned/gotk4/pkg/core"

// Headerer is an interface of something that returns its internal header.
type Headerer interface {
	Header() *Header
}

// ApplyHeader applies the given src headers to dst.
func ApplyHeader(dst Headerer, srcs ...Headerer) {
	dstHeader := dst.Header()
	for _, src := range srcs {
		dstHeader.ApplyFrom(src.Header())
	}
}

// Header describes the side effects of the conversion, such as importing new
// things or modifying the CGo preamble. A zero-value instance is a valid
// instance.
type Header struct {
	Marshalers     []string
	Imports        map[string]string
	CIncludes      map[string]struct{}
	Packages       map[string]struct{} // for pkg-config
	Callbacks      map[string]struct{}
	CallbackDelete bool

	stop bool
}

// NoopHeader is a header instance where its methods do nothing. This instance
// is useful for functions that both validate and generate, but generation is
// not wanted.
var NoopHeader = &Header{stop: true}

// Reset resets the file header to a zero state.
func (h *Header) Reset() {
	*h = Header{stop: h.stop}
}

// HasImport returns true if the header imports the given path.
func (h *Header) HasImport(path string) bool {
	_, ok := h.Imports[path]
	return ok
}

// ImportCore returns the path to import a core package.
func ImportCore(core string) string {
	return CoreImportPath + "/" + core
}

// ImportCore adds a core import.
func (h *Header) ImportCore(core string) {
	h.Import(ImportCore(core))
}

func (h *Header) Import(path string) {
	h.ImportAlias(path, "")
}

func (h *Header) ImportAlias(path, alias string) {
	if h.stop {
		return
	}

	if h.Imports == nil {
		h.Imports = map[string]string{}
	}

	// if old, ok := h.Imports[path]; ok {
	// 	if old != alias {
	// 		log.Panicf("duplicate alias old %q != new %q", old, alias)
	// 	}
	// }

	h.Imports[path] = alias
}

// needsExternGLib adds the external gotk3/glib import.
func (h *Header) NeedsExternGLib() {
	h.ImportAlias("github.com/diamondburned/gotk4/pkg/core/glib", "externglib")
}

func (h *Header) ImportPubl(resolved *types.Resolved) {
	if h.stop {
		return
	}

	if resolved == nil {
		return
	}

	if resolved.IsAbstract() {
		h.ImportResolvedType(resolved.PublImport)
	} else {
		h.ImportResolvedType(resolved.ImplImport)
	}

	if resolved.Extern != nil {
		callback, ok := resolved.Extern.Type.(*gir.Callback)
		if ok {
			h.AddCallback(resolved.Extern.NamespaceFindResult, callback)
		}
	}
}

func (h *Header) ImportImpl(resolved *types.Resolved) {
	if h.stop {
		return
	}

	if resolved == nil {
		return
	}

	h.ImportResolvedType(resolved.ImplImport)

	if resolved.Extern != nil {
		callback, ok := resolved.Extern.Type.(*gir.Callback)
		if ok {
			h.AddCallback(resolved.Extern.NamespaceFindResult, callback)
		}
	}
}

// DashImport imports the given path if it's not already imported.
func (h *Header) DashImport(path string) {
	if h.stop {
		return
	}

	if h.Imports == nil {
		h.Imports = map[string]string{}
	}

	if _, ok := h.Imports[path]; !ok {
		h.Imports[path] = "_"
	}
}

func (h *Header) ImportResolvedType(imports types.ResolvedImport) {
	if imports.Path != "" {
		h.ImportAlias(imports.Path, imports.Package)
	}
}

// AddMarshaler adds the type marshaler into the init header. It also adds
// imports.
func (h *Header) AddMarshaler(glibGetType, goName string) {
	if h.stop {
		return
	}

	h.Marshalers = append(h.Marshalers, fmt.Sprintf(
		`{T: externglib.Type(C.%s()), F: marshal%s},`, glibGetType, goName,
	))
	// Need this for g_value functions inside marshal.
	h.NeedsGLibObject()
	// Need this for the pointer cast.
	h.Import("unsafe")
}

func (h *Header) AddCallback(source *gir.NamespaceFindResult, callback *gir.Callback) {
	h.AddCallbackHeader(CallbackCHeader(source, callback))
}

const callbackPrefix = "_gotk4"

// CallbackExportedName creates the exported C name of the given callback from
// the given namespace.
func CallbackExportedName(source *gir.NamespaceFindResult, callback *gir.Callback) string {
	namespaceName := strings.ToLower(source.Namespace.Name)
	if source.Namespace.Version != "" {
		namespaceName += gir.MajorVersion(source.Namespace.Version)
	}

	goName := strcases.PascalToGo(callback.Name)

	return fmt.Sprintf("%s_%s_%s", callbackPrefix, namespaceName, goName)
}

// CallbackCHeader renders the C function signature.
func CallbackCHeader(source *gir.NamespaceFindResult, callback *gir.Callback) string {
	var ctail pen.Joints
	if callback.Parameters != nil {
		ctail = pen.NewJoints(", ", len(callback.Parameters.Parameters))

		for _, param := range callback.Parameters.Parameters {
			ctail.Add(types.AnyTypeC(param.AnyType))
		}
	}

	cReturn := "void"
	if callback.ReturnValue != nil {
		cReturn = types.AnyTypeC(callback.ReturnValue.AnyType)
	}

	return fmt.Sprintf(
		"%s %s(%s);",
		cReturn, CallbackExportedName(source, callback), ctail.Join(),
	)
}

// AddCallbackHeader adds a callback header raw.
func (h *Header) AddCallbackHeader(header string) {
	if h.stop {
		return
	}

	if h.Callbacks == nil {
		h.Callbacks = map[string]struct{}{}
	}

	h.Callbacks[header] = struct{}{}
}

// SortedCallbackHeaders returns the sorted C callback headers.
func (h *Header) SortedCallbackHeaders() []string {
	headers := make([]string, 0, len(h.Callbacks))
	for callback := range h.Callbacks {
		headers = append(headers, callback)
	}

	sort.Strings(headers)
	return headers
}

// AddPackage adds a pkg-config package.
func (h *Header) AddPackage(pkg string) {
	if h.stop {
		return
	}

	if h.Packages == nil {
		h.Packages = map[string]struct{}{}
	}

	h.Packages[pkg] = struct{}{}
}

// IncludeC adds a C header file into the cgo preamble.
func (h *Header) IncludeC(include string) {
	if h.stop {
		return
	}

	if h.CIncludes == nil {
		h.CIncludes = map[string]struct{}{}
	}

	h.CIncludes[include] = struct{}{}
}

// SortedCIncludes returns the list of C includes sorted.
func (h *Header) SortedCIncludes() []string {
	includes := make([]string, 0, len(h.CIncludes))
	for incl := range h.CIncludes {
		includes = append(includes, incl)
	}

	sort.Strings(includes)
	return includes
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

// ApplyFrom is ApplyTo but reversed.
func (h *Header) ApplyFrom(src *Header) {
	src.ApplyTo(h)
}

// ApplyTo applies the headers into the given one. The caller is responsible for
// calling this.
func (h *Header) ApplyTo(dst *Header) {
	if h.stop || dst.stop {
		return
	}

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
