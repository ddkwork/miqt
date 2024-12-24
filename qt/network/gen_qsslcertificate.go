package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QSslCertificate__SubjectInfo int

const (
	QSslCertificate__Organization               QSslCertificate__SubjectInfo = 0
	QSslCertificate__CommonName                 QSslCertificate__SubjectInfo = 1
	QSslCertificate__LocalityName               QSslCertificate__SubjectInfo = 2
	QSslCertificate__OrganizationalUnitName     QSslCertificate__SubjectInfo = 3
	QSslCertificate__CountryName                QSslCertificate__SubjectInfo = 4
	QSslCertificate__StateOrProvinceName        QSslCertificate__SubjectInfo = 5
	QSslCertificate__DistinguishedNameQualifier QSslCertificate__SubjectInfo = 6
	QSslCertificate__SerialNumber               QSslCertificate__SubjectInfo = 7
	QSslCertificate__EmailAddress               QSslCertificate__SubjectInfo = 8
)

type QSslCertificate__PatternSyntax int

const (
	QSslCertificate__RegularExpression QSslCertificate__PatternSyntax = 0
	QSslCertificate__Wildcard          QSslCertificate__PatternSyntax = 1
	QSslCertificate__FixedString       QSslCertificate__PatternSyntax = 2
)

type QSslCertificate struct {
	h          uintptr
	isSubclass bool
}

