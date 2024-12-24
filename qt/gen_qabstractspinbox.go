package qt

import (
	"unsafe"
)

type QAbstractSpinBox__StepEnabledFlag int

const (
	QAbstractSpinBox__StepNone        QAbstractSpinBox__StepEnabledFlag = 0
	QAbstractSpinBox__StepUpEnabled   QAbstractSpinBox__StepEnabledFlag = 1
	QAbstractSpinBox__StepDownEnabled QAbstractSpinBox__StepEnabledFlag = 2
)

type QAbstractSpinBox__ButtonSymbols int

const (
	QAbstractSpinBox__UpDownArrows QAbstractSpinBox__ButtonSymbols = 0
	QAbstractSpinBox__PlusMinus    QAbstractSpinBox__ButtonSymbols = 1
	QAbstractSpinBox__NoButtons    QAbstractSpinBox__ButtonSymbols = 2
)

type QAbstractSpinBox__CorrectionMode int

const (
	QAbstractSpinBox__CorrectToPreviousValue QAbstractSpinBox__CorrectionMode = 0
	QAbstractSpinBox__CorrectToNearestValue  QAbstractSpinBox__CorrectionMode = 1
)

type QAbstractSpinBox__StepType int

const (
	QAbstractSpinBox__DefaultStepType         QAbstractSpinBox__StepType = 0
	QAbstractSpinBox__AdaptiveDecimalStepType QAbstractSpinBox__StepType = 1
)

type QAbstractSpinBox struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractSpinBox constructs a new QAbstractSpinBox object.
func NewQAbstractSpinBox(parent *QWidget) *QAbstractSpinBox {

	ret := newQAbstractSpinBox(QAbstractSpinBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAbstractSpinBox2 constructs a new QAbstractSpinBox object.
func NewQAbstractSpinBox2() *QAbstractSpinBox {

	ret := newQAbstractSpinBox(QAbstractSpinBox_new2())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractSpinBox) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractSpinBox_MetaObject(this.h))
}

func (this *QAbstractSpinBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractSpinBox_Metacast(this.h, param1_Cstring))
}

func QAbstractSpinBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractSpinBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSpinBox) ButtonSymbols() ButtonSymbols {
	xxxxxxxxx
}

func (this *QAbstractSpinBox) SetButtonSymbols(bs ButtonSymbols) {
	QAbstractSpinBox_SetButtonSymbols(this.h, bs)
}

func (this *QAbstractSpinBox) SetCorrectionMode(cm CorrectionMode) {
	QAbstractSpinBox_SetCorrectionMode(this.h, cm)
}

func (this *QAbstractSpinBox) CorrectionMode() CorrectionMode {
	xxxxxxxxx
}

func (this *QAbstractSpinBox) HasAcceptableInput() bool {
	return (bool)(QAbstractSpinBox_HasAcceptableInput(this.h))
}

func (this *QAbstractSpinBox) Text() string {
	var _ms struct_miqt_string = QAbstractSpinBox_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSpinBox) SpecialValueText() string {
	var _ms struct_miqt_string = QAbstractSpinBox_SpecialValueText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSpinBox) SetSpecialValueText(txt string) {
	txt_ms := struct_miqt_string{}
	txt_ms.data = CString(txt)
	txt_ms.len = size_t(len(txt))
	defer free(unsafe.Pointer(txt_ms.data))
	QAbstractSpinBox_SetSpecialValueText(this.h, txt_ms)
}

func (this *QAbstractSpinBox) Wrapping() bool {
	return (bool)(QAbstractSpinBox_Wrapping(this.h))
}

func (this *QAbstractSpinBox) SetWrapping(w bool) {
	QAbstractSpinBox_SetWrapping(this.h, (bool)(w))
}

func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	QAbstractSpinBox_SetReadOnly(this.h, (bool)(r))
}

func (this *QAbstractSpinBox) IsReadOnly() bool {
	return (bool)(QAbstractSpinBox_IsReadOnly(this.h))
}

func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	QAbstractSpinBox_SetKeyboardTracking(this.h, (bool)(kt))
}

func (this *QAbstractSpinBox) KeyboardTracking() bool {
	return (bool)(QAbstractSpinBox_KeyboardTracking(this.h))
}

func (this *QAbstractSpinBox) SetAlignment(flag AlignmentFlag) {
	QAbstractSpinBox_SetAlignment(this.h, (int)(flag))
}

