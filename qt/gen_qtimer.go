package qt

import (
	"unsafe"
)

type QTimer struct {
	h          uintptr
	isSubclass bool
}

// NewQTimer constructs a new QTimer object.
func NewQTimer() *QTimer {
	ret := newQTimer(QTimer_new())
	ret.isSubclass = true
	return ret
}

// NewQTimer2 constructs a new QTimer object.
func NewQTimer2(parent *QObject) *QTimer {
	ret := newQTimer(QTimer_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTimer) MetaObject() *QMetaObject {
	return newQMetaObject(QTimer_MetaObject(this.h))
}

func (this *QTimer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTimer_Metacast(this.h, param1_Cstring))
}

func QTimer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTimer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimer) IsActive() bool {
	return (bool)(QTimer_IsActive(this.h))
}

func (this *QTimer) TimerId() int {
	return (int)(QTimer_TimerId(this.h))
}

func (this *QTimer) Id() TimerId {
	return (TimerId)(QTimer_Id(this.h))
}

func (this *QTimer) SetInterval(msec int) {
	QTimer_SetInterval(this.h, (int)(msec))
}

func (this *QTimer) Interval() int {
	return (int)(QTimer_Interval(this.h))
}

func (this *QTimer) RemainingTime() int {
	return (int)(QTimer_RemainingTime(this.h))
}

func (this *QTimer) SetTimerType(atype TimerType) {
	QTimer_SetTimerType(this.h, (int)(atype))
}

func (this *QTimer) TimerType() TimerType {
	return (TimerType)(QTimer_TimerType(this.h))
}

func (this *QTimer) SetSingleShot(singleShot bool) {
	QTimer_SetSingleShot(this.h, (bool)(singleShot))
}

func (this *QTimer) IsSingleShot() bool {
	return (bool)(QTimer_IsSingleShot(this.h))
}

func (this *QTimer) Start(msec int) {
	QTimer_Start(this.h, (int)(msec))
}

func (this *QTimer) Start2() {
	QTimer_Start2(this.h)
}

func (this *QTimer) Stop() {
	QTimer_Stop(this.h)
}

func QTimer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTimer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimer) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTimer_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTimer) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_MetaObject
func miqt_exec_callback_QTimer_MetaObject(self QTimer, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTimer{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTimer) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTimer_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTimer) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimer_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimer_Metacast
func miqt_exec_callback_QTimer_Metacast(self QTimer, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTimer{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
