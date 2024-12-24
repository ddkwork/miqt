package qt

import (
	"unsafe"
)

type QSharedMemory__AccessMode int

const (
	QSharedMemory__ReadOnly  QSharedMemory__AccessMode = 0
	QSharedMemory__ReadWrite QSharedMemory__AccessMode = 1
)

type QSharedMemory__SharedMemoryError int

const (
	QSharedMemory__NoError          QSharedMemory__SharedMemoryError = 0
	QSharedMemory__PermissionDenied QSharedMemory__SharedMemoryError = 1
	QSharedMemory__InvalidSize      QSharedMemory__SharedMemoryError = 2
	QSharedMemory__KeyError         QSharedMemory__SharedMemoryError = 3
	QSharedMemory__AlreadyExists    QSharedMemory__SharedMemoryError = 4
	QSharedMemory__NotFound         QSharedMemory__SharedMemoryError = 5
	QSharedMemory__LockError        QSharedMemory__SharedMemoryError = 6
	QSharedMemory__OutOfResources   QSharedMemory__SharedMemoryError = 7
	QSharedMemory__UnknownError     QSharedMemory__SharedMemoryError = 8
)

type QSharedMemory struct {
	h          uintptr
	isSubclass bool
}

// NewQSharedMemory constructs a new QSharedMemory object.
func NewQSharedMemory() *QSharedMemory {

	ret := newQSharedMemory(QSharedMemory_new())
	ret.isSubclass = true
	return ret
}

