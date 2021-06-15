package main

import "github.com/diamondburned/gotk4/gir/girgen"

// packages lists pkg-config packages and optionally the namespaces to be
// generated. If the list of namespaces is nil, then everything is generated.
var packages = []Package{
	{"gobject-introspection-1.0", []string{
		"GLib-2.0",
		"GObject-2.0",
		"Gio-2.0",
		"cairo-1.0",
	}},
	{"gdk-pixbuf-2.0", nil},
	{"graphene-1.0", nil},
	{"pango", nil},
	{"gtk4", nil},     // includes Gdk
	{"gtk+-3.0", nil}, // includes Gdk
}

// filters defines a list of GIR types to be filtered. The map key is the
// namespace, and the values are list of names.
var filters = []girgen.FilterMatcher{
	girgen.AbsoluteFilter("C.cairo_image_surface_create"),

	// Broadway is not included, so we don't generate code for it.
	girgen.FileFilter("../gsk/broadway/gskbroadwayrenderer.h"),

	// Output buffer parameter is not actually array.
	girgen.AbsoluteFilter("GLib.unichar_to_utf8"),

	// Collision due to case conversions.
	girgen.TypeRenamer("GLib.file_test", "test_file"),

	girgen.FileFilter("garray.h"),
	girgen.FileFilter("gasyncqueue.h"),
	girgen.FileFilter("gatomic.h"),
	girgen.FileFilter("gbacktrace.h"),
	girgen.FileFilter("gbitlock.h"),
	girgen.FileFilter("gbytes.h"),
	girgen.FileFilter("gdataset.h"),
	girgen.FileFilter("gdate.h"),
	girgen.FileFilter("gdatetime.h"),
	girgen.FileFilter("gerror.h"), // already handled internally
	girgen.FileFilter("ghook.h"),
	girgen.FileFilter("glib-unix.h"),
	girgen.FileFilter("glist.h"),
	girgen.FileFilter("gmacros.h"),
	girgen.FileFilter("gmem.h"),
	girgen.FileFilter("gprintf.h"),
	girgen.FileFilter("grcbox.h"),
	girgen.FileFilter("grefcount.h"),
	girgen.FileFilter("grefstring.h"),
	girgen.FileFilter("gslice.h"),
	girgen.FileFilter("gslist.h"),
	girgen.FileFilter("gstdio.h"),
	girgen.FileFilter("gstrfuncs.h"),
	girgen.FileFilter("gstringchunk.h"),
	girgen.FileFilter("gstring.h"),
	girgen.FileFilter("gstrvbuilder.h"),
	girgen.FileFilter("gtestutils.h"),
	girgen.FileFilter("gthread.h"),
	girgen.FileFilter("gthreadpool.h"),
	girgen.FileFilter("gtrashstack.h"),

	// These are missing on build for some reason.
	girgen.AbsoluteFilter("C.g_array_get_type"),
	girgen.AbsoluteFilter("C.g_byte_array_get_type"),
	girgen.AbsoluteFilter("C.g_bytes_get_type"),
	girgen.AbsoluteFilter("C.g_ptr_array_get_type"),
}
