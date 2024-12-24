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
	g := newQGesture(QGesture_new())
	g.isSubclass = true
	return g
}

// NewQGesture2 constructs a new QGesture object.
func NewQGesture2(parent *QObject) *QGesture {
	g := newQGesture(QGesture_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QGesture) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGesture_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGesture) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_MetaObject
func miqt_exec_callback_QGesture_MetaObject(self QGesture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGesture{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGesture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGesture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGesture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGesture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGesture_Metacast
func miqt_exec_callback_QGesture_Metacast(self QGesture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGesture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QPanGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQPanGesture constructs a new QPanGesture object.
func NewQPanGesture() *QPanGesture {
	g := newQPanGesture(QPanGesture_new())
	g.isSubclass = true
	return g
}

// NewQPanGesture2 constructs a new QPanGesture object.
func NewQPanGesture2(parent *QObject) *QPanGesture {
	g := newQPanGesture(QPanGesture_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QPanGesture) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPanGesture_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPanGesture) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPanGesture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPanGesture_MetaObject
func miqt_exec_callback_QPanGesture_MetaObject(self QPanGesture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPanGesture{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPanGesture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPanGesture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPanGesture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPanGesture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPanGesture_Metacast
func miqt_exec_callback_QPanGesture_Metacast(self QPanGesture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QPanGesture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QPinchGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQPinchGesture constructs a new QPinchGesture object.
func NewQPinchGesture() *QPinchGesture {
	g := newQPinchGesture(QPinchGesture_new())
	g.isSubclass = true
	return g
}

// NewQPinchGesture2 constructs a new QPinchGesture object.
func NewQPinchGesture2(parent *QObject) *QPinchGesture {
	g := newQPinchGesture(QPinchGesture_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QPinchGesture) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPinchGesture_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPinchGesture) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPinchGesture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPinchGesture_MetaObject
func miqt_exec_callback_QPinchGesture_MetaObject(self QPinchGesture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPinchGesture{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPinchGesture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPinchGesture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPinchGesture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPinchGesture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPinchGesture_Metacast
func miqt_exec_callback_QPinchGesture_Metacast(self QPinchGesture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QPinchGesture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QSwipeGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQSwipeGesture constructs a new QSwipeGesture object.
func NewQSwipeGesture() *QSwipeGesture {
	g := newQSwipeGesture(QSwipeGesture_new())
	g.isSubclass = true
	return g
}

// NewQSwipeGesture2 constructs a new QSwipeGesture object.
func NewQSwipeGesture2(parent *QObject) *QSwipeGesture {
	g := newQSwipeGesture(QSwipeGesture_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QSwipeGesture) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSwipeGesture_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSwipeGesture) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSwipeGesture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSwipeGesture_MetaObject
func miqt_exec_callback_QSwipeGesture_MetaObject(self QSwipeGesture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSwipeGesture{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSwipeGesture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSwipeGesture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSwipeGesture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSwipeGesture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSwipeGesture_Metacast
func miqt_exec_callback_QSwipeGesture_Metacast(self QSwipeGesture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSwipeGesture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QTapGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQTapGesture constructs a new QTapGesture object.
func NewQTapGesture() *QTapGesture {
	g := newQTapGesture(QTapGesture_new())
	g.isSubclass = true
	return g
}

// NewQTapGesture2 constructs a new QTapGesture object.
func NewQTapGesture2(parent *QObject) *QTapGesture {
	g := newQTapGesture(QTapGesture_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QTapGesture) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTapGesture_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTapGesture) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTapGesture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTapGesture_MetaObject
func miqt_exec_callback_QTapGesture_MetaObject(self QTapGesture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTapGesture{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTapGesture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTapGesture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTapGesture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTapGesture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTapGesture_Metacast
func miqt_exec_callback_QTapGesture_Metacast(self QTapGesture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTapGesture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QTapAndHoldGesture struct {
	h          uintptr
	isSubclass bool
}

// NewQTapAndHoldGesture constructs a new QTapAndHoldGesture object.
func NewQTapAndHoldGesture() *QTapAndHoldGesture {
	g := newQTapAndHoldGesture(QTapAndHoldGesture_new())
	g.isSubclass = true
	return g
}

// NewQTapAndHoldGesture2 constructs a new QTapAndHoldGesture object.
func NewQTapAndHoldGesture2(parent *QObject) *QTapAndHoldGesture {
	g := newQTapAndHoldGesture(QTapAndHoldGesture_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QTapAndHoldGesture) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTapAndHoldGesture_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTapAndHoldGesture) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTapAndHoldGesture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTapAndHoldGesture_MetaObject
func miqt_exec_callback_QTapAndHoldGesture_MetaObject(self QTapAndHoldGesture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTapAndHoldGesture{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTapAndHoldGesture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTapAndHoldGesture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTapAndHoldGesture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTapAndHoldGesture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTapAndHoldGesture_Metacast
func miqt_exec_callback_QTapAndHoldGesture_Metacast(self QTapAndHoldGesture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTapAndHoldGesture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
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
	g := newQGestureEvent(QGestureEvent_new(gestures_ma))
	g.isSubclass = true
	return g
}

// NewQGestureEvent2 constructs a new QGestureEvent object.
func NewQGestureEvent2(param1 *QGestureEvent) *QGestureEvent {
	g := newQGestureEvent(QGestureEvent_new2(param1.cPointer()))
	g.isSubclass = true
	return g
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
