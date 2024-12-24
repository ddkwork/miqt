package qt

import (
	"unsafe"
)

type QSizeGrip struct {
	h          uintptr
	isSubclass bool
}

// NewQSizeGrip constructs a new QSizeGrip object.
func NewQSizeGrip(parent *QWidget) *QSizeGrip {

	ret := newQSizeGrip(QSizeGrip_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSizeGrip) MetaObject() *QMetaObject {
	return newQMetaObject(QSizeGrip_MetaObject(this.h))
}

func (this *QSizeGrip) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSizeGrip_Metacast(this.h, param1_Cstring))
}

func QSizeGrip_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSizeGrip_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSizeGrip) SizeHint() *QSize {
	_goptr := newQSize(QSizeGrip_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeGrip) SetVisible(visible bool) {
	QSizeGrip_SetVisible(this.h, (bool)(visible))
}

func QSizeGrip_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSizeGrip_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSizeGrip_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSizeGrip_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSizeGrip) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QSizeGrip_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSizeGrip) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_SizeHint
func miqt_exec_callback_QSizeGrip_SizeHint(self QSizeGrip, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QSizeGrip) callVirtualBase_SetVisible(visible bool) {

	QSizeGrip_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QSizeGrip) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_SetVisible
func miqt_exec_callback_QSizeGrip_SetVisible(self QSizeGrip, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QSizeGrip) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QSizeGrip_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSizeGrip) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_PaintEvent
func miqt_exec_callback_QSizeGrip_PaintEvent(self QSizeGrip, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QSizeGrip_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSizeGrip) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_MousePressEvent
func miqt_exec_callback_QSizeGrip_MousePressEvent(self QSizeGrip, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QSizeGrip_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSizeGrip) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_MouseMoveEvent
func miqt_exec_callback_QSizeGrip_MouseMoveEvent(self QSizeGrip, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_MouseReleaseEvent(mouseEvent *QMouseEvent) {

	QSizeGrip_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), mouseEvent.cPointer())

}
func (this *QSizeGrip) OnMouseReleaseEvent(slot func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_MouseReleaseEvent
func miqt_exec_callback_QSizeGrip_MouseReleaseEvent(self QSizeGrip, cb intptr_t, mouseEvent *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mouseEvent *QMouseEvent), mouseEvent *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(mouseEvent)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_MoveEvent(moveEvent *QMoveEvent) {

	QSizeGrip_virtualbase_MoveEvent(unsafe.Pointer(this.h), moveEvent.cPointer())

}
func (this *QSizeGrip) OnMoveEvent(slot func(super func(moveEvent *QMoveEvent), moveEvent *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_MoveEvent
func miqt_exec_callback_QSizeGrip_MoveEvent(self QSizeGrip, cb intptr_t, moveEvent *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(moveEvent *QMoveEvent), moveEvent *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(moveEvent)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_ShowEvent(showEvent *QShowEvent) {

	QSizeGrip_virtualbase_ShowEvent(unsafe.Pointer(this.h), showEvent.cPointer())

}
func (this *QSizeGrip) OnShowEvent(slot func(super func(showEvent *QShowEvent), showEvent *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_ShowEvent
func miqt_exec_callback_QSizeGrip_ShowEvent(self QSizeGrip, cb intptr_t, showEvent *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(showEvent *QShowEvent), showEvent *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(showEvent)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_HideEvent(hideEvent *QHideEvent) {

	QSizeGrip_virtualbase_HideEvent(unsafe.Pointer(this.h), hideEvent.cPointer())

}
func (this *QSizeGrip) OnHideEvent(slot func(super func(hideEvent *QHideEvent), hideEvent *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_HideEvent
func miqt_exec_callback_QSizeGrip_HideEvent(self QSizeGrip, cb intptr_t, hideEvent *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(hideEvent *QHideEvent), hideEvent *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(hideEvent)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QSizeGrip_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QSizeGrip) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_EventFilter
func miqt_exec_callback_QSizeGrip_EventFilter(self QSizeGrip, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSizeGrip) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QSizeGrip_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QSizeGrip) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_Event
func miqt_exec_callback_QSizeGrip_Event(self QSizeGrip, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSizeGrip) callVirtualBase_DevType() int {

	return (int)(QSizeGrip_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QSizeGrip) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_DevType
func miqt_exec_callback_QSizeGrip_DevType(self QSizeGrip, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QSizeGrip) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QSizeGrip_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSizeGrip) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_MinimumSizeHint
func miqt_exec_callback_QSizeGrip_MinimumSizeHint(self QSizeGrip, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QSizeGrip) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QSizeGrip_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QSizeGrip) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_HeightForWidth
func miqt_exec_callback_QSizeGrip_HeightForWidth(self QSizeGrip, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QSizeGrip) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QSizeGrip_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QSizeGrip) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_HasHeightForWidth
func miqt_exec_callback_QSizeGrip_HasHeightForWidth(self QSizeGrip, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QSizeGrip) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QSizeGrip_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QSizeGrip) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_PaintEngine
func miqt_exec_callback_QSizeGrip_PaintEngine(self QSizeGrip, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QSizeGrip) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QSizeGrip_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_MouseDoubleClickEvent
func miqt_exec_callback_QSizeGrip_MouseDoubleClickEvent(self QSizeGrip, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QSizeGrip_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_WheelEvent
func miqt_exec_callback_QSizeGrip_WheelEvent(self QSizeGrip, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QSizeGrip_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_KeyPressEvent
func miqt_exec_callback_QSizeGrip_KeyPressEvent(self QSizeGrip, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QSizeGrip_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_KeyReleaseEvent
func miqt_exec_callback_QSizeGrip_KeyReleaseEvent(self QSizeGrip, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QSizeGrip_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_FocusInEvent
func miqt_exec_callback_QSizeGrip_FocusInEvent(self QSizeGrip, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QSizeGrip_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_FocusOutEvent
func miqt_exec_callback_QSizeGrip_FocusOutEvent(self QSizeGrip, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QSizeGrip_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_EnterEvent
func miqt_exec_callback_QSizeGrip_EnterEvent(self QSizeGrip, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_LeaveEvent(event *QEvent) {

	QSizeGrip_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_LeaveEvent
func miqt_exec_callback_QSizeGrip_LeaveEvent(self QSizeGrip, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QSizeGrip_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_ResizeEvent
func miqt_exec_callback_QSizeGrip_ResizeEvent(self QSizeGrip, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QSizeGrip_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_CloseEvent
func miqt_exec_callback_QSizeGrip_CloseEvent(self QSizeGrip, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QSizeGrip_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_ContextMenuEvent
func miqt_exec_callback_QSizeGrip_ContextMenuEvent(self QSizeGrip, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QSizeGrip_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_TabletEvent
func miqt_exec_callback_QSizeGrip_TabletEvent(self QSizeGrip, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_ActionEvent(event *QActionEvent) {

	QSizeGrip_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_ActionEvent
func miqt_exec_callback_QSizeGrip_ActionEvent(self QSizeGrip, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QSizeGrip_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_DragEnterEvent
func miqt_exec_callback_QSizeGrip_DragEnterEvent(self QSizeGrip, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QSizeGrip_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_DragMoveEvent
func miqt_exec_callback_QSizeGrip_DragMoveEvent(self QSizeGrip, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QSizeGrip_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_DragLeaveEvent
func miqt_exec_callback_QSizeGrip_DragLeaveEvent(self QSizeGrip, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_DropEvent(event *QDropEvent) {

	QSizeGrip_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSizeGrip) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_DropEvent
func miqt_exec_callback_QSizeGrip_DropEvent(self QSizeGrip, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QSizeGrip_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QSizeGrip) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_NativeEvent
func miqt_exec_callback_QSizeGrip_NativeEvent(self QSizeGrip, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QSizeGrip) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QSizeGrip_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSizeGrip) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_ChangeEvent
func miqt_exec_callback_QSizeGrip_ChangeEvent(self QSizeGrip, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QSizeGrip_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QSizeGrip) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_Metric
func miqt_exec_callback_QSizeGrip_Metric(self QSizeGrip, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QSizeGrip) callVirtualBase_InitPainter(painter *QPainter) {

	QSizeGrip_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QSizeGrip) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_InitPainter
func miqt_exec_callback_QSizeGrip_InitPainter(self QSizeGrip, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QSizeGrip) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QSizeGrip_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QSizeGrip) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_Redirected
func miqt_exec_callback_QSizeGrip_Redirected(self QSizeGrip, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QSizeGrip) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QSizeGrip_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QSizeGrip) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_SharedPainter
func miqt_exec_callback_QSizeGrip_SharedPainter(self QSizeGrip, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QSizeGrip) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QSizeGrip_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSizeGrip) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_InputMethodEvent
func miqt_exec_callback_QSizeGrip_InputMethodEvent(self QSizeGrip, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QSizeGrip{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QSizeGrip) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QSizeGrip_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSizeGrip) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_InputMethodQuery
func miqt_exec_callback_QSizeGrip_InputMethodQuery(self QSizeGrip, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QSizeGrip) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QSizeGrip_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QSizeGrip) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_FocusNextPrevChild
func miqt_exec_callback_QSizeGrip_FocusNextPrevChild(self QSizeGrip, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
