#pragma once
#ifndef MIQT_QT_GEN_QTIMEZONE_H
#define MIQT_QT_GEN_QTIMEZONE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTimeZone__OffsetData)
typedef QTimeZone::OffsetData QTimeZone__OffsetData;
typedef struct QByteArrayView QByteArrayView;
typedef struct QDateTime QDateTime;
typedef struct QLocale QLocale;
typedef struct QTimeZone QTimeZone;
typedef struct QTimeZone__OffsetData QTimeZone__OffsetData;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTimeZone* QTimeZone_new();
extern __declspec(dllexport) QTimeZone* QTimeZone_new2(Initialization spec);
extern __declspec(dllexport) QTimeZone* QTimeZone_new3(int offsetSeconds);
extern __declspec(dllexport) QTimeZone* QTimeZone_new4(struct miqt_string ianaId);
extern __declspec(dllexport) QTimeZone* QTimeZone_new5(struct miqt_string zoneId, int offsetSeconds, struct miqt_string name, struct miqt_string abbreviation);
extern __declspec(dllexport) QTimeZone* QTimeZone_new6(QTimeZone* other);
extern __declspec(dllexport) QTimeZone* QTimeZone_new7(struct miqt_string zoneId, int offsetSeconds, struct miqt_string name, struct miqt_string abbreviation, Country territory);
extern __declspec(dllexport) QTimeZone* QTimeZone_new8(struct miqt_string zoneId, int offsetSeconds, struct miqt_string name, struct miqt_string abbreviation, Country territory, struct miqt_string comment);
extern __declspec(dllexport) void QTimeZone_OperatorAssign(QTimeZone* self, QTimeZone* other);
extern __declspec(dllexport) void QTimeZone_Swap(QTimeZone* self, QTimeZone* other);
extern __declspec(dllexport) bool QTimeZone_IsValid(const QTimeZone* self);
extern __declspec(dllexport) QTimeZone* QTimeZone_FromSecondsAheadOfUtc(int offset);
extern __declspec(dllexport) int QTimeZone_TimeSpec(const QTimeZone* self);
extern __declspec(dllexport) int QTimeZone_FixedSecondsAheadOfUtc(const QTimeZone* self);
extern __declspec(dllexport) bool QTimeZone_IsUtcOrFixedOffset(int spec);
extern __declspec(dllexport) bool QTimeZone_IsUtcOrFixedOffset2(const QTimeZone* self);
extern __declspec(dllexport) QTimeZone* QTimeZone_AsBackendZone(const QTimeZone* self);
extern __declspec(dllexport) bool QTimeZone_HasAlternativeName(const QTimeZone* self, QByteArrayView* alias);
extern __declspec(dllexport) struct miqt_string QTimeZone_Id(const QTimeZone* self);
extern __declspec(dllexport) Country QTimeZone_Territory(const QTimeZone* self);
extern __declspec(dllexport) uint16_t QTimeZone_Country(const QTimeZone* self);
extern __declspec(dllexport) struct miqt_string QTimeZone_Comment(const QTimeZone* self);
extern __declspec(dllexport) struct miqt_string QTimeZone_DisplayName(const QTimeZone* self, QDateTime* atDateTime);
extern __declspec(dllexport) struct miqt_string QTimeZone_DisplayNameWithTimeType(const QTimeZone* self, TimeType timeType);
extern __declspec(dllexport) struct miqt_string QTimeZone_Abbreviation(const QTimeZone* self, QDateTime* atDateTime);
extern __declspec(dllexport) int QTimeZone_OffsetFromUtc(const QTimeZone* self, QDateTime* atDateTime);
extern __declspec(dllexport) int QTimeZone_StandardTimeOffset(const QTimeZone* self, QDateTime* atDateTime);
extern __declspec(dllexport) int QTimeZone_DaylightTimeOffset(const QTimeZone* self, QDateTime* atDateTime);
extern __declspec(dllexport) bool QTimeZone_HasDaylightTime(const QTimeZone* self);
extern __declspec(dllexport) bool QTimeZone_IsDaylightTime(const QTimeZone* self, QDateTime* atDateTime);
extern __declspec(dllexport) OffsetData QTimeZone_OffsetData(const QTimeZone* self, QDateTime* forDateTime);
extern __declspec(dllexport) bool QTimeZone_HasTransitions(const QTimeZone* self);
extern __declspec(dllexport) OffsetData QTimeZone_NextTransition(const QTimeZone* self, QDateTime* afterDateTime);
extern __declspec(dllexport) OffsetData QTimeZone_PreviousTransition(const QTimeZone* self, QDateTime* beforeDateTime);
extern __declspec(dllexport) OffsetDataList QTimeZone_Transitions(const QTimeZone* self, QDateTime* fromDateTime, QDateTime* toDateTime);
extern __declspec(dllexport) struct miqt_string QTimeZone_SystemTimeZoneId();
extern __declspec(dllexport) QTimeZone* QTimeZone_SystemTimeZone();
extern __declspec(dllexport) QTimeZone* QTimeZone_Utc();
extern __declspec(dllexport) bool QTimeZone_IsTimeZoneIdAvailable(struct miqt_string ianaId);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QTimeZone_AvailableTimeZoneIds();
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QTimeZone_AvailableTimeZoneIdsWithTerritory(Country territory);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QTimeZone_AvailableTimeZoneIdsWithOffsetSeconds(int offsetSeconds);
extern __declspec(dllexport) struct miqt_string QTimeZone_IanaIdToWindowsId(struct miqt_string ianaId);
extern __declspec(dllexport) struct miqt_string QTimeZone_WindowsIdToDefaultIanaId(struct miqt_string windowsId);
extern __declspec(dllexport) struct miqt_string QTimeZone_WindowsIdToDefaultIanaId2(struct miqt_string windowsId, Country territory);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QTimeZone_WindowsIdToIanaIds(struct miqt_string windowsId);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QTimeZone_WindowsIdToIanaIds2(struct miqt_string windowsId, Country territory);
extern __declspec(dllexport) struct miqt_string QTimeZone_DisplayName2(const QTimeZone* self, QDateTime* atDateTime, NameType nameType);
extern __declspec(dllexport) struct miqt_string QTimeZone_DisplayName3(const QTimeZone* self, QDateTime* atDateTime, NameType nameType, QLocale* locale);
extern __declspec(dllexport) struct miqt_string QTimeZone_DisplayName22(const QTimeZone* self, TimeType timeType, NameType nameType);
extern __declspec(dllexport) struct miqt_string QTimeZone_DisplayName32(const QTimeZone* self, TimeType timeType, NameType nameType, QLocale* locale);
extern __declspec(dllexport) void QTimeZone_Delete(QTimeZone* self, bool isSubclass);

extern __declspec(dllexport) QTimeZone__OffsetData* QTimeZone__OffsetData_new(const OffsetData* param1);
extern __declspec(dllexport) QTimeZone__OffsetData* QTimeZone__OffsetData_new2();
extern __declspec(dllexport) void QTimeZone__OffsetData_OperatorAssign(QTimeZone__OffsetData* self, const OffsetData* param1);
extern __declspec(dllexport) void QTimeZone__OffsetData_Delete(QTimeZone__OffsetData* self, bool isSubclass);

} 
