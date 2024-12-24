package qt

import (
	"unsafe"
)

type QLibrary__LoadHint int

const (
	QLibrary__ResolveAllSymbolsHint     QLibrary__LoadHint = 1
	QLibrary__ExportExternalSymbolsHint QLibrary__LoadHint = 2
	QLibrary__LoadArchiveMemberHint     QLibrary__LoadHint = 4
	QLibrary__PreventUnloadHint         QLibrary__LoadHint = 8
	QLibrary__DeepBindHint              QLibrary__LoadHint = 16
)

type QLibrary struct {
	h          uintptr
	isSubclass bool
}

// NewQLibrary constructs a new QLibrary object.
func NewQLibrary() *QLibrary {

	ret := newQLibrary(QLibrary_new())
	ret.isSubclass = true
	return ret
}

// NewQLibrary2 constructs a new QLibrary object.
func NewQLibrary2(fileName string) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQLibrary(QLibrary_new2(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQLibrary3 constructs a new QLibrary object.
func NewQLibrary3(fileName string, verNum int) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQLibrary(QLibrary_new3(fileName_ms, (int)(verNum)))
	ret.isSubclass = true
	return ret
}

// NewQLibrary4 constructs a new QLibrary object.
func NewQLibrary4(fileName string, version string) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	version_ms := struct_miqt_string{}
	version_ms.data = CString(version)
	version_ms.len = size_t(len(version))
	defer free(unsafe.Pointer(version_ms.data))

	ret := newQLibrary(QLibrary_new4(fileName_ms, version_ms))
	ret.isSubclass = true
	return ret
}

// NewQLibrary5 constructs a new QLibrary object.
func NewQLibrary5(parent *QObject) *QLibrary {

	ret := newQLibrary(QLibrary_new5(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLibrary6 constructs a new QLibrary object.
func NewQLibrary6(fileName string, parent *QObject) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQLibrary(QLibrary_new6(fileName_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLibrary7 constructs a new QLibrary object.
func NewQLibrary7(fileName string, verNum int, parent *QObject) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQLibrary(QLibrary_new7(fileName_ms, (int)(verNum), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLibrary8 constructs a new QLibrary object.
func NewQLibrary8(fileName string, version string, parent *QObject) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	version_ms := struct_miqt_string{}
	version_ms.data = CString(version)
	version_ms.len = size_t(len(version))
	defer free(unsafe.Pointer(version_ms.data))

	ret := newQLibrary(QLibrary_new8(fileName_ms, version_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QLibrary) MetaObject() *QMetaObject {
	return newQMetaObject(QLibrary_MetaObject(this.h))
}

func (this *QLibrary) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLibrary_Metacast(this.h, param1_Cstring))
}

func QLibrary_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLibrary_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) Load() bool {
	return (bool)(QLibrary_Load(this.h))
}

func (this *QLibrary) Unload() bool {
	return (bool)(QLibrary_Unload(this.h))
}

func (this *QLibrary) IsLoaded() bool {
	return (bool)(QLibrary_IsLoaded(this.h))
}

func QLibrary_IsLibrary(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QLibrary_IsLibrary(fileName_ms))
}

func (this *QLibrary) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QLibrary_SetFileName(this.h, fileName_ms)
}

func (this *QLibrary) FileName() string {
	var _ms struct_miqt_string = QLibrary_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) SetFileNameAndVersion(fileName string, verNum int) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QLibrary_SetFileNameAndVersion(this.h, fileName_ms, (int)(verNum))
}

func (this *QLibrary) SetFileNameAndVersion2(fileName string, version string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	version_ms := struct_miqt_string{}
	version_ms.data = CString(version)
	version_ms.len = size_t(len(version))
	defer free(unsafe.Pointer(version_ms.data))
	QLibrary_SetFileNameAndVersion2(this.h, fileName_ms, version_ms)
}

func (this *QLibrary) ErrorString() string {
	var _ms struct_miqt_string = QLibrary_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) SetLoadHints(hints LoadHints) {
	QLibrary_SetLoadHints(this.h, hints)
}

func (this *QLibrary) LoadHints() LoadHints {
	xxxxxxxxx
}

func QLibrary_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLibrary_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLibrary_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLibrary_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QLibrary_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QLibrary) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_Event
func miqt_exec_callback_QLibrary_Event(self QLibrary, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QLibrary{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLibrary) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QLibrary_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QLibrary) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_EventFilter
func miqt_exec_callback_QLibrary_EventFilter(self QLibrary, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QLibrary{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QLibrary) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QLibrary_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLibrary) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_TimerEvent
func miqt_exec_callback_QLibrary_TimerEvent(self QLibrary, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QLibrary{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QLibrary) callVirtualBase_ChildEvent(event *QChildEvent) {

	QLibrary_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLibrary) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_ChildEvent
func miqt_exec_callback_QLibrary_ChildEvent(self QLibrary, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QLibrary{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QLibrary) callVirtualBase_CustomEvent(event *QEvent) {

	QLibrary_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLibrary) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_CustomEvent
func miqt_exec_callback_QLibrary_CustomEvent(self QLibrary, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QLibrary{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QLibrary) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QLibrary_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QLibrary) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_ConnectNotify
func miqt_exec_callback_QLibrary_ConnectNotify(self QLibrary, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QLibrary{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QLibrary) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QLibrary_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QLibrary) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_DisconnectNotify
func miqt_exec_callback_QLibrary_DisconnectNotify(self QLibrary, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QLibrary{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
