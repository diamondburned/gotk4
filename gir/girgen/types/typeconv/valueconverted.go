package typeconv

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// ValueConverted is the result of conversion for a single value.
//
// Quick convention note:
//
//    - {In,Out}Name is for the original name with no modifications.
//    - {In,Out}Call is used for the C or Go function arguments.
//
// Usually, these are the same, but they're sometimes different depending on the
// edge case.
type ValueConverted struct {
	*ConversionValue // original

	In  ValueName
	Out ValueName

	Conversion string

	// internal states
	ValueType
	Inner []ValueType

	conv *Converter

	// output writers
	p       *pen.PaperString
	inDecl  *pen.PaperString
	outDecl *pen.PaperString
	header  file.Header

	inArray bool // decides if record is ptr or not
	fail    bool
	final   bool
}

// ValueType contains the type information for one value.
type ValueType struct {
	Resolved       *types.Resolved // only for type conversions
	GoType         string
	IsPublic       bool
	NeedsNamespace bool
}

func (vt *ValueType) Import(h *file.Header, public bool) {
	if vt.NeedsNamespace {
		if public {
			h.ImportPubl(vt.Resolved)
		} else {
			h.ImportImpl(vt.Resolved)
		}
	}
}

// ValueName contains the different names to use during and after conversion.
type ValueName struct {
	Name    string
	Call    string
	Set     string
	Type    string
	Declare string
}

func newValueConverted(conv *Converter, value *ConversionValue) ValueConverted {
	return ValueConverted{
		ConversionValue: value,
		In: ValueName{
			Name: value.InName,
			Call: value.InName,
			Set:  value.InName,
		},
		Out: ValueName{
			Name: value.OutName,
			Call: value.OutName,
			Set:  value.OutName,
		},
		conv:    conv,
		p:       pen.NewPaperStringSize(1024), // 1KB
		inDecl:  pen.NewPaperStringSize(128),  // 0.1KB
		outDecl: pen.NewPaperStringSize(128),  // 0.1KB
	}
}

// Header returns the header of the current value.
func (value *ValueConverted) Header() *file.Header { return &value.header }

func (value *ValueConverted) isDone() bool { return value.final }

// finalize finalizes the value and returns true, or false is returned if the
// value is already finalized.
func (value *ValueConverted) finalize() bool {
	if value.final {
		return false
	}

	value.flush()

	// Allow GC to collect the internal buffers.
	value.inDecl = nil
	value.outDecl = nil
	value.p = nil

	value.final = true
	return true
}

// flush commits the writers.
func (value *ValueConverted) flush() {
	value.In.Declare = value.inDecl.String()
	value.Out.Declare = value.outDecl.String()
	value.Conversion = value.p.String()
}

func (value *ValueConverted) vtmpl(tmpl string) {
	value.p.LineTmpl(value, tmpl)
}

func (value *ValueConverted) Logln(lvl logger.Level, v ...interface{}) {
	value.conv.Logln(lvl, logger.Prefix(v, value.logPrefix())...)
}

