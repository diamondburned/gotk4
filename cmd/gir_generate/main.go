package main

import (
	"log"

	"github.com/diamondburned/gotk4/gir"
)

type Package struct {
	Name       string
	Namespaces []string
}

// DBus version 1.0
// DBusGLib version 1.0
// GIRepository version 2.0
// GL version 1.0
// GLib version 2.0
// GModule version 2.0
// GObject version 2.0
// Gio version 2.0
// Vulkan version 1.0
// cairo version 1.0
// fontconfig version 2.0
// freetype2 version 2.0
// libxml2 version 2.0
// win32 version 1.0
// xfixes version 4.0
// xft version 2.0
// xlib version 2.0
// xrandr version 1.3

var packages = []Package{
	{"gobject-introspection-1.0", []string{"GLib", "Gio", "Vulkan", "cairo"}},
	{"pango", nil},
	{"gtk4", nil},
}

func main() {
	var repos gir.Repositories
	var err error

	for _, pkg := range packages {
		if pkg.Namespaces != nil {
			err = repos.AddSelected(pkg.Name, pkg.Namespaces)
		} else {
			err = repos.Add(pkg.Name)
		}

		if err != nil {
			log.Fatalln("error adding packages:", err)
		}
	}

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			log.Println("got namespace", namespace.Name, "version", namespace.Version)
		}
	}
}
