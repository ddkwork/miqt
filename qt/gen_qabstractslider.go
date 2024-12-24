package qt

import (
	"unsafe"
)

type QAbstractSlider__SliderAction int

const (
	QAbstractSlider__SliderNoAction      QAbstractSlider__SliderAction = 0
	QAbstractSlider__SliderSingleStepAdd QAbstractSlider__SliderAction = 1
	QAbstractSlider__SliderSingleStepSub QAbstractSlider__SliderAction = 2
	QAbstractSlider__SliderPageStepAdd   QAbstractSlider__SliderAction = 3
	QAbstractSlider__SliderPageStepSub   QAbstractSlider__SliderAction = 4
	QAbstractSlider__SliderToMinimum     QAbstractSlider__SliderAction = 5
	QAbstractSlider__SliderToMaximum     QAbstractSlider__SliderAction = 6
	QAbstractSlider__SliderMove          QAbstractSlider__SliderAction = 7
)

type QAbstractSlider__SliderChange int

const (
	QAbstractSlider__SliderRangeChange       QAbstractSlider__SliderChange = 0
	QAbstractSlider__SliderOrientationChange QAbstractSlider__SliderChange = 1
	QAbstractSlider__SliderStepsChange       QAbstractSlider__SliderChange = 2
	QAbstractSlider__SliderValueChange       QAbstractSlider__SliderChange = 3
)

type QAbstractSlider struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractSlider constructs a new QAbstractSlider object.
func NewQAbstractSlider(parent *QWidget) *QAbstractSlider {

	ret := newQAbstractSlider(QAbstractSlider_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAbstractSlider2 constructs a new QAbstractSlider object.
func NewQAbstractSlider2() *QAbstractSlider {

	ret := newQAbstractSlider(QAbstractSlider_new2())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractSlider) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractSlider_MetaObject(this.h))
}

func (this *QAbstractSlider) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractSlider_Metacast(this.h, param1_Cstring))
}

func QAbstractSlider_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractSlider_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSlider) Orientation() Orientation {
	return (Orientation)(QAbstractSlider_Orientation(this.h))
}

func (this *QAbstractSlider) SetMinimum(minimum int) {
	QAbstractSlider_SetMinimum(this.h, (int)(minimum))
}

func (this *QAbstractSlider) Minimum() int {
	return (int)(QAbstractSlider_Minimum(this.h))
}

func (this *QAbstractSlider) SetMaximum(maximum int) {
	QAbstractSlider_SetMaximum(this.h, (int)(maximum))
}

func (this *QAbstractSlider) Maximum() int {
	return (int)(QAbstractSlider_Maximum(this.h))
}

func (this *QAbstractSlider) SetSingleStep(singleStep int) {
	QAbstractSlider_SetSingleStep(this.h, (int)(singleStep))
}

func (this *QAbstractSlider) SingleStep() int {
	return (int)(QAbstractSlider_SingleStep(this.h))
}

func (this *QAbstractSlider) SetPageStep(pageStep int) {
	QAbstractSlider_SetPageStep(this.h, (int)(pageStep))
}

func (this *QAbstractSlider) PageStep() int {
	return (int)(QAbstractSlider_PageStep(this.h))
}

func (this *QAbstractSlider) SetTracking(enable bool) {
	QAbstractSlider_SetTracking(this.h, (bool)(enable))
}

func (this *QAbstractSlider) HasTracking() bool {
	return (bool)(QAbstractSlider_HasTracking(this.h))
}

func (this *QAbstractSlider) SetSliderDown(sliderDown bool) {
	QAbstractSlider_SetSliderDown(this.h, (bool)(sliderDown))
}

func (this *QAbstractSlider) IsSliderDown() bool {
	return (bool)(QAbstractSlider_IsSliderDown(this.h))
}

func (this *QAbstractSlider) SetSliderPosition(sliderPosition int) {
	QAbstractSlider_SetSliderPosition(this.h, (int)(sliderPosition))
}

