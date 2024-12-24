package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioDecoder__Error int

const (
	QAudioDecoder__NoError           QAudioDecoder__Error = 0
	QAudioDecoder__ResourceError     QAudioDecoder__Error = 1
	QAudioDecoder__FormatError       QAudioDecoder__Error = 2
	QAudioDecoder__AccessDeniedError QAudioDecoder__Error = 3
	QAudioDecoder__NotSupportedError QAudioDecoder__Error = 4
)

type QAudioDecoder struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioDecoder constructs a new QAudioDecoder object.
func NewQAudioDecoder() *QAudioDecoder {
	ret := newQAudioDecoder(QAudioDecoder_new())
	ret.isSubclass = true
	return ret
}

// NewQAudioDecoder2 constructs a new QAudioDecoder object.
func NewQAudioDecoder2(parent *qt.QObject) *QAudioDecoder {
	ret := newQAudioDecoder(QAudioDecoder_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAudioDecoder) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioDecoder_MetaObject(this.h)))
}

func (this *QAudioDecoder) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioDecoder_Metacast(this.h, param1_Cstring))
}

func QAudioDecoder_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioDecoder_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioDecoder) IsSupported() bool {
	return (bool)(QAudioDecoder_IsSupported(this.h))
}

func (this *QAudioDecoder) IsDecoding() bool {
	return (bool)(QAudioDecoder_IsDecoding(this.h))
}

func (this *QAudioDecoder) Source() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QAudioDecoder_Source(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioDecoder) SetSource(fileName *qt.QUrl) {
	QAudioDecoder_SetSource(this.h, (*QUrl)(fileName.UnsafePointer()))
}

func (this *QAudioDecoder) SourceDevice() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QAudioDecoder_SourceDevice(this.h)))
}

func (this *QAudioDecoder) SetSourceDevice(device *qt.QIODevice) {
	QAudioDecoder_SetSourceDevice(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QAudioDecoder) AudioFormat() *QAudioFormat {
	_goptr := newQAudioFormat(QAudioDecoder_AudioFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioDecoder) SetAudioFormat(format *QAudioFormat) {
	QAudioDecoder_SetAudioFormat(this.h, format.cPointer())
}

func (this *QAudioDecoder) Error() Error {
	xxxxxxxxx
}

func (this *QAudioDecoder) ErrorString() string {
	var _ms struct_miqt_string = QAudioDecoder_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioDecoder) Read() *QAudioBuffer {
	_goptr := newQAudioBuffer(QAudioDecoder_Read(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioDecoder) BufferAvailable() bool {
	return (bool)(QAudioDecoder_BufferAvailable(this.h))
}

func (this *QAudioDecoder) Position() int64 {
	return (int64)(QAudioDecoder_Position(this.h))
}

func (this *QAudioDecoder) Duration() int64 {
	return (int64)(QAudioDecoder_Duration(this.h))
}

func (this *QAudioDecoder) Start() {
	QAudioDecoder_Start(this.h)
}

func (this *QAudioDecoder) Stop() {
	QAudioDecoder_Stop(this.h)
}

func (this *QAudioDecoder) BufferAvailableChanged(param1 bool) {
	QAudioDecoder_BufferAvailableChanged(this.h, (bool)(param1))
}

func (this *QAudioDecoder) OnBufferAvailableChanged(slot func(param1 bool)) {
	QAudioDecoder_connect_BufferAvailableChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_BufferAvailableChanged
func miqt_exec_callback_QAudioDecoder_BufferAvailableChanged(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QAudioDecoder) BufferReady() {
	QAudioDecoder_BufferReady(this.h)
}

func (this *QAudioDecoder) OnBufferReady(slot func()) {
	QAudioDecoder_connect_BufferReady(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_BufferReady
func miqt_exec_callback_QAudioDecoder_BufferReady(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioDecoder) Finished() {
	QAudioDecoder_Finished(this.h)
}

func (this *QAudioDecoder) OnFinished(slot func()) {
	QAudioDecoder_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_Finished
func miqt_exec_callback_QAudioDecoder_Finished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioDecoder) IsDecodingChanged(param1 bool) {
	QAudioDecoder_IsDecodingChanged(this.h, (bool)(param1))
}

func (this *QAudioDecoder) OnIsDecodingChanged(slot func(param1 bool)) {
	QAudioDecoder_connect_IsDecodingChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_IsDecodingChanged
func miqt_exec_callback_QAudioDecoder_IsDecodingChanged(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QAudioDecoder) FormatChanged(format *QAudioFormat) {
	QAudioDecoder_FormatChanged(this.h, format.cPointer())
}

func (this *QAudioDecoder) OnFormatChanged(slot func(format *QAudioFormat)) {
	QAudioDecoder_connect_FormatChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_FormatChanged
func miqt_exec_callback_QAudioDecoder_FormatChanged(cb intptr_t, format *QAudioFormat) {
	gofunc, ok := cgo.Handle(cb).Value().(func(format *QAudioFormat))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAudioFormat(format)

	gofunc(slotval1)
}

func (this *QAudioDecoder) ErrorWithError(error QAudioDecoder__Error) {
	QAudioDecoder_ErrorWithError(this.h, (int)(error))
}

func (this *QAudioDecoder) OnErrorWithError(slot func(error QAudioDecoder__Error)) {
	QAudioDecoder_connect_ErrorWithError(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_ErrorWithError
func miqt_exec_callback_QAudioDecoder_ErrorWithError(cb intptr_t, error int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QAudioDecoder__Error))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudioDecoder__Error)(error)

	gofunc(slotval1)
}

func (this *QAudioDecoder) SourceChanged() {
	QAudioDecoder_SourceChanged(this.h)
}

func (this *QAudioDecoder) OnSourceChanged(slot func()) {
	QAudioDecoder_connect_SourceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_SourceChanged
func miqt_exec_callback_QAudioDecoder_SourceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioDecoder) PositionChanged(position int64) {
	QAudioDecoder_PositionChanged(this.h, (longlong)(position))
}

func (this *QAudioDecoder) OnPositionChanged(slot func(position int64)) {
	QAudioDecoder_connect_PositionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_PositionChanged
func miqt_exec_callback_QAudioDecoder_PositionChanged(cb intptr_t, position longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(position int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(position)

	gofunc(slotval1)
}

func (this *QAudioDecoder) DurationChanged(duration int64) {
	QAudioDecoder_DurationChanged(this.h, (longlong)(duration))
}

func (this *QAudioDecoder) OnDurationChanged(slot func(duration int64)) {
	QAudioDecoder_connect_DurationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_DurationChanged
func miqt_exec_callback_QAudioDecoder_DurationChanged(cb intptr_t, duration longlong) {
	gofunc, ok := cgo.Handle(cb).Value().(func(duration int64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(duration)

	gofunc(slotval1)
}

func QAudioDecoder_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioDecoder_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioDecoder_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioDecoder_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioDecoder) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioDecoder_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioDecoder) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioDecoder_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_MetaObject
func miqt_exec_callback_QAudioDecoder_MetaObject(self QAudioDecoder, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioDecoder{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioDecoder) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioDecoder_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioDecoder) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioDecoder_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioDecoder_Metacast
func miqt_exec_callback_QAudioDecoder_Metacast(self QAudioDecoder, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAudioDecoder{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
