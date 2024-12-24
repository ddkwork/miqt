package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSslSocket__SslMode int

const (
	QSslSocket__UnencryptedMode QSslSocket__SslMode = 0
	QSslSocket__SslClientMode   QSslSocket__SslMode = 1
	QSslSocket__SslServerMode   QSslSocket__SslMode = 2
)

type QSslSocket__PeerVerifyMode int

const (
	QSslSocket__VerifyNone     QSslSocket__PeerVerifyMode = 0
	QSslSocket__QueryPeer      QSslSocket__PeerVerifyMode = 1
	QSslSocket__VerifyPeer     QSslSocket__PeerVerifyMode = 2
	QSslSocket__AutoVerifyPeer QSslSocket__PeerVerifyMode = 3
)

type QSslSocket struct {
	h          uintptr
	isSubclass bool
}

// NewQSslSocket constructs a new QSslSocket object.
func NewQSslSocket() *QSslSocket {
	ret := newQSslSocket(QSslSocket_new())
	ret.isSubclass = true
	return ret
}

// NewQSslSocket2 constructs a new QSslSocket object.
func NewQSslSocket2(parent *qt.QObject) *QSslSocket {
	ret := newQSslSocket(QSslSocket_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QSslSocket) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSslSocket_MetaObject(this.h)))
}

func (this *QSslSocket) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSslSocket_Metacast(this.h, param1_Cstring))
}

func QSslSocket_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSslSocket_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslSocket) Resume() {
	QSslSocket_Resume(this.h)
}

func (this *QSslSocket) ConnectToHostEncrypted(hostName string, port uint16) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QSslSocket_ConnectToHostEncrypted(this.h, hostName_ms, (uint16_t)(port))
}

func (this *QSslSocket) ConnectToHostEncrypted2(hostName string, port uint16, sslPeerName string) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	sslPeerName_ms := struct_miqt_string{}
	sslPeerName_ms.data = CString(sslPeerName)
	sslPeerName_ms.len = size_t(len(sslPeerName))
	defer free(unsafe.Pointer(sslPeerName_ms.data))
	QSslSocket_ConnectToHostEncrypted2(this.h, hostName_ms, (uint16_t)(port), sslPeerName_ms)
}

func (this *QSslSocket) SetSocketDescriptor(socketDescriptor uintptr, state SocketState, openMode OpenMode) bool {
	return (bool)(QSslSocket_SetSocketDescriptor(this.h, (intptr_t)(socketDescriptor), state, openMode))
}

func (this *QSslSocket) ConnectToHost(hostName string, port uint16, openMode OpenMode, protocol NetworkLayerProtocol) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QSslSocket_ConnectToHost(this.h, hostName_ms, (uint16_t)(port), openMode, protocol)
}

func (this *QSslSocket) DisconnectFromHost() {
	QSslSocket_DisconnectFromHost(this.h)
}

func (this *QSslSocket) SetSocketOption(option QAbstractSocket__SocketOption, value *qt.QVariant) {
	QSslSocket_SetSocketOption(this.h, (int)(option), (*QVariant)(value.UnsafePointer()))
}

func (this *QSslSocket) SocketOption(option QAbstractSocket__SocketOption) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QSslSocket_SocketOption(this.h, (int)(option))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslSocket) Mode() SslMode {
	xxxxxxxxx
}

func (this *QSslSocket) IsEncrypted() bool {
	return (bool)(QSslSocket_IsEncrypted(this.h))
}

func (this *QSslSocket) Protocol() QSsl__SslProtocol {
	return (QSsl__SslProtocol)(QSslSocket_Protocol(this.h))
}

func (this *QSslSocket) SetProtocol(protocol QSsl__SslProtocol) {
	QSslSocket_SetProtocol(this.h, (int)(protocol))
}

func (this *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	return (QSslSocket__PeerVerifyMode)(QSslSocket_PeerVerifyMode(this.h))
}

func (this *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	QSslSocket_SetPeerVerifyMode(this.h, (int)(mode))
}

func (this *QSslSocket) PeerVerifyDepth() int {
	return (int)(QSslSocket_PeerVerifyDepth(this.h))
}

func (this *QSslSocket) SetPeerVerifyDepth(depth int) {
	QSslSocket_SetPeerVerifyDepth(this.h, (int)(depth))
}

