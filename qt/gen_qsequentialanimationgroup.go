package qt

import (
	"unsafe"
)

type QSequentialAnimationGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQSequentialAnimationGroup constructs a new QSequentialAnimationGroup object.
func NewQSequentialAnimationGroup() *QSequentialAnimationGroup {

	ret := newQSequentialAnimationGroup(QSequentialAnimationGroup_new())
	ret.isSubclass = true
	return ret
}

// NewQSequentialAnimationGroup2 constructs a new QSequentialAnimationGroup object.
func NewQSequentialAnimationGroup2(parent *QObject) *QSequentialAnimationGroup {

	ret := newQSequentialAnimationGroup(QSequentialAnimationGroup_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSequentialAnimationGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QSequentialAnimationGroup_MetaObject(this.h))
}

func (this *QSequentialAnimationGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSequentialAnimationGroup_Metacast(this.h, param1_Cstring))
}

func QSequentialAnimationGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSequentialAnimationGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	return newQPauseAnimation(QSequentialAnimationGroup_AddPause(this.h, (int)(msecs)))
}

func (this *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation {
	return newQPauseAnimation(QSequentialAnimationGroup_InsertPause(this.h, (int)(index), (int)(msecs)))
}

func (this *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	return newQAbstractAnimation(QSequentialAnimationGroup_CurrentAnimation(this.h))
}

func (this *QSequentialAnimationGroup) Duration() int {
	return (int)(QSequentialAnimationGroup_Duration(this.h))
}

func (this *QSequentialAnimationGroup) CurrentAnimationChanged(current *QAbstractAnimation) {
	QSequentialAnimationGroup_CurrentAnimationChanged(this.h, current.cPointer())
}
func (this *QSequentialAnimationGroup) OnCurrentAnimationChanged(slot func(current *QAbstractAnimation)) {
	QSequentialAnimationGroup_connect_CurrentAnimationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_CurrentAnimationChanged
func miqt_exec_callback_QSequentialAnimationGroup_CurrentAnimationChanged(cb intptr_t, current *QAbstractAnimation) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QAbstractAnimation))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractAnimation(current)

	gofunc(slotval1)
}

func QSequentialAnimationGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSequentialAnimationGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSequentialAnimationGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSequentialAnimationGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSequentialAnimationGroup) callVirtualBase_Duration() int {

	return (int)(QSequentialAnimationGroup_virtualbase_Duration(unsafe.Pointer(this.h)))

}
func (this *QSequentialAnimationGroup) OnDuration(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSequentialAnimationGroup_override_virtual_Duration(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_Duration
func miqt_exec_callback_QSequentialAnimationGroup_Duration(self QSequentialAnimationGroup, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSequentialAnimationGroup{h: self}).callVirtualBase_Duration)

	return (int)(virtualReturn)

}

func (this *QSequentialAnimationGroup) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QSequentialAnimationGroup_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QSequentialAnimationGroup) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSequentialAnimationGroup_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_Event
func miqt_exec_callback_QSequentialAnimationGroup_Event(self QSequentialAnimationGroup, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QSequentialAnimationGroup{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSequentialAnimationGroup) callVirtualBase_UpdateCurrentTime(param1 int) {

	QSequentialAnimationGroup_virtualbase_UpdateCurrentTime(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QSequentialAnimationGroup) OnUpdateCurrentTime(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSequentialAnimationGroup_override_virtual_UpdateCurrentTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_UpdateCurrentTime
func miqt_exec_callback_QSequentialAnimationGroup_UpdateCurrentTime(self QSequentialAnimationGroup, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QSequentialAnimationGroup{h: self}).callVirtualBase_UpdateCurrentTime, slotval1)

}

func (this *QSequentialAnimationGroup) callVirtualBase_UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {

	QSequentialAnimationGroup_virtualbase_UpdateState(unsafe.Pointer(this.h), (int)(newState), (int)(oldState))

}
func (this *QSequentialAnimationGroup) OnUpdateState(slot func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSequentialAnimationGroup_override_virtual_UpdateState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_UpdateState
func miqt_exec_callback_QSequentialAnimationGroup_UpdateState(self QSequentialAnimationGroup, cb intptr_t, newState int, oldState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__State)(newState)

	slotval2 := (QAbstractAnimation__State)(oldState)

	gofunc((&QSequentialAnimationGroup{h: self}).callVirtualBase_UpdateState, slotval1, slotval2)

}

func (this *QSequentialAnimationGroup) callVirtualBase_UpdateDirection(direction QAbstractAnimation__Direction) {

	QSequentialAnimationGroup_virtualbase_UpdateDirection(unsafe.Pointer(this.h), (int)(direction))

}
func (this *QSequentialAnimationGroup) OnUpdateDirection(slot func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSequentialAnimationGroup_override_virtual_UpdateDirection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_UpdateDirection
func miqt_exec_callback_QSequentialAnimationGroup_UpdateDirection(self QSequentialAnimationGroup, cb intptr_t, direction int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__Direction)(direction)

	gofunc((&QSequentialAnimationGroup{h: self}).callVirtualBase_UpdateDirection, slotval1)

}
