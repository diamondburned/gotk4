// Package intern implements value interning for Cgo sharing.
package intern

// #cgo pkg-config: gobject-2.0
// #include "intern.h"
import "C"

import (
	"fmt"
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
)

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
	return slog.Group(
		"gobject",
		slog.String("ptr", fmt.Sprintf("%p", obj)),
		slog.String("type", C.GoString(C.gotk4_object_type_name(C.gpointer(obj)))),
		slog.Int("refs", objRefCount(obj)))
}

func objRefCount(obj unsafe.Pointer) int {
	return int(C.g_atomic_int_get((*C.gint)(unsafe.Pointer(&(*C.GObject)(obj).ref_count))))
}

// Box is the shared Go state for one generation of a native object.
type Box struct {
	dummy      *boxDummy
	closures   atomic.Pointer[closure.Registry]
	gobject    unsafe.Pointer
	token      uintptr
	generation uint64
	freed      atomic.Bool
}

// boxDummy is kept outside the Go closure cycle. The anchor makes the dummy
// pointer-scannable so it cannot be placed in the runtime's tiny allocator,
// whose coalesced allocation slots can suppress individual finalizers.
type boxDummy struct {
	anchor     *boxDummyAnchor
	token      uintptr
	generation uint64
}

type boxDummyAnchor struct{ _ byte }

var boxDummyScannedAnchor boxDummyAnchor

type entryPhase uint8

const (
	entryInitializing entryPhase = iota
	entryActive
	entryCleanupQueued
	entryRemoving
)

type entry struct {
	object         unsafe.Pointer
	token          uintptr
	phase          entryPhase
	box            *Box
	weak           weak.Pointer[Box]
	generation     uint64
	finalizerArmed bool
}

var shared = struct {
	mu       sync.Mutex
	byObject map[unsafe.Pointer]*entry
	byToken  map[uintptr]*entry
}{
	byObject: make(map[unsafe.Pointer]*entry, 1024),
	byToken:  make(map[uintptr]*entry, 1024),
}

var nextToken atomic.Uintptr

// beforeToggleInstallHook is used by lifecycle tests to pause initialization
// while the registry mutex is held. It is nil in normal builds.
var beforeToggleInstallHook func()

func newToken() uintptr { return nextToken.Add(1) }

func newBox(e *entry) *Box {
	e.generation++
	box := &Box{
		gobject:    e.object,
		token:      e.token,
		generation: e.generation,
	}
	box.dummy = &boxDummy{
		anchor:     &boxDummyScannedAnchor,
		token:      e.token,
		generation: e.generation,
	}
	if traceObjects {
		slog.Debug("allocating new box for object", "stack", string(debug.Stack()), objInfo(e.object))
	}
	return box
}

// Object returns Box's native GObject pointer.
func (b *Box) GObject() unsafe.Pointer { return b.gobject }

// Closures returns the closure registry for this Box.
func (b *Box) Closures() *closure.Registry {
	if reg := b.closures.Load(); reg != nil {
		return reg
	}
	reg := closure.NewRegistry()
	if b.closures.CompareAndSwap(nil, reg) {
		return reg
	}
	return b.closures.Load()
}

func consumeTransferredRef(gobject unsafe.Pointer, take bool) {
	if !take {
		C.g_object_unref(C.gpointer(gobject))
	}
}

func boxForEntry(e *entry) *Box {
	if e == nil {
		return nil
	}
	if e.box != nil {
		return e.box
	}
	return e.weak.Value()
}

// TryGet gets the currently interned box, or nil if no safe wrapper exists.
func TryGet(gobject unsafe.Pointer) *Box {
	if gobject == nil {
		return nil
	}
	shared.mu.Lock()
	e := shared.byObject[gobject]
	if e == nil || e.phase != entryActive {
		shared.mu.Unlock()
		return nil
	}
	box := boxForEntry(e)
	shared.mu.Unlock()
	return box
}

// Get interns a GObject. take=true means transfer-none; take=false consumes
// one caller-owned transfer-full reference on every outcome.
func Get(gobject unsafe.Pointer, take bool) *Box {
	if gobject == nil {
		return nil
	}

	shared.mu.Lock()
	if e := shared.byObject[gobject]; e != nil {
		if e.phase == entryActive || e.phase == entryCleanupQueued {
			// Get's transfer contract guarantees that the native object is
			// valid. Promote the stable entry and cancel queued cleanup rather
			// than installing a second toggle reference.
			box := makeStrong(e)
			shared.mu.Unlock()
			consumeTransferredRef(gobject, take)
			return box
		}
		shared.mu.Unlock()
		consumeTransferredRef(gobject, take)
		return nil
	}

	e := &entry{object: gobject, token: newToken(), phase: entryInitializing}
	if objectProfile != nil {
		// Profiles track native registry entries, not transient Box
		// generations. An expired weak Box may be replaced under this entry.
		objectProfile.Add(gobject, 3)
	}
	box := newBox(e)
	e.box = box
	shared.byObject[gobject] = e
	shared.byToken[e.token] = e

	// GLib requires a normal reference before adding a toggle reference and
	// guarantees the initial toggle state is strong. No toggle callback is
	// delivered during this installation, so publishing remains atomic.
	if beforeToggleInstallHook != nil {
		beforeToggleInstallHook()
	}
	C.gotk4_intern_add_toggle_ref(C.gpointer(gobject), C.guintptr(e.token))
	e.phase = entryActive
	shared.mu.Unlock()

	if C.g_object_is_floating(C.gpointer(gobject)) != C.FALSE {
		C.g_object_ref_sink(C.gpointer(gobject))
		C.g_object_unref(C.gpointer(gobject))
	}
	consumeTransferredRef(gobject, take)
	return box
}

