package qt

import (
	"unsafe"
)

type QScrollArea struct {
	h          uintptr
	isSubclass bool
}

// NewQScrollArea constructs a new QScrollArea object.
func NewQScrollArea(parent *QWidget) *QScrollArea {

	ret := newQScrollArea(QScrollArea_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQScrollArea2 constructs a new QScrollArea object.
func NewQScrollArea2() *QScrollArea {

	ret := newQScrollArea(QScrollArea_new2())
	ret.isSubclass = true
	return ret
}

func (this *QScrollArea) MetaObject() *QMetaObject {
	return newQMetaObject(QScrollArea_MetaObject(this.h))
}

func (this *QScrollArea) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QScrollArea_Metacast(this.h, param1_Cstring))
}

func QScrollArea_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QScrollArea_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScrollArea) Widget() *QWidget {
	return newQWidget(QScrollArea_Widget(this.h))
}

func (this *QScrollArea) SetWidget(widget *QWidget) {
	QScrollArea_SetWidget(this.h, widget.cPointer())
}

func (this *QScrollArea) TakeWidget() *QWidget {
	return newQWidget(QScrollArea_TakeWidget(this.h))
}

func (this *QScrollArea) WidgetResizable() bool {
	return (bool)(QScrollArea_WidgetResizable(this.h))
}

func (this *QScrollArea) SetWidgetResizable(resizable bool) {
	QScrollArea_SetWidgetResizable(this.h, (bool)(resizable))
}

func (this *QScrollArea) SizeHint() *QSize {
	_goptr := newQSize(QScrollArea_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollArea) FocusNextPrevChild(next bool) bool {
	return (bool)(QScrollArea_FocusNextPrevChild(this.h, (bool)(next)))
}

func (this *QScrollArea) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QScrollArea_Alignment(this.h))
}

func (this *QScrollArea) SetAlignment(alignment AlignmentFlag) {
	QScrollArea_SetAlignment(this.h, (int)(alignment))
}

func (this *QScrollArea) EnsureVisible(x int, y int) {
	QScrollArea_EnsureVisible(this.h, (int)(x), (int)(y))
}

func (this *QScrollArea) EnsureWidgetVisible(childWidget *QWidget) {
	QScrollArea_EnsureWidgetVisible(this.h, childWidget.cPointer())
}

func QScrollArea_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScrollArea_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScrollArea_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScrollArea_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScrollArea) EnsureVisible3(x int, y int, xmargin int) {
	QScrollArea_EnsureVisible3(this.h, (int)(x), (int)(y), (int)(xmargin))
}

func (this *QScrollArea) EnsureVisible4(x int, y int, xmargin int, ymargin int) {
	QScrollArea_EnsureVisible4(this.h, (int)(x), (int)(y), (int)(xmargin), (int)(ymargin))
}

func (this *QScrollArea) EnsureWidgetVisible2(childWidget *QWidget, xmargin int) {
	QScrollArea_EnsureWidgetVisible2(this.h, childWidget.cPointer(), (int)(xmargin))
}

func (this *QScrollArea) EnsureWidgetVisible3(childWidget *QWidget, xmargin int, ymargin int) {
	QScrollArea_EnsureWidgetVisible3(this.h, childWidget.cPointer(), (int)(xmargin), (int)(ymargin))
}

func (this *QScrollArea) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QScrollArea_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QScrollArea) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_SizeHint
func miqt_exec_callback_QScrollArea_SizeHint(self QScrollArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QScrollArea) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QScrollArea_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QScrollArea) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_FocusNextPrevChild
func miqt_exec_callback_QScrollArea_FocusNextPrevChild(self QScrollArea, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QScrollArea) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QScrollArea_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QScrollArea) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_Event
func miqt_exec_callback_QScrollArea_Event(self QScrollArea, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QScrollArea) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QScrollArea_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QScrollArea) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_EventFilter
func miqt_exec_callback_QScrollArea_EventFilter(self QScrollArea, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QScrollArea) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QScrollArea_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_ResizeEvent
func miqt_exec_callback_QScrollArea_ResizeEvent(self QScrollArea, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QScrollArea_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QScrollArea) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_ScrollContentsBy
func miqt_exec_callback_QScrollArea_ScrollContentsBy(self QScrollArea, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QScrollArea{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QScrollArea) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QScrollArea_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QScrollArea) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_ViewportSizeHint
func miqt_exec_callback_QScrollArea_ViewportSizeHint(self QScrollArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}

