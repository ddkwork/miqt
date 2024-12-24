#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKINTERFACE_H
#define MIQT_QT_NETWORK_GEN_QNETWORKINTERFACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QHostAddress QHostAddress;
typedef struct QNetworkAddressEntry QNetworkAddressEntry;
typedef struct QNetworkInterface QNetworkInterface;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QNetworkAddressEntry* QNetworkAddressEntry_new();
extern __declspec(dllexport) 
QNetworkAddressEntry* QNetworkAddressEntry_new2(QNetworkAddressEntry* other);
extern __declspec(dllexport) 
void QNetworkAddressEntry_OperatorAssign(QNetworkAddressEntry* self, QNetworkAddressEntry* other);
extern __declspec(dllexport) 
void QNetworkAddressEntry_Swap(QNetworkAddressEntry* self, QNetworkAddressEntry* other);
extern __declspec(dllexport) 
bool QNetworkAddressEntry_OperatorEqual(const QNetworkAddressEntry* self, QNetworkAddressEntry* other);
extern __declspec(dllexport) 
bool QNetworkAddressEntry_OperatorNotEqual(const QNetworkAddressEntry* self, QNetworkAddressEntry* other);
extern __declspec(dllexport) 
DnsEligibilityStatus QNetworkAddressEntry_DnsEligibility(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
void QNetworkAddressEntry_SetDnsEligibility(QNetworkAddressEntry* self, DnsEligibilityStatus status);
extern __declspec(dllexport) 
QHostAddress* QNetworkAddressEntry_Ip(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
void QNetworkAddressEntry_SetIp(QNetworkAddressEntry* self, QHostAddress* newIp);
extern __declspec(dllexport) 
QHostAddress* QNetworkAddressEntry_Netmask(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
void QNetworkAddressEntry_SetNetmask(QNetworkAddressEntry* self, QHostAddress* newNetmask);
extern __declspec(dllexport) 
int QNetworkAddressEntry_PrefixLength(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
void QNetworkAddressEntry_SetPrefixLength(QNetworkAddressEntry* self, int length);
extern __declspec(dllexport) 
QHostAddress* QNetworkAddressEntry_Broadcast(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
void QNetworkAddressEntry_SetBroadcast(QNetworkAddressEntry* self, QHostAddress* newBroadcast);
extern __declspec(dllexport) 
bool QNetworkAddressEntry_IsLifetimeKnown(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
QDeadlineTimer* QNetworkAddressEntry_PreferredLifetime(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
QDeadlineTimer* QNetworkAddressEntry_ValidityLifetime(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
void QNetworkAddressEntry_SetAddressLifetime(QNetworkAddressEntry* self, QDeadlineTimer* preferred, QDeadlineTimer* validity);
extern __declspec(dllexport) 
void QNetworkAddressEntry_ClearAddressLifetime(QNetworkAddressEntry* self);
extern __declspec(dllexport) 
bool QNetworkAddressEntry_IsPermanent(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
bool QNetworkAddressEntry_IsTemporary(const QNetworkAddressEntry* self);
extern __declspec(dllexport) 
void QNetworkAddressEntry_Delete(QNetworkAddressEntry* self, bool isSubclass);

extern __declspec(dllexport) 
QNetworkInterface* QNetworkInterface_new();
extern __declspec(dllexport) 
QNetworkInterface* QNetworkInterface_new2(QNetworkInterface* other);
extern __declspec(dllexport) 
void QNetworkInterface_OperatorAssign(QNetworkInterface* self, QNetworkInterface* other);
extern __declspec(dllexport) 
void QNetworkInterface_Swap(QNetworkInterface* self, QNetworkInterface* other);
extern __declspec(dllexport) 
bool QNetworkInterface_IsValid(const QNetworkInterface* self);
extern __declspec(dllexport) 
int QNetworkInterface_Index(const QNetworkInterface* self);
extern __declspec(dllexport) 
int QNetworkInterface_MaximumTransmissionUnit(const QNetworkInterface* self);
extern __declspec(dllexport) 
struct miqt_string QNetworkInterface_Name(const QNetworkInterface* self);
extern __declspec(dllexport) 
struct miqt_string QNetworkInterface_HumanReadableName(const QNetworkInterface* self);
extern __declspec(dllexport) 
InterfaceFlags QNetworkInterface_Flags(const QNetworkInterface* self);
extern __declspec(dllexport) 
InterfaceType QNetworkInterface_Type(const QNetworkInterface* self);
extern __declspec(dllexport) 
struct miqt_string QNetworkInterface_HardwareAddress(const QNetworkInterface* self);
extern __declspec(dllexport) 
struct miqt_array /* of QNetworkAddressEntry* */  QNetworkInterface_AddressEntries(const QNetworkInterface* self);
extern __declspec(dllexport) 
int QNetworkInterface_InterfaceIndexFromName(struct miqt_string name);
extern __declspec(dllexport) 
QNetworkInterface* QNetworkInterface_InterfaceFromName(struct miqt_string name);
extern __declspec(dllexport) 
QNetworkInterface* QNetworkInterface_InterfaceFromIndex(int index);
extern __declspec(dllexport) 
struct miqt_string QNetworkInterface_InterfaceNameFromIndex(int index);
extern __declspec(dllexport) 
struct miqt_array /* of QNetworkInterface* */  QNetworkInterface_AllInterfaces();
extern __declspec(dllexport) 
struct miqt_array /* of QHostAddress* */  QNetworkInterface_AllAddresses();
extern __declspec(dllexport) 
void QNetworkInterface_Delete(QNetworkInterface* self, bool isSubclass);

}
