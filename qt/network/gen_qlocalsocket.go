package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QLocalSocket__LocalSocketError int

const (
	QLocalSocket__ConnectionRefusedError          QLocalSocket__LocalSocketError = 0
	QLocalSocket__PeerClosedError                 QLocalSocket__LocalSocketError = 1
	QLocalSocket__ServerNotFoundError             QLocalSocket__LocalSocketError = 2
	QLocalSocket__SocketAccessError               QLocalSocket__LocalSocketError = 3
	QLocalSocket__SocketResourceError             QLocalSocket__LocalSocketError = 4
	QLocalSocket__SocketTimeoutError              QLocalSocket__LocalSocketError = 5
	QLocalSocket__DatagramTooLargeError           QLocalSocket__LocalSocketError = 6
	QLocalSocket__ConnectionError                 QLocalSocket__LocalSocketError = 7
	QLocalSocket__UnsupportedSocketOperationError QLocalSocket__LocalSocketError = 10
	QLocalSocket__UnknownSocketError              QLocalSocket__LocalSocketError = -1
	QLocalSocket__OperationError                  QLocalSocket__LocalSocketError = 19
)

type QLocalSocket__LocalSocketState int

const (
	QLocalSocket__UnconnectedState QLocalSocket__LocalSocketState = 0
	QLocalSocket__ConnectingState  QLocalSocket__LocalSocketState = 2
	QLocalSocket__ConnectedState   QLocalSocket__LocalSocketState = 3
	QLocalSocket__ClosingState     QLocalSocket__LocalSocketState = 6
)

type QLocalSocket__SocketOption int

const (
	QLocalSocket__NoOptions               QLocalSocket__SocketOption = 0
	QLocalSocket__AbstractNamespaceOption QLocalSocket__SocketOption = 1
)

type QLocalSocket struct {
	h          uintptr
	isSubclass bool
}

// NewQLocalSocket constructs a new QLocalSocket object.
func NewQLocalSocket() *QLocalSocket {
	g := newQLocalSocket(QLocalSocket_new())
	g.isSubclass = true
	return g
}

// NewQLocalSocket2 constructs a new QLocalSocket object.
func NewQLocalSocket2(parent *qt.QObject) *QLocalSocket {
	g := newQLocalSocket(QLocalSocket_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QLocalSocket) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QLocalSocket_MetaObject(this.h)))
}

func (this *QLocalSocket) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLocalSocket_Metacast(this.h, param1_Cstring))
}

func QLocalSocket_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLocalSocket_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalSocket) ConnectToServer() {
	QLocalSocket_ConnectToServer(this.h)
}

func (this *QLocalSocket) ConnectToServerWithName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QLocalSocket_ConnectToServerWithName(this.h, name_ms)
}

func (this *QLocalSocket) DisconnectFromServer() {
	QLocalSocket_DisconnectFromServer(this.h)
}

func (this *QLocalSocket) SetServerName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QLocalSocket_SetServerName(this.h, name_ms)
}

func (this *QLocalSocket) ServerName() string {
	var _ms struct_miqt_string = QLocalSocket_ServerName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalSocket) FullServerName() string {
	var _ms struct_miqt_string = QLocalSocket_FullServerName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalSocket) Abort() {
	QLocalSocket_Abort(this.h)
}

func (this *QLocalSocket) IsSequential() bool {
	return (bool)(QLocalSocket_IsSequential(this.h))
}

func (this *QLocalSocket) BytesAvailable() int64 {
	return (int64)(QLocalSocket_BytesAvailable(this.h))
}

func (this *QLocalSocket) BytesToWrite() int64 {
	return (int64)(QLocalSocket_BytesToWrite(this.h))
}

func (this *QLocalSocket) CanReadLine() bool {
	return (bool)(QLocalSocket_CanReadLine(this.h))
}

func (this *QLocalSocket) Open(openMode OpenMode) bool {
	return (bool)(QLocalSocket_Open(this.h, openMode))
}

func (this *QLocalSocket) Close() {
	QLocalSocket_Close(this.h)
}

func (this *QLocalSocket) Error() LocalSocketError {
	xxxxxxxxx
}

func (this *QLocalSocket) Flush() bool {
	return (bool)(QLocalSocket_Flush(this.h))
}

func (this *QLocalSocket) IsValid() bool {
	return (bool)(QLocalSocket_IsValid(this.h))
}

func (this *QLocalSocket) ReadBufferSize() int64 {
	return (int64)(QLocalSocket_ReadBufferSize(this.h))
}

func (this *QLocalSocket) SetReadBufferSize(size int64) {
	QLocalSocket_SetReadBufferSize(this.h, (longlong)(size))
}

