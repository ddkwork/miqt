package qt

import (
	"unsafe"
)

type QInputDevice__DeviceType int

const (
	QInputDevice__Unknown     QInputDevice__DeviceType = 0
	QInputDevice__Mouse       QInputDevice__DeviceType = 1
	QInputDevice__TouchScreen QInputDevice__DeviceType = 2
	QInputDevice__TouchPad    QInputDevice__DeviceType = 4
	QInputDevice__Puck        QInputDevice__DeviceType = 8
	QInputDevice__Stylus      QInputDevice__DeviceType = 16
	QInputDevice__Airbrush    QInputDevice__DeviceType = 32
	QInputDevice__Keyboard    QInputDevice__DeviceType = 4096
	QInputDevice__AllDevices  QInputDevice__DeviceType = 2147483647
)

type QInputDevice__Capability int

const (
	QInputDevice__None               QInputDevice__Capability = 0
	QInputDevice__Position           QInputDevice__Capability = 1
	QInputDevice__Area               QInputDevice__Capability = 2
	QInputDevice__Pressure           QInputDevice__Capability = 4
	QInputDevice__Velocity           QInputDevice__Capability = 8
	QInputDevice__NormalizedPosition QInputDevice__Capability = 32
	QInputDevice__MouseEmulation     QInputDevice__Capability = 64
	QInputDevice__PixelScroll        QInputDevice__Capability = 128
	QInputDevice__Scroll             QInputDevice__Capability = 256
	QInputDevice__Hover              QInputDevice__Capability = 512
	QInputDevice__Rotation           QInputDevice__Capability = 1024
	QInputDevice__XTilt              QInputDevice__Capability = 2048
	QInputDevice__YTilt              QInputDevice__Capability = 4096
	QInputDevice__TangentialPressure QInputDevice__Capability = 8192
	QInputDevice__ZPosition          QInputDevice__Capability = 16384
	QInputDevice__All                QInputDevice__Capability = 2147483647
)

type QInputDevice struct {
	h          uintptr
	isSubclass bool
}

// NewQInputDevice constructs a new QInputDevice object.
func NewQInputDevice() *QInputDevice {
	ret := newQInputDevice(QInputDevice_new())
	ret.isSubclass = true
	return ret
}

// NewQInputDevice2 constructs a new QInputDevice object.
func NewQInputDevice2(name string, systemId int64, typeVal DeviceType) *QInputDevice {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQInputDevice(QInputDevice_new2(name_ms, (longlong)(systemId), typeVal))
	ret.isSubclass = true
	return ret
}

// NewQInputDevice3 constructs a new QInputDevice object.
func NewQInputDevice3(parent *QObject) *QInputDevice {
	ret := newQInputDevice(QInputDevice_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQInputDevice4 constructs a new QInputDevice object.
func NewQInputDevice4(name string, systemId int64, typeVal DeviceType, seatName string) *QInputDevice {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	seatName_ms := struct_miqt_string{}
	seatName_ms.data = CString(seatName)
	seatName_ms.len = size_t(len(seatName))
	defer free(unsafe.Pointer(seatName_ms.data))

	ret := newQInputDevice(QInputDevice_new4(name_ms, (longlong)(systemId), typeVal, seatName_ms))
	ret.isSubclass = true
	return ret
}

// NewQInputDevice5 constructs a new QInputDevice object.
func NewQInputDevice5(name string, systemId int64, typeVal DeviceType, seatName string, parent *QObject) *QInputDevice {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	seatName_ms := struct_miqt_string{}
	seatName_ms.data = CString(seatName)
	seatName_ms.len = size_t(len(seatName))
	defer free(unsafe.Pointer(seatName_ms.data))

	ret := newQInputDevice(QInputDevice_new5(name_ms, (longlong)(systemId), typeVal, seatName_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QInputDevice) MetaObject() *QMetaObject {
	return newQMetaObject(QInputDevice_MetaObject(this.h))
}

func (this *QInputDevice) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QInputDevice_Metacast(this.h, param1_Cstring))
}

func QInputDevice_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QInputDevice_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDevice) Name() string {
	var _ms struct_miqt_string = QInputDevice_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDevice) Type() DeviceType {
	xxxxxxxxx
}

func (this *QInputDevice) Capabilities() Capabilities {
	xxxxxxxxx
}

func (this *QInputDevice) HasCapability(cap Capability) bool {
	return (bool)(QInputDevice_HasCapability(this.h, cap))
}

func (this *QInputDevice) SystemId() int64 {
	return (int64)(QInputDevice_SystemId(this.h))
}

func (this *QInputDevice) SeatName() string {
	var _ms struct_miqt_string = QInputDevice_SeatName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDevice) AvailableVirtualGeometry() *QRect {
	_goptr := newQRect(QInputDevice_AvailableVirtualGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QInputDevice_SeatNames() []string {
	var _ma struct_miqt_array = QInputDevice_SeatNames()
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QInputDevice_Devices() []*QInputDevice {
	var _ma struct_miqt_array = QInputDevice_Devices()
	_ret := make([]*QInputDevice, int(_ma.len))
	_outCast := (*[0xffff]*QInputDevice)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQInputDevice(_outCast[i])
	}
	return _ret
}

func QInputDevice_PrimaryKeyboard() *QInputDevice {
	return newQInputDevice(QInputDevice_PrimaryKeyboard())
}

func (this *QInputDevice) OperatorEqual(other *QInputDevice) bool {
	return (bool)(QInputDevice_OperatorEqual(this.h, other.cPointer()))
}

func (this *QInputDevice) AvailableVirtualGeometryChanged(area QRect) {
	QInputDevice_AvailableVirtualGeometryChanged(this.h, area.cPointer())
}

func (this *QInputDevice) OnAvailableVirtualGeometryChanged(slot func(area QRect)) {
	QInputDevice_connect_AvailableVirtualGeometryChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDevice_AvailableVirtualGeometryChanged
func miqt_exec_callback_QInputDevice_AvailableVirtualGeometryChanged(cb intptr_t, area *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(area QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	area_goptr := newQRect(area)
	area_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *area_goptr

	gofunc(slotval1)
}

func QInputDevice_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QInputDevice_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDevice_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QInputDevice_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDevice_PrimaryKeyboard1(seatName string) *QInputDevice {
	seatName_ms := struct_miqt_string{}
	seatName_ms.data = CString(seatName)
	seatName_ms.len = size_t(len(seatName))
	defer free(unsafe.Pointer(seatName_ms.data))
	return newQInputDevice(QInputDevice_PrimaryKeyboard1(seatName_ms))
}

func (this *QInputDevice) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QInputDevice_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QInputDevice) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputDevice_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDevice_MetaObject
func miqt_exec_callback_QInputDevice_MetaObject(self QInputDevice, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputDevice{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QInputDevice) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QInputDevice_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QInputDevice) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputDevice_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDevice_Metacast
func miqt_exec_callback_QInputDevice_Metacast(self QInputDevice, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QInputDevice{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
