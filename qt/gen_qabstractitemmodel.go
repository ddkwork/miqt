package qt

import (
	"unsafe"
)

type QAbstractItemModel__LayoutChangeHint int

const (
	QAbstractItemModel__NoLayoutChangeHint QAbstractItemModel__LayoutChangeHint = 0
	QAbstractItemModel__VerticalSortHint   QAbstractItemModel__LayoutChangeHint = 1
	QAbstractItemModel__HorizontalSortHint QAbstractItemModel__LayoutChangeHint = 2
)

type QAbstractItemModel__CheckIndexOption int

const (
	QAbstractItemModel__NoOption        QAbstractItemModel__CheckIndexOption = 0
	QAbstractItemModel__IndexIsValid    QAbstractItemModel__CheckIndexOption = 1
	QAbstractItemModel__DoNotUseParent  QAbstractItemModel__CheckIndexOption = 2
	QAbstractItemModel__ParentIsInvalid QAbstractItemModel__CheckIndexOption = 4
)

type QModelRoleData struct {
	h          uintptr
	isSubclass bool
}

// NewQModelRoleData constructs a new QModelRoleData object.
func NewQModelRoleData(role int) *QModelRoleData {
	ret := newQModelRoleData(QModelRoleData_new((int)(role)))
	ret.isSubclass = true
	return ret
}

func (this *QModelRoleData) Role() int {
	return (int)(QModelRoleData_Role(this.h))
}

func (this *QModelRoleData) Data() *QVariant {
	return newQVariant(QModelRoleData_Data(this.h))
}

func (this *QModelRoleData) Data2() *QVariant {
	return newQVariant(QModelRoleData_Data2(this.h))
}

func (this *QModelRoleData) ClearData() {
	QModelRoleData_ClearData(this.h)
}

func (this *QModelRoleData) OperatorAssign(param1 *QModelRoleData) {
	QModelRoleData_OperatorAssign(this.h, param1.cPointer())
}

type QModelRoleDataSpan struct {
	h          uintptr
	isSubclass bool
}

// NewQModelRoleDataSpan constructs a new QModelRoleDataSpan object.
func NewQModelRoleDataSpan() *QModelRoleDataSpan {
	ret := newQModelRoleDataSpan(QModelRoleDataSpan_new())
	ret.isSubclass = true
	return ret
}

