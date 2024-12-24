package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioSource struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioSource constructs a new QAudioSource object.
func NewQAudioSource() *QAudioSource {
	ret := newQAudioSource(QAudioSource_new())
	ret.isSubclass = true
	return ret
}

// NewQAudioSource2 constructs a new QAudioSource object.
func NewQAudioSource2(audioDeviceInfo *QAudioDevice) *QAudioSource {
	ret := newQAudioSource(QAudioSource_new2(audioDeviceInfo.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAudioSource3 constructs a new QAudioSource object.
func NewQAudioSource3(format *QAudioFormat) *QAudioSource {
	ret := newQAudioSource(QAudioSource_new3(format.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAudioSource4 constructs a new QAudioSource object.
func NewQAudioSource4(format *QAudioFormat, parent *qt.QObject) *QAudioSource {
	ret := newQAudioSource(QAudioSource_new4(format.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQAudioSource5 constructs a new QAudioSource object.
func NewQAudioSource5(audioDeviceInfo *QAudioDevice, format *QAudioFormat) *QAudioSource {
	ret := newQAudioSource(QAudioSource_new5(audioDeviceInfo.cPointer(), format.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAudioSource6 constructs a new QAudioSource object.
func NewQAudioSource6(audioDeviceInfo *QAudioDevice, format *QAudioFormat, parent *qt.QObject) *QAudioSource {
	ret := newQAudioSource(QAudioSource_new6(audioDeviceInfo.cPointer(), format.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAudioSource) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioSource_MetaObject(this.h)))
}

func (this *QAudioSource) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioSource_Metacast(this.h, param1_Cstring))
}

func QAudioSource_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioSource_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioSource) IsNull() bool {
	return (bool)(QAudioSource_IsNull(this.h))
}

func (this *QAudioSource) Format() *QAudioFormat {
	_goptr := newQAudioFormat(QAudioSource_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioSource) Start(device *qt.QIODevice) {
	QAudioSource_Start(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QAudioSource) Start2() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QAudioSource_Start2(this.h)))
}

func (this *QAudioSource) Stop() {
	QAudioSource_Stop(this.h)
}

func (this *QAudioSource) Reset() {
	QAudioSource_Reset(this.h)
}

func (this *QAudioSource) Suspend() {
	QAudioSource_Suspend(this.h)
}

func (this *QAudioSource) Resume() {
	QAudioSource_Resume(this.h)
}

func (this *QAudioSource) SetBufferSize(bytes int64) {
	QAudioSource_SetBufferSize(this.h, (ptrdiff_t)(bytes))
}

func (this *QAudioSource) BufferSize() int64 {
	return (int64)(QAudioSource_BufferSize(this.h))
}

func (this *QAudioSource) BytesAvailable() int64 {
	return (int64)(QAudioSource_BytesAvailable(this.h))
}

func (this *QAudioSource) SetVolume(volume float64) {
	QAudioSource_SetVolume(this.h, (double)(volume))
}

func (this *QAudioSource) Volume() float64 {
	return (float64)(QAudioSource_Volume(this.h))
}

func (this *QAudioSource) ProcessedUSecs() int64 {
	return (int64)(QAudioSource_ProcessedUSecs(this.h))
}

func (this *QAudioSource) ElapsedUSecs() int64 {
	return (int64)(QAudioSource_ElapsedUSecs(this.h))
}

func (this *QAudioSource) Error() QtAudio__Error {
	xxxxxxxxx
}

func (this *QAudioSource) State() QtAudio__State {
	xxxxxxxxx
}

func (this *QAudioSource) StateChanged(state QAudio__State) {
	QAudioSource_StateChanged(this.h, (int)(state))
}

func (this *QAudioSource) OnStateChanged(slot func(state QAudio__State)) {
	QAudioSource_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioSource_StateChanged
func miqt_exec_callback_QAudioSource_StateChanged(cb intptr_t, state int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state QAudio__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudio__State)(state)

	gofunc(slotval1)
}

func QAudioSource_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioSource_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioSource_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioSource_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioSource) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioSource_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioSource) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioSource_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioSource_MetaObject
func miqt_exec_callback_QAudioSource_MetaObject(self QAudioSource, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioSource{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioSource) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioSource_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioSource) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioSource_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioSource_Metacast
func miqt_exec_callback_QAudioSource_Metacast(self QAudioSource, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAudioSource{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
