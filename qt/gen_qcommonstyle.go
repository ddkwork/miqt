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
	g := newQCommonStyle(QCommonStyle_new())
	g.isSubclass = true
	return g
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

func (this *QCommonStyle) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QCommonStyle_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QCommonStyle) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_MetaObject
func miqt_exec_callback_QCommonStyle_MetaObject(self QCommonStyle, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QCommonStyle) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QCommonStyle_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QCommonStyle) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommonStyle_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommonStyle_Metacast
func miqt_exec_callback_QCommonStyle_Metacast(self QCommonStyle, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QCommonStyle{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
