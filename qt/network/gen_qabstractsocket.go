package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
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

func (this *QAbstractSocket) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAbstractSocket_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAbstractSocket) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_MetaObject
func miqt_exec_callback_QAbstractSocket_MetaObject(self QAbstractSocket, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAbstractSocket) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractSocket_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractSocket) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSocket_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSocket_Metacast
func miqt_exec_callback_QAbstractSocket_Metacast(self QAbstractSocket, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractSocket{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
