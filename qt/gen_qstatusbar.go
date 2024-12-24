package qt

import (
	"unsafe"
)

type QStatusBar struct {
	h          uintptr
	isSubclass bool
}

// NewQStatusBar constructs a new QStatusBar object.
func NewQStatusBar(parent *QWidget) *QStatusBar {
	ret := newQStatusBar(QStatusBar_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStatusBar2 constructs a new QStatusBar object.
func NewQStatusBar2() *QStatusBar {
	ret := newQStatusBar(QStatusBar_new2())
	ret.isSubclass = true
	return ret
}

func (this *QStatusBar) MetaObject() *QMetaObject {
	return newQMetaObject(QStatusBar_MetaObject(this.h))
}

func (this *QStatusBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStatusBar_Metacast(this.h, param1_Cstring))
}

func QStatusBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStatusBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStatusBar) AddWidget(widget *QWidget) {
	QStatusBar_AddWidget(this.h, widget.cPointer())
}

func (this *QStatusBar) InsertWidget(index int, widget *QWidget) int {
	return (int)(QStatusBar_InsertWidget(this.h, (int)(index), widget.cPointer()))
}

func (this *QStatusBar) AddPermanentWidget(widget *QWidget) {
	QStatusBar_AddPermanentWidget(this.h, widget.cPointer())
}

func (this *QStatusBar) InsertPermanentWidget(index int, widget *QWidget) int {
	return (int)(QStatusBar_InsertPermanentWidget(this.h, (int)(index), widget.cPointer()))
}

func (this *QStatusBar) RemoveWidget(widget *QWidget) {
	QStatusBar_RemoveWidget(this.h, widget.cPointer())
}

func (this *QStatusBar) SetSizeGripEnabled(sizeGripEnabled bool) {
	QStatusBar_SetSizeGripEnabled(this.h, (bool)(sizeGripEnabled))
}

func (this *QStatusBar) IsSizeGripEnabled() bool {
	return (bool)(QStatusBar_IsSizeGripEnabled(this.h))
}

func (this *QStatusBar) CurrentMessage() string {
	var _ms struct_miqt_string = QStatusBar_CurrentMessage(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStatusBar) ShowMessage(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStatusBar_ShowMessage(this.h, text_ms)
}

func (this *QStatusBar) ClearMessage() {
	QStatusBar_ClearMessage(this.h)
}

func (this *QStatusBar) MessageChanged(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStatusBar_MessageChanged(this.h, text_ms)
}

func (this *QStatusBar) OnMessageChanged(slot func(text string)) {
	QStatusBar_connect_MessageChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MessageChanged
func miqt_exec_callback_QStatusBar_MessageChanged(cb intptr_t, text struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)
}

func QStatusBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStatusBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStatusBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStatusBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStatusBar) AddWidget2(widget *QWidget, stretch int) {
	QStatusBar_AddWidget2(this.h, widget.cPointer(), (int)(stretch))
}

func (this *QStatusBar) InsertWidget3(index int, widget *QWidget, stretch int) int {
	return (int)(QStatusBar_InsertWidget3(this.h, (int)(index), widget.cPointer(), (int)(stretch)))
}

func (this *QStatusBar) AddPermanentWidget2(widget *QWidget, stretch int) {
	QStatusBar_AddPermanentWidget2(this.h, widget.cPointer(), (int)(stretch))
}

func (this *QStatusBar) InsertPermanentWidget3(index int, widget *QWidget, stretch int) int {
	return (int)(QStatusBar_InsertPermanentWidget3(this.h, (int)(index), widget.cPointer(), (int)(stretch)))
}

func (this *QStatusBar) ShowMessage2(text string, timeout int) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QStatusBar_ShowMessage2(this.h, text_ms, (int)(timeout))
}

func (this *QStatusBar) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QStatusBar_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QStatusBar) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_MetaObject
func miqt_exec_callback_QStatusBar_MetaObject(self QStatusBar, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QStatusBar) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QStatusBar_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QStatusBar) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStatusBar_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStatusBar_Metacast
func miqt_exec_callback_QStatusBar_Metacast(self QStatusBar, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QStatusBar{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
