package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioSink struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioSink constructs a new QAudioSink object.
func NewQAudioSink() *QAudioSink {
	g := newQAudioSink(QAudioSink_new())
	g.isSubclass = true
	return g
}

// NewQAudioSink2 constructs a new QAudioSink object.
func NewQAudioSink2(audioDeviceInfo *QAudioDevice) *QAudioSink {
	g := newQAudioSink(QAudioSink_new2(audioDeviceInfo.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAudioSink3 constructs a new QAudioSink object.
func NewQAudioSink3(format *QAudioFormat) *QAudioSink {
	g := newQAudioSink(QAudioSink_new3(format.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAudioSink4 constructs a new QAudioSink object.
func NewQAudioSink4(format *QAudioFormat, parent *qt.QObject) *QAudioSink {
	g := newQAudioSink(QAudioSink_new4(format.cPointer(), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQAudioSink5 constructs a new QAudioSink object.
func NewQAudioSink5(audioDeviceInfo *QAudioDevice, format *QAudioFormat) *QAudioSink {
	g := newQAudioSink(QAudioSink_new5(audioDeviceInfo.cPointer(), format.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAudioSink6 constructs a new QAudioSink object.
func NewQAudioSink6(audioDeviceInfo *QAudioDevice, format *QAudioFormat, parent *qt.QObject) *QAudioSink {
	g := newQAudioSink(QAudioSink_new6(audioDeviceInfo.cPointer(), format.cPointer(), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QAudioSink) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioSink_MetaObject(this.h)))
}

func (this *QAudioSink) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioSink_Metacast(this.h, param1_Cstring))
}

func QAudioSink_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioSink_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioSink) IsNull() bool {
	return (bool)(QAudioSink_IsNull(this.h))
}

func (this *QAudioSink) Format() *QAudioFormat {
	_goptr := newQAudioFormat(QAudioSink_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioSink) Start(device *qt.QIODevice) {
	QAudioSink_Start(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QAudioSink) Start2() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QAudioSink_Start2(this.h)))
}

func (this *QAudioSink) Stop() {
	QAudioSink_Stop(this.h)
}

func (this *QAudioSink) Reset() {
	QAudioSink_Reset(this.h)
}

func (this *QAudioSink) Suspend() {
	QAudioSink_Suspend(this.h)
}

func (this *QAudioSink) Resume() {
	QAudioSink_Resume(this.h)
}

func (this *QAudioSink) SetBufferSize(bytes int64) {
	QAudioSink_SetBufferSize(this.h, (ptrdiff_t)(bytes))
}

func (this *QAudioSink) BufferSize() int64 {
	return (int64)(QAudioSink_BufferSize(this.h))
}

func (this *QAudioSink) BytesFree() int64 {
	return (int64)(QAudioSink_BytesFree(this.h))
}

func (this *QAudioSink) ProcessedUSecs() int64 {
	return (int64)(QAudioSink_ProcessedUSecs(this.h))
}

func (this *QAudioSink) ElapsedUSecs() int64 {
	return (int64)(QAudioSink_ElapsedUSecs(this.h))
}

func (this *QAudioSink) Error() QtAudio__Error {
	xxxxxxxxx
}

func (this *QAudioSink) State() QtAudio__State {
	xxxxxxxxx
}

func (this *QAudioSink) SetVolume(volume float64) {
	QAudioSink_SetVolume(this.h, (double)(volume))
}

func (this *QAudioSink) Volume() float64 {
	return (float64)(QAudioSink_Volume(this.h))
}

func (this *QAudioSink) StateChanged(state QAudio__State) {
	QAudioSink_StateChanged(this.h, (int)(state))
}

func (this *QAudioSink) OnStateChanged(slot func(state QAudio__State)) {
	QAudioSink_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioSink_StateChanged
func miqt_exec_callback_QAudioSink_StateChanged(cb intptr_t, state int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state QAudio__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudio__State)(state)

	gofunc(slotval1)
}

func QAudioSink_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioSink_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioSink_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioSink_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioSink) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioSink_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioSink) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioSink_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioSink_MetaObject
func miqt_exec_callback_QAudioSink_MetaObject(self QAudioSink, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioSink{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioSink) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioSink_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioSink) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioSink_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioSink_Metacast
func miqt_exec_callback_QAudioSink_Metacast(self QAudioSink, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QAudioSink{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
