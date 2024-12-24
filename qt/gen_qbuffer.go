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

func (this *QBuffer) callVirtualBase_Open(openMode OpenMode) bool {

	return (bool)(QBuffer_virtualbase_Open(unsafe.Pointer(this.h), openMode))

}
func (this *QBuffer) OnOpen(slot func(super func(openMode OpenMode) bool, openMode OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_Open
func miqt_exec_callback_QBuffer_Open(self QBuffer, cb intptr_t, openMode OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(openMode OpenMode) bool, openMode OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_Open, slotval1)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_Close() {

	QBuffer_virtualbase_Close(unsafe.Pointer(this.h))

}
func (this *QBuffer) OnClose(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_Close(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_Close
func miqt_exec_callback_QBuffer_Close(self QBuffer, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QBuffer{h: self}).callVirtualBase_Close)

}

func (this *QBuffer) callVirtualBase_Size() int64 {

	return (int64)(QBuffer_virtualbase_Size(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnSize(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_Size
func miqt_exec_callback_QBuffer_Size(self QBuffer, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_Size)

	return (longlong)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_Pos() int64 {

	return (int64)(QBuffer_virtualbase_Pos(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnPos(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_Pos(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_Pos
func miqt_exec_callback_QBuffer_Pos(self QBuffer, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_Pos)

	return (longlong)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_Seek(off int64) bool {

	return (bool)(QBuffer_virtualbase_Seek(unsafe.Pointer(this.h), (longlong)(off)))

}
func (this *QBuffer) OnSeek(slot func(super func(off int64) bool, off int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_Seek(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_Seek
func miqt_exec_callback_QBuffer_Seek(self QBuffer, cb intptr_t, off longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(off int64) bool, off int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(off)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_Seek, slotval1)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_AtEnd() bool {

	return (bool)(QBuffer_virtualbase_AtEnd(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnAtEnd(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_AtEnd(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_AtEnd
func miqt_exec_callback_QBuffer_AtEnd(self QBuffer, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_AtEnd)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_CanReadLine() bool {

	return (bool)(QBuffer_virtualbase_CanReadLine(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnCanReadLine(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_CanReadLine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_CanReadLine
func miqt_exec_callback_QBuffer_CanReadLine(self QBuffer, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_CanReadLine)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_ConnectNotify(param1 *QMetaMethod) {

	QBuffer_virtualbase_ConnectNotify(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QBuffer) OnConnectNotify(slot func(super func(param1 *QMetaMethod), param1 *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_ConnectNotify
func miqt_exec_callback_QBuffer_ConnectNotify(self QBuffer, cb intptr_t, param1 *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMetaMethod), param1 *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(param1)

	gofunc((&QBuffer{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QBuffer) callVirtualBase_DisconnectNotify(param1 *QMetaMethod) {

	QBuffer_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QBuffer) OnDisconnectNotify(slot func(super func(param1 *QMetaMethod), param1 *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_DisconnectNotify
func miqt_exec_callback_QBuffer_DisconnectNotify(self QBuffer, cb intptr_t, param1 *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMetaMethod), param1 *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(param1)

	gofunc((&QBuffer{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

func (this *QBuffer) callVirtualBase_ReadData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QBuffer_virtualbase_ReadData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QBuffer) OnReadData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_ReadData
func miqt_exec_callback_QBuffer_ReadData(self QBuffer, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_ReadData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_WriteData(data string, lenVal int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QBuffer_virtualbase_WriteData(unsafe.Pointer(this.h), data_Cstring, (longlong)(lenVal)))

}
func (this *QBuffer) OnWriteData(slot func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_WriteData
func miqt_exec_callback_QBuffer_WriteData(self QBuffer, cb intptr_t, data *const_char, lenVal longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(lenVal)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_WriteData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_IsSequential() bool {

	return (bool)(QBuffer_virtualbase_IsSequential(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnIsSequential(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_IsSequential
func miqt_exec_callback_QBuffer_IsSequential(self QBuffer, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_IsSequential)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_Reset() bool {

	return (bool)(QBuffer_virtualbase_Reset(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnReset(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_Reset
func miqt_exec_callback_QBuffer_Reset(self QBuffer, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_Reset)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_BytesAvailable() int64 {

	return (int64)(QBuffer_virtualbase_BytesAvailable(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnBytesAvailable(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_BytesAvailable(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_BytesAvailable
func miqt_exec_callback_QBuffer_BytesAvailable(self QBuffer, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_BytesAvailable)

	return (longlong)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_BytesToWrite() int64 {

	return (int64)(QBuffer_virtualbase_BytesToWrite(unsafe.Pointer(this.h)))

}
func (this *QBuffer) OnBytesToWrite(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_BytesToWrite(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_BytesToWrite
func miqt_exec_callback_QBuffer_BytesToWrite(self QBuffer, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_BytesToWrite)

	return (longlong)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_WaitForReadyRead(msecs int) bool {

	return (bool)(QBuffer_virtualbase_WaitForReadyRead(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QBuffer) OnWaitForReadyRead(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_WaitForReadyRead(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_WaitForReadyRead
func miqt_exec_callback_QBuffer_WaitForReadyRead(self QBuffer, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_WaitForReadyRead, slotval1)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_WaitForBytesWritten(msecs int) bool {

	return (bool)(QBuffer_virtualbase_WaitForBytesWritten(unsafe.Pointer(this.h), (int)(msecs)))

}
func (this *QBuffer) OnWaitForBytesWritten(slot func(super func(msecs int) bool, msecs int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_WaitForBytesWritten(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_WaitForBytesWritten
func miqt_exec_callback_QBuffer_WaitForBytesWritten(self QBuffer, cb intptr_t, msecs int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(msecs)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_WaitForBytesWritten, slotval1)

	return (bool)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_ReadLineData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QBuffer_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QBuffer) OnReadLineData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_ReadLineData
func miqt_exec_callback_QBuffer_ReadLineData(self QBuffer, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QBuffer) callVirtualBase_SkipData(maxSize int64) int64 {

	return (int64)(QBuffer_virtualbase_SkipData(unsafe.Pointer(this.h), (longlong)(maxSize)))

}
func (this *QBuffer) OnSkipData(slot func(super func(maxSize int64) int64, maxSize int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBuffer_override_virtual_SkipData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBuffer_SkipData
func miqt_exec_callback_QBuffer_SkipData(self QBuffer, cb intptr_t, maxSize longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(maxSize int64) int64, maxSize int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(maxSize)

	virtualReturn := gofunc((&QBuffer{h: self}).callVirtualBase_SkipData, slotval1)

	return (longlong)(virtualReturn)

}
