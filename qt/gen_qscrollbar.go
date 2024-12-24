package qt

import (
	"unsafe"
)

type QScrollBar struct {
	h          uintptr
	isSubclass bool
}

// NewQScrollBar constructs a new QScrollBar object.
func NewQScrollBar(parent *QWidget) *QScrollBar {

	ret := newQScrollBar(QScrollBar_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQScrollBar2 constructs a new QScrollBar object.
func NewQScrollBar2() *QScrollBar {

	ret := newQScrollBar(QScrollBar_new2())
	ret.isSubclass = true
	return ret
}

// NewQScrollBar3 constructs a new QScrollBar object.
func NewQScrollBar3(param1 Orientation) *QScrollBar {

	ret := newQScrollBar(QScrollBar_new3((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQScrollBar4 constructs a new QScrollBar object.
func NewQScrollBar4(param1 Orientation, parent *QWidget) *QScrollBar {

	ret := newQScrollBar(QScrollBar_new4((int)(param1), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QScrollBar) MetaObject() *QMetaObject {
	return newQMetaObject(QScrollBar_MetaObject(this.h))
}

func (this *QScrollBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QScrollBar_Metacast(this.h, param1_Cstring))
}

func QScrollBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QScrollBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScrollBar) SizeHint() *QSize {
	_goptr := newQSize(QScrollBar_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollBar) Event(event *QEvent) bool {
	return (bool)(QScrollBar_Event(this.h, event.cPointer()))
}

func QScrollBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScrollBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScrollBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScrollBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScrollBar) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QScrollBar_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QScrollBar) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_SizeHint
func miqt_exec_callback_QScrollBar_SizeHint(self QScrollBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollBar{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QScrollBar) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QScrollBar_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QScrollBar) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_Event
func miqt_exec_callback_QScrollBar_Event(self QScrollBar, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QScrollBar{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QScrollBar) callVirtualBase_WheelEvent(param1 *QWheelEvent) {

	QScrollBar_virtualbase_WheelEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnWheelEvent(slot func(super func(param1 *QWheelEvent), param1 *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_WheelEvent
func miqt_exec_callback_QScrollBar_WheelEvent(self QScrollBar, cb intptr_t, param1 *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWheelEvent), param1 *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QScrollBar_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_PaintEvent
func miqt_exec_callback_QScrollBar_PaintEvent(self QScrollBar, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QScrollBar_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_MousePressEvent
func miqt_exec_callback_QScrollBar_MousePressEvent(self QScrollBar, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QScrollBar_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_MouseReleaseEvent
func miqt_exec_callback_QScrollBar_MouseReleaseEvent(self QScrollBar, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QScrollBar_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_MouseMoveEvent
func miqt_exec_callback_QScrollBar_MouseMoveEvent(self QScrollBar, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_HideEvent(param1 *QHideEvent) {

	QScrollBar_virtualbase_HideEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnHideEvent(slot func(super func(param1 *QHideEvent), param1 *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_HideEvent
func miqt_exec_callback_QScrollBar_HideEvent(self QScrollBar, cb intptr_t, param1 *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QHideEvent), param1 *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_SliderChange(change SliderChange) {

	QScrollBar_virtualbase_SliderChange(unsafe.Pointer(this.h), change)

}
func (this *QScrollBar) OnSliderChange(slot func(super func(change SliderChange), change SliderChange)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_SliderChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_SliderChange
func miqt_exec_callback_QScrollBar_SliderChange(self QScrollBar, cb intptr_t, change SliderChange) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(change SliderChange), change SliderChange))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QScrollBar{h: self}).callVirtualBase_SliderChange, slotval1)

}

func (this *QScrollBar) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QScrollBar_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_ContextMenuEvent
func miqt_exec_callback_QScrollBar_ContextMenuEvent(self QScrollBar, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_InitStyleOption(option *QStyleOptionSlider) {

	QScrollBar_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QScrollBar) OnInitStyleOption(slot func(super func(option *QStyleOptionSlider), option *QStyleOptionSlider)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_InitStyleOption
func miqt_exec_callback_QScrollBar_InitStyleOption(self QScrollBar, cb intptr_t, option *QStyleOptionSlider) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSlider), option *QStyleOptionSlider))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSlider(option)

	gofunc((&QScrollBar{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QScrollBar) callVirtualBase_KeyPressEvent(ev *QKeyEvent) {

	QScrollBar_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QScrollBar) OnKeyPressEvent(slot func(super func(ev *QKeyEvent), ev *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_KeyPressEvent
func miqt_exec_callback_QScrollBar_KeyPressEvent(self QScrollBar, cb intptr_t, ev *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QKeyEvent), ev *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(ev)

	gofunc((&QScrollBar{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QScrollBar_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollBar) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_TimerEvent
func miqt_exec_callback_QScrollBar_TimerEvent(self QScrollBar, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QScrollBar{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QScrollBar) callVirtualBase_ChangeEvent(e *QEvent) {

	QScrollBar_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QScrollBar) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollBar_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollBar_ChangeEvent
func miqt_exec_callback_QScrollBar_ChangeEvent(self QScrollBar, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QScrollBar{h: self}).callVirtualBase_ChangeEvent, slotval1)

}
