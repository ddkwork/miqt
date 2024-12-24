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

	ret := newQTableView(QTableView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTableView2 constructs a new QTableView object.
func NewQTableView2() *QTableView {

	ret := newQTableView(QTableView_new2())
	ret.isSubclass = true
	return ret
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

func (this *QTableView) callVirtualBase_SetModel(model *QAbstractItemModel) {

	QTableView_virtualbase_SetModel(unsafe.Pointer(this.h), model.cPointer())

}
func (this *QTableView) OnSetModel(slot func(super func(model *QAbstractItemModel), model *QAbstractItemModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SetModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SetModel
func miqt_exec_callback_QTableView_SetModel(self QTableView, cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(model *QAbstractItemModel), model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc((&QTableView{h: self}).callVirtualBase_SetModel, slotval1)

}

func (this *QTableView) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QTableView_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QTableView) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SetRootIndex
func miqt_exec_callback_QTableView_SetRootIndex(self QTableView, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QTableView{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QTableView) callVirtualBase_SetSelectionModel(selectionModel *QItemSelectionModel) {

	QTableView_virtualbase_SetSelectionModel(unsafe.Pointer(this.h), selectionModel.cPointer())

}
func (this *QTableView) OnSetSelectionModel(slot func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SetSelectionModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SetSelectionModel
func miqt_exec_callback_QTableView_SetSelectionModel(self QTableView, cb intptr_t, selectionModel *QItemSelectionModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelectionModel(selectionModel)

	gofunc((&QTableView{h: self}).callVirtualBase_SetSelectionModel, slotval1)

}

func (this *QTableView) callVirtualBase_DoItemsLayout() {

	QTableView_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QTableView) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_DoItemsLayout
func miqt_exec_callback_QTableView_DoItemsLayout(self QTableView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTableView{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QTableView) callVirtualBase_VisualRect(index *QModelIndex) *QRect {

	_goptr := newQRect(QTableView_virtualbase_VisualRect(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTableView) OnVisualRect(slot func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_VisualRect
func miqt_exec_callback_QTableView_VisualRect(self QTableView, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_VisualRect, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTableView) callVirtualBase_ScrollTo(index *QModelIndex, hint ScrollHint) {

	QTableView_virtualbase_ScrollTo(unsafe.Pointer(this.h), index.cPointer(), hint)

}
func (this *QTableView) OnScrollTo(slot func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_ScrollTo
func miqt_exec_callback_QTableView_ScrollTo(self QTableView, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc((&QTableView{h: self}).callVirtualBase_ScrollTo, slotval1, slotval2)

}

func (this *QTableView) callVirtualBase_IndexAt(p *QPoint) *QModelIndex {

	_goptr := newQModelIndex(QTableView_virtualbase_IndexAt(unsafe.Pointer(this.h), p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTableView) OnIndexAt(slot func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_IndexAt
func miqt_exec_callback_QTableView_IndexAt(self QTableView, cb intptr_t, p *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(p)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_IndexAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTableView) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QTableView_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QTableView) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_ScrollContentsBy
func miqt_exec_callback_QTableView_ScrollContentsBy(self QTableView, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QTableView{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QTableView) callVirtualBase_InitViewItemOption(option *QStyleOptionViewItem) {

	QTableView_virtualbase_InitViewItemOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QTableView) OnInitViewItemOption(slot func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_InitViewItemOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_InitViewItemOption
func miqt_exec_callback_QTableView_InitViewItemOption(self QTableView, cb intptr_t, option *QStyleOptionViewItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	gofunc((&QTableView{h: self}).callVirtualBase_InitViewItemOption, slotval1)

}

func (this *QTableView) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QTableView_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTableView) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_PaintEvent
func miqt_exec_callback_QTableView_PaintEvent(self QTableView, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QTableView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTableView) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QTableView_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_TimerEvent
func miqt_exec_callback_QTableView_TimerEvent(self QTableView, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTableView) callVirtualBase_DropEvent(event *QDropEvent) {

	QTableView_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_DropEvent
func miqt_exec_callback_QTableView_DropEvent(self QTableView, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTableView) callVirtualBase_HorizontalOffset() int {

	return (int)(QTableView_virtualbase_HorizontalOffset(unsafe.Pointer(this.h)))

}
func (this *QTableView) OnHorizontalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_HorizontalOffset
func miqt_exec_callback_QTableView_HorizontalOffset(self QTableView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_HorizontalOffset)

	return (int)(virtualReturn)

}

func (this *QTableView) callVirtualBase_VerticalOffset() int {

	return (int)(QTableView_virtualbase_VerticalOffset(unsafe.Pointer(this.h)))

}
func (this *QTableView) OnVerticalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_VerticalOffset
func miqt_exec_callback_QTableView_VerticalOffset(self QTableView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_VerticalOffset)

	return (int)(virtualReturn)

}

func (this *QTableView) callVirtualBase_MoveCursor(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex {

	_goptr := newQModelIndex(QTableView_virtualbase_MoveCursor(unsafe.Pointer(this.h), cursorAction, (int)(modifiers)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTableView) OnMoveCursor(slot func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_MoveCursor
func miqt_exec_callback_QTableView_MoveCursor(self QTableView, cb intptr_t, cursorAction CursorAction, modifiers int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(modifiers)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_MoveCursor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QTableView) callVirtualBase_SetSelection(rect *QRect, command SelectionFlag) {

	QTableView_virtualbase_SetSelection(unsafe.Pointer(this.h), rect.cPointer(), (int)(command))

}
func (this *QTableView) OnSetSelection(slot func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SetSelection
func miqt_exec_callback_QTableView_SetSelection(self QTableView, cb intptr_t, rect *QRect, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QTableView{h: self}).callVirtualBase_SetSelection, slotval1, slotval2)

}

func (this *QTableView) callVirtualBase_VisualRegionForSelection(selection *QItemSelection) *QRegion {

	_goptr := newQRegion(QTableView_virtualbase_VisualRegionForSelection(unsafe.Pointer(this.h), selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTableView) OnVisualRegionForSelection(slot func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_VisualRegionForSelection
func miqt_exec_callback_QTableView_VisualRegionForSelection(self QTableView, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_VisualRegionForSelection, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTableView) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QTableView_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QTableView) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SelectedIndexes
func miqt_exec_callback_QTableView_SelectedIndexes(self QTableView, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QTableView) callVirtualBase_UpdateGeometries() {

	QTableView_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QTableView) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_UpdateGeometries
func miqt_exec_callback_QTableView_UpdateGeometries(self QTableView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTableView{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QTableView) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QTableView_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTableView) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_ViewportSizeHint
func miqt_exec_callback_QTableView_ViewportSizeHint(self QTableView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}

func (this *QTableView) callVirtualBase_SizeHintForRow(row int) int {

	return (int)(QTableView_virtualbase_SizeHintForRow(unsafe.Pointer(this.h), (int)(row)))

}
func (this *QTableView) OnSizeHintForRow(slot func(super func(row int) int, row int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SizeHintForRow(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SizeHintForRow
func miqt_exec_callback_QTableView_SizeHintForRow(self QTableView, cb intptr_t, row int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int) int, row int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_SizeHintForRow, slotval1)

	return (int)(virtualReturn)

}

func (this *QTableView) callVirtualBase_SizeHintForColumn(column int) int {

	return (int)(QTableView_virtualbase_SizeHintForColumn(unsafe.Pointer(this.h), (int)(column)))

}
func (this *QTableView) OnSizeHintForColumn(slot func(super func(column int) int, column int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SizeHintForColumn(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SizeHintForColumn
func miqt_exec_callback_QTableView_SizeHintForColumn(self QTableView, cb intptr_t, column int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int) int, column int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_SizeHintForColumn, slotval1)

	return (int)(virtualReturn)

}

func (this *QTableView) callVirtualBase_VerticalScrollbarAction(action int) {

	QTableView_virtualbase_VerticalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QTableView) OnVerticalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_VerticalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_VerticalScrollbarAction
func miqt_exec_callback_QTableView_VerticalScrollbarAction(self QTableView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QTableView{h: self}).callVirtualBase_VerticalScrollbarAction, slotval1)

}

func (this *QTableView) callVirtualBase_HorizontalScrollbarAction(action int) {

	QTableView_virtualbase_HorizontalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QTableView) OnHorizontalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_HorizontalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_HorizontalScrollbarAction
func miqt_exec_callback_QTableView_HorizontalScrollbarAction(self QTableView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QTableView{h: self}).callVirtualBase_HorizontalScrollbarAction, slotval1)

}

func (this *QTableView) callVirtualBase_IsIndexHidden(index *QModelIndex) bool {

	return (bool)(QTableView_virtualbase_IsIndexHidden(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QTableView) OnIsIndexHidden(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_IsIndexHidden
func miqt_exec_callback_QTableView_IsIndexHidden(self QTableView, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_IsIndexHidden, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTableView) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QTableView_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QTableView) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SelectionChanged
func miqt_exec_callback_QTableView_SelectionChanged(self QTableView, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QTableView{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QTableView) callVirtualBase_CurrentChanged(current *QModelIndex, previous *QModelIndex) {

	QTableView_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), previous.cPointer())

}
func (this *QTableView) OnCurrentChanged(slot func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_CurrentChanged
func miqt_exec_callback_QTableView_CurrentChanged(self QTableView, cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc((&QTableView{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}

func (this *QTableView) callVirtualBase_KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))

	QTableView_virtualbase_KeyboardSearch(unsafe.Pointer(this.h), search_ms)

}
func (this *QTableView) OnKeyboardSearch(slot func(super func(search string), search string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_KeyboardSearch(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_KeyboardSearch
func miqt_exec_callback_QTableView_KeyboardSearch(self QTableView, cb intptr_t, search struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(search string), search string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var search_ms struct_miqt_string = search
	search_ret := GoStringN(search_ms.data, int(int64(search_ms.len)))
	free(unsafe.Pointer(search_ms.data))
	slotval1 := search_ret

	gofunc((&QTableView{h: self}).callVirtualBase_KeyboardSearch, slotval1)

}

func (this *QTableView) callVirtualBase_ItemDelegateForIndex(index *QModelIndex) *QAbstractItemDelegate {

	return newQAbstractItemDelegate(QTableView_virtualbase_ItemDelegateForIndex(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QTableView) OnItemDelegateForIndex(slot func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_ItemDelegateForIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_ItemDelegateForIndex
func miqt_exec_callback_QTableView_ItemDelegateForIndex(self QTableView, cb intptr_t, index *QModelIndex) *QAbstractItemDelegate {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_ItemDelegateForIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTableView) callVirtualBase_InputMethodQuery(query InputMethodQuery) *QVariant {

	_goptr := newQVariant(QTableView_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTableView) OnInputMethodQuery(slot func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_InputMethodQuery
func miqt_exec_callback_QTableView_InputMethodQuery(self QTableView, cb intptr_t, query int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(query)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTableView) callVirtualBase_Reset() {

	QTableView_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QTableView) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_Reset
func miqt_exec_callback_QTableView_Reset(self QTableView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTableView{h: self}).callVirtualBase_Reset)

}

func (this *QTableView) callVirtualBase_SelectAll() {

	QTableView_virtualbase_SelectAll(unsafe.Pointer(this.h))

}
func (this *QTableView) OnSelectAll(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SelectAll(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SelectAll
func miqt_exec_callback_QTableView_SelectAll(self QTableView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTableView{h: self}).callVirtualBase_SelectAll)

}

func (this *QTableView) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QTableView_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QTableView) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_DataChanged
func miqt_exec_callback_QTableView_DataChanged(self QTableView, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(topLeft)

	slotval2 := newQModelIndex(bottomRight)

	var roles_ma struct_miqt_array = roles
	roles_ret := make([]int, int(roles_ma.len))
	roles_outCast := (*[0xffff]int)(unsafe.Pointer(roles_ma.data)) // hey ya
	for i := 0; i < int(roles_ma.len); i++ {
		roles_ret[i] = (int)(roles_outCast[i])
	}
	slotval3 := roles_ret

	gofunc((&QTableView{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QTableView) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QTableView_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QTableView) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_RowsInserted
func miqt_exec_callback_QTableView_RowsInserted(self QTableView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QTableView{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QTableView) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QTableView_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QTableView) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_RowsAboutToBeRemoved
func miqt_exec_callback_QTableView_RowsAboutToBeRemoved(self QTableView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QTableView{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QTableView) callVirtualBase_UpdateEditorData() {

	QTableView_virtualbase_UpdateEditorData(unsafe.Pointer(this.h))

}
func (this *QTableView) OnUpdateEditorData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_UpdateEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_UpdateEditorData
func miqt_exec_callback_QTableView_UpdateEditorData(self QTableView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTableView{h: self}).callVirtualBase_UpdateEditorData)

}

func (this *QTableView) callVirtualBase_UpdateEditorGeometries() {

	QTableView_virtualbase_UpdateEditorGeometries(unsafe.Pointer(this.h))

}
func (this *QTableView) OnUpdateEditorGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_UpdateEditorGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_UpdateEditorGeometries
func miqt_exec_callback_QTableView_UpdateEditorGeometries(self QTableView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTableView{h: self}).callVirtualBase_UpdateEditorGeometries)

}

func (this *QTableView) callVirtualBase_VerticalScrollbarValueChanged(value int) {

	QTableView_virtualbase_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QTableView) OnVerticalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_VerticalScrollbarValueChanged
func miqt_exec_callback_QTableView_VerticalScrollbarValueChanged(self QTableView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QTableView{h: self}).callVirtualBase_VerticalScrollbarValueChanged, slotval1)

}

func (this *QTableView) callVirtualBase_HorizontalScrollbarValueChanged(value int) {

	QTableView_virtualbase_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QTableView) OnHorizontalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_HorizontalScrollbarValueChanged
func miqt_exec_callback_QTableView_HorizontalScrollbarValueChanged(self QTableView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QTableView{h: self}).callVirtualBase_HorizontalScrollbarValueChanged, slotval1)

}

func (this *QTableView) callVirtualBase_CloseEditor(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {

	QTableView_virtualbase_CloseEditor(unsafe.Pointer(this.h), editor.cPointer(), (int)(hint))

}
func (this *QTableView) OnCloseEditor(slot func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_CloseEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_CloseEditor
func miqt_exec_callback_QTableView_CloseEditor(self QTableView, cb intptr_t, editor *QWidget, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := (QAbstractItemDelegate__EndEditHint)(hint)

	gofunc((&QTableView{h: self}).callVirtualBase_CloseEditor, slotval1, slotval2)

}

func (this *QTableView) callVirtualBase_CommitData(editor *QWidget) {

	QTableView_virtualbase_CommitData(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QTableView) OnCommitData(slot func(super func(editor *QWidget), editor *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_CommitData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_CommitData
func miqt_exec_callback_QTableView_CommitData(self QTableView, cb intptr_t, editor *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget), editor *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	gofunc((&QTableView{h: self}).callVirtualBase_CommitData, slotval1)

}

func (this *QTableView) callVirtualBase_EditorDestroyed(editor *QObject) {

	QTableView_virtualbase_EditorDestroyed(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QTableView) OnEditorDestroyed(slot func(super func(editor *QObject), editor *QObject)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_EditorDestroyed(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_EditorDestroyed
func miqt_exec_callback_QTableView_EditorDestroyed(self QTableView, cb intptr_t, editor *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QObject), editor *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(editor)

	gofunc((&QTableView{h: self}).callVirtualBase_EditorDestroyed, slotval1)

}

func (this *QTableView) callVirtualBase_Edit2(index *QModelIndex, trigger EditTrigger, event *QEvent) bool {

	return (bool)(QTableView_virtualbase_Edit2(unsafe.Pointer(this.h), index.cPointer(), trigger, event.cPointer()))

}
func (this *QTableView) OnEdit2(slot func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_Edit2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_Edit2
func miqt_exec_callback_QTableView_Edit2(self QTableView, cb intptr_t, index *QModelIndex, trigger EditTrigger, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx
	slotval3 := newQEvent(event)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_Edit2, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QTableView) callVirtualBase_SelectionCommand(index *QModelIndex, event *QEvent) SelectionFlag {

	return (SelectionFlag)(QTableView_virtualbase_SelectionCommand(unsafe.Pointer(this.h), index.cPointer(), event.cPointer()))

}
func (this *QTableView) OnSelectionCommand(slot func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_SelectionCommand(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_SelectionCommand
func miqt_exec_callback_QTableView_SelectionCommand(self QTableView, cb intptr_t, index *QModelIndex, event *QEvent) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_SelectionCommand, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QTableView) callVirtualBase_StartDrag(supportedActions DropAction) {

	QTableView_virtualbase_StartDrag(unsafe.Pointer(this.h), (int)(supportedActions))

}
func (this *QTableView) OnStartDrag(slot func(super func(supportedActions DropAction), supportedActions DropAction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_StartDrag(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_StartDrag
func miqt_exec_callback_QTableView_StartDrag(self QTableView, cb intptr_t, supportedActions int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(supportedActions DropAction), supportedActions DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(supportedActions)

	gofunc((&QTableView{h: self}).callVirtualBase_StartDrag, slotval1)

}

func (this *QTableView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QTableView_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QTableView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_FocusNextPrevChild
func miqt_exec_callback_QTableView_FocusNextPrevChild(self QTableView, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTableView) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QTableView_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTableView) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_Event
func miqt_exec_callback_QTableView_Event(self QTableView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTableView) callVirtualBase_ViewportEvent(event *QEvent) bool {

	return (bool)(QTableView_virtualbase_ViewportEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTableView) OnViewportEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_ViewportEvent
func miqt_exec_callback_QTableView_ViewportEvent(self QTableView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTableView) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QTableView_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_MousePressEvent
func miqt_exec_callback_QTableView_MousePressEvent(self QTableView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTableView) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QTableView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_MouseMoveEvent
func miqt_exec_callback_QTableView_MouseMoveEvent(self QTableView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTableView) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QTableView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_MouseReleaseEvent
func miqt_exec_callback_QTableView_MouseReleaseEvent(self QTableView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTableView) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QTableView_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_MouseDoubleClickEvent
func miqt_exec_callback_QTableView_MouseDoubleClickEvent(self QTableView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTableView) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QTableView_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_DragEnterEvent
func miqt_exec_callback_QTableView_DragEnterEvent(self QTableView, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QTableView) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QTableView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_DragMoveEvent
func miqt_exec_callback_QTableView_DragMoveEvent(self QTableView, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTableView) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QTableView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_DragLeaveEvent
func miqt_exec_callback_QTableView_DragLeaveEvent(self QTableView, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QTableView) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QTableView_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_FocusInEvent
func miqt_exec_callback_QTableView_FocusInEvent(self QTableView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTableView) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QTableView_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_FocusOutEvent
func miqt_exec_callback_QTableView_FocusOutEvent(self QTableView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QTableView) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QTableView_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_KeyPressEvent
func miqt_exec_callback_QTableView_KeyPressEvent(self QTableView, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTableView) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QTableView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_ResizeEvent
func miqt_exec_callback_QTableView_ResizeEvent(self QTableView, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QTableView) callVirtualBase_InputMethodEvent(event *QInputMethodEvent) {

	QTableView_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTableView) OnInputMethodEvent(slot func(super func(event *QInputMethodEvent), event *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_InputMethodEvent
func miqt_exec_callback_QTableView_InputMethodEvent(self QTableView, cb intptr_t, event *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QInputMethodEvent), event *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(event)

	gofunc((&QTableView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QTableView) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QTableView_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QTableView) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableView_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableView_EventFilter
func miqt_exec_callback_QTableView_EventFilter(self QTableView, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTableView{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
