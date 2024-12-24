package qt

import (
	"unsafe"
)

type QThread__Priority int

const (
	QThread__IdlePriority         QThread__Priority = 0
	QThread__LowestPriority       QThread__Priority = 1
	QThread__LowPriority          QThread__Priority = 2
	QThread__NormalPriority       QThread__Priority = 3
	QThread__HighPriority         QThread__Priority = 4
	QThread__HighestPriority      QThread__Priority = 5
	QThread__TimeCriticalPriority QThread__Priority = 6
	QThread__InheritPriority      QThread__Priority = 7
)

type QThread__QualityOfService int

const (
	QThread__Auto QThread__QualityOfService = 0
	QThread__High QThread__QualityOfService = 1
	QThread__Eco  QThread__QualityOfService = 2
)

type QThread struct {
	h          uintptr
	isSubclass bool
}

// NewQThread constructs a new QThread object.
func NewQThread() *QThread {
	g := newQThread(QThread_new())
	g.isSubclass = true
	return g
}

// NewQThread2 constructs a new QThread object.
func NewQThread2(parent *QObject) *QThread {
	g := newQThread(QThread_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QThread) MetaObject() *QMetaObject {
	return newQMetaObject(QThread_MetaObject(this.h))
}

func (this *QThread) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QThread_Metacast(this.h, param1_Cstring))
}

func QThread_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QThread_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThread_CurrentThreadId() unsafe.Pointer {
	return (unsafe.Pointer)(QThread_CurrentThreadId())
}

func QThread_CurrentThread() *QThread {
	return newQThread(QThread_CurrentThread())
}

func QThread_IsMainThread() bool {
	return (bool)(QThread_IsMainThread())
}

func QThread_IdealThreadCount() int {
	return (int)(QThread_IdealThreadCount())
}

func QThread_YieldCurrentThread() {
	QThread_YieldCurrentThread()
}

func (this *QThread) SetPriority(priority Priority) {
	QThread_SetPriority(this.h, priority)
}

func (this *QThread) Priority() Priority {
	xxxxxxxxx
}

func (this *QThread) IsFinished() bool {
	return (bool)(QThread_IsFinished(this.h))
}

func (this *QThread) IsRunning() bool {
	return (bool)(QThread_IsRunning(this.h))
}

func (this *QThread) RequestInterruption() {
	QThread_RequestInterruption(this.h)
}

func (this *QThread) IsInterruptionRequested() bool {
	return (bool)(QThread_IsInterruptionRequested(this.h))
}

func (this *QThread) SetStackSize(stackSize uint) {
	QThread_SetStackSize(this.h, (uint)(stackSize))
}

func (this *QThread) StackSize() uint {
	return (uint)(QThread_StackSize(this.h))
}

func (this *QThread) EventDispatcher() *QAbstractEventDispatcher {
	return newQAbstractEventDispatcher(QThread_EventDispatcher(this.h))
}

func (this *QThread) SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher) {
	QThread_SetEventDispatcher(this.h, eventDispatcher.cPointer())
}

func (this *QThread) Event(event *QEvent) bool {
	return (bool)(QThread_Event(this.h, event.cPointer()))
}

func (this *QThread) LoopLevel() int {
	return (int)(QThread_LoopLevel(this.h))
}

func (this *QThread) IsCurrentThread() bool {
	return (bool)(QThread_IsCurrentThread(this.h))
}

func (this *QThread) SetServiceLevel(serviceLevel QualityOfService) {
	QThread_SetServiceLevel(this.h, serviceLevel)
}

func (this *QThread) ServiceLevel() QualityOfService {
	xxxxxxxxx
}

func (this *QThread) Start() {
	QThread_Start(this.h)
}

func (this *QThread) Terminate() {
	QThread_Terminate(this.h)
}

func (this *QThread) Exit() {
	QThread_Exit(this.h)
}

func (this *QThread) Quit() {
	QThread_Quit(this.h)
}

func (this *QThread) Wait() bool {
	return (bool)(QThread_Wait(this.h))
}

func (this *QThread) WaitWithTime(time uint32) bool {
	return (bool)(QThread_WaitWithTime(this.h, (ulong)(time)))
}

func QThread_Sleep(param1 uint32) {
	QThread_Sleep((ulong)(param1))
}

func QThread_Msleep(param1 uint32) {
	QThread_Msleep((ulong)(param1))
}

func QThread_Usleep(param1 uint32) {
	QThread_Usleep((ulong)(param1))
}

func QThread_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThread_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QThread_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QThread_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QThread) Start1(param1 Priority) {
	QThread_Start1(this.h, param1)
}

func (this *QThread) Exit1(retcode int) {
	QThread_Exit1(this.h, (int)(retcode))
}

func (this *QThread) Wait1(deadline QDeadlineTimer) bool {
	return (bool)(QThread_Wait1(this.h, deadline.cPointer()))
}

func (this *QThread) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QThread_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QThread) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_MetaObject
func miqt_exec_callback_QThread_MetaObject(self QThread, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QThread{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QThread) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QThread_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QThread) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QThread_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QThread_Metacast
func miqt_exec_callback_QThread_Metacast(self QThread, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QThread{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
