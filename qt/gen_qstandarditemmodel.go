package qt

import (
	"unsafe"
)

type QStandardItem__ItemType int

const (
	QStandardItem__Type     QStandardItem__ItemType = 0
	QStandardItem__UserType QStandardItem__ItemType = 1000
)

type QStandardItem struct {
	h          uintptr
	isSubclass bool
}

// NewQStandardItem constructs a new QStandardItem object.
func NewQStandardItem() *QStandardItem {
	ret := newQStandardItem(QStandardItem_new())
	ret.isSubclass = true
	return ret
}

// NewQStandardItem2 constructs a new QStandardItem object.
func NewQStandardItem2(text string) *QStandardItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQStandardItem(QStandardItem_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQStandardItem3 constructs a new QStandardItem object.
func NewQStandardItem3(icon *QIcon, text string) *QStandardItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQStandardItem(QStandardItem_new3(icon.cPointer(), text_ms))
	ret.isSubclass = true
	return ret
}

// NewQStandardItem4 constructs a new QStandardItem object.
func NewQStandardItem4(rows int) *QStandardItem {
	ret := newQStandardItem(QStandardItem_new4((int)(rows)))
	ret.isSubclass = true
	return ret
}

// NewQStandardItem5 constructs a new QStandardItem object.
func NewQStandardItem5(rows int, columns int) *QStandardItem {
	ret := newQStandardItem(QStandardItem_new5((int)(rows), (int)(columns)))
	ret.isSubclass = true
	return ret
}

