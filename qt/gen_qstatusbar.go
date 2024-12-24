package qt

import (
	"unsafe"
)

type QStatusBar struct {
	h          uintptr
	isSubclass bool
}

// NewQStatusBar constructs a new QStatusBar object.
func NewQStatusBar(parent *QWidget) *QStatusBar {

	ret := newQStatusBar(QStatusBar_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStatusBar2 constructs a new QStatusBar object.
func NewQStatusBar2() *QStatusBar {

	ret := newQStatusBar(QStatusBar_new2())
	ret.isSubclass = true
	return ret
}

func (this *QStatusBar) MetaObject() *QMetaObject {
	return newQMetaObject(QStatusBar_MetaObject(this.h))
}

func (this *QStatusBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStatusBar_Metacast(this.h, param1_Cstring))
}

func QStatusBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStatusBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStatusBar) AddWidget(widget *QWidget) {
	QStatusBar_AddWidget(this.h, widget.cPointer())
}

func (this *QStatusBar) InsertWidget(index int, widget *QWidget) int {
	return (int)(QStatusBar_InsertWidget(this.h, (int)(index), widget.cPointer()))
}

func (this *QStatusBar) AddPermanentWidget(widget *QWidget) {
	QStatusBar_AddPermanentWidget(this.h, widget.cPointer())
}

func (this *QStatusBar) InsertPermanentWidget(index int, widget *QWidget) int {
	return (int)(QStatusBar_InsertPermanentWidget(this.h, (int)(index), widget.cPointer()))
}

func (this *QStatusBar) RemoveWidget(widget *QWidget) {
	QStatusBar_RemoveWidget(this.h, widget.cPointer())
}

func (this *QStatusBar) SetSizeGripEnabled(sizeGripEnabled bool) {
	QStatusBar_SetSizeGripEnabled(this.h, (bool)(sizeGripEnabled))
}

func (this *QStatusBar) IsSizeGripEnabled() bool {
	return (bool)(QStatusBar_IsSizeGripEnabled(this.h))
}

func (this *QStatusBar) CurrentMessage() string {
	var _ms struct_miqt_string = QStatusBar_CurrentMessage(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStatusBar) ShowMessage(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStatusBar_ShowMessage(this.h, text_ms)
}

func (this *QStatusBar) ClearMessage() {
	QStatusBar_ClearMessage(this.h)
}

func (this *QStatusBar) MessageChanged(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStatusBar_MessageChanged(this.h, text_ms)
}
func (this *QStatusBar) OnMessageChanged(slot func(text string)) {
	QStatusBar_connect_MessageChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MessageChanged
func miqt_exec_callback_QStatusBar_MessageChanged(cb intptr_t, text struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)
}

func QStatusBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStatusBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStatusBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStatusBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStatusBar) AddWidget2(widget *QWidget, stretch int) {
	QStatusBar_AddWidget2(this.h, widget.cPointer(), (int)(stretch))
}

func (this *QStatusBar) InsertWidget3(index int, widget *QWidget, stretch int) int {
	return (int)(QStatusBar_InsertWidget3(this.h, (int)(index), widget.cPointer(), (int)(stretch)))
}

func (this *QStatusBar) AddPermanentWidget2(widget *QWidget, stretch int) {
	QStatusBar_AddPermanentWidget2(this.h, widget.cPointer(), (int)(stretch))
}

func (this *QStatusBar) InsertPermanentWidget3(index int, widget *QWidget, stretch int) int {
	return (int)(QStatusBar_InsertPermanentWidget3(this.h, (int)(index), widget.cPointer(), (int)(stretch)))
}

func (this *QStatusBar) ShowMessage2(text string, timeout int) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStatusBar_ShowMessage2(this.h, text_ms, (int)(timeout))
}

