package qt

import (
	"unsafe"
)

type QCommonStyle struct {
	h          uintptr
	isSubclass bool
}

// NewQCommonStyle constructs a new QCommonStyle object.
func NewQCommonStyle() *QCommonStyle {

	ret := newQCommonStyle(QCommonStyle_new())
	ret.isSubclass = true
	return ret
}

func (this *QCommonStyle) MetaObject() *QMetaObject {
	return newQMetaObject(QCommonStyle_MetaObject(this.h))
}

func (this *QCommonStyle) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCommonStyle_Metacast(this.h, param1_Cstring))
}

func QCommonStyle_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCommonStyle_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommonStyle) DrawPrimitive(pe PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	QCommonStyle_DrawPrimitive(this.h, pe, opt.cPointer(), p.cPointer(), w.cPointer())
}

func (this *QCommonStyle) DrawControl(element ControlElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	QCommonStyle_DrawControl(this.h, element, opt.cPointer(), p.cPointer(), w.cPointer())
}

func (this *QCommonStyle) SubElementRect(r SubElement, opt *QStyleOption, widget *QWidget) *QRect {
	_goptr := newQRect(QCommonStyle_SubElementRect(this.h, r, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommonStyle) DrawComplexControl(cc ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget) {
	QCommonStyle_DrawComplexControl(this.h, cc, opt.cPointer(), p.cPointer(), w.cPointer())
}

func (this *QCommonStyle) HitTestComplexControl(cc ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) SubControl {
	xxxxxxxxx
}

func (this *QCommonStyle) SubControlRect(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, w *QWidget) *QRect {
	_goptr := newQRect(QCommonStyle_SubControlRect(this.h, cc, opt.cPointer(), sc, w.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommonStyle) SizeFromContents(ct ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize {
	_goptr := newQSize(QCommonStyle_SizeFromContents(this.h, ct, opt.cPointer(), contentsSize.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommonStyle) PixelMetric(m PixelMetric, opt *QStyleOption, widget *QWidget) int {
	return (int)(QCommonStyle_PixelMetric(this.h, m, opt.cPointer(), widget.cPointer()))
}

func (this *QCommonStyle) StyleHint(sh StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int {
	return (int)(QCommonStyle_StyleHint(this.h, sh, opt.cPointer(), w.cPointer(), shret.cPointer()))
}

func (this *QCommonStyle) StandardIcon(standardIcon StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon {
	_goptr := newQIcon(QCommonStyle_StandardIcon(this.h, standardIcon, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommonStyle) StandardPixmap(sp StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	_goptr := newQPixmap(QCommonStyle_StandardPixmap(this.h, sp, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommonStyle) GeneratedIconPixmap(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	_goptr := newQPixmap(QCommonStyle_GeneratedIconPixmap(this.h, (int)(iconMode), pixmap.cPointer(), opt.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommonStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {
	return (int)(QCommonStyle_LayoutSpacing(this.h, (int)(control1), (int)(control2), (int)(orientation), option.cPointer(), widget.cPointer()))
}

func (this *QCommonStyle) Polish(param1 *QPalette) {
	QCommonStyle_Polish(this.h, param1.cPointer())
}

func (this *QCommonStyle) PolishWithApp(app *QApplication) {
	QCommonStyle_PolishWithApp(this.h, app.cPointer())
}

func (this *QCommonStyle) PolishWithWidget(widget *QWidget) {
	QCommonStyle_PolishWithWidget(this.h, widget.cPointer())
}

func (this *QCommonStyle) Unpolish(widget *QWidget) {
	QCommonStyle_Unpolish(this.h, widget.cPointer())
}

func (this *QCommonStyle) UnpolishWithApplication(application *QApplication) {
	QCommonStyle_UnpolishWithApplication(this.h, application.cPointer())
}

func QCommonStyle_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCommonStyle_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCommonStyle_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCommonStyle_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommonStyle) callVirtualBase_DrawPrimitive(pe PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget) {

	QCommonStyle_virtualbase_DrawPrimitive(unsafe.Pointer(this.h), pe, opt.cPointer(), p.cPointer(), w.cPointer())

}
func (this *QCommonStyle) OnDrawPrimitive(slot func(super func(pe PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget), pe PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_DrawPrimitive(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_DrawPrimitive
func miqt_exec_callback_QCommonStyle_DrawPrimitive(self QCommonStyle, cb intptr_t, pe PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pe PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget), pe PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQPainter(p)

	slotval4 := newQWidget(w)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_DrawPrimitive, slotval1, slotval2, slotval3, slotval4)

}

func (this *QCommonStyle) callVirtualBase_DrawControl(element ControlElement, opt *QStyleOption, p *QPainter, w *QWidget) {

	QCommonStyle_virtualbase_DrawControl(unsafe.Pointer(this.h), element, opt.cPointer(), p.cPointer(), w.cPointer())

}
func (this *QCommonStyle) OnDrawControl(slot func(super func(element ControlElement, opt *QStyleOption, p *QPainter, w *QWidget), element ControlElement, opt *QStyleOption, p *QPainter, w *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_DrawControl(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_DrawControl
func miqt_exec_callback_QCommonStyle_DrawControl(self QCommonStyle, cb intptr_t, element ControlElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(element ControlElement, opt *QStyleOption, p *QPainter, w *QWidget), element ControlElement, opt *QStyleOption, p *QPainter, w *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQPainter(p)

	slotval4 := newQWidget(w)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_DrawControl, slotval1, slotval2, slotval3, slotval4)

}

func (this *QCommonStyle) callVirtualBase_SubElementRect(r SubElement, opt *QStyleOption, widget *QWidget) *QRect {

	_goptr := newQRect(QCommonStyle_virtualbase_SubElementRect(unsafe.Pointer(this.h), r, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnSubElementRect(slot func(super func(r SubElement, opt *QStyleOption, widget *QWidget) *QRect, r SubElement, opt *QStyleOption, widget *QWidget) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_SubElementRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_SubElementRect
func miqt_exec_callback_QCommonStyle_SubElementRect(self QCommonStyle, cb intptr_t, r SubElement, opt *QStyleOption, widget *QWidget) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r SubElement, opt *QStyleOption, widget *QWidget) *QRect, r SubElement, opt *QStyleOption, widget *QWidget) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_SubElementRect, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_DrawComplexControl(cc ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget) {

	QCommonStyle_virtualbase_DrawComplexControl(unsafe.Pointer(this.h), cc, opt.cPointer(), p.cPointer(), w.cPointer())

}
func (this *QCommonStyle) OnDrawComplexControl(slot func(super func(cc ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget), cc ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_DrawComplexControl(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_DrawComplexControl
func miqt_exec_callback_QCommonStyle_DrawComplexControl(self QCommonStyle, cb intptr_t, cc ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cc ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget), cc ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOptionComplex(opt)

	slotval3 := newQPainter(p)

	slotval4 := newQWidget(w)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_DrawComplexControl, slotval1, slotval2, slotval3, slotval4)

}

func (this *QCommonStyle) callVirtualBase_HitTestComplexControl(cc ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) SubControl {

	xxxxxxxxx
}
func (this *QCommonStyle) OnHitTestComplexControl(slot func(super func(cc ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) SubControl, cc ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) SubControl) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_HitTestComplexControl(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_HitTestComplexControl
func miqt_exec_callback_QCommonStyle_HitTestComplexControl(self QCommonStyle, cb intptr_t, cc ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) SubControl {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cc ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) SubControl, cc ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) SubControl)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOptionComplex(opt)

	slotval3 := newQPoint(pt)

	slotval4 := newQWidget(w)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_HitTestComplexControl, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn

}

func (this *QCommonStyle) callVirtualBase_SubControlRect(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, w *QWidget) *QRect {

	_goptr := newQRect(QCommonStyle_virtualbase_SubControlRect(unsafe.Pointer(this.h), cc, opt.cPointer(), sc, w.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnSubControlRect(slot func(super func(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, w *QWidget) *QRect, cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, w *QWidget) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_SubControlRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_SubControlRect
func miqt_exec_callback_QCommonStyle_SubControlRect(self QCommonStyle, cb intptr_t, cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, w *QWidget) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, w *QWidget) *QRect, cc ComplexControl, opt *QStyleOptionComplex, sc SubControl, w *QWidget) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOptionComplex(opt)

	xxxxxxxxx
	slotval4 := newQWidget(w)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_SubControlRect, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_SizeFromContents(ct ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize {

	_goptr := newQSize(QCommonStyle_virtualbase_SizeFromContents(unsafe.Pointer(this.h), ct, opt.cPointer(), contentsSize.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnSizeFromContents(slot func(super func(ct ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize, ct ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_SizeFromContents(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_SizeFromContents
func miqt_exec_callback_QCommonStyle_SizeFromContents(self QCommonStyle, cb intptr_t, ct ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ct ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize, ct ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQSize(contentsSize)

	slotval4 := newQWidget(widget)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_SizeFromContents, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_PixelMetric(m PixelMetric, opt *QStyleOption, widget *QWidget) int {

	return (int)(QCommonStyle_virtualbase_PixelMetric(unsafe.Pointer(this.h), m, opt.cPointer(), widget.cPointer()))

}
func (this *QCommonStyle) OnPixelMetric(slot func(super func(m PixelMetric, opt *QStyleOption, widget *QWidget) int, m PixelMetric, opt *QStyleOption, widget *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_PixelMetric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_PixelMetric
func miqt_exec_callback_QCommonStyle_PixelMetric(self QCommonStyle, cb intptr_t, m PixelMetric, opt *QStyleOption, widget *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(m PixelMetric, opt *QStyleOption, widget *QWidget) int, m PixelMetric, opt *QStyleOption, widget *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_PixelMetric, slotval1, slotval2, slotval3)

	return (int)(virtualReturn)

}

func (this *QCommonStyle) callVirtualBase_StyleHint(sh StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int {

	return (int)(QCommonStyle_virtualbase_StyleHint(unsafe.Pointer(this.h), sh, opt.cPointer(), w.cPointer(), shret.cPointer()))

}
func (this *QCommonStyle) OnStyleHint(slot func(super func(sh StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int, sh StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_StyleHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_StyleHint
func miqt_exec_callback_QCommonStyle_StyleHint(self QCommonStyle, cb intptr_t, sh StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sh StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int, sh StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQWidget(w)

	slotval4 := newQStyleHintReturn(shret)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_StyleHint, slotval1, slotval2, slotval3, slotval4)

	return (int)(virtualReturn)

}

func (this *QCommonStyle) callVirtualBase_StandardIcon(standardIcon StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon {

	_goptr := newQIcon(QCommonStyle_virtualbase_StandardIcon(unsafe.Pointer(this.h), standardIcon, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnStandardIcon(slot func(super func(standardIcon StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon, standardIcon StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_StandardIcon(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_StandardIcon
func miqt_exec_callback_QCommonStyle_StandardIcon(self QCommonStyle, cb intptr_t, standardIcon StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(standardIcon StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon, standardIcon StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_StandardIcon, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_StandardPixmap(sp StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {

	_goptr := newQPixmap(QCommonStyle_virtualbase_StandardPixmap(unsafe.Pointer(this.h), sp, opt.cPointer(), widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnStandardPixmap(slot func(super func(sp StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap, sp StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_StandardPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_StandardPixmap
func miqt_exec_callback_QCommonStyle_StandardPixmap(self QCommonStyle, cb intptr_t, sp StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sp StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap, sp StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQStyleOption(opt)

	slotval3 := newQWidget(widget)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_StandardPixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_GeneratedIconPixmap(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {

	_goptr := newQPixmap(QCommonStyle_virtualbase_GeneratedIconPixmap(unsafe.Pointer(this.h), (int)(iconMode), pixmap.cPointer(), opt.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnGeneratedIconPixmap(slot func(super func(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap, iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_GeneratedIconPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_GeneratedIconPixmap
func miqt_exec_callback_QCommonStyle_GeneratedIconPixmap(self QCommonStyle, cb intptr_t, iconMode int, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap, iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QIcon__Mode)(iconMode)

	slotval2 := newQPixmap(pixmap)

	slotval3 := newQStyleOption(opt)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_GeneratedIconPixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {

	return (int)(QCommonStyle_virtualbase_LayoutSpacing(unsafe.Pointer(this.h), (int)(control1), (int)(control2), (int)(orientation), option.cPointer(), widget.cPointer()))

}
func (this *QCommonStyle) OnLayoutSpacing(slot func(super func(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int, control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_LayoutSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_LayoutSpacing
func miqt_exec_callback_QCommonStyle_LayoutSpacing(self QCommonStyle, cb intptr_t, control1 int, control2 int, orientation int, option *QStyleOption, widget *QWidget) int {
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

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_LayoutSpacing, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (int)(virtualReturn)

}

func (this *QCommonStyle) callVirtualBase_Polish(param1 *QPalette) {

	QCommonStyle_virtualbase_Polish(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCommonStyle) OnPolish(slot func(super func(param1 *QPalette), param1 *QPalette)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_Polish(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_Polish
func miqt_exec_callback_QCommonStyle_Polish(self QCommonStyle, cb intptr_t, param1 *QPalette) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPalette), param1 *QPalette))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPalette(param1)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_Polish, slotval1)

}

func (this *QCommonStyle) callVirtualBase_PolishWithApp(app *QApplication) {

	QCommonStyle_virtualbase_PolishWithApp(unsafe.Pointer(this.h), app.cPointer())

}
func (this *QCommonStyle) OnPolishWithApp(slot func(super func(app *QApplication), app *QApplication)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_PolishWithApp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_PolishWithApp
func miqt_exec_callback_QCommonStyle_PolishWithApp(self QCommonStyle, cb intptr_t, app *QApplication) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(app *QApplication), app *QApplication))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQApplication(app)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_PolishWithApp, slotval1)

}

func (this *QCommonStyle) callVirtualBase_PolishWithWidget(widget *QWidget) {

	QCommonStyle_virtualbase_PolishWithWidget(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QCommonStyle) OnPolishWithWidget(slot func(super func(widget *QWidget), widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_PolishWithWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_PolishWithWidget
func miqt_exec_callback_QCommonStyle_PolishWithWidget(self QCommonStyle, cb intptr_t, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(widget)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_PolishWithWidget, slotval1)

}

func (this *QCommonStyle) callVirtualBase_Unpolish(widget *QWidget) {

	QCommonStyle_virtualbase_Unpolish(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QCommonStyle) OnUnpolish(slot func(super func(widget *QWidget), widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_Unpolish(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_Unpolish
func miqt_exec_callback_QCommonStyle_Unpolish(self QCommonStyle, cb intptr_t, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(widget)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_Unpolish, slotval1)

}

func (this *QCommonStyle) callVirtualBase_UnpolishWithApplication(application *QApplication) {

	QCommonStyle_virtualbase_UnpolishWithApplication(unsafe.Pointer(this.h), application.cPointer())

}
func (this *QCommonStyle) OnUnpolishWithApplication(slot func(super func(application *QApplication), application *QApplication)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_UnpolishWithApplication(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_UnpolishWithApplication
func miqt_exec_callback_QCommonStyle_UnpolishWithApplication(self QCommonStyle, cb intptr_t, application *QApplication) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(application *QApplication), application *QApplication))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQApplication(application)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_UnpolishWithApplication, slotval1)

}

func (this *QCommonStyle) callVirtualBase_ItemTextRect(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	_goptr := newQRect(QCommonStyle_virtualbase_ItemTextRect(unsafe.Pointer(this.h), fm.cPointer(), r.cPointer(), (int)(flags), (bool)(enabled), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnItemTextRect(slot func(super func(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect, fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_ItemTextRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_ItemTextRect
func miqt_exec_callback_QCommonStyle_ItemTextRect(self QCommonStyle, cb intptr_t, fm *QFontMetrics, r *QRect, flags int, enabled bool, text struct_miqt_string) *QRect {
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

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_ItemTextRect, slotval1, slotval2, slotval3, slotval4, slotval5)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_ItemPixmapRect(r *QRect, flags int, pixmap *QPixmap) *QRect {

	_goptr := newQRect(QCommonStyle_virtualbase_ItemPixmapRect(unsafe.Pointer(this.h), r.cPointer(), (int)(flags), pixmap.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnItemPixmapRect(slot func(super func(r *QRect, flags int, pixmap *QPixmap) *QRect, r *QRect, flags int, pixmap *QPixmap) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_ItemPixmapRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_ItemPixmapRect
func miqt_exec_callback_QCommonStyle_ItemPixmapRect(self QCommonStyle, cb intptr_t, r *QRect, flags int, pixmap *QPixmap) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r *QRect, flags int, pixmap *QPixmap) *QRect, r *QRect, flags int, pixmap *QPixmap) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(r)

	slotval2 := (int)(flags)

	slotval3 := newQPixmap(pixmap)

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_ItemPixmapRect, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QCommonStyle) callVirtualBase_DrawItemText(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	QCommonStyle_virtualbase_DrawItemText(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), (int)(flags), pal.cPointer(), (bool)(enabled), text_ms, (int)(textRole))

}
func (this *QCommonStyle) OnDrawItemText(slot func(super func(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole), painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_DrawItemText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_DrawItemText
func miqt_exec_callback_QCommonStyle_DrawItemText(self QCommonStyle, cb intptr_t, painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text struct_miqt_string, textRole int) {
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

	gofunc((&QCommonStyle{h: self}).callVirtualBase_DrawItemText, slotval1, slotval2, slotval3, slotval4, slotval5, slotval6, slotval7)

}

func (this *QCommonStyle) callVirtualBase_DrawItemPixmap(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {

	QCommonStyle_virtualbase_DrawItemPixmap(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), (int)(alignment), pixmap.cPointer())

}
func (this *QCommonStyle) OnDrawItemPixmap(slot func(super func(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap), painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_DrawItemPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_DrawItemPixmap
func miqt_exec_callback_QCommonStyle_DrawItemPixmap(self QCommonStyle, cb intptr_t, painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap), painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	slotval3 := (int)(alignment)

	slotval4 := newQPixmap(pixmap)

	gofunc((&QCommonStyle{h: self}).callVirtualBase_DrawItemPixmap, slotval1, slotval2, slotval3, slotval4)

}

func (this *QCommonStyle) callVirtualBase_StandardPalette() *QPalette {

	_goptr := newQPalette(QCommonStyle_virtualbase_StandardPalette(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommonStyle) OnStandardPalette(slot func(super func() *QPalette) *QPalette) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_StandardPalette(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_StandardPalette
func miqt_exec_callback_QCommonStyle_StandardPalette(self QCommonStyle, cb intptr_t) *QPalette {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPalette) *QPalette)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_StandardPalette)

	return virtualReturn.cPointer()

}
