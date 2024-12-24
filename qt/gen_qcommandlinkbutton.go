package qt

import (
	"unsafe"
)

type QCommandLinkButton struct {
	h          uintptr
	isSubclass bool
}

// NewQCommandLinkButton constructs a new QCommandLinkButton object.
func NewQCommandLinkButton(parent *QWidget) *QCommandLinkButton {

	ret := newQCommandLinkButton(QCommandLinkButton_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCommandLinkButton2 constructs a new QCommandLinkButton object.
func NewQCommandLinkButton2() *QCommandLinkButton {

	ret := newQCommandLinkButton(QCommandLinkButton_new2())
	ret.isSubclass = true
	return ret
}

// NewQCommandLinkButton3 constructs a new QCommandLinkButton object.
func NewQCommandLinkButton3(text string) *QCommandLinkButton {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQCommandLinkButton(QCommandLinkButton_new3(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLinkButton4 constructs a new QCommandLinkButton object.
func NewQCommandLinkButton4(text string, description string) *QCommandLinkButton {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))

	ret := newQCommandLinkButton(QCommandLinkButton_new4(text_ms, description_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLinkButton5 constructs a new QCommandLinkButton object.
func NewQCommandLinkButton5(text string, parent *QWidget) *QCommandLinkButton {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQCommandLinkButton(QCommandLinkButton_new5(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCommandLinkButton6 constructs a new QCommandLinkButton object.
func NewQCommandLinkButton6(text string, description string, parent *QWidget) *QCommandLinkButton {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))

	ret := newQCommandLinkButton(QCommandLinkButton_new6(text_ms, description_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCommandLinkButton) MetaObject() *QMetaObject {
	return newQMetaObject(QCommandLinkButton_MetaObject(this.h))
}

func (this *QCommandLinkButton) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCommandLinkButton_Metacast(this.h, param1_Cstring))
}

func QCommandLinkButton_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCommandLinkButton_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLinkButton) Description() string {
	var _ms struct_miqt_string = QCommandLinkButton_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLinkButton) SetDescription(description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QCommandLinkButton_SetDescription(this.h, description_ms)
}

func (this *QCommandLinkButton) SizeHint() *QSize {
	_goptr := newQSize(QCommandLinkButton_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommandLinkButton) HeightForWidth(param1 int) int {
	return (int)(QCommandLinkButton_HeightForWidth(this.h, (int)(param1)))
}

func (this *QCommandLinkButton) MinimumSizeHint() *QSize {
	_goptr := newQSize(QCommandLinkButton_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCommandLinkButton) InitStyleOption(option *QStyleOptionButton) {
	QCommandLinkButton_InitStyleOption(this.h, option.cPointer())
}

func QCommandLinkButton_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCommandLinkButton_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCommandLinkButton_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCommandLinkButton_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLinkButton) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QCommandLinkButton_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommandLinkButton) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_SizeHint
func miqt_exec_callback_QCommandLinkButton_SizeHint(self QCommandLinkButton, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCommandLinkButton{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QCommandLinkButton) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QCommandLinkButton_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QCommandLinkButton) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_HeightForWidth
func miqt_exec_callback_QCommandLinkButton_HeightForWidth(self QCommandLinkButton, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QCommandLinkButton{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QCommandLinkButton) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QCommandLinkButton_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCommandLinkButton) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_MinimumSizeHint
func miqt_exec_callback_QCommandLinkButton_MinimumSizeHint(self QCommandLinkButton, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCommandLinkButton{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QCommandLinkButton) callVirtualBase_InitStyleOption(option *QStyleOptionButton) {

	QCommandLinkButton_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QCommandLinkButton) OnInitStyleOption(slot func(super func(option *QStyleOptionButton), option *QStyleOptionButton)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_InitStyleOption
func miqt_exec_callback_QCommandLinkButton_InitStyleOption(self QCommandLinkButton, cb intptr_t, option *QStyleOptionButton) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionButton), option *QStyleOptionButton))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionButton(option)

	gofunc((&QCommandLinkButton{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QCommandLinkButton) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QCommandLinkButton_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QCommandLinkButton) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_Event
func miqt_exec_callback_QCommandLinkButton_Event(self QCommandLinkButton, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QCommandLinkButton{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QCommandLinkButton) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QCommandLinkButton_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCommandLinkButton) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_PaintEvent
func miqt_exec_callback_QCommandLinkButton_PaintEvent(self QCommandLinkButton, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QCommandLinkButton{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QCommandLinkButton) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QCommandLinkButton_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCommandLinkButton) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_KeyPressEvent
func miqt_exec_callback_QCommandLinkButton_KeyPressEvent(self QCommandLinkButton, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QCommandLinkButton{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QCommandLinkButton) callVirtualBase_FocusInEvent(param1 *QFocusEvent) {

	QCommandLinkButton_virtualbase_FocusInEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCommandLinkButton) OnFocusInEvent(slot func(super func(param1 *QFocusEvent), param1 *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_FocusInEvent
func miqt_exec_callback_QCommandLinkButton_FocusInEvent(self QCommandLinkButton, cb intptr_t, param1 *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFocusEvent), param1 *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(param1)

	gofunc((&QCommandLinkButton{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QCommandLinkButton) callVirtualBase_FocusOutEvent(param1 *QFocusEvent) {

	QCommandLinkButton_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCommandLinkButton) OnFocusOutEvent(slot func(super func(param1 *QFocusEvent), param1 *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_FocusOutEvent
func miqt_exec_callback_QCommandLinkButton_FocusOutEvent(self QCommandLinkButton, cb intptr_t, param1 *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFocusEvent), param1 *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(param1)

	gofunc((&QCommandLinkButton{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QCommandLinkButton) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QCommandLinkButton_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCommandLinkButton) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_MouseMoveEvent
func miqt_exec_callback_QCommandLinkButton_MouseMoveEvent(self QCommandLinkButton, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QCommandLinkButton{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QCommandLinkButton) callVirtualBase_HitButton(pos *QPoint) bool {

	return (bool)(QCommandLinkButton_virtualbase_HitButton(unsafe.Pointer(this.h), pos.cPointer()))

}
func (this *QCommandLinkButton) OnHitButton(slot func(super func(pos *QPoint) bool, pos *QPoint) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCommandLinkButton_override_virtual_HitButton(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCommandLinkButton_HitButton
func miqt_exec_callback_QCommandLinkButton_HitButton(self QCommandLinkButton, cb intptr_t, pos *QPoint) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos *QPoint) bool, pos *QPoint) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(pos)

	virtualReturn := gofunc((&QCommandLinkButton{h: self}).callVirtualBase_HitButton, slotval1)

	return (bool)(virtualReturn)

}
