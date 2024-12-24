package qt

import (
	"unsafe"
)

type QTextStream__RealNumberNotation int

const (
	QTextStream__SmartNotation      QTextStream__RealNumberNotation = 0
	QTextStream__FixedNotation      QTextStream__RealNumberNotation = 1
	QTextStream__ScientificNotation QTextStream__RealNumberNotation = 2
)

type QTextStream__FieldAlignment int

const (
	QTextStream__AlignLeft            QTextStream__FieldAlignment = 0
	QTextStream__AlignRight           QTextStream__FieldAlignment = 1
	QTextStream__AlignCenter          QTextStream__FieldAlignment = 2
	QTextStream__AlignAccountingStyle QTextStream__FieldAlignment = 3
)

type QTextStream__Status int

const (
	QTextStream__Ok              QTextStream__Status = 0
	QTextStream__ReadPastEnd     QTextStream__Status = 1
	QTextStream__ReadCorruptData QTextStream__Status = 2
	QTextStream__WriteFailed     QTextStream__Status = 3
)

type QTextStream__NumberFlag int

const (
	QTextStream__ShowBase        QTextStream__NumberFlag = 1
	QTextStream__ForcePoint      QTextStream__NumberFlag = 2
	QTextStream__ForceSign       QTextStream__NumberFlag = 4
	QTextStream__UppercaseBase   QTextStream__NumberFlag = 8
	QTextStream__UppercaseDigits QTextStream__NumberFlag = 16
)

type QTextStream struct {
	h          uintptr
	isSubclass bool
}

// NewQTextStream constructs a new QTextStream object.
func NewQTextStream() *QTextStream {

	ret := newQTextStream(QTextStream_new())
	ret.isSubclass = true
	return ret
}

