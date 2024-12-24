#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QTCPSOCKET_H
#define MIQT_QT_NETWORK_GEN_QTCPSOCKET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSocket QAbstractSocket;
typedef struct QHostAddress QHostAddress;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTcpSocket QTcpSocket;
typedef struct QVariant QVariant;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTcpSocket* QTcpSocket_new();
extern __declspec(dllexport) QTcpSocket* QTcpSocket_new2(QObject* parent);
extern __declspec(dllexport) void QTcpSocket_virtbase(QTcpSocket* src, QAbstractSocket** outptr_QAbstractSocket);
extern __declspec(dllexport) QMetaObject* QTcpSocket_MetaObject(const QTcpSocket* self);
extern __declspec(dllexport) void* QTcpSocket_Metacast(QTcpSocket* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTcpSocket_Tr(const char* s);
extern __declspec(dllexport) bool QTcpSocket_Bind(QTcpSocket* self, int addr);
extern __declspec(dllexport) struct miqt_string QTcpSocket_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTcpSocket_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) bool QTcpSocket_Bind2(QTcpSocket* self, int addr, uint16_t port);
extern __declspec(dllexport) bool QTcpSocket_Bind3(QTcpSocket* self, int addr, uint16_t port, BindMode mode);
extern __declspec(dllexport) void QTcpSocket_override_virtual_Resume(void* self, intptr_t slot);
void QTcpSocket_virtualbase_Resume(void* self);
extern __declspec(dllexport) void QTcpSocket_override_virtual_Bind(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_Bind(void* self, QHostAddress* address, uint16_t port, BindMode mode);
extern __declspec(dllexport) void QTcpSocket_override_virtual_ConnectToHost(void* self, intptr_t slot);
void QTcpSocket_virtualbase_ConnectToHost(void* self, struct miqt_string hostName, uint16_t port, OpenMode mode, NetworkLayerProtocol protocol);
extern __declspec(dllexport) void QTcpSocket_override_virtual_DisconnectFromHost(void* self, intptr_t slot);
void QTcpSocket_virtualbase_DisconnectFromHost(void* self);
extern __declspec(dllexport) void QTcpSocket_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_BytesAvailable(const void* self);
extern __declspec(dllexport) void QTcpSocket_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_BytesToWrite(const void* self);
extern __declspec(dllexport) void QTcpSocket_override_virtual_SetReadBufferSize(void* self, intptr_t slot);
void QTcpSocket_virtualbase_SetReadBufferSize(void* self, long long size);
extern __declspec(dllexport) void QTcpSocket_override_virtual_SocketDescriptor(void* self, intptr_t slot);
intptr_t QTcpSocket_virtualbase_SocketDescriptor(const void* self);
extern __declspec(dllexport) void QTcpSocket_override_virtual_SetSocketDescriptor(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_SetSocketDescriptor(void* self, intptr_t socketDescriptor, SocketState state, OpenMode openMode);
extern __declspec(dllexport) void QTcpSocket_override_virtual_SetSocketOption(void* self, intptr_t slot);
void QTcpSocket_virtualbase_SetSocketOption(void* self, int option, QVariant* value);
extern __declspec(dllexport) void QTcpSocket_override_virtual_SocketOption(void* self, intptr_t slot);
QVariant* QTcpSocket_virtualbase_SocketOption(void* self, int option);
extern __declspec(dllexport) void QTcpSocket_override_virtual_Close(void* self, intptr_t slot);
void QTcpSocket_virtualbase_Close(void* self);
extern __declspec(dllexport) void QTcpSocket_override_virtual_IsSequential(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QTcpSocket_override_virtual_WaitForConnected(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForConnected(void* self, int msecs);
extern __declspec(dllexport) void QTcpSocket_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForReadyRead(void* self, int msecs);
extern __declspec(dllexport) void QTcpSocket_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForBytesWritten(void* self, int msecs);
extern __declspec(dllexport) void QTcpSocket_override_virtual_WaitForDisconnected(void* self, intptr_t slot);
bool QTcpSocket_virtualbase_WaitForDisconnected(void* self, int msecs);
extern __declspec(dllexport) void QTcpSocket_override_virtual_ReadData(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_ReadData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QTcpSocket_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QTcpSocket_override_virtual_SkipData(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_SkipData(void* self, long long maxSize);
extern __declspec(dllexport) void QTcpSocket_override_virtual_WriteData(void* self, intptr_t slot);
long long QTcpSocket_virtualbase_WriteData(void* self, const char* data, long long lenVal);
extern __declspec(dllexport) void QTcpSocket_Delete(QTcpSocket* self, bool isSubclass);

} 
