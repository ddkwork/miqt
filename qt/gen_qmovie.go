package qt

import (
	"unsafe"
)

type QMovie__MovieState int

const (
	QMovie__NotRunning QMovie__MovieState = 0
	QMovie__Paused     QMovie__MovieState = 1
	QMovie__Running    QMovie__MovieState = 2
)

type QMovie__CacheMode int

const (
	QMovie__CacheNone QMovie__CacheMode = 0
	QMovie__CacheAll  QMovie__CacheMode = 1
)

type QMovie struct {
	h          uintptr
	isSubclass bool
}

// NewQMovie constructs a new QMovie object.
func NewQMovie() *QMovie {

	ret := newQMovie(QMovie_new())
	ret.isSubclass = true
	return ret
}

// NewQMovie2 constructs a new QMovie object.
func NewQMovie2(device *QIODevice) *QMovie {

	ret := newQMovie(QMovie_new2(device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMovie3 constructs a new QMovie object.
func NewQMovie3(fileName string) *QMovie {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQMovie(QMovie_new3(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQMovie4 constructs a new QMovie object.
func NewQMovie4(parent *QObject) *QMovie {

	ret := newQMovie(QMovie_new4(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMovie5 constructs a new QMovie object.
func NewQMovie5(device *QIODevice, format []byte) *QMovie {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQMovie(QMovie_new5(device.cPointer(), format_alias))
	ret.isSubclass = true
	return ret
}

// NewQMovie6 constructs a new QMovie object.
func NewQMovie6(device *QIODevice, format []byte, parent *QObject) *QMovie {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQMovie(QMovie_new6(device.cPointer(), format_alias, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMovie7 constructs a new QMovie object.
func NewQMovie7(fileName string, format []byte) *QMovie {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQMovie(QMovie_new7(fileName_ms, format_alias))
	ret.isSubclass = true
	return ret
}

// NewQMovie8 constructs a new QMovie object.
func NewQMovie8(fileName string, format []byte, parent *QObject) *QMovie {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQMovie(QMovie_new8(fileName_ms, format_alias, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMovie) MetaObject() *QMetaObject {
	return newQMetaObject(QMovie_MetaObject(this.h))
}

func (this *QMovie) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMovie_Metacast(this.h, param1_Cstring))
}

func QMovie_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMovie_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMovie_SupportedFormats() [][]byte {
	var _ma struct_miqt_array = QMovie_SupportedFormats()
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QMovie) SetDevice(device *QIODevice) {
	QMovie_SetDevice(this.h, device.cPointer())
}

func (this *QMovie) Device() *QIODevice {
	return newQIODevice(QMovie_Device(this.h))
}

func (this *QMovie) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QMovie_SetFileName(this.h, fileName_ms)
}

func (this *QMovie) FileName() string {
	var _ms struct_miqt_string = QMovie_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMovie) SetFormat(format []byte) {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	QMovie_SetFormat(this.h, format_alias)
}

func (this *QMovie) Format() []byte {
	var _bytearray struct_miqt_string = QMovie_Format(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMovie) SetBackgroundColor(color *QColor) {
	QMovie_SetBackgroundColor(this.h, color.cPointer())
}

func (this *QMovie) BackgroundColor() *QColor {
	_goptr := newQColor(QMovie_BackgroundColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMovie) State() MovieState {
	xxxxxxxxx
}

func (this *QMovie) FrameRect() *QRect {
	_goptr := newQRect(QMovie_FrameRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMovie) CurrentImage() *QImage {
	_goptr := newQImage(QMovie_CurrentImage(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMovie) CurrentPixmap() *QPixmap {
	_goptr := newQPixmap(QMovie_CurrentPixmap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMovie) IsValid() bool {
	return (bool)(QMovie_IsValid(this.h))
}

func (this *QMovie) LastError() QImageReader__ImageReaderError {
	return (QImageReader__ImageReaderError)(QMovie_LastError(this.h))
}

func (this *QMovie) LastErrorString() string {
	var _ms struct_miqt_string = QMovie_LastErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMovie) JumpToFrame(frameNumber int) bool {
	return (bool)(QMovie_JumpToFrame(this.h, (int)(frameNumber)))
}

func (this *QMovie) LoopCount() int {
	return (int)(QMovie_LoopCount(this.h))
}

func (this *QMovie) FrameCount() int {
	return (int)(QMovie_FrameCount(this.h))
}

func (this *QMovie) NextFrameDelay() int {
	return (int)(QMovie_NextFrameDelay(this.h))
}

func (this *QMovie) CurrentFrameNumber() int {
	return (int)(QMovie_CurrentFrameNumber(this.h))
}

func (this *QMovie) Speed() int {
	return (int)(QMovie_Speed(this.h))
}

func (this *QMovie) ScaledSize() *QSize {
	_goptr := newQSize(QMovie_ScaledSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMovie) SetScaledSize(size *QSize) {
	QMovie_SetScaledSize(this.h, size.cPointer())
}

func (this *QMovie) CacheMode() CacheMode {
	xxxxxxxxx
}

func (this *QMovie) SetCacheMode(mode CacheMode) {
	QMovie_SetCacheMode(this.h, mode)
}

func (this *QMovie) Started() {
	QMovie_Started(this.h)
}
func (this *QMovie) OnStarted(slot func()) {
	QMovie_connect_Started(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_Started
func miqt_exec_callback_QMovie_Started(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMovie) Resized(size *QSize) {
	QMovie_Resized(this.h, size.cPointer())
}
func (this *QMovie) OnResized(slot func(size *QSize)) {
	QMovie_connect_Resized(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_Resized
func miqt_exec_callback_QMovie_Resized(cb intptr_t, size *QSize) {
	gofunc, ok := cgo.Handle(cb).Value().(func(size *QSize))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(size)

	gofunc(slotval1)
}

func (this *QMovie) Updated(rect *QRect) {
	QMovie_Updated(this.h, rect.cPointer())
}
func (this *QMovie) OnUpdated(slot func(rect *QRect)) {
	QMovie_connect_Updated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_Updated
func miqt_exec_callback_QMovie_Updated(cb intptr_t, rect *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(rect *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	gofunc(slotval1)
}

func (this *QMovie) StateChanged(state QMovie__MovieState) {
	QMovie_StateChanged(this.h, (int)(state))
}
func (this *QMovie) OnStateChanged(slot func(state QMovie__MovieState)) {
	QMovie_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_StateChanged
func miqt_exec_callback_QMovie_StateChanged(cb intptr_t, state int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state QMovie__MovieState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QMovie__MovieState)(state)

	gofunc(slotval1)
}

func (this *QMovie) Error(error QImageReader__ImageReaderError) {
	QMovie_Error(this.h, (int)(error))
}
func (this *QMovie) OnError(slot func(error QImageReader__ImageReaderError)) {
	QMovie_connect_Error(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_Error
func miqt_exec_callback_QMovie_Error(cb intptr_t, error int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QImageReader__ImageReaderError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QImageReader__ImageReaderError)(error)

	gofunc(slotval1)
}

func (this *QMovie) Finished() {
	QMovie_Finished(this.h)
}
func (this *QMovie) OnFinished(slot func()) {
	QMovie_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_Finished
func miqt_exec_callback_QMovie_Finished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMovie) FrameChanged(frameNumber int) {
	QMovie_FrameChanged(this.h, (int)(frameNumber))
}
func (this *QMovie) OnFrameChanged(slot func(frameNumber int)) {
	QMovie_connect_FrameChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_FrameChanged
func miqt_exec_callback_QMovie_FrameChanged(cb intptr_t, frameNumber int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(frameNumber int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(frameNumber)

	gofunc(slotval1)
}

func (this *QMovie) Start() {
	QMovie_Start(this.h)
}

func (this *QMovie) JumpToNextFrame() bool {
	return (bool)(QMovie_JumpToNextFrame(this.h))
}

func (this *QMovie) SetPaused(paused bool) {
	QMovie_SetPaused(this.h, (bool)(paused))
}

func (this *QMovie) Stop() {
	QMovie_Stop(this.h)
}

func (this *QMovie) SetSpeed(percentSpeed int) {
	QMovie_SetSpeed(this.h, (int)(percentSpeed))
}

func QMovie_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMovie_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMovie_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMovie_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMovie) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QMovie_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QMovie) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMovie_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_Event
func miqt_exec_callback_QMovie_Event(self QMovie, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QMovie{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QMovie) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QMovie_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QMovie) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMovie_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_EventFilter
func miqt_exec_callback_QMovie_EventFilter(self QMovie, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QMovie{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QMovie) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QMovie_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMovie) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMovie_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_TimerEvent
func miqt_exec_callback_QMovie_TimerEvent(self QMovie, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QMovie{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QMovie) callVirtualBase_ChildEvent(event *QChildEvent) {

	QMovie_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMovie) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMovie_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_ChildEvent
func miqt_exec_callback_QMovie_ChildEvent(self QMovie, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QMovie{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QMovie) callVirtualBase_CustomEvent(event *QEvent) {

	QMovie_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMovie) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMovie_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_CustomEvent
func miqt_exec_callback_QMovie_CustomEvent(self QMovie, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QMovie{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QMovie) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QMovie_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QMovie) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMovie_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_ConnectNotify
func miqt_exec_callback_QMovie_ConnectNotify(self QMovie, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QMovie{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QMovie) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QMovie_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QMovie) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMovie_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMovie_DisconnectNotify
func miqt_exec_callback_QMovie_DisconnectNotify(self QMovie, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QMovie{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
