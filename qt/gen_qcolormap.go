package qt

import (
	"unsafe"
)

type QColormap__Mode int

const (
	QColormap__Direct  QColormap__Mode = 0
	QColormap__Indexed QColormap__Mode = 1
	QColormap__Gray    QColormap__Mode = 2
)

type QColormap struct {
	h          uintptr
	isSubclass bool
}

// NewQColormap constructs a new QColormap object.
func NewQColormap(colormap *QColormap) *QColormap {
	g := newQColormap(QColormap_new(colormap.cPointer()))
	g.isSubclass = true
	return g
}

func QColormap_Initialize() {
	QColormap_Initialize()
}

func QColormap_Cleanup() {
	QColormap_Cleanup()
}

func QColormap_Instance() *QColormap {
	_goptr := newQColormap(QColormap_Instance())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColormap) OperatorAssign(colormap *QColormap) {
	QColormap_OperatorAssign(this.h, colormap.cPointer())
}

func (this *QColormap) Mode() Mode {
	xxxxxxxxx
}

func (this *QColormap) Depth() int {
	return (int)(QColormap_Depth(this.h))
}

func (this *QColormap) Size() int {
	return (int)(QColormap_Size(this.h))
}

func (this *QColormap) Pixel(color *QColor) uint {
	return (uint)(QColormap_Pixel(this.h, color.cPointer()))
}

func (this *QColormap) ColorAt(pixel uint) *QColor {
	_goptr := newQColor(QColormap_ColorAt(this.h, (uint)(pixel)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColormap) Colormap() []QColor {
	var _ma struct_miqt_array = QColormap_Colormap(this.h)
	_ret := make([]QColor, int(_ma.len))
	_outCast := (*[0xffff]*QColor)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQColor(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QColormap_Instance1(screen int) *QColormap {
	_goptr := newQColormap(QColormap_Instance1((int)(screen)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
