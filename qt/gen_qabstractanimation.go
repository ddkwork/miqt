package qt

import (
	"unsafe"
)

type QAbstractAnimation__Direction int

const (
	QAbstractAnimation__Forward  QAbstractAnimation__Direction = 0
	QAbstractAnimation__Backward QAbstractAnimation__Direction = 1
)

type QAbstractAnimation__State int

const (
	QAbstractAnimation__Stopped QAbstractAnimation__State = 0
	QAbstractAnimation__Paused  QAbstractAnimation__State = 1
	QAbstractAnimation__Running QAbstractAnimation__State = 2
)

type QAbstractAnimation__DeletionPolicy int

const (
	QAbstractAnimation__KeepWhenStopped   QAbstractAnimation__DeletionPolicy = 0
	QAbstractAnimation__DeleteWhenStopped QAbstractAnimation__DeletionPolicy = 1
)

type QAbstractAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractAnimation constructs a new QAbstractAnimation object.
func NewQAbstractAnimation() *QAbstractAnimation {
	ret := newQAbstractAnimation(QAbstractAnimation_new())
	ret.isSubclass = true
	return ret
}

// NewQAbstractAnimation2 constructs a new QAbstractAnimation object.
func NewQAbstractAnimation2(parent *QObject) *QAbstractAnimation {
	ret := newQAbstractAnimation(QAbstractAnimation_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractAnimation_MetaObject(this.h))
}

func (this *QAbstractAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractAnimation_Metacast(this.h, param1_Cstring))
}

func QAbstractAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractAnimation) State() State {
	xxxxxxxxx
}

func (this *QAbstractAnimation) Group() *QAnimationGroup {
	return newQAnimationGroup(QAbstractAnimation_Group(this.h))
}

func (this *QAbstractAnimation) Direction() Direction {
	xxxxxxxxx
}

func (this *QAbstractAnimation) SetDirection(direction Direction) {
	QAbstractAnimation_SetDirection(this.h, direction)
}

func (this *QAbstractAnimation) CurrentTime() int {
	return (int)(QAbstractAnimation_CurrentTime(this.h))
}

func (this *QAbstractAnimation) CurrentLoopTime() int {
	return (int)(QAbstractAnimation_CurrentLoopTime(this.h))
}

func (this *QAbstractAnimation) LoopCount() int {
	return (int)(QAbstractAnimation_LoopCount(this.h))
}

func (this *QAbstractAnimation) SetLoopCount(loopCount int) {
	QAbstractAnimation_SetLoopCount(this.h, (int)(loopCount))
}

func (this *QAbstractAnimation) CurrentLoop() int {
	return (int)(QAbstractAnimation_CurrentLoop(this.h))
}

func (this *QAbstractAnimation) Duration() int {
	return (int)(QAbstractAnimation_Duration(this.h))
}

func (this *QAbstractAnimation) TotalDuration() int {
	return (int)(QAbstractAnimation_TotalDuration(this.h))
}

func (this *QAbstractAnimation) Finished() {
	QAbstractAnimation_Finished(this.h)
}

func (this *QAbstractAnimation) OnFinished(slot func()) {
	QAbstractAnimation_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAnimation_Finished
func miqt_exec_callback_QAbstractAnimation_Finished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractAnimation) StateChanged(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	QAbstractAnimation_StateChanged(this.h, (int)(newState), (int)(oldState))
}

func (this *QAbstractAnimation) OnStateChanged(slot func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	QAbstractAnimation_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAnimation_StateChanged
func miqt_exec_callback_QAbstractAnimation_StateChanged(cb intptr_t, newState int, oldState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__State)(newState)

	slotval2 := (QAbstractAnimation__State)(oldState)

	gofunc(slotval1, slotval2)
}

func (this *QAbstractAnimation) CurrentLoopChanged(currentLoop int) {
	QAbstractAnimation_CurrentLoopChanged(this.h, (int)(currentLoop))
}

