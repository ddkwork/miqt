package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QNetworkAccessManager__Operation int

const (
	QNetworkAccessManager__HeadOperation    QNetworkAccessManager__Operation = 1
	QNetworkAccessManager__GetOperation     QNetworkAccessManager__Operation = 2
	QNetworkAccessManager__PutOperation     QNetworkAccessManager__Operation = 3
	QNetworkAccessManager__PostOperation    QNetworkAccessManager__Operation = 4
	QNetworkAccessManager__DeleteOperation  QNetworkAccessManager__Operation = 5
	QNetworkAccessManager__CustomOperation  QNetworkAccessManager__Operation = 6
	QNetworkAccessManager__UnknownOperation QNetworkAccessManager__Operation = 0
)

type QNetworkAccessManager struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkAccessManager constructs a new QNetworkAccessManager object.
func NewQNetworkAccessManager() *QNetworkAccessManager {
	ret := newQNetworkAccessManager(QNetworkAccessManager_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkAccessManager2 constructs a new QNetworkAccessManager object.
func NewQNetworkAccessManager2(parent *qt.QObject) *QNetworkAccessManager {
	ret := newQNetworkAccessManager(QNetworkAccessManager_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkAccessManager) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QNetworkAccessManager_MetaObject(this.h)))
}

func (this *QNetworkAccessManager) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QNetworkAccessManager_Metacast(this.h, param1_Cstring))
}

func QNetworkAccessManager_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QNetworkAccessManager_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkAccessManager) SupportedSchemes() []string {
	var _ma struct_miqt_array = QNetworkAccessManager_SupportedSchemes(this.h)
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

func (this *QNetworkAccessManager) ClearAccessCache() {
	QNetworkAccessManager_ClearAccessCache(this.h)
}

func (this *QNetworkAccessManager) ClearConnectionCache() {
	QNetworkAccessManager_ClearConnectionCache(this.h)
}

func (this *QNetworkAccessManager) Proxy() *QNetworkProxy {
	_goptr := newQNetworkProxy(QNetworkAccessManager_Proxy(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkAccessManager) SetProxy(proxy *QNetworkProxy) {
	QNetworkAccessManager_SetProxy(this.h, proxy.cPointer())
}

func (this *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory {
	return newQNetworkProxyFactory(QNetworkAccessManager_ProxyFactory(this.h))
}

func (this *QNetworkAccessManager) SetProxyFactory(factory *QNetworkProxyFactory) {
	QNetworkAccessManager_SetProxyFactory(this.h, factory.cPointer())
}

func (this *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	return newQAbstractNetworkCache(QNetworkAccessManager_Cache(this.h))
}

func (this *QNetworkAccessManager) SetCache(cache *QAbstractNetworkCache) {
	QNetworkAccessManager_SetCache(this.h, cache.cPointer())
}

func (this *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	return newQNetworkCookieJar(QNetworkAccessManager_CookieJar(this.h))
}

func (this *QNetworkAccessManager) SetCookieJar(cookieJar *QNetworkCookieJar) {
	QNetworkAccessManager_SetCookieJar(this.h, cookieJar.cPointer())
}

func (this *QNetworkAccessManager) SetStrictTransportSecurityEnabled(enabled bool) {
	QNetworkAccessManager_SetStrictTransportSecurityEnabled(this.h, (bool)(enabled))
}

func (this *QNetworkAccessManager) IsStrictTransportSecurityEnabled() bool {
	return (bool)(QNetworkAccessManager_IsStrictTransportSecurityEnabled(this.h))
}

func (this *QNetworkAccessManager) EnableStrictTransportSecurityStore(enabled bool) {
	QNetworkAccessManager_EnableStrictTransportSecurityStore(this.h, (bool)(enabled))
}

func (this *QNetworkAccessManager) IsStrictTransportSecurityStoreEnabled() bool {
	return (bool)(QNetworkAccessManager_IsStrictTransportSecurityStoreEnabled(this.h))
}

func (this *QNetworkAccessManager) AddStrictTransportSecurityHosts(knownHosts []QHstsPolicy) {
	knownHosts_CArray := (*[0xffff]*QHstsPolicy)(malloc(size_t(8 * len(knownHosts))))
	defer free(unsafe.Pointer(knownHosts_CArray))
	for i := range knownHosts {
		knownHosts_CArray[i] = knownHosts[i].cPointer()
	}
	knownHosts_ma := struct_miqt_array{len: size_t(len(knownHosts)), data: unsafe.Pointer(knownHosts_CArray)}
	QNetworkAccessManager_AddStrictTransportSecurityHosts(this.h, knownHosts_ma)
}

func (this *QNetworkAccessManager) StrictTransportSecurityHosts() []QHstsPolicy {
	var _ma struct_miqt_array = QNetworkAccessManager_StrictTransportSecurityHosts(this.h)
	_ret := make([]QHstsPolicy, int(_ma.len))
	_outCast := (*[0xffff]*QHstsPolicy)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQHstsPolicy(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QNetworkAccessManager) Head(request *QNetworkRequest) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_Head(this.h, request.cPointer()))
}

func (this *QNetworkAccessManager) Get(request *QNetworkRequest) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_Get(this.h, request.cPointer()))
}

func (this *QNetworkAccessManager) Get2(request *QNetworkRequest, data *qt.QIODevice) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_Get2(this.h, request.cPointer(), (*QIODevice)(data.UnsafePointer())))
}

