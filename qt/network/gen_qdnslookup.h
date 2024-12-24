#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QDNSLOOKUP_H
#define MIQT_QT_NETWORK_GEN_QDNSLOOKUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QDnsDomainNameRecord QDnsDomainNameRecord;
typedef struct QDnsHostAddressRecord QDnsHostAddressRecord;
typedef struct QDnsLookup QDnsLookup;
typedef struct QDnsMailExchangeRecord QDnsMailExchangeRecord;
typedef struct QDnsServiceRecord QDnsServiceRecord;
typedef struct QDnsTextRecord QDnsTextRecord;
typedef struct QDnsTlsAssociationRecord QDnsTlsAssociationRecord;
typedef struct QEvent QEvent;
typedef struct QHostAddress QHostAddress;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSslConfiguration QSslConfiguration;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDnsDomainNameRecord* QDnsDomainNameRecord_new();
extern __declspec(dllexport) QDnsDomainNameRecord* QDnsDomainNameRecord_new2(QDnsDomainNameRecord* other);
extern __declspec(dllexport) void QDnsDomainNameRecord_OperatorAssign(QDnsDomainNameRecord* self, QDnsDomainNameRecord* other);
extern __declspec(dllexport) void QDnsDomainNameRecord_Swap(QDnsDomainNameRecord* self, QDnsDomainNameRecord* other);
extern __declspec(dllexport) struct miqt_string QDnsDomainNameRecord_Name(const QDnsDomainNameRecord* self);
extern __declspec(dllexport) unsigned int QDnsDomainNameRecord_TimeToLive(const QDnsDomainNameRecord* self);
extern __declspec(dllexport) struct miqt_string QDnsDomainNameRecord_Value(const QDnsDomainNameRecord* self);
extern __declspec(dllexport) void QDnsDomainNameRecord_Delete(QDnsDomainNameRecord* self, bool isSubclass);

extern __declspec(dllexport) QDnsHostAddressRecord* QDnsHostAddressRecord_new();
extern __declspec(dllexport) QDnsHostAddressRecord* QDnsHostAddressRecord_new2(QDnsHostAddressRecord* other);
extern __declspec(dllexport) void QDnsHostAddressRecord_OperatorAssign(QDnsHostAddressRecord* self, QDnsHostAddressRecord* other);
extern __declspec(dllexport) void QDnsHostAddressRecord_Swap(QDnsHostAddressRecord* self, QDnsHostAddressRecord* other);
extern __declspec(dllexport) struct miqt_string QDnsHostAddressRecord_Name(const QDnsHostAddressRecord* self);
extern __declspec(dllexport) unsigned int QDnsHostAddressRecord_TimeToLive(const QDnsHostAddressRecord* self);
extern __declspec(dllexport) QHostAddress* QDnsHostAddressRecord_Value(const QDnsHostAddressRecord* self);
extern __declspec(dllexport) void QDnsHostAddressRecord_Delete(QDnsHostAddressRecord* self, bool isSubclass);

extern __declspec(dllexport) QDnsMailExchangeRecord* QDnsMailExchangeRecord_new();
extern __declspec(dllexport) QDnsMailExchangeRecord* QDnsMailExchangeRecord_new2(QDnsMailExchangeRecord* other);
extern __declspec(dllexport) void QDnsMailExchangeRecord_OperatorAssign(QDnsMailExchangeRecord* self, QDnsMailExchangeRecord* other);
extern __declspec(dllexport) void QDnsMailExchangeRecord_Swap(QDnsMailExchangeRecord* self, QDnsMailExchangeRecord* other);
extern __declspec(dllexport) struct miqt_string QDnsMailExchangeRecord_Exchange(const QDnsMailExchangeRecord* self);
extern __declspec(dllexport) struct miqt_string QDnsMailExchangeRecord_Name(const QDnsMailExchangeRecord* self);
extern __declspec(dllexport) uint16_t QDnsMailExchangeRecord_Preference(const QDnsMailExchangeRecord* self);
extern __declspec(dllexport) unsigned int QDnsMailExchangeRecord_TimeToLive(const QDnsMailExchangeRecord* self);
extern __declspec(dllexport) void QDnsMailExchangeRecord_Delete(QDnsMailExchangeRecord* self, bool isSubclass);

