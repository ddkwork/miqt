package qt

import (
	"unsafe"
)

type QAnimationGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQAnimationGroup constructs a new QAnimationGroup object.
func NewQAnimationGroup() *QAnimationGroup {

	ret := newQAnimationGroup(QAnimationGroup_new())
	ret.isSubclass = true
	return ret
}

// NewQAnimationGroup2 constructs a new QAnimationGroup object.
func NewQAnimationGroup2(parent *QObject) *QAnimationGroup {

	ret := newQAnimationGroup(QAnimationGroup_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAnimationGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QAnimationGroup_MetaObject(this.h))
}

func (this *QAnimationGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAnimationGroup_Metacast(this.h, param1_Cstring))
}

func QAnimationGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAnimationGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAnimationGroup) AnimationAt(index int) *QAbstractAnimation {
	return newQAbstractAnimation(QAnimationGroup_AnimationAt(this.h, (int)(index)))
}

func (this *QAnimationGroup) AnimationCount() int {
	return (int)(QAnimationGroup_AnimationCount(this.h))
}

func (this *QAnimationGroup) IndexOfAnimation(animation *QAbstractAnimation) int {
	return (int)(QAnimationGroup_IndexOfAnimation(this.h, animation.cPointer()))
}

func (this *QAnimationGroup) AddAnimation(animation *QAbstractAnimation) {
	QAnimationGroup_AddAnimation(this.h, animation.cPointer())
}

func (this *QAnimationGroup) InsertAnimation(index int, animation *QAbstractAnimation) {
	QAnimationGroup_InsertAnimation(this.h, (int)(index), animation.cPointer())
}

func (this *QAnimationGroup) RemoveAnimation(animation *QAbstractAnimation) {
	QAnimationGroup_RemoveAnimation(this.h, animation.cPointer())
}

func (this *QAnimationGroup) TakeAnimation(index int) *QAbstractAnimation {
	return newQAbstractAnimation(QAnimationGroup_TakeAnimation(this.h, (int)(index)))
}

func (this *QAnimationGroup) Clear() {
	QAnimationGroup_Clear(this.h)
}

func QAnimationGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAnimationGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAnimationGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAnimationGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAnimationGroup) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QAnimationGroup_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QAnimationGroup) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAnimationGroup_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationGroup_Event
func miqt_exec_callback_QAnimationGroup_Event(self QAnimationGroup, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QAnimationGroup{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}
func (this *QAnimationGroup) OnDuration(slot func() int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAnimationGroup_override_virtual_Duration(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationGroup_Duration
func miqt_exec_callback_QAnimationGroup_Duration(self QAnimationGroup, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func() int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (int)(virtualReturn)

}
func (this *QAnimationGroup) OnUpdateCurrentTime(slot func(currentTime int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAnimationGroup_override_virtual_UpdateCurrentTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationGroup_UpdateCurrentTime
func miqt_exec_callback_QAnimationGroup_UpdateCurrentTime(self QAnimationGroup, cb intptr_t, currentTime int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(currentTime int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(currentTime)

	gofunc(slotval1)

}

func (this *QAnimationGroup) callVirtualBase_UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {

	QAnimationGroup_virtualbase_UpdateState(unsafe.Pointer(this.h), (int)(newState), (int)(oldState))

}
func (this *QAnimationGroup) OnUpdateState(slot func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAnimationGroup_override_virtual_UpdateState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationGroup_UpdateState
func miqt_exec_callback_QAnimationGroup_UpdateState(self QAnimationGroup, cb intptr_t, newState int, oldState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State), newState QAbstractAnimation__State, oldState QAbstractAnimation__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__State)(newState)

	slotval2 := (QAbstractAnimation__State)(oldState)

	gofunc((&QAnimationGroup{h: self}).callVirtualBase_UpdateState, slotval1, slotval2)

}

func (this *QAnimationGroup) callVirtualBase_UpdateDirection(direction QAbstractAnimation__Direction) {

	QAnimationGroup_virtualbase_UpdateDirection(unsafe.Pointer(this.h), (int)(direction))

}
func (this *QAnimationGroup) OnUpdateDirection(slot func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAnimationGroup_override_virtual_UpdateDirection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationGroup_UpdateDirection
func miqt_exec_callback_QAnimationGroup_UpdateDirection(self QAnimationGroup, cb intptr_t, direction int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(direction QAbstractAnimation__Direction), direction QAbstractAnimation__Direction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__Direction)(direction)

	gofunc((&QAnimationGroup{h: self}).callVirtualBase_UpdateDirection, slotval1)

}
