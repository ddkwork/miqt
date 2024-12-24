package spatialaudio

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
	"github.com/mappu/miqt/qt/multimedia"
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
	g := newQAudioEngine(QAudioEngine_new())
	g.isSubclass = true
	return g
}

// NewQAudioEngine2 constructs a new QAudioEngine object.
func NewQAudioEngine2(parent *qt.QObject) *QAudioEngine {
	g := newQAudioEngine(QAudioEngine_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQAudioEngine3 constructs a new QAudioEngine object.
func NewQAudioEngine3(sampleRate int) *QAudioEngine {
	g := newQAudioEngine(QAudioEngine_new3((int)(sampleRate)))
	g.isSubclass = true
	return g
}

// NewQAudioEngine4 constructs a new QAudioEngine object.
func NewQAudioEngine4(sampleRate int, parent *qt.QObject) *QAudioEngine {
	g := newQAudioEngine(QAudioEngine_new4((int)(sampleRate), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
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

func (this *QAudioEngine) SetOutputDevice(device *multimedia.QAudioDevice) {
	QAudioEngine_SetOutputDevice(this.h, (*QAudioDevice)(device.UnsafePointer()))
}

func (this *QAudioEngine) OutputDevice() *multimedia.QAudioDevice {
	_goptr := multimedia.UnsafeNewQAudioDevice(unsafe.Pointer(QAudioEngine_OutputDevice(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
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

func (this *QAudioEngine) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioEngine_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioEngine) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_MetaObject
func miqt_exec_callback_QAudioEngine_MetaObject(self QAudioEngine, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioEngine{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioEngine) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioEngine_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioEngine) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioEngine_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioEngine_Metacast
func miqt_exec_callback_QAudioEngine_Metacast(self QAudioEngine, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QAudioEngine{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
