// Package intern implements value interning for Cgo sharing.
package intern

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/closure"

	// Require a non-moving GC for heap pointers. Current GC is moving only by
	// the stack. See https://github.com/go4org/intern.
	_ "go4.org/unsafe/assume-no-moving-gc"
)

// Box contains possible interned values for each GObject.
type Box struct {
	Closures closure.Registry

	obj uintptr
}

// Hack to force an object on the heap.
var never bool
var sink interface{}

// newBox creates a zero-value instance of Box.
func newBox(obj unsafe.Pointer) *Box {
	box := &Box{obj: uintptr(obj)}

	// Force box on the heap. Objects on the stack can move, but not objects on
	// the heap. At least not for now; the assume-no-moving-gc import will
	// guard against that.
	if never {
		sink = box
	}

	runtime.SetFinalizer(box, finalizeBox)

	return box
}

func finalizeBox(box *Box) {
	shared.mu.Lock()
	defer shared.mu.Unlock()

	if box.obj != 0 {
		runtime.SetFinalizer(box, finalizeBox)
		return
	}
}

// shared contains shared closure data.
var shared = struct {
	mu sync.RWMutex
	// weak stores *Box while the object is in Go's heap. The finalizer will
	// move *Box to strong if the reference is toggled. This is only the case,
	// because the finalizer will not run otherwise.
	weak map[unsafe.Pointer]uintptr
	// strong stores *Box while the object is still referenced by C but not Go.
	strong map[unsafe.Pointer]*Box
	// sharing keeps toggle notifies.
	sharing map[unsafe.Pointer]struct{}
}{
	weak:    make(map[unsafe.Pointer]uintptr),
	strong:  make(map[unsafe.Pointer]*Box),
	sharing: make(map[unsafe.Pointer]struct{}),
}

// ObjectClosure gets the FuncStack instance from the given GObject and GClosure
// pointers. The given unsafe.Pointers MUST be C pointers.
func ObjectClosure(gobject, gclosure unsafe.Pointer) *closure.FuncStack {
	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()

	if box == nil || box.obj == 0 {
		// log.Println("gobject", gobject, "requesting destroyed box")
		return nil
	}

	return box.Closures.Load(gclosure)
}

// RemoveClosure removes the given GClosure callback.
func RemoveClosure(gobject, gclosure unsafe.Pointer) {
	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()

	if box != nil {
		box.Closures.Delete(gclosure)
		// log.Println("deleting object", unsafe.Pointer(gobject), "closure", unsafe.Pointer(gclosure))
	}

	// The closure missing here isn't very important. It can happen when the
	// function is called not because the user explicitly wanted to detach a
	// signal handler, but because the object is destroyed. In that case, Go
	// will have already handled it.
}

// ObjectBox gets the interned box for the given GObject C pointer. If the
// object is new or unknown, then a new box is made.
func ObjectBox(gobject unsafe.Pointer) *Box {
	if box := weakCheck(gobject); box != nil {
		return box
	}

	// If the registry does not exist, then we'll have to globally register it.
	// If the registry is currently strongly referenced, then we must move it to
	// a weak reference.

	shared.mu.Lock()
	defer shared.mu.Unlock()

	box, strong := gets(gobject)
	if box == nil {
		box = newBox(gobject)
	} else if strong {
		// Ensure that this box is weakly referenced.
		delete(shared.strong, gobject)
	}

	shared.weak[gobject] = uintptr(unsafe.Pointer(box))
	return box
}

// weakCheck is a fast path if the given object is already a known object that
// has a gobject pointer and a weak reference.
//
// TODO: this stage can be lazily delegated to each objectNative instance having
// its own sync.Once.
func weakCheck(gobject unsafe.Pointer) *Box {
	shared.mu.RLock()
	defer shared.mu.RUnlock()

	// Fast path if if this is a known object.
	box, strong := gets(gobject)
	if box != nil && !strong {
		return box
	}

	return nil
}

