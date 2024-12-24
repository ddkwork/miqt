package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QNetworkAddressEntry__DnsEligibilityStatus int8

const (
	QNetworkAddressEntry__DnsEligibilityUnknown QNetworkAddressEntry__DnsEligibilityStatus = -1
	QNetworkAddressEntry__DnsIneligible         QNetworkAddressEntry__DnsEligibilityStatus = 0
	QNetworkAddressEntry__DnsEligible           QNetworkAddressEntry__DnsEligibilityStatus = 1
)

type QNetworkInterface__InterfaceFlag int

const (
	QNetworkInterface__IsUp           QNetworkInterface__InterfaceFlag = 1
	QNetworkInterface__IsRunning      QNetworkInterface__InterfaceFlag = 2
	QNetworkInterface__CanBroadcast   QNetworkInterface__InterfaceFlag = 4
	QNetworkInterface__IsLoopBack     QNetworkInterface__InterfaceFlag = 8
	QNetworkInterface__IsPointToPoint QNetworkInterface__InterfaceFlag = 16
	QNetworkInterface__CanMulticast   QNetworkInterface__InterfaceFlag = 32
)

type QNetworkInterface__InterfaceType int

const (
	QNetworkInterface__Loopback   QNetworkInterface__InterfaceType = 1
	QNetworkInterface__Virtual    QNetworkInterface__InterfaceType = 2
	QNetworkInterface__Ethernet   QNetworkInterface__InterfaceType = 3
	QNetworkInterface__Slip       QNetworkInterface__InterfaceType = 4
	QNetworkInterface__CanBus     QNetworkInterface__InterfaceType = 5
	QNetworkInterface__Ppp        QNetworkInterface__InterfaceType = 6
	QNetworkInterface__Fddi       QNetworkInterface__InterfaceType = 7
	QNetworkInterface__Wifi       QNetworkInterface__InterfaceType = 8
	QNetworkInterface__Ieee80211  QNetworkInterface__InterfaceType = 8
	QNetworkInterface__Phonet     QNetworkInterface__InterfaceType = 9
	QNetworkInterface__Ieee802154 QNetworkInterface__InterfaceType = 10
	QNetworkInterface__SixLoWPAN  QNetworkInterface__InterfaceType = 11
	QNetworkInterface__Ieee80216  QNetworkInterface__InterfaceType = 12
	QNetworkInterface__Ieee1394   QNetworkInterface__InterfaceType = 13
	QNetworkInterface__Unknown    QNetworkInterface__InterfaceType = 0
)

type QNetworkAddressEntry struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkAddressEntry constructs a new QNetworkAddressEntry object.
func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	g := newQNetworkAddressEntry(QNetworkAddressEntry_new())
	g.isSubclass = true
	return g
}

