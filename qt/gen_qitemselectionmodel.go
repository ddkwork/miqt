package qt

import (
	"unsafe"
)

type QItemSelectionModel__SelectionFlag int

const (
	QItemSelectionModel__NoUpdate       QItemSelectionModel__SelectionFlag = 0
	QItemSelectionModel__Clear          QItemSelectionModel__SelectionFlag = 1
	QItemSelectionModel__Select         QItemSelectionModel__SelectionFlag = 2
	QItemSelectionModel__Deselect       QItemSelectionModel__SelectionFlag = 4
	QItemSelectionModel__Toggle         QItemSelectionModel__SelectionFlag = 8
	QItemSelectionModel__Current        QItemSelectionModel__SelectionFlag = 16
	QItemSelectionModel__Rows           QItemSelectionModel__SelectionFlag = 32
	QItemSelectionModel__Columns        QItemSelectionModel__SelectionFlag = 64
	QItemSelectionModel__SelectCurrent  QItemSelectionModel__SelectionFlag = 18
	QItemSelectionModel__ToggleCurrent  QItemSelectionModel__SelectionFlag = 24
	QItemSelectionModel__ClearAndSelect QItemSelectionModel__SelectionFlag = 3
)

type QItemSelectionRange struct {
	h          uintptr
	isSubclass bool
}

// NewQItemSelectionRange constructs a new QItemSelectionRange object.
func NewQItemSelectionRange() *QItemSelectionRange {

	ret := newQItemSelectionRange(QItemSelectionRange_new())
	ret.isSubclass = true
	return ret
}

