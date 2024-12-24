#pragma once
#ifndef MIQT_QT_GEN_QOBJECTCLEANUPHANDLER_H
#define MIQT_QT_GEN_QOBJECTCLEANUPHANDLER_H

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
typedef struct QObjectCleanupHandler QObjectCleanupHandler;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QObjectCleanupHandler* QObjectCleanupHandler_new();
extern __declspec(dllexport) void QObjectCleanupHandler_virtbase(QObjectCleanupHandler* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QObjectCleanupHandler_MetaObject(const QObjectCleanupHandler* self);
extern __declspec(dllexport) void* QObjectCleanupHandler_Metacast(QObjectCleanupHandler* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QObjectCleanupHandler_Tr(const char* s);
extern __declspec(dllexport) QObject* QObjectCleanupHandler_Add(QObjectCleanupHandler* self, QObject* object);
extern __declspec(dllexport) void QObjectCleanupHandler_Remove(QObjectCleanupHandler* self, QObject* object);
extern __declspec(dllexport) bool QObjectCleanupHandler_IsEmpty(const QObjectCleanupHandler* self);
extern __declspec(dllexport) void QObjectCleanupHandler_Clear(QObjectCleanupHandler* self);
extern __declspec(dllexport) struct miqt_string QObjectCleanupHandler_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QObjectCleanupHandler_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QObjectCleanupHandler_override_virtual_Event(void* self, intptr_t slot);
bool QObjectCleanupHandler_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QObjectCleanupHandler_override_virtual_EventFilter(void* self, intptr_t slot);
bool QObjectCleanupHandler_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QObjectCleanupHandler_override_virtual_TimerEvent(void* self, intptr_t slot);
void QObjectCleanupHandler_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QObjectCleanupHandler_override_virtual_ChildEvent(void* self, intptr_t slot);
void QObjectCleanupHandler_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QObjectCleanupHandler_override_virtual_CustomEvent(void* self, intptr_t slot);
void QObjectCleanupHandler_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QObjectCleanupHandler_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QObjectCleanupHandler_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QObjectCleanupHandler_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QObjectCleanupHandler_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QObjectCleanupHandler_Delete(QObjectCleanupHandler* self, bool isSubclass);

} 
