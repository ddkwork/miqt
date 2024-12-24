package qt

import (
	"unsafe"
)

type QDial struct {
	h          uintptr
	isSubclass bool
}

// NewQDial constructs a new QDial object.
func NewQDial(parent *QWidget) *QDial {
	g := newQDial(QDial_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQDial2 constructs a new QDial object.
func NewQDial2() *QDial {
	g := newQDial(QDial_new2())
	g.isSubclass = true
	return g
}

func (this *QDial) MetaObject() *QMetaObject {
	return newQMetaObject(QDial_MetaObject(this.h))
}

func (this *QDial) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDial_Metacast(this.h, param1_Cstring))
}

func QDial_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDial_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDial) Wrapping() bool {
	return (bool)(QDial_Wrapping(this.h))
}

func (this *QDial) NotchSize() int {
	return (int)(QDial_NotchSize(this.h))
}

func (this *QDial) SetNotchTarget(target float64) {
	QDial_SetNotchTarget(this.h, (double)(target))
}

func (this *QDial) NotchTarget() float64 {
	return (float64)(QDial_NotchTarget(this.h))
}

func (this *QDial) NotchesVisible() bool {
	return (bool)(QDial_NotchesVisible(this.h))
}

func (this *QDial) SizeHint() *QSize {
	_goptr := newQSize(QDial_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDial) MinimumSizeHint() *QSize {
	_goptr := newQSize(QDial_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDial) SetNotchesVisible(visible bool) {
	QDial_SetNotchesVisible(this.h, (bool)(visible))
}

func (this *QDial) SetWrapping(on bool) {
	QDial_SetWrapping(this.h, (bool)(on))
}

func QDial_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDial_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDial_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDial_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDial) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDial_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDial) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_MetaObject
func miqt_exec_callback_QDial_MetaObject(self QDial, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDial{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDial) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDial_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDial) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDial_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDial_Metacast
func miqt_exec_callback_QDial_Metacast(self QDial, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QDial{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
