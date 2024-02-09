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

	box, strong := gets(gobject)
	if box == nil {
		// Extremely weird error. This should never happen.
		log.Printf(
			"gotk4: critical: %p: finishRemovingToggleRef called on unknown object",
			gobject)
		return
	}

	if strong {
		// Panic here, else we're memory leaking.
		log.Panicf(
			"gotk4: critical: %p: finishRemovingToggleRef cannot be called on strongly-referenced object (unexpectedly resurrected?)",
			gobject)
	}

	if !box.done {
		log.Panicf(
			"gotk4: critical: %p: finishRemovingToggleRef cannot be called with finalizer still set",
			gobject)
	}

	// If the closures are weak-referenced, then the object reference hasn't
	// been toggled yet. Since the object is going away and we're still
	// weakly referenced, we can wipe the closures away.
	//
	// Finally clear the object data off the registry.
	delete(shared.weak, gobject)

	// Clear the finalizer.
	runtime.SetFinalizer(box.dummy, nil)

	// Keep the box alive until the end of the function just in case the
	// finalizer is called again.
	runtime.KeepAlive(box.dummy)

	if objectProfile != nil {
		objectProfile.Remove(gobject)
	}
}
