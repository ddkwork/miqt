package qt

import (
	"unsafe"
)

type QLineEdit__ActionPosition int

const (
	QLineEdit__LeadingPosition  QLineEdit__ActionPosition = 0
	QLineEdit__TrailingPosition QLineEdit__ActionPosition = 1
)

type QLineEdit__EchoMode int

const (
	QLineEdit__Normal             QLineEdit__EchoMode = 0
	QLineEdit__NoEcho             QLineEdit__EchoMode = 1
	QLineEdit__Password           QLineEdit__EchoMode = 2
	QLineEdit__PasswordEchoOnEdit QLineEdit__EchoMode = 3
)

type QLineEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQLineEdit constructs a new QLineEdit object.
func NewQLineEdit(parent *QWidget) *QLineEdit {

	ret := newQLineEdit(QLineEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLineEdit2 constructs a new QLineEdit object.
func NewQLineEdit2() *QLineEdit {

	ret := newQLineEdit(QLineEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQLineEdit3 constructs a new QLineEdit object.
func NewQLineEdit3(param1 string) *QLineEdit {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))

	ret := newQLineEdit(QLineEdit_new3(param1_ms))
	ret.isSubclass = true
	return ret
}

// NewQLineEdit4 constructs a new QLineEdit object.
func NewQLineEdit4(param1 string, parent *QWidget) *QLineEdit {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))

	ret := newQLineEdit(QLineEdit_new4(param1_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QLineEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QLineEdit_MetaObject(this.h))
}

func (this *QLineEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLineEdit_Metacast(this.h, param1_Cstring))
}

func QLineEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLineEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLineEdit) Text() string {
	var _ms struct_miqt_string = QLineEdit_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLineEdit) DisplayText() string {
	var _ms struct_miqt_string = QLineEdit_DisplayText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLineEdit) PlaceholderText() string {
	var _ms struct_miqt_string = QLineEdit_PlaceholderText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLineEdit) SetPlaceholderText(placeholderText string) {
	placeholderText_ms := struct_miqt_string{}
	placeholderText_ms.data = CString(placeholderText)
	placeholderText_ms.len = size_t(len(placeholderText))
	defer free(unsafe.Pointer(placeholderText_ms.data))
	QLineEdit_SetPlaceholderText(this.h, placeholderText_ms)
}

func (this *QLineEdit) MaxLength() int {
	return (int)(QLineEdit_MaxLength(this.h))
}

func (this *QLineEdit) SetMaxLength(maxLength int) {
	QLineEdit_SetMaxLength(this.h, (int)(maxLength))
}

func (this *QLineEdit) SetFrame(frame bool) {
	QLineEdit_SetFrame(this.h, (bool)(frame))
}

func (this *QLineEdit) HasFrame() bool {
	return (bool)(QLineEdit_HasFrame(this.h))
}

func (this *QLineEdit) SetClearButtonEnabled(enable bool) {
	QLineEdit_SetClearButtonEnabled(this.h, (bool)(enable))
}

func (this *QLineEdit) IsClearButtonEnabled() bool {
	return (bool)(QLineEdit_IsClearButtonEnabled(this.h))
}

func (this *QLineEdit) EchoMode() EchoMode {
	xxxxxxxxx
}

func (this *QLineEdit) SetEchoMode(echoMode EchoMode) {
	QLineEdit_SetEchoMode(this.h, echoMode)
}

func (this *QLineEdit) IsReadOnly() bool {
	return (bool)(QLineEdit_IsReadOnly(this.h))
}

func (this *QLineEdit) SetReadOnly(readOnly bool) {
	QLineEdit_SetReadOnly(this.h, (bool)(readOnly))
}

func (this *QLineEdit) SetValidator(validator *QValidator) {
	QLineEdit_SetValidator(this.h, validator.cPointer())
}

func (this *QLineEdit) Validator() *QValidator {
	return newQValidator(QLineEdit_Validator(this.h))
}

func (this *QLineEdit) SetCompleter(completer *QCompleter) {
	QLineEdit_SetCompleter(this.h, completer.cPointer())
}

func (this *QLineEdit) Completer() *QCompleter {
	return newQCompleter(QLineEdit_Completer(this.h))
}

