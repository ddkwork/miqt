package qt

import (
	"unsafe"
)

type QTreeView struct {
	h          uintptr
	isSubclass bool
}

// NewQTreeView constructs a new QTreeView object.
func NewQTreeView(parent *QWidget) *QTreeView {

	ret := newQTreeView(QTreeView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeView2 constructs a new QTreeView object.
func NewQTreeView2() *QTreeView {

	ret := newQTreeView(QTreeView_new2())
	ret.isSubclass = true
	return ret
}

func (this *QTreeView) MetaObject() *QMetaObject {
	return newQMetaObject(QTreeView_MetaObject(this.h))
}

func (this *QTreeView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTreeView_Metacast(this.h, param1_Cstring))
}

func QTreeView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTreeView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeView) SetModel(model *QAbstractItemModel) {
	QTreeView_SetModel(this.h, model.cPointer())
}

func (this *QTreeView) SetRootIndex(index *QModelIndex) {
	QTreeView_SetRootIndex(this.h, index.cPointer())
}

func (this *QTreeView) SetSelectionModel(selectionModel *QItemSelectionModel) {
	QTreeView_SetSelectionModel(this.h, selectionModel.cPointer())
}

func (this *QTreeView) Header() *QHeaderView {
	return newQHeaderView(QTreeView_Header(this.h))
}

func (this *QTreeView) SetHeader(header *QHeaderView) {
	QTreeView_SetHeader(this.h, header.cPointer())
}

func (this *QTreeView) AutoExpandDelay() int {
	return (int)(QTreeView_AutoExpandDelay(this.h))
}

func (this *QTreeView) SetAutoExpandDelay(delay int) {
	QTreeView_SetAutoExpandDelay(this.h, (int)(delay))
}

func (this *QTreeView) Indentation() int {
	return (int)(QTreeView_Indentation(this.h))
}

func (this *QTreeView) SetIndentation(i int) {
	QTreeView_SetIndentation(this.h, (int)(i))
}

func (this *QTreeView) ResetIndentation() {
	QTreeView_ResetIndentation(this.h)
}

func (this *QTreeView) RootIsDecorated() bool {
	return (bool)(QTreeView_RootIsDecorated(this.h))
}

func (this *QTreeView) SetRootIsDecorated(show bool) {
	QTreeView_SetRootIsDecorated(this.h, (bool)(show))
}

func (this *QTreeView) UniformRowHeights() bool {
	return (bool)(QTreeView_UniformRowHeights(this.h))
}

func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	QTreeView_SetUniformRowHeights(this.h, (bool)(uniform))
}

func (this *QTreeView) ItemsExpandable() bool {
	return (bool)(QTreeView_ItemsExpandable(this.h))
}

func (this *QTreeView) SetItemsExpandable(enable bool) {
	QTreeView_SetItemsExpandable(this.h, (bool)(enable))
}

func (this *QTreeView) ExpandsOnDoubleClick() bool {
	return (bool)(QTreeView_ExpandsOnDoubleClick(this.h))
}

func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	QTreeView_SetExpandsOnDoubleClick(this.h, (bool)(enable))
}

func (this *QTreeView) ColumnViewportPosition(column int) int {
	return (int)(QTreeView_ColumnViewportPosition(this.h, (int)(column)))
}

func (this *QTreeView) ColumnWidth(column int) int {
	return (int)(QTreeView_ColumnWidth(this.h, (int)(column)))
}

func (this *QTreeView) SetColumnWidth(column int, width int) {
	QTreeView_SetColumnWidth(this.h, (int)(column), (int)(width))
}

func (this *QTreeView) ColumnAt(x int) int {
	return (int)(QTreeView_ColumnAt(this.h, (int)(x)))
}

func (this *QTreeView) IsColumnHidden(column int) bool {
	return (bool)(QTreeView_IsColumnHidden(this.h, (int)(column)))
}

