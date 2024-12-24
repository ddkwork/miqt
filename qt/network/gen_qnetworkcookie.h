#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKCOOKIE_H
#define MIQT_QT_NETWORK_GEN_QNETWORKCOOKIE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArrayView;
class QDateTime;
class QNetworkCookie;
class QUrl;
class _GUID;
class type_info;
#else
typedef struct QByteArrayView QByteArrayView;
typedef struct QDateTime QDateTime;
typedef struct QNetworkCookie QNetworkCookie;
typedef struct QUrl QUrl;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNetworkCookie* QNetworkCookie_new();
extern __declspec(dllexport) QNetworkCookie* QNetworkCookie_new2(QNetworkCookie* other);
extern __declspec(dllexport) QNetworkCookie* QNetworkCookie_new3(struct miqt_string name);
extern __declspec(dllexport) QNetworkCookie* QNetworkCookie_new4(struct miqt_string name, struct miqt_string value);
extern __declspec(dllexport) void QNetworkCookie_OperatorAssign(QNetworkCookie* self, QNetworkCookie* other);
extern __declspec(dllexport) void QNetworkCookie_Swap(QNetworkCookie* self, QNetworkCookie* other);
extern __declspec(dllexport) bool QNetworkCookie_OperatorEqual(const QNetworkCookie* self, QNetworkCookie* other);
extern __declspec(dllexport) bool QNetworkCookie_OperatorNotEqual(const QNetworkCookie* self, QNetworkCookie* other);
extern __declspec(dllexport) bool QNetworkCookie_IsSecure(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetSecure(QNetworkCookie* self, bool enable);
extern __declspec(dllexport) bool QNetworkCookie_IsHttpOnly(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetHttpOnly(QNetworkCookie* self, bool enable);
extern __declspec(dllexport) SameSite QNetworkCookie_SameSitePolicy(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetSameSitePolicy(QNetworkCookie* self, SameSite sameSite);
extern __declspec(dllexport) bool QNetworkCookie_IsSessionCookie(const QNetworkCookie* self);
extern __declspec(dllexport) QDateTime* QNetworkCookie_ExpirationDate(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetExpirationDate(QNetworkCookie* self, QDateTime* date);
extern __declspec(dllexport) struct miqt_string QNetworkCookie_Domain(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetDomain(QNetworkCookie* self, struct miqt_string domain);
extern __declspec(dllexport) struct miqt_string QNetworkCookie_Path(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetPath(QNetworkCookie* self, struct miqt_string path);
extern __declspec(dllexport) struct miqt_string QNetworkCookie_Name(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetName(QNetworkCookie* self, struct miqt_string cookieName);
extern __declspec(dllexport) struct miqt_string QNetworkCookie_Value(const QNetworkCookie* self);
extern __declspec(dllexport) void QNetworkCookie_SetValue(QNetworkCookie* self, struct miqt_string value);
extern __declspec(dllexport) struct miqt_string QNetworkCookie_ToRawForm(const QNetworkCookie* self);
extern __declspec(dllexport) bool QNetworkCookie_HasSameIdentifier(const QNetworkCookie* self, QNetworkCookie* other);
extern __declspec(dllexport) void QNetworkCookie_Normalize(QNetworkCookie* self, QUrl* url);
extern __declspec(dllexport) struct miqt_array /* of QNetworkCookie* */  QNetworkCookie_ParseCookies(QByteArrayView* cookieString);
extern __declspec(dllexport) struct miqt_string QNetworkCookie_ToRawForm1(const QNetworkCookie* self, RawForm form);
extern __declspec(dllexport) void QNetworkCookie_Delete(QNetworkCookie* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
