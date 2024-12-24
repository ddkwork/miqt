package qt

import (
	"unsafe"
)

type QGraphicsTransform struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsTransform constructs a new QGraphicsTransform object.
func NewQGraphicsTransform() *QGraphicsTransform {

	ret := newQGraphicsTransform(QGraphicsTransform_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsTransform2 constructs a new QGraphicsTransform object.
func NewQGraphicsTransform2(parent *QObject) *QGraphicsTransform {

	ret := newQGraphicsTransform(QGraphicsTransform_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsTransform) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsTransform_MetaObject(this.h))
}

func (this *QGraphicsTransform) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsTransform_Metacast(this.h, param1_Cstring))
}

func QGraphicsTransform_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsTransform_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsTransform) ApplyTo(matrix *QMatrix4x4) {
	QGraphicsTransform_ApplyTo(this.h, matrix.cPointer())
}

func QGraphicsTransform_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsTransform_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsTransform_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsTransform_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QGraphicsTransform) OnApplyTo(slot func(matrix *QMatrix4x4)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_ApplyTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_ApplyTo
func miqt_exec_callback_QGraphicsTransform_ApplyTo(self QGraphicsTransform, cb intptr_t, matrix *QMatrix4x4) {
	gofunc, ok := cgo.Handle(cb).Value().(func(matrix *QMatrix4x4))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMatrix4x4(matrix)

	gofunc(slotval1)

}

func (this *QGraphicsTransform) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGraphicsTransform_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsTransform) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_Event
func miqt_exec_callback_QGraphicsTransform_Event(self QGraphicsTransform, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsTransform{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsTransform) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QGraphicsTransform_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QGraphicsTransform) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_EventFilter
func miqt_exec_callback_QGraphicsTransform_EventFilter(self QGraphicsTransform, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsTransform{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QGraphicsTransform) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QGraphicsTransform_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsTransform) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_TimerEvent
func miqt_exec_callback_QGraphicsTransform_TimerEvent(self QGraphicsTransform, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QGraphicsTransform{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QGraphicsTransform) callVirtualBase_ChildEvent(event *QChildEvent) {

	QGraphicsTransform_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsTransform) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_ChildEvent
func miqt_exec_callback_QGraphicsTransform_ChildEvent(self QGraphicsTransform, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QGraphicsTransform{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QGraphicsTransform) callVirtualBase_CustomEvent(event *QEvent) {

	QGraphicsTransform_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsTransform) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_CustomEvent
func miqt_exec_callback_QGraphicsTransform_CustomEvent(self QGraphicsTransform, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsTransform{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QGraphicsTransform) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QGraphicsTransform_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGraphicsTransform) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_ConnectNotify
func miqt_exec_callback_QGraphicsTransform_ConnectNotify(self QGraphicsTransform, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGraphicsTransform{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QGraphicsTransform) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QGraphicsTransform_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGraphicsTransform) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_DisconnectNotify
func miqt_exec_callback_QGraphicsTransform_DisconnectNotify(self QGraphicsTransform, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGraphicsTransform{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

type QGraphicsScale struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsScale constructs a new QGraphicsScale object.
func NewQGraphicsScale() *QGraphicsScale {

	ret := newQGraphicsScale(QGraphicsScale_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsScale2 constructs a new QGraphicsScale object.
func NewQGraphicsScale2(parent *QObject) *QGraphicsScale {

	ret := newQGraphicsScale(QGraphicsScale_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsScale) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsScale_MetaObject(this.h))
}

func (this *QGraphicsScale) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsScale_Metacast(this.h, param1_Cstring))
}

func QGraphicsScale_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsScale_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsScale) Origin() *QVector3D {
	_goptr := newQVector3D(QGraphicsScale_Origin(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScale) SetOrigin(point *QVector3D) {
	QGraphicsScale_SetOrigin(this.h, point.cPointer())
}

func (this *QGraphicsScale) XScale() float64 {
	return (float64)(QGraphicsScale_XScale(this.h))
}

func (this *QGraphicsScale) SetXScale(xScale float64) {
	QGraphicsScale_SetXScale(this.h, (double)(xScale))
}

func (this *QGraphicsScale) YScale() float64 {
	return (float64)(QGraphicsScale_YScale(this.h))
}

func (this *QGraphicsScale) SetYScale(yScale float64) {
	QGraphicsScale_SetYScale(this.h, (double)(yScale))
}

func (this *QGraphicsScale) ZScale() float64 {
	return (float64)(QGraphicsScale_ZScale(this.h))
}

func (this *QGraphicsScale) SetZScale(zScale float64) {
	QGraphicsScale_SetZScale(this.h, (double)(zScale))
}

func (this *QGraphicsScale) ApplyTo(matrix *QMatrix4x4) {
	QGraphicsScale_ApplyTo(this.h, matrix.cPointer())
}

func (this *QGraphicsScale) OriginChanged() {
	QGraphicsScale_OriginChanged(this.h)
}
func (this *QGraphicsScale) OnOriginChanged(slot func()) {
	QGraphicsScale_connect_OriginChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_OriginChanged
func miqt_exec_callback_QGraphicsScale_OriginChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsScale) XScaleChanged() {
	QGraphicsScale_XScaleChanged(this.h)
}
func (this *QGraphicsScale) OnXScaleChanged(slot func()) {
	QGraphicsScale_connect_XScaleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_XScaleChanged
func miqt_exec_callback_QGraphicsScale_XScaleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsScale) YScaleChanged() {
	QGraphicsScale_YScaleChanged(this.h)
}
func (this *QGraphicsScale) OnYScaleChanged(slot func()) {
	QGraphicsScale_connect_YScaleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_YScaleChanged
func miqt_exec_callback_QGraphicsScale_YScaleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsScale) ZScaleChanged() {
	QGraphicsScale_ZScaleChanged(this.h)
}
func (this *QGraphicsScale) OnZScaleChanged(slot func()) {
	QGraphicsScale_connect_ZScaleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_ZScaleChanged
func miqt_exec_callback_QGraphicsScale_ZScaleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsScale) ScaleChanged() {
	QGraphicsScale_ScaleChanged(this.h)
}
func (this *QGraphicsScale) OnScaleChanged(slot func()) {
	QGraphicsScale_connect_ScaleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_ScaleChanged
func miqt_exec_callback_QGraphicsScale_ScaleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QGraphicsScale_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsScale_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsScale_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsScale_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsScale) callVirtualBase_ApplyTo(matrix *QMatrix4x4) {

	QGraphicsScale_virtualbase_ApplyTo(unsafe.Pointer(this.h), matrix.cPointer())

}
func (this *QGraphicsScale) OnApplyTo(slot func(super func(matrix *QMatrix4x4), matrix *QMatrix4x4)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsScale_override_virtual_ApplyTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_ApplyTo
func miqt_exec_callback_QGraphicsScale_ApplyTo(self QGraphicsScale, cb intptr_t, matrix *QMatrix4x4) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(matrix *QMatrix4x4), matrix *QMatrix4x4))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMatrix4x4(matrix)

	gofunc((&QGraphicsScale{h: self}).callVirtualBase_ApplyTo, slotval1)

}

type QGraphicsRotation struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsRotation constructs a new QGraphicsRotation object.
func NewQGraphicsRotation() *QGraphicsRotation {

	ret := newQGraphicsRotation(QGraphicsRotation_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsRotation2 constructs a new QGraphicsRotation object.
func NewQGraphicsRotation2(parent *QObject) *QGraphicsRotation {

	ret := newQGraphicsRotation(QGraphicsRotation_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsRotation) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsRotation_MetaObject(this.h))
}

func (this *QGraphicsRotation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsRotation_Metacast(this.h, param1_Cstring))
}

func QGraphicsRotation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsRotation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsRotation) Origin() *QVector3D {
	_goptr := newQVector3D(QGraphicsRotation_Origin(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsRotation) SetOrigin(point *QVector3D) {
	QGraphicsRotation_SetOrigin(this.h, point.cPointer())
}

func (this *QGraphicsRotation) Angle() float64 {
	return (float64)(QGraphicsRotation_Angle(this.h))
}

func (this *QGraphicsRotation) SetAngle(angle float64) {
	QGraphicsRotation_SetAngle(this.h, (double)(angle))
}

func (this *QGraphicsRotation) Axis() *QVector3D {
	_goptr := newQVector3D(QGraphicsRotation_Axis(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsRotation) SetAxis(axis *QVector3D) {
	QGraphicsRotation_SetAxis(this.h, axis.cPointer())
}

func (this *QGraphicsRotation) SetAxisWithAxis(axis Axis) {
	QGraphicsRotation_SetAxisWithAxis(this.h, (int)(axis))
}

func (this *QGraphicsRotation) ApplyTo(matrix *QMatrix4x4) {
	QGraphicsRotation_ApplyTo(this.h, matrix.cPointer())
}

func (this *QGraphicsRotation) OriginChanged() {
	QGraphicsRotation_OriginChanged(this.h)
}
func (this *QGraphicsRotation) OnOriginChanged(slot func()) {
	QGraphicsRotation_connect_OriginChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsRotation_OriginChanged
func miqt_exec_callback_QGraphicsRotation_OriginChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsRotation) AngleChanged() {
	QGraphicsRotation_AngleChanged(this.h)
}
func (this *QGraphicsRotation) OnAngleChanged(slot func()) {
	QGraphicsRotation_connect_AngleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsRotation_AngleChanged
func miqt_exec_callback_QGraphicsRotation_AngleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsRotation) AxisChanged() {
	QGraphicsRotation_AxisChanged(this.h)
}
func (this *QGraphicsRotation) OnAxisChanged(slot func()) {
	QGraphicsRotation_connect_AxisChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsRotation_AxisChanged
func miqt_exec_callback_QGraphicsRotation_AxisChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QGraphicsRotation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsRotation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsRotation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsRotation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsRotation) callVirtualBase_ApplyTo(matrix *QMatrix4x4) {

	QGraphicsRotation_virtualbase_ApplyTo(unsafe.Pointer(this.h), matrix.cPointer())

}
func (this *QGraphicsRotation) OnApplyTo(slot func(super func(matrix *QMatrix4x4), matrix *QMatrix4x4)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsRotation_override_virtual_ApplyTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsRotation_ApplyTo
func miqt_exec_callback_QGraphicsRotation_ApplyTo(self QGraphicsRotation, cb intptr_t, matrix *QMatrix4x4) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(matrix *QMatrix4x4), matrix *QMatrix4x4))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMatrix4x4(matrix)

	gofunc((&QGraphicsRotation{h: self}).callVirtualBase_ApplyTo, slotval1)

}
