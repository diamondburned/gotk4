package girepository

// #include <girepository.h>
import "C"

import (
	"log"
	"reflect"
	"strings"
	"unsafe"
)

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

func cpyCStr(cptr *C.gchar, gostr string) (next *C.gchar) {
	cbytes := unsafe.Slice((*byte)(unsafe.Pointer(cptr)), len(gostr)+1)
	copy(cbytes, []byte(gostr))
	cbytes[len(gostr)] = 0

	return (*C.gchar)(unsafe.Add(unsafe.Pointer(cptr), len(cbytes)))
}

func cstr2bytes[T IntegerType](cstr *C.gchar, len int) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(cptr)), len)
}
