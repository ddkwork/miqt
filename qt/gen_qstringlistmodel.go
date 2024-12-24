package qt

import (
	"unsafe"
)

type QStringListModel struct {
	h          uintptr
	isSubclass bool
}

// NewQStringListModel constructs a new QStringListModel object.
func NewQStringListModel() *QStringListModel {
	g := newQStringListModel(QStringListModel_new())
	g.isSubclass = true
	return g
}

// NewQStringListModel2 constructs a new QStringListModel object.
func NewQStringListModel2(strings []string) *QStringListModel {
	strings_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(strings))))
	defer free(unsafe.Pointer(strings_CArray))
	for i := range strings {
		strings_i_ms := struct_miqt_string{}
		strings_i_ms.data = CString(strings[i])
		strings_i_ms.len = size_t(len(strings[i]))
		defer free(unsafe.Pointer(strings_i_ms.data))
		strings_CArray[i] = strings_i_ms
	}
	strings_ma := struct_miqt_array{len: size_t(len(strings)), data: unsafe.Pointer(strings_CArray)}
	g := newQStringListModel(QStringListModel_new2(strings_ma))
	g.isSubclass = true
	return g
}

// NewQStringListModel3 constructs a new QStringListModel object.
func NewQStringListModel3(parent *QObject) *QStringListModel {
	g := newQStringListModel(QStringListModel_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQStringListModel4 constructs a new QStringListModel object.
func NewQStringListModel4(strings []string, parent *QObject) *QStringListModel {
	strings_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(strings))))
	defer free(unsafe.Pointer(strings_CArray))
	for i := range strings {
		strings_i_ms := struct_miqt_string{}
		strings_i_ms.data = CString(strings[i])
		strings_i_ms.len = size_t(len(strings[i]))
		defer free(unsafe.Pointer(strings_i_ms.data))
		strings_CArray[i] = strings_i_ms
	}
	strings_ma := struct_miqt_array{len: size_t(len(strings)), data: unsafe.Pointer(strings_CArray)}
	g := newQStringListModel(QStringListModel_new4(strings_ma, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QStringListModel) MetaObject() *QMetaObject {
	return newQMetaObject(QStringListModel_MetaObject(this.h))
}

func (this *QStringListModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStringListModel_Metacast(this.h, param1_Cstring))
}

func QStringListModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStringListModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStringListModel) RowCount(parent *QModelIndex) int {
	return (int)(QStringListModel_RowCount(this.h, parent.cPointer()))
}

func (this *QStringListModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QStringListModel_Sibling(this.h, (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStringListModel) Data(index *QModelIndex, role int) *QVariant {
	_goptr := newQVariant(QStringListModel_Data(this.h, index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStringListModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	return (bool)(QStringListModel_SetData(this.h, index.cPointer(), value.cPointer(), (int)(role)))
}

func (this *QStringListModel) ClearItemData(index *QModelIndex) bool {
	return (bool)(QStringListModel_ClearItemData(this.h, index.cPointer()))
}

func (this *QStringListModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QStringListModel_Flags(this.h, index.cPointer()))
}

func (this *QStringListModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QStringListModel_InsertRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QStringListModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QStringListModel_RemoveRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QStringListModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QStringListModel_MoveRows(this.h, sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QStringListModel) ItemData(index *QModelIndex) map[int]QVariant {
	var _mm struct_miqt_map = QStringListModel_ItemData(this.h, index.cPointer())
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

func (this *QStringListModel) SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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
	return (bool)(QStringListModel_SetItemData(this.h, index.cPointer(), roles_mm))
}

func (this *QStringListModel) Sort(column int, order SortOrder) {
	QStringListModel_Sort(this.h, (int)(column), (int)(order))
}

func (this *QStringListModel) StringList() []string {
	var _ma struct_miqt_array = QStringListModel_StringList(this.h)
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

func (this *QStringListModel) SetStringList(strings []string) {
	strings_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(strings))))
	defer free(unsafe.Pointer(strings_CArray))
	for i := range strings {
		strings_i_ms := struct_miqt_string{}
		strings_i_ms.data = CString(strings[i])
		strings_i_ms.len = size_t(len(strings[i]))
		defer free(unsafe.Pointer(strings_i_ms.data))
		strings_CArray[i] = strings_i_ms
	}
	strings_ma := struct_miqt_array{len: size_t(len(strings)), data: unsafe.Pointer(strings_CArray)}
	QStringListModel_SetStringList(this.h, strings_ma)
}

func (this *QStringListModel) SupportedDropActions() DropAction {
	return (DropAction)(QStringListModel_SupportedDropActions(this.h))
}

func QStringListModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStringListModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStringListModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStringListModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStringListModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QStringListModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QStringListModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_MetaObject
func miqt_exec_callback_QStringListModel_MetaObject(self QStringListModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QStringListModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QStringListModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QStringListModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_Metacast
func miqt_exec_callback_QStringListModel_Metacast(self QStringListModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
