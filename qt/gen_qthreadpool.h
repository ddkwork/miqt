#pragma once
#ifndef MIQT_QT_GEN_QTHREADPOOL_H
#define MIQT_QT_GEN_QTHREADPOOL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRunnable QRunnable;
typedef struct QThread QThread;
typedef struct QThreadPool QThreadPool;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QThreadPool* QThreadPool_new();
extern __declspec(dllexport) QThreadPool* QThreadPool_new2(QObject* parent);
extern __declspec(dllexport) void QThreadPool_virtbase(QThreadPool* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QThreadPool_MetaObject(const QThreadPool* self);
extern __declspec(dllexport) void* QThreadPool_Metacast(QThreadPool* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QThreadPool_Tr(const char* s);
extern __declspec(dllexport) QThreadPool* QThreadPool_GlobalInstance();
extern __declspec(dllexport) void QThreadPool_Start(QThreadPool* self, QRunnable* runnable);
extern __declspec(dllexport) bool QThreadPool_TryStart(QThreadPool* self, QRunnable* runnable);
extern __declspec(dllexport) void QThreadPool_StartOnReservedThread(QThreadPool* self, QRunnable* runnable);
extern __declspec(dllexport) int QThreadPool_ExpiryTimeout(const QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_SetExpiryTimeout(QThreadPool* self, int expiryTimeout);
extern __declspec(dllexport) int QThreadPool_MaxThreadCount(const QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_SetMaxThreadCount(QThreadPool* self, int maxThreadCount);
extern __declspec(dllexport) int QThreadPool_ActiveThreadCount(const QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_SetStackSize(QThreadPool* self, unsigned int stackSize);
extern __declspec(dllexport) unsigned int QThreadPool_StackSize(const QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_SetThreadPriority(QThreadPool* self, int priority);
extern __declspec(dllexport) int QThreadPool_ThreadPriority(const QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_ReserveThread(QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_ReleaseThread(QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_SetServiceLevel(QThreadPool* self, int serviceLevel);
extern __declspec(dllexport) int QThreadPool_ServiceLevel(const QThreadPool* self);
extern __declspec(dllexport) bool QThreadPool_WaitForDone(QThreadPool* self, int msecs);
extern __declspec(dllexport) bool QThreadPool_WaitForDone2(QThreadPool* self);
extern __declspec(dllexport) void QThreadPool_Clear(QThreadPool* self);
extern __declspec(dllexport) bool QThreadPool_Contains(const QThreadPool* self, QThread* thread);
extern __declspec(dllexport) bool QThreadPool_TryTake(QThreadPool* self, QRunnable* runnable);
extern __declspec(dllexport) struct miqt_string QThreadPool_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QThreadPool_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QThreadPool_Start2(QThreadPool* self, QRunnable* runnable, int priority);
extern __declspec(dllexport) bool QThreadPool_WaitForDone1(QThreadPool* self, QDeadlineTimer* deadline);
extern __declspec(dllexport) void QThreadPool_override_virtual_Event(void* self, intptr_t slot);
bool QThreadPool_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QThreadPool_override_virtual_EventFilter(void* self, intptr_t slot);
bool QThreadPool_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QThreadPool_override_virtual_TimerEvent(void* self, intptr_t slot);
void QThreadPool_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QThreadPool_override_virtual_ChildEvent(void* self, intptr_t slot);
void QThreadPool_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QThreadPool_override_virtual_CustomEvent(void* self, intptr_t slot);
void QThreadPool_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QThreadPool_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QThreadPool_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QThreadPool_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QThreadPool_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QThreadPool_Delete(QThreadPool* self, bool isSubclass);

} 
