package glib

import "reflect"

// Initializer is an interface that a class can optionally implement.
type Initializer interface {
	Init()
}

type Disposer interface {
	Dispose()
}

type Finalizer interface {
	Finalize()
}

type typeRegistry struct {
	m map[reflect.Type]Type
}

// GoObject is the base type for all Go types that are meant to be shared to
// GLib for subclassing or implementing an interface. It is different from the
// regular Object type in that it is an object from Go, not from C.
//
//
type GoObject struct {
	base *Object
}

// IsGoObject returns true if the given value is a GoObject.
func IsGoObject(v interface{}) bool {
	_, ok := v.(interface{ goobject() })
	return ok
}

func NewClass(v interface{}) Class {
	return NewClassWithProperties(v, nil)
}

func NewClassWithProperties(v interface{}, props map[string]interface{}) Class {}

// Type returns the Go Object's type.
func (g *GoObject) Type() Type {
	return g.base.Type()
}

func (g *GoObject) goobject() {}
