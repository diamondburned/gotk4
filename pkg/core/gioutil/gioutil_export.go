package gioutil

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #include <glib.h>
// #include <gio/gio.h>
// #include <gio/ginputstream.h>
import "C"

//export goInputStreamRead
func goInputStreamRead(id C.guintptr, buf *C.void, len C.gsize, errOut **C.GError) C.gssize {
	rd, ok := gbox.Get(uintptr(id)).(io.Reader)
	if !ok {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(fmt.Errorf("unknown reader ID %d", id)))
		}
		return -1
	}

	var bytes []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Data = uintptr(unsafe.Pointer(buf))
	header.Len = int(len)
	header.Cap = int(len)

	n, err := rd.Read(bytes)
	if err != nil && !errors.Is(err, io.EOF) {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(err))
		}
		return -1
	}

	return C.gssize(n)
}

//export goOutputStreamWrite
func goOutputStreamWrite(id C.guintptr, buf *C.void, len C.gsize, errOut **C.GError) C.gssize {
	wr, ok := gbox.Get(uintptr(id)).(io.Writer)
	if !ok {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(fmt.Errorf("unknown writer ID %d", id)))
		}
		return -1
	}

	var bytes []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Data = uintptr(unsafe.Pointer(buf))
	header.Len = int(len)
	header.Cap = int(len)

	n, err := wr.Write(bytes)
	if err != nil {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(err))
		}
		return -1
	}

	return C.gssize(n)
}

//export goStreamSeek
func goStreamSeek(id C.guintptr, offset C.goffset, whence C.gint, errOut **C.GError) C.gssize {
	sk, ok := gbox.Get(uintptr(id)).(io.Seeker)
	if !ok {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(fmt.Errorf("unknown seeker ID %d", id)))
		}
		return -1
	}

	n, err := sk.Seek(int64(offset), int(whence))
	if err != nil {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(err))
		}
		return -1
	}

	return C.gssize(n)
}

//export goStreamClose
func goStreamClose(id C.guintptr, errOut **C.GError) C.gboolean {
	rd, ok := gbox.Pop(uintptr(id)).(io.Reader)
	if !ok {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(fmt.Errorf("unknown reader ID %d", id)))
		}
		return C.FALSE
	}

	closer, ok := rd.(io.Closer)
	if !ok {
		return C.TRUE
	}

	if err := closer.Close(); err != nil {
		if errOut != nil {
			*errOut = (*C.GError)(gerror.New(err))
		}
		return C.FALSE
	}

	return C.TRUE
}
