// Package intern implements value interning for Cgo sharing.
package intern

// #cgo pkg-config: gobject-2.0
// #include <glib-object.h>
//
// extern void goToggleNotify(gpointer, GObject*, gboolean);
// static const gchar* gotk4_object_type_name(gpointer obj) { return G_OBJECT_TYPE_NAME(obj); };
import "C"

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	// Require a non-moving GC for heap pointers. Current GC is moving only by
	// the stack. See https://github.com/go4org/intern.
	_ "go4.org/unsafe/assume-no-moving-gc"
)

const maxTypesAllowed = 3

var knownTypes uint32

type BoxedType[T any] struct {
	ctor func(*Box) *T
	id   uint32
}

func RegisterType[T any](ctor func(*Box) *T) BoxedType[T] {
	t := BoxedType[T]{
		ctor: ctor,
		id:   atomic.AddUint32(&knownTypes, 1),
	}
	if t.id > maxTypesAllowed {
		panic("BoxedType ID overflow")
	}
	return t
}

func (t *BoxedType[T]) Get(box *Box) *T {
	if t.ctor == nil {
		ptr := atomic.LoadPointer(&box.data[t.id])
		return (*T)(ptr)
	}

	old := atomic.LoadPointer(&box.data[t.id])
	if old != nil {
		return (*T)(old)
	}

	new := t.ctor(box)

	if atomic.CompareAndSwapPointer(&box.data[t.id], nil, unsafe.Pointer(new)) {
		return new
	}

	ptr := atomic.LoadPointer(&box.data[t.id])
	if ptr != nil {
		return (*T)(ptr)
	}

	panic("Load returned nil after CompareAndSwap(old = nil) failed")
}

func (t *BoxedType[T]) Set(box *Box, v *T) {
	atomic.StorePointer(&box.data[t.id], unsafe.Pointer(v))
}

func (t *BoxedType[T]) Delete(box *Box) {
	atomic.StorePointer(&box.data[t.id], nil)
}

// Box is an opaque type holding extra data.
type Box struct {
	data [maxTypesAllowed]unsafe.Pointer
}

// Object returns Box's C GObject pointer.
func (b *Box) GObject() unsafe.Pointer {
	return atomic.LoadPointer(&b.data[0])
}

// Hack to force an object on the heap.
var never bool
var sink interface{}

var (
	traceObjects  *log.Logger
	toggleRefs    *log.Logger
	objectProfile *pprof.Profile
)

func init() {
	debug := os.Getenv("GOTK4_DEBUG")
	if debug == "" {
		return
	}

	for _, flag := range strings.Split(debug, ",") {
		switch flag {
		case "trace-objects":
			traceObjects = mustDebugLogger("trace-objects")
		case "toggle-refs":
			toggleRefs = mustDebugLogger("toggle-refs")
		case "profile-objects":
			objectProfile = pprof.NewProfile("gotk4-object-box")
		default:
			log.Panicf("unknown GOTK4_DEBUG flag %q", flag)
		}
	}
}

func mustDebugLogger(name string) *log.Logger {
	f, err := os.CreateTemp(os.TempDir(), fmt.Sprintf("gotk4-%s-%d-*", name, os.Getpid()))
	if err != nil {
		log.Panicln("cannot create temp", name, "file:", err)
	}

	log.Println("gotk4: intern: enabled debug file at", f.Name())
	return log.New(f, "", log.LstdFlags)
}

func objInfo(obj unsafe.Pointer) string {
	return fmt.Sprintf("%p (%s):", obj, C.GoString(C.gotk4_object_type_name(C.gpointer(obj))))
}