// NewQTextStream2 constructs a new QTextStream object.
func NewQTextStream2(device *QIODevice) *QTextStream {

	ret := newQTextStream(QTextStream_new2(device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextStream3 constructs a new QTextStream object.
func NewQTextStream3(array []byte) *QTextStream {
	array_alias := struct_miqt_string{}
	array_alias.data = (char)(unsafe.Pointer(&array[0]))
	array_alias.len = size_t(len(array))

	ret := newQTextStream(QTextStream_new3(array_alias))
	ret.isSubclass = true
	return ret
}

// NewQTextStream4 constructs a new QTextStream object.
func NewQTextStream4(array []byte, openMode OpenMode) *QTextStream {
	array_alias := struct_miqt_string{}
	array_alias.data = (char)(unsafe.Pointer(&array[0]))
	array_alias.len = size_t(len(array))

	ret := newQTextStream(QTextStream_new4(array_alias, openMode))
	ret.isSubclass = true
	return ret
}

func (this *QTextStream) SetEncoding(encoding QStringConverter__Encoding) {
	QTextStream_SetEncoding(this.h, (int)(encoding))
}

func (this *QTextStream) Encoding() QStringConverter__Encoding {
	return (QStringConverter__Encoding)(QTextStream_Encoding(this.h))
}

func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	QTextStream_SetAutoDetectUnicode(this.h, (bool)(enabled))
}

func (this *QTextStream) AutoDetectUnicode() bool {
	return (bool)(QTextStream_AutoDetectUnicode(this.h))
}

func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	QTextStream_SetGenerateByteOrderMark(this.h, (bool)(generate))
}

func (this *QTextStream) GenerateByteOrderMark() bool {
	return (bool)(QTextStream_GenerateByteOrderMark(this.h))
}

func (this *QTextStream) SetLocale(locale *QLocale) {
	QTextStream_SetLocale(this.h, locale.cPointer())
}

func (this *QTextStream) Locale() *QLocale {
	_goptr := newQLocale(QTextStream_Locale(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextStream) SetDevice(device *QIODevice) {
	QTextStream_SetDevice(this.h, device.cPointer())
}

func (this *QTextStream) Device() *QIODevice {
	return newQIODevice(QTextStream_Device(this.h))
}

func (this *QTextStream) String() string {
	var _ms struct_miqt_string = QTextStream_String(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextStream) Status() Status {
	xxxxxxxxx
}

func (this *QTextStream) SetStatus(status Status) {
	QTextStream_SetStatus(this.h, status)
}

func (this *QTextStream) ResetStatus() {
	QTextStream_ResetStatus(this.h)
}

func (this *QTextStream) AtEnd() bool {
	return (bool)(QTextStream_AtEnd(this.h))
}

func (this *QTextStream) Reset() {
	QTextStream_Reset(this.h)
}

func (this *QTextStream) Flush() {
	QTextStream_Flush(this.h)
}

func (this *QTextStream) Seek(pos int64) bool {
	return (bool)(QTextStream_Seek(this.h, (longlong)(pos)))
}

func (this *QTextStream) Pos() int64 {
	return (int64)(QTextStream_Pos(this.h))
}

func (this *QTextStream) SkipWhiteSpace() {
	QTextStream_SkipWhiteSpace(this.h)
}

func (this *QTextStream) ReadLine() string {
	var _ms struct_miqt_string = QTextStream_ReadLine(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextStream) ReadAll() string {
	var _ms struct_miqt_string = QTextStream_ReadAll(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextStream) Read(maxlen int64) string {
	var _ms struct_miqt_string = QTextStream_Read(this.h, (longlong)(maxlen))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextStream) SetFieldAlignment(alignment FieldAlignment) {
	QTextStream_SetFieldAlignment(this.h, alignment)
}

func (this *QTextStream) FieldAlignment() FieldAlignment {
	xxxxxxxxx
}

func (this *QTextStream) SetPadChar(ch QChar) {
	QTextStream_SetPadChar(this.h, ch.cPointer())
}

func (this *QTextStream) PadChar() *QChar {
	_goptr := newQChar(QTextStream_PadChar(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextStream) SetFieldWidth(width int) {
	QTextStream_SetFieldWidth(this.h, (int)(width))
}

func (this *QTextStream) FieldWidth() int {
	return (int)(QTextStream_FieldWidth(this.h))
}

func (this *QTextStream) SetNumberFlags(flags NumberFlags) {
	QTextStream_SetNumberFlags(this.h, flags)
}

func (this *QTextStream) NumberFlags() NumberFlags {
	xxxxxxxxx
}

func (this *QTextStream) SetIntegerBase(base int) {
	QTextStream_SetIntegerBase(this.h, (int)(base))
}

func (this *QTextStream) IntegerBase() int {
	return (int)(QTextStream_IntegerBase(this.h))
}

func (this *QTextStream) SetRealNumberNotation(notation RealNumberNotation) {
	QTextStream_SetRealNumberNotation(this.h, notation)
}

func (this *QTextStream) RealNumberNotation() RealNumberNotation {
	xxxxxxxxx
}

func (this *QTextStream) SetRealNumberPrecision(precision int) {
	QTextStream_SetRealNumberPrecision(this.h, (int)(precision))
}

func (this *QTextStream) RealNumberPrecision() int {
	return (int)(QTextStream_RealNumberPrecision(this.h))
}

func (this *QTextStream) OperatorShiftRight(ch *QChar) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRight(this.h, ch.cPointer()))
}

func (this *QTextStream) OperatorShiftRightWithCh(ch *int8) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithCh(this.h, (*char)(unsafe.Pointer(ch))))
}

func (this *QTextStream) OperatorShiftRightWithShort(i *int16) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithShort(this.h, (*int16_t)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithUnsignedshort(i *uint16) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithUnsignedshort(this.h, (*uint16_t)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithInt(i *int) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithInt(this.h, (*int)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithUnsignedint(i *uint) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithUnsignedint(this.h, (*uint)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithLong(i *int32) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithLong(this.h, (*long)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithUnsignedlong(i *uint32) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithUnsignedlong(this.h, (*ulong)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithQlonglong(i *int64) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithQlonglong(this.h, (*longlong)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithQulonglong(i *uint64) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithQulonglong(this.h, (*ulonglong)(unsafe.Pointer(i))))
}

func (this *QTextStream) OperatorShiftRightWithFloat(f *float32) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithFloat(this.h, (*float)(unsafe.Pointer(f))))
}

func (this *QTextStream) OperatorShiftRightWithDouble(f *float64) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftRightWithDouble(this.h, (*double)(unsafe.Pointer(f))))
}

func (this *QTextStream) OperatorShiftRightWithQString(s string) *QTextStream {
	s_ms := struct_miqt_string{}
	s_ms.data = CString(s)
	s_ms.len = size_t(len(s))
	defer free(unsafe.Pointer(s_ms.data))
	return newQTextStream(QTextStream_OperatorShiftRightWithQString(this.h, s_ms))
}

func (this *QTextStream) OperatorShiftRightWithArray(array []byte) *QTextStream {
	array_alias := struct_miqt_string{}
	array_alias.data = (char)(unsafe.Pointer(&array[0]))
	array_alias.len = size_t(len(array))
	return newQTextStream(QTextStream_OperatorShiftRightWithArray(this.h, array_alias))
}

func (this *QTextStream) OperatorShiftRightWithChar(c string) *QTextStream {
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	return newQTextStream(QTextStream_OperatorShiftRightWithChar(this.h, c_Cstring))
}

func (this *QTextStream) OperatorShiftLeft(ch QChar) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeft(this.h, ch.cPointer()))
}

func (this *QTextStream) OperatorShiftLeftWithCh(ch int8) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithCh(this.h, (char)(ch)))
}

func (this *QTextStream) OperatorShiftLeftWithShort(i int16) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithShort(this.h, (int16_t)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedshort(i uint16) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithUnsignedshort(this.h, (uint16_t)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithInt(i int) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithInt(this.h, (int)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedint(i uint) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithUnsignedint(this.h, (uint)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithLong(i int32) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithLong(this.h, (long)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedlong(i uint32) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithUnsignedlong(this.h, (ulong)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithQlonglong(i int64) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithQlonglong(this.h, (longlong)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithQulonglong(i uint64) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithQulonglong(this.h, (ulonglong)(i)))
}

func (this *QTextStream) OperatorShiftLeftWithFloat(f float32) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithFloat(this.h, (float)(f)))
}

func (this *QTextStream) OperatorShiftLeftWithDouble(f float64) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithDouble(this.h, (double)(f)))
}

func (this *QTextStream) OperatorShiftLeftWithQString(s string) *QTextStream {
	s_ms := struct_miqt_string{}
	s_ms.data = CString(s)
	s_ms.len = size_t(len(s))
	defer free(unsafe.Pointer(s_ms.data))
	return newQTextStream(QTextStream_OperatorShiftLeftWithQString(this.h, s_ms))
}

func (this *QTextStream) OperatorShiftLeftWithArray(array []byte) *QTextStream {
	array_alias := struct_miqt_string{}
	array_alias.data = (char)(unsafe.Pointer(&array[0]))
	array_alias.len = size_t(len(array))
	return newQTextStream(QTextStream_OperatorShiftLeftWithArray(this.h, array_alias))
}

func (this *QTextStream) OperatorShiftLeftWithChar(c string) *QTextStream {
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	return newQTextStream(QTextStream_OperatorShiftLeftWithChar(this.h, c_Cstring))
}

func (this *QTextStream) OperatorShiftLeftWithPtr(ptr unsafe.Pointer) *QTextStream {
	return newQTextStream(QTextStream_OperatorShiftLeftWithPtr(this.h, ptr))
}

func (this *QTextStream) ReadLine1(maxlen int64) string {
	var _ms struct_miqt_string = QTextStream_ReadLine1(this.h, (longlong)(maxlen))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
