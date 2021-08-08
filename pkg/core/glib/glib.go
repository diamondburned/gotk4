// Package glib provides some hand-written GObject and GLib bindings.
package glib

// #cgo pkg-config: gio-2.0 glib-2.0 gobject-2.0
// #include <gio/gio.h>
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"

import (
	"errors"
	"log"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/closure"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/intern"
)

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	return b != 0
}

// InitI18n initializes the i18n subsystem. It runs the following C code:
//
//    setlocale(LC_ALL, "");
//    bindtextdomain(domain, dir);
//    bind_textdomain_codeset(domain, "UTF-8");
//    textdomain(domain);
//
func InitI18n(domain, dir string) {
	domainStr := C.CString(domain)
	defer C.free(unsafe.Pointer(domainStr))

	dirStr := C.CString(dir)
	defer C.free(unsafe.Pointer(dirStr))

	C.init_i18n(domainStr, dirStr)
}

// Local localizes a string using gettext.
func Local(input string) string {
	cstr := C.CString(input)
	defer C.free(unsafe.Pointer(cstr))

	return C.GoString(C.localize(cstr))
}

// Type is a representation of GLib's GType.
type Type uint

const (
	TypeInvalid   Type = C.G_TYPE_INVALID
	TypeNone      Type = C.G_TYPE_NONE
	TypeInterface Type = C.G_TYPE_INTERFACE
	TypeChar      Type = C.G_TYPE_CHAR
	TypeUchar     Type = C.G_TYPE_UCHAR
	TypeBoolean   Type = C.G_TYPE_BOOLEAN
	TypeInt       Type = C.G_TYPE_INT
	TypeUint      Type = C.G_TYPE_UINT
	TypeLong      Type = C.G_TYPE_LONG
	TypeUlong     Type = C.G_TYPE_ULONG
	TypeInt64     Type = C.G_TYPE_INT64
	TypeUint64    Type = C.G_TYPE_UINT64
	TypeEnum      Type = C.G_TYPE_ENUM
	TypeFlags     Type = C.G_TYPE_FLAGS
	TypeFloat     Type = C.G_TYPE_FLOAT
	TypeDouble    Type = C.G_TYPE_DOUBLE
	TypeString    Type = C.G_TYPE_STRING
	TypePointer   Type = C.G_TYPE_POINTER
	TypeBoxed     Type = C.G_TYPE_BOXED
	TypeParam     Type = C.G_TYPE_PARAM
	TypeObject    Type = C.G_TYPE_OBJECT
	TypeVariant   Type = C.G_TYPE_VARIANT
)

// FundamentalType returns the fundamental type of the given actual type.
func FundamentalType(actual Type) Type {
	return Type(C._g_value_fundamental(C.GType(actual)))
}

// IsValue checks whether the passed in type can be used for g_value_init().
func (t Type) IsValue() bool {
	return gobool(C._g_type_is_value(C.GType(t)))
}

// Name is a wrapper around g_type_name().
func (t Type) Name() string {
	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
}

// String calls t.Name(). It satisfies fmt.Stringer.
func (t Type) String() string {
	return t.Name()
}

// Depth is a wrapper around g_type_depth().
func (t Type) Depth() uint {
	return uint(C.g_type_depth(C.GType(t)))
}

// Parent is a wrapper around g_type_parent().
func (t Type) Parent() Type {
	return Type(C.g_type_parent(C.GType(t)))
}

// IsA is a wrapper around g_type_is_a().
func (t Type) IsA(isAType Type) bool {
	return gobool(C.g_type_is_a(C.GType(t), C.GType(isAType)))
}

// TypeFromName is a wrapper around g_type_from_name().
func TypeFromName(typeName string) Type {
	cstr := (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(cstr))

	return Type(C.g_type_from_name(cstr))
}

// TypeNextBase is a wrapper around g_type_next_base.
func TypeNextBase(leafType, rootType Type) Type {
	return Type(C.g_type_next_base(C.GType(leafType), C.GType(rootType)))
}

