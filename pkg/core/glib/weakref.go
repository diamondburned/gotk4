package glib

// #include <glib.h>
// #include <glib-object.h>
//
// extern void _gotk4_glib_weak_notify(gpointer, GObject*);
import "C"

import "github.com/diamondburned/gotk4/pkg/core/gbox"

// WeakRefObject is like SetFinalizer, except it's not thread-safe (so notify
// SHOULD NOT REFERENCE OBJECT). It is best that you just do not use this at
// all.
func WeakRefObject(obj Objector, notify func()) {
	data := gbox.AssignOnce(notify)
	C.g_object_weak_ref(
		BaseObject(obj).native(),
		C.GWeakNotify(C._gotk4_glib_weak_notify),
		C.gpointer(data))
}

//export _gotk4_glib_weak_notify
func _gotk4_glib_weak_notify(data C.gpointer, _ *C.GObject) {
	notify := gbox.Get(uintptr(data)).(func())
	notify()
}
