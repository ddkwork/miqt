package qt

import (
	"unsafe"
)

type QScreen struct {
	h          uintptr
	isSubclass bool
}

func (this *QScreen) MetaObject() *QMetaObject {
	return newQMetaObject(QScreen_MetaObject(this.h))
}

func (this *QScreen) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QScreen_Metacast(this.h, param1_Cstring))
}

func QScreen_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QScreen_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreen) Name() string {
	var _ms struct_miqt_string = QScreen_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreen) Manufacturer() string {
	var _ms struct_miqt_string = QScreen_Manufacturer(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreen) Model() string {
	var _ms struct_miqt_string = QScreen_Model(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreen) SerialNumber() string {
	var _ms struct_miqt_string = QScreen_SerialNumber(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreen) Depth() int {
	return (int)(QScreen_Depth(this.h))
}

func (this *QScreen) Size() *QSize {
	_goptr := newQSize(QScreen_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) Geometry() *QRect {
	_goptr := newQRect(QScreen_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) PhysicalSize() *QSizeF {
	_goptr := newQSizeF(QScreen_PhysicalSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) PhysicalDotsPerInchX() float64 {
	return (float64)(QScreen_PhysicalDotsPerInchX(this.h))
}

func (this *QScreen) PhysicalDotsPerInchY() float64 {
	return (float64)(QScreen_PhysicalDotsPerInchY(this.h))
}

func (this *QScreen) PhysicalDotsPerInch() float64 {
	return (float64)(QScreen_PhysicalDotsPerInch(this.h))
}

func (this *QScreen) LogicalDotsPerInchX() float64 {
	return (float64)(QScreen_LogicalDotsPerInchX(this.h))
}

func (this *QScreen) LogicalDotsPerInchY() float64 {
	return (float64)(QScreen_LogicalDotsPerInchY(this.h))
}

func (this *QScreen) LogicalDotsPerInch() float64 {
	return (float64)(QScreen_LogicalDotsPerInch(this.h))
}

func (this *QScreen) DevicePixelRatio() float64 {
	return (float64)(QScreen_DevicePixelRatio(this.h))
}

func (this *QScreen) AvailableSize() *QSize {
	_goptr := newQSize(QScreen_AvailableSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) AvailableGeometry() *QRect {
	_goptr := newQRect(QScreen_AvailableGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) VirtualSiblings() []*QScreen {
	var _ma struct_miqt_array = QScreen_VirtualSiblings(this.h)
	_ret := make([]*QScreen, int(_ma.len))
	_outCast := (*[0xffff]*QScreen)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQScreen(_outCast[i])
	}
	return _ret
}

func (this *QScreen) VirtualSiblingAt(point QPoint) *QScreen {
	return newQScreen(QScreen_VirtualSiblingAt(this.h, point.cPointer()))
}

func (this *QScreen) VirtualSize() *QSize {
	_goptr := newQSize(QScreen_VirtualSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) VirtualGeometry() *QRect {
	_goptr := newQRect(QScreen_VirtualGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) AvailableVirtualSize() *QSize {
	_goptr := newQSize(QScreen_AvailableVirtualSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) AvailableVirtualGeometry() *QRect {
	_goptr := newQRect(QScreen_AvailableVirtualGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) PrimaryOrientation() ScreenOrientation {
	return (ScreenOrientation)(QScreen_PrimaryOrientation(this.h))
}

func (this *QScreen) Orientation() ScreenOrientation {
	return (ScreenOrientation)(QScreen_Orientation(this.h))
}

func (this *QScreen) NativeOrientation() ScreenOrientation {
	return (ScreenOrientation)(QScreen_NativeOrientation(this.h))
}

func (this *QScreen) AngleBetween(a ScreenOrientation, b ScreenOrientation) int {
	return (int)(QScreen_AngleBetween(this.h, (int)(a), (int)(b)))
}

func (this *QScreen) TransformBetween(a ScreenOrientation, b ScreenOrientation, target *QRect) *QTransform {
	_goptr := newQTransform(QScreen_TransformBetween(this.h, (int)(a), (int)(b), target.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) MapBetween(a ScreenOrientation, b ScreenOrientation, rect *QRect) *QRect {
	_goptr := newQRect(QScreen_MapBetween(this.h, (int)(a), (int)(b), rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) IsPortrait(orientation ScreenOrientation) bool {
	return (bool)(QScreen_IsPortrait(this.h, (int)(orientation)))
}

func (this *QScreen) IsLandscape(orientation ScreenOrientation) bool {
	return (bool)(QScreen_IsLandscape(this.h, (int)(orientation)))
}

func (this *QScreen) GrabWindow() *QPixmap {
	_goptr := newQPixmap(QScreen_GrabWindow(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) RefreshRate() float64 {
	return (float64)(QScreen_RefreshRate(this.h))
}

func (this *QScreen) GeometryChanged(geometry *QRect) {
	QScreen_GeometryChanged(this.h, geometry.cPointer())
}
func (this *QScreen) OnGeometryChanged(slot func(geometry *QRect)) {
	QScreen_connect_GeometryChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_GeometryChanged
func miqt_exec_callback_QScreen_GeometryChanged(cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc(slotval1)
}

func (this *QScreen) AvailableGeometryChanged(geometry *QRect) {
	QScreen_AvailableGeometryChanged(this.h, geometry.cPointer())
}
func (this *QScreen) OnAvailableGeometryChanged(slot func(geometry *QRect)) {
	QScreen_connect_AvailableGeometryChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_AvailableGeometryChanged
func miqt_exec_callback_QScreen_AvailableGeometryChanged(cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc(slotval1)
}

func (this *QScreen) PhysicalSizeChanged(size *QSizeF) {
	QScreen_PhysicalSizeChanged(this.h, size.cPointer())
}
func (this *QScreen) OnPhysicalSizeChanged(slot func(size *QSizeF)) {
	QScreen_connect_PhysicalSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_PhysicalSizeChanged
func miqt_exec_callback_QScreen_PhysicalSizeChanged(cb intptr_t, size *QSizeF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(size *QSizeF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSizeF(size)

	gofunc(slotval1)
}

func (this *QScreen) PhysicalDotsPerInchChanged(dpi float64) {
	QScreen_PhysicalDotsPerInchChanged(this.h, (double)(dpi))
}
func (this *QScreen) OnPhysicalDotsPerInchChanged(slot func(dpi float64)) {
	QScreen_connect_PhysicalDotsPerInchChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_PhysicalDotsPerInchChanged
func miqt_exec_callback_QScreen_PhysicalDotsPerInchChanged(cb intptr_t, dpi double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(dpi float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(dpi)

	gofunc(slotval1)
}

func (this *QScreen) LogicalDotsPerInchChanged(dpi float64) {
	QScreen_LogicalDotsPerInchChanged(this.h, (double)(dpi))
}
func (this *QScreen) OnLogicalDotsPerInchChanged(slot func(dpi float64)) {
	QScreen_connect_LogicalDotsPerInchChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_LogicalDotsPerInchChanged
func miqt_exec_callback_QScreen_LogicalDotsPerInchChanged(cb intptr_t, dpi double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(dpi float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(dpi)

	gofunc(slotval1)
}

func (this *QScreen) VirtualGeometryChanged(rect *QRect) {
	QScreen_VirtualGeometryChanged(this.h, rect.cPointer())
}
func (this *QScreen) OnVirtualGeometryChanged(slot func(rect *QRect)) {
	QScreen_connect_VirtualGeometryChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_VirtualGeometryChanged
func miqt_exec_callback_QScreen_VirtualGeometryChanged(cb intptr_t, rect *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(rect *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	gofunc(slotval1)
}

func (this *QScreen) PrimaryOrientationChanged(orientation ScreenOrientation) {
	QScreen_PrimaryOrientationChanged(this.h, (int)(orientation))
}
func (this *QScreen) OnPrimaryOrientationChanged(slot func(orientation ScreenOrientation)) {
	QScreen_connect_PrimaryOrientationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_PrimaryOrientationChanged
func miqt_exec_callback_QScreen_PrimaryOrientationChanged(cb intptr_t, orientation int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(orientation ScreenOrientation))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ScreenOrientation)(orientation)

	gofunc(slotval1)
}

func (this *QScreen) OrientationChanged(orientation ScreenOrientation) {
	QScreen_OrientationChanged(this.h, (int)(orientation))
}
func (this *QScreen) OnOrientationChanged(slot func(orientation ScreenOrientation)) {
	QScreen_connect_OrientationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_OrientationChanged
func miqt_exec_callback_QScreen_OrientationChanged(cb intptr_t, orientation int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(orientation ScreenOrientation))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ScreenOrientation)(orientation)

	gofunc(slotval1)
}

func (this *QScreen) RefreshRateChanged(refreshRate float64) {
	QScreen_RefreshRateChanged(this.h, (double)(refreshRate))
}
func (this *QScreen) OnRefreshRateChanged(slot func(refreshRate float64)) {
	QScreen_connect_RefreshRateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreen_RefreshRateChanged
func miqt_exec_callback_QScreen_RefreshRateChanged(cb intptr_t, refreshRate double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(refreshRate float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(refreshRate)

	gofunc(slotval1)
}

func QScreen_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScreen_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScreen_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScreen_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreen) GrabWindow1(window uintptr) *QPixmap {
	_goptr := newQPixmap(QScreen_GrabWindow1(this.h, (uintptr_t)(window)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) GrabWindow2(window uintptr, x int) *QPixmap {
	_goptr := newQPixmap(QScreen_GrabWindow2(this.h, (uintptr_t)(window), (int)(x)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) GrabWindow3(window uintptr, x int, y int) *QPixmap {
	_goptr := newQPixmap(QScreen_GrabWindow3(this.h, (uintptr_t)(window), (int)(x), (int)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) GrabWindow4(window uintptr, x int, y int, w int) *QPixmap {
	_goptr := newQPixmap(QScreen_GrabWindow4(this.h, (uintptr_t)(window), (int)(x), (int)(y), (int)(w)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScreen) GrabWindow5(window uintptr, x int, y int, w int, h int) *QPixmap {
	_goptr := newQPixmap(QScreen_GrabWindow5(this.h, (uintptr_t)(window), (int)(x), (int)(y), (int)(w), (int)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
