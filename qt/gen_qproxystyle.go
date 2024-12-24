package qt

import (
	"unsafe"
)

type QProxyStyle struct {
	h          uintptr
	isSubclass bool
}

// NewQProxyStyle constructs a new QProxyStyle object.
func NewQProxyStyle() *QProxyStyle {
	ret := newQProxyStyle(QProxyStyle_new())
	ret.isSubclass = true
	return ret
}

// NewQProxyStyle2 constructs a new QProxyStyle object.
func NewQProxyStyle2(key string) *QProxyStyle {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))

	ret := newQProxyStyle(QProxyStyle_new2(key_ms))
	ret.isSubclass = true
	return ret
}

// NewQProxyStyle3 constructs a new QProxyStyle object.
func NewQProxyStyle3(style *QStyle) *QProxyStyle {
	ret := newQProxyStyle(QProxyStyle_new3(style.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QProxyStyle) MetaObject() *QMetaObject {
	return newQMetaObject(QProxyStyle_MetaObject(this.h))
}

func (this *QProxyStyle) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QProxyStyle_Metacast(this.h, param1_Cstring))
}

func QProxyStyle_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QProxyStyle_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProxyStyle) BaseStyle() *QStyle {
	return newQStyle(QProxyStyle_BaseStyle(this.h))
}

func (this *QProxyStyle) SetBaseStyle(style *QStyle) {
	QProxyStyle_SetBaseStyle(this.h, style.cPointer())
}

func (this *QProxyStyle) DrawPrimitive(element PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	QProxyStyle_DrawPrimitive(this.h, element, option.cPointer(), painter.cPointer(), widget.cPointer())
}

func (this *QProxyStyle) DrawControl(element ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	QProxyStyle_DrawControl(this.h, element, option.cPointer(), painter.cPointer(), widget.cPointer())
}

func (this *QProxyStyle) DrawComplexControl(control ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget) {
	QProxyStyle_DrawComplexControl(this.h, control, option.cPointer(), painter.cPointer(), widget.cPointer())
}

func (this *QProxyStyle) DrawItemText(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QProxyStyle_DrawItemText(this.h, painter.cPointer(), rect.cPointer(), (int)(flags), pal.cPointer(), (bool)(enabled), text_ms, (int)(textRole))
}

func (this *QProxyStyle) DrawItemPixmap(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	QProxyStyle_DrawItemPixmap(this.h, painter.cPointer(), rect.cPointer(), (int)(alignment), pixmap.cPointer())
}

func (this *QProxyStyle) SizeFromContents(typeVal ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize {
	_goptr := newQSize(QProxyStyle_SizeFromContents(this.h, typeVal, option.cPointer(), size.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) SubElementRect(element SubElement, option *QStyleOption, widget *QWidget) *QRect {
	_goptr := newQRect(QProxyStyle_SubElementRect(this.h, element, option.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) SubControlRect(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, widget *QWidget) *QRect {
	_goptr := newQRect(QProxyStyle_SubControlRect(this.h, cc, opt.cPointer(), sc, widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) ItemTextRect(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QProxyStyle_ItemTextRect(this.h, fm.cPointer(), r.cPointer(), (int)(flags), (bool)(enabled), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) ItemPixmapRect(r *QRect, flags int, pixmap *QPixmap) *QRect {
	_goptr := newQRect(QProxyStyle_ItemPixmapRect(this.h, r.cPointer(), (int)(flags), pixmap.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) HitTestComplexControl(control ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) SubControl {
	xxxxxxxxx
}

func (this *QProxyStyle) StyleHint(hint StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {
	return (int)(QProxyStyle_StyleHint(this.h, hint, option.cPointer(), widget.cPointer(), returnData.cPointer()))
}

func (this *QProxyStyle) PixelMetric(metric PixelMetric, option *QStyleOption, widget *QWidget) int {
	return (int)(QProxyStyle_PixelMetric(this.h, metric, option.cPointer(), widget.cPointer()))
}

func (this *QProxyStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {
	return (int)(QProxyStyle_LayoutSpacing(this.h, (int)(control1), (int)(control2), (int)(orientation), option.cPointer(), widget.cPointer()))
}

func (this *QProxyStyle) StandardIcon(standardIcon StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon {
	_goptr := newQIcon(QProxyStyle_StandardIcon(this.h, standardIcon, option.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) StandardPixmap(standardPixmap StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	_goptr := newQPixmap(QProxyStyle_StandardPixmap(this.h, standardPixmap, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) GeneratedIconPixmap(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	_goptr := newQPixmap(QProxyStyle_GeneratedIconPixmap(this.h, (int)(iconMode), pixmap.cPointer(), opt.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) StandardPalette() *QPalette {
	_goptr := newQPalette(QProxyStyle_StandardPalette(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProxyStyle) Polish(widget *QWidget) {
	QProxyStyle_Polish(this.h, widget.cPointer())
}

func (this *QProxyStyle) PolishWithPal(pal *QPalette) {
	QProxyStyle_PolishWithPal(this.h, pal.cPointer())
}

func (this *QProxyStyle) PolishWithApp(app *QApplication) {
	QProxyStyle_PolishWithApp(this.h, app.cPointer())
}

func (this *QProxyStyle) Unpolish(widget *QWidget) {
	QProxyStyle_Unpolish(this.h, widget.cPointer())
}

func (this *QProxyStyle) UnpolishWithApp(app *QApplication) {
	QProxyStyle_UnpolishWithApp(this.h, app.cPointer())
}

func QProxyStyle_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProxyStyle_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProxyStyle_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProxyStyle_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProxyStyle) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QProxyStyle_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QProxyStyle) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_MetaObject
func miqt_exec_callback_QProxyStyle_MetaObject(self QProxyStyle, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QProxyStyle) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QProxyStyle_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QProxyStyle) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_Metacast
func miqt_exec_callback_QProxyStyle_Metacast(self QProxyStyle, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
