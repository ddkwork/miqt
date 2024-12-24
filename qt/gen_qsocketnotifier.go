package qt

import (
	"unsafe"
)

type QSocketNotifier__Type int

const (
	QSocketNotifier__Read      QSocketNotifier__Type = 0
	QSocketNotifier__Write     QSocketNotifier__Type = 1
	QSocketNotifier__Exception QSocketNotifier__Type = 2
)

type QSocketNotifier struct {
	h          uintptr
	isSubclass bool
}

// NewQSocketNotifier constructs a new QSocketNotifier object.
func NewQSocketNotifier(param1 Type) *QSocketNotifier {
	ret := newQSocketNotifier(QSocketNotifier_new(param1))
	ret.isSubclass = true
	return ret
}

// NewQSocketNotifier2 constructs a new QSocketNotifier object.
func NewQSocketNotifier2(socket uintptr, param2 Type) *QSocketNotifier {
	ret := newQSocketNotifier(QSocketNotifier_new2((intptr_t)(socket), param2))
	ret.isSubclass = true
	return ret
}

// NewQSocketNotifier3 constructs a new QSocketNotifier object.
func NewQSocketNotifier3(param1 Type, parent *QObject) *QSocketNotifier {
	ret := newQSocketNotifier(QSocketNotifier_new3(param1, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSocketNotifier4 constructs a new QSocketNotifier object.
func NewQSocketNotifier4(socket uintptr, param2 Type, parent *QObject) *QSocketNotifier {
	ret := newQSocketNotifier(QSocketNotifier_new4((intptr_t)(socket), param2, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSocketNotifier) MetaObject() *QMetaObject {
	return newQMetaObject(QSocketNotifier_MetaObject(this.h))
}

func (this *QSocketNotifier) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSocketNotifier_Metacast(this.h, param1_Cstring))
}

func QSocketNotifier_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSocketNotifier_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSocketNotifier) SetSocket(socket uintptr) {
	QSocketNotifier_SetSocket(this.h, (intptr_t)(socket))
}

func (this *QSocketNotifier) Socket() uintptr {
	return (uintptr)(QSocketNotifier_Socket(this.h))
}

func (this *QSocketNotifier) Type() Type {
	xxxxxxxxx
}

func (this *QSocketNotifier) IsValid() bool {
	return (bool)(QSocketNotifier_IsValid(this.h))
}

func (this *QSocketNotifier) IsEnabled() bool {
	return (bool)(QSocketNotifier_IsEnabled(this.h))
}

func (this *QSocketNotifier) SetEnabled(enabled bool) {
	QSocketNotifier_SetEnabled(this.h, (bool)(enabled))
}

func QSocketNotifier_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSocketNotifier_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSocketNotifier_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSocketNotifier_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSocketNotifier) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSocketNotifier_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSocketNotifier) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_MetaObject
func miqt_exec_callback_QSocketNotifier_MetaObject(self QSocketNotifier, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSocketNotifier{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSocketNotifier) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSocketNotifier_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSocketNotifier) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSocketNotifier_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSocketNotifier_Metacast
func miqt_exec_callback_QSocketNotifier_Metacast(self QSocketNotifier, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QSocketNotifier{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QSocketDescriptor struct {
	h          uintptr
	isSubclass bool
}

// NewQSocketDescriptor constructs a new QSocketDescriptor object.
func NewQSocketDescriptor() *QSocketDescriptor {
	ret := newQSocketDescriptor(QSocketDescriptor_new())
	ret.isSubclass = true
	return ret
}

// NewQSocketDescriptor2 constructs a new QSocketDescriptor object.
func NewQSocketDescriptor2(desc uintptr) *QSocketDescriptor {
	ret := newQSocketDescriptor(QSocketDescriptor_new2((intptr_t)(desc)))
	ret.isSubclass = true
	return ret
}

// NewQSocketDescriptor3 constructs a new QSocketDescriptor object.
func NewQSocketDescriptor3(param1 *QSocketDescriptor) *QSocketDescriptor {
	ret := newQSocketDescriptor(QSocketDescriptor_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSocketDescriptor4 constructs a new QSocketDescriptor object.
func NewQSocketDescriptor4(descriptor DescriptorType) *QSocketDescriptor {
	ret := newQSocketDescriptor(QSocketDescriptor_new4(descriptor))
	ret.isSubclass = true
	return ret
}

func (this *QSocketDescriptor) WinHandle() unsafe.Pointer {
	return (unsafe.Pointer)(QSocketDescriptor_WinHandle(this.h))
}

func (this *QSocketDescriptor) IsValid() bool {
	return (bool)(QSocketDescriptor_IsValid(this.h))
}
