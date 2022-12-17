// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeProxyAddress = coreglib.Type(C.g_proxy_address_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeProxyAddress, F: marshalProxyAddress},
	})
}

// ProxyAddressOverrides contains methods that are overridable.
type ProxyAddressOverrides struct {
}

func defaultProxyAddressOverrides(v *ProxyAddress) ProxyAddressOverrides {
	return ProxyAddressOverrides{}
}

// ProxyAddress: support for proxied SocketAddress.
type ProxyAddress struct {
	_ [0]func() // equal guard
	InetSocketAddress
}

var (
	_ SocketAddresser = (*ProxyAddress)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ProxyAddress, *ProxyAddressClass, ProxyAddressOverrides](
		GTypeProxyAddress,
		initProxyAddressClass,
		wrapProxyAddress,
		defaultProxyAddressOverrides,
	)
}

func initProxyAddressClass(gclass unsafe.Pointer, overrides ProxyAddressOverrides, classInitFunc func(*ProxyAddressClass)) {
	if classInitFunc != nil {
		class := (*ProxyAddressClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapProxyAddress(obj *coreglib.Object) *ProxyAddress {
	return &ProxyAddress{
		InetSocketAddress: InetSocketAddress{
			SocketAddress: SocketAddress{
				Object: obj,
				SocketConnectable: SocketConnectable{
					Object: obj,
				},
			},
		},
	}
}

func marshalProxyAddress(p uintptr) (interface{}, error) {
	return wrapProxyAddress(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewProxyAddress creates a new Address for inetaddr with protocol that should
// tunnel through dest_hostname and dest_port.
//
// (Note that this method doesn't set the Address:uri or
// Address:destination-protocol fields; use g_object_new() directly if you want
// to set those.).
//
// The function takes the following parameters:
//
//    - inetaddr: proxy server Address.
//    - port: proxy server port.
//    - protocol: proxy protocol to support, in lower case (e.g. socks, http).
//    - destHostname: destination hostname the proxy should tunnel to.
//    - destPort: destination port to tunnel to.
//    - username (optional) to authenticate to the proxy server (or NULL).
//    - password (optional) to authenticate to the proxy server (or NULL).
//
// The function returns the following values:
//
//    - proxyAddress: new Address.
//
func NewProxyAddress(inetaddr *InetAddress, port uint16, protocol, destHostname string, destPort uint16, username, password string) *ProxyAddress {
	var _arg1 *C.GInetAddress   // out
	var _arg2 C.guint16         // out
	var _arg3 *C.gchar          // out
	var _arg4 *C.gchar          // out
	var _arg5 C.guint16         // out
	var _arg6 *C.gchar          // out
	var _arg7 *C.gchar          // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.GInetAddress)(unsafe.Pointer(coreglib.InternObject(inetaddr).Native()))
	_arg2 = C.guint16(port)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(destHostname)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.guint16(destPort)
	if username != "" {
		_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(username)))
		defer C.free(unsafe.Pointer(_arg6))
	}
	if password != "" {
		_arg7 = (*C.gchar)(unsafe.Pointer(C.CString(password)))
		defer C.free(unsafe.Pointer(_arg7))
	}

	_cret = C.g_proxy_address_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(inetaddr)
	runtime.KeepAlive(port)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(destHostname)
	runtime.KeepAlive(destPort)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)

	var _proxyAddress *ProxyAddress // out

	_proxyAddress = wrapProxyAddress(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _proxyAddress
}

// DestinationHostname gets proxy's destination hostname; that is, the name of
// the host that will be connected to via the proxy, not the name of the proxy
// itself.
//
// The function returns the following values:
//
//    - utf8 proxy's destination hostname.
//
func (proxy *ProxyAddress) DestinationHostname() string {
	var _arg0 *C.GProxyAddress // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_cret = C.g_proxy_address_get_destination_hostname(_arg0)
	runtime.KeepAlive(proxy)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// DestinationPort gets proxy's destination port; that is, the port on the
// destination host that will be connected to via the proxy, not the port number
// of the proxy itself.
//
// The function returns the following values:
//
//    - guint16 proxy's destination port.
//
func (proxy *ProxyAddress) DestinationPort() uint16 {
	var _arg0 *C.GProxyAddress // out
	var _cret C.guint16        // in

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_cret = C.g_proxy_address_get_destination_port(_arg0)
	runtime.KeepAlive(proxy)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// DestinationProtocol gets the protocol that is being spoken to the destination
// server; eg, "http" or "ftp".
//
// The function returns the following values:
//
//    - utf8 proxy's destination protocol.
//
func (proxy *ProxyAddress) DestinationProtocol() string {
	var _arg0 *C.GProxyAddress // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_cret = C.g_proxy_address_get_destination_protocol(_arg0)
	runtime.KeepAlive(proxy)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Password gets proxy's password.
//
// The function returns the following values:
//
//    - utf8 (optional) proxy's password.
//
func (proxy *ProxyAddress) Password() string {
	var _arg0 *C.GProxyAddress // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_cret = C.g_proxy_address_get_password(_arg0)
	runtime.KeepAlive(proxy)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Protocol gets proxy's protocol. eg, "socks" or "http".
//
// The function returns the following values:
//
//    - utf8 proxy's protocol.
//
func (proxy *ProxyAddress) Protocol() string {
	var _arg0 *C.GProxyAddress // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_cret = C.g_proxy_address_get_protocol(_arg0)
	runtime.KeepAlive(proxy)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// URI gets the proxy URI that proxy was constructed from.
//
// The function returns the following values:
//
//    - utf8 (optional) proxy's URI, or NULL if unknown.
//
func (proxy *ProxyAddress) URI() string {
	var _arg0 *C.GProxyAddress // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_cret = C.g_proxy_address_get_uri(_arg0)
	runtime.KeepAlive(proxy)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Username gets proxy's username.
//
// The function returns the following values:
//
//    - utf8 (optional) proxy's username.
//
func (proxy *ProxyAddress) Username() string {
	var _arg0 *C.GProxyAddress // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GProxyAddress)(unsafe.Pointer(coreglib.InternObject(proxy).Native()))

	_cret = C.g_proxy_address_get_username(_arg0)
	runtime.KeepAlive(proxy)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ProxyAddressClass class structure for Address.
//
// An instance of this type is always passed by reference.
type ProxyAddressClass struct {
	*proxyAddressClass
}

// proxyAddressClass is the struct that's finalized.
type proxyAddressClass struct {
	native *C.GProxyAddressClass
}

func (p *ProxyAddressClass) ParentClass() *InetSocketAddressClass {
	valptr := &p.native.parent_class
	var _v *InetSocketAddressClass // out
	_v = (*InetSocketAddressClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}