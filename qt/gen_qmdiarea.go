package qt

import (
	"unsafe"
)

type QMdiArea__AreaOption int

const (
	QMdiArea__DontMaximizeSubWindowOnActivation QMdiArea__AreaOption = 1
)

type QMdiArea__WindowOrder int

const (
	QMdiArea__CreationOrder          QMdiArea__WindowOrder = 0
	QMdiArea__StackingOrder          QMdiArea__WindowOrder = 1
	QMdiArea__ActivationHistoryOrder QMdiArea__WindowOrder = 2
)

type QMdiArea__ViewMode int

const (
	QMdiArea__SubWindowView QMdiArea__ViewMode = 0
	QMdiArea__TabbedView    QMdiArea__ViewMode = 1
)

type QMdiArea struct {
	h          uintptr
	isSubclass bool
}

// NewQMdiArea constructs a new QMdiArea object.
func NewQMdiArea(parent *QWidget) *QMdiArea {

	ret := newQMdiArea(QMdiArea_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMdiArea2 constructs a new QMdiArea object.
func NewQMdiArea2() *QMdiArea {

	ret := newQMdiArea(QMdiArea_new2())
	ret.isSubclass = true
	return ret
}

func (this *QMdiArea) MetaObject() *QMetaObject {
	return newQMetaObject(QMdiArea_MetaObject(this.h))
}

func (this *QMdiArea) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMdiArea_Metacast(this.h, param1_Cstring))
}

func QMdiArea_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMdiArea_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiArea) SizeHint() *QSize {
	_goptr := newQSize(QMdiArea_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiArea) MinimumSizeHint() *QSize {
	_goptr := newQSize(QMdiArea_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiArea) CurrentSubWindow() *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_CurrentSubWindow(this.h))
}

func (this *QMdiArea) ActiveSubWindow() *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_ActiveSubWindow(this.h))
}

func (this *QMdiArea) SubWindowList() []*QMdiSubWindow {
	var _ma struct_miqt_array = QMdiArea_SubWindowList(this.h)
	_ret := make([]*QMdiSubWindow, int(_ma.len))
	_outCast := (*[0xffff]*QMdiSubWindow)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQMdiSubWindow(_outCast[i])
	}
	return _ret
}

func (this *QMdiArea) AddSubWindow(widget *QWidget) *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_AddSubWindow(this.h, widget.cPointer()))
}

func (this *QMdiArea) RemoveSubWindow(widget *QWidget) {
	QMdiArea_RemoveSubWindow(this.h, widget.cPointer())
}

func (this *QMdiArea) Background() *QBrush {
	_goptr := newQBrush(QMdiArea_Background(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiArea) SetBackground(background *QBrush) {
	QMdiArea_SetBackground(this.h, background.cPointer())
}

func (this *QMdiArea) ActivationOrder() WindowOrder {
	xxxxxxxxx
}

func (this *QMdiArea) SetActivationOrder(order WindowOrder) {
	QMdiArea_SetActivationOrder(this.h, order)
}

func (this *QMdiArea) SetOption(option AreaOption) {
	QMdiArea_SetOption(this.h, option)
}

func (this *QMdiArea) TestOption(opton AreaOption) bool {
	return (bool)(QMdiArea_TestOption(this.h, opton))
}

func (this *QMdiArea) SetViewMode(mode ViewMode) {
	QMdiArea_SetViewMode(this.h, mode)
}

func (this *QMdiArea) ViewMode() ViewMode {
	xxxxxxxxx
}

func (this *QMdiArea) DocumentMode() bool {
	return (bool)(QMdiArea_DocumentMode(this.h))
}

func (this *QMdiArea) SetDocumentMode(enabled bool) {
	QMdiArea_SetDocumentMode(this.h, (bool)(enabled))
}

func (this *QMdiArea) SetTabsClosable(closable bool) {
	QMdiArea_SetTabsClosable(this.h, (bool)(closable))
}

func (this *QMdiArea) TabsClosable() bool {
	return (bool)(QMdiArea_TabsClosable(this.h))
}

func (this *QMdiArea) SetTabsMovable(movable bool) {
	QMdiArea_SetTabsMovable(this.h, (bool)(movable))
}

func (this *QMdiArea) TabsMovable() bool {
	return (bool)(QMdiArea_TabsMovable(this.h))
}

func (this *QMdiArea) SetTabShape(shape QTabWidget__TabShape) {
	QMdiArea_SetTabShape(this.h, (int)(shape))
}

func (this *QMdiArea) TabShape() QTabWidget__TabShape {
	return (QTabWidget__TabShape)(QMdiArea_TabShape(this.h))
}

func (this *QMdiArea) SetTabPosition(position QTabWidget__TabPosition) {
	QMdiArea_SetTabPosition(this.h, (int)(position))
}

func (this *QMdiArea) TabPosition() QTabWidget__TabPosition {
	return (QTabWidget__TabPosition)(QMdiArea_TabPosition(this.h))
}

func (this *QMdiArea) SubWindowActivated(param1 *QMdiSubWindow) {
	QMdiArea_SubWindowActivated(this.h, param1.cPointer())
}
func (this *QMdiArea) OnSubWindowActivated(slot func(param1 *QMdiSubWindow)) {
	QMdiArea_connect_SubWindowActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_SubWindowActivated
func miqt_exec_callback_QMdiArea_SubWindowActivated(cb intptr_t, param1 *QMdiSubWindow) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QMdiSubWindow))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMdiSubWindow(param1)

	gofunc(slotval1)
}

