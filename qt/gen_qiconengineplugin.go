package qt

import (
	"unsafe"
)

type QIconEnginePlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQIconEnginePlugin constructs a new QIconEnginePlugin object.
func NewQIconEnginePlugin() *QIconEnginePlugin {

	ret := newQIconEnginePlugin(QIconEnginePlugin_new())
	ret.isSubclass = true
	return ret
}

// NewQIconEnginePlugin2 constructs a new QIconEnginePlugin object.
func NewQIconEnginePlugin2(parent *QObject) *QIconEnginePlugin {

	ret := newQIconEnginePlugin(QIconEnginePlugin_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QIconEnginePlugin) MetaObject() *QMetaObject {
	return newQMetaObject(QIconEnginePlugin_MetaObject(this.h))
}

func (this *QIconEnginePlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QIconEnginePlugin_Metacast(this.h, param1_Cstring))
}

func QIconEnginePlugin_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QIconEnginePlugin_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEnginePlugin) Create(filename string) *QIconEngine {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return newQIconEngine(QIconEnginePlugin_Create(this.h, filename_ms))
}

func QIconEnginePlugin_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIconEnginePlugin_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIconEnginePlugin_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIconEnginePlugin_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QIconEnginePlugin) OnCreate(slot func(filename string) *QIconEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_Create(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_Create
func miqt_exec_callback_QIconEnginePlugin_Create(self QIconEnginePlugin, cb intptr_t, filename struct_miqt_string) *QIconEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(filename string) *QIconEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var filename_ms struct_miqt_string = filename
	filename_ret := GoStringN(filename_ms.data, int(int64(filename_ms.len)))
	free(unsafe.Pointer(filename_ms.data))
	slotval1 := filename_ret

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}

func (this *QIconEnginePlugin) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QIconEnginePlugin_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QIconEnginePlugin) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_Event
func miqt_exec_callback_QIconEnginePlugin_Event(self QIconEnginePlugin, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QIconEnginePlugin) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QIconEnginePlugin_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QIconEnginePlugin) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_EventFilter
func miqt_exec_callback_QIconEnginePlugin_EventFilter(self QIconEnginePlugin, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QIconEnginePlugin) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QIconEnginePlugin_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QIconEnginePlugin) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_TimerEvent
func miqt_exec_callback_QIconEnginePlugin_TimerEvent(self QIconEnginePlugin, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QIconEnginePlugin) callVirtualBase_ChildEvent(event *QChildEvent) {

	QIconEnginePlugin_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QIconEnginePlugin) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_ChildEvent
func miqt_exec_callback_QIconEnginePlugin_ChildEvent(self QIconEnginePlugin, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QIconEnginePlugin) callVirtualBase_CustomEvent(event *QEvent) {

	QIconEnginePlugin_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QIconEnginePlugin) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_CustomEvent
func miqt_exec_callback_QIconEnginePlugin_CustomEvent(self QIconEnginePlugin, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QIconEnginePlugin) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QIconEnginePlugin_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QIconEnginePlugin) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_ConnectNotify
func miqt_exec_callback_QIconEnginePlugin_ConnectNotify(self QIconEnginePlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QIconEnginePlugin) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QIconEnginePlugin_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QIconEnginePlugin) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_DisconnectNotify
func miqt_exec_callback_QIconEnginePlugin_DisconnectNotify(self QIconEnginePlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
