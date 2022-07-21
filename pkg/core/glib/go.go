package glib

import (
	"reflect"
	"unsafe"
)

type subclassType struct {
	goType   reflect.Type
	glibType Type
}

// subclassInstance is stored inside an intern.Box.
type subclassInstance struct {
	goObject   Objector
	glibObject uintptr

	subclassType subclassType
}

// GoPrivateFromObject panics.
func GoPrivateFromObject(obj unsafe.Pointer) interface{} {
	panic("implement me")
}

// Subclass initializes the given object as a new object that subclasses from
// the given type. An example of such a type would be:
//
//    type Gadget struct {
//        gtk.Widget // must be first and must be embedded
//        state string
//    }
//
//    func NewGadget(state string) *Gadget {
//        g := &Gadget{state: state}
//        glib.Subclass(g, gtk.GTypeWidget())
//        return g
//    }
//
func Subclass(obj Objector, parent Type) {

}
