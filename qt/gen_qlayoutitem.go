package qt

import (
	"unsafe"
)

type QLayoutItem struct {
	h          uintptr
	isSubclass bool
}

// NewQLayoutItem constructs a new QLayoutItem object.
func NewQLayoutItem() *QLayoutItem {

	ret := newQLayoutItem(QLayoutItem_new())
	ret.isSubclass = true
	return ret
}

// NewQLayoutItem2 constructs a new QLayoutItem object.
func NewQLayoutItem2(param1 *QLayoutItem) *QLayoutItem {

	ret := newQLayoutItem(QLayoutItem_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLayoutItem3 constructs a new QLayoutItem object.
func NewQLayoutItem3(alignment AlignmentFlag) *QLayoutItem {

	ret := newQLayoutItem(QLayoutItem_new3((int)(alignment)))
	ret.isSubclass = true
	return ret
}

func (this *QLayoutItem) SizeHint() *QSize {
	_goptr := newQSize(QLayoutItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) MinimumSize() *QSize {
	_goptr := newQSize(QLayoutItem_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) MaximumSize() *QSize {
	_goptr := newQSize(QLayoutItem_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) ExpandingDirections() Orientation {
	return (Orientation)(QLayoutItem_ExpandingDirections(this.h))
}

func (this *QLayoutItem) SetGeometry(geometry *QRect) {
	QLayoutItem_SetGeometry(this.h, geometry.cPointer())
}

func (this *QLayoutItem) Geometry() *QRect {
	_goptr := newQRect(QLayoutItem_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) IsEmpty() bool {
	return (bool)(QLayoutItem_IsEmpty(this.h))
}

func (this *QLayoutItem) HasHeightForWidth() bool {
	return (bool)(QLayoutItem_HasHeightForWidth(this.h))
}

func (this *QLayoutItem) HeightForWidth(param1 int) int {
	return (int)(QLayoutItem_HeightForWidth(this.h, (int)(param1)))
}

func (this *QLayoutItem) MinimumHeightForWidth(param1 int) int {
	return (int)(QLayoutItem_MinimumHeightForWidth(this.h, (int)(param1)))
}

func (this *QLayoutItem) Invalidate() {
	QLayoutItem_Invalidate(this.h)
}

func (this *QLayoutItem) Widget() *QWidget {
	return newQWidget(QLayoutItem_Widget(this.h))
}

func (this *QLayoutItem) Layout() *QLayout {
	return newQLayout(QLayoutItem_Layout(this.h))
}

func (this *QLayoutItem) SpacerItem() *QSpacerItem {
	return newQSpacerItem(QLayoutItem_SpacerItem(this.h))
}

func (this *QLayoutItem) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QLayoutItem_Alignment(this.h))
}

func (this *QLayoutItem) SetAlignment(a AlignmentFlag) {
	QLayoutItem_SetAlignment(this.h, (int)(a))
}

func (this *QLayoutItem) ControlTypes() ControlType {
	return (ControlType)(QLayoutItem_ControlTypes(this.h))
}
func (this *QLayoutItem) OnSizeHint(slot func() *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_SizeHint
func miqt_exec_callback_QLayoutItem_SizeHint(self QLayoutItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}
func (this *QLayoutItem) OnMinimumSize(slot func() *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_MinimumSize
func miqt_exec_callback_QLayoutItem_MinimumSize(self QLayoutItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}
func (this *QLayoutItem) OnMaximumSize(slot func() *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_MaximumSize
func miqt_exec_callback_QLayoutItem_MaximumSize(self QLayoutItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}
func (this *QLayoutItem) OnExpandingDirections(slot func() Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_ExpandingDirections
func miqt_exec_callback_QLayoutItem_ExpandingDirections(self QLayoutItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func() Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (int)(virtualReturn)

}
func (this *QLayoutItem) OnSetGeometry(slot func(geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_SetGeometry
func miqt_exec_callback_QLayoutItem_SetGeometry(self QLayoutItem, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc(slotval1)

}
func (this *QLayoutItem) OnGeometry(slot func() *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_Geometry
func miqt_exec_callback_QLayoutItem_Geometry(self QLayoutItem, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}
func (this *QLayoutItem) OnIsEmpty(slot func() bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_IsEmpty
func miqt_exec_callback_QLayoutItem_IsEmpty(self QLayoutItem, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func() bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (bool)(virtualReturn)

}

func (this *QLayoutItem) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QLayoutItem_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QLayoutItem) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_HasHeightForWidth
func miqt_exec_callback_QLayoutItem_HasHeightForWidth(self QLayoutItem, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayoutItem{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QLayoutItem) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QLayoutItem_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QLayoutItem) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_HeightForWidth
func miqt_exec_callback_QLayoutItem_HeightForWidth(self QLayoutItem, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QLayoutItem{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QLayoutItem) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QLayoutItem_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QLayoutItem) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_MinimumHeightForWidth
func miqt_exec_callback_QLayoutItem_MinimumHeightForWidth(self QLayoutItem, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QLayoutItem{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QLayoutItem) callVirtualBase_Invalidate() {

	QLayoutItem_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QLayoutItem) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_Invalidate
func miqt_exec_callback_QLayoutItem_Invalidate(self QLayoutItem, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QLayoutItem{h: self}).callVirtualBase_Invalidate)

}

func (this *QLayoutItem) callVirtualBase_Widget() *QWidget {

	return newQWidget(QLayoutItem_virtualbase_Widget(unsafe.Pointer(this.h)))

}
func (this *QLayoutItem) OnWidget(slot func(super func() *QWidget) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_Widget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_Widget
func miqt_exec_callback_QLayoutItem_Widget(self QLayoutItem, cb intptr_t) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWidget) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayoutItem{h: self}).callVirtualBase_Widget)

	return virtualReturn.cPointer()

}

func (this *QLayoutItem) callVirtualBase_Layout() *QLayout {

	return newQLayout(QLayoutItem_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QLayoutItem) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_Layout
func miqt_exec_callback_QLayoutItem_Layout(self QLayoutItem, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayoutItem{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QLayoutItem) callVirtualBase_SpacerItem() *QSpacerItem {

	return newQSpacerItem(QLayoutItem_virtualbase_SpacerItem(unsafe.Pointer(this.h)))

}
func (this *QLayoutItem) OnSpacerItem(slot func(super func() *QSpacerItem) *QSpacerItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_SpacerItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_SpacerItem
func miqt_exec_callback_QLayoutItem_SpacerItem(self QLayoutItem, cb intptr_t) *QSpacerItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSpacerItem) *QSpacerItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayoutItem{h: self}).callVirtualBase_SpacerItem)

	return virtualReturn.cPointer()

}

func (this *QLayoutItem) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QLayoutItem_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QLayoutItem) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayoutItem_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayoutItem_ControlTypes
func miqt_exec_callback_QLayoutItem_ControlTypes(self QLayoutItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayoutItem{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

type QSpacerItem struct {
	h          uintptr
	isSubclass bool
}

// NewQSpacerItem constructs a new QSpacerItem object.
func NewQSpacerItem(w int, h int) *QSpacerItem {

	ret := newQSpacerItem(QSpacerItem_new((int)(w), (int)(h)))
	ret.isSubclass = true
	return ret
}

// NewQSpacerItem2 constructs a new QSpacerItem object.
func NewQSpacerItem2(param1 *QSpacerItem) *QSpacerItem {

	ret := newQSpacerItem(QSpacerItem_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSpacerItem3 constructs a new QSpacerItem object.
func NewQSpacerItem3(w int, h int, hData QSizePolicy__Policy) *QSpacerItem {

	ret := newQSpacerItem(QSpacerItem_new3((int)(w), (int)(h), (int)(hData)))
	ret.isSubclass = true
	return ret
}

// NewQSpacerItem4 constructs a new QSpacerItem object.
func NewQSpacerItem4(w int, h int, hData QSizePolicy__Policy, vData QSizePolicy__Policy) *QSpacerItem {

	ret := newQSpacerItem(QSpacerItem_new4((int)(w), (int)(h), (int)(hData), (int)(vData)))
	ret.isSubclass = true
	return ret
}

func (this *QSpacerItem) ChangeSize(w int, h int) {
	QSpacerItem_ChangeSize(this.h, (int)(w), (int)(h))
}

func (this *QSpacerItem) SizeHint() *QSize {
	_goptr := newQSize(QSpacerItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) MinimumSize() *QSize {
	_goptr := newQSize(QSpacerItem_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) MaximumSize() *QSize {
	_goptr := newQSize(QSpacerItem_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) ExpandingDirections() Orientation {
	return (Orientation)(QSpacerItem_ExpandingDirections(this.h))
}

func (this *QSpacerItem) IsEmpty() bool {
	return (bool)(QSpacerItem_IsEmpty(this.h))
}

func (this *QSpacerItem) SetGeometry(geometry *QRect) {
	QSpacerItem_SetGeometry(this.h, geometry.cPointer())
}

func (this *QSpacerItem) Geometry() *QRect {
	_goptr := newQRect(QSpacerItem_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) SpacerItem() *QSpacerItem {
	return newQSpacerItem(QSpacerItem_SpacerItem(this.h))
}

func (this *QSpacerItem) SizePolicy() *QSizePolicy {
	_goptr := newQSizePolicy(QSpacerItem_SizePolicy(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) ChangeSize3(w int, h int, hData QSizePolicy__Policy) {
	QSpacerItem_ChangeSize3(this.h, (int)(w), (int)(h), (int)(hData))
}

func (this *QSpacerItem) ChangeSize4(w int, h int, hData QSizePolicy__Policy, vData QSizePolicy__Policy) {
	QSpacerItem_ChangeSize4(this.h, (int)(w), (int)(h), (int)(hData), (int)(vData))
}

func (this *QSpacerItem) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QSpacerItem_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSpacerItem) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_SizeHint
func miqt_exec_callback_QSpacerItem_SizeHint(self QSpacerItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QSpacerItem) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QSpacerItem_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSpacerItem) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_MinimumSize
func miqt_exec_callback_QSpacerItem_MinimumSize(self QSpacerItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QSpacerItem) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QSpacerItem_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSpacerItem) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_MaximumSize
func miqt_exec_callback_QSpacerItem_MaximumSize(self QSpacerItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QSpacerItem) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QSpacerItem_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QSpacerItem) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_ExpandingDirections
func miqt_exec_callback_QSpacerItem_ExpandingDirections(self QSpacerItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QSpacerItem) callVirtualBase_IsEmpty() bool {

	return (bool)(QSpacerItem_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QSpacerItem) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_IsEmpty
func miqt_exec_callback_QSpacerItem_IsEmpty(self QSpacerItem, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QSpacerItem) callVirtualBase_SetGeometry(geometry *QRect) {

	QSpacerItem_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QSpacerItem) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_SetGeometry
func miqt_exec_callback_QSpacerItem_SetGeometry(self QSpacerItem, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QSpacerItem{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QSpacerItem) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QSpacerItem_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSpacerItem) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_Geometry
func miqt_exec_callback_QSpacerItem_Geometry(self QSpacerItem, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QSpacerItem) callVirtualBase_SpacerItem() *QSpacerItem {

	return newQSpacerItem(QSpacerItem_virtualbase_SpacerItem(unsafe.Pointer(this.h)))

}
func (this *QSpacerItem) OnSpacerItem(slot func(super func() *QSpacerItem) *QSpacerItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_SpacerItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_SpacerItem
func miqt_exec_callback_QSpacerItem_SpacerItem(self QSpacerItem, cb intptr_t) *QSpacerItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSpacerItem) *QSpacerItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_SpacerItem)

	return virtualReturn.cPointer()

}

func (this *QSpacerItem) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QSpacerItem_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QSpacerItem) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_HasHeightForWidth
func miqt_exec_callback_QSpacerItem_HasHeightForWidth(self QSpacerItem, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QSpacerItem) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QSpacerItem_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QSpacerItem) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_HeightForWidth
func miqt_exec_callback_QSpacerItem_HeightForWidth(self QSpacerItem, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QSpacerItem) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QSpacerItem_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QSpacerItem) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_MinimumHeightForWidth
func miqt_exec_callback_QSpacerItem_MinimumHeightForWidth(self QSpacerItem, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QSpacerItem) callVirtualBase_Invalidate() {

	QSpacerItem_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QSpacerItem) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_Invalidate
func miqt_exec_callback_QSpacerItem_Invalidate(self QSpacerItem, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QSpacerItem{h: self}).callVirtualBase_Invalidate)

}

func (this *QSpacerItem) callVirtualBase_Widget() *QWidget {

	return newQWidget(QSpacerItem_virtualbase_Widget(unsafe.Pointer(this.h)))

}
func (this *QSpacerItem) OnWidget(slot func(super func() *QWidget) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_Widget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_Widget
func miqt_exec_callback_QSpacerItem_Widget(self QSpacerItem, cb intptr_t) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWidget) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_Widget)

	return virtualReturn.cPointer()

}

func (this *QSpacerItem) callVirtualBase_Layout() *QLayout {

	return newQLayout(QSpacerItem_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QSpacerItem) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_Layout
func miqt_exec_callback_QSpacerItem_Layout(self QSpacerItem, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QSpacerItem) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QSpacerItem_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QSpacerItem) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpacerItem_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpacerItem_ControlTypes
func miqt_exec_callback_QSpacerItem_ControlTypes(self QSpacerItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpacerItem{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

type QWidgetItem struct {
	h          uintptr
	isSubclass bool
}

// NewQWidgetItem constructs a new QWidgetItem object.
func NewQWidgetItem(w *QWidget) *QWidgetItem {

	ret := newQWidgetItem(QWidgetItem_new(w.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QWidgetItem) SizeHint() *QSize {
	_goptr := newQSize(QWidgetItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) MinimumSize() *QSize {
	_goptr := newQSize(QWidgetItem_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) MaximumSize() *QSize {
	_goptr := newQSize(QWidgetItem_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) ExpandingDirections() Orientation {
	return (Orientation)(QWidgetItem_ExpandingDirections(this.h))
}

func (this *QWidgetItem) IsEmpty() bool {
	return (bool)(QWidgetItem_IsEmpty(this.h))
}

func (this *QWidgetItem) SetGeometry(geometry *QRect) {
	QWidgetItem_SetGeometry(this.h, geometry.cPointer())
}

func (this *QWidgetItem) Geometry() *QRect {
	_goptr := newQRect(QWidgetItem_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) Widget() *QWidget {
	return newQWidget(QWidgetItem_Widget(this.h))
}

func (this *QWidgetItem) HasHeightForWidth() bool {
	return (bool)(QWidgetItem_HasHeightForWidth(this.h))
}

func (this *QWidgetItem) HeightForWidth(param1 int) int {
	return (int)(QWidgetItem_HeightForWidth(this.h, (int)(param1)))
}

func (this *QWidgetItem) MinimumHeightForWidth(param1 int) int {
	return (int)(QWidgetItem_MinimumHeightForWidth(this.h, (int)(param1)))
}

func (this *QWidgetItem) ControlTypes() ControlType {
	return (ControlType)(QWidgetItem_ControlTypes(this.h))
}

func (this *QWidgetItem) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QWidgetItem_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItem) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_SizeHint
func miqt_exec_callback_QWidgetItem_SizeHint(self QWidgetItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QWidgetItem) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QWidgetItem_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItem) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_MinimumSize
func miqt_exec_callback_QWidgetItem_MinimumSize(self QWidgetItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QWidgetItem) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QWidgetItem_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItem) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_MaximumSize
func miqt_exec_callback_QWidgetItem_MaximumSize(self QWidgetItem, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QWidgetItem) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QWidgetItem_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QWidgetItem) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_ExpandingDirections
func miqt_exec_callback_QWidgetItem_ExpandingDirections(self QWidgetItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QWidgetItem) callVirtualBase_IsEmpty() bool {

	return (bool)(QWidgetItem_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QWidgetItem) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_IsEmpty
func miqt_exec_callback_QWidgetItem_IsEmpty(self QWidgetItem, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QWidgetItem) callVirtualBase_SetGeometry(geometry *QRect) {

	QWidgetItem_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QWidgetItem) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_SetGeometry
func miqt_exec_callback_QWidgetItem_SetGeometry(self QWidgetItem, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QWidgetItem{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QWidgetItem) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QWidgetItem_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItem) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_Geometry
func miqt_exec_callback_QWidgetItem_Geometry(self QWidgetItem, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QWidgetItem) callVirtualBase_Widget() *QWidget {

	return newQWidget(QWidgetItem_virtualbase_Widget(unsafe.Pointer(this.h)))

}
func (this *QWidgetItem) OnWidget(slot func(super func() *QWidget) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_Widget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_Widget
func miqt_exec_callback_QWidgetItem_Widget(self QWidgetItem, cb intptr_t) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWidget) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_Widget)

	return virtualReturn.cPointer()

}

func (this *QWidgetItem) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QWidgetItem_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QWidgetItem) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_HasHeightForWidth
func miqt_exec_callback_QWidgetItem_HasHeightForWidth(self QWidgetItem, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QWidgetItem) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QWidgetItem_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QWidgetItem) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_HeightForWidth
func miqt_exec_callback_QWidgetItem_HeightForWidth(self QWidgetItem, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QWidgetItem) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QWidgetItem_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QWidgetItem) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_MinimumHeightForWidth
func miqt_exec_callback_QWidgetItem_MinimumHeightForWidth(self QWidgetItem, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QWidgetItem) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QWidgetItem_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QWidgetItem) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_ControlTypes
func miqt_exec_callback_QWidgetItem_ControlTypes(self QWidgetItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

func (this *QWidgetItem) callVirtualBase_Invalidate() {

	QWidgetItem_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QWidgetItem) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_Invalidate
func miqt_exec_callback_QWidgetItem_Invalidate(self QWidgetItem, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QWidgetItem{h: self}).callVirtualBase_Invalidate)

}

func (this *QWidgetItem) callVirtualBase_Layout() *QLayout {

	return newQLayout(QWidgetItem_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QWidgetItem) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_Layout
func miqt_exec_callback_QWidgetItem_Layout(self QWidgetItem, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QWidgetItem) callVirtualBase_SpacerItem() *QSpacerItem {

	return newQSpacerItem(QWidgetItem_virtualbase_SpacerItem(unsafe.Pointer(this.h)))

}
func (this *QWidgetItem) OnSpacerItem(slot func(super func() *QSpacerItem) *QSpacerItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItem_override_virtual_SpacerItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItem_SpacerItem
func miqt_exec_callback_QWidgetItem_SpacerItem(self QWidgetItem, cb intptr_t) *QSpacerItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSpacerItem) *QSpacerItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItem{h: self}).callVirtualBase_SpacerItem)

	return virtualReturn.cPointer()

}

