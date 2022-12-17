// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// GTlsDatabase* _gotk4_gio2_TLSBackend_virtual_get_default_database(void* fnptr, GTlsBackend* arg0) {
//   return ((GTlsDatabase* (*)(GTlsBackend*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_TLSBackend_virtual_supports_dtls(void* fnptr, GTlsBackend* arg0) {
//   return ((gboolean (*)(GTlsBackend*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_TLSBackend_virtual_supports_tls(void* fnptr, GTlsBackend* arg0) {
//   return ((gboolean (*)(GTlsBackend*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeTLSBackend = coreglib.Type(C.g_tls_backend_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTLSBackend, F: marshalTLSBackend},
	})
}

// TLSBackend: TLS (Transport Layer Security, aka SSL) and DTLS backend.
//
// TLSBackend wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TLSBackend struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TLSBackend)(nil)
)

// TLSBackender describes TLSBackend's interface methods.
type TLSBackender interface {
	coreglib.Objector

	// CertificateType gets the #GType of backend's Certificate implementation.
	CertificateType() coreglib.Type
	// ClientConnectionType gets the #GType of backend's ClientConnection
	// implementation.
	ClientConnectionType() coreglib.Type
	// DefaultDatabase gets the default Database used to verify TLS connections.
	DefaultDatabase() TLSDatabaser
	// DTLSClientConnectionType gets the #GType of backend’s ClientConnection
	// implementation.
	DTLSClientConnectionType() coreglib.Type
	// DTLSServerConnectionType gets the #GType of backend’s ServerConnection
	// implementation.
	DTLSServerConnectionType() coreglib.Type
	// FileDatabaseType gets the #GType of backend's FileDatabase
	// implementation.
	FileDatabaseType() coreglib.Type
	// ServerConnectionType gets the #GType of backend's ServerConnection
	// implementation.
	ServerConnectionType() coreglib.Type
	// SetDefaultDatabase: set the default Database used to verify TLS
	// connections Any subsequent call to g_tls_backend_get_default_database()
	// will return the database set in this call.
	SetDefaultDatabase(database TLSDatabaser)
	// SupportsDTLS checks if DTLS is supported.
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns FALSE for the
	// default Backend, it means no "real" TLS backend is available.
	SupportsTLS() bool
}

var _ TLSBackender = (*TLSBackend)(nil)

func wrapTLSBackend(obj *coreglib.Object) *TLSBackend {
	return &TLSBackend{
		Object: obj,
	}
}

func marshalTLSBackend(p uintptr) (interface{}, error) {
	return wrapTLSBackend(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CertificateType gets the #GType of backend's Certificate implementation.
//
// The function returns the following values:
//
//    - gType of backend's Certificate implementation.
//
func (backend *TLSBackend) CertificateType() coreglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_certificate_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// ClientConnectionType gets the #GType of backend's ClientConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend's ClientConnection implementation.
//
func (backend *TLSBackend) ClientConnectionType() coreglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_client_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// DefaultDatabase gets the default Database used to verify TLS connections.
//
// The function returns the following values:
//
//    - tlsDatabase: default database, which should be unreffed when done.
//
func (backend *TLSBackend) DefaultDatabase() TLSDatabaser {
	var _arg0 *C.GTlsBackend  // out
	var _cret *C.GTlsDatabase // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_default_database(_arg0)
	runtime.KeepAlive(backend)

	var _tlsDatabase TLSDatabaser // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSDatabaser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TLSDatabaser)
			return ok
		})
		rv, ok := casted.(TLSDatabaser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSDatabaser")
		}
		_tlsDatabase = rv
	}

	return _tlsDatabase
}

