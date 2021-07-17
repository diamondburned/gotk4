// Package gcancel provides a converter between gio.Cancellable and
// context.Context.
package gcancel

// #cgo pkg-config: gio-2.0
// #include <gio/gio.h>
import "C"

import (
	"context"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/gotk3/gotk3/glib"
)

// FromContext creates a *gio.Cancellable from the given context. It is mostly
// for internal use. It costs a goroutine to do this, but it should be fairly
// cheap.
func FromContext(ctx context.Context) *glib.Object {
	cancellable := C.g_cancellable_new()
	go cancelOnDone(ctx, cancellable)

	return glib.AssumeOwnership(unsafe.Pointer(cancellable))
}

func cancelOnDone(ctx context.Context, cancellable *C.GCancellable) {
	<-ctx.Done()
	g_cancellable_cancel(cancellable)
}

// WithCancellable creates a new context from the given cancellable.
func WithCancellable(cancellable gextras.Objector) (context.Context, context.CancelFunc) {}
