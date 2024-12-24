package qt

import (
	"unsafe"
)

type QPlainTextEdit__LineWrapMode int

const (
	QPlainTextEdit__NoWrap      QPlainTextEdit__LineWrapMode = 0
	QPlainTextEdit__WidgetWidth QPlainTextEdit__LineWrapMode = 1
)

type QPlainTextEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQPlainTextEdit constructs a new QPlainTextEdit object.
func NewQPlainTextEdit(parent *QWidget) *QPlainTextEdit {
	ret := newQPlainTextEdit(QPlainTextEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPlainTextEdit2 constructs a new QPlainTextEdit object.
func NewQPlainTextEdit2() *QPlainTextEdit {
	ret := newQPlainTextEdit(QPlainTextEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQPlainTextEdit3 constructs a new QPlainTextEdit object.
func NewQPlainTextEdit3(text string) *QPlainTextEdit {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQPlainTextEdit(QPlainTextEdit_new3(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQPlainTextEdit4 constructs a new QPlainTextEdit object.
func NewQPlainTextEdit4(text string, parent *QWidget) *QPlainTextEdit {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQPlainTextEdit(QPlainTextEdit_new4(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPlainTextEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QPlainTextEdit_MetaObject(this.h))
}

func (this *QPlainTextEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPlainTextEdit_Metacast(this.h, param1_Cstring))
}

func QPlainTextEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPlainTextEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextEdit) SetDocument(document *QTextDocument) {
	QPlainTextEdit_SetDocument(this.h, document.cPointer())
}

func (this *QPlainTextEdit) Document() *QTextDocument {
	return newQTextDocument(QPlainTextEdit_Document(this.h))
}

func (this *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	placeholderText_ms := struct_miqt_string{}
	placeholderText_ms.data = CString(placeholderText)
	placeholderText_ms.len = size_t(len(placeholderText))
	defer free(unsafe.Pointer(placeholderText_ms.data))
	QPlainTextEdit_SetPlaceholderText(this.h, placeholderText_ms)
}

func (this *QPlainTextEdit) PlaceholderText() string {
	var _ms struct_miqt_string = QPlainTextEdit_PlaceholderText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextEdit) SetTextCursor(cursor *QTextCursor) {
	QPlainTextEdit_SetTextCursor(this.h, cursor.cPointer())
}

func (this *QPlainTextEdit) TextCursor() *QTextCursor {
	_goptr := newQTextCursor(QPlainTextEdit_TextCursor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) IsReadOnly() bool {
	return (bool)(QPlainTextEdit_IsReadOnly(this.h))
}

func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	QPlainTextEdit_SetReadOnly(this.h, (bool)(ro))
}

func (this *QPlainTextEdit) SetTextInteractionFlags(flags TextInteractionFlag) {
	QPlainTextEdit_SetTextInteractionFlags(this.h, (int)(flags))
}

func (this *QPlainTextEdit) TextInteractionFlags() TextInteractionFlag {
	return (TextInteractionFlag)(QPlainTextEdit_TextInteractionFlags(this.h))
}

func (this *QPlainTextEdit) MergeCurrentCharFormat(modifier *QTextCharFormat) {
	QPlainTextEdit_MergeCurrentCharFormat(this.h, modifier.cPointer())
}

func (this *QPlainTextEdit) SetCurrentCharFormat(format *QTextCharFormat) {
	QPlainTextEdit_SetCurrentCharFormat(this.h, format.cPointer())
}

func (this *QPlainTextEdit) CurrentCharFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QPlainTextEdit_CurrentCharFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) TabChangesFocus() bool {
	return (bool)(QPlainTextEdit_TabChangesFocus(this.h))
}

func (this *QPlainTextEdit) SetTabChangesFocus(b bool) {
	QPlainTextEdit_SetTabChangesFocus(this.h, (bool)(b))
}

func (this *QPlainTextEdit) SetDocumentTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QPlainTextEdit_SetDocumentTitle(this.h, title_ms)
}

func (this *QPlainTextEdit) DocumentTitle() string {
	var _ms struct_miqt_string = QPlainTextEdit_DocumentTitle(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextEdit) IsUndoRedoEnabled() bool {
	return (bool)(QPlainTextEdit_IsUndoRedoEnabled(this.h))
}

func (this *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	QPlainTextEdit_SetUndoRedoEnabled(this.h, (bool)(enable))
}

func (this *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	QPlainTextEdit_SetMaximumBlockCount(this.h, (int)(maximum))
}

func (this *QPlainTextEdit) MaximumBlockCount() int {
	return (int)(QPlainTextEdit_MaximumBlockCount(this.h))
}

func (this *QPlainTextEdit) LineWrapMode() LineWrapMode {
	xxxxxxxxx
}

func (this *QPlainTextEdit) SetLineWrapMode(mode LineWrapMode) {
	QPlainTextEdit_SetLineWrapMode(this.h, mode)
}

func (this *QPlainTextEdit) WordWrapMode() QTextOption__WrapMode {
	return (QTextOption__WrapMode)(QPlainTextEdit_WordWrapMode(this.h))
}

func (this *QPlainTextEdit) SetWordWrapMode(policy QTextOption__WrapMode) {
	QPlainTextEdit_SetWordWrapMode(this.h, (int)(policy))
}

func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	QPlainTextEdit_SetBackgroundVisible(this.h, (bool)(visible))
}

func (this *QPlainTextEdit) BackgroundVisible() bool {
	return (bool)(QPlainTextEdit_BackgroundVisible(this.h))
}

func (this *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	QPlainTextEdit_SetCenterOnScroll(this.h, (bool)(enabled))
}

func (this *QPlainTextEdit) CenterOnScroll() bool {
	return (bool)(QPlainTextEdit_CenterOnScroll(this.h))
}

func (this *QPlainTextEdit) Find(exp string) bool {
	exp_ms := struct_miqt_string{}
	exp_ms.data = CString(exp)
	exp_ms.len = size_t(len(exp))
	defer free(unsafe.Pointer(exp_ms.data))
	return (bool)(QPlainTextEdit_Find(this.h, exp_ms))
}

func (this *QPlainTextEdit) FindWithExp(exp *QRegularExpression) bool {
	return (bool)(QPlainTextEdit_FindWithExp(this.h, exp.cPointer()))
}

func (this *QPlainTextEdit) ToPlainText() string {
	var _ms struct_miqt_string = QPlainTextEdit_ToPlainText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextEdit) EnsureCursorVisible() {
	QPlainTextEdit_EnsureCursorVisible(this.h)
}

func (this *QPlainTextEdit) LoadResource(typeVal int, name *QUrl) *QVariant {
	_goptr := newQVariant(QPlainTextEdit_LoadResource(this.h, (int)(typeVal), name.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) CreateStandardContextMenu() *QMenu {
	return newQMenu(QPlainTextEdit_CreateStandardContextMenu(this.h))
}

func (this *QPlainTextEdit) CreateStandardContextMenuWithPosition(position *QPoint) *QMenu {
	return newQMenu(QPlainTextEdit_CreateStandardContextMenuWithPosition(this.h, position.cPointer()))
}

func (this *QPlainTextEdit) CursorForPosition(pos *QPoint) *QTextCursor {
	_goptr := newQTextCursor(QPlainTextEdit_CursorForPosition(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) CursorRect(cursor *QTextCursor) *QRect {
	_goptr := newQRect(QPlainTextEdit_CursorRect(this.h, cursor.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) CursorRect2() *QRect {
	_goptr := newQRect(QPlainTextEdit_CursorRect2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) AnchorAt(pos *QPoint) string {
	var _ms struct_miqt_string = QPlainTextEdit_AnchorAt(this.h, pos.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextEdit) OverwriteMode() bool {
	return (bool)(QPlainTextEdit_OverwriteMode(this.h))
}

func (this *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	QPlainTextEdit_SetOverwriteMode(this.h, (bool)(overwrite))
}

func (this *QPlainTextEdit) TabStopDistance() float64 {
	return (float64)(QPlainTextEdit_TabStopDistance(this.h))
}

func (this *QPlainTextEdit) SetTabStopDistance(distance float64) {
	QPlainTextEdit_SetTabStopDistance(this.h, (double)(distance))
}

func (this *QPlainTextEdit) CursorWidth() int {
	return (int)(QPlainTextEdit_CursorWidth(this.h))
}

func (this *QPlainTextEdit) SetCursorWidth(width int) {
	QPlainTextEdit_SetCursorWidth(this.h, (int)(width))
}

func (this *QPlainTextEdit) SetExtraSelections(selections []QTextEdit__ExtraSelection) {
	selections_CArray := (*[0xffff]*QTextEdit__ExtraSelection)(malloc(size_t(8 * len(selections))))
	defer free(unsafe.Pointer(selections_CArray))
	for i := range selections {
		selections_CArray[i] = selections[i].cPointer()
	}
	selections_ma := struct_miqt_array{len: size_t(len(selections)), data: unsafe.Pointer(selections_CArray)}
	QPlainTextEdit_SetExtraSelections(this.h, selections_ma)
}

func (this *QPlainTextEdit) ExtraSelections() []QTextEdit__ExtraSelection {
	var _ma struct_miqt_array = QPlainTextEdit_ExtraSelections(this.h)
	_ret := make([]QTextEdit__ExtraSelection, int(_ma.len))
	_outCast := (*[0xffff]*QTextEdit__ExtraSelection)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQTextEdit__ExtraSelection(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPlainTextEdit) MoveCursor(operation QTextCursor__MoveOperation) {
	QPlainTextEdit_MoveCursor(this.h, (int)(operation))
}

func (this *QPlainTextEdit) CanPaste() bool {
	return (bool)(QPlainTextEdit_CanPaste(this.h))
}

func (this *QPlainTextEdit) Print(printer *QPagedPaintDevice) {
	QPlainTextEdit_Print(this.h, printer.cPointer())
}

func (this *QPlainTextEdit) BlockCount() int {
	return (int)(QPlainTextEdit_BlockCount(this.h))
}

func (this *QPlainTextEdit) InputMethodQuery(property InputMethodQuery) *QVariant {
	_goptr := newQVariant(QPlainTextEdit_InputMethodQuery(this.h, (int)(property)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) InputMethodQuery2(query InputMethodQuery, argument QVariant) *QVariant {
	_goptr := newQVariant(QPlainTextEdit_InputMethodQuery2(this.h, (int)(query), argument.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextEdit) SetPlainText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPlainTextEdit_SetPlainText(this.h, text_ms)
}

func (this *QPlainTextEdit) Cut() {
	QPlainTextEdit_Cut(this.h)
}

func (this *QPlainTextEdit) Copy() {
	QPlainTextEdit_Copy(this.h)
}

func (this *QPlainTextEdit) Paste() {
	QPlainTextEdit_Paste(this.h)
}

func (this *QPlainTextEdit) Undo() {
	QPlainTextEdit_Undo(this.h)
}

func (this *QPlainTextEdit) Redo() {
	QPlainTextEdit_Redo(this.h)
}

func (this *QPlainTextEdit) Clear() {
	QPlainTextEdit_Clear(this.h)
}

func (this *QPlainTextEdit) SelectAll() {
	QPlainTextEdit_SelectAll(this.h)
}

func (this *QPlainTextEdit) InsertPlainText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPlainTextEdit_InsertPlainText(this.h, text_ms)
}

func (this *QPlainTextEdit) AppendPlainText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPlainTextEdit_AppendPlainText(this.h, text_ms)
}

func (this *QPlainTextEdit) AppendHtml(html string) {
	html_ms := struct_miqt_string{}
	html_ms.data = CString(html)
	html_ms.len = size_t(len(html))
	defer free(unsafe.Pointer(html_ms.data))
	QPlainTextEdit_AppendHtml(this.h, html_ms)
}

func (this *QPlainTextEdit) CenterCursor() {
	QPlainTextEdit_CenterCursor(this.h)
}

func (this *QPlainTextEdit) ZoomIn() {
	QPlainTextEdit_ZoomIn(this.h)
}

func (this *QPlainTextEdit) ZoomOut() {
	QPlainTextEdit_ZoomOut(this.h)
}

func (this *QPlainTextEdit) TextChanged() {
	QPlainTextEdit_TextChanged(this.h)
}

func (this *QPlainTextEdit) OnTextChanged(slot func()) {
	QPlainTextEdit_connect_TextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_TextChanged
func miqt_exec_callback_QPlainTextEdit_TextChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QPlainTextEdit) UndoAvailable(b bool) {
	QPlainTextEdit_UndoAvailable(this.h, (bool)(b))
}

func (this *QPlainTextEdit) OnUndoAvailable(slot func(b bool)) {
	QPlainTextEdit_connect_UndoAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_UndoAvailable
func miqt_exec_callback_QPlainTextEdit_UndoAvailable(cb intptr_t, b bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(b bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(b)

	gofunc(slotval1)
}

func (this *QPlainTextEdit) RedoAvailable(b bool) {
	QPlainTextEdit_RedoAvailable(this.h, (bool)(b))
}

func (this *QPlainTextEdit) OnRedoAvailable(slot func(b bool)) {
	QPlainTextEdit_connect_RedoAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_RedoAvailable
func miqt_exec_callback_QPlainTextEdit_RedoAvailable(cb intptr_t, b bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(b bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(b)

	gofunc(slotval1)
}

func (this *QPlainTextEdit) CopyAvailable(b bool) {
	QPlainTextEdit_CopyAvailable(this.h, (bool)(b))
}

func (this *QPlainTextEdit) OnCopyAvailable(slot func(b bool)) {
	QPlainTextEdit_connect_CopyAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_CopyAvailable
func miqt_exec_callback_QPlainTextEdit_CopyAvailable(cb intptr_t, b bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(b bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(b)

	gofunc(slotval1)
}

func (this *QPlainTextEdit) SelectionChanged() {
	QPlainTextEdit_SelectionChanged(this.h)
}

func (this *QPlainTextEdit) OnSelectionChanged(slot func()) {
	QPlainTextEdit_connect_SelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_SelectionChanged
func miqt_exec_callback_QPlainTextEdit_SelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QPlainTextEdit) CursorPositionChanged() {
	QPlainTextEdit_CursorPositionChanged(this.h)
}

func (this *QPlainTextEdit) OnCursorPositionChanged(slot func()) {
	QPlainTextEdit_connect_CursorPositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_CursorPositionChanged
func miqt_exec_callback_QPlainTextEdit_CursorPositionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QPlainTextEdit) UpdateRequest(rect *QRect, dy int) {
	QPlainTextEdit_UpdateRequest(this.h, rect.cPointer(), (int)(dy))
}

func (this *QPlainTextEdit) OnUpdateRequest(slot func(rect *QRect, dy int)) {
	QPlainTextEdit_connect_UpdateRequest(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_UpdateRequest
func miqt_exec_callback_QPlainTextEdit_UpdateRequest(cb intptr_t, rect *QRect, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(rect *QRect, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (int)(dy)

	gofunc(slotval1, slotval2)
}

func (this *QPlainTextEdit) BlockCountChanged(newBlockCount int) {
	QPlainTextEdit_BlockCountChanged(this.h, (int)(newBlockCount))
}

func (this *QPlainTextEdit) OnBlockCountChanged(slot func(newBlockCount int)) {
	QPlainTextEdit_connect_BlockCountChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_BlockCountChanged
func miqt_exec_callback_QPlainTextEdit_BlockCountChanged(cb intptr_t, newBlockCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newBlockCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(newBlockCount)

	gofunc(slotval1)
}

func (this *QPlainTextEdit) ModificationChanged(param1 bool) {
	QPlainTextEdit_ModificationChanged(this.h, (bool)(param1))
}

func (this *QPlainTextEdit) OnModificationChanged(slot func(param1 bool)) {
	QPlainTextEdit_connect_ModificationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_ModificationChanged
func miqt_exec_callback_QPlainTextEdit_ModificationChanged(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func QPlainTextEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPlainTextEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPlainTextEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPlainTextEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextEdit) Find2(exp string, options FindFlag) bool {
	exp_ms := struct_miqt_string{}
	exp_ms.data = CString(exp)
	exp_ms.len = size_t(len(exp))
	defer free(unsafe.Pointer(exp_ms.data))
	return (bool)(QPlainTextEdit_Find2(this.h, exp_ms, (int)(options)))
}

func (this *QPlainTextEdit) Find22(exp *QRegularExpression, options FindFlag) bool {
	return (bool)(QPlainTextEdit_Find22(this.h, exp.cPointer(), (int)(options)))
}

func (this *QPlainTextEdit) MoveCursor2(operation QTextCursor__MoveOperation, mode QTextCursor__MoveMode) {
	QPlainTextEdit_MoveCursor2(this.h, (int)(operation), (int)(mode))
}

func (this *QPlainTextEdit) ZoomIn1(rangeVal int) {
	QPlainTextEdit_ZoomIn1(this.h, (int)(rangeVal))
}

func (this *QPlainTextEdit) ZoomOut1(rangeVal int) {
	QPlainTextEdit_ZoomOut1(this.h, (int)(rangeVal))
}

func (this *QPlainTextEdit) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPlainTextEdit_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPlainTextEdit) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPlainTextEdit_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_MetaObject
func miqt_exec_callback_QPlainTextEdit_MetaObject(self QPlainTextEdit, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPlainTextEdit{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPlainTextEdit) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPlainTextEdit_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPlainTextEdit) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPlainTextEdit_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextEdit_Metacast
func miqt_exec_callback_QPlainTextEdit_Metacast(self QPlainTextEdit, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QPlainTextEdit{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QPlainTextDocumentLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQPlainTextDocumentLayout constructs a new QPlainTextDocumentLayout object.
func NewQPlainTextDocumentLayout(document *QTextDocument) *QPlainTextDocumentLayout {
	ret := newQPlainTextDocumentLayout(QPlainTextDocumentLayout_new(document.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPlainTextDocumentLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QPlainTextDocumentLayout_MetaObject(this.h))
}

func (this *QPlainTextDocumentLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPlainTextDocumentLayout_Metacast(this.h, param1_Cstring))
}

func QPlainTextDocumentLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPlainTextDocumentLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextDocumentLayout) Draw(param1 *QPainter, param2 *PaintContext) {
	QPlainTextDocumentLayout_Draw(this.h, param1.cPointer(), param2)
}

func (this *QPlainTextDocumentLayout) HitTest(param1 *QPointF, param2 HitTestAccuracy) int {
	return (int)(QPlainTextDocumentLayout_HitTest(this.h, param1.cPointer(), (int)(param2)))
}

func (this *QPlainTextDocumentLayout) PageCount() int {
	return (int)(QPlainTextDocumentLayout_PageCount(this.h))
}

func (this *QPlainTextDocumentLayout) DocumentSize() *QSizeF {
	_goptr := newQSizeF(QPlainTextDocumentLayout_DocumentSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextDocumentLayout) FrameBoundingRect(param1 *QTextFrame) *QRectF {
	_goptr := newQRectF(QPlainTextDocumentLayout_FrameBoundingRect(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextDocumentLayout) BlockBoundingRect(block *QTextBlock) *QRectF {
	_goptr := newQRectF(QPlainTextDocumentLayout_BlockBoundingRect(this.h, block.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPlainTextDocumentLayout) EnsureBlockLayout(block *QTextBlock) {
	QPlainTextDocumentLayout_EnsureBlockLayout(this.h, block.cPointer())
}

func (this *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	QPlainTextDocumentLayout_SetCursorWidth(this.h, (int)(width))
}

func (this *QPlainTextDocumentLayout) CursorWidth() int {
	return (int)(QPlainTextDocumentLayout_CursorWidth(this.h))
}

func (this *QPlainTextDocumentLayout) RequestUpdate() {
	QPlainTextDocumentLayout_RequestUpdate(this.h)
}

func QPlainTextDocumentLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPlainTextDocumentLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPlainTextDocumentLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPlainTextDocumentLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPlainTextDocumentLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPlainTextDocumentLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPlainTextDocumentLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPlainTextDocumentLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextDocumentLayout_MetaObject
func miqt_exec_callback_QPlainTextDocumentLayout_MetaObject(self QPlainTextDocumentLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPlainTextDocumentLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPlainTextDocumentLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPlainTextDocumentLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPlainTextDocumentLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPlainTextDocumentLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPlainTextDocumentLayout_Metacast
func miqt_exec_callback_QPlainTextDocumentLayout_Metacast(self QPlainTextDocumentLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QPlainTextDocumentLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
