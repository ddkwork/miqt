package qt

import (
	"unsafe"
)

type QThreadPool struct {
	h          uintptr
	isSubclass bool
}

// NewQThreadPool constructs a new QThreadPool object.
func NewQThreadPool() *QThreadPool {

	ret := newQThreadPool(QThreadPool_new())
	ret.isSubclass = true
	return ret
}

// NewQThreadPool2 constructs a new QThreadPool object.
func NewQThreadPool2(parent *QObject) *QThreadPool {

	ret := newQThreadPool(QThreadPool_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QThreadPool) MetaObject() *QMetaObject {
	return newQMetaObject(QThreadPool_MetaObject(this.h))
}

func (this *QThreadPool) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QThreadPool_Metacast(this.h, param1_Cstring))
}

func QThreadPool_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QThreadPool_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThreadPool_GlobalInstance() *QThreadPool {
	return newQThreadPool(QThreadPool_GlobalInstance())
}

func (this *QThreadPool) Start(runnable *QRunnable) {
	QThreadPool_Start(this.h, runnable.cPointer())
}

func (this *QThreadPool) TryStart(runnable *QRunnable) bool {
	return (bool)(QThreadPool_TryStart(this.h, runnable.cPointer()))
}

func (this *QThreadPool) StartOnReservedThread(runnable *QRunnable) {
	QThreadPool_StartOnReservedThread(this.h, runnable.cPointer())
}

func (this *QThreadPool) ExpiryTimeout() int {
	return (int)(QThreadPool_ExpiryTimeout(this.h))
}

func (this *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	QThreadPool_SetExpiryTimeout(this.h, (int)(expiryTimeout))
}

func (this *QThreadPool) MaxThreadCount() int {
	return (int)(QThreadPool_MaxThreadCount(this.h))
}

func (this *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	QThreadPool_SetMaxThreadCount(this.h, (int)(maxThreadCount))
}

func (this *QThreadPool) ActiveThreadCount() int {
	return (int)(QThreadPool_ActiveThreadCount(this.h))
}

func (this *QThreadPool) SetStackSize(stackSize uint) {
	QThreadPool_SetStackSize(this.h, (uint)(stackSize))
}

func (this *QThreadPool) StackSize() uint {
	return (uint)(QThreadPool_StackSize(this.h))
}

func (this *QThreadPool) SetThreadPriority(priority QThread__Priority) {
	QThreadPool_SetThreadPriority(this.h, (int)(priority))
}

func (this *QThreadPool) ThreadPriority() QThread__Priority {
	return (QThread__Priority)(QThreadPool_ThreadPriority(this.h))
}

func (this *QThreadPool) ReserveThread() {
	QThreadPool_ReserveThread(this.h)
}

func (this *QThreadPool) ReleaseThread() {
	QThreadPool_ReleaseThread(this.h)
}

func (this *QThreadPool) SetServiceLevel(serviceLevel QThread__QualityOfService) {
	QThreadPool_SetServiceLevel(this.h, (int)(serviceLevel))
}

func (this *QThreadPool) ServiceLevel() QThread__QualityOfService {
	return (QThread__QualityOfService)(QThreadPool_ServiceLevel(this.h))
}

func (this *QThreadPool) WaitForDone(msecs int) bool {
	return (bool)(QThreadPool_WaitForDone(this.h, (int)(msecs)))
}

func (this *QThreadPool) WaitForDone2() bool {
	return (bool)(QThreadPool_WaitForDone2(this.h))
}

func (this *QThreadPool) Clear() {
	QThreadPool_Clear(this.h)
}

func (this *QThreadPool) Contains(thread *QThread) bool {
	return (bool)(QThreadPool_Contains(this.h, thread.cPointer()))
}

func (this *QThreadPool) TryTake(runnable *QRunnable) bool {
	return (bool)(QThreadPool_TryTake(this.h, runnable.cPointer()))
}

func QThreadPool_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThreadPool_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThreadPool_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThreadPool_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QThreadPool) Start2(runnable *QRunnable, priority int) {
	QThreadPool_Start2(this.h, runnable.cPointer(), (int)(priority))
}

func (this *QThreadPool) WaitForDone1(deadline QDeadlineTimer) bool {
	return (bool)(QThreadPool_WaitForDone1(this.h, deadline.cPointer()))
}

func (this *QThreadPool) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QThreadPool_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QThreadPool) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_Event
func miqt_exec_callback_QThreadPool_Event(self QThreadPool, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QThreadPool{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QThreadPool) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QThreadPool_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QThreadPool) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_EventFilter
func miqt_exec_callback_QThreadPool_EventFilter(self QThreadPool, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QThreadPool{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QThreadPool) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QThreadPool_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QThreadPool) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_TimerEvent
func miqt_exec_callback_QThreadPool_TimerEvent(self QThreadPool, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QThreadPool{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QThreadPool) callVirtualBase_ChildEvent(event *QChildEvent) {

	QThreadPool_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QThreadPool) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_ChildEvent
func miqt_exec_callback_QThreadPool_ChildEvent(self QThreadPool, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QThreadPool{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QThreadPool) callVirtualBase_CustomEvent(event *QEvent) {

	QThreadPool_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QThreadPool) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_CustomEvent
func miqt_exec_callback_QThreadPool_CustomEvent(self QThreadPool, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QThreadPool{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QThreadPool) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QThreadPool_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QThreadPool) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_ConnectNotify
func miqt_exec_callback_QThreadPool_ConnectNotify(self QThreadPool, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QThreadPool{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QThreadPool) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QThreadPool_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QThreadPool) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_DisconnectNotify
func miqt_exec_callback_QThreadPool_DisconnectNotify(self QThreadPool, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QThreadPool{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
