#pragma once
#ifndef MIQT_QT_GEN_QWAITCONDITION_H
#define MIQT_QT_GEN_QWAITCONDITION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDeadlineTimer;
class QMutex;
class QReadWriteLock;
class QWaitCondition;
class _GUID;
class type_info;
#else
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QMutex QMutex;
typedef struct QReadWriteLock QReadWriteLock;
typedef struct QWaitCondition QWaitCondition;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QWaitCondition* QWaitCondition_new();
extern __declspec(dllexport) bool QWaitCondition_Wait(QWaitCondition* self, QMutex* lockedMutex);
extern __declspec(dllexport) bool QWaitCondition_Wait2(QWaitCondition* self, QMutex* lockedMutex, unsigned long time);
extern __declspec(dllexport) bool QWaitCondition_WaitWithLockedReadWriteLock(QWaitCondition* self, QReadWriteLock* lockedReadWriteLock);
extern __declspec(dllexport) bool QWaitCondition_Wait3(QWaitCondition* self, QReadWriteLock* lockedReadWriteLock, unsigned long time);
extern __declspec(dllexport) void QWaitCondition_WakeOne(QWaitCondition* self);
extern __declspec(dllexport) void QWaitCondition_WakeAll(QWaitCondition* self);
extern __declspec(dllexport) void QWaitCondition_NotifyOne(QWaitCondition* self);
extern __declspec(dllexport) void QWaitCondition_NotifyAll(QWaitCondition* self);
extern __declspec(dllexport) bool QWaitCondition_Wait22(QWaitCondition* self, QMutex* lockedMutex, QDeadlineTimer* deadline);
extern __declspec(dllexport) bool QWaitCondition_Wait23(QWaitCondition* self, QReadWriteLock* lockedReadWriteLock, QDeadlineTimer* deadline);
extern __declspec(dllexport) void QWaitCondition_Delete(QWaitCondition* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
