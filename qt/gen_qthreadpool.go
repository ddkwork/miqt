package qt

import (
	"unsafe"
)

type QThreadPool struct {
	h          uintptr
	isSubclass bool
}

// NewQThreadPool constructs a new QThreadPool object.
func NewQThreadPool() *QThreadPool {
	g := newQThreadPool(QThreadPool_new())
	g.isSubclass = true
	return g
}

// NewQThreadPool2 constructs a new QThreadPool object.
func NewQThreadPool2(parent *QObject) *QThreadPool {
	g := newQThreadPool(QThreadPool_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QThreadPool) MetaObject() *QMetaObject {
	return newQMetaObject(QThreadPool_MetaObject(this.h))
}

func (this *QThreadPool) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QThreadPool_Metacast(this.h, param1_Cstring))
}

func QThreadPool_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QThreadPool_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThreadPool_GlobalInstance() *QThreadPool {
	return newQThreadPool(QThreadPool_GlobalInstance())
}

func (this *QThreadPool) Start(runnable *QRunnable) {
	QThreadPool_Start(this.h, runnable.cPointer())
}

func (this *QThreadPool) TryStart(runnable *QRunnable) bool {
	return (bool)(QThreadPool_TryStart(this.h, runnable.cPointer()))
}

func (this *QThreadPool) StartOnReservedThread(runnable *QRunnable) {
	QThreadPool_StartOnReservedThread(this.h, runnable.cPointer())
}

func (this *QThreadPool) ExpiryTimeout() int {
	return (int)(QThreadPool_ExpiryTimeout(this.h))
}

func (this *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	QThreadPool_SetExpiryTimeout(this.h, (int)(expiryTimeout))
}

func (this *QThreadPool) MaxThreadCount() int {
	return (int)(QThreadPool_MaxThreadCount(this.h))
}

func (this *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	QThreadPool_SetMaxThreadCount(this.h, (int)(maxThreadCount))
}

func (this *QThreadPool) ActiveThreadCount() int {
	return (int)(QThreadPool_ActiveThreadCount(this.h))
}

func (this *QThreadPool) SetStackSize(stackSize uint) {
	QThreadPool_SetStackSize(this.h, (uint)(stackSize))
}

func (this *QThreadPool) StackSize() uint {
	return (uint)(QThreadPool_StackSize(this.h))
}

func (this *QThreadPool) SetThreadPriority(priority QThread__Priority) {
	QThreadPool_SetThreadPriority(this.h, (int)(priority))
}

func (this *QThreadPool) ThreadPriority() QThread__Priority {
	return (QThread__Priority)(QThreadPool_ThreadPriority(this.h))
}

func (this *QThreadPool) ReserveThread() {
	QThreadPool_ReserveThread(this.h)
}

func (this *QThreadPool) ReleaseThread() {
	QThreadPool_ReleaseThread(this.h)
}

func (this *QThreadPool) SetServiceLevel(serviceLevel QThread__QualityOfService) {
	QThreadPool_SetServiceLevel(this.h, (int)(serviceLevel))
}

func (this *QThreadPool) ServiceLevel() QThread__QualityOfService {
	return (QThread__QualityOfService)(QThreadPool_ServiceLevel(this.h))
}

func (this *QThreadPool) WaitForDone(msecs int) bool {
	return (bool)(QThreadPool_WaitForDone(this.h, (int)(msecs)))
}

func (this *QThreadPool) WaitForDone2() bool {
	return (bool)(QThreadPool_WaitForDone2(this.h))
}

func (this *QThreadPool) Clear() {
	QThreadPool_Clear(this.h)
}

func (this *QThreadPool) Contains(thread *QThread) bool {
	return (bool)(QThreadPool_Contains(this.h, thread.cPointer()))
}

func (this *QThreadPool) TryTake(runnable *QRunnable) bool {
	return (bool)(QThreadPool_TryTake(this.h, runnable.cPointer()))
}

func QThreadPool_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThreadPool_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThreadPool_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThreadPool_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QThreadPool) Start2(runnable *QRunnable, priority int) {
	QThreadPool_Start2(this.h, runnable.cPointer(), (int)(priority))
}

func (this *QThreadPool) WaitForDone1(deadline QDeadlineTimer) bool {
	return (bool)(QThreadPool_WaitForDone1(this.h, deadline.cPointer()))
}

func (this *QThreadPool) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QThreadPool_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QThreadPool) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_MetaObject
func miqt_exec_callback_QThreadPool_MetaObject(self QThreadPool, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QThreadPool{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QThreadPool) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QThreadPool_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QThreadPool) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThreadPool_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThreadPool_Metacast
func miqt_exec_callback_QThreadPool_Metacast(self QThreadPool, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QThreadPool{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