// NewQModelRoleDataSpan2 constructs a new QModelRoleDataSpan object.
func NewQModelRoleDataSpan2(modelRoleData *QModelRoleData) *QModelRoleDataSpan {
	ret := newQModelRoleDataSpan(QModelRoleDataSpan_new2(modelRoleData.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQModelRoleDataSpan3 constructs a new QModelRoleDataSpan object.
func NewQModelRoleDataSpan3(modelRoleData *QModelRoleData, lenVal int64) *QModelRoleDataSpan {
	ret := newQModelRoleDataSpan(QModelRoleDataSpan_new3(modelRoleData.cPointer(), (ptrdiff_t)(lenVal)))
	ret.isSubclass = true
	return ret
}

// NewQModelRoleDataSpan4 constructs a new QModelRoleDataSpan object.
func NewQModelRoleDataSpan4(param1 *QModelRoleDataSpan) *QModelRoleDataSpan {
	ret := newQModelRoleDataSpan(QModelRoleDataSpan_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QModelRoleDataSpan) Size() int64 {
	return (int64)(QModelRoleDataSpan_Size(this.h))
}

func (this *QModelRoleDataSpan) Length() int64 {
	return (int64)(QModelRoleDataSpan_Length(this.h))
}

func (this *QModelRoleDataSpan) Data() *QModelRoleData {
	return newQModelRoleData(QModelRoleDataSpan_Data(this.h))
}

func (this *QModelRoleDataSpan) Begin() *QModelRoleData {
	return newQModelRoleData(QModelRoleDataSpan_Begin(this.h))
}

func (this *QModelRoleDataSpan) End() *QModelRoleData {
	return newQModelRoleData(QModelRoleDataSpan_End(this.h))
}

func (this *QModelRoleDataSpan) OperatorSubscript(index int64) *QModelRoleData {
	return newQModelRoleData(QModelRoleDataSpan_OperatorSubscript(this.h, (ptrdiff_t)(index)))
}

func (this *QModelRoleDataSpan) DataForRole(role int) *QVariant {
	return newQVariant(QModelRoleDataSpan_DataForRole(this.h, (int)(role)))
}

type QModelIndex struct {
	h          uintptr
	isSubclass bool
}

// NewQModelIndex constructs a new QModelIndex object.
func NewQModelIndex() *QModelIndex {
	ret := newQModelIndex(QModelIndex_new())
	ret.isSubclass = true
	return ret
}

// NewQModelIndex2 constructs a new QModelIndex object.
func NewQModelIndex2(param1 *QModelIndex) *QModelIndex {
	ret := newQModelIndex(QModelIndex_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QModelIndex) Row() int {
	return (int)(QModelIndex_Row(this.h))
}

func (this *QModelIndex) Column() int {
	return (int)(QModelIndex_Column(this.h))
}

func (this *QModelIndex) InternalId() uintptr {
	return (uintptr)(QModelIndex_InternalId(this.h))
}

func (this *QModelIndex) InternalPointer() unsafe.Pointer {
	return (unsafe.Pointer)(QModelIndex_InternalPointer(this.h))
}

func (this *QModelIndex) ConstInternalPointer() unsafe.Pointer {
	return (unsafe.Pointer)(QModelIndex_ConstInternalPointer(this.h))
}

func (this *QModelIndex) Parent() *QModelIndex {
	_goptr := newQModelIndex(QModelIndex_Parent(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QModelIndex) Sibling(row int, column int) *QModelIndex {
	_goptr := newQModelIndex(QModelIndex_Sibling(this.h, (int)(row), (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QModelIndex) SiblingAtColumn(column int) *QModelIndex {
	_goptr := newQModelIndex(QModelIndex_SiblingAtColumn(this.h, (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QModelIndex) SiblingAtRow(row int) *QModelIndex {
	_goptr := newQModelIndex(QModelIndex_SiblingAtRow(this.h, (int)(row)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QModelIndex) Data() *QVariant {
	_goptr := newQVariant(QModelIndex_Data(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QModelIndex) MultiData(roleDataSpan QModelRoleDataSpan) {
	QModelIndex_MultiData(this.h, roleDataSpan.cPointer())
}

func (this *QModelIndex) Flags() ItemFlag {
	return (ItemFlag)(QModelIndex_Flags(this.h))
}

func (this *QModelIndex) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QModelIndex_Model(this.h))
}

func (this *QModelIndex) IsValid() bool {
	return (bool)(QModelIndex_IsValid(this.h))
}

func (this *QModelIndex) Data1(role int) *QVariant {
	_goptr := newQVariant(QModelIndex_Data1(this.h, (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QPersistentModelIndex struct {
	h          uintptr
	isSubclass bool
}

// NewQPersistentModelIndex constructs a new QPersistentModelIndex object.
func NewQPersistentModelIndex() *QPersistentModelIndex {
	ret := newQPersistentModelIndex(QPersistentModelIndex_new())
	ret.isSubclass = true
	return ret
}

// NewQPersistentModelIndex2 constructs a new QPersistentModelIndex object.
func NewQPersistentModelIndex2(index *QModelIndex) *QPersistentModelIndex {
	ret := newQPersistentModelIndex(QPersistentModelIndex_new2(index.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPersistentModelIndex3 constructs a new QPersistentModelIndex object.
func NewQPersistentModelIndex3(other *QPersistentModelIndex) *QPersistentModelIndex {
	ret := newQPersistentModelIndex(QPersistentModelIndex_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPersistentModelIndex) OperatorAssign(other *QPersistentModelIndex) {
	QPersistentModelIndex_OperatorAssign(this.h, other.cPointer())
}

func (this *QPersistentModelIndex) Swap(other *QPersistentModelIndex) {
	QPersistentModelIndex_Swap(this.h, other.cPointer())
}

func (this *QPersistentModelIndex) OperatorAssignWithOther(other *QModelIndex) {
	QPersistentModelIndex_OperatorAssignWithOther(this.h, other.cPointer())
}

func (this *QPersistentModelIndex) Row() int {
	return (int)(QPersistentModelIndex_Row(this.h))
}

func (this *QPersistentModelIndex) Column() int {
	return (int)(QPersistentModelIndex_Column(this.h))
}

func (this *QPersistentModelIndex) InternalPointer() unsafe.Pointer {
	return (unsafe.Pointer)(QPersistentModelIndex_InternalPointer(this.h))
}

func (this *QPersistentModelIndex) ConstInternalPointer() unsafe.Pointer {
	return (unsafe.Pointer)(QPersistentModelIndex_ConstInternalPointer(this.h))
}

func (this *QPersistentModelIndex) InternalId() uintptr {
	return (uintptr)(QPersistentModelIndex_InternalId(this.h))
}

func (this *QPersistentModelIndex) Parent() *QModelIndex {
	_goptr := newQModelIndex(QPersistentModelIndex_Parent(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPersistentModelIndex) Sibling(row int, column int) *QModelIndex {
	_goptr := newQModelIndex(QPersistentModelIndex_Sibling(this.h, (int)(row), (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPersistentModelIndex) Data() *QVariant {
	_goptr := newQVariant(QPersistentModelIndex_Data(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPersistentModelIndex) MultiData(roleDataSpan QModelRoleDataSpan) {
	QPersistentModelIndex_MultiData(this.h, roleDataSpan.cPointer())
}

func (this *QPersistentModelIndex) Flags() ItemFlag {
	return (ItemFlag)(QPersistentModelIndex_Flags(this.h))
}

func (this *QPersistentModelIndex) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QPersistentModelIndex_Model(this.h))
}

func (this *QPersistentModelIndex) IsValid() bool {
	return (bool)(QPersistentModelIndex_IsValid(this.h))
}

func (this *QPersistentModelIndex) Data1(role int) *QVariant {
	_goptr := newQVariant(QPersistentModelIndex_Data1(this.h, (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QAbstractItemModel struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractItemModel constructs a new QAbstractItemModel object.
func NewQAbstractItemModel() *QAbstractItemModel {
	ret := newQAbstractItemModel(QAbstractItemModel_new())
	ret.isSubclass = true
	return ret
}

// NewQAbstractItemModel2 constructs a new QAbstractItemModel object.
func NewQAbstractItemModel2(parent *QObject) *QAbstractItemModel {
	ret := newQAbstractItemModel(QAbstractItemModel_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractItemModel) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractItemModel_MetaObject(this.h))
}

func (this *QAbstractItemModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractItemModel_Metacast(this.h, param1_Cstring))
}

func QAbstractItemModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractItemModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractItemModel) HasIndex(row int, column int) bool {
	return (bool)(QAbstractItemModel_HasIndex(this.h, (int)(row), (int)(column)))
}

func (this *QAbstractItemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractItemModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemModel) Parent(child *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractItemModel_Parent(this.h, child.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractItemModel_Sibling(this.h, (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemModel) RowCount(parent *QModelIndex) int {
	return (int)(QAbstractItemModel_RowCount(this.h, parent.cPointer()))
}

func (this *QAbstractItemModel) ColumnCount(parent *QModelIndex) int {
	return (int)(QAbstractItemModel_ColumnCount(this.h, parent.cPointer()))
}

func (this *QAbstractItemModel) HasChildren(parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_HasChildren(this.h, parent.cPointer()))
}

func (this *QAbstractItemModel) Data(index *QModelIndex, role int) *QVariant {
	_goptr := newQVariant(QAbstractItemModel_Data(this.h, index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	return (bool)(QAbstractItemModel_SetData(this.h, index.cPointer(), value.cPointer(), (int)(role)))
}

func (this *QAbstractItemModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	_goptr := newQVariant(QAbstractItemModel_HeaderData(this.h, (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemModel) SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {
	return (bool)(QAbstractItemModel_SetHeaderData(this.h, (int)(section), (int)(orientation), value.cPointer(), (int)(role)))
}

func (this *QAbstractItemModel) ItemData(index *QModelIndex) map[int]QVariant {
	var _mm struct_miqt_map = QAbstractItemModel_ItemData(this.h, index.cPointer())
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

func (this *QAbstractItemModel) SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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
	return (bool)(QAbstractItemModel_SetItemData(this.h, index.cPointer(), roles_mm))
}

func (this *QAbstractItemModel) ClearItemData(index *QModelIndex) bool {
	return (bool)(QAbstractItemModel_ClearItemData(this.h, index.cPointer()))
}

func (this *QAbstractItemModel) MimeTypes() []string {
	var _ma struct_miqt_array = QAbstractItemModel_MimeTypes(this.h)
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

func (this *QAbstractItemModel) MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}
	return newQMimeData(QAbstractItemModel_MimeData(this.h, indexes_ma))
}

func (this *QAbstractItemModel) CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_CanDropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QAbstractItemModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QAbstractItemModel) SupportedDropActions() DropAction {
	return (DropAction)(QAbstractItemModel_SupportedDropActions(this.h))
}

func (this *QAbstractItemModel) SupportedDragActions() DropAction {
	return (DropAction)(QAbstractItemModel_SupportedDragActions(this.h))
}

func (this *QAbstractItemModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_InsertRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QAbstractItemModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_InsertColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QAbstractItemModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_RemoveRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QAbstractItemModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_RemoveColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QAbstractItemModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QAbstractItemModel_MoveRows(this.h, sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QAbstractItemModel) MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QAbstractItemModel_MoveColumns(this.h, sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QAbstractItemModel) InsertRow(row int) bool {
	return (bool)(QAbstractItemModel_InsertRow(this.h, (int)(row)))
}

func (this *QAbstractItemModel) InsertColumn(column int) bool {
	return (bool)(QAbstractItemModel_InsertColumn(this.h, (int)(column)))
}

func (this *QAbstractItemModel) RemoveRow(row int) bool {
	return (bool)(QAbstractItemModel_RemoveRow(this.h, (int)(row)))
}

func (this *QAbstractItemModel) RemoveColumn(column int) bool {
	return (bool)(QAbstractItemModel_RemoveColumn(this.h, (int)(column)))
}

func (this *QAbstractItemModel) MoveRow(sourceParent *QModelIndex, sourceRow int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QAbstractItemModel_MoveRow(this.h, sourceParent.cPointer(), (int)(sourceRow), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QAbstractItemModel) MoveColumn(sourceParent *QModelIndex, sourceColumn int, destinationParent *QModelIndex, destinationChild int) bool {
	return (bool)(QAbstractItemModel_MoveColumn(this.h, sourceParent.cPointer(), (int)(sourceColumn), destinationParent.cPointer(), (int)(destinationChild)))
}

func (this *QAbstractItemModel) FetchMore(parent *QModelIndex) {
	QAbstractItemModel_FetchMore(this.h, parent.cPointer())
}

func (this *QAbstractItemModel) CanFetchMore(parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_CanFetchMore(this.h, parent.cPointer()))
}

func (this *QAbstractItemModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QAbstractItemModel_Flags(this.h, index.cPointer()))
}

func (this *QAbstractItemModel) Sort(column int, order SortOrder) {
	QAbstractItemModel_Sort(this.h, (int)(column), (int)(order))
}

func (this *QAbstractItemModel) Buddy(index *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractItemModel_Buddy(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemModel) Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {
	var _ma struct_miqt_array = QAbstractItemModel_Match(this.h, start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QAbstractItemModel) Span(index *QModelIndex) *QSize {
	_goptr := newQSize(QAbstractItemModel_Span(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemModel) RoleNames() map[int][]byte {
	var _mm struct_miqt_map = QAbstractItemModel_RoleNames(this.h)
	_ret := make(map[int][]byte, int(_mm.len))
	_Keys := (*[0xffff]int)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		_entry_Key := (int)(_Keys[i])

		var _hashval_bytearray struct_miqt_string = _Values[i]
		_hashval_ret := GoBytes(unsafe.Pointer(_hashval_bytearray.data), int(int64(_hashval_bytearray.len)))
		free(unsafe.Pointer(_hashval_bytearray.data))
		_entry_Value := _hashval_ret
		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QAbstractItemModel) CheckIndex(index *QModelIndex) bool {
	return (bool)(QAbstractItemModel_CheckIndex(this.h, index.cPointer()))
}

func (this *QAbstractItemModel) MultiData(index *QModelIndex, roleDataSpan QModelRoleDataSpan) {
	QAbstractItemModel_MultiData(this.h, index.cPointer(), roleDataSpan.cPointer())
}

func (this *QAbstractItemModel) DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex) {
	QAbstractItemModel_DataChanged(this.h, topLeft.cPointer(), bottomRight.cPointer())
}

func (this *QAbstractItemModel) OnDataChanged(slot func(topLeft *QModelIndex, bottomRight *QModelIndex)) {
	QAbstractItemModel_connect_DataChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_DataChanged
func miqt_exec_callback_QAbstractItemModel_DataChanged(cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(topLeft *QModelIndex, bottomRight *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(topLeft)

	slotval2 := newQModelIndex(bottomRight)

	gofunc(slotval1, slotval2)
}

func (this *QAbstractItemModel) HeaderDataChanged(orientation Orientation, first int, last int) {
	QAbstractItemModel_HeaderDataChanged(this.h, (int)(orientation), (int)(first), (int)(last))
}

func (this *QAbstractItemModel) OnHeaderDataChanged(slot func(orientation Orientation, first int, last int)) {
	QAbstractItemModel_connect_HeaderDataChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_HeaderDataChanged
func miqt_exec_callback_QAbstractItemModel_HeaderDataChanged(cb intptr_t, orientation int, first int, last int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(orientation Orientation, first int, last int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (Orientation)(orientation)

	slotval2 := (int)(first)

	slotval3 := (int)(last)

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QAbstractItemModel) LayoutChanged() {
	QAbstractItemModel_LayoutChanged(this.h)
}

func (this *QAbstractItemModel) OnLayoutChanged(slot func()) {
	QAbstractItemModel_connect_LayoutChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_LayoutChanged
func miqt_exec_callback_QAbstractItemModel_LayoutChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractItemModel) LayoutAboutToBeChanged() {
	QAbstractItemModel_LayoutAboutToBeChanged(this.h)
}

func (this *QAbstractItemModel) OnLayoutAboutToBeChanged(slot func()) {
	QAbstractItemModel_connect_LayoutAboutToBeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_LayoutAboutToBeChanged
func miqt_exec_callback_QAbstractItemModel_LayoutAboutToBeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractItemModel) Submit() bool {
	return (bool)(QAbstractItemModel_Submit(this.h))
}

func (this *QAbstractItemModel) Revert() {
	QAbstractItemModel_Revert(this.h)
}

func QAbstractItemModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractItemModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractItemModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractItemModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractItemModel) HasIndex3(row int, column int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_HasIndex3(this.h, (int)(row), (int)(column), parent.cPointer()))
}

func (this *QAbstractItemModel) InsertRow2(row int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_InsertRow2(this.h, (int)(row), parent.cPointer()))
}

func (this *QAbstractItemModel) InsertColumn2(column int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_InsertColumn2(this.h, (int)(column), parent.cPointer()))
}

func (this *QAbstractItemModel) RemoveRow2(row int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_RemoveRow2(this.h, (int)(row), parent.cPointer()))
}

func (this *QAbstractItemModel) RemoveColumn2(column int, parent *QModelIndex) bool {
	return (bool)(QAbstractItemModel_RemoveColumn2(this.h, (int)(column), parent.cPointer()))
}

func (this *QAbstractItemModel) CheckIndex2(index *QModelIndex, options CheckIndexOptions) bool {
	return (bool)(QAbstractItemModel_CheckIndex2(this.h, index.cPointer(), options))
}

func (this *QAbstractItemModel) DataChanged3(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}
	QAbstractItemModel_DataChanged3(this.h, topLeft.cPointer(), bottomRight.cPointer(), roles_ma)
}

func (this *QAbstractItemModel) OnDataChanged3(slot func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	QAbstractItemModel_connect_DataChanged3(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_DataChanged3
func miqt_exec_callback_QAbstractItemModel_DataChanged3(cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int))
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

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QAbstractItemModel) LayoutChanged1(parents []QPersistentModelIndex) {
	parents_CArray := (*[0xffff]*QPersistentModelIndex)(malloc(size_t(8 * len(parents))))
	defer free(unsafe.Pointer(parents_CArray))
	for i := range parents {
		parents_CArray[i] = parents[i].cPointer()
	}
	parents_ma := struct_miqt_array{len: size_t(len(parents)), data: unsafe.Pointer(parents_CArray)}
	QAbstractItemModel_LayoutChanged1(this.h, parents_ma)
}

func (this *QAbstractItemModel) OnLayoutChanged1(slot func(parents []QPersistentModelIndex)) {
	QAbstractItemModel_connect_LayoutChanged1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_LayoutChanged1
func miqt_exec_callback_QAbstractItemModel_LayoutChanged1(cb intptr_t, parents struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(parents []QPersistentModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var parents_ma struct_miqt_array = parents
	parents_ret := make([]QPersistentModelIndex, int(parents_ma.len))
	parents_outCast := (*[0xffff]*QPersistentModelIndex)(unsafe.Pointer(parents_ma.data)) // hey ya
	for i := 0; i < int(parents_ma.len); i++ {
		parents_lv_goptr := newQPersistentModelIndex(parents_outCast[i])
		parents_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		parents_ret[i] = *parents_lv_goptr
	}
	slotval1 := parents_ret

	gofunc(slotval1)
}

func (this *QAbstractItemModel) LayoutChanged2(parents []QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint) {
	parents_CArray := (*[0xffff]*QPersistentModelIndex)(malloc(size_t(8 * len(parents))))
	defer free(unsafe.Pointer(parents_CArray))
	for i := range parents {
		parents_CArray[i] = parents[i].cPointer()
	}
	parents_ma := struct_miqt_array{len: size_t(len(parents)), data: unsafe.Pointer(parents_CArray)}
	QAbstractItemModel_LayoutChanged2(this.h, parents_ma, (int)(hint))
}

func (this *QAbstractItemModel) OnLayoutChanged2(slot func(parents []QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint)) {
	QAbstractItemModel_connect_LayoutChanged2(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_LayoutChanged2
func miqt_exec_callback_QAbstractItemModel_LayoutChanged2(cb intptr_t, parents struct_miqt_array, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(parents []QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var parents_ma struct_miqt_array = parents
	parents_ret := make([]QPersistentModelIndex, int(parents_ma.len))
	parents_outCast := (*[0xffff]*QPersistentModelIndex)(unsafe.Pointer(parents_ma.data)) // hey ya
	for i := 0; i < int(parents_ma.len); i++ {
		parents_lv_goptr := newQPersistentModelIndex(parents_outCast[i])
		parents_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		parents_ret[i] = *parents_lv_goptr
	}
	slotval1 := parents_ret

	slotval2 := (QAbstractItemModel__LayoutChangeHint)(hint)

	gofunc(slotval1, slotval2)
}

func (this *QAbstractItemModel) LayoutAboutToBeChanged1(parents []QPersistentModelIndex) {
	parents_CArray := (*[0xffff]*QPersistentModelIndex)(malloc(size_t(8 * len(parents))))
	defer free(unsafe.Pointer(parents_CArray))
	for i := range parents {
		parents_CArray[i] = parents[i].cPointer()
	}
	parents_ma := struct_miqt_array{len: size_t(len(parents)), data: unsafe.Pointer(parents_CArray)}
	QAbstractItemModel_LayoutAboutToBeChanged1(this.h, parents_ma)
}

func (this *QAbstractItemModel) OnLayoutAboutToBeChanged1(slot func(parents []QPersistentModelIndex)) {
	QAbstractItemModel_connect_LayoutAboutToBeChanged1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_LayoutAboutToBeChanged1
func miqt_exec_callback_QAbstractItemModel_LayoutAboutToBeChanged1(cb intptr_t, parents struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(parents []QPersistentModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var parents_ma struct_miqt_array = parents
	parents_ret := make([]QPersistentModelIndex, int(parents_ma.len))
	parents_outCast := (*[0xffff]*QPersistentModelIndex)(unsafe.Pointer(parents_ma.data)) // hey ya
	for i := 0; i < int(parents_ma.len); i++ {
		parents_lv_goptr := newQPersistentModelIndex(parents_outCast[i])
		parents_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		parents_ret[i] = *parents_lv_goptr
	}
	slotval1 := parents_ret

	gofunc(slotval1)
}

func (this *QAbstractItemModel) LayoutAboutToBeChanged2(parents []QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint) {
	parents_CArray := (*[0xffff]*QPersistentModelIndex)(malloc(size_t(8 * len(parents))))
	defer free(unsafe.Pointer(parents_CArray))
	for i := range parents {
		parents_CArray[i] = parents[i].cPointer()
	}
	parents_ma := struct_miqt_array{len: size_t(len(parents)), data: unsafe.Pointer(parents_CArray)}
	QAbstractItemModel_LayoutAboutToBeChanged2(this.h, parents_ma, (int)(hint))
}

func (this *QAbstractItemModel) OnLayoutAboutToBeChanged2(slot func(parents []QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint)) {
	QAbstractItemModel_connect_LayoutAboutToBeChanged2(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_LayoutAboutToBeChanged2
func miqt_exec_callback_QAbstractItemModel_LayoutAboutToBeChanged2(cb intptr_t, parents struct_miqt_array, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(parents []QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var parents_ma struct_miqt_array = parents
	parents_ret := make([]QPersistentModelIndex, int(parents_ma.len))
	parents_outCast := (*[0xffff]*QPersistentModelIndex)(unsafe.Pointer(parents_ma.data)) // hey ya
	for i := 0; i < int(parents_ma.len); i++ {
		parents_lv_goptr := newQPersistentModelIndex(parents_outCast[i])
		parents_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		parents_ret[i] = *parents_lv_goptr
	}
	slotval1 := parents_ret

	slotval2 := (QAbstractItemModel__LayoutChangeHint)(hint)

	gofunc(slotval1, slotval2)
}

func (this *QAbstractItemModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractItemModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractItemModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_MetaObject
func miqt_exec_callback_QAbstractItemModel_MetaObject(self QAbstractItemModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractItemModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractItemModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractItemModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Metacast
func miqt_exec_callback_QAbstractItemModel_Metacast(self QAbstractItemModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QAbstractTableModel struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractTableModel constructs a new QAbstractTableModel object.
func NewQAbstractTableModel() *QAbstractTableModel {
	ret := newQAbstractTableModel(QAbstractTableModel_new())
	ret.isSubclass = true
	return ret
}

// NewQAbstractTableModel2 constructs a new QAbstractTableModel object.
func NewQAbstractTableModel2(parent *QObject) *QAbstractTableModel {
	ret := newQAbstractTableModel(QAbstractTableModel_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractTableModel) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractTableModel_MetaObject(this.h))
}

func (this *QAbstractTableModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractTableModel_Metacast(this.h, param1_Cstring))
}

func QAbstractTableModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractTableModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractTableModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractTableModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractTableModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractTableModel_Sibling(this.h, (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractTableModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QAbstractTableModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QAbstractTableModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QAbstractTableModel_Flags(this.h, index.cPointer()))
}

func QAbstractTableModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractTableModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractTableModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractTableModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractTableModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractTableModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractTableModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_MetaObject
func miqt_exec_callback_QAbstractTableModel_MetaObject(self QAbstractTableModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractTableModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractTableModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractTableModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Metacast
func miqt_exec_callback_QAbstractTableModel_Metacast(self QAbstractTableModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QAbstractListModel struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractListModel constructs a new QAbstractListModel object.
func NewQAbstractListModel() *QAbstractListModel {
	ret := newQAbstractListModel(QAbstractListModel_new())
	ret.isSubclass = true
	return ret
}

// NewQAbstractListModel2 constructs a new QAbstractListModel object.
func NewQAbstractListModel2(parent *QObject) *QAbstractListModel {
	ret := newQAbstractListModel(QAbstractListModel_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractListModel) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractListModel_MetaObject(this.h))
}

func (this *QAbstractListModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractListModel_Metacast(this.h, param1_Cstring))
}

func QAbstractListModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractListModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractListModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractListModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractListModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QAbstractListModel_Sibling(this.h, (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractListModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QAbstractListModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QAbstractListModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QAbstractListModel_Flags(this.h, index.cPointer()))
}

func QAbstractListModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractListModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractListModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractListModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractListModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractListModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractListModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_MetaObject
func miqt_exec_callback_QAbstractListModel_MetaObject(self QAbstractListModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractListModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractListModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractListModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Metacast
func miqt_exec_callback_QAbstractListModel_Metacast(self QAbstractListModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
