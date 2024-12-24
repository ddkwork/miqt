package qt

import (
	"unsafe"
)

type QMdiSubWindow__SubWindowOption int

const (
	QMdiSubWindow__AllowOutsideAreaHorizontally QMdiSubWindow__SubWindowOption = 1
	QMdiSubWindow__AllowOutsideAreaVertically   QMdiSubWindow__SubWindowOption = 2
	QMdiSubWindow__RubberBandResize             QMdiSubWindow__SubWindowOption = 4
	QMdiSubWindow__RubberBandMove               QMdiSubWindow__SubWindowOption = 8
)

type QMdiSubWindow struct {
	h          uintptr
	isSubclass bool
}

// NewQMdiSubWindow constructs a new QMdiSubWindow object.
func NewQMdiSubWindow(parent *QWidget) *QMdiSubWindow {

	ret := newQMdiSubWindow(QMdiSubWindow_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMdiSubWindow2 constructs a new QMdiSubWindow object.
func NewQMdiSubWindow2() *QMdiSubWindow {

	ret := newQMdiSubWindow(QMdiSubWindow_new2())
	ret.isSubclass = true
	return ret
}

// NewQMdiSubWindow3 constructs a new QMdiSubWindow object.
func NewQMdiSubWindow3(parent *QWidget, flags WindowType) *QMdiSubWindow {

	ret := newQMdiSubWindow(QMdiSubWindow_new3(parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QMdiSubWindow) MetaObject() *QMetaObject {
	return newQMetaObject(QMdiSubWindow_MetaObject(this.h))
}

func (this *QMdiSubWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMdiSubWindow_Metacast(this.h, param1_Cstring))
}

func QMdiSubWindow_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMdiSubWindow_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiSubWindow) SizeHint() *QSize {
	_goptr := newQSize(QMdiSubWindow_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiSubWindow) MinimumSizeHint() *QSize {
	_goptr := newQSize(QMdiSubWindow_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiSubWindow) SetWidget(widget *QWidget) {
	QMdiSubWindow_SetWidget(this.h, widget.cPointer())
}

func (this *QMdiSubWindow) Widget() *QWidget {
	return newQWidget(QMdiSubWindow_Widget(this.h))
}

func (this *QMdiSubWindow) MaximizedButtonsWidget() *QWidget {
	return newQWidget(QMdiSubWindow_MaximizedButtonsWidget(this.h))
}

func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() *QWidget {
	return newQWidget(QMdiSubWindow_MaximizedSystemMenuIconWidget(this.h))
}

func (this *QMdiSubWindow) IsShaded() bool {
	return (bool)(QMdiSubWindow_IsShaded(this.h))
}

func (this *QMdiSubWindow) SetOption(option SubWindowOption) {
	QMdiSubWindow_SetOption(this.h, option)
}

func (this *QMdiSubWindow) TestOption(param1 SubWindowOption) bool {
	return (bool)(QMdiSubWindow_TestOption(this.h, param1))
}

func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	QMdiSubWindow_SetKeyboardSingleStep(this.h, (int)(step))
}

func (this *QMdiSubWindow) KeyboardSingleStep() int {
	return (int)(QMdiSubWindow_KeyboardSingleStep(this.h))
}

func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	QMdiSubWindow_SetKeyboardPageStep(this.h, (int)(step))
}

func (this *QMdiSubWindow) KeyboardPageStep() int {
	return (int)(QMdiSubWindow_KeyboardPageStep(this.h))
}

func (this *QMdiSubWindow) SetSystemMenu(systemMenu *QMenu) {
	QMdiSubWindow_SetSystemMenu(this.h, systemMenu.cPointer())
}

func (this *QMdiSubWindow) SystemMenu() *QMenu {
	return newQMenu(QMdiSubWindow_SystemMenu(this.h))
}

func (this *QMdiSubWindow) MdiArea() *QMdiArea {
	return newQMdiArea(QMdiSubWindow_MdiArea(this.h))
}

func (this *QMdiSubWindow) WindowStateChanged(oldState WindowState, newState WindowState) {
	QMdiSubWindow_WindowStateChanged(this.h, (int)(oldState), (int)(newState))
}
func (this *QMdiSubWindow) OnWindowStateChanged(slot func(oldState WindowState, newState WindowState)) {
	QMdiSubWindow_connect_WindowStateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_WindowStateChanged
func miqt_exec_callback_QMdiSubWindow_WindowStateChanged(cb intptr_t, oldState int, newState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(oldState WindowState, newState WindowState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (WindowState)(oldState)

	slotval2 := (WindowState)(newState)

	gofunc(slotval1, slotval2)
}

func (this *QMdiSubWindow) AboutToActivate() {
	QMdiSubWindow_AboutToActivate(this.h)
}
func (this *QMdiSubWindow) OnAboutToActivate(slot func()) {
	QMdiSubWindow_connect_AboutToActivate(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_AboutToActivate
func miqt_exec_callback_QMdiSubWindow_AboutToActivate(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMdiSubWindow) ShowSystemMenu() {
	QMdiSubWindow_ShowSystemMenu(this.h)
}

func (this *QMdiSubWindow) ShowShaded() {
	QMdiSubWindow_ShowShaded(this.h)
}

func QMdiSubWindow_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiSubWindow_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMdiSubWindow_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiSubWindow_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiSubWindow) SetOption2(option SubWindowOption, on bool) {
	QMdiSubWindow_SetOption2(this.h, option, (bool)(on))
}

func (this *QMdiSubWindow) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QMdiSubWindow_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMdiSubWindow) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_SizeHint
func miqt_exec_callback_QMdiSubWindow_SizeHint(self QMdiSubWindow, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QMdiSubWindow) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QMdiSubWindow_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMdiSubWindow) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_MinimumSizeHint
func miqt_exec_callback_QMdiSubWindow_MinimumSizeHint(self QMdiSubWindow, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QMdiSubWindow) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QMdiSubWindow_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QMdiSubWindow) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_EventFilter
func miqt_exec_callback_QMdiSubWindow_EventFilter(self QMdiSubWindow, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QMdiSubWindow) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QMdiSubWindow_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QMdiSubWindow) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_Event
func miqt_exec_callback_QMdiSubWindow_Event(self QMdiSubWindow, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QMdiSubWindow) callVirtualBase_ShowEvent(showEvent *QShowEvent) {

	QMdiSubWindow_virtualbase_ShowEvent(unsafe.Pointer(this.h), showEvent.cPointer())

}
func (this *QMdiSubWindow) OnShowEvent(slot func(super func(showEvent *QShowEvent), showEvent *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_ShowEvent
func miqt_exec_callback_QMdiSubWindow_ShowEvent(self QMdiSubWindow, cb intptr_t, showEvent *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(showEvent *QShowEvent), showEvent *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(showEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_HideEvent(hideEvent *QHideEvent) {

	QMdiSubWindow_virtualbase_HideEvent(unsafe.Pointer(this.h), hideEvent.cPointer())

}
func (this *QMdiSubWindow) OnHideEvent(slot func(super func(hideEvent *QHideEvent), hideEvent *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_HideEvent
func miqt_exec_callback_QMdiSubWindow_HideEvent(self QMdiSubWindow, cb intptr_t, hideEvent *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(hideEvent *QHideEvent), hideEvent *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(hideEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_ChangeEvent(changeEvent *QEvent) {

	QMdiSubWindow_virtualbase_ChangeEvent(unsafe.Pointer(this.h), changeEvent.cPointer())

}
func (this *QMdiSubWindow) OnChangeEvent(slot func(super func(changeEvent *QEvent), changeEvent *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_ChangeEvent
func miqt_exec_callback_QMdiSubWindow_ChangeEvent(self QMdiSubWindow, cb intptr_t, changeEvent *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(changeEvent *QEvent), changeEvent *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(changeEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_CloseEvent(closeEvent *QCloseEvent) {

	QMdiSubWindow_virtualbase_CloseEvent(unsafe.Pointer(this.h), closeEvent.cPointer())

}
func (this *QMdiSubWindow) OnCloseEvent(slot func(super func(closeEvent *QCloseEvent), closeEvent *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_CloseEvent
func miqt_exec_callback_QMdiSubWindow_CloseEvent(self QMdiSubWindow, cb intptr_t, closeEvent *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(closeEvent *QCloseEvent), closeEvent *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(closeEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_LeaveEvent(leaveEvent *QEvent) {

	QMdiSubWindow_virtualbase_LeaveEvent(unsafe.Pointer(this.h), leaveEvent.cPointer())

}
func (this *QMdiSubWindow) OnLeaveEvent(slot func(super func(leaveEvent *QEvent), leaveEvent *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_LeaveEvent
func miqt_exec_callback_QMdiSubWindow_LeaveEvent(self QMdiSubWindow, cb intptr_t, leaveEvent *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(leaveEvent *QEvent), leaveEvent *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(leaveEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_ResizeEvent(resizeEvent *QResizeEvent) {

	QMdiSubWindow_virtualbase_ResizeEvent(unsafe.Pointer(this.h), resizeEvent.cPointer())

}
func (this *QMdiSubWindow) OnResizeEvent(slot func(super func(resizeEvent *QResizeEvent), resizeEvent *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_ResizeEvent
func miqt_exec_callback_QMdiSubWindow_ResizeEvent(self QMdiSubWindow, cb intptr_t, resizeEvent *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(resizeEvent *QResizeEvent), resizeEvent *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(resizeEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_TimerEvent(timerEvent *QTimerEvent) {

	QMdiSubWindow_virtualbase_TimerEvent(unsafe.Pointer(this.h), timerEvent.cPointer())

}
func (this *QMdiSubWindow) OnTimerEvent(slot func(super func(timerEvent *QTimerEvent), timerEvent *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_TimerEvent
func miqt_exec_callback_QMdiSubWindow_TimerEvent(self QMdiSubWindow, cb intptr_t, timerEvent *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(timerEvent *QTimerEvent), timerEvent *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(timerEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_MoveEvent(moveEvent *QMoveEvent) {

	QMdiSubWindow_virtualbase_MoveEvent(unsafe.Pointer(this.h), moveEvent.cPointer())

}
func (this *QMdiSubWindow) OnMoveEvent(slot func(super func(moveEvent *QMoveEvent), moveEvent *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_MoveEvent
func miqt_exec_callback_QMdiSubWindow_MoveEvent(self QMdiSubWindow, cb intptr_t, moveEvent *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(moveEvent *QMoveEvent), moveEvent *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(moveEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_PaintEvent(paintEvent *QPaintEvent) {

	QMdiSubWindow_virtualbase_PaintEvent(unsafe.Pointer(this.h), paintEvent.cPointer())

}
func (this *QMdiSubWindow) OnPaintEvent(slot func(super func(paintEvent *QPaintEvent), paintEvent *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_PaintEvent
func miqt_exec_callback_QMdiSubWindow_PaintEvent(self QMdiSubWindow, cb intptr_t, paintEvent *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(paintEvent *QPaintEvent), paintEvent *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(paintEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_MousePressEvent(mouseEvent *QMouseEvent) {

	QMdiSubWindow_virtualbase_MousePressEvent(unsafe.Pointer(this.h), mouseEvent.cPointer())

}
func (this *QMdiSubWindow) OnMousePressEvent(slot func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_MousePressEvent
func miqt_exec_callback_QMdiSubWindow_MousePressEvent(self QMdiSubWindow, cb intptr_t, mouseEvent *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(mouseEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_MouseDoubleClickEvent(mouseEvent *QMouseEvent) {

	QMdiSubWindow_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), mouseEvent.cPointer())

}
func (this *QMdiSubWindow) OnMouseDoubleClickEvent(slot func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_MouseDoubleClickEvent
func miqt_exec_callback_QMdiSubWindow_MouseDoubleClickEvent(self QMdiSubWindow, cb intptr_t, mouseEvent *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(mouseEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_MouseReleaseEvent(mouseEvent *QMouseEvent) {

	QMdiSubWindow_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), mouseEvent.cPointer())

}
func (this *QMdiSubWindow) OnMouseReleaseEvent(slot func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_MouseReleaseEvent
func miqt_exec_callback_QMdiSubWindow_MouseReleaseEvent(self QMdiSubWindow, cb intptr_t, mouseEvent *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(mouseEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_MouseMoveEvent(mouseEvent *QMouseEvent) {

	QMdiSubWindow_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), mouseEvent.cPointer())

}
func (this *QMdiSubWindow) OnMouseMoveEvent(slot func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_MouseMoveEvent
func miqt_exec_callback_QMdiSubWindow_MouseMoveEvent(self QMdiSubWindow, cb intptr_t, mouseEvent *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(mouseEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_KeyPressEvent(keyEvent *QKeyEvent) {

	QMdiSubWindow_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), keyEvent.cPointer())

}
func (this *QMdiSubWindow) OnKeyPressEvent(slot func(super func(keyEvent *QKeyEvent), keyEvent *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_KeyPressEvent
func miqt_exec_callback_QMdiSubWindow_KeyPressEvent(self QMdiSubWindow, cb intptr_t, keyEvent *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(keyEvent *QKeyEvent), keyEvent *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(keyEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_ContextMenuEvent(contextMenuEvent *QContextMenuEvent) {

	QMdiSubWindow_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), contextMenuEvent.cPointer())

}
func (this *QMdiSubWindow) OnContextMenuEvent(slot func(super func(contextMenuEvent *QContextMenuEvent), contextMenuEvent *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_ContextMenuEvent
func miqt_exec_callback_QMdiSubWindow_ContextMenuEvent(self QMdiSubWindow, cb intptr_t, contextMenuEvent *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(contextMenuEvent *QContextMenuEvent), contextMenuEvent *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(contextMenuEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_FocusInEvent(focusInEvent *QFocusEvent) {

	QMdiSubWindow_virtualbase_FocusInEvent(unsafe.Pointer(this.h), focusInEvent.cPointer())

}
func (this *QMdiSubWindow) OnFocusInEvent(slot func(super func(focusInEvent *QFocusEvent), focusInEvent *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_FocusInEvent
func miqt_exec_callback_QMdiSubWindow_FocusInEvent(self QMdiSubWindow, cb intptr_t, focusInEvent *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(focusInEvent *QFocusEvent), focusInEvent *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(focusInEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_FocusOutEvent(focusOutEvent *QFocusEvent) {

	QMdiSubWindow_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), focusOutEvent.cPointer())

}
func (this *QMdiSubWindow) OnFocusOutEvent(slot func(super func(focusOutEvent *QFocusEvent), focusOutEvent *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_FocusOutEvent
func miqt_exec_callback_QMdiSubWindow_FocusOutEvent(self QMdiSubWindow, cb intptr_t, focusOutEvent *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(focusOutEvent *QFocusEvent), focusOutEvent *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(focusOutEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_ChildEvent(childEvent *QChildEvent) {

	QMdiSubWindow_virtualbase_ChildEvent(unsafe.Pointer(this.h), childEvent.cPointer())

}
func (this *QMdiSubWindow) OnChildEvent(slot func(super func(childEvent *QChildEvent), childEvent *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_ChildEvent
func miqt_exec_callback_QMdiSubWindow_ChildEvent(self QMdiSubWindow, cb intptr_t, childEvent *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(childEvent *QChildEvent), childEvent *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(childEvent)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_DevType() int {

	return (int)(QMdiSubWindow_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QMdiSubWindow) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_DevType
func miqt_exec_callback_QMdiSubWindow_DevType(self QMdiSubWindow, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QMdiSubWindow) callVirtualBase_SetVisible(visible bool) {

	QMdiSubWindow_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QMdiSubWindow) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_SetVisible
func miqt_exec_callback_QMdiSubWindow_SetVisible(self QMdiSubWindow, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QMdiSubWindow_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QMdiSubWindow) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_HeightForWidth
func miqt_exec_callback_QMdiSubWindow_HeightForWidth(self QMdiSubWindow, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QMdiSubWindow) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QMdiSubWindow_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QMdiSubWindow) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_HasHeightForWidth
func miqt_exec_callback_QMdiSubWindow_HasHeightForWidth(self QMdiSubWindow, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QMdiSubWindow) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QMdiSubWindow_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QMdiSubWindow) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_PaintEngine
func miqt_exec_callback_QMdiSubWindow_PaintEngine(self QMdiSubWindow, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QMdiSubWindow) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QMdiSubWindow_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_WheelEvent
func miqt_exec_callback_QMdiSubWindow_WheelEvent(self QMdiSubWindow, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QMdiSubWindow_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_KeyReleaseEvent
func miqt_exec_callback_QMdiSubWindow_KeyReleaseEvent(self QMdiSubWindow, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QMdiSubWindow_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_EnterEvent
func miqt_exec_callback_QMdiSubWindow_EnterEvent(self QMdiSubWindow, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QMdiSubWindow_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_TabletEvent
func miqt_exec_callback_QMdiSubWindow_TabletEvent(self QMdiSubWindow, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_ActionEvent(event *QActionEvent) {

	QMdiSubWindow_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_ActionEvent
func miqt_exec_callback_QMdiSubWindow_ActionEvent(self QMdiSubWindow, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QMdiSubWindow_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_DragEnterEvent
func miqt_exec_callback_QMdiSubWindow_DragEnterEvent(self QMdiSubWindow, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QMdiSubWindow_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_DragMoveEvent
func miqt_exec_callback_QMdiSubWindow_DragMoveEvent(self QMdiSubWindow, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QMdiSubWindow_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_DragLeaveEvent
func miqt_exec_callback_QMdiSubWindow_DragLeaveEvent(self QMdiSubWindow, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_DropEvent(event *QDropEvent) {

	QMdiSubWindow_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMdiSubWindow) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_DropEvent
func miqt_exec_callback_QMdiSubWindow_DropEvent(self QMdiSubWindow, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QMdiSubWindow_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QMdiSubWindow) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_NativeEvent
func miqt_exec_callback_QMdiSubWindow_NativeEvent(self QMdiSubWindow, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QMdiSubWindow) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QMdiSubWindow_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QMdiSubWindow) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_Metric
func miqt_exec_callback_QMdiSubWindow_Metric(self QMdiSubWindow, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QMdiSubWindow) callVirtualBase_InitPainter(painter *QPainter) {

	QMdiSubWindow_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QMdiSubWindow) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_InitPainter
func miqt_exec_callback_QMdiSubWindow_InitPainter(self QMdiSubWindow, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QMdiSubWindow_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QMdiSubWindow) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_Redirected
func miqt_exec_callback_QMdiSubWindow_Redirected(self QMdiSubWindow, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QMdiSubWindow) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QMdiSubWindow_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QMdiSubWindow) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_SharedPainter
func miqt_exec_callback_QMdiSubWindow_SharedPainter(self QMdiSubWindow, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QMdiSubWindow) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QMdiSubWindow_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMdiSubWindow) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_InputMethodEvent
func miqt_exec_callback_QMdiSubWindow_InputMethodEvent(self QMdiSubWindow, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QMdiSubWindow{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QMdiSubWindow) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QMdiSubWindow_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMdiSubWindow) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_InputMethodQuery
func miqt_exec_callback_QMdiSubWindow_InputMethodQuery(self QMdiSubWindow, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QMdiSubWindow) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QMdiSubWindow_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QMdiSubWindow) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_FocusNextPrevChild
func miqt_exec_callback_QMdiSubWindow_FocusNextPrevChild(self QMdiSubWindow, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
