#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QMEDIATIMERANGE_H
#define MIQT_QT_MULTIMEDIA_GEN_QMEDIATIMERANGE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QMediaTimeRange__Interval)
typedef QMediaTimeRange::Interval QMediaTimeRange__Interval;
typedef struct QMediaTimeRange QMediaTimeRange;
typedef struct QMediaTimeRange__Interval QMediaTimeRange__Interval;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_new();
extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_new2(long long start, long long end);
extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_new3(const Interval* param1);
extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_new4(QMediaTimeRange* rangeVal);
extern __declspec(dllexport) 
void QMediaTimeRange_OperatorAssign(QMediaTimeRange* self, QMediaTimeRange* param1);
extern __declspec(dllexport) 
void QMediaTimeRange_Swap(QMediaTimeRange* self, QMediaTimeRange* other);
extern __declspec(dllexport) 
void QMediaTimeRange_Detach(QMediaTimeRange* self);
extern __declspec(dllexport) 
void QMediaTimeRange_OperatorAssignWithInterval(QMediaTimeRange* self, const Interval* param1);
extern __declspec(dllexport) 
long long QMediaTimeRange_EarliestTime(const QMediaTimeRange* self);
extern __declspec(dllexport) 
long long QMediaTimeRange_LatestTime(const QMediaTimeRange* self);
extern __declspec(dllexport) 
struct miqt_array /* of QMediaTimeRange__Interval* */  QMediaTimeRange_Intervals(const QMediaTimeRange* self);
extern __declspec(dllexport) 
bool QMediaTimeRange_IsEmpty(const QMediaTimeRange* self);
extern __declspec(dllexport) 
bool QMediaTimeRange_IsContinuous(const QMediaTimeRange* self);
extern __declspec(dllexport) 
bool QMediaTimeRange_Contains(const QMediaTimeRange* self, long long time);
extern __declspec(dllexport) 
void QMediaTimeRange_AddInterval(QMediaTimeRange* self, long long start, long long end);
extern __declspec(dllexport) 
void QMediaTimeRange_AddIntervalWithInterval(QMediaTimeRange* self, const Interval* interval);
extern __declspec(dllexport) 
void QMediaTimeRange_AddTimeRange(QMediaTimeRange* self, QMediaTimeRange* param1);
extern __declspec(dllexport) 
void QMediaTimeRange_RemoveInterval(QMediaTimeRange* self, long long start, long long end);
extern __declspec(dllexport) 
void QMediaTimeRange_RemoveIntervalWithInterval(QMediaTimeRange* self, const Interval* interval);
extern __declspec(dllexport) 
void QMediaTimeRange_RemoveTimeRange(QMediaTimeRange* self, QMediaTimeRange* param1);
extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_OperatorPlusAssign(QMediaTimeRange* self, QMediaTimeRange* param1);
extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_OperatorPlusAssignWithInterval(QMediaTimeRange* self, const Interval* param1);
extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_OperatorMinusAssign(QMediaTimeRange* self, QMediaTimeRange* param1);
extern __declspec(dllexport) 
QMediaTimeRange* QMediaTimeRange_OperatorMinusAssignWithInterval(QMediaTimeRange* self, const Interval* param1);
extern __declspec(dllexport) 
void QMediaTimeRange_Clear(QMediaTimeRange* self);
extern __declspec(dllexport) 
void QMediaTimeRange_Delete(QMediaTimeRange* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaTimeRange__Interval* QMediaTimeRange__Interval_new();
extern __declspec(dllexport) 
QMediaTimeRange__Interval* QMediaTimeRange__Interval_new2(long long start, long long end);
extern __declspec(dllexport) 
QMediaTimeRange__Interval* QMediaTimeRange__Interval_new3(const Interval* param1);
extern __declspec(dllexport) 
long long QMediaTimeRange__Interval_Start(const QMediaTimeRange__Interval* self);
extern __declspec(dllexport) 
long long QMediaTimeRange__Interval_End(const QMediaTimeRange__Interval* self);
extern __declspec(dllexport) 
bool QMediaTimeRange__Interval_Contains(const QMediaTimeRange__Interval* self, long long time);
extern __declspec(dllexport) 
bool QMediaTimeRange__Interval_IsNormal(const QMediaTimeRange__Interval* self);
extern __declspec(dllexport) 
Interval QMediaTimeRange__Interval_Normalized(const QMediaTimeRange__Interval* self);
extern __declspec(dllexport) 
Interval QMediaTimeRange__Interval_Translated(const QMediaTimeRange__Interval* self, long long offset);
extern __declspec(dllexport) 
void QMediaTimeRange__Interval_Delete(QMediaTimeRange__Interval* self, bool isSubclass);

}
