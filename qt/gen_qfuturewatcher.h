#pragma once
#ifndef MIQT_QT_GEN_QFUTUREWATCHER_H
#define MIQT_QT_GEN_QFUTUREWATCHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QFutureWatcherBase QFutureWatcherBase;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QFutureWatcherBase_virtbase(QFutureWatcherBase* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QFutureWatcherBase_MetaObject(const QFutureWatcherBase* self);
extern __declspec(dllexport) void* QFutureWatcherBase_Metacast(QFutureWatcherBase* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QFutureWatcherBase_Tr(const char* s);
extern __declspec(dllexport) int QFutureWatcherBase_ProgressValue(const QFutureWatcherBase* self);
extern __declspec(dllexport) int QFutureWatcherBase_ProgressMinimum(const QFutureWatcherBase* self);
extern __declspec(dllexport) int QFutureWatcherBase_ProgressMaximum(const QFutureWatcherBase* self);
extern __declspec(dllexport) struct miqt_string QFutureWatcherBase_ProgressText(const QFutureWatcherBase* self);
extern __declspec(dllexport) bool QFutureWatcherBase_IsStarted(const QFutureWatcherBase* self);
extern __declspec(dllexport) bool QFutureWatcherBase_IsFinished(const QFutureWatcherBase* self);
extern __declspec(dllexport) bool QFutureWatcherBase_IsRunning(const QFutureWatcherBase* self);
extern __declspec(dllexport) bool QFutureWatcherBase_IsCanceled(const QFutureWatcherBase* self);
extern __declspec(dllexport) bool QFutureWatcherBase_IsPaused(const QFutureWatcherBase* self);
extern __declspec(dllexport) bool QFutureWatcherBase_IsSuspending(const QFutureWatcherBase* self);
extern __declspec(dllexport) bool QFutureWatcherBase_IsSuspended(const QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_WaitForFinished(QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_SetPendingResultsLimit(QFutureWatcherBase* self, int limit);
extern __declspec(dllexport) bool QFutureWatcherBase_Event(QFutureWatcherBase* self, QEvent* event);
extern __declspec(dllexport) void QFutureWatcherBase_Started(QFutureWatcherBase* self);
void QFutureWatcherBase_connect_Started(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_Finished(QFutureWatcherBase* self);
void QFutureWatcherBase_connect_Finished(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_Canceled(QFutureWatcherBase* self);
void QFutureWatcherBase_connect_Canceled(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_Paused(QFutureWatcherBase* self);
void QFutureWatcherBase_connect_Paused(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_Suspending(QFutureWatcherBase* self);
void QFutureWatcherBase_connect_Suspending(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_Suspended(QFutureWatcherBase* self);
void QFutureWatcherBase_connect_Suspended(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_Resumed(QFutureWatcherBase* self);
void QFutureWatcherBase_connect_Resumed(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_ResultReadyAt(QFutureWatcherBase* self, int resultIndex);
void QFutureWatcherBase_connect_ResultReadyAt(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_ResultsReadyAt(QFutureWatcherBase* self, int beginIndex, int endIndex);
void QFutureWatcherBase_connect_ResultsReadyAt(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_ProgressRangeChanged(QFutureWatcherBase* self, int minimum, int maximum);
void QFutureWatcherBase_connect_ProgressRangeChanged(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_ProgressValueChanged(QFutureWatcherBase* self, int progressValue);
void QFutureWatcherBase_connect_ProgressValueChanged(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_ProgressTextChanged(QFutureWatcherBase* self, struct miqt_string progressText);
void QFutureWatcherBase_connect_ProgressTextChanged(QFutureWatcherBase* self, intptr_t slot);
extern __declspec(dllexport) void QFutureWatcherBase_Cancel(QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_SetSuspended(QFutureWatcherBase* self, bool suspend);
extern __declspec(dllexport) void QFutureWatcherBase_Suspend(QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_Resume(QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_ToggleSuspended(QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_SetPaused(QFutureWatcherBase* self, bool paused);
extern __declspec(dllexport) void QFutureWatcherBase_Pause(QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_TogglePaused(QFutureWatcherBase* self);
extern __declspec(dllexport) void QFutureWatcherBase_ConnectNotify(QFutureWatcherBase* self, QMetaMethod* signal);
extern __declspec(dllexport) void QFutureWatcherBase_DisconnectNotify(QFutureWatcherBase* self, QMetaMethod* signal);
extern __declspec(dllexport) struct miqt_string QFutureWatcherBase_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QFutureWatcherBase_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QFutureWatcherBase_Delete(QFutureWatcherBase* self, bool isSubclass);

} 
