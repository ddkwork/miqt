package qt

import (
	"unsafe"
)

type QTimer struct {
	h          uintptr
	isSubclass bool
}

// NewQTimer constructs a new QTimer object.
func NewQTimer() *QTimer {

	ret := newQTimer(QTimer_new())
	ret.isSubclass = true
	return ret
}

// NewQTimer2 constructs a new QTimer object.
func NewQTimer2(parent *QObject) *QTimer {

	ret := newQTimer(QTimer_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTimer) MetaObject() *QMetaObject {
	return newQMetaObject(QTimer_MetaObject(this.h))
}

func (this *QTimer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTimer_Metacast(this.h, param1_Cstring))
}

func QTimer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTimer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimer) IsActive() bool {
	return (bool)(QTimer_IsActive(this.h))
}

func (this *QTimer) TimerId() int {
	return (int)(QTimer_TimerId(this.h))
}

func (this *QTimer) Id() TimerId {
	return (TimerId)(QTimer_Id(this.h))
}

func (this *QTimer) SetInterval(msec int) {
	QTimer_SetInterval(this.h, (int)(msec))
}

func (this *QTimer) Interval() int {
	return (int)(QTimer_Interval(this.h))
}

func (this *QTimer) RemainingTime() int {
	return (int)(QTimer_RemainingTime(this.h))
}

func (this *QTimer) SetTimerType(atype TimerType) {
	QTimer_SetTimerType(this.h, (int)(atype))
}

func (this *QTimer) TimerType() TimerType {
	return (TimerType)(QTimer_TimerType(this.h))
}

func (this *QTimer) SetSingleShot(singleShot bool) {
	QTimer_SetSingleShot(this.h, (bool)(singleShot))
}

func (this *QTimer) IsSingleShot() bool {
	return (bool)(QTimer_IsSingleShot(this.h))
}

func (this *QTimer) Start(msec int) {
	QTimer_Start(this.h, (int)(msec))
}

func (this *QTimer) Start2() {
	QTimer_Start2(this.h)
}

func (this *QTimer) Stop() {
	QTimer_Stop(this.h)
}

func QTimer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTimer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimer) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QTimer_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTimer) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_TimerEvent
func miqt_exec_callback_QTimer_TimerEvent(self QTimer, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QTimer{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTimer) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QTimer_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTimer) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_Event
func miqt_exec_callback_QTimer_Event(self QTimer, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTimer{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTimer) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QTimer_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QTimer) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_EventFilter
func miqt_exec_callback_QTimer_EventFilter(self QTimer, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTimer{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QTimer) callVirtualBase_ChildEvent(event *QChildEvent) {

	QTimer_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimer) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_ChildEvent
func miqt_exec_callback_QTimer_ChildEvent(self QTimer, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QTimer{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QTimer) callVirtualBase_CustomEvent(event *QEvent) {

	QTimer_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimer) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_CustomEvent
func miqt_exec_callback_QTimer_CustomEvent(self QTimer, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTimer{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QTimer) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QTimer_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTimer) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_ConnectNotify
func miqt_exec_callback_QTimer_ConnectNotify(self QTimer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTimer{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QTimer) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QTimer_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTimer) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_DisconnectNotify
func miqt_exec_callback_QTimer_DisconnectNotify(self QTimer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTimer{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
