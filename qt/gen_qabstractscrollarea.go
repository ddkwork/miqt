package qt

import (
	"unsafe"
)

type QAbstractScrollArea__SizeAdjustPolicy int

const (
	QAbstractScrollArea__AdjustIgnored               QAbstractScrollArea__SizeAdjustPolicy = 0
	QAbstractScrollArea__AdjustToContentsOnFirstShow QAbstractScrollArea__SizeAdjustPolicy = 1
	QAbstractScrollArea__AdjustToContents            QAbstractScrollArea__SizeAdjustPolicy = 2
)

type QAbstractScrollArea struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractScrollArea constructs a new QAbstractScrollArea object.
func NewQAbstractScrollArea(parent *QWidget) *QAbstractScrollArea {

	ret := newQAbstractScrollArea(QAbstractScrollArea_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAbstractScrollArea2 constructs a new QAbstractScrollArea object.
func NewQAbstractScrollArea2() *QAbstractScrollArea {

	ret := newQAbstractScrollArea(QAbstractScrollArea_new2())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractScrollArea) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractScrollArea_MetaObject(this.h))
}

func (this *QAbstractScrollArea) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractScrollArea_Metacast(this.h, param1_Cstring))
}

func QAbstractScrollArea_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractScrollArea_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractScrollArea) VerticalScrollBarPolicy() ScrollBarPolicy {
	return (ScrollBarPolicy)(QAbstractScrollArea_VerticalScrollBarPolicy(this.h))
}

func (this *QAbstractScrollArea) SetVerticalScrollBarPolicy(verticalScrollBarPolicy ScrollBarPolicy) {
	QAbstractScrollArea_SetVerticalScrollBarPolicy(this.h, (int)(verticalScrollBarPolicy))
}

func (this *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	return newQScrollBar(QAbstractScrollArea_VerticalScrollBar(this.h))
}

func (this *QAbstractScrollArea) SetVerticalScrollBar(scrollbar *QScrollBar) {
	QAbstractScrollArea_SetVerticalScrollBar(this.h, scrollbar.cPointer())
}

func (this *QAbstractScrollArea) HorizontalScrollBarPolicy() ScrollBarPolicy {
	return (ScrollBarPolicy)(QAbstractScrollArea_HorizontalScrollBarPolicy(this.h))
}

func (this *QAbstractScrollArea) SetHorizontalScrollBarPolicy(horizontalScrollBarPolicy ScrollBarPolicy) {
	QAbstractScrollArea_SetHorizontalScrollBarPolicy(this.h, (int)(horizontalScrollBarPolicy))
}

func (this *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	return newQScrollBar(QAbstractScrollArea_HorizontalScrollBar(this.h))
}

func (this *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar *QScrollBar) {
	QAbstractScrollArea_SetHorizontalScrollBar(this.h, scrollbar.cPointer())
}

func (this *QAbstractScrollArea) CornerWidget() *QWidget {
	return newQWidget(QAbstractScrollArea_CornerWidget(this.h))
}

func (this *QAbstractScrollArea) SetCornerWidget(widget *QWidget) {
	QAbstractScrollArea_SetCornerWidget(this.h, widget.cPointer())
}

func (this *QAbstractScrollArea) AddScrollBarWidget(widget *QWidget, alignment AlignmentFlag) {
	QAbstractScrollArea_AddScrollBarWidget(this.h, widget.cPointer(), (int)(alignment))
}

func (this *QAbstractScrollArea) ScrollBarWidgets(alignment AlignmentFlag) []*QWidget {
	var _ma struct_miqt_array = QAbstractScrollArea_ScrollBarWidgets(this.h, (int)(alignment))
	_ret := make([]*QWidget, int(_ma.len))
	_outCast := (*[0xffff]*QWidget)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQWidget(_outCast[i])
	}
	return _ret
}

func (this *QAbstractScrollArea) Viewport() *QWidget {
	return newQWidget(QAbstractScrollArea_Viewport(this.h))
}

func (this *QAbstractScrollArea) SetViewport(widget *QWidget) {
	QAbstractScrollArea_SetViewport(this.h, widget.cPointer())
}

