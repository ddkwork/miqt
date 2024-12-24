package qt

import (
	"unsafe"
)

type QTextCursor__MoveMode int

const (
	QTextCursor__MoveAnchor QTextCursor__MoveMode = 0
	QTextCursor__KeepAnchor QTextCursor__MoveMode = 1
)

type QTextCursor__MoveOperation int

const (
	QTextCursor__NoMove            QTextCursor__MoveOperation = 0
	QTextCursor__Start             QTextCursor__MoveOperation = 1
	QTextCursor__Up                QTextCursor__MoveOperation = 2
	QTextCursor__StartOfLine       QTextCursor__MoveOperation = 3
	QTextCursor__StartOfBlock      QTextCursor__MoveOperation = 4
	QTextCursor__StartOfWord       QTextCursor__MoveOperation = 5
	QTextCursor__PreviousBlock     QTextCursor__MoveOperation = 6
	QTextCursor__PreviousCharacter QTextCursor__MoveOperation = 7
	QTextCursor__PreviousWord      QTextCursor__MoveOperation = 8
	QTextCursor__Left              QTextCursor__MoveOperation = 9
	QTextCursor__WordLeft          QTextCursor__MoveOperation = 10
	QTextCursor__End               QTextCursor__MoveOperation = 11
	QTextCursor__Down              QTextCursor__MoveOperation = 12
	QTextCursor__EndOfLine         QTextCursor__MoveOperation = 13
	QTextCursor__EndOfWord         QTextCursor__MoveOperation = 14
	QTextCursor__EndOfBlock        QTextCursor__MoveOperation = 15
	QTextCursor__NextBlock         QTextCursor__MoveOperation = 16
	QTextCursor__NextCharacter     QTextCursor__MoveOperation = 17
	QTextCursor__NextWord          QTextCursor__MoveOperation = 18
	QTextCursor__Right             QTextCursor__MoveOperation = 19
	QTextCursor__WordRight         QTextCursor__MoveOperation = 20
	QTextCursor__NextCell          QTextCursor__MoveOperation = 21
	QTextCursor__PreviousCell      QTextCursor__MoveOperation = 22
	QTextCursor__NextRow           QTextCursor__MoveOperation = 23
	QTextCursor__PreviousRow       QTextCursor__MoveOperation = 24
)

type QTextCursor__SelectionType int

const (
	QTextCursor__WordUnderCursor  QTextCursor__SelectionType = 0
	QTextCursor__LineUnderCursor  QTextCursor__SelectionType = 1
	QTextCursor__BlockUnderCursor QTextCursor__SelectionType = 2
	QTextCursor__Document         QTextCursor__SelectionType = 3
)

type QTextCursor struct {
	h          uintptr
	isSubclass bool
}

// NewQTextCursor constructs a new QTextCursor object.
func NewQTextCursor() *QTextCursor {

	ret := newQTextCursor(QTextCursor_new())
	ret.isSubclass = true
	return ret
}

