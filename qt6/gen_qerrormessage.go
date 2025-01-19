package qt6

/*

#include "gen_qerrormessage.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QErrorMessage struct {
	h *C.QErrorMessage
	*QDialog
}

func (this *QErrorMessage) cPointer() *C.QErrorMessage {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QErrorMessage) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQErrorMessage constructs the type using only CGO pointers.
func newQErrorMessage(h *C.QErrorMessage) *QErrorMessage {
	if h == nil {
		return nil
	}
	var outptr_QDialog *C.QDialog = nil
	C.QErrorMessage_virtbase(h, &outptr_QDialog)

	return &QErrorMessage{h: h,
		QDialog: newQDialog(outptr_QDialog)}
}

// UnsafeNewQErrorMessage constructs the type using only unsafe pointers.
func UnsafeNewQErrorMessage(h unsafe.Pointer) *QErrorMessage {
	return newQErrorMessage((*C.QErrorMessage)(h))
}

// NewQErrorMessage constructs a new QErrorMessage object.
func NewQErrorMessage(parent *QWidget) *QErrorMessage {

	return newQErrorMessage(C.QErrorMessage_new(parent.cPointer()))
}

// NewQErrorMessage2 constructs a new QErrorMessage object.
func NewQErrorMessage2() *QErrorMessage {

	return newQErrorMessage(C.QErrorMessage_new2())
}

func (this *QErrorMessage) MetaObject() *QMetaObject {
	return newQMetaObject(C.QErrorMessage_MetaObject(this.h))
}

func (this *QErrorMessage) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QErrorMessage_Metacast(this.h, param1_Cstring))
}

func QErrorMessage_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QErrorMessage_Tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QErrorMessage_QtHandler() *QErrorMessage {
	return newQErrorMessage(C.QErrorMessage_QtHandler())
}

func (this *QErrorMessage) ShowMessage(message string) {
	message_ms := C.struct_miqt_string{}
	message_ms.data = C.CString(message)
	message_ms.len = C.size_t(len(message))
	defer C.free(unsafe.Pointer(message_ms.data))
	C.QErrorMessage_ShowMessage(this.h, message_ms)
}

func (this *QErrorMessage) ShowMessage2(message string, typeVal string) {
	message_ms := C.struct_miqt_string{}
	message_ms.data = C.CString(message)
	message_ms.len = C.size_t(len(message))
	defer C.free(unsafe.Pointer(message_ms.data))
	typeVal_ms := C.struct_miqt_string{}
	typeVal_ms.data = C.CString(typeVal)
	typeVal_ms.len = C.size_t(len(typeVal))
	defer C.free(unsafe.Pointer(typeVal_ms.data))
	C.QErrorMessage_ShowMessage2(this.h, message_ms, typeVal_ms)
}

func QErrorMessage_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QErrorMessage_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QErrorMessage_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QErrorMessage_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QErrorMessage) callVirtualBase_Done(param1 int) {

	C.QErrorMessage_virtualbase_Done(unsafe.Pointer(this.h), (C.int)(param1))

}
func (this *QErrorMessage) OnDone(slot func(super func(param1 int), param1 int)) {
	ok := C.QErrorMessage_override_virtual_Done(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Done
func miqt_exec_callback_QErrorMessage_Done(self *C.QErrorMessage, cb C.intptr_t, param1 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ChangeEvent(e *QEvent) {

	C.QErrorMessage_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QErrorMessage) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	ok := C.QErrorMessage_override_virtual_ChangeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_ChangeEvent
func miqt_exec_callback_QErrorMessage_ChangeEvent(self *C.QErrorMessage, cb C.intptr_t, e *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_SetVisible(visible bool) {

	C.QErrorMessage_virtualbase_SetVisible(unsafe.Pointer(this.h), (C.bool)(visible))

}
func (this *QErrorMessage) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	ok := C.QErrorMessage_override_virtual_SetVisible(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_SetVisible
func miqt_exec_callback_QErrorMessage_SetVisible(self *C.QErrorMessage, cb C.intptr_t, visible C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QErrorMessage) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(C.QErrorMessage_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QErrorMessage) OnSizeHint(slot func(super func() *QSize) *QSize) {
	ok := C.QErrorMessage_override_virtual_SizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_SizeHint
func miqt_exec_callback_QErrorMessage_SizeHint(self *C.QErrorMessage, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(C.QErrorMessage_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QErrorMessage) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	ok := C.QErrorMessage_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_MinimumSizeHint
func miqt_exec_callback_QErrorMessage_MinimumSizeHint(self *C.QErrorMessage, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_Open() {

	C.QErrorMessage_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QErrorMessage) OnOpen(slot func(super func())) {
	ok := C.QErrorMessage_override_virtual_Open(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Open
func miqt_exec_callback_QErrorMessage_Open(self *C.QErrorMessage, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Open)

}

func (this *QErrorMessage) callVirtualBase_Exec() int {

	return (int)(C.QErrorMessage_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QErrorMessage) OnExec(slot func(super func() int) int) {
	ok := C.QErrorMessage_override_virtual_Exec(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Exec
func miqt_exec_callback_QErrorMessage_Exec(self *C.QErrorMessage, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_Exec)

	return (C.int)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_Accept() {

	C.QErrorMessage_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QErrorMessage) OnAccept(slot func(super func())) {
	ok := C.QErrorMessage_override_virtual_Accept(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Accept
func miqt_exec_callback_QErrorMessage_Accept(self *C.QErrorMessage, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Accept)

}

func (this *QErrorMessage) callVirtualBase_Reject() {

	C.QErrorMessage_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QErrorMessage) OnReject(slot func(super func())) {
	ok := C.QErrorMessage_override_virtual_Reject(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Reject
func miqt_exec_callback_QErrorMessage_Reject(self *C.QErrorMessage, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QErrorMessage{h: self}).callVirtualBase_Reject)

}

func (this *QErrorMessage) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	C.QErrorMessage_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	ok := C.QErrorMessage_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_KeyPressEvent
func miqt_exec_callback_QErrorMessage_KeyPressEvent(self *C.QErrorMessage, cb C.intptr_t, param1 *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_CloseEvent(param1 *QCloseEvent) {

	C.QErrorMessage_virtualbase_CloseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnCloseEvent(slot func(super func(param1 *QCloseEvent), param1 *QCloseEvent)) {
	ok := C.QErrorMessage_override_virtual_CloseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_CloseEvent
func miqt_exec_callback_QErrorMessage_CloseEvent(self *C.QErrorMessage, cb C.intptr_t, param1 *C.QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QCloseEvent), param1 *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	C.QErrorMessage_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	ok := C.QErrorMessage_override_virtual_ShowEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_ShowEvent
func miqt_exec_callback_QErrorMessage_ShowEvent(self *C.QErrorMessage, cb C.intptr_t, param1 *C.QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	C.QErrorMessage_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	ok := C.QErrorMessage_override_virtual_ResizeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_ResizeEvent
func miqt_exec_callback_QErrorMessage_ResizeEvent(self *C.QErrorMessage, cb C.intptr_t, param1 *C.QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	C.QErrorMessage_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	ok := C.QErrorMessage_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_ContextMenuEvent
func miqt_exec_callback_QErrorMessage_ContextMenuEvent(self *C.QErrorMessage, cb C.intptr_t, param1 *C.QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(C.QErrorMessage_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QErrorMessage) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	ok := C.QErrorMessage_override_virtual_EventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_EventFilter
func miqt_exec_callback_QErrorMessage_EventFilter(self *C.QErrorMessage, cb C.intptr_t, param1 *C.QObject, param2 *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_DevType() int {

	return (int)(C.QErrorMessage_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QErrorMessage) OnDevType(slot func(super func() int) int) {
	ok := C.QErrorMessage_override_virtual_DevType(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_DevType
func miqt_exec_callback_QErrorMessage_DevType(self *C.QErrorMessage, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_DevType)

	return (C.int)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(C.QErrorMessage_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QErrorMessage) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	ok := C.QErrorMessage_override_virtual_HeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_HeightForWidth
func miqt_exec_callback_QErrorMessage_HeightForWidth(self *C.QErrorMessage, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(C.QErrorMessage_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QErrorMessage) OnHasHeightForWidth(slot func(super func() bool) bool) {
	ok := C.QErrorMessage_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_HasHeightForWidth
func miqt_exec_callback_QErrorMessage_HasHeightForWidth(self *C.QErrorMessage, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_HasHeightForWidth)

	return (C.bool)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(C.QErrorMessage_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QErrorMessage) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	ok := C.QErrorMessage_override_virtual_PaintEngine(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_PaintEngine
func miqt_exec_callback_QErrorMessage_PaintEngine(self *C.QErrorMessage, cb C.intptr_t) *C.QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(C.QErrorMessage_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QErrorMessage) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	ok := C.QErrorMessage_override_virtual_Event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Event
func miqt_exec_callback_QErrorMessage_Event(self *C.QErrorMessage, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	C.QErrorMessage_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QErrorMessage_override_virtual_MousePressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_MousePressEvent
func miqt_exec_callback_QErrorMessage_MousePressEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	C.QErrorMessage_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QErrorMessage_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_MouseReleaseEvent
func miqt_exec_callback_QErrorMessage_MouseReleaseEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	C.QErrorMessage_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QErrorMessage_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_MouseDoubleClickEvent
func miqt_exec_callback_QErrorMessage_MouseDoubleClickEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	C.QErrorMessage_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QErrorMessage_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_MouseMoveEvent
func miqt_exec_callback_QErrorMessage_MouseMoveEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_WheelEvent(event *QWheelEvent) {

	C.QErrorMessage_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	ok := C.QErrorMessage_override_virtual_WheelEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_WheelEvent
func miqt_exec_callback_QErrorMessage_WheelEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	C.QErrorMessage_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	ok := C.QErrorMessage_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_KeyReleaseEvent
func miqt_exec_callback_QErrorMessage_KeyReleaseEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	C.QErrorMessage_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	ok := C.QErrorMessage_override_virtual_FocusInEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_FocusInEvent
func miqt_exec_callback_QErrorMessage_FocusInEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	C.QErrorMessage_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	ok := C.QErrorMessage_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_FocusOutEvent
func miqt_exec_callback_QErrorMessage_FocusOutEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_EnterEvent(event *QEnterEvent) {

	C.QErrorMessage_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	ok := C.QErrorMessage_override_virtual_EnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_EnterEvent
func miqt_exec_callback_QErrorMessage_EnterEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_LeaveEvent(event *QEvent) {

	C.QErrorMessage_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QErrorMessage_override_virtual_LeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_LeaveEvent
func miqt_exec_callback_QErrorMessage_LeaveEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_PaintEvent(event *QPaintEvent) {

	C.QErrorMessage_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	ok := C.QErrorMessage_override_virtual_PaintEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_PaintEvent
func miqt_exec_callback_QErrorMessage_PaintEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_MoveEvent(event *QMoveEvent) {

	C.QErrorMessage_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	ok := C.QErrorMessage_override_virtual_MoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_MoveEvent
func miqt_exec_callback_QErrorMessage_MoveEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_TabletEvent(event *QTabletEvent) {

	C.QErrorMessage_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	ok := C.QErrorMessage_override_virtual_TabletEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_TabletEvent
func miqt_exec_callback_QErrorMessage_TabletEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ActionEvent(event *QActionEvent) {

	C.QErrorMessage_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	ok := C.QErrorMessage_override_virtual_ActionEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_ActionEvent
func miqt_exec_callback_QErrorMessage_ActionEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	C.QErrorMessage_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	ok := C.QErrorMessage_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_DragEnterEvent
func miqt_exec_callback_QErrorMessage_DragEnterEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	C.QErrorMessage_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	ok := C.QErrorMessage_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_DragMoveEvent
func miqt_exec_callback_QErrorMessage_DragMoveEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	C.QErrorMessage_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	ok := C.QErrorMessage_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_DragLeaveEvent
func miqt_exec_callback_QErrorMessage_DragLeaveEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_DropEvent(event *QDropEvent) {

	C.QErrorMessage_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	ok := C.QErrorMessage_override_virtual_DropEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_DropEvent
func miqt_exec_callback_QErrorMessage_DropEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_HideEvent(event *QHideEvent) {

	C.QErrorMessage_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	ok := C.QErrorMessage_override_virtual_HideEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_HideEvent
func miqt_exec_callback_QErrorMessage_HideEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := C.struct_miqt_string{}
	eventType_alias.data = (*C.char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = C.size_t(len(eventType))

	return (bool)(C.QErrorMessage_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*C.intptr_t)(unsafe.Pointer(result))))

}
func (this *QErrorMessage) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	ok := C.QErrorMessage_override_virtual_NativeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_NativeEvent
func miqt_exec_callback_QErrorMessage_NativeEvent(self *C.QErrorMessage, cb C.intptr_t, eventType C.struct_miqt_string, message unsafe.Pointer, result *C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray C.struct_miqt_string = eventType
	eventType_ret := C.GoBytes(unsafe.Pointer(eventType_bytearray.data), C.int(int64(eventType_bytearray.len)))
	C.free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (C.bool)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_Metric(param1 QPaintDevice__PaintDeviceMetric) int {

	return (int)(C.QErrorMessage_virtualbase_Metric(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QErrorMessage) OnMetric(slot func(super func(param1 QPaintDevice__PaintDeviceMetric) int, param1 QPaintDevice__PaintDeviceMetric) int) {
	ok := C.QErrorMessage_override_virtual_Metric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Metric
func miqt_exec_callback_QErrorMessage_Metric(self *C.QErrorMessage, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 QPaintDevice__PaintDeviceMetric) int, param1 QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QPaintDevice__PaintDeviceMetric)(param1)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_Metric, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_InitPainter(painter *QPainter) {

	C.QErrorMessage_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QErrorMessage) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	ok := C.QErrorMessage_override_virtual_InitPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_InitPainter
func miqt_exec_callback_QErrorMessage_InitPainter(self *C.QErrorMessage, cb C.intptr_t, painter *C.QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QErrorMessage) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(C.QErrorMessage_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QErrorMessage) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	ok := C.QErrorMessage_override_virtual_Redirected(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_Redirected
func miqt_exec_callback_QErrorMessage_Redirected(self *C.QErrorMessage, cb C.intptr_t, offset *C.QPoint) *C.QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(C.QErrorMessage_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QErrorMessage) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	ok := C.QErrorMessage_override_virtual_SharedPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_SharedPainter
func miqt_exec_callback_QErrorMessage_SharedPainter(self *C.QErrorMessage, cb C.intptr_t) *C.QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	C.QErrorMessage_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QErrorMessage) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	ok := C.QErrorMessage_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_InputMethodEvent
func miqt_exec_callback_QErrorMessage_InputMethodEvent(self *C.QErrorMessage, cb C.intptr_t, param1 *C.QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(C.QErrorMessage_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (C.int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QErrorMessage) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	ok := C.QErrorMessage_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_InputMethodQuery
func miqt_exec_callback_QErrorMessage_InputMethodQuery(self *C.QErrorMessage, cb C.intptr_t, param1 C.int) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QErrorMessage) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(C.QErrorMessage_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (C.bool)(next)))

}
func (this *QErrorMessage) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	ok := C.QErrorMessage_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_FocusNextPrevChild
func miqt_exec_callback_QErrorMessage_FocusNextPrevChild(self *C.QErrorMessage, cb C.intptr_t, next C.bool) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QErrorMessage{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QErrorMessage) callVirtualBase_TimerEvent(event *QTimerEvent) {

	C.QErrorMessage_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	ok := C.QErrorMessage_override_virtual_TimerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_TimerEvent
func miqt_exec_callback_QErrorMessage_TimerEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ChildEvent(event *QChildEvent) {

	C.QErrorMessage_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	ok := C.QErrorMessage_override_virtual_ChildEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_ChildEvent
func miqt_exec_callback_QErrorMessage_ChildEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_CustomEvent(event *QEvent) {

	C.QErrorMessage_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QErrorMessage) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QErrorMessage_override_virtual_CustomEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_CustomEvent
func miqt_exec_callback_QErrorMessage_CustomEvent(self *C.QErrorMessage, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QErrorMessage) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	C.QErrorMessage_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QErrorMessage) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QErrorMessage_override_virtual_ConnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_ConnectNotify
func miqt_exec_callback_QErrorMessage_ConnectNotify(self *C.QErrorMessage, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QErrorMessage) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	C.QErrorMessage_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QErrorMessage) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QErrorMessage_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QErrorMessage_DisconnectNotify
func miqt_exec_callback_QErrorMessage_DisconnectNotify(self *C.QErrorMessage, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QErrorMessage{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QErrorMessage) Delete() {
	C.QErrorMessage_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QErrorMessage) GoGC() {
	runtime.SetFinalizer(this, func(this *QErrorMessage) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
