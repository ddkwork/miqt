package qt

import (
	"unsafe"
)

type QChronoTimer struct {
	h          uintptr
	isSubclass bool
}

// NewQChronoTimer constructs a new QChronoTimer object.
func NewQChronoTimer() *QChronoTimer {

	ret := newQChronoTimer(QChronoTimer_new())
	ret.isSubclass = true
	return ret
}

// NewQChronoTimer2 constructs a new QChronoTimer object.
func NewQChronoTimer2(parent *QObject) *QChronoTimer {

	ret := newQChronoTimer(QChronoTimer_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QChronoTimer) MetaObject() *QMetaObject {
	return newQMetaObject(QChronoTimer_MetaObject(this.h))
}

func (this *QChronoTimer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QChronoTimer_Metacast(this.h, param1_Cstring))
}

func QChronoTimer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QChronoTimer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QChronoTimer) IsActive() bool {
	return (bool)(QChronoTimer_IsActive(this.h))
}

func (this *QChronoTimer) Id() TimerId {
	return (TimerId)(QChronoTimer_Id(this.h))
}

func (this *QChronoTimer) SetTimerType(atype TimerType) {
	QChronoTimer_SetTimerType(this.h, (int)(atype))
}

func (this *QChronoTimer) TimerType() TimerType {
	return (TimerType)(QChronoTimer_TimerType(this.h))
}

func (this *QChronoTimer) SetSingleShot(singleShot bool) {
	QChronoTimer_SetSingleShot(this.h, (bool)(singleShot))
}

func (this *QChronoTimer) IsSingleShot() bool {
	return (bool)(QChronoTimer_IsSingleShot(this.h))
}

func (this *QChronoTimer) Start() {
	QChronoTimer_Start(this.h)
}

func (this *QChronoTimer) Stop() {
	QChronoTimer_Stop(this.h)
}

func QChronoTimer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QChronoTimer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QChronoTimer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QChronoTimer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QChronoTimer) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QChronoTimer_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QChronoTimer) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_TimerEvent
func miqt_exec_callback_QChronoTimer_TimerEvent(self QChronoTimer, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QChronoTimer{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QChronoTimer) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QChronoTimer_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QChronoTimer) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_Event
func miqt_exec_callback_QChronoTimer_Event(self QChronoTimer, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QChronoTimer{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QChronoTimer) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QChronoTimer_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QChronoTimer) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_EventFilter
func miqt_exec_callback_QChronoTimer_EventFilter(self QChronoTimer, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QChronoTimer{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QChronoTimer) callVirtualBase_ChildEvent(event *QChildEvent) {

	QChronoTimer_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QChronoTimer) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_ChildEvent
func miqt_exec_callback_QChronoTimer_ChildEvent(self QChronoTimer, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QChronoTimer{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QChronoTimer) callVirtualBase_CustomEvent(event *QEvent) {

	QChronoTimer_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QChronoTimer) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_CustomEvent
func miqt_exec_callback_QChronoTimer_CustomEvent(self QChronoTimer, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QChronoTimer{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QChronoTimer) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QChronoTimer_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QChronoTimer) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_ConnectNotify
func miqt_exec_callback_QChronoTimer_ConnectNotify(self QChronoTimer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QChronoTimer{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QChronoTimer) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QChronoTimer_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QChronoTimer) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_DisconnectNotify
func miqt_exec_callback_QChronoTimer_DisconnectNotify(self QChronoTimer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QChronoTimer{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
