package qt

import (
	"unsafe"
)

type QTextDocument__MetaInformation int

const (
	QTextDocument__DocumentTitle QTextDocument__MetaInformation = 0
	QTextDocument__DocumentUrl   QTextDocument__MetaInformation = 1
	QTextDocument__CssMedia      QTextDocument__MetaInformation = 2
	QTextDocument__FrontMatter   QTextDocument__MetaInformation = 3
)

type QTextDocument__MarkdownFeature int

const (
	QTextDocument__MarkdownNoHTML            QTextDocument__MarkdownFeature = 96
	QTextDocument__MarkdownDialectCommonMark QTextDocument__MarkdownFeature = 0
	QTextDocument__MarkdownDialectGitHub     QTextDocument__MarkdownFeature = 1068812
)

type QTextDocument__FindFlag int

const (
	QTextDocument__FindBackward        QTextDocument__FindFlag = 1
	QTextDocument__FindCaseSensitively QTextDocument__FindFlag = 2
	QTextDocument__FindWholeWords      QTextDocument__FindFlag = 4
)

type QTextDocument__ResourceType int

const (
	QTextDocument__UnknownResource    QTextDocument__ResourceType = 0
	QTextDocument__HtmlResource       QTextDocument__ResourceType = 1
	QTextDocument__ImageResource      QTextDocument__ResourceType = 2
	QTextDocument__StyleSheetResource QTextDocument__ResourceType = 3
	QTextDocument__MarkdownResource   QTextDocument__ResourceType = 4
	QTextDocument__UserResource       QTextDocument__ResourceType = 100
)

type QTextDocument__Stacks int

const (
	QTextDocument__UndoStack         QTextDocument__Stacks = 1
	QTextDocument__RedoStack         QTextDocument__Stacks = 2
	QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = 3
)

type QAbstractUndoItem struct {
	h          uintptr
	isSubclass bool
}

func (this *QAbstractUndoItem) Undo() {
	QAbstractUndoItem_Undo(this.h)
}

func (this *QAbstractUndoItem) Redo() {
	QAbstractUndoItem_Redo(this.h)
}

func (this *QAbstractUndoItem) OperatorAssign(param1 *QAbstractUndoItem) {
	QAbstractUndoItem_OperatorAssign(this.h, param1.cPointer())
}

type QTextDocument struct {
	h          uintptr
	isSubclass bool
}

// NewQTextDocument constructs a new QTextDocument object.
func NewQTextDocument() *QTextDocument {

	ret := newQTextDocument(QTextDocument_new())
	ret.isSubclass = true
	return ret
}

