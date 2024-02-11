package gioutil

// #cgo pkg-config: gio-2.0
// #include "listmodel.h"
import "C"

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/core/slab"
	"github.com/diamondburned/gotk4/pkg/gio/v2"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

var registry slab.Slab

func init() {
	registry.Grow(128)
}

//export gotk4_gboxlist_remove
func gotk4_gboxlist_remove(id uintptr) { registry.Delete(id) }

// ListModelType is a type-safe wrapper around [ListModel] and [ObjectValue].
// For an example, see [ListModel]'s example.
type ListModelType[T any] struct{}

// New creates a new list model of type T.
func (t ListModelType[T]) New() *ListModel[T] {
	return NewListModel[T]()
}

// ObjectValue returns the value of the given object as type T.
// The object must originate from a [ListModel].
func (t ListModelType[T]) ObjectValue(obj glib.Objector) T {
	return ObjectValue[T](obj)
}

// ListModel is a wrapper around an internal ListModel that allows any Go value
// to be used as a list item. Internally, it uses core/gbox to store the values
// in a global registry for later retrieval.
type ListModel[T any] struct {
	*gio.ListModel
}

// NewListModel creates a new list model.
func NewListModel[T any]() *ListModel[T] {
	obj := coreglib.AssumeOwnership(unsafe.Pointer(C.gotk4_gbox_list_new())) // C.Gotk4GboxList
	return &ListModel[T]{
		ListModel: &gio.ListModel{
			Object: obj,
		},
	}
}

func (l *ListModel[T]) native() *C.Gotk4GboxList {
	return (*C.Gotk4GboxList)(unsafe.Pointer(l.Native()))
}

// Append appends a value to the list.
func (l *ListModel[T]) Append(v T) {
	C.gotk4_gbox_list_append(l.native(), C.guintptr(registry.Put(v, false)))
}

// Remove removes the value at the given index.
func (l *ListModel[T]) Remove(index int) {
	C.gotk4_gbox_list_remove(l.native(), C.guint(index))
}

// Splice removes the values in the given range and replaces them with the
// given values.
func (l *ListModel[T]) Splice(position, removals int, values ...T) {
	var idsPtr *C.guintptr
	if len(values) > 0 {
		idsPtr = (*C.guintptr)(unsafe.Pointer(C.g_malloc_n(C.gsize(len(values)+1), C.sizeof_gpointer)))
		defer C.g_free(C.gpointer(idsPtr))

		ids := unsafe.Slice(idsPtr, len(values)+1)
		ids[len(values)] = 0
		for i, v := range values {
			ids[i] = C.guintptr(registry.Put(v, false))
		}
	}

	C.gotk4_gbox_list_splice(l.native(), C.guint(position), C.guint(removals), idsPtr)
}

// Item returns the value at the given index.
func (l *ListModel[T]) Item(index int) T {
	id := uintptr(C.gotk4_gbox_list_get_id(l.native(), C.guint(index)))
	return registry.Get(id).(T)
}

// NItems returns the number of items in the list.
func (l *ListModel[T]) NItems() int {
	return int(l.ListModel.NItems())
}

// AllItems returns an iterator over all values in the list.
func (l *ListModel[T]) AllItems() func(yield func(T) bool) {
	return func(yield func(T) bool) {
		nItems := l.NItems()
		for i := 0; i < nItems; i++ {
			if !yield(l.Item(i)) {
				break
			}
		}
	}
}

// RangeItems returns an iterator over the values in the given range.
// If j is greater than the length of the list, it will be clamped to that.
func (l *ListModel[T]) RangeItems(i, j int) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		nItems := l.NItems()
		for j = min(j, nItems); i < j; i++ {
			if !yield(l.Item(i)) {
				break
			}
		}
	}
}

// ObjectValue returns the value of the given object.
// The object must originate from a [ListModel].
func ObjectValue[T any](obj glib.Objector) T {
	object := glib.BaseObject(obj)
	if object.Type() != coreglib.Type(C.gotk4_gbox_object_get_type()) {
		panic("StringObjectValue: obj must be a *Gotk4GboxObject")
	}

	native := (*C.Gotk4GboxObject)(unsafe.Pointer(object.Native()))
	id := uintptr(C.gotk4_gbox_object_get_id(native))

	return registry.Get(id).(T)
}
