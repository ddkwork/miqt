package qt

import (
	"unsafe"
)

type QFrame__Shape int

const (
	QFrame__NoFrame     QFrame__Shape = 0
	QFrame__Box         QFrame__Shape = 1
	QFrame__Panel       QFrame__Shape = 2
	QFrame__WinPanel    QFrame__Shape = 3
	QFrame__HLine       QFrame__Shape = 4
	QFrame__VLine       QFrame__Shape = 5
	QFrame__StyledPanel QFrame__Shape = 6
)

type QFrame__Shadow int

const (
	QFrame__Plain  QFrame__Shadow = 16
	QFrame__Raised QFrame__Shadow = 32
	QFrame__Sunken QFrame__Shadow = 48
)

type QFrame__StyleMask int

const (
	QFrame__Shadow_Mask QFrame__StyleMask = 240
	QFrame__Shape_Mask  QFrame__StyleMask = 15
)

type QFrame struct {
	h          uintptr
	isSubclass bool
}

// NewQFrame constructs a new QFrame object.
func NewQFrame(parent *QWidget) *QFrame {

	ret := newQFrame(QFrame_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFrame2 constructs a new QFrame object.
func NewQFrame2() *QFrame {

	ret := newQFrame(QFrame_new2())
	ret.isSubclass = true
	return ret
}

// NewQFrame3 constructs a new QFrame object.
func NewQFrame3(parent *QWidget, f WindowType) *QFrame {

	ret := newQFrame(QFrame_new3(parent.cPointer(), (int)(f)))
	ret.isSubclass = true
	return ret
}

func (this *QFrame) MetaObject() *QMetaObject {
	return newQMetaObject(QFrame_MetaObject(this.h))
}

func (this *QFrame) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFrame_Metacast(this.h, param1_Cstring))
}

func QFrame_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFrame_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFrame) FrameStyle() int {
	return (int)(QFrame_FrameStyle(this.h))
}

func (this *QFrame) SetFrameStyle(frameStyle int) {
	QFrame_SetFrameStyle(this.h, (int)(frameStyle))
}

func (this *QFrame) FrameWidth() int {
	return (int)(QFrame_FrameWidth(this.h))
}

