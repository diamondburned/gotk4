package ptr

import (
	"reflect"
	"unsafe"
)

// Add adds into pointer p the offset given.
func Add(p unsafe.Pointer, offset uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + offset)
}

// SetSlice sets slice's backing array pointer, length and capacity to the
// given parameters.
//
// Example usage:
//
//    var strings []string
//    ptr.SetSlice(unsafe.Pointer(&strings), pointer, 10)
//
func SetSlice(slice, data unsafe.Pointer, len int) {
	h := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	h.Data = uintptr(data)
	h.Len = len
	h.Cap = len
}

// Slice gets the backing data pointer of a slice.
func Slice(slice unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(slice)).Data)
}

// Strlen calculates the length of a null-terminated string in the given
// pointer.
func Strlen(str unsafe.Pointer) int {
	var cum int
	for p := str; *(*byte)(p) != 0; p = Add(p, 1) {
		cum++
	}
	return cum
}
