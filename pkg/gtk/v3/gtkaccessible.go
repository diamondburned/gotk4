// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_AccessibleClass_connect_widget_destroyed(GtkAccessible*);
// extern void _gotk4_gtk3_AccessibleClass_widget_set(GtkAccessible*);
// extern void _gotk4_gtk3_AccessibleClass_widget_unset(GtkAccessible*);
import "C"

// GType values.
var (
	GTypeAccessible = coreglib.Type(C.gtk_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAccessible, F: marshalAccessible},
	})
}

// AccessibleOverrider contains methods that are overridable.
type AccessibleOverrider interface {
	// ConnectWidgetDestroyed: this function specifies the callback function to
	// be called when the widget corresponding to a GtkAccessible is destroyed.
	//
	// Deprecated: Use gtk_accessible_set_widget() and its vfuncs.
	ConnectWidgetDestroyed()
	WidgetSet()
	WidgetUnset()
}

// Accessible class is the base class for accessible implementations for Widget
// subclasses. It is a thin wrapper around Object, which adds facilities for
// associating a widget with its accessible object.
//
// An accessible implementation for a third-party widget should derive from
// Accessible and implement the suitable interfaces from ATK, such as Text or
// Selection. To establish the connection between the widget class and its
// corresponding acccessible implementation, override the get_accessible vfunc
// in WidgetClass.
type Accessible struct {
	_ [0]func() // equal guard
	atk.ObjectClass
}

var (
	_ coreglib.Objector = (*Accessible)(nil)
)

func classInitAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkAccessibleClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ ConnectWidgetDestroyed() }); ok {
		pclass.connect_widget_destroyed = (*[0]byte)(C._gotk4_gtk3_AccessibleClass_connect_widget_destroyed)
	}

	if _, ok := goval.(interface{ WidgetSet() }); ok {
		pclass.widget_set = (*[0]byte)(C._gotk4_gtk3_AccessibleClass_widget_set)
	}

	if _, ok := goval.(interface{ WidgetUnset() }); ok {
		pclass.widget_unset = (*[0]byte)(C._gotk4_gtk3_AccessibleClass_widget_unset)
	}
}

//export _gotk4_gtk3_AccessibleClass_connect_widget_destroyed
func _gotk4_gtk3_AccessibleClass_connect_widget_destroyed(arg0 *C.GtkAccessible) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ConnectWidgetDestroyed() })

	iface.ConnectWidgetDestroyed()
}

//export _gotk4_gtk3_AccessibleClass_widget_set
func _gotk4_gtk3_AccessibleClass_widget_set(arg0 *C.GtkAccessible) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ WidgetSet() })

	iface.WidgetSet()
}

//export _gotk4_gtk3_AccessibleClass_widget_unset
func _gotk4_gtk3_AccessibleClass_widget_unset(arg0 *C.GtkAccessible) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ WidgetUnset() })

	iface.WidgetUnset()
}

func wrapAccessible(obj *coreglib.Object) *Accessible {
	return &Accessible{
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
	}
}

func marshalAccessible(p uintptr) (interface{}, error) {
	return wrapAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectWidgetDestroyed: this function specifies the callback function to be
// called when the widget corresponding to a GtkAccessible is destroyed.
//
// Deprecated: Use gtk_accessible_set_widget() and its vfuncs.
func (accessible *Accessible) ConnectWidgetDestroyed() {
	var _arg0 *C.GtkAccessible // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))

	C.gtk_accessible_connect_widget_destroyed(_arg0)
	runtime.KeepAlive(accessible)
}

// Widget gets the Widget corresponding to the Accessible. The returned widget
// does not have a reference added, so you do not need to unref it.
//
// The function returns the following values:
//
//    - widget (optional): pointer to the Widget corresponding to the Accessible,
//      or NULL.
//
func (accessible *Accessible) Widget() Widgetter {
	var _arg0 *C.GtkAccessible // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))

	_cret = C.gtk_accessible_get_widget(_arg0)
	runtime.KeepAlive(accessible)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetWidget sets the Widget corresponding to the Accessible.
//
// accessible will not hold a reference to widget. It is the caller???s
// responsibility to ensure that when widget is destroyed, the widget is unset
// by calling this function again with widget set to NULL.
//
// The function takes the following parameters:
//
//    - widget (optional) or NULL to unset.
//
func (accessible *Accessible) SetWidget(widget Widgetter) {
	var _arg0 *C.GtkAccessible // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkAccessible)(unsafe.Pointer(coreglib.InternObject(accessible).Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	C.gtk_accessible_set_widget(_arg0, _arg1)
	runtime.KeepAlive(accessible)
	runtime.KeepAlive(widget)
}