func (this *QFrame) SizeHint() *QSize {
	_goptr := newQSize(QFrame_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFrame) FrameShape() Shape {
	xxxxxxxxx
}

func (this *QFrame) SetFrameShape(frameShape Shape) {
	QFrame_SetFrameShape(this.h, frameShape)
}

func (this *QFrame) FrameShadow() Shadow {
	xxxxxxxxx
}

func (this *QFrame) SetFrameShadow(frameShadow Shadow) {
	QFrame_SetFrameShadow(this.h, frameShadow)
}

func (this *QFrame) LineWidth() int {
	return (int)(QFrame_LineWidth(this.h))
}

func (this *QFrame) SetLineWidth(lineWidth int) {
	QFrame_SetLineWidth(this.h, (int)(lineWidth))
}

func (this *QFrame) MidLineWidth() int {
	return (int)(QFrame_MidLineWidth(this.h))
}

func (this *QFrame) SetMidLineWidth(midLineWidth int) {
	QFrame_SetMidLineWidth(this.h, (int)(midLineWidth))
}

func (this *QFrame) FrameRect() *QRect {
	_goptr := newQRect(QFrame_FrameRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFrame) SetFrameRect(frameRect *QRect) {
	QFrame_SetFrameRect(this.h, frameRect.cPointer())
}

func QFrame_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFrame_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFrame_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFrame_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFrame) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QFrame_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFrame) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_SizeHint
func miqt_exec_callback_QFrame_SizeHint(self QFrame, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QFrame) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QFrame_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QFrame) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_Event
func miqt_exec_callback_QFrame_Event(self QFrame, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFrame) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QFrame_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFrame) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_PaintEvent
func miqt_exec_callback_QFrame_PaintEvent(self QFrame, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QFrame{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QFrame) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QFrame_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFrame) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_ChangeEvent
func miqt_exec_callback_QFrame_ChangeEvent(self QFrame, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QFrame{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QFrame) callVirtualBase_InitStyleOption(option *QStyleOptionFrame) {

	QFrame_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QFrame) OnInitStyleOption(slot func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_InitStyleOption
func miqt_exec_callback_QFrame_InitStyleOption(self QFrame, cb intptr_t, option *QStyleOptionFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionFrame(option)

	gofunc((&QFrame{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QFrame) callVirtualBase_DevType() int {

	return (int)(QFrame_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QFrame) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_DevType
func miqt_exec_callback_QFrame_DevType(self QFrame, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QFrame) callVirtualBase_SetVisible(visible bool) {

	QFrame_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QFrame) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_SetVisible
func miqt_exec_callback_QFrame_SetVisible(self QFrame, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QFrame{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QFrame) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QFrame_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFrame) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_MinimumSizeHint
func miqt_exec_callback_QFrame_MinimumSizeHint(self QFrame, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QFrame) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QFrame_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QFrame) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_HeightForWidth
func miqt_exec_callback_QFrame_HeightForWidth(self QFrame, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QFrame) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QFrame_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QFrame) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_HasHeightForWidth
func miqt_exec_callback_QFrame_HasHeightForWidth(self QFrame, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QFrame) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QFrame_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QFrame) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_PaintEngine
func miqt_exec_callback_QFrame_PaintEngine(self QFrame, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QFrame) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QFrame_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_MousePressEvent
func miqt_exec_callback_QFrame_MousePressEvent(self QFrame, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QFrame) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QFrame_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_MouseReleaseEvent
func miqt_exec_callback_QFrame_MouseReleaseEvent(self QFrame, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QFrame) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QFrame_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_MouseDoubleClickEvent
func miqt_exec_callback_QFrame_MouseDoubleClickEvent(self QFrame, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QFrame) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QFrame_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_MouseMoveEvent
func miqt_exec_callback_QFrame_MouseMoveEvent(self QFrame, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QFrame) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QFrame_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_WheelEvent
func miqt_exec_callback_QFrame_WheelEvent(self QFrame, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QFrame) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QFrame_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_KeyPressEvent
func miqt_exec_callback_QFrame_KeyPressEvent(self QFrame, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QFrame) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QFrame_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_KeyReleaseEvent
func miqt_exec_callback_QFrame_KeyReleaseEvent(self QFrame, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QFrame) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QFrame_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_FocusInEvent
func miqt_exec_callback_QFrame_FocusInEvent(self QFrame, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QFrame) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QFrame_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_FocusOutEvent
func miqt_exec_callback_QFrame_FocusOutEvent(self QFrame, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QFrame) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QFrame_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_EnterEvent
func miqt_exec_callback_QFrame_EnterEvent(self QFrame, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QFrame) callVirtualBase_LeaveEvent(event *QEvent) {

	QFrame_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_LeaveEvent
func miqt_exec_callback_QFrame_LeaveEvent(self QFrame, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QFrame) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QFrame_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_MoveEvent
func miqt_exec_callback_QFrame_MoveEvent(self QFrame, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QFrame) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QFrame_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_ResizeEvent
func miqt_exec_callback_QFrame_ResizeEvent(self QFrame, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QFrame) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QFrame_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_CloseEvent
func miqt_exec_callback_QFrame_CloseEvent(self QFrame, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QFrame) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QFrame_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_ContextMenuEvent
func miqt_exec_callback_QFrame_ContextMenuEvent(self QFrame, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QFrame) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QFrame_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_TabletEvent
func miqt_exec_callback_QFrame_TabletEvent(self QFrame, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QFrame) callVirtualBase_ActionEvent(event *QActionEvent) {

	QFrame_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_ActionEvent
func miqt_exec_callback_QFrame_ActionEvent(self QFrame, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QFrame) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QFrame_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_DragEnterEvent
func miqt_exec_callback_QFrame_DragEnterEvent(self QFrame, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QFrame) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QFrame_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_DragMoveEvent
func miqt_exec_callback_QFrame_DragMoveEvent(self QFrame, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QFrame) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QFrame_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_DragLeaveEvent
func miqt_exec_callback_QFrame_DragLeaveEvent(self QFrame, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QFrame) callVirtualBase_DropEvent(event *QDropEvent) {

	QFrame_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_DropEvent
func miqt_exec_callback_QFrame_DropEvent(self QFrame, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QFrame) callVirtualBase_ShowEvent(event *QShowEvent) {

	QFrame_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_ShowEvent
func miqt_exec_callback_QFrame_ShowEvent(self QFrame, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QFrame) callVirtualBase_HideEvent(event *QHideEvent) {

	QFrame_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFrame) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_HideEvent
func miqt_exec_callback_QFrame_HideEvent(self QFrame, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QFrame{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QFrame) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QFrame_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QFrame) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_NativeEvent
func miqt_exec_callback_QFrame_NativeEvent(self QFrame, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QFrame) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QFrame_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QFrame) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_Metric
func miqt_exec_callback_QFrame_Metric(self QFrame, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QFrame) callVirtualBase_InitPainter(painter *QPainter) {

	QFrame_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QFrame) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_InitPainter
func miqt_exec_callback_QFrame_InitPainter(self QFrame, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QFrame{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QFrame) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QFrame_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QFrame) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_Redirected
func miqt_exec_callback_QFrame_Redirected(self QFrame, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFrame) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QFrame_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QFrame) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_SharedPainter
func miqt_exec_callback_QFrame_SharedPainter(self QFrame, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QFrame) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QFrame_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFrame) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_InputMethodEvent
func miqt_exec_callback_QFrame_InputMethodEvent(self QFrame, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QFrame{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QFrame) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QFrame_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFrame) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_InputMethodQuery
func miqt_exec_callback_QFrame_InputMethodQuery(self QFrame, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFrame) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QFrame_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QFrame) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_FocusNextPrevChild
func miqt_exec_callback_QFrame_FocusNextPrevChild(self QFrame, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
