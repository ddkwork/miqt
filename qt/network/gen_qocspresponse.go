package network

import (
	"unsafe"
)

type QOcspCertificateStatus int

const (
	QOcspCertificateStatus__Good    QOcspCertificateStatus = 0
	QOcspCertificateStatus__Revoked QOcspCertificateStatus = 1
	QOcspCertificateStatus__Unknown QOcspCertificateStatus = 2
)

type QOcspRevocationReason int

const (
	QOcspRevocationReason__None                 QOcspRevocationReason = -1
	QOcspRevocationReason__Unspecified          QOcspRevocationReason = 0
	QOcspRevocationReason__KeyCompromise        QOcspRevocationReason = 1
	QOcspRevocationReason__CACompromise         QOcspRevocationReason = 2
	QOcspRevocationReason__AffiliationChanged   QOcspRevocationReason = 3
	QOcspRevocationReason__Superseded           QOcspRevocationReason = 4
	QOcspRevocationReason__CessationOfOperation QOcspRevocationReason = 5
	QOcspRevocationReason__CertificateHold      QOcspRevocationReason = 6
	QOcspRevocationReason__RemoveFromCRL        QOcspRevocationReason = 7
)

type QOcspResponse struct {
	h          uintptr
	isSubclass bool
}

// NewQOcspResponse constructs a new QOcspResponse object.
func NewQOcspResponse() *QOcspResponse {
	ret := newQOcspResponse(QOcspResponse_new())
	ret.isSubclass = true
	return ret
}

// NewQOcspResponse2 constructs a new QOcspResponse object.
func NewQOcspResponse2(other *QOcspResponse) *QOcspResponse {
	ret := newQOcspResponse(QOcspResponse_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QOcspResponse) OperatorAssign(other *QOcspResponse) {
	QOcspResponse_OperatorAssign(this.h, other.cPointer())
}

func (this *QOcspResponse) CertificateStatus() QOcspCertificateStatus {
	return (QOcspCertificateStatus)(QOcspResponse_CertificateStatus(this.h))
}

func (this *QOcspResponse) RevocationReason() QOcspRevocationReason {
	return (QOcspRevocationReason)(QOcspResponse_RevocationReason(this.h))
}

func (this *QOcspResponse) Responder() *QSslCertificate {
	_goptr := newQSslCertificate(QOcspResponse_Responder(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOcspResponse) Subject() *QSslCertificate {
	_goptr := newQSslCertificate(QOcspResponse_Subject(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOcspResponse) Swap(other *QOcspResponse) {
	QOcspResponse_Swap(this.h, other.cPointer())
}
