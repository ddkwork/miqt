package qt

/*

#cgo CFLAGS: -fPIC
#cgo pkg-config: Qt5Widgets
#include "gen_qcolortransform.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QColorTransform struct {
	h *C.QColorTransform
}

func (this *QColorTransform) cPointer() *C.QColorTransform {
	if this == nil {
		return nil
	}
	return this.h
}

func newQColorTransform(h *C.QColorTransform) *QColorTransform {
	if h == nil {
		return nil
	}
	return &QColorTransform{h: h}
}

func newQColorTransform_U(h unsafe.Pointer) *QColorTransform {
	return newQColorTransform((*C.QColorTransform)(h))
}

// NewQColorTransform constructs a new QColorTransform object.
func NewQColorTransform() *QColorTransform {
	ret := C.QColorTransform_new()
	return newQColorTransform(ret)
}

// NewQColorTransform2 constructs a new QColorTransform object.
func NewQColorTransform2(colorTransform *QColorTransform) *QColorTransform {
	ret := C.QColorTransform_new2(colorTransform.cPointer())
	return newQColorTransform(ret)
}

func (this *QColorTransform) OperatorAssign(other *QColorTransform) {
	C.QColorTransform_OperatorAssign(this.h, other.cPointer())
}

func (this *QColorTransform) Swap(other *QColorTransform) {
	C.QColorTransform_Swap(this.h, other.cPointer())
}

func (this *QColorTransform) Map(argb uint) uint {
	ret := C.QColorTransform_Map(this.h, (C.uint)(argb))
	return (uint)(ret)
}

func (this *QColorTransform) MapWithRgba64(rgba64 QRgba64) *QRgba64 {
	ret := C.QColorTransform_MapWithRgba64(this.h, rgba64.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQRgba64(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QRgba64) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QColorTransform) MapWithColor(color *QColor) *QColor {
	ret := C.QColorTransform_MapWithColor(this.h, color.cPointer())
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQColor(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QColor) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QColorTransform) Delete() {
	C.QColorTransform_Delete(this.h)
}