// goMarshal is called by the GLib runtime when a closure needs to be invoked.
// The closure will be invoked with as many arguments as it can take, from 0 to
// the full amount provided by the call. If the closure asks for more parameters
// than there are to give, then a runtime panic will occur.
//
//export goMarshal
func goMarshal(
	gclosure *C.GClosure,
	retValue *C.GValue,
	nParams C.guint,
	params *C.GValue,
	invocationHint C.gpointer,
	gobject *C.GObject) {

	// Get the function value associated with this callback closure.
	fs := intern.ObjectClosure(unsafe.Pointer(gobject), unsafe.Pointer(gclosure))
	if fs == nil {
		// Possible data race, bail.
		log.Println("warning:",
			"gobject", unsafe.Pointer(gobject),
			"gclosure", unsafe.Pointer(gclosure), "not found")
		return
	}

	// Fast path for an empty function.
	if fn, ok := fs.Func.Interface().(func()); ok {
		fn()
		return
	}

	fsType := fs.Func.Type()

	// Get number of parameters passed in.
	nGLibParams := int(nParams)
	nTotalParams := nGLibParams

	// Reflect may panic, so we defer recover here to re-panic with our trace.
	defer fs.TryRepanic()

	// Get number of parameters from the callback closure. If this exceeds
	// the total number of marshaled parameters, trigger a runtime panic.
	nCbParams := fsType.NumIn()
	if nCbParams > nTotalParams {
		fs.Panicf("too many closure args: have %d, max %d", nCbParams, nTotalParams)
	}

	// Create a slice of reflect.Values as arguments to call the function.
	gValues := unsafe.Slice(params, nCbParams)
	args := make([]reflect.Value, 0, nCbParams)

	// Fill beginning of args, up to the minimum of the total number of callback
	// parameters and parameters from the glib runtime.
	for i := 0; i < nCbParams && i < nGLibParams; i++ {
		v := (*Value)(unsafe.Pointer(&gValues[i]))
		val := v.GoValue()

		// Parameters that are descendants of GObject come wrapped in another
		// GObject. For C applications, the default marshaller
		// (g_cclosure_marshal_VOID__VOID in gmarshal.c in the GTK glib library)
		// 'peeks' into the enclosing object and passes the wrapped object to
		// the handler. Use the Object.Cast() function to emulate that for Go
		// signal handlers.
		switch v := val.(type) {
		case *Object:
			val = v.Cast()
		case *Variant:
			if g := v.GoValue(); g != nil {
				val = g
			}
		}

		args = append(args, reflect.ValueOf(val).Convert(fsType.In(i)))
	}

	// Call closure with args. If the callback returns one or more values, save
	// the GValue equivalent of the first.
	rv := fs.Func.Call(args)
	if retValue != nil && len(rv) > 0 {
		gv := NewValue(rv[0].Interface())
		ok := C.g_value_transform(gv.native(), retValue) != 0

		if !ok {
			fs.Panicf(
				"failed to transform return value from %s to %s",
				gv.Type(), (*Value)(unsafe.Pointer(retValue)).Type())
		}
	}
}

// Priority is the enumerated type for GLib priority event sources.
type Priority int

const (
	PriorityHigh        Priority = C.G_PRIORITY_HIGH
	PriorityDefault     Priority = C.G_PRIORITY_DEFAULT // TimeoutAdd
	PriorityHighIdle    Priority = C.G_PRIORITY_HIGH_IDLE
	PriorityDefaultIdle Priority = C.G_PRIORITY_DEFAULT_IDLE // IdleAdd
	PriorityLow         Priority = C.G_PRIORITY_LOW
)

type SourceHandle uint

// sourceFunc is the callback for g_idle_add_full and g_timeout_add_full that
// replaces the GClosure API.
//
//export sourceFunc
func sourceFunc(data C.gpointer) C.gboolean {
	v := gbox.Get(uintptr(data))
	fs := v.(*closure.FuncStack)

	rv := fs.Func.Call(nil)
	if len(rv) == 1 && rv[0].Bool() {
		return C.TRUE
	}

	return C.FALSE
}

//export removeSourceFunc
func removeSourceFunc(data C.gpointer) {
	gbox.Delete(uintptr(data))
}

var (
	_sourceFunc       = (*[0]byte)(C.sourceFunc)
	_removeSourceFunc = (*[0]byte)(C.removeSourceFunc)
)

// IdleAdd adds an idle source to the default main event loop context with the
// DefaultIdle priority. If f is not a function with no parameter, then IdleAdd
// will panic.
//
// After running once, the source func will be removed from the main event loop,
// unless f returns a single bool true.
func IdleAdd(f interface{}) SourceHandle {
	return idleAdd(PriorityDefaultIdle, f)
}