func (this *QLineEdit) SizeHint() *QSize {
	_goptr := newQSize(QLineEdit_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineEdit) MinimumSizeHint() *QSize {
	_goptr := newQSize(QLineEdit_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineEdit) CursorPosition() int {
	return (int)(QLineEdit_CursorPosition(this.h))
}

func (this *QLineEdit) SetCursorPosition(cursorPosition int) {
	QLineEdit_SetCursorPosition(this.h, (int)(cursorPosition))
}

func (this *QLineEdit) CursorPositionAt(pos *QPoint) int {
	return (int)(QLineEdit_CursorPositionAt(this.h, pos.cPointer()))
}

func (this *QLineEdit) SetAlignment(flag AlignmentFlag) {
	QLineEdit_SetAlignment(this.h, (int)(flag))
}

func (this *QLineEdit) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QLineEdit_Alignment(this.h))
}

func (this *QLineEdit) CursorForward(mark bool) {
	QLineEdit_CursorForward(this.h, (bool)(mark))
}

func (this *QLineEdit) CursorBackward(mark bool) {
	QLineEdit_CursorBackward(this.h, (bool)(mark))
}

func (this *QLineEdit) CursorWordForward(mark bool) {
	QLineEdit_CursorWordForward(this.h, (bool)(mark))
}

func (this *QLineEdit) CursorWordBackward(mark bool) {
	QLineEdit_CursorWordBackward(this.h, (bool)(mark))
}

func (this *QLineEdit) Backspace() {
	QLineEdit_Backspace(this.h)
}

func (this *QLineEdit) Del() {
	QLineEdit_Del(this.h)
}

func (this *QLineEdit) Home(mark bool) {
	QLineEdit_Home(this.h, (bool)(mark))
}

func (this *QLineEdit) End(mark bool) {
	QLineEdit_End(this.h, (bool)(mark))
}

func (this *QLineEdit) IsModified() bool {
	return (bool)(QLineEdit_IsModified(this.h))
}

func (this *QLineEdit) SetModified(modified bool) {
	QLineEdit_SetModified(this.h, (bool)(modified))
}

func (this *QLineEdit) SetSelection(param1 int, param2 int) {
	QLineEdit_SetSelection(this.h, (int)(param1), (int)(param2))
}

func (this *QLineEdit) HasSelectedText() bool {
	return (bool)(QLineEdit_HasSelectedText(this.h))
}

func (this *QLineEdit) SelectedText() string {
	var _ms struct_miqt_string = QLineEdit_SelectedText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLineEdit) SelectionStart() int {
	return (int)(QLineEdit_SelectionStart(this.h))
}

func (this *QLineEdit) SelectionEnd() int {
	return (int)(QLineEdit_SelectionEnd(this.h))
}

func (this *QLineEdit) SelectionLength() int {
	return (int)(QLineEdit_SelectionLength(this.h))
}

func (this *QLineEdit) IsUndoAvailable() bool {
	return (bool)(QLineEdit_IsUndoAvailable(this.h))
}

func (this *QLineEdit) IsRedoAvailable() bool {
	return (bool)(QLineEdit_IsRedoAvailable(this.h))
}

func (this *QLineEdit) SetDragEnabled(b bool) {
	QLineEdit_SetDragEnabled(this.h, (bool)(b))
}

func (this *QLineEdit) DragEnabled() bool {
	return (bool)(QLineEdit_DragEnabled(this.h))
}

func (this *QLineEdit) SetCursorMoveStyle(style CursorMoveStyle) {
	QLineEdit_SetCursorMoveStyle(this.h, (int)(style))
}

func (this *QLineEdit) CursorMoveStyle() CursorMoveStyle {
	return (CursorMoveStyle)(QLineEdit_CursorMoveStyle(this.h))
}

