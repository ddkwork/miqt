#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QABSTRACTSOCKET_H
#define MIQT_QT_NETWORK_GEN_QABSTRACTSOCKET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSocket QAbstractSocket;
typedef struct QAuthenticator QAuthenticator;
typedef struct QHostAddress QHostAddress;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkProxy QNetworkProxy;
typedef struct QObject QObject;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractSocket* QAbstractSocket_new(SocketType socketType, QObject* parent);
extern __declspec(dllexport) 
void QAbstractSocket_virtbase(QAbstractSocket* src
, QIODevice** outptr_QIODevice
);
extern __declspec(dllexport) 
QMetaObject* QAbstractSocket_MetaObject(const QAbstractSocket* self);
extern __declspec(dllexport) 
void* QAbstractSocket_Metacast(QAbstractSocket* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractSocket_Tr(const char* s);
extern __declspec(dllexport) 
void QAbstractSocket_Resume(QAbstractSocket* self);
extern __declspec(dllexport) 
PauseModes QAbstractSocket_PauseMode(const QAbstractSocket* self);
extern __declspec(dllexport) 
void QAbstractSocket_SetPauseMode(QAbstractSocket* self, PauseModes pauseMode);
extern __declspec(dllexport) 
bool QAbstractSocket_Bind(QAbstractSocket* self, QHostAddress* address, uint16_t port, BindMode mode);
extern __declspec(dllexport) 
bool QAbstractSocket_Bind2(QAbstractSocket* self);
extern __declspec(dllexport) 
void QAbstractSocket_ConnectToHost(QAbstractSocket* self, struct miqt_string hostName, uint16_t port, OpenMode mode, NetworkLayerProtocol protocol);
extern __declspec(dllexport) 
void QAbstractSocket_ConnectToHost2(QAbstractSocket* self, QHostAddress* address, uint16_t port);
extern __declspec(dllexport) 
void QAbstractSocket_DisconnectFromHost(QAbstractSocket* self);
extern __declspec(dllexport) 
bool QAbstractSocket_IsValid(const QAbstractSocket* self);
extern __declspec(dllexport) 
long long QAbstractSocket_BytesAvailable(const QAbstractSocket* self);
extern __declspec(dllexport) 
long long QAbstractSocket_BytesToWrite(const QAbstractSocket* self);
extern __declspec(dllexport) 
uint16_t QAbstractSocket_LocalPort(const QAbstractSocket* self);
extern __declspec(dllexport) 
QHostAddress* QAbstractSocket_LocalAddress(const QAbstractSocket* self);
extern __declspec(dllexport) 
uint16_t QAbstractSocket_PeerPort(const QAbstractSocket* self);
extern __declspec(dllexport) 
QHostAddress* QAbstractSocket_PeerAddress(const QAbstractSocket* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractSocket_PeerName(const QAbstractSocket* self);
extern __declspec(dllexport) 
long long QAbstractSocket_ReadBufferSize(const QAbstractSocket* self);
extern __declspec(dllexport) 
void QAbstractSocket_SetReadBufferSize(QAbstractSocket* self, long long size);
extern __declspec(dllexport) 
void QAbstractSocket_Abort(QAbstractSocket* self);
extern __declspec(dllexport) 
intptr_t QAbstractSocket_SocketDescriptor(const QAbstractSocket* self);
extern __declspec(dllexport) 
bool QAbstractSocket_SetSocketDescriptor(QAbstractSocket* self, intptr_t socketDescriptor, SocketState state, OpenMode openMode);
extern __declspec(dllexport) 
void QAbstractSocket_SetSocketOption(QAbstractSocket* self, int option, QVariant* value);
extern __declspec(dllexport) 
QVariant* QAbstractSocket_SocketOption(QAbstractSocket* self, int option);
extern __declspec(dllexport) 
SocketType QAbstractSocket_SocketType(const QAbstractSocket* self);
extern __declspec(dllexport) 
SocketState QAbstractSocket_State(const QAbstractSocket* self);
extern __declspec(dllexport) 
SocketError QAbstractSocket_Error(const QAbstractSocket* self);
extern __declspec(dllexport) 
void QAbstractSocket_Close(QAbstractSocket* self);
extern __declspec(dllexport) 
bool QAbstractSocket_IsSequential(const QAbstractSocket* self);
extern __declspec(dllexport) 
bool QAbstractSocket_Flush(QAbstractSocket* self);
extern __declspec(dllexport) 
bool QAbstractSocket_WaitForConnected(QAbstractSocket* self, int msecs);
extern __declspec(dllexport) 
bool QAbstractSocket_WaitForReadyRead(QAbstractSocket* self, int msecs);
extern __declspec(dllexport) 
bool QAbstractSocket_WaitForBytesWritten(QAbstractSocket* self, int msecs);
extern __declspec(dllexport) 
bool QAbstractSocket_WaitForDisconnected(QAbstractSocket* self, int msecs);
extern __declspec(dllexport) 
void QAbstractSocket_SetProxy(QAbstractSocket* self, QNetworkProxy* networkProxy);
extern __declspec(dllexport) 
QNetworkProxy* QAbstractSocket_Proxy(const QAbstractSocket* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractSocket_ProtocolTag(const QAbstractSocket* self);
extern __declspec(dllexport) 
void QAbstractSocket_SetProtocolTag(QAbstractSocket* self, struct miqt_string tag);
extern __declspec(dllexport) 
void QAbstractSocket_HostFound(QAbstractSocket* self);
void QAbstractSocket_connect_HostFound(QAbstractSocket* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSocket_Connected(QAbstractSocket* self);
void QAbstractSocket_connect_Connected(QAbstractSocket* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSocket_Disconnected(QAbstractSocket* self);
void QAbstractSocket_connect_Disconnected(QAbstractSocket* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSocket_StateChanged(QAbstractSocket* self, int param1);
void QAbstractSocket_connect_StateChanged(QAbstractSocket* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSocket_ErrorOccurred(QAbstractSocket* self, int param1);
void QAbstractSocket_connect_ErrorOccurred(QAbstractSocket* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractSocket_ProxyAuthenticationRequired(QAbstractSocket* self, QNetworkProxy* proxy, QAuthenticator* authenticator);
void QAbstractSocket_connect_ProxyAuthenticationRequired(QAbstractSocket* self, intptr_t slot);
extern __declspec(dllexport) 
long long QAbstractSocket_ReadData(QAbstractSocket* self, char* data, long long maxlen);
extern __declspec(dllexport) 
long long QAbstractSocket_ReadLineData(QAbstractSocket* self, char* data, long long maxlen);
extern __declspec(dllexport) 
long long QAbstractSocket_SkipData(QAbstractSocket* self, long long maxSize);
extern __declspec(dllexport) 
long long QAbstractSocket_WriteData(QAbstractSocket* self, const char* data, long long lenVal);
extern __declspec(dllexport) 
struct miqt_string QAbstractSocket_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractSocket_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
bool QAbstractSocket_Bind1(QAbstractSocket* self, uint16_t port);
extern __declspec(dllexport) 
bool QAbstractSocket_Bind22(QAbstractSocket* self, uint16_t port, BindMode mode);
extern __declspec(dllexport) 
void QAbstractSocket_ConnectToHost3(QAbstractSocket* self, QHostAddress* address, uint16_t port, OpenMode mode);
extern __declspec(dllexport) 
void QAbstractSocket_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractSocket_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractSocket_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractSocket_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractSocket_Delete(QAbstractSocket* self, bool isSubclass);

}
