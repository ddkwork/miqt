package printsupport

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
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

func (this *QPrintPreviewDialog) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QPrintPreviewDialog_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QPrintPreviewDialog) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_MetaObject
func miqt_exec_callback_QPrintPreviewDialog_MetaObject(self QPrintPreviewDialog, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QPrintPreviewDialog) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPrintPreviewDialog_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPrintPreviewDialog) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPrintPreviewDialog_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPrintPreviewDialog_Metacast
func miqt_exec_callback_QPrintPreviewDialog_Metacast(self QPrintPreviewDialog, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QPrintPreviewDialog{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