func (this *QNetworkAccessManager) Get3(request *QNetworkRequest, data []byte) *QNetworkReply {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QNetworkAccessManager_Get3(this.h, request.cPointer(), data_alias))
}

func (this *QNetworkAccessManager) Post(request *QNetworkRequest, data *qt.QIODevice) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_Post(this.h, request.cPointer(), (*QIODevice)(data.UnsafePointer())))
}

func (this *QNetworkAccessManager) Post2(request *QNetworkRequest, data []byte) *QNetworkReply {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QNetworkAccessManager_Post2(this.h, request.cPointer(), data_alias))
}

func (this *QNetworkAccessManager) Put(request *QNetworkRequest, data *qt.QIODevice) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_Put(this.h, request.cPointer(), (*QIODevice)(data.UnsafePointer())))
}

func (this *QNetworkAccessManager) Put2(request *QNetworkRequest, data []byte) *QNetworkReply {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QNetworkAccessManager_Put2(this.h, request.cPointer(), data_alias))
}

func (this *QNetworkAccessManager) DeleteResource(request *QNetworkRequest) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_DeleteResource(this.h, request.cPointer()))
}

func (this *QNetworkAccessManager) SendCustomRequest(request *QNetworkRequest, verb []byte) *QNetworkReply {
	verb_alias := struct_miqt_string{}
	verb_alias.data = (char)(unsafe.Pointer(&verb[0]))
	verb_alias.len = size_t(len(verb))
	return newQNetworkReply(QNetworkAccessManager_SendCustomRequest(this.h, request.cPointer(), verb_alias))
}

func (this *QNetworkAccessManager) SendCustomRequest2(request *QNetworkRequest, verb []byte, data []byte) *QNetworkReply {
	verb_alias := struct_miqt_string{}
	verb_alias.data = (char)(unsafe.Pointer(&verb[0]))
	verb_alias.len = size_t(len(verb))
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QNetworkAccessManager_SendCustomRequest2(this.h, request.cPointer(), verb_alias, data_alias))
}

func (this *QNetworkAccessManager) Post4(request *QNetworkRequest, multiPart *QHttpMultiPart) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_Post4(this.h, request.cPointer(), multiPart.cPointer()))
}

func (this *QNetworkAccessManager) Put4(request *QNetworkRequest, multiPart *QHttpMultiPart) *QNetworkReply {
	return newQNetworkReply(QNetworkAccessManager_Put4(this.h, request.cPointer(), multiPart.cPointer()))
}

func (this *QNetworkAccessManager) SendCustomRequest3(request *QNetworkRequest, verb []byte, multiPart *QHttpMultiPart) *QNetworkReply {
	verb_alias := struct_miqt_string{}
	verb_alias.data = (char)(unsafe.Pointer(&verb[0]))
	verb_alias.len = size_t(len(verb))
	return newQNetworkReply(QNetworkAccessManager_SendCustomRequest3(this.h, request.cPointer(), verb_alias, multiPart.cPointer()))
}

func (this *QNetworkAccessManager) ConnectToHostEncrypted(hostName string) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QNetworkAccessManager_ConnectToHostEncrypted(this.h, hostName_ms)
}

