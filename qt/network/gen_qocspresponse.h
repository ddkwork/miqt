#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QOCSPRESPONSE_H
#define MIQT_QT_NETWORK_GEN_QOCSPRESPONSE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QOcspResponse;
class QSslCertificate;
class _GUID;
class type_info;
#else
typedef struct QOcspResponse QOcspResponse;
typedef struct QSslCertificate QSslCertificate;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QOcspResponse* QOcspResponse_new();
extern __declspec(dllexport) QOcspResponse* QOcspResponse_new2(QOcspResponse* other);
extern __declspec(dllexport) void QOcspResponse_OperatorAssign(QOcspResponse* self, QOcspResponse* other);
extern __declspec(dllexport) int QOcspResponse_CertificateStatus(const QOcspResponse* self);
extern __declspec(dllexport) int QOcspResponse_RevocationReason(const QOcspResponse* self);
extern __declspec(dllexport) QSslCertificate* QOcspResponse_Responder(const QOcspResponse* self);
extern __declspec(dllexport) QSslCertificate* QOcspResponse_Subject(const QOcspResponse* self);
extern __declspec(dllexport) void QOcspResponse_Swap(QOcspResponse* self, QOcspResponse* other);
extern __declspec(dllexport) void QOcspResponse_Delete(QOcspResponse* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
