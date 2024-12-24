package qt

import (
	"unsafe"
)

type QPointingDevice__PointerType int

const (
	QPointingDevice__Unknown         QPointingDevice__PointerType = 0
	QPointingDevice__Generic         QPointingDevice__PointerType = 1
	QPointingDevice__Finger          QPointingDevice__PointerType = 2
	QPointingDevice__Pen             QPointingDevice__PointerType = 4
	QPointingDevice__Eraser          QPointingDevice__PointerType = 8
	QPointingDevice__Cursor          QPointingDevice__PointerType = 16
	QPointingDevice__AllPointerTypes QPointingDevice__PointerType = 32767
)

type QPointingDevice__GrabTransition int

const (
	QPointingDevice__GrabPassive         QPointingDevice__GrabTransition = 1
	QPointingDevice__UngrabPassive       QPointingDevice__GrabTransition = 2
	QPointingDevice__CancelGrabPassive   QPointingDevice__GrabTransition = 3
	QPointingDevice__OverrideGrabPassive QPointingDevice__GrabTransition = 4
	QPointingDevice__GrabExclusive       QPointingDevice__GrabTransition = 16
	QPointingDevice__UngrabExclusive     QPointingDevice__GrabTransition = 32
	QPointingDevice__CancelGrabExclusive QPointingDevice__GrabTransition = 48
)

type QPointingDeviceUniqueId struct {
	h          uintptr
	isSubclass bool
}

// NewQPointingDeviceUniqueId constructs a new QPointingDeviceUniqueId object.
func NewQPointingDeviceUniqueId() *QPointingDeviceUniqueId {
	g := newQPointingDeviceUniqueId(QPointingDeviceUniqueId_new())
	g.isSubclass = true
	return g
}

