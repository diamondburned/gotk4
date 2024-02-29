// Package gendata contains data used to generate GTK4 bindings for Go. It
// exists primarily to be used externally.
package gendata

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	. "github.com/diamondburned/gotk4/gir/girgen/types"
	. "github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

const Module = "github.com/diamondburned/gotk4/pkg"

// Main contains all of gendata's variables as one big genmain.Data value.
//
// This stays ugly just because it's the main gotk4 package with exposed
// gendata. Don't actually do this; just make a global genmain.Data instead.
var Main = genmain.Data{
	Module:                Module,
	Packages:              Packages,
	ImportOverrides:       ImportOverrides,
	PkgExceptions:         PkgExceptions,
	GenerateExceptions:    GenerateExceptions,
	PkgGenerated:          PkgGenerated,
	Preprocessors:         Preprocessors,
	Postprocessors:        Postprocessors,
	ExtraGoContents:       ExtraGoContents,
	Filters:               Filters,
	ProcessConverters:     ConversionProcessors,
	DynamicLinkNamespaces: DynamicLinkNamespaces,
	SingleFile:            true,
}

type Package struct {
	PkgName    string   // pkg-config name
	Namespaces []string // refer to ./cmd/gir_namespaces
}

// HasNamespace returns true if the package allows all namespaces or has the
// given namespace in the list.
func (pkg *Package) HasNamespace(n *gir.Namespace) bool {
	if pkg.Namespaces == nil {
		return true
	}

	namespace := gir.VersionedNamespace(n)
	for _, name := range pkg.Namespaces {
		if name == namespace {
			return true
		}
	}

	return false
}

// PkgExceptions contains a list of file names that won't be deleted off of
// pkg/.
var PkgExceptions = []string{
	"core",
	"cairo",
	"go.mod",
	"go.sum",
	"LICENSE",
}

// PkgGenerated contains a list of file names that are packages generated using
// the given Packages list. It is manually updated.
var PkgGenerated = []string{
	"atk",
	"gdk",
	"gdkpixbuf",
	"gdkpixdata",
	"gdkwayland",
	"gdkx11",
	"gio",
	"glib",
	"gobject",
	"graphene",
	"gsk",
	"gtk",
	"pango",
	"pangocairo",
}

// GenerateExceptions contains the keys of the underneath ImportOverrides map.
var GenerateExceptions = []string{
	"cairo-1",
}

// ImportOverrides is the list of imports to defer to another library, usually
// because it's tedious or impossible to generate.
//
// Not included: coreglib (gotk3/gotk3/glib).
var ImportOverrides = map[string]string{}

// Packages lists pkg-config packages and optionally the namespaces to be
// generated. If the list of namespaces is nil, then everything is generated.
var Packages = []genmain.Package{
	{Name: "gobject-introspection-1.0", Namespaces: []string{
		"GLib-2",
		"GObject-2",
		"Gio-2",
		"cairo-1",
	}},
	{Name: "gdk-pixbuf-2.0"},
	{Name: "graphene-1.0"},
	{Name: "atk"},
	{Name: "pango", Namespaces: []string{
		"Pango-1",
		"PangoCairo-1",
	}},
	{Name: "gtk4"},     // includes Gdk
	{Name: "gtk+-3.0"}, // includes Gdk
}

// DynamicLinkNamespaces lists namespaces that should be generated directly
// using Cgo. It includes important core packages as well as packages that are
// small but performance-sensitive.
var DynamicLinkNamespaces = []string{
	"GLib-2",
	"GObject-2",
	"Graphene-1",
	"GdkPixbuf-2",
}