// resolveType resolves the value type to the resolved field. If inputC is true,
// then the input type is set to the CGo type, otherwise the Go type is set.
func (value *ValueConverted) resolveType(conv *Converter) bool {
	if value.Resolved != nil && value.Inner != nil && value.In.Type != "" && value.Out.Type != "" {
		// already resolved
		return true
	}

	// Treat this specific case differently. We only care about when we index
	// the slice, not the slice itself, so setting the inArray below is not what
	// we want.
	callbackInArray := value.inArray && conv.Callback

	// Ensure that the array type matches the inner type. Some functions violate
	// this, e.g. g_spawn_command_line_sync().
	if value.Array != nil {
		value.inArray = true

		if value.Array.Type == nil {
			value.Logln(logger.Debug, "nested array not supported")
			return false
		}

		if value.Array.Type.CType == "" {
			// Copy the inner type so we don't accidentally change a reference.
			array := *value.Array
			typ := *array.Type

			// TODO: delegate this routine to after resolving.

			if types.CleanCType(value.Array.CType, true) == "void" {
				// We can't dereference the array type to get a void type, so we
				// have to guess from the GIR type. Thanks, GIR.
				typ.CType = types.CTypeFallback("", typ.Name)
			} else {
				// Dereference the inner type.
				typ.CType = strings.Replace(value.Array.CType, "*", "", 1)
			}

			array.Type = &typ
			value.Array = &array
		}
	}

	var typ *gir.Type
	var cType string

	switch {
	case value.Array != nil:
		// Resolve the array's inner type.
		typ = value.Array.Type
		cType = value.Array.CType
	case value.Type != nil:
		typ = value.Type
		cType = value.Type.CType
	}

	if typ == nil {
		value.Logln(logger.Debug, "value missing both array and type")
		return false
	}

	if value.ValueType == (ValueType{}) {
		var ok bool
		value.ValueType, ok = value.resolveTypeInner(conv, typ)
		if !ok {
			return false
		}
	}

	value.Inner = make([]ValueType, 0, len(typ.Types))

	for i := range typ.Types {
		it, ok := value.resolveTypeInner(conv, &typ.Types[i])
		if !ok {
			value.Logln(logger.Debug, "resolveTypeInner at inner type", i)
			return false
		}

		value.Inner = append(value.Inner, it)
	}

	if cType == "" {
		// Fix missing CGo type, which sometimes happens when we're in a
		// subtype.
		cType = value.Resolved.CType
	}

	cgoType := types.CGoTypeFromC(cType)
	if cgoType == "" {
		value.Logln(logger.Debug, "empty CGoType")
		return false
	}

	if !strings.Contains(cgoType, "*") && value.InContainer {
		// The inner type inside a container must always be a pointer, so if
		// it's not, then make it one.
		cgoType = "*" + cgoType
	}

	if value.ParameterIsOutput() {
		// Output parameter, so neutralize the pointer type.
		cgoType = strings.Replace(cgoType, "*", "", 1)
	}

	if value.conv.isRuntimeLinking() {
		// Runtime-linking means we do not have access to any of the actual C
		// types, since we intentionally don't want to import them.
		// Instead, we must mask those types into primitive GLib types.
		if ptr := types.CountPtr(cgoType); ptr > 0 {
			// Using gpointer doesn't quite work right here, so we just use
			// C.void.
			cgoType = types.MovePtr(cgoType, "C.void")
		} else if types.GIRIsPrimitive(cType) {
			// cgoType OK as it is.
		} else {
			value.Logln(logger.Debug, "cType", cType, "is not primitive")
			return false
		}
	}

	switch value.Resolved.GType {
	case "GLib.HashTable":
		if value.Resolved.Ptr != 1 {
			// Unknown ptr rule.
			value.Logln(logger.Debug, "unknown HashTable pointer rule", value.Resolved.Ptr)
			return false
		}

		if len(value.Inner) != 2 {
			value.Logln(logger.Debug, "HashTable types != 2")
			return false
		}

		value.GoType = fmt.Sprintf(
			"map[%s]%s",
			value.Inner[0].GoType,
			value.Inner[1].GoType,
		)
	case "GLib.List", "GLib.SList":
		if value.Resolved.Ptr != 1 {
			value.Logln(logger.Debug, "unknown List pointer rule", value.Resolved.Ptr)
			return false
		}

		if len(value.Inner) != 1 {
			value.Logln(logger.Debug, "List missing inner type")
			return false
		}

		value.GoType = fmt.Sprintf("[]%s", value.Inner[0].GoType)
	default:
		if value.Array != nil {
			if value.Array.FixedSize > 0 {
				value.GoType = fmt.Sprintf("[%d]%s", value.Array.FixedSize, value.GoType)
			} else {
				value.GoType = fmt.Sprintf("[]%s", value.GoType)
			}
		}
	}

	switch value.Direction {
	case ConvertCToGo:
		value.In.Type = cgoType
		value.Out.Type = value.GoType
	case ConvertGoToC:
		value.Out.Type = cgoType
		value.In.Type = value.GoType
	default:
		value.Logln(logger.Error, "unknown direction", value.Direction)
		return false
	}

	if !callbackInArray && value.ParameterIsOutput() {
		switch value.Direction {
		case ConvertCToGo:
			value.In.Call = "&" + value.In.Call
		case ConvertGoToC:
			value.Out.Set = "*" + value.Out.Set
		}
	}

	if value.GoType == "Allocation" {
		value.Logln(logger.Error, "2: allocation value while record, ptr =", value.Resolved.Ptr)
	}

	value.inDecl.Linef("var %s %s // in", value.InName, value.In.Type)
	value.outDecl.Linef("var %s %s // out", value.OutName, value.Out.Type)

	return true
}