// newBox creates a zero-value instance of Box.
func newBox(obj unsafe.Pointer) *Box {
	box := &Box{}
	box.data[0] = obj

	if objectProfile != nil {
		objectProfile.Add(obj, 3)
	}

	if traceObjects != nil {
		traceObjects.Printf("%p: %s", obj, debug.Stack())
	}

	// Force box on the heap. Objects on the stack can move, but not objects on
	// the heap. At least not for now; the assume-no-moving-gc import will
	// guard against that.
	if never {
		sink = box
	}

	return box
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
}{
	weak:   make(map[unsafe.Pointer]uintptr, 1024),
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

	shared.mu.RLock()
	box, _ := gets(gobject)
	shared.mu.RUnlock()

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
	runtime.SetFinalizer(box, finalizeBox)

	shared.mu.Unlock()

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(gobject),
			"Get: will introduce new box, current ref =",
			C.g_atomic_int_get((*C.gint)(unsafe.Pointer(&(*C.GObject)(gobject).ref_count))))
	}

	C.g_object_add_toggle_ref(
		(*C.GObject)(gobject),
		(*[0]byte)(C.goToggleNotify), nil,
	)

	// If we're "not taking," then we can assume our ownership over the object,
	// meaning the strong reference is now ours. That means we need to replace
	// it, not add.
	if !take {
		C.g_object_unref(C.gpointer(gobject))
	}

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(gobject),
			"Get: introduced new box, current ref =",
			C.g_atomic_int_get((*C.gint)(unsafe.Pointer(&(*C.GObject)(gobject).ref_count))))
	}

	// Undo the initial ref_sink.
	// C.g_object_unref(C.gpointer(gobject))

	return box
}

// Free explicitly frees the box permanently. It must not be resurrected after
// this.
func Free(box *Box) {
	obj := box.GObject()
	if obj == nil {
		panic("bug: Free called on already freed object")
	}

	shared.mu.Lock()
	delete(shared.strong, obj)
	delete(shared.weak, obj)
	for i := range box.data {
		atomic.StorePointer(&box.data[i], nil)
	}
	shared.mu.Unlock()

	C.g_object_remove_toggle_ref(
		(*C.GObject)(unsafe.Pointer(obj)),
		(*[0]byte)(C.goToggleNotify), nil,
	)

	if toggleRefs != nil {
		toggleRefs.Println(objInfo(obj), "Free: explicitly removed toggle ref")
	}

	if objectProfile != nil {
		objectProfile.Remove(obj)
	}
}

// finalizeBox only delays its finalization until GLib notifies us a toggle. It
// does so for as long as an object is stored only in the Go heap. Once the
// object is also shared, the toggle notifier will strongly reference the Box.
func finalizeBox(box *Box) {
	obj := box.GObject()
	if obj == nil {
		return
	}

	var objInfoRes string
	if toggleRefs != nil {
		objInfoRes = objInfo(obj)
		toggleRefs.Println(objInfoRes, "finalizeBox: acquiring lock...")
	}

	shared.mu.Lock()

	if !freeBox(box) {
		// Delegate finalizing to the next cycle.
		runtime.SetFinalizer(box, finalizeBox)
		shared.mu.Unlock()
		if toggleRefs != nil {
			toggleRefs.Println(objInfoRes, "finalizeBox: moving finalize to next GC cycle")
		}
	} else {
		shared.mu.Unlock()
		// Unreference the object. This will potentially free the object as
		// well. The closures are definitely gone at this point.
		C.g_object_remove_toggle_ref(
			(*C.GObject)(unsafe.Pointer(obj)),
			(*[0]byte)(C.goToggleNotify), nil,
		)
		if toggleRefs != nil {
			toggleRefs.Println(objInfoRes, "finalizeBox: removed toggle ref during GC")
		}

		if objectProfile != nil {
			objectProfile.Remove(obj)
		}
	}
}

// freeBox must only be called during finalizing of a box. It's used to know if
// a box should be freed or not during finalization. If false is returned, then
// the object must not be freed yet.
//
//go:nocheckptr
func freeBox(box *Box) bool {
	b, ok := shared.strong[box.GObject()]
	if ok {
		if b != box {
			panic("bug: multiple Box found for same GObject")
		}
		// If the closures are strong-referenced, then they might still be
		// referenced from the C side, and those closures might access this
		// object. Don't free.
		return false
	}

	_, ok = shared.weak[box.GObject()]
	if ok {
		// If the closures are weak-referenced, then the object reference hasn't
		// been toggled yet. Since the object is going away and we're still
		// weakly referenced, we can wipe the closures away.
		delete(shared.weak, box.GObject())

		// By setting *box to a zero-value of closures, we're nilling out the
		// maps, which will signal to Go that these cyclical objects can be
		// freed altogether.
		for i := range box.data {
			atomic.StorePointer(&box.data[i], nil)
		}
	}

	// We can proceed to free the object.
	return true
}

// goToggleNotify is called by GLib on each toggle notification. It doesn't
// actually free anything and relies on Box's finalizer to free both the box and
// the C GObject.
//
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
