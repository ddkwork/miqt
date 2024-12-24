package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QLocalServer__SocketOption int

const (
	QLocalServer__NoOptions               QLocalServer__SocketOption = 0
	QLocalServer__UserAccessOption        QLocalServer__SocketOption = 1
	QLocalServer__GroupAccessOption       QLocalServer__SocketOption = 2
	QLocalServer__OtherAccessOption       QLocalServer__SocketOption = 4
	QLocalServer__WorldAccessOption       QLocalServer__SocketOption = 7
	QLocalServer__AbstractNamespaceOption QLocalServer__SocketOption = 8
)

type QLocalServer struct {
	h          uintptr
	isSubclass bool
}

// NewQLocalServer constructs a new QLocalServer object.
func NewQLocalServer() *QLocalServer {

	ret := newQLocalServer(QLocalServer_new())
	ret.isSubclass = true
	return ret
}

// NewQLocalServer2 constructs a new QLocalServer object.
func NewQLocalServer2(parent *qt.QObject) *QLocalServer {

	ret := newQLocalServer(QLocalServer_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QLocalServer) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QLocalServer_MetaObject(this.h)))
}

func (this *QLocalServer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLocalServer_Metacast(this.h, param1_Cstring))
}

func QLocalServer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLocalServer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalServer) NewConnection() {
	QLocalServer_NewConnection(this.h)
}
func (this *QLocalServer) OnNewConnection(slot func()) {
	QLocalServer_connect_NewConnection(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_NewConnection
func miqt_exec_callback_QLocalServer_NewConnection(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QLocalServer) Close() {
	QLocalServer_Close(this.h)
}

func (this *QLocalServer) ErrorString() string {
	var _ms struct_miqt_string = QLocalServer_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalServer) HasPendingConnections() bool {
	return (bool)(QLocalServer_HasPendingConnections(this.h))
}

func (this *QLocalServer) IsListening() bool {
	return (bool)(QLocalServer_IsListening(this.h))
}

func (this *QLocalServer) Listen(name string) bool {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	return (bool)(QLocalServer_Listen(this.h, name_ms))
}

func (this *QLocalServer) ListenWithSocketDescriptor(socketDescriptor uintptr) bool {
	return (bool)(QLocalServer_ListenWithSocketDescriptor(this.h, (intptr_t)(socketDescriptor)))
}

func (this *QLocalServer) MaxPendingConnections() int {
	return (int)(QLocalServer_MaxPendingConnections(this.h))
}

func (this *QLocalServer) NextPendingConnection() *QLocalSocket {
	return newQLocalSocket(QLocalServer_NextPendingConnection(this.h))
}

func (this *QLocalServer) ServerName() string {
	var _ms struct_miqt_string = QLocalServer_ServerName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalServer) FullServerName() string {
	var _ms struct_miqt_string = QLocalServer_FullServerName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLocalServer_RemoveServer(name string) bool {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	return (bool)(QLocalServer_RemoveServer(name_ms))
}

func (this *QLocalServer) ServerError() QAbstractSocket__SocketError {
	return (QAbstractSocket__SocketError)(QLocalServer_ServerError(this.h))
}

func (this *QLocalServer) SetMaxPendingConnections(numConnections int) {
	QLocalServer_SetMaxPendingConnections(this.h, (int)(numConnections))
}

func (this *QLocalServer) WaitForNewConnection() bool {
	return (bool)(QLocalServer_WaitForNewConnection(this.h))
}

func (this *QLocalServer) SetListenBacklogSize(size int) {
	QLocalServer_SetListenBacklogSize(this.h, (int)(size))
}

func (this *QLocalServer) ListenBacklogSize() int {
	return (int)(QLocalServer_ListenBacklogSize(this.h))
}

func (this *QLocalServer) SetSocketOptions(options SocketOptions) {
	QLocalServer_SetSocketOptions(this.h, options)
}

func (this *QLocalServer) SocketOptions() SocketOptions {
	xxxxxxxxx
}

func (this *QLocalServer) SocketDescriptor() uintptr {
	return (uintptr)(QLocalServer_SocketDescriptor(this.h))
}

func QLocalServer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLocalServer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLocalServer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLocalServer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalServer) WaitForNewConnection1(msec int) bool {
	return (bool)(QLocalServer_WaitForNewConnection1(this.h, (int)(msec)))
}

func (this *QLocalServer) WaitForNewConnection2(msec int, timedOut *bool) bool {
	return (bool)(QLocalServer_WaitForNewConnection2(this.h, (int)(msec), (*bool)(unsafe.Pointer(timedOut))))
}

func (this *QLocalServer) callVirtualBase_HasPendingConnections() bool {

	return (bool)(QLocalServer_virtualbase_HasPendingConnections(unsafe.Pointer(this.h)))

}
func (this *QLocalServer) OnHasPendingConnections(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_HasPendingConnections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_HasPendingConnections
func miqt_exec_callback_QLocalServer_HasPendingConnections(self QLocalServer, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalServer{h: self}).callVirtualBase_HasPendingConnections)

	return (bool)(virtualReturn)

}

func (this *QLocalServer) callVirtualBase_NextPendingConnection() *QLocalSocket {

	return newQLocalSocket(QLocalServer_virtualbase_NextPendingConnection(unsafe.Pointer(this.h)))

}
func (this *QLocalServer) OnNextPendingConnection(slot func(super func() *QLocalSocket) *QLocalSocket) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_NextPendingConnection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_NextPendingConnection
func miqt_exec_callback_QLocalServer_NextPendingConnection(self QLocalServer, cb intptr_t) *QLocalSocket {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLocalSocket) *QLocalSocket)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalServer{h: self}).callVirtualBase_NextPendingConnection)

	return virtualReturn.cPointer()

}

