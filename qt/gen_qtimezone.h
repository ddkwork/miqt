#ifndef GEN_QTIMEZONE_H
#define GEN_QTIMEZONE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QDateTime;
class QLocale;
class QTimeZone;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTimeZone__OffsetData)
typedef QTimeZone::OffsetData QTimeZone__OffsetData;
#else
class QTimeZone__OffsetData;
#endif
#else
typedef struct QByteArray QByteArray;
typedef struct QDateTime QDateTime;
typedef struct QLocale QLocale;
typedef struct QTimeZone QTimeZone;
typedef struct QTimeZone__OffsetData QTimeZone__OffsetData;
#endif

QTimeZone* QTimeZone_new();
QTimeZone* QTimeZone_new2(QByteArray* ianaId);
QTimeZone* QTimeZone_new3(int offsetSeconds);
QTimeZone* QTimeZone_new4(QByteArray* zoneId, int offsetSeconds, const char* name, size_t name_Strlen, const char* abbreviation, size_t abbreviation_Strlen);
QTimeZone* QTimeZone_new5(QTimeZone* other);
QTimeZone* QTimeZone_new6(QByteArray* zoneId, int offsetSeconds, const char* name, size_t name_Strlen, const char* abbreviation, size_t abbreviation_Strlen, uintptr_t country);
QTimeZone* QTimeZone_new7(QByteArray* zoneId, int offsetSeconds, const char* name, size_t name_Strlen, const char* abbreviation, size_t abbreviation_Strlen, uintptr_t country, const char* comment, size_t comment_Strlen);
void QTimeZone_OperatorAssign(QTimeZone* self, QTimeZone* other);
void QTimeZone_Swap(QTimeZone* self, QTimeZone* other);
bool QTimeZone_OperatorEqual(QTimeZone* self, QTimeZone* other);
bool QTimeZone_OperatorNotEqual(QTimeZone* self, QTimeZone* other);
bool QTimeZone_IsValid(QTimeZone* self);
QByteArray* QTimeZone_Id(QTimeZone* self);
uintptr_t QTimeZone_Country(QTimeZone* self);
void QTimeZone_Comment(QTimeZone* self, char** _out, int* _out_Strlen);
void QTimeZone_DisplayName(QTimeZone* self, QDateTime* atDateTime, char** _out, int* _out_Strlen);
void QTimeZone_DisplayNameWithTimeType(QTimeZone* self, uintptr_t timeType, char** _out, int* _out_Strlen);
void QTimeZone_Abbreviation(QTimeZone* self, QDateTime* atDateTime, char** _out, int* _out_Strlen);
int QTimeZone_OffsetFromUtc(QTimeZone* self, QDateTime* atDateTime);
int QTimeZone_StandardTimeOffset(QTimeZone* self, QDateTime* atDateTime);
int QTimeZone_DaylightTimeOffset(QTimeZone* self, QDateTime* atDateTime);
bool QTimeZone_HasDaylightTime(QTimeZone* self);
bool QTimeZone_IsDaylightTime(QTimeZone* self, QDateTime* atDateTime);
QTimeZone__OffsetData* QTimeZone_OffsetData(QTimeZone* self, QDateTime* forDateTime);
bool QTimeZone_HasTransitions(QTimeZone* self);
QTimeZone__OffsetData* QTimeZone_NextTransition(QTimeZone* self, QDateTime* afterDateTime);
QTimeZone__OffsetData* QTimeZone_PreviousTransition(QTimeZone* self, QDateTime* beforeDateTime);
void QTimeZone_Transitions(QTimeZone* self, QDateTime* fromDateTime, QDateTime* toDateTime, QTimeZone__OffsetData*** _out, size_t* _out_len);
QByteArray* QTimeZone_SystemTimeZoneId();
QTimeZone* QTimeZone_SystemTimeZone();
QTimeZone* QTimeZone_Utc();
bool QTimeZone_IsTimeZoneIdAvailable(QByteArray* ianaId);
void QTimeZone_AvailableTimeZoneIds(QByteArray*** _out, size_t* _out_len);
void QTimeZone_AvailableTimeZoneIdsWithCountry(uintptr_t country, QByteArray*** _out, size_t* _out_len);
void QTimeZone_AvailableTimeZoneIdsWithOffsetSeconds(int offsetSeconds, QByteArray*** _out, size_t* _out_len);
QByteArray* QTimeZone_IanaIdToWindowsId(QByteArray* ianaId);
QByteArray* QTimeZone_WindowsIdToDefaultIanaId(QByteArray* windowsId);
QByteArray* QTimeZone_WindowsIdToDefaultIanaId2(QByteArray* windowsId, uintptr_t country);
void QTimeZone_WindowsIdToIanaIds(QByteArray* windowsId, QByteArray*** _out, size_t* _out_len);
void QTimeZone_WindowsIdToIanaIds2(QByteArray* windowsId, uintptr_t country, QByteArray*** _out, size_t* _out_len);
void QTimeZone_DisplayName2(QTimeZone* self, QDateTime* atDateTime, uintptr_t nameType, char** _out, int* _out_Strlen);
void QTimeZone_DisplayName3(QTimeZone* self, QDateTime* atDateTime, uintptr_t nameType, QLocale* locale, char** _out, int* _out_Strlen);
void QTimeZone_DisplayName22(QTimeZone* self, uintptr_t timeType, uintptr_t nameType, char** _out, int* _out_Strlen);
void QTimeZone_DisplayName32(QTimeZone* self, uintptr_t timeType, uintptr_t nameType, QLocale* locale, char** _out, int* _out_Strlen);
void QTimeZone_Delete(QTimeZone* self);

QTimeZone__OffsetData* QTimeZone__OffsetData_new(QTimeZone__OffsetData* param1);
void QTimeZone__OffsetData_Delete(QTimeZone__OffsetData* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif