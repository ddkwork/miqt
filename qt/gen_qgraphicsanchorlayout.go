package qt

import (
	"unsafe"
)

type QGraphicsAnchor struct {
	h          uintptr
	isSubclass bool
}

func (this *QGraphicsAnchor) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsAnchor_MetaObject(this.h))
}

func (this *QGraphicsAnchor) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsAnchor_Metacast(this.h, param1_Cstring))
}

func QGraphicsAnchor_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsAnchor_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsAnchor) SetSpacing(spacing float64) {
	QGraphicsAnchor_SetSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchor) UnsetSpacing() {
	QGraphicsAnchor_UnsetSpacing(this.h)
}

func (this *QGraphicsAnchor) Spacing() float64 {
	return (float64)(QGraphicsAnchor_Spacing(this.h))
}

func (this *QGraphicsAnchor) SetSizePolicy(policy QSizePolicy__Policy) {
	QGraphicsAnchor_SetSizePolicy(this.h, (int)(policy))
}

func (this *QGraphicsAnchor) SizePolicy() QSizePolicy__Policy {
	return (QSizePolicy__Policy)(QGraphicsAnchor_SizePolicy(this.h))
}

func QGraphicsAnchor_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsAnchor_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsAnchor_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsAnchor_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QGraphicsAnchorLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsAnchorLayout constructs a new QGraphicsAnchorLayout object.
func NewQGraphicsAnchorLayout() *QGraphicsAnchorLayout {

	ret := newQGraphicsAnchorLayout(QGraphicsAnchorLayout_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsAnchorLayout2 constructs a new QGraphicsAnchorLayout object.
func NewQGraphicsAnchorLayout2(parent *QGraphicsLayoutItem) *QGraphicsAnchorLayout {

	ret := newQGraphicsAnchorLayout(QGraphicsAnchorLayout_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsAnchorLayout) AddAnchor(firstItem *QGraphicsLayoutItem, firstEdge AnchorPoint, secondItem *QGraphicsLayoutItem, secondEdge AnchorPoint) *QGraphicsAnchor {
	return newQGraphicsAnchor(QGraphicsAnchorLayout_AddAnchor(this.h, firstItem.cPointer(), (int)(firstEdge), secondItem.cPointer(), (int)(secondEdge)))
}

func (this *QGraphicsAnchorLayout) Anchor(firstItem *QGraphicsLayoutItem, firstEdge AnchorPoint, secondItem *QGraphicsLayoutItem, secondEdge AnchorPoint) *QGraphicsAnchor {
	return newQGraphicsAnchor(QGraphicsAnchorLayout_Anchor(this.h, firstItem.cPointer(), (int)(firstEdge), secondItem.cPointer(), (int)(secondEdge)))
}

func (this *QGraphicsAnchorLayout) AddCornerAnchors(firstItem *QGraphicsLayoutItem, firstCorner Corner, secondItem *QGraphicsLayoutItem, secondCorner Corner) {
	QGraphicsAnchorLayout_AddCornerAnchors(this.h, firstItem.cPointer(), (int)(firstCorner), secondItem.cPointer(), (int)(secondCorner))
}

func (this *QGraphicsAnchorLayout) AddAnchors(firstItem *QGraphicsLayoutItem, secondItem *QGraphicsLayoutItem) {
	QGraphicsAnchorLayout_AddAnchors(this.h, firstItem.cPointer(), secondItem.cPointer())
}

func (this *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	QGraphicsAnchorLayout_SetHorizontalSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	QGraphicsAnchorLayout_SetVerticalSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	QGraphicsAnchorLayout_SetSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchorLayout) HorizontalSpacing() float64 {
	return (float64)(QGraphicsAnchorLayout_HorizontalSpacing(this.h))
}

func (this *QGraphicsAnchorLayout) VerticalSpacing() float64 {
	return (float64)(QGraphicsAnchorLayout_VerticalSpacing(this.h))
}

func (this *QGraphicsAnchorLayout) RemoveAt(index int) {
	QGraphicsAnchorLayout_RemoveAt(this.h, (int)(index))
}

func (this *QGraphicsAnchorLayout) SetGeometry(rect *QRectF) {
	QGraphicsAnchorLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QGraphicsAnchorLayout) Count() int {
	return (int)(QGraphicsAnchorLayout_Count(this.h))
}

func (this *QGraphicsAnchorLayout) ItemAt(index int) *QGraphicsLayoutItem {
	return newQGraphicsLayoutItem(QGraphicsAnchorLayout_ItemAt(this.h, (int)(index)))
}

func (this *QGraphicsAnchorLayout) Invalidate() {
	QGraphicsAnchorLayout_Invalidate(this.h)
}

func (this *QGraphicsAnchorLayout) AddAnchors3(firstItem *QGraphicsLayoutItem, secondItem *QGraphicsLayoutItem, orientations Orientation) {
	QGraphicsAnchorLayout_AddAnchors3(this.h, firstItem.cPointer(), secondItem.cPointer(), (int)(orientations))
}

func (this *QGraphicsAnchorLayout) callVirtualBase_RemoveAt(index int) {

	QGraphicsAnchorLayout_virtualbase_RemoveAt(unsafe.Pointer(this.h), (int)(index))

}
func (this *QGraphicsAnchorLayout) OnRemoveAt(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_RemoveAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_RemoveAt
func miqt_exec_callback_QGraphicsAnchorLayout_RemoveAt(self QGraphicsAnchorLayout, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_RemoveAt, slotval1)

}

func (this *QGraphicsAnchorLayout) callVirtualBase_SetGeometry(rect *QRectF) {

	QGraphicsAnchorLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QGraphicsAnchorLayout) OnSetGeometry(slot func(super func(rect *QRectF), rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_SetGeometry
func miqt_exec_callback_QGraphicsAnchorLayout_SetGeometry(self QGraphicsAnchorLayout, cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF), rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGraphicsAnchorLayout) callVirtualBase_Count() int {

	return (int)(QGraphicsAnchorLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QGraphicsAnchorLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_Count
func miqt_exec_callback_QGraphicsAnchorLayout_Count(self QGraphicsAnchorLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QGraphicsAnchorLayout) callVirtualBase_ItemAt(index int) *QGraphicsLayoutItem {

	return newQGraphicsLayoutItem(QGraphicsAnchorLayout_virtualbase_ItemAt(unsafe.Pointer(this.h), (int)(index)))

}
func (this *QGraphicsAnchorLayout) OnItemAt(slot func(super func(index int) *QGraphicsLayoutItem, index int) *QGraphicsLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_ItemAt
func miqt_exec_callback_QGraphicsAnchorLayout_ItemAt(self QGraphicsAnchorLayout, cb intptr_t, index int) *QGraphicsLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QGraphicsLayoutItem, index int) *QGraphicsLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_ItemAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsAnchorLayout) callVirtualBase_Invalidate() {

	QGraphicsAnchorLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QGraphicsAnchorLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_Invalidate
func miqt_exec_callback_QGraphicsAnchorLayout_Invalidate(self QGraphicsAnchorLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QGraphicsAnchorLayout) callVirtualBase_SizeHint(which SizeHint, constraint *QSizeF) *QSizeF {

	_goptr := newQSizeF(QGraphicsAnchorLayout_virtualbase_SizeHint(unsafe.Pointer(this.h), (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsAnchorLayout) OnSizeHint(slot func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_SizeHint
func miqt_exec_callback_QGraphicsAnchorLayout_SizeHint(self QGraphicsAnchorLayout, cb intptr_t, which int, constraint *QSizeF) *QSizeF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (SizeHint)(which)

	slotval2 := newQSizeF(constraint)

	virtualReturn := gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_SizeHint, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsAnchorLayout) callVirtualBase_GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {

	QGraphicsAnchorLayout_virtualbase_GetContentsMargins(unsafe.Pointer(this.h), (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))

}
func (this *QGraphicsAnchorLayout) OnGetContentsMargins(slot func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_GetContentsMargins(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_GetContentsMargins
func miqt_exec_callback_QGraphicsAnchorLayout_GetContentsMargins(self QGraphicsAnchorLayout, cb intptr_t, left *double, top *double, right *double, bottom *double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*float64)(unsafe.Pointer(left))

	slotval2 := (*float64)(unsafe.Pointer(top))

	slotval3 := (*float64)(unsafe.Pointer(right))

	slotval4 := (*float64)(unsafe.Pointer(bottom))

	gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_GetContentsMargins, slotval1, slotval2, slotval3, slotval4)

}

func (this *QGraphicsAnchorLayout) callVirtualBase_UpdateGeometry() {

	QGraphicsAnchorLayout_virtualbase_UpdateGeometry(unsafe.Pointer(this.h))

}
func (this *QGraphicsAnchorLayout) OnUpdateGeometry(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_UpdateGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_UpdateGeometry
func miqt_exec_callback_QGraphicsAnchorLayout_UpdateGeometry(self QGraphicsAnchorLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_UpdateGeometry)

}

func (this *QGraphicsAnchorLayout) callVirtualBase_WidgetEvent(e *QEvent) {

	QGraphicsAnchorLayout_virtualbase_WidgetEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QGraphicsAnchorLayout) OnWidgetEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsAnchorLayout_override_virtual_WidgetEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsAnchorLayout_WidgetEvent
func miqt_exec_callback_QGraphicsAnchorLayout_WidgetEvent(self QGraphicsAnchorLayout, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QGraphicsAnchorLayout{h: self}).callVirtualBase_WidgetEvent, slotval1)

}
