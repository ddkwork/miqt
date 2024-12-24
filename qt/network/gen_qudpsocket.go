package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QUdpSocket struct {
	h          uintptr
	isSubclass bool
}

// NewQUdpSocket constructs a new QUdpSocket object.
func NewQUdpSocket() *QUdpSocket {

	ret := newQUdpSocket(QUdpSocket_new())
	ret.isSubclass = true
	return ret
}

// NewQUdpSocket2 constructs a new QUdpSocket object.
func NewQUdpSocket2(parent *qt.QObject) *QUdpSocket {

	ret := newQUdpSocket(QUdpSocket_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QUdpSocket) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QUdpSocket_MetaObject(this.h)))
}

func (this *QUdpSocket) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QUdpSocket_Metacast(this.h, param1_Cstring))
}

func QUdpSocket_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QUdpSocket_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUdpSocket) Bind(addr QHostAddress__SpecialAddress) bool {
	return (bool)(QUdpSocket_Bind(this.h, (int)(addr)))
}

func (this *QUdpSocket) JoinMulticastGroup(groupAddress *QHostAddress) bool {
	return (bool)(QUdpSocket_JoinMulticastGroup(this.h, groupAddress.cPointer()))
}

func (this *QUdpSocket) JoinMulticastGroup2(groupAddress *QHostAddress, iface *QNetworkInterface) bool {
	return (bool)(QUdpSocket_JoinMulticastGroup2(this.h, groupAddress.cPointer(), iface.cPointer()))
}

func (this *QUdpSocket) LeaveMulticastGroup(groupAddress *QHostAddress) bool {
	return (bool)(QUdpSocket_LeaveMulticastGroup(this.h, groupAddress.cPointer()))
}

func (this *QUdpSocket) LeaveMulticastGroup2(groupAddress *QHostAddress, iface *QNetworkInterface) bool {
	return (bool)(QUdpSocket_LeaveMulticastGroup2(this.h, groupAddress.cPointer(), iface.cPointer()))
}

func (this *QUdpSocket) MulticastInterface() *QNetworkInterface {
	_goptr := newQNetworkInterface(QUdpSocket_MulticastInterface(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUdpSocket) SetMulticastInterface(iface *QNetworkInterface) {
	QUdpSocket_SetMulticastInterface(this.h, iface.cPointer())
}

func (this *QUdpSocket) HasPendingDatagrams() bool {
	return (bool)(QUdpSocket_HasPendingDatagrams(this.h))
}

func (this *QUdpSocket) PendingDatagramSize() int64 {
	return (int64)(QUdpSocket_PendingDatagramSize(this.h))
}

func (this *QUdpSocket) ReceiveDatagram() *QNetworkDatagram {
	_goptr := newQNetworkDatagram(QUdpSocket_ReceiveDatagram(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUdpSocket) ReadDatagram(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_ReadDatagram(this.h, data_Cstring, (longlong)(maxlen)))
}

func (this *QUdpSocket) WriteDatagram(datagram *QNetworkDatagram) int64 {
	return (int64)(QUdpSocket_WriteDatagram(this.h, datagram.cPointer()))
}

func (this *QUdpSocket) WriteDatagram2(data string, lenVal int64, host *QHostAddress, port uint16) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_WriteDatagram2(this.h, data_Cstring, (longlong)(lenVal), host.cPointer(), (uint16_t)(port)))
}

func (this *QUdpSocket) WriteDatagram3(datagram []byte, host *QHostAddress, port uint16) int64 {
	datagram_alias := struct_miqt_string{}
	datagram_alias.data = (char)(unsafe.Pointer(&datagram[0]))
	datagram_alias.len = size_t(len(datagram))
	return (int64)(QUdpSocket_WriteDatagram3(this.h, datagram_alias, host.cPointer(), (uint16_t)(port)))
}

func QUdpSocket_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUdpSocket_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUdpSocket_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUdpSocket_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUdpSocket) Bind2(addr QHostAddress__SpecialAddress, port uint16) bool {
	return (bool)(QUdpSocket_Bind2(this.h, (int)(addr), (uint16_t)(port)))
}

