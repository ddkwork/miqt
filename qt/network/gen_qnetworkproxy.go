package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QNetworkProxyQuery__QueryType int

const (
	QNetworkProxyQuery__TcpSocket  QNetworkProxyQuery__QueryType = 0
	QNetworkProxyQuery__UdpSocket  QNetworkProxyQuery__QueryType = 1
	QNetworkProxyQuery__SctpSocket QNetworkProxyQuery__QueryType = 2
	QNetworkProxyQuery__TcpServer  QNetworkProxyQuery__QueryType = 100
	QNetworkProxyQuery__UrlRequest QNetworkProxyQuery__QueryType = 101
	QNetworkProxyQuery__SctpServer QNetworkProxyQuery__QueryType = 102
)

type QNetworkProxy__ProxyType int

const (
	QNetworkProxy__DefaultProxy     QNetworkProxy__ProxyType = 0
	QNetworkProxy__Socks5Proxy      QNetworkProxy__ProxyType = 1
	QNetworkProxy__NoProxy          QNetworkProxy__ProxyType = 2
	QNetworkProxy__HttpProxy        QNetworkProxy__ProxyType = 3
	QNetworkProxy__HttpCachingProxy QNetworkProxy__ProxyType = 4
	QNetworkProxy__FtpCachingProxy  QNetworkProxy__ProxyType = 5
)

type QNetworkProxy__Capability int

const (
	QNetworkProxy__TunnelingCapability      QNetworkProxy__Capability = 1
	QNetworkProxy__ListeningCapability      QNetworkProxy__Capability = 2
	QNetworkProxy__UdpTunnelingCapability   QNetworkProxy__Capability = 4
	QNetworkProxy__CachingCapability        QNetworkProxy__Capability = 8
	QNetworkProxy__HostNameLookupCapability QNetworkProxy__Capability = 16
	QNetworkProxy__SctpTunnelingCapability  QNetworkProxy__Capability = 32
	QNetworkProxy__SctpListeningCapability  QNetworkProxy__Capability = 64
)

