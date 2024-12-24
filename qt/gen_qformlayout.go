package qt

import (
	"unsafe"
)

type QFormLayout__FieldGrowthPolicy int

const (
	QFormLayout__FieldsStayAtSizeHint  QFormLayout__FieldGrowthPolicy = 0
	QFormLayout__ExpandingFieldsGrow   QFormLayout__FieldGrowthPolicy = 1
	QFormLayout__AllNonFixedFieldsGrow QFormLayout__FieldGrowthPolicy = 2
)

type QFormLayout__RowWrapPolicy int

const (
	QFormLayout__DontWrapRows QFormLayout__RowWrapPolicy = 0
	QFormLayout__WrapLongRows QFormLayout__RowWrapPolicy = 1
	QFormLayout__WrapAllRows  QFormLayout__RowWrapPolicy = 2
)

type QFormLayout__ItemRole int

const (
	QFormLayout__LabelRole    QFormLayout__ItemRole = 0
	QFormLayout__FieldRole    QFormLayout__ItemRole = 1
	QFormLayout__SpanningRole QFormLayout__ItemRole = 2
)

type QFormLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQFormLayout constructs a new QFormLayout object.
func NewQFormLayout(parent *QWidget) *QFormLayout {

	ret := newQFormLayout(QFormLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFormLayout2 constructs a new QFormLayout object.
func NewQFormLayout2() *QFormLayout {

	ret := newQFormLayout(QFormLayout_new2())
	ret.isSubclass = true
	return ret
}

func (this *QFormLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QFormLayout_MetaObject(this.h))
}

func (this *QFormLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFormLayout_Metacast(this.h, param1_Cstring))
}

func QFormLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFormLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFormLayout) SetFieldGrowthPolicy(policy FieldGrowthPolicy) {
	QFormLayout_SetFieldGrowthPolicy(this.h, policy)
}

func (this *QFormLayout) FieldGrowthPolicy() FieldGrowthPolicy {
	xxxxxxxxx
}

func (this *QFormLayout) SetRowWrapPolicy(policy RowWrapPolicy) {
	QFormLayout_SetRowWrapPolicy(this.h, policy)
}

func (this *QFormLayout) RowWrapPolicy() RowWrapPolicy {
	xxxxxxxxx
}

func (this *QFormLayout) SetLabelAlignment(alignment AlignmentFlag) {
	QFormLayout_SetLabelAlignment(this.h, (int)(alignment))
}

func (this *QFormLayout) LabelAlignment() AlignmentFlag {
	return (AlignmentFlag)(QFormLayout_LabelAlignment(this.h))
}

func (this *QFormLayout) SetFormAlignment(alignment AlignmentFlag) {
	QFormLayout_SetFormAlignment(this.h, (int)(alignment))
}

func (this *QFormLayout) FormAlignment() AlignmentFlag {
	return (AlignmentFlag)(QFormLayout_FormAlignment(this.h))
}

func (this *QFormLayout) SetHorizontalSpacing(spacing int) {
	QFormLayout_SetHorizontalSpacing(this.h, (int)(spacing))
}

func (this *QFormLayout) HorizontalSpacing() int {
	return (int)(QFormLayout_HorizontalSpacing(this.h))
}

func (this *QFormLayout) SetVerticalSpacing(spacing int) {
	QFormLayout_SetVerticalSpacing(this.h, (int)(spacing))
}

func (this *QFormLayout) VerticalSpacing() int {
	return (int)(QFormLayout_VerticalSpacing(this.h))
}

func (this *QFormLayout) Spacing() int {
	return (int)(QFormLayout_Spacing(this.h))
}

func (this *QFormLayout) SetSpacing(spacing int) {
	QFormLayout_SetSpacing(this.h, (int)(spacing))
}

func (this *QFormLayout) AddRow(label *QWidget, field *QWidget) {
	QFormLayout_AddRow(this.h, label.cPointer(), field.cPointer())
}

func (this *QFormLayout) AddRow2(label *QWidget, field *QLayout) {
	QFormLayout_AddRow2(this.h, label.cPointer(), field.cPointer())
}

func (this *QFormLayout) AddRow3(labelText string, field *QWidget) {
	labelText_ms := struct_miqt_string{}
	labelText_ms.data = CString(labelText)
	labelText_ms.len = size_t(len(labelText))
	defer free(unsafe.Pointer(labelText_ms.data))
	QFormLayout_AddRow3(this.h, labelText_ms, field.cPointer())
}

func (this *QFormLayout) AddRow4(labelText string, field *QLayout) {
	labelText_ms := struct_miqt_string{}
	labelText_ms.data = CString(labelText)
	labelText_ms.len = size_t(len(labelText))
	defer free(unsafe.Pointer(labelText_ms.data))
	QFormLayout_AddRow4(this.h, labelText_ms, field.cPointer())
}

