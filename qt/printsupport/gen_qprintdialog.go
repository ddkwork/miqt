package printsupport

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QPrintDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQPrintDialog constructs a new QPrintDialog object.
func NewQPrintDialog(parent *qt.QWidget) *QPrintDialog {

	ret := newQPrintDialog(QPrintDialog_new((*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQPrintDialog2 constructs a new QPrintDialog object.
func NewQPrintDialog2(printer *QPrinter) *QPrintDialog {

	ret := newQPrintDialog(QPrintDialog_new2(printer.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPrintDialog3 constructs a new QPrintDialog object.
func NewQPrintDialog3() *QPrintDialog {

	ret := newQPrintDialog(QPrintDialog_new3())
	ret.isSubclass = true
	return ret
}

// NewQPrintDialog4 constructs a new QPrintDialog object.
func NewQPrintDialog4(printer *QPrinter, parent *qt.QWidget) *QPrintDialog {

	ret := newQPrintDialog(QPrintDialog_new4(printer.cPointer(), (*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QPrintDialog) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QPrintDialog_MetaObject(this.h)))
}

func (this *QPrintDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPrintDialog_Metacast(this.h, param1_Cstring))
}

func QPrintDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPrintDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrintDialog) Exec() int {
	return (int)(QPrintDialog_Exec(this.h))
}

func (this *QPrintDialog) Done(result int) {
	QPrintDialog_Done(this.h, (int)(result))
}

func (this *QPrintDialog) SetOption(option PrintDialogOption) {
	QPrintDialog_SetOption(this.h, option)
}

func (this *QPrintDialog) TestOption(option PrintDialogOption) bool {
	return (bool)(QPrintDialog_TestOption(this.h, option))
}

func (this *QPrintDialog) SetOptions(options PrintDialogOptions) {
	QPrintDialog_SetOptions(this.h, options)
}

func (this *QPrintDialog) Options() PrintDialogOptions {
	xxxxxxxxx
}

func (this *QPrintDialog) SetVisible(visible bool) {
	QPrintDialog_SetVisible(this.h, (bool)(visible))
}

func (this *QPrintDialog) Accepted(printer *QPrinter) {
	QPrintDialog_Accepted(this.h, printer.cPointer())
}
func (this *QPrintDialog) OnAccepted(slot func(printer *QPrinter)) {
	QPrintDialog_connect_Accepted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintDialog_Accepted
func miqt_exec_callback_QPrintDialog_Accepted(cb intptr_t, printer *QPrinter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(printer *QPrinter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPrinter(printer)

	gofunc(slotval1)
}

func QPrintDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPrintDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPrintDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPrintDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrintDialog) SetOption2(option PrintDialogOption, on bool) {
	QPrintDialog_SetOption2(this.h, option, (bool)(on))
}

func (this *QPrintDialog) callVirtualBase_Exec() int {

	return (int)(QPrintDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QPrintDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintDialog_Exec
func miqt_exec_callback_QPrintDialog_Exec(self QPrintDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QPrintDialog) callVirtualBase_Done(result int) {

	QPrintDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(result))

}
func (this *QPrintDialog) OnDone(slot func(super func(result int), result int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintDialog_Done
func miqt_exec_callback_QPrintDialog_Done(self QPrintDialog, cb intptr_t, result int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(result int), result int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(result)

	gofunc((&QPrintDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QPrintDialog) callVirtualBase_SetVisible(visible bool) {

	QPrintDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QPrintDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintDialog_SetVisible
func miqt_exec_callback_QPrintDialog_SetVisible(self QPrintDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QPrintDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}