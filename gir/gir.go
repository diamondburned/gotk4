package gir

import (
	"fmt"
	"path"
	"strings"
	"unicode"

	"github.com/diamondburned/gotk4/gir/pkgconfig"
	"github.com/pkg/errors"
)

// PackageRoot is the root of the gotk4 (this) package.
const PackageRoot = "github.com/diamondburned/gotk4"

// ImportPath generates the full import path from the package root.
func ImportPath(pkgPath string) string {
	return path.Join(PackageRoot, pkgPath)
}

// GoPackageName converts a GIR package name to a Go package name. It's only
// tested against a known set of GIR files.
func GoPackageName(girPkgName string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, girPkgName)
}

// GoNamespace converts a namespace's name to a Go package name.
func GoNamespace(namespace Namespace) string {
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

	girs, err := pkgconfig.FindGirFiles(pkg)
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

		*repos = append(*repos, PkgRepository{
			Repository: *repo,
			Pkg:        pkg,
			Path:       gir,
		})
	}

	if found != len(namespaces) {
		return fmt.Errorf("only %d girs found", found)
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

		*repos = append(*repos, PkgRepository{
			Repository: *repo,
			Pkg:        pkg,
			Path:       gir,
		})
	}

	return nil
}

// NamespaceFindResult is the result returned from FindNamespace.
type NamespaceFindResult struct {
	Repository PkgRepository
	Namespace  Namespace
}

// Eq compares that the resulting namespace's name and version match.
func (res *NamespaceFindResult) Eq(other *NamespaceFindResult) bool {
	return true &&
		res.Namespace.Name == other.Namespace.Name &&
		res.Namespace.Version == other.Namespace.Version
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

	SameNamespace bool

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
		return res.Function.Name, res.Function.CType
	case res.Union != nil:
		return res.Union.Name, res.Union.CType
	case res.Bitfield != nil:
		return res.Bitfield.Name, res.Bitfield.CType
	case res.Callback != nil:
		return res.Callback.Name, res.Callback.CType
	}

	panic("TypeFindResult has all fields nil")
}

// FindType finds a type in the repositories from the given current namespace
// name, version, and the GIR type name.
func (repos *Repositories) FindType(nspName, nspVersion, typ string) *TypeFindResult {
	var r TypeFindResult

	// need this for the version
	currentNamespace := repos.FindNamespace(nspName, nspVersion)

	if namespace, typeName := SplitGIRType(typ); namespace != "" {
		// Search the namespace's version, if possible or available.
		var version string
		for _, incl := range currentNamespace.Repository.Includes {
			if incl.Name == nspName && incl.Version != "" {
				version = incl.Version
				break
			}
		}

		r.NamespaceFindResult = repos.FindNamespace(namespace, version)
		typ = typeName

	} else {
		r.NamespaceFindResult = currentNamespace
		r.SameNamespace = true
		typ = typeName
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

// // FindCType works like FindType but for C types.
// func (repos *Repositories) FindCType(cTyp string) *TypeFindResult {
// 	r := TypeFindResult{
// 		NamespaceFindResult: &NamespaceFindResult{},
// 	}

// 	for _, r.Repository = range *repos {
// 		for _, r.Namespace = range r.Repository.Namespaces {
// 			// Avoid searching the whole namespace by verifying the prefix.
// 			if !strings.HasPrefix(cTyp, r.Namespace.CIdentifierPrefixes) {
// 				continue
// 			}

// 			for _, class := range r.Namespace.Classes {
// 				if class.Name == cTyp {
// 					r.Class = &class
// 					return &r
// 				}
// 			}

// 			for _, enum := range r.Namespace.Enums {
// 				if enum.Name == cTyp {
// 					r.Enum = &enum
// 					return &r
// 				}
// 			}

// 			for _, function := range r.Namespace.Functions {
// 				if function.Name == cTyp {
// 					r.Function = &function
// 					return &r
// 				}
// 			}

// 			for _, callback := range r.Namespace.Callbacks {
// 				if callback.Name == cTyp {
// 					r.Callback = &callback
// 					return &r
// 				}
// 			}

// 			for _, iface := range r.Namespace.Interfaces {
// 				if iface.Name == cTyp {
// 					r.Interface = &iface
// 					return &r
// 				}
// 			}
// 		}
// 	}

// 	return nil
// }

// PkgRepository wraps a Repository to add additional information.
type PkgRepository struct {
	Repository
	Pkg  string // arg for pkg-config
	Path string // gir file path
}
