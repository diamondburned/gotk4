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

// MajorVersion returns the major version number only.
func MajorVersion(version string) string {
	parts := strings.SplitN(version, ".", 2)
	return parts[0]
}

// EqVersion compares major versions.
func EqVersion(v1, v2 string) bool { return MajorVersion(v1) == MajorVersion(v2) }

// VersionedNamespace returns the versioned name for the given namespace.
func VersionedNamespace(namespace *Namespace) string {
	return VersionedName(namespace.Name, namespace.Version)
}

// VersionedName returns the name appended with the version suffix.
func VersionedName(name, version string) string {
	return name + "-" + MajorVersion(version)
}

// ParseVersionName parses the given fullName to return the original name and
// the version separately. If no versions are available, then the empty string
// is returned.
func ParseVersionName(fullName string) (name, majorVersion string) {
	parts := strings.SplitN(fullName, "-", 2)
	if len(parts) != 2 {
		return parts[0], ""
	}

	return parts[0], parts[1]
}

// Repositories contains a list of known repositories.
type Repositories []PkgRepository

// PkgRepository wraps a Repository to add additional information.
type PkgRepository struct {
	Repository
	Pkg  string // arg for pkg-config
	Path string // gir file path
}

// AddSelected adds a single package but only searches for the given list of
// GIR files.
func (repos *Repositories) AddSelected(pkg string, namespaces []string) error {
	found := 0

	filter := func(r *Repository) bool {
		repoNames := r.Namespaces
		r.Namespaces = repoNames[:0]

		for _, fullName := range namespaces {
			nsp, version := ParseVersionName(fullName)

			for _, namespace := range repoNames {
				if namespace.Name != nsp {
					continue
				}
				if version != "" && !EqVersion(namespace.Version, version) {
					continue
				}

				r.Namespaces = append(r.Namespaces, namespace)
				found++
				break
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
				if addingNsp.Name != repoNsp.Name || !EqVersion(addingNsp.Version, repoNsp.Version) {
					continue
				}

				return fmt.Errorf(
					"colliding namespace %s, got v%s, add v%s",
					addingNsp.Name, repoNsp.Version, addingNsp.Version,
				)
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
		res.Namespace.Version == other.Namespace.Version &&
		res.Repository.Pkg == other.Repository.Pkg &&
		res.Repository.Version == other.Repository.Version
}

// FindNamespace finds the repository and namespace with the given name and
// version. If name doesn't have the version bits, then it panics.
func (repos Repositories) FindNamespace(name string) *NamespaceFindResult {
	name, version := ParseVersionName(name)
	if version == "" {
		panic("FindNamespace given namespace unversioned: " + name)
	}

	for i := range repos {
		repository := &repos[i]

		for j := range repository.Namespaces {
			namespace := &repository.Namespaces[j]
			if namespace.Name != name || MajorVersion(namespace.Version) != version {
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
}

// Info gets the name and C type of the resulting type. The name returned is in
// camel case.
//
// TODO: split Info into Name() and CType().
func (res *TypeFindResult) Info() (name, ctype string) {
	return res.Name(false), res.CType()
}

// CType returns the resulting type's C type identifier.
func (res *TypeFindResult) CType() string {
	switch {
	case res.Alias != nil:
		return res.Alias.CType
	case res.Class != nil:
		return res.Class.GLibTypeName
	case res.Interface != nil:
		return res.Interface.CType
	case res.Record != nil:
		return res.Record.CType
	case res.Enum != nil:
		return res.Enum.CType
	case res.Function != nil:
		return res.Function.CIdentifier
	case res.Union != nil:
		return res.Union.CType
	case res.Bitfield != nil:
		return res.Bitfield.CType
	case res.Callback != nil:
		return res.Callback.CIdentifier
	}

	panic("TypeFindResult has all fields nil")
}

// Name returns the resulting type's GIR name, with or without the namespace.
func (res *TypeFindResult) Name(needsNamespace bool) string {
	var typ string
	switch {
	case res.Alias != nil:
		typ = res.Alias.Name
	case res.Class != nil:
		typ = res.Class.Name
	case res.Interface != nil:
		typ = res.Interface.Name
	case res.Record != nil:
		typ = res.Record.Name
	case res.Enum != nil:
		typ = res.Enum.Name
	case res.Function != nil:
		typ = res.Function.Name
	case res.Union != nil:
		typ = res.Union.Name
	case res.Bitfield != nil:
		typ = res.Bitfield.Name
	case res.Callback != nil:
		typ = res.Callback.Name
	}

	if typ == "" {
		panic("TypeFindResult has all fields nil")
	}

	if needsNamespace {
		typ = res.Namespace.Name + "." + typ
	}

	return typ
}

// IsIntrospectable returns true if the type inside the result is
// introspectable.
func (res *TypeFindResult) IsIntrospectable() bool {
	switch {
	case res.Alias != nil:
		return res.Alias.IsIntrospectable()
	case res.Class != nil:
		return res.Class.IsIntrospectable()
	case res.Interface != nil:
		return res.Interface.IsIntrospectable()
	case res.Record != nil:
		return res.Record.IsIntrospectable()
	case res.Enum != nil:
		return res.Enum.IsIntrospectable()
	case res.Function != nil:
		return res.Function.IsIntrospectable()
	case res.Union != nil:
		return res.Union.IsIntrospectable()
	case res.Bitfield != nil:
		return res.Bitfield.IsIntrospectable()
	case res.Callback != nil:
		return res.Callback.IsIntrospectable()
	}

	panic("TypeFindResult has all fields nil")
}

// FindInclude returns the namespace that the given namespace includes. It
// resolves imports recursively. This function is primarily used to ensure that
// proper versions are imported.
func (repos Repositories) FindInclude(res *NamespaceFindResult, includes string) *NamespaceFindResult {
	foundIncludes := make([]*NamespaceFindResult, 0, len(res.Repository.Includes))

	for _, incl := range res.Repository.Includes {
		nspIncl := repos.FindNamespace(VersionedName(incl.Name, incl.Version))
		if nspIncl == nil {
			// We've already seen the import and it's not available, so early
			// bail.
			if incl.Name == includes {
				return nil
			}

			continue
		}

		if incl.Name == includes {
			return nspIncl
		}

		foundIncludes = append(foundIncludes, nspIncl)
	}

	for _, nspIncl := range foundIncludes {
		if res := repos.FindInclude(nspIncl, includes); res != nil {
			return res
		}
	}

	return nil
}

var (
	typeResultCache  sync.Map // full GIR type -> *TypeFindResult
	typeResultFlight singleflight.Group
)

// FindType finds a type in the repositories from the given current namespace
// and the GIR type name. The function will cache the returned TypeFindResult.
func (repos Repositories) FindType(nsp *NamespaceFindResult, typ string) *TypeFindResult {
	// Build the full type name for cache querying.
	namespace, typName := SplitGIRType(typ)

	// Ensure the new fullType string has the version.
	if namespace != "" {
		n, version := ParseVersionName(namespace)
		if version != "" {
			goto gotNamespace
		}

		// No versions provided; check if the namespace is the same (redundant).
		if n == nsp.Namespace.Name {
			namespace = VersionedNamespace(nsp.Namespace)
			goto gotNamespace
		}

		if result := repos.FindInclude(nsp, n); result != nil {
			namespace = VersionedNamespace(result.Namespace)
			goto gotNamespace
		}

		return nil
	}

	// Type comes from the same namespace. Append the current one.
	namespace = VersionedNamespace(nsp.Namespace)

gotNamespace:
	fullType := namespace + "." + typName

	v, ok := typeResultCache.Load(fullType)
	if ok {
		return v.(*TypeFindResult)
	}

	v, _, _ = typeResultFlight.Do(fullType, func() (interface{}, error) {
		result := repos.findType(fullType)
		if result != nil {
			typeResultCache.Store(fullType, result)
		}

		return result, nil
	})

	return v.(*TypeFindResult)
}

func (repos Repositories) findType(fullType string) *TypeFindResult {
	vNamespace, typ := SplitGIRType(fullType)

	r := TypeFindResult{NamespaceFindResult: repos.FindNamespace(vNamespace)}
	if r.NamespaceFindResult == nil {
		return nil
	}

	v := SearchNamespace(r.Namespace, func(name, _ string) bool { return name == typ })
	if v == nil {
		return nil
	}

	switch v := v.(type) {
	case *Alias:
		r.Alias = v
	case *Class:
		r.Class = v
	case *Interface:
		r.Interface = v
	case *Record:
		r.Record = v
	case *Enum:
		r.Enum = v
	case *Function:
		r.Function = v
	case *Union:
		r.Union = v
	case *Bitfield:
		r.Bitfield = v
	case *Callback:
		r.Callback = v
	}

	return &r
}

// SearchNamespace searches the namespace for the given type name. The returned
// interface may be any of the types in TypeFindResult.
func SearchNamespace(namespace *Namespace, f func(typ, ctyp string) bool) interface{} {
	for _, alias := range namespace.Aliases {
		if f(alias.Name, alias.CType) {
			return &alias
		}
	}

	for _, class := range namespace.Classes {
		if f(class.Name, class.GLibTypeName) {
			return &class
		}
	}

	for _, enum := range namespace.Enums {
		if f(enum.Name, enum.CType) {
			return &enum
		}
	}

	for _, record := range namespace.Records {
		if f(record.Name, record.CType) {
			return &record
		}
	}

	for _, function := range namespace.Functions {
		if f(function.Name, function.CIdentifier) {
			return &function
		}
	}

	for _, union := range namespace.Unions {
		if f(union.Name, union.CType) {
			return &union
		}
	}

	for _, bitfield := range namespace.Bitfields {
		if f(bitfield.Name, bitfield.CType) {
			return &bitfield
		}
	}

	for _, callback := range namespace.Callbacks {
		if f(callback.Name, callback.CIdentifier) {
			return &callback
		}
	}

	for _, iface := range namespace.Interfaces {
		if f(iface.Name, iface.CType) {
			return &iface
		}
	}

	return nil
}
