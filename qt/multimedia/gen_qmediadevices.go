package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QMediaDevices struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaDevices constructs a new QMediaDevices object.
func NewQMediaDevices() *QMediaDevices {
	g := newQMediaDevices(QMediaDevices_new())
	g.isSubclass = true
	return g
}

// NewQMediaDevices2 constructs a new QMediaDevices object.
func NewQMediaDevices2(parent *qt.QObject) *QMediaDevices {
	g := newQMediaDevices(QMediaDevices_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QMediaDevices) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaDevices_MetaObject(this.h)))
}

func (this *QMediaDevices) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMediaDevices_Metacast(this.h, param1_Cstring))
}

func QMediaDevices_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMediaDevices_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaDevices_AudioInputs() []QAudioDevice {
	var _ma struct_miqt_array = QMediaDevices_AudioInputs()
	_ret := make([]QAudioDevice, int(_ma.len))
	_outCast := (*[0xffff]*QAudioDevice)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQAudioDevice(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QMediaDevices_AudioOutputs() []QAudioDevice {
	var _ma struct_miqt_array = QMediaDevices_AudioOutputs()
	_ret := make([]QAudioDevice, int(_ma.len))
	_outCast := (*[0xffff]*QAudioDevice)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQAudioDevice(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QMediaDevices_VideoInputs() []QCameraDevice {
	var _ma struct_miqt_array = QMediaDevices_VideoInputs()
	_ret := make([]QCameraDevice, int(_ma.len))
	_outCast := (*[0xffff]*QCameraDevice)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQCameraDevice(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QMediaDevices_DefaultAudioInput() *QAudioDevice {
	_goptr := newQAudioDevice(QMediaDevices_DefaultAudioInput())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QMediaDevices_DefaultAudioOutput() *QAudioDevice {
	_goptr := newQAudioDevice(QMediaDevices_DefaultAudioOutput())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QMediaDevices_DefaultVideoInput() *QCameraDevice {
	_goptr := newQCameraDevice(QMediaDevices_DefaultVideoInput())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaDevices) AudioInputsChanged() {
	QMediaDevices_AudioInputsChanged(this.h)
}

func (this *QMediaDevices) OnAudioInputsChanged(slot func()) {
	QMediaDevices_connect_AudioInputsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaDevices_AudioInputsChanged
func miqt_exec_callback_QMediaDevices_AudioInputsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaDevices) AudioOutputsChanged() {
	QMediaDevices_AudioOutputsChanged(this.h)
}

func (this *QMediaDevices) OnAudioOutputsChanged(slot func()) {
	QMediaDevices_connect_AudioOutputsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaDevices_AudioOutputsChanged
func miqt_exec_callback_QMediaDevices_AudioOutputsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMediaDevices) VideoInputsChanged() {
	QMediaDevices_VideoInputsChanged(this.h)
}

func (this *QMediaDevices) OnVideoInputsChanged(slot func()) {
	QMediaDevices_connect_VideoInputsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaDevices_VideoInputsChanged
func miqt_exec_callback_QMediaDevices_VideoInputsChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QMediaDevices_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaDevices_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaDevices_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMediaDevices_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaDevices) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QMediaDevices_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QMediaDevices) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaDevices_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaDevices_MetaObject
func miqt_exec_callback_QMediaDevices_MetaObject(self QMediaDevices, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMediaDevices{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QMediaDevices) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMediaDevices_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMediaDevices) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMediaDevices_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaDevices_Metacast
func miqt_exec_callback_QMediaDevices_Metacast(self QMediaDevices, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QMediaDevices{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
