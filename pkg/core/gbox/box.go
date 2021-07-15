package gbox

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

import (
	"github.com/diamondburned/gotk4/pkg/core/slab"
)

var registry slab.Slab

// Assign assigns the given value and returns the fake pointer.
func Assign(v interface{}) uintptr {
	return registry.Put(v, false)
}

// AssignOnce stores the given value so that, when the value is retrieved, it
// will immediately be deleted.
func AssignOnce(v interface{}) uintptr {
	return registry.Put(v, true)
}

// Get gets the value from the given fake pointer. The context must match the
// given value in Assign.
func Get(ptr uintptr) interface{} {
	return registry.Get(ptr)
}

// Delete deletes a boxed value. It is exposed to C under the name
// "callbackDelete".
func Delete(ptr uintptr) {
	registry.Delete(ptr)
}

//export callbackDelete
func callbackDelete(ptr uintptr) {
	registry.Delete(ptr)
}

// Pop gets a value and deletes it atomically.
func Pop(ptr uintptr) interface{} {
	return registry.Pop(ptr)
}
