#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKPROXY_H
#define MIQT_QT_NETWORK_GEN_QNETWORKPROXY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHttpHeaders QHttpHeaders;
typedef struct QNetworkProxy QNetworkProxy;
typedef struct QNetworkProxyFactory QNetworkProxyFactory;
typedef struct QNetworkProxyQuery QNetworkProxyQuery;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new();
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new2(QUrl* requestUrl);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new3(struct miqt_string hostname, int port);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new4(uint16_t bindPort);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new5(QNetworkProxyQuery* other);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new6(QUrl* requestUrl, QueryType queryType);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new7(struct miqt_string hostname, int port, struct miqt_string protocolTag);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new8(struct miqt_string hostname, int port, struct miqt_string protocolTag, QueryType queryType);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new9(uint16_t bindPort, struct miqt_string protocolTag);
extern __declspec(dllexport) 
QNetworkProxyQuery* QNetworkProxyQuery_new10(uint16_t bindPort, struct miqt_string protocolTag, QueryType queryType);
extern __declspec(dllexport) 
void QNetworkProxyQuery_OperatorAssign(QNetworkProxyQuery* self, QNetworkProxyQuery* other);
extern __declspec(dllexport) 
void QNetworkProxyQuery_Swap(QNetworkProxyQuery* self, QNetworkProxyQuery* other);
extern __declspec(dllexport) 
bool QNetworkProxyQuery_OperatorEqual(const QNetworkProxyQuery* self, QNetworkProxyQuery* other);
extern __declspec(dllexport) 
bool QNetworkProxyQuery_OperatorNotEqual(const QNetworkProxyQuery* self, QNetworkProxyQuery* other);
extern __declspec(dllexport) 
QueryType QNetworkProxyQuery_QueryType(const QNetworkProxyQuery* self);
extern __declspec(dllexport) 
void QNetworkProxyQuery_SetQueryType(QNetworkProxyQuery* self, QueryType typeVal);
extern __declspec(dllexport) 
int QNetworkProxyQuery_PeerPort(const QNetworkProxyQuery* self);
extern __declspec(dllexport) 
void QNetworkProxyQuery_SetPeerPort(QNetworkProxyQuery* self, int port);
extern __declspec(dllexport) 
struct miqt_string QNetworkProxyQuery_PeerHostName(const QNetworkProxyQuery* self);
extern __declspec(dllexport) 
void QNetworkProxyQuery_SetPeerHostName(QNetworkProxyQuery* self, struct miqt_string hostname);
extern __declspec(dllexport) 
int QNetworkProxyQuery_LocalPort(const QNetworkProxyQuery* self);
extern __declspec(dllexport) 
void QNetworkProxyQuery_SetLocalPort(QNetworkProxyQuery* self, int port);
extern __declspec(dllexport) 
struct miqt_string QNetworkProxyQuery_ProtocolTag(const QNetworkProxyQuery* self);
extern __declspec(dllexport) 
void QNetworkProxyQuery_SetProtocolTag(QNetworkProxyQuery* self, struct miqt_string protocolTag);
extern __declspec(dllexport) 
QUrl* QNetworkProxyQuery_Url(const QNetworkProxyQuery* self);
extern __declspec(dllexport) 
void QNetworkProxyQuery_SetUrl(QNetworkProxyQuery* self, QUrl* url);
extern __declspec(dllexport) 
void QNetworkProxyQuery_Delete(QNetworkProxyQuery* self, bool isSubclass);

extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_new();
extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_new2(ProxyType typeVal);
extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_new3(QNetworkProxy* other);
extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_new4(ProxyType typeVal, struct miqt_string hostName);
extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_new5(ProxyType typeVal, struct miqt_string hostName, uint16_t port);
extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_new6(ProxyType typeVal, struct miqt_string hostName, uint16_t port, struct miqt_string user);
extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_new7(ProxyType typeVal, struct miqt_string hostName, uint16_t port, struct miqt_string user, struct miqt_string password);
extern __declspec(dllexport) 
void QNetworkProxy_OperatorAssign(QNetworkProxy* self, QNetworkProxy* other);
extern __declspec(dllexport) 
void QNetworkProxy_Swap(QNetworkProxy* self, QNetworkProxy* other);
extern __declspec(dllexport) 
bool QNetworkProxy_OperatorEqual(const QNetworkProxy* self, QNetworkProxy* other);
extern __declspec(dllexport) 
bool QNetworkProxy_OperatorNotEqual(const QNetworkProxy* self, QNetworkProxy* other);
extern __declspec(dllexport) 
void QNetworkProxy_SetType(QNetworkProxy* self, int typeVal);
extern __declspec(dllexport) 
int QNetworkProxy_Type(const QNetworkProxy* self);
extern __declspec(dllexport) 
void QNetworkProxy_SetCapabilities(QNetworkProxy* self, Capabilities capab);
extern __declspec(dllexport) 
Capabilities QNetworkProxy_Capabilities(const QNetworkProxy* self);
extern __declspec(dllexport) 
bool QNetworkProxy_IsCachingProxy(const QNetworkProxy* self);
extern __declspec(dllexport) 
bool QNetworkProxy_IsTransparentProxy(const QNetworkProxy* self);
extern __declspec(dllexport) 
void QNetworkProxy_SetUser(QNetworkProxy* self, struct miqt_string userName);
extern __declspec(dllexport) 
struct miqt_string QNetworkProxy_User(const QNetworkProxy* self);
extern __declspec(dllexport) 
void QNetworkProxy_SetPassword(QNetworkProxy* self, struct miqt_string password);
extern __declspec(dllexport) 
struct miqt_string QNetworkProxy_Password(const QNetworkProxy* self);
extern __declspec(dllexport) 
void QNetworkProxy_SetHostName(QNetworkProxy* self, struct miqt_string hostName);
extern __declspec(dllexport) 
struct miqt_string QNetworkProxy_HostName(const QNetworkProxy* self);
extern __declspec(dllexport) 
void QNetworkProxy_SetPort(QNetworkProxy* self, uint16_t port);
extern __declspec(dllexport) 
uint16_t QNetworkProxy_Port(const QNetworkProxy* self);
extern __declspec(dllexport) 
void QNetworkProxy_SetApplicationProxy(QNetworkProxy* proxy);
extern __declspec(dllexport) 
QNetworkProxy* QNetworkProxy_ApplicationProxy();
extern __declspec(dllexport) 
QHttpHeaders* QNetworkProxy_Headers(const QNetworkProxy* self);
extern __declspec(dllexport) 
void QNetworkProxy_SetHeaders(QNetworkProxy* self, QHttpHeaders* newHeaders);
extern __declspec(dllexport) 
QVariant* QNetworkProxy_Header(const QNetworkProxy* self, int header);
extern __declspec(dllexport) 
void QNetworkProxy_SetHeader(QNetworkProxy* self, int header, QVariant* value);
extern __declspec(dllexport) 
bool QNetworkProxy_HasRawHeader(const QNetworkProxy* self, struct miqt_string headerName);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QNetworkProxy_RawHeaderList(const QNetworkProxy* self);
extern __declspec(dllexport) 
struct miqt_string QNetworkProxy_RawHeader(const QNetworkProxy* self, struct miqt_string headerName);
extern __declspec(dllexport) 
void QNetworkProxy_SetRawHeader(QNetworkProxy* self, struct miqt_string headerName, struct miqt_string value);
extern __declspec(dllexport) 
void QNetworkProxy_Delete(QNetworkProxy* self, bool isSubclass);

extern __declspec(dllexport) 
QNetworkProxyFactory* QNetworkProxyFactory_new();
extern __declspec(dllexport) 
struct miqt_array /* of QNetworkProxy* */  QNetworkProxyFactory_QueryProxy(QNetworkProxyFactory* self, QNetworkProxyQuery* query);
extern __declspec(dllexport) 
bool QNetworkProxyFactory_UsesSystemConfiguration();
extern __declspec(dllexport) 
void QNetworkProxyFactory_SetUseSystemConfiguration(bool enable);
extern __declspec(dllexport) 
void QNetworkProxyFactory_SetApplicationProxyFactory(QNetworkProxyFactory* factory);
extern __declspec(dllexport) 
struct miqt_array /* of QNetworkProxy* */  QNetworkProxyFactory_ProxyForQuery(QNetworkProxyQuery* query);
extern __declspec(dllexport) 
struct miqt_array /* of QNetworkProxy* */  QNetworkProxyFactory_SystemProxyForQuery();
extern __declspec(dllexport) 
void QNetworkProxyFactory_OperatorAssign(QNetworkProxyFactory* self, QNetworkProxyFactory* param1);
extern __declspec(dllexport) 
struct miqt_array /* of QNetworkProxy* */  QNetworkProxyFactory_SystemProxyForQuery1(QNetworkProxyQuery* query);
extern __declspec(dllexport) 
void QNetworkProxyFactory_Delete(QNetworkProxyFactory* self, bool isSubclass);

}