// manualTypes is a set of GTypes that are manually converted internally, so
// they don't actually reference the package that they belong to.
var manualTypes = map[string]func() bool{
	"GLib.HashTable": nil,
	"GLib.ByteArray": nil,
	"GLib.List":      nil,
	"GLib.SList":     nil,
}

func (value *ValueConverted) resolveTypeInner(conv *Converter, typ *gir.Type) (ValueType, bool) {
	if typ == nil {
		return ValueType{}, false
	}

	// ResolveType already checks this, but we can early bail.
	if !typ.IsIntrospectable() {
		return ValueType{}, false
	}

	// Check nested types.
	for _, unsupported := range types.UnsupportedCTypes {
		if unsupported == typ.Name {
			return ValueType{}, false
		}
	}

	resolved := types.Resolve(conv.rgen(), *typ)
	if resolved == nil {
		value.Logln(logger.Debug, "can't resolve", types.TypeCGo(typ))
		return ValueType{}, false
	}

	// Allow gio.Cancellable to be a context.Context. We're not doing this
	// inside Resolve, because this is the only case where we actually want
	// this.
	if !value.ParameterIsOutput() && value.ParameterIndex.Index() != -1 &&
		resolved.Ptr == 1 && resolved.GType == "Gio.Cancellable" {

		resolved = types.BuiltinType("context", "Context", *typ)
		resolved.Ptr--
	}

	vType := ValueType{
		Resolved:       resolved,
		NeedsNamespace: resolved.NeedsNamespace(conv.fgen.Namespace()),
	}

	var intentionalPtr bool
	var needsImporting bool

	// HashTable's handling doesn't import the glib package's implementation, so
	// we don't import it if that's the case. Ideally, HashTable should be
	// resolved to a map directly in types/resolved.go, but that requires a
	// refactor.
	f, ok := manualTypes[vType.Resolved.GType]
	if ok && (f == nil || f()) {
		// Hack so the caller doesn't add the import.
		vType.NeedsNamespace = false
		needsImporting = false
	} else {
		needsImporting = true

		alwaysPointer := false ||
			vType.Resolved.IsUnion() ||
			vType.Resolved.IsRecord()

		// TODO: do the same thing for classes as well.
		if !value.inArray && alwaysPointer {
			// Fix the mismatch that would happen if the parameter is a C
			// output parameter, since we normalize that elsewhere.
			intentionalPtr = true

			// Account for some edge cases where the struct already has a
			// pointer. The intentionalPtr logic is a bit flawed, but this check
			// makes it work for cases like TreeView.DestRowAtPos.
			if vType.Resolved.Ptr > 0 {
				vType.Resolved.Ptr--
			}

			// Records are implicitly pointers, and methods of records have
			// pointer receivers, so it's more correct this way. Only do this if
			// it's not in an array, though.
			//
			// TODO: REMOVE ME IF THINGS DON'T BUILD BECAUSE OF RECORD POINTERS.
			if vType.Resolved.Ptr == 0 {
				vType.Resolved.Ptr = 1
			}
		}
	}

	// If this is the output parameter, then the pointer count should be less.
	// This only affects the Go type.
	if value.ParameterIsOutput() && vType.Resolved.Ptr > 0 && !intentionalPtr {
		vType.Resolved.Ptr--
	}

	vType.GoType = vType.Resolved.ImplType(vType.NeedsNamespace)

	// Only use the PublicType (an interface) when the parameter is not a
	// returning/output parameter.
	//
	// This change makes sense (even though it doesn't seem to with Go),
	// because GLib interfaces don't actually describe what a class has to
	// implement publicly; it only describes private methods that are made to
	// explicitly implement an interface.
	//
	// The above design is different from Go, where interfaces describe what a
	// class (struct) has to implement publicly, and so its polymorphism is far
	// simpler: there is no need to worry about "virtual" methods, and all
	// methods are the same (with the exception of exporting).
	//
	// Because we try to map GLib class methods to Go struct methods, we can't
	// really expect GLib interfaces to translate the same, so we have to
	// return the concrete type instead. The user can still get the underlying
	// type by calling Cast(), but it won't be a simple Go assertion.
	if !value.KeepType {
		if resolved.IsAbstractClass() || (resolved.IsInterface() && !value.ParameterIsReturn()) {
			vType.GoType = vType.Resolved.PublicType(vType.NeedsNamespace)
			vType.IsPublic = true
		}

		if (resolved.IsClass() || resolved.IsInterface()) && !vType.IsPublic && vType.Resolved.Ptr == 0 {
			// Non-abstract classes should always be a pointer.
			vType.GoType = "*" + vType.GoType
			vType.Resolved.Ptr = 1
		}
	}

	if needsImporting {
		vType.Import(&value.header, value.IsPublic)
	}

	if vType.Resolved.IsCallback() {
		value.header.AddCallback(
			vType.Resolved.Extern.NamespaceFindResult,
			vType.Resolved.Extern.Type.(*gir.Callback),
		)
	}

	return vType, true
}