func (this *QUdpSocket) Bind3(addr QHostAddress__SpecialAddress, port uint16, mode BindMode) bool {
	return (bool)(QUdpSocket_Bind3(this.h, (int)(addr), (uint16_t)(port), mode))
}

func (this *QUdpSocket) ReceiveDatagram1(maxSize int64) *QNetworkDatagram {
	_goptr := newQNetworkDatagram(QUdpSocket_ReceiveDatagram1(this.h, (longlong)(maxSize)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUdpSocket) ReadDatagram3(data string, maxlen int64, host *QHostAddress) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_ReadDatagram3(this.h, data_Cstring, (longlong)(maxlen), host.cPointer()))
}

func (this *QUdpSocket) ReadDatagram4(data string, maxlen int64, host *QHostAddress, port *uint16) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_ReadDatagram4(this.h, data_Cstring, (longlong)(maxlen), host.cPointer(), (*uint16_t)(unsafe.Pointer(port))))
}

func (this *QUdpSocket) callVirtualBase_Resume() {

	QUdpSocket_virtualbase_Resume(unsafe.Pointer(this.h))

}
func (this *QUdpSocket) OnResume(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_Resume(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_Resume
func miqt_exec_callback_QUdpSocket_Resume(self QUdpSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUdpSocket{h: self}).callVirtualBase_Resume)

}

