package qt

import (
	"unsafe"
)

type QBoxLayout__Direction int

const (
	QBoxLayout__LeftToRight QBoxLayout__Direction = 0
	QBoxLayout__RightToLeft QBoxLayout__Direction = 1
	QBoxLayout__TopToBottom QBoxLayout__Direction = 2
	QBoxLayout__BottomToTop QBoxLayout__Direction = 3
	QBoxLayout__Down        QBoxLayout__Direction = 2
	QBoxLayout__Up          QBoxLayout__Direction = 3
)

type QBoxLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQBoxLayout constructs a new QBoxLayout object.
func NewQBoxLayout(param1 Direction) *QBoxLayout {

	ret := newQBoxLayout(QBoxLayout_new(param1))
	ret.isSubclass = true
	return ret
}

// NewQBoxLayout2 constructs a new QBoxLayout object.
func NewQBoxLayout2(param1 Direction, parent *QWidget) *QBoxLayout {

	ret := newQBoxLayout(QBoxLayout_new2(param1, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QBoxLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QBoxLayout_MetaObject(this.h))
}

func (this *QBoxLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QBoxLayout_Metacast(this.h, param1_Cstring))
}

func QBoxLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QBoxLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QBoxLayout) Direction() Direction {
	xxxxxxxxx
}

func (this *QBoxLayout) SetDirection(direction Direction) {
	QBoxLayout_SetDirection(this.h, direction)
}

func (this *QBoxLayout) AddSpacing(size int) {
	QBoxLayout_AddSpacing(this.h, (int)(size))
}

func (this *QBoxLayout) AddStretch() {
	QBoxLayout_AddStretch(this.h)
}

func (this *QBoxLayout) AddSpacerItem(spacerItem *QSpacerItem) {
	QBoxLayout_AddSpacerItem(this.h, spacerItem.cPointer())
}

func (this *QBoxLayout) AddWidget(param1 *QWidget) {
	QBoxLayout_AddWidget(this.h, param1.cPointer())
}

func (this *QBoxLayout) AddLayout(layout *QLayout) {
	QBoxLayout_AddLayout(this.h, layout.cPointer())
}

func (this *QBoxLayout) AddStrut(param1 int) {
	QBoxLayout_AddStrut(this.h, (int)(param1))
}

func (this *QBoxLayout) AddItem(param1 *QLayoutItem) {
	QBoxLayout_AddItem(this.h, param1.cPointer())
}

func (this *QBoxLayout) InsertSpacing(index int, size int) {
	QBoxLayout_InsertSpacing(this.h, (int)(index), (int)(size))
}

func (this *QBoxLayout) InsertStretch(index int) {
	QBoxLayout_InsertStretch(this.h, (int)(index))
}

func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem *QSpacerItem) {
	QBoxLayout_InsertSpacerItem(this.h, (int)(index), spacerItem.cPointer())
}

func (this *QBoxLayout) InsertWidget(index int, widget *QWidget) {
	QBoxLayout_InsertWidget(this.h, (int)(index), widget.cPointer())
}

func (this *QBoxLayout) InsertLayout(index int, layout *QLayout) {
	QBoxLayout_InsertLayout(this.h, (int)(index), layout.cPointer())
}

func (this *QBoxLayout) InsertItem(index int, param2 *QLayoutItem) {
	QBoxLayout_InsertItem(this.h, (int)(index), param2.cPointer())
}

func (this *QBoxLayout) Spacing() int {
	return (int)(QBoxLayout_Spacing(this.h))
}

func (this *QBoxLayout) SetSpacing(spacing int) {
	QBoxLayout_SetSpacing(this.h, (int)(spacing))
}

func (this *QBoxLayout) SetStretchFactor(w *QWidget, stretch int) bool {
	return (bool)(QBoxLayout_SetStretchFactor(this.h, w.cPointer(), (int)(stretch)))
}

func (this *QBoxLayout) SetStretchFactor2(l *QLayout, stretch int) bool {
	return (bool)(QBoxLayout_SetStretchFactor2(this.h, l.cPointer(), (int)(stretch)))
}

func (this *QBoxLayout) SetStretch(index int, stretch int) {
	QBoxLayout_SetStretch(this.h, (int)(index), (int)(stretch))
}

