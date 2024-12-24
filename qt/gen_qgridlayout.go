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

	ret := newQGridLayout(QGridLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGridLayout2 constructs a new QGridLayout object.
func NewQGridLayout2() *QGridLayout {

	ret := newQGridLayout(QGridLayout_new2())
	ret.isSubclass = true
	return ret
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

func (this *QGridLayout) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QGridLayout_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGridLayout) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_SizeHint
func miqt_exec_callback_QGridLayout_SizeHint(self QGridLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QGridLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGridLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_MinimumSize
func miqt_exec_callback_QGridLayout_MinimumSize(self QGridLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QGridLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGridLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_MaximumSize
func miqt_exec_callback_QGridLayout_MaximumSize(self QGridLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_SetSpacing(spacing int) {

	QGridLayout_virtualbase_SetSpacing(unsafe.Pointer(this.h), (int)(spacing))

}
func (this *QGridLayout) OnSetSpacing(slot func(super func(spacing int), spacing int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_SetSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_SetSpacing
func miqt_exec_callback_QGridLayout_SetSpacing(self QGridLayout, cb intptr_t, spacing int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(spacing int), spacing int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(spacing)

	gofunc((&QGridLayout{h: self}).callVirtualBase_SetSpacing, slotval1)

}

func (this *QGridLayout) callVirtualBase_Spacing() int {

	return (int)(QGridLayout_virtualbase_Spacing(unsafe.Pointer(this.h)))

}
func (this *QGridLayout) OnSpacing(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_Spacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_Spacing
func miqt_exec_callback_QGridLayout_Spacing(self QGridLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_Spacing)

	return (int)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QGridLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QGridLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_HasHeightForWidth
func miqt_exec_callback_QGridLayout_HasHeightForWidth(self QGridLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QGridLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QGridLayout) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_HeightForWidth
func miqt_exec_callback_QGridLayout_HeightForWidth(self QGridLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QGridLayout_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QGridLayout) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_MinimumHeightForWidth
func miqt_exec_callback_QGridLayout_MinimumHeightForWidth(self QGridLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QGridLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QGridLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_ExpandingDirections
func miqt_exec_callback_QGridLayout_ExpandingDirections(self QGridLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_Invalidate() {

	QGridLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QGridLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_Invalidate
func miqt_exec_callback_QGridLayout_Invalidate(self QGridLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGridLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QGridLayout) callVirtualBase_ItemAt(index int) *QLayoutItem {

	return newQLayoutItem(QGridLayout_virtualbase_ItemAt(unsafe.Pointer(this.h), (int)(index)))

}
func (this *QGridLayout) OnItemAt(slot func(super func(index int) *QLayoutItem, index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_ItemAt
func miqt_exec_callback_QGridLayout_ItemAt(self QGridLayout, cb intptr_t, index int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QLayoutItem, index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_ItemAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_TakeAt(index int) *QLayoutItem {

	return newQLayoutItem(QGridLayout_virtualbase_TakeAt(unsafe.Pointer(this.h), (int)(index)))

}
func (this *QGridLayout) OnTakeAt(slot func(super func(index int) *QLayoutItem, index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_TakeAt
func miqt_exec_callback_QGridLayout_TakeAt(self QGridLayout, cb intptr_t, index int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QLayoutItem, index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_TakeAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_Count() int {

	return (int)(QGridLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QGridLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_Count
func miqt_exec_callback_QGridLayout_Count(self QGridLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_SetGeometry(geometry *QRect) {

	QGridLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QGridLayout) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_SetGeometry
func miqt_exec_callback_QGridLayout_SetGeometry(self QGridLayout, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QGridLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGridLayout) callVirtualBase_AddItemWithQLayoutItem(param1 *QLayoutItem) {

	QGridLayout_virtualbase_AddItemWithQLayoutItem(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QGridLayout) OnAddItemWithQLayoutItem(slot func(super func(param1 *QLayoutItem), param1 *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_AddItemWithQLayoutItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_AddItemWithQLayoutItem
func miqt_exec_callback_QGridLayout_AddItemWithQLayoutItem(self QGridLayout, cb intptr_t, param1 *QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QLayoutItem), param1 *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(param1)

	gofunc((&QGridLayout{h: self}).callVirtualBase_AddItemWithQLayoutItem, slotval1)

}

func (this *QGridLayout) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QGridLayout_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGridLayout) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_Geometry
func miqt_exec_callback_QGridLayout_Geometry(self QGridLayout, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_IndexOf(param1 *QWidget) int {

	return (int)(QGridLayout_virtualbase_IndexOf(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QGridLayout) OnIndexOf(slot func(super func(param1 *QWidget) int, param1 *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_IndexOf(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_IndexOf
func miqt_exec_callback_QGridLayout_IndexOf(self QGridLayout, cb intptr_t, param1 *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWidget) int, param1 *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(param1)

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_IndexOf, slotval1)

	return (int)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_IsEmpty() bool {

	return (bool)(QGridLayout_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QGridLayout) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_IsEmpty
func miqt_exec_callback_QGridLayout_IsEmpty(self QGridLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QGridLayout_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QGridLayout) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_ControlTypes
func miqt_exec_callback_QGridLayout_ControlTypes(self QGridLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

func (this *QGridLayout) callVirtualBase_ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {

	return newQLayoutItem(QGridLayout_virtualbase_ReplaceWidget(unsafe.Pointer(this.h), from.cPointer(), to.cPointer(), (int)(options)))

}
func (this *QGridLayout) OnReplaceWidget(slot func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_ReplaceWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_ReplaceWidget
func miqt_exec_callback_QGridLayout_ReplaceWidget(self QGridLayout, cb intptr_t, from *QWidget, to *QWidget, options int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(from)

	slotval2 := newQWidget(to)

	slotval3 := (FindChildOption)(options)

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_ReplaceWidget, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_Layout() *QLayout {

	return newQLayout(QGridLayout_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QGridLayout) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_Layout
func miqt_exec_callback_QGridLayout_Layout(self QGridLayout, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGridLayout{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QGridLayout) callVirtualBase_ChildEvent(e *QChildEvent) {

	QGridLayout_virtualbase_ChildEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QGridLayout) OnChildEvent(slot func(super func(e *QChildEvent), e *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGridLayout_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGridLayout_ChildEvent
func miqt_exec_callback_QGridLayout_ChildEvent(self QGridLayout, cb intptr_t, e *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QChildEvent), e *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(e)

	gofunc((&QGridLayout{h: self}).callVirtualBase_ChildEvent, slotval1)

}
