package qt

import (
	"unsafe"
)

type QImageWriter__ImageWriterError int

const (
	QImageWriter__UnknownError           QImageWriter__ImageWriterError = 0
	QImageWriter__DeviceError            QImageWriter__ImageWriterError = 1
	QImageWriter__UnsupportedFormatError QImageWriter__ImageWriterError = 2
	QImageWriter__InvalidImageError      QImageWriter__ImageWriterError = 3
)

type QImageWriter struct {
	h          uintptr
	isSubclass bool
}

// NewQImageWriter constructs a new QImageWriter object.
func NewQImageWriter() *QImageWriter {

	ret := newQImageWriter(QImageWriter_new())
	ret.isSubclass = true
	return ret
}

// NewQImageWriter2 constructs a new QImageWriter object.
func NewQImageWriter2(device *QIODevice, format []byte) *QImageWriter {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQImageWriter(QImageWriter_new2(device.cPointer(), format_alias))
	ret.isSubclass = true
	return ret
}

// NewQImageWriter3 constructs a new QImageWriter object.
func NewQImageWriter3(fileName string) *QImageWriter {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQImageWriter(QImageWriter_new3(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQImageWriter4 constructs a new QImageWriter object.
func NewQImageWriter4(fileName string, format []byte) *QImageWriter {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQImageWriter(QImageWriter_new4(fileName_ms, format_alias))
	ret.isSubclass = true
	return ret
}

func QImageWriter_Tr(sourceText string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	var _ms struct_miqt_string = QImageWriter_Tr(sourceText_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageWriter) SetFormat(format []byte) {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	QImageWriter_SetFormat(this.h, format_alias)
}

func (this *QImageWriter) Format() []byte {
	var _bytearray struct_miqt_string = QImageWriter_Format(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QImageWriter) SetDevice(device *QIODevice) {
	QImageWriter_SetDevice(this.h, device.cPointer())
}

func (this *QImageWriter) Device() *QIODevice {
	return newQIODevice(QImageWriter_Device(this.h))
}

func (this *QImageWriter) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QImageWriter_SetFileName(this.h, fileName_ms)
}

func (this *QImageWriter) FileName() string {
	var _ms struct_miqt_string = QImageWriter_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageWriter) SetQuality(quality int) {
	QImageWriter_SetQuality(this.h, (int)(quality))
}

func (this *QImageWriter) Quality() int {
	return (int)(QImageWriter_Quality(this.h))
}

func (this *QImageWriter) SetCompression(compression int) {
	QImageWriter_SetCompression(this.h, (int)(compression))
}

func (this *QImageWriter) Compression() int {
	return (int)(QImageWriter_Compression(this.h))
}

func (this *QImageWriter) SetSubType(typeVal []byte) {
	typeVal_alias := struct_miqt_string{}
	typeVal_alias.data = (char)(unsafe.Pointer(&typeVal[0]))
	typeVal_alias.len = size_t(len(typeVal))
	QImageWriter_SetSubType(this.h, typeVal_alias)
}

func (this *QImageWriter) SubType() []byte {
	var _bytearray struct_miqt_string = QImageWriter_SubType(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QImageWriter) SupportedSubTypes() [][]byte {
	var _ma struct_miqt_array = QImageWriter_SupportedSubTypes(this.h)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QImageWriter) SetOptimizedWrite(optimize bool) {
	QImageWriter_SetOptimizedWrite(this.h, (bool)(optimize))
}

func (this *QImageWriter) OptimizedWrite() bool {
	return (bool)(QImageWriter_OptimizedWrite(this.h))
}

func (this *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	QImageWriter_SetProgressiveScanWrite(this.h, (bool)(progressive))
}

func (this *QImageWriter) ProgressiveScanWrite() bool {
	return (bool)(QImageWriter_ProgressiveScanWrite(this.h))
}

func (this *QImageWriter) Transformation() Transformation {
	return (Transformation)(QImageWriter_Transformation(this.h))
}

func (this *QImageWriter) SetTransformation(orientation Transformation) {
	QImageWriter_SetTransformation(this.h, (int)(orientation))
}

func (this *QImageWriter) SetText(key string, text string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QImageWriter_SetText(this.h, key_ms, text_ms)
}

func (this *QImageWriter) CanWrite() bool {
	return (bool)(QImageWriter_CanWrite(this.h))
}

func (this *QImageWriter) Write(image *QImage) bool {
	return (bool)(QImageWriter_Write(this.h, image.cPointer()))
}

func (this *QImageWriter) Error() ImageWriterError {
	xxxxxxxxx
}

func (this *QImageWriter) ErrorString() string {
	var _ms struct_miqt_string = QImageWriter_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageWriter) SupportsOption(option QImageIOHandler__ImageOption) bool {
	return (bool)(QImageWriter_SupportsOption(this.h, (int)(option)))
}

func QImageWriter_SupportedImageFormats() [][]byte {
	var _ma struct_miqt_array = QImageWriter_SupportedImageFormats()
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QImageWriter_SupportedMimeTypes() [][]byte {
	var _ma struct_miqt_array = QImageWriter_SupportedMimeTypes()
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QImageWriter_ImageFormatsForMimeType(mimeType []byte) [][]byte {
	mimeType_alias := struct_miqt_string{}
	mimeType_alias.data = (char)(unsafe.Pointer(&mimeType[0]))
	mimeType_alias.len = size_t(len(mimeType))
	var _ma struct_miqt_array = QImageWriter_ImageFormatsForMimeType(mimeType_alias)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QImageWriter_Tr2(sourceText string, disambiguation string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QImageWriter_Tr2(sourceText_Cstring, disambiguation_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QImageWriter_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QImageWriter_Tr3(sourceText_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
