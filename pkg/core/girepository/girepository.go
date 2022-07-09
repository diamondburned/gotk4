// Packagee girepository provides convenient bindings around libgirepository's C
// API. It is handwritten, so the API surface is very abstracted and is kept to
// the bare minimum. As such, it is not a 1:1 binding.
package girepository

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #include <girepository.h>
import "C"

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

var repository = C.g_irepository_get_default()

// LoadFlags are flags that control how a typelib is loaded.
type LoadFlags C.GIRepositoryLoadFlags

// LoadFlags constants.
const (
	LoadFlagLazy = LoadFlags(C.G_IREPOSITORY_LOAD_FLAG_LAZY)
)

// Argument is a type to a GIArgument union.
type Argument C.GIArgument

/*
// SetArgument sets v's underlying value into the Argument union.
func SetArgument[T any](arg *Argument, v T) {
	if uintptr(len(Argument{})) != unsafe.Sizeof(uintptr(0)) {
		panic("BUG: GIArgument is not as big as a pointer")
	}

	// Treat arg as a *T, which if T = gpointer, would mean that we're setting
	// the v_pointer field. Dereference the pointer and set it as expected.
	*(*T)(unsafe.Pointer(arg)) = v
}

// GetArgument gets an argument from arg.
func GetArgument[T any](arg *Argument) T {
	// Go actually optimizes this sizeof() into a constant. See
	// https://godbolt.org/z/dYxcM8ssq.
	var zero T
	if unsafe.Sizeof(zero) > uintptr(len(Argument{})) {
		panic("BUG: sizeof(T) larger than size of GIArgument")
	}

	return *(*T)(unsafe.Pointer(arg))
}
*/

func (a Argument) String() string {
	return fmt.Sprint(*(*unsafe.Pointer)(unsafe.Pointer(&a)))
}

// Require forces the namespace to be loaded if it isn't already. If namespace
// is not loaded, this function will search for a ".typelib" file using the
// repository search path. In addition, a version version of namespace may be
// specified. If version is not specified, the latest will be used.
func Require(namespace, version string, flags LoadFlags) {
	cnamespace := unsafeCStr(&namespace)
	cversion := unsafeCStr(&version)

	var gerr *C.GError
	C.g_irepository_require(repository, cnamespace, cversion, C.GIRepositoryLoadFlags(flags), &gerr)
	if gerr != nil {
		log.Panicf(
			"GIRepository: cannot require %q v%s: %v",
			trimNull(namespace), trimNull(version), gerror.Take(unsafe.Pointer(gerr)),
		)
	}
}

type (
	infoKey  [3]string
	infoCKey [3]*C.gchar
)

func (k infoKey) toC() infoCKey {
	bigCStr := unsafe.Pointer(C.malloc(
		// Add known lengths and the n null terminators.
		C.size_t(len(k[0]) + len(k[1]) + len(k[2]) + len(infoKey{})),
	))

	var ckey infoCKey
	ckey[0] = (*C.gchar)(bigCStr)
	ckey[1] = cpyCStr(ckey[0], k[0])
	ckey[2] = cpyCStr(ckey[1], k[1])
	cpyCStr(ckey[2], k[2])

	return ckey
}

func (k infoKey) String() string {
	var name string
	switch {
	case k[1] == "":
		name = k[0]
	case k[2] == "":
		name = k[0] + "." + k[1]
	default:
		name = k[0] + "." + k[1] + "." + k[2]
	}
	return name
}

func cpyCStr(cptr *C.gchar, gostr string) (next *C.gchar) {
	cbytes := unsafe.Slice((*byte)(unsafe.Pointer(cptr)), len(gostr)+1)
	copy(cbytes, []byte(gostr))
	cbytes[len(gostr)] = 0

	return (*C.gchar)(unsafe.Add(unsafe.Pointer(cptr), len(cbytes)))
}

func (k infoCKey) free() {
	C.free(unsafe.Pointer(k[0]))
}

var infoCache = struct {
	sync.RWMutex
	m map[infoKey]*Info // {namespace, funcName}
}{
	m: make(map[infoKey]*Info, 1024),
}

type Info struct {
	info unsafe.Pointer
	keys infoKey
}

// MustFind finds or panics.
func MustFind(namespace, name string) *Info {
	info := Find(namespace, name)
	if info == nil {
		log.Panicf("girepository: unknown %s.%s", namespace, name)
	}
	return info
}

// Find finds name from the given namespace.
func Find(namespace, name string) *Info {
	return findInfo([3]string{namespace, name}, func(k infoCKey) unsafe.Pointer {
		return unsafe.Pointer(C.g_irepository_find_by_name(repository, k[0], k[1]))
	})
}

// String formats the type Info as a string describing its name.
func (i *Info) String() string {
	if i == nil {
		return "<nil>"
	}

	baseInfo := (*C.GIBaseInfo)(i.info)
	infoType := C.g_info_type_to_string(C.g_base_info_get_type(baseInfo))

	return fmt.Sprintf("%s %s", C.GoString(infoType), i.keys.String())
}

