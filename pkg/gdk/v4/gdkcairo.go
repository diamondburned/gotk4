// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// CairoDrawFromGL: this is the main way to draw GL content in GTK. It takes a
// render buffer ID (@source_type == RENDERBUFFER) or a texture id (@source_type
// == TEXTURE) and draws it onto @cr with an OVER operation, respecting the
// current clip. The top left corner of the rectangle specified by @x, @y,
// @width and @height will be drawn at the current (0,0) position of the
// cairo_t.
//
// This will work for *all* cairo_t, as long as @surface is realized, but the
// fallback implementation that reads back the pixels from the buffer may be
// used in the general case. In the case of direct drawing to a surface with no
// special effects applied to @cr it will however use a more efficient approach.
//
// For RENDERBUFFER the code will always fall back to software for buffers with
// alpha components, so make sure you use TEXTURE if using alpha.
//
// Calling this may change the current GL context.
func CairoDrawFromGL(cr *cairo.Context, surface Surface, source int, sourceType int, bufferScale int, x int, y int, width int, height int) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkSurface
	var arg3 C.int
	var arg4 C.int
	var arg5 C.int
	var arg6 C.int
	var arg7 C.int
	var arg8 C.int
	var arg9 C.int

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	arg3 = C.int(source)
	arg4 = C.int(sourceType)
	arg5 = C.int(bufferScale)
	arg6 = C.int(x)
	arg7 = C.int(y)
	arg8 = C.int(width)
	arg9 = C.int(height)

	C.gdk_cairo_draw_from_gl(cr, surface, source, sourceType, bufferScale, x, y, width, height)
}

// CairoRectangle adds the given rectangle to the current path of @cr.
func CairoRectangle(cr *cairo.Context, rectangle *Rectangle) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkRectangle

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkRectangle)(unsafe.Pointer(rectangle.Native()))

	C.gdk_cairo_rectangle(cr, rectangle)
}

// CairoRegion adds the given region to the current path of @cr.
func CairoRegion(cr *cairo.Context, region *cairo.Region) {
	var arg1 *C.cairo_t
	var arg2 *C.cairo_region_t

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_cairo_region(cr, region)
}

// CairoRegionCreateFromSurface creates region that describes covers the area
// where the given @surface is more than 50% opaque.
//
// This function takes into account device offsets that might be set with
// cairo_surface_set_device_offset().
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	var arg1 *C.cairo_surface_t

	arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	var cret *C.cairo_region_t
	var goret1 *cairo.Region

	cret = C.gdk_cairo_region_create_from_surface(surface)

	goret1 = cairo.WrapRegion(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *cairo.Region) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// CairoSetSourcePixbuf sets the given pixbuf as the source pattern for @cr.
//
// The pattern has an extend mode of CAIRO_EXTEND_NONE and is aligned so that
// the origin of @pixbuf is @pixbuf_x, @pixbuf_y.
func CairoSetSourcePixbuf(cr *cairo.Context, pixbuf gdkpixbuf.Pixbuf, pixbufX float64, pixbufY float64) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkPixbuf
	var arg3 C.double
	var arg4 C.double

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	arg3 = C.double(pixbufX)
	arg4 = C.double(pixbufY)

	C.gdk_cairo_set_source_pixbuf(cr, pixbuf, pixbufX, pixbufY)
}

// CairoSetSourceRGBA sets the specified RGBA as the source color of @cr.
func CairoSetSourceRGBA(cr *cairo.Context, rgba *RGBA) {
	var arg1 *C.cairo_t
	var arg2 *C.GdkRGBA

	arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	arg2 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	C.gdk_cairo_set_source_rgba(cr, rgba)
}