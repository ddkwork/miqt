#pragma once
#ifndef MIQT_QT_GEN_QDATETIME_H
#define MIQT_QT_GEN_QDATETIME_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCalendar QCalendar;
typedef struct QDate QDate;
typedef struct QDateTime QDateTime;
typedef struct QTime QTime;
typedef struct QTimeZone QTimeZone;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDate* QDate_new();
extern __declspec(dllexport) QDate* QDate_new2(int y, int m, int d);
extern __declspec(dllexport) QDate* QDate_new3(int y, int m, int d, QCalendar* cal);
extern __declspec(dllexport) QDate* QDate_new4(QDate* param1);
extern __declspec(dllexport) bool QDate_IsNull(const QDate* self);
extern __declspec(dllexport) bool QDate_IsValid(const QDate* self);
extern __declspec(dllexport) int QDate_Year(const QDate* self);
extern __declspec(dllexport) int QDate_Month(const QDate* self);
extern __declspec(dllexport) int QDate_Day(const QDate* self);
extern __declspec(dllexport) int QDate_DayOfWeek(const QDate* self);
extern __declspec(dllexport) int QDate_DayOfYear(const QDate* self);
extern __declspec(dllexport) int QDate_DaysInMonth(const QDate* self);
extern __declspec(dllexport) int QDate_DaysInYear(const QDate* self);
extern __declspec(dllexport) int QDate_WeekNumber(const QDate* self);
extern __declspec(dllexport) int QDate_YearWithCal(const QDate* self, QCalendar* cal);
extern __declspec(dllexport) int QDate_MonthWithCal(const QDate* self, QCalendar* cal);
extern __declspec(dllexport) int QDate_DayWithCal(const QDate* self, QCalendar* cal);
extern __declspec(dllexport) int QDate_DayOfWeekWithCal(const QDate* self, QCalendar* cal);
extern __declspec(dllexport) int QDate_DayOfYearWithCal(const QDate* self, QCalendar* cal);
extern __declspec(dllexport) int QDate_DaysInMonthWithCal(const QDate* self, QCalendar* cal);
extern __declspec(dllexport) int QDate_DaysInYearWithCal(const QDate* self, QCalendar* cal);
extern __declspec(dllexport) QDateTime* QDate_StartOfDay(const QDate* self, int spec);
extern __declspec(dllexport) QDateTime* QDate_EndOfDay(const QDate* self, int spec);
extern __declspec(dllexport) QDateTime* QDate_StartOfDayWithZone(const QDate* self, QTimeZone* zone);
extern __declspec(dllexport) QDateTime* QDate_EndOfDayWithZone(const QDate* self, QTimeZone* zone);
extern __declspec(dllexport) QDateTime* QDate_StartOfDay2(const QDate* self);
extern __declspec(dllexport) QDateTime* QDate_EndOfDay2(const QDate* self);
extern __declspec(dllexport) struct miqt_string QDate_ToString(const QDate* self);
extern __declspec(dllexport) struct miqt_string QDate_ToStringWithFormat(const QDate* self, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QDate_ToString2(const QDate* self, struct miqt_string format, QCalendar* cal);
extern __declspec(dllexport) bool QDate_SetDate(QDate* self, int year, int month, int day);
extern __declspec(dllexport) bool QDate_SetDate2(QDate* self, int year, int month, int day, QCalendar* cal);
extern __declspec(dllexport) void QDate_GetDate(const QDate* self, int* year, int* month, int* day);
extern __declspec(dllexport) QDate* QDate_AddDays(const QDate* self, long long days);
extern __declspec(dllexport) QDate* QDate_AddMonths(const QDate* self, int months);
extern __declspec(dllexport) QDate* QDate_AddYears(const QDate* self, int years);
extern __declspec(dllexport) QDate* QDate_AddMonths2(const QDate* self, int months, QCalendar* cal);
extern __declspec(dllexport) QDate* QDate_AddYears2(const QDate* self, int years, QCalendar* cal);
extern __declspec(dllexport) long long QDate_DaysTo(const QDate* self, QDate* d);
extern __declspec(dllexport) QDate* QDate_CurrentDate();
extern __declspec(dllexport) QDate* QDate_FromStringWithStringVal(struct miqt_string stringVal);
extern __declspec(dllexport) QDate* QDate_FromString4(struct miqt_string stringVal, struct miqt_string format, QCalendar* cal);
extern __declspec(dllexport) QDate* QDate_FromString9(struct miqt_string stringVal, struct miqt_string format);
extern __declspec(dllexport) QDate* QDate_FromString10(struct miqt_string stringVal, struct miqt_string format, int baseYear, QCalendar* cal);
extern __declspec(dllexport) bool QDate_IsValid2(int y, int m, int d);
extern __declspec(dllexport) bool QDate_IsLeapYear(int year);
extern __declspec(dllexport) QDate* QDate_FromJulianDay(long long jd_);
extern __declspec(dllexport) long long QDate_ToJulianDay(const QDate* self);
extern __declspec(dllexport) int QDate_WeekNumber1(const QDate* self, int* yearNum);
extern __declspec(dllexport) QDateTime* QDate_StartOfDay22(const QDate* self, int spec, int offsetSeconds);
extern __declspec(dllexport) QDateTime* QDate_EndOfDay22(const QDate* self, int spec, int offsetSeconds);
extern __declspec(dllexport) struct miqt_string QDate_ToString1(const QDate* self, int format);
extern __declspec(dllexport) QDate* QDate_FromString23(struct miqt_string stringVal, int format);
extern __declspec(dllexport) QDate* QDate_FromString34(struct miqt_string stringVal, struct miqt_string format, int baseYear);
extern __declspec(dllexport) void QDate_Delete(QDate* self, bool isSubclass);

extern __declspec(dllexport) QTime* QTime_new();
extern __declspec(dllexport) QTime* QTime_new2(int h, int m);
extern __declspec(dllexport) QTime* QTime_new3(QTime* param1);
extern __declspec(dllexport) QTime* QTime_new4(int h, int m, int s);
extern __declspec(dllexport) QTime* QTime_new5(int h, int m, int s, int ms);
extern __declspec(dllexport) bool QTime_IsNull(const QTime* self);
extern __declspec(dllexport) bool QTime_IsValid(const QTime* self);
extern __declspec(dllexport) int QTime_Hour(const QTime* self);
extern __declspec(dllexport) int QTime_Minute(const QTime* self);
extern __declspec(dllexport) int QTime_Second(const QTime* self);
extern __declspec(dllexport) int QTime_Msec(const QTime* self);
extern __declspec(dllexport) struct miqt_string QTime_ToString(const QTime* self);
extern __declspec(dllexport) struct miqt_string QTime_ToStringWithFormat(const QTime* self, struct miqt_string format);
extern __declspec(dllexport) bool QTime_SetHMS(QTime* self, int h, int m, int s);
extern __declspec(dllexport) QTime* QTime_AddSecs(const QTime* self, int secs);
extern __declspec(dllexport) int QTime_SecsTo(const QTime* self, QTime* t);
extern __declspec(dllexport) QTime* QTime_AddMSecs(const QTime* self, int ms);
extern __declspec(dllexport) int QTime_MsecsTo(const QTime* self, QTime* t);
extern __declspec(dllexport) QTime* QTime_FromMSecsSinceStartOfDay(int msecs);
extern __declspec(dllexport) int QTime_MsecsSinceStartOfDay(const QTime* self);
extern __declspec(dllexport) QTime* QTime_CurrentTime();
extern __declspec(dllexport) QTime* QTime_FromStringWithStringVal(struct miqt_string stringVal);
extern __declspec(dllexport) QTime* QTime_FromString4(struct miqt_string stringVal, struct miqt_string format);
extern __declspec(dllexport) bool QTime_IsValid2(int h, int m, int s);
extern __declspec(dllexport) struct miqt_string QTime_ToString1(const QTime* self, int f);
extern __declspec(dllexport) bool QTime_SetHMS4(QTime* self, int h, int m, int s, int ms);
extern __declspec(dllexport) QTime* QTime_FromString23(struct miqt_string stringVal, int format);
extern __declspec(dllexport) bool QTime_IsValid4(int h, int m, int s, int ms);
extern __declspec(dllexport) void QTime_Delete(QTime* self, bool isSubclass);

extern __declspec(dllexport) QDateTime* QDateTime_new();
extern __declspec(dllexport) QDateTime* QDateTime_new2(QDate* date, QTime* time, int spec);
extern __declspec(dllexport) QDateTime* QDateTime_new3(QDate* date, QTime* time, QTimeZone* timeZone);
extern __declspec(dllexport) QDateTime* QDateTime_new4(QDate* date, QTime* time);
extern __declspec(dllexport) QDateTime* QDateTime_new5(QDateTime* other);
extern __declspec(dllexport) QDateTime* QDateTime_new6(QDate* date, QTime* time, int spec, int offsetSeconds);
extern __declspec(dllexport) QDateTime* QDateTime_new7(QDate* date, QTime* time, QTimeZone* timeZone, TransitionResolution resolve);
extern __declspec(dllexport) QDateTime* QDateTime_new8(QDate* date, QTime* time, TransitionResolution resolve);
extern __declspec(dllexport) void QDateTime_OperatorAssign(QDateTime* self, QDateTime* other);
extern __declspec(dllexport) void QDateTime_Swap(QDateTime* self, QDateTime* other);
extern __declspec(dllexport) bool QDateTime_IsNull(const QDateTime* self);
extern __declspec(dllexport) bool QDateTime_IsValid(const QDateTime* self);
extern __declspec(dllexport) QDate* QDateTime_Date(const QDateTime* self);
extern __declspec(dllexport) QTime* QDateTime_Time(const QDateTime* self);
extern __declspec(dllexport) int QDateTime_TimeSpec(const QDateTime* self);
extern __declspec(dllexport) int QDateTime_OffsetFromUtc(const QDateTime* self);
extern __declspec(dllexport) QTimeZone* QDateTime_TimeRepresentation(const QDateTime* self);
extern __declspec(dllexport) QTimeZone* QDateTime_TimeZone(const QDateTime* self);
extern __declspec(dllexport) struct miqt_string QDateTime_TimeZoneAbbreviation(const QDateTime* self);
extern __declspec(dllexport) bool QDateTime_IsDaylightTime(const QDateTime* self);
extern __declspec(dllexport) long long QDateTime_ToMSecsSinceEpoch(const QDateTime* self);
extern __declspec(dllexport) long long QDateTime_ToSecsSinceEpoch(const QDateTime* self);
extern __declspec(dllexport) void QDateTime_SetDate(QDateTime* self, QDate* date);
extern __declspec(dllexport) void QDateTime_SetTime(QDateTime* self, QTime* time);
extern __declspec(dllexport) void QDateTime_SetTimeSpec(QDateTime* self, int spec);
extern __declspec(dllexport) void QDateTime_SetOffsetFromUtc(QDateTime* self, int offsetSeconds);
extern __declspec(dllexport) void QDateTime_SetTimeZone(QDateTime* self, QTimeZone* toZone);
extern __declspec(dllexport) void QDateTime_SetMSecsSinceEpoch(QDateTime* self, long long msecs);
extern __declspec(dllexport) void QDateTime_SetSecsSinceEpoch(QDateTime* self, long long secs);
extern __declspec(dllexport) struct miqt_string QDateTime_ToString(const QDateTime* self);
extern __declspec(dllexport) struct miqt_string QDateTime_ToStringWithFormat(const QDateTime* self, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QDateTime_ToString2(const QDateTime* self, struct miqt_string format, QCalendar* cal);
extern __declspec(dllexport) QDateTime* QDateTime_AddDays(const QDateTime* self, long long days);
extern __declspec(dllexport) QDateTime* QDateTime_AddMonths(const QDateTime* self, int months);
extern __declspec(dllexport) QDateTime* QDateTime_AddYears(const QDateTime* self, int years);
extern __declspec(dllexport) QDateTime* QDateTime_AddSecs(const QDateTime* self, long long secs);
extern __declspec(dllexport) QDateTime* QDateTime_AddMSecs(const QDateTime* self, long long msecs);
extern __declspec(dllexport) QDateTime* QDateTime_ToTimeSpec(const QDateTime* self, int spec);
extern __declspec(dllexport) QDateTime* QDateTime_ToLocalTime(const QDateTime* self);
extern __declspec(dllexport) QDateTime* QDateTime_ToUTC(const QDateTime* self);
extern __declspec(dllexport) QDateTime* QDateTime_ToOffsetFromUtc(const QDateTime* self, int offsetSeconds);
extern __declspec(dllexport) QDateTime* QDateTime_ToTimeZone(const QDateTime* self, QTimeZone* toZone);
extern __declspec(dllexport) long long QDateTime_DaysTo(const QDateTime* self, QDateTime* param1);
extern __declspec(dllexport) long long QDateTime_SecsTo(const QDateTime* self, QDateTime* param1);
extern __declspec(dllexport) long long QDateTime_MsecsTo(const QDateTime* self, QDateTime* param1);
extern __declspec(dllexport) QDateTime* QDateTime_CurrentDateTime(QTimeZone* zone);
extern __declspec(dllexport) QDateTime* QDateTime_CurrentDateTime2();
extern __declspec(dllexport) QDateTime* QDateTime_CurrentDateTimeUtc();
extern __declspec(dllexport) QDateTime* QDateTime_FromStringWithStringVal(struct miqt_string stringVal);
extern __declspec(dllexport) QDateTime* QDateTime_FromString4(struct miqt_string stringVal, struct miqt_string format, QCalendar* cal);
extern __declspec(dllexport) QDateTime* QDateTime_FromString9(struct miqt_string stringVal, struct miqt_string format);
extern __declspec(dllexport) QDateTime* QDateTime_FromString10(struct miqt_string stringVal, struct miqt_string format, int baseYear, QCalendar* cal);
extern __declspec(dllexport) QDateTime* QDateTime_FromMSecsSinceEpoch(long long msecs, int spec);
extern __declspec(dllexport) QDateTime* QDateTime_FromSecsSinceEpoch(long long secs, int spec);
extern __declspec(dllexport) QDateTime* QDateTime_FromMSecsSinceEpoch2(long long msecs, QTimeZone* timeZone);
extern __declspec(dllexport) QDateTime* QDateTime_FromSecsSinceEpoch2(long long secs, QTimeZone* timeZone);
extern __declspec(dllexport) QDateTime* QDateTime_FromMSecsSinceEpochWithMsecs(long long msecs);
extern __declspec(dllexport) QDateTime* QDateTime_FromSecsSinceEpochWithSecs(long long secs);
extern __declspec(dllexport) long long QDateTime_CurrentMSecsSinceEpoch();
extern __declspec(dllexport) long long QDateTime_CurrentSecsSinceEpoch();
extern __declspec(dllexport) void QDateTime_SetDate2(QDateTime* self, QDate* date, TransitionResolution resolve);
extern __declspec(dllexport) void QDateTime_SetTime2(QDateTime* self, QTime* time, TransitionResolution resolve);
extern __declspec(dllexport) void QDateTime_SetTimeZone2(QDateTime* self, QTimeZone* toZone, TransitionResolution resolve);
extern __declspec(dllexport) struct miqt_string QDateTime_ToString1(const QDateTime* self, int format);
extern __declspec(dllexport) QDateTime* QDateTime_FromString23(struct miqt_string stringVal, int format);
extern __declspec(dllexport) QDateTime* QDateTime_FromString34(struct miqt_string stringVal, struct miqt_string format, int baseYear);
extern __declspec(dllexport) QDateTime* QDateTime_FromMSecsSinceEpoch3(long long msecs, int spec, int offsetFromUtc);
extern __declspec(dllexport) QDateTime* QDateTime_FromSecsSinceEpoch3(long long secs, int spec, int offsetFromUtc);
extern __declspec(dllexport) void QDateTime_Delete(QDateTime* self, bool isSubclass);

} 
