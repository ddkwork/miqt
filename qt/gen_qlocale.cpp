// +build ignore

#include <QCalendar>
#include <QDate>
#include <QDateTime>
#include <QList>
#include <QLocale>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTextStream>
#include <QTime>
#include <qlocale.h>
#include "gen_qlocale.h"

void QTextStream_Delete(QTextStream* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTextStream*>( self );
	} else {
		delete self;
	}
}

QLocale* QLocale_new() {
	return new QLocale();
}

QLocale* QLocale_new2(struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new QLocale(name_QString);
}

QLocale* QLocale_new3(Language language, Territory territory) {
	return new QLocale(language, territory);
}

QLocale* QLocale_new4(Language language) {
	return new QLocale(language);
}

QLocale* QLocale_new5(QLocale* other) {
	return new QLocale(*other);
}

QLocale* QLocale_new6(Language language, Script script) {
	return new QLocale(language, script);
}

QLocale* QLocale_new7(Language language, Script script, Territory territory) {
	return new QLocale(language, script, territory);
}

void QLocale_OperatorAssign(QLocale* self, QLocale* other) {
	self->operator=(*other);
}

void QLocale_Swap(QLocale* self, QLocale* other) {
	self->swap(*other);
}

Language QLocale_Language(const QLocale* self) {
	return self->language();
}

Script QLocale_Script(const QLocale* self) {
	return self->script();
}

Territory QLocale_Territory(const QLocale* self) {
	return self->territory();
}

Country QLocale_Country(const QLocale* self) {
	return self->country();
}

