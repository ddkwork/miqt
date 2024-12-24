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

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QUdpSocket* QUdpSocket_new();
extern __declspec(dllexport) 
QUdpSocket* QUdpSocket_new2(QObject* parent);
extern __declspec(dllexport) 
void QUdpSocket_virtbase(QUdpSocket* src
, QAbstractSocket** outptr_QAbstractSocket
);
extern __declspec(dllexport) 
QMetaObject* QUdpSocket_MetaObject(const QUdpSocket* self);
extern __declspec(dllexport) 
void* QUdpSocket_Metacast(QUdpSocket* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QUdpSocket_Tr(const char* s);
extern __declspec(dllexport) 
bool QUdpSocket_Bind(QUdpSocket* self, int addr);
extern __declspec(dllexport) 
bool QUdpSocket_JoinMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress);
extern __declspec(dllexport) 
bool QUdpSocket_JoinMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface);
extern __declspec(dllexport) 
bool QUdpSocket_LeaveMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress);
extern __declspec(dllexport) 
bool QUdpSocket_LeaveMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface);
extern __declspec(dllexport) 
QNetworkInterface* QUdpSocket_MulticastInterface(const QUdpSocket* self);
extern __declspec(dllexport) 
void QUdpSocket_SetMulticastInterface(QUdpSocket* self, QNetworkInterface* iface);
extern __declspec(dllexport) 
bool QUdpSocket_HasPendingDatagrams(const QUdpSocket* self);
extern __declspec(dllexport) 
long long QUdpSocket_PendingDatagramSize(const QUdpSocket* self);
extern __declspec(dllexport) 
QNetworkDatagram* QUdpSocket_ReceiveDatagram(QUdpSocket* self);
extern __declspec(dllexport) 
long long QUdpSocket_ReadDatagram(QUdpSocket* self, char* data, long long maxlen);
extern __declspec(dllexport) 
long long QUdpSocket_WriteDatagram(QUdpSocket* self, QNetworkDatagram* datagram);
extern __declspec(dllexport) 
long long QUdpSocket_WriteDatagram2(QUdpSocket* self, const char* data, long long lenVal, QHostAddress* host, uint16_t port);
extern __declspec(dllexport) 
long long QUdpSocket_WriteDatagram3(QUdpSocket* self, struct miqt_string datagram, QHostAddress* host, uint16_t port);
extern __declspec(dllexport) 
struct miqt_string QUdpSocket_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QUdpSocket_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
bool QUdpSocket_Bind2(QUdpSocket* self, int addr, uint16_t port);
extern __declspec(dllexport) 
bool QUdpSocket_Bind3(QUdpSocket* self, int addr, uint16_t port, BindMode mode);
extern __declspec(dllexport) 
QNetworkDatagram* QUdpSocket_ReceiveDatagram1(QUdpSocket* self, long long maxSize);
extern __declspec(dllexport) 
long long QUdpSocket_ReadDatagram3(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host);
extern __declspec(dllexport) 
long long QUdpSocket_ReadDatagram4(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host, uint16_t* port);
extern __declspec(dllexport) 
void QUdpSocket_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QUdpSocket_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QUdpSocket_override_virtual_Metacast(void* self, intptr_t slot);
void* QUdpSocket_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QUdpSocket_Delete(QUdpSocket* self, bool isSubclass);

}
