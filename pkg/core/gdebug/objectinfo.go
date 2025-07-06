package gdebug

// #cgo pkg-config: gobject-2.0
// #include <glib-object.h>
//
// const gchar *gotk4_object_type_name(gpointer obj) {
//   return G_OBJECT_TYPE_NAME(obj);
// };
import "C"
import (
	"fmt"
	"log/slog"
	"unsafe"
)

// ObjectInfo returns a slog.Attr with the GObject's pointer and type.
func ObjectInfo(obj unsafe.Pointer) slog.Attr {
	return slog.Group(
		"gobject",
		slog.String("ptr", fmt.Sprintf("%p", obj)),
		slog.String("type", C.GoString(C.gotk4_object_type_name(C.gpointer(obj)))),
		slog.Int("refs", objRefCount(obj)))
}

func objRefCount(obj unsafe.Pointer) int {
	return int(C.g_atomic_int_get((*C.gint)(unsafe.Pointer(&(*C.GObject)(obj).ref_count))))
}
