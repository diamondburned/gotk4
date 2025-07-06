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
func goToggleNotify(_ C.gpointer, obj *C.GObject, isLastInt C.gboolean) {
	gobject := unsafe.Pointer(obj)
	isLast := isLastInt != C.FALSE

	shared.mu.Lock()
	defer shared.mu.Unlock()

	var box *Box
	if isLast {
		box = makeWeak(gobject)
	} else {
		box = makeStrong(gobject)
	}

	if box == nil {
		if toggleRefs {
			slog.Debug(
				"goToggleNotify: box not found",
				objInfo(unsafe.Pointer(obj)))
		}
		return
	}

	if toggleRefs {
		slog.Debug(
			"goToggleNotify: toggled notify for object",
			"is_last", isLast,
			"finalizing", box.finalizing.Load(),
			objInfo(unsafe.Pointer(obj)))
	}

	if box.finalizing.Load() {
		if !isLast {
			panic("goToggleNotify: object is already finalizing but is not last")
		}

		delete(shared.weak, gobject)

		if toggleRefs {
			slog.Debug(
				"goToggleNotify: object is already finalizing so we removed it",
				"is_last", isLast, // should be true
				objInfo(unsafe.Pointer(obj)))
		}
	}
}

// finishRemovingToggleRef is called after the toggle reference removal routine
// is dispatched in the main loop. It removes the GObject from the global maps.
//
//export goFinishRemovingToggleRef
func goFinishRemovingToggleRef(gobject unsafe.Pointer) {
	/*
		shared.mu.Lock()
		defer shared.mu.Unlock()

		box, strong := gets(gobject)
		if box == nil {
			if toggleRefs {
				slog.Debug(
					"goFinishRemovingToggleRef: object not found in weak map",
					"box", false,
					objInfo(gobject))
			}
			return
		}

		if toggleRefs {
			slog.Debug(
				"goFinishRemovingToggleRef: object found in weak map",
				"box", true,
				objInfo(gobject))
		}

		if strong {
			if toggleRefs {
				slog.Debug(
					"goFinishRemovingToggleRef: object still strong",
					objInfo(gobject))
			}
			return
		}

		if !box.finalize {
			if toggleRefs {
				slog.Debug(
					"goFinishRemovingToggleRef: object resurrected",
					objInfo(gobject))
			}
			return
		}

		delete(shared.weak, gobject)

		if toggleRefs {
			slog.Debug(
				"goFinishRemovingToggleRef: removed from weak ref",
				objInfo(gobject))
		}

		if objectProfile != nil {
			objectProfile.Remove(gobject)
		}
	*/
}
