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

func (this *QLineEdit) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QLineEdit_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QLineEdit) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_MetaObject
func miqt_exec_callback_QLineEdit_MetaObject(self QLineEdit, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QLineEdit) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QLineEdit_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QLineEdit) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLineEdit_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLineEdit_Metacast
func miqt_exec_callback_QLineEdit_Metacast(self QLineEdit, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QLineEdit{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
