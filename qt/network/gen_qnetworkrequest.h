#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKREQUEST_H
#define MIQT_QT_NETWORK_GEN_QNETWORKREQUEST_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QHttp1Configuration QHttp1Configuration;
typedef struct QHttp2Configuration QHttp2Configuration;
typedef struct QHttpHeaders QHttpHeaders;
typedef struct QNetworkRequest QNetworkRequest;
typedef struct QObject QObject;
typedef struct QSslConfiguration QSslConfiguration;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNetworkRequest* QNetworkRequest_new();
extern __declspec(dllexport) QNetworkRequest* QNetworkRequest_new2(QUrl* url);
extern __declspec(dllexport) QNetworkRequest* QNetworkRequest_new3(QNetworkRequest* other);
extern __declspec(dllexport) void QNetworkRequest_OperatorAssign(QNetworkRequest* self, QNetworkRequest* other);
extern __declspec(dllexport) void QNetworkRequest_Swap(QNetworkRequest* self, QNetworkRequest* other);
extern __declspec(dllexport) bool QNetworkRequest_OperatorEqual(const QNetworkRequest* self, QNetworkRequest* other);
extern __declspec(dllexport) bool QNetworkRequest_OperatorNotEqual(const QNetworkRequest* self, QNetworkRequest* other);
extern __declspec(dllexport) QUrl* QNetworkRequest_Url(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetUrl(QNetworkRequest* self, QUrl* url);
extern __declspec(dllexport) QHttpHeaders* QNetworkRequest_Headers(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetHeaders(QNetworkRequest* self, QHttpHeaders* newHeaders);
extern __declspec(dllexport) QVariant* QNetworkRequest_Header(const QNetworkRequest* self, KnownHeaders header);
extern __declspec(dllexport) void QNetworkRequest_SetHeader(QNetworkRequest* self, KnownHeaders header, QVariant* value);
extern __declspec(dllexport) bool QNetworkRequest_HasRawHeader(const QNetworkRequest* self, QAnyStringView* headerName);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QNetworkRequest_RawHeaderList(const QNetworkRequest* self);
extern __declspec(dllexport) struct miqt_string QNetworkRequest_RawHeader(const QNetworkRequest* self, QAnyStringView* headerName);
extern __declspec(dllexport) void QNetworkRequest_SetRawHeader(QNetworkRequest* self, struct miqt_string headerName, struct miqt_string value);
extern __declspec(dllexport) QVariant* QNetworkRequest_Attribute(const QNetworkRequest* self, Attribute code);
extern __declspec(dllexport) void QNetworkRequest_SetAttribute(QNetworkRequest* self, Attribute code, QVariant* value);
extern __declspec(dllexport) QSslConfiguration* QNetworkRequest_SslConfiguration(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetSslConfiguration(QNetworkRequest* self, QSslConfiguration* configuration);
extern __declspec(dllexport) void QNetworkRequest_SetOriginatingObject(QNetworkRequest* self, QObject* object);
extern __declspec(dllexport) QObject* QNetworkRequest_OriginatingObject(const QNetworkRequest* self);
extern __declspec(dllexport) Priority QNetworkRequest_Priority(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetPriority(QNetworkRequest* self, Priority priority);
extern __declspec(dllexport) int QNetworkRequest_MaximumRedirectsAllowed(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetMaximumRedirectsAllowed(QNetworkRequest* self, int maximumRedirectsAllowed);
extern __declspec(dllexport) struct miqt_string QNetworkRequest_PeerVerifyName(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetPeerVerifyName(QNetworkRequest* self, struct miqt_string peerName);
extern __declspec(dllexport) QHttp1Configuration* QNetworkRequest_Http1Configuration(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetHttp1Configuration(QNetworkRequest* self, QHttp1Configuration* configuration);
extern __declspec(dllexport) QHttp2Configuration* QNetworkRequest_Http2Configuration(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetHttp2Configuration(QNetworkRequest* self, QHttp2Configuration* configuration);
extern __declspec(dllexport) long long QNetworkRequest_DecompressedSafetyCheckThreshold(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetDecompressedSafetyCheckThreshold(QNetworkRequest* self, long long threshold);
extern __declspec(dllexport) int QNetworkRequest_TransferTimeout(const QNetworkRequest* self);
extern __declspec(dllexport) void QNetworkRequest_SetTransferTimeout(QNetworkRequest* self, int timeout);
extern __declspec(dllexport) void QNetworkRequest_SetTransferTimeout2(QNetworkRequest* self);
extern __declspec(dllexport) QVariant* QNetworkRequest_Attribute2(const QNetworkRequest* self, Attribute code, QVariant* defaultValue);
extern __declspec(dllexport) void QNetworkRequest_Delete(QNetworkRequest* self, bool isSubclass);

} 
