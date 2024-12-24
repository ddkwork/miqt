package webchannel

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QWebChannelAbstractTransport struct {
	h          uintptr
	isSubclass bool
}

// NewQWebChannelAbstractTransport constructs a new QWebChannelAbstractTransport object.
func NewQWebChannelAbstractTransport() *QWebChannelAbstractTransport {

	ret := newQWebChannelAbstractTransport(QWebChannelAbstractTransport_new())
	ret.isSubclass = true
	return ret
}

// NewQWebChannelAbstractTransport2 constructs a new QWebChannelAbstractTransport object.
func NewQWebChannelAbstractTransport2(parent *qt.QObject) *QWebChannelAbstractTransport {

	ret := newQWebChannelAbstractTransport(QWebChannelAbstractTransport_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QWebChannelAbstractTransport) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWebChannelAbstractTransport_MetaObject(this.h)))
}

func (this *QWebChannelAbstractTransport) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWebChannelAbstractTransport_Metacast(this.h, param1_Cstring))
}

func QWebChannelAbstractTransport_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWebChannelAbstractTransport_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebChannelAbstractTransport) SendMessage(message *qt.QJsonObject) {
	QWebChannelAbstractTransport_SendMessage(this.h, (*QJsonObject)(message.UnsafePointer()))
}

func (this *QWebChannelAbstractTransport) MessageReceived(message *qt.QJsonObject, transport *QWebChannelAbstractTransport) {
	QWebChannelAbstractTransport_MessageReceived(this.h, (*QJsonObject)(message.UnsafePointer()), transport.cPointer())
}
func (this *QWebChannelAbstractTransport) OnMessageReceived(slot func(message *qt.QJsonObject, transport *QWebChannelAbstractTransport)) {
	QWebChannelAbstractTransport_connect_MessageReceived(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_MessageReceived
func miqt_exec_callback_QWebChannelAbstractTransport_MessageReceived(cb intptr_t, message *QJsonObject, transport *QWebChannelAbstractTransport) {
	gofunc, ok := cgo.Handle(cb).Value().(func(message *qt.QJsonObject, transport *QWebChannelAbstractTransport))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQJsonObject(unsafe.Pointer(message))

	slotval2 := newQWebChannelAbstractTransport(transport)

	gofunc(slotval1, slotval2)
}

func QWebChannelAbstractTransport_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWebChannelAbstractTransport_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebChannelAbstractTransport_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWebChannelAbstractTransport_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QWebChannelAbstractTransport) OnSendMessage(slot func(message *qt.QJsonObject)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_SendMessage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_SendMessage
func miqt_exec_callback_QWebChannelAbstractTransport_SendMessage(self QWebChannelAbstractTransport, cb intptr_t, message *QJsonObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(message *qt.QJsonObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQJsonObject(unsafe.Pointer(message))

	gofunc(slotval1)

}

func (this *QWebChannelAbstractTransport) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QWebChannelAbstractTransport_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QWebChannelAbstractTransport) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_Event
func miqt_exec_callback_QWebChannelAbstractTransport_Event(self QWebChannelAbstractTransport, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QWebChannelAbstractTransport) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QWebChannelAbstractTransport_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QWebChannelAbstractTransport) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_EventFilter
func miqt_exec_callback_QWebChannelAbstractTransport_EventFilter(self QWebChannelAbstractTransport, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QWebChannelAbstractTransport) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QWebChannelAbstractTransport_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QWebChannelAbstractTransport) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_TimerEvent
func miqt_exec_callback_QWebChannelAbstractTransport_TimerEvent(self QWebChannelAbstractTransport, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QWebChannelAbstractTransport) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QWebChannelAbstractTransport_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QWebChannelAbstractTransport) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_ChildEvent
func miqt_exec_callback_QWebChannelAbstractTransport_ChildEvent(self QWebChannelAbstractTransport, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QWebChannelAbstractTransport) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QWebChannelAbstractTransport_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QWebChannelAbstractTransport) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_CustomEvent
func miqt_exec_callback_QWebChannelAbstractTransport_CustomEvent(self QWebChannelAbstractTransport, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QWebChannelAbstractTransport) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QWebChannelAbstractTransport_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebChannelAbstractTransport) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_ConnectNotify
func miqt_exec_callback_QWebChannelAbstractTransport_ConnectNotify(self QWebChannelAbstractTransport, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QWebChannelAbstractTransport) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QWebChannelAbstractTransport_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebChannelAbstractTransport) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_DisconnectNotify
func miqt_exec_callback_QWebChannelAbstractTransport_DisconnectNotify(self QWebChannelAbstractTransport, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
