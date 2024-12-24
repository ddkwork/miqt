package qt

import (
	"unsafe"
)

type QGraphicsLayoutItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsLayoutItem constructs a new QGraphicsLayoutItem object.
func NewQGraphicsLayoutItem() *QGraphicsLayoutItem {
	g := newQGraphicsLayoutItem(QGraphicsLayoutItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsLayoutItem2 constructs a new QGraphicsLayoutItem object.
func NewQGraphicsLayoutItem2(parent *QGraphicsLayoutItem) *QGraphicsLayoutItem {
	g := newQGraphicsLayoutItem(QGraphicsLayoutItem_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsLayoutItem3 constructs a new QGraphicsLayoutItem object.
func NewQGraphicsLayoutItem3(parent *QGraphicsLayoutItem, isLayout bool) *QGraphicsLayoutItem {
	g := newQGraphicsLayoutItem(QGraphicsLayoutItem_new3(parent.cPointer(), (bool)(isLayout)))
	g.isSubclass = true
	return g
}

func (this *QGraphicsLayoutItem) SetSizePolicy(policy *QSizePolicy) {
	QGraphicsLayoutItem_SetSizePolicy(this.h, policy.cPointer())
}

func (this *QGraphicsLayoutItem) SetSizePolicy2(hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) {
	QGraphicsLayoutItem_SetSizePolicy2(this.h, (int)(hPolicy), (int)(vPolicy))
}

func (this *QGraphicsLayoutItem) SizePolicy() *QSizePolicy {
	_goptr := newQSizePolicy(QGraphicsLayoutItem_SizePolicy(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLayoutItem) SetMinimumSize(size *QSizeF) {
	QGraphicsLayoutItem_SetMinimumSize(this.h, size.cPointer())
}

func (this *QGraphicsLayoutItem) SetMinimumSize2(w float64, h float64) {
	QGraphicsLayoutItem_SetMinimumSize2(this.h, (double)(w), (double)(h))
}

func (this *QGraphicsLayoutItem) MinimumSize() *QSizeF {
	_goptr := newQSizeF(QGraphicsLayoutItem_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLayoutItem) SetMinimumWidth(width float64) {
	QGraphicsLayoutItem_SetMinimumWidth(this.h, (double)(width))
}

func (this *QGraphicsLayoutItem) MinimumWidth() float64 {
	return (float64)(QGraphicsLayoutItem_MinimumWidth(this.h))
}

func (this *QGraphicsLayoutItem) SetMinimumHeight(height float64) {
	QGraphicsLayoutItem_SetMinimumHeight(this.h, (double)(height))
}

func (this *QGraphicsLayoutItem) MinimumHeight() float64 {
	return (float64)(QGraphicsLayoutItem_MinimumHeight(this.h))
}

func (this *QGraphicsLayoutItem) SetPreferredSize(size *QSizeF) {
	QGraphicsLayoutItem_SetPreferredSize(this.h, size.cPointer())
}

func (this *QGraphicsLayoutItem) SetPreferredSize2(w float64, h float64) {
	QGraphicsLayoutItem_SetPreferredSize2(this.h, (double)(w), (double)(h))
}

func (this *QGraphicsLayoutItem) PreferredSize() *QSizeF {
	_goptr := newQSizeF(QGraphicsLayoutItem_PreferredSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLayoutItem) SetPreferredWidth(width float64) {
	QGraphicsLayoutItem_SetPreferredWidth(this.h, (double)(width))
}

func (this *QGraphicsLayoutItem) PreferredWidth() float64 {
	return (float64)(QGraphicsLayoutItem_PreferredWidth(this.h))
}

func (this *QGraphicsLayoutItem) SetPreferredHeight(height float64) {
	QGraphicsLayoutItem_SetPreferredHeight(this.h, (double)(height))
}

func (this *QGraphicsLayoutItem) PreferredHeight() float64 {
	return (float64)(QGraphicsLayoutItem_PreferredHeight(this.h))
}

func (this *QGraphicsLayoutItem) SetMaximumSize(size *QSizeF) {
	QGraphicsLayoutItem_SetMaximumSize(this.h, size.cPointer())
}

func (this *QGraphicsLayoutItem) SetMaximumSize2(w float64, h float64) {
	QGraphicsLayoutItem_SetMaximumSize2(this.h, (double)(w), (double)(h))
}

func (this *QGraphicsLayoutItem) MaximumSize() *QSizeF {
	_goptr := newQSizeF(QGraphicsLayoutItem_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLayoutItem) SetMaximumWidth(width float64) {
	QGraphicsLayoutItem_SetMaximumWidth(this.h, (double)(width))
}

func (this *QGraphicsLayoutItem) MaximumWidth() float64 {
	return (float64)(QGraphicsLayoutItem_MaximumWidth(this.h))
}

func (this *QGraphicsLayoutItem) SetMaximumHeight(height float64) {
	QGraphicsLayoutItem_SetMaximumHeight(this.h, (double)(height))
}

func (this *QGraphicsLayoutItem) MaximumHeight() float64 {
	return (float64)(QGraphicsLayoutItem_MaximumHeight(this.h))
}

func (this *QGraphicsLayoutItem) SetGeometry(rect *QRectF) {
	QGraphicsLayoutItem_SetGeometry(this.h, rect.cPointer())
}

func (this *QGraphicsLayoutItem) Geometry() *QRectF {
	_goptr := newQRectF(QGraphicsLayoutItem_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLayoutItem) GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {
	QGraphicsLayoutItem_GetContentsMargins(this.h, (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))
}

func (this *QGraphicsLayoutItem) ContentsRect() *QRectF {
	_goptr := newQRectF(QGraphicsLayoutItem_ContentsRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLayoutItem) EffectiveSizeHint(which SizeHint) *QSizeF {
	_goptr := newQSizeF(QGraphicsLayoutItem_EffectiveSizeHint(this.h, (int)(which)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLayoutItem) UpdateGeometry() {
	QGraphicsLayoutItem_UpdateGeometry(this.h)
}

func (this *QGraphicsLayoutItem) IsEmpty() bool {
	return (bool)(QGraphicsLayoutItem_IsEmpty(this.h))
}

func (this *QGraphicsLayoutItem) ParentLayoutItem() *QGraphicsLayoutItem {
	return newQGraphicsLayoutItem(QGraphicsLayoutItem_ParentLayoutItem(this.h))
}

func (this *QGraphicsLayoutItem) SetParentLayoutItem(parent *QGraphicsLayoutItem) {
	QGraphicsLayoutItem_SetParentLayoutItem(this.h, parent.cPointer())
}

func (this *QGraphicsLayoutItem) IsLayout() bool {
	return (bool)(QGraphicsLayoutItem_IsLayout(this.h))
}

func (this *QGraphicsLayoutItem) GraphicsItem() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsLayoutItem_GraphicsItem(this.h))
}

func (this *QGraphicsLayoutItem) OwnedByLayout() bool {
	return (bool)(QGraphicsLayoutItem_OwnedByLayout(this.h))
}

func (this *QGraphicsLayoutItem) SetSizePolicy3(hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy, controlType QSizePolicy__ControlType) {
	QGraphicsLayoutItem_SetSizePolicy3(this.h, (int)(hPolicy), (int)(vPolicy), (int)(controlType))
}

func (this *QGraphicsLayoutItem) EffectiveSizeHint2(which SizeHint, constraint *QSizeF) *QSizeF {
	_goptr := newQSizeF(QGraphicsLayoutItem_EffectiveSizeHint2(this.h, (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
