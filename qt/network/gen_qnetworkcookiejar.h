#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKCOOKIEJAR_H
#define MIQT_QT_NETWORK_GEN_QNETWORKCOOKIEJAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QNetworkCookie;
class QNetworkCookieJar;
class QObject;
class QTimerEvent;
class QUrl;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkCookie QNetworkCookie;
typedef struct QNetworkCookieJar QNetworkCookieJar;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUrl QUrl;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNetworkCookieJar* QNetworkCookieJar_new();
extern __declspec(dllexport) QNetworkCookieJar* QNetworkCookieJar_new2(QObject* parent);
extern __declspec(dllexport) void QNetworkCookieJar_virtbase(QNetworkCookieJar* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QNetworkCookieJar_MetaObject(const QNetworkCookieJar* self);
extern __declspec(dllexport) void* QNetworkCookieJar_Metacast(QNetworkCookieJar* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QNetworkCookieJar_Tr(const char* s);
extern __declspec(dllexport) struct miqt_array /* of QNetworkCookie* */  QNetworkCookieJar_CookiesForUrl(const QNetworkCookieJar* self, QUrl* url);
extern __declspec(dllexport) bool QNetworkCookieJar_SetCookiesFromUrl(QNetworkCookieJar* self, struct miqt_array /* of QNetworkCookie* */  cookieList, QUrl* url);
extern __declspec(dllexport) bool QNetworkCookieJar_InsertCookie(QNetworkCookieJar* self, QNetworkCookie* cookie);
extern __declspec(dllexport) bool QNetworkCookieJar_UpdateCookie(QNetworkCookieJar* self, QNetworkCookie* cookie);
extern __declspec(dllexport) bool QNetworkCookieJar_DeleteCookie(QNetworkCookieJar* self, QNetworkCookie* cookie);
extern __declspec(dllexport) bool QNetworkCookieJar_ValidateCookie(const QNetworkCookieJar* self, QNetworkCookie* cookie, QUrl* url);
extern __declspec(dllexport) struct miqt_string QNetworkCookieJar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QNetworkCookieJar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_CookiesForUrl(void* self, intptr_t slot);
struct miqt_array /* of QNetworkCookie* */  QNetworkCookieJar_virtualbase_CookiesForUrl(const void* self, QUrl* url);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_SetCookiesFromUrl(void* self, intptr_t slot);
bool QNetworkCookieJar_virtualbase_SetCookiesFromUrl(void* self, struct miqt_array /* of QNetworkCookie* */  cookieList, QUrl* url);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_InsertCookie(void* self, intptr_t slot);
bool QNetworkCookieJar_virtualbase_InsertCookie(void* self, QNetworkCookie* cookie);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_UpdateCookie(void* self, intptr_t slot);
bool QNetworkCookieJar_virtualbase_UpdateCookie(void* self, QNetworkCookie* cookie);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_DeleteCookie(void* self, intptr_t slot);
bool QNetworkCookieJar_virtualbase_DeleteCookie(void* self, QNetworkCookie* cookie);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_ValidateCookie(void* self, intptr_t slot);
bool QNetworkCookieJar_virtualbase_ValidateCookie(const void* self, QNetworkCookie* cookie, QUrl* url);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_Event(void* self, intptr_t slot);
bool QNetworkCookieJar_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_EventFilter(void* self, intptr_t slot);
bool QNetworkCookieJar_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_TimerEvent(void* self, intptr_t slot);
void QNetworkCookieJar_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_ChildEvent(void* self, intptr_t slot);
void QNetworkCookieJar_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_CustomEvent(void* self, intptr_t slot);
void QNetworkCookieJar_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QNetworkCookieJar_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QNetworkCookieJar_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QNetworkCookieJar_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QNetworkCookieJar_Delete(QNetworkCookieJar* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
