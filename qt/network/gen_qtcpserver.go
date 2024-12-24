package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QTcpServer struct {
	h          uintptr
	isSubclass bool
}

// NewQTcpServer constructs a new QTcpServer object.
func NewQTcpServer() *QTcpServer {

	ret := newQTcpServer(QTcpServer_new())
	ret.isSubclass = true
	return ret
}

// NewQTcpServer2 constructs a new QTcpServer object.
func NewQTcpServer2(parent *qt.QObject) *QTcpServer {

	ret := newQTcpServer(QTcpServer_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QTcpServer) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QTcpServer_MetaObject(this.h)))
}

func (this *QTcpServer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTcpServer_Metacast(this.h, param1_Cstring))
}

func QTcpServer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTcpServer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTcpServer) Listen() bool {
	return (bool)(QTcpServer_Listen(this.h))
}

func (this *QTcpServer) Close() {
	QTcpServer_Close(this.h)
}

func (this *QTcpServer) IsListening() bool {
	return (bool)(QTcpServer_IsListening(this.h))
}

func (this *QTcpServer) SetMaxPendingConnections(numConnections int) {
	QTcpServer_SetMaxPendingConnections(this.h, (int)(numConnections))
}

func (this *QTcpServer) MaxPendingConnections() int {
	return (int)(QTcpServer_MaxPendingConnections(this.h))
}

func (this *QTcpServer) SetListenBacklogSize(size int) {
	QTcpServer_SetListenBacklogSize(this.h, (int)(size))
}

func (this *QTcpServer) ListenBacklogSize() int {
	return (int)(QTcpServer_ListenBacklogSize(this.h))
}

func (this *QTcpServer) ServerPort() uint16 {
	return (uint16)(QTcpServer_ServerPort(this.h))
}

