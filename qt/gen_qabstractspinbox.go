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

func (this *QAbstractSpinBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractSpinBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractSpinBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_MetaObject
func miqt_exec_callback_QAbstractSpinBox_MetaObject(self QAbstractSpinBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractSpinBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractSpinBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractSpinBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractSpinBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractSpinBox_Metacast
func miqt_exec_callback_QAbstractSpinBox_Metacast(self QAbstractSpinBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractSpinBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