// NewQItemSelectionRange2 constructs a new QItemSelectionRange object.
func NewQItemSelectionRange2(topL *QModelIndex, bottomR *QModelIndex) *QItemSelectionRange {

	ret := newQItemSelectionRange(QItemSelectionRange_new2(topL.cPointer(), bottomR.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQItemSelectionRange3 constructs a new QItemSelectionRange object.
func NewQItemSelectionRange3(index *QModelIndex) *QItemSelectionRange {

	ret := newQItemSelectionRange(QItemSelectionRange_new3(index.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QItemSelectionRange) Swap(other *QItemSelectionRange) {
	QItemSelectionRange_Swap(this.h, other.cPointer())
}

func (this *QItemSelectionRange) Top() int {
	return (int)(QItemSelectionRange_Top(this.h))
}

func (this *QItemSelectionRange) Left() int {
	return (int)(QItemSelectionRange_Left(this.h))
}

func (this *QItemSelectionRange) Bottom() int {
	return (int)(QItemSelectionRange_Bottom(this.h))
}

func (this *QItemSelectionRange) Right() int {
	return (int)(QItemSelectionRange_Right(this.h))
}

func (this *QItemSelectionRange) Width() int {
	return (int)(QItemSelectionRange_Width(this.h))
}

func (this *QItemSelectionRange) Height() int {
	return (int)(QItemSelectionRange_Height(this.h))
}

func (this *QItemSelectionRange) TopLeft() *QPersistentModelIndex {
	return newQPersistentModelIndex(QItemSelectionRange_TopLeft(this.h))
}

func (this *QItemSelectionRange) BottomRight() *QPersistentModelIndex {
	return newQPersistentModelIndex(QItemSelectionRange_BottomRight(this.h))
}

func (this *QItemSelectionRange) Parent() *QModelIndex {
	_goptr := newQModelIndex(QItemSelectionRange_Parent(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QItemSelectionRange) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QItemSelectionRange_Model(this.h))
}

func (this *QItemSelectionRange) Contains(index *QModelIndex) bool {
	return (bool)(QItemSelectionRange_Contains(this.h, index.cPointer()))
}

func (this *QItemSelectionRange) Contains2(row int, column int, parentIndex *QModelIndex) bool {
	return (bool)(QItemSelectionRange_Contains2(this.h, (int)(row), (int)(column), parentIndex.cPointer()))
}

func (this *QItemSelectionRange) Intersects(other *QItemSelectionRange) bool {
	return (bool)(QItemSelectionRange_Intersects(this.h, other.cPointer()))
}

func (this *QItemSelectionRange) Intersected(other *QItemSelectionRange) *QItemSelectionRange {
	_goptr := newQItemSelectionRange(QItemSelectionRange_Intersected(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QItemSelectionRange) IsValid() bool {
	return (bool)(QItemSelectionRange_IsValid(this.h))
}

func (this *QItemSelectionRange) IsEmpty() bool {
	return (bool)(QItemSelectionRange_IsEmpty(this.h))
}

func (this *QItemSelectionRange) Indexes() []QModelIndex {
	var _ma struct_miqt_array = QItemSelectionRange_Indexes(this.h)
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

type QItemSelectionModel struct {
	h          uintptr
	isSubclass bool
}

// NewQItemSelectionModel constructs a new QItemSelectionModel object.
func NewQItemSelectionModel() *QItemSelectionModel {

	ret := newQItemSelectionModel(QItemSelectionModel_new())
	ret.isSubclass = true
	return ret
}

// NewQItemSelectionModel2 constructs a new QItemSelectionModel object.
func NewQItemSelectionModel2(model *QAbstractItemModel, parent *QObject) *QItemSelectionModel {

	ret := newQItemSelectionModel(QItemSelectionModel_new2(model.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQItemSelectionModel3 constructs a new QItemSelectionModel object.
func NewQItemSelectionModel3(model *QAbstractItemModel) *QItemSelectionModel {

	ret := newQItemSelectionModel(QItemSelectionModel_new3(model.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QItemSelectionModel) MetaObject() *QMetaObject {
	return newQMetaObject(QItemSelectionModel_MetaObject(this.h))
}

func (this *QItemSelectionModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QItemSelectionModel_Metacast(this.h, param1_Cstring))
}

func QItemSelectionModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QItemSelectionModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QItemSelectionModel) CurrentIndex() *QModelIndex {
	_goptr := newQModelIndex(QItemSelectionModel_CurrentIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QItemSelectionModel) IsSelected(index *QModelIndex) bool {
	return (bool)(QItemSelectionModel_IsSelected(this.h, index.cPointer()))
}

func (this *QItemSelectionModel) IsRowSelected(row int) bool {
	return (bool)(QItemSelectionModel_IsRowSelected(this.h, (int)(row)))
}

func (this *QItemSelectionModel) IsColumnSelected(column int) bool {
	return (bool)(QItemSelectionModel_IsColumnSelected(this.h, (int)(column)))
}

func (this *QItemSelectionModel) RowIntersectsSelection(row int) bool {
	return (bool)(QItemSelectionModel_RowIntersectsSelection(this.h, (int)(row)))
}

func (this *QItemSelectionModel) ColumnIntersectsSelection(column int) bool {
	return (bool)(QItemSelectionModel_ColumnIntersectsSelection(this.h, (int)(column)))
}

func (this *QItemSelectionModel) HasSelection() bool {
	return (bool)(QItemSelectionModel_HasSelection(this.h))
}

func (this *QItemSelectionModel) SelectedIndexes() []QModelIndex {
	var _ma struct_miqt_array = QItemSelectionModel_SelectedIndexes(this.h)
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QItemSelectionModel) SelectedRows() []QModelIndex {
	var _ma struct_miqt_array = QItemSelectionModel_SelectedRows(this.h)
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QItemSelectionModel) SelectedColumns() []QModelIndex {
	var _ma struct_miqt_array = QItemSelectionModel_SelectedColumns(this.h)
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QItemSelectionModel) Selection() *QItemSelection {
	_goptr := newQItemSelection(QItemSelectionModel_Selection(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QItemSelectionModel) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QItemSelectionModel_Model(this.h))
}

func (this *QItemSelectionModel) Model2() *QAbstractItemModel {
	return newQAbstractItemModel(QItemSelectionModel_Model2(this.h))
}

func (this *QItemSelectionModel) SetModel(model *QAbstractItemModel) {
	QItemSelectionModel_SetModel(this.h, model.cPointer())
}

func (this *QItemSelectionModel) SetCurrentIndex(index *QModelIndex, command SelectionFlag) {
	QItemSelectionModel_SetCurrentIndex(this.h, index.cPointer(), (int)(command))
}

func (this *QItemSelectionModel) Select(index *QModelIndex, command SelectionFlag) {
	QItemSelectionModel_Select(this.h, index.cPointer(), (int)(command))
}

func (this *QItemSelectionModel) Select2(selection *QItemSelection, command SelectionFlag) {
	QItemSelectionModel_Select2(this.h, selection.cPointer(), (int)(command))
}

func (this *QItemSelectionModel) Clear() {
	QItemSelectionModel_Clear(this.h)
}

func (this *QItemSelectionModel) Reset() {
	QItemSelectionModel_Reset(this.h)
}

func (this *QItemSelectionModel) ClearSelection() {
	QItemSelectionModel_ClearSelection(this.h)
}

func (this *QItemSelectionModel) ClearCurrentIndex() {
	QItemSelectionModel_ClearCurrentIndex(this.h)
}

func (this *QItemSelectionModel) SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {
	QItemSelectionModel_SelectionChanged(this.h, selected.cPointer(), deselected.cPointer())
}
func (this *QItemSelectionModel) OnSelectionChanged(slot func(selected *QItemSelection, deselected *QItemSelection)) {
	QItemSelectionModel_connect_SelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_SelectionChanged
func miqt_exec_callback_QItemSelectionModel_SelectionChanged(cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc(slotval1, slotval2)
}

func (this *QItemSelectionModel) CurrentChanged(current *QModelIndex, previous *QModelIndex) {
	QItemSelectionModel_CurrentChanged(this.h, current.cPointer(), previous.cPointer())
}
func (this *QItemSelectionModel) OnCurrentChanged(slot func(current *QModelIndex, previous *QModelIndex)) {
	QItemSelectionModel_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_CurrentChanged
func miqt_exec_callback_QItemSelectionModel_CurrentChanged(cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc(slotval1, slotval2)
}

func (this *QItemSelectionModel) CurrentRowChanged(current *QModelIndex, previous *QModelIndex) {
	QItemSelectionModel_CurrentRowChanged(this.h, current.cPointer(), previous.cPointer())
}
func (this *QItemSelectionModel) OnCurrentRowChanged(slot func(current *QModelIndex, previous *QModelIndex)) {
	QItemSelectionModel_connect_CurrentRowChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_CurrentRowChanged
func miqt_exec_callback_QItemSelectionModel_CurrentRowChanged(cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc(slotval1, slotval2)
}

func (this *QItemSelectionModel) CurrentColumnChanged(current *QModelIndex, previous *QModelIndex) {
	QItemSelectionModel_CurrentColumnChanged(this.h, current.cPointer(), previous.cPointer())
}
func (this *QItemSelectionModel) OnCurrentColumnChanged(slot func(current *QModelIndex, previous *QModelIndex)) {
	QItemSelectionModel_connect_CurrentColumnChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_CurrentColumnChanged
func miqt_exec_callback_QItemSelectionModel_CurrentColumnChanged(cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc(slotval1, slotval2)
}

func (this *QItemSelectionModel) ModelChanged(model *QAbstractItemModel) {
	QItemSelectionModel_ModelChanged(this.h, model.cPointer())
}
func (this *QItemSelectionModel) OnModelChanged(slot func(model *QAbstractItemModel)) {
	QItemSelectionModel_connect_ModelChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_ModelChanged
func miqt_exec_callback_QItemSelectionModel_ModelChanged(cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc(slotval1)
}

func QItemSelectionModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QItemSelectionModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QItemSelectionModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QItemSelectionModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QItemSelectionModel) IsRowSelected2(row int, parent *QModelIndex) bool {
	return (bool)(QItemSelectionModel_IsRowSelected2(this.h, (int)(row), parent.cPointer()))
}

func (this *QItemSelectionModel) IsColumnSelected2(column int, parent *QModelIndex) bool {
	return (bool)(QItemSelectionModel_IsColumnSelected2(this.h, (int)(column), parent.cPointer()))
}

func (this *QItemSelectionModel) RowIntersectsSelection2(row int, parent *QModelIndex) bool {
	return (bool)(QItemSelectionModel_RowIntersectsSelection2(this.h, (int)(row), parent.cPointer()))
}

func (this *QItemSelectionModel) ColumnIntersectsSelection2(column int, parent *QModelIndex) bool {
	return (bool)(QItemSelectionModel_ColumnIntersectsSelection2(this.h, (int)(column), parent.cPointer()))
}

func (this *QItemSelectionModel) SelectedRows1(column int) []QModelIndex {
	var _ma struct_miqt_array = QItemSelectionModel_SelectedRows1(this.h, (int)(column))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QItemSelectionModel) SelectedColumns1(row int) []QModelIndex {
	var _ma struct_miqt_array = QItemSelectionModel_SelectedColumns1(this.h, (int)(row))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QItemSelectionModel) callVirtualBase_SetCurrentIndex(index *QModelIndex, command SelectionFlag) {

	QItemSelectionModel_virtualbase_SetCurrentIndex(unsafe.Pointer(this.h), index.cPointer(), (int)(command))

}
func (this *QItemSelectionModel) OnSetCurrentIndex(slot func(super func(index *QModelIndex, command SelectionFlag), index *QModelIndex, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_SetCurrentIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_SetCurrentIndex
func miqt_exec_callback_QItemSelectionModel_SetCurrentIndex(self QItemSelectionModel, cb intptr_t, index *QModelIndex, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, command SelectionFlag), index *QModelIndex, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_SetCurrentIndex, slotval1, slotval2)

}

func (this *QItemSelectionModel) callVirtualBase_Select(index *QModelIndex, command SelectionFlag) {

	QItemSelectionModel_virtualbase_Select(unsafe.Pointer(this.h), index.cPointer(), (int)(command))

}
func (this *QItemSelectionModel) OnSelect(slot func(super func(index *QModelIndex, command SelectionFlag), index *QModelIndex, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_Select(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_Select
func miqt_exec_callback_QItemSelectionModel_Select(self QItemSelectionModel, cb intptr_t, index *QModelIndex, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, command SelectionFlag), index *QModelIndex, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_Select, slotval1, slotval2)

}

func (this *QItemSelectionModel) callVirtualBase_Select2(selection *QItemSelection, command SelectionFlag) {

	QItemSelectionModel_virtualbase_Select2(unsafe.Pointer(this.h), selection.cPointer(), (int)(command))

}
func (this *QItemSelectionModel) OnSelect2(slot func(super func(selection *QItemSelection, command SelectionFlag), selection *QItemSelection, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_Select2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_Select2
func miqt_exec_callback_QItemSelectionModel_Select2(self QItemSelectionModel, cb intptr_t, selection *QItemSelection, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection, command SelectionFlag), selection *QItemSelection, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_Select2, slotval1, slotval2)

}

func (this *QItemSelectionModel) callVirtualBase_Clear() {

	QItemSelectionModel_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QItemSelectionModel) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_Clear
func miqt_exec_callback_QItemSelectionModel_Clear(self QItemSelectionModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_Clear)

}

func (this *QItemSelectionModel) callVirtualBase_Reset() {

	QItemSelectionModel_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QItemSelectionModel) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_Reset
func miqt_exec_callback_QItemSelectionModel_Reset(self QItemSelectionModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_Reset)

}

func (this *QItemSelectionModel) callVirtualBase_ClearCurrentIndex() {

	QItemSelectionModel_virtualbase_ClearCurrentIndex(unsafe.Pointer(this.h))

}
func (this *QItemSelectionModel) OnClearCurrentIndex(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_ClearCurrentIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_ClearCurrentIndex
func miqt_exec_callback_QItemSelectionModel_ClearCurrentIndex(self QItemSelectionModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_ClearCurrentIndex)

}

func (this *QItemSelectionModel) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QItemSelectionModel_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QItemSelectionModel) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_Event
func miqt_exec_callback_QItemSelectionModel_Event(self QItemSelectionModel, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QItemSelectionModel{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QItemSelectionModel) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QItemSelectionModel_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QItemSelectionModel) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_EventFilter
func miqt_exec_callback_QItemSelectionModel_EventFilter(self QItemSelectionModel, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QItemSelectionModel{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QItemSelectionModel) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QItemSelectionModel_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QItemSelectionModel) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_TimerEvent
func miqt_exec_callback_QItemSelectionModel_TimerEvent(self QItemSelectionModel, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QItemSelectionModel) callVirtualBase_ChildEvent(event *QChildEvent) {

	QItemSelectionModel_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QItemSelectionModel) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_ChildEvent
func miqt_exec_callback_QItemSelectionModel_ChildEvent(self QItemSelectionModel, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QItemSelectionModel) callVirtualBase_CustomEvent(event *QEvent) {

	QItemSelectionModel_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QItemSelectionModel) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_CustomEvent
func miqt_exec_callback_QItemSelectionModel_CustomEvent(self QItemSelectionModel, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QItemSelectionModel) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QItemSelectionModel_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QItemSelectionModel) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_ConnectNotify
func miqt_exec_callback_QItemSelectionModel_ConnectNotify(self QItemSelectionModel, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QItemSelectionModel) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QItemSelectionModel_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QItemSelectionModel) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_DisconnectNotify
func miqt_exec_callback_QItemSelectionModel_DisconnectNotify(self QItemSelectionModel, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QItemSelectionModel{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

type QItemSelection struct {
	h          uintptr
	isSubclass bool
}

// NewQItemSelection constructs a new QItemSelection object.
func NewQItemSelection(topLeft *QModelIndex, bottomRight *QModelIndex) *QItemSelection {

	ret := newQItemSelection(QItemSelection_new(topLeft.cPointer(), bottomRight.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQItemSelection2 constructs a new QItemSelection object.
func NewQItemSelection2() *QItemSelection {

	ret := newQItemSelection(QItemSelection_new2())
	ret.isSubclass = true
	return ret
}

// NewQItemSelection3 constructs a new QItemSelection object.
func NewQItemSelection3(param1 *QItemSelection) *QItemSelection {

	ret := newQItemSelection(QItemSelection_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QItemSelection) Select(topLeft *QModelIndex, bottomRight *QModelIndex) {
	QItemSelection_Select(this.h, topLeft.cPointer(), bottomRight.cPointer())
}

func (this *QItemSelection) Contains(index *QModelIndex) bool {
	return (bool)(QItemSelection_Contains(this.h, index.cPointer()))
}

func (this *QItemSelection) Indexes() []QModelIndex {
	var _ma struct_miqt_array = QItemSelection_Indexes(this.h)
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QItemSelection) Merge(other *QItemSelection, command SelectionFlag) {
	QItemSelection_Merge(this.h, other.cPointer(), (int)(command))
}

func QItemSelection_Split(rangeVal *QItemSelectionRange, other *QItemSelectionRange, result *QItemSelection) {
	QItemSelection_Split(rangeVal.cPointer(), other.cPointer(), result.cPointer())
}
