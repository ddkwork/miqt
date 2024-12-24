#pragma once
#ifndef MIQT_QT_GEN_QSHAREDMEMORY_H
#define MIQT_QT_GEN_QSHAREDMEMORY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QNativeIpcKey;
class QObject;
class QSharedMemory;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QNativeIpcKey QNativeIpcKey;
typedef struct QObject QObject;
typedef struct QSharedMemory QSharedMemory;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSharedMemory* QSharedMemory_new();
extern __declspec(dllexport) QSharedMemory* QSharedMemory_new2(QNativeIpcKey* key);
extern __declspec(dllexport) QSharedMemory* QSharedMemory_new3(struct miqt_string key);
extern __declspec(dllexport) QSharedMemory* QSharedMemory_new4(QObject* parent);
extern __declspec(dllexport) QSharedMemory* QSharedMemory_new5(QNativeIpcKey* key, QObject* parent);
extern __declspec(dllexport) QSharedMemory* QSharedMemory_new6(struct miqt_string key, QObject* parent);
extern __declspec(dllexport) void QSharedMemory_virtbase(QSharedMemory* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSharedMemory_MetaObject(const QSharedMemory* self);
extern __declspec(dllexport) void* QSharedMemory_Metacast(QSharedMemory* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSharedMemory_Tr(const char* s);
extern __declspec(dllexport) void QSharedMemory_SetKey(QSharedMemory* self, struct miqt_string key);
extern __declspec(dllexport) struct miqt_string QSharedMemory_Key(const QSharedMemory* self);
extern __declspec(dllexport) void QSharedMemory_SetNativeKey(QSharedMemory* self, QNativeIpcKey* key);
extern __declspec(dllexport) void QSharedMemory_SetNativeKeyWithKey(QSharedMemory* self, struct miqt_string key);
extern __declspec(dllexport) struct miqt_string QSharedMemory_NativeKey(const QSharedMemory* self);
extern __declspec(dllexport) QNativeIpcKey* QSharedMemory_NativeIpcKey(const QSharedMemory* self);
extern __declspec(dllexport) bool QSharedMemory_Create(QSharedMemory* self, ptrdiff_t size);
extern __declspec(dllexport) ptrdiff_t QSharedMemory_Size(const QSharedMemory* self);
extern __declspec(dllexport) bool QSharedMemory_Attach(QSharedMemory* self);
extern __declspec(dllexport) bool QSharedMemory_IsAttached(const QSharedMemory* self);
extern __declspec(dllexport) bool QSharedMemory_Detach(QSharedMemory* self);
extern __declspec(dllexport) void* QSharedMemory_Data(QSharedMemory* self);
extern __declspec(dllexport) const void* QSharedMemory_ConstData(const QSharedMemory* self);
extern __declspec(dllexport) const void* QSharedMemory_Data2(const QSharedMemory* self);
extern __declspec(dllexport) bool QSharedMemory_Lock(QSharedMemory* self);
extern __declspec(dllexport) bool QSharedMemory_Unlock(QSharedMemory* self);
extern __declspec(dllexport) SharedMemoryError QSharedMemory_Error(const QSharedMemory* self);
extern __declspec(dllexport) struct miqt_string QSharedMemory_ErrorString(const QSharedMemory* self);
extern __declspec(dllexport) bool QSharedMemory_IsKeyTypeSupported(uint16_t typeVal);
extern __declspec(dllexport) QNativeIpcKey* QSharedMemory_PlatformSafeKey(struct miqt_string key);
extern __declspec(dllexport) QNativeIpcKey* QSharedMemory_LegacyNativeKey(struct miqt_string key);
extern __declspec(dllexport) struct miqt_string QSharedMemory_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSharedMemory_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSharedMemory_SetNativeKey2(QSharedMemory* self, struct miqt_string key, uint16_t typeVal);
extern __declspec(dllexport) bool QSharedMemory_Create2(QSharedMemory* self, ptrdiff_t size, AccessMode mode);
extern __declspec(dllexport) bool QSharedMemory_Attach1(QSharedMemory* self, AccessMode mode);
extern __declspec(dllexport) QNativeIpcKey* QSharedMemory_PlatformSafeKey2(struct miqt_string key, uint16_t typeVal);
extern __declspec(dllexport) QNativeIpcKey* QSharedMemory_LegacyNativeKey2(struct miqt_string key, uint16_t typeVal);
extern __declspec(dllexport) void QSharedMemory_override_virtual_Event(void* self, intptr_t slot);
bool QSharedMemory_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSharedMemory_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSharedMemory_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QSharedMemory_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSharedMemory_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QSharedMemory_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSharedMemory_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QSharedMemory_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSharedMemory_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSharedMemory_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSharedMemory_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSharedMemory_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSharedMemory_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSharedMemory_Delete(QSharedMemory* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
