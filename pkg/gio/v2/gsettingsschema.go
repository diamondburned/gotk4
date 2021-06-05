// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_settings_schema_get_type()), F: marshalSettingsSchema},
		{T: externglib.Type(C.g_settings_schema_key_get_type()), F: marshalSettingsSchemaKey},
		{T: externglib.Type(C.g_settings_schema_source_get_type()), F: marshalSettingsSchemaSource},
	})
}

// SettingsSchemaSourceGetDefault gets the default system schema source.
//
// This function is not required for normal uses of #GSettings but it may be
// useful to authors of plugin management systems or to those who want to
// introspect the content of schemas.
//
// If no schemas are installed, nil will be returned.
//
// The returned source may actually consist of multiple schema sources from
// different directories, depending on which directories were given in
// `XDG_DATA_DIRS` and `GSETTINGS_SCHEMA_DIR`. For this reason, all lookups
// performed against the default source should probably be done recursively.
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	var cret *C.GSettingsSchemaSource
	var goret1 *SettingsSchemaSource

	cret = C.g_settings_schema_source_get_default()

	goret1 = WrapSettingsSchemaSource(unsafe.Pointer(cret))

	return goret1
}

// SettingsSchema: the SchemaSource and Schema APIs provide a mechanism for
// advanced control over the loading of schemas and a mechanism for
// introspecting their content.
//
// Plugin loading systems that wish to provide plugins a way to access settings
// face the problem of how to make the schemas for these settings visible to
// GSettings. Typically, a plugin will want to ship the schema along with itself
// and it won't be installed into the standard system directories for schemas.
//
// SchemaSource provides a mechanism for dealing with this by allowing the
// creation of a new 'schema source' from which schemas can be acquired. This
// schema source can then become part of the metadata associated with the plugin
// and queried whenever the plugin requires access to some settings.
//
// Consider the following example:
//
//    {
//      GSettings *settings;
//      gint some_value;
//
//      settings = plugin_get_settings (self, NULL);
//      some_value = g_settings_get_int (settings, "some-value");
//      ...
//    }
//
// It's also possible that the plugin system expects the schema source files
// (ie: .gschema.xml files) instead of a gschemas.compiled file. In that case,
// the plugin loading system must compile the schemas for itself before
// attempting to create the settings source.
type SettingsSchema struct {
	native C.GSettingsSchema
}

// WrapSettingsSchema wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSettingsSchema(ptr unsafe.Pointer) *SettingsSchema {
	if ptr == nil {
		return nil
	}

	return (*SettingsSchema)(ptr)
}

