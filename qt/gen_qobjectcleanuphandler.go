package qt

import (
	"unsafe"
)

type QObjectCleanupHandler struct {
	h          uintptr
	isSubclass bool
}

// NewQObjectCleanupHandler constructs a new QObjectCleanupHandler object.
func NewQObjectCleanupHandler() *QObjectCleanupHandler {

	ret := newQObjectCleanupHandler(QObjectCleanupHandler_new())
	ret.isSubclass = true
	return ret
}

func (this *QObjectCleanupHandler) MetaObject() *QMetaObject {
	return newQMetaObject(QObjectCleanupHandler_MetaObject(this.h))
}

func (this *QObjectCleanupHandler) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QObjectCleanupHandler_Metacast(this.h, param1_Cstring))
}

func QObjectCleanupHandler_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QObjectCleanupHandler_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObjectCleanupHandler) Add(object *QObject) *QObject {
	return newQObject(QObjectCleanupHandler_Add(this.h, object.cPointer()))
}

func (this *QObjectCleanupHandler) Remove(object *QObject) {
	QObjectCleanupHandler_Remove(this.h, object.cPointer())
}

func (this *QObjectCleanupHandler) IsEmpty() bool {
	return (bool)(QObjectCleanupHandler_IsEmpty(this.h))
}

func (this *QObjectCleanupHandler) Clear() {
	QObjectCleanupHandler_Clear(this.h)
}

func QObjectCleanupHandler_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QObjectCleanupHandler_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QObjectCleanupHandler_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QObjectCleanupHandler_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObjectCleanupHandler) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QObjectCleanupHandler_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QObjectCleanupHandler) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_Event
func miqt_exec_callback_QObjectCleanupHandler_Event(self QObjectCleanupHandler, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QObjectCleanupHandler) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QObjectCleanupHandler_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QObjectCleanupHandler) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_EventFilter
func miqt_exec_callback_QObjectCleanupHandler_EventFilter(self QObjectCleanupHandler, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QObjectCleanupHandler) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QObjectCleanupHandler_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QObjectCleanupHandler) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_TimerEvent
func miqt_exec_callback_QObjectCleanupHandler_TimerEvent(self QObjectCleanupHandler, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QObjectCleanupHandler) callVirtualBase_ChildEvent(event *QChildEvent) {

	QObjectCleanupHandler_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QObjectCleanupHandler) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_ChildEvent
func miqt_exec_callback_QObjectCleanupHandler_ChildEvent(self QObjectCleanupHandler, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QObjectCleanupHandler) callVirtualBase_CustomEvent(event *QEvent) {

	QObjectCleanupHandler_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QObjectCleanupHandler) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_CustomEvent
func miqt_exec_callback_QObjectCleanupHandler_CustomEvent(self QObjectCleanupHandler, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QObjectCleanupHandler) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QObjectCleanupHandler_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QObjectCleanupHandler) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_ConnectNotify
func miqt_exec_callback_QObjectCleanupHandler_ConnectNotify(self QObjectCleanupHandler, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QObjectCleanupHandler) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QObjectCleanupHandler_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QObjectCleanupHandler) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_DisconnectNotify
func miqt_exec_callback_QObjectCleanupHandler_DisconnectNotify(self QObjectCleanupHandler, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
