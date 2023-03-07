//go:build !no_string_connect

package glib

// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"log"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/closure"
)

// Connect is a wrapper around g_signal_connect_closure(). f must be a function
// with at least one parameter matching the type it is connected to.
//
// It is optional to list the rest of the required types from GTK, as values
// that don't fit into the function parameter will simply be ignored; however,
// extraneous types will trigger a runtime panic. Arguments for f must be a
// matching Go equivalent type for the C callback, or an interface type which
// the value may be packed in. If the type is not suitable, a runtime panic will
// occur when the signal is emitted.
func (v *Object) Connect(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(false, detailedSignal, f)
}

// ConnectAfter is a wrapper around g_signal_connect_closure(). The difference
// between Connect and ConnectAfter is that the latter will be invoked after the
// default handler, not before. For more information, refer to Connect.
func (v *Object) ConnectAfter(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(true, detailedSignal, f)
}

func (v *Object) connectClosure(after bool, detailedSignal string, f interface{}) SignalHandle {
	fs := closure.NewFuncStack(f, 2)

	cstr := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(cstr))

	gclosure := closureNew(v, fs)
	c := C.g_signal_connect_closure(C.gpointer(v.Native()), (*C.gchar)(cstr), gclosure, gbool(after))

	runtime.KeepAlive(v)

	return SignalHandle(c)
}

// NewClosure creates a new closure for the given object.
func NewClosure(v *Object, f interface{}) unsafe.Pointer {
	return unsafe.Pointer(closureNew(v, f))
}

// closureNew creates a new GClosure that's bound to the current object and adds
// its callback function to the internal registry. It's exported for visibility
// to other gotk3 packages and should not be used in a regular application.
func closureNew(v *Object, f interface{}) *C.GClosure {
	fs, ok := f.(*closure.FuncStack)
	if !ok {
		fs = closure.NewFuncStack(f, 2)
	}

	gclosure := C.g_closure_new_simple(C.sizeof_GClosure, nil)

	closures := closure.RegistryType.Get(v.box)
	closures.Register(unsafe.Pointer(gclosure), fs)

	// C.g_object_watch_closure(v.native(), gclosure)
	C.g_closure_set_meta_marshal(gclosure, C.gpointer(v.Native()), (*[0]byte)(C._gotk4_goMarshal))
	C.g_closure_add_finalize_notifier(gclosure, C.gpointer(v.Native()), (*[0]byte)(C._gotk4_removeClosure))

	return gclosure
}

// WipeAllClosures wipes all the Go closures associated with the given object
// away. BE EXTREMELY CAREFUL WHEN USING THIS FUNCTION! If not careful, your
// program WILL crash!
//
// There is only one specific use case for this function: if your object has
// closures connected to it where these closures capture the object itself, then
// it might create a cyclic dependency on the GC, preventing its finalizer from
// ever running. This will cause the program to leak memory. As a temporary
// hack, this function is introduced for cases where the programmer knows for
// sure that the object will never be used again, and it is significant enough
// of a leak that having a workaround is better than not.
func WipeAllClosures(objector Objector) {
	obj := BaseObject(objector)

	closures := closure.RegistryType.TryGet(obj.box)
	if closures == nil {
		return
	}

	log.Printf("wiping all closures (%v) for %v (%s)", closures, obj, obj.Type())

	closures.Range(func(gclosure unsafe.Pointer) {
		cclosure := (*C.GClosure)(gclosure)
		closures.Delete(gclosure)
		C.g_closure_invalidate(cclosure)
	})

	closure.RegistryType.Delete(obj.box)
	log.Printf("wiped all closures (%v) for %v (%s)", closures, obj, obj.Type())
}
