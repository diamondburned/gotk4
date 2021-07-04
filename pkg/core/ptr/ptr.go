package ptr

import (
	"unsafe"
)

// Strlen calculates the length of a null-terminated string in the given
// pointer.
func Strlen(str unsafe.Pointer) int {
	var cum int
	for p := str; *(*byte)(p) != 0; p = unsafe.Add(p, 1) {
		cum++
	}
	return cum
}
