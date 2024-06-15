package types

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
)

// LinkMode describes the mode that determines how the generator generates
// code to be linked. See the possible constants for more information.
type LinkMode uint8

const (
	// RuntimeLinkMode is the new generation mode. It generates Go code that
	// uses the girepository library to link the library at runtime instead of
	// compile-time.
	//
	// Since all generated calls are actually Go function calls and not C
	// function calls, compile times are far better than when using
	// DynamicLinkMode, but runtime performance may be slightly lower and more
	// memory-intensive.
	//
	// Use RuntimeLinkMode for large packages, such as GTK; use DynamicLinkMode
	// for small and important packages, such as the core glib/v2 package.
	RuntimeLinkMode LinkMode = iota
	// DynamicLinkMode is the old generation mode. It generates Cgo code that
	// contains compile-time link flags to dynamically link libraries to the
	// output binary.
	//
	// This mode forces the use of cmd/cgo on every call during compilation,
	// making it very slow. Only use this on core packages such as glib.
	DynamicLinkMode
)

// GoAnyType generates the Go type signature for the AnyType union. An empty
// string returned is an invalid type. If pub is true, then the returned string
// will use public interface types for classes instead of implementation types.
func GoAnyType(gen FileGenerator, any gir.AnyType, pub bool) (string, bool) {
	switch {
	case any.Array != nil:
		return goArrayType(gen, *any.Array, pub)
	case any.Type != nil:
		return GoType(gen, *any.Type, pub)
	}

	// Probably varargs, ignore because CGo.
	return "", false
}

// goArrayType generates the Go type signature for the given array.
func goArrayType(gen FileGenerator, array gir.Array, pub bool) (string, bool) {
	if array.Type == nil {
		return "", false
	}

	arrayPrefix := "[]"
	if array.FixedSize > 0 {
		arrayPrefix = fmt.Sprintf("[%d]", array.FixedSize)
	}

	child, _ := GoType(gen, *array.Type, pub)
	// There can't be []void, so this check ensures there can only be valid
	// array types.
	if child == "" {
		return "", false
	}

	return arrayPrefix + child, true
}

// GoType is a convenient function that wraps around ResolveType and returns the
// Go type.
func GoType(gen FileGenerator, typ gir.Type, pub bool) (string, bool) {
	resolved := Resolve(gen, typ)
	if resolved == nil {
		return "", false
	}

	needsNamespace := resolved.NeedsNamespace(gen.Namespace())
	if pub {
		return resolved.PublicType(needsNamespace), true
	}
	return resolved.ImplType(needsNamespace), true
}

// RecordIsOpaque returns true if the record has no fields in the GIR schema.
// These records must always be referenced using a pointer.
func RecordIsOpaque(rec gir.Record) bool {
	return len(rec.Fields) == 0 || rec.GLibGetType == "intern" || rec.Foreign
}

// methodCanCallDirectly returns true if the method is generated, has no
// arguments (sans the receiver) and has no returns.
func methodCanCallDirectly(method *gir.Method) bool {
	return true &&
		method.CIdentifier != "" &&
		method.ShadowedBy == "" &&
		method.MovedTo == "" &&
		method.IsIntrospectable() &&
		method.Parameters != nil &&
		method.Parameters.InstanceParameter != nil &&
		method.Parameters.InstanceParameter.Type != nil &&
		len(method.Parameters.Parameters) == 0
}

// RecordHasFree returns the free/unref method if it has one.
func RecordHasFree(record *gir.Record) *gir.Method {
	for _, name := range []string{"unref", "free", "destroy"} {
		if m := FindMethodName(record.Methods, name); m != nil {
			return m
		}
	}
	return nil
}

// RecordPrintFree prints the call to the record's free function OR an empty
// string.
func RecordPrintFree(gen FileGenerator, parent *gir.TypeFindResult, value string) string {
	return RecordPrintFreeMethod(gen, parent, value)
}

