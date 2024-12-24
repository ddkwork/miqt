package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QMediaCaptureSession struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaCaptureSession constructs a new QMediaCaptureSession object.
func NewQMediaCaptureSession() *QMediaCaptureSession {
	ret := newQMediaCaptureSession(QMediaCaptureSession_new())
	ret.isSubclass = true
	return ret
}

// NewQMediaCaptureSession2 constructs a new QMediaCaptureSession object.
func NewQMediaCaptureSession2(parent *qt.QObject) *QMediaCaptureSession {
	ret := newQMediaCaptureSession(QMediaCaptureSession_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QMediaCaptureSession) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaCaptureSession_MetaObject(this.h)))
}

func (this *QMediaCaptureSession) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMediaCaptureSession_Metacast(this.h, param1_Cstring))
}

func QMediaCaptureSession_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMediaCaptureSession_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaCaptureSession) AudioInput() *QAudioInput {
	return newQAudioInput(QMediaCaptureSession_AudioInput(this.h))
}

func (this *QMediaCaptureSession) SetAudioInput(input *QAudioInput) {
	QMediaCaptureSession_SetAudioInput(this.h, input.cPointer())
}

func (this *QMediaCaptureSession) AudioBufferInput() *QAudioBufferInput {
	return newQAudioBufferInput(QMediaCaptureSession_AudioBufferInput(this.h))
}

func (this *QMediaCaptureSession) SetAudioBufferInput(input *QAudioBufferInput) {
	QMediaCaptureSession_SetAudioBufferInput(this.h, input.cPointer())
}

func (this *QMediaCaptureSession) Camera() *QCamera {
	return newQCamera(QMediaCaptureSession_Camera(this.h))
}

func (this *QMediaCaptureSession) SetCamera(camera *QCamera) {
	QMediaCaptureSession_SetCamera(this.h, camera.cPointer())
}

func (this *QMediaCaptureSession) ImageCapture() *QImageCapture {
	return newQImageCapture(QMediaCaptureSession_ImageCapture(this.h))
}

func (this *QMediaCaptureSession) SetImageCapture(imageCapture *QImageCapture) {
	QMediaCaptureSession_SetImageCapture(this.h, imageCapture.cPointer())
}

func (this *QMediaCaptureSession) ScreenCapture() *QScreenCapture {
	return newQScreenCapture(QMediaCaptureSession_ScreenCapture(this.h))
}

func (this *QMediaCaptureSession) SetScreenCapture(screenCapture *QScreenCapture) {
	QMediaCaptureSession_SetScreenCapture(this.h, screenCapture.cPointer())
}

func (this *QMediaCaptureSession) WindowCapture() *QWindowCapture {
	return newQWindowCapture(QMediaCaptureSession_WindowCapture(this.h))
}

func (this *QMediaCaptureSession) SetWindowCapture(windowCapture *QWindowCapture) {
	QMediaCaptureSession_SetWindowCapture(this.h, windowCapture.cPointer())
}

func (this *QMediaCaptureSession) VideoFrameInput() *QVideoFrameInput {
	return newQVideoFrameInput(QMediaCaptureSession_VideoFrameInput(this.h))
}

func (this *QMediaCaptureSession) SetVideoFrameInput(input *QVideoFrameInput) {
	QMediaCaptureSession_SetVideoFrameInput(this.h, input.cPointer())
}

func (this *QMediaCaptureSession) Recorder() *QMediaRecorder {
	return newQMediaRecorder(QMediaCaptureSession_Recorder(this.h))
}

func (this *QMediaCaptureSession) SetRecorder(recorder *QMediaRecorder) {
	QMediaCaptureSession_SetRecorder(this.h, recorder.cPointer())
}

func (this *QMediaCaptureSession) SetVideoOutput(output *qt.QObject) {
	QMediaCaptureSession_SetVideoOutput(this.h, (*QObject)(output.UnsafePointer()))
}

func (this *QMediaCaptureSession) VideoOutput() *qt.QObject {
	return qt.UnsafeNewQObject(unsafe.Pointer(QMediaCaptureSession_VideoOutput(this.h)))
}

func (this *QMediaCaptureSession) SetVideoSink(sink *QVideoSink) {
	QMediaCaptureSession_SetVideoSink(this.h, sink.cPointer())
}

func (this *QMediaCaptureSession) VideoSink() *QVideoSink {
	return newQVideoSink(QMediaCaptureSession_VideoSink(this.h))
}

