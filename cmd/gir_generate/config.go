package main

import "github.com/diamondburned/gotk4/gir/girgen"

// packages lists pkg-config packages and optionally the namespaces to be
// generated. If the list of namespaces is nil, then everything is generated.
var packages = []Package{
	{"gobject-introspection-1.0", []string{
		"GLib-2.0",
		"GObject-2.0",
		"GModule-2.0",
		"Gio-2.0",
		"cairo-1.0",
		"DBus-1.0",
		"DBusGLib-1.0",
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
	girgen.RegexFilter("GLib.str.+"),
	girgen.AbsoluteFilter("C.cairo_image_surface_create"),

	// girgen.RegexFilter("GLib.[Vv]ariant.*"),
}
