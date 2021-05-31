// Package gextras contains supplemental types to gotk3.
package gextras

// #cgo pkg-config: glib-2.0 gobject-2.0
// #include <glib.h>
// #include <glib-object.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// Objector is an interface that describes partially the glib.Object type.
type Objector interface {
	Connect(string, interface{}) glib.SignalHandle
	ConnectAfter(string, interface{}) glib.SignalHandle
	Emit(string, ...interface{}) (interface{}, error)
	HandlerBlock(glib.SignalHandle)
	HandlerDisconnect(glib.SignalHandle)
	HandlerUnblock(glib.SignalHandle)
	IsA(glib.Type) bool
	Native() uintptr
	GetProperty(string) (interface{}, error)
	SetProperty(string, interface{}) error
	StopEmission(string)
	TypeFromInstance() glib.Type
}

var _ Objector = (*glib.Object)(nil)

// CastObject casts the given object pointer to the class name. The caller is
// responsible for recasting the interface to the wanted type.
func CastObject(obj *glib.Object) interface{} {
	var gvalue C.GValue

	C.g_value_init(&gvalue, C.GType(obj.TypeFromInstance()))
	defer C.g_value_unset(&gvalue)

	v, err := glib.ValueFromNative(unsafe.Pointer(&gvalue)).GoValue()
	if err != nil {
		return Objector(obj)
	}

	return v
}
