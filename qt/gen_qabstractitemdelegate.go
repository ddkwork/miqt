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

func (this *QAbstractItemDelegate) OnPaint(slot func(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_Paint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_Paint
func miqt_exec_callback_QAbstractItemDelegate_Paint(self QAbstractItemDelegate, cb intptr_t, painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	gofunc(slotval1, slotval2, slotval3)

}
func (this *QAbstractItemDelegate) OnSizeHint(slot func(option *QStyleOptionViewItem, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_SizeHint
func miqt_exec_callback_QAbstractItemDelegate_SizeHint(self QAbstractItemDelegate, cb intptr_t, option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(option *QStyleOptionViewItem, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	slotval2 := newQModelIndex(index)

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemDelegate) callVirtualBase_CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {

	return newQWidget(QAbstractItemDelegate_virtualbase_CreateEditor(unsafe.Pointer(this.h), parent.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QAbstractItemDelegate) OnCreateEditor(slot func(super func(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_CreateEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_CreateEditor
func miqt_exec_callback_QAbstractItemDelegate_CreateEditor(self QAbstractItemDelegate, cb intptr_t, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(parent)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_CreateEditor, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemDelegate) callVirtualBase_DestroyEditor(editor *QWidget, index *QModelIndex) {

	QAbstractItemDelegate_virtualbase_DestroyEditor(unsafe.Pointer(this.h), editor.cPointer(), index.cPointer())

}
func (this *QAbstractItemDelegate) OnDestroyEditor(slot func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_DestroyEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_DestroyEditor
func miqt_exec_callback_QAbstractItemDelegate_DestroyEditor(self QAbstractItemDelegate, cb intptr_t, editor *QWidget, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQModelIndex(index)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_DestroyEditor, slotval1, slotval2)

}

func (this *QAbstractItemDelegate) callVirtualBase_SetEditorData(editor *QWidget, index *QModelIndex) {

	QAbstractItemDelegate_virtualbase_SetEditorData(unsafe.Pointer(this.h), editor.cPointer(), index.cPointer())

}
func (this *QAbstractItemDelegate) OnSetEditorData(slot func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_SetEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_SetEditorData
func miqt_exec_callback_QAbstractItemDelegate_SetEditorData(self QAbstractItemDelegate, cb intptr_t, editor *QWidget, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQModelIndex(index)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_SetEditorData, slotval1, slotval2)

}

func (this *QAbstractItemDelegate) callVirtualBase_SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {

	QAbstractItemDelegate_virtualbase_SetModelData(unsafe.Pointer(this.h), editor.cPointer(), model.cPointer(), index.cPointer())

}
func (this *QAbstractItemDelegate) OnSetModelData(slot func(super func(editor *QWidget, model *QAbstractItemModel, index *QModelIndex), editor *QWidget, model *QAbstractItemModel, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_SetModelData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_SetModelData
func miqt_exec_callback_QAbstractItemDelegate_SetModelData(self QAbstractItemDelegate, cb intptr_t, editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, model *QAbstractItemModel, index *QModelIndex), editor *QWidget, model *QAbstractItemModel, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQAbstractItemModel(model)

	slotval3 := newQModelIndex(index)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_SetModelData, slotval1, slotval2, slotval3)

}

func (this *QAbstractItemDelegate) callVirtualBase_UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {

	QAbstractItemDelegate_virtualbase_UpdateEditorGeometry(unsafe.Pointer(this.h), editor.cPointer(), option.cPointer(), index.cPointer())

}
func (this *QAbstractItemDelegate) OnUpdateEditorGeometry(slot func(super func(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex), editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_UpdateEditorGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_UpdateEditorGeometry
func miqt_exec_callback_QAbstractItemDelegate_UpdateEditorGeometry(self QAbstractItemDelegate, cb intptr_t, editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex), editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_UpdateEditorGeometry, slotval1, slotval2, slotval3)

}

func (this *QAbstractItemDelegate) callVirtualBase_EditorEvent(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {

	return (bool)(QAbstractItemDelegate_virtualbase_EditorEvent(unsafe.Pointer(this.h), event.cPointer(), model.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QAbstractItemDelegate) OnEditorEvent(slot func(super func(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_EditorEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_EditorEvent
func miqt_exec_callback_QAbstractItemDelegate_EditorEvent(self QAbstractItemDelegate, cb intptr_t, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	slotval2 := newQAbstractItemModel(model)

	slotval3 := newQStyleOptionViewItem(option)

	slotval4 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_EditorEvent, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemDelegate) callVirtualBase_HelpEvent(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {

	return (bool)(QAbstractItemDelegate_virtualbase_HelpEvent(unsafe.Pointer(this.h), event.cPointer(), view.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QAbstractItemDelegate) OnHelpEvent(slot func(super func(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_HelpEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_HelpEvent
func miqt_exec_callback_QAbstractItemDelegate_HelpEvent(self QAbstractItemDelegate, cb intptr_t, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHelpEvent(event)

	slotval2 := newQAbstractItemView(view)

	slotval3 := newQStyleOptionViewItem(option)

	slotval4 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_HelpEvent, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemDelegate) callVirtualBase_PaintingRoles() []int {

	var _ma struct_miqt_array = QAbstractItemDelegate_virtualbase_PaintingRoles(unsafe.Pointer(this.h))
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret

}
func (this *QAbstractItemDelegate) OnPaintingRoles(slot func(super func() []int) []int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_PaintingRoles(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_PaintingRoles
func miqt_exec_callback_QAbstractItemDelegate_PaintingRoles(self QAbstractItemDelegate, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []int) []int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_PaintingRoles)
	virtualReturn_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = (int)(virtualReturn[i])
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractItemDelegate) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QAbstractItemDelegate_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QAbstractItemDelegate) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_Event
func miqt_exec_callback_QAbstractItemDelegate_Event(self QAbstractItemDelegate, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemDelegate) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QAbstractItemDelegate_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QAbstractItemDelegate) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_EventFilter
func miqt_exec_callback_QAbstractItemDelegate_EventFilter(self QAbstractItemDelegate, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemDelegate) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QAbstractItemDelegate_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemDelegate) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_TimerEvent
func miqt_exec_callback_QAbstractItemDelegate_TimerEvent(self QAbstractItemDelegate, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAbstractItemDelegate) callVirtualBase_ChildEvent(event *QChildEvent) {

	QAbstractItemDelegate_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemDelegate) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_ChildEvent
func miqt_exec_callback_QAbstractItemDelegate_ChildEvent(self QAbstractItemDelegate, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QAbstractItemDelegate) callVirtualBase_CustomEvent(event *QEvent) {

	QAbstractItemDelegate_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemDelegate) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_CustomEvent
func miqt_exec_callback_QAbstractItemDelegate_CustomEvent(self QAbstractItemDelegate, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QAbstractItemDelegate) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QAbstractItemDelegate_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QAbstractItemDelegate) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_ConnectNotify
func miqt_exec_callback_QAbstractItemDelegate_ConnectNotify(self QAbstractItemDelegate, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QAbstractItemDelegate) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QAbstractItemDelegate_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QAbstractItemDelegate) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemDelegate_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemDelegate_DisconnectNotify
func miqt_exec_callback_QAbstractItemDelegate_DisconnectNotify(self QAbstractItemDelegate, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QAbstractItemDelegate{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
