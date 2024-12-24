package cbor

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QCborStreamReader__Type byte

const (
	QCborStreamReader__UnsignedInteger QCborStreamReader__Type = 0
	QCborStreamReader__NegativeInteger QCborStreamReader__Type = 32
	QCborStreamReader__ByteString      QCborStreamReader__Type = 64
	QCborStreamReader__ByteArray       QCborStreamReader__Type = 64
	QCborStreamReader__TextString      QCborStreamReader__Type = 96
	QCborStreamReader__String          QCborStreamReader__Type = 96
	QCborStreamReader__Array           QCborStreamReader__Type = 128
	QCborStreamReader__Map             QCborStreamReader__Type = 160
	QCborStreamReader__Tag             QCborStreamReader__Type = 192
	QCborStreamReader__SimpleType      QCborStreamReader__Type = 224
	QCborStreamReader__HalfFloat       QCborStreamReader__Type = 249
	QCborStreamReader__Float16         QCborStreamReader__Type = 249
	QCborStreamReader__Float           QCborStreamReader__Type = 250
	QCborStreamReader__Double          QCborStreamReader__Type = 251
	QCborStreamReader__Invalid         QCborStreamReader__Type = 255
)

type QCborStreamReader__StringResultCode int

const (
	QCborStreamReader__EndOfString QCborStreamReader__StringResultCode = 0
	QCborStreamReader__Ok          QCborStreamReader__StringResultCode = 1
	QCborStreamReader__Error       QCborStreamReader__StringResultCode = -1
)

type QCborStreamReader struct {
	h          uintptr
	isSubclass bool
}

// NewQCborStreamReader constructs a new QCborStreamReader object.
func NewQCborStreamReader() *QCborStreamReader {

	ret := newQCborStreamReader(QCborStreamReader_new())
	ret.isSubclass = true
	return ret
}

// NewQCborStreamReader2 constructs a new QCborStreamReader object.
func NewQCborStreamReader2(data string, lenVal int64) *QCborStreamReader {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	ret := newQCborStreamReader(QCborStreamReader_new2(data_Cstring, (ptrdiff_t)(lenVal)))
	ret.isSubclass = true
	return ret
}

// NewQCborStreamReader3 constructs a new QCborStreamReader object.
func NewQCborStreamReader3(data *byte, lenVal int64) *QCborStreamReader {

	ret := newQCborStreamReader(QCborStreamReader_new3((*uchar)(unsafe.Pointer(data)), (ptrdiff_t)(lenVal)))
	ret.isSubclass = true
	return ret
}

// NewQCborStreamReader4 constructs a new QCborStreamReader object.
func NewQCborStreamReader4(data []byte) *QCborStreamReader {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))

	ret := newQCborStreamReader(QCborStreamReader_new4(data_alias))
	ret.isSubclass = true
	return ret
}

