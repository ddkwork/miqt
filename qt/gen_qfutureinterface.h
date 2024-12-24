#pragma once
#ifndef MIQT_QT_GEN_QFUTUREINTERFACE_H
#define MIQT_QT_GEN_QFUTUREINTERFACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QFutureInterfaceBase;
class QMutex;
class QRunnable;
class QThreadPool;
class _GUID;
class type_info;
#else
typedef struct QFutureInterfaceBase QFutureInterfaceBase;
typedef struct QMutex QMutex;
typedef struct QRunnable QRunnable;
typedef struct QThreadPool QThreadPool;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFutureInterfaceBase* QFutureInterfaceBase_new();
extern __declspec(dllexport) QFutureInterfaceBase* QFutureInterfaceBase_new2(QFutureInterfaceBase* other);
extern __declspec(dllexport) QFutureInterfaceBase* QFutureInterfaceBase_new3(State initialState);
extern __declspec(dllexport) void QFutureInterfaceBase_OperatorAssign(QFutureInterfaceBase* self, QFutureInterfaceBase* other);
extern __declspec(dllexport) void QFutureInterfaceBase_ReportStarted(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_ReportFinished(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_ReportCanceled(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_ReportResultsReady(QFutureInterfaceBase* self, int beginIndex, int endIndex);
extern __declspec(dllexport) void QFutureInterfaceBase_SetRunnable(QFutureInterfaceBase* self, QRunnable* runnable);
extern __declspec(dllexport) void QFutureInterfaceBase_SetThreadPool(QFutureInterfaceBase* self, QThreadPool* pool);
extern __declspec(dllexport) QThreadPool* QFutureInterfaceBase_ThreadPool(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SetFilterMode(QFutureInterfaceBase* self, bool enable);
extern __declspec(dllexport) void QFutureInterfaceBase_SetProgressRange(QFutureInterfaceBase* self, int minimum, int maximum);
extern __declspec(dllexport) int QFutureInterfaceBase_ProgressMinimum(const QFutureInterfaceBase* self);
extern __declspec(dllexport) int QFutureInterfaceBase_ProgressMaximum(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsProgressUpdateNeeded(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SetProgressValue(QFutureInterfaceBase* self, int progressValue);
extern __declspec(dllexport) int QFutureInterfaceBase_ProgressValue(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SetProgressValueAndText(QFutureInterfaceBase* self, int progressValue, struct miqt_string progressText);
extern __declspec(dllexport) struct miqt_string QFutureInterfaceBase_ProgressText(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SetExpectedResultCount(QFutureInterfaceBase* self, int resultCount);
extern __declspec(dllexport) int QFutureInterfaceBase_ExpectedResultCount(QFutureInterfaceBase* self);
extern __declspec(dllexport) int QFutureInterfaceBase_ResultCount(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_QueryState(const QFutureInterfaceBase* self, State state);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsRunning(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsStarted(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsCanceled(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsFinished(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsPaused(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SetPaused(QFutureInterfaceBase* self, bool paused);
extern __declspec(dllexport) void QFutureInterfaceBase_TogglePaused(QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsSuspending(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsSuspended(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsThrottled(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsResultReadyAt(const QFutureInterfaceBase* self, int index);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsValid(const QFutureInterfaceBase* self);
extern __declspec(dllexport) int QFutureInterfaceBase_LoadState(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_Cancel(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_CancelAndFinish(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SetSuspended(QFutureInterfaceBase* self, bool suspend);
extern __declspec(dllexport) void QFutureInterfaceBase_ToggleSuspended(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_ReportSuspended(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SetThrottled(QFutureInterfaceBase* self, bool enable);
extern __declspec(dllexport) void QFutureInterfaceBase_WaitForFinished(QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_WaitForNextResult(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_WaitForResult(QFutureInterfaceBase* self, int resultIndex);
extern __declspec(dllexport) void QFutureInterfaceBase_WaitForResume(QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_SuspendIfRequested(QFutureInterfaceBase* self);
extern __declspec(dllexport) QMutex* QFutureInterfaceBase_Mutex(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_HasException(const QFutureInterfaceBase* self);
extern __declspec(dllexport) bool QFutureInterfaceBase_OperatorEqual(const QFutureInterfaceBase* self, QFutureInterfaceBase* other);
extern __declspec(dllexport) bool QFutureInterfaceBase_OperatorNotEqual(const QFutureInterfaceBase* self, QFutureInterfaceBase* other);
extern __declspec(dllexport) void QFutureInterfaceBase_Swap(QFutureInterfaceBase* self, QFutureInterfaceBase* other);
extern __declspec(dllexport) bool QFutureInterfaceBase_IsChainCanceled(const QFutureInterfaceBase* self);
extern __declspec(dllexport) void QFutureInterfaceBase_Delete(QFutureInterfaceBase* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
