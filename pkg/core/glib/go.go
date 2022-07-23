package glib

// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
//
// extern void _gotk4_gobject_init_class(gpointer, gpointer);
// extern void _gotk4_gobject_init_instance(GTypeInstance*, gpointer);
// extern void _gotk4_gobject_finalize_class(gpointer, gpointer);
import "C"

import (
	"log"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// RegisteredSubclass is a type that described a registered Go subclass type.
type RegisteredSubclass[T any] registeredSubclass

func rtypeElem[T any]() reflect.Type {
	var zero T
	rtype := reflect.TypeOf(zero)
	if rtype.Kind() != reflect.Ptr {
		log.Panicln("expected type pointer")
	}
	rtype = rtype.Elem()
	return rtype
}

type registerOpts struct {
	paramSpecs []*ParamSpec
}

// RegisterOptsFunc is a function type that modifies the behavior of a
// RegisterSubclass call.
type RegisterOptsFunc func(*registerOpts)

// WithParamSpecs adds additional ParamSpecs into a type for its properties.
func WithParamSpecs(paramSpecs []*ParamSpec) RegisterOptsFunc {
	return func(opts *registerOpts) { opts.paramSpecs = append(opts.paramSpecs, paramSpecs...) }
}

// RegisterSubclass is RegisterSubclassWithConstructor, but a zero-value
// instance of T is automatically created.
func RegisterSubclass[T any](opts ...RegisterOptsFunc) *RegisteredSubclass[T] {
	rtype := rtypeElem[T]()
	return RegisterSubclassWithConstructor(func() T {
		return reflect.New(rtype).Interface().(T)
	}, opts...)
}

// RegisterSubclassWithConstructor registers a new type T that is a subclass of
// its parent type, which is the first field that must be embedded.
//
// ctor has to be idempotent (i.e. can be called multiple times w/o side
// effects).
func RegisterSubclassWithConstructor[T any](ctor func() T, opts ...RegisterOptsFunc) *RegisteredSubclass[T] {
	rtype := rtypeElem[T]()

	knownTypesMut.RLock()
	subclass, ok := knownTypes[rtype]
	knownTypesMut.RUnlock()
	if ok {
		return castRegisteredSubclass[T](subclass)
	}

	knownTypesMut.Lock()
	subclass, ok = knownTypes[rtype]
	if ok {
		knownTypesMut.Unlock()
		return castRegisteredSubclass[T](subclass)
	}

	subclass = registerSubclass(rtype, func() any { return ctor() }, opts)
	knownTypes[rtype] = subclass
	knownTypesMut.Unlock()

	return castRegisteredSubclass[T](subclass)
}

func castRegisteredSubclass[T any](src *registeredSubclass) *RegisteredSubclass[T] {
	return (*RegisteredSubclass[T])(unsafe.Pointer(src))
}

// New creates an instance of the subclass object with no properties.
func (r *RegisteredSubclass[T]) New() T {
	return r.NewWithProperties(nil)
}

// NewWithProperties creates an instance of the subclass object with the given
// properties.
func (r *RegisteredSubclass[T]) NewWithProperties(properties map[string]any) T {
	var names []*C.char
	var values []C.GValue

	if len(properties) > 0 {
		names = make([]*C.char, 0, len(properties))
		values = make([]C.GValue, 0, len(properties))

		for name, value := range properties {
			cname := (*C.char)(C.CString(name))
			defer C.free(unsafe.Pointer(cname))

			gvalue := NewValue(value)
			defer runtime.KeepAlive(gvalue)

			names = append(names, cname)
			values = append(values, *gvalue.gvalue)
		}
	}

	cval := C.g_object_new_with_properties(
		C.GType(r.gType),
		C.guint(len(properties)),
		&names[0],
		&values[0],
	)

	gobject := AssumeOwnership(unsafe.Pointer(cval))
	return gobject.Cast().(T)
}

// Type returns the GType of the registered Go subclass.
func (r *RegisteredSubclass[T]) Type() Type {
	return r.gType
}

type registeredSubclass struct {
	goType      reflect.Type
	gType       Type
	parentType  ClassTypeInfo
	constructor func() any
	properties  []subclassProperty
	paramSpecs  []*C.GParamSpec
}

type subclassProperty struct {
	name     string
	minit    reflect.Method
	fieldIdx []int
}

var (
	knownTypes    = map[reflect.Type]*registeredSubclass{}
	knownTypesMut sync.RWMutex
)

var (
	knownTypePtrs    []reflect.Type
	knownTypePtrsMut sync.RWMutex
)

const knownTypePtrsOffset = 4096 // see gbox.minLegalPointer

func rtypeNewData(rtype reflect.Type) C.gpointer {
	knownTypePtrsMut.Lock()
	data := uintptr(len(knownTypePtrs) + knownTypePtrsOffset)
	knownTypePtrs = append(knownTypePtrs, rtype)
	knownTypePtrsMut.Unlock()
	return C.gpointer(unsafe.Pointer(data))
}

func rtypeFromData(data C.gpointer) reflect.Type {
	knownTypePtrsMut.RLock()
	rtype := knownTypePtrs[uintptr(data)-knownTypePtrsOffset]
	knownTypePtrsMut.RUnlock()
	return rtype
}

func subclassFromData(data C.gpointer) *registeredSubclass {
	rtype := rtypeFromData(data)

	knownTypesMut.RLock()
	subclass, ok := knownTypes[rtype]
	knownTypesMut.RUnlock()
	if !ok {
		log.Panicf("type %s is not registered", rtype)
	}

	return subclass
}

func registerSubclass(rtype reflect.Type, ctor func() any, optsFuncs []RegisterOptsFunc) *registeredSubclass {
	subclass := &registeredSubclass{
		goType:      rtype,
		parentType:  extractParentType(rtype),
		constructor: ctor,
	}

	var opts registerOpts
	for _, fn := range optsFuncs {
		fn(&opts)
	}

	if len(opts.paramSpecs) > 0 {
		subclass.paramSpecs = make([]*C.GParamSpec, len(opts.paramSpecs)+1)
		for i, spec := range opts.paramSpecs {
			// [0] is reserved by GLib.
			subclass.paramSpecs[i+1] = spec.intern
			// Permanently take a reference, since we're globalling this forever
			// anyway.
			C.g_param_spec_ref(spec.intern)
		}
	}

	// Scan for Property fields.
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if !strings.HasSuffix(field.PkgPath, "/core/glib") {
			continue
		}
		// Hack!
		if !strings.HasPrefix(field.Type.Name(), "Property[") {
			continue
		}

		name := field.Tag.Get("glib")
		if name == "" {
			log.Panicf("field %s of type %s has no glib tag", field.Name, rtype)
		}

		minit, ok := field.Type.MethodByName("init")
		if !ok {
			log.Panicf("BUG: (*Property[T]).init missing")
		}

		subclass.properties = append(subclass.properties, subclassProperty{
			name:     name,
			minit:    minit,
			fieldIdx: field.Index,
		})
	}

	typeName := subclass.goType.String()

	var typeInfo C.GTypeInfo

	// Why are these ushort anyway?
	typeInfo.class_size = C.gushort(subclass.parentType.ClassSize)
	typeInfo.class_data = C.gconstpointer(rtypeNewData(rtype))
	typeInfo.class_init = C.GClassInitFunc(C._gotk4_gobject_init_class)
	typeInfo.class_finalize = C.GClassFinalizeFunc(C._gotk4_gobject_finalize_class)

	typeInfo.instance_size = C.gushort(subclass.parentType.InstanceSize)
	typeInfo.instance_init = C.GInstanceInitFunc(C._gotk4_gobject_init_instance)

	cTypeName := (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(
		C.GType(subclass.parentType.GType),
		cTypeName,
		&typeInfo,
		C.GTypeFlags(0))
	subclass.gType = Type(gtype)

	knownTypes[rtype] = subclass
	return subclass
}

func extractParentType(rtype reflect.Type) ClassTypeInfo {
	if rtype.Kind() != reflect.Struct {
		log.Panicln("given type is not a struct or a *struct")
	}

	field := rtype.Field(0)
	if !field.Anonymous {
		log.Panicln("first field (parent) must be embedded")
	}

	typeInfo, ok := classTypeInfos[field.Type]
	if !ok {
		// TODO: allow inheriting from a Go type.
		log.Panicln("unknown type", field.Type)
	}

	return typeInfo
}

func (r *registeredSubclass) setParent(instance any, parent unsafe.Pointer) {
	gobject := Take(unsafe.Pointer(parent))

	// We want to set the first field of our new Go class instance, which is the
	// parent, to be the initialized parent. There's a pretty nifty way of doing
	// this: we can hijack our base methods and use that.

	// We have to hijack anything that stores a pointer to a GObject.
	// InitiallyUnowned is one. The user can also do that themselves,
	// but calling gobject.baseObject on that will explode.
	if initiallyUnowned, ok := instance.(initiallyUnownedor); ok {
		v := initiallyUnowned.baseInitiallyUnowned()
		v.Object = gobject
		return
	}

	// Probably an actual Objector. Grab the baseObject and hijack its internal
	// pointer.
	base := gobject.baseObject()
	if base.box != nil {
		log.Panicf("cannot construct subclass %s instance with non-nil parent %s (%p)",
			r.goType,
			r.parentType.GoType,
			base.box.GObject(),
		)
	}

	base.box = gobject.box
}

// ClassTypeInfo contains autogenerated data made by the generator that binds
// relevant subclassing information.
type ClassTypeInfo struct {
	GType  Type
	GoType reflect.Type
	// InitClass will be called to initialize a class using the given Go value.
	// It should type assert goValue and sets the available functions into the
	// gclass (*GObjectTypeClass) type parameter.
	InitClass func(gclass unsafe.Pointer, goValue any)
	// FinalizeClass will be called right before the class is finalized.
	FinalizeClass func(gclass unsafe.Pointer, goValue any)
	// ClassSize is the size of an ObjectClass. (optional)
	ClassSize uint32
	// InstanceSize is the size of an Object. (optional)
	InstanceSize uint32
}

var classTypeInfos = make(map[reflect.Type]ClassTypeInfo, 1024)

// RegisterClassInfo registers the given class type info.
func RegisterClassInfo(typeInfo ClassTypeInfo) {
	if typeInfo.ClassSize == 0 || typeInfo.InstanceSize == 0 {
		var query C.GTypeQuery
		C.g_type_query(C.GType(typeInfo.GType), &query)

		if query._type == 0 {
			log.Panicln("unknown GType", typeInfo.GType)
		}

		typeInfo.ClassSize = uint32(query.class_size)
		typeInfo.InstanceSize = uint32(query.instance_size)
	}

	typeInfo.GoType = typeInfo.GoType.Elem()
	classTypeInfos[typeInfo.GoType] = typeInfo
}

// registeredGClass binds gclass (which is a *GTypeClassInfo describing XClass
// structs) to data (which is the ID that gets us a *registeredSubclass).
// It exists for init_instance.
//
// gpointer (*GTypeClassInfo) -> *privateGoInstance
var registeredPrivateInstances sync.Map

// privateGoInstance maps the two private fields.
type privateGoInstance struct {
	data       C.gpointer // constant
	instanceID C.gpointer // used for finalizing
}

func privateFromInstance(obj unsafe.Pointer) *privateGoInstance {
	gtype := typeFromObject(obj)

	private := C.g_type_instance_get_private((*C.GTypeInstance)(obj), C.GType(gtype))
	if private == nil {
		log.Panicf("cannot get private from unknown object %s (%p)", Type(gtype), obj)
	}

	return (*privateGoInstance)(unsafe.Pointer(private))
	// return privateGoInstance{
	// 	data:       *(*C.gpointer)(unsafe.Add(unsafe.Pointer(private), 0*ptrsz)),
	// 	instanceID: *(*C.gpointer)(unsafe.Add(unsafe.Pointer(private), 1*ptrsz)),
	// }
}

func (p *privateGoInstance) subclass() *registeredSubclass {
	return subclassFromData(p.data)
}

func (p *privateGoInstance) instance() any {
	return gbox.Get(uintptr(p.instanceID))
}

// GoObjectFromInstance returns the Go value from a given GObject pointer whose
// type is that of a Go subclass type. This function is only ever used
// internally.
func GoObjectFromInstance(instance unsafe.Pointer) any {
	private := privateFromInstance(instance)
	return private.instance()
}

//export _gotk4_gobject_init_class
func _gotk4_gobject_init_class(gclass, data C.gpointer) {
	subclass := subclassFromData(data)

	_, dup := registeredPrivateInstances.LoadOrStore(gclass, &privateGoInstance{data: data})
	if dup {
		log.Panicf("init_class called on the same gclass %s (%p) twice", subclass.goType, gclass)
	}

	// Add 2 pointers, one for our constant RegisteredSubclass type info, one
	// for the boxed instantiated value.
	// C.g_type_class_add_private(gclass, unsafe.Sizeof(privateGoInstance{}))
	C.g_type_add_instance_private(C.GType(subclass.gType), C.size_t(unsafe.Sizeof(privateGoInstance{})))

	// Install properties, if any.
	if len(subclass.paramSpecs) > 0 {
		C.g_object_class_install_properties(
			(*C.GObjectClass)(gclass),
			C.guint(len(subclass.paramSpecs)),
			&subclass.paramSpecs[0],
		)
	}

	// Have our generated code crawl through our Go type and set whatever method
	// it can into the given gclass field.
	nilValue := reflect.NewAt(subclass.goType, nil).Elem()
	subclass.parentType.InitClass(unsafe.Pointer(gclass), nilValue.Interface())
}

//export _gotk4_gobject_finalize_class
func _gotk4_gobject_finalize_class(gclass, data C.gpointer) {
	subclass := subclassFromData(data)

	// Unregister our gclass.
	privateV, ok := registeredPrivateInstances.LoadAndDelete(gclass)
	if !ok {
		log.Panicf("cannot delete known gclass %s (%p)", subclass.goType, gclass)
		return
	}

	private := privateV.(*privateGoInstance)

	// Call the custom finalizer if any.
	instance := private.instance()
	subclass.parentType.FinalizeClass(unsafe.Pointer(gclass), instance)

	// Unbind our instance from the global store.
	gbox.Delete(uintptr(private.instanceID))
}

//export _gotk4_gobject_init_instance
func _gotk4_gobject_init_instance(obj *C.GTypeInstance, gclass C.gpointer) {
	// Reminder: obj of type *GTypeInstance IS a regular *GObject if we're
	// initializing a class! We can consider it as such.

	// Grab our registeredSubclass ID.
	privateV, ok := registeredPrivateInstances.Load(gclass)
	if !ok {
		log.Panicf(
			"init_instance called on unregistered gclass %s (%p)",
			typeFromObject(unsafe.Pointer(obj)), gclass)
	}

	private := privateV.(*privateGoInstance)
	subclass := private.subclass()

	// Allocate and construct a new instance.
	instance := subclass.constructor()
	instanceID := gbox.Assign(instance)

	// Initialize its properties.
	if len(subclass.properties) > 0 {
		rval := reflect.ValueOf(instance)
		for _, propProto := range subclass.properties {
			propProto.minit.Func.Call([]reflect.Value{
				reflect.ValueOf(rval.FieldByIndex(propProto.fieldIdx).Addr()),
				reflect.ValueOf(rval.Convert(rtypeObjector)),
				reflect.ValueOf(unsafe.Pointer(gclass)),
				reflect.ValueOf(reflect.ValueOf(propProto.name)),
			})
		}
	}

	private.instanceID = C.gpointer(instanceID)

	// Copy our fully initialized private instance values to GLib's allocated
	// object private one.
	*privateFromInstance(unsafe.Pointer(obj)) = *private

	// Bind our new Go class' parent field.
	subclass.setParent(private.instance, unsafe.Pointer(obj))
}

// Property describes a Go GObject property. It is used as an alternative to
// manually written property getter/setters.
type Property[T any] struct {
	parent Objector
	name   string
	rtype  reflect.Type
}

var rtypeObjector = reflect.TypeOf(Objector(nil))

func (p *Property[T]) init(obj Objector, gclass unsafe.Pointer, name string) {
	p.parent = obj
	p.name = name

	var z T
	p.rtype = reflect.TypeOf(z)

	// Verify that this property exist. Ideally, we want to type check as well.
	cname := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cname))

	spec := C.g_object_class_find_property((*C.GObjectClass)(gclass), cname)
	if spec == nil {
		log.Panicf("unknown property %q for type %T", name, obj)
	}
}

// Set sets the new property.
func (p *Property[T]) Set(v T) {
	base := BaseObject(p.parent)
	base.SetObjectProperty(p.name, v)
}

// Get gets the value of the property.
func (p *Property[T]) Get() T {
	base := BaseObject(p.parent)
	pval := base.ObjectProperty(p.name)
	return reflect.ValueOf(pval).Convert(p.rtype).Interface().(T)
}

// Notify calls f everytime the property changes.
func (p *Property[T]) Notify(f func(T)) SignalHandle {
	base := BaseObject(p.parent)
	return base.NotifyProperty(p.name, func() { f(p.Get()) })
}