func (this *QScrollArea) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QScrollArea_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QScrollArea) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_MinimumSizeHint
func miqt_exec_callback_QScrollArea_MinimumSizeHint(self QScrollArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QScrollArea) callVirtualBase_SetupViewport(viewport *QWidget) {

	QScrollArea_virtualbase_SetupViewport(unsafe.Pointer(this.h), viewport.cPointer())

}
func (this *QScrollArea) OnSetupViewport(slot func(super func(viewport *QWidget), viewport *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_SetupViewport(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_SetupViewport
func miqt_exec_callback_QScrollArea_SetupViewport(self QScrollArea, cb intptr_t, viewport *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(viewport *QWidget), viewport *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(viewport)

	gofunc((&QScrollArea{h: self}).callVirtualBase_SetupViewport, slotval1)

}

func (this *QScrollArea) callVirtualBase_ViewportEvent(param1 *QEvent) bool {

	return (bool)(QScrollArea_virtualbase_ViewportEvent(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QScrollArea) OnViewportEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_ViewportEvent
func miqt_exec_callback_QScrollArea_ViewportEvent(self QScrollArea, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QScrollArea) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QScrollArea_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_PaintEvent
func miqt_exec_callback_QScrollArea_PaintEvent(self QScrollArea, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QScrollArea_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_MousePressEvent
func miqt_exec_callback_QScrollArea_MousePressEvent(self QScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QScrollArea_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_MouseReleaseEvent
func miqt_exec_callback_QScrollArea_MouseReleaseEvent(self QScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_MouseDoubleClickEvent(param1 *QMouseEvent) {

	QScrollArea_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnMouseDoubleClickEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_MouseDoubleClickEvent
func miqt_exec_callback_QScrollArea_MouseDoubleClickEvent(self QScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QScrollArea_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_MouseMoveEvent
func miqt_exec_callback_QScrollArea_MouseMoveEvent(self QScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_WheelEvent(param1 *QWheelEvent) {

	QScrollArea_virtualbase_WheelEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnWheelEvent(slot func(super func(param1 *QWheelEvent), param1 *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_WheelEvent
func miqt_exec_callback_QScrollArea_WheelEvent(self QScrollArea, cb intptr_t, param1 *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWheelEvent), param1 *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QScrollArea_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_ContextMenuEvent
func miqt_exec_callback_QScrollArea_ContextMenuEvent(self QScrollArea, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_DragEnterEvent(param1 *QDragEnterEvent) {

	QScrollArea_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnDragEnterEvent(slot func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_DragEnterEvent
func miqt_exec_callback_QScrollArea_DragEnterEvent(self QScrollArea, cb intptr_t, param1 *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_DragMoveEvent(param1 *QDragMoveEvent) {

	QScrollArea_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnDragMoveEvent(slot func(super func(param1 *QDragMoveEvent), param1 *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_DragMoveEvent
func miqt_exec_callback_QScrollArea_DragMoveEvent(self QScrollArea, cb intptr_t, param1 *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragMoveEvent), param1 *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_DragLeaveEvent(param1 *QDragLeaveEvent) {

	QScrollArea_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnDragLeaveEvent(slot func(super func(param1 *QDragLeaveEvent), param1 *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_DragLeaveEvent
func miqt_exec_callback_QScrollArea_DragLeaveEvent(self QScrollArea, cb intptr_t, param1 *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragLeaveEvent), param1 *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_DropEvent(param1 *QDropEvent) {

	QScrollArea_virtualbase_DropEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnDropEvent(slot func(super func(param1 *QDropEvent), param1 *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_DropEvent
func miqt_exec_callback_QScrollArea_DropEvent(self QScrollArea, cb intptr_t, param1 *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDropEvent), param1 *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QScrollArea) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QScrollArea_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QScrollArea) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_KeyPressEvent
func miqt_exec_callback_QScrollArea_KeyPressEvent(self QScrollArea, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QScrollArea{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}
