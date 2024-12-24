package webchannel

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QWebChannel struct {
	h          uintptr
	isSubclass bool
}

// NewQWebChannel constructs a new QWebChannel object.
func NewQWebChannel() *QWebChannel {

	ret := newQWebChannel(QWebChannel_new())
	ret.isSubclass = true
	return ret
}

// NewQWebChannel2 constructs a new QWebChannel object.
func NewQWebChannel2(parent *qt.QObject) *QWebChannel {

	ret := newQWebChannel(QWebChannel_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QWebChannel) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWebChannel_MetaObject(this.h)))
}

func (this *QWebChannel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWebChannel_Metacast(this.h, param1_Cstring))
}

func QWebChannel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWebChannel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebChannel) RegisterObjects(objects map[string]*qt.QObject) {
	objects_Keys_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(objects))))
	defer free(unsafe.Pointer(objects_Keys_CArray))
	objects_Values_CArray := (*[0xffff]*QObject)(malloc(size_t(8 * len(objects))))
	defer free(unsafe.Pointer(objects_Values_CArray))
	objects_ctr := 0
	for objects_k, objects_v := range objects {
		objects_k_ms := struct_miqt_string{}
		objects_k_ms.data = CString(objects_k)
		objects_k_ms.len = size_t(len(objects_k))
		defer free(unsafe.Pointer(objects_k_ms.data))
		objects_Keys_CArray[objects_ctr] = objects_k_ms
		objects_Values_CArray[objects_ctr] = (*QObject)(objects_v.UnsafePointer())
		objects_ctr++
	}
	objects_mm := struct_miqt_map{
		len:    size_t(len(objects)),
		keys:   unsafe.Pointer(objects_Keys_CArray),
		values: unsafe.Pointer(objects_Values_CArray),
	}
	QWebChannel_RegisterObjects(this.h, objects_mm)
}

func (this *QWebChannel) RegisteredObjects() map[string]*qt.QObject {
	var _mm struct_miqt_map = QWebChannel_RegisteredObjects(this.h)
	_ret := make(map[string]*qt.QObject, int(_mm.len))
	_Keys := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QObject)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		var _hashkey_ms struct_miqt_string = _Keys[i]
		_hashkey_ret := GoStringN(_hashkey_ms.data, int(int64(_hashkey_ms.len)))
		free(unsafe.Pointer(_hashkey_ms.data))
		_entry_Key := _hashkey_ret
		_entry_Value := qt.UnsafeNewQObject(unsafe.Pointer(_Values[i]))

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QWebChannel) RegisterObject(id string, object *qt.QObject) {
	id_ms := struct_miqt_string{}
	id_ms.data = CString(id)
	id_ms.len = size_t(len(id))
	defer free(unsafe.Pointer(id_ms.data))
	QWebChannel_RegisterObject(this.h, id_ms, (*QObject)(object.UnsafePointer()))
}

func (this *QWebChannel) DeregisterObject(object *qt.QObject) {
	QWebChannel_DeregisterObject(this.h, (*QObject)(object.UnsafePointer()))
}

func (this *QWebChannel) BlockUpdates() bool {
	return (bool)(QWebChannel_BlockUpdates(this.h))
}

func (this *QWebChannel) SetBlockUpdates(block bool) {
	QWebChannel_SetBlockUpdates(this.h, (bool)(block))
}

func (this *QWebChannel) PropertyUpdateInterval() int {
	return (int)(QWebChannel_PropertyUpdateInterval(this.h))
}

func (this *QWebChannel) SetPropertyUpdateInterval(ms int) {
	QWebChannel_SetPropertyUpdateInterval(this.h, (int)(ms))
}

func (this *QWebChannel) BlockUpdatesChanged(block bool) {
	QWebChannel_BlockUpdatesChanged(this.h, (bool)(block))
}
func (this *QWebChannel) OnBlockUpdatesChanged(slot func(block bool)) {
	QWebChannel_connect_BlockUpdatesChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_BlockUpdatesChanged
func miqt_exec_callback_QWebChannel_BlockUpdatesChanged(cb intptr_t, block bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(block bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(block)

	gofunc(slotval1)
}

func (this *QWebChannel) ConnectTo(transport *QWebChannelAbstractTransport) {
	QWebChannel_ConnectTo(this.h, transport.cPointer())
}

func (this *QWebChannel) DisconnectFrom(transport *QWebChannelAbstractTransport) {
	QWebChannel_DisconnectFrom(this.h, transport.cPointer())
}

func QWebChannel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWebChannel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebChannel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWebChannel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebChannel) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QWebChannel_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QWebChannel) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannel_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_Event
func miqt_exec_callback_QWebChannel_Event(self QWebChannel, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QWebChannel{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QWebChannel) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QWebChannel_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QWebChannel) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannel_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_EventFilter
func miqt_exec_callback_QWebChannel_EventFilter(self QWebChannel, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QWebChannel{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QWebChannel) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QWebChannel_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QWebChannel) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannel_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_TimerEvent
func miqt_exec_callback_QWebChannel_TimerEvent(self QWebChannel, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QWebChannel{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QWebChannel) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QWebChannel_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QWebChannel) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannel_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_ChildEvent
func miqt_exec_callback_QWebChannel_ChildEvent(self QWebChannel, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QWebChannel{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QWebChannel) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QWebChannel_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QWebChannel) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannel_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_CustomEvent
func miqt_exec_callback_QWebChannel_CustomEvent(self QWebChannel, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QWebChannel{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QWebChannel) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QWebChannel_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebChannel) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannel_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_ConnectNotify
func miqt_exec_callback_QWebChannel_ConnectNotify(self QWebChannel, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebChannel{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QWebChannel) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QWebChannel_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebChannel) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannel_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_DisconnectNotify
func miqt_exec_callback_QWebChannel_DisconnectNotify(self QWebChannel, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebChannel{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
