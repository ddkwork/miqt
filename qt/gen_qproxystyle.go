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

func (this *QProxyStyle) callVirtualBase_DrawPrimitive(element PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget) {

	QProxyStyle_virtualbase_DrawPrimitive(unsafe.Pointer(this.h), element, option.cPointer(), painter.cPointer(), widget.cPointer())

}
func (this *QProxyStyle) OnDrawPrimitive(slot func(super func(element PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget), element PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_DrawPrimitive(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawPrimitive
func miqt_exec_callback_QProxyStyle_DrawPrimitive(self QProxyStyle, cb intptr_t, element PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(element PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget), element PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(option)

	slotval3 := newQPainter(painter)

	slotval4 := newQWidget(widget)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawPrimitive, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_DrawControl(element ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget) {

	QProxyStyle_virtualbase_DrawControl(unsafe.Pointer(this.h), element, option.cPointer(), painter.cPointer(), widget.cPointer())

}
func (this *QProxyStyle) OnDrawControl(slot func(super func(element ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget), element ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_DrawControl(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawControl
func miqt_exec_callback_QProxyStyle_DrawControl(self QProxyStyle, cb intptr_t, element ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(element ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget), element ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(option)

	slotval3 := newQPainter(painter)

	slotval4 := newQWidget(widget)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawControl, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_DrawComplexControl(control ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget) {

	QProxyStyle_virtualbase_DrawComplexControl(unsafe.Pointer(this.h), control, option.cPointer(), painter.cPointer(), widget.cPointer())

}
func (this *QProxyStyle) OnDrawComplexControl(slot func(super func(control ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget), control ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_DrawComplexControl(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawComplexControl
func miqt_exec_callback_QProxyStyle_DrawComplexControl(self QProxyStyle, cb intptr_t, control ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(control ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget), control ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOptionComplex(option)

	slotval3 := newQPainter(painter)

	slotval4 := newQWidget(widget)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawComplexControl, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_DrawItemText(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	QProxyStyle_virtualbase_DrawItemText(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), (int)(flags), pal.cPointer(), (bool)(enabled), text_ms, (int)(textRole))

}
func (this *QProxyStyle) OnDrawItemText(slot func(super func(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole), painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_DrawItemText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawItemText
func miqt_exec_callback_QProxyStyle_DrawItemText(self QProxyStyle, cb intptr_t, painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text struct_miqt_string, textRole int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole), painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	slotval3 := (int)(flags)

	slotval4 := newQPalette(pal)

	slotval5 := (bool)(enabled)

	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval6 := text_ret
	slotval7 := (QPalette__ColorRole)(textRole)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawItemText, slotval1, slotval2, slotval3, slotval4, slotval5, slotval6, slotval7)

}

func (this *QProxyStyle) callVirtualBase_DrawItemPixmap(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {

	QProxyStyle_virtualbase_DrawItemPixmap(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), (int)(alignment), pixmap.cPointer())

}
func (this *QProxyStyle) OnDrawItemPixmap(slot func(super func(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap), painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_DrawItemPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_DrawItemPixmap
func miqt_exec_callback_QProxyStyle_DrawItemPixmap(self QProxyStyle, cb intptr_t, painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap), painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	slotval3 := (int)(alignment)

	slotval4 := newQPixmap(pixmap)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_DrawItemPixmap, slotval1, slotval2, slotval3, slotval4)

}

func (this *QProxyStyle) callVirtualBase_SizeFromContents(typeVal ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize {

	_goptr := newQSize(QProxyStyle_virtualbase_SizeFromContents(unsafe.Pointer(this.h), typeVal, option.cPointer(), size.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnSizeFromContents(slot func(super func(typeVal ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize, typeVal ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_SizeFromContents(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_SizeFromContents
func miqt_exec_callback_QProxyStyle_SizeFromContents(self QProxyStyle, cb intptr_t, typeVal ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize, typeVal ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(option)

	slotval3 := newQSize(size)

	slotval4 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_SizeFromContents, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_SubElementRect(element SubElement, option *QStyleOption, widget *QWidget) *QRect {

	_goptr := newQRect(QProxyStyle_virtualbase_SubElementRect(unsafe.Pointer(this.h), element, option.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnSubElementRect(slot func(super func(element SubElement, option *QStyleOption, widget *QWidget) *QRect, element SubElement, option *QStyleOption, widget *QWidget) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_SubElementRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_SubElementRect
func miqt_exec_callback_QProxyStyle_SubElementRect(self QProxyStyle, cb intptr_t, element SubElement, option *QStyleOption, widget *QWidget) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(element SubElement, option *QStyleOption, widget *QWidget) *QRect, element SubElement, option *QStyleOption, widget *QWidget) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(option)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_SubElementRect, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_SubControlRect(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, widget *QWidget) *QRect {

	_goptr := newQRect(QProxyStyle_virtualbase_SubControlRect(unsafe.Pointer(this.h), cc, opt.cPointer(), sc, widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnSubControlRect(slot func(super func(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, widget *QWidget) *QRect, cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, widget *QWidget) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_SubControlRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_SubControlRect
func miqt_exec_callback_QProxyStyle_SubControlRect(self QProxyStyle, cb intptr_t, cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, widget *QWidget) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, widget *QWidget) *QRect, cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, widget *QWidget) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOptionComplex(opt)

	xxxxxxxxx
	slotval4 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_SubControlRect, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_ItemTextRect(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	_goptr := newQRect(QProxyStyle_virtualbase_ItemTextRect(unsafe.Pointer(this.h), fm.cPointer(), r.cPointer(), (int)(flags), (bool)(enabled), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnItemTextRect(slot func(super func(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect, fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_ItemTextRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_ItemTextRect
func miqt_exec_callback_QProxyStyle_ItemTextRect(self QProxyStyle, cb intptr_t, fm *QFontMetrics, r *QRect, flags int, enabled bool, text struct_miqt_string) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect, fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFontMetrics(fm)

	slotval2 := newQRect(r)

	slotval3 := (int)(flags)

	slotval4 := (bool)(enabled)

	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval5 := text_ret

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_ItemTextRect, slotval1, slotval2, slotval3, slotval4, slotval5)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_ItemPixmapRect(r *QRect, flags int, pixmap *QPixmap) *QRect {

	_goptr := newQRect(QProxyStyle_virtualbase_ItemPixmapRect(unsafe.Pointer(this.h), r.cPointer(), (int)(flags), pixmap.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnItemPixmapRect(slot func(super func(r *QRect, flags int, pixmap *QPixmap) *QRect, r *QRect, flags int, pixmap *QPixmap) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_ItemPixmapRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_ItemPixmapRect
func miqt_exec_callback_QProxyStyle_ItemPixmapRect(self QProxyStyle, cb intptr_t, r *QRect, flags int, pixmap *QPixmap) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r *QRect, flags int, pixmap *QPixmap) *QRect, r *QRect, flags int, pixmap *QPixmap) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(r)

	slotval2 := (int)(flags)

	slotval3 := newQPixmap(pixmap)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_ItemPixmapRect, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_HitTestComplexControl(control ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) SubControl {

	xxxxxxxxx
}
func (this *QProxyStyle) OnHitTestComplexControl(slot func(super func(control ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) SubControl, control ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) SubControl) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_HitTestComplexControl(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_HitTestComplexControl
func miqt_exec_callback_QProxyStyle_HitTestComplexControl(self QProxyStyle, cb intptr_t, control ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) SubControl {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(control ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) SubControl, control ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) SubControl)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOptionComplex(option)

	slotval3 := newQPoint(pos)

	slotval4 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_HitTestComplexControl, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn

}

func (this *QProxyStyle) callVirtualBase_StyleHint(hint StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {

	return (int)(QProxyStyle_virtualbase_StyleHint(unsafe.Pointer(this.h), hint, option.cPointer(), widget.cPointer(), returnData.cPointer()))

}
func (this *QProxyStyle) OnStyleHint(slot func(super func(hint StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int, hint StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_StyleHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StyleHint
func miqt_exec_callback_QProxyStyle_StyleHint(self QProxyStyle, cb intptr_t, hint StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(hint StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int, hint StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(option)

	slotval3 := newQWidget(widget)

	slotval4 := newQStyleHintReturn(returnData)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StyleHint, slotval1, slotval2, slotval3, slotval4)

	return (int)(virtualReturn)

}

func (this *QProxyStyle) callVirtualBase_PixelMetric(metric PixelMetric, option *QStyleOption, widget *QWidget) int {

	return (int)(QProxyStyle_virtualbase_PixelMetric(unsafe.Pointer(this.h), metric, option.cPointer(), widget.cPointer()))

}
func (this *QProxyStyle) OnPixelMetric(slot func(super func(metric PixelMetric, option *QStyleOption, widget *QWidget) int, metric PixelMetric, option *QStyleOption, widget *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_PixelMetric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_PixelMetric
func miqt_exec_callback_QProxyStyle_PixelMetric(self QProxyStyle, cb intptr_t, metric PixelMetric, option *QStyleOption, widget *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(metric PixelMetric, option *QStyleOption, widget *QWidget) int, metric PixelMetric, option *QStyleOption, widget *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(option)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_PixelMetric, slotval1, slotval2, slotval3)

	return (int)(virtualReturn)

}

func (this *QProxyStyle) callVirtualBase_LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {

	return (int)(QProxyStyle_virtualbase_LayoutSpacing(unsafe.Pointer(this.h), (int)(control1), (int)(control2), (int)(orientation), option.cPointer(), widget.cPointer()))

}
func (this *QProxyStyle) OnLayoutSpacing(slot func(super func(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int, control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_LayoutSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_LayoutSpacing
func miqt_exec_callback_QProxyStyle_LayoutSpacing(self QProxyStyle, cb intptr_t, control1 int, control2 int, orientation int, option *QStyleOption, widget *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int, control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QSizePolicy__ControlType)(control1)

	slotval2 := (QSizePolicy__ControlType)(control2)

	slotval3 := (Orientation)(orientation)

	slotval4 := newQStyleOption(option)

	slotval5 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_LayoutSpacing, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (int)(virtualReturn)

}

func (this *QProxyStyle) callVirtualBase_StandardIcon(standardIcon StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon {

	_goptr := newQIcon(QProxyStyle_virtualbase_StandardIcon(unsafe.Pointer(this.h), standardIcon, option.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnStandardIcon(slot func(super func(standardIcon StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon, standardIcon StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_StandardIcon(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StandardIcon
func miqt_exec_callback_QProxyStyle_StandardIcon(self QProxyStyle, cb intptr_t, standardIcon StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(standardIcon StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon, standardIcon StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(option)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StandardIcon, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_StandardPixmap(standardPixmap StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {

	_goptr := newQPixmap(QProxyStyle_virtualbase_StandardPixmap(unsafe.Pointer(this.h), standardPixmap, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnStandardPixmap(slot func(super func(standardPixmap StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap, standardPixmap StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_StandardPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StandardPixmap
func miqt_exec_callback_QProxyStyle_StandardPixmap(self QProxyStyle, cb intptr_t, standardPixmap StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(standardPixmap StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap, standardPixmap StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StandardPixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_GeneratedIconPixmap(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {

	_goptr := newQPixmap(QProxyStyle_virtualbase_GeneratedIconPixmap(unsafe.Pointer(this.h), (int)(iconMode), pixmap.cPointer(), opt.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnGeneratedIconPixmap(slot func(super func(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap, iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_GeneratedIconPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_GeneratedIconPixmap
func miqt_exec_callback_QProxyStyle_GeneratedIconPixmap(self QProxyStyle, cb intptr_t, iconMode int, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap, iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QIcon__Mode)(iconMode)

	slotval2 := newQPixmap(pixmap)

	slotval3 := newQStyleOption(opt)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_GeneratedIconPixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_StandardPalette() *QPalette {

	_goptr := newQPalette(QProxyStyle_virtualbase_StandardPalette(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProxyStyle) OnStandardPalette(slot func(super func() *QPalette) *QPalette) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_StandardPalette(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_StandardPalette
func miqt_exec_callback_QProxyStyle_StandardPalette(self QProxyStyle, cb intptr_t) *QPalette {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPalette) *QPalette)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_StandardPalette)

	return virtualReturn.cPointer()

}

func (this *QProxyStyle) callVirtualBase_Polish(widget *QWidget) {

	QProxyStyle_virtualbase_Polish(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QProxyStyle) OnPolish(slot func(super func(widget *QWidget), widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_Polish(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_Polish
func miqt_exec_callback_QProxyStyle_Polish(self QProxyStyle, cb intptr_t, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(widget)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_Polish, slotval1)

}

func (this *QProxyStyle) callVirtualBase_PolishWithPal(pal *QPalette) {

	QProxyStyle_virtualbase_PolishWithPal(unsafe.Pointer(this.h), pal.cPointer())

}
func (this *QProxyStyle) OnPolishWithPal(slot func(super func(pal *QPalette), pal *QPalette)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_PolishWithPal(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_PolishWithPal
func miqt_exec_callback_QProxyStyle_PolishWithPal(self QProxyStyle, cb intptr_t, pal *QPalette) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pal *QPalette), pal *QPalette))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPalette(pal)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_PolishWithPal, slotval1)

}

func (this *QProxyStyle) callVirtualBase_PolishWithApp(app *QApplication) {

	QProxyStyle_virtualbase_PolishWithApp(unsafe.Pointer(this.h), app.cPointer())

}
func (this *QProxyStyle) OnPolishWithApp(slot func(super func(app *QApplication), app *QApplication)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_PolishWithApp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_PolishWithApp
func miqt_exec_callback_QProxyStyle_PolishWithApp(self QProxyStyle, cb intptr_t, app *QApplication) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(app *QApplication), app *QApplication))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQApplication(app)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_PolishWithApp, slotval1)

}

func (this *QProxyStyle) callVirtualBase_Unpolish(widget *QWidget) {

	QProxyStyle_virtualbase_Unpolish(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QProxyStyle) OnUnpolish(slot func(super func(widget *QWidget), widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_Unpolish(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_Unpolish
func miqt_exec_callback_QProxyStyle_Unpolish(self QProxyStyle, cb intptr_t, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(widget)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_Unpolish, slotval1)

}

func (this *QProxyStyle) callVirtualBase_UnpolishWithApp(app *QApplication) {

	QProxyStyle_virtualbase_UnpolishWithApp(unsafe.Pointer(this.h), app.cPointer())

}
func (this *QProxyStyle) OnUnpolishWithApp(slot func(super func(app *QApplication), app *QApplication)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_UnpolishWithApp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_UnpolishWithApp
func miqt_exec_callback_QProxyStyle_UnpolishWithApp(self QProxyStyle, cb intptr_t, app *QApplication) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(app *QApplication), app *QApplication))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQApplication(app)

	gofunc((&QProxyStyle{h: self}).callVirtualBase_UnpolishWithApp, slotval1)

}

func (this *QProxyStyle) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QProxyStyle_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QProxyStyle) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProxyStyle_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProxyStyle_Event
func miqt_exec_callback_QProxyStyle_Event(self QProxyStyle, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QProxyStyle{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}
