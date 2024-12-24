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

	ret := newQGraphicsLinearLayout(QGraphicsLinearLayout_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsLinearLayout2 constructs a new QGraphicsLinearLayout object.
func NewQGraphicsLinearLayout2(orientation Orientation) *QGraphicsLinearLayout {

	ret := newQGraphicsLinearLayout(QGraphicsLinearLayout_new2((int)(orientation)))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsLinearLayout3 constructs a new QGraphicsLinearLayout object.
func NewQGraphicsLinearLayout3(parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {

	ret := newQGraphicsLinearLayout(QGraphicsLinearLayout_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsLinearLayout4 constructs a new QGraphicsLinearLayout object.
func NewQGraphicsLinearLayout4(orientation Orientation, parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {

	ret := newQGraphicsLinearLayout(QGraphicsLinearLayout_new4((int)(orientation), parent.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QGraphicsLinearLayout) callVirtualBase_RemoveAt(index int) {

	QGraphicsLinearLayout_virtualbase_RemoveAt(unsafe.Pointer(this.h), (int)(index))

}
func (this *QGraphicsLinearLayout) OnRemoveAt(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_RemoveAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_RemoveAt
func miqt_exec_callback_QGraphicsLinearLayout_RemoveAt(self QGraphicsLinearLayout, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_RemoveAt, slotval1)

}

func (this *QGraphicsLinearLayout) callVirtualBase_SetGeometry(rect *QRectF) {

	QGraphicsLinearLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QGraphicsLinearLayout) OnSetGeometry(slot func(super func(rect *QRectF), rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_SetGeometry
func miqt_exec_callback_QGraphicsLinearLayout_SetGeometry(self QGraphicsLinearLayout, cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF), rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGraphicsLinearLayout) callVirtualBase_Count() int {

	return (int)(QGraphicsLinearLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QGraphicsLinearLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_Count
func miqt_exec_callback_QGraphicsLinearLayout_Count(self QGraphicsLinearLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QGraphicsLinearLayout) callVirtualBase_ItemAt(index int) *QGraphicsLayoutItem {

	return newQGraphicsLayoutItem(QGraphicsLinearLayout_virtualbase_ItemAt(unsafe.Pointer(this.h), (int)(index)))

}
func (this *QGraphicsLinearLayout) OnItemAt(slot func(super func(index int) *QGraphicsLayoutItem, index int) *QGraphicsLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_ItemAt
func miqt_exec_callback_QGraphicsLinearLayout_ItemAt(self QGraphicsLinearLayout, cb intptr_t, index int) *QGraphicsLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QGraphicsLayoutItem, index int) *QGraphicsLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_ItemAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsLinearLayout) callVirtualBase_Invalidate() {

	QGraphicsLinearLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QGraphicsLinearLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_Invalidate
func miqt_exec_callback_QGraphicsLinearLayout_Invalidate(self QGraphicsLinearLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QGraphicsLinearLayout) callVirtualBase_SizeHint(which SizeHint, constraint *QSizeF) *QSizeF {

	_goptr := newQSizeF(QGraphicsLinearLayout_virtualbase_SizeHint(unsafe.Pointer(this.h), (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsLinearLayout) OnSizeHint(slot func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_SizeHint
func miqt_exec_callback_QGraphicsLinearLayout_SizeHint(self QGraphicsLinearLayout, cb intptr_t, which int, constraint *QSizeF) *QSizeF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (SizeHint)(which)

	slotval2 := newQSizeF(constraint)

	virtualReturn := gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_SizeHint, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsLinearLayout) callVirtualBase_GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {

	QGraphicsLinearLayout_virtualbase_GetContentsMargins(unsafe.Pointer(this.h), (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))

}
func (this *QGraphicsLinearLayout) OnGetContentsMargins(slot func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_GetContentsMargins(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_GetContentsMargins
func miqt_exec_callback_QGraphicsLinearLayout_GetContentsMargins(self QGraphicsLinearLayout, cb intptr_t, left *double, top *double, right *double, bottom *double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*float64)(unsafe.Pointer(left))

	slotval2 := (*float64)(unsafe.Pointer(top))

	slotval3 := (*float64)(unsafe.Pointer(right))

	slotval4 := (*float64)(unsafe.Pointer(bottom))

	gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_GetContentsMargins, slotval1, slotval2, slotval3, slotval4)

}

func (this *QGraphicsLinearLayout) callVirtualBase_UpdateGeometry() {

	QGraphicsLinearLayout_virtualbase_UpdateGeometry(unsafe.Pointer(this.h))

}
func (this *QGraphicsLinearLayout) OnUpdateGeometry(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_UpdateGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_UpdateGeometry
func miqt_exec_callback_QGraphicsLinearLayout_UpdateGeometry(self QGraphicsLinearLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_UpdateGeometry)

}

func (this *QGraphicsLinearLayout) callVirtualBase_WidgetEvent(e *QEvent) {

	QGraphicsLinearLayout_virtualbase_WidgetEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QGraphicsLinearLayout) OnWidgetEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLinearLayout_override_virtual_WidgetEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLinearLayout_WidgetEvent
func miqt_exec_callback_QGraphicsLinearLayout_WidgetEvent(self QGraphicsLinearLayout, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QGraphicsLinearLayout{h: self}).callVirtualBase_WidgetEvent, slotval1)

}
