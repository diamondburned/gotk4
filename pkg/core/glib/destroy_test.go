package glib_test

import (
	"testing"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/core/intern/leaktest"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// TestDestroyBreaksCycle verifies that glib.Destroy manually breaks a
// closure cycle that GC alone cannot collect, by freeing the C object
// and asserting it is actually finalized.
//
// Use glib.AssumeOwnership here because leaktest.NewGObject returns an
// owned ref (a transfer-full-style C pointer): Get balances that ref with
// a single g_object_unref, leaving only the toggle ref owned by the box,
// so removing the toggle in Free drives the object to ref_count 0. (Take
// is the transfer-none codegen pattern; using it with an owned ref would
// leave the object alive forever here.)
func TestDestroyBreaksCycle(t *testing.T) {
	btn := gtk.NewButton()
	obj := glib.BaseObject(btn)
	tracked := leaktest.Track(unsafe.Pointer(obj.Native()))

	// Connect a self-referencing handler — this creates a cycle
	// (Box → closures → FuncStack → this func → obj → Object → box).
	btn.ConnectClicked(func() {
		_ = obj
	})

	// Explicitly destroy: removes the toggle ref, which is the only ref
	// the box owns, driving the object to ref_count 0.
	glib.Destroy(obj)

	leaktest.AssertFreed(t, tracked, "destroy-breaks-cycle")
}

// TestDestroyIdempotent verifies that calling Destroy twice doesn't crash
// and that the first Destroy actually freed the object.
func TestDestroyIdempotent(t *testing.T) {
	btn := gtk.NewButton()
	obj := glib.BaseObject(btn)
	tracked := leaktest.Track(unsafe.Pointer(obj.Native()))

	glib.Destroy(obj)
	leaktest.AssertFreed(t, tracked, "destroy-freed")

	// Repeated Destroy must be a no-op, not a crash or double-free.
	glib.Destroy(obj)
}

// TestDestroyDisconnectsSignalClosures verifies that explicit cleanup does
// not leave ordinary signal closures attached to a surviving C object.
func TestDestroyDisconnectsSignalClosures(t *testing.T) {
	btn := gtk.NewButton()
	obj := glib.BaseObject(btn)
	p := unsafe.Pointer(obj.Native())
	held := leaktest.Hold(p)
	tracked := leaktest.Track(p)

	called := 0
	btn.ConnectClicked(func() { called++ })
	glib.Destroy(obj)

	leaktest.EmitClicked(held)
	if called != 0 {
		t.Fatalf("destroyed signal closure was called %d times", called)
	}

	leaktest.Release(held)
	leaktest.AssertFreed(t, tracked, "destroy-disconnects-signal")
}