// RecordPrintFreeMethod generates a call with 1 argument for either free or
// unref. If method is nil, then a C.free call is generated. Value is assumed to
// be an unsafe.Pointer.
//
// The caller must import girepository.Argument manually.
//
// Deprecated: Use RecordPrintFree.
func RecordPrintFreeMethod(gen FileGenerator, parent *gir.TypeFindResult, value string) string {
	rec := parent.Type.(*gir.Record)

	free := RecordHasFree(rec)
	if free == nil {
		return fmt.Sprintf("C.free(%s)", value)
	}

	// TODO: refactor thoughts: should typeconv and girgen/callable be combined?
	// typeconv has to generate function calling code in some cases for freeing,
	// and we have various different ways of doing that scattered throughout the
	// program. It would be far better to have 1 place of doing them all.

	switch gen.LinkMode() {
	case RuntimeLinkMode:
		p := pen.NewBlock()
		p.Linef("var args [1]girepository.Argument")
		p.Linef("*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(%s)", value)
		p.Linef("girepository.MustFind(%q, %q).InvokeRecordMethod(%q, args[:], nil)",
			parent.Namespace.Name, rec.Name, "free")
		return p.String()
	case DynamicLinkMode:
		return fmt.Sprintf(
			"C.%s((%s)(%s))",
			free.CIdentifier,
			AnyTypeCGo(free.Parameters.InstanceParameter.AnyType),
			value)
	default:
		panic("unreachable")
	}
}

// RecordHasUnref returns the unref method if it has one.
func RecordHasUnref(record *gir.Record) *gir.Method {
	// TODO: runtime link mode support
	return FindMethodName(record.Methods, "unref")
}

// RecordHasRef returns the ref method if it has one.
func RecordHasRef(record *gir.Record) *gir.Method {
	// TODO: runtime link mode support
	return FindMethodName(record.Methods, "ref")
}

// FindMethodName finds from the method the given name.
func FindMethodName(methods []gir.Method, name string) *gir.Method {
	for i, method := range methods {
		if method.Name == name && methodCanCallDirectly(&method) {
			return &methods[i]
		}
	}
	return nil
}

// EnsureNamespace ensures that exported, non-primitive types have the namespace
// prepended. This is useful for matching hard-coded types.
func EnsureNamespace(nsp *gir.NamespaceFindResult, girType string) string {
	if _, ok := girToBuiltin[girType]; ok {
		return girType
	}
	if _, ok := gpointerTypes[girType]; ok {
		return girType
	}

	// Special cases, because GIR is very unusual.
	switch girType {
	case "GType", "GValue":
		return girType
	}

	if strings.Contains(girType, ".") {
		return girType
	}

	return nsp.Namespace.Name + "." + girType
}

func countPtrs(typ gir.Type, result *gir.TypeFindResult) uint8 {
	return uint8(strings.Count(typ.CType, "*"))
}

var objectorMethods = map[string]struct{}{
	"Connect":           {},
	"ConnectAfter":      {},
	"HandlerBlock":      {},
	"HandlerUnblock":    {},
	"HandlerDisconnect": {},
	"NotifyProperty":    {},
	"ObjectProperty":    {},
	"SetObjectProperty": {},
	"FreezeNotify":      {},
	"ThawNotify":        {},
	"StopEmission":      {},
	"Cast":              {},
	"baseObject":        {},
}

func IsObjectorMethod(goName string) bool {
	_, is := objectorMethods[goName]
	return is
}

var (
	cTypePrefixes     = []string{"const", "volatile"}
	cTypePrefixEraser *strings.Replacer
)

func init() {
	replacers := make([]string, len(cTypePrefixes)*4)
	for _, prefix := range cTypePrefixes {
		// Use trailing spaces to prevent casting to the wrong type, like
		// gpointer instead of gconstpointer.
		replacers = append(replacers, prefix+" ", "")
		replacers = append(replacers, " "+prefix, "")
	}
	cTypePrefixEraser = strings.NewReplacer(replacers...)
}

// CountPtr counts the number of pointers in the given type string. If the
// string contains "[]", then the pointer is counted up to that.
func CountPtr(typ string) int {
	if typ == "" {
		return 0
	}

	sliceIx := strings.Index(typ, "[]")
	if sliceIx == -1 {
		sliceIx = len(typ)
	}

	return strings.Count(typ[:sliceIx], "*")
}

// DecPtr decrements a pointer in the type.
func DecPtr(t string) string {
	return strings.Replace(t, "*", "", 1)
}

// StripPtr removes all pointers from a type.
func StripPtr(t string) string {
	return strings.ReplaceAll(t, "*", "")
}

// MovePtr moves the same number of pointers from the given orig string into
// another string.
func MovePtr(orig, into string) string {
	ptr := strings.Count(orig, "*")
	return strings.Repeat("*", ptr) + into
}

// MoveCPtr moves the same number of pointers from the given orig string into
// another string as prefix, for C  types.
func MoveCPtr(orig, into string) string {
	ptr := strings.Count(orig, "*")
	return into + strings.Repeat("*", ptr)
}

// MoveTypePtr moves the pointer from src to dst. It overrides dst's pointer.
// A copy of dst is returned.
func MoveTypePtr(src, dst gir.Type) *gir.Type {
	ptr := strings.Count(src.CType, "*")
	dst.CType = strings.ReplaceAll(dst.CType, "*", "")
	dst.CType = dst.CType + strings.Repeat("*", ptr)
	return &dst
}

