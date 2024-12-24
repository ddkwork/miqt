package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSslKey struct {
	h          uintptr
	isSubclass bool
}

// NewQSslKey constructs a new QSslKey object.
func NewQSslKey() *QSslKey {
	g := newQSslKey(QSslKey_new())
	g.isSubclass = true
	return g
}

// NewQSslKey2 constructs a new QSslKey object.
func NewQSslKey2(encoded []byte, algorithm QSsl__KeyAlgorithm) *QSslKey {
	encoded_alias := struct_miqt_string{}
	encoded_alias.data = (char)(unsafe.Pointer(&encoded[0]))
	encoded_alias.len = size_t(len(encoded))
	g := newQSslKey(QSslKey_new2(encoded_alias, (int)(algorithm)))
	g.isSubclass = true
	return g
}

// NewQSslKey3 constructs a new QSslKey object.
func NewQSslKey3(device *qt.QIODevice, algorithm QSsl__KeyAlgorithm) *QSslKey {
	g := newQSslKey(QSslKey_new3((*QIODevice)(device.UnsafePointer()), (int)(algorithm)))
	g.isSubclass = true
	return g
}

// NewQSslKey4 constructs a new QSslKey object.
func NewQSslKey4(handle unsafe.Pointer) *QSslKey {
	g := newQSslKey(QSslKey_new4(handle))
	g.isSubclass = true
	return g
}

// NewQSslKey5 constructs a new QSslKey object.
func NewQSslKey5(other *QSslKey) *QSslKey {
	g := newQSslKey(QSslKey_new5(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSslKey6 constructs a new QSslKey object.
func NewQSslKey6(encoded []byte, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat) *QSslKey {
	encoded_alias := struct_miqt_string{}
	encoded_alias.data = (char)(unsafe.Pointer(&encoded[0]))
	encoded_alias.len = size_t(len(encoded))
	g := newQSslKey(QSslKey_new6(encoded_alias, (int)(algorithm), (int)(format)))
	g.isSubclass = true
	return g
}

// NewQSslKey7 constructs a new QSslKey object.
func NewQSslKey7(encoded []byte, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, typeVal QSsl__KeyType) *QSslKey {
	encoded_alias := struct_miqt_string{}
	encoded_alias.data = (char)(unsafe.Pointer(&encoded[0]))
	encoded_alias.len = size_t(len(encoded))
	g := newQSslKey(QSslKey_new7(encoded_alias, (int)(algorithm), (int)(format), (int)(typeVal)))
	g.isSubclass = true
	return g
}

// NewQSslKey8 constructs a new QSslKey object.
func NewQSslKey8(encoded []byte, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, typeVal QSsl__KeyType, passPhrase []byte) *QSslKey {
	encoded_alias := struct_miqt_string{}
	encoded_alias.data = (char)(unsafe.Pointer(&encoded[0]))
	encoded_alias.len = size_t(len(encoded))
	passPhrase_alias := struct_miqt_string{}
	passPhrase_alias.data = (char)(unsafe.Pointer(&passPhrase[0]))
	passPhrase_alias.len = size_t(len(passPhrase))
	g := newQSslKey(QSslKey_new8(encoded_alias, (int)(algorithm), (int)(format), (int)(typeVal), passPhrase_alias))
	g.isSubclass = true
	return g
}

// NewQSslKey9 constructs a new QSslKey object.
func NewQSslKey9(device *qt.QIODevice, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat) *QSslKey {
	g := newQSslKey(QSslKey_new9((*QIODevice)(device.UnsafePointer()), (int)(algorithm), (int)(format)))
	g.isSubclass = true
	return g
}

// NewQSslKey10 constructs a new QSslKey object.
func NewQSslKey10(device *qt.QIODevice, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, typeVal QSsl__KeyType) *QSslKey {
	g := newQSslKey(QSslKey_new10((*QIODevice)(device.UnsafePointer()), (int)(algorithm), (int)(format), (int)(typeVal)))
	g.isSubclass = true
	return g
}

// NewQSslKey11 constructs a new QSslKey object.
func NewQSslKey11(device *qt.QIODevice, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, typeVal QSsl__KeyType, passPhrase []byte) *QSslKey {
	passPhrase_alias := struct_miqt_string{}
	passPhrase_alias.data = (char)(unsafe.Pointer(&passPhrase[0]))
	passPhrase_alias.len = size_t(len(passPhrase))
	g := newQSslKey(QSslKey_new11((*QIODevice)(device.UnsafePointer()), (int)(algorithm), (int)(format), (int)(typeVal), passPhrase_alias))
	g.isSubclass = true
	return g
}

// NewQSslKey12 constructs a new QSslKey object.
func NewQSslKey12(handle unsafe.Pointer, typeVal QSsl__KeyType) *QSslKey {
	g := newQSslKey(QSslKey_new12(handle, (int)(typeVal)))
	g.isSubclass = true
	return g
}

func (this *QSslKey) OperatorAssign(other *QSslKey) {
	QSslKey_OperatorAssign(this.h, other.cPointer())
}

func (this *QSslKey) Swap(other *QSslKey) {
	QSslKey_Swap(this.h, other.cPointer())
}

func (this *QSslKey) IsNull() bool {
	return (bool)(QSslKey_IsNull(this.h))
}

func (this *QSslKey) Clear() {
	QSslKey_Clear(this.h)
}

func (this *QSslKey) Length() int {
	return (int)(QSslKey_Length(this.h))
}

func (this *QSslKey) Type() QSsl__KeyType {
	return (QSsl__KeyType)(QSslKey_Type(this.h))
}

func (this *QSslKey) Algorithm() QSsl__KeyAlgorithm {
	return (QSsl__KeyAlgorithm)(QSslKey_Algorithm(this.h))
}

func (this *QSslKey) ToPem() []byte {
	var _bytearray struct_miqt_string = QSslKey_ToPem(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslKey) ToDer() []byte {
	var _bytearray struct_miqt_string = QSslKey_ToDer(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslKey) Handle() unsafe.Pointer {
	return (unsafe.Pointer)(QSslKey_Handle(this.h))
}

func (this *QSslKey) OperatorEqual(key *QSslKey) bool {
	return (bool)(QSslKey_OperatorEqual(this.h, key.cPointer()))
}

func (this *QSslKey) OperatorNotEqual(key *QSslKey) bool {
	return (bool)(QSslKey_OperatorNotEqual(this.h, key.cPointer()))
}

func (this *QSslKey) ToPem1(passPhrase []byte) []byte {
	passPhrase_alias := struct_miqt_string{}
	passPhrase_alias.data = (char)(unsafe.Pointer(&passPhrase[0]))
	passPhrase_alias.len = size_t(len(passPhrase))
	var _bytearray struct_miqt_string = QSslKey_ToPem1(this.h, passPhrase_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslKey) ToDer1(passPhrase []byte) []byte {
	passPhrase_alias := struct_miqt_string{}
	passPhrase_alias.data = (char)(unsafe.Pointer(&passPhrase[0]))
	passPhrase_alias.len = size_t(len(passPhrase))
	var _bytearray struct_miqt_string = QSslKey_ToDer1(this.h, passPhrase_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}
