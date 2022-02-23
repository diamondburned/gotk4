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

// Marshaler describes a marshaler.
type Marshaler struct {
	GLibGetType string
	GoTypeName  string
}

// GLibType returns a type-casted get_type function call.
func (m Marshaler) GLibType() string {
	// Small hack to allow overriding.
	if strings.Contains(m.GLibGetType, ".") {
		return fmt.Sprintf("externglib.Type(%s)", m.GLibGetType)
	}
	return fmt.Sprintf("externglib.Type(C.%s())", m.GLibGetType)
}

// Header describes the side effects of the conversion, such as importing new
// things or modifying the CGo preamble. A zero-value instance is a valid
// instance.
type Header struct {
	Marshalers     []Marshaler
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

	h.Marshalers = append(h.Marshalers, Marshaler{glibGetType, goName})
	// Need this for g_value functions inside marshal.
	h.NeedsGLibObject()
	// Need this for the pointer cast (which one?).
	h.Import("unsafe")
}

func (h *Header) AddCallback(source *gir.NamespaceFindResult, callback *gir.Callback) {
	h.AddCallable(source, "", &callback.CallableAttrs)
}

// AddCallable adds an extern C function header from the callable. The extern
// function will have the given name.
func (h *Header) AddCallable(source *gir.NamespaceFindResult, name string, callable *gir.CallableAttrs) {
	h.AddCallbackHeader(CallableCHeader(source, name, callable))
}

const callbackPrefix = "_gotk4"

// CallableExportedName creates the exported C name of the given callback from
// the given namespace.
func CallableExportedName(source *gir.NamespaceFindResult, callable *gir.CallableAttrs) string {
	return ExportedName(source, strcases.PascalToGo(callable.Name))
}

// ExportedName generates the full name for any exported function or variable
// generated by gotk4. It prepends the right strings to make the exported
// identifier unique in the C namespace.
func ExportedName(source *gir.NamespaceFindResult, suffixes ...string) string {
	namespaceName := strings.ToLower(source.Namespace.Name)
	if source.Namespace.Version != "" {
		namespaceName += gir.MajorVersion(source.Namespace.Version)
	}

	return callbackPrefix + "_" + namespaceName + "_" + strings.Join(suffixes, "_")
}

// CallableCHeader renders the C function signature.
func CallableCHeader(source *gir.NamespaceFindResult, name string, callable *gir.CallableAttrs) string {
	var ctail pen.Joints
	if callable.Parameters != nil {
		ctail = pen.NewJoints(", ", len(callable.Parameters.Parameters)+1)

		if callable.Parameters.InstanceParameter != nil {
			ctail.Add(types.AnyTypeC(callable.Parameters.InstanceParameter.AnyType))
		}
		for _, param := range callable.Parameters.Parameters {
			ctail.Add(types.AnyTypeC(param.AnyType))
		}
		if callable.Throws {
			ctail.Add("GError**")
		}
	}

	cReturn := "void"
	if callable.ReturnValue != nil {
		cReturn = types.AnyTypeC(callable.ReturnValue.AnyType)
	}

	if name == "" {
		name = CallableExportedName(source, callable)
	}

	return fmt.Sprintf("extern %s %s(%s);", cReturn, name, ctail.Join())
}

// AddCBlock adds a block of C code into the Cgo preamble.
func (h *Header) AddCBlock(block string) {
	if block == "" {
		// Bound check to prevent panic in line logic.
		return
	}

	block = strings.TrimSpace(block)

	// Guess the indentation: grab the last line and count its leading tabs.
	lines := strings.Split(block, "\n")
	lastLine := lines[len(lines)-1]

	var tabs int
	for tabs < len(lastLine) && lastLine[tabs] == '\t' {
		tabs++
	}

	// Trim all lines except for the first, since TrimSpace will take care of
	// that.
	indent := strings.Repeat("\t", tabs)
	for i := 1; i < len(lines); i++ {
		lines[i] = strings.TrimPrefix(lines[i], indent)
	}

	// Rejoin and add.
	block = strings.Join(lines, "\n")
	h.AddCallbackHeader(block)
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

// NeedsGLibObject adds the glib-object.h include and the glib-2.0 package.
func (h *Header) NeedsGLibObject() {
	// Need this for g_value_get_boxed.
	h.IncludeC("glib-object.h")
	// Need this for the above header.
	h.AddPackage("glib-2.0")
}

// CopyHeader copies the given header.
func CopyHeader(src *Header) Header {
	var dst Header
	dst.ApplyFrom(src)
	return dst
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
