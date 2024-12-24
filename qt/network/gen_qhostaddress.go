package network

import (
	"unsafe"
)

type QHostAddress__SpecialAddress int

const (
	QHostAddress__Null          QHostAddress__SpecialAddress = 0
	QHostAddress__Broadcast     QHostAddress__SpecialAddress = 1
	QHostAddress__LocalHost     QHostAddress__SpecialAddress = 2
	QHostAddress__LocalHostIPv6 QHostAddress__SpecialAddress = 3
	QHostAddress__Any           QHostAddress__SpecialAddress = 4
	QHostAddress__AnyIPv6       QHostAddress__SpecialAddress = 5
	QHostAddress__AnyIPv4       QHostAddress__SpecialAddress = 6
)

type QHostAddress__ConversionModeFlag int

const (
	QHostAddress__ConvertV4MappedToIPv4     QHostAddress__ConversionModeFlag = 1
	QHostAddress__ConvertV4CompatToIPv4     QHostAddress__ConversionModeFlag = 2
	QHostAddress__ConvertUnspecifiedAddress QHostAddress__ConversionModeFlag = 4
	QHostAddress__ConvertLocalHost          QHostAddress__ConversionModeFlag = 8
	QHostAddress__TolerantConversion        QHostAddress__ConversionModeFlag = 255
	QHostAddress__StrictConversion          QHostAddress__ConversionModeFlag = 0
)

type QIPv6Address struct {
	h          uintptr
	isSubclass bool
}

// NewQIPv6Address constructs a new QIPv6Address object.
func NewQIPv6Address() *QIPv6Address {
	ret := newQIPv6Address(QIPv6Address_new())
	ret.isSubclass = true
	return ret
}

