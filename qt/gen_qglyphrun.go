package qt

import (
	"unsafe"
)

type QGlyphRun__GlyphRunFlag int

const (
	QGlyphRun__Overline      QGlyphRun__GlyphRunFlag = 1
	QGlyphRun__Underline     QGlyphRun__GlyphRunFlag = 2
	QGlyphRun__StrikeOut     QGlyphRun__GlyphRunFlag = 4
	QGlyphRun__RightToLeft   QGlyphRun__GlyphRunFlag = 8
	QGlyphRun__SplitLigature QGlyphRun__GlyphRunFlag = 16
)

type QGlyphRun struct {
	h          uintptr
	isSubclass bool
}

// NewQGlyphRun constructs a new QGlyphRun object.
func NewQGlyphRun() *QGlyphRun {

	ret := newQGlyphRun(QGlyphRun_new())
	ret.isSubclass = true
	return ret
}

// NewQGlyphRun2 constructs a new QGlyphRun object.
func NewQGlyphRun2(other *QGlyphRun) *QGlyphRun {

	ret := newQGlyphRun(QGlyphRun_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGlyphRun) OperatorAssign(other *QGlyphRun) {
	QGlyphRun_OperatorAssign(this.h, other.cPointer())
}

func (this *QGlyphRun) Swap(other *QGlyphRun) {
	QGlyphRun_Swap(this.h, other.cPointer())
}

func (this *QGlyphRun) RawFont() *QRawFont {
	_goptr := newQRawFont(QGlyphRun_RawFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGlyphRun) SetRawFont(rawFont *QRawFont) {
	QGlyphRun_SetRawFont(this.h, rawFont.cPointer())
}

func (this *QGlyphRun) SetRawData(glyphIndexArray *uint, glyphPositionArray *QPointF, size int) {
	QGlyphRun_SetRawData(this.h, (*uint)(unsafe.Pointer(glyphIndexArray)), glyphPositionArray.cPointer(), (int)(size))
}

func (this *QGlyphRun) GlyphIndexes() []uint {
	var _ma struct_miqt_array = QGlyphRun_GlyphIndexes(this.h)
	_ret := make([]uint, int(_ma.len))
	_outCast := (*[0xffff]uint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (uint)(_outCast[i])
	}
	return _ret
}

func (this *QGlyphRun) SetGlyphIndexes(glyphIndexes []uint) {
	glyphIndexes_CArray := (*[0xffff]uint)(malloc(size_t(8 * len(glyphIndexes))))
	defer free(unsafe.Pointer(glyphIndexes_CArray))
	for i := range glyphIndexes {
		glyphIndexes_CArray[i] = (uint)(glyphIndexes[i])
	}
	glyphIndexes_ma := struct_miqt_array{len: size_t(len(glyphIndexes)), data: unsafe.Pointer(glyphIndexes_CArray)}
	QGlyphRun_SetGlyphIndexes(this.h, glyphIndexes_ma)
}

func (this *QGlyphRun) Positions() []QPointF {
	var _ma struct_miqt_array = QGlyphRun_Positions(this.h)
	_ret := make([]QPointF, int(_ma.len))
	_outCast := (*[0xffff]*QPointF)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQPointF(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QGlyphRun) SetPositions(positions []QPointF) {
	positions_CArray := (*[0xffff]*QPointF)(malloc(size_t(8 * len(positions))))
	defer free(unsafe.Pointer(positions_CArray))
	for i := range positions {
		positions_CArray[i] = positions[i].cPointer()
	}
	positions_ma := struct_miqt_array{len: size_t(len(positions)), data: unsafe.Pointer(positions_CArray)}
	QGlyphRun_SetPositions(this.h, positions_ma)
}

func (this *QGlyphRun) Clear() {
	QGlyphRun_Clear(this.h)
}

func (this *QGlyphRun) OperatorEqual(other *QGlyphRun) bool {
	return (bool)(QGlyphRun_OperatorEqual(this.h, other.cPointer()))
}

func (this *QGlyphRun) OperatorNotEqual(other *QGlyphRun) bool {
	return (bool)(QGlyphRun_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QGlyphRun) SetOverline(overline bool) {
	QGlyphRun_SetOverline(this.h, (bool)(overline))
}

func (this *QGlyphRun) Overline() bool {
	return (bool)(QGlyphRun_Overline(this.h))
}

func (this *QGlyphRun) SetUnderline(underline bool) {
	QGlyphRun_SetUnderline(this.h, (bool)(underline))
}

func (this *QGlyphRun) Underline() bool {
	return (bool)(QGlyphRun_Underline(this.h))
}

func (this *QGlyphRun) SetStrikeOut(strikeOut bool) {
	QGlyphRun_SetStrikeOut(this.h, (bool)(strikeOut))
}

func (this *QGlyphRun) StrikeOut() bool {
	return (bool)(QGlyphRun_StrikeOut(this.h))
}

func (this *QGlyphRun) SetRightToLeft(on bool) {
	QGlyphRun_SetRightToLeft(this.h, (bool)(on))
}

func (this *QGlyphRun) IsRightToLeft() bool {
	return (bool)(QGlyphRun_IsRightToLeft(this.h))
}

func (this *QGlyphRun) SetFlag(flag GlyphRunFlag) {
	QGlyphRun_SetFlag(this.h, flag)
}

func (this *QGlyphRun) SetFlags(flags GlyphRunFlags) {
	QGlyphRun_SetFlags(this.h, flags)
}

func (this *QGlyphRun) Flags() GlyphRunFlags {
	xxxxxxxxx
}

func (this *QGlyphRun) SetBoundingRect(boundingRect *QRectF) {
	QGlyphRun_SetBoundingRect(this.h, boundingRect.cPointer())
}

func (this *QGlyphRun) BoundingRect() *QRectF {
	_goptr := newQRectF(QGlyphRun_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGlyphRun) StringIndexes() []int64 {
	var _ma struct_miqt_array = QGlyphRun_StringIndexes(this.h)
	_ret := make([]int64, int(_ma.len))
	_outCast := (*[0xffff]ptrdiff_t)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int64)(_outCast[i])
	}
	return _ret
}

func (this *QGlyphRun) SetStringIndexes(stringIndexes []int64) {
	stringIndexes_CArray := (*[0xffff]ptrdiff_t)(malloc(size_t(8 * len(stringIndexes))))
	defer free(unsafe.Pointer(stringIndexes_CArray))
	for i := range stringIndexes {
		stringIndexes_CArray[i] = (ptrdiff_t)(stringIndexes[i])
	}
	stringIndexes_ma := struct_miqt_array{len: size_t(len(stringIndexes)), data: unsafe.Pointer(stringIndexes_CArray)}
	QGlyphRun_SetStringIndexes(this.h, stringIndexes_ma)
}

func (this *QGlyphRun) SetSourceString(sourceString string) {
	sourceString_ms := struct_miqt_string{}
	sourceString_ms.data = CString(sourceString)
	sourceString_ms.len = size_t(len(sourceString))
	defer free(unsafe.Pointer(sourceString_ms.data))
	QGlyphRun_SetSourceString(this.h, sourceString_ms)
}

func (this *QGlyphRun) SourceString() string {
	var _ms struct_miqt_string = QGlyphRun_SourceString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGlyphRun) IsEmpty() bool {
	return (bool)(QGlyphRun_IsEmpty(this.h))
}

func (this *QGlyphRun) SetFlag2(flag GlyphRunFlag, enabled bool) {
	QGlyphRun_SetFlag2(this.h, flag, (bool)(enabled))
}
