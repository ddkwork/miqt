#pragma once
#ifndef MIQT_QT_GEN_QSYSTEMSEMAPHORE_H
#define MIQT_QT_GEN_QSYSTEMSEMAPHORE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QNativeIpcKey QNativeIpcKey;
typedef struct QSystemSemaphore QSystemSemaphore;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSystemSemaphore* QSystemSemaphore_new(QNativeIpcKey* key);
extern __declspec(dllexport) QSystemSemaphore* QSystemSemaphore_new2(struct miqt_string key);
extern __declspec(dllexport) QSystemSemaphore* QSystemSemaphore_new3(QNativeIpcKey* key, int initialValue);
extern __declspec(dllexport) QSystemSemaphore* QSystemSemaphore_new4(QNativeIpcKey* key, int initialValue, AccessMode param3);
extern __declspec(dllexport) QSystemSemaphore* QSystemSemaphore_new5(struct miqt_string key, int initialValue);
extern __declspec(dllexport) QSystemSemaphore* QSystemSemaphore_new6(struct miqt_string key, int initialValue, AccessMode mode);
extern __declspec(dllexport) struct miqt_string QSystemSemaphore_Tr(const char* sourceText);
extern __declspec(dllexport) void QSystemSemaphore_SetNativeKey(QSystemSemaphore* self, QNativeIpcKey* key);
extern __declspec(dllexport) void QSystemSemaphore_SetNativeKeyWithKey(QSystemSemaphore* self, struct miqt_string key);
extern __declspec(dllexport) QNativeIpcKey* QSystemSemaphore_NativeIpcKey(const QSystemSemaphore* self);
extern __declspec(dllexport) void QSystemSemaphore_SetKey(QSystemSemaphore* self, struct miqt_string key);
extern __declspec(dllexport) struct miqt_string QSystemSemaphore_Key(const QSystemSemaphore* self);
extern __declspec(dllexport) bool QSystemSemaphore_Acquire(QSystemSemaphore* self);
extern __declspec(dllexport) bool QSystemSemaphore_Release(QSystemSemaphore* self);
extern __declspec(dllexport) SystemSemaphoreError QSystemSemaphore_Error(const QSystemSemaphore* self);
extern __declspec(dllexport) struct miqt_string QSystemSemaphore_ErrorString(const QSystemSemaphore* self);
extern __declspec(dllexport) bool QSystemSemaphore_IsKeyTypeSupported(uint16_t typeVal);
extern __declspec(dllexport) QNativeIpcKey* QSystemSemaphore_PlatformSafeKey(struct miqt_string key);
extern __declspec(dllexport) QNativeIpcKey* QSystemSemaphore_LegacyNativeKey(struct miqt_string key);
extern __declspec(dllexport) struct miqt_string QSystemSemaphore_Tr2(const char* sourceText, const char* disambiguation);
extern __declspec(dllexport) struct miqt_string QSystemSemaphore_Tr3(const char* sourceText, const char* disambiguation, int n);
extern __declspec(dllexport) void QSystemSemaphore_SetNativeKey2(QSystemSemaphore* self, QNativeIpcKey* key, int initialValue);
extern __declspec(dllexport) void QSystemSemaphore_SetNativeKey3(QSystemSemaphore* self, QNativeIpcKey* key, int initialValue, AccessMode param3);
extern __declspec(dllexport) void QSystemSemaphore_SetNativeKey22(QSystemSemaphore* self, struct miqt_string key, int initialValue);
extern __declspec(dllexport) void QSystemSemaphore_SetNativeKey32(QSystemSemaphore* self, struct miqt_string key, int initialValue, AccessMode mode);
extern __declspec(dllexport) void QSystemSemaphore_SetNativeKey4(QSystemSemaphore* self, struct miqt_string key, int initialValue, AccessMode mode, uint16_t typeVal);
extern __declspec(dllexport) void QSystemSemaphore_SetKey2(QSystemSemaphore* self, struct miqt_string key, int initialValue);
extern __declspec(dllexport) void QSystemSemaphore_SetKey3(QSystemSemaphore* self, struct miqt_string key, int initialValue, AccessMode mode);
extern __declspec(dllexport) bool QSystemSemaphore_Release1(QSystemSemaphore* self, int n);
extern __declspec(dllexport) QNativeIpcKey* QSystemSemaphore_PlatformSafeKey2(struct miqt_string key, uint16_t typeVal);
extern __declspec(dllexport) QNativeIpcKey* QSystemSemaphore_LegacyNativeKey2(struct miqt_string key, uint16_t typeVal);
extern __declspec(dllexport) void QSystemSemaphore_Delete(QSystemSemaphore* self, bool isSubclass);

} 
