package qt

import (
	"unsafe"
)

type QObjectCleanupHandler struct {
	h          uintptr
	isSubclass bool
}

// NewQObjectCleanupHandler constructs a new QObjectCleanupHandler object.
func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	g := newQObjectCleanupHandler(QObjectCleanupHandler_new())
	g.isSubclass = true
	return g
}

func (this *QObjectCleanupHandler) MetaObject() *QMetaObject {
	return newQMetaObject(QObjectCleanupHandler_MetaObject(this.h))
}

func (this *QObjectCleanupHandler) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QObjectCleanupHandler_Metacast(this.h, param1_Cstring))
}

func QObjectCleanupHandler_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QObjectCleanupHandler_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObjectCleanupHandler) Add(object *QObject) *QObject {
	return newQObject(QObjectCleanupHandler_Add(this.h, object.cPointer()))
}

func (this *QObjectCleanupHandler) Remove(object *QObject) {
	QObjectCleanupHandler_Remove(this.h, object.cPointer())
}

func (this *QObjectCleanupHandler) IsEmpty() bool {
	return (bool)(QObjectCleanupHandler_IsEmpty(this.h))
}

func (this *QObjectCleanupHandler) Clear() {
	QObjectCleanupHandler_Clear(this.h)
}

func QObjectCleanupHandler_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QObjectCleanupHandler_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QObjectCleanupHandler_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QObjectCleanupHandler_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QObjectCleanupHandler) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QObjectCleanupHandler_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QObjectCleanupHandler) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_MetaObject
func miqt_exec_callback_QObjectCleanupHandler_MetaObject(self QObjectCleanupHandler, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QObjectCleanupHandler) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QObjectCleanupHandler_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QObjectCleanupHandler) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QObjectCleanupHandler_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QObjectCleanupHandler_Metacast
func miqt_exec_callback_QObjectCleanupHandler_Metacast(self QObjectCleanupHandler, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QObjectCleanupHandler{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
