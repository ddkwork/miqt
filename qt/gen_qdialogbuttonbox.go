package qt

import (
	"unsafe"
)

type QDialogButtonBox__ButtonRole int

const (
	QDialogButtonBox__InvalidRole     QDialogButtonBox__ButtonRole = -1
	QDialogButtonBox__AcceptRole      QDialogButtonBox__ButtonRole = 0
	QDialogButtonBox__RejectRole      QDialogButtonBox__ButtonRole = 1
	QDialogButtonBox__DestructiveRole QDialogButtonBox__ButtonRole = 2
	QDialogButtonBox__ActionRole      QDialogButtonBox__ButtonRole = 3
	QDialogButtonBox__HelpRole        QDialogButtonBox__ButtonRole = 4
	QDialogButtonBox__YesRole         QDialogButtonBox__ButtonRole = 5
	QDialogButtonBox__NoRole          QDialogButtonBox__ButtonRole = 6
	QDialogButtonBox__ResetRole       QDialogButtonBox__ButtonRole = 7
	QDialogButtonBox__ApplyRole       QDialogButtonBox__ButtonRole = 8
	QDialogButtonBox__NRoles          QDialogButtonBox__ButtonRole = 9
)

type QDialogButtonBox__StandardButton int

const (
	QDialogButtonBox__NoButton        QDialogButtonBox__StandardButton = 0
	QDialogButtonBox__Ok              QDialogButtonBox__StandardButton = 1024
	QDialogButtonBox__Save            QDialogButtonBox__StandardButton = 2048
	QDialogButtonBox__SaveAll         QDialogButtonBox__StandardButton = 4096
	QDialogButtonBox__Open            QDialogButtonBox__StandardButton = 8192
	QDialogButtonBox__Yes             QDialogButtonBox__StandardButton = 16384
	QDialogButtonBox__YesToAll        QDialogButtonBox__StandardButton = 32768
	QDialogButtonBox__No              QDialogButtonBox__StandardButton = 65536
	QDialogButtonBox__NoToAll         QDialogButtonBox__StandardButton = 131072
	QDialogButtonBox__Abort           QDialogButtonBox__StandardButton = 262144
	QDialogButtonBox__Retry           QDialogButtonBox__StandardButton = 524288
	QDialogButtonBox__Ignore          QDialogButtonBox__StandardButton = 1048576
	QDialogButtonBox__Close           QDialogButtonBox__StandardButton = 2097152
	QDialogButtonBox__Cancel          QDialogButtonBox__StandardButton = 4194304
	QDialogButtonBox__Discard         QDialogButtonBox__StandardButton = 8388608
	QDialogButtonBox__Help            QDialogButtonBox__StandardButton = 16777216
	QDialogButtonBox__Apply           QDialogButtonBox__StandardButton = 33554432
	QDialogButtonBox__Reset           QDialogButtonBox__StandardButton = 67108864
	QDialogButtonBox__RestoreDefaults QDialogButtonBox__StandardButton = 134217728
	QDialogButtonBox__FirstButton     QDialogButtonBox__StandardButton = 1024
	QDialogButtonBox__LastButton      QDialogButtonBox__StandardButton = 134217728
)

type QDialogButtonBox__ButtonLayout int

const (
	QDialogButtonBox__WinLayout     QDialogButtonBox__ButtonLayout = 0
	QDialogButtonBox__MacLayout     QDialogButtonBox__ButtonLayout = 1
	QDialogButtonBox__KdeLayout     QDialogButtonBox__ButtonLayout = 2
	QDialogButtonBox__GnomeLayout   QDialogButtonBox__ButtonLayout = 3
	QDialogButtonBox__AndroidLayout QDialogButtonBox__ButtonLayout = 4
)

type QDialogButtonBox struct {
	h          uintptr
	isSubclass bool
}

// NewQDialogButtonBox constructs a new QDialogButtonBox object.
func NewQDialogButtonBox(parent *QWidget) *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQDialogButtonBox2 constructs a new QDialogButtonBox object.
func NewQDialogButtonBox2() *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new2())
	g.isSubclass = true
	return g
}

// NewQDialogButtonBox3 constructs a new QDialogButtonBox object.
func NewQDialogButtonBox3(orientation Orientation) *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new3((int)(orientation)))
	g.isSubclass = true
	return g
}

// NewQDialogButtonBox4 constructs a new QDialogButtonBox object.
func NewQDialogButtonBox4(buttons StandardButtons) *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new4(buttons))
	g.isSubclass = true
	return g
}

// NewQDialogButtonBox5 constructs a new QDialogButtonBox object.
func NewQDialogButtonBox5(buttons StandardButtons, orientation Orientation) *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new5(buttons, (int)(orientation)))
	g.isSubclass = true
	return g
}

