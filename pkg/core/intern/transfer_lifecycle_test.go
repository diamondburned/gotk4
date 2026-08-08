package intern_test

import (
	"runtime"
	"testing"
	"time"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/core/intern"
	"github.com/diamondburned/gotk4/pkg/core/intern/leaktest"
)

type finalizerBlocker [1024]byte

func blockFinalizers(t *testing.T) func() {
	t.Helper()
	started := make(chan struct{})
	release := make(chan struct{})
	v := new(finalizerBlocker)
	runtime.SetFinalizer(v, func(*finalizerBlocker) {
		close(started)
		<-release
	})
	v = nil
	for i := 0; i < 20; i++ {
		runtime.GC()
		select {
		case <-started:
			return func() { close(release) }
		default:
			time.Sleep(time.Millisecond)
		}
	}
	t.Fatal("timed out waiting for finalizer blocker")
	return func() { close(release) }
}

// TestTakeTransferNoneNonFloating reproduces the critical regression: a
// transfer-none, non-floating object held by a C-side owner that is wrapped
// with glib.Take (the codegen pattern, e.g. GtkCellLayout.Area()). With the
// buggy baseRef logic, the box kept an owned base ref that finishCleanup
// tried to drop AFTER g_object_remove_toggle_ref had already freed the
// object — a use-after-free that either crashed or corrupted memory.
//
// Reproduction: build a non-floating GObject (the base GObject type never
// floats), treat leaktest.NewGObject's return as "the C-side owner's ref"
// (one ref on the C side — that's the transfer-none caller's situation),
// wrap it with glib.Take (the literal transfer-none code pattern), then drop
// first the Go wrapper and secondly the C-side owner's ref. The toggle ref
// added in Take must drop the object to ref_count 0 on its own;
// finishCleanup must NOT issue an extra g_object_unref — that extra unref
// is the use-after-free we are guarding against. Verifying "no crash" is the
// regression assertion; leaktest.Track also proves the object was actually
// freed rather than leaked.
func TestTakeTransferNoneNonFloating(t *testing.T) {
	p := leaktest.NewGObject()
	tracked := leaktest.Track(p)

	// N = 1 (the C-side owner's ref, modeled by the g_object_new return).
	if tracked.RefCount() != 1 {
		t.Fatalf("setup ref_count = %d, want 1", tracked.RefCount())
	}

	// The transfer-none codegen pattern: glib.Take. Take must NOT add an
	// owned base ref — only a toggle ref. After this, N = 2.
	obj := glib.Take(p)
	if obj == nil {
		t.Fatal("Take returned nil")
	}
	if got := tracked.RefCount(); got != 2 {
		t.Fatalf("post-Take ref_count = %d, want 2 (toggle + 1 C-owned)", got)
	}

	// Drop the Go wrapper reference. The box stays strong (the C-side
	// owner's ref keeps ref_count above 1), so no teardown should run yet.
	obj = nil
	leaktest.ForceGC()
	if tracked.IsFinalized() {
		t.Fatal("object finalized while C-side owner still holds it")
	}

	// Drop the C-side owner's reference: N 2 -> 1 -> is_last -> weak
	// transition -> GC -> idle removes toggle -> N 1 -> 0 -> freed.
	// With the buggy baseRef code, finishCleanup would g_object_unref the
	// already-freed object here and crash.
	leaktest.Release(p)

	leaktest.AssertFreed(t, tracked, "take-transfer-none-nonfloating")
}

