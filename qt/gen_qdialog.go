package qt

import (
	"unsafe"
)

type QDialog__DialogCode int

const (
	QDialog__Rejected QDialog__DialogCode = 0
	QDialog__Accepted QDialog__DialogCode = 1
)

type QDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQDialog constructs a new QDialog object.
func NewQDialog(parent *QWidget) *QDialog {

	ret := newQDialog(QDialog_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDialog2 constructs a new QDialog object.
func NewQDialog2() *QDialog {

	ret := newQDialog(QDialog_new2())
	ret.isSubclass = true
	return ret
}

// NewQDialog3 constructs a new QDialog object.
func NewQDialog3(parent *QWidget, f WindowType) *QDialog {

	ret := newQDialog(QDialog_new3(parent.cPointer(), (int)(f)))
	ret.isSubclass = true
	return ret
}

func (this *QDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QDialog_MetaObject(this.h))
}

func (this *QDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDialog_Metacast(this.h, param1_Cstring))
}

func QDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDialog) Result() int {
	return (int)(QDialog_Result(this.h))
}

func (this *QDialog) SetVisible(visible bool) {
	QDialog_SetVisible(this.h, (bool)(visible))
}

func (this *QDialog) SizeHint() *QSize {
	_goptr := newQSize(QDialog_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDialog) MinimumSizeHint() *QSize {
	_goptr := newQSize(QDialog_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDialog) SetSizeGripEnabled(sizeGripEnabled bool) {
	QDialog_SetSizeGripEnabled(this.h, (bool)(sizeGripEnabled))
}

func (this *QDialog) IsSizeGripEnabled() bool {
	return (bool)(QDialog_IsSizeGripEnabled(this.h))
}

func (this *QDialog) SetModal(modal bool) {
	QDialog_SetModal(this.h, (bool)(modal))
}

func (this *QDialog) SetResult(r int) {
	QDialog_SetResult(this.h, (int)(r))
}

func (this *QDialog) Finished(result int) {
	QDialog_Finished(this.h, (int)(result))
}
func (this *QDialog) OnFinished(slot func(result int)) {
	QDialog_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Finished
func miqt_exec_callback_QDialog_Finished(cb intptr_t, result int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(result int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(result)

	gofunc(slotval1)
}

func (this *QDialog) Accepted() {
	QDialog_Accepted(this.h)
}
func (this *QDialog) OnAccepted(slot func()) {
	QDialog_connect_Accepted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Accepted
func miqt_exec_callback_QDialog_Accepted(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QDialog) Rejected() {
	QDialog_Rejected(this.h)
}
func (this *QDialog) OnRejected(slot func()) {
	QDialog_connect_Rejected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Rejected
func miqt_exec_callback_QDialog_Rejected(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QDialog) Open() {
	QDialog_Open(this.h)
}

func (this *QDialog) Exec() int {
	return (int)(QDialog_Exec(this.h))
}

func (this *QDialog) Done(param1 int) {
	QDialog_Done(this.h, (int)(param1))
}

func (this *QDialog) Accept() {
	QDialog_Accept(this.h)
}

func (this *QDialog) Reject() {
	QDialog_Reject(this.h)
}

func QDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDialog) callVirtualBase_SetVisible(visible bool) {

	QDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_SetVisible
func miqt_exec_callback_QDialog_SetVisible(self QDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QDialog) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QDialog_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDialog) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_SizeHint
func miqt_exec_callback_QDialog_SizeHint(self QDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QDialog) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QDialog_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDialog) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_MinimumSizeHint
func miqt_exec_callback_QDialog_MinimumSizeHint(self QDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QDialog) callVirtualBase_Open() {

	QDialog_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QDialog) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Open
func miqt_exec_callback_QDialog_Open(self QDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QDialog{h: self}).callVirtualBase_Open)

}

func (this *QDialog) callVirtualBase_Exec() int {

	return (int)(QDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Exec
func miqt_exec_callback_QDialog_Exec(self QDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QDialog) callVirtualBase_Done(param1 int) {

	QDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QDialog) OnDone(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Done
func miqt_exec_callback_QDialog_Done(self QDialog, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QDialog) callVirtualBase_Accept() {

	QDialog_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QDialog) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Accept
func miqt_exec_callback_QDialog_Accept(self QDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QDialog{h: self}).callVirtualBase_Accept)

}

func (this *QDialog) callVirtualBase_Reject() {

	QDialog_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QDialog) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Reject
func miqt_exec_callback_QDialog_Reject(self QDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QDialog{h: self}).callVirtualBase_Reject)

}

func (this *QDialog) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QDialog_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDialog) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_KeyPressEvent
func miqt_exec_callback_QDialog_KeyPressEvent(self QDialog, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QDialog) callVirtualBase_CloseEvent(param1 *QCloseEvent) {

	QDialog_virtualbase_CloseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDialog) OnCloseEvent(slot func(super func(param1 *QCloseEvent), param1 *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_CloseEvent
func miqt_exec_callback_QDialog_CloseEvent(self QDialog, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QCloseEvent), param1 *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QDialog) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QDialog_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDialog) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_ShowEvent
func miqt_exec_callback_QDialog_ShowEvent(self QDialog, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QDialog) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QDialog_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDialog) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_ResizeEvent
func miqt_exec_callback_QDialog_ResizeEvent(self QDialog, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QDialog) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QDialog_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDialog) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_ContextMenuEvent
func miqt_exec_callback_QDialog_ContextMenuEvent(self QDialog, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QDialog) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QDialog_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QDialog) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_EventFilter
func miqt_exec_callback_QDialog_EventFilter(self QDialog, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QDialog) callVirtualBase_DevType() int {

	return (int)(QDialog_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QDialog) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_DevType
func miqt_exec_callback_QDialog_DevType(self QDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QDialog) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QDialog_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QDialog) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_HeightForWidth
func miqt_exec_callback_QDialog_HeightForWidth(self QDialog, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QDialog) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QDialog_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QDialog) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_HasHeightForWidth
func miqt_exec_callback_QDialog_HasHeightForWidth(self QDialog, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QDialog) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QDialog_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QDialog) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_PaintEngine
func miqt_exec_callback_QDialog_PaintEngine(self QDialog, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QDialog) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QDialog_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QDialog) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Event
func miqt_exec_callback_QDialog_Event(self QDialog, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDialog) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QDialog_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_MousePressEvent
func miqt_exec_callback_QDialog_MousePressEvent(self QDialog, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QDialog) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QDialog_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_MouseReleaseEvent
func miqt_exec_callback_QDialog_MouseReleaseEvent(self QDialog, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QDialog) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QDialog_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_MouseDoubleClickEvent
func miqt_exec_callback_QDialog_MouseDoubleClickEvent(self QDialog, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QDialog) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QDialog_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_MouseMoveEvent
func miqt_exec_callback_QDialog_MouseMoveEvent(self QDialog, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QDialog) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QDialog_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_WheelEvent
func miqt_exec_callback_QDialog_WheelEvent(self QDialog, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QDialog) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QDialog_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_KeyReleaseEvent
func miqt_exec_callback_QDialog_KeyReleaseEvent(self QDialog, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QDialog) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QDialog_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_FocusInEvent
func miqt_exec_callback_QDialog_FocusInEvent(self QDialog, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QDialog) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QDialog_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_FocusOutEvent
func miqt_exec_callback_QDialog_FocusOutEvent(self QDialog, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QDialog) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QDialog_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_EnterEvent
func miqt_exec_callback_QDialog_EnterEvent(self QDialog, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QDialog) callVirtualBase_LeaveEvent(event *QEvent) {

	QDialog_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_LeaveEvent
func miqt_exec_callback_QDialog_LeaveEvent(self QDialog, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QDialog) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QDialog_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_PaintEvent
func miqt_exec_callback_QDialog_PaintEvent(self QDialog, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QDialog) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QDialog_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_MoveEvent
func miqt_exec_callback_QDialog_MoveEvent(self QDialog, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QDialog) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QDialog_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_TabletEvent
func miqt_exec_callback_QDialog_TabletEvent(self QDialog, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QDialog) callVirtualBase_ActionEvent(event *QActionEvent) {

	QDialog_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_ActionEvent
func miqt_exec_callback_QDialog_ActionEvent(self QDialog, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QDialog) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QDialog_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_DragEnterEvent
func miqt_exec_callback_QDialog_DragEnterEvent(self QDialog, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QDialog) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QDialog_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_DragMoveEvent
func miqt_exec_callback_QDialog_DragMoveEvent(self QDialog, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QDialog) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QDialog_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_DragLeaveEvent
func miqt_exec_callback_QDialog_DragLeaveEvent(self QDialog, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QDialog) callVirtualBase_DropEvent(event *QDropEvent) {

	QDialog_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_DropEvent
func miqt_exec_callback_QDialog_DropEvent(self QDialog, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QDialog) callVirtualBase_HideEvent(event *QHideEvent) {

	QDialog_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDialog) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_HideEvent
func miqt_exec_callback_QDialog_HideEvent(self QDialog, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QDialog{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QDialog) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QDialog_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QDialog) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_NativeEvent
func miqt_exec_callback_QDialog_NativeEvent(self QDialog, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QDialog) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QDialog_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDialog) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_ChangeEvent
func miqt_exec_callback_QDialog_ChangeEvent(self QDialog, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QDialog) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QDialog_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QDialog) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Metric
func miqt_exec_callback_QDialog_Metric(self QDialog, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QDialog) callVirtualBase_InitPainter(painter *QPainter) {

	QDialog_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QDialog) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_InitPainter
func miqt_exec_callback_QDialog_InitPainter(self QDialog, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QDialog{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QDialog) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QDialog_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QDialog) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_Redirected
func miqt_exec_callback_QDialog_Redirected(self QDialog, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDialog) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QDialog_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QDialog) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_SharedPainter
func miqt_exec_callback_QDialog_SharedPainter(self QDialog, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QDialog) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QDialog_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDialog) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_InputMethodEvent
func miqt_exec_callback_QDialog_InputMethodEvent(self QDialog, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QDialog{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QDialog) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QDialog_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDialog) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_InputMethodQuery
func miqt_exec_callback_QDialog_InputMethodQuery(self QDialog, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDialog) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QDialog_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QDialog) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDialog_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDialog_FocusNextPrevChild
func miqt_exec_callback_QDialog_FocusNextPrevChild(self QDialog, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QDialog{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
