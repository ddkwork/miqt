package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QVideoFrameInput struct {
	h          uintptr
	isSubclass bool
}

// NewQVideoFrameInput constructs a new QVideoFrameInput object.
func NewQVideoFrameInput() *QVideoFrameInput {
	g := newQVideoFrameInput(QVideoFrameInput_new())
	g.isSubclass = true
	return g
}

// NewQVideoFrameInput2 constructs a new QVideoFrameInput object.
func NewQVideoFrameInput2(format *QVideoFrameFormat) *QVideoFrameInput {
	g := newQVideoFrameInput(QVideoFrameInput_new2(format.cPointer()))
	g.isSubclass = true
	return g
}

// NewQVideoFrameInput3 constructs a new QVideoFrameInput object.
func NewQVideoFrameInput3(parent *qt.QObject) *QVideoFrameInput {
	g := newQVideoFrameInput(QVideoFrameInput_new3((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQVideoFrameInput4 constructs a new QVideoFrameInput object.
func NewQVideoFrameInput4(format *QVideoFrameFormat, parent *qt.QObject) *QVideoFrameInput {
	g := newQVideoFrameInput(QVideoFrameInput_new4(format.cPointer(), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QVideoFrameInput) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QVideoFrameInput_MetaObject(this.h)))
}

func (this *QVideoFrameInput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QVideoFrameInput_Metacast(this.h, param1_Cstring))
}

func QVideoFrameInput_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QVideoFrameInput_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoFrameInput) SendVideoFrame(frame *QVideoFrame) bool {
	return (bool)(QVideoFrameInput_SendVideoFrame(this.h, frame.cPointer()))
}

func (this *QVideoFrameInput) Format() *QVideoFrameFormat {
	_goptr := newQVideoFrameFormat(QVideoFrameInput_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrameInput) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(QVideoFrameInput_CaptureSession(this.h))
}

func (this *QVideoFrameInput) ReadyToSendVideoFrame() {
	QVideoFrameInput_ReadyToSendVideoFrame(this.h)
}

func (this *QVideoFrameInput) OnReadyToSendVideoFrame(slot func()) {
	QVideoFrameInput_connect_ReadyToSendVideoFrame(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoFrameInput_ReadyToSendVideoFrame
func miqt_exec_callback_QVideoFrameInput_ReadyToSendVideoFrame(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QVideoFrameInput_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVideoFrameInput_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoFrameInput_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVideoFrameInput_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoFrameInput) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QVideoFrameInput_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QVideoFrameInput) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVideoFrameInput_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoFrameInput_MetaObject
func miqt_exec_callback_QVideoFrameInput_MetaObject(self QVideoFrameInput, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVideoFrameInput{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QVideoFrameInput) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QVideoFrameInput_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QVideoFrameInput) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVideoFrameInput_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoFrameInput_Metacast
func miqt_exec_callback_QVideoFrameInput_Metacast(self QVideoFrameInput, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QVideoFrameInput{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
