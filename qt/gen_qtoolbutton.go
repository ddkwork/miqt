package qt

import (
	"unsafe"
)

type QToolButton__ToolButtonPopupMode int

const (
	QToolButton__DelayedPopup    QToolButton__ToolButtonPopupMode = 0
	QToolButton__MenuButtonPopup QToolButton__ToolButtonPopupMode = 1
	QToolButton__InstantPopup    QToolButton__ToolButtonPopupMode = 2
)

type QToolButton struct {
	h          uintptr
	isSubclass bool
}

// NewQToolButton constructs a new QToolButton object.
func NewQToolButton(parent *QWidget) *QToolButton {
	ret := newQToolButton(QToolButton_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQToolButton2 constructs a new QToolButton object.
func NewQToolButton2() *QToolButton {
	ret := newQToolButton(QToolButton_new2())
	ret.isSubclass = true
	return ret
}

func (this *QToolButton) MetaObject() *QMetaObject {
	return newQMetaObject(QToolButton_MetaObject(this.h))
}

func (this *QToolButton) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QToolButton_Metacast(this.h, param1_Cstring))
}

func QToolButton_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QToolButton_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolButton) SizeHint() *QSize {
	_goptr := newQSize(QToolButton_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QToolButton) MinimumSizeHint() *QSize {
	_goptr := newQSize(QToolButton_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QToolButton) ToolButtonStyle() ToolButtonStyle {
	return (ToolButtonStyle)(QToolButton_ToolButtonStyle(this.h))
}

func (this *QToolButton) ArrowType() ArrowType {
	return (ArrowType)(QToolButton_ArrowType(this.h))
}

func (this *QToolButton) SetArrowType(typeVal ArrowType) {
	QToolButton_SetArrowType(this.h, (int)(typeVal))
}

func (this *QToolButton) SetMenu(menu *QMenu) {
	QToolButton_SetMenu(this.h, menu.cPointer())
}

func (this *QToolButton) Menu() *QMenu {
	return newQMenu(QToolButton_Menu(this.h))
}

func (this *QToolButton) SetPopupMode(mode ToolButtonPopupMode) {
	QToolButton_SetPopupMode(this.h, mode)
}

func (this *QToolButton) PopupMode() ToolButtonPopupMode {
	xxxxxxxxx
}

func (this *QToolButton) DefaultAction() *QAction {
	return newQAction(QToolButton_DefaultAction(this.h))
}

func (this *QToolButton) SetAutoRaise(enable bool) {
	QToolButton_SetAutoRaise(this.h, (bool)(enable))
}

func (this *QToolButton) AutoRaise() bool {
	return (bool)(QToolButton_AutoRaise(this.h))
}

func (this *QToolButton) ShowMenu() {
	QToolButton_ShowMenu(this.h)
}

func (this *QToolButton) SetToolButtonStyle(style ToolButtonStyle) {
	QToolButton_SetToolButtonStyle(this.h, (int)(style))
}

func (this *QToolButton) SetDefaultAction(defaultAction *QAction) {
	QToolButton_SetDefaultAction(this.h, defaultAction.cPointer())
}

func (this *QToolButton) Triggered(param1 *QAction) {
	QToolButton_Triggered(this.h, param1.cPointer())
}

func (this *QToolButton) OnTriggered(slot func(param1 *QAction)) {
	QToolButton_connect_Triggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolButton_Triggered
func miqt_exec_callback_QToolButton_Triggered(cb intptr_t, param1 *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(param1)

	gofunc(slotval1)
}

func QToolButton_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QToolButton_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QToolButton_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QToolButton_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolButton) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QToolButton_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QToolButton) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolButton_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolButton_MetaObject
func miqt_exec_callback_QToolButton_MetaObject(self QToolButton, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolButton{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QToolButton) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QToolButton_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QToolButton) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolButton_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolButton_Metacast
func miqt_exec_callback_QToolButton_Metacast(self QToolButton, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QToolButton{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