extern __declspec(dllexport) QDnsServiceRecord* QDnsServiceRecord_new();
extern __declspec(dllexport) QDnsServiceRecord* QDnsServiceRecord_new2(QDnsServiceRecord* other);
extern __declspec(dllexport) void QDnsServiceRecord_OperatorAssign(QDnsServiceRecord* self, QDnsServiceRecord* other);
extern __declspec(dllexport) void QDnsServiceRecord_Swap(QDnsServiceRecord* self, QDnsServiceRecord* other);
extern __declspec(dllexport) struct miqt_string QDnsServiceRecord_Name(const QDnsServiceRecord* self);
extern __declspec(dllexport) uint16_t QDnsServiceRecord_Port(const QDnsServiceRecord* self);
extern __declspec(dllexport) uint16_t QDnsServiceRecord_Priority(const QDnsServiceRecord* self);
extern __declspec(dllexport) struct miqt_string QDnsServiceRecord_Target(const QDnsServiceRecord* self);
extern __declspec(dllexport) unsigned int QDnsServiceRecord_TimeToLive(const QDnsServiceRecord* self);
extern __declspec(dllexport) uint16_t QDnsServiceRecord_Weight(const QDnsServiceRecord* self);
extern __declspec(dllexport) void QDnsServiceRecord_Delete(QDnsServiceRecord* self, bool isSubclass);

