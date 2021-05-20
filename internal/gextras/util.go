package gextras

// #include <glib/glib.h>
// #include "util.h"
import "C"

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func Gpointer(ptr unsafe.Pointer) C.gpointer {
	return C.conptr(ptr)
}

func Cbool(val bool) C.gboolean {
	if val {
		return C.TRUE
	}
	return C.FALSE
}

func Gobool(val C.gboolean) bool {
	return val != C.FALSE
}

// These functions are copied from gotk3. For some reasons, they weren't
// exported?

// CastInternal casts the given object to the appropriate Go struct, but returns it as interface for later type assertions.
// The className is the results of C.object_get_class_name(c) called on the native object.
// The obj is the result of glib.Take(unsafe.Pointer(c)), used as a parameter for the wrapper functions.
func CastInternal(className string, obj *glib.Object) (interface{}, error) {
	fn, ok := gtk.WrapMap[className]
	if !ok {
		return nil, errors.New("unrecognized class name '" + className + "'")
	}

	// Check that the wrapper function is actually a function
	rf := reflect.ValueOf(fn)
	if rf.Type().Kind() != reflect.Func {
		return nil, errors.New("wraper is not a function")
	}

	// Call the wraper function with the *glib.Object as first parameter
	// e.g. "wrapWindow(obj)"
	v := reflect.ValueOf(obj)
	rv := rf.Call([]reflect.Value{v})

	// At most/max 1 return value
	if len(rv) != 1 {
		return nil, errors.New("wrapper did not return")
	}

	// Needs to be a pointer of some sort
	if k := rv[0].Kind(); k != reflect.Ptr {
		return nil, fmt.Errorf("wrong return type %s", k)
	}

	// Only get an interface value, type check will be done in more specific functions
	return rv[0].Interface(), nil
}

// Cast takes a native GObject and casts it to the appropriate Go struct.
//TODO change all wrapFns to return an IObject
//^- not sure about this TODO. This may make some usages of the wrapper functions quite verbose, no?
func Cast(c *C.GObject) (glib.IObject, error) {
	ptr := unsafe.Pointer(c)
	var (
		className = C.GoString(C.object_get_class_name(c))
		obj       = glib.Take(ptr)
	)

	intf, err := castInternal(className, obj)
	if err != nil {
		return nil, err
	}

	ret, ok := intf.(glib.IObject)
	if !ok {
		return nil, errors.New("did not return an IObject")
	}

	return ret, nil
}

// CastWidget takes a native GtkWidget and casts it to the appropriate Go struct.
func CastWidget(c *C.GtkWidget) (gtk.IWidget, error) {
	ptr := unsafe.Pointer(c)
	var (
		className = C.GoString(C.object_get_class_name(C.toGObject(ptr)))
		obj       = glib.Take(ptr)
	)

	intf, err := castInternal(className, obj)
	if err != nil {
		return nil, err
	}

	ret, ok := intf.(gtk.IWidget)
	if !ok {
		return nil, fmt.Errorf("expected value of type IWidget, got %T", intf)
	}

	return ret, nil
}