// cgoSetObject generates a glib.Take or glib.AssumeOwnership into a new
// function. This should only be used for C to Go conversion.
func (value *ValueConverted) cgoSetObject(conv *Converter) bool {
	var gobjectFunction string
	if value.ownershipIsTransferring() {
		// Full or container means we implicitly own the object, so we must
		// not take another reference.
		gobjectFunction = "AssumeOwnership"
	} else {
		// Else the object is either unowned by us or it's a floating
		// reference. Take our own or sink the object.
		gobjectFunction = "Take"
	}

	value.header.NeedsExternGLib()
	value.header.Import("unsafe")

	m := gotmpl.M{
		"Value": value,
		"Func":  gobjectFunction,
	}

	if value.Resolved.IsExternGLib("Object") {
		// Shortcut for GObject.
		value.p.LineTmpl(m, `
			<.Value.Out.Set> = coreglib.<.Func>(unsafe.Pointer(<.Value.InPtr 1><.Value.InName>))
		`)
		return true
	}

	if value.IsPublic {
		// Used for verbosity.
		m["GoType"] = value.Resolved.PublicType(true)
		// Require the abstract cast if we have an abstract type.
		value.header.NeedsExternGLib()
		value.p.Descend()
		value.p.LineTmpl(m, `
			objptr := unsafe.Pointer(<.Value.InPtr 1><.Value.InName>)
			<if not (or .Value.Optional .Value.Nullable) ->
			if objptr == nil {
				panic("object of type <.GoType> is nil")
			}
			<end>
			object := coreglib.<.Func>(objptr)
			casted := <.Value.OutPtr 0>object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(<.Value.Out.Type>)
				return ok
			})
			rv, ok := casted.(<.Value.Out.Type>)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching <.GoType>")
			}
			<.Value.Out.Set> = rv
		`)
		value.p.Ascend()
		return true
	}

	if !value.NeedsNamespace {
		value.p.LineTmpl(m, `
			<.Value.Out.Set> = <.Value.OutPtr 1><.Value.Resolved.WrapName false ->
				(coreglib.<.Func>(unsafe.Pointer(<.Value.InPtr 1><.Value.InName>)))
		`)
		return true
	}

	if tree := types.NewTree(conv.fgen); tree.ResolveFromType(value.Resolved) {
		wrap := tree.WrapInNamespace("obj", &value.header, conv.Parent.NamespaceFindResult)
		if value.OutPtr(1) == "*" {
			// Dereference the wrapped struct value by removing the &.
			wrap = strings.TrimPrefix(wrap, "&")
		}
		m["Wrap"] = wrap

		value.p.LineTmpl(m, `{
			obj := coreglib.<.Func>(unsafe.Pointer(<.Value.InPtr 1><.Value.InName>))
			<.Value.Out.Set> = <.Wrap>
		}`)

		return true
	}

	return false
}

// writeMalloc ensures bound-checking.
func (value *ValueConverted) writeMalloc(inner *ValueConverted, lenOf string, add1 bool) {
	lenOf = "len(" + lenOf + ")"
	if add1 {
		lenOf = "(" + lenOf + " + 1)"
	}

	value.p.Linef(
		"%s = (%s)(C.calloc(C.size_t(%s), C.size_t(%s)))",
		value.Out.Set, value.Out.Type, lenOf, inner.csizeof(),
	)

	// I give up.

	// value.header.Import("strconv")
	// value.header.IncludeC("stdint.h") // for SIZE_MAX

	// value.p.Linef("if size := uint(%s) * %s; size > C.SIZE_MAX {", lenOf, inner.csizeof())
	// value.p.Linef(`  panic("malloc overflow: " + strconv.FormatUint(uint64(size), 10))`)
	// value.p.Linef("} else {")
	// value.p.Linef("  %s = (%s)(C.malloc(C.size_t(size)))", value.Out.Set, value.Out.Type)
	// value.p.Linef("}")
}

