package qt

import (
	"unsafe"
)

type QStyledItemDelegate struct {
	h          uintptr
	isSubclass bool
}

// NewQStyledItemDelegate constructs a new QStyledItemDelegate object.
func NewQStyledItemDelegate() *QStyledItemDelegate {

	ret := newQStyledItemDelegate(QStyledItemDelegate_new())
	ret.isSubclass = true
	return ret
}

// NewQStyledItemDelegate2 constructs a new QStyledItemDelegate object.
func NewQStyledItemDelegate2(parent *QObject) *QStyledItemDelegate {

	ret := newQStyledItemDelegate(QStyledItemDelegate_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStyledItemDelegate) MetaObject() *QMetaObject {
	return newQMetaObject(QStyledItemDelegate_MetaObject(this.h))
}

func (this *QStyledItemDelegate) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStyledItemDelegate_Metacast(this.h, param1_Cstring))
}

func QStyledItemDelegate_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStyledItemDelegate_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStyledItemDelegate) Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	QStyledItemDelegate_Paint(this.h, painter.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	_goptr := newQSize(QStyledItemDelegate_SizeHint(this.h, option.cPointer(), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStyledItemDelegate) CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	return newQWidget(QStyledItemDelegate_CreateEditor(this.h, parent.cPointer(), option.cPointer(), index.cPointer()))
}

func (this *QStyledItemDelegate) SetEditorData(editor *QWidget, index *QModelIndex) {
	QStyledItemDelegate_SetEditorData(this.h, editor.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	QStyledItemDelegate_SetModelData(this.h, editor.cPointer(), model.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	QStyledItemDelegate_UpdateEditorGeometry(this.h, editor.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	return newQItemEditorFactory(QStyledItemDelegate_ItemEditorFactory(this.h))
}

func (this *QStyledItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory) {
	QStyledItemDelegate_SetItemEditorFactory(this.h, factory.cPointer())
}

func (this *QStyledItemDelegate) DisplayText(value *QVariant, locale *QLocale) string {
	var _ms struct_miqt_string = QStyledItemDelegate_DisplayText(this.h, value.cPointer(), locale.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStyledItemDelegate_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStyledItemDelegate_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStyledItemDelegate_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStyledItemDelegate_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStyledItemDelegate) callVirtualBase_Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {

	QStyledItemDelegate_virtualbase_Paint(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), index.cPointer())

}
func (this *QStyledItemDelegate) OnPaint(slot func(super func(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_Paint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_Paint
func miqt_exec_callback_QStyledItemDelegate_Paint(self QStyledItemDelegate, cb intptr_t, painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_Paint, slotval1, slotval2, slotval3)

}

func (this *QStyledItemDelegate) callVirtualBase_SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {

	_goptr := newQSize(QStyledItemDelegate_virtualbase_SizeHint(unsafe.Pointer(this.h), option.cPointer(), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QStyledItemDelegate) OnSizeHint(slot func(super func(option *QStyleOptionViewItem, index *QModelIndex) *QSize, option *QStyleOptionViewItem, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_SizeHint
func miqt_exec_callback_QStyledItemDelegate_SizeHint(self QStyledItemDelegate, cb intptr_t, option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem, index *QModelIndex) *QSize, option *QStyleOptionViewItem, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	slotval2 := newQModelIndex(index)

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_SizeHint, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QStyledItemDelegate) callVirtualBase_CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {

	return newQWidget(QStyledItemDelegate_virtualbase_CreateEditor(unsafe.Pointer(this.h), parent.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QStyledItemDelegate) OnCreateEditor(slot func(super func(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_CreateEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_CreateEditor
func miqt_exec_callback_QStyledItemDelegate_CreateEditor(self QStyledItemDelegate, cb intptr_t, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(parent)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_CreateEditor, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QStyledItemDelegate) callVirtualBase_SetEditorData(editor *QWidget, index *QModelIndex) {

	QStyledItemDelegate_virtualbase_SetEditorData(unsafe.Pointer(this.h), editor.cPointer(), index.cPointer())

}
func (this *QStyledItemDelegate) OnSetEditorData(slot func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_SetEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_SetEditorData
func miqt_exec_callback_QStyledItemDelegate_SetEditorData(self QStyledItemDelegate, cb intptr_t, editor *QWidget, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQModelIndex(index)

	gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_SetEditorData, slotval1, slotval2)

}

func (this *QStyledItemDelegate) callVirtualBase_SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {

	QStyledItemDelegate_virtualbase_SetModelData(unsafe.Pointer(this.h), editor.cPointer(), model.cPointer(), index.cPointer())

}
func (this *QStyledItemDelegate) OnSetModelData(slot func(super func(editor *QWidget, model *QAbstractItemModel, index *QModelIndex), editor *QWidget, model *QAbstractItemModel, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_SetModelData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_SetModelData
func miqt_exec_callback_QStyledItemDelegate_SetModelData(self QStyledItemDelegate, cb intptr_t, editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, model *QAbstractItemModel, index *QModelIndex), editor *QWidget, model *QAbstractItemModel, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQAbstractItemModel(model)

	slotval3 := newQModelIndex(index)

	gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_SetModelData, slotval1, slotval2, slotval3)

}

func (this *QStyledItemDelegate) callVirtualBase_UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {

	QStyledItemDelegate_virtualbase_UpdateEditorGeometry(unsafe.Pointer(this.h), editor.cPointer(), option.cPointer(), index.cPointer())

}
func (this *QStyledItemDelegate) OnUpdateEditorGeometry(slot func(super func(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex), editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_UpdateEditorGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_UpdateEditorGeometry
func miqt_exec_callback_QStyledItemDelegate_UpdateEditorGeometry(self QStyledItemDelegate, cb intptr_t, editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex), editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_UpdateEditorGeometry, slotval1, slotval2, slotval3)

}

func (this *QStyledItemDelegate) callVirtualBase_DisplayText(value *QVariant, locale *QLocale) string {

	var _ms struct_miqt_string = QStyledItemDelegate_virtualbase_DisplayText(unsafe.Pointer(this.h), value.cPointer(), locale.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QStyledItemDelegate) OnDisplayText(slot func(super func(value *QVariant, locale *QLocale) string, value *QVariant, locale *QLocale) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_DisplayText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_DisplayText
func miqt_exec_callback_QStyledItemDelegate_DisplayText(self QStyledItemDelegate, cb intptr_t, value *QVariant, locale *QLocale) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value *QVariant, locale *QLocale) string, value *QVariant, locale *QLocale) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(value)

	slotval2 := newQLocale(locale)

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_DisplayText, slotval1, slotval2)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QStyledItemDelegate) callVirtualBase_InitStyleOption(option *QStyleOptionViewItem, index *QModelIndex) {

	QStyledItemDelegate_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer(), index.cPointer())

}
func (this *QStyledItemDelegate) OnInitStyleOption(slot func(super func(option *QStyleOptionViewItem, index *QModelIndex), option *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_InitStyleOption
func miqt_exec_callback_QStyledItemDelegate_InitStyleOption(self QStyledItemDelegate, cb intptr_t, option *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem, index *QModelIndex), option *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	slotval2 := newQModelIndex(index)

	gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_InitStyleOption, slotval1, slotval2)

}

func (this *QStyledItemDelegate) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QStyledItemDelegate_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QStyledItemDelegate) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_EventFilter
func miqt_exec_callback_QStyledItemDelegate_EventFilter(self QStyledItemDelegate, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QStyledItemDelegate) callVirtualBase_EditorEvent(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {

	return (bool)(QStyledItemDelegate_virtualbase_EditorEvent(unsafe.Pointer(this.h), event.cPointer(), model.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QStyledItemDelegate) OnEditorEvent(slot func(super func(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_EditorEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_EditorEvent
func miqt_exec_callback_QStyledItemDelegate_EditorEvent(self QStyledItemDelegate, cb intptr_t, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	slotval2 := newQAbstractItemModel(model)

	slotval3 := newQStyleOptionViewItem(option)

	slotval4 := newQModelIndex(index)

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_EditorEvent, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QStyledItemDelegate) callVirtualBase_DestroyEditor(editor *QWidget, index *QModelIndex) {

	QStyledItemDelegate_virtualbase_DestroyEditor(unsafe.Pointer(this.h), editor.cPointer(), index.cPointer())

}
func (this *QStyledItemDelegate) OnDestroyEditor(slot func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_DestroyEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_DestroyEditor
func miqt_exec_callback_QStyledItemDelegate_DestroyEditor(self QStyledItemDelegate, cb intptr_t, editor *QWidget, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQModelIndex(index)

	gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_DestroyEditor, slotval1, slotval2)

}

func (this *QStyledItemDelegate) callVirtualBase_HelpEvent(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {

	return (bool)(QStyledItemDelegate_virtualbase_HelpEvent(unsafe.Pointer(this.h), event.cPointer(), view.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QStyledItemDelegate) OnHelpEvent(slot func(super func(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_HelpEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_HelpEvent
func miqt_exec_callback_QStyledItemDelegate_HelpEvent(self QStyledItemDelegate, cb intptr_t, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHelpEvent(event)

	slotval2 := newQAbstractItemView(view)

	slotval3 := newQStyleOptionViewItem(option)

	slotval4 := newQModelIndex(index)

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_HelpEvent, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QStyledItemDelegate) callVirtualBase_PaintingRoles() []int {

	var _ma struct_miqt_array = QStyledItemDelegate_virtualbase_PaintingRoles(unsafe.Pointer(this.h))
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret

}
func (this *QStyledItemDelegate) OnPaintingRoles(slot func(super func() []int) []int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_PaintingRoles(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_PaintingRoles
func miqt_exec_callback_QStyledItemDelegate_PaintingRoles(self QStyledItemDelegate, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []int) []int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_PaintingRoles)
	virtualReturn_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = (int)(virtualReturn[i])
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}