type QWidgetItemV2 struct {
	h          uintptr
	isSubclass bool
}

// NewQWidgetItemV2 constructs a new QWidgetItemV2 object.
func NewQWidgetItemV2(widget *QWidget) *QWidgetItemV2 {

	ret := newQWidgetItemV2(QWidgetItemV2_new(widget.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QWidgetItemV2) SizeHint() *QSize {
	_goptr := newQSize(QWidgetItemV2_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItemV2) MinimumSize() *QSize {
	_goptr := newQSize(QWidgetItemV2_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItemV2) MaximumSize() *QSize {
	_goptr := newQSize(QWidgetItemV2_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItemV2) HeightForWidth(width int) int {
	return (int)(QWidgetItemV2_HeightForWidth(this.h, (int)(width)))
}

func (this *QWidgetItemV2) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QWidgetItemV2_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItemV2) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_SizeHint
func miqt_exec_callback_QWidgetItemV2_SizeHint(self QWidgetItemV2, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QWidgetItemV2) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QWidgetItemV2_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItemV2) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_MinimumSize
func miqt_exec_callback_QWidgetItemV2_MinimumSize(self QWidgetItemV2, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QWidgetItemV2) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QWidgetItemV2_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItemV2) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_MaximumSize
func miqt_exec_callback_QWidgetItemV2_MaximumSize(self QWidgetItemV2, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QWidgetItemV2) callVirtualBase_HeightForWidth(width int) int {

	return (int)(QWidgetItemV2_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(width)))

}
func (this *QWidgetItemV2) OnHeightForWidth(slot func(super func(width int) int, width int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_HeightForWidth
func miqt_exec_callback_QWidgetItemV2_HeightForWidth(self QWidgetItemV2, cb intptr_t, width int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(width int) int, width int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(width)

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QWidgetItemV2) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QWidgetItemV2_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QWidgetItemV2) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_ExpandingDirections
func miqt_exec_callback_QWidgetItemV2_ExpandingDirections(self QWidgetItemV2, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QWidgetItemV2) callVirtualBase_IsEmpty() bool {

	return (bool)(QWidgetItemV2_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QWidgetItemV2) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_IsEmpty
func miqt_exec_callback_QWidgetItemV2_IsEmpty(self QWidgetItemV2, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QWidgetItemV2) callVirtualBase_SetGeometry(geometry *QRect) {

	QWidgetItemV2_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QWidgetItemV2) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_SetGeometry
func miqt_exec_callback_QWidgetItemV2_SetGeometry(self QWidgetItemV2, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QWidgetItemV2{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QWidgetItemV2) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QWidgetItemV2_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWidgetItemV2) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_Geometry
func miqt_exec_callback_QWidgetItemV2_Geometry(self QWidgetItemV2, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QWidgetItemV2) callVirtualBase_Widget() *QWidget {

	return newQWidget(QWidgetItemV2_virtualbase_Widget(unsafe.Pointer(this.h)))

}
func (this *QWidgetItemV2) OnWidget(slot func(super func() *QWidget) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_Widget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_Widget
func miqt_exec_callback_QWidgetItemV2_Widget(self QWidgetItemV2, cb intptr_t) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QWidget) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_Widget)

	return virtualReturn.cPointer()

}

func (this *QWidgetItemV2) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QWidgetItemV2_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QWidgetItemV2) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_HasHeightForWidth
func miqt_exec_callback_QWidgetItemV2_HasHeightForWidth(self QWidgetItemV2, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QWidgetItemV2) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QWidgetItemV2_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QWidgetItemV2) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_MinimumHeightForWidth
func miqt_exec_callback_QWidgetItemV2_MinimumHeightForWidth(self QWidgetItemV2, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QWidgetItemV2) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QWidgetItemV2_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QWidgetItemV2) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetItemV2_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetItemV2_ControlTypes
func miqt_exec_callback_QWidgetItemV2_ControlTypes(self QWidgetItemV2, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetItemV2{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}