// CleanCType cleans the underlying C type of and special keywords for
// comparison.
func CleanCType(cType string, stripPtr bool) string {
	cType = cTypePrefixEraser.Replace(cType)
	cType = strings.TrimSpace(cType)
	if stripPtr {
		cType = StripPtr(cType)
	}
	return cType
}

// AnyTypeIsVoid returns true if AnyType is a void type.
func AnyTypeIsVoid(any gir.AnyType) bool {
	return any.Type != nil && any.Type.Name == "none"
}

// cgoPrimitiveTypes contains edge cases for referencing C primitive types from
// CGo.
//
// See https://gist.github.com/zchee/b9c99695463d8902cd33.
var cgoPrimitiveTypes = map[string]string{
	"long long": "longlong",

	"unsigned int":       "uint",
	"unsigned short":     "ushort",
	"unsigned long":      "ulong",
	"unsigned long long": "ulonglong",
}

// CGoTypeFromC converts a C type to a CGo type.
func CGoTypeFromC(cType string) string {
	originalCType := cType

	cType = cTypePrefixEraser.Replace(cType)
	cType = strings.ReplaceAll(cType, "*", "")
	cType = strings.TrimSpace(cType)

	if replace, ok := cgoPrimitiveTypes[cType]; ok {
		cType = replace
	}

	return MovePtr(originalCType, "C."+cType)
}

// TypeCGo is a helper function that invokes AnyTypeCGo.
func TypeCGo(typ *gir.Type) string {
	return AnyTypeCGo(gir.AnyType{Type: typ})
}

// AnyTypeCGo returns the CGo type for a GIR AnyType. An empty string is
// returned if none is made.
func AnyTypeCGo(any gir.AnyType) string {
	return CGoTypeFromC(AnyTypeC(any))
}

// AnyTypeC returns the C type for a GIR AnyType. An empty string is returned if
// none is made.
func AnyTypeC(any gir.AnyType) string {
	switch {
	case any.Array != nil:
		if any.Array.CType != "" {
			return CTypeFallback(any.Array.CType, any.Array.Name)
		}
		innerType := AnyTypeC(gir.AnyType{Type: any.Array.Type})
		if any.Array.FixedSize > 0 {
			return fmt.Sprintf("%s[%d]", innerType, any.Array.FixedSize)
		}
		return innerType + "*"
	case any.Type != nil:
		return CTypeFallback(any.Type.CType, any.Type.Name)
	default:
		return ""
	}
}

// AnyTypeIsPtr returns true if the AnyType contains a pointer.
func AnyTypeIsPtr(any gir.AnyType) bool {
	return strings.Contains(AnyTypeC(any), "*")
}

// AnyTypeCPrimitive converts AnyType to a C primitive type.
func AnyTypeCPrimitive(gen FileGenerator, any gir.AnyType) string {
	switch {
	case any.Array != nil:
		innerType := AnyTypeCPrimitive(gen, gir.AnyType{Type: any.Array.Type})
		if any.Array.FixedSize > 0 {
			return fmt.Sprintf("%s[%d]", innerType, any.Array.FixedSize)
		}
		return innerType + "*"
	case any.Type != nil:
		cType := AnyTypeC(any)

		if CountPtr(cType) == 0 && GIRIsPrimitive(cType) {
			return cType
		}

		resolved := Resolve(gen, *any.Type)
		if resolved != nil {
			if prim := resolved.AsDynamicLinkedCType(gen); prim != "" {
				return prim
			}
		}

		if prim := CTypeToPrimitive(cType); prim != "" {
			return prim
		}
	}

	return ""
}

// AnyTypeCGoPrimitive converts AnyType to a CGo primitive type.
func AnyTypeCGoPrimitive(gen FileGenerator, any gir.AnyType) string {
	return CGoTypeFromC(AnyTypeCPrimitive(gen, any))
}

// CTypeToPrimitive converts a C type to the primitive type.
func CTypeToPrimitive(cType string) string {
	// Keep this in sync with valueconverted.go's isRuntimeLinking check inside
	// resolveType. We probably can't make the C callback header generator use
	// the type converter, so we'll have to keep them in sync.

	switch cType {
	case "gpointer", "gconstpointer":
		return "gpointer"
	}

	if CountPtr(cType) > 0 {
		baseCType := CleanCType(cType, true)
		if GIRIsPrimitive(baseCType) {
			return MoveCPtr(cType, baseCType)
		} else {
			return MoveCPtr(cType, "void")
		}
	}

	if GIRIsPrimitive(cType) {
		return cType
	}

	return ""
}

