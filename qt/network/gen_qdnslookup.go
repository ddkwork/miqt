package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QDnsTlsAssociationRecord__CertificateUsage byte

const (
	QDnsTlsAssociationRecord__CertificateUsage__CertificateAuthorityConstrait QDnsTlsAssociationRecord__CertificateUsage = 0
	QDnsTlsAssociationRecord__CertificateUsage__ServiceCertificateConstraint  QDnsTlsAssociationRecord__CertificateUsage = 1
	QDnsTlsAssociationRecord__CertificateUsage__TrustAnchorAssertion          QDnsTlsAssociationRecord__CertificateUsage = 2
	QDnsTlsAssociationRecord__CertificateUsage__DomainIssuedCertificate       QDnsTlsAssociationRecord__CertificateUsage = 3
	QDnsTlsAssociationRecord__CertificateUsage__PrivateUse                    QDnsTlsAssociationRecord__CertificateUsage = 255
	QDnsTlsAssociationRecord__CertificateUsage__PKIX_TA                       QDnsTlsAssociationRecord__CertificateUsage = 0
	QDnsTlsAssociationRecord__CertificateUsage__PKIX_EE                       QDnsTlsAssociationRecord__CertificateUsage = 1
	QDnsTlsAssociationRecord__CertificateUsage__DANE_TA                       QDnsTlsAssociationRecord__CertificateUsage = 2
	QDnsTlsAssociationRecord__CertificateUsage__DANE_EE                       QDnsTlsAssociationRecord__CertificateUsage = 3
	QDnsTlsAssociationRecord__CertificateUsage__PrivCert                      QDnsTlsAssociationRecord__CertificateUsage = 255
)

type QDnsTlsAssociationRecord__Selector byte

const (
	QDnsTlsAssociationRecord__Selector__FullCertificate      QDnsTlsAssociationRecord__Selector = 0
	QDnsTlsAssociationRecord__Selector__SubjectPublicKeyInfo QDnsTlsAssociationRecord__Selector = 1
	QDnsTlsAssociationRecord__Selector__PrivateUse           QDnsTlsAssociationRecord__Selector = 255
	QDnsTlsAssociationRecord__Selector__Cert                 QDnsTlsAssociationRecord__Selector = 0
	QDnsTlsAssociationRecord__Selector__SPKI                 QDnsTlsAssociationRecord__Selector = 1
	QDnsTlsAssociationRecord__Selector__PrivSel              QDnsTlsAssociationRecord__Selector = 255
)

type QDnsTlsAssociationRecord__MatchingType byte

const (
	QDnsTlsAssociationRecord__MatchingType__Exact      QDnsTlsAssociationRecord__MatchingType = 0
	QDnsTlsAssociationRecord__MatchingType__Sha256     QDnsTlsAssociationRecord__MatchingType = 1
	QDnsTlsAssociationRecord__MatchingType__Sha512     QDnsTlsAssociationRecord__MatchingType = 2
	QDnsTlsAssociationRecord__MatchingType__PrivateUse QDnsTlsAssociationRecord__MatchingType = 255
	QDnsTlsAssociationRecord__MatchingType__PrivMatch  QDnsTlsAssociationRecord__MatchingType = 255
)

type QDnsLookup__Error int

const (
	QDnsLookup__NoError                 QDnsLookup__Error = 0
	QDnsLookup__ResolverError           QDnsLookup__Error = 1
	QDnsLookup__OperationCancelledError QDnsLookup__Error = 2
	QDnsLookup__InvalidRequestError     QDnsLookup__Error = 3
	QDnsLookup__InvalidReplyError       QDnsLookup__Error = 4
	QDnsLookup__ServerFailureError      QDnsLookup__Error = 5
	QDnsLookup__ServerRefusedError      QDnsLookup__Error = 6
	QDnsLookup__NotFoundError           QDnsLookup__Error = 7
	QDnsLookup__TimeoutError            QDnsLookup__Error = 8
)

type QDnsLookup__Type int

