package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAuthenticator struct {
	h          uintptr
	isSubclass bool
}

// NewQAuthenticator constructs a new QAuthenticator object.
func NewQAuthenticator() *QAuthenticator {
	ret := newQAuthenticator(QAuthenticator_new())
	ret.isSubclass = true
	return ret
}

// NewQAuthenticator2 constructs a new QAuthenticator object.
func NewQAuthenticator2(other *QAuthenticator) *QAuthenticator {
	ret := newQAuthenticator(QAuthenticator_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAuthenticator) OperatorAssign(other *QAuthenticator) {
	QAuthenticator_OperatorAssign(this.h, other.cPointer())
}

func (this *QAuthenticator) OperatorEqual(other *QAuthenticator) bool {
	return (bool)(QAuthenticator_OperatorEqual(this.h, other.cPointer()))
}

func (this *QAuthenticator) OperatorNotEqual(other *QAuthenticator) bool {
	return (bool)(QAuthenticator_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QAuthenticator) User() string {
	var _ms struct_miqt_string = QAuthenticator_User(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAuthenticator) SetUser(user string) {
	user_ms := struct_miqt_string{}
	user_ms.data = CString(user)
	user_ms.len = size_t(len(user))
	defer free(unsafe.Pointer(user_ms.data))
	QAuthenticator_SetUser(this.h, user_ms)
}

func (this *QAuthenticator) Password() string {
	var _ms struct_miqt_string = QAuthenticator_Password(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAuthenticator) SetPassword(password string) {
	password_ms := struct_miqt_string{}
	password_ms.data = CString(password)
	password_ms.len = size_t(len(password))
	defer free(unsafe.Pointer(password_ms.data))
	QAuthenticator_SetPassword(this.h, password_ms)
}

func (this *QAuthenticator) Realm() string {
	var _ms struct_miqt_string = QAuthenticator_Realm(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAuthenticator) SetRealm(realm string) {
	realm_ms := struct_miqt_string{}
	realm_ms.data = CString(realm)
	realm_ms.len = size_t(len(realm))
	defer free(unsafe.Pointer(realm_ms.data))
	QAuthenticator_SetRealm(this.h, realm_ms)
}

func (this *QAuthenticator) Option(opt string) *qt.QVariant {
	opt_ms := struct_miqt_string{}
	opt_ms.data = CString(opt)
	opt_ms.len = size_t(len(opt))
	defer free(unsafe.Pointer(opt_ms.data))
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QAuthenticator_Option(this.h, opt_ms)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAuthenticator) Options() map[string]qt.QVariant {
	var _mm struct_miqt_map = QAuthenticator_Options(this.h)
	_ret := make(map[string]qt.QVariant, int(_mm.len))
	_Keys := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		var _hashkey_ms struct_miqt_string = _Keys[i]
		_hashkey_ret := GoStringN(_hashkey_ms.data, int(int64(_hashkey_ms.len)))
		free(unsafe.Pointer(_hashkey_ms.data))
		_entry_Key := _hashkey_ret
		_hashval_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(_Values[i]))
		_hashval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_hashval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QAuthenticator) SetOption(opt string, value *qt.QVariant) {
	opt_ms := struct_miqt_string{}
	opt_ms.data = CString(opt)
	opt_ms.len = size_t(len(opt))
	defer free(unsafe.Pointer(opt_ms.data))
	QAuthenticator_SetOption(this.h, opt_ms, (*QVariant)(value.UnsafePointer()))
}

func (this *QAuthenticator) IsNull() bool {
	return (bool)(QAuthenticator_IsNull(this.h))
}

func (this *QAuthenticator) Detach() {
	QAuthenticator_Detach(this.h)
}
