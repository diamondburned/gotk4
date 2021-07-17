// Package gcancel provides a converter between gio.Cancellable and
// context.Context.
package gcancel

// #cgo pkg-config: gio-2.0
// #include <gio/gio.h>
// extern void cancelContextCallback(GCancellable*, gpointer);
import "C"

import (
	"context"
	"time"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/gotk3/gotk3/glib"
)

// Cancellable is a wrapper around the GCancellable object. It satisfies the
// context.Context interface.
type Cancellable struct {
	*glib.Object
	done <-chan struct{}
}

var _ context.Context = (*Cancellable)(nil)

// Cancel will set cancellable to cancelled. It is the same as calling the
// cancel callback given after context creation.
func (c *Cancellable) Cancel() {
	native := (*C.GCancellable)(unsafe.Pointer(c.Native()))
	C.g_cancellable_cancel(native)
}

// IsCancelled checks if a cancellable job has been cancelled.
func (c *Cancellable) IsCancelled() bool {
	native := (*C.GCancellable)(unsafe.Pointer(c.Native()))
	return C.g_cancellable_is_cancelled(native) != 0
}

// Deadline always returns a zero-value of time.Time and false.
func (c *Cancellable) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

// Value always returns nil.
func (c *Cancellable) Value(key interface{}) interface{} {
	return nil
}

// Done returns the channel that's closed once the cancellable is cancelled.
func (c *Cancellable) Done() <-chan struct{} {
	return c.done
}

// Err returns context.Canceled if the cancellable is already cancelled,
// otherwise nil is returned.
func (c *Cancellable) Err() error {
	if c.IsCancelled() {
		return context.Canceled
	}
	return nil
}

// FromContext creates a *gio.Cancellable from the given context. It is mostly
// for internal use. It costs a goroutine to do this, but it should be fairly
// cheap. If FronContext is given a context returned from WithCancellable, then
// the original Cancellable object is returned.
func FromContext(ctx context.Context) *Cancellable {
	// If the context is already a cancellable, then return that.
	if v, ok := ctx.(*Cancellable); ok {
		return v
	}

	cancellable := &Cancellable{
		Object: glib.AssumeOwnership(unsafe.Pointer(C.g_cancellable_new())),
		done:   ctx.Done(),
	}

	go cancelOnDone(ctx, cancellable)
	return cancellable
}

func cancelOnDone(ctx context.Context, cancellable *Cancellable) {
	<-ctx.Done()
	cancellable.Cancel()
}

// WithCancellable creates a new context from the given cancellable object. It
// is mostly for internal use.
func WithCancellable(obj *glib.Object) (context.Context, context.CancelFunc) {
	done := make(chan struct{})
	obj.Connect("cancelled", func() { close(done) })

	cancellable := Cancellable{
		Object: obj,
		done:   done,
	}

	return &cancellable, cancellable.Cancel
}

//export cancelContextCallback
func cancelContextCallback(cancellable *C.GCancellable, ptr C.gpointer) {
	cancel := gbox.Get(uintptr(ptr)).(context.CancelFunc)
	cancel()
}
