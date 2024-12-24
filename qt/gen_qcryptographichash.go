package qt
import (
	"unsafe"
)
type QCryptographicHash__Algorithm int
const (
QCryptographicHash__Md4 QCryptographicHash__Algorithm = 0
QCryptographicHash__Md5 QCryptographicHash__Algorithm = 1
QCryptographicHash__Sha1 QCryptographicHash__Algorithm = 2
QCryptographicHash__Sha224 QCryptographicHash__Algorithm = 3
QCryptographicHash__Sha256 QCryptographicHash__Algorithm = 4
QCryptographicHash__Sha384 QCryptographicHash__Algorithm = 5
QCryptographicHash__Sha512 QCryptographicHash__Algorithm = 6
QCryptographicHash__Keccak_224 QCryptographicHash__Algorithm = 7
QCryptographicHash__Keccak_256 QCryptographicHash__Algorithm = 8
QCryptographicHash__Keccak_384 QCryptographicHash__Algorithm = 9
QCryptographicHash__Keccak_512 QCryptographicHash__Algorithm = 10
QCryptographicHash__RealSha3_224 QCryptographicHash__Algorithm = 11
QCryptographicHash__RealSha3_256 QCryptographicHash__Algorithm = 12
QCryptographicHash__RealSha3_384 QCryptographicHash__Algorithm = 13
QCryptographicHash__RealSha3_512 QCryptographicHash__Algorithm = 14
QCryptographicHash__Sha3_224 QCryptographicHash__Algorithm = 11
QCryptographicHash__Sha3_256 QCryptographicHash__Algorithm = 12
QCryptographicHash__Sha3_384 QCryptographicHash__Algorithm = 13
QCryptographicHash__Sha3_512 QCryptographicHash__Algorithm = 14
QCryptographicHash__Blake2b_160 QCryptographicHash__Algorithm = 15
QCryptographicHash__Blake2b_256 QCryptographicHash__Algorithm = 16
QCryptographicHash__Blake2b_384 QCryptographicHash__Algorithm = 17
QCryptographicHash__Blake2b_512 QCryptographicHash__Algorithm = 18
QCryptographicHash__Blake2s_128 QCryptographicHash__Algorithm = 19
QCryptographicHash__Blake2s_160 QCryptographicHash__Algorithm = 20
QCryptographicHash__Blake2s_224 QCryptographicHash__Algorithm = 21
QCryptographicHash__Blake2s_256 QCryptographicHash__Algorithm = 22
QCryptographicHash__NumAlgorithms QCryptographicHash__Algorithm = 23
)
type QCryptographicHash struct {
			h uintptr
			isSubclass bool
}
// NewQCryptographicHash constructs a new QCryptographicHash object.
			func NewQCryptographicHash(method Algorithm) *QCryptographicHash {

g := newQCryptographicHash(QCryptographicHash_new(method))
				g.isSubclass = true
				return g
			}
func (this *QCryptographicHash) Swap(other *QCryptographicHash)  {
QCryptographicHash_Swap(this.h, other.cPointer())
}
func (this *QCryptographicHash) Reset()  {
QCryptographicHash_Reset(this.h)
}
func (this *QCryptographicHash) Algorithm() Algorithm {
xxxxxxxxx}
func (this *QCryptographicHash) AddData(data string, length int64)  {
data_Cstring := CString(data)
defer free(unsafe.Pointer(data_Cstring))
 QCryptographicHash_AddData(this.h, data_Cstring, (ptrdiff_t)(length))
}
func (this *QCryptographicHash) AddDataWithData(data QByteArrayView)  {
QCryptographicHash_AddDataWithData(this.h, data.cPointer())
}
func (this *QCryptographicHash) AddDataWithDevice(device *QIODevice) bool {
return (bool)(QCryptographicHash_AddDataWithDevice(this.h, device.cPointer()))
}
func (this *QCryptographicHash) Result() []byte {
var _bytearray struct_miqt_string =  QCryptographicHash_Result(this.h)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
func (this *QCryptographicHash) ResultView() *QByteArrayView {
_goptr := newQByteArrayView(QCryptographicHash_ResultView(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func QCryptographicHash_Hash(data QByteArrayView, method Algorithm) []byte {
var _bytearray struct_miqt_string =  QCryptographicHash_Hash(data.cPointer(), method)
_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
free(unsafe.Pointer(_bytearray.data))
return _ret}
func QCryptographicHash_HashInto(buffer QSpan<char>, data QByteArrayView, method Algorithm) *QByteArrayView {
_goptr := newQByteArrayView(QCryptographicHash_HashInto(buffer, data.cPointer(), method))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func QCryptographicHash_HashInto2(buffer QSpan<uchar>, data QByteArrayView, method Algorithm) *QByteArrayView {
_goptr := newQByteArrayView(QCryptographicHash_HashInto2(buffer, data.cPointer(), method))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func QCryptographicHash_HashInto3(buffer QSpan<std__byte>, data QByteArrayView, method Algorithm) *QByteArrayView {
_goptr := newQByteArrayView(QCryptographicHash_HashInto3(buffer, data.cPointer(), method))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func QCryptographicHash_HashInto4(buffer QSpan<char>, data QSpan<const QByteArrayView>, method Algorithm) *QByteArrayView {
_goptr := newQByteArrayView(QCryptographicHash_HashInto4(buffer, data, method))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func QCryptographicHash_HashInto5(buffer QSpan<uchar>, data QSpan<const QByteArrayView>, method Algorithm) *QByteArrayView {
_goptr := newQByteArrayView(QCryptographicHash_HashInto5(buffer, data, method))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func QCryptographicHash_HashInto6(buffer QSpan<std__byte>, data QSpan<const QByteArrayView>, method Algorithm) *QByteArrayView {
_goptr := newQByteArrayView(QCryptographicHash_HashInto6(buffer, data, method))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
func QCryptographicHash_HashLength(method Algorithm) int {
return (int)(QCryptographicHash_HashLength(method))
}
func QCryptographicHash_SupportsAlgorithm(method Algorithm) bool {
return (bool)(QCryptographicHash_SupportsAlgorithm(method))
}
