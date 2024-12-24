package qt

import (
	"unsafe"
)

type QDataStream__Version int

const (
	QDataStream__Qt_1_0                    QDataStream__Version = 1
	QDataStream__Qt_2_0                    QDataStream__Version = 2
	QDataStream__Qt_2_1                    QDataStream__Version = 3
	QDataStream__Qt_3_0                    QDataStream__Version = 4
	QDataStream__Qt_3_1                    QDataStream__Version = 5
	QDataStream__Qt_3_3                    QDataStream__Version = 6
	QDataStream__Qt_4_0                    QDataStream__Version = 7
	QDataStream__Qt_4_1                    QDataStream__Version = 7
	QDataStream__Qt_4_2                    QDataStream__Version = 8
	QDataStream__Qt_4_3                    QDataStream__Version = 9
	QDataStream__Qt_4_4                    QDataStream__Version = 10
	QDataStream__Qt_4_5                    QDataStream__Version = 11
	QDataStream__Qt_4_6                    QDataStream__Version = 12
	QDataStream__Qt_4_7                    QDataStream__Version = 12
	QDataStream__Qt_4_8                    QDataStream__Version = 12
	QDataStream__Qt_4_9                    QDataStream__Version = 12
	QDataStream__Qt_5_0                    QDataStream__Version = 13
	QDataStream__Qt_5_1                    QDataStream__Version = 14
	QDataStream__Qt_5_2                    QDataStream__Version = 15
	QDataStream__Qt_5_3                    QDataStream__Version = 15
	QDataStream__Qt_5_4                    QDataStream__Version = 16
	QDataStream__Qt_5_5                    QDataStream__Version = 16
	QDataStream__Qt_5_6                    QDataStream__Version = 17
	QDataStream__Qt_5_7                    QDataStream__Version = 17
	QDataStream__Qt_5_8                    QDataStream__Version = 17
	QDataStream__Qt_5_9                    QDataStream__Version = 17
	QDataStream__Qt_5_10                   QDataStream__Version = 17
	QDataStream__Qt_5_11                   QDataStream__Version = 17
	QDataStream__Qt_5_12                   QDataStream__Version = 18
	QDataStream__Qt_5_13                   QDataStream__Version = 19
	QDataStream__Qt_5_14                   QDataStream__Version = 19
	QDataStream__Qt_5_15                   QDataStream__Version = 19
	QDataStream__Qt_6_0                    QDataStream__Version = 20
	QDataStream__Qt_6_1                    QDataStream__Version = 20
	QDataStream__Qt_6_2                    QDataStream__Version = 20
	QDataStream__Qt_6_3                    QDataStream__Version = 20
	QDataStream__Qt_6_4                    QDataStream__Version = 20
	QDataStream__Qt_6_5                    QDataStream__Version = 20
	QDataStream__Qt_6_6                    QDataStream__Version = 21
	QDataStream__Qt_6_7                    QDataStream__Version = 22
	QDataStream__Qt_6_8                    QDataStream__Version = 22
	QDataStream__Qt_6_9                    QDataStream__Version = 22
	QDataStream__Qt_DefaultCompiledVersion QDataStream__Version = 22
)

type QDataStream__ByteOrder int

const (
	QDataStream__BigEndian    QDataStream__ByteOrder = 0
	QDataStream__LittleEndian QDataStream__ByteOrder = 1
)

type QDataStream__Status int

const (
	QDataStream__Ok                QDataStream__Status = 0
	QDataStream__ReadPastEnd       QDataStream__Status = 1
	QDataStream__ReadCorruptData   QDataStream__Status = 2
	QDataStream__WriteFailed       QDataStream__Status = 3
	QDataStream__SizeLimitExceeded QDataStream__Status = 4
)

type QDataStream__FloatingPointPrecision int

const (
	QDataStream__SinglePrecision QDataStream__FloatingPointPrecision = 0
	QDataStream__DoublePrecision QDataStream__FloatingPointPrecision = 1
)

type QDataStream struct {
	h          uintptr
	isSubclass bool
}

// NewQDataStream constructs a new QDataStream object.
func NewQDataStream() *QDataStream {

	ret := newQDataStream(QDataStream_new())
	ret.isSubclass = true
	return ret
}

