package ptr

import (
	"unsafe"
)

// Strlen calculates the length of a null-terminated string in the given
// pointer.
func Strlen(str unsafe.Pointer) int {
	var cum int
	for p := str; *(*byte)(p) != nil; p = unsafe.Add(p, 1) {
		cum++
	}
	return cum
}
