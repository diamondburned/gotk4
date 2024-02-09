package intern

// #cgo pkg-config: gobject-2.0
// #include "intern.h"
import "C"

import (
	"log"
	"runtime"
	"unsafe"
)

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
	defer shared.mu.Unlock()

	if isLast {
		// delete(shared.sharing, gobject)
		makeWeak(gobject)
	} else {
		// shared.sharing[gobject] = struct{}{}
		makeStrong(gobject)
	}

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(unsafe.Pointer(obj)), "goToggleNotify: is last =", isLast)
	}
}

// finishRemovingToggleRef is called after the toggle reference removal routine
// is dispatched in the main loop. It removes the GObject from the strong and
// weak global maps and unsets the finalizer.
//
//go:nosplit
//export goFinishRemovingToggleRef
func goFinishRemovingToggleRef(gobject unsafe.Pointer) {
	if toggleRefs != nil {
		toggleRefs.Printf("goFinishRemovingToggleRef: called on %p", gobject)
	}

	shared.mu.Lock()
	defer shared.mu.Unlock()

	box, ok := shared.finalizing[gobject]
	if !ok {
		// Extremely weird error. This should never happen.
		log.Printf(
			"gotk4: critical: %p: finishRemovingToggleRef called on unknown object",
			gobject)
		return
	}

	// Clear the finalizer.
	// runtime.SetFinalizer(box.dummy, nil)

	// Finally clear the object data off the registry.
	delete(shared.finalizing, gobject)

	// Keep the box alive until the end of the function just in case the
	// finalizer is called again.
	runtime.KeepAlive(box.dummy)
	runtime.KeepAlive(box)

	if toggleRefs != nil {
		toggleRefs.Printf("goFinishRemovingToggleRef: removed %p", gobject)
	}

	if objectProfile != nil {
		objectProfile.Remove(gobject)
	}
}