func (this *QSslSocket) PeerVerifyName() string {
	var _ms struct_miqt_string = QSslSocket_PeerVerifyName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslSocket) SetPeerVerifyName(hostName string) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QSslSocket_SetPeerVerifyName(this.h, hostName_ms)
}

func (this *QSslSocket) BytesAvailable() int64 {
	return (int64)(QSslSocket_BytesAvailable(this.h))
}

func (this *QSslSocket) BytesToWrite() int64 {
	return (int64)(QSslSocket_BytesToWrite(this.h))
}

func (this *QSslSocket) CanReadLine() bool {
	return (bool)(QSslSocket_CanReadLine(this.h))
}

func (this *QSslSocket) Close() {
	QSslSocket_Close(this.h)
}

func (this *QSslSocket) AtEnd() bool {
	return (bool)(QSslSocket_AtEnd(this.h))
}

func (this *QSslSocket) SetReadBufferSize(size int64) {
	QSslSocket_SetReadBufferSize(this.h, (longlong)(size))
}

func (this *QSslSocket) EncryptedBytesAvailable() int64 {
	return (int64)(QSslSocket_EncryptedBytesAvailable(this.h))
}

func (this *QSslSocket) EncryptedBytesToWrite() int64 {
	return (int64)(QSslSocket_EncryptedBytesToWrite(this.h))
}

func (this *QSslSocket) SslConfiguration() *QSslConfiguration {
	_goptr := newQSslConfiguration(QSslSocket_SslConfiguration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslSocket) SetSslConfiguration(config *QSslConfiguration) {
	QSslSocket_SetSslConfiguration(this.h, config.cPointer())
}

func (this *QSslSocket) SetLocalCertificateChain(localChain []QSslCertificate) {
	localChain_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(localChain))))
	defer free(unsafe.Pointer(localChain_CArray))
	for i := range localChain {
		localChain_CArray[i] = localChain[i].cPointer()
	}
	localChain_ma := struct_miqt_array{len: size_t(len(localChain)), data: unsafe.Pointer(localChain_CArray)}
	QSslSocket_SetLocalCertificateChain(this.h, localChain_ma)
}

func (this *QSslSocket) LocalCertificateChain() []QSslCertificate {
	var _ma struct_miqt_array = QSslSocket_LocalCertificateChain(this.h)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslSocket) SetLocalCertificate(certificate *QSslCertificate) {
	QSslSocket_SetLocalCertificate(this.h, certificate.cPointer())
}

func (this *QSslSocket) SetLocalCertificateWithFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QSslSocket_SetLocalCertificateWithFileName(this.h, fileName_ms)
}

