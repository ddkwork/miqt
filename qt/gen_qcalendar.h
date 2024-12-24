#pragma once
#ifndef MIQT_QT_GEN_QCALENDAR_H
#define MIQT_QT_GEN_QCALENDAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QCalendar__SystemId)
typedef QCalendar::SystemId QCalendar__SystemId;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QCalendar__YearMonthDay)
typedef QCalendar::YearMonthDay QCalendar__YearMonthDay;
typedef struct QAnyStringView QAnyStringView;
typedef struct QCalendar QCalendar;
typedef struct QCalendar__SystemId QCalendar__SystemId;
typedef struct QCalendar__YearMonthDay QCalendar__YearMonthDay;
typedef struct QDate QDate;
typedef struct QLocale QLocale;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCalendar* QCalendar_new();
extern __declspec(dllexport) 
QCalendar* QCalendar_new2(System system);
extern __declspec(dllexport) 
QCalendar* QCalendar_new3(QAnyStringView* name);
extern __declspec(dllexport) 
QCalendar* QCalendar_new4(SystemId id);
extern __declspec(dllexport) 
bool QCalendar_IsValid(const QCalendar* self);
extern __declspec(dllexport) 
int QCalendar_DaysInMonth(const QCalendar* self, int month);
extern __declspec(dllexport) 
int QCalendar_DaysInYear(const QCalendar* self, int year);
extern __declspec(dllexport) 
int QCalendar_MonthsInYear(const QCalendar* self, int year);
extern __declspec(dllexport) 
bool QCalendar_IsDateValid(const QCalendar* self, int year, int month, int day);
extern __declspec(dllexport) 
bool QCalendar_IsLeapYear(const QCalendar* self, int year);
extern __declspec(dllexport) 
bool QCalendar_IsGregorian(const QCalendar* self);
extern __declspec(dllexport) 
bool QCalendar_IsLunar(const QCalendar* self);
extern __declspec(dllexport) 
bool QCalendar_IsLuniSolar(const QCalendar* self);
extern __declspec(dllexport) 
bool QCalendar_IsSolar(const QCalendar* self);
extern __declspec(dllexport) 
bool QCalendar_IsProleptic(const QCalendar* self);
extern __declspec(dllexport) 
bool QCalendar_HasYearZero(const QCalendar* self);
extern __declspec(dllexport) 
int QCalendar_MaximumDaysInMonth(const QCalendar* self);
extern __declspec(dllexport) 
int QCalendar_MinimumDaysInMonth(const QCalendar* self);
extern __declspec(dllexport) 
int QCalendar_MaximumMonthsInYear(const QCalendar* self);
extern __declspec(dllexport) 
struct miqt_string QCalendar_Name(const QCalendar* self);
extern __declspec(dllexport) 
QDate* QCalendar_DateFromParts(const QCalendar* self, int year, int month, int day);
extern __declspec(dllexport) 
QDate* QCalendar_DateFromPartsWithParts(const QCalendar* self, const YearMonthDay* parts);
extern __declspec(dllexport) 
QDate* QCalendar_MatchCenturyToWeekday(const QCalendar* self, const YearMonthDay* parts, int dow);
extern __declspec(dllexport) 
YearMonthDay QCalendar_PartsFromDate(const QCalendar* self, QDate* date);
extern __declspec(dllexport) 
int QCalendar_DayOfWeek(const QCalendar* self, QDate* date);
extern __declspec(dllexport) 
struct miqt_string QCalendar_MonthName(const QCalendar* self, QLocale* locale, int month);
extern __declspec(dllexport) 
struct miqt_string QCalendar_StandaloneMonthName(const QCalendar* self, QLocale* locale, int month);
extern __declspec(dllexport) 
struct miqt_string QCalendar_WeekDayName(const QCalendar* self, QLocale* locale, int day);
extern __declspec(dllexport) 
struct miqt_string QCalendar_StandaloneWeekDayName(const QCalendar* self, QLocale* locale, int day);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCalendar_AvailableCalendars();
extern __declspec(dllexport) 
int QCalendar_DaysInMonth2(const QCalendar* self, int month, int year);
extern __declspec(dllexport) 
struct miqt_string QCalendar_MonthName3(const QCalendar* self, QLocale* locale, int month, int year);
extern __declspec(dllexport) 
struct miqt_string QCalendar_MonthName4(const QCalendar* self, QLocale* locale, int month, int year, int format);
extern __declspec(dllexport) 
struct miqt_string QCalendar_StandaloneMonthName3(const QCalendar* self, QLocale* locale, int month, int year);
extern __declspec(dllexport) 
struct miqt_string QCalendar_StandaloneMonthName4(const QCalendar* self, QLocale* locale, int month, int year, int format);
extern __declspec(dllexport) 
struct miqt_string QCalendar_WeekDayName3(const QCalendar* self, QLocale* locale, int day, int format);
extern __declspec(dllexport) 
struct miqt_string QCalendar_StandaloneWeekDayName3(const QCalendar* self, QLocale* locale, int day, int format);
extern __declspec(dllexport) 
void QCalendar_Delete(QCalendar* self, bool isSubclass);

extern __declspec(dllexport) 
QCalendar__YearMonthDay* QCalendar__YearMonthDay_new();
extern __declspec(dllexport) 
QCalendar__YearMonthDay* QCalendar__YearMonthDay_new2(int y);
extern __declspec(dllexport) 
QCalendar__YearMonthDay* QCalendar__YearMonthDay_new3(int y, int m);
extern __declspec(dllexport) 
QCalendar__YearMonthDay* QCalendar__YearMonthDay_new4(int y, int m, int d);
extern __declspec(dllexport) 
bool QCalendar__YearMonthDay_IsValid(const QCalendar__YearMonthDay* self);
extern __declspec(dllexport) 
void QCalendar__YearMonthDay_Delete(QCalendar__YearMonthDay* self, bool isSubclass);

extern __declspec(dllexport) 
QCalendar__SystemId* QCalendar__SystemId_new();
extern __declspec(dllexport) 
unsigned long long QCalendar__SystemId_Index(const QCalendar__SystemId* self);
extern __declspec(dllexport) 
bool QCalendar__SystemId_IsValid(const QCalendar__SystemId* self);
extern __declspec(dllexport) 
void QCalendar__SystemId_Delete(QCalendar__SystemId* self, bool isSubclass);

}
