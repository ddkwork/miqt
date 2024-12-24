package qt

import (
	"unsafe"
)

type QListView__Movement int

const (
	QListView__Static QListView__Movement = 0
	QListView__Free   QListView__Movement = 1
	QListView__Snap   QListView__Movement = 2
)

type QListView__Flow int

const (
	QListView__LeftToRight QListView__Flow = 0
	QListView__TopToBottom QListView__Flow = 1
)

type QListView__ResizeMode int

const (
	QListView__Fixed  QListView__ResizeMode = 0
	QListView__Adjust QListView__ResizeMode = 1
)

type QListView__LayoutMode int

const (
	QListView__SinglePass QListView__LayoutMode = 0
	QListView__Batched    QListView__LayoutMode = 1
)

type QListView__ViewMode int

const (
	QListView__ListMode QListView__ViewMode = 0
	QListView__IconMode QListView__ViewMode = 1
)

type QListView struct {
	h          uintptr
	isSubclass bool
}

// NewQListView constructs a new QListView object.
func NewQListView(parent *QWidget) *QListView {
	ret := newQListView(QListView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQListView2 constructs a new QListView object.
func NewQListView2() *QListView {
	ret := newQListView(QListView_new2())
	ret.isSubclass = true
	return ret
}

func (this *QListView) MetaObject() *QMetaObject {
	return newQMetaObject(QListView_MetaObject(this.h))
}

func (this *QListView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QListView_Metacast(this.h, param1_Cstring))
}

func QListView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QListView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListView) SetMovement(movement Movement) {
	QListView_SetMovement(this.h, movement)
}

func (this *QListView) Movement() Movement {
	xxxxxxxxx
}

func (this *QListView) SetFlow(flow Flow) {
	QListView_SetFlow(this.h, flow)
}

func (this *QListView) Flow() Flow {
	xxxxxxxxx
}

func (this *QListView) SetWrapping(enable bool) {
	QListView_SetWrapping(this.h, (bool)(enable))
}

func (this *QListView) IsWrapping() bool {
	return (bool)(QListView_IsWrapping(this.h))
}

func (this *QListView) SetResizeMode(mode ResizeMode) {
	QListView_SetResizeMode(this.h, mode)
}

func (this *QListView) ResizeMode() ResizeMode {
	xxxxxxxxx
}

func (this *QListView) SetLayoutMode(mode LayoutMode) {
	QListView_SetLayoutMode(this.h, mode)
}

func (this *QListView) LayoutMode() LayoutMode {
	xxxxxxxxx
}

func (this *QListView) SetSpacing(space int) {
	QListView_SetSpacing(this.h, (int)(space))
}

func (this *QListView) Spacing() int {
	return (int)(QListView_Spacing(this.h))
}

func (this *QListView) SetBatchSize(batchSize int) {
	QListView_SetBatchSize(this.h, (int)(batchSize))
}

func (this *QListView) BatchSize() int {
	return (int)(QListView_BatchSize(this.h))
}

func (this *QListView) SetGridSize(size *QSize) {
	QListView_SetGridSize(this.h, size.cPointer())
}

func (this *QListView) GridSize() *QSize {
	_goptr := newQSize(QListView_GridSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListView) SetViewMode(mode ViewMode) {
	QListView_SetViewMode(this.h, mode)
}

func (this *QListView) ViewMode() ViewMode {
	xxxxxxxxx
}

func (this *QListView) ClearPropertyFlags() {
	QListView_ClearPropertyFlags(this.h)
}

func (this *QListView) IsRowHidden(row int) bool {
	return (bool)(QListView_IsRowHidden(this.h, (int)(row)))
}

func (this *QListView) SetRowHidden(row int, hide bool) {
	QListView_SetRowHidden(this.h, (int)(row), (bool)(hide))
}

func (this *QListView) SetModelColumn(column int) {
	QListView_SetModelColumn(this.h, (int)(column))
}

func (this *QListView) ModelColumn() int {
	return (int)(QListView_ModelColumn(this.h))
}

func (this *QListView) SetUniformItemSizes(enable bool) {
	QListView_SetUniformItemSizes(this.h, (bool)(enable))
}

func (this *QListView) UniformItemSizes() bool {
	return (bool)(QListView_UniformItemSizes(this.h))
}

func (this *QListView) SetWordWrap(on bool) {
	QListView_SetWordWrap(this.h, (bool)(on))
}

func (this *QListView) WordWrap() bool {
	return (bool)(QListView_WordWrap(this.h))
}

func (this *QListView) SetSelectionRectVisible(show bool) {
	QListView_SetSelectionRectVisible(this.h, (bool)(show))
}

func (this *QListView) IsSelectionRectVisible() bool {
	return (bool)(QListView_IsSelectionRectVisible(this.h))
}

func (this *QListView) SetItemAlignment(alignment AlignmentFlag) {
	QListView_SetItemAlignment(this.h, (int)(alignment))
}

func (this *QListView) ItemAlignment() AlignmentFlag {
	return (AlignmentFlag)(QListView_ItemAlignment(this.h))
}

func (this *QListView) VisualRect(index *QModelIndex) *QRect {
	_goptr := newQRect(QListView_VisualRect(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListView) ScrollTo(index *QModelIndex, hint ScrollHint) {
	QListView_ScrollTo(this.h, index.cPointer(), hint)
}

func (this *QListView) IndexAt(p *QPoint) *QModelIndex {
	_goptr := newQModelIndex(QListView_IndexAt(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QListView) DoItemsLayout() {
	QListView_DoItemsLayout(this.h)
}

func (this *QListView) Reset() {
	QListView_Reset(this.h)
}

func (this *QListView) SetRootIndex(index *QModelIndex) {
	QListView_SetRootIndex(this.h, index.cPointer())
}

func (this *QListView) IndexesMoved(indexes []QModelIndex) {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}
	QListView_IndexesMoved(this.h, indexes_ma)
}

func (this *QListView) OnIndexesMoved(slot func(indexes []QModelIndex)) {
	QListView_connect_IndexesMoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListView_IndexesMoved
func miqt_exec_callback_QListView_IndexesMoved(cb intptr_t, indexes struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(indexes []QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var indexes_ma struct_miqt_array = indexes
	indexes_ret := make([]QModelIndex, int(indexes_ma.len))
	indexes_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(indexes_ma.data)) // hey ya
	for i := 0; i < int(indexes_ma.len); i++ {
		indexes_lv_goptr := newQModelIndex(indexes_outCast[i])
		indexes_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		indexes_ret[i] = *indexes_lv_goptr
	}
	slotval1 := indexes_ret

	gofunc(slotval1)
}

func QListView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QListView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QListView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QListView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QListView) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QListView_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QListView) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListView_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListView_MetaObject
func miqt_exec_callback_QListView_MetaObject(self QListView, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QListView{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QListView) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QListView_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QListView) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QListView_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QListView_Metacast
func miqt_exec_callback_QListView_Metacast(self QListView, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QListView{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
