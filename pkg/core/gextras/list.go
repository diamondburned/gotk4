package gextras

// #cgo pkg-config: glib-2.0
// #include <glib.h>
// #include <gmodule.h> // HashTable
import "C"

import (
	"runtime"
	"unsafe"
)

// List wraps a GList, which is a doubly-linked list.
type List[T any] struct {
	*list[T]
}

type list[T any] struct {
	opts   ListOpts[T]
	native *C.GList
}

// ListOpts describes the options when constructing a list. It is mostly for
// internal use.
type ListOpts[T any] struct {
	// Convert must fully copy the pointer's value.
	Convert  func(unsafe.Pointer) T
	FreeData func(unsafe.Pointer)
}

// NewList creates a new List instance from a given GList pointer.
func NewList[T any](ptr unsafe.Pointer, opts ListOpts[T], free bool) *List[T] {
	if ptr == nil {
		return nil
	}
	v := &list[T]{
		native: (*C.GList)(ptr),
		opts:   opts,
	}
	if free {
		runtime.SetFinalizer(v, (*list[T]).free)
	}
	return &List[T]{v}
}

func (l *list[T]) free() {
	if l.opts.FreeData != nil {
		for v := l.native; v != nil; v = v.next {
			l.opts.FreeData(unsafe.Pointer(v.data))
		}
	}
	C.g_list_free(l.native)
	l.native = nil
	runtime.KeepAlive(l)
}

// ForEach iterates over the List. If the callback returns true, then the loop
// is broken. The value is allocated on each call and the callback can take it
// outside.
func (l *List[T]) ForEach(f func(T) (stop bool)) {
	for v := l.native; v != nil; v = v.next {
		if f(l.opts.Convert(unsafe.Pointer(v.data))) {
			break
		}
	}
	runtime.KeepAlive(l)
}

// Find returns the T item that f returns true on or nil.
func (l *List[T]) Find(f func(T) (found bool)) T {
	for v := l.native; v != nil; v = v.next {
		d := l.opts.Convert(unsafe.Pointer(v.data))
		if f(d) {
			runtime.KeepAlive(l)
			return d
		}
	}
	runtime.KeepAlive(l)
	var zero T
	return zero
}

// First returns the first element in the GList.
func (l *List[T]) First() T {
	node := l.native
	for node.prev != nil {
		node = node.prev
	}
	return l.convert(node)
}

// Last returns the last element in the GList.
func (l *List[T]) Last() T {
	node := l.native
	for node.next != nil {
		node = node.next
	}
	return l.convert(node)
}

// Next returns the next element of the GList.
func (l *List[T]) Next() T {
	return l.convert(l.native.next)
}

// Prev returns the previous element of the GList.
func (l *List[T]) Prev() T {
	return l.convert(l.native.prev)
}

// Nth returns the nth element in the GList.
func (l *List[T]) Nth(n int) T {
	if node := l.nth(n); node != nil {
		return l.convert(node)
	}
	var zero T
	return zero
}

// nth returns the nth GList node or nil.
func (l *List[T]) nth(n int) *C.GList {
	node := l.native
	for node != nil && n > 0 {
		node = node.next
		n--
	}
	return node
}

// Length returns the number of eleemnts in the GList. The current node is
// assumed to be the head.
func (l *List[T]) Length() int {
	var i int
	for node := l.native; node != nil; node = node.next {
		i++
	}
	runtime.KeepAlive(l)
	return i
}

// convert wraps around Convert with a KeepAlive.
func (l *List[T]) convert(c *C.GList) T {
	if c == nil {
		var zero T
		return zero
	}
	v := l.opts.Convert(unsafe.Pointer(c.data))
	runtime.KeepAlive(v)
	return v
}

// Convert returns a new Go slice that contains all values. This method might be
// more useful if the user needs to use the list more than once.
func (l *List[T]) Convert() []T {
	s := make([]T, 0, l.Length())
	l.ForEach(func(v T) bool {
		s = append(s, v)
		return false
	})
	return s
}

// ListSize returns the length of the list.
func ListSize(ptr unsafe.Pointer) int {
	return int(C.g_list_length((*C.GList)(ptr)))
}

// MoveList calls f on every value of the given *GList. If rm is true, then the
// GList is freed.
func MoveList(ptr unsafe.Pointer, rm bool, f func(v unsafe.Pointer)) {
	for v := (*C.GList)(ptr); v != nil; v = v.next {
		f(unsafe.Pointer(v.data))
	}

	if rm {
		C.g_list_free((*C.GList)(ptr))
	}
}
