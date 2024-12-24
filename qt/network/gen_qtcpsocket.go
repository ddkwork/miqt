package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QTcpSocket struct {
	h          uintptr
	isSubclass bool
}

// NewQTcpSocket constructs a new QTcpSocket object.
func NewQTcpSocket() *QTcpSocket {

	ret := newQTcpSocket(QTcpSocket_new())
	ret.isSubclass = true
	return ret
}

// NewQTcpSocket2 constructs a new QTcpSocket object.
func NewQTcpSocket2(parent *qt.QObject) *QTcpSocket {

	ret := newQTcpSocket(QTcpSocket_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QTcpSocket) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QTcpSocket_MetaObject(this.h)))
}

func (this *QTcpSocket) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTcpSocket_Metacast(this.h, param1_Cstring))
}

func QTcpSocket_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTcpSocket_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTcpSocket) Bind(addr QHostAddress__SpecialAddress) bool {
	return (bool)(QTcpSocket_Bind(this.h, (int)(addr)))
}

func QTcpSocket_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTcpSocket_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTcpSocket_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTcpSocket_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTcpSocket) Bind2(addr QHostAddress__SpecialAddress, port uint16) bool {
	return (bool)(QTcpSocket_Bind2(this.h, (int)(addr), (uint16_t)(port)))
}

func (this *QTcpSocket) Bind3(addr QHostAddress__SpecialAddress, port uint16, mode BindMode) bool {
	return (bool)(QTcpSocket_Bind3(this.h, (int)(addr), (uint16_t)(port), mode))
}

func (this *QTcpSocket) callVirtualBase_Resume() {

	QTcpSocket_virtualbase_Resume(unsafe.Pointer(this.h))

}
func (this *QTcpSocket) OnResume(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_Resume(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_Resume
func miqt_exec_callback_QTcpSocket_Resume(self QTcpSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTcpSocket{h: self}).callVirtualBase_Resume)

}