// Preprocessors defines a list of preprocessors that the main generator will
// use. It's mostly used for renaming colliding types/identifiers.
var Preprocessors = []Preprocessor{
	// Collision due to case conversions.
	TypeRenamer("GLib-2.file_test", "test_file"),
	// This collides with Native().
	TypeRenamer("Gtk-4.Native", "NativeSurface"),
	// This collides with Editable()
	TypeRenamer("Gtk-4.Editable", "EditableTextWidget"),
	// These collide with structs of the same names.
	RenameEnumMembers("Pango-1.AttrType", "ATTR_(.*)", "ATTR_TYPE_$1"),
	RenameEnumMembers("Gsk-4.RenderNodeType", ".*", "${0}_TYPE"),
	RenameEnumMembers("Gdk-3.EventType", ".*", "${0}_TYPE"),
	// See #28.
	RemoveCIncludes("Gio-2.0.gir", "gio/gdesktopappinfo.h"),
	// These probably shouldn't be built on Windows.
	RemovePkgconfig("Gio-2.0.gir", "gio-unix-2.0"),
	RemoveCIncludes("Gio-2.0.gir", "gio/gfiledescriptorbased.h", `/gio/gunix.*\.h/`),

	// Length and Value are invalid in Go. We manually handle them in GLibLogs.
	RemoveRecordFields("GLib-2.LogField", "length", "value"),

	ModifyParamDirections("Gio-2.InputStream.read", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.InputStream.read_async", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.InputStream.read_all", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.InputStream.read_all_async", map[string]string{
		"buffer": "in",
		"count":  "in",
	}),
	ModifyParamDirections("Gio-2.Socket.receive", map[string]string{
		"buffer": "in",
		"size":   "in",
	}),
	ModifyParamDirections("Gio-2.Socket.receive_from", map[string]string{
		"buffer": "in",
		"size":   "in",
	}),
	ModifyParamDirections("Gio-2.Socket.receive_with_blocking", map[string]string{
		"buffer": "in",
		"size":   "in",
	}),
	ModifyParamDirections("Gio-2.DBusInterfaceGetPropertyFunc", map[string]string{
		"error": "out",
	}),

	ModifyCallable("Gdk-4.Clipboard.read_async", func(c *gir.CallableAttrs) {
		// Fix this parameter's type not being a proper array.
		p := FindParameter(c, "mime_types")
		p.Array = &gir.Array{
			CType: "const char**",
			Type:  &gir.Type{Name: "utf8"},
		}
	}),

	// These are not introspectable for some reason, even though their
	// signatures look correct.
	MustIntrospect("Gdk-4.Clipboard.set_text"),
	MustIntrospect("Gdk-4.Clipboard.set_texture"),

	// Fix up the return array type for (*Variant).String().
	ModifyCallable("GLib-2.Variant.get_string", func(c *gir.CallableAttrs) {
		c.ReturnValue.Array = &gir.Array{
			CType:          "const gchar*",
			Type:           &gir.Type{Name: "gchar"},
			Length:         new(int),  // 0
			ZeroTerminated: new(bool), // false
		}
	}),

	// Fix up Application::open's File type. It's supposed to be a GFile** from
	// the source code, but that's missing from the GIR data.
	ModifySignal("Gio-2.Application::open", func(sig *gir.Signal) {
		param := FindParameterFromSlice(sig.Parameters.Parameters, "files")
		param.Array.CType = "GFile**"
	}),

	// Fix up GVariant methods to have nullable returns.
	PreprocessorFunc(func(repos gir.Repositories) {
		variant := repos.FindFullType("GLib-2.Variant").Type.(*gir.Record)
		for _, method := range variant.Methods {
			returnsGVariant := true &&
				method.ReturnValue != nil &&
				method.ReturnValue.Type != nil &&
				method.ReturnValue.Type.CType == "GVariant*"

			if returnsGVariant && !method.ReturnValue.Nullable {
				// GVariant pointers can be null.
				method.ReturnValue.Nullable = true
			}
		}
	}),

	modifyBufferInsert("Gtk-4.TextBuffer.insert"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_markup"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_at_cursor"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_interactive"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_interactive_at_cursor"),
	modifyBufferInsert("Gtk-4.TextBuffer.set_text"),

	modifyBufferInsert("Gtk-3.TextBuffer.insert"),
	modifyBufferInsert("Gtk-3.TextBuffer.insert_markup"),
	modifyBufferInsert("Gtk-3.TextBuffer.insert_at_cursor"),
	modifyBufferInsert("Gtk-3.TextBuffer.insert_interactive"),
	modifyBufferInsert("Gtk-3.TextBuffer.insert_interactive_at_cursor"),
	modifyBufferInsert("Gtk-3.TextBuffer.set_text"),
}

func modifyBufferInsert(name string) Preprocessor {
	names := []string{"text", "markup"}

	return ModifyCallable(name, func(c *gir.CallableAttrs) {
		var p *gir.ParameterAttrs

		for _, name := range names {
			if p = FindParameter(c, name); p != nil {
				break
			}
		}

		if p == nil {
			return
		}

		lenIx := findTextLenParam(c.Parameters.Parameters)
		if lenIx == -1 {
			return
		}

		p.Type = nil
		p.Array = &gir.Array{
			CType:          "const char*",
			Type:           &gir.Type{Name: "gchar"},
			Length:         &lenIx,
			ZeroTerminated: new(bool), // false
		}
	})
}

func findTextLenParam(params []gir.Parameter) int {
	const doc = "length of"

	for i, param := range params {
		if param.Doc != nil && strings.Contains(param.Doc.String, doc) {
			return i
		}
	}

	return -1
}

