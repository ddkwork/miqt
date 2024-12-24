package qt

import (
	"unsafe"
)

type QGenericPlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQGenericPlugin constructs a new QGenericPlugin object.
func NewQGenericPlugin() *QGenericPlugin {

	ret := newQGenericPlugin(QGenericPlugin_new())
	ret.isSubclass = true
	return ret
}

// NewQGenericPlugin2 constructs a new QGenericPlugin object.
func NewQGenericPlugin2(parent *QObject) *QGenericPlugin {

	ret := newQGenericPlugin(QGenericPlugin_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGenericPlugin) MetaObject() *QMetaObject {
	return newQMetaObject(QGenericPlugin_MetaObject(this.h))
}

func (this *QGenericPlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGenericPlugin_Metacast(this.h, param1_Cstring))
}

func QGenericPlugin_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGenericPlugin_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGenericPlugin) Create(name string, spec string) *QObject {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	spec_ms := struct_miqt_string{}
	spec_ms.data = CString(spec)
	spec_ms.len = size_t(len(spec))
	defer free(unsafe.Pointer(spec_ms.data))
	return newQObject(QGenericPlugin_Create(this.h, name_ms, spec_ms))
}

func QGenericPlugin_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGenericPlugin_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGenericPlugin_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGenericPlugin_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QGenericPlugin) OnCreate(slot func(name string, spec string) *QObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_Create(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_Create
func miqt_exec_callback_QGenericPlugin_Create(self QGenericPlugin, cb intptr_t, name struct_miqt_string, spec struct_miqt_string) *QObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(name string, spec string) *QObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var name_ms struct_miqt_string = name
	name_ret := GoStringN(name_ms.data, int(int64(name_ms.len)))
	free(unsafe.Pointer(name_ms.data))
	slotval1 := name_ret
	var spec_ms struct_miqt_string = spec
	spec_ret := GoStringN(spec_ms.data, int(int64(spec_ms.len)))
	free(unsafe.Pointer(spec_ms.data))
	slotval2 := spec_ret

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGenericPlugin) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGenericPlugin_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGenericPlugin) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_Event
func miqt_exec_callback_QGenericPlugin_Event(self QGenericPlugin, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGenericPlugin{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGenericPlugin) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QGenericPlugin_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QGenericPlugin) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_EventFilter
func miqt_exec_callback_QGenericPlugin_EventFilter(self QGenericPlugin, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QGenericPlugin{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QGenericPlugin) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QGenericPlugin_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGenericPlugin) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_TimerEvent
func miqt_exec_callback_QGenericPlugin_TimerEvent(self QGenericPlugin, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QGenericPlugin{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QGenericPlugin) callVirtualBase_ChildEvent(event *QChildEvent) {

	QGenericPlugin_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGenericPlugin) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_ChildEvent
func miqt_exec_callback_QGenericPlugin_ChildEvent(self QGenericPlugin, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QGenericPlugin{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QGenericPlugin) callVirtualBase_CustomEvent(event *QEvent) {

	QGenericPlugin_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGenericPlugin) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_CustomEvent
func miqt_exec_callback_QGenericPlugin_CustomEvent(self QGenericPlugin, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGenericPlugin{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QGenericPlugin) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QGenericPlugin_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGenericPlugin) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_ConnectNotify
func miqt_exec_callback_QGenericPlugin_ConnectNotify(self QGenericPlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGenericPlugin{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QGenericPlugin) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QGenericPlugin_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGenericPlugin) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_DisconnectNotify
func miqt_exec_callback_QGenericPlugin_DisconnectNotify(self QGenericPlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGenericPlugin{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