// NewQDataStream2 constructs a new QDataStream object.
func NewQDataStream2(param1 *QIODevice) *QDataStream {

	ret := newQDataStream(QDataStream_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDataStream3 constructs a new QDataStream object.
func NewQDataStream3(param1 []byte) *QDataStream {
	param1_alias := struct_miqt_string{}
	param1_alias.data = (char)(unsafe.Pointer(&param1[0]))
	param1_alias.len = size_t(len(param1))

	ret := newQDataStream(QDataStream_new3(param1_alias))
	ret.isSubclass = true
	return ret
}

func (this *QDataStream) Device() *QIODevice {
	return newQIODevice(QDataStream_Device(this.h))
}

func (this *QDataStream) SetDevice(device *QIODevice) {
	QDataStream_SetDevice(this.h, device.cPointer())
}

func (this *QDataStream) AtEnd() bool {
	return (bool)(QDataStream_AtEnd(this.h))
}

func (this *QDataStream) Status() Status {
	xxxxxxxxx
}

func (this *QDataStream) SetStatus(status Status) {
	QDataStream_SetStatus(this.h, status)
}

func (this *QDataStream) ResetStatus() {
	QDataStream_ResetStatus(this.h)
}

func (this *QDataStream) FloatingPointPrecision() FloatingPointPrecision {
	xxxxxxxxx
}

func (this *QDataStream) SetFloatingPointPrecision(precision FloatingPointPrecision) {
	QDataStream_SetFloatingPointPrecision(this.h, precision)
}

func (this *QDataStream) ByteOrder() ByteOrder {
	xxxxxxxxx
}

func (this *QDataStream) SetByteOrder(byteOrder ByteOrder) {
	QDataStream_SetByteOrder(this.h, byteOrder)
}

func (this *QDataStream) Version() int {
	return (int)(QDataStream_Version(this.h))
}

func (this *QDataStream) SetVersion(version int) {
	QDataStream_SetVersion(this.h, (int)(version))
}

func (this *QDataStream) OperatorShiftRight(i *int8) {
	QDataStream_OperatorShiftRight(this.h, (*char)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQint8(i *int8) {
	QDataStream_OperatorShiftRightWithQint8(this.h, (*schar)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQuint8(i *byte) {
	QDataStream_OperatorShiftRightWithQuint8(this.h, (*uchar)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQint16(i *int16) {
	QDataStream_OperatorShiftRightWithQint16(this.h, (*int16_t)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQuint16(i *uint16) {
	QDataStream_OperatorShiftRightWithQuint16(this.h, (*uint16_t)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQint32(i *int) {
	QDataStream_OperatorShiftRightWithQint32(this.h, (*int)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQuint32(i *uint) {
	QDataStream_OperatorShiftRightWithQuint32(this.h, (*uint)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQint64(i *int64) {
	QDataStream_OperatorShiftRightWithQint64(this.h, (*longlong)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithQuint64(i *uint64) {
	QDataStream_OperatorShiftRightWithQuint64(this.h, (*ulonglong)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithBool(i *bool) {
	QDataStream_OperatorShiftRightWithBool(this.h, (*bool)(unsafe.Pointer(i)))
}

func (this *QDataStream) OperatorShiftRightWithFloat(f *float32) {
	QDataStream_OperatorShiftRightWithFloat(this.h, (*float)(unsafe.Pointer(f)))
}

func (this *QDataStream) OperatorShiftRightWithDouble(f *float64) {
	QDataStream_OperatorShiftRightWithDouble(this.h, (*double)(unsafe.Pointer(f)))
}

func (this *QDataStream) OperatorShiftRightWithStr(str string) {
	str_Cstring := CString(str)
	defer free(unsafe.Pointer(str_Cstring))
	QDataStream_OperatorShiftRightWithStr(this.h, str_Cstring)
}

func (this *QDataStream) OperatorShiftLeft(i int8) {
	QDataStream_OperatorShiftLeft(this.h, (char)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQint8(i int8) {
	QDataStream_OperatorShiftLeftWithQint8(this.h, (schar)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQuint8(i byte) {
	QDataStream_OperatorShiftLeftWithQuint8(this.h, (uchar)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQint16(i int16) {
	QDataStream_OperatorShiftLeftWithQint16(this.h, (int16_t)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQuint16(i uint16) {
	QDataStream_OperatorShiftLeftWithQuint16(this.h, (uint16_t)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQint32(i int) {
	QDataStream_OperatorShiftLeftWithQint32(this.h, (int)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQuint32(i uint) {
	QDataStream_OperatorShiftLeftWithQuint32(this.h, (uint)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQint64(i int64) {
	QDataStream_OperatorShiftLeftWithQint64(this.h, (longlong)(i))
}

func (this *QDataStream) OperatorShiftLeftWithQuint64(i uint64) {
	QDataStream_OperatorShiftLeftWithQuint64(this.h, (ulonglong)(i))
}

func (this *QDataStream) OperatorShiftLeftWithFloat(f float32) {
	QDataStream_OperatorShiftLeftWithFloat(this.h, (float)(f))
}

func (this *QDataStream) OperatorShiftLeftWithDouble(f float64) {
	QDataStream_OperatorShiftLeftWithDouble(this.h, (double)(f))
}

func (this *QDataStream) OperatorShiftLeftWithStr(str string) {
	str_Cstring := CString(str)
	defer free(unsafe.Pointer(str_Cstring))
	QDataStream_OperatorShiftLeftWithStr(this.h, str_Cstring)
}

func (this *QDataStream) ReadBytes(param1 string, lenVal *uint) *QDataStream {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return newQDataStream(QDataStream_ReadBytes(this.h, param1_Cstring, (*uint)(unsafe.Pointer(lenVal))))
}

func (this *QDataStream) ReadBytes2(param1 string, lenVal *int64) *QDataStream {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return newQDataStream(QDataStream_ReadBytes2(this.h, param1_Cstring, (*longlong)(unsafe.Pointer(lenVal))))
}

func (this *QDataStream) ReadRawData(param1 string, lenVal int64) int64 {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (int64)(QDataStream_ReadRawData(this.h, param1_Cstring, (longlong)(lenVal)))
}

func (this *QDataStream) WriteBytes(param1 string, lenVal int64) {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	QDataStream_WriteBytes(this.h, param1_Cstring, (longlong)(lenVal))
}

func (this *QDataStream) WriteRawData(param1 string, lenVal int64) int64 {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (int64)(QDataStream_WriteRawData(this.h, param1_Cstring, (longlong)(lenVal)))
}

func (this *QDataStream) SkipRawData(lenVal int64) int64 {
	return (int64)(QDataStream_SkipRawData(this.h, (longlong)(lenVal)))
}

func (this *QDataStream) StartTransaction() {
	QDataStream_StartTransaction(this.h)
}

func (this *QDataStream) CommitTransaction() bool {
	return (bool)(QDataStream_CommitTransaction(this.h))
}

func (this *QDataStream) RollbackTransaction() {
	QDataStream_RollbackTransaction(this.h)
}

func (this *QDataStream) AbortTransaction() {
	QDataStream_AbortTransaction(this.h)
}

func (this *QDataStream) IsDeviceTransactionStarted() bool {
	return (bool)(QDataStream_IsDeviceTransactionStarted(this.h))
}