func (this *QStandardItem) Data(role int) *QVariant {
	_goptr := newQVariant(QStandardItem_Data(this.h, (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItem) MultiData(roleDataSpan QModelRoleDataSpan) {
	QStandardItem_MultiData(this.h, roleDataSpan.cPointer())
}

func (this *QStandardItem) SetData(value *QVariant, role int) {
	QStandardItem_SetData(this.h, value.cPointer(), (int)(role))
}

func (this *QStandardItem) ClearData() {
	QStandardItem_ClearData(this.h)
}

func (this *QStandardItem) Text() string {
	var _ms struct_miqt_string = QStandardItem_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItem) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStandardItem_SetText(this.h, text_ms)
}

func (this *QStandardItem) Icon() *QIcon {
	_goptr := newQIcon(QStandardItem_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItem) SetIcon(icon *QIcon) {
	QStandardItem_SetIcon(this.h, icon.cPointer())
}

func (this *QStandardItem) ToolTip() string {
	var _ms struct_miqt_string = QStandardItem_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItem) SetToolTip(toolTip string) {
	toolTip_ms := struct_miqt_string{}
	toolTip_ms.data = CString(toolTip)
	toolTip_ms.len = size_t(len(toolTip))
	defer free(unsafe.Pointer(toolTip_ms.data))
	QStandardItem_SetToolTip(this.h, toolTip_ms)
}

func (this *QStandardItem) StatusTip() string {
	var _ms struct_miqt_string = QStandardItem_StatusTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItem) SetStatusTip(statusTip string) {
	statusTip_ms := struct_miqt_string{}
	statusTip_ms.data = CString(statusTip)
	statusTip_ms.len = size_t(len(statusTip))
	defer free(unsafe.Pointer(statusTip_ms.data))
	QStandardItem_SetStatusTip(this.h, statusTip_ms)
}

func (this *QStandardItem) WhatsThis() string {
	var _ms struct_miqt_string = QStandardItem_WhatsThis(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItem) SetWhatsThis(whatsThis string) {
	whatsThis_ms := struct_miqt_string{}
	whatsThis_ms.data = CString(whatsThis)
	whatsThis_ms.len = size_t(len(whatsThis))
	defer free(unsafe.Pointer(whatsThis_ms.data))
	QStandardItem_SetWhatsThis(this.h, whatsThis_ms)
}

func (this *QStandardItem) SizeHint() *QSize {
	_goptr := newQSize(QStandardItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItem) SetSizeHint(sizeHint *QSize) {
	QStandardItem_SetSizeHint(this.h, sizeHint.cPointer())
}

func (this *QStandardItem) Font() *QFont {
	_goptr := newQFont(QStandardItem_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItem) SetFont(font *QFont) {
	QStandardItem_SetFont(this.h, font.cPointer())
}

func (this *QStandardItem) TextAlignment() AlignmentFlag {
	return (AlignmentFlag)(QStandardItem_TextAlignment(this.h))
}

func (this *QStandardItem) SetTextAlignment(textAlignment AlignmentFlag) {
	QStandardItem_SetTextAlignment(this.h, (int)(textAlignment))
}

func (this *QStandardItem) Background() *QBrush {
	_goptr := newQBrush(QStandardItem_Background(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItem) SetBackground(brush *QBrush) {
	QStandardItem_SetBackground(this.h, brush.cPointer())
}

func (this *QStandardItem) Foreground() *QBrush {
	_goptr := newQBrush(QStandardItem_Foreground(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItem) SetForeground(brush *QBrush) {
	QStandardItem_SetForeground(this.h, brush.cPointer())
}

func (this *QStandardItem) CheckState() CheckState {
	return (CheckState)(QStandardItem_CheckState(this.h))
}

func (this *QStandardItem) SetCheckState(checkState CheckState) {
	QStandardItem_SetCheckState(this.h, (int)(checkState))
}

func (this *QStandardItem) AccessibleText() string {
	var _ms struct_miqt_string = QStandardItem_AccessibleText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItem) SetAccessibleText(accessibleText string) {
	accessibleText_ms := struct_miqt_string{}
	accessibleText_ms.data = CString(accessibleText)
	accessibleText_ms.len = size_t(len(accessibleText))
	defer free(unsafe.Pointer(accessibleText_ms.data))
	QStandardItem_SetAccessibleText(this.h, accessibleText_ms)
}

func (this *QStandardItem) AccessibleDescription() string {
	var _ms struct_miqt_string = QStandardItem_AccessibleDescription(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	accessibleDescription_ms := struct_miqt_string{}
	accessibleDescription_ms.data = CString(accessibleDescription)
	accessibleDescription_ms.len = size_t(len(accessibleDescription))
	defer free(unsafe.Pointer(accessibleDescription_ms.data))
	QStandardItem_SetAccessibleDescription(this.h, accessibleDescription_ms)
}

func (this *QStandardItem) Flags() ItemFlag {
	return (ItemFlag)(QStandardItem_Flags(this.h))
}

func (this *QStandardItem) SetFlags(flags ItemFlag) {
	QStandardItem_SetFlags(this.h, (int)(flags))
}

func (this *QStandardItem) IsEnabled() bool {
	return (bool)(QStandardItem_IsEnabled(this.h))
}

func (this *QStandardItem) SetEnabled(enabled bool) {
	QStandardItem_SetEnabled(this.h, (bool)(enabled))
}

func (this *QStandardItem) IsEditable() bool {
	return (bool)(QStandardItem_IsEditable(this.h))
}

func (this *QStandardItem) SetEditable(editable bool) {
	QStandardItem_SetEditable(this.h, (bool)(editable))
}

func (this *QStandardItem) IsSelectable() bool {
	return (bool)(QStandardItem_IsSelectable(this.h))
}

func (this *QStandardItem) SetSelectable(selectable bool) {
	QStandardItem_SetSelectable(this.h, (bool)(selectable))
}

func (this *QStandardItem) IsCheckable() bool {
	return (bool)(QStandardItem_IsCheckable(this.h))
}

func (this *QStandardItem) SetCheckable(checkable bool) {
	QStandardItem_SetCheckable(this.h, (bool)(checkable))
}

func (this *QStandardItem) IsAutoTristate() bool {
	return (bool)(QStandardItem_IsAutoTristate(this.h))
}

func (this *QStandardItem) SetAutoTristate(tristate bool) {
	QStandardItem_SetAutoTristate(this.h, (bool)(tristate))
}

func (this *QStandardItem) IsUserTristate() bool {
	return (bool)(QStandardItem_IsUserTristate(this.h))
}

func (this *QStandardItem) SetUserTristate(tristate bool) {
	QStandardItem_SetUserTristate(this.h, (bool)(tristate))
}

func (this *QStandardItem) IsDragEnabled() bool {
	return (bool)(QStandardItem_IsDragEnabled(this.h))
}

func (this *QStandardItem) SetDragEnabled(dragEnabled bool) {
	QStandardItem_SetDragEnabled(this.h, (bool)(dragEnabled))
}

func (this *QStandardItem) IsDropEnabled() bool {
	return (bool)(QStandardItem_IsDropEnabled(this.h))
}

func (this *QStandardItem) SetDropEnabled(dropEnabled bool) {
	QStandardItem_SetDropEnabled(this.h, (bool)(dropEnabled))
}

func (this *QStandardItem) Parent() *QStandardItem {
	return newQStandardItem(QStandardItem_Parent(this.h))
}

func (this *QStandardItem) Row() int {
	return (int)(QStandardItem_Row(this.h))
}

func (this *QStandardItem) Column() int {
	return (int)(QStandardItem_Column(this.h))
}

func (this *QStandardItem) Index() *QModelIndex {
	_goptr := newQModelIndex(QStandardItem_Index(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItem) Model() *QStandardItemModel {
	return newQStandardItemModel(QStandardItem_Model(this.h))
}

func (this *QStandardItem) RowCount() int {
	return (int)(QStandardItem_RowCount(this.h))
}

func (this *QStandardItem) SetRowCount(rows int) {
	QStandardItem_SetRowCount(this.h, (int)(rows))
}

func (this *QStandardItem) ColumnCount() int {
	return (int)(QStandardItem_ColumnCount(this.h))
}

func (this *QStandardItem) SetColumnCount(columns int) {
	QStandardItem_SetColumnCount(this.h, (int)(columns))
}

func (this *QStandardItem) HasChildren() bool {
	return (bool)(QStandardItem_HasChildren(this.h))
}

func (this *QStandardItem) Child(row int) *QStandardItem {
	return newQStandardItem(QStandardItem_Child(this.h, (int)(row)))
}

func (this *QStandardItem) SetChild(row int, column int, item *QStandardItem) {
	QStandardItem_SetChild(this.h, (int)(row), (int)(column), item.cPointer())
}

func (this *QStandardItem) SetChild2(row int, item *QStandardItem) {
	QStandardItem_SetChild2(this.h, (int)(row), item.cPointer())
}

func (this *QStandardItem) InsertRow(row int, items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItem_InsertRow(this.h, (int)(row), items_ma)
}

func (this *QStandardItem) InsertColumn(column int, items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItem_InsertColumn(this.h, (int)(column), items_ma)
}

func (this *QStandardItem) InsertRows(row int, items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItem_InsertRows(this.h, (int)(row), items_ma)
}

func (this *QStandardItem) InsertRows2(row int, count int) {
	QStandardItem_InsertRows2(this.h, (int)(row), (int)(count))
}

func (this *QStandardItem) InsertColumns(column int, count int) {
	QStandardItem_InsertColumns(this.h, (int)(column), (int)(count))
}

func (this *QStandardItem) RemoveRow(row int) {
	QStandardItem_RemoveRow(this.h, (int)(row))
}

func (this *QStandardItem) RemoveColumn(column int) {
	QStandardItem_RemoveColumn(this.h, (int)(column))
}

func (this *QStandardItem) RemoveRows(row int, count int) {
	QStandardItem_RemoveRows(this.h, (int)(row), (int)(count))
}

func (this *QStandardItem) RemoveColumns(column int, count int) {
	QStandardItem_RemoveColumns(this.h, (int)(column), (int)(count))
}

func (this *QStandardItem) AppendRow(items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItem_AppendRow(this.h, items_ma)
}

func (this *QStandardItem) AppendRows(items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItem_AppendRows(this.h, items_ma)
}

func (this *QStandardItem) AppendColumn(items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItem_AppendColumn(this.h, items_ma)
}

func (this *QStandardItem) InsertRow2(row int, item *QStandardItem) {
	QStandardItem_InsertRow2(this.h, (int)(row), item.cPointer())
}

func (this *QStandardItem) AppendRowWithItem(item *QStandardItem) {
	QStandardItem_AppendRowWithItem(this.h, item.cPointer())
}

func (this *QStandardItem) TakeChild(row int) *QStandardItem {
	return newQStandardItem(QStandardItem_TakeChild(this.h, (int)(row)))
}

func (this *QStandardItem) TakeRow(row int) []*QStandardItem {
	var _ma struct_miqt_array = QStandardItem_TakeRow(this.h, (int)(row))
	_ret := make([]*QStandardItem, int(_ma.len))
	_outCast := (*[0xffff]*QStandardItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQStandardItem(_outCast[i])
	}
	return _ret
}

func (this *QStandardItem) TakeColumn(column int) []*QStandardItem {
	var _ma struct_miqt_array = QStandardItem_TakeColumn(this.h, (int)(column))
	_ret := make([]*QStandardItem, int(_ma.len))
	_outCast := (*[0xffff]*QStandardItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQStandardItem(_outCast[i])
	}
	return _ret
}

func (this *QStandardItem) SortChildren(column int) {
	QStandardItem_SortChildren(this.h, (int)(column))
}

func (this *QStandardItem) Clone() *QStandardItem {
	return newQStandardItem(QStandardItem_Clone(this.h))
}

func (this *QStandardItem) Type() int {
	return (int)(QStandardItem_Type(this.h))
}

func (this *QStandardItem) Read(in *QDataStream) {
	QStandardItem_Read(this.h, in.cPointer())
}

func (this *QStandardItem) Write(out *QDataStream) {
	QStandardItem_Write(this.h, out.cPointer())
}

func (this *QStandardItem) OperatorLesser(other *QStandardItem) bool {
	return (bool)(QStandardItem_OperatorLesser(this.h, other.cPointer()))
}

func (this *QStandardItem) Child2(row int, column int) *QStandardItem {
	return newQStandardItem(QStandardItem_Child2(this.h, (int)(row), (int)(column)))
}

func (this *QStandardItem) TakeChild2(row int, column int) *QStandardItem {
	return newQStandardItem(QStandardItem_TakeChild2(this.h, (int)(row), (int)(column)))
}

func (this *QStandardItem) SortChildren2(column int, order SortOrder) {
	QStandardItem_SortChildren2(this.h, (int)(column), (int)(order))
}

type QStandardItemModel struct {
	h          uintptr
	isSubclass bool
}

// NewQStandardItemModel constructs a new QStandardItemModel object.
func NewQStandardItemModel() *QStandardItemModel {
	ret := newQStandardItemModel(QStandardItemModel_new())
	ret.isSubclass = true
	return ret
}

// NewQStandardItemModel2 constructs a new QStandardItemModel object.
func NewQStandardItemModel2(rows int, columns int) *QStandardItemModel {
	ret := newQStandardItemModel(QStandardItemModel_new2((int)(rows), (int)(columns)))
	ret.isSubclass = true
	return ret
}

// NewQStandardItemModel3 constructs a new QStandardItemModel object.
func NewQStandardItemModel3(parent *QObject) *QStandardItemModel {
	ret := newQStandardItemModel(QStandardItemModel_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStandardItemModel4 constructs a new QStandardItemModel object.
func NewQStandardItemModel4(rows int, columns int, parent *QObject) *QStandardItemModel {
	ret := newQStandardItemModel(QStandardItemModel_new4((int)(rows), (int)(columns), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStandardItemModel) MetaObject() *QMetaObject {
	return newQMetaObject(QStandardItemModel_MetaObject(this.h))
}

func (this *QStandardItemModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStandardItemModel_Metacast(this.h, param1_Cstring))
}

func QStandardItemModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStandardItemModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItemModel) SetItemRoleNames(roleNames map[int][]byte) {
	roleNames_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roleNames))))
	defer free(unsafe.Pointer(roleNames_Keys_CArray))
	roleNames_Values_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(roleNames))))
	defer free(unsafe.Pointer(roleNames_Values_CArray))
	roleNames_ctr := 0
	for roleNames_k, roleNames_v := range roleNames {
		roleNames_Keys_CArray[roleNames_ctr] = (int)(roleNames_k)
		roleNames_v_alias := struct_miqt_string{}
		roleNames_v_alias.data = (char)(unsafe.Pointer(&roleNames_v[0]))
		roleNames_v_alias.len = size_t(len(roleNames_v))
		roleNames_Values_CArray[roleNames_ctr] = roleNames_v_alias
		roleNames_ctr++
	}
	roleNames_mm := struct_miqt_map{
		len:    size_t(len(roleNames)),
		keys:   unsafe.Pointer(roleNames_Keys_CArray),
		values: unsafe.Pointer(roleNames_Values_CArray),
	}
	QStandardItemModel_SetItemRoleNames(this.h, roleNames_mm)
}

func (this *QStandardItemModel) RoleNames() map[int][]byte {
	var _mm struct_miqt_map = QStandardItemModel_RoleNames(this.h)
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

func (this *QStandardItemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QStandardItemModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItemModel) Parent(child *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QStandardItemModel_Parent(this.h, child.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItemModel) RowCount(parent *QModelIndex) int {
	return (int)(QStandardItemModel_RowCount(this.h, parent.cPointer()))
}

func (this *QStandardItemModel) ColumnCount(parent *QModelIndex) int {
	return (int)(QStandardItemModel_ColumnCount(this.h, parent.cPointer()))
}

func (this *QStandardItemModel) HasChildren(parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_HasChildren(this.h, parent.cPointer()))
}

func (this *QStandardItemModel) Data(index *QModelIndex, role int) *QVariant {
	_goptr := newQVariant(QStandardItemModel_Data(this.h, index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItemModel) MultiData(index *QModelIndex, roleDataSpan QModelRoleDataSpan) {
	QStandardItemModel_MultiData(this.h, index.cPointer(), roleDataSpan.cPointer())
}

func (this *QStandardItemModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	return (bool)(QStandardItemModel_SetData(this.h, index.cPointer(), value.cPointer(), (int)(role)))
}

func (this *QStandardItemModel) ClearItemData(index *QModelIndex) bool {
	return (bool)(QStandardItemModel_ClearItemData(this.h, index.cPointer()))
}

func (this *QStandardItemModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	_goptr := newQVariant(QStandardItemModel_HeaderData(this.h, (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItemModel) SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {
	return (bool)(QStandardItemModel_SetHeaderData(this.h, (int)(section), (int)(orientation), value.cPointer(), (int)(role)))
}

func (this *QStandardItemModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_InsertRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QStandardItemModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_InsertColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QStandardItemModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_RemoveRows(this.h, (int)(row), (int)(count), parent.cPointer()))
}

func (this *QStandardItemModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_RemoveColumns(this.h, (int)(column), (int)(count), parent.cPointer()))
}

func (this *QStandardItemModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QStandardItemModel_Flags(this.h, index.cPointer()))
}

func (this *QStandardItemModel) SupportedDropActions() DropAction {
	return (DropAction)(QStandardItemModel_SupportedDropActions(this.h))
}

func (this *QStandardItemModel) ItemData(index *QModelIndex) map[int]QVariant {
	var _mm struct_miqt_map = QStandardItemModel_ItemData(this.h, index.cPointer())
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

func (this *QStandardItemModel) SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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
	return (bool)(QStandardItemModel_SetItemData(this.h, index.cPointer(), roles_mm))
}

func (this *QStandardItemModel) Clear() {
	QStandardItemModel_Clear(this.h)
}

func (this *QStandardItemModel) Sort(column int, order SortOrder) {
	QStandardItemModel_Sort(this.h, (int)(column), (int)(order))
}

func (this *QStandardItemModel) ItemFromIndex(index *QModelIndex) *QStandardItem {
	return newQStandardItem(QStandardItemModel_ItemFromIndex(this.h, index.cPointer()))
}

func (this *QStandardItemModel) IndexFromItem(item *QStandardItem) *QModelIndex {
	_goptr := newQModelIndex(QStandardItemModel_IndexFromItem(this.h, item.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStandardItemModel) Item(row int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_Item(this.h, (int)(row)))
}

func (this *QStandardItemModel) SetItem(row int, column int, item *QStandardItem) {
	QStandardItemModel_SetItem(this.h, (int)(row), (int)(column), item.cPointer())
}

func (this *QStandardItemModel) SetItem2(row int, item *QStandardItem) {
	QStandardItemModel_SetItem2(this.h, (int)(row), item.cPointer())
}

func (this *QStandardItemModel) InvisibleRootItem() *QStandardItem {
	return newQStandardItem(QStandardItemModel_InvisibleRootItem(this.h))
}

func (this *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_HorizontalHeaderItem(this.h, (int)(column)))
}

func (this *QStandardItemModel) SetHorizontalHeaderItem(column int, item *QStandardItem) {
	QStandardItemModel_SetHorizontalHeaderItem(this.h, (int)(column), item.cPointer())
}

func (this *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_VerticalHeaderItem(this.h, (int)(row)))
}

func (this *QStandardItemModel) SetVerticalHeaderItem(row int, item *QStandardItem) {
	QStandardItemModel_SetVerticalHeaderItem(this.h, (int)(row), item.cPointer())
}

func (this *QStandardItemModel) SetHorizontalHeaderLabels(labels []string) {
	labels_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(labels))))
	defer free(unsafe.Pointer(labels_CArray))
	for i := range labels {
		labels_i_ms := struct_miqt_string{}
		labels_i_ms.data = CString(labels[i])
		labels_i_ms.len = size_t(len(labels[i]))
		defer free(unsafe.Pointer(labels_i_ms.data))
		labels_CArray[i] = labels_i_ms
	}
	labels_ma := struct_miqt_array{len: size_t(len(labels)), data: unsafe.Pointer(labels_CArray)}
	QStandardItemModel_SetHorizontalHeaderLabels(this.h, labels_ma)
}

func (this *QStandardItemModel) SetVerticalHeaderLabels(labels []string) {
	labels_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(labels))))
	defer free(unsafe.Pointer(labels_CArray))
	for i := range labels {
		labels_i_ms := struct_miqt_string{}
		labels_i_ms.data = CString(labels[i])
		labels_i_ms.len = size_t(len(labels[i]))
		defer free(unsafe.Pointer(labels_i_ms.data))
		labels_CArray[i] = labels_i_ms
	}
	labels_ma := struct_miqt_array{len: size_t(len(labels)), data: unsafe.Pointer(labels_CArray)}
	QStandardItemModel_SetVerticalHeaderLabels(this.h, labels_ma)
}

func (this *QStandardItemModel) SetRowCount(rows int) {
	QStandardItemModel_SetRowCount(this.h, (int)(rows))
}

func (this *QStandardItemModel) SetColumnCount(columns int) {
	QStandardItemModel_SetColumnCount(this.h, (int)(columns))
}

func (this *QStandardItemModel) AppendRow(items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItemModel_AppendRow(this.h, items_ma)
}

func (this *QStandardItemModel) AppendColumn(items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItemModel_AppendColumn(this.h, items_ma)
}

func (this *QStandardItemModel) AppendRowWithItem(item *QStandardItem) {
	QStandardItemModel_AppendRowWithItem(this.h, item.cPointer())
}

func (this *QStandardItemModel) InsertRow(row int, items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItemModel_InsertRow(this.h, (int)(row), items_ma)
}

func (this *QStandardItemModel) InsertColumn(column int, items []*QStandardItem) {
	items_CArray := (*[0xffff]*QStandardItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QStandardItemModel_InsertColumn(this.h, (int)(column), items_ma)
}

func (this *QStandardItemModel) InsertRow2(row int, item *QStandardItem) {
	QStandardItemModel_InsertRow2(this.h, (int)(row), item.cPointer())
}

func (this *QStandardItemModel) InsertRowWithRow(row int) bool {
	return (bool)(QStandardItemModel_InsertRowWithRow(this.h, (int)(row)))
}

func (this *QStandardItemModel) InsertColumnWithColumn(column int) bool {
	return (bool)(QStandardItemModel_InsertColumnWithColumn(this.h, (int)(column)))
}

func (this *QStandardItemModel) TakeItem(row int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_TakeItem(this.h, (int)(row)))
}

func (this *QStandardItemModel) TakeRow(row int) []*QStandardItem {
	var _ma struct_miqt_array = QStandardItemModel_TakeRow(this.h, (int)(row))
	_ret := make([]*QStandardItem, int(_ma.len))
	_outCast := (*[0xffff]*QStandardItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQStandardItem(_outCast[i])
	}
	return _ret
}

func (this *QStandardItemModel) TakeColumn(column int) []*QStandardItem {
	var _ma struct_miqt_array = QStandardItemModel_TakeColumn(this.h, (int)(column))
	_ret := make([]*QStandardItem, int(_ma.len))
	_outCast := (*[0xffff]*QStandardItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQStandardItem(_outCast[i])
	}
	return _ret
}

func (this *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_TakeHorizontalHeaderItem(this.h, (int)(column)))
}

func (this *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_TakeVerticalHeaderItem(this.h, (int)(row)))
}

func (this *QStandardItemModel) ItemPrototype() *QStandardItem {
	return newQStandardItem(QStandardItemModel_ItemPrototype(this.h))
}

func (this *QStandardItemModel) SetItemPrototype(item *QStandardItem) {
	QStandardItemModel_SetItemPrototype(this.h, item.cPointer())
}

func (this *QStandardItemModel) FindItems(text string) []*QStandardItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QStandardItemModel_FindItems(this.h, text_ms)
	_ret := make([]*QStandardItem, int(_ma.len))
	_outCast := (*[0xffff]*QStandardItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQStandardItem(_outCast[i])
	}
	return _ret
}

func (this *QStandardItemModel) SortRole() int {
	return (int)(QStandardItemModel_SortRole(this.h))
}

func (this *QStandardItemModel) SetSortRole(role int) {
	QStandardItemModel_SetSortRole(this.h, (int)(role))
}

func (this *QStandardItemModel) MimeTypes() []string {
	var _ma struct_miqt_array = QStandardItemModel_MimeTypes(this.h)
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

func (this *QStandardItemModel) MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}
	return newQMimeData(QStandardItemModel_MimeData(this.h, indexes_ma))
}

