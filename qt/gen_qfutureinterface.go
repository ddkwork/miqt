package qt

import (
	"unsafe"
)

type QFutureInterfaceBase__State int

const (
	QFutureInterfaceBase__NoState    QFutureInterfaceBase__State = 0
	QFutureInterfaceBase__Running    QFutureInterfaceBase__State = 1
	QFutureInterfaceBase__Started    QFutureInterfaceBase__State = 2
	QFutureInterfaceBase__Finished   QFutureInterfaceBase__State = 4
	QFutureInterfaceBase__Canceled   QFutureInterfaceBase__State = 8
	QFutureInterfaceBase__Suspending QFutureInterfaceBase__State = 16
	QFutureInterfaceBase__Suspended  QFutureInterfaceBase__State = 32
	QFutureInterfaceBase__Throttled  QFutureInterfaceBase__State = 64
	QFutureInterfaceBase__Pending    QFutureInterfaceBase__State = 128
)

type QFutureInterfaceBase__CancelMode int

const (
	QFutureInterfaceBase__CancelOnly      QFutureInterfaceBase__CancelMode = 0
	QFutureInterfaceBase__CancelAndFinish QFutureInterfaceBase__CancelMode = 1
)

type QFutureInterfaceBase struct {
	h          uintptr
	isSubclass bool
}

// NewQFutureInterfaceBase constructs a new QFutureInterfaceBase object.
func NewQFutureInterfaceBase() *QFutureInterfaceBase {

	ret := newQFutureInterfaceBase(QFutureInterfaceBase_new())
	ret.isSubclass = true
	return ret
}