func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	QTreeView_SetColumnHidden(this.h, (int)(column), (bool)(hide))
}

func (this *QTreeView) IsHeaderHidden() bool {
	return (bool)(QTreeView_IsHeaderHidden(this.h))
}

func (this *QTreeView) SetHeaderHidden(hide bool) {
	QTreeView_SetHeaderHidden(this.h, (bool)(hide))
}

func (this *QTreeView) IsRowHidden(row int, parent *QModelIndex) bool {
	return (bool)(QTreeView_IsRowHidden(this.h, (int)(row), parent.cPointer()))
}

func (this *QTreeView) SetRowHidden(row int, parent *QModelIndex, hide bool) {
	QTreeView_SetRowHidden(this.h, (int)(row), parent.cPointer(), (bool)(hide))
}

func (this *QTreeView) IsFirstColumnSpanned(row int, parent *QModelIndex) bool {
	return (bool)(QTreeView_IsFirstColumnSpanned(this.h, (int)(row), parent.cPointer()))
}

func (this *QTreeView) SetFirstColumnSpanned(row int, parent *QModelIndex, span bool) {
	QTreeView_SetFirstColumnSpanned(this.h, (int)(row), parent.cPointer(), (bool)(span))
}

func (this *QTreeView) IsExpanded(index *QModelIndex) bool {
	return (bool)(QTreeView_IsExpanded(this.h, index.cPointer()))
}

func (this *QTreeView) SetExpanded(index *QModelIndex, expand bool) {
	QTreeView_SetExpanded(this.h, index.cPointer(), (bool)(expand))
}

func (this *QTreeView) SetSortingEnabled(enable bool) {
	QTreeView_SetSortingEnabled(this.h, (bool)(enable))
}

func (this *QTreeView) IsSortingEnabled() bool {
	return (bool)(QTreeView_IsSortingEnabled(this.h))
}

func (this *QTreeView) SetAnimated(enable bool) {
	QTreeView_SetAnimated(this.h, (bool)(enable))
}

func (this *QTreeView) IsAnimated() bool {
	return (bool)(QTreeView_IsAnimated(this.h))
}

func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	QTreeView_SetAllColumnsShowFocus(this.h, (bool)(enable))
}

func (this *QTreeView) AllColumnsShowFocus() bool {
	return (bool)(QTreeView_AllColumnsShowFocus(this.h))
}

func (this *QTreeView) SetWordWrap(on bool) {
	QTreeView_SetWordWrap(this.h, (bool)(on))
}

func (this *QTreeView) WordWrap() bool {
	return (bool)(QTreeView_WordWrap(this.h))
}

func (this *QTreeView) SetTreePosition(logicalIndex int) {
	QTreeView_SetTreePosition(this.h, (int)(logicalIndex))
}

func (this *QTreeView) TreePosition() int {
	return (int)(QTreeView_TreePosition(this.h))
}

func (this *QTreeView) KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))
	QTreeView_KeyboardSearch(this.h, search_ms)
}

