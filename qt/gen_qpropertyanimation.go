package qt

import (
	"unsafe"
)

type QPropertyAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQPropertyAnimation constructs a new QPropertyAnimation object.
func NewQPropertyAnimation() *QPropertyAnimation {
	g := newQPropertyAnimation(QPropertyAnimation_new())
	g.isSubclass = true
	return g
}

// NewQPropertyAnimation2 constructs a new QPropertyAnimation object.
func NewQPropertyAnimation2(target *QObject, propertyName []byte) *QPropertyAnimation {
	propertyName_alias := struct_miqt_string{}
	propertyName_alias.data = (char)(unsafe.Pointer(&propertyName[0]))
	propertyName_alias.len = size_t(len(propertyName))
	g := newQPropertyAnimation(QPropertyAnimation_new2(target.cPointer(), propertyName_alias))
	g.isSubclass = true
	return g
}

// NewQPropertyAnimation3 constructs a new QPropertyAnimation object.
func NewQPropertyAnimation3(parent *QObject) *QPropertyAnimation {
	g := newQPropertyAnimation(QPropertyAnimation_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPropertyAnimation4 constructs a new QPropertyAnimation object.
func NewQPropertyAnimation4(target *QObject, propertyName []byte, parent *QObject) *QPropertyAnimation {
	propertyName_alias := struct_miqt_string{}
	propertyName_alias.data = (char)(unsafe.Pointer(&propertyName[0]))
	propertyName_alias.len = size_t(len(propertyName))
	g := newQPropertyAnimation(QPropertyAnimation_new4(target.cPointer(), propertyName_alias, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPropertyAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QPropertyAnimation_MetaObject(this.h))
}

func (this *QPropertyAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPropertyAnimation_Metacast(this.h, param1_Cstring))
}

func QPropertyAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPropertyAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPropertyAnimation) TargetObject() *QObject {
	return newQObject(QPropertyAnimation_TargetObject(this.h))
}

func (this *QPropertyAnimation) SetTargetObject(target *QObject) {
	QPropertyAnimation_SetTargetObject(this.h, target.cPointer())
}

func (this *QPropertyAnimation) PropertyName() []byte {
	var _bytearray struct_miqt_string = QPropertyAnimation_PropertyName(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QPropertyAnimation) SetPropertyName(propertyName []byte) {
	propertyName_alias := struct_miqt_string{}
	propertyName_alias.data = (char)(unsafe.Pointer(&propertyName[0]))
	propertyName_alias.len = size_t(len(propertyName))
	QPropertyAnimation_SetPropertyName(this.h, propertyName_alias)
}

func QPropertyAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPropertyAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPropertyAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPropertyAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPropertyAnimation) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPropertyAnimation_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPropertyAnimation) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_MetaObject
func miqt_exec_callback_QPropertyAnimation_MetaObject(self QPropertyAnimation, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPropertyAnimation{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPropertyAnimation) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPropertyAnimation_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPropertyAnimation) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPropertyAnimation_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPropertyAnimation_Metacast
func miqt_exec_callback_QPropertyAnimation_Metacast(self QPropertyAnimation, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QPropertyAnimation{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
