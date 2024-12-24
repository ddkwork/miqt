package qt

import (
	"unsafe"
)

type QSystemTrayIcon__ActivationReason int

const (
	QSystemTrayIcon__Unknown     QSystemTrayIcon__ActivationReason = 0
	QSystemTrayIcon__Context     QSystemTrayIcon__ActivationReason = 1
	QSystemTrayIcon__DoubleClick QSystemTrayIcon__ActivationReason = 2
	QSystemTrayIcon__Trigger     QSystemTrayIcon__ActivationReason = 3
	QSystemTrayIcon__MiddleClick QSystemTrayIcon__ActivationReason = 4
)

type QSystemTrayIcon__MessageIcon int

const (
	QSystemTrayIcon__NoIcon      QSystemTrayIcon__MessageIcon = 0
	QSystemTrayIcon__Information QSystemTrayIcon__MessageIcon = 1
	QSystemTrayIcon__Warning     QSystemTrayIcon__MessageIcon = 2
	QSystemTrayIcon__Critical    QSystemTrayIcon__MessageIcon = 3
)

type QSystemTrayIcon struct {
	h          uintptr
	isSubclass bool
}

// NewQSystemTrayIcon constructs a new QSystemTrayIcon object.
func NewQSystemTrayIcon() *QSystemTrayIcon {

	ret := newQSystemTrayIcon(QSystemTrayIcon_new())
	ret.isSubclass = true
	return ret
}

