#pragma once
#ifndef MIQT_QT_GEN_QTIMER_H
#define MIQT_QT_GEN_QTIMER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimer QTimer;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTimer* QTimer_new();
extern __declspec(dllexport) QTimer* QTimer_new2(QObject* parent);
extern __declspec(dllexport) void QTimer_virtbase(QTimer* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QTimer_MetaObject(const QTimer* self);
extern __declspec(dllexport) void* QTimer_Metacast(QTimer* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTimer_Tr(const char* s);
extern __declspec(dllexport) bool QTimer_IsActive(const QTimer* self);
extern __declspec(dllexport) int QTimer_TimerId(const QTimer* self);
extern __declspec(dllexport) int QTimer_Id(const QTimer* self);
extern __declspec(dllexport) void QTimer_SetInterval(QTimer* self, int msec);
extern __declspec(dllexport) int QTimer_Interval(const QTimer* self);
extern __declspec(dllexport) int QTimer_RemainingTime(const QTimer* self);
extern __declspec(dllexport) void QTimer_SetTimerType(QTimer* self, int atype);
extern __declspec(dllexport) int QTimer_TimerType(const QTimer* self);
extern __declspec(dllexport) void QTimer_SetSingleShot(QTimer* self, bool singleShot);
extern __declspec(dllexport) bool QTimer_IsSingleShot(const QTimer* self);
extern __declspec(dllexport) void QTimer_Start(QTimer* self, int msec);
extern __declspec(dllexport) void QTimer_Start2(QTimer* self);
extern __declspec(dllexport) void QTimer_Stop(QTimer* self);
extern __declspec(dllexport) void QTimer_TimerEvent(QTimer* self, QTimerEvent* param1);
extern __declspec(dllexport) struct miqt_string QTimer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTimer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QTimer_override_virtual_TimerEvent(void* self, intptr_t slot);
void QTimer_virtualbase_TimerEvent(void* self, QTimerEvent* param1);
extern __declspec(dllexport) void QTimer_override_virtual_Event(void* self, intptr_t slot);
bool QTimer_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QTimer_override_virtual_EventFilter(void* self, intptr_t slot);
bool QTimer_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QTimer_override_virtual_ChildEvent(void* self, intptr_t slot);
void QTimer_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QTimer_override_virtual_CustomEvent(void* self, intptr_t slot);
void QTimer_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QTimer_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QTimer_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QTimer_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QTimer_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QTimer_Delete(QTimer* self, bool isSubclass);

} 
