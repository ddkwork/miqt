#pragma once
#ifndef MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNELABSTRACTTRANSPORT_H
#define MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNELABSTRACTTRANSPORT_H

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
class QJsonObject;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTimerEvent;
class QWebChannelAbstractTransport;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QJsonObject QJsonObject;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWebChannelAbstractTransport QWebChannelAbstractTransport;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QWebChannelAbstractTransport* QWebChannelAbstractTransport_new();
extern __declspec(dllexport) QWebChannelAbstractTransport* QWebChannelAbstractTransport_new2(QObject* parent);
extern __declspec(dllexport) void QWebChannelAbstractTransport_virtbase(QWebChannelAbstractTransport* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QWebChannelAbstractTransport_MetaObject(const QWebChannelAbstractTransport* self);
extern __declspec(dllexport) void* QWebChannelAbstractTransport_Metacast(QWebChannelAbstractTransport* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QWebChannelAbstractTransport_Tr(const char* s);
extern __declspec(dllexport) void QWebChannelAbstractTransport_SendMessage(QWebChannelAbstractTransport* self, QJsonObject* message);
extern __declspec(dllexport) void QWebChannelAbstractTransport_MessageReceived(QWebChannelAbstractTransport* self, QJsonObject* message, QWebChannelAbstractTransport* transport);
void QWebChannelAbstractTransport_connect_MessageReceived(QWebChannelAbstractTransport* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QWebChannelAbstractTransport_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QWebChannelAbstractTransport_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_SendMessage(void* self, intptr_t slot);
void QWebChannelAbstractTransport_virtualbase_SendMessage(void* self, QJsonObject* message);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_Event(void* self, intptr_t slot);
bool QWebChannelAbstractTransport_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_EventFilter(void* self, intptr_t slot);
bool QWebChannelAbstractTransport_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_TimerEvent(void* self, intptr_t slot);
void QWebChannelAbstractTransport_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_ChildEvent(void* self, intptr_t slot);
void QWebChannelAbstractTransport_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_CustomEvent(void* self, intptr_t slot);
void QWebChannelAbstractTransport_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QWebChannelAbstractTransport_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QWebChannelAbstractTransport_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QWebChannelAbstractTransport_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QWebChannelAbstractTransport_Delete(QWebChannelAbstractTransport* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
