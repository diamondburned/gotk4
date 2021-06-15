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
	"github.com/pkg/errors"
)

var (
	output  string
	module  string
	verbose bool
	listPkg bool
)

func init() {
	flag.StringVar(&output, "o", "", "output directory to mkdir in")
	flag.StringVar(&module, "m", "github.com/diamondburned/gotk4", "go module name")
	flag.BoolVar(&verbose, "v", verbose, "log verbosely (debug mode)")
	flag.BoolVar(&listPkg, "l", listPkg, "only list packages and exit")
	flag.Parse()

	if !listPkg && output == "" {
		log.Fatalln("Missing -o output directory.")
	}

	if !verbose {
		verbose = os.Getenv("GIR_VERBOSE") == "1"
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

	if listPkg {
		return
	}

	var wg sync.WaitGroup

	var errMut sync.Mutex
	var errors []error

	sema := make(chan struct{}, runtime.GOMAXPROCS(-1))

	gen := girgen.NewGenerator(repos, modulePath)
	gen.Color = true
	gen.Logger = log.New(os.Stderr, "girgen: ", log.Lmsgprefix)
	gen.AddFilters(filters)

	if verbose {
		gen.LogLevel = girgen.LogDebug
	}

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
				if err := writeNamespace(ng); err != nil {
					errMut.Lock()
					errors = append(errors, err)
					errMut.Unlock()
				}

				<-sema
				wg.Done()
			}()
		}
	}

	wg.Wait()

	if len(errors) > 0 {
		for _, err := range errors {
			log.Println("generation error:", err)
		}

		os.Exit(1)
	}
}

func writeNamespace(ng *girgen.NamespaceGenerator) error {
	pkg := ng.PackageName()
	dir := filepath.Join(output, pkg)

	if version := majorVer(ng.Namespace().Namespace); version > 1 {
		// Follow Go's convention of a versioned package, so we can generate
		// multiple versions.
		dir = filepath.Join(dir, fmt.Sprintf("v%d", version))
	}

	files, genErr := ng.Generate()

	if len(files) > 0 {
		if err := os.MkdirAll(dir, os.ModePerm|os.ModeDir); err != nil {
			return errors.Wrap(err, "mkdir -p failed")
		}

		for name, data := range files {
			if err := os.WriteFile(filepath.Join(dir, name), data, 0666); err != nil {
				return errors.Wrapf(err, "failed to write %s", name)
			}
		}
	}

	// Preserve the generation error, but give it last priority.
	return genErr
}

func modulePath(namespace *gir.Namespace) string {
	modulePath := path.Join(module, output, gir.GoPackageName(namespace.Name))
	if version := majorVer(namespace); version > 1 {
		modulePath = path.Join(modulePath, fmt.Sprintf("v%d", version))
	}

	return modulePath
}

func majorVer(nsp *gir.Namespace) int {
	v, err := strconv.Atoi(gir.MajorVersion(nsp.Version))
	if err != nil {
		return 0
	}
	return v
}
