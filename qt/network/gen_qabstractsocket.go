package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QAbstractSocket__SocketType int

const (
	QAbstractSocket__TcpSocket         QAbstractSocket__SocketType = 0
	QAbstractSocket__UdpSocket         QAbstractSocket__SocketType = 1
	QAbstractSocket__SctpSocket        QAbstractSocket__SocketType = 2
	QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = -1
)

type QAbstractSocket__NetworkLayerProtocol int

const (
	QAbstractSocket__IPv4Protocol                QAbstractSocket__NetworkLayerProtocol = 0
	QAbstractSocket__IPv6Protocol                QAbstractSocket__NetworkLayerProtocol = 1
	QAbstractSocket__AnyIPProtocol               QAbstractSocket__NetworkLayerProtocol = 2
	QAbstractSocket__UnknownNetworkLayerProtocol QAbstractSocket__NetworkLayerProtocol = -1
)

type QAbstractSocket__SocketError int

const (
	QAbstractSocket__ConnectionRefusedError           QAbstractSocket__SocketError = 0
	QAbstractSocket__RemoteHostClosedError            QAbstractSocket__SocketError = 1
	QAbstractSocket__HostNotFoundError                QAbstractSocket__SocketError = 2
	QAbstractSocket__SocketAccessError                QAbstractSocket__SocketError = 3
	QAbstractSocket__SocketResourceError              QAbstractSocket__SocketError = 4
	QAbstractSocket__SocketTimeoutError               QAbstractSocket__SocketError = 5
	QAbstractSocket__DatagramTooLargeError            QAbstractSocket__SocketError = 6
	QAbstractSocket__NetworkError                     QAbstractSocket__SocketError = 7
	QAbstractSocket__AddressInUseError                QAbstractSocket__SocketError = 8
	QAbstractSocket__SocketAddressNotAvailableError   QAbstractSocket__SocketError = 9
	QAbstractSocket__UnsupportedSocketOperationError  QAbstractSocket__SocketError = 10
	QAbstractSocket__UnfinishedSocketOperationError   QAbstractSocket__SocketError = 11
	QAbstractSocket__ProxyAuthenticationRequiredError QAbstractSocket__SocketError = 12
	QAbstractSocket__SslHandshakeFailedError          QAbstractSocket__SocketError = 13
	QAbstractSocket__ProxyConnectionRefusedError      QAbstractSocket__SocketError = 14
	QAbstractSocket__ProxyConnectionClosedError       QAbstractSocket__SocketError = 15
	QAbstractSocket__ProxyConnectionTimeoutError      QAbstractSocket__SocketError = 16
	QAbstractSocket__ProxyNotFoundError               QAbstractSocket__SocketError = 17
	QAbstractSocket__ProxyProtocolError               QAbstractSocket__SocketError = 18
	QAbstractSocket__OperationError                   QAbstractSocket__SocketError = 19
	QAbstractSocket__SslInternalError                 QAbstractSocket__SocketError = 20
	QAbstractSocket__SslInvalidUserDataError          QAbstractSocket__SocketError = 21
	QAbstractSocket__TemporaryError                   QAbstractSocket__SocketError = 22
	QAbstractSocket__UnknownSocketError               QAbstractSocket__SocketError = -1
)

type QAbstractSocket__SocketState int

const (
	QAbstractSocket__UnconnectedState QAbstractSocket__SocketState = 0
	QAbstractSocket__HostLookupState  QAbstractSocket__SocketState = 1
	QAbstractSocket__ConnectingState  QAbstractSocket__SocketState = 2
	QAbstractSocket__ConnectedState   QAbstractSocket__SocketState = 3
	QAbstractSocket__BoundState       QAbstractSocket__SocketState = 4
	QAbstractSocket__ListeningState   QAbstractSocket__SocketState = 5
	QAbstractSocket__ClosingState     QAbstractSocket__SocketState = 6
)

