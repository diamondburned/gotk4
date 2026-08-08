package intern

// #cgo pkg-config: gobject-2.0
// #include "intern.h"
import "C"

import (
	"log/slog"
	"unsafe"
)

// goToggleNotify is called by GLib on each toggle notification. It doesn't
// actually free anything and relies on Box's finalizer to free both the box and
// the C GObject.
//
//export goToggleNotify
func goToggleNotify(data C.guintptr, obj *C.GObject, isLastInt C.gboolean) {
	toggleNotify(uintptr(data), unsafe.Pointer(obj), isLastInt != C.FALSE)
}

func toggleNotify(token uintptr, gobject unsafe.Pointer, isLast bool) {

	shared.mu.Lock()
	e := shared.byToken[token]
	if e == nil || e.object != gobject || e.phase == entryRemoving {
		shared.mu.Unlock()
		return
	}
	if e.phase == entryCleanupQueued && isLast {
		// The finalizer already owns the queued removal and there is no strong
		// Box to re-weakify. A later non-last notification may still promote
		// the entry through makeStrong.
		shared.mu.Unlock()
		return
	}
	var box *Box
	if isLast {
		box = makeWeak(e)
	} else {
		box = makeStrong(e)
	}
	shared.mu.Unlock()

	if box == nil {
		if toggleRefs {
			slog.Debug(
				"goToggleNotify: box not found",
				"object", gobject)
		}
		return
	}

	if toggleRefs {
		slog.Debug(
			"goToggleNotify: finished",
			"is_last", isLast,
			"generation", box.generation,
			objInfo(gobject))
	}

}

// finishRemovingToggleRef is called after the toggle reference removal routine
// has returned. It removes the GObject from the global maps. The toggle ref
// was the only reference the box owned, so no additional unref is needed.
//
//export goFinishRemovingToggleRef
func goFinishRemovingToggleRef(token C.guintptr) {
	finishCleanup(uintptr(token))
}

//export goRemoveQueuedToggleRef
func goRemoveQueuedToggleRef(token C.guintptr) {
	beginQueuedCleanup(uintptr(token))
}
