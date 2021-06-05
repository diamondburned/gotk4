// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_range_accessible_get_type()), F: marshalRangeAccessible},
	})
}

type RangeAccessible interface {
	WidgetAccessible
}

// rangeAccessible implements the RangeAccessible interface.
type rangeAccessible struct {
	WidgetAccessible
}

var _ RangeAccessible = (*rangeAccessible)(nil)

// WrapRangeAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapRangeAccessible(obj *externglib.Object) RangeAccessible {
	return RangeAccessible{
		WidgetAccessible: WrapWidgetAccessible(obj),
	}
}

func marshalRangeAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRangeAccessible(obj), nil
}

type RangeAccessiblePrivate struct {
	native C.GtkRangeAccessiblePrivate
}

// WrapRangeAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRangeAccessiblePrivate(ptr unsafe.Pointer) *RangeAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*RangeAccessiblePrivate)(ptr)
}

func marshalRangeAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRangeAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RangeAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}