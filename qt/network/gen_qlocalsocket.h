#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QLOCALSOCKET_H
#define MIQT_QT_NETWORK_GEN_QLOCALSOCKET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QLocalSocket QLocalSocket;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLocalSocket* QLocalSocket_new();
extern __declspec(dllexport) QLocalSocket* QLocalSocket_new2(QObject* parent);
extern __declspec(dllexport) void QLocalSocket_virtbase(QLocalSocket* src, QIODevice** outptr_QIODevice);
extern __declspec(dllexport) QMetaObject* QLocalSocket_MetaObject(const QLocalSocket* self);
extern __declspec(dllexport) void* QLocalSocket_Metacast(QLocalSocket* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QLocalSocket_Tr(const char* s);
extern __declspec(dllexport) void QLocalSocket_ConnectToServer(QLocalSocket* self);
extern __declspec(dllexport) void QLocalSocket_ConnectToServerWithName(QLocalSocket* self, struct miqt_string name);
extern __declspec(dllexport) void QLocalSocket_DisconnectFromServer(QLocalSocket* self);
extern __declspec(dllexport) void QLocalSocket_SetServerName(QLocalSocket* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QLocalSocket_ServerName(const QLocalSocket* self);
extern __declspec(dllexport) struct miqt_string QLocalSocket_FullServerName(const QLocalSocket* self);
extern __declspec(dllexport) void QLocalSocket_Abort(QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_IsSequential(const QLocalSocket* self);
extern __declspec(dllexport) long long QLocalSocket_BytesAvailable(const QLocalSocket* self);
extern __declspec(dllexport) long long QLocalSocket_BytesToWrite(const QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_CanReadLine(const QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_Open(QLocalSocket* self, OpenMode openMode);
extern __declspec(dllexport) void QLocalSocket_Close(QLocalSocket* self);
extern __declspec(dllexport) LocalSocketError QLocalSocket_Error(const QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_Flush(QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_IsValid(const QLocalSocket* self);
extern __declspec(dllexport) long long QLocalSocket_ReadBufferSize(const QLocalSocket* self);
extern __declspec(dllexport) void QLocalSocket_SetReadBufferSize(QLocalSocket* self, long long size);
extern __declspec(dllexport) bool QLocalSocket_SetSocketDescriptor(QLocalSocket* self, intptr_t socketDescriptor);
extern __declspec(dllexport) intptr_t QLocalSocket_SocketDescriptor(const QLocalSocket* self);
extern __declspec(dllexport) void QLocalSocket_SetSocketOptions(QLocalSocket* self, SocketOptions option);
extern __declspec(dllexport) SocketOptions QLocalSocket_SocketOptions(const QLocalSocket* self);
extern __declspec(dllexport) LocalSocketState QLocalSocket_State(const QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_WaitForBytesWritten(QLocalSocket* self, int msecs);
extern __declspec(dllexport) bool QLocalSocket_WaitForConnected(QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_WaitForDisconnected(QLocalSocket* self);
extern __declspec(dllexport) bool QLocalSocket_WaitForReadyRead(QLocalSocket* self, int msecs);
extern __declspec(dllexport) void QLocalSocket_Connected(QLocalSocket* self);
void QLocalSocket_connect_Connected(QLocalSocket* self, intptr_t slot);
extern __declspec(dllexport) void QLocalSocket_Disconnected(QLocalSocket* self);
void QLocalSocket_connect_Disconnected(QLocalSocket* self, intptr_t slot);
extern __declspec(dllexport) void QLocalSocket_ErrorOccurred(QLocalSocket* self, int socketError);
void QLocalSocket_connect_ErrorOccurred(QLocalSocket* self, intptr_t slot);
extern __declspec(dllexport) void QLocalSocket_StateChanged(QLocalSocket* self, int socketState);
void QLocalSocket_connect_StateChanged(QLocalSocket* self, intptr_t slot);
extern __declspec(dllexport) long long QLocalSocket_ReadData(QLocalSocket* self, char* param1, long long param2);
extern __declspec(dllexport) long long QLocalSocket_ReadLineData(QLocalSocket* self, char* data, long long maxSize);
extern __declspec(dllexport) long long QLocalSocket_SkipData(QLocalSocket* self, long long maxSize);
extern __declspec(dllexport) long long QLocalSocket_WriteData(QLocalSocket* self, const char* param1, long long param2);
extern __declspec(dllexport) struct miqt_string QLocalSocket_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QLocalSocket_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QLocalSocket_ConnectToServer1(QLocalSocket* self, OpenMode openMode);
extern __declspec(dllexport) void QLocalSocket_ConnectToServer2(QLocalSocket* self, struct miqt_string name, OpenMode openMode);
extern __declspec(dllexport) bool QLocalSocket_SetSocketDescriptor2(QLocalSocket* self, intptr_t socketDescriptor, LocalSocketState socketState);
extern __declspec(dllexport) bool QLocalSocket_SetSocketDescriptor3(QLocalSocket* self, intptr_t socketDescriptor, LocalSocketState socketState, OpenMode openMode);
extern __declspec(dllexport) bool QLocalSocket_WaitForConnected1(QLocalSocket* self, int msecs);
extern __declspec(dllexport) bool QLocalSocket_WaitForDisconnected1(QLocalSocket* self, int msecs);
extern __declspec(dllexport) void QLocalSocket_override_virtual_IsSequential(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_BytesAvailable(const void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_BytesToWrite(const void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_CanReadLine(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_CanReadLine(const void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_Open(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_Open(void* self, OpenMode openMode);
extern __declspec(dllexport) void QLocalSocket_override_virtual_Close(void* self, intptr_t slot);
void QLocalSocket_virtualbase_Close(void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_WaitForBytesWritten(void* self, int msecs);
extern __declspec(dllexport) void QLocalSocket_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_WaitForReadyRead(void* self, int msecs);
extern __declspec(dllexport) void QLocalSocket_override_virtual_ReadData(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_ReadData(void* self, char* param1, long long param2);
extern __declspec(dllexport) void QLocalSocket_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_ReadLineData(void* self, char* data, long long maxSize);
extern __declspec(dllexport) void QLocalSocket_override_virtual_SkipData(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_SkipData(void* self, long long maxSize);
extern __declspec(dllexport) void QLocalSocket_override_virtual_WriteData(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_WriteData(void* self, const char* param1, long long param2);
extern __declspec(dllexport) void QLocalSocket_override_virtual_Pos(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_Pos(const void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_Size(void* self, intptr_t slot);
long long QLocalSocket_virtualbase_Size(const void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_Seek(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_Seek(void* self, long long pos);
extern __declspec(dllexport) void QLocalSocket_override_virtual_AtEnd(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_AtEnd(const void* self);
extern __declspec(dllexport) void QLocalSocket_override_virtual_Reset(void* self, intptr_t slot);
bool QLocalSocket_virtualbase_Reset(void* self);
extern __declspec(dllexport) void QLocalSocket_Delete(QLocalSocket* self, bool isSubclass);

} 
