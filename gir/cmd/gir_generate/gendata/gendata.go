// Package gendata contains data used to generate GTK4 bindings for Go. It
// exists primarily to be used externally.
package gendata

import (
	"errors"
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	. "github.com/diamondburned/gotk4/gir/girgen/types"
	. "github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

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
// Not included: externglib (gotk3/gotk3/glib).
var ImportOverrides = map[string]string{}

// Packages lists pkg-config packages and optionally the namespaces to be
// generated. If the list of namespaces is nil, then everything is generated.
var Packages = []Package{
	{"gobject-introspection-1.0", []string{
		"GLib-2",
		"GObject-2",
		"Gio-2",
		"cairo-1",
	}},
	{"gdk-pixbuf-2.0", nil},
	{"graphene-1.0", nil},
	{"atk", nil},
	{"pango", []string{
		"Pango-1",
		"PangoCairo-1",
	}},
	{"gtk4", nil},     // includes Gdk
	{"gtk+-3.0", nil}, // includes Gdk
}

// Preprocessors defines a list of preprocessors that the main generator will
// use. It's mostly used for renaming colliding types/identifiers.
var Preprocessors = []Preprocessor{
	// Collision due to case conversions.
	TypeRenamer("GLib-2.file_test", "test_file"),
	// This collides with Native().
	TypeRenamer("Gtk-4.Native", "NativeSurface"),
	// These collide with structs of the same names.
	RenameEnumMembers("Pango-1.AttrType", "ATTR_(.*)", "ATTR_TYPE_$1"),
	RenameEnumMembers("Gsk-4.RenderNodeType", ".*", "${0}_TYPE"),
	RenameEnumMembers("Gdk-3.EventType", ".*", "${0}_TYPE"),
	// See #28.
	RemoveCIncludes("Gio-2.0.gir", "gio/gdesktopappinfo.h"),

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

	modifyBufferInsert("Gtk-4.TextBuffer.insert"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_markup"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_at_cursor"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_interactive"),
	modifyBufferInsert("Gtk-4.TextBuffer.insert_interactive_at_cursor"),
	modifyBufferInsert("Gtk-4.TextBuffer.set_text"),
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
		log.Printf("doc %#v\n", param)
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

	// These are not in gotk3/cairo.
	AbsoluteFilter("cairo.ScaledFont"),
	AbsoluteFilter("cairo.FontType"),

	// Broadway is not included, so we don't generate code for it.
	FileFilter("gsk/broadway/gskbroadwayrenderer.h"),
	// Output buffer parameter is not actually array.
	AbsoluteFilter("GLib.unichar_to_utf8"),
	// This is useless.
	AbsoluteFilter("GLib.nullify_pointer"),
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

	FileFilter("gasyncqueue."),
	FileFilter("gatomic."),
	FileFilter("gbacktrace."),
	FileFilter("gbase64."),
	FileFilter("gbitlock."),
	FileFilter("gdataset."),
	FileFilter("gdate."),
	FileFilter("gerror."), // already handled internally
	FileFilter("ghook."),
	FileFilter("glib-unix."),
	FileFilter("glist."),
	FileFilter("gmacros."),
	FileFilter("gmem."),
	FileFilter("gnetworking."), // needs header
	FileFilter("gprintf."),
	FileFilter("grcbox."),
	FileFilter("grefcount."),
	FileFilter("grefstring."),
	FileFilter("gsettingsbackend."),
	FileFilter("gslice."),
	FileFilter("gslist."),
	FileFilter("gstdio."),
	FileFilter("gstrfuncs."),
	FileFilter("gstringchunk."),
	FileFilter("gstring."),
	FileFilter("gstrvbuilder."),
	FileFilter("gtestutils."),
	FileFilter("gthread."),
	FileFilter("gthreadpool."),
	FileFilter("gtrashstack."),

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

// GLibDateTime generates NewTimeZoneFromGo and NewDateTimeFromGo.
func GLibDateTime(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gdatetime.go"]
	if !ok {
		return nil
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

			return NewDateTime(tz, Y, int(M), D, h, m, seconds.Seconds())
		}
	`)

	return nil
}

// GioArrayUseBytes is the postprocessor that adds gio/v2.UseBytes.
func GioArrayUseBytes(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["garray.go"]
	if !ok {
		return nil
	}

	h := fg.Header()
	h.Import("runtime")
	h.ImportCore("gbox")
	h.ImportCore("gextras")
	h.CallbackDelete = true

	// We can use the gbox.Assign API for this. The type doesn't matter much,
	// since we're not actually going to access the data through it.

	p := fg.Pen()
	p.Line(`
		// UseBytes is similar to NewBytes, except the given Go byte slice is
		// not copied, but will be kept alive for the lifetime of the GBytes.
		// Note that the user must NOT modify data.
		//
		// Refer to g_bytes_new_with_free_func() for more information.
		func UseBytes(data []byte) *Bytes {
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
	`)

	return nil
}

// GLibAliases generates aliases in the glib/v2 package to the core/glib
// package. It is generated so that users don't have to import both glib
// packages.
func GLibAliases(nsgen *girgen.NamespaceGenerator) error {
	fg := nsgen.MakeFile("externglib.go")

	h := fg.Header()
	h.NeedsExternGLib()

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
	}

	p := fg.Pen()

	for _, fn := range fns {
		names := make([]string, len(fn.Params))
		for i := range fn.Params {
			names[i] = strings.Split(fn.Params[i], " ")[0]
		}

		p.Linef("// %s is an alias for pkg/core/glib.%[1]s.", fn.Name)
		p.Linef("func %s(%s) %s {", fn.Name, strings.Join(fn.Params, ", "), fn.Return)
		p.Linef("  return externglib.%s(%s)", fn.Name, strings.Join(names, ", "))
		p.Linef("}")
	}

	// TODO: right now, we have both externglib.Variant and glib.Variant.
	// externglib's implementation is more idiomatic and clean, but glib's
	// generated implementation is more faithful.
	//
	// For now, we'll keep the generated implementation, since it appears more
	// complete, but in the future, if there are too many incorrect methods that
	// users may fall for, then it's better to switch to externglib.

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
		p.Linef("type %s = externglib.%[1]s", t)
	}

	return nil
}

// GLibLogs adds the following g_log_* functions:
//
//  - g_log_set_handler
//  - g_log_set_handler_full
//
func GLibLogs(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gmessages.go"]
	if !ok {
		return errors.New("missing file gmessages.go")
	}

	h := fg.Header()
	h.Import("os")
	h.Import("log")
	h.Import("strings")
	h.ImportCore("gbox")
	h.CallbackDelete = true

	n := nsgen.Namespace()
	r := nsgen.Repositories()

	h.AddCallback(n, r.FindFullType("GLib-2.LogFunc").Type.(*gir.Callback))
	h.AddCallback(n, r.FindFullType("GLib-2.LogWriterFunc").Type.(*gir.Callback))

	p := fg.Pen()
	p.Line(`
		// LogSetHandler sets the handler used for GLib logging and returns the
		// new handler ID. It is a wrapper around g_log_set_handler and
		// g_log_set_handler_full.
		//
		// To detach a log handler, use LogRemoveHandler.
		func LogSetHandler(domain string, level LogLevelFlags, f LogFunc) uint {
			var log_domain *C.gchar
			if domain != "" {
				log_domain = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
				defer C.free(unsafe.Pointer(log_domain))
			}

			data := gbox.Assign(f)

			h := C.g_log_set_handler_full(
				log_domain,
				C.GLogLevelFlags(level), 
				C.GLogFunc((*[0]byte)(C._gotk4_glib2_LogFunc)),
				C.gpointer(data),
				C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
			)

			return uint(h)
		}

		// Value returns the field's value.
		func (l *LogField) Value() string {
			if l.native.length == -1 {
				return C.GoString((*C.gchar)(unsafe.Pointer(l.native.value)))
			}
			return C.GoStringN((*C.gchar)(unsafe.Pointer(l.native.value)), C.int(l.native.length))
		}

		// LogSetWriter sets the log writer to the given callback, which should
		// take in a list of pair of key-value strings and return true if the
		// log has been successfully written. It is a wrapper around
		// g_log_set_writer_func.
		//
		// Note that this function must ONLY be called ONCE. The GLib
		// documentation states that it is an error to call it more than once.
		func LogSetWriter(f LogWriterFunc) {
			data := gbox.Assign(f)
			C.g_log_set_writer_func(
				C.GLogWriterFunc((*[0]byte)(C._gotk4_glib2_LogWriterFunc)),
				C.gpointer(data),
				C.GDestroyNotify((*[0]byte)(C.callbackDelete)),
			)
		}

		// LogUseDefaultLogger calls LogUseLogger with Go's default standard
		// logger. It is a convenient function for log.Default().
		func LogUseDefaultLogger() { LogUseLogger(log.Default()) }

		// LogUseLogger calls LogSetWriter with the given Go's standard logger.
		// Note that either this or LogSetWriter must only be called once.
		// The method will ignore all fields excet for "MESSAGE"; for more
		// sophisticated, structured log writing, use LogSetWriter.
		// The output format of the logs printed using this function is not
		// guaranteed to not change. Users who rely on the format are better off
		// using LogSetWriter.
		func LogUseLogger(l *log.Logger) { LogSetWriter(LoggerHandler(l)) }

		// LoggerHandler creates a new LogWriterFunc that LogUseLogger uses. For
		// more information, see LogUseLogger's documentation.
		func LoggerHandler(l *log.Logger) LogWriterFunc {
			// Treat Lshortfile and Llongfile the same, because we don't have
			// the full path in codeFile anyway.
			Lfile := l.Flags() & (log.Lshortfile | log.Llongfile) != 0

			// Support $G_MESSAGES_DEBUG.
			debugDomains := make(map[string]struct{})
			for _, debugDomain := range strings.Fields(os.Getenv("G_MESSAGES_DEBUG")) {
				debugDomains[debugDomain] = struct{}{}
			}

			// Special case: G_MESSAGES_DEBUG=all.
			_, debugAll := debugDomains["all"]

			return func(lvl LogLevelFlags, fields []LogField) LogWriterOutput {
				var message, codeFile, codeLine, codeFunc string
				domain := "GLib (no domain)"

				for _, field := range fields {
					if !Lfile {
						switch field.Key() {
						case "MESSAGE":     message  = field.Value()
						case "GLIB_DOMAIN": domain   = field.Value()
						}
						// Skip setting code* if we don't have to.
						continue
					}

					switch field.Key() {
					case "MESSAGE":     message  = field.Value()
					case "CODE_FILE":   codeFile = field.Value()
					case "CODE_LINE":   codeLine = field.Value()
					case "CODE_FUNC":   codeFunc = field.Value()
					case "GLIB_DOMAIN": domain   = field.Value()
					}
				}

				if !debugAll && (lvl&LogLevelDebug != 0) && domain != "" {
					if _, ok := debugDomains[domain]; !ok {
						return LogWriterHandled
					}
				}

				f := l.Printf
				if lvl&LogFlagFatal != 0 {
					f = l.Fatalf
				}

				if !Lfile || (codeFile == "" && codeLine == "") {
					f("%s: %s: %s", lvl, domain, message)
					return LogWriterHandled
				}

				if codeFunc == "" {
					f("%s: %s: %s:%s: %s", lvl, domain, codeFile, codeLine, message)
					return LogWriterHandled
				}

				f("%s: %s: %s:%s (%s): %s", lvl, domain, codeFile, codeLine, codeFunc, message)
				return LogWriterHandled
			}
		}
	`)

	return nil
}

func GtkTextBufferInsert(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gtktextbuffer.go"]
	if !ok {
		return errors.New("missing file gmessages.go")
	}

	h := fg.Header()
	h.Import("reflect")
	h.Import("runtime")
	h.Import("unsafe")

	p := fg.Pen()
	p.Line(`
		// Insert is a convenient method around InsertBytes.
		func (buffer *TextBuffer) Insert(iter *TextIter, str string) {
			var bytes []byte
			bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
			bytesHeader.Len = len(str)
			bytesHeader.Cap = len(str)
			bytesHeader.Data = (*reflect.StringHeader)(unsafe.Pointer(&str)).Data

			buffer.InsertBytes(iter, bytes)
			runtime.KeepAlive(str)
		}
	`)

	return nil
}

const cGTKDialogNew2 = `
	GtkWidget* _gotk4_gtk_dialog_new2(const gchar* title, GtkWindow* parent, GtkDialogFlags flags) {
		return gtk_dialog_new_with_buttons(title, parent, flags, NULL, NULL);
	}
`

func GtkNewDialog(nsgen *girgen.NamespaceGenerator) error {
	fg, ok := nsgen.Files["gtkdialog.go"]
	if !ok {
		return nil
	}

	h := fg.Header()
	h.Import("unsafe")
	h.Import("runtime")
	h.AddCBlock(cGTKDialogNew2)

	p := fg.Pen()
	p.Line(`
		// NewDialogWithFlags is a slightly more advanced version of NewDialog,
		// allowing the user to construct a new dialog with the given
		// constructor-only dialog flags.
		//
		// It is a wrapper around Gtk.Dialog.new_with_buttons in C.
		func NewDialogWithFlags(title string, parent *Window, flags DialogFlags) *Dialog {
			ctitle := C.CString(title)
			defer C.free(unsafe.Pointer(ctitle))

			w := C._gotk4_gtk_dialog_new2(
				(*C.gchar)(unsafe.Pointer(ctitle)),
				(*C.GtkWindow)(unsafe.Pointer(parent.Native())),
				(C.GtkDialogFlags)(flags),
			)
			runtime.KeepAlive(parent)

			return wrapDialog(externglib.Take(unsafe.Pointer(w)))
		}
	`)

	return nil
}

// Postprocessors is similar to Append, except the caller can mutate the package
// in a more flexible manner.
var Postprocessors = map[string][]girgen.Postprocessor{
	"GLib-2": {ImportGError, GioArrayUseBytes, GLibAliases, GLibLogs, GLibDateTime},
	"Gio-2":  {ImportGError},
	"Gtk-3":  {ImportGError, GtkNewDialog},
	"Gtk-4":  {ImportGError, GtkNewDialog},
}

// Appends contains the contents of files that are appended into generated
// outputs. It is used to add custom implementations of missing functions.
var Appends = map[string]string{
	"gtk/v3/gtk.go": `
		// Init binds to the gtk_init() function. Argument parsing is not
		// supported.
		func Init() {
			C.gtk_init(nil, nil)
		}
	`,
}