func (this *QLocalServer) callVirtualBase_IncomingConnection(socketDescriptor uintptr) {

	QLocalServer_virtualbase_IncomingConnection(unsafe.Pointer(this.h), (uintptr_t)(socketDescriptor))

}
func (this *QLocalServer) OnIncomingConnection(slot func(super func(socketDescriptor uintptr), socketDescriptor uintptr)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_IncomingConnection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_IncomingConnection
func miqt_exec_callback_QLocalServer_IncomingConnection(self QLocalServer, cb intptr_t, socketDescriptor uintptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(socketDescriptor uintptr), socketDescriptor uintptr))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uintptr)(socketDescriptor)

	gofunc((&QLocalServer{h: self}).callVirtualBase_IncomingConnection, slotval1)

}

func (this *QLocalServer) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QLocalServer_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QLocalServer) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_Event
func miqt_exec_callback_QLocalServer_Event(self QLocalServer, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QLocalServer{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLocalServer) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QLocalServer_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QLocalServer) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_EventFilter
func miqt_exec_callback_QLocalServer_EventFilter(self QLocalServer, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QLocalServer{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QLocalServer) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QLocalServer_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QLocalServer) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_TimerEvent
func miqt_exec_callback_QLocalServer_TimerEvent(self QLocalServer, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QLocalServer{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QLocalServer) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QLocalServer_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QLocalServer) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_ChildEvent
func miqt_exec_callback_QLocalServer_ChildEvent(self QLocalServer, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QLocalServer{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QLocalServer) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QLocalServer_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QLocalServer) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_CustomEvent
func miqt_exec_callback_QLocalServer_CustomEvent(self QLocalServer, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QLocalServer{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QLocalServer) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QLocalServer_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QLocalServer) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_ConnectNotify
func miqt_exec_callback_QLocalServer_ConnectNotify(self QLocalServer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QLocalServer{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QLocalServer) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QLocalServer_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QLocalServer) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_DisconnectNotify
func miqt_exec_callback_QLocalServer_DisconnectNotify(self QLocalServer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QLocalServer{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