// IdleAddPriority adds an idle source to the default main event loop context
// with the given priority. Its behavior is the same as IdleAdd.
func IdleAddPriority(priority Priority, f interface{}) SourceHandle {
	return idleAdd(priority, f)
}

func idleAdd(priority Priority, f interface{}) SourceHandle {
	fs := closure.NewIdleFuncStack(f, 2)
	id := C.gpointer(gbox.Assign(fs))
	h := C.g_idle_add_full(C.gint(priority), _sourceFunc, id, _removeSourceFunc)

	return SourceHandle(h)
}

// TimeoutAdd adds an timeout source to the default main event loop context.
// Timeout is in milliseconds. If f is not a function with no parameter, then it
// will panic.
//
// After running once, the source func will be removed from the main event loop,
// unless f returns a single bool true.
func TimeoutAdd(milliseconds uint, f interface{}) SourceHandle {
	return timeoutAdd(milliseconds, false, PriorityDefault, f)
}

// TimeoutAddPriority is similar to TimeoutAdd with the given priority. Refer to
// TimeoutAdd for more information.
func TimeoutAddPriority(milliseconds uint, priority Priority, f interface{}) SourceHandle {
	return timeoutAdd(milliseconds, false, priority, f)
}

// TimeoutSecondsAdd is similar to TimeoutAdd, except with seconds granularity.
func TimeoutSecondsAdd(seconds uint, f interface{}) SourceHandle {
	return timeoutAdd(seconds, true, PriorityDefault, f)
}

// TimeoutSecondsAddPriority adds a timeout source with the given priority.
// Refer to TimeoutSecondsAdd for more information.
func TimeoutSecondsAddPriority(seconds uint, priority Priority, f interface{}) SourceHandle {
	return timeoutAdd(seconds, true, priority, f)
}

func timeoutAdd(time uint, sec bool, priority Priority, f interface{}) SourceHandle {
	fs := closure.NewIdleFuncStack(f, 2)
	id := C.gpointer(gbox.Assign(fs))

	var h C.guint
	if sec {
		h = C.g_timeout_add_seconds_full(C.gint(priority), C.guint(time), _sourceFunc, id, _removeSourceFunc)
	} else {
		h = C.g_timeout_add_full(C.gint(priority), C.guint(time), _sourceFunc, id, _removeSourceFunc)
	}

	return SourceHandle(h)
}

// SourceRemove is a wrapper around g_source_remove()
func SourceRemove(src SourceHandle) bool {
	return gobool(C.g_source_remove(C.guint(src)))
}

// Objector is an interface that describes the Object type's methods. Only this
// package's Object types and types that embed it can satisfy this interface.
type Objector interface {
	Connect(string, interface{}) SignalHandle
	ConnectAfter(string, interface{}) SignalHandle

	HandlerBlock(SignalHandle)
	HandlerUnblock(SignalHandle)
	HandlerDisconnect(SignalHandle)

	ObjectProperty(string) interface{}
	SetObjectProperty(string, interface{})

	Cast() Objector
	Native() uintptr

	baseObject() *Object
}

var _ Objector = (*Object)(nil)

// Object is a representation of GLib's GObject.
type Object struct {
	*objectNative
	box *intern.Box
}

// InternObject gets the internal Object type. This is used for calling methods
// not in the Objector.
func InternObject(obj Objector) *Object {
	return obj.baseObject()
}

// Take wraps a unsafe.Pointer as a Object, taking ownership of it.
// This function is exported for visibility in other gotk3 packages and
// is not meant to be used by applications.
//
// To be clear, this should mostly be used when Gtk says "transfer none". Refer
// to AssumeOwnership for more details.
func Take(ptr unsafe.Pointer) *Object {
	obj := newObject(ptr)
	if obj == nil {
		return nil
	}

	// Ensure that the reference is sunken.
	obj.RefSink()
	defer obj.Unref()

	obj.addToggleRef()

	return obj
}

// AssumeOwnership is similar to Take, except the function does not take a
// reference. This is usually used for newly constructed objects that for some
// reason does not have an initial floating reference.
//
// To be clear, this should often be used when Gtk says "transfer full", as it
// means the ownership is transferred to us (Go), so we can assume that much.
// This is in contrary to Take, which is used when Gtk says "transfer none", as
// we're now referencing an object that might possibly be kept by C, so we
// should take our own.
func AssumeOwnership(ptr unsafe.Pointer) *Object {
	obj := newObject(ptr)
	if obj == nil {
		return nil
	}

	obj.RefSink()
	defer obj.Unref()

	obj.addToggleRef()
	obj.Unref()

	return obj
}