func (value *ValueConverted) csizeof() string {
	// Arrays are lists of pointers.
	if strings.Contains(types.AnyTypeC(value.AnyType), "*") {
		return value.ptrsz()
	}

	if value.Resolved == nil {
		// Erroneous case.
		return "/* uncertain */" + value.ptrsz()
	}

	return "C.sizeof_" + types.CleanCType(value.Resolved.CType, true)
}

func (value *ValueConverted) ptrsz() string {
	value.header.Import("unsafe")
	// Size of a pointer is the same as uint.
	return "unsafe.Sizeof(uint(0))"
}

func (value *ValueConverted) logPrefix() string {
	var prefix string
	var cgoType string

	switch value.Direction {
	case ConvertCToGo:
		prefix = fmt.Sprintf("C %s -> Go %s", value.InName, value.OutName)
		cgoType = value.In.Type
	case ConvertGoToC:
		prefix = fmt.Sprintf("Go %s -> C %s", value.InName, value.OutName)
		cgoType = value.Out.Type
	default:
		return ""
	}

	prefix += fmt.Sprintf(" (%s)", value.ParameterAttrs.Direction)

	if cgoType != "" {
		prefix += fmt.Sprintf(" (%s)", cgoType)
	}

	if value.GoType != "" {
		prefix += fmt.Sprintf(" (%s)", value.GoType)
	}

	return prefix + ":"
}

// castPrimitive is used to cast primitives of any pointer level from C to Go or
// vice-versa. It is useful for casting enums and bitfields.
func (value *ValueConverted) castPrimitive() bool {
	// if value.Resolved.Ptr == 0 {
	// 	value.p.Linef("%s = %s(%s)", value.Out.Set, value.Out.Type, value.InName)
	// 	return true
	// }

	// // BEWARE: We cannot just cast the unsafe.Pointer of this type, because Go's
	// // integer types might be different from C. We MUST dereference and copy it
	// // back, if needed.
	// if value.Resolved.Ptr == 1 {
	// 	value.p.Descend()
	// 	defer value.p.Ascend()

	// 	// We know this is a pointer of depth 1, so we can trim a pointer from
	// 	// the C type.
	// 	valueCType := strings.Replace(value.Out.Type, "*", "", 1)

	// 	value.p.Linef("tmp := %s(*%s)", valueCType, value.InName)
	// 	value.p.Linef("%s = &tmp", value.Out.Set)
	// }

	// return false

	hasIn := value.InNPtr()
	hasOut := value.OutNPtr()

	if value._ptr(hasIn, hasOut) == "" && value.fail {
		return false
	}

	// With the way enums and bitfields are generated, this cast is completely
	// safe, since we're casting back and forth the same exact types.
	if hasIn == hasOut {
		if value.Resolved.Ptr > 0 {
			value.header.Import("unsafe")
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(%s))",
				value.Out.Set, value.Out.Type, value.In.Name,
			)
		} else {
			value.p.Linef("%s = %s(%s)", value.Out.Set, value.Out.Type, value.In.Name)
		}
		return true
	}

	// Difference is 1. Let the code figure it out.
	value.p.Linef(
		"%s = %s(unsafe.Pointer(%s))",
		value.Out.Set,
		// What have I done?! I think this only works if in is a pointer. It
		// will explode otherwise. Which is fine, I guess. OutCast will fail
		// otherwise.
		value.OutCast(1),
		value.InName,
	)

	return !value.fail
}

// isPtr checks pointer coherency for C types and Go types. It's mostly used to
// guarantee that conversion routines get what they expect.
func (value *ValueConverted) isPtr(wantC int) bool {
	// See this same piece of code in convertRef for more information.
	if value.Resolved.IsGpointer() && wantC > 0 {
		wantC--
	}

	switch value.Direction {
	case ConvertCToGo:
		return types.CountPtr(value.In.Type) == wantC
	case ConvertGoToC:
		return types.CountPtr(value.Out.Type) == wantC
	default:
		return false
	}

	// Rationale for not verifying Go pointer offset is that the pointer offset
	// is already determined in the type resolver routine, so repeating that
	// information is redundant.
	//
	// Edit: this rationale does NOT work because the type resolver only has the
	// wanted Go pointer information up to the point of creating a new
	// ResolvedType, and there's no way we can get it back. This routine may not
	// need to verify the Go pointer, but the conversiron routine will.
}

