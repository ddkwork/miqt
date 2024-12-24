#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QUDPSOCKET_H
#define MIQT_QT_NETWORK_GEN_QUDPSOCKET_H

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
typedef struct QNetworkDatagram QNetworkDatagram;
typedef struct QNetworkInterface QNetworkInterface;
typedef struct QObject QObject;
typedef struct QUdpSocket QUdpSocket;
typedef struct QVariant QVariant;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QUdpSocket* QUdpSocket_new();
extern __declspec(dllexport) QUdpSocket* QUdpSocket_new2(QObject* parent);
extern __declspec(dllexport) void QUdpSocket_virtbase(QUdpSocket* src, QAbstractSocket** outptr_QAbstractSocket);
extern __declspec(dllexport) QMetaObject* QUdpSocket_MetaObject(const QUdpSocket* self);
extern __declspec(dllexport) void* QUdpSocket_Metacast(QUdpSocket* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QUdpSocket_Tr(const char* s);
extern __declspec(dllexport) bool QUdpSocket_Bind(QUdpSocket* self, int addr);
extern __declspec(dllexport) bool QUdpSocket_JoinMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress);
extern __declspec(dllexport) bool QUdpSocket_JoinMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface);
extern __declspec(dllexport) bool QUdpSocket_LeaveMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress);
extern __declspec(dllexport) bool QUdpSocket_LeaveMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface);
extern __declspec(dllexport) QNetworkInterface* QUdpSocket_MulticastInterface(const QUdpSocket* self);
extern __declspec(dllexport) void QUdpSocket_SetMulticastInterface(QUdpSocket* self, QNetworkInterface* iface);
extern __declspec(dllexport) bool QUdpSocket_HasPendingDatagrams(const QUdpSocket* self);
extern __declspec(dllexport) long long QUdpSocket_PendingDatagramSize(const QUdpSocket* self);
extern __declspec(dllexport) QNetworkDatagram* QUdpSocket_ReceiveDatagram(QUdpSocket* self);
extern __declspec(dllexport) long long QUdpSocket_ReadDatagram(QUdpSocket* self, char* data, long long maxlen);
extern __declspec(dllexport) long long QUdpSocket_WriteDatagram(QUdpSocket* self, QNetworkDatagram* datagram);
extern __declspec(dllexport) long long QUdpSocket_WriteDatagram2(QUdpSocket* self, const char* data, long long lenVal, QHostAddress* host, uint16_t port);
extern __declspec(dllexport) long long QUdpSocket_WriteDatagram3(QUdpSocket* self, struct miqt_string datagram, QHostAddress* host, uint16_t port);
extern __declspec(dllexport) struct miqt_string QUdpSocket_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QUdpSocket_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) bool QUdpSocket_Bind2(QUdpSocket* self, int addr, uint16_t port);
extern __declspec(dllexport) bool QUdpSocket_Bind3(QUdpSocket* self, int addr, uint16_t port, BindMode mode);
extern __declspec(dllexport) QNetworkDatagram* QUdpSocket_ReceiveDatagram1(QUdpSocket* self, long long maxSize);
extern __declspec(dllexport) long long QUdpSocket_ReadDatagram3(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host);
extern __declspec(dllexport) long long QUdpSocket_ReadDatagram4(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host, uint16_t* port);
extern __declspec(dllexport) void QUdpSocket_override_virtual_Resume(void* self, intptr_t slot);
void QUdpSocket_virtualbase_Resume(void* self);
extern __declspec(dllexport) void QUdpSocket_override_virtual_Bind(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_Bind(void* self, QHostAddress* address, uint16_t port, BindMode mode);
extern __declspec(dllexport) void QUdpSocket_override_virtual_ConnectToHost(void* self, intptr_t slot);
void QUdpSocket_virtualbase_ConnectToHost(void* self, struct miqt_string hostName, uint16_t port, OpenMode mode, NetworkLayerProtocol protocol);
extern __declspec(dllexport) void QUdpSocket_override_virtual_DisconnectFromHost(void* self, intptr_t slot);
void QUdpSocket_virtualbase_DisconnectFromHost(void* self);
extern __declspec(dllexport) void QUdpSocket_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_BytesAvailable(const void* self);
extern __declspec(dllexport) void QUdpSocket_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_BytesToWrite(const void* self);
extern __declspec(dllexport) void QUdpSocket_override_virtual_SetReadBufferSize(void* self, intptr_t slot);
void QUdpSocket_virtualbase_SetReadBufferSize(void* self, long long size);
extern __declspec(dllexport) void QUdpSocket_override_virtual_SocketDescriptor(void* self, intptr_t slot);
intptr_t QUdpSocket_virtualbase_SocketDescriptor(const void* self);
extern __declspec(dllexport) void QUdpSocket_override_virtual_SetSocketDescriptor(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_SetSocketDescriptor(void* self, intptr_t socketDescriptor, SocketState state, OpenMode openMode);
extern __declspec(dllexport) void QUdpSocket_override_virtual_SetSocketOption(void* self, intptr_t slot);
void QUdpSocket_virtualbase_SetSocketOption(void* self, int option, QVariant* value);
extern __declspec(dllexport) void QUdpSocket_override_virtual_SocketOption(void* self, intptr_t slot);
QVariant* QUdpSocket_virtualbase_SocketOption(void* self, int option);
extern __declspec(dllexport) void QUdpSocket_override_virtual_Close(void* self, intptr_t slot);
void QUdpSocket_virtualbase_Close(void* self);
extern __declspec(dllexport) void QUdpSocket_override_virtual_IsSequential(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QUdpSocket_override_virtual_WaitForConnected(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForConnected(void* self, int msecs);
extern __declspec(dllexport) void QUdpSocket_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForReadyRead(void* self, int msecs);
extern __declspec(dllexport) void QUdpSocket_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForBytesWritten(void* self, int msecs);
extern __declspec(dllexport) void QUdpSocket_override_virtual_WaitForDisconnected(void* self, intptr_t slot);
bool QUdpSocket_virtualbase_WaitForDisconnected(void* self, int msecs);
extern __declspec(dllexport) void QUdpSocket_override_virtual_ReadData(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_ReadData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QUdpSocket_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QUdpSocket_override_virtual_SkipData(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_SkipData(void* self, long long maxSize);
extern __declspec(dllexport) void QUdpSocket_override_virtual_WriteData(void* self, intptr_t slot);
long long QUdpSocket_virtualbase_WriteData(void* self, const char* data, long long lenVal);
extern __declspec(dllexport) void QUdpSocket_Delete(QUdpSocket* self, bool isSubclass);

} 
