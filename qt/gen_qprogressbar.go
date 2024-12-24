package qt

import (
	"unsafe"
)

type QProgressBar__Direction int

const (
	QProgressBar__TopToBottom QProgressBar__Direction = 0
	QProgressBar__BottomToTop QProgressBar__Direction = 1
)

type QProgressBar struct {
	h          uintptr
	isSubclass bool
}

// NewQProgressBar constructs a new QProgressBar object.
func NewQProgressBar(parent *QWidget) *QProgressBar {

	ret := newQProgressBar(QProgressBar_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQProgressBar2 constructs a new QProgressBar object.
func NewQProgressBar2() *QProgressBar {

	ret := newQProgressBar(QProgressBar_new2())
	ret.isSubclass = true
	return ret
}

func (this *QProgressBar) MetaObject() *QMetaObject {
	return newQMetaObject(QProgressBar_MetaObject(this.h))
}

func (this *QProgressBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QProgressBar_Metacast(this.h, param1_Cstring))
}

func QProgressBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QProgressBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) Minimum() int {
	return (int)(QProgressBar_Minimum(this.h))
}

func (this *QProgressBar) Maximum() int {
	return (int)(QProgressBar_Maximum(this.h))
}

func (this *QProgressBar) Value() int {
	return (int)(QProgressBar_Value(this.h))
}

func (this *QProgressBar) Text() string {
	var _ms struct_miqt_string = QProgressBar_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) SetTextVisible(visible bool) {
	QProgressBar_SetTextVisible(this.h, (bool)(visible))
}

func (this *QProgressBar) IsTextVisible() bool {
	return (bool)(QProgressBar_IsTextVisible(this.h))
}

func (this *QProgressBar) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QProgressBar_Alignment(this.h))
}

func (this *QProgressBar) SetAlignment(alignment AlignmentFlag) {
	QProgressBar_SetAlignment(this.h, (int)(alignment))
}

