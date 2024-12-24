#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QHSTSPOLICY_H
#define MIQT_QT_NETWORK_GEN_QHSTSPOLICY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDateTime QDateTime;
typedef struct QHstsPolicy QHstsPolicy;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QHstsPolicy* QHstsPolicy_new();
extern __declspec(dllexport) 
QHstsPolicy* QHstsPolicy_new2(QDateTime* expiry, PolicyFlags flags, struct miqt_string host);
extern __declspec(dllexport) 
QHstsPolicy* QHstsPolicy_new3(QHstsPolicy* rhs);
extern __declspec(dllexport) 
QHstsPolicy* QHstsPolicy_new4(QDateTime* expiry, PolicyFlags flags, struct miqt_string host, int mode);
extern __declspec(dllexport) 
void QHstsPolicy_OperatorAssign(QHstsPolicy* self, QHstsPolicy* rhs);
extern __declspec(dllexport) 
void QHstsPolicy_Swap(QHstsPolicy* self, QHstsPolicy* other);
extern __declspec(dllexport) 
void QHstsPolicy_SetHost(QHstsPolicy* self, struct miqt_string host);
extern __declspec(dllexport) 
struct miqt_string QHstsPolicy_Host(const QHstsPolicy* self);
extern __declspec(dllexport) 
void QHstsPolicy_SetExpiry(QHstsPolicy* self, QDateTime* expiry);
extern __declspec(dllexport) 
QDateTime* QHstsPolicy_Expiry(const QHstsPolicy* self);
extern __declspec(dllexport) 
void QHstsPolicy_SetIncludesSubDomains(QHstsPolicy* self, bool include);
extern __declspec(dllexport) 
bool QHstsPolicy_IncludesSubDomains(const QHstsPolicy* self);
extern __declspec(dllexport) 
bool QHstsPolicy_IsExpired(const QHstsPolicy* self);
extern __declspec(dllexport) 
void QHstsPolicy_SetHost2(QHstsPolicy* self, struct miqt_string host, int mode);
extern __declspec(dllexport) 
struct miqt_string QHstsPolicy_Host1(const QHstsPolicy* self, int options);
extern __declspec(dllexport) 
void QHstsPolicy_Delete(QHstsPolicy* self, bool isSubclass);

}
