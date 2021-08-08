package types

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
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
	return len(rec.Fields) == 0 || rec.GLibGetType == "intern"
}

var acceptableFreeNames = []string{
	"unref",
	"free",
	"destroy",
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
	if m := findMethodName(record, "free"); m != nil {
		return m
	}
	if m := findMethodName(record, "destroy"); m != nil {
		return m
	}
	return nil
}

// RecordHasUnref returns the unref method if it has one.
func RecordHasUnref(record *gir.Record) *gir.Method {
	return findMethodName(record, "unref")
}

// RecordHasRef returns the ref method if it has one.
func RecordHasRef(record *gir.Record) *gir.Method {
	return findMethodName(record, "ref")
}

func findMethodName(record *gir.Record, name string) *gir.Method {
	for i, method := range record.Methods {
		if method.Name == name && methodCanCallDirectly(&method) {
			return &record.Methods[i]
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
	"ObjectProperty":    {},
	"SetObjectProperty": {},
	"Cast":              {},
	"Native":            {},
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
	sliceIx := strings.Index(typ, "[]")
	if sliceIx == -1 {
		sliceIx = len(typ) - 1
	}

	return strings.Count(typ[:sliceIx], "*")
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
		cType = strings.ReplaceAll(cType, "*", "")
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
		return CTypeFallback(any.Array.CType, any.Array.Name)
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

	return cTypePrefixEraser.Replace(gir)
}

// ReturnIsVoid returns true if the return type is void.
func ReturnIsVoid(ret *gir.ReturnValue) bool {
	return ret == nil || (ret != nil && AnyTypeIsVoid(ret.AnyType))
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
	"utf8":     "string",
	"filename": "string",
}

// GIRPrimitiveGo returns Go built-in types (primitive types and string). It
// returns an empty string if there's none.
func GIRBuiltinGo(typ string) string {
	v, _ := girToBuiltin[typ]
	return v
}

// GIRPrimitiveGo returns Go primitive types that can be copied by-value without
// doing any pointer work. It returns an empty string if there's none.
func GIRPrimitiveGo(typ string) string {
	gp, ok := girToBuiltin[typ]
	if !ok || gp == "string" {
		return ""
	}
	return gp
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

	for i, param := range c.Parameters.Parameters {
		if param.Name == paramName {
			return &c.Parameters.Parameters[i].ParameterAttrs
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
		return "out"
	}

	// default
	return "in"
}
