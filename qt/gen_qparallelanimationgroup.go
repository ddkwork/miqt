package qt

import (
	"unsafe"
)

type QParallelAnimationGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQParallelAnimationGroup constructs a new QParallelAnimationGroup object.
func NewQParallelAnimationGroup() *QParallelAnimationGroup {

	ret := newQParallelAnimationGroup(QParallelAnimationGroup_new())
	ret.isSubclass = true
	return ret
}

// NewQParallelAnimationGroup2 constructs a new QParallelAnimationGroup object.
func NewQParallelAnimationGroup2(parent *QObject) *QParallelAnimationGroup {

	ret := newQParallelAnimationGroup(QParallelAnimationGroup_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QParallelAnimationGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QParallelAnimationGroup_MetaObject(this.h))
}

func (this *QParallelAnimationGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QParallelAnimationGroup_Metacast(this.h, param1_Cstring))
}

func QParallelAnimationGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QParallelAnimationGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QParallelAnimationGroup) Duration() int {
	return (int)(QParallelAnimationGroup_Duration(this.h))
}

func QParallelAnimationGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QParallelAnimationGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QParallelAnimationGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QParallelAnimationGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QParallelAnimationGroup) callVirtualBase_Duration() int {

	return (int)(QParallelAnimationGroup_virtualbase_Duration(unsafe.Pointer(this.h)))

}
func (this *QParallelAnimationGroup) OnDuration(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QParallelAnimationGroup_override_virtual_Duration(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QParallelAnimationGroup_Duration
func miqt_exec_callback_QParallelAnimationGroup_Duration(self QParallelAnimationGroup, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QParallelAnimationGroup{h: self}).callVirtualBase_Duration)

	return (int)(virtualReturn)

}

func (this *QParallelAnimationGroup) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QParallelAnimationGroup_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QParallelAnimationGroup) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QParallelAnimationGroup_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QParallelAnimationGroup_Event
func miqt_exec_callback_QParallelAnimationGroup_Event(self QParallelAnimationGroup, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QParallelAnimationGroup{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QParallelAnimationGroup) callVirtualBase_UpdateCurrentTime(currentTime int) {

	QParallelAnimationGroup_virtualbase_UpdateCurrentTime(unsafe.Pointer(this.h), (int)(currentTime))

}
func (this *QParallelAnimationGroup) OnUpdateCurrentTime(slot func(super func(currentTime int), currentTime int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QParallelAnimationGroup_override_virtual_UpdateCurrentTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QParallelAnimationGroup_UpdateCurrentTime
func miqt_exec_callback_QParallelAnimationGroup_UpdateCurrentTime(self QParallelAnimationGroup, cb intptr_t, currentTime int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(currentTime int), currentTime int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(currentTime)

	gofunc((&QParallelAnimationGroup{h: self}).callVirtualBase_UpdateCurrentTime, slotval1)

}

func (this *QParallelAnimationGroup) callVirtualBase_UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {

	QParallelAnimationGroup_virtualbase_UpdateState(unsafe.Pointer(this.h), (int)(newState), (int)(oldState))

}
func (this *QParallelAnimationGroup) OnUpdateState(slot func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QParallelAnimationGroup_override_virtual_UpdateState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QParallelAnimationGroup_UpdateState
func miqt_exec_callback_QParallelAnimationGroup_UpdateState(self QParallelAnimationGroup, cb intptr_t, newState int, oldState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__State)(newState)

	slotval2 := (QAbstractAnimation__State)(oldState)

	gofunc((&QParallelAnimationGroup{h: self}).callVirtualBase_UpdateState, slotval1, slotval2)

}

func (this *QParallelAnimationGroup) callVirtualBase_UpdateDirection(direction QAbstractAnimation__Direction) {

	QParallelAnimationGroup_virtualbase_UpdateDirection(unsafe.Pointer(this.h), (int)(direction))

}
func (this *QParallelAnimationGroup) OnUpdateDirection(slot func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QParallelAnimationGroup_override_virtual_UpdateDirection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QParallelAnimationGroup_UpdateDirection
func miqt_exec_callback_QParallelAnimationGroup_UpdateDirection(self QParallelAnimationGroup, cb intptr_t, direction int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__Direction)(direction)

	gofunc((&QParallelAnimationGroup{h: self}).callVirtualBase_UpdateDirection, slotval1)

}
