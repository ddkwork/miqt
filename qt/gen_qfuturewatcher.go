package qt

import (
	"unsafe"
)

type QFutureWatcherBase struct {
	h          uintptr
	isSubclass bool
}

func (this *QFutureWatcherBase) MetaObject() *QMetaObject {
	return newQMetaObject(QFutureWatcherBase_MetaObject(this.h))
}

func (this *QFutureWatcherBase) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFutureWatcherBase_Metacast(this.h, param1_Cstring))
}

func QFutureWatcherBase_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFutureWatcherBase_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFutureWatcherBase) ProgressValue() int {
	return (int)(QFutureWatcherBase_ProgressValue(this.h))
}

func (this *QFutureWatcherBase) ProgressMinimum() int {
	return (int)(QFutureWatcherBase_ProgressMinimum(this.h))
}

func (this *QFutureWatcherBase) ProgressMaximum() int {
	return (int)(QFutureWatcherBase_ProgressMaximum(this.h))
}

func (this *QFutureWatcherBase) ProgressText() string {
	var _ms struct_miqt_string = QFutureWatcherBase_ProgressText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFutureWatcherBase) IsStarted() bool {
	return (bool)(QFutureWatcherBase_IsStarted(this.h))
}

func (this *QFutureWatcherBase) IsFinished() bool {
	return (bool)(QFutureWatcherBase_IsFinished(this.h))
}

func (this *QFutureWatcherBase) IsRunning() bool {
	return (bool)(QFutureWatcherBase_IsRunning(this.h))
}

func (this *QFutureWatcherBase) IsCanceled() bool {
	return (bool)(QFutureWatcherBase_IsCanceled(this.h))
}

func (this *QFutureWatcherBase) IsPaused() bool {
	return (bool)(QFutureWatcherBase_IsPaused(this.h))
}

func (this *QFutureWatcherBase) IsSuspending() bool {
	return (bool)(QFutureWatcherBase_IsSuspending(this.h))
}

func (this *QFutureWatcherBase) IsSuspended() bool {
	return (bool)(QFutureWatcherBase_IsSuspended(this.h))
}

func (this *QFutureWatcherBase) WaitForFinished() {
	QFutureWatcherBase_WaitForFinished(this.h)
}

func (this *QFutureWatcherBase) SetPendingResultsLimit(limit int) {
	QFutureWatcherBase_SetPendingResultsLimit(this.h, (int)(limit))
}

func (this *QFutureWatcherBase) Event(event *QEvent) bool {
	return (bool)(QFutureWatcherBase_Event(this.h, event.cPointer()))
}

func (this *QFutureWatcherBase) Started() {
	QFutureWatcherBase_Started(this.h)
}

