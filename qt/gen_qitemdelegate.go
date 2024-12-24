package qt

import (
	"unsafe"
)

type QItemDelegate struct {
	h          uintptr
	isSubclass bool
}

// NewQItemDelegate constructs a new QItemDelegate object.
func NewQItemDelegate() *QItemDelegate {

	ret := newQItemDelegate(QItemDelegate_new())
	ret.isSubclass = true
	return ret
}

// NewQItemDelegate2 constructs a new QItemDelegate object.
func NewQItemDelegate2(parent *QObject) *QItemDelegate {

	ret := newQItemDelegate(QItemDelegate_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QItemDelegate) MetaObject() *QMetaObject {
	return newQMetaObject(QItemDelegate_MetaObject(this.h))
}

func (this *QItemDelegate) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QItemDelegate_Metacast(this.h, param1_Cstring))
}

func QItemDelegate_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QItemDelegate_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QItemDelegate) HasClipping() bool {
	return (bool)(QItemDelegate_HasClipping(this.h))
}

func (this *QItemDelegate) SetClipping(clip bool) {
	QItemDelegate_SetClipping(this.h, (bool)(clip))
}

func (this *QItemDelegate) Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	QItemDelegate_Paint(this.h, painter.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QItemDelegate) SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	_goptr := newQSize(QItemDelegate_SizeHint(this.h, option.cPointer(), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QItemDelegate) CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	return newQWidget(QItemDelegate_CreateEditor(this.h, parent.cPointer(), option.cPointer(), index.cPointer()))
}

func (this *QItemDelegate) SetEditorData(editor *QWidget, index *QModelIndex) {
	QItemDelegate_SetEditorData(this.h, editor.cPointer(), index.cPointer())
}

func (this *QItemDelegate) SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	QItemDelegate_SetModelData(this.h, editor.cPointer(), model.cPointer(), index.cPointer())
}

func (this *QItemDelegate) UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	QItemDelegate_UpdateEditorGeometry(this.h, editor.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	return newQItemEditorFactory(QItemDelegate_ItemEditorFactory(this.h))
}

func (this *QItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory) {
	QItemDelegate_SetItemEditorFactory(this.h, factory.cPointer())
}

