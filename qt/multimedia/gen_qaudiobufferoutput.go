package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioBufferOutput struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioBufferOutput constructs a new QAudioBufferOutput object.
func NewQAudioBufferOutput() *QAudioBufferOutput {
	ret := newQAudioBufferOutput(QAudioBufferOutput_new())
	ret.isSubclass = true
	return ret
}

// NewQAudioBufferOutput2 constructs a new QAudioBufferOutput object.
func NewQAudioBufferOutput2(format *QAudioFormat) *QAudioBufferOutput {
	ret := newQAudioBufferOutput(QAudioBufferOutput_new2(format.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAudioBufferOutput3 constructs a new QAudioBufferOutput object.
func NewQAudioBufferOutput3(parent *qt.QObject) *QAudioBufferOutput {
	ret := newQAudioBufferOutput(QAudioBufferOutput_new3((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQAudioBufferOutput4 constructs a new QAudioBufferOutput object.
func NewQAudioBufferOutput4(format *QAudioFormat, parent *qt.QObject) *QAudioBufferOutput {
	ret := newQAudioBufferOutput(QAudioBufferOutput_new4(format.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAudioBufferOutput) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioBufferOutput_MetaObject(this.h)))
}

func (this *QAudioBufferOutput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioBufferOutput_Metacast(this.h, param1_Cstring))
}

func QAudioBufferOutput_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioBufferOutput_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioBufferOutput) Format() *QAudioFormat {
	_goptr := newQAudioFormat(QAudioBufferOutput_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioBufferOutput) AudioBufferReceived(buffer *QAudioBuffer) {
	QAudioBufferOutput_AudioBufferReceived(this.h, buffer.cPointer())
}

func (this *QAudioBufferOutput) OnAudioBufferReceived(slot func(buffer *QAudioBuffer)) {
	QAudioBufferOutput_connect_AudioBufferReceived(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioBufferOutput_AudioBufferReceived
func miqt_exec_callback_QAudioBufferOutput_AudioBufferReceived(cb intptr_t, buffer *QAudioBuffer) {
	gofunc, ok := cgo.Handle(cb).Value().(func(buffer *QAudioBuffer))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAudioBuffer(buffer)

	gofunc(slotval1)
}

func QAudioBufferOutput_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioBufferOutput_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioBufferOutput_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioBufferOutput_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioBufferOutput) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioBufferOutput_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioBufferOutput) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioBufferOutput_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioBufferOutput_MetaObject
func miqt_exec_callback_QAudioBufferOutput_MetaObject(self QAudioBufferOutput, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioBufferOutput{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioBufferOutput) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioBufferOutput_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioBufferOutput) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioBufferOutput_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioBufferOutput_Metacast
func miqt_exec_callback_QAudioBufferOutput_Metacast(self QAudioBufferOutput, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAudioBufferOutput{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
