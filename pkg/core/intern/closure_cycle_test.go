package intern_test

import (
	"os"
	"sync"
	"testing"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/core/intern"
	"github.com/diamondburned/gotk4/pkg/core/intern/leaktest"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func TestMain(m *testing.M) {
	gtk.Init()
	os.Exit(m.Run())
}

func TestClosureSelfReferenceCycle(t *testing.T) {
	var tracked *leaktest.TrackedObject

	func() {
		btn := gtk.NewButton()
		tracked = leaktest.Track(unsafe.Pointer(glib.BaseObject(btn).Native()))

		btn.ConnectClicked(func() {
			btn.SetLabel("clicked")
		})
	}()

	leaktest.AssertFreed(t, tracked, "closure-self-reference")
}

func TestClosureCapturesParent(t *testing.T) {
	var trackedChild *leaktest.TrackedObject
	var trackedParent *leaktest.TrackedObject

	func() {
		parent := gtk.NewBox(gtk.OrientationVertical, 0)
		child := gtk.NewButton()
		trackedParent = leaktest.Track(unsafe.Pointer(glib.BaseObject(parent).Native()))
		trackedChild = leaktest.Track(unsafe.Pointer(glib.BaseObject(child).Native()))

		parent.Append(child)

		child.ConnectClicked(func() {
			parent.Remove(child)
			parent.Append(gtk.NewLabel("replaced"))
		})

		glib.Destroy(parent)
	}()

	leaktest.AssertFreed(t, trackedChild, "child-with-parent-capture")
	leaktest.AssertFreed(t, trackedParent, "parent-captured-by-child-handler")
}

func TestClosureCycleGrowth(t *testing.T) {
	const iterations = 100
	var objects []*leaktest.TrackedObject

	for i := 0; i < iterations; i++ {
		func() {
			btn := gtk.NewButton()
			objects = append(objects,
				leaktest.Track(unsafe.Pointer(glib.BaseObject(btn).Native())))
			btn.ConnectClicked(func() {
				btn.SetLabel("leak")
			})
		}()
	}

	leaktest.ForceGC()

	leaked := 0
	for _, obj := range objects {
		if !obj.IsFinalized() {
			leaked++
		}
	}

	if leaked > 0 {
		t.Errorf("LEAK: %d/%d objects not freed after GC", leaked, iterations)
	}
}

// TestConcurrentGetTryGet exercises the shared-map mutex discipline
// under the race detector. Hammer Get/TryGet from many goroutines while
// a finalizer may be running. Run with: go test -race.
//
// We use Take (intern.Get with take=true) to exercise the existing-entry
// transfer-none path. Get serializes access with shared.mu and reuses the
// existing toggle, so the ref count stays stable across the loop.
func TestConcurrentGetTryGet(t *testing.T) {
	btn := gtk.NewButton()
	defer glib.Destroy(glib.BaseObject(btn))
	ptr := unsafe.Pointer(glib.BaseObject(btn).Native())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				_ = intern.TryGet(ptr)
				_ = intern.Get(ptr, true)
			}
		}()
	}
	wg.Wait()
}

// BenchmarkGet measures the hot path: every wrapper creation calls
// intern.Get. A regression here affects every gotk4 application.
func BenchmarkGet(b *testing.B) {
	btn := gtk.NewButton()
	defer glib.Destroy(glib.BaseObject(btn))
	ptr := unsafe.Pointer(glib.BaseObject(btn).Native())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = intern.Get(ptr, true)
	}
}

// TestApplicationWindowLifecycle simulates a real application: create
// windows with self-referencing signal handlers, then explicitly
// destroy the windows. Verifies no leaks after 10 cycles.
//
// Explicit destruction must release closure-held references, claim cleanup
// exactly once, and release every reference owned by the wrapper so the
// underlying C object reaches finalization.
func TestApplicationWindowLifecycle(t *testing.T) {
	var windowTrackers []*leaktest.TrackedObject

	for cycle := 0; cycle < 10; cycle++ {
		func() {
			win := gtk.NewWindow()
			windowTrackers = append(windowTrackers,
				leaktest.Track(unsafe.Pointer(glib.BaseObject(win).Native())))

			// A self-referencing handler on the window creates the closure cycle
			// exercised by this lifecycle test.
			win.Connect("notify", func() {
				_ = win
			})

			// GTK destruction releases the window's widget-side references.
			// glib.Destroy separately removes the Go interning toggle ref (the
			// remaining reference owned by the wrapper), allowing finalization.
			win.Destroy()
			glib.Destroy(glib.BaseObject(win))
		}()
	}

	leaktest.ForceGC()

	leakedWindows := 0
	for _, tr := range windowTrackers {
		if !tr.IsFinalized() {
			leakedWindows++
		}
	}

	if leakedWindows > 0 {
		t.Errorf("LEAK: %d/%d windows not freed",
			leakedWindows, len(windowTrackers))
	}
}
