// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeCheckMenuItemAccessible = coreglib.Type(C.gtk_check_menu_item_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCheckMenuItemAccessible, F: marshalCheckMenuItemAccessible},
	})
}

// CheckMenuItemAccessibleOverrider contains methods that are overridable.
type CheckMenuItemAccessibleOverrider interface {
}

type CheckMenuItemAccessible struct {
	_ [0]func() // equal guard
	MenuItemAccessible
}

var (
	_ coreglib.Objector = (*CheckMenuItemAccessible)(nil)
)

func classInitCheckMenuItemAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCheckMenuItemAccessible(obj *coreglib.Object) *CheckMenuItemAccessible {
	return &CheckMenuItemAccessible{
		MenuItemAccessible: MenuItemAccessible{
			ContainerAccessible: ContainerAccessible{
				WidgetAccessible: WidgetAccessible{
					Accessible: Accessible{
						ObjectClass: atk.ObjectClass{
							Object: obj,
						},
					},
					Component: atk.Component{
						Object: obj,
					},
				},
			},
			Object: obj,
			Action: atk.Action{
				Object: obj,
			},
			Selection: atk.Selection{
				Object: obj,
			},
		},
	}
}

func marshalCheckMenuItemAccessible(p uintptr) (interface{}, error) {
	return wrapCheckMenuItemAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
