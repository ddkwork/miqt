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

	ret := newQInputEvent(QInputEvent_new(typeVal, m_dev.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQInputEvent2 constructs a new QInputEvent object.
func NewQInputEvent2(typeVal Type, m_dev *QInputDevice, modifiers KeyboardModifier) *QInputEvent {

	ret := newQInputEvent(QInputEvent_new2(typeVal, m_dev.cPointer(), (int)(modifiers)))
	ret.isSubclass = true
	return ret
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

func (this *QInputEvent) callVirtualBase_Clone() *QInputEvent {

	return newQInputEvent(QInputEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QInputEvent) OnClone(slot func(super func() *QInputEvent) *QInputEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputEvent_Clone
func miqt_exec_callback_QInputEvent_Clone(self QInputEvent, cb intptr_t) *QInputEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QInputEvent) *QInputEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QInputEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	QInputEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (ulonglong)(timestamp))

}
func (this *QInputEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputEvent_SetTimestamp
func miqt_exec_callback_QInputEvent_SetTimestamp(self QInputEvent, cb intptr_t, timestamp ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QInputEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

func (this *QInputEvent) callVirtualBase_SetAccepted(accepted bool) {

	QInputEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QInputEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputEvent_SetAccepted
func miqt_exec_callback_QInputEvent_SetAccepted(self QInputEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QInputEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QPointerEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQPointerEvent constructs a new QPointerEvent object.
func NewQPointerEvent(typeVal Type, dev *QPointingDevice) *QPointerEvent {

	ret := newQPointerEvent(QPointerEvent_new(typeVal, dev.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPointerEvent2 constructs a new QPointerEvent object.
func NewQPointerEvent2(typeVal Type, dev *QPointingDevice, modifiers KeyboardModifier) *QPointerEvent {

	ret := newQPointerEvent(QPointerEvent_new2(typeVal, dev.cPointer(), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQPointerEvent3 constructs a new QPointerEvent object.
func NewQPointerEvent3(typeVal Type, dev *QPointingDevice, modifiers KeyboardModifier, points []QEventPoint) *QPointerEvent {
	points_CArray := (*[0xffff]*QEventPoint)(malloc(size_t(8 * len(points))))
	defer free(unsafe.Pointer(points_CArray))
	for i := range points {
		points_CArray[i] = points[i].cPointer()
	}
	points_ma := struct_miqt_array{len: size_t(len(points)), data: unsafe.Pointer(points_CArray)}

	ret := newQPointerEvent(QPointerEvent_new3(typeVal, dev.cPointer(), (int)(modifiers), points_ma))
	ret.isSubclass = true
	return ret
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

func (this *QPointerEvent) callVirtualBase_Clone() *QPointerEvent {

	return newQPointerEvent(QPointerEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QPointerEvent) OnClone(slot func(super func() *QPointerEvent) *QPointerEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointerEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_Clone
func miqt_exec_callback_QPointerEvent_Clone(self QPointerEvent, cb intptr_t) *QPointerEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPointerEvent) *QPointerEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QPointerEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	QPointerEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (ulonglong)(timestamp))

}
func (this *QPointerEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointerEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_SetTimestamp
func miqt_exec_callback_QPointerEvent_SetTimestamp(self QPointerEvent, cb intptr_t, timestamp ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QPointerEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

func (this *QPointerEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QPointerEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QPointerEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointerEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_IsBeginEvent
func miqt_exec_callback_QPointerEvent_IsBeginEvent(self QPointerEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QPointerEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QPointerEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QPointerEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointerEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_IsUpdateEvent
func miqt_exec_callback_QPointerEvent_IsUpdateEvent(self QPointerEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QPointerEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QPointerEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QPointerEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointerEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_IsEndEvent
func miqt_exec_callback_QPointerEvent_IsEndEvent(self QPointerEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointerEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

func (this *QPointerEvent) callVirtualBase_SetAccepted(accepted bool) {

	QPointerEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QPointerEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointerEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointerEvent_SetAccepted
func miqt_exec_callback_QPointerEvent_SetAccepted(self QPointerEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QPointerEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

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

	ret := newQEnterEvent(QEnterEvent_new(localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQEnterEvent2 constructs a new QEnterEvent object.
func NewQEnterEvent2(localPos *QPointF, scenePos *QPointF, globalPos *QPointF, device *QPointingDevice) *QEnterEvent {

	ret := newQEnterEvent(QEnterEvent_new2(localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), device.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QEnterEvent) callVirtualBase_Clone() *QEnterEvent {

	return newQEnterEvent(QEnterEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QEnterEvent) OnClone(slot func(super func() *QEnterEvent) *QEnterEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEnterEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_Clone
func miqt_exec_callback_QEnterEvent_Clone(self QEnterEvent, cb intptr_t) *QEnterEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QEnterEvent) *QEnterEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QEnterEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QEnterEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QEnterEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEnterEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_IsBeginEvent
func miqt_exec_callback_QEnterEvent_IsBeginEvent(self QEnterEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QEnterEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QEnterEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QEnterEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEnterEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_IsUpdateEvent
func miqt_exec_callback_QEnterEvent_IsUpdateEvent(self QEnterEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QEnterEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QEnterEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QEnterEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEnterEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEnterEvent_IsEndEvent
func miqt_exec_callback_QEnterEvent_IsEndEvent(self QEnterEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEnterEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

type QMouseEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQMouseEvent constructs a new QMouseEvent object.
func NewQMouseEvent(typeVal Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new(typeVal, localPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent2 constructs a new QMouseEvent object.
func NewQMouseEvent2(typeVal Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new2(typeVal, localPos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent3 constructs a new QMouseEvent object.
func NewQMouseEvent3(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new3(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent4 constructs a new QMouseEvent object.
func NewQMouseEvent4(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new4(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), (int)(source)))
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent5 constructs a new QMouseEvent object.
func NewQMouseEvent5(typeVal Type, localPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new5(typeVal, localPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent6 constructs a new QMouseEvent object.
func NewQMouseEvent6(typeVal Type, localPos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new6(typeVal, localPos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent7 constructs a new QMouseEvent object.
func NewQMouseEvent7(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, device *QPointingDevice) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new7(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMouseEvent8 constructs a new QMouseEvent object.
func NewQMouseEvent8(typeVal Type, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, button MouseButton, buttons MouseButton, modifiers KeyboardModifier, source MouseEventSource, device *QPointingDevice) *QMouseEvent {

	ret := newQMouseEvent(QMouseEvent_new8(typeVal, localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (int)(button), (int)(buttons), (int)(modifiers), (int)(source), device.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QMouseEvent) callVirtualBase_Clone() *QMouseEvent {

	return newQMouseEvent(QMouseEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QMouseEvent) OnClone(slot func(super func() *QMouseEvent) *QMouseEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMouseEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_Clone
func miqt_exec_callback_QMouseEvent_Clone(self QMouseEvent, cb intptr_t) *QMouseEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMouseEvent) *QMouseEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QMouseEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QMouseEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QMouseEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMouseEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_IsBeginEvent
func miqt_exec_callback_QMouseEvent_IsBeginEvent(self QMouseEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QMouseEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QMouseEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QMouseEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMouseEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_IsUpdateEvent
func miqt_exec_callback_QMouseEvent_IsUpdateEvent(self QMouseEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QMouseEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QMouseEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QMouseEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMouseEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMouseEvent_IsEndEvent
func miqt_exec_callback_QMouseEvent_IsEndEvent(self QMouseEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMouseEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

type QHoverEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQHoverEvent constructs a new QHoverEvent object.
func NewQHoverEvent(typeVal Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF) *QHoverEvent {

	ret := newQHoverEvent(QHoverEvent_new(typeVal, scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent2 constructs a new QHoverEvent object.
func NewQHoverEvent2(typeVal Type, pos *QPointF, oldPos *QPointF) *QHoverEvent {

	ret := newQHoverEvent(QHoverEvent_new2(typeVal, pos.cPointer(), oldPos.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent3 constructs a new QHoverEvent object.
func NewQHoverEvent3(typeVal Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {

	ret := newQHoverEvent(QHoverEvent_new3(typeVal, scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer(), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent4 constructs a new QHoverEvent object.
func NewQHoverEvent4(typeVal Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {

	ret := newQHoverEvent(QHoverEvent_new4(typeVal, scenePos.cPointer(), globalPos.cPointer(), oldPos.cPointer(), (int)(modifiers), device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent5 constructs a new QHoverEvent object.
func NewQHoverEvent5(typeVal Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {

	ret := newQHoverEvent(QHoverEvent_new5(typeVal, pos.cPointer(), oldPos.cPointer(), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQHoverEvent6 constructs a new QHoverEvent object.
func NewQHoverEvent6(typeVal Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {

	ret := newQHoverEvent(QHoverEvent_new6(typeVal, pos.cPointer(), oldPos.cPointer(), (int)(modifiers), device.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QHoverEvent) callVirtualBase_Clone() *QHoverEvent {

	return newQHoverEvent(QHoverEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QHoverEvent) OnClone(slot func(super func() *QHoverEvent) *QHoverEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHoverEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_Clone
func miqt_exec_callback_QHoverEvent_Clone(self QHoverEvent, cb intptr_t) *QHoverEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QHoverEvent) *QHoverEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QHoverEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QHoverEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QHoverEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHoverEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_IsUpdateEvent
func miqt_exec_callback_QHoverEvent_IsUpdateEvent(self QHoverEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QHoverEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QHoverEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QHoverEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHoverEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_IsBeginEvent
func miqt_exec_callback_QHoverEvent_IsBeginEvent(self QHoverEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QHoverEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QHoverEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QHoverEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHoverEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHoverEvent_IsEndEvent
func miqt_exec_callback_QHoverEvent_IsEndEvent(self QHoverEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHoverEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

type QWheelEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQWheelEvent constructs a new QWheelEvent object.
func NewQWheelEvent(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool) *QWheelEvent {

	ret := newQWheelEvent(QWheelEvent_new(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (int)(buttons), (int)(modifiers), (int)(phase), (bool)(inverted)))
	ret.isSubclass = true
	return ret
}

// NewQWheelEvent2 constructs a new QWheelEvent object.
func NewQWheelEvent2(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource) *QWheelEvent {

	ret := newQWheelEvent(QWheelEvent_new2(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (int)(buttons), (int)(modifiers), (int)(phase), (bool)(inverted), (int)(source)))
	ret.isSubclass = true
	return ret
}

// NewQWheelEvent3 constructs a new QWheelEvent object.
func NewQWheelEvent3(pos *QPointF, globalPos *QPointF, pixelDelta QPoint, angleDelta QPoint, buttons MouseButton, modifiers KeyboardModifier, phase ScrollPhase, inverted bool, source MouseEventSource, device *QPointingDevice) *QWheelEvent {

	ret := newQWheelEvent(QWheelEvent_new3(pos.cPointer(), globalPos.cPointer(), pixelDelta.cPointer(), angleDelta.cPointer(), (int)(buttons), (int)(modifiers), (int)(phase), (bool)(inverted), (int)(source), device.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QWheelEvent) callVirtualBase_Clone() *QWheelEvent {

	return newQWheelEvent(QWheelEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QWheelEvent) OnClone(slot func(super func() *QWheelEvent) *QWheelEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWheelEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_Clone
func miqt_exec_callback_QWheelEvent_Clone(self QWheelEvent, cb intptr_t) *QWheelEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWheelEvent) *QWheelEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QWheelEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QWheelEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QWheelEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWheelEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_IsBeginEvent
func miqt_exec_callback_QWheelEvent_IsBeginEvent(self QWheelEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QWheelEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QWheelEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QWheelEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWheelEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_IsUpdateEvent
func miqt_exec_callback_QWheelEvent_IsUpdateEvent(self QWheelEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QWheelEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QWheelEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QWheelEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWheelEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWheelEvent_IsEndEvent
func miqt_exec_callback_QWheelEvent_IsEndEvent(self QWheelEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWheelEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

type QTabletEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQTabletEvent constructs a new QTabletEvent object.
func NewQTabletEvent(t Type, device *QPointingDevice, pos *QPointF, globalPos *QPointF, pressure float64, xTilt float32, yTilt float32, tangentialPressure float32, rotation float64, z float32, keyState KeyboardModifier, button MouseButton, buttons MouseButton) *QTabletEvent {

	ret := newQTabletEvent(QTabletEvent_new(t, device.cPointer(), pos.cPointer(), globalPos.cPointer(), (double)(pressure), (float)(xTilt), (float)(yTilt), (float)(tangentialPressure), (double)(rotation), (float)(z), (int)(keyState), (int)(button), (int)(buttons)))
	ret.isSubclass = true
	return ret
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

func (this *QTabletEvent) callVirtualBase_Clone() *QTabletEvent {

	return newQTabletEvent(QTabletEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QTabletEvent) OnClone(slot func(super func() *QTabletEvent) *QTabletEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabletEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_Clone
func miqt_exec_callback_QTabletEvent_Clone(self QTabletEvent, cb intptr_t) *QTabletEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QTabletEvent) *QTabletEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QTabletEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QTabletEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QTabletEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabletEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_IsBeginEvent
func miqt_exec_callback_QTabletEvent_IsBeginEvent(self QTabletEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QTabletEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QTabletEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QTabletEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabletEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_IsUpdateEvent
func miqt_exec_callback_QTabletEvent_IsUpdateEvent(self QTabletEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QTabletEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QTabletEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QTabletEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabletEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabletEvent_IsEndEvent
func miqt_exec_callback_QTabletEvent_IsEndEvent(self QTabletEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabletEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

type QNativeGestureEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQNativeGestureEvent constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent(typeVal NativeGestureType, dev *QPointingDevice, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, sequenceId uint64, intArgument uint64) *QNativeGestureEvent {

	ret := newQNativeGestureEvent(QNativeGestureEvent_new((int)(typeVal), dev.cPointer(), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (double)(value), (ulonglong)(sequenceId), (ulonglong)(intArgument)))
	ret.isSubclass = true
	return ret
}

// NewQNativeGestureEvent2 constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent2(typeVal NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF) *QNativeGestureEvent {

	ret := newQNativeGestureEvent(QNativeGestureEvent_new2((int)(typeVal), dev.cPointer(), (int)(fingerCount), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (double)(value), delta.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQNativeGestureEvent3 constructs a new QNativeGestureEvent object.
func NewQNativeGestureEvent3(typeVal NativeGestureType, dev *QPointingDevice, fingerCount int, localPos *QPointF, scenePos *QPointF, globalPos *QPointF, value float64, delta *QPointF, sequenceId uint64) *QNativeGestureEvent {

	ret := newQNativeGestureEvent(QNativeGestureEvent_new3((int)(typeVal), dev.cPointer(), (int)(fingerCount), localPos.cPointer(), scenePos.cPointer(), globalPos.cPointer(), (double)(value), delta.cPointer(), (ulonglong)(sequenceId)))
	ret.isSubclass = true
	return ret
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

func (this *QNativeGestureEvent) callVirtualBase_Clone() *QNativeGestureEvent {

	return newQNativeGestureEvent(QNativeGestureEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QNativeGestureEvent) OnClone(slot func(super func() *QNativeGestureEvent) *QNativeGestureEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNativeGestureEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_Clone
func miqt_exec_callback_QNativeGestureEvent_Clone(self QNativeGestureEvent, cb intptr_t) *QNativeGestureEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QNativeGestureEvent) *QNativeGestureEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QNativeGestureEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QNativeGestureEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QNativeGestureEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNativeGestureEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_IsBeginEvent
func miqt_exec_callback_QNativeGestureEvent_IsBeginEvent(self QNativeGestureEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QNativeGestureEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QNativeGestureEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QNativeGestureEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNativeGestureEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_IsUpdateEvent
func miqt_exec_callback_QNativeGestureEvent_IsUpdateEvent(self QNativeGestureEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QNativeGestureEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QNativeGestureEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QNativeGestureEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNativeGestureEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNativeGestureEvent_IsEndEvent
func miqt_exec_callback_QNativeGestureEvent_IsEndEvent(self QNativeGestureEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNativeGestureEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

type QKeyEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQKeyEvent constructs a new QKeyEvent object.
func NewQKeyEvent(typeVal Type, key int, modifiers KeyboardModifier) *QKeyEvent {

	ret := newQKeyEvent(QKeyEvent_new(typeVal, (int)(key), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent2 constructs a new QKeyEvent object.
func NewQKeyEvent2(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint) *QKeyEvent {

	ret := newQKeyEvent(QKeyEvent_new2(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers)))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent3 constructs a new QKeyEvent object.
func NewQKeyEvent3(typeVal Type, key int, modifiers KeyboardModifier, text string) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQKeyEvent(QKeyEvent_new3(typeVal, (int)(key), (int)(modifiers), text_ms))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent4 constructs a new QKeyEvent object.
func NewQKeyEvent4(typeVal Type, key int, modifiers KeyboardModifier, text string, autorep bool) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQKeyEvent(QKeyEvent_new4(typeVal, (int)(key), (int)(modifiers), text_ms, (bool)(autorep)))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent5 constructs a new QKeyEvent object.
func NewQKeyEvent5(typeVal Type, key int, modifiers KeyboardModifier, text string, autorep bool, count uint16) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQKeyEvent(QKeyEvent_new5(typeVal, (int)(key), (int)(modifiers), text_ms, (bool)(autorep), (uint16_t)(count)))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent6 constructs a new QKeyEvent object.
func NewQKeyEvent6(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQKeyEvent(QKeyEvent_new6(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent7 constructs a new QKeyEvent object.
func NewQKeyEvent7(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQKeyEvent(QKeyEvent_new7(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms, (bool)(autorep)))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent8 constructs a new QKeyEvent object.
func NewQKeyEvent8(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQKeyEvent(QKeyEvent_new8(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms, (bool)(autorep), (uint16_t)(count)))
	ret.isSubclass = true
	return ret
}

// NewQKeyEvent9 constructs a new QKeyEvent object.
func NewQKeyEvent9(typeVal Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16, device *QInputDevice) *QKeyEvent {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQKeyEvent(QKeyEvent_new9(typeVal, (int)(key), (int)(modifiers), (uint)(nativeScanCode), (uint)(nativeVirtualKey), (uint)(nativeModifiers), text_ms, (bool)(autorep), (uint16_t)(count), device.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QKeyEvent) callVirtualBase_Clone() *QKeyEvent {

	return newQKeyEvent(QKeyEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QKeyEvent) OnClone(slot func(super func() *QKeyEvent) *QKeyEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeyEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeyEvent_Clone
func miqt_exec_callback_QKeyEvent_Clone(self QKeyEvent, cb intptr_t) *QKeyEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QKeyEvent) *QKeyEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeyEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QKeyEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	QKeyEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (ulonglong)(timestamp))

}
func (this *QKeyEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeyEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeyEvent_SetTimestamp
func miqt_exec_callback_QKeyEvent_SetTimestamp(self QKeyEvent, cb intptr_t, timestamp ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QKeyEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

type QFocusEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQFocusEvent constructs a new QFocusEvent object.
func NewQFocusEvent(typeVal Type) *QFocusEvent {

	ret := newQFocusEvent(QFocusEvent_new(typeVal))
	ret.isSubclass = true
	return ret
}

// NewQFocusEvent2 constructs a new QFocusEvent object.
func NewQFocusEvent2(typeVal Type, reason FocusReason) *QFocusEvent {

	ret := newQFocusEvent(QFocusEvent_new2(typeVal, (int)(reason)))
	ret.isSubclass = true
	return ret
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

func (this *QFocusEvent) callVirtualBase_Clone() *QFocusEvent {

	return newQFocusEvent(QFocusEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QFocusEvent) OnClone(slot func(super func() *QFocusEvent) *QFocusEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFocusEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFocusEvent_Clone
func miqt_exec_callback_QFocusEvent_Clone(self QFocusEvent, cb intptr_t) *QFocusEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QFocusEvent) *QFocusEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFocusEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QFocusEvent) callVirtualBase_SetAccepted(accepted bool) {

	QFocusEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QFocusEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFocusEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFocusEvent_SetAccepted
func miqt_exec_callback_QFocusEvent_SetAccepted(self QFocusEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QFocusEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QPaintEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQPaintEvent constructs a new QPaintEvent object.
func NewQPaintEvent(paintRegion *QRegion) *QPaintEvent {

	ret := newQPaintEvent(QPaintEvent_new(paintRegion.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPaintEvent2 constructs a new QPaintEvent object.
func NewQPaintEvent2(paintRect *QRect) *QPaintEvent {

	ret := newQPaintEvent(QPaintEvent_new2(paintRect.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QPaintEvent) callVirtualBase_Clone() *QPaintEvent {

	return newQPaintEvent(QPaintEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QPaintEvent) OnClone(slot func(super func() *QPaintEvent) *QPaintEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEvent_Clone
func miqt_exec_callback_QPaintEvent_Clone(self QPaintEvent, cb intptr_t) *QPaintEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEvent) *QPaintEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPaintEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QPaintEvent) callVirtualBase_SetAccepted(accepted bool) {

	QPaintEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QPaintEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEvent_SetAccepted
func miqt_exec_callback_QPaintEvent_SetAccepted(self QPaintEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QPaintEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QMoveEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQMoveEvent constructs a new QMoveEvent object.
func NewQMoveEvent(pos *QPoint, oldPos *QPoint) *QMoveEvent {

	ret := newQMoveEvent(QMoveEvent_new(pos.cPointer(), oldPos.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QMoveEvent) callVirtualBase_Clone() *QMoveEvent {

	return newQMoveEvent(QMoveEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QMoveEvent) OnClone(slot func(super func() *QMoveEvent) *QMoveEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMoveEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMoveEvent_Clone
func miqt_exec_callback_QMoveEvent_Clone(self QMoveEvent, cb intptr_t) *QMoveEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMoveEvent) *QMoveEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMoveEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QMoveEvent) callVirtualBase_SetAccepted(accepted bool) {

	QMoveEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QMoveEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMoveEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMoveEvent_SetAccepted
func miqt_exec_callback_QMoveEvent_SetAccepted(self QMoveEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QMoveEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QExposeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQExposeEvent constructs a new QExposeEvent object.
func NewQExposeEvent(m_region *QRegion) *QExposeEvent {

	ret := newQExposeEvent(QExposeEvent_new(m_region.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QExposeEvent) Clone() *QExposeEvent {
	return newQExposeEvent(QExposeEvent_Clone(this.h))
}

func (this *QExposeEvent) Region() *QRegion {
	return newQRegion(QExposeEvent_Region(this.h))
}

func (this *QExposeEvent) callVirtualBase_Clone() *QExposeEvent {

	return newQExposeEvent(QExposeEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QExposeEvent) OnClone(slot func(super func() *QExposeEvent) *QExposeEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QExposeEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QExposeEvent_Clone
func miqt_exec_callback_QExposeEvent_Clone(self QExposeEvent, cb intptr_t) *QExposeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QExposeEvent) *QExposeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QExposeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QExposeEvent) callVirtualBase_SetAccepted(accepted bool) {

	QExposeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QExposeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QExposeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QExposeEvent_SetAccepted
func miqt_exec_callback_QExposeEvent_SetAccepted(self QExposeEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QExposeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QPlatformSurfaceEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQPlatformSurfaceEvent constructs a new QPlatformSurfaceEvent object.
func NewQPlatformSurfaceEvent(surfaceEventType SurfaceEventType) *QPlatformSurfaceEvent {

	ret := newQPlatformSurfaceEvent(QPlatformSurfaceEvent_new(surfaceEventType))
	ret.isSubclass = true
	return ret
}

func (this *QPlatformSurfaceEvent) Clone() *QPlatformSurfaceEvent {
	return newQPlatformSurfaceEvent(QPlatformSurfaceEvent_Clone(this.h))
}

func (this *QPlatformSurfaceEvent) SurfaceEventType() SurfaceEventType {
	xxxxxxxxx
}

func (this *QPlatformSurfaceEvent) callVirtualBase_Clone() *QPlatformSurfaceEvent {

	return newQPlatformSurfaceEvent(QPlatformSurfaceEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QPlatformSurfaceEvent) OnClone(slot func(super func() *QPlatformSurfaceEvent) *QPlatformSurfaceEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPlatformSurfaceEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlatformSurfaceEvent_Clone
func miqt_exec_callback_QPlatformSurfaceEvent_Clone(self QPlatformSurfaceEvent, cb intptr_t) *QPlatformSurfaceEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPlatformSurfaceEvent) *QPlatformSurfaceEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPlatformSurfaceEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QPlatformSurfaceEvent) callVirtualBase_SetAccepted(accepted bool) {

	QPlatformSurfaceEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QPlatformSurfaceEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPlatformSurfaceEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlatformSurfaceEvent_SetAccepted
func miqt_exec_callback_QPlatformSurfaceEvent_SetAccepted(self QPlatformSurfaceEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QPlatformSurfaceEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QResizeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQResizeEvent constructs a new QResizeEvent object.
func NewQResizeEvent(size *QSize, oldSize *QSize) *QResizeEvent {

	ret := newQResizeEvent(QResizeEvent_new(size.cPointer(), oldSize.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QResizeEvent) callVirtualBase_Clone() *QResizeEvent {

	return newQResizeEvent(QResizeEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QResizeEvent) OnClone(slot func(super func() *QResizeEvent) *QResizeEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QResizeEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QResizeEvent_Clone
func miqt_exec_callback_QResizeEvent_Clone(self QResizeEvent, cb intptr_t) *QResizeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QResizeEvent) *QResizeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QResizeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QResizeEvent) callVirtualBase_SetAccepted(accepted bool) {

	QResizeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QResizeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QResizeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QResizeEvent_SetAccepted
func miqt_exec_callback_QResizeEvent_SetAccepted(self QResizeEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QResizeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QCloseEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQCloseEvent constructs a new QCloseEvent object.
func NewQCloseEvent() *QCloseEvent {

	ret := newQCloseEvent(QCloseEvent_new())
	ret.isSubclass = true
	return ret
}

func (this *QCloseEvent) Clone() *QCloseEvent {
	return newQCloseEvent(QCloseEvent_Clone(this.h))
}

func (this *QCloseEvent) callVirtualBase_Clone() *QCloseEvent {

	return newQCloseEvent(QCloseEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QCloseEvent) OnClone(slot func(super func() *QCloseEvent) *QCloseEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCloseEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCloseEvent_Clone
func miqt_exec_callback_QCloseEvent_Clone(self QCloseEvent, cb intptr_t) *QCloseEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QCloseEvent) *QCloseEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCloseEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QCloseEvent) callVirtualBase_SetAccepted(accepted bool) {

	QCloseEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QCloseEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCloseEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCloseEvent_SetAccepted
func miqt_exec_callback_QCloseEvent_SetAccepted(self QCloseEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QCloseEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QIconDragEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQIconDragEvent constructs a new QIconDragEvent object.
func NewQIconDragEvent() *QIconDragEvent {

	ret := newQIconDragEvent(QIconDragEvent_new())
	ret.isSubclass = true
	return ret
}

func (this *QIconDragEvent) Clone() *QIconDragEvent {
	return newQIconDragEvent(QIconDragEvent_Clone(this.h))
}

func (this *QIconDragEvent) callVirtualBase_Clone() *QIconDragEvent {

	return newQIconDragEvent(QIconDragEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QIconDragEvent) OnClone(slot func(super func() *QIconDragEvent) *QIconDragEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconDragEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconDragEvent_Clone
func miqt_exec_callback_QIconDragEvent_Clone(self QIconDragEvent, cb intptr_t) *QIconDragEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QIconDragEvent) *QIconDragEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconDragEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QIconDragEvent) callVirtualBase_SetAccepted(accepted bool) {

	QIconDragEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QIconDragEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconDragEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconDragEvent_SetAccepted
func miqt_exec_callback_QIconDragEvent_SetAccepted(self QIconDragEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QIconDragEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QShowEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQShowEvent constructs a new QShowEvent object.
func NewQShowEvent() *QShowEvent {

	ret := newQShowEvent(QShowEvent_new())
	ret.isSubclass = true
	return ret
}

func (this *QShowEvent) Clone() *QShowEvent {
	return newQShowEvent(QShowEvent_Clone(this.h))
}

func (this *QShowEvent) callVirtualBase_Clone() *QShowEvent {

	return newQShowEvent(QShowEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QShowEvent) OnClone(slot func(super func() *QShowEvent) *QShowEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShowEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShowEvent_Clone
func miqt_exec_callback_QShowEvent_Clone(self QShowEvent, cb intptr_t) *QShowEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QShowEvent) *QShowEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QShowEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QShowEvent) callVirtualBase_SetAccepted(accepted bool) {

	QShowEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QShowEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShowEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShowEvent_SetAccepted
func miqt_exec_callback_QShowEvent_SetAccepted(self QShowEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QShowEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QHideEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQHideEvent constructs a new QHideEvent object.
func NewQHideEvent() *QHideEvent {

	ret := newQHideEvent(QHideEvent_new())
	ret.isSubclass = true
	return ret
}

func (this *QHideEvent) Clone() *QHideEvent {
	return newQHideEvent(QHideEvent_Clone(this.h))
}

func (this *QHideEvent) callVirtualBase_Clone() *QHideEvent {

	return newQHideEvent(QHideEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QHideEvent) OnClone(slot func(super func() *QHideEvent) *QHideEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHideEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHideEvent_Clone
func miqt_exec_callback_QHideEvent_Clone(self QHideEvent, cb intptr_t) *QHideEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QHideEvent) *QHideEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHideEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QHideEvent) callVirtualBase_SetAccepted(accepted bool) {

	QHideEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QHideEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHideEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHideEvent_SetAccepted
func miqt_exec_callback_QHideEvent_SetAccepted(self QHideEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QHideEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QContextMenuEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQContextMenuEvent constructs a new QContextMenuEvent object.
func NewQContextMenuEvent(reason Reason, pos *QPoint, globalPos *QPoint) *QContextMenuEvent {

	ret := newQContextMenuEvent(QContextMenuEvent_new(reason, pos.cPointer(), globalPos.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQContextMenuEvent2 constructs a new QContextMenuEvent object.
func NewQContextMenuEvent2(reason Reason, pos *QPoint) *QContextMenuEvent {

	ret := newQContextMenuEvent(QContextMenuEvent_new2(reason, pos.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQContextMenuEvent3 constructs a new QContextMenuEvent object.
func NewQContextMenuEvent3(reason Reason, pos *QPoint, globalPos *QPoint, modifiers KeyboardModifier) *QContextMenuEvent {

	ret := newQContextMenuEvent(QContextMenuEvent_new3(reason, pos.cPointer(), globalPos.cPointer(), (int)(modifiers)))
	ret.isSubclass = true
	return ret
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

func (this *QContextMenuEvent) callVirtualBase_Clone() *QContextMenuEvent {

	return newQContextMenuEvent(QContextMenuEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QContextMenuEvent) OnClone(slot func(super func() *QContextMenuEvent) *QContextMenuEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QContextMenuEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QContextMenuEvent_Clone
func miqt_exec_callback_QContextMenuEvent_Clone(self QContextMenuEvent, cb intptr_t) *QContextMenuEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QContextMenuEvent) *QContextMenuEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QContextMenuEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QContextMenuEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	QContextMenuEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (ulonglong)(timestamp))

}
func (this *QContextMenuEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QContextMenuEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QContextMenuEvent_SetTimestamp
func miqt_exec_callback_QContextMenuEvent_SetTimestamp(self QContextMenuEvent, cb intptr_t, timestamp ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QContextMenuEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

type QInputMethodEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQInputMethodEvent constructs a new QInputMethodEvent object.
func NewQInputMethodEvent() *QInputMethodEvent {

	ret := newQInputMethodEvent(QInputMethodEvent_new())
	ret.isSubclass = true
	return ret
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

	ret := newQInputMethodEvent(QInputMethodEvent_new2(preeditText_ms, attributes_ma))
	ret.isSubclass = true
	return ret
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

func (this *QInputMethodEvent) callVirtualBase_Clone() *QInputMethodEvent {

	return newQInputMethodEvent(QInputMethodEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QInputMethodEvent) OnClone(slot func(super func() *QInputMethodEvent) *QInputMethodEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputMethodEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodEvent_Clone
func miqt_exec_callback_QInputMethodEvent_Clone(self QInputMethodEvent, cb intptr_t) *QInputMethodEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QInputMethodEvent) *QInputMethodEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputMethodEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QInputMethodEvent) callVirtualBase_SetAccepted(accepted bool) {

	QInputMethodEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QInputMethodEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputMethodEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodEvent_SetAccepted
func miqt_exec_callback_QInputMethodEvent_SetAccepted(self QInputMethodEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QInputMethodEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QInputMethodQueryEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQInputMethodQueryEvent constructs a new QInputMethodQueryEvent object.
func NewQInputMethodQueryEvent(queries InputMethodQuery) *QInputMethodQueryEvent {

	ret := newQInputMethodQueryEvent(QInputMethodQueryEvent_new((int)(queries)))
	ret.isSubclass = true
	return ret
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

func (this *QInputMethodQueryEvent) callVirtualBase_Clone() *QInputMethodQueryEvent {

	return newQInputMethodQueryEvent(QInputMethodQueryEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QInputMethodQueryEvent) OnClone(slot func(super func() *QInputMethodQueryEvent) *QInputMethodQueryEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputMethodQueryEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodQueryEvent_Clone
func miqt_exec_callback_QInputMethodQueryEvent_Clone(self QInputMethodQueryEvent, cb intptr_t) *QInputMethodQueryEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QInputMethodQueryEvent) *QInputMethodQueryEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputMethodQueryEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QInputMethodQueryEvent) callVirtualBase_SetAccepted(accepted bool) {

	QInputMethodQueryEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QInputMethodQueryEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputMethodQueryEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethodQueryEvent_SetAccepted
func miqt_exec_callback_QInputMethodQueryEvent_SetAccepted(self QInputMethodQueryEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QInputMethodQueryEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QDropEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDropEvent constructs a new QDropEvent object.
func NewQDropEvent(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDropEvent {

	ret := newQDropEvent(QDropEvent_new(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQDropEvent2 constructs a new QDropEvent object.
func NewQDropEvent2(pos *QPointF, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, typeVal Type) *QDropEvent {

	ret := newQDropEvent(QDropEvent_new2(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers), typeVal))
	ret.isSubclass = true
	return ret
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

func (this *QDropEvent) callVirtualBase_Clone() *QDropEvent {

	return newQDropEvent(QDropEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QDropEvent) OnClone(slot func(super func() *QDropEvent) *QDropEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDropEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDropEvent_Clone
func miqt_exec_callback_QDropEvent_Clone(self QDropEvent, cb intptr_t) *QDropEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDropEvent) *QDropEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDropEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QDropEvent) callVirtualBase_SetAccepted(accepted bool) {

	QDropEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QDropEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDropEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDropEvent_SetAccepted
func miqt_exec_callback_QDropEvent_SetAccepted(self QDropEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QDropEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QDragMoveEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDragMoveEvent constructs a new QDragMoveEvent object.
func NewQDragMoveEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragMoveEvent {

	ret := newQDragMoveEvent(QDragMoveEvent_new(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQDragMoveEvent2 constructs a new QDragMoveEvent object.
func NewQDragMoveEvent2(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier, typeVal Type) *QDragMoveEvent {

	ret := newQDragMoveEvent(QDragMoveEvent_new2(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers), typeVal))
	ret.isSubclass = true
	return ret
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

func (this *QDragMoveEvent) callVirtualBase_Clone() *QDragMoveEvent {

	return newQDragMoveEvent(QDragMoveEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QDragMoveEvent) OnClone(slot func(super func() *QDragMoveEvent) *QDragMoveEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDragMoveEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragMoveEvent_Clone
func miqt_exec_callback_QDragMoveEvent_Clone(self QDragMoveEvent, cb intptr_t) *QDragMoveEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDragMoveEvent) *QDragMoveEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDragMoveEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

type QDragEnterEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDragEnterEvent constructs a new QDragEnterEvent object.
func NewQDragEnterEvent(pos *QPoint, actions DropAction, data *QMimeData, buttons MouseButton, modifiers KeyboardModifier) *QDragEnterEvent {

	ret := newQDragEnterEvent(QDragEnterEvent_new(pos.cPointer(), (int)(actions), data.cPointer(), (int)(buttons), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

func (this *QDragEnterEvent) Clone() *QDragEnterEvent {
	return newQDragEnterEvent(QDragEnterEvent_Clone(this.h))
}

func (this *QDragEnterEvent) callVirtualBase_Clone() *QDragEnterEvent {

	return newQDragEnterEvent(QDragEnterEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QDragEnterEvent) OnClone(slot func(super func() *QDragEnterEvent) *QDragEnterEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDragEnterEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragEnterEvent_Clone
func miqt_exec_callback_QDragEnterEvent_Clone(self QDragEnterEvent, cb intptr_t) *QDragEnterEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDragEnterEvent) *QDragEnterEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDragEnterEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

type QDragLeaveEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQDragLeaveEvent constructs a new QDragLeaveEvent object.
func NewQDragLeaveEvent() *QDragLeaveEvent {

	ret := newQDragLeaveEvent(QDragLeaveEvent_new())
	ret.isSubclass = true
	return ret
}

func (this *QDragLeaveEvent) Clone() *QDragLeaveEvent {
	return newQDragLeaveEvent(QDragLeaveEvent_Clone(this.h))
}

func (this *QDragLeaveEvent) callVirtualBase_Clone() *QDragLeaveEvent {

	return newQDragLeaveEvent(QDragLeaveEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QDragLeaveEvent) OnClone(slot func(super func() *QDragLeaveEvent) *QDragLeaveEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDragLeaveEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragLeaveEvent_Clone
func miqt_exec_callback_QDragLeaveEvent_Clone(self QDragLeaveEvent, cb intptr_t) *QDragLeaveEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QDragLeaveEvent) *QDragLeaveEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDragLeaveEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QDragLeaveEvent) callVirtualBase_SetAccepted(accepted bool) {

	QDragLeaveEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QDragLeaveEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDragLeaveEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDragLeaveEvent_SetAccepted
func miqt_exec_callback_QDragLeaveEvent_SetAccepted(self QDragLeaveEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QDragLeaveEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QHelpEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQHelpEvent constructs a new QHelpEvent object.
func NewQHelpEvent(typeVal Type, pos *QPoint, globalPos *QPoint) *QHelpEvent {

	ret := newQHelpEvent(QHelpEvent_new(typeVal, pos.cPointer(), globalPos.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QHelpEvent) callVirtualBase_Clone() *QHelpEvent {

	return newQHelpEvent(QHelpEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QHelpEvent) OnClone(slot func(super func() *QHelpEvent) *QHelpEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHelpEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHelpEvent_Clone
func miqt_exec_callback_QHelpEvent_Clone(self QHelpEvent, cb intptr_t) *QHelpEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QHelpEvent) *QHelpEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHelpEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QHelpEvent) callVirtualBase_SetAccepted(accepted bool) {

	QHelpEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QHelpEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHelpEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHelpEvent_SetAccepted
func miqt_exec_callback_QHelpEvent_SetAccepted(self QHelpEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QHelpEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

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

	ret := newQStatusTipEvent(QStatusTipEvent_new(tip_ms))
	ret.isSubclass = true
	return ret
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

func (this *QStatusTipEvent) callVirtualBase_Clone() *QStatusTipEvent {

	return newQStatusTipEvent(QStatusTipEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QStatusTipEvent) OnClone(slot func(super func() *QStatusTipEvent) *QStatusTipEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusTipEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusTipEvent_Clone
func miqt_exec_callback_QStatusTipEvent_Clone(self QStatusTipEvent, cb intptr_t) *QStatusTipEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QStatusTipEvent) *QStatusTipEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusTipEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QStatusTipEvent) callVirtualBase_SetAccepted(accepted bool) {

	QStatusTipEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QStatusTipEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusTipEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusTipEvent_SetAccepted
func miqt_exec_callback_QStatusTipEvent_SetAccepted(self QStatusTipEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QStatusTipEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

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

	ret := newQWhatsThisClickedEvent(QWhatsThisClickedEvent_new(href_ms))
	ret.isSubclass = true
	return ret
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

func (this *QWhatsThisClickedEvent) callVirtualBase_Clone() *QWhatsThisClickedEvent {

	return newQWhatsThisClickedEvent(QWhatsThisClickedEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QWhatsThisClickedEvent) OnClone(slot func(super func() *QWhatsThisClickedEvent) *QWhatsThisClickedEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWhatsThisClickedEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWhatsThisClickedEvent_Clone
func miqt_exec_callback_QWhatsThisClickedEvent_Clone(self QWhatsThisClickedEvent, cb intptr_t) *QWhatsThisClickedEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWhatsThisClickedEvent) *QWhatsThisClickedEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWhatsThisClickedEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QWhatsThisClickedEvent) callVirtualBase_SetAccepted(accepted bool) {

	QWhatsThisClickedEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QWhatsThisClickedEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWhatsThisClickedEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWhatsThisClickedEvent_SetAccepted
func miqt_exec_callback_QWhatsThisClickedEvent_SetAccepted(self QWhatsThisClickedEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QWhatsThisClickedEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QActionEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQActionEvent constructs a new QActionEvent object.
func NewQActionEvent(typeVal int, action *QAction) *QActionEvent {

	ret := newQActionEvent(QActionEvent_new((int)(typeVal), action.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQActionEvent2 constructs a new QActionEvent object.
func NewQActionEvent2(typeVal int, action *QAction, before *QAction) *QActionEvent {

	ret := newQActionEvent(QActionEvent_new2((int)(typeVal), action.cPointer(), before.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QActionEvent) callVirtualBase_Clone() *QActionEvent {

	return newQActionEvent(QActionEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QActionEvent) OnClone(slot func(super func() *QActionEvent) *QActionEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionEvent_Clone
func miqt_exec_callback_QActionEvent_Clone(self QActionEvent, cb intptr_t) *QActionEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QActionEvent) *QActionEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QActionEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QActionEvent) callVirtualBase_SetAccepted(accepted bool) {

	QActionEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QActionEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionEvent_SetAccepted
func miqt_exec_callback_QActionEvent_SetAccepted(self QActionEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QActionEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

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

	ret := newQFileOpenEvent(QFileOpenEvent_new(file_ms))
	ret.isSubclass = true
	return ret
}

// NewQFileOpenEvent2 constructs a new QFileOpenEvent object.
func NewQFileOpenEvent2(url *QUrl) *QFileOpenEvent {

	ret := newQFileOpenEvent(QFileOpenEvent_new2(url.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QFileOpenEvent) callVirtualBase_Clone() *QFileOpenEvent {

	return newQFileOpenEvent(QFileOpenEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QFileOpenEvent) OnClone(slot func(super func() *QFileOpenEvent) *QFileOpenEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileOpenEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileOpenEvent_Clone
func miqt_exec_callback_QFileOpenEvent_Clone(self QFileOpenEvent, cb intptr_t) *QFileOpenEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QFileOpenEvent) *QFileOpenEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileOpenEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QFileOpenEvent) callVirtualBase_SetAccepted(accepted bool) {

	QFileOpenEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QFileOpenEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileOpenEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileOpenEvent_SetAccepted
func miqt_exec_callback_QFileOpenEvent_SetAccepted(self QFileOpenEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QFileOpenEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QToolBarChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQToolBarChangeEvent constructs a new QToolBarChangeEvent object.
func NewQToolBarChangeEvent(t bool) *QToolBarChangeEvent {

	ret := newQToolBarChangeEvent(QToolBarChangeEvent_new((bool)(t)))
	ret.isSubclass = true
	return ret
}

func (this *QToolBarChangeEvent) Clone() *QToolBarChangeEvent {
	return newQToolBarChangeEvent(QToolBarChangeEvent_Clone(this.h))
}

func (this *QToolBarChangeEvent) Toggle() bool {
	return (bool)(QToolBarChangeEvent_Toggle(this.h))
}

func (this *QToolBarChangeEvent) callVirtualBase_Clone() *QToolBarChangeEvent {

	return newQToolBarChangeEvent(QToolBarChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QToolBarChangeEvent) OnClone(slot func(super func() *QToolBarChangeEvent) *QToolBarChangeEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBarChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBarChangeEvent_Clone
func miqt_exec_callback_QToolBarChangeEvent_Clone(self QToolBarChangeEvent, cb intptr_t) *QToolBarChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QToolBarChangeEvent) *QToolBarChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBarChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QToolBarChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	QToolBarChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QToolBarChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBarChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBarChangeEvent_SetAccepted
func miqt_exec_callback_QToolBarChangeEvent_SetAccepted(self QToolBarChangeEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QToolBarChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QShortcutEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQShortcutEvent constructs a new QShortcutEvent object.
func NewQShortcutEvent(key *QKeySequence, id int) *QShortcutEvent {

	ret := newQShortcutEvent(QShortcutEvent_new(key.cPointer(), (int)(id)))
	ret.isSubclass = true
	return ret
}

// NewQShortcutEvent2 constructs a new QShortcutEvent object.
func NewQShortcutEvent2(key *QKeySequence) *QShortcutEvent {

	ret := newQShortcutEvent(QShortcutEvent_new2(key.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQShortcutEvent3 constructs a new QShortcutEvent object.
func NewQShortcutEvent3(key *QKeySequence, id int, ambiguous bool) *QShortcutEvent {

	ret := newQShortcutEvent(QShortcutEvent_new3(key.cPointer(), (int)(id), (bool)(ambiguous)))
	ret.isSubclass = true
	return ret
}

// NewQShortcutEvent4 constructs a new QShortcutEvent object.
func NewQShortcutEvent4(key *QKeySequence, shortcut *QShortcut) *QShortcutEvent {

	ret := newQShortcutEvent(QShortcutEvent_new4(key.cPointer(), shortcut.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQShortcutEvent5 constructs a new QShortcutEvent object.
func NewQShortcutEvent5(key *QKeySequence, shortcut *QShortcut, ambiguous bool) *QShortcutEvent {

	ret := newQShortcutEvent(QShortcutEvent_new5(key.cPointer(), shortcut.cPointer(), (bool)(ambiguous)))
	ret.isSubclass = true
	return ret
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

func (this *QShortcutEvent) callVirtualBase_Clone() *QShortcutEvent {

	return newQShortcutEvent(QShortcutEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QShortcutEvent) OnClone(slot func(super func() *QShortcutEvent) *QShortcutEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcutEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcutEvent_Clone
func miqt_exec_callback_QShortcutEvent_Clone(self QShortcutEvent, cb intptr_t) *QShortcutEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QShortcutEvent) *QShortcutEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QShortcutEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QShortcutEvent) callVirtualBase_SetAccepted(accepted bool) {

	QShortcutEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QShortcutEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QShortcutEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QShortcutEvent_SetAccepted
func miqt_exec_callback_QShortcutEvent_SetAccepted(self QShortcutEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QShortcutEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QWindowStateChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQWindowStateChangeEvent constructs a new QWindowStateChangeEvent object.
func NewQWindowStateChangeEvent(oldState WindowState) *QWindowStateChangeEvent {

	ret := newQWindowStateChangeEvent(QWindowStateChangeEvent_new((int)(oldState)))
	ret.isSubclass = true
	return ret
}

// NewQWindowStateChangeEvent2 constructs a new QWindowStateChangeEvent object.
func NewQWindowStateChangeEvent2(oldState WindowState, isOverride bool) *QWindowStateChangeEvent {

	ret := newQWindowStateChangeEvent(QWindowStateChangeEvent_new2((int)(oldState), (bool)(isOverride)))
	ret.isSubclass = true
	return ret
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

func (this *QWindowStateChangeEvent) callVirtualBase_Clone() *QWindowStateChangeEvent {

	return newQWindowStateChangeEvent(QWindowStateChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QWindowStateChangeEvent) OnClone(slot func(super func() *QWindowStateChangeEvent) *QWindowStateChangeEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowStateChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowStateChangeEvent_Clone
func miqt_exec_callback_QWindowStateChangeEvent_Clone(self QWindowStateChangeEvent, cb intptr_t) *QWindowStateChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWindowStateChangeEvent) *QWindowStateChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindowStateChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QWindowStateChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	QWindowStateChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QWindowStateChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowStateChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowStateChangeEvent_SetAccepted
func miqt_exec_callback_QWindowStateChangeEvent_SetAccepted(self QWindowStateChangeEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QWindowStateChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QTouchEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQTouchEvent constructs a new QTouchEvent object.
func NewQTouchEvent(eventType QEvent__Type) *QTouchEvent {

	ret := newQTouchEvent(QTouchEvent_new((int)(eventType)))
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent2 constructs a new QTouchEvent object.
func NewQTouchEvent2(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates State) *QTouchEvent {

	ret := newQTouchEvent(QTouchEvent_new2((int)(eventType), device.cPointer(), (int)(modifiers), (int)(touchPointStates)))
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent3 constructs a new QTouchEvent object.
func NewQTouchEvent3(eventType QEvent__Type, device *QPointingDevice) *QTouchEvent {

	ret := newQTouchEvent(QTouchEvent_new3((int)(eventType), device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent4 constructs a new QTouchEvent object.
func NewQTouchEvent4(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier) *QTouchEvent {

	ret := newQTouchEvent(QTouchEvent_new4((int)(eventType), device.cPointer(), (int)(modifiers)))
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent5 constructs a new QTouchEvent object.
func NewQTouchEvent5(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPoints []QEventPoint) *QTouchEvent {
	touchPoints_CArray := (*[0xffff]*QEventPoint)(malloc(size_t(8 * len(touchPoints))))
	defer free(unsafe.Pointer(touchPoints_CArray))
	for i := range touchPoints {
		touchPoints_CArray[i] = touchPoints[i].cPointer()
	}
	touchPoints_ma := struct_miqt_array{len: size_t(len(touchPoints)), data: unsafe.Pointer(touchPoints_CArray)}

	ret := newQTouchEvent(QTouchEvent_new5((int)(eventType), device.cPointer(), (int)(modifiers), touchPoints_ma))
	ret.isSubclass = true
	return ret
}

// NewQTouchEvent6 constructs a new QTouchEvent object.
func NewQTouchEvent6(eventType QEvent__Type, device *QPointingDevice, modifiers KeyboardModifier, touchPointStates State, touchPoints []QEventPoint) *QTouchEvent {
	touchPoints_CArray := (*[0xffff]*QEventPoint)(malloc(size_t(8 * len(touchPoints))))
	defer free(unsafe.Pointer(touchPoints_CArray))
	for i := range touchPoints {
		touchPoints_CArray[i] = touchPoints[i].cPointer()
	}
	touchPoints_ma := struct_miqt_array{len: size_t(len(touchPoints)), data: unsafe.Pointer(touchPoints_CArray)}

	ret := newQTouchEvent(QTouchEvent_new6((int)(eventType), device.cPointer(), (int)(modifiers), (int)(touchPointStates), touchPoints_ma))
	ret.isSubclass = true
	return ret
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

func (this *QTouchEvent) callVirtualBase_Clone() *QTouchEvent {

	return newQTouchEvent(QTouchEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QTouchEvent) OnClone(slot func(super func() *QTouchEvent) *QTouchEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTouchEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_Clone
func miqt_exec_callback_QTouchEvent_Clone(self QTouchEvent, cb intptr_t) *QTouchEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QTouchEvent) *QTouchEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QTouchEvent) callVirtualBase_IsBeginEvent() bool {

	return (bool)(QTouchEvent_virtualbase_IsBeginEvent(unsafe.Pointer(this.h)))

}
func (this *QTouchEvent) OnIsBeginEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTouchEvent_override_virtual_IsBeginEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_IsBeginEvent
func miqt_exec_callback_QTouchEvent_IsBeginEvent(self QTouchEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_IsBeginEvent)

	return (bool)(virtualReturn)

}

func (this *QTouchEvent) callVirtualBase_IsUpdateEvent() bool {

	return (bool)(QTouchEvent_virtualbase_IsUpdateEvent(unsafe.Pointer(this.h)))

}
func (this *QTouchEvent) OnIsUpdateEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTouchEvent_override_virtual_IsUpdateEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_IsUpdateEvent
func miqt_exec_callback_QTouchEvent_IsUpdateEvent(self QTouchEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_IsUpdateEvent)

	return (bool)(virtualReturn)

}

func (this *QTouchEvent) callVirtualBase_IsEndEvent() bool {

	return (bool)(QTouchEvent_virtualbase_IsEndEvent(unsafe.Pointer(this.h)))

}
func (this *QTouchEvent) OnIsEndEvent(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTouchEvent_override_virtual_IsEndEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_IsEndEvent
func miqt_exec_callback_QTouchEvent_IsEndEvent(self QTouchEvent, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTouchEvent{h: self}).callVirtualBase_IsEndEvent)

	return (bool)(virtualReturn)

}

func (this *QTouchEvent) callVirtualBase_SetTimestamp(timestamp uint64) {

	QTouchEvent_virtualbase_SetTimestamp(unsafe.Pointer(this.h), (ulonglong)(timestamp))

}
func (this *QTouchEvent) OnSetTimestamp(slot func(super func(timestamp uint64), timestamp uint64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTouchEvent_override_virtual_SetTimestamp(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_SetTimestamp
func miqt_exec_callback_QTouchEvent_SetTimestamp(self QTouchEvent, cb intptr_t, timestamp ulonglong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timestamp uint64), timestamp uint64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (uint64)(timestamp)

	gofunc((&QTouchEvent{h: self}).callVirtualBase_SetTimestamp, slotval1)

}

func (this *QTouchEvent) callVirtualBase_SetAccepted(accepted bool) {

	QTouchEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QTouchEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTouchEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTouchEvent_SetAccepted
func miqt_exec_callback_QTouchEvent_SetAccepted(self QTouchEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QTouchEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QScrollPrepareEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQScrollPrepareEvent constructs a new QScrollPrepareEvent object.
func NewQScrollPrepareEvent(startPos *QPointF) *QScrollPrepareEvent {

	ret := newQScrollPrepareEvent(QScrollPrepareEvent_new(startPos.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QScrollPrepareEvent) callVirtualBase_Clone() *QScrollPrepareEvent {

	return newQScrollPrepareEvent(QScrollPrepareEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QScrollPrepareEvent) OnClone(slot func(super func() *QScrollPrepareEvent) *QScrollPrepareEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollPrepareEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollPrepareEvent_Clone
func miqt_exec_callback_QScrollPrepareEvent_Clone(self QScrollPrepareEvent, cb intptr_t) *QScrollPrepareEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QScrollPrepareEvent) *QScrollPrepareEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollPrepareEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QScrollPrepareEvent) callVirtualBase_SetAccepted(accepted bool) {

	QScrollPrepareEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QScrollPrepareEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollPrepareEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollPrepareEvent_SetAccepted
func miqt_exec_callback_QScrollPrepareEvent_SetAccepted(self QScrollPrepareEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QScrollPrepareEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QScrollEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQScrollEvent constructs a new QScrollEvent object.
func NewQScrollEvent(contentPos *QPointF, overshoot *QPointF, scrollState ScrollState) *QScrollEvent {

	ret := newQScrollEvent(QScrollEvent_new(contentPos.cPointer(), overshoot.cPointer(), scrollState))
	ret.isSubclass = true
	return ret
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

func (this *QScrollEvent) callVirtualBase_Clone() *QScrollEvent {

	return newQScrollEvent(QScrollEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QScrollEvent) OnClone(slot func(super func() *QScrollEvent) *QScrollEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollEvent_Clone
func miqt_exec_callback_QScrollEvent_Clone(self QScrollEvent, cb intptr_t) *QScrollEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QScrollEvent) *QScrollEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QScrollEvent) callVirtualBase_SetAccepted(accepted bool) {

	QScrollEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QScrollEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollEvent_SetAccepted
func miqt_exec_callback_QScrollEvent_SetAccepted(self QScrollEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QScrollEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QScreenOrientationChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQScreenOrientationChangeEvent constructs a new QScreenOrientationChangeEvent object.
func NewQScreenOrientationChangeEvent(screen *QScreen, orientation ScreenOrientation) *QScreenOrientationChangeEvent {

	ret := newQScreenOrientationChangeEvent(QScreenOrientationChangeEvent_new(screen.cPointer(), (int)(orientation)))
	ret.isSubclass = true
	return ret
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

func (this *QScreenOrientationChangeEvent) callVirtualBase_Clone() *QScreenOrientationChangeEvent {

	return newQScreenOrientationChangeEvent(QScreenOrientationChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QScreenOrientationChangeEvent) OnClone(slot func(super func() *QScreenOrientationChangeEvent) *QScreenOrientationChangeEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScreenOrientationChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenOrientationChangeEvent_Clone
func miqt_exec_callback_QScreenOrientationChangeEvent_Clone(self QScreenOrientationChangeEvent, cb intptr_t) *QScreenOrientationChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QScreenOrientationChangeEvent) *QScreenOrientationChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScreenOrientationChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QScreenOrientationChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	QScreenOrientationChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QScreenOrientationChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScreenOrientationChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenOrientationChangeEvent_SetAccepted
func miqt_exec_callback_QScreenOrientationChangeEvent_SetAccepted(self QScreenOrientationChangeEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QScreenOrientationChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QApplicationStateChangeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQApplicationStateChangeEvent constructs a new QApplicationStateChangeEvent object.
func NewQApplicationStateChangeEvent(state ApplicationState) *QApplicationStateChangeEvent {

	ret := newQApplicationStateChangeEvent(QApplicationStateChangeEvent_new((int)(state)))
	ret.isSubclass = true
	return ret
}

func (this *QApplicationStateChangeEvent) Clone() *QApplicationStateChangeEvent {
	return newQApplicationStateChangeEvent(QApplicationStateChangeEvent_Clone(this.h))
}

func (this *QApplicationStateChangeEvent) ApplicationState() ApplicationState {
	return (ApplicationState)(QApplicationStateChangeEvent_ApplicationState(this.h))
}

func (this *QApplicationStateChangeEvent) callVirtualBase_Clone() *QApplicationStateChangeEvent {

	return newQApplicationStateChangeEvent(QApplicationStateChangeEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QApplicationStateChangeEvent) OnClone(slot func(super func() *QApplicationStateChangeEvent) *QApplicationStateChangeEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QApplicationStateChangeEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QApplicationStateChangeEvent_Clone
func miqt_exec_callback_QApplicationStateChangeEvent_Clone(self QApplicationStateChangeEvent, cb intptr_t) *QApplicationStateChangeEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QApplicationStateChangeEvent) *QApplicationStateChangeEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QApplicationStateChangeEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QApplicationStateChangeEvent) callVirtualBase_SetAccepted(accepted bool) {

	QApplicationStateChangeEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QApplicationStateChangeEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QApplicationStateChangeEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QApplicationStateChangeEvent_SetAccepted
func miqt_exec_callback_QApplicationStateChangeEvent_SetAccepted(self QApplicationStateChangeEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QApplicationStateChangeEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QChildWindowEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQChildWindowEvent constructs a new QChildWindowEvent object.
func NewQChildWindowEvent(typeVal Type, childWindow *QWindow) *QChildWindowEvent {

	ret := newQChildWindowEvent(QChildWindowEvent_new(typeVal, childWindow.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QChildWindowEvent) Clone() *QChildWindowEvent {
	return newQChildWindowEvent(QChildWindowEvent_Clone(this.h))
}

func (this *QChildWindowEvent) Child() *QWindow {
	return newQWindow(QChildWindowEvent_Child(this.h))
}

func (this *QChildWindowEvent) callVirtualBase_Clone() *QChildWindowEvent {

	return newQChildWindowEvent(QChildWindowEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QChildWindowEvent) OnClone(slot func(super func() *QChildWindowEvent) *QChildWindowEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChildWindowEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChildWindowEvent_Clone
func miqt_exec_callback_QChildWindowEvent_Clone(self QChildWindowEvent, cb intptr_t) *QChildWindowEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QChildWindowEvent) *QChildWindowEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QChildWindowEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QChildWindowEvent) callVirtualBase_SetAccepted(accepted bool) {

	QChildWindowEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QChildWindowEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChildWindowEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChildWindowEvent_SetAccepted
func miqt_exec_callback_QChildWindowEvent_SetAccepted(self QChildWindowEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QChildWindowEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

type QInputMethodEvent__Attribute struct {
	h          uintptr
	isSubclass bool
}

// NewQInputMethodEvent__Attribute constructs a new QInputMethodEvent::Attribute object.
func NewQInputMethodEvent__Attribute(typ AttributeType, s int, l int, val QVariant) *QInputMethodEvent__Attribute {

	ret := newQInputMethodEvent__Attribute(QInputMethodEvent__Attribute_new(typ, (int)(s), (int)(l), val.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQInputMethodEvent__Attribute2 constructs a new QInputMethodEvent::Attribute object.
func NewQInputMethodEvent__Attribute2(typ AttributeType, s int, l int) *QInputMethodEvent__Attribute {

	ret := newQInputMethodEvent__Attribute(QInputMethodEvent__Attribute_new2(typ, (int)(s), (int)(l)))
	ret.isSubclass = true
	return ret
}

func (this *QInputMethodEvent__Attribute) OperatorAssign(param1 *Attribute) {
	QInputMethodEvent__Attribute_OperatorAssign(this.h, param1)
}
