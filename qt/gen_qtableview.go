package qt

import (
	"unsafe"
)

type QTableView struct {
	h          uintptr
	isSubclass bool
}

// NewQTableView constructs a new QTableView object.
func NewQTableView(parent *QWidget) *QTableView {
	g := newQTableView(QTableView_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQTableView2 constructs a new QTableView object.
func NewQTableView2() *QTableView {
	g := newQTableView(QTableView_new2())
	g.isSubclass = true
	return g
}

func (this *QTableView) MetaObject() *QMetaObject {
	return newQMetaObject(QTableView_MetaObject(this.h))
}

func (this *QTableView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTableView_Metacast(this.h, param1_Cstring))
}

func QTableView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTableView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableView) SetModel(model *QAbstractItemModel) {
	QTableView_SetModel(this.h, model.cPointer())
}

func (this *QTableView) SetRootIndex(index *QModelIndex) {
	QTableView_SetRootIndex(this.h, index.cPointer())
}

func (this *QTableView) SetSelectionModel(selectionModel *QItemSelectionModel) {
	QTableView_SetSelectionModel(this.h, selectionModel.cPointer())
}

func (this *QTableView) DoItemsLayout() {
	QTableView_DoItemsLayout(this.h)
}

func (this *QTableView) HorizontalHeader() *QHeaderView {
	return newQHeaderView(QTableView_HorizontalHeader(this.h))
}

func (this *QTableView) VerticalHeader() *QHeaderView {
	return newQHeaderView(QTableView_VerticalHeader(this.h))
}

func (this *QTableView) SetHorizontalHeader(header *QHeaderView) {
	QTableView_SetHorizontalHeader(this.h, header.cPointer())
}

func (this *QTableView) SetVerticalHeader(header *QHeaderView) {
	QTableView_SetVerticalHeader(this.h, header.cPointer())
}

func (this *QTableView) RowViewportPosition(row int) int {
	return (int)(QTableView_RowViewportPosition(this.h, (int)(row)))
}

func (this *QTableView) RowAt(y int) int {
	return (int)(QTableView_RowAt(this.h, (int)(y)))
}

func (this *QTableView) SetRowHeight(row int, height int) {
	QTableView_SetRowHeight(this.h, (int)(row), (int)(height))
}

func (this *QTableView) RowHeight(row int) int {
	return (int)(QTableView_RowHeight(this.h, (int)(row)))
}

func (this *QTableView) ColumnViewportPosition(column int) int {
	return (int)(QTableView_ColumnViewportPosition(this.h, (int)(column)))
}

func (this *QTableView) ColumnAt(x int) int {
	return (int)(QTableView_ColumnAt(this.h, (int)(x)))
}

func (this *QTableView) SetColumnWidth(column int, width int) {
	QTableView_SetColumnWidth(this.h, (int)(column), (int)(width))
}

func (this *QTableView) ColumnWidth(column int) int {
	return (int)(QTableView_ColumnWidth(this.h, (int)(column)))
}

func (this *QTableView) IsRowHidden(row int) bool {
	return (bool)(QTableView_IsRowHidden(this.h, (int)(row)))
}

func (this *QTableView) SetRowHidden(row int, hide bool) {
	QTableView_SetRowHidden(this.h, (int)(row), (bool)(hide))
}

func (this *QTableView) IsColumnHidden(column int) bool {
	return (bool)(QTableView_IsColumnHidden(this.h, (int)(column)))
}

func (this *QTableView) SetColumnHidden(column int, hide bool) {
	QTableView_SetColumnHidden(this.h, (int)(column), (bool)(hide))
}

func (this *QTableView) SetSortingEnabled(enable bool) {
	QTableView_SetSortingEnabled(this.h, (bool)(enable))
}

func (this *QTableView) IsSortingEnabled() bool {
	return (bool)(QTableView_IsSortingEnabled(this.h))
}

func (this *QTableView) ShowGrid() bool {
	return (bool)(QTableView_ShowGrid(this.h))
}

func (this *QTableView) GridStyle() PenStyle {
	return (PenStyle)(QTableView_GridStyle(this.h))
}

func (this *QTableView) SetGridStyle(style PenStyle) {
	QTableView_SetGridStyle(this.h, (int)(style))
}

func (this *QTableView) SetWordWrap(on bool) {
	QTableView_SetWordWrap(this.h, (bool)(on))
}

func (this *QTableView) WordWrap() bool {
	return (bool)(QTableView_WordWrap(this.h))
}

func (this *QTableView) SetCornerButtonEnabled(enable bool) {
	QTableView_SetCornerButtonEnabled(this.h, (bool)(enable))
}

func (this *QTableView) IsCornerButtonEnabled() bool {
	return (bool)(QTableView_IsCornerButtonEnabled(this.h))
}

func (this *QTableView) VisualRect(index *QModelIndex) *QRect {
	_goptr := newQRect(QTableView_VisualRect(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableView) ScrollTo(index *QModelIndex, hint ScrollHint) {
	QTableView_ScrollTo(this.h, index.cPointer(), hint)
}

func (this *QTableView) IndexAt(p *QPoint) *QModelIndex {
	_goptr := newQModelIndex(QTableView_IndexAt(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableView) SetSpan(row int, column int, rowSpan int, columnSpan int) {
	QTableView_SetSpan(this.h, (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan))
}

func (this *QTableView) RowSpan(row int, column int) int {
	return (int)(QTableView_RowSpan(this.h, (int)(row), (int)(column)))
}

func (this *QTableView) ColumnSpan(row int, column int) int {
	return (int)(QTableView_ColumnSpan(this.h, (int)(row), (int)(column)))
}

func (this *QTableView) ClearSpans() {
	QTableView_ClearSpans(this.h)
}

func (this *QTableView) SelectRow(row int) {
	QTableView_SelectRow(this.h, (int)(row))
}

func (this *QTableView) SelectColumn(column int) {
	QTableView_SelectColumn(this.h, (int)(column))
}

func (this *QTableView) HideRow(row int) {
	QTableView_HideRow(this.h, (int)(row))
}

func (this *QTableView) HideColumn(column int) {
	QTableView_HideColumn(this.h, (int)(column))
}

func (this *QTableView) ShowRow(row int) {
	QTableView_ShowRow(this.h, (int)(row))
}

func (this *QTableView) ShowColumn(column int) {
	QTableView_ShowColumn(this.h, (int)(column))
}

func (this *QTableView) ResizeRowToContents(row int) {
	QTableView_ResizeRowToContents(this.h, (int)(row))
}

func (this *QTableView) ResizeRowsToContents() {
	QTableView_ResizeRowsToContents(this.h)
}

func (this *QTableView) ResizeColumnToContents(column int) {
	QTableView_ResizeColumnToContents(this.h, (int)(column))
}

func (this *QTableView) ResizeColumnsToContents() {
	QTableView_ResizeColumnsToContents(this.h)
}

func (this *QTableView) SortByColumn(column int, order SortOrder) {
	QTableView_SortByColumn(this.h, (int)(column), (int)(order))
}

func (this *QTableView) SetShowGrid(show bool) {
	QTableView_SetShowGrid(this.h, (bool)(show))
}

func QTableView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTableView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTableView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTableView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableView) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTableView_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTableView) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_MetaObject
func miqt_exec_callback_QTableView_MetaObject(self QTableView, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTableView) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTableView_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTableView) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_Metacast
func miqt_exec_callback_QTableView_Metacast(self QTableView, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
