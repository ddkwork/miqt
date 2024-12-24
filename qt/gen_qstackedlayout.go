package qt

import (
	"unsafe"
)

type QStackedLayout__StackingMode int

const (
	QStackedLayout__StackOne QStackedLayout__StackingMode = 0
	QStackedLayout__StackAll QStackedLayout__StackingMode = 1
)

type QStackedLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQStackedLayout constructs a new QStackedLayout object.
func NewQStackedLayout(parent *QWidget) *QStackedLayout {

	ret := newQStackedLayout(QStackedLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStackedLayout2 constructs a new QStackedLayout object.
func NewQStackedLayout2() *QStackedLayout {

	ret := newQStackedLayout(QStackedLayout_new2())
	ret.isSubclass = true
	return ret
}

// NewQStackedLayout3 constructs a new QStackedLayout object.
func NewQStackedLayout3(parentLayout *QLayout) *QStackedLayout {

	ret := newQStackedLayout(QStackedLayout_new3(parentLayout.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStackedLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QStackedLayout_MetaObject(this.h))
}

func (this *QStackedLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStackedLayout_Metacast(this.h, param1_Cstring))
}

func QStackedLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStackedLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStackedLayout) AddWidget(w *QWidget) int {
	return (int)(QStackedLayout_AddWidget(this.h, w.cPointer()))
}

func (this *QStackedLayout) InsertWidget(index int, w *QWidget) int {
	return (int)(QStackedLayout_InsertWidget(this.h, (int)(index), w.cPointer()))
}

func (this *QStackedLayout) CurrentWidget() *QWidget {
	return newQWidget(QStackedLayout_CurrentWidget(this.h))
}

func (this *QStackedLayout) CurrentIndex() int {
	return (int)(QStackedLayout_CurrentIndex(this.h))
}

func (this *QStackedLayout) Widget(param1 int) *QWidget {
	return newQWidget(QStackedLayout_Widget(this.h, (int)(param1)))
}

func (this *QStackedLayout) Count() int {
	return (int)(QStackedLayout_Count(this.h))
}

func (this *QStackedLayout) StackingMode() StackingMode {
	xxxxxxxxx
}

func (this *QStackedLayout) SetStackingMode(stackingMode StackingMode) {
	QStackedLayout_SetStackingMode(this.h, stackingMode)
}

func (this *QStackedLayout) AddItem(item *QLayoutItem) {
	QStackedLayout_AddItem(this.h, item.cPointer())
}

func (this *QStackedLayout) SizeHint() *QSize {
	_goptr := newQSize(QStackedLayout_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStackedLayout) MinimumSize() *QSize {
	_goptr := newQSize(QStackedLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStackedLayout) ItemAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QStackedLayout_ItemAt(this.h, (int)(param1)))
}

func (this *QStackedLayout) TakeAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QStackedLayout_TakeAt(this.h, (int)(param1)))
}

func (this *QStackedLayout) SetGeometry(rect *QRect) {
	QStackedLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QStackedLayout) HasHeightForWidth() bool {
	return (bool)(QStackedLayout_HasHeightForWidth(this.h))
}

func (this *QStackedLayout) HeightForWidth(width int) int {
	return (int)(QStackedLayout_HeightForWidth(this.h, (int)(width)))
}

