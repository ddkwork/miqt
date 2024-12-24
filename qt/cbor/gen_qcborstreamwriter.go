package cbor

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QCborStreamWriter struct {
	h          uintptr
	isSubclass bool
}

// NewQCborStreamWriter constructs a new QCborStreamWriter object.
func NewQCborStreamWriter(device *qt.QIODevice) *QCborStreamWriter {

	ret := newQCborStreamWriter(QCborStreamWriter_new((*QIODevice)(device.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QCborStreamWriter) SetDevice(device *qt.QIODevice) {
	QCborStreamWriter_SetDevice(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QCborStreamWriter) Device() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QCborStreamWriter_Device(this.h)))
}

func (this *QCborStreamWriter) Append(u uint64) {
	QCborStreamWriter_Append(this.h, (ulonglong)(u))
}

func (this *QCborStreamWriter) AppendWithQint64(i int64) {
	QCborStreamWriter_AppendWithQint64(this.h, (longlong)(i))
}

func (this *QCborStreamWriter) AppendWithQCborNegativeInteger(n QCborNegativeInteger) {
	QCborStreamWriter_AppendWithQCborNegativeInteger(this.h, (uint64_t)(n))
}

func (this *QCborStreamWriter) AppendWithBa(ba []byte) {
	ba_alias := struct_miqt_string{}
	ba_alias.data = (char)(unsafe.Pointer(&ba[0]))
	ba_alias.len = size_t(len(ba))
	QCborStreamWriter_AppendWithBa(this.h, ba_alias)
}

func (this *QCborStreamWriter) AppendWithTag(tag QCborTag) {
	QCborStreamWriter_AppendWithTag(this.h, (uint64_t)(tag))
}

func (this *QCborStreamWriter) Append3(tag QCborKnownTags) {
	QCborStreamWriter_Append3(this.h, (int)(tag))
}

func (this *QCborStreamWriter) AppendWithSt(st QCborSimpleType) {
	QCborStreamWriter_AppendWithSt(this.h, (uint8_t)(st))
}

func (this *QCborStreamWriter) AppendWithFloat(f float32) {
	QCborStreamWriter_AppendWithFloat(this.h, (float)(f))
}

func (this *QCborStreamWriter) AppendWithDouble(d float64) {
	QCborStreamWriter_AppendWithDouble(this.h, (double)(d))
}

func (this *QCborStreamWriter) AppendByteString(data string, lenVal int64) {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	QCborStreamWriter_AppendByteString(this.h, data_Cstring, (ptrdiff_t)(lenVal))
}

func (this *QCborStreamWriter) AppendTextString(utf8 string, lenVal int64) {
	utf8_Cstring := CString(utf8)
	defer free(unsafe.Pointer(utf8_Cstring))
	QCborStreamWriter_AppendTextString(this.h, utf8_Cstring, (ptrdiff_t)(lenVal))
}

func (this *QCborStreamWriter) AppendWithBool(b bool) {
	QCborStreamWriter_AppendWithBool(this.h, (bool)(b))
}

func (this *QCborStreamWriter) AppendNull() {
	QCborStreamWriter_AppendNull(this.h)
}

func (this *QCborStreamWriter) AppendUndefined() {
	QCborStreamWriter_AppendUndefined(this.h)
}

func (this *QCborStreamWriter) AppendWithInt(i int) {
	QCborStreamWriter_AppendWithInt(this.h, (int)(i))
}

func (this *QCborStreamWriter) AppendWithUint(u uint) {
	QCborStreamWriter_AppendWithUint(this.h, (uint)(u))
}

func (this *QCborStreamWriter) Append4(str string) {
	str_Cstring := CString(str)
	defer free(unsafe.Pointer(str_Cstring))
	QCborStreamWriter_Append4(this.h, str_Cstring)
}

func (this *QCborStreamWriter) StartArray() {
	QCborStreamWriter_StartArray(this.h)
}

func (this *QCborStreamWriter) StartArrayWithCount(count uint64) {
	QCborStreamWriter_StartArrayWithCount(this.h, (ulonglong)(count))
}

func (this *QCborStreamWriter) EndArray() bool {
	return (bool)(QCborStreamWriter_EndArray(this.h))
}

func (this *QCborStreamWriter) StartMap() {
	QCborStreamWriter_StartMap(this.h)
}

func (this *QCborStreamWriter) StartMapWithCount(count uint64) {
	QCborStreamWriter_StartMapWithCount(this.h, (ulonglong)(count))
}

func (this *QCborStreamWriter) EndMap() bool {
	return (bool)(QCborStreamWriter_EndMap(this.h))
}

func (this *QCborStreamWriter) Append22(str string, size int64) {
	str_Cstring := CString(str)
	defer free(unsafe.Pointer(str_Cstring))
	QCborStreamWriter_Append22(this.h, str_Cstring, (ptrdiff_t)(size))
}
