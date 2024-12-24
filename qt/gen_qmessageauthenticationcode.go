package qt
import (
	"unsafe"
)

		type QMessageAuthenticationCode struct {
			h uintptr
			isSubclass bool
}
		
			// NewQMessageAuthenticationCode constructs a new QMessageAuthenticationCode object.
			func NewQMessageAuthenticationCode(method QCryptographicHash__Algorithm) *QMessageAuthenticationCode {
								
				ret := newQMessageAuthenticationCode(QMessageAuthenticationCode_new((int)(method)))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQMessageAuthenticationCode2 constructs a new QMessageAuthenticationCode object.
			func NewQMessageAuthenticationCode2(method QCryptographicHash__Algorithm, key QByteArrayView) *QMessageAuthenticationCode {
								
				ret := newQMessageAuthenticationCode(QMessageAuthenticationCode_new2((int)(method), key.cPointer()))
				ret.isSubclass = true
				return ret
			}
			
			
			func (this *QMessageAuthenticationCode) Swap(other *QMessageAuthenticationCode)  {
				 QMessageAuthenticationCode_Swap(this.h, other.cPointer())
}
			
			func (this *QMessageAuthenticationCode) Reset()  {
				 QMessageAuthenticationCode_Reset(this.h)
}
			
			func (this *QMessageAuthenticationCode) SetKey(key QByteArrayView)  {
				 QMessageAuthenticationCode_SetKey(this.h, key.cPointer())
}
			
			func (this *QMessageAuthenticationCode) AddData(data string, length int64)  {
				data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))
 QMessageAuthenticationCode_AddData(this.h, data_Cstring, (ptrdiff_t)(length))
}
			
			func (this *QMessageAuthenticationCode) AddDataWithData(data QByteArrayView)  {
				 QMessageAuthenticationCode_AddDataWithData(this.h, data.cPointer())
}
			
			func (this *QMessageAuthenticationCode) AddDataWithDevice(device *QIODevice) bool {
				return (bool)(QMessageAuthenticationCode_AddDataWithDevice(this.h, device.cPointer()))
}
			
			func (this *QMessageAuthenticationCode) ResultView() *QByteArrayView {
				_goptr := newQByteArrayView(QMessageAuthenticationCode_ResultView(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QMessageAuthenticationCode) Result() []byte {
				var _bytearray struct_miqt_string =  QMessageAuthenticationCode_Result(this.h)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
			
			func QMessageAuthenticationCode_Hash(message QByteArrayView, key QByteArrayView, method QCryptographicHash__Algorithm) []byte {
				var _bytearray struct_miqt_string =  QMessageAuthenticationCode_Hash(message.cPointer(), key.cPointer(), (int)(method))
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
			
			func QMessageAuthenticationCode_HashInto(buffer QSpan<char>, message QByteArrayView, key QByteArrayView, method QCryptographicHash__Algorithm) *QByteArrayView {
				_goptr := newQByteArrayView(QMessageAuthenticationCode_HashInto(buffer, message.cPointer(), key.cPointer(), (int)(method)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QMessageAuthenticationCode_HashInto2(buffer QSpan<uchar>, message QByteArrayView, key QByteArrayView, method QCryptographicHash__Algorithm) *QByteArrayView {
				_goptr := newQByteArrayView(QMessageAuthenticationCode_HashInto2(buffer, message.cPointer(), key.cPointer(), (int)(method)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QMessageAuthenticationCode_HashInto3(buffer QSpan<std__byte>, message QByteArrayView, key QByteArrayView, method QCryptographicHash__Algorithm) *QByteArrayView {
				_goptr := newQByteArrayView(QMessageAuthenticationCode_HashInto3(buffer, message.cPointer(), key.cPointer(), (int)(method)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QMessageAuthenticationCode_HashInto4(buffer QSpan<char>, messageParts QSpan<const QByteArrayView>, key QByteArrayView, method QCryptographicHash__Algorithm) *QByteArrayView {
				_goptr := newQByteArrayView(QMessageAuthenticationCode_HashInto4(buffer, messageParts, key.cPointer(), (int)(method)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QMessageAuthenticationCode_HashInto5(buffer QSpan<uchar>, messageParts QSpan<const QByteArrayView>, key QByteArrayView, method QCryptographicHash__Algorithm) *QByteArrayView {
				_goptr := newQByteArrayView(QMessageAuthenticationCode_HashInto5(buffer, messageParts, key.cPointer(), (int)(method)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func QMessageAuthenticationCode_HashInto6(buffer QSpan<std__byte>, message QSpan<const QByteArrayView>, key QByteArrayView, method QCryptographicHash__Algorithm) *QByteArrayView {
				_goptr := newQByteArrayView(QMessageAuthenticationCode_HashInto6(buffer, message, key.cPointer(), (int)(method)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			