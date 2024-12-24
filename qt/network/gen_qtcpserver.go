package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QTcpServer struct {
	h          uintptr
	isSubclass bool
}

// NewQTcpServer constructs a new QTcpServer object.
func NewQTcpServer() *QTcpServer {
	g := newQTcpServer(QTcpServer_new())
	g.isSubclass = true
	return g
}

// NewQTcpServer2 constructs a new QTcpServer object.
func NewQTcpServer2(parent *qt.QObject) *QTcpServer {
	g := newQTcpServer(QTcpServer_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
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

func (this *QTcpServer) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QTcpServer_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QTcpServer) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_MetaObject
func miqt_exec_callback_QTcpServer_MetaObject(self QTcpServer, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTcpServer{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QTcpServer) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTcpServer_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTcpServer) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpServer_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpServer_Metacast
func miqt_exec_callback_QTcpServer_Metacast(self QTcpServer, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTcpServer{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