func (this *QAbstractSlider) SliderPosition() int {
	return (int)(QAbstractSlider_SliderPosition(this.h))
}

func (this *QAbstractSlider) SetInvertedAppearance(invertedAppearance bool) {
	QAbstractSlider_SetInvertedAppearance(this.h, (bool)(invertedAppearance))
}

func (this *QAbstractSlider) InvertedAppearance() bool {
	return (bool)(QAbstractSlider_InvertedAppearance(this.h))
}

func (this *QAbstractSlider) SetInvertedControls(invertedControls bool) {
	QAbstractSlider_SetInvertedControls(this.h, (bool)(invertedControls))
}

func (this *QAbstractSlider) InvertedControls() bool {
	return (bool)(QAbstractSlider_InvertedControls(this.h))
}

func (this *QAbstractSlider) Value() int {
	return (int)(QAbstractSlider_Value(this.h))
}

func (this *QAbstractSlider) TriggerAction(action SliderAction) {
	QAbstractSlider_TriggerAction(this.h, action)
}

func (this *QAbstractSlider) SetValue(value int) {
	QAbstractSlider_SetValue(this.h, (int)(value))
}

func (this *QAbstractSlider) SetOrientation(orientation Orientation) {
	QAbstractSlider_SetOrientation(this.h, (int)(orientation))
}

func (this *QAbstractSlider) SetRange(min int, max int) {
	QAbstractSlider_SetRange(this.h, (int)(min), (int)(max))
}

