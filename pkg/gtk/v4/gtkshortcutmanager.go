// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_ShortcutManagerInterface_add_controller(GtkShortcutManager*, GtkShortcutController*);
// extern void _gotk4_gtk4_ShortcutManagerInterface_remove_controller(GtkShortcutManager*, GtkShortcutController*);
import "C"

// GType values.
var (
	GTypeShortcutManager = coreglib.Type(C.gtk_shortcut_manager_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutManager, F: marshalShortcutManager},
	})
}

// ShortcutManagerOverrider contains methods that are overridable.
type ShortcutManagerOverrider interface {
	// The function takes the following parameters:
	//
	AddController(controller *ShortcutController)
	// The function takes the following parameters:
	//
	RemoveController(controller *ShortcutController)
}

// ShortcutManager: GtkShortcutManager interface is used to implement shortcut
// scopes.
//
// This is important for gtk.Native widgets that have their own surface, since
// the event controllers that are used to implement managed and global scopes
// are limited to the same native.
//
// Examples for widgets implementing GtkShortcutManager are gtk.Window and
// gtk.Popover.
//
// Every widget that implements GtkShortcutManager will be used as a
// GTK_SHORTCUT_SCOPE_MANAGED.
//
// ShortcutManager wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ShortcutManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ShortcutManager)(nil)
)

// ShortcutManagerer describes ShortcutManager's interface methods.
type ShortcutManagerer interface {
	coreglib.Objector

	baseShortcutManager() *ShortcutManager
}

var _ ShortcutManagerer = (*ShortcutManager)(nil)

func ifaceInitShortcutManagerer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkShortcutManagerInterface)(unsafe.Pointer(gifacePtr))
	iface.add_controller = (*[0]byte)(C._gotk4_gtk4_ShortcutManagerInterface_add_controller)
	iface.remove_controller = (*[0]byte)(C._gotk4_gtk4_ShortcutManagerInterface_remove_controller)
}

//export _gotk4_gtk4_ShortcutManagerInterface_add_controller
func _gotk4_gtk4_ShortcutManagerInterface_add_controller(arg0 *C.GtkShortcutManager, arg1 *C.GtkShortcutController) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ShortcutManagerOverrider)

	var _controller *ShortcutController // out

	_controller = wrapShortcutController(coreglib.Take(unsafe.Pointer(arg1)))

	iface.AddController(_controller)
}

//export _gotk4_gtk4_ShortcutManagerInterface_remove_controller
func _gotk4_gtk4_ShortcutManagerInterface_remove_controller(arg0 *C.GtkShortcutManager, arg1 *C.GtkShortcutController) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(ShortcutManagerOverrider)

	var _controller *ShortcutController // out

	_controller = wrapShortcutController(coreglib.Take(unsafe.Pointer(arg1)))

	iface.RemoveController(_controller)
}

func wrapShortcutManager(obj *coreglib.Object) *ShortcutManager {
	return &ShortcutManager{
		Object: obj,
	}
}

func marshalShortcutManager(p uintptr) (interface{}, error) {
	return wrapShortcutManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *ShortcutManager) baseShortcutManager() *ShortcutManager {
	return self
}

// BaseShortcutManager returns the underlying base object.
func BaseShortcutManager(obj ShortcutManagerer) *ShortcutManager {
	return obj.baseShortcutManager()
}
