package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QNetworkCookie__RawForm int

const (
	QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = 0
	QNetworkCookie__Full             QNetworkCookie__RawForm = 1
)

type QNetworkCookie__SameSite int

const (
	QNetworkCookie__Default QNetworkCookie__SameSite = 0
	QNetworkCookie__None    QNetworkCookie__SameSite = 1
	QNetworkCookie__Lax     QNetworkCookie__SameSite = 2
	QNetworkCookie__Strict  QNetworkCookie__SameSite = 3
)

type QNetworkCookie struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkCookie constructs a new QNetworkCookie object.
func NewQNetworkCookie() *QNetworkCookie {

	ret := newQNetworkCookie(QNetworkCookie_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkCookie2 constructs a new QNetworkCookie object.
func NewQNetworkCookie2(other *QNetworkCookie) *QNetworkCookie {

	ret := newQNetworkCookie(QNetworkCookie_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQNetworkCookie3 constructs a new QNetworkCookie object.
func NewQNetworkCookie3(name []byte) *QNetworkCookie {
	name_alias := struct_miqt_string{}
	name_alias.data = (char)(unsafe.Pointer(&name[0]))
	name_alias.len = size_t(len(name))

	ret := newQNetworkCookie(QNetworkCookie_new3(name_alias))
	ret.isSubclass = true
	return ret
}

// NewQNetworkCookie4 constructs a new QNetworkCookie object.
func NewQNetworkCookie4(name []byte, value []byte) *QNetworkCookie {
	name_alias := struct_miqt_string{}
	name_alias.data = (char)(unsafe.Pointer(&name[0]))
	name_alias.len = size_t(len(name))
	value_alias := struct_miqt_string{}
	value_alias.data = (char)(unsafe.Pointer(&value[0]))
	value_alias.len = size_t(len(value))

	ret := newQNetworkCookie(QNetworkCookie_new4(name_alias, value_alias))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkCookie) OperatorAssign(other *QNetworkCookie) {
	QNetworkCookie_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkCookie) Swap(other *QNetworkCookie) {
	QNetworkCookie_Swap(this.h, other.cPointer())
}

func (this *QNetworkCookie) OperatorEqual(other *QNetworkCookie) bool {
	return (bool)(QNetworkCookie_OperatorEqual(this.h, other.cPointer()))
}

func (this *QNetworkCookie) OperatorNotEqual(other *QNetworkCookie) bool {
	return (bool)(QNetworkCookie_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QNetworkCookie) IsSecure() bool {
	return (bool)(QNetworkCookie_IsSecure(this.h))
}

func (this *QNetworkCookie) SetSecure(enable bool) {
	QNetworkCookie_SetSecure(this.h, (bool)(enable))
}

func (this *QNetworkCookie) IsHttpOnly() bool {
	return (bool)(QNetworkCookie_IsHttpOnly(this.h))
}

func (this *QNetworkCookie) SetHttpOnly(enable bool) {
	QNetworkCookie_SetHttpOnly(this.h, (bool)(enable))
}

func (this *QNetworkCookie) SameSitePolicy() SameSite {
	xxxxxxxxx
}

func (this *QNetworkCookie) SetSameSitePolicy(sameSite SameSite) {
	QNetworkCookie_SetSameSitePolicy(this.h, sameSite)
}

func (this *QNetworkCookie) IsSessionCookie() bool {
	return (bool)(QNetworkCookie_IsSessionCookie(this.h))
}

func (this *QNetworkCookie) ExpirationDate() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QNetworkCookie_ExpirationDate(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkCookie) SetExpirationDate(date *qt.QDateTime) {
	QNetworkCookie_SetExpirationDate(this.h, (*QDateTime)(date.UnsafePointer()))
}

func (this *QNetworkCookie) Domain() string {
	var _ms struct_miqt_string = QNetworkCookie_Domain(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkCookie) SetDomain(domain string) {
	domain_ms := struct_miqt_string{}
	domain_ms.data = CString(domain)
	domain_ms.len = size_t(len(domain))
	defer free(unsafe.Pointer(domain_ms.data))
	QNetworkCookie_SetDomain(this.h, domain_ms)
}

func (this *QNetworkCookie) Path() string {
	var _ms struct_miqt_string = QNetworkCookie_Path(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkCookie) SetPath(path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QNetworkCookie_SetPath(this.h, path_ms)
}

func (this *QNetworkCookie) Name() []byte {
	var _bytearray struct_miqt_string = QNetworkCookie_Name(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkCookie) SetName(cookieName []byte) {
	cookieName_alias := struct_miqt_string{}
	cookieName_alias.data = (char)(unsafe.Pointer(&cookieName[0]))
	cookieName_alias.len = size_t(len(cookieName))
	QNetworkCookie_SetName(this.h, cookieName_alias)
}

func (this *QNetworkCookie) Value() []byte {
	var _bytearray struct_miqt_string = QNetworkCookie_Value(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkCookie) SetValue(value []byte) {
	value_alias := struct_miqt_string{}
	value_alias.data = (char)(unsafe.Pointer(&value[0]))
	value_alias.len = size_t(len(value))
	QNetworkCookie_SetValue(this.h, value_alias)
}

func (this *QNetworkCookie) ToRawForm() []byte {
	var _bytearray struct_miqt_string = QNetworkCookie_ToRawForm(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkCookie) HasSameIdentifier(other *QNetworkCookie) bool {
	return (bool)(QNetworkCookie_HasSameIdentifier(this.h, other.cPointer()))
}

func (this *QNetworkCookie) Normalize(url *qt.QUrl) {
	QNetworkCookie_Normalize(this.h, (*QUrl)(url.UnsafePointer()))
}

func QNetworkCookie_ParseCookies(cookieString qt.QByteArrayView) []QNetworkCookie {
	var _ma struct_miqt_array = QNetworkCookie_ParseCookies((*QByteArrayView)(cookieString.UnsafePointer()))
	_ret := make([]QNetworkCookie, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkCookie)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkCookie(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QNetworkCookie) ToRawForm1(form RawForm) []byte {
	var _bytearray struct_miqt_string = QNetworkCookie_ToRawForm1(this.h, form)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}
