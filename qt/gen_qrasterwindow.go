package qt

import (
	"unsafe"
)

type QRasterWindow struct {
	h          uintptr
	isSubclass bool
}

// NewQRasterWindow constructs a new QRasterWindow object.
func NewQRasterWindow() *QRasterWindow {
	g := newQRasterWindow(QRasterWindow_new())
	g.isSubclass = true
	return g
}

// NewQRasterWindow2 constructs a new QRasterWindow object.
func NewQRasterWindow2(parent *QWindow) *QRasterWindow {
	g := newQRasterWindow(QRasterWindow_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QRasterWindow) MetaObject() *QMetaObject {
	return newQMetaObject(QRasterWindow_MetaObject(this.h))
}

func (this *QRasterWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRasterWindow_Metacast(this.h, param1_Cstring))
}

func QRasterWindow_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRasterWindow_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRasterWindow_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRasterWindow_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRasterWindow) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QRasterWindow_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QRasterWindow) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_MetaObject
func miqt_exec_callback_QRasterWindow_MetaObject(self QRasterWindow, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QRasterWindow) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QRasterWindow_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QRasterWindow) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_Metacast
func miqt_exec_callback_QRasterWindow_Metacast(self QRasterWindow, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