func (this *QTcpSocket) callVirtualBase_Bind(address *QHostAddress, port uint16, mode BindMode) bool {

	return (bool)(QTcpSocket_virtualbase_Bind(unsafe.Pointer(this.h), address.cPointer(), (uint16_t)(port), mode))

}
func (this *QTcpSocket) OnBind(slot func(super func(address *QHostAddress, port uint16, mode BindMode) bool, address *QHostAddress, port uint16, mode BindMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_Bind(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_Bind
func miqt_exec_callback_QTcpSocket_Bind(self QTcpSocket, cb intptr_t, address *QHostAddress, port uint16_t, mode BindMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(address *QHostAddress, port uint16, mode BindMode) bool, address *QHostAddress, port uint16, mode BindMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHostAddress(address)

	slotval2 := (uint16)(port)

	xxxxxxxxx

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_Bind, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_ConnectToHost(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))

	QTcpSocket_virtualbase_ConnectToHost(unsafe.Pointer(this.h), hostName_ms, (uint16_t)(port), mode, protocol)

}
func (this *QTcpSocket) OnConnectToHost(slot func(super func(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol), hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_ConnectToHost(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_ConnectToHost
func miqt_exec_callback_QTcpSocket_ConnectToHost(self QTcpSocket, cb intptr_t, hostName struct_miqt_string, port uint16_t, mode OpenMode, protocol NetworkLayerProtocol) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol), hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var hostName_ms struct_miqt_string = hostName
	hostName_ret := GoStringN(hostName_ms.data, int(int64(hostName_ms.len)))
	free(unsafe.Pointer(hostName_ms.data))
	slotval1 := hostName_ret
	slotval2 := (uint16)(port)

	xxxxxxxxx
	xxxxxxxxx

	gofunc((&QTcpSocket{h: self}).callVirtualBase_ConnectToHost, slotval1, slotval2, slotval3, slotval4)

}

func (this *QTcpSocket) callVirtualBase_DisconnectFromHost() {

	QTcpSocket_virtualbase_DisconnectFromHost(unsafe.Pointer(this.h))

}
func (this *QTcpSocket) OnDisconnectFromHost(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_DisconnectFromHost(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_DisconnectFromHost
func miqt_exec_callback_QTcpSocket_DisconnectFromHost(self QTcpSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTcpSocket{h: self}).callVirtualBase_DisconnectFromHost)

}

func (this *QTcpSocket) callVirtualBase_BytesAvailable() int64 {

	return (int64)(QTcpSocket_virtualbase_BytesAvailable(unsafe.Pointer(this.h)))

}
func (this *QTcpSocket) OnBytesAvailable(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_BytesAvailable(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_BytesAvailable
func miqt_exec_callback_QTcpSocket_BytesAvailable(self QTcpSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_BytesAvailable)

	return (longlong)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_BytesToWrite() int64 {

	return (int64)(QTcpSocket_virtualbase_BytesToWrite(unsafe.Pointer(this.h)))

}
func (this *QTcpSocket) OnBytesToWrite(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_BytesToWrite(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_BytesToWrite
func miqt_exec_callback_QTcpSocket_BytesToWrite(self QTcpSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_BytesToWrite)

	return (longlong)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_SetReadBufferSize(size int64) {

	QTcpSocket_virtualbase_SetReadBufferSize(unsafe.Pointer(this.h), (longlong)(size))

}
func (this *QTcpSocket) OnSetReadBufferSize(slot func(super func(size int64), size int64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_SetReadBufferSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_SetReadBufferSize
func miqt_exec_callback_QTcpSocket_SetReadBufferSize(self QTcpSocket, cb intptr_t, size longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size int64), size int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(size)

	gofunc((&QTcpSocket{h: self}).callVirtualBase_SetReadBufferSize, slotval1)

}

func (this *QTcpSocket) callVirtualBase_SocketDescriptor() uintptr {

	return (uintptr)(QTcpSocket_virtualbase_SocketDescriptor(unsafe.Pointer(this.h)))

}
func (this *QTcpSocket) OnSocketDescriptor(slot func(super func() uintptr) uintptr) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_SocketDescriptor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_SocketDescriptor
func miqt_exec_callback_QTcpSocket_SocketDescriptor(self QTcpSocket, cb intptr_t) intptr_t {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() uintptr) uintptr)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_SocketDescriptor)

	return (intptr_t)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_SetSocketDescriptor(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool {

	return (bool)(QTcpSocket_virtualbase_SetSocketDescriptor(unsafe.Pointer(this.h), (intptr_t)(socketDescriptor), state, openMode))

}
func (this *QTcpSocket) OnSetSocketDescriptor(slot func(super func(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool, socketDescriptor uintptr, state SocketState, openMode OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_SetSocketDescriptor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_SetSocketDescriptor
func miqt_exec_callback_QTcpSocket_SetSocketDescriptor(self QTcpSocket, cb intptr_t, socketDescriptor intptr_t, state SocketState, openMode OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool, socketDescriptor uintptr, state SocketState, openMode OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uintptr)(socketDescriptor)

	xxxxxxxxx
	xxxxxxxxx

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_SetSocketDescriptor, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_SetSocketOption(option QAbstractSocket__SocketOption, value *qt.QVariant) {

	QTcpSocket_virtualbase_SetSocketOption(unsafe.Pointer(this.h), (int)(option), (*QVariant)(value.UnsafePointer()))

}
func (this *QTcpSocket) OnSetSocketOption(slot func(super func(option QAbstractSocket__SocketOption, value *qt.QVariant), option QAbstractSocket__SocketOption, value *qt.QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_SetSocketOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_SetSocketOption
func miqt_exec_callback_QTcpSocket_SetSocketOption(self QTcpSocket, cb intptr_t, option int, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option QAbstractSocket__SocketOption, value *qt.QVariant), option QAbstractSocket__SocketOption, value *qt.QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketOption)(option)

	slotval2 := qt.UnsafeNewQVariant(unsafe.Pointer(value))

	gofunc((&QTcpSocket{h: self}).callVirtualBase_SetSocketOption, slotval1, slotval2)

}

func (this *QTcpSocket) callVirtualBase_SocketOption(option QAbstractSocket__SocketOption) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QTcpSocket_virtualbase_SocketOption(unsafe.Pointer(this.h), (int)(option))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTcpSocket) OnSocketOption(slot func(super func(option QAbstractSocket__SocketOption) *qt.QVariant, option QAbstractSocket__SocketOption) *qt.QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_SocketOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_SocketOption
func miqt_exec_callback_QTcpSocket_SocketOption(self QTcpSocket, cb intptr_t, option int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option QAbstractSocket__SocketOption) *qt.QVariant, option QAbstractSocket__SocketOption) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketOption)(option)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_SocketOption, slotval1)

	return (*QVariant)(virtualReturn.UnsafePointer())

}

func (this *QTcpSocket) callVirtualBase_Close() {

	QTcpSocket_virtualbase_Close(unsafe.Pointer(this.h))

}
func (this *QTcpSocket) OnClose(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_Close(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_Close
func miqt_exec_callback_QTcpSocket_Close(self QTcpSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTcpSocket{h: self}).callVirtualBase_Close)

}

func (this *QTcpSocket) callVirtualBase_IsSequential() bool {

	return (bool)(QTcpSocket_virtualbase_IsSequential(unsafe.Pointer(this.h)))

}
func (this *QTcpSocket) OnIsSequential(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_IsSequential
func miqt_exec_callback_QTcpSocket_IsSequential(self QTcpSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_IsSequential)

	return (bool)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_WaitForConnected(msecs int) bool {

	return (bool)(QTcpSocket_virtualbase_WaitForConnected(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QTcpSocket) OnWaitForConnected(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_WaitForConnected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_WaitForConnected
func miqt_exec_callback_QTcpSocket_WaitForConnected(self QTcpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_WaitForConnected, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_WaitForReadyRead(msecs int) bool {

	return (bool)(QTcpSocket_virtualbase_WaitForReadyRead(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QTcpSocket) OnWaitForReadyRead(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_WaitForReadyRead(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_WaitForReadyRead
func miqt_exec_callback_QTcpSocket_WaitForReadyRead(self QTcpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_WaitForReadyRead, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_WaitForBytesWritten(msecs int) bool {

	return (bool)(QTcpSocket_virtualbase_WaitForBytesWritten(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QTcpSocket) OnWaitForBytesWritten(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_WaitForBytesWritten(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_WaitForBytesWritten
func miqt_exec_callback_QTcpSocket_WaitForBytesWritten(self QTcpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_WaitForBytesWritten, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_WaitForDisconnected(msecs int) bool {

	return (bool)(QTcpSocket_virtualbase_WaitForDisconnected(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QTcpSocket) OnWaitForDisconnected(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_WaitForDisconnected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_WaitForDisconnected
func miqt_exec_callback_QTcpSocket_WaitForDisconnected(self QTcpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_WaitForDisconnected, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_ReadData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QTcpSocket_virtualbase_ReadData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QTcpSocket) OnReadData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_ReadData
func miqt_exec_callback_QTcpSocket_ReadData(self QTcpSocket, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_ReadData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_ReadLineData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QTcpSocket_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QTcpSocket) OnReadLineData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_ReadLineData
func miqt_exec_callback_QTcpSocket_ReadLineData(self QTcpSocket, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_SkipData(maxSize int64) int64 {

	return (int64)(QTcpSocket_virtualbase_SkipData(unsafe.Pointer(this.h), (longlong)(maxSize)))

}
func (this *QTcpSocket) OnSkipData(slot func(super func(maxSize int64) int64, maxSize int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_SkipData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_SkipData
func miqt_exec_callback_QTcpSocket_SkipData(self QTcpSocket, cb intptr_t, maxSize longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(maxSize int64) int64, maxSize int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(maxSize)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_SkipData, slotval1)

	return (longlong)(virtualReturn)

}

func (this *QTcpSocket) callVirtualBase_WriteData(data string, lenVal int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QTcpSocket_virtualbase_WriteData(unsafe.Pointer(this.h), data_Cstring, (longlong)(lenVal)))

}
func (this *QTcpSocket) OnWriteData(slot func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTcpSocket_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTcpSocket_WriteData
func miqt_exec_callback_QTcpSocket_WriteData(self QTcpSocket, cb intptr_t, data *const_char, lenVal longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(lenVal)

	virtualReturn := gofunc((&QTcpSocket{h: self}).callVirtualBase_WriteData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}
