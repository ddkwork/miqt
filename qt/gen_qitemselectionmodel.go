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

func (this *QItemSelectionModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QItemSelectionModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QItemSelectionModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_MetaObject
func miqt_exec_callback_QItemSelectionModel_MetaObject(self QItemSelectionModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QItemSelectionModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QItemSelectionModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QItemSelectionModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QItemSelectionModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemSelectionModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemSelectionModel_Metacast
func miqt_exec_callback_QItemSelectionModel_Metacast(self QItemSelectionModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QItemSelectionModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
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