// NewQSystemTrayIcon2 constructs a new QSystemTrayIcon object.
func NewQSystemTrayIcon2(icon *QIcon) *QSystemTrayIcon {

	ret := newQSystemTrayIcon(QSystemTrayIcon_new2(icon.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSystemTrayIcon3 constructs a new QSystemTrayIcon object.
func NewQSystemTrayIcon3(parent *QObject) *QSystemTrayIcon {

	ret := newQSystemTrayIcon(QSystemTrayIcon_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSystemTrayIcon4 constructs a new QSystemTrayIcon object.
func NewQSystemTrayIcon4(icon *QIcon, parent *QObject) *QSystemTrayIcon {

	ret := newQSystemTrayIcon(QSystemTrayIcon_new4(icon.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSystemTrayIcon) MetaObject() *QMetaObject {
	return newQMetaObject(QSystemTrayIcon_MetaObject(this.h))
}

func (this *QSystemTrayIcon) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSystemTrayIcon_Metacast(this.h, param1_Cstring))
}

func QSystemTrayIcon_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSystemTrayIcon_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSystemTrayIcon) SetContextMenu(menu *QMenu) {
	QSystemTrayIcon_SetContextMenu(this.h, menu.cPointer())
}

func (this *QSystemTrayIcon) ContextMenu() *QMenu {
	return newQMenu(QSystemTrayIcon_ContextMenu(this.h))
}

func (this *QSystemTrayIcon) Icon() *QIcon {
	_goptr := newQIcon(QSystemTrayIcon_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSystemTrayIcon) SetIcon(icon *QIcon) {
	QSystemTrayIcon_SetIcon(this.h, icon.cPointer())
}

func (this *QSystemTrayIcon) ToolTip() string {
	var _ms struct_miqt_string = QSystemTrayIcon_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSystemTrayIcon) SetToolTip(tip string) {
	tip_ms := struct_miqt_string{}
	tip_ms.data = CString(tip)
	tip_ms.len = size_t(len(tip))
	defer free(unsafe.Pointer(tip_ms.data))
	QSystemTrayIcon_SetToolTip(this.h, tip_ms)
}

func QSystemTrayIcon_IsSystemTrayAvailable() bool {
	return (bool)(QSystemTrayIcon_IsSystemTrayAvailable())
}

func QSystemTrayIcon_SupportsMessages() bool {
	return (bool)(QSystemTrayIcon_SupportsMessages())
}

func (this *QSystemTrayIcon) Geometry() *QRect {
	_goptr := newQRect(QSystemTrayIcon_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSystemTrayIcon) IsVisible() bool {
	return (bool)(QSystemTrayIcon_IsVisible(this.h))
}

func (this *QSystemTrayIcon) SetVisible(visible bool) {
	QSystemTrayIcon_SetVisible(this.h, (bool)(visible))
}

func (this *QSystemTrayIcon) Show() {
	QSystemTrayIcon_Show(this.h)
}

func (this *QSystemTrayIcon) Hide() {
	QSystemTrayIcon_Hide(this.h)
}

func (this *QSystemTrayIcon) ShowMessage(title string, msg string, icon *QIcon) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	msg_ms := struct_miqt_string{}
	msg_ms.data = CString(msg)
	msg_ms.len = size_t(len(msg))
	defer free(unsafe.Pointer(msg_ms.data))
	QSystemTrayIcon_ShowMessage(this.h, title_ms, msg_ms, icon.cPointer())
}

func (this *QSystemTrayIcon) ShowMessage2(title string, msg string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	msg_ms := struct_miqt_string{}
	msg_ms.data = CString(msg)
	msg_ms.len = size_t(len(msg))
	defer free(unsafe.Pointer(msg_ms.data))
	QSystemTrayIcon_ShowMessage2(this.h, title_ms, msg_ms)
}

func (this *QSystemTrayIcon) Activated(reason QSystemTrayIcon__ActivationReason) {
	QSystemTrayIcon_Activated(this.h, (int)(reason))
}
func (this *QSystemTrayIcon) OnActivated(slot func(reason QSystemTrayIcon__ActivationReason)) {
	QSystemTrayIcon_connect_Activated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_Activated
func miqt_exec_callback_QSystemTrayIcon_Activated(cb intptr_t, reason int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(reason QSystemTrayIcon__ActivationReason))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QSystemTrayIcon__ActivationReason)(reason)

	gofunc(slotval1)
}

func (this *QSystemTrayIcon) MessageClicked() {
	QSystemTrayIcon_MessageClicked(this.h)
}
func (this *QSystemTrayIcon) OnMessageClicked(slot func()) {
	QSystemTrayIcon_connect_MessageClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_MessageClicked
func miqt_exec_callback_QSystemTrayIcon_MessageClicked(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QSystemTrayIcon_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSystemTrayIcon_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSystemTrayIcon_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSystemTrayIcon_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSystemTrayIcon) ShowMessage4(title string, msg string, icon *QIcon, msecs int) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	msg_ms := struct_miqt_string{}
	msg_ms.data = CString(msg)
	msg_ms.len = size_t(len(msg))
	defer free(unsafe.Pointer(msg_ms.data))
	QSystemTrayIcon_ShowMessage4(this.h, title_ms, msg_ms, icon.cPointer(), (int)(msecs))
}

func (this *QSystemTrayIcon) ShowMessage3(title string, msg string, icon QSystemTrayIcon__MessageIcon) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	msg_ms := struct_miqt_string{}
	msg_ms.data = CString(msg)
	msg_ms.len = size_t(len(msg))
	defer free(unsafe.Pointer(msg_ms.data))
	QSystemTrayIcon_ShowMessage3(this.h, title_ms, msg_ms, (int)(icon))
}

func (this *QSystemTrayIcon) ShowMessage42(title string, msg string, icon QSystemTrayIcon__MessageIcon, msecs int) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	msg_ms := struct_miqt_string{}
	msg_ms.data = CString(msg)
	msg_ms.len = size_t(len(msg))
	defer free(unsafe.Pointer(msg_ms.data))
	QSystemTrayIcon_ShowMessage42(this.h, title_ms, msg_ms, (int)(icon), (int)(msecs))
}

func (this *QSystemTrayIcon) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QSystemTrayIcon_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QSystemTrayIcon) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSystemTrayIcon_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_Event
func miqt_exec_callback_QSystemTrayIcon_Event(self QSystemTrayIcon, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QSystemTrayIcon{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSystemTrayIcon) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QSystemTrayIcon_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QSystemTrayIcon) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSystemTrayIcon_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_EventFilter
func miqt_exec_callback_QSystemTrayIcon_EventFilter(self QSystemTrayIcon, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QSystemTrayIcon{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSystemTrayIcon) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QSystemTrayIcon_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSystemTrayIcon) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSystemTrayIcon_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_TimerEvent
func miqt_exec_callback_QSystemTrayIcon_TimerEvent(self QSystemTrayIcon, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QSystemTrayIcon{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSystemTrayIcon) callVirtualBase_ChildEvent(event *QChildEvent) {

	QSystemTrayIcon_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSystemTrayIcon) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSystemTrayIcon_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_ChildEvent
func miqt_exec_callback_QSystemTrayIcon_ChildEvent(self QSystemTrayIcon, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QSystemTrayIcon{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSystemTrayIcon) callVirtualBase_CustomEvent(event *QEvent) {

	QSystemTrayIcon_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSystemTrayIcon) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSystemTrayIcon_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_CustomEvent
func miqt_exec_callback_QSystemTrayIcon_CustomEvent(self QSystemTrayIcon, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSystemTrayIcon{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QSystemTrayIcon) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QSystemTrayIcon_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSystemTrayIcon) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSystemTrayIcon_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_ConnectNotify
func miqt_exec_callback_QSystemTrayIcon_ConnectNotify(self QSystemTrayIcon, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSystemTrayIcon{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QSystemTrayIcon) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QSystemTrayIcon_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSystemTrayIcon) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSystemTrayIcon_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSystemTrayIcon_DisconnectNotify
func miqt_exec_callback_QSystemTrayIcon_DisconnectNotify(self QSystemTrayIcon, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSystemTrayIcon{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
