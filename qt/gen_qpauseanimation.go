package qt

import (
	"unsafe"
)

type QPauseAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQPauseAnimation constructs a new QPauseAnimation object.
func NewQPauseAnimation() *QPauseAnimation {
	g := newQPauseAnimation(QPauseAnimation_new())
	g.isSubclass = true
	return g
}

// NewQPauseAnimation2 constructs a new QPauseAnimation object.
func NewQPauseAnimation2(msecs int) *QPauseAnimation {
	g := newQPauseAnimation(QPauseAnimation_new2((int)(msecs)))
	g.isSubclass = true
	return g
}

// NewQPauseAnimation3 constructs a new QPauseAnimation object.
func NewQPauseAnimation3(parent *QObject) *QPauseAnimation {
	g := newQPauseAnimation(QPauseAnimation_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPauseAnimation4 constructs a new QPauseAnimation object.
func NewQPauseAnimation4(msecs int, parent *QObject) *QPauseAnimation {
	g := newQPauseAnimation(QPauseAnimation_new4((int)(msecs), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPauseAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QPauseAnimation_MetaObject(this.h))
}

func (this *QPauseAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPauseAnimation_Metacast(this.h, param1_Cstring))
}

func QPauseAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPauseAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPauseAnimation) Duration() int {
	return (int)(QPauseAnimation_Duration(this.h))
}

func (this *QPauseAnimation) SetDuration(msecs int) {
	QPauseAnimation_SetDuration(this.h, (int)(msecs))
}

func QPauseAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPauseAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPauseAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPauseAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPauseAnimation) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPauseAnimation_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPauseAnimation) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPauseAnimation_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPauseAnimation_MetaObject
func miqt_exec_callback_QPauseAnimation_MetaObject(self QPauseAnimation, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPauseAnimation{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPauseAnimation) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPauseAnimation_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPauseAnimation) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPauseAnimation_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPauseAnimation_Metacast
func miqt_exec_callback_QPauseAnimation_Metacast(self QPauseAnimation, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QPauseAnimation{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
