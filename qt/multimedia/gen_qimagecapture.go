package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QImageCapture__Error int

const (
	QImageCapture__NoError                  QImageCapture__Error = 0
	QImageCapture__NotReadyError            QImageCapture__Error = 1
	QImageCapture__ResourceError            QImageCapture__Error = 2
	QImageCapture__OutOfSpaceError          QImageCapture__Error = 3
	QImageCapture__NotSupportedFeatureError QImageCapture__Error = 4
	QImageCapture__FormatError              QImageCapture__Error = 5
)

type QImageCapture__Quality int

const (
	QImageCapture__VeryLowQuality  QImageCapture__Quality = 0
	QImageCapture__LowQuality      QImageCapture__Quality = 1
	QImageCapture__NormalQuality   QImageCapture__Quality = 2
	QImageCapture__HighQuality     QImageCapture__Quality = 3
	QImageCapture__VeryHighQuality QImageCapture__Quality = 4
)

type QImageCapture__FileFormat int

const (
	QImageCapture__UnspecifiedFormat QImageCapture__FileFormat = 0
	QImageCapture__JPEG              QImageCapture__FileFormat = 1
	QImageCapture__PNG               QImageCapture__FileFormat = 2
	QImageCapture__WebP              QImageCapture__FileFormat = 3
	QImageCapture__Tiff              QImageCapture__FileFormat = 4
	QImageCapture__LastFileFormat    QImageCapture__FileFormat = 4
)

type QImageCapture struct {
	h          uintptr
	isSubclass bool
}

// NewQImageCapture constructs a new QImageCapture object.
func NewQImageCapture() *QImageCapture {
	g := newQImageCapture(QImageCapture_new())
	g.isSubclass = true
	return g
}

// NewQImageCapture2 constructs a new QImageCapture object.
func NewQImageCapture2(parent *qt.QObject) *QImageCapture {
	g := newQImageCapture(QImageCapture_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QImageCapture) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QImageCapture_MetaObject(this.h)))
}

func (this *QImageCapture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QImageCapture_Metacast(this.h, param1_Cstring))
}

func QImageCapture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QImageCapture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) IsAvailable() bool {
	return (bool)(QImageCapture_IsAvailable(this.h))
}

func (this *QImageCapture) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(QImageCapture_CaptureSession(this.h))
}

func (this *QImageCapture) Error() Error {
	xxxxxxxxx
}

func (this *QImageCapture) ErrorString() string {
	var _ms struct_miqt_string = QImageCapture_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) IsReadyForCapture() bool {
	return (bool)(QImageCapture_IsReadyForCapture(this.h))
}

func (this *QImageCapture) FileFormat() FileFormat {
	xxxxxxxxx
}

func (this *QImageCapture) SetFileFormat(format FileFormat) {
	QImageCapture_SetFileFormat(this.h, format)
}

