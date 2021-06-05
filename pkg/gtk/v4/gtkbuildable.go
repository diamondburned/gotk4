// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_buildable_get_type()), F: marshalBuildable},
	})
}

// BuildableOverrider contains methods that are overridable. This
// interface is a subset of the interface Buildable.
type BuildableOverrider interface {
	// AddChild adds a child to @buildable. @type is an optional string
	// describing how the child should be added.
	AddChild(builder Builder, child gextras.Objector, typ string)

	ConstructChild(builder Builder, name string) gextras.Objector
	// CustomFinished: similar to gtk_buildable_parser_finished() but is called
	// once for each custom tag handled by the @buildable.
	CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagEnd: called at the end of each custom element handled by the
	// buildable.
	CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{})

	ID() string
	// InternalChild retrieves the internal child called @childname of the
	// @buildable object.
	InternalChild(builder Builder, childname string) gextras.Objector

	ParserFinished(builder Builder)

	SetBuildableProperty(builder Builder, name string, value *externglib.Value)

	SetID(id string)
}

// Buildable: gtkBuildable allows objects to extend and customize their
// deserialization from [GtkBuilder UI descriptions][BUILDER-UI]. The interface
// includes methods for setting names and properties of objects, parsing custom
// tags and constructing child objects.
//
// The GtkBuildable interface is implemented by all widgets and many of the
// non-widget objects that are provided by GTK. The main user of this interface
// is Builder. There should be very little need for applications to call any of
// these functions directly.
//
// An object only needs to implement this interface if it needs to extend the
// Builder format or run any extra routines at deserialization time.
type Buildable interface {
	gextras.Objector
	BuildableOverrider

	// BuildableID gets the ID of the @buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the @buildable.
	BuildableID() string
}

// buildable implements the Buildable interface.
type buildable struct {
	gextras.Objector
}

var _ Buildable = (*buildable)(nil)

// WrapBuildable wraps a GObject to a type that implements interface
// Buildable. It is primarily used internally.
func WrapBuildable(obj *externglib.Object) Buildable {
	return Buildable{
		Objector: obj,
	}
}

func marshalBuildable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuildable(obj), nil
}

// BuildableID gets the ID of the @buildable object.
//
// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
// used to construct the @buildable.
func (b buildable) BuildableID() string {
	var arg0 *C.GtkBuildable

	arg0 = (*C.GtkBuildable)(unsafe.Pointer(b.Native()))

	var cret *C.char
	var goret1 string

	cret = C.gtk_buildable_get_buildable_id(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

type BuildableParseContext struct {
	native C.GtkBuildableParseContext
}

// WrapBuildableParseContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBuildableParseContext(ptr unsafe.Pointer) *BuildableParseContext {
	if ptr == nil {
		return nil
	}

	return (*BuildableParseContext)(ptr)
}

func marshalBuildableParseContext(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBuildableParseContext(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *BuildableParseContext) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Element retrieves the name of the currently open element.
//
// If called from the start_element or end_element handlers this will give the
// element_name as passed to those functions. For the parent elements, see
// gtk_buildable_parse_context_get_element_stack().
func (c *BuildableParseContext) Element() string {
	var arg0 *C.GtkBuildableParseContext

	arg0 = (*C.GtkBuildableParseContext)(unsafe.Pointer(c.Native()))

	var cret *C.char
	var goret1 string

	cret = C.gtk_buildable_parse_context_get_element(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// ElementStack retrieves the element stack from the internal state of the
// parser.
//
// The returned Array is an array of strings where the last item is the
// currently open tag (as would be returned by
// gtk_buildable_parse_context_get_element()) and the previous item is its
// immediate parent.
//
// This function is intended to be used in the start_element and end_element
// handlers where gtk_buildable_parse_context_get_element() would merely return
// the name of the element that is being processed.
func (c *BuildableParseContext) ElementStack() []string {
	var arg0 *C.GtkBuildableParseContext

	arg0 = (*C.GtkBuildableParseContext)(unsafe.Pointer(c.Native()))

	var cret *C.GPtrArray
	var goret1 []string

	cret = C.gtk_buildable_parse_context_get_element_stack(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (*C.GPtrArray)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			goret1[i] = C.GoString(src)
		}
	}

	return goret1
}

// Position retrieves the current line number and the number of the character on
// that line. Intended for use in error messages; there are no strict semantics
// for what constitutes the "current" line number other than "the best number we
// could come up with for error messages."
func (c *BuildableParseContext) Position() (lineNumber int, charNumber int) {
	var arg0 *C.GtkBuildableParseContext

	arg0 = (*C.GtkBuildableParseContext)(unsafe.Pointer(c.Native()))

	var arg1 *C.int
	var ret1 int
	var arg2 *C.int
	var ret2 int

	C.gtk_buildable_parse_context_get_position(arg0, &arg1, &arg2)

	ret1 = *C.int(arg1)
	ret2 = *C.int(arg2)

	return ret1, ret2
}

// Pop completes the process of a temporary sub-parser redirection.
//
// This function exists to collect the user_data allocated by a matching call to
// gtk_buildable_parse_context_push(). It must be called in the end_element
// handler corresponding to the start_element handler during which
// gtk_buildable_parse_context_push() was called. You must not call this
// function from the error callback -- the @user_data is provided directly to
// the callback in that case.
//
// This function is not intended to be directly called by users interested in
// invoking subparsers. Instead, it is intended to be used by the subparsers
// themselves to implement a higher-level interface.
func (c *BuildableParseContext) Pop() interface{} {
	var arg0 *C.GtkBuildableParseContext

	arg0 = (*C.GtkBuildableParseContext)(unsafe.Pointer(c.Native()))

	var cret C.gpointer
	var goret1 interface{}

	cret = C.gtk_buildable_parse_context_pop(arg0)

	goret1 = C.gpointer(cret)

	return goret1
}