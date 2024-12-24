#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLELLIPTICCURVE_H
#define MIQT_QT_NETWORK_GEN_QSSLELLIPTICCURVE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QSslEllipticCurve QSslEllipticCurve;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSslEllipticCurve* QSslEllipticCurve_new();
extern __declspec(dllexport) QSslEllipticCurve* QSslEllipticCurve_new2(QSslEllipticCurve* param1);
extern __declspec(dllexport) QSslEllipticCurve* QSslEllipticCurve_FromShortName(struct miqt_string name);
extern __declspec(dllexport) QSslEllipticCurve* QSslEllipticCurve_FromLongName(struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QSslEllipticCurve_ShortName(const QSslEllipticCurve* self);
extern __declspec(dllexport) struct miqt_string QSslEllipticCurve_LongName(const QSslEllipticCurve* self);
extern __declspec(dllexport) bool QSslEllipticCurve_IsValid(const QSslEllipticCurve* self);
extern __declspec(dllexport) bool QSslEllipticCurve_IsTlsNamedCurve(const QSslEllipticCurve* self);
extern __declspec(dllexport) void QSslEllipticCurve_Delete(QSslEllipticCurve* self, bool isSubclass);

} 
