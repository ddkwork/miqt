package qt

import (
	"unsafe"
)

type QGridLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQGridLayout constructs a new QGridLayout object.
func NewQGridLayout(parent *QWidget) *QGridLayout {
	g := newQGridLayout(QGridLayout_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGridLayout2 constructs a new QGridLayout object.
func NewQGridLayout2() *QGridLayout {
	g := newQGridLayout(QGridLayout_new2())
	g.isSubclass = true
	return g
}

func (this *QGridLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QGridLayout_MetaObject(this.h))
}

func (this *QGridLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGridLayout_Metacast(this.h, param1_Cstring))
}

func QGridLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGridLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGridLayout) SizeHint() *QSize {
	_goptr := newQSize(QGridLayout_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGridLayout) MinimumSize() *QSize {
	_goptr := newQSize(QGridLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGridLayout) MaximumSize() *QSize {
	_goptr := newQSize(QGridLayout_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGridLayout) SetHorizontalSpacing(spacing int) {
	QGridLayout_SetHorizontalSpacing(this.h, (int)(spacing))
}

func (this *QGridLayout) HorizontalSpacing() int {
	return (int)(QGridLayout_HorizontalSpacing(this.h))
}

func (this *QGridLayout) SetVerticalSpacing(spacing int) {
	QGridLayout_SetVerticalSpacing(this.h, (int)(spacing))
}

func (this *QGridLayout) VerticalSpacing() int {
	return (int)(QGridLayout_VerticalSpacing(this.h))
}

func (this *QGridLayout) SetSpacing(spacing int) {
	QGridLayout_SetSpacing(this.h, (int)(spacing))
}

func (this *QGridLayout) Spacing() int {
	return (int)(QGridLayout_Spacing(this.h))
}

func (this *QGridLayout) SetRowStretch(row int, stretch int) {
	QGridLayout_SetRowStretch(this.h, (int)(row), (int)(stretch))
}

func (this *QGridLayout) SetColumnStretch(column int, stretch int) {
	QGridLayout_SetColumnStretch(this.h, (int)(column), (int)(stretch))
}

func (this *QGridLayout) RowStretch(row int) int {
	return (int)(QGridLayout_RowStretch(this.h, (int)(row)))
}

func (this *QGridLayout) ColumnStretch(column int) int {
	return (int)(QGridLayout_ColumnStretch(this.h, (int)(column)))
}

func (this *QGridLayout) SetRowMinimumHeight(row int, minSize int) {
	QGridLayout_SetRowMinimumHeight(this.h, (int)(row), (int)(minSize))
}

func (this *QGridLayout) SetColumnMinimumWidth(column int, minSize int) {
	QGridLayout_SetColumnMinimumWidth(this.h, (int)(column), (int)(minSize))
}

func (this *QGridLayout) RowMinimumHeight(row int) int {
	return (int)(QGridLayout_RowMinimumHeight(this.h, (int)(row)))
}

func (this *QGridLayout) ColumnMinimumWidth(column int) int {
	return (int)(QGridLayout_ColumnMinimumWidth(this.h, (int)(column)))
}

func (this *QGridLayout) ColumnCount() int {
	return (int)(QGridLayout_ColumnCount(this.h))
}

func (this *QGridLayout) RowCount() int {
	return (int)(QGridLayout_RowCount(this.h))
}

func (this *QGridLayout) CellRect(row int, column int) *QRect {
	_goptr := newQRect(QGridLayout_CellRect(this.h, (int)(row), (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGridLayout) HasHeightForWidth() bool {
	return (bool)(QGridLayout_HasHeightForWidth(this.h))
}

func (this *QGridLayout) HeightForWidth(param1 int) int {
	return (int)(QGridLayout_HeightForWidth(this.h, (int)(param1)))
}

func (this *QGridLayout) MinimumHeightForWidth(param1 int) int {
	return (int)(QGridLayout_MinimumHeightForWidth(this.h, (int)(param1)))
}

func (this *QGridLayout) ExpandingDirections() Orientation {
	return (Orientation)(QGridLayout_ExpandingDirections(this.h))
}

func (this *QGridLayout) Invalidate() {
	QGridLayout_Invalidate(this.h)
}

func (this *QGridLayout) AddWidget(w *QWidget) {
	QGridLayout_AddWidget(this.h, w.cPointer())
}

func (this *QGridLayout) AddWidget2(param1 *QWidget, row int, column int) {
	QGridLayout_AddWidget2(this.h, param1.cPointer(), (int)(row), (int)(column))
}

func (this *QGridLayout) AddWidget3(param1 *QWidget, row int, column int, rowSpan int, columnSpan int) {
	QGridLayout_AddWidget3(this.h, param1.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan))
}

func (this *QGridLayout) AddLayout(param1 *QLayout, row int, column int) {
	QGridLayout_AddLayout(this.h, param1.cPointer(), (int)(row), (int)(column))
}

func (this *QGridLayout) AddLayout2(param1 *QLayout, row int, column int, rowSpan int, columnSpan int) {
	QGridLayout_AddLayout2(this.h, param1.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan))
}

func (this *QGridLayout) SetOriginCorner(originCorner Corner) {
	QGridLayout_SetOriginCorner(this.h, (int)(originCorner))
}

func (this *QGridLayout) OriginCorner() Corner {
	return (Corner)(QGridLayout_OriginCorner(this.h))
}

func (this *QGridLayout) ItemAt(index int) *QLayoutItem {
	return newQLayoutItem(QGridLayout_ItemAt(this.h, (int)(index)))
}

func (this *QGridLayout) ItemAtPosition(row int, column int) *QLayoutItem {
	return newQLayoutItem(QGridLayout_ItemAtPosition(this.h, (int)(row), (int)(column)))
}

func (this *QGridLayout) TakeAt(index int) *QLayoutItem {
	return newQLayoutItem(QGridLayout_TakeAt(this.h, (int)(index)))
}

func (this *QGridLayout) Count() int {
	return (int)(QGridLayout_Count(this.h))
}

func (this *QGridLayout) SetGeometry(geometry *QRect) {
	QGridLayout_SetGeometry(this.h, geometry.cPointer())
}

func (this *QGridLayout) AddItem(item *QLayoutItem, row int, column int) {
	QGridLayout_AddItem(this.h, item.cPointer(), (int)(row), (int)(column))
}

func (this *QGridLayout) SetDefaultPositioning(n int, orient Orientation) {
	QGridLayout_SetDefaultPositioning(this.h, (int)(n), (int)(orient))
}

func (this *QGridLayout) GetItemPosition(idx int, row *int, column *int, rowSpan *int, columnSpan *int) {
	QGridLayout_GetItemPosition(this.h, (int)(idx), (*int)(unsafe.Pointer(row)), (*int)(unsafe.Pointer(column)), (*int)(unsafe.Pointer(rowSpan)), (*int)(unsafe.Pointer(columnSpan)))
}

func QGridLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGridLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGridLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGridLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGridLayout) AddWidget4(param1 *QWidget, row int, column int, param4 AlignmentFlag) {
	QGridLayout_AddWidget4(this.h, param1.cPointer(), (int)(row), (int)(column), (int)(param4))
}

func (this *QGridLayout) AddWidget6(param1 *QWidget, row int, column int, rowSpan int, columnSpan int, param6 AlignmentFlag) {
	QGridLayout_AddWidget6(this.h, param1.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan), (int)(param6))
}

func (this *QGridLayout) AddLayout4(param1 *QLayout, row int, column int, param4 AlignmentFlag) {
	QGridLayout_AddLayout4(this.h, param1.cPointer(), (int)(row), (int)(column), (int)(param4))
}

func (this *QGridLayout) AddLayout6(param1 *QLayout, row int, column int, rowSpan int, columnSpan int, param6 AlignmentFlag) {
	QGridLayout_AddLayout6(this.h, param1.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan), (int)(param6))
}

func (this *QGridLayout) AddItem4(item *QLayoutItem, row int, column int, rowSpan int) {
	QGridLayout_AddItem4(this.h, item.cPointer(), (int)(row), (int)(column), (int)(rowSpan))
}

func (this *QGridLayout) AddItem5(item *QLayoutItem, row int, column int, rowSpan int, columnSpan int) {
	QGridLayout_AddItem5(this.h, item.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan))
}

func (this *QGridLayout) AddItem6(item *QLayoutItem, row int, column int, rowSpan int, columnSpan int, param6 AlignmentFlag) {
	QGridLayout_AddItem6(this.h, item.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan), (int)(param6))
}

func (this *QGridLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGridLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGridLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_MetaObject
func miqt_exec_callback_QGridLayout_MetaObject(self QGridLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGridLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGridLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGridLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_Metacast
func miqt_exec_callback_QGridLayout_Metacast(self QGridLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
