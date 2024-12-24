package qt

import (
	"unsafe"
)

type QObjectData__ int

const (
	QObjectData__CheckForParentChildLoopsWarnDepth QObjectData__ = 4096
)

type QObjectData struct {
	h          uintptr
	isSubclass bool
}

func (this *QObjectData) DynamicMetaObject() *QMetaObject {
	return newQMetaObject(QObjectData_DynamicMetaObject(this.h))
}

type QObject struct {
	h          uintptr
	isSubclass bool
}

// NewQObject constructs a new QObject object.
func NewQObject() *QObject {
	g := newQObject(QObject_new())
	g.isSubclass = true
	return g
}

// NewQObject2 constructs a new QObject object.
func NewQObject2(parent *QObject) *QObject {
	g := newQObject(QObject_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QObject) MetaObject() *QMetaObject {
	return newQMetaObject(QObject_MetaObject(this.h))
}

func (this *QObject) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QObject_Metacast(this.h, param1_Cstring))
}

func QObject_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QObject_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObject) Event(event *QEvent) bool {
	return (bool)(QObject_Event(this.h, event.cPointer()))
}

func (this *QObject) EventFilter(watched *QObject, event *QEvent) bool {
	return (bool)(QObject_EventFilter(this.h, watched.cPointer(), event.cPointer()))
}

func (this *QObject) ObjectName() string {
	var _ms struct_miqt_string = QObject_ObjectName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObject) SetObjectName(name QAnyStringView) {
	QObject_SetObjectName(this.h, name.cPointer())
}

func (this *QObject) IsWidgetType() bool {
	return (bool)(QObject_IsWidgetType(this.h))
}

func (this *QObject) IsWindowType() bool {
	return (bool)(QObject_IsWindowType(this.h))
}

func (this *QObject) IsQuickItemType() bool {
	return (bool)(QObject_IsQuickItemType(this.h))
}

func (this *QObject) SignalsBlocked() bool {
	return (bool)(QObject_SignalsBlocked(this.h))
}

func (this *QObject) BlockSignals(b bool) bool {
	return (bool)(QObject_BlockSignals(this.h, (bool)(b)))
}

func (this *QObject) Thread() *QThread {
	return newQThread(QObject_Thread(this.h))
}

func (this *QObject) MoveToThread(thread *QThread) bool {
	return (bool)(QObject_MoveToThread(this.h, thread.cPointer()))
}

func (this *QObject) StartTimer(interval int) int {
	return (int)(QObject_StartTimer(this.h, (int)(interval)))
}

func (this *QObject) KillTimer(id int) {
	QObject_KillTimer(this.h, (int)(id))
}

func (this *QObject) KillTimerWithId(id TimerId) {
	QObject_KillTimerWithId(this.h, (int)(id))
}

func (this *QObject) Children() []*QObject {
	var _ma struct_miqt_array = QObject_Children(this.h)
	_ret := make([]*QObject, int(_ma.len))
	_outCast := (*[0xffff]*QObject)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQObject(_outCast[i])
	}
	return _ret
}

func (this *QObject) SetParent(parent *QObject) {
	QObject_SetParent(this.h, parent.cPointer())
}

func (this *QObject) InstallEventFilter(filterObj *QObject) {
	QObject_InstallEventFilter(this.h, filterObj.cPointer())
}

func (this *QObject) RemoveEventFilter(obj *QObject) {
	QObject_RemoveEventFilter(this.h, obj.cPointer())
}

