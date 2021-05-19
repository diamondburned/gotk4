package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/goimports"
)

var (
	output string
	module string
)

func init() {
	flag.StringVar(&output, "o", "", "output directory to mkdir in")
	flag.StringVar(&module, "m", "github.com/diamondburned/gotk4", "go module name")
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
		"GLib", "Gio", "Vulkan", "cairo", "xft", "xlib", "freetype2", "fontconfig",
	}},
	{"pango", nil},
	{"gdk-pixbuf-2.0", []string{"GdkPixbuf", "GdkPixdata"}},
	{"graphene-1.0", nil},
	{"gtk4", nil}, // includes Gdk
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

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			log.Println(
				"added from package", repo.Pkg,
				"namespace", namespace.Name,
				"v"+namespace.Version, repo.Path,
			)
		}
	}

	var wg sync.WaitGroup

	sema := make(chan struct{}, runtime.GOMAXPROCS(-1))

	gen := girgen.NewGenerator(repos, path.Join(module, output))
	gen.WithLogger(log.New(os.Stderr, "girgen: ", log.LstdFlags|log.Lmsgprefix), true)

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

	log.Println("running goimports on output...")
	if err := goimports.Dir(output); err != nil {
		log.Println("failed to run goimports on "+output+":", err)
	}
}

func writeNamespace(ng *girgen.NamespaceGenerator) {
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
