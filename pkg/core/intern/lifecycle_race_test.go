package intern

import (
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/intern/leaktest"
)

func TestGetFreeWaitsForToggleInstallation(t *testing.T) {
	p := leaktest.NewGObject()
	entered := make(chan struct{})
	release := make(chan struct{})
	freeStarted := make(chan struct{})
	freeDone := make(chan struct{})
	beforeToggleInstallHook = func() {
		close(entered)
		candidate := shared.byObject[p].box
		go func() {
			close(freeStarted)
			Free(candidate)
			close(freeDone)
		}()
		<-release
	}
	defer func() { beforeToggleInstallHook = nil }()

	var first *Box
	done := make(chan struct{})
	go func() {
		first = Get(p, true)
		close(done)
	}()
	<-entered
	<-freeStarted

	select {
	case <-freeDone:
		t.Fatal("Free completed before toggle installation completed")
	case <-time.After(10 * time.Millisecond):
	}

	close(release)
	<-done
	<-freeDone
	if first == nil {
		t.Fatal("Get returned nil after toggle installation")
	}
	leaktest.Release(p)
	leaktest.ForceGC()
}

func TestRetiredToggleTokenIgnored(t *testing.T) {
	p := leaktest.NewGObject()
	box := Get(p, true)
	if box == nil {
		t.Fatal("Get returned nil")
	}
	token := box.token

	Free(box)
	if got := TryGet(p); got != nil {
		t.Fatal("retired entry remained addressable")
	}

	// A late callback carrying the retired token must not affect any entry,
	// even if the native address is still allocated.
	toggleNotify(token, p, false)
	if got := TryGet(p); got != nil {
		t.Fatal("late retired-token callback resurrected an entry")
	}

	leaktest.Release(p)
	leaktest.ForceGC()
}

func TestConcurrentExpiredWeakLookup(t *testing.T) {
	p := leaktest.NewGObject()
	box := Get(p, true)
	if box == nil {
		t.Fatal("initial Get returned nil")
	}
	// Drop the native owner, leaving only the toggle. This makes the entry
	// weak, after which the box can expire before the native object is held
	// again.
	box = nil
	leaktest.Release(p)
	for i := 0; i < 20 && TryGet(p) != nil; i++ {
		runtime.GC()
		time.Sleep(time.Millisecond)
	}
	if TryGet(p) != nil {
		t.Fatal("weak box did not expire")
	}

	held := leaktest.Hold(p)
	const workers = 8
	var wg sync.WaitGroup
	boxes := make(chan *Box, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			boxes <- Get(unsafe.Pointer(held), true)
		}()
	}
	wg.Wait()
	close(boxes)
	var recovered *Box
	for got := range boxes {
		if got == nil {
			t.Fatal("concurrent lookup returned nil")
		}
		if recovered == nil {
			recovered = got
		} else if recovered != got {
			t.Fatal("concurrent lookup created multiple boxes")
		}
	}

	Free(recovered)
	leaktest.Release(held)
	leaktest.ForceGC()
}

func TestSignalHandlerTrackingStates(t *testing.T) {
	var nilBox *Box
	nilBox.TrackSignalHandler(1)
	nilBox.UntrackSignalHandler(1)

	p := leaktest.NewGObject()
	box := Get(p, true)
	if box == nil {
		t.Fatal("Get returned nil")
	}

	// Handler ID 0 is GLib's invalid-handler sentinel and must be ignored.
	box.TrackSignalHandler(0)
	box.UntrackSignalHandler(0)

	shared.mu.Lock()
	e := shared.byToken[box.token]
	if e == nil || len(e.signalHandlers) != 0 {
		shared.mu.Unlock()
		t.Fatalf("expected empty signalHandlers for ID 0, got %v", e)
	}
	shared.mu.Unlock()

	box.TrackSignalHandler(42)
	box.TrackSignalHandler(99)

	shared.mu.Lock()
	if _, ok := e.signalHandlers[42]; !ok {
		shared.mu.Unlock()
		t.Fatal("handler 42 was not tracked")
	}
	if _, ok := e.signalHandlers[99]; !ok {
		shared.mu.Unlock()
		t.Fatal("handler 99 was not tracked")
	}
	shared.mu.Unlock()

	box.UntrackSignalHandler(42)
	shared.mu.Lock()
	if _, ok := e.signalHandlers[42]; ok {
		shared.mu.Unlock()
		t.Fatal("handler 42 was not untracked")
	}
	if _, ok := e.signalHandlers[99]; !ok {
		shared.mu.Unlock()
		t.Fatal("handler 99 should remain tracked")
	}
	shared.mu.Unlock()

	Free(box)

	shared.mu.Lock()
	if e.signalHandlers != nil {
		shared.mu.Unlock()
		t.Fatal("expected Free to clear signalHandlers")
	}
	shared.mu.Unlock()

	// Tracking against a retired Box must be ignored.
	box.TrackSignalHandler(123)
	shared.mu.Lock()
	if shared.byToken[box.token] != nil {
		shared.mu.Unlock()
		t.Fatal("retired token still present in shared.byToken")
	}
	shared.mu.Unlock()

	leaktest.Release(p)
	leaktest.ForceGC()
}
