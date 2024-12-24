package printsupport

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QPageSetupDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQPageSetupDialog constructs a new QPageSetupDialog object.
func NewQPageSetupDialog(parent *qt.QWidget) *QPageSetupDialog {

	ret := newQPageSetupDialog(QPageSetupDialog_new((*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQPageSetupDialog2 constructs a new QPageSetupDialog object.
func NewQPageSetupDialog2(printer *QPrinter) *QPageSetupDialog {

	ret := newQPageSetupDialog(QPageSetupDialog_new2(printer.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPageSetupDialog3 constructs a new QPageSetupDialog object.
func NewQPageSetupDialog3() *QPageSetupDialog {

	ret := newQPageSetupDialog(QPageSetupDialog_new3())
	ret.isSubclass = true
	return ret
}

// NewQPageSetupDialog4 constructs a new QPageSetupDialog object.
func NewQPageSetupDialog4(printer *QPrinter, parent *qt.QWidget) *QPageSetupDialog {

	ret := newQPageSetupDialog(QPageSetupDialog_new4(printer.cPointer(), (*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QPageSetupDialog) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QPageSetupDialog_MetaObject(this.h)))
}

func (this *QPageSetupDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPageSetupDialog_Metacast(this.h, param1_Cstring))
}

func QPageSetupDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPageSetupDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPageSetupDialog) SetVisible(visible bool) {
	QPageSetupDialog_SetVisible(this.h, (bool)(visible))
}

func (this *QPageSetupDialog) Exec() int {
	return (int)(QPageSetupDialog_Exec(this.h))
}

func (this *QPageSetupDialog) Done(result int) {
	QPageSetupDialog_Done(this.h, (int)(result))
}

func (this *QPageSetupDialog) Printer() *QPrinter {
	return newQPrinter(QPageSetupDialog_Printer(this.h))
}

func QPageSetupDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPageSetupDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPageSetupDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPageSetupDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPageSetupDialog) callVirtualBase_SetVisible(visible bool) {

	QPageSetupDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QPageSetupDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_SetVisible
func miqt_exec_callback_QPageSetupDialog_SetVisible(self QPageSetupDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QPageSetupDialog) callVirtualBase_Exec() int {

	return (int)(QPageSetupDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QPageSetupDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_Exec
func miqt_exec_callback_QPageSetupDialog_Exec(self QPageSetupDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPageSetupDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QPageSetupDialog) callVirtualBase_Done(result int) {

	QPageSetupDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(result))

}
func (this *QPageSetupDialog) OnDone(slot func(super func(result int), result int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_Done
func miqt_exec_callback_QPageSetupDialog_Done(self QPageSetupDialog, cb intptr_t, result int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(result int), result int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(result)

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QPageSetupDialog) callVirtualBase_SizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QPageSetupDialog_virtualbase_SizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPageSetupDialog) OnSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_SizeHint
func miqt_exec_callback_QPageSetupDialog_SizeHint(self QPageSetupDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPageSetupDialog{h: self}).callVirtualBase_SizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QPageSetupDialog) callVirtualBase_MinimumSizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QPageSetupDialog_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPageSetupDialog) OnMinimumSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_MinimumSizeHint
func miqt_exec_callback_QPageSetupDialog_MinimumSizeHint(self QPageSetupDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPageSetupDialog{h: self}).callVirtualBase_MinimumSizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QPageSetupDialog) callVirtualBase_Open() {

	QPageSetupDialog_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QPageSetupDialog) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_Open
func miqt_exec_callback_QPageSetupDialog_Open(self QPageSetupDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_Open)

}

func (this *QPageSetupDialog) callVirtualBase_Accept() {

	QPageSetupDialog_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QPageSetupDialog) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_Accept
func miqt_exec_callback_QPageSetupDialog_Accept(self QPageSetupDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_Accept)

}

func (this *QPageSetupDialog) callVirtualBase_Reject() {

	QPageSetupDialog_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QPageSetupDialog) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_Reject
func miqt_exec_callback_QPageSetupDialog_Reject(self QPageSetupDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_Reject)

}

func (this *QPageSetupDialog) callVirtualBase_KeyPressEvent(param1 *qt.QKeyEvent) {

	QPageSetupDialog_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), (*QKeyEvent)(param1.UnsafePointer()))

}
func (this *QPageSetupDialog) OnKeyPressEvent(slot func(super func(param1 *qt.QKeyEvent), param1 *qt.QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_KeyPressEvent
func miqt_exec_callback_QPageSetupDialog_KeyPressEvent(self QPageSetupDialog, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QKeyEvent), param1 *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(param1))

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QPageSetupDialog) callVirtualBase_CloseEvent(param1 *qt.QCloseEvent) {

	QPageSetupDialog_virtualbase_CloseEvent(unsafe.Pointer(this.h), (*QCloseEvent)(param1.UnsafePointer()))

}
func (this *QPageSetupDialog) OnCloseEvent(slot func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_CloseEvent
func miqt_exec_callback_QPageSetupDialog_CloseEvent(self QPageSetupDialog, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQCloseEvent(unsafe.Pointer(param1))

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QPageSetupDialog) callVirtualBase_ShowEvent(param1 *qt.QShowEvent) {

	QPageSetupDialog_virtualbase_ShowEvent(unsafe.Pointer(this.h), (*QShowEvent)(param1.UnsafePointer()))

}
func (this *QPageSetupDialog) OnShowEvent(slot func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_ShowEvent
func miqt_exec_callback_QPageSetupDialog_ShowEvent(self QPageSetupDialog, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQShowEvent(unsafe.Pointer(param1))

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QPageSetupDialog) callVirtualBase_ResizeEvent(param1 *qt.QResizeEvent) {

	QPageSetupDialog_virtualbase_ResizeEvent(unsafe.Pointer(this.h), (*QResizeEvent)(param1.UnsafePointer()))

}
func (this *QPageSetupDialog) OnResizeEvent(slot func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_ResizeEvent
func miqt_exec_callback_QPageSetupDialog_ResizeEvent(self QPageSetupDialog, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQResizeEvent(unsafe.Pointer(param1))

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QPageSetupDialog) callVirtualBase_ContextMenuEvent(param1 *qt.QContextMenuEvent) {

	QPageSetupDialog_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), (*QContextMenuEvent)(param1.UnsafePointer()))

}
func (this *QPageSetupDialog) OnContextMenuEvent(slot func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_ContextMenuEvent
func miqt_exec_callback_QPageSetupDialog_ContextMenuEvent(self QPageSetupDialog, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQContextMenuEvent(unsafe.Pointer(param1))

	gofunc((&QPageSetupDialog{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QPageSetupDialog) callVirtualBase_EventFilter(param1 *qt.QObject, param2 *qt.QEvent) bool {

	return (bool)(QPageSetupDialog_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(param1.UnsafePointer()), (*QEvent)(param2.UnsafePointer())))

}
func (this *QPageSetupDialog) OnEventFilter(slot func(super func(param1 *qt.QObject, param2 *qt.QEvent) bool, param1 *qt.QObject, param2 *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPageSetupDialog_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPageSetupDialog_EventFilter
func miqt_exec_callback_QPageSetupDialog_EventFilter(self QPageSetupDialog, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QObject, param2 *qt.QEvent) bool, param1 *qt.QObject, param2 *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(param1))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(param2))

	virtualReturn := gofunc((&QPageSetupDialog{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