func (this *QFutureWatcherBase) OnStarted(slot func()) {
	QFutureWatcherBase_connect_Started(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_Started
func miqt_exec_callback_QFutureWatcherBase_Started(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QFutureWatcherBase) Finished() {
	QFutureWatcherBase_Finished(this.h)
}

func (this *QFutureWatcherBase) OnFinished(slot func()) {
	QFutureWatcherBase_connect_Finished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_Finished
func miqt_exec_callback_QFutureWatcherBase_Finished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QFutureWatcherBase) Canceled() {
	QFutureWatcherBase_Canceled(this.h)
}

func (this *QFutureWatcherBase) OnCanceled(slot func()) {
	QFutureWatcherBase_connect_Canceled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_Canceled
func miqt_exec_callback_QFutureWatcherBase_Canceled(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QFutureWatcherBase) Paused() {
	QFutureWatcherBase_Paused(this.h)
}

func (this *QFutureWatcherBase) OnPaused(slot func()) {
	QFutureWatcherBase_connect_Paused(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_Paused
func miqt_exec_callback_QFutureWatcherBase_Paused(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QFutureWatcherBase) Suspending() {
	QFutureWatcherBase_Suspending(this.h)
}

func (this *QFutureWatcherBase) OnSuspending(slot func()) {
	QFutureWatcherBase_connect_Suspending(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_Suspending
func miqt_exec_callback_QFutureWatcherBase_Suspending(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QFutureWatcherBase) Suspended() {
	QFutureWatcherBase_Suspended(this.h)
}

func (this *QFutureWatcherBase) OnSuspended(slot func()) {
	QFutureWatcherBase_connect_Suspended(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_Suspended
func miqt_exec_callback_QFutureWatcherBase_Suspended(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QFutureWatcherBase) Resumed() {
	QFutureWatcherBase_Resumed(this.h)
}

func (this *QFutureWatcherBase) OnResumed(slot func()) {
	QFutureWatcherBase_connect_Resumed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_Resumed
func miqt_exec_callback_QFutureWatcherBase_Resumed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QFutureWatcherBase) ResultReadyAt(resultIndex int) {
	QFutureWatcherBase_ResultReadyAt(this.h, (int)(resultIndex))
}

func (this *QFutureWatcherBase) OnResultReadyAt(slot func(resultIndex int)) {
	QFutureWatcherBase_connect_ResultReadyAt(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_ResultReadyAt
func miqt_exec_callback_QFutureWatcherBase_ResultReadyAt(cb intptr_t, resultIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(resultIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(resultIndex)

	gofunc(slotval1)
}

func (this *QFutureWatcherBase) ResultsReadyAt(beginIndex int, endIndex int) {
	QFutureWatcherBase_ResultsReadyAt(this.h, (int)(beginIndex), (int)(endIndex))
}

func (this *QFutureWatcherBase) OnResultsReadyAt(slot func(beginIndex int, endIndex int)) {
	QFutureWatcherBase_connect_ResultsReadyAt(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_ResultsReadyAt
func miqt_exec_callback_QFutureWatcherBase_ResultsReadyAt(cb intptr_t, beginIndex int, endIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(beginIndex int, endIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(beginIndex)

	slotval2 := (int)(endIndex)

	gofunc(slotval1, slotval2)
}

func (this *QFutureWatcherBase) ProgressRangeChanged(minimum int, maximum int) {
	QFutureWatcherBase_ProgressRangeChanged(this.h, (int)(minimum), (int)(maximum))
}

func (this *QFutureWatcherBase) OnProgressRangeChanged(slot func(minimum int, maximum int)) {
	QFutureWatcherBase_connect_ProgressRangeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_ProgressRangeChanged
func miqt_exec_callback_QFutureWatcherBase_ProgressRangeChanged(cb intptr_t, minimum int, maximum int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(minimum int, maximum int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(minimum)

	slotval2 := (int)(maximum)

	gofunc(slotval1, slotval2)
}

func (this *QFutureWatcherBase) ProgressValueChanged(progressValue int) {
	QFutureWatcherBase_ProgressValueChanged(this.h, (int)(progressValue))
}

func (this *QFutureWatcherBase) OnProgressValueChanged(slot func(progressValue int)) {
	QFutureWatcherBase_connect_ProgressValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_ProgressValueChanged
func miqt_exec_callback_QFutureWatcherBase_ProgressValueChanged(cb intptr_t, progressValue int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(progressValue int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(progressValue)

	gofunc(slotval1)
}

func (this *QFutureWatcherBase) ProgressTextChanged(progressText string) {
	progressText_ms := struct_miqt_string{}
	progressText_ms.data = CString(progressText)
	progressText_ms.len = size_t(len(progressText))
	defer free(unsafe.Pointer(progressText_ms.data))
	QFutureWatcherBase_ProgressTextChanged(this.h, progressText_ms)
}

func (this *QFutureWatcherBase) OnProgressTextChanged(slot func(progressText string)) {
	QFutureWatcherBase_connect_ProgressTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFutureWatcherBase_ProgressTextChanged
func miqt_exec_callback_QFutureWatcherBase_ProgressTextChanged(cb intptr_t, progressText struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(progressText string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var progressText_ms struct_miqt_string = progressText
	progressText_ret := GoStringN(progressText_ms.data, int(int64(progressText_ms.len)))
	free(unsafe.Pointer(progressText_ms.data))
	slotval1 := progressText_ret

	gofunc(slotval1)
}

func (this *QFutureWatcherBase) Cancel() {
	QFutureWatcherBase_Cancel(this.h)
}

func (this *QFutureWatcherBase) SetSuspended(suspend bool) {
	QFutureWatcherBase_SetSuspended(this.h, (bool)(suspend))
}

func (this *QFutureWatcherBase) Suspend() {
	QFutureWatcherBase_Suspend(this.h)
}

func (this *QFutureWatcherBase) Resume() {
	QFutureWatcherBase_Resume(this.h)
}

func (this *QFutureWatcherBase) ToggleSuspended() {
	QFutureWatcherBase_ToggleSuspended(this.h)
}

func (this *QFutureWatcherBase) SetPaused(paused bool) {
	QFutureWatcherBase_SetPaused(this.h, (bool)(paused))
}

func (this *QFutureWatcherBase) Pause() {
	QFutureWatcherBase_Pause(this.h)
}

func (this *QFutureWatcherBase) TogglePaused() {
	QFutureWatcherBase_TogglePaused(this.h)
}

func QFutureWatcherBase_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFutureWatcherBase_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFutureWatcherBase_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFutureWatcherBase_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
