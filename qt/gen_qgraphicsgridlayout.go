package qt

import (
	"unsafe"
)

type QGraphicsGridLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsGridLayout constructs a new QGraphicsGridLayout object.
func NewQGraphicsGridLayout() *QGraphicsGridLayout {

	ret := newQGraphicsGridLayout(QGraphicsGridLayout_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsGridLayout2 constructs a new QGraphicsGridLayout object.
func NewQGraphicsGridLayout2(parent *QGraphicsLayoutItem) *QGraphicsGridLayout {

	ret := newQGraphicsGridLayout(QGraphicsGridLayout_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsGridLayout) AddItem(item *QGraphicsLayoutItem, row int, column int, rowSpan int, columnSpan int) {
	QGraphicsGridLayout_AddItem(this.h, item.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan))
}

func (this *QGraphicsGridLayout) AddItem2(item *QGraphicsLayoutItem, row int, column int) {
	QGraphicsGridLayout_AddItem2(this.h, item.cPointer(), (int)(row), (int)(column))
}

func (this *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	QGraphicsGridLayout_SetHorizontalSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsGridLayout) HorizontalSpacing() float64 {
	return (float64)(QGraphicsGridLayout_HorizontalSpacing(this.h))
}

func (this *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	QGraphicsGridLayout_SetVerticalSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsGridLayout) VerticalSpacing() float64 {
	return (float64)(QGraphicsGridLayout_VerticalSpacing(this.h))
}

func (this *QGraphicsGridLayout) SetSpacing(spacing float64) {
	QGraphicsGridLayout_SetSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	QGraphicsGridLayout_SetRowSpacing(this.h, (int)(row), (double)(spacing))
}

func (this *QGraphicsGridLayout) RowSpacing(row int) float64 {
	return (float64)(QGraphicsGridLayout_RowSpacing(this.h, (int)(row)))
}

func (this *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	QGraphicsGridLayout_SetColumnSpacing(this.h, (int)(column), (double)(spacing))
}

func (this *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	return (float64)(QGraphicsGridLayout_ColumnSpacing(this.h, (int)(column)))
}

func (this *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	QGraphicsGridLayout_SetRowStretchFactor(this.h, (int)(row), (int)(stretch))
}

func (this *QGraphicsGridLayout) RowStretchFactor(row int) int {
	return (int)(QGraphicsGridLayout_RowStretchFactor(this.h, (int)(row)))
}

func (this *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	QGraphicsGridLayout_SetColumnStretchFactor(this.h, (int)(column), (int)(stretch))
}

func (this *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	return (int)(QGraphicsGridLayout_ColumnStretchFactor(this.h, (int)(column)))
}

func (this *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	QGraphicsGridLayout_SetRowMinimumHeight(this.h, (int)(row), (double)(height))
}

func (this *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	return (float64)(QGraphicsGridLayout_RowMinimumHeight(this.h, (int)(row)))
}

func (this *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	QGraphicsGridLayout_SetRowPreferredHeight(this.h, (int)(row), (double)(height))
}

func (this *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	return (float64)(QGraphicsGridLayout_RowPreferredHeight(this.h, (int)(row)))
}

func (this *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	QGraphicsGridLayout_SetRowMaximumHeight(this.h, (int)(row), (double)(height))
}

func (this *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	return (float64)(QGraphicsGridLayout_RowMaximumHeight(this.h, (int)(row)))
}

func (this *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	QGraphicsGridLayout_SetRowFixedHeight(this.h, (int)(row), (double)(height))
}

func (this *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	QGraphicsGridLayout_SetColumnMinimumWidth(this.h, (int)(column), (double)(width))
}

func (this *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	return (float64)(QGraphicsGridLayout_ColumnMinimumWidth(this.h, (int)(column)))
}

func (this *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	QGraphicsGridLayout_SetColumnPreferredWidth(this.h, (int)(column), (double)(width))
}

func (this *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	return (float64)(QGraphicsGridLayout_ColumnPreferredWidth(this.h, (int)(column)))
}

func (this *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	QGraphicsGridLayout_SetColumnMaximumWidth(this.h, (int)(column), (double)(width))
}

func (this *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	return (float64)(QGraphicsGridLayout_ColumnMaximumWidth(this.h, (int)(column)))
}

func (this *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	QGraphicsGridLayout_SetColumnFixedWidth(this.h, (int)(column), (double)(width))
}

func (this *QGraphicsGridLayout) SetRowAlignment(row int, alignment AlignmentFlag) {
	QGraphicsGridLayout_SetRowAlignment(this.h, (int)(row), (int)(alignment))
}

func (this *QGraphicsGridLayout) RowAlignment(row int) AlignmentFlag {
	return (AlignmentFlag)(QGraphicsGridLayout_RowAlignment(this.h, (int)(row)))
}

func (this *QGraphicsGridLayout) SetColumnAlignment(column int, alignment AlignmentFlag) {
	QGraphicsGridLayout_SetColumnAlignment(this.h, (int)(column), (int)(alignment))
}

func (this *QGraphicsGridLayout) ColumnAlignment(column int) AlignmentFlag {
	return (AlignmentFlag)(QGraphicsGridLayout_ColumnAlignment(this.h, (int)(column)))
}

func (this *QGraphicsGridLayout) SetAlignment(item *QGraphicsLayoutItem, alignment AlignmentFlag) {
	QGraphicsGridLayout_SetAlignment(this.h, item.cPointer(), (int)(alignment))
}

func (this *QGraphicsGridLayout) Alignment(item *QGraphicsLayoutItem) AlignmentFlag {
	return (AlignmentFlag)(QGraphicsGridLayout_Alignment(this.h, item.cPointer()))
}

func (this *QGraphicsGridLayout) RowCount() int {
	return (int)(QGraphicsGridLayout_RowCount(this.h))
}

func (this *QGraphicsGridLayout) ColumnCount() int {
	return (int)(QGraphicsGridLayout_ColumnCount(this.h))
}

func (this *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem {
	return newQGraphicsLayoutItem(QGraphicsGridLayout_ItemAt(this.h, (int)(row), (int)(column)))
}

func (this *QGraphicsGridLayout) Count() int {
	return (int)(QGraphicsGridLayout_Count(this.h))
}

func (this *QGraphicsGridLayout) ItemAtWithIndex(index int) *QGraphicsLayoutItem {
	return newQGraphicsLayoutItem(QGraphicsGridLayout_ItemAtWithIndex(this.h, (int)(index)))
}

func (this *QGraphicsGridLayout) RemoveAt(index int) {
	QGraphicsGridLayout_RemoveAt(this.h, (int)(index))
}

func (this *QGraphicsGridLayout) RemoveItem(item *QGraphicsLayoutItem) {
	QGraphicsGridLayout_RemoveItem(this.h, item.cPointer())
}

func (this *QGraphicsGridLayout) Invalidate() {
	QGraphicsGridLayout_Invalidate(this.h)
}

func (this *QGraphicsGridLayout) SetGeometry(rect *QRectF) {
	QGraphicsGridLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QGraphicsGridLayout) SizeHint(which SizeHint, constraint *QSizeF) *QSizeF {
	_goptr := newQSizeF(QGraphicsGridLayout_SizeHint(this.h, (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsGridLayout) AddItem6(item *QGraphicsLayoutItem, row int, column int, rowSpan int, columnSpan int, alignment AlignmentFlag) {
	QGraphicsGridLayout_AddItem6(this.h, item.cPointer(), (int)(row), (int)(column), (int)(rowSpan), (int)(columnSpan), (int)(alignment))
}

func (this *QGraphicsGridLayout) AddItem4(item *QGraphicsLayoutItem, row int, column int, alignment AlignmentFlag) {
	QGraphicsGridLayout_AddItem4(this.h, item.cPointer(), (int)(row), (int)(column), (int)(alignment))
}

func (this *QGraphicsGridLayout) callVirtualBase_Count() int {

	return (int)(QGraphicsGridLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QGraphicsGridLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_Count
func miqt_exec_callback_QGraphicsGridLayout_Count(self QGraphicsGridLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QGraphicsGridLayout) callVirtualBase_ItemAtWithIndex(index int) *QGraphicsLayoutItem {

	return newQGraphicsLayoutItem(QGraphicsGridLayout_virtualbase_ItemAtWithIndex(unsafe.Pointer(this.h), (int)(index)))

}
func (this *QGraphicsGridLayout) OnItemAtWithIndex(slot func(super func(index int) *QGraphicsLayoutItem, index int) *QGraphicsLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_ItemAtWithIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_ItemAtWithIndex
func miqt_exec_callback_QGraphicsGridLayout_ItemAtWithIndex(self QGraphicsGridLayout, cb intptr_t, index int) *QGraphicsLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QGraphicsLayoutItem, index int) *QGraphicsLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_ItemAtWithIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsGridLayout) callVirtualBase_RemoveAt(index int) {

	QGraphicsGridLayout_virtualbase_RemoveAt(unsafe.Pointer(this.h), (int)(index))

}
func (this *QGraphicsGridLayout) OnRemoveAt(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_RemoveAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_RemoveAt
func miqt_exec_callback_QGraphicsGridLayout_RemoveAt(self QGraphicsGridLayout, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_RemoveAt, slotval1)

}

func (this *QGraphicsGridLayout) callVirtualBase_Invalidate() {

	QGraphicsGridLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QGraphicsGridLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_Invalidate
func miqt_exec_callback_QGraphicsGridLayout_Invalidate(self QGraphicsGridLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QGraphicsGridLayout) callVirtualBase_SetGeometry(rect *QRectF) {

	QGraphicsGridLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QGraphicsGridLayout) OnSetGeometry(slot func(super func(rect *QRectF), rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_SetGeometry
func miqt_exec_callback_QGraphicsGridLayout_SetGeometry(self QGraphicsGridLayout, cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF), rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGraphicsGridLayout) callVirtualBase_SizeHint(which SizeHint, constraint *QSizeF) *QSizeF {

	_goptr := newQSizeF(QGraphicsGridLayout_virtualbase_SizeHint(unsafe.Pointer(this.h), (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsGridLayout) OnSizeHint(slot func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_SizeHint
func miqt_exec_callback_QGraphicsGridLayout_SizeHint(self QGraphicsGridLayout, cb intptr_t, which int, constraint *QSizeF) *QSizeF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (SizeHint)(which)

	slotval2 := newQSizeF(constraint)

	virtualReturn := gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_SizeHint, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsGridLayout) callVirtualBase_GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {

	QGraphicsGridLayout_virtualbase_GetContentsMargins(unsafe.Pointer(this.h), (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))

}
func (this *QGraphicsGridLayout) OnGetContentsMargins(slot func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_GetContentsMargins(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_GetContentsMargins
func miqt_exec_callback_QGraphicsGridLayout_GetContentsMargins(self QGraphicsGridLayout, cb intptr_t, left *double, top *double, right *double, bottom *double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*float64)(unsafe.Pointer(left))

	slotval2 := (*float64)(unsafe.Pointer(top))

	slotval3 := (*float64)(unsafe.Pointer(right))

	slotval4 := (*float64)(unsafe.Pointer(bottom))

	gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_GetContentsMargins, slotval1, slotval2, slotval3, slotval4)

}

func (this *QGraphicsGridLayout) callVirtualBase_UpdateGeometry() {

	QGraphicsGridLayout_virtualbase_UpdateGeometry(unsafe.Pointer(this.h))

}
func (this *QGraphicsGridLayout) OnUpdateGeometry(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_UpdateGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_UpdateGeometry
func miqt_exec_callback_QGraphicsGridLayout_UpdateGeometry(self QGraphicsGridLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_UpdateGeometry)

}

func (this *QGraphicsGridLayout) callVirtualBase_WidgetEvent(e *QEvent) {

	QGraphicsGridLayout_virtualbase_WidgetEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QGraphicsGridLayout) OnWidgetEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsGridLayout_override_virtual_WidgetEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsGridLayout_WidgetEvent
func miqt_exec_callback_QGraphicsGridLayout_WidgetEvent(self QGraphicsGridLayout, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QGraphicsGridLayout{h: self}).callVirtualBase_WidgetEvent, slotval1)

}