func (this *QAbstractSlider) ValueChanged(value int) {
	QAbstractSlider_ValueChanged(this.h, (int)(value))
}
func (this *QAbstractSlider) OnValueChanged(slot func(value int)) {
	QAbstractSlider_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_ValueChanged
func miqt_exec_callback_QAbstractSlider_ValueChanged(cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc(slotval1)
}

func (this *QAbstractSlider) SliderPressed() {
	QAbstractSlider_SliderPressed(this.h)
}
func (this *QAbstractSlider) OnSliderPressed(slot func()) {
	QAbstractSlider_connect_SliderPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_SliderPressed
func miqt_exec_callback_QAbstractSlider_SliderPressed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractSlider) SliderMoved(position int) {
	QAbstractSlider_SliderMoved(this.h, (int)(position))
}
func (this *QAbstractSlider) OnSliderMoved(slot func(position int)) {
	QAbstractSlider_connect_SliderMoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_SliderMoved
func miqt_exec_callback_QAbstractSlider_SliderMoved(cb intptr_t, position int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(position int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(position)

	gofunc(slotval1)
}

func (this *QAbstractSlider) SliderReleased() {
	QAbstractSlider_SliderReleased(this.h)
}
func (this *QAbstractSlider) OnSliderReleased(slot func()) {
	QAbstractSlider_connect_SliderReleased(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_SliderReleased
func miqt_exec_callback_QAbstractSlider_SliderReleased(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractSlider) RangeChanged(min int, max int) {
	QAbstractSlider_RangeChanged(this.h, (int)(min), (int)(max))
}
func (this *QAbstractSlider) OnRangeChanged(slot func(min int, max int)) {
	QAbstractSlider_connect_RangeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_RangeChanged
func miqt_exec_callback_QAbstractSlider_RangeChanged(cb intptr_t, min int, max int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(min int, max int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(min)

	slotval2 := (int)(max)

	gofunc(slotval1, slotval2)
}

func (this *QAbstractSlider) ActionTriggered(action int) {
	QAbstractSlider_ActionTriggered(this.h, (int)(action))
}
func (this *QAbstractSlider) OnActionTriggered(slot func(action int)) {
	QAbstractSlider_connect_ActionTriggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_ActionTriggered
func miqt_exec_callback_QAbstractSlider_ActionTriggered(cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc(slotval1)
}

func QAbstractSlider_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractSlider_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractSlider_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractSlider_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSlider) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QAbstractSlider_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QAbstractSlider) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_Event
func miqt_exec_callback_QAbstractSlider_Event(self QAbstractSlider, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSlider) callVirtualBase_SliderChange(change SliderChange) {

	QAbstractSlider_virtualbase_SliderChange(unsafe.Pointer(this.h), change)

}
func (this *QAbstractSlider) OnSliderChange(slot func(super func(change SliderChange), change SliderChange)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_SliderChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_SliderChange
func miqt_exec_callback_QAbstractSlider_SliderChange(self QAbstractSlider, cb intptr_t, change SliderChange) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(change SliderChange), change SliderChange))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_SliderChange, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_KeyPressEvent(ev *QKeyEvent) {

	QAbstractSlider_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QAbstractSlider) OnKeyPressEvent(slot func(super func(ev *QKeyEvent), ev *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_KeyPressEvent
func miqt_exec_callback_QAbstractSlider_KeyPressEvent(self QAbstractSlider, cb intptr_t, ev *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QKeyEvent), ev *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(ev)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QAbstractSlider_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractSlider) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_TimerEvent
func miqt_exec_callback_QAbstractSlider_TimerEvent(self QAbstractSlider, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QAbstractSlider_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractSlider) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_WheelEvent
func miqt_exec_callback_QAbstractSlider_WheelEvent(self QAbstractSlider, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_ChangeEvent(e *QEvent) {

	QAbstractSlider_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractSlider) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_ChangeEvent
func miqt_exec_callback_QAbstractSlider_ChangeEvent(self QAbstractSlider, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_DevType() int {

	return (int)(QAbstractSlider_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QAbstractSlider) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_DevType
func miqt_exec_callback_QAbstractSlider_DevType(self QAbstractSlider, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QAbstractSlider) callVirtualBase_SetVisible(visible bool) {

	QAbstractSlider_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QAbstractSlider) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_SetVisible
func miqt_exec_callback_QAbstractSlider_SetVisible(self QAbstractSlider, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QAbstractSlider_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractSlider) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_SizeHint
func miqt_exec_callback_QAbstractSlider_SizeHint(self QAbstractSlider, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractSlider) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QAbstractSlider_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractSlider) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_MinimumSizeHint
func miqt_exec_callback_QAbstractSlider_MinimumSizeHint(self QAbstractSlider, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractSlider) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QAbstractSlider_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QAbstractSlider) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_HeightForWidth
func miqt_exec_callback_QAbstractSlider_HeightForWidth(self QAbstractSlider, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractSlider) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QAbstractSlider_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QAbstractSlider) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_HasHeightForWidth
func miqt_exec_callback_QAbstractSlider_HasHeightForWidth(self QAbstractSlider, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QAbstractSlider) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QAbstractSlider_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QAbstractSlider) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_PaintEngine
func miqt_exec_callback_QAbstractSlider_PaintEngine(self QAbstractSlider, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QAbstractSlider) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QAbstractSlider_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_MousePressEvent
func miqt_exec_callback_QAbstractSlider_MousePressEvent(self QAbstractSlider, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QAbstractSlider_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_MouseReleaseEvent
func miqt_exec_callback_QAbstractSlider_MouseReleaseEvent(self QAbstractSlider, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QAbstractSlider_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_MouseDoubleClickEvent
func miqt_exec_callback_QAbstractSlider_MouseDoubleClickEvent(self QAbstractSlider, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QAbstractSlider_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_MouseMoveEvent
func miqt_exec_callback_QAbstractSlider_MouseMoveEvent(self QAbstractSlider, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QAbstractSlider_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_KeyReleaseEvent
func miqt_exec_callback_QAbstractSlider_KeyReleaseEvent(self QAbstractSlider, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QAbstractSlider_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_FocusInEvent
func miqt_exec_callback_QAbstractSlider_FocusInEvent(self QAbstractSlider, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QAbstractSlider_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_FocusOutEvent
func miqt_exec_callback_QAbstractSlider_FocusOutEvent(self QAbstractSlider, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QAbstractSlider_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_EnterEvent
func miqt_exec_callback_QAbstractSlider_EnterEvent(self QAbstractSlider, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_LeaveEvent(event *QEvent) {

	QAbstractSlider_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_LeaveEvent
func miqt_exec_callback_QAbstractSlider_LeaveEvent(self QAbstractSlider, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QAbstractSlider_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_PaintEvent
func miqt_exec_callback_QAbstractSlider_PaintEvent(self QAbstractSlider, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QAbstractSlider_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_MoveEvent
func miqt_exec_callback_QAbstractSlider_MoveEvent(self QAbstractSlider, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QAbstractSlider_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_ResizeEvent
func miqt_exec_callback_QAbstractSlider_ResizeEvent(self QAbstractSlider, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QAbstractSlider_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_CloseEvent
func miqt_exec_callback_QAbstractSlider_CloseEvent(self QAbstractSlider, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QAbstractSlider_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_ContextMenuEvent
func miqt_exec_callback_QAbstractSlider_ContextMenuEvent(self QAbstractSlider, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QAbstractSlider_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_TabletEvent
func miqt_exec_callback_QAbstractSlider_TabletEvent(self QAbstractSlider, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_ActionEvent(event *QActionEvent) {

	QAbstractSlider_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_ActionEvent
func miqt_exec_callback_QAbstractSlider_ActionEvent(self QAbstractSlider, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QAbstractSlider_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_DragEnterEvent
func miqt_exec_callback_QAbstractSlider_DragEnterEvent(self QAbstractSlider, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QAbstractSlider_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_DragMoveEvent
func miqt_exec_callback_QAbstractSlider_DragMoveEvent(self QAbstractSlider, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QAbstractSlider_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_DragLeaveEvent
func miqt_exec_callback_QAbstractSlider_DragLeaveEvent(self QAbstractSlider, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_DropEvent(event *QDropEvent) {

	QAbstractSlider_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_DropEvent
func miqt_exec_callback_QAbstractSlider_DropEvent(self QAbstractSlider, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_ShowEvent(event *QShowEvent) {

	QAbstractSlider_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_ShowEvent
func miqt_exec_callback_QAbstractSlider_ShowEvent(self QAbstractSlider, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_HideEvent(event *QHideEvent) {

	QAbstractSlider_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSlider) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_HideEvent
func miqt_exec_callback_QAbstractSlider_HideEvent(self QAbstractSlider, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QAbstractSlider_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QAbstractSlider) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_NativeEvent
func miqt_exec_callback_QAbstractSlider_NativeEvent(self QAbstractSlider, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractSlider) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QAbstractSlider_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QAbstractSlider) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_Metric
func miqt_exec_callback_QAbstractSlider_Metric(self QAbstractSlider, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractSlider) callVirtualBase_InitPainter(painter *QPainter) {

	QAbstractSlider_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QAbstractSlider) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_InitPainter
func miqt_exec_callback_QAbstractSlider_InitPainter(self QAbstractSlider, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QAbstractSlider_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QAbstractSlider) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_Redirected
func miqt_exec_callback_QAbstractSlider_Redirected(self QAbstractSlider, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractSlider) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QAbstractSlider_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QAbstractSlider) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_SharedPainter
func miqt_exec_callback_QAbstractSlider_SharedPainter(self QAbstractSlider, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QAbstractSlider) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QAbstractSlider_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractSlider) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_InputMethodEvent
func miqt_exec_callback_QAbstractSlider_InputMethodEvent(self QAbstractSlider, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QAbstractSlider{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QAbstractSlider) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QAbstractSlider_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractSlider) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_InputMethodQuery
func miqt_exec_callback_QAbstractSlider_InputMethodQuery(self QAbstractSlider, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractSlider) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QAbstractSlider_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QAbstractSlider) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_FocusNextPrevChild
func miqt_exec_callback_QAbstractSlider_FocusNextPrevChild(self QAbstractSlider, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
