#pragma once
#ifndef MIQT_QT_GEN_QSEMAPHORE_H
#define MIQT_QT_GEN_QSEMAPHORE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QSemaphore QSemaphore;
typedef struct QSemaphoreReleaser QSemaphoreReleaser;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSemaphore* QSemaphore_new();
extern __declspec(dllexport) QSemaphore* QSemaphore_new2(int n);
extern __declspec(dllexport) void QSemaphore_Acquire(QSemaphore* self);
extern __declspec(dllexport) bool QSemaphore_TryAcquire(QSemaphore* self);
extern __declspec(dllexport) bool QSemaphore_TryAcquire2(QSemaphore* self, int n, int timeout);
extern __declspec(dllexport) bool QSemaphore_TryAcquire3(QSemaphore* self, int n, QDeadlineTimer* timeout);
extern __declspec(dllexport) void QSemaphore_Release(QSemaphore* self);
extern __declspec(dllexport) int QSemaphore_Available(const QSemaphore* self);
extern __declspec(dllexport) bool QSemaphore_TryAcquire4(QSemaphore* self);
extern __declspec(dllexport) void QSemaphore_Acquire1(QSemaphore* self, int n);
extern __declspec(dllexport) bool QSemaphore_TryAcquire1(QSemaphore* self, int n);
extern __declspec(dllexport) void QSemaphore_Release1(QSemaphore* self, int n);
extern __declspec(dllexport) void QSemaphore_Delete(QSemaphore* self, bool isSubclass);

extern __declspec(dllexport) QSemaphoreReleaser* QSemaphoreReleaser_new();
extern __declspec(dllexport) QSemaphoreReleaser* QSemaphoreReleaser_new2(QSemaphore* sem);
extern __declspec(dllexport) QSemaphoreReleaser* QSemaphoreReleaser_new3(QSemaphore* sem);
extern __declspec(dllexport) QSemaphoreReleaser* QSemaphoreReleaser_new4(QSemaphore* sem, int n);
extern __declspec(dllexport) QSemaphoreReleaser* QSemaphoreReleaser_new5(QSemaphore* sem, int n);
extern __declspec(dllexport) void QSemaphoreReleaser_Swap(QSemaphoreReleaser* self, QSemaphoreReleaser* other);
extern __declspec(dllexport) QSemaphore* QSemaphoreReleaser_Semaphore(const QSemaphoreReleaser* self);
extern __declspec(dllexport) QSemaphore* QSemaphoreReleaser_Cancel(QSemaphoreReleaser* self);
extern __declspec(dllexport) void QSemaphoreReleaser_Delete(QSemaphoreReleaser* self, bool isSubclass);

} 
