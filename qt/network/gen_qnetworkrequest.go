package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QNetworkRequest__KnownHeaders int

const (
	QNetworkRequest__ContentTypeHeader        QNetworkRequest__KnownHeaders = 0
	QNetworkRequest__ContentLengthHeader      QNetworkRequest__KnownHeaders = 1
	QNetworkRequest__LocationHeader           QNetworkRequest__KnownHeaders = 2
	QNetworkRequest__LastModifiedHeader       QNetworkRequest__KnownHeaders = 3
	QNetworkRequest__CookieHeader             QNetworkRequest__KnownHeaders = 4
	QNetworkRequest__SetCookieHeader          QNetworkRequest__KnownHeaders = 5
	QNetworkRequest__ContentDispositionHeader QNetworkRequest__KnownHeaders = 6
	QNetworkRequest__UserAgentHeader          QNetworkRequest__KnownHeaders = 7
	QNetworkRequest__ServerHeader             QNetworkRequest__KnownHeaders = 8
	QNetworkRequest__IfModifiedSinceHeader    QNetworkRequest__KnownHeaders = 9
	QNetworkRequest__ETagHeader               QNetworkRequest__KnownHeaders = 10
	QNetworkRequest__IfMatchHeader            QNetworkRequest__KnownHeaders = 11
	QNetworkRequest__IfNoneMatchHeader        QNetworkRequest__KnownHeaders = 12
	QNetworkRequest__NumKnownHeaders          QNetworkRequest__KnownHeaders = 13
)

type QNetworkRequest__Attribute int

const (
	QNetworkRequest__HttpStatusCodeAttribute                      QNetworkRequest__Attribute = 0
	QNetworkRequest__HttpReasonPhraseAttribute                    QNetworkRequest__Attribute = 1
	QNetworkRequest__RedirectionTargetAttribute                   QNetworkRequest__Attribute = 2
	QNetworkRequest__ConnectionEncryptedAttribute                 QNetworkRequest__Attribute = 3
	QNetworkRequest__CacheLoadControlAttribute                    QNetworkRequest__Attribute = 4
	QNetworkRequest__CacheSaveControlAttribute                    QNetworkRequest__Attribute = 5
	QNetworkRequest__SourceIsFromCacheAttribute                   QNetworkRequest__Attribute = 6
	QNetworkRequest__DoNotBufferUploadDataAttribute               QNetworkRequest__Attribute = 7
	QNetworkRequest__HttpPipeliningAllowedAttribute               QNetworkRequest__Attribute = 8
	QNetworkRequest__HttpPipeliningWasUsedAttribute               QNetworkRequest__Attribute = 9
	QNetworkRequest__CustomVerbAttribute                          QNetworkRequest__Attribute = 10
	QNetworkRequest__CookieLoadControlAttribute                   QNetworkRequest__Attribute = 11
	QNetworkRequest__AuthenticationReuseAttribute                 QNetworkRequest__Attribute = 12
	QNetworkRequest__CookieSaveControlAttribute                   QNetworkRequest__Attribute = 13
	QNetworkRequest__MaximumDownloadBufferSizeAttribute           QNetworkRequest__Attribute = 14
	QNetworkRequest__DownloadBufferAttribute                      QNetworkRequest__Attribute = 15
	QNetworkRequest__SynchronousRequestAttribute                  QNetworkRequest__Attribute = 16
	QNetworkRequest__BackgroundRequestAttribute                   QNetworkRequest__Attribute = 17
	QNetworkRequest__EmitAllUploadProgressSignalsAttribute        QNetworkRequest__Attribute = 18
	QNetworkRequest__Http2AllowedAttribute                        QNetworkRequest__Attribute = 19
	QNetworkRequest__Http2WasUsedAttribute                        QNetworkRequest__Attribute = 20
	QNetworkRequest__OriginalContentLengthAttribute               QNetworkRequest__Attribute = 21
	QNetworkRequest__RedirectPolicyAttribute                      QNetworkRequest__Attribute = 22
	QNetworkRequest__Http2DirectAttribute                         QNetworkRequest__Attribute = 23
	QNetworkRequest__ResourceTypeAttribute                        QNetworkRequest__Attribute = 24
	QNetworkRequest__AutoDeleteReplyOnFinishAttribute             QNetworkRequest__Attribute = 25
	QNetworkRequest__ConnectionCacheExpiryTimeoutSecondsAttribute QNetworkRequest__Attribute = 26
	QNetworkRequest__Http2CleartextAllowedAttribute               QNetworkRequest__Attribute = 27
	QNetworkRequest__UseCredentialsAttribute                      QNetworkRequest__Attribute = 28
	QNetworkRequest__FullLocalServerNameAttribute                 QNetworkRequest__Attribute = 29
	QNetworkRequest__User                                         QNetworkRequest__Attribute = 1000
	QNetworkRequest__UserMax                                      QNetworkRequest__Attribute = 32767
)

