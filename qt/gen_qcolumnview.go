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
	g := newQColumnView(QColumnView_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQColumnView2 constructs a new QColumnView object.
func NewQColumnView2() *QColumnView {
	g := newQColumnView(QColumnView_new2())
	g.isSubclass = true
	return g
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

func (this *QColumnView) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QColumnView_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QColumnView) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_MetaObject
func miqt_exec_callback_QColumnView_MetaObject(self QColumnView, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QColumnView) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QColumnView_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QColumnView) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColumnView_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColumnView_Metacast
func miqt_exec_callback_QColumnView_Metacast(self QColumnView, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QColumnView{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
