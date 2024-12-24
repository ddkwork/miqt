package qt

import (
	"unsafe"
)

type QShortcut struct {
	h          uintptr
	isSubclass bool
}

// NewQShortcut constructs a new QShortcut object.
func NewQShortcut(parent *QObject) *QShortcut {

	ret := newQShortcut(QShortcut_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQShortcut2 constructs a new QShortcut object.
func NewQShortcut2(key *QKeySequence, parent *QObject) *QShortcut {

	ret := newQShortcut(QShortcut_new2(key.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQShortcut3 constructs a new QShortcut object.
func NewQShortcut3(key QKeySequence__StandardKey, parent *QObject) *QShortcut {

	ret := newQShortcut(QShortcut_new3((int)(key), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQShortcut4 constructs a new QShortcut object.
func NewQShortcut4(key *QKeySequence, parent *QObject, member string) *QShortcut {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))

	ret := newQShortcut(QShortcut_new4(key.cPointer(), parent.cPointer(), member_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQShortcut5 constructs a new QShortcut object.
func NewQShortcut5(key *QKeySequence, parent *QObject, member string, ambiguousMember string) *QShortcut {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	ambiguousMember_Cstring := CString(ambiguousMember)
	defer free(unsafe.Pointer(ambiguousMember_Cstring))

	ret := newQShortcut(QShortcut_new5(key.cPointer(), parent.cPointer(), member_Cstring, ambiguousMember_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQShortcut6 constructs a new QShortcut object.
func NewQShortcut6(key *QKeySequence, parent *QObject, member string, ambiguousMember string, context ShortcutContext) *QShortcut {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	ambiguousMember_Cstring := CString(ambiguousMember)
	defer free(unsafe.Pointer(ambiguousMember_Cstring))

	ret := newQShortcut(QShortcut_new6(key.cPointer(), parent.cPointer(), member_Cstring, ambiguousMember_Cstring, (int)(context)))
	ret.isSubclass = true
	return ret
}

// NewQShortcut7 constructs a new QShortcut object.
func NewQShortcut7(key QKeySequence__StandardKey, parent *QObject, member string) *QShortcut {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))

	ret := newQShortcut(QShortcut_new7((int)(key), parent.cPointer(), member_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQShortcut8 constructs a new QShortcut object.
func NewQShortcut8(key QKeySequence__StandardKey, parent *QObject, member string, ambiguousMember string) *QShortcut {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	ambiguousMember_Cstring := CString(ambiguousMember)
	defer free(unsafe.Pointer(ambiguousMember_Cstring))

	ret := newQShortcut(QShortcut_new8((int)(key), parent.cPointer(), member_Cstring, ambiguousMember_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQShortcut9 constructs a new QShortcut object.
func NewQShortcut9(key QKeySequence__StandardKey, parent *QObject, member string, ambiguousMember string, context ShortcutContext) *QShortcut {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	ambiguousMember_Cstring := CString(ambiguousMember)
	defer free(unsafe.Pointer(ambiguousMember_Cstring))

	ret := newQShortcut(QShortcut_new9((int)(key), parent.cPointer(), member_Cstring, ambiguousMember_Cstring, (int)(context)))
	ret.isSubclass = true
	return ret
}

func (this *QShortcut) MetaObject() *QMetaObject {
	return newQMetaObject(QShortcut_MetaObject(this.h))
}

func (this *QShortcut) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QShortcut_Metacast(this.h, param1_Cstring))
}

func QShortcut_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QShortcut_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QShortcut) SetKey(key *QKeySequence) {
	QShortcut_SetKey(this.h, key.cPointer())
}

func (this *QShortcut) Key() *QKeySequence {
	_goptr := newQKeySequence(QShortcut_Key(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QShortcut) SetKeys(key QKeySequence__StandardKey) {
	QShortcut_SetKeys(this.h, (int)(key))
}

func (this *QShortcut) SetKeysWithKeys(keys []QKeySequence) {
	keys_CArray := (*[0xffff]*QKeySequence)(malloc(size_t(8 * len(keys))))
	defer free(unsafe.Pointer(keys_CArray))
	for i := range keys {
		keys_CArray[i] = keys[i].cPointer()
	}
	keys_ma := struct_miqt_array{len: size_t(len(keys)), data: unsafe.Pointer(keys_CArray)}
	QShortcut_SetKeysWithKeys(this.h, keys_ma)
}

func (this *QShortcut) Keys() []QKeySequence {
	var _ma struct_miqt_array = QShortcut_Keys(this.h)
	_ret := make([]QKeySequence, int(_ma.len))
	_outCast := (*[0xffff]*QKeySequence)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQKeySequence(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QShortcut) SetEnabled(enable bool) {
	QShortcut_SetEnabled(this.h, (bool)(enable))
}

func (this *QShortcut) IsEnabled() bool {
	return (bool)(QShortcut_IsEnabled(this.h))
}

func (this *QShortcut) SetContext(context ShortcutContext) {
	QShortcut_SetContext(this.h, (int)(context))
}

func (this *QShortcut) Context() ShortcutContext {
	return (ShortcutContext)(QShortcut_Context(this.h))
}

func (this *QShortcut) SetAutoRepeat(on bool) {
	QShortcut_SetAutoRepeat(this.h, (bool)(on))
}

func (this *QShortcut) AutoRepeat() bool {
	return (bool)(QShortcut_AutoRepeat(this.h))
}

func (this *QShortcut) Id() int {
	return (int)(QShortcut_Id(this.h))
}

func (this *QShortcut) SetWhatsThis(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QShortcut_SetWhatsThis(this.h, text_ms)
}

func (this *QShortcut) WhatsThis() string {
	var _ms struct_miqt_string = QShortcut_WhatsThis(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QShortcut) Activated() {
	QShortcut_Activated(this.h)
}
func (this *QShortcut) OnActivated(slot func()) {
	QShortcut_connect_Activated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_Activated
func miqt_exec_callback_QShortcut_Activated(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QShortcut) ActivatedAmbiguously() {
	QShortcut_ActivatedAmbiguously(this.h)
}
func (this *QShortcut) OnActivatedAmbiguously(slot func()) {
	QShortcut_connect_ActivatedAmbiguously(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_ActivatedAmbiguously
func miqt_exec_callback_QShortcut_ActivatedAmbiguously(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QShortcut_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QShortcut_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QShortcut_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QShortcut_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QShortcut) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QShortcut_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QShortcut) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcut_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_Event
func miqt_exec_callback_QShortcut_Event(self QShortcut, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QShortcut{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QShortcut) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QShortcut_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QShortcut) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcut_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_EventFilter
func miqt_exec_callback_QShortcut_EventFilter(self QShortcut, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QShortcut{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QShortcut) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QShortcut_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QShortcut) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcut_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_TimerEvent
func miqt_exec_callback_QShortcut_TimerEvent(self QShortcut, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QShortcut{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QShortcut) callVirtualBase_ChildEvent(event *QChildEvent) {

	QShortcut_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QShortcut) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcut_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_ChildEvent
func miqt_exec_callback_QShortcut_ChildEvent(self QShortcut, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QShortcut{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QShortcut) callVirtualBase_CustomEvent(event *QEvent) {

	QShortcut_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QShortcut) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcut_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_CustomEvent
func miqt_exec_callback_QShortcut_CustomEvent(self QShortcut, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QShortcut{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QShortcut) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QShortcut_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QShortcut) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcut_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_ConnectNotify
func miqt_exec_callback_QShortcut_ConnectNotify(self QShortcut, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QShortcut{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QShortcut) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QShortcut_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QShortcut) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcut_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcut_DisconnectNotify
func miqt_exec_callback_QShortcut_DisconnectNotify(self QShortcut, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QShortcut{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
