#pragma once
#ifndef MIQT_QT_GEN_QLOCALE_H
#define MIQT_QT_GEN_QLOCALE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCalendar;
class QDate;
class QDateTime;
class QLocale;
class QTextStream;
class QTime;
class _GUID;
class type_info;
#else
typedef struct QCalendar QCalendar;
typedef struct QDate QDate;
typedef struct QDateTime QDateTime;
typedef struct QLocale QLocale;
typedef struct QTextStream QTextStream;
typedef struct QTime QTime;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QTextStream_Delete(QTextStream* self, bool isSubclass);

extern __declspec(dllexport) QLocale* QLocale_new();
extern __declspec(dllexport) QLocale* QLocale_new2(struct miqt_string name);
extern __declspec(dllexport) QLocale* QLocale_new3(Language language, Territory territory);
extern __declspec(dllexport) QLocale* QLocale_new4(Language language);
extern __declspec(dllexport) QLocale* QLocale_new5(QLocale* other);
extern __declspec(dllexport) QLocale* QLocale_new6(Language language, Script script);
extern __declspec(dllexport) QLocale* QLocale_new7(Language language, Script script, Territory territory);
extern __declspec(dllexport) void QLocale_OperatorAssign(QLocale* self, QLocale* other);
extern __declspec(dllexport) void QLocale_Swap(QLocale* self, QLocale* other);
extern __declspec(dllexport) Language QLocale_Language(const QLocale* self);
extern __declspec(dllexport) Script QLocale_Script(const QLocale* self);
extern __declspec(dllexport) Territory QLocale_Territory(const QLocale* self);
extern __declspec(dllexport) Country QLocale_Country(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_Name(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_Bcp47Name(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_NativeLanguageName(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_NativeTerritoryName(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_NativeCountryName(const QLocale* self);
extern __declspec(dllexport) int16_t QLocale_ToShort(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) uint16_t QLocale_ToUShort(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) int QLocale_ToInt(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) unsigned int QLocale_ToUInt(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) long QLocale_ToLong(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) unsigned long QLocale_ToULong(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) long long QLocale_ToLongLong(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) unsigned long long QLocale_ToULongLong(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) float QLocale_ToFloat(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) double QLocale_ToDouble(const QLocale* self, struct miqt_string s);
extern __declspec(dllexport) struct miqt_string QLocale_ToString(const QLocale* self, long long i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithQulonglong(const QLocale* self, unsigned long long i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithLong(const QLocale* self, long i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithUlong(const QLocale* self, unsigned long i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithShort(const QLocale* self, int16_t i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithUshort(const QLocale* self, uint16_t i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithInt(const QLocale* self, int i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithUint(const QLocale* self, unsigned int i);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithDouble(const QLocale* self, double f);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithFloat(const QLocale* self, float f);
extern __declspec(dllexport) struct miqt_string QLocale_ToString2(const QLocale* self, QDate* date, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QLocale_ToString3(const QLocale* self, QTime* time, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QLocale_ToString4(const QLocale* self, QDateTime* dateTime, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithDate(const QLocale* self, QDate* date);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithTime(const QLocale* self, QTime* time);
extern __declspec(dllexport) struct miqt_string QLocale_ToStringWithDateTime(const QLocale* self, QDateTime* dateTime);
extern __declspec(dllexport) struct miqt_string QLocale_ToString9(const QLocale* self, QDate* date, FormatType format, QCalendar* cal);
extern __declspec(dllexport) struct miqt_string QLocale_ToString10(const QLocale* self, QDateTime* dateTime, FormatType format, QCalendar* cal);
extern __declspec(dllexport) struct miqt_string QLocale_DateFormat(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_TimeFormat(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_DateTimeFormat(const QLocale* self);
extern __declspec(dllexport) QTime* QLocale_ToTime(const QLocale* self, struct miqt_string stringVal);
extern __declspec(dllexport) QTime* QLocale_ToTime2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format);
extern __declspec(dllexport) QDate* QLocale_ToDate(const QLocale* self, struct miqt_string stringVal);
extern __declspec(dllexport) QDate* QLocale_ToDate2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime(const QLocale* self, struct miqt_string stringVal);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format);
extern __declspec(dllexport) QDate* QLocale_ToDate3(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal);
extern __declspec(dllexport) QDate* QLocale_ToDate4(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime3(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime4(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal);
extern __declspec(dllexport) struct miqt_string QLocale_DecimalPoint(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_GroupSeparator(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_Percent(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_ZeroDigit(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_NegativeSign(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_PositiveSign(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_Exponential(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_MonthName(const QLocale* self, int param1);
extern __declspec(dllexport) struct miqt_string QLocale_StandaloneMonthName(const QLocale* self, int param1);
extern __declspec(dllexport) struct miqt_string QLocale_DayName(const QLocale* self, int param1);
extern __declspec(dllexport) struct miqt_string QLocale_StandaloneDayName(const QLocale* self, int param1);
extern __declspec(dllexport) int QLocale_FirstDayOfWeek(const QLocale* self);
extern __declspec(dllexport) struct miqt_array /* of int */  QLocale_Weekdays(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_AmText(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_PmText(const QLocale* self);
extern __declspec(dllexport) MeasurementSystem QLocale_MeasurementSystem(const QLocale* self);
extern __declspec(dllexport) QLocale* QLocale_Collation(const QLocale* self);
extern __declspec(dllexport) int QLocale_TextDirection(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_ToUpper(const QLocale* self, struct miqt_string str);
extern __declspec(dllexport) struct miqt_string QLocale_ToLower(const QLocale* self, struct miqt_string str);
extern __declspec(dllexport) struct miqt_string QLocale_CurrencySymbol(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString(const QLocale* self, long long param1);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyStringWithQulonglong(const QLocale* self, unsigned long long param1);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyStringWithShort(const QLocale* self, int16_t i);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyStringWithUshort(const QLocale* self, uint16_t i);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyStringWithInt(const QLocale* self, int i);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyStringWithUint(const QLocale* self, unsigned int i);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyStringWithDouble(const QLocale* self, double param1);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyStringWithFloat(const QLocale* self, float i);
extern __declspec(dllexport) struct miqt_string QLocale_FormattedDataSize(const QLocale* self, long long bytes);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QLocale_UiLanguages(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_LanguageToCode(Language language);
extern __declspec(dllexport) struct miqt_string QLocale_TerritoryToCode(Territory territory);
extern __declspec(dllexport) struct miqt_string QLocale_CountryToCode(Country country);
extern __declspec(dllexport) struct miqt_string QLocale_ScriptToCode(Script script);
extern __declspec(dllexport) struct miqt_string QLocale_LanguageToString(Language language);
extern __declspec(dllexport) struct miqt_string QLocale_TerritoryToString(Territory territory);
extern __declspec(dllexport) struct miqt_string QLocale_CountryToString(Country country);
extern __declspec(dllexport) struct miqt_string QLocale_ScriptToString(Script script);
extern __declspec(dllexport) void QLocale_SetDefault(QLocale* locale);
extern __declspec(dllexport) QLocale* QLocale_C();
extern __declspec(dllexport) QLocale* QLocale_System();
extern __declspec(dllexport) struct miqt_array /* of QLocale* */  QLocale_MatchingLocales(uint16_t language, uint16_t script, Country territory);
extern __declspec(dllexport) struct miqt_array /* of Country */  QLocale_CountriesForLanguage(Language lang);
extern __declspec(dllexport) void QLocale_SetNumberOptions(QLocale* self, NumberOptions options);
extern __declspec(dllexport) NumberOptions QLocale_NumberOptions(const QLocale* self);
extern __declspec(dllexport) struct miqt_string QLocale_QuoteString(const QLocale* self, struct miqt_string str);
extern __declspec(dllexport) struct miqt_string QLocale_CreateSeparatedList(const QLocale* self, struct miqt_array /* of struct miqt_string */  strl);
extern __declspec(dllexport) struct miqt_string QLocale_Name1(const QLocale* self, TagSeparator separator);
extern __declspec(dllexport) struct miqt_string QLocale_Bcp47Name1(const QLocale* self, TagSeparator separator);
extern __declspec(dllexport) int16_t QLocale_ToShort2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) uint16_t QLocale_ToUShort2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) int QLocale_ToInt2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) unsigned int QLocale_ToUInt2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) long QLocale_ToLong2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) unsigned long QLocale_ToULong2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) long long QLocale_ToLongLong2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) unsigned long long QLocale_ToULongLong2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) float QLocale_ToFloat2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) double QLocale_ToDouble2(const QLocale* self, struct miqt_string s, bool* ok);
extern __declspec(dllexport) struct miqt_string QLocale_ToString22(const QLocale* self, double f, char format);
extern __declspec(dllexport) struct miqt_string QLocale_ToString32(const QLocale* self, double f, char format, int precision);
extern __declspec(dllexport) struct miqt_string QLocale_ToString23(const QLocale* self, float f, char format);
extern __declspec(dllexport) struct miqt_string QLocale_ToString33(const QLocale* self, float f, char format, int precision);
extern __declspec(dllexport) struct miqt_string QLocale_ToString24(const QLocale* self, QDate* date, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_ToString25(const QLocale* self, QTime* time, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_ToString26(const QLocale* self, QDateTime* dateTime, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_DateFormat1(const QLocale* self, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_TimeFormat1(const QLocale* self, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_DateTimeFormat1(const QLocale* self, FormatType format);
extern __declspec(dllexport) QTime* QLocale_ToTime22(const QLocale* self, struct miqt_string stringVal, FormatType param2);
extern __declspec(dllexport) QDate* QLocale_ToDate22(const QLocale* self, struct miqt_string stringVal, FormatType param2);
extern __declspec(dllexport) QDate* QLocale_ToDate32(const QLocale* self, struct miqt_string stringVal, FormatType param2, int baseYear);
extern __declspec(dllexport) QDate* QLocale_ToDate33(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, int baseYear);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime22(const QLocale* self, struct miqt_string stringVal, FormatType format);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime32(const QLocale* self, struct miqt_string stringVal, FormatType format, int baseYear);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime33(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, int baseYear);
extern __declspec(dllexport) QDate* QLocale_ToDate42(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal, int baseYear);
extern __declspec(dllexport) QDate* QLocale_ToDate43(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal, int baseYear);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime42(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal, int baseYear);
extern __declspec(dllexport) QDateTime* QLocale_ToDateTime43(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal, int baseYear);
extern __declspec(dllexport) struct miqt_string QLocale_MonthName2(const QLocale* self, int param1, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_StandaloneMonthName2(const QLocale* self, int param1, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_DayName2(const QLocale* self, int param1, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_StandaloneDayName2(const QLocale* self, int param1, FormatType format);
extern __declspec(dllexport) struct miqt_string QLocale_CurrencySymbol1(const QLocale* self, CurrencySymbolFormat param1);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString2(const QLocale* self, long long param1, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString22(const QLocale* self, unsigned long long param1, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString23(const QLocale* self, int16_t i, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString24(const QLocale* self, uint16_t i, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString25(const QLocale* self, int i, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString26(const QLocale* self, unsigned int i, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString27(const QLocale* self, double param1, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString3(const QLocale* self, double param1, struct miqt_string symbol, int precision);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString28(const QLocale* self, float i, struct miqt_string symbol);
extern __declspec(dllexport) struct miqt_string QLocale_ToCurrencyString32(const QLocale* self, float i, struct miqt_string symbol, int precision);
extern __declspec(dllexport) struct miqt_string QLocale_FormattedDataSize2(const QLocale* self, long long bytes, int precision);
extern __declspec(dllexport) struct miqt_string QLocale_FormattedDataSize3(const QLocale* self, long long bytes, int precision, DataSizeFormats format);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QLocale_UiLanguages1(const QLocale* self, TagSeparator separator);
extern __declspec(dllexport) struct miqt_string QLocale_LanguageToCode2(Language language, LanguageCodeTypes codeTypes);
extern __declspec(dllexport) struct miqt_string QLocale_QuoteString2(const QLocale* self, struct miqt_string str, QuotationStyle style);
extern __declspec(dllexport) void QLocale_Delete(QLocale* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
