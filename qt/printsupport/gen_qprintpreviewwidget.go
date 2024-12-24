package printsupport

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
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
	g := newQPrintPreviewWidget(QPrintPreviewWidget_new((*QWidget)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQPrintPreviewWidget2 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget2(printer *QPrinter) *QPrintPreviewWidget {
	g := newQPrintPreviewWidget(QPrintPreviewWidget_new2(printer.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPrintPreviewWidget3 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget3() *QPrintPreviewWidget {
	g := newQPrintPreviewWidget(QPrintPreviewWidget_new3())
	g.isSubclass = true
	return g
}

// NewQPrintPreviewWidget4 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget4(printer *QPrinter, parent *qt.QWidget) *QPrintPreviewWidget {
	g := newQPrintPreviewWidget(QPrintPreviewWidget_new4(printer.cPointer(), (*QWidget)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQPrintPreviewWidget5 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget5(printer *QPrinter, parent *qt.QWidget, flags WindowType) *QPrintPreviewWidget {
	g := newQPrintPreviewWidget(QPrintPreviewWidget_new5(printer.cPointer(), (*QWidget)(parent.UnsafePointer()), (int)(flags)))
	g.isSubclass = true
	return g
}

// NewQPrintPreviewWidget6 constructs a new QPrintPreviewWidget object.
func NewQPrintPreviewWidget6(parent *qt.QWidget, flags WindowType) *QPrintPreviewWidget {
	g := newQPrintPreviewWidget(QPrintPreviewWidget_new6((*QWidget)(parent.UnsafePointer()), (int)(flags)))
	g.isSubclass = true
	return g
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

func (this *QPrintPreviewWidget) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QPrintPreviewWidget_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QPrintPreviewWidget) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_MetaObject
func miqt_exec_callback_QPrintPreviewWidget_MetaObject(self QPrintPreviewWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QPrintPreviewWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPrintPreviewWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPrintPreviewWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewWidget_Metacast
func miqt_exec_callback_QPrintPreviewWidget_Metacast(self QPrintPreviewWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QPrintPreviewWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
