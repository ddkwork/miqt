package network

import (
	"unsafe"
)

type QSslPreSharedKeyAuthenticator struct {
	h          uintptr
	isSubclass bool
}

// NewQSslPreSharedKeyAuthenticator constructs a new QSslPreSharedKeyAuthenticator object.
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	ret := newQSslPreSharedKeyAuthenticator(QSslPreSharedKeyAuthenticator_new())
	ret.isSubclass = true
	return ret
}

// NewQSslPreSharedKeyAuthenticator2 constructs a new QSslPreSharedKeyAuthenticator object.
func NewQSslPreSharedKeyAuthenticator2(authenticator *QSslPreSharedKeyAuthenticator) *QSslPreSharedKeyAuthenticator {
	ret := newQSslPreSharedKeyAuthenticator(QSslPreSharedKeyAuthenticator_new2(authenticator.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSslPreSharedKeyAuthenticator) OperatorAssign(authenticator *QSslPreSharedKeyAuthenticator) {
	QSslPreSharedKeyAuthenticator_OperatorAssign(this.h, authenticator.cPointer())
}

func (this *QSslPreSharedKeyAuthenticator) Swap(other *QSslPreSharedKeyAuthenticator) {
	QSslPreSharedKeyAuthenticator_Swap(this.h, other.cPointer())
}

func (this *QSslPreSharedKeyAuthenticator) IdentityHint() []byte {
	var _bytearray struct_miqt_string = QSslPreSharedKeyAuthenticator_IdentityHint(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslPreSharedKeyAuthenticator) SetIdentity(identity []byte) {
	identity_alias := struct_miqt_string{}
	identity_alias.data = (char)(unsafe.Pointer(&identity[0]))
	identity_alias.len = size_t(len(identity))
	QSslPreSharedKeyAuthenticator_SetIdentity(this.h, identity_alias)
}

func (this *QSslPreSharedKeyAuthenticator) Identity() []byte {
	var _bytearray struct_miqt_string = QSslPreSharedKeyAuthenticator_Identity(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	return (int)(QSslPreSharedKeyAuthenticator_MaximumIdentityLength(this.h))
}

func (this *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey []byte) {
	preSharedKey_alias := struct_miqt_string{}
	preSharedKey_alias.data = (char)(unsafe.Pointer(&preSharedKey[0]))
	preSharedKey_alias.len = size_t(len(preSharedKey))
	QSslPreSharedKeyAuthenticator_SetPreSharedKey(this.h, preSharedKey_alias)
}

func (this *QSslPreSharedKeyAuthenticator) PreSharedKey() []byte {
	var _bytearray struct_miqt_string = QSslPreSharedKeyAuthenticator_PreSharedKey(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	return (int)(QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(this.h))
}
