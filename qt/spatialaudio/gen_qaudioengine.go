package spatialaudio

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QAudioEngine__OutputMode int

const (
	QAudioEngine__Surround  QAudioEngine__OutputMode = 0
	QAudioEngine__Stereo    QAudioEngine__OutputMode = 1
	QAudioEngine__Headphone QAudioEngine__OutputMode = 2
)

type QAudioEngine struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioEngine constructs a new QAudioEngine object.
func NewQAudioEngine() *QAudioEngine {

	ret := newQAudioEngine(QAudioEngine_new())
	ret.isSubclass = true
	return ret
}

// NewQAudioEngine2 constructs a new QAudioEngine object.
func NewQAudioEngine2(parent *qt.QObject) *QAudioEngine {

	ret := newQAudioEngine(QAudioEngine_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQAudioEngine3 constructs a new QAudioEngine object.
func NewQAudioEngine3(sampleRate int) *QAudioEngine {

	ret := newQAudioEngine(QAudioEngine_new3((int)(sampleRate)))
	ret.isSubclass = true
	return ret
}

// NewQAudioEngine4 constructs a new QAudioEngine object.
func NewQAudioEngine4(sampleRate int, parent *qt.QObject) *QAudioEngine {

	ret := newQAudioEngine(QAudioEngine_new4((int)(sampleRate), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAudioEngine) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioEngine_MetaObject(this.h)))
}

func (this *QAudioEngine) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioEngine_Metacast(this.h, param1_Cstring))
}

func QAudioEngine_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioEngine_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioEngine) SetOutputMode(mode OutputMode) {
	QAudioEngine_SetOutputMode(this.h, mode)
}

func (this *QAudioEngine) OutputMode() OutputMode {
	xxxxxxxxx
}

func (this *QAudioEngine) SampleRate() int {
	return (int)(QAudioEngine_SampleRate(this.h))
}

func (this *QAudioEngine) SetOutputDevice(device *QAudioDevice) {
	QAudioEngine_SetOutputDevice(this.h, device)
}

func (this *QAudioEngine) OutputDevice() QAudioDevice {
	xxxxxxxxx
}

func (this *QAudioEngine) SetMasterVolume(volume float32) {
	QAudioEngine_SetMasterVolume(this.h, (float)(volume))
}

func (this *QAudioEngine) MasterVolume() float32 {
	return (float32)(QAudioEngine_MasterVolume(this.h))
}

func (this *QAudioEngine) SetPaused(paused bool) {
	QAudioEngine_SetPaused(this.h, (bool)(paused))
}

func (this *QAudioEngine) Paused() bool {
	return (bool)(QAudioEngine_Paused(this.h))
}

func (this *QAudioEngine) SetRoomEffectsEnabled(enabled bool) {
	QAudioEngine_SetRoomEffectsEnabled(this.h, (bool)(enabled))
}

func (this *QAudioEngine) RoomEffectsEnabled() bool {
	return (bool)(QAudioEngine_RoomEffectsEnabled(this.h))
}

func (this *QAudioEngine) SetDistanceScale(scale float32) {
	QAudioEngine_SetDistanceScale(this.h, (float)(scale))
}

func (this *QAudioEngine) DistanceScale() float32 {
	return (float32)(QAudioEngine_DistanceScale(this.h))
}

func (this *QAudioEngine) OutputModeChanged() {
	QAudioEngine_OutputModeChanged(this.h)
}
func (this *QAudioEngine) OnOutputModeChanged(slot func()) {
	QAudioEngine_connect_OutputModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_OutputModeChanged
func miqt_exec_callback_QAudioEngine_OutputModeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioEngine) OutputDeviceChanged() {
	QAudioEngine_OutputDeviceChanged(this.h)
}
func (this *QAudioEngine) OnOutputDeviceChanged(slot func()) {
	QAudioEngine_connect_OutputDeviceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_OutputDeviceChanged
func miqt_exec_callback_QAudioEngine_OutputDeviceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioEngine) MasterVolumeChanged() {
	QAudioEngine_MasterVolumeChanged(this.h)
}
func (this *QAudioEngine) OnMasterVolumeChanged(slot func()) {
	QAudioEngine_connect_MasterVolumeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_MasterVolumeChanged
func miqt_exec_callback_QAudioEngine_MasterVolumeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioEngine) PausedChanged() {
	QAudioEngine_PausedChanged(this.h)
}
func (this *QAudioEngine) OnPausedChanged(slot func()) {
	QAudioEngine_connect_PausedChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_PausedChanged
func miqt_exec_callback_QAudioEngine_PausedChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioEngine) DistanceScaleChanged() {
	QAudioEngine_DistanceScaleChanged(this.h)
}
func (this *QAudioEngine) OnDistanceScaleChanged(slot func()) {
	QAudioEngine_connect_DistanceScaleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_DistanceScaleChanged
func miqt_exec_callback_QAudioEngine_DistanceScaleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioEngine) Start() {
	QAudioEngine_Start(this.h)
}

func (this *QAudioEngine) Stop() {
	QAudioEngine_Stop(this.h)
}

func (this *QAudioEngine) Pause() {
	QAudioEngine_Pause(this.h)
}

func (this *QAudioEngine) Resume() {
	QAudioEngine_Resume(this.h)
}

func QAudioEngine_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioEngine_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioEngine_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioEngine_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioEngine) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QAudioEngine_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QAudioEngine) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_Event
func miqt_exec_callback_QAudioEngine_Event(self QAudioEngine, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QAudioEngine{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAudioEngine) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QAudioEngine_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QAudioEngine) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_EventFilter
func miqt_exec_callback_QAudioEngine_EventFilter(self QAudioEngine, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QAudioEngine{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAudioEngine) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QAudioEngine_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QAudioEngine) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_TimerEvent
func miqt_exec_callback_QAudioEngine_TimerEvent(self QAudioEngine, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QAudioEngine{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAudioEngine) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QAudioEngine_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QAudioEngine) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_ChildEvent
func miqt_exec_callback_QAudioEngine_ChildEvent(self QAudioEngine, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QAudioEngine{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QAudioEngine) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QAudioEngine_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QAudioEngine) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_CustomEvent
func miqt_exec_callback_QAudioEngine_CustomEvent(self QAudioEngine, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QAudioEngine{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QAudioEngine) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QAudioEngine_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QAudioEngine) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_ConnectNotify
func miqt_exec_callback_QAudioEngine_ConnectNotify(self QAudioEngine, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QAudioEngine{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QAudioEngine) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QAudioEngine_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QAudioEngine) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_DisconnectNotify
func miqt_exec_callback_QAudioEngine_DisconnectNotify(self QAudioEngine, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QAudioEngine{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
