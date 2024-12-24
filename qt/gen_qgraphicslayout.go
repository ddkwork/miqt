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

func (this *QGraphicsLayout) callVirtualBase_GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {

	QGraphicsLayout_virtualbase_GetContentsMargins(unsafe.Pointer(this.h), (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))

}
func (this *QGraphicsLayout) OnGetContentsMargins(slot func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_GetContentsMargins(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_GetContentsMargins
func miqt_exec_callback_QGraphicsLayout_GetContentsMargins(self QGraphicsLayout, cb intptr_t, left *double, top *double, right *double, bottom *double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*float64)(unsafe.Pointer(left))

	slotval2 := (*float64)(unsafe.Pointer(top))

	slotval3 := (*float64)(unsafe.Pointer(right))

	slotval4 := (*float64)(unsafe.Pointer(bottom))

	gofunc((&QGraphicsLayout{h: self}).callVirtualBase_GetContentsMargins, slotval1, slotval2, slotval3, slotval4)

}

func (this *QGraphicsLayout) callVirtualBase_Invalidate() {

	QGraphicsLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QGraphicsLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_Invalidate
func miqt_exec_callback_QGraphicsLayout_Invalidate(self QGraphicsLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QGraphicsLayout) callVirtualBase_UpdateGeometry() {

	QGraphicsLayout_virtualbase_UpdateGeometry(unsafe.Pointer(this.h))

}
func (this *QGraphicsLayout) OnUpdateGeometry(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_UpdateGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_UpdateGeometry
func miqt_exec_callback_QGraphicsLayout_UpdateGeometry(self QGraphicsLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsLayout{h: self}).callVirtualBase_UpdateGeometry)

}

func (this *QGraphicsLayout) callVirtualBase_WidgetEvent(e *QEvent) {

	QGraphicsLayout_virtualbase_WidgetEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QGraphicsLayout) OnWidgetEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_WidgetEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_WidgetEvent
func miqt_exec_callback_QGraphicsLayout_WidgetEvent(self QGraphicsLayout, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QGraphicsLayout{h: self}).callVirtualBase_WidgetEvent, slotval1)

}
func (this *QGraphicsLayout) OnCount(slot func() int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_Count
func miqt_exec_callback_QGraphicsLayout_Count(self QGraphicsLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func() int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (int)(virtualReturn)

}
func (this *QGraphicsLayout) OnItemAt(slot func(i int) *QGraphicsLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_ItemAt
func miqt_exec_callback_QGraphicsLayout_ItemAt(self QGraphicsLayout, cb intptr_t, i int) *QGraphicsLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(i int) *QGraphicsLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(i)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}
func (this *QGraphicsLayout) OnRemoveAt(slot func(index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_RemoveAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_RemoveAt
func miqt_exec_callback_QGraphicsLayout_RemoveAt(self QGraphicsLayout, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)

}

func (this *QGraphicsLayout) callVirtualBase_SetGeometry(rect *QRectF) {

	QGraphicsLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QGraphicsLayout) OnSetGeometry(slot func(super func(rect *QRectF), rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_SetGeometry
func miqt_exec_callback_QGraphicsLayout_SetGeometry(self QGraphicsLayout, cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF), rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc((&QGraphicsLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGraphicsLayout) callVirtualBase_IsEmpty() bool {

	return (bool)(QGraphicsLayout_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QGraphicsLayout) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_IsEmpty
func miqt_exec_callback_QGraphicsLayout_IsEmpty(self QGraphicsLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsLayout{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}
func (this *QGraphicsLayout) OnSizeHint(slot func(which SizeHint, constraint *QSizeF) *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayout_SizeHint
func miqt_exec_callback_QGraphicsLayout_SizeHint(self QGraphicsLayout, cb intptr_t, which int, constraint *QSizeF) *QSizeF {
	gofunc, ok := cgo.Handle(cb).Value().(func(which SizeHint, constraint *QSizeF) *QSizeF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (SizeHint)(which)

	slotval2 := newQSizeF(constraint)

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}