func (this *QLocalSocket) SetSocketDescriptor(socketDescriptor uintptr) bool {
	return (bool)(QLocalSocket_SetSocketDescriptor(this.h, (intptr_t)(socketDescriptor)))
}

func (this *QLocalSocket) SocketDescriptor() uintptr {
	return (uintptr)(QLocalSocket_SocketDescriptor(this.h))
}

func (this *QLocalSocket) SetSocketOptions(option SocketOptions) {
	QLocalSocket_SetSocketOptions(this.h, option)
}

func (this *QLocalSocket) SocketOptions() SocketOptions {
	xxxxxxxxx
}

func (this *QLocalSocket) State() LocalSocketState {
	xxxxxxxxx
}

func (this *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	return (bool)(QLocalSocket_WaitForBytesWritten(this.h, (int)(msecs)))
}

func (this *QLocalSocket) WaitForConnected() bool {
	return (bool)(QLocalSocket_WaitForConnected(this.h))
}

func (this *QLocalSocket) WaitForDisconnected() bool {
	return (bool)(QLocalSocket_WaitForDisconnected(this.h))
}

func (this *QLocalSocket) WaitForReadyRead(msecs int) bool {
	return (bool)(QLocalSocket_WaitForReadyRead(this.h, (int)(msecs)))
}

func (this *QLocalSocket) Connected() {
	QLocalSocket_Connected(this.h)
}

func (this *QLocalSocket) OnConnected(slot func()) {
	QLocalSocket_connect_Connected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Connected
func miqt_exec_callback_QLocalSocket_Connected(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QLocalSocket) Disconnected() {
	QLocalSocket_Disconnected(this.h)
}

func (this *QLocalSocket) OnDisconnected(slot func()) {
	QLocalSocket_connect_Disconnected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Disconnected
func miqt_exec_callback_QLocalSocket_Disconnected(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QLocalSocket) ErrorOccurred(socketError QLocalSocket__LocalSocketError) {
	QLocalSocket_ErrorOccurred(this.h, (int)(socketError))
}

func (this *QLocalSocket) OnErrorOccurred(slot func(socketError QLocalSocket__LocalSocketError)) {
	QLocalSocket_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_ErrorOccurred
func miqt_exec_callback_QLocalSocket_ErrorOccurred(cb intptr_t, socketError int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socketError QLocalSocket__LocalSocketError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QLocalSocket__LocalSocketError)(socketError)

	gofunc(slotval1)
}

func (this *QLocalSocket) StateChanged(socketState QLocalSocket__LocalSocketState) {
	QLocalSocket_StateChanged(this.h, (int)(socketState))
}

func (this *QLocalSocket) OnStateChanged(slot func(socketState QLocalSocket__LocalSocketState)) {
	QLocalSocket_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_StateChanged
func miqt_exec_callback_QLocalSocket_StateChanged(cb intptr_t, socketState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socketState QLocalSocket__LocalSocketState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QLocalSocket__LocalSocketState)(socketState)

	gofunc(slotval1)
}

func QLocalSocket_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLocalSocket_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLocalSocket_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLocalSocket_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLocalSocket) ConnectToServer1(openMode OpenMode) {
	QLocalSocket_ConnectToServer1(this.h, openMode)
}

func (this *QLocalSocket) ConnectToServer2(name string, openMode OpenMode) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QLocalSocket_ConnectToServer2(this.h, name_ms, openMode)
}

func (this *QLocalSocket) SetSocketDescriptor2(socketDescriptor uintptr, socketState LocalSocketState) bool {
	return (bool)(QLocalSocket_SetSocketDescriptor2(this.h, (intptr_t)(socketDescriptor), socketState))
}

func (this *QLocalSocket) SetSocketDescriptor3(socketDescriptor uintptr, socketState LocalSocketState, openMode OpenMode) bool {
	return (bool)(QLocalSocket_SetSocketDescriptor3(this.h, (intptr_t)(socketDescriptor), socketState, openMode))
}

func (this *QLocalSocket) WaitForConnected1(msecs int) bool {
	return (bool)(QLocalSocket_WaitForConnected1(this.h, (int)(msecs)))
}

func (this *QLocalSocket) WaitForDisconnected1(msecs int) bool {
	return (bool)(QLocalSocket_WaitForDisconnected1(this.h, (int)(msecs)))
}

func (this *QLocalSocket) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QLocalSocket_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QLocalSocket) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_MetaObject
func miqt_exec_callback_QLocalSocket_MetaObject(self QLocalSocket, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QLocalSocket) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QLocalSocket_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QLocalSocket) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Metacast
func miqt_exec_callback_QLocalSocket_Metacast(self QLocalSocket, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