// girCTypes maps some primitive GIR types to C types, because people who write
// GIR generators are lazy fucks who can't be bothered to fill in the right
// information.
var girCTypes = map[string]string{
	"utf8":     "gchar*",
	"filename": "gchar*",
}

// CTypeFallback returns the C type OR the GIR type if it's empty.
func CTypeFallback(c, gir string) string {
	if c != "" {
		return cTypePrefixEraser.Replace(c)
	}

	// Handle edge cases with a hard-coded type map. Thanks, GIR, for evading
	// needed information.
	ctyp, ok := girCTypes[gir]
	if ok {
		return ctyp
	}

	ctyp = cTypePrefixEraser.Replace(gir)
	if strings.IndexFunc(ctyp, unicode.IsUpper) != -1 {
		// CType contains an upper character, which implies that it's a GIR type
		// name that isn't a C type, so we can't convert it to C.
		return ""
	}

	return ctyp
}

// ReturnIsVoid returns true if the return type is void.
func ReturnIsVoid(ret *gir.ReturnValue) bool {
	return ret == nil || (ret != nil && AnyTypeIsVoid(ret.AnyType))
}

// noCasting contains types that must not be casted, usually because its
// equivalent type in Go has a different size and/or structure altogether.
var noCasting = map[string]struct{}{
	// C.uint and C.int have different sizes in Go.
	"gint":     {},
	"guint":    {},
	"utf8":     {},
	"filename": {},
}

// girToBuiltin maps the given GIR primitive type to a Go builtin type.
var girToBuiltin = map[string]string{
	"none":     "",
	"gboolean": "bool",
	"gfloat":   "float32",
	"gdouble":  "float64",
	"gint":     "int",
	"gssize":   "int",
	"gint8":    "int8",
	"gint16":   "int16",
	"gshort":   "int16",
	"gint32":   "int32",
	"glong":    "int32",
	"int32":    "int32",
	"gint64":   "int64",
	"guint":    "uint",
	"gsize":    "uint",
	"guchar":   "byte",
	"gchar":    "byte",
	"guint8":   "byte", // some weird cases
	"guint16":  "uint16",
	"gushort":  "uint16",
	"guint32":  "uint32",
	"gulong":   "uint32",
	"gunichar": "uint32",
	"guint64":  "uint64",
	"guintptr": "uintptr",
	"utf8":     "string",
	"filename": "string",
}

// GIRPrimitiveGo returns Go built-in types (primitive types and string). It
// returns an empty string if there's none.
func GIRBuiltinGo(typ string) string {
	return girToBuiltin[typ]
}

// GIRPrimitiveGo returns Go primitive types that can be copied by-value without
// doing any pointer work. It returns an empty string if there's none.
func GIRPrimitiveGo(typ string) string {
	t, ok := girToBuiltin[typ]
	if ok {
		return t
	}

	if !strings.HasPrefix(typ, "g") {
		return GIRPrimitiveGo("g" + typ) // maybe?
	}
	return ""
}

// GIRIsPrimitive returns true if the Go primitive type of the GIR type is valid
// and is not a string.
func GIRIsPrimitive(typ string) bool {
	t := GIRPrimitiveGo(typ)
	return t != "" && t != "string"
}

// FindParameter finds a parameter.
func FindParameter(c *gir.CallableAttrs, paramName string) *gir.ParameterAttrs {
	if c.Parameters == nil {
		return nil
	}

	if c.Parameters.InstanceParameter != nil {
		if c.Parameters.InstanceParameter.Name == paramName {
			return &c.Parameters.InstanceParameter.ParameterAttrs
		}
	}

	return FindParameterFromSlice(c.Parameters.Parameters, paramName)
}

// FindParameter finds a parameter from a slice of gir.Parameters.
func FindParameterFromSlice(params []gir.Parameter, paramName string) *gir.ParameterAttrs {
	for i, param := range params {
		if param.Name == paramName {
			return &params[i].ParameterAttrs
		}
	}

	return nil
}

// GuessParameterOutput guesses the parameter output using various clues to make
// up for GIR's painful shortcomings.
func GuessParameterOutput(param *gir.Parameter) string {
	switch param.Direction {
	case "out", "in", "inout":
		return param.Direction
	}

	// GIR is fucking miserable. Why can't they get these properly?
	if param.Doc != nil && strings.HasPrefix(param.Doc.String, "Return location") {
		param.Direction = "out"
		return "out"
	}

	// default
	param.Direction = "in"
	return "in"
}
