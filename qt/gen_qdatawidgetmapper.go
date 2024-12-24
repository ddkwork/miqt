package qt

import (
	"unsafe"
)

type QDataWidgetMapper__SubmitPolicy int

const (
	QDataWidgetMapper__AutoSubmit   QDataWidgetMapper__SubmitPolicy = 0
	QDataWidgetMapper__ManualSubmit QDataWidgetMapper__SubmitPolicy = 1
)

type QDataWidgetMapper struct {
	h          uintptr
	isSubclass bool
}

// NewQDataWidgetMapper constructs a new QDataWidgetMapper object.
func NewQDataWidgetMapper() *QDataWidgetMapper {

	ret := newQDataWidgetMapper(QDataWidgetMapper_new())
	ret.isSubclass = true
	return ret
}

// NewQDataWidgetMapper2 constructs a new QDataWidgetMapper object.
func NewQDataWidgetMapper2(parent *QObject) *QDataWidgetMapper {

	ret := newQDataWidgetMapper(QDataWidgetMapper_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDataWidgetMapper) MetaObject() *QMetaObject {
	return newQMetaObject(QDataWidgetMapper_MetaObject(this.h))
}

func (this *QDataWidgetMapper) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDataWidgetMapper_Metacast(this.h, param1_Cstring))
}

func QDataWidgetMapper_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDataWidgetMapper_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDataWidgetMapper) SetModel(model *QAbstractItemModel) {
	QDataWidgetMapper_SetModel(this.h, model.cPointer())
}

func (this *QDataWidgetMapper) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QDataWidgetMapper_Model(this.h))
}

func (this *QDataWidgetMapper) SetItemDelegate(delegate *QAbstractItemDelegate) {
	QDataWidgetMapper_SetItemDelegate(this.h, delegate.cPointer())
}

func (this *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QDataWidgetMapper_ItemDelegate(this.h))
}

func (this *QDataWidgetMapper) SetRootIndex(index *QModelIndex) {
	QDataWidgetMapper_SetRootIndex(this.h, index.cPointer())
}