func (this *QProgressBar) SizeHint() *QSize {
	_goptr := newQSize(QProgressBar_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProgressBar) MinimumSizeHint() *QSize {
	_goptr := newQSize(QProgressBar_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProgressBar) Orientation() Orientation {
	return (Orientation)(QProgressBar_Orientation(this.h))
}

func (this *QProgressBar) SetInvertedAppearance(invert bool) {
	QProgressBar_SetInvertedAppearance(this.h, (bool)(invert))
}

func (this *QProgressBar) InvertedAppearance() bool {
	return (bool)(QProgressBar_InvertedAppearance(this.h))
}

func (this *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	QProgressBar_SetTextDirection(this.h, (int)(textDirection))
}

func (this *QProgressBar) TextDirection() QProgressBar__Direction {
	return (QProgressBar__Direction)(QProgressBar_TextDirection(this.h))
}

func (this *QProgressBar) SetFormat(format string) {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	QProgressBar_SetFormat(this.h, format_ms)
}

func (this *QProgressBar) ResetFormat() {
	QProgressBar_ResetFormat(this.h)
}

func (this *QProgressBar) Format() string {
	var _ms struct_miqt_string = QProgressBar_Format(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) Reset() {
	QProgressBar_Reset(this.h)
}

func (this *QProgressBar) SetRange(minimum int, maximum int) {
	QProgressBar_SetRange(this.h, (int)(minimum), (int)(maximum))
}

func (this *QProgressBar) SetMinimum(minimum int) {
	QProgressBar_SetMinimum(this.h, (int)(minimum))
}

func (this *QProgressBar) SetMaximum(maximum int) {
	QProgressBar_SetMaximum(this.h, (int)(maximum))
}

func (this *QProgressBar) SetValue(value int) {
	QProgressBar_SetValue(this.h, (int)(value))
}

func (this *QProgressBar) SetOrientation(orientation Orientation) {
	QProgressBar_SetOrientation(this.h, (int)(orientation))
}

func (this *QProgressBar) ValueChanged(value int) {
	QProgressBar_ValueChanged(this.h, (int)(value))
}
func (this *QProgressBar) OnValueChanged(slot func(value int)) {
	QProgressBar_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_ValueChanged
func miqt_exec_callback_QProgressBar_ValueChanged(cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc(slotval1)
}

func QProgressBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProgressBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProgressBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProgressBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) callVirtualBase_Text() string {

	var _ms struct_miqt_string = QProgressBar_virtualbase_Text(unsafe.Pointer(this.h))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QProgressBar) OnText(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_Text(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_Text
func miqt_exec_callback_QProgressBar_Text(self QProgressBar, cb intptr_t) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_Text)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QProgressBar) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QProgressBar_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProgressBar) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_SizeHint
func miqt_exec_callback_QProgressBar_SizeHint(self QProgressBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QProgressBar) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QProgressBar_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProgressBar) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_MinimumSizeHint
func miqt_exec_callback_QProgressBar_MinimumSizeHint(self QProgressBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QProgressBar) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QProgressBar_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QProgressBar) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_Event
func miqt_exec_callback_QProgressBar_Event(self QProgressBar, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QProgressBar) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QProgressBar_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QProgressBar) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_PaintEvent
func miqt_exec_callback_QProgressBar_PaintEvent(self QProgressBar, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QProgressBar{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_InitStyleOption(option *QStyleOptionProgressBar) {

	QProgressBar_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QProgressBar) OnInitStyleOption(slot func(super func(option *QStyleOptionProgressBar), option *QStyleOptionProgressBar)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_InitStyleOption
func miqt_exec_callback_QProgressBar_InitStyleOption(self QProgressBar, cb intptr_t, option *QStyleOptionProgressBar) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionProgressBar), option *QStyleOptionProgressBar))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionProgressBar(option)

	gofunc((&QProgressBar{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QProgressBar) callVirtualBase_DevType() int {

	return (int)(QProgressBar_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QProgressBar) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_DevType
func miqt_exec_callback_QProgressBar_DevType(self QProgressBar, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QProgressBar) callVirtualBase_SetVisible(visible bool) {

	QProgressBar_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QProgressBar) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_SetVisible
func miqt_exec_callback_QProgressBar_SetVisible(self QProgressBar, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QProgressBar{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QProgressBar) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QProgressBar_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QProgressBar) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_HeightForWidth
func miqt_exec_callback_QProgressBar_HeightForWidth(self QProgressBar, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QProgressBar) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QProgressBar_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QProgressBar) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_HasHeightForWidth
func miqt_exec_callback_QProgressBar_HasHeightForWidth(self QProgressBar, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QProgressBar) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QProgressBar_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QProgressBar) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_PaintEngine
func miqt_exec_callback_QProgressBar_PaintEngine(self QProgressBar, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QProgressBar) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QProgressBar_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_MousePressEvent
func miqt_exec_callback_QProgressBar_MousePressEvent(self QProgressBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QProgressBar_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_MouseReleaseEvent
func miqt_exec_callback_QProgressBar_MouseReleaseEvent(self QProgressBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QProgressBar_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_MouseDoubleClickEvent
func miqt_exec_callback_QProgressBar_MouseDoubleClickEvent(self QProgressBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QProgressBar_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_MouseMoveEvent
func miqt_exec_callback_QProgressBar_MouseMoveEvent(self QProgressBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QProgressBar_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_WheelEvent
func miqt_exec_callback_QProgressBar_WheelEvent(self QProgressBar, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QProgressBar_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_KeyPressEvent
func miqt_exec_callback_QProgressBar_KeyPressEvent(self QProgressBar, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QProgressBar_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_KeyReleaseEvent
func miqt_exec_callback_QProgressBar_KeyReleaseEvent(self QProgressBar, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QProgressBar_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_FocusInEvent
func miqt_exec_callback_QProgressBar_FocusInEvent(self QProgressBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QProgressBar_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_FocusOutEvent
func miqt_exec_callback_QProgressBar_FocusOutEvent(self QProgressBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QProgressBar_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_EnterEvent
func miqt_exec_callback_QProgressBar_EnterEvent(self QProgressBar, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_LeaveEvent(event *QEvent) {

	QProgressBar_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_LeaveEvent
func miqt_exec_callback_QProgressBar_LeaveEvent(self QProgressBar, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QProgressBar_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_MoveEvent
func miqt_exec_callback_QProgressBar_MoveEvent(self QProgressBar, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QProgressBar_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_ResizeEvent
func miqt_exec_callback_QProgressBar_ResizeEvent(self QProgressBar, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QProgressBar_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_CloseEvent
func miqt_exec_callback_QProgressBar_CloseEvent(self QProgressBar, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QProgressBar_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_ContextMenuEvent
func miqt_exec_callback_QProgressBar_ContextMenuEvent(self QProgressBar, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QProgressBar_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_TabletEvent
func miqt_exec_callback_QProgressBar_TabletEvent(self QProgressBar, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_ActionEvent(event *QActionEvent) {

	QProgressBar_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_ActionEvent
func miqt_exec_callback_QProgressBar_ActionEvent(self QProgressBar, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QProgressBar_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_DragEnterEvent
func miqt_exec_callback_QProgressBar_DragEnterEvent(self QProgressBar, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QProgressBar_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_DragMoveEvent
func miqt_exec_callback_QProgressBar_DragMoveEvent(self QProgressBar, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QProgressBar_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_DragLeaveEvent
func miqt_exec_callback_QProgressBar_DragLeaveEvent(self QProgressBar, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_DropEvent(event *QDropEvent) {

	QProgressBar_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_DropEvent
func miqt_exec_callback_QProgressBar_DropEvent(self QProgressBar, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_ShowEvent(event *QShowEvent) {

	QProgressBar_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_ShowEvent
func miqt_exec_callback_QProgressBar_ShowEvent(self QProgressBar, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_HideEvent(event *QHideEvent) {

	QProgressBar_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressBar) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_HideEvent
func miqt_exec_callback_QProgressBar_HideEvent(self QProgressBar, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QProgressBar{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QProgressBar_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QProgressBar) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_NativeEvent
func miqt_exec_callback_QProgressBar_NativeEvent(self QProgressBar, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QProgressBar) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QProgressBar_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QProgressBar) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_ChangeEvent
func miqt_exec_callback_QProgressBar_ChangeEvent(self QProgressBar, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QProgressBar{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QProgressBar_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QProgressBar) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_Metric
func miqt_exec_callback_QProgressBar_Metric(self QProgressBar, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QProgressBar) callVirtualBase_InitPainter(painter *QPainter) {

	QProgressBar_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QProgressBar) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_InitPainter
func miqt_exec_callback_QProgressBar_InitPainter(self QProgressBar, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QProgressBar{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QProgressBar) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QProgressBar_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QProgressBar) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_Redirected
func miqt_exec_callback_QProgressBar_Redirected(self QProgressBar, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QProgressBar) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QProgressBar_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QProgressBar) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_SharedPainter
func miqt_exec_callback_QProgressBar_SharedPainter(self QProgressBar, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QProgressBar) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QProgressBar_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QProgressBar) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_InputMethodEvent
func miqt_exec_callback_QProgressBar_InputMethodEvent(self QProgressBar, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QProgressBar{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QProgressBar) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QProgressBar_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProgressBar) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_InputMethodQuery
func miqt_exec_callback_QProgressBar_InputMethodQuery(self QProgressBar, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QProgressBar) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QProgressBar_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QProgressBar) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_FocusNextPrevChild
func miqt_exec_callback_QProgressBar_FocusNextPrevChild(self QProgressBar, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
