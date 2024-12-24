package qt

import (
	"unsafe"
)

type QSocketNotifier__Type int

const (
	QSocketNotifier__Read      QSocketNotifier__Type = 0
	QSocketNotifier__Write     QSocketNotifier__Type = 1
	QSocketNotifier__Exception QSocketNotifier__Type = 2
)

type QSocketNotifier struct {
	h          uintptr
	isSubclass bool
}

// NewQSocketNotifier constructs a new QSocketNotifier object.
func NewQSocketNotifier(param1 Type) *QSocketNotifier {

	ret := newQSocketNotifier(QSocketNotifier_new(param1))
	ret.isSubclass = true
	return ret
}

// NewQSocketNotifier2 constructs a new QSocketNotifier object.
func NewQSocketNotifier2(socket uintptr, param2 Type) *QSocketNotifier {

	ret := newQSocketNotifier(QSocketNotifier_new2((intptr_t)(socket), param2))
	ret.isSubclass = true
	return ret
}

// NewQSocketNotifier3 constructs a new QSocketNotifier object.
func NewQSocketNotifier3(param1 Type, parent *QObject) *QSocketNotifier {

	ret := newQSocketNotifier(QSocketNotifier_new3(param1, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSocketNotifier4 constructs a new QSocketNotifier object.
func NewQSocketNotifier4(socket uintptr, param2 Type, parent *QObject) *QSocketNotifier {

	ret := newQSocketNotifier(QSocketNotifier_new4((intptr_t)(socket), param2, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSocketNotifier) MetaObject() *QMetaObject {
	return newQMetaObject(QSocketNotifier_MetaObject(this.h))
}

func (this *QSocketNotifier) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSocketNotifier_Metacast(this.h, param1_Cstring))
}

func QSocketNotifier_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSocketNotifier_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSocketNotifier) SetSocket(socket uintptr) {
	QSocketNotifier_SetSocket(this.h, (intptr_t)(socket))
}

func (this *QSocketNotifier) Socket() uintptr {
	return (uintptr)(QSocketNotifier_Socket(this.h))
}

func (this *QSocketNotifier) Type() Type {
	xxxxxxxxx
}

func (this *QSocketNotifier) IsValid() bool {
	return (bool)(QSocketNotifier_IsValid(this.h))
}

func (this *QSocketNotifier) IsEnabled() bool {
	return (bool)(QSocketNotifier_IsEnabled(this.h))
}

func (this *QSocketNotifier) SetEnabled(enabled bool) {
	QSocketNotifier_SetEnabled(this.h, (bool)(enabled))
}

func QSocketNotifier_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSocketNotifier_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSocketNotifier_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSocketNotifier_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSocketNotifier) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QSocketNotifier_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QSocketNotifier) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_Event
func miqt_exec_callback_QSocketNotifier_Event(self QSocketNotifier, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QSocketNotifier{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSocketNotifier) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QSocketNotifier_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QSocketNotifier) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_EventFilter
func miqt_exec_callback_QSocketNotifier_EventFilter(self QSocketNotifier, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QSocketNotifier{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSocketNotifier) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QSocketNotifier_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSocketNotifier) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_TimerEvent
func miqt_exec_callback_QSocketNotifier_TimerEvent(self QSocketNotifier, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QSocketNotifier{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSocketNotifier) callVirtualBase_ChildEvent(event *QChildEvent) {

	QSocketNotifier_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSocketNotifier) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_ChildEvent
func miqt_exec_callback_QSocketNotifier_ChildEvent(self QSocketNotifier, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QSocketNotifier{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSocketNotifier) callVirtualBase_CustomEvent(event *QEvent) {

	QSocketNotifier_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSocketNotifier) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_CustomEvent
func miqt_exec_callback_QSocketNotifier_CustomEvent(self QSocketNotifier, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSocketNotifier{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QSocketNotifier) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QSocketNotifier_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSocketNotifier) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_ConnectNotify
func miqt_exec_callback_QSocketNotifier_ConnectNotify(self QSocketNotifier, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSocketNotifier{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QSocketNotifier) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QSocketNotifier_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSocketNotifier) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_DisconnectNotify
func miqt_exec_callback_QSocketNotifier_DisconnectNotify(self QSocketNotifier, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSocketNotifier{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

type QSocketDescriptor struct {
	h          uintptr
	isSubclass bool
}

// NewQSocketDescriptor constructs a new QSocketDescriptor object.
func NewQSocketDescriptor() *QSocketDescriptor {

	ret := newQSocketDescriptor(QSocketDescriptor_new())
	ret.isSubclass = true
	return ret
}

// NewQSocketDescriptor2 constructs a new QSocketDescriptor object.
func NewQSocketDescriptor2(desc uintptr) *QSocketDescriptor {

	ret := newQSocketDescriptor(QSocketDescriptor_new2((intptr_t)(desc)))
	ret.isSubclass = true
	return ret
}

// NewQSocketDescriptor3 constructs a new QSocketDescriptor object.
func NewQSocketDescriptor3(param1 *QSocketDescriptor) *QSocketDescriptor {

	ret := newQSocketDescriptor(QSocketDescriptor_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSocketDescriptor4 constructs a new QSocketDescriptor object.
func NewQSocketDescriptor4(descriptor DescriptorType) *QSocketDescriptor {

	ret := newQSocketDescriptor(QSocketDescriptor_new4(descriptor))
	ret.isSubclass = true
	return ret
}

func (this *QSocketDescriptor) WinHandle() unsafe.Pointer {
	return (unsafe.Pointer)(QSocketDescriptor_WinHandle(this.h))
}

func (this *QSocketDescriptor) IsValid() bool {
	return (bool)(QSocketDescriptor_IsValid(this.h))
}
