package qt

import (
	"unsafe"
)

type QWinEventNotifier struct {
	h          uintptr
	isSubclass bool
}

// NewQWinEventNotifier constructs a new QWinEventNotifier object.
func NewQWinEventNotifier() *QWinEventNotifier {
	g := newQWinEventNotifier(QWinEventNotifier_new())
	g.isSubclass = true
	return g
}

// NewQWinEventNotifier2 constructs a new QWinEventNotifier object.
func NewQWinEventNotifier2(hEvent HANDLE) *QWinEventNotifier {
	g := newQWinEventNotifier(QWinEventNotifier_new2(hEvent))
	g.isSubclass = true
	return g
}

// NewQWinEventNotifier3 constructs a new QWinEventNotifier object.
func NewQWinEventNotifier3(parent *QObject) *QWinEventNotifier {
	g := newQWinEventNotifier(QWinEventNotifier_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQWinEventNotifier4 constructs a new QWinEventNotifier object.
func NewQWinEventNotifier4(hEvent HANDLE, parent *QObject) *QWinEventNotifier {
	g := newQWinEventNotifier(QWinEventNotifier_new4(hEvent, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QWinEventNotifier) MetaObject() *QMetaObject {
	return newQMetaObject(QWinEventNotifier_MetaObject(this.h))
}

func (this *QWinEventNotifier) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWinEventNotifier_Metacast(this.h, param1_Cstring))
}

func QWinEventNotifier_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWinEventNotifier_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWinEventNotifier) SetHandle(hEvent HANDLE) {
	QWinEventNotifier_SetHandle(this.h, hEvent)
}

func (this *QWinEventNotifier) Handle() HANDLE {
	xxxxxxxxx
}

func (this *QWinEventNotifier) IsEnabled() bool {
	return (bool)(QWinEventNotifier_IsEnabled(this.h))
}

func (this *QWinEventNotifier) SetEnabled(enable bool) {
	QWinEventNotifier_SetEnabled(this.h, (bool)(enable))
}

func QWinEventNotifier_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWinEventNotifier_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWinEventNotifier_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWinEventNotifier_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWinEventNotifier) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QWinEventNotifier_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QWinEventNotifier) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_MetaObject
func miqt_exec_callback_QWinEventNotifier_MetaObject(self QWinEventNotifier, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWinEventNotifier{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QWinEventNotifier) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWinEventNotifier_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWinEventNotifier) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWinEventNotifier_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWinEventNotifier_Metacast
func miqt_exec_callback_QWinEventNotifier_Metacast(self QWinEventNotifier, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QWinEventNotifier{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
