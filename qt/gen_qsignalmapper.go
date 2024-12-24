package qt

import (
	"unsafe"
)

type QSignalMapper struct {
	h          uintptr
	isSubclass bool
}

// NewQSignalMapper constructs a new QSignalMapper object.
func NewQSignalMapper() *QSignalMapper {

	ret := newQSignalMapper(QSignalMapper_new())
	ret.isSubclass = true
	return ret
}

// NewQSignalMapper2 constructs a new QSignalMapper object.
func NewQSignalMapper2(parent *QObject) *QSignalMapper {

	ret := newQSignalMapper(QSignalMapper_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSignalMapper) MetaObject() *QMetaObject {
	return newQMetaObject(QSignalMapper_MetaObject(this.h))
}

func (this *QSignalMapper) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSignalMapper_Metacast(this.h, param1_Cstring))
}

func QSignalMapper_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSignalMapper_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSignalMapper) SetMapping(sender *QObject, id int) {
	QSignalMapper_SetMapping(this.h, sender.cPointer(), (int)(id))
}

func (this *QSignalMapper) SetMapping2(sender *QObject, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QSignalMapper_SetMapping2(this.h, sender.cPointer(), text_ms)
}

func (this *QSignalMapper) SetMapping3(sender *QObject, object *QObject) {
	QSignalMapper_SetMapping3(this.h, sender.cPointer(), object.cPointer())
}

func (this *QSignalMapper) RemoveMappings(sender *QObject) {
	QSignalMapper_RemoveMappings(this.h, sender.cPointer())
}

func (this *QSignalMapper) Mapping(id int) *QObject {
	return newQObject(QSignalMapper_Mapping(this.h, (int)(id)))
}

func (this *QSignalMapper) MappingWithText(text string) *QObject {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQObject(QSignalMapper_MappingWithText(this.h, text_ms))
}

func (this *QSignalMapper) MappingWithObject(object *QObject) *QObject {
	return newQObject(QSignalMapper_MappingWithObject(this.h, object.cPointer()))
}

func (this *QSignalMapper) MappedInt(param1 int) {
	QSignalMapper_MappedInt(this.h, (int)(param1))
}
func (this *QSignalMapper) OnMappedInt(slot func(param1 int)) {
	QSignalMapper_connect_MappedInt(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_MappedInt
func miqt_exec_callback_QSignalMapper_MappedInt(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QSignalMapper) MappedString(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QSignalMapper_MappedString(this.h, param1_ms)
}
func (this *QSignalMapper) OnMappedString(slot func(param1 string)) {
	QSignalMapper_connect_MappedString(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_MappedString
func miqt_exec_callback_QSignalMapper_MappedString(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func (this *QSignalMapper) MappedObject(param1 *QObject) {
	QSignalMapper_MappedObject(this.h, param1.cPointer())
}
func (this *QSignalMapper) OnMappedObject(slot func(param1 *QObject)) {
	QSignalMapper_connect_MappedObject(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_MappedObject
func miqt_exec_callback_QSignalMapper_MappedObject(cb intptr_t, param1 *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	gofunc(slotval1)
}

func (this *QSignalMapper) Map() {
	QSignalMapper_Map(this.h)
}

func (this *QSignalMapper) MapWithSender(sender *QObject) {
	QSignalMapper_MapWithSender(this.h, sender.cPointer())
}

func QSignalMapper_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSignalMapper_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSignalMapper_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSignalMapper_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSignalMapper) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QSignalMapper_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QSignalMapper) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_Event
func miqt_exec_callback_QSignalMapper_Event(self QSignalMapper, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QSignalMapper{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSignalMapper) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QSignalMapper_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QSignalMapper) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_EventFilter
func miqt_exec_callback_QSignalMapper_EventFilter(self QSignalMapper, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QSignalMapper{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSignalMapper) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QSignalMapper_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSignalMapper) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_TimerEvent
func miqt_exec_callback_QSignalMapper_TimerEvent(self QSignalMapper, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QSignalMapper{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSignalMapper) callVirtualBase_ChildEvent(event *QChildEvent) {

	QSignalMapper_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSignalMapper) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_ChildEvent
func miqt_exec_callback_QSignalMapper_ChildEvent(self QSignalMapper, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QSignalMapper{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSignalMapper) callVirtualBase_CustomEvent(event *QEvent) {

	QSignalMapper_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSignalMapper) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_CustomEvent
func miqt_exec_callback_QSignalMapper_CustomEvent(self QSignalMapper, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSignalMapper{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QSignalMapper) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QSignalMapper_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSignalMapper) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_ConnectNotify
func miqt_exec_callback_QSignalMapper_ConnectNotify(self QSignalMapper, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSignalMapper{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QSignalMapper) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QSignalMapper_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSignalMapper) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_DisconnectNotify
func miqt_exec_callback_QSignalMapper_DisconnectNotify(self QSignalMapper, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSignalMapper{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
