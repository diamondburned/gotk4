package test

import (
	"runtime"
	"testing"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

func TestObjectData(t *testing.T) {
	testObjectData(t)

	runtime.GC()
}

func testObjectData(t *testing.T) {
	app := gio.NewApplication("foo.bar", gio.ApplicationFlagsNone)

	glib.Bind(app, "foo")

	if value := glib.Bounded[string](app); value == nil || *value != "foo" {
		t.Fatal("returned data did not match expected data")
	}

	glib.Unbind[string](app)
}
