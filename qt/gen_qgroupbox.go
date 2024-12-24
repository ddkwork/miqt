package qt

import (
	"unsafe"
)

type QGroupBox struct {
	h          uintptr
	isSubclass bool
}

// NewQGroupBox constructs a new QGroupBox object.
func NewQGroupBox(parent *QWidget) *QGroupBox {

	ret := newQGroupBox(QGroupBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGroupBox2 constructs a new QGroupBox object.
func NewQGroupBox2() *QGroupBox {

	ret := newQGroupBox(QGroupBox_new2())
	ret.isSubclass = true
	return ret
}

// NewQGroupBox3 constructs a new QGroupBox object.
func NewQGroupBox3(title string) *QGroupBox {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))

	ret := newQGroupBox(QGroupBox_new3(title_ms))
	ret.isSubclass = true
	return ret
}

// NewQGroupBox4 constructs a new QGroupBox object.
func NewQGroupBox4(title string, parent *QWidget) *QGroupBox {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))

	ret := newQGroupBox(QGroupBox_new4(title_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGroupBox) MetaObject() *QMetaObject {
	return newQMetaObject(QGroupBox_MetaObject(this.h))
}

func (this *QGroupBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGroupBox_Metacast(this.h, param1_Cstring))
}

func QGroupBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGroupBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGroupBox) Title() string {
	var _ms struct_miqt_string = QGroupBox_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGroupBox) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QGroupBox_SetTitle(this.h, title_ms)
}

func (this *QGroupBox) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QGroupBox_Alignment(this.h))
}

func (this *QGroupBox) SetAlignment(alignment int) {
	QGroupBox_SetAlignment(this.h, (int)(alignment))
}

