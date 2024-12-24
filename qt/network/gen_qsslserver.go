package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QSslServer struct {
	h          uintptr
	isSubclass bool
}

// NewQSslServer constructs a new QSslServer object.
func NewQSslServer() *QSslServer {

	ret := newQSslServer(QSslServer_new())
	ret.isSubclass = true
	return ret
}

// NewQSslServer2 constructs a new QSslServer object.
func NewQSslServer2(parent *qt.QObject) *QSslServer {

	ret := newQSslServer(QSslServer_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QSslServer) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSslServer_MetaObject(this.h)))
}

func (this *QSslServer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSslServer_Metacast(this.h, param1_Cstring))
}

func QSslServer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSslServer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslServer) SetSslConfiguration(sslConfiguration *QSslConfiguration) {
	QSslServer_SetSslConfiguration(this.h, sslConfiguration.cPointer())
}

func (this *QSslServer) SslConfiguration() *QSslConfiguration {
	_goptr := newQSslConfiguration(QSslServer_SslConfiguration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslServer) SetHandshakeTimeout(timeout int) {
	QSslServer_SetHandshakeTimeout(this.h, (int)(timeout))
}

func (this *QSslServer) HandshakeTimeout() int {
	return (int)(QSslServer_HandshakeTimeout(this.h))
}

func (this *QSslServer) SslErrors(socket *QSslSocket, errors []QSslError) {
	errors_CArray := (*[0xffff]*QSslError)(malloc(size_t(8 * len(errors))))
	defer free(unsafe.Pointer(errors_CArray))
	for i := range errors {
		errors_CArray[i] = errors[i].cPointer()
	}
	errors_ma := struct_miqt_array{len: size_t(len(errors)), data: unsafe.Pointer(errors_CArray)}
	QSslServer_SslErrors(this.h, socket.cPointer(), errors_ma)
}
func (this *QSslServer) OnSslErrors(slot func(socket *QSslSocket, errors []QSslError)) {
	QSslServer_connect_SslErrors(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_SslErrors
func miqt_exec_callback_QSslServer_SslErrors(cb intptr_t, socket *QSslSocket, errors struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket, errors []QSslError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	var errors_ma struct_miqt_array = errors
	errors_ret := make([]QSslError, int(errors_ma.len))
	errors_outCast := (*[0xffff]*QSslError)(unsafe.Pointer(errors_ma.data)) // hey ya
	for i := 0; i < int(errors_ma.len); i++ {
		errors_lv_goptr := newQSslError(errors_outCast[i])
		errors_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		errors_ret[i] = *errors_lv_goptr
	}
	slotval2 := errors_ret

	gofunc(slotval1, slotval2)
}

func (this *QSslServer) PeerVerifyError(socket *QSslSocket, error *QSslError) {
	QSslServer_PeerVerifyError(this.h, socket.cPointer(), error.cPointer())
}
func (this *QSslServer) OnPeerVerifyError(slot func(socket *QSslSocket, error *QSslError)) {
	QSslServer_connect_PeerVerifyError(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_PeerVerifyError
func miqt_exec_callback_QSslServer_PeerVerifyError(cb intptr_t, socket *QSslSocket, error *QSslError) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket, error *QSslError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	slotval2 := newQSslError(error)

	gofunc(slotval1, slotval2)
}

func (this *QSslServer) ErrorOccurred(socket *QSslSocket, error QAbstractSocket__SocketError) {
	QSslServer_ErrorOccurred(this.h, socket.cPointer(), (int)(error))
}
func (this *QSslServer) OnErrorOccurred(slot func(socket *QSslSocket, error QAbstractSocket__SocketError)) {
	QSslServer_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_ErrorOccurred
func miqt_exec_callback_QSslServer_ErrorOccurred(cb intptr_t, socket *QSslSocket, error int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket, error QAbstractSocket__SocketError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	slotval2 := (QAbstractSocket__SocketError)(error)

	gofunc(slotval1, slotval2)
}

func (this *QSslServer) PreSharedKeyAuthenticationRequired(socket *QSslSocket, authenticator *QSslPreSharedKeyAuthenticator) {
	QSslServer_PreSharedKeyAuthenticationRequired(this.h, socket.cPointer(), authenticator.cPointer())
}
func (this *QSslServer) OnPreSharedKeyAuthenticationRequired(slot func(socket *QSslSocket, authenticator *QSslPreSharedKeyAuthenticator)) {
	QSslServer_connect_PreSharedKeyAuthenticationRequired(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_PreSharedKeyAuthenticationRequired
func miqt_exec_callback_QSslServer_PreSharedKeyAuthenticationRequired(cb intptr_t, socket *QSslSocket, authenticator *QSslPreSharedKeyAuthenticator) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket, authenticator *QSslPreSharedKeyAuthenticator))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	slotval2 := newQSslPreSharedKeyAuthenticator(authenticator)

	gofunc(slotval1, slotval2)
}

func (this *QSslServer) AlertSent(socket *QSslSocket, level QSsl__AlertLevel, typeVal QSsl__AlertType, description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QSslServer_AlertSent(this.h, socket.cPointer(), (int)(level), (int)(typeVal), description_ms)
}
func (this *QSslServer) OnAlertSent(slot func(socket *QSslSocket, level QSsl__AlertLevel, typeVal QSsl__AlertType, description string)) {
	QSslServer_connect_AlertSent(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_AlertSent
func miqt_exec_callback_QSslServer_AlertSent(cb intptr_t, socket *QSslSocket, level int, typeVal int, description struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket, level QSsl__AlertLevel, typeVal QSsl__AlertType, description string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	slotval2 := (QSsl__AlertLevel)(level)

	slotval3 := (QSsl__AlertType)(typeVal)

	var description_ms struct_miqt_string = description
	description_ret := GoStringN(description_ms.data, int(int64(description_ms.len)))
	free(unsafe.Pointer(description_ms.data))
	slotval4 := description_ret

	gofunc(slotval1, slotval2, slotval3, slotval4)
}

func (this *QSslServer) AlertReceived(socket *QSslSocket, level QSsl__AlertLevel, typeVal QSsl__AlertType, description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QSslServer_AlertReceived(this.h, socket.cPointer(), (int)(level), (int)(typeVal), description_ms)
}
func (this *QSslServer) OnAlertReceived(slot func(socket *QSslSocket, level QSsl__AlertLevel, typeVal QSsl__AlertType, description string)) {
	QSslServer_connect_AlertReceived(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_AlertReceived
func miqt_exec_callback_QSslServer_AlertReceived(cb intptr_t, socket *QSslSocket, level int, typeVal int, description struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket, level QSsl__AlertLevel, typeVal QSsl__AlertType, description string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	slotval2 := (QSsl__AlertLevel)(level)

	slotval3 := (QSsl__AlertType)(typeVal)

	var description_ms struct_miqt_string = description
	description_ret := GoStringN(description_ms.data, int(int64(description_ms.len)))
	free(unsafe.Pointer(description_ms.data))
	slotval4 := description_ret

	gofunc(slotval1, slotval2, slotval3, slotval4)
}

func (this *QSslServer) HandshakeInterruptedOnError(socket *QSslSocket, error *QSslError) {
	QSslServer_HandshakeInterruptedOnError(this.h, socket.cPointer(), error.cPointer())
}
func (this *QSslServer) OnHandshakeInterruptedOnError(slot func(socket *QSslSocket, error *QSslError)) {
	QSslServer_connect_HandshakeInterruptedOnError(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_HandshakeInterruptedOnError
func miqt_exec_callback_QSslServer_HandshakeInterruptedOnError(cb intptr_t, socket *QSslSocket, error *QSslError) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket, error *QSslError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	slotval2 := newQSslError(error)

	gofunc(slotval1, slotval2)
}

func (this *QSslServer) StartedEncryptionHandshake(socket *QSslSocket) {
	QSslServer_StartedEncryptionHandshake(this.h, socket.cPointer())
}
func (this *QSslServer) OnStartedEncryptionHandshake(slot func(socket *QSslSocket)) {
	QSslServer_connect_StartedEncryptionHandshake(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_StartedEncryptionHandshake
func miqt_exec_callback_QSslServer_StartedEncryptionHandshake(cb intptr_t, socket *QSslSocket) {
	gofunc, ok := cgo.Handle(cb).Value().(func(socket *QSslSocket))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslSocket(socket)

	gofunc(slotval1)
}

func QSslServer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSslServer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSslServer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSslServer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslServer) callVirtualBase_IncomingConnection(socket uintptr) {

	QSslServer_virtualbase_IncomingConnection(unsafe.Pointer(this.h), (intptr_t)(socket))

}
func (this *QSslServer) OnIncomingConnection(slot func(super func(socket uintptr), socket uintptr)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSslServer_override_virtual_IncomingConnection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_IncomingConnection
func miqt_exec_callback_QSslServer_IncomingConnection(self QSslServer, cb intptr_t, socket intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(socket uintptr), socket uintptr))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uintptr)(socket)

	gofunc((&QSslServer{h: self}).callVirtualBase_IncomingConnection, slotval1)

}

func (this *QSslServer) callVirtualBase_HasPendingConnections() bool {

	return (bool)(QSslServer_virtualbase_HasPendingConnections(unsafe.Pointer(this.h)))

}
func (this *QSslServer) OnHasPendingConnections(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSslServer_override_virtual_HasPendingConnections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_HasPendingConnections
func miqt_exec_callback_QSslServer_HasPendingConnections(self QSslServer, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSslServer{h: self}).callVirtualBase_HasPendingConnections)

	return (bool)(virtualReturn)

}

func (this *QSslServer) callVirtualBase_NextPendingConnection() *QTcpSocket {

	return newQTcpSocket(QSslServer_virtualbase_NextPendingConnection(unsafe.Pointer(this.h)))

}
func (this *QSslServer) OnNextPendingConnection(slot func(super func() *QTcpSocket) *QTcpSocket) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSslServer_override_virtual_NextPendingConnection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSslServer_NextPendingConnection
func miqt_exec_callback_QSslServer_NextPendingConnection(self QSslServer, cb intptr_t) *QTcpSocket {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QTcpSocket) *QTcpSocket)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSslServer{h: self}).callVirtualBase_NextPendingConnection)

	return virtualReturn.cPointer()

}