type QAbstractSocket__SocketOption int

const (
	QAbstractSocket__LowDelayOption                QAbstractSocket__SocketOption = 0
	QAbstractSocket__KeepAliveOption               QAbstractSocket__SocketOption = 1
	QAbstractSocket__MulticastTtlOption            QAbstractSocket__SocketOption = 2
	QAbstractSocket__MulticastLoopbackOption       QAbstractSocket__SocketOption = 3
	QAbstractSocket__TypeOfServiceOption           QAbstractSocket__SocketOption = 4
	QAbstractSocket__SendBufferSizeSocketOption    QAbstractSocket__SocketOption = 5
	QAbstractSocket__ReceiveBufferSizeSocketOption QAbstractSocket__SocketOption = 6
	QAbstractSocket__PathMtuSocketOption           QAbstractSocket__SocketOption = 7
)

type QAbstractSocket__BindFlag int

const (
	QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = 0
	QAbstractSocket__ShareAddress       QAbstractSocket__BindFlag = 1
	QAbstractSocket__DontShareAddress   QAbstractSocket__BindFlag = 2
	QAbstractSocket__ReuseAddressHint   QAbstractSocket__BindFlag = 4
)

type QAbstractSocket__PauseMode int

const (
	QAbstractSocket__PauseNever       QAbstractSocket__PauseMode = 0
	QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = 1
)

type QAbstractSocket struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractSocket constructs a new QAbstractSocket object.
func NewQAbstractSocket(socketType SocketType, parent *qt.QObject) *QAbstractSocket {

	ret := newQAbstractSocket(QAbstractSocket_new(socketType, (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractSocket) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAbstractSocket_MetaObject(this.h)))
}

func (this *QAbstractSocket) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractSocket_Metacast(this.h, param1_Cstring))
}

func QAbstractSocket_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractSocket_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSocket) Resume() {
	QAbstractSocket_Resume(this.h)
}

func (this *QAbstractSocket) PauseMode() PauseModes {
	xxxxxxxxx
}

func (this *QAbstractSocket) SetPauseMode(pauseMode PauseModes) {
	QAbstractSocket_SetPauseMode(this.h, pauseMode)
}

func (this *QAbstractSocket) Bind(address *QHostAddress, port uint16, mode BindMode) bool {
	return (bool)(QAbstractSocket_Bind(this.h, address.cPointer(), (uint16_t)(port), mode))
}

func (this *QAbstractSocket) Bind2() bool {
	return (bool)(QAbstractSocket_Bind2(this.h))
}

func (this *QAbstractSocket) ConnectToHost(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QAbstractSocket_ConnectToHost(this.h, hostName_ms, (uint16_t)(port), mode, protocol)
}

func (this *QAbstractSocket) ConnectToHost2(address *QHostAddress, port uint16) {
	QAbstractSocket_ConnectToHost2(this.h, address.cPointer(), (uint16_t)(port))
}

func (this *QAbstractSocket) DisconnectFromHost() {
	QAbstractSocket_DisconnectFromHost(this.h)
}

func (this *QAbstractSocket) IsValid() bool {
	return (bool)(QAbstractSocket_IsValid(this.h))
}

func (this *QAbstractSocket) BytesAvailable() int64 {
	return (int64)(QAbstractSocket_BytesAvailable(this.h))
}

func (this *QAbstractSocket) BytesToWrite() int64 {
	return (int64)(QAbstractSocket_BytesToWrite(this.h))
}

func (this *QAbstractSocket) LocalPort() uint16 {
	return (uint16)(QAbstractSocket_LocalPort(this.h))
}

