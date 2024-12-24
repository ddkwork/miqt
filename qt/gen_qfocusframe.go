package qt

import (
	"unsafe"
)

type QFocusFrame struct {
	h          uintptr
	isSubclass bool
}

// NewQFocusFrame constructs a new QFocusFrame object.
func NewQFocusFrame(parent *QWidget) *QFocusFrame {
	ret := newQFocusFrame(QFocusFrame_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFocusFrame2 constructs a new QFocusFrame object.
func NewQFocusFrame2() *QFocusFrame {
	ret := newQFocusFrame(QFocusFrame_new2())
	ret.isSubclass = true
	return ret
}

func (this *QFocusFrame) MetaObject() *QMetaObject {
	return newQMetaObject(QFocusFrame_MetaObject(this.h))
}

func (this *QFocusFrame) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFocusFrame_Metacast(this.h, param1_Cstring))
}

func QFocusFrame_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFocusFrame_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFocusFrame) SetWidget(widget *QWidget) {
	QFocusFrame_SetWidget(this.h, widget.cPointer())
}

func (this *QFocusFrame) Widget() *QWidget {
	return newQWidget(QFocusFrame_Widget(this.h))
}

func QFocusFrame_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFocusFrame_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFocusFrame_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFocusFrame_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFocusFrame) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFocusFrame_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFocusFrame) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFocusFrame_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFocusFrame_MetaObject
func miqt_exec_callback_QFocusFrame_MetaObject(self QFocusFrame, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFocusFrame{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFocusFrame) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFocusFrame_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFocusFrame) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFocusFrame_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFocusFrame_Metacast
func miqt_exec_callback_QFocusFrame_Metacast(self QFocusFrame, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QFocusFrame{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
