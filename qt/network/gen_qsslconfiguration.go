package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QSslConfiguration__NextProtocolNegotiationStatus int

const (
	QSslConfiguration__NextProtocolNegotiationNone        QSslConfiguration__NextProtocolNegotiationStatus = 0
	QSslConfiguration__NextProtocolNegotiationNegotiated  QSslConfiguration__NextProtocolNegotiationStatus = 1
	QSslConfiguration__NextProtocolNegotiationUnsupported QSslConfiguration__NextProtocolNegotiationStatus = 2
)

type QSslConfiguration struct {
	h          uintptr
	isSubclass bool
}

// NewQSslConfiguration constructs a new QSslConfiguration object.
func NewQSslConfiguration() *QSslConfiguration {

	ret := newQSslConfiguration(QSslConfiguration_new())
	ret.isSubclass = true
	return ret
}

// NewQSslConfiguration2 constructs a new QSslConfiguration object.
func NewQSslConfiguration2(other *QSslConfiguration) *QSslConfiguration {

	ret := newQSslConfiguration(QSslConfiguration_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSslConfiguration) OperatorAssign(other *QSslConfiguration) {
	QSslConfiguration_OperatorAssign(this.h, other.cPointer())
}

func (this *QSslConfiguration) Swap(other *QSslConfiguration) {
	QSslConfiguration_Swap(this.h, other.cPointer())
}

func (this *QSslConfiguration) OperatorEqual(other *QSslConfiguration) bool {
	return (bool)(QSslConfiguration_OperatorEqual(this.h, other.cPointer()))
}

func (this *QSslConfiguration) OperatorNotEqual(other *QSslConfiguration) bool {
	return (bool)(QSslConfiguration_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QSslConfiguration) IsNull() bool {
	return (bool)(QSslConfiguration_IsNull(this.h))
}

func (this *QSslConfiguration) Protocol() QSsl__SslProtocol {
	return (QSsl__SslProtocol)(QSslConfiguration_Protocol(this.h))
}

func (this *QSslConfiguration) SetProtocol(protocol QSsl__SslProtocol) {
	QSslConfiguration_SetProtocol(this.h, (int)(protocol))
}

func (this *QSslConfiguration) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	return (QSslSocket__PeerVerifyMode)(QSslConfiguration_PeerVerifyMode(this.h))
}

func (this *QSslConfiguration) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	QSslConfiguration_SetPeerVerifyMode(this.h, (int)(mode))
}

func (this *QSslConfiguration) PeerVerifyDepth() int {
	return (int)(QSslConfiguration_PeerVerifyDepth(this.h))
}

func (this *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	QSslConfiguration_SetPeerVerifyDepth(this.h, (int)(depth))
}

func (this *QSslConfiguration) LocalCertificateChain() []QSslCertificate {
	var _ma struct_miqt_array = QSslConfiguration_LocalCertificateChain(this.h)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) SetLocalCertificateChain(localChain []QSslCertificate) {
	localChain_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(localChain))))
	defer free(unsafe.Pointer(localChain_CArray))
	for i := range localChain {
		localChain_CArray[i] = localChain[i].cPointer()
	}
	localChain_ma := struct_miqt_array{len: size_t(len(localChain)), data: unsafe.Pointer(localChain_CArray)}
	QSslConfiguration_SetLocalCertificateChain(this.h, localChain_ma)
}

