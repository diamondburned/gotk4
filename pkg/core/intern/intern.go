// Package intern implements value interning for Cgo sharing.
package intern

// #cgo pkg-config: gobject-2.0
// #include "intern.h"
import "C"

import (
	"log/slog"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"sync"
	"sync/atomic"
	"unsafe"
	"weak"

	"github.com/diamondburned/gotk4/pkg/core/closure"
	"github.com/diamondburned/gotk4/pkg/core/gdebug"

	// Require a non-moving GC for heap pointers. Current GC is moving only by
	// the stack. See https://github.com/go4org/intern.
	_ "go4.org/unsafe/assume-no-moving-gc"
)

// Box is an opaque type holding extra data.
type Box struct {
	closures   atomic.Pointer[closure.Registry]
	gobject    unsafe.Pointer
	finalizing atomic.Bool
}

// Object returns Box's C GObject pointer.
func (b *Box) GObject() unsafe.Pointer {
	return b.gobject
}

// Closures returns the closure registry for this Box.
func (b *Box) Closures() *closure.Registry {
	closures := b.closures.Load()
	if closures != nil {
		return closures
	}

	// If the closures are nil, then we'll have to create a new one.
	closures = closure.NewRegistry()
	if !b.closures.CompareAndSwap(nil, closures) {
		// If the CAS failed, then we'll read the value again.
		return b.closures.Load()
	}

	return closures
}

// Hack to force an object on the heap.
var never bool
var sink_ interface{}

//go:nosplit
func sink(v interface{}) {
	if never {
		sink_ = v
	}
}

var (
	traceObjects  = gdebug.HasKey("trace-objects")
	toggleRefs    = gdebug.HasKey("toggle-refs")
	objectProfile *pprof.Profile
)

func init() {
	if gdebug.HasKey("profile-objects") {
		objectProfile = pprof.NewProfile("gotk4-object-box")
	}
}

func objInfo(obj unsafe.Pointer) slog.Attr {
	return gdebug.ObjectInfo(obj)
}

/*
func boxFinalizerKeepAlive(box *Box) {
	shared.mu.RLock()
	defer shared.mu.RUnlock()

	if resurrected, _ := gets(box.gobject); resurrected != nil {
		// We managed to resurrect a weakly referenced object, so it's still
		// being used somewhere. We'll keep the box alive.
		// Until this stops happening, AddCleanup will never be called.
		runtime.SetFinalizer(box, boxFinalizerKeepAlive)
	}
}
*/

// newBox creates a zero-value instance of Box.
func newBox(obj unsafe.Pointer) *Box {
	box := &Box{}
	box.gobject = obj

	// runtime.SetFinalizer(box, boxFinalizerKeepAlive)

	// runtime.AddCleanup(box, func(obj unsafe.Pointer) {
	// 	if toggleRefs {
	// 		slog.Debug(
	// 			"cleaning up object",
	// 			"box", TryGet(obj) != nil,
	// 			objInfo(obj))
	// 	}
	//
	// 	C.g_object_remove_toggle_ref((*C.GObject)(obj), (*[0]byte)(C.goToggleNotify), nil)
	//
	// 	if objectProfile != nil {
	// 		objectProfile.Remove(obj)
	// 	}
	// }, obj)

	runtime.SetFinalizer(box, func(box *Box) {
		box.finalizing.Store(true)

		obj := box.gobject

		var objInfoSaved slog.Attr
		if toggleRefs {
			objInfoSaved = objInfo(obj)
		}

		shared.mu.Lock()

		if toggleRefs {
			slog.Debug(
				"cleaning up object",
				"box", box != nil,
				objInfoSaved)
		}

		// weak.Pointer's behavior is to be invalidated by the time the
		// finalizer is called, so we temporarily resurrect the box so that
		// destroy signal handlers can obtain it. We'll purge it from the
		// registry after the destroy callbacks.
		shared.weak[obj] = weak.Make(box)

		shared.mu.Unlock()

		C.g_object_remove_toggle_ref((*C.GObject)(obj), (*[0]byte)(C.goToggleNotify), nil)

		if objectProfile != nil {
			objectProfile.Remove(obj)
		}

		if toggleRefs {
			shared.mu.RLock()
			defer shared.mu.RUnlock()

			_, weak := shared.weak[obj]
			_, strong := shared.strong[obj]

			slog.Debug(
				"post-finalizer aftermath for box",
				"is_weak", weak,
				"is_strong", strong,
				objInfoSaved)
		}
	})

	if objectProfile != nil {
		objectProfile.Add(obj, 3)
	}

	if traceObjects {
		slog.Debug(
			"allocating new box for object",
			"stack", string(debug.Stack()),
			objInfo(obj))
	}

	// Force box on the heap. Objects on the stack can move, but not objects on
	// the heap. At least not for now; the assume-no-moving-gc import will
	// guard against that.
	sink(box)

	return box
}