// NewQCborStreamReader5 constructs a new QCborStreamReader object.
func NewQCborStreamReader5(device *qt.QIODevice) *QCborStreamReader {

	ret := newQCborStreamReader(QCborStreamReader_new5((*QIODevice)(device.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QCborStreamReader) SetDevice(device *qt.QIODevice) {
	QCborStreamReader_SetDevice(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QCborStreamReader) Device() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QCborStreamReader_Device(this.h)))
}

func (this *QCborStreamReader) AddData(data []byte) {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	QCborStreamReader_AddData(this.h, data_alias)
}

func (this *QCborStreamReader) AddData2(data string, lenVal int64) {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	QCborStreamReader_AddData2(this.h, data_Cstring, (ptrdiff_t)(lenVal))
}

func (this *QCborStreamReader) AddData3(data *byte, lenVal int64) {
	QCborStreamReader_AddData3(this.h, (*uchar)(unsafe.Pointer(data)), (ptrdiff_t)(lenVal))
}

func (this *QCborStreamReader) Reparse() {
	QCborStreamReader_Reparse(this.h)
}

func (this *QCborStreamReader) Clear() {
	QCborStreamReader_Clear(this.h)
}

func (this *QCborStreamReader) Reset() {
	QCborStreamReader_Reset(this.h)
}

func (this *QCborStreamReader) LastError() *QCborError {
	_goptr := newQCborError(QCborStreamReader_LastError(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborStreamReader) CurrentOffset() int64 {
	return (int64)(QCborStreamReader_CurrentOffset(this.h))
}

func (this *QCborStreamReader) IsValid() bool {
	return (bool)(QCborStreamReader_IsValid(this.h))
}

func (this *QCborStreamReader) ContainerDepth() int {
	return (int)(QCborStreamReader_ContainerDepth(this.h))
}

func (this *QCborStreamReader) ParentContainerType() QCborStreamReader__Type {
	return (QCborStreamReader__Type)(QCborStreamReader_ParentContainerType(this.h))
}

func (this *QCborStreamReader) HasNext() bool {
	return (bool)(QCborStreamReader_HasNext(this.h))
}

func (this *QCborStreamReader) Next() bool {
	return (bool)(QCborStreamReader_Next(this.h))
}

func (this *QCborStreamReader) Type() Type {
	xxxxxxxxx
}

func (this *QCborStreamReader) IsUnsignedInteger() bool {
	return (bool)(QCborStreamReader_IsUnsignedInteger(this.h))
}

func (this *QCborStreamReader) IsNegativeInteger() bool {
	return (bool)(QCborStreamReader_IsNegativeInteger(this.h))
}

func (this *QCborStreamReader) IsInteger() bool {
	return (bool)(QCborStreamReader_IsInteger(this.h))
}

func (this *QCborStreamReader) IsByteArray() bool {
	return (bool)(QCborStreamReader_IsByteArray(this.h))
}

func (this *QCborStreamReader) IsString() bool {
	return (bool)(QCborStreamReader_IsString(this.h))
}

func (this *QCborStreamReader) IsArray() bool {
	return (bool)(QCborStreamReader_IsArray(this.h))
}

func (this *QCborStreamReader) IsMap() bool {
	return (bool)(QCborStreamReader_IsMap(this.h))
}

func (this *QCborStreamReader) IsTag() bool {
	return (bool)(QCborStreamReader_IsTag(this.h))
}

func (this *QCborStreamReader) IsSimpleType() bool {
	return (bool)(QCborStreamReader_IsSimpleType(this.h))
}

func (this *QCborStreamReader) IsFloat16() bool {
	return (bool)(QCborStreamReader_IsFloat16(this.h))
}

func (this *QCborStreamReader) IsFloat() bool {
	return (bool)(QCborStreamReader_IsFloat(this.h))
}

func (this *QCborStreamReader) IsDouble() bool {
	return (bool)(QCborStreamReader_IsDouble(this.h))
}

func (this *QCborStreamReader) IsInvalid() bool {
	return (bool)(QCborStreamReader_IsInvalid(this.h))
}

func (this *QCborStreamReader) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	return (bool)(QCborStreamReader_IsSimpleTypeWithSt(this.h, (uint8_t)(st)))
}

func (this *QCborStreamReader) IsFalse() bool {
	return (bool)(QCborStreamReader_IsFalse(this.h))
}

func (this *QCborStreamReader) IsTrue() bool {
	return (bool)(QCborStreamReader_IsTrue(this.h))
}

func (this *QCborStreamReader) IsBool() bool {
	return (bool)(QCborStreamReader_IsBool(this.h))
}

func (this *QCborStreamReader) IsNull() bool {
	return (bool)(QCborStreamReader_IsNull(this.h))
}

func (this *QCborStreamReader) IsUndefined() bool {
	return (bool)(QCborStreamReader_IsUndefined(this.h))
}

func (this *QCborStreamReader) IsLengthKnown() bool {
	return (bool)(QCborStreamReader_IsLengthKnown(this.h))
}

func (this *QCborStreamReader) Length() uint64 {
	return (uint64)(QCborStreamReader_Length(this.h))
}

func (this *QCborStreamReader) IsContainer() bool {
	return (bool)(QCborStreamReader_IsContainer(this.h))
}

func (this *QCborStreamReader) EnterContainer() bool {
	return (bool)(QCborStreamReader_EnterContainer(this.h))
}

func (this *QCborStreamReader) LeaveContainer() bool {
	return (bool)(QCborStreamReader_LeaveContainer(this.h))
}

func (this *QCborStreamReader) ReadAndAppendToString(dst string) bool {
	dst_ms := struct_miqt_string{}
	dst_ms.data = CString(dst)
	dst_ms.len = size_t(len(dst))
	defer free(unsafe.Pointer(dst_ms.data))
	return (bool)(QCborStreamReader_ReadAndAppendToString(this.h, dst_ms))
}

func (this *QCborStreamReader) ReadAndAppendToUtf8String(dst []byte) bool {
	dst_alias := struct_miqt_string{}
	dst_alias.data = (char)(unsafe.Pointer(&dst[0]))
	dst_alias.len = size_t(len(dst))
	return (bool)(QCborStreamReader_ReadAndAppendToUtf8String(this.h, dst_alias))
}

func (this *QCborStreamReader) ReadAndAppendToByteArray(dst []byte) bool {
	dst_alias := struct_miqt_string{}
	dst_alias.data = (char)(unsafe.Pointer(&dst[0]))
	dst_alias.len = size_t(len(dst))
	return (bool)(QCborStreamReader_ReadAndAppendToByteArray(this.h, dst_alias))
}

func (this *QCborStreamReader) CurrentStringChunkSize() int64 {
	return (int64)(QCborStreamReader_CurrentStringChunkSize(this.h))
}

func (this *QCborStreamReader) ToBool() bool {
	return (bool)(QCborStreamReader_ToBool(this.h))
}

func (this *QCborStreamReader) ToTag() QCborTag {
	return (QCborTag)(QCborStreamReader_ToTag(this.h))
}

func (this *QCborStreamReader) ToUnsignedInteger() uint64 {
	return (uint64)(QCborStreamReader_ToUnsignedInteger(this.h))
}

func (this *QCborStreamReader) ToNegativeInteger() QCborNegativeInteger {
	return (QCborNegativeInteger)(QCborStreamReader_ToNegativeInteger(this.h))
}

func (this *QCborStreamReader) ToSimpleType() QCborSimpleType {
	return (QCborSimpleType)(QCborStreamReader_ToSimpleType(this.h))
}

func (this *QCborStreamReader) ToFloat() float32 {
	return (float32)(QCborStreamReader_ToFloat(this.h))
}

func (this *QCborStreamReader) ToDouble() float64 {
	return (float64)(QCborStreamReader_ToDouble(this.h))
}

func (this *QCborStreamReader) ToInteger() int64 {
	return (int64)(QCborStreamReader_ToInteger(this.h))
}

func (this *QCborStreamReader) ReadAllString() string {
	var _ms struct_miqt_string = QCborStreamReader_ReadAllString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborStreamReader) ReadAllUtf8String() []byte {
	var _bytearray struct_miqt_string = QCborStreamReader_ReadAllUtf8String(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborStreamReader) ReadAllByteArray() []byte {
	var _bytearray struct_miqt_string = QCborStreamReader_ReadAllByteArray(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborStreamReader) Next1(maxRecursion int) bool {
	return (bool)(QCborStreamReader_Next1(this.h, (int)(maxRecursion)))
}
