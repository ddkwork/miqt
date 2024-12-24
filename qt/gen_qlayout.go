package qt

import (
	"unsafe"
)

type QLayout__SizeConstraint int

const (
	QLayout__SetDefaultConstraint QLayout__SizeConstraint = 0
	QLayout__SetNoConstraint      QLayout__SizeConstraint = 1
	QLayout__SetMinimumSize       QLayout__SizeConstraint = 2
	QLayout__SetFixedSize         QLayout__SizeConstraint = 3
	QLayout__SetMaximumSize       QLayout__SizeConstraint = 4
	QLayout__SetMinAndMaxSize     QLayout__SizeConstraint = 5
)

type QLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQLayout constructs a new QLayout object.
func NewQLayout(parent *QWidget) *QLayout {

	ret := newQLayout(QLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLayout2 constructs a new QLayout object.
func NewQLayout2() *QLayout {

	ret := newQLayout(QLayout_new2())
	ret.isSubclass = true
	return ret
}

func (this *QLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QLayout_MetaObject(this.h))
}

func (this *QLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLayout_Metacast(this.h, param1_Cstring))
}

func QLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLayout) Spacing() int {
	return (int)(QLayout_Spacing(this.h))
}

func (this *QLayout) SetSpacing(spacing int) {
	QLayout_SetSpacing(this.h, (int)(spacing))
}

func (this *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	QLayout_SetContentsMargins(this.h, (int)(left), (int)(top), (int)(right), (int)(bottom))
}

func (this *QLayout) SetContentsMarginsWithMargins(margins *QMargins) {
	QLayout_SetContentsMarginsWithMargins(this.h, margins.cPointer())
}

func (this *QLayout) UnsetContentsMargins() {
	QLayout_UnsetContentsMargins(this.h)
}

func (this *QLayout) GetContentsMargins(left *int, top *int, right *int, bottom *int) {
	QLayout_GetContentsMargins(this.h, (*int)(unsafe.Pointer(left)), (*int)(unsafe.Pointer(top)), (*int)(unsafe.Pointer(right)), (*int)(unsafe.Pointer(bottom)))
}

