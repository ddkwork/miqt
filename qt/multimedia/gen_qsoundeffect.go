package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSoundEffect__Loop int

const (
	QSoundEffect__Infinite QSoundEffect__Loop = -2
)

type QSoundEffect__Status int

const (
	QSoundEffect__Null    QSoundEffect__Status = 0
	QSoundEffect__Loading QSoundEffect__Status = 1
	QSoundEffect__Ready   QSoundEffect__Status = 2
	QSoundEffect__Error   QSoundEffect__Status = 3
)

type QSoundEffect struct {
	h          uintptr
	isSubclass bool
}

// NewQSoundEffect constructs a new QSoundEffect object.
func NewQSoundEffect() *QSoundEffect {
	ret := newQSoundEffect(QSoundEffect_new())
	ret.isSubclass = true
	return ret
}

// NewQSoundEffect2 constructs a new QSoundEffect object.
func NewQSoundEffect2(audioDevice *QAudioDevice) *QSoundEffect {
	ret := newQSoundEffect(QSoundEffect_new2(audioDevice.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSoundEffect3 constructs a new QSoundEffect object.
func NewQSoundEffect3(parent *qt.QObject) *QSoundEffect {
	ret := newQSoundEffect(QSoundEffect_new3((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQSoundEffect4 constructs a new QSoundEffect object.
func NewQSoundEffect4(audioDevice *QAudioDevice, parent *qt.QObject) *QSoundEffect {
	ret := newQSoundEffect(QSoundEffect_new4(audioDevice.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QSoundEffect) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSoundEffect_MetaObject(this.h)))
}

func (this *QSoundEffect) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSoundEffect_Metacast(this.h, param1_Cstring))
}

func QSoundEffect_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSoundEffect_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSoundEffect_SupportedMimeTypes() []string {
	var _ma struct_miqt_array = QSoundEffect_SupportedMimeTypes()
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

func (this *QSoundEffect) Source() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QSoundEffect_Source(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSoundEffect) SetSource(url *qt.QUrl) {
	QSoundEffect_SetSource(this.h, (*QUrl)(url.UnsafePointer()))
}

func (this *QSoundEffect) LoopCount() int {
	return (int)(QSoundEffect_LoopCount(this.h))
}

func (this *QSoundEffect) LoopsRemaining() int {
	return (int)(QSoundEffect_LoopsRemaining(this.h))
}

func (this *QSoundEffect) SetLoopCount(loopCount int) {
	QSoundEffect_SetLoopCount(this.h, (int)(loopCount))
}

func (this *QSoundEffect) AudioDevice() *QAudioDevice {
	_goptr := newQAudioDevice(QSoundEffect_AudioDevice(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSoundEffect) SetAudioDevice(device *QAudioDevice) {
	QSoundEffect_SetAudioDevice(this.h, device.cPointer())
}

func (this *QSoundEffect) Volume() float32 {
	return (float32)(QSoundEffect_Volume(this.h))
}

func (this *QSoundEffect) SetVolume(volume float32) {
	QSoundEffect_SetVolume(this.h, (float)(volume))
}

func (this *QSoundEffect) IsMuted() bool {
	return (bool)(QSoundEffect_IsMuted(this.h))
}

func (this *QSoundEffect) SetMuted(muted bool) {
	QSoundEffect_SetMuted(this.h, (bool)(muted))
}

func (this *QSoundEffect) IsLoaded() bool {
	return (bool)(QSoundEffect_IsLoaded(this.h))
}

func (this *QSoundEffect) IsPlaying() bool {
	return (bool)(QSoundEffect_IsPlaying(this.h))
}

func (this *QSoundEffect) Status() Status {
	xxxxxxxxx
}

func (this *QSoundEffect) SourceChanged() {
	QSoundEffect_SourceChanged(this.h)
}

func (this *QSoundEffect) OnSourceChanged(slot func()) {
	QSoundEffect_connect_SourceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_SourceChanged
func miqt_exec_callback_QSoundEffect_SourceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) LoopCountChanged() {
	QSoundEffect_LoopCountChanged(this.h)
}

func (this *QSoundEffect) OnLoopCountChanged(slot func()) {
	QSoundEffect_connect_LoopCountChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_LoopCountChanged
func miqt_exec_callback_QSoundEffect_LoopCountChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) LoopsRemainingChanged() {
	QSoundEffect_LoopsRemainingChanged(this.h)
}

func (this *QSoundEffect) OnLoopsRemainingChanged(slot func()) {
	QSoundEffect_connect_LoopsRemainingChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_LoopsRemainingChanged
func miqt_exec_callback_QSoundEffect_LoopsRemainingChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) VolumeChanged() {
	QSoundEffect_VolumeChanged(this.h)
}

func (this *QSoundEffect) OnVolumeChanged(slot func()) {
	QSoundEffect_connect_VolumeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_VolumeChanged
func miqt_exec_callback_QSoundEffect_VolumeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) MutedChanged() {
	QSoundEffect_MutedChanged(this.h)
}

func (this *QSoundEffect) OnMutedChanged(slot func()) {
	QSoundEffect_connect_MutedChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_MutedChanged
func miqt_exec_callback_QSoundEffect_MutedChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) LoadedChanged() {
	QSoundEffect_LoadedChanged(this.h)
}

func (this *QSoundEffect) OnLoadedChanged(slot func()) {
	QSoundEffect_connect_LoadedChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_LoadedChanged
func miqt_exec_callback_QSoundEffect_LoadedChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) PlayingChanged() {
	QSoundEffect_PlayingChanged(this.h)
}

func (this *QSoundEffect) OnPlayingChanged(slot func()) {
	QSoundEffect_connect_PlayingChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_PlayingChanged
func miqt_exec_callback_QSoundEffect_PlayingChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) StatusChanged() {
	QSoundEffect_StatusChanged(this.h)
}

func (this *QSoundEffect) OnStatusChanged(slot func()) {
	QSoundEffect_connect_StatusChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_StatusChanged
func miqt_exec_callback_QSoundEffect_StatusChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) AudioDeviceChanged() {
	QSoundEffect_AudioDeviceChanged(this.h)
}

func (this *QSoundEffect) OnAudioDeviceChanged(slot func()) {
	QSoundEffect_connect_AudioDeviceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_AudioDeviceChanged
func miqt_exec_callback_QSoundEffect_AudioDeviceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSoundEffect) Play() {
	QSoundEffect_Play(this.h)
}

func (this *QSoundEffect) Stop() {
	QSoundEffect_Stop(this.h)
}

func QSoundEffect_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSoundEffect_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSoundEffect_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSoundEffect_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSoundEffect) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSoundEffect_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QSoundEffect) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSoundEffect_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_MetaObject
func miqt_exec_callback_QSoundEffect_MetaObject(self QSoundEffect, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSoundEffect{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QSoundEffect) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSoundEffect_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSoundEffect) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSoundEffect_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSoundEffect_Metacast
func miqt_exec_callback_QSoundEffect_Metacast(self QSoundEffect, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QSoundEffect{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