func (this *QGroupBox) MinimumSizeHint() *QSize {
	_goptr := newQSize(QGroupBox_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGroupBox) IsFlat() bool {
	return (bool)(QGroupBox_IsFlat(this.h))
}

func (this *QGroupBox) SetFlat(flat bool) {
	QGroupBox_SetFlat(this.h, (bool)(flat))
}

func (this *QGroupBox) IsCheckable() bool {
	return (bool)(QGroupBox_IsCheckable(this.h))
}

func (this *QGroupBox) SetCheckable(checkable bool) {
	QGroupBox_SetCheckable(this.h, (bool)(checkable))
}

func (this *QGroupBox) IsChecked() bool {
	return (bool)(QGroupBox_IsChecked(this.h))
}

func (this *QGroupBox) SetChecked(checked bool) {
	QGroupBox_SetChecked(this.h, (bool)(checked))
}

func (this *QGroupBox) Clicked() {
	QGroupBox_Clicked(this.h)
}
func (this *QGroupBox) OnClicked(slot func()) {
	QGroupBox_connect_Clicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Clicked
func miqt_exec_callback_QGroupBox_Clicked(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGroupBox) Toggled(param1 bool) {
	QGroupBox_Toggled(this.h, (bool)(param1))
}
func (this *QGroupBox) OnToggled(slot func(param1 bool)) {
	QGroupBox_connect_Toggled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Toggled
func miqt_exec_callback_QGroupBox_Toggled(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func QGroupBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGroupBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGroupBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGroupBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGroupBox) Clicked1(checked bool) {
	QGroupBox_Clicked1(this.h, (bool)(checked))
}
func (this *QGroupBox) OnClicked1(slot func(checked bool)) {
	QGroupBox_connect_Clicked1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Clicked1
func miqt_exec_callback_QGroupBox_Clicked1(cb intptr_t, checked bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checked bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checked)

	gofunc(slotval1)
}

func (this *QGroupBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QGroupBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGroupBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_MinimumSizeHint
func miqt_exec_callback_QGroupBox_MinimumSizeHint(self QGroupBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QGroupBox) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGroupBox_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGroupBox) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Event
func miqt_exec_callback_QGroupBox_Event(self QGroupBox, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGroupBox) callVirtualBase_ChildEvent(event *QChildEvent) {

	QGroupBox_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_ChildEvent
func miqt_exec_callback_QGroupBox_ChildEvent(self QGroupBox, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QGroupBox_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_ResizeEvent
func miqt_exec_callback_QGroupBox_ResizeEvent(self QGroupBox, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QGroupBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_PaintEvent
func miqt_exec_callback_QGroupBox_PaintEvent(self QGroupBox, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QGroupBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_FocusInEvent
func miqt_exec_callback_QGroupBox_FocusInEvent(self QGroupBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_ChangeEvent(event *QEvent) {

	QGroupBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_ChangeEvent
func miqt_exec_callback_QGroupBox_ChangeEvent(self QGroupBox, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QGroupBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_MousePressEvent
func miqt_exec_callback_QGroupBox_MousePressEvent(self QGroupBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QGroupBox_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_MouseMoveEvent
func miqt_exec_callback_QGroupBox_MouseMoveEvent(self QGroupBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QGroupBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_MouseReleaseEvent
func miqt_exec_callback_QGroupBox_MouseReleaseEvent(self QGroupBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_InitStyleOption(option *QStyleOptionGroupBox) {

	QGroupBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QGroupBox) OnInitStyleOption(slot func(super func(option *QStyleOptionGroupBox), option *QStyleOptionGroupBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_InitStyleOption
func miqt_exec_callback_QGroupBox_InitStyleOption(self QGroupBox, cb intptr_t, option *QStyleOptionGroupBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionGroupBox), option *QStyleOptionGroupBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionGroupBox(option)

	gofunc((&QGroupBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QGroupBox) callVirtualBase_DevType() int {

	return (int)(QGroupBox_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QGroupBox) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_DevType
func miqt_exec_callback_QGroupBox_DevType(self QGroupBox, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QGroupBox) callVirtualBase_SetVisible(visible bool) {

	QGroupBox_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QGroupBox) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_SetVisible
func miqt_exec_callback_QGroupBox_SetVisible(self QGroupBox, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QGroupBox{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QGroupBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QGroupBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGroupBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_SizeHint
func miqt_exec_callback_QGroupBox_SizeHint(self QGroupBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QGroupBox) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QGroupBox_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QGroupBox) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_HeightForWidth
func miqt_exec_callback_QGroupBox_HeightForWidth(self QGroupBox, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QGroupBox) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QGroupBox_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QGroupBox) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_HasHeightForWidth
func miqt_exec_callback_QGroupBox_HasHeightForWidth(self QGroupBox, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QGroupBox) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QGroupBox_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QGroupBox) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_PaintEngine
func miqt_exec_callback_QGroupBox_PaintEngine(self QGroupBox, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QGroupBox) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QGroupBox_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_MouseDoubleClickEvent
func miqt_exec_callback_QGroupBox_MouseDoubleClickEvent(self QGroupBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QGroupBox_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_WheelEvent
func miqt_exec_callback_QGroupBox_WheelEvent(self QGroupBox, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QGroupBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_KeyPressEvent
func miqt_exec_callback_QGroupBox_KeyPressEvent(self QGroupBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QGroupBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_KeyReleaseEvent
func miqt_exec_callback_QGroupBox_KeyReleaseEvent(self QGroupBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QGroupBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_FocusOutEvent
func miqt_exec_callback_QGroupBox_FocusOutEvent(self QGroupBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QGroupBox_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_EnterEvent
func miqt_exec_callback_QGroupBox_EnterEvent(self QGroupBox, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_LeaveEvent(event *QEvent) {

	QGroupBox_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_LeaveEvent
func miqt_exec_callback_QGroupBox_LeaveEvent(self QGroupBox, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QGroupBox_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_MoveEvent
func miqt_exec_callback_QGroupBox_MoveEvent(self QGroupBox, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QGroupBox_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_CloseEvent
func miqt_exec_callback_QGroupBox_CloseEvent(self QGroupBox, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QGroupBox_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_ContextMenuEvent
func miqt_exec_callback_QGroupBox_ContextMenuEvent(self QGroupBox, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QGroupBox_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_TabletEvent
func miqt_exec_callback_QGroupBox_TabletEvent(self QGroupBox, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_ActionEvent(event *QActionEvent) {

	QGroupBox_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_ActionEvent
func miqt_exec_callback_QGroupBox_ActionEvent(self QGroupBox, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QGroupBox_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_DragEnterEvent
func miqt_exec_callback_QGroupBox_DragEnterEvent(self QGroupBox, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QGroupBox_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_DragMoveEvent
func miqt_exec_callback_QGroupBox_DragMoveEvent(self QGroupBox, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QGroupBox_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_DragLeaveEvent
func miqt_exec_callback_QGroupBox_DragLeaveEvent(self QGroupBox, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_DropEvent(event *QDropEvent) {

	QGroupBox_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_DropEvent
func miqt_exec_callback_QGroupBox_DropEvent(self QGroupBox, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_ShowEvent(event *QShowEvent) {

	QGroupBox_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_ShowEvent
func miqt_exec_callback_QGroupBox_ShowEvent(self QGroupBox, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_HideEvent(event *QHideEvent) {

	QGroupBox_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGroupBox) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_HideEvent
func miqt_exec_callback_QGroupBox_HideEvent(self QGroupBox, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QGroupBox{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QGroupBox_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QGroupBox) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_NativeEvent
func miqt_exec_callback_QGroupBox_NativeEvent(self QGroupBox, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QGroupBox) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QGroupBox_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QGroupBox) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Metric
func miqt_exec_callback_QGroupBox_Metric(self QGroupBox, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QGroupBox) callVirtualBase_InitPainter(painter *QPainter) {

	QGroupBox_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QGroupBox) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_InitPainter
func miqt_exec_callback_QGroupBox_InitPainter(self QGroupBox, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QGroupBox{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QGroupBox) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QGroupBox_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QGroupBox) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Redirected
func miqt_exec_callback_QGroupBox_Redirected(self QGroupBox, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGroupBox) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QGroupBox_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QGroupBox) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_SharedPainter
func miqt_exec_callback_QGroupBox_SharedPainter(self QGroupBox, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QGroupBox) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QGroupBox_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QGroupBox) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_InputMethodEvent
func miqt_exec_callback_QGroupBox_InputMethodEvent(self QGroupBox, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QGroupBox{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QGroupBox) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QGroupBox_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGroupBox) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_InputMethodQuery
func miqt_exec_callback_QGroupBox_InputMethodQuery(self QGroupBox, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGroupBox) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QGroupBox_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QGroupBox) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_FocusNextPrevChild
func miqt_exec_callback_QGroupBox_FocusNextPrevChild(self QGroupBox, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
