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
								
				ret := newQIODevice(QIODevice_new())
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQIODevice2 constructs a new QIODevice object.
			func NewQIODevice2(parent *QObject) *QIODevice {
								
				ret := newQIODevice(QIODevice_new2(parent.cPointer()))
				ret.isSubclass = true
				return ret
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
			
				func (this *QIODevice) callVirtualBase_IsSequential() bool {
					
					return (bool)(QIODevice_virtualbase_IsSequential(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnIsSequential(slot func(super func() bool) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_IsSequential
				func miqt_exec_callback_QIODevice_IsSequential(self QIODevice, cb intptr_t) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_IsSequential )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_Open(mode OpenModeFlag) bool {
					
					return (bool)(QIODevice_virtualbase_Open(unsafe.Pointer(this.h), (int)(mode)))

				}
			func (this *QIODevice) OnOpen(slot func(super func(mode OpenModeFlag) bool, mode OpenModeFlag) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Open
				func miqt_exec_callback_QIODevice_Open(self QIODevice, cb intptr_t, mode int) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(mode OpenModeFlag) bool, mode OpenModeFlag) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := (OpenModeFlag)(mode)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_Open, slotval1 )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_Close()  {
					
					 QIODevice_virtualbase_Close(unsafe.Pointer(this.h))

				}
			func (this *QIODevice) OnClose(slot func(super func() ) ) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Close(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Close
				func miqt_exec_callback_QIODevice_Close(self QIODevice, cb intptr_t) {
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() ) )
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
gofunc((&QIODevice{h: self}).callVirtualBase_Close )

				}
			
				func (this *QIODevice) callVirtualBase_Pos() int64 {
					
					return (int64)(QIODevice_virtualbase_Pos(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnPos(slot func(super func() int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Pos(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Pos
				func miqt_exec_callback_QIODevice_Pos(self QIODevice, cb intptr_t) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_Pos )

return (longlong)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_Size() int64 {
					
					return (int64)(QIODevice_virtualbase_Size(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnSize(slot func(super func() int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Size
				func miqt_exec_callback_QIODevice_Size(self QIODevice, cb intptr_t) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_Size )

return (longlong)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_Seek(pos int64) bool {
					
					return (bool)(QIODevice_virtualbase_Seek(unsafe.Pointer(this.h), (longlong)(pos)))

				}
			func (this *QIODevice) OnSeek(slot func(super func(pos int64) bool, pos int64) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Seek(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Seek
				func miqt_exec_callback_QIODevice_Seek(self QIODevice, cb intptr_t, pos longlong) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos int64) bool, pos int64) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := (int64)(pos)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_Seek, slotval1 )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_AtEnd() bool {
					
					return (bool)(QIODevice_virtualbase_AtEnd(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnAtEnd(slot func(super func() bool) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_AtEnd(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_AtEnd
				func miqt_exec_callback_QIODevice_AtEnd(self QIODevice, cb intptr_t) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_AtEnd )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_Reset() bool {
					
					return (bool)(QIODevice_virtualbase_Reset(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnReset(slot func(super func() bool) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Reset
				func miqt_exec_callback_QIODevice_Reset(self QIODevice, cb intptr_t) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_Reset )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_BytesAvailable() int64 {
					
					return (int64)(QIODevice_virtualbase_BytesAvailable(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnBytesAvailable(slot func(super func() int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_BytesAvailable(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_BytesAvailable
				func miqt_exec_callback_QIODevice_BytesAvailable(self QIODevice, cb intptr_t) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_BytesAvailable )

return (longlong)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_BytesToWrite() int64 {
					
					return (int64)(QIODevice_virtualbase_BytesToWrite(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnBytesToWrite(slot func(super func() int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_BytesToWrite(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_BytesToWrite
				func miqt_exec_callback_QIODevice_BytesToWrite(self QIODevice, cb intptr_t) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_BytesToWrite )

return (longlong)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_CanReadLine() bool {
					
					return (bool)(QIODevice_virtualbase_CanReadLine(unsafe.Pointer(this.h)))

				}
			func (this *QIODevice) OnCanReadLine(slot func(super func() bool) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_CanReadLine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_CanReadLine
				func miqt_exec_callback_QIODevice_CanReadLine(self QIODevice, cb intptr_t) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			
virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_CanReadLine )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_WaitForReadyRead(msecs int) bool {
					
					return (bool)(QIODevice_virtualbase_WaitForReadyRead(unsafe.Pointer(this.h), (int)(msecs)))

				}
			func (this *QIODevice) OnWaitForReadyRead(slot func(super func(msecs int) bool, msecs int) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_WaitForReadyRead(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_WaitForReadyRead
				func miqt_exec_callback_QIODevice_WaitForReadyRead(self QIODevice, cb intptr_t, msecs int) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := (int)(msecs)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_WaitForReadyRead, slotval1 )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_WaitForBytesWritten(msecs int) bool {
					
					return (bool)(QIODevice_virtualbase_WaitForBytesWritten(unsafe.Pointer(this.h), (int)(msecs)))

				}
			func (this *QIODevice) OnWaitForBytesWritten(slot func(super func(msecs int) bool, msecs int) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_WaitForBytesWritten(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_WaitForBytesWritten
				func miqt_exec_callback_QIODevice_WaitForBytesWritten(self QIODevice, cb intptr_t, msecs int) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(msecs int) bool, msecs int) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := (int)(msecs)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_WaitForBytesWritten, slotval1 )

return (bool)(virtualReturn)

				}
			func (this *QIODevice) OnReadData(slot func(data string, maxlen int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ReadData
				func miqt_exec_callback_QIODevice_ReadData(self QIODevice, cb intptr_t, data *char, maxlen longlong) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(data string, maxlen int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
data_ret :=  data
slotval1 :=  GoString(data_ret)

slotval2 := (int64)(maxlen)


virtualReturn := gofunc(slotval1, slotval2 )

return (longlong)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_ReadLineData(data string, maxlen int64) int64 {
					data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))

					return (int64)(QIODevice_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

				}
			func (this *QIODevice) OnReadLineData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ReadLineData
				func miqt_exec_callback_QIODevice_ReadLineData(self QIODevice, cb intptr_t, data *char, maxlen longlong) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
data_ret :=  data
slotval1 :=  GoString(data_ret)

slotval2 := (int64)(maxlen)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2 )

return (longlong)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_SkipData(maxSize int64) int64 {
					
					return (int64)(QIODevice_virtualbase_SkipData(unsafe.Pointer(this.h), (longlong)(maxSize)))

				}
			func (this *QIODevice) OnSkipData(slot func(super func(maxSize int64) int64, maxSize int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_SkipData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_SkipData
				func miqt_exec_callback_QIODevice_SkipData(self QIODevice, cb intptr_t, maxSize longlong) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(maxSize int64) int64, maxSize int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := (int64)(maxSize)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_SkipData, slotval1 )

return (longlong)(virtualReturn)

				}
			func (this *QIODevice) OnWriteData(slot func(data string, lenVal int64) int64) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_WriteData
				func miqt_exec_callback_QIODevice_WriteData(self QIODevice, cb intptr_t, data *const_char, lenVal longlong) longlong{
					gofunc, ok := cgo.Handle(cb).Value().(func(data string, lenVal int64) int64)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
data_ret :=  data
slotval1 :=  GoString(data_ret)

slotval2 := (int64)(lenVal)


virtualReturn := gofunc(slotval1, slotval2 )

return (longlong)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_Event(event *QEvent) bool {
					
					return (bool)(QIODevice_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

				}
			func (this *QIODevice) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_Event
				func miqt_exec_callback_QIODevice_Event(self QIODevice, cb intptr_t, event *QEvent) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQEvent(event)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_Event, slotval1 )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {
					
					return (bool)(QIODevice_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

				}
			func (this *QIODevice) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_EventFilter
				func miqt_exec_callback_QIODevice_EventFilter(self QIODevice, cb intptr_t, watched *QObject, event *QEvent) bool{
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQObject(watched)

slotval2 := newQEvent(event)


virtualReturn := gofunc((&QIODevice{h: self}).callVirtualBase_EventFilter, slotval1, slotval2 )

return (bool)(virtualReturn)

				}
			
				func (this *QIODevice) callVirtualBase_TimerEvent(event *QTimerEvent)  {
					
					 QIODevice_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

				}
			func (this *QIODevice) OnTimerEvent(slot func(super func(event *QTimerEvent) , event *QTimerEvent) ) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_TimerEvent
				func miqt_exec_callback_QIODevice_TimerEvent(self QIODevice, cb intptr_t, event *QTimerEvent) {
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent) , event *QTimerEvent) )
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQTimerEvent(event)


gofunc((&QIODevice{h: self}).callVirtualBase_TimerEvent, slotval1 )

				}
			
				func (this *QIODevice) callVirtualBase_ChildEvent(event *QChildEvent)  {
					
					 QIODevice_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

				}
			func (this *QIODevice) OnChildEvent(slot func(super func(event *QChildEvent) , event *QChildEvent) ) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ChildEvent
				func miqt_exec_callback_QIODevice_ChildEvent(self QIODevice, cb intptr_t, event *QChildEvent) {
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent) , event *QChildEvent) )
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQChildEvent(event)


gofunc((&QIODevice{h: self}).callVirtualBase_ChildEvent, slotval1 )

				}
			
				func (this *QIODevice) callVirtualBase_CustomEvent(event *QEvent)  {
					
					 QIODevice_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

				}
			func (this *QIODevice) OnCustomEvent(slot func(super func(event *QEvent) , event *QEvent) ) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_CustomEvent
				func miqt_exec_callback_QIODevice_CustomEvent(self QIODevice, cb intptr_t, event *QEvent) {
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) , event *QEvent) )
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQEvent(event)


gofunc((&QIODevice{h: self}).callVirtualBase_CustomEvent, slotval1 )

				}
			
				func (this *QIODevice) callVirtualBase_ConnectNotify(signal *QMetaMethod)  {
					
					 QIODevice_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

				}
			func (this *QIODevice) OnConnectNotify(slot func(super func(signal *QMetaMethod) , signal *QMetaMethod) ) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_ConnectNotify
				func miqt_exec_callback_QIODevice_ConnectNotify(self QIODevice, cb intptr_t, signal *QMetaMethod) {
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod) , signal *QMetaMethod) )
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQMetaMethod(signal)


gofunc((&QIODevice{h: self}).callVirtualBase_ConnectNotify, slotval1 )

				}
			
				func (this *QIODevice) callVirtualBase_DisconnectNotify(signal *QMetaMethod)  {
					
					 QIODevice_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

				}
			func (this *QIODevice) OnDisconnectNotify(slot func(super func(signal *QMetaMethod) , signal *QMetaMethod) ) {
					if ! this.isSubclass {
						panic("miqt: can only override virtual methods for directly constructed types")
					}
					QIODevice_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)) )
				}
				
				//export miqt_exec_callback_QIODevice_DisconnectNotify
				func miqt_exec_callback_QIODevice_DisconnectNotify(self QIODevice, cb intptr_t, signal *QMetaMethod) {
					gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod) , signal *QMetaMethod) )
					if !ok {
						panic("miqt: callback of non-callback type (heap corruption?)")
					}
				
			// Convert all CABI parameters to Go parameters
slotval1 := newQMetaMethod(signal)


gofunc((&QIODevice{h: self}).callVirtualBase_DisconnectNotify, slotval1 )

				}
			