package intern

// #cgo pkg-config: gobject-2.0
// #include <glib-object.h>
import "C"

import "unsafe"

// goToggleNotify is called by GLib on each toggle notification. It doesn't
// actually free anything and relies on Box's finalizer to free both the box and
// the C GObject.
//
//go:nosplit
//export goToggleNotify
func goToggleNotify(_ C.gpointer, obj *C.GObject, isLastInt C.gboolean) {
	gobject := unsafe.Pointer(obj)
	isLast := isLastInt != C.FALSE

	shared.mu.Lock()

	if isLast {
		// delete(shared.sharing, gobject)
		makeWeak(gobject)
	} else {
		// shared.sharing[gobject] = struct{}{}
		makeStrong(gobject)
	}

	shared.mu.Unlock()

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(unsafe.Pointer(obj)), "goToggleNotify: is last =", isLast)
	}
}
