package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QNetworkRequestFactory struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkRequestFactory constructs a new QNetworkRequestFactory object.
func NewQNetworkRequestFactory() *QNetworkRequestFactory {

	ret := newQNetworkRequestFactory(QNetworkRequestFactory_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkRequestFactory2 constructs a new QNetworkRequestFactory object.
func NewQNetworkRequestFactory2(baseUrl *qt.QUrl) *QNetworkRequestFactory {

	ret := newQNetworkRequestFactory(QNetworkRequestFactory_new2((*QUrl)(baseUrl.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQNetworkRequestFactory3 constructs a new QNetworkRequestFactory object.
func NewQNetworkRequestFactory3(other *QNetworkRequestFactory) *QNetworkRequestFactory {

	ret := newQNetworkRequestFactory(QNetworkRequestFactory_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkRequestFactory) OperatorAssign(other *QNetworkRequestFactory) {
	QNetworkRequestFactory_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkRequestFactory) Swap(other *QNetworkRequestFactory) {
	QNetworkRequestFactory_Swap(this.h, other.cPointer())
}

func (this *QNetworkRequestFactory) BaseUrl() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QNetworkRequestFactory_BaseUrl(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) SetBaseUrl(url *qt.QUrl) {
	QNetworkRequestFactory_SetBaseUrl(this.h, (*QUrl)(url.UnsafePointer()))
}

func (this *QNetworkRequestFactory) SslConfiguration() *QSslConfiguration {
	_goptr := newQSslConfiguration(QNetworkRequestFactory_SslConfiguration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) SetSslConfiguration(configuration *QSslConfiguration) {
	QNetworkRequestFactory_SetSslConfiguration(this.h, configuration.cPointer())
}

func (this *QNetworkRequestFactory) CreateRequest() *QNetworkRequest {
	_goptr := newQNetworkRequest(QNetworkRequestFactory_CreateRequest(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) CreateRequestWithQuery(query *qt.QUrlQuery) *QNetworkRequest {
	_goptr := newQNetworkRequest(QNetworkRequestFactory_CreateRequestWithQuery(this.h, (*QUrlQuery)(query.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) CreateRequestWithPath(path string) *QNetworkRequest {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	_goptr := newQNetworkRequest(QNetworkRequestFactory_CreateRequestWithPath(this.h, path_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) CreateRequest2(path string, query *qt.QUrlQuery) *QNetworkRequest {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	_goptr := newQNetworkRequest(QNetworkRequestFactory_CreateRequest2(this.h, path_ms, (*QUrlQuery)(query.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) SetCommonHeaders(headers *QHttpHeaders) {
	QNetworkRequestFactory_SetCommonHeaders(this.h, headers.cPointer())
}

func (this *QNetworkRequestFactory) CommonHeaders() *QHttpHeaders {
	_goptr := newQHttpHeaders(QNetworkRequestFactory_CommonHeaders(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) ClearCommonHeaders() {
	QNetworkRequestFactory_ClearCommonHeaders(this.h)
}

func (this *QNetworkRequestFactory) BearerToken() []byte {
	var _bytearray struct_miqt_string = QNetworkRequestFactory_BearerToken(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkRequestFactory) SetBearerToken(token []byte) {
	token_alias := struct_miqt_string{}
	token_alias.data = (char)(unsafe.Pointer(&token[0]))
	token_alias.len = size_t(len(token))
	QNetworkRequestFactory_SetBearerToken(this.h, token_alias)
}

func (this *QNetworkRequestFactory) ClearBearerToken() {
	QNetworkRequestFactory_ClearBearerToken(this.h)
}

func (this *QNetworkRequestFactory) UserName() string {
	var _ms struct_miqt_string = QNetworkRequestFactory_UserName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkRequestFactory) SetUserName(userName string) {
	userName_ms := struct_miqt_string{}
	userName_ms.data = CString(userName)
	userName_ms.len = size_t(len(userName))
	defer free(unsafe.Pointer(userName_ms.data))
	QNetworkRequestFactory_SetUserName(this.h, userName_ms)
}

func (this *QNetworkRequestFactory) ClearUserName() {
	QNetworkRequestFactory_ClearUserName(this.h)
}

func (this *QNetworkRequestFactory) Password() string {
	var _ms struct_miqt_string = QNetworkRequestFactory_Password(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkRequestFactory) SetPassword(password string) {
	password_ms := struct_miqt_string{}
	password_ms.data = CString(password)
	password_ms.len = size_t(len(password))
	defer free(unsafe.Pointer(password_ms.data))
	QNetworkRequestFactory_SetPassword(this.h, password_ms)
}

func (this *QNetworkRequestFactory) ClearPassword() {
	QNetworkRequestFactory_ClearPassword(this.h)
}

func (this *QNetworkRequestFactory) QueryParameters() *qt.QUrlQuery {
	_goptr := qt.UnsafeNewQUrlQuery(unsafe.Pointer(QNetworkRequestFactory_QueryParameters(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) SetQueryParameters(query *qt.QUrlQuery) {
	QNetworkRequestFactory_SetQueryParameters(this.h, (*QUrlQuery)(query.UnsafePointer()))
}

func (this *QNetworkRequestFactory) ClearQueryParameters() {
	QNetworkRequestFactory_ClearQueryParameters(this.h)
}

func (this *QNetworkRequestFactory) SetPriority(priority QNetworkRequest__Priority) {
	QNetworkRequestFactory_SetPriority(this.h, (int)(priority))
}

func (this *QNetworkRequestFactory) Priority() QNetworkRequest__Priority {
	return (QNetworkRequest__Priority)(QNetworkRequestFactory_Priority(this.h))
}

func (this *QNetworkRequestFactory) Attribute(attribute QNetworkRequest__Attribute) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkRequestFactory_Attribute(this.h, (int)(attribute))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) Attribute2(attribute QNetworkRequest__Attribute, defaultValue *qt.QVariant) *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QNetworkRequestFactory_Attribute2(this.h, (int)(attribute), (*QVariant)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkRequestFactory) SetAttribute(attribute QNetworkRequest__Attribute, value *qt.QVariant) {
	QNetworkRequestFactory_SetAttribute(this.h, (int)(attribute), (*QVariant)(value.UnsafePointer()))
}

func (this *QNetworkRequestFactory) ClearAttribute(attribute QNetworkRequest__Attribute) {
	QNetworkRequestFactory_ClearAttribute(this.h, (int)(attribute))
}

func (this *QNetworkRequestFactory) ClearAttributes() {
	QNetworkRequestFactory_ClearAttributes(this.h)
}
