package qt

import (
	"unsafe"
)

type QListWidgetItem__ItemType int

const (
	QListWidgetItem__Type     QListWidgetItem__ItemType = 0
	QListWidgetItem__UserType QListWidgetItem__ItemType = 1000
)

type QListWidgetItem struct {
	h          uintptr
	isSubclass bool
}

// NewQListWidgetItem constructs a new QListWidgetItem object.
func NewQListWidgetItem() *QListWidgetItem {
	g := newQListWidgetItem(QListWidgetItem_new())
	g.isSubclass = true
	return g
}

// NewQListWidgetItem2 constructs a new QListWidgetItem object.
func NewQListWidgetItem2(text string) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQListWidgetItem(QListWidgetItem_new2(text_ms))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem3 constructs a new QListWidgetItem object.
func NewQListWidgetItem3(icon *QIcon, text string) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQListWidgetItem(QListWidgetItem_new3(icon.cPointer(), text_ms))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem4 constructs a new QListWidgetItem object.
func NewQListWidgetItem4(other *QListWidgetItem) *QListWidgetItem {
	g := newQListWidgetItem(QListWidgetItem_new4(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem5 constructs a new QListWidgetItem object.
func NewQListWidgetItem5(listview *QListWidget) *QListWidgetItem {
	g := newQListWidgetItem(QListWidgetItem_new5(listview.cPointer()))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem6 constructs a new QListWidgetItem object.
func NewQListWidgetItem6(listview *QListWidget, typeVal int) *QListWidgetItem {
	g := newQListWidgetItem(QListWidgetItem_new6(listview.cPointer(), (int)(typeVal)))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem7 constructs a new QListWidgetItem object.
func NewQListWidgetItem7(text string, listview *QListWidget) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQListWidgetItem(QListWidgetItem_new7(text_ms, listview.cPointer()))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem8 constructs a new QListWidgetItem object.
func NewQListWidgetItem8(text string, listview *QListWidget, typeVal int) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQListWidgetItem(QListWidgetItem_new8(text_ms, listview.cPointer(), (int)(typeVal)))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem9 constructs a new QListWidgetItem object.
func NewQListWidgetItem9(icon *QIcon, text string, listview *QListWidget) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQListWidgetItem(QListWidgetItem_new9(icon.cPointer(), text_ms, listview.cPointer()))
	g.isSubclass = true
	return g
}

// NewQListWidgetItem10 constructs a new QListWidgetItem object.
func NewQListWidgetItem10(icon *QIcon, text string, listview *QListWidget, typeVal int) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQListWidgetItem(QListWidgetItem_new10(icon.cPointer(), text_ms, listview.cPointer(), (int)(typeVal)))
	g.isSubclass = true
	return g
}

func (this *QListWidgetItem) Clone() *QListWidgetItem {
	return newQListWidgetItem(QListWidgetItem_Clone(this.h))
}

func (this *QListWidgetItem) ListWidget() *QListWidget {
	return newQListWidget(QListWidgetItem_ListWidget(this.h))
}

func (this *QListWidgetItem) SetSelected(selectVal bool) {
	QListWidgetItem_SetSelected(this.h, (bool)(selectVal))
}

func (this *QListWidgetItem) IsSelected() bool {
	return (bool)(QListWidgetItem_IsSelected(this.h))
}

func (this *QListWidgetItem) SetHidden(hide bool) {
	QListWidgetItem_SetHidden(this.h, (bool)(hide))
}

func (this *QListWidgetItem) IsHidden() bool {
	return (bool)(QListWidgetItem_IsHidden(this.h))
}

func (this *QListWidgetItem) Flags() ItemFlag {
	return (ItemFlag)(QListWidgetItem_Flags(this.h))
}

func (this *QListWidgetItem) SetFlags(flags ItemFlag) {
	QListWidgetItem_SetFlags(this.h, (int)(flags))
}

func (this *QListWidgetItem) Text() string {
	var _ms struct_miqt_string = QListWidgetItem_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListWidgetItem) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QListWidgetItem_SetText(this.h, text_ms)
}

func (this *QListWidgetItem) Icon() *QIcon {
	_goptr := newQIcon(QListWidgetItem_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidgetItem) SetIcon(icon *QIcon) {
	QListWidgetItem_SetIcon(this.h, icon.cPointer())
}

func (this *QListWidgetItem) StatusTip() string {
	var _ms struct_miqt_string = QListWidgetItem_StatusTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListWidgetItem) SetStatusTip(statusTip string) {
	statusTip_ms := struct_miqt_string{}
	statusTip_ms.data = CString(statusTip)
	statusTip_ms.len = size_t(len(statusTip))
	defer free(unsafe.Pointer(statusTip_ms.data))
	QListWidgetItem_SetStatusTip(this.h, statusTip_ms)
}

func (this *QListWidgetItem) ToolTip() string {
	var _ms struct_miqt_string = QListWidgetItem_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListWidgetItem) SetToolTip(toolTip string) {
	toolTip_ms := struct_miqt_string{}
	toolTip_ms.data = CString(toolTip)
	toolTip_ms.len = size_t(len(toolTip))
	defer free(unsafe.Pointer(toolTip_ms.data))
	QListWidgetItem_SetToolTip(this.h, toolTip_ms)
}

func (this *QListWidgetItem) WhatsThis() string {
	var _ms struct_miqt_string = QListWidgetItem_WhatsThis(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListWidgetItem) SetWhatsThis(whatsThis string) {
	whatsThis_ms := struct_miqt_string{}
	whatsThis_ms.data = CString(whatsThis)
	whatsThis_ms.len = size_t(len(whatsThis))
	defer free(unsafe.Pointer(whatsThis_ms.data))
	QListWidgetItem_SetWhatsThis(this.h, whatsThis_ms)
}

func (this *QListWidgetItem) Font() *QFont {
	_goptr := newQFont(QListWidgetItem_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidgetItem) SetFont(font *QFont) {
	QListWidgetItem_SetFont(this.h, font.cPointer())
}

func (this *QListWidgetItem) TextAlignment() int {
	return (int)(QListWidgetItem_TextAlignment(this.h))
}

func (this *QListWidgetItem) SetTextAlignment(alignment int) {
	QListWidgetItem_SetTextAlignment(this.h, (int)(alignment))
}

func (this *QListWidgetItem) SetTextAlignmentWithAlignment(alignment AlignmentFlag) {
	QListWidgetItem_SetTextAlignmentWithAlignment(this.h, (int)(alignment))
}

func (this *QListWidgetItem) SetTextAlignment2(alignment AlignmentFlag) {
	QListWidgetItem_SetTextAlignment2(this.h, (int)(alignment))
}

func (this *QListWidgetItem) Background() *QBrush {
	_goptr := newQBrush(QListWidgetItem_Background(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidgetItem) SetBackground(brush *QBrush) {
	QListWidgetItem_SetBackground(this.h, brush.cPointer())
}

func (this *QListWidgetItem) Foreground() *QBrush {
	_goptr := newQBrush(QListWidgetItem_Foreground(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidgetItem) SetForeground(brush *QBrush) {
	QListWidgetItem_SetForeground(this.h, brush.cPointer())
}

func (this *QListWidgetItem) CheckState() CheckState {
	return (CheckState)(QListWidgetItem_CheckState(this.h))
}

func (this *QListWidgetItem) SetCheckState(state CheckState) {
	QListWidgetItem_SetCheckState(this.h, (int)(state))
}

func (this *QListWidgetItem) SizeHint() *QSize {
	_goptr := newQSize(QListWidgetItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidgetItem) SetSizeHint(size *QSize) {
	QListWidgetItem_SetSizeHint(this.h, size.cPointer())
}

func (this *QListWidgetItem) Data(role int) *QVariant {
	_goptr := newQVariant(QListWidgetItem_Data(this.h, (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidgetItem) SetData(role int, value *QVariant) {
	QListWidgetItem_SetData(this.h, (int)(role), value.cPointer())
}

func (this *QListWidgetItem) OperatorLesser(other *QListWidgetItem) bool {
	return (bool)(QListWidgetItem_OperatorLesser(this.h, other.cPointer()))
}

func (this *QListWidgetItem) Read(in *QDataStream) {
	QListWidgetItem_Read(this.h, in.cPointer())
}

func (this *QListWidgetItem) Write(out *QDataStream) {
	QListWidgetItem_Write(this.h, out.cPointer())
}

func (this *QListWidgetItem) OperatorAssign(other *QListWidgetItem) {
	QListWidgetItem_OperatorAssign(this.h, other.cPointer())
}

func (this *QListWidgetItem) Type() int {
	return (int)(QListWidgetItem_Type(this.h))
}

type QListWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQListWidget constructs a new QListWidget object.
func NewQListWidget(parent *QWidget) *QListWidget {
	g := newQListWidget(QListWidget_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQListWidget2 constructs a new QListWidget object.
func NewQListWidget2() *QListWidget {
	g := newQListWidget(QListWidget_new2())
	g.isSubclass = true
	return g
}

func (this *QListWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QListWidget_MetaObject(this.h))
}

func (this *QListWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QListWidget_Metacast(this.h, param1_Cstring))
}

func QListWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QListWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListWidget) SetSelectionModel(selectionModel *QItemSelectionModel) {
	QListWidget_SetSelectionModel(this.h, selectionModel.cPointer())
}

func (this *QListWidget) Item(row int) *QListWidgetItem {
	return newQListWidgetItem(QListWidget_Item(this.h, (int)(row)))
}

func (this *QListWidget) Row(item *QListWidgetItem) int {
	return (int)(QListWidget_Row(this.h, item.cPointer()))
}

func (this *QListWidget) InsertItem(row int, item *QListWidgetItem) {
	QListWidget_InsertItem(this.h, (int)(row), item.cPointer())
}

func (this *QListWidget) InsertItem2(row int, label string) {
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	QListWidget_InsertItem2(this.h, (int)(row), label_ms)
}

func (this *QListWidget) InsertItems(row int, labels []string) {
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
	QListWidget_InsertItems(this.h, (int)(row), labels_ma)
}

func (this *QListWidget) AddItem(label string) {
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	QListWidget_AddItem(this.h, label_ms)
}

func (this *QListWidget) AddItemWithItem(item *QListWidgetItem) {
	QListWidget_AddItemWithItem(this.h, item.cPointer())
}

func (this *QListWidget) AddItems(labels []string) {
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
	QListWidget_AddItems(this.h, labels_ma)
}

func (this *QListWidget) TakeItem(row int) *QListWidgetItem {
	return newQListWidgetItem(QListWidget_TakeItem(this.h, (int)(row)))
}

func (this *QListWidget) Count() int {
	return (int)(QListWidget_Count(this.h))
}

func (this *QListWidget) CurrentItem() *QListWidgetItem {
	return newQListWidgetItem(QListWidget_CurrentItem(this.h))
}

func (this *QListWidget) SetCurrentItem(item *QListWidgetItem) {
	QListWidget_SetCurrentItem(this.h, item.cPointer())
}

func (this *QListWidget) SetCurrentItem2(item *QListWidgetItem, command SelectionFlag) {
	QListWidget_SetCurrentItem2(this.h, item.cPointer(), (int)(command))
}

func (this *QListWidget) CurrentRow() int {
	return (int)(QListWidget_CurrentRow(this.h))
}

func (this *QListWidget) SetCurrentRow(row int) {
	QListWidget_SetCurrentRow(this.h, (int)(row))
}

func (this *QListWidget) SetCurrentRow2(row int, command SelectionFlag) {
	QListWidget_SetCurrentRow2(this.h, (int)(row), (int)(command))
}

func (this *QListWidget) ItemAt(p *QPoint) *QListWidgetItem {
	return newQListWidgetItem(QListWidget_ItemAt(this.h, p.cPointer()))
}

func (this *QListWidget) ItemAt2(x int, y int) *QListWidgetItem {
	return newQListWidgetItem(QListWidget_ItemAt2(this.h, (int)(x), (int)(y)))
}

func (this *QListWidget) VisualItemRect(item *QListWidgetItem) *QRect {
	_goptr := newQRect(QListWidget_VisualItemRect(this.h, item.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidget) SortItems() {
	QListWidget_SortItems(this.h)
}

func (this *QListWidget) SetSortingEnabled(enable bool) {
	QListWidget_SetSortingEnabled(this.h, (bool)(enable))
}

func (this *QListWidget) IsSortingEnabled() bool {
	return (bool)(QListWidget_IsSortingEnabled(this.h))
}

func (this *QListWidget) EditItem(item *QListWidgetItem) {
	QListWidget_EditItem(this.h, item.cPointer())
}

func (this *QListWidget) OpenPersistentEditor(item *QListWidgetItem) {
	QListWidget_OpenPersistentEditor(this.h, item.cPointer())
}

func (this *QListWidget) ClosePersistentEditor(item *QListWidgetItem) {
	QListWidget_ClosePersistentEditor(this.h, item.cPointer())
}

func (this *QListWidget) IsPersistentEditorOpen(item *QListWidgetItem) bool {
	return (bool)(QListWidget_IsPersistentEditorOpen(this.h, item.cPointer()))
}

func (this *QListWidget) ItemWidget(item *QListWidgetItem) *QWidget {
	return newQWidget(QListWidget_ItemWidget(this.h, item.cPointer()))
}

func (this *QListWidget) SetItemWidget(item *QListWidgetItem, widget *QWidget) {
	QListWidget_SetItemWidget(this.h, item.cPointer(), widget.cPointer())
}

func (this *QListWidget) RemoveItemWidget(item *QListWidgetItem) {
	QListWidget_RemoveItemWidget(this.h, item.cPointer())
}

func (this *QListWidget) SelectedItems() []*QListWidgetItem {
	var _ma struct_miqt_array = QListWidget_SelectedItems(this.h)
	_ret := make([]*QListWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QListWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQListWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QListWidget) FindItems(text string, flags MatchFlag) []*QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QListWidget_FindItems(this.h, text_ms, (int)(flags))
	_ret := make([]*QListWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QListWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQListWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QListWidget) Items(data *QMimeData) []*QListWidgetItem {
	var _ma struct_miqt_array = QListWidget_Items(this.h, data.cPointer())
	_ret := make([]*QListWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QListWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQListWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QListWidget) IndexFromItem(item *QListWidgetItem) *QModelIndex {
	_goptr := newQModelIndex(QListWidget_IndexFromItem(this.h, item.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListWidget) ItemFromIndex(index *QModelIndex) *QListWidgetItem {
	return newQListWidgetItem(QListWidget_ItemFromIndex(this.h, index.cPointer()))
}

func (this *QListWidget) ScrollToItem(item *QListWidgetItem) {
	QListWidget_ScrollToItem(this.h, item.cPointer())
}

func (this *QListWidget) Clear() {
	QListWidget_Clear(this.h)
}

func (this *QListWidget) ItemPressed(item *QListWidgetItem) {
	QListWidget_ItemPressed(this.h, item.cPointer())
}

func (this *QListWidget) OnItemPressed(slot func(item *QListWidgetItem)) {
	QListWidget_connect_ItemPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ItemPressed
func miqt_exec_callback_QListWidget_ItemPressed(cb intptr_t, item *QListWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QListWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(item)

	gofunc(slotval1)
}

func (this *QListWidget) ItemClicked(item *QListWidgetItem) {
	QListWidget_ItemClicked(this.h, item.cPointer())
}

func (this *QListWidget) OnItemClicked(slot func(item *QListWidgetItem)) {
	QListWidget_connect_ItemClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ItemClicked
func miqt_exec_callback_QListWidget_ItemClicked(cb intptr_t, item *QListWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QListWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(item)

	gofunc(slotval1)
}

func (this *QListWidget) ItemDoubleClicked(item *QListWidgetItem) {
	QListWidget_ItemDoubleClicked(this.h, item.cPointer())
}

func (this *QListWidget) OnItemDoubleClicked(slot func(item *QListWidgetItem)) {
	QListWidget_connect_ItemDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ItemDoubleClicked
func miqt_exec_callback_QListWidget_ItemDoubleClicked(cb intptr_t, item *QListWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QListWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(item)

	gofunc(slotval1)
}

func (this *QListWidget) ItemActivated(item *QListWidgetItem) {
	QListWidget_ItemActivated(this.h, item.cPointer())
}

func (this *QListWidget) OnItemActivated(slot func(item *QListWidgetItem)) {
	QListWidget_connect_ItemActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ItemActivated
func miqt_exec_callback_QListWidget_ItemActivated(cb intptr_t, item *QListWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QListWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(item)

	gofunc(slotval1)
}

func (this *QListWidget) ItemEntered(item *QListWidgetItem) {
	QListWidget_ItemEntered(this.h, item.cPointer())
}

func (this *QListWidget) OnItemEntered(slot func(item *QListWidgetItem)) {
	QListWidget_connect_ItemEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ItemEntered
func miqt_exec_callback_QListWidget_ItemEntered(cb intptr_t, item *QListWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QListWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(item)

	gofunc(slotval1)
}

func (this *QListWidget) ItemChanged(item *QListWidgetItem) {
	QListWidget_ItemChanged(this.h, item.cPointer())
}

func (this *QListWidget) OnItemChanged(slot func(item *QListWidgetItem)) {
	QListWidget_connect_ItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ItemChanged
func miqt_exec_callback_QListWidget_ItemChanged(cb intptr_t, item *QListWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QListWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(item)

	gofunc(slotval1)
}

func (this *QListWidget) CurrentItemChanged(current *QListWidgetItem, previous *QListWidgetItem) {
	QListWidget_CurrentItemChanged(this.h, current.cPointer(), previous.cPointer())
}

func (this *QListWidget) OnCurrentItemChanged(slot func(current *QListWidgetItem, previous *QListWidgetItem)) {
	QListWidget_connect_CurrentItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_CurrentItemChanged
func miqt_exec_callback_QListWidget_CurrentItemChanged(cb intptr_t, current *QListWidgetItem, previous *QListWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QListWidgetItem, previous *QListWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(current)

	slotval2 := newQListWidgetItem(previous)

	gofunc(slotval1, slotval2)
}

func (this *QListWidget) CurrentTextChanged(currentText string) {
	currentText_ms := struct_miqt_string{}
	currentText_ms.data = CString(currentText)
	currentText_ms.len = size_t(len(currentText))
	defer free(unsafe.Pointer(currentText_ms.data))
	QListWidget_CurrentTextChanged(this.h, currentText_ms)
}

func (this *QListWidget) OnCurrentTextChanged(slot func(currentText string)) {
	QListWidget_connect_CurrentTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_CurrentTextChanged
func miqt_exec_callback_QListWidget_CurrentTextChanged(cb intptr_t, currentText struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(currentText string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var currentText_ms struct_miqt_string = currentText
	currentText_ret := GoStringN(currentText_ms.data, int(int64(currentText_ms.len)))
	free(unsafe.Pointer(currentText_ms.data))
	slotval1 := currentText_ret

	gofunc(slotval1)
}

func (this *QListWidget) CurrentRowChanged(currentRow int) {
	QListWidget_CurrentRowChanged(this.h, (int)(currentRow))
}

func (this *QListWidget) OnCurrentRowChanged(slot func(currentRow int)) {
	QListWidget_connect_CurrentRowChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_CurrentRowChanged
func miqt_exec_callback_QListWidget_CurrentRowChanged(cb intptr_t, currentRow int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(currentRow int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(currentRow)

	gofunc(slotval1)
}

func (this *QListWidget) ItemSelectionChanged() {
	QListWidget_ItemSelectionChanged(this.h)
}

func (this *QListWidget) OnItemSelectionChanged(slot func()) {
	QListWidget_connect_ItemSelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ItemSelectionChanged
func miqt_exec_callback_QListWidget_ItemSelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QListWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QListWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QListWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QListWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListWidget) SortItems1(order SortOrder) {
	QListWidget_SortItems1(this.h, (int)(order))
}

func (this *QListWidget) ScrollToItem2(item *QListWidgetItem, hint QAbstractItemView__ScrollHint) {
	QListWidget_ScrollToItem2(this.h, item.cPointer(), (int)(hint))
}

func (this *QListWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QListWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QListWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_MetaObject
func miqt_exec_callback_QListWidget_MetaObject(self QListWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QListWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QListWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QListWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_Metacast
func miqt_exec_callback_QListWidget_Metacast(self QListWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