// NewQIPv6Address2 constructs a new QIPv6Address object.
func NewQIPv6Address2(param1 *QIPv6Address) *QIPv6Address {
	ret := newQIPv6Address(QIPv6Address_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QIPv6Address) OperatorSubscript(index int) byte {
	return (byte)(QIPv6Address_OperatorSubscript(this.h, (int)(index)))
}

type QHostAddress struct {
	h          uintptr
	isSubclass bool
}

// NewQHostAddress constructs a new QHostAddress object.
func NewQHostAddress() *QHostAddress {
	ret := newQHostAddress(QHostAddress_new())
	ret.isSubclass = true
	return ret
}

// NewQHostAddress2 constructs a new QHostAddress object.
func NewQHostAddress2(ip4Addr uint) *QHostAddress {
	ret := newQHostAddress(QHostAddress_new2((uint)(ip4Addr)))
	ret.isSubclass = true
	return ret
}

// NewQHostAddress3 constructs a new QHostAddress object.
func NewQHostAddress3(ip6Addr *byte) *QHostAddress {
	ret := newQHostAddress(QHostAddress_new3((*uchar)(unsafe.Pointer(ip6Addr))))
	ret.isSubclass = true
	return ret
}

// NewQHostAddress4 constructs a new QHostAddress object.
func NewQHostAddress4(ip6Addr *QIPv6Address) *QHostAddress {
	ret := newQHostAddress(QHostAddress_new4(ip6Addr.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQHostAddress5 constructs a new QHostAddress object.
func NewQHostAddress5(address string) *QHostAddress {
	address_ms := struct_miqt_string{}
	address_ms.data = CString(address)
	address_ms.len = size_t(len(address))
	defer free(unsafe.Pointer(address_ms.data))

	ret := newQHostAddress(QHostAddress_new5(address_ms))
	ret.isSubclass = true
	return ret
}

// NewQHostAddress6 constructs a new QHostAddress object.
func NewQHostAddress6(copyVal *QHostAddress) *QHostAddress {
	ret := newQHostAddress(QHostAddress_new6(copyVal.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQHostAddress7 constructs a new QHostAddress object.
func NewQHostAddress7(address SpecialAddress) *QHostAddress {
	ret := newQHostAddress(QHostAddress_new7(address))
	ret.isSubclass = true
	return ret
}

func (this *QHostAddress) OperatorAssign(other *QHostAddress) {
	QHostAddress_OperatorAssign(this.h, other.cPointer())
}

func (this *QHostAddress) OperatorAssignWithAddress(address SpecialAddress) {
	QHostAddress_OperatorAssignWithAddress(this.h, address)
}

func (this *QHostAddress) Swap(other *QHostAddress) {
	QHostAddress_Swap(this.h, other.cPointer())
}

func (this *QHostAddress) SetAddress(ip4Addr uint) {
	QHostAddress_SetAddress(this.h, (uint)(ip4Addr))
}

func (this *QHostAddress) SetAddressWithIp6Addr(ip6Addr *byte) {
	QHostAddress_SetAddressWithIp6Addr(this.h, (*uchar)(unsafe.Pointer(ip6Addr)))
}

func (this *QHostAddress) SetAddress2(ip6Addr *QIPv6Address) {
	QHostAddress_SetAddress2(this.h, ip6Addr.cPointer())
}

func (this *QHostAddress) SetAddress3(address string) bool {
	address_ms := struct_miqt_string{}
	address_ms.data = CString(address)
	address_ms.len = size_t(len(address))
	defer free(unsafe.Pointer(address_ms.data))
	return (bool)(QHostAddress_SetAddress3(this.h, address_ms))
}

func (this *QHostAddress) SetAddress4(address SpecialAddress) {
	QHostAddress_SetAddress4(this.h, address)
}

func (this *QHostAddress) Protocol() NetworkLayerProtocol {
	xxxxxxxxx
}

func (this *QHostAddress) ToIPv4Address() uint {
	return (uint)(QHostAddress_ToIPv4Address(this.h))
}

func (this *QHostAddress) ToIPv6Address() *QIPv6Address {
	_goptr := newQIPv6Address(QHostAddress_ToIPv6Address(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHostAddress) ToString() string {
	var _ms struct_miqt_string = QHostAddress_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHostAddress) ScopeId() string {
	var _ms struct_miqt_string = QHostAddress_ScopeId(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHostAddress) SetScopeId(id string) {
	id_ms := struct_miqt_string{}
	id_ms.data = CString(id)
	id_ms.len = size_t(len(id))
	defer free(unsafe.Pointer(id_ms.data))
	QHostAddress_SetScopeId(this.h, id_ms)
}

func (this *QHostAddress) IsEqual(address *QHostAddress) bool {
	return (bool)(QHostAddress_IsEqual(this.h, address.cPointer()))
}

func (this *QHostAddress) OperatorEqual(address *QHostAddress) bool {
	return (bool)(QHostAddress_OperatorEqual(this.h, address.cPointer()))
}

func (this *QHostAddress) OperatorEqualWithAddress(address SpecialAddress) bool {
	return (bool)(QHostAddress_OperatorEqualWithAddress(this.h, address))
}

func (this *QHostAddress) OperatorNotEqual(address *QHostAddress) bool {
	return (bool)(QHostAddress_OperatorNotEqual(this.h, address.cPointer()))
}

func (this *QHostAddress) OperatorNotEqualWithAddress(address SpecialAddress) bool {
	return (bool)(QHostAddress_OperatorNotEqualWithAddress(this.h, address))
}

func (this *QHostAddress) IsNull() bool {
	return (bool)(QHostAddress_IsNull(this.h))
}

func (this *QHostAddress) Clear() {
	QHostAddress_Clear(this.h)
}

func (this *QHostAddress) IsInSubnet(subnet *QHostAddress, netmask int) bool {
	return (bool)(QHostAddress_IsInSubnet(this.h, subnet.cPointer(), (int)(netmask)))
}

func (this *QHostAddress) IsLoopback() bool {
	return (bool)(QHostAddress_IsLoopback(this.h))
}

func (this *QHostAddress) IsGlobal() bool {
	return (bool)(QHostAddress_IsGlobal(this.h))
}

func (this *QHostAddress) IsLinkLocal() bool {
	return (bool)(QHostAddress_IsLinkLocal(this.h))
}

func (this *QHostAddress) IsSiteLocal() bool {
	return (bool)(QHostAddress_IsSiteLocal(this.h))
}

func (this *QHostAddress) IsUniqueLocalUnicast() bool {
	return (bool)(QHostAddress_IsUniqueLocalUnicast(this.h))
}

func (this *QHostAddress) IsMulticast() bool {
	return (bool)(QHostAddress_IsMulticast(this.h))
}

func (this *QHostAddress) IsBroadcast() bool {
	return (bool)(QHostAddress_IsBroadcast(this.h))
}

func (this *QHostAddress) IsPrivateUse() bool {
	return (bool)(QHostAddress_IsPrivateUse(this.h))
}

func (this *QHostAddress) ToIPv4Address1(ok *bool) uint {
	return (uint)(QHostAddress_ToIPv4Address1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QHostAddress) IsEqual2(address *QHostAddress, mode ConversionMode) bool {
	return (bool)(QHostAddress_IsEqual2(this.h, address.cPointer(), mode))
}
