#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLCIPHER_H
#define MIQT_QT_NETWORK_GEN_QSSLCIPHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QSslCipher QSslCipher;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSslCipher* QSslCipher_new();
extern __declspec(dllexport) 
QSslCipher* QSslCipher_new2(struct miqt_string name);
extern __declspec(dllexport) 
QSslCipher* QSslCipher_new3(struct miqt_string name, int protocol);
extern __declspec(dllexport) 
QSslCipher* QSslCipher_new4(QSslCipher* other);
extern __declspec(dllexport) 
void QSslCipher_OperatorAssign(QSslCipher* self, QSslCipher* other);
extern __declspec(dllexport) 
void QSslCipher_Swap(QSslCipher* self, QSslCipher* other);
extern __declspec(dllexport) 
bool QSslCipher_OperatorEqual(const QSslCipher* self, QSslCipher* other);
extern __declspec(dllexport) 
bool QSslCipher_OperatorNotEqual(const QSslCipher* self, QSslCipher* other);
extern __declspec(dllexport) 
bool QSslCipher_IsNull(const QSslCipher* self);
extern __declspec(dllexport) 
struct miqt_string QSslCipher_Name(const QSslCipher* self);
extern __declspec(dllexport) 
int QSslCipher_SupportedBits(const QSslCipher* self);
extern __declspec(dllexport) 
int QSslCipher_UsedBits(const QSslCipher* self);
extern __declspec(dllexport) 
struct miqt_string QSslCipher_KeyExchangeMethod(const QSslCipher* self);
extern __declspec(dllexport) 
struct miqt_string QSslCipher_AuthenticationMethod(const QSslCipher* self);
extern __declspec(dllexport) 
struct miqt_string QSslCipher_EncryptionMethod(const QSslCipher* self);
extern __declspec(dllexport) 
struct miqt_string QSslCipher_ProtocolString(const QSslCipher* self);
extern __declspec(dllexport) 
int QSslCipher_Protocol(const QSslCipher* self);
extern __declspec(dllexport) 
void QSslCipher_Delete(QSslCipher* self, bool isSubclass);

}
