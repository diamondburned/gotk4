package gir

import (
	"fmt"
	"path"
	"strings"
	"sync"
	"unicode"

	"github.com/diamondburned/gotk4/gir/pkgconfig"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
)

// ImportPath generates the full import path from the package root.
func ImportPath(root, pkgPath string) string {
	return path.Join(root, pkgPath)
}

// goPackageNameTrans describes the transformation checks to filter out runes.
var goPackageNameTrans = []func(r rune) bool{
	unicode.IsLetter,
	unicode.IsDigit,
}

// GoPackageName converts a GIR package name to a Go package name. It's only
// tested against a known set of GIR files.
func GoPackageName(girPkgName string) string {
	return strings.Map(func(r rune) rune {
		for _, trans := range goPackageNameTrans {
			if trans(r) {
				return unicode.ToLower(r)
			}
		}

		return -1
	}, girPkgName)
}

// GoNamespace converts a namespace's name to a Go package name.
func GoNamespace(namespace *Namespace) string {
	return GoPackageName(namespace.Name)
}

// SplitGIRType splits the given GIR type string into 2 parts: the namespace,
// which preceeds the period, and the type name. If there is no period, then an
// empty string is returned for namespace.
func SplitGIRType(typ string) (namespace, typeName string) {
	parts := strings.SplitN(typ, ".", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}

	return "", parts[0]
}

// Repositories contains a list of known repositories.
type Repositories []PkgRepository

// AddSelected adds a single package but only searches for the given list of
// GIR files.
func (repos *Repositories) AddSelected(pkg string, namespaces []string) error {
	found := 0

	filter := func(r *Repository) bool {
		repoNames := r.Namespaces
		r.Namespaces = repoNames[:0]

		for _, namespace := range repoNames {
			for _, nsp := range namespaces {
				if nsp == namespace.Name {
					r.Namespaces = append(r.Namespaces, namespace)
					found++
					break
				}
			}
		}

		return len(r.Namespaces) > 0
	}

	girs, err := pkgconfig.FindGIRFiles(pkg)
	if err != nil {
		return errors.Wrapf(err, "failed to get gir files for %q", pkg)
	}

	for _, gir := range girs {
		repo, err := ParseRepository(gir)
		if err != nil {
			return errors.Wrapf(err, "failed to parse file %q", gir)
		}

		if !filter(repo) {
			continue
		}

		if err := repos.add(*repo, pkg, gir); err != nil {
			return errors.Wrapf(err, "failed to add file %q", gir)
		}
	}

	if found != len(namespaces) {
		return fmt.Errorf("only %d girs found", found)
	}

	return nil
}

// Add finds the given pkg name to be searched using pkg-config and added into
// the list of repositories.
func (repos *Repositories) Add(pkg string) error {
	girs, err := pkgconfig.FindGIRFiles(pkg)
	if err != nil {
		return errors.Wrapf(err, "failed to get gir files for %q", pkg)
	}

	for _, gir := range girs {
		repo, err := ParseRepository(gir)
		if err != nil {
			return errors.Wrapf(err, "failed to parse file %q", gir)
		}

		if err := repos.add(*repo, pkg, gir); err != nil {
			return errors.Wrapf(err, "failed to add file %q", gir)
		}
	}

	return nil
}

// add adds the given PkgRepository.
func (repos *Repositories) add(r Repository, pkg, path string) error {
	for _, repo := range *repos {
		for _, repoNsp := range repo.Namespaces {
			for _, addingNsp := range r.Namespaces {
				if addingNsp.Name == repoNsp.Name {
					return fmt.Errorf(
						"colliding namespace %s, got v%s, add v%s",
						addingNsp.Name, repoNsp.Version, addingNsp.Version,
					)
				}
			}
		}
	}

	*repos = append(*repos, PkgRepository{
		Repository: r,
		Pkg:        pkg,
		Path:       path,
	})

	return nil
}

// NamespaceFindResult is the result returned from FindNamespace.
type NamespaceFindResult struct {
	Repository *PkgRepository
	Namespace  *Namespace
}

