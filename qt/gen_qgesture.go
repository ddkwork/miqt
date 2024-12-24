package qt

import (
	"unsafe"
)

type QGesture__GestureCancelPolicy int

const (
	QGesture__CancelNone         QGesture__GestureCancelPolicy = 0
	QGesture__CancelAllInContext QGesture__GestureCancelPolicy = 1
)

type QPinchGesture__ChangeFlag int

const (
	QPinchGesture__ScaleFactorChanged   QPinchGesture__ChangeFlag = 1
	QPinchGesture__RotationAngleChanged QPinchGesture__ChangeFlag = 2
	QPinchGesture__CenterPointChanged   QPinchGesture__ChangeFlag = 4
)

type QSwipeGesture__SwipeDirection int

const (
	QSwipeGesture__NoDirection QSwipeGesture__SwipeDirection = 0
	QSwipeGesture__Left        QSwipeGesture__SwipeDirection = 1
	QSwipeGesture__Right       QSwipeGesture__SwipeDirection = 2
	QSwipeGesture__Up          QSwipeGesture__SwipeDirection = 3
	QSwipeGesture__Down        QSwipeGesture__SwipeDirection = 4
)

type QGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQGesture constructs a new QGesture object.
func NewQGesture() *QGesture {

	ret := newQGesture(QGesture_new())
	ret.isSubclass = true
	return ret
}

// NewQGesture2 constructs a new QGesture object.
func NewQGesture2(parent *QObject) *QGesture {

	ret := newQGesture(QGesture_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGesture) MetaObject() *QMetaObject {
	return newQMetaObject(QGesture_MetaObject(this.h))
}

func (this *QGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGesture_Metacast(this.h, param1_Cstring))
}

func QGesture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGesture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGesture) GestureType() GestureType {
	return (GestureType)(QGesture_GestureType(this.h))
}

func (this *QGesture) State() GestureState {
	return (GestureState)(QGesture_State(this.h))
}