const (
	QDnsLookup__A     QDnsLookup__Type = 1
	QDnsLookup__AAAA  QDnsLookup__Type = 28
	QDnsLookup__ANY   QDnsLookup__Type = 255
	QDnsLookup__CNAME QDnsLookup__Type = 5
	QDnsLookup__MX    QDnsLookup__Type = 15
	QDnsLookup__NS    QDnsLookup__Type = 2
	QDnsLookup__PTR   QDnsLookup__Type = 12
	QDnsLookup__SRV   QDnsLookup__Type = 33
	QDnsLookup__TLSA  QDnsLookup__Type = 52
	QDnsLookup__TXT   QDnsLookup__Type = 16
)

type QDnsLookup__Protocol byte

const (
	QDnsLookup__Standard   QDnsLookup__Protocol = 0
	QDnsLookup__DnsOverTls QDnsLookup__Protocol = 1
)

type QDnsDomainNameRecord struct {
	h          uintptr
	isSubclass bool
}

// NewQDnsDomainNameRecord constructs a new QDnsDomainNameRecord object.
func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	ret := newQDnsDomainNameRecord(QDnsDomainNameRecord_new())
	ret.isSubclass = true
	return ret
}

// NewQDnsDomainNameRecord2 constructs a new QDnsDomainNameRecord object.
func NewQDnsDomainNameRecord2(other *QDnsDomainNameRecord) *QDnsDomainNameRecord {
	ret := newQDnsDomainNameRecord(QDnsDomainNameRecord_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDnsDomainNameRecord) OperatorAssign(other *QDnsDomainNameRecord) {
	QDnsDomainNameRecord_OperatorAssign(this.h, other.cPointer())
}

func (this *QDnsDomainNameRecord) Swap(other *QDnsDomainNameRecord) {
	QDnsDomainNameRecord_Swap(this.h, other.cPointer())
}

func (this *QDnsDomainNameRecord) Name() string {
	var _ms struct_miqt_string = QDnsDomainNameRecord_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsDomainNameRecord) TimeToLive() uint {
	return (uint)(QDnsDomainNameRecord_TimeToLive(this.h))
}

func (this *QDnsDomainNameRecord) Value() string {
	var _ms struct_miqt_string = QDnsDomainNameRecord_Value(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QDnsHostAddressRecord struct {
	h          uintptr
	isSubclass bool
}

// NewQDnsHostAddressRecord constructs a new QDnsHostAddressRecord object.
func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	ret := newQDnsHostAddressRecord(QDnsHostAddressRecord_new())
	ret.isSubclass = true
	return ret
}

// NewQDnsHostAddressRecord2 constructs a new QDnsHostAddressRecord object.
func NewQDnsHostAddressRecord2(other *QDnsHostAddressRecord) *QDnsHostAddressRecord {
	ret := newQDnsHostAddressRecord(QDnsHostAddressRecord_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDnsHostAddressRecord) OperatorAssign(other *QDnsHostAddressRecord) {
	QDnsHostAddressRecord_OperatorAssign(this.h, other.cPointer())
}

func (this *QDnsHostAddressRecord) Swap(other *QDnsHostAddressRecord) {
	QDnsHostAddressRecord_Swap(this.h, other.cPointer())
}

func (this *QDnsHostAddressRecord) Name() string {
	var _ms struct_miqt_string = QDnsHostAddressRecord_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsHostAddressRecord) TimeToLive() uint {
	return (uint)(QDnsHostAddressRecord_TimeToLive(this.h))
}

func (this *QDnsHostAddressRecord) Value() *QHostAddress {
	_goptr := newQHostAddress(QDnsHostAddressRecord_Value(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QDnsMailExchangeRecord struct {
	h          uintptr
	isSubclass bool
}

// NewQDnsMailExchangeRecord constructs a new QDnsMailExchangeRecord object.
func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	ret := newQDnsMailExchangeRecord(QDnsMailExchangeRecord_new())
	ret.isSubclass = true
	return ret
}

// NewQDnsMailExchangeRecord2 constructs a new QDnsMailExchangeRecord object.
func NewQDnsMailExchangeRecord2(other *QDnsMailExchangeRecord) *QDnsMailExchangeRecord {
	ret := newQDnsMailExchangeRecord(QDnsMailExchangeRecord_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDnsMailExchangeRecord) OperatorAssign(other *QDnsMailExchangeRecord) {
	QDnsMailExchangeRecord_OperatorAssign(this.h, other.cPointer())
}

func (this *QDnsMailExchangeRecord) Swap(other *QDnsMailExchangeRecord) {
	QDnsMailExchangeRecord_Swap(this.h, other.cPointer())
}

func (this *QDnsMailExchangeRecord) Exchange() string {
	var _ms struct_miqt_string = QDnsMailExchangeRecord_Exchange(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsMailExchangeRecord) Name() string {
	var _ms struct_miqt_string = QDnsMailExchangeRecord_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsMailExchangeRecord) Preference() uint16 {
	return (uint16)(QDnsMailExchangeRecord_Preference(this.h))
}

func (this *QDnsMailExchangeRecord) TimeToLive() uint {
	return (uint)(QDnsMailExchangeRecord_TimeToLive(this.h))
}

type QDnsServiceRecord struct {
	h          uintptr
	isSubclass bool
}

// NewQDnsServiceRecord constructs a new QDnsServiceRecord object.
func NewQDnsServiceRecord() *QDnsServiceRecord {
	ret := newQDnsServiceRecord(QDnsServiceRecord_new())
	ret.isSubclass = true
	return ret
}

// NewQDnsServiceRecord2 constructs a new QDnsServiceRecord object.
func NewQDnsServiceRecord2(other *QDnsServiceRecord) *QDnsServiceRecord {
	ret := newQDnsServiceRecord(QDnsServiceRecord_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDnsServiceRecord) OperatorAssign(other *QDnsServiceRecord) {
	QDnsServiceRecord_OperatorAssign(this.h, other.cPointer())
}

func (this *QDnsServiceRecord) Swap(other *QDnsServiceRecord) {
	QDnsServiceRecord_Swap(this.h, other.cPointer())
}

func (this *QDnsServiceRecord) Name() string {
	var _ms struct_miqt_string = QDnsServiceRecord_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsServiceRecord) Port() uint16 {
	return (uint16)(QDnsServiceRecord_Port(this.h))
}

func (this *QDnsServiceRecord) Priority() uint16 {
	return (uint16)(QDnsServiceRecord_Priority(this.h))
}

func (this *QDnsServiceRecord) Target() string {
	var _ms struct_miqt_string = QDnsServiceRecord_Target(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsServiceRecord) TimeToLive() uint {
	return (uint)(QDnsServiceRecord_TimeToLive(this.h))
}

func (this *QDnsServiceRecord) Weight() uint16 {
	return (uint16)(QDnsServiceRecord_Weight(this.h))
}

type QDnsTextRecord struct {
	h          uintptr
	isSubclass bool
}

// NewQDnsTextRecord constructs a new QDnsTextRecord object.
func NewQDnsTextRecord() *QDnsTextRecord {
	ret := newQDnsTextRecord(QDnsTextRecord_new())
	ret.isSubclass = true
	return ret
}

// NewQDnsTextRecord2 constructs a new QDnsTextRecord object.
func NewQDnsTextRecord2(other *QDnsTextRecord) *QDnsTextRecord {
	ret := newQDnsTextRecord(QDnsTextRecord_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDnsTextRecord) OperatorAssign(other *QDnsTextRecord) {
	QDnsTextRecord_OperatorAssign(this.h, other.cPointer())
}

func (this *QDnsTextRecord) Swap(other *QDnsTextRecord) {
	QDnsTextRecord_Swap(this.h, other.cPointer())
}

func (this *QDnsTextRecord) Name() string {
	var _ms struct_miqt_string = QDnsTextRecord_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsTextRecord) TimeToLive() uint {
	return (uint)(QDnsTextRecord_TimeToLive(this.h))
}

func (this *QDnsTextRecord) Values() [][]byte {
	var _ma struct_miqt_array = QDnsTextRecord_Values(this.h)
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

type QDnsTlsAssociationRecord struct {
	h          uintptr
	isSubclass bool
}

// NewQDnsTlsAssociationRecord constructs a new QDnsTlsAssociationRecord object.
func NewQDnsTlsAssociationRecord() *QDnsTlsAssociationRecord {
	ret := newQDnsTlsAssociationRecord(QDnsTlsAssociationRecord_new())
	ret.isSubclass = true
	return ret
}

// NewQDnsTlsAssociationRecord2 constructs a new QDnsTlsAssociationRecord object.
func NewQDnsTlsAssociationRecord2(other *QDnsTlsAssociationRecord) *QDnsTlsAssociationRecord {
	ret := newQDnsTlsAssociationRecord(QDnsTlsAssociationRecord_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDnsTlsAssociationRecord) OperatorAssign(other *QDnsTlsAssociationRecord) {
	QDnsTlsAssociationRecord_OperatorAssign(this.h, other.cPointer())
}

func (this *QDnsTlsAssociationRecord) Swap(other *QDnsTlsAssociationRecord) {
	QDnsTlsAssociationRecord_Swap(this.h, other.cPointer())
}

func (this *QDnsTlsAssociationRecord) Name() string {
	var _ms struct_miqt_string = QDnsTlsAssociationRecord_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsTlsAssociationRecord) TimeToLive() uint {
	return (uint)(QDnsTlsAssociationRecord_TimeToLive(this.h))
}

func (this *QDnsTlsAssociationRecord) Usage() CertificateUsage {
	xxxxxxxxx
}

func (this *QDnsTlsAssociationRecord) Selector() Selector {
	xxxxxxxxx
}

func (this *QDnsTlsAssociationRecord) MatchType() MatchingType {
	xxxxxxxxx
}

func (this *QDnsTlsAssociationRecord) Value() []byte {
	var _bytearray struct_miqt_string = QDnsTlsAssociationRecord_Value(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

type QDnsLookup struct {
	h          uintptr
	isSubclass bool
}

// NewQDnsLookup constructs a new QDnsLookup object.
func NewQDnsLookup() *QDnsLookup {
	ret := newQDnsLookup(QDnsLookup_new())
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup2 constructs a new QDnsLookup object.
func NewQDnsLookup2(typeVal Type, name string) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new2(typeVal, name_ms))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup3 constructs a new QDnsLookup object.
func NewQDnsLookup3(typeVal Type, name string, nameserver *QHostAddress) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new3(typeVal, name_ms, nameserver.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup4 constructs a new QDnsLookup object.
func NewQDnsLookup4(typeVal Type, name string, nameserver *QHostAddress, port uint16) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new4(typeVal, name_ms, nameserver.cPointer(), (uint16_t)(port)))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup5 constructs a new QDnsLookup object.
func NewQDnsLookup5(typeVal Type, name string, protocol Protocol, nameserver *QHostAddress) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new5(typeVal, name_ms, protocol, nameserver.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup6 constructs a new QDnsLookup object.
func NewQDnsLookup6(parent *qt.QObject) *QDnsLookup {
	ret := newQDnsLookup(QDnsLookup_new6((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup7 constructs a new QDnsLookup object.
func NewQDnsLookup7(typeVal Type, name string, parent *qt.QObject) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new7(typeVal, name_ms, (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup8 constructs a new QDnsLookup object.
func NewQDnsLookup8(typeVal Type, name string, nameserver *QHostAddress, parent *qt.QObject) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new8(typeVal, name_ms, nameserver.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup9 constructs a new QDnsLookup object.
func NewQDnsLookup9(typeVal Type, name string, nameserver *QHostAddress, port uint16, parent *qt.QObject) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new9(typeVal, name_ms, nameserver.cPointer(), (uint16_t)(port), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup10 constructs a new QDnsLookup object.
func NewQDnsLookup10(typeVal Type, name string, protocol Protocol, nameserver *QHostAddress, port uint16) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new10(typeVal, name_ms, protocol, nameserver.cPointer(), (uint16_t)(port)))
	ret.isSubclass = true
	return ret
}

// NewQDnsLookup11 constructs a new QDnsLookup object.
func NewQDnsLookup11(typeVal Type, name string, protocol Protocol, nameserver *QHostAddress, port uint16, parent *qt.QObject) *QDnsLookup {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQDnsLookup(QDnsLookup_new11(typeVal, name_ms, protocol, nameserver.cPointer(), (uint16_t)(port), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QDnsLookup) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QDnsLookup_MetaObject(this.h)))
}

func (this *QDnsLookup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDnsLookup_Metacast(this.h, param1_Cstring))
}

func QDnsLookup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDnsLookup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsLookup) IsAuthenticData() bool {
	return (bool)(QDnsLookup_IsAuthenticData(this.h))
}

func (this *QDnsLookup) Error() Error {
	xxxxxxxxx
}

func (this *QDnsLookup) ErrorString() string {
	var _ms struct_miqt_string = QDnsLookup_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsLookup) IsFinished() bool {
	return (bool)(QDnsLookup_IsFinished(this.h))
}

func (this *QDnsLookup) Name() string {
	var _ms struct_miqt_string = QDnsLookup_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsLookup) SetName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QDnsLookup_SetName(this.h, name_ms)
}

func (this *QDnsLookup) Type() Type {
	xxxxxxxxx
}

func (this *QDnsLookup) SetType(typeVal QDnsLookup__Type) {
	QDnsLookup_SetType(this.h, (int)(typeVal))
}

func (this *QDnsLookup) Nameserver() *QHostAddress {
	_goptr := newQHostAddress(QDnsLookup_Nameserver(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDnsLookup) SetNameserver(nameserver *QHostAddress) {
	QDnsLookup_SetNameserver(this.h, nameserver.cPointer())
}

func (this *QDnsLookup) NameserverPort() uint16 {
	return (uint16)(QDnsLookup_NameserverPort(this.h))
}

func (this *QDnsLookup) SetNameserverPort(port uint16) {
	QDnsLookup_SetNameserverPort(this.h, (uint16_t)(port))
}

func (this *QDnsLookup) NameserverProtocol() Protocol {
	xxxxxxxxx
}

func (this *QDnsLookup) SetNameserverProtocol(protocol Protocol) {
	QDnsLookup_SetNameserverProtocol(this.h, protocol)
}

func (this *QDnsLookup) SetNameserver2(protocol Protocol, nameserver *QHostAddress) {
	QDnsLookup_SetNameserver2(this.h, protocol, nameserver.cPointer())
}

func (this *QDnsLookup) SetNameserver3(nameserver *QHostAddress, port uint16) {
	QDnsLookup_SetNameserver3(this.h, nameserver.cPointer(), (uint16_t)(port))
}

func (this *QDnsLookup) CanonicalNameRecords() []QDnsDomainNameRecord {
	var _ma struct_miqt_array = QDnsLookup_CanonicalNameRecords(this.h)
	_ret := make([]QDnsDomainNameRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsDomainNameRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsDomainNameRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) HostAddressRecords() []QDnsHostAddressRecord {
	var _ma struct_miqt_array = QDnsLookup_HostAddressRecords(this.h)
	_ret := make([]QDnsHostAddressRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsHostAddressRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsHostAddressRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) MailExchangeRecords() []QDnsMailExchangeRecord {
	var _ma struct_miqt_array = QDnsLookup_MailExchangeRecords(this.h)
	_ret := make([]QDnsMailExchangeRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsMailExchangeRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsMailExchangeRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) NameServerRecords() []QDnsDomainNameRecord {
	var _ma struct_miqt_array = QDnsLookup_NameServerRecords(this.h)
	_ret := make([]QDnsDomainNameRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsDomainNameRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsDomainNameRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) PointerRecords() []QDnsDomainNameRecord {
	var _ma struct_miqt_array = QDnsLookup_PointerRecords(this.h)
	_ret := make([]QDnsDomainNameRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsDomainNameRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsDomainNameRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) ServiceRecords() []QDnsServiceRecord {
	var _ma struct_miqt_array = QDnsLookup_ServiceRecords(this.h)
	_ret := make([]QDnsServiceRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsServiceRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsServiceRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) TextRecords() []QDnsTextRecord {
	var _ma struct_miqt_array = QDnsLookup_TextRecords(this.h)
	_ret := make([]QDnsTextRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsTextRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsTextRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) TlsAssociationRecords() []QDnsTlsAssociationRecord {
	var _ma struct_miqt_array = QDnsLookup_TlsAssociationRecords(this.h)
	_ret := make([]QDnsTlsAssociationRecord, int(_ma.len))
	_outCast := (*[0xffff]*QDnsTlsAssociationRecord)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQDnsTlsAssociationRecord(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDnsLookup) SetSslConfiguration(sslConfiguration *QSslConfiguration) {
	QDnsLookup_SetSslConfiguration(this.h, sslConfiguration.cPointer())
}

func (this *QDnsLookup) SslConfiguration() *QSslConfiguration {
	_goptr := newQSslConfiguration(QDnsLookup_SslConfiguration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDnsLookup_IsProtocolSupported(protocol Protocol) bool {
	return (bool)(QDnsLookup_IsProtocolSupported(protocol))
}

func QDnsLookup_DefaultPortForProtocol(protocol Protocol) uint16 {
	return (uint16)(QDnsLookup_DefaultPortForProtocol(protocol))
}

func (this *QDnsLookup) Abort() {
	QDnsLookup_Abort(this.h)
}

func (this *QDnsLookup) Lookup() {
	QDnsLookup_Lookup(this.h)
}

func (this *QDnsLookup) Finished() {
	QDnsLookup_Finished(this.h)
}

func (this *QDnsLookup) OnFinished(slot func()) {
	QDnsLookup_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_Finished
func miqt_exec_callback_QDnsLookup_Finished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QDnsLookup) NameChanged(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QDnsLookup_NameChanged(this.h, name_ms)
}

func (this *QDnsLookup) OnNameChanged(slot func(name string)) {
	QDnsLookup_connect_NameChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_NameChanged
func miqt_exec_callback_QDnsLookup_NameChanged(cb intptr_t, name struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(name string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var name_ms struct_miqt_string = name
	name_ret := GoStringN(name_ms.data, int(int64(name_ms.len)))
	free(unsafe.Pointer(name_ms.data))
	slotval1 := name_ret

	gofunc(slotval1)
}

func (this *QDnsLookup) TypeChanged(typeVal QDnsLookup__Type) {
	QDnsLookup_TypeChanged(this.h, (int)(typeVal))
}

func (this *QDnsLookup) OnTypeChanged(slot func(typeVal QDnsLookup__Type)) {
	QDnsLookup_connect_TypeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_TypeChanged
func miqt_exec_callback_QDnsLookup_TypeChanged(cb intptr_t, typeVal int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(typeVal QDnsLookup__Type))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QDnsLookup__Type)(typeVal)

	gofunc(slotval1)
}

func (this *QDnsLookup) NameserverChanged(nameserver *QHostAddress) {
	QDnsLookup_NameserverChanged(this.h, nameserver.cPointer())
}

func (this *QDnsLookup) OnNameserverChanged(slot func(nameserver *QHostAddress)) {
	QDnsLookup_connect_NameserverChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_NameserverChanged
func miqt_exec_callback_QDnsLookup_NameserverChanged(cb intptr_t, nameserver *QHostAddress) {
	gofunc, ok := cgo.Handle(cb).Value().(func(nameserver *QHostAddress))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHostAddress(nameserver)

	gofunc(slotval1)
}

func (this *QDnsLookup) NameserverPortChanged(port uint16) {
	QDnsLookup_NameserverPortChanged(this.h, (uint16_t)(port))
}

func (this *QDnsLookup) OnNameserverPortChanged(slot func(port uint16)) {
	QDnsLookup_connect_NameserverPortChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_NameserverPortChanged
func miqt_exec_callback_QDnsLookup_NameserverPortChanged(cb intptr_t, port uint16_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(port uint16))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint16)(port)

	gofunc(slotval1)
}

func (this *QDnsLookup) NameserverProtocolChanged(protocol QDnsLookup__Protocol) {
	QDnsLookup_NameserverProtocolChanged(this.h, (uint8_t)(protocol))
}

func (this *QDnsLookup) OnNameserverProtocolChanged(slot func(protocol QDnsLookup__Protocol)) {
	QDnsLookup_connect_NameserverProtocolChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_NameserverProtocolChanged
func miqt_exec_callback_QDnsLookup_NameserverProtocolChanged(cb intptr_t, protocol uint8_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(protocol QDnsLookup__Protocol))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QDnsLookup__Protocol)(protocol)

	gofunc(slotval1)
}

func QDnsLookup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDnsLookup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDnsLookup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDnsLookup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDnsLookup) SetNameserver32(protocol Protocol, nameserver *QHostAddress, port uint16) {
	QDnsLookup_SetNameserver32(this.h, protocol, nameserver.cPointer(), (uint16_t)(port))
}

func (this *QDnsLookup) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QDnsLookup_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QDnsLookup) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDnsLookup_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_MetaObject
func miqt_exec_callback_QDnsLookup_MetaObject(self QDnsLookup, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDnsLookup{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QDnsLookup) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDnsLookup_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDnsLookup) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDnsLookup_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDnsLookup_Metacast
func miqt_exec_callback_QDnsLookup_Metacast(self QDnsLookup, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QDnsLookup{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
