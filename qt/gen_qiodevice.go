package qt
import (
	"unsafe"
)
type QIODevice struct {
			h uintptr
			isSubclass bool
}
// NewQIODevice constructs a new QIODevice object.
			func NewQIODevice() *QIODevice {

g := newQIODevice(QIODevice_new())
				g.isSubclass = true
				return g
			}
// NewQIODevice2 constructs a new QIODevice object.
			func NewQIODevice2(parent *QObject) *QIODevice {

g := newQIODevice(QIODevice_new2(parent.cPointer()))
				g.isSubclass = true
				return g
			}
func (this *QIODevice) MetaObject() *QMetaObject {
return newQMetaObject(QIODevice_MetaObject(this.h))
}
func (this *QIODevice) Metacast(param1 string) unsafe.Pointer {
param1_Cstring := CString(param1)
defer free(unsafe.Pointer(param1_Cstring))
return  (unsafe.Pointer)(QIODevice_Metacast(this.h, param1_Cstring))
}
func QIODevice_Tr(s string) string {
s_Cstring := CString(s)
defer free(unsafe.Pointer(s_Cstring))
var _ms struct_miqt_string =  QIODevice_Tr(s_Cstring)
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
func (this *QIODevice) OpenMode() OpenModeFlag {
return (OpenModeFlag)(QIODevice_OpenMode(this.h))
}
func (this *QIODevice) SetTextModeEnabled(enabled bool)  {
QIODevice_SetTextModeEnabled(this.h, (bool)(enabled))
}
func (this *QIODevice) IsTextModeEnabled() bool {
return (bool)(QIODevice_IsTextModeEnabled(this.h))
}
func (this *QIODevice) IsOpen() bool {
return (bool)(QIODevice_IsOpen(this.h))
}
func (this *QIODevice) IsReadable() bool {
return (bool)(QIODevice_IsReadable(this.h))
}
func (this *QIODevice) IsWritable() bool {
return (bool)(QIODevice_IsWritable(this.h))
}
func (this *QIODevice) IsSequential() bool {
return (bool)(QIODevice_IsSequential(this.h))
}
func (this *QIODevice) ReadChannelCount() int {
return (int)(QIODevice_ReadChannelCount(this.h))
}
func (this *QIODevice) WriteChannelCount() int {
return (int)(QIODevice_WriteChannelCount(this.h))
}
func (this *QIODevice) CurrentReadChannel() int {
return (int)(QIODevice_CurrentReadChannel(this.h))
}
func (this *QIODevice) SetCurrentReadChannel(channel int)  {
QIODevice_SetCurrentReadChannel(this.h, (int)(channel))
}
func (this *QIODevice) CurrentWriteChannel() int {
return (int)(QIODevice_CurrentWriteChannel(this.h))
}
func (this *QIODevice) SetCurrentWriteChannel(channel int)  {
QIODevice_SetCurrentWriteChannel(this.h, (int)(channel))
}
func (this *QIODevice) Open(mode OpenModeFlag) bool {
return (bool)(QIODevice_Open(this.h, (int)(mode)))
}
func (this *QIODevice) Close()  {
QIODevice_Close(this.h)
}
func (this *QIODevice) Pos() int64 {
return (int64)(QIODevice_Pos(this.h))
}
func (this *QIODevice) Size() int64 {
return (int64)(QIODevice_Size(this.h))
}
func (this *QIODevice) Seek(pos int64) bool {
return (bool)(QIODevice_Seek(this.h, (longlong)(pos)))
}
func (this *QIODevice) AtEnd() bool {
return (bool)(QIODevice_AtEnd(this.h))
}
func (this *QIODevice) Reset() bool {
return (bool)(QIODevice_Reset(this.h))
}
func (this *QIODevice) BytesAvailable() int64 {
return (int64)(QIODevice_BytesAvailable(this.h))
}
func (this *QIODevice) BytesToWrite() int64 {
return (int64)(QIODevice_BytesToWrite(this.h))
}
func (this *QIODevice) Read(data string, maxlen int64) int64 {
data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))
return (int64)(QIODevice_Read(this.h, data_Cstring, (longlong)(maxlen)))
}
func (this *QIODevice) ReadWithMaxlen(maxlen int64) []byte {
var _bytearray struct_miqt_string =  QIODevice_ReadWithMaxlen(this.h, (longlong)(maxlen))
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
func (this *QIODevice) ReadAll() []byte {
var _bytearray struct_miqt_string =  QIODevice_ReadAll(this.h)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
func (this *QIODevice) ReadLine(data string, maxlen int64) int64 {
data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))
return (int64)(QIODevice_ReadLine(this.h, data_Cstring, (longlong)(maxlen)))
}
func (this *QIODevice) ReadLine2() []byte {
var _bytearray struct_miqt_string =  QIODevice_ReadLine2(this.h)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
func (this *QIODevice) ReadLineIntoWithBuffer(buffer QSpan<char>) *QByteArrayView {
_goptr := newQByteArrayView(QIODevice_ReadLineIntoWithBuffer(this.h, buffer))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func (this *QIODevice) ReadLineInto2(buffer QSpan<uchar>) *QByteArrayView {
_goptr := newQByteArrayView(QIODevice_ReadLineInto2(this.h, buffer))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func (this *QIODevice) ReadLineInto3(buffer QSpan<std__byte>) *QByteArrayView {
_goptr := newQByteArrayView(QIODevice_ReadLineInto3(this.h, buffer))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func (this *QIODevice) CanReadLine() bool {
return (bool)(QIODevice_CanReadLine(this.h))
}
func (this *QIODevice) StartTransaction()  {
QIODevice_StartTransaction(this.h)
}
func (this *QIODevice) CommitTransaction()  {
QIODevice_CommitTransaction(this.h)
}
func (this *QIODevice) RollbackTransaction()  {
QIODevice_RollbackTransaction(this.h)
}
func (this *QIODevice) IsTransactionStarted() bool {
return (bool)(QIODevice_IsTransactionStarted(this.h))
}
func (this *QIODevice) Write(data string, lenVal int64) int64 {
data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))
return (int64)(QIODevice_Write(this.h, data_Cstring, (longlong)(lenVal)))
}
func (this *QIODevice) WriteWithData(data string) int64 {
data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))
return (int64)(QIODevice_WriteWithData(this.h, data_Cstring))
}
func (this *QIODevice) Write2(data []byte) int64 {
data_alias := struct_miqt_string{}
data_alias.data = (char)(unsafe.Pointer(&data[0]))
data_alias.len = size_t(len(data))
return (int64)(QIODevice_Write2(this.h, data_alias))
}
func (this *QIODevice) Peek(data string, maxlen int64) int64 {
data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))
return (int64)(QIODevice_Peek(this.h, data_Cstring, (longlong)(maxlen)))
}
func (this *QIODevice) PeekWithMaxlen(maxlen int64) []byte {
var _bytearray struct_miqt_string =  QIODevice_PeekWithMaxlen(this.h, (longlong)(maxlen))
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
func (this *QIODevice) Skip(maxSize int64) int64 {
return (int64)(QIODevice_Skip(this.h, (longlong)(maxSize)))
}
func (this *QIODevice) WaitForReadyRead(msecs int) bool {
return (bool)(QIODevice_WaitForReadyRead(this.h, (int)(msecs)))
}
func (this *QIODevice) WaitForBytesWritten(msecs int) bool {
return (bool)(QIODevice_WaitForBytesWritten(this.h, (int)(msecs)))
}
func (this *QIODevice) UngetChar(c int8)  {
QIODevice_UngetChar(this.h, (char)(c))
}
func (this *QIODevice) PutChar(c int8) bool {
return (bool)(QIODevice_PutChar(this.h, (char)(c)))
}
func (this *QIODevice) GetChar(c string) bool {
c_Cstring := CString(c)
defer free(unsafe.Pointer(c_Cstring))
return (bool)(QIODevice_GetChar(this.h, c_Cstring))
}
func (this *QIODevice) ErrorString() string {
var _ms struct_miqt_string =  QIODevice_ErrorString(this.h)
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
func (this *QIODevice) ReadyRead()  {
QIODevice_ReadyRead(this.h)
}
func (this *QIODevice) OnReadyRead(slot func()) {
					QIODevice_connect_ReadyRead(this.h, intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ReadyRead
				func miqt_exec_callback_QIODevice_ReadyRead(cb intptr_t) {
					gofunc, ok := cgo.Handle(cb).Value().(func())
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
					
						
					gofunc( )
				}
func (this *QIODevice) ChannelReadyRead(channel int)  {
QIODevice_ChannelReadyRead(this.h, (int)(channel))
}
func (this *QIODevice) OnChannelReadyRead(slot func(channel int)) {
					QIODevice_connect_ChannelReadyRead(this.h, intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ChannelReadyRead
				func miqt_exec_callback_QIODevice_ChannelReadyRead(cb intptr_t, channel int) {
					gofunc, ok := cgo.Handle(cb).Value().(func(channel int))
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
					// Convert all CABI parameters to Go parameters
slotval1 := (int)(channel)


						
					gofunc(slotval1 )
				}
func (this *QIODevice) BytesWritten(bytes int64)  {
QIODevice_BytesWritten(this.h, (longlong)(bytes))
}
func (this *QIODevice) OnBytesWritten(slot func(bytes int64)) {
					QIODevice_connect_BytesWritten(this.h, intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_BytesWritten
				func miqt_exec_callback_QIODevice_BytesWritten(cb intptr_t, bytes longlong) {
					gofunc, ok := cgo.Handle(cb).Value().(func(bytes int64))
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
					// Convert all CABI parameters to Go parameters
slotval1 := (int64)(bytes)


						
					gofunc(slotval1 )
				}
func (this *QIODevice) ChannelBytesWritten(channel int, bytes int64)  {
QIODevice_ChannelBytesWritten(this.h, (int)(channel), (longlong)(bytes))
}
func (this *QIODevice) OnChannelBytesWritten(slot func(channel int, bytes int64)) {
					QIODevice_connect_ChannelBytesWritten(this.h, intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ChannelBytesWritten
				func miqt_exec_callback_QIODevice_ChannelBytesWritten(cb intptr_t, channel int, bytes longlong) {
					gofunc, ok := cgo.Handle(cb).Value().(func(channel int, bytes int64))
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
					// Convert all CABI parameters to Go parameters
slotval1 := (int)(channel)

slotval2 := (int64)(bytes)


						
					gofunc(slotval1, slotval2 )
				}
func (this *QIODevice) AboutToClose()  {
QIODevice_AboutToClose(this.h)
}
func (this *QIODevice) OnAboutToClose(slot func()) {
					QIODevice_connect_AboutToClose(this.h, intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_AboutToClose
				func miqt_exec_callback_QIODevice_AboutToClose(cb intptr_t) {
					gofunc, ok := cgo.Handle(cb).Value().(func())
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
					
						
					gofunc( )
				}
func (this *QIODevice) ReadChannelFinished()  {
QIODevice_ReadChannelFinished(this.h)
}
func (this *QIODevice) OnReadChannelFinished(slot func()) {
					QIODevice_connect_ReadChannelFinished(this.h, intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ReadChannelFinished
				func miqt_exec_callback_QIODevice_ReadChannelFinished(cb intptr_t) {
					gofunc, ok := cgo.Handle(cb).Value().(func())
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
					
						
					gofunc( )
				}
func QIODevice_Tr2(s string, c string) string {
s_Cstring := CString(s)
defer free(unsafe.Pointer(s_Cstring))
c_Cstring := CString(c)
defer free(unsafe.Pointer(c_Cstring))
var _ms struct_miqt_string =  QIODevice_Tr2(s_Cstring, c_Cstring)
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
func QIODevice_Tr3(s string, c string, n int) string {
s_Cstring := CString(s)
defer free(unsafe.Pointer(s_Cstring))
c_Cstring := CString(c)
defer free(unsafe.Pointer(c_Cstring))
var _ms struct_miqt_string =  QIODevice_Tr3(s_Cstring, c_Cstring, (int)(n))
_ret := GoStringN(_ms.data, int(int64(_ms.len)))
free(unsafe.Pointer(_ms.data))
return _ret}
func (this *QIODevice) ReadLine1(maxlen int64) []byte {
var _bytearray struct_miqt_string =  QIODevice_ReadLine1(this.h, (longlong)(maxlen))
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
func (this *QIODevice) callVirtualBase_MetaObject() *QMetaObject {
					
					return newQMetaObject(QIODevice_virtualbase_MetaObject(unsafe.Pointer(this.h)))

				}
func (this *QIODevice) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_MetaObject
				func miqt_exec_callback_QIODevice_MetaObject(self QIODevice, cb intptr_t) *QMetaObject{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}

virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_MetaObject )

return virtualReturn.cPointer()
}
func (this *QIODevice) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
					param1_Cstring := CString(param1)
defer free(unsafe.Pointer(param1_Cstring))

					return  (unsafe.Pointer)(QIODevice_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))

				}
func (this *QIODevice) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Metacast
				func miqt_exec_callback_QIODevice_Metacast(self QIODevice, cb intptr_t, param1 *const_char) unsafe.Pointer{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
// Convert all CABI parameters to Go parameters
param1_ret :=  param1
slotval1 :=  GoString(param1_ret)
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_Metacast, slotval1 )

return virtualReturn
}
