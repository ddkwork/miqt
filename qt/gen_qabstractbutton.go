package qt

import (
	"unsafe"
)

type QAbstractButton struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractButton constructs a new QAbstractButton object.
func NewQAbstractButton(parent *QWidget) *QAbstractButton {

	ret := newQAbstractButton(QAbstractButton_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAbstractButton2 constructs a new QAbstractButton object.
func NewQAbstractButton2() *QAbstractButton {

	ret := newQAbstractButton(QAbstractButton_new2())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractButton) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractButton_MetaObject(this.h))
}

func (this *QAbstractButton) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractButton_Metacast(this.h, param1_Cstring))
}

func QAbstractButton_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractButton_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractButton) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QAbstractButton_SetText(this.h, text_ms)
}

func (this *QAbstractButton) Text() string {
	var _ms struct_miqt_string = QAbstractButton_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractButton) SetIcon(icon *QIcon) {
	QAbstractButton_SetIcon(this.h, icon.cPointer())
}

func (this *QAbstractButton) Icon() *QIcon {
	_goptr := newQIcon(QAbstractButton_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractButton) IconSize() *QSize {
	_goptr := newQSize(QAbstractButton_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractButton) SetShortcut(key *QKeySequence) {
	QAbstractButton_SetShortcut(this.h, key.cPointer())
}

func (this *QAbstractButton) Shortcut() *QKeySequence {
	_goptr := newQKeySequence(QAbstractButton_Shortcut(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractButton) SetCheckable(checkable bool) {
	QAbstractButton_SetCheckable(this.h, (bool)(checkable))
}

func (this *QAbstractButton) IsCheckable() bool {
	return (bool)(QAbstractButton_IsCheckable(this.h))
}

func (this *QAbstractButton) IsChecked() bool {
	return (bool)(QAbstractButton_IsChecked(this.h))
}

func (this *QAbstractButton) SetDown(down bool) {
	QAbstractButton_SetDown(this.h, (bool)(down))
}

func (this *QAbstractButton) IsDown() bool {
	return (bool)(QAbstractButton_IsDown(this.h))
}

func (this *QAbstractButton) SetAutoRepeat(autoRepeat bool) {
	QAbstractButton_SetAutoRepeat(this.h, (bool)(autoRepeat))
}

func (this *QAbstractButton) AutoRepeat() bool {
	return (bool)(QAbstractButton_AutoRepeat(this.h))
}

func (this *QAbstractButton) SetAutoRepeatDelay(autoRepeatDelay int) {
	QAbstractButton_SetAutoRepeatDelay(this.h, (int)(autoRepeatDelay))
}

func (this *QAbstractButton) AutoRepeatDelay() int {
	return (int)(QAbstractButton_AutoRepeatDelay(this.h))
}

func (this *QAbstractButton) SetAutoRepeatInterval(autoRepeatInterval int) {
	QAbstractButton_SetAutoRepeatInterval(this.h, (int)(autoRepeatInterval))
}

func (this *QAbstractButton) AutoRepeatInterval() int {
	return (int)(QAbstractButton_AutoRepeatInterval(this.h))
}

func (this *QAbstractButton) SetAutoExclusive(autoExclusive bool) {
	QAbstractButton_SetAutoExclusive(this.h, (bool)(autoExclusive))
}

func (this *QAbstractButton) AutoExclusive() bool {
	return (bool)(QAbstractButton_AutoExclusive(this.h))
}

func (this *QAbstractButton) Group() *QButtonGroup {
	return newQButtonGroup(QAbstractButton_Group(this.h))
}

func (this *QAbstractButton) SetIconSize(size *QSize) {
	QAbstractButton_SetIconSize(this.h, size.cPointer())
}

func (this *QAbstractButton) AnimateClick() {
	QAbstractButton_AnimateClick(this.h)
}

func (this *QAbstractButton) Click() {
	QAbstractButton_Click(this.h)
}

func (this *QAbstractButton) Toggle() {
	QAbstractButton_Toggle(this.h)
}

func (this *QAbstractButton) SetChecked(checked bool) {
	QAbstractButton_SetChecked(this.h, (bool)(checked))
}

func (this *QAbstractButton) Pressed() {
	QAbstractButton_Pressed(this.h)
}
func (this *QAbstractButton) OnPressed(slot func()) {
	QAbstractButton_connect_Pressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Pressed
func miqt_exec_callback_QAbstractButton_Pressed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractButton) Released() {
	QAbstractButton_Released(this.h)
}
func (this *QAbstractButton) OnReleased(slot func()) {
	QAbstractButton_connect_Released(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Released
func miqt_exec_callback_QAbstractButton_Released(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractButton) Clicked() {
	QAbstractButton_Clicked(this.h)
}
func (this *QAbstractButton) OnClicked(slot func()) {
	QAbstractButton_connect_Clicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Clicked
func miqt_exec_callback_QAbstractButton_Clicked(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractButton) Toggled(checked bool) {
	QAbstractButton_Toggled(this.h, (bool)(checked))
}
func (this *QAbstractButton) OnToggled(slot func(checked bool)) {
	QAbstractButton_connect_Toggled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Toggled
func miqt_exec_callback_QAbstractButton_Toggled(cb intptr_t, checked bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checked bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checked)

	gofunc(slotval1)
}

func QAbstractButton_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractButton_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractButton_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractButton_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractButton) Clicked1(checked bool) {
	QAbstractButton_Clicked1(this.h, (bool)(checked))
}
func (this *QAbstractButton) OnClicked1(slot func(checked bool)) {
	QAbstractButton_connect_Clicked1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Clicked1
func miqt_exec_callback_QAbstractButton_Clicked1(cb intptr_t, checked bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checked bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checked)

	gofunc(slotval1)
}

func (this *QAbstractButton) OnPaintEvent(slot func(e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_PaintEvent
func miqt_exec_callback_QAbstractButton_PaintEvent(self QAbstractButton, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc(slotval1)

}

func (this *QAbstractButton) callVirtualBase_HitButton(pos *QPoint) bool {

	return (bool)(QAbstractButton_virtualbase_HitButton(unsafe.Pointer(this.h), pos.cPointer()))

}
func (this *QAbstractButton) OnHitButton(slot func(super func(pos *QPoint) bool, pos *QPoint) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_HitButton(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_HitButton
func miqt_exec_callback_QAbstractButton_HitButton(self QAbstractButton, cb intptr_t, pos *QPoint) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos *QPoint) bool, pos *QPoint) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(pos)

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_HitButton, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractButton) callVirtualBase_CheckStateSet() {

	QAbstractButton_virtualbase_CheckStateSet(unsafe.Pointer(this.h))

}
func (this *QAbstractButton) OnCheckStateSet(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_CheckStateSet(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_CheckStateSet
func miqt_exec_callback_QAbstractButton_CheckStateSet(self QAbstractButton, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractButton{h: self}).callVirtualBase_CheckStateSet)

}

func (this *QAbstractButton) callVirtualBase_NextCheckState() {

	QAbstractButton_virtualbase_NextCheckState(unsafe.Pointer(this.h))

}
func (this *QAbstractButton) OnNextCheckState(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_NextCheckState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_NextCheckState
func miqt_exec_callback_QAbstractButton_NextCheckState(self QAbstractButton, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractButton{h: self}).callVirtualBase_NextCheckState)

}

func (this *QAbstractButton) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QAbstractButton_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QAbstractButton) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Event
func miqt_exec_callback_QAbstractButton_Event(self QAbstractButton, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractButton) callVirtualBase_KeyPressEvent(e *QKeyEvent) {

	QAbstractButton_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnKeyPressEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_KeyPressEvent
func miqt_exec_callback_QAbstractButton_KeyPressEvent(self QAbstractButton, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_KeyReleaseEvent(e *QKeyEvent) {

	QAbstractButton_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnKeyReleaseEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_KeyReleaseEvent
func miqt_exec_callback_QAbstractButton_KeyReleaseEvent(self QAbstractButton, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_MousePressEvent(e *QMouseEvent) {

	QAbstractButton_virtualbase_MousePressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnMousePressEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_MousePressEvent
func miqt_exec_callback_QAbstractButton_MousePressEvent(self QAbstractButton, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QAbstractButton_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_MouseReleaseEvent
func miqt_exec_callback_QAbstractButton_MouseReleaseEvent(self QAbstractButton, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_MouseMoveEvent(e *QMouseEvent) {

	QAbstractButton_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnMouseMoveEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_MouseMoveEvent
func miqt_exec_callback_QAbstractButton_MouseMoveEvent(self QAbstractButton, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_FocusInEvent(e *QFocusEvent) {

	QAbstractButton_virtualbase_FocusInEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnFocusInEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_FocusInEvent
func miqt_exec_callback_QAbstractButton_FocusInEvent(self QAbstractButton, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_FocusOutEvent(e *QFocusEvent) {

	QAbstractButton_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnFocusOutEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_FocusOutEvent
func miqt_exec_callback_QAbstractButton_FocusOutEvent(self QAbstractButton, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_ChangeEvent(e *QEvent) {

	QAbstractButton_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_ChangeEvent
func miqt_exec_callback_QAbstractButton_ChangeEvent(self QAbstractButton, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_TimerEvent(e *QTimerEvent) {

	QAbstractButton_virtualbase_TimerEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QAbstractButton) OnTimerEvent(slot func(super func(e *QTimerEvent), e *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_TimerEvent
func miqt_exec_callback_QAbstractButton_TimerEvent(self QAbstractButton, cb intptr_t, e *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QTimerEvent), e *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(e)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_DevType() int {

	return (int)(QAbstractButton_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QAbstractButton) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_DevType
func miqt_exec_callback_QAbstractButton_DevType(self QAbstractButton, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QAbstractButton) callVirtualBase_SetVisible(visible bool) {

	QAbstractButton_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QAbstractButton) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_SetVisible
func miqt_exec_callback_QAbstractButton_SetVisible(self QAbstractButton, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QAbstractButton) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QAbstractButton_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractButton) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_SizeHint
func miqt_exec_callback_QAbstractButton_SizeHint(self QAbstractButton, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractButton) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QAbstractButton_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractButton) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_MinimumSizeHint
func miqt_exec_callback_QAbstractButton_MinimumSizeHint(self QAbstractButton, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractButton) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QAbstractButton_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QAbstractButton) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_HeightForWidth
func miqt_exec_callback_QAbstractButton_HeightForWidth(self QAbstractButton, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractButton) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QAbstractButton_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QAbstractButton) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_HasHeightForWidth
func miqt_exec_callback_QAbstractButton_HasHeightForWidth(self QAbstractButton, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QAbstractButton) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QAbstractButton_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QAbstractButton) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_PaintEngine
func miqt_exec_callback_QAbstractButton_PaintEngine(self QAbstractButton, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QAbstractButton) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QAbstractButton_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_MouseDoubleClickEvent
func miqt_exec_callback_QAbstractButton_MouseDoubleClickEvent(self QAbstractButton, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QAbstractButton_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_WheelEvent
func miqt_exec_callback_QAbstractButton_WheelEvent(self QAbstractButton, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QAbstractButton_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_EnterEvent
func miqt_exec_callback_QAbstractButton_EnterEvent(self QAbstractButton, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_LeaveEvent(event *QEvent) {

	QAbstractButton_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_LeaveEvent
func miqt_exec_callback_QAbstractButton_LeaveEvent(self QAbstractButton, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QAbstractButton_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_MoveEvent
func miqt_exec_callback_QAbstractButton_MoveEvent(self QAbstractButton, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QAbstractButton_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_ResizeEvent
func miqt_exec_callback_QAbstractButton_ResizeEvent(self QAbstractButton, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QAbstractButton_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_CloseEvent
func miqt_exec_callback_QAbstractButton_CloseEvent(self QAbstractButton, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QAbstractButton_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_ContextMenuEvent
func miqt_exec_callback_QAbstractButton_ContextMenuEvent(self QAbstractButton, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QAbstractButton_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_TabletEvent
func miqt_exec_callback_QAbstractButton_TabletEvent(self QAbstractButton, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_ActionEvent(event *QActionEvent) {

	QAbstractButton_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_ActionEvent
func miqt_exec_callback_QAbstractButton_ActionEvent(self QAbstractButton, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QAbstractButton_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_DragEnterEvent
func miqt_exec_callback_QAbstractButton_DragEnterEvent(self QAbstractButton, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QAbstractButton_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_DragMoveEvent
func miqt_exec_callback_QAbstractButton_DragMoveEvent(self QAbstractButton, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QAbstractButton_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_DragLeaveEvent
func miqt_exec_callback_QAbstractButton_DragLeaveEvent(self QAbstractButton, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_DropEvent(event *QDropEvent) {

	QAbstractButton_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_DropEvent
func miqt_exec_callback_QAbstractButton_DropEvent(self QAbstractButton, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_ShowEvent(event *QShowEvent) {

	QAbstractButton_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_ShowEvent
func miqt_exec_callback_QAbstractButton_ShowEvent(self QAbstractButton, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_HideEvent(event *QHideEvent) {

	QAbstractButton_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractButton) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_HideEvent
func miqt_exec_callback_QAbstractButton_HideEvent(self QAbstractButton, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QAbstractButton_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QAbstractButton) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_NativeEvent
func miqt_exec_callback_QAbstractButton_NativeEvent(self QAbstractButton, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractButton) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QAbstractButton_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QAbstractButton) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Metric
func miqt_exec_callback_QAbstractButton_Metric(self QAbstractButton, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractButton) callVirtualBase_InitPainter(painter *QPainter) {

	QAbstractButton_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QAbstractButton) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_InitPainter
func miqt_exec_callback_QAbstractButton_InitPainter(self QAbstractButton, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QAbstractButton) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QAbstractButton_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QAbstractButton) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Redirected
func miqt_exec_callback_QAbstractButton_Redirected(self QAbstractButton, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractButton) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QAbstractButton_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QAbstractButton) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_SharedPainter
func miqt_exec_callback_QAbstractButton_SharedPainter(self QAbstractButton, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QAbstractButton) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QAbstractButton_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractButton) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_InputMethodEvent
func miqt_exec_callback_QAbstractButton_InputMethodEvent(self QAbstractButton, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QAbstractButton{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QAbstractButton) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QAbstractButton_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractButton) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_InputMethodQuery
func miqt_exec_callback_QAbstractButton_InputMethodQuery(self QAbstractButton, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractButton) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QAbstractButton_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QAbstractButton) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_FocusNextPrevChild
func miqt_exec_callback_QAbstractButton_FocusNextPrevChild(self QAbstractButton, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
