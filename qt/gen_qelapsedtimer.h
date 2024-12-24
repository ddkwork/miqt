#pragma once
#ifndef MIQT_QT_GEN_QELAPSEDTIMER_H
#define MIQT_QT_GEN_QELAPSEDTIMER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QElapsedTimer QElapsedTimer;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QElapsedTimer* QElapsedTimer_new();
extern __declspec(dllexport) ClockType QElapsedTimer_ClockType();
extern __declspec(dllexport) bool QElapsedTimer_IsMonotonic();
extern __declspec(dllexport) void QElapsedTimer_Start(QElapsedTimer* self);
extern __declspec(dllexport) long long QElapsedTimer_Restart(QElapsedTimer* self);
extern __declspec(dllexport) void QElapsedTimer_Invalidate(QElapsedTimer* self);
extern __declspec(dllexport) bool QElapsedTimer_IsValid(const QElapsedTimer* self);
extern __declspec(dllexport) Duration QElapsedTimer_DurationElapsed(const QElapsedTimer* self);
extern __declspec(dllexport) long long QElapsedTimer_NsecsElapsed(const QElapsedTimer* self);
extern __declspec(dllexport) long long QElapsedTimer_Elapsed(const QElapsedTimer* self);
extern __declspec(dllexport) bool QElapsedTimer_HasExpired(const QElapsedTimer* self, long long timeout);
extern __declspec(dllexport) long long QElapsedTimer_MsecsSinceReference(const QElapsedTimer* self);
extern __declspec(dllexport) Duration QElapsedTimer_DurationTo(const QElapsedTimer* self, QElapsedTimer* other);
extern __declspec(dllexport) long long QElapsedTimer_MsecsTo(const QElapsedTimer* self, QElapsedTimer* other);
extern __declspec(dllexport) long long QElapsedTimer_SecsTo(const QElapsedTimer* self, QElapsedTimer* other);
extern __declspec(dllexport) void QElapsedTimer_Delete(QElapsedTimer* self, bool isSubclass);

} 