func (this *QStackedLayout) WidgetRemoved(index int) {
	QStackedLayout_WidgetRemoved(this.h, (int)(index))
}
func (this *QStackedLayout) OnWidgetRemoved(slot func(index int)) {
	QStackedLayout_connect_WidgetRemoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_WidgetRemoved
func miqt_exec_callback_QStackedLayout_WidgetRemoved(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QStackedLayout) CurrentChanged(index int) {
	QStackedLayout_CurrentChanged(this.h, (int)(index))
}
func (this *QStackedLayout) OnCurrentChanged(slot func(index int)) {
	QStackedLayout_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_CurrentChanged
func miqt_exec_callback_QStackedLayout_CurrentChanged(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QStackedLayout) WidgetAdded(index int) {
	QStackedLayout_WidgetAdded(this.h, (int)(index))
}
func (this *QStackedLayout) OnWidgetAdded(slot func(index int)) {
	QStackedLayout_connect_WidgetAdded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_WidgetAdded
func miqt_exec_callback_QStackedLayout_WidgetAdded(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QStackedLayout) SetCurrentIndex(index int) {
	QStackedLayout_SetCurrentIndex(this.h, (int)(index))
}

func (this *QStackedLayout) SetCurrentWidget(w *QWidget) {
	QStackedLayout_SetCurrentWidget(this.h, w.cPointer())
}

func QStackedLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStackedLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStackedLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStackedLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStackedLayout) callVirtualBase_Count() int {

	return (int)(QStackedLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QStackedLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_Count
func miqt_exec_callback_QStackedLayout_Count(self QStackedLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_AddItem(item *QLayoutItem) {

	QStackedLayout_virtualbase_AddItem(unsafe.Pointer(this.h), item.cPointer())

}
func (this *QStackedLayout) OnAddItem(slot func(super func(item *QLayoutItem), item *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_AddItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_AddItem
func miqt_exec_callback_QStackedLayout_AddItem(self QStackedLayout, cb intptr_t, item *QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(item *QLayoutItem), item *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(item)

	gofunc((&QStackedLayout{h: self}).callVirtualBase_AddItem, slotval1)

}

func (this *QStackedLayout) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QStackedLayout_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStackedLayout) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_SizeHint
func miqt_exec_callback_QStackedLayout_SizeHint(self QStackedLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QStackedLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStackedLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_MinimumSize
func miqt_exec_callback_QStackedLayout_MinimumSize(self QStackedLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_ItemAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QStackedLayout_virtualbase_ItemAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QStackedLayout) OnItemAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_ItemAt
func miqt_exec_callback_QStackedLayout_ItemAt(self QStackedLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_ItemAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_TakeAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QStackedLayout_virtualbase_TakeAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QStackedLayout) OnTakeAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_TakeAt
func miqt_exec_callback_QStackedLayout_TakeAt(self QStackedLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_TakeAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_SetGeometry(rect *QRect) {

	QStackedLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QStackedLayout) OnSetGeometry(slot func(super func(rect *QRect), rect *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_SetGeometry
func miqt_exec_callback_QStackedLayout_SetGeometry(self QStackedLayout, cb intptr_t, rect *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect), rect *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	gofunc((&QStackedLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QStackedLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QStackedLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QStackedLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_HasHeightForWidth
func miqt_exec_callback_QStackedLayout_HasHeightForWidth(self QStackedLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_HeightForWidth(width int) int {

	return (int)(QStackedLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(width)))

}
func (this *QStackedLayout) OnHeightForWidth(slot func(super func(width int) int, width int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_HeightForWidth
func miqt_exec_callback_QStackedLayout_HeightForWidth(self QStackedLayout, cb intptr_t, width int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(width int) int, width int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(width)

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_Spacing() int {

	return (int)(QStackedLayout_virtualbase_Spacing(unsafe.Pointer(this.h)))

}
func (this *QStackedLayout) OnSpacing(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_Spacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_Spacing
func miqt_exec_callback_QStackedLayout_Spacing(self QStackedLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_Spacing)

	return (int)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_SetSpacing(spacing int) {

	QStackedLayout_virtualbase_SetSpacing(unsafe.Pointer(this.h), (int)(spacing))

}
func (this *QStackedLayout) OnSetSpacing(slot func(super func(spacing int), spacing int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_SetSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_SetSpacing
func miqt_exec_callback_QStackedLayout_SetSpacing(self QStackedLayout, cb intptr_t, spacing int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(spacing int), spacing int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(spacing)

	gofunc((&QStackedLayout{h: self}).callVirtualBase_SetSpacing, slotval1)

}

func (this *QStackedLayout) callVirtualBase_Invalidate() {

	QStackedLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QStackedLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_Invalidate
func miqt_exec_callback_QStackedLayout_Invalidate(self QStackedLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QStackedLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QStackedLayout) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QStackedLayout_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStackedLayout) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_Geometry
func miqt_exec_callback_QStackedLayout_Geometry(self QStackedLayout, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QStackedLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QStackedLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_ExpandingDirections
func miqt_exec_callback_QStackedLayout_ExpandingDirections(self QStackedLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QStackedLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStackedLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_MaximumSize
func miqt_exec_callback_QStackedLayout_MaximumSize(self QStackedLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_IndexOf(param1 *QWidget) int {

	return (int)(QStackedLayout_virtualbase_IndexOf(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QStackedLayout) OnIndexOf(slot func(super func(param1 *QWidget) int, param1 *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_IndexOf(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_IndexOf
func miqt_exec_callback_QStackedLayout_IndexOf(self QStackedLayout, cb intptr_t, param1 *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWidget) int, param1 *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(param1)

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_IndexOf, slotval1)

	return (int)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_IsEmpty() bool {

	return (bool)(QStackedLayout_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QStackedLayout) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_IsEmpty
func miqt_exec_callback_QStackedLayout_IsEmpty(self QStackedLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QStackedLayout_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QStackedLayout) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_ControlTypes
func miqt_exec_callback_QStackedLayout_ControlTypes(self QStackedLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

func (this *QStackedLayout) callVirtualBase_ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {

	return newQLayoutItem(QStackedLayout_virtualbase_ReplaceWidget(unsafe.Pointer(this.h), from.cPointer(), to.cPointer(), (int)(options)))

}
func (this *QStackedLayout) OnReplaceWidget(slot func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_ReplaceWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_ReplaceWidget
func miqt_exec_callback_QStackedLayout_ReplaceWidget(self QStackedLayout, cb intptr_t, from *QWidget, to *QWidget, options int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(from)

	slotval2 := newQWidget(to)

	slotval3 := (FindChildOption)(options)

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_ReplaceWidget, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_Layout() *QLayout {

	return newQLayout(QStackedLayout_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QStackedLayout) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_Layout
func miqt_exec_callback_QStackedLayout_Layout(self QStackedLayout, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QStackedLayout) callVirtualBase_ChildEvent(e *QChildEvent) {

	QStackedLayout_virtualbase_ChildEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QStackedLayout) OnChildEvent(slot func(super func(e *QChildEvent), e *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_ChildEvent
func miqt_exec_callback_QStackedLayout_ChildEvent(self QStackedLayout, cb intptr_t, e *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QChildEvent), e *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(e)

	gofunc((&QStackedLayout{h: self}).callVirtualBase_ChildEvent, slotval1)

}
