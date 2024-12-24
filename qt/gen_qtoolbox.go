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
	g := newQToolBox(QToolBox_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQToolBox2 constructs a new QToolBox object.
func NewQToolBox2() *QToolBox {
	g := newQToolBox(QToolBox_new2())
	g.isSubclass = true
	return g
}

// NewQToolBox3 constructs a new QToolBox object.
func NewQToolBox3(parent *QWidget, f WindowType) *QToolBox {
	g := newQToolBox(QToolBox_new3(parent.cPointer(), (int)(f)))
	g.isSubclass = true
	return g
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

func (this *QToolBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QToolBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QToolBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_MetaObject
func miqt_exec_callback_QToolBox_MetaObject(self QToolBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QToolBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QToolBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QToolBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBox_Metacast
func miqt_exec_callback_QToolBox_Metacast(self QToolBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QToolBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
