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
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeTLSServerConnection = coreglib.Type(C.g_tls_server_connection_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTLSServerConnection, F: marshalTLSServerConnection},
	})
}

// TLSServerConnectionOverrider contains methods that are overridable.
type TLSServerConnectionOverrider interface {
}

// TLSServerConnection is the server-side subclass of Connection, representing a
// server-side TLS connection.
//
// TLSServerConnection wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TLSServerConnection struct {
	_ [0]func() // equal guard
	TLSConnection
}

var (
	_ TLSConnectioner = (*TLSServerConnection)(nil)
)

// TLSServerConnectioner describes TLSServerConnection's interface methods.
type TLSServerConnectioner interface {
	coreglib.Objector

	baseTLSServerConnection() *TLSServerConnection
}

var _ TLSServerConnectioner = (*TLSServerConnection)(nil)

func ifaceInitTLSServerConnectioner(gifacePtr, data C.gpointer) {
}

func wrapTLSServerConnection(obj *coreglib.Object) *TLSServerConnection {
	return &TLSServerConnection{
		TLSConnection: TLSConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTLSServerConnection(p uintptr) (interface{}, error) {
	return wrapTLSServerConnection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *TLSServerConnection) baseTLSServerConnection() *TLSServerConnection {
	return v
}

// BaseTLSServerConnection returns the underlying base object.
func BaseTLSServerConnection(obj TLSServerConnectioner) *TLSServerConnection {
	return obj.baseTLSServerConnection()
}

// NewTLSServerConnection creates a new ServerConnection wrapping base_io_stream
// (which must have pollable input and output streams).
//
// See the documentation for Connection:base-io-stream for restrictions on when
// application code can run operations on the base_io_stream after this function
// has returned.
//
// The function takes the following parameters:
//
//    - baseIoStream to wrap.
//    - certificate (optional): default server certificate, or NULL.
//
// The function returns the following values:
//
//    - tlsServerConnection: new ServerConnection, or NULL on error.
//
func NewTLSServerConnection(baseIoStream IOStreamer, certificate TLSCertificater) (*TLSServerConnection, error) {
	var _arg1 *C.GIOStream       // out
	var _arg2 *C.GTlsCertificate // out
	var _cret *C.GIOStream       // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(coreglib.InternObject(baseIoStream).Native()))
	if certificate != nil {
		_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(certificate).Native()))
	}

	_cret = C.g_tls_server_connection_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(certificate)

	var _tlsServerConnection *TLSServerConnection // out
	var _goerr error                              // out

	_tlsServerConnection = wrapTLSServerConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsServerConnection, _goerr
}