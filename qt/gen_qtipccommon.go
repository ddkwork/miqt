package qt

import (
	"unsafe"
)

type QNativeIpcKey__Type uint16

const (
	QNativeIpcKey__SystemV       QNativeIpcKey__Type = 81
	QNativeIpcKey__PosixRealtime QNativeIpcKey__Type = 256
	QNativeIpcKey__Windows       QNativeIpcKey__Type = 257
)

type QNativeIpcKey struct {
	h          uintptr
	isSubclass bool
}

// NewQNativeIpcKey constructs a new QNativeIpcKey object.
func NewQNativeIpcKey() *QNativeIpcKey {
	ret := newQNativeIpcKey(QNativeIpcKey_new())
	ret.isSubclass = true
	return ret
}

// NewQNativeIpcKey2 constructs a new QNativeIpcKey object.
func NewQNativeIpcKey2(typeVal Type) *QNativeIpcKey {
	ret := newQNativeIpcKey(QNativeIpcKey_new2(typeVal))
	ret.isSubclass = true
	return ret
}

// NewQNativeIpcKey3 constructs a new QNativeIpcKey object.
func NewQNativeIpcKey3(k string) *QNativeIpcKey {
	k_ms := struct_miqt_string{}
	k_ms.data = CString(k)
	k_ms.len = size_t(len(k))
	defer free(unsafe.Pointer(k_ms.data))

	ret := newQNativeIpcKey(QNativeIpcKey_new3(k_ms))
	ret.isSubclass = true
	return ret
}

// NewQNativeIpcKey4 constructs a new QNativeIpcKey object.
func NewQNativeIpcKey4(other *QNativeIpcKey) *QNativeIpcKey {
	ret := newQNativeIpcKey(QNativeIpcKey_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQNativeIpcKey5 constructs a new QNativeIpcKey object.
func NewQNativeIpcKey5(k string, typeVal Type) *QNativeIpcKey {
	k_ms := struct_miqt_string{}
	k_ms.data = CString(k)
	k_ms.len = size_t(len(k))
	defer free(unsafe.Pointer(k_ms.data))

	ret := newQNativeIpcKey(QNativeIpcKey_new5(k_ms, typeVal))
	ret.isSubclass = true
	return ret
}

func QNativeIpcKey_LegacyDefaultTypeForOs() Type {
	xxxxxxxxx
}

func (this *QNativeIpcKey) OperatorAssign(other *QNativeIpcKey) {
	QNativeIpcKey_OperatorAssign(this.h, other.cPointer())
}

func (this *QNativeIpcKey) Swap(other *QNativeIpcKey) {
	QNativeIpcKey_Swap(this.h, other.cPointer())
}

func (this *QNativeIpcKey) IsEmpty() bool {
	return (bool)(QNativeIpcKey_IsEmpty(this.h))
}

func (this *QNativeIpcKey) IsValid() bool {
	return (bool)(QNativeIpcKey_IsValid(this.h))
}

func (this *QNativeIpcKey) Type() Type {
	xxxxxxxxx
}

func (this *QNativeIpcKey) SetType(typeVal Type) {
	QNativeIpcKey_SetType(this.h, typeVal)
}

func (this *QNativeIpcKey) NativeKey() string {
	var _ms struct_miqt_string = QNativeIpcKey_NativeKey(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNativeIpcKey) SetNativeKey(newKey string) {
	newKey_ms := struct_miqt_string{}
	newKey_ms.data = CString(newKey)
	newKey_ms.len = size_t(len(newKey))
	defer free(unsafe.Pointer(newKey_ms.data))
	QNativeIpcKey_SetNativeKey(this.h, newKey_ms)
}

func (this *QNativeIpcKey) ToString() string {
	var _ms struct_miqt_string = QNativeIpcKey_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNativeIpcKey_FromString(stringVal string) *QNativeIpcKey {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQNativeIpcKey(QNativeIpcKey_FromString(stringVal_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
