package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSslCertificateExtension struct {
	h          uintptr
	isSubclass bool
}

// NewQSslCertificateExtension constructs a new QSslCertificateExtension object.
func NewQSslCertificateExtension() *QSslCertificateExtension {
	g := newQSslCertificateExtension(QSslCertificateExtension_new())
	g.isSubclass = true
	return g
}

// NewQSslCertificateExtension2 constructs a new QSslCertificateExtension object.
func NewQSslCertificateExtension2(other *QSslCertificateExtension) *QSslCertificateExtension {
	g := newQSslCertificateExtension(QSslCertificateExtension_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSslCertificateExtension) OperatorAssign(other *QSslCertificateExtension) {
	QSslCertificateExtension_OperatorAssign(this.h, other.cPointer())
}

func (this *QSslCertificateExtension) Swap(other *QSslCertificateExtension) {
	QSslCertificateExtension_Swap(this.h, other.cPointer())
}

func (this *QSslCertificateExtension) Oid() string {
	var _ms struct_miqt_string = QSslCertificateExtension_Oid(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCertificateExtension) Name() string {
	var _ms struct_miqt_string = QSslCertificateExtension_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCertificateExtension) Value() *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QSslCertificateExtension_Value(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslCertificateExtension) IsCritical() bool {
	return (bool)(QSslCertificateExtension_IsCritical(this.h))
}

func (this *QSslCertificateExtension) IsSupported() bool {
	return (bool)(QSslCertificateExtension_IsSupported(this.h))
}
