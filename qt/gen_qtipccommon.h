#pragma once
#ifndef MIQT_QT_GEN_QTIPCCOMMON_H
#define MIQT_QT_GEN_QTIPCCOMMON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QNativeIpcKey;
class _GUID;
class type_info;
#else
typedef struct QNativeIpcKey QNativeIpcKey;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNativeIpcKey* QNativeIpcKey_new();
extern __declspec(dllexport) QNativeIpcKey* QNativeIpcKey_new2(Type typeVal);
extern __declspec(dllexport) QNativeIpcKey* QNativeIpcKey_new3(struct miqt_string k);
extern __declspec(dllexport) QNativeIpcKey* QNativeIpcKey_new4(QNativeIpcKey* other);
extern __declspec(dllexport) QNativeIpcKey* QNativeIpcKey_new5(struct miqt_string k, Type typeVal);
extern __declspec(dllexport) Type QNativeIpcKey_LegacyDefaultTypeForOs();
extern __declspec(dllexport) void QNativeIpcKey_OperatorAssign(QNativeIpcKey* self, QNativeIpcKey* other);
extern __declspec(dllexport) void QNativeIpcKey_Swap(QNativeIpcKey* self, QNativeIpcKey* other);
extern __declspec(dllexport) bool QNativeIpcKey_IsEmpty(const QNativeIpcKey* self);
extern __declspec(dllexport) bool QNativeIpcKey_IsValid(const QNativeIpcKey* self);
extern __declspec(dllexport) Type QNativeIpcKey_Type(const QNativeIpcKey* self);
extern __declspec(dllexport) void QNativeIpcKey_SetType(QNativeIpcKey* self, Type typeVal);
extern __declspec(dllexport) struct miqt_string QNativeIpcKey_NativeKey(const QNativeIpcKey* self);
extern __declspec(dllexport) void QNativeIpcKey_SetNativeKey(QNativeIpcKey* self, struct miqt_string newKey);
extern __declspec(dllexport) struct miqt_string QNativeIpcKey_ToString(const QNativeIpcKey* self);
extern __declspec(dllexport) QNativeIpcKey* QNativeIpcKey_FromString(struct miqt_string stringVal);
extern __declspec(dllexport) void QNativeIpcKey_Delete(QNativeIpcKey* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