func (this *QTcpServer) ServerAddress() *QHostAddress {
	_goptr := newQHostAddress(QTcpServer_ServerAddress(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTcpServer) SocketDescriptor() uintptr {
	return (uintptr)(QTcpServer_SocketDescriptor(this.h))
}

func (this *QTcpServer) SetSocketDescriptor(socketDescriptor uintptr) bool {
	return (bool)(QTcpServer_SetSocketDescriptor(this.h, (intptr_t)(socketDescriptor)))
}

func (this *QTcpServer) WaitForNewConnection() bool {
	return (bool)(QTcpServer_WaitForNewConnection(this.h))
}

func (this *QTcpServer) HasPendingConnections() bool {
	return (bool)(QTcpServer_HasPendingConnections(this.h))
}

func (this *QTcpServer) NextPendingConnection() *QTcpSocket {
	return newQTcpSocket(QTcpServer_NextPendingConnection(this.h))
}

func (this *QTcpServer) ServerError() QAbstractSocket__SocketError {
	return (QAbstractSocket__SocketError)(QTcpServer_ServerError(this.h))
}

func (this *QTcpServer) ErrorString() string {
	var _ms struct_miqt_string = QTcpServer_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTcpServer) PauseAccepting() {
	QTcpServer_PauseAccepting(this.h)
}

func (this *QTcpServer) ResumeAccepting() {
	QTcpServer_ResumeAccepting(this.h)
}

func (this *QTcpServer) SetProxy(networkProxy *QNetworkProxy) {
	QTcpServer_SetProxy(this.h, networkProxy.cPointer())
}

func (this *QTcpServer) Proxy() *QNetworkProxy {
	_goptr := newQNetworkProxy(QTcpServer_Proxy(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTcpServer) NewConnection() {
	QTcpServer_NewConnection(this.h)
}
func (this *QTcpServer) OnNewConnection(slot func()) {
	QTcpServer_connect_NewConnection(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_NewConnection
func miqt_exec_callback_QTcpServer_NewConnection(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTcpServer) AcceptError(socketError QAbstractSocket__SocketError) {
	QTcpServer_AcceptError(this.h, (int)(socketError))
}
func (this *QTcpServer) OnAcceptError(slot func(socketError QAbstractSocket__SocketError)) {
	QTcpServer_connect_AcceptError(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_AcceptError
func miqt_exec_callback_QTcpServer_AcceptError(cb intptr_t, socketError int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socketError QAbstractSocket__SocketError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketError)(socketError)

	gofunc(slotval1)
}

func QTcpServer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTcpServer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTcpServer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTcpServer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTcpServer) Listen1(address *QHostAddress) bool {
	return (bool)(QTcpServer_Listen1(this.h, address.cPointer()))
}

func (this *QTcpServer) Listen2(address *QHostAddress, port uint16) bool {
	return (bool)(QTcpServer_Listen2(this.h, address.cPointer(), (uint16_t)(port)))
}

func (this *QTcpServer) WaitForNewConnection1(msec int) bool {
	return (bool)(QTcpServer_WaitForNewConnection1(this.h, (int)(msec)))
}

func (this *QTcpServer) WaitForNewConnection2(msec int, timedOut *bool) bool {
	return (bool)(QTcpServer_WaitForNewConnection2(this.h, (int)(msec), (*bool)(unsafe.Pointer(timedOut))))
}

func (this *QTcpServer) callVirtualBase_HasPendingConnections() bool {

	return (bool)(QTcpServer_virtualbase_HasPendingConnections(unsafe.Pointer(this.h)))

}
func (this *QTcpServer) OnHasPendingConnections(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_HasPendingConnections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_HasPendingConnections
func miqt_exec_callback_QTcpServer_HasPendingConnections(self QTcpServer, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTcpServer{h: self}).callVirtualBase_HasPendingConnections)

	return (bool)(virtualReturn)

}

func (this *QTcpServer) callVirtualBase_NextPendingConnection() *QTcpSocket {

	return newQTcpSocket(QTcpServer_virtualbase_NextPendingConnection(unsafe.Pointer(this.h)))

}
func (this *QTcpServer) OnNextPendingConnection(slot func(super func() *QTcpSocket) *QTcpSocket) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_NextPendingConnection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_NextPendingConnection
func miqt_exec_callback_QTcpServer_NextPendingConnection(self QTcpServer, cb intptr_t) *QTcpSocket {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QTcpSocket) *QTcpSocket)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTcpServer{h: self}).callVirtualBase_NextPendingConnection)

	return virtualReturn.cPointer()

}

func (this *QTcpServer) callVirtualBase_IncomingConnection(handle uintptr) {

	QTcpServer_virtualbase_IncomingConnection(unsafe.Pointer(this.h), (intptr_t)(handle))

}
func (this *QTcpServer) OnIncomingConnection(slot func(super func(handle uintptr), handle uintptr)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_IncomingConnection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_IncomingConnection
func miqt_exec_callback_QTcpServer_IncomingConnection(self QTcpServer, cb intptr_t, handle intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(handle uintptr), handle uintptr))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uintptr)(handle)

	gofunc((&QTcpServer{h: self}).callVirtualBase_IncomingConnection, slotval1)

}

func (this *QTcpServer) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QTcpServer_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QTcpServer) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_Event
func miqt_exec_callback_QTcpServer_Event(self QTcpServer, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QTcpServer{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTcpServer) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QTcpServer_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QTcpServer) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_EventFilter
func miqt_exec_callback_QTcpServer_EventFilter(self QTcpServer, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QTcpServer{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QTcpServer) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QTcpServer_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QTcpServer) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_TimerEvent
func miqt_exec_callback_QTcpServer_TimerEvent(self QTcpServer, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QTcpServer{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTcpServer) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QTcpServer_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QTcpServer) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_ChildEvent
func miqt_exec_callback_QTcpServer_ChildEvent(self QTcpServer, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QTcpServer{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QTcpServer) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QTcpServer_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QTcpServer) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_CustomEvent
func miqt_exec_callback_QTcpServer_CustomEvent(self QTcpServer, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QTcpServer{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QTcpServer) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QTcpServer_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QTcpServer) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_ConnectNotify
func miqt_exec_callback_QTcpServer_ConnectNotify(self QTcpServer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QTcpServer{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QTcpServer) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QTcpServer_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QTcpServer) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_DisconnectNotify
func miqt_exec_callback_QTcpServer_DisconnectNotify(self QTcpServer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QTcpServer{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
