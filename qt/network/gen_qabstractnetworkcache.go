package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QNetworkCacheMetaData struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkCacheMetaData constructs a new QNetworkCacheMetaData object.
func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	ret := newQNetworkCacheMetaData(QNetworkCacheMetaData_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkCacheMetaData2 constructs a new QNetworkCacheMetaData object.
func NewQNetworkCacheMetaData2(other *QNetworkCacheMetaData) *QNetworkCacheMetaData {
	ret := newQNetworkCacheMetaData(QNetworkCacheMetaData_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkCacheMetaData) OperatorAssign(other *QNetworkCacheMetaData) {
	QNetworkCacheMetaData_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkCacheMetaData) Swap(other *QNetworkCacheMetaData) {
	QNetworkCacheMetaData_Swap(this.h, other.cPointer())
}

func (this *QNetworkCacheMetaData) OperatorEqual(other *QNetworkCacheMetaData) bool {
	return (bool)(QNetworkCacheMetaData_OperatorEqual(this.h, other.cPointer()))
}

func (this *QNetworkCacheMetaData) OperatorNotEqual(other *QNetworkCacheMetaData) bool {
	return (bool)(QNetworkCacheMetaData_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QNetworkCacheMetaData) IsValid() bool {
	return (bool)(QNetworkCacheMetaData_IsValid(this.h))
}

func (this *QNetworkCacheMetaData) Url() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QNetworkCacheMetaData_Url(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkCacheMetaData) SetUrl(url *qt.QUrl) {
	QNetworkCacheMetaData_SetUrl(this.h, (*QUrl)(url.UnsafePointer()))
}

func (this *QNetworkCacheMetaData) RawHeaders() RawHeaderList {
	xxxxxxxxx
}

func (this *QNetworkCacheMetaData) SetRawHeaders(headers *RawHeaderList) {
	QNetworkCacheMetaData_SetRawHeaders(this.h, headers)
}

func (this *QNetworkCacheMetaData) Headers() *QHttpHeaders {
	_goptr := newQHttpHeaders(QNetworkCacheMetaData_Headers(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkCacheMetaData) SetHeaders(headers *QHttpHeaders) {
	QNetworkCacheMetaData_SetHeaders(this.h, headers.cPointer())
}

func (this *QNetworkCacheMetaData) LastModified() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QNetworkCacheMetaData_LastModified(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkCacheMetaData) SetLastModified(dateTime *qt.QDateTime) {
	QNetworkCacheMetaData_SetLastModified(this.h, (*QDateTime)(dateTime.UnsafePointer()))
}

func (this *QNetworkCacheMetaData) ExpirationDate() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QNetworkCacheMetaData_ExpirationDate(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkCacheMetaData) SetExpirationDate(dateTime *qt.QDateTime) {
	QNetworkCacheMetaData_SetExpirationDate(this.h, (*QDateTime)(dateTime.UnsafePointer()))
}

func (this *QNetworkCacheMetaData) SaveToDisk() bool {
	return (bool)(QNetworkCacheMetaData_SaveToDisk(this.h))
}

func (this *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	QNetworkCacheMetaData_SetSaveToDisk(this.h, (bool)(allow))
}

func (this *QNetworkCacheMetaData) Attributes() AttributesMap {
	xxxxxxxxx
}

func (this *QNetworkCacheMetaData) SetAttributes(attributes *AttributesMap) {
	QNetworkCacheMetaData_SetAttributes(this.h, attributes)
}

type QAbstractNetworkCache struct {
	h          uintptr
	isSubclass bool
}

func (this *QAbstractNetworkCache) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAbstractNetworkCache_MetaObject(this.h)))
}

func (this *QAbstractNetworkCache) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractNetworkCache_Metacast(this.h, param1_Cstring))
}

func QAbstractNetworkCache_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractNetworkCache_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractNetworkCache) MetaData(url *qt.QUrl) *QNetworkCacheMetaData {
	_goptr := newQNetworkCacheMetaData(QAbstractNetworkCache_MetaData(this.h, (*QUrl)(url.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractNetworkCache) UpdateMetaData(metaData *QNetworkCacheMetaData) {
	QAbstractNetworkCache_UpdateMetaData(this.h, metaData.cPointer())
}

func (this *QAbstractNetworkCache) Data(url *qt.QUrl) *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QAbstractNetworkCache_Data(this.h, (*QUrl)(url.UnsafePointer()))))
}

func (this *QAbstractNetworkCache) Remove(url *qt.QUrl) bool {
	return (bool)(QAbstractNetworkCache_Remove(this.h, (*QUrl)(url.UnsafePointer())))
}

func (this *QAbstractNetworkCache) CacheSize() int64 {
	return (int64)(QAbstractNetworkCache_CacheSize(this.h))
}

func (this *QAbstractNetworkCache) Prepare(metaData *QNetworkCacheMetaData) *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QAbstractNetworkCache_Prepare(this.h, metaData.cPointer())))
}

func (this *QAbstractNetworkCache) Insert(device *qt.QIODevice) {
	QAbstractNetworkCache_Insert(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QAbstractNetworkCache) Clear() {
	QAbstractNetworkCache_Clear(this.h)
}

func QAbstractNetworkCache_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractNetworkCache_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractNetworkCache_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractNetworkCache_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
