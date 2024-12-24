#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLSERVER_H
#define MIQT_QT_NETWORK_GEN_QSSLSERVER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMetaObject;
class QObject;
class QSslConfiguration;
class QSslError;
class QSslPreSharedKeyAuthenticator;
class QSslServer;
class QSslSocket;
class QTcpServer;
class QTcpSocket;
class _GUID;
class type_info;
#else
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSslConfiguration QSslConfiguration;
typedef struct QSslError QSslError;
typedef struct QSslPreSharedKeyAuthenticator QSslPreSharedKeyAuthenticator;
typedef struct QSslServer QSslServer;
typedef struct QSslSocket QSslSocket;
typedef struct QTcpServer QTcpServer;
typedef struct QTcpSocket QTcpSocket;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSslServer* QSslServer_new();
extern __declspec(dllexport) QSslServer* QSslServer_new2(QObject* parent);
extern __declspec(dllexport) void QSslServer_virtbase(QSslServer* src, QTcpServer** outptr_QTcpServer);
extern __declspec(dllexport) QMetaObject* QSslServer_MetaObject(const QSslServer* self);
extern __declspec(dllexport) void* QSslServer_Metacast(QSslServer* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSslServer_Tr(const char* s);
extern __declspec(dllexport) void QSslServer_SetSslConfiguration(QSslServer* self, QSslConfiguration* sslConfiguration);
extern __declspec(dllexport) QSslConfiguration* QSslServer_SslConfiguration(const QSslServer* self);
extern __declspec(dllexport) void QSslServer_SetHandshakeTimeout(QSslServer* self, int timeout);
extern __declspec(dllexport) int QSslServer_HandshakeTimeout(const QSslServer* self);
extern __declspec(dllexport) void QSslServer_SslErrors(QSslServer* self, QSslSocket* socket, struct miqt_array /* of QSslError* */  errors);
void QSslServer_connect_SslErrors(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_PeerVerifyError(QSslServer* self, QSslSocket* socket, QSslError* error);
void QSslServer_connect_PeerVerifyError(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_ErrorOccurred(QSslServer* self, QSslSocket* socket, int error);
void QSslServer_connect_ErrorOccurred(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_PreSharedKeyAuthenticationRequired(QSslServer* self, QSslSocket* socket, QSslPreSharedKeyAuthenticator* authenticator);
void QSslServer_connect_PreSharedKeyAuthenticationRequired(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_AlertSent(QSslServer* self, QSslSocket* socket, int level, int typeVal, struct miqt_string description);
void QSslServer_connect_AlertSent(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_AlertReceived(QSslServer* self, QSslSocket* socket, int level, int typeVal, struct miqt_string description);
void QSslServer_connect_AlertReceived(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_HandshakeInterruptedOnError(QSslServer* self, QSslSocket* socket, QSslError* error);
void QSslServer_connect_HandshakeInterruptedOnError(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_StartedEncryptionHandshake(QSslServer* self, QSslSocket* socket);
void QSslServer_connect_StartedEncryptionHandshake(QSslServer* self, intptr_t slot);
extern __declspec(dllexport) void QSslServer_IncomingConnection(QSslServer* self, intptr_t socket);
extern __declspec(dllexport) struct miqt_string QSslServer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSslServer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSslServer_override_virtual_IncomingConnection(void* self, intptr_t slot);
void QSslServer_virtualbase_IncomingConnection(void* self, intptr_t socket);
extern __declspec(dllexport) void QSslServer_override_virtual_HasPendingConnections(void* self, intptr_t slot);
bool QSslServer_virtualbase_HasPendingConnections(const void* self);
extern __declspec(dllexport) void QSslServer_override_virtual_NextPendingConnection(void* self, intptr_t slot);
QTcpSocket* QSslServer_virtualbase_NextPendingConnection(void* self);
extern __declspec(dllexport) void QSslServer_Delete(QSslServer* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
