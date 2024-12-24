package qt

import (
	"unsafe"
)

type QTableWidgetItem__ItemType int

const (
	QTableWidgetItem__Type     QTableWidgetItem__ItemType = 0
	QTableWidgetItem__UserType QTableWidgetItem__ItemType = 1000
)

type QTableWidgetSelectionRange struct {
	h          uintptr
	isSubclass bool
}

// NewQTableWidgetSelectionRange constructs a new QTableWidgetSelectionRange object.
func NewQTableWidgetSelectionRange() *QTableWidgetSelectionRange {
	ret := newQTableWidgetSelectionRange(QTableWidgetSelectionRange_new())
	ret.isSubclass = true
	return ret
}

// NewQTableWidgetSelectionRange2 constructs a new QTableWidgetSelectionRange object.
func NewQTableWidgetSelectionRange2(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	ret := newQTableWidgetSelectionRange(QTableWidgetSelectionRange_new2((int)(top), (int)(left), (int)(bottom), (int)(right)))
	ret.isSubclass = true
	return ret
}

func (this *QTableWidgetSelectionRange) TopRow() int {
	return (int)(QTableWidgetSelectionRange_TopRow(this.h))
}

func (this *QTableWidgetSelectionRange) BottomRow() int {
	return (int)(QTableWidgetSelectionRange_BottomRow(this.h))
}

func (this *QTableWidgetSelectionRange) LeftColumn() int {
	return (int)(QTableWidgetSelectionRange_LeftColumn(this.h))
}

func (this *QTableWidgetSelectionRange) RightColumn() int {
	return (int)(QTableWidgetSelectionRange_RightColumn(this.h))
}

func (this *QTableWidgetSelectionRange) RowCount() int {
	return (int)(QTableWidgetSelectionRange_RowCount(this.h))
}

func (this *QTableWidgetSelectionRange) ColumnCount() int {
	return (int)(QTableWidgetSelectionRange_ColumnCount(this.h))
}

type QTableWidgetItem struct {
	h          uintptr
	isSubclass bool
}

// NewQTableWidgetItem constructs a new QTableWidgetItem object.
func NewQTableWidgetItem() *QTableWidgetItem {
	ret := newQTableWidgetItem(QTableWidgetItem_new())
	ret.isSubclass = true
	return ret
}

// NewQTableWidgetItem2 constructs a new QTableWidgetItem object.
func NewQTableWidgetItem2(text string) *QTableWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTableWidgetItem(QTableWidgetItem_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQTableWidgetItem3 constructs a new QTableWidgetItem object.
func NewQTableWidgetItem3(icon *QIcon, text string) *QTableWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTableWidgetItem(QTableWidgetItem_new3(icon.cPointer(), text_ms))
	ret.isSubclass = true
	return ret
}

