package qt

import (
	"unsafe"
)

type QActionGroup__ExclusionPolicy int

const (
	QActionGroup__None              QActionGroup__ExclusionPolicy = 0
	QActionGroup__Exclusive         QActionGroup__ExclusionPolicy = 1
	QActionGroup__ExclusiveOptional QActionGroup__ExclusionPolicy = 2
)

type QActionGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQActionGroup constructs a new QActionGroup object.
func NewQActionGroup(parent *QObject) *QActionGroup {
	ret := newQActionGroup(QActionGroup_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QActionGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QActionGroup_MetaObject(this.h))
}

func (this *QActionGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QActionGroup_Metacast(this.h, param1_Cstring))
}

func QActionGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QActionGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QActionGroup) AddAction(a *QAction) *QAction {
	return newQAction(QActionGroup_AddAction(this.h, a.cPointer()))
}

func (this *QActionGroup) AddActionWithText(text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QActionGroup_AddActionWithText(this.h, text_ms))
}

func (this *QActionGroup) AddAction2(icon *QIcon, text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QActionGroup_AddAction2(this.h, icon.cPointer(), text_ms))
}

func (this *QActionGroup) RemoveAction(a *QAction) {
	QActionGroup_RemoveAction(this.h, a.cPointer())
}

func (this *QActionGroup) Actions() []*QAction {
	var _ma struct_miqt_array = QActionGroup_Actions(this.h)
	_ret := make([]*QAction, int(_ma.len))
	_outCast := (*[0xffff]*QAction)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAction(_outCast[i])
	}
	return _ret
}

func (this *QActionGroup) CheckedAction() *QAction {
	return newQAction(QActionGroup_CheckedAction(this.h))
}

func (this *QActionGroup) IsExclusive() bool {
	return (bool)(QActionGroup_IsExclusive(this.h))
}

func (this *QActionGroup) IsEnabled() bool {
	return (bool)(QActionGroup_IsEnabled(this.h))
}

func (this *QActionGroup) IsVisible() bool {
	return (bool)(QActionGroup_IsVisible(this.h))
}

func (this *QActionGroup) ExclusionPolicy() ExclusionPolicy {
	xxxxxxxxx
}

func (this *QActionGroup) SetEnabled(enabled bool) {
	QActionGroup_SetEnabled(this.h, (bool)(enabled))
}

func (this *QActionGroup) SetDisabled(b bool) {
	QActionGroup_SetDisabled(this.h, (bool)(b))
}

func (this *QActionGroup) SetVisible(visible bool) {
	QActionGroup_SetVisible(this.h, (bool)(visible))
}

func (this *QActionGroup) SetExclusive(exclusive bool) {
	QActionGroup_SetExclusive(this.h, (bool)(exclusive))
}

func (this *QActionGroup) SetExclusionPolicy(policy ExclusionPolicy) {
	QActionGroup_SetExclusionPolicy(this.h, policy)
}

func (this *QActionGroup) Triggered(param1 *QAction) {
	QActionGroup_Triggered(this.h, param1.cPointer())
}

func (this *QActionGroup) OnTriggered(slot func(param1 *QAction)) {
	QActionGroup_connect_Triggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_Triggered
func miqt_exec_callback_QActionGroup_Triggered(cb intptr_t, param1 *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(param1)

	gofunc(slotval1)
}

func (this *QActionGroup) Hovered(param1 *QAction) {
	QActionGroup_Hovered(this.h, param1.cPointer())
}

func (this *QActionGroup) OnHovered(slot func(param1 *QAction)) {
	QActionGroup_connect_Hovered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_Hovered
func miqt_exec_callback_QActionGroup_Hovered(cb intptr_t, param1 *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(param1)

	gofunc(slotval1)
}

func QActionGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QActionGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QActionGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QActionGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QActionGroup) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QActionGroup_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QActionGroup) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_MetaObject
func miqt_exec_callback_QActionGroup_MetaObject(self QActionGroup, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QActionGroup{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QActionGroup) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QActionGroup_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QActionGroup) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QActionGroup_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QActionGroup_Metacast
func miqt_exec_callback_QActionGroup_Metacast(self QActionGroup, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QActionGroup{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
