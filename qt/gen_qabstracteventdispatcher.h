#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTEVENTDISPATCHER_H
#define MIQT_QT_GEN_QABSTRACTEVENTDISPATCHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QAbstractEventDispatcher__TimerInfo)
typedef QAbstractEventDispatcher::TimerInfo QAbstractEventDispatcher__TimerInfo;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QAbstractEventDispatcher__TimerInfoV2)
typedef QAbstractEventDispatcher::TimerInfoV2 QAbstractEventDispatcher__TimerInfoV2;
typedef struct QAbstractEventDispatcher QAbstractEventDispatcher;
typedef struct QAbstractEventDispatcher__TimerInfo QAbstractEventDispatcher__TimerInfo;
typedef struct QAbstractEventDispatcher__TimerInfoV2 QAbstractEventDispatcher__TimerInfoV2;
typedef struct QAbstractEventDispatcherV2 QAbstractEventDispatcherV2;
typedef struct QAbstractNativeEventFilter QAbstractNativeEventFilter;
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSocketNotifier QSocketNotifier;
typedef struct QThread QThread;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QAbstractEventDispatcher_virtbase(QAbstractEventDispatcher* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QAbstractEventDispatcher_MetaObject(const QAbstractEventDispatcher* self);
extern __declspec(dllexport) void* QAbstractEventDispatcher_Metacast(QAbstractEventDispatcher* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAbstractEventDispatcher_Tr(const char* s);
extern __declspec(dllexport) QAbstractEventDispatcher* QAbstractEventDispatcher_Instance();
extern __declspec(dllexport) bool QAbstractEventDispatcher_ProcessEvents(QAbstractEventDispatcher* self, int flags);
extern __declspec(dllexport) void QAbstractEventDispatcher_RegisterSocketNotifier(QAbstractEventDispatcher* self, QSocketNotifier* notifier);
extern __declspec(dllexport) void QAbstractEventDispatcher_UnregisterSocketNotifier(QAbstractEventDispatcher* self, QSocketNotifier* notifier);
extern __declspec(dllexport) int QAbstractEventDispatcher_RegisterTimer(QAbstractEventDispatcher* self, Duration interval, int timerType, QObject* object);
extern __declspec(dllexport) int QAbstractEventDispatcher_RegisterTimer2(QAbstractEventDispatcher* self, long long interval, int timerType, QObject* object);
extern __declspec(dllexport) void QAbstractEventDispatcher_RegisterTimer3(QAbstractEventDispatcher* self, int timerId, long long interval, int timerType, QObject* object);
extern __declspec(dllexport) bool QAbstractEventDispatcher_UnregisterTimer(QAbstractEventDispatcher* self, int timerId);
extern __declspec(dllexport) bool QAbstractEventDispatcher_UnregisterTimers(QAbstractEventDispatcher* self, QObject* object);
extern __declspec(dllexport) struct miqt_array /* of TimerInfo */  QAbstractEventDispatcher_RegisteredTimers(const QAbstractEventDispatcher* self, QObject* object);
extern __declspec(dllexport) int QAbstractEventDispatcher_RemainingTime(QAbstractEventDispatcher* self, int timerId);
extern __declspec(dllexport) void QAbstractEventDispatcher_RegisterTimer4(QAbstractEventDispatcher* self, int timerId, Duration interval, int timerType, QObject* object);
extern __declspec(dllexport) bool QAbstractEventDispatcher_UnregisterTimerWithTimerId(QAbstractEventDispatcher* self, int timerId);
extern __declspec(dllexport) struct miqt_array /* of TimerInfoV2 */  QAbstractEventDispatcher_TimersForObject(const QAbstractEventDispatcher* self, QObject* object);
extern __declspec(dllexport) Duration QAbstractEventDispatcher_RemainingTimeWithTimerId(const QAbstractEventDispatcher* self, int timerId);
extern __declspec(dllexport) void QAbstractEventDispatcher_WakeUp(QAbstractEventDispatcher* self);
extern __declspec(dllexport) void QAbstractEventDispatcher_Interrupt(QAbstractEventDispatcher* self);
extern __declspec(dllexport) void QAbstractEventDispatcher_StartingUp(QAbstractEventDispatcher* self);
extern __declspec(dllexport) void QAbstractEventDispatcher_ClosingDown(QAbstractEventDispatcher* self);
extern __declspec(dllexport) void QAbstractEventDispatcher_InstallNativeEventFilter(QAbstractEventDispatcher* self, QAbstractNativeEventFilter* filterObj);
extern __declspec(dllexport) void QAbstractEventDispatcher_RemoveNativeEventFilter(QAbstractEventDispatcher* self, QAbstractNativeEventFilter* filterObj);
extern __declspec(dllexport) bool QAbstractEventDispatcher_FilterNativeEvent(QAbstractEventDispatcher* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) void QAbstractEventDispatcher_AboutToBlock(QAbstractEventDispatcher* self);
void QAbstractEventDispatcher_connect_AboutToBlock(QAbstractEventDispatcher* self, intptr_t slot);
extern __declspec(dllexport) void QAbstractEventDispatcher_Awake(QAbstractEventDispatcher* self);
void QAbstractEventDispatcher_connect_Awake(QAbstractEventDispatcher* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QAbstractEventDispatcher_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAbstractEventDispatcher_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) QAbstractEventDispatcher* QAbstractEventDispatcher_Instance1(QThread* thread);
extern __declspec(dllexport) void QAbstractEventDispatcher_Delete(QAbstractEventDispatcher* self, bool isSubclass);

extern __declspec(dllexport) QAbstractEventDispatcherV2* QAbstractEventDispatcherV2_new();
extern __declspec(dllexport) QAbstractEventDispatcherV2* QAbstractEventDispatcherV2_new2(QObject* parent);
extern __declspec(dllexport) void QAbstractEventDispatcherV2_virtbase(QAbstractEventDispatcherV2* src, QAbstractEventDispatcher** outptr_QAbstractEventDispatcher);
extern __declspec(dllexport) QMetaObject* QAbstractEventDispatcherV2_MetaObject(const QAbstractEventDispatcherV2* self);
extern __declspec(dllexport) void* QAbstractEventDispatcherV2_Metacast(QAbstractEventDispatcherV2* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAbstractEventDispatcherV2_Tr(const char* s);
extern __declspec(dllexport) void QAbstractEventDispatcherV2_RegisterTimer(QAbstractEventDispatcherV2* self, int timerId, Duration interval, int timerType, QObject* object);
extern __declspec(dllexport) bool QAbstractEventDispatcherV2_UnregisterTimer(QAbstractEventDispatcherV2* self, int timerId);
extern __declspec(dllexport) struct miqt_array /* of TimerInfoV2 */  QAbstractEventDispatcherV2_TimersForObject(const QAbstractEventDispatcherV2* self, QObject* object);
extern __declspec(dllexport) Duration QAbstractEventDispatcherV2_RemainingTime(const QAbstractEventDispatcherV2* self, int timerId);
extern __declspec(dllexport) bool QAbstractEventDispatcherV2_ProcessEventsWithDeadline(QAbstractEventDispatcherV2* self, int flags, QDeadlineTimer* deadline);
extern __declspec(dllexport) struct miqt_string QAbstractEventDispatcherV2_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAbstractEventDispatcherV2_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QAbstractEventDispatcherV2_Delete(QAbstractEventDispatcherV2* self, bool isSubclass);

extern __declspec(dllexport) QAbstractEventDispatcher__TimerInfo* QAbstractEventDispatcher__TimerInfo_new(int id, int i, int t);
extern __declspec(dllexport) QAbstractEventDispatcher__TimerInfo* QAbstractEventDispatcher__TimerInfo_new2(const TimerInfo* param1);
extern __declspec(dllexport) void QAbstractEventDispatcher__TimerInfo_Delete(QAbstractEventDispatcher__TimerInfo* self, bool isSubclass);

extern __declspec(dllexport) QAbstractEventDispatcher__TimerInfoV2* QAbstractEventDispatcher__TimerInfoV2_new();
extern __declspec(dllexport) QAbstractEventDispatcher__TimerInfoV2* QAbstractEventDispatcher__TimerInfoV2_new2(const TimerInfoV2* param1);
extern __declspec(dllexport) void QAbstractEventDispatcher__TimerInfoV2_Delete(QAbstractEventDispatcher__TimerInfoV2* self, bool isSubclass);

} 
