package qt

import (
	"unsafe"
)

type QFontMetrics struct {
	h          uintptr
	isSubclass bool
}

// NewQFontMetrics constructs a new QFontMetrics object.
func NewQFontMetrics(param1 *QFont) *QFontMetrics {
	ret := newQFontMetrics(QFontMetrics_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontMetrics2 constructs a new QFontMetrics object.
func NewQFontMetrics2(font *QFont, pd *QPaintDevice) *QFontMetrics {
	ret := newQFontMetrics(QFontMetrics_new2(font.cPointer(), pd.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontMetrics3 constructs a new QFontMetrics object.
func NewQFontMetrics3(param1 *QFontMetrics) *QFontMetrics {
	ret := newQFontMetrics(QFontMetrics_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFontMetrics) OperatorAssign(param1 *QFontMetrics) {
	QFontMetrics_OperatorAssign(this.h, param1.cPointer())
}

func (this *QFontMetrics) Swap(other *QFontMetrics) {
	QFontMetrics_Swap(this.h, other.cPointer())
}

func (this *QFontMetrics) Ascent() int {
	return (int)(QFontMetrics_Ascent(this.h))
}

func (this *QFontMetrics) CapHeight() int {
	return (int)(QFontMetrics_CapHeight(this.h))
}

func (this *QFontMetrics) Descent() int {
	return (int)(QFontMetrics_Descent(this.h))
}

func (this *QFontMetrics) Height() int {
	return (int)(QFontMetrics_Height(this.h))
}

func (this *QFontMetrics) Leading() int {
	return (int)(QFontMetrics_Leading(this.h))
}

func (this *QFontMetrics) LineSpacing() int {
	return (int)(QFontMetrics_LineSpacing(this.h))
}

func (this *QFontMetrics) MinLeftBearing() int {
	return (int)(QFontMetrics_MinLeftBearing(this.h))
}

func (this *QFontMetrics) MinRightBearing() int {
	return (int)(QFontMetrics_MinRightBearing(this.h))
}

func (this *QFontMetrics) MaxWidth() int {
	return (int)(QFontMetrics_MaxWidth(this.h))
}

func (this *QFontMetrics) XHeight() int {
	return (int)(QFontMetrics_XHeight(this.h))
}

func (this *QFontMetrics) AverageCharWidth() int {
	return (int)(QFontMetrics_AverageCharWidth(this.h))
}

func (this *QFontMetrics) InFont(param1 QChar) bool {
	return (bool)(QFontMetrics_InFont(this.h, param1.cPointer()))
}

func (this *QFontMetrics) InFontUcs4(ucs4 uint) bool {
	return (bool)(QFontMetrics_InFontUcs4(this.h, (uint)(ucs4)))
}

func (this *QFontMetrics) LeftBearing(param1 QChar) int {
	return (int)(QFontMetrics_LeftBearing(this.h, param1.cPointer()))
}

func (this *QFontMetrics) RightBearing(param1 QChar) int {
	return (int)(QFontMetrics_RightBearing(this.h, param1.cPointer()))
}

func (this *QFontMetrics) HorizontalAdvance(param1 string) int {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	return (int)(QFontMetrics_HorizontalAdvance(this.h, param1_ms))
}

func (this *QFontMetrics) HorizontalAdvance2(param1 string, textOption *QTextOption) int {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	return (int)(QFontMetrics_HorizontalAdvance2(this.h, param1_ms, textOption.cPointer()))
}

func (this *QFontMetrics) HorizontalAdvanceWithQChar(param1 QChar) int {
	return (int)(QFontMetrics_HorizontalAdvanceWithQChar(this.h, param1.cPointer()))
}

func (this *QFontMetrics) BoundingRect(param1 QChar) *QRect {
	_goptr := newQRect(QFontMetrics_BoundingRect(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) BoundingRectWithText(text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRectWithText(this.h, text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) BoundingRect2(text string, textOption *QTextOption) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRect2(this.h, text_ms, textOption.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) BoundingRect3(r *QRect, flags int, text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRect3(this.h, r.cPointer(), (int)(flags), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) BoundingRect4(x int, y int, w int, h int, flags int, text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRect4(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(flags), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) Size(flags int, str string) *QSize {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	_goptr := newQSize(QFontMetrics_Size(this.h, (int)(flags), str_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) TightBoundingRect(text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_TightBoundingRect(this.h, text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) TightBoundingRect2(text string, textOption *QTextOption) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_TightBoundingRect2(this.h, text_ms, textOption.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) ElidedText(text string, mode TextElideMode, width int) string {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QFontMetrics_ElidedText(this.h, text_ms, (int)(mode), (int)(width))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontMetrics) UnderlinePos() int {
	return (int)(QFontMetrics_UnderlinePos(this.h))
}

func (this *QFontMetrics) OverlinePos() int {
	return (int)(QFontMetrics_OverlinePos(this.h))
}

func (this *QFontMetrics) StrikeOutPos() int {
	return (int)(QFontMetrics_StrikeOutPos(this.h))
}

func (this *QFontMetrics) LineWidth() int {
	return (int)(QFontMetrics_LineWidth(this.h))
}

func (this *QFontMetrics) FontDpi() float64 {
	return (float64)(QFontMetrics_FontDpi(this.h))
}

func (this *QFontMetrics) OperatorEqual(other *QFontMetrics) bool {
	return (bool)(QFontMetrics_OperatorEqual(this.h, other.cPointer()))
}

func (this *QFontMetrics) OperatorNotEqual(other *QFontMetrics) bool {
	return (bool)(QFontMetrics_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QFontMetrics) HorizontalAdvance22(param1 string, lenVal int) int {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	return (int)(QFontMetrics_HorizontalAdvance22(this.h, param1_ms, (int)(lenVal)))
}

func (this *QFontMetrics) BoundingRect42(r *QRect, flags int, text string, tabstops int) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRect42(this.h, r.cPointer(), (int)(flags), text_ms, (int)(tabstops)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) BoundingRect5(r *QRect, flags int, text string, tabstops int, tabarray *int) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRect5(this.h, r.cPointer(), (int)(flags), text_ms, (int)(tabstops), (*int)(unsafe.Pointer(tabarray))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) BoundingRect7(x int, y int, w int, h int, flags int, text string, tabstops int) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRect7(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(flags), text_ms, (int)(tabstops)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) BoundingRect8(x int, y int, w int, h int, flags int, text string, tabstops int, tabarray *int) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QFontMetrics_BoundingRect8(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(flags), text_ms, (int)(tabstops), (*int)(unsafe.Pointer(tabarray))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) Size3(flags int, str string, tabstops int) *QSize {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	_goptr := newQSize(QFontMetrics_Size3(this.h, (int)(flags), str_ms, (int)(tabstops)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) Size4(flags int, str string, tabstops int, tabarray *int) *QSize {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	_goptr := newQSize(QFontMetrics_Size4(this.h, (int)(flags), str_ms, (int)(tabstops), (*int)(unsafe.Pointer(tabarray))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetrics) ElidedText4(text string, mode TextElideMode, width int, flags int) string {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QFontMetrics_ElidedText4(this.h, text_ms, (int)(mode), (int)(width), (int)(flags))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QFontMetricsF struct {
	h          uintptr
	isSubclass bool
}

// NewQFontMetricsF constructs a new QFontMetricsF object.
func NewQFontMetricsF(font *QFont) *QFontMetricsF {
	ret := newQFontMetricsF(QFontMetricsF_new(font.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontMetricsF2 constructs a new QFontMetricsF object.
func NewQFontMetricsF2(font *QFont, pd *QPaintDevice) *QFontMetricsF {
	ret := newQFontMetricsF(QFontMetricsF_new2(font.cPointer(), pd.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontMetricsF3 constructs a new QFontMetricsF object.
func NewQFontMetricsF3(param1 *QFontMetrics) *QFontMetricsF {
	ret := newQFontMetricsF(QFontMetricsF_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontMetricsF4 constructs a new QFontMetricsF object.
func NewQFontMetricsF4(param1 *QFontMetricsF) *QFontMetricsF {
	ret := newQFontMetricsF(QFontMetricsF_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFontMetricsF) OperatorAssign(param1 *QFontMetricsF) {
	QFontMetricsF_OperatorAssign(this.h, param1.cPointer())
}

func (this *QFontMetricsF) OperatorAssignWithQFontMetrics(param1 *QFontMetrics) {
	QFontMetricsF_OperatorAssignWithQFontMetrics(this.h, param1.cPointer())
}

func (this *QFontMetricsF) Swap(other *QFontMetricsF) {
	QFontMetricsF_Swap(this.h, other.cPointer())
}

func (this *QFontMetricsF) Ascent() float64 {
	return (float64)(QFontMetricsF_Ascent(this.h))
}

func (this *QFontMetricsF) CapHeight() float64 {
	return (float64)(QFontMetricsF_CapHeight(this.h))
}

func (this *QFontMetricsF) Descent() float64 {
	return (float64)(QFontMetricsF_Descent(this.h))
}

func (this *QFontMetricsF) Height() float64 {
	return (float64)(QFontMetricsF_Height(this.h))
}

func (this *QFontMetricsF) Leading() float64 {
	return (float64)(QFontMetricsF_Leading(this.h))
}

func (this *QFontMetricsF) LineSpacing() float64 {
	return (float64)(QFontMetricsF_LineSpacing(this.h))
}

func (this *QFontMetricsF) MinLeftBearing() float64 {
	return (float64)(QFontMetricsF_MinLeftBearing(this.h))
}

func (this *QFontMetricsF) MinRightBearing() float64 {
	return (float64)(QFontMetricsF_MinRightBearing(this.h))
}

func (this *QFontMetricsF) MaxWidth() float64 {
	return (float64)(QFontMetricsF_MaxWidth(this.h))
}

func (this *QFontMetricsF) XHeight() float64 {
	return (float64)(QFontMetricsF_XHeight(this.h))
}

func (this *QFontMetricsF) AverageCharWidth() float64 {
	return (float64)(QFontMetricsF_AverageCharWidth(this.h))
}

func (this *QFontMetricsF) InFont(param1 QChar) bool {
	return (bool)(QFontMetricsF_InFont(this.h, param1.cPointer()))
}

func (this *QFontMetricsF) InFontUcs4(ucs4 uint) bool {
	return (bool)(QFontMetricsF_InFontUcs4(this.h, (uint)(ucs4)))
}

func (this *QFontMetricsF) LeftBearing(param1 QChar) float64 {
	return (float64)(QFontMetricsF_LeftBearing(this.h, param1.cPointer()))
}

func (this *QFontMetricsF) RightBearing(param1 QChar) float64 {
	return (float64)(QFontMetricsF_RightBearing(this.h, param1.cPointer()))
}

func (this *QFontMetricsF) HorizontalAdvance(stringVal string) float64 {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	return (float64)(QFontMetricsF_HorizontalAdvance(this.h, stringVal_ms))
}

func (this *QFontMetricsF) HorizontalAdvanceWithQChar(param1 QChar) float64 {
	return (float64)(QFontMetricsF_HorizontalAdvanceWithQChar(this.h, param1.cPointer()))
}

func (this *QFontMetricsF) HorizontalAdvance2(stringVal string, textOption *QTextOption) float64 {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	return (float64)(QFontMetricsF_HorizontalAdvance2(this.h, stringVal_ms, textOption.cPointer()))
}

func (this *QFontMetricsF) BoundingRect(stringVal string) *QRectF {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQRectF(QFontMetricsF_BoundingRect(this.h, stringVal_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) BoundingRect2(text string, textOption *QTextOption) *QRectF {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRectF(QFontMetricsF_BoundingRect2(this.h, text_ms, textOption.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) BoundingRectWithQChar(param1 QChar) *QRectF {
	_goptr := newQRectF(QFontMetricsF_BoundingRectWithQChar(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) BoundingRect3(r *QRectF, flags int, stringVal string) *QRectF {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQRectF(QFontMetricsF_BoundingRect3(this.h, r.cPointer(), (int)(flags), stringVal_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) Size(flags int, str string) *QSizeF {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	_goptr := newQSizeF(QFontMetricsF_Size(this.h, (int)(flags), str_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) TightBoundingRect(text string) *QRectF {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRectF(QFontMetricsF_TightBoundingRect(this.h, text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) TightBoundingRect2(text string, textOption *QTextOption) *QRectF {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRectF(QFontMetricsF_TightBoundingRect2(this.h, text_ms, textOption.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) ElidedText(text string, mode TextElideMode, width float64) string {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QFontMetricsF_ElidedText(this.h, text_ms, (int)(mode), (double)(width))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontMetricsF) UnderlinePos() float64 {
	return (float64)(QFontMetricsF_UnderlinePos(this.h))
}

func (this *QFontMetricsF) OverlinePos() float64 {
	return (float64)(QFontMetricsF_OverlinePos(this.h))
}

func (this *QFontMetricsF) StrikeOutPos() float64 {
	return (float64)(QFontMetricsF_StrikeOutPos(this.h))
}

func (this *QFontMetricsF) LineWidth() float64 {
	return (float64)(QFontMetricsF_LineWidth(this.h))
}

func (this *QFontMetricsF) FontDpi() float64 {
	return (float64)(QFontMetricsF_FontDpi(this.h))
}

func (this *QFontMetricsF) OperatorEqual(other *QFontMetricsF) bool {
	return (bool)(QFontMetricsF_OperatorEqual(this.h, other.cPointer()))
}

func (this *QFontMetricsF) OperatorNotEqual(other *QFontMetricsF) bool {
	return (bool)(QFontMetricsF_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QFontMetricsF) HorizontalAdvance22(stringVal string, length int) float64 {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	return (float64)(QFontMetricsF_HorizontalAdvance22(this.h, stringVal_ms, (int)(length)))
}

func (this *QFontMetricsF) BoundingRect4(r *QRectF, flags int, stringVal string, tabstops int) *QRectF {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQRectF(QFontMetricsF_BoundingRect4(this.h, r.cPointer(), (int)(flags), stringVal_ms, (int)(tabstops)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) BoundingRect5(r *QRectF, flags int, stringVal string, tabstops int, tabarray *int) *QRectF {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQRectF(QFontMetricsF_BoundingRect5(this.h, r.cPointer(), (int)(flags), stringVal_ms, (int)(tabstops), (*int)(unsafe.Pointer(tabarray))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) Size3(flags int, str string, tabstops int) *QSizeF {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	_goptr := newQSizeF(QFontMetricsF_Size3(this.h, (int)(flags), str_ms, (int)(tabstops)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) Size4(flags int, str string, tabstops int, tabarray *int) *QSizeF {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	_goptr := newQSizeF(QFontMetricsF_Size4(this.h, (int)(flags), str_ms, (int)(tabstops), (*int)(unsafe.Pointer(tabarray))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontMetricsF) ElidedText4(text string, mode TextElideMode, width float64, flags int) string {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QFontMetricsF_ElidedText4(this.h, text_ms, (int)(mode), (double)(width), (int)(flags))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
