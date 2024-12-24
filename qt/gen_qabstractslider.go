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

func (this *QAbstractSlider) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractSlider_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractSlider) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_MetaObject
func miqt_exec_callback_QAbstractSlider_MetaObject(self QAbstractSlider, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractSlider) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractSlider_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractSlider) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSlider_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSlider_Metacast
func miqt_exec_callback_QAbstractSlider_Metacast(self QAbstractSlider, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractSlider{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