func (this *QFormLayout) AddRowWithWidget(widget *QWidget) {
	QFormLayout_AddRowWithWidget(this.h, widget.cPointer())
}

func (this *QFormLayout) AddRowWithLayout(layout *QLayout) {
	QFormLayout_AddRowWithLayout(this.h, layout.cPointer())
}

func (this *QFormLayout) InsertRow(row int, label *QWidget, field *QWidget) {
	QFormLayout_InsertRow(this.h, (int)(row), label.cPointer(), field.cPointer())
}

func (this *QFormLayout) InsertRow2(row int, label *QWidget, field *QLayout) {
	QFormLayout_InsertRow2(this.h, (int)(row), label.cPointer(), field.cPointer())
}

func (this *QFormLayout) InsertRow3(row int, labelText string, field *QWidget) {
	labelText_ms := struct_miqt_string{}
	labelText_ms.data = CString(labelText)
	labelText_ms.len = size_t(len(labelText))
	defer free(unsafe.Pointer(labelText_ms.data))
	QFormLayout_InsertRow3(this.h, (int)(row), labelText_ms, field.cPointer())
}

func (this *QFormLayout) InsertRow4(row int, labelText string, field *QLayout) {
	labelText_ms := struct_miqt_string{}
	labelText_ms.data = CString(labelText)
	labelText_ms.len = size_t(len(labelText))
	defer free(unsafe.Pointer(labelText_ms.data))
	QFormLayout_InsertRow4(this.h, (int)(row), labelText_ms, field.cPointer())
}

func (this *QFormLayout) InsertRow5(row int, widget *QWidget) {
	QFormLayout_InsertRow5(this.h, (int)(row), widget.cPointer())
}

func (this *QFormLayout) InsertRow6(row int, layout *QLayout) {
	QFormLayout_InsertRow6(this.h, (int)(row), layout.cPointer())
}

func (this *QFormLayout) RemoveRow(row int) {
	QFormLayout_RemoveRow(this.h, (int)(row))
}

func (this *QFormLayout) RemoveRowWithWidget(widget *QWidget) {
	QFormLayout_RemoveRowWithWidget(this.h, widget.cPointer())
}

func (this *QFormLayout) RemoveRowWithLayout(layout *QLayout) {
	QFormLayout_RemoveRowWithLayout(this.h, layout.cPointer())
}

func (this *QFormLayout) TakeRow(row int) TakeRowResult {
	xxxxxxxxx
}

func (this *QFormLayout) TakeRowWithWidget(widget *QWidget) TakeRowResult {
	xxxxxxxxx
}

func (this *QFormLayout) TakeRowWithLayout(layout *QLayout) TakeRowResult {
	xxxxxxxxx
}

func (this *QFormLayout) SetItem(row int, role ItemRole, item *QLayoutItem) {
	QFormLayout_SetItem(this.h, (int)(row), role, item.cPointer())
}

func (this *QFormLayout) SetWidget(row int, role ItemRole, widget *QWidget) {
	QFormLayout_SetWidget(this.h, (int)(row), role, widget.cPointer())
}

func (this *QFormLayout) SetLayout(row int, role ItemRole, layout *QLayout) {
	QFormLayout_SetLayout(this.h, (int)(row), role, layout.cPointer())
}

func (this *QFormLayout) SetRowVisible(row int, on bool) {
	QFormLayout_SetRowVisible(this.h, (int)(row), (bool)(on))
}

func (this *QFormLayout) SetRowVisible2(widget *QWidget, on bool) {
	QFormLayout_SetRowVisible2(this.h, widget.cPointer(), (bool)(on))
}

func (this *QFormLayout) SetRowVisible3(layout *QLayout, on bool) {
	QFormLayout_SetRowVisible3(this.h, layout.cPointer(), (bool)(on))
}

func (this *QFormLayout) IsRowVisible(row int) bool {
	return (bool)(QFormLayout_IsRowVisible(this.h, (int)(row)))
}

func (this *QFormLayout) IsRowVisibleWithWidget(widget *QWidget) bool {
	return (bool)(QFormLayout_IsRowVisibleWithWidget(this.h, widget.cPointer()))
}

func (this *QFormLayout) IsRowVisibleWithLayout(layout *QLayout) bool {
	return (bool)(QFormLayout_IsRowVisibleWithLayout(this.h, layout.cPointer()))
}

func (this *QFormLayout) ItemAt(row int, role ItemRole) *QLayoutItem {
	return newQLayoutItem(QFormLayout_ItemAt(this.h, (int)(row), role))
}

func (this *QFormLayout) GetItemPosition(index int, rowPtr *int, rolePtr *ItemRole) {
	QFormLayout_GetItemPosition(this.h, (int)(index), (*int)(unsafe.Pointer(rowPtr)), rolePtr)
}

