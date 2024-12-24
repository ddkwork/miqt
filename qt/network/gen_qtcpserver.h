#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QTCPSERVER_H
#define MIQT_QT_NETWORK_GEN_QTCPSERVER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QHostAddress QHostAddress;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkProxy QNetworkProxy;
typedef struct QObject QObject;
typedef struct QTcpServer QTcpServer;
typedef struct QTcpSocket QTcpSocket;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTcpServer* QTcpServer_new();
extern __declspec(dllexport) QTcpServer* QTcpServer_new2(QObject* parent);
extern __declspec(dllexport) void QTcpServer_virtbase(QTcpServer* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QTcpServer_MetaObject(const QTcpServer* self);
extern __declspec(dllexport) void* QTcpServer_Metacast(QTcpServer* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTcpServer_Tr(const char* s);
extern __declspec(dllexport) bool QTcpServer_Listen(QTcpServer* self);
extern __declspec(dllexport) void QTcpServer_Close(QTcpServer* self);
extern __declspec(dllexport) bool QTcpServer_IsListening(const QTcpServer* self);
extern __declspec(dllexport) void QTcpServer_SetMaxPendingConnections(QTcpServer* self, int numConnections);
extern __declspec(dllexport) int QTcpServer_MaxPendingConnections(const QTcpServer* self);
extern __declspec(dllexport) void QTcpServer_SetListenBacklogSize(QTcpServer* self, int size);
extern __declspec(dllexport) int QTcpServer_ListenBacklogSize(const QTcpServer* self);
extern __declspec(dllexport) uint16_t QTcpServer_ServerPort(const QTcpServer* self);
extern __declspec(dllexport) QHostAddress* QTcpServer_ServerAddress(const QTcpServer* self);
extern __declspec(dllexport) intptr_t QTcpServer_SocketDescriptor(const QTcpServer* self);
extern __declspec(dllexport) bool QTcpServer_SetSocketDescriptor(QTcpServer* self, intptr_t socketDescriptor);
extern __declspec(dllexport) bool QTcpServer_WaitForNewConnection(QTcpServer* self);
extern __declspec(dllexport) bool QTcpServer_HasPendingConnections(const QTcpServer* self);
extern __declspec(dllexport) QTcpSocket* QTcpServer_NextPendingConnection(QTcpServer* self);
extern __declspec(dllexport) int QTcpServer_ServerError(const QTcpServer* self);
extern __declspec(dllexport) struct miqt_string QTcpServer_ErrorString(const QTcpServer* self);
extern __declspec(dllexport) void QTcpServer_PauseAccepting(QTcpServer* self);
extern __declspec(dllexport) void QTcpServer_ResumeAccepting(QTcpServer* self);
extern __declspec(dllexport) void QTcpServer_SetProxy(QTcpServer* self, QNetworkProxy* networkProxy);
extern __declspec(dllexport) QNetworkProxy* QTcpServer_Proxy(const QTcpServer* self);
extern __declspec(dllexport) void QTcpServer_IncomingConnection(QTcpServer* self, intptr_t handle);
extern __declspec(dllexport) void QTcpServer_NewConnection(QTcpServer* self);
void QTcpServer_connect_NewConnection(QTcpServer* self, intptr_t slot);
extern __declspec(dllexport) void QTcpServer_AcceptError(QTcpServer* self, int socketError);
void QTcpServer_connect_AcceptError(QTcpServer* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QTcpServer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTcpServer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) bool QTcpServer_Listen1(QTcpServer* self, QHostAddress* address);
extern __declspec(dllexport) bool QTcpServer_Listen2(QTcpServer* self, QHostAddress* address, uint16_t port);
extern __declspec(dllexport) bool QTcpServer_WaitForNewConnection1(QTcpServer* self, int msec);
extern __declspec(dllexport) bool QTcpServer_WaitForNewConnection2(QTcpServer* self, int msec, bool* timedOut);
extern __declspec(dllexport) void QTcpServer_override_virtual_HasPendingConnections(void* self, intptr_t slot);
bool QTcpServer_virtualbase_HasPendingConnections(const void* self);
extern __declspec(dllexport) void QTcpServer_override_virtual_NextPendingConnection(void* self, intptr_t slot);
QTcpSocket* QTcpServer_virtualbase_NextPendingConnection(void* self);
extern __declspec(dllexport) void QTcpServer_override_virtual_IncomingConnection(void* self, intptr_t slot);
void QTcpServer_virtualbase_IncomingConnection(void* self, intptr_t handle);
extern __declspec(dllexport) void QTcpServer_override_virtual_Event(void* self, intptr_t slot);
bool QTcpServer_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QTcpServer_override_virtual_EventFilter(void* self, intptr_t slot);
bool QTcpServer_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QTcpServer_override_virtual_TimerEvent(void* self, intptr_t slot);
void QTcpServer_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QTcpServer_override_virtual_ChildEvent(void* self, intptr_t slot);
void QTcpServer_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QTcpServer_override_virtual_CustomEvent(void* self, intptr_t slot);
void QTcpServer_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QTcpServer_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QTcpServer_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QTcpServer_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QTcpServer_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QTcpServer_Delete(QTcpServer* self, bool isSubclass);

} 