var ConversionProcessors = []ConversionProcessor{
	ProcessCallback("Gio-2.AsyncReadyCallback", func(conv *Converter) {
		// Don't include the first parameter in Go.
		conv.Results[0].Skip = true
	}),
}

// Filters defines a list of GIR types to be filtered. The map key is the
// namespace, and the values are list of names.
var Filters = []FilterMatcher{
	AbsoluteFilter("C.cairo_image_surface_create"),

	// This seems to be macro-guarded between x86 and arm64.
	AbsoluteFilter("GLib.VA_COPY_AS_ARRAY"),

	// These are not in gotk3/cairo.
	AbsoluteFilter("cairo.ScaledFont"),
	AbsoluteFilter("cairo.FontType"),

	// Broadway is not included, so we don't generate code for it.
	FileFilter("gsk/broadway/gskbroadwayrenderer.h"),
	// Output buffer parameter is not actually array.
	AbsoluteFilter("GLib.unichar_to_utf8"),
	// This is useless.
	AbsoluteFilter("GLib.nullify_pointer"),
	// We already alias this from coreglib.
	AbsoluteFilter("GLib.idle_add_full"),
	AbsoluteFilter("GLib.timeout_add_full"),
	AbsoluteFilter("GLib.timeout_add_seconds_full"),
	// We manually wrote these before the code was able to generate them.
	AbsoluteFilter("GLib.log_set_writer_func"),
	AbsoluteFilter("GLib.log_set_handler_full"),
	// Requires special header, is optional function.
	AbsoluteFilter("GLib.unix_error_quark"),
	AbsoluteFilter("Gio.networking_init"),
	// Not an array type but expects an array.
	AbsoluteFilter("Gio.SimpleProxyResolver.set_ignore_hosts"),
	// These are not found.
	AbsoluteFilter("C.GdkPixbufModule"),
	AbsoluteFilter("GdkPixbuf.PixbufNonAnim"),
	AbsoluteFilter("GdkPixbuf.PixbufModulePattern"),
	AbsoluteFilter("GdkPixbuf.PixbufFormat.domain"),
	AbsoluteFilter("GdkPixbuf.PixbufFormat.flags"),
	AbsoluteFilter("GdkPixbuf.PixbufFormat.disabled"),
	// Dangerous.
	AbsoluteFilter("GLib.IOChannel.read"),
	AbsoluteFilter("GLib.Bytes.new_take"),
	AbsoluteFilter("GLib.Bytes.new_static"),
	AbsoluteFilter("GLib.Bytes.unref_to_data"),
	AbsoluteFilter("GLib.Bytes.unref_to_array"),
	// Not available on Windows.
	RegexFilter(`GLib.Source\..*unix.*`),
	RegexFilter(`Gio.Subprocess`),
	RegexFilter(`Gio.SubprocessLauncher`), // useless without Subprocess
	// Nothing "Unix" is going to be available on Windows.
	RegexFilter(`Gio..*[Uu]nix.*`),
	// Type is different across platforms, which the generator isn't prepared
	// for. Use os/exec instead.
	AbsoluteFilter("GLib.Pid"),
	// PollFD needs fd to be int on Unix and syscall.Handle on Windows, We can
	// handle this later. Go doesn't need this.
	AbsoluteFilter("GLib.PollFD"),
	// These are removed in Preprocessors.
	FileFilter("gfiledescriptorbased."),
	FileFilter("gunix"),

	FileFilterNamespace("GLib", "gasyncqueue."),
	FileFilterNamespace("GLib", "gatomic."),
	FileFilterNamespace("GLib", "gbacktrace."),
	FileFilterNamespace("GLib", "gbase64."),
	FileFilterNamespace("GLib", "gbitlock."),
	FileFilterNamespace("GLib", "gdataset."),
	FileFilterNamespace("GLib", "gdate."),
	FileFilterNamespace("GLib", "gerror."), // already handled internally
	FileFilterNamespace("GLib", "ghook."),
	FileFilterNamespace("GLib", "glib-unix."),
	FileFilterNamespace("GLib", "glist."),
	FileFilterNamespace("GLib", "gmacros."),
	FileFilterNamespace("GLib", "gmem."),
	FileFilterNamespace("GLib", "gnetworking."), // needs header
	FileFilterNamespace("GLib", "gprintf."),
	FileFilterNamespace("GLib", "grcbox."),
	FileFilterNamespace("GLib", "grefcount."),
	FileFilterNamespace("GLib", "grefstring."),
	FileFilterNamespace("GLib", "gslice."),
	FileFilterNamespace("GLib", "gslist."),
	FileFilterNamespace("GLib", "gstdio."),
	FileFilterNamespace("GLib", "gstrfuncs."),
	FileFilterNamespace("GLib", "gstringchunk."),
	FileFilterNamespace("GLib", "gstring."),
	FileFilterNamespace("GLib", "gstrvbuilder."),
	FileFilterNamespace("GLib", "gtestutils."),
	FileFilterNamespace("GLib", "gthread."),
	FileFilterNamespace("GLib", "gthreadpool."),
	FileFilterNamespace("GLib", "gtrashstack."),
	FileFilterNamespace("Gio", "gsettingsbackend."),

	// Header-specific.
	FileFilter("gskglrenderer."),
	FileFilter("gsknglrenderer."),
	FileFilter("gskvulkanrenderer."),
	FileFilter("gdesktopappinfo."), // See #28.
	// These are not found in GTK4 for some reason, but we're ignoring it for
	// GTK3 as well.
	FileFilter("gtkpagesetupunixdialog"),
	FileFilter("gtkprintunixdialog"),
	FileFilter("gtkprinter"),
	FileFilter("gtkprintjob"),
	FileFilter("gdkprivate"),

	// These are missing on build for some reason.
	AbsoluteFilter("C.g_array_get_type"),
	AbsoluteFilter("C.g_byte_array_get_type"),
	AbsoluteFilter("C.g_bytes_get_type"),
	AbsoluteFilter("C.g_ptr_array_get_type"),
	AbsoluteFilter("C.gtk_header_bar_accessible_get_type"),
	AbsoluteFilter("C.gdk_pixbuf_non_anim_get_type"),
	AbsoluteFilter("C.gdk_window_destroy_notify"),
	AbsoluteFilter("C.gtk_print_capabilities_get_type"),
	AbsoluteFilter("C.GdkPixbufAnimationClass"),
	AbsoluteFilter("C.GdkPixbufAnimationIterClass"),

	// Missing.
	AbsoluteFilter("Gtk-3.HeaderBarAccessibleClass"),

	// Already handled in GLibAliases.
	AbsoluteFilter("C.g_source_remove"),
}

