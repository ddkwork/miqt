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

func (this *QTreeWidgetItem) callVirtualBase_Clone() *QTreeWidgetItem {

	return newQTreeWidgetItem(QTreeWidgetItem_virtualbase_Clone(unsafe.Pointer(this.h)))

}
func (this *QTreeWidgetItem) OnClone(slot func(super func() *QTreeWidgetItem) *QTreeWidgetItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidgetItem_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidgetItem_Clone
func miqt_exec_callback_QTreeWidgetItem_Clone(self QTreeWidgetItem, cb intptr_t) *QTreeWidgetItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QTreeWidgetItem) *QTreeWidgetItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidgetItem{h: self}).callVirtualBase_Clone)

	return virtualReturn.cPointer()

}

func (this *QTreeWidgetItem) callVirtualBase_Data(column int, role int) *QVariant {

	_goptr := newQVariant(QTreeWidgetItem_virtualbase_Data(unsafe.Pointer(this.h), (int)(column), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeWidgetItem) OnData(slot func(super func(column int, role int) *QVariant, column int, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidgetItem_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidgetItem_Data
func miqt_exec_callback_QTreeWidgetItem_Data(self QTreeWidgetItem, cb intptr_t, column int, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, role int) *QVariant, column int, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(role)

	virtualReturn := gofunc((&QTreeWidgetItem{h: self}).callVirtualBase_Data, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QTreeWidgetItem) callVirtualBase_SetData(column int, role int, value *QVariant) {

	QTreeWidgetItem_virtualbase_SetData(unsafe.Pointer(this.h), (int)(column), (int)(role), value.cPointer())

}
func (this *QTreeWidgetItem) OnSetData(slot func(super func(column int, role int, value *QVariant), column int, role int, value *QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidgetItem_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidgetItem_SetData
func miqt_exec_callback_QTreeWidgetItem_SetData(self QTreeWidgetItem, cb intptr_t, column int, role int, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, role int, value *QVariant), column int, role int, value *QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(role)

	slotval3 := newQVariant(value)

	gofunc((&QTreeWidgetItem{h: self}).callVirtualBase_SetData, slotval1, slotval2, slotval3)

}

func (this *QTreeWidgetItem) callVirtualBase_OperatorLesser(other *QTreeWidgetItem) bool {

	return (bool)(QTreeWidgetItem_virtualbase_OperatorLesser(unsafe.Pointer(this.h), other.cPointer()))

}
func (this *QTreeWidgetItem) OnOperatorLesser(slot func(super func(other *QTreeWidgetItem) bool, other *QTreeWidgetItem) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidgetItem_override_virtual_OperatorLesser(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidgetItem_OperatorLesser
func miqt_exec_callback_QTreeWidgetItem_OperatorLesser(self QTreeWidgetItem, cb intptr_t, other *QTreeWidgetItem) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(other *QTreeWidgetItem) bool, other *QTreeWidgetItem) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(other)

	virtualReturn := gofunc((&QTreeWidgetItem{h: self}).callVirtualBase_OperatorLesser, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeWidgetItem) callVirtualBase_Read(in *QDataStream) {

	QTreeWidgetItem_virtualbase_Read(unsafe.Pointer(this.h), in.cPointer())

}
func (this *QTreeWidgetItem) OnRead(slot func(super func(in *QDataStream), in *QDataStream)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidgetItem_override_virtual_Read(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidgetItem_Read
func miqt_exec_callback_QTreeWidgetItem_Read(self QTreeWidgetItem, cb intptr_t, in *QDataStream) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(in *QDataStream), in *QDataStream))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(in)

	gofunc((&QTreeWidgetItem{h: self}).callVirtualBase_Read, slotval1)

}

func (this *QTreeWidgetItem) callVirtualBase_Write(out *QDataStream) {

	QTreeWidgetItem_virtualbase_Write(unsafe.Pointer(this.h), out.cPointer())

}
func (this *QTreeWidgetItem) OnWrite(slot func(super func(out *QDataStream), out *QDataStream)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidgetItem_override_virtual_Write(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidgetItem_Write
func miqt_exec_callback_QTreeWidgetItem_Write(self QTreeWidgetItem, cb intptr_t, out *QDataStream) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(out *QDataStream), out *QDataStream))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(out)

	gofunc((&QTreeWidgetItem{h: self}).callVirtualBase_Write, slotval1)

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

func (this *QTreeWidget) callVirtualBase_SetSelectionModel(selectionModel *QItemSelectionModel) {

	QTreeWidget_virtualbase_SetSelectionModel(unsafe.Pointer(this.h), selectionModel.cPointer())

}
func (this *QTreeWidget) OnSetSelectionModel(slot func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SetSelectionModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SetSelectionModel
func miqt_exec_callback_QTreeWidget_SetSelectionModel(self QTreeWidget, cb intptr_t, selectionModel *QItemSelectionModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelectionModel(selectionModel)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_SetSelectionModel, slotval1)

}

func (this *QTreeWidget) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QTreeWidget_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QTreeWidget) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_Event
func miqt_exec_callback_QTreeWidget_Event(self QTreeWidget, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_MimeTypes() []string {

	var _ma struct_miqt_array = QTreeWidget_virtualbase_MimeTypes(unsafe.Pointer(this.h))
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
func (this *QTreeWidget) OnMimeTypes(slot func(super func() []string) []string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MimeTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MimeTypes
func miqt_exec_callback_QTreeWidget_MimeTypes(self QTreeWidget, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_MimeTypes)
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

func (this *QTreeWidget) callVirtualBase_MimeData(items []*QTreeWidgetItem) *QMimeData {
	items_CArray := (*[0xffff]*QTreeWidgetItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}

	return newQMimeData(QTreeWidget_virtualbase_MimeData(unsafe.Pointer(this.h), items_ma))

}
func (this *QTreeWidget) OnMimeData(slot func(super func(items []*QTreeWidgetItem) *QMimeData, items []*QTreeWidgetItem) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MimeData
func miqt_exec_callback_QTreeWidget_MimeData(self QTreeWidget, cb intptr_t, items struct_miqt_array) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(items []*QTreeWidgetItem) *QMimeData, items []*QTreeWidgetItem) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var items_ma struct_miqt_array = items
	items_ret := make([]*QTreeWidgetItem, int(items_ma.len))
	items_outCast := (*[0xffff]*QTreeWidgetItem)(unsafe.Pointer(items_ma.data)) // hey ya
	for i := 0; i < int(items_ma.len); i++ {
		items_ret[i] = newQTreeWidgetItem(items_outCast[i])
	}
	slotval1 := items_ret

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_MimeData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeWidget) callVirtualBase_DropMimeData(parent *QTreeWidgetItem, index int, data *QMimeData, action DropAction) bool {

	return (bool)(QTreeWidget_virtualbase_DropMimeData(unsafe.Pointer(this.h), parent.cPointer(), (int)(index), data.cPointer(), (int)(action)))

}
func (this *QTreeWidget) OnDropMimeData(slot func(super func(parent *QTreeWidgetItem, index int, data *QMimeData, action DropAction) bool, parent *QTreeWidgetItem, index int, data *QMimeData, action DropAction) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_DropMimeData
func miqt_exec_callback_QTreeWidget_DropMimeData(self QTreeWidget, cb intptr_t, parent *QTreeWidgetItem, index int, data *QMimeData, action int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QTreeWidgetItem, index int, data *QMimeData, action DropAction) bool, parent *QTreeWidgetItem, index int, data *QMimeData, action DropAction) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTreeWidgetItem(parent)

	slotval2 := (int)(index)

	slotval3 := newQMimeData(data)

	slotval4 := (DropAction)(action)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QTreeWidget_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QTreeWidget) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SupportedDropActions
func miqt_exec_callback_QTreeWidget_SupportedDropActions(self QTreeWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_DropEvent(event *QDropEvent) {

	QTreeWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_DropEvent
func miqt_exec_callback_QTreeWidget_DropEvent(self QTreeWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QTreeWidget_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QTreeWidget) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SetRootIndex
func miqt_exec_callback_QTreeWidget_SetRootIndex(self QTreeWidget, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QTreeWidget) callVirtualBase_KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))

	QTreeWidget_virtualbase_KeyboardSearch(unsafe.Pointer(this.h), search_ms)

}
func (this *QTreeWidget) OnKeyboardSearch(slot func(super func(search string), search string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_KeyboardSearch(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_KeyboardSearch
func miqt_exec_callback_QTreeWidget_KeyboardSearch(self QTreeWidget, cb intptr_t, search struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(search string), search string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var search_ms struct_miqt_string = search
	search_ret := GoStringN(search_ms.data, int(int64(search_ms.len)))
	free(unsafe.Pointer(search_ms.data))
	slotval1 := search_ret

	gofunc((&QTreeWidget{h: self}).callVirtualBase_KeyboardSearch, slotval1)

}

func (this *QTreeWidget) callVirtualBase_VisualRect(index *QModelIndex) *QRect {

	_goptr := newQRect(QTreeWidget_virtualbase_VisualRect(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeWidget) OnVisualRect(slot func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_VisualRect
func miqt_exec_callback_QTreeWidget_VisualRect(self QTreeWidget, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_VisualRect, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeWidget) callVirtualBase_ScrollTo(index *QModelIndex, hint ScrollHint) {

	QTreeWidget_virtualbase_ScrollTo(unsafe.Pointer(this.h), index.cPointer(), hint)

}
func (this *QTreeWidget) OnScrollTo(slot func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ScrollTo
func miqt_exec_callback_QTreeWidget_ScrollTo(self QTreeWidget, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc((&QTreeWidget{h: self}).callVirtualBase_ScrollTo, slotval1, slotval2)

}

func (this *QTreeWidget) callVirtualBase_IndexAt(p *QPoint) *QModelIndex {

	_goptr := newQModelIndex(QTreeWidget_virtualbase_IndexAt(unsafe.Pointer(this.h), p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeWidget) OnIndexAt(slot func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_IndexAt
func miqt_exec_callback_QTreeWidget_IndexAt(self QTreeWidget, cb intptr_t, p *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(p)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_IndexAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeWidget) callVirtualBase_DoItemsLayout() {

	QTreeWidget_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QTreeWidget) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_DoItemsLayout
func miqt_exec_callback_QTreeWidget_DoItemsLayout(self QTreeWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeWidget{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QTreeWidget) callVirtualBase_Reset() {

	QTreeWidget_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QTreeWidget) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_Reset
func miqt_exec_callback_QTreeWidget_Reset(self QTreeWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeWidget{h: self}).callVirtualBase_Reset)

}

func (this *QTreeWidget) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QTreeWidget_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QTreeWidget) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_DataChanged
func miqt_exec_callback_QTreeWidget_DataChanged(self QTreeWidget, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
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

	gofunc((&QTreeWidget{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QTreeWidget) callVirtualBase_SelectAll() {

	QTreeWidget_virtualbase_SelectAll(unsafe.Pointer(this.h))

}
func (this *QTreeWidget) OnSelectAll(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SelectAll(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SelectAll
func miqt_exec_callback_QTreeWidget_SelectAll(self QTreeWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeWidget{h: self}).callVirtualBase_SelectAll)

}

func (this *QTreeWidget) callVirtualBase_VerticalScrollbarValueChanged(value int) {

	QTreeWidget_virtualbase_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QTreeWidget) OnVerticalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_VerticalScrollbarValueChanged
func miqt_exec_callback_QTreeWidget_VerticalScrollbarValueChanged(self QTreeWidget, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_VerticalScrollbarValueChanged, slotval1)

}

func (this *QTreeWidget) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QTreeWidget_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QTreeWidget) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ScrollContentsBy
func miqt_exec_callback_QTreeWidget_ScrollContentsBy(self QTreeWidget, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QTreeWidget) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QTreeWidget_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QTreeWidget) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_RowsInserted
func miqt_exec_callback_QTreeWidget_RowsInserted(self QTreeWidget, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QTreeWidget) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QTreeWidget_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QTreeWidget) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_RowsAboutToBeRemoved
func miqt_exec_callback_QTreeWidget_RowsAboutToBeRemoved(self QTreeWidget, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QTreeWidget) callVirtualBase_MoveCursor(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex {

	_goptr := newQModelIndex(QTreeWidget_virtualbase_MoveCursor(unsafe.Pointer(this.h), cursorAction, (int)(modifiers)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeWidget) OnMoveCursor(slot func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MoveCursor
func miqt_exec_callback_QTreeWidget_MoveCursor(self QTreeWidget, cb intptr_t, cursorAction CursorAction, modifiers int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(modifiers)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_MoveCursor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QTreeWidget) callVirtualBase_HorizontalOffset() int {

	return (int)(QTreeWidget_virtualbase_HorizontalOffset(unsafe.Pointer(this.h)))

}
func (this *QTreeWidget) OnHorizontalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_HorizontalOffset
func miqt_exec_callback_QTreeWidget_HorizontalOffset(self QTreeWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_HorizontalOffset)

	return (int)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_VerticalOffset() int {

	return (int)(QTreeWidget_virtualbase_VerticalOffset(unsafe.Pointer(this.h)))

}
func (this *QTreeWidget) OnVerticalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_VerticalOffset
func miqt_exec_callback_QTreeWidget_VerticalOffset(self QTreeWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_VerticalOffset)

	return (int)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_SetSelection(rect *QRect, command SelectionFlag) {

	QTreeWidget_virtualbase_SetSelection(unsafe.Pointer(this.h), rect.cPointer(), (int)(command))

}
func (this *QTreeWidget) OnSetSelection(slot func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SetSelection
func miqt_exec_callback_QTreeWidget_SetSelection(self QTreeWidget, cb intptr_t, rect *QRect, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_SetSelection, slotval1, slotval2)

}

func (this *QTreeWidget) callVirtualBase_VisualRegionForSelection(selection *QItemSelection) *QRegion {

	_goptr := newQRegion(QTreeWidget_virtualbase_VisualRegionForSelection(unsafe.Pointer(this.h), selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeWidget) OnVisualRegionForSelection(slot func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_VisualRegionForSelection
func miqt_exec_callback_QTreeWidget_VisualRegionForSelection(self QTreeWidget, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_VisualRegionForSelection, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTreeWidget) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QTreeWidget_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QTreeWidget) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SelectedIndexes
func miqt_exec_callback_QTreeWidget_SelectedIndexes(self QTreeWidget, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QTreeWidget) callVirtualBase_ChangeEvent(event *QEvent) {

	QTreeWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ChangeEvent
func miqt_exec_callback_QTreeWidget_ChangeEvent(self QTreeWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QTreeWidget_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_TimerEvent
func miqt_exec_callback_QTreeWidget_TimerEvent(self QTreeWidget, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QTreeWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_PaintEvent
func miqt_exec_callback_QTreeWidget_PaintEvent(self QTreeWidget, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_DrawRow(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex) {

	QTreeWidget_virtualbase_DrawRow(unsafe.Pointer(this.h), painter.cPointer(), options.cPointer(), index.cPointer())

}
func (this *QTreeWidget) OnDrawRow(slot func(super func(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_DrawRow(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_DrawRow
func miqt_exec_callback_QTreeWidget_DrawRow(self QTreeWidget, cb intptr_t, painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(options)

	slotval3 := newQModelIndex(index)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_DrawRow, slotval1, slotval2, slotval3)

}

func (this *QTreeWidget) callVirtualBase_DrawBranches(painter *QPainter, rect *QRect, index *QModelIndex) {

	QTreeWidget_virtualbase_DrawBranches(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), index.cPointer())

}
func (this *QTreeWidget) OnDrawBranches(slot func(super func(painter *QPainter, rect *QRect, index *QModelIndex), painter *QPainter, rect *QRect, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_DrawBranches(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_DrawBranches
func miqt_exec_callback_QTreeWidget_DrawBranches(self QTreeWidget, cb intptr_t, painter *QPainter, rect *QRect, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, index *QModelIndex), painter *QPainter, rect *QRect, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	slotval3 := newQModelIndex(index)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_DrawBranches, slotval1, slotval2, slotval3)

}

func (this *QTreeWidget) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QTreeWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MousePressEvent
func miqt_exec_callback_QTreeWidget_MousePressEvent(self QTreeWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QTreeWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MouseReleaseEvent
func miqt_exec_callback_QTreeWidget_MouseReleaseEvent(self QTreeWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QTreeWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MouseDoubleClickEvent
func miqt_exec_callback_QTreeWidget_MouseDoubleClickEvent(self QTreeWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QTreeWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_MouseMoveEvent
func miqt_exec_callback_QTreeWidget_MouseMoveEvent(self QTreeWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QTreeWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_KeyPressEvent
func miqt_exec_callback_QTreeWidget_KeyPressEvent(self QTreeWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QTreeWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTreeWidget) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_DragMoveEvent
func miqt_exec_callback_QTreeWidget_DragMoveEvent(self QTreeWidget, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTreeWidget) callVirtualBase_ViewportEvent(event *QEvent) bool {

	return (bool)(QTreeWidget_virtualbase_ViewportEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTreeWidget) OnViewportEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ViewportEvent
func miqt_exec_callback_QTreeWidget_ViewportEvent(self QTreeWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_UpdateGeometries() {

	QTreeWidget_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QTreeWidget) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_UpdateGeometries
func miqt_exec_callback_QTreeWidget_UpdateGeometries(self QTreeWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTreeWidget{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QTreeWidget) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QTreeWidget_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTreeWidget) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_ViewportSizeHint
func miqt_exec_callback_QTreeWidget_ViewportSizeHint(self QTreeWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}

func (this *QTreeWidget) callVirtualBase_SizeHintForColumn(column int) int {

	return (int)(QTreeWidget_virtualbase_SizeHintForColumn(unsafe.Pointer(this.h), (int)(column)))

}
func (this *QTreeWidget) OnSizeHintForColumn(slot func(super func(column int) int, column int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SizeHintForColumn(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SizeHintForColumn
func miqt_exec_callback_QTreeWidget_SizeHintForColumn(self QTreeWidget, cb intptr_t, column int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int) int, column int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_SizeHintForColumn, slotval1)

	return (int)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_HorizontalScrollbarAction(action int) {

	QTreeWidget_virtualbase_HorizontalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QTreeWidget) OnHorizontalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_HorizontalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_HorizontalScrollbarAction
func miqt_exec_callback_QTreeWidget_HorizontalScrollbarAction(self QTreeWidget, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_HorizontalScrollbarAction, slotval1)

}

func (this *QTreeWidget) callVirtualBase_IsIndexHidden(index *QModelIndex) bool {

	return (bool)(QTreeWidget_virtualbase_IsIndexHidden(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QTreeWidget) OnIsIndexHidden(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_IsIndexHidden
func miqt_exec_callback_QTreeWidget_IsIndexHidden(self QTreeWidget, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QTreeWidget{h: self}).callVirtualBase_IsIndexHidden, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTreeWidget) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QTreeWidget_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QTreeWidget) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_SelectionChanged
func miqt_exec_callback_QTreeWidget_SelectionChanged(self QTreeWidget, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QTreeWidget) callVirtualBase_CurrentChanged(current *QModelIndex, previous *QModelIndex) {

	QTreeWidget_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), previous.cPointer())

}
func (this *QTreeWidget) OnCurrentChanged(slot func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTreeWidget_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTreeWidget_CurrentChanged
func miqt_exec_callback_QTreeWidget_CurrentChanged(self QTreeWidget, cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc((&QTreeWidget{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}
