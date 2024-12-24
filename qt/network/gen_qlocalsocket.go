package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
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

	ret := newQLocalSocket(QLocalSocket_new())
	ret.isSubclass = true
	return ret
}

// NewQLocalSocket2 constructs a new QLocalSocket object.
func NewQLocalSocket2(parent *qt.QObject) *QLocalSocket {

	ret := newQLocalSocket(QLocalSocket_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
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

func (this *QLocalSocket) callVirtualBase_IsSequential() bool {

	return (bool)(QLocalSocket_virtualbase_IsSequential(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnIsSequential(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_IsSequential
func miqt_exec_callback_QLocalSocket_IsSequential(self QLocalSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_IsSequential)

	return (bool)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_BytesAvailable() int64 {

	return (int64)(QLocalSocket_virtualbase_BytesAvailable(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnBytesAvailable(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_BytesAvailable(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_BytesAvailable
func miqt_exec_callback_QLocalSocket_BytesAvailable(self QLocalSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_BytesAvailable)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_BytesToWrite() int64 {

	return (int64)(QLocalSocket_virtualbase_BytesToWrite(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnBytesToWrite(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_BytesToWrite(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_BytesToWrite
func miqt_exec_callback_QLocalSocket_BytesToWrite(self QLocalSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_BytesToWrite)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_CanReadLine() bool {

	return (bool)(QLocalSocket_virtualbase_CanReadLine(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnCanReadLine(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_CanReadLine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_CanReadLine
func miqt_exec_callback_QLocalSocket_CanReadLine(self QLocalSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_CanReadLine)

	return (bool)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_Open(openMode OpenMode) bool {

	return (bool)(QLocalSocket_virtualbase_Open(unsafe.Pointer(this.h), openMode))

}
func (this *QLocalSocket) OnOpen(slot func(super func(openMode OpenMode) bool, openMode OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Open
func miqt_exec_callback_QLocalSocket_Open(self QLocalSocket, cb intptr_t, openMode OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(openMode OpenMode) bool, openMode OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_Open, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_Close() {

	QLocalSocket_virtualbase_Close(unsafe.Pointer(this.h))

}
func (this *QLocalSocket) OnClose(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_Close(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Close
func miqt_exec_callback_QLocalSocket_Close(self QLocalSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QLocalSocket{h: self}).callVirtualBase_Close)

}

func (this *QLocalSocket) callVirtualBase_WaitForBytesWritten(msecs int) bool {

	return (bool)(QLocalSocket_virtualbase_WaitForBytesWritten(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QLocalSocket) OnWaitForBytesWritten(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_WaitForBytesWritten(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_WaitForBytesWritten
func miqt_exec_callback_QLocalSocket_WaitForBytesWritten(self QLocalSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_WaitForBytesWritten, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_WaitForReadyRead(msecs int) bool {

	return (bool)(QLocalSocket_virtualbase_WaitForReadyRead(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QLocalSocket) OnWaitForReadyRead(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_WaitForReadyRead(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_WaitForReadyRead
func miqt_exec_callback_QLocalSocket_WaitForReadyRead(self QLocalSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_WaitForReadyRead, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_ReadData(param1 string, param2 int64) int64 {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (int64)(QLocalSocket_virtualbase_ReadData(unsafe.Pointer(this.h), param1_Cstring, (longlong)(param2)))

}
func (this *QLocalSocket) OnReadData(slot func(super func(param1 string, param2 int64) int64, param1 string, param2 int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_ReadData
func miqt_exec_callback_QLocalSocket_ReadData(self QLocalSocket, cb intptr_t, param1 *char, param2 longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string, param2 int64) int64, param1 string, param2 int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	slotval2 := (int64)(param2)

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_ReadData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_ReadLineData(data string, maxSize int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QLocalSocket_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxSize)))

}
func (this *QLocalSocket) OnReadLineData(slot func(super func(data string, maxSize int64) int64, data string, maxSize int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_ReadLineData
func miqt_exec_callback_QLocalSocket_ReadLineData(self QLocalSocket, cb intptr_t, data *char, maxSize longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxSize int64) int64, data string, maxSize int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxSize)

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_SkipData(maxSize int64) int64 {

	return (int64)(QLocalSocket_virtualbase_SkipData(unsafe.Pointer(this.h), (longlong)(maxSize)))

}
func (this *QLocalSocket) OnSkipData(slot func(super func(maxSize int64) int64, maxSize int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_SkipData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_SkipData
func miqt_exec_callback_QLocalSocket_SkipData(self QLocalSocket, cb intptr_t, maxSize longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(maxSize int64) int64, maxSize int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(maxSize)

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_SkipData, slotval1)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_WriteData(param1 string, param2 int64) int64 {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (int64)(QLocalSocket_virtualbase_WriteData(unsafe.Pointer(this.h), param1_Cstring, (longlong)(param2)))

}
func (this *QLocalSocket) OnWriteData(slot func(super func(param1 string, param2 int64) int64, param1 string, param2 int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_WriteData
func miqt_exec_callback_QLocalSocket_WriteData(self QLocalSocket, cb intptr_t, param1 *const_char, param2 longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string, param2 int64) int64, param1 string, param2 int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	slotval2 := (int64)(param2)

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_WriteData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_Pos() int64 {

	return (int64)(QLocalSocket_virtualbase_Pos(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnPos(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_Pos(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Pos
func miqt_exec_callback_QLocalSocket_Pos(self QLocalSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_Pos)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_Size() int64 {

	return (int64)(QLocalSocket_virtualbase_Size(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnSize(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Size
func miqt_exec_callback_QLocalSocket_Size(self QLocalSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_Size)

	return (longlong)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_Seek(pos int64) bool {

	return (bool)(QLocalSocket_virtualbase_Seek(unsafe.Pointer(this.h), (longlong)(pos)))

}
func (this *QLocalSocket) OnSeek(slot func(super func(pos int64) bool, pos int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_Seek(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Seek
func miqt_exec_callback_QLocalSocket_Seek(self QLocalSocket, cb intptr_t, pos longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos int64) bool, pos int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(pos)

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_Seek, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_AtEnd() bool {

	return (bool)(QLocalSocket_virtualbase_AtEnd(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnAtEnd(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_AtEnd(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_AtEnd
func miqt_exec_callback_QLocalSocket_AtEnd(self QLocalSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_AtEnd)

	return (bool)(virtualReturn)

}

func (this *QLocalSocket) callVirtualBase_Reset() bool {

	return (bool)(QLocalSocket_virtualbase_Reset(unsafe.Pointer(this.h)))

}
func (this *QLocalSocket) OnReset(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLocalSocket_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocalSocket_Reset
func miqt_exec_callback_QLocalSocket_Reset(self QLocalSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLocalSocket{h: self}).callVirtualBase_Reset)

	return (bool)(virtualReturn)

}