func (this *QTreeView) VisualRect(index *QModelIndex) *QRect {
	_goptr := newQRect(QTreeView_VisualRect(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeView) ScrollTo(index *QModelIndex, hint ScrollHint) {
	QTreeView_ScrollTo(this.h, index.cPointer(), hint)
}

func (this *QTreeView) IndexAt(p *QPoint) *QModelIndex {
	_goptr := newQModelIndex(QTreeView_IndexAt(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeView) IndexAbove(index *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QTreeView_IndexAbove(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeView) IndexBelow(index *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QTreeView_IndexBelow(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeView) DoItemsLayout() {
	QTreeView_DoItemsLayout(this.h)
}

func (this *QTreeView) Reset() {
	QTreeView_Reset(this.h)
}

func (this *QTreeView) DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}
	QTreeView_DataChanged(this.h, topLeft.cPointer(), bottomRight.cPointer(), roles_ma)
}

func (this *QTreeView) SelectAll() {
	QTreeView_SelectAll(this.h)
}

func (this *QTreeView) Expanded(index *QModelIndex) {
	QTreeView_Expanded(this.h, index.cPointer())
}
func (this *QTreeView) OnExpanded(slot func(index *QModelIndex)) {
	QTreeView_connect_Expanded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_Expanded
func miqt_exec_callback_QTreeView_Expanded(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QTreeView) Collapsed(index *QModelIndex) {
	QTreeView_Collapsed(this.h, index.cPointer())
}
func (this *QTreeView) OnCollapsed(slot func(index *QModelIndex)) {
	QTreeView_connect_Collapsed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_Collapsed
func miqt_exec_callback_QTreeView_Collapsed(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QTreeView) HideColumn(column int) {
	QTreeView_HideColumn(this.h, (int)(column))
}

func (this *QTreeView) ShowColumn(column int) {
	QTreeView_ShowColumn(this.h, (int)(column))
}

func (this *QTreeView) Expand(index *QModelIndex) {
	QTreeView_Expand(this.h, index.cPointer())
}

func (this *QTreeView) Collapse(index *QModelIndex) {
	QTreeView_Collapse(this.h, index.cPointer())
}

func (this *QTreeView) ResizeColumnToContents(column int) {
	QTreeView_ResizeColumnToContents(this.h, (int)(column))
}

func (this *QTreeView) SortByColumn(column int, order SortOrder) {
	QTreeView_SortByColumn(this.h, (int)(column), (int)(order))
}

func (this *QTreeView) ExpandAll() {
	QTreeView_ExpandAll(this.h)
}

func (this *QTreeView) ExpandRecursively(index *QModelIndex) {
	QTreeView_ExpandRecursively(this.h, index.cPointer())
}

func (this *QTreeView) CollapseAll() {
	QTreeView_CollapseAll(this.h)
}

func (this *QTreeView) ExpandToDepth(depth int) {
	QTreeView_ExpandToDepth(this.h, (int)(depth))
}

func QTreeView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTreeView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTreeView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTreeView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeView) ExpandRecursively2(index *QModelIndex, depth int) {
	QTreeView_ExpandRecursively2(this.h, index.cPointer(), (int)(depth))
}

func (this *QTreeView) callVirtualBase_SetModel(model *QAbstractItemModel) {

	QTreeView_virtualbase_SetModel(unsafe.Pointer(this.h), model.cPointer())

}
func (this *QTreeView) OnSetModel(slot func(super func(model *QAbstractItemModel), model *QAbstractItemModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SetModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SetModel
func miqt_exec_callback_QTreeView_SetModel(self QTreeView, cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(model *QAbstractItemModel), model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc((&QTreeView{h: self}).callVirtualBase_SetModel, slotval1)

}

func (this *QTreeView) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QTreeView_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QTreeView) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SetRootIndex
func miqt_exec_callback_QTreeView_SetRootIndex(self QTreeView, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QTreeView{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QTreeView) callVirtualBase_SetSelectionModel(selectionModel *QItemSelectionModel) {

	QTreeView_virtualbase_SetSelectionModel(unsafe.Pointer(this.h), selectionModel.cPointer())

}
func (this *QTreeView) OnSetSelectionModel(slot func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SetSelectionModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SetSelectionModel
func miqt_exec_callback_QTreeView_SetSelectionModel(self QTreeView, cb intptr_t, selectionModel *QItemSelectionModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelectionModel(selectionModel)

	gofunc((&QTreeView{h: self}).callVirtualBase_SetSelectionModel, slotval1)

}

func (this *QTreeView) callVirtualBase_KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))

	QTreeView_virtualbase_KeyboardSearch(unsafe.Pointer(this.h), search_ms)

}
func (this *QTreeView) OnKeyboardSearch(slot func(super func(search string), search string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_KeyboardSearch(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_KeyboardSearch
func miqt_exec_callback_QTreeView_KeyboardSearch(self QTreeView, cb intptr_t, search struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(search string), search string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var search_ms struct_miqt_string = search
	search_ret := GoStringN(search_ms.data, int(int64(search_ms.len)))
	free(unsafe.Pointer(search_ms.data))
	slotval1 := search_ret

	gofunc((&QTreeView{h: self}).callVirtualBase_KeyboardSearch, slotval1)

}

func (this *QTreeView) callVirtualBase_VisualRect(index *QModelIndex) *QRect {

	_goptr := newQRect(QTreeView_virtualbase_VisualRect(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeView) OnVisualRect(slot func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_VisualRect
func miqt_exec_callback_QTreeView_VisualRect(self QTreeView, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_VisualRect, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeView) callVirtualBase_ScrollTo(index *QModelIndex, hint ScrollHint) {

	QTreeView_virtualbase_ScrollTo(unsafe.Pointer(this.h), index.cPointer(), hint)

}
func (this *QTreeView) OnScrollTo(slot func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_ScrollTo
func miqt_exec_callback_QTreeView_ScrollTo(self QTreeView, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc((&QTreeView{h: self}).callVirtualBase_ScrollTo, slotval1, slotval2)

}

func (this *QTreeView) callVirtualBase_IndexAt(p *QPoint) *QModelIndex {

	_goptr := newQModelIndex(QTreeView_virtualbase_IndexAt(unsafe.Pointer(this.h), p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeView) OnIndexAt(slot func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_IndexAt
func miqt_exec_callback_QTreeView_IndexAt(self QTreeView, cb intptr_t, p *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(p)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_IndexAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeView) callVirtualBase_DoItemsLayout() {

	QTreeView_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QTreeView) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DoItemsLayout
func miqt_exec_callback_QTreeView_DoItemsLayout(self QTreeView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeView{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QTreeView) callVirtualBase_Reset() {

	QTreeView_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QTreeView) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_Reset
func miqt_exec_callback_QTreeView_Reset(self QTreeView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeView{h: self}).callVirtualBase_Reset)

}

func (this *QTreeView) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QTreeView_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QTreeView) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DataChanged
func miqt_exec_callback_QTreeView_DataChanged(self QTreeView, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
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

	gofunc((&QTreeView{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QTreeView) callVirtualBase_SelectAll() {

	QTreeView_virtualbase_SelectAll(unsafe.Pointer(this.h))

}
func (this *QTreeView) OnSelectAll(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SelectAll(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SelectAll
func miqt_exec_callback_QTreeView_SelectAll(self QTreeView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeView{h: self}).callVirtualBase_SelectAll)

}

func (this *QTreeView) callVirtualBase_VerticalScrollbarValueChanged(value int) {

	QTreeView_virtualbase_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QTreeView) OnVerticalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_VerticalScrollbarValueChanged
func miqt_exec_callback_QTreeView_VerticalScrollbarValueChanged(self QTreeView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QTreeView{h: self}).callVirtualBase_VerticalScrollbarValueChanged, slotval1)

}

func (this *QTreeView) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QTreeView_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QTreeView) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_ScrollContentsBy
func miqt_exec_callback_QTreeView_ScrollContentsBy(self QTreeView, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QTreeView{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QTreeView) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QTreeView_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QTreeView) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_RowsInserted
func miqt_exec_callback_QTreeView_RowsInserted(self QTreeView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QTreeView{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QTreeView) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QTreeView_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QTreeView) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_RowsAboutToBeRemoved
func miqt_exec_callback_QTreeView_RowsAboutToBeRemoved(self QTreeView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QTreeView{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QTreeView) callVirtualBase_MoveCursor(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex {

	_goptr := newQModelIndex(QTreeView_virtualbase_MoveCursor(unsafe.Pointer(this.h), cursorAction, (int)(modifiers)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeView) OnMoveCursor(slot func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_MoveCursor
func miqt_exec_callback_QTreeView_MoveCursor(self QTreeView, cb intptr_t, cursorAction CursorAction, modifiers int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(modifiers)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_MoveCursor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QTreeView) callVirtualBase_HorizontalOffset() int {

	return (int)(QTreeView_virtualbase_HorizontalOffset(unsafe.Pointer(this.h)))

}
func (this *QTreeView) OnHorizontalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_HorizontalOffset
func miqt_exec_callback_QTreeView_HorizontalOffset(self QTreeView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_HorizontalOffset)

	return (int)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_VerticalOffset() int {

	return (int)(QTreeView_virtualbase_VerticalOffset(unsafe.Pointer(this.h)))

}
func (this *QTreeView) OnVerticalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_VerticalOffset
func miqt_exec_callback_QTreeView_VerticalOffset(self QTreeView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_VerticalOffset)

	return (int)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_SetSelection(rect *QRect, command SelectionFlag) {

	QTreeView_virtualbase_SetSelection(unsafe.Pointer(this.h), rect.cPointer(), (int)(command))

}
func (this *QTreeView) OnSetSelection(slot func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SetSelection
func miqt_exec_callback_QTreeView_SetSelection(self QTreeView, cb intptr_t, rect *QRect, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QTreeView{h: self}).callVirtualBase_SetSelection, slotval1, slotval2)

}

func (this *QTreeView) callVirtualBase_VisualRegionForSelection(selection *QItemSelection) *QRegion {

	_goptr := newQRegion(QTreeView_virtualbase_VisualRegionForSelection(unsafe.Pointer(this.h), selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeView) OnVisualRegionForSelection(slot func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_VisualRegionForSelection
func miqt_exec_callback_QTreeView_VisualRegionForSelection(self QTreeView, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_VisualRegionForSelection, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeView) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QTreeView_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QTreeView) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SelectedIndexes
func miqt_exec_callback_QTreeView_SelectedIndexes(self QTreeView, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QTreeView) callVirtualBase_ChangeEvent(event *QEvent) {

	QTreeView_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_ChangeEvent
func miqt_exec_callback_QTreeView_ChangeEvent(self QTreeView, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QTreeView_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_TimerEvent
func miqt_exec_callback_QTreeView_TimerEvent(self QTreeView, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QTreeView_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_PaintEvent
func miqt_exec_callback_QTreeView_PaintEvent(self QTreeView, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_DrawRow(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex) {

	QTreeView_virtualbase_DrawRow(unsafe.Pointer(this.h), painter.cPointer(), options.cPointer(), index.cPointer())

}
func (this *QTreeView) OnDrawRow(slot func(super func(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DrawRow(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DrawRow
func miqt_exec_callback_QTreeView_DrawRow(self QTreeView, cb intptr_t, painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(options)

	slotval3 := newQModelIndex(index)

	gofunc((&QTreeView{h: self}).callVirtualBase_DrawRow, slotval1, slotval2, slotval3)

}

func (this *QTreeView) callVirtualBase_DrawBranches(painter *QPainter, rect *QRect, index *QModelIndex) {

	QTreeView_virtualbase_DrawBranches(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), index.cPointer())

}
func (this *QTreeView) OnDrawBranches(slot func(super func(painter *QPainter, rect *QRect, index *QModelIndex), painter *QPainter, rect *QRect, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DrawBranches(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DrawBranches
func miqt_exec_callback_QTreeView_DrawBranches(self QTreeView, cb intptr_t, painter *QPainter, rect *QRect, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, index *QModelIndex), painter *QPainter, rect *QRect, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	slotval3 := newQModelIndex(index)

	gofunc((&QTreeView{h: self}).callVirtualBase_DrawBranches, slotval1, slotval2, slotval3)

}

func (this *QTreeView) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QTreeView_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_MousePressEvent
func miqt_exec_callback_QTreeView_MousePressEvent(self QTreeView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QTreeView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_MouseReleaseEvent
func miqt_exec_callback_QTreeView_MouseReleaseEvent(self QTreeView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QTreeView_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_MouseDoubleClickEvent
func miqt_exec_callback_QTreeView_MouseDoubleClickEvent(self QTreeView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QTreeView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_MouseMoveEvent
func miqt_exec_callback_QTreeView_MouseMoveEvent(self QTreeView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QTreeView_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_KeyPressEvent
func miqt_exec_callback_QTreeView_KeyPressEvent(self QTreeView, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QTreeView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DragMoveEvent
func miqt_exec_callback_QTreeView_DragMoveEvent(self QTreeView, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_ViewportEvent(event *QEvent) bool {

	return (bool)(QTreeView_virtualbase_ViewportEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTreeView) OnViewportEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_ViewportEvent
func miqt_exec_callback_QTreeView_ViewportEvent(self QTreeView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_UpdateGeometries() {

	QTreeView_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QTreeView) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_UpdateGeometries
func miqt_exec_callback_QTreeView_UpdateGeometries(self QTreeView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeView{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QTreeView) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QTreeView_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeView) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_ViewportSizeHint
func miqt_exec_callback_QTreeView_ViewportSizeHint(self QTreeView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}

func (this *QTreeView) callVirtualBase_SizeHintForColumn(column int) int {

	return (int)(QTreeView_virtualbase_SizeHintForColumn(unsafe.Pointer(this.h), (int)(column)))

}
func (this *QTreeView) OnSizeHintForColumn(slot func(super func(column int) int, column int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SizeHintForColumn(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SizeHintForColumn
func miqt_exec_callback_QTreeView_SizeHintForColumn(self QTreeView, cb intptr_t, column int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int) int, column int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_SizeHintForColumn, slotval1)

	return (int)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_HorizontalScrollbarAction(action int) {

	QTreeView_virtualbase_HorizontalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QTreeView) OnHorizontalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_HorizontalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_HorizontalScrollbarAction
func miqt_exec_callback_QTreeView_HorizontalScrollbarAction(self QTreeView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QTreeView{h: self}).callVirtualBase_HorizontalScrollbarAction, slotval1)

}

func (this *QTreeView) callVirtualBase_IsIndexHidden(index *QModelIndex) bool {

	return (bool)(QTreeView_virtualbase_IsIndexHidden(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QTreeView) OnIsIndexHidden(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_IsIndexHidden
func miqt_exec_callback_QTreeView_IsIndexHidden(self QTreeView, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_IsIndexHidden, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QTreeView_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QTreeView) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SelectionChanged
func miqt_exec_callback_QTreeView_SelectionChanged(self QTreeView, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QTreeView{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QTreeView) callVirtualBase_CurrentChanged(current *QModelIndex, previous *QModelIndex) {

	QTreeView_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), previous.cPointer())

}
func (this *QTreeView) OnCurrentChanged(slot func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_CurrentChanged
func miqt_exec_callback_QTreeView_CurrentChanged(self QTreeView, cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc((&QTreeView{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}

func (this *QTreeView) callVirtualBase_SizeHintForRow(row int) int {

	return (int)(QTreeView_virtualbase_SizeHintForRow(unsafe.Pointer(this.h), (int)(row)))

}
func (this *QTreeView) OnSizeHintForRow(slot func(super func(row int) int, row int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SizeHintForRow(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SizeHintForRow
func miqt_exec_callback_QTreeView_SizeHintForRow(self QTreeView, cb intptr_t, row int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int) int, row int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_SizeHintForRow, slotval1)

	return (int)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_ItemDelegateForIndex(index *QModelIndex) *QAbstractItemDelegate {

	return newQAbstractItemDelegate(QTreeView_virtualbase_ItemDelegateForIndex(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QTreeView) OnItemDelegateForIndex(slot func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_ItemDelegateForIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_ItemDelegateForIndex
func miqt_exec_callback_QTreeView_ItemDelegateForIndex(self QTreeView, cb intptr_t, index *QModelIndex) *QAbstractItemDelegate {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_ItemDelegateForIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeView) callVirtualBase_InputMethodQuery(query InputMethodQuery) *QVariant {

	_goptr := newQVariant(QTreeView_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeView) OnInputMethodQuery(slot func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_InputMethodQuery
func miqt_exec_callback_QTreeView_InputMethodQuery(self QTreeView, cb intptr_t, query int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(query)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeView) callVirtualBase_UpdateEditorData() {

	QTreeView_virtualbase_UpdateEditorData(unsafe.Pointer(this.h))

}
func (this *QTreeView) OnUpdateEditorData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_UpdateEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_UpdateEditorData
func miqt_exec_callback_QTreeView_UpdateEditorData(self QTreeView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeView{h: self}).callVirtualBase_UpdateEditorData)

}

func (this *QTreeView) callVirtualBase_UpdateEditorGeometries() {

	QTreeView_virtualbase_UpdateEditorGeometries(unsafe.Pointer(this.h))

}
func (this *QTreeView) OnUpdateEditorGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_UpdateEditorGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_UpdateEditorGeometries
func miqt_exec_callback_QTreeView_UpdateEditorGeometries(self QTreeView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeView{h: self}).callVirtualBase_UpdateEditorGeometries)

}

func (this *QTreeView) callVirtualBase_VerticalScrollbarAction(action int) {

	QTreeView_virtualbase_VerticalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QTreeView) OnVerticalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_VerticalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_VerticalScrollbarAction
func miqt_exec_callback_QTreeView_VerticalScrollbarAction(self QTreeView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QTreeView{h: self}).callVirtualBase_VerticalScrollbarAction, slotval1)

}

func (this *QTreeView) callVirtualBase_HorizontalScrollbarValueChanged(value int) {

	QTreeView_virtualbase_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QTreeView) OnHorizontalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_HorizontalScrollbarValueChanged
func miqt_exec_callback_QTreeView_HorizontalScrollbarValueChanged(self QTreeView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QTreeView{h: self}).callVirtualBase_HorizontalScrollbarValueChanged, slotval1)

}

func (this *QTreeView) callVirtualBase_CloseEditor(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {

	QTreeView_virtualbase_CloseEditor(unsafe.Pointer(this.h), editor.cPointer(), (int)(hint))

}
func (this *QTreeView) OnCloseEditor(slot func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_CloseEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_CloseEditor
func miqt_exec_callback_QTreeView_CloseEditor(self QTreeView, cb intptr_t, editor *QWidget, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := (QAbstractItemDelegate__EndEditHint)(hint)

	gofunc((&QTreeView{h: self}).callVirtualBase_CloseEditor, slotval1, slotval2)

}

func (this *QTreeView) callVirtualBase_CommitData(editor *QWidget) {

	QTreeView_virtualbase_CommitData(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QTreeView) OnCommitData(slot func(super func(editor *QWidget), editor *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_CommitData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_CommitData
func miqt_exec_callback_QTreeView_CommitData(self QTreeView, cb intptr_t, editor *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget), editor *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	gofunc((&QTreeView{h: self}).callVirtualBase_CommitData, slotval1)

}

func (this *QTreeView) callVirtualBase_EditorDestroyed(editor *QObject) {

	QTreeView_virtualbase_EditorDestroyed(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QTreeView) OnEditorDestroyed(slot func(super func(editor *QObject), editor *QObject)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_EditorDestroyed(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_EditorDestroyed
func miqt_exec_callback_QTreeView_EditorDestroyed(self QTreeView, cb intptr_t, editor *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QObject), editor *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(editor)

	gofunc((&QTreeView{h: self}).callVirtualBase_EditorDestroyed, slotval1)

}

func (this *QTreeView) callVirtualBase_Edit2(index *QModelIndex, trigger EditTrigger, event *QEvent) bool {

	return (bool)(QTreeView_virtualbase_Edit2(unsafe.Pointer(this.h), index.cPointer(), trigger, event.cPointer()))

}
func (this *QTreeView) OnEdit2(slot func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_Edit2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_Edit2
func miqt_exec_callback_QTreeView_Edit2(self QTreeView, cb intptr_t, index *QModelIndex, trigger EditTrigger, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx
	slotval3 := newQEvent(event)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_Edit2, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_SelectionCommand(index *QModelIndex, event *QEvent) SelectionFlag {

	return (SelectionFlag)(QTreeView_virtualbase_SelectionCommand(unsafe.Pointer(this.h), index.cPointer(), event.cPointer()))

}
func (this *QTreeView) OnSelectionCommand(slot func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_SelectionCommand(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_SelectionCommand
func miqt_exec_callback_QTreeView_SelectionCommand(self QTreeView, cb intptr_t, index *QModelIndex, event *QEvent) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_SelectionCommand, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_StartDrag(supportedActions DropAction) {

	QTreeView_virtualbase_StartDrag(unsafe.Pointer(this.h), (int)(supportedActions))

}
func (this *QTreeView) OnStartDrag(slot func(super func(supportedActions DropAction), supportedActions DropAction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_StartDrag(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_StartDrag
func miqt_exec_callback_QTreeView_StartDrag(self QTreeView, cb intptr_t, supportedActions int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(supportedActions DropAction), supportedActions DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(supportedActions)

	gofunc((&QTreeView{h: self}).callVirtualBase_StartDrag, slotval1)

}

func (this *QTreeView) callVirtualBase_InitViewItemOption(option *QStyleOptionViewItem) {

	QTreeView_virtualbase_InitViewItemOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QTreeView) OnInitViewItemOption(slot func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_InitViewItemOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_InitViewItemOption
func miqt_exec_callback_QTreeView_InitViewItemOption(self QTreeView, cb intptr_t, option *QStyleOptionViewItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	gofunc((&QTreeView{h: self}).callVirtualBase_InitViewItemOption, slotval1)

}

func (this *QTreeView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QTreeView_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QTreeView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_FocusNextPrevChild
func miqt_exec_callback_QTreeView_FocusNextPrevChild(self QTreeView, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QTreeView_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTreeView) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_Event
func miqt_exec_callback_QTreeView_Event(self QTreeView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeView) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QTreeView_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DragEnterEvent
func miqt_exec_callback_QTreeView_DragEnterEvent(self QTreeView, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QTreeView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DragLeaveEvent
func miqt_exec_callback_QTreeView_DragLeaveEvent(self QTreeView, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_DropEvent(event *QDropEvent) {

	QTreeView_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_DropEvent
func miqt_exec_callback_QTreeView_DropEvent(self QTreeView, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QTreeView_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_FocusInEvent
func miqt_exec_callback_QTreeView_FocusInEvent(self QTreeView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QTreeView_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_FocusOutEvent
func miqt_exec_callback_QTreeView_FocusOutEvent(self QTreeView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QTreeView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_ResizeEvent
func miqt_exec_callback_QTreeView_ResizeEvent(self QTreeView, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_InputMethodEvent(event *QInputMethodEvent) {

	QTreeView_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeView) OnInputMethodEvent(slot func(super func(event *QInputMethodEvent), event *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_InputMethodEvent
func miqt_exec_callback_QTreeView_InputMethodEvent(self QTreeView, cb intptr_t, event *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QInputMethodEvent), event *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(event)

	gofunc((&QTreeView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QTreeView) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QTreeView_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QTreeView) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_EventFilter
func miqt_exec_callback_QTreeView_EventFilter(self QTreeView, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
