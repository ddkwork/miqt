package qt

/*

#include "gen_qtextcursor.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QTextCursor__MoveMode int

const (
	QTextCursor__MoveMode__MoveAnchor QTextCursor__MoveMode = 0
	QTextCursor__MoveMode__KeepAnchor QTextCursor__MoveMode = 1
)

type QTextCursor__MoveOperation int

const (
	QTextCursor__MoveOperation__NoMove            QTextCursor__MoveOperation = 0
	QTextCursor__MoveOperation__Start             QTextCursor__MoveOperation = 1
	QTextCursor__MoveOperation__Up                QTextCursor__MoveOperation = 2
	QTextCursor__MoveOperation__StartOfLine       QTextCursor__MoveOperation = 3
	QTextCursor__MoveOperation__StartOfBlock      QTextCursor__MoveOperation = 4
	QTextCursor__MoveOperation__StartOfWord       QTextCursor__MoveOperation = 5
	QTextCursor__MoveOperation__PreviousBlock     QTextCursor__MoveOperation = 6
	QTextCursor__MoveOperation__PreviousCharacter QTextCursor__MoveOperation = 7
	QTextCursor__MoveOperation__PreviousWord      QTextCursor__MoveOperation = 8
	QTextCursor__MoveOperation__Left              QTextCursor__MoveOperation = 9
	QTextCursor__MoveOperation__WordLeft          QTextCursor__MoveOperation = 10
	QTextCursor__MoveOperation__End               QTextCursor__MoveOperation = 11
	QTextCursor__MoveOperation__Down              QTextCursor__MoveOperation = 12
	QTextCursor__MoveOperation__EndOfLine         QTextCursor__MoveOperation = 13
	QTextCursor__MoveOperation__EndOfWord         QTextCursor__MoveOperation = 14
	QTextCursor__MoveOperation__EndOfBlock        QTextCursor__MoveOperation = 15
	QTextCursor__MoveOperation__NextBlock         QTextCursor__MoveOperation = 16
	QTextCursor__MoveOperation__NextCharacter     QTextCursor__MoveOperation = 17
	QTextCursor__MoveOperation__NextWord          QTextCursor__MoveOperation = 18
	QTextCursor__MoveOperation__Right             QTextCursor__MoveOperation = 19
	QTextCursor__MoveOperation__WordRight         QTextCursor__MoveOperation = 20
	QTextCursor__MoveOperation__NextCell          QTextCursor__MoveOperation = 21
	QTextCursor__MoveOperation__PreviousCell      QTextCursor__MoveOperation = 22
	QTextCursor__MoveOperation__NextRow           QTextCursor__MoveOperation = 23
	QTextCursor__MoveOperation__PreviousRow       QTextCursor__MoveOperation = 24
)

type QTextCursor__SelectionType int

const (
	QTextCursor__SelectionType__WordUnderCursor  QTextCursor__SelectionType = 0
	QTextCursor__SelectionType__LineUnderCursor  QTextCursor__SelectionType = 1
	QTextCursor__SelectionType__BlockUnderCursor QTextCursor__SelectionType = 2
	QTextCursor__SelectionType__Document         QTextCursor__SelectionType = 3
)

type QTextCursor struct {
	h *C.QTextCursor
}

func (this *QTextCursor) cPointer() *C.QTextCursor {
	if this == nil {
		return nil
	}
	return this.h
}

func newQTextCursor(h *C.QTextCursor) *QTextCursor {
	if h == nil {
		return nil
	}
	return &QTextCursor{h: h}
}

func newQTextCursor_U(h unsafe.Pointer) *QTextCursor {
	return newQTextCursor((*C.QTextCursor)(h))
}

// NewQTextCursor constructs a new QTextCursor object.
func NewQTextCursor() *QTextCursor {
	ret := C.QTextCursor_new()
	return newQTextCursor(ret)
}

// NewQTextCursor2 constructs a new QTextCursor object.
func NewQTextCursor2(document *QTextDocument) *QTextCursor {
	ret := C.QTextCursor_new2(document.cPointer())
	return newQTextCursor(ret)
}

// NewQTextCursor3 constructs a new QTextCursor object.
func NewQTextCursor3(frame *QTextFrame) *QTextCursor {
	ret := C.QTextCursor_new3(frame.cPointer())
	return newQTextCursor(ret)
}

// NewQTextCursor4 constructs a new QTextCursor object.
func NewQTextCursor4(block *QTextBlock) *QTextCursor {
	ret := C.QTextCursor_new4(block.cPointer())
	return newQTextCursor(ret)
}

// NewQTextCursor5 constructs a new QTextCursor object.
func NewQTextCursor5(cursor *QTextCursor) *QTextCursor {
	ret := C.QTextCursor_new5(cursor.cPointer())
	return newQTextCursor(ret)
}

func (this *QTextCursor) OperatorAssign(other *QTextCursor) {
	C.QTextCursor_OperatorAssign(this.h, other.cPointer())
}

func (this *QTextCursor) Swap(other *QTextCursor) {
	C.QTextCursor_Swap(this.h, other.cPointer())
}

func (this *QTextCursor) IsNull() bool {
	ret := C.QTextCursor_IsNull(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) SetPosition(pos int) {
	C.QTextCursor_SetPosition(this.h, (C.int)(pos))
}

func (this *QTextCursor) Position() int {
	ret := C.QTextCursor_Position(this.h)
	return (int)(ret)
}

func (this *QTextCursor) PositionInBlock() int {
	ret := C.QTextCursor_PositionInBlock(this.h)
	return (int)(ret)
}

func (this *QTextCursor) Anchor() int {
	ret := C.QTextCursor_Anchor(this.h)
	return (int)(ret)
}

func (this *QTextCursor) InsertText(text string) {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	C.QTextCursor_InsertText(this.h, text_Cstring, C.size_t(len(text)))
}

func (this *QTextCursor) InsertText2(text string, format *QTextCharFormat) {
	text_Cstring := C.CString(text)
	defer C.free(unsafe.Pointer(text_Cstring))
	C.QTextCursor_InsertText2(this.h, text_Cstring, C.size_t(len(text)), format.cPointer())
}

func (this *QTextCursor) MovePosition(op QTextCursor__MoveOperation) bool {
	ret := C.QTextCursor_MovePosition(this.h, (C.uintptr_t)(op))
	return (bool)(ret)
}

func (this *QTextCursor) VisualNavigation() bool {
	ret := C.QTextCursor_VisualNavigation(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) SetVisualNavigation(b bool) {
	C.QTextCursor_SetVisualNavigation(this.h, (C.bool)(b))
}

func (this *QTextCursor) SetVerticalMovementX(x int) {
	C.QTextCursor_SetVerticalMovementX(this.h, (C.int)(x))
}

func (this *QTextCursor) VerticalMovementX() int {
	ret := C.QTextCursor_VerticalMovementX(this.h)
	return (int)(ret)
}

func (this *QTextCursor) SetKeepPositionOnInsert(b bool) {
	C.QTextCursor_SetKeepPositionOnInsert(this.h, (C.bool)(b))
}

func (this *QTextCursor) KeepPositionOnInsert() bool {
	ret := C.QTextCursor_KeepPositionOnInsert(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) DeleteChar() {
	C.QTextCursor_DeleteChar(this.h)
}

func (this *QTextCursor) DeletePreviousChar() {
	C.QTextCursor_DeletePreviousChar(this.h)
}

func (this *QTextCursor) Select(selection QTextCursor__SelectionType) {
	C.QTextCursor_Select(this.h, (C.uintptr_t)(selection))
}

func (this *QTextCursor) HasSelection() bool {
	ret := C.QTextCursor_HasSelection(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) HasComplexSelection() bool {
	ret := C.QTextCursor_HasComplexSelection(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) RemoveSelectedText() {
	C.QTextCursor_RemoveSelectedText(this.h)
}

func (this *QTextCursor) ClearSelection() {
	C.QTextCursor_ClearSelection(this.h)
}

func (this *QTextCursor) SelectionStart() int {
	ret := C.QTextCursor_SelectionStart(this.h)
	return (int)(ret)
}

func (this *QTextCursor) SelectionEnd() int {
	ret := C.QTextCursor_SelectionEnd(this.h)
	return (int)(ret)
}

func (this *QTextCursor) SelectedText() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QTextCursor_SelectedText(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QTextCursor) Selection() *QTextDocumentFragment {
	ret := C.QTextCursor_Selection(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextDocumentFragment(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextDocumentFragment) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextCursor) SelectedTableCells(firstRow *int, numRows *int, firstColumn *int, numColumns *int) {
	C.QTextCursor_SelectedTableCells(this.h, (*C.int)(unsafe.Pointer(firstRow)), (*C.int)(unsafe.Pointer(numRows)), (*C.int)(unsafe.Pointer(firstColumn)), (*C.int)(unsafe.Pointer(numColumns)))
}

func (this *QTextCursor) Block() *QTextBlock {
	ret := C.QTextCursor_Block(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextBlock(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextBlock) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextCursor) CharFormat() *QTextCharFormat {
	ret := C.QTextCursor_CharFormat(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextCharFormat(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextCharFormat) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextCursor) SetCharFormat(format *QTextCharFormat) {
	C.QTextCursor_SetCharFormat(this.h, format.cPointer())
}

func (this *QTextCursor) MergeCharFormat(modifier *QTextCharFormat) {
	C.QTextCursor_MergeCharFormat(this.h, modifier.cPointer())
}

func (this *QTextCursor) BlockFormat() *QTextBlockFormat {
	ret := C.QTextCursor_BlockFormat(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextBlockFormat(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextBlockFormat) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextCursor) SetBlockFormat(format *QTextBlockFormat) {
	C.QTextCursor_SetBlockFormat(this.h, format.cPointer())
}

func (this *QTextCursor) MergeBlockFormat(modifier *QTextBlockFormat) {
	C.QTextCursor_MergeBlockFormat(this.h, modifier.cPointer())
}

func (this *QTextCursor) BlockCharFormat() *QTextCharFormat {
	ret := C.QTextCursor_BlockCharFormat(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQTextCharFormat(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QTextCharFormat) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QTextCursor) SetBlockCharFormat(format *QTextCharFormat) {
	C.QTextCursor_SetBlockCharFormat(this.h, format.cPointer())
}

func (this *QTextCursor) MergeBlockCharFormat(modifier *QTextCharFormat) {
	C.QTextCursor_MergeBlockCharFormat(this.h, modifier.cPointer())
}

func (this *QTextCursor) AtBlockStart() bool {
	ret := C.QTextCursor_AtBlockStart(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) AtBlockEnd() bool {
	ret := C.QTextCursor_AtBlockEnd(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) AtStart() bool {
	ret := C.QTextCursor_AtStart(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) AtEnd() bool {
	ret := C.QTextCursor_AtEnd(this.h)
	return (bool)(ret)
}

func (this *QTextCursor) InsertBlock() {
	C.QTextCursor_InsertBlock(this.h)
}

func (this *QTextCursor) InsertBlockWithFormat(format *QTextBlockFormat) {
	C.QTextCursor_InsertBlockWithFormat(this.h, format.cPointer())
}

func (this *QTextCursor) InsertBlock2(format *QTextBlockFormat, charFormat *QTextCharFormat) {
	C.QTextCursor_InsertBlock2(this.h, format.cPointer(), charFormat.cPointer())
}

func (this *QTextCursor) InsertList(format *QTextListFormat) *QTextList {
	ret := C.QTextCursor_InsertList(this.h, format.cPointer())
	return newQTextList_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) InsertListWithStyle(style QTextListFormat__Style) *QTextList {
	ret := C.QTextCursor_InsertListWithStyle(this.h, (C.uintptr_t)(style))
	return newQTextList_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) CreateList(format *QTextListFormat) *QTextList {
	ret := C.QTextCursor_CreateList(this.h, format.cPointer())
	return newQTextList_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) CreateListWithStyle(style QTextListFormat__Style) *QTextList {
	ret := C.QTextCursor_CreateListWithStyle(this.h, (C.uintptr_t)(style))
	return newQTextList_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) CurrentList() *QTextList {
	ret := C.QTextCursor_CurrentList(this.h)
	return newQTextList_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) InsertTable(rows int, cols int, format *QTextTableFormat) *QTextTable {
	ret := C.QTextCursor_InsertTable(this.h, (C.int)(rows), (C.int)(cols), format.cPointer())
	return newQTextTable_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) InsertTable2(rows int, cols int) *QTextTable {
	ret := C.QTextCursor_InsertTable2(this.h, (C.int)(rows), (C.int)(cols))
	return newQTextTable_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) CurrentTable() *QTextTable {
	ret := C.QTextCursor_CurrentTable(this.h)
	return newQTextTable_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) InsertFrame(format *QTextFrameFormat) *QTextFrame {
	ret := C.QTextCursor_InsertFrame(this.h, format.cPointer())
	return newQTextFrame_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) CurrentFrame() *QTextFrame {
	ret := C.QTextCursor_CurrentFrame(this.h)
	return newQTextFrame_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) InsertFragment(fragment *QTextDocumentFragment) {
	C.QTextCursor_InsertFragment(this.h, fragment.cPointer())
}

func (this *QTextCursor) InsertHtml(html string) {
	html_Cstring := C.CString(html)
	defer C.free(unsafe.Pointer(html_Cstring))
	C.QTextCursor_InsertHtml(this.h, html_Cstring, C.size_t(len(html)))
}

func (this *QTextCursor) InsertImage(format *QTextImageFormat, alignment QTextFrameFormat__Position) {
	C.QTextCursor_InsertImage(this.h, format.cPointer(), (C.uintptr_t)(alignment))
}

func (this *QTextCursor) InsertImageWithFormat(format *QTextImageFormat) {
	C.QTextCursor_InsertImageWithFormat(this.h, format.cPointer())
}

func (this *QTextCursor) InsertImageWithName(name string) {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	C.QTextCursor_InsertImageWithName(this.h, name_Cstring, C.size_t(len(name)))
}

func (this *QTextCursor) InsertImageWithImage(image *QImage) {
	C.QTextCursor_InsertImageWithImage(this.h, image.cPointer())
}

func (this *QTextCursor) BeginEditBlock() {
	C.QTextCursor_BeginEditBlock(this.h)
}

func (this *QTextCursor) JoinPreviousEditBlock() {
	C.QTextCursor_JoinPreviousEditBlock(this.h)
}

func (this *QTextCursor) EndEditBlock() {
	C.QTextCursor_EndEditBlock(this.h)
}

func (this *QTextCursor) OperatorNotEqual(rhs *QTextCursor) bool {
	ret := C.QTextCursor_OperatorNotEqual(this.h, rhs.cPointer())
	return (bool)(ret)
}

func (this *QTextCursor) OperatorLesser(rhs *QTextCursor) bool {
	ret := C.QTextCursor_OperatorLesser(this.h, rhs.cPointer())
	return (bool)(ret)
}

func (this *QTextCursor) OperatorLesserOrEqual(rhs *QTextCursor) bool {
	ret := C.QTextCursor_OperatorLesserOrEqual(this.h, rhs.cPointer())
	return (bool)(ret)
}

func (this *QTextCursor) OperatorEqual(rhs *QTextCursor) bool {
	ret := C.QTextCursor_OperatorEqual(this.h, rhs.cPointer())
	return (bool)(ret)
}

func (this *QTextCursor) OperatorGreaterOrEqual(rhs *QTextCursor) bool {
	ret := C.QTextCursor_OperatorGreaterOrEqual(this.h, rhs.cPointer())
	return (bool)(ret)
}

func (this *QTextCursor) OperatorGreater(rhs *QTextCursor) bool {
	ret := C.QTextCursor_OperatorGreater(this.h, rhs.cPointer())
	return (bool)(ret)
}

func (this *QTextCursor) IsCopyOf(other *QTextCursor) bool {
	ret := C.QTextCursor_IsCopyOf(this.h, other.cPointer())
	return (bool)(ret)
}

func (this *QTextCursor) BlockNumber() int {
	ret := C.QTextCursor_BlockNumber(this.h)
	return (int)(ret)
}

func (this *QTextCursor) ColumnNumber() int {
	ret := C.QTextCursor_ColumnNumber(this.h)
	return (int)(ret)
}

func (this *QTextCursor) Document() *QTextDocument {
	ret := C.QTextCursor_Document(this.h)
	return newQTextDocument_U(unsafe.Pointer(ret))
}

func (this *QTextCursor) SetPosition2(pos int, mode QTextCursor__MoveMode) {
	C.QTextCursor_SetPosition2(this.h, (C.int)(pos), (C.uintptr_t)(mode))
}

func (this *QTextCursor) MovePosition2(op QTextCursor__MoveOperation, param2 QTextCursor__MoveMode) bool {
	ret := C.QTextCursor_MovePosition2(this.h, (C.uintptr_t)(op), (C.uintptr_t)(param2))
	return (bool)(ret)
}

func (this *QTextCursor) MovePosition3(op QTextCursor__MoveOperation, param2 QTextCursor__MoveMode, n int) bool {
	ret := C.QTextCursor_MovePosition3(this.h, (C.uintptr_t)(op), (C.uintptr_t)(param2), (C.int)(n))
	return (bool)(ret)
}

func (this *QTextCursor) InsertImage2(image *QImage, name string) {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	C.QTextCursor_InsertImage2(this.h, image.cPointer(), name_Cstring, C.size_t(len(name)))
}

func (this *QTextCursor) Delete() {
	C.QTextCursor_Delete(this.h)
}