//export goToggleNotify
func goToggleNotify(_ C.gpointer, obj *C.GObject, isLastInt C.gboolean) {
	intern.Toggle(unsafe.Pointer(obj), isLastInt != 0)
}

// objectNative wraps around a native C object. It exists to work around
// runtime.SetFinalizer's cyclic restriction.
type objectNative struct {
	GObject *C.GObject
}

// newObject creates a new Object from a GObject pointer with the finalizer set.
func newObject(ptr unsafe.Pointer) *Object {
	if ptr == nil {
		return nil
	}

	native := &objectNative{GObject: (*C.GObject)(ptr)}
	native.attachFinalizer()

	return &Object{
		objectNative: native,
		box:          intern.ObjectBox(ptr),
	}
}

func (native *objectNative) addToggleRef() {
	C.g_object_add_toggle_ref(native.GObject, (*[0]byte)(C.goToggleNotify), nil)
}

func (native *objectNative) removeToggleRef() {
	C.g_object_remove_toggle_ref(native.GObject, (*[0]byte)(C.goToggleNotify), nil)
}

func (native *objectNative) attachFinalizer() {
	runtime.SetFinalizer(native, finalizeObjectNative)
}

func finalizeObjectNative(native *objectNative) {
	if !intern.ShouldFree(unsafe.Pointer(native.GObject)) {
		// Delegate finalizing to the next cycle.
		native.attachFinalizer()
		return
	}

	native.removeToggleRef()
}

// Cast casts v to the concrete Go type (e.g. *Object to *gtk.Entry).
func (v *Object) Cast() Objector {
	var gvalue C.GValue
	C.g_value_init_from_instance(&gvalue, C.gpointer(v.Native()))

	value := ValueFromNative(unsafe.Pointer(&gvalue))
	defer value.Unset()

	// Note that if GoValue successfully unmarshaled into a nil object type,
	// then the interface would actually be non-nil.
	if gv := value.GoValue(); gv != nil && gv != InvalidValue {
		return gv.(Objector)
	}

	return v
}

func (v *Object) baseObject() *Object {
	return v
}

// native returns a pointer to the underlying GObject.
func (v *Object) native() *C.GObject {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GObject)(p)
}

// Native returns a pointer to the underlying GObject.
func (v *Object) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// IsA is a wrapper around g_type_is_a().
func (v *Object) IsA(typ Type) bool {
	return gobool(C.g_type_is_a(C.GType(v.TypeFromInstance()), C.GType(typ)))
}

// TypeFromInstance is a wrapper around g_type_from_instance().
func (v *Object) TypeFromInstance() Type {
	c := C._g_type_from_instance(C.gpointer(unsafe.Pointer(v.native())))
	return Type(c)
}

// // ToGObject type converts an unsafe.Pointer as a native C GObject.
// // This function is exported for visibility in other gotk3 packages and
// // is not meant to be used by applications.
// func ToGObject(p unsafe.Pointer) *C.GObject {
// 	return (*C.GObject)(p)
// }

// Ref is a wrapper around g_object_ref().
func (v *Object) Ref() {
	C.g_object_ref(C.gpointer(v.GObject))
}

// Unref is a wrapper around g_object_unref().
func (v *Object) Unref() {
	C.g_object_unref(C.gpointer(v.GObject))
}

// RefSink is a wrapper around g_object_ref_sink().
func (v *Object) RefSink() {
	C.g_object_ref_sink(C.gpointer(v.GObject))
}

// IsFloating is a wrapper around g_object_is_floating().
func (v *Object) IsFloating() bool {
	c := C.g_object_is_floating(C.gpointer(v.GObject))
	return gobool(c)
}

// ForceFloating is a wrapper around g_object_force_floating().
func (v *Object) ForceFloating() {
	C.g_object_force_floating(v.GObject)
}

// StopEmission is a wrapper around g_signal_stop_emission_by_name().
func (v *Object) StopEmission(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.g_signal_stop_emission_by_name((C.gpointer)(v.GObject),
		(*C.gchar)(cstr))
}

// PropertyType returns the Type of a property of the underlying GObject. If the
// property is missing, it will return TypeInvalid.
func (v *Object) PropertyType(name string) Type {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	return v.propertyType(cstr)
}

