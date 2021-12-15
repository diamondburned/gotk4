package gextras

// #cgo pkg-config: glib-2.0
// #include <glib.h>
// #include <gmodule.h> // HashTable
import "C"
import (
	"runtime"
	"unsafe"
)

// HashTable wraps around a GHashTable.
type HashTable[K comparable, V any] struct {
	*hashTable[K, V]
}

type hashTable[K comparable, V any] struct {
	opts   HashTableOpts[K, V]
	native *C.GHashTable
}

// TODO: these opts only allow C to Go conversion, so we can't do HashTable
// lookups without converting all keys.

// HashTableOpts describes the options when constructing a HashTable. It is
// mostly for internal use. Each Convert function must return a full copy of the
// given pointer's value.
type HashTableOpts[K comparable, V any] struct {
	ConvertKey   func(unsafe.Pointer) K
	ConvertValue func(unsafe.Pointer) V
	HashKey      func(K, func(unsafe.Pointer))
	FreeData     func(key, value unsafe.Pointer)
}

// NewHashTable creates a new HashTable instance from a given GHashTable
// pointer.
func NewHashTable[K comparable, V any](ptr unsafe.Pointer, opts HashTableOpts[K, V], free bool) *HashTable[K, V] {
	if ptr == nil {
		return nil
	}
	v := &hashTable[K, V]{
		opts:   opts,
		native: (*C.GHashTable)(ptr),
	}
	if free {
		runtime.SetFinalizer(v, (*hashTable[K, V]).free)
	}
	return &HashTable[K, V]{v}
}

func (h *hashTable[K, V]) free() {
	if h.opts.FreeData != nil {
		var iter C.GHashTableIter
		C.g_hash_table_iter_init(&iter, h.native)

		var k, v C.gpointer
		for C.g_hash_table_iter_next(&iter, &k, &v) != 0 {
			h.opts.FreeData(unsafe.Pointer(k), unsafe.Pointer(v))
		}
	}
	C.g_hash_table_unref((*C.GHashTable)(h.native))
	h.native = nil
	runtime.KeepAlive(h)
}

// ForEach iterates over the HashTable. If the callback returns true, then the
// loop is broken. The key and value will be allocated on each call and the
// callback can take them outside.
func (h *HashTable[K, V]) ForEach(f func(K, V) (stop bool)) {
	var iter C.GHashTableIter
	C.g_hash_table_iter_init(&iter, h.native)

	var k, v C.gpointer
	for C.g_hash_table_iter_next(&iter, &k, &v) != 0 {
		gk := h.opts.ConvertKey(unsafe.Pointer(k))
		gv := h.opts.ConvertValue(unsafe.Pointer(v))
		if f(gk, gv) {
			break
		}
	}
	runtime.KeepAlive(h)
}

// Find returns the T item that f returns true on or nil.
func (h *HashTable[K, V]) Find(f func(K, V) (found bool)) (K, V) {
	var key K
	var value V
	h.ForEach(func(k K, v V) bool {
		if f(k, v) {
			key = k
			value = v
			return true
		}
		return false
	})
	return key, value
}

// Lookup looks up the given key in the hash table. Note that not all HashTables
// can be looked up in constant time.
func (h *HashTable[K, V]) Lookup(key K) (V, bool) {
	var value V
	var found bool

	if h.opts.HashKey == nil {
		// Slow path.
		h.ForEach(func(k K, v V) bool {
			value = v
			found = key == k
			return found
		})
		return value, found
	}

	h.opts.HashKey(key, func(ckey unsafe.Pointer) {
		var cvalue C.gpointer
		found = C.g_hash_table_lookup_extended(h.native, C.gconstpointer(ckey), nil, &cvalue) != C.FALSE
		if found {
			value = h.opts.ConvertValue(unsafe.Pointer(cvalue))
		}
	})

	return value, found
}

// Length returns the size of the GHashTable.
func (h *HashTable[K, V]) Length() int {
	return int(C.g_hash_table_size(h.native))
}

// Convert returns a new Go map that contains all of HashTable's values. This
// method is useful if the Go program intends to do complicated operations on
// the hash table, such as storing it or passing it around. The keys and values
// will be allocated once, so it is likely cheaper than calling ForEach multiple
// times.
func (h *HashTable[K, V]) Convert() map[K]V {
	m := make(map[K]V, h.Length())

	var iter C.GHashTableIter
	C.g_hash_table_iter_init(&iter, h.native)

	var k, v C.gpointer
	for C.g_hash_table_iter_next(&iter, &k, &v) != 0 {
		gk := h.opts.ConvertKey(unsafe.Pointer(k))
		gv := h.opts.ConvertValue(unsafe.Pointer(v))
		m[gk] = gv

		if h.opts.FreeData != nil {
			h.opts.FreeData(unsafe.Pointer(k), unsafe.Pointer(v))
		}
	}

	return m
}

// MoveHashTable calls f on every value of the given *GHashTable and frees each
// element in the process if rm is true.
func MoveHashTable(ptr unsafe.Pointer, rm bool, f func(k, v unsafe.Pointer)) {
	var k, v C.gpointer
	var iter C.GHashTableIter
	C.g_hash_table_iter_init(&iter, (*C.GHashTable)(ptr))

	for C.g_hash_table_iter_next(&iter, &k, &v) != 0 {
		f(unsafe.Pointer(k), unsafe.Pointer(v))
	}

	if rm {
		C.g_hash_table_unref((*C.GHashTable)(ptr))
	}
}
