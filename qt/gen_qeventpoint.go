package qt

import (
	"unsafe"
)

type QEventPoint__State byte

const (
	QEventPoint__Unknown    QEventPoint__State = 0
	QEventPoint__Stationary QEventPoint__State = 4
	QEventPoint__Pressed    QEventPoint__State = 1
	QEventPoint__Updated    QEventPoint__State = 2
	QEventPoint__Released   QEventPoint__State = 8
)

type QEventPoint struct {
	h          uintptr
	isSubclass bool
}

// NewQEventPoint constructs a new QEventPoint object.
func NewQEventPoint() *QEventPoint {
	ret := newQEventPoint(QEventPoint_new())
	ret.isSubclass = true
	return ret
}

// NewQEventPoint2 constructs a new QEventPoint object.
func NewQEventPoint2(pointId int, state State, scenePosition *QPointF, globalPosition *QPointF) *QEventPoint {
	ret := newQEventPoint(QEventPoint_new2((int)(pointId), state, scenePosition.cPointer(), globalPosition.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQEventPoint3 constructs a new QEventPoint object.
func NewQEventPoint3(other *QEventPoint) *QEventPoint {
	ret := newQEventPoint(QEventPoint_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQEventPoint4 constructs a new QEventPoint object.
func NewQEventPoint4(id int) *QEventPoint {
	ret := newQEventPoint(QEventPoint_new4((int)(id)))
	ret.isSubclass = true
	return ret
}

// NewQEventPoint5 constructs a new QEventPoint object.
func NewQEventPoint5(id int, device *QPointingDevice) *QEventPoint {
	ret := newQEventPoint(QEventPoint_new5((int)(id), device.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QEventPoint) OperatorAssign(other *QEventPoint) {
	QEventPoint_OperatorAssign(this.h, other.cPointer())
}

func (this *QEventPoint) OperatorEqual(other *QEventPoint) bool {
	return (bool)(QEventPoint_OperatorEqual(this.h, other.cPointer()))
}

func (this *QEventPoint) OperatorNotEqual(other *QEventPoint) bool {
	return (bool)(QEventPoint_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QEventPoint) Swap(other *QEventPoint) {
	QEventPoint_Swap(this.h, other.cPointer())
}

func (this *QEventPoint) Position() *QPointF {
	_goptr := newQPointF(QEventPoint_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) PressPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_PressPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GrabPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_GrabPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_LastPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScenePosition() *QPointF {
	_goptr := newQPointF(QEventPoint_ScenePosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScenePressPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_ScenePressPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) SceneGrabPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_SceneGrabPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) SceneLastPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_SceneLastPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_GlobalPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalPressPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_GlobalPressPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalGrabPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_GlobalGrabPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) GlobalLastPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_GlobalLastPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) NormalizedPosition() *QPointF {
	_goptr := newQPointF(QEventPoint_NormalizedPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) Pos() *QPointF {
	_goptr := newQPointF(QEventPoint_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartPos() *QPointF {
	_goptr := newQPointF(QEventPoint_StartPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScenePos() *QPointF {
	_goptr := newQPointF(QEventPoint_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartScenePos() *QPointF {
	_goptr := newQPointF(QEventPoint_StartScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) ScreenPos() *QPointF {
	_goptr := newQPointF(QEventPoint_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartScreenPos() *QPointF {
	_goptr := newQPointF(QEventPoint_StartScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) StartNormalizedPos() *QPointF {
	_goptr := newQPointF(QEventPoint_StartNormalizedPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) NormalizedPos() *QPointF {
	_goptr := newQPointF(QEventPoint_NormalizedPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastPos() *QPointF {
	_goptr := newQPointF(QEventPoint_LastPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastScenePos() *QPointF {
	_goptr := newQPointF(QEventPoint_LastScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastScreenPos() *QPointF {
	_goptr := newQPointF(QEventPoint_LastScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) LastNormalizedPos() *QPointF {
	_goptr := newQPointF(QEventPoint_LastNormalizedPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) Velocity() *QVector2D {
	_goptr := newQVector2D(QEventPoint_Velocity(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) State() State {
	xxxxxxxxx
}

func (this *QEventPoint) Device() *QPointingDevice {
	return newQPointingDevice(QEventPoint_Device(this.h))
}

func (this *QEventPoint) Id() int {
	return (int)(QEventPoint_Id(this.h))
}

func (this *QEventPoint) UniqueId() *QPointingDeviceUniqueId {
	_goptr := newQPointingDeviceUniqueId(QEventPoint_UniqueId(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) Timestamp() uint32 {
	return (uint32)(QEventPoint_Timestamp(this.h))
}

func (this *QEventPoint) LastTimestamp() uint32 {
	return (uint32)(QEventPoint_LastTimestamp(this.h))
}

func (this *QEventPoint) PressTimestamp() uint32 {
	return (uint32)(QEventPoint_PressTimestamp(this.h))
}

func (this *QEventPoint) TimeHeld() float64 {
	return (float64)(QEventPoint_TimeHeld(this.h))
}

func (this *QEventPoint) Pressure() float64 {
	return (float64)(QEventPoint_Pressure(this.h))
}

func (this *QEventPoint) Rotation() float64 {
	return (float64)(QEventPoint_Rotation(this.h))
}

func (this *QEventPoint) EllipseDiameters() *QSizeF {
	_goptr := newQSizeF(QEventPoint_EllipseDiameters(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEventPoint) IsAccepted() bool {
	return (bool)(QEventPoint_IsAccepted(this.h))
}

func (this *QEventPoint) SetAccepted() {
	QEventPoint_SetAccepted(this.h)
}

func (this *QEventPoint) SetAccepted1(accepted bool) {
	QEventPoint_SetAccepted1(this.h, (bool)(accepted))
}
