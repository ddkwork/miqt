package qt

import (
	"unsafe"
)

type QPropertyAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQPropertyAnimation constructs a new QPropertyAnimation object.
func NewQPropertyAnimation() *QPropertyAnimation {

	ret := newQPropertyAnimation(QPropertyAnimation_new())
	ret.isSubclass = true
	return ret
}

// NewQPropertyAnimation2 constructs a new QPropertyAnimation object.
func NewQPropertyAnimation2(target *QObject, propertyName []byte) *QPropertyAnimation {
	propertyName_alias := struct_miqt_string{}
	propertyName_alias.data = (char)(unsafe.Pointer(&propertyName[0]))
	propertyName_alias.len = size_t(len(propertyName))

	ret := newQPropertyAnimation(QPropertyAnimation_new2(target.cPointer(), propertyName_alias))
	ret.isSubclass = true
	return ret
}

// NewQPropertyAnimation3 constructs a new QPropertyAnimation object.
func NewQPropertyAnimation3(parent *QObject) *QPropertyAnimation {

	ret := newQPropertyAnimation(QPropertyAnimation_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPropertyAnimation4 constructs a new QPropertyAnimation object.
func NewQPropertyAnimation4(target *QObject, propertyName []byte, parent *QObject) *QPropertyAnimation {
	propertyName_alias := struct_miqt_string{}
	propertyName_alias.data = (char)(unsafe.Pointer(&propertyName[0]))
	propertyName_alias.len = size_t(len(propertyName))

	ret := newQPropertyAnimation(QPropertyAnimation_new4(target.cPointer(), propertyName_alias, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPropertyAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QPropertyAnimation_MetaObject(this.h))
}

func (this *QPropertyAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPropertyAnimation_Metacast(this.h, param1_Cstring))
}

func QPropertyAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPropertyAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPropertyAnimation) TargetObject() *QObject {
	return newQObject(QPropertyAnimation_TargetObject(this.h))
}

func (this *QPropertyAnimation) SetTargetObject(target *QObject) {
	QPropertyAnimation_SetTargetObject(this.h, target.cPointer())
}

func (this *QPropertyAnimation) PropertyName() []byte {
	var _bytearray struct_miqt_string = QPropertyAnimation_PropertyName(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QPropertyAnimation) SetPropertyName(propertyName []byte) {
	propertyName_alias := struct_miqt_string{}
	propertyName_alias.data = (char)(unsafe.Pointer(&propertyName[0]))
	propertyName_alias.len = size_t(len(propertyName))
	QPropertyAnimation_SetPropertyName(this.h, propertyName_alias)
}

func QPropertyAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPropertyAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPropertyAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPropertyAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPropertyAnimation) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QPropertyAnimation_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QPropertyAnimation) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_Event
func miqt_exec_callback_QPropertyAnimation_Event(self QPropertyAnimation, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QPropertyAnimation{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QPropertyAnimation) callVirtualBase_UpdateCurrentValue(value *QVariant) {

	QPropertyAnimation_virtualbase_UpdateCurrentValue(unsafe.Pointer(this.h), value.cPointer())

}
func (this *QPropertyAnimation) OnUpdateCurrentValue(slot func(super func(value *QVariant), value *QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_UpdateCurrentValue(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_UpdateCurrentValue
func miqt_exec_callback_QPropertyAnimation_UpdateCurrentValue(self QPropertyAnimation, cb intptr_t, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value *QVariant), value *QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(value)

	gofunc((&QPropertyAnimation{h: self}).callVirtualBase_UpdateCurrentValue, slotval1)

}

func (this *QPropertyAnimation) callVirtualBase_UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {

	QPropertyAnimation_virtualbase_UpdateState(unsafe.Pointer(this.h), (int)(newState), (int)(oldState))

}
func (this *QPropertyAnimation) OnUpdateState(slot func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_UpdateState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_UpdateState
func miqt_exec_callback_QPropertyAnimation_UpdateState(self QPropertyAnimation, cb intptr_t, newState int, oldState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__State)(newState)

	slotval2 := (QAbstractAnimation__State)(oldState)

	gofunc((&QPropertyAnimation{h: self}).callVirtualBase_UpdateState, slotval1, slotval2)

}

func (this *QPropertyAnimation) callVirtualBase_Duration() int {

	return (int)(QPropertyAnimation_virtualbase_Duration(unsafe.Pointer(this.h)))

}
func (this *QPropertyAnimation) OnDuration(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_Duration(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_Duration
func miqt_exec_callback_QPropertyAnimation_Duration(self QPropertyAnimation, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPropertyAnimation{h: self}).callVirtualBase_Duration)

	return (int)(virtualReturn)

}

func (this *QPropertyAnimation) callVirtualBase_UpdateCurrentTime(param1 int) {

	QPropertyAnimation_virtualbase_UpdateCurrentTime(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QPropertyAnimation) OnUpdateCurrentTime(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_UpdateCurrentTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_UpdateCurrentTime
func miqt_exec_callback_QPropertyAnimation_UpdateCurrentTime(self QPropertyAnimation, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QPropertyAnimation{h: self}).callVirtualBase_UpdateCurrentTime, slotval1)

}

func (this *QPropertyAnimation) callVirtualBase_Interpolated(from *QVariant, to *QVariant, progress float64) *QVariant {

	_goptr := newQVariant(QPropertyAnimation_virtualbase_Interpolated(unsafe.Pointer(this.h), from.cPointer(), to.cPointer(), (double)(progress)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPropertyAnimation) OnInterpolated(slot func(super func(from *QVariant, to *QVariant, progress float64) *QVariant, from *QVariant, to *QVariant, progress float64) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_Interpolated(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_Interpolated
func miqt_exec_callback_QPropertyAnimation_Interpolated(self QPropertyAnimation, cb intptr_t, from *QVariant, to *QVariant, progress double) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(from *QVariant, to *QVariant, progress float64) *QVariant, from *QVariant, to *QVariant, progress float64) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(from)

	slotval2 := newQVariant(to)

	slotval3 := (float64)(progress)

	virtualReturn := gofunc((&QPropertyAnimation{h: self}).callVirtualBase_Interpolated, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}
