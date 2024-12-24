package qt

import (
	"unsafe"
)

type QLocationPermission__Accuracy byte

const (
	QLocationPermission__Approximate QLocationPermission__Accuracy = 0
	QLocationPermission__Precise     QLocationPermission__Accuracy = 1
)

type QLocationPermission__Availability byte

const (
	QLocationPermission__WhenInUse QLocationPermission__Availability = 0
	QLocationPermission__Always    QLocationPermission__Availability = 1
)

type QCalendarPermission__AccessMode byte

const (
	QCalendarPermission__ReadOnly  QCalendarPermission__AccessMode = 0
	QCalendarPermission__ReadWrite QCalendarPermission__AccessMode = 1
)

type QContactsPermission__AccessMode byte

const (
	QContactsPermission__ReadOnly  QContactsPermission__AccessMode = 0
	QContactsPermission__ReadWrite QContactsPermission__AccessMode = 1
)

type QBluetoothPermission__CommunicationMode byte

const (
	QBluetoothPermission__Access    QBluetoothPermission__CommunicationMode = 1
	QBluetoothPermission__Advertise QBluetoothPermission__CommunicationMode = 2
	QBluetoothPermission__Default   QBluetoothPermission__CommunicationMode = 3
)

type QPermission struct {
	h          uintptr
	isSubclass bool
}

// NewQPermission constructs a new QPermission object.
func NewQPermission() *QPermission {
	g := newQPermission(QPermission_new())
	g.isSubclass = true
	return g
}

