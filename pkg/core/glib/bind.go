package glib

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

var bindingNames sync.Map // map[reflect.Type]C.GQuark

// Associate value with object
func Bind[T any](obj Objector, value T) {
	object := BaseObject(obj)
	name := bindingName[T]()

	ptr := C.gpointer(gbox.Assign(value))

	C.g_object_set_data_full(object.native(), (*C.gchar)(name), ptr, (*[0]byte)(C._gotk4_data_destroy))
}

// Disassociate value from object
func Unbind[T any](obj Objector) {
	name := bindingName[T]()

	ptr := C.g_object_steal_data(BaseObject(obj).native(), (*C.gchar)(name))
	defer gbox.Delete(uintptr(ptr))
}

// Obtain value associated with object
func Bounded[T any](obj Objector) *T {
	name := bindingName[T]()

	ptr := C.g_object_get_data(BaseObject(obj).native(), name)

	value, ok := gbox.Get(uintptr(ptr)).(T)
	if !ok {
		return nil
	}

	return &value
}

func bindingName[T any]() *C.gchar {
	t := reflect.TypeFor[T]()

	if v, ok := bindingNames.Load(t); ok {
		quark := v.(C.GQuark)
		return C.g_quark_to_string(quark)
	}

	name := "_gotk4_" + t.String()

	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	quark := C.g_quark_from_string(nameC)
	if v, lost := bindingNames.LoadOrStore(t, quark); lost {
		quark = v.(C.GQuark)
	}

	return C.g_quark_to_string(quark)
}

//export _gotk4_data_destroy
func _gotk4_data_destroy(ptr C.gpointer) {
	gbox.Delete(uintptr(ptr))
}
