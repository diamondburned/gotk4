package main

import "github.com/diamondburned/gotk4/gir/girgen"

// packages lists pkg-config packages and optionally the namespaces to be
// generated. If the list of namespaces is nil, then everything is generated.
var packages = []Package{
	{"gobject-introspection-1.0", []string{
		"GLib", "Gio", "cairo", "xft", "xlib", "freetype2", "fontconfig",
	}},
	{"gdk-pixbuf-2.0", []string{"GdkPixbuf", "GdkPixdata"}},
	{"graphene-1.0", nil},
	{"pango", nil},
	{"gtk4", nil}, // includes Gdk
}

// filters defines a list of GIR types to be filtered. The map key is the
// namespace, and the values are list of names.
var filters = []girgen.FilterMatcher{
	girgen.RegexFilter("GLib.str.+"),
	girgen.RegexFilter("GLib.Variant.*"),
}
