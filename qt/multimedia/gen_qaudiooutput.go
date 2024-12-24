package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioOutput struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioOutput constructs a new QAudioOutput object.
func NewQAudioOutput() *QAudioOutput {
	g := newQAudioOutput(QAudioOutput_new())
	g.isSubclass = true
	return g
}

// NewQAudioOutput2 constructs a new QAudioOutput object.
func NewQAudioOutput2(device *QAudioDevice) *QAudioOutput {
	g := newQAudioOutput(QAudioOutput_new2(device.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAudioOutput3 constructs a new QAudioOutput object.
func NewQAudioOutput3(parent *qt.QObject) *QAudioOutput {
	g := newQAudioOutput(QAudioOutput_new3((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQAudioOutput4 constructs a new QAudioOutput object.
func NewQAudioOutput4(device *QAudioDevice, parent *qt.QObject) *QAudioOutput {
	g := newQAudioOutput(QAudioOutput_new4(device.cPointer(), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QAudioOutput) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioOutput_MetaObject(this.h)))
}

func (this *QAudioOutput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAudioOutput_Metacast(this.h, param1_Cstring))
}

func QAudioOutput_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAudioOutput_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioOutput) Device() *QAudioDevice {
	_goptr := newQAudioDevice(QAudioOutput_Device(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioOutput) Volume() float32 {
	return (float32)(QAudioOutput_Volume(this.h))
}

func (this *QAudioOutput) IsMuted() bool {
	return (bool)(QAudioOutput_IsMuted(this.h))
}

func (this *QAudioOutput) SetDevice(device *QAudioDevice) {
	QAudioOutput_SetDevice(this.h, device.cPointer())
}

func (this *QAudioOutput) SetVolume(volume float32) {
	QAudioOutput_SetVolume(this.h, (float)(volume))
}

func (this *QAudioOutput) SetMuted(muted bool) {
	QAudioOutput_SetMuted(this.h, (bool)(muted))
}

func (this *QAudioOutput) DeviceChanged() {
	QAudioOutput_DeviceChanged(this.h)
}

func (this *QAudioOutput) OnDeviceChanged(slot func()) {
	QAudioOutput_connect_DeviceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioOutput_DeviceChanged
func miqt_exec_callback_QAudioOutput_DeviceChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioOutput) VolumeChanged(volume float32) {
	QAudioOutput_VolumeChanged(this.h, (float)(volume))
}

func (this *QAudioOutput) OnVolumeChanged(slot func(volume float32)) {
	QAudioOutput_connect_VolumeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioOutput_VolumeChanged
func miqt_exec_callback_QAudioOutput_VolumeChanged(cb intptr_t, volume float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(volume float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(volume)

	gofunc(slotval1)
}

func (this *QAudioOutput) MutedChanged(muted bool) {
	QAudioOutput_MutedChanged(this.h, (bool)(muted))
}

func (this *QAudioOutput) OnMutedChanged(slot func(muted bool)) {
	QAudioOutput_connect_MutedChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioOutput_MutedChanged
func miqt_exec_callback_QAudioOutput_MutedChanged(cb intptr_t, muted bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(muted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(muted)

	gofunc(slotval1)
}

func QAudioOutput_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioOutput_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioOutput_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAudioOutput_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioOutput) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioOutput_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioOutput) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioOutput_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioOutput_MetaObject
func miqt_exec_callback_QAudioOutput_MetaObject(self QAudioOutput, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioOutput{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioOutput) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioOutput_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioOutput) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioOutput_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioOutput_Metacast
func miqt_exec_callback_QAudioOutput_Metacast(self QAudioOutput, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QAudioOutput{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
