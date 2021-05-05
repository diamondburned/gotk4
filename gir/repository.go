package gir

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"strings"

	"github.com/diamondburned/gotk4/gir/pkgconfig"
	"github.com/pkg/errors"
)

// Debug, change to true to enable verbose logging.
var Debug = false

func debugln(v ...interface{}) {
	if Debug {
		log.Println(v...)
	}
}

// Repositories contains a list of known repositories.
type Repositories []Repository

// AddAll adds multiple packages using Add.
func (repos *Repositories) AddAll(pkgs ...string) error {
	for _, pkg := range pkgs {
		if err := repos.Add(pkg); err != nil {
			return errors.Wrapf(err, "failed to add %s", pkg)
		}
	}

	return nil
}

// Add finds the given pkg name to be searched using pkg-config and added into
// the list of repositories.
func (repos *Repositories) Add(pkg string) error {
	girs, err := pkgconfig.FindGirFiles(pkg)
	if err != nil {
		return errors.Wrapf(err, "failed to get gir files for %q", pkg)
	}

	for _, gir := range girs {
		repo, err := ParseRepository(gir)
		if err != nil {
			return errors.Wrapf(err, "failed to parse file %q", gir)
		}

		*repos = append(*repos, *repo)
	}

	return nil
}

// NamespaceFindResult is the result returned from FindNamespace.
type NamespaceFindResult struct {
	Repository Repository
	Namespace  Namespace
}

// FindNamespace finds the repository and namespace with the given name and
// version.
func (repos *Repositories) FindNamespace(name, version string) *NamespaceFindResult {
	for _, repo := range *repos {
		for _, nsp := range repo.Namespaces {
			if nsp.Name != name {
				continue
			}
			// Only skip the namespace if the version is not empty AND it
			// doesn't match, in case a namespace doesn't actually have a
			// version.
			if nsp.Version != version && version != "" {
				continue
			}

			return &NamespaceFindResult{
				Repository: repo,
				Namespace:  nsp,
			}
		}
	}

	return nil
}

// TypeFindResult is the result
type TypeFindResult struct {
	*NamespaceFindResult

	// Only one of these fields are not nil. They should also be read-only.
	Class     *Class
	Enum      *Enum
	Function  *Function
	Callback  *Callback
	Interface *Interface
}

// FindType finds a type in the repositories from the given current namespace
// name, version, and the type name.
func (repos *Repositories) FindType(nspName, nspVersion, typ string) *TypeFindResult {
	var r TypeFindResult

	// need this for the version
	currentNamespace := repos.FindNamespace(nspName, nspVersion)

	parts := strings.SplitN(typ, ".", 2)
	if len(parts) != 2 {
		r.NamespaceFindResult = currentNamespace
	} else {
		// Search the namespace's version, if possible or available.
		var version string
		for _, incl := range currentNamespace.Repository.Includes {
			if incl.Name == nspName && incl.Version != nil {
				version = *incl.Version
				break
			}
		}

		r.NamespaceFindResult = repos.FindNamespace(parts[1], version)
		typ = parts[1]
	}

	if r.NamespaceFindResult == nil {
		// TODO: This is shadowing the namespace-in-type-not-found error.
		debugln("namespace", nspName, "not found")
		return nil
	}

	for _, class := range r.Namespace.Classes {
		if class.Name == typ {
			r.Class = &class
			return &r
		}
	}

	for _, enum := range r.Namespace.Enums {
		if enum.Name == typ {
			r.Enum = &enum
			return &r
		}
	}

	for _, function := range r.Namespace.Functions {
		if function.Name == typ {
			r.Function = &function
			return &r
		}
	}

	for _, callback := range r.Namespace.Callbacks {
		if callback.Name == typ {
			r.Callback = &callback
			return &r
		}
	}

	for _, iface := range r.Namespace.Interfaces {
		if iface.Name == typ {
			r.Interface = &iface
			return &r
		}
	}

	debugln("type", typ, "not found")
	return nil
}

// Repository represents a GObject Introspection Repository, which contains the
// includes, C includes and namespaces of a single gir file.
type Repository struct {
	Includes   []Include   `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	CIncludes  []CInclude  `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Namespaces []Namespace `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
}

// ParseRepository parses a repository from the given file path.
func ParseRepository(file string) (*Repository, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open file")
	}
	defer f.Close()

	return ParseRepositoryFromReader(f)
}

// ParseRepositoryFromReader parses a repository from the given reader.
func ParseRepositoryFromReader(r io.Reader) (*Repository, error) {
	var repo Repository

	if err := xml.NewDecoder(r).Decode(&repo); err != nil {
		return nil, errors.Wrap(err, "Failed to decode gir XML")
	}

	return &repo, nil
}
