#pragma once
#ifndef MIQT_QT_GEN_QTHREAD_H
#define MIQT_QT_GEN_QTHREAD_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractEventDispatcher QAbstractEventDispatcher;
typedef struct QChildEvent QChildEvent;
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QThread QThread;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QThread* QThread_new();
extern __declspec(dllexport) QThread* QThread_new2(QObject* parent);
extern __declspec(dllexport) void QThread_virtbase(QThread* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QThread_MetaObject(const QThread* self);
extern __declspec(dllexport) void* QThread_Metacast(QThread* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QThread_Tr(const char* s);
extern __declspec(dllexport) void* QThread_CurrentThreadId();
extern __declspec(dllexport) QThread* QThread_CurrentThread();
extern __declspec(dllexport) bool QThread_IsMainThread();
extern __declspec(dllexport) int QThread_IdealThreadCount();
extern __declspec(dllexport) void QThread_YieldCurrentThread();
extern __declspec(dllexport) void QThread_SetPriority(QThread* self, Priority priority);
extern __declspec(dllexport) Priority QThread_Priority(const QThread* self);
extern __declspec(dllexport) bool QThread_IsFinished(const QThread* self);
extern __declspec(dllexport) bool QThread_IsRunning(const QThread* self);
extern __declspec(dllexport) void QThread_RequestInterruption(QThread* self);
extern __declspec(dllexport) bool QThread_IsInterruptionRequested(const QThread* self);
extern __declspec(dllexport) void QThread_SetStackSize(QThread* self, unsigned int stackSize);
extern __declspec(dllexport) unsigned int QThread_StackSize(const QThread* self);
extern __declspec(dllexport) QAbstractEventDispatcher* QThread_EventDispatcher(const QThread* self);
extern __declspec(dllexport) void QThread_SetEventDispatcher(QThread* self, QAbstractEventDispatcher* eventDispatcher);
extern __declspec(dllexport) bool QThread_Event(QThread* self, QEvent* event);
extern __declspec(dllexport) int QThread_LoopLevel(const QThread* self);
extern __declspec(dllexport) bool QThread_IsCurrentThread(const QThread* self);
extern __declspec(dllexport) void QThread_SetServiceLevel(QThread* self, QualityOfService serviceLevel);
extern __declspec(dllexport) QualityOfService QThread_ServiceLevel(const QThread* self);
extern __declspec(dllexport) void QThread_Start(QThread* self);
extern __declspec(dllexport) void QThread_Terminate(QThread* self);
extern __declspec(dllexport) void QThread_Exit(QThread* self);
extern __declspec(dllexport) void QThread_Quit(QThread* self);
extern __declspec(dllexport) bool QThread_Wait(QThread* self);
extern __declspec(dllexport) bool QThread_WaitWithTime(QThread* self, unsigned long time);
extern __declspec(dllexport) void QThread_Sleep(unsigned long param1);
extern __declspec(dllexport) void QThread_Msleep(unsigned long param1);
extern __declspec(dllexport) void QThread_Usleep(unsigned long param1);
extern __declspec(dllexport) void QThread_Run(QThread* self);
extern __declspec(dllexport) struct miqt_string QThread_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QThread_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QThread_Start1(QThread* self, Priority param1);
extern __declspec(dllexport) void QThread_Exit1(QThread* self, int retcode);
extern __declspec(dllexport) bool QThread_Wait1(QThread* self, QDeadlineTimer* deadline);
extern __declspec(dllexport) void QThread_override_virtual_Event(void* self, intptr_t slot);
bool QThread_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QThread_override_virtual_Run(void* self, intptr_t slot);
void QThread_virtualbase_Run(void* self);
extern __declspec(dllexport) void QThread_override_virtual_EventFilter(void* self, intptr_t slot);
bool QThread_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QThread_override_virtual_TimerEvent(void* self, intptr_t slot);
void QThread_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QThread_override_virtual_ChildEvent(void* self, intptr_t slot);
void QThread_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QThread_override_virtual_CustomEvent(void* self, intptr_t slot);
void QThread_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QThread_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QThread_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QThread_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QThread_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QThread_Delete(QThread* self, bool isSubclass);

} 