func (v *Object) propertyType(cstr *C.gchar) Type {
	paramSpec := C.g_object_class_find_property(C._g_object_get_class(v.native()), (*C.gchar)(cstr))
	if paramSpec == nil {
		return TypeInvalid
	}

	return Type(paramSpec.value_type)
}

// ObjectProperty is a wrapper around g_object_get_property(). If the property's
// type cannot be resolved to a Go type, then InvalidValue is returned.
func (v *Object) ObjectProperty(name string) interface{} {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	t := v.propertyType((*C.gchar)(cstr))
	if t == TypeInvalid {
		return InvalidValue
	}

	p := InitValue(t)
	C.g_object_get_property(v.GObject, (*C.gchar)(cstr), p.native())
	return p.GoValue()
}

// SetObjectProperty is a wrapper around g_object_set_property().
func (v *Object) SetObjectProperty(name string, value interface{}) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	p := NewValue(value)
	C.g_object_set_property(v.GObject, (*C.gchar)(cstr), p.native())
}

// Emit emits the signal specified by the string s to an Object. Arguments to
// callback functions connected to this signal must be specified in args. Emit
// returns an interface{} which must be type asserted to an equivalent Go type.
// The return value is the return value for the native C callback.
//
// Note that this code is unsafe in that the types of values in args are not
// checked against whether they are suitable for the callback.
func (v *Object) Emit(s string, args ...interface{}) interface{} {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))

	// Create array of this instance and arguments
	valv := (*C.GValue)(C.malloc(C.sizeof_GValue * C.ulong(len(args)+1)))
	defer C.free(unsafe.Pointer(valv))

	// Add args and valv
	val := NewValue(v)
	C.val_list_insert(valv, C.int(0), val.native())

	for i := range args {
		val := NewValue(args[i])
		C.val_list_insert(valv, C.int(i+1), val.native())
	}

	t := v.TypeFromInstance()
	// TODO: use just the signal name
	id := C.g_signal_lookup((*C.gchar)(cstr), C.GType(t))

	ret := AllocateValue()
	C.g_signal_emitv(valv, id, C.GQuark(0), ret.native())

	return ret.GoValue()
}

