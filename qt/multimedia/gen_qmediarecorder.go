package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QMediaRecorder__Quality int

const (
	QMediaRecorder__VeryLowQuality  QMediaRecorder__Quality = 0
	QMediaRecorder__LowQuality      QMediaRecorder__Quality = 1
	QMediaRecorder__NormalQuality   QMediaRecorder__Quality = 2
	QMediaRecorder__HighQuality     QMediaRecorder__Quality = 3
	QMediaRecorder__VeryHighQuality QMediaRecorder__Quality = 4
)

type QMediaRecorder__EncodingMode int

const (
	QMediaRecorder__ConstantQualityEncoding QMediaRecorder__EncodingMode = 0
	QMediaRecorder__ConstantBitRateEncoding QMediaRecorder__EncodingMode = 1
	QMediaRecorder__AverageBitRateEncoding  QMediaRecorder__EncodingMode = 2
	QMediaRecorder__TwoPassEncoding         QMediaRecorder__EncodingMode = 3
)

type QMediaRecorder__RecorderState int

const (
	QMediaRecorder__StoppedState   QMediaRecorder__RecorderState = 0
	QMediaRecorder__RecordingState QMediaRecorder__RecorderState = 1
	QMediaRecorder__PausedState    QMediaRecorder__RecorderState = 2
)

type QMediaRecorder__Error int

const (
	QMediaRecorder__NoError             QMediaRecorder__Error = 0
	QMediaRecorder__ResourceError       QMediaRecorder__Error = 1
	QMediaRecorder__FormatError         QMediaRecorder__Error = 2
	QMediaRecorder__OutOfSpaceError     QMediaRecorder__Error = 3
	QMediaRecorder__LocationNotWritable QMediaRecorder__Error = 4
)

type QMediaRecorder struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaRecorder constructs a new QMediaRecorder object.
func NewQMediaRecorder() *QMediaRecorder {
	ret := newQMediaRecorder(QMediaRecorder_new())
	ret.isSubclass = true
	return ret
}

// NewQMediaRecorder2 constructs a new QMediaRecorder object.
func NewQMediaRecorder2(parent *qt.QObject) *QMediaRecorder {
	ret := newQMediaRecorder(QMediaRecorder_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QMediaRecorder) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaRecorder_MetaObject(this.h)))
}

func (this *QMediaRecorder) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMediaRecorder_Metacast(this.h, param1_Cstring))
}

func QMediaRecorder_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMediaRecorder_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaRecorder) IsAvailable() bool {
	return (bool)(QMediaRecorder_IsAvailable(this.h))
}

func (this *QMediaRecorder) OutputLocation() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QMediaRecorder_OutputLocation(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaRecorder) SetOutputLocation(location *qt.QUrl) {
	QMediaRecorder_SetOutputLocation(this.h, (*QUrl)(location.UnsafePointer()))
}

func (this *QMediaRecorder) SetOutputDevice(device *qt.QIODevice) {
	QMediaRecorder_SetOutputDevice(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QMediaRecorder) OutputDevice() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QMediaRecorder_OutputDevice(this.h)))
}

func (this *QMediaRecorder) ActualLocation() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QMediaRecorder_ActualLocation(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaRecorder) RecorderState() RecorderState {
	xxxxxxxxx
}

func (this *QMediaRecorder) Error() Error {
	xxxxxxxxx
}

func (this *QMediaRecorder) ErrorString() string {
	var _ms struct_miqt_string = QMediaRecorder_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaRecorder) Duration() int64 {
	return (int64)(QMediaRecorder_Duration(this.h))
}

