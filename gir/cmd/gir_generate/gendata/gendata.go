// Package gendata contains data used to generate GTK4 bindings for Go. It
// exists primarily to be used externally.
package gendata

import (
	. "github.com/diamondburned/gotk4/gir/girgen/types"
)

type Package struct {
	PkgName    string   // pkg-config name
	Namespaces []string // refer to ./cmd/gir_namespaces
}

// PkgExceptions contains a list of file names that won't be deleted off of
// pkg/.
var PkgExceptions = []string{
	"core",
	"go.mod",
	"go.sum",
	"LICENSE",
}

// PkgGenerated contains a list of file names that are packages generated using
// the given Packages list. It is manually updated.
var PkgGenerated = []string{
	"atk",
	"cairo",
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
	// CellRendererSpinner (generated interface) collides with an actual
	// CellRendererSpinner class.
	TypeRenamer("Gtk-3.CellRendererSpin", "CellRendererSpinButton"),
	TypeRenamer("Gtk-4.CellRendererSpin", "CellRendererSpinButton"),
	// This collides with Native().
	TypeRenamer("Gtk-4.Native", "NativeSurface"),

	// Fix incorrect parameter direction.
	// ModifyCallable("Gio-2.DBusInterfaceGetPropertyFunc",
	// 	func(cattrs *gir.CallableAttrs) {
	// 		FindParameter(cattrs, "error").Direction = "out"
	// 	},
	// ),
}

// Filters defines a list of GIR types to be filtered. The map key is the
// namespace, and the values are list of names.
var Filters = []FilterMatcher{
	AbsoluteFilter("C.cairo_image_surface_create"),

	// Broadway is not included, so we don't generate code for it.
	FileFilter("gsk/broadway/gskbroadwayrenderer.h"),
	// Output buffer parameter is not actually array.
	AbsoluteFilter("GLib.unichar_to_utf8"),
	// This is useless.
	AbsoluteFilter("GLib.nullify_pointer"),
	// Requires special header, is optional function.
	AbsoluteFilter("Gio.networking_init"),
	// Not an array type but expects an array.
	AbsoluteFilter("Gio.SimpleProxyResolver.set_ignore_hosts"),
	// These are not found.
	AbsoluteFilter("C.GdkPixbufModule"),
	AbsoluteFilter("GdkPixbuf.PixbufNonAnim"),
	AbsoluteFilter("GdkPixbuf.PixbufModulePattern"),

	FileFilter("garray."),
	FileFilter("gasyncqueue."),
	FileFilter("gatomic."),
	FileFilter("gbacktrace."),
	FileFilter("gbase64."),
	FileFilter("gbitlock."),
	FileFilter("gbytes."),
	FileFilter("gdataset."),
	FileFilter("gdate."),
	FileFilter("gdatetime."),
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
}