// ImportGError ensures that gerror is imported.
func ImportGError(nsgen *girgen.NamespaceGenerator) error {
	core := file.ImportCore("gerror")

	for _, f := range nsgen.Files {
		if f.Header().HasImport(core) {
			return nil
		}
	}

	f := nsgen.MakeFile("")
	f.Header().DashImport(core)

	return nil
}

func GLibVariantIter(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gvariant.go"]
	if !ok {
		fg = nsgen.MakeFile("")
	}

	h := fg.Header()
	h.Import("unsafe")
	h.Import("runtime")
	h.ImportCore("gextras")
	h.NeedsExternGLib()
	h.AddMarshaler("coreglib.TypeVariant", "Variant")

	p := fg.Pen()
	p.Line(`
		func marshalVariant(p uintptr) (interface{}, error) {
			_cret := C.g_value_dup_variant((*C.GValue)(unsafe.Pointer(p)))
			if _cret == nil {
				return (*Variant)(nil), nil
			}

			_variant := (*Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(_variant)),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_variant_unref((*C.GVariant)(intern.C))
				},
			)
			return _variant, nil
		}

		// ForEach iterates over items in value. The iteration breaks out once f
		// returns true. This method wraps around g_variant_iter_new.
		func (value *Variant) ForEach(f func(*Variant) (stop bool)) {
			valueNative := (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

			var iter C.GVariantIter
			C.g_variant_iter_init(&iter, valueNative)

			next := func() *Variant {
				item := C.g_variant_iter_next_value(&iter)
				if item == nil {
					return nil
				}

				variant := (*Variant)(gextras.NewStructNative(unsafe.Pointer(item)))
				runtime.SetFinalizer(
					gextras.StructIntern(unsafe.Pointer(variant)),
					func(intern *struct{ C unsafe.Pointer }) {
						C.g_variant_unref((*C.GVariant)(intern.C))
					},
				)

				return variant
			}

			for item := next(); item != nil; item = next() {
				if f(item) {
					break
				}
			}

			runtime.KeepAlive(value)
		}
	`)

	return nil
}

