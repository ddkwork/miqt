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

func (this *QAbstractItemModel) OnIndex(slot func(row int, column int, parent *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Index(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Index
func miqt_exec_callback_QAbstractItemModel_Index(self QAbstractItemModel, cb intptr_t, row int, column int, parent *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(row int, column int, parent *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc(slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}
func (this *QAbstractItemModel) OnParent(slot func(child *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Parent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Parent
func miqt_exec_callback_QAbstractItemModel_Parent(self QAbstractItemModel, cb intptr_t, child *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(child *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(child)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemModel) callVirtualBase_Sibling(row int, column int, idx *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractItemModel_virtualbase_Sibling(unsafe.Pointer(this.h), (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemModel) OnSibling(slot func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Sibling(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Sibling
func miqt_exec_callback_QAbstractItemModel_Sibling(self QAbstractItemModel, cb intptr_t, row int, column int, idx *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(idx)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Sibling, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}
func (this *QAbstractItemModel) OnRowCount(slot func(parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_RowCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_RowCount
func miqt_exec_callback_QAbstractItemModel_RowCount(self QAbstractItemModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc(slotval1)

	return (int)(virtualReturn)

}
func (this *QAbstractItemModel) OnColumnCount(slot func(parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_ColumnCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_ColumnCount
func miqt_exec_callback_QAbstractItemModel_ColumnCount(self QAbstractItemModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc(slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_HasChildren(parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_HasChildren(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QAbstractItemModel) OnHasChildren(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_HasChildren(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_HasChildren
func miqt_exec_callback_QAbstractItemModel_HasChildren(self QAbstractItemModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_HasChildren, slotval1)

	return (bool)(virtualReturn)

}
func (this *QAbstractItemModel) OnData(slot func(index *QModelIndex, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Data
func miqt_exec_callback_QAbstractItemModel_Data(self QAbstractItemModel, cb intptr_t, index *QModelIndex, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (int)(role)

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemModel) callVirtualBase_SetData(index *QModelIndex, value *QVariant, role int) bool {

	return (bool)(QAbstractItemModel_virtualbase_SetData(unsafe.Pointer(this.h), index.cPointer(), value.cPointer(), (int)(role)))

}
func (this *QAbstractItemModel) OnSetData(slot func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_SetData
func miqt_exec_callback_QAbstractItemModel_SetData(self QAbstractItemModel, cb intptr_t, index *QModelIndex, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQVariant(value)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_SetData, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_HeaderData(section int, orientation Orientation, role int) *QVariant {

	_goptr := newQVariant(QAbstractItemModel_virtualbase_HeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemModel) OnHeaderData(slot func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_HeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_HeaderData
func miqt_exec_callback_QAbstractItemModel_HeaderData(self QAbstractItemModel, cb intptr_t, section int, orientation int, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_HeaderData, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemModel) callVirtualBase_SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {

	return (bool)(QAbstractItemModel_virtualbase_SetHeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), value.cPointer(), (int)(role)))

}
func (this *QAbstractItemModel) OnSetHeaderData(slot func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_SetHeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_SetHeaderData
func miqt_exec_callback_QAbstractItemModel_SetHeaderData(self QAbstractItemModel, cb intptr_t, section int, orientation int, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := newQVariant(value)

	slotval4 := (int)(role)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_SetHeaderData, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_ItemData(index *QModelIndex) map[int]QVariant {

	var _mm struct_miqt_map = QAbstractItemModel_virtualbase_ItemData(unsafe.Pointer(this.h), index.cPointer())
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
func (this *QAbstractItemModel) OnItemData(slot func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_ItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_ItemData
func miqt_exec_callback_QAbstractItemModel_ItemData(self QAbstractItemModel, cb intptr_t, index *QModelIndex) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_ItemData, slotval1)
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

func (this *QAbstractItemModel) callVirtualBase_SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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

	return (bool)(QAbstractItemModel_virtualbase_SetItemData(unsafe.Pointer(this.h), index.cPointer(), roles_mm))

}
func (this *QAbstractItemModel) OnSetItemData(slot func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_SetItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_SetItemData
func miqt_exec_callback_QAbstractItemModel_SetItemData(self QAbstractItemModel, cb intptr_t, index *QModelIndex, roles struct_miqt_map) bool {
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

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_SetItemData, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_ClearItemData(index *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_ClearItemData(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QAbstractItemModel) OnClearItemData(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_ClearItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_ClearItemData
func miqt_exec_callback_QAbstractItemModel_ClearItemData(self QAbstractItemModel, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_ClearItemData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_MimeTypes() []string {

	var _ma struct_miqt_array = QAbstractItemModel_virtualbase_MimeTypes(unsafe.Pointer(this.h))
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
func (this *QAbstractItemModel) OnMimeTypes(slot func(super func() []string) []string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_MimeTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_MimeTypes
func miqt_exec_callback_QAbstractItemModel_MimeTypes(self QAbstractItemModel, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_MimeTypes)
	virtualReturn_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_i_ms := struct_miqt_string{}
		virtualReturn_i_ms.data = CString(virtualReturn[i])
		virtualReturn_i_ms.len = size_t(len(virtualReturn[i]))
		defer free(unsafe.Pointer(virtualReturn_i_ms.data))
		virtualReturn_CArray[i] = virtualReturn_i_ms
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractItemModel) callVirtualBase_MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}

	return newQMimeData(QAbstractItemModel_virtualbase_MimeData(unsafe.Pointer(this.h), indexes_ma))

}
func (this *QAbstractItemModel) OnMimeData(slot func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_MimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_MimeData
func miqt_exec_callback_QAbstractItemModel_MimeData(self QAbstractItemModel, cb intptr_t, indexes struct_miqt_array) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var indexes_ma struct_miqt_array = indexes
	indexes_ret := make([]QModelIndex, int(indexes_ma.len))
	indexes_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(indexes_ma.data)) // hey ya
	for i := 0; i < int(indexes_ma.len); i++ {
		indexes_lv_goptr := newQModelIndex(indexes_outCast[i])
		indexes_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		indexes_ret[i] = *indexes_lv_goptr
	}
	slotval1 := indexes_ret

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_MimeData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemModel) callVirtualBase_CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_CanDropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QAbstractItemModel) OnCanDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_CanDropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_CanDropMimeData
func miqt_exec_callback_QAbstractItemModel_CanDropMimeData(self QAbstractItemModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_CanDropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_DropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QAbstractItemModel) OnDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_DropMimeData
func miqt_exec_callback_QAbstractItemModel_DropMimeData(self QAbstractItemModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QAbstractItemModel_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QAbstractItemModel) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_SupportedDropActions
func miqt_exec_callback_QAbstractItemModel_SupportedDropActions(self QAbstractItemModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_SupportedDragActions() DropAction {

	return (DropAction)(QAbstractItemModel_virtualbase_SupportedDragActions(unsafe.Pointer(this.h)))

}
func (this *QAbstractItemModel) OnSupportedDragActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_SupportedDragActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_SupportedDragActions
func miqt_exec_callback_QAbstractItemModel_SupportedDragActions(self QAbstractItemModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_SupportedDragActions)

	return (int)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_InsertRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_InsertRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QAbstractItemModel) OnInsertRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_InsertRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_InsertRows
func miqt_exec_callback_QAbstractItemModel_InsertRows(self QAbstractItemModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_InsertRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_InsertColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_InsertColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QAbstractItemModel) OnInsertColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_InsertColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_InsertColumns
func miqt_exec_callback_QAbstractItemModel_InsertColumns(self QAbstractItemModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_InsertColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_RemoveRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_RemoveRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QAbstractItemModel) OnRemoveRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_RemoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_RemoveRows
func miqt_exec_callback_QAbstractItemModel_RemoveRows(self QAbstractItemModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_RemoveRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_RemoveColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_RemoveColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QAbstractItemModel) OnRemoveColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_RemoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_RemoveColumns
func miqt_exec_callback_QAbstractItemModel_RemoveColumns(self QAbstractItemModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_RemoveColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QAbstractItemModel_virtualbase_MoveRows(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QAbstractItemModel) OnMoveRows(slot func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_MoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_MoveRows
func miqt_exec_callback_QAbstractItemModel_MoveRows(self QAbstractItemModel, cb intptr_t, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
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

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_MoveRows, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QAbstractItemModel_virtualbase_MoveColumns(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QAbstractItemModel) OnMoveColumns(slot func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_MoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_MoveColumns
func miqt_exec_callback_QAbstractItemModel_MoveColumns(self QAbstractItemModel, cb intptr_t, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(sourceParent)

	slotval2 := (int)(sourceColumn)

	slotval3 := (int)(count)

	slotval4 := newQModelIndex(destinationParent)

	slotval5 := (int)(destinationChild)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_MoveColumns, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_FetchMore(parent *QModelIndex) {

	QAbstractItemModel_virtualbase_FetchMore(unsafe.Pointer(this.h), parent.cPointer())

}
func (this *QAbstractItemModel) OnFetchMore(slot func(super func(parent *QModelIndex), parent *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_FetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_FetchMore
func miqt_exec_callback_QAbstractItemModel_FetchMore(self QAbstractItemModel, cb intptr_t, parent *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex), parent *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_FetchMore, slotval1)

}

func (this *QAbstractItemModel) callVirtualBase_CanFetchMore(parent *QModelIndex) bool {

	return (bool)(QAbstractItemModel_virtualbase_CanFetchMore(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QAbstractItemModel) OnCanFetchMore(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_CanFetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_CanFetchMore
func miqt_exec_callback_QAbstractItemModel_CanFetchMore(self QAbstractItemModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_CanFetchMore, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_Flags(index *QModelIndex) ItemFlag {

	return (ItemFlag)(QAbstractItemModel_virtualbase_Flags(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QAbstractItemModel) OnFlags(slot func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Flags(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Flags
func miqt_exec_callback_QAbstractItemModel_Flags(self QAbstractItemModel, cb intptr_t, index *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Flags, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_Sort(column int, order SortOrder) {

	QAbstractItemModel_virtualbase_Sort(unsafe.Pointer(this.h), (int)(column), (int)(order))

}
func (this *QAbstractItemModel) OnSort(slot func(super func(column int, order SortOrder), column int, order SortOrder)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Sort(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Sort
func miqt_exec_callback_QAbstractItemModel_Sort(self QAbstractItemModel, cb intptr_t, column int, order int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, order SortOrder), column int, order SortOrder))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (SortOrder)(order)

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Sort, slotval1, slotval2)

}

func (this *QAbstractItemModel) callVirtualBase_Buddy(index *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractItemModel_virtualbase_Buddy(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemModel) OnBuddy(slot func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Buddy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Buddy
func miqt_exec_callback_QAbstractItemModel_Buddy(self QAbstractItemModel, cb intptr_t, index *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Buddy, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemModel) callVirtualBase_Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {

	var _ma struct_miqt_array = QAbstractItemModel_virtualbase_Match(unsafe.Pointer(this.h), start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QAbstractItemModel) OnMatch(slot func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Match(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Match
func miqt_exec_callback_QAbstractItemModel_Match(self QAbstractItemModel, cb intptr_t, start *QModelIndex, role int, value *QVariant, hits int, flags int) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(start)

	slotval2 := (int)(role)

	slotval3 := newQVariant(value)

	slotval4 := (int)(hits)

	slotval5 := (MatchFlag)(flags)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Match, slotval1, slotval2, slotval3, slotval4, slotval5)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractItemModel) callVirtualBase_Span(index *QModelIndex) *QSize {

	_goptr := newQSize(QAbstractItemModel_virtualbase_Span(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemModel) OnSpan(slot func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Span(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Span
func miqt_exec_callback_QAbstractItemModel_Span(self QAbstractItemModel, cb intptr_t, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Span, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemModel) callVirtualBase_RoleNames() map[int][]byte {

	var _mm struct_miqt_map = QAbstractItemModel_virtualbase_RoleNames(unsafe.Pointer(this.h))
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
func (this *QAbstractItemModel) OnRoleNames(slot func(super func() map[int][]byte) map[int][]byte) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_RoleNames(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_RoleNames
func miqt_exec_callback_QAbstractItemModel_RoleNames(self QAbstractItemModel, cb intptr_t) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() map[int][]byte) map[int][]byte)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_RoleNames)
	virtualReturn_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Keys_CArray))
	virtualReturn_Values_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Values_CArray))
	virtualReturn_ctr := 0
	for virtualReturn_k, virtualReturn_v := range virtualReturn {
		virtualReturn_Keys_CArray[virtualReturn_ctr] = (int)(virtualReturn_k)
		virtualReturn_v_alias := struct_miqt_string{}
		virtualReturn_v_alias.data = (char)(unsafe.Pointer(&virtualReturn_v[0]))
		virtualReturn_v_alias.len = size_t(len(virtualReturn_v))
		virtualReturn_Values_CArray[virtualReturn_ctr] = virtualReturn_v_alias
		virtualReturn_ctr++
	}
	virtualReturn_mm := struct_miqt_map{
		len:    size_t(len(virtualReturn)),
		keys:   unsafe.Pointer(virtualReturn_Keys_CArray),
		values: unsafe.Pointer(virtualReturn_Values_CArray),
	}

	return virtualReturn_mm

}

func (this *QAbstractItemModel) callVirtualBase_MultiData(index *QModelIndex, roleDataSpan QModelRoleDataSpan) {

	QAbstractItemModel_virtualbase_MultiData(unsafe.Pointer(this.h), index.cPointer(), roleDataSpan.cPointer())

}
func (this *QAbstractItemModel) OnMultiData(slot func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_MultiData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_MultiData
func miqt_exec_callback_QAbstractItemModel_MultiData(self QAbstractItemModel, cb intptr_t, index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	roleDataSpan_goptr := newQModelRoleDataSpan(roleDataSpan)
	roleDataSpan_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval2 := *roleDataSpan_goptr

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_MultiData, slotval1, slotval2)

}

func (this *QAbstractItemModel) callVirtualBase_Submit() bool {

	return (bool)(QAbstractItemModel_virtualbase_Submit(unsafe.Pointer(this.h)))

}
func (this *QAbstractItemModel) OnSubmit(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Submit(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Submit
func miqt_exec_callback_QAbstractItemModel_Submit(self QAbstractItemModel, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Submit)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_Revert() {

	QAbstractItemModel_virtualbase_Revert(unsafe.Pointer(this.h))

}
func (this *QAbstractItemModel) OnRevert(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Revert(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Revert
func miqt_exec_callback_QAbstractItemModel_Revert(self QAbstractItemModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Revert)

}

func (this *QAbstractItemModel) callVirtualBase_ResetInternalData() {

	QAbstractItemModel_virtualbase_ResetInternalData(unsafe.Pointer(this.h))

}
func (this *QAbstractItemModel) OnResetInternalData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_ResetInternalData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_ResetInternalData
func miqt_exec_callback_QAbstractItemModel_ResetInternalData(self QAbstractItemModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_ResetInternalData)

}

func (this *QAbstractItemModel) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QAbstractItemModel_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QAbstractItemModel) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_Event
func miqt_exec_callback_QAbstractItemModel_Event(self QAbstractItemModel, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QAbstractItemModel_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QAbstractItemModel) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_EventFilter
func miqt_exec_callback_QAbstractItemModel_EventFilter(self QAbstractItemModel, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemModel{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemModel) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QAbstractItemModel_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemModel) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_TimerEvent
func miqt_exec_callback_QAbstractItemModel_TimerEvent(self QAbstractItemModel, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAbstractItemModel) callVirtualBase_ChildEvent(event *QChildEvent) {

	QAbstractItemModel_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemModel) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_ChildEvent
func miqt_exec_callback_QAbstractItemModel_ChildEvent(self QAbstractItemModel, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QAbstractItemModel) callVirtualBase_CustomEvent(event *QEvent) {

	QAbstractItemModel_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemModel) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_CustomEvent
func miqt_exec_callback_QAbstractItemModel_CustomEvent(self QAbstractItemModel, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QAbstractItemModel) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QAbstractItemModel_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QAbstractItemModel) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_ConnectNotify
func miqt_exec_callback_QAbstractItemModel_ConnectNotify(self QAbstractItemModel, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QAbstractItemModel) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QAbstractItemModel_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QAbstractItemModel) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemModel_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemModel_DisconnectNotify
func miqt_exec_callback_QAbstractItemModel_DisconnectNotify(self QAbstractItemModel, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QAbstractItemModel{h: self}).callVirtualBase_DisconnectNotify, slotval1)

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

func (this *QAbstractTableModel) callVirtualBase_Index(row int, column int, parent *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractTableModel_virtualbase_Index(unsafe.Pointer(this.h), (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractTableModel) OnIndex(slot func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Index(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Index
func miqt_exec_callback_QAbstractTableModel_Index(self QAbstractTableModel, cb intptr_t, row int, column int, parent *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Index, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractTableModel) callVirtualBase_Sibling(row int, column int, idx *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractTableModel_virtualbase_Sibling(unsafe.Pointer(this.h), (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractTableModel) OnSibling(slot func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Sibling(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Sibling
func miqt_exec_callback_QAbstractTableModel_Sibling(self QAbstractTableModel, cb intptr_t, row int, column int, idx *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(idx)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Sibling, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractTableModel) callVirtualBase_DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_DropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QAbstractTableModel) OnDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_DropMimeData
func miqt_exec_callback_QAbstractTableModel_DropMimeData(self QAbstractTableModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_Flags(index *QModelIndex) ItemFlag {

	return (ItemFlag)(QAbstractTableModel_virtualbase_Flags(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QAbstractTableModel) OnFlags(slot func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Flags(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Flags
func miqt_exec_callback_QAbstractTableModel_Flags(self QAbstractTableModel, cb intptr_t, index *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Flags, slotval1)

	return (int)(virtualReturn)

}
func (this *QAbstractTableModel) OnRowCount(slot func(parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_RowCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_RowCount
func miqt_exec_callback_QAbstractTableModel_RowCount(self QAbstractTableModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc(slotval1)

	return (int)(virtualReturn)

}
func (this *QAbstractTableModel) OnColumnCount(slot func(parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_ColumnCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_ColumnCount
func miqt_exec_callback_QAbstractTableModel_ColumnCount(self QAbstractTableModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc(slotval1)

	return (int)(virtualReturn)

}
func (this *QAbstractTableModel) OnData(slot func(index *QModelIndex, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Data
func miqt_exec_callback_QAbstractTableModel_Data(self QAbstractTableModel, cb intptr_t, index *QModelIndex, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (int)(role)

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QAbstractTableModel) callVirtualBase_SetData(index *QModelIndex, value *QVariant, role int) bool {

	return (bool)(QAbstractTableModel_virtualbase_SetData(unsafe.Pointer(this.h), index.cPointer(), value.cPointer(), (int)(role)))

}
func (this *QAbstractTableModel) OnSetData(slot func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_SetData
func miqt_exec_callback_QAbstractTableModel_SetData(self QAbstractTableModel, cb intptr_t, index *QModelIndex, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQVariant(value)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_SetData, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_HeaderData(section int, orientation Orientation, role int) *QVariant {

	_goptr := newQVariant(QAbstractTableModel_virtualbase_HeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractTableModel) OnHeaderData(slot func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_HeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_HeaderData
func miqt_exec_callback_QAbstractTableModel_HeaderData(self QAbstractTableModel, cb intptr_t, section int, orientation int, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_HeaderData, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractTableModel) callVirtualBase_SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {

	return (bool)(QAbstractTableModel_virtualbase_SetHeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), value.cPointer(), (int)(role)))

}
func (this *QAbstractTableModel) OnSetHeaderData(slot func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_SetHeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_SetHeaderData
func miqt_exec_callback_QAbstractTableModel_SetHeaderData(self QAbstractTableModel, cb intptr_t, section int, orientation int, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := newQVariant(value)

	slotval4 := (int)(role)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_SetHeaderData, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_ItemData(index *QModelIndex) map[int]QVariant {

	var _mm struct_miqt_map = QAbstractTableModel_virtualbase_ItemData(unsafe.Pointer(this.h), index.cPointer())
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
func (this *QAbstractTableModel) OnItemData(slot func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_ItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_ItemData
func miqt_exec_callback_QAbstractTableModel_ItemData(self QAbstractTableModel, cb intptr_t, index *QModelIndex) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_ItemData, slotval1)
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

func (this *QAbstractTableModel) callVirtualBase_SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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

	return (bool)(QAbstractTableModel_virtualbase_SetItemData(unsafe.Pointer(this.h), index.cPointer(), roles_mm))

}
func (this *QAbstractTableModel) OnSetItemData(slot func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_SetItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_SetItemData
func miqt_exec_callback_QAbstractTableModel_SetItemData(self QAbstractTableModel, cb intptr_t, index *QModelIndex, roles struct_miqt_map) bool {
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

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_SetItemData, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_ClearItemData(index *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_ClearItemData(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QAbstractTableModel) OnClearItemData(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_ClearItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_ClearItemData
func miqt_exec_callback_QAbstractTableModel_ClearItemData(self QAbstractTableModel, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_ClearItemData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_MimeTypes() []string {

	var _ma struct_miqt_array = QAbstractTableModel_virtualbase_MimeTypes(unsafe.Pointer(this.h))
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
func (this *QAbstractTableModel) OnMimeTypes(slot func(super func() []string) []string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_MimeTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_MimeTypes
func miqt_exec_callback_QAbstractTableModel_MimeTypes(self QAbstractTableModel, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_MimeTypes)
	virtualReturn_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_i_ms := struct_miqt_string{}
		virtualReturn_i_ms.data = CString(virtualReturn[i])
		virtualReturn_i_ms.len = size_t(len(virtualReturn[i]))
		defer free(unsafe.Pointer(virtualReturn_i_ms.data))
		virtualReturn_CArray[i] = virtualReturn_i_ms
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractTableModel) callVirtualBase_MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}

	return newQMimeData(QAbstractTableModel_virtualbase_MimeData(unsafe.Pointer(this.h), indexes_ma))

}
func (this *QAbstractTableModel) OnMimeData(slot func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_MimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_MimeData
func miqt_exec_callback_QAbstractTableModel_MimeData(self QAbstractTableModel, cb intptr_t, indexes struct_miqt_array) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var indexes_ma struct_miqt_array = indexes
	indexes_ret := make([]QModelIndex, int(indexes_ma.len))
	indexes_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(indexes_ma.data)) // hey ya
	for i := 0; i < int(indexes_ma.len); i++ {
		indexes_lv_goptr := newQModelIndex(indexes_outCast[i])
		indexes_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		indexes_ret[i] = *indexes_lv_goptr
	}
	slotval1 := indexes_ret

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_MimeData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractTableModel) callVirtualBase_CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_CanDropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QAbstractTableModel) OnCanDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_CanDropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_CanDropMimeData
func miqt_exec_callback_QAbstractTableModel_CanDropMimeData(self QAbstractTableModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_CanDropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QAbstractTableModel_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QAbstractTableModel) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_SupportedDropActions
func miqt_exec_callback_QAbstractTableModel_SupportedDropActions(self QAbstractTableModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_SupportedDragActions() DropAction {

	return (DropAction)(QAbstractTableModel_virtualbase_SupportedDragActions(unsafe.Pointer(this.h)))

}
func (this *QAbstractTableModel) OnSupportedDragActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_SupportedDragActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_SupportedDragActions
func miqt_exec_callback_QAbstractTableModel_SupportedDragActions(self QAbstractTableModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_SupportedDragActions)

	return (int)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_InsertRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_InsertRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QAbstractTableModel) OnInsertRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_InsertRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_InsertRows
func miqt_exec_callback_QAbstractTableModel_InsertRows(self QAbstractTableModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_InsertRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_InsertColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_InsertColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QAbstractTableModel) OnInsertColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_InsertColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_InsertColumns
func miqt_exec_callback_QAbstractTableModel_InsertColumns(self QAbstractTableModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_InsertColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_RemoveRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_RemoveRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QAbstractTableModel) OnRemoveRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_RemoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_RemoveRows
func miqt_exec_callback_QAbstractTableModel_RemoveRows(self QAbstractTableModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_RemoveRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_RemoveColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_RemoveColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QAbstractTableModel) OnRemoveColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_RemoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_RemoveColumns
func miqt_exec_callback_QAbstractTableModel_RemoveColumns(self QAbstractTableModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_RemoveColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QAbstractTableModel_virtualbase_MoveRows(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QAbstractTableModel) OnMoveRows(slot func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_MoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_MoveRows
func miqt_exec_callback_QAbstractTableModel_MoveRows(self QAbstractTableModel, cb intptr_t, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
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

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_MoveRows, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QAbstractTableModel_virtualbase_MoveColumns(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QAbstractTableModel) OnMoveColumns(slot func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_MoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_MoveColumns
func miqt_exec_callback_QAbstractTableModel_MoveColumns(self QAbstractTableModel, cb intptr_t, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(sourceParent)

	slotval2 := (int)(sourceColumn)

	slotval3 := (int)(count)

	slotval4 := newQModelIndex(destinationParent)

	slotval5 := (int)(destinationChild)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_MoveColumns, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_FetchMore(parent *QModelIndex) {

	QAbstractTableModel_virtualbase_FetchMore(unsafe.Pointer(this.h), parent.cPointer())

}
func (this *QAbstractTableModel) OnFetchMore(slot func(super func(parent *QModelIndex), parent *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_FetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_FetchMore
func miqt_exec_callback_QAbstractTableModel_FetchMore(self QAbstractTableModel, cb intptr_t, parent *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex), parent *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	gofunc((&QAbstractTableModel{h: self}).callVirtualBase_FetchMore, slotval1)

}

func (this *QAbstractTableModel) callVirtualBase_CanFetchMore(parent *QModelIndex) bool {

	return (bool)(QAbstractTableModel_virtualbase_CanFetchMore(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QAbstractTableModel) OnCanFetchMore(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_CanFetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_CanFetchMore
func miqt_exec_callback_QAbstractTableModel_CanFetchMore(self QAbstractTableModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_CanFetchMore, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_Sort(column int, order SortOrder) {

	QAbstractTableModel_virtualbase_Sort(unsafe.Pointer(this.h), (int)(column), (int)(order))

}
func (this *QAbstractTableModel) OnSort(slot func(super func(column int, order SortOrder), column int, order SortOrder)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Sort(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Sort
func miqt_exec_callback_QAbstractTableModel_Sort(self QAbstractTableModel, cb intptr_t, column int, order int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, order SortOrder), column int, order SortOrder))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (SortOrder)(order)

	gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Sort, slotval1, slotval2)

}

func (this *QAbstractTableModel) callVirtualBase_Buddy(index *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractTableModel_virtualbase_Buddy(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractTableModel) OnBuddy(slot func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Buddy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Buddy
func miqt_exec_callback_QAbstractTableModel_Buddy(self QAbstractTableModel, cb intptr_t, index *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Buddy, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractTableModel) callVirtualBase_Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {

	var _ma struct_miqt_array = QAbstractTableModel_virtualbase_Match(unsafe.Pointer(this.h), start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QAbstractTableModel) OnMatch(slot func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Match(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Match
func miqt_exec_callback_QAbstractTableModel_Match(self QAbstractTableModel, cb intptr_t, start *QModelIndex, role int, value *QVariant, hits int, flags int) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(start)

	slotval2 := (int)(role)

	slotval3 := newQVariant(value)

	slotval4 := (int)(hits)

	slotval5 := (MatchFlag)(flags)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Match, slotval1, slotval2, slotval3, slotval4, slotval5)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractTableModel) callVirtualBase_Span(index *QModelIndex) *QSize {

	_goptr := newQSize(QAbstractTableModel_virtualbase_Span(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractTableModel) OnSpan(slot func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Span(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Span
func miqt_exec_callback_QAbstractTableModel_Span(self QAbstractTableModel, cb intptr_t, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Span, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractTableModel) callVirtualBase_RoleNames() map[int][]byte {

	var _mm struct_miqt_map = QAbstractTableModel_virtualbase_RoleNames(unsafe.Pointer(this.h))
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
func (this *QAbstractTableModel) OnRoleNames(slot func(super func() map[int][]byte) map[int][]byte) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_RoleNames(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_RoleNames
func miqt_exec_callback_QAbstractTableModel_RoleNames(self QAbstractTableModel, cb intptr_t) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() map[int][]byte) map[int][]byte)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_RoleNames)
	virtualReturn_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Keys_CArray))
	virtualReturn_Values_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Values_CArray))
	virtualReturn_ctr := 0
	for virtualReturn_k, virtualReturn_v := range virtualReturn {
		virtualReturn_Keys_CArray[virtualReturn_ctr] = (int)(virtualReturn_k)
		virtualReturn_v_alias := struct_miqt_string{}
		virtualReturn_v_alias.data = (char)(unsafe.Pointer(&virtualReturn_v[0]))
		virtualReturn_v_alias.len = size_t(len(virtualReturn_v))
		virtualReturn_Values_CArray[virtualReturn_ctr] = virtualReturn_v_alias
		virtualReturn_ctr++
	}
	virtualReturn_mm := struct_miqt_map{
		len:    size_t(len(virtualReturn)),
		keys:   unsafe.Pointer(virtualReturn_Keys_CArray),
		values: unsafe.Pointer(virtualReturn_Values_CArray),
	}

	return virtualReturn_mm

}

func (this *QAbstractTableModel) callVirtualBase_MultiData(index *QModelIndex, roleDataSpan QModelRoleDataSpan) {

	QAbstractTableModel_virtualbase_MultiData(unsafe.Pointer(this.h), index.cPointer(), roleDataSpan.cPointer())

}
func (this *QAbstractTableModel) OnMultiData(slot func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_MultiData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_MultiData
func miqt_exec_callback_QAbstractTableModel_MultiData(self QAbstractTableModel, cb intptr_t, index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	roleDataSpan_goptr := newQModelRoleDataSpan(roleDataSpan)
	roleDataSpan_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval2 := *roleDataSpan_goptr

	gofunc((&QAbstractTableModel{h: self}).callVirtualBase_MultiData, slotval1, slotval2)

}

func (this *QAbstractTableModel) callVirtualBase_Submit() bool {

	return (bool)(QAbstractTableModel_virtualbase_Submit(unsafe.Pointer(this.h)))

}
func (this *QAbstractTableModel) OnSubmit(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Submit(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Submit
func miqt_exec_callback_QAbstractTableModel_Submit(self QAbstractTableModel, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Submit)

	return (bool)(virtualReturn)

}

func (this *QAbstractTableModel) callVirtualBase_Revert() {

	QAbstractTableModel_virtualbase_Revert(unsafe.Pointer(this.h))

}
func (this *QAbstractTableModel) OnRevert(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_Revert(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_Revert
func miqt_exec_callback_QAbstractTableModel_Revert(self QAbstractTableModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractTableModel{h: self}).callVirtualBase_Revert)

}

func (this *QAbstractTableModel) callVirtualBase_ResetInternalData() {

	QAbstractTableModel_virtualbase_ResetInternalData(unsafe.Pointer(this.h))

}
func (this *QAbstractTableModel) OnResetInternalData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTableModel_override_virtual_ResetInternalData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTableModel_ResetInternalData
func miqt_exec_callback_QAbstractTableModel_ResetInternalData(self QAbstractTableModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractTableModel{h: self}).callVirtualBase_ResetInternalData)

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

func (this *QAbstractListModel) callVirtualBase_Index(row int, column int, parent *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractListModel_virtualbase_Index(unsafe.Pointer(this.h), (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractListModel) OnIndex(slot func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Index(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Index
func miqt_exec_callback_QAbstractListModel_Index(self QAbstractListModel, cb intptr_t, row int, column int, parent *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Index, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractListModel) callVirtualBase_Sibling(row int, column int, idx *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractListModel_virtualbase_Sibling(unsafe.Pointer(this.h), (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractListModel) OnSibling(slot func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Sibling(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Sibling
func miqt_exec_callback_QAbstractListModel_Sibling(self QAbstractListModel, cb intptr_t, row int, column int, idx *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(idx)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Sibling, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractListModel) callVirtualBase_DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_DropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QAbstractListModel) OnDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_DropMimeData
func miqt_exec_callback_QAbstractListModel_DropMimeData(self QAbstractListModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_Flags(index *QModelIndex) ItemFlag {

	return (ItemFlag)(QAbstractListModel_virtualbase_Flags(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QAbstractListModel) OnFlags(slot func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Flags(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Flags
func miqt_exec_callback_QAbstractListModel_Flags(self QAbstractListModel, cb intptr_t, index *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Flags, slotval1)

	return (int)(virtualReturn)

}
func (this *QAbstractListModel) OnRowCount(slot func(parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_RowCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_RowCount
func miqt_exec_callback_QAbstractListModel_RowCount(self QAbstractListModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc(slotval1)

	return (int)(virtualReturn)

}
func (this *QAbstractListModel) OnData(slot func(index *QModelIndex, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Data
func miqt_exec_callback_QAbstractListModel_Data(self QAbstractListModel, cb intptr_t, index *QModelIndex, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (int)(role)

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QAbstractListModel) callVirtualBase_SetData(index *QModelIndex, value *QVariant, role int) bool {

	return (bool)(QAbstractListModel_virtualbase_SetData(unsafe.Pointer(this.h), index.cPointer(), value.cPointer(), (int)(role)))

}
func (this *QAbstractListModel) OnSetData(slot func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_SetData
func miqt_exec_callback_QAbstractListModel_SetData(self QAbstractListModel, cb intptr_t, index *QModelIndex, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQVariant(value)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_SetData, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_HeaderData(section int, orientation Orientation, role int) *QVariant {

	_goptr := newQVariant(QAbstractListModel_virtualbase_HeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractListModel) OnHeaderData(slot func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_HeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_HeaderData
func miqt_exec_callback_QAbstractListModel_HeaderData(self QAbstractListModel, cb intptr_t, section int, orientation int, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_HeaderData, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractListModel) callVirtualBase_SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {

	return (bool)(QAbstractListModel_virtualbase_SetHeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), value.cPointer(), (int)(role)))

}
func (this *QAbstractListModel) OnSetHeaderData(slot func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_SetHeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_SetHeaderData
func miqt_exec_callback_QAbstractListModel_SetHeaderData(self QAbstractListModel, cb intptr_t, section int, orientation int, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := newQVariant(value)

	slotval4 := (int)(role)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_SetHeaderData, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_ItemData(index *QModelIndex) map[int]QVariant {

	var _mm struct_miqt_map = QAbstractListModel_virtualbase_ItemData(unsafe.Pointer(this.h), index.cPointer())
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
func (this *QAbstractListModel) OnItemData(slot func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_ItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_ItemData
func miqt_exec_callback_QAbstractListModel_ItemData(self QAbstractListModel, cb intptr_t, index *QModelIndex) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_ItemData, slotval1)
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

func (this *QAbstractListModel) callVirtualBase_SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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

	return (bool)(QAbstractListModel_virtualbase_SetItemData(unsafe.Pointer(this.h), index.cPointer(), roles_mm))

}
func (this *QAbstractListModel) OnSetItemData(slot func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_SetItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_SetItemData
func miqt_exec_callback_QAbstractListModel_SetItemData(self QAbstractListModel, cb intptr_t, index *QModelIndex, roles struct_miqt_map) bool {
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

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_SetItemData, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_ClearItemData(index *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_ClearItemData(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QAbstractListModel) OnClearItemData(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_ClearItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_ClearItemData
func miqt_exec_callback_QAbstractListModel_ClearItemData(self QAbstractListModel, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_ClearItemData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_MimeTypes() []string {

	var _ma struct_miqt_array = QAbstractListModel_virtualbase_MimeTypes(unsafe.Pointer(this.h))
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
func (this *QAbstractListModel) OnMimeTypes(slot func(super func() []string) []string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_MimeTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_MimeTypes
func miqt_exec_callback_QAbstractListModel_MimeTypes(self QAbstractListModel, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_MimeTypes)
	virtualReturn_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_i_ms := struct_miqt_string{}
		virtualReturn_i_ms.data = CString(virtualReturn[i])
		virtualReturn_i_ms.len = size_t(len(virtualReturn[i]))
		defer free(unsafe.Pointer(virtualReturn_i_ms.data))
		virtualReturn_CArray[i] = virtualReturn_i_ms
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractListModel) callVirtualBase_MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}

	return newQMimeData(QAbstractListModel_virtualbase_MimeData(unsafe.Pointer(this.h), indexes_ma))

}
func (this *QAbstractListModel) OnMimeData(slot func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_MimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_MimeData
func miqt_exec_callback_QAbstractListModel_MimeData(self QAbstractListModel, cb intptr_t, indexes struct_miqt_array) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var indexes_ma struct_miqt_array = indexes
	indexes_ret := make([]QModelIndex, int(indexes_ma.len))
	indexes_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(indexes_ma.data)) // hey ya
	for i := 0; i < int(indexes_ma.len); i++ {
		indexes_lv_goptr := newQModelIndex(indexes_outCast[i])
		indexes_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		indexes_ret[i] = *indexes_lv_goptr
	}
	slotval1 := indexes_ret

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_MimeData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractListModel) callVirtualBase_CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_CanDropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QAbstractListModel) OnCanDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_CanDropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_CanDropMimeData
func miqt_exec_callback_QAbstractListModel_CanDropMimeData(self QAbstractListModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_CanDropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QAbstractListModel_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QAbstractListModel) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_SupportedDropActions
func miqt_exec_callback_QAbstractListModel_SupportedDropActions(self QAbstractListModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_SupportedDragActions() DropAction {

	return (DropAction)(QAbstractListModel_virtualbase_SupportedDragActions(unsafe.Pointer(this.h)))

}
func (this *QAbstractListModel) OnSupportedDragActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_SupportedDragActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_SupportedDragActions
func miqt_exec_callback_QAbstractListModel_SupportedDragActions(self QAbstractListModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_SupportedDragActions)

	return (int)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_InsertRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_InsertRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QAbstractListModel) OnInsertRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_InsertRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_InsertRows
func miqt_exec_callback_QAbstractListModel_InsertRows(self QAbstractListModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_InsertRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_InsertColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_InsertColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QAbstractListModel) OnInsertColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_InsertColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_InsertColumns
func miqt_exec_callback_QAbstractListModel_InsertColumns(self QAbstractListModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_InsertColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_RemoveRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_RemoveRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QAbstractListModel) OnRemoveRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_RemoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_RemoveRows
func miqt_exec_callback_QAbstractListModel_RemoveRows(self QAbstractListModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_RemoveRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_RemoveColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_RemoveColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QAbstractListModel) OnRemoveColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_RemoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_RemoveColumns
func miqt_exec_callback_QAbstractListModel_RemoveColumns(self QAbstractListModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_RemoveColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QAbstractListModel_virtualbase_MoveRows(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QAbstractListModel) OnMoveRows(slot func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_MoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_MoveRows
func miqt_exec_callback_QAbstractListModel_MoveRows(self QAbstractListModel, cb intptr_t, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
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

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_MoveRows, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QAbstractListModel_virtualbase_MoveColumns(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QAbstractListModel) OnMoveColumns(slot func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_MoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_MoveColumns
func miqt_exec_callback_QAbstractListModel_MoveColumns(self QAbstractListModel, cb intptr_t, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(sourceParent)

	slotval2 := (int)(sourceColumn)

	slotval3 := (int)(count)

	slotval4 := newQModelIndex(destinationParent)

	slotval5 := (int)(destinationChild)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_MoveColumns, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_FetchMore(parent *QModelIndex) {

	QAbstractListModel_virtualbase_FetchMore(unsafe.Pointer(this.h), parent.cPointer())

}
func (this *QAbstractListModel) OnFetchMore(slot func(super func(parent *QModelIndex), parent *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_FetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_FetchMore
func miqt_exec_callback_QAbstractListModel_FetchMore(self QAbstractListModel, cb intptr_t, parent *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex), parent *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	gofunc((&QAbstractListModel{h: self}).callVirtualBase_FetchMore, slotval1)

}

func (this *QAbstractListModel) callVirtualBase_CanFetchMore(parent *QModelIndex) bool {

	return (bool)(QAbstractListModel_virtualbase_CanFetchMore(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QAbstractListModel) OnCanFetchMore(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_CanFetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_CanFetchMore
func miqt_exec_callback_QAbstractListModel_CanFetchMore(self QAbstractListModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_CanFetchMore, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_Sort(column int, order SortOrder) {

	QAbstractListModel_virtualbase_Sort(unsafe.Pointer(this.h), (int)(column), (int)(order))

}
func (this *QAbstractListModel) OnSort(slot func(super func(column int, order SortOrder), column int, order SortOrder)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Sort(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Sort
func miqt_exec_callback_QAbstractListModel_Sort(self QAbstractListModel, cb intptr_t, column int, order int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, order SortOrder), column int, order SortOrder))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (SortOrder)(order)

	gofunc((&QAbstractListModel{h: self}).callVirtualBase_Sort, slotval1, slotval2)

}

func (this *QAbstractListModel) callVirtualBase_Buddy(index *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QAbstractListModel_virtualbase_Buddy(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractListModel) OnBuddy(slot func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Buddy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Buddy
func miqt_exec_callback_QAbstractListModel_Buddy(self QAbstractListModel, cb intptr_t, index *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Buddy, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractListModel) callVirtualBase_Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {

	var _ma struct_miqt_array = QAbstractListModel_virtualbase_Match(unsafe.Pointer(this.h), start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QAbstractListModel) OnMatch(slot func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Match(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Match
func miqt_exec_callback_QAbstractListModel_Match(self QAbstractListModel, cb intptr_t, start *QModelIndex, role int, value *QVariant, hits int, flags int) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(start)

	slotval2 := (int)(role)

	slotval3 := newQVariant(value)

	slotval4 := (int)(hits)

	slotval5 := (MatchFlag)(flags)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Match, slotval1, slotval2, slotval3, slotval4, slotval5)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractListModel) callVirtualBase_Span(index *QModelIndex) *QSize {

	_goptr := newQSize(QAbstractListModel_virtualbase_Span(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractListModel) OnSpan(slot func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Span(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Span
func miqt_exec_callback_QAbstractListModel_Span(self QAbstractListModel, cb intptr_t, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Span, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractListModel) callVirtualBase_RoleNames() map[int][]byte {

	var _mm struct_miqt_map = QAbstractListModel_virtualbase_RoleNames(unsafe.Pointer(this.h))
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
func (this *QAbstractListModel) OnRoleNames(slot func(super func() map[int][]byte) map[int][]byte) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_RoleNames(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_RoleNames
func miqt_exec_callback_QAbstractListModel_RoleNames(self QAbstractListModel, cb intptr_t) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() map[int][]byte) map[int][]byte)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_RoleNames)
	virtualReturn_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Keys_CArray))
	virtualReturn_Values_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Values_CArray))
	virtualReturn_ctr := 0
	for virtualReturn_k, virtualReturn_v := range virtualReturn {
		virtualReturn_Keys_CArray[virtualReturn_ctr] = (int)(virtualReturn_k)
		virtualReturn_v_alias := struct_miqt_string{}
		virtualReturn_v_alias.data = (char)(unsafe.Pointer(&virtualReturn_v[0]))
		virtualReturn_v_alias.len = size_t(len(virtualReturn_v))
		virtualReturn_Values_CArray[virtualReturn_ctr] = virtualReturn_v_alias
		virtualReturn_ctr++
	}
	virtualReturn_mm := struct_miqt_map{
		len:    size_t(len(virtualReturn)),
		keys:   unsafe.Pointer(virtualReturn_Keys_CArray),
		values: unsafe.Pointer(virtualReturn_Values_CArray),
	}

	return virtualReturn_mm

}

func (this *QAbstractListModel) callVirtualBase_MultiData(index *QModelIndex, roleDataSpan QModelRoleDataSpan) {

	QAbstractListModel_virtualbase_MultiData(unsafe.Pointer(this.h), index.cPointer(), roleDataSpan.cPointer())

}
func (this *QAbstractListModel) OnMultiData(slot func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_MultiData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_MultiData
func miqt_exec_callback_QAbstractListModel_MultiData(self QAbstractListModel, cb intptr_t, index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	roleDataSpan_goptr := newQModelRoleDataSpan(roleDataSpan)
	roleDataSpan_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval2 := *roleDataSpan_goptr

	gofunc((&QAbstractListModel{h: self}).callVirtualBase_MultiData, slotval1, slotval2)

}

func (this *QAbstractListModel) callVirtualBase_Submit() bool {

	return (bool)(QAbstractListModel_virtualbase_Submit(unsafe.Pointer(this.h)))

}
func (this *QAbstractListModel) OnSubmit(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Submit(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Submit
func miqt_exec_callback_QAbstractListModel_Submit(self QAbstractListModel, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractListModel{h: self}).callVirtualBase_Submit)

	return (bool)(virtualReturn)

}

func (this *QAbstractListModel) callVirtualBase_Revert() {

	QAbstractListModel_virtualbase_Revert(unsafe.Pointer(this.h))

}
func (this *QAbstractListModel) OnRevert(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_Revert(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_Revert
func miqt_exec_callback_QAbstractListModel_Revert(self QAbstractListModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractListModel{h: self}).callVirtualBase_Revert)

}

func (this *QAbstractListModel) callVirtualBase_ResetInternalData() {

	QAbstractListModel_virtualbase_ResetInternalData(unsafe.Pointer(this.h))

}
func (this *QAbstractListModel) OnResetInternalData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractListModel_override_virtual_ResetInternalData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractListModel_ResetInternalData
func miqt_exec_callback_QAbstractListModel_ResetInternalData(self QAbstractListModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractListModel{h: self}).callVirtualBase_ResetInternalData)

}
