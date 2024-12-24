package qt

import (
	"unsafe"
)

type QTreeWidgetItem__ItemType int

const (
	QTreeWidgetItem__Type     QTreeWidgetItem__ItemType = 0
	QTreeWidgetItem__UserType QTreeWidgetItem__ItemType = 1000
)

type QTreeWidgetItem__ChildIndicatorPolicy int

const (
	QTreeWidgetItem__ShowIndicator                  QTreeWidgetItem__ChildIndicatorPolicy = 0
	QTreeWidgetItem__DontShowIndicator              QTreeWidgetItem__ChildIndicatorPolicy = 1
	QTreeWidgetItem__DontShowIndicatorWhenChildless QTreeWidgetItem__ChildIndicatorPolicy = 2
)

type QTreeWidgetItem struct {
	h          uintptr
	isSubclass bool
}

// NewQTreeWidgetItem constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem() *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new())
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem2 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem2(strings []string) *QTreeWidgetItem {
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

	ret := newQTreeWidgetItem(QTreeWidgetItem_new2(strings_ma))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem3 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem3(treeview *QTreeWidget) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new3(treeview.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem4 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem4(treeview *QTreeWidget, strings []string) *QTreeWidgetItem {
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

	ret := newQTreeWidgetItem(QTreeWidgetItem_new4(treeview.cPointer(), strings_ma))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem5 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem5(treeview *QTreeWidget, after *QTreeWidgetItem) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new5(treeview.cPointer(), after.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem6 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem6(parent *QTreeWidgetItem) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new6(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem7 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem7(parent *QTreeWidgetItem, strings []string) *QTreeWidgetItem {
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

	ret := newQTreeWidgetItem(QTreeWidgetItem_new7(parent.cPointer(), strings_ma))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem8 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem8(parent *QTreeWidgetItem, after *QTreeWidgetItem) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new8(parent.cPointer(), after.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem9 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem9(other *QTreeWidgetItem) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new9(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem10 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem10(typeVal int) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new10((int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem11 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem11(strings []string, typeVal int) *QTreeWidgetItem {
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

	ret := newQTreeWidgetItem(QTreeWidgetItem_new11(strings_ma, (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem12 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem12(treeview *QTreeWidget, typeVal int) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new12(treeview.cPointer(), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem13 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem13(treeview *QTreeWidget, strings []string, typeVal int) *QTreeWidgetItem {
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

	ret := newQTreeWidgetItem(QTreeWidgetItem_new13(treeview.cPointer(), strings_ma, (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem14 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem14(treeview *QTreeWidget, after *QTreeWidgetItem, typeVal int) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new14(treeview.cPointer(), after.cPointer(), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem15 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem15(parent *QTreeWidgetItem, typeVal int) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new15(parent.cPointer(), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem16 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem16(parent *QTreeWidgetItem, strings []string, typeVal int) *QTreeWidgetItem {
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

	ret := newQTreeWidgetItem(QTreeWidgetItem_new16(parent.cPointer(), strings_ma, (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItem17 constructs a new QTreeWidgetItem object.
func NewQTreeWidgetItem17(parent *QTreeWidgetItem, after *QTreeWidgetItem, typeVal int) *QTreeWidgetItem {
	ret := newQTreeWidgetItem(QTreeWidgetItem_new17(parent.cPointer(), after.cPointer(), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

func (this *QTreeWidgetItem) Clone() *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidgetItem_Clone(this.h))
}

func (this *QTreeWidgetItem) TreeWidget() *QTreeWidget {
	return newQTreeWidget(QTreeWidgetItem_TreeWidget(this.h))
}

func (this *QTreeWidgetItem) SetSelected(selectVal bool) {
	QTreeWidgetItem_SetSelected(this.h, (bool)(selectVal))
}

func (this *QTreeWidgetItem) IsSelected() bool {
	return (bool)(QTreeWidgetItem_IsSelected(this.h))
}

func (this *QTreeWidgetItem) SetHidden(hide bool) {
	QTreeWidgetItem_SetHidden(this.h, (bool)(hide))
}

func (this *QTreeWidgetItem) IsHidden() bool {
	return (bool)(QTreeWidgetItem_IsHidden(this.h))
}

func (this *QTreeWidgetItem) SetExpanded(expand bool) {
	QTreeWidgetItem_SetExpanded(this.h, (bool)(expand))
}

func (this *QTreeWidgetItem) IsExpanded() bool {
	return (bool)(QTreeWidgetItem_IsExpanded(this.h))
}

func (this *QTreeWidgetItem) SetFirstColumnSpanned(span bool) {
	QTreeWidgetItem_SetFirstColumnSpanned(this.h, (bool)(span))
}

func (this *QTreeWidgetItem) IsFirstColumnSpanned() bool {
	return (bool)(QTreeWidgetItem_IsFirstColumnSpanned(this.h))
}

func (this *QTreeWidgetItem) SetDisabled(disabled bool) {
	QTreeWidgetItem_SetDisabled(this.h, (bool)(disabled))
}

func (this *QTreeWidgetItem) IsDisabled() bool {
	return (bool)(QTreeWidgetItem_IsDisabled(this.h))
}

func (this *QTreeWidgetItem) SetChildIndicatorPolicy(policy QTreeWidgetItem__ChildIndicatorPolicy) {
	QTreeWidgetItem_SetChildIndicatorPolicy(this.h, (int)(policy))
}

func (this *QTreeWidgetItem) ChildIndicatorPolicy() QTreeWidgetItem__ChildIndicatorPolicy {
	return (QTreeWidgetItem__ChildIndicatorPolicy)(QTreeWidgetItem_ChildIndicatorPolicy(this.h))
}

func (this *QTreeWidgetItem) Flags() ItemFlag {
	return (ItemFlag)(QTreeWidgetItem_Flags(this.h))
}

func (this *QTreeWidgetItem) SetFlags(flags ItemFlag) {
	QTreeWidgetItem_SetFlags(this.h, (int)(flags))
}

func (this *QTreeWidgetItem) Text(column int) string {
	var _ms struct_miqt_string = QTreeWidgetItem_Text(this.h, (int)(column))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeWidgetItem) SetText(column int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTreeWidgetItem_SetText(this.h, (int)(column), text_ms)
}

func (this *QTreeWidgetItem) Icon(column int) *QIcon {
	_goptr := newQIcon(QTreeWidgetItem_Icon(this.h, (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItem) SetIcon(column int, icon *QIcon) {
	QTreeWidgetItem_SetIcon(this.h, (int)(column), icon.cPointer())
}

func (this *QTreeWidgetItem) StatusTip(column int) string {
	var _ms struct_miqt_string = QTreeWidgetItem_StatusTip(this.h, (int)(column))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeWidgetItem) SetStatusTip(column int, statusTip string) {
	statusTip_ms := struct_miqt_string{}
	statusTip_ms.data = CString(statusTip)
	statusTip_ms.len = size_t(len(statusTip))
	defer free(unsafe.Pointer(statusTip_ms.data))
	QTreeWidgetItem_SetStatusTip(this.h, (int)(column), statusTip_ms)
}

func (this *QTreeWidgetItem) ToolTip(column int) string {
	var _ms struct_miqt_string = QTreeWidgetItem_ToolTip(this.h, (int)(column))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeWidgetItem) SetToolTip(column int, toolTip string) {
	toolTip_ms := struct_miqt_string{}
	toolTip_ms.data = CString(toolTip)
	toolTip_ms.len = size_t(len(toolTip))
	defer free(unsafe.Pointer(toolTip_ms.data))
	QTreeWidgetItem_SetToolTip(this.h, (int)(column), toolTip_ms)
}

func (this *QTreeWidgetItem) WhatsThis(column int) string {
	var _ms struct_miqt_string = QTreeWidgetItem_WhatsThis(this.h, (int)(column))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeWidgetItem) SetWhatsThis(column int, whatsThis string) {
	whatsThis_ms := struct_miqt_string{}
	whatsThis_ms.data = CString(whatsThis)
	whatsThis_ms.len = size_t(len(whatsThis))
	defer free(unsafe.Pointer(whatsThis_ms.data))
	QTreeWidgetItem_SetWhatsThis(this.h, (int)(column), whatsThis_ms)
}

func (this *QTreeWidgetItem) Font(column int) *QFont {
	_goptr := newQFont(QTreeWidgetItem_Font(this.h, (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItem) SetFont(column int, font *QFont) {
	QTreeWidgetItem_SetFont(this.h, (int)(column), font.cPointer())
}

func (this *QTreeWidgetItem) TextAlignment(column int) int {
	return (int)(QTreeWidgetItem_TextAlignment(this.h, (int)(column)))
}

func (this *QTreeWidgetItem) SetTextAlignment(column int, alignment int) {
	QTreeWidgetItem_SetTextAlignment(this.h, (int)(column), (int)(alignment))
}

func (this *QTreeWidgetItem) SetTextAlignment2(column int, alignment AlignmentFlag) {
	QTreeWidgetItem_SetTextAlignment2(this.h, (int)(column), (int)(alignment))
}

func (this *QTreeWidgetItem) SetTextAlignment3(column int, alignment AlignmentFlag) {
	QTreeWidgetItem_SetTextAlignment3(this.h, (int)(column), (int)(alignment))
}

func (this *QTreeWidgetItem) Background(column int) *QBrush {
	_goptr := newQBrush(QTreeWidgetItem_Background(this.h, (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItem) SetBackground(column int, brush *QBrush) {
	QTreeWidgetItem_SetBackground(this.h, (int)(column), brush.cPointer())
}

func (this *QTreeWidgetItem) Foreground(column int) *QBrush {
	_goptr := newQBrush(QTreeWidgetItem_Foreground(this.h, (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItem) SetForeground(column int, brush *QBrush) {
	QTreeWidgetItem_SetForeground(this.h, (int)(column), brush.cPointer())
}

func (this *QTreeWidgetItem) CheckState(column int) CheckState {
	return (CheckState)(QTreeWidgetItem_CheckState(this.h, (int)(column)))
}

func (this *QTreeWidgetItem) SetCheckState(column int, state CheckState) {
	QTreeWidgetItem_SetCheckState(this.h, (int)(column), (int)(state))
}

func (this *QTreeWidgetItem) SizeHint(column int) *QSize {
	_goptr := newQSize(QTreeWidgetItem_SizeHint(this.h, (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItem) SetSizeHint(column int, size *QSize) {
	QTreeWidgetItem_SetSizeHint(this.h, (int)(column), size.cPointer())
}

func (this *QTreeWidgetItem) Data(column int, role int) *QVariant {
	_goptr := newQVariant(QTreeWidgetItem_Data(this.h, (int)(column), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItem) SetData(column int, role int, value *QVariant) {
	QTreeWidgetItem_SetData(this.h, (int)(column), (int)(role), value.cPointer())
}

func (this *QTreeWidgetItem) OperatorLesser(other *QTreeWidgetItem) bool {
	return (bool)(QTreeWidgetItem_OperatorLesser(this.h, other.cPointer()))
}

func (this *QTreeWidgetItem) Read(in *QDataStream) {
	QTreeWidgetItem_Read(this.h, in.cPointer())
}

func (this *QTreeWidgetItem) Write(out *QDataStream) {
	QTreeWidgetItem_Write(this.h, out.cPointer())
}

func (this *QTreeWidgetItem) OperatorAssign(other *QTreeWidgetItem) {
	QTreeWidgetItem_OperatorAssign(this.h, other.cPointer())
}

func (this *QTreeWidgetItem) Parent() *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidgetItem_Parent(this.h))
}

func (this *QTreeWidgetItem) Child(index int) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidgetItem_Child(this.h, (int)(index)))
}

func (this *QTreeWidgetItem) ChildCount() int {
	return (int)(QTreeWidgetItem_ChildCount(this.h))
}

func (this *QTreeWidgetItem) ColumnCount() int {
	return (int)(QTreeWidgetItem_ColumnCount(this.h))
}

func (this *QTreeWidgetItem) IndexOfChild(child *QTreeWidgetItem) int {
	return (int)(QTreeWidgetItem_IndexOfChild(this.h, child.cPointer()))
}

func (this *QTreeWidgetItem) AddChild(child *QTreeWidgetItem) {
	QTreeWidgetItem_AddChild(this.h, child.cPointer())
}

func (this *QTreeWidgetItem) InsertChild(index int, child *QTreeWidgetItem) {
	QTreeWidgetItem_InsertChild(this.h, (int)(index), child.cPointer())
}

func (this *QTreeWidgetItem) RemoveChild(child *QTreeWidgetItem) {
	QTreeWidgetItem_RemoveChild(this.h, child.cPointer())
}

func (this *QTreeWidgetItem) TakeChild(index int) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidgetItem_TakeChild(this.h, (int)(index)))
}

func (this *QTreeWidgetItem) AddChildren(children []*QTreeWidgetItem) {
	children_CArray := (*[0xffff]*QTreeWidgetItem)(malloc(size_t(8 * len(children))))
	defer free(unsafe.Pointer(children_CArray))
	for i := range children {
		children_CArray[i] = children[i].cPointer()
	}
	children_ma := struct_miqt_array{len: size_t(len(children)), data: unsafe.Pointer(children_CArray)}
	QTreeWidgetItem_AddChildren(this.h, children_ma)
}

func (this *QTreeWidgetItem) InsertChildren(index int, children []*QTreeWidgetItem) {
	children_CArray := (*[0xffff]*QTreeWidgetItem)(malloc(size_t(8 * len(children))))
	defer free(unsafe.Pointer(children_CArray))
	for i := range children {
		children_CArray[i] = children[i].cPointer()
	}
	children_ma := struct_miqt_array{len: size_t(len(children)), data: unsafe.Pointer(children_CArray)}
	QTreeWidgetItem_InsertChildren(this.h, (int)(index), children_ma)
}

func (this *QTreeWidgetItem) TakeChildren() []*QTreeWidgetItem {
	var _ma struct_miqt_array = QTreeWidgetItem_TakeChildren(this.h)
	_ret := make([]*QTreeWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QTreeWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTreeWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QTreeWidgetItem) Type() int {
	return (int)(QTreeWidgetItem_Type(this.h))
}

func (this *QTreeWidgetItem) SortChildren(column int, order SortOrder) {
	QTreeWidgetItem_SortChildren(this.h, (int)(column), (int)(order))
}

type QTreeWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQTreeWidget constructs a new QTreeWidget object.
func NewQTreeWidget(parent *QWidget) *QTreeWidget {
	ret := newQTreeWidget(QTreeWidget_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidget2 constructs a new QTreeWidget object.
func NewQTreeWidget2() *QTreeWidget {
	ret := newQTreeWidget(QTreeWidget_new2())
	ret.isSubclass = true
	return ret
}

func (this *QTreeWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QTreeWidget_MetaObject(this.h))
}

func (this *QTreeWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTreeWidget_Metacast(this.h, param1_Cstring))
}

func QTreeWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTreeWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeWidget) ColumnCount() int {
	return (int)(QTreeWidget_ColumnCount(this.h))
}

func (this *QTreeWidget) SetColumnCount(columns int) {
	QTreeWidget_SetColumnCount(this.h, (int)(columns))
}

func (this *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_InvisibleRootItem(this.h))
}

func (this *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_TopLevelItem(this.h, (int)(index)))
}

func (this *QTreeWidget) TopLevelItemCount() int {
	return (int)(QTreeWidget_TopLevelItemCount(this.h))
}

func (this *QTreeWidget) InsertTopLevelItem(index int, item *QTreeWidgetItem) {
	QTreeWidget_InsertTopLevelItem(this.h, (int)(index), item.cPointer())
}

func (this *QTreeWidget) AddTopLevelItem(item *QTreeWidgetItem) {
	QTreeWidget_AddTopLevelItem(this.h, item.cPointer())
}

func (this *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_TakeTopLevelItem(this.h, (int)(index)))
}

func (this *QTreeWidget) IndexOfTopLevelItem(item *QTreeWidgetItem) int {
	return (int)(QTreeWidget_IndexOfTopLevelItem(this.h, item.cPointer()))
}

func (this *QTreeWidget) InsertTopLevelItems(index int, items []*QTreeWidgetItem) {
	items_CArray := (*[0xffff]*QTreeWidgetItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QTreeWidget_InsertTopLevelItems(this.h, (int)(index), items_ma)
}

func (this *QTreeWidget) AddTopLevelItems(items []*QTreeWidgetItem) {
	items_CArray := (*[0xffff]*QTreeWidgetItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QTreeWidget_AddTopLevelItems(this.h, items_ma)
}

func (this *QTreeWidget) HeaderItem() *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_HeaderItem(this.h))
}

func (this *QTreeWidget) SetHeaderItem(item *QTreeWidgetItem) {
	QTreeWidget_SetHeaderItem(this.h, item.cPointer())
}

func (this *QTreeWidget) SetHeaderLabels(labels []string) {
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
	QTreeWidget_SetHeaderLabels(this.h, labels_ma)
}

func (this *QTreeWidget) SetHeaderLabel(label string) {
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	QTreeWidget_SetHeaderLabel(this.h, label_ms)
}

func (this *QTreeWidget) CurrentItem() *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_CurrentItem(this.h))
}

func (this *QTreeWidget) CurrentColumn() int {
	return (int)(QTreeWidget_CurrentColumn(this.h))
}

func (this *QTreeWidget) SetCurrentItem(item *QTreeWidgetItem) {
	QTreeWidget_SetCurrentItem(this.h, item.cPointer())
}

func (this *QTreeWidget) SetCurrentItem2(item *QTreeWidgetItem, column int) {
	QTreeWidget_SetCurrentItem2(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) SetCurrentItem3(item *QTreeWidgetItem, column int, command SelectionFlag) {
	QTreeWidget_SetCurrentItem3(this.h, item.cPointer(), (int)(column), (int)(command))
}

func (this *QTreeWidget) ItemAt(p *QPoint) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_ItemAt(this.h, p.cPointer()))
}

func (this *QTreeWidget) ItemAt2(x int, y int) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_ItemAt2(this.h, (int)(x), (int)(y)))
}

func (this *QTreeWidget) VisualItemRect(item *QTreeWidgetItem) *QRect {
	_goptr := newQRect(QTreeWidget_VisualItemRect(this.h, item.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidget) SortColumn() int {
	return (int)(QTreeWidget_SortColumn(this.h))
}

func (this *QTreeWidget) SortItems(column int, order SortOrder) {
	QTreeWidget_SortItems(this.h, (int)(column), (int)(order))
}

func (this *QTreeWidget) EditItem(item *QTreeWidgetItem) {
	QTreeWidget_EditItem(this.h, item.cPointer())
}

func (this *QTreeWidget) OpenPersistentEditor(item *QTreeWidgetItem) {
	QTreeWidget_OpenPersistentEditor(this.h, item.cPointer())
}

func (this *QTreeWidget) ClosePersistentEditor(item *QTreeWidgetItem) {
	QTreeWidget_ClosePersistentEditor(this.h, item.cPointer())
}

func (this *QTreeWidget) IsPersistentEditorOpen(item *QTreeWidgetItem) bool {
	return (bool)(QTreeWidget_IsPersistentEditorOpen(this.h, item.cPointer()))
}

func (this *QTreeWidget) ItemWidget(item *QTreeWidgetItem, column int) *QWidget {
	return newQWidget(QTreeWidget_ItemWidget(this.h, item.cPointer(), (int)(column)))
}

func (this *QTreeWidget) SetItemWidget(item *QTreeWidgetItem, column int, widget *QWidget) {
	QTreeWidget_SetItemWidget(this.h, item.cPointer(), (int)(column), widget.cPointer())
}

func (this *QTreeWidget) RemoveItemWidget(item *QTreeWidgetItem, column int) {
	QTreeWidget_RemoveItemWidget(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) SelectedItems() []*QTreeWidgetItem {
	var _ma struct_miqt_array = QTreeWidget_SelectedItems(this.h)
	_ret := make([]*QTreeWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QTreeWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTreeWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QTreeWidget) FindItems(text string, flags MatchFlag) []*QTreeWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QTreeWidget_FindItems(this.h, text_ms, (int)(flags))
	_ret := make([]*QTreeWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QTreeWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTreeWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QTreeWidget) ItemAbove(item *QTreeWidgetItem) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_ItemAbove(this.h, item.cPointer()))
}

func (this *QTreeWidget) ItemBelow(item *QTreeWidgetItem) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_ItemBelow(this.h, item.cPointer()))
}

func (this *QTreeWidget) IndexFromItem(item *QTreeWidgetItem) *QModelIndex {
	_goptr := newQModelIndex(QTreeWidget_IndexFromItem(this.h, item.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidget) ItemFromIndex(index *QModelIndex) *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidget_ItemFromIndex(this.h, index.cPointer()))
}

func (this *QTreeWidget) SetSelectionModel(selectionModel *QItemSelectionModel) {
	QTreeWidget_SetSelectionModel(this.h, selectionModel.cPointer())
}

func (this *QTreeWidget) ScrollToItem(item *QTreeWidgetItem) {
	QTreeWidget_ScrollToItem(this.h, item.cPointer())
}

func (this *QTreeWidget) ExpandItem(item *QTreeWidgetItem) {
	QTreeWidget_ExpandItem(this.h, item.cPointer())
}

func (this *QTreeWidget) CollapseItem(item *QTreeWidgetItem) {
	QTreeWidget_CollapseItem(this.h, item.cPointer())
}

func (this *QTreeWidget) Clear() {
	QTreeWidget_Clear(this.h)
}

func (this *QTreeWidget) ItemPressed(item *QTreeWidgetItem, column int) {
	QTreeWidget_ItemPressed(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) OnItemPressed(slot func(item *QTreeWidgetItem, column int)) {
	QTreeWidget_connect_ItemPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemPressed
func miqt_exec_callback_QTreeWidget_ItemPressed(cb intptr_t, item *QTreeWidgetItem, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTreeWidget) ItemClicked(item *QTreeWidgetItem, column int) {
	QTreeWidget_ItemClicked(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) OnItemClicked(slot func(item *QTreeWidgetItem, column int)) {
	QTreeWidget_connect_ItemClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemClicked
func miqt_exec_callback_QTreeWidget_ItemClicked(cb intptr_t, item *QTreeWidgetItem, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTreeWidget) ItemDoubleClicked(item *QTreeWidgetItem, column int) {
	QTreeWidget_ItemDoubleClicked(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) OnItemDoubleClicked(slot func(item *QTreeWidgetItem, column int)) {
	QTreeWidget_connect_ItemDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemDoubleClicked
func miqt_exec_callback_QTreeWidget_ItemDoubleClicked(cb intptr_t, item *QTreeWidgetItem, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTreeWidget) ItemActivated(item *QTreeWidgetItem, column int) {
	QTreeWidget_ItemActivated(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) OnItemActivated(slot func(item *QTreeWidgetItem, column int)) {
	QTreeWidget_connect_ItemActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemActivated
func miqt_exec_callback_QTreeWidget_ItemActivated(cb intptr_t, item *QTreeWidgetItem, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTreeWidget) ItemEntered(item *QTreeWidgetItem, column int) {
	QTreeWidget_ItemEntered(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) OnItemEntered(slot func(item *QTreeWidgetItem, column int)) {
	QTreeWidget_connect_ItemEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemEntered
func miqt_exec_callback_QTreeWidget_ItemEntered(cb intptr_t, item *QTreeWidgetItem, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTreeWidget) ItemChanged(item *QTreeWidgetItem, column int) {
	QTreeWidget_ItemChanged(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) OnItemChanged(slot func(item *QTreeWidgetItem, column int)) {
	QTreeWidget_connect_ItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemChanged
func miqt_exec_callback_QTreeWidget_ItemChanged(cb intptr_t, item *QTreeWidgetItem, column int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem, column int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	slotval2 := (int)(column)

	gofunc(slotval1, slotval2)
}

func (this *QTreeWidget) ItemExpanded(item *QTreeWidgetItem) {
	QTreeWidget_ItemExpanded(this.h, item.cPointer())
}

func (this *QTreeWidget) OnItemExpanded(slot func(item *QTreeWidgetItem)) {
	QTreeWidget_connect_ItemExpanded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemExpanded
func miqt_exec_callback_QTreeWidget_ItemExpanded(cb intptr_t, item *QTreeWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTreeWidget) ItemCollapsed(item *QTreeWidgetItem) {
	QTreeWidget_ItemCollapsed(this.h, item.cPointer())
}

func (this *QTreeWidget) OnItemCollapsed(slot func(item *QTreeWidgetItem)) {
	QTreeWidget_connect_ItemCollapsed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemCollapsed
func miqt_exec_callback_QTreeWidget_ItemCollapsed(cb intptr_t, item *QTreeWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(item *QTreeWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(item)

	gofunc(slotval1)
}

func (this *QTreeWidget) CurrentItemChanged(current *QTreeWidgetItem, previous *QTreeWidgetItem) {
	QTreeWidget_CurrentItemChanged(this.h, current.cPointer(), previous.cPointer())
}

func (this *QTreeWidget) OnCurrentItemChanged(slot func(current *QTreeWidgetItem, previous *QTreeWidgetItem)) {
	QTreeWidget_connect_CurrentItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_CurrentItemChanged
func miqt_exec_callback_QTreeWidget_CurrentItemChanged(cb intptr_t, current *QTreeWidgetItem, previous *QTreeWidgetItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QTreeWidgetItem, previous *QTreeWidgetItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(current)

	slotval2 := newQTreeWidgetItem(previous)

	gofunc(slotval1, slotval2)
}

func (this *QTreeWidget) ItemSelectionChanged() {
	QTreeWidget_ItemSelectionChanged(this.h)
}

func (this *QTreeWidget) OnItemSelectionChanged(slot func()) {
	QTreeWidget_connect_ItemSelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ItemSelectionChanged
func miqt_exec_callback_QTreeWidget_ItemSelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QTreeWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTreeWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTreeWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTreeWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTreeWidget) EditItem2(item *QTreeWidgetItem, column int) {
	QTreeWidget_EditItem2(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) OpenPersistentEditor2(item *QTreeWidgetItem, column int) {
	QTreeWidget_OpenPersistentEditor2(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) ClosePersistentEditor2(item *QTreeWidgetItem, column int) {
	QTreeWidget_ClosePersistentEditor2(this.h, item.cPointer(), (int)(column))
}

func (this *QTreeWidget) IsPersistentEditorOpen2(item *QTreeWidgetItem, column int) bool {
	return (bool)(QTreeWidget_IsPersistentEditorOpen2(this.h, item.cPointer(), (int)(column)))
}

func (this *QTreeWidget) FindItems3(text string, flags MatchFlag, column int) []*QTreeWidgetItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QTreeWidget_FindItems3(this.h, text_ms, (int)(flags), (int)(column))
	_ret := make([]*QTreeWidgetItem, int(_ma.len))
	_outCast := (*[0xffff]*QTreeWidgetItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTreeWidgetItem(_outCast[i])
	}
	return _ret
}

func (this *QTreeWidget) IndexFromItem2(item *QTreeWidgetItem, column int) *QModelIndex {
	_goptr := newQModelIndex(QTreeWidget_IndexFromItem2(this.h, item.cPointer(), (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidget) ScrollToItem2(item *QTreeWidgetItem, hint QAbstractItemView__ScrollHint) {
	QTreeWidget_ScrollToItem2(this.h, item.cPointer(), (int)(hint))
}

func (this *QTreeWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTreeWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTreeWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MetaObject
func miqt_exec_callback_QTreeWidget_MetaObject(self QTreeWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTreeWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTreeWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTreeWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_Metacast
func miqt_exec_callback_QTreeWidget_Metacast(self QTreeWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
