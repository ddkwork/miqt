package qt

import (
	"unsafe"
)

type QTextTableCell struct {
	h          uintptr
	isSubclass bool
}

// NewQTextTableCell constructs a new QTextTableCell object.
func NewQTextTableCell() *QTextTableCell {
	g := newQTextTableCell(QTextTableCell_new())
	g.isSubclass = true
	return g
}

// NewQTextTableCell2 constructs a new QTextTableCell object.
func NewQTextTableCell2(o *QTextTableCell) *QTextTableCell {
	g := newQTextTableCell(QTextTableCell_new2(o.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QTextTableCell) OperatorAssign(o *QTextTableCell) {
	QTextTableCell_OperatorAssign(this.h, o.cPointer())
}

func (this *QTextTableCell) SetFormat(format *QTextCharFormat) {
	QTextTableCell_SetFormat(this.h, format.cPointer())
}

func (this *QTextTableCell) Format() *QTextCharFormat {
	_goptr := newQTextCharFormat(QTextTableCell_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCell) Row() int {
	return (int)(QTextTableCell_Row(this.h))
}

func (this *QTextTableCell) Column() int {
	return (int)(QTextTableCell_Column(this.h))
}

func (this *QTextTableCell) RowSpan() int {
	return (int)(QTextTableCell_RowSpan(this.h))
}

func (this *QTextTableCell) ColumnSpan() int {
	return (int)(QTextTableCell_ColumnSpan(this.h))
}

func (this *QTextTableCell) IsValid() bool {
	return (bool)(QTextTableCell_IsValid(this.h))
}

func (this *QTextTableCell) FirstCursorPosition() *QTextCursor {
	_goptr := newQTextCursor(QTextTableCell_FirstCursorPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCell) LastCursorPosition() *QTextCursor {
	_goptr := newQTextCursor(QTextTableCell_LastCursorPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCell) FirstPosition() int {
	return (int)(QTextTableCell_FirstPosition(this.h))
}

func (this *QTextTableCell) LastPosition() int {
	return (int)(QTextTableCell_LastPosition(this.h))
}

func (this *QTextTableCell) OperatorEqual(other *QTextTableCell) bool {
	return (bool)(QTextTableCell_OperatorEqual(this.h, other.cPointer()))
}

func (this *QTextTableCell) OperatorNotEqual(other *QTextTableCell) bool {
	return (bool)(QTextTableCell_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QTextTableCell) Begin() *QTextFrame__iterator {
	_goptr := newQTextFrame__iterator(QTextTableCell_Begin(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCell) End() *QTextFrame__iterator {
	_goptr := newQTextFrame__iterator(QTextTableCell_End(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCell) TableCellFormatIndex() int {
	return (int)(QTextTableCell_TableCellFormatIndex(this.h))
}

type QTextTable struct {
	h          uintptr
	isSubclass bool
}

// NewQTextTable constructs a new QTextTable object.
func NewQTextTable(doc *QTextDocument) *QTextTable {
	g := newQTextTable(QTextTable_new(doc.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QTextTable) MetaObject() *QMetaObject {
	return newQMetaObject(QTextTable_MetaObject(this.h))
}

func (this *QTextTable) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextTable_Metacast(this.h, param1_Cstring))
}

func QTextTable_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextTable_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextTable) Resize(rows int, cols int) {
	QTextTable_Resize(this.h, (int)(rows), (int)(cols))
}

func (this *QTextTable) InsertRows(pos int, num int) {
	QTextTable_InsertRows(this.h, (int)(pos), (int)(num))
}

func (this *QTextTable) InsertColumns(pos int, num int) {
	QTextTable_InsertColumns(this.h, (int)(pos), (int)(num))
}

func (this *QTextTable) AppendRows(count int) {
	QTextTable_AppendRows(this.h, (int)(count))
}

func (this *QTextTable) AppendColumns(count int) {
	QTextTable_AppendColumns(this.h, (int)(count))
}

func (this *QTextTable) RemoveRows(pos int, num int) {
	QTextTable_RemoveRows(this.h, (int)(pos), (int)(num))
}

func (this *QTextTable) RemoveColumns(pos int, num int) {
	QTextTable_RemoveColumns(this.h, (int)(pos), (int)(num))
}

func (this *QTextTable) MergeCells(row int, col int, numRows int, numCols int) {
	QTextTable_MergeCells(this.h, (int)(row), (int)(col), (int)(numRows), (int)(numCols))
}

func (this *QTextTable) MergeCellsWithCursor(cursor *QTextCursor) {
	QTextTable_MergeCellsWithCursor(this.h, cursor.cPointer())
}

func (this *QTextTable) SplitCell(row int, col int, numRows int, numCols int) {
	QTextTable_SplitCell(this.h, (int)(row), (int)(col), (int)(numRows), (int)(numCols))
}

func (this *QTextTable) Rows() int {
	return (int)(QTextTable_Rows(this.h))
}

func (this *QTextTable) Columns() int {
	return (int)(QTextTable_Columns(this.h))
}

func (this *QTextTable) CellAt(row int, col int) *QTextTableCell {
	_goptr := newQTextTableCell(QTextTable_CellAt(this.h, (int)(row), (int)(col)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTable) CellAtWithPosition(position int) *QTextTableCell {
	_goptr := newQTextTableCell(QTextTable_CellAtWithPosition(this.h, (int)(position)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTable) CellAtWithQTextCursor(c *QTextCursor) *QTextTableCell {
	_goptr := newQTextTableCell(QTextTable_CellAtWithQTextCursor(this.h, c.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTable) RowStart(c *QTextCursor) *QTextCursor {
	_goptr := newQTextCursor(QTextTable_RowStart(this.h, c.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTable) RowEnd(c *QTextCursor) *QTextCursor {
	_goptr := newQTextCursor(QTextTable_RowEnd(this.h, c.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTable) SetFormat(format *QTextTableFormat) {
	QTextTable_SetFormat(this.h, format.cPointer())
}

func (this *QTextTable) Format() *QTextTableFormat {
	_goptr := newQTextTableFormat(QTextTable_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTextTable_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextTable_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextTable_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextTable_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextTable) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTextTable_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTextTable) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextTable_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextTable_MetaObject
func miqt_exec_callback_QTextTable_MetaObject(self QTextTable, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextTable{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTextTable) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTextTable_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTextTable) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextTable_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextTable_Metacast
func miqt_exec_callback_QTextTable_Metacast(self QTextTable, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTextTable{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