// Free explicitly retires a Box and its exact toggle generation. It is
// synchronous; callers must use the object's owning GLib/GTK thread because
// removing the toggle reference can run native destruction on the caller.
func Free(box *Box) {
	if box == nil || !box.freed.CompareAndSwap(false, true) {
		return
	}

	shared.mu.Lock()
	e := shared.byToken[box.token]
	if e == nil || e.generation != box.generation || boxForEntry(e) != box ||
		e.phase != entryActive {
		shared.mu.Unlock()
		return
	}
	e.phase = entryRemoving
	if e.box == box {
		e.box = nil
	}
	e.weak = weak.Pointer[Box]{}
	runtime.SetFinalizer(box.dummy, nil)
	shared.mu.Unlock()

	disconnectClosures(box, box.gobject)
	C.gotk4_intern_remove_toggle_ref(C.gpointer(box.gobject), C.guintptr(box.token))
	finishCleanup(box.token)
}

func finalizeBox(dummy *boxDummy) {
	if dummy == nil {
		return
	}
	shared.mu.Lock()
	e := shared.byToken[dummy.token]
	if e == nil || e.generation != dummy.generation ||
		!shouldQueueFinalizerCleanup(e.phase, e.box != nil) {
		shared.mu.Unlock()
		return
	}
	e.phase = entryCleanupQueued
	token := e.token
	shared.mu.Unlock()

	C.gotk4_intern_queue_remove_toggle_ref(C.guintptr(token))
}

// shouldQueueFinalizerCleanup decides whether a boxDummy finalizer owns the
// active entry's cleanup. A finalizer proves its Box was unreachable at mark
// time; weak handles are cleared later during sweeping and must not veto it.
func shouldQueueFinalizerCleanup(phase entryPhase, hasStrongBox bool) bool {
	return phase == entryActive && !hasStrongBox
}

func beginQueuedCleanup(token uintptr) {
	shared.mu.Lock()
	e := shared.byToken[token]
	if e == nil || e.token != token || e.phase != entryCleanupQueued {
		shared.mu.Unlock()
		return
	}
	e.phase = entryRemoving
	obj := e.object
	shared.mu.Unlock()

	C.gotk4_intern_remove_toggle_ref(C.gpointer(obj), C.guintptr(token))
	finishCleanup(token)
}

func finishCleanup(token uintptr) {
	shared.mu.Lock()
	e := shared.byToken[token]
	if e == nil || e.phase != entryRemoving {
		shared.mu.Unlock()
		return
	}
	delete(shared.byToken, token)
	if shared.byObject[e.object] == e {
		delete(shared.byObject, e.object)
	}
	obj := e.object
	shared.mu.Unlock()
	if objectProfile != nil {
		objectProfile.Remove(obj)
	}
}

func disconnectClosures(box *Box, gobject unsafe.Pointer) {
	if box == nil {
		return
	}
	if reg := box.closures.Load(); reg != nil {
		reg.RangeSignals(func(gclosure unsafe.Pointer, _ *closure.FuncStack) bool {
			C.gotk4_intern_disconnect_closure(C.gpointer(gobject), (*C.GClosure)(gclosure))
			reg.Delete(gclosure)
			return true
		})
	}
}

func makeStrong(e *entry) *Box {
	if e == nil || e.phase == entryRemoving {
		return nil
	}
	if e.box == nil {
		box := e.weak.Value()
		if box != nil {
			if e.finalizerArmed {
				runtime.SetFinalizer(box.dummy, nil)
				e.finalizerArmed = false
			}
			e.box = box
		} else {
			// A queued finalizer for an expired box is harmless because it
			// carries the old generation. The replacement gets a new sentinel.
			e.finalizerArmed = false
			e.box = newBox(e)
		}
	}
	e.weak = weak.Pointer[Box]{}
	e.phase = entryActive
	return e.box
}

func makeWeak(e *entry) *Box {
	if e == nil || e.phase != entryActive || e.box == nil {
		return nil
	}
	box := e.box
	e.weak = weak.Make(box)
	e.box = nil
	if !e.finalizerArmed {
		runtime.SetFinalizer(box.dummy, finalizeBox)
		e.finalizerArmed = true
	}
	return box
}
