package glib

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/closure"
	"github.com/diamondburned/gotk4/pkg/core/intern"
)

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

var typeRegistry struct {
	sync.RWMutex
	types map[reflect.Type]Type
}

// GoTypeInfo is a struct value that's crafted once by every generated
// overriding function to provide context on what it's supposed to be
// overriding, extending, or subclassing.
type GoTypeInfo struct {
	Extends interface{} // GObject{}
	Sizes   struct {
		Instance uint16 // C.sizeof_T
		Class    uint16 // C.sizeof_TClass
	}
}

type goInternTypeInfo struct {
}

// subclassInstance describes the value that goes into the SubclassIx field of
// the intern.Box. It keeps the original reference to the Go type for invokation
// as well as the type that it inherits from.
type subclassInstance struct{}

func registerObjectType(v reflect.Type, typeInfo GoTypeInfo) Type {
	typeRegistry.RLock()
	typ, ok := typeRegistry.types[v]
	typeRegistry.RUnlock()

	if ok {
		return typ
	}

	typeName := C.CString(v.String())
	defer C.free(unsafe.Pointer(typeName))

	// typeInfo
	typeInfo := (*C.GTypeInfo)(C.malloc(C.sizeof_GTypeInfo))
	defer C.free(unsafe.Pointer(typeInfo))

	typeRegistry.Lock()
	defer typeRegistry.Unlock()

	if typeRegistry.types == nil {
		typeRegistry.types = make(map[reflect.Type]Type)
	} else {
		// Recheck. Rare path so prior allocations are fine.
		typ, ok = typeRegistry.types[v]
		if ok {
			return typ
		}
	}

	if typeInfo.Sizes.Class == 0 && typeInfo.Sizes.Instance == 0 {
		if _, ok := typeInfo.Extends.(*GObject); ok {
			typeInfo.Sizes.Instance = C.sizeof_GObject
			typeInfo.Sizes.Class = C.sizeof_GObjectClass
		} else {
			panic("missing typeInfo.Sizes")
		}
	}

	typeInfo.base_init = nil
	typeInfo.base_finalize = nil
	typeInfo.class_finalize = nil // TODO
	typeInfo.n_preallocs = 0
	typeInfo.value_table = nil
	typeInfo.instance_size = C.guint16(typeInfo.Sizes.Instance)
	typeInfo.class_size = C.guint16(typeInfo.Sizes.Class)

	// TODO:
	// - instance_init
	// - class_init

	// NOTE that we're making this function ONLY FOR REGISTERING TYPES!
	// {class,instance}_init functions will take in the GObject class pointer
	// and the class data pointer, in which the class data pointer is the fake
	// pointer to our data stored on gbox. What we should do is store the TYPE
	// information into the class data, then combine that with the class pointer
	// to grab the intern box. The intern box is only made in the new object
	// function.

	// Delegated types:
	//   class_data -> goInternTypeInfo
	//   intern -> subclassInstance

	// TODO: since GObjectTypeClass and GObject are different, maybe we can
	// allocate a pointer fields in it for the GObject (set after the fact).
}

func implementInterfaceType(v interface{}) {
	typeRegistry.RLock()
	typ, ok := typeRegistry.types[v]
	typeRegistry.RUnlock()

	if !ok {
		panic("value type " + v.String() + " is not registered")
	}

	panic("TODO implement implementInterfaceType")
}

func newGoObject(v Objector, typeInfo GoTypeInfo, props map[string]interface{}) *Object {
	var obj *C.GObject
	typ := registerObjectType(reflect.TypeOf(v), typeInfo)

	if len(props) > 0 {
		// We can allocate this on the Go heap, probably.
		names := make([]*C.char, len(props)+1)
		values := make([]value, len(props)+1)

		intern.Escape(names)
		intern.Escape(values)

		defer func() {
			for _, n := range names {
				C.free(unsafe.Pointer(n))
			}
			for _, v := range values {
				v.unset()
			}
		}()

		ix := 0
		for k, v := range props {
			names[ix] = C.CString(k)
			InitValueFromGo(&Value{&values[ix]}, v)
			ix++
		}

		obj = C.g_object_new_with_properties(
			typ,
			C.guint(len(props)),
			(**C.char)(unsafe.Pointer(&names[0])),
			(*C.GValue)(unsafe.Pointer(&values[0])),
		)
	} else {
		obj = C.g_object_newv(typ, 0, nil)
	}

	box := intern.Get(unsafe.Pointer(obj), true)
	box.Set(intern.ClosureIx, closure.NewRegistry())
	box.Set(intern.SubclassIx, &subclassInstance{})
	box.Seal()

	return &Object{box: box}
}

//export _gotk4_glibObjectFinalize
func _gotk4_glibObjectFinalize(obj *C.GObjectType)

// isGoSubclass returns true if the object is a Go subclassed object.
func (o *Object) isGoSubclass() bool {
	return o.box.Get(intern.SubclassIx) != nil
}

func NewClass(v interface{}) Class {
	return NewClassWithProperties(v, nil)
}

func NewClassWithProperties(v interface{}, props map[string]interface{}) Class {}
