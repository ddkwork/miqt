package qt

import (
	"unsafe"
)

type QWinEventNotifier struct {
	h          uintptr
	isSubclass bool
}

// NewQWinEventNotifier constructs a new QWinEventNotifier object.
func NewQWinEventNotifier() *QWinEventNotifier {

	ret := newQWinEventNotifier(QWinEventNotifier_new())
	ret.isSubclass = true
	return ret
}

// NewQWinEventNotifier2 constructs a new QWinEventNotifier object.
func NewQWinEventNotifier2(hEvent HANDLE) *QWinEventNotifier {

	ret := newQWinEventNotifier(QWinEventNotifier_new2(hEvent))
	ret.isSubclass = true
	return ret
}

// NewQWinEventNotifier3 constructs a new QWinEventNotifier object.
func NewQWinEventNotifier3(parent *QObject) *QWinEventNotifier {

	ret := newQWinEventNotifier(QWinEventNotifier_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQWinEventNotifier4 constructs a new QWinEventNotifier object.
func NewQWinEventNotifier4(hEvent HANDLE, parent *QObject) *QWinEventNotifier {

	ret := newQWinEventNotifier(QWinEventNotifier_new4(hEvent, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QWinEventNotifier) MetaObject() *QMetaObject {
	return newQMetaObject(QWinEventNotifier_MetaObject(this.h))
}

func (this *QWinEventNotifier) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWinEventNotifier_Metacast(this.h, param1_Cstring))
}

func QWinEventNotifier_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWinEventNotifier_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWinEventNotifier) SetHandle(hEvent HANDLE) {
	QWinEventNotifier_SetHandle(this.h, hEvent)
}

func (this *QWinEventNotifier) Handle() HANDLE {
	xxxxxxxxx
}

func (this *QWinEventNotifier) IsEnabled() bool {
	return (bool)(QWinEventNotifier_IsEnabled(this.h))
}

func (this *QWinEventNotifier) SetEnabled(enable bool) {
	QWinEventNotifier_SetEnabled(this.h, (bool)(enable))
}

func QWinEventNotifier_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWinEventNotifier_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWinEventNotifier_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWinEventNotifier_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWinEventNotifier) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QWinEventNotifier_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QWinEventNotifier) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_Event
func miqt_exec_callback_QWinEventNotifier_Event(self QWinEventNotifier, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QWinEventNotifier{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QWinEventNotifier) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QWinEventNotifier_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QWinEventNotifier) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_EventFilter
func miqt_exec_callback_QWinEventNotifier_EventFilter(self QWinEventNotifier, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QWinEventNotifier{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QWinEventNotifier) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QWinEventNotifier_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWinEventNotifier) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_TimerEvent
func miqt_exec_callback_QWinEventNotifier_TimerEvent(self QWinEventNotifier, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QWinEventNotifier{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QWinEventNotifier) callVirtualBase_ChildEvent(event *QChildEvent) {

	QWinEventNotifier_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWinEventNotifier) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_ChildEvent
func miqt_exec_callback_QWinEventNotifier_ChildEvent(self QWinEventNotifier, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QWinEventNotifier{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QWinEventNotifier) callVirtualBase_CustomEvent(event *QEvent) {

	QWinEventNotifier_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWinEventNotifier) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_CustomEvent
func miqt_exec_callback_QWinEventNotifier_CustomEvent(self QWinEventNotifier, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QWinEventNotifier{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QWinEventNotifier) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QWinEventNotifier_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QWinEventNotifier) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_ConnectNotify
func miqt_exec_callback_QWinEventNotifier_ConnectNotify(self QWinEventNotifier, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QWinEventNotifier{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QWinEventNotifier) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QWinEventNotifier_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QWinEventNotifier) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_DisconnectNotify
func miqt_exec_callback_QWinEventNotifier_DisconnectNotify(self QWinEventNotifier, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QWinEventNotifier{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