func QObject_Connect(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	_goptr := newQMetaObject__Connection(QObject_Connect(sender.cPointer(), signal.cPointer(), receiver.cPointer(), method.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) Connect2(sender *QObject, signal string, member string) *QMetaObject__Connection {
	signal_Cstring := CString(signal)
	defer free(unsafe.Pointer(signal_Cstring))
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	_goptr := newQMetaObject__Connection(QObject_Connect2(this.h, sender.cPointer(), signal_Cstring, member_Cstring))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QObject_Disconnect(sender *QObject, signal *QMetaMethod, receiver *QObject, member *QMetaMethod) bool {
	return (bool)(QObject_Disconnect(sender.cPointer(), signal.cPointer(), receiver.cPointer(), member.cPointer()))
}

func QObject_DisconnectWithQMetaObjectConnection(param1 *QMetaObject__Connection) bool {
	return (bool)(QObject_DisconnectWithQMetaObjectConnection(param1.cPointer()))
}

func (this *QObject) DumpObjectTree() {
	QObject_DumpObjectTree(this.h)
}

func (this *QObject) DumpObjectInfo() {
	QObject_DumpObjectInfo(this.h)
}

func (this *QObject) SetProperty(name string, value *QVariant) bool {
	name_Cstring := CString(name)
	defer free(unsafe.Pointer(name_Cstring))
	return (bool)(QObject_SetProperty(this.h, name_Cstring, value.cPointer()))
}

func (this *QObject) Property(name string) *QVariant {
	name_Cstring := CString(name)
	defer free(unsafe.Pointer(name_Cstring))
	_goptr := newQVariant(QObject_Property(this.h, name_Cstring))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) DynamicPropertyNames() [][]byte {
	var _ma struct_miqt_array = QObject_DynamicPropertyNames(this.h)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QObject) BindingStorage() *QBindingStorage {
	return newQBindingStorage(QObject_BindingStorage(this.h))
}

func (this *QObject) BindingStorage2() *QBindingStorage {
	return newQBindingStorage(QObject_BindingStorage2(this.h))
}

func (this *QObject) Destroyed() {
	QObject_Destroyed(this.h)
}

func (this *QObject) OnDestroyed(slot func()) {
	QObject_connect_Destroyed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObject_Destroyed
func miqt_exec_callback_QObject_Destroyed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QObject) Parent() *QObject {
	return newQObject(QObject_Parent(this.h))
}

func (this *QObject) Inherits(classname string) bool {
	classname_Cstring := CString(classname)
	defer free(unsafe.Pointer(classname_Cstring))
	return (bool)(QObject_Inherits(this.h, classname_Cstring))
}

func (this *QObject) DeleteLater() {
	QObject_DeleteLater(this.h)
}

func QObject_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QObject_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QObject_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QObject_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObject) MoveToThread2(thread *QThread, param2 Disambiguated_t) bool {
	return (bool)(QObject_MoveToThread2(this.h, thread.cPointer(), param2.cPointer()))
}

func (this *QObject) StartTimer2(interval int, timerType TimerType) int {
	return (int)(QObject_StartTimer2(this.h, (int)(interval), (int)(timerType)))
}

func QObject_Connect5(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, typeVal ConnectionType) *QMetaObject__Connection {
	_goptr := newQMetaObject__Connection(QObject_Connect5(sender.cPointer(), signal.cPointer(), receiver.cPointer(), method.cPointer(), (int)(typeVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) Connect4(sender *QObject, signal string, member string, typeVal ConnectionType) *QMetaObject__Connection {
	signal_Cstring := CString(signal)
	defer free(unsafe.Pointer(signal_Cstring))
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	_goptr := newQMetaObject__Connection(QObject_Connect4(this.h, sender.cPointer(), signal_Cstring, member_Cstring, (int)(typeVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QObject) Destroyed1(param1 *QObject) {
	QObject_Destroyed1(this.h, param1.cPointer())
}

func (this *QObject) OnDestroyed1(slot func(param1 *QObject)) {
	QObject_connect_Destroyed1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObject_Destroyed1
func miqt_exec_callback_QObject_Destroyed1(cb intptr_t, param1 *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	gofunc(slotval1)
}

func (this *QObject) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QObject_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QObject) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObject_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObject_MetaObject
func miqt_exec_callback_QObject_MetaObject(self QObject, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QObject{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QObject) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QObject_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QObject) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObject_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObject_Metacast
func miqt_exec_callback_QObject_Metacast(self QObject, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QObject{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QSignalBlocker struct {
	h          uintptr
	isSubclass bool
}

// NewQSignalBlocker constructs a new QSignalBlocker object.
func NewQSignalBlocker(o *QObject) *QSignalBlocker {
	g := newQSignalBlocker(QSignalBlocker_new(o.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSignalBlocker2 constructs a new QSignalBlocker object.
func NewQSignalBlocker2(o *QObject) *QSignalBlocker {
	g := newQSignalBlocker(QSignalBlocker_new2(o.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSignalBlocker) Reblock() {
	QSignalBlocker_Reblock(this.h)
}

func (this *QSignalBlocker) Unblock() {
	QSignalBlocker_Unblock(this.h)
}

func (this *QSignalBlocker) Dismiss() {
	QSignalBlocker_Dismiss(this.h)
}