func (this *QUdpSocket) callVirtualBase_Bind(address *QHostAddress, port uint16, mode BindMode) bool {

	return (bool)(QUdpSocket_virtualbase_Bind(unsafe.Pointer(this.h), address.cPointer(), (uint16_t)(port), mode))

}
func (this *QUdpSocket) OnBind(slot func(super func(address *QHostAddress, port uint16, mode BindMode) bool, address *QHostAddress, port uint16, mode BindMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_Bind(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_Bind
func miqt_exec_callback_QUdpSocket_Bind(self QUdpSocket, cb intptr_t, address *QHostAddress, port uint16_t, mode BindMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(address *QHostAddress, port uint16, mode BindMode) bool, address *QHostAddress, port uint16, mode BindMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHostAddress(address)

	slotval2 := (uint16)(port)

	xxxxxxxxx

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_Bind, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_ConnectToHost(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))

	QUdpSocket_virtualbase_ConnectToHost(unsafe.Pointer(this.h), hostName_ms, (uint16_t)(port), mode, protocol)

}
func (this *QUdpSocket) OnConnectToHost(slot func(super func(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol), hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_ConnectToHost(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_ConnectToHost
func miqt_exec_callback_QUdpSocket_ConnectToHost(self QUdpSocket, cb intptr_t, hostName struct_miqt_string, port uint16_t, mode OpenMode, protocol NetworkLayerProtocol) {
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

	gofunc((&QUdpSocket{h: self}).callVirtualBase_ConnectToHost, slotval1, slotval2, slotval3, slotval4)

}

func (this *QUdpSocket) callVirtualBase_DisconnectFromHost() {

	QUdpSocket_virtualbase_DisconnectFromHost(unsafe.Pointer(this.h))

}
func (this *QUdpSocket) OnDisconnectFromHost(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_DisconnectFromHost(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_DisconnectFromHost
func miqt_exec_callback_QUdpSocket_DisconnectFromHost(self QUdpSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUdpSocket{h: self}).callVirtualBase_DisconnectFromHost)

}

func (this *QUdpSocket) callVirtualBase_BytesAvailable() int64 {

	return (int64)(QUdpSocket_virtualbase_BytesAvailable(unsafe.Pointer(this.h)))

}
func (this *QUdpSocket) OnBytesAvailable(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_BytesAvailable(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_BytesAvailable
func miqt_exec_callback_QUdpSocket_BytesAvailable(self QUdpSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_BytesAvailable)

	return (longlong)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_BytesToWrite() int64 {

	return (int64)(QUdpSocket_virtualbase_BytesToWrite(unsafe.Pointer(this.h)))

}
func (this *QUdpSocket) OnBytesToWrite(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_BytesToWrite(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_BytesToWrite
func miqt_exec_callback_QUdpSocket_BytesToWrite(self QUdpSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_BytesToWrite)

	return (longlong)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_SetReadBufferSize(size int64) {

	QUdpSocket_virtualbase_SetReadBufferSize(unsafe.Pointer(this.h), (longlong)(size))

}
func (this *QUdpSocket) OnSetReadBufferSize(slot func(super func(size int64), size int64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_SetReadBufferSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_SetReadBufferSize
func miqt_exec_callback_QUdpSocket_SetReadBufferSize(self QUdpSocket, cb intptr_t, size longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size int64), size int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(size)

	gofunc((&QUdpSocket{h: self}).callVirtualBase_SetReadBufferSize, slotval1)

}

func (this *QUdpSocket) callVirtualBase_SocketDescriptor() uintptr {

	return (uintptr)(QUdpSocket_virtualbase_SocketDescriptor(unsafe.Pointer(this.h)))

}
func (this *QUdpSocket) OnSocketDescriptor(slot func(super func() uintptr) uintptr) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_SocketDescriptor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_SocketDescriptor
func miqt_exec_callback_QUdpSocket_SocketDescriptor(self QUdpSocket, cb intptr_t) intptr_t {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() uintptr) uintptr)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_SocketDescriptor)

	return (intptr_t)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_SetSocketDescriptor(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool {

	return (bool)(QUdpSocket_virtualbase_SetSocketDescriptor(unsafe.Pointer(this.h), (intptr_t)(socketDescriptor), state, openMode))

}
func (this *QUdpSocket) OnSetSocketDescriptor(slot func(super func(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool, socketDescriptor uintptr, state SocketState, openMode OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_SetSocketDescriptor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_SetSocketDescriptor
func miqt_exec_callback_QUdpSocket_SetSocketDescriptor(self QUdpSocket, cb intptr_t, socketDescriptor intptr_t, state SocketState, openMode OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool, socketDescriptor uintptr, state SocketState, openMode OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uintptr)(socketDescriptor)

	xxxxxxxxx
	xxxxxxxxx

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_SetSocketDescriptor, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_SetSocketOption(option QAbstractSocket__SocketOption, value *qt.QVariant) {

	QUdpSocket_virtualbase_SetSocketOption(unsafe.Pointer(this.h), (int)(option), (*QVariant)(value.UnsafePointer()))

}
func (this *QUdpSocket) OnSetSocketOption(slot func(super func(option QAbstractSocket__SocketOption, value *qt.QVariant), option QAbstractSocket__SocketOption, value *qt.QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_SetSocketOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_SetSocketOption
func miqt_exec_callback_QUdpSocket_SetSocketOption(self QUdpSocket, cb intptr_t, option int, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option QAbstractSocket__SocketOption, value *qt.QVariant), option QAbstractSocket__SocketOption, value *qt.QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketOption)(option)

	slotval2 := qt.UnsafeNewQVariant(unsafe.Pointer(value))

	gofunc((&QUdpSocket{h: self}).callVirtualBase_SetSocketOption, slotval1, slotval2)

}

func (this *QUdpSocket) callVirtualBase_SocketOption(option QAbstractSocket__SocketOption) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QUdpSocket_virtualbase_SocketOption(unsafe.Pointer(this.h), (int)(option))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QUdpSocket) OnSocketOption(slot func(super func(option QAbstractSocket__SocketOption) *qt.QVariant, option QAbstractSocket__SocketOption) *qt.QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_SocketOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_SocketOption
func miqt_exec_callback_QUdpSocket_SocketOption(self QUdpSocket, cb intptr_t, option int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option QAbstractSocket__SocketOption) *qt.QVariant, option QAbstractSocket__SocketOption) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketOption)(option)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_SocketOption, slotval1)

	return (*QVariant)(virtualReturn.UnsafePointer())

}

func (this *QUdpSocket) callVirtualBase_Close() {

	QUdpSocket_virtualbase_Close(unsafe.Pointer(this.h))

}
func (this *QUdpSocket) OnClose(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_Close(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_Close
func miqt_exec_callback_QUdpSocket_Close(self QUdpSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUdpSocket{h: self}).callVirtualBase_Close)

}

func (this *QUdpSocket) callVirtualBase_IsSequential() bool {

	return (bool)(QUdpSocket_virtualbase_IsSequential(unsafe.Pointer(this.h)))

}
func (this *QUdpSocket) OnIsSequential(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_IsSequential
func miqt_exec_callback_QUdpSocket_IsSequential(self QUdpSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_IsSequential)

	return (bool)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_WaitForConnected(msecs int) bool {

	return (bool)(QUdpSocket_virtualbase_WaitForConnected(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QUdpSocket) OnWaitForConnected(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_WaitForConnected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_WaitForConnected
func miqt_exec_callback_QUdpSocket_WaitForConnected(self QUdpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_WaitForConnected, slotval1)

	return (bool)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_WaitForReadyRead(msecs int) bool {

	return (bool)(QUdpSocket_virtualbase_WaitForReadyRead(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QUdpSocket) OnWaitForReadyRead(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_WaitForReadyRead(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_WaitForReadyRead
func miqt_exec_callback_QUdpSocket_WaitForReadyRead(self QUdpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_WaitForReadyRead, slotval1)

	return (bool)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_WaitForBytesWritten(msecs int) bool {

	return (bool)(QUdpSocket_virtualbase_WaitForBytesWritten(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QUdpSocket) OnWaitForBytesWritten(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_WaitForBytesWritten(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_WaitForBytesWritten
func miqt_exec_callback_QUdpSocket_WaitForBytesWritten(self QUdpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_WaitForBytesWritten, slotval1)

	return (bool)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_WaitForDisconnected(msecs int) bool {

	return (bool)(QUdpSocket_virtualbase_WaitForDisconnected(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QUdpSocket) OnWaitForDisconnected(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_WaitForDisconnected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_WaitForDisconnected
func miqt_exec_callback_QUdpSocket_WaitForDisconnected(self QUdpSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_WaitForDisconnected, slotval1)

	return (bool)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_ReadData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QUdpSocket_virtualbase_ReadData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QUdpSocket) OnReadData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_ReadData
func miqt_exec_callback_QUdpSocket_ReadData(self QUdpSocket, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_ReadData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_ReadLineData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QUdpSocket_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QUdpSocket) OnReadLineData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_ReadLineData
func miqt_exec_callback_QUdpSocket_ReadLineData(self QUdpSocket, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_SkipData(maxSize int64) int64 {

	return (int64)(QUdpSocket_virtualbase_SkipData(unsafe.Pointer(this.h), (longlong)(maxSize)))

}
func (this *QUdpSocket) OnSkipData(slot func(super func(maxSize int64) int64, maxSize int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_SkipData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_SkipData
func miqt_exec_callback_QUdpSocket_SkipData(self QUdpSocket, cb intptr_t, maxSize longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(maxSize int64) int64, maxSize int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(maxSize)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_SkipData, slotval1)

	return (longlong)(virtualReturn)

}

func (this *QUdpSocket) callVirtualBase_WriteData(data string, lenVal int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QUdpSocket_virtualbase_WriteData(unsafe.Pointer(this.h), data_Cstring, (longlong)(lenVal)))

}
func (this *QUdpSocket) OnWriteData(slot func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_WriteData
func miqt_exec_callback_QUdpSocket_WriteData(self QUdpSocket, cb intptr_t, data *const_char, lenVal longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(lenVal)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_WriteData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}
