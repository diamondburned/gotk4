// Package file provides per-file state helpers, such as for tracking imports.
package file

import (
	"fmt"
	"sort"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
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

// EachHeader applies f on all of headerers' headers.
func EachHeader(headerers []Headerer, f func(*Header)) {
	for _, h := range headerers {
		f(h.Header())
	}
}

// AggregateCgoStubs aggregates all the given headerers' CgoStubs maps as well
// as CgoHeader imports and sort them appropriatetly.
func AggregateCgoStubs(headerers ...Headerer) string {
	cgoStubs := map[string]cgoHeaderOrder{}
	var hasStubs bool

	for _, headerer := range headerers {
		h := headerer.Header()

		for header, ord := range h.CgoHeader {
			if ord == cgoImports {
				cgoStubs[header] = ord
			}
		}
		for stub, ord := range h.CgoStubs {
			hasStubs = true
			cgoStubs[stub] = ord
		}

		h.CgoStubs = nil
	}

	if !hasStubs {
		return ""
	}

	stubParts := [cgoHeaderLen][]string{}
	for ord := range stubParts {
		// Preallocate.
		stubParts[ord] = make([]string, 0, len(cgoStubs))
	}
	for stub, ord := range cgoStubs {
		if strings.Count(stub, "\n") > 1 {
			stub += "\n"
		}
		stubParts[ord] = append(stubParts[ord], stub)
	}

	var stubs strings.Builder

	for _, part := range stubParts {
		if len(part) == 0 {
			continue
		}
		sort.Strings(part)
		stubs.WriteString(strings.Join(part, "\n"))
		stubs.WriteString("\n\n")
	}

	return strings.TrimSuffix(stubs.String(), "\n")
}

// Header describes the side effects of the conversion, such as importing new
// things or modifying the CGo preamble. A zero-value instance is a valid
// instance.
type Header struct {
	Marshalers []string
	Imports    map[string]string
	Packages   map[string]struct{} // for pkg-config
	CgoStubs   map[string]cgoHeaderOrder
	CgoHeader  map[string]cgoHeaderOrder

	stop bool
}

type cgoHeaderOrder uint8

const (
	cgoImports cgoHeaderOrder = iota
	cgoExterns
	cgoExtras
	cgoHeaderLen // internal
)

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
	h.addCgoHeader(CallbackCHeader(source, callback), cgoExterns)
}

