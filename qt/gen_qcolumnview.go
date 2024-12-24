package qt

import (
	"unsafe"
)

type QColumnView struct {
	h          uintptr
	isSubclass bool
}

// NewQColumnView constructs a new QColumnView object.
func NewQColumnView(parent *QWidget) *QColumnView {

	ret := newQColumnView(QColumnView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQColumnView2 constructs a new QColumnView object.
func NewQColumnView2() *QColumnView {

	ret := newQColumnView(QColumnView_new2())
	ret.isSubclass = true
	return ret
}

func (this *QColumnView) MetaObject() *QMetaObject {
	return newQMetaObject(QColumnView_MetaObject(this.h))
}

func (this *QColumnView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QColumnView_Metacast(this.h, param1_Cstring))
}

func QColumnView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QColumnView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QColumnView) UpdatePreviewWidget(index *QModelIndex) {
	QColumnView_UpdatePreviewWidget(this.h, index.cPointer())
}
func (this *QColumnView) OnUpdatePreviewWidget(slot func(index *QModelIndex)) {
	QColumnView_connect_UpdatePreviewWidget(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_UpdatePreviewWidget
func miqt_exec_callback_QColumnView_UpdatePreviewWidget(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QColumnView) IndexAt(point *QPoint) *QModelIndex {
	_goptr := newQModelIndex(QColumnView_IndexAt(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColumnView) ScrollTo(index *QModelIndex, hint ScrollHint) {
	QColumnView_ScrollTo(this.h, index.cPointer(), hint)
}

func (this *QColumnView) SizeHint() *QSize {
	_goptr := newQSize(QColumnView_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColumnView) VisualRect(index *QModelIndex) *QRect {
	_goptr := newQRect(QColumnView_VisualRect(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColumnView) SetModel(model *QAbstractItemModel) {
	QColumnView_SetModel(this.h, model.cPointer())
}

func (this *QColumnView) SetSelectionModel(selectionModel *QItemSelectionModel) {
	QColumnView_SetSelectionModel(this.h, selectionModel.cPointer())
}

func (this *QColumnView) SetRootIndex(index *QModelIndex) {
	QColumnView_SetRootIndex(this.h, index.cPointer())
}

func (this *QColumnView) SelectAll() {
	QColumnView_SelectAll(this.h)
}

func (this *QColumnView) SetResizeGripsVisible(visible bool) {
	QColumnView_SetResizeGripsVisible(this.h, (bool)(visible))
}

func (this *QColumnView) ResizeGripsVisible() bool {
	return (bool)(QColumnView_ResizeGripsVisible(this.h))
}

func (this *QColumnView) PreviewWidget() *QWidget {
	return newQWidget(QColumnView_PreviewWidget(this.h))
}

func (this *QColumnView) SetPreviewWidget(widget *QWidget) {
	QColumnView_SetPreviewWidget(this.h, widget.cPointer())
}

func (this *QColumnView) SetColumnWidths(list []int) {
	list_CArray := (*[0xffff]int)(malloc(size_t(8 * len(list))))
	defer free(unsafe.Pointer(list_CArray))
	for i := range list {
		list_CArray[i] = (int)(list[i])
	}
	list_ma := struct_miqt_array{len: size_t(len(list)), data: unsafe.Pointer(list_CArray)}
	QColumnView_SetColumnWidths(this.h, list_ma)
}

func (this *QColumnView) ColumnWidths() []int {
	var _ma struct_miqt_array = QColumnView_ColumnWidths(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func QColumnView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QColumnView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QColumnView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QColumnView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QColumnView) callVirtualBase_IndexAt(point *QPoint) *QModelIndex {

	_goptr := newQModelIndex(QColumnView_virtualbase_IndexAt(unsafe.Pointer(this.h), point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColumnView) OnIndexAt(slot func(super func(point *QPoint) *QModelIndex, point *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_IndexAt
func miqt_exec_callback_QColumnView_IndexAt(self QColumnView, cb intptr_t, point *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(point *QPoint) *QModelIndex, point *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(point)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_IndexAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_ScrollTo(index *QModelIndex, hint ScrollHint) {

	QColumnView_virtualbase_ScrollTo(unsafe.Pointer(this.h), index.cPointer(), hint)

}
func (this *QColumnView) OnScrollTo(slot func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_ScrollTo
func miqt_exec_callback_QColumnView_ScrollTo(self QColumnView, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc((&QColumnView{h: self}).callVirtualBase_ScrollTo, slotval1, slotval2)

}

func (this *QColumnView) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QColumnView_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColumnView) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SizeHint
func miqt_exec_callback_QColumnView_SizeHint(self QColumnView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_VisualRect(index *QModelIndex) *QRect {

	_goptr := newQRect(QColumnView_virtualbase_VisualRect(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColumnView) OnVisualRect(slot func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_VisualRect
func miqt_exec_callback_QColumnView_VisualRect(self QColumnView, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_VisualRect, slotval1)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_SetModel(model *QAbstractItemModel) {

	QColumnView_virtualbase_SetModel(unsafe.Pointer(this.h), model.cPointer())

}
func (this *QColumnView) OnSetModel(slot func(super func(model *QAbstractItemModel), model *QAbstractItemModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SetModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SetModel
func miqt_exec_callback_QColumnView_SetModel(self QColumnView, cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(model *QAbstractItemModel), model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc((&QColumnView{h: self}).callVirtualBase_SetModel, slotval1)

}

func (this *QColumnView) callVirtualBase_SetSelectionModel(selectionModel *QItemSelectionModel) {

	QColumnView_virtualbase_SetSelectionModel(unsafe.Pointer(this.h), selectionModel.cPointer())

}
func (this *QColumnView) OnSetSelectionModel(slot func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SetSelectionModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SetSelectionModel
func miqt_exec_callback_QColumnView_SetSelectionModel(self QColumnView, cb intptr_t, selectionModel *QItemSelectionModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelectionModel(selectionModel)

	gofunc((&QColumnView{h: self}).callVirtualBase_SetSelectionModel, slotval1)

}

func (this *QColumnView) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QColumnView_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QColumnView) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SetRootIndex
func miqt_exec_callback_QColumnView_SetRootIndex(self QColumnView, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QColumnView{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QColumnView) callVirtualBase_SelectAll() {

	QColumnView_virtualbase_SelectAll(unsafe.Pointer(this.h))

}
func (this *QColumnView) OnSelectAll(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SelectAll(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SelectAll
func miqt_exec_callback_QColumnView_SelectAll(self QColumnView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColumnView{h: self}).callVirtualBase_SelectAll)

}

func (this *QColumnView) callVirtualBase_IsIndexHidden(index *QModelIndex) bool {

	return (bool)(QColumnView_virtualbase_IsIndexHidden(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QColumnView) OnIsIndexHidden(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_IsIndexHidden
func miqt_exec_callback_QColumnView_IsIndexHidden(self QColumnView, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_IsIndexHidden, slotval1)

	return (bool)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_MoveCursor(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex {

	_goptr := newQModelIndex(QColumnView_virtualbase_MoveCursor(unsafe.Pointer(this.h), cursorAction, (int)(modifiers)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColumnView) OnMoveCursor(slot func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_MoveCursor
func miqt_exec_callback_QColumnView_MoveCursor(self QColumnView, cb intptr_t, cursorAction CursorAction, modifiers int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex, cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(modifiers)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_MoveCursor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QColumnView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_ResizeEvent
func miqt_exec_callback_QColumnView_ResizeEvent(self QColumnView, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_SetSelection(rect *QRect, command SelectionFlag) {

	QColumnView_virtualbase_SetSelection(unsafe.Pointer(this.h), rect.cPointer(), (int)(command))

}
func (this *QColumnView) OnSetSelection(slot func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SetSelection
func miqt_exec_callback_QColumnView_SetSelection(self QColumnView, cb intptr_t, rect *QRect, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect, command SelectionFlag), rect *QRect, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(command)

	gofunc((&QColumnView{h: self}).callVirtualBase_SetSelection, slotval1, slotval2)

}

func (this *QColumnView) callVirtualBase_VisualRegionForSelection(selection *QItemSelection) *QRegion {

	_goptr := newQRegion(QColumnView_virtualbase_VisualRegionForSelection(unsafe.Pointer(this.h), selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColumnView) OnVisualRegionForSelection(slot func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_VisualRegionForSelection
func miqt_exec_callback_QColumnView_VisualRegionForSelection(self QColumnView, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_VisualRegionForSelection, slotval1)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_HorizontalOffset() int {

	return (int)(QColumnView_virtualbase_HorizontalOffset(unsafe.Pointer(this.h)))

}
func (this *QColumnView) OnHorizontalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_HorizontalOffset
func miqt_exec_callback_QColumnView_HorizontalOffset(self QColumnView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_HorizontalOffset)

	return (int)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_VerticalOffset() int {

	return (int)(QColumnView_virtualbase_VerticalOffset(unsafe.Pointer(this.h)))

}
func (this *QColumnView) OnVerticalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_VerticalOffset
func miqt_exec_callback_QColumnView_VerticalOffset(self QColumnView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_VerticalOffset)

	return (int)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QColumnView_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QColumnView) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_RowsInserted
func miqt_exec_callback_QColumnView_RowsInserted(self QColumnView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QColumnView{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QColumnView) callVirtualBase_CurrentChanged(current *QModelIndex, previous *QModelIndex) {

	QColumnView_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), previous.cPointer())

}
func (this *QColumnView) OnCurrentChanged(slot func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_CurrentChanged
func miqt_exec_callback_QColumnView_CurrentChanged(self QColumnView, cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc((&QColumnView{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}

func (this *QColumnView) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QColumnView_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QColumnView) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_ScrollContentsBy
func miqt_exec_callback_QColumnView_ScrollContentsBy(self QColumnView, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QColumnView{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QColumnView) callVirtualBase_CreateColumn(rootIndex *QModelIndex) *QAbstractItemView {

	return newQAbstractItemView(QColumnView_virtualbase_CreateColumn(unsafe.Pointer(this.h), rootIndex.cPointer()))

}
func (this *QColumnView) OnCreateColumn(slot func(super func(rootIndex *QModelIndex) *QAbstractItemView, rootIndex *QModelIndex) *QAbstractItemView) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_CreateColumn(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_CreateColumn
func miqt_exec_callback_QColumnView_CreateColumn(self QColumnView, cb intptr_t, rootIndex *QModelIndex) *QAbstractItemView {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rootIndex *QModelIndex) *QAbstractItemView, rootIndex *QModelIndex) *QAbstractItemView)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(rootIndex)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_CreateColumn, slotval1)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))

	QColumnView_virtualbase_KeyboardSearch(unsafe.Pointer(this.h), search_ms)

}
func (this *QColumnView) OnKeyboardSearch(slot func(super func(search string), search string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_KeyboardSearch(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_KeyboardSearch
func miqt_exec_callback_QColumnView_KeyboardSearch(self QColumnView, cb intptr_t, search struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(search string), search string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var search_ms struct_miqt_string = search
	search_ret := GoStringN(search_ms.data, int(int64(search_ms.len)))
	free(unsafe.Pointer(search_ms.data))
	slotval1 := search_ret

	gofunc((&QColumnView{h: self}).callVirtualBase_KeyboardSearch, slotval1)

}

func (this *QColumnView) callVirtualBase_SizeHintForRow(row int) int {

	return (int)(QColumnView_virtualbase_SizeHintForRow(unsafe.Pointer(this.h), (int)(row)))

}
func (this *QColumnView) OnSizeHintForRow(slot func(super func(row int) int, row int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SizeHintForRow(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SizeHintForRow
func miqt_exec_callback_QColumnView_SizeHintForRow(self QColumnView, cb intptr_t, row int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int) int, row int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_SizeHintForRow, slotval1)

	return (int)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_SizeHintForColumn(column int) int {

	return (int)(QColumnView_virtualbase_SizeHintForColumn(unsafe.Pointer(this.h), (int)(column)))

}
func (this *QColumnView) OnSizeHintForColumn(slot func(super func(column int) int, column int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SizeHintForColumn(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SizeHintForColumn
func miqt_exec_callback_QColumnView_SizeHintForColumn(self QColumnView, cb intptr_t, column int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int) int, column int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_SizeHintForColumn, slotval1)

	return (int)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_ItemDelegateForIndex(index *QModelIndex) *QAbstractItemDelegate {

	return newQAbstractItemDelegate(QColumnView_virtualbase_ItemDelegateForIndex(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QColumnView) OnItemDelegateForIndex(slot func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_ItemDelegateForIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_ItemDelegateForIndex
func miqt_exec_callback_QColumnView_ItemDelegateForIndex(self QColumnView, cb intptr_t, index *QModelIndex) *QAbstractItemDelegate {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_ItemDelegateForIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_InputMethodQuery(query InputMethodQuery) *QVariant {

	_goptr := newQVariant(QColumnView_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColumnView) OnInputMethodQuery(slot func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_InputMethodQuery
func miqt_exec_callback_QColumnView_InputMethodQuery(self QColumnView, cb intptr_t, query int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(query)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QColumnView) callVirtualBase_Reset() {

	QColumnView_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QColumnView) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_Reset
func miqt_exec_callback_QColumnView_Reset(self QColumnView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColumnView{h: self}).callVirtualBase_Reset)

}

func (this *QColumnView) callVirtualBase_DoItemsLayout() {

	QColumnView_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QColumnView) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_DoItemsLayout
func miqt_exec_callback_QColumnView_DoItemsLayout(self QColumnView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColumnView{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QColumnView) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QColumnView_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QColumnView) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_DataChanged
func miqt_exec_callback_QColumnView_DataChanged(self QColumnView, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
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

	gofunc((&QColumnView{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QColumnView) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QColumnView_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QColumnView) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_RowsAboutToBeRemoved
func miqt_exec_callback_QColumnView_RowsAboutToBeRemoved(self QColumnView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QColumnView{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QColumnView) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QColumnView_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QColumnView) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SelectionChanged
func miqt_exec_callback_QColumnView_SelectionChanged(self QColumnView, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QColumnView{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QColumnView) callVirtualBase_UpdateEditorData() {

	QColumnView_virtualbase_UpdateEditorData(unsafe.Pointer(this.h))

}
func (this *QColumnView) OnUpdateEditorData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_UpdateEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_UpdateEditorData
func miqt_exec_callback_QColumnView_UpdateEditorData(self QColumnView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColumnView{h: self}).callVirtualBase_UpdateEditorData)

}

func (this *QColumnView) callVirtualBase_UpdateEditorGeometries() {

	QColumnView_virtualbase_UpdateEditorGeometries(unsafe.Pointer(this.h))

}
func (this *QColumnView) OnUpdateEditorGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_UpdateEditorGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_UpdateEditorGeometries
func miqt_exec_callback_QColumnView_UpdateEditorGeometries(self QColumnView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColumnView{h: self}).callVirtualBase_UpdateEditorGeometries)

}

func (this *QColumnView) callVirtualBase_UpdateGeometries() {

	QColumnView_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QColumnView) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_UpdateGeometries
func miqt_exec_callback_QColumnView_UpdateGeometries(self QColumnView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColumnView{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QColumnView) callVirtualBase_VerticalScrollbarAction(action int) {

	QColumnView_virtualbase_VerticalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QColumnView) OnVerticalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_VerticalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_VerticalScrollbarAction
func miqt_exec_callback_QColumnView_VerticalScrollbarAction(self QColumnView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QColumnView{h: self}).callVirtualBase_VerticalScrollbarAction, slotval1)

}

func (this *QColumnView) callVirtualBase_HorizontalScrollbarAction(action int) {

	QColumnView_virtualbase_HorizontalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QColumnView) OnHorizontalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_HorizontalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_HorizontalScrollbarAction
func miqt_exec_callback_QColumnView_HorizontalScrollbarAction(self QColumnView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QColumnView{h: self}).callVirtualBase_HorizontalScrollbarAction, slotval1)

}

func (this *QColumnView) callVirtualBase_VerticalScrollbarValueChanged(value int) {

	QColumnView_virtualbase_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QColumnView) OnVerticalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_VerticalScrollbarValueChanged
func miqt_exec_callback_QColumnView_VerticalScrollbarValueChanged(self QColumnView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QColumnView{h: self}).callVirtualBase_VerticalScrollbarValueChanged, slotval1)

}

func (this *QColumnView) callVirtualBase_HorizontalScrollbarValueChanged(value int) {

	QColumnView_virtualbase_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QColumnView) OnHorizontalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_HorizontalScrollbarValueChanged
func miqt_exec_callback_QColumnView_HorizontalScrollbarValueChanged(self QColumnView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QColumnView{h: self}).callVirtualBase_HorizontalScrollbarValueChanged, slotval1)

}

func (this *QColumnView) callVirtualBase_CloseEditor(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {

	QColumnView_virtualbase_CloseEditor(unsafe.Pointer(this.h), editor.cPointer(), (int)(hint))

}
func (this *QColumnView) OnCloseEditor(slot func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_CloseEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_CloseEditor
func miqt_exec_callback_QColumnView_CloseEditor(self QColumnView, cb intptr_t, editor *QWidget, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := (QAbstractItemDelegate__EndEditHint)(hint)

	gofunc((&QColumnView{h: self}).callVirtualBase_CloseEditor, slotval1, slotval2)

}

func (this *QColumnView) callVirtualBase_CommitData(editor *QWidget) {

	QColumnView_virtualbase_CommitData(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QColumnView) OnCommitData(slot func(super func(editor *QWidget), editor *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_CommitData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_CommitData
func miqt_exec_callback_QColumnView_CommitData(self QColumnView, cb intptr_t, editor *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget), editor *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	gofunc((&QColumnView{h: self}).callVirtualBase_CommitData, slotval1)

}

func (this *QColumnView) callVirtualBase_EditorDestroyed(editor *QObject) {

	QColumnView_virtualbase_EditorDestroyed(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QColumnView) OnEditorDestroyed(slot func(super func(editor *QObject), editor *QObject)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_EditorDestroyed(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_EditorDestroyed
func miqt_exec_callback_QColumnView_EditorDestroyed(self QColumnView, cb intptr_t, editor *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QObject), editor *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(editor)

	gofunc((&QColumnView{h: self}).callVirtualBase_EditorDestroyed, slotval1)

}

func (this *QColumnView) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QColumnView_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QColumnView) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SelectedIndexes
func miqt_exec_callback_QColumnView_SelectedIndexes(self QColumnView, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QColumnView) callVirtualBase_Edit2(index *QModelIndex, trigger EditTrigger, event *QEvent) bool {

	return (bool)(QColumnView_virtualbase_Edit2(unsafe.Pointer(this.h), index.cPointer(), trigger, event.cPointer()))

}
func (this *QColumnView) OnEdit2(slot func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_Edit2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_Edit2
func miqt_exec_callback_QColumnView_Edit2(self QColumnView, cb intptr_t, index *QModelIndex, trigger EditTrigger, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx
	slotval3 := newQEvent(event)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_Edit2, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_SelectionCommand(index *QModelIndex, event *QEvent) SelectionFlag {

	return (SelectionFlag)(QColumnView_virtualbase_SelectionCommand(unsafe.Pointer(this.h), index.cPointer(), event.cPointer()))

}
func (this *QColumnView) OnSelectionCommand(slot func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_SelectionCommand(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_SelectionCommand
func miqt_exec_callback_QColumnView_SelectionCommand(self QColumnView, cb intptr_t, index *QModelIndex, event *QEvent) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_SelectionCommand, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_StartDrag(supportedActions DropAction) {

	QColumnView_virtualbase_StartDrag(unsafe.Pointer(this.h), (int)(supportedActions))

}
func (this *QColumnView) OnStartDrag(slot func(super func(supportedActions DropAction), supportedActions DropAction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_StartDrag(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_StartDrag
func miqt_exec_callback_QColumnView_StartDrag(self QColumnView, cb intptr_t, supportedActions int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(supportedActions DropAction), supportedActions DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(supportedActions)

	gofunc((&QColumnView{h: self}).callVirtualBase_StartDrag, slotval1)

}

func (this *QColumnView) callVirtualBase_InitViewItemOption(option *QStyleOptionViewItem) {

	QColumnView_virtualbase_InitViewItemOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QColumnView) OnInitViewItemOption(slot func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_InitViewItemOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_InitViewItemOption
func miqt_exec_callback_QColumnView_InitViewItemOption(self QColumnView, cb intptr_t, option *QStyleOptionViewItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	gofunc((&QColumnView{h: self}).callVirtualBase_InitViewItemOption, slotval1)

}

func (this *QColumnView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QColumnView_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QColumnView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_FocusNextPrevChild
func miqt_exec_callback_QColumnView_FocusNextPrevChild(self QColumnView, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QColumnView_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QColumnView) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_Event
func miqt_exec_callback_QColumnView_Event(self QColumnView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_ViewportEvent(event *QEvent) bool {

	return (bool)(QColumnView_virtualbase_ViewportEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QColumnView) OnViewportEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_ViewportEvent
func miqt_exec_callback_QColumnView_ViewportEvent(self QColumnView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QColumnView_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_MousePressEvent
func miqt_exec_callback_QColumnView_MousePressEvent(self QColumnView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QColumnView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_MouseMoveEvent
func miqt_exec_callback_QColumnView_MouseMoveEvent(self QColumnView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QColumnView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_MouseReleaseEvent
func miqt_exec_callback_QColumnView_MouseReleaseEvent(self QColumnView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QColumnView_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_MouseDoubleClickEvent
func miqt_exec_callback_QColumnView_MouseDoubleClickEvent(self QColumnView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QColumnView_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_DragEnterEvent
func miqt_exec_callback_QColumnView_DragEnterEvent(self QColumnView, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QColumnView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_DragMoveEvent
func miqt_exec_callback_QColumnView_DragMoveEvent(self QColumnView, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QColumnView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_DragLeaveEvent
func miqt_exec_callback_QColumnView_DragLeaveEvent(self QColumnView, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_DropEvent(event *QDropEvent) {

	QColumnView_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_DropEvent
func miqt_exec_callback_QColumnView_DropEvent(self QColumnView, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QColumnView_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_FocusInEvent
func miqt_exec_callback_QColumnView_FocusInEvent(self QColumnView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QColumnView_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_FocusOutEvent
func miqt_exec_callback_QColumnView_FocusOutEvent(self QColumnView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QColumnView_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_KeyPressEvent
func miqt_exec_callback_QColumnView_KeyPressEvent(self QColumnView, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QColumnView_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_TimerEvent
func miqt_exec_callback_QColumnView_TimerEvent(self QColumnView, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_InputMethodEvent(event *QInputMethodEvent) {

	QColumnView_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColumnView) OnInputMethodEvent(slot func(super func(event *QInputMethodEvent), event *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_InputMethodEvent
func miqt_exec_callback_QColumnView_InputMethodEvent(self QColumnView, cb intptr_t, event *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QInputMethodEvent), event *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(event)

	gofunc((&QColumnView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QColumnView) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QColumnView_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QColumnView) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_EventFilter
func miqt_exec_callback_QColumnView_EventFilter(self QColumnView, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QColumnView) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QColumnView_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColumnView) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_ViewportSizeHint
func miqt_exec_callback_QColumnView_ViewportSizeHint(self QColumnView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}
