package spatialaudio

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioRoom__Material int

const (
	QAudioRoom__Transparent            QAudioRoom__Material = 0
	QAudioRoom__AcousticCeilingTiles   QAudioRoom__Material = 1
	QAudioRoom__BrickBare              QAudioRoom__Material = 2
	QAudioRoom__BrickPainted           QAudioRoom__Material = 3
	QAudioRoom__ConcreteBlockCoarse    QAudioRoom__Material = 4
	QAudioRoom__ConcreteBlockPainted   QAudioRoom__Material = 5
	QAudioRoom__CurtainHeavy           QAudioRoom__Material = 6
	QAudioRoom__FiberGlassInsulation   QAudioRoom__Material = 7
	QAudioRoom__GlassThin              QAudioRoom__Material = 8
	QAudioRoom__GlassThick             QAudioRoom__Material = 9
	QAudioRoom__Grass                  QAudioRoom__Material = 10
	QAudioRoom__LinoleumOnConcrete     QAudioRoom__Material = 11
	QAudioRoom__Marble                 QAudioRoom__Material = 12
	QAudioRoom__Metal                  QAudioRoom__Material = 13
	QAudioRoom__ParquetOnConcrete      QAudioRoom__Material = 14
	QAudioRoom__PlasterRough           QAudioRoom__Material = 15
	QAudioRoom__PlasterSmooth          QAudioRoom__Material = 16
	QAudioRoom__PlywoodPanel           QAudioRoom__Material = 17
	QAudioRoom__PolishedConcreteOrTile QAudioRoom__Material = 18
	QAudioRoom__Sheetrock              QAudioRoom__Material = 19
	QAudioRoom__WaterOrIceSurface      QAudioRoom__Material = 20
	QAudioRoom__WoodCeiling            QAudioRoom__Material = 21
	QAudioRoom__WoodPanel              QAudioRoom__Material = 22
	QAudioRoom__UniformMaterial        QAudioRoom__Material = 23
)

type QAudioRoom__Wall int

const (
	QAudioRoom__LeftWall  QAudioRoom__Wall = 0
	QAudioRoom__RightWall QAudioRoom__Wall = 1
	QAudioRoom__Floor     QAudioRoom__Wall = 2
	QAudioRoom__Ceiling   QAudioRoom__Wall = 3
	QAudioRoom__FrontWall QAudioRoom__Wall = 4
	QAudioRoom__BackWall  QAudioRoom__Wall = 5
)

type QAudioRoom struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioRoom constructs a new QAudioRoom object.
func NewQAudioRoom(engine *QAudioEngine) *QAudioRoom {
	ret := newQAudioRoom(QAudioRoom_new(engine.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAudioRoom) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioRoom_MetaObject(this.h)))
}

func (this *QAudioRoom) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioRoom_Metacast(this.h, param1_Cstring))
}

func QAudioRoom_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioRoom_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioRoom) SetPosition(pos qt.QVector3D) {
	QAudioRoom_SetPosition(this.h, (*QVector3D)(pos.UnsafePointer()))
}

