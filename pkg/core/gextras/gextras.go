// Package gextras contains supplemental types to gotk3.
package gextras

import (
	"unsafe"
)

type record struct{ intern *internRecord }

type internRecord struct{ c unsafe.Pointer }

// StructNative returns the underlying C pointer of the given Go record struct
// pointer. It can be used like so:
//
//    rec := NewRecord(...) // T = *Record
//    c := (*namespace_record)(StructPtr(unsafe.Pointer(rec)))
//
func StructNative(ptr unsafe.Pointer) unsafe.Pointer {
	return (*record)(ptr).intern.c
}

// StructIntern returns the given struct's internal struct pointer.
func StructIntern(ptr unsafe.Pointer) *struct{ C unsafe.Pointer } {
	return (*struct{ C unsafe.Pointer })(unsafe.Pointer((*record)(ptr).intern))
}

// SetStructNative sets the native value inside the Go struct value that the
// given dst pointer points to. It can be used like so:
//
//    var rec Record
//    SetStructNative(&rec, cvalue) // T(cvalue) = *namespace_record
//
func SetStructNative(dst, native unsafe.Pointer) {
	(*record)(dst).intern.c = native
}

// NewStructNative creates a new Go struct from the given native pointer. The
// finalizer is NOT set.
func NewStructNative(native unsafe.Pointer) unsafe.Pointer {
	r := record{intern: &internRecord{native}}
	return unsafe.Pointer(&r)
}
