package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSslDiffieHellmanParameters__Error int

const (
	QSslDiffieHellmanParameters__NoError               QSslDiffieHellmanParameters__Error = 0
	QSslDiffieHellmanParameters__InvalidInputDataError QSslDiffieHellmanParameters__Error = 1
	QSslDiffieHellmanParameters__UnsafeParametersError QSslDiffieHellmanParameters__Error = 2
)

type QSslDiffieHellmanParameters struct {
	h          uintptr
	isSubclass bool
}

// NewQSslDiffieHellmanParameters constructs a new QSslDiffieHellmanParameters object.
func NewQSslDiffieHellmanParameters() *QSslDiffieHellmanParameters {
	ret := newQSslDiffieHellmanParameters(QSslDiffieHellmanParameters_new())
	ret.isSubclass = true
	return ret
}

// NewQSslDiffieHellmanParameters2 constructs a new QSslDiffieHellmanParameters object.
func NewQSslDiffieHellmanParameters2(other *QSslDiffieHellmanParameters) *QSslDiffieHellmanParameters {
	ret := newQSslDiffieHellmanParameters(QSslDiffieHellmanParameters_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func QSslDiffieHellmanParameters_DefaultParameters() *QSslDiffieHellmanParameters {
	_goptr := newQSslDiffieHellmanParameters(QSslDiffieHellmanParameters_DefaultParameters())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslDiffieHellmanParameters) OperatorAssign(other *QSslDiffieHellmanParameters) {
	QSslDiffieHellmanParameters_OperatorAssign(this.h, other.cPointer())
}

func (this *QSslDiffieHellmanParameters) Swap(other *QSslDiffieHellmanParameters) {
	QSslDiffieHellmanParameters_Swap(this.h, other.cPointer())
}

func QSslDiffieHellmanParameters_FromEncoded(encoded []byte) *QSslDiffieHellmanParameters {
	encoded_alias := struct_miqt_string{}
	encoded_alias.data = (char)(unsafe.Pointer(&encoded[0]))
	encoded_alias.len = size_t(len(encoded))
	_goptr := newQSslDiffieHellmanParameters(QSslDiffieHellmanParameters_FromEncoded(encoded_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSslDiffieHellmanParameters_FromEncodedWithDevice(device *qt.QIODevice) *QSslDiffieHellmanParameters {
	_goptr := newQSslDiffieHellmanParameters(QSslDiffieHellmanParameters_FromEncodedWithDevice((*QIODevice)(device.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslDiffieHellmanParameters) IsEmpty() bool {
	return (bool)(QSslDiffieHellmanParameters_IsEmpty(this.h))
}

func (this *QSslDiffieHellmanParameters) IsValid() bool {
	return (bool)(QSslDiffieHellmanParameters_IsValid(this.h))
}

func (this *QSslDiffieHellmanParameters) Error() Error {
	xxxxxxxxx
}

func (this *QSslDiffieHellmanParameters) ErrorString() string {
	var _ms struct_miqt_string = QSslDiffieHellmanParameters_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSslDiffieHellmanParameters_FromEncoded2(encoded []byte, format QSsl__EncodingFormat) *QSslDiffieHellmanParameters {
	encoded_alias := struct_miqt_string{}
	encoded_alias.data = (char)(unsafe.Pointer(&encoded[0]))
	encoded_alias.len = size_t(len(encoded))
	_goptr := newQSslDiffieHellmanParameters(QSslDiffieHellmanParameters_FromEncoded2(encoded_alias, (int)(format)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSslDiffieHellmanParameters_FromEncoded22(device *qt.QIODevice, format QSsl__EncodingFormat) *QSslDiffieHellmanParameters {
	_goptr := newQSslDiffieHellmanParameters(QSslDiffieHellmanParameters_FromEncoded22((*QIODevice)(device.UnsafePointer()), (int)(format)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