// InNamePtrPubl adds in an edge case if the value being inputted is possibly a
// Go interface.
func (value *ValueConverted) InNamePtrPubl(want int) string {
	if want > 0 && value.IsPublic {
		want--
	}
	return value.InNamePtr(want)
}

// InNamePtr returns the name with the pointer prefixed using InPtr.
func (value *ValueConverted) InNamePtr(want int) string {
	ptr := value.InPtr(want)
	if ptr == "" {
		return value.InName
	}
	return fmt.Sprintf("(%s%s)", ptr, value.InName)
}

func (value *ValueConverted) InNPtr() int {
	// Account for gpointer.
	has := types.CountPtr(value.In.Type)
	if value.Direction == ConvertCToGo && value.Resolved.IsGpointer() {
		has++
	}
	return has
}

func (value *ValueConverted) InPtr(want int) string {
	has := value.InNPtr()
	return value._ptr(has, want)
}

// OutCast returns the left-hand side consisting of the pointer dereference and
// the pointer-prefixed type.
func (value *ValueConverted) OutCast(want int) string {
	ptr := value.OutPtr(want)
	if ptr == "" && !strings.Contains(value.Out.Type, "*") {
		return value.Out.Type
	}
	return fmt.Sprintf("%s(%s%s)", ptr, ptr, value.Out.Type)
}

// OutInNamePtr is like InNamePtr but for OutInPtr.
func (value *ValueConverted) OutInNamePtr(want int) string {
	ptr := value.OutInPtr(want)
	if ptr == "" {
		return value.Out.Name
	}
	return fmt.Sprintf("(%s%s)", ptr, value.Out.Name)
}

// OutInPtr returns the left-hand side for the output name and type SPECIFICALLY
// for inputting the output name elsewhere, like giving it to SetFinalizer.
func (value *ValueConverted) OutInPtr(want int) string {
	has := types.CountPtr(value.Out.Type)
	return value._ptr(has, want)
}

func (value *ValueConverted) OutNPtr() int {
	has := types.CountPtr(value.Out.Type)
	// Account for gpointer.
	if value.Direction == ConvertGoToC && value.Resolved.IsGpointer() {
		has++
	}
	return has
}

func (value *ValueConverted) OutPtr(want int) string {
	has := value.OutNPtr()

	ptr := value._ptr(want, has)
	if ptr == "&" {
		// Refuse to reference the value we converted, since that requires a
		// temporary variable.
		value.fail = true
		value.Logln(logger.Debug, "OutPtr refusing to reference, has", has, "want", want)
		return ""
	}

	return ptr
}

func (value *ValueConverted) _ptr(has, want int) string {
	if difference(has, want) > 1 {
		value.fail = true
		value.Logln(logger.Debug, "pointer difference too high, has", has, "want", want)
		return ""
	}

	switch {
	case has < want:
		return "&"
	case has > want:
		return "*"
	default:
		return ""
	}
}

func difference(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

// ShouldFree returns true if the C value must be freed once we're done.
func (value *ValueConverted) ShouldFree() bool {
	// goReceiving is true when we're receiving the C value.
	goReceiving := value.ParameterIndex == ReturnValueIndex || value.ParameterIsOutput()
	if value.conv.Callback {
		goReceiving = !goReceiving
	}
	// If we're not converting C to Go, then we're probably in a callback, so
	// the ownership is flipped.
	// if value.Direction != ConvertCToGo {
	// 	goReceiving = !goReceiving
	// }

	if goReceiving {
		return value.ownershipIsTransferring()
	}

	return !value.ownershipIsTransferring()
}

// MustRealloc returns true if we need to malloc our values to give it to C.
// Generally, if a conversion routine has a no-alloc path, it should check
// MustRealloc first. If MustRealloc is true, then it must check ShouldFree.
//
//    if value.MustAlloc() {
//        v = &oldValue
//    } else {
//        v = malloc()
//        if value.ShouldFree() {
//            defer free(v)
//        }
//    }
//
func (value *ValueConverted) MustRealloc() bool {
	// goGiving is true when we're giving the C value.
	goGiving := value.ParameterIndex > -1 && !value.ParameterIsOutput()
	if value.conv.Callback {
		goGiving = !goGiving
	}
	// If we're not converting Go to C, then we're probably in a callback, so
	// the ownership is flipped.
	// if value.Direction != ConvertGoToC {
	// 	goGiving = !goGiving
	// }

	if goGiving {
		return value.ownershipIsTransferring()
	}

	return !value.ownershipIsTransferring()
}