// Eq compares that the resulting namespace's name and version match.
func (res *NamespaceFindResult) Eq(other *NamespaceFindResult) bool {
	if other == res {
		return true
	}

	return true &&
		res.Namespace.Name == other.Namespace.Name &&
		res.Repository.Pkg == other.Repository.Pkg
}

// FindNamespace finds the repository and namespace with the given name and
// version.
func (repos *Repositories) FindNamespace(name string) *NamespaceFindResult {
	for i := range *repos {
		repository := &(*repos)[i]

		for j := range repository.Namespaces {
			namespace := &repository.Namespaces[j]
			if namespace.Name != name {
				continue
			}

			return &NamespaceFindResult{
				Repository: repository,
				Namespace:  namespace,
			}
		}
	}

	return nil
}

// TypeFindResult is the result
type TypeFindResult struct {
	*NamespaceFindResult

	// Only one of these fields are not nil. They should also be read-only.
	Alias     *Alias
	Class     *Class
	Interface *Interface
	Record    *Record
	Enum      *Enum
	Function  *Function
	Union     *Union
	Bitfield  *Bitfield
	Callback  *Callback

	// TODO: Constant, Annotations, Boxed
	// TODO: Methods
	// TODO: Enum Members
}

// Info gets the name and C type of the resulting type. The name returned is in
// camel case.
func (res *TypeFindResult) Info() (name, ctype string) {
	switch {
	case res.Alias != nil:
		return res.Alias.Name, res.Alias.CType
	case res.Class != nil:
		return res.Class.Name, res.Class.CType
	case res.Interface != nil:
		return res.Interface.Name, res.Interface.CType
	case res.Record != nil:
		return res.Record.Name, res.Record.CType
	case res.Enum != nil:
		return res.Enum.Name, res.Enum.CType
	case res.Function != nil:
		return res.Function.Name, ""
	case res.Union != nil:
		return res.Union.Name, res.Union.CType
	case res.Bitfield != nil:
		return res.Bitfield.Name, res.Bitfield.CType
	case res.Callback != nil:
		return res.Callback.Name, ""
	}

	panic("TypeFindResult has all fields nil")
}

var (
	typeResultCache  sync.Map // full GIR type -> *TypeFindResult
	typeResultFlight singleflight.Group
)

// FindType finds a type in the repositories from the given current namespace
// name, version, and the GIR type name. The function will cache the returned
// TypeFindResult.
func (repos *Repositories) FindType(nspName, typ string) *TypeFindResult {
	// Build the full type name for cache querying. Use a faster method to check
	// if the type name already has a namespace by detecting for a period (".").
	fullType := typ
	if !strings.Contains(typ, ".") {
		fullType = nspName + "." + typ
	}

	v, ok := typeResultCache.Load(fullType)
	if ok {
		return v.(*TypeFindResult)
	}

	v, _, _ = typeResultFlight.Do(fullType, func() (interface{}, error) {
		result := repos.findType(nspName, typ)
		if result != nil {
			typeResultCache.Store(fullType, result)
		}

		return result, nil
	})

	return v.(*TypeFindResult)
}

func (repos *Repositories) findType(nspName, typ string) *TypeFindResult {
	var r TypeFindResult

	if namespace, typeName := SplitGIRType(typ); namespace != "" {
		r.NamespaceFindResult = repos.FindNamespace(namespace)
		typ = typeName
	} else {
		r.NamespaceFindResult = repos.FindNamespace(nspName)
	}

	if r.NamespaceFindResult == nil {
		return nil
	}

	for _, alias := range r.Namespace.Aliases {
		if alias.Name == typ {
			r.Alias = &alias
			return &r
		}
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

	for _, record := range r.Namespace.Records {
		if record.Name == typ {
			r.Record = &record
			return &r
		}
	}

	for _, function := range r.Namespace.Functions {
		if function.Name == typ {
			r.Function = &function
			return &r
		}
	}

	for _, union := range r.Namespace.Unions {
		if union.Name == typ {
			r.Union = &union
			return &r
		}
	}

	for _, bitfield := range r.Namespace.Bitfields {
		if bitfield.Name == typ {
			r.Bitfield = &bitfield
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

	return nil
}

// PkgRepository wraps a Repository to add additional information.
type PkgRepository struct {
	Repository
	Pkg  string // arg for pkg-config
	Path string // gir file path
}
