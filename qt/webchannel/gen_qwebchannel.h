#pragma once
#ifndef MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNEL_H
#define MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTimerEvent;
class QWebChannel;
class QWebChannelAbstractTransport;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWebChannel QWebChannel;
typedef struct QWebChannelAbstractTransport QWebChannelAbstractTransport;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QWebChannel* QWebChannel_new();
extern __declspec(dllexport) QWebChannel* QWebChannel_new2(QObject* parent);
extern __declspec(dllexport) void QWebChannel_virtbase(QWebChannel* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QWebChannel_MetaObject(const QWebChannel* self);
extern __declspec(dllexport) void* QWebChannel_Metacast(QWebChannel* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QWebChannel_Tr(const char* s);
extern __declspec(dllexport) void QWebChannel_RegisterObjects(QWebChannel* self, struct miqt_map /* of struct miqt_string to QObject* */  objects);
extern __declspec(dllexport) struct miqt_map /* of struct miqt_string to QObject* */  QWebChannel_RegisteredObjects(const QWebChannel* self);
extern __declspec(dllexport) void QWebChannel_RegisterObject(QWebChannel* self, struct miqt_string id, QObject* object);
extern __declspec(dllexport) void QWebChannel_DeregisterObject(QWebChannel* self, QObject* object);
extern __declspec(dllexport) bool QWebChannel_BlockUpdates(const QWebChannel* self);
extern __declspec(dllexport) void QWebChannel_SetBlockUpdates(QWebChannel* self, bool block);
extern __declspec(dllexport) int QWebChannel_PropertyUpdateInterval(const QWebChannel* self);
extern __declspec(dllexport) void QWebChannel_SetPropertyUpdateInterval(QWebChannel* self, int ms);
extern __declspec(dllexport) void QWebChannel_BlockUpdatesChanged(QWebChannel* self, bool block);
void QWebChannel_connect_BlockUpdatesChanged(QWebChannel* self, intptr_t slot);
extern __declspec(dllexport) void QWebChannel_ConnectTo(QWebChannel* self, QWebChannelAbstractTransport* transport);
extern __declspec(dllexport) void QWebChannel_DisconnectFrom(QWebChannel* self, QWebChannelAbstractTransport* transport);
extern __declspec(dllexport) struct miqt_string QWebChannel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QWebChannel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QWebChannel_override_virtual_Event(void* self, intptr_t slot);
bool QWebChannel_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QWebChannel_override_virtual_EventFilter(void* self, intptr_t slot);
bool QWebChannel_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QWebChannel_override_virtual_TimerEvent(void* self, intptr_t slot);
void QWebChannel_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QWebChannel_override_virtual_ChildEvent(void* self, intptr_t slot);
void QWebChannel_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QWebChannel_override_virtual_CustomEvent(void* self, intptr_t slot);
void QWebChannel_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QWebChannel_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QWebChannel_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QWebChannel_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QWebChannel_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QWebChannel_Delete(QWebChannel* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
