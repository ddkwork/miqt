#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKREQUESTFACTORY_H
#define MIQT_QT_NETWORK_GEN_QNETWORKREQUESTFACTORY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHttpHeaders QHttpHeaders;
typedef struct QNetworkRequest QNetworkRequest;
typedef struct QNetworkRequestFactory QNetworkRequestFactory;
typedef struct QSslConfiguration QSslConfiguration;
typedef struct QUrl QUrl;
typedef struct QUrlQuery QUrlQuery;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNetworkRequestFactory* QNetworkRequestFactory_new();
extern __declspec(dllexport) QNetworkRequestFactory* QNetworkRequestFactory_new2(QUrl* baseUrl);
extern __declspec(dllexport) QNetworkRequestFactory* QNetworkRequestFactory_new3(QNetworkRequestFactory* other);
extern __declspec(dllexport) void QNetworkRequestFactory_OperatorAssign(QNetworkRequestFactory* self, QNetworkRequestFactory* other);
extern __declspec(dllexport) void QNetworkRequestFactory_Swap(QNetworkRequestFactory* self, QNetworkRequestFactory* other);
extern __declspec(dllexport) QUrl* QNetworkRequestFactory_BaseUrl(const QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_SetBaseUrl(QNetworkRequestFactory* self, QUrl* url);
extern __declspec(dllexport) QSslConfiguration* QNetworkRequestFactory_SslConfiguration(const QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_SetSslConfiguration(QNetworkRequestFactory* self, QSslConfiguration* configuration);
extern __declspec(dllexport) QNetworkRequest* QNetworkRequestFactory_CreateRequest(const QNetworkRequestFactory* self);
extern __declspec(dllexport) QNetworkRequest* QNetworkRequestFactory_CreateRequestWithQuery(const QNetworkRequestFactory* self, QUrlQuery* query);
extern __declspec(dllexport) QNetworkRequest* QNetworkRequestFactory_CreateRequestWithPath(const QNetworkRequestFactory* self, struct miqt_string path);
extern __declspec(dllexport) QNetworkRequest* QNetworkRequestFactory_CreateRequest2(const QNetworkRequestFactory* self, struct miqt_string path, QUrlQuery* query);
extern __declspec(dllexport) void QNetworkRequestFactory_SetCommonHeaders(QNetworkRequestFactory* self, QHttpHeaders* headers);
extern __declspec(dllexport) QHttpHeaders* QNetworkRequestFactory_CommonHeaders(const QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_ClearCommonHeaders(QNetworkRequestFactory* self);
extern __declspec(dllexport) struct miqt_string QNetworkRequestFactory_BearerToken(const QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_SetBearerToken(QNetworkRequestFactory* self, struct miqt_string token);
extern __declspec(dllexport) void QNetworkRequestFactory_ClearBearerToken(QNetworkRequestFactory* self);
extern __declspec(dllexport) struct miqt_string QNetworkRequestFactory_UserName(const QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_SetUserName(QNetworkRequestFactory* self, struct miqt_string userName);
extern __declspec(dllexport) void QNetworkRequestFactory_ClearUserName(QNetworkRequestFactory* self);
extern __declspec(dllexport) struct miqt_string QNetworkRequestFactory_Password(const QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_SetPassword(QNetworkRequestFactory* self, struct miqt_string password);
extern __declspec(dllexport) void QNetworkRequestFactory_ClearPassword(QNetworkRequestFactory* self);
extern __declspec(dllexport) QUrlQuery* QNetworkRequestFactory_QueryParameters(const QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_SetQueryParameters(QNetworkRequestFactory* self, QUrlQuery* query);
extern __declspec(dllexport) void QNetworkRequestFactory_ClearQueryParameters(QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_SetPriority(QNetworkRequestFactory* self, int priority);
extern __declspec(dllexport) int QNetworkRequestFactory_Priority(const QNetworkRequestFactory* self);
extern __declspec(dllexport) QVariant* QNetworkRequestFactory_Attribute(const QNetworkRequestFactory* self, int attribute);
extern __declspec(dllexport) QVariant* QNetworkRequestFactory_Attribute2(const QNetworkRequestFactory* self, int attribute, QVariant* defaultValue);
extern __declspec(dllexport) void QNetworkRequestFactory_SetAttribute(QNetworkRequestFactory* self, int attribute, QVariant* value);
extern __declspec(dllexport) void QNetworkRequestFactory_ClearAttribute(QNetworkRequestFactory* self, int attribute);
extern __declspec(dllexport) void QNetworkRequestFactory_ClearAttributes(QNetworkRequestFactory* self);
extern __declspec(dllexport) void QNetworkRequestFactory_Delete(QNetworkRequestFactory* self, bool isSubclass);

} 