type QNetworkProxyQuery struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkProxyQuery constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery2 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery2(requestUrl *qt.QUrl) *QNetworkProxyQuery {
	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new2((*QUrl)(requestUrl.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery3 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery3(hostname string, port int) *QNetworkProxyQuery {
	hostname_ms := struct_miqt_string{}
	hostname_ms.data = CString(hostname)
	hostname_ms.len = size_t(len(hostname))
	defer free(unsafe.Pointer(hostname_ms.data))

	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new3(hostname_ms, (int)(port)))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery4 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery4(bindPort uint16) *QNetworkProxyQuery {
	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new4((uint16_t)(bindPort)))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery5 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery5(other *QNetworkProxyQuery) *QNetworkProxyQuery {
	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new5(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery6 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery6(requestUrl *qt.QUrl, queryType QueryType) *QNetworkProxyQuery {
	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new6((*QUrl)(requestUrl.UnsafePointer()), queryType))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery7 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery7(hostname string, port int, protocolTag string) *QNetworkProxyQuery {
	hostname_ms := struct_miqt_string{}
	hostname_ms.data = CString(hostname)
	hostname_ms.len = size_t(len(hostname))
	defer free(unsafe.Pointer(hostname_ms.data))
	protocolTag_ms := struct_miqt_string{}
	protocolTag_ms.data = CString(protocolTag)
	protocolTag_ms.len = size_t(len(protocolTag))
	defer free(unsafe.Pointer(protocolTag_ms.data))

	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new7(hostname_ms, (int)(port), protocolTag_ms))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery8 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery8(hostname string, port int, protocolTag string, queryType QueryType) *QNetworkProxyQuery {
	hostname_ms := struct_miqt_string{}
	hostname_ms.data = CString(hostname)
	hostname_ms.len = size_t(len(hostname))
	defer free(unsafe.Pointer(hostname_ms.data))
	protocolTag_ms := struct_miqt_string{}
	protocolTag_ms.data = CString(protocolTag)
	protocolTag_ms.len = size_t(len(protocolTag))
	defer free(unsafe.Pointer(protocolTag_ms.data))

	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new8(hostname_ms, (int)(port), protocolTag_ms, queryType))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery9 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery9(bindPort uint16, protocolTag string) *QNetworkProxyQuery {
	protocolTag_ms := struct_miqt_string{}
	protocolTag_ms.data = CString(protocolTag)
	protocolTag_ms.len = size_t(len(protocolTag))
	defer free(unsafe.Pointer(protocolTag_ms.data))

	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new9((uint16_t)(bindPort), protocolTag_ms))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxyQuery10 constructs a new QNetworkProxyQuery object.
func NewQNetworkProxyQuery10(bindPort uint16, protocolTag string, queryType QueryType) *QNetworkProxyQuery {
	protocolTag_ms := struct_miqt_string{}
	protocolTag_ms.data = CString(protocolTag)
	protocolTag_ms.len = size_t(len(protocolTag))
	defer free(unsafe.Pointer(protocolTag_ms.data))

	ret := newQNetworkProxyQuery(QNetworkProxyQuery_new10((uint16_t)(bindPort), protocolTag_ms, queryType))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkProxyQuery) OperatorAssign(other *QNetworkProxyQuery) {
	QNetworkProxyQuery_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkProxyQuery) Swap(other *QNetworkProxyQuery) {
	QNetworkProxyQuery_Swap(this.h, other.cPointer())
}

func (this *QNetworkProxyQuery) OperatorEqual(other *QNetworkProxyQuery) bool {
	return (bool)(QNetworkProxyQuery_OperatorEqual(this.h, other.cPointer()))
}

func (this *QNetworkProxyQuery) OperatorNotEqual(other *QNetworkProxyQuery) bool {
	return (bool)(QNetworkProxyQuery_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QNetworkProxyQuery) QueryType() QueryType {
	xxxxxxxxx
}

func (this *QNetworkProxyQuery) SetQueryType(typeVal QueryType) {
	QNetworkProxyQuery_SetQueryType(this.h, typeVal)
}

func (this *QNetworkProxyQuery) PeerPort() int {
	return (int)(QNetworkProxyQuery_PeerPort(this.h))
}

func (this *QNetworkProxyQuery) SetPeerPort(port int) {
	QNetworkProxyQuery_SetPeerPort(this.h, (int)(port))
}

func (this *QNetworkProxyQuery) PeerHostName() string {
	var _ms struct_miqt_string = QNetworkProxyQuery_PeerHostName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	hostname_ms := struct_miqt_string{}
	hostname_ms.data = CString(hostname)
	hostname_ms.len = size_t(len(hostname))
	defer free(unsafe.Pointer(hostname_ms.data))
	QNetworkProxyQuery_SetPeerHostName(this.h, hostname_ms)
}

func (this *QNetworkProxyQuery) LocalPort() int {
	return (int)(QNetworkProxyQuery_LocalPort(this.h))
}

func (this *QNetworkProxyQuery) SetLocalPort(port int) {
	QNetworkProxyQuery_SetLocalPort(this.h, (int)(port))
}

func (this *QNetworkProxyQuery) ProtocolTag() string {
	var _ms struct_miqt_string = QNetworkProxyQuery_ProtocolTag(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	protocolTag_ms := struct_miqt_string{}
	protocolTag_ms.data = CString(protocolTag)
	protocolTag_ms.len = size_t(len(protocolTag))
	defer free(unsafe.Pointer(protocolTag_ms.data))
	QNetworkProxyQuery_SetProtocolTag(this.h, protocolTag_ms)
}

func (this *QNetworkProxyQuery) Url() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QNetworkProxyQuery_Url(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkProxyQuery) SetUrl(url *qt.QUrl) {
	QNetworkProxyQuery_SetUrl(this.h, (*QUrl)(url.UnsafePointer()))
}

type QNetworkProxy struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkProxy constructs a new QNetworkProxy object.
func NewQNetworkProxy() *QNetworkProxy {
	ret := newQNetworkProxy(QNetworkProxy_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxy2 constructs a new QNetworkProxy object.
func NewQNetworkProxy2(typeVal ProxyType) *QNetworkProxy {
	ret := newQNetworkProxy(QNetworkProxy_new2(typeVal))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxy3 constructs a new QNetworkProxy object.
func NewQNetworkProxy3(other *QNetworkProxy) *QNetworkProxy {
	ret := newQNetworkProxy(QNetworkProxy_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxy4 constructs a new QNetworkProxy object.
func NewQNetworkProxy4(typeVal ProxyType, hostName string) *QNetworkProxy {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))

	ret := newQNetworkProxy(QNetworkProxy_new4(typeVal, hostName_ms))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxy5 constructs a new QNetworkProxy object.
func NewQNetworkProxy5(typeVal ProxyType, hostName string, port uint16) *QNetworkProxy {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))

	ret := newQNetworkProxy(QNetworkProxy_new5(typeVal, hostName_ms, (uint16_t)(port)))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxy6 constructs a new QNetworkProxy object.
func NewQNetworkProxy6(typeVal ProxyType, hostName string, port uint16, user string) *QNetworkProxy {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	user_ms := struct_miqt_string{}
	user_ms.data = CString(user)
	user_ms.len = size_t(len(user))
	defer free(unsafe.Pointer(user_ms.data))

	ret := newQNetworkProxy(QNetworkProxy_new6(typeVal, hostName_ms, (uint16_t)(port), user_ms))
	ret.isSubclass = true
	return ret
}

// NewQNetworkProxy7 constructs a new QNetworkProxy object.
func NewQNetworkProxy7(typeVal ProxyType, hostName string, port uint16, user string, password string) *QNetworkProxy {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	user_ms := struct_miqt_string{}
	user_ms.data = CString(user)
	user_ms.len = size_t(len(user))
	defer free(unsafe.Pointer(user_ms.data))
	password_ms := struct_miqt_string{}
	password_ms.data = CString(password)
	password_ms.len = size_t(len(password))
	defer free(unsafe.Pointer(password_ms.data))

	ret := newQNetworkProxy(QNetworkProxy_new7(typeVal, hostName_ms, (uint16_t)(port), user_ms, password_ms))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkProxy) OperatorAssign(other *QNetworkProxy) {
	QNetworkProxy_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkProxy) Swap(other *QNetworkProxy) {
	QNetworkProxy_Swap(this.h, other.cPointer())
}

func (this *QNetworkProxy) OperatorEqual(other *QNetworkProxy) bool {
	return (bool)(QNetworkProxy_OperatorEqual(this.h, other.cPointer()))
}

func (this *QNetworkProxy) OperatorNotEqual(other *QNetworkProxy) bool {
	return (bool)(QNetworkProxy_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QNetworkProxy) SetType(typeVal QNetworkProxy__ProxyType) {
	QNetworkProxy_SetType(this.h, (int)(typeVal))
}

func (this *QNetworkProxy) Type() QNetworkProxy__ProxyType {
	return (QNetworkProxy__ProxyType)(QNetworkProxy_Type(this.h))
}

func (this *QNetworkProxy) SetCapabilities(capab Capabilities) {
	QNetworkProxy_SetCapabilities(this.h, capab)
}

func (this *QNetworkProxy) Capabilities() Capabilities {
	xxxxxxxxx
}

func (this *QNetworkProxy) IsCachingProxy() bool {
	return (bool)(QNetworkProxy_IsCachingProxy(this.h))
}

func (this *QNetworkProxy) IsTransparentProxy() bool {
	return (bool)(QNetworkProxy_IsTransparentProxy(this.h))
}

func (this *QNetworkProxy) SetUser(userName string) {
	userName_ms := struct_miqt_string{}
	userName_ms.data = CString(userName)
	userName_ms.len = size_t(len(userName))
	defer free(unsafe.Pointer(userName_ms.data))
	QNetworkProxy_SetUser(this.h, userName_ms)
}

func (this *QNetworkProxy) User() string {
	var _ms struct_miqt_string = QNetworkProxy_User(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkProxy) SetPassword(password string) {
	password_ms := struct_miqt_string{}
	password_ms.data = CString(password)
	password_ms.len = size_t(len(password))
	defer free(unsafe.Pointer(password_ms.data))
	QNetworkProxy_SetPassword(this.h, password_ms)
}

func (this *QNetworkProxy) Password() string {
	var _ms struct_miqt_string = QNetworkProxy_Password(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkProxy) SetHostName(hostName string) {
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	QNetworkProxy_SetHostName(this.h, hostName_ms)
}

func (this *QNetworkProxy) HostName() string {
	var _ms struct_miqt_string = QNetworkProxy_HostName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkProxy) SetPort(port uint16) {
	QNetworkProxy_SetPort(this.h, (uint16_t)(port))
}

func (this *QNetworkProxy) Port() uint16 {
	return (uint16)(QNetworkProxy_Port(this.h))
}

func QNetworkProxy_SetApplicationProxy(proxy *QNetworkProxy) {
	QNetworkProxy_SetApplicationProxy(proxy.cPointer())
}

func QNetworkProxy_ApplicationProxy() *QNetworkProxy {
	_goptr := newQNetworkProxy(QNetworkProxy_ApplicationProxy())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkProxy) Headers() *QHttpHeaders {
	_goptr := newQHttpHeaders(QNetworkProxy_Headers(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkProxy) SetHeaders(newHeaders *QHttpHeaders) {
	QNetworkProxy_SetHeaders(this.h, newHeaders.cPointer())
}

func (this *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkProxy_Header(this.h, (int)(header))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkProxy) SetHeader(header QNetworkRequest__KnownHeaders, value *qt.QVariant) {
	QNetworkProxy_SetHeader(this.h, (int)(header), (*QVariant)(value.UnsafePointer()))
}

func (this *QNetworkProxy) HasRawHeader(headerName []byte) bool {
	headerName_alias := struct_miqt_string{}
	headerName_alias.data = (char)(unsafe.Pointer(&headerName[0]))
	headerName_alias.len = size_t(len(headerName))
	return (bool)(QNetworkProxy_HasRawHeader(this.h, headerName_alias))
}

func (this *QNetworkProxy) RawHeaderList() [][]byte {
	var _ma struct_miqt_array = QNetworkProxy_RawHeaderList(this.h)
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

func (this *QNetworkProxy) RawHeader(headerName []byte) []byte {
	headerName_alias := struct_miqt_string{}
	headerName_alias.data = (char)(unsafe.Pointer(&headerName[0]))
	headerName_alias.len = size_t(len(headerName))
	var _bytearray struct_miqt_string = QNetworkProxy_RawHeader(this.h, headerName_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkProxy) SetRawHeader(headerName []byte, value []byte) {
	headerName_alias := struct_miqt_string{}
	headerName_alias.data = (char)(unsafe.Pointer(&headerName[0]))
	headerName_alias.len = size_t(len(headerName))
	value_alias := struct_miqt_string{}
	value_alias.data = (char)(unsafe.Pointer(&value[0]))
	value_alias.len = size_t(len(value))
	QNetworkProxy_SetRawHeader(this.h, headerName_alias, value_alias)
}

type QNetworkProxyFactory struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkProxyFactory constructs a new QNetworkProxyFactory object.
func NewQNetworkProxyFactory() *QNetworkProxyFactory {
	ret := newQNetworkProxyFactory(QNetworkProxyFactory_new())
	ret.isSubclass = true
	return ret
}

func (this *QNetworkProxyFactory) QueryProxy(query *QNetworkProxyQuery) []QNetworkProxy {
	var _ma struct_miqt_array = QNetworkProxyFactory_QueryProxy(this.h, query.cPointer())
	_ret := make([]QNetworkProxy, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkProxy)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkProxy(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QNetworkProxyFactory_UsesSystemConfiguration() bool {
	return (bool)(QNetworkProxyFactory_UsesSystemConfiguration())
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	QNetworkProxyFactory_SetUseSystemConfiguration((bool)(enable))
}

func QNetworkProxyFactory_SetApplicationProxyFactory(factory *QNetworkProxyFactory) {
	QNetworkProxyFactory_SetApplicationProxyFactory(factory.cPointer())
}

func QNetworkProxyFactory_ProxyForQuery(query *QNetworkProxyQuery) []QNetworkProxy {
	var _ma struct_miqt_array = QNetworkProxyFactory_ProxyForQuery(query.cPointer())
	_ret := make([]QNetworkProxy, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkProxy)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkProxy(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QNetworkProxyFactory_SystemProxyForQuery() []QNetworkProxy {
	var _ma struct_miqt_array = QNetworkProxyFactory_SystemProxyForQuery()
	_ret := make([]QNetworkProxy, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkProxy)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkProxy(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QNetworkProxyFactory) OperatorAssign(param1 *QNetworkProxyFactory) {
	QNetworkProxyFactory_OperatorAssign(this.h, param1.cPointer())
}

func QNetworkProxyFactory_SystemProxyForQuery1(query *QNetworkProxyQuery) []QNetworkProxy {
	var _ma struct_miqt_array = QNetworkProxyFactory_SystemProxyForQuery1(query.cPointer())
	_ret := make([]QNetworkProxy, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkProxy)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkProxy(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}
