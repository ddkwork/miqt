package qt

import (
	"unsafe"
)

type QStylePainter struct {
	h          uintptr
	isSubclass bool
}

// NewQStylePainter constructs a new QStylePainter object.
func NewQStylePainter(w *QWidget) *QStylePainter {
	ret := newQStylePainter(QStylePainter_new(w.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStylePainter2 constructs a new QStylePainter object.
func NewQStylePainter2() *QStylePainter {
	ret := newQStylePainter(QStylePainter_new2())
	ret.isSubclass = true
	return ret
}

// NewQStylePainter3 constructs a new QStylePainter object.
func NewQStylePainter3(pd *QPaintDevice, w *QWidget) *QStylePainter {
	ret := newQStylePainter(QStylePainter_new3(pd.cPointer(), w.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStylePainter) Begin(w *QWidget) bool {
	return (bool)(QStylePainter_Begin(this.h, w.cPointer()))
}

func (this *QStylePainter) Begin2(pd *QPaintDevice, w *QWidget) bool {
	return (bool)(QStylePainter_Begin2(this.h, pd.cPointer(), w.cPointer()))
}

func (this *QStylePainter) DrawPrimitive(pe QStyle__PrimitiveElement, opt *QStyleOption) {
	QStylePainter_DrawPrimitive(this.h, (int)(pe), opt.cPointer())
}

func (this *QStylePainter) DrawControl(ce QStyle__ControlElement, opt *QStyleOption) {
	QStylePainter_DrawControl(this.h, (int)(ce), opt.cPointer())
}

func (this *QStylePainter) DrawComplexControl(cc QStyle__ComplexControl, opt *QStyleOptionComplex) {
	QStylePainter_DrawComplexControl(this.h, (int)(cc), opt.cPointer())
}

func (this *QStylePainter) DrawItemText(r *QRect, flags int, pal *QPalette, enabled bool, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStylePainter_DrawItemText(this.h, r.cPointer(), (int)(flags), pal.cPointer(), (bool)(enabled), text_ms)
}

func (this *QStylePainter) DrawItemPixmap(r *QRect, flags int, pixmap *QPixmap) {
	QStylePainter_DrawItemPixmap(this.h, r.cPointer(), (int)(flags), pixmap.cPointer())
}

func (this *QStylePainter) Style() *QStyle {
	return newQStyle(QStylePainter_Style(this.h))
}

func (this *QStylePainter) DrawItemText6(r *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStylePainter_DrawItemText6(this.h, r.cPointer(), (int)(flags), pal.cPointer(), (bool)(enabled), text_ms, (int)(textRole))
}
