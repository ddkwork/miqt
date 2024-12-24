package network

import (
	"unsafe"
)

type QSslCipher struct {
	h          uintptr
	isSubclass bool
}

// NewQSslCipher constructs a new QSslCipher object.
func NewQSslCipher() *QSslCipher {
	g := newQSslCipher(QSslCipher_new())
	g.isSubclass = true
	return g
}

// NewQSslCipher2 constructs a new QSslCipher object.
func NewQSslCipher2(name string) *QSslCipher {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	g := newQSslCipher(QSslCipher_new2(name_ms))
	g.isSubclass = true
	return g
}

// NewQSslCipher3 constructs a new QSslCipher object.
func NewQSslCipher3(name string, protocol QSsl__SslProtocol) *QSslCipher {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	g := newQSslCipher(QSslCipher_new3(name_ms, (int)(protocol)))
	g.isSubclass = true
	return g
}

// NewQSslCipher4 constructs a new QSslCipher object.
func NewQSslCipher4(other *QSslCipher) *QSslCipher {
	g := newQSslCipher(QSslCipher_new4(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSslCipher) OperatorAssign(other *QSslCipher) {
	QSslCipher_OperatorAssign(this.h, other.cPointer())
}

func (this *QSslCipher) Swap(other *QSslCipher) {
	QSslCipher_Swap(this.h, other.cPointer())
}

func (this *QSslCipher) OperatorEqual(other *QSslCipher) bool {
	return (bool)(QSslCipher_OperatorEqual(this.h, other.cPointer()))
}

func (this *QSslCipher) OperatorNotEqual(other *QSslCipher) bool {
	return (bool)(QSslCipher_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QSslCipher) IsNull() bool {
	return (bool)(QSslCipher_IsNull(this.h))
}

func (this *QSslCipher) Name() string {
	var _ms struct_miqt_string = QSslCipher_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCipher) SupportedBits() int {
	return (int)(QSslCipher_SupportedBits(this.h))
}

func (this *QSslCipher) UsedBits() int {
	return (int)(QSslCipher_UsedBits(this.h))
}

func (this *QSslCipher) KeyExchangeMethod() string {
	var _ms struct_miqt_string = QSslCipher_KeyExchangeMethod(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCipher) AuthenticationMethod() string {
	var _ms struct_miqt_string = QSslCipher_AuthenticationMethod(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCipher) EncryptionMethod() string {
	var _ms struct_miqt_string = QSslCipher_EncryptionMethod(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCipher) ProtocolString() string {
	var _ms struct_miqt_string = QSslCipher_ProtocolString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCipher) Protocol() QSsl__SslProtocol {
	return (QSsl__SslProtocol)(QSslCipher_Protocol(this.h))
}