// HandlerBlock is a wrapper around g_signal_handler_block().
func (v *Object) HandlerBlock(handle SignalHandle) {
	C.g_signal_handler_block(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerUnblock is a wrapper around g_signal_handler_unblock().
func (v *Object) HandlerUnblock(handle SignalHandle) {
	C.g_signal_handler_unblock(C.gpointer(v.GObject), C.gulong(handle))
}

// HandlerDisconnect is a wrapper around g_signal_handler_disconnect().
func (v *Object) HandlerDisconnect(handle SignalHandle) {
	// Ensure that Gtk will not use the closure beforehand.
	C.g_signal_handler_disconnect(C.gpointer(v.GObject), C.gulong(handle))
}

// InitiallyUnowned is a representation of GLib's GInitiallyUnowned.
type InitiallyUnowned struct {
	*Object
}

type invalidValueType struct{}

var (
	// InvalidValue is returned from Value methods, such as GoValue, to indicate
	// that the value obtained is invalid.
	InvalidValue = invalidValueType{}
)

// Value is a representation of GLib's GValue.
type Value struct {
	gvalue C.GValue
}

func marshalValue(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Value)(unsafe.Pointer(c)), nil
}

// AllocateValue allocates a Value but does not initialize it. It sets a
// runtime finalizer to call g_value_unset() on the underlying GValue after
// leaving scope. ValueAlloc() returns a non-nil error if the allocation failed.
func AllocateValue() *Value {
	v := new(Value)

	//An allocated GValue is not guaranteed to hold a value that can be unset
	//We need to double check before unsetting, to prevent:
	//`g_value_unset: assertion 'G_IS_VALUE (value)' failed`
	runtime.SetFinalizer(v, func(f *Value) {
		if f.IsValue() {
			f.Unset()
		}
	})

	return v
}

// InitValue is a wrapper around g_value_init() and allocates and initializes a
// new Value with the Type t. A runtime finalizer is set to call g_value_unset()
// on the underlying GValue after leaving scope.  ValueInit() returns a non-nil
// error if the allocation failed.
func InitValue(t Type) *Value {
	v := new(Value)
	C.g_value_init(v.native(), C.GType(t))

	runtime.SetFinalizer(v, (*Value).Unset)
	return v
}

// ValueFromNative returns a type-asserted pointer to the Value.
func ValueFromNative(l unsafe.Pointer) *Value {
	//TODO why it does not add finalizer to the value?
	return (*Value)(l)
}

// NewValue converts a Go type to a comparable GValue. It will panic if the
// given type is unknown. Most Go primitive types and all Object types are
// supported.
func NewValue(v interface{}) *Value {
	if v == nil {
		val := InitValue(TypePointer)
		val.SetPointer(uintptr(unsafe.Pointer(nil)))
		return val
	}

	if val := newValuePrimitive(v); val != nil {
		return val
	}

	if v == InvalidValue {
		return InitValue(TypeInvalid)
	}

	// Try this since above doesn't catch constants under other types.
	rval := reflect.Indirect(reflect.ValueOf(v))

	switch rval.Kind() {
	case reflect.Bool:
		return newValuePrimitive(rval.Bool())
	case reflect.Int8:
		return newValuePrimitive(int8(rval.Int()))
	case reflect.Int32:
		return newValuePrimitive(int32(rval.Int()))
	case reflect.Int64:
		return newValuePrimitive(int64(rval.Int()))
	case reflect.Int:
		return newValuePrimitive(int(rval.Int()))
	case reflect.Uint8:
		return newValuePrimitive(uint8(rval.Uint()))
	case reflect.Uint32:
		return newValuePrimitive(uint32(rval.Uint()))
	case reflect.Uint64:
		return newValuePrimitive(uint64(rval.Uint()))
	case reflect.Uint:
		return newValuePrimitive(uint(rval.Uint()))
	case reflect.Float32:
		return newValuePrimitive(float32(rval.Float()))
	case reflect.Float64:
		return newValuePrimitive(float64(rval.Float()))
	case reflect.String:
		return newValuePrimitive(rval.String())
	}

	log.Panicf("type %T not implemented", v)
	return nil
}

func newValuePrimitive(v interface{}) *Value {
	var val *Value
	switch e := v.(type) {
	case *Value:
		val = e
	case bool:
		val = InitValue(TypeBoolean)
		val.SetBool(e)
	case int8:
		val = InitValue(TypeChar)
		val.SetSchar(e)
	case int32:
		val = InitValue(TypeLong)
		val.SetLong(e)
	case int64:
		val = InitValue(TypeInt64)
		val.SetInt64(e)
	case int:
		val = InitValue(TypeInt)
		val.SetInt(e)
	case uint8:
		val = InitValue(TypeUchar)
		val.SetUchar(e)
	case uint32:
		val = InitValue(TypeUlong)
		val.SetUlong(e)
	case uint64:
		val = InitValue(TypeUint64)
		val.SetUint64(e)
	case uint:
		val = InitValue(TypeUint)
		val.SetUint(e)
	case float32:
		val = InitValue(TypeFloat)
		val.SetFloat(e)
	case float64:
		val = InitValue(TypeDouble)
		val.SetDouble(e)
	case string:
		val = InitValue(TypeString)
		val.SetString(e)
	case Objector:
		val = InitValue(TypeObject)
		val.SetInstance(uintptr(unsafe.Pointer(e.Native())))
	}
	return val
}

func (v *Value) native() *C.GValue {
	if v == nil {
		return nil
	}
	return &v.gvalue
}

// Native returns a pointer to the underlying GValue.
func (v *Value) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

// IsValue checks if value is a valid and initialized GValue structure.
func (v *Value) IsValue() bool {
	return gobool(C._g_is_value(v.native()))
}

// TypeName gets the type name of value.
func (v *Value) TypeName() string {
	return C.GoString((*C.char)(C._g_value_type_name(v.native())))
}

// Unset is wrapper for g_value_unset
func (v *Value) Unset() {
	C.g_value_unset(v.native())
}

// Type returns the Value's actual type. is a wrapper around the G_VALUE_TYPE()
// macro. It returns TYPE_INVALID if v does not hold a Type, or otherwise
// returns the Type of v.
//
// To get the fundamental type, use FundamentalType.
func (v *Value) Type() (actual Type) {
	if v == nil || !v.IsValue() {
		return TypeInvalid
	}
	return Type(C._g_value_type(v.native()))
}

// CastObject casts the given object pointer to the Go concrete type. The caller
// is responsible for recasting the interface to the wanted type.
//
// Deprecated: Use obj.Cast() instead.
func CastObject(obj *Object) interface{} {
	return obj.Cast()
}

// GValueMarshaler is a marshal function to convert a GValue into an
// appropriate Go type.  The uintptr parameter is a *C.GValue.
type GValueMarshaler func(uintptr) (interface{}, error)

// TypeMarshaler represents an actual type and it's associated marshaler.
type TypeMarshaler struct {
	T Type
	F GValueMarshaler
}

// RegisterGValueMarshalers adds marshalers for several types to the
// internal marshalers map. Once registered, calling GoValue on any
// Value with a registered type will return the data returned by the
// marshaler.
func RegisterGValueMarshalers(tm []TypeMarshaler) {
	gValueMarshalers.register(tm)
}

type marshalMap map[Type]GValueMarshaler

// gValueMarshalers is a map of Glib types to functions to marshal a
// GValue to a native Go type.
var gValueMarshalers = marshalMap{
	TypeInvalid:   marshalInvalid,
	TypeNone:      marshalNone,
	TypeInterface: marshalInterface,
	TypeChar:      marshalChar,
	TypeUchar:     marshalUchar,
	TypeBoolean:   marshalBoolean,
	TypeInt:       marshalInt,
	TypeLong:      marshalLong,
	TypeEnum:      marshalEnum,
	TypeInt64:     marshalInt64,
	TypeUint:      marshalUint,
	TypeUlong:     marshalUlong,
	TypeFlags:     marshalFlags,
	TypeUint64:    marshalUint64,
	TypeFloat:     marshalFloat,
	TypeDouble:    marshalDouble,
	TypeString:    marshalString,
	TypePointer:   marshalPointer,
	TypeBoxed:     marshalBoxed,
	TypeObject:    marshalObject,
	TypeVariant:   marshalVariant,
}

func init() {
	gValueMarshalers.register([]TypeMarshaler{
		{Type(C.g_value_get_type()), marshalValue},
	})
}

func (m marshalMap) register(tm []TypeMarshaler) {
	for i := range tm {
		m[tm[i].T] = tm[i].F
	}
}

func (m marshalMap) lookup(v *Value) GValueMarshaler {
	actual := v.Type()
	if f, ok := m[actual]; ok {
		return f
	}

	fundamental := FundamentalType(actual)
	if f, ok := m[fundamental]; ok {
		return f
	}

	return nil
}

func (m marshalMap) lookupType(t Type) GValueMarshaler {
	if f, ok := m[t]; ok {
		return f
	}
	return nil
}

func marshalInvalid(uintptr) (interface{}, error) {
	return nil, errors.New("invalid type")
}

func marshalNone(uintptr) (interface{}, error) {
	return nil, nil
}

func marshalInterface(uintptr) (interface{}, error) {
	return nil, errors.New("interface conversion not yet implemented")
}

func marshalChar(p uintptr) (interface{}, error) {
	c := C.g_value_get_schar((*C.GValue)(unsafe.Pointer(p)))
	return int8(c), nil
}

func marshalUchar(p uintptr) (interface{}, error) {
	c := C.g_value_get_uchar((*C.GValue)(unsafe.Pointer(p)))
	return uint8(c), nil
}

func marshalBoolean(p uintptr) (interface{}, error) {
	c := C.g_value_get_boolean((*C.GValue)(unsafe.Pointer(p)))
	return gobool(c), nil
}

func marshalInt(p uintptr) (interface{}, error) {
	c := C.g_value_get_int((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalLong(p uintptr) (interface{}, error) {
	c := C.g_value_get_long((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalEnum(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return int(c), nil
}

func marshalInt64(p uintptr) (interface{}, error) {
	c := C.g_value_get_int64((*C.GValue)(unsafe.Pointer(p)))
	return int64(c), nil
}

func marshalUint(p uintptr) (interface{}, error) {
	c := C.g_value_get_uint((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalUlong(p uintptr) (interface{}, error) {
	c := C.g_value_get_ulong((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))
	return uint(c), nil
}

func marshalUint64(p uintptr) (interface{}, error) {
	c := C.g_value_get_uint64((*C.GValue)(unsafe.Pointer(p)))
	return uint64(c), nil
}

func marshalFloat(p uintptr) (interface{}, error) {
	c := C.g_value_get_float((*C.GValue)(unsafe.Pointer(p)))
	return float32(c), nil
}

func marshalDouble(p uintptr) (interface{}, error) {
	c := C.g_value_get_double((*C.GValue)(unsafe.Pointer(p)))
	return float64(c), nil
}

func marshalString(p uintptr) (interface{}, error) {
	c := C.g_value_get_string((*C.GValue)(unsafe.Pointer(p)))
	return C.GoString((*C.char)(c)), nil
}

func marshalBoxed(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return unsafe.Pointer(c), nil
}

func marshalPointer(p uintptr) (interface{}, error) {
	c := C.g_value_get_pointer((*C.GValue)(unsafe.Pointer(p)))
	return unsafe.Pointer(c), nil
}

func marshalObject(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return Take(unsafe.Pointer(c)), nil
}

func marshalVariant(p uintptr) (interface{}, error) {
	c := C.g_value_get_variant((*C.GValue)(unsafe.Pointer(p)))
	return newVariant((*C.GVariant)(c)), nil
}

// GoValue converts a Value to comparable Go type. GoValue() returns
// InvalidValue if the conversion was unsuccessful. The returned value must be
// type asserted to the actual Go type.
//
// Unlike the type getters, although this method is not type-safe, it can get
// any concrete object type out, as long as there exists a marshaler for it.
func (v *Value) GoValue() interface{} {
	f := gValueMarshalers.lookup(v)
	if f == nil {
		return InvalidValue
	}

	// No need to add finalizer because it is already done by AllocateValue and
	// InitValue. (?)
	g, err := f(uintptr(unsafe.Pointer(v.native())))
	if err != nil {
		return InvalidValue
	}

	return g
}

// SetBool is a wrapper around g_value_set_boolean().
func (v *Value) SetBool(val bool) {
	C.g_value_set_boolean(v.native(), gbool(val))
}

// SetSChar is a wrapper around g_value_set_schar().
func (v *Value) SetSchar(val int8) {
	C.g_value_set_schar(v.native(), C.gint8(val))
}

// SetInt64 is a wrapper around g_value_set_int64().
func (v *Value) SetInt64(val int64) {
	C.g_value_set_int64(v.native(), C.gint64(val))
}

// SetLong is a wrapper around g_value_set_long().
func (v *Value) SetLong(long int32) {
	C.g_value_set_long(v.native(), C.glong(long))
}

// SetInt is a wrapper around g_value_set_int().
func (v *Value) SetInt(val int) {
	C.g_value_set_int(v.native(), C.gint(val))
}

// SetUchar is a wrapper around g_value_set_uchar().
func (v *Value) SetUchar(val uint8) {
	C.g_value_set_uchar(v.native(), C.guchar(val))
}

// SetUint64 is a wrapper around g_value_set_uint64().
func (v *Value) SetUint64(val uint64) {
	C.g_value_set_uint64(v.native(), C.guint64(val))
}

// SetUlong is a wrapper around g_value_set_ulong().
func (v *Value) SetUlong(ulong uint32) {
	C.g_value_set_ulong(v.native(), C.gulong(ulong))
}

// SetUint is a wrapper around g_value_set_uint().
func (v *Value) SetUint(val uint) {
	C.g_value_set_uint(v.native(), C.guint(val))
}

// SetFloat is a wrapper around g_value_set_float().
func (v *Value) SetFloat(val float32) {
	C.g_value_set_float(v.native(), C.gfloat(val))
}

// SetDouble is a wrapper around g_value_set_double().
func (v *Value) SetDouble(val float64) {
	C.g_value_set_double(v.native(), C.gdouble(val))
}

// SetString is a wrapper around g_value_set_string().
func (v *Value) SetString(val string) {
	cstr := C.CString(val)
	defer C.free(unsafe.Pointer(cstr))
	C.g_value_set_string(v.native(), (*C.gchar)(cstr))
}

// SetInstance is a wrapper around g_value_set_instance().
func (v *Value) SetInstance(instance uintptr) {
	C.g_value_set_instance(v.native(), C.gpointer(instance))
}

// SetPointer is a wrapper around g_value_set_pointer().
func (v *Value) SetPointer(p uintptr) {
	C.g_value_set_pointer(v.native(), C.gpointer(p))
}

// Pointer is a wrapper around g_value_get_pointer().
func (v *Value) Pointer() unsafe.Pointer {
	return unsafe.Pointer(C.g_value_get_pointer(v.native()))
}

// String is a wrapper around g_value_get_string().  String() returns a non-nil
// error if g_value_get_string() returned a NULL pointer to distinguish between
// returning a NULL pointer and returning an empty string.
func (v *Value) String() string {
	c := C.g_value_get_string(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}