// NewQFutureInterfaceBase2 constructs a new QFutureInterfaceBase object.
func NewQFutureInterfaceBase2(other *QFutureInterfaceBase) *QFutureInterfaceBase {

	ret := newQFutureInterfaceBase(QFutureInterfaceBase_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFutureInterfaceBase3 constructs a new QFutureInterfaceBase object.
func NewQFutureInterfaceBase3(initialState State) *QFutureInterfaceBase {

	ret := newQFutureInterfaceBase(QFutureInterfaceBase_new3(initialState))
	ret.isSubclass = true
	return ret
}

func (this *QFutureInterfaceBase) OperatorAssign(other *QFutureInterfaceBase) {
	QFutureInterfaceBase_OperatorAssign(this.h, other.cPointer())
}

func (this *QFutureInterfaceBase) ReportStarted() {
	QFutureInterfaceBase_ReportStarted(this.h)
}

func (this *QFutureInterfaceBase) ReportFinished() {
	QFutureInterfaceBase_ReportFinished(this.h)
}

func (this *QFutureInterfaceBase) ReportCanceled() {
	QFutureInterfaceBase_ReportCanceled(this.h)
}

func (this *QFutureInterfaceBase) ReportResultsReady(beginIndex int, endIndex int) {
	QFutureInterfaceBase_ReportResultsReady(this.h, (int)(beginIndex), (int)(endIndex))
}

func (this *QFutureInterfaceBase) SetRunnable(runnable *QRunnable) {
	QFutureInterfaceBase_SetRunnable(this.h, runnable.cPointer())
}

func (this *QFutureInterfaceBase) SetThreadPool(pool *QThreadPool) {
	QFutureInterfaceBase_SetThreadPool(this.h, pool.cPointer())
}

func (this *QFutureInterfaceBase) ThreadPool() *QThreadPool {
	return newQThreadPool(QFutureInterfaceBase_ThreadPool(this.h))
}

func (this *QFutureInterfaceBase) SetFilterMode(enable bool) {
	QFutureInterfaceBase_SetFilterMode(this.h, (bool)(enable))
}

func (this *QFutureInterfaceBase) SetProgressRange(minimum int, maximum int) {
	QFutureInterfaceBase_SetProgressRange(this.h, (int)(minimum), (int)(maximum))
}

func (this *QFutureInterfaceBase) ProgressMinimum() int {
	return (int)(QFutureInterfaceBase_ProgressMinimum(this.h))
}

func (this *QFutureInterfaceBase) ProgressMaximum() int {
	return (int)(QFutureInterfaceBase_ProgressMaximum(this.h))
}

func (this *QFutureInterfaceBase) IsProgressUpdateNeeded() bool {
	return (bool)(QFutureInterfaceBase_IsProgressUpdateNeeded(this.h))
}

func (this *QFutureInterfaceBase) SetProgressValue(progressValue int) {
	QFutureInterfaceBase_SetProgressValue(this.h, (int)(progressValue))
}

func (this *QFutureInterfaceBase) ProgressValue() int {
	return (int)(QFutureInterfaceBase_ProgressValue(this.h))
}

func (this *QFutureInterfaceBase) SetProgressValueAndText(progressValue int, progressText string) {
	progressText_ms := struct_miqt_string{}
	progressText_ms.data = CString(progressText)
	progressText_ms.len = size_t(len(progressText))
	defer free(unsafe.Pointer(progressText_ms.data))
	QFutureInterfaceBase_SetProgressValueAndText(this.h, (int)(progressValue), progressText_ms)
}

func (this *QFutureInterfaceBase) ProgressText() string {
	var _ms struct_miqt_string = QFutureInterfaceBase_ProgressText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFutureInterfaceBase) SetExpectedResultCount(resultCount int) {
	QFutureInterfaceBase_SetExpectedResultCount(this.h, (int)(resultCount))
}

func (this *QFutureInterfaceBase) ExpectedResultCount() int {
	return (int)(QFutureInterfaceBase_ExpectedResultCount(this.h))
}

func (this *QFutureInterfaceBase) ResultCount() int {
	return (int)(QFutureInterfaceBase_ResultCount(this.h))
}

func (this *QFutureInterfaceBase) QueryState(state State) bool {
	return (bool)(QFutureInterfaceBase_QueryState(this.h, state))
}

func (this *QFutureInterfaceBase) IsRunning() bool {
	return (bool)(QFutureInterfaceBase_IsRunning(this.h))
}

func (this *QFutureInterfaceBase) IsStarted() bool {
	return (bool)(QFutureInterfaceBase_IsStarted(this.h))
}

func (this *QFutureInterfaceBase) IsCanceled() bool {
	return (bool)(QFutureInterfaceBase_IsCanceled(this.h))
}

func (this *QFutureInterfaceBase) IsFinished() bool {
	return (bool)(QFutureInterfaceBase_IsFinished(this.h))
}

func (this *QFutureInterfaceBase) IsPaused() bool {
	return (bool)(QFutureInterfaceBase_IsPaused(this.h))
}

func (this *QFutureInterfaceBase) SetPaused(paused bool) {
	QFutureInterfaceBase_SetPaused(this.h, (bool)(paused))
}

func (this *QFutureInterfaceBase) TogglePaused() {
	QFutureInterfaceBase_TogglePaused(this.h)
}

func (this *QFutureInterfaceBase) IsSuspending() bool {
	return (bool)(QFutureInterfaceBase_IsSuspending(this.h))
}

func (this *QFutureInterfaceBase) IsSuspended() bool {
	return (bool)(QFutureInterfaceBase_IsSuspended(this.h))
}

func (this *QFutureInterfaceBase) IsThrottled() bool {
	return (bool)(QFutureInterfaceBase_IsThrottled(this.h))
}

func (this *QFutureInterfaceBase) IsResultReadyAt(index int) bool {
	return (bool)(QFutureInterfaceBase_IsResultReadyAt(this.h, (int)(index)))
}

func (this *QFutureInterfaceBase) IsValid() bool {
	return (bool)(QFutureInterfaceBase_IsValid(this.h))
}

func (this *QFutureInterfaceBase) LoadState() int {
	return (int)(QFutureInterfaceBase_LoadState(this.h))
}

func (this *QFutureInterfaceBase) Cancel() {
	QFutureInterfaceBase_Cancel(this.h)
}

func (this *QFutureInterfaceBase) CancelAndFinish() {
	QFutureInterfaceBase_CancelAndFinish(this.h)
}

func (this *QFutureInterfaceBase) SetSuspended(suspend bool) {
	QFutureInterfaceBase_SetSuspended(this.h, (bool)(suspend))
}

func (this *QFutureInterfaceBase) ToggleSuspended() {
	QFutureInterfaceBase_ToggleSuspended(this.h)
}

func (this *QFutureInterfaceBase) ReportSuspended() {
	QFutureInterfaceBase_ReportSuspended(this.h)
}

func (this *QFutureInterfaceBase) SetThrottled(enable bool) {
	QFutureInterfaceBase_SetThrottled(this.h, (bool)(enable))
}

func (this *QFutureInterfaceBase) WaitForFinished() {
	QFutureInterfaceBase_WaitForFinished(this.h)
}

func (this *QFutureInterfaceBase) WaitForNextResult() bool {
	return (bool)(QFutureInterfaceBase_WaitForNextResult(this.h))
}

func (this *QFutureInterfaceBase) WaitForResult(resultIndex int) {
	QFutureInterfaceBase_WaitForResult(this.h, (int)(resultIndex))
}

func (this *QFutureInterfaceBase) WaitForResume() {
	QFutureInterfaceBase_WaitForResume(this.h)
}

func (this *QFutureInterfaceBase) SuspendIfRequested() {
	QFutureInterfaceBase_SuspendIfRequested(this.h)
}

func (this *QFutureInterfaceBase) Mutex() *QMutex {
	return newQMutex(QFutureInterfaceBase_Mutex(this.h))
}

func (this *QFutureInterfaceBase) HasException() bool {
	return (bool)(QFutureInterfaceBase_HasException(this.h))
}

func (this *QFutureInterfaceBase) OperatorEqual(other *QFutureInterfaceBase) bool {
	return (bool)(QFutureInterfaceBase_OperatorEqual(this.h, other.cPointer()))
}

func (this *QFutureInterfaceBase) OperatorNotEqual(other *QFutureInterfaceBase) bool {
	return (bool)(QFutureInterfaceBase_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QFutureInterfaceBase) Swap(other *QFutureInterfaceBase) {
	QFutureInterfaceBase_Swap(this.h, other.cPointer())
}

func (this *QFutureInterfaceBase) IsChainCanceled() bool {
	return (bool)(QFutureInterfaceBase_IsChainCanceled(this.h))
}
