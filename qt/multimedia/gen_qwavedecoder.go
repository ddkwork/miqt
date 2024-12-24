package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QWaveDecoder struct {
	h          uintptr
	isSubclass bool
}

// NewQWaveDecoder constructs a new QWaveDecoder object.
func NewQWaveDecoder(device *qt.QIODevice) *QWaveDecoder {
	ret := newQWaveDecoder(QWaveDecoder_new((*QIODevice)(device.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQWaveDecoder2 constructs a new QWaveDecoder object.
func NewQWaveDecoder2(device *qt.QIODevice, format *QAudioFormat) *QWaveDecoder {
	ret := newQWaveDecoder(QWaveDecoder_new2((*QIODevice)(device.UnsafePointer()), format.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQWaveDecoder3 constructs a new QWaveDecoder object.
func NewQWaveDecoder3(device *qt.QIODevice, parent *qt.QObject) *QWaveDecoder {
	ret := newQWaveDecoder(QWaveDecoder_new3((*QIODevice)(device.UnsafePointer()), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQWaveDecoder4 constructs a new QWaveDecoder object.
func NewQWaveDecoder4(device *qt.QIODevice, format *QAudioFormat, parent *qt.QObject) *QWaveDecoder {
	ret := newQWaveDecoder(QWaveDecoder_new4((*QIODevice)(device.UnsafePointer()), format.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QWaveDecoder) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWaveDecoder_MetaObject(this.h)))
}

func (this *QWaveDecoder) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWaveDecoder_Metacast(this.h, param1_Cstring))
}

func QWaveDecoder_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWaveDecoder_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWaveDecoder) AudioFormat() *QAudioFormat {
	_goptr := newQAudioFormat(QWaveDecoder_AudioFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWaveDecoder) GetDevice() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QWaveDecoder_GetDevice(this.h)))
}

func (this *QWaveDecoder) Duration() int {
	return (int)(QWaveDecoder_Duration(this.h))
}

func QWaveDecoder_HeaderLength() int64 {
	return (int64)(QWaveDecoder_HeaderLength())
}

func (this *QWaveDecoder) Open(mode OpenModeFlag) bool {
	return (bool)(QWaveDecoder_Open(this.h, (int)(mode)))
}

func (this *QWaveDecoder) Close() {
	QWaveDecoder_Close(this.h)
}

func (this *QWaveDecoder) Seek(pos int64) bool {
	return (bool)(QWaveDecoder_Seek(this.h, (longlong)(pos)))
}

func (this *QWaveDecoder) Pos() int64 {
	return (int64)(QWaveDecoder_Pos(this.h))
}

func (this *QWaveDecoder) Size() int64 {
	return (int64)(QWaveDecoder_Size(this.h))
}

func (this *QWaveDecoder) IsSequential() bool {
	return (bool)(QWaveDecoder_IsSequential(this.h))
}

func (this *QWaveDecoder) BytesAvailable() int64 {
	return (int64)(QWaveDecoder_BytesAvailable(this.h))
}

func (this *QWaveDecoder) FormatKnown() {
	QWaveDecoder_FormatKnown(this.h)
}

func (this *QWaveDecoder) OnFormatKnown(slot func()) {
	QWaveDecoder_connect_FormatKnown(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWaveDecoder_FormatKnown
func miqt_exec_callback_QWaveDecoder_FormatKnown(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWaveDecoder) ParsingError() {
	QWaveDecoder_ParsingError(this.h)
}

func (this *QWaveDecoder) OnParsingError(slot func()) {
	QWaveDecoder_connect_ParsingError(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWaveDecoder_ParsingError
func miqt_exec_callback_QWaveDecoder_ParsingError(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QWaveDecoder_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWaveDecoder_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWaveDecoder_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWaveDecoder_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWaveDecoder) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWaveDecoder_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QWaveDecoder) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWaveDecoder_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWaveDecoder_MetaObject
func miqt_exec_callback_QWaveDecoder_MetaObject(self QWaveDecoder, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWaveDecoder{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QWaveDecoder) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWaveDecoder_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWaveDecoder) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWaveDecoder_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWaveDecoder_Metacast
func miqt_exec_callback_QWaveDecoder_Metacast(self QWaveDecoder, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QWaveDecoder{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
