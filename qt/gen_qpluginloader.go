package qt

import (
	"unsafe"
)

type QPluginLoader struct {
	h          uintptr
	isSubclass bool
}

// NewQPluginLoader constructs a new QPluginLoader object.
func NewQPluginLoader() *QPluginLoader {

	ret := newQPluginLoader(QPluginLoader_new())
	ret.isSubclass = true
	return ret
}

// NewQPluginLoader2 constructs a new QPluginLoader object.
func NewQPluginLoader2(fileName string) *QPluginLoader {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQPluginLoader(QPluginLoader_new2(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQPluginLoader3 constructs a new QPluginLoader object.
func NewQPluginLoader3(parent *QObject) *QPluginLoader {

	ret := newQPluginLoader(QPluginLoader_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPluginLoader4 constructs a new QPluginLoader object.
func NewQPluginLoader4(fileName string, parent *QObject) *QPluginLoader {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQPluginLoader(QPluginLoader_new4(fileName_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPluginLoader) MetaObject() *QMetaObject {
	return newQMetaObject(QPluginLoader_MetaObject(this.h))
}

func (this *QPluginLoader) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPluginLoader_Metacast(this.h, param1_Cstring))
}

func QPluginLoader_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPluginLoader_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) Instance() *QObject {
	return newQObject(QPluginLoader_Instance(this.h))
}

func (this *QPluginLoader) MetaData() *QJsonObject {
	_goptr := newQJsonObject(QPluginLoader_MetaData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPluginLoader_StaticInstances() []*QObject {
	var _ma struct_miqt_array = QPluginLoader_StaticInstances()
	_ret := make([]*QObject, int(_ma.len))
	_outCast := (*[0xffff]*QObject)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQObject(_outCast[i])
	}
	return _ret
}

func QPluginLoader_StaticPlugins() []QStaticPlugin {
	var _ma struct_miqt_array = QPluginLoader_StaticPlugins()
	_ret := make([]QStaticPlugin, int(_ma.len))
	_outCast := (*[0xffff]*QStaticPlugin)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQStaticPlugin(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPluginLoader) Load() bool {
	return (bool)(QPluginLoader_Load(this.h))
}

func (this *QPluginLoader) Unload() bool {
	return (bool)(QPluginLoader_Unload(this.h))
}

func (this *QPluginLoader) IsLoaded() bool {
	return (bool)(QPluginLoader_IsLoaded(this.h))
}

func (this *QPluginLoader) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QPluginLoader_SetFileName(this.h, fileName_ms)
}

func (this *QPluginLoader) FileName() string {
	var _ms struct_miqt_string = QPluginLoader_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) ErrorString() string {
	var _ms struct_miqt_string = QPluginLoader_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) SetLoadHints(loadHints LoadHint) {
	QPluginLoader_SetLoadHints(this.h, (int)(loadHints))
}

func (this *QPluginLoader) LoadHints() LoadHint {
	return (LoadHint)(QPluginLoader_LoadHints(this.h))
}

func QPluginLoader_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPluginLoader_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPluginLoader_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPluginLoader_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QPluginLoader_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QPluginLoader) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_Event
func miqt_exec_callback_QPluginLoader_Event(self QPluginLoader, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QPluginLoader{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QPluginLoader) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QPluginLoader_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QPluginLoader) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_EventFilter
func miqt_exec_callback_QPluginLoader_EventFilter(self QPluginLoader, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QPluginLoader{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QPluginLoader) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QPluginLoader_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QPluginLoader) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_TimerEvent
func miqt_exec_callback_QPluginLoader_TimerEvent(self QPluginLoader, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QPluginLoader{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QPluginLoader) callVirtualBase_ChildEvent(event *QChildEvent) {

	QPluginLoader_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QPluginLoader) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_ChildEvent
func miqt_exec_callback_QPluginLoader_ChildEvent(self QPluginLoader, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QPluginLoader{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QPluginLoader) callVirtualBase_CustomEvent(event *QEvent) {

	QPluginLoader_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QPluginLoader) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_CustomEvent
func miqt_exec_callback_QPluginLoader_CustomEvent(self QPluginLoader, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QPluginLoader{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QPluginLoader) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QPluginLoader_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QPluginLoader) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_ConnectNotify
func miqt_exec_callback_QPluginLoader_ConnectNotify(self QPluginLoader, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QPluginLoader{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QPluginLoader) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QPluginLoader_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QPluginLoader) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_DisconnectNotify
func miqt_exec_callback_QPluginLoader_DisconnectNotify(self QPluginLoader, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QPluginLoader{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
