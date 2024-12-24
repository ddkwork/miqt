package qt

import (
	"unsafe"
)

type QAbstractItemDelegate__EndEditHint int

const (
	QAbstractItemDelegate__NoHint           QAbstractItemDelegate__EndEditHint = 0
	QAbstractItemDelegate__EditNextItem     QAbstractItemDelegate__EndEditHint = 1
	QAbstractItemDelegate__EditPreviousItem QAbstractItemDelegate__EndEditHint = 2
	QAbstractItemDelegate__SubmitModelCache QAbstractItemDelegate__EndEditHint = 3
	QAbstractItemDelegate__RevertModelCache QAbstractItemDelegate__EndEditHint = 4
)

type QAbstractItemDelegate struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractItemDelegate constructs a new QAbstractItemDelegate object.
func NewQAbstractItemDelegate() *QAbstractItemDelegate {
	ret := newQAbstractItemDelegate(QAbstractItemDelegate_new())
	ret.isSubclass = true
	return ret
}

// NewQAbstractItemDelegate2 constructs a new QAbstractItemDelegate object.
func NewQAbstractItemDelegate2(parent *QObject) *QAbstractItemDelegate {
	ret := newQAbstractItemDelegate(QAbstractItemDelegate_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractItemDelegate) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractItemDelegate_MetaObject(this.h))
}

func (this *QAbstractItemDelegate) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractItemDelegate_Metacast(this.h, param1_Cstring))
}

func QAbstractItemDelegate_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractItemDelegate_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractItemDelegate) Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	QAbstractItemDelegate_Paint(this.h, painter.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QAbstractItemDelegate) SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	_goptr := newQSize(QAbstractItemDelegate_SizeHint(this.h, option.cPointer(), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemDelegate) CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	return newQWidget(QAbstractItemDelegate_CreateEditor(this.h, parent.cPointer(), option.cPointer(), index.cPointer()))
}

func (this *QAbstractItemDelegate) DestroyEditor(editor *QWidget, index *QModelIndex) {
	QAbstractItemDelegate_DestroyEditor(this.h, editor.cPointer(), index.cPointer())
}

func (this *QAbstractItemDelegate) SetEditorData(editor *QWidget, index *QModelIndex) {
	QAbstractItemDelegate_SetEditorData(this.h, editor.cPointer(), index.cPointer())
}

func (this *QAbstractItemDelegate) SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	QAbstractItemDelegate_SetModelData(this.h, editor.cPointer(), model.cPointer(), index.cPointer())
}

func (this *QAbstractItemDelegate) UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	QAbstractItemDelegate_UpdateEditorGeometry(this.h, editor.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QAbstractItemDelegate) EditorEvent(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	return (bool)(QAbstractItemDelegate_EditorEvent(this.h, event.cPointer(), model.cPointer(), option.cPointer(), index.cPointer()))
}

func (this *QAbstractItemDelegate) HelpEvent(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	return (bool)(QAbstractItemDelegate_HelpEvent(this.h, event.cPointer(), view.cPointer(), option.cPointer(), index.cPointer()))
}

func (this *QAbstractItemDelegate) PaintingRoles() []int {
	var _ma struct_miqt_array = QAbstractItemDelegate_PaintingRoles(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QAbstractItemDelegate) CommitData(editor *QWidget) {
	QAbstractItemDelegate_CommitData(this.h, editor.cPointer())
}

func (this *QAbstractItemDelegate) OnCommitData(slot func(editor *QWidget)) {
	QAbstractItemDelegate_connect_CommitData(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_CommitData
func miqt_exec_callback_QAbstractItemDelegate_CommitData(cb intptr_t, editor *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(editor *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	gofunc(slotval1)
}

func (this *QAbstractItemDelegate) CloseEditor(editor *QWidget) {
	QAbstractItemDelegate_CloseEditor(this.h, editor.cPointer())
}

func (this *QAbstractItemDelegate) OnCloseEditor(slot func(editor *QWidget)) {
	QAbstractItemDelegate_connect_CloseEditor(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_CloseEditor
func miqt_exec_callback_QAbstractItemDelegate_CloseEditor(cb intptr_t, editor *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(editor *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	gofunc(slotval1)
}

func (this *QAbstractItemDelegate) SizeHintChanged(param1 *QModelIndex) {
	QAbstractItemDelegate_SizeHintChanged(this.h, param1.cPointer())
}

func (this *QAbstractItemDelegate) OnSizeHintChanged(slot func(param1 *QModelIndex)) {
	QAbstractItemDelegate_connect_SizeHintChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_SizeHintChanged
func miqt_exec_callback_QAbstractItemDelegate_SizeHintChanged(cb intptr_t, param1 *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(param1)

	gofunc(slotval1)
}

func QAbstractItemDelegate_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractItemDelegate_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractItemDelegate_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractItemDelegate_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractItemDelegate) CloseEditor2(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	QAbstractItemDelegate_CloseEditor2(this.h, editor.cPointer(), (int)(hint))
}

func (this *QAbstractItemDelegate) OnCloseEditor2(slot func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	QAbstractItemDelegate_connect_CloseEditor2(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_CloseEditor2
func miqt_exec_callback_QAbstractItemDelegate_CloseEditor2(cb intptr_t, editor *QWidget, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := (QAbstractItemDelegate__EndEditHint)(hint)

	gofunc(slotval1, slotval2)
}

func (this *QAbstractItemDelegate) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractItemDelegate_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractItemDelegate) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_MetaObject
func miqt_exec_callback_QAbstractItemDelegate_MetaObject(self QAbstractItemDelegate, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractItemDelegate) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractItemDelegate_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractItemDelegate) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_Metacast
func miqt_exec_callback_QAbstractItemDelegate_Metacast(self QAbstractItemDelegate, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
