#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QAUTHENTICATOR_H
#define MIQT_QT_NETWORK_GEN_QAUTHENTICATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAuthenticator;
class QVariant;
class _GUID;
class type_info;
#else
typedef struct QAuthenticator QAuthenticator;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAuthenticator* QAuthenticator_new();
extern __declspec(dllexport) QAuthenticator* QAuthenticator_new2(QAuthenticator* other);
extern __declspec(dllexport) void QAuthenticator_OperatorAssign(QAuthenticator* self, QAuthenticator* other);
extern __declspec(dllexport) bool QAuthenticator_OperatorEqual(const QAuthenticator* self, QAuthenticator* other);
extern __declspec(dllexport) bool QAuthenticator_OperatorNotEqual(const QAuthenticator* self, QAuthenticator* other);
extern __declspec(dllexport) struct miqt_string QAuthenticator_User(const QAuthenticator* self);
extern __declspec(dllexport) void QAuthenticator_SetUser(QAuthenticator* self, struct miqt_string user);
extern __declspec(dllexport) struct miqt_string QAuthenticator_Password(const QAuthenticator* self);
extern __declspec(dllexport) void QAuthenticator_SetPassword(QAuthenticator* self, struct miqt_string password);
extern __declspec(dllexport) struct miqt_string QAuthenticator_Realm(const QAuthenticator* self);
extern __declspec(dllexport) void QAuthenticator_SetRealm(QAuthenticator* self, struct miqt_string realm);
extern __declspec(dllexport) QVariant* QAuthenticator_Option(const QAuthenticator* self, struct miqt_string opt);
extern __declspec(dllexport) struct miqt_map /* of struct miqt_string to QVariant* */  QAuthenticator_Options(const QAuthenticator* self);
extern __declspec(dllexport) void QAuthenticator_SetOption(QAuthenticator* self, struct miqt_string opt, QVariant* value);
extern __declspec(dllexport) bool QAuthenticator_IsNull(const QAuthenticator* self);
extern __declspec(dllexport) void QAuthenticator_Detach(QAuthenticator* self);
extern __declspec(dllexport) void QAuthenticator_Delete(QAuthenticator* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