// DTLSClientConnectionType gets the #GType of backend’s ClientConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend’s ClientConnection implementation, or G_TYPE_INVALID if
//      this backend doesn’t support DTLS.
//
func (backend *TLSBackend) DTLSClientConnectionType() coreglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_dtls_client_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// DTLSServerConnectionType gets the #GType of backend’s ServerConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend’s ServerConnection implementation, or G_TYPE_INVALID if
//      this backend doesn’t support DTLS.
//
func (backend *TLSBackend) DTLSServerConnectionType() coreglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_dtls_server_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// FileDatabaseType gets the #GType of backend's FileDatabase implementation.
//
// The function returns the following values:
//
//    - gType of backend's FileDatabase implementation.
//
func (backend *TLSBackend) FileDatabaseType() coreglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_file_database_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// ServerConnectionType gets the #GType of backend's ServerConnection
// implementation.
//
// The function returns the following values:
//
//    - gType of backend's ServerConnection implementation.
//
func (backend *TLSBackend) ServerConnectionType() coreglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_get_server_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// SetDefaultDatabase: set the default Database used to verify TLS connections
//
// Any subsequent call to g_tls_backend_get_default_database() will return the
// database set in this call. Existing databases and connections are not
// modified.
//
// Setting a NULL default database will reset to using the system default
// database as if g_tls_backend_set_default_database() had never been called.
//
// The function takes the following parameters:
//
//    - database (optional): Database.
//
func (backend *TLSBackend) SetDefaultDatabase(database TLSDatabaser) {
	var _arg0 *C.GTlsBackend  // out
	var _arg1 *C.GTlsDatabase // out

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))
	if database != nil {
		_arg1 = (*C.GTlsDatabase)(unsafe.Pointer(coreglib.InternObject(database).Native()))
	}

	C.g_tls_backend_set_default_database(_arg0, _arg1)
	runtime.KeepAlive(backend)
	runtime.KeepAlive(database)
}

// SupportsDTLS checks if DTLS is supported. DTLS support may not be available
// even if TLS support is available, and vice-versa.
//
// The function returns the following values:
//
//    - ok: whether DTLS is supported.
//
func (backend *TLSBackend) SupportsDTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_supports_dtls(_arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsTLS checks if TLS is supported; if this returns FALSE for the default
// Backend, it means no "real" TLS backend is available.
//
// The function returns the following values:
//
//    - ok: whether or not TLS is supported.
//
func (backend *TLSBackend) SupportsTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C.g_tls_backend_supports_tls(_arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// defaultDatabase gets the default Database used to verify TLS connections.
//
// The function returns the following values:
//
//    - tlsDatabase: default database, which should be unreffed when done.
//
func (backend *TLSBackend) defaultDatabase() TLSDatabaser {
	gclass := (*C.GTlsBackendInterface)(coreglib.PeekParentClass(backend))
	fnarg := gclass.get_default_database

	var _arg0 *C.GTlsBackend  // out
	var _cret *C.GTlsDatabase // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C._gotk4_gio2_TLSBackend_virtual_get_default_database(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(backend)

	var _tlsDatabase TLSDatabaser // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSDatabaser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TLSDatabaser)
			return ok
		})
		rv, ok := casted.(TLSDatabaser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSDatabaser")
		}
		_tlsDatabase = rv
	}

	return _tlsDatabase
}

// supportsDTLS checks if DTLS is supported. DTLS support may not be available
// even if TLS support is available, and vice-versa.
//
// The function returns the following values:
//
//    - ok: whether DTLS is supported.
//
func (backend *TLSBackend) supportsDTLS() bool {
	gclass := (*C.GTlsBackendInterface)(coreglib.PeekParentClass(backend))
	fnarg := gclass.supports_dtls

	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C._gotk4_gio2_TLSBackend_virtual_supports_dtls(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// supportsTLS checks if TLS is supported; if this returns FALSE for the default
// Backend, it means no "real" TLS backend is available.
//
// The function returns the following values:
//
//    - ok: whether or not TLS is supported.
//
func (backend *TLSBackend) supportsTLS() bool {
	gclass := (*C.GTlsBackendInterface)(coreglib.PeekParentClass(backend))
	fnarg := gclass.supports_tls

	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(coreglib.InternObject(backend).Native()))

	_cret = C._gotk4_gio2_TLSBackend_virtual_supports_tls(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TLSBackendGetDefault gets the default Backend for the system.
//
// The function returns the following values:
//
//    - tlsBackend which will be a dummy object if no TLS backend is available.
//
func TLSBackendGetDefault() *TLSBackend {
	var _cret *C.GTlsBackend // in

	_cret = C.g_tls_backend_get_default()

	var _tlsBackend *TLSBackend // out

	_tlsBackend = wrapTLSBackend(coreglib.Take(unsafe.Pointer(_cret)))

	return _tlsBackend
}

// TLSBackendInterface provides an interface for describing TLS-related types.
//
// An instance of this type is always passed by reference.
type TLSBackendInterface struct {
	*tlsBackendInterface
}

// tlsBackendInterface is the struct that's finalized.
type tlsBackendInterface struct {
	native *C.GTlsBackendInterface
}