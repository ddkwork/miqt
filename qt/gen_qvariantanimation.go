package qt

import (
	"unsafe"
)

type QVariantAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQVariantAnimation constructs a new QVariantAnimation object.
func NewQVariantAnimation() *QVariantAnimation {

	ret := newQVariantAnimation(QVariantAnimation_new())
	ret.isSubclass = true
	return ret
}

// NewQVariantAnimation2 constructs a new QVariantAnimation object.
func NewQVariantAnimation2(parent *QObject) *QVariantAnimation {

	ret := newQVariantAnimation(QVariantAnimation_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QVariantAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QVariantAnimation_MetaObject(this.h))
}

func (this *QVariantAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QVariantAnimation_Metacast(this.h, param1_Cstring))
}

func QVariantAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QVariantAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVariantAnimation) StartValue() *QVariant {
	_goptr := newQVariant(QVariantAnimation_StartValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetStartValue(value *QVariant) {
	QVariantAnimation_SetStartValue(this.h, value.cPointer())
}

func (this *QVariantAnimation) EndValue() *QVariant {
	_goptr := newQVariant(QVariantAnimation_EndValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetEndValue(value *QVariant) {
	QVariantAnimation_SetEndValue(this.h, value.cPointer())
}

func (this *QVariantAnimation) KeyValueAt(step float64) *QVariant {
	_goptr := newQVariant(QVariantAnimation_KeyValueAt(this.h, (double)(step)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetKeyValueAt(step float64, value *QVariant) {
	QVariantAnimation_SetKeyValueAt(this.h, (double)(step), value.cPointer())
}

func (this *QVariantAnimation) KeyValues() KeyValues {
	xxxxxxxxx
}

func (this *QVariantAnimation) SetKeyValues(values *KeyValues) {
	QVariantAnimation_SetKeyValues(this.h, values)
}

func (this *QVariantAnimation) CurrentValue() *QVariant {
	_goptr := newQVariant(QVariantAnimation_CurrentValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) Duration() int {
	return (int)(QVariantAnimation_Duration(this.h))
}

func (this *QVariantAnimation) SetDuration(msecs int) {
	QVariantAnimation_SetDuration(this.h, (int)(msecs))
}

func (this *QVariantAnimation) EasingCurve() *QEasingCurve {
	_goptr := newQEasingCurve(QVariantAnimation_EasingCurve(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetEasingCurve(easing *QEasingCurve) {
	QVariantAnimation_SetEasingCurve(this.h, easing.cPointer())
}

func (this *QVariantAnimation) ValueChanged(value *QVariant) {
	QVariantAnimation_ValueChanged(this.h, value.cPointer())
}
func (this *QVariantAnimation) OnValueChanged(slot func(value *QVariant)) {
	QVariantAnimation_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_ValueChanged
func miqt_exec_callback_QVariantAnimation_ValueChanged(cb intptr_t, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value *QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(value)

	gofunc(slotval1)
}

func QVariantAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVariantAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVariantAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVariantAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVariantAnimation) callVirtualBase_Duration() int {

	return (int)(QVariantAnimation_virtualbase_Duration(unsafe.Pointer(this.h)))

}
func (this *QVariantAnimation) OnDuration(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_Duration(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_Duration
func miqt_exec_callback_QVariantAnimation_Duration(self QVariantAnimation, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVariantAnimation{h: self}).callVirtualBase_Duration)

	return (int)(virtualReturn)

}

func (this *QVariantAnimation) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QVariantAnimation_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QVariantAnimation) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_Event
func miqt_exec_callback_QVariantAnimation_Event(self QVariantAnimation, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QVariantAnimation{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QVariantAnimation) callVirtualBase_UpdateCurrentTime(param1 int) {

	QVariantAnimation_virtualbase_UpdateCurrentTime(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QVariantAnimation) OnUpdateCurrentTime(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_UpdateCurrentTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_UpdateCurrentTime
func miqt_exec_callback_QVariantAnimation_UpdateCurrentTime(self QVariantAnimation, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QVariantAnimation{h: self}).callVirtualBase_UpdateCurrentTime, slotval1)

}

func (this *QVariantAnimation) callVirtualBase_UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {

	QVariantAnimation_virtualbase_UpdateState(unsafe.Pointer(this.h), (int)(newState), (int)(oldState))

}
func (this *QVariantAnimation) OnUpdateState(slot func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_UpdateState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_UpdateState
func miqt_exec_callback_QVariantAnimation_UpdateState(self QVariantAnimation, cb intptr_t, newState int, oldState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__State)(newState)

	slotval2 := (QAbstractAnimation__State)(oldState)

	gofunc((&QVariantAnimation{h: self}).callVirtualBase_UpdateState, slotval1, slotval2)

}

func (this *QVariantAnimation) callVirtualBase_UpdateCurrentValue(value *QVariant) {

	QVariantAnimation_virtualbase_UpdateCurrentValue(unsafe.Pointer(this.h), value.cPointer())

}
func (this *QVariantAnimation) OnUpdateCurrentValue(slot func(super func(value *QVariant), value *QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_UpdateCurrentValue(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_UpdateCurrentValue
func miqt_exec_callback_QVariantAnimation_UpdateCurrentValue(self QVariantAnimation, cb intptr_t, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value *QVariant), value *QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(value)

	gofunc((&QVariantAnimation{h: self}).callVirtualBase_UpdateCurrentValue, slotval1)

}

func (this *QVariantAnimation) callVirtualBase_Interpolated(from *QVariant, to *QVariant, progress float64) *QVariant {

	_goptr := newQVariant(QVariantAnimation_virtualbase_Interpolated(unsafe.Pointer(this.h), from.cPointer(), to.cPointer(), (double)(progress)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QVariantAnimation) OnInterpolated(slot func(super func(from *QVariant, to *QVariant, progress float64) *QVariant, from *QVariant, to *QVariant, progress float64) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_Interpolated(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_Interpolated
func miqt_exec_callback_QVariantAnimation_Interpolated(self QVariantAnimation, cb intptr_t, from *QVariant, to *QVariant, progress double) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(from *QVariant, to *QVariant, progress float64) *QVariant, from *QVariant, to *QVariant, progress float64) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(from)

	slotval2 := newQVariant(to)

	slotval3 := (float64)(progress)

	virtualReturn := gofunc((&QVariantAnimation{h: self}).callVirtualBase_Interpolated, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QVariantAnimation) callVirtualBase_UpdateDirection(direction QAbstractAnimation__Direction) {

	QVariantAnimation_virtualbase_UpdateDirection(unsafe.Pointer(this.h), (int)(direction))

}
func (this *QVariantAnimation) OnUpdateDirection(slot func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_UpdateDirection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_UpdateDirection
func miqt_exec_callback_QVariantAnimation_UpdateDirection(self QVariantAnimation, cb intptr_t, direction int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__Direction)(direction)

	gofunc((&QVariantAnimation{h: self}).callVirtualBase_UpdateDirection, slotval1)

}
