#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLPRESHAREDKEYAUTHENTICATOR_H
#define MIQT_QT_NETWORK_GEN_QSSLPRESHAREDKEYAUTHENTICATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QSslPreSharedKeyAuthenticator QSslPreSharedKeyAuthenticator;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSslPreSharedKeyAuthenticator* QSslPreSharedKeyAuthenticator_new();
extern __declspec(dllexport) 
QSslPreSharedKeyAuthenticator* QSslPreSharedKeyAuthenticator_new2(QSslPreSharedKeyAuthenticator* authenticator);
extern __declspec(dllexport) 
void QSslPreSharedKeyAuthenticator_OperatorAssign(QSslPreSharedKeyAuthenticator* self, QSslPreSharedKeyAuthenticator* authenticator);
extern __declspec(dllexport) 
void QSslPreSharedKeyAuthenticator_Swap(QSslPreSharedKeyAuthenticator* self, QSslPreSharedKeyAuthenticator* other);
extern __declspec(dllexport) 
struct miqt_string QSslPreSharedKeyAuthenticator_IdentityHint(const QSslPreSharedKeyAuthenticator* self);
extern __declspec(dllexport) 
void QSslPreSharedKeyAuthenticator_SetIdentity(QSslPreSharedKeyAuthenticator* self, struct miqt_string identity);
extern __declspec(dllexport) 
struct miqt_string QSslPreSharedKeyAuthenticator_Identity(const QSslPreSharedKeyAuthenticator* self);
extern __declspec(dllexport) 
int QSslPreSharedKeyAuthenticator_MaximumIdentityLength(const QSslPreSharedKeyAuthenticator* self);
extern __declspec(dllexport) 
void QSslPreSharedKeyAuthenticator_SetPreSharedKey(QSslPreSharedKeyAuthenticator* self, struct miqt_string preSharedKey);
extern __declspec(dllexport) 
struct miqt_string QSslPreSharedKeyAuthenticator_PreSharedKey(const QSslPreSharedKeyAuthenticator* self);
extern __declspec(dllexport) 
int QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(const QSslPreSharedKeyAuthenticator* self);
extern __declspec(dllexport) 
void QSslPreSharedKeyAuthenticator_Delete(QSslPreSharedKeyAuthenticator* self, bool isSubclass);

}