func (this *QMediaRecorder) MediaFormat() *QMediaFormat {
	_goptr := newQMediaFormat(QMediaRecorder_MediaFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaRecorder) SetMediaFormat(format *QMediaFormat) {
	QMediaRecorder_SetMediaFormat(this.h, format.cPointer())
}

func (this *QMediaRecorder) EncodingMode() EncodingMode {
	xxxxxxxxx
}

func (this *QMediaRecorder) SetEncodingMode(encodingMode EncodingMode) {
	QMediaRecorder_SetEncodingMode(this.h, encodingMode)
}

func (this *QMediaRecorder) Quality() Quality {
	xxxxxxxxx
}

func (this *QMediaRecorder) SetQuality(quality Quality) {
	QMediaRecorder_SetQuality(this.h, quality)
}

func (this *QMediaRecorder) VideoResolution() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QMediaRecorder_VideoResolution(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaRecorder) SetVideoResolution(videoResolution *qt.QSize) {
	QMediaRecorder_SetVideoResolution(this.h, (*QSize)(videoResolution.UnsafePointer()))
}

func (this *QMediaRecorder) SetVideoResolution2(width int, height int) {
	QMediaRecorder_SetVideoResolution2(this.h, (int)(width), (int)(height))
}

func (this *QMediaRecorder) VideoFrameRate() float64 {
	return (float64)(QMediaRecorder_VideoFrameRate(this.h))
}

func (this *QMediaRecorder) SetVideoFrameRate(frameRate float64) {
	QMediaRecorder_SetVideoFrameRate(this.h, (double)(frameRate))
}

func (this *QMediaRecorder) VideoBitRate() int {
	return (int)(QMediaRecorder_VideoBitRate(this.h))
}

func (this *QMediaRecorder) SetVideoBitRate(bitRate int) {
	QMediaRecorder_SetVideoBitRate(this.h, (int)(bitRate))
}

func (this *QMediaRecorder) AudioBitRate() int {
	return (int)(QMediaRecorder_AudioBitRate(this.h))
}

func (this *QMediaRecorder) SetAudioBitRate(bitRate int) {
	QMediaRecorder_SetAudioBitRate(this.h, (int)(bitRate))
}

func (this *QMediaRecorder) AudioChannelCount() int {
	return (int)(QMediaRecorder_AudioChannelCount(this.h))
}

func (this *QMediaRecorder) SetAudioChannelCount(channels int) {
	QMediaRecorder_SetAudioChannelCount(this.h, (int)(channels))
}

func (this *QMediaRecorder) AudioSampleRate() int {
	return (int)(QMediaRecorder_AudioSampleRate(this.h))
}

func (this *QMediaRecorder) SetAudioSampleRate(sampleRate int) {
	QMediaRecorder_SetAudioSampleRate(this.h, (int)(sampleRate))
}

func (this *QMediaRecorder) MetaData() *QMediaMetaData {
	_goptr := newQMediaMetaData(QMediaRecorder_MetaData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaRecorder) SetMetaData(metaData *QMediaMetaData) {
	QMediaRecorder_SetMetaData(this.h, metaData.cPointer())
}

func (this *QMediaRecorder) AddMetaData(metaData *QMediaMetaData) {
	QMediaRecorder_AddMetaData(this.h, metaData.cPointer())
}

func (this *QMediaRecorder) AutoStop() bool {
	return (bool)(QMediaRecorder_AutoStop(this.h))
}

func (this *QMediaRecorder) SetAutoStop(autoStop bool) {
	QMediaRecorder_SetAutoStop(this.h, (bool)(autoStop))
}

func (this *QMediaRecorder) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(QMediaRecorder_CaptureSession(this.h))
}

func (this *QMediaRecorder) Record() {
	QMediaRecorder_Record(this.h)
}

func (this *QMediaRecorder) Pause() {
	QMediaRecorder_Pause(this.h)
}

func (this *QMediaRecorder) Stop() {
	QMediaRecorder_Stop(this.h)
}

func (this *QMediaRecorder) RecorderStateChanged(state RecorderState) {
	QMediaRecorder_RecorderStateChanged(this.h, state)
}

func (this *QMediaRecorder) OnRecorderStateChanged(slot func(state RecorderState)) {
	QMediaRecorder_connect_RecorderStateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_RecorderStateChanged
func miqt_exec_callback_QMediaRecorder_RecorderStateChanged(cb intptr_t, state RecorderState) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state RecorderState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc(slotval1)
}

func (this *QMediaRecorder) DurationChanged(duration int64) {
	QMediaRecorder_DurationChanged(this.h, (longlong)(duration))
}

func (this *QMediaRecorder) OnDurationChanged(slot func(duration int64)) {
	QMediaRecorder_connect_DurationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_DurationChanged
func miqt_exec_callback_QMediaRecorder_DurationChanged(cb intptr_t, duration longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(duration int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(duration)

	gofunc(slotval1)
}

func (this *QMediaRecorder) ActualLocationChanged(location *qt.QUrl) {
	QMediaRecorder_ActualLocationChanged(this.h, (*QUrl)(location.UnsafePointer()))
}

func (this *QMediaRecorder) OnActualLocationChanged(slot func(location *qt.QUrl)) {
	QMediaRecorder_connect_ActualLocationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_ActualLocationChanged
func miqt_exec_callback_QMediaRecorder_ActualLocationChanged(cb intptr_t, location *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(location *qt.QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(location))

	gofunc(slotval1)
}

func (this *QMediaRecorder) EncoderSettingsChanged() {
	QMediaRecorder_EncoderSettingsChanged(this.h)
}

func (this *QMediaRecorder) OnEncoderSettingsChanged(slot func()) {
	QMediaRecorder_connect_EncoderSettingsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_EncoderSettingsChanged
func miqt_exec_callback_QMediaRecorder_EncoderSettingsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) ErrorOccurred(error Error, errorString string) {
	errorString_ms := struct_miqt_string{}
	errorString_ms.data = CString(errorString)
	errorString_ms.len = size_t(len(errorString))
	defer free(unsafe.Pointer(errorString_ms.data))
	QMediaRecorder_ErrorOccurred(this.h, error, errorString_ms)
}

func (this *QMediaRecorder) OnErrorOccurred(slot func(error Error, errorString string)) {
	QMediaRecorder_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_ErrorOccurred
func miqt_exec_callback_QMediaRecorder_ErrorOccurred(cb intptr_t, error Error, errorString struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error Error, errorString string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	var errorString_ms struct_miqt_string = errorString
	errorString_ret := GoStringN(errorString_ms.data, int(int64(errorString_ms.len)))
	free(unsafe.Pointer(errorString_ms.data))
	slotval2 := errorString_ret

	gofunc(slotval1, slotval2)
}

func (this *QMediaRecorder) ErrorChanged() {
	QMediaRecorder_ErrorChanged(this.h)
}

func (this *QMediaRecorder) OnErrorChanged(slot func()) {
	QMediaRecorder_connect_ErrorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_ErrorChanged
func miqt_exec_callback_QMediaRecorder_ErrorChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) MetaDataChanged() {
	QMediaRecorder_MetaDataChanged(this.h)
}

func (this *QMediaRecorder) OnMetaDataChanged(slot func()) {
	QMediaRecorder_connect_MetaDataChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_MetaDataChanged
func miqt_exec_callback_QMediaRecorder_MetaDataChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) MediaFormatChanged() {
	QMediaRecorder_MediaFormatChanged(this.h)
}

func (this *QMediaRecorder) OnMediaFormatChanged(slot func()) {
	QMediaRecorder_connect_MediaFormatChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_MediaFormatChanged
func miqt_exec_callback_QMediaRecorder_MediaFormatChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) EncodingModeChanged() {
	QMediaRecorder_EncodingModeChanged(this.h)
}

func (this *QMediaRecorder) OnEncodingModeChanged(slot func()) {
	QMediaRecorder_connect_EncodingModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_EncodingModeChanged
func miqt_exec_callback_QMediaRecorder_EncodingModeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) QualityChanged() {
	QMediaRecorder_QualityChanged(this.h)
}

func (this *QMediaRecorder) OnQualityChanged(slot func()) {
	QMediaRecorder_connect_QualityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_QualityChanged
func miqt_exec_callback_QMediaRecorder_QualityChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) VideoResolutionChanged() {
	QMediaRecorder_VideoResolutionChanged(this.h)
}

func (this *QMediaRecorder) OnVideoResolutionChanged(slot func()) {
	QMediaRecorder_connect_VideoResolutionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_VideoResolutionChanged
func miqt_exec_callback_QMediaRecorder_VideoResolutionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) VideoFrameRateChanged() {
	QMediaRecorder_VideoFrameRateChanged(this.h)
}

func (this *QMediaRecorder) OnVideoFrameRateChanged(slot func()) {
	QMediaRecorder_connect_VideoFrameRateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_VideoFrameRateChanged
func miqt_exec_callback_QMediaRecorder_VideoFrameRateChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) VideoBitRateChanged() {
	QMediaRecorder_VideoBitRateChanged(this.h)
}

func (this *QMediaRecorder) OnVideoBitRateChanged(slot func()) {
	QMediaRecorder_connect_VideoBitRateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_VideoBitRateChanged
func miqt_exec_callback_QMediaRecorder_VideoBitRateChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) AudioBitRateChanged() {
	QMediaRecorder_AudioBitRateChanged(this.h)
}

func (this *QMediaRecorder) OnAudioBitRateChanged(slot func()) {
	QMediaRecorder_connect_AudioBitRateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_AudioBitRateChanged
func miqt_exec_callback_QMediaRecorder_AudioBitRateChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) AudioChannelCountChanged() {
	QMediaRecorder_AudioChannelCountChanged(this.h)
}

func (this *QMediaRecorder) OnAudioChannelCountChanged(slot func()) {
	QMediaRecorder_connect_AudioChannelCountChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_AudioChannelCountChanged
func miqt_exec_callback_QMediaRecorder_AudioChannelCountChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) AudioSampleRateChanged() {
	QMediaRecorder_AudioSampleRateChanged(this.h)
}

func (this *QMediaRecorder) OnAudioSampleRateChanged(slot func()) {
	QMediaRecorder_connect_AudioSampleRateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_AudioSampleRateChanged
func miqt_exec_callback_QMediaRecorder_AudioSampleRateChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaRecorder) AutoStopChanged() {
	QMediaRecorder_AutoStopChanged(this.h)
}

func (this *QMediaRecorder) OnAutoStopChanged(slot func()) {
	QMediaRecorder_connect_AutoStopChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_AutoStopChanged
func miqt_exec_callback_QMediaRecorder_AutoStopChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QMediaRecorder_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaRecorder_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaRecorder_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaRecorder_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaRecorder) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaRecorder_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QMediaRecorder) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaRecorder_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_MetaObject
func miqt_exec_callback_QMediaRecorder_MetaObject(self QMediaRecorder, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMediaRecorder{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QMediaRecorder) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMediaRecorder_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMediaRecorder) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaRecorder_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaRecorder_Metacast
func miqt_exec_callback_QMediaRecorder_Metacast(self QMediaRecorder, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QMediaRecorder{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