// shared contains shared closure data.
var shared = struct {
	mu sync.RWMutex
	// weak stores *Box while the object is in Go's heap. The finalizer will
	// move *Box to strong if the reference is toggled. This is only the case,
	// because the finalizer will not run otherwise.
	weak map[unsafe.Pointer]weak.Pointer[Box]
	// strong stores *Box while the object is still referenced by C but not Go.
	strong map[unsafe.Pointer]*Box
}{
	weak:   make(map[unsafe.Pointer]weak.Pointer[Box], 1024),
	strong: make(map[unsafe.Pointer]*Box, 1024),
}

// TryGet gets the Box associated with the GObject or nil if it's gone. The
// caller must not retain the Box pointer anywhere.
func TryGet(gobject unsafe.Pointer) *Box {
	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()
	return box
}

// Get gets the interned box for the given GObject C pointer. If the object is
// new or unknown, then a new box is made. If the intern box already exists for
// a given C pointer, then that box is weakly referenced and returned. The box
// will be reference-counted; the caller must use ShouldFree to unreference it.
func Get(gobject unsafe.Pointer, take bool) *Box {
	// If the registry does not exist, then we'll have to globally register it.
	// If the registry is currently strongly referenced, then we must move it to
	// a weak reference.

	box := TryGet(gobject)
	if box != nil {
		return box
	}

	shared.mu.Lock()

	box, _ = gets(gobject)
	if box != nil {
		shared.mu.Unlock()
		return box
	}

	box = newBox(gobject)

	// add_toggle_ref's documentation states:
	//
	//    Since a (normal) reference must be held to the object before
	//    calling g_object_add_toggle_ref(), the initial state of the
	//    reverse link is always strong.
	//
	shared.strong[gobject] = box

	if toggleRefs {
		slog.Debug(
			"Get: will introduce new box for object",
			objInfo(gobject))
	}

	shared.mu.Unlock()

	C.g_object_add_toggle_ref(
		(*C.GObject)(gobject),
		(*[0]byte)(C.goToggleNotify), nil,
	)

	if toggleRefs {
		slog.Debug(
			"Get: added toggle reference to object",
			objInfo(gobject))
	}

	// We should already have a strong reference. Sink the object in case. This
	// will force the reference to be truly strong.
	if C.g_object_is_floating(C.gpointer(gobject)) != C.FALSE {
		// First, we need to ref_sink the object to convert the floating
		// reference to a strong reference.
		C.g_object_ref_sink(C.gpointer(gobject))
		// Then, we need to unref it to balance the ref_sink.
		C.g_object_unref(C.gpointer(gobject))

		if toggleRefs {
			slog.Debug(
				"Get: ref_sink'd the object",
				objInfo(gobject))
		}
	}

	// If we're "not taking," then we can assume our ownership over the object,
	// meaning the strong reference is now ours. That means we need to replace
	// it, not add.
	if !take {
		C.g_object_unref(C.gpointer(gobject))
		if toggleRefs {
			slog.Debug(
				"Get: not taking, so unref'd the object",
				objInfo(gobject))
		}
	}

	// Undo the initial ref_sink.
	// C.g_object_unref(C.gpointer(gobject))

	return box
}

// finalizeBox only delays its finalization until GLib notifies us a toggle. It
// does so for as long as an object is stored only in the Go heap. Once the
// object is also shared, the toggle notifier will strongly reference the Box.

