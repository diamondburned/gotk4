package genmain

import (
	"context"
	"fmt"
	"go/format"
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
	"golang.org/x/sync/semaphore"
)

// StringSet joins the given slices of strings into a map with the keys as the
// values of each of the given slices.
func StringSet(strs ...[]string) map[string]struct{} {
	var length int
	for _, str := range strs {
		length += len(str)
	}

	set := make(map[string]struct{}, length)

	for _, str := range strs {
		for _, s := range str {
			set[s] = struct{}{}
		}
	}

	return set
}

// ModulePath crafts the full module path from the given base module path. If
// the overrides map is given, then the function will use that as the list of
// special cases.
func ModulePath(module string, overrides map[string]string) func(*gir.Namespace) string {
	return func(namespace *gir.Namespace) string {
		if path, ok := overrides[gir.VersionedNamespace(namespace)]; ok {
			return path
		}

		modulePath := path.Join(module, gir.GoPackageName(namespace.Name))
		if version := MajorVersion(namespace); version > 1 {
			modulePath = path.Join(modulePath, fmt.Sprintf("v%d", version))
		}

		return modulePath
	}
}

// MajorVersion returns the major version of the GIR namespace in int.
func MajorVersion(nsp *gir.Namespace) int {
	version := gir.MajorVersion(nsp.Version)

	v, err := strconv.Atoi(version)
	if err != nil {
		log.Panicf("invalid version %q", version)
	}

	return v
}

// LoadExternOverrides creates a map of GIR versioned namespace names to the Go
// import path from the given list of all repositories.
func LoadExternOverrides(mod string, repos gir.Repositories) map[string]string {
	externs := map[string]string{}
	modPath := ModulePath(mod, nil)

	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			externs[gir.VersionedNamespace(&namespace)] = modPath(&namespace)
		}
	}

	return externs
}

