#pragma once
#ifndef MIQT_QT_GEN_QWINEVENTNOTIFIER_H
#define MIQT_QT_GEN_QWINEVENTNOTIFIER_H

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
typedef struct QTimerEvent QTimerEvent;
typedef struct QWinEventNotifier QWinEventNotifier;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QWinEventNotifier* QWinEventNotifier_new();
extern __declspec(dllexport) QWinEventNotifier* QWinEventNotifier_new2(HANDLE hEvent);
extern __declspec(dllexport) QWinEventNotifier* QWinEventNotifier_new3(QObject* parent);
extern __declspec(dllexport) QWinEventNotifier* QWinEventNotifier_new4(HANDLE hEvent, QObject* parent);
extern __declspec(dllexport) void QWinEventNotifier_virtbase(QWinEventNotifier* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QWinEventNotifier_MetaObject(const QWinEventNotifier* self);
extern __declspec(dllexport) void* QWinEventNotifier_Metacast(QWinEventNotifier* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QWinEventNotifier_Tr(const char* s);
extern __declspec(dllexport) void QWinEventNotifier_SetHandle(QWinEventNotifier* self, HANDLE hEvent);
extern __declspec(dllexport) HANDLE QWinEventNotifier_Handle(const QWinEventNotifier* self);
extern __declspec(dllexport) bool QWinEventNotifier_IsEnabled(const QWinEventNotifier* self);
extern __declspec(dllexport) void QWinEventNotifier_SetEnabled(QWinEventNotifier* self, bool enable);
extern __declspec(dllexport) bool QWinEventNotifier_Event(QWinEventNotifier* self, QEvent* e);
extern __declspec(dllexport) struct miqt_string QWinEventNotifier_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QWinEventNotifier_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QWinEventNotifier_override_virtual_Event(void* self, intptr_t slot);
bool QWinEventNotifier_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QWinEventNotifier_override_virtual_EventFilter(void* self, intptr_t slot);
bool QWinEventNotifier_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QWinEventNotifier_override_virtual_TimerEvent(void* self, intptr_t slot);
void QWinEventNotifier_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QWinEventNotifier_override_virtual_ChildEvent(void* self, intptr_t slot);
void QWinEventNotifier_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QWinEventNotifier_override_virtual_CustomEvent(void* self, intptr_t slot);
void QWinEventNotifier_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QWinEventNotifier_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QWinEventNotifier_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QWinEventNotifier_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QWinEventNotifier_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QWinEventNotifier_Delete(QWinEventNotifier* self, bool isSubclass);

} 
