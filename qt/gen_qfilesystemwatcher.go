package qt

import (
	"unsafe"
)

type QFileSystemWatcher struct {
	h          uintptr
	isSubclass bool
}

// NewQFileSystemWatcher constructs a new QFileSystemWatcher object.
func NewQFileSystemWatcher() *QFileSystemWatcher {

	ret := newQFileSystemWatcher(QFileSystemWatcher_new())
	ret.isSubclass = true
	return ret
}

// NewQFileSystemWatcher2 constructs a new QFileSystemWatcher object.
func NewQFileSystemWatcher2(paths []string) *QFileSystemWatcher {
	paths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(paths))))
	defer free(unsafe.Pointer(paths_CArray))
	for i := range paths {
		paths_i_ms := struct_miqt_string{}
		paths_i_ms.data = CString(paths[i])
		paths_i_ms.len = size_t(len(paths[i]))
		defer free(unsafe.Pointer(paths_i_ms.data))
		paths_CArray[i] = paths_i_ms
	}
	paths_ma := struct_miqt_array{len: size_t(len(paths)), data: unsafe.Pointer(paths_CArray)}

	ret := newQFileSystemWatcher(QFileSystemWatcher_new2(paths_ma))
	ret.isSubclass = true
	return ret
}

// NewQFileSystemWatcher3 constructs a new QFileSystemWatcher object.
func NewQFileSystemWatcher3(parent *QObject) *QFileSystemWatcher {

	ret := newQFileSystemWatcher(QFileSystemWatcher_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFileSystemWatcher4 constructs a new QFileSystemWatcher object.
func NewQFileSystemWatcher4(paths []string, parent *QObject) *QFileSystemWatcher {
	paths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(paths))))
	defer free(unsafe.Pointer(paths_CArray))
	for i := range paths {
		paths_i_ms := struct_miqt_string{}
		paths_i_ms.data = CString(paths[i])
		paths_i_ms.len = size_t(len(paths[i]))
		defer free(unsafe.Pointer(paths_i_ms.data))
		paths_CArray[i] = paths_i_ms
	}
	paths_ma := struct_miqt_array{len: size_t(len(paths)), data: unsafe.Pointer(paths_CArray)}

	ret := newQFileSystemWatcher(QFileSystemWatcher_new4(paths_ma, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFileSystemWatcher) MetaObject() *QMetaObject {
	return newQMetaObject(QFileSystemWatcher_MetaObject(this.h))
}

func (this *QFileSystemWatcher) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFileSystemWatcher_Metacast(this.h, param1_Cstring))
}

func QFileSystemWatcher_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFileSystemWatcher_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemWatcher) AddPath(file string) bool {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	return (bool)(QFileSystemWatcher_AddPath(this.h, file_ms))
}

func (this *QFileSystemWatcher) AddPaths(files []string) []string {
	files_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(files))))
	defer free(unsafe.Pointer(files_CArray))
	for i := range files {
		files_i_ms := struct_miqt_string{}
		files_i_ms.data = CString(files[i])
		files_i_ms.len = size_t(len(files[i]))
		defer free(unsafe.Pointer(files_i_ms.data))
		files_CArray[i] = files_i_ms
	}
	files_ma := struct_miqt_array{len: size_t(len(files)), data: unsafe.Pointer(files_CArray)}
	var _ma struct_miqt_array = QFileSystemWatcher_AddPaths(this.h, files_ma)
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

func (this *QFileSystemWatcher) RemovePath(file string) bool {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	return (bool)(QFileSystemWatcher_RemovePath(this.h, file_ms))
}

func (this *QFileSystemWatcher) RemovePaths(files []string) []string {
	files_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(files))))
	defer free(unsafe.Pointer(files_CArray))
	for i := range files {
		files_i_ms := struct_miqt_string{}
		files_i_ms.data = CString(files[i])
		files_i_ms.len = size_t(len(files[i]))
		defer free(unsafe.Pointer(files_i_ms.data))
		files_CArray[i] = files_i_ms
	}
	files_ma := struct_miqt_array{len: size_t(len(files)), data: unsafe.Pointer(files_CArray)}
	var _ma struct_miqt_array = QFileSystemWatcher_RemovePaths(this.h, files_ma)
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

func (this *QFileSystemWatcher) Files() []string {
	var _ma struct_miqt_array = QFileSystemWatcher_Files(this.h)
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

func (this *QFileSystemWatcher) Directories() []string {
	var _ma struct_miqt_array = QFileSystemWatcher_Directories(this.h)
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

func QFileSystemWatcher_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSystemWatcher_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileSystemWatcher_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSystemWatcher_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemWatcher) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QFileSystemWatcher_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QFileSystemWatcher) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemWatcher_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemWatcher_Event
func miqt_exec_callback_QFileSystemWatcher_Event(self QFileSystemWatcher, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QFileSystemWatcher{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFileSystemWatcher) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QFileSystemWatcher_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QFileSystemWatcher) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemWatcher_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemWatcher_EventFilter
func miqt_exec_callback_QFileSystemWatcher_EventFilter(self QFileSystemWatcher, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QFileSystemWatcher{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QFileSystemWatcher) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QFileSystemWatcher_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFileSystemWatcher) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemWatcher_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemWatcher_TimerEvent
func miqt_exec_callback_QFileSystemWatcher_TimerEvent(self QFileSystemWatcher, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QFileSystemWatcher{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QFileSystemWatcher) callVirtualBase_ChildEvent(event *QChildEvent) {

	QFileSystemWatcher_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFileSystemWatcher) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemWatcher_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemWatcher_ChildEvent
func miqt_exec_callback_QFileSystemWatcher_ChildEvent(self QFileSystemWatcher, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QFileSystemWatcher{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QFileSystemWatcher) callVirtualBase_CustomEvent(event *QEvent) {

	QFileSystemWatcher_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFileSystemWatcher) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemWatcher_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemWatcher_CustomEvent
func miqt_exec_callback_QFileSystemWatcher_CustomEvent(self QFileSystemWatcher, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QFileSystemWatcher{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QFileSystemWatcher) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QFileSystemWatcher_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QFileSystemWatcher) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemWatcher_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemWatcher_ConnectNotify
func miqt_exec_callback_QFileSystemWatcher_ConnectNotify(self QFileSystemWatcher, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QFileSystemWatcher{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QFileSystemWatcher) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QFileSystemWatcher_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QFileSystemWatcher) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemWatcher_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemWatcher_DisconnectNotify
func miqt_exec_callback_QFileSystemWatcher_DisconnectNotify(self QFileSystemWatcher, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QFileSystemWatcher{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