func (h *Header) NeedsCallbackDelete() {
	h.addCgoHeader("extern void callbackDelete(gpointer);", cgoExterns)
	h.DashImport(ImportCore("gbox"))
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

// AddCgoExtras adds extra Cgo header raw.
func (h *Header) AddCgoExtras(header string) {
	h.addCgoHeader(header, cgoExtras)
}

func (h *Header) addCgoHeader(header string, ord cgoHeaderOrder) {
	if h.stop {
		return
	}

	if h.CgoHeader == nil {
		h.CgoHeader = map[string]cgoHeaderOrder{}
	}

	h.CgoHeader[header] = ord
}

// SortedCgoHeaders returns the sorted Cgo headers.
func (h *Header) SortedCgoHeaders() []string {
	headers := make([]string, 0, len(h.CgoHeader))
	for callback := range h.CgoHeader {
		headers = append(headers, callback)
	}

	sort.Slice(headers, func(i, j int) bool {
		iord := h.CgoHeader[headers[i]]
		jord := h.CgoHeader[headers[j]]

		if iord == jord {
			return headers[i] < headers[j]
		}

		return iord < jord
	})

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
	h.addCgoHeader(fmt.Sprintf("#include <%s>", include), cgoImports)
}

// IncludeLocalC is like IncludeC, but the given include string is a local file.
func (h *Header) IncludeLocalC(include string) {
	h.addCgoHeader(fmt.Sprintf(`#include "%s"`, include), cgoImports)
}

// NeedsGLibObject adds the glib-object.h include and the glib-2.0 package.
func (h *Header) NeedsGLibObject() {
	// Need this for g_value_get_boxed.
	h.IncludeC("glib-object.h")
	// Need this for the above header.
	h.AddPackage("glib-2.0")
}

// StubCFuncHeader generates a stub C function signature for use in generating
// stub panic calls.
func StubCFuncHeader(cattrs *gir.CallableAttrs) string {
	ret := "void"
	if cattrs.ReturnValue != nil {
		// You don't actually need to write a return statement for the C
		// function. Funny!
		ret = rawCType(cattrs.ReturnValue.AnyType)
	}

	args := "void"
	if cattrs.Parameters != nil {
		joints := pen.NewJoints(", ", len(cattrs.Parameters.Parameters)+1)
		if cattrs.Parameters.InstanceParameter != nil {
			joints.Addf("%s v", rawCType(cattrs.Parameters.InstanceParameter.AnyType))
		}
		for i, param := range cattrs.Parameters.Parameters {
			joints.Addf("%s _%d", rawCType(param.AnyType), i)
		}
		args = joints.Join()
	}

	return fmt.Sprintf("%s %s(%s)", ret, cattrs.CIdentifier, args)
}

func rawCType(any gir.AnyType) string {
	switch {
	case any.Array != nil:
		return any.Array.CType
	case any.Type != nil:
		return any.Type.CType
	default:
		return "..." // possibly variadic?
	}
}

// StubFn generates a stub C version that is available if the C library's
// version is older than a certain version. This is a crude way to allow users
// with older libraries to compile an application that was developed on a newer
// library but doesn't use any of the new features.
//
// True is returned if the stub is successfully generated. The result of this is
// usually constant within the same namespace.
func (h *Header) StubFn(source *gir.NamespaceFindResult, cattrs *gir.CallableAttrs) bool {
	if cattrs.Version == "" {
		return true
	}

	if h.checkVersion(source, cattrs) {
		h.DashImport(ImportCore("glib"))
		return true
	}

	return false
}

func (h *Header) addCgoStub(stub string, ord cgoHeaderOrder) {
	if h.stop {
		return
	}

	if h.CgoStubs == nil {
		h.CgoStubs = map[string]cgoHeaderOrder{}
	}

	h.CgoStubs[stub] = ord
}

const checkVersionTmpl = `
#if (<.has_major> <"<"> <.major> || (<.has_major> == <.major> && <.has_minor> <"<"> <.minor>))
<.stub> {
	goPanic("<.function>: library too old: needs at least <.major>.<.minor>");
}
#endif
`

func (h *Header) checkVersion(source *gir.NamespaceFindResult, cattrs *gir.CallableAttrs) bool {
	major := constMacro(source, "MAJOR_VERSION")
	minor := constMacro(source, "MINOR_VERSION")
	if major == "" || minor == "" {
		return false
	}

	parts := strings.Split(cattrs.Version, ".")
	if len(parts) < 2 {
		return false
	}

	m := gotmpl.M{
		"has_major": major,
		"has_minor": minor,
		"stub":      StubCFuncHeader(cattrs),
		"major":     parts[0],
		"minor":     parts[1],
		"function":  cattrs.CIdentifier,
	}

	var out strings.Builder
	gotmpl.Render(&out, checkVersionTmpl, m)

	h.addCgoStub("extern void goPanic(const char*);", cgoExterns)
	h.addCgoStub(out.String(), cgoExtras)

	return true
}

func constMacro(source *gir.NamespaceFindResult, name string) string {
	for _, macro := range source.Namespace.Constants {
		if macro.Name == name {
			return macro.CType
		}
	}
	return ""
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

	for path, alias := range h.Imports {
		dst.ImportAlias(path, alias)
	}
	for header, order := range h.CgoStubs {
		dst.addCgoStub(header, order)
	}
	for header, order := range h.CgoHeader {
		dst.addCgoHeader(header, order)
	}
	for pkg := range h.Packages {
		dst.AddPackage(pkg)
	}
}