// GLibDateTime generates NewTimeZoneFromGo and NewDateTimeFromGo.
func GLibDateTime(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gdatetime.go"]
	if !ok {
		fg = nsgen.MakeFile("")
	}

	h := fg.Header()
	h.Import("time")

	p := fg.Pen()
	p.Line(`
		// NewTimeZoneFromGo creates a new TimeZone instance from Go's Location.
		// The location's accuracy is down to the second.
		func NewTimeZoneFromGo(loc *time.Location) *TimeZone {
			switch loc {
			case time.UTC:
				return NewTimeZoneUTC()
			case time.Local:
				return NewTimeZoneLocal()
			}

			t1 := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
			t2 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
			return NewTimeZoneOffset(int32(t2.Sub(t1) / time.Second))
		}

		// NewDateTimeFromGo creates a new DateTime instance from Go's Time. The
		// TimeZone of the DateTime will be implicitly converted from the Time.
		func NewDateTimeFromGo(t time.Time) *DateTime {
			tz := NewTimeZoneFromGo(t.Location())

			Y, M, D := t.Date()
			h, m, s := t.Clock()

			// Second offset within a minute in nanoseconds.
			seconds := (time.Duration(s) * time.Second) + time.Duration(t.Nanosecond())

			return NewDateTime(tz, int(Y), int(M), int(D), int(h), int(m), seconds.Seconds())
		}
	`)

	return nil
}

func GLibObjectComparer(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gobjectcomparer.go"]
	if !ok {
		fg = nsgen.MakeFile("")
	}

	h := fg.Header()
	h.NeedsExternGLib()

	p := fg.Pen()
	p.Line(`
		// NewObjectComparer returns a CompareDataFunc that uses the given function to
		// compare two objects of type T. If the underlying objects are not of type T,
		// then the function panics. If the underlying pointers aren't objects, then the
		// behavior is undefined.
		func NewObjectComparer[T Objector](f func(a, b T) int) CompareDataFunc {
			return func(a, b unsafe.Pointer) int {
				var aobj, bobj T
				if a != nil {
					aobj = coreglib.Take(a).Cast().(T)
				}
				if b != nil {
					bobj = coreglib.Take(b).Cast().(T)
				}
				return f(aobj, bobj)
			}
		}
	`)

	return nil
}

// GioArrayUseBytes is the postprocessor that adds gio/v2.UseBytes.
func GioArrayUseBytes(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["garray.go"]
	if !ok {
		fg = nsgen.MakeFile("")
	}

	h := fg.Header()
	h.Import("runtime")
	h.Import("reflect")
	h.ImportCore("gbox")
	h.ImportCore("gextras")
	h.CallbackDelete = true

	// We can use the gbox.Assign API for this. The type doesn't matter much,
	// since we're not actually going to access the data through it.

	p := fg.Pen()
	p.Line(`
		// NewBytesWithGo is similar to NewBytes, except the given Go byte slice
		// is not copied, but will be kept alive for the lifetime of the GBytes.
		// Note that the user must NOT modify data.
		//
		// Refer to g_bytes_new_with_free_func() for more information.
		func NewBytesWithGo(data []byte) *Bytes {
			byteID := gbox.Assign(data)

			v := C.g_bytes_new_with_free_func(
				C.gconstpointer(unsafe.Pointer(&data[0])),
				C.gsize(len(data)),
				C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
				C.gpointer(byteID),
			)

			_bytes := (*Bytes)(gextras.NewStructNative(unsafe.Pointer(v)))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(_bytes)),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_bytes_unref((*C.GBytes)(intern.C))
				},
			)

			return _bytes
		}

		// Use calls f with Bytes' internal byte slice without making a copy. f
		// must NOT move the byte slice to outside of the closure, since the
		// slice's internal array buffer may be freed after.
		func (b *Bytes) Use(f func([]byte)) {
			var ptr C.gconstpointer // in
			var len C.gsize         // in

			ptr = C.g_bytes_get_data(
				(*C.GBytes)(gextras.StructNative(unsafe.Pointer(b))),
				&len,
			)

			var buf []byte

			h := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
			h.Data = uintptr(ptr)
			h.Len = int(len)
			h.Cap = int(len)

			f(buf)
			runtime.KeepAlive(b)
		}
	`)

	return nil
}

