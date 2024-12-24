package qt

import (
	"unsafe"
)

type QRgba64 struct {
	h          uintptr
	isSubclass bool
}

// NewQRgba64 constructs a new QRgba64 object.
func NewQRgba64() *QRgba64 {

	ret := newQRgba64(QRgba64_new())
	ret.isSubclass = true
	return ret
}

// NewQRgba642 constructs a new QRgba64 object.
func NewQRgba642(param1 *QRgba64) *QRgba64 {

	ret := newQRgba64(QRgba64_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func QRgba64_FromRgba64(c uint64) *QRgba64 {
	_goptr := newQRgba64(QRgba64_FromRgba64((ulonglong)(c)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QRgba64_FromRgba642(red uint16, green uint16, blue uint16, alpha uint16) *QRgba64 {
	_goptr := newQRgba64(QRgba64_FromRgba642((uint16_t)(red), (uint16_t)(green), (uint16_t)(blue), (uint16_t)(alpha)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QRgba64_FromRgba(red byte, green byte, blue byte, alpha byte) *QRgba64 {
	_goptr := newQRgba64(QRgba64_FromRgba((uchar)(red), (uchar)(green), (uchar)(blue), (uchar)(alpha)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QRgba64_FromArgb32(rgb uint) *QRgba64 {
	_goptr := newQRgba64(QRgba64_FromArgb32((uint)(rgb)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRgba64) IsOpaque() bool {
	return (bool)(QRgba64_IsOpaque(this.h))
}

func (this *QRgba64) IsTransparent() bool {
	return (bool)(QRgba64_IsTransparent(this.h))
}

func (this *QRgba64) Red() uint16 {
	return (uint16)(QRgba64_Red(this.h))
}

func (this *QRgba64) Green() uint16 {
	return (uint16)(QRgba64_Green(this.h))
}

func (this *QRgba64) Blue() uint16 {
	return (uint16)(QRgba64_Blue(this.h))
}

func (this *QRgba64) Alpha() uint16 {
	return (uint16)(QRgba64_Alpha(this.h))
}

func (this *QRgba64) SetRed(_red uint16) {
	QRgba64_SetRed(this.h, (uint16_t)(_red))
}

func (this *QRgba64) SetGreen(_green uint16) {
	QRgba64_SetGreen(this.h, (uint16_t)(_green))
}

func (this *QRgba64) SetBlue(_blue uint16) {
	QRgba64_SetBlue(this.h, (uint16_t)(_blue))
}

func (this *QRgba64) SetAlpha(_alpha uint16) {
	QRgba64_SetAlpha(this.h, (uint16_t)(_alpha))
}

func (this *QRgba64) Red8() byte {
	return (byte)(QRgba64_Red8(this.h))
}

func (this *QRgba64) Green8() byte {
	return (byte)(QRgba64_Green8(this.h))
}

func (this *QRgba64) Blue8() byte {
	return (byte)(QRgba64_Blue8(this.h))
}

func (this *QRgba64) Alpha8() byte {
	return (byte)(QRgba64_Alpha8(this.h))
}

func (this *QRgba64) ToArgb32() uint {
	return (uint)(QRgba64_ToArgb32(this.h))
}

func (this *QRgba64) ToRgb16() uint16 {
	return (uint16)(QRgba64_ToRgb16(this.h))
}

func (this *QRgba64) Premultiplied() *QRgba64 {
	_goptr := newQRgba64(QRgba64_Premultiplied(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRgba64) Unpremultiplied() *QRgba64 {
	_goptr := newQRgba64(QRgba64_Unpremultiplied(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRgba64) OperatorAssign(_rgba uint64) {
	QRgba64_OperatorAssign(this.h, (ulonglong)(_rgba))
}
