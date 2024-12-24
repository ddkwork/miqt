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
