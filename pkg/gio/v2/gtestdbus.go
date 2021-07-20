// Code generated by girgen. DO NOT EDIT.

package gio

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

// TestDBusUnset: unset DISPLAY and DBUS_SESSION_BUS_ADDRESS env variables to
// ensure the test won't use user's session bus.
//
// This is useful for unit tests that want to verify behaviour when no session
// bus is running. It is not necessary to call this if unit test already calls
// g_test_dbus_up() before acquiring the session bus.
func TestDBusUnset() {
	C.g_test_dbus_unset()
}