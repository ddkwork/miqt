#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QHOSTINFO_H
#define MIQT_QT_NETWORK_GEN_QHOSTINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHostAddress QHostAddress;
typedef struct QHostInfo QHostInfo;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QHostInfo* QHostInfo_new();
extern __declspec(dllexport) QHostInfo* QHostInfo_new2(QHostInfo* d);
extern __declspec(dllexport) QHostInfo* QHostInfo_new3(int lookupId);
extern __declspec(dllexport) void QHostInfo_OperatorAssign(QHostInfo* self, QHostInfo* d);
extern __declspec(dllexport) void QHostInfo_Swap(QHostInfo* self, QHostInfo* other);
extern __declspec(dllexport) struct miqt_string QHostInfo_HostName(const QHostInfo* self);
extern __declspec(dllexport) void QHostInfo_SetHostName(QHostInfo* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_array /* of QHostAddress* */  QHostInfo_Addresses(const QHostInfo* self);
extern __declspec(dllexport) void QHostInfo_SetAddresses(QHostInfo* self, struct miqt_array /* of QHostAddress* */  addresses);
extern __declspec(dllexport) HostInfoError QHostInfo_Error(const QHostInfo* self);
extern __declspec(dllexport) void QHostInfo_SetError(QHostInfo* self, HostInfoError error);
extern __declspec(dllexport) struct miqt_string QHostInfo_ErrorString(const QHostInfo* self);
extern __declspec(dllexport) void QHostInfo_SetErrorString(QHostInfo* self, struct miqt_string errorString);
extern __declspec(dllexport) void QHostInfo_SetLookupId(QHostInfo* self, int id);
extern __declspec(dllexport) int QHostInfo_LookupId(const QHostInfo* self);
extern __declspec(dllexport) void QHostInfo_AbortHostLookup(int lookupId);
extern __declspec(dllexport) QHostInfo* QHostInfo_FromName(struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QHostInfo_LocalHostName();
extern __declspec(dllexport) struct miqt_string QHostInfo_LocalDomainName();
extern __declspec(dllexport) void QHostInfo_Delete(QHostInfo* self, bool isSubclass);

} 
