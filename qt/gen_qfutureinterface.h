#pragma once
#ifndef MIQT_QT_GEN_QFUTUREINTERFACE_H
#define MIQT_QT_GEN_QFUTUREINTERFACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QFutureInterfaceBase;
class QMutex;
class QRunnable;
class QThreadPool;
#else
typedef struct QFutureInterfaceBase QFutureInterfaceBase;
typedef struct QMutex QMutex;
typedef struct QRunnable QRunnable;
typedef struct QThreadPool QThreadPool;
#endif

void QFutureInterfaceBase_new(QFutureInterfaceBase** outptr_QFutureInterfaceBase);
void QFutureInterfaceBase_new2(QFutureInterfaceBase* other, QFutureInterfaceBase** outptr_QFutureInterfaceBase);
void QFutureInterfaceBase_new3(int initialState, QFutureInterfaceBase** outptr_QFutureInterfaceBase);
void QFutureInterfaceBase_ReportStarted(QFutureInterfaceBase* self);
void QFutureInterfaceBase_ReportFinished(QFutureInterfaceBase* self);
void QFutureInterfaceBase_ReportCanceled(QFutureInterfaceBase* self);
void QFutureInterfaceBase_ReportResultsReady(QFutureInterfaceBase* self, int beginIndex, int endIndex);
void QFutureInterfaceBase_SetRunnable(QFutureInterfaceBase* self, QRunnable* runnable);
void QFutureInterfaceBase_SetThreadPool(QFutureInterfaceBase* self, QThreadPool* pool);
void QFutureInterfaceBase_SetFilterMode(QFutureInterfaceBase* self, bool enable);
void QFutureInterfaceBase_SetProgressRange(QFutureInterfaceBase* self, int minimum, int maximum);
int QFutureInterfaceBase_ProgressMinimum(const QFutureInterfaceBase* self);
int QFutureInterfaceBase_ProgressMaximum(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_IsProgressUpdateNeeded(const QFutureInterfaceBase* self);
void QFutureInterfaceBase_SetProgressValue(QFutureInterfaceBase* self, int progressValue);
int QFutureInterfaceBase_ProgressValue(const QFutureInterfaceBase* self);
void QFutureInterfaceBase_SetProgressValueAndText(QFutureInterfaceBase* self, int progressValue, struct miqt_string progressText);
struct miqt_string QFutureInterfaceBase_ProgressText(const QFutureInterfaceBase* self);
void QFutureInterfaceBase_SetExpectedResultCount(QFutureInterfaceBase* self, int resultCount);
int QFutureInterfaceBase_ExpectedResultCount(QFutureInterfaceBase* self);
int QFutureInterfaceBase_ResultCount(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_QueryState(const QFutureInterfaceBase* self, int state);
bool QFutureInterfaceBase_IsRunning(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_IsStarted(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_IsCanceled(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_IsFinished(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_IsPaused(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_IsThrottled(const QFutureInterfaceBase* self);
bool QFutureInterfaceBase_IsResultReadyAt(const QFutureInterfaceBase* self, int index);
void QFutureInterfaceBase_Cancel(QFutureInterfaceBase* self);
void QFutureInterfaceBase_SetPaused(QFutureInterfaceBase* self, bool paused);
void QFutureInterfaceBase_TogglePaused(QFutureInterfaceBase* self);
void QFutureInterfaceBase_SetThrottled(QFutureInterfaceBase* self, bool enable);
void QFutureInterfaceBase_WaitForFinished(QFutureInterfaceBase* self);
bool QFutureInterfaceBase_WaitForNextResult(QFutureInterfaceBase* self);
void QFutureInterfaceBase_WaitForResult(QFutureInterfaceBase* self, int resultIndex);
void QFutureInterfaceBase_WaitForResume(QFutureInterfaceBase* self);
QMutex* QFutureInterfaceBase_Mutex(const QFutureInterfaceBase* self);
QMutex* QFutureInterfaceBase_MutexWithInt(const QFutureInterfaceBase* self, int param1);
bool QFutureInterfaceBase_OperatorEqual(const QFutureInterfaceBase* self, QFutureInterfaceBase* other);
bool QFutureInterfaceBase_OperatorNotEqual(const QFutureInterfaceBase* self, QFutureInterfaceBase* other);
void QFutureInterfaceBase_OperatorAssign(QFutureInterfaceBase* self, QFutureInterfaceBase* other);
void QFutureInterfaceBase_Delete(QFutureInterfaceBase* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
