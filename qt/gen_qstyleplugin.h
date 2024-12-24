#pragma once
#ifndef MIQT_QT_GEN_QSTYLEPLUGIN_H
#define MIQT_QT_GEN_QSTYLEPLUGIN_H

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
typedef struct QStyle QStyle;
typedef struct QStylePlugin QStylePlugin;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStylePlugin* QStylePlugin_new();
extern __declspec(dllexport) QStylePlugin* QStylePlugin_new2(QObject* parent);
extern __declspec(dllexport) void QStylePlugin_virtbase(QStylePlugin* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QStylePlugin_MetaObject(const QStylePlugin* self);
extern __declspec(dllexport) void* QStylePlugin_Metacast(QStylePlugin* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QStylePlugin_Tr(const char* s);
extern __declspec(dllexport) QStyle* QStylePlugin_Create(QStylePlugin* self, struct miqt_string key);
extern __declspec(dllexport) struct miqt_string QStylePlugin_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QStylePlugin_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QStylePlugin_override_virtual_Create(void* self, intptr_t slot);
QStyle* QStylePlugin_virtualbase_Create(void* self, struct miqt_string key);
extern __declspec(dllexport) void QStylePlugin_override_virtual_Event(void* self, intptr_t slot);
bool QStylePlugin_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QStylePlugin_override_virtual_EventFilter(void* self, intptr_t slot);
bool QStylePlugin_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QStylePlugin_override_virtual_TimerEvent(void* self, intptr_t slot);
void QStylePlugin_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QStylePlugin_override_virtual_ChildEvent(void* self, intptr_t slot);
void QStylePlugin_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QStylePlugin_override_virtual_CustomEvent(void* self, intptr_t slot);
void QStylePlugin_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QStylePlugin_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QStylePlugin_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QStylePlugin_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QStylePlugin_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QStylePlugin_Delete(QStylePlugin* self, bool isSubclass);

} 
