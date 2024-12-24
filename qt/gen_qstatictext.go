package qt

import (
	"unsafe"
)

type QStaticText__PerformanceHint int

const (
	QStaticText__ModerateCaching   QStaticText__PerformanceHint = 0
	QStaticText__AggressiveCaching QStaticText__PerformanceHint = 1
)

type QStaticText struct {
	h          uintptr
	isSubclass bool
}

// NewQStaticText constructs a new QStaticText object.
func NewQStaticText() *QStaticText {

	ret := newQStaticText(QStaticText_new())
	ret.isSubclass = true
	return ret
}

// NewQStaticText2 constructs a new QStaticText object.
func NewQStaticText2(text string) *QStaticText {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQStaticText(QStaticText_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQStaticText3 constructs a new QStaticText object.
func NewQStaticText3(other *QStaticText) *QStaticText {

	ret := newQStaticText(QStaticText_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStaticText) OperatorAssign(param1 *QStaticText) {
	QStaticText_OperatorAssign(this.h, param1.cPointer())
}

func (this *QStaticText) Swap(other *QStaticText) {
	QStaticText_Swap(this.h, other.cPointer())
}

func (this *QStaticText) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStaticText_SetText(this.h, text_ms)
}

func (this *QStaticText) Text() string {
	var _ms struct_miqt_string = QStaticText_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStaticText) SetTextFormat(textFormat TextFormat) {
	QStaticText_SetTextFormat(this.h, (int)(textFormat))
}

func (this *QStaticText) TextFormat() TextFormat {
	return (TextFormat)(QStaticText_TextFormat(this.h))
}

func (this *QStaticText) SetTextWidth(textWidth float64) {
	QStaticText_SetTextWidth(this.h, (double)(textWidth))
}

func (this *QStaticText) TextWidth() float64 {
	return (float64)(QStaticText_TextWidth(this.h))
}

func (this *QStaticText) SetTextOption(textOption *QTextOption) {
	QStaticText_SetTextOption(this.h, textOption.cPointer())
}

func (this *QStaticText) TextOption() *QTextOption {
	_goptr := newQTextOption(QStaticText_TextOption(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStaticText) Size() *QSizeF {
	_goptr := newQSizeF(QStaticText_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStaticText) Prepare() {
	QStaticText_Prepare(this.h)
}

func (this *QStaticText) SetPerformanceHint(performanceHint PerformanceHint) {
	QStaticText_SetPerformanceHint(this.h, performanceHint)
}

func (this *QStaticText) PerformanceHint() PerformanceHint {
	xxxxxxxxx
}

func (this *QStaticText) OperatorEqual(param1 *QStaticText) bool {
	return (bool)(QStaticText_OperatorEqual(this.h, param1.cPointer()))
}

func (this *QStaticText) OperatorNotEqual(param1 *QStaticText) bool {
	return (bool)(QStaticText_OperatorNotEqual(this.h, param1.cPointer()))
}

func (this *QStaticText) Prepare1(matrix *QTransform) {
	QStaticText_Prepare1(this.h, matrix.cPointer())
}

func (this *QStaticText) Prepare2(matrix *QTransform, font *QFont) {
	QStaticText_Prepare2(this.h, matrix.cPointer(), font.cPointer())
}