func marshalSettingsSchema(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSettingsSchema(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SettingsSchema) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// ID: get the ID of @schema.
func (s *SettingsSchema) ID() string {
	var arg0 *C.GSettingsSchema

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_settings_schema_get_id(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// Key gets the key named @name from @schema.
//
// It is a programmer error to request a key that does not exist. See
// g_settings_schema_list_keys().
func (s *SettingsSchema) Key(name string) *SettingsSchemaKey {
	var arg0 *C.GSettingsSchema
	var arg1 *C.gchar

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GSettingsSchemaKey
	var goret1 *SettingsSchemaKey

	cret = C.g_settings_schema_get_key(arg0, name)

	goret1 = WrapSettingsSchemaKey(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *SettingsSchemaKey) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Path gets the path associated with @schema, or nil.
//
// Schemas may be single-instance or relocatable. Single-instance schemas
// correspond to exactly one set of keys in the backend database: those located
// at the path returned by this function.
//
// Relocatable schemas can be referenced by other schemas and can therefore
// describe multiple sets of keys at different locations. For relocatable
// schemas, this function will return nil.
func (s *SettingsSchema) Path() string {
	var arg0 *C.GSettingsSchema

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_settings_schema_get_path(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// HasKey checks if @schema has a key named @name.
func (s *SettingsSchema) HasKey(name string) bool {
	var arg0 *C.GSettingsSchema
	var arg1 *C.gchar

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_settings_schema_has_key(arg0, name)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// ListChildren gets the list of children in @schema.
//
// You should free the return value with g_strfreev() when you are done with it.
func (s *SettingsSchema) ListChildren() []string {
	var arg0 *C.GSettingsSchema

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))

	var cret **C.gchar
	var goret1 []string

	cret = C.g_settings_schema_list_children(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			goret1[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}

	return goret1
}

// ListKeys introspects the list of keys on @schema.
//
// You should probably not be calling this function from "normal" code (since
// you should already know what keys are in your schema). This function is
// intended for introspection reasons.
func (s *SettingsSchema) ListKeys() []string {
	var arg0 *C.GSettingsSchema

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))

	var cret **C.gchar
	var goret1 []string

	cret = C.g_settings_schema_list_keys(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			goret1[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}

	return goret1
}

// Ref: increase the reference count of @schema, returning a new reference.
func (s *SettingsSchema) Ref() *SettingsSchema {
	var arg0 *C.GSettingsSchema

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))

	var cret *C.GSettingsSchema
	var goret1 *SettingsSchema

	cret = C.g_settings_schema_ref(arg0)

	goret1 = WrapSettingsSchema(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *SettingsSchema) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Unref: decrease the reference count of @schema, possibly freeing it.
func (s *SettingsSchema) Unref() {
	var arg0 *C.GSettingsSchema

	arg0 = (*C.GSettingsSchema)(unsafe.Pointer(s.Native()))

	C.g_settings_schema_unref(arg0)
}

// SettingsSchemaKey is an opaque data structure and can only be accessed using
// the following functions.
type SettingsSchemaKey struct {
	native C.GSettingsSchemaKey
}

// WrapSettingsSchemaKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSettingsSchemaKey(ptr unsafe.Pointer) *SettingsSchemaKey {
	if ptr == nil {
		return nil
	}

	return (*SettingsSchemaKey)(ptr)
}

func marshalSettingsSchemaKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSettingsSchemaKey(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SettingsSchemaKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// DefaultValue gets the default value for @key.
//
// Note that this is the default value according to the schema. System
// administrator defaults and lockdown are not visible via this API.
func (k *SettingsSchemaKey) DefaultValue() *glib.Variant {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	var cret *C.GVariant
	var goret1 *glib.Variant

	cret = C.g_settings_schema_key_get_default_value(arg0)

	goret1 = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Description gets the description for @key.
//
// If no description has been provided in the schema for @key, returns nil.
//
// The description can be one sentence to several paragraphs in length.
// Paragraphs are delimited with a double newline. Descriptions can be
// translated and the value returned from this function is is the current
// locale.
//
// This function is slow. The summary and description information for the
// schemas is not stored in the compiled schema database so this function has to
// parse all of the source XML files in the schema directory.
func (k *SettingsSchemaKey) Description() string {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_settings_schema_key_get_description(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// Name gets the name of @key.
func (k *SettingsSchemaKey) Name() string {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_settings_schema_key_get_name(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// Range queries the range of a key.
//
// This function will return a #GVariant that fully describes the range of
// values that are valid for @key.
//
// The type of #GVariant returned is `(sv)`. The string describes the type of
// range restriction in effect. The type and meaning of the value contained in
// the variant depends on the string.
//
// If the string is `'type'` then the variant contains an empty array. The
// element type of that empty array is the expected type of value and all values
// of that type are valid.
//
// If the string is `'enum'` then the variant contains an array enumerating the
// possible values. Each item in the array is a possible valid value and no
// other values are valid.
//
// If the string is `'flags'` then the variant contains an array. Each item in
// the array is a value that may appear zero or one times in an array to be used
// as the value for this key. For example, if the variant contained the array
// `['x', 'y']` then the valid values for the key would be `[]`, `['x']`,
// `['y']`, `['x', 'y']` and `['y', 'x']`.
//
// Finally, if the string is `'range'` then the variant contains a pair of
// like-typed values -- the minimum and maximum permissible values for this key.
//
// This information should not be used by normal programs. It is considered to
// be a hint for introspection purposes. Normal programs should already know
// what is permitted by their own schema. The format may change in any way in
// the future -- but particularly, new forms may be added to the possibilities
// described above.
//
// You should free the returned value with g_variant_unref() when it is no
// longer needed.
func (k *SettingsSchemaKey) Range() *glib.Variant {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	var cret *C.GVariant
	var goret1 *glib.Variant

	cret = C.g_settings_schema_key_get_range(arg0)

	goret1 = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Summary gets the summary for @key.
//
// If no summary has been provided in the schema for @key, returns nil.
//
// The summary is a short description of the purpose of the key; usually one
// short sentence. Summaries can be translated and the value returned from this
// function is is the current locale.
//
// This function is slow. The summary and description information for the
// schemas is not stored in the compiled schema database so this function has to
// parse all of the source XML files in the schema directory.
func (k *SettingsSchemaKey) Summary() string {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_settings_schema_key_get_summary(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// ValueType gets the Type of @key.
func (k *SettingsSchemaKey) ValueType() *glib.VariantType {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	var cret *C.GVariantType
	var goret1 *glib.VariantType

	cret = C.g_settings_schema_key_get_value_type(arg0)

	goret1 = glib.WrapVariantType(unsafe.Pointer(cret))

	return goret1
}

// RangeCheck checks if the given @value is of the correct type and within the
// permitted range for @key.
//
// It is a programmer error if @value is not of the correct type -- you must
// check for this first.
func (k *SettingsSchemaKey) RangeCheck(value *glib.Variant) bool {
	var arg0 *C.GSettingsSchemaKey
	var arg1 *C.GVariant

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))
	arg1 = (*C.GVariant)(unsafe.Pointer(value.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_settings_schema_key_range_check(arg0, value)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Ref: increase the reference count of @key, returning a new reference.
func (k *SettingsSchemaKey) Ref() *SettingsSchemaKey {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	var cret *C.GSettingsSchemaKey
	var goret1 *SettingsSchemaKey

	cret = C.g_settings_schema_key_ref(arg0)

	goret1 = WrapSettingsSchemaKey(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *SettingsSchemaKey) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Unref: decrease the reference count of @key, possibly freeing it.
func (k *SettingsSchemaKey) Unref() {
	var arg0 *C.GSettingsSchemaKey

	arg0 = (*C.GSettingsSchemaKey)(unsafe.Pointer(k.Native()))

	C.g_settings_schema_key_unref(arg0)
}

// SettingsSchemaSource: this is an opaque structure type. You may not access it
// directly.
type SettingsSchemaSource struct {
	native C.GSettingsSchemaSource
}

// WrapSettingsSchemaSource wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSettingsSchemaSource(ptr unsafe.Pointer) *SettingsSchemaSource {
	if ptr == nil {
		return nil
	}

	return (*SettingsSchemaSource)(ptr)
}

func marshalSettingsSchemaSource(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSettingsSchemaSource(unsafe.Pointer(b)), nil
}

// NewSettingsSchemaSourceFromDirectory constructs a struct SettingsSchemaSource.
func NewSettingsSchemaSourceFromDirectory(directory string, parent *SettingsSchemaSource, trusted bool) (settingsSchemaSource *SettingsSchemaSource, err error) {
	var arg1 *C.gchar
	var arg2 *C.GSettingsSchemaSource
	var arg3 C.gboolean
	var errout *C.GError

	arg1 = (*C.gchar)(C.CString(directory))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GSettingsSchemaSource)(unsafe.Pointer(parent.Native()))
	if trusted {
		arg3 = C.gboolean(1)
	}

	var cret *C.GSettingsSchemaSource
	var goret1 *SettingsSchemaSource
	var goerr error

	cret = C.g_settings_schema_source_new_from_directory(directory, parent, trusted, &errout)

	goret1 = WrapSettingsSchemaSource(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *SettingsSchemaSource) {
		C.free(unsafe.Pointer(v.Native()))
	})
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// Native returns the underlying C source pointer.
func (s *SettingsSchemaSource) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// ListSchemas lists the schemas in a given source.
//
// If @recursive is true then include parent sources. If false then only include
// the schemas from one source (ie: one directory). You probably want true.
//
// Non-relocatable schemas are those for which you can call g_settings_new().
// Relocatable schemas are those for which you must use
// g_settings_new_with_path().
//
// Do not call this function from normal programs. This is designed for use by
// database editors, commandline tools, etc.
func (s *SettingsSchemaSource) ListSchemas(recursive bool) (nonRelocatable []string, relocatable []string) {
	var arg0 *C.GSettingsSchemaSource
	var arg1 C.gboolean

	arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(s.Native()))
	if recursive {
		arg1 = C.gboolean(1)
	}

	var arg2 ***C.gchar
	var ret2 []string
	var arg3 ***C.gchar
	var ret3 []string

	C.g_settings_schema_source_list_schemas(arg0, recursive, &arg2, &arg3)

	{
		var length int
		for p := arg2; *p != 0; p = (***C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret2 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (**C.gchar)(ptr.Add(unsafe.Pointer(arg2), i))
			ret2[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}
	{
		var length int
		for p := arg3; *p != 0; p = (***C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret3 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (**C.gchar)(ptr.Add(unsafe.Pointer(arg3), i))
			ret3[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}

	return ret2, ret3
}

// Lookup looks up a schema with the identifier @schema_id in @source.
//
// This function is not required for normal uses of #GSettings but it may be
// useful to authors of plugin management systems or to those who want to
// introspect the content of schemas.
//
// If the schema isn't found directly in @source and @recursive is true then the
// parent sources will also be checked.
//
// If the schema isn't found, nil is returned.
func (s *SettingsSchemaSource) Lookup(schemaID string, recursive bool) *SettingsSchema {
	var arg0 *C.GSettingsSchemaSource
	var arg1 *C.gchar
	var arg2 C.gboolean

	arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(schemaID))
	defer C.free(unsafe.Pointer(arg1))
	if recursive {
		arg2 = C.gboolean(1)
	}

	var cret *C.GSettingsSchema
	var goret1 *SettingsSchema

	cret = C.g_settings_schema_source_lookup(arg0, schemaID, recursive)

	goret1 = WrapSettingsSchema(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *SettingsSchema) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Ref: increase the reference count of @source, returning a new reference.
func (s *SettingsSchemaSource) Ref() *SettingsSchemaSource {
	var arg0 *C.GSettingsSchemaSource

	arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(s.Native()))

	var cret *C.GSettingsSchemaSource
	var goret1 *SettingsSchemaSource

	cret = C.g_settings_schema_source_ref(arg0)

	goret1 = WrapSettingsSchemaSource(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *SettingsSchemaSource) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Unref: decrease the reference count of @source, possibly freeing it.
func (s *SettingsSchemaSource) Unref() {
	var arg0 *C.GSettingsSchemaSource

	arg0 = (*C.GSettingsSchemaSource)(unsafe.Pointer(s.Native()))

	C.g_settings_schema_source_unref(arg0)
}