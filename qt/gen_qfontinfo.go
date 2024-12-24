package qt

import (
	"unsafe"
)

type QFontInfo struct {
	h          uintptr
	isSubclass bool
}

// NewQFontInfo constructs a new QFontInfo object.
func NewQFontInfo(param1 *QFont) *QFontInfo {
	ret := newQFontInfo(QFontInfo_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontInfo2 constructs a new QFontInfo object.
func NewQFontInfo2(param1 *QFontInfo) *QFontInfo {
	ret := newQFontInfo(QFontInfo_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFontInfo) OperatorAssign(param1 *QFontInfo) {
	QFontInfo_OperatorAssign(this.h, param1.cPointer())
}

func (this *QFontInfo) Swap(other *QFontInfo) {
	QFontInfo_Swap(this.h, other.cPointer())
}

func (this *QFontInfo) Family() string {
	var _ms struct_miqt_string = QFontInfo_Family(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontInfo) StyleName() string {
	var _ms struct_miqt_string = QFontInfo_StyleName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontInfo) PixelSize() int {
	return (int)(QFontInfo_PixelSize(this.h))
}

func (this *QFontInfo) PointSize() int {
	return (int)(QFontInfo_PointSize(this.h))
}

func (this *QFontInfo) PointSizeF() float64 {
	return (float64)(QFontInfo_PointSizeF(this.h))
}

func (this *QFontInfo) Italic() bool {
	return (bool)(QFontInfo_Italic(this.h))
}

func (this *QFontInfo) Style() QFont__Style {
	return (QFont__Style)(QFontInfo_Style(this.h))
}

func (this *QFontInfo) Weight() int {
	return (int)(QFontInfo_Weight(this.h))
}

func (this *QFontInfo) Bold() bool {
	return (bool)(QFontInfo_Bold(this.h))
}

func (this *QFontInfo) Underline() bool {
	return (bool)(QFontInfo_Underline(this.h))
}

func (this *QFontInfo) Overline() bool {
	return (bool)(QFontInfo_Overline(this.h))
}

func (this *QFontInfo) StrikeOut() bool {
	return (bool)(QFontInfo_StrikeOut(this.h))
}

func (this *QFontInfo) FixedPitch() bool {
	return (bool)(QFontInfo_FixedPitch(this.h))
}

func (this *QFontInfo) StyleHint() QFont__StyleHint {
	return (QFont__StyleHint)(QFontInfo_StyleHint(this.h))
}

func (this *QFontInfo) VariableAxes() []QFontVariableAxis {
	var _ma struct_miqt_array = QFontInfo_VariableAxes(this.h)
	_ret := make([]QFontVariableAxis, int(_ma.len))
	_outCast := (*[0xffff]*QFontVariableAxis)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFontVariableAxis(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QFontInfo) LegacyWeight() int {
	return (int)(QFontInfo_LegacyWeight(this.h))
}

func (this *QFontInfo) ExactMatch() bool {
	return (bool)(QFontInfo_ExactMatch(this.h))
}
