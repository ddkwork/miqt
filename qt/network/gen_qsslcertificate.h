#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLCERTIFICATE_H
#define MIQT_QT_NETWORK_GEN_QSSLCERTIFICATE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDateTime QDateTime;
typedef struct QIODevice QIODevice;
typedef struct QSslCertificate QSslCertificate;
typedef struct QSslCertificateExtension QSslCertificateExtension;
typedef struct QSslError QSslError;
typedef struct QSslKey QSslKey;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSslCertificate* QSslCertificate_new(QIODevice* device);
extern __declspec(dllexport) 
QSslCertificate* QSslCertificate_new2();
extern __declspec(dllexport) 
QSslCertificate* QSslCertificate_new3(QSslCertificate* other);
extern __declspec(dllexport) 
QSslCertificate* QSslCertificate_new4(QIODevice* device, int format);
extern __declspec(dllexport) 
QSslCertificate* QSslCertificate_new5(struct miqt_string data);
extern __declspec(dllexport) 
QSslCertificate* QSslCertificate_new6(struct miqt_string data, int format);
extern __declspec(dllexport) 
void QSslCertificate_OperatorAssign(QSslCertificate* self, QSslCertificate* other);
extern __declspec(dllexport) 
void QSslCertificate_Swap(QSslCertificate* self, QSslCertificate* other);
extern __declspec(dllexport) 
bool QSslCertificate_OperatorEqual(const QSslCertificate* self, QSslCertificate* other);
extern __declspec(dllexport) 
bool QSslCertificate_OperatorNotEqual(const QSslCertificate* self, QSslCertificate* other);
extern __declspec(dllexport) 
bool QSslCertificate_IsNull(const QSslCertificate* self);
extern __declspec(dllexport) 
bool QSslCertificate_IsBlacklisted(const QSslCertificate* self);
extern __declspec(dllexport) 
bool QSslCertificate_IsSelfSigned(const QSslCertificate* self);
extern __declspec(dllexport) 
void QSslCertificate_Clear(QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_Version(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_SerialNumber(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_Digest(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSslCertificate_IssuerInfo(const QSslCertificate* self, SubjectInfo info);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSslCertificate_IssuerInfoWithAttribute(const QSslCertificate* self, struct miqt_string attribute);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSslCertificate_SubjectInfo(const QSslCertificate* self, SubjectInfo info);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSslCertificate_SubjectInfoWithAttribute(const QSslCertificate* self, struct miqt_string attribute);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_IssuerDisplayName(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_SubjectDisplayName(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSslCertificate_SubjectInfoAttributes(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSslCertificate_IssuerInfoAttributes(const QSslCertificate* self);
extern __declspec(dllexport) 
QDateTime* QSslCertificate_EffectiveDate(const QSslCertificate* self);
extern __declspec(dllexport) 
QDateTime* QSslCertificate_ExpiryDate(const QSslCertificate* self);
extern __declspec(dllexport) 
QSslKey* QSslCertificate_PublicKey(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificateExtension* */  QSslCertificate_Extensions(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_ToPem(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_ToDer(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_ToText(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslCertificate_FromPath(struct miqt_string path);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslCertificate_FromDevice(QIODevice* device);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslCertificate_FromData(struct miqt_string data);
extern __declspec(dllexport) 
struct miqt_array /* of QSslError* */  QSslCertificate_Verify(struct miqt_array /* of QSslCertificate* */  certificateChain);
extern __declspec(dllexport) 
bool QSslCertificate_ImportPkcs12(QIODevice* device, QSslKey* key, QSslCertificate* cert);
extern __declspec(dllexport) 
void* QSslCertificate_Handle(const QSslCertificate* self);
extern __declspec(dllexport) 
struct miqt_string QSslCertificate_Digest1(const QSslCertificate* self, int algorithm);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslCertificate_FromPath2(struct miqt_string path, int format);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslCertificate_FromPath3(struct miqt_string path, int format, PatternSyntax syntax);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslCertificate_FromDevice2(QIODevice* device, int format);
extern __declspec(dllexport) 
struct miqt_array /* of QSslCertificate* */  QSslCertificate_FromData2(struct miqt_string data, int format);
extern __declspec(dllexport) 
struct miqt_array /* of QSslError* */  QSslCertificate_Verify2(struct miqt_array /* of QSslCertificate* */  certificateChain, struct miqt_string hostName);
extern __declspec(dllexport) 
bool QSslCertificate_ImportPkcs124(QIODevice* device, QSslKey* key, QSslCertificate* cert, struct miqt_array /* of QSslCertificate* */  caCertificates);
extern __declspec(dllexport) 
bool QSslCertificate_ImportPkcs125(QIODevice* device, QSslKey* key, QSslCertificate* cert, struct miqt_array /* of QSslCertificate* */  caCertificates, struct miqt_string passPhrase);
extern __declspec(dllexport) 
void QSslCertificate_Delete(QSslCertificate* self, bool isSubclass);

}