// GLibAliases generates aliases in the glib/v2 package to the core/glib
// package. It is generated so that users don't have to import both glib
// packages.
func GLibAliases(nsgen *girgen.NamespaceGenerator) error {
	fg := nsgen.MakeFile("coreglib.go")

	h := fg.Header()
	h.NeedsExternGLib()

	h.Import("unsafe")
	h.ImportCore("gextras")

	// Needed for GValue.
	h.IncludeC("glib-object.h")
	h.AddPackage("gobject-2.0")

	type fn struct {
		Name   string
		Params []string
		Return string
	}

	fns := []fn{
		{"IdleAdd", []string{"f interface{}"}, "SourceHandle"},
		{"IdleAddPriority", []string{"p Priority", "f interface{}"}, "SourceHandle"},
		{"TimeoutAdd", []string{"ms uint", "f interface{}"}, "SourceHandle"},
		{"TimeoutAddPriority", []string{"ms uint", "p Priority", "f interface{}"}, "SourceHandle"},
		{"TimeoutSecondsAdd", []string{"s uint", "f interface{}"}, "SourceHandle"},
		{"TimeoutSecondsAddPriority", []string{"s uint", "p Priority", "f interface{}"}, "SourceHandle"},
		{"TypeFromName", []string{"typeName string"}, "Type"},
		{"NewValue", []string{"v interface{}"}, "*Value"},
		{"SourceRemove", []string{"src SourceHandle"}, "bool"},
		{"ObjectEq", []string{"obj1 Objector", "obj2 Objector"}, "bool"},
		{"BaseObject", []string{"obj Objector"}, "*Object"},
	}

	p := fg.Pen()

	for _, fn := range fns {
		names := make([]string, len(fn.Params))
		for i := range fn.Params {
			names[i] = strings.Split(fn.Params[i], " ")[0]
		}

		p.Linef("// %s is an alias for pkg/core/glib.%[1]s.", fn.Name)
		p.Linef("func %s(%s) %s {", fn.Name, strings.Join(fn.Params, ", "), fn.Return)
		p.Linef("  return coreglib.%s(%s)", fn.Name, strings.Join(names, ", "))
		p.Linef("}")
	}

	// TODO: right now, we have both coreglib.Variant and glib.Variant.
	// coreglib's implementation is more idiomatic and clean, but glib's
	// generated implementation is more faithful.
	//
	// For now, we'll keep the generated implementation, since it appears more
	// complete, but in the future, if there are too many incorrect methods that
	// users may fall for, then it's better to switch to coreglib.

	types := []string{
		"Object",
		"Objector",
		"Type",
		"Value",
		"Priority",
		"SourceHandle",
		"SignalHandle",
	}

	for _, t := range types {
		p.Linef("// %s is an alias for pkg/core/glib.%[1]s.", t)
		p.Linef("type %s = coreglib.%[1]s", t)
	}

	consts := []string{
		"TypeInvalid",
		"TypeNone",
		"TypeInterface",
		"TypeChar",
		"TypeUchar",
		"TypeBoolean",
		"TypeInt",
		"TypeUint",
		"TypeLong",
		"TypeUlong",
		"TypeInt64",
		"TypeUint64",
		"TypeEnum",
		"TypeFlags",
		"TypeFloat",
		"TypeDouble",
		"TypeString",
		"TypePointer",
		"TypeBoxed",
		"TypeParam",
		"TypeObject",
		"TypeVariant",
		"",
		"PriorityHigh",
		"PriorityDefault",
		"PriorityHighIdle",
		"PriorityDefaultIdle",
		"PriorityLow",
	}

	p.Linef("// Constant aliases from pkg/core/glib.")
	p.Linef("const (")
	for _, c := range consts {
		if c == "" {
			p.EmptyLine()
			continue
		}
		p.Linef("%s = coreglib.%[1]s", c)
	}
	p.Linef(")")

	p.Linef("// NewVariantValue creates a new GValue from a GVariant. This function")
	p.Linef("// only exists as a workaround for coreglib's cyclical imports. It")
	p.Linef("// be removed in the future once coreglib is merged in.")
	p.Linef("func NewVariantValue(variant *Variant) *coreglib.Value {")
	p.Linef("  value := coreglib.InitValue(coreglib.TypeVariant)")
	p.Linef("  C.g_value_set_variant(")
	p.Linef("    (*C.GValue)(unsafe.Pointer(value.Native())),")
	p.Linef("    (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant))),")
	p.Linef("  )")
	p.Linef("  return value")
	p.Linef("}")

	return nil
}

