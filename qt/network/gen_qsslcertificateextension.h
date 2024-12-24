#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLCERTIFICATEEXTENSION_H
#define MIQT_QT_NETWORK_GEN_QSSLCERTIFICATEEXTENSION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QSslCertificateExtension QSslCertificateExtension;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSslCertificateExtension* QSslCertificateExtension_new();
extern __declspec(dllexport) QSslCertificateExtension* QSslCertificateExtension_new2(QSslCertificateExtension* other);
extern __declspec(dllexport) void QSslCertificateExtension_OperatorAssign(QSslCertificateExtension* self, QSslCertificateExtension* other);
extern __declspec(dllexport) void QSslCertificateExtension_Swap(QSslCertificateExtension* self, QSslCertificateExtension* other);
extern __declspec(dllexport) struct miqt_string QSslCertificateExtension_Oid(const QSslCertificateExtension* self);
extern __declspec(dllexport) struct miqt_string QSslCertificateExtension_Name(const QSslCertificateExtension* self);
extern __declspec(dllexport) QVariant* QSslCertificateExtension_Value(const QSslCertificateExtension* self);
extern __declspec(dllexport) bool QSslCertificateExtension_IsCritical(const QSslCertificateExtension* self);
extern __declspec(dllexport) bool QSslCertificateExtension_IsSupported(const QSslCertificateExtension* self);
extern __declspec(dllexport) void QSslCertificateExtension_Delete(QSslCertificateExtension* self, bool isSubclass);

} 