func (this *QMediaCaptureSession) SetAudioOutput(output *QAudioOutput) {
	QMediaCaptureSession_SetAudioOutput(this.h, output.cPointer())
}

func (this *QMediaCaptureSession) AudioOutput() *QAudioOutput {
	return newQAudioOutput(QMediaCaptureSession_AudioOutput(this.h))
}

func (this *QMediaCaptureSession) AudioInputChanged() {
	QMediaCaptureSession_AudioInputChanged(this.h)
}

func (this *QMediaCaptureSession) OnAudioInputChanged(slot func()) {
	QMediaCaptureSession_connect_AudioInputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_AudioInputChanged
func miqt_exec_callback_QMediaCaptureSession_AudioInputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) AudioBufferInputChanged() {
	QMediaCaptureSession_AudioBufferInputChanged(this.h)
}

func (this *QMediaCaptureSession) OnAudioBufferInputChanged(slot func()) {
	QMediaCaptureSession_connect_AudioBufferInputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_AudioBufferInputChanged
func miqt_exec_callback_QMediaCaptureSession_AudioBufferInputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) CameraChanged() {
	QMediaCaptureSession_CameraChanged(this.h)
}

func (this *QMediaCaptureSession) OnCameraChanged(slot func()) {
	QMediaCaptureSession_connect_CameraChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_CameraChanged
func miqt_exec_callback_QMediaCaptureSession_CameraChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) ScreenCaptureChanged() {
	QMediaCaptureSession_ScreenCaptureChanged(this.h)
}

func (this *QMediaCaptureSession) OnScreenCaptureChanged(slot func()) {
	QMediaCaptureSession_connect_ScreenCaptureChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_ScreenCaptureChanged
func miqt_exec_callback_QMediaCaptureSession_ScreenCaptureChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) WindowCaptureChanged() {
	QMediaCaptureSession_WindowCaptureChanged(this.h)
}

func (this *QMediaCaptureSession) OnWindowCaptureChanged(slot func()) {
	QMediaCaptureSession_connect_WindowCaptureChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_WindowCaptureChanged
func miqt_exec_callback_QMediaCaptureSession_WindowCaptureChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) VideoFrameInputChanged() {
	QMediaCaptureSession_VideoFrameInputChanged(this.h)
}

func (this *QMediaCaptureSession) OnVideoFrameInputChanged(slot func()) {
	QMediaCaptureSession_connect_VideoFrameInputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_VideoFrameInputChanged
func miqt_exec_callback_QMediaCaptureSession_VideoFrameInputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) ImageCaptureChanged() {
	QMediaCaptureSession_ImageCaptureChanged(this.h)
}

func (this *QMediaCaptureSession) OnImageCaptureChanged(slot func()) {
	QMediaCaptureSession_connect_ImageCaptureChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_ImageCaptureChanged
func miqt_exec_callback_QMediaCaptureSession_ImageCaptureChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) RecorderChanged() {
	QMediaCaptureSession_RecorderChanged(this.h)
}

func (this *QMediaCaptureSession) OnRecorderChanged(slot func()) {
	QMediaCaptureSession_connect_RecorderChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_RecorderChanged
func miqt_exec_callback_QMediaCaptureSession_RecorderChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) VideoOutputChanged() {
	QMediaCaptureSession_VideoOutputChanged(this.h)
}

func (this *QMediaCaptureSession) OnVideoOutputChanged(slot func()) {
	QMediaCaptureSession_connect_VideoOutputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_VideoOutputChanged
func miqt_exec_callback_QMediaCaptureSession_VideoOutputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaCaptureSession) AudioOutputChanged() {
	QMediaCaptureSession_AudioOutputChanged(this.h)
}

func (this *QMediaCaptureSession) OnAudioOutputChanged(slot func()) {
	QMediaCaptureSession_connect_AudioOutputChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_AudioOutputChanged
func miqt_exec_callback_QMediaCaptureSession_AudioOutputChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QMediaCaptureSession_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaCaptureSession_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaCaptureSession_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaCaptureSession_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaCaptureSession) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaCaptureSession_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QMediaCaptureSession) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaCaptureSession_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_MetaObject
func miqt_exec_callback_QMediaCaptureSession_MetaObject(self QMediaCaptureSession, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMediaCaptureSession{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QMediaCaptureSession) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMediaCaptureSession_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMediaCaptureSession) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaCaptureSession_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaCaptureSession_Metacast
func miqt_exec_callback_QMediaCaptureSession_Metacast(self QMediaCaptureSession, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QMediaCaptureSession{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