func (this *QLineEdit) InputMask() string {
	var _ms struct_miqt_string = QLineEdit_InputMask(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLineEdit) SetInputMask(inputMask string) {
	inputMask_ms := struct_miqt_string{}
	inputMask_ms.data = CString(inputMask)
	inputMask_ms.len = size_t(len(inputMask))
	defer free(unsafe.Pointer(inputMask_ms.data))
	QLineEdit_SetInputMask(this.h, inputMask_ms)
}

func (this *QLineEdit) HasAcceptableInput() bool {
	return (bool)(QLineEdit_HasAcceptableInput(this.h))
}

func (this *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	QLineEdit_SetTextMargins(this.h, (int)(left), (int)(top), (int)(right), (int)(bottom))
}

func (this *QLineEdit) SetTextMarginsWithMargins(margins *QMargins) {
	QLineEdit_SetTextMarginsWithMargins(this.h, margins.cPointer())
}

func (this *QLineEdit) TextMargins() *QMargins {
	_goptr := newQMargins(QLineEdit_TextMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineEdit) AddAction(action *QAction, position ActionPosition) {
	QLineEdit_AddAction(this.h, action.cPointer(), position)
}

func (this *QLineEdit) AddAction2(icon *QIcon, position ActionPosition) *QAction {
	return newQAction(QLineEdit_AddAction2(this.h, icon.cPointer(), position))
}

func (this *QLineEdit) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QLineEdit_SetText(this.h, text_ms)
}

func (this *QLineEdit) Clear() {
	QLineEdit_Clear(this.h)
}

func (this *QLineEdit) SelectAll() {
	QLineEdit_SelectAll(this.h)
}

func (this *QLineEdit) Undo() {
	QLineEdit_Undo(this.h)
}

func (this *QLineEdit) Redo() {
	QLineEdit_Redo(this.h)
}

func (this *QLineEdit) Cut() {
	QLineEdit_Cut(this.h)
}

func (this *QLineEdit) Copy() {
	QLineEdit_Copy(this.h)
}

func (this *QLineEdit) Paste() {
	QLineEdit_Paste(this.h)
}

func (this *QLineEdit) Deselect() {
	QLineEdit_Deselect(this.h)
}

func (this *QLineEdit) Insert(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QLineEdit_Insert(this.h, param1_ms)
}

func (this *QLineEdit) CreateStandardContextMenu() *QMenu {
	return newQMenu(QLineEdit_CreateStandardContextMenu(this.h))
}

func (this *QLineEdit) TextChanged(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QLineEdit_TextChanged(this.h, param1_ms)
}
func (this *QLineEdit) OnTextChanged(slot func(param1 string)) {
	QLineEdit_connect_TextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_TextChanged
func miqt_exec_callback_QLineEdit_TextChanged(cb intptr_t, param1 struct_miqt_string) {
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

func (this *QLineEdit) TextEdited(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QLineEdit_TextEdited(this.h, param1_ms)
}
func (this *QLineEdit) OnTextEdited(slot func(param1 string)) {
	QLineEdit_connect_TextEdited(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_TextEdited
func miqt_exec_callback_QLineEdit_TextEdited(cb intptr_t, param1 struct_miqt_string) {
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

func (this *QLineEdit) CursorPositionChanged(param1 int, param2 int) {
	QLineEdit_CursorPositionChanged(this.h, (int)(param1), (int)(param2))
}
func (this *QLineEdit) OnCursorPositionChanged(slot func(param1 int, param2 int)) {
	QLineEdit_connect_CursorPositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_CursorPositionChanged
func miqt_exec_callback_QLineEdit_CursorPositionChanged(cb intptr_t, param1 int, param2 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int, param2 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	slotval2 := (int)(param2)

	gofunc(slotval1, slotval2)
}

func (this *QLineEdit) ReturnPressed() {
	QLineEdit_ReturnPressed(this.h)
}
func (this *QLineEdit) OnReturnPressed(slot func()) {
	QLineEdit_connect_ReturnPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_ReturnPressed
func miqt_exec_callback_QLineEdit_ReturnPressed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QLineEdit) EditingFinished() {
	QLineEdit_EditingFinished(this.h)
}
func (this *QLineEdit) OnEditingFinished(slot func()) {
	QLineEdit_connect_EditingFinished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_EditingFinished
func miqt_exec_callback_QLineEdit_EditingFinished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QLineEdit) SelectionChanged() {
	QLineEdit_SelectionChanged(this.h)
}
func (this *QLineEdit) OnSelectionChanged(slot func()) {
	QLineEdit_connect_SelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_SelectionChanged
func miqt_exec_callback_QLineEdit_SelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QLineEdit) InputRejected() {
	QLineEdit_InputRejected(this.h)
}
func (this *QLineEdit) OnInputRejected(slot func()) {
	QLineEdit_connect_InputRejected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_InputRejected
func miqt_exec_callback_QLineEdit_InputRejected(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QLineEdit) InputMethodQuery(param1 InputMethodQuery) *QVariant {
	_goptr := newQVariant(QLineEdit_InputMethodQuery(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineEdit) InputMethodQuery2(property InputMethodQuery, argument QVariant) *QVariant {
	_goptr := newQVariant(QLineEdit_InputMethodQuery2(this.h, (int)(property), argument.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineEdit) TimerEvent(param1 *QTimerEvent) {
	QLineEdit_TimerEvent(this.h, param1.cPointer())
}

func (this *QLineEdit) Event(param1 *QEvent) bool {
	return (bool)(QLineEdit_Event(this.h, param1.cPointer()))
}

func QLineEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLineEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLineEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLineEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLineEdit) CursorForward2(mark bool, steps int) {
	QLineEdit_CursorForward2(this.h, (bool)(mark), (int)(steps))
}

func (this *QLineEdit) CursorBackward2(mark bool, steps int) {
	QLineEdit_CursorBackward2(this.h, (bool)(mark), (int)(steps))
}

func (this *QLineEdit) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QLineEdit_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QLineEdit) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_SizeHint
func miqt_exec_callback_QLineEdit_SizeHint(self QLineEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QLineEdit) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QLineEdit_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QLineEdit) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_MinimumSizeHint
func miqt_exec_callback_QLineEdit_MinimumSizeHint(self QLineEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QLineEdit) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QLineEdit_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_MousePressEvent
func miqt_exec_callback_QLineEdit_MousePressEvent(self QLineEdit, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QLineEdit_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_MouseMoveEvent
func miqt_exec_callback_QLineEdit_MouseMoveEvent(self QLineEdit, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QLineEdit_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_MouseReleaseEvent
func miqt_exec_callback_QLineEdit_MouseReleaseEvent(self QLineEdit, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_MouseDoubleClickEvent(param1 *QMouseEvent) {

	QLineEdit_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnMouseDoubleClickEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_MouseDoubleClickEvent
func miqt_exec_callback_QLineEdit_MouseDoubleClickEvent(self QLineEdit, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QLineEdit_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_KeyPressEvent
func miqt_exec_callback_QLineEdit_KeyPressEvent(self QLineEdit, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_KeyReleaseEvent(param1 *QKeyEvent) {

	QLineEdit_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnKeyReleaseEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_KeyReleaseEvent
func miqt_exec_callback_QLineEdit_KeyReleaseEvent(self QLineEdit, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_FocusInEvent(param1 *QFocusEvent) {

	QLineEdit_virtualbase_FocusInEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnFocusInEvent(slot func(super func(param1 *QFocusEvent), param1 *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_FocusInEvent
func miqt_exec_callback_QLineEdit_FocusInEvent(self QLineEdit, cb intptr_t, param1 *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFocusEvent), param1 *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_FocusOutEvent(param1 *QFocusEvent) {

	QLineEdit_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnFocusOutEvent(slot func(super func(param1 *QFocusEvent), param1 *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_FocusOutEvent
func miqt_exec_callback_QLineEdit_FocusOutEvent(self QLineEdit, cb intptr_t, param1 *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFocusEvent), param1 *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QLineEdit_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_PaintEvent
func miqt_exec_callback_QLineEdit_PaintEvent(self QLineEdit, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_DragEnterEvent(param1 *QDragEnterEvent) {

	QLineEdit_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnDragEnterEvent(slot func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_DragEnterEvent
func miqt_exec_callback_QLineEdit_DragEnterEvent(self QLineEdit, cb intptr_t, param1 *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDragEnterEvent), param1 *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_DragMoveEvent(e *QDragMoveEvent) {

	QLineEdit_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QLineEdit) OnDragMoveEvent(slot func(super func(e *QDragMoveEvent), e *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_DragMoveEvent
func miqt_exec_callback_QLineEdit_DragMoveEvent(self QLineEdit, cb intptr_t, e *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragMoveEvent), e *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(e)

	gofunc((&QLineEdit{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_DragLeaveEvent(e *QDragLeaveEvent) {

	QLineEdit_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QLineEdit) OnDragLeaveEvent(slot func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_DragLeaveEvent
func miqt_exec_callback_QLineEdit_DragLeaveEvent(self QLineEdit, cb intptr_t, e *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(e)

	gofunc((&QLineEdit{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_DropEvent(param1 *QDropEvent) {

	QLineEdit_virtualbase_DropEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnDropEvent(slot func(super func(param1 *QDropEvent), param1 *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_DropEvent
func miqt_exec_callback_QLineEdit_DropEvent(self QLineEdit, cb intptr_t, param1 *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QDropEvent), param1 *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QLineEdit_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_ChangeEvent
func miqt_exec_callback_QLineEdit_ChangeEvent(self QLineEdit, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QLineEdit_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_ContextMenuEvent
func miqt_exec_callback_QLineEdit_ContextMenuEvent(self QLineEdit, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QLineEdit_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_InputMethodEvent
func miqt_exec_callback_QLineEdit_InputMethodEvent(self QLineEdit, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_InitStyleOption(option *QStyleOptionFrame) {

	QLineEdit_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QLineEdit) OnInitStyleOption(slot func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_InitStyleOption
func miqt_exec_callback_QLineEdit_InitStyleOption(self QLineEdit, cb intptr_t, option *QStyleOptionFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionFrame(option)

	gofunc((&QLineEdit{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QLineEdit) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QLineEdit_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QLineEdit) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_InputMethodQuery
func miqt_exec_callback_QLineEdit_InputMethodQuery(self QLineEdit, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QLineEdit) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QLineEdit_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QLineEdit) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_TimerEvent
func miqt_exec_callback_QLineEdit_TimerEvent(self QLineEdit, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QLineEdit{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QLineEdit_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QLineEdit) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_Event
func miqt_exec_callback_QLineEdit_Event(self QLineEdit, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLineEdit) callVirtualBase_DevType() int {

	return (int)(QLineEdit_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QLineEdit) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_DevType
func miqt_exec_callback_QLineEdit_DevType(self QLineEdit, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QLineEdit) callVirtualBase_SetVisible(visible bool) {

	QLineEdit_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QLineEdit) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_SetVisible
func miqt_exec_callback_QLineEdit_SetVisible(self QLineEdit, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QLineEdit{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QLineEdit) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QLineEdit_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QLineEdit) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_HeightForWidth
func miqt_exec_callback_QLineEdit_HeightForWidth(self QLineEdit, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QLineEdit) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QLineEdit_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QLineEdit) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_HasHeightForWidth
func miqt_exec_callback_QLineEdit_HasHeightForWidth(self QLineEdit, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QLineEdit) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QLineEdit_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QLineEdit) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_PaintEngine
func miqt_exec_callback_QLineEdit_PaintEngine(self QLineEdit, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QLineEdit) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QLineEdit_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_WheelEvent
func miqt_exec_callback_QLineEdit_WheelEvent(self QLineEdit, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QLineEdit_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_EnterEvent
func miqt_exec_callback_QLineEdit_EnterEvent(self QLineEdit, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_LeaveEvent(event *QEvent) {

	QLineEdit_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_LeaveEvent
func miqt_exec_callback_QLineEdit_LeaveEvent(self QLineEdit, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QLineEdit_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_MoveEvent
func miqt_exec_callback_QLineEdit_MoveEvent(self QLineEdit, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QLineEdit_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_ResizeEvent
func miqt_exec_callback_QLineEdit_ResizeEvent(self QLineEdit, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QLineEdit_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_CloseEvent
func miqt_exec_callback_QLineEdit_CloseEvent(self QLineEdit, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QLineEdit_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_TabletEvent
func miqt_exec_callback_QLineEdit_TabletEvent(self QLineEdit, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_ActionEvent(event *QActionEvent) {

	QLineEdit_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_ActionEvent
func miqt_exec_callback_QLineEdit_ActionEvent(self QLineEdit, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_ShowEvent(event *QShowEvent) {

	QLineEdit_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_ShowEvent
func miqt_exec_callback_QLineEdit_ShowEvent(self QLineEdit, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_HideEvent(event *QHideEvent) {

	QLineEdit_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLineEdit) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_HideEvent
func miqt_exec_callback_QLineEdit_HideEvent(self QLineEdit, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QLineEdit{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QLineEdit) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QLineEdit_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QLineEdit) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_NativeEvent
func miqt_exec_callback_QLineEdit_NativeEvent(self QLineEdit, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QLineEdit) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QLineEdit_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QLineEdit) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_Metric
func miqt_exec_callback_QLineEdit_Metric(self QLineEdit, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QLineEdit) callVirtualBase_InitPainter(painter *QPainter) {

	QLineEdit_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QLineEdit) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_InitPainter
func miqt_exec_callback_QLineEdit_InitPainter(self QLineEdit, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QLineEdit{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QLineEdit) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QLineEdit_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QLineEdit) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_Redirected
func miqt_exec_callback_QLineEdit_Redirected(self QLineEdit, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QLineEdit) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QLineEdit_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QLineEdit) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_SharedPainter
func miqt_exec_callback_QLineEdit_SharedPainter(self QLineEdit, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QLineEdit) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QLineEdit_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QLineEdit) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_FocusNextPrevChild
func miqt_exec_callback_QLineEdit_FocusNextPrevChild(self QLineEdit, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
