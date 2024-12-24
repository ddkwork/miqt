package qt

import (
	"unsafe"
)

type QGraphicsLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsLayout constructs a new QGraphicsLayout object.
func NewQGraphicsLayout() *QGraphicsLayout {
	ret := newQGraphicsLayout(QGraphicsLayout_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsLayout2 constructs a new QGraphicsLayout object.
func NewQGraphicsLayout2(parent *QGraphicsLayoutItem) *QGraphicsLayout {
	ret := newQGraphicsLayout(QGraphicsLayout_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	QGraphicsLayout_SetContentsMargins(this.h, (double)(left), (double)(top), (double)(right), (double)(bottom))
}

func (this *QGraphicsLayout) GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {
	QGraphicsLayout_GetContentsMargins(this.h, (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))
}

func (this *QGraphicsLayout) Activate() {
	QGraphicsLayout_Activate(this.h)
}

func (this *QGraphicsLayout) IsActivated() bool {
	return (bool)(QGraphicsLayout_IsActivated(this.h))
}

func (this *QGraphicsLayout) Invalidate() {
	QGraphicsLayout_Invalidate(this.h)
}

func (this *QGraphicsLayout) UpdateGeometry() {
	QGraphicsLayout_UpdateGeometry(this.h)
}

func (this *QGraphicsLayout) WidgetEvent(e *QEvent) {
	QGraphicsLayout_WidgetEvent(this.h, e.cPointer())
}

func (this *QGraphicsLayout) Count() int {
	return (int)(QGraphicsLayout_Count(this.h))
}

func (this *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem {
	return newQGraphicsLayoutItem(QGraphicsLayout_ItemAt(this.h, (int)(i)))
}

func (this *QGraphicsLayout) RemoveAt(index int) {
	QGraphicsLayout_RemoveAt(this.h, (int)(index))
}

func QGraphicsLayout_SetInstantInvalidatePropagation(enable bool) {
	QGraphicsLayout_SetInstantInvalidatePropagation((bool)(enable))
}

func QGraphicsLayout_InstantInvalidatePropagation() bool {
	return (bool)(QGraphicsLayout_InstantInvalidatePropagation())
}
