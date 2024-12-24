package qt

import (
	"unsafe"
)

type QSpinBox struct {
	h          uintptr
	isSubclass bool
}

// NewQSpinBox constructs a new QSpinBox object.
func NewQSpinBox(parent *QWidget) *QSpinBox {

	ret := newQSpinBox(QSpinBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSpinBox2 constructs a new QSpinBox object.
func NewQSpinBox2() *QSpinBox {

	ret := newQSpinBox(QSpinBox_new2())
	ret.isSubclass = true
	return ret
}

func (this *QSpinBox) MetaObject() *QMetaObject {
	return newQMetaObject(QSpinBox_MetaObject(this.h))
}

func (this *QSpinBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSpinBox_Metacast(this.h, param1_Cstring))
}

func QSpinBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSpinBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) Value() int {
	return (int)(QSpinBox_Value(this.h))
}

func (this *QSpinBox) Prefix() string {
	var _ms struct_miqt_string = QSpinBox_Prefix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) SetPrefix(prefix string) {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	QSpinBox_SetPrefix(this.h, prefix_ms)
}

func (this *QSpinBox) Suffix() string {
	var _ms struct_miqt_string = QSpinBox_Suffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) SetSuffix(suffix string) {
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	QSpinBox_SetSuffix(this.h, suffix_ms)
}

func (this *QSpinBox) CleanText() string {
	var _ms struct_miqt_string = QSpinBox_CleanText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) SingleStep() int {
	return (int)(QSpinBox_SingleStep(this.h))
}

func (this *QSpinBox) SetSingleStep(val int) {
	QSpinBox_SetSingleStep(this.h, (int)(val))
}

func (this *QSpinBox) Minimum() int {
	return (int)(QSpinBox_Minimum(this.h))
}

func (this *QSpinBox) SetMinimum(min int) {
	QSpinBox_SetMinimum(this.h, (int)(min))
}

func (this *QSpinBox) Maximum() int {
	return (int)(QSpinBox_Maximum(this.h))
}

func (this *QSpinBox) SetMaximum(max int) {
	QSpinBox_SetMaximum(this.h, (int)(max))
}

func (this *QSpinBox) SetRange(min int, max int) {
	QSpinBox_SetRange(this.h, (int)(min), (int)(max))
}

func (this *QSpinBox) StepType() StepType {
	xxxxxxxxx
}

func (this *QSpinBox) SetStepType(stepType StepType) {
	QSpinBox_SetStepType(this.h, stepType)
}

func (this *QSpinBox) DisplayIntegerBase() int {
	return (int)(QSpinBox_DisplayIntegerBase(this.h))
}

func (this *QSpinBox) SetDisplayIntegerBase(base int) {
	QSpinBox_SetDisplayIntegerBase(this.h, (int)(base))
}

func (this *QSpinBox) SetValue(val int) {
	QSpinBox_SetValue(this.h, (int)(val))
}

func (this *QSpinBox) ValueChanged(param1 int) {
	QSpinBox_ValueChanged(this.h, (int)(param1))
}
func (this *QSpinBox) OnValueChanged(slot func(param1 int)) {
	QSpinBox_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_ValueChanged
func miqt_exec_callback_QSpinBox_ValueChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QSpinBox) TextChanged(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QSpinBox_TextChanged(this.h, param1_ms)
}
func (this *QSpinBox) OnTextChanged(slot func(param1 string)) {
	QSpinBox_connect_TextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_TextChanged
func miqt_exec_callback_QSpinBox_TextChanged(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func QSpinBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSpinBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSpinBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSpinBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QSpinBox_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QSpinBox) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_Event
func miqt_exec_callback_QSpinBox_Event(self QSpinBox, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSpinBox) callVirtualBase_Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	return (QValidator__State)(QSpinBox_virtualbase_Validate(unsafe.Pointer(this.h), input_ms, (*int)(unsafe.Pointer(pos))))

}
func (this *QSpinBox) OnValidate(slot func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_Validate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_Validate
func miqt_exec_callback_QSpinBox_Validate(self QSpinBox, cb intptr_t, input struct_miqt_string, pos *int) int {
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

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_Validate, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QSpinBox) callVirtualBase_ValueFromText(text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	return (int)(QSpinBox_virtualbase_ValueFromText(unsafe.Pointer(this.h), text_ms))

}
func (this *QSpinBox) OnValueFromText(slot func(super func(text string) int, text string) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_ValueFromText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_ValueFromText
func miqt_exec_callback_QSpinBox_ValueFromText(self QSpinBox, cb intptr_t, text struct_miqt_string) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(text string) int, text string) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_ValueFromText, slotval1)

	return (int)(virtualReturn)

}

