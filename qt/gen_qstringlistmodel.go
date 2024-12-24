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

	ret := newQStringListModel(QStringListModel_new())
	ret.isSubclass = true
	return ret
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

	ret := newQStringListModel(QStringListModel_new2(strings_ma))
	ret.isSubclass = true
	return ret
}

// NewQStringListModel3 constructs a new QStringListModel object.
func NewQStringListModel3(parent *QObject) *QStringListModel {

	ret := newQStringListModel(QStringListModel_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
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

	ret := newQStringListModel(QStringListModel_new4(strings_ma, parent.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QStringListModel) callVirtualBase_RowCount(parent *QModelIndex) int {

	return (int)(QStringListModel_virtualbase_RowCount(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QStringListModel) OnRowCount(slot func(super func(parent *QModelIndex) int, parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_RowCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_RowCount
func miqt_exec_callback_QStringListModel_RowCount(self QStringListModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) int, parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_RowCount, slotval1)

	return (int)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_Sibling(row int, column int, idx *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QStringListModel_virtualbase_Sibling(unsafe.Pointer(this.h), (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStringListModel) OnSibling(slot func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_Sibling(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_Sibling
func miqt_exec_callback_QStringListModel_Sibling(self QStringListModel, cb intptr_t, row int, column int, idx *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(idx)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_Sibling, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QStringListModel) callVirtualBase_Data(index *QModelIndex, role int) *QVariant {

	_goptr := newQVariant(QStringListModel_virtualbase_Data(unsafe.Pointer(this.h), index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStringListModel) OnData(slot func(super func(index *QModelIndex, role int) *QVariant, index *QModelIndex, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_Data
func miqt_exec_callback_QStringListModel_Data(self QStringListModel, cb intptr_t, index *QModelIndex, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, role int) *QVariant, index *QModelIndex, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (int)(role)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_Data, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QStringListModel) callVirtualBase_SetData(index *QModelIndex, value *QVariant, role int) bool {

	return (bool)(QStringListModel_virtualbase_SetData(unsafe.Pointer(this.h), index.cPointer(), value.cPointer(), (int)(role)))

}
func (this *QStringListModel) OnSetData(slot func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_SetData
func miqt_exec_callback_QStringListModel_SetData(self QStringListModel, cb intptr_t, index *QModelIndex, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQVariant(value)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_SetData, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_ClearItemData(index *QModelIndex) bool {

	return (bool)(QStringListModel_virtualbase_ClearItemData(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QStringListModel) OnClearItemData(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_ClearItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_ClearItemData
func miqt_exec_callback_QStringListModel_ClearItemData(self QStringListModel, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_ClearItemData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_Flags(index *QModelIndex) ItemFlag {

	return (ItemFlag)(QStringListModel_virtualbase_Flags(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QStringListModel) OnFlags(slot func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_Flags(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_Flags
func miqt_exec_callback_QStringListModel_Flags(self QStringListModel, cb intptr_t, index *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_Flags, slotval1)

	return (int)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_InsertRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QStringListModel_virtualbase_InsertRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QStringListModel) OnInsertRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_InsertRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_InsertRows
func miqt_exec_callback_QStringListModel_InsertRows(self QStringListModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_InsertRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_RemoveRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QStringListModel_virtualbase_RemoveRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QStringListModel) OnRemoveRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_RemoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_RemoveRows
func miqt_exec_callback_QStringListModel_RemoveRows(self QStringListModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_RemoveRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QStringListModel_virtualbase_MoveRows(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QStringListModel) OnMoveRows(slot func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_MoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_MoveRows
func miqt_exec_callback_QStringListModel_MoveRows(self QStringListModel, cb intptr_t, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(sourceParent)

	slotval2 := (int)(sourceRow)

	slotval3 := (int)(count)

	slotval4 := newQModelIndex(destinationParent)

	slotval5 := (int)(destinationChild)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_MoveRows, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_ItemData(index *QModelIndex) map[int]QVariant {

	var _mm struct_miqt_map = QStringListModel_virtualbase_ItemData(unsafe.Pointer(this.h), index.cPointer())
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
func (this *QStringListModel) OnItemData(slot func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_ItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_ItemData
func miqt_exec_callback_QStringListModel_ItemData(self QStringListModel, cb intptr_t, index *QModelIndex) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_ItemData, slotval1)
	virtualReturn_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Keys_CArray))
	virtualReturn_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Values_CArray))
	virtualReturn_ctr := 0
	for virtualReturn_k, virtualReturn_v := range virtualReturn {
		virtualReturn_Keys_CArray[virtualReturn_ctr] = (int)(virtualReturn_k)
		virtualReturn_Values_CArray[virtualReturn_ctr] = virtualReturn_v.cPointer()
		virtualReturn_ctr++
	}
	virtualReturn_mm := struct_miqt_map{
		len:    size_t(len(virtualReturn)),
		keys:   unsafe.Pointer(virtualReturn_Keys_CArray),
		values: unsafe.Pointer(virtualReturn_Values_CArray),
	}

	return virtualReturn_mm

}

func (this *QStringListModel) callVirtualBase_SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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

	return (bool)(QStringListModel_virtualbase_SetItemData(unsafe.Pointer(this.h), index.cPointer(), roles_mm))

}
func (this *QStringListModel) OnSetItemData(slot func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_SetItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_SetItemData
func miqt_exec_callback_QStringListModel_SetItemData(self QStringListModel, cb intptr_t, index *QModelIndex, roles struct_miqt_map) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	var roles_mm struct_miqt_map = roles
	roles_ret := make(map[int]QVariant, int(roles_mm.len))
	roles_Keys := (*[0xffff]int)(unsafe.Pointer(roles_mm.keys))
	roles_Values := (*[0xffff]*QVariant)(unsafe.Pointer(roles_mm.values))
	for i := 0; i < int(roles_mm.len); i++ {
		roles_entry_Key := (int)(roles_Keys[i])

		roles_mapval_goptr := newQVariant(roles_Values[i])
		roles_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		roles_entry_Value := *roles_mapval_goptr

		roles_ret[roles_entry_Key] = roles_entry_Value
	}
	slotval2 := roles_ret

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_SetItemData, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_Sort(column int, order SortOrder) {

	QStringListModel_virtualbase_Sort(unsafe.Pointer(this.h), (int)(column), (int)(order))

}
func (this *QStringListModel) OnSort(slot func(super func(column int, order SortOrder), column int, order SortOrder)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_Sort(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_Sort
func miqt_exec_callback_QStringListModel_Sort(self QStringListModel, cb intptr_t, column int, order int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, order SortOrder), column int, order SortOrder))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (SortOrder)(order)

	gofunc((&QStringListModel{h: self}).callVirtualBase_Sort, slotval1, slotval2)

}

func (this *QStringListModel) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QStringListModel_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QStringListModel) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_SupportedDropActions
func miqt_exec_callback_QStringListModel_SupportedDropActions(self QStringListModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QStringListModel) callVirtualBase_Index(row int, column int, parent *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QStringListModel_virtualbase_Index(unsafe.Pointer(this.h), (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStringListModel) OnIndex(slot func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_Index(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_Index
func miqt_exec_callback_QStringListModel_Index(self QStringListModel, cb intptr_t, row int, column int, parent *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_Index, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QStringListModel) callVirtualBase_DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QStringListModel_virtualbase_DropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QStringListModel) OnDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStringListModel_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStringListModel_DropMimeData
func miqt_exec_callback_QStringListModel_DropMimeData(self QStringListModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMimeData(data)

	slotval2 := (DropAction)(action)

	slotval3 := (int)(row)

	slotval4 := (int)(column)

	slotval5 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStringListModel{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}