// NewQSharedMemory2 constructs a new QSharedMemory object.
func NewQSharedMemory2(key *QNativeIpcKey) *QSharedMemory {

	ret := newQSharedMemory(QSharedMemory_new2(key.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSharedMemory3 constructs a new QSharedMemory object.
func NewQSharedMemory3(key string) *QSharedMemory {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))

	ret := newQSharedMemory(QSharedMemory_new3(key_ms))
	ret.isSubclass = true
	return ret
}

// NewQSharedMemory4 constructs a new QSharedMemory object.
func NewQSharedMemory4(parent *QObject) *QSharedMemory {

	ret := newQSharedMemory(QSharedMemory_new4(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSharedMemory5 constructs a new QSharedMemory object.
func NewQSharedMemory5(key *QNativeIpcKey, parent *QObject) *QSharedMemory {

	ret := newQSharedMemory(QSharedMemory_new5(key.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSharedMemory6 constructs a new QSharedMemory object.
func NewQSharedMemory6(key string, parent *QObject) *QSharedMemory {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))

	ret := newQSharedMemory(QSharedMemory_new6(key_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSharedMemory) MetaObject() *QMetaObject {
	return newQMetaObject(QSharedMemory_MetaObject(this.h))
}

func (this *QSharedMemory) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSharedMemory_Metacast(this.h, param1_Cstring))
}

func QSharedMemory_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSharedMemory_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSharedMemory) SetKey(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSharedMemory_SetKey(this.h, key_ms)
}

func (this *QSharedMemory) Key() string {
	var _ms struct_miqt_string = QSharedMemory_Key(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSharedMemory) SetNativeKey(key *QNativeIpcKey) {
	QSharedMemory_SetNativeKey(this.h, key.cPointer())
}

func (this *QSharedMemory) SetNativeKeyWithKey(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSharedMemory_SetNativeKeyWithKey(this.h, key_ms)
}

func (this *QSharedMemory) NativeKey() string {
	var _ms struct_miqt_string = QSharedMemory_NativeKey(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSharedMemory) NativeIpcKey() *QNativeIpcKey {
	_goptr := newQNativeIpcKey(QSharedMemory_NativeIpcKey(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSharedMemory) Create(size int64) bool {
	return (bool)(QSharedMemory_Create(this.h, (ptrdiff_t)(size)))
}

func (this *QSharedMemory) Size() int64 {
	return (int64)(QSharedMemory_Size(this.h))
}

func (this *QSharedMemory) Attach() bool {
	return (bool)(QSharedMemory_Attach(this.h))
}

func (this *QSharedMemory) IsAttached() bool {
	return (bool)(QSharedMemory_IsAttached(this.h))
}

func (this *QSharedMemory) Detach() bool {
	return (bool)(QSharedMemory_Detach(this.h))
}

func (this *QSharedMemory) Data() unsafe.Pointer {
	return (unsafe.Pointer)(QSharedMemory_Data(this.h))
}

func (this *QSharedMemory) ConstData() unsafe.Pointer {
	return (unsafe.Pointer)(QSharedMemory_ConstData(this.h))
}

func (this *QSharedMemory) Data2() unsafe.Pointer {
	return (unsafe.Pointer)(QSharedMemory_Data2(this.h))
}

func (this *QSharedMemory) Lock() bool {
	return (bool)(QSharedMemory_Lock(this.h))
}

func (this *QSharedMemory) Unlock() bool {
	return (bool)(QSharedMemory_Unlock(this.h))
}

func (this *QSharedMemory) Error() SharedMemoryError {
	xxxxxxxxx
}

func (this *QSharedMemory) ErrorString() string {
	var _ms struct_miqt_string = QSharedMemory_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSharedMemory_IsKeyTypeSupported(typeVal QNativeIpcKey__Type) bool {
	return (bool)(QSharedMemory_IsKeyTypeSupported((uint16_t)(typeVal)))
}

func QSharedMemory_PlatformSafeKey(key string) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSharedMemory_PlatformSafeKey(key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSharedMemory_LegacyNativeKey(key string) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSharedMemory_LegacyNativeKey(key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSharedMemory_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSharedMemory_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSharedMemory_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSharedMemory_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSharedMemory) SetNativeKey2(key string, typeVal QNativeIpcKey__Type) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QSharedMemory_SetNativeKey2(this.h, key_ms, (uint16_t)(typeVal))
}

func (this *QSharedMemory) Create2(size int64, mode AccessMode) bool {
	return (bool)(QSharedMemory_Create2(this.h, (ptrdiff_t)(size), mode))
}

func (this *QSharedMemory) Attach1(mode AccessMode) bool {
	return (bool)(QSharedMemory_Attach1(this.h, mode))
}

func QSharedMemory_PlatformSafeKey2(key string, typeVal QNativeIpcKey__Type) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSharedMemory_PlatformSafeKey2(key_ms, (uint16_t)(typeVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSharedMemory_LegacyNativeKey2(key string, typeVal QNativeIpcKey__Type) *QNativeIpcKey {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQNativeIpcKey(QSharedMemory_LegacyNativeKey2(key_ms, (uint16_t)(typeVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSharedMemory) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QSharedMemory_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QSharedMemory) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSharedMemory_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSharedMemory_Event
func miqt_exec_callback_QSharedMemory_Event(self QSharedMemory, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QSharedMemory{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSharedMemory) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QSharedMemory_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QSharedMemory) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSharedMemory_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSharedMemory_EventFilter
func miqt_exec_callback_QSharedMemory_EventFilter(self QSharedMemory, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QSharedMemory{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSharedMemory) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QSharedMemory_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSharedMemory) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSharedMemory_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSharedMemory_TimerEvent
func miqt_exec_callback_QSharedMemory_TimerEvent(self QSharedMemory, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QSharedMemory{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSharedMemory) callVirtualBase_ChildEvent(event *QChildEvent) {

	QSharedMemory_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSharedMemory) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSharedMemory_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSharedMemory_ChildEvent
func miqt_exec_callback_QSharedMemory_ChildEvent(self QSharedMemory, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QSharedMemory{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSharedMemory) callVirtualBase_CustomEvent(event *QEvent) {

	QSharedMemory_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSharedMemory) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSharedMemory_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSharedMemory_CustomEvent
func miqt_exec_callback_QSharedMemory_CustomEvent(self QSharedMemory, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSharedMemory{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QSharedMemory) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QSharedMemory_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSharedMemory) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSharedMemory_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSharedMemory_ConnectNotify
func miqt_exec_callback_QSharedMemory_ConnectNotify(self QSharedMemory, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSharedMemory{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QSharedMemory) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QSharedMemory_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSharedMemory) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSharedMemory_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSharedMemory_DisconnectNotify
func miqt_exec_callback_QSharedMemory_DisconnectNotify(self QSharedMemory, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSharedMemory{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
