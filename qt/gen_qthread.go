package qt

import (
	"unsafe"
)

type QThread__Priority int

const (
	QThread__IdlePriority         QThread__Priority = 0
	QThread__LowestPriority       QThread__Priority = 1
	QThread__LowPriority          QThread__Priority = 2
	QThread__NormalPriority       QThread__Priority = 3
	QThread__HighPriority         QThread__Priority = 4
	QThread__HighestPriority      QThread__Priority = 5
	QThread__TimeCriticalPriority QThread__Priority = 6
	QThread__InheritPriority      QThread__Priority = 7
)

type QThread__QualityOfService int

const (
	QThread__Auto QThread__QualityOfService = 0
	QThread__High QThread__QualityOfService = 1
	QThread__Eco  QThread__QualityOfService = 2
)

type QThread struct {
	h          uintptr
	isSubclass bool
}

// NewQThread constructs a new QThread object.
func NewQThread() *QThread {

	ret := newQThread(QThread_new())
	ret.isSubclass = true
	return ret
}

// NewQThread2 constructs a new QThread object.
func NewQThread2(parent *QObject) *QThread {

	ret := newQThread(QThread_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QThread) MetaObject() *QMetaObject {
	return newQMetaObject(QThread_MetaObject(this.h))
}

func (this *QThread) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QThread_Metacast(this.h, param1_Cstring))
}

func QThread_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QThread_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThread_CurrentThreadId() unsafe.Pointer {
	return (unsafe.Pointer)(QThread_CurrentThreadId())
}

func QThread_CurrentThread() *QThread {
	return newQThread(QThread_CurrentThread())
}

func QThread_IsMainThread() bool {
	return (bool)(QThread_IsMainThread())
}

func QThread_IdealThreadCount() int {
	return (int)(QThread_IdealThreadCount())
}

func QThread_YieldCurrentThread() {
	QThread_YieldCurrentThread()
}

func (this *QThread) SetPriority(priority Priority) {
	QThread_SetPriority(this.h, priority)
}

func (this *QThread) Priority() Priority {
	xxxxxxxxx
}

func (this *QThread) IsFinished() bool {
	return (bool)(QThread_IsFinished(this.h))
}

func (this *QThread) IsRunning() bool {
	return (bool)(QThread_IsRunning(this.h))
}

func (this *QThread) RequestInterruption() {
	QThread_RequestInterruption(this.h)
}

func (this *QThread) IsInterruptionRequested() bool {
	return (bool)(QThread_IsInterruptionRequested(this.h))
}

func (this *QThread) SetStackSize(stackSize uint) {
	QThread_SetStackSize(this.h, (uint)(stackSize))
}

func (this *QThread) StackSize() uint {
	return (uint)(QThread_StackSize(this.h))
}

func (this *QThread) EventDispatcher() *QAbstractEventDispatcher {
	return newQAbstractEventDispatcher(QThread_EventDispatcher(this.h))
}

func (this *QThread) SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher) {
	QThread_SetEventDispatcher(this.h, eventDispatcher.cPointer())
}

func (this *QThread) Event(event *QEvent) bool {
	return (bool)(QThread_Event(this.h, event.cPointer()))
}

func (this *QThread) LoopLevel() int {
	return (int)(QThread_LoopLevel(this.h))
}

func (this *QThread) IsCurrentThread() bool {
	return (bool)(QThread_IsCurrentThread(this.h))
}

func (this *QThread) SetServiceLevel(serviceLevel QualityOfService) {
	QThread_SetServiceLevel(this.h, serviceLevel)
}

func (this *QThread) ServiceLevel() QualityOfService {
	xxxxxxxxx
}

func (this *QThread) Start() {
	QThread_Start(this.h)
}

func (this *QThread) Terminate() {
	QThread_Terminate(this.h)
}

func (this *QThread) Exit() {
	QThread_Exit(this.h)
}

func (this *QThread) Quit() {
	QThread_Quit(this.h)
}

func (this *QThread) Wait() bool {
	return (bool)(QThread_Wait(this.h))
}

func (this *QThread) WaitWithTime(time uint32) bool {
	return (bool)(QThread_WaitWithTime(this.h, (ulong)(time)))
}

func QThread_Sleep(param1 uint32) {
	QThread_Sleep((ulong)(param1))
}

func QThread_Msleep(param1 uint32) {
	QThread_Msleep((ulong)(param1))
}

func QThread_Usleep(param1 uint32) {
	QThread_Usleep((ulong)(param1))
}

func QThread_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThread_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThread_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThread_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QThread) Start1(param1 Priority) {
	QThread_Start1(this.h, param1)
}

func (this *QThread) Exit1(retcode int) {
	QThread_Exit1(this.h, (int)(retcode))
}

func (this *QThread) Wait1(deadline QDeadlineTimer) bool {
	return (bool)(QThread_Wait1(this.h, deadline.cPointer()))
}

func (this *QThread) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QThread_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QThread) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_Event
func miqt_exec_callback_QThread_Event(self QThread, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QThread{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QThread) callVirtualBase_Run() {

	QThread_virtualbase_Run(unsafe.Pointer(this.h))

}
func (this *QThread) OnRun(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_Run(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_Run
func miqt_exec_callback_QThread_Run(self QThread, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QThread{h: self}).callVirtualBase_Run)

}

func (this *QThread) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QThread_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QThread) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_EventFilter
func miqt_exec_callback_QThread_EventFilter(self QThread, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QThread{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QThread) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QThread_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QThread) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_TimerEvent
func miqt_exec_callback_QThread_TimerEvent(self QThread, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QThread{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QThread) callVirtualBase_ChildEvent(event *QChildEvent) {

	QThread_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QThread) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_ChildEvent
func miqt_exec_callback_QThread_ChildEvent(self QThread, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QThread{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QThread) callVirtualBase_CustomEvent(event *QEvent) {

	QThread_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QThread) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_CustomEvent
func miqt_exec_callback_QThread_CustomEvent(self QThread, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QThread{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QThread) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QThread_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QThread) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_ConnectNotify
func miqt_exec_callback_QThread_ConnectNotify(self QThread, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QThread{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QThread) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QThread_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QThread) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_DisconnectNotify
func miqt_exec_callback_QThread_DisconnectNotify(self QThread, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QThread{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
