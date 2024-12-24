#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKINFORMATION_H
#define MIQT_QT_NETWORK_GEN_QNETWORKINFORMATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMetaObject;
class QNetworkInformation;
class QObject;
class _GUID;
class type_info;
#else
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkInformation QNetworkInformation;
typedef struct QObject QObject;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QNetworkInformation_virtbase(QNetworkInformation* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QNetworkInformation_MetaObject(const QNetworkInformation* self);
extern __declspec(dllexport) void* QNetworkInformation_Metacast(QNetworkInformation* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QNetworkInformation_Tr(const char* s);
extern __declspec(dllexport) Reachability QNetworkInformation_Reachability(const QNetworkInformation* self);
extern __declspec(dllexport) bool QNetworkInformation_IsBehindCaptivePortal(const QNetworkInformation* self);
extern __declspec(dllexport) TransportMedium QNetworkInformation_TransportMedium(const QNetworkInformation* self);
extern __declspec(dllexport) bool QNetworkInformation_IsMetered(const QNetworkInformation* self);
extern __declspec(dllexport) struct miqt_string QNetworkInformation_BackendName(const QNetworkInformation* self);
extern __declspec(dllexport) bool QNetworkInformation_Supports(const QNetworkInformation* self, Features features);
extern __declspec(dllexport) Features QNetworkInformation_SupportedFeatures(const QNetworkInformation* self);
extern __declspec(dllexport) bool QNetworkInformation_LoadDefaultBackend();
extern __declspec(dllexport) bool QNetworkInformation_LoadBackendByFeatures(Features features);
extern __declspec(dllexport) bool QNetworkInformation_LoadWithFeatures(Features features);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QNetworkInformation_AvailableBackends();
extern __declspec(dllexport) QNetworkInformation* QNetworkInformation_Instance();
extern __declspec(dllexport) void QNetworkInformation_ReachabilityChanged(QNetworkInformation* self, int newReachability);
void QNetworkInformation_connect_ReachabilityChanged(QNetworkInformation* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkInformation_IsBehindCaptivePortalChanged(QNetworkInformation* self, bool state);
void QNetworkInformation_connect_IsBehindCaptivePortalChanged(QNetworkInformation* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkInformation_TransportMediumChanged(QNetworkInformation* self, int current);
void QNetworkInformation_connect_TransportMediumChanged(QNetworkInformation* self, intptr_t slot);
extern __declspec(dllexport) void QNetworkInformation_IsMeteredChanged(QNetworkInformation* self, bool isMetered);
void QNetworkInformation_connect_IsMeteredChanged(QNetworkInformation* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QNetworkInformation_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QNetworkInformation_Tr3(const char* s, const char* c, int n);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
