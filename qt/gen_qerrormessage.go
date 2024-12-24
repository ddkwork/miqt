package qt

import (
	"unsafe"
)

type QErrorMessage struct {
	h          uintptr
	isSubclass bool
}

// NewQErrorMessage constructs a new QErrorMessage object.
func NewQErrorMessage(parent *QWidget) *QErrorMessage {
	ret := newQErrorMessage(QErrorMessage_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQErrorMessage2 constructs a new QErrorMessage object.
func NewQErrorMessage2() *QErrorMessage {
	ret := newQErrorMessage(QErrorMessage_new2())
	ret.isSubclass = true
	return ret
}

func (this *QErrorMessage) MetaObject() *QMetaObject {
	return newQMetaObject(QErrorMessage_MetaObject(this.h))
}

func (this *QErrorMessage) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QErrorMessage_Metacast(this.h, param1_Cstring))
}

func QErrorMessage_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QErrorMessage_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QErrorMessage_QtHandler() *QErrorMessage {
	return newQErrorMessage(QErrorMessage_QtHandler())
}

func (this *QErrorMessage) ShowMessage(message string) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QErrorMessage_ShowMessage(this.h, message_ms)
}

func (this *QErrorMessage) ShowMessage2(message string, typeVal string) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	typeVal_ms := struct_miqt_string{}
	typeVal_ms.data = CString(typeVal)
	typeVal_ms.len = size_t(len(typeVal))
	defer free(unsafe.Pointer(typeVal_ms.data))
	QErrorMessage_ShowMessage2(this.h, message_ms, typeVal_ms)
}

func QErrorMessage_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QErrorMessage_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QErrorMessage_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QErrorMessage_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QErrorMessage) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QErrorMessage_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QErrorMessage) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_MetaObject
func miqt_exec_callback_QErrorMessage_MetaObject(self QErrorMessage, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QErrorMessage) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QErrorMessage_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QErrorMessage) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_Metacast
func miqt_exec_callback_QErrorMessage_Metacast(self QErrorMessage, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