// NewQNetworkAddressEntry2 constructs a new QNetworkAddressEntry object.
func NewQNetworkAddressEntry2(other *QNetworkAddressEntry) *QNetworkAddressEntry {
	g := newQNetworkAddressEntry(QNetworkAddressEntry_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QNetworkAddressEntry) OperatorAssign(other *QNetworkAddressEntry) {
	QNetworkAddressEntry_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkAddressEntry) Swap(other *QNetworkAddressEntry) {
	QNetworkAddressEntry_Swap(this.h, other.cPointer())
}

func (this *QNetworkAddressEntry) OperatorEqual(other *QNetworkAddressEntry) bool {
	return (bool)(QNetworkAddressEntry_OperatorEqual(this.h, other.cPointer()))
}

func (this *QNetworkAddressEntry) OperatorNotEqual(other *QNetworkAddressEntry) bool {
	return (bool)(QNetworkAddressEntry_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QNetworkAddressEntry) DnsEligibility() DnsEligibilityStatus {
	xxxxxxxxx
}

func (this *QNetworkAddressEntry) SetDnsEligibility(status DnsEligibilityStatus) {
	QNetworkAddressEntry_SetDnsEligibility(this.h, status)
}

func (this *QNetworkAddressEntry) Ip() *QHostAddress {
	_goptr := newQHostAddress(QNetworkAddressEntry_Ip(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkAddressEntry) SetIp(newIp *QHostAddress) {
	QNetworkAddressEntry_SetIp(this.h, newIp.cPointer())
}

func (this *QNetworkAddressEntry) Netmask() *QHostAddress {
	_goptr := newQHostAddress(QNetworkAddressEntry_Netmask(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkAddressEntry) SetNetmask(newNetmask *QHostAddress) {
	QNetworkAddressEntry_SetNetmask(this.h, newNetmask.cPointer())
}

func (this *QNetworkAddressEntry) PrefixLength() int {
	return (int)(QNetworkAddressEntry_PrefixLength(this.h))
}

func (this *QNetworkAddressEntry) SetPrefixLength(length int) {
	QNetworkAddressEntry_SetPrefixLength(this.h, (int)(length))
}

func (this *QNetworkAddressEntry) Broadcast() *QHostAddress {
	_goptr := newQHostAddress(QNetworkAddressEntry_Broadcast(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkAddressEntry) SetBroadcast(newBroadcast *QHostAddress) {
	QNetworkAddressEntry_SetBroadcast(this.h, newBroadcast.cPointer())
}

func (this *QNetworkAddressEntry) IsLifetimeKnown() bool {
	return (bool)(QNetworkAddressEntry_IsLifetimeKnown(this.h))
}

func (this *QNetworkAddressEntry) PreferredLifetime() *qt.QDeadlineTimer {
	_goptr := qt.UnsafeNewQDeadlineTimer(unsafe.Pointer(QNetworkAddressEntry_PreferredLifetime(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkAddressEntry) ValidityLifetime() *qt.QDeadlineTimer {
	_goptr := qt.UnsafeNewQDeadlineTimer(unsafe.Pointer(QNetworkAddressEntry_ValidityLifetime(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkAddressEntry) SetAddressLifetime(preferred qt.QDeadlineTimer, validity qt.QDeadlineTimer) {
	QNetworkAddressEntry_SetAddressLifetime(this.h, (*QDeadlineTimer)(preferred.UnsafePointer()), (*QDeadlineTimer)(validity.UnsafePointer()))
}

func (this *QNetworkAddressEntry) ClearAddressLifetime() {
	QNetworkAddressEntry_ClearAddressLifetime(this.h)
}

func (this *QNetworkAddressEntry) IsPermanent() bool {
	return (bool)(QNetworkAddressEntry_IsPermanent(this.h))
}

func (this *QNetworkAddressEntry) IsTemporary() bool {
	return (bool)(QNetworkAddressEntry_IsTemporary(this.h))
}

type QNetworkInterface struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkInterface constructs a new QNetworkInterface object.
func NewQNetworkInterface() *QNetworkInterface {
	g := newQNetworkInterface(QNetworkInterface_new())
	g.isSubclass = true
	return g
}

// NewQNetworkInterface2 constructs a new QNetworkInterface object.
func NewQNetworkInterface2(other *QNetworkInterface) *QNetworkInterface {
	g := newQNetworkInterface(QNetworkInterface_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QNetworkInterface) OperatorAssign(other *QNetworkInterface) {
	QNetworkInterface_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkInterface) Swap(other *QNetworkInterface) {
	QNetworkInterface_Swap(this.h, other.cPointer())
}

func (this *QNetworkInterface) IsValid() bool {
	return (bool)(QNetworkInterface_IsValid(this.h))
}

func (this *QNetworkInterface) Index() int {
	return (int)(QNetworkInterface_Index(this.h))
}

func (this *QNetworkInterface) MaximumTransmissionUnit() int {
	return (int)(QNetworkInterface_MaximumTransmissionUnit(this.h))
}

func (this *QNetworkInterface) Name() string {
	var _ms struct_miqt_string = QNetworkInterface_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkInterface) HumanReadableName() string {
	var _ms struct_miqt_string = QNetworkInterface_HumanReadableName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkInterface) Flags() InterfaceFlags {
	xxxxxxxxx
}

func (this *QNetworkInterface) Type() InterfaceType {
	xxxxxxxxx
}

func (this *QNetworkInterface) HardwareAddress() string {
	var _ms struct_miqt_string = QNetworkInterface_HardwareAddress(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkInterface) AddressEntries() []QNetworkAddressEntry {
	var _ma struct_miqt_array = QNetworkInterface_AddressEntries(this.h)
	_ret := make([]QNetworkAddressEntry, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkAddressEntry)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkAddressEntry(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QNetworkInterface_InterfaceIndexFromName(name string) int {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	return (int)(QNetworkInterface_InterfaceIndexFromName(name_ms))
}

func QNetworkInterface_InterfaceFromName(name string) *QNetworkInterface {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQNetworkInterface(QNetworkInterface_InterfaceFromName(name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QNetworkInterface_InterfaceFromIndex(index int) *QNetworkInterface {
	_goptr := newQNetworkInterface(QNetworkInterface_InterfaceFromIndex((int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QNetworkInterface_InterfaceNameFromIndex(index int) string {
	var _ms struct_miqt_string = QNetworkInterface_InterfaceNameFromIndex((int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkInterface_AllInterfaces() []QNetworkInterface {
	var _ma struct_miqt_array = QNetworkInterface_AllInterfaces()
	_ret := make([]QNetworkInterface, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkInterface)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkInterface(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QNetworkInterface_AllAddresses() []QHostAddress {
	var _ma struct_miqt_array = QNetworkInterface_AllAddresses()
	_ret := make([]QHostAddress, int(_ma.len))
	_outCast := (*[0xffff]*QHostAddress)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQHostAddress(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}