func QItemDelegate_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QItemDelegate_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QItemDelegate_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QItemDelegate_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QItemDelegate) callVirtualBase_Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {

	QItemDelegate_virtualbase_Paint(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), index.cPointer())

}
func (this *QItemDelegate) OnPaint(slot func(super func(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_Paint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_Paint
func miqt_exec_callback_QItemDelegate_Paint(self QItemDelegate, cb intptr_t, painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex), painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_Paint, slotval1, slotval2, slotval3)

}

func (this *QItemDelegate) callVirtualBase_SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {

	_goptr := newQSize(QItemDelegate_virtualbase_SizeHint(unsafe.Pointer(this.h), option.cPointer(), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QItemDelegate) OnSizeHint(slot func(super func(option *QStyleOptionViewItem, index *QModelIndex) *QSize, option *QStyleOptionViewItem, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_SizeHint
func miqt_exec_callback_QItemDelegate_SizeHint(self QItemDelegate, cb intptr_t, option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem, index *QModelIndex) *QSize, option *QStyleOptionViewItem, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	slotval2 := newQModelIndex(index)

	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_SizeHint, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QItemDelegate) callVirtualBase_CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {

	return newQWidget(QItemDelegate_virtualbase_CreateEditor(unsafe.Pointer(this.h), parent.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QItemDelegate) OnCreateEditor(slot func(super func(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_CreateEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_CreateEditor
func miqt_exec_callback_QItemDelegate_CreateEditor(self QItemDelegate, cb intptr_t, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget, parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(parent)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_CreateEditor, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QItemDelegate) callVirtualBase_SetEditorData(editor *QWidget, index *QModelIndex) {

	QItemDelegate_virtualbase_SetEditorData(unsafe.Pointer(this.h), editor.cPointer(), index.cPointer())

}
func (this *QItemDelegate) OnSetEditorData(slot func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_SetEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_SetEditorData
func miqt_exec_callback_QItemDelegate_SetEditorData(self QItemDelegate, cb intptr_t, editor *QWidget, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQModelIndex(index)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_SetEditorData, slotval1, slotval2)

}

func (this *QItemDelegate) callVirtualBase_SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {

	QItemDelegate_virtualbase_SetModelData(unsafe.Pointer(this.h), editor.cPointer(), model.cPointer(), index.cPointer())

}
func (this *QItemDelegate) OnSetModelData(slot func(super func(editor *QWidget, model *QAbstractItemModel, index *QModelIndex), editor *QWidget, model *QAbstractItemModel, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_SetModelData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_SetModelData
func miqt_exec_callback_QItemDelegate_SetModelData(self QItemDelegate, cb intptr_t, editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, model *QAbstractItemModel, index *QModelIndex), editor *QWidget, model *QAbstractItemModel, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQAbstractItemModel(model)

	slotval3 := newQModelIndex(index)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_SetModelData, slotval1, slotval2, slotval3)

}

func (this *QItemDelegate) callVirtualBase_UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {

	QItemDelegate_virtualbase_UpdateEditorGeometry(unsafe.Pointer(this.h), editor.cPointer(), option.cPointer(), index.cPointer())

}
func (this *QItemDelegate) OnUpdateEditorGeometry(slot func(super func(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex), editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_UpdateEditorGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_UpdateEditorGeometry
func miqt_exec_callback_QItemDelegate_UpdateEditorGeometry(self QItemDelegate, cb intptr_t, editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex), editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQModelIndex(index)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_UpdateEditorGeometry, slotval1, slotval2, slotval3)

}

func (this *QItemDelegate) callVirtualBase_DrawDisplay(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	QItemDelegate_virtualbase_DrawDisplay(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), rect.cPointer(), text_ms)

}
func (this *QItemDelegate) OnDrawDisplay(slot func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, text string), painter *QPainter, option *QStyleOptionViewItem, rect *QRect, text string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_DrawDisplay(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_DrawDisplay
func miqt_exec_callback_QItemDelegate_DrawDisplay(self QItemDelegate, cb intptr_t, painter *QPainter, option *QStyleOptionViewItem, rect *QRect, text struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, text string), painter *QPainter, option *QStyleOptionViewItem, rect *QRect, text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQRect(rect)

	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval4 := text_ret

	gofunc((&QItemDelegate{h: self}).callVirtualBase_DrawDisplay, slotval1, slotval2, slotval3, slotval4)

}

func (this *QItemDelegate) callVirtualBase_DrawDecoration(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, pixmap *QPixmap) {

	QItemDelegate_virtualbase_DrawDecoration(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), rect.cPointer(), pixmap.cPointer())

}
func (this *QItemDelegate) OnDrawDecoration(slot func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, pixmap *QPixmap), painter *QPainter, option *QStyleOptionViewItem, rect *QRect, pixmap *QPixmap)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_DrawDecoration(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_DrawDecoration
func miqt_exec_callback_QItemDelegate_DrawDecoration(self QItemDelegate, cb intptr_t, painter *QPainter, option *QStyleOptionViewItem, rect *QRect, pixmap *QPixmap) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, pixmap *QPixmap), painter *QPainter, option *QStyleOptionViewItem, rect *QRect, pixmap *QPixmap))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQRect(rect)

	slotval4 := newQPixmap(pixmap)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_DrawDecoration, slotval1, slotval2, slotval3, slotval4)

}

func (this *QItemDelegate) callVirtualBase_DrawFocus(painter *QPainter, option *QStyleOptionViewItem, rect *QRect) {

	QItemDelegate_virtualbase_DrawFocus(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), rect.cPointer())

}
func (this *QItemDelegate) OnDrawFocus(slot func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect), painter *QPainter, option *QStyleOptionViewItem, rect *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_DrawFocus(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_DrawFocus
func miqt_exec_callback_QItemDelegate_DrawFocus(self QItemDelegate, cb intptr_t, painter *QPainter, option *QStyleOptionViewItem, rect *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect), painter *QPainter, option *QStyleOptionViewItem, rect *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQRect(rect)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_DrawFocus, slotval1, slotval2, slotval3)

}

