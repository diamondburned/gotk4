// Package gcancel provides a converter between gio.Cancellable and
// context.Context.
package gcancel

// #cgo pkg-config: gio-2.0
// #include <gio/gio.h>
import "C"

import (
	"context"
	"runtime"
	"time"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/glib"
)

// Cancellable is a wrapper around the GCancellable object. It satisfies the
// context.Context interface.
type Cancellable struct {
	obj  *glib.Object
	done <-chan struct{}
}

var _ context.Context = (*Cancellable)(nil)

// Object returns the underlying GLib.Object, which might be nil.
func (c *Cancellable) Object() *glib.Object {
	return c.obj
}

// Cancel will set cancellable to cancelled. It is the same as calling the
// cancel callback given after context creation.
func (c *Cancellable) Cancel() {
	defer runtime.KeepAlive(c.obj)

	native := (*C.GCancellable)(unsafe.Pointer(c.obj.Native()))
	C.g_cancellable_cancel(native)
}

// IsCancelled checks if a cancellable job has been cancelled.
func (c *Cancellable) IsCancelled() bool {
	defer runtime.KeepAlive(c.obj)

	native := (*C.GCancellable)(unsafe.Pointer(c.obj.Native()))
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

// CancellableFromContext creates a *gio.Cancellable from the given context. It
// is mostly for internal use; users should use WithCancel instead.
func GCancellableFromContext(ctx context.Context) *glib.Object {
	return fromContext(ctx).obj
}

func fromContext(ctx context.Context) *Cancellable {
	if ctx == nil {
		panic("given ctx is nil")
	}

	// If the context is already a cancellable, then return that.
	if v, ok := ctx.(*Cancellable); ok {
		return v
	}

	// If the context is Background or TODO, then return a nil Cancellable
	// object.
	if ctx == context.Background() || ctx == context.TODO() {
		return &Cancellable{obj: nil, done: nil}
	}

	cancellable := &Cancellable{
		obj:  glib.AssumeOwnership(unsafe.Pointer(C.g_cancellable_new())),
		done: ctx.Done(),
	}

	go cancelOnDone(ctx, cancellable)
	return cancellable
}

func cancelOnDone(ctx context.Context, cancellable *Cancellable) {
	<-ctx.Done()
	cancellable.Cancel()
}

// WithCancel behaves similarly to context.WithCancel, except the created
// context is of type Cancellable. This is useful if the user wants to reuse the
// same Cancellable instance for multiple calls.
//
// This function costs a goroutine to do this unless the given context is
// previously created with WithCancel, is otherwise a Cancellable instance, or
// is an instance from context.Background() or context.TODO(), but it should be
// fairly cheap otherwise.
func WithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	cancellable := fromContext(ctx)
	return cancellable, cancellable.Cancel
}

// WithCancellable creates a new context from the given cancellable object. It
// is mostly for internal use.
func WithCancellable(obj *glib.Object) (context.Context, context.CancelFunc) {
	done := make(chan struct{})
	obj.Connect("cancelled", func() { close(done) })

	cancellable := Cancellable{
		obj:  obj,
		done: done,
	}

	return &cancellable, cancellable.Cancel
}
