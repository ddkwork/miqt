#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKREPLY_H
#define MIQT_QT_NETWORK_GEN_QNETWORKREPLY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QHttpHeaders QHttpHeaders;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkAccessManager QNetworkAccessManager;
typedef struct QNetworkReply QNetworkReply;
typedef struct QNetworkRequest QNetworkRequest;
typedef struct QObject QObject;
typedef struct QSslConfiguration QSslConfiguration;
typedef struct QSslError QSslError;
typedef struct QSslPreSharedKeyAuthenticator QSslPreSharedKeyAuthenticator;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QNetworkReply_virtbase(QNetworkReply* src
, QIODevice** outptr_QIODevice
);
extern __declspec(dllexport) 
QMetaObject* QNetworkReply_MetaObject(const QNetworkReply* self);
extern __declspec(dllexport) 
void* QNetworkReply_Metacast(QNetworkReply* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QNetworkReply_Tr(const char* s);
extern __declspec(dllexport) 
void QNetworkReply_Close(QNetworkReply* self);
extern __declspec(dllexport) 
bool QNetworkReply_IsSequential(const QNetworkReply* self);
extern __declspec(dllexport) 
long long QNetworkReply_ReadBufferSize(const QNetworkReply* self);
extern __declspec(dllexport) 
void QNetworkReply_SetReadBufferSize(QNetworkReply* self, long long size);
extern __declspec(dllexport) 
QNetworkAccessManager* QNetworkReply_Manager(const QNetworkReply* self);
extern __declspec(dllexport) 
int QNetworkReply_Operation(const QNetworkReply* self);
extern __declspec(dllexport) 
QNetworkRequest* QNetworkReply_Request(const QNetworkReply* self);
extern __declspec(dllexport) 
NetworkError QNetworkReply_Error(const QNetworkReply* self);
extern __declspec(dllexport) 
bool QNetworkReply_IsFinished(const QNetworkReply* self);
extern __declspec(dllexport) 
bool QNetworkReply_IsRunning(const QNetworkReply* self);
extern __declspec(dllexport) 
QUrl* QNetworkReply_Url(const QNetworkReply* self);
extern __declspec(dllexport) 
QVariant* QNetworkReply_Header(const QNetworkReply* self, int header);
extern __declspec(dllexport) 
bool QNetworkReply_HasRawHeader(const QNetworkReply* self, QAnyStringView* headerName);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QNetworkReply_RawHeaderList(const QNetworkReply* self);
extern __declspec(dllexport) 
struct miqt_string QNetworkReply_RawHeader(const QNetworkReply* self, QAnyStringView* headerName);
extern __declspec(dllexport) 
struct miqt_array /* of RawHeaderPair */  QNetworkReply_RawHeaderPairs(const QNetworkReply* self);
extern __declspec(dllexport) 
QHttpHeaders* QNetworkReply_Headers(const QNetworkReply* self);
extern __declspec(dllexport) 
QVariant* QNetworkReply_Attribute(const QNetworkReply* self, int code);
extern __declspec(dllexport) 
QSslConfiguration* QNetworkReply_SslConfiguration(const QNetworkReply* self);
extern __declspec(dllexport) 
void QNetworkReply_SetSslConfiguration(QNetworkReply* self, QSslConfiguration* configuration);
extern __declspec(dllexport) 
void QNetworkReply_IgnoreSslErrors(QNetworkReply* self, struct miqt_array /* of QSslError* */  errors);
extern __declspec(dllexport) 
void QNetworkReply_Abort(QNetworkReply* self);
extern __declspec(dllexport) 
void QNetworkReply_IgnoreSslErrors2(QNetworkReply* self);
extern __declspec(dllexport) 
void QNetworkReply_SocketStartedConnecting(QNetworkReply* self);
void QNetworkReply_connect_SocketStartedConnecting(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_RequestSent(QNetworkReply* self);
void QNetworkReply_connect_RequestSent(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_MetaDataChanged(QNetworkReply* self);
void QNetworkReply_connect_MetaDataChanged(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_Finished(QNetworkReply* self);
void QNetworkReply_connect_Finished(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_ErrorOccurred(QNetworkReply* self, int param1);
void QNetworkReply_connect_ErrorOccurred(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_Encrypted(QNetworkReply* self);
void QNetworkReply_connect_Encrypted(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_SslErrors(QNetworkReply* self, struct miqt_array /* of QSslError* */  errors);
void QNetworkReply_connect_SslErrors(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_PreSharedKeyAuthenticationRequired(QNetworkReply* self, QSslPreSharedKeyAuthenticator* authenticator);
void QNetworkReply_connect_PreSharedKeyAuthenticationRequired(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_Redirected(QNetworkReply* self, QUrl* url);
void QNetworkReply_connect_Redirected(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_RedirectAllowed(QNetworkReply* self);
void QNetworkReply_connect_RedirectAllowed(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_UploadProgress(QNetworkReply* self, long long bytesSent, long long bytesTotal);
void QNetworkReply_connect_UploadProgress(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
void QNetworkReply_DownloadProgress(QNetworkReply* self, long long bytesReceived, long long bytesTotal);
void QNetworkReply_connect_DownloadProgress(QNetworkReply* self, intptr_t slot);
extern __declspec(dllexport) 
long long QNetworkReply_WriteData(QNetworkReply* self, const char* data, long long lenVal);
extern __declspec(dllexport) 
void QNetworkReply_SslConfigurationImplementation(const QNetworkReply* self, QSslConfiguration* param1);
extern __declspec(dllexport) 
void QNetworkReply_SetSslConfigurationImplementation(QNetworkReply* self, QSslConfiguration* sslConfigurationImplementation);
extern __declspec(dllexport) 
void QNetworkReply_IgnoreSslErrorsImplementation(QNetworkReply* self, struct miqt_array /* of QSslError* */  param1);
extern __declspec(dllexport) 
struct miqt_string QNetworkReply_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QNetworkReply_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QNetworkReply_Delete(QNetworkReply* self, bool isSubclass);

}