func (this *QMdiArea) SetActiveSubWindow(window *QMdiSubWindow) {
	QMdiArea_SetActiveSubWindow(this.h, window.cPointer())
}

func (this *QMdiArea) TileSubWindows() {
	QMdiArea_TileSubWindows(this.h)
}

func (this *QMdiArea) CascadeSubWindows() {
	QMdiArea_CascadeSubWindows(this.h)
}

func (this *QMdiArea) CloseActiveSubWindow() {
	QMdiArea_CloseActiveSubWindow(this.h)
}

func (this *QMdiArea) CloseAllSubWindows() {
	QMdiArea_CloseAllSubWindows(this.h)
}

func (this *QMdiArea) ActivateNextSubWindow() {
	QMdiArea_ActivateNextSubWindow(this.h)
}

func (this *QMdiArea) ActivatePreviousSubWindow() {
	QMdiArea_ActivatePreviousSubWindow(this.h)
}

func QMdiArea_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiArea_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMdiArea_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiArea_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiArea) SubWindowList1(order WindowOrder) []*QMdiSubWindow {
	var _ma struct_miqt_array = QMdiArea_SubWindowList1(this.h, order)
	_ret := make([]*QMdiSubWindow, int(_ma.len))
	_outCast := (*[0xffff]*QMdiSubWindow)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQMdiSubWindow(_outCast[i])
	}
	return _ret
}

func (this *QMdiArea) AddSubWindow2(widget *QWidget, flags WindowType) *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_AddSubWindow2(this.h, widget.cPointer(), (int)(flags)))
}

func (this *QMdiArea) SetOption2(option AreaOption, on bool) {
	QMdiArea_SetOption2(this.h, option, (bool)(on))
}

