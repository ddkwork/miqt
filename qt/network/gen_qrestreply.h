#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QRESTREPLY_H
#define MIQT_QT_NETWORK_GEN_QRESTREPLY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QNetworkReply;
class QRestReply;
class _GUID;
class type_info;
#else
typedef struct QNetworkReply QNetworkReply;
typedef struct QRestReply QRestReply;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRestReply* QRestReply_new(QNetworkReply* reply);
extern __declspec(dllexport) void QRestReply_Swap(QRestReply* self, QRestReply* other);
extern __declspec(dllexport) QNetworkReply* QRestReply_NetworkReply(const QRestReply* self);
extern __declspec(dllexport) struct miqt_string QRestReply_ReadBody(QRestReply* self);
extern __declspec(dllexport) struct miqt_string QRestReply_ReadText(QRestReply* self);
extern __declspec(dllexport) bool QRestReply_IsSuccess(const QRestReply* self);
extern __declspec(dllexport) int QRestReply_HttpStatus(const QRestReply* self);
extern __declspec(dllexport) bool QRestReply_IsHttpStatusSuccess(const QRestReply* self);
extern __declspec(dllexport) bool QRestReply_HasError(const QRestReply* self);
extern __declspec(dllexport) int QRestReply_Error(const QRestReply* self);
extern __declspec(dllexport) struct miqt_string QRestReply_ErrorString(const QRestReply* self);
extern __declspec(dllexport) void QRestReply_Delete(QRestReply* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
