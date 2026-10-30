package glib_test

import (
	"testing"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

type launchContextWithParentOverrides struct {
	gio.AppLaunchContext
}

type launchContextWithObjectOverrides struct {
	gio.AppLaunchContext
	initCalled bool
}

// TestSubclassOverrides verifies that subclasses can add overrides for any
// class in its parent chain, and that adding overrides for a class other than
// GObject doesn't break GObject's own overrides.
func TestSubclassOverrides(t *testing.T) {
	t.Run("parent overrides only", func(t *testing.T) {
		subclass := glib.RegisterSubclass[*launchContextWithParentOverrides](
			glib.WithOverrides(func(*launchContextWithParentOverrides) gio.AppLaunchContextOverrides {
				return gio.AppLaunchContextOverrides{}
			}),
		)

		// This used to panic, because GObject's init and dispose asked for
		// ObjectOverrides from the gio.AppLaunchContextOverrides.
		obj := subclass.New()
		glib.Destroy(obj)
	})

	t.Run("parent and object overrides", func(t *testing.T) {
		subclass := glib.RegisterSubclass[*launchContextWithObjectOverrides](
			glib.WithOverrides(func(*launchContextWithObjectOverrides) gio.AppLaunchContextOverrides {
				return gio.AppLaunchContextOverrides{}
			}),
			glib.WithOverrides(func(obj *launchContextWithObjectOverrides) glib.ObjectOverrides {
				return glib.ObjectOverrides{
					Init:    func() { obj.initCalled = true },
					Dispose: func() {},
				}
			}),
		)

		obj := subclass.New()
		if !obj.initCalled {
			t.Error("ObjectOverrides.Init was not called")
		}

		overrides := glib.Overrides[glib.ObjectOverrides](obj)
		if overrides.Init == nil || overrides.Dispose == nil {
			t.Error("Overrides did not return the ObjectOverrides")
		}

		glib.Destroy(obj)
	})
}
