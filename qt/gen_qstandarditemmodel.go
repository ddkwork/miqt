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

func (this *QStandardItem) callVirtualBase_Data(role int) *QVariant {

	_goptr := newQVariant(QStandardItem_virtualbase_Data(unsafe.Pointer(this.h), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItem) OnData(slot func(super func(role int) *QVariant, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_Data
func miqt_exec_callback_QStandardItem_Data(self QStandardItem, cb intptr_t, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(role int) *QVariant, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(role)

	virtualReturn := gofunc((&QStandardItem{h: self}).callVirtualBase_Data, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStandardItem) callVirtualBase_MultiData(roleDataSpan QModelRoleDataSpan) {

	QStandardItem_virtualbase_MultiData(unsafe.Pointer(this.h), roleDataSpan.cPointer())

}
func (this *QStandardItem) OnMultiData(slot func(super func(roleDataSpan QModelRoleDataSpan), roleDataSpan QModelRoleDataSpan)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_MultiData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_MultiData
func miqt_exec_callback_QStandardItem_MultiData(self QStandardItem, cb intptr_t, roleDataSpan *QModelRoleDataSpan) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(roleDataSpan QModelRoleDataSpan), roleDataSpan QModelRoleDataSpan))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	roleDataSpan_goptr := newQModelRoleDataSpan(roleDataSpan)
	roleDataSpan_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *roleDataSpan_goptr

	gofunc((&QStandardItem{h: self}).callVirtualBase_MultiData, slotval1)

}

func (this *QStandardItem) callVirtualBase_SetData(value *QVariant, role int) {

	QStandardItem_virtualbase_SetData(unsafe.Pointer(this.h), value.cPointer(), (int)(role))

}
func (this *QStandardItem) OnSetData(slot func(super func(value *QVariant, role int), value *QVariant, role int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_SetData
func miqt_exec_callback_QStandardItem_SetData(self QStandardItem, cb intptr_t, value *QVariant, role int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value *QVariant, role int), value *QVariant, role int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(value)

	slotval2 := (int)(role)

	gofunc((&QStandardItem{h: self}).callVirtualBase_SetData, slotval1, slotval2)

}

func (this *QStandardItem) callVirtualBase_Clone() *QStandardItem {

	return newQStandardItem(QStandardItem_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QStandardItem) OnClone(slot func(super func() *QStandardItem) *QStandardItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_Clone
func miqt_exec_callback_QStandardItem_Clone(self QStandardItem, cb intptr_t) *QStandardItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QStandardItem) *QStandardItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItem{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QStandardItem) callVirtualBase_Type() int {

	return (int)(QStandardItem_virtualbase_Type(unsafe.Pointer(this.h)))

}
func (this *QStandardItem) OnType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_Type(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_Type
func miqt_exec_callback_QStandardItem_Type(self QStandardItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItem{h: self}).callVirtualBase_Type)

	return (int)(virtualReturn)

}

func (this *QStandardItem) callVirtualBase_Read(in *QDataStream) {

	QStandardItem_virtualbase_Read(unsafe.Pointer(this.h), in.cPointer())

}
func (this *QStandardItem) OnRead(slot func(super func(in *QDataStream), in *QDataStream)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_Read(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_Read
func miqt_exec_callback_QStandardItem_Read(self QStandardItem, cb intptr_t, in *QDataStream) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(in *QDataStream), in *QDataStream))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(in)

	gofunc((&QStandardItem{h: self}).callVirtualBase_Read, slotval1)

}

func (this *QStandardItem) callVirtualBase_Write(out *QDataStream) {

	QStandardItem_virtualbase_Write(unsafe.Pointer(this.h), out.cPointer())

}
func (this *QStandardItem) OnWrite(slot func(super func(out *QDataStream), out *QDataStream)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_Write(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_Write
func miqt_exec_callback_QStandardItem_Write(self QStandardItem, cb intptr_t, out *QDataStream) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(out *QDataStream), out *QDataStream))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(out)

	gofunc((&QStandardItem{h: self}).callVirtualBase_Write, slotval1)

}

func (this *QStandardItem) callVirtualBase_OperatorLesser(other *QStandardItem) bool {

	return (bool)(QStandardItem_virtualbase_OperatorLesser(unsafe.Pointer(this.h), other.cPointer()))

}
func (this *QStandardItem) OnOperatorLesser(slot func(super func(other *QStandardItem) bool, other *QStandardItem) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItem_override_virtual_OperatorLesser(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItem_OperatorLesser
func miqt_exec_callback_QStandardItem_OperatorLesser(self QStandardItem, cb intptr_t, other *QStandardItem) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(other *QStandardItem) bool, other *QStandardItem) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStandardItem(other)

	virtualReturn := gofunc((&QStandardItem{h: self}).callVirtualBase_OperatorLesser, slotval1)

	return (bool)(virtualReturn)

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

func (this *QStandardItemModel) callVirtualBase_RoleNames() map[int][]byte {

	var _mm struct_miqt_map = QStandardItemModel_virtualbase_RoleNames(unsafe.Pointer(this.h))
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
func (this *QStandardItemModel) OnRoleNames(slot func(super func() map[int][]byte) map[int][]byte) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_RoleNames(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_RoleNames
func miqt_exec_callback_QStandardItemModel_RoleNames(self QStandardItemModel, cb intptr_t) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() map[int][]byte) map[int][]byte)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_RoleNames)
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

func (this *QStandardItemModel) callVirtualBase_Index(row int, column int, parent *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QStandardItemModel_virtualbase_Index(unsafe.Pointer(this.h), (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItemModel) OnIndex(slot func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Index(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Index
func miqt_exec_callback_QStandardItemModel_Index(self QStandardItemModel, cb intptr_t, row int, column int, parent *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Index, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_Parent(child *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QStandardItemModel_virtualbase_Parent(unsafe.Pointer(this.h), child.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItemModel) OnParent(slot func(super func(child *QModelIndex) *QModelIndex, child *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Parent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Parent
func miqt_exec_callback_QStandardItemModel_Parent(self QStandardItemModel, cb intptr_t, child *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(child *QModelIndex) *QModelIndex, child *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(child)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Parent, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_RowCount(parent *QModelIndex) int {

	return (int)(QStandardItemModel_virtualbase_RowCount(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QStandardItemModel) OnRowCount(slot func(super func(parent *QModelIndex) int, parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_RowCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_RowCount
func miqt_exec_callback_QStandardItemModel_RowCount(self QStandardItemModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) int, parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_RowCount, slotval1)

	return (int)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_ColumnCount(parent *QModelIndex) int {

	return (int)(QStandardItemModel_virtualbase_ColumnCount(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QStandardItemModel) OnColumnCount(slot func(super func(parent *QModelIndex) int, parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_ColumnCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_ColumnCount
func miqt_exec_callback_QStandardItemModel_ColumnCount(self QStandardItemModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) int, parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_ColumnCount, slotval1)

	return (int)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_HasChildren(parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_HasChildren(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QStandardItemModel) OnHasChildren(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_HasChildren(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_HasChildren
func miqt_exec_callback_QStandardItemModel_HasChildren(self QStandardItemModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_HasChildren, slotval1)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_Data(index *QModelIndex, role int) *QVariant {

	_goptr := newQVariant(QStandardItemModel_virtualbase_Data(unsafe.Pointer(this.h), index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItemModel) OnData(slot func(super func(index *QModelIndex, role int) *QVariant, index *QModelIndex, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Data
func miqt_exec_callback_QStandardItemModel_Data(self QStandardItemModel, cb intptr_t, index *QModelIndex, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, role int) *QVariant, index *QModelIndex, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (int)(role)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Data, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_MultiData(index *QModelIndex, roleDataSpan QModelRoleDataSpan) {

	QStandardItemModel_virtualbase_MultiData(unsafe.Pointer(this.h), index.cPointer(), roleDataSpan.cPointer())

}
func (this *QStandardItemModel) OnMultiData(slot func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_MultiData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_MultiData
func miqt_exec_callback_QStandardItemModel_MultiData(self QStandardItemModel, cb intptr_t, index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	roleDataSpan_goptr := newQModelRoleDataSpan(roleDataSpan)
	roleDataSpan_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval2 := *roleDataSpan_goptr

	gofunc((&QStandardItemModel{h: self}).callVirtualBase_MultiData, slotval1, slotval2)

}

func (this *QStandardItemModel) callVirtualBase_SetData(index *QModelIndex, value *QVariant, role int) bool {

	return (bool)(QStandardItemModel_virtualbase_SetData(unsafe.Pointer(this.h), index.cPointer(), value.cPointer(), (int)(role)))

}
func (this *QStandardItemModel) OnSetData(slot func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_SetData
func miqt_exec_callback_QStandardItemModel_SetData(self QStandardItemModel, cb intptr_t, index *QModelIndex, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQVariant(value)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_SetData, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_ClearItemData(index *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_ClearItemData(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QStandardItemModel) OnClearItemData(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_ClearItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_ClearItemData
func miqt_exec_callback_QStandardItemModel_ClearItemData(self QStandardItemModel, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_ClearItemData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_HeaderData(section int, orientation Orientation, role int) *QVariant {

	_goptr := newQVariant(QStandardItemModel_virtualbase_HeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItemModel) OnHeaderData(slot func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_HeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_HeaderData
func miqt_exec_callback_QStandardItemModel_HeaderData(self QStandardItemModel, cb intptr_t, section int, orientation int, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_HeaderData, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {

	return (bool)(QStandardItemModel_virtualbase_SetHeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), value.cPointer(), (int)(role)))

}
func (this *QStandardItemModel) OnSetHeaderData(slot func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_SetHeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_SetHeaderData
func miqt_exec_callback_QStandardItemModel_SetHeaderData(self QStandardItemModel, cb intptr_t, section int, orientation int, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := newQVariant(value)

	slotval4 := (int)(role)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_SetHeaderData, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_InsertRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_InsertRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QStandardItemModel) OnInsertRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_InsertRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_InsertRows
func miqt_exec_callback_QStandardItemModel_InsertRows(self QStandardItemModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_InsertRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_InsertColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_InsertColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QStandardItemModel) OnInsertColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_InsertColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_InsertColumns
func miqt_exec_callback_QStandardItemModel_InsertColumns(self QStandardItemModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_InsertColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_RemoveRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_RemoveRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QStandardItemModel) OnRemoveRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_RemoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_RemoveRows
func miqt_exec_callback_QStandardItemModel_RemoveRows(self QStandardItemModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_RemoveRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_RemoveColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_RemoveColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QStandardItemModel) OnRemoveColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_RemoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_RemoveColumns
func miqt_exec_callback_QStandardItemModel_RemoveColumns(self QStandardItemModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_RemoveColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_Flags(index *QModelIndex) ItemFlag {

	return (ItemFlag)(QStandardItemModel_virtualbase_Flags(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QStandardItemModel) OnFlags(slot func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Flags(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Flags
func miqt_exec_callback_QStandardItemModel_Flags(self QStandardItemModel, cb intptr_t, index *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Flags, slotval1)

	return (int)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QStandardItemModel_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QStandardItemModel) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_SupportedDropActions
func miqt_exec_callback_QStandardItemModel_SupportedDropActions(self QStandardItemModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_ItemData(index *QModelIndex) map[int]QVariant {

	var _mm struct_miqt_map = QStandardItemModel_virtualbase_ItemData(unsafe.Pointer(this.h), index.cPointer())
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
func (this *QStandardItemModel) OnItemData(slot func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_ItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_ItemData
func miqt_exec_callback_QStandardItemModel_ItemData(self QStandardItemModel, cb intptr_t, index *QModelIndex) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_ItemData, slotval1)
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

func (this *QStandardItemModel) callVirtualBase_SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
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

	return (bool)(QStandardItemModel_virtualbase_SetItemData(unsafe.Pointer(this.h), index.cPointer(), roles_mm))

}
func (this *QStandardItemModel) OnSetItemData(slot func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_SetItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_SetItemData
func miqt_exec_callback_QStandardItemModel_SetItemData(self QStandardItemModel, cb intptr_t, index *QModelIndex, roles struct_miqt_map) bool {
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

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_SetItemData, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_Sort(column int, order SortOrder) {

	QStandardItemModel_virtualbase_Sort(unsafe.Pointer(this.h), (int)(column), (int)(order))

}
func (this *QStandardItemModel) OnSort(slot func(super func(column int, order SortOrder), column int, order SortOrder)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Sort(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Sort
func miqt_exec_callback_QStandardItemModel_Sort(self QStandardItemModel, cb intptr_t, column int, order int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, order SortOrder), column int, order SortOrder))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (SortOrder)(order)

	gofunc((&QStandardItemModel{h: self}).callVirtualBase_Sort, slotval1, slotval2)

}

func (this *QStandardItemModel) callVirtualBase_MimeTypes() []string {

	var _ma struct_miqt_array = QStandardItemModel_virtualbase_MimeTypes(unsafe.Pointer(this.h))
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
func (this *QStandardItemModel) OnMimeTypes(slot func(super func() []string) []string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_MimeTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_MimeTypes
func miqt_exec_callback_QStandardItemModel_MimeTypes(self QStandardItemModel, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_MimeTypes)
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

func (this *QStandardItemModel) callVirtualBase_MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}

	return newQMimeData(QStandardItemModel_virtualbase_MimeData(unsafe.Pointer(this.h), indexes_ma))

}
func (this *QStandardItemModel) OnMimeData(slot func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_MimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_MimeData
func miqt_exec_callback_QStandardItemModel_MimeData(self QStandardItemModel, cb intptr_t, indexes struct_miqt_array) *QMimeData {
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

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_MimeData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_DropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QStandardItemModel) OnDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_DropMimeData
func miqt_exec_callback_QStandardItemModel_DropMimeData(self QStandardItemModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_Sibling(row int, column int, idx *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QStandardItemModel_virtualbase_Sibling(unsafe.Pointer(this.h), (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItemModel) OnSibling(slot func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Sibling(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Sibling
func miqt_exec_callback_QStandardItemModel_Sibling(self QStandardItemModel, cb intptr_t, row int, column int, idx *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(idx)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Sibling, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_CanDropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QStandardItemModel) OnCanDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_CanDropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_CanDropMimeData
func miqt_exec_callback_QStandardItemModel_CanDropMimeData(self QStandardItemModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
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

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_CanDropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_SupportedDragActions() DropAction {

	return (DropAction)(QStandardItemModel_virtualbase_SupportedDragActions(unsafe.Pointer(this.h)))

}
func (this *QStandardItemModel) OnSupportedDragActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_SupportedDragActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_SupportedDragActions
func miqt_exec_callback_QStandardItemModel_SupportedDragActions(self QStandardItemModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_SupportedDragActions)

	return (int)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QStandardItemModel_virtualbase_MoveRows(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QStandardItemModel) OnMoveRows(slot func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_MoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_MoveRows
func miqt_exec_callback_QStandardItemModel_MoveRows(self QStandardItemModel, cb intptr_t, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
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

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_MoveRows, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QStandardItemModel_virtualbase_MoveColumns(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QStandardItemModel) OnMoveColumns(slot func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_MoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_MoveColumns
func miqt_exec_callback_QStandardItemModel_MoveColumns(self QStandardItemModel, cb intptr_t, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
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

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_MoveColumns, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_FetchMore(parent *QModelIndex) {

	QStandardItemModel_virtualbase_FetchMore(unsafe.Pointer(this.h), parent.cPointer())

}
func (this *QStandardItemModel) OnFetchMore(slot func(super func(parent *QModelIndex), parent *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_FetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_FetchMore
func miqt_exec_callback_QStandardItemModel_FetchMore(self QStandardItemModel, cb intptr_t, parent *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex), parent *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	gofunc((&QStandardItemModel{h: self}).callVirtualBase_FetchMore, slotval1)

}

func (this *QStandardItemModel) callVirtualBase_CanFetchMore(parent *QModelIndex) bool {

	return (bool)(QStandardItemModel_virtualbase_CanFetchMore(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QStandardItemModel) OnCanFetchMore(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_CanFetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_CanFetchMore
func miqt_exec_callback_QStandardItemModel_CanFetchMore(self QStandardItemModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_CanFetchMore, slotval1)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_Buddy(index *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QStandardItemModel_virtualbase_Buddy(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItemModel) OnBuddy(slot func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Buddy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Buddy
func miqt_exec_callback_QStandardItemModel_Buddy(self QStandardItemModel, cb intptr_t, index *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Buddy, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {

	var _ma struct_miqt_array = QStandardItemModel_virtualbase_Match(unsafe.Pointer(this.h), start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QStandardItemModel) OnMatch(slot func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Match(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Match
func miqt_exec_callback_QStandardItemModel_Match(self QStandardItemModel, cb intptr_t, start *QModelIndex, role int, value *QVariant, hits int, flags int) struct_miqt_array {
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

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Match, slotval1, slotval2, slotval3, slotval4, slotval5)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QStandardItemModel) callVirtualBase_Span(index *QModelIndex) *QSize {

	_goptr := newQSize(QStandardItemModel_virtualbase_Span(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStandardItemModel) OnSpan(slot func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Span(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Span
func miqt_exec_callback_QStandardItemModel_Span(self QStandardItemModel, cb intptr_t, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Span, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStandardItemModel) callVirtualBase_Submit() bool {

	return (bool)(QStandardItemModel_virtualbase_Submit(unsafe.Pointer(this.h)))

}
func (this *QStandardItemModel) OnSubmit(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Submit(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Submit
func miqt_exec_callback_QStandardItemModel_Submit(self QStandardItemModel, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStandardItemModel{h: self}).callVirtualBase_Submit)

	return (bool)(virtualReturn)

}

func (this *QStandardItemModel) callVirtualBase_Revert() {

	QStandardItemModel_virtualbase_Revert(unsafe.Pointer(this.h))

}
func (this *QStandardItemModel) OnRevert(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_Revert(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_Revert
func miqt_exec_callback_QStandardItemModel_Revert(self QStandardItemModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QStandardItemModel{h: self}).callVirtualBase_Revert)

}

func (this *QStandardItemModel) callVirtualBase_ResetInternalData() {

	QStandardItemModel_virtualbase_ResetInternalData(unsafe.Pointer(this.h))

}
func (this *QStandardItemModel) OnResetInternalData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStandardItemModel_override_virtual_ResetInternalData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStandardItemModel_ResetInternalData
func miqt_exec_callback_QStandardItemModel_ResetInternalData(self QStandardItemModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QStandardItemModel{h: self}).callVirtualBase_ResetInternalData)

}
