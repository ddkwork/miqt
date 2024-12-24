#pragma once
#ifndef MIQT_QT_GEN_QBASICTIMER_H
#define MIQT_QT_GEN_QBASICTIMER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBasicTimer QBasicTimer;
typedef struct QObject QObject;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QBasicTimer* QBasicTimer_new();
extern __declspec(dllexport) void QBasicTimer_Swap(QBasicTimer* self, QBasicTimer* other);
extern __declspec(dllexport) bool QBasicTimer_IsActive(const QBasicTimer* self);
extern __declspec(dllexport) int QBasicTimer_TimerId(const QBasicTimer* self);
extern __declspec(dllexport) int QBasicTimer_Id(const QBasicTimer* self);
extern __declspec(dllexport) void QBasicTimer_Start(QBasicTimer* self, int msec, QObject* obj);
extern __declspec(dllexport) void QBasicTimer_Start2(QBasicTimer* self, int msec, int timerType, QObject* obj);
extern __declspec(dllexport) void QBasicTimer_Start3(QBasicTimer* self, Duration duration, QObject* obj);
extern __declspec(dllexport) void QBasicTimer_Start4(QBasicTimer* self, Duration duration, int timerType, QObject* obj);
extern __declspec(dllexport) void QBasicTimer_Stop(QBasicTimer* self);
extern __declspec(dllexport) void QBasicTimer_Delete(QBasicTimer* self, bool isSubclass);

} 
