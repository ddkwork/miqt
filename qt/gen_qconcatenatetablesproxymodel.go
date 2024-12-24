package qt

import (
	"unsafe"
)

type QConcatenateTablesProxyModel struct {
	h          uintptr
	isSubclass bool
}

// NewQConcatenateTablesProxyModel constructs a new QConcatenateTablesProxyModel object.
func NewQConcatenateTablesProxyModel() *QConcatenateTablesProxyModel {
	ret := newQConcatenateTablesProxyModel(QConcatenateTablesProxyModel_new())
	ret.isSubclass = true
	return ret
}

// NewQConcatenateTablesProxyModel2 constructs a new QConcatenateTablesProxyModel object.
func NewQConcatenateTablesProxyModel2(parent *QObject) *QConcatenateTablesProxyModel {
	ret := newQConcatenateTablesProxyModel(QConcatenateTablesProxyModel_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QConcatenateTablesProxyModel) MetaObject() *QMetaObject {
	return newQMetaObject(QConcatenateTablesProxyModel_MetaObject(this.h))
}

func (this *QConcatenateTablesProxyModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QConcatenateTablesProxyModel_Metacast(this.h, param1_Cstring))
}

func QConcatenateTablesProxyModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QConcatenateTablesProxyModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QConcatenateTablesProxyModel) SourceModels() []*QAbstractItemModel {
	var _ma struct_miqt_array = QConcatenateTablesProxyModel_SourceModels(this.h)
	_ret := make([]*QAbstractItemModel, int(_ma.len))
	_outCast := (*[0xffff]*QAbstractItemModel)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAbstractItemModel(_outCast[i])
	}
	return _ret
}

func (this *QConcatenateTablesProxyModel) AddSourceModel(sourceModel *QAbstractItemModel) {
	QConcatenateTablesProxyModel_AddSourceModel(this.h, sourceModel.cPointer())
}

func (this *QConcatenateTablesProxyModel) RemoveSourceModel(sourceModel *QAbstractItemModel) {
	QConcatenateTablesProxyModel_RemoveSourceModel(this.h, sourceModel.cPointer())
}

func (this *QConcatenateTablesProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QConcatenateTablesProxyModel_MapFromSource(this.h, sourceIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QConcatenateTablesProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QConcatenateTablesProxyModel_MapToSource(this.h, proxyIndex.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QConcatenateTablesProxyModel) Data(index *QModelIndex, role int) *QVariant {
	_goptr := newQVariant(QConcatenateTablesProxyModel_Data(this.h, index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QConcatenateTablesProxyModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	return (bool)(QConcatenateTablesProxyModel_SetData(this.h, index.cPointer(), value.cPointer(), (int)(role)))
}

func (this *QConcatenateTablesProxyModel) ItemData(proxyIndex *QModelIndex) map[int]QVariant {
	var _mm struct_miqt_map = QConcatenateTablesProxyModel_ItemData(this.h, proxyIndex.cPointer())
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

func (this *QConcatenateTablesProxyModel) SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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
	return (bool)(QConcatenateTablesProxyModel_SetItemData(this.h, index.cPointer(), roles_mm))
}

func (this *QConcatenateTablesProxyModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QConcatenateTablesProxyModel_Flags(this.h, index.cPointer()))
}

func (this *QConcatenateTablesProxyModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QConcatenateTablesProxyModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QConcatenateTablesProxyModel) Parent(index *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QConcatenateTablesProxyModel_Parent(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QConcatenateTablesProxyModel) RowCount(parent *QModelIndex) int {
	return (int)(QConcatenateTablesProxyModel_RowCount(this.h, parent.cPointer()))
}

func (this *QConcatenateTablesProxyModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	_goptr := newQVariant(QConcatenateTablesProxyModel_HeaderData(this.h, (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QConcatenateTablesProxyModel) ColumnCount(parent *QModelIndex) int {
	return (int)(QConcatenateTablesProxyModel_ColumnCount(this.h, parent.cPointer()))
}

func (this *QConcatenateTablesProxyModel) MimeTypes() []string {
	var _ma struct_miqt_array = QConcatenateTablesProxyModel_MimeTypes(this.h)
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

func (this *QConcatenateTablesProxyModel) MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}
	return newQMimeData(QConcatenateTablesProxyModel_MimeData(this.h, indexes_ma))
}

func (this *QConcatenateTablesProxyModel) CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QConcatenateTablesProxyModel_CanDropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QConcatenateTablesProxyModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QConcatenateTablesProxyModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QConcatenateTablesProxyModel) Span(index *QModelIndex) *QSize {
	_goptr := newQSize(QConcatenateTablesProxyModel_Span(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QConcatenateTablesProxyModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QConcatenateTablesProxyModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QConcatenateTablesProxyModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QConcatenateTablesProxyModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QConcatenateTablesProxyModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QConcatenateTablesProxyModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QConcatenateTablesProxyModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QConcatenateTablesProxyModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QConcatenateTablesProxyModel_MetaObject
func miqt_exec_callback_QConcatenateTablesProxyModel_MetaObject(self QConcatenateTablesProxyModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QConcatenateTablesProxyModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QConcatenateTablesProxyModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QConcatenateTablesProxyModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QConcatenateTablesProxyModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QConcatenateTablesProxyModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QConcatenateTablesProxyModel_Metacast
func miqt_exec_callback_QConcatenateTablesProxyModel_Metacast(self QConcatenateTablesProxyModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QConcatenateTablesProxyModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
