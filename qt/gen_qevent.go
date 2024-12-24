package qt

import (
	"unsafe"
)

type QWheelEvent__ int

const (
	QWheelEvent__DefaultDeltasPerStep QWheelEvent__ = 120
)

type QPlatformSurfaceEvent__SurfaceEventType int

const (
	QPlatformSurfaceEvent__SurfaceCreated            QPlatformSurfaceEvent__SurfaceEventType = 0
	QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed QPlatformSurfaceEvent__SurfaceEventType = 1
)

type QContextMenuEvent__Reason int

const (
	QContextMenuEvent__Mouse    QContextMenuEvent__Reason = 0
	QContextMenuEvent__Keyboard QContextMenuEvent__Reason = 1
	QContextMenuEvent__Other    QContextMenuEvent__Reason = 2
)

type QInputMethodEvent__AttributeType int

const (
	QInputMethodEvent__TextFormat QInputMethodEvent__AttributeType = 0
	QInputMethodEvent__Cursor     QInputMethodEvent__AttributeType = 1
	QInputMethodEvent__Language   QInputMethodEvent__AttributeType = 2
	QInputMethodEvent__Ruby       QInputMethodEvent__AttributeType = 3
	QInputMethodEvent__Selection  QInputMethodEvent__AttributeType = 4
)

type QScrollEvent__ScrollState int

const (
	QScrollEvent__ScrollStarted  QScrollEvent__ScrollState = 0
	QScrollEvent__ScrollUpdated  QScrollEvent__ScrollState = 1
	QScrollEvent__ScrollFinished QScrollEvent__ScrollState = 2
)

type QInputEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQInputEvent constructs a new QInputEvent object.
func NewQInputEvent(typeVal Type, m_dev *QInputDevice) *QInputEvent {
	g := newQInputEvent(QInputEvent_new(typeVal, m_dev.cPointer()))
	g.isSubclass = true
	return g
}

// NewQInputEvent2 constructs a new QInputEvent object.
func NewQInputEvent2(typeVal Type, m_dev *QInputDevice, modifiers KeyboardModifier) *QInputEvent {
	g := newQInputEvent(QInputEvent_new2(typeVal, m_dev.cPointer(), (int)(modifiers)))
	g.isSubclass = true
	return g
}

func (this *QInputEvent) Clone() *QInputEvent {
	return newQInputEvent(QInputEvent_Clone(this.h))
}

func (this *QInputEvent) Device() *QInputDevice {
	return newQInputDevice(QInputEvent_Device(this.h))
}

func (this *QInputEvent) DeviceType() QInputDevice__DeviceType {
	return (QInputDevice__DeviceType)(QInputEvent_DeviceType(this.h))
}

func (this *QInputEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QInputEvent_Modifiers(this.h))
}

func (this *QInputEvent) SetModifiers(modifiers KeyboardModifier) {
	QInputEvent_SetModifiers(this.h, (int)(modifiers))
}

func (this *QInputEvent) Timestamp() uint64 {
	return (uint64)(QInputEvent_Timestamp(this.h))
}

func (this *QInputEvent) SetTimestamp(timestamp uint64) {
	QInputEvent_SetTimestamp(this.h, (ulonglong)(timestamp))
}

type QPointerEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQPointerEvent constructs a new QPointerEvent object.
func NewQPointerEvent(typeVal Type, dev *QPointingDevice) *QPointerEvent {
	g := newQPointerEvent(QPointerEvent_new(typeVal, dev.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPointerEvent2 constructs a new QPointerEvent object.
func NewQPointerEvent2(typeVal Type, dev *QPointingDevice, modifiers KeyboardModifier) *QPointerEvent {
	g := newQPointerEvent(QPointerEvent_new2(typeVal, dev.cPointer(), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQPointerEvent3 constructs a new QPointerEvent object.
func NewQPointerEvent3(typeVal Type, dev *QPointingDevice, modifiers KeyboardModifier, points []QEventPoint) *QPointerEvent {
	points_CArray := (*[0xffff]*QEventPoint)(malloc(size_t(8 * len(points))))
	defer free(unsafe.Pointer(points_CArray))
	for i := range points {
		points_CArray[i] = points[i].cPointer()
	}
	points_ma := struct_miqt_array{len: size_t(len(points)), data: unsafe.Pointer(points_CArray)}
	g := newQPointerEvent(QPointerEvent_new3(typeVal, dev.cPointer(), (int)(modifiers), points_ma))
	g.isSubclass = true
	return g
}

func (this *QPointerEvent) Clone() *QPointerEvent {
	return newQPointerEvent(QPointerEvent_Clone(this.h))
}

func (this *QPointerEvent) PointingDevice() *QPointingDevice {
	return newQPointingDevice(QPointerEvent_PointingDevice(this.h))
}

func (this *QPointerEvent) PointerType() QPointingDevice__PointerType {
	return (QPointingDevice__PointerType)(QPointerEvent_PointerType(this.h))
}

func (this *QPointerEvent) SetTimestamp(timestamp uint64) {
	QPointerEvent_SetTimestamp(this.h, (ulonglong)(timestamp))
}

func (this *QPointerEvent) PointCount() int64 {
	return (int64)(QPointerEvent_PointCount(this.h))
}

func (this *QPointerEvent) Point(i int64) *QEventPoint {
	return newQEventPoint(QPointerEvent_Point(this.h, (ptrdiff_t)(i)))
}

func (this *QPointerEvent) Points() []QEventPoint {
	var _ma struct_miqt_array = QPointerEvent_Points(this.h)
	_ret := make([]QEventPoint, int(_ma.len))
	_outCast := (*[0xffff]*QEventPoint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQEventPoint(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPointerEvent) PointById(id int) *QEventPoint {
	return newQEventPoint(QPointerEvent_PointById(this.h, (int)(id)))
}

func (this *QPointerEvent) AllPointsGrabbed() bool {
	return (bool)(QPointerEvent_AllPointsGrabbed(this.h))
}

func (this *QPointerEvent) IsBeginEvent() bool {
	return (bool)(QPointerEvent_IsBeginEvent(this.h))
}

func (this *QPointerEvent) IsUpdateEvent() bool {
	return (bool)(QPointerEvent_IsUpdateEvent(this.h))
}

func (this *QPointerEvent) IsEndEvent() bool {
	return (bool)(QPointerEvent_IsEndEvent(this.h))
}

func (this *QPointerEvent) AllPointsAccepted() bool {
	return (bool)(QPointerEvent_AllPointsAccepted(this.h))
}

func (this *QPointerEvent) SetAccepted(accepted bool) {
	QPointerEvent_SetAccepted(this.h, (bool)(accepted))
}

func (this *QPointerEvent) ExclusiveGrabber(point *QEventPoint) *QObject {
	return newQObject(QPointerEvent_ExclusiveGrabber(this.h, point.cPointer()))
}

func (this *QPointerEvent) SetExclusiveGrabber(point *QEventPoint, exclusiveGrabber *QObject) {
	QPointerEvent_SetExclusiveGrabber(this.h, point.cPointer(), exclusiveGrabber.cPointer())
}

func (this *QPointerEvent) ClearPassiveGrabbers(point *QEventPoint) {
	QPointerEvent_ClearPassiveGrabbers(this.h, point.cPointer())
}

func (this *QPointerEvent) AddPassiveGrabber(point *QEventPoint, grabber *QObject) bool {
	return (bool)(QPointerEvent_AddPassiveGrabber(this.h, point.cPointer(), grabber.cPointer()))
}

func (this *QPointerEvent) RemovePassiveGrabber(point *QEventPoint, grabber *QObject) bool {
	return (bool)(QPointerEvent_RemovePassiveGrabber(this.h, point.cPointer(), grabber.cPointer()))
}

type QSinglePointEvent struct {
	h          uintptr
	isSubclass bool
}

func (this *QSinglePointEvent) Clone() *QSinglePointEvent {
	return newQSinglePointEvent(QSinglePointEvent_Clone(this.h))
}

func (this *QSinglePointEvent) Button() MouseButton {
	return (MouseButton)(QSinglePointEvent_Button(this.h))
}

func (this *QSinglePointEvent) Buttons() MouseButton {
	return (MouseButton)(QSinglePointEvent_Buttons(this.h))
}

func (this *QSinglePointEvent) Position() *QPointF {
	_goptr := newQPointF(QSinglePointEvent_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSinglePointEvent) ScenePosition() *QPointF {
	_goptr := newQPointF(QSinglePointEvent_ScenePosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSinglePointEvent) GlobalPosition() *QPointF {
	_goptr := newQPointF(QSinglePointEvent_GlobalPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSinglePointEvent) IsBeginEvent() bool {
	return (bool)(QSinglePointEvent_IsBeginEvent(this.h))
}

func (this *QSinglePointEvent) IsUpdateEvent() bool {
	return (bool)(QSinglePointEvent_IsUpdateEvent(this.h))
}

func (this *QSinglePointEvent) IsEndEvent() bool {
	return (bool)(QSinglePointEvent_IsEndEvent(this.h))
}

func (this *QSinglePointEvent) ExclusivePointGrabber() *QObject {
	return newQObject(QSinglePointEvent_ExclusivePointGrabber(this.h))
}

func (this *QSinglePointEvent) SetExclusivePointGrabber(exclusiveGrabber *QObject) {
	QSinglePointEvent_SetExclusivePointGrabber(this.h, exclusiveGrabber.cPointer())
}

type QEnterEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQEnterEvent constructs a new QEnterEvent object.
func NewQEnterEvent(localPos *QPointF, scenePos *QPointF, globalPos *QPointF) *QEnterEvent {
	g := newQEnterEvent(QEnterEvent_new(localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer()))
	g.isSubclass = true
	return g
}

// NewQEnterEvent2 constructs a new QEnterEvent object.
func NewQEnterEvent2(localPos *QPointF, scenePos *QPointF, globalPos *QPointF, device *QPointingDevice) *QEnterEvent {
	g := newQEnterEvent(QEnterEvent_new2(localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), device.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QEnterEvent) Clone() *QEnterEvent {
	return newQEnterEvent(QEnterEvent_Clone(this.h))
}

func (this *QEnterEvent) Pos() *QPoint {
	_goptr := newQPoint(QEnterEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) GlobalPos() *QPoint {
	_goptr := newQPoint(QEnterEvent_GlobalPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) X() int {
	return (int)(QEnterEvent_X(this.h))
}

func (this *QEnterEvent) Y() int {
	return (int)(QEnterEvent_Y(this.h))
}

func (this *QEnterEvent) GlobalX() int {
	return (int)(QEnterEvent_GlobalX(this.h))
}

func (this *QEnterEvent) GlobalY() int {
	return (int)(QEnterEvent_GlobalY(this.h))
}

func (this *QEnterEvent) LocalPos() *QPointF {
	_goptr := newQPointF(QEnterEvent_LocalPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) WindowPos() *QPointF {
	_goptr := newQPointF(QEnterEvent_WindowPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QEnterEvent) ScreenPos() *QPointF {
	_goptr := newQPointF(QEnterEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QMouseEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQMouseEvent constructs a new QMouseEvent object.
func NewQMouseEvent(typeVal Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new(typeVal, localPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQMouseEvent2 constructs a new QMouseEvent object.
func NewQMouseEvent2(typeVal Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new2(typeVal, localPos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQMouseEvent3 constructs a new QMouseEvent object.
func NewQMouseEvent3(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new3(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQMouseEvent4 constructs a new QMouseEvent object.
func NewQMouseEvent4(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new4(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), (int)(source)))
	g.isSubclass = true
	return g
}

// NewQMouseEvent5 constructs a new QMouseEvent object.
func NewQMouseEvent5(typeVal Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new5(typeVal, localPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), device.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMouseEvent6 constructs a new QMouseEvent object.
func NewQMouseEvent6(typeVal Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new6(typeVal, localPos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), device.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMouseEvent7 constructs a new QMouseEvent object.
func NewQMouseEvent7(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new7(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), device.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMouseEvent8 constructs a new QMouseEvent object.
func NewQMouseEvent8(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource, device *QPointingDevice) *QMouseEvent {
	g := newQMouseEvent(QMouseEvent_new8(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), (int)(source), device.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QMouseEvent) Clone() *QMouseEvent {
	return newQMouseEvent(QMouseEvent_Clone(this.h))
}

func (this *QMouseEvent) Pos() *QPoint {
	_goptr := newQPoint(QMouseEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) GlobalPos() *QPoint {
	_goptr := newQPoint(QMouseEvent_GlobalPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) X() int {
	return (int)(QMouseEvent_X(this.h))
}

func (this *QMouseEvent) Y() int {
	return (int)(QMouseEvent_Y(this.h))
}

func (this *QMouseEvent) GlobalX() int {
	return (int)(QMouseEvent_GlobalX(this.h))
}

func (this *QMouseEvent) GlobalY() int {
	return (int)(QMouseEvent_GlobalY(this.h))
}

func (this *QMouseEvent) LocalPos() *QPointF {
	_goptr := newQPointF(QMouseEvent_LocalPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) WindowPos() *QPointF {
	_goptr := newQPointF(QMouseEvent_WindowPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) ScreenPos() *QPointF {
	_goptr := newQPointF(QMouseEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMouseEvent) Source() MouseEventSource {
	return (MouseEventSource)(QMouseEvent_Source(this.h))
}

func (this *QMouseEvent) Flags() MouseEventFlag {
	return (MouseEventFlag)(QMouseEvent_Flags(this.h))
}

type QHoverEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQHoverEvent constructs a new QHoverEvent object.
func NewQHoverEvent(typeVal Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF) *QHoverEvent {
	g := newQHoverEvent(QHoverEvent_new(typeVal, scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer()))
	g.isSubclass = true
	return g
}

// NewQHoverEvent2 constructs a new QHoverEvent object.
func NewQHoverEvent2(typeVal Type, pos *QPointF, oldPos *QPointF) *QHoverEvent {
	g := newQHoverEvent(QHoverEvent_new2(typeVal, pos.cPointer(), oldPos.cPointer()))
	g.isSubclass = true
	return g
}

// NewQHoverEvent3 constructs a new QHoverEvent object.
func NewQHoverEvent3(typeVal Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {
	g := newQHoverEvent(QHoverEvent_new3(typeVal, scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer(), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQHoverEvent4 constructs a new QHoverEvent object.
func NewQHoverEvent4(typeVal Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {
	g := newQHoverEvent(QHoverEvent_new4(typeVal, scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer(), (int)(modifiers), device.cPointer()))
	g.isSubclass = true
	return g
}

// NewQHoverEvent5 constructs a new QHoverEvent object.
func NewQHoverEvent5(typeVal Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {
	g := newQHoverEvent(QHoverEvent_new5(typeVal, pos.cPointer(), oldPos.cPointer(), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQHoverEvent6 constructs a new QHoverEvent object.
func NewQHoverEvent6(typeVal Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {
	g := newQHoverEvent(QHoverEvent_new6(typeVal, pos.cPointer(), oldPos.cPointer(), (int)(modifiers), device.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QHoverEvent) Clone() *QHoverEvent {
	return newQHoverEvent(QHoverEvent_Clone(this.h))
}

func (this *QHoverEvent) Pos() *QPoint {
	_goptr := newQPoint(QHoverEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHoverEvent) PosF() *QPointF {
	_goptr := newQPointF(QHoverEvent_PosF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHoverEvent) IsUpdateEvent() bool {
	return (bool)(QHoverEvent_IsUpdateEvent(this.h))
}

func (this *QHoverEvent) OldPos() *QPoint {
	_goptr := newQPoint(QHoverEvent_OldPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHoverEvent) OldPosF() *QPointF {
	_goptr := newQPointF(QHoverEvent_OldPosF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QWheelEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQWheelEvent constructs a new QWheelEvent object.
func NewQWheelEvent(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool) *QWheelEvent {
	g := newQWheelEvent(QWheelEvent_new(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (int)(buttons), (int)(modifiers), (int)(phase), (bool)(inverted)))
	g.isSubclass = true
	return g
}

// NewQWheelEvent2 constructs a new QWheelEvent object.
func NewQWheelEvent2(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource) *QWheelEvent {
	g := newQWheelEvent(QWheelEvent_new2(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (int)(buttons), (int)(modifiers), (int)(phase), (bool)(inverted), (int)(source)))
	g.isSubclass = true
	return g
}

// NewQWheelEvent3 constructs a new QWheelEvent object.
func NewQWheelEvent3(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource, device *QPointingDevice) *QWheelEvent {
	g := newQWheelEvent(QWheelEvent_new3(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (int)(buttons), (int)(modifiers), (int)(phase), (bool)(inverted), (int)(source), device.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QWheelEvent) Clone() *QWheelEvent {
	return newQWheelEvent(QWheelEvent_Clone(this.h))
}

func (this *QWheelEvent) PixelDelta() *QPoint {
	_goptr := newQPoint(QWheelEvent_PixelDelta(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWheelEvent) AngleDelta() *QPoint {
	_goptr := newQPoint(QWheelEvent_AngleDelta(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWheelEvent) Phase() ScrollPhase {
	return (ScrollPhase)(QWheelEvent_Phase(this.h))
}

func (this *QWheelEvent) Inverted() bool {
	return (bool)(QWheelEvent_Inverted(this.h))
}

func (this *QWheelEvent) IsInverted() bool {
	return (bool)(QWheelEvent_IsInverted(this.h))
}

func (this *QWheelEvent) HasPixelDelta() bool {
	return (bool)(QWheelEvent_HasPixelDelta(this.h))
}

func (this *QWheelEvent) IsBeginEvent() bool {
	return (bool)(QWheelEvent_IsBeginEvent(this.h))
}

func (this *QWheelEvent) IsUpdateEvent() bool {
	return (bool)(QWheelEvent_IsUpdateEvent(this.h))
}

func (this *QWheelEvent) IsEndEvent() bool {
	return (bool)(QWheelEvent_IsEndEvent(this.h))
}

func (this *QWheelEvent) Source() MouseEventSource {
	return (MouseEventSource)(QWheelEvent_Source(this.h))
}

type QTabletEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQTabletEvent constructs a new QTabletEvent object.
func NewQTabletEvent(t Type, device *QPointingDevice, pos *QPointF, globalPos *QPointF, pressure float64, xTilt float32, yTilt float32, tangentialPressure float32, rotation float64, z float32, keyState KeyboardModifier, button MouseButton, buttons MouseButton) *QTabletEvent {
	g := newQTabletEvent(QTabletEvent_new(t, device.cPointer(), pos.cPointer(), globalPos.cPointer(), (double)(pressure), (float)(xTilt), (float)(yTilt), (float)(tangentialPressure), (double)(rotation), (float)(z), (int)(keyState), (int)(button), (int)(buttons)))
	g.isSubclass = true
	return g
}

func (this *QTabletEvent) Clone() *QTabletEvent {
	return newQTabletEvent(QTabletEvent_Clone(this.h))
}

func (this *QTabletEvent) Pos() *QPoint {
	_goptr := newQPoint(QTabletEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) GlobalPos() *QPoint {
	_goptr := newQPoint(QTabletEvent_GlobalPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) PosF() *QPointF {
	_goptr := newQPointF(QTabletEvent_PosF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) GlobalPosF() *QPointF {
	_goptr := newQPointF(QTabletEvent_GlobalPosF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabletEvent) X() int {
	return (int)(QTabletEvent_X(this.h))
}

func (this *QTabletEvent) Y() int {
	return (int)(QTabletEvent_Y(this.h))
}

func (this *QTabletEvent) GlobalX() int {
	return (int)(QTabletEvent_GlobalX(this.h))
}

func (this *QTabletEvent) GlobalY() int {
	return (int)(QTabletEvent_GlobalY(this.h))
}

func (this *QTabletEvent) HiResGlobalX() float64 {
	return (float64)(QTabletEvent_HiResGlobalX(this.h))
}

func (this *QTabletEvent) HiResGlobalY() float64 {
	return (float64)(QTabletEvent_HiResGlobalY(this.h))
}

func (this *QTabletEvent) UniqueId() int64 {
	return (int64)(QTabletEvent_UniqueId(this.h))
}

func (this *QTabletEvent) Pressure() float64 {
	return (float64)(QTabletEvent_Pressure(this.h))
}

func (this *QTabletEvent) Rotation() float64 {
	return (float64)(QTabletEvent_Rotation(this.h))
}

func (this *QTabletEvent) Z() float64 {
	return (float64)(QTabletEvent_Z(this.h))
}

func (this *QTabletEvent) TangentialPressure() float64 {
	return (float64)(QTabletEvent_TangentialPressure(this.h))
}

func (this *QTabletEvent) XTilt() float64 {
	return (float64)(QTabletEvent_XTilt(this.h))
}

func (this *QTabletEvent) YTilt() float64 {
	return (float64)(QTabletEvent_YTilt(this.h))
}

type QNativeGestureEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQNativeGestureEvent constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent(typeVal NativeGestureType, dev *QPointingDevice, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, sequenceId uint64, intArgument uint64) *QNativeGestureEvent {
	g := newQNativeGestureEvent(QNativeGestureEvent_new((int)(typeVal), dev.cPointer(), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (double)(value), (ulonglong)(sequenceId), (ulonglong)(intArgument)))
	g.isSubclass = true
	return g
}

// NewQNativeGestureEvent2 constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent2(typeVal NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF) *QNativeGestureEvent {
	g := newQNativeGestureEvent(QNativeGestureEvent_new2((int)(typeVal), dev.cPointer(), (int)(fingerCount), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (double)(value), delta.cPointer()))
	g.isSubclass = true
	return g
}

// NewQNativeGestureEvent3 constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent3(typeVal NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF, sequenceId uint64) *QNativeGestureEvent {
	g := newQNativeGestureEvent(QNativeGestureEvent_new3((int)(typeVal), dev.cPointer(), (int)(fingerCount), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (double)(value), delta.cPointer(), (ulonglong)(sequenceId)))
	g.isSubclass = true
	return g
}

func (this *QNativeGestureEvent) Clone() *QNativeGestureEvent {
	return newQNativeGestureEvent(QNativeGestureEvent_Clone(this.h))
}

func (this *QNativeGestureEvent) GestureType() NativeGestureType {
	return (NativeGestureType)(QNativeGestureEvent_GestureType(this.h))
}

func (this *QNativeGestureEvent) FingerCount() int {
	return (int)(QNativeGestureEvent_FingerCount(this.h))
}

func (this *QNativeGestureEvent) Value() float64 {
	return (float64)(QNativeGestureEvent_Value(this.h))
}

func (this *QNativeGestureEvent) Delta() *QPointF {
	_goptr := newQPointF(QNativeGestureEvent_Delta(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) Pos() *QPoint {
	_goptr := newQPoint(QNativeGestureEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) GlobalPos() *QPoint {
	_goptr := newQPoint(QNativeGestureEvent_GlobalPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) LocalPos() *QPointF {
	_goptr := newQPointF(QNativeGestureEvent_LocalPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) WindowPos() *QPointF {
	_goptr := newQPointF(QNativeGestureEvent_WindowPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNativeGestureEvent) ScreenPos() *QPointF {
	_goptr := newQPointF(QNativeGestureEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QKeyEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQKeyEvent constructs a new QKeyEvent object.
func NewQKeyEvent(typeVal Type, key int, modifiers KeyboardModifier) *QKeyEvent {
	g := newQKeyEvent(QKeyEvent_new(typeVal, (int)(key), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQKeyEvent2 constructs a new QKeyEvent object.
func NewQKeyEvent2(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint) *QKeyEvent {
	g := newQKeyEvent(QKeyEvent_new2(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers)))
	g.isSubclass = true
	return g
}

// NewQKeyEvent3 constructs a new QKeyEvent object.
func NewQKeyEvent3(typeVal Type, key int, modifiers KeyboardModifier, text string) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQKeyEvent(QKeyEvent_new3(typeVal, (int)(key), (int)(modifiers), text_ms))
	g.isSubclass = true
	return g
}

// NewQKeyEvent4 constructs a new QKeyEvent object.
func NewQKeyEvent4(typeVal Type, key int, modifiers KeyboardModifier, text string, autorep bool) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQKeyEvent(QKeyEvent_new4(typeVal, (int)(key), (int)(modifiers), text_ms, (bool)(autorep)))
	g.isSubclass = true
	return g
}

// NewQKeyEvent5 constructs a new QKeyEvent object.
func NewQKeyEvent5(typeVal Type, key int, modifiers KeyboardModifier, text string, autorep bool, count uint16) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQKeyEvent(QKeyEvent_new5(typeVal, (int)(key), (int)(modifiers), text_ms, (bool)(autorep), (uint16_t)(count)))
	g.isSubclass = true
	return g
}

// NewQKeyEvent6 constructs a new QKeyEvent object.
func NewQKeyEvent6(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQKeyEvent(QKeyEvent_new6(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms))
	g.isSubclass = true
	return g
}

// NewQKeyEvent7 constructs a new QKeyEvent object.
func NewQKeyEvent7(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQKeyEvent(QKeyEvent_new7(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms, (bool)(autorep)))
	g.isSubclass = true
	return g
}

// NewQKeyEvent8 constructs a new QKeyEvent object.
func NewQKeyEvent8(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQKeyEvent(QKeyEvent_new8(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms, (bool)(autorep), (uint16_t)(count)))
	g.isSubclass = true
	return g
}

// NewQKeyEvent9 constructs a new QKeyEvent object.
func NewQKeyEvent9(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16, device *QInputDevice) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQKeyEvent(QKeyEvent_new9(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms, (bool)(autorep), (uint16_t)(count), device.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QKeyEvent) Clone() *QKeyEvent {
	return newQKeyEvent(QKeyEvent_Clone(this.h))
}

func (this *QKeyEvent) Key() int {
	return (int)(QKeyEvent_Key(this.h))
}

func (this *QKeyEvent) Matches(key QKeySequence__StandardKey) bool {
	return (bool)(QKeyEvent_Matches(this.h, (int)(key)))
}

func (this *QKeyEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QKeyEvent_Modifiers(this.h))
}

func (this *QKeyEvent) KeyCombination() *QKeyCombination {
	_goptr := newQKeyCombination(QKeyEvent_KeyCombination(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QKeyEvent) Text() string {
	var _ms struct_miqt_string = QKeyEvent_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QKeyEvent) IsAutoRepeat() bool {
	return (bool)(QKeyEvent_IsAutoRepeat(this.h))
}

func (this *QKeyEvent) Count() int {
	return (int)(QKeyEvent_Count(this.h))
}

func (this *QKeyEvent) NativeScanCode() uint {
	return (uint)(QKeyEvent_NativeScanCode(this.h))
}

func (this *QKeyEvent) NativeVirtualKey() uint {
	return (uint)(QKeyEvent_NativeVirtualKey(this.h))
}

func (this *QKeyEvent) NativeModifiers() uint {
	return (uint)(QKeyEvent_NativeModifiers(this.h))
}

type QFocusEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQFocusEvent constructs a new QFocusEvent object.
func NewQFocusEvent(typeVal Type) *QFocusEvent {
	g := newQFocusEvent(QFocusEvent_new(typeVal))
	g.isSubclass = true
	return g
}

// NewQFocusEvent2 constructs a new QFocusEvent object.
func NewQFocusEvent2(typeVal Type, reason FocusReason) *QFocusEvent {
	g := newQFocusEvent(QFocusEvent_new2(typeVal, (int)(reason)))
	g.isSubclass = true
	return g
}

func (this *QFocusEvent) Clone() *QFocusEvent {
	return newQFocusEvent(QFocusEvent_Clone(this.h))
}

func (this *QFocusEvent) GotFocus() bool {
	return (bool)(QFocusEvent_GotFocus(this.h))
}

func (this *QFocusEvent) LostFocus() bool {
	return (bool)(QFocusEvent_LostFocus(this.h))
}

func (this *QFocusEvent) Reason() FocusReason {
	return (FocusReason)(QFocusEvent_Reason(this.h))
}

type QPaintEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQPaintEvent constructs a new QPaintEvent object.
func NewQPaintEvent(paintRegion *QRegion) *QPaintEvent {
	g := newQPaintEvent(QPaintEvent_new(paintRegion.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPaintEvent2 constructs a new QPaintEvent object.
func NewQPaintEvent2(paintRect *QRect) *QPaintEvent {
	g := newQPaintEvent(QPaintEvent_new2(paintRect.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPaintEvent) Clone() *QPaintEvent {
	return newQPaintEvent(QPaintEvent_Clone(this.h))
}

func (this *QPaintEvent) Rect() *QRect {
	return newQRect(QPaintEvent_Rect(this.h))
}

func (this *QPaintEvent) Region() *QRegion {
	return newQRegion(QPaintEvent_Region(this.h))
}

type QMoveEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQMoveEvent constructs a new QMoveEvent object.
func NewQMoveEvent(pos *QPoint, oldPos *QPoint) *QMoveEvent {
	g := newQMoveEvent(QMoveEvent_new(pos.cPointer(), oldPos.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QMoveEvent) Clone() *QMoveEvent {
	return newQMoveEvent(QMoveEvent_Clone(this.h))
}

func (this *QMoveEvent) Pos() *QPoint {
	return newQPoint(QMoveEvent_Pos(this.h))
}

func (this *QMoveEvent) OldPos() *QPoint {
	return newQPoint(QMoveEvent_OldPos(this.h))
}

type QExposeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQExposeEvent constructs a new QExposeEvent object.
func NewQExposeEvent(m_region *QRegion) *QExposeEvent {
	g := newQExposeEvent(QExposeEvent_new(m_region.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QExposeEvent) Clone() *QExposeEvent {
	return newQExposeEvent(QExposeEvent_Clone(this.h))
}

func (this *QExposeEvent) Region() *QRegion {
	return newQRegion(QExposeEvent_Region(this.h))
}

type QPlatformSurfaceEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQPlatformSurfaceEvent constructs a new QPlatformSurfaceEvent object.
func NewQPlatformSurfaceEvent(surfaceEventType SurfaceEventType) *QPlatformSurfaceEvent {
	g := newQPlatformSurfaceEvent(QPlatformSurfaceEvent_new(surfaceEventType))
	g.isSubclass = true
	return g
}

func (this *QPlatformSurfaceEvent) Clone() *QPlatformSurfaceEvent {
	return newQPlatformSurfaceEvent(QPlatformSurfaceEvent_Clone(this.h))
}

func (this *QPlatformSurfaceEvent) SurfaceEventType() SurfaceEventType {
	xxxxxxxxx
}

type QResizeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQResizeEvent constructs a new QResizeEvent object.
func NewQResizeEvent(size *QSize, oldSize *QSize) *QResizeEvent {
	g := newQResizeEvent(QResizeEvent_new(size.cPointer(), oldSize.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QResizeEvent) Clone() *QResizeEvent {
	return newQResizeEvent(QResizeEvent_Clone(this.h))
}

func (this *QResizeEvent) Size() *QSize {
	return newQSize(QResizeEvent_Size(this.h))
}

func (this *QResizeEvent) OldSize() *QSize {
	return newQSize(QResizeEvent_OldSize(this.h))
}

type QCloseEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQCloseEvent constructs a new QCloseEvent object.
func NewQCloseEvent() *QCloseEvent {
	g := newQCloseEvent(QCloseEvent_new())
	g.isSubclass = true
	return g
}

func (this *QCloseEvent) Clone() *QCloseEvent {
	return newQCloseEvent(QCloseEvent_Clone(this.h))
}

type QIconDragEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQIconDragEvent constructs a new QIconDragEvent object.
func NewQIconDragEvent() *QIconDragEvent {
	g := newQIconDragEvent(QIconDragEvent_new())
	g.isSubclass = true
	return g
}

func (this *QIconDragEvent) Clone() *QIconDragEvent {
	return newQIconDragEvent(QIconDragEvent_Clone(this.h))
}

type QShowEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQShowEvent constructs a new QShowEvent object.
func NewQShowEvent() *QShowEvent {
	g := newQShowEvent(QShowEvent_new())
	g.isSubclass = true
	return g
}

func (this *QShowEvent) Clone() *QShowEvent {
	return newQShowEvent(QShowEvent_Clone(this.h))
}

type QHideEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQHideEvent constructs a new QHideEvent object.
func NewQHideEvent() *QHideEvent {
	g := newQHideEvent(QHideEvent_new())
	g.isSubclass = true
	return g
}

func (this *QHideEvent) Clone() *QHideEvent {
	return newQHideEvent(QHideEvent_Clone(this.h))
}

type QContextMenuEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQContextMenuEvent constructs a new QContextMenuEvent object.
func NewQContextMenuEvent(reason Reason, pos *QPoint, globalPos *QPoint) *QContextMenuEvent {
	g := newQContextMenuEvent(QContextMenuEvent_new(reason, pos.cPointer(), globalPos.cPointer()))
	g.isSubclass = true
	return g
}

// NewQContextMenuEvent2 constructs a new QContextMenuEvent object.
func NewQContextMenuEvent2(reason Reason, pos *QPoint) *QContextMenuEvent {
	g := newQContextMenuEvent(QContextMenuEvent_new2(reason, pos.cPointer()))
	g.isSubclass = true
	return g
}

// NewQContextMenuEvent3 constructs a new QContextMenuEvent object.
func NewQContextMenuEvent3(reason Reason, pos *QPoint, globalPos *QPoint, modifiers KeyboardModifier) *QContextMenuEvent {
	g := newQContextMenuEvent(QContextMenuEvent_new3(reason, pos.cPointer(), globalPos.cPointer(), (int)(modifiers)))
	g.isSubclass = true
	return g
}

func (this *QContextMenuEvent) Clone() *QContextMenuEvent {
	return newQContextMenuEvent(QContextMenuEvent_Clone(this.h))
}

func (this *QContextMenuEvent) X() int {
	return (int)(QContextMenuEvent_X(this.h))
}

func (this *QContextMenuEvent) Y() int {
	return (int)(QContextMenuEvent_Y(this.h))
}

func (this *QContextMenuEvent) GlobalX() int {
	return (int)(QContextMenuEvent_GlobalX(this.h))
}

func (this *QContextMenuEvent) GlobalY() int {
	return (int)(QContextMenuEvent_GlobalY(this.h))
}

func (this *QContextMenuEvent) Pos() *QPoint {
	return newQPoint(QContextMenuEvent_Pos(this.h))
}

func (this *QContextMenuEvent) GlobalPos() *QPoint {
	return newQPoint(QContextMenuEvent_GlobalPos(this.h))
}

func (this *QContextMenuEvent) Reason() Reason {
	xxxxxxxxx
}

type QInputMethodEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQInputMethodEvent constructs a new QInputMethodEvent object.
func NewQInputMethodEvent() *QInputMethodEvent {
	g := newQInputMethodEvent(QInputMethodEvent_new())
	g.isSubclass = true
	return g
}

// NewQInputMethodEvent2 constructs a new QInputMethodEvent object.
func NewQInputMethodEvent2(preeditText string, attributes []Attribute) *QInputMethodEvent {
	preeditText_ms := struct_miqt_string{}
	preeditText_ms.data = CString(preeditText)
	preeditText_ms.len = size_t(len(preeditText))
	defer free(unsafe.Pointer(preeditText_ms.data))
	attributes_CArray := (*[0xffff]Attribute)(malloc(size_t(8 * len(attributes))))
	defer free(unsafe.Pointer(attributes_CArray))
	for i := range attributes {
		attributes_CArray[i] = attributes[i]
	}
	attributes_ma := struct_miqt_array{len: size_t(len(attributes)), data: unsafe.Pointer(attributes_CArray)}
	g := newQInputMethodEvent(QInputMethodEvent_new2(preeditText_ms, attributes_ma))
	g.isSubclass = true
	return g
}

func (this *QInputMethodEvent) Clone() *QInputMethodEvent {
	return newQInputMethodEvent(QInputMethodEvent_Clone(this.h))
}

func (this *QInputMethodEvent) SetCommitString(commitString string) {
	commitString_ms := struct_miqt_string{}
	commitString_ms.data = CString(commitString)
	commitString_ms.len = size_t(len(commitString))
	defer free(unsafe.Pointer(commitString_ms.data))
	QInputMethodEvent_SetCommitString(this.h, commitString_ms)
}

func (this *QInputMethodEvent) Attributes() []Attribute {
	var _ma struct_miqt_array = QInputMethodEvent_Attributes(this.h)
	_ret := make([]Attribute, int(_ma.len))
	_outCast := (*[0xffff]Attribute)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QInputMethodEvent) PreeditString() string {
	var _ms struct_miqt_string = QInputMethodEvent_PreeditString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputMethodEvent) CommitString() string {
	var _ms struct_miqt_string = QInputMethodEvent_CommitString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputMethodEvent) ReplacementStart() int {
	return (int)(QInputMethodEvent_ReplacementStart(this.h))
}

func (this *QInputMethodEvent) ReplacementLength() int {
	return (int)(QInputMethodEvent_ReplacementLength(this.h))
}

func (this *QInputMethodEvent) SetCommitString2(commitString string, replaceFrom int) {
	commitString_ms := struct_miqt_string{}
	commitString_ms.data = CString(commitString)
	commitString_ms.len = size_t(len(commitString))
	defer free(unsafe.Pointer(commitString_ms.data))
	QInputMethodEvent_SetCommitString2(this.h, commitString_ms, (int)(replaceFrom))
}

func (this *QInputMethodEvent) SetCommitString3(commitString string, replaceFrom int, replaceLength int) {
	commitString_ms := struct_miqt_string{}
	commitString_ms.data = CString(commitString)
	commitString_ms.len = size_t(len(commitString))
	defer free(unsafe.Pointer(commitString_ms.data))
	QInputMethodEvent_SetCommitString3(this.h, commitString_ms, (int)(replaceFrom), (int)(replaceLength))
}

type QInputMethodQueryEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQInputMethodQueryEvent constructs a new QInputMethodQueryEvent object.
func NewQInputMethodQueryEvent(queries InputMethodQuery) *QInputMethodQueryEvent {
	g := newQInputMethodQueryEvent(QInputMethodQueryEvent_new((int)(queries)))
	g.isSubclass = true
	return g
}

func (this *QInputMethodQueryEvent) Clone() *QInputMethodQueryEvent {
	return newQInputMethodQueryEvent(QInputMethodQueryEvent_Clone(this.h))
}

func (this *QInputMethodQueryEvent) Queries() InputMethodQuery {
	return (InputMethodQuery)(QInputMethodQueryEvent_Queries(this.h))
}

func (this *QInputMethodQueryEvent) SetValue(query InputMethodQuery, value *QVariant) {
	QInputMethodQueryEvent_SetValue(this.h, (int)(query), value.cPointer())
}

func (this *QInputMethodQueryEvent) Value(query InputMethodQuery) *QVariant {
	_goptr := newQVariant(QInputMethodQueryEvent_Value(this.h, (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QDropEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDropEvent constructs a new QDropEvent object.
func NewQDropEvent(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDropEvent {
	g := newQDropEvent(QDropEvent_new(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQDropEvent2 constructs a new QDropEvent object.
func NewQDropEvent2(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, typeVal Type) *QDropEvent {
	g := newQDropEvent(QDropEvent_new2(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers), typeVal))
	g.isSubclass = true
	return g
}

func (this *QDropEvent) Clone() *QDropEvent {
	return newQDropEvent(QDropEvent_Clone(this.h))
}

func (this *QDropEvent) Pos() *QPoint {
	_goptr := newQPoint(QDropEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDropEvent) PosF() *QPointF {
	_goptr := newQPointF(QDropEvent_PosF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDropEvent) MouseButtons() MouseButton {
	return (MouseButton)(QDropEvent_MouseButtons(this.h))
}

func (this *QDropEvent) KeyboardModifiers() KeyboardModifier {
	return (KeyboardModifier)(QDropEvent_KeyboardModifiers(this.h))
}

func (this *QDropEvent) Position() *QPointF {
	_goptr := newQPointF(QDropEvent_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDropEvent) Buttons() MouseButton {
	return (MouseButton)(QDropEvent_Buttons(this.h))
}

func (this *QDropEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QDropEvent_Modifiers(this.h))
}

func (this *QDropEvent) PossibleActions() DropAction {
	return (DropAction)(QDropEvent_PossibleActions(this.h))
}

func (this *QDropEvent) ProposedAction() DropAction {
	return (DropAction)(QDropEvent_ProposedAction(this.h))
}

func (this *QDropEvent) AcceptProposedAction() {
	QDropEvent_AcceptProposedAction(this.h)
}

func (this *QDropEvent) DropAction() DropAction {
	return (DropAction)(QDropEvent_DropAction(this.h))
}

func (this *QDropEvent) SetDropAction(action DropAction) {
	QDropEvent_SetDropAction(this.h, (int)(action))
}

func (this *QDropEvent) Source() *QObject {
	return newQObject(QDropEvent_Source(this.h))
}

func (this *QDropEvent) MimeData() *QMimeData {
	return newQMimeData(QDropEvent_MimeData(this.h))
}

type QDragMoveEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDragMoveEvent constructs a new QDragMoveEvent object.
func NewQDragMoveEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragMoveEvent {
	g := newQDragMoveEvent(QDragMoveEvent_new(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQDragMoveEvent2 constructs a new QDragMoveEvent object.
func NewQDragMoveEvent2(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, typeVal Type) *QDragMoveEvent {
	g := newQDragMoveEvent(QDragMoveEvent_new2(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers), typeVal))
	g.isSubclass = true
	return g
}

func (this *QDragMoveEvent) Clone() *QDragMoveEvent {
	return newQDragMoveEvent(QDragMoveEvent_Clone(this.h))
}

func (this *QDragMoveEvent) AnswerRect() *QRect {
	_goptr := newQRect(QDragMoveEvent_AnswerRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDragMoveEvent) Accept() {
	QDragMoveEvent_Accept(this.h)
}

func (this *QDragMoveEvent) Ignore() {
	QDragMoveEvent_Ignore(this.h)
}

func (this *QDragMoveEvent) AcceptWithQRect(r *QRect) {
	QDragMoveEvent_AcceptWithQRect(this.h, r.cPointer())
}

func (this *QDragMoveEvent) IgnoreWithQRect(r *QRect) {
	QDragMoveEvent_IgnoreWithQRect(this.h, r.cPointer())
}

type QDragEnterEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDragEnterEvent constructs a new QDragEnterEvent object.
func NewQDragEnterEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragEnterEvent {
	g := newQDragEnterEvent(QDragEnterEvent_new(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers)))
	g.isSubclass = true
	return g
}

func (this *QDragEnterEvent) Clone() *QDragEnterEvent {
	return newQDragEnterEvent(QDragEnterEvent_Clone(this.h))
}

type QDragLeaveEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDragLeaveEvent constructs a new QDragLeaveEvent object.
func NewQDragLeaveEvent() *QDragLeaveEvent {
	g := newQDragLeaveEvent(QDragLeaveEvent_new())
	g.isSubclass = true
	return g
}

func (this *QDragLeaveEvent) Clone() *QDragLeaveEvent {
	return newQDragLeaveEvent(QDragLeaveEvent_Clone(this.h))
}

type QHelpEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQHelpEvent constructs a new QHelpEvent object.
func NewQHelpEvent(typeVal Type, pos *QPoint, globalPos *QPoint) *QHelpEvent {
	g := newQHelpEvent(QHelpEvent_new(typeVal, pos.cPointer(), globalPos.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QHelpEvent) Clone() *QHelpEvent {
	return newQHelpEvent(QHelpEvent_Clone(this.h))
}

func (this *QHelpEvent) X() int {
	return (int)(QHelpEvent_X(this.h))
}

func (this *QHelpEvent) Y() int {
	return (int)(QHelpEvent_Y(this.h))
}

func (this *QHelpEvent) GlobalX() int {
	return (int)(QHelpEvent_GlobalX(this.h))
}

func (this *QHelpEvent) GlobalY() int {
	return (int)(QHelpEvent_GlobalY(this.h))
}

func (this *QHelpEvent) Pos() *QPoint {
	return newQPoint(QHelpEvent_Pos(this.h))
}

func (this *QHelpEvent) GlobalPos() *QPoint {
	return newQPoint(QHelpEvent_GlobalPos(this.h))
}

type QStatusTipEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQStatusTipEvent constructs a new QStatusTipEvent object.
func NewQStatusTipEvent(tip string) *QStatusTipEvent {
	tip_ms := struct_miqt_string{}
	tip_ms.data = CString(tip)
	tip_ms.len = size_t(len(tip))
	defer free(unsafe.Pointer(tip_ms.data))
	g := newQStatusTipEvent(QStatusTipEvent_new(tip_ms))
	g.isSubclass = true
	return g
}

func (this *QStatusTipEvent) Clone() *QStatusTipEvent {
	return newQStatusTipEvent(QStatusTipEvent_Clone(this.h))
}

func (this *QStatusTipEvent) Tip() string {
	var _ms struct_miqt_string = QStatusTipEvent_Tip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QWhatsThisClickedEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQWhatsThisClickedEvent constructs a new QWhatsThisClickedEvent object.
func NewQWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	href_ms := struct_miqt_string{}
	href_ms.data = CString(href)
	href_ms.len = size_t(len(href))
	defer free(unsafe.Pointer(href_ms.data))
	g := newQWhatsThisClickedEvent(QWhatsThisClickedEvent_new(href_ms))
	g.isSubclass = true
	return g
}

func (this *QWhatsThisClickedEvent) Clone() *QWhatsThisClickedEvent {
	return newQWhatsThisClickedEvent(QWhatsThisClickedEvent_Clone(this.h))
}

func (this *QWhatsThisClickedEvent) Href() string {
	var _ms struct_miqt_string = QWhatsThisClickedEvent_Href(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QActionEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQActionEvent constructs a new QActionEvent object.
func NewQActionEvent(typeVal int, action *QAction) *QActionEvent {
	g := newQActionEvent(QActionEvent_new((int)(typeVal), action.cPointer()))
	g.isSubclass = true
	return g
}

// NewQActionEvent2 constructs a new QActionEvent object.
func NewQActionEvent2(typeVal int, action *QAction, before *QAction) *QActionEvent {
	g := newQActionEvent(QActionEvent_new2((int)(typeVal), action.cPointer(), before.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QActionEvent) Clone() *QActionEvent {
	return newQActionEvent(QActionEvent_Clone(this.h))
}

func (this *QActionEvent) Action() *QAction {
	return newQAction(QActionEvent_Action(this.h))
}

func (this *QActionEvent) Before() *QAction {
	return newQAction(QActionEvent_Before(this.h))
}

type QFileOpenEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQFileOpenEvent constructs a new QFileOpenEvent object.
func NewQFileOpenEvent(file string) *QFileOpenEvent {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	g := newQFileOpenEvent(QFileOpenEvent_new(file_ms))
	g.isSubclass = true
	return g
}

// NewQFileOpenEvent2 constructs a new QFileOpenEvent object.
func NewQFileOpenEvent2(url *QUrl) *QFileOpenEvent {
	g := newQFileOpenEvent(QFileOpenEvent_new2(url.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QFileOpenEvent) Clone() *QFileOpenEvent {
	return newQFileOpenEvent(QFileOpenEvent_Clone(this.h))
}

func (this *QFileOpenEvent) File() string {
	var _ms struct_miqt_string = QFileOpenEvent_File(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileOpenEvent) Url() *QUrl {
	_goptr := newQUrl(QFileOpenEvent_Url(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileOpenEvent) OpenFile(file *QFile, flags OpenModeFlag) bool {
	return (bool)(QFileOpenEvent_OpenFile(this.h, file.cPointer(), (int)(flags)))
}

type QToolBarChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQToolBarChangeEvent constructs a new QToolBarChangeEvent object.
func NewQToolBarChangeEvent(t bool) *QToolBarChangeEvent {
	g := newQToolBarChangeEvent(QToolBarChangeEvent_new((bool)(t)))
	g.isSubclass = true
	return g
}

func (this *QToolBarChangeEvent) Clone() *QToolBarChangeEvent {
	return newQToolBarChangeEvent(QToolBarChangeEvent_Clone(this.h))
}

func (this *QToolBarChangeEvent) Toggle() bool {
	return (bool)(QToolBarChangeEvent_Toggle(this.h))
}

type QShortcutEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQShortcutEvent constructs a new QShortcutEvent object.
func NewQShortcutEvent(key *QKeySequence, id int) *QShortcutEvent {
	g := newQShortcutEvent(QShortcutEvent_new(key.cPointer(), (int)(id)))
	g.isSubclass = true
	return g
}

// NewQShortcutEvent2 constructs a new QShortcutEvent object.
func NewQShortcutEvent2(key *QKeySequence) *QShortcutEvent {
	g := newQShortcutEvent(QShortcutEvent_new2(key.cPointer()))
	g.isSubclass = true
	return g
}

// NewQShortcutEvent3 constructs a new QShortcutEvent object.
func NewQShortcutEvent3(key *QKeySequence, id int, ambiguous bool) *QShortcutEvent {
	g := newQShortcutEvent(QShortcutEvent_new3(key.cPointer(), (int)(id), (bool)(ambiguous)))
	g.isSubclass = true
	return g
}

// NewQShortcutEvent4 constructs a new QShortcutEvent object.
func NewQShortcutEvent4(key *QKeySequence, shortcut *QShortcut) *QShortcutEvent {
	g := newQShortcutEvent(QShortcutEvent_new4(key.cPointer(), shortcut.cPointer()))
	g.isSubclass = true
	return g
}

// NewQShortcutEvent5 constructs a new QShortcutEvent object.
func NewQShortcutEvent5(key *QKeySequence, shortcut *QShortcut, ambiguous bool) *QShortcutEvent {
	g := newQShortcutEvent(QShortcutEvent_new5(key.cPointer(), shortcut.cPointer(), (bool)(ambiguous)))
	g.isSubclass = true
	return g
}

func (this *QShortcutEvent) Clone() *QShortcutEvent {
	return newQShortcutEvent(QShortcutEvent_Clone(this.h))
}

func (this *QShortcutEvent) Key() *QKeySequence {
	return newQKeySequence(QShortcutEvent_Key(this.h))
}

func (this *QShortcutEvent) ShortcutId() int {
	return (int)(QShortcutEvent_ShortcutId(this.h))
}

func (this *QShortcutEvent) IsAmbiguous() bool {
	return (bool)(QShortcutEvent_IsAmbiguous(this.h))
}

type QWindowStateChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQWindowStateChangeEvent constructs a new QWindowStateChangeEvent object.
func NewQWindowStateChangeEvent(oldState WindowState) *QWindowStateChangeEvent {
	g := newQWindowStateChangeEvent(QWindowStateChangeEvent_new((int)(oldState)))
	g.isSubclass = true
	return g
}

// NewQWindowStateChangeEvent2 constructs a new QWindowStateChangeEvent object.
func NewQWindowStateChangeEvent2(oldState WindowState, isOverride bool) *QWindowStateChangeEvent {
	g := newQWindowStateChangeEvent(QWindowStateChangeEvent_new2((int)(oldState), (bool)(isOverride)))
	g.isSubclass = true
	return g
}

func (this *QWindowStateChangeEvent) Clone() *QWindowStateChangeEvent {
	return newQWindowStateChangeEvent(QWindowStateChangeEvent_Clone(this.h))
}

func (this *QWindowStateChangeEvent) OldState() WindowState {
	return (WindowState)(QWindowStateChangeEvent_OldState(this.h))
}

func (this *QWindowStateChangeEvent) IsOverride() bool {
	return (bool)(QWindowStateChangeEvent_IsOverride(this.h))
}

type QTouchEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQTouchEvent constructs a new QTouchEvent object.
func NewQTouchEvent(eventType QEvent__Type) *QTouchEvent {
	g := newQTouchEvent(QTouchEvent_new((int)(eventType)))
	g.isSubclass = true
	return g
}

// NewQTouchEvent2 constructs a new QTouchEvent object.
func NewQTouchEvent2(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates State) *QTouchEvent {
	g := newQTouchEvent(QTouchEvent_new2((int)(eventType), device.cPointer(), (int)(modifiers), (int)(touchPointStates)))
	g.isSubclass = true
	return g
}

// NewQTouchEvent3 constructs a new QTouchEvent object.
func NewQTouchEvent3(eventType QEvent__Type, device *QPointingDevice) *QTouchEvent {
	g := newQTouchEvent(QTouchEvent_new3((int)(eventType), device.cPointer()))
	g.isSubclass = true
	return g
}

// NewQTouchEvent4 constructs a new QTouchEvent object.
func NewQTouchEvent4(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier) *QTouchEvent {
	g := newQTouchEvent(QTouchEvent_new4((int)(eventType), device.cPointer(), (int)(modifiers)))
	g.isSubclass = true
	return g
}

// NewQTouchEvent5 constructs a new QTouchEvent object.
func NewQTouchEvent5(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPoints []QEventPoint) *QTouchEvent {
	touchPoints_CArray := (*[0xffff]*QEventPoint)(malloc(size_t(8 * len(touchPoints))))
	defer free(unsafe.Pointer(touchPoints_CArray))
	for i := range touchPoints {
		touchPoints_CArray[i] = touchPoints[i].cPointer()
	}
	touchPoints_ma := struct_miqt_array{len: size_t(len(touchPoints)), data: unsafe.Pointer(touchPoints_CArray)}
	g := newQTouchEvent(QTouchEvent_new5((int)(eventType), device.cPointer(), (int)(modifiers), touchPoints_ma))
	g.isSubclass = true
	return g
}

// NewQTouchEvent6 constructs a new QTouchEvent object.
func NewQTouchEvent6(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates State, touchPoints []QEventPoint) *QTouchEvent {
	touchPoints_CArray := (*[0xffff]*QEventPoint)(malloc(size_t(8 * len(touchPoints))))
	defer free(unsafe.Pointer(touchPoints_CArray))
	for i := range touchPoints {
		touchPoints_CArray[i] = touchPoints[i].cPointer()
	}
	touchPoints_ma := struct_miqt_array{len: size_t(len(touchPoints)), data: unsafe.Pointer(touchPoints_CArray)}
	g := newQTouchEvent(QTouchEvent_new6((int)(eventType), device.cPointer(), (int)(modifiers), (int)(touchPointStates), touchPoints_ma))
	g.isSubclass = true
	return g
}

func (this *QTouchEvent) Clone() *QTouchEvent {
	return newQTouchEvent(QTouchEvent_Clone(this.h))
}

func (this *QTouchEvent) Target() *QObject {
	return newQObject(QTouchEvent_Target(this.h))
}

func (this *QTouchEvent) TouchPointStates() State {
	return (State)(QTouchEvent_TouchPointStates(this.h))
}

func (this *QTouchEvent) TouchPoints() []QEventPoint {
	var _ma struct_miqt_array = QTouchEvent_TouchPoints(this.h)
	_ret := make([]QEventPoint, int(_ma.len))
	_outCast := (*[0xffff]*QEventPoint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQEventPoint(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTouchEvent) IsBeginEvent() bool {
	return (bool)(QTouchEvent_IsBeginEvent(this.h))
}

func (this *QTouchEvent) IsUpdateEvent() bool {
	return (bool)(QTouchEvent_IsUpdateEvent(this.h))
}

func (this *QTouchEvent) IsEndEvent() bool {
	return (bool)(QTouchEvent_IsEndEvent(this.h))
}

type QScrollPrepareEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQScrollPrepareEvent constructs a new QScrollPrepareEvent object.
func NewQScrollPrepareEvent(startPos *QPointF) *QScrollPrepareEvent {
	g := newQScrollPrepareEvent(QScrollPrepareEvent_new(startPos.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QScrollPrepareEvent) Clone() *QScrollPrepareEvent {
	return newQScrollPrepareEvent(QScrollPrepareEvent_Clone(this.h))
}

func (this *QScrollPrepareEvent) StartPos() *QPointF {
	_goptr := newQPointF(QScrollPrepareEvent_StartPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) ViewportSize() *QSizeF {
	_goptr := newQSizeF(QScrollPrepareEvent_ViewportSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) ContentPosRange() *QRectF {
	_goptr := newQRectF(QScrollPrepareEvent_ContentPosRange(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) ContentPos() *QPointF {
	_goptr := newQPointF(QScrollPrepareEvent_ContentPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollPrepareEvent) SetViewportSize(size *QSizeF) {
	QScrollPrepareEvent_SetViewportSize(this.h, size.cPointer())
}

func (this *QScrollPrepareEvent) SetContentPosRange(rect *QRectF) {
	QScrollPrepareEvent_SetContentPosRange(this.h, rect.cPointer())
}

func (this *QScrollPrepareEvent) SetContentPos(pos *QPointF) {
	QScrollPrepareEvent_SetContentPos(this.h, pos.cPointer())
}

type QScrollEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQScrollEvent constructs a new QScrollEvent object.
func NewQScrollEvent(contentPos *QPointF, overshoot *QPointF, scrollState ScrollState) *QScrollEvent {
	g := newQScrollEvent(QScrollEvent_new(contentPos.cPointer(), overshoot.cPointer(), scrollState))
	g.isSubclass = true
	return g
}

func (this *QScrollEvent) Clone() *QScrollEvent {
	return newQScrollEvent(QScrollEvent_Clone(this.h))
}

func (this *QScrollEvent) ContentPos() *QPointF {
	_goptr := newQPointF(QScrollEvent_ContentPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollEvent) OvershootDistance() *QPointF {
	_goptr := newQPointF(QScrollEvent_OvershootDistance(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollEvent) ScrollState() ScrollState {
	xxxxxxxxx
}

type QScreenOrientationChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQScreenOrientationChangeEvent constructs a new QScreenOrientationChangeEvent object.
func NewQScreenOrientationChangeEvent(screen *QScreen, orientation ScreenOrientation) *QScreenOrientationChangeEvent {
	g := newQScreenOrientationChangeEvent(QScreenOrientationChangeEvent_new(screen.cPointer(), (int)(orientation)))
	g.isSubclass = true
	return g
}

func (this *QScreenOrientationChangeEvent) Clone() *QScreenOrientationChangeEvent {
	return newQScreenOrientationChangeEvent(QScreenOrientationChangeEvent_Clone(this.h))
}

func (this *QScreenOrientationChangeEvent) Screen() *QScreen {
	return newQScreen(QScreenOrientationChangeEvent_Screen(this.h))
}

func (this *QScreenOrientationChangeEvent) Orientation() ScreenOrientation {
	return (ScreenOrientation)(QScreenOrientationChangeEvent_Orientation(this.h))
}

type QApplicationStateChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQApplicationStateChangeEvent constructs a new QApplicationStateChangeEvent object.
func NewQApplicationStateChangeEvent(state ApplicationState) *QApplicationStateChangeEvent {
	g := newQApplicationStateChangeEvent(QApplicationStateChangeEvent_new((int)(state)))
	g.isSubclass = true
	return g
}

func (this *QApplicationStateChangeEvent) Clone() *QApplicationStateChangeEvent {
	return newQApplicationStateChangeEvent(QApplicationStateChangeEvent_Clone(this.h))
}

func (this *QApplicationStateChangeEvent) ApplicationState() ApplicationState {
	return (ApplicationState)(QApplicationStateChangeEvent_ApplicationState(this.h))
}

type QChildWindowEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQChildWindowEvent constructs a new QChildWindowEvent object.
func NewQChildWindowEvent(typeVal Type, childWindow *QWindow) *QChildWindowEvent {
	g := newQChildWindowEvent(QChildWindowEvent_new(typeVal, childWindow.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QChildWindowEvent) Clone() *QChildWindowEvent {
	return newQChildWindowEvent(QChildWindowEvent_Clone(this.h))
}

func (this *QChildWindowEvent) Child() *QWindow {
	return newQWindow(QChildWindowEvent_Child(this.h))
}

type QInputMethodEvent__Attribute struct {
	h          uintptr
	isSubclass bool
}

// NewQInputMethodEvent__Attribute constructs a new QInputMethodEvent::Attribute object.
func NewQInputMethodEvent__Attribute(typ AttributeType, s int, l int, val QVariant) *QInputMethodEvent__Attribute {
	g := newQInputMethodEvent__Attribute(QInputMethodEvent__Attribute_new(typ, (int)(s), (int)(l), val.cPointer()))
	g.isSubclass = true
	return g
}

// NewQInputMethodEvent__Attribute2 constructs a new QInputMethodEvent::Attribute object.
func NewQInputMethodEvent__Attribute2(typ AttributeType, s int, l int) *QInputMethodEvent__Attribute {
	g := newQInputMethodEvent__Attribute(QInputMethodEvent__Attribute_new2(typ, (int)(s), (int)(l)))
	g.isSubclass = true
	return g
}

func (this *QInputMethodEvent__Attribute) OperatorAssign(param1 *Attribute) {
	QInputMethodEvent__Attribute_OperatorAssign(this.h, param1)
}