func (this *QNetworkAccessManager) ConnectToHostEncrypted2(hostName string, port uint16, sslConfiguration *QSslConfiguration, peerName string) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	peerName_ms := struct_miqt_string{}
	peerName_ms.data = CString(peerName)
	peerName_ms.len = size_t(len(peerName))
	defer free(unsafe.Pointer(peerName_ms.data))
	QNetworkAccessManager_ConnectToHostEncrypted2(this.h, hostName_ms, (uint16_t)(port), sslConfiguration.cPointer(), peerName_ms)
}

func (this *QNetworkAccessManager) ConnectToHost(hostName string) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QNetworkAccessManager_ConnectToHost(this.h, hostName_ms)
}

func (this *QNetworkAccessManager) SetRedirectPolicy(policy QNetworkRequest__RedirectPolicy) {
	QNetworkAccessManager_SetRedirectPolicy(this.h, (int)(policy))
}

func (this *QNetworkAccessManager) RedirectPolicy() QNetworkRequest__RedirectPolicy {
	return (QNetworkRequest__RedirectPolicy)(QNetworkAccessManager_RedirectPolicy(this.h))
}

func (this *QNetworkAccessManager) AutoDeleteReplies() bool {
	return (bool)(QNetworkAccessManager_AutoDeleteReplies(this.h))
}

func (this *QNetworkAccessManager) SetAutoDeleteReplies(autoDelete bool) {
	QNetworkAccessManager_SetAutoDeleteReplies(this.h, (bool)(autoDelete))
}

func (this *QNetworkAccessManager) TransferTimeout() int {
	return (int)(QNetworkAccessManager_TransferTimeout(this.h))
}

func (this *QNetworkAccessManager) SetTransferTimeout(timeout int) {
	QNetworkAccessManager_SetTransferTimeout(this.h, (int)(timeout))
}

func (this *QNetworkAccessManager) SetTransferTimeout2() {
	QNetworkAccessManager_SetTransferTimeout2(this.h)
}

func (this *QNetworkAccessManager) ProxyAuthenticationRequired(proxy *QNetworkProxy, authenticator *QAuthenticator) {
	QNetworkAccessManager_ProxyAuthenticationRequired(this.h, proxy.cPointer(), authenticator.cPointer())
}

