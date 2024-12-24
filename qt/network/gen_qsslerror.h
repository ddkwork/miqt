#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLERROR_H
#define MIQT_QT_NETWORK_GEN_QSSLERROR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QSslCertificate;
class QSslError;
class _GUID;
class type_info;
#else
typedef struct QSslCertificate QSslCertificate;
typedef struct QSslError QSslError;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSslError* QSslError_new();
extern __declspec(dllexport) QSslError* QSslError_new2(SslError error);
extern __declspec(dllexport) QSslError* QSslError_new3(SslError error, QSslCertificate* certificate);
extern __declspec(dllexport) QSslError* QSslError_new4(QSslError* other);
extern __declspec(dllexport) void QSslError_Swap(QSslError* self, QSslError* other);
extern __declspec(dllexport) void QSslError_OperatorAssign(QSslError* self, QSslError* other);
extern __declspec(dllexport) bool QSslError_OperatorEqual(const QSslError* self, QSslError* other);
extern __declspec(dllexport) bool QSslError_OperatorNotEqual(const QSslError* self, QSslError* other);
extern __declspec(dllexport) SslError QSslError_Error(const QSslError* self);
extern __declspec(dllexport) struct miqt_string QSslError_ErrorString(const QSslError* self);
extern __declspec(dllexport) QSslCertificate* QSslError_Certificate(const QSslError* self);
extern __declspec(dllexport) void QSslError_Delete(QSslError* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