// GLibLogs adds the following g_log_* functions:
//
//   - g_log_set_handler
//   - g_log_set_handler_full
func GLibLogs(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.File("gmessages.go")
	if !ok {
		fg = nsgen.MakeFile("")
	}

	h := fg.Header()
	h.Import("os")
	h.Import("context")
	h.Import("log/slog")
	h.Import("strings")
	h.ImportCore("gbox")
	h.CallbackDelete = true

	r := nsgen.Repositories()

	AddCallbackHeader(nsgen, h, r.FindFullType("GLib-2.LogFunc").Type.(*gir.Callback))
	AddCallbackHeader(nsgen, h, r.FindFullType("GLib-2.LogWriterFunc").Type.(*gir.Callback))

	p := fg.Pen()
	p.Line(`
		// Value returns the field's value.
		func (l *LogField) Value() string {
			if l.native.length == -1 {
				return C.GoString((*C.gchar)(unsafe.Pointer(l.native.value)))
			}
			return C.GoStringN((*C.gchar)(unsafe.Pointer(l.native.value)), C.int(l.native.length))
		}

		// logSetWriter sets the log writer to the given callback, which should
		// take in a list of pair of key-value strings and return true if the
		// log has been successfully written. It is a wrapper around
		// g_log_set_writer_func.
		func logSetWriter(f LogWriterFunc) {
			data := gbox.Assign(f)
			C.g_log_set_writer_func(
				C.GLogWriterFunc((*[0]byte)(C._gotk4_glib2_LogWriterFunc)),
				C.gpointer(data),
				C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
			)
		}

		func init() {
			logSetWriter(func(lvl LogLevelFlags, fields []LogField) LogWriterOutput {
				handle := newSlogWriterFunc(slog.Default())
				handle(lvl, fields)
				return LogWriterHandled
			})
		}

		// Support $G_MESSAGES_DEBUG.
		var debugDomains = func() map[string]struct{} {
			debugDomains := make(map[string]struct{})
			for _, debugDomain := range strings.Fields(os.Getenv("G_MESSAGES_DEBUG")) {
				debugDomains[debugDomain] = struct{}{}
			}
			return debugDomains
		}()

		// Special case: G_MESSAGES_DEBUG=all.
		var _, debugAllDomains = debugDomains["all"]

		// newSlogWriterFunc returns a new LogWriterFunc that writes to the given
		// slog.Logger.
		func newSlogWriterFunc(l *slog.Logger) LogWriterFunc {
			return func(lvl LogLevelFlags, fields []LogField) LogWriterOutput {
				attrs := make([]slog.Attr, 0, len(fields))
				var message, domain string

				for _, field := range fields {
					k := field.Key()
					v := field.Value()
					if k == "MESSAGE" {
						message = v
					} else {
						if k == "GLIB_DOMAIN" {
							domain = v
						}
						k = strings.ToLower(k)
						attrs = append(attrs, slog.String(k, v))
					}
				}

				if !debugAllDomains && (lvl&LogLevelDebug != 0) && domain != "" {
					if _, ok := debugDomains[domain]; !ok {
						return LogWriterHandled
					}
				}

				slogLevel := slog.LevelInfo
				switch {
				case lvl.Has(LogLevelError), lvl.Has(LogLevelCritical):
					slogLevel = slog.LevelError
				case lvl.Has(LogLevelWarning):
					slogLevel = slog.LevelWarn
				case lvl.Has(LogLevelMessage), lvl.Has(LogLevelInfo):
					slogLevel = slog.LevelInfo
				case lvl.Has(LogLevelDebug):
					slogLevel = slog.LevelDebug
				}

				l.LogAttrs(context.Background(), slogLevel, message, attrs...)

				if lvl.Has(LogFlagFatal) {
					panic(message)
				}

				return LogWriterHandled
			}
		}
	`)

	return nil
}

func GtkNewDialog(nsgen *girgen.NamespaceGenerator) error {
	name := fmt.Sprintf(
		"_gotk4_gtk%s_dialog_new2",
		gir.MajorVersion(nsgen.Namespace().Namespace.Version),
	)

	fg := nsgen.MakeFile("gtkdialog-new.go")

	h := fg.Header()
	h.Import("unsafe")
	h.Import("runtime")
	h.NeedsExternGLib()

	h.AddCBlock(fmt.Sprintf(`
		GtkWidget* %s(const gchar* title, GtkWindow* parent, GtkDialogFlags flags) {
			return gtk_dialog_new_with_buttons(title, parent, flags, NULL, NULL);
		}`,
		name,
	))

	p := fg.Pen()
	p.Linef(`
		// NewDialogWithFlags is a slightly more advanced version of NewDialog,
		// allowing the user to construct a new dialog with the given
		// constructor-only dialog flags.
		//
		// It is a wrapper around Gtk.Dialog.new_with_buttons in C.
		func NewDialogWithFlags(title string, parent *Window, flags DialogFlags) *Dialog {
			ctitle := C.CString(title)
			defer C.free(unsafe.Pointer(ctitle))

			w := C.%s(
				(*C.gchar)(unsafe.Pointer(ctitle)),
				(*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native())),
				(C.GtkDialogFlags)(flags),
			)
			runtime.KeepAlive(parent)

			return wrapDialog(coreglib.Take(unsafe.Pointer(w)))
		}
	`, name)

	return nil
}

const cGTKMessageDialogNew2 = `
	GtkWidget* _gotk4_gtk_message_dialog_new2(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons) {
		return gtk_message_dialog_new_with_markup(parent, flags, type, buttons, NULL);
	}
`

