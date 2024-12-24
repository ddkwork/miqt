package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioInput struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioInput constructs a new QAudioInput object.
func NewQAudioInput() *QAudioInput {
	ret := newQAudioInput(QAudioInput_new())
	ret.isSubclass = true
	return ret
}

// NewQAudioInput2 constructs a new QAudioInput object.
func NewQAudioInput2(deviceInfo *QAudioDevice) *QAudioInput {
	ret := newQAudioInput(QAudioInput_new2(deviceInfo.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAudioInput3 constructs a new QAudioInput object.
func NewQAudioInput3(parent *qt.QObject) *QAudioInput {
	ret := newQAudioInput(QAudioInput_new3((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQAudioInput4 constructs a new QAudioInput object.
func NewQAudioInput4(deviceInfo *QAudioDevice, parent *qt.QObject) *QAudioInput {
	ret := newQAudioInput(QAudioInput_new4(deviceInfo.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAudioInput) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioInput_MetaObject(this.h)))
}

func (this *QAudioInput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioInput_Metacast(this.h, param1_Cstring))
}

func QAudioInput_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioInput_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioInput) Device() *QAudioDevice {
	_goptr := newQAudioDevice(QAudioInput_Device(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioInput) Volume() float32 {
	return (float32)(QAudioInput_Volume(this.h))
}

func (this *QAudioInput) IsMuted() bool {
	return (bool)(QAudioInput_IsMuted(this.h))
}

func (this *QAudioInput) SetDevice(device *QAudioDevice) {
	QAudioInput_SetDevice(this.h, device.cPointer())
}

func (this *QAudioInput) SetVolume(volume float32) {
	QAudioInput_SetVolume(this.h, (float)(volume))
}

func (this *QAudioInput) SetMuted(muted bool) {
	QAudioInput_SetMuted(this.h, (bool)(muted))
}

func (this *QAudioInput) DeviceChanged() {
	QAudioInput_DeviceChanged(this.h)
}

func (this *QAudioInput) OnDeviceChanged(slot func()) {
	QAudioInput_connect_DeviceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_DeviceChanged
func miqt_exec_callback_QAudioInput_DeviceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioInput) VolumeChanged(volume float32) {
	QAudioInput_VolumeChanged(this.h, (float)(volume))
}

func (this *QAudioInput) OnVolumeChanged(slot func(volume float32)) {
	QAudioInput_connect_VolumeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_VolumeChanged
func miqt_exec_callback_QAudioInput_VolumeChanged(cb intptr_t, volume float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(volume float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(volume)

	gofunc(slotval1)
}

func (this *QAudioInput) MutedChanged(muted bool) {
	QAudioInput_MutedChanged(this.h, (bool)(muted))
}

func (this *QAudioInput) OnMutedChanged(slot func(muted bool)) {
	QAudioInput_connect_MutedChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_MutedChanged
func miqt_exec_callback_QAudioInput_MutedChanged(cb intptr_t, muted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(muted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(muted)

	gofunc(slotval1)
}

func QAudioInput_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioInput_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioInput_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioInput_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioInput) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioInput_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioInput) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioInput_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_MetaObject
func miqt_exec_callback_QAudioInput_MetaObject(self QAudioInput, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioInput{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioInput) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioInput_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioInput) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioInput_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_Metacast
func miqt_exec_callback_QAudioInput_Metacast(self QAudioInput, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAudioInput{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
