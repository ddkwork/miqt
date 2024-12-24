package spatialaudio

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
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

func (this *QAmbientSound) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QAmbientSound_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QAmbientSound) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_Event
func miqt_exec_callback_QAmbientSound_Event(self QAmbientSound, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QAmbientSound{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAmbientSound) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QAmbientSound_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QAmbientSound) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_EventFilter
func miqt_exec_callback_QAmbientSound_EventFilter(self QAmbientSound, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QAmbientSound{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAmbientSound) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QAmbientSound_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QAmbientSound) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_TimerEvent
func miqt_exec_callback_QAmbientSound_TimerEvent(self QAmbientSound, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QAmbientSound{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAmbientSound) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QAmbientSound_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QAmbientSound) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_ChildEvent
func miqt_exec_callback_QAmbientSound_ChildEvent(self QAmbientSound, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QAmbientSound{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QAmbientSound) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QAmbientSound_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QAmbientSound) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_CustomEvent
func miqt_exec_callback_QAmbientSound_CustomEvent(self QAmbientSound, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QAmbientSound{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QAmbientSound) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QAmbientSound_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QAmbientSound) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_ConnectNotify
func miqt_exec_callback_QAmbientSound_ConnectNotify(self QAmbientSound, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QAmbientSound{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QAmbientSound) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QAmbientSound_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QAmbientSound) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAmbientSound_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAmbientSound_DisconnectNotify
func miqt_exec_callback_QAmbientSound_DisconnectNotify(self QAmbientSound, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QAmbientSound{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
