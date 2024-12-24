package qt

import (
	"unsafe"
)

type QBuffer struct {
	h          uintptr
	isSubclass bool
}

// NewQBuffer constructs a new QBuffer object.
func NewQBuffer() *QBuffer {
	ret := newQBuffer(QBuffer_new())
	ret.isSubclass = true
	return ret
}

// NewQBuffer2 constructs a new QBuffer object.
func NewQBuffer2(parent *QObject) *QBuffer {
	ret := newQBuffer(QBuffer_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QBuffer) MetaObject() *QMetaObject {
	return newQMetaObject(QBuffer_MetaObject(this.h))
}

func (this *QBuffer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QBuffer_Metacast(this.h, param1_Cstring))
}

func QBuffer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QBuffer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QBuffer) Buffer() []byte {
	var _bytearray struct_miqt_string = QBuffer_Buffer(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QBuffer) Buffer2() []byte {
	var _bytearray struct_miqt_string = QBuffer_Buffer2(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QBuffer) SetData(data []byte) {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	QBuffer_SetData(this.h, data_alias)
}

func (this *QBuffer) SetData2(data string, lenVal int64) {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	QBuffer_SetData2(this.h, data_Cstring, (ptrdiff_t)(lenVal))
}

func (this *QBuffer) Data() []byte {
	var _bytearray struct_miqt_string = QBuffer_Data(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QBuffer) Open(openMode OpenMode) bool {
	return (bool)(QBuffer_Open(this.h, openMode))
}

func (this *QBuffer) Close() {
	QBuffer_Close(this.h)
}

func (this *QBuffer) Size() int64 {
	return (int64)(QBuffer_Size(this.h))
}

func (this *QBuffer) Pos() int64 {
	return (int64)(QBuffer_Pos(this.h))
}

func (this *QBuffer) Seek(off int64) bool {
	return (bool)(QBuffer_Seek(this.h, (longlong)(off)))
}

func (this *QBuffer) AtEnd() bool {
	return (bool)(QBuffer_AtEnd(this.h))
}

func (this *QBuffer) CanReadLine() bool {
	return (bool)(QBuffer_CanReadLine(this.h))
}

func QBuffer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QBuffer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QBuffer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QBuffer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QBuffer) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QBuffer_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QBuffer) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_MetaObject
func miqt_exec_callback_QBuffer_MetaObject(self QBuffer, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QBuffer) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QBuffer_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QBuffer) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_Metacast
func miqt_exec_callback_QBuffer_Metacast(self QBuffer, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
