package qt

import (
	"unsafe"
)

type QGraphicsLinearLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsLinearLayout constructs a new QGraphicsLinearLayout object.
func NewQGraphicsLinearLayout() *QGraphicsLinearLayout {
	g := newQGraphicsLinearLayout(QGraphicsLinearLayout_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsLinearLayout2 constructs a new QGraphicsLinearLayout object.
func NewQGraphicsLinearLayout2(orientation Orientation) *QGraphicsLinearLayout {
	g := newQGraphicsLinearLayout(QGraphicsLinearLayout_new2((int)(orientation)))
	g.isSubclass = true
	return g
}

// NewQGraphicsLinearLayout3 constructs a new QGraphicsLinearLayout object.
func NewQGraphicsLinearLayout3(parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {
	g := newQGraphicsLinearLayout(QGraphicsLinearLayout_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsLinearLayout4 constructs a new QGraphicsLinearLayout object.
func NewQGraphicsLinearLayout4(orientation Orientation, parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {
	g := newQGraphicsLinearLayout(QGraphicsLinearLayout_new4((int)(orientation), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsLinearLayout) SetOrientation(orientation Orientation) {
	QGraphicsLinearLayout_SetOrientation(this.h, (int)(orientation))
}

func (this *QGraphicsLinearLayout) Orientation() Orientation {
	return (Orientation)(QGraphicsLinearLayout_Orientation(this.h))
}

func (this *QGraphicsLinearLayout) AddItem(item *QGraphicsLayoutItem) {
	QGraphicsLinearLayout_AddItem(this.h, item.cPointer())
}

func (this *QGraphicsLinearLayout) AddStretch() {
	QGraphicsLinearLayout_AddStretch(this.h)
}

func (this *QGraphicsLinearLayout) InsertItem(index int, item *QGraphicsLayoutItem) {
	QGraphicsLinearLayout_InsertItem(this.h, (int)(index), item.cPointer())
}

func (this *QGraphicsLinearLayout) InsertStretch(index int) {
	QGraphicsLinearLayout_InsertStretch(this.h, (int)(index))
}

func (this *QGraphicsLinearLayout) RemoveItem(item *QGraphicsLayoutItem) {
	QGraphicsLinearLayout_RemoveItem(this.h, item.cPointer())
}

func (this *QGraphicsLinearLayout) RemoveAt(index int) {
	QGraphicsLinearLayout_RemoveAt(this.h, (int)(index))
}

func (this *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	QGraphicsLinearLayout_SetSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsLinearLayout) Spacing() float64 {
	return (float64)(QGraphicsLinearLayout_Spacing(this.h))
}

func (this *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	QGraphicsLinearLayout_SetItemSpacing(this.h, (int)(index), (double)(spacing))
}

func (this *QGraphicsLinearLayout) ItemSpacing(index int) float64 {
	return (float64)(QGraphicsLinearLayout_ItemSpacing(this.h, (int)(index)))
}

func (this *QGraphicsLinearLayout) SetStretchFactor(item *QGraphicsLayoutItem, stretch int) {
	QGraphicsLinearLayout_SetStretchFactor(this.h, item.cPointer(), (int)(stretch))
}

func (this *QGraphicsLinearLayout) StretchFactor(item *QGraphicsLayoutItem) int {
	return (int)(QGraphicsLinearLayout_StretchFactor(this.h, item.cPointer()))
}

func (this *QGraphicsLinearLayout) SetAlignment(item *QGraphicsLayoutItem, alignment AlignmentFlag) {
	QGraphicsLinearLayout_SetAlignment(this.h, item.cPointer(), (int)(alignment))
}

func (this *QGraphicsLinearLayout) Alignment(item *QGraphicsLayoutItem) AlignmentFlag {
	return (AlignmentFlag)(QGraphicsLinearLayout_Alignment(this.h, item.cPointer()))
}

func (this *QGraphicsLinearLayout) SetGeometry(rect *QRectF) {
	QGraphicsLinearLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QGraphicsLinearLayout) Count() int {
	return (int)(QGraphicsLinearLayout_Count(this.h))
}

func (this *QGraphicsLinearLayout) ItemAt(index int) *QGraphicsLayoutItem {
	return newQGraphicsLayoutItem(QGraphicsLinearLayout_ItemAt(this.h, (int)(index)))
}

func (this *QGraphicsLinearLayout) Invalidate() {
	QGraphicsLinearLayout_Invalidate(this.h)
}

func (this *QGraphicsLinearLayout) SizeHint(which SizeHint, constraint *QSizeF) *QSizeF {
	_goptr := newQSizeF(QGraphicsLinearLayout_SizeHint(this.h, (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLinearLayout) Dump() {
	QGraphicsLinearLayout_Dump(this.h)
}

func (this *QGraphicsLinearLayout) AddStretch1(stretch int) {
	QGraphicsLinearLayout_AddStretch1(this.h, (int)(stretch))
}

func (this *QGraphicsLinearLayout) InsertStretch2(index int, stretch int) {
	QGraphicsLinearLayout_InsertStretch2(this.h, (int)(index), (int)(stretch))
}

func (this *QGraphicsLinearLayout) Dump1(indent int) {
	QGraphicsLinearLayout_Dump1(this.h, (int)(indent))
}
