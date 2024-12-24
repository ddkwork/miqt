package printsupport

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QPrintPreviewDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQPrintPreviewDialog constructs a new QPrintPreviewDialog object.
func NewQPrintPreviewDialog(parent *qt.QWidget) *QPrintPreviewDialog {

	ret := newQPrintPreviewDialog(QPrintPreviewDialog_new((*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewDialog2 constructs a new QPrintPreviewDialog object.
func NewQPrintPreviewDialog2() *QPrintPreviewDialog {

	ret := newQPrintPreviewDialog(QPrintPreviewDialog_new2())
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewDialog3 constructs a new QPrintPreviewDialog object.
func NewQPrintPreviewDialog3(printer *QPrinter) *QPrintPreviewDialog {

	ret := newQPrintPreviewDialog(QPrintPreviewDialog_new3(printer.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewDialog4 constructs a new QPrintPreviewDialog object.
func NewQPrintPreviewDialog4(parent *qt.QWidget, flags WindowType) *QPrintPreviewDialog {

	ret := newQPrintPreviewDialog(QPrintPreviewDialog_new4((*QWidget)(parent.UnsafePointer()), (int)(flags)))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewDialog5 constructs a new QPrintPreviewDialog object.
func NewQPrintPreviewDialog5(printer *QPrinter, parent *qt.QWidget) *QPrintPreviewDialog {

	ret := newQPrintPreviewDialog(QPrintPreviewDialog_new5(printer.cPointer(), (*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQPrintPreviewDialog6 constructs a new QPrintPreviewDialog object.
func NewQPrintPreviewDialog6(printer *QPrinter, parent *qt.QWidget, flags WindowType) *QPrintPreviewDialog {

	ret := newQPrintPreviewDialog(QPrintPreviewDialog_new6(printer.cPointer(), (*QWidget)(parent.UnsafePointer()), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QPrintPreviewDialog) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QPrintPreviewDialog_MetaObject(this.h)))
}

func (this *QPrintPreviewDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPrintPreviewDialog_Metacast(this.h, param1_Cstring))
}

func QPrintPreviewDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPrintPreviewDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrintPreviewDialog) Printer() *QPrinter {
	return newQPrinter(QPrintPreviewDialog_Printer(this.h))
}

func (this *QPrintPreviewDialog) SetVisible(visible bool) {
	QPrintPreviewDialog_SetVisible(this.h, (bool)(visible))
}

func (this *QPrintPreviewDialog) Done(result int) {
	QPrintPreviewDialog_Done(this.h, (int)(result))
}

func (this *QPrintPreviewDialog) PaintRequested(printer *QPrinter) {
	QPrintPreviewDialog_PaintRequested(this.h, printer.cPointer())
}
func (this *QPrintPreviewDialog) OnPaintRequested(slot func(printer *QPrinter)) {
	QPrintPreviewDialog_connect_PaintRequested(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_PaintRequested
func miqt_exec_callback_QPrintPreviewDialog_PaintRequested(cb intptr_t, printer *QPrinter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(printer *QPrinter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPrinter(printer)

	gofunc(slotval1)
}

func QPrintPreviewDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPrintPreviewDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPrintPreviewDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPrintPreviewDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrintPreviewDialog) callVirtualBase_SetVisible(visible bool) {

	QPrintPreviewDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QPrintPreviewDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_SetVisible
func miqt_exec_callback_QPrintPreviewDialog_SetVisible(self QPrintPreviewDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QPrintPreviewDialog) callVirtualBase_Done(result int) {

	QPrintPreviewDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(result))

}
func (this *QPrintPreviewDialog) OnDone(slot func(super func(result int), result int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_Done
func miqt_exec_callback_QPrintPreviewDialog_Done(self QPrintPreviewDialog, cb intptr_t, result int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(result int), result int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(result)

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QPrintPreviewDialog) callVirtualBase_SizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QPrintPreviewDialog_virtualbase_SizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPrintPreviewDialog) OnSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_SizeHint
func miqt_exec_callback_QPrintPreviewDialog_SizeHint(self QPrintPreviewDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_SizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewDialog) callVirtualBase_MinimumSizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QPrintPreviewDialog_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPrintPreviewDialog) OnMinimumSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_MinimumSizeHint
func miqt_exec_callback_QPrintPreviewDialog_MinimumSizeHint(self QPrintPreviewDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_MinimumSizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QPrintPreviewDialog) callVirtualBase_Open() {

	QPrintPreviewDialog_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QPrintPreviewDialog) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_Open
func miqt_exec_callback_QPrintPreviewDialog_Open(self QPrintPreviewDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_Open)

}

func (this *QPrintPreviewDialog) callVirtualBase_Exec() int {

	return (int)(QPrintPreviewDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QPrintPreviewDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_Exec
func miqt_exec_callback_QPrintPreviewDialog_Exec(self QPrintPreviewDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QPrintPreviewDialog) callVirtualBase_Accept() {

	QPrintPreviewDialog_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QPrintPreviewDialog) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_Accept
func miqt_exec_callback_QPrintPreviewDialog_Accept(self QPrintPreviewDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_Accept)

}

func (this *QPrintPreviewDialog) callVirtualBase_Reject() {

	QPrintPreviewDialog_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QPrintPreviewDialog) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_Reject
func miqt_exec_callback_QPrintPreviewDialog_Reject(self QPrintPreviewDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_Reject)

}

func (this *QPrintPreviewDialog) callVirtualBase_KeyPressEvent(param1 *qt.QKeyEvent) {

	QPrintPreviewDialog_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), (*QKeyEvent)(param1.UnsafePointer()))

}
func (this *QPrintPreviewDialog) OnKeyPressEvent(slot func(super func(param1 *qt.QKeyEvent), param1 *qt.QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_KeyPressEvent
func miqt_exec_callback_QPrintPreviewDialog_KeyPressEvent(self QPrintPreviewDialog, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QKeyEvent), param1 *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(param1))

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QPrintPreviewDialog) callVirtualBase_CloseEvent(param1 *qt.QCloseEvent) {

	QPrintPreviewDialog_virtualbase_CloseEvent(unsafe.Pointer(this.h), (*QCloseEvent)(param1.UnsafePointer()))

}
func (this *QPrintPreviewDialog) OnCloseEvent(slot func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_CloseEvent
func miqt_exec_callback_QPrintPreviewDialog_CloseEvent(self QPrintPreviewDialog, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQCloseEvent(unsafe.Pointer(param1))

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QPrintPreviewDialog) callVirtualBase_ShowEvent(param1 *qt.QShowEvent) {

	QPrintPreviewDialog_virtualbase_ShowEvent(unsafe.Pointer(this.h), (*QShowEvent)(param1.UnsafePointer()))

}
func (this *QPrintPreviewDialog) OnShowEvent(slot func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_ShowEvent
func miqt_exec_callback_QPrintPreviewDialog_ShowEvent(self QPrintPreviewDialog, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQShowEvent(unsafe.Pointer(param1))

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QPrintPreviewDialog) callVirtualBase_ResizeEvent(param1 *qt.QResizeEvent) {

	QPrintPreviewDialog_virtualbase_ResizeEvent(unsafe.Pointer(this.h), (*QResizeEvent)(param1.UnsafePointer()))

}
func (this *QPrintPreviewDialog) OnResizeEvent(slot func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_ResizeEvent
func miqt_exec_callback_QPrintPreviewDialog_ResizeEvent(self QPrintPreviewDialog, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQResizeEvent(unsafe.Pointer(param1))

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QPrintPreviewDialog) callVirtualBase_ContextMenuEvent(param1 *qt.QContextMenuEvent) {

	QPrintPreviewDialog_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), (*QContextMenuEvent)(param1.UnsafePointer()))

}
func (this *QPrintPreviewDialog) OnContextMenuEvent(slot func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_ContextMenuEvent
func miqt_exec_callback_QPrintPreviewDialog_ContextMenuEvent(self QPrintPreviewDialog, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQContextMenuEvent(unsafe.Pointer(param1))

	gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QPrintPreviewDialog) callVirtualBase_EventFilter(param1 *qt.QObject, param2 *qt.QEvent) bool {

	return (bool)(QPrintPreviewDialog_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(param1.UnsafePointer()), (*QEvent)(param2.UnsafePointer())))

}
func (this *QPrintPreviewDialog) OnEventFilter(slot func(super func(param1 *qt.QObject, param2 *qt.QEvent) bool, param1 *qt.QObject, param2 *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_EventFilter
func miqt_exec_callback_QPrintPreviewDialog_EventFilter(self QPrintPreviewDialog, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QObject, param2 *qt.QEvent) bool, param1 *qt.QObject, param2 *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(param1))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(param2))

	virtualReturn := gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
