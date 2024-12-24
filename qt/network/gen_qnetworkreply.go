package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QNetworkReply__NetworkError int

const (
	QNetworkReply__NoError                           QNetworkReply__NetworkError = 0
	QNetworkReply__ConnectionRefusedError            QNetworkReply__NetworkError = 1
	QNetworkReply__RemoteHostClosedError             QNetworkReply__NetworkError = 2
	QNetworkReply__HostNotFoundError                 QNetworkReply__NetworkError = 3
	QNetworkReply__TimeoutError                      QNetworkReply__NetworkError = 4
	QNetworkReply__OperationCanceledError            QNetworkReply__NetworkError = 5
	QNetworkReply__SslHandshakeFailedError           QNetworkReply__NetworkError = 6
	QNetworkReply__TemporaryNetworkFailureError      QNetworkReply__NetworkError = 7
	QNetworkReply__NetworkSessionFailedError         QNetworkReply__NetworkError = 8
	QNetworkReply__BackgroundRequestNotAllowedError  QNetworkReply__NetworkError = 9
	QNetworkReply__TooManyRedirectsError             QNetworkReply__NetworkError = 10
	QNetworkReply__InsecureRedirectError             QNetworkReply__NetworkError = 11
	QNetworkReply__UnknownNetworkError               QNetworkReply__NetworkError = 99
	QNetworkReply__ProxyConnectionRefusedError       QNetworkReply__NetworkError = 101
	QNetworkReply__ProxyConnectionClosedError        QNetworkReply__NetworkError = 102
	QNetworkReply__ProxyNotFoundError                QNetworkReply__NetworkError = 103
	QNetworkReply__ProxyTimeoutError                 QNetworkReply__NetworkError = 104
	QNetworkReply__ProxyAuthenticationRequiredError  QNetworkReply__NetworkError = 105
	QNetworkReply__UnknownProxyError                 QNetworkReply__NetworkError = 199
	QNetworkReply__ContentAccessDenied               QNetworkReply__NetworkError = 201
	QNetworkReply__ContentOperationNotPermittedError QNetworkReply__NetworkError = 202
	QNetworkReply__ContentNotFoundError              QNetworkReply__NetworkError = 203
	QNetworkReply__AuthenticationRequiredError       QNetworkReply__NetworkError = 204
	QNetworkReply__ContentReSendError                QNetworkReply__NetworkError = 205
	QNetworkReply__ContentConflictError              QNetworkReply__NetworkError = 206
	QNetworkReply__ContentGoneError                  QNetworkReply__NetworkError = 207
	QNetworkReply__UnknownContentError               QNetworkReply__NetworkError = 299
	QNetworkReply__ProtocolUnknownError              QNetworkReply__NetworkError = 301
	QNetworkReply__ProtocolInvalidOperationError     QNetworkReply__NetworkError = 302
	QNetworkReply__ProtocolFailure                   QNetworkReply__NetworkError = 399
	QNetworkReply__InternalServerError               QNetworkReply__NetworkError = 401
	QNetworkReply__OperationNotImplementedError      QNetworkReply__NetworkError = 402
	QNetworkReply__ServiceUnavailableError           QNetworkReply__NetworkError = 403
	QNetworkReply__UnknownServerError                QNetworkReply__NetworkError = 499
)

type QNetworkReply struct {
	h          uintptr
	isSubclass bool
}

func (this *QNetworkReply) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QNetworkReply_MetaObject(this.h)))
}

func (this *QNetworkReply) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QNetworkReply_Metacast(this.h, param1_Cstring))
}

func QNetworkReply_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QNetworkReply_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkReply) Close() {
	QNetworkReply_Close(this.h)
}

func (this *QNetworkReply) IsSequential() bool {
	return (bool)(QNetworkReply_IsSequential(this.h))
}

func (this *QNetworkReply) ReadBufferSize() int64 {
	return (int64)(QNetworkReply_ReadBufferSize(this.h))
}

func (this *QNetworkReply) SetReadBufferSize(size int64) {
	QNetworkReply_SetReadBufferSize(this.h, (longlong)(size))
}

func (this *QNetworkReply) Manager() *QNetworkAccessManager {
	return newQNetworkAccessManager(QNetworkReply_Manager(this.h))
}

func (this *QNetworkReply) Operation() QNetworkAccessManager__Operation {
	return (QNetworkAccessManager__Operation)(QNetworkReply_Operation(this.h))
}

