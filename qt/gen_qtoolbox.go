package qt

import (
	"unsafe"
)

type QToolBox struct {
	h          uintptr
	isSubclass bool
}

// NewQToolBox constructs a new QToolBox object.
func NewQToolBox(parent *QWidget) *QToolBox {

	ret := newQToolBox(QToolBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQToolBox2 constructs a new QToolBox object.
func NewQToolBox2() *QToolBox {

	ret := newQToolBox(QToolBox_new2())
	ret.isSubclass = true
	return ret
}

// NewQToolBox3 constructs a new QToolBox object.
func NewQToolBox3(parent *QWidget, f WindowType) *QToolBox {

	ret := newQToolBox(QToolBox_new3(parent.cPointer(), (int)(f)))
	ret.isSubclass = true
	return ret
}

func (this *QToolBox) MetaObject() *QMetaObject {
	return newQMetaObject(QToolBox_MetaObject(this.h))
}

func (this *QToolBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QToolBox_Metacast(this.h, param1_Cstring))
}

func QToolBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QToolBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolBox) AddItem(widget *QWidget, text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QToolBox_AddItem(this.h, widget.cPointer(), text_ms))
}

func (this *QToolBox) AddItem2(widget *QWidget, icon *QIcon, text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QToolBox_AddItem2(this.h, widget.cPointer(), icon.cPointer(), text_ms))
}

func (this *QToolBox) InsertItem(index int, widget *QWidget, text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QToolBox_InsertItem(this.h, (int)(index), widget.cPointer(), text_ms))
}

func (this *QToolBox) InsertItem2(index int, widget *QWidget, icon *QIcon, text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QToolBox_InsertItem2(this.h, (int)(index), widget.cPointer(), icon.cPointer(), text_ms))
}

func (this *QToolBox) RemoveItem(index int) {
	QToolBox_RemoveItem(this.h, (int)(index))
}

func (this *QToolBox) SetItemEnabled(index int, enabled bool) {
	QToolBox_SetItemEnabled(this.h, (int)(index), (bool)(enabled))
}

func (this *QToolBox) IsItemEnabled(index int) bool {
	return (bool)(QToolBox_IsItemEnabled(this.h, (int)(index)))
}

func (this *QToolBox) SetItemText(index int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QToolBox_SetItemText(this.h, (int)(index), text_ms)
}

func (this *QToolBox) ItemText(index int) string {
	var _ms struct_miqt_string = QToolBox_ItemText(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolBox) SetItemIcon(index int, icon *QIcon) {
	QToolBox_SetItemIcon(this.h, (int)(index), icon.cPointer())
}

func (this *QToolBox) ItemIcon(index int) *QIcon {
	_goptr := newQIcon(QToolBox_ItemIcon(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QToolBox) SetItemToolTip(index int, toolTip string) {
	toolTip_ms := struct_miqt_string{}
	toolTip_ms.data = CString(toolTip)
	toolTip_ms.len = size_t(len(toolTip))
	defer free(unsafe.Pointer(toolTip_ms.data))
	QToolBox_SetItemToolTip(this.h, (int)(index), toolTip_ms)
}

func (this *QToolBox) ItemToolTip(index int) string {
	var _ms struct_miqt_string = QToolBox_ItemToolTip(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolBox) CurrentIndex() int {
	return (int)(QToolBox_CurrentIndex(this.h))
}

func (this *QToolBox) CurrentWidget() *QWidget {
	return newQWidget(QToolBox_CurrentWidget(this.h))
}

func (this *QToolBox) Widget(index int) *QWidget {
	return newQWidget(QToolBox_Widget(this.h, (int)(index)))
}

func (this *QToolBox) IndexOf(widget *QWidget) int {
	return (int)(QToolBox_IndexOf(this.h, widget.cPointer()))
}

func (this *QToolBox) Count() int {
	return (int)(QToolBox_Count(this.h))
}

func (this *QToolBox) SetCurrentIndex(index int) {
	QToolBox_SetCurrentIndex(this.h, (int)(index))
}

func (this *QToolBox) SetCurrentWidget(widget *QWidget) {
	QToolBox_SetCurrentWidget(this.h, widget.cPointer())
}

func (this *QToolBox) CurrentChanged(index int) {
	QToolBox_CurrentChanged(this.h, (int)(index))
}
func (this *QToolBox) OnCurrentChanged(slot func(index int)) {
	QToolBox_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_CurrentChanged
func miqt_exec_callback_QToolBox_CurrentChanged(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func QToolBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QToolBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QToolBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QToolBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolBox) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QToolBox_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QToolBox) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_Event
func miqt_exec_callback_QToolBox_Event(self QToolBox, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QToolBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QToolBox) callVirtualBase_ItemInserted(index int) {

	QToolBox_virtualbase_ItemInserted(unsafe.Pointer(this.h), (int)(index))

}
func (this *QToolBox) OnItemInserted(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_ItemInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_ItemInserted
func miqt_exec_callback_QToolBox_ItemInserted(self QToolBox, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QToolBox{h: self}).callVirtualBase_ItemInserted, slotval1)

}

func (this *QToolBox) callVirtualBase_ItemRemoved(index int) {

	QToolBox_virtualbase_ItemRemoved(unsafe.Pointer(this.h), (int)(index))

}
func (this *QToolBox) OnItemRemoved(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_ItemRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_ItemRemoved
func miqt_exec_callback_QToolBox_ItemRemoved(self QToolBox, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QToolBox{h: self}).callVirtualBase_ItemRemoved, slotval1)

}

func (this *QToolBox) callVirtualBase_ShowEvent(e *QShowEvent) {

	QToolBox_virtualbase_ShowEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QToolBox) OnShowEvent(slot func(super func(e *QShowEvent), e *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_ShowEvent
func miqt_exec_callback_QToolBox_ShowEvent(self QToolBox, cb intptr_t, e *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QShowEvent), e *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(e)

	gofunc((&QToolBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QToolBox) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QToolBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QToolBox) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_ChangeEvent
func miqt_exec_callback_QToolBox_ChangeEvent(self QToolBox, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QToolBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QToolBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QToolBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QToolBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_SizeHint
func miqt_exec_callback_QToolBox_SizeHint(self QToolBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QToolBox) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QToolBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QToolBox) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_PaintEvent
func miqt_exec_callback_QToolBox_PaintEvent(self QToolBox, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QToolBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QToolBox) callVirtualBase_InitStyleOption(option *QStyleOptionFrame) {

	QToolBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QToolBox) OnInitStyleOption(slot func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_InitStyleOption
func miqt_exec_callback_QToolBox_InitStyleOption(self QToolBox, cb intptr_t, option *QStyleOptionFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionFrame), option *QStyleOptionFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionFrame(option)

	gofunc((&QToolBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}
