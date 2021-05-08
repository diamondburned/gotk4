package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/goimports"
)

var (
	output string
)

func init() {
	flag.StringVar(&output, "o", "", "output directory to mkdir in")
	flag.Parse()

	if output == "" {
		log.Fatalln("Missing -o output directory.")
	}
}

type Package struct {
	PkgName    string   // pkg-config name
	Namespaces []string // refer to ./cmd/gir_namespaces
}

var packages = []Package{
	{"gobject-introspection-1.0", []string{
		"GLib", "Gio", "Vulkan", "cairo", "xft", "xlib", "freetype2",
	}},
	{"pango", nil},
	{"gdk-pixbuf-2.0", []string{"GdkPixbuf", "GdkPixdata"}},
	{"gdk-wayland-3.0", []string{"Gdk", "GdkX11"}},
	{"graphene-1.0", nil},
	{"gtk4", nil},
}

func main() {
	var repos gir.Repositories
	var err error

	for _, pkg := range packages {
		if pkg.Namespaces != nil {
			err = repos.AddSelected(pkg.PkgName, pkg.Namespaces)
		} else {
			err = repos.Add(pkg.PkgName)
		}

		if err != nil {
			log.Fatalln("error adding packages:", err)
		}
	}

	var wg sync.WaitGroup

	sema := make(chan struct{}, runtime.GOMAXPROCS(-1))

	gen := girgen.NewGenerator(repos)
	gen.WithLogger(log.New(os.Stderr, "[girgen] ", log.LstdFlags|log.Lmsgprefix))

	// Do a clean-up of the target directory.
	if err := os.RemoveAll(output); err != nil {
		log.Println("non-fatal: failed to rm -rf output dir:", err)
	}

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			ng := gen.UseNamespace(namespace.Name)

			sema <- struct{}{}
			wg.Add(1)

			go func() {
				writeNamespace(ng)

				<-sema
				wg.Done()
			}()
		}
	}

	wg.Wait()

	if err := goimports.Dir(output); err != nil {
		log.Println("failed to run goimports on "+output+":", err)
	}
}

func writeNamespace(ng *girgen.NamespaceGenerator) {
	log.Println("generating", ng.PackageName(), "at", ng.Repository().Path)

	pkg := ng.PackageName()
	dir := filepath.Join(output, pkg)
	out := filepath.Join(dir, pkg+".go")

	if err := os.MkdirAll(dir, os.ModePerm|os.ModeDir); err != nil {
		log.Println("mkdir -p failed:", err)
		return
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

}
