package qt

import (
	"unsafe"
)

type QPauseAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQPauseAnimation constructs a new QPauseAnimation object.
func NewQPauseAnimation() *QPauseAnimation {

	ret := newQPauseAnimation(QPauseAnimation_new())
	ret.isSubclass = true
	return ret
}

// NewQPauseAnimation2 constructs a new QPauseAnimation object.
func NewQPauseAnimation2(msecs int) *QPauseAnimation {

	ret := newQPauseAnimation(QPauseAnimation_new2((int)(msecs)))
	ret.isSubclass = true
	return ret
}

// NewQPauseAnimation3 constructs a new QPauseAnimation object.
func NewQPauseAnimation3(parent *QObject) *QPauseAnimation {

	ret := newQPauseAnimation(QPauseAnimation_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPauseAnimation4 constructs a new QPauseAnimation object.
func NewQPauseAnimation4(msecs int, parent *QObject) *QPauseAnimation {

	ret := newQPauseAnimation(QPauseAnimation_new4((int)(msecs), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPauseAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QPauseAnimation_MetaObject(this.h))
}

func (this *QPauseAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPauseAnimation_Metacast(this.h, param1_Cstring))
}

func QPauseAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPauseAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPauseAnimation) Duration() int {
	return (int)(QPauseAnimation_Duration(this.h))
}

func (this *QPauseAnimation) SetDuration(msecs int) {
	QPauseAnimation_SetDuration(this.h, (int)(msecs))
}

func QPauseAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPauseAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPauseAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPauseAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPauseAnimation) callVirtualBase_Duration() int {

	return (int)(QPauseAnimation_virtualbase_Duration(unsafe.Pointer(this.h)))

}
func (this *QPauseAnimation) OnDuration(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPauseAnimation_override_virtual_Duration(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPauseAnimation_Duration
func miqt_exec_callback_QPauseAnimation_Duration(self QPauseAnimation, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPauseAnimation{h: self}).callVirtualBase_Duration)

	return (int)(virtualReturn)

}

func (this *QPauseAnimation) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QPauseAnimation_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QPauseAnimation) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPauseAnimation_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPauseAnimation_Event
func miqt_exec_callback_QPauseAnimation_Event(self QPauseAnimation, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QPauseAnimation{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QPauseAnimation) callVirtualBase_UpdateCurrentTime(param1 int) {

	QPauseAnimation_virtualbase_UpdateCurrentTime(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QPauseAnimation) OnUpdateCurrentTime(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPauseAnimation_override_virtual_UpdateCurrentTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPauseAnimation_UpdateCurrentTime
func miqt_exec_callback_QPauseAnimation_UpdateCurrentTime(self QPauseAnimation, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QPauseAnimation{h: self}).callVirtualBase_UpdateCurrentTime, slotval1)

}

func (this *QPauseAnimation) callVirtualBase_UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {

	QPauseAnimation_virtualbase_UpdateState(unsafe.Pointer(this.h), (int)(newState), (int)(oldState))

}
func (this *QPauseAnimation) OnUpdateState(slot func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPauseAnimation_override_virtual_UpdateState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPauseAnimation_UpdateState
func miqt_exec_callback_QPauseAnimation_UpdateState(self QPauseAnimation, cb intptr_t, newState int, oldState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__State)(newState)

	slotval2 := (QAbstractAnimation__State)(oldState)

	gofunc((&QPauseAnimation{h: self}).callVirtualBase_UpdateState, slotval1, slotval2)

}

func (this *QPauseAnimation) callVirtualBase_UpdateDirection(direction QAbstractAnimation__Direction) {

	QPauseAnimation_virtualbase_UpdateDirection(unsafe.Pointer(this.h), (int)(direction))

}
func (this *QPauseAnimation) OnUpdateDirection(slot func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPauseAnimation_override_virtual_UpdateDirection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPauseAnimation_UpdateDirection
func miqt_exec_callback_QPauseAnimation_UpdateDirection(self QPauseAnimation, cb intptr_t, direction int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__Direction)(direction)

	gofunc((&QPauseAnimation{h: self}).callVirtualBase_UpdateDirection, slotval1)

}
