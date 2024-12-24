package qt

import (
	"unsafe"
)

type QScroller__State int

const (
	QScroller__Inactive  QScroller__State = 0
	QScroller__Pressed   QScroller__State = 1
	QScroller__Dragging  QScroller__State = 2
	QScroller__Scrolling QScroller__State = 3
)

type QScroller__ScrollerGestureType int

const (
	QScroller__TouchGesture             QScroller__ScrollerGestureType = 0
	QScroller__LeftMouseButtonGesture   QScroller__ScrollerGestureType = 1
	QScroller__RightMouseButtonGesture  QScroller__ScrollerGestureType = 2
	QScroller__MiddleMouseButtonGesture QScroller__ScrollerGestureType = 3
)

type QScroller__Input int

const (
	QScroller__InputPress   QScroller__Input = 1
	QScroller__InputMove    QScroller__Input = 2
	QScroller__InputRelease QScroller__Input = 3
)

type QScroller struct {
	h          uintptr
	isSubclass bool
}

func (this *QScroller) MetaObject() *QMetaObject {
	return newQMetaObject(QScroller_MetaObject(this.h))
}

func (this *QScroller) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QScroller_Metacast(this.h, param1_Cstring))
}

func QScroller_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QScroller_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScroller_HasScroller(target *QObject) bool {
	return (bool)(QScroller_HasScroller(target.cPointer()))
}

func QScroller_Scroller(target *QObject) *QScroller {
	return newQScroller(QScroller_Scroller(target.cPointer()))
}

func QScroller_ScrollerWithTarget(target *QObject) *QScroller {
	return newQScroller(QScroller_ScrollerWithTarget(target.cPointer()))
}

func QScroller_GrabGesture(target *QObject) GestureType {
	return (GestureType)(QScroller_GrabGesture(target.cPointer()))
}

func QScroller_GrabbedGesture(target *QObject) GestureType {
	return (GestureType)(QScroller_GrabbedGesture(target.cPointer()))
}

func QScroller_UngrabGesture(target *QObject) {
	QScroller_UngrabGesture(target.cPointer())
}

func QScroller_ActiveScrollers() []*QScroller {
	var _ma struct_miqt_array = QScroller_ActiveScrollers()
	_ret := make([]*QScroller, int(_ma.len))
	_outCast := (*[0xffff]*QScroller)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQScroller(_outCast[i])
	}
	return _ret
}

func (this *QScroller) Target() *QObject {
	return newQObject(QScroller_Target(this.h))
}

func (this *QScroller) State() State {
	xxxxxxxxx
}

func (this *QScroller) HandleInput(input Input, position *QPointF) bool {
	return (bool)(QScroller_HandleInput(this.h, input, position.cPointer()))
}

func (this *QScroller) Stop() {
	QScroller_Stop(this.h)
}

func (this *QScroller) Velocity() *QPointF {
	_goptr := newQPointF(QScroller_Velocity(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScroller) FinalPosition() *QPointF {
	_goptr := newQPointF(QScroller_FinalPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScroller) PixelPerMeter() *QPointF {
	_goptr := newQPointF(QScroller_PixelPerMeter(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScroller) ScrollerProperties() *QScrollerProperties {
	_goptr := newQScrollerProperties(QScroller_ScrollerProperties(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScroller) SetSnapPositionsX(positions []float64) {
	positions_CArray := (*[0xffff]double)(malloc(size_t(8 * len(positions))))
	defer free(unsafe.Pointer(positions_CArray))
	for i := range positions {
		positions_CArray[i] = (double)(positions[i])
	}
	positions_ma := struct_miqt_array{len: size_t(len(positions)), data: unsafe.Pointer(positions_CArray)}
	QScroller_SetSnapPositionsX(this.h, positions_ma)
}

func (this *QScroller) SetSnapPositionsX2(first float64, interval float64) {
	QScroller_SetSnapPositionsX2(this.h, (double)(first), (double)(interval))
}

func (this *QScroller) SetSnapPositionsY(positions []float64) {
	positions_CArray := (*[0xffff]double)(malloc(size_t(8 * len(positions))))
	defer free(unsafe.Pointer(positions_CArray))
	for i := range positions {
		positions_CArray[i] = (double)(positions[i])
	}
	positions_ma := struct_miqt_array{len: size_t(len(positions)), data: unsafe.Pointer(positions_CArray)}
	QScroller_SetSnapPositionsY(this.h, positions_ma)
}

func (this *QScroller) SetSnapPositionsY2(first float64, interval float64) {
	QScroller_SetSnapPositionsY2(this.h, (double)(first), (double)(interval))
}

func (this *QScroller) SetScrollerProperties(prop *QScrollerProperties) {
	QScroller_SetScrollerProperties(this.h, prop.cPointer())
}

func (this *QScroller) ScrollTo(pos *QPointF) {
	QScroller_ScrollTo(this.h, pos.cPointer())
}

func (this *QScroller) ScrollTo2(pos *QPointF, scrollTime int) {
	QScroller_ScrollTo2(this.h, pos.cPointer(), (int)(scrollTime))
}

func (this *QScroller) EnsureVisible(rect *QRectF, xmargin float64, ymargin float64) {
	QScroller_EnsureVisible(this.h, rect.cPointer(), (double)(xmargin), (double)(ymargin))
}

func (this *QScroller) EnsureVisible2(rect *QRectF, xmargin float64, ymargin float64, scrollTime int) {
	QScroller_EnsureVisible2(this.h, rect.cPointer(), (double)(xmargin), (double)(ymargin), (int)(scrollTime))
}

func (this *QScroller) ResendPrepareEvent() {
	QScroller_ResendPrepareEvent(this.h)
}

func (this *QScroller) StateChanged(newstate QScroller__State) {
	QScroller_StateChanged(this.h, (int)(newstate))
}
func (this *QScroller) OnStateChanged(slot func(newstate QScroller__State)) {
	QScroller_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScroller_StateChanged
func miqt_exec_callback_QScroller_StateChanged(cb intptr_t, newstate int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newstate QScroller__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QScroller__State)(newstate)

	gofunc(slotval1)
}

func (this *QScroller) ScrollerPropertiesChanged(param1 *QScrollerProperties) {
	QScroller_ScrollerPropertiesChanged(this.h, param1.cPointer())
}
func (this *QScroller) OnScrollerPropertiesChanged(slot func(param1 *QScrollerProperties)) {
	QScroller_connect_ScrollerPropertiesChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScroller_ScrollerPropertiesChanged
func miqt_exec_callback_QScroller_ScrollerPropertiesChanged(cb intptr_t, param1 *QScrollerProperties) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QScrollerProperties))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQScrollerProperties(param1)

	gofunc(slotval1)
}

func QScroller_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScroller_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScroller_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScroller_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScroller_GrabGesture2(target *QObject, gestureType ScrollerGestureType) GestureType {
	return (GestureType)(QScroller_GrabGesture2(target.cPointer(), gestureType))
}

func (this *QScroller) HandleInput3(input Input, position *QPointF, timestamp int64) bool {
	return (bool)(QScroller_HandleInput3(this.h, input, position.cPointer(), (longlong)(timestamp)))
}
