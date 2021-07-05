package main

import (
	"context"
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
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/pkg/errors"
	"golang.org/x/sync/semaphore"
)

var (
	output   string
	module   string
	verbose  bool
	listPkg  bool
	parallel = int64(runtime.GOMAXPROCS(-1))
)

func init() {
	flag.StringVar(&output, "o", "", "output directory to mkdir in")
	flag.StringVar(&module, "m", "github.com/diamondburned/gotk4", "go module name")
	flag.BoolVar(&verbose, "v", verbose, "log verbosely (debug mode)")
	flag.BoolVar(&listPkg, "l", listPkg, "only list packages and exit")
	flag.Int64Var(&parallel, "p", parallel, "number of generator goroutines to spawn")
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

	gen := girgen.NewGenerator(repos, modulePath)
	gen.Logger = log.New(os.Stderr, "girgen: ", log.Lmsgprefix)
	gen.AddFilters(filters)
	gen.ApplyPreprocessors(preprocessors)

	if verbose {
		gen.LogLevel = logger.Debug
	}

	// Do a clean-up of the target directory.
	oldFiles, _ := os.ReadDir(output)
deleteLoop:
	for _, oldFile := range oldFiles {
		for _, except := range pkgExceptions {
			if except == oldFile.Name() {
				continue deleteLoop
			}
		}

		fullPath := filepath.Join(output, oldFile.Name())
		if err := os.RemoveAll(fullPath); err != nil {
			log.Printf("non-fatal: failed to rm -rf %s/: %v", oldFile, err)
		}
	}

	sema := semaphore.NewWeighted(parallel)

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			ng := gen.UseNamespace(namespace.Name, namespace.Version)
			if ng == nil {
				log.Fatalln("cannot find namespace", namespace.Name, "v"+namespace.Version)
			}

			sema.Acquire(context.Background(), 1)
			wg.Add(1)

			go func() {
				if err := writeNamespace(ng); err != nil {
					errMut.Lock()
					errors = append(errors, err)
					errMut.Unlock()
				}

				sema.Release(1)
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
	dir := filepath.Join(output, ng.PkgName)

	if version := majorVer(ng.Namespace().Namespace); version > 1 {
		// Follow Go's convention of a versioned package, so we can generate
		// multiple versions.
		dir = filepath.Join(dir, fmt.Sprintf("v%d", version))
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		return errors.Wrapf(err, "failed to mkdir -p %q", dir)
	}

	files, err := ng.Generate()

	for name, file := range files {
		dst := filepath.Join(dir, name)
		if err := os.WriteFile(dst, file, 0666); err != nil {
			return errors.Wrapf(err, "failed to write to %s", dst)
		}
	}

	// Preserve the generation error, but give it last priority.
	return err
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