func (this *QFormLayout) GetWidgetPosition(widget *QWidget, rowPtr *int, rolePtr *ItemRole) {
	QFormLayout_GetWidgetPosition(this.h, widget.cPointer(), (*int)(unsafe.Pointer(rowPtr)), rolePtr)
}

func (this *QFormLayout) GetLayoutPosition(layout *QLayout, rowPtr *int, rolePtr *ItemRole) {
	QFormLayout_GetLayoutPosition(this.h, layout.cPointer(), (*int)(unsafe.Pointer(rowPtr)), rolePtr)
}

func (this *QFormLayout) LabelForField(field *QWidget) *QWidget {
	return newQWidget(QFormLayout_LabelForField(this.h, field.cPointer()))
}

func (this *QFormLayout) LabelForFieldWithField(field *QLayout) *QWidget {
	return newQWidget(QFormLayout_LabelForFieldWithField(this.h, field.cPointer()))
}

func (this *QFormLayout) AddItem(item *QLayoutItem) {
	QFormLayout_AddItem(this.h, item.cPointer())
}

func (this *QFormLayout) ItemAtWithIndex(index int) *QLayoutItem {
	return newQLayoutItem(QFormLayout_ItemAtWithIndex(this.h, (int)(index)))
}

func (this *QFormLayout) TakeAt(index int) *QLayoutItem {
	return newQLayoutItem(QFormLayout_TakeAt(this.h, (int)(index)))
}

