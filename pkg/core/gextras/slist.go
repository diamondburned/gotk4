package gextras

// #cgo pkg-config: glib-2.0
// #include <glib.h>
// #include <gmodule.h> // HashTable
import "C"

import (
	"runtime"
	"unsafe"
)

// SList wraps a GSList, which is a singly-linked list.
type SList[T any] struct {
	*slist[T]
}

type slist[T any] struct {
	opts   ListOpts[T]
	native *C.GSList
}

// NewSList creates a new SList instance from a given GSList pointer.
func NewSList[T any](ptr unsafe.Pointer, opts ListOpts[T], free bool) *SList[T] {
	if ptr == nil {
		return nil
	}
	v := &slist[T]{
		opts:   opts,
		native: (*C.GSList)(ptr),
	}
	if free {
		runtime.SetFinalizer(v, (*slist[T]).free)
	}
	return &SList[T]{v}
}

func (l *slist[T]) free() {
	if l.opts.FreeData != nil {
		for v := l.native; v != nil; v = v.next {
			l.opts.FreeData(unsafe.Pointer(v.data))
		}
	}
	C.g_slist_free(l.native)
	l.native = nil
	runtime.KeepAlive(l)
}

// ForEach iterates over the SList. If the callback returns true, then the loop
// is broken. The value is allocated on each call and the callback can take it
// outside.
func (l *SList[T]) ForEach(f func(T) (stop bool)) {
	for v := l.native; v != nil; v = v.next {
		if f(l.opts.Convert(unsafe.Pointer(v.data))) {
			break
		}
	}
	runtime.KeepAlive(l)
}

// Find returns the T item that f returns true on or nil.
func (l *SList[T]) Find(f func(T) (found bool)) T {
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

// Last returns the last element in the GSList.
func (l *SList[T]) Last() T {
	node := l.native
	for node.next != nil {
		node = node.next
	}
	return l.convert(node)
}

// Next returns the next element of the GSList.
func (l *SList[T]) Next() T {
	return l.convert(l.native.next)
}

// Nth returns the nth element in the GSList.
func (l *SList[T]) Nth(n int) T {
	if node := l.nth(n); node != nil {
		return l.convert(node)
	}
	var zero T
	return zero
}

// nth returns the nth GSList node or nil.
func (l *SList[T]) nth(n int) *C.GSList {
	node := l.native
	for node != nil && n > 0 {
		node = node.next
		n--
	}
	return node
}

// Length returns the number of eleemnts in the GSList. The current node is
// assumed to be the head.
func (l *SList[T]) Length() int {
	var i int
	for node := l.native; node != nil; node = node.next {
		i++
	}
	runtime.KeepAlive(l)
	return i
}

// convert wraps around Convert with a KeepAlive.
func (l *SList[T]) convert(c *C.GSList) T {
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
func (l *SList[T]) Convert() []T {
	s := make([]T, 0, l.Length())
	l.ForEach(func(v T) bool {
		s = append(s, v)
		return false
	})
	return s
}

// SListSize returns the length of the singly-linked list.
func SListSize(ptr unsafe.Pointer) int {
	return int(C.g_slist_length((*C.GSList)(ptr)))
}

// MoveSList is similar to MoveList, except it's used for singly-linked lists.
func MoveSList(ptr unsafe.Pointer, rm bool, f func(v unsafe.Pointer)) {
	for v := (*C.GSList)(ptr); v != nil; v = v.next {
		f(unsafe.Pointer(v.data))
	}

	if rm {
		C.g_slist_free((*C.GSList)(ptr))
	}
}
