package main

import (
	. "github.com/diamondburned/gotk4/gir/girgen/types"
)

// pkgExceptions contains a list of file names that won't be deleted off of
// pkg/.
var pkgExceptions = []string{
	"core",
	"go.mod",
	"go.sum",
	"LICENSE",
}

// packages lists pkg-config packages and optionally the namespaces to be
// generated. If the list of namespaces is nil, then everything is generated.
var packages = []Package{
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

// preprocessors defines a list of preprocessors that the main generator will
// use. It's mostly used for renaming colliding types/identifiers.
var preprocessors = []Preprocessor{
	// Collision due to case conversions.
	TypeRenamer("GLib-2.file_test", "test_file"),
	// Fix incorrect parameter direction.
	// ModifyCallable("Gio-2.DBusInterfaceGetPropertyFunc",
	// 	func(cattrs *gir.CallableAttrs) {
	// 		FindParameter(cattrs, "error").Direction = "out"
	// 	},
	// ),
}

// filters defines a list of GIR types to be filtered. The map key is the
// namespace, and the values are list of names.
var filters = []FilterMatcher{
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

	FileFilter("garray.h"),
	FileFilter("gasyncqueue.h"),
	FileFilter("gatomic.h"),
	FileFilter("gbacktrace.h"),
	FileFilter("gbitlock.h"),
	FileFilter("gbytes.h"),
	FileFilter("gdataset.h"),
	FileFilter("gdate.h"),
	FileFilter("gdatetime.h"),
	FileFilter("gerror.h"), // already handled internally
	FileFilter("ghook.h"),
	FileFilter("glib-unix.h"),
	FileFilter("glist.h"),
	FileFilter("gmacros.h"),
	FileFilter("gmem.h"),
	FileFilter("gnetworking.h"), // needs header
	FileFilter("gprintf.h"),
	FileFilter("grcbox.h"),
	FileFilter("grefcount.h"),
	FileFilter("grefstring.h"),
	FileFilter("gsettingsbackend.h"),
	FileFilter("gslice.h"),
	FileFilter("gslist.h"),
	FileFilter("gstdio.h"),
	FileFilter("gstrfuncs.h"),
	FileFilter("gstringchunk.h"),
	FileFilter("gstring.h"),
	FileFilter("gstrvbuilder.h"),
	FileFilter("gtestutils.h"),
	FileFilter("gthread.h"),
	FileFilter("gthreadpool.h"),
	FileFilter("gtrashstack.h"),

	// Header-specific.
	FileFilter("gskglrenderer.h"),
	FileFilter("gsknglrenderer.h"),
	FileFilter("gskvulkanrenderer.h"),
	// These are not found in GTK4 for some reason, but we're ignoring it for
	// GTK3 as well.
	FileFilter("gtkpagesetupunixdialog.c"),
	FileFilter("gtkpagesetupunixdialog.h"),
	FileFilter("gtkprinter.c"),
	FileFilter("gtkprinter.h"),

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
