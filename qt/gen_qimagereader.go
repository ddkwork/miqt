package qt

import (
	"unsafe"
)

type QImageReader__ImageReaderError int

const (
	QImageReader__UnknownError           QImageReader__ImageReaderError = 0
	QImageReader__FileNotFoundError      QImageReader__ImageReaderError = 1
	QImageReader__DeviceError            QImageReader__ImageReaderError = 2
	QImageReader__UnsupportedFormatError QImageReader__ImageReaderError = 3
	QImageReader__InvalidDataError       QImageReader__ImageReaderError = 4
)

type QImageReader struct {
	h          uintptr
	isSubclass bool
}

// NewQImageReader constructs a new QImageReader object.
func NewQImageReader() *QImageReader {
	ret := newQImageReader(QImageReader_new())
	ret.isSubclass = true
	return ret
}

// NewQImageReader2 constructs a new QImageReader object.
func NewQImageReader2(device *QIODevice) *QImageReader {
	ret := newQImageReader(QImageReader_new2(device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQImageReader3 constructs a new QImageReader object.
func NewQImageReader3(fileName string) *QImageReader {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQImageReader(QImageReader_new3(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQImageReader4 constructs a new QImageReader object.
func NewQImageReader4(device *QIODevice, format []byte) *QImageReader {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQImageReader(QImageReader_new4(device.cPointer(), format_alias))
	ret.isSubclass = true
	return ret
}

// NewQImageReader5 constructs a new QImageReader object.
func NewQImageReader5(fileName string, format []byte) *QImageReader {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))

	ret := newQImageReader(QImageReader_new5(fileName_ms, format_alias))
	ret.isSubclass = true
	return ret
}

func QImageReader_Tr(sourceText string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	var _ms struct_miqt_string = QImageReader_Tr(sourceText_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageReader) SetFormat(format []byte) {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	QImageReader_SetFormat(this.h, format_alias)
}

func (this *QImageReader) Format() []byte {
	var _bytearray struct_miqt_string = QImageReader_Format(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	QImageReader_SetAutoDetectImageFormat(this.h, (bool)(enabled))
}

func (this *QImageReader) AutoDetectImageFormat() bool {
	return (bool)(QImageReader_AutoDetectImageFormat(this.h))
}

func (this *QImageReader) SetDecideFormatFromContent(ignored bool) {
	QImageReader_SetDecideFormatFromContent(this.h, (bool)(ignored))
}

func (this *QImageReader) DecideFormatFromContent() bool {
	return (bool)(QImageReader_DecideFormatFromContent(this.h))
}

func (this *QImageReader) SetDevice(device *QIODevice) {
	QImageReader_SetDevice(this.h, device.cPointer())
}

func (this *QImageReader) Device() *QIODevice {
	return newQIODevice(QImageReader_Device(this.h))
}

func (this *QImageReader) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QImageReader_SetFileName(this.h, fileName_ms)
}

func (this *QImageReader) FileName() string {
	var _ms struct_miqt_string = QImageReader_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageReader) Size() *QSize {
	_goptr := newQSize(QImageReader_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageReader) ImageFormat() QImage__Format {
	return (QImage__Format)(QImageReader_ImageFormat(this.h))
}

func (this *QImageReader) TextKeys() []string {
	var _ma struct_miqt_array = QImageReader_TextKeys(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QImageReader) Text(key string) string {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	var _ms struct_miqt_string = QImageReader_Text(this.h, key_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageReader) SetClipRect(rect *QRect) {
	QImageReader_SetClipRect(this.h, rect.cPointer())
}

func (this *QImageReader) ClipRect() *QRect {
	_goptr := newQRect(QImageReader_ClipRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageReader) SetScaledSize(size *QSize) {
	QImageReader_SetScaledSize(this.h, size.cPointer())
}

func (this *QImageReader) ScaledSize() *QSize {
	_goptr := newQSize(QImageReader_ScaledSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageReader) SetQuality(quality int) {
	QImageReader_SetQuality(this.h, (int)(quality))
}

func (this *QImageReader) Quality() int {
	return (int)(QImageReader_Quality(this.h))
}

func (this *QImageReader) SetScaledClipRect(rect *QRect) {
	QImageReader_SetScaledClipRect(this.h, rect.cPointer())
}

func (this *QImageReader) ScaledClipRect() *QRect {
	_goptr := newQRect(QImageReader_ScaledClipRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageReader) SetBackgroundColor(color *QColor) {
	QImageReader_SetBackgroundColor(this.h, color.cPointer())
}

func (this *QImageReader) BackgroundColor() *QColor {
	_goptr := newQColor(QImageReader_BackgroundColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageReader) SupportsAnimation() bool {
	return (bool)(QImageReader_SupportsAnimation(this.h))
}

func (this *QImageReader) Transformation() Transformation {
	return (Transformation)(QImageReader_Transformation(this.h))
}

func (this *QImageReader) SetAutoTransform(enabled bool) {
	QImageReader_SetAutoTransform(this.h, (bool)(enabled))
}

func (this *QImageReader) AutoTransform() bool {
	return (bool)(QImageReader_AutoTransform(this.h))
}

func (this *QImageReader) SubType() []byte {
	var _bytearray struct_miqt_string = QImageReader_SubType(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QImageReader) SupportedSubTypes() [][]byte {
	var _ma struct_miqt_array = QImageReader_SupportedSubTypes(this.h)
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

func (this *QImageReader) CanRead() bool {
	return (bool)(QImageReader_CanRead(this.h))
}

func (this *QImageReader) Read() *QImage {
	_goptr := newQImage(QImageReader_Read(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageReader) ReadWithImage(image *QImage) bool {
	return (bool)(QImageReader_ReadWithImage(this.h, image.cPointer()))
}

func (this *QImageReader) JumpToNextImage() bool {
	return (bool)(QImageReader_JumpToNextImage(this.h))
}

func (this *QImageReader) JumpToImage(imageNumber int) bool {
	return (bool)(QImageReader_JumpToImage(this.h, (int)(imageNumber)))
}

func (this *QImageReader) LoopCount() int {
	return (int)(QImageReader_LoopCount(this.h))
}

func (this *QImageReader) ImageCount() int {
	return (int)(QImageReader_ImageCount(this.h))
}

func (this *QImageReader) NextImageDelay() int {
	return (int)(QImageReader_NextImageDelay(this.h))
}

func (this *QImageReader) CurrentImageNumber() int {
	return (int)(QImageReader_CurrentImageNumber(this.h))
}

func (this *QImageReader) CurrentImageRect() *QRect {
	_goptr := newQRect(QImageReader_CurrentImageRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageReader) Error() ImageReaderError {
	xxxxxxxxx
}

func (this *QImageReader) ErrorString() string {
	var _ms struct_miqt_string = QImageReader_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageReader) SupportsOption(option QImageIOHandler__ImageOption) bool {
	return (bool)(QImageReader_SupportsOption(this.h, (int)(option)))
}

func QImageReader_ImageFormatWithFileName(fileName string) []byte {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _bytearray struct_miqt_string = QImageReader_ImageFormatWithFileName(fileName_ms)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QImageReader_ImageFormatWithDevice(device *QIODevice) []byte {
	var _bytearray struct_miqt_string = QImageReader_ImageFormatWithDevice(device.cPointer())
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QImageReader_SupportedImageFormats() [][]byte {
	var _ma struct_miqt_array = QImageReader_SupportedImageFormats()
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

func QImageReader_SupportedMimeTypes() [][]byte {
	var _ma struct_miqt_array = QImageReader_SupportedMimeTypes()
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

func QImageReader_ImageFormatsForMimeType(mimeType []byte) [][]byte {
	mimeType_alias := struct_miqt_string{}
	mimeType_alias.data = (char)(unsafe.Pointer(&mimeType[0]))
	mimeType_alias.len = size_t(len(mimeType))
	var _ma struct_miqt_array = QImageReader_ImageFormatsForMimeType(mimeType_alias)
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

func QImageReader_AllocationLimit() int {
	return (int)(QImageReader_AllocationLimit())
}

func QImageReader_SetAllocationLimit(mbLimit int) {
	QImageReader_SetAllocationLimit((int)(mbLimit))
}

func QImageReader_Tr2(sourceText string, disambiguation string) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QImageReader_Tr2(sourceText_Cstring, disambiguation_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QImageReader_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QImageReader_Tr3(sourceText_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
