#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QLOCALSERVER_H
#define MIQT_QT_NETWORK_GEN_QLOCALSERVER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLocalServer QLocalServer;
typedef struct QLocalSocket QLocalSocket;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QLocalServer* QLocalServer_new();
extern __declspec(dllexport) 
QLocalServer* QLocalServer_new2(QObject* parent);
extern __declspec(dllexport) 
void QLocalServer_virtbase(QLocalServer* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QLocalServer_MetaObject(const QLocalServer* self);
extern __declspec(dllexport) 
void* QLocalServer_Metacast(QLocalServer* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QLocalServer_Tr(const char* s);
extern __declspec(dllexport) 
void QLocalServer_NewConnection(QLocalServer* self);
void QLocalServer_connect_NewConnection(QLocalServer* self, intptr_t slot);
extern __declspec(dllexport) 
void QLocalServer_Close(QLocalServer* self);
extern __declspec(dllexport) 
struct miqt_string QLocalServer_ErrorString(const QLocalServer* self);
extern __declspec(dllexport) 
bool QLocalServer_HasPendingConnections(const QLocalServer* self);
extern __declspec(dllexport) 
bool QLocalServer_IsListening(const QLocalServer* self);
extern __declspec(dllexport) 
bool QLocalServer_Listen(QLocalServer* self, struct miqt_string name);
extern __declspec(dllexport) 
bool QLocalServer_ListenWithSocketDescriptor(QLocalServer* self, intptr_t socketDescriptor);
extern __declspec(dllexport) 
int QLocalServer_MaxPendingConnections(const QLocalServer* self);
extern __declspec(dllexport) 
QLocalSocket* QLocalServer_NextPendingConnection(QLocalServer* self);
extern __declspec(dllexport) 
struct miqt_string QLocalServer_ServerName(const QLocalServer* self);
extern __declspec(dllexport) 
struct miqt_string QLocalServer_FullServerName(const QLocalServer* self);
extern __declspec(dllexport) 
bool QLocalServer_RemoveServer(struct miqt_string name);
extern __declspec(dllexport) 
int QLocalServer_ServerError(const QLocalServer* self);
extern __declspec(dllexport) 
void QLocalServer_SetMaxPendingConnections(QLocalServer* self, int numConnections);
extern __declspec(dllexport) 
bool QLocalServer_WaitForNewConnection(QLocalServer* self);
extern __declspec(dllexport) 
void QLocalServer_SetListenBacklogSize(QLocalServer* self, int size);
extern __declspec(dllexport) 
int QLocalServer_ListenBacklogSize(const QLocalServer* self);
extern __declspec(dllexport) 
void QLocalServer_SetSocketOptions(QLocalServer* self, SocketOptions options);
extern __declspec(dllexport) 
SocketOptions QLocalServer_SocketOptions(const QLocalServer* self);
extern __declspec(dllexport) 
intptr_t QLocalServer_SocketDescriptor(const QLocalServer* self);
extern __declspec(dllexport) 
void QLocalServer_IncomingConnection(QLocalServer* self, uintptr_t socketDescriptor);
extern __declspec(dllexport) 
struct miqt_string QLocalServer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QLocalServer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
bool QLocalServer_WaitForNewConnection1(QLocalServer* self, int msec);
extern __declspec(dllexport) 
bool QLocalServer_WaitForNewConnection2(QLocalServer* self, int msec, bool* timedOut);
extern __declspec(dllexport) 
void QLocalServer_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QLocalServer_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QLocalServer_override_virtual_Metacast(void* self, intptr_t slot);
void* QLocalServer_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QLocalServer_Delete(QLocalServer* self, bool isSubclass);

}
