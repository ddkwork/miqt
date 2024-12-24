package qt

import (
	"unsafe"
)

type QColorTransform struct {
	h          uintptr
	isSubclass bool
}

// NewQColorTransform constructs a new QColorTransform object.
func NewQColorTransform() *QColorTransform {

	ret := newQColorTransform(QColorTransform_new())
	ret.isSubclass = true
	return ret
}

// NewQColorTransform2 constructs a new QColorTransform object.
func NewQColorTransform2(colorTransform *QColorTransform) *QColorTransform {

	ret := newQColorTransform(QColorTransform_new2(colorTransform.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QColorTransform) OperatorAssign(other *QColorTransform) {
	QColorTransform_OperatorAssign(this.h, other.cPointer())
}

func (this *QColorTransform) Swap(other *QColorTransform) {
	QColorTransform_Swap(this.h, other.cPointer())
}

func (this *QColorTransform) IsIdentity() bool {
	return (bool)(QColorTransform_IsIdentity(this.h))
}

func (this *QColorTransform) Map(argb uint) uint {
	return (uint)(QColorTransform_Map(this.h, (uint)(argb)))
}

func (this *QColorTransform) MapWithRgba64(rgba64 QRgba64) *QRgba64 {
	_goptr := newQRgba64(QColorTransform_MapWithRgba64(this.h, rgba64.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorTransform) MapWithColor(color *QColor) *QColor {
	_goptr := newQColor(QColorTransform_MapWithColor(this.h, color.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
