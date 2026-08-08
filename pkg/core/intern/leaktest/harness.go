package leaktest

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
#include <stdint.h>

extern void goLeakTestFinalized(uintptr_t handle);

static void on_weak_notify(gpointer data, GObject *where) {
	goLeakTestFinalized((uintptr_t)data);
}

static void install_weak_ref(gpointer obj, uintptr_t handle) {
	g_object_weak_ref(G_OBJECT(obj), on_weak_notify, (gpointer)handle);
}

static guint get_ref_count(gpointer obj) {
  return G_OBJECT(obj)->ref_count;
}

static void iterate_default_context(void) {
  while (g_main_context_iteration(NULL, FALSE)) {
  }
}

static gpointer hold_ref(gpointer obj) {
  return g_object_ref(obj);
}

static void release_ref(gpointer obj) {
  g_object_unref(obj);
}

static void notify_name(gpointer obj) {
	g_object_notify(G_OBJECT(obj), "name");
}

static void emit_clicked(gpointer obj) {
	g_signal_emit_by_name(G_OBJECT(obj), "clicked");
}

// new_gobject creates a fresh GObject of the base GObject type and
// returns it without taking ownership (the GC won't free C memory).
static gpointer new_gobject() {
	return (gpointer)g_object_new(G_TYPE_OBJECT, NULL);
}
*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type TrackedObject struct {
	ptr       unsafe.Pointer
	finalized atomic.Bool
	handle    cgo.Handle
}

//export goLeakTestFinalized
func goLeakTestFinalized(handle C.uintptr_t) {
	h := cgo.Handle(handle)
	if obj, ok := h.Value().(*TrackedObject); ok {
		obj.finalized.Store(true)
	}
	h.Delete()
}

func Track(ptr unsafe.Pointer) *TrackedObject {
	t := &TrackedObject{ptr: ptr}
	t.handle = cgo.NewHandle(t)
	C.install_weak_ref(C.gpointer(ptr), C.uintptr_t(t.handle))
	return t
}

// NewGObject creates a fresh GObject (the base GObject type) and returns
// its raw C pointer. Tests use it with glib.Take/glib.AssumeOwnership to
// build a *glib.Object. This lives in a cgo-enabled package so test
// files that are in packages already using cgo (where Go's test
// scaffolding rejects embedded cgo) can construct GObjects.
func NewGObject() unsafe.Pointer {
	return unsafe.Pointer(C.new_gobject())
}

// Hold adds a C-side reference for testing cleanup while a signal target
// remains externally alive.
func Hold(ptr unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.hold_ref(C.gpointer(ptr)))
}

// NotifyName emits GObject::notify::name without going through a Go wrapper.
func NotifyName(ptr unsafe.Pointer) {
	C.notify_name(C.gpointer(ptr))
}

func EmitClicked(ptr unsafe.Pointer) {
	C.emit_clicked(C.gpointer(ptr))
}

// Release drops a reference returned by Hold.
func Release(ptr unsafe.Pointer) {
	C.release_ref(C.gpointer(ptr))
}

func (t *TrackedObject) RefCount() uint {
	if t.finalized.Load() {
		return 0
	}
	return uint(C.get_ref_count(C.gpointer(t.ptr)))
}

func (t *TrackedObject) IsFinalized() bool {
	return t.finalized.Load()
}

func ForceGC() {
	for i := 0; i < 20; i++ {
		runtime.GC()
		time.Sleep(10 * time.Millisecond)
		C.iterate_default_context()
	}
}

func AssertFreed(t *testing.T, obj *TrackedObject, context string) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for !obj.IsFinalized() && time.Now().Before(deadline) {
		runtime.GC()
		time.Sleep(10 * time.Millisecond)
		C.iterate_default_context()
	}
	if !obj.IsFinalized() {
		t.Errorf("LEAK [%s]: object at %p was not freed (ref_count=%d)", context, obj.ptr, obj.RefCount())
	}
}

func AssertRefCount(t *testing.T, obj *TrackedObject, expected uint, context string) {
	t.Helper()
	if obj.IsFinalized() {
		if expected != 0 {
			t.Errorf("REFCOUNT [%s]: object already finalized, expected %d",
				context, expected)
		}
		return
	}
	actual := obj.RefCount()
	if actual != expected {
		t.Errorf("REFCOUNT [%s]: expected %d, got %d (ptr=%p)",
			context, expected, actual, obj.ptr)
	}
}
