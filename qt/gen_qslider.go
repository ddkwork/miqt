package qt

import (
	"unsafe"
)

type QSlider__TickPosition int

const (
	QSlider__NoTicks        QSlider__TickPosition = 0
	QSlider__TicksAbove     QSlider__TickPosition = 1
	QSlider__TicksLeft      QSlider__TickPosition = 1
	QSlider__TicksBelow     QSlider__TickPosition = 2
	QSlider__TicksRight     QSlider__TickPosition = 2
	QSlider__TicksBothSides QSlider__TickPosition = 3
)

type QSlider struct {
	h          uintptr
	isSubclass bool
}

// NewQSlider constructs a new QSlider object.
func NewQSlider(parent *QWidget) *QSlider {

	ret := newQSlider(QSlider_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSlider2 constructs a new QSlider object.
func NewQSlider2() *QSlider {

	ret := newQSlider(QSlider_new2())
	ret.isSubclass = true
	return ret
}

// NewQSlider3 constructs a new QSlider object.
func NewQSlider3(orientation Orientation) *QSlider {

	ret := newQSlider(QSlider_new3((int)(orientation)))
	ret.isSubclass = true
	return ret
}

// NewQSlider4 constructs a new QSlider object.
func NewQSlider4(orientation Orientation, parent *QWidget) *QSlider {

	ret := newQSlider(QSlider_new4((int)(orientation), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSlider) MetaObject() *QMetaObject {
	return newQMetaObject(QSlider_MetaObject(this.h))
}

func (this *QSlider) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSlider_Metacast(this.h, param1_Cstring))
}

func QSlider_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSlider_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSlider) SizeHint() *QSize {
	_goptr := newQSize(QSlider_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSlider) MinimumSizeHint() *QSize {
	_goptr := newQSize(QSlider_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSlider) SetTickPosition(position TickPosition) {
	QSlider_SetTickPosition(this.h, position)
}

func (this *QSlider) TickPosition() TickPosition {
	xxxxxxxxx
}

func (this *QSlider) SetTickInterval(ti int) {
	QSlider_SetTickInterval(this.h, (int)(ti))
}

func (this *QSlider) TickInterval() int {
	return (int)(QSlider_TickInterval(this.h))
}

func (this *QSlider) Event(event *QEvent) bool {
	return (bool)(QSlider_Event(this.h, event.cPointer()))
}

func QSlider_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSlider_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSlider_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSlider_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSlider) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QSlider_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSlider) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_SizeHint
func miqt_exec_callback_QSlider_SizeHint(self QSlider, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSlider{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QSlider) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QSlider_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSlider) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_MinimumSizeHint
func miqt_exec_callback_QSlider_MinimumSizeHint(self QSlider, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSlider{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QSlider) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QSlider_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QSlider) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_Event
func miqt_exec_callback_QSlider_Event(self QSlider, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QSlider{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSlider) callVirtualBase_PaintEvent(ev *QPaintEvent) {

	QSlider_virtualbase_PaintEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QSlider) OnPaintEvent(slot func(super func(ev *QPaintEvent), ev *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_PaintEvent
func miqt_exec_callback_QSlider_PaintEvent(self QSlider, cb intptr_t, ev *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QPaintEvent), ev *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(ev)

	gofunc((&QSlider{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QSlider) callVirtualBase_MousePressEvent(ev *QMouseEvent) {

	QSlider_virtualbase_MousePressEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QSlider) OnMousePressEvent(slot func(super func(ev *QMouseEvent), ev *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_MousePressEvent
func miqt_exec_callback_QSlider_MousePressEvent(self QSlider, cb intptr_t, ev *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QMouseEvent), ev *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(ev)

	gofunc((&QSlider{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QSlider) callVirtualBase_MouseReleaseEvent(ev *QMouseEvent) {

	QSlider_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QSlider) OnMouseReleaseEvent(slot func(super func(ev *QMouseEvent), ev *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_MouseReleaseEvent
func miqt_exec_callback_QSlider_MouseReleaseEvent(self QSlider, cb intptr_t, ev *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QMouseEvent), ev *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(ev)

	gofunc((&QSlider{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QSlider) callVirtualBase_MouseMoveEvent(ev *QMouseEvent) {

	QSlider_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QSlider) OnMouseMoveEvent(slot func(super func(ev *QMouseEvent), ev *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_MouseMoveEvent
func miqt_exec_callback_QSlider_MouseMoveEvent(self QSlider, cb intptr_t, ev *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QMouseEvent), ev *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(ev)

	gofunc((&QSlider{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QSlider) callVirtualBase_InitStyleOption(option *QStyleOptionSlider) {

	QSlider_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QSlider) OnInitStyleOption(slot func(super func(option *QStyleOptionSlider), option *QStyleOptionSlider)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_InitStyleOption
func miqt_exec_callback_QSlider_InitStyleOption(self QSlider, cb intptr_t, option *QStyleOptionSlider) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSlider), option *QStyleOptionSlider))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSlider(option)

	gofunc((&QSlider{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QSlider) callVirtualBase_SliderChange(change SliderChange) {

	QSlider_virtualbase_SliderChange(unsafe.Pointer(this.h), change)

}
func (this *QSlider) OnSliderChange(slot func(super func(change SliderChange), change SliderChange)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_SliderChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_SliderChange
func miqt_exec_callback_QSlider_SliderChange(self QSlider, cb intptr_t, change SliderChange) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(change SliderChange), change SliderChange))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QSlider{h: self}).callVirtualBase_SliderChange, slotval1)

}

func (this *QSlider) callVirtualBase_KeyPressEvent(ev *QKeyEvent) {

	QSlider_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QSlider) OnKeyPressEvent(slot func(super func(ev *QKeyEvent), ev *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_KeyPressEvent
func miqt_exec_callback_QSlider_KeyPressEvent(self QSlider, cb intptr_t, ev *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QKeyEvent), ev *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(ev)

	gofunc((&QSlider{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QSlider) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QSlider_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSlider) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_TimerEvent
func miqt_exec_callback_QSlider_TimerEvent(self QSlider, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QSlider{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSlider) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QSlider_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QSlider) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_WheelEvent
func miqt_exec_callback_QSlider_WheelEvent(self QSlider, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QSlider{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QSlider) callVirtualBase_ChangeEvent(e *QEvent) {

	QSlider_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QSlider) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_ChangeEvent
func miqt_exec_callback_QSlider_ChangeEvent(self QSlider, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QSlider{h: self}).callVirtualBase_ChangeEvent, slotval1)

}
