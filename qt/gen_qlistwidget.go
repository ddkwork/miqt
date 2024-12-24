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

	ret := newQListWidgetItem(QListWidgetItem_new())
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem2 constructs a new QListWidgetItem object.
func NewQListWidgetItem2(text string) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQListWidgetItem(QListWidgetItem_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem3 constructs a new QListWidgetItem object.
func NewQListWidgetItem3(icon *QIcon, text string) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQListWidgetItem(QListWidgetItem_new3(icon.cPointer(), text_ms))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem4 constructs a new QListWidgetItem object.
func NewQListWidgetItem4(other *QListWidgetItem) *QListWidgetItem {

	ret := newQListWidgetItem(QListWidgetItem_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem5 constructs a new QListWidgetItem object.
func NewQListWidgetItem5(listview *QListWidget) *QListWidgetItem {

	ret := newQListWidgetItem(QListWidgetItem_new5(listview.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem6 constructs a new QListWidgetItem object.
func NewQListWidgetItem6(listview *QListWidget, typeVal int) *QListWidgetItem {

	ret := newQListWidgetItem(QListWidgetItem_new6(listview.cPointer(), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem7 constructs a new QListWidgetItem object.
func NewQListWidgetItem7(text string, listview *QListWidget) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQListWidgetItem(QListWidgetItem_new7(text_ms, listview.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem8 constructs a new QListWidgetItem object.
func NewQListWidgetItem8(text string, listview *QListWidget, typeVal int) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQListWidgetItem(QListWidgetItem_new8(text_ms, listview.cPointer(), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem9 constructs a new QListWidgetItem object.
func NewQListWidgetItem9(icon *QIcon, text string, listview *QListWidget) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQListWidgetItem(QListWidgetItem_new9(icon.cPointer(), text_ms, listview.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQListWidgetItem10 constructs a new QListWidgetItem object.
func NewQListWidgetItem10(icon *QIcon, text string, listview *QListWidget, typeVal int) *QListWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQListWidgetItem(QListWidgetItem_new10(icon.cPointer(), text_ms, listview.cPointer(), (int)(typeVal)))
	ret.isSubclass = true
	return ret
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

func (this *QListWidgetItem) callVirtualBase_Clone() *QListWidgetItem {

	return newQListWidgetItem(QListWidgetItem_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QListWidgetItem) OnClone(slot func(super func() *QListWidgetItem) *QListWidgetItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidgetItem_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidgetItem_Clone
func miqt_exec_callback_QListWidgetItem_Clone(self QListWidgetItem, cb intptr_t) *QListWidgetItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QListWidgetItem) *QListWidgetItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidgetItem{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QListWidgetItem) callVirtualBase_Data(role int) *QVariant {

	_goptr := newQVariant(QListWidgetItem_virtualbase_Data(unsafe.Pointer(this.h), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QListWidgetItem) OnData(slot func(super func(role int) *QVariant, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidgetItem_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidgetItem_Data
func miqt_exec_callback_QListWidgetItem_Data(self QListWidgetItem, cb intptr_t, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(role int) *QVariant, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(role)

	virtualReturn := gofunc((&QListWidgetItem{h: self}).callVirtualBase_Data, slotval1)

	return virtualReturn.cPointer()

}

func (this *QListWidgetItem) callVirtualBase_SetData(role int, value *QVariant) {

	QListWidgetItem_virtualbase_SetData(unsafe.Pointer(this.h), (int)(role), value.cPointer())

}
func (this *QListWidgetItem) OnSetData(slot func(super func(role int, value *QVariant), role int, value *QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidgetItem_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidgetItem_SetData
func miqt_exec_callback_QListWidgetItem_SetData(self QListWidgetItem, cb intptr_t, role int, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(role int, value *QVariant), role int, value *QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(role)

	slotval2 := newQVariant(value)

	gofunc((&QListWidgetItem{h: self}).callVirtualBase_SetData, slotval1, slotval2)

}

func (this *QListWidgetItem) callVirtualBase_OperatorLesser(other *QListWidgetItem) bool {

	return (bool)(QListWidgetItem_virtualbase_OperatorLesser(unsafe.Pointer(this.h), other.cPointer()))

}
func (this *QListWidgetItem) OnOperatorLesser(slot func(super func(other *QListWidgetItem) bool, other *QListWidgetItem) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidgetItem_override_virtual_OperatorLesser(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidgetItem_OperatorLesser
func miqt_exec_callback_QListWidgetItem_OperatorLesser(self QListWidgetItem, cb intptr_t, other *QListWidgetItem) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(other *QListWidgetItem) bool, other *QListWidgetItem) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQListWidgetItem(other)

	virtualReturn := gofunc((&QListWidgetItem{h: self}).callVirtualBase_OperatorLesser, slotval1)

	return (bool)(virtualReturn)

}

func (this *QListWidgetItem) callVirtualBase_Read(in *QDataStream) {

	QListWidgetItem_virtualbase_Read(unsafe.Pointer(this.h), in.cPointer())

}
func (this *QListWidgetItem) OnRead(slot func(super func(in *QDataStream), in *QDataStream)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidgetItem_override_virtual_Read(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidgetItem_Read
func miqt_exec_callback_QListWidgetItem_Read(self QListWidgetItem, cb intptr_t, in *QDataStream) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(in *QDataStream), in *QDataStream))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(in)

	gofunc((&QListWidgetItem{h: self}).callVirtualBase_Read, slotval1)

}

func (this *QListWidgetItem) callVirtualBase_Write(out *QDataStream) {

	QListWidgetItem_virtualbase_Write(unsafe.Pointer(this.h), out.cPointer())

}
func (this *QListWidgetItem) OnWrite(slot func(super func(out *QDataStream), out *QDataStream)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidgetItem_override_virtual_Write(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidgetItem_Write
func miqt_exec_callback_QListWidgetItem_Write(self QListWidgetItem, cb intptr_t, out *QDataStream) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(out *QDataStream), out *QDataStream))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(out)

	gofunc((&QListWidgetItem{h: self}).callVirtualBase_Write, slotval1)

}

type QListWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQListWidget constructs a new QListWidget object.
func NewQListWidget(parent *QWidget) *QListWidget {

	ret := newQListWidget(QListWidget_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQListWidget2 constructs a new QListWidget object.
func NewQListWidget2() *QListWidget {

	ret := newQListWidget(QListWidget_new2())
	ret.isSubclass = true
	return ret
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

func (this *QListWidget) callVirtualBase_SetSelectionModel(selectionModel *QItemSelectionModel) {

	QListWidget_virtualbase_SetSelectionModel(unsafe.Pointer(this.h), selectionModel.cPointer())

}
func (this *QListWidget) OnSetSelectionModel(slot func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_SetSelectionModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_SetSelectionModel
func miqt_exec_callback_QListWidget_SetSelectionModel(self QListWidget, cb intptr_t, selectionModel *QItemSelectionModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelectionModel(selectionModel)

	gofunc((&QListWidget{h: self}).callVirtualBase_SetSelectionModel, slotval1)

}

func (this *QListWidget) callVirtualBase_DropEvent(event *QDropEvent) {

	QListWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QListWidget) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_DropEvent
func miqt_exec_callback_QListWidget_DropEvent(self QListWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QListWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QListWidget_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QListWidget) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_Event
func miqt_exec_callback_QListWidget_Event(self QListWidget, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QListWidget) callVirtualBase_MimeTypes() []string {

	var _ma struct_miqt_array = QListWidget_virtualbase_MimeTypes(unsafe.Pointer(this.h))
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
func (this *QListWidget) OnMimeTypes(slot func(super func() []string) []string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_MimeTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_MimeTypes
func miqt_exec_callback_QListWidget_MimeTypes(self QListWidget, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_MimeTypes)
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

func (this *QListWidget) callVirtualBase_MimeData(items []*QListWidgetItem) *QMimeData {
	items_CArray := (*[0xffff]*QListWidgetItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}

	return newQMimeData(QListWidget_virtualbase_MimeData(unsafe.Pointer(this.h), items_ma))

}
func (this *QListWidget) OnMimeData(slot func(super func(items []*QListWidgetItem) *QMimeData, items []*QListWidgetItem) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_MimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_MimeData
func miqt_exec_callback_QListWidget_MimeData(self QListWidget, cb intptr_t, items struct_miqt_array) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(items []*QListWidgetItem) *QMimeData, items []*QListWidgetItem) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var items_ma struct_miqt_array = items
	items_ret := make([]*QListWidgetItem, int(items_ma.len))
	items_outCast := (*[0xffff]*QListWidgetItem)(unsafe.Pointer(items_ma.data)) // hey ya
	for i := 0; i < int(items_ma.len); i++ {
		items_ret[i] = newQListWidgetItem(items_outCast[i])
	}
	slotval1 := items_ret

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_MimeData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QListWidget) callVirtualBase_DropMimeData(index int, data *QMimeData, action DropAction) bool {

	return (bool)(QListWidget_virtualbase_DropMimeData(unsafe.Pointer(this.h), (int)(index), data.cPointer(), (int)(action)))

}
func (this *QListWidget) OnDropMimeData(slot func(super func(index int, data *QMimeData, action DropAction) bool, index int, data *QMimeData, action DropAction) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_DropMimeData
func miqt_exec_callback_QListWidget_DropMimeData(self QListWidget, cb intptr_t, index int, data *QMimeData, action int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int, data *QMimeData, action DropAction) bool, index int, data *QMimeData, action DropAction) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	slotval2 := newQMimeData(data)

	slotval3 := (DropAction)(action)

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QListWidget) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QListWidget_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QListWidget) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_SupportedDropActions
func miqt_exec_callback_QListWidget_SupportedDropActions(self QListWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QListWidget) callVirtualBase_VisualRect(index *QModelIndex) *QRect {

	_goptr := newQRect(QListWidget_virtualbase_VisualRect(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QListWidget) OnVisualRect(slot func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_VisualRect
func miqt_exec_callback_QListWidget_VisualRect(self QListWidget, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_VisualRect, slotval1)

	return virtualReturn.cPointer()

}

func (this *QListWidget) callVirtualBase_ScrollTo(index *QModelIndex, hint ScrollHint) {

	QListWidget_virtualbase_ScrollTo(unsafe.Pointer(this.h), index.cPointer(), hint)

}
func (this *QListWidget) OnScrollTo(slot func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ScrollTo
func miqt_exec_callback_QListWidget_ScrollTo(self QListWidget, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc((&QListWidget{h: self}).callVirtualBase_ScrollTo, slotval1, slotval2)

}

func (this *QListWidget) callVirtualBase_IndexAt(p *QPoint) *QModelIndex {

	_goptr := newQModelIndex(QListWidget_virtualbase_IndexAt(unsafe.Pointer(this.h), p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QListWidget) OnIndexAt(slot func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_IndexAt
func miqt_exec_callback_QListWidget_IndexAt(self QListWidget, cb intptr_t, p *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(p)

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_IndexAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QListWidget) callVirtualBase_DoItemsLayout() {

	QListWidget_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QListWidget) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_DoItemsLayout
func miqt_exec_callback_QListWidget_DoItemsLayout(self QListWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QListWidget{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QListWidget) callVirtualBase_Reset() {

	QListWidget_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QListWidget) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_Reset
func miqt_exec_callback_QListWidget_Reset(self QListWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QListWidget{h: self}).callVirtualBase_Reset)

}

func (this *QListWidget) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QListWidget_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QListWidget) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_SetRootIndex
func miqt_exec_callback_QListWidget_SetRootIndex(self QListWidget, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QListWidget{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QListWidget) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QListWidget_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QListWidget) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ScrollContentsBy
func miqt_exec_callback_QListWidget_ScrollContentsBy(self QListWidget, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QListWidget{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QListWidget) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QListWidget_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QListWidget) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_DataChanged
func miqt_exec_callback_QListWidget_DataChanged(self QListWidget, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int))
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

	gofunc((&QListWidget{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QListWidget) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QListWidget_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QListWidget) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_RowsInserted
func miqt_exec_callback_QListWidget_RowsInserted(self QListWidget, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QListWidget{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QListWidget) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QListWidget_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QListWidget) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_RowsAboutToBeRemoved
func miqt_exec_callback_QListWidget_RowsAboutToBeRemoved(self QListWidget, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QListWidget{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QListWidget) callVirtualBase_MouseMoveEvent(e *QMouseEvent) {

	QListWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnMouseMoveEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_MouseMoveEvent
func miqt_exec_callback_QListWidget_MouseMoveEvent(self QListWidget, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QListWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_MouseReleaseEvent
func miqt_exec_callback_QListWidget_MouseReleaseEvent(self QListWidget, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QListWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_WheelEvent
func miqt_exec_callback_QListWidget_WheelEvent(self QListWidget, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_TimerEvent(e *QTimerEvent) {

	QListWidget_virtualbase_TimerEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnTimerEvent(slot func(super func(e *QTimerEvent), e *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_TimerEvent
func miqt_exec_callback_QListWidget_TimerEvent(self QListWidget, cb intptr_t, e *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QTimerEvent), e *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_ResizeEvent(e *QResizeEvent) {

	QListWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnResizeEvent(slot func(super func(e *QResizeEvent), e *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ResizeEvent
func miqt_exec_callback_QListWidget_ResizeEvent(self QListWidget, cb intptr_t, e *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QResizeEvent), e *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_DragMoveEvent(e *QDragMoveEvent) {

	QListWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnDragMoveEvent(slot func(super func(e *QDragMoveEvent), e *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_DragMoveEvent
func miqt_exec_callback_QListWidget_DragMoveEvent(self QListWidget, cb intptr_t, e *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragMoveEvent), e *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_DragLeaveEvent(e *QDragLeaveEvent) {

	QListWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnDragLeaveEvent(slot func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_DragLeaveEvent
func miqt_exec_callback_QListWidget_DragLeaveEvent(self QListWidget, cb intptr_t, e *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_StartDrag(supportedActions DropAction) {

	QListWidget_virtualbase_StartDrag(unsafe.Pointer(this.h), (int)(supportedActions))

}
func (this *QListWidget) OnStartDrag(slot func(super func(supportedActions DropAction), supportedActions DropAction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_StartDrag(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_StartDrag
func miqt_exec_callback_QListWidget_StartDrag(self QListWidget, cb intptr_t, supportedActions int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(supportedActions DropAction), supportedActions DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(supportedActions)

	gofunc((&QListWidget{h: self}).callVirtualBase_StartDrag, slotval1)

}

func (this *QListWidget) callVirtualBase_InitViewItemOption(option *QStyleOptionViewItem) {

	QListWidget_virtualbase_InitViewItemOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QListWidget) OnInitViewItemOption(slot func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_InitViewItemOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_InitViewItemOption
func miqt_exec_callback_QListWidget_InitViewItemOption(self QListWidget, cb intptr_t, option *QStyleOptionViewItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	gofunc((&QListWidget{h: self}).callVirtualBase_InitViewItemOption, slotval1)

}

func (this *QListWidget) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QListWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QListWidget) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_PaintEvent
func miqt_exec_callback_QListWidget_PaintEvent(self QListWidget, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QListWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QListWidget) callVirtualBase_HorizontalOffset() int {

	return (int)(QListWidget_virtualbase_HorizontalOffset(unsafe.Pointer(this.h)))

}
func (this *QListWidget) OnHorizontalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_HorizontalOffset
func miqt_exec_callback_QListWidget_HorizontalOffset(self QListWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_HorizontalOffset)

	return (int)(virtualReturn)

}

func (this *QListWidget) callVirtualBase_VerticalOffset() int {

	return (int)(QListWidget_virtualbase_VerticalOffset(unsafe.Pointer(this.h)))

}
func (this *QListWidget) OnVerticalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_VerticalOffset
func miqt_exec_callback_QListWidget_VerticalOffset(self QListWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_VerticalOffset)

	return (int)(virtualReturn)

}

func (this *QListWidget) callVirtualBase_MoveCursor(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex {

	_goptr := newQModelIndex(QListWidget_virtualbase_MoveCursor(unsafe.Pointer(this.h), cursorAction, (int)(modifiers)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QListWidget) OnMoveCursor(slot func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_MoveCursor
func miqt_exec_callback_QListWidget_MoveCursor(self QListWidget, cb intptr_t, cursorAction CursorAction, modifiers int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(modifiers)

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_MoveCursor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QListWidget) callVirtualBase_SetSelection(rect *QRect, command SelectionFlag) {

	QListWidget_virtualbase_SetSelection(unsafe.Pointer(this.h), rect.cPointer(), (int)(command))

}
func (this *QListWidget) OnSetSelection(slot func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_SetSelection
func miqt_exec_callback_QListWidget_SetSelection(self QListWidget, cb intptr_t, rect *QRect, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QListWidget{h: self}).callVirtualBase_SetSelection, slotval1, slotval2)

}

func (this *QListWidget) callVirtualBase_VisualRegionForSelection(selection *QItemSelection) *QRegion {

	_goptr := newQRegion(QListWidget_virtualbase_VisualRegionForSelection(unsafe.Pointer(this.h), selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QListWidget) OnVisualRegionForSelection(slot func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_VisualRegionForSelection
func miqt_exec_callback_QListWidget_VisualRegionForSelection(self QListWidget, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_VisualRegionForSelection, slotval1)

	return virtualReturn.cPointer()

}

func (this *QListWidget) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QListWidget_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QListWidget) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_SelectedIndexes
func miqt_exec_callback_QListWidget_SelectedIndexes(self QListWidget, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QListWidget) callVirtualBase_UpdateGeometries() {

	QListWidget_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QListWidget) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_UpdateGeometries
func miqt_exec_callback_QListWidget_UpdateGeometries(self QListWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QListWidget{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QListWidget) callVirtualBase_IsIndexHidden(index *QModelIndex) bool {

	return (bool)(QListWidget_virtualbase_IsIndexHidden(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QListWidget) OnIsIndexHidden(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_IsIndexHidden
func miqt_exec_callback_QListWidget_IsIndexHidden(self QListWidget, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_IsIndexHidden, slotval1)

	return (bool)(virtualReturn)

}

func (this *QListWidget) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QListWidget_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QListWidget) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_SelectionChanged
func miqt_exec_callback_QListWidget_SelectionChanged(self QListWidget, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QListWidget{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QListWidget) callVirtualBase_CurrentChanged(current *QModelIndex, previous *QModelIndex) {

	QListWidget_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), previous.cPointer())

}
func (this *QListWidget) OnCurrentChanged(slot func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_CurrentChanged
func miqt_exec_callback_QListWidget_CurrentChanged(self QListWidget, cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc((&QListWidget{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}

func (this *QListWidget) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QListWidget_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QListWidget) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListWidget_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListWidget_ViewportSizeHint
func miqt_exec_callback_QListWidget_ViewportSizeHint(self QListWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListWidget{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}
