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

	ret := newQGraphicsLayoutItem(QGraphicsLayoutItem_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsLayoutItem2 constructs a new QGraphicsLayoutItem object.
func NewQGraphicsLayoutItem2(parent *QGraphicsLayoutItem) *QGraphicsLayoutItem {

	ret := newQGraphicsLayoutItem(QGraphicsLayoutItem_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsLayoutItem3 constructs a new QGraphicsLayoutItem object.
func NewQGraphicsLayoutItem3(parent *QGraphicsLayoutItem, isLayout bool) *QGraphicsLayoutItem {

	ret := newQGraphicsLayoutItem(QGraphicsLayoutItem_new3(parent.cPointer(), (bool)(isLayout)))
	ret.isSubclass = true
	return ret
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

func (this *QGraphicsLayoutItem) callVirtualBase_SetGeometry(rect *QRectF) {

	QGraphicsLayoutItem_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QGraphicsLayoutItem) OnSetGeometry(slot func(super func(rect *QRectF), rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayoutItem_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayoutItem_SetGeometry
func miqt_exec_callback_QGraphicsLayoutItem_SetGeometry(self QGraphicsLayoutItem, cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF), rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc((&QGraphicsLayoutItem{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGraphicsLayoutItem) callVirtualBase_GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {

	QGraphicsLayoutItem_virtualbase_GetContentsMargins(unsafe.Pointer(this.h), (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))

}
func (this *QGraphicsLayoutItem) OnGetContentsMargins(slot func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayoutItem_override_virtual_GetContentsMargins(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayoutItem_GetContentsMargins
func miqt_exec_callback_QGraphicsLayoutItem_GetContentsMargins(self QGraphicsLayoutItem, cb intptr_t, left *double, top *double, right *double, bottom *double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*float64)(unsafe.Pointer(left))

	slotval2 := (*float64)(unsafe.Pointer(top))

	slotval3 := (*float64)(unsafe.Pointer(right))

	slotval4 := (*float64)(unsafe.Pointer(bottom))

	gofunc((&QGraphicsLayoutItem{h: self}).callVirtualBase_GetContentsMargins, slotval1, slotval2, slotval3, slotval4)

}

func (this *QGraphicsLayoutItem) callVirtualBase_UpdateGeometry() {

	QGraphicsLayoutItem_virtualbase_UpdateGeometry(unsafe.Pointer(this.h))

}
func (this *QGraphicsLayoutItem) OnUpdateGeometry(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayoutItem_override_virtual_UpdateGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayoutItem_UpdateGeometry
func miqt_exec_callback_QGraphicsLayoutItem_UpdateGeometry(self QGraphicsLayoutItem, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsLayoutItem{h: self}).callVirtualBase_UpdateGeometry)

}

func (this *QGraphicsLayoutItem) callVirtualBase_IsEmpty() bool {

	return (bool)(QGraphicsLayoutItem_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QGraphicsLayoutItem) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayoutItem_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayoutItem_IsEmpty
func miqt_exec_callback_QGraphicsLayoutItem_IsEmpty(self QGraphicsLayoutItem, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsLayoutItem{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}
func (this *QGraphicsLayoutItem) OnSizeHint(slot func(which SizeHint, constraint *QSizeF) *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsLayoutItem_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsLayoutItem_SizeHint
func miqt_exec_callback_QGraphicsLayoutItem_SizeHint(self QGraphicsLayoutItem, cb intptr_t, which int, constraint *QSizeF) *QSizeF {
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
