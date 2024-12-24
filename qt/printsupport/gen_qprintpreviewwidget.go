package printsupport

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QPrintPreviewWidget__ViewMode int

const (
	QPrintPreviewWidget__SinglePageView  QPrintPreviewWidget__ViewMode = 0
	QPrintPreviewWidget__FacingPagesView QPrintPreviewWidget__ViewMode = 1
	QPrintPreviewWidget__AllPagesView    QPrintPreviewWidget__ViewMode = 2
)

type QPrintPreviewWidget__ZoomMode int

const (
	QPrintPreviewWidget__CustomZoom QPrintPreviewWidget__ZoomMode = 0
	QPrintPreviewWidget__FitToWidth QPrintPreviewWidget__ZoomMode = 1
	QPrintPreviewWidget__FitInView  QPrintPreviewWidget__ZoomMode = 2
)

type QPrintPreviewWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQPrintPreviewWidget constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget(parent *qt.QWidget) *QPrintPreviewWidget {

	ret := newQPrintPreviewWidget(QPrintPreviewWidget_new((*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewWidget2 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget2(printer *QPrinter) *QPrintPreviewWidget {

	ret := newQPrintPreviewWidget(QPrintPreviewWidget_new2(printer.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewWidget3 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget3() *QPrintPreviewWidget {

	ret := newQPrintPreviewWidget(QPrintPreviewWidget_new3())
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewWidget4 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget4(printer *QPrinter, parent *qt.QWidget) *QPrintPreviewWidget {

	ret := newQPrintPreviewWidget(QPrintPreviewWidget_new4(printer.cPointer(), (*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewWidget5 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget5(printer *QPrinter, parent *qt.QWidget, flags WindowType) *QPrintPreviewWidget {

	ret := newQPrintPreviewWidget(QPrintPreviewWidget_new5(printer.cPointer(), (*QWidget)(parent.UnsafePointer()), (int)(flags)))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewWidget6 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget6(parent *qt.QWidget, flags WindowType) *QPrintPreviewWidget {

	ret := newQPrintPreviewWidget(QPrintPreviewWidget_new6((*QWidget)(parent.UnsafePointer()), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QPrintPreviewWidget) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QPrintPreviewWidget_MetaObject(this.h)))
}

func (this *QPrintPreviewWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPrintPreviewWidget_Metacast(this.h, param1_Cstring))
}

func QPrintPreviewWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPrintPreviewWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrintPreviewWidget) ZoomFactor() float64 {
	return (float64)(QPrintPreviewWidget_ZoomFactor(this.h))
}

func (this *QPrintPreviewWidget) Orientation() qt.QPageLayout__Orientation {
	return (qt.QPageLayout__Orientation)(QPrintPreviewWidget_Orientation(this.h))
}

func (this *QPrintPreviewWidget) ViewMode() ViewMode {
	xxxxxxxxx
}

func (this *QPrintPreviewWidget) ZoomMode() ZoomMode {
	xxxxxxxxx
}

func (this *QPrintPreviewWidget) CurrentPage() int {
	return (int)(QPrintPreviewWidget_CurrentPage(this.h))
}

func (this *QPrintPreviewWidget) PageCount() int {
	return (int)(QPrintPreviewWidget_PageCount(this.h))
}

func (this *QPrintPreviewWidget) SetVisible(visible bool) {
	QPrintPreviewWidget_SetVisible(this.h, (bool)(visible))
}

func (this *QPrintPreviewWidget) Print() {
	QPrintPreviewWidget_Print(this.h)
}

func (this *QPrintPreviewWidget) ZoomIn() {
	QPrintPreviewWidget_ZoomIn(this.h)
}

func (this *QPrintPreviewWidget) ZoomOut() {
	QPrintPreviewWidget_ZoomOut(this.h)
}

func (this *QPrintPreviewWidget) SetZoomFactor(zoomFactor float64) {
	QPrintPreviewWidget_SetZoomFactor(this.h, (double)(zoomFactor))
}

func (this *QPrintPreviewWidget) SetOrientation(orientation qt.QPageLayout__Orientation) {
	QPrintPreviewWidget_SetOrientation(this.h, (int)(orientation))
}

func (this *QPrintPreviewWidget) SetViewMode(viewMode ViewMode) {
	QPrintPreviewWidget_SetViewMode(this.h, viewMode)
}

func (this *QPrintPreviewWidget) SetZoomMode(zoomMode ZoomMode) {
	QPrintPreviewWidget_SetZoomMode(this.h, zoomMode)
}

func (this *QPrintPreviewWidget) SetCurrentPage(pageNumber int) {
	QPrintPreviewWidget_SetCurrentPage(this.h, (int)(pageNumber))
}

func (this *QPrintPreviewWidget) FitToWidth() {
	QPrintPreviewWidget_FitToWidth(this.h)
}

func (this *QPrintPreviewWidget) FitInView() {
	QPrintPreviewWidget_FitInView(this.h)
}

func (this *QPrintPreviewWidget) SetLandscapeOrientation() {
	QPrintPreviewWidget_SetLandscapeOrientation(this.h)
}

func (this *QPrintPreviewWidget) SetPortraitOrientation() {
	QPrintPreviewWidget_SetPortraitOrientation(this.h)
}

func (this *QPrintPreviewWidget) SetSinglePageViewMode() {
	QPrintPreviewWidget_SetSinglePageViewMode(this.h)
}

func (this *QPrintPreviewWidget) SetFacingPagesViewMode() {
	QPrintPreviewWidget_SetFacingPagesViewMode(this.h)
}

func (this *QPrintPreviewWidget) SetAllPagesViewMode() {
	QPrintPreviewWidget_SetAllPagesViewMode(this.h)
}

func (this *QPrintPreviewWidget) UpdatePreview() {
	QPrintPreviewWidget_UpdatePreview(this.h)
}

func (this *QPrintPreviewWidget) PaintRequested(printer *QPrinter) {
	QPrintPreviewWidget_PaintRequested(this.h, printer.cPointer())
}
func (this *QPrintPreviewWidget) OnPaintRequested(slot func(printer *QPrinter)) {
	QPrintPreviewWidget_connect_PaintRequested(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_PaintRequested
func miqt_exec_callback_QPrintPreviewWidget_PaintRequested(cb intptr_t, printer *QPrinter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(printer *QPrinter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPrinter(printer)

	gofunc(slotval1)
}

func (this *QPrintPreviewWidget) PreviewChanged() {
	QPrintPreviewWidget_PreviewChanged(this.h)
}
func (this *QPrintPreviewWidget) OnPreviewChanged(slot func()) {
	QPrintPreviewWidget_connect_PreviewChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_PreviewChanged
func miqt_exec_callback_QPrintPreviewWidget_PreviewChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QPrintPreviewWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPrintPreviewWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPrintPreviewWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPrintPreviewWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrintPreviewWidget) ZoomIn1(zoom float64) {
	QPrintPreviewWidget_ZoomIn1(this.h, (double)(zoom))
}

func (this *QPrintPreviewWidget) ZoomOut1(zoom float64) {
	QPrintPreviewWidget_ZoomOut1(this.h, (double)(zoom))
}

func (this *QPrintPreviewWidget) callVirtualBase_SetVisible(visible bool) {

	QPrintPreviewWidget_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QPrintPreviewWidget) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_SetVisible
func miqt_exec_callback_QPrintPreviewWidget_SetVisible(self QPrintPreviewWidget, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_DevType() int {

	return (int)(QPrintPreviewWidget_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QPrintPreviewWidget) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_DevType
func miqt_exec_callback_QPrintPreviewWidget_DevType(self QPrintPreviewWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QPrintPreviewWidget) callVirtualBase_SizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QPrintPreviewWidget_virtualbase_SizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPrintPreviewWidget) OnSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_SizeHint
func miqt_exec_callback_QPrintPreviewWidget_SizeHint(self QPrintPreviewWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_SizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewWidget) callVirtualBase_MinimumSizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QPrintPreviewWidget_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPrintPreviewWidget) OnMinimumSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_MinimumSizeHint
func miqt_exec_callback_QPrintPreviewWidget_MinimumSizeHint(self QPrintPreviewWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_MinimumSizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewWidget) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QPrintPreviewWidget_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QPrintPreviewWidget) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_HeightForWidth
func miqt_exec_callback_QPrintPreviewWidget_HeightForWidth(self QPrintPreviewWidget, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QPrintPreviewWidget) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QPrintPreviewWidget_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QPrintPreviewWidget) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_HasHeightForWidth
func miqt_exec_callback_QPrintPreviewWidget_HasHeightForWidth(self QPrintPreviewWidget, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QPrintPreviewWidget) callVirtualBase_PaintEngine() *qt.QPaintEngine {

	return qt.UnsafeNewQPaintEngine(unsafe.Pointer(QPrintPreviewWidget_virtualbase_PaintEngine(unsafe.Pointer(this.h))))

}
func (this *QPrintPreviewWidget) OnPaintEngine(slot func(super func() *qt.QPaintEngine) *qt.QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_PaintEngine
func miqt_exec_callback_QPrintPreviewWidget_PaintEngine(self QPrintPreviewWidget, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPaintEngine) *qt.QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_PaintEngine)

	return (*QPaintEngine)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewWidget) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QPrintPreviewWidget_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QPrintPreviewWidget) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_Event
func miqt_exec_callback_QPrintPreviewWidget_Event(self QPrintPreviewWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QPrintPreviewWidget) callVirtualBase_MousePressEvent(event *qt.QMouseEvent) {

	QPrintPreviewWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnMousePressEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_MousePressEvent
func miqt_exec_callback_QPrintPreviewWidget_MousePressEvent(self QPrintPreviewWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_MouseReleaseEvent(event *qt.QMouseEvent) {

	QPrintPreviewWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnMouseReleaseEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_MouseReleaseEvent
func miqt_exec_callback_QPrintPreviewWidget_MouseReleaseEvent(self QPrintPreviewWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_MouseDoubleClickEvent(event *qt.QMouseEvent) {

	QPrintPreviewWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnMouseDoubleClickEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_MouseDoubleClickEvent
func miqt_exec_callback_QPrintPreviewWidget_MouseDoubleClickEvent(self QPrintPreviewWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_MouseMoveEvent(event *qt.QMouseEvent) {

	QPrintPreviewWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnMouseMoveEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_MouseMoveEvent
func miqt_exec_callback_QPrintPreviewWidget_MouseMoveEvent(self QPrintPreviewWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_WheelEvent(event *qt.QWheelEvent) {

	QPrintPreviewWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), (*QWheelEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnWheelEvent(slot func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_WheelEvent
func miqt_exec_callback_QPrintPreviewWidget_WheelEvent(self QPrintPreviewWidget, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQWheelEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_KeyPressEvent(event *qt.QKeyEvent) {

	QPrintPreviewWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), (*QKeyEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnKeyPressEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_KeyPressEvent
func miqt_exec_callback_QPrintPreviewWidget_KeyPressEvent(self QPrintPreviewWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_KeyReleaseEvent(event *qt.QKeyEvent) {

	QPrintPreviewWidget_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), (*QKeyEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnKeyReleaseEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_KeyReleaseEvent
func miqt_exec_callback_QPrintPreviewWidget_KeyReleaseEvent(self QPrintPreviewWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_FocusInEvent(event *qt.QFocusEvent) {

	QPrintPreviewWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), (*QFocusEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnFocusInEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_FocusInEvent
func miqt_exec_callback_QPrintPreviewWidget_FocusInEvent(self QPrintPreviewWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_FocusOutEvent(event *qt.QFocusEvent) {

	QPrintPreviewWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), (*QFocusEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnFocusOutEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_FocusOutEvent
func miqt_exec_callback_QPrintPreviewWidget_FocusOutEvent(self QPrintPreviewWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_EnterEvent(event *qt.QEnterEvent) {

	QPrintPreviewWidget_virtualbase_EnterEvent(unsafe.Pointer(this.h), (*QEnterEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnEnterEvent(slot func(super func(event *qt.QEnterEvent), event *qt.QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_EnterEvent
func miqt_exec_callback_QPrintPreviewWidget_EnterEvent(self QPrintPreviewWidget, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEnterEvent), event *qt.QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEnterEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_LeaveEvent(event *qt.QEvent) {

	QPrintPreviewWidget_virtualbase_LeaveEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnLeaveEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_LeaveEvent
func miqt_exec_callback_QPrintPreviewWidget_LeaveEvent(self QPrintPreviewWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_PaintEvent(event *qt.QPaintEvent) {

	QPrintPreviewWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), (*QPaintEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnPaintEvent(slot func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_PaintEvent
func miqt_exec_callback_QPrintPreviewWidget_PaintEvent(self QPrintPreviewWidget, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPaintEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_MoveEvent(event *qt.QMoveEvent) {

	QPrintPreviewWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), (*QMoveEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnMoveEvent(slot func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_MoveEvent
func miqt_exec_callback_QPrintPreviewWidget_MoveEvent(self QPrintPreviewWidget, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMoveEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_ResizeEvent(event *qt.QResizeEvent) {

	QPrintPreviewWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), (*QResizeEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnResizeEvent(slot func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_ResizeEvent
func miqt_exec_callback_QPrintPreviewWidget_ResizeEvent(self QPrintPreviewWidget, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQResizeEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_CloseEvent(event *qt.QCloseEvent) {

	QPrintPreviewWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), (*QCloseEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnCloseEvent(slot func(super func(event *qt.QCloseEvent), event *qt.QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_CloseEvent
func miqt_exec_callback_QPrintPreviewWidget_CloseEvent(self QPrintPreviewWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QCloseEvent), event *qt.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQCloseEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_ContextMenuEvent(event *qt.QContextMenuEvent) {

	QPrintPreviewWidget_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), (*QContextMenuEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnContextMenuEvent(slot func(super func(event *qt.QContextMenuEvent), event *qt.QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_ContextMenuEvent
func miqt_exec_callback_QPrintPreviewWidget_ContextMenuEvent(self QPrintPreviewWidget, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QContextMenuEvent), event *qt.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQContextMenuEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_TabletEvent(event *qt.QTabletEvent) {

	QPrintPreviewWidget_virtualbase_TabletEvent(unsafe.Pointer(this.h), (*QTabletEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnTabletEvent(slot func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_TabletEvent
func miqt_exec_callback_QPrintPreviewWidget_TabletEvent(self QPrintPreviewWidget, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTabletEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_ActionEvent(event *qt.QActionEvent) {

	QPrintPreviewWidget_virtualbase_ActionEvent(unsafe.Pointer(this.h), (*QActionEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnActionEvent(slot func(super func(event *qt.QActionEvent), event *qt.QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_ActionEvent
func miqt_exec_callback_QPrintPreviewWidget_ActionEvent(self QPrintPreviewWidget, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QActionEvent), event *qt.QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQActionEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_DragEnterEvent(event *qt.QDragEnterEvent) {

	QPrintPreviewWidget_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), (*QDragEnterEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnDragEnterEvent(slot func(super func(event *qt.QDragEnterEvent), event *qt.QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_DragEnterEvent
func miqt_exec_callback_QPrintPreviewWidget_DragEnterEvent(self QPrintPreviewWidget, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragEnterEvent), event *qt.QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragEnterEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_DragMoveEvent(event *qt.QDragMoveEvent) {

	QPrintPreviewWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), (*QDragMoveEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnDragMoveEvent(slot func(super func(event *qt.QDragMoveEvent), event *qt.QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_DragMoveEvent
func miqt_exec_callback_QPrintPreviewWidget_DragMoveEvent(self QPrintPreviewWidget, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragMoveEvent), event *qt.QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragMoveEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_DragLeaveEvent(event *qt.QDragLeaveEvent) {

	QPrintPreviewWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), (*QDragLeaveEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnDragLeaveEvent(slot func(super func(event *qt.QDragLeaveEvent), event *qt.QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_DragLeaveEvent
func miqt_exec_callback_QPrintPreviewWidget_DragLeaveEvent(self QPrintPreviewWidget, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragLeaveEvent), event *qt.QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_DropEvent(event *qt.QDropEvent) {

	QPrintPreviewWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), (*QDropEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnDropEvent(slot func(super func(event *qt.QDropEvent), event *qt.QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_DropEvent
func miqt_exec_callback_QPrintPreviewWidget_DropEvent(self QPrintPreviewWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDropEvent), event *qt.QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDropEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_ShowEvent(event *qt.QShowEvent) {

	QPrintPreviewWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), (*QShowEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnShowEvent(slot func(super func(event *qt.QShowEvent), event *qt.QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_ShowEvent
func miqt_exec_callback_QPrintPreviewWidget_ShowEvent(self QPrintPreviewWidget, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QShowEvent), event *qt.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQShowEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_HideEvent(event *qt.QHideEvent) {

	QPrintPreviewWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), (*QHideEvent)(event.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnHideEvent(slot func(super func(event *qt.QHideEvent), event *qt.QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_HideEvent
func miqt_exec_callback_QPrintPreviewWidget_HideEvent(self QPrintPreviewWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QHideEvent), event *qt.QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQHideEvent(unsafe.Pointer(event))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QPrintPreviewWidget_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QPrintPreviewWidget) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_NativeEvent
func miqt_exec_callback_QPrintPreviewWidget_NativeEvent(self QPrintPreviewWidget, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QPrintPreviewWidget) callVirtualBase_ChangeEvent(param1 *qt.QEvent) {

	QPrintPreviewWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), (*QEvent)(param1.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnChangeEvent(slot func(super func(param1 *qt.QEvent), param1 *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_ChangeEvent
func miqt_exec_callback_QPrintPreviewWidget_ChangeEvent(self QPrintPreviewWidget, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QEvent), param1 *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(param1))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QPrintPreviewWidget_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QPrintPreviewWidget) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_Metric
func miqt_exec_callback_QPrintPreviewWidget_Metric(self QPrintPreviewWidget, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QPrintPreviewWidget) callVirtualBase_InitPainter(painter *qt.QPainter) {

	QPrintPreviewWidget_virtualbase_InitPainter(unsafe.Pointer(this.h), (*QPainter)(painter.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnInitPainter(slot func(super func(painter *qt.QPainter), painter *qt.QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_InitPainter
func miqt_exec_callback_QPrintPreviewWidget_InitPainter(self QPrintPreviewWidget, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter), painter *qt.QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_Redirected(offset *qt.QPoint) *qt.QPaintDevice {

	return qt.UnsafeNewQPaintDevice(unsafe.Pointer(QPrintPreviewWidget_virtualbase_Redirected(unsafe.Pointer(this.h), (*QPoint)(offset.UnsafePointer()))))

}
func (this *QPrintPreviewWidget) OnRedirected(slot func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_Redirected
func miqt_exec_callback_QPrintPreviewWidget_Redirected(self QPrintPreviewWidget, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPoint(unsafe.Pointer(offset))

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_Redirected, slotval1)

	return (*QPaintDevice)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewWidget) callVirtualBase_SharedPainter() *qt.QPainter {

	return qt.UnsafeNewQPainter(unsafe.Pointer(QPrintPreviewWidget_virtualbase_SharedPainter(unsafe.Pointer(this.h))))

}
func (this *QPrintPreviewWidget) OnSharedPainter(slot func(super func() *qt.QPainter) *qt.QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_SharedPainter
func miqt_exec_callback_QPrintPreviewWidget_SharedPainter(self QPrintPreviewWidget, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPainter) *qt.QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_SharedPainter)

	return (*QPainter)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewWidget) callVirtualBase_InputMethodEvent(param1 *qt.QInputMethodEvent) {

	QPrintPreviewWidget_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), (*QInputMethodEvent)(param1.UnsafePointer()))

}
func (this *QPrintPreviewWidget) OnInputMethodEvent(slot func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_InputMethodEvent
func miqt_exec_callback_QPrintPreviewWidget_InputMethodEvent(self QPrintPreviewWidget, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQInputMethodEvent(unsafe.Pointer(param1))

	gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QPrintPreviewWidget) callVirtualBase_InputMethodQuery(param1 qt.InputMethodQuery) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QPrintPreviewWidget_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPrintPreviewWidget) OnInputMethodQuery(slot func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_InputMethodQuery
func miqt_exec_callback_QPrintPreviewWidget_InputMethodQuery(self QPrintPreviewWidget, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.InputMethodQuery)(param1)

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return (*QVariant)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QPrintPreviewWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QPrintPreviewWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_FocusNextPrevChild
func miqt_exec_callback_QPrintPreviewWidget_FocusNextPrevChild(self QPrintPreviewWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
