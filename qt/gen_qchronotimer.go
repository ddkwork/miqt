package qt

import (
	"unsafe"
)

type QChronoTimer struct {
	h          uintptr
	isSubclass bool
}

// NewQChronoTimer constructs a new QChronoTimer object.
func NewQChronoTimer() *QChronoTimer {
	ret := newQChronoTimer(QChronoTimer_new())
	ret.isSubclass = true
	return ret
}

// NewQChronoTimer2 constructs a new QChronoTimer object.
func NewQChronoTimer2(parent *QObject) *QChronoTimer {
	ret := newQChronoTimer(QChronoTimer_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QChronoTimer) MetaObject() *QMetaObject {
	return newQMetaObject(QChronoTimer_MetaObject(this.h))
}

func (this *QChronoTimer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QChronoTimer_Metacast(this.h, param1_Cstring))
}

func QChronoTimer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QChronoTimer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QChronoTimer) IsActive() bool {
	return (bool)(QChronoTimer_IsActive(this.h))
}

func (this *QChronoTimer) Id() TimerId {
	return (TimerId)(QChronoTimer_Id(this.h))
}

func (this *QChronoTimer) SetTimerType(atype TimerType) {
	QChronoTimer_SetTimerType(this.h, (int)(atype))
}

func (this *QChronoTimer) TimerType() TimerType {
	return (TimerType)(QChronoTimer_TimerType(this.h))
}

func (this *QChronoTimer) SetSingleShot(singleShot bool) {
	QChronoTimer_SetSingleShot(this.h, (bool)(singleShot))
}

func (this *QChronoTimer) IsSingleShot() bool {
	return (bool)(QChronoTimer_IsSingleShot(this.h))
}

func (this *QChronoTimer) Start() {
	QChronoTimer_Start(this.h)
}

func (this *QChronoTimer) Stop() {
	QChronoTimer_Stop(this.h)
}

func QChronoTimer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QChronoTimer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QChronoTimer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QChronoTimer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QChronoTimer) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QChronoTimer_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QChronoTimer) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_MetaObject
func miqt_exec_callback_QChronoTimer_MetaObject(self QChronoTimer, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QChronoTimer{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QChronoTimer) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QChronoTimer_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QChronoTimer) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QChronoTimer_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QChronoTimer_Metacast
func miqt_exec_callback_QChronoTimer_Metacast(self QChronoTimer, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QChronoTimer{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