// NewQPointingDeviceUniqueId2 constructs a new QPointingDeviceUniqueId object.
func NewQPointingDeviceUniqueId2(param1 *QPointingDeviceUniqueId) *QPointingDeviceUniqueId {
	g := newQPointingDeviceUniqueId(QPointingDeviceUniqueId_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

func QPointingDeviceUniqueId_FromNumericId(id int64) *QPointingDeviceUniqueId {
	_goptr := newQPointingDeviceUniqueId(QPointingDeviceUniqueId_FromNumericId((longlong)(id)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPointingDeviceUniqueId) IsValid() bool {
	return (bool)(QPointingDeviceUniqueId_IsValid(this.h))
}

func (this *QPointingDeviceUniqueId) NumericId() int64 {
	return (int64)(QPointingDeviceUniqueId_NumericId(this.h))
}

type QPointingDevice struct {
	h          uintptr
	isSubclass bool
}

// NewQPointingDevice constructs a new QPointingDevice object.
func NewQPointingDevice() *QPointingDevice {
	g := newQPointingDevice(QPointingDevice_new())
	g.isSubclass = true
	return g
}

// NewQPointingDevice2 constructs a new QPointingDevice object.
func NewQPointingDevice2(name string, systemId int64, devType QInputDevice__DeviceType, pType PointerType, caps Capabilities, maxPoints int, buttonCount int) *QPointingDevice {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	g := newQPointingDevice(QPointingDevice_new2(name_ms, (longlong)(systemId), (int)(devType), pType, caps, (int)(maxPoints), (int)(buttonCount)))
	g.isSubclass = true
	return g
}

// NewQPointingDevice3 constructs a new QPointingDevice object.
func NewQPointingDevice3(parent *QObject) *QPointingDevice {
	g := newQPointingDevice(QPointingDevice_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPointingDevice4 constructs a new QPointingDevice object.
func NewQPointingDevice4(name string, systemId int64, devType QInputDevice__DeviceType, pType PointerType, caps Capabilities, maxPoints int, buttonCount int, seatName string) *QPointingDevice {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	seatName_ms := struct_miqt_string{}
	seatName_ms.data = CString(seatName)
	seatName_ms.len = size_t(len(seatName))
	defer free(unsafe.Pointer(seatName_ms.data))
	g := newQPointingDevice(QPointingDevice_new4(name_ms, (longlong)(systemId), (int)(devType), pType, caps, (int)(maxPoints), (int)(buttonCount), seatName_ms))
	g.isSubclass = true
	return g
}

// NewQPointingDevice5 constructs a new QPointingDevice object.
func NewQPointingDevice5(name string, systemId int64, devType QInputDevice__DeviceType, pType PointerType, caps Capabilities, maxPoints int, buttonCount int, seatName string, uniqueId QPointingDeviceUniqueId) *QPointingDevice {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	seatName_ms := struct_miqt_string{}
	seatName_ms.data = CString(seatName)
	seatName_ms.len = size_t(len(seatName))
	defer free(unsafe.Pointer(seatName_ms.data))
	g := newQPointingDevice(QPointingDevice_new5(name_ms, (longlong)(systemId), (int)(devType), pType, caps, (int)(maxPoints), (int)(buttonCount), seatName_ms, uniqueId.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPointingDevice6 constructs a new QPointingDevice object.
func NewQPointingDevice6(name string, systemId int64, devType QInputDevice__DeviceType, pType PointerType, caps Capabilities, maxPoints int, buttonCount int, seatName string, uniqueId QPointingDeviceUniqueId, parent *QObject) *QPointingDevice {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	seatName_ms := struct_miqt_string{}
	seatName_ms.data = CString(seatName)
	seatName_ms.len = size_t(len(seatName))
	defer free(unsafe.Pointer(seatName_ms.data))
	g := newQPointingDevice(QPointingDevice_new6(name_ms, (longlong)(systemId), (int)(devType), pType, caps, (int)(maxPoints), (int)(buttonCount), seatName_ms, uniqueId.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPointingDevice) MetaObject() *QMetaObject {
	return newQMetaObject(QPointingDevice_MetaObject(this.h))
}

func (this *QPointingDevice) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPointingDevice_Metacast(this.h, param1_Cstring))
}

func QPointingDevice_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPointingDevice_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPointingDevice) SetType(devType DeviceType) {
	QPointingDevice_SetType(this.h, devType)
}

func (this *QPointingDevice) SetCapabilities(caps Capability) {
	QPointingDevice_SetCapabilities(this.h, (int)(caps))
}

func (this *QPointingDevice) SetMaximumTouchPoints(c int) {
	QPointingDevice_SetMaximumTouchPoints(this.h, (int)(c))
}

func (this *QPointingDevice) PointerType() PointerType {
	xxxxxxxxx
}

func (this *QPointingDevice) MaximumPoints() int {
	return (int)(QPointingDevice_MaximumPoints(this.h))
}

func (this *QPointingDevice) ButtonCount() int {
	return (int)(QPointingDevice_ButtonCount(this.h))
}

func (this *QPointingDevice) UniqueId() *QPointingDeviceUniqueId {
	_goptr := newQPointingDeviceUniqueId(QPointingDevice_UniqueId(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPointingDevice_PrimaryPointingDevice() *QPointingDevice {
	return newQPointingDevice(QPointingDevice_PrimaryPointingDevice())
}

func (this *QPointingDevice) OperatorEqual(other *QPointingDevice) bool {
	return (bool)(QPointingDevice_OperatorEqual(this.h, other.cPointer()))
}

func (this *QPointingDevice) GrabChanged(grabber *QObject, transition GrabTransition, event *QPointerEvent, point *QEventPoint) {
	QPointingDevice_GrabChanged(this.h, grabber.cPointer(), transition, event.cPointer(), point.cPointer())
}

func (this *QPointingDevice) OnGrabChanged(slot func(grabber *QObject, transition GrabTransition, event *QPointerEvent, point *QEventPoint)) {
	QPointingDevice_connect_GrabChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointingDevice_GrabChanged
func miqt_exec_callback_QPointingDevice_GrabChanged(cb intptr_t, grabber *QObject, transition GrabTransition, event *QPointerEvent, point *QEventPoint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(grabber *QObject, transition GrabTransition, event *QPointerEvent, point *QEventPoint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(grabber)

	xxxxxxxxx
	slotval3 := newQPointerEvent(event)

	slotval4 := newQEventPoint(point)

	gofunc(slotval1, slotval2, slotval3, slotval4)
}

func QPointingDevice_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPointingDevice_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPointingDevice_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPointingDevice_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPointingDevice_PrimaryPointingDevice1(seatName string) *QPointingDevice {
	seatName_ms := struct_miqt_string{}
	seatName_ms.data = CString(seatName)
	seatName_ms.len = size_t(len(seatName))
	defer free(unsafe.Pointer(seatName_ms.data))
	return newQPointingDevice(QPointingDevice_PrimaryPointingDevice1(seatName_ms))
}

func (this *QPointingDevice) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPointingDevice_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPointingDevice) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointingDevice_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointingDevice_MetaObject
func miqt_exec_callback_QPointingDevice_MetaObject(self QPointingDevice, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPointingDevice{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPointingDevice) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPointingDevice_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPointingDevice) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPointingDevice_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPointingDevice_Metacast
func miqt_exec_callback_QPointingDevice_Metacast(self QPointingDevice, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QPointingDevice{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
