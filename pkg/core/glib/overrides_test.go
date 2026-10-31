package glib_test

import (
	"testing"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

type launchContextWithParentOverrides struct {
	gio.AppLaunchContext
}

type launchContextWithObjectOverrides struct {
	gio.AppLaunchContext
	initCalled        bool
	constructedCalled bool
	disposeCalled     bool
}

type launchContextWithClassInit struct {
	gio.AppLaunchContext
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
					Init:        func() { obj.initCalled = true },
					Constructed: func() { obj.constructedCalled = true },
					Dispose:     func() { obj.disposeCalled = true },
				}
			}),
		)

		obj := subclass.New()
		if !obj.initCalled {
			t.Error("ObjectOverrides.Init was not called")
		}
		if !obj.constructedCalled {
			t.Error("ObjectOverrides.Constructed was not called")
		}

		overrides := glib.Overrides[glib.ObjectOverrides](obj)
		if overrides.Init == nil || overrides.Constructed == nil || overrides.Dispose == nil {
			t.Error("Overrides did not return the ObjectOverrides")
		}

		glib.Destroy(obj)
		if !obj.disposeCalled {
			t.Error("ObjectOverrides.Dispose was not called")
		}
	})

	t.Run("with class init", func(t *testing.T) {
		var classInitCalled bool
		subclass := glib.RegisterSubclass[*launchContextWithClassInit](
			glib.WithClassInit(func(c *gio.AppLaunchContextClass) {
				classInitCalled = true
			}),
		)

		obj := subclass.New()
		if !classInitCalled {
			t.Error("WithClassInit was not called")
		}
		glib.Destroy(obj)
	})
}

type benchLaunchContext struct {
	gio.AppLaunchContext
	counter int
}

func BenchmarkSubclassLifecycle(b *testing.B) {
	subclass := glib.RegisterSubclass[*benchLaunchContext](
		glib.WithOverrides(func(o *benchLaunchContext) gio.AppLaunchContextOverrides {
			return gio.AppLaunchContextOverrides{
				LaunchFailed: func(startupNotifyId string) {
					o.counter++
				},
			}
		}),
		glib.WithOverrides(func(o *benchLaunchContext) glib.ObjectOverrides {
			return glib.ObjectOverrides{
				Init: func() {
					o.counter++
				},
				Constructed: func() {
					o.counter++
				},
				Dispose: func() {
					o.counter++
				},
			}
		}),
	)

	b.Run("allocate-deallocate", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			obj := subclass.New()
			glib.Destroy(obj)
		}
	})

	b.Run("with-virtual-method-call", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			obj := subclass.New()
			obj.LaunchFailed("bench")
			glib.Destroy(obj)
		}
	})
}

type sizedPaintable struct {
	glib.Object
	width int
}

// TestSubclassInterface verifies that subclasses can implement interfaces, and
// that unimplemented methods keep the interface's default implementation.
func TestSubclassInterface(t *testing.T) {
	subclass := glib.RegisterSubclass[*sizedPaintable](
		glib.WithOverrides(func(p *sizedPaintable) gdk.PaintableOverrides {
			return gdk.PaintableOverrides{
				IntrinsicWidth: func() int { return p.width },
			}
		}),
	)

	if !subclass.Type().IsA(gdk.GTypePaintable) {
		t.Fatal("subclass does not implement GdkPaintable")
	}

	obj := subclass.New()
	obj.width = 42

	paintable := &gdk.Paintable{Object: &obj.Object}
	if w := paintable.IntrinsicWidth(); w != 42 {
		t.Errorf("IntrinsicWidth = %d, want 42", w)
	}
	if h := paintable.IntrinsicHeight(); h != 0 {
		t.Errorf("IntrinsicHeight = %d, want the default 0", h)
	}

	glib.Destroy(obj)
}
