package qt

import (
	"unsafe"
)

type QFileSelector struct {
	h          uintptr
	isSubclass bool
}

// NewQFileSelector constructs a new QFileSelector object.
func NewQFileSelector() *QFileSelector {

	ret := newQFileSelector(QFileSelector_new())
	ret.isSubclass = true
	return ret
}

// NewQFileSelector2 constructs a new QFileSelector object.
func NewQFileSelector2(parent *QObject) *QFileSelector {

	ret := newQFileSelector(QFileSelector_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFileSelector) MetaObject() *QMetaObject {
	return newQMetaObject(QFileSelector_MetaObject(this.h))
}

func (this *QFileSelector) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFileSelector_Metacast(this.h, param1_Cstring))
}

func QFileSelector_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFileSelector_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSelector) Select(filePath string) string {
	filePath_ms := struct_miqt_string{}
	filePath_ms.data = CString(filePath)
	filePath_ms.len = size_t(len(filePath))
	defer free(unsafe.Pointer(filePath_ms.data))
	var _ms struct_miqt_string = QFileSelector_Select(this.h, filePath_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSelector) SelectWithFilePath(filePath *QUrl) *QUrl {
	_goptr := newQUrl(QFileSelector_SelectWithFilePath(this.h, filePath.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSelector) ExtraSelectors() []string {
	var _ma struct_miqt_array = QFileSelector_ExtraSelectors(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QFileSelector) SetExtraSelectors(list []string) {
	list_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(list))))
	defer free(unsafe.Pointer(list_CArray))
	for i := range list {
		list_i_ms := struct_miqt_string{}
		list_i_ms.data = CString(list[i])
		list_i_ms.len = size_t(len(list[i]))
		defer free(unsafe.Pointer(list_i_ms.data))
		list_CArray[i] = list_i_ms
	}
	list_ma := struct_miqt_array{len: size_t(len(list)), data: unsafe.Pointer(list_CArray)}
	QFileSelector_SetExtraSelectors(this.h, list_ma)
}

func (this *QFileSelector) AllSelectors() []string {
	var _ma struct_miqt_array = QFileSelector_AllSelectors(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QFileSelector_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSelector_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileSelector_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSelector_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSelector) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QFileSelector_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QFileSelector) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_Event
func miqt_exec_callback_QFileSelector_Event(self QFileSelector, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QFileSelector{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFileSelector) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QFileSelector_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QFileSelector) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_EventFilter
func miqt_exec_callback_QFileSelector_EventFilter(self QFileSelector, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QFileSelector{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QFileSelector) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QFileSelector_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFileSelector) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_TimerEvent
func miqt_exec_callback_QFileSelector_TimerEvent(self QFileSelector, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QFileSelector{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QFileSelector) callVirtualBase_ChildEvent(event *QChildEvent) {

	QFileSelector_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFileSelector) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_ChildEvent
func miqt_exec_callback_QFileSelector_ChildEvent(self QFileSelector, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QFileSelector{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QFileSelector) callVirtualBase_CustomEvent(event *QEvent) {

	QFileSelector_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFileSelector) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_CustomEvent
func miqt_exec_callback_QFileSelector_CustomEvent(self QFileSelector, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QFileSelector{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QFileSelector) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QFileSelector_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QFileSelector) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_ConnectNotify
func miqt_exec_callback_QFileSelector_ConnectNotify(self QFileSelector, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QFileSelector{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QFileSelector) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QFileSelector_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QFileSelector) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_DisconnectNotify
func miqt_exec_callback_QFileSelector_DisconnectNotify(self QFileSelector, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QFileSelector{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