// NewQDialogButtonBox6 constructs a new QDialogButtonBox object.
func NewQDialogButtonBox6(orientation Orientation, parent *QWidget) *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new6((int)(orientation), parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQDialogButtonBox7 constructs a new QDialogButtonBox object.
func NewQDialogButtonBox7(buttons StandardButtons, parent *QWidget) *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new7(buttons, parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQDialogButtonBox8 constructs a new QDialogButtonBox object.
func NewQDialogButtonBox8(buttons StandardButtons, orientation Orientation, parent *QWidget) *QDialogButtonBox {
	g := newQDialogButtonBox(QDialogButtonBox_new8(buttons, (int)(orientation), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QDialogButtonBox) MetaObject() *QMetaObject {
	return newQMetaObject(QDialogButtonBox_MetaObject(this.h))
}

func (this *QDialogButtonBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDialogButtonBox_Metacast(this.h, param1_Cstring))
}

func QDialogButtonBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDialogButtonBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDialogButtonBox) SetOrientation(orientation Orientation) {
	QDialogButtonBox_SetOrientation(this.h, (int)(orientation))
}

func (this *QDialogButtonBox) Orientation() Orientation {
	return (Orientation)(QDialogButtonBox_Orientation(this.h))
}

func (this *QDialogButtonBox) AddButton(button *QAbstractButton, role ButtonRole) {
	QDialogButtonBox_AddButton(this.h, button.cPointer(), role)
}

func (this *QDialogButtonBox) AddButton2(text string, role ButtonRole) *QPushButton {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQPushButton(QDialogButtonBox_AddButton2(this.h, text_ms, role))
}

func (this *QDialogButtonBox) AddButtonWithButton(button StandardButton) *QPushButton {
	return newQPushButton(QDialogButtonBox_AddButtonWithButton(this.h, button))
}

func (this *QDialogButtonBox) RemoveButton(button *QAbstractButton) {
	QDialogButtonBox_RemoveButton(this.h, button.cPointer())
}

func (this *QDialogButtonBox) Clear() {
	QDialogButtonBox_Clear(this.h)
}

func (this *QDialogButtonBox) Buttons() []*QAbstractButton {
	var _ma struct_miqt_array = QDialogButtonBox_Buttons(this.h)
	_ret := make([]*QAbstractButton, int(_ma.len))
	_outCast := (*[0xffff]*QAbstractButton)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAbstractButton(_outCast[i])
	}
	return _ret
}

func (this *QDialogButtonBox) ButtonRole(button *QAbstractButton) ButtonRole {
	xxxxxxxxx
}

func (this *QDialogButtonBox) SetStandardButtons(buttons StandardButtons) {
	QDialogButtonBox_SetStandardButtons(this.h, buttons)
}

func (this *QDialogButtonBox) StandardButtons() StandardButtons {
	xxxxxxxxx
}

func (this *QDialogButtonBox) StandardButton(button *QAbstractButton) StandardButton {
	xxxxxxxxx
}

func (this *QDialogButtonBox) Button(which StandardButton) *QPushButton {
	return newQPushButton(QDialogButtonBox_Button(this.h, which))
}

func (this *QDialogButtonBox) SetCenterButtons(center bool) {
	QDialogButtonBox_SetCenterButtons(this.h, (bool)(center))
}

func (this *QDialogButtonBox) CenterButtons() bool {
	return (bool)(QDialogButtonBox_CenterButtons(this.h))
}

func (this *QDialogButtonBox) Clicked(button *QAbstractButton) {
	QDialogButtonBox_Clicked(this.h, button.cPointer())
}

func (this *QDialogButtonBox) OnClicked(slot func(button *QAbstractButton)) {
	QDialogButtonBox_connect_Clicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialogButtonBox_Clicked
func miqt_exec_callback_QDialogButtonBox_Clicked(cb intptr_t, button *QAbstractButton) {
	gofunc, ok := cgo.Handle(cb).Value().(func(button *QAbstractButton))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractButton(button)

	gofunc(slotval1)
}

func (this *QDialogButtonBox) Accepted() {
	QDialogButtonBox_Accepted(this.h)
}

func (this *QDialogButtonBox) OnAccepted(slot func()) {
	QDialogButtonBox_connect_Accepted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialogButtonBox_Accepted
func miqt_exec_callback_QDialogButtonBox_Accepted(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QDialogButtonBox) HelpRequested() {
	QDialogButtonBox_HelpRequested(this.h)
}

func (this *QDialogButtonBox) OnHelpRequested(slot func()) {
	QDialogButtonBox_connect_HelpRequested(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialogButtonBox_HelpRequested
func miqt_exec_callback_QDialogButtonBox_HelpRequested(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QDialogButtonBox) Rejected() {
	QDialogButtonBox_Rejected(this.h)
}

func (this *QDialogButtonBox) OnRejected(slot func()) {
	QDialogButtonBox_connect_Rejected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialogButtonBox_Rejected
func miqt_exec_callback_QDialogButtonBox_Rejected(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QDialogButtonBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDialogButtonBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDialogButtonBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDialogButtonBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDialogButtonBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDialogButtonBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDialogButtonBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialogButtonBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialogButtonBox_MetaObject
func miqt_exec_callback_QDialogButtonBox_MetaObject(self QDialogButtonBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialogButtonBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDialogButtonBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDialogButtonBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDialogButtonBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialogButtonBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialogButtonBox_Metacast
func miqt_exec_callback_QDialogButtonBox_Metacast(self QDialogButtonBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QDialogButtonBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
