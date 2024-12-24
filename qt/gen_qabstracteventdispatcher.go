package qt

import (
	"unsafe"
)

type QAbstractEventDispatcher struct {
	h          uintptr
	isSubclass bool
}

func (this *QAbstractEventDispatcher) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractEventDispatcher_MetaObject(this.h))
}

func (this *QAbstractEventDispatcher) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractEventDispatcher_Metacast(this.h, param1_Cstring))
}

func QAbstractEventDispatcher_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractEventDispatcher_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractEventDispatcher_Instance() *QAbstractEventDispatcher {
	return newQAbstractEventDispatcher(QAbstractEventDispatcher_Instance())
}

func (this *QAbstractEventDispatcher) ProcessEvents(flags ProcessEventsFlag) bool {
	return (bool)(QAbstractEventDispatcher_ProcessEvents(this.h, (int)(flags)))
}

func (this *QAbstractEventDispatcher) RegisterSocketNotifier(notifier *QSocketNotifier) {
	QAbstractEventDispatcher_RegisterSocketNotifier(this.h, notifier.cPointer())
}

func (this *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier *QSocketNotifier) {
	QAbstractEventDispatcher_UnregisterSocketNotifier(this.h, notifier.cPointer())
}

func (this *QAbstractEventDispatcher) RegisterTimer(interval Duration, timerType TimerType, object *QObject) TimerId {
	return (TimerId)(QAbstractEventDispatcher_RegisterTimer(this.h, interval, (int)(timerType), object.cPointer()))
}

func (this *QAbstractEventDispatcher) RegisterTimer2(interval int64, timerType TimerType, object *QObject) int {
	return (int)(QAbstractEventDispatcher_RegisterTimer2(this.h, (longlong)(interval), (int)(timerType), object.cPointer()))
}

func (this *QAbstractEventDispatcher) RegisterTimer3(timerId int, interval int64, timerType TimerType, object *QObject) {
	QAbstractEventDispatcher_RegisterTimer3(this.h, (int)(timerId), (longlong)(interval), (int)(timerType), object.cPointer())
}

func (this *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	return (bool)(QAbstractEventDispatcher_UnregisterTimer(this.h, (int)(timerId)))
}

func (this *QAbstractEventDispatcher) UnregisterTimers(object *QObject) bool {
	return (bool)(QAbstractEventDispatcher_UnregisterTimers(this.h, object.cPointer()))
}

func (this *QAbstractEventDispatcher) RegisteredTimers(object *QObject) []TimerInfo {
	var _ma struct_miqt_array = QAbstractEventDispatcher_RegisteredTimers(this.h, object.cPointer())
	_ret := make([]TimerInfo, int(_ma.len))
	_outCast := (*[0xffff]TimerInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	return (int)(QAbstractEventDispatcher_RemainingTime(this.h, (int)(timerId)))
}

func (this *QAbstractEventDispatcher) RegisterTimer4(timerId TimerId, interval Duration, timerType TimerType, object *QObject) {
	QAbstractEventDispatcher_RegisterTimer4(this.h, (int)(timerId), interval, (int)(timerType), object.cPointer())
}

func (this *QAbstractEventDispatcher) UnregisterTimerWithTimerId(timerId TimerId) bool {
	return (bool)(QAbstractEventDispatcher_UnregisterTimerWithTimerId(this.h, (int)(timerId)))
}

func (this *QAbstractEventDispatcher) TimersForObject(object *QObject) []TimerInfoV2 {
	var _ma struct_miqt_array = QAbstractEventDispatcher_TimersForObject(this.h, object.cPointer())
	_ret := make([]TimerInfoV2, int(_ma.len))
	_outCast := (*[0xffff]TimerInfoV2)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QAbstractEventDispatcher) RemainingTimeWithTimerId(timerId TimerId) Duration {
	xxxxxxxxx
}

func (this *QAbstractEventDispatcher) WakeUp() {
	QAbstractEventDispatcher_WakeUp(this.h)
}

func (this *QAbstractEventDispatcher) Interrupt() {
	QAbstractEventDispatcher_Interrupt(this.h)
}

func (this *QAbstractEventDispatcher) StartingUp() {
	QAbstractEventDispatcher_StartingUp(this.h)
}

func (this *QAbstractEventDispatcher) ClosingDown() {
	QAbstractEventDispatcher_ClosingDown(this.h)
}

func (this *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	QAbstractEventDispatcher_InstallNativeEventFilter(this.h, filterObj.cPointer())
}

func (this *QAbstractEventDispatcher) RemoveNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	QAbstractEventDispatcher_RemoveNativeEventFilter(this.h, filterObj.cPointer())
}

func (this *QAbstractEventDispatcher) FilterNativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))
	return (bool)(QAbstractEventDispatcher_FilterNativeEvent(this.h, eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))
}