//go:nocheckptr
func gets(gobject unsafe.Pointer) (b *Box, strong bool) {
	if strong, ok := shared.strong[gobject]; ok {
		return strong, true
	}

	if weak, ok := shared.weak[gobject]; ok {
		// If forObject is false, then that probably means this was called
		// inside goMarshal while the Go object is still alive, otherwise
		// toggleNotify would've moved it over. We don't have to worry about
		// this being freed as long as we acquire the mutex.
		return (*Box)(unsafe.Pointer(weak)), false
	}

	return nil, false
}

// Toggle is called on the GToggleNotify callback.
func Toggle(gobject unsafe.Pointer, isLast bool) {
	shared.mu.Lock()
	defer shared.mu.Unlock()

	if isLast {
		delete(shared.sharing, gobject)
		makeWeak(gobject)
	} else {
		shared.sharing[gobject] = struct{}{}
		makeStrong(gobject)
	}
}

// makeStrong forces the Box instance associated with the given object to be
// strongly referenced.
func makeStrong(gobject unsafe.Pointer) {
	// TODO: double mutex check, similar to ShouldFree.

	box, strong := gets(gobject)
	if box == nil {
		return
	}

	if !strong {
		shared.strong[gobject] = box
		delete(shared.weak, gobject)
	}
}

// makeWeak forces the Box intsance associated with the given object to be
// weakly referenced.
func makeWeak(gobject unsafe.Pointer) {
	box, strong := gets(gobject)
	if box == nil {
		return
	}

	if strong {
		shared.weak[gobject] = uintptr(unsafe.Pointer(box))
		delete(shared.strong, gobject)
	}
}

// ShouldFree must only be called during finalizing of an object. It's used to
// know if an object should be freed or not during finalization. If false is
// returned, then the object must not be freed yet.
//
//go:nocheckptr
func ShouldFree(gobject unsafe.Pointer) bool {
	// shared.mu.RLock()
	// result, weak := preemptiveShouldFree(gobject)
	// shared.mu.RUnlock()

	// if !weak {
	// 	return result
	// }

	// return true

	shared.mu.Lock()
	defer shared.mu.Unlock()

	if _, ok := shared.sharing[gobject]; ok {
		return false
	}

	// Recheck to ensure that the state stayed the same while we couldn't
	// acquire the lock.
	result, weak := preemptiveShouldFree(gobject)
	if !weak {
		return result
	}

	box := (*Box)(unsafe.Pointer(shared.weak[gobject]))
	if box == nil {
		// The weak flag is incorrect, for some reason. Allow freeing.
		return true
	}

	// log.Printf("deleting object %p closure, result=%v, weak=%v", gobject, result, weak)

	// If the closures are weak-referenced, then the object reference hasn't
	// been toggled yet. Since the object is going away and we're still weakly
	// referenced, we can wipe the closures away.
	delete(shared.weak, gobject)

	// By setting *box to a zero-value of closures, we're nilling out the maps,
	// which will signal to Go that these cyclical objects can be freed
	// altogether.
	*box = Box{}

	// We can proceed to free the object.
	return true
}

// preemptiveShouldFree is a fast path that can be executed using just a
// read-only lock. The only edge case that this function cannot fully handle is
// if the GObject is found in the weak reference map, in which weak=true is
// returned.
func preemptiveShouldFree(gobject unsafe.Pointer) (res, weak bool) {
	if len(shared.strong) == 0 && len(shared.weak) == 0 {
		// We have no boxes, so we can free.
		return true, false
	}

	_, ok := shared.strong[gobject]
	if ok {
		// If the closures are strong-referenced, then they might still be
		// referenced from the C side, and those closures might access this
		// object. Don't free.
		return false, false
	}

	_, ok = shared.weak[gobject]
	if ok {
		return false, true
	}

	// We're not seeing any closures belonging to this object, so it can be
	// freed.
	return true, false
}
