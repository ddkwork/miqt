package network

import (
	"unsafe"
)

type QHostInfo__HostInfoError int

const (
	QHostInfo__NoError      QHostInfo__HostInfoError = 0
	QHostInfo__HostNotFound QHostInfo__HostInfoError = 1
	QHostInfo__UnknownError QHostInfo__HostInfoError = 2
)

type QHostInfo struct {
	h          uintptr
	isSubclass bool
}

// NewQHostInfo constructs a new QHostInfo object.
func NewQHostInfo() *QHostInfo {
	ret := newQHostInfo(QHostInfo_new())
	ret.isSubclass = true
	return ret
}

// NewQHostInfo2 constructs a new QHostInfo object.
func NewQHostInfo2(d *QHostInfo) *QHostInfo {
	ret := newQHostInfo(QHostInfo_new2(d.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQHostInfo3 constructs a new QHostInfo object.
func NewQHostInfo3(lookupId int) *QHostInfo {
	ret := newQHostInfo(QHostInfo_new3((int)(lookupId)))
	ret.isSubclass = true
	return ret
}

func (this *QHostInfo) OperatorAssign(d *QHostInfo) {
	QHostInfo_OperatorAssign(this.h, d.cPointer())
}

func (this *QHostInfo) Swap(other *QHostInfo) {
	QHostInfo_Swap(this.h, other.cPointer())
}

func (this *QHostInfo) HostName() string {
	var _ms struct_miqt_string = QHostInfo_HostName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHostInfo) SetHostName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QHostInfo_SetHostName(this.h, name_ms)
}

func (this *QHostInfo) Addresses() []QHostAddress {
	var _ma struct_miqt_array = QHostInfo_Addresses(this.h)
	_ret := make([]QHostAddress, int(_ma.len))
	_outCast := (*[0xffff]*QHostAddress)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQHostAddress(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QHostInfo) SetAddresses(addresses []QHostAddress) {
	addresses_CArray := (*[0xffff]*QHostAddress)(malloc(size_t(8 * len(addresses))))
	defer free(unsafe.Pointer(addresses_CArray))
	for i := range addresses {
		addresses_CArray[i] = addresses[i].cPointer()
	}
	addresses_ma := struct_miqt_array{len: size_t(len(addresses)), data: unsafe.Pointer(addresses_CArray)}
	QHostInfo_SetAddresses(this.h, addresses_ma)
}

func (this *QHostInfo) Error() HostInfoError {
	xxxxxxxxx
}

func (this *QHostInfo) SetError(error HostInfoError) {
	QHostInfo_SetError(this.h, error)
}

func (this *QHostInfo) ErrorString() string {
	var _ms struct_miqt_string = QHostInfo_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHostInfo) SetErrorString(errorString string) {
	errorString_ms := struct_miqt_string{}
	errorString_ms.data = CString(errorString)
	errorString_ms.len = size_t(len(errorString))
	defer free(unsafe.Pointer(errorString_ms.data))
	QHostInfo_SetErrorString(this.h, errorString_ms)
}

func (this *QHostInfo) SetLookupId(id int) {
	QHostInfo_SetLookupId(this.h, (int)(id))
}

func (this *QHostInfo) LookupId() int {
	return (int)(QHostInfo_LookupId(this.h))
}

func QHostInfo_AbortHostLookup(lookupId int) {
	QHostInfo_AbortHostLookup((int)(lookupId))
}

func QHostInfo_FromName(name string) *QHostInfo {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQHostInfo(QHostInfo_FromName(name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QHostInfo_LocalHostName() string {
	var _ms struct_miqt_string = QHostInfo_LocalHostName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QHostInfo_LocalDomainName() string {
	var _ms struct_miqt_string = QHostInfo_LocalDomainName()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
