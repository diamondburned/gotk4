package glib_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/core/intern/leaktest"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func TestMain(m *testing.M) {
	gtk.Init()
	os.Exit(m.Run())
}

// TestSignalLifecycleRefCount verifies that connecting and disconnecting a
// signal doesn't permanently inflate the reference count.
// The signal lifecycle must not leave an extra native reference behind.
func TestSignalLifecycleRefCount(t *testing.T) {
	p := leaktest.NewGObject()
	obj := glib.AssumeOwnership(p)
	defer glib.Destroy(obj)
	tracked := leaktest.Track(p)

	refBefore := tracked.RefCount()

	id := obj.Connect("notify", func() {})
	obj.HandlerDisconnect(id)

	refAfter := tracked.RefCount()
	if refAfter != refBefore {
		t.Errorf("LEAK: ref_count after disconnect is %d, expected %d "+
			"(signal lifecycle retained an extra reference)", refAfter, refBefore)
	}
}

// TestGeneratedClosureRefCount verifies ConnectGeneratedClosure.
// NotifyProperty emits "notify::<prop>"; use GtkButton's real "label"
// property so the detailed signal and its closure are both valid.
func TestGeneratedClosureRefCount(t *testing.T) {
	btn := gtk.NewButton()
	obj := glib.BaseObject(btn)
	defer glib.Destroy(obj)
	tracked := leaktest.Track(unsafe.Pointer(obj.Native()))

	refBefore := tracked.RefCount()

	id := obj.NotifyProperty("label", func() {})
	obj.HandlerDisconnect(id)

	refAfter := tracked.RefCount()
	if refAfter != refBefore {
		t.Errorf("LEAK: ref_count after NotifyProperty+Disconnect is %d, expected %d",
			refAfter, refBefore)
	}
}

// TestSignalCallbackCanDestroyObject exercises the GObject closure watch's
// marshal guards: removing the toggle reference during a callback must not
// destroy the native object until signal emission has unwound.
func TestSignalCallbackCanDestroyObject(t *testing.T) {
	btn := gtk.NewButton()
	obj := glib.BaseObject(btn)
	called := false

	btn.ConnectClicked(func() {
		called = true
		glib.Destroy(obj)
	})
	obj.Emit("clicked")

	if !called {
		t.Fatal("signal callback was not called")
	}
}

// TestSignalStillFires verifies that signal delivery remains functional.
// Use gtk.Button with obj.Emit("clicked")
// (a no-param signal). Emit("notify") would warn without a GParamSpec.
func TestSignalStillFires(t *testing.T) {
	btn := gtk.NewButton()
	defer glib.Destroy(glib.BaseObject(btn))

	var called int
	btn.ConnectClicked(func() {
		called++
	})

	glib.BaseObject(btn).Emit("clicked")

	if called != 1 {
		t.Errorf("signal handler not called: got %d, want 1", called)
	}
}
