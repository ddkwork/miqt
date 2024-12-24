package qt

import (
	"unsafe"
)

type QRubberBand__Shape int

const (
	QRubberBand__Line      QRubberBand__Shape = 0
	QRubberBand__Rectangle QRubberBand__Shape = 1
)

type QRubberBand struct {
	h          uintptr
	isSubclass bool
}

// NewQRubberBand constructs a new QRubberBand object.
func NewQRubberBand(param1 Shape) *QRubberBand {

	ret := newQRubberBand(QRubberBand_new(param1))
	ret.isSubclass = true
	return ret
}

// NewQRubberBand2 constructs a new QRubberBand object.
func NewQRubberBand2(param1 Shape, param2 *QWidget) *QRubberBand {

	ret := newQRubberBand(QRubberBand_new2(param1, param2.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRubberBand) MetaObject() *QMetaObject {
	return newQMetaObject(QRubberBand_MetaObject(this.h))
}

func (this *QRubberBand) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRubberBand_Metacast(this.h, param1_Cstring))
}

func QRubberBand_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRubberBand_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRubberBand) Shape() Shape {
	xxxxxxxxx
}

func (this *QRubberBand) SetGeometry(r *QRect) {
	QRubberBand_SetGeometry(this.h, r.cPointer())
}

func (this *QRubberBand) SetGeometry2(x int, y int, w int, h int) {
	QRubberBand_SetGeometry2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QRubberBand) Move(x int, y int) {
	QRubberBand_Move(this.h, (int)(x), (int)(y))
}

func (this *QRubberBand) MoveWithQPoint(p *QPoint) {
	QRubberBand_MoveWithQPoint(this.h, p.cPointer())
}

func (this *QRubberBand) Resize(w int, h int) {
	QRubberBand_Resize(this.h, (int)(w), (int)(h))
}

func (this *QRubberBand) ResizeWithQSize(s *QSize) {
	QRubberBand_ResizeWithQSize(this.h, s.cPointer())
}

