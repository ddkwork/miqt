package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioBufferInput struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioBufferInput constructs a new QAudioBufferInput object.
func NewQAudioBufferInput() *QAudioBufferInput {
	ret := newQAudioBufferInput(QAudioBufferInput_new())
	ret.isSubclass = true
	return ret
}

// NewQAudioBufferInput2 constructs a new QAudioBufferInput object.
func NewQAudioBufferInput2(format *QAudioFormat) *QAudioBufferInput {
	ret := newQAudioBufferInput(QAudioBufferInput_new2(format.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAudioBufferInput3 constructs a new QAudioBufferInput object.
func NewQAudioBufferInput3(parent *qt.QObject) *QAudioBufferInput {
	ret := newQAudioBufferInput(QAudioBufferInput_new3((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQAudioBufferInput4 constructs a new QAudioBufferInput object.
func NewQAudioBufferInput4(format *QAudioFormat, parent *qt.QObject) *QAudioBufferInput {
	ret := newQAudioBufferInput(QAudioBufferInput_new4(format.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAudioBufferInput) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioBufferInput_MetaObject(this.h)))
}

func (this *QAudioBufferInput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioBufferInput_Metacast(this.h, param1_Cstring))
}

func QAudioBufferInput_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioBufferInput_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioBufferInput) SendAudioBuffer(audioBuffer *QAudioBuffer) bool {
	return (bool)(QAudioBufferInput_SendAudioBuffer(this.h, audioBuffer.cPointer()))
}

func (this *QAudioBufferInput) Format() *QAudioFormat {
	_goptr := newQAudioFormat(QAudioBufferInput_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioBufferInput) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(QAudioBufferInput_CaptureSession(this.h))
}

func (this *QAudioBufferInput) ReadyToSendAudioBuffer() {
	QAudioBufferInput_ReadyToSendAudioBuffer(this.h)
}

func (this *QAudioBufferInput) OnReadyToSendAudioBuffer(slot func()) {
	QAudioBufferInput_connect_ReadyToSendAudioBuffer(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioBufferInput_ReadyToSendAudioBuffer
func miqt_exec_callback_QAudioBufferInput_ReadyToSendAudioBuffer(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QAudioBufferInput_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioBufferInput_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioBufferInput_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioBufferInput_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioBufferInput) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioBufferInput_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioBufferInput) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioBufferInput_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioBufferInput_MetaObject
func miqt_exec_callback_QAudioBufferInput_MetaObject(self QAudioBufferInput, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioBufferInput{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioBufferInput) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioBufferInput_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioBufferInput) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioBufferInput_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioBufferInput_Metacast
func miqt_exec_callback_QAudioBufferInput_Metacast(self QAudioBufferInput, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAudioBufferInput{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
