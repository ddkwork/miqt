package network

import (
	"unsafe"
)

type QSslEllipticCurve struct {
	h          uintptr
	isSubclass bool
}

// NewQSslEllipticCurve constructs a new QSslEllipticCurve object.
func NewQSslEllipticCurve() *QSslEllipticCurve {

	ret := newQSslEllipticCurve(QSslEllipticCurve_new())
	ret.isSubclass = true
	return ret
}

// NewQSslEllipticCurve2 constructs a new QSslEllipticCurve object.
func NewQSslEllipticCurve2(param1 *QSslEllipticCurve) *QSslEllipticCurve {

	ret := newQSslEllipticCurve(QSslEllipticCurve_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func QSslEllipticCurve_FromShortName(name string) *QSslEllipticCurve {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQSslEllipticCurve(QSslEllipticCurve_FromShortName(name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSslEllipticCurve_FromLongName(name string) *QSslEllipticCurve {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQSslEllipticCurve(QSslEllipticCurve_FromLongName(name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslEllipticCurve) ShortName() string {
	var _ms struct_miqt_string = QSslEllipticCurve_ShortName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslEllipticCurve) LongName() string {
	var _ms struct_miqt_string = QSslEllipticCurve_LongName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslEllipticCurve) IsValid() bool {
	return (bool)(QSslEllipticCurve_IsValid(this.h))
}

func (this *QSslEllipticCurve) IsTlsNamedCurve() bool {
	return (bool)(QSslEllipticCurve_IsTlsNamedCurve(this.h))
}
