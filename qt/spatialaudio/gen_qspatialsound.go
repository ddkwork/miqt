package spatialaudio

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QSpatialSound__Loops int

const (
	QSpatialSound__Infinite QSpatialSound__Loops = -1
	QSpatialSound__Once     QSpatialSound__Loops = 1
)

type QSpatialSound__DistanceModel int

const (
	QSpatialSound__Logarithmic       QSpatialSound__DistanceModel = 0
	QSpatialSound__Linear            QSpatialSound__DistanceModel = 1
	QSpatialSound__ManualAttenuation QSpatialSound__DistanceModel = 2
)

type QSpatialSound struct {
	h          uintptr
	isSubclass bool
}

// NewQSpatialSound constructs a new QSpatialSound object.
func NewQSpatialSound(engine *QAudioEngine) *QSpatialSound {

	ret := newQSpatialSound(QSpatialSound_new(engine.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSpatialSound) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSpatialSound_MetaObject(this.h)))
}

func (this *QSpatialSound) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSpatialSound_Metacast(this.h, param1_Cstring))
}

func QSpatialSound_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSpatialSound_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpatialSound) SetSource(url *qt.QUrl) {
	QSpatialSound_SetSource(this.h, (*QUrl)(url.UnsafePointer()))
}

func (this *QSpatialSound) Source() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QSpatialSound_Source(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpatialSound) Loops() int {
	return (int)(QSpatialSound_Loops(this.h))
}

func (this *QSpatialSound) SetLoops(loops int) {
	QSpatialSound_SetLoops(this.h, (int)(loops))
}

func (this *QSpatialSound) AutoPlay() bool {
	return (bool)(QSpatialSound_AutoPlay(this.h))
}

func (this *QSpatialSound) SetAutoPlay(autoPlay bool) {
	QSpatialSound_SetAutoPlay(this.h, (bool)(autoPlay))
}

func (this *QSpatialSound) SetPosition(pos qt.QVector3D) {
	QSpatialSound_SetPosition(this.h, (*QVector3D)(pos.UnsafePointer()))
}

