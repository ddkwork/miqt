package qt

import (
	"unsafe"
)

type QTextLayout__GlyphRunRetrievalFlag uint16

const (
	QTextLayout__RetrieveGlyphIndexes   QTextLayout__GlyphRunRetrievalFlag = 1
	QTextLayout__RetrieveGlyphPositions QTextLayout__GlyphRunRetrievalFlag = 2
	QTextLayout__RetrieveStringIndexes  QTextLayout__GlyphRunRetrievalFlag = 4
	QTextLayout__RetrieveString         QTextLayout__GlyphRunRetrievalFlag = 8
	QTextLayout__DefaultRetrievalFlags  QTextLayout__GlyphRunRetrievalFlag = 3
	QTextLayout__RetrieveAll            QTextLayout__GlyphRunRetrievalFlag = 65535
)

type QTextLayout__CursorMode int

const (
	QTextLayout__SkipCharacters QTextLayout__CursorMode = 0
	QTextLayout__SkipWords      QTextLayout__CursorMode = 1
)

type QTextLine__Edge int

const (
	QTextLine__Leading  QTextLine__Edge = 0
	QTextLine__Trailing QTextLine__Edge = 1
)

type QTextLine__CursorPosition int

const (
	QTextLine__CursorBetweenCharacters QTextLine__CursorPosition = 0
	QTextLine__CursorOnCharacter       QTextLine__CursorPosition = 1
)

type QTextInlineObject struct {
	h          uintptr
	isSubclass bool
}