func (this *QStatusBar) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QStatusBar_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QStatusBar) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_ShowEvent
func miqt_exec_callback_QStatusBar_ShowEvent(self QStatusBar, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QStatusBar{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QStatusBar_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QStatusBar) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_PaintEvent
func miqt_exec_callback_QStatusBar_PaintEvent(self QStatusBar, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QStatusBar{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QStatusBar_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QStatusBar) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_ResizeEvent
func miqt_exec_callback_QStatusBar_ResizeEvent(self QStatusBar, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QStatusBar{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QStatusBar_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QStatusBar) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_Event
func miqt_exec_callback_QStatusBar_Event(self QStatusBar, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QStatusBar) callVirtualBase_DevType() int {

	return (int)(QStatusBar_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QStatusBar) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_DevType
func miqt_exec_callback_QStatusBar_DevType(self QStatusBar, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QStatusBar) callVirtualBase_SetVisible(visible bool) {

	QStatusBar_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QStatusBar) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_SetVisible
func miqt_exec_callback_QStatusBar_SetVisible(self QStatusBar, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QStatusBar{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QStatusBar) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QStatusBar_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStatusBar) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_SizeHint
func miqt_exec_callback_QStatusBar_SizeHint(self QStatusBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QStatusBar) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QStatusBar_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStatusBar) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MinimumSizeHint
func miqt_exec_callback_QStatusBar_MinimumSizeHint(self QStatusBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QStatusBar) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QStatusBar_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QStatusBar) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_HeightForWidth
func miqt_exec_callback_QStatusBar_HeightForWidth(self QStatusBar, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QStatusBar) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QStatusBar_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QStatusBar) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_HasHeightForWidth
func miqt_exec_callback_QStatusBar_HasHeightForWidth(self QStatusBar, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QStatusBar) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QStatusBar_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QStatusBar) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_PaintEngine
func miqt_exec_callback_QStatusBar_PaintEngine(self QStatusBar, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QStatusBar) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QStatusBar_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MousePressEvent
func miqt_exec_callback_QStatusBar_MousePressEvent(self QStatusBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QStatusBar_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MouseReleaseEvent
func miqt_exec_callback_QStatusBar_MouseReleaseEvent(self QStatusBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QStatusBar_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MouseDoubleClickEvent
func miqt_exec_callback_QStatusBar_MouseDoubleClickEvent(self QStatusBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QStatusBar_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MouseMoveEvent
func miqt_exec_callback_QStatusBar_MouseMoveEvent(self QStatusBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QStatusBar_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_WheelEvent
func miqt_exec_callback_QStatusBar_WheelEvent(self QStatusBar, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QStatusBar_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_KeyPressEvent
func miqt_exec_callback_QStatusBar_KeyPressEvent(self QStatusBar, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QStatusBar_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_KeyReleaseEvent
func miqt_exec_callback_QStatusBar_KeyReleaseEvent(self QStatusBar, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QStatusBar_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_FocusInEvent
func miqt_exec_callback_QStatusBar_FocusInEvent(self QStatusBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QStatusBar_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_FocusOutEvent
func miqt_exec_callback_QStatusBar_FocusOutEvent(self QStatusBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QStatusBar_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_EnterEvent
func miqt_exec_callback_QStatusBar_EnterEvent(self QStatusBar, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_LeaveEvent(event *QEvent) {

	QStatusBar_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_LeaveEvent
func miqt_exec_callback_QStatusBar_LeaveEvent(self QStatusBar, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QStatusBar_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MoveEvent
func miqt_exec_callback_QStatusBar_MoveEvent(self QStatusBar, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QStatusBar_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_CloseEvent
func miqt_exec_callback_QStatusBar_CloseEvent(self QStatusBar, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QStatusBar_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_ContextMenuEvent
func miqt_exec_callback_QStatusBar_ContextMenuEvent(self QStatusBar, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QStatusBar_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_TabletEvent
func miqt_exec_callback_QStatusBar_TabletEvent(self QStatusBar, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_ActionEvent(event *QActionEvent) {

	QStatusBar_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_ActionEvent
func miqt_exec_callback_QStatusBar_ActionEvent(self QStatusBar, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QStatusBar_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_DragEnterEvent
func miqt_exec_callback_QStatusBar_DragEnterEvent(self QStatusBar, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QStatusBar_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_DragMoveEvent
func miqt_exec_callback_QStatusBar_DragMoveEvent(self QStatusBar, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QStatusBar_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_DragLeaveEvent
func miqt_exec_callback_QStatusBar_DragLeaveEvent(self QStatusBar, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_DropEvent(event *QDropEvent) {

	QStatusBar_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_DropEvent
func miqt_exec_callback_QStatusBar_DropEvent(self QStatusBar, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_HideEvent(event *QHideEvent) {

	QStatusBar_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QStatusBar) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_HideEvent
func miqt_exec_callback_QStatusBar_HideEvent(self QStatusBar, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QStatusBar{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QStatusBar_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QStatusBar) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_NativeEvent
func miqt_exec_callback_QStatusBar_NativeEvent(self QStatusBar, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStatusBar) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QStatusBar_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QStatusBar) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_ChangeEvent
func miqt_exec_callback_QStatusBar_ChangeEvent(self QStatusBar, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QStatusBar{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QStatusBar_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QStatusBar) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_Metric
func miqt_exec_callback_QStatusBar_Metric(self QStatusBar, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QStatusBar) callVirtualBase_InitPainter(painter *QPainter) {

	QStatusBar_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QStatusBar) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_InitPainter
func miqt_exec_callback_QStatusBar_InitPainter(self QStatusBar, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QStatusBar{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QStatusBar) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QStatusBar_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QStatusBar) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_Redirected
func miqt_exec_callback_QStatusBar_Redirected(self QStatusBar, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStatusBar) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QStatusBar_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QStatusBar) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_SharedPainter
func miqt_exec_callback_QStatusBar_SharedPainter(self QStatusBar, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QStatusBar) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QStatusBar_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QStatusBar) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_InputMethodEvent
func miqt_exec_callback_QStatusBar_InputMethodEvent(self QStatusBar, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QStatusBar{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QStatusBar) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QStatusBar_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStatusBar) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_InputMethodQuery
func miqt_exec_callback_QStatusBar_InputMethodQuery(self QStatusBar, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStatusBar) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QStatusBar_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QStatusBar) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_FocusNextPrevChild
func miqt_exec_callback_QStatusBar_FocusNextPrevChild(self QStatusBar, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