// NewQPermission2 constructs a new QPermission object.
func NewQPermission2(param1 *QPermission) *QPermission {
	g := newQPermission(QPermission_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPermission) Status() PermissionStatus {
	return (PermissionStatus)(QPermission_Status(this.h))
}

func (this *QPermission) Type() *QMetaType {
	_goptr := newQMetaType(QPermission_Type(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QLocationPermission struct {
	h          uintptr
	isSubclass bool
}

// NewQLocationPermission constructs a new QLocationPermission object.
func NewQLocationPermission() *QLocationPermission {
	g := newQLocationPermission(QLocationPermission_new())
	g.isSubclass = true
	return g
}

// NewQLocationPermission2 constructs a new QLocationPermission object.
func NewQLocationPermission2(other *QLocationPermission) *QLocationPermission {
	g := newQLocationPermission(QLocationPermission_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QLocationPermission) SetAccuracy(accuracy Accuracy) {
	QLocationPermission_SetAccuracy(this.h, accuracy)
}

func (this *QLocationPermission) Accuracy() Accuracy {
	xxxxxxxxx
}

func (this *QLocationPermission) SetAvailability(availability Availability) {
	QLocationPermission_SetAvailability(this.h, availability)
}

func (this *QLocationPermission) Availability() Availability {
	xxxxxxxxx
}

func (this *QLocationPermission) OperatorAssign(other *QLocationPermission) {
	QLocationPermission_OperatorAssign(this.h, other.cPointer())
}

func (this *QLocationPermission) OnOperatorAssign(slot func(other *QLocationPermission)) {
	QLocationPermission_connect_OperatorAssign(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocationPermission_OperatorAssign
func miqt_exec_callback_QLocationPermission_OperatorAssign(cb intptr_t, other *QLocationPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QLocationPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLocationPermission(other)

	gofunc(slotval1)
}

func (this *QLocationPermission) Swap(other *QLocationPermission) {
	QLocationPermission_Swap(this.h, other.cPointer())
}

func (this *QLocationPermission) OnSwap(slot func(other *QLocationPermission)) {
	QLocationPermission_connect_Swap(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLocationPermission_Swap
func miqt_exec_callback_QLocationPermission_Swap(cb intptr_t, other *QLocationPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QLocationPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLocationPermission(other)

	gofunc(slotval1)
}

type QCalendarPermission struct {
	h          uintptr
	isSubclass bool
}

// NewQCalendarPermission constructs a new QCalendarPermission object.
func NewQCalendarPermission() *QCalendarPermission {
	g := newQCalendarPermission(QCalendarPermission_new())
	g.isSubclass = true
	return g
}

// NewQCalendarPermission2 constructs a new QCalendarPermission object.
func NewQCalendarPermission2(other *QCalendarPermission) *QCalendarPermission {
	g := newQCalendarPermission(QCalendarPermission_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QCalendarPermission) SetAccessMode(mode AccessMode) {
	QCalendarPermission_SetAccessMode(this.h, mode)
}

func (this *QCalendarPermission) AccessMode() AccessMode {
	xxxxxxxxx
}

func (this *QCalendarPermission) OperatorAssign(other *QCalendarPermission) {
	QCalendarPermission_OperatorAssign(this.h, other.cPointer())
}

func (this *QCalendarPermission) OnOperatorAssign(slot func(other *QCalendarPermission)) {
	QCalendarPermission_connect_OperatorAssign(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarPermission_OperatorAssign
func miqt_exec_callback_QCalendarPermission_OperatorAssign(cb intptr_t, other *QCalendarPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QCalendarPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCalendarPermission(other)

	gofunc(slotval1)
}

func (this *QCalendarPermission) Swap(other *QCalendarPermission) {
	QCalendarPermission_Swap(this.h, other.cPointer())
}

func (this *QCalendarPermission) OnSwap(slot func(other *QCalendarPermission)) {
	QCalendarPermission_connect_Swap(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarPermission_Swap
func miqt_exec_callback_QCalendarPermission_Swap(cb intptr_t, other *QCalendarPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QCalendarPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCalendarPermission(other)

	gofunc(slotval1)
}

type QContactsPermission struct {
	h          uintptr
	isSubclass bool
}

// NewQContactsPermission constructs a new QContactsPermission object.
func NewQContactsPermission() *QContactsPermission {
	g := newQContactsPermission(QContactsPermission_new())
	g.isSubclass = true
	return g
}

// NewQContactsPermission2 constructs a new QContactsPermission object.
func NewQContactsPermission2(other *QContactsPermission) *QContactsPermission {
	g := newQContactsPermission(QContactsPermission_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QContactsPermission) SetAccessMode(mode AccessMode) {
	QContactsPermission_SetAccessMode(this.h, mode)
}

func (this *QContactsPermission) AccessMode() AccessMode {
	xxxxxxxxx
}

func (this *QContactsPermission) OperatorAssign(other *QContactsPermission) {
	QContactsPermission_OperatorAssign(this.h, other.cPointer())
}

func (this *QContactsPermission) OnOperatorAssign(slot func(other *QContactsPermission)) {
	QContactsPermission_connect_OperatorAssign(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QContactsPermission_OperatorAssign
func miqt_exec_callback_QContactsPermission_OperatorAssign(cb intptr_t, other *QContactsPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QContactsPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContactsPermission(other)

	gofunc(slotval1)
}

func (this *QContactsPermission) Swap(other *QContactsPermission) {
	QContactsPermission_Swap(this.h, other.cPointer())
}

func (this *QContactsPermission) OnSwap(slot func(other *QContactsPermission)) {
	QContactsPermission_connect_Swap(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QContactsPermission_Swap
func miqt_exec_callback_QContactsPermission_Swap(cb intptr_t, other *QContactsPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QContactsPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContactsPermission(other)

	gofunc(slotval1)
}

type QBluetoothPermission struct {
	h          uintptr
	isSubclass bool
}

// NewQBluetoothPermission constructs a new QBluetoothPermission object.
func NewQBluetoothPermission() *QBluetoothPermission {
	g := newQBluetoothPermission(QBluetoothPermission_new())
	g.isSubclass = true
	return g
}

// NewQBluetoothPermission2 constructs a new QBluetoothPermission object.
func NewQBluetoothPermission2(other *QBluetoothPermission) *QBluetoothPermission {
	g := newQBluetoothPermission(QBluetoothPermission_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QBluetoothPermission) SetCommunicationModes(modes CommunicationModes) {
	QBluetoothPermission_SetCommunicationModes(this.h, modes)
}

func (this *QBluetoothPermission) CommunicationModes() CommunicationModes {
	xxxxxxxxx
}

func (this *QBluetoothPermission) OperatorAssign(other *QBluetoothPermission) {
	QBluetoothPermission_OperatorAssign(this.h, other.cPointer())
}

func (this *QBluetoothPermission) OnOperatorAssign(slot func(other *QBluetoothPermission)) {
	QBluetoothPermission_connect_OperatorAssign(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBluetoothPermission_OperatorAssign
func miqt_exec_callback_QBluetoothPermission_OperatorAssign(cb intptr_t, other *QBluetoothPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QBluetoothPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQBluetoothPermission(other)

	gofunc(slotval1)
}

func (this *QBluetoothPermission) Swap(other *QBluetoothPermission) {
	QBluetoothPermission_Swap(this.h, other.cPointer())
}

func (this *QBluetoothPermission) OnSwap(slot func(other *QBluetoothPermission)) {
	QBluetoothPermission_connect_Swap(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBluetoothPermission_Swap
func miqt_exec_callback_QBluetoothPermission_Swap(cb intptr_t, other *QBluetoothPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QBluetoothPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQBluetoothPermission(other)

	gofunc(slotval1)
}

type QCameraPermission struct {
	h          uintptr
	isSubclass bool
}

// NewQCameraPermission constructs a new QCameraPermission object.
func NewQCameraPermission() *QCameraPermission {
	g := newQCameraPermission(QCameraPermission_new())
	g.isSubclass = true
	return g
}

// NewQCameraPermission2 constructs a new QCameraPermission object.
func NewQCameraPermission2(other *QCameraPermission) *QCameraPermission {
	g := newQCameraPermission(QCameraPermission_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QCameraPermission) OperatorAssign(other *QCameraPermission) {
	QCameraPermission_OperatorAssign(this.h, other.cPointer())
}

func (this *QCameraPermission) OnOperatorAssign(slot func(other *QCameraPermission)) {
	QCameraPermission_connect_OperatorAssign(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCameraPermission_OperatorAssign
func miqt_exec_callback_QCameraPermission_OperatorAssign(cb intptr_t, other *QCameraPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QCameraPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCameraPermission(other)

	gofunc(slotval1)
}

func (this *QCameraPermission) Swap(other *QCameraPermission) {
	QCameraPermission_Swap(this.h, other.cPointer())
}

func (this *QCameraPermission) OnSwap(slot func(other *QCameraPermission)) {
	QCameraPermission_connect_Swap(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCameraPermission_Swap
func miqt_exec_callback_QCameraPermission_Swap(cb intptr_t, other *QCameraPermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QCameraPermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCameraPermission(other)

	gofunc(slotval1)
}

type QMicrophonePermission struct {
	h          uintptr
	isSubclass bool
}

// NewQMicrophonePermission constructs a new QMicrophonePermission object.
func NewQMicrophonePermission() *QMicrophonePermission {
	g := newQMicrophonePermission(QMicrophonePermission_new())
	g.isSubclass = true
	return g
}

// NewQMicrophonePermission2 constructs a new QMicrophonePermission object.
func NewQMicrophonePermission2(other *QMicrophonePermission) *QMicrophonePermission {
	g := newQMicrophonePermission(QMicrophonePermission_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QMicrophonePermission) OperatorAssign(other *QMicrophonePermission) {
	QMicrophonePermission_OperatorAssign(this.h, other.cPointer())
}

func (this *QMicrophonePermission) OnOperatorAssign(slot func(other *QMicrophonePermission)) {
	QMicrophonePermission_connect_OperatorAssign(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMicrophonePermission_OperatorAssign
func miqt_exec_callback_QMicrophonePermission_OperatorAssign(cb intptr_t, other *QMicrophonePermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QMicrophonePermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMicrophonePermission(other)

	gofunc(slotval1)
}

func (this *QMicrophonePermission) Swap(other *QMicrophonePermission) {
	QMicrophonePermission_Swap(this.h, other.cPointer())
}

func (this *QMicrophonePermission) OnSwap(slot func(other *QMicrophonePermission)) {
	QMicrophonePermission_connect_Swap(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMicrophonePermission_Swap
func miqt_exec_callback_QMicrophonePermission_Swap(cb intptr_t, other *QMicrophonePermission) {
	gofunc, ok := cgo.Handle(cb).Value().(func(other *QMicrophonePermission))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMicrophonePermission(other)

	gofunc(slotval1)
}