func (this *QLayout) ContentsMargins() *QMargins {
	_goptr := newQMargins(QLayout_ContentsMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) ContentsRect() *QRect {
	_goptr := newQRect(QLayout_ContentsRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) SetAlignment(w *QWidget, alignment AlignmentFlag) bool {
	return (bool)(QLayout_SetAlignment(this.h, w.cPointer(), (int)(alignment)))
}

func (this *QLayout) SetAlignment2(l *QLayout, alignment AlignmentFlag) bool {
	return (bool)(QLayout_SetAlignment2(this.h, l.cPointer(), (int)(alignment)))
}

func (this *QLayout) SetSizeConstraint(sizeConstraint SizeConstraint) {
	QLayout_SetSizeConstraint(this.h, sizeConstraint)
}

func (this *QLayout) SizeConstraint() SizeConstraint {
	xxxxxxxxx
}

func (this *QLayout) SetMenuBar(w *QWidget) {
	QLayout_SetMenuBar(this.h, w.cPointer())
}

func (this *QLayout) MenuBar() *QWidget {
	return newQWidget(QLayout_MenuBar(this.h))
}

func (this *QLayout) ParentWidget() *QWidget {
	return newQWidget(QLayout_ParentWidget(this.h))
}

func (this *QLayout) Invalidate() {
	QLayout_Invalidate(this.h)
}

func (this *QLayout) Geometry() *QRect {
	_goptr := newQRect(QLayout_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) Activate() bool {
	return (bool)(QLayout_Activate(this.h))
}

func (this *QLayout) Update() {
	QLayout_Update(this.h)
}

func (this *QLayout) AddWidget(w *QWidget) {
	QLayout_AddWidget(this.h, w.cPointer())
}

func (this *QLayout) AddItem(param1 *QLayoutItem) {
	QLayout_AddItem(this.h, param1.cPointer())
}

func (this *QLayout) RemoveWidget(w *QWidget) {
	QLayout_RemoveWidget(this.h, w.cPointer())
}

func (this *QLayout) RemoveItem(param1 *QLayoutItem) {
	QLayout_RemoveItem(this.h, param1.cPointer())
}

func (this *QLayout) ExpandingDirections() Orientation {
	return (Orientation)(QLayout_ExpandingDirections(this.h))
}

func (this *QLayout) MinimumSize() *QSize {
	_goptr := newQSize(QLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) MaximumSize() *QSize {
	_goptr := newQSize(QLayout_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) SetGeometry(geometry *QRect) {
	QLayout_SetGeometry(this.h, geometry.cPointer())
}

func (this *QLayout) ItemAt(index int) *QLayoutItem {
	return newQLayoutItem(QLayout_ItemAt(this.h, (int)(index)))
}

func (this *QLayout) TakeAt(index int) *QLayoutItem {
	return newQLayoutItem(QLayout_TakeAt(this.h, (int)(index)))
}

func (this *QLayout) IndexOf(param1 *QWidget) int {
	return (int)(QLayout_IndexOf(this.h, param1.cPointer()))
}

func (this *QLayout) IndexOfWithQLayoutItem(param1 *QLayoutItem) int {
	return (int)(QLayout_IndexOfWithQLayoutItem(this.h, param1.cPointer()))
}

func (this *QLayout) Count() int {
	return (int)(QLayout_Count(this.h))
}

func (this *QLayout) IsEmpty() bool {
	return (bool)(QLayout_IsEmpty(this.h))
}

func (this *QLayout) ControlTypes() ControlType {
	return (ControlType)(QLayout_ControlTypes(this.h))
}

func (this *QLayout) ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	return newQLayoutItem(QLayout_ReplaceWidget(this.h, from.cPointer(), to.cPointer(), (int)(options)))
}

func (this *QLayout) TotalMinimumHeightForWidth(w int) int {
	return (int)(QLayout_TotalMinimumHeightForWidth(this.h, (int)(w)))
}

func (this *QLayout) TotalHeightForWidth(w int) int {
	return (int)(QLayout_TotalHeightForWidth(this.h, (int)(w)))
}

func (this *QLayout) TotalMinimumSize() *QSize {
	_goptr := newQSize(QLayout_TotalMinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) TotalMaximumSize() *QSize {
	_goptr := newQSize(QLayout_TotalMaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) TotalSizeHint() *QSize {
	_goptr := newQSize(QLayout_TotalSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) Layout() *QLayout {
	return newQLayout(QLayout_Layout(this.h))
}

func (this *QLayout) SetEnabled(enabled bool) {
	QLayout_SetEnabled(this.h, (bool)(enabled))
}

func (this *QLayout) IsEnabled() bool {
	return (bool)(QLayout_IsEnabled(this.h))
}

func QLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	_goptr := newQSize(QLayout_ClosestAcceptableSize(w.cPointer(), s.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLayout) callVirtualBase_Spacing() int {

	return (int)(QLayout_virtualbase_Spacing(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnSpacing(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Spacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Spacing
func miqt_exec_callback_QLayout_Spacing(self QLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_Spacing)

	return (int)(virtualReturn)

}

func (this *QLayout) callVirtualBase_SetSpacing(spacing int) {

	QLayout_virtualbase_SetSpacing(unsafe.Pointer(this.h), (int)(spacing))

}
func (this *QLayout) OnSetSpacing(slot func(super func(spacing int), spacing int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_SetSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_SetSpacing
func miqt_exec_callback_QLayout_SetSpacing(self QLayout, cb intptr_t, spacing int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(spacing int), spacing int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(spacing)

	gofunc((&QLayout{h: self}).callVirtualBase_SetSpacing, slotval1)

}

func (this *QLayout) callVirtualBase_Invalidate() {

	QLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Invalidate
func miqt_exec_callback_QLayout_Invalidate(self QLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QLayout) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QLayout_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QLayout) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Geometry
func miqt_exec_callback_QLayout_Geometry(self QLayout, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}
func (this *QLayout) OnAddItem(slot func(param1 *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_AddItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_AddItem
func miqt_exec_callback_QLayout_AddItem(self QLayout, cb intptr_t, param1 *QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(param1)

	gofunc(slotval1)

}

func (this *QLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_ExpandingDirections
func miqt_exec_callback_QLayout_ExpandingDirections(self QLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_MinimumSize
func miqt_exec_callback_QLayout_MinimumSize(self QLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_MaximumSize
func miqt_exec_callback_QLayout_MaximumSize(self QLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QLayout) callVirtualBase_SetGeometry(geometry *QRect) {

	QLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QLayout) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_SetGeometry
func miqt_exec_callback_QLayout_SetGeometry(self QLayout, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}
func (this *QLayout) OnItemAt(slot func(index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_ItemAt
func miqt_exec_callback_QLayout_ItemAt(self QLayout, cb intptr_t, index int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}
func (this *QLayout) OnTakeAt(slot func(index int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_TakeAt
func miqt_exec_callback_QLayout_TakeAt(self QLayout, cb intptr_t, index int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}

func (this *QLayout) callVirtualBase_IndexOf(param1 *QWidget) int {

	return (int)(QLayout_virtualbase_IndexOf(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QLayout) OnIndexOf(slot func(super func(param1 *QWidget) int, param1 *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_IndexOf(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_IndexOf
func miqt_exec_callback_QLayout_IndexOf(self QLayout, cb intptr_t, param1 *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWidget) int, param1 *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(param1)

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_IndexOf, slotval1)

	return (int)(virtualReturn)

}

func (this *QLayout) callVirtualBase_IndexOfWithQLayoutItem(param1 *QLayoutItem) int {

	return (int)(QLayout_virtualbase_IndexOfWithQLayoutItem(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QLayout) OnIndexOfWithQLayoutItem(slot func(super func(param1 *QLayoutItem) int, param1 *QLayoutItem) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_IndexOfWithQLayoutItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_IndexOfWithQLayoutItem
func miqt_exec_callback_QLayout_IndexOfWithQLayoutItem(self QLayout, cb intptr_t, param1 *QLayoutItem) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QLayoutItem) int, param1 *QLayoutItem) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(param1)

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_IndexOfWithQLayoutItem, slotval1)

	return (int)(virtualReturn)

}
func (this *QLayout) OnCount(slot func() int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Count
func miqt_exec_callback_QLayout_Count(self QLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func() int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (int)(virtualReturn)

}

func (this *QLayout) callVirtualBase_IsEmpty() bool {

	return (bool)(QLayout_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_IsEmpty
func miqt_exec_callback_QLayout_IsEmpty(self QLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QLayout) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QLayout_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_ControlTypes
func miqt_exec_callback_QLayout_ControlTypes(self QLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

func (this *QLayout) callVirtualBase_ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {

	return newQLayoutItem(QLayout_virtualbase_ReplaceWidget(unsafe.Pointer(this.h), from.cPointer(), to.cPointer(), (int)(options)))

}
func (this *QLayout) OnReplaceWidget(slot func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_ReplaceWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_ReplaceWidget
func miqt_exec_callback_QLayout_ReplaceWidget(self QLayout, cb intptr_t, from *QWidget, to *QWidget, options int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(from)

	slotval2 := newQWidget(to)

	slotval3 := (FindChildOption)(options)

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_ReplaceWidget, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QLayout) callVirtualBase_Layout() *QLayout {

	return newQLayout(QLayout_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Layout
func miqt_exec_callback_QLayout_Layout(self QLayout, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QLayout) callVirtualBase_ChildEvent(e *QChildEvent) {

	QLayout_virtualbase_ChildEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QLayout) OnChildEvent(slot func(super func(e *QChildEvent), e *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_ChildEvent
func miqt_exec_callback_QLayout_ChildEvent(self QLayout, cb intptr_t, e *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QChildEvent), e *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(e)

	gofunc((&QLayout{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QLayout) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QLayout_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QLayout) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Event
func miqt_exec_callback_QLayout_Event(self QLayout, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QLayout) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QLayout_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QLayout) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_EventFilter
func miqt_exec_callback_QLayout_EventFilter(self QLayout, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QLayout) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QLayout_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLayout) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_TimerEvent
func miqt_exec_callback_QLayout_TimerEvent(self QLayout, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QLayout{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QLayout) callVirtualBase_CustomEvent(event *QEvent) {

	QLayout_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QLayout) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_CustomEvent
func miqt_exec_callback_QLayout_CustomEvent(self QLayout, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QLayout{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QLayout) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QLayout_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QLayout) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_ConnectNotify
func miqt_exec_callback_QLayout_ConnectNotify(self QLayout, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QLayout{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QLayout) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QLayout_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QLayout) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_DisconnectNotify
func miqt_exec_callback_QLayout_DisconnectNotify(self QLayout, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QLayout{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
func (this *QLayout) OnSizeHint(slot func() *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_SizeHint
func miqt_exec_callback_QLayout_SizeHint(self QLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}

func (this *QLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_HasHeightForWidth
func miqt_exec_callback_QLayout_HasHeightForWidth(self QLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QLayout) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QLayout) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_HeightForWidth
func miqt_exec_callback_QLayout_HeightForWidth(self QLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QLayout) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QLayout_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QLayout) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_MinimumHeightForWidth
func miqt_exec_callback_QLayout_MinimumHeightForWidth(self QLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QLayout) callVirtualBase_Widget() *QWidget {

	return newQWidget(QLayout_virtualbase_Widget(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnWidget(slot func(super func() *QWidget) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Widget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Widget
func miqt_exec_callback_QLayout_Widget(self QLayout, cb intptr_t) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWidget) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_Widget)

	return virtualReturn.cPointer()

}

func (this *QLayout) callVirtualBase_SpacerItem() *QSpacerItem {

	return newQSpacerItem(QLayout_virtualbase_SpacerItem(unsafe.Pointer(this.h)))

}
func (this *QLayout) OnSpacerItem(slot func(super func() *QSpacerItem) *QSpacerItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_SpacerItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_SpacerItem
func miqt_exec_callback_QLayout_SpacerItem(self QLayout, cb intptr_t) *QSpacerItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSpacerItem) *QSpacerItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_SpacerItem)

	return virtualReturn.cPointer()

}
