package cairo

// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
// #include <cairo-pdf.h>
import "C"

import (
	"image"
	"image/draw"
	"runtime"
	"unsafe"
)

// CreatePNGSurface is a wrapper around cairo_image_surface_create_from_png().
func CreatePNGSurfaceFromPNG(fileName string) (*Surface, error) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))

	surfaceNative := C.cairo_image_surface_create_from_png(cstr)

	status := Status(C.cairo_surface_status(surfaceNative))
	if status != STATUS_SUCCESS {
		return nil, ErrorStatus(status)
	}

	return &Surface{surface: surfaceNative}, nil
}

// CreateImageSurfaceForData is a wrapper around cairo_image_surface_create_for_data().
func CreateImageSurfaceForData(data []byte, format Format, width, height, stride int) *Surface {
	surfaceNative := C.cairo_image_surface_create_for_data((*C.uchar)(unsafe.Pointer(&data[0])),
		C.cairo_format_t(format), C.int(width), C.int(height), C.int(stride))

	status := Status(C.cairo_surface_status(surfaceNative))
	if status != STATUS_SUCCESS {
		panic("cairo_image_surface_create_for_data: " + ErrorStatus(status).Error())
	}

	s := wrapSurface(surfaceNative)
	runtime.SetFinalizer(s, (*Surface).destroy)

	return s
}

// CreateImageSurface is a wrapper around cairo_image_surface_create().
func CreateImageSurface(format Format, width, height int) *Surface {
	surfaceNative := C.cairo_image_surface_create(C.cairo_format_t(format),
		C.int(width), C.int(height))

	status := Status(C.cairo_surface_status(surfaceNative))
	if status != STATUS_SUCCESS {
		panic("cairo_image_surface_create: " + ErrorStatus(status).Error())
	}

	s := wrapSurface(surfaceNative)
	runtime.SetFinalizer(s, (*Surface).destroy)

	return s
}

// CreateSurfaceFromImage is a better wrapper around cairo_image_surface_create_for_data().
func CreateSurfaceFromImage(img image.Image) *Surface {
	var s *Surface

	switch img := img.(type) {
	case *image.RGBA:
		s = CreateImageSurface(FORMAT_ARGB32, img.Rect.Dx(), img.Rect.Dy())

		pix := s.GetData()
		for i := 0; i < len(pix); i += 4 {
			pix[i+0] = img.Pix[i+2]
			pix[i+1] = img.Pix[i+1]
			pix[i+2] = img.Pix[i+0]
			pix[i+3] = img.Pix[i+3]
		}

	case *image.NRGBA:
		s = CreateImageSurface(FORMAT_ARGB32, img.Rect.Dx(), img.Rect.Dy())

		pix := s.GetData()
		for i := 0; i < len(pix); i += 4 {
			alpha := uint16(img.Pix[i+3])
			pix[i+0] = uint8(uint16(img.Pix[i+2]) * alpha / 0xFF)
			pix[i+1] = uint8(uint16(img.Pix[i+1]) * alpha / 0xFF)
			pix[i+2] = uint8(uint16(img.Pix[i+0]) * alpha / 0xFF)
			pix[i+3] = img.Pix[i+3]
		}

	case *image.Alpha:
		s = CreateImageSurface(FORMAT_A8, img.Rect.Dx(), img.Rect.Dy())
		s.Flush()
		pix := s.GetData()
		copy(pix, img.Pix)

	default:
		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, img.Bounds(), img, image.Point{}, draw.Over)
		return CreateSurfaceFromImage(rgba)
	}

	s.MarkDirty()
	return s
}