func (this *QSslSocket) LocalCertificate() *QSslCertificate {
	_goptr := newQSslCertificate(QSslSocket_LocalCertificate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslSocket) PeerCertificate() *QSslCertificate {
	_goptr := newQSslCertificate(QSslSocket_PeerCertificate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslSocket) PeerCertificateChain() []QSslCertificate {
	var _ma struct_miqt_array = QSslSocket_PeerCertificateChain(this.h)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslSocket) SessionCipher() *QSslCipher {
	_goptr := newQSslCipher(QSslSocket_SessionCipher(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslSocket) SessionProtocol() QSsl__SslProtocol {
	return (QSsl__SslProtocol)(QSslSocket_SessionProtocol(this.h))
}

func (this *QSslSocket) OcspResponses() []QOcspResponse {
	var _ma struct_miqt_array = QSslSocket_OcspResponses(this.h)
	_ret := make([]QOcspResponse, int(_ma.len))
	_outCast := (*[0xffff]*QOcspResponse)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQOcspResponse(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslSocket) SetPrivateKey(key *QSslKey) {
	QSslSocket_SetPrivateKey(this.h, key.cPointer())
}

func (this *QSslSocket) SetPrivateKeyWithFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QSslSocket_SetPrivateKeyWithFileName(this.h, fileName_ms)
}

func (this *QSslSocket) PrivateKey() *QSslKey {
	_goptr := newQSslKey(QSslSocket_PrivateKey(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslSocket) WaitForConnected(msecs int) bool {
	return (bool)(QSslSocket_WaitForConnected(this.h, (int)(msecs)))
}

func (this *QSslSocket) WaitForEncrypted() bool {
	return (bool)(QSslSocket_WaitForEncrypted(this.h))
}

func (this *QSslSocket) WaitForReadyRead(msecs int) bool {
	return (bool)(QSslSocket_WaitForReadyRead(this.h, (int)(msecs)))
}

func (this *QSslSocket) WaitForBytesWritten(msecs int) bool {
	return (bool)(QSslSocket_WaitForBytesWritten(this.h, (int)(msecs)))
}

func (this *QSslSocket) WaitForDisconnected(msecs int) bool {
	return (bool)(QSslSocket_WaitForDisconnected(this.h, (int)(msecs)))
}

func (this *QSslSocket) SslHandshakeErrors() []QSslError {
	var _ma struct_miqt_array = QSslSocket_SslHandshakeErrors(this.h)
	_ret := make([]QSslError, int(_ma.len))
	_outCast := (*[0xffff]*QSslError)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslError(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslSocket_SupportsSsl() bool {
	return (bool)(QSslSocket_SupportsSsl())
}

func QSslSocket_SslLibraryVersionNumber() int32 {
	return (int32)(QSslSocket_SslLibraryVersionNumber())
}

func QSslSocket_SslLibraryVersionString() string {
	var _ms struct_miqt_string = QSslSocket_SslLibraryVersionString()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSslSocket_SslLibraryBuildVersionNumber() int32 {
	return (int32)(QSslSocket_SslLibraryBuildVersionNumber())
}

func QSslSocket_SslLibraryBuildVersionString() string {
	var _ms struct_miqt_string = QSslSocket_SslLibraryBuildVersionString()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSslSocket_AvailableBackends() []string {
	var _ma struct_miqt_array = QSslSocket_AvailableBackends()
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QSslSocket_ActiveBackend() string {
	var _ms struct_miqt_string = QSslSocket_ActiveBackend()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSslSocket_SetActiveBackend(backendName string) bool {
	backendName_ms := struct_miqt_string{}
	backendName_ms.data = CString(backendName)
	backendName_ms.len = size_t(len(backendName))
	defer free(unsafe.Pointer(backendName_ms.data))
	return (bool)(QSslSocket_SetActiveBackend(backendName_ms))
}

func QSslSocket_SupportedProtocols() []QSsl__SslProtocol {
	var _ma struct_miqt_array = QSslSocket_SupportedProtocols()
	_ret := make([]QSsl__SslProtocol, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QSsl__SslProtocol)(_outCast[i])
	}
	return _ret
}

func QSslSocket_IsProtocolSupported(protocol QSsl__SslProtocol) bool {
	return (bool)(QSslSocket_IsProtocolSupported((int)(protocol)))
}

func QSslSocket_ImplementedClasses() []QSsl__ImplementedClass {
	var _ma struct_miqt_array = QSslSocket_ImplementedClasses()
	_ret := make([]QSsl__ImplementedClass, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QSsl__ImplementedClass)(_outCast[i])
	}
	return _ret
}

func QSslSocket_IsClassImplemented(cl QSsl__ImplementedClass) bool {
	return (bool)(QSslSocket_IsClassImplemented((int)(cl)))
}

func QSslSocket_SupportedFeatures() []QSsl__SupportedFeature {
	var _ma struct_miqt_array = QSslSocket_SupportedFeatures()
	_ret := make([]QSsl__SupportedFeature, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QSsl__SupportedFeature)(_outCast[i])
	}
	return _ret
}

func QSslSocket_IsFeatureSupported(feat QSsl__SupportedFeature) bool {
	return (bool)(QSslSocket_IsFeatureSupported((int)(feat)))
}

func (this *QSslSocket) IgnoreSslErrors(errors []QSslError) {
	errors_CArray := (*[0xffff]*QSslError)(malloc(size_t(8 * len(errors))))
	defer free(unsafe.Pointer(errors_CArray))
	for i := range errors {
		errors_CArray[i] = errors[i].cPointer()
	}
	errors_ma := struct_miqt_array{len: size_t(len(errors)), data: unsafe.Pointer(errors_CArray)}
	QSslSocket_IgnoreSslErrors(this.h, errors_ma)
}

func (this *QSslSocket) ContinueInterruptedHandshake() {
	QSslSocket_ContinueInterruptedHandshake(this.h)
}

func (this *QSslSocket) StartClientEncryption() {
	QSslSocket_StartClientEncryption(this.h)
}

func (this *QSslSocket) StartServerEncryption() {
	QSslSocket_StartServerEncryption(this.h)
}

func (this *QSslSocket) IgnoreSslErrors2() {
	QSslSocket_IgnoreSslErrors2(this.h)
}

func (this *QSslSocket) Encrypted() {
	QSslSocket_Encrypted(this.h)
}

func (this *QSslSocket) OnEncrypted(slot func()) {
	QSslSocket_connect_Encrypted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_Encrypted
func miqt_exec_callback_QSslSocket_Encrypted(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSslSocket) PeerVerifyError(error *QSslError) {
	QSslSocket_PeerVerifyError(this.h, error.cPointer())
}

func (this *QSslSocket) OnPeerVerifyError(slot func(error *QSslError)) {
	QSslSocket_connect_PeerVerifyError(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_PeerVerifyError
func miqt_exec_callback_QSslSocket_PeerVerifyError(cb intptr_t, error *QSslError) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error *QSslError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslError(error)

	gofunc(slotval1)
}

func (this *QSslSocket) SslErrors(errors []QSslError) {
	errors_CArray := (*[0xffff]*QSslError)(malloc(size_t(8 * len(errors))))
	defer free(unsafe.Pointer(errors_CArray))
	for i := range errors {
		errors_CArray[i] = errors[i].cPointer()
	}
	errors_ma := struct_miqt_array{len: size_t(len(errors)), data: unsafe.Pointer(errors_CArray)}
	QSslSocket_SslErrors(this.h, errors_ma)
}

func (this *QSslSocket) OnSslErrors(slot func(errors []QSslError)) {
	QSslSocket_connect_SslErrors(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_SslErrors
func miqt_exec_callback_QSslSocket_SslErrors(cb intptr_t, errors struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(errors []QSslError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var errors_ma struct_miqt_array = errors
	errors_ret := make([]QSslError, int(errors_ma.len))
	errors_outCast := (*[0xffff]*QSslError)(unsafe.Pointer(errors_ma.data)) // hey ya
	for i := 0; i < int(errors_ma.len); i++ {
		errors_lv_goptr := newQSslError(errors_outCast[i])
		errors_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		errors_ret[i] = *errors_lv_goptr
	}
	slotval1 := errors_ret

	gofunc(slotval1)
}

func (this *QSslSocket) ModeChanged(newMode QSslSocket__SslMode) {
	QSslSocket_ModeChanged(this.h, (int)(newMode))
}

func (this *QSslSocket) OnModeChanged(slot func(newMode QSslSocket__SslMode)) {
	QSslSocket_connect_ModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_ModeChanged
func miqt_exec_callback_QSslSocket_ModeChanged(cb intptr_t, newMode int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newMode QSslSocket__SslMode))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QSslSocket__SslMode)(newMode)

	gofunc(slotval1)
}

func (this *QSslSocket) EncryptedBytesWritten(totalBytes int64) {
	QSslSocket_EncryptedBytesWritten(this.h, (longlong)(totalBytes))
}

func (this *QSslSocket) OnEncryptedBytesWritten(slot func(totalBytes int64)) {
	QSslSocket_connect_EncryptedBytesWritten(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_EncryptedBytesWritten
func miqt_exec_callback_QSslSocket_EncryptedBytesWritten(cb intptr_t, totalBytes longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(totalBytes int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(totalBytes)

	gofunc(slotval1)
}

func (this *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator *QSslPreSharedKeyAuthenticator) {
	QSslSocket_PreSharedKeyAuthenticationRequired(this.h, authenticator.cPointer())
}

func (this *QSslSocket) OnPreSharedKeyAuthenticationRequired(slot func(authenticator *QSslPreSharedKeyAuthenticator)) {
	QSslSocket_connect_PreSharedKeyAuthenticationRequired(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_PreSharedKeyAuthenticationRequired
func miqt_exec_callback_QSslSocket_PreSharedKeyAuthenticationRequired(cb intptr_t, authenticator *QSslPreSharedKeyAuthenticator) {
	gofunc, ok := cgo.Handle(cb).Value().(func(authenticator *QSslPreSharedKeyAuthenticator))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslPreSharedKeyAuthenticator(authenticator)

	gofunc(slotval1)
}

func (this *QSslSocket) NewSessionTicketReceived() {
	QSslSocket_NewSessionTicketReceived(this.h)
}

func (this *QSslSocket) OnNewSessionTicketReceived(slot func()) {
	QSslSocket_connect_NewSessionTicketReceived(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_NewSessionTicketReceived
func miqt_exec_callback_QSslSocket_NewSessionTicketReceived(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSslSocket) AlertSent(level QSsl__AlertLevel, typeVal QSsl__AlertType, description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QSslSocket_AlertSent(this.h, (int)(level), (int)(typeVal), description_ms)
}

func (this *QSslSocket) OnAlertSent(slot func(level QSsl__AlertLevel, typeVal QSsl__AlertType, description string)) {
	QSslSocket_connect_AlertSent(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_AlertSent
func miqt_exec_callback_QSslSocket_AlertSent(cb intptr_t, level int, typeVal int, description struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(level QSsl__AlertLevel, typeVal QSsl__AlertType, description string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QSsl__AlertLevel)(level)

	slotval2 := (QSsl__AlertType)(typeVal)

	var description_ms struct_miqt_string = description
	description_ret := GoStringN(description_ms.data, int(int64(description_ms.len)))
	free(unsafe.Pointer(description_ms.data))
	slotval3 := description_ret

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QSslSocket) AlertReceived(level QSsl__AlertLevel, typeVal QSsl__AlertType, description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QSslSocket_AlertReceived(this.h, (int)(level), (int)(typeVal), description_ms)
}

func (this *QSslSocket) OnAlertReceived(slot func(level QSsl__AlertLevel, typeVal QSsl__AlertType, description string)) {
	QSslSocket_connect_AlertReceived(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_AlertReceived
func miqt_exec_callback_QSslSocket_AlertReceived(cb intptr_t, level int, typeVal int, description struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(level QSsl__AlertLevel, typeVal QSsl__AlertType, description string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QSsl__AlertLevel)(level)

	slotval2 := (QSsl__AlertType)(typeVal)

	var description_ms struct_miqt_string = description
	description_ret := GoStringN(description_ms.data, int(int64(description_ms.len)))
	free(unsafe.Pointer(description_ms.data))
	slotval3 := description_ret

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QSslSocket) HandshakeInterruptedOnError(error *QSslError) {
	QSslSocket_HandshakeInterruptedOnError(this.h, error.cPointer())
}

func (this *QSslSocket) OnHandshakeInterruptedOnError(slot func(error *QSslError)) {
	QSslSocket_connect_HandshakeInterruptedOnError(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_HandshakeInterruptedOnError
func miqt_exec_callback_QSslSocket_HandshakeInterruptedOnError(cb intptr_t, error *QSslError) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error *QSslError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslError(error)

	gofunc(slotval1)
}

func QSslSocket_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSslSocket_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSslSocket_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSslSocket_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslSocket) ConnectToHostEncrypted3(hostName string, port uint16, mode OpenMode) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QSslSocket_ConnectToHostEncrypted3(this.h, hostName_ms, (uint16_t)(port), mode)
}

func (this *QSslSocket) ConnectToHostEncrypted4(hostName string, port uint16, mode OpenMode, protocol NetworkLayerProtocol) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QSslSocket_ConnectToHostEncrypted4(this.h, hostName_ms, (uint16_t)(port), mode, protocol)
}

func (this *QSslSocket) ConnectToHostEncrypted42(hostName string, port uint16, sslPeerName string, mode OpenMode) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	sslPeerName_ms := struct_miqt_string{}
	sslPeerName_ms.data = CString(sslPeerName)
	sslPeerName_ms.len = size_t(len(sslPeerName))
	defer free(unsafe.Pointer(sslPeerName_ms.data))
	QSslSocket_ConnectToHostEncrypted42(this.h, hostName_ms, (uint16_t)(port), sslPeerName_ms, mode)
}

func (this *QSslSocket) ConnectToHostEncrypted5(hostName string, port uint16, sslPeerName string, mode OpenMode, protocol NetworkLayerProtocol) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	sslPeerName_ms := struct_miqt_string{}
	sslPeerName_ms.data = CString(sslPeerName)
	sslPeerName_ms.len = size_t(len(sslPeerName))
	defer free(unsafe.Pointer(sslPeerName_ms.data))
	QSslSocket_ConnectToHostEncrypted5(this.h, hostName_ms, (uint16_t)(port), sslPeerName_ms, mode, protocol)
}

func (this *QSslSocket) SetLocalCertificate2(fileName string, format QSsl__EncodingFormat) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QSslSocket_SetLocalCertificate2(this.h, fileName_ms, (int)(format))
}

func (this *QSslSocket) SetPrivateKey2(fileName string, algorithm QSsl__KeyAlgorithm) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QSslSocket_SetPrivateKey2(this.h, fileName_ms, (int)(algorithm))
}

func (this *QSslSocket) SetPrivateKey3(fileName string, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QSslSocket_SetPrivateKey3(this.h, fileName_ms, (int)(algorithm), (int)(format))
}

func (this *QSslSocket) SetPrivateKey4(fileName string, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, passPhrase []byte) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	passPhrase_alias := struct_miqt_string{}
	passPhrase_alias.data = (char)(unsafe.Pointer(&passPhrase[0]))
	passPhrase_alias.len = size_t(len(passPhrase))
	QSslSocket_SetPrivateKey4(this.h, fileName_ms, (int)(algorithm), (int)(format), passPhrase_alias)
}

func (this *QSslSocket) WaitForEncrypted1(msecs int) bool {
	return (bool)(QSslSocket_WaitForEncrypted1(this.h, (int)(msecs)))
}

func QSslSocket_SupportedProtocols1(backendName string) []QSsl__SslProtocol {
	backendName_ms := struct_miqt_string{}
	backendName_ms.data = CString(backendName)
	backendName_ms.len = size_t(len(backendName))
	defer free(unsafe.Pointer(backendName_ms.data))
	var _ma struct_miqt_array = QSslSocket_SupportedProtocols1(backendName_ms)
	_ret := make([]QSsl__SslProtocol, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QSsl__SslProtocol)(_outCast[i])
	}
	return _ret
}

func QSslSocket_IsProtocolSupported2(protocol QSsl__SslProtocol, backendName string) bool {
	backendName_ms := struct_miqt_string{}
	backendName_ms.data = CString(backendName)
	backendName_ms.len = size_t(len(backendName))
	defer free(unsafe.Pointer(backendName_ms.data))
	return (bool)(QSslSocket_IsProtocolSupported2((int)(protocol), backendName_ms))
}

func QSslSocket_ImplementedClasses1(backendName string) []QSsl__ImplementedClass {
	backendName_ms := struct_miqt_string{}
	backendName_ms.data = CString(backendName)
	backendName_ms.len = size_t(len(backendName))
	defer free(unsafe.Pointer(backendName_ms.data))
	var _ma struct_miqt_array = QSslSocket_ImplementedClasses1(backendName_ms)
	_ret := make([]QSsl__ImplementedClass, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QSsl__ImplementedClass)(_outCast[i])
	}
	return _ret
}

func QSslSocket_IsClassImplemented2(cl QSsl__ImplementedClass, backendName string) bool {
	backendName_ms := struct_miqt_string{}
	backendName_ms.data = CString(backendName)
	backendName_ms.len = size_t(len(backendName))
	defer free(unsafe.Pointer(backendName_ms.data))
	return (bool)(QSslSocket_IsClassImplemented2((int)(cl), backendName_ms))
}

func QSslSocket_SupportedFeatures1(backendName string) []QSsl__SupportedFeature {
	backendName_ms := struct_miqt_string{}
	backendName_ms.data = CString(backendName)
	backendName_ms.len = size_t(len(backendName))
	defer free(unsafe.Pointer(backendName_ms.data))
	var _ma struct_miqt_array = QSslSocket_SupportedFeatures1(backendName_ms)
	_ret := make([]QSsl__SupportedFeature, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QSsl__SupportedFeature)(_outCast[i])
	}
	return _ret
}

func QSslSocket_IsFeatureSupported2(feat QSsl__SupportedFeature, backendName string) bool {
	backendName_ms := struct_miqt_string{}
	backendName_ms.data = CString(backendName)
	backendName_ms.len = size_t(len(backendName))
	defer free(unsafe.Pointer(backendName_ms.data))
	return (bool)(QSslSocket_IsFeatureSupported2((int)(feat), backendName_ms))
}

func (this *QSslSocket) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSslSocket_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QSslSocket) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSslSocket_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_MetaObject
func miqt_exec_callback_QSslSocket_MetaObject(self QSslSocket, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSslSocket{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QSslSocket) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSslSocket_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSslSocket) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSslSocket_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslSocket_Metacast
func miqt_exec_callback_QSslSocket_Metacast(self QSslSocket, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QSslSocket{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