// NewQTextCursor2 constructs a new QTextCursor object.
func NewQTextCursor2(document *QTextDocument) *QTextCursor {

	ret := newQTextCursor(QTextCursor_new2(document.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextCursor3 constructs a new QTextCursor object.
func NewQTextCursor3(frame *QTextFrame) *QTextCursor {

	ret := newQTextCursor(QTextCursor_new3(frame.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextCursor4 constructs a new QTextCursor object.
func NewQTextCursor4(block *QTextBlock) *QTextCursor {

	ret := newQTextCursor(QTextCursor_new4(block.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextCursor5 constructs a new QTextCursor object.
func NewQTextCursor5(cursor *QTextCursor) *QTextCursor {

	ret := newQTextCursor(QTextCursor_new5(cursor.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextCursor) OperatorAssign(other *QTextCursor) {
	QTextCursor_OperatorAssign(this.h, other.cPointer())
}

func (this *QTextCursor) Swap(other *QTextCursor) {
	QTextCursor_Swap(this.h, other.cPointer())
}

func (this *QTextCursor) IsNull() bool {
	return (bool)(QTextCursor_IsNull(this.h))
}

func (this *QTextCursor) SetPosition(pos int) {
	QTextCursor_SetPosition(this.h, (int)(pos))
}

func (this *QTextCursor) Position() int {
	return (int)(QTextCursor_Position(this.h))
}

func (this *QTextCursor) PositionInBlock() int {
	return (int)(QTextCursor_PositionInBlock(this.h))
}

func (this *QTextCursor) Anchor() int {
	return (int)(QTextCursor_Anchor(this.h))
}

func (this *QTextCursor) InsertText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextCursor_InsertText(this.h, text_ms)
}

func (this *QTextCursor) InsertText2(text string, format *QTextCharFormat) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTextCursor_InsertText2(this.h, text_ms, format.cPointer())
}

func (this *QTextCursor) MovePosition(op MoveOperation) bool {
	return (bool)(QTextCursor_MovePosition(this.h, op))
}

func (this *QTextCursor) VisualNavigation() bool {
	return (bool)(QTextCursor_VisualNavigation(this.h))
}

func (this *QTextCursor) SetVisualNavigation(b bool) {
	QTextCursor_SetVisualNavigation(this.h, (bool)(b))
}

func (this *QTextCursor) SetVerticalMovementX(x int) {
	QTextCursor_SetVerticalMovementX(this.h, (int)(x))
}

func (this *QTextCursor) VerticalMovementX() int {
	return (int)(QTextCursor_VerticalMovementX(this.h))
}

func (this *QTextCursor) SetKeepPositionOnInsert(b bool) {
	QTextCursor_SetKeepPositionOnInsert(this.h, (bool)(b))
}

func (this *QTextCursor) KeepPositionOnInsert() bool {
	return (bool)(QTextCursor_KeepPositionOnInsert(this.h))
}

func (this *QTextCursor) DeleteChar() {
	QTextCursor_DeleteChar(this.h)
}

func (this *QTextCursor) DeletePreviousChar() {
	QTextCursor_DeletePreviousChar(this.h)
}

func (this *QTextCursor) Select(selection SelectionType) {
	QTextCursor_Select(this.h, selection)
}

func (this *QTextCursor) HasSelection() bool {
	return (bool)(QTextCursor_HasSelection(this.h))
}

func (this *QTextCursor) HasComplexSelection() bool {
	return (bool)(QTextCursor_HasComplexSelection(this.h))
}

func (this *QTextCursor) RemoveSelectedText() {
	QTextCursor_RemoveSelectedText(this.h)
}

func (this *QTextCursor) ClearSelection() {
	QTextCursor_ClearSelection(this.h)
}

func (this *QTextCursor) SelectionStart() int {
	return (int)(QTextCursor_SelectionStart(this.h))
}

func (this *QTextCursor) SelectionEnd() int {
	return (int)(QTextCursor_SelectionEnd(this.h))
}

func (this *QTextCursor) SelectedText() string {
	var _ms struct_miqt_string = QTextCursor_SelectedText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextCursor) Selection() *QTextDocumentFragment {
	_goptr := newQTextDocumentFragment(QTextCursor_Selection(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCursor) SelectedTableCells(firstRow *int, numRows *int, firstColumn *int, numColumns *int) {
	QTextCursor_SelectedTableCells(this.h, (*int)(unsafe.Pointer(firstRow)), (*int)(unsafe.Pointer(numRows)), (*int)(unsafe.Pointer(firstColumn)), (*int)(unsafe.Pointer(numColumns)))
}

func (this *QTextCursor) Block() *QTextBlock {
	_goptr := newQTextBlock(QTextCursor_Block(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCursor) CharFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QTextCursor_CharFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCursor) SetCharFormat(format *QTextCharFormat) {
	QTextCursor_SetCharFormat(this.h, format.cPointer())
}

func (this *QTextCursor) MergeCharFormat(modifier *QTextCharFormat) {
	QTextCursor_MergeCharFormat(this.h, modifier.cPointer())
}

func (this *QTextCursor) BlockFormat() *QTextBlockFormat {
	_goptr := newQTextBlockFormat(QTextCursor_BlockFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCursor) SetBlockFormat(format *QTextBlockFormat) {
	QTextCursor_SetBlockFormat(this.h, format.cPointer())
}

func (this *QTextCursor) MergeBlockFormat(modifier *QTextBlockFormat) {
	QTextCursor_MergeBlockFormat(this.h, modifier.cPointer())
}

func (this *QTextCursor) BlockCharFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QTextCursor_BlockCharFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCursor) SetBlockCharFormat(format *QTextCharFormat) {
	QTextCursor_SetBlockCharFormat(this.h, format.cPointer())
}

func (this *QTextCursor) MergeBlockCharFormat(modifier *QTextCharFormat) {
	QTextCursor_MergeBlockCharFormat(this.h, modifier.cPointer())
}

func (this *QTextCursor) AtBlockStart() bool {
	return (bool)(QTextCursor_AtBlockStart(this.h))
}

func (this *QTextCursor) AtBlockEnd() bool {
	return (bool)(QTextCursor_AtBlockEnd(this.h))
}

func (this *QTextCursor) AtStart() bool {
	return (bool)(QTextCursor_AtStart(this.h))
}

func (this *QTextCursor) AtEnd() bool {
	return (bool)(QTextCursor_AtEnd(this.h))
}

func (this *QTextCursor) InsertBlock() {
	QTextCursor_InsertBlock(this.h)
}

func (this *QTextCursor) InsertBlockWithFormat(format *QTextBlockFormat) {
	QTextCursor_InsertBlockWithFormat(this.h, format.cPointer())
}

func (this *QTextCursor) InsertBlock2(format *QTextBlockFormat, charFormat *QTextCharFormat) {
	QTextCursor_InsertBlock2(this.h, format.cPointer(), charFormat.cPointer())
}

func (this *QTextCursor) InsertList(format *QTextListFormat) *QTextList {
	return newQTextList(QTextCursor_InsertList(this.h, format.cPointer()))
}

func (this *QTextCursor) InsertListWithStyle(style QTextListFormat__Style) *QTextList {
	return newQTextList(QTextCursor_InsertListWithStyle(this.h, (int)(style)))
}

func (this *QTextCursor) CreateList(format *QTextListFormat) *QTextList {
	return newQTextList(QTextCursor_CreateList(this.h, format.cPointer()))
}

func (this *QTextCursor) CreateListWithStyle(style QTextListFormat__Style) *QTextList {
	return newQTextList(QTextCursor_CreateListWithStyle(this.h, (int)(style)))
}

func (this *QTextCursor) CurrentList() *QTextList {
	return newQTextList(QTextCursor_CurrentList(this.h))
}

func (this *QTextCursor) InsertTable(rows int, cols int, format *QTextTableFormat) *QTextTable {
	return newQTextTable(QTextCursor_InsertTable(this.h, (int)(rows), (int)(cols), format.cPointer()))
}

func (this *QTextCursor) InsertTable2(rows int, cols int) *QTextTable {
	return newQTextTable(QTextCursor_InsertTable2(this.h, (int)(rows), (int)(cols)))
}

func (this *QTextCursor) CurrentTable() *QTextTable {
	return newQTextTable(QTextCursor_CurrentTable(this.h))
}

func (this *QTextCursor) InsertFrame(format *QTextFrameFormat) *QTextFrame {
	return newQTextFrame(QTextCursor_InsertFrame(this.h, format.cPointer()))
}

func (this *QTextCursor) CurrentFrame() *QTextFrame {
	return newQTextFrame(QTextCursor_CurrentFrame(this.h))
}

func (this *QTextCursor) InsertFragment(fragment *QTextDocumentFragment) {
	QTextCursor_InsertFragment(this.h, fragment.cPointer())
}

func (this *QTextCursor) InsertHtml(html string) {
	html_ms := struct_miqt_string{}
	html_ms.data = CString(html)
	html_ms.len = size_t(len(html))
	defer free(unsafe.Pointer(html_ms.data))
	QTextCursor_InsertHtml(this.h, html_ms)
}

func (this *QTextCursor) InsertMarkdown(markdown string) {
	markdown_ms := struct_miqt_string{}
	markdown_ms.data = CString(markdown)
	markdown_ms.len = size_t(len(markdown))
	defer free(unsafe.Pointer(markdown_ms.data))
	QTextCursor_InsertMarkdown(this.h, markdown_ms)
}

func (this *QTextCursor) InsertImage(format *QTextImageFormat, alignment QTextFrameFormat__Position) {
	QTextCursor_InsertImage(this.h, format.cPointer(), (int)(alignment))
}

func (this *QTextCursor) InsertImageWithFormat(format *QTextImageFormat) {
	QTextCursor_InsertImageWithFormat(this.h, format.cPointer())
}

func (this *QTextCursor) InsertImageWithName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QTextCursor_InsertImageWithName(this.h, name_ms)
}

func (this *QTextCursor) InsertImageWithImage(image *QImage) {
	QTextCursor_InsertImageWithImage(this.h, image.cPointer())
}

func (this *QTextCursor) BeginEditBlock() {
	QTextCursor_BeginEditBlock(this.h)
}

func (this *QTextCursor) JoinPreviousEditBlock() {
	QTextCursor_JoinPreviousEditBlock(this.h)
}

func (this *QTextCursor) EndEditBlock() {
	QTextCursor_EndEditBlock(this.h)
}

func (this *QTextCursor) OperatorNotEqual(rhs *QTextCursor) bool {
	return (bool)(QTextCursor_OperatorNotEqual(this.h, rhs.cPointer()))
}

func (this *QTextCursor) OperatorLesser(rhs *QTextCursor) bool {
	return (bool)(QTextCursor_OperatorLesser(this.h, rhs.cPointer()))
}

func (this *QTextCursor) OperatorLesserOrEqual(rhs *QTextCursor) bool {
	return (bool)(QTextCursor_OperatorLesserOrEqual(this.h, rhs.cPointer()))
}

func (this *QTextCursor) OperatorEqual(rhs *QTextCursor) bool {
	return (bool)(QTextCursor_OperatorEqual(this.h, rhs.cPointer()))
}

func (this *QTextCursor) OperatorGreaterOrEqual(rhs *QTextCursor) bool {
	return (bool)(QTextCursor_OperatorGreaterOrEqual(this.h, rhs.cPointer()))
}

func (this *QTextCursor) OperatorGreater(rhs *QTextCursor) bool {
	return (bool)(QTextCursor_OperatorGreater(this.h, rhs.cPointer()))
}

func (this *QTextCursor) IsCopyOf(other *QTextCursor) bool {
	return (bool)(QTextCursor_IsCopyOf(this.h, other.cPointer()))
}

func (this *QTextCursor) BlockNumber() int {
	return (int)(QTextCursor_BlockNumber(this.h))
}

func (this *QTextCursor) ColumnNumber() int {
	return (int)(QTextCursor_ColumnNumber(this.h))
}

func (this *QTextCursor) Document() *QTextDocument {
	return newQTextDocument(QTextCursor_Document(this.h))
}

func (this *QTextCursor) SetPosition2(pos int, mode MoveMode) {
	QTextCursor_SetPosition2(this.h, (int)(pos), mode)
}

func (this *QTextCursor) MovePosition2(op MoveOperation, param2 MoveMode) bool {
	return (bool)(QTextCursor_MovePosition2(this.h, op, param2))
}

func (this *QTextCursor) MovePosition3(op MoveOperation, param2 MoveMode, n int) bool {
	return (bool)(QTextCursor_MovePosition3(this.h, op, param2, (int)(n)))
}

func (this *QTextCursor) InsertMarkdown2(markdown string, features MarkdownFeature) {
	markdown_ms := struct_miqt_string{}
	markdown_ms.data = CString(markdown)
	markdown_ms.len = size_t(len(markdown))
	defer free(unsafe.Pointer(markdown_ms.data))
	QTextCursor_InsertMarkdown2(this.h, markdown_ms, (int)(features))
}

func (this *QTextCursor) InsertImage2(image *QImage, name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QTextCursor_InsertImage2(this.h, image.cPointer(), name_ms)
}