func (this *QAbstractAnimation) OnCurrentLoopChanged(slot func(currentLoop int)) {
	QAbstractAnimation_connect_CurrentLoopChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAnimation_CurrentLoopChanged
func miqt_exec_callback_QAbstractAnimation_CurrentLoopChanged(cb intptr_t, currentLoop int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(currentLoop int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(currentLoop)

	gofunc(slotval1)
}

func (this *QAbstractAnimation) DirectionChanged(param1 QAbstractAnimation__Direction) {
	QAbstractAnimation_DirectionChanged(this.h, (int)(param1))
}

func (this *QAbstractAnimation) OnDirectionChanged(slot func(param1 QAbstractAnimation__Direction)) {
	QAbstractAnimation_connect_DirectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAnimation_DirectionChanged
func miqt_exec_callback_QAbstractAnimation_DirectionChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QAbstractAnimation__Direction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAbstractAnimation__Direction)(param1)

	gofunc(slotval1)
}

func (this *QAbstractAnimation) Start() {
	QAbstractAnimation_Start(this.h)
}

func (this *QAbstractAnimation) Pause() {
	QAbstractAnimation_Pause(this.h)
}

func (this *QAbstractAnimation) Resume() {
	QAbstractAnimation_Resume(this.h)
}

func (this *QAbstractAnimation) SetPaused(paused bool) {
	QAbstractAnimation_SetPaused(this.h, (bool)(paused))
}

func (this *QAbstractAnimation) Stop() {
	QAbstractAnimation_Stop(this.h)
}

func (this *QAbstractAnimation) SetCurrentTime(msecs int) {
	QAbstractAnimation_SetCurrentTime(this.h, (int)(msecs))
}

func QAbstractAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractAnimation) Start1(policy QAbstractAnimation__DeletionPolicy) {
	QAbstractAnimation_Start1(this.h, (int)(policy))
}

func (this *QAbstractAnimation) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractAnimation_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractAnimation) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractAnimation_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAnimation_MetaObject
func miqt_exec_callback_QAbstractAnimation_MetaObject(self QAbstractAnimation, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractAnimation{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractAnimation) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractAnimation_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractAnimation) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractAnimation_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAnimation_Metacast
func miqt_exec_callback_QAbstractAnimation_Metacast(self QAbstractAnimation, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractAnimation{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QAnimationDriver struct {
	h          uintptr
	isSubclass bool
}

// NewQAnimationDriver constructs a new QAnimationDriver object.
func NewQAnimationDriver() *QAnimationDriver {
	ret := newQAnimationDriver(QAnimationDriver_new())
	ret.isSubclass = true
	return ret
}

// NewQAnimationDriver2 constructs a new QAnimationDriver object.
func NewQAnimationDriver2(parent *QObject) *QAnimationDriver {
	ret := newQAnimationDriver(QAnimationDriver_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAnimationDriver) MetaObject() *QMetaObject {
	return newQMetaObject(QAnimationDriver_MetaObject(this.h))
}

func (this *QAnimationDriver) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAnimationDriver_Metacast(this.h, param1_Cstring))
}

func QAnimationDriver_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAnimationDriver_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAnimationDriver) Advance() {
	QAnimationDriver_Advance(this.h)
}

func (this *QAnimationDriver) Install() {
	QAnimationDriver_Install(this.h)
}

func (this *QAnimationDriver) Uninstall() {
	QAnimationDriver_Uninstall(this.h)
}

func (this *QAnimationDriver) IsRunning() bool {
	return (bool)(QAnimationDriver_IsRunning(this.h))
}

func (this *QAnimationDriver) Elapsed() int64 {
	return (int64)(QAnimationDriver_Elapsed(this.h))
}

func (this *QAnimationDriver) Started() {
	QAnimationDriver_Started(this.h)
}

func (this *QAnimationDriver) OnStarted(slot func()) {
	QAnimationDriver_connect_Started(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationDriver_Started
func miqt_exec_callback_QAnimationDriver_Started(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAnimationDriver) Stopped() {
	QAnimationDriver_Stopped(this.h)
}

func (this *QAnimationDriver) OnStopped(slot func()) {
	QAnimationDriver_connect_Stopped(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationDriver_Stopped
func miqt_exec_callback_QAnimationDriver_Stopped(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QAnimationDriver_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAnimationDriver_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAnimationDriver_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAnimationDriver_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAnimationDriver) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAnimationDriver_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAnimationDriver) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAnimationDriver_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationDriver_MetaObject
func miqt_exec_callback_QAnimationDriver_MetaObject(self QAnimationDriver, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAnimationDriver{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAnimationDriver) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAnimationDriver_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAnimationDriver) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAnimationDriver_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAnimationDriver_Metacast
func miqt_exec_callback_QAnimationDriver_Metacast(self QAnimationDriver, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAnimationDriver{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
