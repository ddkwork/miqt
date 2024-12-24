#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLCONFIGURATION_H
#define MIQT_QT_NETWORK_GEN_QSSLCONFIGURATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QSslCertificate QSslCertificate;
typedef struct QSslCipher QSslCipher;
typedef struct QSslConfiguration QSslConfiguration;
typedef struct QSslDiffieHellmanParameters QSslDiffieHellmanParameters;
typedef struct QSslEllipticCurve QSslEllipticCurve;
typedef struct QSslKey QSslKey;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSslConfiguration* QSslConfiguration_new();
extern __declspec(dllexport) 
QSslConfiguration* QSslConfiguration_new2(QSslConfiguration* other);
extern __declspec(dllexport) 
void QSslConfiguration_OperatorAssign(QSslConfiguration* self, QSslConfiguration* other);
extern __declspec(dllexport) 
void QSslConfiguration_Swap(QSslConfiguration* self, QSslConfiguration* other);
extern __declspec(dllexport) 
bool QSslConfiguration_OperatorEqual(const QSslConfiguration* self, QSslConfiguration* other);
extern __declspec(dllexport) 
bool QSslConfiguration_OperatorNotEqual(const QSslConfiguration* self, QSslConfiguration* other);
extern __declspec(dllexport) 
bool QSslConfiguration_IsNull(const QSslConfiguration* self);
extern __declspec(dllexport) 
int QSslConfiguration_Protocol(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetProtocol(QSslConfiguration* self, int protocol);
extern __declspec(dllexport) 
int QSslConfiguration_PeerVerifyMode(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetPeerVerifyMode(QSslConfiguration* self, int mode);
extern __declspec(dllexport) 
int QSslConfiguration_PeerVerifyDepth(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetPeerVerifyDepth(QSslConfiguration* self, int depth);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslConfiguration_LocalCertificateChain(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetLocalCertificateChain(QSslConfiguration* self, struct miqt_array /* of QSslCertificate* */  localChain);
extern __declspec(dllexport) 
QSslCertificate* QSslConfiguration_LocalCertificate(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetLocalCertificate(QSslConfiguration* self, QSslCertificate* certificate);
extern __declspec(dllexport) 
QSslCertificate* QSslConfiguration_PeerCertificate(const QSslConfiguration* self);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslConfiguration_PeerCertificateChain(const QSslConfiguration* self);
extern __declspec(dllexport) 
QSslCipher* QSslConfiguration_SessionCipher(const QSslConfiguration* self);
extern __declspec(dllexport) 
int QSslConfiguration_SessionProtocol(const QSslConfiguration* self);
extern __declspec(dllexport) 
QSslKey* QSslConfiguration_PrivateKey(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetPrivateKey(QSslConfiguration* self, QSslKey* key);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCipher* */  QSslConfiguration_Ciphers(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetCiphers(QSslConfiguration* self, struct miqt_array /* of QSslCipher* */  ciphers);
extern __declspec(dllexport) 
void QSslConfiguration_SetCiphersWithCiphers(QSslConfiguration* self, struct miqt_string ciphers);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCipher* */  QSslConfiguration_SupportedCiphers();
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslConfiguration_CaCertificates(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetCaCertificates(QSslConfiguration* self, struct miqt_array /* of QSslCertificate* */  certificates);
extern __declspec(dllexport) 
bool QSslConfiguration_AddCaCertificates(QSslConfiguration* self, struct miqt_string path);
extern __declspec(dllexport) 
void QSslConfiguration_AddCaCertificate(QSslConfiguration* self, QSslCertificate* certificate);
extern __declspec(dllexport) 
void QSslConfiguration_AddCaCertificatesWithCertificates(QSslConfiguration* self, struct miqt_array /* of QSslCertificate* */  certificates);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslConfiguration_SystemCaCertificates();
extern __declspec(dllexport) 
void QSslConfiguration_SetSslOption(QSslConfiguration* self, int option, bool on);
extern __declspec(dllexport) 
bool QSslConfiguration_TestSslOption(const QSslConfiguration* self, int option);
extern __declspec(dllexport) 
struct miqt_string QSslConfiguration_SessionTicket(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetSessionTicket(QSslConfiguration* self, struct miqt_string sessionTicket);
extern __declspec(dllexport) 
int QSslConfiguration_SessionTicketLifeTimeHint(const QSslConfiguration* self);
extern __declspec(dllexport) 
QSslKey* QSslConfiguration_EphemeralServerKey(const QSslConfiguration* self);
extern __declspec(dllexport) 
struct miqt_array /* of QSslEllipticCurve* */  QSslConfiguration_EllipticCurves(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetEllipticCurves(QSslConfiguration* self, struct miqt_array /* of QSslEllipticCurve* */  curves);
extern __declspec(dllexport) 
struct miqt_array /* of QSslEllipticCurve* */  QSslConfiguration_SupportedEllipticCurves();
extern __declspec(dllexport) 
struct miqt_string QSslConfiguration_PreSharedKeyIdentityHint(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetPreSharedKeyIdentityHint(QSslConfiguration* self, struct miqt_string hint);
extern __declspec(dllexport) 
QSslDiffieHellmanParameters* QSslConfiguration_DiffieHellmanParameters(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetDiffieHellmanParameters(QSslConfiguration* self, QSslDiffieHellmanParameters* dhparams);
extern __declspec(dllexport) 
void QSslConfiguration_SetBackendConfigurationOption(QSslConfiguration* self, struct miqt_string name, QVariant* value);
extern __declspec(dllexport) 
void QSslConfiguration_SetBackendConfiguration(QSslConfiguration* self);
extern __declspec(dllexport) 
QSslConfiguration* QSslConfiguration_DefaultConfiguration();
extern __declspec(dllexport) 
void QSslConfiguration_SetDefaultConfiguration(QSslConfiguration* configuration);
extern __declspec(dllexport) 
bool QSslConfiguration_HandshakeMustInterruptOnError(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetHandshakeMustInterruptOnError(QSslConfiguration* self, bool interrupt);
extern __declspec(dllexport) 
bool QSslConfiguration_MissingCertificateIsFatal(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetMissingCertificateIsFatal(QSslConfiguration* self, bool cannotRecover);
extern __declspec(dllexport) 
void QSslConfiguration_SetOcspStaplingEnabled(QSslConfiguration* self, bool enable);
extern __declspec(dllexport) 
bool QSslConfiguration_OcspStaplingEnabled(const QSslConfiguration* self);
extern __declspec(dllexport) 
void QSslConfiguration_SetAllowedNextProtocols(QSslConfiguration* self, struct miqt_array /* of struct miqt_string */  protocols);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSslConfiguration_AllowedNextProtocols(const QSslConfiguration* self);
extern __declspec(dllexport) 
struct miqt_string QSslConfiguration_NextNegotiatedProtocol(const QSslConfiguration* self);
extern __declspec(dllexport) 
NextProtocolNegotiationStatus QSslConfiguration_NextProtocolNegotiationStatus(const QSslConfiguration* self);
extern __declspec(dllexport) 
bool QSslConfiguration_AddCaCertificates2(QSslConfiguration* self, struct miqt_string path, int format);
extern __declspec(dllexport) 
bool QSslConfiguration_AddCaCertificates3(QSslConfiguration* self, struct miqt_string path, int format, int syntax);
extern __declspec(dllexport) 
void QSslConfiguration_Delete(QSslConfiguration* self, bool isSubclass);

}