type QNetworkRequest__CacheLoadControl int

const (
	QNetworkRequest__AlwaysNetwork QNetworkRequest__CacheLoadControl = 0
	QNetworkRequest__PreferNetwork QNetworkRequest__CacheLoadControl = 1
	QNetworkRequest__PreferCache   QNetworkRequest__CacheLoadControl = 2
	QNetworkRequest__AlwaysCache   QNetworkRequest__CacheLoadControl = 3
)

type QNetworkRequest__LoadControl int

const (
	QNetworkRequest__Automatic QNetworkRequest__LoadControl = 0
	QNetworkRequest__Manual    QNetworkRequest__LoadControl = 1
)

type QNetworkRequest__Priority int

const (
	QNetworkRequest__HighPriority   QNetworkRequest__Priority = 1
	QNetworkRequest__NormalPriority QNetworkRequest__Priority = 3
	QNetworkRequest__LowPriority    QNetworkRequest__Priority = 5
)

type QNetworkRequest__RedirectPolicy int

const (
	QNetworkRequest__ManualRedirectPolicy       QNetworkRequest__RedirectPolicy = 0
	QNetworkRequest__NoLessSafeRedirectPolicy   QNetworkRequest__RedirectPolicy = 1
	QNetworkRequest__SameOriginRedirectPolicy   QNetworkRequest__RedirectPolicy = 2
	QNetworkRequest__UserVerifiedRedirectPolicy QNetworkRequest__RedirectPolicy = 3
)

type QNetworkRequest__TransferTimeoutConstant int

const (
	QNetworkRequest__DefaultTransferTimeoutConstant QNetworkRequest__TransferTimeoutConstant = 30000
)