// NewQTableWidgetItem4 constructs a new QTableWidgetItem object.
func NewQTableWidgetItem4(other *QTableWidgetItem) *QTableWidgetItem {
	ret := newQTableWidgetItem(QTableWidgetItem_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTableWidgetItem5 constructs a new QTableWidgetItem object.
func NewQTableWidgetItem5(typeVal int) *QTableWidgetItem {
	ret := newQTableWidgetItem(QTableWidgetItem_new5((int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTableWidgetItem6 constructs a new QTableWidgetItem object.
func NewQTableWidgetItem6(text string, typeVal int) *QTableWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTableWidgetItem(QTableWidgetItem_new6(text_ms, (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTableWidgetItem7 constructs a new QTableWidgetItem object.
func NewQTableWidgetItem7(icon *QIcon, text string, typeVal int) *QTableWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQTableWidgetItem(QTableWidgetItem_new7(icon.cPointer(), text_ms, (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

func (this *QTableWidgetItem) Clone() *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidgetItem_Clone(this.h))
}

func (this *QTableWidgetItem) TableWidget() *QTableWidget {
	return newQTableWidget(QTableWidgetItem_TableWidget(this.h))
}

func (this *QTableWidgetItem) Row() int {
	return (int)(QTableWidgetItem_Row(this.h))
}

func (this *QTableWidgetItem) Column() int {
	return (int)(QTableWidgetItem_Column(this.h))
}

func (this *QTableWidgetItem) SetSelected(selectVal bool) {
	QTableWidgetItem_SetSelected(this.h, (bool)(selectVal))
}

func (this *QTableWidgetItem) IsSelected() bool {
	return (bool)(QTableWidgetItem_IsSelected(this.h))
}

func (this *QTableWidgetItem) Flags() ItemFlag {
	return (ItemFlag)(QTableWidgetItem_Flags(this.h))
}

func (this *QTableWidgetItem) SetFlags(flags ItemFlag) {
	QTableWidgetItem_SetFlags(this.h, (int)(flags))
}

func (this *QTableWidgetItem) Text() string {
	var _ms struct_miqt_string = QTableWidgetItem_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableWidgetItem) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTableWidgetItem_SetText(this.h, text_ms)
}

func (this *QTableWidgetItem) Icon() *QIcon {
	_goptr := newQIcon(QTableWidgetItem_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidgetItem) SetIcon(icon *QIcon) {
	QTableWidgetItem_SetIcon(this.h, icon.cPointer())
}

func (this *QTableWidgetItem) StatusTip() string {
	var _ms struct_miqt_string = QTableWidgetItem_StatusTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableWidgetItem) SetStatusTip(statusTip string) {
	statusTip_ms := struct_miqt_string{}
	statusTip_ms.data = CString(statusTip)
	statusTip_ms.len = size_t(len(statusTip))
	defer free(unsafe.Pointer(statusTip_ms.data))
	QTableWidgetItem_SetStatusTip(this.h, statusTip_ms)
}

func (this *QTableWidgetItem) ToolTip() string {
	var _ms struct_miqt_string = QTableWidgetItem_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableWidgetItem) SetToolTip(toolTip string) {
	toolTip_ms := struct_miqt_string{}
	toolTip_ms.data = CString(toolTip)
	toolTip_ms.len = size_t(len(toolTip))
	defer free(unsafe.Pointer(toolTip_ms.data))
	QTableWidgetItem_SetToolTip(this.h, toolTip_ms)
}

func (this *QTableWidgetItem) WhatsThis() string {
	var _ms struct_miqt_string = QTableWidgetItem_WhatsThis(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableWidgetItem) SetWhatsThis(whatsThis string) {
	whatsThis_ms := struct_miqt_string{}
	whatsThis_ms.data = CString(whatsThis)
	whatsThis_ms.len = size_t(len(whatsThis))
	defer free(unsafe.Pointer(whatsThis_ms.data))
	QTableWidgetItem_SetWhatsThis(this.h, whatsThis_ms)
}

func (this *QTableWidgetItem) Font() *QFont {
	_goptr := newQFont(QTableWidgetItem_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidgetItem) SetFont(font *QFont) {
	QTableWidgetItem_SetFont(this.h, font.cPointer())
}

func (this *QTableWidgetItem) TextAlignment() int {
	return (int)(QTableWidgetItem_TextAlignment(this.h))
}

func (this *QTableWidgetItem) SetTextAlignment(alignment int) {
	QTableWidgetItem_SetTextAlignment(this.h, (int)(alignment))
}

func (this *QTableWidgetItem) SetTextAlignmentWithAlignment(alignment AlignmentFlag) {
	QTableWidgetItem_SetTextAlignmentWithAlignment(this.h, (int)(alignment))
}

func (this *QTableWidgetItem) SetTextAlignment2(alignment AlignmentFlag) {
	QTableWidgetItem_SetTextAlignment2(this.h, (int)(alignment))
}

func (this *QTableWidgetItem) Background() *QBrush {
	_goptr := newQBrush(QTableWidgetItem_Background(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidgetItem) SetBackground(brush *QBrush) {
	QTableWidgetItem_SetBackground(this.h, brush.cPointer())
}

func (this *QTableWidgetItem) Foreground() *QBrush {
	_goptr := newQBrush(QTableWidgetItem_Foreground(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidgetItem) SetForeground(brush *QBrush) {
	QTableWidgetItem_SetForeground(this.h, brush.cPointer())
}

func (this *QTableWidgetItem) CheckState() CheckState {
	return (CheckState)(QTableWidgetItem_CheckState(this.h))
}

func (this *QTableWidgetItem) SetCheckState(state CheckState) {
	QTableWidgetItem_SetCheckState(this.h, (int)(state))
}

func (this *QTableWidgetItem) SizeHint() *QSize {
	_goptr := newQSize(QTableWidgetItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidgetItem) SetSizeHint(size *QSize) {
	QTableWidgetItem_SetSizeHint(this.h, size.cPointer())
}

func (this *QTableWidgetItem) Data(role int) *QVariant {
	_goptr := newQVariant(QTableWidgetItem_Data(this.h, (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidgetItem) SetData(role int, value *QVariant) {
	QTableWidgetItem_SetData(this.h, (int)(role), value.cPointer())
}

func (this *QTableWidgetItem) OperatorLesser(other *QTableWidgetItem) bool {
	return (bool)(QTableWidgetItem_OperatorLesser(this.h, other.cPointer()))
}

func (this *QTableWidgetItem) Read(in *QDataStream) {
	QTableWidgetItem_Read(this.h, in.cPointer())
}

func (this *QTableWidgetItem) Write(out *QDataStream) {
	QTableWidgetItem_Write(this.h, out.cPointer())
}

func (this *QTableWidgetItem) OperatorAssign(other *QTableWidgetItem) {
	QTableWidgetItem_OperatorAssign(this.h, other.cPointer())
}

func (this *QTableWidgetItem) Type() int {
	return (int)(QTableWidgetItem_Type(this.h))
}

type QTableWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQTableWidget constructs a new QTableWidget object.
func NewQTableWidget(parent *QWidget) *QTableWidget {
	ret := newQTableWidget(QTableWidget_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTableWidget2 constructs a new QTableWidget object.
func NewQTableWidget2() *QTableWidget {
	ret := newQTableWidget(QTableWidget_new2())
	ret.isSubclass = true
	return ret
}

// NewQTableWidget3 constructs a new QTableWidget object.
func NewQTableWidget3(rows int, columns int) *QTableWidget {
	ret := newQTableWidget(QTableWidget_new3((int)(rows), (int)(columns)))
	ret.isSubclass = true
	return ret
}

// NewQTableWidget4 constructs a new QTableWidget object.
func NewQTableWidget4(rows int, columns int, parent *QWidget) *QTableWidget {
	ret := newQTableWidget(QTableWidget_new4((int)(rows), (int)(columns), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTableWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QTableWidget_MetaObject(this.h))
}

func (this *QTableWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTableWidget_Metacast(this.h, param1_Cstring))
}

func QTableWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTableWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableWidget) SetRowCount(rows int) {
	QTableWidget_SetRowCount(this.h, (int)(rows))
}

func (this *QTableWidget) RowCount() int {
	return (int)(QTableWidget_RowCount(this.h))
}

func (this *QTableWidget) SetColumnCount(columns int) {
	QTableWidget_SetColumnCount(this.h, (int)(columns))
}

func (this *QTableWidget) ColumnCount() int {
	return (int)(QTableWidget_ColumnCount(this.h))
}

func (this *QTableWidget) Row(item *QTableWidgetItem) int {
	return (int)(QTableWidget_Row(this.h, item.cPointer()))
}

func (this *QTableWidget) Column(item *QTableWidgetItem) int {
	return (int)(QTableWidget_Column(this.h, item.cPointer()))
}

func (this *QTableWidget) Item(row int, column int) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_Item(this.h, (int)(row), (int)(column)))
}

func (this *QTableWidget) SetItem(row int, column int, item *QTableWidgetItem) {
	QTableWidget_SetItem(this.h, (int)(row), (int)(column), item.cPointer())
}

func (this *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_TakeItem(this.h, (int)(row), (int)(column)))
}

func (this *QTableWidget) Items(data *QMimeData) []*QTableWidgetItem {
	var _ma struct_miqt_array = QTableWidget_Items(this.h, data.cPointer())
	_ret := make([]*QTableWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QTableWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTableWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QTableWidget) IndexFromItem(item *QTableWidgetItem) *QModelIndex {
	_goptr := newQModelIndex(QTableWidget_IndexFromItem(this.h, item.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidget) ItemFromIndex(index *QModelIndex) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_ItemFromIndex(this.h, index.cPointer()))
}

func (this *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_VerticalHeaderItem(this.h, (int)(row)))
}

func (this *QTableWidget) SetVerticalHeaderItem(row int, item *QTableWidgetItem) {
	QTableWidget_SetVerticalHeaderItem(this.h, (int)(row), item.cPointer())
}

func (this *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_TakeVerticalHeaderItem(this.h, (int)(row)))
}

func (this *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_HorizontalHeaderItem(this.h, (int)(column)))
}

func (this *QTableWidget) SetHorizontalHeaderItem(column int, item *QTableWidgetItem) {
	QTableWidget_SetHorizontalHeaderItem(this.h, (int)(column), item.cPointer())
}

func (this *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_TakeHorizontalHeaderItem(this.h, (int)(column)))
}

func (this *QTableWidget) SetVerticalHeaderLabels(labels []string) {
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
	QTableWidget_SetVerticalHeaderLabels(this.h, labels_ma)
}

func (this *QTableWidget) SetHorizontalHeaderLabels(labels []string) {
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
	QTableWidget_SetHorizontalHeaderLabels(this.h, labels_ma)
}

func (this *QTableWidget) CurrentRow() int {
	return (int)(QTableWidget_CurrentRow(this.h))
}

func (this *QTableWidget) CurrentColumn() int {
	return (int)(QTableWidget_CurrentColumn(this.h))
}

func (this *QTableWidget) CurrentItem() *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_CurrentItem(this.h))
}

func (this *QTableWidget) SetCurrentItem(item *QTableWidgetItem) {
	QTableWidget_SetCurrentItem(this.h, item.cPointer())
}

func (this *QTableWidget) SetCurrentItem2(item *QTableWidgetItem, command SelectionFlag) {
	QTableWidget_SetCurrentItem2(this.h, item.cPointer(), (int)(command))
}

func (this *QTableWidget) SetCurrentCell(row int, column int) {
	QTableWidget_SetCurrentCell(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) SetCurrentCell2(row int, column int, command SelectionFlag) {
	QTableWidget_SetCurrentCell2(this.h, (int)(row), (int)(column), (int)(command))
}

func (this *QTableWidget) SortItems(column int) {
	QTableWidget_SortItems(this.h, (int)(column))
}

func (this *QTableWidget) SetSortingEnabled(enable bool) {
	QTableWidget_SetSortingEnabled(this.h, (bool)(enable))
}

func (this *QTableWidget) IsSortingEnabled() bool {
	return (bool)(QTableWidget_IsSortingEnabled(this.h))
}

func (this *QTableWidget) EditItem(item *QTableWidgetItem) {
	QTableWidget_EditItem(this.h, item.cPointer())
}

func (this *QTableWidget) OpenPersistentEditor(item *QTableWidgetItem) {
	QTableWidget_OpenPersistentEditor(this.h, item.cPointer())
}

func (this *QTableWidget) ClosePersistentEditor(item *QTableWidgetItem) {
	QTableWidget_ClosePersistentEditor(this.h, item.cPointer())
}

func (this *QTableWidget) IsPersistentEditorOpen(item *QTableWidgetItem) bool {
	return (bool)(QTableWidget_IsPersistentEditorOpen(this.h, item.cPointer()))
}

func (this *QTableWidget) CellWidget(row int, column int) *QWidget {
	return newQWidget(QTableWidget_CellWidget(this.h, (int)(row), (int)(column)))
}

func (this *QTableWidget) SetCellWidget(row int, column int, widget *QWidget) {
	QTableWidget_SetCellWidget(this.h, (int)(row), (int)(column), widget.cPointer())
}

func (this *QTableWidget) RemoveCellWidget(row int, column int) {
	QTableWidget_RemoveCellWidget(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) SetRangeSelected(rangeVal *QTableWidgetSelectionRange, selectVal bool) {
	QTableWidget_SetRangeSelected(this.h, rangeVal.cPointer(), (bool)(selectVal))
}

func (this *QTableWidget) SelectedRanges() []QTableWidgetSelectionRange {
	var _ma struct_miqt_array = QTableWidget_SelectedRanges(this.h)
	_ret := make([]QTableWidgetSelectionRange, int(_ma.len))
	_outCast := (*[0xffff]*QTableWidgetSelectionRange)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQTableWidgetSelectionRange(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTableWidget) SelectedItems() []*QTableWidgetItem {
	var _ma struct_miqt_array = QTableWidget_SelectedItems(this.h)
	_ret := make([]*QTableWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QTableWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTableWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QTableWidget) FindItems(text string, flags MatchFlag) []*QTableWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QTableWidget_FindItems(this.h, text_ms, (int)(flags))
	_ret := make([]*QTableWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QTableWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTableWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QTableWidget) VisualRow(logicalRow int) int {
	return (int)(QTableWidget_VisualRow(this.h, (int)(logicalRow)))
}

func (this *QTableWidget) VisualColumn(logicalColumn int) int {
	return (int)(QTableWidget_VisualColumn(this.h, (int)(logicalColumn)))
}

func (this *QTableWidget) ItemAt(p *QPoint) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_ItemAt(this.h, p.cPointer()))
}

func (this *QTableWidget) ItemAt2(x int, y int) *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_ItemAt2(this.h, (int)(x), (int)(y)))
}

func (this *QTableWidget) VisualItemRect(item *QTableWidgetItem) *QRect {
	_goptr := newQRect(QTableWidget_VisualItemRect(this.h, item.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTableWidget) ItemPrototype() *QTableWidgetItem {
	return newQTableWidgetItem(QTableWidget_ItemPrototype(this.h))
}

func (this *QTableWidget) SetItemPrototype(item *QTableWidgetItem) {
	QTableWidget_SetItemPrototype(this.h, item.cPointer())
}

func (this *QTableWidget) ScrollToItem(item *QTableWidgetItem) {
	QTableWidget_ScrollToItem(this.h, item.cPointer())
}

func (this *QTableWidget) InsertRow(row int) {
	QTableWidget_InsertRow(this.h, (int)(row))
}

func (this *QTableWidget) InsertColumn(column int) {
	QTableWidget_InsertColumn(this.h, (int)(column))
}

func (this *QTableWidget) RemoveRow(row int) {
	QTableWidget_RemoveRow(this.h, (int)(row))
}

func (this *QTableWidget) RemoveColumn(column int) {
	QTableWidget_RemoveColumn(this.h, (int)(column))
}

func (this *QTableWidget) Clear() {
	QTableWidget_Clear(this.h)
}

func (this *QTableWidget) ClearContents() {
	QTableWidget_ClearContents(this.h)
}

func (this *QTableWidget) ItemPressed(item *QTableWidgetItem) {
	QTableWidget_ItemPressed(this.h, item.cPointer())
}

func (this *QTableWidget) OnItemPressed(slot func(item *QTableWidgetItem)) {
	QTableWidget_connect_ItemPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_ItemPressed
func miqt_exec_callback_QTableWidget_ItemPressed(cb intptr_t, item *QTableWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTableWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTableWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTableWidget) ItemClicked(item *QTableWidgetItem) {
	QTableWidget_ItemClicked(this.h, item.cPointer())
}

func (this *QTableWidget) OnItemClicked(slot func(item *QTableWidgetItem)) {
	QTableWidget_connect_ItemClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_ItemClicked
func miqt_exec_callback_QTableWidget_ItemClicked(cb intptr_t, item *QTableWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTableWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTableWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTableWidget) ItemDoubleClicked(item *QTableWidgetItem) {
	QTableWidget_ItemDoubleClicked(this.h, item.cPointer())
}

func (this *QTableWidget) OnItemDoubleClicked(slot func(item *QTableWidgetItem)) {
	QTableWidget_connect_ItemDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_ItemDoubleClicked
func miqt_exec_callback_QTableWidget_ItemDoubleClicked(cb intptr_t, item *QTableWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTableWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTableWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTableWidget) ItemActivated(item *QTableWidgetItem) {
	QTableWidget_ItemActivated(this.h, item.cPointer())
}

func (this *QTableWidget) OnItemActivated(slot func(item *QTableWidgetItem)) {
	QTableWidget_connect_ItemActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_ItemActivated
func miqt_exec_callback_QTableWidget_ItemActivated(cb intptr_t, item *QTableWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTableWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTableWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTableWidget) ItemEntered(item *QTableWidgetItem) {
	QTableWidget_ItemEntered(this.h, item.cPointer())
}

func (this *QTableWidget) OnItemEntered(slot func(item *QTableWidgetItem)) {
	QTableWidget_connect_ItemEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_ItemEntered
func miqt_exec_callback_QTableWidget_ItemEntered(cb intptr_t, item *QTableWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTableWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTableWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTableWidget) ItemChanged(item *QTableWidgetItem) {
	QTableWidget_ItemChanged(this.h, item.cPointer())
}

func (this *QTableWidget) OnItemChanged(slot func(item *QTableWidgetItem)) {
	QTableWidget_connect_ItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_ItemChanged
func miqt_exec_callback_QTableWidget_ItemChanged(cb intptr_t, item *QTableWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTableWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTableWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTableWidget) CurrentItemChanged(current *QTableWidgetItem, previous *QTableWidgetItem) {
	QTableWidget_CurrentItemChanged(this.h, current.cPointer(), previous.cPointer())
}

func (this *QTableWidget) OnCurrentItemChanged(slot func(current *QTableWidgetItem, previous *QTableWidgetItem)) {
	QTableWidget_connect_CurrentItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CurrentItemChanged
func miqt_exec_callback_QTableWidget_CurrentItemChanged(cb intptr_t, current *QTableWidgetItem, previous *QTableWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QTableWidgetItem, previous *QTableWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTableWidgetItem(current)

	slotval2 := newQTableWidgetItem(previous)

	gofunc(slotval1, slotval2)
}

func (this *QTableWidget) ItemSelectionChanged() {
	QTableWidget_ItemSelectionChanged(this.h)
}

func (this *QTableWidget) OnItemSelectionChanged(slot func()) {
	QTableWidget_connect_ItemSelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_ItemSelectionChanged
func miqt_exec_callback_QTableWidget_ItemSelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTableWidget) CellPressed(row int, column int) {
	QTableWidget_CellPressed(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) OnCellPressed(slot func(row int, column int)) {
	QTableWidget_connect_CellPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CellPressed
func miqt_exec_callback_QTableWidget_CellPressed(cb intptr_t, row int, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(row int, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTableWidget) CellClicked(row int, column int) {
	QTableWidget_CellClicked(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) OnCellClicked(slot func(row int, column int)) {
	QTableWidget_connect_CellClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CellClicked
func miqt_exec_callback_QTableWidget_CellClicked(cb intptr_t, row int, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(row int, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTableWidget) CellDoubleClicked(row int, column int) {
	QTableWidget_CellDoubleClicked(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) OnCellDoubleClicked(slot func(row int, column int)) {
	QTableWidget_connect_CellDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CellDoubleClicked
func miqt_exec_callback_QTableWidget_CellDoubleClicked(cb intptr_t, row int, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(row int, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTableWidget) CellActivated(row int, column int) {
	QTableWidget_CellActivated(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) OnCellActivated(slot func(row int, column int)) {
	QTableWidget_connect_CellActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CellActivated
func miqt_exec_callback_QTableWidget_CellActivated(cb intptr_t, row int, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(row int, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTableWidget) CellEntered(row int, column int) {
	QTableWidget_CellEntered(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) OnCellEntered(slot func(row int, column int)) {
	QTableWidget_connect_CellEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CellEntered
func miqt_exec_callback_QTableWidget_CellEntered(cb intptr_t, row int, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(row int, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTableWidget) CellChanged(row int, column int) {
	QTableWidget_CellChanged(this.h, (int)(row), (int)(column))
}

func (this *QTableWidget) OnCellChanged(slot func(row int, column int)) {
	QTableWidget_connect_CellChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CellChanged
func miqt_exec_callback_QTableWidget_CellChanged(cb intptr_t, row int, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(row int, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTableWidget) CurrentCellChanged(currentRow int, currentColumn int, previousRow int, previousColumn int) {
	QTableWidget_CurrentCellChanged(this.h, (int)(currentRow), (int)(currentColumn), (int)(previousRow), (int)(previousColumn))
}

func (this *QTableWidget) OnCurrentCellChanged(slot func(currentRow int, currentColumn int, previousRow int, previousColumn int)) {
	QTableWidget_connect_CurrentCellChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_CurrentCellChanged
func miqt_exec_callback_QTableWidget_CurrentCellChanged(cb intptr_t, currentRow int, currentColumn int, previousRow int, previousColumn int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(currentRow int, currentColumn int, previousRow int, previousColumn int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(currentRow)

	slotval2 := (int)(currentColumn)

	slotval3 := (int)(previousRow)

	slotval4 := (int)(previousColumn)

	gofunc(slotval1, slotval2, slotval3, slotval4)
}

func QTableWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTableWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTableWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTableWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTableWidget) SortItems2(column int, order SortOrder) {
	QTableWidget_SortItems2(this.h, (int)(column), (int)(order))
}

func (this *QTableWidget) ScrollToItem2(item *QTableWidgetItem, hint QAbstractItemView__ScrollHint) {
	QTableWidget_ScrollToItem2(this.h, item.cPointer(), (int)(hint))
}

func (this *QTableWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTableWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTableWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_MetaObject
func miqt_exec_callback_QTableWidget_MetaObject(self QTableWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTableWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTableWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTableWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTableWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTableWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTableWidget_Metacast
func miqt_exec_callback_QTableWidget_Metacast(self QTableWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTableWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
