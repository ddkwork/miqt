package spatialaudio

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAmbientSound__Loops int

const (
	QAmbientSound__Infinite QAmbientSound__Loops = -1
	QAmbientSound__Once     QAmbientSound__Loops = 1
)

type QAmbientSound struct {
	h          uintptr
	isSubclass bool
}

// NewQAmbientSound constructs a new QAmbientSound object.
func NewQAmbientSound(engine *QAudioEngine) *QAmbientSound {
	ret := newQAmbientSound(QAmbientSound_new(engine.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAmbientSound) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAmbientSound_MetaObject(this.h)))
}

func (this *QAmbientSound) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAmbientSound_Metacast(this.h, param1_Cstring))
}

func QAmbientSound_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAmbientSound_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAmbientSound) SetSource(url *qt.QUrl) {
	QAmbientSound_SetSource(this.h, (*QUrl)(url.UnsafePointer()))
}

func (this *QAmbientSound) Source() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QAmbientSound_Source(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAmbientSound) Loops() int {
	return (int)(QAmbientSound_Loops(this.h))
}

func (this *QAmbientSound) SetLoops(loops int) {
	QAmbientSound_SetLoops(this.h, (int)(loops))
}

func (this *QAmbientSound) AutoPlay() bool {
	return (bool)(QAmbientSound_AutoPlay(this.h))
}

func (this *QAmbientSound) SetAutoPlay(autoPlay bool) {
	QAmbientSound_SetAutoPlay(this.h, (bool)(autoPlay))
}

func (this *QAmbientSound) SetVolume(volume float32) {
	QAmbientSound_SetVolume(this.h, (float)(volume))
}

func (this *QAmbientSound) Volume() float32 {
	return (float32)(QAmbientSound_Volume(this.h))
}

func (this *QAmbientSound) Engine() *QAudioEngine {
	return newQAudioEngine(QAmbientSound_Engine(this.h))
}

func (this *QAmbientSound) SourceChanged() {
	QAmbientSound_SourceChanged(this.h)
}

func (this *QAmbientSound) OnSourceChanged(slot func()) {
	QAmbientSound_connect_SourceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_SourceChanged
func miqt_exec_callback_QAmbientSound_SourceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAmbientSound) LoopsChanged() {
	QAmbientSound_LoopsChanged(this.h)
}

func (this *QAmbientSound) OnLoopsChanged(slot func()) {
	QAmbientSound_connect_LoopsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_LoopsChanged
func miqt_exec_callback_QAmbientSound_LoopsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAmbientSound) AutoPlayChanged() {
	QAmbientSound_AutoPlayChanged(this.h)
}

func (this *QAmbientSound) OnAutoPlayChanged(slot func()) {
	QAmbientSound_connect_AutoPlayChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_AutoPlayChanged
func miqt_exec_callback_QAmbientSound_AutoPlayChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAmbientSound) VolumeChanged() {
	QAmbientSound_VolumeChanged(this.h)
}

func (this *QAmbientSound) OnVolumeChanged(slot func()) {
	QAmbientSound_connect_VolumeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_VolumeChanged
func miqt_exec_callback_QAmbientSound_VolumeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAmbientSound) Play() {
	QAmbientSound_Play(this.h)
}

func (this *QAmbientSound) Pause() {
	QAmbientSound_Pause(this.h)
}

func (this *QAmbientSound) Stop() {
	QAmbientSound_Stop(this.h)
}

func QAmbientSound_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAmbientSound_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAmbientSound_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAmbientSound_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAmbientSound) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAmbientSound_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAmbientSound) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_MetaObject
func miqt_exec_callback_QAmbientSound_MetaObject(self QAmbientSound, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAmbientSound{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAmbientSound) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAmbientSound_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAmbientSound) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_Metacast
func miqt_exec_callback_QAmbientSound_Metacast(self QAmbientSound, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAmbientSound{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
