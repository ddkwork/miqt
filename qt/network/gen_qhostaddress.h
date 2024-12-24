#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QHOSTADDRESS_H
#define MIQT_QT_NETWORK_GEN_QHOSTADDRESS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHostAddress QHostAddress;
typedef struct QIPv6Address QIPv6Address;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QIPv6Address* QIPv6Address_new();
extern __declspec(dllexport) 
QIPv6Address* QIPv6Address_new2(QIPv6Address* param1);
extern __declspec(dllexport) 
unsigned char QIPv6Address_OperatorSubscript(const QIPv6Address* self, int index);
extern __declspec(dllexport) 
void QIPv6Address_Delete(QIPv6Address* self, bool isSubclass);

extern __declspec(dllexport) 
QHostAddress* QHostAddress_new();
extern __declspec(dllexport) 
QHostAddress* QHostAddress_new2(unsigned int ip4Addr);
extern __declspec(dllexport) 
QHostAddress* QHostAddress_new3(const unsigned char* ip6Addr);
extern __declspec(dllexport) 
QHostAddress* QHostAddress_new4(QIPv6Address* ip6Addr);
extern __declspec(dllexport) 
QHostAddress* QHostAddress_new5(struct miqt_string address);
extern __declspec(dllexport) 
QHostAddress* QHostAddress_new6(QHostAddress* copyVal);
extern __declspec(dllexport) 
QHostAddress* QHostAddress_new7(SpecialAddress address);
extern __declspec(dllexport) 
void QHostAddress_OperatorAssign(QHostAddress* self, QHostAddress* other);
extern __declspec(dllexport) 
void QHostAddress_OperatorAssignWithAddress(QHostAddress* self, SpecialAddress address);
extern __declspec(dllexport) 
void QHostAddress_Swap(QHostAddress* self, QHostAddress* other);
extern __declspec(dllexport) 
void QHostAddress_SetAddress(QHostAddress* self, unsigned int ip4Addr);
extern __declspec(dllexport) 
void QHostAddress_SetAddressWithIp6Addr(QHostAddress* self, const unsigned char* ip6Addr);
extern __declspec(dllexport) 
void QHostAddress_SetAddress2(QHostAddress* self, QIPv6Address* ip6Addr);
extern __declspec(dllexport) 
bool QHostAddress_SetAddress3(QHostAddress* self, struct miqt_string address);
extern __declspec(dllexport) 
void QHostAddress_SetAddress4(QHostAddress* self, SpecialAddress address);
extern __declspec(dllexport) 
NetworkLayerProtocol QHostAddress_Protocol(const QHostAddress* self);
extern __declspec(dllexport) 
unsigned int QHostAddress_ToIPv4Address(const QHostAddress* self);
extern __declspec(dllexport) 
QIPv6Address* QHostAddress_ToIPv6Address(const QHostAddress* self);
extern __declspec(dllexport) 
struct miqt_string QHostAddress_ToString(const QHostAddress* self);
extern __declspec(dllexport) 
struct miqt_string QHostAddress_ScopeId(const QHostAddress* self);
extern __declspec(dllexport) 
void QHostAddress_SetScopeId(QHostAddress* self, struct miqt_string id);
extern __declspec(dllexport) 
bool QHostAddress_IsEqual(const QHostAddress* self, QHostAddress* address);
extern __declspec(dllexport) 
bool QHostAddress_OperatorEqual(const QHostAddress* self, QHostAddress* address);
extern __declspec(dllexport) 
bool QHostAddress_OperatorEqualWithAddress(const QHostAddress* self, SpecialAddress address);
extern __declspec(dllexport) 
bool QHostAddress_OperatorNotEqual(const QHostAddress* self, QHostAddress* address);
extern __declspec(dllexport) 
bool QHostAddress_OperatorNotEqualWithAddress(const QHostAddress* self, SpecialAddress address);
extern __declspec(dllexport) 
bool QHostAddress_IsNull(const QHostAddress* self);
extern __declspec(dllexport) 
void QHostAddress_Clear(QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsInSubnet(const QHostAddress* self, QHostAddress* subnet, int netmask);
extern __declspec(dllexport) 
bool QHostAddress_IsLoopback(const QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsGlobal(const QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsLinkLocal(const QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsSiteLocal(const QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsUniqueLocalUnicast(const QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsMulticast(const QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsBroadcast(const QHostAddress* self);
extern __declspec(dllexport) 
bool QHostAddress_IsPrivateUse(const QHostAddress* self);
extern __declspec(dllexport) 
unsigned int QHostAddress_ToIPv4Address1(const QHostAddress* self, bool* ok);
extern __declspec(dllexport) 
bool QHostAddress_IsEqual2(const QHostAddress* self, QHostAddress* address, ConversionMode mode);
extern __declspec(dllexport) 
void QHostAddress_Delete(QHostAddress* self, bool isSubclass);

}