func (this *QBoxLayout) Stretch(index int) int {
	return (int)(QBoxLayout_Stretch(this.h, (int)(index)))
}

func (this *QBoxLayout) SizeHint() *QSize {
	_goptr := newQSize(QBoxLayout_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBoxLayout) MinimumSize() *QSize {
	_goptr := newQSize(QBoxLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBoxLayout) MaximumSize() *QSize {
	_goptr := newQSize(QBoxLayout_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBoxLayout) HasHeightForWidth() bool {
	return (bool)(QBoxLayout_HasHeightForWidth(this.h))
}

func (this *QBoxLayout) HeightForWidth(param1 int) int {
	return (int)(QBoxLayout_HeightForWidth(this.h, (int)(param1)))
}

func (this *QBoxLayout) MinimumHeightForWidth(param1 int) int {
	return (int)(QBoxLayout_MinimumHeightForWidth(this.h, (int)(param1)))
}

func (this *QBoxLayout) ExpandingDirections() Orientation {
	return (Orientation)(QBoxLayout_ExpandingDirections(this.h))
}

func (this *QBoxLayout) Invalidate() {
	QBoxLayout_Invalidate(this.h)
}

func (this *QBoxLayout) ItemAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QBoxLayout_ItemAt(this.h, (int)(param1)))
}

func (this *QBoxLayout) TakeAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QBoxLayout_TakeAt(this.h, (int)(param1)))
}

func (this *QBoxLayout) Count() int {
	return (int)(QBoxLayout_Count(this.h))
}

func (this *QBoxLayout) SetGeometry(geometry *QRect) {
	QBoxLayout_SetGeometry(this.h, geometry.cPointer())
}

func QBoxLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QBoxLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QBoxLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QBoxLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QBoxLayout) AddStretch1(stretch int) {
	QBoxLayout_AddStretch1(this.h, (int)(stretch))
}

