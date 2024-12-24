package qt

import (
	"unsafe"
)

type QStylePlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQStylePlugin constructs a new QStylePlugin object.
func NewQStylePlugin() *QStylePlugin {

	ret := newQStylePlugin(QStylePlugin_new())
	ret.isSubclass = true
	return ret
}

// NewQStylePlugin2 constructs a new QStylePlugin object.
func NewQStylePlugin2(parent *QObject) *QStylePlugin {

	ret := newQStylePlugin(QStylePlugin_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStylePlugin) MetaObject() *QMetaObject {
	return newQMetaObject(QStylePlugin_MetaObject(this.h))
}

func (this *QStylePlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStylePlugin_Metacast(this.h, param1_Cstring))
}

func QStylePlugin_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStylePlugin_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStylePlugin) Create(key string) *QStyle {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	return newQStyle(QStylePlugin_Create(this.h, key_ms))
}

func QStylePlugin_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStylePlugin_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStylePlugin_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStylePlugin_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QStylePlugin) OnCreate(slot func(key string) *QStyle) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_Create(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_Create
func miqt_exec_callback_QStylePlugin_Create(self QStylePlugin, cb intptr_t, key struct_miqt_string) *QStyle {
	gofunc, ok := cgo.Handle(cb).Value().(func(key string) *QStyle)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var key_ms struct_miqt_string = key
	key_ret := GoStringN(key_ms.data, int(int64(key_ms.len)))
	free(unsafe.Pointer(key_ms.data))
	slotval1 := key_ret

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}

func (this *QStylePlugin) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QStylePlugin_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QStylePlugin) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_Event
func miqt_exec_callback_QStylePlugin_Event(self QStylePlugin, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QStylePlugin{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QStylePlugin) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QStylePlugin_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QStylePlugin) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_EventFilter
func miqt_exec_callback_QStylePlugin_EventFilter(self QStylePlugin, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QStylePlugin{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QStylePlugin) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QStylePlugin_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStylePlugin) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_TimerEvent
func miqt_exec_callback_QStylePlugin_TimerEvent(self QStylePlugin, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QStylePlugin{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QStylePlugin) callVirtualBase_ChildEvent(event *QChildEvent) {

	QStylePlugin_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStylePlugin) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_ChildEvent
func miqt_exec_callback_QStylePlugin_ChildEvent(self QStylePlugin, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QStylePlugin{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QStylePlugin) callVirtualBase_CustomEvent(event *QEvent) {

	QStylePlugin_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStylePlugin) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_CustomEvent
func miqt_exec_callback_QStylePlugin_CustomEvent(self QStylePlugin, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QStylePlugin{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QStylePlugin) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QStylePlugin_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QStylePlugin) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_ConnectNotify
func miqt_exec_callback_QStylePlugin_ConnectNotify(self QStylePlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QStylePlugin{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QStylePlugin) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QStylePlugin_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QStylePlugin) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_DisconnectNotify
func miqt_exec_callback_QStylePlugin_DisconnectNotify(self QStylePlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QStylePlugin{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