func (this *QGesture) HotSpot() *QPointF {
	_goptr := newQPointF(QGesture_HotSpot(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGesture) SetHotSpot(value *QPointF) {
	QGesture_SetHotSpot(this.h, value.cPointer())
}

func (this *QGesture) HasHotSpot() bool {
	return (bool)(QGesture_HasHotSpot(this.h))
}

func (this *QGesture) UnsetHotSpot() {
	QGesture_UnsetHotSpot(this.h)
}

func (this *QGesture) SetGestureCancelPolicy(policy GestureCancelPolicy) {
	QGesture_SetGestureCancelPolicy(this.h, policy)
}

func (this *QGesture) GestureCancelPolicy() GestureCancelPolicy {
	xxxxxxxxx
}

func QGesture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGesture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGesture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGesture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGesture) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGesture_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGesture) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_Event
func miqt_exec_callback_QGesture_Event(self QGesture, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGesture{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGesture) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QGesture_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QGesture) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_EventFilter
func miqt_exec_callback_QGesture_EventFilter(self QGesture, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QGesture{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QGesture) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QGesture_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGesture) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_TimerEvent
func miqt_exec_callback_QGesture_TimerEvent(self QGesture, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QGesture{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QGesture) callVirtualBase_ChildEvent(event *QChildEvent) {

	QGesture_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGesture) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_ChildEvent
func miqt_exec_callback_QGesture_ChildEvent(self QGesture, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QGesture{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QGesture) callVirtualBase_CustomEvent(event *QEvent) {

	QGesture_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGesture) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_CustomEvent
func miqt_exec_callback_QGesture_CustomEvent(self QGesture, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGesture{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QGesture) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QGesture_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGesture) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_ConnectNotify
func miqt_exec_callback_QGesture_ConnectNotify(self QGesture, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGesture{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QGesture) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QGesture_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGesture) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_DisconnectNotify
func miqt_exec_callback_QGesture_DisconnectNotify(self QGesture, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGesture{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

type QPanGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQPanGesture constructs a new QPanGesture object.
func NewQPanGesture() *QPanGesture {

	ret := newQPanGesture(QPanGesture_new())
	ret.isSubclass = true
	return ret
}

// NewQPanGesture2 constructs a new QPanGesture object.
func NewQPanGesture2(parent *QObject) *QPanGesture {

	ret := newQPanGesture(QPanGesture_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPanGesture) MetaObject() *QMetaObject {
	return newQMetaObject(QPanGesture_MetaObject(this.h))
}

func (this *QPanGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPanGesture_Metacast(this.h, param1_Cstring))
}

func QPanGesture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPanGesture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPanGesture) LastOffset() *QPointF {
	_goptr := newQPointF(QPanGesture_LastOffset(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPanGesture) Offset() *QPointF {
	_goptr := newQPointF(QPanGesture_Offset(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPanGesture) Delta() *QPointF {
	_goptr := newQPointF(QPanGesture_Delta(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPanGesture) Acceleration() float64 {
	return (float64)(QPanGesture_Acceleration(this.h))
}

func (this *QPanGesture) SetLastOffset(value *QPointF) {
	QPanGesture_SetLastOffset(this.h, value.cPointer())
}

func (this *QPanGesture) SetOffset(value *QPointF) {
	QPanGesture_SetOffset(this.h, value.cPointer())
}

func (this *QPanGesture) SetAcceleration(value float64) {
	QPanGesture_SetAcceleration(this.h, (double)(value))
}

func QPanGesture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPanGesture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPanGesture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPanGesture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QPinchGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQPinchGesture constructs a new QPinchGesture object.
func NewQPinchGesture() *QPinchGesture {

	ret := newQPinchGesture(QPinchGesture_new())
	ret.isSubclass = true
	return ret
}

// NewQPinchGesture2 constructs a new QPinchGesture object.
func NewQPinchGesture2(parent *QObject) *QPinchGesture {

	ret := newQPinchGesture(QPinchGesture_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPinchGesture) MetaObject() *QMetaObject {
	return newQMetaObject(QPinchGesture_MetaObject(this.h))
}

func (this *QPinchGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPinchGesture_Metacast(this.h, param1_Cstring))
}

func QPinchGesture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPinchGesture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPinchGesture) TotalChangeFlags() ChangeFlags {
	xxxxxxxxx
}

func (this *QPinchGesture) SetTotalChangeFlags(value ChangeFlags) {
	QPinchGesture_SetTotalChangeFlags(this.h, value)
}

func (this *QPinchGesture) ChangeFlags() ChangeFlags {
	xxxxxxxxx
}

func (this *QPinchGesture) SetChangeFlags(value ChangeFlags) {
	QPinchGesture_SetChangeFlags(this.h, value)
}

func (this *QPinchGesture) StartCenterPoint() *QPointF {
	_goptr := newQPointF(QPinchGesture_StartCenterPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPinchGesture) LastCenterPoint() *QPointF {
	_goptr := newQPointF(QPinchGesture_LastCenterPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPinchGesture) CenterPoint() *QPointF {
	_goptr := newQPointF(QPinchGesture_CenterPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPinchGesture) SetStartCenterPoint(value *QPointF) {
	QPinchGesture_SetStartCenterPoint(this.h, value.cPointer())
}

func (this *QPinchGesture) SetLastCenterPoint(value *QPointF) {
	QPinchGesture_SetLastCenterPoint(this.h, value.cPointer())
}

func (this *QPinchGesture) SetCenterPoint(value *QPointF) {
	QPinchGesture_SetCenterPoint(this.h, value.cPointer())
}

func (this *QPinchGesture) TotalScaleFactor() float64 {
	return (float64)(QPinchGesture_TotalScaleFactor(this.h))
}

func (this *QPinchGesture) LastScaleFactor() float64 {
	return (float64)(QPinchGesture_LastScaleFactor(this.h))
}

func (this *QPinchGesture) ScaleFactor() float64 {
	return (float64)(QPinchGesture_ScaleFactor(this.h))
}

func (this *QPinchGesture) SetTotalScaleFactor(value float64) {
	QPinchGesture_SetTotalScaleFactor(this.h, (double)(value))
}

func (this *QPinchGesture) SetLastScaleFactor(value float64) {
	QPinchGesture_SetLastScaleFactor(this.h, (double)(value))
}

func (this *QPinchGesture) SetScaleFactor(value float64) {
	QPinchGesture_SetScaleFactor(this.h, (double)(value))
}

func (this *QPinchGesture) TotalRotationAngle() float64 {
	return (float64)(QPinchGesture_TotalRotationAngle(this.h))
}

func (this *QPinchGesture) LastRotationAngle() float64 {
	return (float64)(QPinchGesture_LastRotationAngle(this.h))
}

func (this *QPinchGesture) RotationAngle() float64 {
	return (float64)(QPinchGesture_RotationAngle(this.h))
}

func (this *QPinchGesture) SetTotalRotationAngle(value float64) {
	QPinchGesture_SetTotalRotationAngle(this.h, (double)(value))
}

func (this *QPinchGesture) SetLastRotationAngle(value float64) {
	QPinchGesture_SetLastRotationAngle(this.h, (double)(value))
}

func (this *QPinchGesture) SetRotationAngle(value float64) {
	QPinchGesture_SetRotationAngle(this.h, (double)(value))
}

func QPinchGesture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPinchGesture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPinchGesture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPinchGesture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QSwipeGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQSwipeGesture constructs a new QSwipeGesture object.
func NewQSwipeGesture() *QSwipeGesture {

	ret := newQSwipeGesture(QSwipeGesture_new())
	ret.isSubclass = true
	return ret
}

// NewQSwipeGesture2 constructs a new QSwipeGesture object.
func NewQSwipeGesture2(parent *QObject) *QSwipeGesture {

	ret := newQSwipeGesture(QSwipeGesture_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSwipeGesture) MetaObject() *QMetaObject {
	return newQMetaObject(QSwipeGesture_MetaObject(this.h))
}

func (this *QSwipeGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSwipeGesture_Metacast(this.h, param1_Cstring))
}

func QSwipeGesture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSwipeGesture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSwipeGesture) HorizontalDirection() SwipeDirection {
	xxxxxxxxx
}

func (this *QSwipeGesture) VerticalDirection() SwipeDirection {
	xxxxxxxxx
}

func (this *QSwipeGesture) SwipeAngle() float64 {
	return (float64)(QSwipeGesture_SwipeAngle(this.h))
}

func (this *QSwipeGesture) SetSwipeAngle(value float64) {
	QSwipeGesture_SetSwipeAngle(this.h, (double)(value))
}

func QSwipeGesture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSwipeGesture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSwipeGesture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSwipeGesture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QTapGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQTapGesture constructs a new QTapGesture object.
func NewQTapGesture() *QTapGesture {

	ret := newQTapGesture(QTapGesture_new())
	ret.isSubclass = true
	return ret
}

// NewQTapGesture2 constructs a new QTapGesture object.
func NewQTapGesture2(parent *QObject) *QTapGesture {

	ret := newQTapGesture(QTapGesture_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTapGesture) MetaObject() *QMetaObject {
	return newQMetaObject(QTapGesture_MetaObject(this.h))
}

func (this *QTapGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTapGesture_Metacast(this.h, param1_Cstring))
}

func QTapGesture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTapGesture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTapGesture) Position() *QPointF {
	_goptr := newQPointF(QTapGesture_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTapGesture) SetPosition(pos *QPointF) {
	QTapGesture_SetPosition(this.h, pos.cPointer())
}

func QTapGesture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTapGesture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTapGesture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTapGesture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QTapAndHoldGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQTapAndHoldGesture constructs a new QTapAndHoldGesture object.
func NewQTapAndHoldGesture() *QTapAndHoldGesture {

	ret := newQTapAndHoldGesture(QTapAndHoldGesture_new())
	ret.isSubclass = true
	return ret
}

// NewQTapAndHoldGesture2 constructs a new QTapAndHoldGesture object.
func NewQTapAndHoldGesture2(parent *QObject) *QTapAndHoldGesture {

	ret := newQTapAndHoldGesture(QTapAndHoldGesture_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTapAndHoldGesture) MetaObject() *QMetaObject {
	return newQMetaObject(QTapAndHoldGesture_MetaObject(this.h))
}

func (this *QTapAndHoldGesture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTapAndHoldGesture_Metacast(this.h, param1_Cstring))
}

func QTapAndHoldGesture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTapAndHoldGesture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTapAndHoldGesture) Position() *QPointF {
	_goptr := newQPointF(QTapAndHoldGesture_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTapAndHoldGesture) SetPosition(pos *QPointF) {
	QTapAndHoldGesture_SetPosition(this.h, pos.cPointer())
}

func QTapAndHoldGesture_SetTimeout(msecs int) {
	QTapAndHoldGesture_SetTimeout((int)(msecs))
}

func QTapAndHoldGesture_Timeout() int {
	return (int)(QTapAndHoldGesture_Timeout())
}

func QTapAndHoldGesture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTapAndHoldGesture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTapAndHoldGesture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTapAndHoldGesture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QGestureEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGestureEvent constructs a new QGestureEvent object.
func NewQGestureEvent(gestures []*QGesture) *QGestureEvent {
	gestures_CArray := (*[0xffff]*QGesture)(malloc(size_t(8 * len(gestures))))
	defer free(unsafe.Pointer(gestures_CArray))
	for i := range gestures {
		gestures_CArray[i] = gestures[i].cPointer()
	}
	gestures_ma := struct_miqt_array{len: size_t(len(gestures)), data: unsafe.Pointer(gestures_CArray)}

	ret := newQGestureEvent(QGestureEvent_new(gestures_ma))
	ret.isSubclass = true
	return ret
}

// NewQGestureEvent2 constructs a new QGestureEvent object.
func NewQGestureEvent2(param1 *QGestureEvent) *QGestureEvent {

	ret := newQGestureEvent(QGestureEvent_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGestureEvent) Gestures() []*QGesture {
	var _ma struct_miqt_array = QGestureEvent_Gestures(this.h)
	_ret := make([]*QGesture, int(_ma.len))
	_outCast := (*[0xffff]*QGesture)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGesture(_outCast[i])
	}
	return _ret
}

func (this *QGestureEvent) Gesture(typeVal GestureType) *QGesture {
	return newQGesture(QGestureEvent_Gesture(this.h, (int)(typeVal)))
}

func (this *QGestureEvent) ActiveGestures() []*QGesture {
	var _ma struct_miqt_array = QGestureEvent_ActiveGestures(this.h)
	_ret := make([]*QGesture, int(_ma.len))
	_outCast := (*[0xffff]*QGesture)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGesture(_outCast[i])
	}
	return _ret
}

func (this *QGestureEvent) CanceledGestures() []*QGesture {
	var _ma struct_miqt_array = QGestureEvent_CanceledGestures(this.h)
	_ret := make([]*QGesture, int(_ma.len))
	_outCast := (*[0xffff]*QGesture)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGesture(_outCast[i])
	}
	return _ret
}

func (this *QGestureEvent) SetAccepted(param1 *QGesture, param2 bool) {
	QGestureEvent_SetAccepted(this.h, param1.cPointer(), (bool)(param2))
}

func (this *QGestureEvent) Accept(param1 *QGesture) {
	QGestureEvent_Accept(this.h, param1.cPointer())
}

func (this *QGestureEvent) Ignore(param1 *QGesture) {
	QGestureEvent_Ignore(this.h, param1.cPointer())
}

func (this *QGestureEvent) IsAccepted(param1 *QGesture) bool {
	return (bool)(QGestureEvent_IsAccepted(this.h, param1.cPointer()))
}

func (this *QGestureEvent) SetAccepted2(param1 GestureType, param2 bool) {
	QGestureEvent_SetAccepted2(this.h, (int)(param1), (bool)(param2))
}

func (this *QGestureEvent) AcceptWithQtGestureType(param1 GestureType) {
	QGestureEvent_AcceptWithQtGestureType(this.h, (int)(param1))
}

func (this *QGestureEvent) IgnoreWithQtGestureType(param1 GestureType) {
	QGestureEvent_IgnoreWithQtGestureType(this.h, (int)(param1))
}

func (this *QGestureEvent) IsAcceptedWithQtGestureType(param1 GestureType) bool {
	return (bool)(QGestureEvent_IsAcceptedWithQtGestureType(this.h, (int)(param1)))
}

func (this *QGestureEvent) SetWidget(widget *QWidget) {
	QGestureEvent_SetWidget(this.h, widget.cPointer())
}

func (this *QGestureEvent) Widget() *QWidget {
	return newQWidget(QGestureEvent_Widget(this.h))
}

func (this *QGestureEvent) MapToGraphicsScene(gesturePoint *QPointF) *QPointF {
	_goptr := newQPointF(QGestureEvent_MapToGraphicsScene(this.h, gesturePoint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGestureEvent) callVirtualBase_SetAccepted(accepted bool) {

	QGestureEvent_virtualbase_SetAccepted(unsafe.Pointer(this.h), (bool)(accepted))

}
func (this *QGestureEvent) OnSetAccepted(slot func(super func(accepted bool), accepted bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGestureEvent_override_virtual_SetAccepted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGestureEvent_SetAccepted
func miqt_exec_callback_QGestureEvent_SetAccepted(self QGestureEvent, cb intptr_t, accepted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(accepted bool), accepted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(accepted)

	gofunc((&QGestureEvent{h: self}).callVirtualBase_SetAccepted, slotval1)

}

func (this *QGestureEvent) callVirtualBase_Clone() *QEvent {

	return newQEvent(QGestureEvent_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QGestureEvent) OnClone(slot func(super func() *QEvent) *QEvent) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGestureEvent_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGestureEvent_Clone
func miqt_exec_callback_QGestureEvent_Clone(self QGestureEvent, cb intptr_t) *QEvent {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QEvent) *QEvent)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGestureEvent{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}
