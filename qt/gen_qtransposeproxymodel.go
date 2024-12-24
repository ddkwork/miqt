package qt

import (
	"unsafe"
)

type QTransposeProxyModel struct {
	h          uintptr
	isSubclass bool
}

// NewQTransposeProxyModel constructs a new QTransposeProxyModel object.
func NewQTransposeProxyModel() *QTransposeProxyModel {
	g := newQTransposeProxyModel(QTransposeProxyModel_new())
	g.isSubclass = true
	return g
}

// NewQTransposeProxyModel2 constructs a new QTransposeProxyModel object.
func NewQTransposeProxyModel2(parent *QObject) *QTransposeProxyModel {
	g := newQTransposeProxyModel(QTransposeProxyModel_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QTransposeProxyModel) MetaObject() *QMetaObject {
	return newQMetaObject(QTransposeProxyModel_MetaObject(this.h))
}

func (this *QTransposeProxyModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTransposeProxyModel_Metacast(this.h, param1_Cstring))
}

func QTransposeProxyModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTransposeProxyModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTransposeProxyModel) SetSourceModel(newSourceModel *QAbstractItemModel) {
	QTransposeProxyModel_SetSourceModel(this.h, newSourceModel.cPointer())
}

func (this *QTransposeProxyModel) RowCount(parent *QModelIndex) int {
	return (int)(QTransposeProxyModel_RowCount(this.h, parent.cPointer()))
}

func (this *QTransposeProxyModel) ColumnCount(parent *QModelIndex) int {
	return (int)(QTransposeProxyModel_ColumnCount(this.h, parent.cPointer()))
}

func (this *QTransposeProxyModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	_goptr := newQVariant(QTransposeProxyModel_HeaderData(this.h, (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransposeProxyModel) SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {
	return (bool)(QTransposeProxyModel_SetHeaderData(this.h, (int)(section), (int)(orientation), value.cPointer(), (int)(role)))
}

func (this *QTransposeProxyModel) SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
	roles_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_Keys_CArray))
	roles_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_Values_CArray))
	roles_ctr := 0
	for roles_k, roles_v := range roles {
		roles_Keys_CArray[roles_ctr] = (int)(roles_k)
		roles_Values_CArray[roles_ctr] = roles_v.cPointer()
		roles_ctr++
	}
	roles_mm := struct_miqt_map{
		len:    size_t(len(roles)),
		keys:   unsafe.Pointer(roles_Keys_CArray),
		values: unsafe.Pointer(roles_Values_CArray),
	}
	return (bool)(QTransposeProxyModel_SetItemData(this.h, index.cPointer(), roles_mm))
}

func (this *QTransposeProxyModel) Span(index *QModelIndex) *QSize {
	_goptr := newQSize(QTransposeProxyModel_Span(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransposeProxyModel) ItemData(index *QModelIndex) map[int]QVariant {
	var _mm struct_miqt_map = QTransposeProxyModel_ItemData(this.h, index.cPointer())
	_ret := make(map[int]QVariant, int(_mm.len))
	_Keys := (*[0xffff]int)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		_entry_Key := (int)(_Keys[i])

		_mapval_goptr := newQVariant(_Values[i])
		_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_mapval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QTransposeProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QTransposeProxyModel_MapFromSource(this.h, sourceIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransposeProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QTransposeProxyModel_MapToSource(this.h, proxyIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransposeProxyModel) Parent(index *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QTransposeProxyModel_Parent(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransposeProxyModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QTransposeProxyModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransposeProxyModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QTransposeProxyModel_InsertRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QTransposeProxyModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QTransposeProxyModel_RemoveRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QTransposeProxyModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QTransposeProxyModel_MoveRows(this.h, sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QTransposeProxyModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QTransposeProxyModel_InsertColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QTransposeProxyModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QTransposeProxyModel_RemoveColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QTransposeProxyModel) MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QTransposeProxyModel_MoveColumns(this.h, sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QTransposeProxyModel) Sort(column int, order SortOrder) {
	QTransposeProxyModel_Sort(this.h, (int)(column), (int)(order))
}

func QTransposeProxyModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTransposeProxyModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTransposeProxyModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTransposeProxyModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTransposeProxyModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTransposeProxyModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTransposeProxyModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTransposeProxyModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTransposeProxyModel_MetaObject
func miqt_exec_callback_QTransposeProxyModel_MetaObject(self QTransposeProxyModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTransposeProxyModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTransposeProxyModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTransposeProxyModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTransposeProxyModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTransposeProxyModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTransposeProxyModel_Metacast
func miqt_exec_callback_QTransposeProxyModel_Metacast(self QTransposeProxyModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTransposeProxyModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
