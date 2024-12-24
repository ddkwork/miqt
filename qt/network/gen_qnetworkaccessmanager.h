#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKACCESSMANAGER_H
#define MIQT_QT_NETWORK_GEN_QNETWORKACCESSMANAGER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractNetworkCache;
class QAuthenticator;
class QChildEvent;
class QEvent;
class QHstsPolicy;
class QHttpMultiPart;
class QIODevice;
class QMetaMethod;
class QMetaObject;
class QNetworkAccessManager;
class QNetworkCookieJar;
class QNetworkProxy;
class QNetworkProxyFactory;
class QNetworkReply;
class QNetworkRequest;
class QObject;
class QSslConfiguration;
class QSslError;
class QSslPreSharedKeyAuthenticator;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QAbstractNetworkCache QAbstractNetworkCache;
typedef struct QAuthenticator QAuthenticator;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QHstsPolicy QHstsPolicy;
typedef struct QHttpMultiPart QHttpMultiPart;
typedef struct QIODevice QIODevice;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkAccessManager QNetworkAccessManager;
typedef struct QNetworkCookieJar QNetworkCookieJar;
typedef struct QNetworkProxy QNetworkProxy;
typedef struct QNetworkProxyFactory QNetworkProxyFactory;
typedef struct QNetworkReply QNetworkReply;
typedef struct QNetworkRequest QNetworkRequest;
typedef struct QObject QObject;
typedef struct QSslConfiguration QSslConfiguration;
typedef struct QSslError QSslError;
typedef struct QSslPreSharedKeyAuthenticator QSslPreSharedKeyAuthenticator;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNetworkAccessManager* QNetworkAccessManager_new();
extern __declspec(dllexport) QNetworkAccessManager* QNetworkAccessManager_new2(QObject* parent);
extern __declspec(dllexport) void QNetworkAccessManager_virtbase(QNetworkAccessManager* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QNetworkAccessManager_MetaObject(const QNetworkAccessManager* self);
extern __declspec(dllexport) void* QNetworkAccessManager_Metacast(QNetworkAccessManager* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QNetworkAccessManager_Tr(const char* s);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QNetworkAccessManager_SupportedSchemes(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_ClearAccessCache(QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_ClearConnectionCache(QNetworkAccessManager* self);
extern __declspec(dllexport) QNetworkProxy* QNetworkAccessManager_Proxy(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_SetProxy(QNetworkAccessManager* self, QNetworkProxy* proxy);
extern __declspec(dllexport) QNetworkProxyFactory* QNetworkAccessManager_ProxyFactory(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_SetProxyFactory(QNetworkAccessManager* self, QNetworkProxyFactory* factory);
extern __declspec(dllexport) QAbstractNetworkCache* QNetworkAccessManager_Cache(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_SetCache(QNetworkAccessManager* self, QAbstractNetworkCache* cache);
extern __declspec(dllexport) QNetworkCookieJar* QNetworkAccessManager_CookieJar(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_SetCookieJar(QNetworkAccessManager* self, QNetworkCookieJar* cookieJar);
extern __declspec(dllexport) void QNetworkAccessManager_SetStrictTransportSecurityEnabled(QNetworkAccessManager* self, bool enabled);
extern __declspec(dllexport) bool QNetworkAccessManager_IsStrictTransportSecurityEnabled(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_EnableStrictTransportSecurityStore(QNetworkAccessManager* self, bool enabled);
extern __declspec(dllexport) bool QNetworkAccessManager_IsStrictTransportSecurityStoreEnabled(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_AddStrictTransportSecurityHosts(QNetworkAccessManager* self, struct miqt_array /* of QHstsPolicy* */  knownHosts);
extern __declspec(dllexport) struct miqt_array /* of QHstsPolicy* */  QNetworkAccessManager_StrictTransportSecurityHosts(const QNetworkAccessManager* self);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Head(QNetworkAccessManager* self, QNetworkRequest* request);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Get(QNetworkAccessManager* self, QNetworkRequest* request);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Get2(QNetworkAccessManager* self, QNetworkRequest* request, QIODevice* data);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Get3(QNetworkAccessManager* self, QNetworkRequest* request, struct miqt_string data);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Post(QNetworkAccessManager* self, QNetworkRequest* request, QIODevice* data);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Post2(QNetworkAccessManager* self, QNetworkRequest* request, struct miqt_string data);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Put(QNetworkAccessManager* self, QNetworkRequest* request, QIODevice* data);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Put2(QNetworkAccessManager* self, QNetworkRequest* request, struct miqt_string data);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_DeleteResource(QNetworkAccessManager* self, QNetworkRequest* request);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_SendCustomRequest(QNetworkAccessManager* self, QNetworkRequest* request, struct miqt_string verb);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_SendCustomRequest2(QNetworkAccessManager* self, QNetworkRequest* request, struct miqt_string verb, struct miqt_string data);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Post4(QNetworkAccessManager* self, QNetworkRequest* request, QHttpMultiPart* multiPart);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_Put4(QNetworkAccessManager* self, QNetworkRequest* request, QHttpMultiPart* multiPart);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_SendCustomRequest3(QNetworkAccessManager* self, QNetworkRequest* request, struct miqt_string verb, QHttpMultiPart* multiPart);
extern __declspec(dllexport) void QNetworkAccessManager_ConnectToHostEncrypted(QNetworkAccessManager* self, struct miqt_string hostName);
extern __declspec(dllexport) void QNetworkAccessManager_ConnectToHostEncrypted2(QNetworkAccessManager* self, struct miqt_string hostName, uint16_t port, QSslConfiguration* sslConfiguration, struct miqt_string peerName);
extern __declspec(dllexport) void QNetworkAccessManager_ConnectToHost(QNetworkAccessManager* self, struct miqt_string hostName);
extern __declspec(dllexport) void QNetworkAccessManager_SetRedirectPolicy(QNetworkAccessManager* self, int policy);
extern __declspec(dllexport) int QNetworkAccessManager_RedirectPolicy(const QNetworkAccessManager* self);
extern __declspec(dllexport) bool QNetworkAccessManager_AutoDeleteReplies(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_SetAutoDeleteReplies(QNetworkAccessManager* self, bool autoDelete);
extern __declspec(dllexport) int QNetworkAccessManager_TransferTimeout(const QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_SetTransferTimeout(QNetworkAccessManager* self, int timeout);
extern __declspec(dllexport) void QNetworkAccessManager_SetTransferTimeout2(QNetworkAccessManager* self);
extern __declspec(dllexport) void QNetworkAccessManager_ProxyAuthenticationRequired(QNetworkAccessManager* self, QNetworkProxy* proxy, QAuthenticator* authenticator);
void QNetworkAccessManager_connect_ProxyAuthenticationRequired(QNetworkAccessManager* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkAccessManager_AuthenticationRequired(QNetworkAccessManager* self, QNetworkReply* reply, QAuthenticator* authenticator);
void QNetworkAccessManager_connect_AuthenticationRequired(QNetworkAccessManager* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkAccessManager_Finished(QNetworkAccessManager* self, QNetworkReply* reply);
void QNetworkAccessManager_connect_Finished(QNetworkAccessManager* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkAccessManager_Encrypted(QNetworkAccessManager* self, QNetworkReply* reply);
void QNetworkAccessManager_connect_Encrypted(QNetworkAccessManager* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkAccessManager_SslErrors(QNetworkAccessManager* self, QNetworkReply* reply, struct miqt_array /* of QSslError* */  errors);
void QNetworkAccessManager_connect_SslErrors(QNetworkAccessManager* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkAccessManager_PreSharedKeyAuthenticationRequired(QNetworkAccessManager* self, QNetworkReply* reply, QSslPreSharedKeyAuthenticator* authenticator);
void QNetworkAccessManager_connect_PreSharedKeyAuthenticationRequired(QNetworkAccessManager* self, intptr_t slot);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_CreateRequest(QNetworkAccessManager* self, Operation op, QNetworkRequest* request, QIODevice* outgoingData);
extern __declspec(dllexport) struct miqt_string QNetworkAccessManager_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QNetworkAccessManager_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QNetworkAccessManager_EnableStrictTransportSecurityStore2(QNetworkAccessManager* self, bool enabled, struct miqt_string storeDir);
extern __declspec(dllexport) QNetworkReply* QNetworkAccessManager_SendCustomRequest32(QNetworkAccessManager* self, QNetworkRequest* request, struct miqt_string verb, QIODevice* data);
extern __declspec(dllexport) void QNetworkAccessManager_ConnectToHostEncrypted22(QNetworkAccessManager* self, struct miqt_string hostName, uint16_t port);
extern __declspec(dllexport) void QNetworkAccessManager_ConnectToHostEncrypted3(QNetworkAccessManager* self, struct miqt_string hostName, uint16_t port, QSslConfiguration* sslConfiguration);
extern __declspec(dllexport) void QNetworkAccessManager_ConnectToHost2(QNetworkAccessManager* self, struct miqt_string hostName, uint16_t port);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_SupportedSchemes(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QNetworkAccessManager_virtualbase_SupportedSchemes(const void* self);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_CreateRequest(void* self, intptr_t slot);
QNetworkReply* QNetworkAccessManager_virtualbase_CreateRequest(void* self, Operation op, QNetworkRequest* request, QIODevice* outgoingData);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_Event(void* self, intptr_t slot);
bool QNetworkAccessManager_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_EventFilter(void* self, intptr_t slot);
bool QNetworkAccessManager_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_TimerEvent(void* self, intptr_t slot);
void QNetworkAccessManager_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_ChildEvent(void* self, intptr_t slot);
void QNetworkAccessManager_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_CustomEvent(void* self, intptr_t slot);
void QNetworkAccessManager_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QNetworkAccessManager_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QNetworkAccessManager_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QNetworkAccessManager_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QNetworkAccessManager_Delete(QNetworkAccessManager* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
