package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
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

func (this *QLocalServer) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QLocalServer_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QLocalServer) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_MetaObject
func miqt_exec_callback_QLocalServer_MetaObject(self QLocalServer, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalServer{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QLocalServer) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QLocalServer_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QLocalServer) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalServer_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalServer_Metacast
func miqt_exec_callback_QLocalServer_Metacast(self QLocalServer, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QLocalServer{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
