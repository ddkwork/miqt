package qt

import (
	"unsafe"
)

type QDial struct {
	h          uintptr
	isSubclass bool
}

// NewQDial constructs a new QDial object.
func NewQDial(parent *QWidget) *QDial {

	ret := newQDial(QDial_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDial2 constructs a new QDial object.
func NewQDial2() *QDial {

	ret := newQDial(QDial_new2())
	ret.isSubclass = true
	return ret
}

func (this *QDial) MetaObject() *QMetaObject {
	return newQMetaObject(QDial_MetaObject(this.h))
}

func (this *QDial) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDial_Metacast(this.h, param1_Cstring))
}

func QDial_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDial_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDial) Wrapping() bool {
	return (bool)(QDial_Wrapping(this.h))
}

func (this *QDial) NotchSize() int {
	return (int)(QDial_NotchSize(this.h))
}

func (this *QDial) SetNotchTarget(target float64) {
	QDial_SetNotchTarget(this.h, (double)(target))
}

func (this *QDial) NotchTarget() float64 {
	return (float64)(QDial_NotchTarget(this.h))
}

func (this *QDial) NotchesVisible() bool {
	return (bool)(QDial_NotchesVisible(this.h))
}

func (this *QDial) SizeHint() *QSize {
	_goptr := newQSize(QDial_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDial) MinimumSizeHint() *QSize {
	_goptr := newQSize(QDial_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDial) SetNotchesVisible(visible bool) {
	QDial_SetNotchesVisible(this.h, (bool)(visible))
}

func (this *QDial) SetWrapping(on bool) {
	QDial_SetWrapping(this.h, (bool)(on))
}

func QDial_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDial_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDial_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDial_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDial) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QDial_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDial) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_SizeHint
func miqt_exec_callback_QDial_SizeHint(self QDial, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDial{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QDial) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QDial_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDial) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_MinimumSizeHint
func miqt_exec_callback_QDial_MinimumSizeHint(self QDial, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDial{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QDial) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QDial_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QDial) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_Event
func miqt_exec_callback_QDial_Event(self QDial, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QDial{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDial) callVirtualBase_ResizeEvent(re *QResizeEvent) {

	QDial_virtualbase_ResizeEvent(unsafe.Pointer(this.h), re.cPointer())

}
func (this *QDial) OnResizeEvent(slot func(super func(re *QResizeEvent), re *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_ResizeEvent
func miqt_exec_callback_QDial_ResizeEvent(self QDial, cb intptr_t, re *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(re *QResizeEvent), re *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(re)

	gofunc((&QDial{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QDial) callVirtualBase_PaintEvent(pe *QPaintEvent) {

	QDial_virtualbase_PaintEvent(unsafe.Pointer(this.h), pe.cPointer())

}
func (this *QDial) OnPaintEvent(slot func(super func(pe *QPaintEvent), pe *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_PaintEvent
func miqt_exec_callback_QDial_PaintEvent(self QDial, cb intptr_t, pe *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pe *QPaintEvent), pe *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(pe)

	gofunc((&QDial{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QDial) callVirtualBase_MousePressEvent(me *QMouseEvent) {

	QDial_virtualbase_MousePressEvent(unsafe.Pointer(this.h), me.cPointer())

}
func (this *QDial) OnMousePressEvent(slot func(super func(me *QMouseEvent), me *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_MousePressEvent
func miqt_exec_callback_QDial_MousePressEvent(self QDial, cb intptr_t, me *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(me *QMouseEvent), me *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(me)

	gofunc((&QDial{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QDial) callVirtualBase_MouseReleaseEvent(me *QMouseEvent) {

	QDial_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), me.cPointer())

}
func (this *QDial) OnMouseReleaseEvent(slot func(super func(me *QMouseEvent), me *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_MouseReleaseEvent
func miqt_exec_callback_QDial_MouseReleaseEvent(self QDial, cb intptr_t, me *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(me *QMouseEvent), me *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(me)

	gofunc((&QDial{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QDial) callVirtualBase_MouseMoveEvent(me *QMouseEvent) {

	QDial_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), me.cPointer())

}
func (this *QDial) OnMouseMoveEvent(slot func(super func(me *QMouseEvent), me *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_MouseMoveEvent
func miqt_exec_callback_QDial_MouseMoveEvent(self QDial, cb intptr_t, me *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(me *QMouseEvent), me *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(me)

	gofunc((&QDial{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QDial) callVirtualBase_SliderChange(change SliderChange) {

	QDial_virtualbase_SliderChange(unsafe.Pointer(this.h), change)

}
func (this *QDial) OnSliderChange(slot func(super func(change SliderChange), change SliderChange)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_SliderChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_SliderChange
func miqt_exec_callback_QDial_SliderChange(self QDial, cb intptr_t, change SliderChange) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(change SliderChange), change SliderChange))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QDial{h: self}).callVirtualBase_SliderChange, slotval1)

}

func (this *QDial) callVirtualBase_InitStyleOption(option *QStyleOptionSlider) {

	QDial_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QDial) OnInitStyleOption(slot func(super func(option *QStyleOptionSlider), option *QStyleOptionSlider)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_InitStyleOption
func miqt_exec_callback_QDial_InitStyleOption(self QDial, cb intptr_t, option *QStyleOptionSlider) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSlider), option *QStyleOptionSlider))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSlider(option)

	gofunc((&QDial{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QDial) callVirtualBase_KeyPressEvent(ev *QKeyEvent) {

	QDial_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QDial) OnKeyPressEvent(slot func(super func(ev *QKeyEvent), ev *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_KeyPressEvent
func miqt_exec_callback_QDial_KeyPressEvent(self QDial, cb intptr_t, ev *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QKeyEvent), ev *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(ev)

	gofunc((&QDial{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QDial) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QDial_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDial) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_TimerEvent
func miqt_exec_callback_QDial_TimerEvent(self QDial, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QDial{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QDial) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QDial_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QDial) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_WheelEvent
func miqt_exec_callback_QDial_WheelEvent(self QDial, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QDial{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QDial) callVirtualBase_ChangeEvent(e *QEvent) {

	QDial_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QDial) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_ChangeEvent
func miqt_exec_callback_QDial_ChangeEvent(self QDial, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QDial{h: self}).callVirtualBase_ChangeEvent, slotval1)

}
