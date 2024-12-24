#pragma once
#ifndef MIQT_QT_GEN_QICONENGINEPLUGIN_H
#define MIQT_QT_GEN_QICONENGINEPLUGIN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QIconEngine;
class QIconEnginePlugin;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QIconEngine QIconEngine;
typedef struct QIconEnginePlugin QIconEnginePlugin;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QIconEnginePlugin* QIconEnginePlugin_new();
extern __declspec(dllexport) QIconEnginePlugin* QIconEnginePlugin_new2(QObject* parent);
extern __declspec(dllexport) void QIconEnginePlugin_virtbase(QIconEnginePlugin* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QIconEnginePlugin_MetaObject(const QIconEnginePlugin* self);
extern __declspec(dllexport) void* QIconEnginePlugin_Metacast(QIconEnginePlugin* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QIconEnginePlugin_Tr(const char* s);
extern __declspec(dllexport) QIconEngine* QIconEnginePlugin_Create(QIconEnginePlugin* self, struct miqt_string filename);
extern __declspec(dllexport) struct miqt_string QIconEnginePlugin_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QIconEnginePlugin_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_Create(void* self, intptr_t slot);
QIconEngine* QIconEnginePlugin_virtualbase_Create(void* self, struct miqt_string filename);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_Event(void* self, intptr_t slot);
bool QIconEnginePlugin_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_EventFilter(void* self, intptr_t slot);
bool QIconEnginePlugin_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_TimerEvent(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_ChildEvent(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_CustomEvent(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QIconEnginePlugin_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QIconEnginePlugin_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QIconEnginePlugin_Delete(QIconEnginePlugin* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