func (this *QMdiArea) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QMdiArea_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMdiArea) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_SizeHint
func miqt_exec_callback_QMdiArea_SizeHint(self QMdiArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QMdiArea) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QMdiArea_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMdiArea) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_MinimumSizeHint
func miqt_exec_callback_QMdiArea_MinimumSizeHint(self QMdiArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QMdiArea) callVirtualBase_SetupViewport(viewport *QWidget) {

	QMdiArea_virtualbase_SetupViewport(unsafe.Pointer(this.h), viewport.cPointer())

}
func (this *QMdiArea) OnSetupViewport(slot func(super func(viewport *QWidget), viewport *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_SetupViewport(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_SetupViewport
func miqt_exec_callback_QMdiArea_SetupViewport(self QMdiArea, cb intptr_t, viewport *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(viewport *QWidget), viewport *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(viewport)

	gofunc((&QMdiArea{h: self}).callVirtualBase_SetupViewport, slotval1)

}

func (this *QMdiArea) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QMdiArea_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QMdiArea) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_Event
func miqt_exec_callback_QMdiArea_Event(self QMdiArea, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QMdiArea) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QMdiArea_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QMdiArea) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_EventFilter
func miqt_exec_callback_QMdiArea_EventFilter(self QMdiArea, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QMdiArea) callVirtualBase_PaintEvent(paintEvent *QPaintEvent) {

	QMdiArea_virtualbase_PaintEvent(unsafe.Pointer(this.h), paintEvent.cPointer())

}
func (this *QMdiArea) OnPaintEvent(slot func(super func(paintEvent *QPaintEvent), paintEvent *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_PaintEvent
func miqt_exec_callback_QMdiArea_PaintEvent(self QMdiArea, cb intptr_t, paintEvent *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(paintEvent *QPaintEvent), paintEvent *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(paintEvent)

	gofunc((&QMdiArea{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_ChildEvent(childEvent *QChildEvent) {

	QMdiArea_virtualbase_ChildEvent(unsafe.Pointer(this.h), childEvent.cPointer())

}
func (this *QMdiArea) OnChildEvent(slot func(super func(childEvent *QChildEvent), childEvent *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_ChildEvent
func miqt_exec_callback_QMdiArea_ChildEvent(self QMdiArea, cb intptr_t, childEvent *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(childEvent *QChildEvent), childEvent *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(childEvent)

	gofunc((&QMdiArea{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_ResizeEvent(resizeEvent *QResizeEvent) {

	QMdiArea_virtualbase_ResizeEvent(unsafe.Pointer(this.h), resizeEvent.cPointer())

}
func (this *QMdiArea) OnResizeEvent(slot func(super func(resizeEvent *QResizeEvent), resizeEvent *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_ResizeEvent
func miqt_exec_callback_QMdiArea_ResizeEvent(self QMdiArea, cb intptr_t, resizeEvent *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(resizeEvent *QResizeEvent), resizeEvent *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(resizeEvent)

	gofunc((&QMdiArea{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_TimerEvent(timerEvent *QTimerEvent) {

	QMdiArea_virtualbase_TimerEvent(unsafe.Pointer(this.h), timerEvent.cPointer())

}
func (this *QMdiArea) OnTimerEvent(slot func(super func(timerEvent *QTimerEvent), timerEvent *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_TimerEvent
func miqt_exec_callback_QMdiArea_TimerEvent(self QMdiArea, cb intptr_t, timerEvent *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timerEvent *QTimerEvent), timerEvent *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(timerEvent)

	gofunc((&QMdiArea{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_ShowEvent(showEvent *QShowEvent) {

	QMdiArea_virtualbase_ShowEvent(unsafe.Pointer(this.h), showEvent.cPointer())

}
func (this *QMdiArea) OnShowEvent(slot func(super func(showEvent *QShowEvent), showEvent *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_ShowEvent
func miqt_exec_callback_QMdiArea_ShowEvent(self QMdiArea, cb intptr_t, showEvent *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(showEvent *QShowEvent), showEvent *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(showEvent)

	gofunc((&QMdiArea{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_ViewportEvent(event *QEvent) bool {

	return (bool)(QMdiArea_virtualbase_ViewportEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QMdiArea) OnViewportEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_ViewportEvent
func miqt_exec_callback_QMdiArea_ViewportEvent(self QMdiArea, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QMdiArea) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QMdiArea_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QMdiArea) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_ScrollContentsBy
func miqt_exec_callback_QMdiArea_ScrollContentsBy(self QMdiArea, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QMdiArea{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QMdiArea) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QMdiArea_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_MousePressEvent
func miqt_exec_callback_QMdiArea_MousePressEvent(self QMdiArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QMdiArea_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_MouseReleaseEvent
func miqt_exec_callback_QMdiArea_MouseReleaseEvent(self QMdiArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_MouseDoubleClickEvent(param1 *QMouseEvent) {

	QMdiArea_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnMouseDoubleClickEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_MouseDoubleClickEvent
func miqt_exec_callback_QMdiArea_MouseDoubleClickEvent(self QMdiArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QMdiArea_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_MouseMoveEvent
func miqt_exec_callback_QMdiArea_MouseMoveEvent(self QMdiArea, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_WheelEvent(param1 *QWheelEvent) {

	QMdiArea_virtualbase_WheelEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnWheelEvent(slot func(super func(param1 *QWheelEvent), param1 *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_WheelEvent
func miqt_exec_callback_QMdiArea_WheelEvent(self QMdiArea, cb intptr_t, param1 *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWheelEvent), param1 *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QMdiArea_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_ContextMenuEvent
func miqt_exec_callback_QMdiArea_ContextMenuEvent(self QMdiArea, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_DragEnterEvent(param1 *QDragEnterEvent) {

	QMdiArea_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnDragEnterEvent(slot func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_DragEnterEvent
func miqt_exec_callback_QMdiArea_DragEnterEvent(self QMdiArea, cb intptr_t, param1 *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_DragMoveEvent(param1 *QDragMoveEvent) {

	QMdiArea_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnDragMoveEvent(slot func(super func(param1 *QDragMoveEvent), param1 *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_DragMoveEvent
func miqt_exec_callback_QMdiArea_DragMoveEvent(self QMdiArea, cb intptr_t, param1 *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragMoveEvent), param1 *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_DragLeaveEvent(param1 *QDragLeaveEvent) {

	QMdiArea_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnDragLeaveEvent(slot func(super func(param1 *QDragLeaveEvent), param1 *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_DragLeaveEvent
func miqt_exec_callback_QMdiArea_DragLeaveEvent(self QMdiArea, cb intptr_t, param1 *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragLeaveEvent), param1 *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_DropEvent(param1 *QDropEvent) {

	QMdiArea_virtualbase_DropEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnDropEvent(slot func(super func(param1 *QDropEvent), param1 *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_DropEvent
func miqt_exec_callback_QMdiArea_DropEvent(self QMdiArea, cb intptr_t, param1 *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDropEvent), param1 *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QMdiArea_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiArea) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_KeyPressEvent
func miqt_exec_callback_QMdiArea_KeyPressEvent(self QMdiArea, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QMdiArea{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QMdiArea) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QMdiArea_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMdiArea) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_ViewportSizeHint
func miqt_exec_callback_QMdiArea_ViewportSizeHint(self QMdiArea, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}
