#pragma once
#ifndef MIQT_QT_GEN_QREADWRITELOCK_H
#define MIQT_QT_GEN_QREADWRITELOCK_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QReadLocker QReadLocker;
typedef struct QReadWriteLock QReadWriteLock;
typedef struct QWriteLocker QWriteLocker;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QReadWriteLock* QReadWriteLock_new();
extern __declspec(dllexport) QReadWriteLock* QReadWriteLock_new2(RecursionMode recursionMode);
extern __declspec(dllexport) void QReadWriteLock_LockForRead(QReadWriteLock* self);
extern __declspec(dllexport) bool QReadWriteLock_TryLockForRead(QReadWriteLock* self, int timeout);
extern __declspec(dllexport) bool QReadWriteLock_TryLockForRead2(QReadWriteLock* self);
extern __declspec(dllexport) void QReadWriteLock_LockForWrite(QReadWriteLock* self);
extern __declspec(dllexport) bool QReadWriteLock_TryLockForWrite(QReadWriteLock* self, int timeout);
extern __declspec(dllexport) bool QReadWriteLock_TryLockForWrite2(QReadWriteLock* self);
extern __declspec(dllexport) void QReadWriteLock_Unlock(QReadWriteLock* self);
extern __declspec(dllexport) bool QReadWriteLock_TryLockForRead1(QReadWriteLock* self, QDeadlineTimer* timeout);
extern __declspec(dllexport) bool QReadWriteLock_TryLockForWrite1(QReadWriteLock* self, QDeadlineTimer* timeout);
extern __declspec(dllexport) void QReadWriteLock_Delete(QReadWriteLock* self, bool isSubclass);

extern __declspec(dllexport) QReadLocker* QReadLocker_new(QReadWriteLock* readWriteLock);
extern __declspec(dllexport) void QReadLocker_Unlock(QReadLocker* self);
extern __declspec(dllexport) void QReadLocker_Relock(QReadLocker* self);
extern __declspec(dllexport) QReadWriteLock* QReadLocker_ReadWriteLock(const QReadLocker* self);
extern __declspec(dllexport) void QReadLocker_Delete(QReadLocker* self, bool isSubclass);

extern __declspec(dllexport) QWriteLocker* QWriteLocker_new(QReadWriteLock* readWriteLock);
extern __declspec(dllexport) void QWriteLocker_Unlock(QWriteLocker* self);
extern __declspec(dllexport) void QWriteLocker_Relock(QWriteLocker* self);
extern __declspec(dllexport) QReadWriteLock* QWriteLocker_ReadWriteLock(const QWriteLocker* self);
extern __declspec(dllexport) void QWriteLocker_Delete(QWriteLocker* self, bool isSubclass);

} 
