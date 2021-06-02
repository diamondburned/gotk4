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

	girgen.FileFilter("gdate.h"),
	girgen.FileFilter("glib-unix.h"),

	// These are missing on build for some reason.
	girgen.AbsoluteFilter("C.g_array_get_type"),
	girgen.AbsoluteFilter("C.g_byte_array_get_type"),
	girgen.AbsoluteFilter("C.g_bytes_get_type"),
	girgen.AbsoluteFilter("C.g_ptr_array_get_type"),
}