struct miqt_string QLocale_Name(const QLocale* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_Bcp47Name(const QLocale* self) {
	QString _ret = self->bcp47Name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_NativeLanguageName(const QLocale* self) {
	QString _ret = self->nativeLanguageName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_NativeTerritoryName(const QLocale* self) {
	QString _ret = self->nativeTerritoryName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_NativeCountryName(const QLocale* self) {
	QString _ret = self->nativeCountryName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int16_t QLocale_ToShort(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toShort(s_QString);
}

uint16_t QLocale_ToUShort(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	ushort _ret = self->toUShort(s_QString);
	return static_cast<uint16_t>(_ret);
}

int QLocale_ToInt(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toInt(s_QString);
}

unsigned int QLocale_ToUInt(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	uint _ret = self->toUInt(s_QString);
	return static_cast<unsigned int>(_ret);
}

long QLocale_ToLong(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toLong(s_QString);
}

unsigned long QLocale_ToULong(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	ulong _ret = self->toULong(s_QString);
	return static_cast<unsigned long>(_ret);
}

long long QLocale_ToLongLong(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	qlonglong _ret = self->toLongLong(s_QString);
	return static_cast<long long>(_ret);
}

unsigned long long QLocale_ToULongLong(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	qulonglong _ret = self->toULongLong(s_QString);
	return static_cast<unsigned long long>(_ret);
}

float QLocale_ToFloat(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toFloat(s_QString);
}

double QLocale_ToDouble(const QLocale* self, struct miqt_string s) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toDouble(s_QString);
}

struct miqt_string QLocale_ToString(const QLocale* self, long long i) {
	QString _ret = self->toString(static_cast<qlonglong>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithQulonglong(const QLocale* self, unsigned long long i) {
	QString _ret = self->toString(static_cast<qulonglong>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithLong(const QLocale* self, long i) {
	QString _ret = self->toString(static_cast<long>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithUlong(const QLocale* self, unsigned long i) {
	QString _ret = self->toString(static_cast<ulong>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithShort(const QLocale* self, int16_t i) {
	QString _ret = self->toString(static_cast<short>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithUshort(const QLocale* self, uint16_t i) {
	QString _ret = self->toString(static_cast<ushort>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithInt(const QLocale* self, int i) {
	QString _ret = self->toString(static_cast<int>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithUint(const QLocale* self, unsigned int i) {
	QString _ret = self->toString(static_cast<uint>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithDouble(const QLocale* self, double f) {
	QString _ret = self->toString(static_cast<double>(f));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithFloat(const QLocale* self, float f) {
	QString _ret = self->toString(static_cast<float>(f));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString2(const QLocale* self, QDate* date, struct miqt_string format) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	QString _ret = self->toString(*date, format_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString3(const QLocale* self, QTime* time, struct miqt_string format) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	QString _ret = self->toString(*time, format_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString4(const QLocale* self, QDateTime* dateTime, struct miqt_string format) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	QString _ret = self->toString(*dateTime, format_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithDate(const QLocale* self, QDate* date) {
	QString _ret = self->toString(*date);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithTime(const QLocale* self, QTime* time) {
	QString _ret = self->toString(*time);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToStringWithDateTime(const QLocale* self, QDateTime* dateTime) {
	QString _ret = self->toString(*dateTime);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString9(const QLocale* self, QDate* date, FormatType format, QCalendar* cal) {
	QString _ret = self->toString(*date, format, *cal);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString10(const QLocale* self, QDateTime* dateTime, FormatType format, QCalendar* cal) {
	QString _ret = self->toString(*dateTime, format, *cal);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_DateFormat(const QLocale* self) {
	QString _ret = self->dateFormat();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_TimeFormat(const QLocale* self) {
	QString _ret = self->timeFormat();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_DateTimeFormat(const QLocale* self) {
	QString _ret = self->dateTimeFormat();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QTime* QLocale_ToTime(const QLocale* self, struct miqt_string stringVal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QTime(self->toTime(stringVal_QString));
}

QTime* QLocale_ToTime2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QTime(self->toTime(stringVal_QString, format_QString));
}

QDate* QLocale_ToDate(const QLocale* self, struct miqt_string stringVal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDate(self->toDate(stringVal_QString));
}

QDate* QLocale_ToDate2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDate(self->toDate(stringVal_QString, format_QString));
}

QDateTime* QLocale_ToDateTime(const QLocale* self, struct miqt_string stringVal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDateTime(self->toDateTime(stringVal_QString));
}

QDateTime* QLocale_ToDateTime2(const QLocale* self, struct miqt_string stringVal, struct miqt_string format) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format_QString));
}

QDate* QLocale_ToDate3(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDate(self->toDate(stringVal_QString, format, *cal));
}

QDate* QLocale_ToDate4(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDate(self->toDate(stringVal_QString, format_QString, *cal));
}

QDateTime* QLocale_ToDateTime3(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format, *cal));
}

QDateTime* QLocale_ToDateTime4(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format_QString, *cal));
}

struct miqt_string QLocale_DecimalPoint(const QLocale* self) {
	QString _ret = self->decimalPoint();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_GroupSeparator(const QLocale* self) {
	QString _ret = self->groupSeparator();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_Percent(const QLocale* self) {
	QString _ret = self->percent();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ZeroDigit(const QLocale* self) {
	QString _ret = self->zeroDigit();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_NegativeSign(const QLocale* self) {
	QString _ret = self->negativeSign();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_PositiveSign(const QLocale* self) {
	QString _ret = self->positiveSign();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_Exponential(const QLocale* self) {
	QString _ret = self->exponential();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_MonthName(const QLocale* self, int param1) {
	QString _ret = self->monthName(static_cast<int>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_StandaloneMonthName(const QLocale* self, int param1) {
	QString _ret = self->standaloneMonthName(static_cast<int>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_DayName(const QLocale* self, int param1) {
	QString _ret = self->dayName(static_cast<int>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_StandaloneDayName(const QLocale* self, int param1) {
	QString _ret = self->standaloneDayName(static_cast<int>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QLocale_FirstDayOfWeek(const QLocale* self) {
	Qt::DayOfWeek _ret = self->firstDayOfWeek();
	return static_cast<int>(_ret);
}

struct miqt_array /* of int */  QLocale_Weekdays(const QLocale* self) {
	QList<Qt::DayOfWeek> _ret = self->weekdays();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		Qt::DayOfWeek _lv_ret = _ret[i];
		_arr[i] = static_cast<int>(_lv_ret);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QLocale_AmText(const QLocale* self) {
	QString _ret = self->amText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_PmText(const QLocale* self) {
	QString _ret = self->pmText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

MeasurementSystem QLocale_MeasurementSystem(const QLocale* self) {
	return self->measurementSystem();
}

QLocale* QLocale_Collation(const QLocale* self) {
	return new QLocale(self->collation());
}

int QLocale_TextDirection(const QLocale* self) {
	Qt::LayoutDirection _ret = self->textDirection();
	return static_cast<int>(_ret);
}

struct miqt_string QLocale_ToUpper(const QLocale* self, struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QString _ret = self->toUpper(str_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToLower(const QLocale* self, struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QString _ret = self->toLower(str_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_CurrencySymbol(const QLocale* self) {
	QString _ret = self->currencySymbol();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString(const QLocale* self, long long param1) {
	QString _ret = self->toCurrencyString(static_cast<qlonglong>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyStringWithQulonglong(const QLocale* self, unsigned long long param1) {
	QString _ret = self->toCurrencyString(static_cast<qulonglong>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyStringWithShort(const QLocale* self, int16_t i) {
	QString _ret = self->toCurrencyString(static_cast<short>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyStringWithUshort(const QLocale* self, uint16_t i) {
	QString _ret = self->toCurrencyString(static_cast<ushort>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyStringWithInt(const QLocale* self, int i) {
	QString _ret = self->toCurrencyString(static_cast<int>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyStringWithUint(const QLocale* self, unsigned int i) {
	QString _ret = self->toCurrencyString(static_cast<uint>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyStringWithDouble(const QLocale* self, double param1) {
	QString _ret = self->toCurrencyString(static_cast<double>(param1));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyStringWithFloat(const QLocale* self, float i) {
	QString _ret = self->toCurrencyString(static_cast<float>(i));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_FormattedDataSize(const QLocale* self, long long bytes) {
	QString _ret = self->formattedDataSize(static_cast<qint64>(bytes));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of struct miqt_string */  QLocale_UiLanguages(const QLocale* self) {
	QStringList _ret = self->uiLanguages();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QLocale_LanguageToCode(Language language) {
	QString _ret = QLocale::languageToCode(language);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_TerritoryToCode(Territory territory) {
	QString _ret = QLocale::territoryToCode(territory);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_CountryToCode(Country country) {
	QString _ret = QLocale::countryToCode(country);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ScriptToCode(Script script) {
	QString _ret = QLocale::scriptToCode(script);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_LanguageToString(Language language) {
	QString _ret = QLocale::languageToString(language);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_TerritoryToString(Territory territory) {
	QString _ret = QLocale::territoryToString(territory);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_CountryToString(Country country) {
	QString _ret = QLocale::countryToString(country);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ScriptToString(Script script) {
	QString _ret = QLocale::scriptToString(script);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLocale_SetDefault(QLocale* locale) {
	QLocale::setDefault(*locale);
}

QLocale* QLocale_C() {
	return new QLocale(QLocale::c());
}

QLocale* QLocale_System() {
	return new QLocale(QLocale::system());
}

struct miqt_array /* of QLocale* */  QLocale_MatchingLocales(uint16_t language, uint16_t script, Country territory) {
	QList<QLocale> _ret = QLocale::matchingLocales(static_cast<QLocale::Language>(language), static_cast<QLocale::Script>(script), territory);
	// Convert QList<> from C++ memory to manually-managed C memory
	QLocale** _arr = static_cast<QLocale**>(malloc(sizeof(QLocale*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QLocale(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of Country */  QLocale_CountriesForLanguage(Language lang) {
	QList<Country> _ret = QLocale::countriesForLanguage(lang);
	// Convert QList<> from C++ memory to manually-managed C memory
	Country* _arr = static_cast<Country*>(malloc(sizeof(Country) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QLocale_SetNumberOptions(QLocale* self, NumberOptions options) {
	self->setNumberOptions(options);
}

NumberOptions QLocale_NumberOptions(const QLocale* self) {
	return self->numberOptions();
}

struct miqt_string QLocale_QuoteString(const QLocale* self, struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QString _ret = self->quoteString(str_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_CreateSeparatedList(const QLocale* self, struct miqt_array /* of struct miqt_string */  strl) {
	QStringList strl_QList;
	strl_QList.reserve(strl.len);
	struct miqt_string* strl_arr = static_cast<struct miqt_string*>(strl.data);
	for(size_t i = 0; i < strl.len; ++i) {
		QString strl_arr_i_QString = QString::fromUtf8(strl_arr[i].data, strl_arr[i].len);
		strl_QList.push_back(strl_arr_i_QString);
	}
	QString _ret = self->createSeparatedList(strl_QList);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_Name1(const QLocale* self, TagSeparator separator) {
	QString _ret = self->name(separator);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_Bcp47Name1(const QLocale* self, TagSeparator separator) {
	QString _ret = self->bcp47Name(separator);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int16_t QLocale_ToShort2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toShort(s_QString, ok);
}

uint16_t QLocale_ToUShort2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	ushort _ret = self->toUShort(s_QString, ok);
	return static_cast<uint16_t>(_ret);
}

int QLocale_ToInt2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toInt(s_QString, ok);
}

unsigned int QLocale_ToUInt2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	uint _ret = self->toUInt(s_QString, ok);
	return static_cast<unsigned int>(_ret);
}

long QLocale_ToLong2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toLong(s_QString, ok);
}

unsigned long QLocale_ToULong2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	ulong _ret = self->toULong(s_QString, ok);
	return static_cast<unsigned long>(_ret);
}

long long QLocale_ToLongLong2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	qlonglong _ret = self->toLongLong(s_QString, ok);
	return static_cast<long long>(_ret);
}

unsigned long long QLocale_ToULongLong2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	qulonglong _ret = self->toULongLong(s_QString, ok);
	return static_cast<unsigned long long>(_ret);
}

float QLocale_ToFloat2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toFloat(s_QString, ok);
}

double QLocale_ToDouble2(const QLocale* self, struct miqt_string s, bool* ok) {
	QString s_QString = QString::fromUtf8(s.data, s.len);
	return self->toDouble(s_QString, ok);
}

struct miqt_string QLocale_ToString22(const QLocale* self, double f, char format) {
	QString _ret = self->toString(static_cast<double>(f), static_cast<char>(format));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString32(const QLocale* self, double f, char format, int precision) {
	QString _ret = self->toString(static_cast<double>(f), static_cast<char>(format), static_cast<int>(precision));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString23(const QLocale* self, float f, char format) {
	QString _ret = self->toString(static_cast<float>(f), static_cast<char>(format));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString33(const QLocale* self, float f, char format, int precision) {
	QString _ret = self->toString(static_cast<float>(f), static_cast<char>(format), static_cast<int>(precision));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString24(const QLocale* self, QDate* date, FormatType format) {
	QString _ret = self->toString(*date, format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString25(const QLocale* self, QTime* time, FormatType format) {
	QString _ret = self->toString(*time, format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToString26(const QLocale* self, QDateTime* dateTime, FormatType format) {
	QString _ret = self->toString(*dateTime, format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_DateFormat1(const QLocale* self, FormatType format) {
	QString _ret = self->dateFormat(format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_TimeFormat1(const QLocale* self, FormatType format) {
	QString _ret = self->timeFormat(format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_DateTimeFormat1(const QLocale* self, FormatType format) {
	QString _ret = self->dateTimeFormat(format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QTime* QLocale_ToTime22(const QLocale* self, struct miqt_string stringVal, FormatType param2) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QTime(self->toTime(stringVal_QString, param2));
}

QDate* QLocale_ToDate22(const QLocale* self, struct miqt_string stringVal, FormatType param2) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDate(self->toDate(stringVal_QString, param2));
}

QDate* QLocale_ToDate32(const QLocale* self, struct miqt_string stringVal, FormatType param2, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDate(self->toDate(stringVal_QString, param2, static_cast<int>(baseYear)));
}

QDate* QLocale_ToDate33(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDate(self->toDate(stringVal_QString, format_QString, static_cast<int>(baseYear)));
}

QDateTime* QLocale_ToDateTime22(const QLocale* self, struct miqt_string stringVal, FormatType format) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format));
}

QDateTime* QLocale_ToDateTime32(const QLocale* self, struct miqt_string stringVal, FormatType format, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format, static_cast<int>(baseYear)));
}

QDateTime* QLocale_ToDateTime33(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format_QString, static_cast<int>(baseYear)));
}

QDate* QLocale_ToDate42(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDate(self->toDate(stringVal_QString, format, *cal, static_cast<int>(baseYear)));
}

QDate* QLocale_ToDate43(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDate(self->toDate(stringVal_QString, format_QString, *cal, static_cast<int>(baseYear)));
}

QDateTime* QLocale_ToDateTime42(const QLocale* self, struct miqt_string stringVal, FormatType format, QCalendar* cal, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format, *cal, static_cast<int>(baseYear)));
}

QDateTime* QLocale_ToDateTime43(const QLocale* self, struct miqt_string stringVal, struct miqt_string format, QCalendar* cal, int baseYear) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	QString format_QString = QString::fromUtf8(format.data, format.len);
	return new QDateTime(self->toDateTime(stringVal_QString, format_QString, *cal, static_cast<int>(baseYear)));
}

struct miqt_string QLocale_MonthName2(const QLocale* self, int param1, FormatType format) {
	QString _ret = self->monthName(static_cast<int>(param1), format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_StandaloneMonthName2(const QLocale* self, int param1, FormatType format) {
	QString _ret = self->standaloneMonthName(static_cast<int>(param1), format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_DayName2(const QLocale* self, int param1, FormatType format) {
	QString _ret = self->dayName(static_cast<int>(param1), format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_StandaloneDayName2(const QLocale* self, int param1, FormatType format) {
	QString _ret = self->standaloneDayName(static_cast<int>(param1), format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_CurrencySymbol1(const QLocale* self, CurrencySymbolFormat param1) {
	QString _ret = self->currencySymbol(param1);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString2(const QLocale* self, long long param1, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<qlonglong>(param1), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString22(const QLocale* self, unsigned long long param1, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<qulonglong>(param1), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString23(const QLocale* self, int16_t i, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<short>(i), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString24(const QLocale* self, uint16_t i, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<ushort>(i), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString25(const QLocale* self, int i, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<int>(i), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString26(const QLocale* self, unsigned int i, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<uint>(i), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString27(const QLocale* self, double param1, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<double>(param1), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString3(const QLocale* self, double param1, struct miqt_string symbol, int precision) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<double>(param1), symbol_QString, static_cast<int>(precision));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString28(const QLocale* self, float i, struct miqt_string symbol) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<float>(i), symbol_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_ToCurrencyString32(const QLocale* self, float i, struct miqt_string symbol, int precision) {
	QString symbol_QString = QString::fromUtf8(symbol.data, symbol.len);
	QString _ret = self->toCurrencyString(static_cast<float>(i), symbol_QString, static_cast<int>(precision));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_FormattedDataSize2(const QLocale* self, long long bytes, int precision) {
	QString _ret = self->formattedDataSize(static_cast<qint64>(bytes), static_cast<int>(precision));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_FormattedDataSize3(const QLocale* self, long long bytes, int precision, DataSizeFormats format) {
	QString _ret = self->formattedDataSize(static_cast<qint64>(bytes), static_cast<int>(precision), format);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of struct miqt_string */  QLocale_UiLanguages1(const QLocale* self, TagSeparator separator) {
	QStringList _ret = self->uiLanguages(separator);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QLocale_LanguageToCode2(Language language, LanguageCodeTypes codeTypes) {
	QString _ret = QLocale::languageToCode(language, codeTypes);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocale_QuoteString2(const QLocale* self, struct miqt_string str, QuotationStyle style) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QString _ret = self->quoteString(str_QString, style);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLocale_Delete(QLocale* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QLocale*>( self );
	} else {
		delete self;
	}
}