// NewQTextInlineObject constructs a new QTextInlineObject object.
func NewQTextInlineObject() *QTextInlineObject {

	ret := newQTextInlineObject(QTextInlineObject_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextInlineObject) IsValid() bool {
	return (bool)(QTextInlineObject_IsValid(this.h))
}

func (this *QTextInlineObject) Rect() *QRectF {
	_goptr := newQRectF(QTextInlineObject_Rect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextInlineObject) Width() float64 {
	return (float64)(QTextInlineObject_Width(this.h))
}

func (this *QTextInlineObject) Ascent() float64 {
	return (float64)(QTextInlineObject_Ascent(this.h))
}

func (this *QTextInlineObject) Descent() float64 {
	return (float64)(QTextInlineObject_Descent(this.h))
}

func (this *QTextInlineObject) Height() float64 {
	return (float64)(QTextInlineObject_Height(this.h))
}

func (this *QTextInlineObject) TextDirection() LayoutDirection {
	return (LayoutDirection)(QTextInlineObject_TextDirection(this.h))
}

func (this *QTextInlineObject) SetWidth(w float64) {
	QTextInlineObject_SetWidth(this.h, (double)(w))
}

func (this *QTextInlineObject) SetAscent(a float64) {
	QTextInlineObject_SetAscent(this.h, (double)(a))
}

func (this *QTextInlineObject) SetDescent(d float64) {
	QTextInlineObject_SetDescent(this.h, (double)(d))
}

func (this *QTextInlineObject) TextPosition() int {
	return (int)(QTextInlineObject_TextPosition(this.h))
}

func (this *QTextInlineObject) FormatIndex() int {
	return (int)(QTextInlineObject_FormatIndex(this.h))
}

func (this *QTextInlineObject) Format() *QTextFormat {
	_goptr := newQTextFormat(QTextInlineObject_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QTextLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQTextLayout constructs a new QTextLayout object.
func NewQTextLayout() *QTextLayout {

	ret := newQTextLayout(QTextLayout_new())
	ret.isSubclass = true
	return ret
}

// NewQTextLayout2 constructs a new QTextLayout object.
func NewQTextLayout2(text string) *QTextLayout {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTextLayout(QTextLayout_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQTextLayout3 constructs a new QTextLayout object.
func NewQTextLayout3(text string, font *QFont) *QTextLayout {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTextLayout(QTextLayout_new3(text_ms, font.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextLayout4 constructs a new QTextLayout object.
func NewQTextLayout4(b *QTextBlock) *QTextLayout {

	ret := newQTextLayout(QTextLayout_new4(b.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextLayout5 constructs a new QTextLayout object.
func NewQTextLayout5(text string, font *QFont, paintdevice *QPaintDevice) *QTextLayout {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTextLayout(QTextLayout_new5(text_ms, font.cPointer(), paintdevice.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextLayout) SetFont(f *QFont) {
	QTextLayout_SetFont(this.h, f.cPointer())
}

func (this *QTextLayout) Font() *QFont {
	_goptr := newQFont(QTextLayout_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLayout) SetRawFont(rawFont *QRawFont) {
	QTextLayout_SetRawFont(this.h, rawFont.cPointer())
}

func (this *QTextLayout) SetText(stringVal string) {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	QTextLayout_SetText(this.h, stringVal_ms)
}

func (this *QTextLayout) Text() string {
	var _ms struct_miqt_string = QTextLayout_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextLayout) SetTextOption(option *QTextOption) {
	QTextLayout_SetTextOption(this.h, option.cPointer())
}

func (this *QTextLayout) TextOption() *QTextOption {
	return newQTextOption(QTextLayout_TextOption(this.h))
}

func (this *QTextLayout) SetPreeditArea(position int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextLayout_SetPreeditArea(this.h, (int)(position), text_ms)
}

func (this *QTextLayout) PreeditAreaPosition() int {
	return (int)(QTextLayout_PreeditAreaPosition(this.h))
}

func (this *QTextLayout) PreeditAreaText() string {
	var _ms struct_miqt_string = QTextLayout_PreeditAreaText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextLayout) SetFormats(overrides []FormatRange) {
	overrides_CArray := (*[0xffff]FormatRange)(malloc(size_t(8 * len(overrides))))
	defer free(unsafe.Pointer(overrides_CArray))
	for i := range overrides {
		overrides_CArray[i] = overrides[i]
	}
	overrides_ma := struct_miqt_array{len: size_t(len(overrides)), data: unsafe.Pointer(overrides_CArray)}
	QTextLayout_SetFormats(this.h, overrides_ma)
}

func (this *QTextLayout) Formats() []FormatRange {
	var _ma struct_miqt_array = QTextLayout_Formats(this.h)
	_ret := make([]FormatRange, int(_ma.len))
	_outCast := (*[0xffff]FormatRange)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QTextLayout) ClearFormats() {
	QTextLayout_ClearFormats(this.h)
}

func (this *QTextLayout) SetCacheEnabled(enable bool) {
	QTextLayout_SetCacheEnabled(this.h, (bool)(enable))
}

func (this *QTextLayout) CacheEnabled() bool {
	return (bool)(QTextLayout_CacheEnabled(this.h))
}

func (this *QTextLayout) SetCursorMoveStyle(style CursorMoveStyle) {
	QTextLayout_SetCursorMoveStyle(this.h, (int)(style))
}

func (this *QTextLayout) CursorMoveStyle() CursorMoveStyle {
	return (CursorMoveStyle)(QTextLayout_CursorMoveStyle(this.h))
}

func (this *QTextLayout) BeginLayout() {
	QTextLayout_BeginLayout(this.h)
}

func (this *QTextLayout) EndLayout() {
	QTextLayout_EndLayout(this.h)
}

func (this *QTextLayout) ClearLayout() {
	QTextLayout_ClearLayout(this.h)
}

func (this *QTextLayout) CreateLine() *QTextLine {
	_goptr := newQTextLine(QTextLayout_CreateLine(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLayout) LineCount() int {
	return (int)(QTextLayout_LineCount(this.h))
}

func (this *QTextLayout) LineAt(i int) *QTextLine {
	_goptr := newQTextLine(QTextLayout_LineAt(this.h, (int)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLayout) LineForTextPosition(pos int) *QTextLine {
	_goptr := newQTextLine(QTextLayout_LineForTextPosition(this.h, (int)(pos)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLayout) IsValidCursorPosition(pos int) bool {
	return (bool)(QTextLayout_IsValidCursorPosition(this.h, (int)(pos)))
}

func (this *QTextLayout) NextCursorPosition(oldPos int) int {
	return (int)(QTextLayout_NextCursorPosition(this.h, (int)(oldPos)))
}

func (this *QTextLayout) PreviousCursorPosition(oldPos int) int {
	return (int)(QTextLayout_PreviousCursorPosition(this.h, (int)(oldPos)))
}

func (this *QTextLayout) LeftCursorPosition(oldPos int) int {
	return (int)(QTextLayout_LeftCursorPosition(this.h, (int)(oldPos)))
}

func (this *QTextLayout) RightCursorPosition(oldPos int) int {
	return (int)(QTextLayout_RightCursorPosition(this.h, (int)(oldPos)))
}

func (this *QTextLayout) Draw(p *QPainter, pos *QPointF) {
	QTextLayout_Draw(this.h, p.cPointer(), pos.cPointer())
}

func (this *QTextLayout) DrawCursor(p *QPainter, pos *QPointF, cursorPosition int) {
	QTextLayout_DrawCursor(this.h, p.cPointer(), pos.cPointer(), (int)(cursorPosition))
}

func (this *QTextLayout) DrawCursor2(p *QPainter, pos *QPointF, cursorPosition int, width int) {
	QTextLayout_DrawCursor2(this.h, p.cPointer(), pos.cPointer(), (int)(cursorPosition), (int)(width))
}

func (this *QTextLayout) Position() *QPointF {
	_goptr := newQPointF(QTextLayout_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLayout) SetPosition(p *QPointF) {
	QTextLayout_SetPosition(this.h, p.cPointer())
}

func (this *QTextLayout) BoundingRect() *QRectF {
	_goptr := newQRectF(QTextLayout_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLayout) MinimumWidth() float64 {
	return (float64)(QTextLayout_MinimumWidth(this.h))
}

func (this *QTextLayout) MaximumWidth() float64 {
	return (float64)(QTextLayout_MaximumWidth(this.h))
}

func (this *QTextLayout) GlyphRuns(from int, length int, flags GlyphRunRetrievalFlags) []QGlyphRun {
	var _ma struct_miqt_array = QTextLayout_GlyphRuns(this.h, (int)(from), (int)(length), flags)
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextLayout) GlyphRuns2() []QGlyphRun {
	var _ma struct_miqt_array = QTextLayout_GlyphRuns2(this.h)
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextLayout) SetFlags(flags int) {
	QTextLayout_SetFlags(this.h, (int)(flags))
}

func (this *QTextLayout) NextCursorPosition2(oldPos int, mode CursorMode) int {
	return (int)(QTextLayout_NextCursorPosition2(this.h, (int)(oldPos), mode))
}

func (this *QTextLayout) PreviousCursorPosition2(oldPos int, mode CursorMode) int {
	return (int)(QTextLayout_PreviousCursorPosition2(this.h, (int)(oldPos), mode))
}

func (this *QTextLayout) Draw3(p *QPainter, pos *QPointF, selections []FormatRange) {
	selections_CArray := (*[0xffff]FormatRange)(malloc(size_t(8 * len(selections))))
	defer free(unsafe.Pointer(selections_CArray))
	for i := range selections {
		selections_CArray[i] = selections[i]
	}
	selections_ma := struct_miqt_array{len: size_t(len(selections)), data: unsafe.Pointer(selections_CArray)}
	QTextLayout_Draw3(this.h, p.cPointer(), pos.cPointer(), selections_ma)
}

func (this *QTextLayout) Draw4(p *QPainter, pos *QPointF, selections []FormatRange, clip *QRectF) {
	selections_CArray := (*[0xffff]FormatRange)(malloc(size_t(8 * len(selections))))
	defer free(unsafe.Pointer(selections_CArray))
	for i := range selections {
		selections_CArray[i] = selections[i]
	}
	selections_ma := struct_miqt_array{len: size_t(len(selections)), data: unsafe.Pointer(selections_CArray)}
	QTextLayout_Draw4(this.h, p.cPointer(), pos.cPointer(), selections_ma, clip.cPointer())
}

func (this *QTextLayout) GlyphRuns1(from int) []QGlyphRun {
	var _ma struct_miqt_array = QTextLayout_GlyphRuns1(this.h, (int)(from))
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextLayout) GlyphRuns22(from int, length int) []QGlyphRun {
	var _ma struct_miqt_array = QTextLayout_GlyphRuns22(this.h, (int)(from), (int)(length))
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

type QTextLine struct {
	h          uintptr
	isSubclass bool
}

// NewQTextLine constructs a new QTextLine object.
func NewQTextLine() *QTextLine {

	ret := newQTextLine(QTextLine_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextLine) IsValid() bool {
	return (bool)(QTextLine_IsValid(this.h))
}

func (this *QTextLine) Rect() *QRectF {
	_goptr := newQRectF(QTextLine_Rect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLine) X() float64 {
	return (float64)(QTextLine_X(this.h))
}

func (this *QTextLine) Y() float64 {
	return (float64)(QTextLine_Y(this.h))
}

func (this *QTextLine) Width() float64 {
	return (float64)(QTextLine_Width(this.h))
}

func (this *QTextLine) Ascent() float64 {
	return (float64)(QTextLine_Ascent(this.h))
}

func (this *QTextLine) Descent() float64 {
	return (float64)(QTextLine_Descent(this.h))
}

func (this *QTextLine) Height() float64 {
	return (float64)(QTextLine_Height(this.h))
}

func (this *QTextLine) Leading() float64 {
	return (float64)(QTextLine_Leading(this.h))
}

func (this *QTextLine) SetLeadingIncluded(included bool) {
	QTextLine_SetLeadingIncluded(this.h, (bool)(included))
}

func (this *QTextLine) LeadingIncluded() bool {
	return (bool)(QTextLine_LeadingIncluded(this.h))
}

func (this *QTextLine) NaturalTextWidth() float64 {
	return (float64)(QTextLine_NaturalTextWidth(this.h))
}

func (this *QTextLine) HorizontalAdvance() float64 {
	return (float64)(QTextLine_HorizontalAdvance(this.h))
}

func (this *QTextLine) NaturalTextRect() *QRectF {
	_goptr := newQRectF(QTextLine_NaturalTextRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLine) CursorToX(cursorPos *int) float64 {
	return (float64)(QTextLine_CursorToX(this.h, (*int)(unsafe.Pointer(cursorPos))))
}

func (this *QTextLine) CursorToXWithCursorPos(cursorPos int) float64 {
	return (float64)(QTextLine_CursorToXWithCursorPos(this.h, (int)(cursorPos)))
}

func (this *QTextLine) XToCursor(x float64) int {
	return (int)(QTextLine_XToCursor(this.h, (double)(x)))
}

func (this *QTextLine) SetLineWidth(width float64) {
	QTextLine_SetLineWidth(this.h, (double)(width))
}

func (this *QTextLine) SetNumColumns(columns int) {
	QTextLine_SetNumColumns(this.h, (int)(columns))
}

func (this *QTextLine) SetNumColumns2(columns int, alignmentWidth float64) {
	QTextLine_SetNumColumns2(this.h, (int)(columns), (double)(alignmentWidth))
}

func (this *QTextLine) SetPosition(pos *QPointF) {
	QTextLine_SetPosition(this.h, pos.cPointer())
}

func (this *QTextLine) Position() *QPointF {
	_goptr := newQPointF(QTextLine_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextLine) TextStart() int {
	return (int)(QTextLine_TextStart(this.h))
}

func (this *QTextLine) TextLength() int {
	return (int)(QTextLine_TextLength(this.h))
}

func (this *QTextLine) LineNumber() int {
	return (int)(QTextLine_LineNumber(this.h))
}

func (this *QTextLine) Draw(painter *QPainter, position *QPointF) {
	QTextLine_Draw(this.h, painter.cPointer(), position.cPointer())
}

func (this *QTextLine) GlyphRuns(from int, length int, flags GlyphRunRetrievalFlag) []QGlyphRun {
	var _ma struct_miqt_array = QTextLine_GlyphRuns(this.h, (int)(from), (int)(length), (int)(flags))
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextLine) GlyphRuns2() []QGlyphRun {
	var _ma struct_miqt_array = QTextLine_GlyphRuns2(this.h)
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextLine) CursorToX2(cursorPos *int, edge Edge) float64 {
	return (float64)(QTextLine_CursorToX2(this.h, (*int)(unsafe.Pointer(cursorPos)), edge))
}

func (this *QTextLine) CursorToX22(cursorPos int, edge Edge) float64 {
	return (float64)(QTextLine_CursorToX22(this.h, (int)(cursorPos), edge))
}

func (this *QTextLine) XToCursor2(x float64, param2 CursorPosition) int {
	return (int)(QTextLine_XToCursor2(this.h, (double)(x), param2))
}

func (this *QTextLine) GlyphRuns1(from int) []QGlyphRun {
	var _ma struct_miqt_array = QTextLine_GlyphRuns1(this.h, (int)(from))
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextLine) GlyphRuns22(from int, length int) []QGlyphRun {
	var _ma struct_miqt_array = QTextLine_GlyphRuns22(this.h, (int)(from), (int)(length))
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

type QTextLayout__FormatRange struct {
	h          uintptr
	isSubclass bool
}

// NewQTextLayout__FormatRange constructs a new QTextLayout::FormatRange object.
func NewQTextLayout__FormatRange() *QTextLayout__FormatRange {

	ret := newQTextLayout__FormatRange(QTextLayout__FormatRange_new())
	ret.isSubclass = true
	return ret
}
