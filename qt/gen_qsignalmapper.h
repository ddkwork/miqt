#pragma once
#ifndef MIQT_QT_GEN_QSIGNALMAPPER_H
#define MIQT_QT_GEN_QSIGNALMAPPER_H

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
class QMetaMethod;
class QMetaObject;
class QObject;
class QSignalMapper;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSignalMapper QSignalMapper;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSignalMapper* QSignalMapper_new();
extern __declspec(dllexport) QSignalMapper* QSignalMapper_new2(QObject* parent);
extern __declspec(dllexport) void QSignalMapper_virtbase(QSignalMapper* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSignalMapper_MetaObject(const QSignalMapper* self);
extern __declspec(dllexport) void* QSignalMapper_Metacast(QSignalMapper* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSignalMapper_Tr(const char* s);
extern __declspec(dllexport) void QSignalMapper_SetMapping(QSignalMapper* self, QObject* sender, int id);
extern __declspec(dllexport) void QSignalMapper_SetMapping2(QSignalMapper* self, QObject* sender, struct miqt_string text);
extern __declspec(dllexport) void QSignalMapper_SetMapping3(QSignalMapper* self, QObject* sender, QObject* object);
extern __declspec(dllexport) void QSignalMapper_RemoveMappings(QSignalMapper* self, QObject* sender);
extern __declspec(dllexport) QObject* QSignalMapper_Mapping(const QSignalMapper* self, int id);
extern __declspec(dllexport) QObject* QSignalMapper_MappingWithText(const QSignalMapper* self, struct miqt_string text);
extern __declspec(dllexport) QObject* QSignalMapper_MappingWithObject(const QSignalMapper* self, QObject* object);
extern __declspec(dllexport) void QSignalMapper_MappedInt(QSignalMapper* self, int param1);
void QSignalMapper_connect_MappedInt(QSignalMapper* self, intptr_t slot);
extern __declspec(dllexport) void QSignalMapper_MappedString(QSignalMapper* self, struct miqt_string param1);
void QSignalMapper_connect_MappedString(QSignalMapper* self, intptr_t slot);
extern __declspec(dllexport) void QSignalMapper_MappedObject(QSignalMapper* self, QObject* param1);
void QSignalMapper_connect_MappedObject(QSignalMapper* self, intptr_t slot);
extern __declspec(dllexport) void QSignalMapper_Map(QSignalMapper* self);
extern __declspec(dllexport) void QSignalMapper_MapWithSender(QSignalMapper* self, QObject* sender);
extern __declspec(dllexport) struct miqt_string QSignalMapper_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSignalMapper_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSignalMapper_override_virtual_Event(void* self, intptr_t slot);
bool QSignalMapper_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSignalMapper_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSignalMapper_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QSignalMapper_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSignalMapper_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QSignalMapper_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSignalMapper_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QSignalMapper_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSignalMapper_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSignalMapper_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSignalMapper_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSignalMapper_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSignalMapper_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSignalMapper_Delete(QSignalMapper* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
