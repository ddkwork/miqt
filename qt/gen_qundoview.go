package qt

import (
	"unsafe"
)

type QUndoView struct {
	h          uintptr
	isSubclass bool
}

// NewQUndoView constructs a new QUndoView object.
func NewQUndoView(parent *QWidget) *QUndoView {

	ret := newQUndoView(QUndoView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView2 constructs a new QUndoView object.
func NewQUndoView2() *QUndoView {

	ret := newQUndoView(QUndoView_new2())
	ret.isSubclass = true
	return ret
}

// NewQUndoView3 constructs a new QUndoView object.
func NewQUndoView3(stack *QUndoStack) *QUndoView {

	ret := newQUndoView(QUndoView_new3(stack.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView4 constructs a new QUndoView object.
func NewQUndoView4(group *QUndoGroup) *QUndoView {

	ret := newQUndoView(QUndoView_new4(group.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView5 constructs a new QUndoView object.
func NewQUndoView5(stack *QUndoStack, parent *QWidget) *QUndoView {

	ret := newQUndoView(QUndoView_new5(stack.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView6 constructs a new QUndoView object.
func NewQUndoView6(group *QUndoGroup, parent *QWidget) *QUndoView {

	ret := newQUndoView(QUndoView_new6(group.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QUndoView) MetaObject() *QMetaObject {
	return newQMetaObject(QUndoView_MetaObject(this.h))
}

func (this *QUndoView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QUndoView_Metacast(this.h, param1_Cstring))
}

func QUndoView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QUndoView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoView) Stack() *QUndoStack {
	return newQUndoStack(QUndoView_Stack(this.h))
}

func (this *QUndoView) Group() *QUndoGroup {
	return newQUndoGroup(QUndoView_Group(this.h))
}

func (this *QUndoView) SetEmptyLabel(label string) {
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	QUndoView_SetEmptyLabel(this.h, label_ms)
}

func (this *QUndoView) EmptyLabel() string {
	var _ms struct_miqt_string = QUndoView_EmptyLabel(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoView) SetCleanIcon(icon *QIcon) {
	QUndoView_SetCleanIcon(this.h, icon.cPointer())
}

func (this *QUndoView) CleanIcon() *QIcon {
	_goptr := newQIcon(QUndoView_CleanIcon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUndoView) SetStack(stack *QUndoStack) {
	QUndoView_SetStack(this.h, stack.cPointer())
}

func (this *QUndoView) SetGroup(group *QUndoGroup) {
	QUndoView_SetGroup(this.h, group.cPointer())
}

func QUndoView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUndoView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoView) callVirtualBase_VisualRect(index *QModelIndex) *QRect {

	_goptr := newQRect(QUndoView_virtualbase_VisualRect(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QUndoView) OnVisualRect(slot func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_VisualRect
func miqt_exec_callback_QUndoView_VisualRect(self QUndoView, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_VisualRect, slotval1)

	return virtualReturn.cPointer()

}

func (this *QUndoView) callVirtualBase_ScrollTo(index *QModelIndex, hint ScrollHint) {

	QUndoView_virtualbase_ScrollTo(unsafe.Pointer(this.h), index.cPointer(), hint)

}
func (this *QUndoView) OnScrollTo(slot func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_ScrollTo
func miqt_exec_callback_QUndoView_ScrollTo(self QUndoView, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc((&QUndoView{h: self}).callVirtualBase_ScrollTo, slotval1, slotval2)

}

func (this *QUndoView) callVirtualBase_IndexAt(p *QPoint) *QModelIndex {

	_goptr := newQModelIndex(QUndoView_virtualbase_IndexAt(unsafe.Pointer(this.h), p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QUndoView) OnIndexAt(slot func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_IndexAt
func miqt_exec_callback_QUndoView_IndexAt(self QUndoView, cb intptr_t, p *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(p)

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_IndexAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QUndoView) callVirtualBase_DoItemsLayout() {

	QUndoView_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QUndoView) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_DoItemsLayout
func miqt_exec_callback_QUndoView_DoItemsLayout(self QUndoView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUndoView{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QUndoView) callVirtualBase_Reset() {

	QUndoView_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QUndoView) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_Reset
func miqt_exec_callback_QUndoView_Reset(self QUndoView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUndoView{h: self}).callVirtualBase_Reset)

}

func (this *QUndoView) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QUndoView_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QUndoView) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_SetRootIndex
func miqt_exec_callback_QUndoView_SetRootIndex(self QUndoView, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QUndoView{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QUndoView) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QUndoView_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QUndoView) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_Event
func miqt_exec_callback_QUndoView_Event(self QUndoView, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QUndoView) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QUndoView_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QUndoView) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_ScrollContentsBy
func miqt_exec_callback_QUndoView_ScrollContentsBy(self QUndoView, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QUndoView{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QUndoView) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QUndoView_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QUndoView) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_DataChanged
func miqt_exec_callback_QUndoView_DataChanged(self QUndoView, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(topLeft)

	slotval2 := newQModelIndex(bottomRight)

	var roles_ma struct_miqt_array = roles
	roles_ret := make([]int, int(roles_ma.len))
	roles_outCast := (*[0xffff]int)(unsafe.Pointer(roles_ma.data)) // hey ya
	for i := 0; i < int(roles_ma.len); i++ {
		roles_ret[i] = (int)(roles_outCast[i])
	}
	slotval3 := roles_ret

	gofunc((&QUndoView{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QUndoView) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QUndoView_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QUndoView) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_RowsInserted
func miqt_exec_callback_QUndoView_RowsInserted(self QUndoView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QUndoView{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QUndoView) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QUndoView_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QUndoView) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_RowsAboutToBeRemoved
func miqt_exec_callback_QUndoView_RowsAboutToBeRemoved(self QUndoView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QUndoView{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QUndoView) callVirtualBase_MouseMoveEvent(e *QMouseEvent) {

	QUndoView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnMouseMoveEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_MouseMoveEvent
func miqt_exec_callback_QUndoView_MouseMoveEvent(self QUndoView, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QUndoView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_MouseReleaseEvent
func miqt_exec_callback_QUndoView_MouseReleaseEvent(self QUndoView, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QUndoView_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_WheelEvent
func miqt_exec_callback_QUndoView_WheelEvent(self QUndoView, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_TimerEvent(e *QTimerEvent) {

	QUndoView_virtualbase_TimerEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnTimerEvent(slot func(super func(e *QTimerEvent), e *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_TimerEvent
func miqt_exec_callback_QUndoView_TimerEvent(self QUndoView, cb intptr_t, e *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QTimerEvent), e *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_ResizeEvent(e *QResizeEvent) {

	QUndoView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnResizeEvent(slot func(super func(e *QResizeEvent), e *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_ResizeEvent
func miqt_exec_callback_QUndoView_ResizeEvent(self QUndoView, cb intptr_t, e *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QResizeEvent), e *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_DragMoveEvent(e *QDragMoveEvent) {

	QUndoView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnDragMoveEvent(slot func(super func(e *QDragMoveEvent), e *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_DragMoveEvent
func miqt_exec_callback_QUndoView_DragMoveEvent(self QUndoView, cb intptr_t, e *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragMoveEvent), e *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_DragLeaveEvent(e *QDragLeaveEvent) {

	QUndoView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnDragLeaveEvent(slot func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_DragLeaveEvent
func miqt_exec_callback_QUndoView_DragLeaveEvent(self QUndoView, cb intptr_t, e *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_DropEvent(e *QDropEvent) {

	QUndoView_virtualbase_DropEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnDropEvent(slot func(super func(e *QDropEvent), e *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_DropEvent
func miqt_exec_callback_QUndoView_DropEvent(self QUndoView, cb intptr_t, e *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDropEvent), e *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_StartDrag(supportedActions DropAction) {

	QUndoView_virtualbase_StartDrag(unsafe.Pointer(this.h), (int)(supportedActions))

}
func (this *QUndoView) OnStartDrag(slot func(super func(supportedActions DropAction), supportedActions DropAction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_StartDrag(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_StartDrag
func miqt_exec_callback_QUndoView_StartDrag(self QUndoView, cb intptr_t, supportedActions int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(supportedActions DropAction), supportedActions DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(supportedActions)

	gofunc((&QUndoView{h: self}).callVirtualBase_StartDrag, slotval1)

}

func (this *QUndoView) callVirtualBase_InitViewItemOption(option *QStyleOptionViewItem) {

	QUndoView_virtualbase_InitViewItemOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QUndoView) OnInitViewItemOption(slot func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_InitViewItemOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_InitViewItemOption
func miqt_exec_callback_QUndoView_InitViewItemOption(self QUndoView, cb intptr_t, option *QStyleOptionViewItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	gofunc((&QUndoView{h: self}).callVirtualBase_InitViewItemOption, slotval1)

}

func (this *QUndoView) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QUndoView_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QUndoView) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_PaintEvent
func miqt_exec_callback_QUndoView_PaintEvent(self QUndoView, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QUndoView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QUndoView) callVirtualBase_HorizontalOffset() int {

	return (int)(QUndoView_virtualbase_HorizontalOffset(unsafe.Pointer(this.h)))

}
func (this *QUndoView) OnHorizontalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_HorizontalOffset
func miqt_exec_callback_QUndoView_HorizontalOffset(self QUndoView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_HorizontalOffset)

	return (int)(virtualReturn)

}

func (this *QUndoView) callVirtualBase_VerticalOffset() int {

	return (int)(QUndoView_virtualbase_VerticalOffset(unsafe.Pointer(this.h)))

}
func (this *QUndoView) OnVerticalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_VerticalOffset
func miqt_exec_callback_QUndoView_VerticalOffset(self QUndoView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_VerticalOffset)

	return (int)(virtualReturn)

}

func (this *QUndoView) callVirtualBase_MoveCursor(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex {

	_goptr := newQModelIndex(QUndoView_virtualbase_MoveCursor(unsafe.Pointer(this.h), cursorAction, (int)(modifiers)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QUndoView) OnMoveCursor(slot func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_MoveCursor
func miqt_exec_callback_QUndoView_MoveCursor(self QUndoView, cb intptr_t, cursorAction CursorAction, modifiers int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(modifiers)

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_MoveCursor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QUndoView) callVirtualBase_SetSelection(rect *QRect, command SelectionFlag) {

	QUndoView_virtualbase_SetSelection(unsafe.Pointer(this.h), rect.cPointer(), (int)(command))

}
func (this *QUndoView) OnSetSelection(slot func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_SetSelection
func miqt_exec_callback_QUndoView_SetSelection(self QUndoView, cb intptr_t, rect *QRect, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QUndoView{h: self}).callVirtualBase_SetSelection, slotval1, slotval2)

}

func (this *QUndoView) callVirtualBase_VisualRegionForSelection(selection *QItemSelection) *QRegion {

	_goptr := newQRegion(QUndoView_virtualbase_VisualRegionForSelection(unsafe.Pointer(this.h), selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QUndoView) OnVisualRegionForSelection(slot func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_VisualRegionForSelection
func miqt_exec_callback_QUndoView_VisualRegionForSelection(self QUndoView, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_VisualRegionForSelection, slotval1)

	return virtualReturn.cPointer()

}

func (this *QUndoView) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QUndoView_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QUndoView) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_SelectedIndexes
func miqt_exec_callback_QUndoView_SelectedIndexes(self QUndoView, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QUndoView) callVirtualBase_UpdateGeometries() {

	QUndoView_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QUndoView) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_UpdateGeometries
func miqt_exec_callback_QUndoView_UpdateGeometries(self QUndoView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUndoView{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QUndoView) callVirtualBase_IsIndexHidden(index *QModelIndex) bool {

	return (bool)(QUndoView_virtualbase_IsIndexHidden(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QUndoView) OnIsIndexHidden(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_IsIndexHidden
func miqt_exec_callback_QUndoView_IsIndexHidden(self QUndoView, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_IsIndexHidden, slotval1)

	return (bool)(virtualReturn)

}

func (this *QUndoView) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QUndoView_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QUndoView) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_SelectionChanged
func miqt_exec_callback_QUndoView_SelectionChanged(self QUndoView, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QUndoView{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QUndoView) callVirtualBase_CurrentChanged(current *QModelIndex, previous *QModelIndex) {

	QUndoView_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), previous.cPointer())

}
func (this *QUndoView) OnCurrentChanged(slot func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_CurrentChanged
func miqt_exec_callback_QUndoView_CurrentChanged(self QUndoView, cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc((&QUndoView{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}

func (this *QUndoView) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QUndoView_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QUndoView) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_ViewportSizeHint
func miqt_exec_callback_QUndoView_ViewportSizeHint(self QUndoView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}