func (this *QAbstractScrollArea) MaximumViewportSize() *QSize {
	_goptr := newQSize(QAbstractScrollArea_MaximumViewportSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractScrollArea) MinimumSizeHint() *QSize {
	_goptr := newQSize(QAbstractScrollArea_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractScrollArea) SizeHint() *QSize {
	_goptr := newQSize(QAbstractScrollArea_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractScrollArea) SetupViewport(viewport *QWidget) {
	QAbstractScrollArea_SetupViewport(this.h, viewport.cPointer())
}

func (this *QAbstractScrollArea) SizeAdjustPolicy() SizeAdjustPolicy {
	xxxxxxxxx
}

func (this *QAbstractScrollArea) SetSizeAdjustPolicy(policy SizeAdjustPolicy) {
	QAbstractScrollArea_SetSizeAdjustPolicy(this.h, policy)
}

func QAbstractScrollArea_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractScrollArea_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractScrollArea_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractScrollArea_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractScrollArea) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QAbstractScrollArea_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractScrollArea) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_MinimumSizeHint
func miqt_exec_callback_QAbstractScrollArea_MinimumSizeHint(self QAbstractScrollArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractScrollArea) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QAbstractScrollArea_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractScrollArea) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_SizeHint
func miqt_exec_callback_QAbstractScrollArea_SizeHint(self QAbstractScrollArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractScrollArea) callVirtualBase_SetupViewport(viewport *QWidget) {

	QAbstractScrollArea_virtualbase_SetupViewport(unsafe.Pointer(this.h), viewport.cPointer())

}
func (this *QAbstractScrollArea) OnSetupViewport(slot func(super func(viewport *QWidget), viewport *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_SetupViewport(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_SetupViewport
func miqt_exec_callback_QAbstractScrollArea_SetupViewport(self QAbstractScrollArea, cb intptr_t, viewport *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(viewport *QWidget), viewport *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(viewport)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_SetupViewport, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QAbstractScrollArea_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QAbstractScrollArea) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_EventFilter
func miqt_exec_callback_QAbstractScrollArea_EventFilter(self QAbstractScrollArea, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractScrollArea) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QAbstractScrollArea_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QAbstractScrollArea) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_Event
func miqt_exec_callback_QAbstractScrollArea_Event(self QAbstractScrollArea, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractScrollArea) callVirtualBase_ViewportEvent(param1 *QEvent) bool {

	return (bool)(QAbstractScrollArea_virtualbase_ViewportEvent(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QAbstractScrollArea) OnViewportEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_ViewportEvent
func miqt_exec_callback_QAbstractScrollArea_ViewportEvent(self QAbstractScrollArea, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractScrollArea) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QAbstractScrollArea_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_ResizeEvent
func miqt_exec_callback_QAbstractScrollArea_ResizeEvent(self QAbstractScrollArea, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QAbstractScrollArea_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_PaintEvent
func miqt_exec_callback_QAbstractScrollArea_PaintEvent(self QAbstractScrollArea, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QAbstractScrollArea_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_MousePressEvent
func miqt_exec_callback_QAbstractScrollArea_MousePressEvent(self QAbstractScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QAbstractScrollArea_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_MouseReleaseEvent
func miqt_exec_callback_QAbstractScrollArea_MouseReleaseEvent(self QAbstractScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_MouseDoubleClickEvent(param1 *QMouseEvent) {

	QAbstractScrollArea_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnMouseDoubleClickEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_MouseDoubleClickEvent
func miqt_exec_callback_QAbstractScrollArea_MouseDoubleClickEvent(self QAbstractScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QAbstractScrollArea_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_MouseMoveEvent
func miqt_exec_callback_QAbstractScrollArea_MouseMoveEvent(self QAbstractScrollArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_WheelEvent(param1 *QWheelEvent) {

	QAbstractScrollArea_virtualbase_WheelEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnWheelEvent(slot func(super func(param1 *QWheelEvent), param1 *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_WheelEvent
func miqt_exec_callback_QAbstractScrollArea_WheelEvent(self QAbstractScrollArea, cb intptr_t, param1 *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWheelEvent), param1 *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QAbstractScrollArea_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_ContextMenuEvent
func miqt_exec_callback_QAbstractScrollArea_ContextMenuEvent(self QAbstractScrollArea, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_DragEnterEvent(param1 *QDragEnterEvent) {

	QAbstractScrollArea_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnDragEnterEvent(slot func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_DragEnterEvent
func miqt_exec_callback_QAbstractScrollArea_DragEnterEvent(self QAbstractScrollArea, cb intptr_t, param1 *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_DragMoveEvent(param1 *QDragMoveEvent) {

	QAbstractScrollArea_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnDragMoveEvent(slot func(super func(param1 *QDragMoveEvent), param1 *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_DragMoveEvent
func miqt_exec_callback_QAbstractScrollArea_DragMoveEvent(self QAbstractScrollArea, cb intptr_t, param1 *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragMoveEvent), param1 *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_DragLeaveEvent(param1 *QDragLeaveEvent) {

	QAbstractScrollArea_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnDragLeaveEvent(slot func(super func(param1 *QDragLeaveEvent), param1 *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_DragLeaveEvent
func miqt_exec_callback_QAbstractScrollArea_DragLeaveEvent(self QAbstractScrollArea, cb intptr_t, param1 *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragLeaveEvent), param1 *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_DropEvent(param1 *QDropEvent) {

	QAbstractScrollArea_virtualbase_DropEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnDropEvent(slot func(super func(param1 *QDropEvent), param1 *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_DropEvent
func miqt_exec_callback_QAbstractScrollArea_DropEvent(self QAbstractScrollArea, cb intptr_t, param1 *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDropEvent), param1 *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QAbstractScrollArea_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_KeyPressEvent
func miqt_exec_callback_QAbstractScrollArea_KeyPressEvent(self QAbstractScrollArea, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QAbstractScrollArea_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QAbstractScrollArea) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_ScrollContentsBy
func miqt_exec_callback_QAbstractScrollArea_ScrollContentsBy(self QAbstractScrollArea, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QAbstractScrollArea) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QAbstractScrollArea_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractScrollArea) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_ViewportSizeHint
func miqt_exec_callback_QAbstractScrollArea_ViewportSizeHint(self QAbstractScrollArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractScrollArea) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QAbstractScrollArea_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractScrollArea) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_ChangeEvent
func miqt_exec_callback_QAbstractScrollArea_ChangeEvent(self QAbstractScrollArea, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QAbstractScrollArea) callVirtualBase_InitStyleOption(option *QStyleOptionFrame) {

	QAbstractScrollArea_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QAbstractScrollArea) OnInitStyleOption(slot func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_InitStyleOption
func miqt_exec_callback_QAbstractScrollArea_InitStyleOption(self QAbstractScrollArea, cb intptr_t, option *QStyleOptionFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionFrame(option)

	gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_InitStyleOption, slotval1)

}
