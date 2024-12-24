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
	g := newQGraphicsTransform(QGraphicsTransform_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsTransform2 constructs a new QGraphicsTransform object.
func NewQGraphicsTransform2(parent *QObject) *QGraphicsTransform {
	g := newQGraphicsTransform(QGraphicsTransform_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QGraphicsTransform) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsTransform_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsTransform) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_MetaObject
func miqt_exec_callback_QGraphicsTransform_MetaObject(self QGraphicsTransform, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsTransform{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsTransform) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsTransform_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsTransform) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTransform_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTransform_Metacast
func miqt_exec_callback_QGraphicsTransform_Metacast(self QGraphicsTransform, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGraphicsTransform{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QGraphicsScale struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsScale constructs a new QGraphicsScale object.
func NewQGraphicsScale() *QGraphicsScale {
	g := newQGraphicsScale(QGraphicsScale_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsScale2 constructs a new QGraphicsScale object.
func NewQGraphicsScale2(parent *QObject) *QGraphicsScale {
	g := newQGraphicsScale(QGraphicsScale_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QGraphicsScale) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsScale_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsScale) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsScale_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_MetaObject
func miqt_exec_callback_QGraphicsScale_MetaObject(self QGraphicsScale, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsScale{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsScale) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsScale_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsScale) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsScale_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScale_Metacast
func miqt_exec_callback_QGraphicsScale_Metacast(self QGraphicsScale, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGraphicsScale{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QGraphicsRotation struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsRotation constructs a new QGraphicsRotation object.
func NewQGraphicsRotation() *QGraphicsRotation {
	g := newQGraphicsRotation(QGraphicsRotation_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsRotation2 constructs a new QGraphicsRotation object.
func NewQGraphicsRotation2(parent *QObject) *QGraphicsRotation {
	g := newQGraphicsRotation(QGraphicsRotation_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QGraphicsRotation) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsRotation_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsRotation) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsRotation_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsRotation_MetaObject
func miqt_exec_callback_QGraphicsRotation_MetaObject(self QGraphicsRotation, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsRotation{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsRotation) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsRotation_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsRotation) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsRotation_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsRotation_Metacast
func miqt_exec_callback_QGraphicsRotation_Metacast(self QGraphicsRotation, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGraphicsRotation{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
