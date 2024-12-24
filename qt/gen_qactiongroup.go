package qt

import (
	"unsafe"
)

type QActionGroup__ExclusionPolicy int

const (
	QActionGroup__None              QActionGroup__ExclusionPolicy = 0
	QActionGroup__Exclusive         QActionGroup__ExclusionPolicy = 1
	QActionGroup__ExclusiveOptional QActionGroup__ExclusionPolicy = 2
)

type QActionGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQActionGroup constructs a new QActionGroup object.
func NewQActionGroup(parent *QObject) *QActionGroup {

	ret := newQActionGroup(QActionGroup_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QActionGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QActionGroup_MetaObject(this.h))
}

func (this *QActionGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QActionGroup_Metacast(this.h, param1_Cstring))
}

func QActionGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QActionGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QActionGroup) AddAction(a *QAction) *QAction {
	return newQAction(QActionGroup_AddAction(this.h, a.cPointer()))
}

func (this *QActionGroup) AddActionWithText(text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QActionGroup_AddActionWithText(this.h, text_ms))
}

func (this *QActionGroup) AddAction2(icon *QIcon, text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QActionGroup_AddAction2(this.h, icon.cPointer(), text_ms))
}

func (this *QActionGroup) RemoveAction(a *QAction) {
	QActionGroup_RemoveAction(this.h, a.cPointer())
}

func (this *QActionGroup) Actions() []*QAction {
	var _ma struct_miqt_array = QActionGroup_Actions(this.h)
	_ret := make([]*QAction, int(_ma.len))
	_outCast := (*[0xffff]*QAction)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAction(_outCast[i])
	}
	return _ret
}

func (this *QActionGroup) CheckedAction() *QAction {
	return newQAction(QActionGroup_CheckedAction(this.h))
}

func (this *QActionGroup) IsExclusive() bool {
	return (bool)(QActionGroup_IsExclusive(this.h))
}

func (this *QActionGroup) IsEnabled() bool {
	return (bool)(QActionGroup_IsEnabled(this.h))
}

func (this *QActionGroup) IsVisible() bool {
	return (bool)(QActionGroup_IsVisible(this.h))
}

func (this *QActionGroup) ExclusionPolicy() ExclusionPolicy {
	xxxxxxxxx
}

func (this *QActionGroup) SetEnabled(enabled bool) {
	QActionGroup_SetEnabled(this.h, (bool)(enabled))
}

func (this *QActionGroup) SetDisabled(b bool) {
	QActionGroup_SetDisabled(this.h, (bool)(b))
}

func (this *QActionGroup) SetVisible(visible bool) {
	QActionGroup_SetVisible(this.h, (bool)(visible))
}

func (this *QActionGroup) SetExclusive(exclusive bool) {
	QActionGroup_SetExclusive(this.h, (bool)(exclusive))
}

func (this *QActionGroup) SetExclusionPolicy(policy ExclusionPolicy) {
	QActionGroup_SetExclusionPolicy(this.h, policy)
}

func (this *QActionGroup) Triggered(param1 *QAction) {
	QActionGroup_Triggered(this.h, param1.cPointer())
}
func (this *QActionGroup) OnTriggered(slot func(param1 *QAction)) {
	QActionGroup_connect_Triggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_Triggered
func miqt_exec_callback_QActionGroup_Triggered(cb intptr_t, param1 *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(param1)

	gofunc(slotval1)
}

func (this *QActionGroup) Hovered(param1 *QAction) {
	QActionGroup_Hovered(this.h, param1.cPointer())
}
func (this *QActionGroup) OnHovered(slot func(param1 *QAction)) {
	QActionGroup_connect_Hovered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_Hovered
func miqt_exec_callback_QActionGroup_Hovered(cb intptr_t, param1 *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(param1)

	gofunc(slotval1)
}

func QActionGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QActionGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QActionGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QActionGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QActionGroup) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QActionGroup_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QActionGroup) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_Event
func miqt_exec_callback_QActionGroup_Event(self QActionGroup, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QActionGroup{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QActionGroup) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QActionGroup_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QActionGroup) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_EventFilter
func miqt_exec_callback_QActionGroup_EventFilter(self QActionGroup, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QActionGroup{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QActionGroup) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QActionGroup_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QActionGroup) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_TimerEvent
func miqt_exec_callback_QActionGroup_TimerEvent(self QActionGroup, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QActionGroup{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QActionGroup) callVirtualBase_ChildEvent(event *QChildEvent) {

	QActionGroup_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QActionGroup) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_ChildEvent
func miqt_exec_callback_QActionGroup_ChildEvent(self QActionGroup, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QActionGroup{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QActionGroup) callVirtualBase_CustomEvent(event *QEvent) {

	QActionGroup_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QActionGroup) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_CustomEvent
func miqt_exec_callback_QActionGroup_CustomEvent(self QActionGroup, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QActionGroup{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QActionGroup) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QActionGroup_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QActionGroup) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_ConnectNotify
func miqt_exec_callback_QActionGroup_ConnectNotify(self QActionGroup, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QActionGroup{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QActionGroup) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QActionGroup_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QActionGroup) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_DisconnectNotify
func miqt_exec_callback_QActionGroup_DisconnectNotify(self QActionGroup, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QActionGroup{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
