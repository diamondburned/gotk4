package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
)

// arrayType generates the Go type signature for the given array.
func (ng *NamespaceGenerator) resolveArrayType(array gir.Array) string {
	arrayPrefix := "[]"
	if array.FixedSize > 0 {
		arrayPrefix = fmt.Sprintf("[%d]", array.FixedSize)
	}

	child := ng.resolveAnyType(array.AnyType)
	if child == "" {
		return ""
	}

	return arrayPrefix + child
}

// anyType generates the Go type signature for the AnyType union.
func (ng *NamespaceGenerator) resolveAnyType(any gir.AnyType) string {
	switch {
	case any.Array != nil:
		return ng.resolveArrayType(*any.Array)

	case any.Type != nil:
		if r := ng.resolveType(*any.Type); r != nil {
			return r.GoType
		}
		return ""

	default:
		ng.gen.debugln("anyType missing both array and type")
		return ""
	}
}