func (this *QAudioRoom) Position() *qt.QVector3D {
	_goptr := qt.UnsafeNewQVector3D(unsafe.Pointer(QAudioRoom_Position(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioRoom) SetDimensions(dim qt.QVector3D) {
	QAudioRoom_SetDimensions(this.h, (*QVector3D)(dim.UnsafePointer()))
}

func (this *QAudioRoom) Dimensions() *qt.QVector3D {
	_goptr := qt.UnsafeNewQVector3D(unsafe.Pointer(QAudioRoom_Dimensions(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioRoom) SetRotation(q *qt.QQuaternion) {
	QAudioRoom_SetRotation(this.h, (*QQuaternion)(q.UnsafePointer()))
}

func (this *QAudioRoom) Rotation() *qt.QQuaternion {
	_goptr := qt.UnsafeNewQQuaternion(unsafe.Pointer(QAudioRoom_Rotation(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioRoom) SetWallMaterial(wall Wall, material Material) {
	QAudioRoom_SetWallMaterial(this.h, wall, material)
}

func (this *QAudioRoom) WallMaterial(wall Wall) Material {
	xxxxxxxxx
}

func (this *QAudioRoom) SetReflectionGain(factor float32) {
	QAudioRoom_SetReflectionGain(this.h, (float)(factor))
}

func (this *QAudioRoom) ReflectionGain() float32 {
	return (float32)(QAudioRoom_ReflectionGain(this.h))
}

func (this *QAudioRoom) SetReverbGain(factor float32) {
	QAudioRoom_SetReverbGain(this.h, (float)(factor))
}

func (this *QAudioRoom) ReverbGain() float32 {
	return (float32)(QAudioRoom_ReverbGain(this.h))
}

func (this *QAudioRoom) SetReverbTime(factor float32) {
	QAudioRoom_SetReverbTime(this.h, (float)(factor))
}

func (this *QAudioRoom) ReverbTime() float32 {
	return (float32)(QAudioRoom_ReverbTime(this.h))
}

func (this *QAudioRoom) SetReverbBrightness(factor float32) {
	QAudioRoom_SetReverbBrightness(this.h, (float)(factor))
}

func (this *QAudioRoom) ReverbBrightness() float32 {
	return (float32)(QAudioRoom_ReverbBrightness(this.h))
}

func (this *QAudioRoom) PositionChanged() {
	QAudioRoom_PositionChanged(this.h)
}

func (this *QAudioRoom) OnPositionChanged(slot func()) {
	QAudioRoom_connect_PositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_PositionChanged
func miqt_exec_callback_QAudioRoom_PositionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioRoom) DimensionsChanged() {
	QAudioRoom_DimensionsChanged(this.h)
}

func (this *QAudioRoom) OnDimensionsChanged(slot func()) {
	QAudioRoom_connect_DimensionsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_DimensionsChanged
func miqt_exec_callback_QAudioRoom_DimensionsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioRoom) RotationChanged() {
	QAudioRoom_RotationChanged(this.h)
}

func (this *QAudioRoom) OnRotationChanged(slot func()) {
	QAudioRoom_connect_RotationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_RotationChanged
func miqt_exec_callback_QAudioRoom_RotationChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioRoom) WallsChanged() {
	QAudioRoom_WallsChanged(this.h)
}

func (this *QAudioRoom) OnWallsChanged(slot func()) {
	QAudioRoom_connect_WallsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_WallsChanged
func miqt_exec_callback_QAudioRoom_WallsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioRoom) ReflectionGainChanged() {
	QAudioRoom_ReflectionGainChanged(this.h)
}

func (this *QAudioRoom) OnReflectionGainChanged(slot func()) {
	QAudioRoom_connect_ReflectionGainChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_ReflectionGainChanged
func miqt_exec_callback_QAudioRoom_ReflectionGainChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioRoom) ReverbGainChanged() {
	QAudioRoom_ReverbGainChanged(this.h)
}

func (this *QAudioRoom) OnReverbGainChanged(slot func()) {
	QAudioRoom_connect_ReverbGainChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_ReverbGainChanged
func miqt_exec_callback_QAudioRoom_ReverbGainChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioRoom) ReverbTimeChanged() {
	QAudioRoom_ReverbTimeChanged(this.h)
}

func (this *QAudioRoom) OnReverbTimeChanged(slot func()) {
	QAudioRoom_connect_ReverbTimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_ReverbTimeChanged
func miqt_exec_callback_QAudioRoom_ReverbTimeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioRoom) ReverbBrightnessChanged() {
	QAudioRoom_ReverbBrightnessChanged(this.h)
}

func (this *QAudioRoom) OnReverbBrightnessChanged(slot func()) {
	QAudioRoom_connect_ReverbBrightnessChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_ReverbBrightnessChanged
func miqt_exec_callback_QAudioRoom_ReverbBrightnessChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QAudioRoom_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioRoom_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioRoom_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioRoom_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioRoom) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioRoom_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioRoom) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioRoom_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_MetaObject
func miqt_exec_callback_QAudioRoom_MetaObject(self QAudioRoom, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioRoom{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioRoom) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioRoom_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioRoom) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioRoom_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioRoom_Metacast
func miqt_exec_callback_QAudioRoom_Metacast(self QAudioRoom, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAudioRoom{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
