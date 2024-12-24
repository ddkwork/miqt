package qt

import (
	"unsafe"
)

type QIdentityProxyModel struct {
	h          uintptr
	isSubclass bool
}

// NewQIdentityProxyModel constructs a new QIdentityProxyModel object.
func NewQIdentityProxyModel() *QIdentityProxyModel {
	ret := newQIdentityProxyModel(QIdentityProxyModel_new())
	ret.isSubclass = true
	return ret
}

// NewQIdentityProxyModel2 constructs a new QIdentityProxyModel object.
func NewQIdentityProxyModel2(parent *QObject) *QIdentityProxyModel {
	ret := newQIdentityProxyModel(QIdentityProxyModel_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QIdentityProxyModel) MetaObject() *QMetaObject {
	return newQMetaObject(QIdentityProxyModel_MetaObject(this.h))
}

func (this *QIdentityProxyModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QIdentityProxyModel_Metacast(this.h, param1_Cstring))
}

func QIdentityProxyModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QIdentityProxyModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIdentityProxyModel) ColumnCount(parent *QModelIndex) int {
	return (int)(QIdentityProxyModel_ColumnCount(this.h, parent.cPointer()))
}

func (this *QIdentityProxyModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QIdentityProxyModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QIdentityProxyModel_MapFromSource(this.h, sourceIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QIdentityProxyModel_MapToSource(this.h, proxyIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) Parent(child *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QIdentityProxyModel_Parent(this.h, child.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) RowCount(parent *QModelIndex) int {
	return (int)(QIdentityProxyModel_RowCount(this.h, parent.cPointer()))
}

func (this *QIdentityProxyModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	_goptr := newQVariant(QIdentityProxyModel_HeaderData(this.h, (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QIdentityProxyModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QIdentityProxyModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QIdentityProxyModel_Sibling(this.h, (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) MapSelectionFromSource(selection *QItemSelection) *QItemSelection {
	_goptr := newQItemSelection(QIdentityProxyModel_MapSelectionFromSource(this.h, selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) MapSelectionToSource(selection *QItemSelection) *QItemSelection {
	_goptr := newQItemSelection(QIdentityProxyModel_MapSelectionToSource(this.h, selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIdentityProxyModel) Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {
	var _ma struct_miqt_array = QIdentityProxyModel_Match(this.h, start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QIdentityProxyModel) SetSourceModel(sourceModel *QAbstractItemModel) {
	QIdentityProxyModel_SetSourceModel(this.h, sourceModel.cPointer())
}

func (this *QIdentityProxyModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QIdentityProxyModel_InsertColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QIdentityProxyModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QIdentityProxyModel_InsertRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QIdentityProxyModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QIdentityProxyModel_RemoveColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QIdentityProxyModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QIdentityProxyModel_RemoveRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QIdentityProxyModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QIdentityProxyModel_MoveRows(this.h, sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QIdentityProxyModel) MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QIdentityProxyModel_MoveColumns(this.h, sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QIdentityProxyModel) HandleSourceLayoutChanges() bool {
	return (bool)(QIdentityProxyModel_HandleSourceLayoutChanges(this.h))
}

func (this *QIdentityProxyModel) HandleSourceDataChanges() bool {
	return (bool)(QIdentityProxyModel_HandleSourceDataChanges(this.h))
}

func QIdentityProxyModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIdentityProxyModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIdentityProxyModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIdentityProxyModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIdentityProxyModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QIdentityProxyModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QIdentityProxyModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIdentityProxyModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIdentityProxyModel_MetaObject
func miqt_exec_callback_QIdentityProxyModel_MetaObject(self QIdentityProxyModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIdentityProxyModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QIdentityProxyModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QIdentityProxyModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QIdentityProxyModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIdentityProxyModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIdentityProxyModel_Metacast
func miqt_exec_callback_QIdentityProxyModel_Metacast(self QIdentityProxyModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QIdentityProxyModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