func GtkNewMessageDialog(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.File("gtkmessagedialog.go")
	if !ok {
		fg = nsgen.MakeFile("")
	}

	h := fg.Header()
	h.Import("unsafe")
	h.Import("runtime")
	h.AddCBlock(cGTKMessageDialogNew2)

	p := fg.Pen()
	p.Line(`
		// NewMessageDialog creates a new message dialog. This is a simple
		// dialog with some text taht the user may want to see. When the user
		// clicks a button, a "response" signal is emitted with response IDs
		// from ResponseType.
		func NewMessageDialog(parent *Window, flags DialogFlags, typ MessageType, buttons ButtonsType) *MessageDialog {
			w := C._gotk4_gtk_message_dialog_new2(
				(*C.GtkWindow)(unsafe.Pointer(coreglib.InternObject(parent).Native())),
				(C.GtkDialogFlags)(flags),
				(C.GtkMessageType)(typ),
				(C.GtkButtonsType)(buttons),
			)
			runtime.KeepAlive(parent)

			return wrapMessageDialog(coreglib.Take(unsafe.Pointer(w)))
		}
	`)

	return nil
}

func GtkLockOSThread(nsgen *girgen.NamespaceGenerator) error {
	// LockOSThread potentially induces additional overhead, so we're limiting
	// it to only the platforms that need it.
	fg := nsgen.MakeFile("gtk_darwin.go")
	fg.Header().Import("runtime")

	p := fg.Pen()
	p.Line(`
		func init() {
			runtime.LockOSThread()
		}
	`)

	return nil
}

func GtkInvalidListItem(nsgen *girgen.NamespaceGenerator) error {
	fg := nsgen.MakeFile("gtk.go")
	fg.Header().Import("math")

	p := fg.Pen()
	p.Line(`
		// InvalidListPosition is the value used to refer to a guaranteed
		// invalid position in a [gio.ListModel].
		//
		// This value may be returned from some functions, others may accept it
		// as input. Its interpretation may differ for different functions.
		//
		// Refer to each function’s documentation for if this value is allowed
		// and what it does.
		const InvalidListPosition = math.MaxUint32
	`)

	return nil
}

func GdkPixbufFromImage(nsgen *girgen.NamespaceGenerator) error {
	fg := nsgen.MakeFile("gdk-pixbuf-core-go.go")
	fg.Header().Import("image")
	fg.Header().Import("image/draw")
	fg.Header().Import(Module + "/glib/v2")
	fg.Header().Import(Module + "/cairo/swizzle")

	p := fg.Pen()
	p.Line(`
		// NewPixbufFromImage creates a new Pixbuf from a stdlib image.Image. It
		// contains a fast path for *image.RGBA while resorting to
		// copying/converting the image otherwise.
		func NewPixbufFromImage(img image.Image) *Pixbuf {
			bounds := img.Bounds()
			var pixbuf *Pixbuf

			switch img := img.(type) {
			case *image.RGBA:
				bytes := glib.NewBytesWithGo(img.Pix)
				pixbuf = NewPixbufFromBytes(bytes, ColorspaceRGB, true, 8, bounds.Dx(), bounds.Dy(), img.Stride)
			default:
				pixbuf = NewPixbuf(ColorspaceRGB, true, 8, bounds.Dx(), bounds.Dy())
				pixbuf.ReadPixelBytes().Use(func(b []byte) {
					// For information on how this works, refer to
					// pkg/cairo/surface_image.go.
					rgba := image.RGBA{
						Pix:    b,
						Stride: bounds.Dx(),
						Rect:   bounds,
					}
					draw.Draw(&rgba, rgba.Rect, img, image.Point{}, draw.Over)
					swizzle.BGRA(rgba.Pix)
				})
			}

			return pixbuf
		}
	`)

	return nil
}

// Postprocessors is similar to Append, except the caller can mutate the package
// in a more flexible manner.
var Postprocessors = map[string][]girgen.Postprocessor{
	"GLib-2":      {ImportGError, GioArrayUseBytes, GLibVariantIter, GLibAliases, GLibLogs, GLibDateTime, GLibObjectComparer},
	"GdkPixbuf-2": {GdkPixbufFromImage},
	"Gio-2":       {ImportGError},
	"Gtk-3":       {ImportGError, GtkNewDialog, GtkNewMessageDialog, GtkLockOSThread},
	"Gtk-4":       {ImportGError, GtkNewDialog, GtkNewMessageDialog, GtkLockOSThread, GtkInvalidListItem},
}

// ExtraGoContents contains the contents of files that are appended into
// generated outputs. It is used to add custom implementations of missing
// functions.
var ExtraGoContents = map[string]string{
	"gtk/v3/gtk.go": `
		// Init binds to the gtk_init() function. Argument parsing is not
		// supported.
		func Init() {
			C.gtk_init(nil, nil)
		}
	`,
}