func (this *QItemDelegate) callVirtualBase_DrawCheck(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, state CheckState) {

	QItemDelegate_virtualbase_DrawCheck(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), rect.cPointer(), (int)(state))

}
func (this *QItemDelegate) OnDrawCheck(slot func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, state CheckState), painter *QPainter, option *QStyleOptionViewItem, rect *QRect, state CheckState)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_DrawCheck(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_DrawCheck
func miqt_exec_callback_QItemDelegate_DrawCheck(self QItemDelegate, cb intptr_t, painter *QPainter, option *QStyleOptionViewItem, rect *QRect, state int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, state CheckState), painter *QPainter, option *QStyleOptionViewItem, rect *QRect, state CheckState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionViewItem(option)

	slotval3 := newQRect(rect)

	slotval4 := (CheckState)(state)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_DrawCheck, slotval1, slotval2, slotval3, slotval4)

}

func (this *QItemDelegate) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QItemDelegate_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QItemDelegate) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_EventFilter
func miqt_exec_callback_QItemDelegate_EventFilter(self QItemDelegate, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QItemDelegate) callVirtualBase_EditorEvent(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {

	return (bool)(QItemDelegate_virtualbase_EditorEvent(unsafe.Pointer(this.h), event.cPointer(), model.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QItemDelegate) OnEditorEvent(slot func(super func(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_EditorEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_EditorEvent
func miqt_exec_callback_QItemDelegate_EditorEvent(self QItemDelegate, cb intptr_t, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	slotval2 := newQAbstractItemModel(model)

	slotval3 := newQStyleOptionViewItem(option)

	slotval4 := newQModelIndex(index)

	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_EditorEvent, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QItemDelegate) callVirtualBase_DestroyEditor(editor *QWidget, index *QModelIndex) {

	QItemDelegate_virtualbase_DestroyEditor(unsafe.Pointer(this.h), editor.cPointer(), index.cPointer())

}
func (this *QItemDelegate) OnDestroyEditor(slot func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_DestroyEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_DestroyEditor
func miqt_exec_callback_QItemDelegate_DestroyEditor(self QItemDelegate, cb intptr_t, editor *QWidget, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, index *QModelIndex), editor *QWidget, index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := newQModelIndex(index)

	gofunc((&QItemDelegate{h: self}).callVirtualBase_DestroyEditor, slotval1, slotval2)

}

func (this *QItemDelegate) callVirtualBase_HelpEvent(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {

	return (bool)(QItemDelegate_virtualbase_HelpEvent(unsafe.Pointer(this.h), event.cPointer(), view.cPointer(), option.cPointer(), index.cPointer()))

}
func (this *QItemDelegate) OnHelpEvent(slot func(super func(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_HelpEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_HelpEvent
func miqt_exec_callback_QItemDelegate_HelpEvent(self QItemDelegate, cb intptr_t, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool, event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHelpEvent(event)

	slotval2 := newQAbstractItemView(view)

	slotval3 := newQStyleOptionViewItem(option)

	slotval4 := newQModelIndex(index)

	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_HelpEvent, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QItemDelegate) callVirtualBase_PaintingRoles() []int {

	var _ma struct_miqt_array = QItemDelegate_virtualbase_PaintingRoles(unsafe.Pointer(this.h))
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret

}
func (this *QItemDelegate) OnPaintingRoles(slot func(super func() []int) []int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_PaintingRoles(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_PaintingRoles
func miqt_exec_callback_QItemDelegate_PaintingRoles(self QItemDelegate, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []int) []int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_PaintingRoles)
	virtualReturn_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = (int)(virtualReturn[i])
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}
