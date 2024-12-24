package qt

import (
	"unsafe"
)

type QErrorMessage struct {
	h          uintptr
	isSubclass bool
}

// NewQErrorMessage constructs a new QErrorMessage object.
func NewQErrorMessage(parent *QWidget) *QErrorMessage {

	ret := newQErrorMessage(QErrorMessage_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQErrorMessage2 constructs a new QErrorMessage object.
func NewQErrorMessage2() *QErrorMessage {

	ret := newQErrorMessage(QErrorMessage_new2())
	ret.isSubclass = true
	return ret
}

func (this *QErrorMessage) MetaObject() *QMetaObject {
	return newQMetaObject(QErrorMessage_MetaObject(this.h))
}

func (this *QErrorMessage) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QErrorMessage_Metacast(this.h, param1_Cstring))
}

func QErrorMessage_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QErrorMessage_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QErrorMessage_QtHandler() *QErrorMessage {
	return newQErrorMessage(QErrorMessage_QtHandler())
}

func (this *QErrorMessage) ShowMessage(message string) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QErrorMessage_ShowMessage(this.h, message_ms)
}

func (this *QErrorMessage) ShowMessage2(message string, typeVal string) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	typeVal_ms := struct_miqt_string{}
	typeVal_ms.data = CString(typeVal)
	typeVal_ms.len = size_t(len(typeVal))
	defer free(unsafe.Pointer(typeVal_ms.data))
	QErrorMessage_ShowMessage2(this.h, message_ms, typeVal_ms)
}

func QErrorMessage_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QErrorMessage_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QErrorMessage_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QErrorMessage_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QErrorMessage) callVirtualBase_Done(param1 int) {

	QErrorMessage_virtualbase_Done(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QErrorMessage) OnDone(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_Done
func miqt_exec_callback_QErrorMessage_Done(self QErrorMessage, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ChangeEvent(e *QEvent) {

	QErrorMessage_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QErrorMessage) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_ChangeEvent
func miqt_exec_callback_QErrorMessage_ChangeEvent(self QErrorMessage, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_SetVisible(visible bool) {

	QErrorMessage_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QErrorMessage) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_SetVisible
func miqt_exec_callback_QErrorMessage_SetVisible(self QErrorMessage, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QErrorMessage) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QErrorMessage_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QErrorMessage) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_SizeHint
func miqt_exec_callback_QErrorMessage_SizeHint(self QErrorMessage, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QErrorMessage_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QErrorMessage) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_MinimumSizeHint
func miqt_exec_callback_QErrorMessage_MinimumSizeHint(self QErrorMessage, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_Open() {

	QErrorMessage_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QErrorMessage) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_Open
func miqt_exec_callback_QErrorMessage_Open(self QErrorMessage, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Open)

}

func (this *QErrorMessage) callVirtualBase_Exec() int {

	return (int)(QErrorMessage_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QErrorMessage) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_Exec
func miqt_exec_callback_QErrorMessage_Exec(self QErrorMessage, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_Accept() {

	QErrorMessage_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QErrorMessage) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_Accept
func miqt_exec_callback_QErrorMessage_Accept(self QErrorMessage, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Accept)

}

func (this *QErrorMessage) callVirtualBase_Reject() {

	QErrorMessage_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QErrorMessage) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_Reject
func miqt_exec_callback_QErrorMessage_Reject(self QErrorMessage, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Reject)

}

func (this *QErrorMessage) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QErrorMessage_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_KeyPressEvent
func miqt_exec_callback_QErrorMessage_KeyPressEvent(self QErrorMessage, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_CloseEvent(param1 *QCloseEvent) {

	QErrorMessage_virtualbase_CloseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnCloseEvent(slot func(super func(param1 *QCloseEvent), param1 *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_CloseEvent
func miqt_exec_callback_QErrorMessage_CloseEvent(self QErrorMessage, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QCloseEvent), param1 *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QErrorMessage_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_ShowEvent
func miqt_exec_callback_QErrorMessage_ShowEvent(self QErrorMessage, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QErrorMessage_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_ResizeEvent
func miqt_exec_callback_QErrorMessage_ResizeEvent(self QErrorMessage, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QErrorMessage_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_ContextMenuEvent
func miqt_exec_callback_QErrorMessage_ContextMenuEvent(self QErrorMessage, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QErrorMessage_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QErrorMessage) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QErrorMessage_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QErrorMessage_EventFilter
func miqt_exec_callback_QErrorMessage_EventFilter(self QErrorMessage, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