func (this *QFormLayout) SetGeometry(rect *QRect) {
	QFormLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QFormLayout) MinimumSize() *QSize {
	_goptr := newQSize(QFormLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormLayout) SizeHint() *QSize {
	_goptr := newQSize(QFormLayout_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFormLayout) Invalidate() {
	QFormLayout_Invalidate(this.h)
}

func (this *QFormLayout) HasHeightForWidth() bool {
	return (bool)(QFormLayout_HasHeightForWidth(this.h))
}

func (this *QFormLayout) HeightForWidth(width int) int {
	return (int)(QFormLayout_HeightForWidth(this.h, (int)(width)))
}

func (this *QFormLayout) ExpandingDirections() Orientation {
	return (Orientation)(QFormLayout_ExpandingDirections(this.h))
}

func (this *QFormLayout) Count() int {
	return (int)(QFormLayout_Count(this.h))
}

func (this *QFormLayout) RowCount() int {
	return (int)(QFormLayout_RowCount(this.h))
}

func QFormLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFormLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFormLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFormLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFormLayout) callVirtualBase_Spacing() int {

	return (int)(QFormLayout_virtualbase_Spacing(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnSpacing(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_Spacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Spacing
func miqt_exec_callback_QFormLayout_Spacing(self QFormLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Spacing)

	return (int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_SetSpacing(spacing int) {

	QFormLayout_virtualbase_SetSpacing(unsafe.Pointer(this.h), (int)(spacing))

}
func (this *QFormLayout) OnSetSpacing(slot func(super func(spacing int), spacing int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_SetSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_SetSpacing
func miqt_exec_callback_QFormLayout_SetSpacing(self QFormLayout, cb intptr_t, spacing int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(spacing int), spacing int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(spacing)

	gofunc((&QFormLayout{h: self}).callVirtualBase_SetSpacing, slotval1)

}

func (this *QFormLayout) callVirtualBase_AddItem(item *QLayoutItem) {

	QFormLayout_virtualbase_AddItem(unsafe.Pointer(this.h), item.cPointer())

}
func (this *QFormLayout) OnAddItem(slot func(super func(item *QLayoutItem), item *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_AddItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_AddItem
func miqt_exec_callback_QFormLayout_AddItem(self QFormLayout, cb intptr_t, item *QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(item *QLayoutItem), item *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(item)

	gofunc((&QFormLayout{h: self}).callVirtualBase_AddItem, slotval1)

}

func (this *QFormLayout) callVirtualBase_ItemAtWithIndex(index int) *QLayoutItem {

	return newQLayoutItem(QFormLayout_virtualbase_ItemAtWithIndex(unsafe.Pointer(this.h), (int)(index)))

}
func (this *QFormLayout) OnItemAtWithIndex(slot func(super func(index int) *QLayoutItem, index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_ItemAtWithIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ItemAtWithIndex
func miqt_exec_callback_QFormLayout_ItemAtWithIndex(self QFormLayout, cb intptr_t, index int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QLayoutItem, index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_ItemAtWithIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_TakeAt(index int) *QLayoutItem {

	return newQLayoutItem(QFormLayout_virtualbase_TakeAt(unsafe.Pointer(this.h), (int)(index)))

}
func (this *QFormLayout) OnTakeAt(slot func(super func(index int) *QLayoutItem, index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_TakeAt
func miqt_exec_callback_QFormLayout_TakeAt(self QFormLayout, cb intptr_t, index int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QLayoutItem, index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_TakeAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_SetGeometry(rect *QRect) {

	QFormLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QFormLayout) OnSetGeometry(slot func(super func(rect *QRect), rect *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_SetGeometry
func miqt_exec_callback_QFormLayout_SetGeometry(self QFormLayout, cb intptr_t, rect *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect), rect *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	gofunc((&QFormLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QFormLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QFormLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_MinimumSize
func miqt_exec_callback_QFormLayout_MinimumSize(self QFormLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QFormLayout_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_SizeHint
func miqt_exec_callback_QFormLayout_SizeHint(self QFormLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_Invalidate() {

	QFormLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QFormLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Invalidate
func miqt_exec_callback_QFormLayout_Invalidate(self QFormLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFormLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QFormLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QFormLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_HasHeightForWidth
func miqt_exec_callback_QFormLayout_HasHeightForWidth(self QFormLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_HeightForWidth(width int) int {

	return (int)(QFormLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(width)))

}
func (this *QFormLayout) OnHeightForWidth(slot func(super func(width int) int, width int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_HeightForWidth
func miqt_exec_callback_QFormLayout_HeightForWidth(self QFormLayout, cb intptr_t, width int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(width int) int, width int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(width)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QFormLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ExpandingDirections
func miqt_exec_callback_QFormLayout_ExpandingDirections(self QFormLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_Count() int {

	return (int)(QFormLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Count
func miqt_exec_callback_QFormLayout_Count(self QFormLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QFormLayout_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Geometry
func miqt_exec_callback_QFormLayout_Geometry(self QFormLayout, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QFormLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFormLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_MaximumSize
func miqt_exec_callback_QFormLayout_MaximumSize(self QFormLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_IndexOf(param1 *QWidget) int {

	return (int)(QFormLayout_virtualbase_IndexOf(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QFormLayout) OnIndexOf(slot func(super func(param1 *QWidget) int, param1 *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_IndexOf(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_IndexOf
func miqt_exec_callback_QFormLayout_IndexOf(self QFormLayout, cb intptr_t, param1 *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWidget) int, param1 *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(param1)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_IndexOf, slotval1)

	return (int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_IsEmpty() bool {

	return (bool)(QFormLayout_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_IsEmpty
func miqt_exec_callback_QFormLayout_IsEmpty(self QFormLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QFormLayout_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ControlTypes
func miqt_exec_callback_QFormLayout_ControlTypes(self QFormLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

func (this *QFormLayout) callVirtualBase_ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {

	return newQLayoutItem(QFormLayout_virtualbase_ReplaceWidget(unsafe.Pointer(this.h), from.cPointer(), to.cPointer(), (int)(options)))

}
func (this *QFormLayout) OnReplaceWidget(slot func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_ReplaceWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ReplaceWidget
func miqt_exec_callback_QFormLayout_ReplaceWidget(self QFormLayout, cb intptr_t, from *QWidget, to *QWidget, options int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(from)

	slotval2 := newQWidget(to)

	slotval3 := (FindChildOption)(options)

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_ReplaceWidget, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_Layout() *QLayout {

	return newQLayout(QFormLayout_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QFormLayout) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Layout
func miqt_exec_callback_QFormLayout_Layout(self QFormLayout, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QFormLayout) callVirtualBase_ChildEvent(e *QChildEvent) {

	QFormLayout_virtualbase_ChildEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFormLayout) OnChildEvent(slot func(super func(e *QChildEvent), e *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_ChildEvent
func miqt_exec_callback_QFormLayout_ChildEvent(self QFormLayout, cb intptr_t, e *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QChildEvent), e *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(e)

	gofunc((&QFormLayout{h: self}).callVirtualBase_ChildEvent, slotval1)

}

type QFormLayout__TakeRowResult struct {
	h          uintptr
	isSubclass bool
}

// NewQFormLayout__TakeRowResult constructs a new QFormLayout::TakeRowResult object.
func NewQFormLayout__TakeRowResult() *QFormLayout__TakeRowResult {

	ret := newQFormLayout__TakeRowResult(QFormLayout__TakeRowResult_new())
	ret.isSubclass = true
	return ret
}

// NewQFormLayout__TakeRowResult2 constructs a new QFormLayout::TakeRowResult object.
func NewQFormLayout__TakeRowResult2(param1 *TakeRowResult) *QFormLayout__TakeRowResult {

	ret := newQFormLayout__TakeRowResult(QFormLayout__TakeRowResult_new2(param1))
	ret.isSubclass = true
	return ret
}
