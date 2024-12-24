package qt

import (
	"unsafe"
)

type QSystemSemaphore__AccessMode int

const (
	QSystemSemaphore__Open   QSystemSemaphore__AccessMode = 0
	QSystemSemaphore__Create QSystemSemaphore__AccessMode = 1
)

type QSystemSemaphore__SystemSemaphoreError int

const (
	QSystemSemaphore__NoError          QSystemSemaphore__SystemSemaphoreError = 0
	QSystemSemaphore__PermissionDenied QSystemSemaphore__SystemSemaphoreError = 1
	QSystemSemaphore__KeyError         QSystemSemaphore__SystemSemaphoreError = 2
	QSystemSemaphore__AlreadyExists    QSystemSemaphore__SystemSemaphoreError = 3
	QSystemSemaphore__NotFound         QSystemSemaphore__SystemSemaphoreError = 4
	QSystemSemaphore__OutOfResources   QSystemSemaphore__SystemSemaphoreError = 5
	QSystemSemaphore__UnknownError     QSystemSemaphore__SystemSemaphoreError = 6
)

type QSystemSemaphore struct {
	h          uintptr
	isSubclass bool
}

// NewQSystemSemaphore constructs a new QSystemSemaphore object.
func NewQSystemSemaphore(key *QNativeIpcKey) *QSystemSemaphore {
	ret := newQSystemSemaphore(QSystemSemaphore_new(key.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSystemSemaphore2 constructs a new QSystemSemaphore object.
func NewQSystemSemaphore2(key string) *QSystemSemaphore {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))

	ret := newQSystemSemaphore(QSystemSemaphore_new2(key_ms))
	ret.isSubclass = true
	return ret
}

// NewQSystemSemaphore3 constructs a new QSystemSemaphore object.
func NewQSystemSemaphore3(key *QNativeIpcKey, initialValue int) *QSystemSemaphore {
	ret := newQSystemSemaphore(QSystemSemaphore_new3(key.cPointer(), (int)(initialValue)))
	ret.isSubclass = true
	return ret
}

// NewQSystemSemaphore4 constructs a new QSystemSemaphore object.
func NewQSystemSemaphore4(key *QNativeIpcKey, initialValue int, param3 AccessMode) *QSystemSemaphore {
	ret := newQSystemSemaphore(QSystemSemaphore_new4(key.cPointer(), (int)(initialValue), param3))
	ret.isSubclass = true
	return ret
}

// NewQSystemSemaphore5 constructs a new QSystemSemaphore object.
func NewQSystemSemaphore5(key string, initialValue int) *QSystemSemaphore {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))

	ret := newQSystemSemaphore(QSystemSemaphore_new5(key_ms, (int)(initialValue)))
	ret.isSubclass = true
	return ret
}

// NewQSystemSemaphore6 constructs a new QSystemSemaphore object.
func NewQSystemSemaphore6(key string, initialValue int, mode AccessMode) *QSystemSemaphore {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))

	ret := newQSystemSemaphore(QSystemSemaphore_new6(key_ms, (int)(initialValue), mode))
	ret.isSubclass = true
	return ret
}

func QSystemSemaphore_Tr(sourceText string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	var _ms struct_miqt_string = QSystemSemaphore_Tr(sourceText_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSystemSemaphore) SetNativeKey(key *QNativeIpcKey) {
	QSystemSemaphore_SetNativeKey(this.h, key.cPointer())
}

func (this *QSystemSemaphore) SetNativeKeyWithKey(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSystemSemaphore_SetNativeKeyWithKey(this.h, key_ms)
}

func (this *QSystemSemaphore) NativeIpcKey() *QNativeIpcKey {
	_goptr := newQNativeIpcKey(QSystemSemaphore_NativeIpcKey(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSystemSemaphore) SetKey(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSystemSemaphore_SetKey(this.h, key_ms)
}

func (this *QSystemSemaphore) Key() string {
	var _ms struct_miqt_string = QSystemSemaphore_Key(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSystemSemaphore) Acquire() bool {
	return (bool)(QSystemSemaphore_Acquire(this.h))
}

func (this *QSystemSemaphore) Release() bool {
	return (bool)(QSystemSemaphore_Release(this.h))
}

func (this *QSystemSemaphore) Error() SystemSemaphoreError {
	xxxxxxxxx
}

func (this *QSystemSemaphore) ErrorString() string {
	var _ms struct_miqt_string = QSystemSemaphore_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSystemSemaphore_IsKeyTypeSupported(typeVal QNativeIpcKey__Type) bool {
	return (bool)(QSystemSemaphore_IsKeyTypeSupported((uint16_t)(typeVal)))
}

func QSystemSemaphore_PlatformSafeKey(key string) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSystemSemaphore_PlatformSafeKey(key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSystemSemaphore_LegacyNativeKey(key string) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSystemSemaphore_LegacyNativeKey(key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSystemSemaphore_Tr2(sourceText string, disambiguation string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QSystemSemaphore_Tr2(sourceText_Cstring, disambiguation_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSystemSemaphore_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QSystemSemaphore_Tr3(sourceText_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSystemSemaphore) SetNativeKey2(key *QNativeIpcKey, initialValue int) {
	QSystemSemaphore_SetNativeKey2(this.h, key.cPointer(), (int)(initialValue))
}

func (this *QSystemSemaphore) SetNativeKey3(key *QNativeIpcKey, initialValue int, param3 AccessMode) {
	QSystemSemaphore_SetNativeKey3(this.h, key.cPointer(), (int)(initialValue), param3)
}

func (this *QSystemSemaphore) SetNativeKey22(key string, initialValue int) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSystemSemaphore_SetNativeKey22(this.h, key_ms, (int)(initialValue))
}

func (this *QSystemSemaphore) SetNativeKey32(key string, initialValue int, mode AccessMode) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSystemSemaphore_SetNativeKey32(this.h, key_ms, (int)(initialValue), mode)
}

func (this *QSystemSemaphore) SetNativeKey4(key string, initialValue int, mode AccessMode, typeVal QNativeIpcKey__Type) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSystemSemaphore_SetNativeKey4(this.h, key_ms, (int)(initialValue), mode, (uint16_t)(typeVal))
}

func (this *QSystemSemaphore) SetKey2(key string, initialValue int) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSystemSemaphore_SetKey2(this.h, key_ms, (int)(initialValue))
}

func (this *QSystemSemaphore) SetKey3(key string, initialValue int, mode AccessMode) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSystemSemaphore_SetKey3(this.h, key_ms, (int)(initialValue), mode)
}

func (this *QSystemSemaphore) Release1(n int) bool {
	return (bool)(QSystemSemaphore_Release1(this.h, (int)(n)))
}

func QSystemSemaphore_PlatformSafeKey2(key string, typeVal QNativeIpcKey__Type) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSystemSemaphore_PlatformSafeKey2(key_ms, (uint16_t)(typeVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSystemSemaphore_LegacyNativeKey2(key string, typeVal QNativeIpcKey__Type) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSystemSemaphore_LegacyNativeKey2(key_ms, (uint16_t)(typeVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