func (this *QAbstractEventDispatcher) AboutToBlock() {
	QAbstractEventDispatcher_AboutToBlock(this.h)
}
func (this *QAbstractEventDispatcher) OnAboutToBlock(slot func()) {
	QAbstractEventDispatcher_connect_AboutToBlock(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractEventDispatcher_AboutToBlock
func miqt_exec_callback_QAbstractEventDispatcher_AboutToBlock(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractEventDispatcher) Awake() {
	QAbstractEventDispatcher_Awake(this.h)
}
func (this *QAbstractEventDispatcher) OnAwake(slot func()) {
	QAbstractEventDispatcher_connect_Awake(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractEventDispatcher_Awake
func miqt_exec_callback_QAbstractEventDispatcher_Awake(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QAbstractEventDispatcher_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractEventDispatcher_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractEventDispatcher_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractEventDispatcher_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractEventDispatcher_Instance1(thread *QThread) *QAbstractEventDispatcher {
	return newQAbstractEventDispatcher(QAbstractEventDispatcher_Instance1(thread.cPointer()))
}

type QAbstractEventDispatcherV2 struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractEventDispatcherV2 constructs a new QAbstractEventDispatcherV2 object.
func NewQAbstractEventDispatcherV2() *QAbstractEventDispatcherV2 {

	ret := newQAbstractEventDispatcherV2(QAbstractEventDispatcherV2_new())
	ret.isSubclass = true
	return ret
}

// NewQAbstractEventDispatcherV22 constructs a new QAbstractEventDispatcherV2 object.
func NewQAbstractEventDispatcherV22(parent *QObject) *QAbstractEventDispatcherV2 {

	ret := newQAbstractEventDispatcherV2(QAbstractEventDispatcherV2_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractEventDispatcherV2) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractEventDispatcherV2_MetaObject(this.h))
}

func (this *QAbstractEventDispatcherV2) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractEventDispatcherV2_Metacast(this.h, param1_Cstring))
}

func QAbstractEventDispatcherV2_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractEventDispatcherV2_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractEventDispatcherV2) RegisterTimer(timerId TimerId, interval Duration, timerType TimerType, object *QObject) {
	QAbstractEventDispatcherV2_RegisterTimer(this.h, (int)(timerId), interval, (int)(timerType), object.cPointer())
}

func (this *QAbstractEventDispatcherV2) UnregisterTimer(timerId TimerId) bool {
	return (bool)(QAbstractEventDispatcherV2_UnregisterTimer(this.h, (int)(timerId)))
}

func (this *QAbstractEventDispatcherV2) TimersForObject(object *QObject) []TimerInfoV2 {
	var _ma struct_miqt_array = QAbstractEventDispatcherV2_TimersForObject(this.h, object.cPointer())
	_ret := make([]TimerInfoV2, int(_ma.len))
	_outCast := (*[0xffff]TimerInfoV2)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QAbstractEventDispatcherV2) RemainingTime(timerId TimerId) Duration {
	xxxxxxxxx
}

func (this *QAbstractEventDispatcherV2) ProcessEventsWithDeadline(flags ProcessEventsFlag, deadline QDeadlineTimer) bool {
	return (bool)(QAbstractEventDispatcherV2_ProcessEventsWithDeadline(this.h, (int)(flags), deadline.cPointer()))
}

func QAbstractEventDispatcherV2_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractEventDispatcherV2_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractEventDispatcherV2_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractEventDispatcherV2_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QAbstractEventDispatcher__TimerInfo struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractEventDispatcher__TimerInfo constructs a new QAbstractEventDispatcher::TimerInfo object.
func NewQAbstractEventDispatcher__TimerInfo(id int, i int, t TimerType) *QAbstractEventDispatcher__TimerInfo {

	ret := newQAbstractEventDispatcher__TimerInfo(QAbstractEventDispatcher__TimerInfo_new((int)(id), (int)(i), (int)(t)))
	ret.isSubclass = true
	return ret
}

// NewQAbstractEventDispatcher__TimerInfo2 constructs a new QAbstractEventDispatcher::TimerInfo object.
func NewQAbstractEventDispatcher__TimerInfo2(param1 *TimerInfo) *QAbstractEventDispatcher__TimerInfo {

	ret := newQAbstractEventDispatcher__TimerInfo(QAbstractEventDispatcher__TimerInfo_new2(param1))
	ret.isSubclass = true
	return ret
}

type QAbstractEventDispatcher__TimerInfoV2 struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractEventDispatcher__TimerInfoV2 constructs a new QAbstractEventDispatcher::TimerInfoV2 object.
func NewQAbstractEventDispatcher__TimerInfoV2() *QAbstractEventDispatcher__TimerInfoV2 {

	ret := newQAbstractEventDispatcher__TimerInfoV2(QAbstractEventDispatcher__TimerInfoV2_new())
	ret.isSubclass = true
	return ret
}

// NewQAbstractEventDispatcher__TimerInfoV22 constructs a new QAbstractEventDispatcher::TimerInfoV2 object.
func NewQAbstractEventDispatcher__TimerInfoV22(param1 *TimerInfoV2) *QAbstractEventDispatcher__TimerInfoV2 {

	ret := newQAbstractEventDispatcher__TimerInfoV2(QAbstractEventDispatcher__TimerInfoV2_new2(param1))
	ret.isSubclass = true
	return ret
}