func (this *QBoxLayout) AddWidget2(param1 *QWidget, stretch int) {
	QBoxLayout_AddWidget2(this.h, param1.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) AddWidget3(param1 *QWidget, stretch int, alignment AlignmentFlag) {
	QBoxLayout_AddWidget3(this.h, param1.cPointer(), (int)(stretch), (int)(alignment))
}

func (this *QBoxLayout) AddLayout2(layout *QLayout, stretch int) {
	QBoxLayout_AddLayout2(this.h, layout.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) InsertStretch2(index int, stretch int) {
	QBoxLayout_InsertStretch2(this.h, (int)(index), (int)(stretch))
}

func (this *QBoxLayout) InsertWidget3(index int, widget *QWidget, stretch int) {
	QBoxLayout_InsertWidget3(this.h, (int)(index), widget.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) InsertWidget4(index int, widget *QWidget, stretch int, alignment AlignmentFlag) {
	QBoxLayout_InsertWidget4(this.h, (int)(index), widget.cPointer(), (int)(stretch), (int)(alignment))
}

func (this *QBoxLayout) InsertLayout3(index int, layout *QLayout, stretch int) {
	QBoxLayout_InsertLayout3(this.h, (int)(index), layout.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) callVirtualBase_AddItem(param1 *QLayoutItem) {

	QBoxLayout_virtualbase_AddItem(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QBoxLayout) OnAddItem(slot func(super func(param1 *QLayoutItem), param1 *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_AddItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_AddItem
func miqt_exec_callback_QBoxLayout_AddItem(self QBoxLayout, cb intptr_t, param1 *QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QLayoutItem), param1 *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(param1)

	gofunc((&QBoxLayout{h: self}).callVirtualBase_AddItem, slotval1)

}

func (this *QBoxLayout) callVirtualBase_Spacing() int {

	return (int)(QBoxLayout_virtualbase_Spacing(unsafe.Pointer(this.h)))

}
func (this *QBoxLayout) OnSpacing(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_Spacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_Spacing
func miqt_exec_callback_QBoxLayout_Spacing(self QBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_Spacing)

	return (int)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_SetSpacing(spacing int) {

	QBoxLayout_virtualbase_SetSpacing(unsafe.Pointer(this.h), (int)(spacing))

}
func (this *QBoxLayout) OnSetSpacing(slot func(super func(spacing int), spacing int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_SetSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_SetSpacing
func miqt_exec_callback_QBoxLayout_SetSpacing(self QBoxLayout, cb intptr_t, spacing int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(spacing int), spacing int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(spacing)

	gofunc((&QBoxLayout{h: self}).callVirtualBase_SetSpacing, slotval1)

}

func (this *QBoxLayout) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QBoxLayout_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QBoxLayout) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_SizeHint
func miqt_exec_callback_QBoxLayout_SizeHint(self QBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QBoxLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QBoxLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_MinimumSize
func miqt_exec_callback_QBoxLayout_MinimumSize(self QBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QBoxLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QBoxLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_MaximumSize
func miqt_exec_callback_QBoxLayout_MaximumSize(self QBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QBoxLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QBoxLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_HasHeightForWidth
func miqt_exec_callback_QBoxLayout_HasHeightForWidth(self QBoxLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QBoxLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QBoxLayout) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_HeightForWidth
func miqt_exec_callback_QBoxLayout_HeightForWidth(self QBoxLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QBoxLayout_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QBoxLayout) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_MinimumHeightForWidth
func miqt_exec_callback_QBoxLayout_MinimumHeightForWidth(self QBoxLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QBoxLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QBoxLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_ExpandingDirections
func miqt_exec_callback_QBoxLayout_ExpandingDirections(self QBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_Invalidate() {

	QBoxLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QBoxLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_Invalidate
func miqt_exec_callback_QBoxLayout_Invalidate(self QBoxLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QBoxLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QBoxLayout) callVirtualBase_ItemAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QBoxLayout_virtualbase_ItemAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QBoxLayout) OnItemAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_ItemAt
func miqt_exec_callback_QBoxLayout_ItemAt(self QBoxLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_ItemAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_TakeAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QBoxLayout_virtualbase_TakeAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QBoxLayout) OnTakeAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_TakeAt
func miqt_exec_callback_QBoxLayout_TakeAt(self QBoxLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_TakeAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_Count() int {

	return (int)(QBoxLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QBoxLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_Count
func miqt_exec_callback_QBoxLayout_Count(self QBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_SetGeometry(geometry *QRect) {

	QBoxLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QBoxLayout) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_SetGeometry
func miqt_exec_callback_QBoxLayout_SetGeometry(self QBoxLayout, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QBoxLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QBoxLayout) callVirtualBase_Geometry() *QRect {

	_goptr := newQRect(QBoxLayout_virtualbase_Geometry(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QBoxLayout) OnGeometry(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_Geometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_Geometry
func miqt_exec_callback_QBoxLayout_Geometry(self QBoxLayout, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_Geometry)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_IndexOf(param1 *QWidget) int {

	return (int)(QBoxLayout_virtualbase_IndexOf(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QBoxLayout) OnIndexOf(slot func(super func(param1 *QWidget) int, param1 *QWidget) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_IndexOf(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_IndexOf
func miqt_exec_callback_QBoxLayout_IndexOf(self QBoxLayout, cb intptr_t, param1 *QWidget) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWidget) int, param1 *QWidget) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(param1)

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_IndexOf, slotval1)

	return (int)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_IsEmpty() bool {

	return (bool)(QBoxLayout_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QBoxLayout) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_IsEmpty
func miqt_exec_callback_QBoxLayout_IsEmpty(self QBoxLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_ControlTypes() ControlType {

	return (ControlType)(QBoxLayout_virtualbase_ControlTypes(unsafe.Pointer(this.h)))

}
func (this *QBoxLayout) OnControlTypes(slot func(super func() ControlType) ControlType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_ControlTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_ControlTypes
func miqt_exec_callback_QBoxLayout_ControlTypes(self QBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() ControlType) ControlType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_ControlTypes)

	return (int)(virtualReturn)

}

func (this *QBoxLayout) callVirtualBase_ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {

	return newQLayoutItem(QBoxLayout_virtualbase_ReplaceWidget(unsafe.Pointer(this.h), from.cPointer(), to.cPointer(), (int)(options)))

}
func (this *QBoxLayout) OnReplaceWidget(slot func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_ReplaceWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_ReplaceWidget
func miqt_exec_callback_QBoxLayout_ReplaceWidget(self QBoxLayout, cb intptr_t, from *QWidget, to *QWidget, options int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem, from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(from)

	slotval2 := newQWidget(to)

	slotval3 := (FindChildOption)(options)

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_ReplaceWidget, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_Layout() *QLayout {

	return newQLayout(QBoxLayout_virtualbase_Layout(unsafe.Pointer(this.h)))

}
func (this *QBoxLayout) OnLayout(slot func(super func() *QLayout) *QLayout) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_Layout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_Layout
func miqt_exec_callback_QBoxLayout_Layout(self QBoxLayout, cb intptr_t) *QLayout {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QLayout) *QLayout)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_Layout)

	return virtualReturn.cPointer()

}

func (this *QBoxLayout) callVirtualBase_ChildEvent(e *QChildEvent) {

	QBoxLayout_virtualbase_ChildEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QBoxLayout) OnChildEvent(slot func(super func(e *QChildEvent), e *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_ChildEvent
func miqt_exec_callback_QBoxLayout_ChildEvent(self QBoxLayout, cb intptr_t, e *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QChildEvent), e *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(e)

	gofunc((&QBoxLayout{h: self}).callVirtualBase_ChildEvent, slotval1)

}

type QHBoxLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQHBoxLayout constructs a new QHBoxLayout object.
func NewQHBoxLayout(parent *QWidget) *QHBoxLayout {

	ret := newQHBoxLayout(QHBoxLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQHBoxLayout2 constructs a new QHBoxLayout object.
func NewQHBoxLayout2() *QHBoxLayout {

	ret := newQHBoxLayout(QHBoxLayout_new2())
	ret.isSubclass = true
	return ret
}

func (this *QHBoxLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QHBoxLayout_MetaObject(this.h))
}

func (this *QHBoxLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QHBoxLayout_Metacast(this.h, param1_Cstring))
}

func QHBoxLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QHBoxLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QHBoxLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHBoxLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QHBoxLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHBoxLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHBoxLayout) callVirtualBase_AddItem(param1 *QLayoutItem) {

	QHBoxLayout_virtualbase_AddItem(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QHBoxLayout) OnAddItem(slot func(super func(param1 *QLayoutItem), param1 *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_AddItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_AddItem
func miqt_exec_callback_QHBoxLayout_AddItem(self QHBoxLayout, cb intptr_t, param1 *QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QLayoutItem), param1 *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(param1)

	gofunc((&QHBoxLayout{h: self}).callVirtualBase_AddItem, slotval1)

}

func (this *QHBoxLayout) callVirtualBase_Spacing() int {

	return (int)(QHBoxLayout_virtualbase_Spacing(unsafe.Pointer(this.h)))

}
func (this *QHBoxLayout) OnSpacing(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_Spacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_Spacing
func miqt_exec_callback_QHBoxLayout_Spacing(self QHBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_Spacing)

	return (int)(virtualReturn)

}

func (this *QHBoxLayout) callVirtualBase_SetSpacing(spacing int) {

	QHBoxLayout_virtualbase_SetSpacing(unsafe.Pointer(this.h), (int)(spacing))

}
func (this *QHBoxLayout) OnSetSpacing(slot func(super func(spacing int), spacing int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_SetSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_SetSpacing
func miqt_exec_callback_QHBoxLayout_SetSpacing(self QHBoxLayout, cb intptr_t, spacing int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(spacing int), spacing int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(spacing)

	gofunc((&QHBoxLayout{h: self}).callVirtualBase_SetSpacing, slotval1)

}

func (this *QHBoxLayout) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QHBoxLayout_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHBoxLayout) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_SizeHint
func miqt_exec_callback_QHBoxLayout_SizeHint(self QHBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QHBoxLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QHBoxLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHBoxLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_MinimumSize
func miqt_exec_callback_QHBoxLayout_MinimumSize(self QHBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QHBoxLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QHBoxLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHBoxLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_MaximumSize
func miqt_exec_callback_QHBoxLayout_MaximumSize(self QHBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QHBoxLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QHBoxLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QHBoxLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_HasHeightForWidth
func miqt_exec_callback_QHBoxLayout_HasHeightForWidth(self QHBoxLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QHBoxLayout) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QHBoxLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QHBoxLayout) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_HeightForWidth
func miqt_exec_callback_QHBoxLayout_HeightForWidth(self QHBoxLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QHBoxLayout) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QHBoxLayout_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QHBoxLayout) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_MinimumHeightForWidth
func miqt_exec_callback_QHBoxLayout_MinimumHeightForWidth(self QHBoxLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QHBoxLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QHBoxLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QHBoxLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_ExpandingDirections
func miqt_exec_callback_QHBoxLayout_ExpandingDirections(self QHBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QHBoxLayout) callVirtualBase_Invalidate() {

	QHBoxLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QHBoxLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_Invalidate
func miqt_exec_callback_QHBoxLayout_Invalidate(self QHBoxLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QHBoxLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QHBoxLayout) callVirtualBase_ItemAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QHBoxLayout_virtualbase_ItemAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QHBoxLayout) OnItemAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_ItemAt
func miqt_exec_callback_QHBoxLayout_ItemAt(self QHBoxLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_ItemAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHBoxLayout) callVirtualBase_TakeAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QHBoxLayout_virtualbase_TakeAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QHBoxLayout) OnTakeAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_TakeAt
func miqt_exec_callback_QHBoxLayout_TakeAt(self QHBoxLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_TakeAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHBoxLayout) callVirtualBase_Count() int {

	return (int)(QHBoxLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QHBoxLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_Count
func miqt_exec_callback_QHBoxLayout_Count(self QHBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QHBoxLayout) callVirtualBase_SetGeometry(geometry *QRect) {

	QHBoxLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QHBoxLayout) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_SetGeometry
func miqt_exec_callback_QHBoxLayout_SetGeometry(self QHBoxLayout, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QHBoxLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}

type QVBoxLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQVBoxLayout constructs a new QVBoxLayout object.
func NewQVBoxLayout(parent *QWidget) *QVBoxLayout {

	ret := newQVBoxLayout(QVBoxLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVBoxLayout2 constructs a new QVBoxLayout object.
func NewQVBoxLayout2() *QVBoxLayout {

	ret := newQVBoxLayout(QVBoxLayout_new2())
	ret.isSubclass = true
	return ret
}

func (this *QVBoxLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QVBoxLayout_MetaObject(this.h))
}

func (this *QVBoxLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QVBoxLayout_Metacast(this.h, param1_Cstring))
}

func QVBoxLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QVBoxLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVBoxLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVBoxLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVBoxLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVBoxLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVBoxLayout) callVirtualBase_AddItem(param1 *QLayoutItem) {

	QVBoxLayout_virtualbase_AddItem(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QVBoxLayout) OnAddItem(slot func(super func(param1 *QLayoutItem), param1 *QLayoutItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_AddItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_AddItem
func miqt_exec_callback_QVBoxLayout_AddItem(self QVBoxLayout, cb intptr_t, param1 *QLayoutItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QLayoutItem), param1 *QLayoutItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLayoutItem(param1)

	gofunc((&QVBoxLayout{h: self}).callVirtualBase_AddItem, slotval1)

}

func (this *QVBoxLayout) callVirtualBase_Spacing() int {

	return (int)(QVBoxLayout_virtualbase_Spacing(unsafe.Pointer(this.h)))

}
func (this *QVBoxLayout) OnSpacing(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_Spacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_Spacing
func miqt_exec_callback_QVBoxLayout_Spacing(self QVBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_Spacing)

	return (int)(virtualReturn)

}

func (this *QVBoxLayout) callVirtualBase_SetSpacing(spacing int) {

	QVBoxLayout_virtualbase_SetSpacing(unsafe.Pointer(this.h), (int)(spacing))

}
func (this *QVBoxLayout) OnSetSpacing(slot func(super func(spacing int), spacing int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_SetSpacing(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_SetSpacing
func miqt_exec_callback_QVBoxLayout_SetSpacing(self QVBoxLayout, cb intptr_t, spacing int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(spacing int), spacing int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(spacing)

	gofunc((&QVBoxLayout{h: self}).callVirtualBase_SetSpacing, slotval1)

}

func (this *QVBoxLayout) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QVBoxLayout_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QVBoxLayout) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_SizeHint
func miqt_exec_callback_QVBoxLayout_SizeHint(self QVBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QVBoxLayout) callVirtualBase_MinimumSize() *QSize {

	_goptr := newQSize(QVBoxLayout_virtualbase_MinimumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QVBoxLayout) OnMinimumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_MinimumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_MinimumSize
func miqt_exec_callback_QVBoxLayout_MinimumSize(self QVBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_MinimumSize)

	return virtualReturn.cPointer()

}

func (this *QVBoxLayout) callVirtualBase_MaximumSize() *QSize {

	_goptr := newQSize(QVBoxLayout_virtualbase_MaximumSize(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QVBoxLayout) OnMaximumSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_MaximumSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_MaximumSize
func miqt_exec_callback_QVBoxLayout_MaximumSize(self QVBoxLayout, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_MaximumSize)

	return virtualReturn.cPointer()

}

func (this *QVBoxLayout) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QVBoxLayout_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QVBoxLayout) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_HasHeightForWidth
func miqt_exec_callback_QVBoxLayout_HasHeightForWidth(self QVBoxLayout, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QVBoxLayout) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QVBoxLayout_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QVBoxLayout) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_HeightForWidth
func miqt_exec_callback_QVBoxLayout_HeightForWidth(self QVBoxLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QVBoxLayout) callVirtualBase_MinimumHeightForWidth(param1 int) int {

	return (int)(QVBoxLayout_virtualbase_MinimumHeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QVBoxLayout) OnMinimumHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_MinimumHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_MinimumHeightForWidth
func miqt_exec_callback_QVBoxLayout_MinimumHeightForWidth(self QVBoxLayout, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_MinimumHeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QVBoxLayout) callVirtualBase_ExpandingDirections() Orientation {

	return (Orientation)(QVBoxLayout_virtualbase_ExpandingDirections(unsafe.Pointer(this.h)))

}
func (this *QVBoxLayout) OnExpandingDirections(slot func(super func() Orientation) Orientation) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_ExpandingDirections(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_ExpandingDirections
func miqt_exec_callback_QVBoxLayout_ExpandingDirections(self QVBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Orientation) Orientation)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_ExpandingDirections)

	return (int)(virtualReturn)

}

func (this *QVBoxLayout) callVirtualBase_Invalidate() {

	QVBoxLayout_virtualbase_Invalidate(unsafe.Pointer(this.h))

}
func (this *QVBoxLayout) OnInvalidate(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_Invalidate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_Invalidate
func miqt_exec_callback_QVBoxLayout_Invalidate(self QVBoxLayout, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QVBoxLayout{h: self}).callVirtualBase_Invalidate)

}

func (this *QVBoxLayout) callVirtualBase_ItemAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QVBoxLayout_virtualbase_ItemAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QVBoxLayout) OnItemAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_ItemAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_ItemAt
func miqt_exec_callback_QVBoxLayout_ItemAt(self QVBoxLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_ItemAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QVBoxLayout) callVirtualBase_TakeAt(param1 int) *QLayoutItem {

	return newQLayoutItem(QVBoxLayout_virtualbase_TakeAt(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QVBoxLayout) OnTakeAt(slot func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_TakeAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_TakeAt
func miqt_exec_callback_QVBoxLayout_TakeAt(self QVBoxLayout, cb intptr_t, param1 int) *QLayoutItem {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) *QLayoutItem, param1 int) *QLayoutItem)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_TakeAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QVBoxLayout) callVirtualBase_Count() int {

	return (int)(QVBoxLayout_virtualbase_Count(unsafe.Pointer(this.h)))

}
func (this *QVBoxLayout) OnCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_Count(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_Count
func miqt_exec_callback_QVBoxLayout_Count(self QVBoxLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_Count)

	return (int)(virtualReturn)

}

func (this *QVBoxLayout) callVirtualBase_SetGeometry(geometry *QRect) {

	QVBoxLayout_virtualbase_SetGeometry(unsafe.Pointer(this.h), geometry.cPointer())

}
func (this *QVBoxLayout) OnSetGeometry(slot func(super func(geometry *QRect), geometry *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_SetGeometry
func miqt_exec_callback_QVBoxLayout_SetGeometry(self QVBoxLayout, cb intptr_t, geometry *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(geometry *QRect), geometry *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(geometry)

	gofunc((&QVBoxLayout{h: self}).callVirtualBase_SetGeometry, slotval1)

}
