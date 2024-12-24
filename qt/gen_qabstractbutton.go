package qt

import (
	"unsafe"
)

type QAbstractButton struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractButton constructs a new QAbstractButton object.
func NewQAbstractButton(parent *QWidget) *QAbstractButton {
	ret := newQAbstractButton(QAbstractButton_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAbstractButton2 constructs a new QAbstractButton object.
func NewQAbstractButton2() *QAbstractButton {
	ret := newQAbstractButton(QAbstractButton_new2())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractButton) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractButton_MetaObject(this.h))
}

func (this *QAbstractButton) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractButton_Metacast(this.h, param1_Cstring))
}

func QAbstractButton_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractButton_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractButton) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QAbstractButton_SetText(this.h, text_ms)
}

func (this *QAbstractButton) Text() string {
	var _ms struct_miqt_string = QAbstractButton_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractButton) SetIcon(icon *QIcon) {
	QAbstractButton_SetIcon(this.h, icon.cPointer())
}

func (this *QAbstractButton) Icon() *QIcon {
	_goptr := newQIcon(QAbstractButton_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractButton) IconSize() *QSize {
	_goptr := newQSize(QAbstractButton_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractButton) SetShortcut(key *QKeySequence) {
	QAbstractButton_SetShortcut(this.h, key.cPointer())
}

func (this *QAbstractButton) Shortcut() *QKeySequence {
	_goptr := newQKeySequence(QAbstractButton_Shortcut(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractButton) SetCheckable(checkable bool) {
	QAbstractButton_SetCheckable(this.h, (bool)(checkable))
}

func (this *QAbstractButton) IsCheckable() bool {
	return (bool)(QAbstractButton_IsCheckable(this.h))
}

func (this *QAbstractButton) IsChecked() bool {
	return (bool)(QAbstractButton_IsChecked(this.h))
}

func (this *QAbstractButton) SetDown(down bool) {
	QAbstractButton_SetDown(this.h, (bool)(down))
}

func (this *QAbstractButton) IsDown() bool {
	return (bool)(QAbstractButton_IsDown(this.h))
}

func (this *QAbstractButton) SetAutoRepeat(autoRepeat bool) {
	QAbstractButton_SetAutoRepeat(this.h, (bool)(autoRepeat))
}

func (this *QAbstractButton) AutoRepeat() bool {
	return (bool)(QAbstractButton_AutoRepeat(this.h))
}

func (this *QAbstractButton) SetAutoRepeatDelay(autoRepeatDelay int) {
	QAbstractButton_SetAutoRepeatDelay(this.h, (int)(autoRepeatDelay))
}

func (this *QAbstractButton) AutoRepeatDelay() int {
	return (int)(QAbstractButton_AutoRepeatDelay(this.h))
}

func (this *QAbstractButton) SetAutoRepeatInterval(autoRepeatInterval int) {
	QAbstractButton_SetAutoRepeatInterval(this.h, (int)(autoRepeatInterval))
}

func (this *QAbstractButton) AutoRepeatInterval() int {
	return (int)(QAbstractButton_AutoRepeatInterval(this.h))
}

func (this *QAbstractButton) SetAutoExclusive(autoExclusive bool) {
	QAbstractButton_SetAutoExclusive(this.h, (bool)(autoExclusive))
}

func (this *QAbstractButton) AutoExclusive() bool {
	return (bool)(QAbstractButton_AutoExclusive(this.h))
}

func (this *QAbstractButton) Group() *QButtonGroup {
	return newQButtonGroup(QAbstractButton_Group(this.h))
}

func (this *QAbstractButton) SetIconSize(size *QSize) {
	QAbstractButton_SetIconSize(this.h, size.cPointer())
}

func (this *QAbstractButton) AnimateClick() {
	QAbstractButton_AnimateClick(this.h)
}

func (this *QAbstractButton) Click() {
	QAbstractButton_Click(this.h)
}

func (this *QAbstractButton) Toggle() {
	QAbstractButton_Toggle(this.h)
}

func (this *QAbstractButton) SetChecked(checked bool) {
	QAbstractButton_SetChecked(this.h, (bool)(checked))
}

func (this *QAbstractButton) Pressed() {
	QAbstractButton_Pressed(this.h)
}

func (this *QAbstractButton) OnPressed(slot func()) {
	QAbstractButton_connect_Pressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Pressed
func miqt_exec_callback_QAbstractButton_Pressed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractButton) Released() {
	QAbstractButton_Released(this.h)
}

func (this *QAbstractButton) OnReleased(slot func()) {
	QAbstractButton_connect_Released(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Released
func miqt_exec_callback_QAbstractButton_Released(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractButton) Clicked() {
	QAbstractButton_Clicked(this.h)
}

func (this *QAbstractButton) OnClicked(slot func()) {
	QAbstractButton_connect_Clicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Clicked
func miqt_exec_callback_QAbstractButton_Clicked(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractButton) Toggled(checked bool) {
	QAbstractButton_Toggled(this.h, (bool)(checked))
}

func (this *QAbstractButton) OnToggled(slot func(checked bool)) {
	QAbstractButton_connect_Toggled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Toggled
func miqt_exec_callback_QAbstractButton_Toggled(cb intptr_t, checked bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checked bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checked)

	gofunc(slotval1)
}

func QAbstractButton_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractButton_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractButton_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractButton_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractButton) Clicked1(checked bool) {
	QAbstractButton_Clicked1(this.h, (bool)(checked))
}

func (this *QAbstractButton) OnClicked1(slot func(checked bool)) {
	QAbstractButton_connect_Clicked1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Clicked1
func miqt_exec_callback_QAbstractButton_Clicked1(cb intptr_t, checked bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checked bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checked)

	gofunc(slotval1)
}

func (this *QAbstractButton) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractButton_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractButton) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_MetaObject
func miqt_exec_callback_QAbstractButton_MetaObject(self QAbstractButton, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractButton) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractButton_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractButton) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractButton_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractButton_Metacast
func miqt_exec_callback_QAbstractButton_Metacast(self QAbstractButton, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractButton{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
