package qt

import (
	"unsafe"
)

type QEventLoop__ProcessEventsFlag int

const (
	QEventLoop__AllEvents              QEventLoop__ProcessEventsFlag = 0
	QEventLoop__ExcludeUserInputEvents QEventLoop__ProcessEventsFlag = 1
	QEventLoop__ExcludeSocketNotifiers QEventLoop__ProcessEventsFlag = 2
	QEventLoop__WaitForMoreEvents      QEventLoop__ProcessEventsFlag = 4
	QEventLoop__X11ExcludeTimers       QEventLoop__ProcessEventsFlag = 8
	QEventLoop__EventLoopExec          QEventLoop__ProcessEventsFlag = 32
	QEventLoop__DialogExec             QEventLoop__ProcessEventsFlag = 64
	QEventLoop__ApplicationExec        QEventLoop__ProcessEventsFlag = 128
)

type QEventLoop struct {
	h          uintptr
	isSubclass bool
}

// NewQEventLoop constructs a new QEventLoop object.
func NewQEventLoop() *QEventLoop {

	ret := newQEventLoop(QEventLoop_new())
	ret.isSubclass = true
	return ret
}

// NewQEventLoop2 constructs a new QEventLoop object.
func NewQEventLoop2(parent *QObject) *QEventLoop {

	ret := newQEventLoop(QEventLoop_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QEventLoop) MetaObject() *QMetaObject {
	return newQMetaObject(QEventLoop_MetaObject(this.h))
}

func (this *QEventLoop) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QEventLoop_Metacast(this.h, param1_Cstring))
}

func QEventLoop_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QEventLoop_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QEventLoop) ProcessEvents() bool {
	return (bool)(QEventLoop_ProcessEvents(this.h))
}

func (this *QEventLoop) ProcessEvents2(flags ProcessEventsFlags, maximumTime int) {
	QEventLoop_ProcessEvents2(this.h, flags, (int)(maximumTime))
}

func (this *QEventLoop) ProcessEvents3(flags ProcessEventsFlags, deadline QDeadlineTimer) {
	QEventLoop_ProcessEvents3(this.h, flags, deadline.cPointer())
}

func (this *QEventLoop) Exec() int {
	return (int)(QEventLoop_Exec(this.h))
}

func (this *QEventLoop) IsRunning() bool {
	return (bool)(QEventLoop_IsRunning(this.h))
}

func (this *QEventLoop) WakeUp() {
	QEventLoop_WakeUp(this.h)
}

func (this *QEventLoop) Event(event *QEvent) bool {
	return (bool)(QEventLoop_Event(this.h, event.cPointer()))
}

func (this *QEventLoop) Exit() {
	QEventLoop_Exit(this.h)
}

func (this *QEventLoop) Quit() {
	QEventLoop_Quit(this.h)
}

func QEventLoop_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QEventLoop_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QEventLoop_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QEventLoop_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QEventLoop) ProcessEvents1(flags ProcessEventsFlags) bool {
	return (bool)(QEventLoop_ProcessEvents1(this.h, flags))
}

func (this *QEventLoop) Exec1(flags ProcessEventsFlags) int {
	return (int)(QEventLoop_Exec1(this.h, flags))
}

func (this *QEventLoop) Exit1(returnCode int) {
	QEventLoop_Exit1(this.h, (int)(returnCode))
}

func (this *QEventLoop) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QEventLoop_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QEventLoop) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_Event
func miqt_exec_callback_QEventLoop_Event(self QEventLoop, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QEventLoop{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QEventLoop) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QEventLoop_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QEventLoop) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_EventFilter
func miqt_exec_callback_QEventLoop_EventFilter(self QEventLoop, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QEventLoop{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QEventLoop) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QEventLoop_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QEventLoop) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_TimerEvent
func miqt_exec_callback_QEventLoop_TimerEvent(self QEventLoop, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QEventLoop{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QEventLoop) callVirtualBase_ChildEvent(event *QChildEvent) {

	QEventLoop_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QEventLoop) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_ChildEvent
func miqt_exec_callback_QEventLoop_ChildEvent(self QEventLoop, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QEventLoop{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QEventLoop) callVirtualBase_CustomEvent(event *QEvent) {

	QEventLoop_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QEventLoop) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_CustomEvent
func miqt_exec_callback_QEventLoop_CustomEvent(self QEventLoop, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QEventLoop{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QEventLoop) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QEventLoop_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QEventLoop) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_ConnectNotify
func miqt_exec_callback_QEventLoop_ConnectNotify(self QEventLoop, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QEventLoop{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QEventLoop) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QEventLoop_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QEventLoop) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_DisconnectNotify
func miqt_exec_callback_QEventLoop_DisconnectNotify(self QEventLoop, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QEventLoop{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

type QEventLoopLocker struct {
	h          uintptr
	isSubclass bool
}

// NewQEventLoopLocker constructs a new QEventLoopLocker object.
func NewQEventLoopLocker() *QEventLoopLocker {

	ret := newQEventLoopLocker(QEventLoopLocker_new())
	ret.isSubclass = true
	return ret
}

// NewQEventLoopLocker2 constructs a new QEventLoopLocker object.
func NewQEventLoopLocker2(loop *QEventLoop) *QEventLoopLocker {

	ret := newQEventLoopLocker(QEventLoopLocker_new2(loop.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQEventLoopLocker3 constructs a new QEventLoopLocker object.
func NewQEventLoopLocker3(thread *QThread) *QEventLoopLocker {

	ret := newQEventLoopLocker(QEventLoopLocker_new3(thread.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QEventLoopLocker) Swap(other *QEventLoopLocker) {
	QEventLoopLocker_Swap(this.h, other.cPointer())
}