// NewQSslCertificate constructs a new QSslCertificate object.
func NewQSslCertificate(device *qt.QIODevice) *QSslCertificate {

	ret := newQSslCertificate(QSslCertificate_new((*QIODevice)(device.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQSslCertificate2 constructs a new QSslCertificate object.
func NewQSslCertificate2() *QSslCertificate {

	ret := newQSslCertificate(QSslCertificate_new2())
	ret.isSubclass = true
	return ret
}

// NewQSslCertificate3 constructs a new QSslCertificate object.
func NewQSslCertificate3(other *QSslCertificate) *QSslCertificate {

	ret := newQSslCertificate(QSslCertificate_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSslCertificate4 constructs a new QSslCertificate object.
func NewQSslCertificate4(device *qt.QIODevice, format QSsl__EncodingFormat) *QSslCertificate {

	ret := newQSslCertificate(QSslCertificate_new4((*QIODevice)(device.UnsafePointer()), (int)(format)))
	ret.isSubclass = true
	return ret
}

// NewQSslCertificate5 constructs a new QSslCertificate object.
func NewQSslCertificate5(data []byte) *QSslCertificate {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))

	ret := newQSslCertificate(QSslCertificate_new5(data_alias))
	ret.isSubclass = true
	return ret
}

// NewQSslCertificate6 constructs a new QSslCertificate object.
func NewQSslCertificate6(data []byte, format QSsl__EncodingFormat) *QSslCertificate {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))

	ret := newQSslCertificate(QSslCertificate_new6(data_alias, (int)(format)))
	ret.isSubclass = true
	return ret
}

func (this *QSslCertificate) OperatorAssign(other *QSslCertificate) {
	QSslCertificate_OperatorAssign(this.h, other.cPointer())
}

func (this *QSslCertificate) Swap(other *QSslCertificate) {
	QSslCertificate_Swap(this.h, other.cPointer())
}

func (this *QSslCertificate) OperatorEqual(other *QSslCertificate) bool {
	return (bool)(QSslCertificate_OperatorEqual(this.h, other.cPointer()))
}

func (this *QSslCertificate) OperatorNotEqual(other *QSslCertificate) bool {
	return (bool)(QSslCertificate_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QSslCertificate) IsNull() bool {
	return (bool)(QSslCertificate_IsNull(this.h))
}

func (this *QSslCertificate) IsBlacklisted() bool {
	return (bool)(QSslCertificate_IsBlacklisted(this.h))
}

func (this *QSslCertificate) IsSelfSigned() bool {
	return (bool)(QSslCertificate_IsSelfSigned(this.h))
}

func (this *QSslCertificate) Clear() {
	QSslCertificate_Clear(this.h)
}

func (this *QSslCertificate) Version() []byte {
	var _bytearray struct_miqt_string = QSslCertificate_Version(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslCertificate) SerialNumber() []byte {
	var _bytearray struct_miqt_string = QSslCertificate_SerialNumber(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslCertificate) Digest() []byte {
	var _bytearray struct_miqt_string = QSslCertificate_Digest(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslCertificate) IssuerInfo(info SubjectInfo) []string {
	var _ma struct_miqt_array = QSslCertificate_IssuerInfo(this.h, info)
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

func (this *QSslCertificate) IssuerInfoWithAttribute(attribute []byte) []string {
	attribute_alias := struct_miqt_string{}
	attribute_alias.data = (char)(unsafe.Pointer(&attribute[0]))
	attribute_alias.len = size_t(len(attribute))
	var _ma struct_miqt_array = QSslCertificate_IssuerInfoWithAttribute(this.h, attribute_alias)
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

func (this *QSslCertificate) SubjectInfo(info SubjectInfo) []string {
	var _ma struct_miqt_array = QSslCertificate_SubjectInfo(this.h, info)
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

func (this *QSslCertificate) SubjectInfoWithAttribute(attribute []byte) []string {
	attribute_alias := struct_miqt_string{}
	attribute_alias.data = (char)(unsafe.Pointer(&attribute[0]))
	attribute_alias.len = size_t(len(attribute))
	var _ma struct_miqt_array = QSslCertificate_SubjectInfoWithAttribute(this.h, attribute_alias)
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

func (this *QSslCertificate) IssuerDisplayName() string {
	var _ms struct_miqt_string = QSslCertificate_IssuerDisplayName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCertificate) SubjectDisplayName() string {
	var _ms struct_miqt_string = QSslCertificate_SubjectDisplayName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCertificate) SubjectInfoAttributes() [][]byte {
	var _ma struct_miqt_array = QSslCertificate_SubjectInfoAttributes(this.h)
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

func (this *QSslCertificate) IssuerInfoAttributes() [][]byte {
	var _ma struct_miqt_array = QSslCertificate_IssuerInfoAttributes(this.h)
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

func (this *QSslCertificate) EffectiveDate() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QSslCertificate_EffectiveDate(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslCertificate) ExpiryDate() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QSslCertificate_ExpiryDate(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslCertificate) PublicKey() *QSslKey {
	_goptr := newQSslKey(QSslCertificate_PublicKey(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslCertificate) Extensions() []QSslCertificateExtension {
	var _ma struct_miqt_array = QSslCertificate_Extensions(this.h)
	_ret := make([]QSslCertificateExtension, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificateExtension)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificateExtension(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslCertificate) ToPem() []byte {
	var _bytearray struct_miqt_string = QSslCertificate_ToPem(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslCertificate) ToDer() []byte {
	var _bytearray struct_miqt_string = QSslCertificate_ToDer(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslCertificate) ToText() string {
	var _ms struct_miqt_string = QSslCertificate_ToText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSslCertificate_FromPath(path string) []QSslCertificate {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	var _ma struct_miqt_array = QSslCertificate_FromPath(path_ms)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_FromDevice(device *qt.QIODevice) []QSslCertificate {
	var _ma struct_miqt_array = QSslCertificate_FromDevice((*QIODevice)(device.UnsafePointer()))
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_FromData(data []byte) []QSslCertificate {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	var _ma struct_miqt_array = QSslCertificate_FromData(data_alias)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_Verify(certificateChain []QSslCertificate) []QSslError {
	certificateChain_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(certificateChain))))
	defer free(unsafe.Pointer(certificateChain_CArray))
	for i := range certificateChain {
		certificateChain_CArray[i] = certificateChain[i].cPointer()
	}
	certificateChain_ma := struct_miqt_array{len: size_t(len(certificateChain)), data: unsafe.Pointer(certificateChain_CArray)}
	var _ma struct_miqt_array = QSslCertificate_Verify(certificateChain_ma)
	_ret := make([]QSslError, int(_ma.len))
	_outCast := (*[0xffff]*QSslError)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslError(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_ImportPkcs12(device *qt.QIODevice, key *QSslKey, cert *QSslCertificate) bool {
	return (bool)(QSslCertificate_ImportPkcs12((*QIODevice)(device.UnsafePointer()), key.cPointer(), cert.cPointer()))
}

func (this *QSslCertificate) Handle() unsafe.Pointer {
	return (unsafe.Pointer)(QSslCertificate_Handle(this.h))
}

func (this *QSslCertificate) Digest1(algorithm qt.QCryptographicHash__Algorithm) []byte {
	var _bytearray struct_miqt_string = QSslCertificate_Digest1(this.h, (int)(algorithm))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QSslCertificate_FromPath2(path string, format QSsl__EncodingFormat) []QSslCertificate {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	var _ma struct_miqt_array = QSslCertificate_FromPath2(path_ms, (int)(format))
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_FromPath3(path string, format QSsl__EncodingFormat, syntax PatternSyntax) []QSslCertificate {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	var _ma struct_miqt_array = QSslCertificate_FromPath3(path_ms, (int)(format), syntax)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_FromDevice2(device *qt.QIODevice, format QSsl__EncodingFormat) []QSslCertificate {
	var _ma struct_miqt_array = QSslCertificate_FromDevice2((*QIODevice)(device.UnsafePointer()), (int)(format))
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_FromData2(data []byte, format QSsl__EncodingFormat) []QSslCertificate {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	var _ma struct_miqt_array = QSslCertificate_FromData2(data_alias, (int)(format))
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_Verify2(certificateChain []QSslCertificate, hostName string) []QSslError {
	certificateChain_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(certificateChain))))
	defer free(unsafe.Pointer(certificateChain_CArray))
	for i := range certificateChain {
		certificateChain_CArray[i] = certificateChain[i].cPointer()
	}
	certificateChain_ma := struct_miqt_array{len: size_t(len(certificateChain)), data: unsafe.Pointer(certificateChain_CArray)}
	hostName_ms := struct_miqt_string{}
	hostName_ms.data = CString(hostName)
	hostName_ms.len = size_t(len(hostName))
	defer free(unsafe.Pointer(hostName_ms.data))
	var _ma struct_miqt_array = QSslCertificate_Verify2(certificateChain_ma, hostName_ms)
	_ret := make([]QSslError, int(_ma.len))
	_outCast := (*[0xffff]*QSslError)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslError(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QSslCertificate_ImportPkcs124(device *qt.QIODevice, key *QSslKey, cert *QSslCertificate, caCertificates []QSslCertificate) bool {
	caCertificates_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(caCertificates))))
	defer free(unsafe.Pointer(caCertificates_CArray))
	for i := range caCertificates {
		caCertificates_CArray[i] = caCertificates[i].cPointer()
	}
	caCertificates_ma := struct_miqt_array{len: size_t(len(caCertificates)), data: unsafe.Pointer(caCertificates_CArray)}
	return (bool)(QSslCertificate_ImportPkcs124((*QIODevice)(device.UnsafePointer()), key.cPointer(), cert.cPointer(), caCertificates_ma))
}

func QSslCertificate_ImportPkcs125(device *qt.QIODevice, key *QSslKey, cert *QSslCertificate, caCertificates []QSslCertificate, passPhrase []byte) bool {
	caCertificates_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(caCertificates))))
	defer free(unsafe.Pointer(caCertificates_CArray))
	for i := range caCertificates {
		caCertificates_CArray[i] = caCertificates[i].cPointer()
	}
	caCertificates_ma := struct_miqt_array{len: size_t(len(caCertificates)), data: unsafe.Pointer(caCertificates_CArray)}
	passPhrase_alias := struct_miqt_string{}
	passPhrase_alias.data = (char)(unsafe.Pointer(&passPhrase[0]))
	passPhrase_alias.len = size_t(len(passPhrase))
	return (bool)(QSslCertificate_ImportPkcs125((*QIODevice)(device.UnsafePointer()), key.cPointer(), cert.cPointer(), caCertificates_ma, passPhrase_alias))
}
