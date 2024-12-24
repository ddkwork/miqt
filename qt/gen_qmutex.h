#pragma once
#ifndef MIQT_QT_GEN_QMUTEX_H
#define MIQT_QT_GEN_QMUTEX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBasicMutex;
class QDeadlineTimer;
class QMutex;
class QRecursiveMutex;
class _GUID;
class type_info;
#else
typedef struct QBasicMutex QBasicMutex;
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QMutex QMutex;
typedef struct QRecursiveMutex QRecursiveMutex;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QBasicMutex* QBasicMutex_new();
extern __declspec(dllexport) void QBasicMutex_Lock(QBasicMutex* self);
extern __declspec(dllexport) void QBasicMutex_Unlock(QBasicMutex* self);
extern __declspec(dllexport) bool QBasicMutex_TryLock(QBasicMutex* self);
extern __declspec(dllexport) bool QBasicMutex_TryLock2(QBasicMutex* self);
extern __declspec(dllexport) void QBasicMutex_Delete(QBasicMutex* self, bool isSubclass);

extern __declspec(dllexport) QMutex* QMutex_new();
extern __declspec(dllexport) void QMutex_virtbase(QMutex* src, QBasicMutex** outptr_QBasicMutex);
extern __declspec(dllexport) bool QMutex_TryLock(QMutex* self);
extern __declspec(dllexport) bool QMutex_TryLockWithTimeout(QMutex* self, int timeout);
extern __declspec(dllexport) bool QMutex_TryLock2(QMutex* self, QDeadlineTimer* timeout);
extern __declspec(dllexport) void QMutex_Delete(QMutex* self, bool isSubclass);

extern __declspec(dllexport) QRecursiveMutex* QRecursiveMutex_new();
extern __declspec(dllexport) void QRecursiveMutex_Lock(QRecursiveMutex* self);
extern __declspec(dllexport) bool QRecursiveMutex_TryLock(QRecursiveMutex* self, int timeout);
extern __declspec(dllexport) bool QRecursiveMutex_TryLock2(QRecursiveMutex* self);
extern __declspec(dllexport) void QRecursiveMutex_Unlock(QRecursiveMutex* self);
extern __declspec(dllexport) bool QRecursiveMutex_TryLock3(QRecursiveMutex* self);
extern __declspec(dllexport) bool QRecursiveMutex_TryLock1(QRecursiveMutex* self, QDeadlineTimer* timer);
extern __declspec(dllexport) void QRecursiveMutex_Delete(QRecursiveMutex* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
