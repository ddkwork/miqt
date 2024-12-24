package qt

import (
	"unsafe"
)

type QButtonGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQButtonGroup constructs a new QButtonGroup object.
func NewQButtonGroup() *QButtonGroup {
	g := newQButtonGroup(QButtonGroup_new())
	g.isSubclass = true
	return g
}

// NewQButtonGroup2 constructs a new QButtonGroup object.
func NewQButtonGroup2(parent *QObject) *QButtonGroup {
	g := newQButtonGroup(QButtonGroup_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QButtonGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QButtonGroup_MetaObject(this.h))
}

func (this *QButtonGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QButtonGroup_Metacast(this.h, param1_Cstring))
}

func QButtonGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QButtonGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QButtonGroup) SetExclusive(exclusive bool) {
	QButtonGroup_SetExclusive(this.h, (bool)(exclusive))
}

func (this *QButtonGroup) Exclusive() bool {
	return (bool)(QButtonGroup_Exclusive(this.h))
}

func (this *QButtonGroup) AddButton(param1 *QAbstractButton) {
	QButtonGroup_AddButton(this.h, param1.cPointer())
}

func (this *QButtonGroup) RemoveButton(param1 *QAbstractButton) {
	QButtonGroup_RemoveButton(this.h, param1.cPointer())
}

func (this *QButtonGroup) Buttons() []*QAbstractButton {
	var _ma struct_miqt_array = QButtonGroup_Buttons(this.h)
	_ret := make([]*QAbstractButton, int(_ma.len))
	_outCast := (*[0xffff]*QAbstractButton)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAbstractButton(_outCast[i])
	}
	return _ret
}

func (this *QButtonGroup) CheckedButton() *QAbstractButton {
	return newQAbstractButton(QButtonGroup_CheckedButton(this.h))
}

func (this *QButtonGroup) Button(id int) *QAbstractButton {
	return newQAbstractButton(QButtonGroup_Button(this.h, (int)(id)))
}

func (this *QButtonGroup) SetId(button *QAbstractButton, id int) {
	QButtonGroup_SetId(this.h, button.cPointer(), (int)(id))
}

func (this *QButtonGroup) Id(button *QAbstractButton) int {
	return (int)(QButtonGroup_Id(this.h, button.cPointer()))
}

func (this *QButtonGroup) CheckedId() int {
	return (int)(QButtonGroup_CheckedId(this.h))
}

func (this *QButtonGroup) ButtonClicked(param1 *QAbstractButton) {
	QButtonGroup_ButtonClicked(this.h, param1.cPointer())
}

func (this *QButtonGroup) OnButtonClicked(slot func(param1 *QAbstractButton)) {
	QButtonGroup_connect_ButtonClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_ButtonClicked
func miqt_exec_callback_QButtonGroup_ButtonClicked(cb intptr_t, param1 *QAbstractButton) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAbstractButton))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractButton(param1)

	gofunc(slotval1)
}

func (this *QButtonGroup) ButtonPressed(param1 *QAbstractButton) {
	QButtonGroup_ButtonPressed(this.h, param1.cPointer())
}

func (this *QButtonGroup) OnButtonPressed(slot func(param1 *QAbstractButton)) {
	QButtonGroup_connect_ButtonPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_ButtonPressed
func miqt_exec_callback_QButtonGroup_ButtonPressed(cb intptr_t, param1 *QAbstractButton) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAbstractButton))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractButton(param1)

	gofunc(slotval1)
}

func (this *QButtonGroup) ButtonReleased(param1 *QAbstractButton) {
	QButtonGroup_ButtonReleased(this.h, param1.cPointer())
}

func (this *QButtonGroup) OnButtonReleased(slot func(param1 *QAbstractButton)) {
	QButtonGroup_connect_ButtonReleased(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_ButtonReleased
func miqt_exec_callback_QButtonGroup_ButtonReleased(cb intptr_t, param1 *QAbstractButton) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAbstractButton))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractButton(param1)

	gofunc(slotval1)
}

func (this *QButtonGroup) ButtonToggled(param1 *QAbstractButton, param2 bool) {
	QButtonGroup_ButtonToggled(this.h, param1.cPointer(), (bool)(param2))
}

func (this *QButtonGroup) OnButtonToggled(slot func(param1 *QAbstractButton, param2 bool)) {
	QButtonGroup_connect_ButtonToggled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_ButtonToggled
func miqt_exec_callback_QButtonGroup_ButtonToggled(cb intptr_t, param1 *QAbstractButton, param2 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAbstractButton, param2 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractButton(param1)

	slotval2 := (bool)(param2)

	gofunc(slotval1, slotval2)
}

func (this *QButtonGroup) IdClicked(param1 int) {
	QButtonGroup_IdClicked(this.h, (int)(param1))
}

func (this *QButtonGroup) OnIdClicked(slot func(param1 int)) {
	QButtonGroup_connect_IdClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_IdClicked
func miqt_exec_callback_QButtonGroup_IdClicked(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QButtonGroup) IdPressed(param1 int) {
	QButtonGroup_IdPressed(this.h, (int)(param1))
}

func (this *QButtonGroup) OnIdPressed(slot func(param1 int)) {
	QButtonGroup_connect_IdPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_IdPressed
func miqt_exec_callback_QButtonGroup_IdPressed(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QButtonGroup) IdReleased(param1 int) {
	QButtonGroup_IdReleased(this.h, (int)(param1))
}

func (this *QButtonGroup) OnIdReleased(slot func(param1 int)) {
	QButtonGroup_connect_IdReleased(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_IdReleased
func miqt_exec_callback_QButtonGroup_IdReleased(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QButtonGroup) IdToggled(param1 int, param2 bool) {
	QButtonGroup_IdToggled(this.h, (int)(param1), (bool)(param2))
}

func (this *QButtonGroup) OnIdToggled(slot func(param1 int, param2 bool)) {
	QButtonGroup_connect_IdToggled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_IdToggled
func miqt_exec_callback_QButtonGroup_IdToggled(cb intptr_t, param1 int, param2 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int, param2 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	slotval2 := (bool)(param2)

	gofunc(slotval1, slotval2)
}

func QButtonGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QButtonGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QButtonGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QButtonGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QButtonGroup) AddButton2(param1 *QAbstractButton, id int) {
	QButtonGroup_AddButton2(this.h, param1.cPointer(), (int)(id))
}

func (this *QButtonGroup) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QButtonGroup_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QButtonGroup) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QButtonGroup_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_MetaObject
func miqt_exec_callback_QButtonGroup_MetaObject(self QButtonGroup, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QButtonGroup{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QButtonGroup) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QButtonGroup_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QButtonGroup) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QButtonGroup_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QButtonGroup_Metacast
func miqt_exec_callback_QButtonGroup_Metacast(self QButtonGroup, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QButtonGroup{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