func (this *QNetworkReply) Request() *QNetworkRequest {
	_goptr := newQNetworkRequest(QNetworkReply_Request(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkReply) Error() NetworkError {
	xxxxxxxxx
}

func (this *QNetworkReply) IsFinished() bool {
	return (bool)(QNetworkReply_IsFinished(this.h))
}

func (this *QNetworkReply) IsRunning() bool {
	return (bool)(QNetworkReply_IsRunning(this.h))
}

func (this *QNetworkReply) Url() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QNetworkReply_Url(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkReply) Header(header QNetworkRequest__KnownHeaders) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkReply_Header(this.h, (int)(header))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkReply) HasRawHeader(headerName qt.QAnyStringView) bool {
	return (bool)(QNetworkReply_HasRawHeader(this.h, (*QAnyStringView)(headerName.UnsafePointer())))
}

func (this *QNetworkReply) RawHeaderList() [][]byte {
	var _ma struct_miqt_array = QNetworkReply_RawHeaderList(this.h)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QNetworkReply) RawHeader(headerName qt.QAnyStringView) []byte {
	var _bytearray struct_miqt_string = QNetworkReply_RawHeader(this.h, (*QAnyStringView)(headerName.UnsafePointer()))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkReply) RawHeaderPairs() []RawHeaderPair {
	var _ma struct_miqt_array = QNetworkReply_RawHeaderPairs(this.h)
	_ret := make([]RawHeaderPair, int(_ma.len))
	_outCast := (*[0xffff]RawHeaderPair)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QNetworkReply) Headers() *QHttpHeaders {
	_goptr := newQHttpHeaders(QNetworkReply_Headers(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkReply) Attribute(code QNetworkRequest__Attribute) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkReply_Attribute(this.h, (int)(code))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkReply) SslConfiguration() *QSslConfiguration {
	_goptr := newQSslConfiguration(QNetworkReply_SslConfiguration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkReply) SetSslConfiguration(configuration *QSslConfiguration) {
	QNetworkReply_SetSslConfiguration(this.h, configuration.cPointer())
}

func (this *QNetworkReply) IgnoreSslErrors(errors []QSslError) {
	errors_CArray := (*[0xffff]*QSslError)(malloc(size_t(8 * len(errors))))
	defer free(unsafe.Pointer(errors_CArray))
	for i := range errors {
		errors_CArray[i] = errors[i].cPointer()
	}
	errors_ma := struct_miqt_array{len: size_t(len(errors)), data: unsafe.Pointer(errors_CArray)}
	QNetworkReply_IgnoreSslErrors(this.h, errors_ma)
}

func (this *QNetworkReply) Abort() {
	QNetworkReply_Abort(this.h)
}

func (this *QNetworkReply) IgnoreSslErrors2() {
	QNetworkReply_IgnoreSslErrors2(this.h)
}

func (this *QNetworkReply) SocketStartedConnecting() {
	QNetworkReply_SocketStartedConnecting(this.h)
}

func (this *QNetworkReply) OnSocketStartedConnecting(slot func()) {
	QNetworkReply_connect_SocketStartedConnecting(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_SocketStartedConnecting
func miqt_exec_callback_QNetworkReply_SocketStartedConnecting(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkReply) RequestSent() {
	QNetworkReply_RequestSent(this.h)
}

func (this *QNetworkReply) OnRequestSent(slot func()) {
	QNetworkReply_connect_RequestSent(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_RequestSent
func miqt_exec_callback_QNetworkReply_RequestSent(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkReply) MetaDataChanged() {
	QNetworkReply_MetaDataChanged(this.h)
}

func (this *QNetworkReply) OnMetaDataChanged(slot func()) {
	QNetworkReply_connect_MetaDataChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_MetaDataChanged
func miqt_exec_callback_QNetworkReply_MetaDataChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkReply) Finished() {
	QNetworkReply_Finished(this.h)
}

func (this *QNetworkReply) OnFinished(slot func()) {
	QNetworkReply_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_Finished
func miqt_exec_callback_QNetworkReply_Finished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkReply) ErrorOccurred(param1 QNetworkReply__NetworkError) {
	QNetworkReply_ErrorOccurred(this.h, (int)(param1))
}

func (this *QNetworkReply) OnErrorOccurred(slot func(param1 QNetworkReply__NetworkError)) {
	QNetworkReply_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_ErrorOccurred
func miqt_exec_callback_QNetworkReply_ErrorOccurred(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QNetworkReply__NetworkError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkReply__NetworkError)(param1)

	gofunc(slotval1)
}

func (this *QNetworkReply) Encrypted() {
	QNetworkReply_Encrypted(this.h)
}

func (this *QNetworkReply) OnEncrypted(slot func()) {
	QNetworkReply_connect_Encrypted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_Encrypted
func miqt_exec_callback_QNetworkReply_Encrypted(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkReply) SslErrors(errors []QSslError) {
	errors_CArray := (*[0xffff]*QSslError)(malloc(size_t(8 * len(errors))))
	defer free(unsafe.Pointer(errors_CArray))
	for i := range errors {
		errors_CArray[i] = errors[i].cPointer()
	}
	errors_ma := struct_miqt_array{len: size_t(len(errors)), data: unsafe.Pointer(errors_CArray)}
	QNetworkReply_SslErrors(this.h, errors_ma)
}

func (this *QNetworkReply) OnSslErrors(slot func(errors []QSslError)) {
	QNetworkReply_connect_SslErrors(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_SslErrors
func miqt_exec_callback_QNetworkReply_SslErrors(cb intptr_t, errors struct_miqt_array) {
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

func (this *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator *QSslPreSharedKeyAuthenticator) {
	QNetworkReply_PreSharedKeyAuthenticationRequired(this.h, authenticator.cPointer())
}

func (this *QNetworkReply) OnPreSharedKeyAuthenticationRequired(slot func(authenticator *QSslPreSharedKeyAuthenticator)) {
	QNetworkReply_connect_PreSharedKeyAuthenticationRequired(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_PreSharedKeyAuthenticationRequired
func miqt_exec_callback_QNetworkReply_PreSharedKeyAuthenticationRequired(cb intptr_t, authenticator *QSslPreSharedKeyAuthenticator) {
	gofunc, ok := cgo.Handle(cb).Value().(func(authenticator *QSslPreSharedKeyAuthenticator))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSslPreSharedKeyAuthenticator(authenticator)

	gofunc(slotval1)
}

func (this *QNetworkReply) Redirected(url *qt.QUrl) {
	QNetworkReply_Redirected(this.h, (*QUrl)(url.UnsafePointer()))
}

func (this *QNetworkReply) OnRedirected(slot func(url *qt.QUrl)) {
	QNetworkReply_connect_Redirected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_Redirected
func miqt_exec_callback_QNetworkReply_Redirected(cb intptr_t, url *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(url *qt.QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(url))

	gofunc(slotval1)
}

func (this *QNetworkReply) RedirectAllowed() {
	QNetworkReply_RedirectAllowed(this.h)
}

func (this *QNetworkReply) OnRedirectAllowed(slot func()) {
	QNetworkReply_connect_RedirectAllowed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_RedirectAllowed
func miqt_exec_callback_QNetworkReply_RedirectAllowed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	QNetworkReply_UploadProgress(this.h, (longlong)(bytesSent), (longlong)(bytesTotal))
}

func (this *QNetworkReply) OnUploadProgress(slot func(bytesSent int64, bytesTotal int64)) {
	QNetworkReply_connect_UploadProgress(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_UploadProgress
func miqt_exec_callback_QNetworkReply_UploadProgress(cb intptr_t, bytesSent longlong, bytesTotal longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(bytesSent int64, bytesTotal int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(bytesSent)

	slotval2 := (int64)(bytesTotal)

	gofunc(slotval1, slotval2)
}

func (this *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	QNetworkReply_DownloadProgress(this.h, (longlong)(bytesReceived), (longlong)(bytesTotal))
}

func (this *QNetworkReply) OnDownloadProgress(slot func(bytesReceived int64, bytesTotal int64)) {
	QNetworkReply_connect_DownloadProgress(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkReply_DownloadProgress
func miqt_exec_callback_QNetworkReply_DownloadProgress(cb intptr_t, bytesReceived longlong, bytesTotal longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(bytesReceived int64, bytesTotal int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(bytesReceived)

	slotval2 := (int64)(bytesTotal)

	gofunc(slotval1, slotval2)
}

func QNetworkReply_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkReply_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkReply_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkReply_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