// WriteNamespace generates everything in the given namespace and writes it to
// the given basePath.
func WriteNamespace(ng *girgen.NamespaceGenerator, basePath string) error {
	dir := filepath.Join(basePath, ng.PkgName)

	if version := MajorVersion(ng.Namespace().Namespace); version > 1 {
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

// MustLoadPackages bails if packages fail to load.
func MustLoadPackages(pkgs []Package) gir.Repositories {
	p, err := LoadPackages(pkgs)
	if err != nil {
		log.Fatalln(err)
	}
	return p
}

// LoadPackages loads all GIR repositories from the given list of packages.
func LoadPackages(pkgs []Package) (gir.Repositories, error) {
	var repos gir.Repositories
	return repos, AddPackages(&repos, pkgs)
}

// MustAddPackages bails if packages fail to be added.
func MustAddPackages(repos *gir.Repositories, pkgs []Package) {
	if err := AddPackages(repos, pkgs); err != nil {
		log.Fatalln(err)
	}
}

// AddPackages adds the given list of packages into the repository.
func AddPackages(repos *gir.Repositories, pkgs []Package) error {
	for _, pkg := range pkgs {
		var err error
		if pkg.Namespaces != nil {
			err = repos.AddSelected(pkg.Name, pkg.Namespaces)
		} else {
			err = repos.Add(pkg.Name)
		}

		if err != nil {
			return errors.Wrapf(err, "error adding package %q", pkg.Name)
		}
	}

	return nil
}

// PrintAddedPkgs prints to console the added packages.
func PrintAddedPkgs(repos gir.Repositories) {
	for _, repo := range repos {
		for _, namespace := range repo.Namespaces {
			log.Println(
				"added from package", repo.Pkg,
				"namespace", gir.VersionedNamespace(&namespace),
				repo.Path,
			)
		}
	}
}

// GenerateAll generates all namespaces inside the given generator into the
// given dst path. It uses WriteNamespace to do so. The namespaces will be
// generated in parallel. If a versioned namespace is in the except slice, then
// it is skipped.
func GenerateAll(gen *girgen.Generator, dst string, except []string) []error {
	sema := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(-1)))
	var wg sync.WaitGroup

	exemptions := StringSet(except)

	var errMut sync.Mutex
	var errors []error

	for _, repo := range gen.Repositories() {
		for _, namespace := range repo.Namespaces {
			_, exempted := exemptions[gir.VersionedNamespace(&namespace)]
			if exempted {
				continue
			}

			ng := gen.UseNamespace(namespace.Name, namespace.Version)
			if ng == nil {
				log.Fatalln("cannot find namespace", namespace.Name, "v"+namespace.Version)
			}

			sema.Acquire(context.Background(), 1)
			wg.Add(1)

			go func() {
				if err := WriteNamespace(ng, dst); err != nil {
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

	return errors
}

// GeneratePackages generates the given pkgs list into the given dst directory.
// It uses WriteNamespace to do so. The namespaces will be generated in
// parallel. Most external GIR generators should call this.
func GeneratePackages(gen *girgen.Generator, dst string, pkgs []Package, except []string) []error {
	sema := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(-1)))

	var wg sync.WaitGroup
	defer wg.Wait()

	exemptions := StringSet(except)

	var errMut sync.Mutex
	var errors []error

	genNamespace := func(namespace *gir.Namespace) {
		ng := gen.UseNamespace(namespace.Name, namespace.Version)
		if ng == nil {
			log.Fatalln("cannot find namespace", namespace.Name, "v"+namespace.Version)
		}

		sema.Acquire(context.Background(), 1)
		wg.Add(1)

		go func() {
			if err := WriteNamespace(ng, dst); err != nil {
				errMut.Lock()
				errors = append(errors, err)
				errMut.Unlock()
			}

			sema.Release(1)
			wg.Done()
		}()
	}

	repos := gen.Repositories()

	for _, pkg := range pkgs {
		if pkg.Namespaces != nil {
			for _, wantedName := range pkg.Namespaces {
				namespace := repos.FindNamespace(wantedName)
				if namespace == nil {
					return []error{fmt.Errorf("namespace %q not found", wantedName)}
				}

				_, exempted := exemptions[gir.VersionedNamespace(namespace.Namespace)]
				if exempted {
					continue
				}

				genNamespace(namespace.Namespace)
			}
		}

		repo := repos.FromPkg(pkg.Name)
		if repo == nil {
			return []error{fmt.Errorf("package %q not found", pkg.Name)}
		}

		for _, namespace := range repo.Namespaces {
			genNamespace(&namespace)
		}
	}

	wg.Wait()

	return errors
}

// CleanDirectory cleans up the directory at the given path. Files listed inside
// except will not be wiped.
func CleanDirectory(path string, except []string) error {
	// Do a clean-up of the target directory.
	oldFiles, _ := os.ReadDir(path)
	pkgExceptions := StringSet(except)

	for _, oldFile := range oldFiles {
		_, except := pkgExceptions[oldFile.Name()]
		if except {
			continue
		}

		fullPath := filepath.Join(path, oldFile.Name())
		if err := os.RemoveAll(fullPath); err != nil {
			return errors.Wrapf(err, "failed to rm -rf %s", oldFile)
		}
	}

	return nil
}

// AppendGoFiles appends the value of the given contents map into the files at
// its keys and run go fmt on it.
func AppendGoFiles(path string, contents map[string]string) error {
	for name, content := range contents {
		if err := appendGoFile(path, name, content); err != nil {
			return err
		}
	}
	return nil
}

func appendGoFile(path, filename, content string) error {
	fullPath := filepath.Join(path, filename)

	b, err := os.ReadFile(fullPath)
	if err != nil {
		return errors.Wrapf(err, "failed to read file %q", filename)
	}

	b = append(b, []byte(content)...)

	b, err = format.Source(b)
	if err != nil {
		return errors.Wrapf(err, "failed to go fmt file %q", filename)
	}

	if err := os.WriteFile(fullPath, b, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write file %q", filename)
	}

	return nil
}

// EnsureDirectory ensures that all files inside the given directory path are
// present in the given list of string slices.
func EnsureDirectory(path string, expects ...[]string) error {
	wantedFiles := StringSet(expects...)

	files, err := os.ReadDir(path)
	if err != nil {
		return errors.Wrap(err, "failed to read dir")
	}

	for _, file := range files {
		if _, ok := wantedFiles[file.Name()]; !ok {
			return fmt.Errorf("unexpected file/folder %q", file.Name())
		}

		delete(wantedFiles, file.Name())
	}

	for name := range wantedFiles {
		return fmt.Errorf("missing file/folder %q", name)
	}

	return nil
}
