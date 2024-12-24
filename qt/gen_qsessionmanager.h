#pragma once
#ifndef MIQT_QT_GEN_QSESSIONMANAGER_H
#define MIQT_QT_GEN_QSESSIONMANAGER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMetaObject;
class QObject;
class QSessionManager;
class _GUID;
class type_info;
#else
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSessionManager QSessionManager;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QSessionManager_virtbase(QSessionManager* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSessionManager_MetaObject(const QSessionManager* self);
extern __declspec(dllexport) void* QSessionManager_Metacast(QSessionManager* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSessionManager_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QSessionManager_SessionId(const QSessionManager* self);
extern __declspec(dllexport) struct miqt_string QSessionManager_SessionKey(const QSessionManager* self);
extern __declspec(dllexport) bool QSessionManager_AllowsInteraction(QSessionManager* self);
extern __declspec(dllexport) bool QSessionManager_AllowsErrorInteraction(QSessionManager* self);
extern __declspec(dllexport) void QSessionManager_Release(QSessionManager* self);
extern __declspec(dllexport) void QSessionManager_Cancel(QSessionManager* self);
extern __declspec(dllexport) void QSessionManager_SetRestartHint(QSessionManager* self, RestartHint restartHint);
extern __declspec(dllexport) RestartHint QSessionManager_RestartHint(const QSessionManager* self);
extern __declspec(dllexport) void QSessionManager_SetRestartCommand(QSessionManager* self, struct miqt_array /* of struct miqt_string */  restartCommand);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QSessionManager_RestartCommand(const QSessionManager* self);
extern __declspec(dllexport) void QSessionManager_SetDiscardCommand(QSessionManager* self, struct miqt_array /* of struct miqt_string */  discardCommand);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QSessionManager_DiscardCommand(const QSessionManager* self);
extern __declspec(dllexport) void QSessionManager_SetManagerProperty(QSessionManager* self, struct miqt_string name, struct miqt_string value);
extern __declspec(dllexport) void QSessionManager_SetManagerProperty2(QSessionManager* self, struct miqt_string name, struct miqt_array /* of struct miqt_string */  value);
extern __declspec(dllexport) bool QSessionManager_IsPhase2(const QSessionManager* self);
extern __declspec(dllexport) void QSessionManager_RequestPhase2(QSessionManager* self);
extern __declspec(dllexport) struct miqt_string QSessionManager_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSessionManager_Tr3(const char* s, const char* c, int n);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
