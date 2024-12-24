package qt

import (
	"unsafe"
)

type QEventLoop__ProcessEventsFlag int

const (
	QEventLoop__AllEvents              QEventLoop__ProcessEventsFlag = 0
	QEventLoop__ExcludeUserInputEvents QEventLoop__ProcessEventsFlag = 1
	QEventLoop__ExcludeSocketNotifiers QEventLoop__ProcessEventsFlag = 2
	QEventLoop__WaitForMoreEvents      QEventLoop__ProcessEventsFlag = 4
	QEventLoop__X11ExcludeTimers       QEventLoop__ProcessEventsFlag = 8
	QEventLoop__EventLoopExec          QEventLoop__ProcessEventsFlag = 32
	QEventLoop__DialogExec             QEventLoop__ProcessEventsFlag = 64
	QEventLoop__ApplicationExec        QEventLoop__ProcessEventsFlag = 128
)

type QEventLoop struct {
	h          uintptr
	isSubclass bool
}

// NewQEventLoop constructs a new QEventLoop object.
func NewQEventLoop() *QEventLoop {
	g := newQEventLoop(QEventLoop_new())
	g.isSubclass = true
	return g
}

// NewQEventLoop2 constructs a new QEventLoop object.
func NewQEventLoop2(parent *QObject) *QEventLoop {
	g := newQEventLoop(QEventLoop_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QEventLoop) MetaObject() *QMetaObject {
	return newQMetaObject(QEventLoop_MetaObject(this.h))
}

func (this *QEventLoop) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QEventLoop_Metacast(this.h, param1_Cstring))
}

func QEventLoop_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QEventLoop_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QEventLoop) ProcessEvents() bool {
	return (bool)(QEventLoop_ProcessEvents(this.h))
}

func (this *QEventLoop) ProcessEvents2(flags ProcessEventsFlags, maximumTime int) {
	QEventLoop_ProcessEvents2(this.h, flags, (int)(maximumTime))
}

func (this *QEventLoop) ProcessEvents3(flags ProcessEventsFlags, deadline QDeadlineTimer) {
	QEventLoop_ProcessEvents3(this.h, flags, deadline.cPointer())
}

func (this *QEventLoop) Exec() int {
	return (int)(QEventLoop_Exec(this.h))
}

func (this *QEventLoop) IsRunning() bool {
	return (bool)(QEventLoop_IsRunning(this.h))
}

func (this *QEventLoop) WakeUp() {
	QEventLoop_WakeUp(this.h)
}

func (this *QEventLoop) Event(event *QEvent) bool {
	return (bool)(QEventLoop_Event(this.h, event.cPointer()))
}

func (this *QEventLoop) Exit() {
	QEventLoop_Exit(this.h)
}

func (this *QEventLoop) Quit() {
	QEventLoop_Quit(this.h)
}

func QEventLoop_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QEventLoop_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QEventLoop_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QEventLoop_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QEventLoop) ProcessEvents1(flags ProcessEventsFlags) bool {
	return (bool)(QEventLoop_ProcessEvents1(this.h, flags))
}

func (this *QEventLoop) Exec1(flags ProcessEventsFlags) int {
	return (int)(QEventLoop_Exec1(this.h, flags))
}

func (this *QEventLoop) Exit1(returnCode int) {
	QEventLoop_Exit1(this.h, (int)(returnCode))
}

func (this *QEventLoop) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QEventLoop_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QEventLoop) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_MetaObject
func miqt_exec_callback_QEventLoop_MetaObject(self QEventLoop, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QEventLoop{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QEventLoop) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QEventLoop_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QEventLoop) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QEventLoop_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QEventLoop_Metacast
func miqt_exec_callback_QEventLoop_Metacast(self QEventLoop, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QEventLoop{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QEventLoopLocker struct {
	h          uintptr
	isSubclass bool
}

// NewQEventLoopLocker constructs a new QEventLoopLocker object.
func NewQEventLoopLocker() *QEventLoopLocker {
	g := newQEventLoopLocker(QEventLoopLocker_new())
	g.isSubclass = true
	return g
}

// NewQEventLoopLocker2 constructs a new QEventLoopLocker object.
func NewQEventLoopLocker2(loop *QEventLoop) *QEventLoopLocker {
	g := newQEventLoopLocker(QEventLoopLocker_new2(loop.cPointer()))
	g.isSubclass = true
	return g
}

// NewQEventLoopLocker3 constructs a new QEventLoopLocker object.
func NewQEventLoopLocker3(thread *QThread) *QEventLoopLocker {
	g := newQEventLoopLocker(QEventLoopLocker_new3(thread.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QEventLoopLocker) Swap(other *QEventLoopLocker) {
	QEventLoopLocker_Swap(this.h, other.cPointer())
}