func (this *QSpinBox) callVirtualBase_TextFromValue(val int) string {

	var _ms struct_miqt_string = QSpinBox_virtualbase_TextFromValue(unsafe.Pointer(this.h), (int)(val))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QSpinBox) OnTextFromValue(slot func(super func(val int) string, val int) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_TextFromValue(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_TextFromValue
func miqt_exec_callback_QSpinBox_TextFromValue(self QSpinBox, cb intptr_t, val int) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(val int) string, val int) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(val)

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_TextFromValue, slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QSpinBox) callVirtualBase_Fixup(str string) {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))

	QSpinBox_virtualbase_Fixup(unsafe.Pointer(this.h), str_ms)

}
func (this *QSpinBox) OnFixup(slot func(super func(str string), str string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_Fixup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_Fixup
func miqt_exec_callback_QSpinBox_Fixup(self QSpinBox, cb intptr_t, str struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(str string), str string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var str_ms struct_miqt_string = str
	str_ret := GoStringN(str_ms.data, int(int64(str_ms.len)))
	free(unsafe.Pointer(str_ms.data))
	slotval1 := str_ret

	gofunc((&QSpinBox{h: self}).callVirtualBase_Fixup, slotval1)

}

func (this *QSpinBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QSpinBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSpinBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_SizeHint
func miqt_exec_callback_QSpinBox_SizeHint(self QSpinBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QSpinBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QSpinBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSpinBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_MinimumSizeHint
func miqt_exec_callback_QSpinBox_MinimumSizeHint(self QSpinBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QSpinBox) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QSpinBox_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSpinBox) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_InputMethodQuery
func miqt_exec_callback_QSpinBox_InputMethodQuery(self QSpinBox, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QSpinBox) callVirtualBase_StepBy(steps int) {

	QSpinBox_virtualbase_StepBy(unsafe.Pointer(this.h), (int)(steps))

}
func (this *QSpinBox) OnStepBy(slot func(super func(steps int), steps int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_StepBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_StepBy
func miqt_exec_callback_QSpinBox_StepBy(self QSpinBox, cb intptr_t, steps int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(steps int), steps int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(steps)

	gofunc((&QSpinBox{h: self}).callVirtualBase_StepBy, slotval1)

}

func (this *QSpinBox) callVirtualBase_Clear() {

	QSpinBox_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QSpinBox) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_Clear
func miqt_exec_callback_QSpinBox_Clear(self QSpinBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QSpinBox{h: self}).callVirtualBase_Clear)

}

func (this *QSpinBox) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QSpinBox_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_ResizeEvent
func miqt_exec_callback_QSpinBox_ResizeEvent(self QSpinBox, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QSpinBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_KeyPressEvent
func miqt_exec_callback_QSpinBox_KeyPressEvent(self QSpinBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QSpinBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_KeyReleaseEvent
func miqt_exec_callback_QSpinBox_KeyReleaseEvent(self QSpinBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QSpinBox_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_WheelEvent
func miqt_exec_callback_QSpinBox_WheelEvent(self QSpinBox, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QSpinBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_FocusInEvent
func miqt_exec_callback_QSpinBox_FocusInEvent(self QSpinBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QSpinBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_FocusOutEvent
func miqt_exec_callback_QSpinBox_FocusOutEvent(self QSpinBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QSpinBox_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_ContextMenuEvent
func miqt_exec_callback_QSpinBox_ContextMenuEvent(self QSpinBox, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_ChangeEvent(event *QEvent) {

	QSpinBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_ChangeEvent
func miqt_exec_callback_QSpinBox_ChangeEvent(self QSpinBox, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QSpinBox_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_CloseEvent
func miqt_exec_callback_QSpinBox_CloseEvent(self QSpinBox, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_HideEvent(event *QHideEvent) {

	QSpinBox_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_HideEvent
func miqt_exec_callback_QSpinBox_HideEvent(self QSpinBox, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QSpinBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_MousePressEvent
func miqt_exec_callback_QSpinBox_MousePressEvent(self QSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QSpinBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_MouseReleaseEvent
func miqt_exec_callback_QSpinBox_MouseReleaseEvent(self QSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QSpinBox_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_MouseMoveEvent
func miqt_exec_callback_QSpinBox_MouseMoveEvent(self QSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QSpinBox_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_TimerEvent
func miqt_exec_callback_QSpinBox_TimerEvent(self QSpinBox, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QSpinBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_PaintEvent
func miqt_exec_callback_QSpinBox_PaintEvent(self QSpinBox, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_ShowEvent(event *QShowEvent) {

	QSpinBox_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSpinBox) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_ShowEvent
func miqt_exec_callback_QSpinBox_ShowEvent(self QSpinBox, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QSpinBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QSpinBox) callVirtualBase_InitStyleOption(option *QStyleOptionSpinBox) {

	QSpinBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QSpinBox) OnInitStyleOption(slot func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_InitStyleOption
func miqt_exec_callback_QSpinBox_InitStyleOption(self QSpinBox, cb intptr_t, option *QStyleOptionSpinBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSpinBox(option)

	gofunc((&QSpinBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QSpinBox) callVirtualBase_StepEnabled() StepEnabled {

	xxxxxxxxx
}
func (this *QSpinBox) OnStepEnabled(slot func(super func() StepEnabled) StepEnabled) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_StepEnabled(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_StepEnabled
func miqt_exec_callback_QSpinBox_StepEnabled(self QSpinBox, cb intptr_t) StepEnabled {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() StepEnabled) StepEnabled)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_StepEnabled)

	return virtualReturn

}

type QDoubleSpinBox struct {
	h          uintptr
	isSubclass bool
}

// NewQDoubleSpinBox constructs a new QDoubleSpinBox object.
func NewQDoubleSpinBox(parent *QWidget) *QDoubleSpinBox {

	ret := newQDoubleSpinBox(QDoubleSpinBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDoubleSpinBox2 constructs a new QDoubleSpinBox object.
func NewQDoubleSpinBox2() *QDoubleSpinBox {

	ret := newQDoubleSpinBox(QDoubleSpinBox_new2())
	ret.isSubclass = true
	return ret
}

func (this *QDoubleSpinBox) MetaObject() *QMetaObject {
	return newQMetaObject(QDoubleSpinBox_MetaObject(this.h))
}

func (this *QDoubleSpinBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDoubleSpinBox_Metacast(this.h, param1_Cstring))
}

func QDoubleSpinBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDoubleSpinBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) Value() float64 {
	return (float64)(QDoubleSpinBox_Value(this.h))
}

func (this *QDoubleSpinBox) Prefix() string {
	var _ms struct_miqt_string = QDoubleSpinBox_Prefix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) SetPrefix(prefix string) {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	QDoubleSpinBox_SetPrefix(this.h, prefix_ms)
}

func (this *QDoubleSpinBox) Suffix() string {
	var _ms struct_miqt_string = QDoubleSpinBox_Suffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) SetSuffix(suffix string) {
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	QDoubleSpinBox_SetSuffix(this.h, suffix_ms)
}

func (this *QDoubleSpinBox) CleanText() string {
	var _ms struct_miqt_string = QDoubleSpinBox_CleanText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) SingleStep() float64 {
	return (float64)(QDoubleSpinBox_SingleStep(this.h))
}

func (this *QDoubleSpinBox) SetSingleStep(val float64) {
	QDoubleSpinBox_SetSingleStep(this.h, (double)(val))
}

func (this *QDoubleSpinBox) Minimum() float64 {
	return (float64)(QDoubleSpinBox_Minimum(this.h))
}

func (this *QDoubleSpinBox) SetMinimum(min float64) {
	QDoubleSpinBox_SetMinimum(this.h, (double)(min))
}

func (this *QDoubleSpinBox) Maximum() float64 {
	return (float64)(QDoubleSpinBox_Maximum(this.h))
}

func (this *QDoubleSpinBox) SetMaximum(max float64) {
	QDoubleSpinBox_SetMaximum(this.h, (double)(max))
}

func (this *QDoubleSpinBox) SetRange(min float64, max float64) {
	QDoubleSpinBox_SetRange(this.h, (double)(min), (double)(max))
}

func (this *QDoubleSpinBox) StepType() StepType {
	xxxxxxxxx
}

func (this *QDoubleSpinBox) SetStepType(stepType StepType) {
	QDoubleSpinBox_SetStepType(this.h, stepType)
}

func (this *QDoubleSpinBox) Decimals() int {
	return (int)(QDoubleSpinBox_Decimals(this.h))
}

func (this *QDoubleSpinBox) SetDecimals(prec int) {
	QDoubleSpinBox_SetDecimals(this.h, (int)(prec))
}

func (this *QDoubleSpinBox) Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))
	return (QValidator__State)(QDoubleSpinBox_Validate(this.h, input_ms, (*int)(unsafe.Pointer(pos))))
}

func (this *QDoubleSpinBox) ValueFromText(text string) float64 {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (float64)(QDoubleSpinBox_ValueFromText(this.h, text_ms))
}

func (this *QDoubleSpinBox) TextFromValue(val float64) string {
	var _ms struct_miqt_string = QDoubleSpinBox_TextFromValue(this.h, (double)(val))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) Fixup(str string) {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	QDoubleSpinBox_Fixup(this.h, str_ms)
}

func (this *QDoubleSpinBox) SetValue(val float64) {
	QDoubleSpinBox_SetValue(this.h, (double)(val))
}

func (this *QDoubleSpinBox) ValueChanged(param1 float64) {
	QDoubleSpinBox_ValueChanged(this.h, (double)(param1))
}
func (this *QDoubleSpinBox) OnValueChanged(slot func(param1 float64)) {
	QDoubleSpinBox_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_ValueChanged
func miqt_exec_callback_QDoubleSpinBox_ValueChanged(cb intptr_t, param1 double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(param1)

	gofunc(slotval1)
}

func (this *QDoubleSpinBox) TextChanged(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QDoubleSpinBox_TextChanged(this.h, param1_ms)
}
func (this *QDoubleSpinBox) OnTextChanged(slot func(param1 string)) {
	QDoubleSpinBox_connect_TextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_TextChanged
func miqt_exec_callback_QDoubleSpinBox_TextChanged(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func QDoubleSpinBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDoubleSpinBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDoubleSpinBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDoubleSpinBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) callVirtualBase_Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	return (QValidator__State)(QDoubleSpinBox_virtualbase_Validate(unsafe.Pointer(this.h), input_ms, (*int)(unsafe.Pointer(pos))))

}
func (this *QDoubleSpinBox) OnValidate(slot func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_Validate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_Validate
func miqt_exec_callback_QDoubleSpinBox_Validate(self QDoubleSpinBox, cb intptr_t, input struct_miqt_string, pos *int) int {
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

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_Validate, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QDoubleSpinBox) callVirtualBase_ValueFromText(text string) float64 {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	return (float64)(QDoubleSpinBox_virtualbase_ValueFromText(unsafe.Pointer(this.h), text_ms))

}
func (this *QDoubleSpinBox) OnValueFromText(slot func(super func(text string) float64, text string) float64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_ValueFromText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_ValueFromText
func miqt_exec_callback_QDoubleSpinBox_ValueFromText(self QDoubleSpinBox, cb intptr_t, text struct_miqt_string) double {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(text string) float64, text string) float64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_ValueFromText, slotval1)

	return (double)(virtualReturn)

}

func (this *QDoubleSpinBox) callVirtualBase_TextFromValue(val float64) string {

	var _ms struct_miqt_string = QDoubleSpinBox_virtualbase_TextFromValue(unsafe.Pointer(this.h), (double)(val))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QDoubleSpinBox) OnTextFromValue(slot func(super func(val float64) string, val float64) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_TextFromValue(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_TextFromValue
func miqt_exec_callback_QDoubleSpinBox_TextFromValue(self QDoubleSpinBox, cb intptr_t, val double) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(val float64) string, val float64) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(val)

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_TextFromValue, slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QDoubleSpinBox) callVirtualBase_Fixup(str string) {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))

	QDoubleSpinBox_virtualbase_Fixup(unsafe.Pointer(this.h), str_ms)

}
func (this *QDoubleSpinBox) OnFixup(slot func(super func(str string), str string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_Fixup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_Fixup
func miqt_exec_callback_QDoubleSpinBox_Fixup(self QDoubleSpinBox, cb intptr_t, str struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(str string), str string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var str_ms struct_miqt_string = str
	str_ret := GoStringN(str_ms.data, int(int64(str_ms.len)))
	free(unsafe.Pointer(str_ms.data))
	slotval1 := str_ret

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_Fixup, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QDoubleSpinBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDoubleSpinBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_SizeHint
func miqt_exec_callback_QDoubleSpinBox_SizeHint(self QDoubleSpinBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QDoubleSpinBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QDoubleSpinBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDoubleSpinBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_MinimumSizeHint
func miqt_exec_callback_QDoubleSpinBox_MinimumSizeHint(self QDoubleSpinBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QDoubleSpinBox) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QDoubleSpinBox_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QDoubleSpinBox) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_Event
func miqt_exec_callback_QDoubleSpinBox_Event(self QDoubleSpinBox, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDoubleSpinBox) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QDoubleSpinBox_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDoubleSpinBox) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_InputMethodQuery
func miqt_exec_callback_QDoubleSpinBox_InputMethodQuery(self QDoubleSpinBox, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDoubleSpinBox) callVirtualBase_StepBy(steps int) {

	QDoubleSpinBox_virtualbase_StepBy(unsafe.Pointer(this.h), (int)(steps))

}
func (this *QDoubleSpinBox) OnStepBy(slot func(super func(steps int), steps int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_StepBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_StepBy
func miqt_exec_callback_QDoubleSpinBox_StepBy(self QDoubleSpinBox, cb intptr_t, steps int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(steps int), steps int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(steps)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_StepBy, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_Clear() {

	QDoubleSpinBox_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QDoubleSpinBox) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_Clear
func miqt_exec_callback_QDoubleSpinBox_Clear(self QDoubleSpinBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_Clear)

}

func (this *QDoubleSpinBox) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QDoubleSpinBox_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_ResizeEvent
func miqt_exec_callback_QDoubleSpinBox_ResizeEvent(self QDoubleSpinBox, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QDoubleSpinBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_KeyPressEvent
func miqt_exec_callback_QDoubleSpinBox_KeyPressEvent(self QDoubleSpinBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QDoubleSpinBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_KeyReleaseEvent
func miqt_exec_callback_QDoubleSpinBox_KeyReleaseEvent(self QDoubleSpinBox, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QDoubleSpinBox_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_WheelEvent
func miqt_exec_callback_QDoubleSpinBox_WheelEvent(self QDoubleSpinBox, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QDoubleSpinBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_FocusInEvent
func miqt_exec_callback_QDoubleSpinBox_FocusInEvent(self QDoubleSpinBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QDoubleSpinBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_FocusOutEvent
func miqt_exec_callback_QDoubleSpinBox_FocusOutEvent(self QDoubleSpinBox, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QDoubleSpinBox_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_ContextMenuEvent
func miqt_exec_callback_QDoubleSpinBox_ContextMenuEvent(self QDoubleSpinBox, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_ChangeEvent(event *QEvent) {

	QDoubleSpinBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_ChangeEvent
func miqt_exec_callback_QDoubleSpinBox_ChangeEvent(self QDoubleSpinBox, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QDoubleSpinBox_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_CloseEvent
func miqt_exec_callback_QDoubleSpinBox_CloseEvent(self QDoubleSpinBox, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_HideEvent(event *QHideEvent) {

	QDoubleSpinBox_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_HideEvent
func miqt_exec_callback_QDoubleSpinBox_HideEvent(self QDoubleSpinBox, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QDoubleSpinBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_MousePressEvent
func miqt_exec_callback_QDoubleSpinBox_MousePressEvent(self QDoubleSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QDoubleSpinBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_MouseReleaseEvent
func miqt_exec_callback_QDoubleSpinBox_MouseReleaseEvent(self QDoubleSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QDoubleSpinBox_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_MouseMoveEvent
func miqt_exec_callback_QDoubleSpinBox_MouseMoveEvent(self QDoubleSpinBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QDoubleSpinBox_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_TimerEvent
func miqt_exec_callback_QDoubleSpinBox_TimerEvent(self QDoubleSpinBox, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QDoubleSpinBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_PaintEvent
func miqt_exec_callback_QDoubleSpinBox_PaintEvent(self QDoubleSpinBox, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_ShowEvent(event *QShowEvent) {

	QDoubleSpinBox_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDoubleSpinBox) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_ShowEvent
func miqt_exec_callback_QDoubleSpinBox_ShowEvent(self QDoubleSpinBox, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_InitStyleOption(option *QStyleOptionSpinBox) {

	QDoubleSpinBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QDoubleSpinBox) OnInitStyleOption(slot func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_InitStyleOption
func miqt_exec_callback_QDoubleSpinBox_InitStyleOption(self QDoubleSpinBox, cb intptr_t, option *QStyleOptionSpinBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSpinBox(option)

	gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QDoubleSpinBox) callVirtualBase_StepEnabled() StepEnabled {

	xxxxxxxxx
}
func (this *QDoubleSpinBox) OnStepEnabled(slot func(super func() StepEnabled) StepEnabled) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_StepEnabled(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_StepEnabled
func miqt_exec_callback_QDoubleSpinBox_StepEnabled(self QDoubleSpinBox, cb intptr_t) StepEnabled {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() StepEnabled) StepEnabled)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_StepEnabled)

	return virtualReturn

}