func (this *QAbstractSpinBox) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QAbstractSpinBox_Alignment(this.h))
}

func (this *QAbstractSpinBox) SetFrame(frame bool) {
	QAbstractSpinBox_SetFrame(this.h, (bool)(frame))
}

func (this *QAbstractSpinBox) HasFrame() bool {
	return (bool)(QAbstractSpinBox_HasFrame(this.h))
}

func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	QAbstractSpinBox_SetAccelerated(this.h, (bool)(on))
}

func (this *QAbstractSpinBox) IsAccelerated() bool {
	return (bool)(QAbstractSpinBox_IsAccelerated(this.h))
}

func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	QAbstractSpinBox_SetGroupSeparatorShown(this.h, (bool)(shown))
}

func (this *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	return (bool)(QAbstractSpinBox_IsGroupSeparatorShown(this.h))
}

func (this *QAbstractSpinBox) SizeHint() *QSize {
	_goptr := newQSize(QAbstractSpinBox_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractSpinBox) MinimumSizeHint() *QSize {
	_goptr := newQSize(QAbstractSpinBox_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractSpinBox) InterpretText() {
	QAbstractSpinBox_InterpretText(this.h)
}

func (this *QAbstractSpinBox) Event(event *QEvent) bool {
	return (bool)(QAbstractSpinBox_Event(this.h, event.cPointer()))
}

func (this *QAbstractSpinBox) InputMethodQuery(param1 InputMethodQuery) *QVariant {
	_goptr := newQVariant(QAbstractSpinBox_InputMethodQuery(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractSpinBox) Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))
	return (QValidator__State)(QAbstractSpinBox_Validate(this.h, input_ms, (*int)(unsafe.Pointer(pos))))
}

func (this *QAbstractSpinBox) Fixup(input string) {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))
	QAbstractSpinBox_Fixup(this.h, input_ms)
}

func (this *QAbstractSpinBox) StepBy(steps int) {
	QAbstractSpinBox_StepBy(this.h, (int)(steps))
}

func (this *QAbstractSpinBox) StepUp() {
	QAbstractSpinBox_StepUp(this.h)
}

func (this *QAbstractSpinBox) StepDown() {
	QAbstractSpinBox_StepDown(this.h)
}

func (this *QAbstractSpinBox) SelectAll() {
	QAbstractSpinBox_SelectAll(this.h)
}

func (this *QAbstractSpinBox) Clear() {
	QAbstractSpinBox_Clear(this.h)
}