func (this *QNetworkAccessManager) OnProxyAuthenticationRequired(slot func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	QNetworkAccessManager_connect_ProxyAuthenticationRequired(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_ProxyAuthenticationRequired
func miqt_exec_callback_QNetworkAccessManager_ProxyAuthenticationRequired(cb intptr_t, proxy *QNetworkProxy, authenticator *QAuthenticator) {
	gofunc, ok := cgo.Handle(cb).Value().(func(proxy *QNetworkProxy, authenticator *QAuthenticator))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkProxy(proxy)

	slotval2 := newQAuthenticator(authenticator)

	gofunc(slotval1, slotval2)
}

func (this *QNetworkAccessManager) AuthenticationRequired(reply *QNetworkReply, authenticator *QAuthenticator) {
	QNetworkAccessManager_AuthenticationRequired(this.h, reply.cPointer(), authenticator.cPointer())
}

func (this *QNetworkAccessManager) OnAuthenticationRequired(slot func(reply *QNetworkReply, authenticator *QAuthenticator)) {
	QNetworkAccessManager_connect_AuthenticationRequired(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_AuthenticationRequired
func miqt_exec_callback_QNetworkAccessManager_AuthenticationRequired(cb intptr_t, reply *QNetworkReply, authenticator *QAuthenticator) {
	gofunc, ok := cgo.Handle(cb).Value().(func(reply *QNetworkReply, authenticator *QAuthenticator))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkReply(reply)

	slotval2 := newQAuthenticator(authenticator)

	gofunc(slotval1, slotval2)
}

func (this *QNetworkAccessManager) Finished(reply *QNetworkReply) {
	QNetworkAccessManager_Finished(this.h, reply.cPointer())
}

func (this *QNetworkAccessManager) OnFinished(slot func(reply *QNetworkReply)) {
	QNetworkAccessManager_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_Finished
func miqt_exec_callback_QNetworkAccessManager_Finished(cb intptr_t, reply *QNetworkReply) {
	gofunc, ok := cgo.Handle(cb).Value().(func(reply *QNetworkReply))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkReply(reply)

	gofunc(slotval1)
}

func (this *QNetworkAccessManager) Encrypted(reply *QNetworkReply) {
	QNetworkAccessManager_Encrypted(this.h, reply.cPointer())
}

func (this *QNetworkAccessManager) OnEncrypted(slot func(reply *QNetworkReply)) {
	QNetworkAccessManager_connect_Encrypted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_Encrypted
func miqt_exec_callback_QNetworkAccessManager_Encrypted(cb intptr_t, reply *QNetworkReply) {
	gofunc, ok := cgo.Handle(cb).Value().(func(reply *QNetworkReply))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkReply(reply)

	gofunc(slotval1)
}

func (this *QNetworkAccessManager) SslErrors(reply *QNetworkReply, errors []QSslError) {
	errors_CArray := (*[0xffff]*QSslError)(malloc(size_t(8 * len(errors))))
	defer free(unsafe.Pointer(errors_CArray))
	for i := range errors {
		errors_CArray[i] = errors[i].cPointer()
	}
	errors_ma := struct_miqt_array{len: size_t(len(errors)), data: unsafe.Pointer(errors_CArray)}
	QNetworkAccessManager_SslErrors(this.h, reply.cPointer(), errors_ma)
}

func (this *QNetworkAccessManager) OnSslErrors(slot func(reply *QNetworkReply, errors []QSslError)) {
	QNetworkAccessManager_connect_SslErrors(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_SslErrors
func miqt_exec_callback_QNetworkAccessManager_SslErrors(cb intptr_t, reply *QNetworkReply, errors struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(reply *QNetworkReply, errors []QSslError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkReply(reply)

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

func (this *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator) {
	QNetworkAccessManager_PreSharedKeyAuthenticationRequired(this.h, reply.cPointer(), authenticator.cPointer())
}

func (this *QNetworkAccessManager) OnPreSharedKeyAuthenticationRequired(slot func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {
	QNetworkAccessManager_connect_PreSharedKeyAuthenticationRequired(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_PreSharedKeyAuthenticationRequired
func miqt_exec_callback_QNetworkAccessManager_PreSharedKeyAuthenticationRequired(cb intptr_t, reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator) {
	gofunc, ok := cgo.Handle(cb).Value().(func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkReply(reply)

	slotval2 := newQSslPreSharedKeyAuthenticator(authenticator)

	gofunc(slotval1, slotval2)
}

func QNetworkAccessManager_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkAccessManager_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkAccessManager_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkAccessManager_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkAccessManager) EnableStrictTransportSecurityStore2(enabled bool, storeDir string) {
	storeDir_ms := struct_miqt_string{}
	storeDir_ms.data = CString(storeDir)
	storeDir_ms.len = size_t(len(storeDir))
	defer free(unsafe.Pointer(storeDir_ms.data))
	QNetworkAccessManager_EnableStrictTransportSecurityStore2(this.h, (bool)(enabled), storeDir_ms)
}

func (this *QNetworkAccessManager) SendCustomRequest32(request *QNetworkRequest, verb []byte, data *qt.QIODevice) *QNetworkReply {
	verb_alias := struct_miqt_string{}
	verb_alias.data = (char)(unsafe.Pointer(&verb[0]))
	verb_alias.len = size_t(len(verb))
	return newQNetworkReply(QNetworkAccessManager_SendCustomRequest32(this.h, request.cPointer(), verb_alias, (*QIODevice)(data.UnsafePointer())))
}

func (this *QNetworkAccessManager) ConnectToHostEncrypted22(hostName string, port uint16) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QNetworkAccessManager_ConnectToHostEncrypted22(this.h, hostName_ms, (uint16_t)(port))
}

func (this *QNetworkAccessManager) ConnectToHostEncrypted3(hostName string, port uint16, sslConfiguration *QSslConfiguration) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QNetworkAccessManager_ConnectToHostEncrypted3(this.h, hostName_ms, (uint16_t)(port), sslConfiguration.cPointer())
}

func (this *QNetworkAccessManager) ConnectToHost2(hostName string, port uint16) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QNetworkAccessManager_ConnectToHost2(this.h, hostName_ms, (uint16_t)(port))
}

func (this *QNetworkAccessManager) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QNetworkAccessManager_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QNetworkAccessManager) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkAccessManager_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_MetaObject
func miqt_exec_callback_QNetworkAccessManager_MetaObject(self QNetworkAccessManager, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNetworkAccessManager{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QNetworkAccessManager) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QNetworkAccessManager_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QNetworkAccessManager) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkAccessManager_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkAccessManager_Metacast
func miqt_exec_callback_QNetworkAccessManager_Metacast(self QNetworkAccessManager, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QNetworkAccessManager{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
