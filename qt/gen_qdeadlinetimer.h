#pragma once
#ifndef MIQT_QT_GEN_QDEADLINETIMER_H
#define MIQT_QT_GEN_QDEADLINETIMER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_new();
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_new2(int type_);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_new3(ForeverConstant param1);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_new4(long long msecs);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_new5(QDeadlineTimer* param1);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_new6(ForeverConstant param1, int type_);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_new7(long long msecs, int typeVal);
extern __declspec(dllexport) void QDeadlineTimer_Swap(QDeadlineTimer* self, QDeadlineTimer* other);
extern __declspec(dllexport) bool QDeadlineTimer_IsForever(const QDeadlineTimer* self);
extern __declspec(dllexport) bool QDeadlineTimer_HasExpired(const QDeadlineTimer* self);
extern __declspec(dllexport) int QDeadlineTimer_TimerType(const QDeadlineTimer* self);
extern __declspec(dllexport) void QDeadlineTimer_SetTimerType(QDeadlineTimer* self, int typeVal);
extern __declspec(dllexport) long long QDeadlineTimer_RemainingTime(const QDeadlineTimer* self);
extern __declspec(dllexport) long long QDeadlineTimer_RemainingTimeNSecs(const QDeadlineTimer* self);
extern __declspec(dllexport) void QDeadlineTimer_SetRemainingTime(QDeadlineTimer* self, long long msecs);
extern __declspec(dllexport) void QDeadlineTimer_SetPreciseRemainingTime(QDeadlineTimer* self, long long secs);
extern __declspec(dllexport) long long QDeadlineTimer_Deadline(const QDeadlineTimer* self);
extern __declspec(dllexport) long long QDeadlineTimer_DeadlineNSecs(const QDeadlineTimer* self);
extern __declspec(dllexport) void QDeadlineTimer_SetDeadline(QDeadlineTimer* self, long long msecs);
extern __declspec(dllexport) void QDeadlineTimer_SetPreciseDeadline(QDeadlineTimer* self, long long secs);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_AddNSecs(QDeadlineTimer* dt, long long nsecs);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_Current();
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_OperatorPlusAssign(QDeadlineTimer* self, long long msecs);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_OperatorMinusAssign(QDeadlineTimer* self, long long msecs);
extern __declspec(dllexport) void QDeadlineTimer_SetRemainingTime2(QDeadlineTimer* self, long long msecs, int typeVal);
extern __declspec(dllexport) void QDeadlineTimer_SetPreciseRemainingTime2(QDeadlineTimer* self, long long secs, long long nsecs);
extern __declspec(dllexport) void QDeadlineTimer_SetPreciseRemainingTime3(QDeadlineTimer* self, long long secs, long long nsecs, int typeVal);
extern __declspec(dllexport) void QDeadlineTimer_SetDeadline2(QDeadlineTimer* self, long long msecs, int timerType);
extern __declspec(dllexport) void QDeadlineTimer_SetPreciseDeadline2(QDeadlineTimer* self, long long secs, long long nsecs);
extern __declspec(dllexport) void QDeadlineTimer_SetPreciseDeadline3(QDeadlineTimer* self, long long secs, long long nsecs, int typeVal);
extern __declspec(dllexport) QDeadlineTimer* QDeadlineTimer_Current1(int timerType);
extern __declspec(dllexport) void QDeadlineTimer_Delete(QDeadlineTimer* self, bool isSubclass);

} 
