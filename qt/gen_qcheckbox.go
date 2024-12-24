package qt

import (
	"unsafe"
)

type QCheckBox struct {
	h          uintptr
	isSubclass bool
}

// NewQCheckBox constructs a new QCheckBox object.
func NewQCheckBox(parent *QWidget) *QCheckBox {

	ret := newQCheckBox(QCheckBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCheckBox2 constructs a new QCheckBox object.
func NewQCheckBox2() *QCheckBox {

	ret := newQCheckBox(QCheckBox_new2())
	ret.isSubclass = true
	return ret
}

// NewQCheckBox3 constructs a new QCheckBox object.
func NewQCheckBox3(text string) *QCheckBox {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQCheckBox(QCheckBox_new3(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQCheckBox4 constructs a new QCheckBox object.
func NewQCheckBox4(text string, parent *QWidget) *QCheckBox {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQCheckBox(QCheckBox_new4(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCheckBox) MetaObject() *QMetaObject {
	return newQMetaObject(QCheckBox_MetaObject(this.h))
}

func (this *QCheckBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCheckBox_Metacast(this.h, param1_Cstring))
}

func QCheckBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCheckBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCheckBox) SizeHint() *QSize {
	_goptr := newQSize(QCheckBox_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCheckBox) MinimumSizeHint() *QSize {
	_goptr := newQSize(QCheckBox_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCheckBox) SetTristate() {
	QCheckBox_SetTristate(this.h)
}

func (this *QCheckBox) IsTristate() bool {
	return (bool)(QCheckBox_IsTristate(this.h))
}

func (this *QCheckBox) CheckState() CheckState {
	return (CheckState)(QCheckBox_CheckState(this.h))
}

func (this *QCheckBox) SetCheckState(state CheckState) {
	QCheckBox_SetCheckState(this.h, (int)(state))
}

func (this *QCheckBox) StateChanged(param1 int) {
	QCheckBox_StateChanged(this.h, (int)(param1))
}
func (this *QCheckBox) OnStateChanged(slot func(param1 int)) {
	QCheckBox_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_StateChanged
func miqt_exec_callback_QCheckBox_StateChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QCheckBox) CheckStateChanged(param1 CheckState) {
	QCheckBox_CheckStateChanged(this.h, (int)(param1))
}
func (this *QCheckBox) OnCheckStateChanged(slot func(param1 CheckState)) {
	QCheckBox_connect_CheckStateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_CheckStateChanged
func miqt_exec_callback_QCheckBox_CheckStateChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 CheckState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (CheckState)(param1)

	gofunc(slotval1)
}

func QCheckBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCheckBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCheckBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCheckBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCheckBox) SetTristate1(y bool) {
	QCheckBox_SetTristate1(this.h, (bool)(y))
}

func (this *QCheckBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QCheckBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCheckBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_SizeHint
func miqt_exec_callback_QCheckBox_SizeHint(self QCheckBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QCheckBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QCheckBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCheckBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MinimumSizeHint
func miqt_exec_callback_QCheckBox_MinimumSizeHint(self QCheckBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QCheckBox) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QCheckBox_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QCheckBox) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_Event
func miqt_exec_callback_QCheckBox_Event(self QCheckBox, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QCheckBox) callVirtualBase_HitButton(pos *QPoint) bool {

	return (bool)(QCheckBox_virtualbase_HitButton(unsafe.Pointer(this.h), pos.cPointer()))

}
func (this *QCheckBox) OnHitButton(slot func(super func(pos *QPoint) bool, pos *QPoint) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_HitButton(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_HitButton
func miqt_exec_callback_QCheckBox_HitButton(self QCheckBox, cb intptr_t, pos *QPoint) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos *QPoint) bool, pos *QPoint) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(pos)

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_HitButton, slotval1)

	return (bool)(virtualReturn)

}

func (this *QCheckBox) callVirtualBase_CheckStateSet() {

	QCheckBox_virtualbase_CheckStateSet(unsafe.Pointer(this.h))

}
func (this *QCheckBox) OnCheckStateSet(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_CheckStateSet(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_CheckStateSet
func miqt_exec_callback_QCheckBox_CheckStateSet(self QCheckBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QCheckBox{h: self}).callVirtualBase_CheckStateSet)

}

func (this *QCheckBox) callVirtualBase_NextCheckState() {

	QCheckBox_virtualbase_NextCheckState(unsafe.Pointer(this.h))

}
func (this *QCheckBox) OnNextCheckState(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_NextCheckState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_NextCheckState
func miqt_exec_callback_QCheckBox_NextCheckState(self QCheckBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QCheckBox{h: self}).callVirtualBase_NextCheckState)

}

func (this *QCheckBox) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QCheckBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCheckBox) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_PaintEvent
func miqt_exec_callback_QCheckBox_PaintEvent(self QCheckBox, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QCheckBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QCheckBox_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCheckBox) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MouseMoveEvent
func miqt_exec_callback_QCheckBox_MouseMoveEvent(self QCheckBox, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QCheckBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_InitStyleOption(option *QStyleOptionButton) {

	QCheckBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QCheckBox) OnInitStyleOption(slot func(super func(option *QStyleOptionButton), option *QStyleOptionButton)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_InitStyleOption
func miqt_exec_callback_QCheckBox_InitStyleOption(self QCheckBox, cb intptr_t, option *QStyleOptionButton) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionButton), option *QStyleOptionButton))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionButton(option)

	gofunc((&QCheckBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QCheckBox) callVirtualBase_KeyPressEvent(e *QKeyEvent) {

	QCheckBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnKeyPressEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_KeyPressEvent
func miqt_exec_callback_QCheckBox_KeyPressEvent(self QCheckBox, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_KeyReleaseEvent(e *QKeyEvent) {

	QCheckBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnKeyReleaseEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_KeyReleaseEvent
func miqt_exec_callback_QCheckBox_KeyReleaseEvent(self QCheckBox, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_MousePressEvent(e *QMouseEvent) {

	QCheckBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnMousePressEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MousePressEvent
func miqt_exec_callback_QCheckBox_MousePressEvent(self QCheckBox, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QCheckBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MouseReleaseEvent
func miqt_exec_callback_QCheckBox_MouseReleaseEvent(self QCheckBox, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_FocusInEvent(e *QFocusEvent) {

	QCheckBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnFocusInEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_FocusInEvent
func miqt_exec_callback_QCheckBox_FocusInEvent(self QCheckBox, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_FocusOutEvent(e *QFocusEvent) {

	QCheckBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnFocusOutEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_FocusOutEvent
func miqt_exec_callback_QCheckBox_FocusOutEvent(self QCheckBox, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_ChangeEvent(e *QEvent) {

	QCheckBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_ChangeEvent
func miqt_exec_callback_QCheckBox_ChangeEvent(self QCheckBox, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QCheckBox) callVirtualBase_TimerEvent(e *QTimerEvent) {

	QCheckBox_virtualbase_TimerEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QCheckBox) OnTimerEvent(slot func(super func(e *QTimerEvent), e *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_TimerEvent
func miqt_exec_callback_QCheckBox_TimerEvent(self QCheckBox, cb intptr_t, e *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QTimerEvent), e *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(e)

	gofunc((&QCheckBox{h: self}).callVirtualBase_TimerEvent, slotval1)

}