// NewQTextDocument2 constructs a new QTextDocument object.
func NewQTextDocument2(text string) *QTextDocument {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTextDocument(QTextDocument_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQTextDocument3 constructs a new QTextDocument object.
func NewQTextDocument3(parent *QObject) *QTextDocument {

	ret := newQTextDocument(QTextDocument_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextDocument4 constructs a new QTextDocument object.
func NewQTextDocument4(text string, parent *QObject) *QTextDocument {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTextDocument(QTextDocument_new4(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextDocument) MetaObject() *QMetaObject {
	return newQMetaObject(QTextDocument_MetaObject(this.h))
}

func (this *QTextDocument) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextDocument_Metacast(this.h, param1_Cstring))
}

func QTextDocument_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextDocument_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) Clone() *QTextDocument {
	return newQTextDocument(QTextDocument_Clone(this.h))
}

func (this *QTextDocument) IsEmpty() bool {
	return (bool)(QTextDocument_IsEmpty(this.h))
}

func (this *QTextDocument) Clear() {
	QTextDocument_Clear(this.h)
}

func (this *QTextDocument) SetUndoRedoEnabled(enable bool) {
	QTextDocument_SetUndoRedoEnabled(this.h, (bool)(enable))
}

func (this *QTextDocument) IsUndoRedoEnabled() bool {
	return (bool)(QTextDocument_IsUndoRedoEnabled(this.h))
}

func (this *QTextDocument) IsUndoAvailable() bool {
	return (bool)(QTextDocument_IsUndoAvailable(this.h))
}

func (this *QTextDocument) IsRedoAvailable() bool {
	return (bool)(QTextDocument_IsRedoAvailable(this.h))
}

func (this *QTextDocument) AvailableUndoSteps() int {
	return (int)(QTextDocument_AvailableUndoSteps(this.h))
}

func (this *QTextDocument) AvailableRedoSteps() int {
	return (int)(QTextDocument_AvailableRedoSteps(this.h))
}

func (this *QTextDocument) Revision() int {
	return (int)(QTextDocument_Revision(this.h))
}

func (this *QTextDocument) SetDocumentLayout(layout *QAbstractTextDocumentLayout) {
	QTextDocument_SetDocumentLayout(this.h, layout.cPointer())
}

func (this *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout {
	return newQAbstractTextDocumentLayout(QTextDocument_DocumentLayout(this.h))
}

func (this *QTextDocument) SetMetaInformation(info MetaInformation, param2 string) {
	param2_ms := struct_miqt_string{}
	param2_ms.data = CString(param2)
	param2_ms.len = size_t(len(param2))
	defer free(unsafe.Pointer(param2_ms.data))
	QTextDocument_SetMetaInformation(this.h, info, param2_ms)
}

func (this *QTextDocument) MetaInformation(info MetaInformation) string {
	var _ms struct_miqt_string = QTextDocument_MetaInformation(this.h, info)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) ToHtml() string {
	var _ms struct_miqt_string = QTextDocument_ToHtml(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) SetHtml(html string) {
	html_ms := struct_miqt_string{}
	html_ms.data = CString(html)
	html_ms.len = size_t(len(html))
	defer free(unsafe.Pointer(html_ms.data))
	QTextDocument_SetHtml(this.h, html_ms)
}

func (this *QTextDocument) ToMarkdown() string {
	var _ms struct_miqt_string = QTextDocument_ToMarkdown(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) SetMarkdown(markdown string) {
	markdown_ms := struct_miqt_string{}
	markdown_ms.data = CString(markdown)
	markdown_ms.len = size_t(len(markdown))
	defer free(unsafe.Pointer(markdown_ms.data))
	QTextDocument_SetMarkdown(this.h, markdown_ms)
}

func (this *QTextDocument) ToRawText() string {
	var _ms struct_miqt_string = QTextDocument_ToRawText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) ToPlainText() string {
	var _ms struct_miqt_string = QTextDocument_ToPlainText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) SetPlainText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextDocument_SetPlainText(this.h, text_ms)
}

func (this *QTextDocument) CharacterAt(pos int) *QChar {
	_goptr := newQChar(QTextDocument_CharacterAt(this.h, (int)(pos)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find(subString string) *QTextCursor {
	subString_ms := struct_miqt_string{}
	subString_ms.data = CString(subString)
	subString_ms.len = size_t(len(subString))
	defer free(unsafe.Pointer(subString_ms.data))
	_goptr := newQTextCursor(QTextDocument_Find(this.h, subString_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find2(subString string, cursor *QTextCursor) *QTextCursor {
	subString_ms := struct_miqt_string{}
	subString_ms.data = CString(subString)
	subString_ms.len = size_t(len(subString))
	defer free(unsafe.Pointer(subString_ms.data))
	_goptr := newQTextCursor(QTextDocument_Find2(this.h, subString_ms, cursor.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) FindWithExpr(expr *QRegularExpression) *QTextCursor {
	_goptr := newQTextCursor(QTextDocument_FindWithExpr(this.h, expr.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find3(expr *QRegularExpression, cursor *QTextCursor) *QTextCursor {
	_goptr := newQTextCursor(QTextDocument_Find3(this.h, expr.cPointer(), cursor.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) FrameAt(pos int) *QTextFrame {
	return newQTextFrame(QTextDocument_FrameAt(this.h, (int)(pos)))
}

func (this *QTextDocument) RootFrame() *QTextFrame {
	return newQTextFrame(QTextDocument_RootFrame(this.h))
}

func (this *QTextDocument) Object(objectIndex int) *QTextObject {
	return newQTextObject(QTextDocument_Object(this.h, (int)(objectIndex)))
}

func (this *QTextDocument) ObjectForFormat(param1 *QTextFormat) *QTextObject {
	return newQTextObject(QTextDocument_ObjectForFormat(this.h, param1.cPointer()))
}

func (this *QTextDocument) FindBlock(pos int) *QTextBlock {
	_goptr := newQTextBlock(QTextDocument_FindBlock(this.h, (int)(pos)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) FindBlockByNumber(blockNumber int) *QTextBlock {
	_goptr := newQTextBlock(QTextDocument_FindBlockByNumber(this.h, (int)(blockNumber)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) FindBlockByLineNumber(blockNumber int) *QTextBlock {
	_goptr := newQTextBlock(QTextDocument_FindBlockByLineNumber(this.h, (int)(blockNumber)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Begin() *QTextBlock {
	_goptr := newQTextBlock(QTextDocument_Begin(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) End() *QTextBlock {
	_goptr := newQTextBlock(QTextDocument_End(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) FirstBlock() *QTextBlock {
	_goptr := newQTextBlock(QTextDocument_FirstBlock(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) LastBlock() *QTextBlock {
	_goptr := newQTextBlock(QTextDocument_LastBlock(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) SetPageSize(size *QSizeF) {
	QTextDocument_SetPageSize(this.h, size.cPointer())
}

func (this *QTextDocument) PageSize() *QSizeF {
	_goptr := newQSizeF(QTextDocument_PageSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) SetDefaultFont(font *QFont) {
	QTextDocument_SetDefaultFont(this.h, font.cPointer())
}

func (this *QTextDocument) DefaultFont() *QFont {
	_goptr := newQFont(QTextDocument_DefaultFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) SetSuperScriptBaseline(baseline float64) {
	QTextDocument_SetSuperScriptBaseline(this.h, (double)(baseline))
}

func (this *QTextDocument) SuperScriptBaseline() float64 {
	return (float64)(QTextDocument_SuperScriptBaseline(this.h))
}

func (this *QTextDocument) SetSubScriptBaseline(baseline float64) {
	QTextDocument_SetSubScriptBaseline(this.h, (double)(baseline))
}

func (this *QTextDocument) SubScriptBaseline() float64 {
	return (float64)(QTextDocument_SubScriptBaseline(this.h))
}

func (this *QTextDocument) SetBaselineOffset(baseline float64) {
	QTextDocument_SetBaselineOffset(this.h, (double)(baseline))
}

func (this *QTextDocument) BaselineOffset() float64 {
	return (float64)(QTextDocument_BaselineOffset(this.h))
}

func (this *QTextDocument) PageCount() int {
	return (int)(QTextDocument_PageCount(this.h))
}

func (this *QTextDocument) IsModified() bool {
	return (bool)(QTextDocument_IsModified(this.h))
}

func (this *QTextDocument) Print(printer *QPagedPaintDevice) {
	QTextDocument_Print(this.h, printer.cPointer())
}

func (this *QTextDocument) Resource(typeVal int, name *QUrl) *QVariant {
	_goptr := newQVariant(QTextDocument_Resource(this.h, (int)(typeVal), name.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) AddResource(typeVal int, name *QUrl, resource *QVariant) {
	QTextDocument_AddResource(this.h, (int)(typeVal), name.cPointer(), resource.cPointer())
}

func (this *QTextDocument) SetResourceProvider(provider *ResourceProvider) {
	QTextDocument_SetResourceProvider(this.h, provider)
}

func QTextDocument_SetDefaultResourceProvider(provider *ResourceProvider) {
	QTextDocument_SetDefaultResourceProvider(provider)
}

func (this *QTextDocument) AllFormats() []QTextFormat {
	var _ma struct_miqt_array = QTextDocument_AllFormats(this.h)
	_ret := make([]QTextFormat, int(_ma.len))
	_outCast := (*[0xffff]*QTextFormat)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQTextFormat(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextDocument) MarkContentsDirty(from int, length int) {
	QTextDocument_MarkContentsDirty(this.h, (int)(from), (int)(length))
}

func (this *QTextDocument) SetUseDesignMetrics(b bool) {
	QTextDocument_SetUseDesignMetrics(this.h, (bool)(b))
}

func (this *QTextDocument) UseDesignMetrics() bool {
	return (bool)(QTextDocument_UseDesignMetrics(this.h))
}

func (this *QTextDocument) SetLayoutEnabled(b bool) {
	QTextDocument_SetLayoutEnabled(this.h, (bool)(b))
}

func (this *QTextDocument) IsLayoutEnabled() bool {
	return (bool)(QTextDocument_IsLayoutEnabled(this.h))
}

func (this *QTextDocument) DrawContents(painter *QPainter) {
	QTextDocument_DrawContents(this.h, painter.cPointer())
}

func (this *QTextDocument) SetTextWidth(width float64) {
	QTextDocument_SetTextWidth(this.h, (double)(width))
}

func (this *QTextDocument) TextWidth() float64 {
	return (float64)(QTextDocument_TextWidth(this.h))
}

func (this *QTextDocument) IdealWidth() float64 {
	return (float64)(QTextDocument_IdealWidth(this.h))
}

func (this *QTextDocument) IndentWidth() float64 {
	return (float64)(QTextDocument_IndentWidth(this.h))
}

func (this *QTextDocument) SetIndentWidth(width float64) {
	QTextDocument_SetIndentWidth(this.h, (double)(width))
}

func (this *QTextDocument) DocumentMargin() float64 {
	return (float64)(QTextDocument_DocumentMargin(this.h))
}

func (this *QTextDocument) SetDocumentMargin(margin float64) {
	QTextDocument_SetDocumentMargin(this.h, (double)(margin))
}

func (this *QTextDocument) AdjustSize() {
	QTextDocument_AdjustSize(this.h)
}

func (this *QTextDocument) Size() *QSizeF {
	_goptr := newQSizeF(QTextDocument_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) BlockCount() int {
	return (int)(QTextDocument_BlockCount(this.h))
}

func (this *QTextDocument) LineCount() int {
	return (int)(QTextDocument_LineCount(this.h))
}

func (this *QTextDocument) CharacterCount() int {
	return (int)(QTextDocument_CharacterCount(this.h))
}

func (this *QTextDocument) SetDefaultStyleSheet(sheet string) {
	sheet_ms := struct_miqt_string{}
	sheet_ms.data = CString(sheet)
	sheet_ms.len = size_t(len(sheet))
	defer free(unsafe.Pointer(sheet_ms.data))
	QTextDocument_SetDefaultStyleSheet(this.h, sheet_ms)
}

func (this *QTextDocument) DefaultStyleSheet() string {
	var _ms struct_miqt_string = QTextDocument_DefaultStyleSheet(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) Undo(cursor *QTextCursor) {
	QTextDocument_Undo(this.h, cursor.cPointer())
}

func (this *QTextDocument) Redo(cursor *QTextCursor) {
	QTextDocument_Redo(this.h, cursor.cPointer())
}

func (this *QTextDocument) ClearUndoRedoStacks() {
	QTextDocument_ClearUndoRedoStacks(this.h)
}

func (this *QTextDocument) MaximumBlockCount() int {
	return (int)(QTextDocument_MaximumBlockCount(this.h))
}

func (this *QTextDocument) SetMaximumBlockCount(maximum int) {
	QTextDocument_SetMaximumBlockCount(this.h, (int)(maximum))
}

func (this *QTextDocument) DefaultTextOption() *QTextOption {
	_goptr := newQTextOption(QTextDocument_DefaultTextOption(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) SetDefaultTextOption(option *QTextOption) {
	QTextDocument_SetDefaultTextOption(this.h, option.cPointer())
}

func (this *QTextDocument) BaseUrl() *QUrl {
	_goptr := newQUrl(QTextDocument_BaseUrl(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) SetBaseUrl(url *QUrl) {
	QTextDocument_SetBaseUrl(this.h, url.cPointer())
}

func (this *QTextDocument) DefaultCursorMoveStyle() CursorMoveStyle {
	return (CursorMoveStyle)(QTextDocument_DefaultCursorMoveStyle(this.h))
}

func (this *QTextDocument) SetDefaultCursorMoveStyle(style CursorMoveStyle) {
	QTextDocument_SetDefaultCursorMoveStyle(this.h, (int)(style))
}

func (this *QTextDocument) ContentsChange(from int, charsRemoved int, charsAdded int) {
	QTextDocument_ContentsChange(this.h, (int)(from), (int)(charsRemoved), (int)(charsAdded))
}
func (this *QTextDocument) OnContentsChange(slot func(from int, charsRemoved int, charsAdded int)) {
	QTextDocument_connect_ContentsChange(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_ContentsChange
func miqt_exec_callback_QTextDocument_ContentsChange(cb intptr_t, from int, charsRemoved int, charsAdded int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(from int, charsRemoved int, charsAdded int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(from)

	slotval2 := (int)(charsRemoved)

	slotval3 := (int)(charsAdded)

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QTextDocument) ContentsChanged() {
	QTextDocument_ContentsChanged(this.h)
}
func (this *QTextDocument) OnContentsChanged(slot func()) {
	QTextDocument_connect_ContentsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_ContentsChanged
func miqt_exec_callback_QTextDocument_ContentsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTextDocument) UndoAvailable(param1 bool) {
	QTextDocument_UndoAvailable(this.h, (bool)(param1))
}
func (this *QTextDocument) OnUndoAvailable(slot func(param1 bool)) {
	QTextDocument_connect_UndoAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_UndoAvailable
func miqt_exec_callback_QTextDocument_UndoAvailable(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QTextDocument) RedoAvailable(param1 bool) {
	QTextDocument_RedoAvailable(this.h, (bool)(param1))
}
func (this *QTextDocument) OnRedoAvailable(slot func(param1 bool)) {
	QTextDocument_connect_RedoAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_RedoAvailable
func miqt_exec_callback_QTextDocument_RedoAvailable(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QTextDocument) UndoCommandAdded() {
	QTextDocument_UndoCommandAdded(this.h)
}
func (this *QTextDocument) OnUndoCommandAdded(slot func()) {
	QTextDocument_connect_UndoCommandAdded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_UndoCommandAdded
func miqt_exec_callback_QTextDocument_UndoCommandAdded(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTextDocument) ModificationChanged(m bool) {
	QTextDocument_ModificationChanged(this.h, (bool)(m))
}
func (this *QTextDocument) OnModificationChanged(slot func(m bool)) {
	QTextDocument_connect_ModificationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_ModificationChanged
func miqt_exec_callback_QTextDocument_ModificationChanged(cb intptr_t, m bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(m bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(m)

	gofunc(slotval1)
}

func (this *QTextDocument) CursorPositionChanged(cursor *QTextCursor) {
	QTextDocument_CursorPositionChanged(this.h, cursor.cPointer())
}
func (this *QTextDocument) OnCursorPositionChanged(slot func(cursor *QTextCursor)) {
	QTextDocument_connect_CursorPositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_CursorPositionChanged
func miqt_exec_callback_QTextDocument_CursorPositionChanged(cb intptr_t, cursor *QTextCursor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(cursor *QTextCursor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextCursor(cursor)

	gofunc(slotval1)
}

func (this *QTextDocument) BlockCountChanged(newBlockCount int) {
	QTextDocument_BlockCountChanged(this.h, (int)(newBlockCount))
}
func (this *QTextDocument) OnBlockCountChanged(slot func(newBlockCount int)) {
	QTextDocument_connect_BlockCountChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_BlockCountChanged
func miqt_exec_callback_QTextDocument_BlockCountChanged(cb intptr_t, newBlockCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newBlockCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(newBlockCount)

	gofunc(slotval1)
}

func (this *QTextDocument) BaseUrlChanged(url *QUrl) {
	QTextDocument_BaseUrlChanged(this.h, url.cPointer())
}
func (this *QTextDocument) OnBaseUrlChanged(slot func(url *QUrl)) {
	QTextDocument_connect_BaseUrlChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_BaseUrlChanged
func miqt_exec_callback_QTextDocument_BaseUrlChanged(cb intptr_t, url *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(url *QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(url)

	gofunc(slotval1)
}

func (this *QTextDocument) DocumentLayoutChanged() {
	QTextDocument_DocumentLayoutChanged(this.h)
}
func (this *QTextDocument) OnDocumentLayoutChanged(slot func()) {
	QTextDocument_connect_DocumentLayoutChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_DocumentLayoutChanged
func miqt_exec_callback_QTextDocument_DocumentLayoutChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTextDocument) Undo2() {
	QTextDocument_Undo2(this.h)
}

func (this *QTextDocument) Redo2() {
	QTextDocument_Redo2(this.h)
}

func (this *QTextDocument) AppendUndoItem(param1 *QAbstractUndoItem) {
	QTextDocument_AppendUndoItem(this.h, param1.cPointer())
}

func (this *QTextDocument) SetModified() {
	QTextDocument_SetModified(this.h)
}

func QTextDocument_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextDocument_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextDocument_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextDocument_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) Clone1(parent *QObject) *QTextDocument {
	return newQTextDocument(QTextDocument_Clone1(this.h, parent.cPointer()))
}

func (this *QTextDocument) ToMarkdown1(features MarkdownFeatures) string {
	var _ms struct_miqt_string = QTextDocument_ToMarkdown1(this.h, features)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocument) SetMarkdown2(markdown string, features MarkdownFeatures) {
	markdown_ms := struct_miqt_string{}
	markdown_ms.data = CString(markdown)
	markdown_ms.len = size_t(len(markdown))
	defer free(unsafe.Pointer(markdown_ms.data))
	QTextDocument_SetMarkdown2(this.h, markdown_ms, features)
}

func (this *QTextDocument) Find22(subString string, from int) *QTextCursor {
	subString_ms := struct_miqt_string{}
	subString_ms.data = CString(subString)
	subString_ms.len = size_t(len(subString))
	defer free(unsafe.Pointer(subString_ms.data))
	_goptr := newQTextCursor(QTextDocument_Find22(this.h, subString_ms, (int)(from)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find32(subString string, from int, options FindFlags) *QTextCursor {
	subString_ms := struct_miqt_string{}
	subString_ms.data = CString(subString)
	subString_ms.len = size_t(len(subString))
	defer free(unsafe.Pointer(subString_ms.data))
	_goptr := newQTextCursor(QTextDocument_Find32(this.h, subString_ms, (int)(from), options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find33(subString string, cursor *QTextCursor, options FindFlags) *QTextCursor {
	subString_ms := struct_miqt_string{}
	subString_ms.data = CString(subString)
	subString_ms.len = size_t(len(subString))
	defer free(unsafe.Pointer(subString_ms.data))
	_goptr := newQTextCursor(QTextDocument_Find33(this.h, subString_ms, cursor.cPointer(), options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find23(expr *QRegularExpression, from int) *QTextCursor {
	_goptr := newQTextCursor(QTextDocument_Find23(this.h, expr.cPointer(), (int)(from)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find34(expr *QRegularExpression, from int, options FindFlags) *QTextCursor {
	_goptr := newQTextCursor(QTextDocument_Find34(this.h, expr.cPointer(), (int)(from), options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) Find35(expr *QRegularExpression, cursor *QTextCursor, options FindFlags) *QTextCursor {
	_goptr := newQTextCursor(QTextDocument_Find35(this.h, expr.cPointer(), cursor.cPointer(), options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocument) DrawContents2(painter *QPainter, rect *QRectF) {
	QTextDocument_DrawContents2(this.h, painter.cPointer(), rect.cPointer())
}

func (this *QTextDocument) ClearUndoRedoStacks1(historyToClear Stacks) {
	QTextDocument_ClearUndoRedoStacks1(this.h, historyToClear)
}

func (this *QTextDocument) SetModified1(m bool) {
	QTextDocument_SetModified1(this.h, (bool)(m))
}

func (this *QTextDocument) callVirtualBase_Clear() {

	QTextDocument_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QTextDocument) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_Clear
func miqt_exec_callback_QTextDocument_Clear(self QTextDocument, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTextDocument{h: self}).callVirtualBase_Clear)

}

func (this *QTextDocument) callVirtualBase_CreateObject(f *QTextFormat) *QTextObject {

	return newQTextObject(QTextDocument_virtualbase_CreateObject(unsafe.Pointer(this.h), f.cPointer()))

}
func (this *QTextDocument) OnCreateObject(slot func(super func(f *QTextFormat) *QTextObject, f *QTextFormat) *QTextObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_CreateObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_CreateObject
func miqt_exec_callback_QTextDocument_CreateObject(self QTextDocument, cb intptr_t, f *QTextFormat) *QTextObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(f *QTextFormat) *QTextObject, f *QTextFormat) *QTextObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextFormat(f)

	virtualReturn := gofunc((&QTextDocument{h: self}).callVirtualBase_CreateObject, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTextDocument) callVirtualBase_LoadResource(typeVal int, name *QUrl) *QVariant {

	_goptr := newQVariant(QTextDocument_virtualbase_LoadResource(unsafe.Pointer(this.h), (int)(typeVal), name.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextDocument) OnLoadResource(slot func(super func(typeVal int, name *QUrl) *QVariant, typeVal int, name *QUrl) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_LoadResource(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_LoadResource
func miqt_exec_callback_QTextDocument_LoadResource(self QTextDocument, cb intptr_t, typeVal int, name *QUrl) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal int, name *QUrl) *QVariant, typeVal int, name *QUrl) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(typeVal)

	slotval2 := newQUrl(name)

	virtualReturn := gofunc((&QTextDocument{h: self}).callVirtualBase_LoadResource, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QTextDocument) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QTextDocument_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTextDocument) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_Event
func miqt_exec_callback_QTextDocument_Event(self QTextDocument, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTextDocument{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextDocument) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QTextDocument_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QTextDocument) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_EventFilter
func miqt_exec_callback_QTextDocument_EventFilter(self QTextDocument, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTextDocument{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QTextDocument) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QTextDocument_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTextDocument) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_TimerEvent
func miqt_exec_callback_QTextDocument_TimerEvent(self QTextDocument, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QTextDocument{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTextDocument) callVirtualBase_ChildEvent(event *QChildEvent) {

	QTextDocument_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTextDocument) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_ChildEvent
func miqt_exec_callback_QTextDocument_ChildEvent(self QTextDocument, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QTextDocument{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QTextDocument) callVirtualBase_CustomEvent(event *QEvent) {

	QTextDocument_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTextDocument) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_CustomEvent
func miqt_exec_callback_QTextDocument_CustomEvent(self QTextDocument, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTextDocument{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QTextDocument) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QTextDocument_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTextDocument) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_ConnectNotify
func miqt_exec_callback_QTextDocument_ConnectNotify(self QTextDocument, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTextDocument{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QTextDocument) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QTextDocument_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTextDocument) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextDocument_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextDocument_DisconnectNotify
func miqt_exec_callback_QTextDocument_DisconnectNotify(self QTextDocument, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTextDocument{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