/*
func finalizeGObject(gobject unsafe.Pointer) {
	shared.mu.Lock()
	defer shared.mu.Unlock()

	box, strong := gets(gobject)
	if box == nil {
		// Silently ignore unknown objects.
		//
		// This is a trick to make sure that the box is really finalized. Turns
		// out it hates being finalized in goFinishRemovingToggleRef, so we just
		// don't call it there and let the GC do its thing.

		if traceObjects {
			slog.Debug(
				"finalizeBox: unknown object, possible bug?",
				"gobject.ptr", fmt.Sprintf("%p", dummy.gobject))
		}

		return
	}

	// Always delegate the finalization to the next cycle.
	// This won't be the case once goFinishRemovingToggleRef is called.
	runtime.SetFinalizer(dummy, finalizeBox)

	if box.finalize {
		// If the box is already finalizing, then we don't need to do anything.
		// Repeat this until box is gone from the registry.

		if toggleRefs {
			slog.Debug(
				"finalizeBox: already finalizing, waiting for goFinishRemovingToggleRef",
				"gobject.ptr", fmt.Sprintf("%p", dummy.gobject))
		}

		return
	}

	if strong {
		// If strong: the closures are strong-referenced, then they might still
		// be referenced from the C side, and those closures might access this
		// object. Don't free.

		if toggleRefs {
			slog.Debug(
				"finalizeBox: moving finalize to next GC cycle since object is still strong",
				objInfo(dummy.gobject))
		}

		return
	}

	// Mark the box as finalizing.
	box.finalize = true

	// Do this before we dispatch the remove_toggle_ref, because the
	// remove_toggle_ref might destroy the object.
	var prevObjInfo slog.Attr
	if toggleRefs {
		prevObjInfo = objInfo(dummy.gobject)
	}

	// Do this in the main loop instead. This is because finalizers are
	// called in a finalizer thread, and our remove_toggle_ref might be
	// destroying other main loop objects.
	C.g_main_context_invoke(
		nil, // nil means the default main context
		(*[0]byte)(C.gotk4_intern_remove_toggle_ref),
		C.gpointer(dummy.gobject))

	if toggleRefs {
		slog.Debug(
			"finalizeBox: remove_toggle_ref queued for next main loop iteration",
			prevObjInfo)
	}
}
*/

//go:nosplit
func gets(gobject unsafe.Pointer) (b *Box, strong bool) {
	if strong, ok := shared.strong[gobject]; ok {
		return strong, true
	}

	if weakPtr, ok := shared.weak[gobject]; ok {
		if weak := weakPtr.Value(); weak != nil {
			// If forObject is false, then that probably means this was called
			// inside goMarshal while the Go object is still alive, otherwise
			// toggleNotify would've moved it over. We don't have to worry about
			// this being freed as long as we acquire the mutex.
			//
			// TODO: does this actually resurrect the value properly? We have a
			// mutex to guard this which is also used in the finalizer, so it
			// shouldn't explode, but still.
			return weak, false
		}
	}

	return nil, false
}

// makeStrong forces the Box instance associated with the given object to be
// strongly referenced.
//
//go:nosplit
func makeStrong(gobject unsafe.Pointer) *Box {
	// TODO: double mutex check, similar to ShouldFree.

	box, strong := gets(gobject)
	if toggleRefs {
		slog.Debug(
			"makeStrong: obtained box",
			"strong", strong,
			"box", box != nil,
			objInfo(gobject))
	}
	if box == nil {
		return nil
	}

	if !strong {
		shared.strong[gobject] = box
		delete(shared.weak, gobject)
	}

	return box
}

// makeWeak forces the Box intsance associated with the given object to be
// weakly referenced.
//
//go:nosplit
func makeWeak(gobject unsafe.Pointer) *Box {
	box, strong := gets(gobject)
	if toggleRefs {
		slog.Debug(
			"makeWeak: obtained box",
			"strong", strong,
			"box", box != nil,
			objInfo(gobject))
	}
	if box == nil {
		return nil
	}

	if strong {
		shared.weak[gobject] = weak.Make(box)
		delete(shared.strong, gobject)
	}

	return box
}