func (this *QDataWidgetMapper) RootIndex() *QModelIndex {
	_goptr := newQModelIndex(QDataWidgetMapper_RootIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDataWidgetMapper) SetOrientation(aOrientation Orientation) {
	QDataWidgetMapper_SetOrientation(this.h, (int)(aOrientation))
}

func (this *QDataWidgetMapper) Orientation() Orientation {
	return (Orientation)(QDataWidgetMapper_Orientation(this.h))
}

func (this *QDataWidgetMapper) SetSubmitPolicy(policy SubmitPolicy) {
	QDataWidgetMapper_SetSubmitPolicy(this.h, policy)
}

func (this *QDataWidgetMapper) SubmitPolicy() SubmitPolicy {
	xxxxxxxxx
}

func (this *QDataWidgetMapper) AddMapping(widget *QWidget, section int) {
	QDataWidgetMapper_AddMapping(this.h, widget.cPointer(), (int)(section))
}

func (this *QDataWidgetMapper) AddMapping2(widget *QWidget, section int, propertyName []byte) {
	propertyName_alias := struct_miqt_string{}
	propertyName_alias.data = (char)(unsafe.Pointer(&propertyName[0]))
	propertyName_alias.len = size_t(len(propertyName))
	QDataWidgetMapper_AddMapping2(this.h, widget.cPointer(), (int)(section), propertyName_alias)
}

func (this *QDataWidgetMapper) RemoveMapping(widget *QWidget) {
	QDataWidgetMapper_RemoveMapping(this.h, widget.cPointer())
}

func (this *QDataWidgetMapper) MappedSection(widget *QWidget) int {
	return (int)(QDataWidgetMapper_MappedSection(this.h, widget.cPointer()))
}

func (this *QDataWidgetMapper) MappedPropertyName(widget *QWidget) []byte {
	var _bytearray struct_miqt_string = QDataWidgetMapper_MappedPropertyName(this.h, widget.cPointer())
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget {
	return newQWidget(QDataWidgetMapper_MappedWidgetAt(this.h, (int)(section)))
}

func (this *QDataWidgetMapper) ClearMapping() {
	QDataWidgetMapper_ClearMapping(this.h)
}

func (this *QDataWidgetMapper) CurrentIndex() int {
	return (int)(QDataWidgetMapper_CurrentIndex(this.h))
}

func (this *QDataWidgetMapper) Revert() {
	QDataWidgetMapper_Revert(this.h)
}

func (this *QDataWidgetMapper) Submit() bool {
	return (bool)(QDataWidgetMapper_Submit(this.h))
}

func (this *QDataWidgetMapper) ToFirst() {
	QDataWidgetMapper_ToFirst(this.h)
}

func (this *QDataWidgetMapper) ToLast() {
	QDataWidgetMapper_ToLast(this.h)
}

func (this *QDataWidgetMapper) ToNext() {
	QDataWidgetMapper_ToNext(this.h)
}

func (this *QDataWidgetMapper) ToPrevious() {
	QDataWidgetMapper_ToPrevious(this.h)
}

func (this *QDataWidgetMapper) SetCurrentIndex(index int) {
	QDataWidgetMapper_SetCurrentIndex(this.h, (int)(index))
}

func (this *QDataWidgetMapper) SetCurrentModelIndex(index *QModelIndex) {
	QDataWidgetMapper_SetCurrentModelIndex(this.h, index.cPointer())
}

func (this *QDataWidgetMapper) CurrentIndexChanged(index int) {
	QDataWidgetMapper_CurrentIndexChanged(this.h, (int)(index))
}
func (this *QDataWidgetMapper) OnCurrentIndexChanged(slot func(index int)) {
	QDataWidgetMapper_connect_CurrentIndexChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_CurrentIndexChanged
func miqt_exec_callback_QDataWidgetMapper_CurrentIndexChanged(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func QDataWidgetMapper_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDataWidgetMapper_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDataWidgetMapper_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDataWidgetMapper_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDataWidgetMapper) callVirtualBase_SetCurrentIndex(index int) {

	QDataWidgetMapper_virtualbase_SetCurrentIndex(unsafe.Pointer(this.h), (int)(index))

}
func (this *QDataWidgetMapper) OnSetCurrentIndex(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_SetCurrentIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_SetCurrentIndex
func miqt_exec_callback_QDataWidgetMapper_SetCurrentIndex(self QDataWidgetMapper, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_SetCurrentIndex, slotval1)

}

func (this *QDataWidgetMapper) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QDataWidgetMapper_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QDataWidgetMapper) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_Event
func miqt_exec_callback_QDataWidgetMapper_Event(self QDataWidgetMapper, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDataWidgetMapper) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QDataWidgetMapper_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QDataWidgetMapper) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_EventFilter
func miqt_exec_callback_QDataWidgetMapper_EventFilter(self QDataWidgetMapper, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QDataWidgetMapper) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QDataWidgetMapper_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDataWidgetMapper) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_TimerEvent
func miqt_exec_callback_QDataWidgetMapper_TimerEvent(self QDataWidgetMapper, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QDataWidgetMapper) callVirtualBase_ChildEvent(event *QChildEvent) {

	QDataWidgetMapper_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDataWidgetMapper) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_ChildEvent
func miqt_exec_callback_QDataWidgetMapper_ChildEvent(self QDataWidgetMapper, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QDataWidgetMapper) callVirtualBase_CustomEvent(event *QEvent) {

	QDataWidgetMapper_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDataWidgetMapper) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_CustomEvent
func miqt_exec_callback_QDataWidgetMapper_CustomEvent(self QDataWidgetMapper, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QDataWidgetMapper) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QDataWidgetMapper_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QDataWidgetMapper) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_ConnectNotify
func miqt_exec_callback_QDataWidgetMapper_ConnectNotify(self QDataWidgetMapper, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QDataWidgetMapper) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QDataWidgetMapper_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QDataWidgetMapper) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDataWidgetMapper_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDataWidgetMapper_DisconnectNotify
func miqt_exec_callback_QDataWidgetMapper_DisconnectNotify(self QDataWidgetMapper, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QDataWidgetMapper{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