func (this *QSslConfiguration) LocalCertificate() *QSslCertificate {
	_goptr := newQSslCertificate(QSslConfiguration_LocalCertificate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslConfiguration) SetLocalCertificate(certificate *QSslCertificate) {
	QSslConfiguration_SetLocalCertificate(this.h, certificate.cPointer())
}

func (this *QSslConfiguration) PeerCertificate() *QSslCertificate {
	_goptr := newQSslCertificate(QSslConfiguration_PeerCertificate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslConfiguration) PeerCertificateChain() []QSslCertificate {
	var _ma struct_miqt_array = QSslConfiguration_PeerCertificateChain(this.h)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) SessionCipher() *QSslCipher {
	_goptr := newQSslCipher(QSslConfiguration_SessionCipher(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslConfiguration) SessionProtocol() QSsl__SslProtocol {
	return (QSsl__SslProtocol)(QSslConfiguration_SessionProtocol(this.h))
}

func (this *QSslConfiguration) PrivateKey() *QSslKey {
	_goptr := newQSslKey(QSslConfiguration_PrivateKey(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslConfiguration) SetPrivateKey(key *QSslKey) {
	QSslConfiguration_SetPrivateKey(this.h, key.cPointer())
}

func (this *QSslConfiguration) Ciphers() []QSslCipher {
	var _ma struct_miqt_array = QSslConfiguration_Ciphers(this.h)
	_ret := make([]QSslCipher, int(_ma.len))
	_outCast := (*[0xffff]*QSslCipher)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCipher(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) SetCiphers(ciphers []QSslCipher) {
	ciphers_CArray := (*[0xffff]*QSslCipher)(malloc(size_t(8 * len(ciphers))))
	defer free(unsafe.Pointer(ciphers_CArray))
	for i := range ciphers {
		ciphers_CArray[i] = ciphers[i].cPointer()
	}
	ciphers_ma := struct_miqt_array{len: size_t(len(ciphers)), data: unsafe.Pointer(ciphers_CArray)}
	QSslConfiguration_SetCiphers(this.h, ciphers_ma)
}

func (this *QSslConfiguration) SetCiphersWithCiphers(ciphers string) {
	ciphers_ms := struct_miqt_string{}
	ciphers_ms.data = CString(ciphers)
	ciphers_ms.len = size_t(len(ciphers))
	defer free(unsafe.Pointer(ciphers_ms.data))
	QSslConfiguration_SetCiphersWithCiphers(this.h, ciphers_ms)
}

func QSslConfiguration_SupportedCiphers() []QSslCipher {
	var _ma struct_miqt_array = QSslConfiguration_SupportedCiphers()
	_ret := make([]QSslCipher, int(_ma.len))
	_outCast := (*[0xffff]*QSslCipher)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCipher(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) CaCertificates() []QSslCertificate {
	var _ma struct_miqt_array = QSslConfiguration_CaCertificates(this.h)
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) SetCaCertificates(certificates []QSslCertificate) {
	certificates_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(certificates))))
	defer free(unsafe.Pointer(certificates_CArray))
	for i := range certificates {
		certificates_CArray[i] = certificates[i].cPointer()
	}
	certificates_ma := struct_miqt_array{len: size_t(len(certificates)), data: unsafe.Pointer(certificates_CArray)}
	QSslConfiguration_SetCaCertificates(this.h, certificates_ma)
}

func (this *QSslConfiguration) AddCaCertificates(path string) bool {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	return (bool)(QSslConfiguration_AddCaCertificates(this.h, path_ms))
}

func (this *QSslConfiguration) AddCaCertificate(certificate *QSslCertificate) {
	QSslConfiguration_AddCaCertificate(this.h, certificate.cPointer())
}

func (this *QSslConfiguration) AddCaCertificatesWithCertificates(certificates []QSslCertificate) {
	certificates_CArray := (*[0xffff]*QSslCertificate)(malloc(size_t(8 * len(certificates))))
	defer free(unsafe.Pointer(certificates_CArray))
	for i := range certificates {
		certificates_CArray[i] = certificates[i].cPointer()
	}
	certificates_ma := struct_miqt_array{len: size_t(len(certificates)), data: unsafe.Pointer(certificates_CArray)}
	QSslConfiguration_AddCaCertificatesWithCertificates(this.h, certificates_ma)
}

func QSslConfiguration_SystemCaCertificates() []QSslCertificate {
	var _ma struct_miqt_array = QSslConfiguration_SystemCaCertificates()
	_ret := make([]QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslCertificate(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) SetSslOption(option QSsl__SslOption, on bool) {
	QSslConfiguration_SetSslOption(this.h, (int)(option), (bool)(on))
}

func (this *QSslConfiguration) TestSslOption(option QSsl__SslOption) bool {
	return (bool)(QSslConfiguration_TestSslOption(this.h, (int)(option)))
}

func (this *QSslConfiguration) SessionTicket() []byte {
	var _bytearray struct_miqt_string = QSslConfiguration_SessionTicket(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslConfiguration) SetSessionTicket(sessionTicket []byte) {
	sessionTicket_alias := struct_miqt_string{}
	sessionTicket_alias.data = (char)(unsafe.Pointer(&sessionTicket[0]))
	sessionTicket_alias.len = size_t(len(sessionTicket))
	QSslConfiguration_SetSessionTicket(this.h, sessionTicket_alias)
}

func (this *QSslConfiguration) SessionTicketLifeTimeHint() int {
	return (int)(QSslConfiguration_SessionTicketLifeTimeHint(this.h))
}

func (this *QSslConfiguration) EphemeralServerKey() *QSslKey {
	_goptr := newQSslKey(QSslConfiguration_EphemeralServerKey(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslConfiguration) EllipticCurves() []QSslEllipticCurve {
	var _ma struct_miqt_array = QSslConfiguration_EllipticCurves(this.h)
	_ret := make([]QSslEllipticCurve, int(_ma.len))
	_outCast := (*[0xffff]*QSslEllipticCurve)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslEllipticCurve(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) SetEllipticCurves(curves []QSslEllipticCurve) {
	curves_CArray := (*[0xffff]*QSslEllipticCurve)(malloc(size_t(8 * len(curves))))
	defer free(unsafe.Pointer(curves_CArray))
	for i := range curves {
		curves_CArray[i] = curves[i].cPointer()
	}
	curves_ma := struct_miqt_array{len: size_t(len(curves)), data: unsafe.Pointer(curves_CArray)}
	QSslConfiguration_SetEllipticCurves(this.h, curves_ma)
}

func QSslConfiguration_SupportedEllipticCurves() []QSslEllipticCurve {
	var _ma struct_miqt_array = QSslConfiguration_SupportedEllipticCurves()
	_ret := make([]QSslEllipticCurve, int(_ma.len))
	_outCast := (*[0xffff]*QSslEllipticCurve)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSslEllipticCurve(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSslConfiguration) PreSharedKeyIdentityHint() []byte {
	var _bytearray struct_miqt_string = QSslConfiguration_PreSharedKeyIdentityHint(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslConfiguration) SetPreSharedKeyIdentityHint(hint []byte) {
	hint_alias := struct_miqt_string{}
	hint_alias.data = (char)(unsafe.Pointer(&hint[0]))
	hint_alias.len = size_t(len(hint))
	QSslConfiguration_SetPreSharedKeyIdentityHint(this.h, hint_alias)
}

func (this *QSslConfiguration) DiffieHellmanParameters() *QSslDiffieHellmanParameters {
	_goptr := newQSslDiffieHellmanParameters(QSslConfiguration_DiffieHellmanParameters(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslConfiguration) SetDiffieHellmanParameters(dhparams *QSslDiffieHellmanParameters) {
	QSslConfiguration_SetDiffieHellmanParameters(this.h, dhparams.cPointer())
}

func (this *QSslConfiguration) SetBackendConfigurationOption(name []byte, value *qt.QVariant) {
	name_alias := struct_miqt_string{}
	name_alias.data = (char)(unsafe.Pointer(&name[0]))
	name_alias.len = size_t(len(name))
	QSslConfiguration_SetBackendConfigurationOption(this.h, name_alias, (*QVariant)(value.UnsafePointer()))
}

func (this *QSslConfiguration) SetBackendConfiguration() {
	QSslConfiguration_SetBackendConfiguration(this.h)
}

func QSslConfiguration_DefaultConfiguration() *QSslConfiguration {
	_goptr := newQSslConfiguration(QSslConfiguration_DefaultConfiguration())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSslConfiguration_SetDefaultConfiguration(configuration *QSslConfiguration) {
	QSslConfiguration_SetDefaultConfiguration(configuration.cPointer())
}

func (this *QSslConfiguration) HandshakeMustInterruptOnError() bool {
	return (bool)(QSslConfiguration_HandshakeMustInterruptOnError(this.h))
}

func (this *QSslConfiguration) SetHandshakeMustInterruptOnError(interrupt bool) {
	QSslConfiguration_SetHandshakeMustInterruptOnError(this.h, (bool)(interrupt))
}

func (this *QSslConfiguration) MissingCertificateIsFatal() bool {
	return (bool)(QSslConfiguration_MissingCertificateIsFatal(this.h))
}

func (this *QSslConfiguration) SetMissingCertificateIsFatal(cannotRecover bool) {
	QSslConfiguration_SetMissingCertificateIsFatal(this.h, (bool)(cannotRecover))
}

func (this *QSslConfiguration) SetOcspStaplingEnabled(enable bool) {
	QSslConfiguration_SetOcspStaplingEnabled(this.h, (bool)(enable))
}

func (this *QSslConfiguration) OcspStaplingEnabled() bool {
	return (bool)(QSslConfiguration_OcspStaplingEnabled(this.h))
}

func (this *QSslConfiguration) SetAllowedNextProtocols(protocols [][]byte) {
	protocols_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(protocols))))
	defer free(unsafe.Pointer(protocols_CArray))
	for i := range protocols {
		protocols_i_alias := struct_miqt_string{}
		protocols_i_alias.data = (char)(unsafe.Pointer(&protocols[i][0]))
		protocols_i_alias.len = size_t(len(protocols[i]))
		protocols_CArray[i] = protocols_i_alias
	}
	protocols_ma := struct_miqt_array{len: size_t(len(protocols)), data: unsafe.Pointer(protocols_CArray)}
	QSslConfiguration_SetAllowedNextProtocols(this.h, protocols_ma)
}

func (this *QSslConfiguration) AllowedNextProtocols() [][]byte {
	var _ma struct_miqt_array = QSslConfiguration_AllowedNextProtocols(this.h)
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

func (this *QSslConfiguration) NextNegotiatedProtocol() []byte {
	var _bytearray struct_miqt_string = QSslConfiguration_NextNegotiatedProtocol(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslConfiguration) NextProtocolNegotiationStatus() NextProtocolNegotiationStatus {
	xxxxxxxxx
}

func (this *QSslConfiguration) AddCaCertificates2(path string, format QSsl__EncodingFormat) bool {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	return (bool)(QSslConfiguration_AddCaCertificates2(this.h, path_ms, (int)(format)))
}

func (this *QSslConfiguration) AddCaCertificates3(path string, format QSsl__EncodingFormat, syntax QSslCertificate__PatternSyntax) bool {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	return (bool)(QSslConfiguration_AddCaCertificates3(this.h, path_ms, (int)(format), (int)(syntax)))
}
