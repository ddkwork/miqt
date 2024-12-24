#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKDATAGRAM_H
#define MIQT_QT_NETWORK_GEN_QNETWORKDATAGRAM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHostAddress QHostAddress;
typedef struct QNetworkDatagram QNetworkDatagram;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNetworkDatagram* QNetworkDatagram_new();
extern __declspec(dllexport) QNetworkDatagram* QNetworkDatagram_new2(struct miqt_string data);
extern __declspec(dllexport) QNetworkDatagram* QNetworkDatagram_new3(QNetworkDatagram* other);
extern __declspec(dllexport) QNetworkDatagram* QNetworkDatagram_new4(struct miqt_string data, QHostAddress* destinationAddress);
extern __declspec(dllexport) QNetworkDatagram* QNetworkDatagram_new5(struct miqt_string data, QHostAddress* destinationAddress, uint16_t port);
extern __declspec(dllexport) void QNetworkDatagram_OperatorAssign(QNetworkDatagram* self, QNetworkDatagram* other);
extern __declspec(dllexport) void QNetworkDatagram_Swap(QNetworkDatagram* self, QNetworkDatagram* other);
extern __declspec(dllexport) void QNetworkDatagram_Clear(QNetworkDatagram* self);
extern __declspec(dllexport) bool QNetworkDatagram_IsValid(const QNetworkDatagram* self);
extern __declspec(dllexport) bool QNetworkDatagram_IsNull(const QNetworkDatagram* self);
extern __declspec(dllexport) unsigned int QNetworkDatagram_InterfaceIndex(const QNetworkDatagram* self);
extern __declspec(dllexport) void QNetworkDatagram_SetInterfaceIndex(QNetworkDatagram* self, unsigned int index);
extern __declspec(dllexport) QHostAddress* QNetworkDatagram_SenderAddress(const QNetworkDatagram* self);
extern __declspec(dllexport) QHostAddress* QNetworkDatagram_DestinationAddress(const QNetworkDatagram* self);
extern __declspec(dllexport) int QNetworkDatagram_SenderPort(const QNetworkDatagram* self);
extern __declspec(dllexport) int QNetworkDatagram_DestinationPort(const QNetworkDatagram* self);
extern __declspec(dllexport) void QNetworkDatagram_SetSender(QNetworkDatagram* self, QHostAddress* address);
extern __declspec(dllexport) void QNetworkDatagram_SetDestination(QNetworkDatagram* self, QHostAddress* address, uint16_t port);
extern __declspec(dllexport) int QNetworkDatagram_HopLimit(const QNetworkDatagram* self);
extern __declspec(dllexport) void QNetworkDatagram_SetHopLimit(QNetworkDatagram* self, int count);
extern __declspec(dllexport) struct miqt_string QNetworkDatagram_Data(const QNetworkDatagram* self);
extern __declspec(dllexport) void QNetworkDatagram_SetData(QNetworkDatagram* self, struct miqt_string data);
extern __declspec(dllexport) QNetworkDatagram* QNetworkDatagram_MakeReply(const QNetworkDatagram* self, struct miqt_string payload);
extern __declspec(dllexport) void QNetworkDatagram_SetSender2(QNetworkDatagram* self, QHostAddress* address, uint16_t port);
extern __declspec(dllexport) void QNetworkDatagram_Delete(QNetworkDatagram* self, bool isSubclass);

} 