func (this *QSpatialSound) Position() *qt.QVector3D {
	_goptr := qt.UnsafeNewQVector3D(unsafe.Pointer(QSpatialSound_Position(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpatialSound) SetRotation(q *qt.QQuaternion) {
	QSpatialSound_SetRotation(this.h, (*QQuaternion)(q.UnsafePointer()))
}

func (this *QSpatialSound) Rotation() *qt.QQuaternion {
	_goptr := qt.UnsafeNewQQuaternion(unsafe.Pointer(QSpatialSound_Rotation(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpatialSound) SetVolume(volume float32) {
	QSpatialSound_SetVolume(this.h, (float)(volume))
}

func (this *QSpatialSound) Volume() float32 {
	return (float32)(QSpatialSound_Volume(this.h))
}

func (this *QSpatialSound) SetDistanceModel(model DistanceModel) {
	QSpatialSound_SetDistanceModel(this.h, model)
}

func (this *QSpatialSound) DistanceModel() DistanceModel {
	xxxxxxxxx
}

func (this *QSpatialSound) SetSize(size float32) {
	QSpatialSound_SetSize(this.h, (float)(size))
}

func (this *QSpatialSound) Size() float32 {
	return (float32)(QSpatialSound_Size(this.h))
}

func (this *QSpatialSound) SetDistanceCutoff(cutoff float32) {
	QSpatialSound_SetDistanceCutoff(this.h, (float)(cutoff))
}

func (this *QSpatialSound) DistanceCutoff() float32 {
	return (float32)(QSpatialSound_DistanceCutoff(this.h))
}

func (this *QSpatialSound) SetManualAttenuation(attenuation float32) {
	QSpatialSound_SetManualAttenuation(this.h, (float)(attenuation))
}

func (this *QSpatialSound) ManualAttenuation() float32 {
	return (float32)(QSpatialSound_ManualAttenuation(this.h))
}

func (this *QSpatialSound) SetOcclusionIntensity(occlusion float32) {
	QSpatialSound_SetOcclusionIntensity(this.h, (float)(occlusion))
}

func (this *QSpatialSound) OcclusionIntensity() float32 {
	return (float32)(QSpatialSound_OcclusionIntensity(this.h))
}

func (this *QSpatialSound) SetDirectivity(alpha float32) {
	QSpatialSound_SetDirectivity(this.h, (float)(alpha))
}

func (this *QSpatialSound) Directivity() float32 {
	return (float32)(QSpatialSound_Directivity(this.h))
}

func (this *QSpatialSound) SetDirectivityOrder(alpha float32) {
	QSpatialSound_SetDirectivityOrder(this.h, (float)(alpha))
}

func (this *QSpatialSound) DirectivityOrder() float32 {
	return (float32)(QSpatialSound_DirectivityOrder(this.h))
}

func (this *QSpatialSound) SetNearFieldGain(gain float32) {
	QSpatialSound_SetNearFieldGain(this.h, (float)(gain))
}

func (this *QSpatialSound) NearFieldGain() float32 {
	return (float32)(QSpatialSound_NearFieldGain(this.h))
}

func (this *QSpatialSound) Engine() *QAudioEngine {
	return newQAudioEngine(QSpatialSound_Engine(this.h))
}

func (this *QSpatialSound) SourceChanged() {
	QSpatialSound_SourceChanged(this.h)
}
func (this *QSpatialSound) OnSourceChanged(slot func()) {
	QSpatialSound_connect_SourceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_SourceChanged
func miqt_exec_callback_QSpatialSound_SourceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) LoopsChanged() {
	QSpatialSound_LoopsChanged(this.h)
}
func (this *QSpatialSound) OnLoopsChanged(slot func()) {
	QSpatialSound_connect_LoopsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_LoopsChanged
func miqt_exec_callback_QSpatialSound_LoopsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) AutoPlayChanged() {
	QSpatialSound_AutoPlayChanged(this.h)
}
func (this *QSpatialSound) OnAutoPlayChanged(slot func()) {
	QSpatialSound_connect_AutoPlayChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_AutoPlayChanged
func miqt_exec_callback_QSpatialSound_AutoPlayChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) PositionChanged() {
	QSpatialSound_PositionChanged(this.h)
}
func (this *QSpatialSound) OnPositionChanged(slot func()) {
	QSpatialSound_connect_PositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_PositionChanged
func miqt_exec_callback_QSpatialSound_PositionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) RotationChanged() {
	QSpatialSound_RotationChanged(this.h)
}
func (this *QSpatialSound) OnRotationChanged(slot func()) {
	QSpatialSound_connect_RotationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_RotationChanged
func miqt_exec_callback_QSpatialSound_RotationChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) VolumeChanged() {
	QSpatialSound_VolumeChanged(this.h)
}
func (this *QSpatialSound) OnVolumeChanged(slot func()) {
	QSpatialSound_connect_VolumeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_VolumeChanged
func miqt_exec_callback_QSpatialSound_VolumeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) DistanceModelChanged() {
	QSpatialSound_DistanceModelChanged(this.h)
}
func (this *QSpatialSound) OnDistanceModelChanged(slot func()) {
	QSpatialSound_connect_DistanceModelChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_DistanceModelChanged
func miqt_exec_callback_QSpatialSound_DistanceModelChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) SizeChanged() {
	QSpatialSound_SizeChanged(this.h)
}
func (this *QSpatialSound) OnSizeChanged(slot func()) {
	QSpatialSound_connect_SizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_SizeChanged
func miqt_exec_callback_QSpatialSound_SizeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) DistanceCutoffChanged() {
	QSpatialSound_DistanceCutoffChanged(this.h)
}
func (this *QSpatialSound) OnDistanceCutoffChanged(slot func()) {
	QSpatialSound_connect_DistanceCutoffChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_DistanceCutoffChanged
func miqt_exec_callback_QSpatialSound_DistanceCutoffChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) ManualAttenuationChanged() {
	QSpatialSound_ManualAttenuationChanged(this.h)
}
func (this *QSpatialSound) OnManualAttenuationChanged(slot func()) {
	QSpatialSound_connect_ManualAttenuationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_ManualAttenuationChanged
func miqt_exec_callback_QSpatialSound_ManualAttenuationChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) OcclusionIntensityChanged() {
	QSpatialSound_OcclusionIntensityChanged(this.h)
}
func (this *QSpatialSound) OnOcclusionIntensityChanged(slot func()) {
	QSpatialSound_connect_OcclusionIntensityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_OcclusionIntensityChanged
func miqt_exec_callback_QSpatialSound_OcclusionIntensityChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) DirectivityChanged() {
	QSpatialSound_DirectivityChanged(this.h)
}
func (this *QSpatialSound) OnDirectivityChanged(slot func()) {
	QSpatialSound_connect_DirectivityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_DirectivityChanged
func miqt_exec_callback_QSpatialSound_DirectivityChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) DirectivityOrderChanged() {
	QSpatialSound_DirectivityOrderChanged(this.h)
}
func (this *QSpatialSound) OnDirectivityOrderChanged(slot func()) {
	QSpatialSound_connect_DirectivityOrderChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_DirectivityOrderChanged
func miqt_exec_callback_QSpatialSound_DirectivityOrderChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) NearFieldGainChanged() {
	QSpatialSound_NearFieldGainChanged(this.h)
}
func (this *QSpatialSound) OnNearFieldGainChanged(slot func()) {
	QSpatialSound_connect_NearFieldGainChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_NearFieldGainChanged
func miqt_exec_callback_QSpatialSound_NearFieldGainChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QSpatialSound) Play() {
	QSpatialSound_Play(this.h)
}

func (this *QSpatialSound) Pause() {
	QSpatialSound_Pause(this.h)
}

func (this *QSpatialSound) Stop() {
	QSpatialSound_Stop(this.h)
}

func QSpatialSound_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSpatialSound_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSpatialSound_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSpatialSound_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpatialSound) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QSpatialSound_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QSpatialSound) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpatialSound_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_Event
func miqt_exec_callback_QSpatialSound_Event(self QSpatialSound, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QSpatialSound{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSpatialSound) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QSpatialSound_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QSpatialSound) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpatialSound_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_EventFilter
func miqt_exec_callback_QSpatialSound_EventFilter(self QSpatialSound, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QSpatialSound{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSpatialSound) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QSpatialSound_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QSpatialSound) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpatialSound_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_TimerEvent
func miqt_exec_callback_QSpatialSound_TimerEvent(self QSpatialSound, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QSpatialSound{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSpatialSound) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QSpatialSound_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QSpatialSound) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpatialSound_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_ChildEvent
func miqt_exec_callback_QSpatialSound_ChildEvent(self QSpatialSound, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QSpatialSound{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSpatialSound) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QSpatialSound_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QSpatialSound) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpatialSound_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_CustomEvent
func miqt_exec_callback_QSpatialSound_CustomEvent(self QSpatialSound, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QSpatialSound{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QSpatialSound) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QSpatialSound_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QSpatialSound) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpatialSound_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_ConnectNotify
func miqt_exec_callback_QSpatialSound_ConnectNotify(self QSpatialSound, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QSpatialSound{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QSpatialSound) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QSpatialSound_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QSpatialSound) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpatialSound_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpatialSound_DisconnectNotify
func miqt_exec_callback_QSpatialSound_DisconnectNotify(self QSpatialSound, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QSpatialSound{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}