func (this *QAbstractSpinBox) EditingFinished() {
	QAbstractSpinBox_EditingFinished(this.h)
}
func (this *QAbstractSpinBox) OnEditingFinished(slot func()) {
	QAbstractSpinBox_connect_EditingFinished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_EditingFinished
func miqt_exec_callback_QAbstractSpinBox_EditingFinished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QAbstractSpinBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractSpinBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractSpinBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractSpinBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractSpinBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QAbstractSpinBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractSpinBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_SizeHint
func miqt_exec_callback_QAbstractSpinBox_SizeHint(self QAbstractSpinBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractSpinBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QAbstractSpinBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractSpinBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_MinimumSizeHint
func miqt_exec_callback_QAbstractSpinBox_MinimumSizeHint(self QAbstractSpinBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractSpinBox) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QAbstractSpinBox_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QAbstractSpinBox) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_Event
func miqt_exec_callback_QAbstractSpinBox_Event(self QAbstractSpinBox, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractSpinBox) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QAbstractSpinBox_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractSpinBox) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_InputMethodQuery
func miqt_exec_callback_QAbstractSpinBox_InputMethodQuery(self QAbstractSpinBox, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractSpinBox) callVirtualBase_Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	return (QValidator__State)(QAbstractSpinBox_virtualbase_Validate(unsafe.Pointer(this.h), input_ms, (*int)(unsafe.Pointer(pos))))

}
func (this *QAbstractSpinBox) OnValidate(slot func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_Validate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_Validate
func miqt_exec_callback_QAbstractSpinBox_Validate(self QAbstractSpinBox, cb intptr_t, input struct_miqt_string, pos *int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret
	slotval2 := (*int)(unsafe.Pointer(pos))

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_Validate, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QAbstractSpinBox) callVirtualBase_Fixup(input string) {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	QAbstractSpinBox_virtualbase_Fixup(unsafe.Pointer(this.h), input_ms)

}
func (this *QAbstractSpinBox) OnFixup(slot func(super func(input string), input string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_Fixup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_Fixup
func miqt_exec_callback_QAbstractSpinBox_Fixup(self QAbstractSpinBox, cb intptr_t, input struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string), input string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_Fixup, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_StepBy(steps int) {

	QAbstractSpinBox_virtualbase_StepBy(unsafe.Pointer(this.h), (int)(steps))

}
func (this *QAbstractSpinBox) OnStepBy(slot func(super func(steps int), steps int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_StepBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_StepBy
func miqt_exec_callback_QAbstractSpinBox_StepBy(self QAbstractSpinBox, cb intptr_t, steps int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(steps int), steps int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(steps)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_StepBy, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_Clear() {

	QAbstractSpinBox_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QAbstractSpinBox) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_Clear
func miqt_exec_callback_QAbstractSpinBox_Clear(self QAbstractSpinBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_Clear)

}

func (this *QAbstractSpinBox) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QAbstractSpinBox_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_ResizeEvent
func miqt_exec_callback_QAbstractSpinBox_ResizeEvent(self QAbstractSpinBox, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QAbstractSpinBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_KeyPressEvent
func miqt_exec_callback_QAbstractSpinBox_KeyPressEvent(self QAbstractSpinBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QAbstractSpinBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_KeyReleaseEvent
func miqt_exec_callback_QAbstractSpinBox_KeyReleaseEvent(self QAbstractSpinBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QAbstractSpinBox_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_WheelEvent
func miqt_exec_callback_QAbstractSpinBox_WheelEvent(self QAbstractSpinBox, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QAbstractSpinBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_FocusInEvent
func miqt_exec_callback_QAbstractSpinBox_FocusInEvent(self QAbstractSpinBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QAbstractSpinBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_FocusOutEvent
func miqt_exec_callback_QAbstractSpinBox_FocusOutEvent(self QAbstractSpinBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QAbstractSpinBox_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_ContextMenuEvent
func miqt_exec_callback_QAbstractSpinBox_ContextMenuEvent(self QAbstractSpinBox, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_ChangeEvent(event *QEvent) {

	QAbstractSpinBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_ChangeEvent
func miqt_exec_callback_QAbstractSpinBox_ChangeEvent(self QAbstractSpinBox, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QAbstractSpinBox_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_CloseEvent
func miqt_exec_callback_QAbstractSpinBox_CloseEvent(self QAbstractSpinBox, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_HideEvent(event *QHideEvent) {

	QAbstractSpinBox_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_HideEvent
func miqt_exec_callback_QAbstractSpinBox_HideEvent(self QAbstractSpinBox, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QAbstractSpinBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_MousePressEvent
func miqt_exec_callback_QAbstractSpinBox_MousePressEvent(self QAbstractSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QAbstractSpinBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_MouseReleaseEvent
func miqt_exec_callback_QAbstractSpinBox_MouseReleaseEvent(self QAbstractSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QAbstractSpinBox_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_MouseMoveEvent
func miqt_exec_callback_QAbstractSpinBox_MouseMoveEvent(self QAbstractSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QAbstractSpinBox_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_TimerEvent
func miqt_exec_callback_QAbstractSpinBox_TimerEvent(self QAbstractSpinBox, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QAbstractSpinBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_PaintEvent
func miqt_exec_callback_QAbstractSpinBox_PaintEvent(self QAbstractSpinBox, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_ShowEvent(event *QShowEvent) {

	QAbstractSpinBox_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_ShowEvent
func miqt_exec_callback_QAbstractSpinBox_ShowEvent(self QAbstractSpinBox, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_InitStyleOption(option *QStyleOptionSpinBox) {

	QAbstractSpinBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QAbstractSpinBox) OnInitStyleOption(slot func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_InitStyleOption
func miqt_exec_callback_QAbstractSpinBox_InitStyleOption(self QAbstractSpinBox, cb intptr_t, option *QStyleOptionSpinBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSpinBox(option)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_StepEnabled() StepEnabled {

	xxxxxxxxx
}
func (this *QAbstractSpinBox) OnStepEnabled(slot func(super func() StepEnabled) StepEnabled) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_StepEnabled(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_StepEnabled
func miqt_exec_callback_QAbstractSpinBox_StepEnabled(self QAbstractSpinBox, cb intptr_t) StepEnabled {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() StepEnabled) StepEnabled)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_StepEnabled)

	return virtualReturn

}

func (this *QAbstractSpinBox) callVirtualBase_DevType() int {

	return (int)(QAbstractSpinBox_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QAbstractSpinBox) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_DevType
func miqt_exec_callback_QAbstractSpinBox_DevType(self QAbstractSpinBox, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QAbstractSpinBox) callVirtualBase_SetVisible(visible bool) {

	QAbstractSpinBox_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QAbstractSpinBox) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_SetVisible
func miqt_exec_callback_QAbstractSpinBox_SetVisible(self QAbstractSpinBox, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QAbstractSpinBox_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QAbstractSpinBox) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_HeightForWidth
func miqt_exec_callback_QAbstractSpinBox_HeightForWidth(self QAbstractSpinBox, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractSpinBox) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QAbstractSpinBox_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QAbstractSpinBox) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_HasHeightForWidth
func miqt_exec_callback_QAbstractSpinBox_HasHeightForWidth(self QAbstractSpinBox, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QAbstractSpinBox) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QAbstractSpinBox_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QAbstractSpinBox) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_PaintEngine
func miqt_exec_callback_QAbstractSpinBox_PaintEngine(self QAbstractSpinBox, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QAbstractSpinBox) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QAbstractSpinBox_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_MouseDoubleClickEvent
func miqt_exec_callback_QAbstractSpinBox_MouseDoubleClickEvent(self QAbstractSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QAbstractSpinBox_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_EnterEvent
func miqt_exec_callback_QAbstractSpinBox_EnterEvent(self QAbstractSpinBox, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_LeaveEvent(event *QEvent) {

	QAbstractSpinBox_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_LeaveEvent
func miqt_exec_callback_QAbstractSpinBox_LeaveEvent(self QAbstractSpinBox, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QAbstractSpinBox_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_MoveEvent
func miqt_exec_callback_QAbstractSpinBox_MoveEvent(self QAbstractSpinBox, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QAbstractSpinBox_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_TabletEvent
func miqt_exec_callback_QAbstractSpinBox_TabletEvent(self QAbstractSpinBox, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_ActionEvent(event *QActionEvent) {

	QAbstractSpinBox_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_ActionEvent
func miqt_exec_callback_QAbstractSpinBox_ActionEvent(self QAbstractSpinBox, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QAbstractSpinBox_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_DragEnterEvent
func miqt_exec_callback_QAbstractSpinBox_DragEnterEvent(self QAbstractSpinBox, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QAbstractSpinBox_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_DragMoveEvent
func miqt_exec_callback_QAbstractSpinBox_DragMoveEvent(self QAbstractSpinBox, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QAbstractSpinBox_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_DragLeaveEvent
func miqt_exec_callback_QAbstractSpinBox_DragLeaveEvent(self QAbstractSpinBox, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_DropEvent(event *QDropEvent) {

	QAbstractSpinBox_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractSpinBox) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_DropEvent
func miqt_exec_callback_QAbstractSpinBox_DropEvent(self QAbstractSpinBox, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QAbstractSpinBox_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QAbstractSpinBox) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_NativeEvent
func miqt_exec_callback_QAbstractSpinBox_NativeEvent(self QAbstractSpinBox, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractSpinBox) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QAbstractSpinBox_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QAbstractSpinBox) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_Metric
func miqt_exec_callback_QAbstractSpinBox_Metric(self QAbstractSpinBox, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractSpinBox) callVirtualBase_InitPainter(painter *QPainter) {

	QAbstractSpinBox_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QAbstractSpinBox) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_InitPainter
func miqt_exec_callback_QAbstractSpinBox_InitPainter(self QAbstractSpinBox, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QAbstractSpinBox_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QAbstractSpinBox) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_Redirected
func miqt_exec_callback_QAbstractSpinBox_Redirected(self QAbstractSpinBox, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractSpinBox) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QAbstractSpinBox_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QAbstractSpinBox) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_SharedPainter
func miqt_exec_callback_QAbstractSpinBox_SharedPainter(self QAbstractSpinBox, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QAbstractSpinBox) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QAbstractSpinBox_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractSpinBox) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_InputMethodEvent
func miqt_exec_callback_QAbstractSpinBox_InputMethodEvent(self QAbstractSpinBox, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QAbstractSpinBox) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QAbstractSpinBox_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QAbstractSpinBox) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_FocusNextPrevChild
func miqt_exec_callback_QAbstractSpinBox_FocusNextPrevChild(self QAbstractSpinBox, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
