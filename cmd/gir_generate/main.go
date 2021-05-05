package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/goimports"
)

var (
	output string
)

func init() {
	flag.StringVar(&output, "o", "./", "output directory to mkdir in")
	flag.Parse()
}

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

	var wg sync.WaitGroup

	gen := girgen.NewGenerator(repos)
	gen.WithLogger(log.New(os.Stderr, "[girgen]", log.LstdFlags))

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			ng := gen.UseNamespace(namespace.Name)

			wg.Add(1)
			go func() {
				writeNamespace(ng)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}

func writeNamespace(ng *girgen.NamespaceGenerator) {
	pkg := ng.PackageName()
	dir := filepath.Join(output, pkg)
	out := filepath.Join(dir, pkg+".go")

	if err := os.Mkdir(dir, os.ModePerm|os.ModeDir); err != nil {
		if !os.IsExist(err) {
			log.Println("mkdir failed:", err)
			return
		}
	}

	f, err := os.Create(out)
	if err != nil {
		log.Println("failed to create go file:", err)
		return
	}
	defer f.Close()

	if err := ng.Generate(f); err != nil {
		log.Println("generation error:", err)
		return
	}

	if err := f.Close(); err != nil {
		log.Println("failed to close file:", err)
		return
	}

	if err := goimports.File(out); err != nil {
		log.Println("failed to run goimports on output:", err)
	}
}