func QRubberBand_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRubberBand_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRubberBand_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRubberBand_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRubberBand) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QRubberBand_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QRubberBand) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_Event
func miqt_exec_callback_QRubberBand_Event(self QRubberBand, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QRubberBand) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QRubberBand_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRubberBand) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_PaintEvent
func miqt_exec_callback_QRubberBand_PaintEvent(self QRubberBand, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QRubberBand{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QRubberBand_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRubberBand) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_ChangeEvent
func miqt_exec_callback_QRubberBand_ChangeEvent(self QRubberBand, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QRubberBand{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QRubberBand_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRubberBand) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_ShowEvent
func miqt_exec_callback_QRubberBand_ShowEvent(self QRubberBand, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QRubberBand{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QRubberBand_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRubberBand) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_ResizeEvent
func miqt_exec_callback_QRubberBand_ResizeEvent(self QRubberBand, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QRubberBand{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_MoveEvent(param1 *QMoveEvent) {

	QRubberBand_virtualbase_MoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRubberBand) OnMoveEvent(slot func(super func(param1 *QMoveEvent), param1 *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_MoveEvent
func miqt_exec_callback_QRubberBand_MoveEvent(self QRubberBand, cb intptr_t, param1 *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMoveEvent), param1 *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(param1)

	gofunc((&QRubberBand{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_InitStyleOption(option *QStyleOptionRubberBand) {

	QRubberBand_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QRubberBand) OnInitStyleOption(slot func(super func(option *QStyleOptionRubberBand), option *QStyleOptionRubberBand)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_InitStyleOption
func miqt_exec_callback_QRubberBand_InitStyleOption(self QRubberBand, cb intptr_t, option *QStyleOptionRubberBand) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionRubberBand), option *QStyleOptionRubberBand))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionRubberBand(option)

	gofunc((&QRubberBand{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QRubberBand) callVirtualBase_DevType() int {

	return (int)(QRubberBand_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QRubberBand) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_DevType
func miqt_exec_callback_QRubberBand_DevType(self QRubberBand, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QRubberBand) callVirtualBase_SetVisible(visible bool) {

	QRubberBand_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QRubberBand) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_SetVisible
func miqt_exec_callback_QRubberBand_SetVisible(self QRubberBand, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QRubberBand{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QRubberBand) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QRubberBand_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QRubberBand) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_SizeHint
func miqt_exec_callback_QRubberBand_SizeHint(self QRubberBand, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QRubberBand) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QRubberBand_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QRubberBand) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_MinimumSizeHint
func miqt_exec_callback_QRubberBand_MinimumSizeHint(self QRubberBand, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QRubberBand) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QRubberBand_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QRubberBand) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_HeightForWidth
func miqt_exec_callback_QRubberBand_HeightForWidth(self QRubberBand, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QRubberBand) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QRubberBand_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QRubberBand) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_HasHeightForWidth
func miqt_exec_callback_QRubberBand_HasHeightForWidth(self QRubberBand, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QRubberBand) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QRubberBand_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QRubberBand) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_PaintEngine
func miqt_exec_callback_QRubberBand_PaintEngine(self QRubberBand, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QRubberBand) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QRubberBand_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_MousePressEvent
func miqt_exec_callback_QRubberBand_MousePressEvent(self QRubberBand, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QRubberBand_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_MouseReleaseEvent
func miqt_exec_callback_QRubberBand_MouseReleaseEvent(self QRubberBand, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QRubberBand_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_MouseDoubleClickEvent
func miqt_exec_callback_QRubberBand_MouseDoubleClickEvent(self QRubberBand, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QRubberBand_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_MouseMoveEvent
func miqt_exec_callback_QRubberBand_MouseMoveEvent(self QRubberBand, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QRubberBand_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_WheelEvent
func miqt_exec_callback_QRubberBand_WheelEvent(self QRubberBand, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QRubberBand_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_KeyPressEvent
func miqt_exec_callback_QRubberBand_KeyPressEvent(self QRubberBand, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QRubberBand_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_KeyReleaseEvent
func miqt_exec_callback_QRubberBand_KeyReleaseEvent(self QRubberBand, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QRubberBand_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_FocusInEvent
func miqt_exec_callback_QRubberBand_FocusInEvent(self QRubberBand, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QRubberBand_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_FocusOutEvent
func miqt_exec_callback_QRubberBand_FocusOutEvent(self QRubberBand, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QRubberBand_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_EnterEvent
func miqt_exec_callback_QRubberBand_EnterEvent(self QRubberBand, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_LeaveEvent(event *QEvent) {

	QRubberBand_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_LeaveEvent
func miqt_exec_callback_QRubberBand_LeaveEvent(self QRubberBand, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QRubberBand_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_CloseEvent
func miqt_exec_callback_QRubberBand_CloseEvent(self QRubberBand, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QRubberBand_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_ContextMenuEvent
func miqt_exec_callback_QRubberBand_ContextMenuEvent(self QRubberBand, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QRubberBand_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_TabletEvent
func miqt_exec_callback_QRubberBand_TabletEvent(self QRubberBand, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_ActionEvent(event *QActionEvent) {

	QRubberBand_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_ActionEvent
func miqt_exec_callback_QRubberBand_ActionEvent(self QRubberBand, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QRubberBand_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_DragEnterEvent
func miqt_exec_callback_QRubberBand_DragEnterEvent(self QRubberBand, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QRubberBand_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_DragMoveEvent
func miqt_exec_callback_QRubberBand_DragMoveEvent(self QRubberBand, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QRubberBand_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_DragLeaveEvent
func miqt_exec_callback_QRubberBand_DragLeaveEvent(self QRubberBand, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_DropEvent(event *QDropEvent) {

	QRubberBand_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_DropEvent
func miqt_exec_callback_QRubberBand_DropEvent(self QRubberBand, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_HideEvent(event *QHideEvent) {

	QRubberBand_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRubberBand) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_HideEvent
func miqt_exec_callback_QRubberBand_HideEvent(self QRubberBand, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QRubberBand{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QRubberBand_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QRubberBand) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_NativeEvent
func miqt_exec_callback_QRubberBand_NativeEvent(self QRubberBand, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QRubberBand) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QRubberBand_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QRubberBand) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_Metric
func miqt_exec_callback_QRubberBand_Metric(self QRubberBand, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QRubberBand) callVirtualBase_InitPainter(painter *QPainter) {

	QRubberBand_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QRubberBand) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_InitPainter
func miqt_exec_callback_QRubberBand_InitPainter(self QRubberBand, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QRubberBand{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QRubberBand) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QRubberBand_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QRubberBand) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_Redirected
func miqt_exec_callback_QRubberBand_Redirected(self QRubberBand, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QRubberBand) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QRubberBand_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QRubberBand) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_SharedPainter
func miqt_exec_callback_QRubberBand_SharedPainter(self QRubberBand, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QRubberBand) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QRubberBand_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRubberBand) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_InputMethodEvent
func miqt_exec_callback_QRubberBand_InputMethodEvent(self QRubberBand, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QRubberBand{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QRubberBand) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QRubberBand_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QRubberBand) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_InputMethodQuery
func miqt_exec_callback_QRubberBand_InputMethodQuery(self QRubberBand, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QRubberBand) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QRubberBand_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QRubberBand) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_FocusNextPrevChild
func miqt_exec_callback_QRubberBand_FocusNextPrevChild(self QRubberBand, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
