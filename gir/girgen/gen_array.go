package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
)

// arrayType generates the Go type signature for the given array.
func (ng *NamespaceGenerator) resolveArrayType(array gir.Array) (string, bool) {
	arrayPrefix := "[]"
	if array.FixedSize > 0 {
		arrayPrefix = fmt.Sprintf("[%d]", array.FixedSize)
	}

	child, _ := ng.resolveAnyType(array.AnyType)
	// There can't be []void, so this check ensures there can only be valid
	// array types.
	if child == "" {
		return "", false
	}

	return arrayPrefix + child, true
}

// anyType generates the Go type signature for the AnyType union. An empty
// string returned is an invalid type.
func (ng *NamespaceGenerator) resolveAnyType(any gir.AnyType) (string, bool) {
	switch {
	case any.Array != nil:
		return ng.resolveArrayType(*any.Array)

	case any.Type != nil:
		if r := ng.resolveType(*any.Type); r != nil {
			return r.GoType, true
		}

	default:
		ng.debugln("anyType missing both array and type")
	}

	return "", false
}
