package glib_test

import (
	"testing"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/core/intern/leaktest"
)

func TestEmptyValueObjectIsNil(t *testing.T) {
	v := glib.InitValue(glib.TypeObject)
	if got := v.Object(); got != nil {
		t.Fatalf("empty object value returned %v, want nil", got)
	}
	if got := v.GoValue(); got != nil {
		t.Fatalf("empty object GoValue returned %v, want nil", got)
	}
}

// TestMarshalObjectTransferNone verifies marshalObject (called via
// Value.GoValue for TypeObject) doesn't leak a reference.
// Repeated conversion must preserve the original native reference count.
func TestMarshalObjectTransferNone(t *testing.T) {
	p := leaktest.NewGObject()
	obj := glib.AssumeOwnership(p)
	tracked := leaktest.Track(p)

	v := glib.InitValue(glib.TypeObject)
	v.SetObject(obj)

	refBefore := tracked.RefCount()

	for i := 0; i < 10; i++ {
		got := v.GoValue()
		if got == nil {
			t.Fatal("GoValue returned nil")
		}
	}

	refAfter := tracked.RefCount()
	if refAfter != refBefore {
		t.Errorf("LEAK: ref_count grew from %d to %d after 10 GoValue() calls",
			refBefore, refAfter)
	}
	v = nil
	glib.Destroy(obj)
	leaktest.AssertFreed(t, tracked, "marshal-object-transfer-none")
}

// TestValueObjectTransferNone verifies Value.Object() specifically.
func TestValueObjectTransferNone(t *testing.T) {
	p := leaktest.NewGObject()
	obj := glib.AssumeOwnership(p)
	tracked := leaktest.Track(p)

	v := glib.InitValue(glib.TypeObject)
	v.SetObject(obj)

	refBefore := tracked.RefCount()

	for i := 0; i < 10; i++ {
		o := v.Object()
		_ = o
	}

	refAfter := tracked.RefCount()
	if refAfter != refBefore {
		t.Errorf("LEAK: ref_count grew from %d to %d after 10 Object() calls",
			refBefore, refAfter)
	}
	v = nil
	glib.Destroy(obj)
	leaktest.AssertFreed(t, tracked, "value-object-transfer-none")
}

// TestCastDoesNotLeak verifies Cast() (g_value_init_from_instance +
// GoValue internally) doesn't leak.
func TestCastDoesNotLeak(t *testing.T) {
	p := leaktest.NewGObject()
	obj := glib.AssumeOwnership(p)
	tracked := leaktest.Track(p)

	refBefore := tracked.RefCount()

	for i := 0; i < 20; i++ {
		_ = obj.Cast()
	}

	refAfter := tracked.RefCount()
	if refAfter != refBefore {
		t.Errorf("LEAK: ref_count grew from %d to %d after 20 Cast() calls",
			refBefore, refAfter)
	}
	glib.Destroy(obj)
	leaktest.AssertFreed(t, tracked, "cast-does-not-leak")
}