func QImageCapture_SupportedFormats() []FileFormat {
	var _ma struct_miqt_array = QImageCapture_SupportedFormats()
	_ret := make([]FileFormat, int(_ma.len))
	_outCast := (*[0xffff]FileFormat)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func QImageCapture_FileFormatName(c FileFormat) string {
	var _ms struct_miqt_string = QImageCapture_FileFormatName(c)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QImageCapture_FileFormatDescription(c FileFormat) string {
	var _ms struct_miqt_string = QImageCapture_FileFormatDescription(c)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) Resolution() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QImageCapture_Resolution(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageCapture) SetResolution(resolution *qt.QSize) {
	QImageCapture_SetResolution(this.h, (*QSize)(resolution.UnsafePointer()))
}

func (this *QImageCapture) SetResolution2(width int, height int) {
	QImageCapture_SetResolution2(this.h, (int)(width), (int)(height))
}

func (this *QImageCapture) Quality() Quality {
	xxxxxxxxx
}

func (this *QImageCapture) SetQuality(quality Quality) {
	QImageCapture_SetQuality(this.h, quality)
}

func (this *QImageCapture) MetaData() *QMediaMetaData {
	_goptr := newQMediaMetaData(QImageCapture_MetaData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageCapture) SetMetaData(metaData *QMediaMetaData) {
	QImageCapture_SetMetaData(this.h, metaData.cPointer())
}

func (this *QImageCapture) AddMetaData(metaData *QMediaMetaData) {
	QImageCapture_AddMetaData(this.h, metaData.cPointer())
}

func (this *QImageCapture) CaptureToFile() int {
	return (int)(QImageCapture_CaptureToFile(this.h))
}

func (this *QImageCapture) Capture() int {
	return (int)(QImageCapture_Capture(this.h))
}

func (this *QImageCapture) ErrorChanged() {
	QImageCapture_ErrorChanged(this.h)
}

func (this *QImageCapture) OnErrorChanged(slot func()) {
	QImageCapture_connect_ErrorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ErrorChanged
func miqt_exec_callback_QImageCapture_ErrorChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) ErrorOccurred(id int, error QImageCapture__Error, errorString string) {
	errorString_ms := struct_miqt_string{}
	errorString_ms.data = CString(errorString)
	errorString_ms.len = size_t(len(errorString))
	defer free(unsafe.Pointer(errorString_ms.data))
	QImageCapture_ErrorOccurred(this.h, (int)(id), (int)(error), errorString_ms)
}

func (this *QImageCapture) OnErrorOccurred(slot func(id int, error QImageCapture__Error, errorString string)) {
	QImageCapture_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ErrorOccurred
func miqt_exec_callback_QImageCapture_ErrorOccurred(cb intptr_t, id int, error int, errorString struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, error QImageCapture__Error, errorString string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := (QImageCapture__Error)(error)

	var errorString_ms struct_miqt_string = errorString
	errorString_ret := GoStringN(errorString_ms.data, int(int64(errorString_ms.len)))
	free(unsafe.Pointer(errorString_ms.data))
	slotval3 := errorString_ret

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QImageCapture) ReadyForCaptureChanged(ready bool) {
	QImageCapture_ReadyForCaptureChanged(this.h, (bool)(ready))
}

func (this *QImageCapture) OnReadyForCaptureChanged(slot func(ready bool)) {
	QImageCapture_connect_ReadyForCaptureChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ReadyForCaptureChanged
func miqt_exec_callback_QImageCapture_ReadyForCaptureChanged(cb intptr_t, ready bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(ready bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(ready)

	gofunc(slotval1)
}

func (this *QImageCapture) MetaDataChanged() {
	QImageCapture_MetaDataChanged(this.h)
}

func (this *QImageCapture) OnMetaDataChanged(slot func()) {
	QImageCapture_connect_MetaDataChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_MetaDataChanged
func miqt_exec_callback_QImageCapture_MetaDataChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) FileFormatChanged() {
	QImageCapture_FileFormatChanged(this.h)
}

func (this *QImageCapture) OnFileFormatChanged(slot func()) {
	QImageCapture_connect_FileFormatChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_FileFormatChanged
func miqt_exec_callback_QImageCapture_FileFormatChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) QualityChanged() {
	QImageCapture_QualityChanged(this.h)
}

func (this *QImageCapture) OnQualityChanged(slot func()) {
	QImageCapture_connect_QualityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_QualityChanged
func miqt_exec_callback_QImageCapture_QualityChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) ResolutionChanged() {
	QImageCapture_ResolutionChanged(this.h)
}

func (this *QImageCapture) OnResolutionChanged(slot func()) {
	QImageCapture_connect_ResolutionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ResolutionChanged
func miqt_exec_callback_QImageCapture_ResolutionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) ImageExposed(id int) {
	QImageCapture_ImageExposed(this.h, (int)(id))
}

func (this *QImageCapture) OnImageExposed(slot func(id int)) {
	QImageCapture_connect_ImageExposed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ImageExposed
func miqt_exec_callback_QImageCapture_ImageExposed(cb intptr_t, id int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	gofunc(slotval1)
}

func (this *QImageCapture) ImageCaptured(id int, preview *qt.QImage) {
	QImageCapture_ImageCaptured(this.h, (int)(id), (*QImage)(preview.UnsafePointer()))
}

func (this *QImageCapture) OnImageCaptured(slot func(id int, preview *qt.QImage)) {
	QImageCapture_connect_ImageCaptured(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ImageCaptured
func miqt_exec_callback_QImageCapture_ImageCaptured(cb intptr_t, id int, preview *QImage) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, preview *qt.QImage))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := qt.UnsafeNewQImage(unsafe.Pointer(preview))

	gofunc(slotval1, slotval2)
}

func (this *QImageCapture) ImageMetadataAvailable(id int, metaData *QMediaMetaData) {
	QImageCapture_ImageMetadataAvailable(this.h, (int)(id), metaData.cPointer())
}

func (this *QImageCapture) OnImageMetadataAvailable(slot func(id int, metaData *QMediaMetaData)) {
	QImageCapture_connect_ImageMetadataAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ImageMetadataAvailable
func miqt_exec_callback_QImageCapture_ImageMetadataAvailable(cb intptr_t, id int, metaData *QMediaMetaData) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, metaData *QMediaMetaData))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := newQMediaMetaData(metaData)

	gofunc(slotval1, slotval2)
}

func (this *QImageCapture) ImageAvailable(id int, frame *QVideoFrame) {
	QImageCapture_ImageAvailable(this.h, (int)(id), frame.cPointer())
}

func (this *QImageCapture) OnImageAvailable(slot func(id int, frame *QVideoFrame)) {
	QImageCapture_connect_ImageAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ImageAvailable
func miqt_exec_callback_QImageCapture_ImageAvailable(cb intptr_t, id int, frame *QVideoFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, frame *QVideoFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := newQVideoFrame(frame)

	gofunc(slotval1, slotval2)
}

func (this *QImageCapture) ImageSaved(id int, fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QImageCapture_ImageSaved(this.h, (int)(id), fileName_ms)
}

func (this *QImageCapture) OnImageSaved(slot func(id int, fileName string)) {
	QImageCapture_connect_ImageSaved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_ImageSaved
func miqt_exec_callback_QImageCapture_ImageSaved(cb intptr_t, id int, fileName struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, fileName string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	var fileName_ms struct_miqt_string = fileName
	fileName_ret := GoStringN(fileName_ms.data, int(int64(fileName_ms.len)))
	free(unsafe.Pointer(fileName_ms.data))
	slotval2 := fileName_ret

	gofunc(slotval1, slotval2)
}

func QImageCapture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QImageCapture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QImageCapture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QImageCapture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) CaptureToFile1(location string) int {
	location_ms := struct_miqt_string{}
	location_ms.data = CString(location)
	location_ms.len = size_t(len(location))
	defer free(unsafe.Pointer(location_ms.data))
	return (int)(QImageCapture_CaptureToFile1(this.h, location_ms))
}

func (this *QImageCapture) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QImageCapture_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QImageCapture) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageCapture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_MetaObject
func miqt_exec_callback_QImageCapture_MetaObject(self QImageCapture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageCapture{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QImageCapture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QImageCapture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QImageCapture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageCapture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_Metacast
func miqt_exec_callback_QImageCapture_Metacast(self QImageCapture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QImageCapture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
