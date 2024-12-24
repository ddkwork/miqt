package qt

import (
	"unsafe"
)

type QTextStream struct {
	h          uintptr
	isSubclass bool
}

type QSplitter struct {
	h          uintptr
	isSubclass bool
}

// NewQSplitter constructs a new QSplitter object.
func NewQSplitter(parent *QWidget) *QSplitter {

	ret := newQSplitter(QSplitter_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSplitter2 constructs a new QSplitter object.
func NewQSplitter2() *QSplitter {

	ret := newQSplitter(QSplitter_new2())
	ret.isSubclass = true
	return ret
}

// NewQSplitter3 constructs a new QSplitter object.
func NewQSplitter3(param1 Orientation) *QSplitter {

	ret := newQSplitter(QSplitter_new3((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQSplitter4 constructs a new QSplitter object.
func NewQSplitter4(param1 Orientation, parent *QWidget) *QSplitter {

	ret := newQSplitter(QSplitter_new4((int)(param1), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSplitter) MetaObject() *QMetaObject {
	return newQMetaObject(QSplitter_MetaObject(this.h))
}

func (this *QSplitter) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSplitter_Metacast(this.h, param1_Cstring))
}

func QSplitter_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSplitter_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitter) AddWidget(widget *QWidget) {
	QSplitter_AddWidget(this.h, widget.cPointer())
}

func (this *QSplitter) InsertWidget(index int, widget *QWidget) {
	QSplitter_InsertWidget(this.h, (int)(index), widget.cPointer())
}

func (this *QSplitter) ReplaceWidget(index int, widget *QWidget) *QWidget {
	return newQWidget(QSplitter_ReplaceWidget(this.h, (int)(index), widget.cPointer()))
}

func (this *QSplitter) SetOrientation(orientation Orientation) {
	QSplitter_SetOrientation(this.h, (int)(orientation))
}

func (this *QSplitter) Orientation() Orientation {
	return (Orientation)(QSplitter_Orientation(this.h))
}

func (this *QSplitter) SetChildrenCollapsible(childrenCollapsible bool) {
	QSplitter_SetChildrenCollapsible(this.h, (bool)(childrenCollapsible))
}

func (this *QSplitter) ChildrenCollapsible() bool {
	return (bool)(QSplitter_ChildrenCollapsible(this.h))
}

func (this *QSplitter) SetCollapsible(index int, param2 bool) {
	QSplitter_SetCollapsible(this.h, (int)(index), (bool)(param2))
}

func (this *QSplitter) IsCollapsible(index int) bool {
	return (bool)(QSplitter_IsCollapsible(this.h, (int)(index)))
}

func (this *QSplitter) SetOpaqueResize() {
	QSplitter_SetOpaqueResize(this.h)
}

func (this *QSplitter) OpaqueResize() bool {
	return (bool)(QSplitter_OpaqueResize(this.h))
}

func (this *QSplitter) Refresh() {
	QSplitter_Refresh(this.h)
}

func (this *QSplitter) SizeHint() *QSize {
	_goptr := newQSize(QSplitter_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSplitter) MinimumSizeHint() *QSize {
	_goptr := newQSize(QSplitter_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSplitter) Sizes() []int {
	var _ma struct_miqt_array = QSplitter_Sizes(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QSplitter) SetSizes(list []int) {
	list_CArray := (*[0xffff]int)(malloc(size_t(8 * len(list))))
	defer free(unsafe.Pointer(list_CArray))
	for i := range list {
		list_CArray[i] = (int)(list[i])
	}
	list_ma := struct_miqt_array{len: size_t(len(list)), data: unsafe.Pointer(list_CArray)}
	QSplitter_SetSizes(this.h, list_ma)
}

func (this *QSplitter) SaveState() []byte {
	var _bytearray struct_miqt_string = QSplitter_SaveState(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSplitter) RestoreState(state []byte) bool {
	state_alias := struct_miqt_string{}
	state_alias.data = (char)(unsafe.Pointer(&state[0]))
	state_alias.len = size_t(len(state))
	return (bool)(QSplitter_RestoreState(this.h, state_alias))
}

func (this *QSplitter) HandleWidth() int {
	return (int)(QSplitter_HandleWidth(this.h))
}

func (this *QSplitter) SetHandleWidth(handleWidth int) {
	QSplitter_SetHandleWidth(this.h, (int)(handleWidth))
}

func (this *QSplitter) IndexOf(w *QWidget) int {
	return (int)(QSplitter_IndexOf(this.h, w.cPointer()))
}

func (this *QSplitter) Widget(index int) *QWidget {
	return newQWidget(QSplitter_Widget(this.h, (int)(index)))
}

func (this *QSplitter) Count() int {
	return (int)(QSplitter_Count(this.h))
}

func (this *QSplitter) GetRange(index int, param2 *int, param3 *int) {
	QSplitter_GetRange(this.h, (int)(index), (*int)(unsafe.Pointer(param2)), (*int)(unsafe.Pointer(param3)))
}

func (this *QSplitter) Handle(index int) *QSplitterHandle {
	return newQSplitterHandle(QSplitter_Handle(this.h, (int)(index)))
}

func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	QSplitter_SetStretchFactor(this.h, (int)(index), (int)(stretch))
}

func (this *QSplitter) SplitterMoved(pos int, index int) {
	QSplitter_SplitterMoved(this.h, (int)(pos), (int)(index))
}
func (this *QSplitter) OnSplitterMoved(slot func(pos int, index int)) {
	QSplitter_connect_SplitterMoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_SplitterMoved
func miqt_exec_callback_QSplitter_SplitterMoved(cb intptr_t, pos int, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(pos int, index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(pos)

	slotval2 := (int)(index)

	gofunc(slotval1, slotval2)
}

func QSplitter_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitter_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSplitter_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitter_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitter) SetOpaqueResize1(opaque bool) {
	QSplitter_SetOpaqueResize1(this.h, (bool)(opaque))
}

func (this *QSplitter) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QSplitter_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSplitter) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_SizeHint
func miqt_exec_callback_QSplitter_SizeHint(self QSplitter, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitter{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QSplitter) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QSplitter_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSplitter) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_MinimumSizeHint
func miqt_exec_callback_QSplitter_MinimumSizeHint(self QSplitter, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitter{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QSplitter) callVirtualBase_CreateHandle() *QSplitterHandle {

	return newQSplitterHandle(QSplitter_virtualbase_CreateHandle(unsafe.Pointer(this.h)))

}
func (this *QSplitter) OnCreateHandle(slot func(super func() *QSplitterHandle) *QSplitterHandle) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_CreateHandle(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_CreateHandle
func miqt_exec_callback_QSplitter_CreateHandle(self QSplitter, cb intptr_t) *QSplitterHandle {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSplitterHandle) *QSplitterHandle)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitter{h: self}).callVirtualBase_CreateHandle)

	return virtualReturn.cPointer()

}

func (this *QSplitter) callVirtualBase_ChildEvent(param1 *QChildEvent) {

	QSplitter_virtualbase_ChildEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitter) OnChildEvent(slot func(super func(param1 *QChildEvent), param1 *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_ChildEvent
func miqt_exec_callback_QSplitter_ChildEvent(self QSplitter, cb intptr_t, param1 *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QChildEvent), param1 *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(param1)

	gofunc((&QSplitter{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSplitter) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QSplitter_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QSplitter) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_Event
func miqt_exec_callback_QSplitter_Event(self QSplitter, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QSplitter{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSplitter) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QSplitter_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitter) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_ResizeEvent
func miqt_exec_callback_QSplitter_ResizeEvent(self QSplitter, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QSplitter{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QSplitter) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QSplitter_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitter) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_ChangeEvent
func miqt_exec_callback_QSplitter_ChangeEvent(self QSplitter, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QSplitter{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QSplitter) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QSplitter_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitter) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_PaintEvent
func miqt_exec_callback_QSplitter_PaintEvent(self QSplitter, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QSplitter{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QSplitter) callVirtualBase_InitStyleOption(option *QStyleOptionFrame) {

	QSplitter_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QSplitter) OnInitStyleOption(slot func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_InitStyleOption
func miqt_exec_callback_QSplitter_InitStyleOption(self QSplitter, cb intptr_t, option *QStyleOptionFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionFrame(option)

	gofunc((&QSplitter{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

type QSplitterHandle struct {
	h          uintptr
	isSubclass bool
}

// NewQSplitterHandle constructs a new QSplitterHandle object.
func NewQSplitterHandle(o Orientation, parent *QSplitter) *QSplitterHandle {

	ret := newQSplitterHandle(QSplitterHandle_new((int)(o), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSplitterHandle) MetaObject() *QMetaObject {
	return newQMetaObject(QSplitterHandle_MetaObject(this.h))
}

func (this *QSplitterHandle) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSplitterHandle_Metacast(this.h, param1_Cstring))
}

func QSplitterHandle_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSplitterHandle_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitterHandle) SetOrientation(o Orientation) {
	QSplitterHandle_SetOrientation(this.h, (int)(o))
}

func (this *QSplitterHandle) Orientation() Orientation {
	return (Orientation)(QSplitterHandle_Orientation(this.h))
}

func (this *QSplitterHandle) OpaqueResize() bool {
	return (bool)(QSplitterHandle_OpaqueResize(this.h))
}

func (this *QSplitterHandle) Splitter() *QSplitter {
	return newQSplitter(QSplitterHandle_Splitter(this.h))
}

func (this *QSplitterHandle) SizeHint() *QSize {
	_goptr := newQSize(QSplitterHandle_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSplitterHandle_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitterHandle_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSplitterHandle_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitterHandle_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitterHandle) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QSplitterHandle_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSplitterHandle) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_SizeHint
func miqt_exec_callback_QSplitterHandle_SizeHint(self QSplitterHandle, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QSplitterHandle) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QSplitterHandle_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitterHandle) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_PaintEvent
func miqt_exec_callback_QSplitterHandle_PaintEvent(self QSplitterHandle, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QSplitterHandle_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitterHandle) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_MouseMoveEvent
func miqt_exec_callback_QSplitterHandle_MouseMoveEvent(self QSplitterHandle, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QSplitterHandle_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitterHandle) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_MousePressEvent
func miqt_exec_callback_QSplitterHandle_MousePressEvent(self QSplitterHandle, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QSplitterHandle_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitterHandle) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_MouseReleaseEvent
func miqt_exec_callback_QSplitterHandle_MouseReleaseEvent(self QSplitterHandle, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QSplitterHandle_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitterHandle) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_ResizeEvent
func miqt_exec_callback_QSplitterHandle_ResizeEvent(self QSplitterHandle, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QSplitterHandle_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QSplitterHandle) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_Event
func miqt_exec_callback_QSplitterHandle_Event(self QSplitterHandle, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSplitterHandle) callVirtualBase_DevType() int {

	return (int)(QSplitterHandle_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QSplitterHandle) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_DevType
func miqt_exec_callback_QSplitterHandle_DevType(self QSplitterHandle, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QSplitterHandle) callVirtualBase_SetVisible(visible bool) {

	QSplitterHandle_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QSplitterHandle) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_SetVisible
func miqt_exec_callback_QSplitterHandle_SetVisible(self QSplitterHandle, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QSplitterHandle_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSplitterHandle) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_MinimumSizeHint
func miqt_exec_callback_QSplitterHandle_MinimumSizeHint(self QSplitterHandle, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QSplitterHandle) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QSplitterHandle_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QSplitterHandle) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_HeightForWidth
func miqt_exec_callback_QSplitterHandle_HeightForWidth(self QSplitterHandle, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QSplitterHandle) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QSplitterHandle_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QSplitterHandle) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_HasHeightForWidth
func miqt_exec_callback_QSplitterHandle_HasHeightForWidth(self QSplitterHandle, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QSplitterHandle) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QSplitterHandle_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QSplitterHandle) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_PaintEngine
func miqt_exec_callback_QSplitterHandle_PaintEngine(self QSplitterHandle, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QSplitterHandle) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QSplitterHandle_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_MouseDoubleClickEvent
func miqt_exec_callback_QSplitterHandle_MouseDoubleClickEvent(self QSplitterHandle, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QSplitterHandle_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_WheelEvent
func miqt_exec_callback_QSplitterHandle_WheelEvent(self QSplitterHandle, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QSplitterHandle_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_KeyPressEvent
func miqt_exec_callback_QSplitterHandle_KeyPressEvent(self QSplitterHandle, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QSplitterHandle_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_KeyReleaseEvent
func miqt_exec_callback_QSplitterHandle_KeyReleaseEvent(self QSplitterHandle, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QSplitterHandle_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_FocusInEvent
func miqt_exec_callback_QSplitterHandle_FocusInEvent(self QSplitterHandle, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QSplitterHandle_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_FocusOutEvent
func miqt_exec_callback_QSplitterHandle_FocusOutEvent(self QSplitterHandle, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QSplitterHandle_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_EnterEvent
func miqt_exec_callback_QSplitterHandle_EnterEvent(self QSplitterHandle, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_LeaveEvent(event *QEvent) {

	QSplitterHandle_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_LeaveEvent
func miqt_exec_callback_QSplitterHandle_LeaveEvent(self QSplitterHandle, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QSplitterHandle_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_MoveEvent
func miqt_exec_callback_QSplitterHandle_MoveEvent(self QSplitterHandle, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QSplitterHandle_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_CloseEvent
func miqt_exec_callback_QSplitterHandle_CloseEvent(self QSplitterHandle, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QSplitterHandle_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_ContextMenuEvent
func miqt_exec_callback_QSplitterHandle_ContextMenuEvent(self QSplitterHandle, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QSplitterHandle_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_TabletEvent
func miqt_exec_callback_QSplitterHandle_TabletEvent(self QSplitterHandle, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_ActionEvent(event *QActionEvent) {

	QSplitterHandle_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_ActionEvent
func miqt_exec_callback_QSplitterHandle_ActionEvent(self QSplitterHandle, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QSplitterHandle_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_DragEnterEvent
func miqt_exec_callback_QSplitterHandle_DragEnterEvent(self QSplitterHandle, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QSplitterHandle_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_DragMoveEvent
func miqt_exec_callback_QSplitterHandle_DragMoveEvent(self QSplitterHandle, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QSplitterHandle_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_DragLeaveEvent
func miqt_exec_callback_QSplitterHandle_DragLeaveEvent(self QSplitterHandle, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_DropEvent(event *QDropEvent) {

	QSplitterHandle_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_DropEvent
func miqt_exec_callback_QSplitterHandle_DropEvent(self QSplitterHandle, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_ShowEvent(event *QShowEvent) {

	QSplitterHandle_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_ShowEvent
func miqt_exec_callback_QSplitterHandle_ShowEvent(self QSplitterHandle, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_HideEvent(event *QHideEvent) {

	QSplitterHandle_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSplitterHandle) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_HideEvent
func miqt_exec_callback_QSplitterHandle_HideEvent(self QSplitterHandle, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QSplitterHandle_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QSplitterHandle) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_NativeEvent
func miqt_exec_callback_QSplitterHandle_NativeEvent(self QSplitterHandle, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QSplitterHandle) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QSplitterHandle_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitterHandle) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_ChangeEvent
func miqt_exec_callback_QSplitterHandle_ChangeEvent(self QSplitterHandle, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QSplitterHandle_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QSplitterHandle) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_Metric
func miqt_exec_callback_QSplitterHandle_Metric(self QSplitterHandle, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QSplitterHandle) callVirtualBase_InitPainter(painter *QPainter) {

	QSplitterHandle_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QSplitterHandle) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_InitPainter
func miqt_exec_callback_QSplitterHandle_InitPainter(self QSplitterHandle, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QSplitterHandle_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QSplitterHandle) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_Redirected
func miqt_exec_callback_QSplitterHandle_Redirected(self QSplitterHandle, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QSplitterHandle) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QSplitterHandle_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QSplitterHandle) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_SharedPainter
func miqt_exec_callback_QSplitterHandle_SharedPainter(self QSplitterHandle, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QSplitterHandle) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QSplitterHandle_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QSplitterHandle) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_InputMethodEvent
func miqt_exec_callback_QSplitterHandle_InputMethodEvent(self QSplitterHandle, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QSplitterHandle{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QSplitterHandle) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QSplitterHandle_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSplitterHandle) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_InputMethodQuery
func miqt_exec_callback_QSplitterHandle_InputMethodQuery(self QSplitterHandle, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QSplitterHandle) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QSplitterHandle_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QSplitterHandle) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_FocusNextPrevChild
func miqt_exec_callback_QSplitterHandle_FocusNextPrevChild(self QSplitterHandle, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
