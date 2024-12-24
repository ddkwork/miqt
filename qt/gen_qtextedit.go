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

func (this *QTextEdit) callVirtualBase_LoadResource(typeVal int, name *QUrl) *QVariant {

	_goptr := newQVariant(QTextEdit_virtualbase_LoadResource(unsafe.Pointer(this.h), (int)(typeVal), name.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextEdit) OnLoadResource(slot func(super func(typeVal int, name *QUrl) *QVariant, typeVal int, name *QUrl) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_LoadResource(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_LoadResource
func miqt_exec_callback_QTextEdit_LoadResource(self QTextEdit, cb intptr_t, typeVal int, name *QUrl) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal int, name *QUrl) *QVariant, typeVal int, name *QUrl) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(typeVal)

	slotval2 := newQUrl(name)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_LoadResource, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QTextEdit) callVirtualBase_InputMethodQuery(property InputMethodQuery) *QVariant {

	_goptr := newQVariant(QTextEdit_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(property)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextEdit) OnInputMethodQuery(slot func(super func(property InputMethodQuery) *QVariant, property InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_InputMethodQuery
func miqt_exec_callback_QTextEdit_InputMethodQuery(self QTextEdit, cb intptr_t, property int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(property InputMethodQuery) *QVariant, property InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(property)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTextEdit) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QTextEdit_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QTextEdit) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_Event
func miqt_exec_callback_QTextEdit_Event(self QTextEdit, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextEdit) callVirtualBase_TimerEvent(e *QTimerEvent) {

	QTextEdit_virtualbase_TimerEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnTimerEvent(slot func(super func(e *QTimerEvent), e *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_TimerEvent
func miqt_exec_callback_QTextEdit_TimerEvent(self QTextEdit, cb intptr_t, e *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QTimerEvent), e *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_KeyPressEvent(e *QKeyEvent) {

	QTextEdit_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnKeyPressEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_KeyPressEvent
func miqt_exec_callback_QTextEdit_KeyPressEvent(self QTextEdit, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_KeyReleaseEvent(e *QKeyEvent) {

	QTextEdit_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnKeyReleaseEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_KeyReleaseEvent
func miqt_exec_callback_QTextEdit_KeyReleaseEvent(self QTextEdit, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_ResizeEvent(e *QResizeEvent) {

	QTextEdit_virtualbase_ResizeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnResizeEvent(slot func(super func(e *QResizeEvent), e *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_ResizeEvent
func miqt_exec_callback_QTextEdit_ResizeEvent(self QTextEdit, cb intptr_t, e *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QResizeEvent), e *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QTextEdit_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_PaintEvent
func miqt_exec_callback_QTextEdit_PaintEvent(self QTextEdit, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_MousePressEvent(e *QMouseEvent) {

	QTextEdit_virtualbase_MousePressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnMousePressEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_MousePressEvent
func miqt_exec_callback_QTextEdit_MousePressEvent(self QTextEdit, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_MouseMoveEvent(e *QMouseEvent) {

	QTextEdit_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnMouseMoveEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_MouseMoveEvent
func miqt_exec_callback_QTextEdit_MouseMoveEvent(self QTextEdit, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QTextEdit_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_MouseReleaseEvent
func miqt_exec_callback_QTextEdit_MouseReleaseEvent(self QTextEdit, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_MouseDoubleClickEvent(e *QMouseEvent) {

	QTextEdit_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnMouseDoubleClickEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_MouseDoubleClickEvent
func miqt_exec_callback_QTextEdit_MouseDoubleClickEvent(self QTextEdit, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QTextEdit_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QTextEdit) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_FocusNextPrevChild
func miqt_exec_callback_QTextEdit_FocusNextPrevChild(self QTextEdit, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextEdit) callVirtualBase_ContextMenuEvent(e *QContextMenuEvent) {

	QTextEdit_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnContextMenuEvent(slot func(super func(e *QContextMenuEvent), e *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_ContextMenuEvent
func miqt_exec_callback_QTextEdit_ContextMenuEvent(self QTextEdit, cb intptr_t, e *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QContextMenuEvent), e *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_DragEnterEvent(e *QDragEnterEvent) {

	QTextEdit_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnDragEnterEvent(slot func(super func(e *QDragEnterEvent), e *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_DragEnterEvent
func miqt_exec_callback_QTextEdit_DragEnterEvent(self QTextEdit, cb intptr_t, e *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragEnterEvent), e *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_DragLeaveEvent(e *QDragLeaveEvent) {

	QTextEdit_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnDragLeaveEvent(slot func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_DragLeaveEvent
func miqt_exec_callback_QTextEdit_DragLeaveEvent(self QTextEdit, cb intptr_t, e *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_DragMoveEvent(e *QDragMoveEvent) {

	QTextEdit_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnDragMoveEvent(slot func(super func(e *QDragMoveEvent), e *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_DragMoveEvent
func miqt_exec_callback_QTextEdit_DragMoveEvent(self QTextEdit, cb intptr_t, e *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragMoveEvent), e *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_DropEvent(e *QDropEvent) {

	QTextEdit_virtualbase_DropEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnDropEvent(slot func(super func(e *QDropEvent), e *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_DropEvent
func miqt_exec_callback_QTextEdit_DropEvent(self QTextEdit, cb intptr_t, e *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDropEvent), e *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_FocusInEvent(e *QFocusEvent) {

	QTextEdit_virtualbase_FocusInEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnFocusInEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_FocusInEvent
func miqt_exec_callback_QTextEdit_FocusInEvent(self QTextEdit, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_FocusOutEvent(e *QFocusEvent) {

	QTextEdit_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnFocusOutEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_FocusOutEvent
func miqt_exec_callback_QTextEdit_FocusOutEvent(self QTextEdit, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QTextEdit_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTextEdit) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_ShowEvent
func miqt_exec_callback_QTextEdit_ShowEvent(self QTextEdit, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QTextEdit{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_ChangeEvent(e *QEvent) {

	QTextEdit_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_ChangeEvent
func miqt_exec_callback_QTextEdit_ChangeEvent(self QTextEdit, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QTextEdit_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextEdit) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_WheelEvent
func miqt_exec_callback_QTextEdit_WheelEvent(self QTextEdit, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QTextEdit{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_CreateMimeDataFromSelection() *QMimeData {

	return newQMimeData(QTextEdit_virtualbase_CreateMimeDataFromSelection(unsafe.Pointer(this.h)))

}
func (this *QTextEdit) OnCreateMimeDataFromSelection(slot func(super func() *QMimeData) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_CreateMimeDataFromSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_CreateMimeDataFromSelection
func miqt_exec_callback_QTextEdit_CreateMimeDataFromSelection(self QTextEdit, cb intptr_t) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMimeData) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_CreateMimeDataFromSelection)

	return virtualReturn.cPointer()

}

func (this *QTextEdit) callVirtualBase_CanInsertFromMimeData(source *QMimeData) bool {

	return (bool)(QTextEdit_virtualbase_CanInsertFromMimeData(unsafe.Pointer(this.h), source.cPointer()))

}
func (this *QTextEdit) OnCanInsertFromMimeData(slot func(super func(source *QMimeData) bool, source *QMimeData) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_CanInsertFromMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_CanInsertFromMimeData
func miqt_exec_callback_QTextEdit_CanInsertFromMimeData(self QTextEdit, cb intptr_t, source *QMimeData) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(source *QMimeData) bool, source *QMimeData) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMimeData(source)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_CanInsertFromMimeData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextEdit) callVirtualBase_InsertFromMimeData(source *QMimeData) {

	QTextEdit_virtualbase_InsertFromMimeData(unsafe.Pointer(this.h), source.cPointer())

}
func (this *QTextEdit) OnInsertFromMimeData(slot func(super func(source *QMimeData), source *QMimeData)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_InsertFromMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_InsertFromMimeData
func miqt_exec_callback_QTextEdit_InsertFromMimeData(self QTextEdit, cb intptr_t, source *QMimeData) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(source *QMimeData), source *QMimeData))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMimeData(source)

	gofunc((&QTextEdit{h: self}).callVirtualBase_InsertFromMimeData, slotval1)

}

func (this *QTextEdit) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QTextEdit_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTextEdit) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_InputMethodEvent
func miqt_exec_callback_QTextEdit_InputMethodEvent(self QTextEdit, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QTextEdit{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QTextEdit) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QTextEdit_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QTextEdit) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_ScrollContentsBy
func miqt_exec_callback_QTextEdit_ScrollContentsBy(self QTextEdit, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QTextEdit{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QTextEdit) callVirtualBase_DoSetTextCursor(cursor *QTextCursor) {

	QTextEdit_virtualbase_DoSetTextCursor(unsafe.Pointer(this.h), cursor.cPointer())

}
func (this *QTextEdit) OnDoSetTextCursor(slot func(super func(cursor *QTextCursor), cursor *QTextCursor)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_DoSetTextCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_DoSetTextCursor
func miqt_exec_callback_QTextEdit_DoSetTextCursor(self QTextEdit, cb intptr_t, cursor *QTextCursor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursor *QTextCursor), cursor *QTextCursor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextCursor(cursor)

	gofunc((&QTextEdit{h: self}).callVirtualBase_DoSetTextCursor, slotval1)

}

func (this *QTextEdit) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QTextEdit_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextEdit) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_MinimumSizeHint
func miqt_exec_callback_QTextEdit_MinimumSizeHint(self QTextEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QTextEdit) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QTextEdit_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextEdit) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_SizeHint
func miqt_exec_callback_QTextEdit_SizeHint(self QTextEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QTextEdit) callVirtualBase_SetupViewport(viewport *QWidget) {

	QTextEdit_virtualbase_SetupViewport(unsafe.Pointer(this.h), viewport.cPointer())

}
func (this *QTextEdit) OnSetupViewport(slot func(super func(viewport *QWidget), viewport *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_SetupViewport(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_SetupViewport
func miqt_exec_callback_QTextEdit_SetupViewport(self QTextEdit, cb intptr_t, viewport *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(viewport *QWidget), viewport *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(viewport)

	gofunc((&QTextEdit{h: self}).callVirtualBase_SetupViewport, slotval1)

}

func (this *QTextEdit) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QTextEdit_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QTextEdit) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_EventFilter
func miqt_exec_callback_QTextEdit_EventFilter(self QTextEdit, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QTextEdit) callVirtualBase_ViewportEvent(param1 *QEvent) bool {

	return (bool)(QTextEdit_virtualbase_ViewportEvent(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QTextEdit) OnViewportEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_ViewportEvent
func miqt_exec_callback_QTextEdit_ViewportEvent(self QTextEdit, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextEdit) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QTextEdit_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextEdit) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextEdit_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextEdit_ViewportSizeHint
func miqt_exec_callback_QTextEdit_ViewportSizeHint(self QTextEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextEdit{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

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
