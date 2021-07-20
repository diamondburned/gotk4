// Package gextras contains supplemental types to gotk3.
package gextras

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <gmodule.h> // HashTable
// #include <glib-object.h>
import "C"

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// Nativer is an interface for an object that can return its native pointer.
type Nativer interface {
	Native() uintptr
}

// Objector is an interface that describes partially the glib.Object type.
type Objector interface {
	Nativer

	Connect(string, interface{}) glib.SignalHandle
	ConnectAfter(string, interface{}) glib.SignalHandle

	HandlerBlock(glib.SignalHandle)
	HandlerDisconnect(glib.SignalHandle)
	HandlerUnblock(glib.SignalHandle)

	GetProperty(string) (interface{}, error)
	SetProperty(string, interface{}) error
}

var _ Objector = (*glib.Object)(nil)

// MustSet panics if the given property cannot be set into the given object.
func MustSet(obj Objector, k string, v interface{}) {
	if err := obj.SetProperty(k, v); err != nil {
		log.Panicf("cannot set object property %q: %s", k, err)
	}
}

// MustGet panics if the given property cannot be retrieved from the given
// object.
func MustGet(obj Objector, k string) interface{} {
	v, err := obj.GetProperty(k)
	if err != nil {
		log.Panicf("cannot get object property %q: %s", k, err)
	}
	return v
}

// GetInto gets the given property key from the object into the given pointer.
// An error is returned if it cannot get the property or the type is wrong. This
// method is mostly useful for avoiding type assertions.
func GetInto(obj Objector, k string, ptr interface{}) error {
	return getInto(obj, k, ptr, false)
}

// MustGetInto is similar to GetInfo, except it does not do safety checks and
// will panic on an error. Code that uses constants should use this function
// over GetInto.
func MustGetInto(obj Objector, k string, ptr interface{}) {
	getInto(obj, k, ptr, true)
}

func getInto(obj Objector, k string, ptr interface{}, must bool) error {
	dst := reflect.ValueOf(ptr)
	if !must {
		typ := dst.Type()
		if !(typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Interface) {
			return fmt.Errorf("ptr type %s is not a pointer", typ)
		}
	}

	elem := dst.Elem()
	if !must {
		if !elem.CanSet() {
			return errors.New("ptr's value cannot be set")
		}
	}

	v, err := obj.GetProperty(k)
	if err != nil {
		if !must {
			return fmt.Errorf("cannot get object property %q: %s", k, err)
		}

		log.Panicf("cannot get object property %q: %s", k, err)
	}

	val := reflect.ValueOf(v)

	if !must {
		property := val.Type()
		given := elem.Type()

		if !property.AssignableTo(given) {
			return fmt.Errorf("property type %s not assignable to given %s", property, given)
		}
	}

	elem.Set(val)
	return nil
}

// InternObject gets the internal Object type. This is used for calling methods
// not in the Objector.
func InternObject(nativer Nativer) *glib.Object {
	obj := glib.Object{
		GObject: glib.ToGObject(unsafe.Pointer(nativer.Native())),
	}

	if typ := obj.TypeFromInstance(); typ != glib.TYPE_OBJECT {
		log.Panicf("InternObject cast: expected type object, got %v", typ)
	}

	return &obj
}

// CastObject casts the given object pointer to the class name. The caller is
// responsible for recasting the interface to the wanted type.
func CastObject(obj *glib.Object) interface{} {
	var gvalue C.GValue

	C.g_value_init_from_instance(&gvalue, C.gpointer(unsafe.Pointer(obj.GObject)))
	defer C.g_value_unset(&gvalue)

	v, err := glib.ValueFromNative(unsafe.Pointer(&gvalue)).GoValue()
	if err != nil {
		return Objector(obj)
	}

	return v
}

type record struct {
	_ NoCopy
	p unsafe.Pointer
}

// StructNative returns the underlying C pointer of the given Go record struct
// pointer. It can be used like so:
//
//    rec := NewRecord(...) // T = *Record
//    c := (*namespace_record)(StructPtr(unsafe.Pointer(rec)))
//
func StructNative(ptr unsafe.Pointer) unsafe.Pointer {
	return (*record)(ptr).p
}

// SetStructNative sets the native value inside the Go struct value that the
// given dst pointer points to. It can be used like so:
//
//    var rec Record
//    SetStructNative(&rec, cvalue) // T(cvalue) = *namespace_record
//
func SetStructNative(dst, native unsafe.Pointer) {
	(*record)(dst).p = native
}

// NewStructNative creates a new Go struct from the given native pointer. The
// finalizer is NOT set.
func NewStructNative(native unsafe.Pointer) unsafe.Pointer {
	var r record
	(*record)(&r).p = native
	return unsafe.Pointer(&r)
}

// NoCopy is a zero-sized type that triggers a warning in go vet when a struct
// containing this type is copied.
//
// See https://github.com/golang/go/issues/8005#issuecomment-190753527.
type NoCopy struct{}

func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}

// HashTableSize returns the size of the *GHashTable.
func HashTableSize(ptr unsafe.Pointer) int {
	return int(C.g_hash_table_size((*C.GHashTable)(ptr)))
}

// FreeHashTable frees the given hash table.
func FreeHashTable(ptr unsafe.Pointer) {
	C.g_hash_table_unref((*C.GHashTable)(ptr))
}

// MoveHashTable calls f on every value of the given *GHashTable and frees each
// element in the process if rm is true.
func MoveHashTable(ptr unsafe.Pointer, rm bool, f func(k, v unsafe.Pointer)) {
	var k, v C.gpointer
	var iter C.GHashTableIter
	C.g_hash_table_iter_init(&iter, (*C.GHashTable)(ptr))

	for C.g_hash_table_iter_next(&iter, &k, &v) != 0 {
		f(unsafe.Pointer(k), unsafe.Pointer(v))

		if rm {
			C.g_hash_table_iter_remove(&iter)
		}
	}
}
