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
	g := newQFormLayout(QFormLayout_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQFormLayout2 constructs a new QFormLayout object.
func NewQFormLayout2() *QFormLayout {
	g := newQFormLayout(QFormLayout_new2())
	g.isSubclass = true
	return g
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

func (this *QFormLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFormLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFormLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_MetaObject
func miqt_exec_callback_QFormLayout_MetaObject(self QFormLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFormLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFormLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFormLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFormLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFormLayout_Metacast
func miqt_exec_callback_QFormLayout_Metacast(self QFormLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QFormLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QFormLayout__TakeRowResult struct {
	h          uintptr
	isSubclass bool
}

// NewQFormLayout__TakeRowResult constructs a new QFormLayout::TakeRowResult object.
func NewQFormLayout__TakeRowResult() *QFormLayout__TakeRowResult {
	g := newQFormLayout__TakeRowResult(QFormLayout__TakeRowResult_new())
	g.isSubclass = true
	return g
}

// NewQFormLayout__TakeRowResult2 constructs a new QFormLayout::TakeRowResult object.
func NewQFormLayout__TakeRowResult2(param1 *TakeRowResult) *QFormLayout__TakeRowResult {
	g := newQFormLayout__TakeRowResult(QFormLayout__TakeRowResult_new2(param1))
	g.isSubclass = true
	return g
}
