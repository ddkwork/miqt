package qt

import (
	"unsafe"
)

type QTimeLine__State int

const (
	QTimeLine__NotRunning QTimeLine__State = 0
	QTimeLine__Paused     QTimeLine__State = 1
	QTimeLine__Running    QTimeLine__State = 2
)

type QTimeLine__Direction int

const (
	QTimeLine__Forward  QTimeLine__Direction = 0
	QTimeLine__Backward QTimeLine__Direction = 1
)

type QTimeLine struct {
	h          uintptr
	isSubclass bool
}

// NewQTimeLine constructs a new QTimeLine object.
func NewQTimeLine() *QTimeLine {

	ret := newQTimeLine(QTimeLine_new())
	ret.isSubclass = true
	return ret
}

// NewQTimeLine2 constructs a new QTimeLine object.
func NewQTimeLine2(duration int) *QTimeLine {

	ret := newQTimeLine(QTimeLine_new2((int)(duration)))
	ret.isSubclass = true
	return ret
}

// NewQTimeLine3 constructs a new QTimeLine object.
func NewQTimeLine3(duration int, parent *QObject) *QTimeLine {

	ret := newQTimeLine(QTimeLine_new3((int)(duration), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTimeLine) MetaObject() *QMetaObject {
	return newQMetaObject(QTimeLine_MetaObject(this.h))
}

func (this *QTimeLine) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTimeLine_Metacast(this.h, param1_Cstring))
}

func QTimeLine_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTimeLine_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeLine) State() State {
	xxxxxxxxx
}

func (this *QTimeLine) LoopCount() int {
	return (int)(QTimeLine_LoopCount(this.h))
}

func (this *QTimeLine) SetLoopCount(count int) {
	QTimeLine_SetLoopCount(this.h, (int)(count))
}

func (this *QTimeLine) Direction() Direction {
	xxxxxxxxx
}

func (this *QTimeLine) SetDirection(direction Direction) {
	QTimeLine_SetDirection(this.h, direction)
}

func (this *QTimeLine) Duration() int {
	return (int)(QTimeLine_Duration(this.h))
}

func (this *QTimeLine) SetDuration(duration int) {
	QTimeLine_SetDuration(this.h, (int)(duration))
}

func (this *QTimeLine) StartFrame() int {
	return (int)(QTimeLine_StartFrame(this.h))
}

func (this *QTimeLine) SetStartFrame(frame int) {
	QTimeLine_SetStartFrame(this.h, (int)(frame))
}

func (this *QTimeLine) EndFrame() int {
	return (int)(QTimeLine_EndFrame(this.h))
}

func (this *QTimeLine) SetEndFrame(frame int) {
	QTimeLine_SetEndFrame(this.h, (int)(frame))
}

func (this *QTimeLine) SetFrameRange(startFrame int, endFrame int) {
	QTimeLine_SetFrameRange(this.h, (int)(startFrame), (int)(endFrame))
}

func (this *QTimeLine) UpdateInterval() int {
	return (int)(QTimeLine_UpdateInterval(this.h))
}

func (this *QTimeLine) SetUpdateInterval(interval int) {
	QTimeLine_SetUpdateInterval(this.h, (int)(interval))
}

func (this *QTimeLine) EasingCurve() *QEasingCurve {
	_goptr := newQEasingCurve(QTimeLine_EasingCurve(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTimeLine) SetEasingCurve(curve *QEasingCurve) {
	QTimeLine_SetEasingCurve(this.h, curve.cPointer())
}

func (this *QTimeLine) CurrentTime() int {
	return (int)(QTimeLine_CurrentTime(this.h))
}

func (this *QTimeLine) CurrentFrame() int {
	return (int)(QTimeLine_CurrentFrame(this.h))
}

func (this *QTimeLine) CurrentValue() float64 {
	return (float64)(QTimeLine_CurrentValue(this.h))
}

func (this *QTimeLine) FrameForTime(msec int) int {
	return (int)(QTimeLine_FrameForTime(this.h, (int)(msec)))
}

func (this *QTimeLine) ValueForTime(msec int) float64 {
	return (float64)(QTimeLine_ValueForTime(this.h, (int)(msec)))
}

func (this *QTimeLine) Start() {
	QTimeLine_Start(this.h)
}

func (this *QTimeLine) Resume() {
	QTimeLine_Resume(this.h)
}

func (this *QTimeLine) Stop() {
	QTimeLine_Stop(this.h)
}

func (this *QTimeLine) SetPaused(paused bool) {
	QTimeLine_SetPaused(this.h, (bool)(paused))
}

func (this *QTimeLine) SetCurrentTime(msec int) {
	QTimeLine_SetCurrentTime(this.h, (int)(msec))
}

func (this *QTimeLine) ToggleDirection() {
	QTimeLine_ToggleDirection(this.h)
}

func QTimeLine_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimeLine_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTimeLine_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimeLine_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeLine) callVirtualBase_ValueForTime(msec int) float64 {

	return (float64)(QTimeLine_virtualbase_ValueForTime(unsafe.Pointer(this.h), (int)(msec)))

}
func (this *QTimeLine) OnValueForTime(slot func(super func(msec int) float64, msec int) float64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_ValueForTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_ValueForTime
func miqt_exec_callback_QTimeLine_ValueForTime(self QTimeLine, cb intptr_t, msec int) double {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msec int) float64, msec int) float64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msec)

	virtualReturn := gofunc((&QTimeLine{h: self}).callVirtualBase_ValueForTime, slotval1)

	return (double)(virtualReturn)

}

func (this *QTimeLine) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QTimeLine_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeLine) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_TimerEvent
func miqt_exec_callback_QTimeLine_TimerEvent(self QTimeLine, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QTimeLine{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTimeLine) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QTimeLine_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTimeLine) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_Event
func miqt_exec_callback_QTimeLine_Event(self QTimeLine, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTimeLine{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTimeLine) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QTimeLine_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QTimeLine) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_EventFilter
func miqt_exec_callback_QTimeLine_EventFilter(self QTimeLine, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTimeLine{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QTimeLine) callVirtualBase_ChildEvent(event *QChildEvent) {

	QTimeLine_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeLine) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_ChildEvent
func miqt_exec_callback_QTimeLine_ChildEvent(self QTimeLine, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QTimeLine{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QTimeLine) callVirtualBase_CustomEvent(event *QEvent) {

	QTimeLine_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeLine) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_CustomEvent
func miqt_exec_callback_QTimeLine_CustomEvent(self QTimeLine, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTimeLine{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QTimeLine) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QTimeLine_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTimeLine) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_ConnectNotify
func miqt_exec_callback_QTimeLine_ConnectNotify(self QTimeLine, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTimeLine{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QTimeLine) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QTimeLine_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTimeLine) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeLine_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeLine_DisconnectNotify
func miqt_exec_callback_QTimeLine_DisconnectNotify(self QTimeLine, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTimeLine{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