// RegisteredGType returns the GType of the type belonging to this Info.
func (i *Info) RegisteredGType() coreglib.Type {
	return coreglib.Type(C.g_registered_type_info_get_g_type((*C.GIRegisteredTypeInfo)(i.info)))
}

// StructFieldOffset gets the offset of the field for the record that is i.
func (i *Info) StructFieldOffset(name string) uintptr {
	k := i.keys
	k[2] = name

	field := findInfo(k, func(ckey infoCKey) unsafe.Pointer {
		return unsafe.Pointer(C.g_struct_info_find_field((*C.GIStructInfo)(i.info), ckey[2]))
	})

	offset := C.g_field_info_get_offset((*C.GIFieldInfo)(field.info))
	if offset < 0 {
		panic("ERROR: girepository: field_info_get_offset returned negative")
	}

	return uintptr(offset)
}

// InvokeFunction invokes this BaseInfo as a FunctionInfo.
func (i *Info) InvokeFunction(in, out []Argument) Argument {
	return invokeFunc(i, in, out)
}

// InvokeMethod invokes the method of this ClassInfo with the given name.
func (i *Info) InvokeMethod(name string, in, out []Argument) Argument {
	k := i.keys
	k[2] = name

	method := findInfo(k, func(ckey infoCKey) unsafe.Pointer {
		return unsafe.Pointer(C.g_object_info_find_method((*C.GIObjectInfo)(i.info), ckey[2]))
	})

	return invokeFunc(method, in, out)
}

// InvokeIfaceMethod invokes the method of this InterfaceInfo with the given name.
func (i *Info) InvokeIfaceMethod(name string, in, out []Argument) Argument {
	k := i.keys
	k[2] = name

	method := findInfo(k, func(ckey infoCKey) unsafe.Pointer {
		return unsafe.Pointer(C.g_interface_info_find_method((*C.GIInterfaceInfo)(i.info), ckey[2]))
	})

	return invokeFunc(method, in, out)
}

// InvokeRecordMethod invokes the method of this StructInfo with the given name.
func (i *Info) InvokeRecordMethod(name string, in, out []Argument) Argument {
	k := i.keys
	k[2] = name

	method := findInfo(k, func(ckey infoCKey) unsafe.Pointer {
		return unsafe.Pointer(C.g_struct_info_find_method((*C.GIStructInfo)(i.info), ckey[2]))
	})

	return invokeFunc(method, in, out)
}

func findInfo(k infoKey, f func(infoCKey) unsafe.Pointer) *Info {
	infoCache.RLock()
	info, ok := infoCache.m[k]
	infoCache.RUnlock()

	if ok {
		log.Printf("for %q got %p (cached)", k, info.info)
		return info
	}

	infoCache.Lock()
	defer infoCache.Unlock()

	ckey := k.toC()
	defer ckey.free()

	infoPtr := f(ckey)
	if infoPtr == nil {
		log.Panicln("girepository: cannot find", k.String())
	}

	info = &Info{
		info: infoPtr,
		keys: k,
	}
	log.Printf("for %q got %p", k, infoPtr)

	infoCache.m[k] = info
	return info
}

func invokeFunc(info *Info, in, out []Argument) Argument {
	var (
		argsIn  *C.GIArgument
		argsOut *C.GIArgument
	)

	log.Printf("invoking %s (%s %p)", info, C.GoString(C.g_function_info_get_symbol((*C.GIFunctionInfo)(info.info))), info.info)
	log.Printf("  in:  %v (n=%d)", in, C.g_callable_info_get_n_args((*C.GICallableInfo)(info.info)))
	log.Printf("  out: %v", out)

	if len(in) > 0 {
		argsIn = (*C.GIArgument)(&in[0])
	}
	if len(out) > 0 {
		argsOut = (*C.GIArgument)(&out[0])
	}

	var ret Argument
	var invokeErr *C.GError

	invoked := C.g_function_info_invoke(
		(*C.GIFunctionInfo)(info.info),
		(*C.GIArgument)(argsIn), C.int(len(in)),
		(*C.GIArgument)(argsOut), C.int(len(out)),
		(*C.GIArgument)(&ret),
		&invokeErr,
	) == C.TRUE

	log.Println("  ret:", ret)

	if !invoked {
		var tail string
		if invokeErr != nil {
			tail = ": " + gerror.Take(unsafe.Pointer(invokeErr)).Error()
		}
		log.Panic("cannot invoke ", info.String(), tail)
	}

	return ret
}

func trimNull(str string) string {
	return strings.TrimSuffix(str, "\x00")
}

func unsafeCStr(gostr *string) *C.gchar {
	if len(*gostr) == 0 {
		return nil
	}

	if (*gostr)[len(*gostr)-1] != 0 {
		log.Panicf("Go string %q should be null-terminated with \\x00", *gostr)
	}

	header := (*reflect.StringHeader)(unsafe.Pointer(gostr))
	return (*C.gchar)(unsafe.Pointer(header.Data))
}
