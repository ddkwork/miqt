package qt

import (
	"unsafe"
)

type QSizeGrip struct {
	h          uintptr
	isSubclass bool
}

// NewQSizeGrip constructs a new QSizeGrip object.
func NewQSizeGrip(parent *QWidget) *QSizeGrip {
	g := newQSizeGrip(QSizeGrip_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSizeGrip) MetaObject() *QMetaObject {
	return newQMetaObject(QSizeGrip_MetaObject(this.h))
}

func (this *QSizeGrip) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSizeGrip_Metacast(this.h, param1_Cstring))
}

func QSizeGrip_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSizeGrip_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSizeGrip) SizeHint() *QSize {
	_goptr := newQSize(QSizeGrip_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeGrip) SetVisible(visible bool) {
	QSizeGrip_SetVisible(this.h, (bool)(visible))
}

func QSizeGrip_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSizeGrip_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSizeGrip_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSizeGrip_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSizeGrip) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSizeGrip_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSizeGrip) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_MetaObject
func miqt_exec_callback_QSizeGrip_MetaObject(self QSizeGrip, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSizeGrip) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSizeGrip_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSizeGrip) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSizeGrip_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSizeGrip_Metacast
func miqt_exec_callback_QSizeGrip_Metacast(self QSizeGrip, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSizeGrip{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