func (this *QAbstractSocket) LocalAddress() *QHostAddress {
	_goptr := newQHostAddress(QAbstractSocket_LocalAddress(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractSocket) PeerPort() uint16 {
	return (uint16)(QAbstractSocket_PeerPort(this.h))
}

func (this *QAbstractSocket) PeerAddress() *QHostAddress {
	_goptr := newQHostAddress(QAbstractSocket_PeerAddress(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractSocket) PeerName() string {
	var _ms struct_miqt_string = QAbstractSocket_PeerName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSocket) ReadBufferSize() int64 {
	return (int64)(QAbstractSocket_ReadBufferSize(this.h))
}

func (this *QAbstractSocket) SetReadBufferSize(size int64) {
	QAbstractSocket_SetReadBufferSize(this.h, (longlong)(size))
}

func (this *QAbstractSocket) Abort() {
	QAbstractSocket_Abort(this.h)
}

func (this *QAbstractSocket) SocketDescriptor() uintptr {
	return (uintptr)(QAbstractSocket_SocketDescriptor(this.h))
}

func (this *QAbstractSocket) SetSocketDescriptor(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool {
	return (bool)(QAbstractSocket_SetSocketDescriptor(this.h, (intptr_t)(socketDescriptor), state, openMode))
}

func (this *QAbstractSocket) SetSocketOption(option QAbstractSocket__SocketOption, value *qt.QVariant) {
	QAbstractSocket_SetSocketOption(this.h, (int)(option), (*QVariant)(value.UnsafePointer()))
}

func (this *QAbstractSocket) SocketOption(option QAbstractSocket__SocketOption) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QAbstractSocket_SocketOption(this.h, (int)(option))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractSocket) SocketType() SocketType {
	xxxxxxxxx
}

func (this *QAbstractSocket) State() SocketState {
	xxxxxxxxx
}

func (this *QAbstractSocket) Error() SocketError {
	xxxxxxxxx
}

func (this *QAbstractSocket) Close() {
	QAbstractSocket_Close(this.h)
}

func (this *QAbstractSocket) IsSequential() bool {
	return (bool)(QAbstractSocket_IsSequential(this.h))
}

func (this *QAbstractSocket) Flush() bool {
	return (bool)(QAbstractSocket_Flush(this.h))
}

func (this *QAbstractSocket) WaitForConnected(msecs int) bool {
	return (bool)(QAbstractSocket_WaitForConnected(this.h, (int)(msecs)))
}

func (this *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	return (bool)(QAbstractSocket_WaitForReadyRead(this.h, (int)(msecs)))
}

func (this *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	return (bool)(QAbstractSocket_WaitForBytesWritten(this.h, (int)(msecs)))
}

func (this *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	return (bool)(QAbstractSocket_WaitForDisconnected(this.h, (int)(msecs)))
}

func (this *QAbstractSocket) SetProxy(networkProxy *QNetworkProxy) {
	QAbstractSocket_SetProxy(this.h, networkProxy.cPointer())
}

func (this *QAbstractSocket) Proxy() *QNetworkProxy {
	_goptr := newQNetworkProxy(QAbstractSocket_Proxy(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractSocket) ProtocolTag() string {
	var _ms struct_miqt_string = QAbstractSocket_ProtocolTag(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSocket) SetProtocolTag(tag string) {
	tag_ms := struct_miqt_string{}
	tag_ms.data = CString(tag)
	tag_ms.len = size_t(len(tag))
	defer free(unsafe.Pointer(tag_ms.data))
	QAbstractSocket_SetProtocolTag(this.h, tag_ms)
}

func (this *QAbstractSocket) HostFound() {
	QAbstractSocket_HostFound(this.h)
}
func (this *QAbstractSocket) OnHostFound(slot func()) {
	QAbstractSocket_connect_HostFound(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_HostFound
func miqt_exec_callback_QAbstractSocket_HostFound(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractSocket) Connected() {
	QAbstractSocket_Connected(this.h)
}
func (this *QAbstractSocket) OnConnected(slot func()) {
	QAbstractSocket_connect_Connected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Connected
func miqt_exec_callback_QAbstractSocket_Connected(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractSocket) Disconnected() {
	QAbstractSocket_Disconnected(this.h)
}
func (this *QAbstractSocket) OnDisconnected(slot func()) {
	QAbstractSocket_connect_Disconnected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Disconnected
func miqt_exec_callback_QAbstractSocket_Disconnected(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractSocket) StateChanged(param1 QAbstractSocket__SocketState) {
	QAbstractSocket_StateChanged(this.h, (int)(param1))
}
func (this *QAbstractSocket) OnStateChanged(slot func(param1 QAbstractSocket__SocketState)) {
	QAbstractSocket_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_StateChanged
func miqt_exec_callback_QAbstractSocket_StateChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QAbstractSocket__SocketState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketState)(param1)

	gofunc(slotval1)
}

func (this *QAbstractSocket) ErrorOccurred(param1 QAbstractSocket__SocketError) {
	QAbstractSocket_ErrorOccurred(this.h, (int)(param1))
}
func (this *QAbstractSocket) OnErrorOccurred(slot func(param1 QAbstractSocket__SocketError)) {
	QAbstractSocket_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_ErrorOccurred
func miqt_exec_callback_QAbstractSocket_ErrorOccurred(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QAbstractSocket__SocketError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketError)(param1)

	gofunc(slotval1)
}

func (this *QAbstractSocket) ProxyAuthenticationRequired(proxy *QNetworkProxy, authenticator *QAuthenticator) {
	QAbstractSocket_ProxyAuthenticationRequired(this.h, proxy.cPointer(), authenticator.cPointer())
}
func (this *QAbstractSocket) OnProxyAuthenticationRequired(slot func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	QAbstractSocket_connect_ProxyAuthenticationRequired(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_ProxyAuthenticationRequired
func miqt_exec_callback_QAbstractSocket_ProxyAuthenticationRequired(cb intptr_t, proxy *QNetworkProxy, authenticator *QAuthenticator) {
	gofunc, ok := cgo.Handle(cb).Value().(func(proxy *QNetworkProxy, authenticator *QAuthenticator))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkProxy(proxy)

	slotval2 := newQAuthenticator(authenticator)

	gofunc(slotval1, slotval2)
}

func QAbstractSocket_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractSocket_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractSocket_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractSocket_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSocket) Bind1(port uint16) bool {
	return (bool)(QAbstractSocket_Bind1(this.h, (uint16_t)(port)))
}

func (this *QAbstractSocket) Bind22(port uint16, mode BindMode) bool {
	return (bool)(QAbstractSocket_Bind22(this.h, (uint16_t)(port), mode))
}

func (this *QAbstractSocket) ConnectToHost3(address *QHostAddress, port uint16, mode OpenMode) {
	QAbstractSocket_ConnectToHost3(this.h, address.cPointer(), (uint16_t)(port), mode)
}

func (this *QAbstractSocket) callVirtualBase_Resume() {

	QAbstractSocket_virtualbase_Resume(unsafe.Pointer(this.h))

}
func (this *QAbstractSocket) OnResume(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Resume(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Resume
func miqt_exec_callback_QAbstractSocket_Resume(self QAbstractSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractSocket{h: self}).callVirtualBase_Resume)

}

func (this *QAbstractSocket) callVirtualBase_Bind(address *QHostAddress, port uint16, mode BindMode) bool {

	return (bool)(QAbstractSocket_virtualbase_Bind(unsafe.Pointer(this.h), address.cPointer(), (uint16_t)(port), mode))

}
func (this *QAbstractSocket) OnBind(slot func(super func(address *QHostAddress, port uint16, mode BindMode) bool, address *QHostAddress, port uint16, mode BindMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Bind(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Bind
func miqt_exec_callback_QAbstractSocket_Bind(self QAbstractSocket, cb intptr_t, address *QHostAddress, port uint16_t, mode BindMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(address *QHostAddress, port uint16, mode BindMode) bool, address *QHostAddress, port uint16, mode BindMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHostAddress(address)

	slotval2 := (uint16)(port)

	xxxxxxxxx

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_Bind, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_ConnectToHost(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))

	QAbstractSocket_virtualbase_ConnectToHost(unsafe.Pointer(this.h), hostName_ms, (uint16_t)(port), mode, protocol)

}
func (this *QAbstractSocket) OnConnectToHost(slot func(super func(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol), hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_ConnectToHost(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_ConnectToHost
func miqt_exec_callback_QAbstractSocket_ConnectToHost(self QAbstractSocket, cb intptr_t, hostName struct_miqt_string, port uint16_t, mode OpenMode, protocol NetworkLayerProtocol) {
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

	gofunc((&QAbstractSocket{h: self}).callVirtualBase_ConnectToHost, slotval1, slotval2, slotval3, slotval4)

}

func (this *QAbstractSocket) callVirtualBase_DisconnectFromHost() {

	QAbstractSocket_virtualbase_DisconnectFromHost(unsafe.Pointer(this.h))

}
func (this *QAbstractSocket) OnDisconnectFromHost(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_DisconnectFromHost(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_DisconnectFromHost
func miqt_exec_callback_QAbstractSocket_DisconnectFromHost(self QAbstractSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractSocket{h: self}).callVirtualBase_DisconnectFromHost)

}

func (this *QAbstractSocket) callVirtualBase_BytesAvailable() int64 {

	return (int64)(QAbstractSocket_virtualbase_BytesAvailable(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnBytesAvailable(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_BytesAvailable(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_BytesAvailable
func miqt_exec_callback_QAbstractSocket_BytesAvailable(self QAbstractSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_BytesAvailable)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_BytesToWrite() int64 {

	return (int64)(QAbstractSocket_virtualbase_BytesToWrite(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnBytesToWrite(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_BytesToWrite(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_BytesToWrite
func miqt_exec_callback_QAbstractSocket_BytesToWrite(self QAbstractSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_BytesToWrite)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_SetReadBufferSize(size int64) {

	QAbstractSocket_virtualbase_SetReadBufferSize(unsafe.Pointer(this.h), (longlong)(size))

}
func (this *QAbstractSocket) OnSetReadBufferSize(slot func(super func(size int64), size int64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_SetReadBufferSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_SetReadBufferSize
func miqt_exec_callback_QAbstractSocket_SetReadBufferSize(self QAbstractSocket, cb intptr_t, size longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size int64), size int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(size)

	gofunc((&QAbstractSocket{h: self}).callVirtualBase_SetReadBufferSize, slotval1)

}

func (this *QAbstractSocket) callVirtualBase_SocketDescriptor() uintptr {

	return (uintptr)(QAbstractSocket_virtualbase_SocketDescriptor(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnSocketDescriptor(slot func(super func() uintptr) uintptr) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_SocketDescriptor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_SocketDescriptor
func miqt_exec_callback_QAbstractSocket_SocketDescriptor(self QAbstractSocket, cb intptr_t) intptr_t {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() uintptr) uintptr)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_SocketDescriptor)

	return (intptr_t)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_SetSocketDescriptor(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool {

	return (bool)(QAbstractSocket_virtualbase_SetSocketDescriptor(unsafe.Pointer(this.h), (intptr_t)(socketDescriptor), state, openMode))

}
func (this *QAbstractSocket) OnSetSocketDescriptor(slot func(super func(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool, socketDescriptor uintptr, state SocketState, openMode OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_SetSocketDescriptor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_SetSocketDescriptor
func miqt_exec_callback_QAbstractSocket_SetSocketDescriptor(self QAbstractSocket, cb intptr_t, socketDescriptor intptr_t, state SocketState, openMode OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool, socketDescriptor uintptr, state SocketState, openMode OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uintptr)(socketDescriptor)

	xxxxxxxxx
	xxxxxxxxx

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_SetSocketDescriptor, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_SetSocketOption(option QAbstractSocket__SocketOption, value *qt.QVariant) {

	QAbstractSocket_virtualbase_SetSocketOption(unsafe.Pointer(this.h), (int)(option), (*QVariant)(value.UnsafePointer()))

}
func (this *QAbstractSocket) OnSetSocketOption(slot func(super func(option QAbstractSocket__SocketOption, value *qt.QVariant), option QAbstractSocket__SocketOption, value *qt.QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_SetSocketOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_SetSocketOption
func miqt_exec_callback_QAbstractSocket_SetSocketOption(self QAbstractSocket, cb intptr_t, option int, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option QAbstractSocket__SocketOption, value *qt.QVariant), option QAbstractSocket__SocketOption, value *qt.QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketOption)(option)

	slotval2 := qt.UnsafeNewQVariant(unsafe.Pointer(value))

	gofunc((&QAbstractSocket{h: self}).callVirtualBase_SetSocketOption, slotval1, slotval2)

}

func (this *QAbstractSocket) callVirtualBase_SocketOption(option QAbstractSocket__SocketOption) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QAbstractSocket_virtualbase_SocketOption(unsafe.Pointer(this.h), (int)(option))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractSocket) OnSocketOption(slot func(super func(option QAbstractSocket__SocketOption) *qt.QVariant, option QAbstractSocket__SocketOption) *qt.QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_SocketOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_SocketOption
func miqt_exec_callback_QAbstractSocket_SocketOption(self QAbstractSocket, cb intptr_t, option int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option QAbstractSocket__SocketOption) *qt.QVariant, option QAbstractSocket__SocketOption) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractSocket__SocketOption)(option)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_SocketOption, slotval1)

	return (*QVariant)(virtualReturn.UnsafePointer())

}

func (this *QAbstractSocket) callVirtualBase_Close() {

	QAbstractSocket_virtualbase_Close(unsafe.Pointer(this.h))

}
func (this *QAbstractSocket) OnClose(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Close(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Close
func miqt_exec_callback_QAbstractSocket_Close(self QAbstractSocket, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractSocket{h: self}).callVirtualBase_Close)

}

func (this *QAbstractSocket) callVirtualBase_IsSequential() bool {

	return (bool)(QAbstractSocket_virtualbase_IsSequential(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnIsSequential(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_IsSequential
func miqt_exec_callback_QAbstractSocket_IsSequential(self QAbstractSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_IsSequential)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_WaitForConnected(msecs int) bool {

	return (bool)(QAbstractSocket_virtualbase_WaitForConnected(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QAbstractSocket) OnWaitForConnected(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_WaitForConnected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_WaitForConnected
func miqt_exec_callback_QAbstractSocket_WaitForConnected(self QAbstractSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_WaitForConnected, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_WaitForReadyRead(msecs int) bool {

	return (bool)(QAbstractSocket_virtualbase_WaitForReadyRead(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QAbstractSocket) OnWaitForReadyRead(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_WaitForReadyRead(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_WaitForReadyRead
func miqt_exec_callback_QAbstractSocket_WaitForReadyRead(self QAbstractSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_WaitForReadyRead, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_WaitForBytesWritten(msecs int) bool {

	return (bool)(QAbstractSocket_virtualbase_WaitForBytesWritten(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QAbstractSocket) OnWaitForBytesWritten(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_WaitForBytesWritten(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_WaitForBytesWritten
func miqt_exec_callback_QAbstractSocket_WaitForBytesWritten(self QAbstractSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_WaitForBytesWritten, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_WaitForDisconnected(msecs int) bool {

	return (bool)(QAbstractSocket_virtualbase_WaitForDisconnected(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QAbstractSocket) OnWaitForDisconnected(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_WaitForDisconnected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_WaitForDisconnected
func miqt_exec_callback_QAbstractSocket_WaitForDisconnected(self QAbstractSocket, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_WaitForDisconnected, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_ReadData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QAbstractSocket_virtualbase_ReadData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QAbstractSocket) OnReadData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_ReadData
func miqt_exec_callback_QAbstractSocket_ReadData(self QAbstractSocket, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_ReadData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_ReadLineData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QAbstractSocket_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QAbstractSocket) OnReadLineData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_ReadLineData
func miqt_exec_callback_QAbstractSocket_ReadLineData(self QAbstractSocket, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_SkipData(maxSize int64) int64 {

	return (int64)(QAbstractSocket_virtualbase_SkipData(unsafe.Pointer(this.h), (longlong)(maxSize)))

}
func (this *QAbstractSocket) OnSkipData(slot func(super func(maxSize int64) int64, maxSize int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_SkipData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_SkipData
func miqt_exec_callback_QAbstractSocket_SkipData(self QAbstractSocket, cb intptr_t, maxSize longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(maxSize int64) int64, maxSize int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(maxSize)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_SkipData, slotval1)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_WriteData(data string, lenVal int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QAbstractSocket_virtualbase_WriteData(unsafe.Pointer(this.h), data_Cstring, (longlong)(lenVal)))

}
func (this *QAbstractSocket) OnWriteData(slot func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_WriteData
func miqt_exec_callback_QAbstractSocket_WriteData(self QAbstractSocket, cb intptr_t, data *const_char, lenVal longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(lenVal)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_WriteData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_Open(mode OpenModeFlag) bool {

	return (bool)(QAbstractSocket_virtualbase_Open(unsafe.Pointer(this.h), (int)(mode)))

}
func (this *QAbstractSocket) OnOpen(slot func(super func(mode OpenModeFlag) bool, mode OpenModeFlag) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Open
func miqt_exec_callback_QAbstractSocket_Open(self QAbstractSocket, cb intptr_t, mode int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mode OpenModeFlag) bool, mode OpenModeFlag) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (OpenModeFlag)(mode)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_Open, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_Pos() int64 {

	return (int64)(QAbstractSocket_virtualbase_Pos(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnPos(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Pos(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Pos
func miqt_exec_callback_QAbstractSocket_Pos(self QAbstractSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_Pos)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_Size() int64 {

	return (int64)(QAbstractSocket_virtualbase_Size(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnSize(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Size
func miqt_exec_callback_QAbstractSocket_Size(self QAbstractSocket, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_Size)

	return (longlong)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_Seek(pos int64) bool {

	return (bool)(QAbstractSocket_virtualbase_Seek(unsafe.Pointer(this.h), (longlong)(pos)))

}
func (this *QAbstractSocket) OnSeek(slot func(super func(pos int64) bool, pos int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Seek(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Seek
func miqt_exec_callback_QAbstractSocket_Seek(self QAbstractSocket, cb intptr_t, pos longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos int64) bool, pos int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(pos)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_Seek, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_AtEnd() bool {

	return (bool)(QAbstractSocket_virtualbase_AtEnd(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnAtEnd(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_AtEnd(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_AtEnd
func miqt_exec_callback_QAbstractSocket_AtEnd(self QAbstractSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_AtEnd)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_Reset() bool {

	return (bool)(QAbstractSocket_virtualbase_Reset(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnReset(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Reset
func miqt_exec_callback_QAbstractSocket_Reset(self QAbstractSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_Reset)

	return (bool)(virtualReturn)

}

func (this *QAbstractSocket) callVirtualBase_CanReadLine() bool {

	return (bool)(QAbstractSocket_virtualbase_CanReadLine(unsafe.Pointer(this.h)))

}
func (this *QAbstractSocket) OnCanReadLine(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_CanReadLine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_CanReadLine
func miqt_exec_callback_QAbstractSocket_CanReadLine(self QAbstractSocket, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_CanReadLine)

	return (bool)(virtualReturn)

}
