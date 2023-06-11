// Package gir provides GIR types, as well as functions that parse GIR files.
//
// For reference, see
// https://gitlab.gnome.org/GNOME/gobject-introspection/-/blob/HEAD/docs/gir-1.2.rnc.
package gir

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strconv"
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

// ParseVersionName parses the given full namespace to return the original name
// and the version separately. If no versions are available, then the empty
// string is returned.
func ParseVersionName(fullNamespace string) (name, majorVersion string) {
	parts := strings.SplitN(fullNamespace, "-", 2)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}

	// Trim the minor versions off.
	version := MajorVersion(parts[1])

	// Verify the number is valid.
	_, err := strconv.Atoi(version)
	if err != nil {
		log.Panicf("version %q is invalid int", version)
	}

	return parts[0], version
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
func (repos *Repositories) AddSelected(pkg string, wantedNames []string) error {
	found := 0

	filter := func(r *Repository) bool {
		namespaces := r.Namespaces
		r.Namespaces = namespaces[:0]

		for _, namespace := range namespaces {
			vname := VersionedNamespace(&namespace)

			for _, wantedName := range wantedNames {
				if wantedName != vname {
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

	if found != len(wantedNames) {
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
				if addingNsp.Name == repoNsp.Name && EqVersion(addingNsp.Version, repoNsp.Version) {
					return nil
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

// FromGIRFile finds the repository from the given GIR filename.
func (repos Repositories) FromGIRFile(girFile string) *PkgRepository {
	for i, repo := range repos {
		if filepath.Base(repo.Path) == girFile {
			return &repos[i]
		}
	}
	return nil
}

// FromPkg finds the repository from the given package name.
func (repos Repositories) FromPkg(pkg string) []PkgRepository {
	var found []PkgRepository
	for i, repo := range repos {
		if repo.Pkg == pkg {
			found = append(found, repos[i])
		}
	}
	return found
}

// NamespaceFindResult is the result returned from FindNamespace.
type NamespaceFindResult struct {
	Repository *PkgRepository
	Namespace  *Namespace
}

// Versioned returns the versioned namespace name.
func (res *NamespaceFindResult) Versioned() string {
	return VersionedNamespace(res.Namespace)
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

	// Types:
	//   *Alias
	//   *Class
	//   *Interface
	//   *Record
	//   *Enum
	//   *Function
	//   *Union
	//   *Bitfield
	//   *Callback
	Type interface{}

	// TODO: Constant, Annotations, Boxed
	// TODO: Methods
}

// CType returns the resulting type's C type identifier.
func (res *TypeFindResult) CType() string {
	return *res.cTypePtr()
}

// SetCType sets the type's name. Note that this will change the type inside the
// repositories as well.
func (res *TypeFindResult) SetCType(ctype string) {
	*res.cTypePtr() = ctype
}

func (res *TypeFindResult) cTypePtr() *string {
	switch v := res.Type.(type) {
	case *Class:
		if v.CType != "" {
			return &v.CType
		}
		return &v.GLibTypeName

	case *Interface:
		if v.CType != "" {
			return &v.CType
		}
		return &v.GLibTypeName

	case *Alias:
		return &v.CType
	case *Record:
		return &v.CType
	case *Enum:
		return &v.CType
	case *Function:
		return &v.CIdentifier
	case *Union:
		return &v.CType
	case *Bitfield:
		return &v.CType
	case *Callback:
		return &v.CIdentifier
	}

	panic("TypeFindResult has all fields nil")
}

func (res *TypeFindResult) namePtr() *string {
	var typ *string
	switch v := res.Type.(type) {
	case *Alias:
		typ = &v.Name
	case *Class:
		typ = &v.Name
	case *Interface:
		typ = &v.Name
	case *Record:
		typ = &v.Name
	case *Enum:
		typ = &v.Name
	case *Function:
		typ = &v.Name
	case *Union:
		typ = &v.Name
	case *Bitfield:
		typ = &v.Name
	case *Callback:
		typ = &v.Name
	}

	if typ == nil {
		panic("TypeFindResult has all fields nil")
	}

	return typ
}

// Name returns the copy of the type's name.
func (res *TypeFindResult) Name() string {
	return *res.namePtr()
}

// SetName sets the type's name. Note that this will change the type inside the
// repositories as well.
func (res *TypeFindResult) SetName(name string) {
	*res.namePtr() = name
}

// NamespacedType returns the copy of the type's name with the namespace
// prepended to it.
func (res *TypeFindResult) NamespacedType() string {
	return res.Namespace.Name + "." + res.Name()
}

// VersionedNamespaceType is like NamespacedType, but the returned type has the
// namespace and its version.
func (res *TypeFindResult) VersionedNamespaceType() string {
	return VersionedNamespace(res.Namespace) + "." + res.Name()
}

// IsIntrospectable returns true if the type inside the result is
// introspectable.
func (res *TypeFindResult) IsIntrospectable() bool {
	if res.Type == nil {
		panic("TypeFindResult has all fields nil")
	}

	type isIntrospectabler interface {
		IsIntrospectable() bool
	}

	return res.Type.(isIntrospectabler).IsIntrospectable()
}

// MustFindGIR generates a girepository.MustFind call for the TypeFindResult.
func (res *TypeFindResult) MustFindGIR() string {
	return fmt.Sprintf("girepository.MustFind(%q, %q)", res.Namespace.Name, res.Name())
}

// FindInclude returns the namespace that the given namespace includes. It
// resolves imports recursively. This function is primarily used to ensure that
// proper versions are imported.
func (repos Repositories) FindInclude(
	res *NamespaceFindResult, includes string) *NamespaceFindResult {

	for _, incl := range res.Repository.Includes {
		if incl.Name != includes {
			continue
		}

		nspIncl := repos.FindNamespace(VersionedName(incl.Name, incl.Version))
		if nspIncl == nil {
			// Include found but not the namespace, so it's probably not added
			// at all.
			return nil
		}

		return nspIncl
	}

	for _, incl := range res.Repository.Includes {
		nspIncl := repos.FindNamespace(VersionedName(incl.Name, incl.Version))
		if nspIncl == nil {
			continue
		}

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

		// log.Panicf("namespace %q of type %q in %q not found", namespace, typ, nsp.Namespace.Name)
		return nil
	}

	// Type comes from the same namespace. Append the current one.
	namespace = VersionedNamespace(nsp.Namespace)

gotNamespace:
	return repos.FindFullType(namespace + "." + typName)
}

// FindFullType finds a type from the given fullType string. The fullType string
// MUST have the namespace version, suck as Gdk-2.Item.
func (repos Repositories) FindFullType(fullType string) *TypeFindResult {
	v, ok := typeResultCache.Load(fullType)
	if ok {
		return v.(*TypeFindResult)
	}

	v, _, _ = typeResultFlight.Do(fullType, func() (interface{}, error) {
		result := repos.findFullType(fullType)
		if result != nil {
			typeResultCache.Store(fullType, result)
		}

		return result, nil
	})

	return v.(*TypeFindResult)
}

func (repos Repositories) findFullType(fullType string) *TypeFindResult {
	vNamespace, typ := SplitGIRType(fullType)
	if vNamespace == "" {
		log.Panicf("type %q empty namespace", fullType)
	}

	r := TypeFindResult{NamespaceFindResult: repos.FindNamespace(vNamespace)}
	if r.NamespaceFindResult == nil {
		return nil
	}

	v := SearchNamespace(r.Namespace, func(name, _ string) bool { return name == typ })
	if v == nil {
		return nil
	}

	r.Type = v
	return &r
}

// SearchNamespace searches the namespace for the given type name. The returned
// interface may be any of the types in TypeFindResult.
func SearchNamespace(namespace *Namespace, f func(typ, ctyp string) bool) interface{} {
	for i, alias := range namespace.Aliases {
		if f(alias.Name, alias.CType) {
			return &namespace.Aliases[i]
		}
	}

	for i, class := range namespace.Classes {
		if f(class.Name, class.GLibTypeName) {
			return &namespace.Classes[i]
		}
	}

	for i, enum := range namespace.Enums {
		if f(enum.Name, enum.CType) {
			return &namespace.Enums[i]
		}
	}

	for i, record := range namespace.Records {
		if f(record.Name, record.CType) {
			return &namespace.Records[i]
		}
	}

	for i, function := range namespace.Functions {
		if f(function.Name, function.CIdentifier) {
			return &namespace.Functions[i]
		}
	}

	for i, union := range namespace.Unions {
		if f(union.Name, union.CType) {
			return &namespace.Unions[i]
		}
	}

	for i, bitfield := range namespace.Bitfields {
		if f(bitfield.Name, bitfield.CType) {
			return &namespace.Bitfields[i]
		}
	}

	for i, callback := range namespace.Callbacks {
		if f(callback.Name, callback.CIdentifier) {
			return &namespace.Callbacks[i]
		}
	}

	for i, iface := range namespace.Interfaces {
		if f(iface.Name, iface.CType) {
			return &namespace.Interfaces[i]
		}
	}

	return nil
}
