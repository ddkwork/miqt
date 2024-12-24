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

func (this *QTreeView) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTreeView_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTreeView) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_MetaObject
func miqt_exec_callback_QTreeView_MetaObject(self QTreeView, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTreeView) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTreeView_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTreeView) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeView_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeView_Metacast
func miqt_exec_callback_QTreeView_Metacast(self QTreeView, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTreeView{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
