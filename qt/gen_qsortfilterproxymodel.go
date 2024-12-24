package qt

import (
	"unsafe"
)

type QSortFilterProxyModel struct {
	h          uintptr
	isSubclass bool
}

// NewQSortFilterProxyModel constructs a new QSortFilterProxyModel object.
func NewQSortFilterProxyModel() *QSortFilterProxyModel {
	ret := newQSortFilterProxyModel(QSortFilterProxyModel_new())
	ret.isSubclass = true
	return ret
}

// NewQSortFilterProxyModel2 constructs a new QSortFilterProxyModel object.
func NewQSortFilterProxyModel2(parent *QObject) *QSortFilterProxyModel {
	ret := newQSortFilterProxyModel(QSortFilterProxyModel_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSortFilterProxyModel) MetaObject() *QMetaObject {
	return newQMetaObject(QSortFilterProxyModel_MetaObject(this.h))
}

func (this *QSortFilterProxyModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSortFilterProxyModel_Metacast(this.h, param1_Cstring))
}

func QSortFilterProxyModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSortFilterProxyModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSortFilterProxyModel) SetSourceModel(sourceModel *QAbstractItemModel) {
	QSortFilterProxyModel_SetSourceModel(this.h, sourceModel.cPointer())
}

func (this *QSortFilterProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QSortFilterProxyModel_MapToSource(this.h, proxyIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QSortFilterProxyModel_MapFromSource(this.h, sourceIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) MapSelectionToSource(proxySelection *QItemSelection) *QItemSelection {
	_goptr := newQItemSelection(QSortFilterProxyModel_MapSelectionToSource(this.h, proxySelection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) MapSelectionFromSource(sourceSelection *QItemSelection) *QItemSelection {
	_goptr := newQItemSelection(QSortFilterProxyModel_MapSelectionFromSource(this.h, sourceSelection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) FilterRegularExpression() *QRegularExpression {
	_goptr := newQRegularExpression(QSortFilterProxyModel_FilterRegularExpression(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) FilterKeyColumn() int {
	return (int)(QSortFilterProxyModel_FilterKeyColumn(this.h))
}

func (this *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	QSortFilterProxyModel_SetFilterKeyColumn(this.h, (int)(column))
}

func (this *QSortFilterProxyModel) FilterCaseSensitivity() CaseSensitivity {
	return (CaseSensitivity)(QSortFilterProxyModel_FilterCaseSensitivity(this.h))
}

func (this *QSortFilterProxyModel) SetFilterCaseSensitivity(cs CaseSensitivity) {
	QSortFilterProxyModel_SetFilterCaseSensitivity(this.h, (int)(cs))
}

func (this *QSortFilterProxyModel) SortCaseSensitivity() CaseSensitivity {
	return (CaseSensitivity)(QSortFilterProxyModel_SortCaseSensitivity(this.h))
}

func (this *QSortFilterProxyModel) SetSortCaseSensitivity(cs CaseSensitivity) {
	QSortFilterProxyModel_SetSortCaseSensitivity(this.h, (int)(cs))
}

func (this *QSortFilterProxyModel) IsSortLocaleAware() bool {
	return (bool)(QSortFilterProxyModel_IsSortLocaleAware(this.h))
}

func (this *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	QSortFilterProxyModel_SetSortLocaleAware(this.h, (bool)(on))
}

func (this *QSortFilterProxyModel) SortColumn() int {
	return (int)(QSortFilterProxyModel_SortColumn(this.h))
}

func (this *QSortFilterProxyModel) SortOrder() SortOrder {
	return (SortOrder)(QSortFilterProxyModel_SortOrder(this.h))
}

func (this *QSortFilterProxyModel) DynamicSortFilter() bool {
	return (bool)(QSortFilterProxyModel_DynamicSortFilter(this.h))
}

func (this *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	QSortFilterProxyModel_SetDynamicSortFilter(this.h, (bool)(enable))
}

func (this *QSortFilterProxyModel) SortRole() int {
	return (int)(QSortFilterProxyModel_SortRole(this.h))
}

func (this *QSortFilterProxyModel) SetSortRole(role int) {
	QSortFilterProxyModel_SetSortRole(this.h, (int)(role))
}

func (this *QSortFilterProxyModel) FilterRole() int {
	return (int)(QSortFilterProxyModel_FilterRole(this.h))
}

func (this *QSortFilterProxyModel) SetFilterRole(role int) {
	QSortFilterProxyModel_SetFilterRole(this.h, (int)(role))
}

func (this *QSortFilterProxyModel) IsRecursiveFilteringEnabled() bool {
	return (bool)(QSortFilterProxyModel_IsRecursiveFilteringEnabled(this.h))
}

func (this *QSortFilterProxyModel) SetRecursiveFilteringEnabled(recursive bool) {
	QSortFilterProxyModel_SetRecursiveFilteringEnabled(this.h, (bool)(recursive))
}

func (this *QSortFilterProxyModel) AutoAcceptChildRows() bool {
	return (bool)(QSortFilterProxyModel_AutoAcceptChildRows(this.h))
}

func (this *QSortFilterProxyModel) SetAutoAcceptChildRows(accept bool) {
	QSortFilterProxyModel_SetAutoAcceptChildRows(this.h, (bool)(accept))
}

func (this *QSortFilterProxyModel) SetFilterRegularExpression(pattern string) {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))
	QSortFilterProxyModel_SetFilterRegularExpression(this.h, pattern_ms)
}

func (this *QSortFilterProxyModel) SetFilterRegularExpressionWithRegularExpression(regularExpression *QRegularExpression) {
	QSortFilterProxyModel_SetFilterRegularExpressionWithRegularExpression(this.h, regularExpression.cPointer())
}

func (this *QSortFilterProxyModel) SetFilterWildcard(pattern string) {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))
	QSortFilterProxyModel_SetFilterWildcard(this.h, pattern_ms)
}

func (this *QSortFilterProxyModel) SetFilterFixedString(pattern string) {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))
	QSortFilterProxyModel_SetFilterFixedString(this.h, pattern_ms)
}

func (this *QSortFilterProxyModel) Invalidate() {
	QSortFilterProxyModel_Invalidate(this.h)
}

func (this *QSortFilterProxyModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QSortFilterProxyModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) Parent(child *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QSortFilterProxyModel_Parent(this.h, child.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QSortFilterProxyModel_Sibling(this.h, (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) RowCount(parent *QModelIndex) int {
	return (int)(QSortFilterProxyModel_RowCount(this.h, parent.cPointer()))
}

func (this *QSortFilterProxyModel) ColumnCount(parent *QModelIndex) int {
	return (int)(QSortFilterProxyModel_ColumnCount(this.h, parent.cPointer()))
}

func (this *QSortFilterProxyModel) HasChildren(parent *QModelIndex) bool {
	return (bool)(QSortFilterProxyModel_HasChildren(this.h, parent.cPointer()))
}

func (this *QSortFilterProxyModel) Data(index *QModelIndex, role int) *QVariant {
	_goptr := newQVariant(QSortFilterProxyModel_Data(this.h, index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	return (bool)(QSortFilterProxyModel_SetData(this.h, index.cPointer(), value.cPointer(), (int)(role)))
}

func (this *QSortFilterProxyModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	_goptr := newQVariant(QSortFilterProxyModel_HeaderData(this.h, (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {
	return (bool)(QSortFilterProxyModel_SetHeaderData(this.h, (int)(section), (int)(orientation), value.cPointer(), (int)(role)))
}

func (this *QSortFilterProxyModel) MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}
	return newQMimeData(QSortFilterProxyModel_MimeData(this.h, indexes_ma))
}

func (this *QSortFilterProxyModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QSortFilterProxyModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QSortFilterProxyModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QSortFilterProxyModel_InsertRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QSortFilterProxyModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QSortFilterProxyModel_InsertColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QSortFilterProxyModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QSortFilterProxyModel_RemoveRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QSortFilterProxyModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QSortFilterProxyModel_RemoveColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QSortFilterProxyModel) FetchMore(parent *QModelIndex) {
	QSortFilterProxyModel_FetchMore(this.h, parent.cPointer())
}

func (this *QSortFilterProxyModel) CanFetchMore(parent *QModelIndex) bool {
	return (bool)(QSortFilterProxyModel_CanFetchMore(this.h, parent.cPointer()))
}

func (this *QSortFilterProxyModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QSortFilterProxyModel_Flags(this.h, index.cPointer()))
}

func (this *QSortFilterProxyModel) Buddy(index *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QSortFilterProxyModel_Buddy(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {
	var _ma struct_miqt_array = QSortFilterProxyModel_Match(this.h, start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QSortFilterProxyModel) Span(index *QModelIndex) *QSize {
	_goptr := newQSize(QSortFilterProxyModel_Span(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSortFilterProxyModel) Sort(column int, order SortOrder) {
	QSortFilterProxyModel_Sort(this.h, (int)(column), (int)(order))
}

func (this *QSortFilterProxyModel) MimeTypes() []string {
	var _ma struct_miqt_array = QSortFilterProxyModel_MimeTypes(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QSortFilterProxyModel) SupportedDropActions() DropAction {
	return (DropAction)(QSortFilterProxyModel_SupportedDropActions(this.h))
}

func (this *QSortFilterProxyModel) DynamicSortFilterChanged(dynamicSortFilter bool) {
	QSortFilterProxyModel_DynamicSortFilterChanged(this.h, (bool)(dynamicSortFilter))
}

func (this *QSortFilterProxyModel) OnDynamicSortFilterChanged(slot func(dynamicSortFilter bool)) {
	QSortFilterProxyModel_connect_DynamicSortFilterChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_DynamicSortFilterChanged
func miqt_exec_callback_QSortFilterProxyModel_DynamicSortFilterChanged(cb intptr_t, dynamicSortFilter bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(dynamicSortFilter bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(dynamicSortFilter)

	gofunc(slotval1)
}

func (this *QSortFilterProxyModel) FilterCaseSensitivityChanged(filterCaseSensitivity CaseSensitivity) {
	QSortFilterProxyModel_FilterCaseSensitivityChanged(this.h, (int)(filterCaseSensitivity))
}

func (this *QSortFilterProxyModel) OnFilterCaseSensitivityChanged(slot func(filterCaseSensitivity CaseSensitivity)) {
	QSortFilterProxyModel_connect_FilterCaseSensitivityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_FilterCaseSensitivityChanged
func miqt_exec_callback_QSortFilterProxyModel_FilterCaseSensitivityChanged(cb intptr_t, filterCaseSensitivity int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(filterCaseSensitivity CaseSensitivity))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (CaseSensitivity)(filterCaseSensitivity)

	gofunc(slotval1)
}

func (this *QSortFilterProxyModel) SortCaseSensitivityChanged(sortCaseSensitivity CaseSensitivity) {
	QSortFilterProxyModel_SortCaseSensitivityChanged(this.h, (int)(sortCaseSensitivity))
}

func (this *QSortFilterProxyModel) OnSortCaseSensitivityChanged(slot func(sortCaseSensitivity CaseSensitivity)) {
	QSortFilterProxyModel_connect_SortCaseSensitivityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_SortCaseSensitivityChanged
func miqt_exec_callback_QSortFilterProxyModel_SortCaseSensitivityChanged(cb intptr_t, sortCaseSensitivity int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(sortCaseSensitivity CaseSensitivity))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (CaseSensitivity)(sortCaseSensitivity)

	gofunc(slotval1)
}

func (this *QSortFilterProxyModel) SortLocaleAwareChanged(sortLocaleAware bool) {
	QSortFilterProxyModel_SortLocaleAwareChanged(this.h, (bool)(sortLocaleAware))
}

func (this *QSortFilterProxyModel) OnSortLocaleAwareChanged(slot func(sortLocaleAware bool)) {
	QSortFilterProxyModel_connect_SortLocaleAwareChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_SortLocaleAwareChanged
func miqt_exec_callback_QSortFilterProxyModel_SortLocaleAwareChanged(cb intptr_t, sortLocaleAware bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(sortLocaleAware bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(sortLocaleAware)

	gofunc(slotval1)
}

func (this *QSortFilterProxyModel) SortRoleChanged(sortRole int) {
	QSortFilterProxyModel_SortRoleChanged(this.h, (int)(sortRole))
}

func (this *QSortFilterProxyModel) OnSortRoleChanged(slot func(sortRole int)) {
	QSortFilterProxyModel_connect_SortRoleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_SortRoleChanged
func miqt_exec_callback_QSortFilterProxyModel_SortRoleChanged(cb intptr_t, sortRole int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(sortRole int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(sortRole)

	gofunc(slotval1)
}

func (this *QSortFilterProxyModel) FilterRoleChanged(filterRole int) {
	QSortFilterProxyModel_FilterRoleChanged(this.h, (int)(filterRole))
}

func (this *QSortFilterProxyModel) OnFilterRoleChanged(slot func(filterRole int)) {
	QSortFilterProxyModel_connect_FilterRoleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_FilterRoleChanged
func miqt_exec_callback_QSortFilterProxyModel_FilterRoleChanged(cb intptr_t, filterRole int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(filterRole int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(filterRole)

	gofunc(slotval1)
}

func (this *QSortFilterProxyModel) RecursiveFilteringEnabledChanged(recursiveFilteringEnabled bool) {
	QSortFilterProxyModel_RecursiveFilteringEnabledChanged(this.h, (bool)(recursiveFilteringEnabled))
}

func (this *QSortFilterProxyModel) OnRecursiveFilteringEnabledChanged(slot func(recursiveFilteringEnabled bool)) {
	QSortFilterProxyModel_connect_RecursiveFilteringEnabledChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_RecursiveFilteringEnabledChanged
func miqt_exec_callback_QSortFilterProxyModel_RecursiveFilteringEnabledChanged(cb intptr_t, recursiveFilteringEnabled bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(recursiveFilteringEnabled bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(recursiveFilteringEnabled)

	gofunc(slotval1)
}

func (this *QSortFilterProxyModel) AutoAcceptChildRowsChanged(autoAcceptChildRows bool) {
	QSortFilterProxyModel_AutoAcceptChildRowsChanged(this.h, (bool)(autoAcceptChildRows))
}

func (this *QSortFilterProxyModel) OnAutoAcceptChildRowsChanged(slot func(autoAcceptChildRows bool)) {
	QSortFilterProxyModel_connect_AutoAcceptChildRowsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_AutoAcceptChildRowsChanged
func miqt_exec_callback_QSortFilterProxyModel_AutoAcceptChildRowsChanged(cb intptr_t, autoAcceptChildRows bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(autoAcceptChildRows bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(autoAcceptChildRows)

	gofunc(slotval1)
}

func QSortFilterProxyModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSortFilterProxyModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSortFilterProxyModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSortFilterProxyModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSortFilterProxyModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSortFilterProxyModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSortFilterProxyModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSortFilterProxyModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_MetaObject
func miqt_exec_callback_QSortFilterProxyModel_MetaObject(self QSortFilterProxyModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSortFilterProxyModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSortFilterProxyModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSortFilterProxyModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSortFilterProxyModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSortFilterProxyModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSortFilterProxyModel_Metacast
func miqt_exec_callback_QSortFilterProxyModel_Metacast(self QSortFilterProxyModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QSortFilterProxyModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
