package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen"
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
				"namespace", gir.VersionedNamespace(&namespace),
				repo.Path,
			)
		}
	}

	var wg sync.WaitGroup

	sema := make(chan struct{}, runtime.GOMAXPROCS(-1))

	gen := girgen.NewGenerator(repos, path.Join(module, output))
	gen.WithLogger(log.New(os.Stderr, "girgen: ", log.LstdFlags|log.Lmsgprefix), true)
	gen.AddFilters(filters)

	// Do a clean-up of the target directory.
	if err := os.RemoveAll(output); err != nil {
		log.Println("non-fatal: failed to rm -rf output dir:", err)
	}

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			ng := gen.UseNamespace(namespace.Name, namespace.Version)

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
}

func writeNamespace(ng *girgen.NamespaceGenerator) {
	pkg := ng.PackageName()
	dir := filepath.Join(output, pkg)

	if version := majorVer(ng); version > 1 {
		// Follow Go's convention of a versioned package, so we can generate
		// multiple versions.
		dir = filepath.Join(dir, fmt.Sprintf("v%d", version))
	}

	if err := os.MkdirAll(dir, os.ModePerm|os.ModeDir); err != nil {
		log.Println("mkdir -p failed:", err)
		return
	}

	b, err := ng.Generate()
	if err != nil {
		log.Println("generation error:", err)
	}

	// Write to file any non-empty output.
	if len(b) > 0 {
		goFile := filepath.Join(dir, pkg+".go")

		if err := os.WriteFile(goFile, b, os.ModePerm); err != nil {
			log.Println("failed to write file:", err)
		}
	}
}

func majorVer(ng *girgen.NamespaceGenerator) int {
	v, err := strconv.Atoi(gir.MajorVersion(ng.Namespace().Version))
	if err != nil {
		return 0
	}
	return v
}
