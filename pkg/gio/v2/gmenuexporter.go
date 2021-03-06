// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// ExportMenuModel exports menu on connection at object_path.
//
// The implemented D-Bus API should be considered private. It is subject to
// change in the future.
//
// An object path can only have one menu model exported on it. If this
// constraint is violated, the export will fail and 0 will be returned (with
// error set accordingly).
//
// You can unexport the menu model using g_dbus_connection_unexport_menu_model()
// with the return value of this function.
//
// The function takes the following parameters:
//
//    - objectPath d-Bus object path.
//    - menu: Model.
//
// The function returns the following values:
//
//    - guint: ID of the export (never zero), or 0 in case of failure.
//
func (connection *DBusConnection) ExportMenuModel(objectPath string, menu MenuModeller) (uint, error) {
	var _arg0 *C.GDBusConnection // out
	var _arg1 *C.gchar           // out
	var _arg2 *C.GMenuModel      // out
	var _cret C.guint            // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GDBusConnection)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	_cret = C.g_dbus_connection_export_menu_model(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(objectPath)
	runtime.KeepAlive(menu)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// UnexportMenuModel reverses the effect of a previous call to
// g_dbus_connection_export_menu_model().
//
// It is an error to call this function with an ID that wasn't returned from
// g_dbus_connection_export_menu_model() or to call it with the same ID more
// than once.
//
// The function takes the following parameters:
//
//    - exportId: ID from g_dbus_connection_export_menu_model().
//
func (connection *DBusConnection) UnexportMenuModel(exportId uint) {
	var _arg0 *C.GDBusConnection // out
	var _arg1 C.guint            // out

	_arg0 = (*C.GDBusConnection)(unsafe.Pointer(coreglib.InternObject(connection).Native()))
	_arg1 = C.guint(exportId)

	C.g_dbus_connection_unexport_menu_model(_arg0, _arg1)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(exportId)
}
