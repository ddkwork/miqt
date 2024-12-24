#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLKEY_H
#define MIQT_QT_NETWORK_GEN_QSSLKEY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QSslKey QSslKey;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSslKey* QSslKey_new();
extern __declspec(dllexport) QSslKey* QSslKey_new2(struct miqt_string encoded, int algorithm);
extern __declspec(dllexport) QSslKey* QSslKey_new3(QIODevice* device, int algorithm);
extern __declspec(dllexport) QSslKey* QSslKey_new4(void* handle);
extern __declspec(dllexport) QSslKey* QSslKey_new5(QSslKey* other);
extern __declspec(dllexport) QSslKey* QSslKey_new6(struct miqt_string encoded, int algorithm, int format);
extern __declspec(dllexport) QSslKey* QSslKey_new7(struct miqt_string encoded, int algorithm, int format, int typeVal);
extern __declspec(dllexport) QSslKey* QSslKey_new8(struct miqt_string encoded, int algorithm, int format, int typeVal, struct miqt_string passPhrase);
extern __declspec(dllexport) QSslKey* QSslKey_new9(QIODevice* device, int algorithm, int format);
extern __declspec(dllexport) QSslKey* QSslKey_new10(QIODevice* device, int algorithm, int format, int typeVal);
extern __declspec(dllexport) QSslKey* QSslKey_new11(QIODevice* device, int algorithm, int format, int typeVal, struct miqt_string passPhrase);
extern __declspec(dllexport) QSslKey* QSslKey_new12(void* handle, int typeVal);
extern __declspec(dllexport) void QSslKey_OperatorAssign(QSslKey* self, QSslKey* other);
extern __declspec(dllexport) void QSslKey_Swap(QSslKey* self, QSslKey* other);
extern __declspec(dllexport) bool QSslKey_IsNull(const QSslKey* self);
extern __declspec(dllexport) void QSslKey_Clear(QSslKey* self);
extern __declspec(dllexport) int QSslKey_Length(const QSslKey* self);
extern __declspec(dllexport) int QSslKey_Type(const QSslKey* self);
extern __declspec(dllexport) int QSslKey_Algorithm(const QSslKey* self);
extern __declspec(dllexport) struct miqt_string QSslKey_ToPem(const QSslKey* self);
extern __declspec(dllexport) struct miqt_string QSslKey_ToDer(const QSslKey* self);
extern __declspec(dllexport) void* QSslKey_Handle(const QSslKey* self);
extern __declspec(dllexport) bool QSslKey_OperatorEqual(const QSslKey* self, QSslKey* key);
extern __declspec(dllexport) bool QSslKey_OperatorNotEqual(const QSslKey* self, QSslKey* key);
extern __declspec(dllexport) struct miqt_string QSslKey_ToPem1(const QSslKey* self, struct miqt_string passPhrase);
extern __declspec(dllexport) struct miqt_string QSslKey_ToDer1(const QSslKey* self, struct miqt_string passPhrase);
extern __declspec(dllexport) void QSslKey_Delete(QSslKey* self, bool isSubclass);

} 
