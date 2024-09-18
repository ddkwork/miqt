package qt

/*

#include "gen_qcolormap.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QColormap__Mode int

const (
	QColormap__Mode__Direct  QColormap__Mode = 0
	QColormap__Mode__Indexed QColormap__Mode = 1
	QColormap__Mode__Gray    QColormap__Mode = 2
)

type QColormap struct {
	h *C.QColormap
}

func (this *QColormap) cPointer() *C.QColormap {
	if this == nil {
		return nil
	}
	return this.h
}

func newQColormap(h *C.QColormap) *QColormap {
	if h == nil {
		return nil
	}
	return &QColormap{h: h}
}

func newQColormap_U(h unsafe.Pointer) *QColormap {
	return newQColormap((*C.QColormap)(h))
}

// NewQColormap constructs a new QColormap object.
func NewQColormap(colormap *QColormap) *QColormap {
	ret := C.QColormap_new(colormap.cPointer())
	return newQColormap(ret)
}

func QColormap_Initialize() {
	C.QColormap_Initialize()
}

func QColormap_Cleanup() {
	C.QColormap_Cleanup()
}

func QColormap_Instance() *QColormap {
	_ret := C.QColormap_Instance()
	_goptr := newQColormap(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColormap) OperatorAssign(colormap *QColormap) {
	C.QColormap_OperatorAssign(this.h, colormap.cPointer())
}

func (this *QColormap) Mode() QColormap__Mode {
	return (QColormap__Mode)(C.QColormap_Mode(this.h))
}

func (this *QColormap) Depth() int {
	return (int)(C.QColormap_Depth(this.h))
}

func (this *QColormap) Size() int {
	return (int)(C.QColormap_Size(this.h))
}

func (this *QColormap) Pixel(color *QColor) uint {
	return (uint)(C.QColormap_Pixel(this.h, color.cPointer()))
}

func (this *QColormap) ColorAt(pixel uint) *QColor {
	_ret := C.QColormap_ColorAt(this.h, (C.uint)(pixel))
	_goptr := newQColor(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColormap) Colormap() []QColor {
	var _ma *C.struct_miqt_array = C.QColormap_Colormap(this.h)
	_ret := make([]QColor, int(_ma.len))
	_outCast := (*[0xffff]*C.QColor)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_vv_ret := _outCast[i]
		_vv_goptr := newQColor(_vv_ret)
		_vv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_vv_goptr
	}
	C.free(unsafe.Pointer(_ma))
	return _ret
}

func QColormap_Instance1(screen int) *QColormap {
	_ret := C.QColormap_Instance1((C.int)(screen))
	_goptr := newQColormap(_ret)
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QColormap) Delete() {
	C.QColormap_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QColormap) GoGC() {
	runtime.SetFinalizer(this, func(this *QColormap) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
