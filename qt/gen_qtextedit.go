package qt

import (
	"unsafe"
)

type QTextEdit__LineWrapMode int

const (
	QTextEdit__NoWrap           QTextEdit__LineWrapMode = 0
	QTextEdit__WidgetWidth      QTextEdit__LineWrapMode = 1
	QTextEdit__FixedPixelWidth  QTextEdit__LineWrapMode = 2
	QTextEdit__FixedColumnWidth QTextEdit__LineWrapMode = 3
)

type QTextEdit__AutoFormattingFlag int

const (
	QTextEdit__AutoNone       QTextEdit__AutoFormattingFlag = 0
	QTextEdit__AutoBulletList QTextEdit__AutoFormattingFlag = 1
	QTextEdit__AutoAll        QTextEdit__AutoFormattingFlag = 4294967295
)

type QTextEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQTextEdit constructs a new QTextEdit object.
func NewQTextEdit(parent *QWidget) *QTextEdit {
	ret := newQTextEdit(QTextEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextEdit2 constructs a new QTextEdit object.
func NewQTextEdit2() *QTextEdit {
	ret := newQTextEdit(QTextEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQTextEdit3 constructs a new QTextEdit object.
func NewQTextEdit3(text string) *QTextEdit {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTextEdit(QTextEdit_new3(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQTextEdit4 constructs a new QTextEdit object.
func NewQTextEdit4(text string, parent *QWidget) *QTextEdit {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTextEdit(QTextEdit_new4(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QTextEdit_MetaObject(this.h))
}

func (this *QTextEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextEdit_Metacast(this.h, param1_Cstring))
}

func QTextEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) SetDocument(document *QTextDocument) {
	QTextEdit_SetDocument(this.h, document.cPointer())
}

func (this *QTextEdit) Document() *QTextDocument {
	return newQTextDocument(QTextEdit_Document(this.h))
}

func (this *QTextEdit) SetPlaceholderText(placeholderText string) {
	placeholderText_ms := struct_miqt_string{}
	placeholderText_ms.data = CString(placeholderText)
	placeholderText_ms.len = size_t(len(placeholderText))
	defer free(unsafe.Pointer(placeholderText_ms.data))
	QTextEdit_SetPlaceholderText(this.h, placeholderText_ms)
}

func (this *QTextEdit) PlaceholderText() string {
	var _ms struct_miqt_string = QTextEdit_PlaceholderText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) SetTextCursor(cursor *QTextCursor) {
	QTextEdit_SetTextCursor(this.h, cursor.cPointer())
}

func (this *QTextEdit) TextCursor() *QTextCursor {
	_goptr := newQTextCursor(QTextEdit_TextCursor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) IsReadOnly() bool {
	return (bool)(QTextEdit_IsReadOnly(this.h))
}

func (this *QTextEdit) SetReadOnly(ro bool) {
	QTextEdit_SetReadOnly(this.h, (bool)(ro))
}

func (this *QTextEdit) SetTextInteractionFlags(flags TextInteractionFlag) {
	QTextEdit_SetTextInteractionFlags(this.h, (int)(flags))
}

func (this *QTextEdit) TextInteractionFlags() TextInteractionFlag {
	return (TextInteractionFlag)(QTextEdit_TextInteractionFlags(this.h))
}

func (this *QTextEdit) FontPointSize() float64 {
	return (float64)(QTextEdit_FontPointSize(this.h))
}

func (this *QTextEdit) FontFamily() string {
	var _ms struct_miqt_string = QTextEdit_FontFamily(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) FontWeight() int {
	return (int)(QTextEdit_FontWeight(this.h))
}

func (this *QTextEdit) FontUnderline() bool {
	return (bool)(QTextEdit_FontUnderline(this.h))
}

func (this *QTextEdit) FontItalic() bool {
	return (bool)(QTextEdit_FontItalic(this.h))
}

func (this *QTextEdit) TextColor() *QColor {
	_goptr := newQColor(QTextEdit_TextColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) TextBackgroundColor() *QColor {
	_goptr := newQColor(QTextEdit_TextBackgroundColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) CurrentFont() *QFont {
	_goptr := newQFont(QTextEdit_CurrentFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QTextEdit_Alignment(this.h))
}

func (this *QTextEdit) MergeCurrentCharFormat(modifier *QTextCharFormat) {
	QTextEdit_MergeCurrentCharFormat(this.h, modifier.cPointer())
}

func (this *QTextEdit) SetCurrentCharFormat(format *QTextCharFormat) {
	QTextEdit_SetCurrentCharFormat(this.h, format.cPointer())
}

func (this *QTextEdit) CurrentCharFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QTextEdit_CurrentCharFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) AutoFormatting() AutoFormatting {
	xxxxxxxxx
}

func (this *QTextEdit) SetAutoFormatting(features AutoFormatting) {
	QTextEdit_SetAutoFormatting(this.h, features)
}

func (this *QTextEdit) TabChangesFocus() bool {
	return (bool)(QTextEdit_TabChangesFocus(this.h))
}

func (this *QTextEdit) SetTabChangesFocus(b bool) {
	QTextEdit_SetTabChangesFocus(this.h, (bool)(b))
}

func (this *QTextEdit) SetDocumentTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QTextEdit_SetDocumentTitle(this.h, title_ms)
}

func (this *QTextEdit) DocumentTitle() string {
	var _ms struct_miqt_string = QTextEdit_DocumentTitle(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) IsUndoRedoEnabled() bool {
	return (bool)(QTextEdit_IsUndoRedoEnabled(this.h))
}

func (this *QTextEdit) SetUndoRedoEnabled(enable bool) {
	QTextEdit_SetUndoRedoEnabled(this.h, (bool)(enable))
}

func (this *QTextEdit) LineWrapMode() LineWrapMode {
	xxxxxxxxx
}

func (this *QTextEdit) SetLineWrapMode(mode LineWrapMode) {
	QTextEdit_SetLineWrapMode(this.h, mode)
}

func (this *QTextEdit) LineWrapColumnOrWidth() int {
	return (int)(QTextEdit_LineWrapColumnOrWidth(this.h))
}

func (this *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	QTextEdit_SetLineWrapColumnOrWidth(this.h, (int)(w))
}

func (this *QTextEdit) WordWrapMode() QTextOption__WrapMode {
	return (QTextOption__WrapMode)(QTextEdit_WordWrapMode(this.h))
}

func (this *QTextEdit) SetWordWrapMode(policy QTextOption__WrapMode) {
	QTextEdit_SetWordWrapMode(this.h, (int)(policy))
}

func (this *QTextEdit) Find(exp string) bool {
	exp_ms := struct_miqt_string{}
	exp_ms.data = CString(exp)
	exp_ms.len = size_t(len(exp))
	defer free(unsafe.Pointer(exp_ms.data))
	return (bool)(QTextEdit_Find(this.h, exp_ms))
}

func (this *QTextEdit) FindWithExp(exp *QRegularExpression) bool {
	return (bool)(QTextEdit_FindWithExp(this.h, exp.cPointer()))
}

func (this *QTextEdit) ToPlainText() string {
	var _ms struct_miqt_string = QTextEdit_ToPlainText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) ToHtml() string {
	var _ms struct_miqt_string = QTextEdit_ToHtml(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) ToMarkdown() string {
	var _ms struct_miqt_string = QTextEdit_ToMarkdown(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) EnsureCursorVisible() {
	QTextEdit_EnsureCursorVisible(this.h)
}

func (this *QTextEdit) LoadResource(typeVal int, name *QUrl) *QVariant {
	_goptr := newQVariant(QTextEdit_LoadResource(this.h, (int)(typeVal), name.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) CreateStandardContextMenu() *QMenu {
	return newQMenu(QTextEdit_CreateStandardContextMenu(this.h))
}

func (this *QTextEdit) CreateStandardContextMenuWithPosition(position *QPoint) *QMenu {
	return newQMenu(QTextEdit_CreateStandardContextMenuWithPosition(this.h, position.cPointer()))
}

func (this *QTextEdit) CursorForPosition(pos *QPoint) *QTextCursor {
	_goptr := newQTextCursor(QTextEdit_CursorForPosition(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) CursorRect(cursor *QTextCursor) *QRect {
	_goptr := newQRect(QTextEdit_CursorRect(this.h, cursor.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) CursorRect2() *QRect {
	_goptr := newQRect(QTextEdit_CursorRect2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) AnchorAt(pos *QPoint) string {
	var _ms struct_miqt_string = QTextEdit_AnchorAt(this.h, pos.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) OverwriteMode() bool {
	return (bool)(QTextEdit_OverwriteMode(this.h))
}

func (this *QTextEdit) SetOverwriteMode(overwrite bool) {
	QTextEdit_SetOverwriteMode(this.h, (bool)(overwrite))
}

func (this *QTextEdit) TabStopDistance() float64 {
	return (float64)(QTextEdit_TabStopDistance(this.h))
}

func (this *QTextEdit) SetTabStopDistance(distance float64) {
	QTextEdit_SetTabStopDistance(this.h, (double)(distance))
}

func (this *QTextEdit) CursorWidth() int {
	return (int)(QTextEdit_CursorWidth(this.h))
}

func (this *QTextEdit) SetCursorWidth(width int) {
	QTextEdit_SetCursorWidth(this.h, (int)(width))
}

func (this *QTextEdit) AcceptRichText() bool {
	return (bool)(QTextEdit_AcceptRichText(this.h))
}

func (this *QTextEdit) SetAcceptRichText(accept bool) {
	QTextEdit_SetAcceptRichText(this.h, (bool)(accept))
}

func (this *QTextEdit) SetExtraSelections(selections []ExtraSelection) {
	selections_CArray := (*[0xffff]ExtraSelection)(malloc(size_t(8 * len(selections))))
	defer free(unsafe.Pointer(selections_CArray))
	for i := range selections {
		selections_CArray[i] = selections[i]
	}
	selections_ma := struct_miqt_array{len: size_t(len(selections)), data: unsafe.Pointer(selections_CArray)}
	QTextEdit_SetExtraSelections(this.h, selections_ma)
}

func (this *QTextEdit) ExtraSelections() []ExtraSelection {
	var _ma struct_miqt_array = QTextEdit_ExtraSelections(this.h)
	_ret := make([]ExtraSelection, int(_ma.len))
	_outCast := (*[0xffff]ExtraSelection)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QTextEdit) MoveCursor(operation QTextCursor__MoveOperation) {
	QTextEdit_MoveCursor(this.h, (int)(operation))
}

func (this *QTextEdit) CanPaste() bool {
	return (bool)(QTextEdit_CanPaste(this.h))
}

func (this *QTextEdit) Print(printer *QPagedPaintDevice) {
	QTextEdit_Print(this.h, printer.cPointer())
}

func (this *QTextEdit) InputMethodQuery(property InputMethodQuery) *QVariant {
	_goptr := newQVariant(QTextEdit_InputMethodQuery(this.h, (int)(property)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) InputMethodQuery2(query InputMethodQuery, argument QVariant) *QVariant {
	_goptr := newQVariant(QTextEdit_InputMethodQuery2(this.h, (int)(query), argument.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextEdit) SetFontPointSize(s float64) {
	QTextEdit_SetFontPointSize(this.h, (double)(s))
}

func (this *QTextEdit) SetFontFamily(fontFamily string) {
	fontFamily_ms := struct_miqt_string{}
	fontFamily_ms.data = CString(fontFamily)
	fontFamily_ms.len = size_t(len(fontFamily))
	defer free(unsafe.Pointer(fontFamily_ms.data))
	QTextEdit_SetFontFamily(this.h, fontFamily_ms)
}

func (this *QTextEdit) SetFontWeight(w int) {
	QTextEdit_SetFontWeight(this.h, (int)(w))
}

func (this *QTextEdit) SetFontUnderline(b bool) {
	QTextEdit_SetFontUnderline(this.h, (bool)(b))
}

func (this *QTextEdit) SetFontItalic(b bool) {
	QTextEdit_SetFontItalic(this.h, (bool)(b))
}

func (this *QTextEdit) SetTextColor(c *QColor) {
	QTextEdit_SetTextColor(this.h, c.cPointer())
}

func (this *QTextEdit) SetTextBackgroundColor(c *QColor) {
	QTextEdit_SetTextBackgroundColor(this.h, c.cPointer())
}

func (this *QTextEdit) SetCurrentFont(f *QFont) {
	QTextEdit_SetCurrentFont(this.h, f.cPointer())
}

func (this *QTextEdit) SetAlignment(a AlignmentFlag) {
	QTextEdit_SetAlignment(this.h, (int)(a))
}

func (this *QTextEdit) SetPlainText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextEdit_SetPlainText(this.h, text_ms)
}

func (this *QTextEdit) SetHtml(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextEdit_SetHtml(this.h, text_ms)
}

func (this *QTextEdit) SetMarkdown(markdown string) {
	markdown_ms := struct_miqt_string{}
	markdown_ms.data = CString(markdown)
	markdown_ms.len = size_t(len(markdown))
	defer free(unsafe.Pointer(markdown_ms.data))
	QTextEdit_SetMarkdown(this.h, markdown_ms)
}

func (this *QTextEdit) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextEdit_SetText(this.h, text_ms)
}

func (this *QTextEdit) Cut() {
	QTextEdit_Cut(this.h)
}

func (this *QTextEdit) Copy() {
	QTextEdit_Copy(this.h)
}

func (this *QTextEdit) Paste() {
	QTextEdit_Paste(this.h)
}

func (this *QTextEdit) Undo() {
	QTextEdit_Undo(this.h)
}

func (this *QTextEdit) Redo() {
	QTextEdit_Redo(this.h)
}

func (this *QTextEdit) Clear() {
	QTextEdit_Clear(this.h)
}

func (this *QTextEdit) SelectAll() {
	QTextEdit_SelectAll(this.h)
}

func (this *QTextEdit) InsertPlainText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextEdit_InsertPlainText(this.h, text_ms)
}

func (this *QTextEdit) InsertHtml(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextEdit_InsertHtml(this.h, text_ms)
}

func (this *QTextEdit) Append(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextEdit_Append(this.h, text_ms)
}

func (this *QTextEdit) ScrollToAnchor(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QTextEdit_ScrollToAnchor(this.h, name_ms)
}

func (this *QTextEdit) ZoomIn() {
	QTextEdit_ZoomIn(this.h)
}

func (this *QTextEdit) ZoomOut() {
	QTextEdit_ZoomOut(this.h)
}

func (this *QTextEdit) TextChanged() {
	QTextEdit_TextChanged(this.h)
}

func (this *QTextEdit) OnTextChanged(slot func()) {
	QTextEdit_connect_TextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_TextChanged
func miqt_exec_callback_QTextEdit_TextChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTextEdit) UndoAvailable(b bool) {
	QTextEdit_UndoAvailable(this.h, (bool)(b))
}

func (this *QTextEdit) OnUndoAvailable(slot func(b bool)) {
	QTextEdit_connect_UndoAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_UndoAvailable
func miqt_exec_callback_QTextEdit_UndoAvailable(cb intptr_t, b bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(b bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(b)

	gofunc(slotval1)
}

func (this *QTextEdit) RedoAvailable(b bool) {
	QTextEdit_RedoAvailable(this.h, (bool)(b))
}

func (this *QTextEdit) OnRedoAvailable(slot func(b bool)) {
	QTextEdit_connect_RedoAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_RedoAvailable
func miqt_exec_callback_QTextEdit_RedoAvailable(cb intptr_t, b bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(b bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(b)

	gofunc(slotval1)
}

func (this *QTextEdit) CurrentCharFormatChanged(format *QTextCharFormat) {
	QTextEdit_CurrentCharFormatChanged(this.h, format.cPointer())
}

func (this *QTextEdit) OnCurrentCharFormatChanged(slot func(format *QTextCharFormat)) {
	QTextEdit_connect_CurrentCharFormatChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_CurrentCharFormatChanged
func miqt_exec_callback_QTextEdit_CurrentCharFormatChanged(cb intptr_t, format *QTextCharFormat) {
	gofunc, ok := cgo.Handle(cb).Value().(func(format *QTextCharFormat))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextCharFormat(format)

	gofunc(slotval1)
}

func (this *QTextEdit) CopyAvailable(b bool) {
	QTextEdit_CopyAvailable(this.h, (bool)(b))
}

func (this *QTextEdit) OnCopyAvailable(slot func(b bool)) {
	QTextEdit_connect_CopyAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_CopyAvailable
func miqt_exec_callback_QTextEdit_CopyAvailable(cb intptr_t, b bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(b bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(b)

	gofunc(slotval1)
}

func (this *QTextEdit) SelectionChanged() {
	QTextEdit_SelectionChanged(this.h)
}

func (this *QTextEdit) OnSelectionChanged(slot func()) {
	QTextEdit_connect_SelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_SelectionChanged
func miqt_exec_callback_QTextEdit_SelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTextEdit) CursorPositionChanged() {
	QTextEdit_CursorPositionChanged(this.h)
}

func (this *QTextEdit) OnCursorPositionChanged(slot func()) {
	QTextEdit_connect_CursorPositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_CursorPositionChanged
func miqt_exec_callback_QTextEdit_CursorPositionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QTextEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) Find2(exp string, options FindFlag) bool {
	exp_ms := struct_miqt_string{}
	exp_ms.data = CString(exp)
	exp_ms.len = size_t(len(exp))
	defer free(unsafe.Pointer(exp_ms.data))
	return (bool)(QTextEdit_Find2(this.h, exp_ms, (int)(options)))
}

func (this *QTextEdit) Find22(exp *QRegularExpression, options FindFlag) bool {
	return (bool)(QTextEdit_Find22(this.h, exp.cPointer(), (int)(options)))
}

func (this *QTextEdit) ToMarkdown1(features MarkdownFeature) string {
	var _ms struct_miqt_string = QTextEdit_ToMarkdown1(this.h, (int)(features))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextEdit) MoveCursor2(operation QTextCursor__MoveOperation, mode QTextCursor__MoveMode) {
	QTextEdit_MoveCursor2(this.h, (int)(operation), (int)(mode))
}

func (this *QTextEdit) ZoomIn1(rangeVal int) {
	QTextEdit_ZoomIn1(this.h, (int)(rangeVal))
}

func (this *QTextEdit) ZoomOut1(rangeVal int) {
	QTextEdit_ZoomOut1(this.h, (int)(rangeVal))
}

func (this *QTextEdit) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTextEdit_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTextEdit) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_MetaObject
func miqt_exec_callback_QTextEdit_MetaObject(self QTextEdit, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTextEdit) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTextEdit_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTextEdit) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_Metacast
func miqt_exec_callback_QTextEdit_Metacast(self QTextEdit, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QTextEdit__ExtraSelection struct {
	h          uintptr
	isSubclass bool
}

// NewQTextEdit__ExtraSelection constructs a new QTextEdit::ExtraSelection object.
func NewQTextEdit__ExtraSelection(param1 *ExtraSelection) *QTextEdit__ExtraSelection {
	ret := newQTextEdit__ExtraSelection(QTextEdit__ExtraSelection_new(param1))
	ret.isSubclass = true
	return ret
}

func (this *QTextEdit__ExtraSelection) OperatorAssign(param1 *ExtraSelection) {
	QTextEdit__ExtraSelection_OperatorAssign(this.h, param1)
}