extern __declspec(dllexport) QDnsTextRecord* QDnsTextRecord_new();
extern __declspec(dllexport) QDnsTextRecord* QDnsTextRecord_new2(QDnsTextRecord* other);
extern __declspec(dllexport) void QDnsTextRecord_OperatorAssign(QDnsTextRecord* self, QDnsTextRecord* other);
extern __declspec(dllexport) void QDnsTextRecord_Swap(QDnsTextRecord* self, QDnsTextRecord* other);
extern __declspec(dllexport) struct miqt_string QDnsTextRecord_Name(const QDnsTextRecord* self);
extern __declspec(dllexport) unsigned int QDnsTextRecord_TimeToLive(const QDnsTextRecord* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDnsTextRecord_Values(const QDnsTextRecord* self);
extern __declspec(dllexport) void QDnsTextRecord_Delete(QDnsTextRecord* self, bool isSubclass);

extern __declspec(dllexport) QDnsTlsAssociationRecord* QDnsTlsAssociationRecord_new();
extern __declspec(dllexport) QDnsTlsAssociationRecord* QDnsTlsAssociationRecord_new2(QDnsTlsAssociationRecord* other);
extern __declspec(dllexport) void QDnsTlsAssociationRecord_OperatorAssign(QDnsTlsAssociationRecord* self, QDnsTlsAssociationRecord* other);
extern __declspec(dllexport) void QDnsTlsAssociationRecord_Swap(QDnsTlsAssociationRecord* self, QDnsTlsAssociationRecord* other);
extern __declspec(dllexport) struct miqt_string QDnsTlsAssociationRecord_Name(const QDnsTlsAssociationRecord* self);
extern __declspec(dllexport) unsigned int QDnsTlsAssociationRecord_TimeToLive(const QDnsTlsAssociationRecord* self);
extern __declspec(dllexport) CertificateUsage QDnsTlsAssociationRecord_Usage(const QDnsTlsAssociationRecord* self);
extern __declspec(dllexport) Selector QDnsTlsAssociationRecord_Selector(const QDnsTlsAssociationRecord* self);
extern __declspec(dllexport) MatchingType QDnsTlsAssociationRecord_MatchType(const QDnsTlsAssociationRecord* self);
extern __declspec(dllexport) struct miqt_string QDnsTlsAssociationRecord_Value(const QDnsTlsAssociationRecord* self);
extern __declspec(dllexport) void QDnsTlsAssociationRecord_Delete(QDnsTlsAssociationRecord* self, bool isSubclass);

extern __declspec(dllexport) QDnsLookup* QDnsLookup_new();
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new2(Type typeVal, struct miqt_string name);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new3(Type typeVal, struct miqt_string name, QHostAddress* nameserver);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new4(Type typeVal, struct miqt_string name, QHostAddress* nameserver, uint16_t port);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new5(Type typeVal, struct miqt_string name, Protocol protocol, QHostAddress* nameserver);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new6(QObject* parent);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new7(Type typeVal, struct miqt_string name, QObject* parent);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new8(Type typeVal, struct miqt_string name, QHostAddress* nameserver, QObject* parent);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new9(Type typeVal, struct miqt_string name, QHostAddress* nameserver, uint16_t port, QObject* parent);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new10(Type typeVal, struct miqt_string name, Protocol protocol, QHostAddress* nameserver, uint16_t port);
extern __declspec(dllexport) QDnsLookup* QDnsLookup_new11(Type typeVal, struct miqt_string name, Protocol protocol, QHostAddress* nameserver, uint16_t port, QObject* parent);
extern __declspec(dllexport) void QDnsLookup_virtbase(QDnsLookup* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QDnsLookup_MetaObject(const QDnsLookup* self);
extern __declspec(dllexport) void* QDnsLookup_Metacast(QDnsLookup* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QDnsLookup_Tr(const char* s);
extern __declspec(dllexport) bool QDnsLookup_IsAuthenticData(const QDnsLookup* self);
extern __declspec(dllexport) Error QDnsLookup_Error(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_string QDnsLookup_ErrorString(const QDnsLookup* self);
extern __declspec(dllexport) bool QDnsLookup_IsFinished(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_string QDnsLookup_Name(const QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_SetName(QDnsLookup* self, struct miqt_string name);
extern __declspec(dllexport) Type QDnsLookup_Type(const QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_SetType(QDnsLookup* self, int typeVal);
extern __declspec(dllexport) QHostAddress* QDnsLookup_Nameserver(const QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_SetNameserver(QDnsLookup* self, QHostAddress* nameserver);
extern __declspec(dllexport) uint16_t QDnsLookup_NameserverPort(const QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_SetNameserverPort(QDnsLookup* self, uint16_t port);
extern __declspec(dllexport) Protocol QDnsLookup_NameserverProtocol(const QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_SetNameserverProtocol(QDnsLookup* self, Protocol protocol);
extern __declspec(dllexport) void QDnsLookup_SetNameserver2(QDnsLookup* self, Protocol protocol, QHostAddress* nameserver);
extern __declspec(dllexport) void QDnsLookup_SetNameserver3(QDnsLookup* self, QHostAddress* nameserver, uint16_t port);
extern __declspec(dllexport) struct miqt_array /* of QDnsDomainNameRecord* */  QDnsLookup_CanonicalNameRecords(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_array /* of QDnsHostAddressRecord* */  QDnsLookup_HostAddressRecords(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_array /* of QDnsMailExchangeRecord* */  QDnsLookup_MailExchangeRecords(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_array /* of QDnsDomainNameRecord* */  QDnsLookup_NameServerRecords(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_array /* of QDnsDomainNameRecord* */  QDnsLookup_PointerRecords(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_array /* of QDnsServiceRecord* */  QDnsLookup_ServiceRecords(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_array /* of QDnsTextRecord* */  QDnsLookup_TextRecords(const QDnsLookup* self);
extern __declspec(dllexport) struct miqt_array /* of QDnsTlsAssociationRecord* */  QDnsLookup_TlsAssociationRecords(const QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_SetSslConfiguration(QDnsLookup* self, QSslConfiguration* sslConfiguration);
extern __declspec(dllexport) QSslConfiguration* QDnsLookup_SslConfiguration(const QDnsLookup* self);
extern __declspec(dllexport) bool QDnsLookup_IsProtocolSupported(Protocol protocol);
extern __declspec(dllexport) uint16_t QDnsLookup_DefaultPortForProtocol(Protocol protocol);
extern __declspec(dllexport) void QDnsLookup_Abort(QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_Lookup(QDnsLookup* self);
extern __declspec(dllexport) void QDnsLookup_Finished(QDnsLookup* self);
void QDnsLookup_connect_Finished(QDnsLookup* self, intptr_t slot);
extern __declspec(dllexport) void QDnsLookup_NameChanged(QDnsLookup* self, struct miqt_string name);
void QDnsLookup_connect_NameChanged(QDnsLookup* self, intptr_t slot);
extern __declspec(dllexport) void QDnsLookup_TypeChanged(QDnsLookup* self, int typeVal);
void QDnsLookup_connect_TypeChanged(QDnsLookup* self, intptr_t slot);
extern __declspec(dllexport) void QDnsLookup_NameserverChanged(QDnsLookup* self, QHostAddress* nameserver);
void QDnsLookup_connect_NameserverChanged(QDnsLookup* self, intptr_t slot);
extern __declspec(dllexport) void QDnsLookup_NameserverPortChanged(QDnsLookup* self, uint16_t port);
void QDnsLookup_connect_NameserverPortChanged(QDnsLookup* self, intptr_t slot);
extern __declspec(dllexport) void QDnsLookup_NameserverProtocolChanged(QDnsLookup* self, uint8_t protocol);
void QDnsLookup_connect_NameserverProtocolChanged(QDnsLookup* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QDnsLookup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QDnsLookup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QDnsLookup_SetNameserver32(QDnsLookup* self, Protocol protocol, QHostAddress* nameserver, uint16_t port);
extern __declspec(dllexport) void QDnsLookup_override_virtual_Event(void* self, intptr_t slot);
bool QDnsLookup_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QDnsLookup_override_virtual_EventFilter(void* self, intptr_t slot);
bool QDnsLookup_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QDnsLookup_override_virtual_TimerEvent(void* self, intptr_t slot);
void QDnsLookup_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QDnsLookup_override_virtual_ChildEvent(void* self, intptr_t slot);
void QDnsLookup_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QDnsLookup_override_virtual_CustomEvent(void* self, intptr_t slot);
void QDnsLookup_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QDnsLookup_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QDnsLookup_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QDnsLookup_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QDnsLookup_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QDnsLookup_Delete(QDnsLookup* self, bool isSubclass);

} 