// TestTransferNoneOwnerDropTriggersGC focuses on the GC finalizer path when
// the C-side owner holds a long-lived ref, mirroring real wrapped widgets:
// C owns the object, we Take() a wrapper, eventually C drops its last ref,
// the toggle fires is_last, and the GC frees the wrapper without crashing.
func TestTransferNoneOwnerDropTriggersGC(t *testing.T) {
	p := leaktest.NewGObject()
	tracked := leaktest.Track(p)

	obj := glib.Take(p)
	if obj == nil {
		t.Fatal("Take returned nil")
	}

	// Wrapper reachable, C side holds the only other ref. ref_count = 2.
	if got := tracked.RefCount(); got != 2 {
		t.Fatalf("post-Take ref_count = %d, want 2", got)
	}

	// Drop the explicit Go wrapper reference but keep the C-side ref.
	// Without C dropping its ref, is_last has not fired, so the box must
	// stay strong and the object must remain alive.
	obj = nil
	leaktest.ForceGC()
	if got := tracked.RefCount(); got != 2 {
		t.Fatalf("after wrapper drop ref_count = %d, want 2 (C still owns)", got)
	}
	if tracked.IsFinalized() {
		t.Fatal("object finalized while C-side owner still holds it")
	}

	// Reacquire a wrapper via Take to ensure re-interning of a still-alive,
	// strong-box object works after a transient drop.
	obj2 := glib.Take(p)
	if obj2 == nil {
		t.Fatal("re-Take returned nil")
	}
	if got := tracked.RefCount(); got != 2 {
		t.Fatalf("post re-Take ref_count = %d, want 2 (no extra ref added)", got)
	}

	// Drop the C-side holder's ref: N 2 -> 1 -> is_last -> weak -> GC
	// -> idle removes toggle -> ref_count 0 -> freed.
	leaktest.Release(p)

	// Release the Go wrapper reference and drain idle. The object must be
	// freed without any use-after-free.
	obj2 = nil
	leaktest.AssertFreed(t, tracked, "transfer-none-owner-drop-triggers-gc")
}

func TestAssumeOwnershipExistingConsumesTransferredRef(t *testing.T) {
	p := leaktest.NewGObject()
	tracked := leaktest.Track(p)

	obj := glib.AssumeOwnership(p)
	if obj == nil {
		t.Fatal("first AssumeOwnership returned nil")
	}
	if got := tracked.RefCount(); got != 1 {
		t.Fatalf("first AssumeOwnership ref_count = %d, want 1", got)
	}

	transferred := leaktest.Hold(p)
	if got := tracked.RefCount(); got != 2 {
		t.Fatalf("transferred ref_count = %d, want 2", got)
	}

	obj2 := glib.AssumeOwnership(transferred)
	if obj2 == nil {
		t.Fatal("second AssumeOwnership returned nil")
	}
	if got := tracked.RefCount(); got != 1 {
		t.Fatalf("second AssumeOwnership ref_count = %d, want 1", got)
	}

	glib.Destroy(obj)
	leaktest.AssertFreed(t, tracked, "assume-ownership-existing")
	_ = obj2
}

func TestExpiredWeakReusesExistingToggle(t *testing.T) {
	releaseFinalizers := blockFinalizers(t)
	defer func() { releaseFinalizers() }()

	p := leaktest.NewGObject()
	tracked := leaktest.Track(p)
	func() {
		obj := glib.Take(p)
		leaktest.Release(p)
		runtime.KeepAlive(obj)
	}()
	for i := 0; i < 20 && intern.TryGet(p) != nil; i++ {
		runtime.GC()
	}
	if intern.TryGet(p) != nil {
		t.Fatal("box weak pointer did not expire")
	}

	held := leaktest.Hold(p)
	if got := tracked.RefCount(); got != 2 {
		t.Fatalf("after native reacquire ref_count = %d, want 2", got)
	}

	obj := glib.Take(held)
	if obj == nil {
		t.Fatal("Take after native reacquire returned nil")
	}
	if got := tracked.RefCount(); got != 2 {
		t.Fatalf("re-Take installed another toggle: ref_count = %d, want 2", got)
	}
	runtime.KeepAlive(obj)
	obj = nil
	leaktest.Release(held)
	releaseFinalizers()
	releaseFinalizers = func() {}
	leaktest.AssertFreed(t, tracked, "expired-weak-reuses-toggle")
}
