package main

import (
	"log"

	"github.com/diamondburned/gotk4/gir"
)

var packages = []string{"gobject-introspection-1.0", "glib-2.0", "pango", "gtk4"}

func main() {
	var repos gir.Repositories

	if err := repos.AddAll(packages...); err != nil {
		log.Fatalln("error adding packages:", err)
	}

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			log.Println("got namespace", namespace.Name, "version", namespace.Version)
		}
	}
}