func (this *QStandardItemModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QStandardItemModel) ItemChanged(item *QStandardItem) {
	QStandardItemModel_ItemChanged(this.h, item.cPointer())
}

func (this *QStandardItemModel) OnItemChanged(slot func(item *QStandardItem)) {
	QStandardItemModel_connect_ItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_ItemChanged
func miqt_exec_callback_QStandardItemModel_ItemChanged(cb intptr_t, item *QStandardItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QStandardItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStandardItem(item)

	gofunc(slotval1)
}

func QStandardItemModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStandardItemModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStandardItemModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStandardItemModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStandardItemModel) Item2(row int, column int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_Item2(this.h, (int)(row), (int)(column)))
}

func (this *QStandardItemModel) InsertRow22(row int, parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_InsertRow22(this.h, (int)(row), parent.cPointer()))
}

func (this *QStandardItemModel) InsertColumn2(column int, parent *QModelIndex) bool {
	return (bool)(QStandardItemModel_InsertColumn2(this.h, (int)(column), parent.cPointer()))
}

func (this *QStandardItemModel) TakeItem2(row int, column int) *QStandardItem {
	return newQStandardItem(QStandardItemModel_TakeItem2(this.h, (int)(row), (int)(column)))
}

func (this *QStandardItemModel) FindItems2(text string, flags MatchFlag) []*QStandardItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QStandardItemModel_FindItems2(this.h, text_ms, (int)(flags))
	_ret := make([]*QStandardItem, int(_ma.len))
	_outCast := (*[0xffff]*QStandardItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQStandardItem(_outCast[i])
	}
	return _ret
}

func (this *QStandardItemModel) FindItems3(text string, flags MatchFlag, column int) []*QStandardItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QStandardItemModel_FindItems3(this.h, text_ms, (int)(flags), (int)(column))
	_ret := make([]*QStandardItem, int(_ma.len))
	_outCast := (*[0xffff]*QStandardItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQStandardItem(_outCast[i])
	}
	return _ret
}

func (this *QStandardItemModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QStandardItemModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QStandardItemModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_MetaObject
func miqt_exec_callback_QStandardItemModel_MetaObject(self QStandardItemModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QStandardItemModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QStandardItemModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QStandardItemModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Metacast
func miqt_exec_callback_QStandardItemModel_Metacast(self QStandardItemModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