type QNetworkRequest struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkRequest constructs a new QNetworkRequest object.
func NewQNetworkRequest() *QNetworkRequest {

	ret := newQNetworkRequest(QNetworkRequest_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkRequest2 constructs a new QNetworkRequest object.
func NewQNetworkRequest2(url *qt.QUrl) *QNetworkRequest {

	ret := newQNetworkRequest(QNetworkRequest_new2((*QUrl)(url.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQNetworkRequest3 constructs a new QNetworkRequest object.
func NewQNetworkRequest3(other *QNetworkRequest) *QNetworkRequest {

	ret := newQNetworkRequest(QNetworkRequest_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkRequest) OperatorAssign(other *QNetworkRequest) {
	QNetworkRequest_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkRequest) Swap(other *QNetworkRequest) {
	QNetworkRequest_Swap(this.h, other.cPointer())
}

func (this *QNetworkRequest) OperatorEqual(other *QNetworkRequest) bool {
	return (bool)(QNetworkRequest_OperatorEqual(this.h, other.cPointer()))
}

func (this *QNetworkRequest) OperatorNotEqual(other *QNetworkRequest) bool {
	return (bool)(QNetworkRequest_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QNetworkRequest) Url() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QNetworkRequest_Url(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequest) SetUrl(url *qt.QUrl) {
	QNetworkRequest_SetUrl(this.h, (*QUrl)(url.UnsafePointer()))
}

func (this *QNetworkRequest) Headers() *QHttpHeaders {
	_goptr := newQHttpHeaders(QNetworkRequest_Headers(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequest) SetHeaders(newHeaders *QHttpHeaders) {
	QNetworkRequest_SetHeaders(this.h, newHeaders.cPointer())
}

func (this *QNetworkRequest) Header(header KnownHeaders) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkRequest_Header(this.h, header)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequest) SetHeader(header KnownHeaders, value *qt.QVariant) {
	QNetworkRequest_SetHeader(this.h, header, (*QVariant)(value.UnsafePointer()))
}

func (this *QNetworkRequest) HasRawHeader(headerName qt.QAnyStringView) bool {
	return (bool)(QNetworkRequest_HasRawHeader(this.h, (*QAnyStringView)(headerName.UnsafePointer())))
}

func (this *QNetworkRequest) RawHeaderList() [][]byte {
	var _ma struct_miqt_array = QNetworkRequest_RawHeaderList(this.h)
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

func (this *QNetworkRequest) RawHeader(headerName qt.QAnyStringView) []byte {
	var _bytearray struct_miqt_string = QNetworkRequest_RawHeader(this.h, (*QAnyStringView)(headerName.UnsafePointer()))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkRequest) SetRawHeader(headerName []byte, value []byte) {
	headerName_alias := struct_miqt_string{}
	headerName_alias.data = (char)(unsafe.Pointer(&headerName[0]))
	headerName_alias.len = size_t(len(headerName))
	value_alias := struct_miqt_string{}
	value_alias.data = (char)(unsafe.Pointer(&value[0]))
	value_alias.len = size_t(len(value))
	QNetworkRequest_SetRawHeader(this.h, headerName_alias, value_alias)
}

func (this *QNetworkRequest) Attribute(code Attribute) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkRequest_Attribute(this.h, code)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequest) SetAttribute(code Attribute, value *qt.QVariant) {
	QNetworkRequest_SetAttribute(this.h, code, (*QVariant)(value.UnsafePointer()))
}

func (this *QNetworkRequest) SslConfiguration() *QSslConfiguration {
	_goptr := newQSslConfiguration(QNetworkRequest_SslConfiguration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequest) SetSslConfiguration(configuration *QSslConfiguration) {
	QNetworkRequest_SetSslConfiguration(this.h, configuration.cPointer())
}

func (this *QNetworkRequest) SetOriginatingObject(object *qt.QObject) {
	QNetworkRequest_SetOriginatingObject(this.h, (*QObject)(object.UnsafePointer()))
}

func (this *QNetworkRequest) OriginatingObject() *qt.QObject {
	return qt.UnsafeNewQObject(unsafe.Pointer(QNetworkRequest_OriginatingObject(this.h)))
}

func (this *QNetworkRequest) Priority() Priority {
	xxxxxxxxx
}

func (this *QNetworkRequest) SetPriority(priority Priority) {
	QNetworkRequest_SetPriority(this.h, priority)
}

func (this *QNetworkRequest) MaximumRedirectsAllowed() int {
	return (int)(QNetworkRequest_MaximumRedirectsAllowed(this.h))
}

func (this *QNetworkRequest) SetMaximumRedirectsAllowed(maximumRedirectsAllowed int) {
	QNetworkRequest_SetMaximumRedirectsAllowed(this.h, (int)(maximumRedirectsAllowed))
}

func (this *QNetworkRequest) PeerVerifyName() string {
	var _ms struct_miqt_string = QNetworkRequest_PeerVerifyName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkRequest) SetPeerVerifyName(peerName string) {
	peerName_ms := struct_miqt_string{}
	peerName_ms.data = CString(peerName)
	peerName_ms.len = size_t(len(peerName))
	defer free(unsafe.Pointer(peerName_ms.data))
	QNetworkRequest_SetPeerVerifyName(this.h, peerName_ms)
}

func (this *QNetworkRequest) Http1Configuration() *QHttp1Configuration {
	_goptr := newQHttp1Configuration(QNetworkRequest_Http1Configuration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequest) SetHttp1Configuration(configuration *QHttp1Configuration) {
	QNetworkRequest_SetHttp1Configuration(this.h, configuration.cPointer())
}

func (this *QNetworkRequest) Http2Configuration() *QHttp2Configuration {
	_goptr := newQHttp2Configuration(QNetworkRequest_Http2Configuration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequest) SetHttp2Configuration(configuration *QHttp2Configuration) {
	QNetworkRequest_SetHttp2Configuration(this.h, configuration.cPointer())
}

func (this *QNetworkRequest) DecompressedSafetyCheckThreshold() int64 {
	return (int64)(QNetworkRequest_DecompressedSafetyCheckThreshold(this.h))
}

func (this *QNetworkRequest) SetDecompressedSafetyCheckThreshold(threshold int64) {
	QNetworkRequest_SetDecompressedSafetyCheckThreshold(this.h, (longlong)(threshold))
}

func (this *QNetworkRequest) TransferTimeout() int {
	return (int)(QNetworkRequest_TransferTimeout(this.h))
}

func (this *QNetworkRequest) SetTransferTimeout(timeout int) {
	QNetworkRequest_SetTransferTimeout(this.h, (int)(timeout))
}

func (this *QNetworkRequest) SetTransferTimeout2() {
	QNetworkRequest_SetTransferTimeout2(this.h)
}

func (this *QNetworkRequest) Attribute2(code Attribute, defaultValue *qt.QVariant) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkRequest_Attribute2(this.h, code, (*QVariant)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
