// +build ignore

#include <QByteArray>
#include <QByteArrayView>
#include <QChar>
#include <QLatin1Char>
#include <QLatin1String>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qlatin1stringview.h>
#include "gen_qlatin1stringview.h"

QLatin1String* QLatin1String_new() {
	return new QLatin1String();
}

QLatin1String* QLatin1String_new2(const char* s) {
	return new QLatin1String(s);
}

QLatin1String* QLatin1String_new3(const char* f, const char* l) {
	return new QLatin1String(f, l);
}

QLatin1String* QLatin1String_new4(const char* s, ptrdiff_t sz) {
	return new QLatin1String(s, (qsizetype)(sz));
}

QLatin1String* QLatin1String_new5(struct miqt_string s) {
	QByteArray s_QByteArray(s.data, s.len);
	return new QLatin1String(s_QByteArray);
}

QLatin1String* QLatin1String_new6(QByteArrayView* s) {
	return new QLatin1String(*s);
}

struct miqt_string QLatin1String_ToString(const QLatin1String* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLatin1String_ToUtf8(const QLatin1String* self) {
	QByteArray _qb = self->toUtf8();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

const char* QLatin1String_Latin1(const QLatin1String* self) {
	return (const char*) self->latin1();
}

ptrdiff_t QLatin1String_Size(const QLatin1String* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}

const char* QLatin1String_Data(const QLatin1String* self) {
	return (const char*) self->data();
}

const char* QLatin1String_ConstData(const QLatin1String* self) {
	return (const char*) self->constData();
}

const char* QLatin1String_ConstBegin(const QLatin1String* self) {
	return (const char*) self->constBegin();
}

const char* QLatin1String_ConstEnd(const QLatin1String* self) {
	return (const char*) self->constEnd();
}

QLatin1Char* QLatin1String_First(const QLatin1String* self) {
	return new QLatin1Char(self->first());
}

QLatin1Char* QLatin1String_Last(const QLatin1String* self) {
	return new QLatin1Char(self->last());
}

ptrdiff_t QLatin1String_Length(const QLatin1String* self) {
	qsizetype _ret = self->length();
	return static_cast<ptrdiff_t>(_ret);
}

bool QLatin1String_IsNull(const QLatin1String* self) {
	return self->isNull();
}

bool QLatin1String_IsEmpty(const QLatin1String* self) {
	return self->isEmpty();
}

bool QLatin1String_Empty(const QLatin1String* self) {
	return self->empty();
}

QLatin1Char* QLatin1String_At(const QLatin1String* self, ptrdiff_t i) {
	return new QLatin1Char(self->at((qsizetype)(i)));
}

QLatin1Char* QLatin1String_OperatorSubscript(const QLatin1String* self, ptrdiff_t i) {
	return new QLatin1Char(self->operator[]((qsizetype)(i)));
}

QLatin1Char* QLatin1String_Front(const QLatin1String* self) {
	return new QLatin1Char(self->front());
}

QLatin1Char* QLatin1String_Back(const QLatin1String* self) {
	return new QLatin1Char(self->back());
}

int QLatin1String_CompareWithQChar(const QLatin1String* self, QChar* c) {
	return self->compare(*c);
}

int QLatin1String_Compare3(const QLatin1String* self, QChar* c, int cs) {
	return self->compare(*c, static_cast<Qt::CaseSensitivity>(cs));
}

bool QLatin1String_StartsWithWithQChar(const QLatin1String* self, QChar* c) {
	return self->startsWith(*c);
}

bool QLatin1String_StartsWith2(const QLatin1String* self, QChar* c, int cs) {
	return self->startsWith(*c, static_cast<Qt::CaseSensitivity>(cs));
}

bool QLatin1String_EndsWithWithQChar(const QLatin1String* self, QChar* c) {
	return self->endsWith(*c);
}

bool QLatin1String_EndsWith2(const QLatin1String* self, QChar* c, int cs) {
	return self->endsWith(*c, static_cast<Qt::CaseSensitivity>(cs));
}

ptrdiff_t QLatin1String_IndexOfWithQChar(const QLatin1String* self, QChar* c) {
	qsizetype _ret = self->indexOf(*c);
	return static_cast<ptrdiff_t>(_ret);
}

bool QLatin1String_ContainsWithQChar(const QLatin1String* self, QChar* c) {
	return self->contains(*c);
}

ptrdiff_t QLatin1String_LastIndexOfWithQChar(const QLatin1String* self, QChar* c) {
	qsizetype _ret = self->lastIndexOf(*c);
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QLatin1String_LastIndexOf4(const QLatin1String* self, QChar* c, ptrdiff_t from) {
	qsizetype _ret = self->lastIndexOf(*c, (qsizetype)(from));
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QLatin1String_CountWithCh(const QLatin1String* self, QChar* ch) {
	qsizetype _ret = self->count(*ch);
	return static_cast<ptrdiff_t>(_ret);
}

int16_t QLatin1String_ToShort(const QLatin1String* self) {
	return self->toShort();
}

uint16_t QLatin1String_ToUShort(const QLatin1String* self) {
	ushort _ret = self->toUShort();
	return static_cast<uint16_t>(_ret);
}

int QLatin1String_ToInt(const QLatin1String* self) {
	return self->toInt();
}

unsigned int QLatin1String_ToUInt(const QLatin1String* self) {
	uint _ret = self->toUInt();
	return static_cast<unsigned int>(_ret);
}

long QLatin1String_ToLong(const QLatin1String* self) {
	return self->toLong();
}

unsigned long QLatin1String_ToULong(const QLatin1String* self) {
	ulong _ret = self->toULong();
	return static_cast<unsigned long>(_ret);
}

long long QLatin1String_ToLongLong(const QLatin1String* self) {
	qlonglong _ret = self->toLongLong();
	return static_cast<long long>(_ret);
}

unsigned long long QLatin1String_ToULongLong(const QLatin1String* self) {
	qulonglong _ret = self->toULongLong();
	return static_cast<unsigned long long>(_ret);
}

float QLatin1String_ToFloat(const QLatin1String* self) {
	return self->toFloat();
}

double QLatin1String_ToDouble(const QLatin1String* self) {
	return self->toDouble();
}

const_iterator QLatin1String_Begin(const QLatin1String* self) {
	return self->begin();
}

const_iterator QLatin1String_Cbegin(const QLatin1String* self) {
	return self->cbegin();
}

const_iterator QLatin1String_End(const QLatin1String* self) {
	return self->end();
}

const_iterator QLatin1String_Cend(const QLatin1String* self) {
	return self->cend();
}

const_reverse_iterator QLatin1String_Rbegin(const QLatin1String* self) {
	return self->rbegin();
}

const_reverse_iterator QLatin1String_Crbegin(const QLatin1String* self) {
	return self->crbegin();
}

const_reverse_iterator QLatin1String_Rend(const QLatin1String* self) {
	return self->rend();
}

const_reverse_iterator QLatin1String_Crend(const QLatin1String* self) {
	return self->crend();
}

ptrdiff_t QLatin1String_MaxSize(const QLatin1String* self) {
	qsizetype _ret = self->max_size();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QLatin1String_MaxSize2() {
	qsizetype _ret = QLatin1String::maxSize();
	return static_cast<ptrdiff_t>(_ret);
}

void QLatin1String_Chop(QLatin1String* self, ptrdiff_t n) {
	self->chop((qsizetype)(n));
}

void QLatin1String_Truncate(QLatin1String* self, ptrdiff_t n) {
	self->truncate((qsizetype)(n));
}

ptrdiff_t QLatin1String_IndexOf23(const QLatin1String* self, QChar* c, ptrdiff_t from) {
	qsizetype _ret = self->indexOf(*c, (qsizetype)(from));
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QLatin1String_IndexOf33(const QLatin1String* self, QChar* c, ptrdiff_t from, int cs) {
	qsizetype _ret = self->indexOf(*c, (qsizetype)(from), static_cast<Qt::CaseSensitivity>(cs));
	return static_cast<ptrdiff_t>(_ret);
}

bool QLatin1String_Contains23(const QLatin1String* self, QChar* c, int cs) {
	return self->contains(*c, static_cast<Qt::CaseSensitivity>(cs));
}

ptrdiff_t QLatin1String_LastIndexOf24(const QLatin1String* self, QChar* c, int cs) {
	qsizetype _ret = self->lastIndexOf(*c, static_cast<Qt::CaseSensitivity>(cs));
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QLatin1String_LastIndexOf34(const QLatin1String* self, QChar* c, ptrdiff_t from, int cs) {
	qsizetype _ret = self->lastIndexOf(*c, (qsizetype)(from), static_cast<Qt::CaseSensitivity>(cs));
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QLatin1String_Count23(const QLatin1String* self, QChar* ch, int cs) {
	qsizetype _ret = self->count(*ch, static_cast<Qt::CaseSensitivity>(cs));
	return static_cast<ptrdiff_t>(_ret);
}

int16_t QLatin1String_ToShort1(const QLatin1String* self, bool* ok) {
	return self->toShort(ok);
}

int16_t QLatin1String_ToShort2(const QLatin1String* self, bool* ok, int base) {
	return self->toShort(ok, static_cast<int>(base));
}

uint16_t QLatin1String_ToUShort1(const QLatin1String* self, bool* ok) {
	ushort _ret = self->toUShort(ok);
	return static_cast<uint16_t>(_ret);
}

uint16_t QLatin1String_ToUShort2(const QLatin1String* self, bool* ok, int base) {
	ushort _ret = self->toUShort(ok, static_cast<int>(base));
	return static_cast<uint16_t>(_ret);
}

int QLatin1String_ToInt1(const QLatin1String* self, bool* ok) {
	return self->toInt(ok);
}

int QLatin1String_ToInt2(const QLatin1String* self, bool* ok, int base) {
	return self->toInt(ok, static_cast<int>(base));
}

unsigned int QLatin1String_ToUInt1(const QLatin1String* self, bool* ok) {
	uint _ret = self->toUInt(ok);
	return static_cast<unsigned int>(_ret);
}

unsigned int QLatin1String_ToUInt2(const QLatin1String* self, bool* ok, int base) {
	uint _ret = self->toUInt(ok, static_cast<int>(base));
	return static_cast<unsigned int>(_ret);
}

long QLatin1String_ToLong1(const QLatin1String* self, bool* ok) {
	return self->toLong(ok);
}

long QLatin1String_ToLong2(const QLatin1String* self, bool* ok, int base) {
	return self->toLong(ok, static_cast<int>(base));
}

unsigned long QLatin1String_ToULong1(const QLatin1String* self, bool* ok) {
	ulong _ret = self->toULong(ok);
	return static_cast<unsigned long>(_ret);
}

unsigned long QLatin1String_ToULong2(const QLatin1String* self, bool* ok, int base) {
	ulong _ret = self->toULong(ok, static_cast<int>(base));
	return static_cast<unsigned long>(_ret);
}

long long QLatin1String_ToLongLong1(const QLatin1String* self, bool* ok) {
	qlonglong _ret = self->toLongLong(ok);
	return static_cast<long long>(_ret);
}

long long QLatin1String_ToLongLong2(const QLatin1String* self, bool* ok, int base) {
	qlonglong _ret = self->toLongLong(ok, static_cast<int>(base));
	return static_cast<long long>(_ret);
}

unsigned long long QLatin1String_ToULongLong1(const QLatin1String* self, bool* ok) {
	qulonglong _ret = self->toULongLong(ok);
	return static_cast<unsigned long long>(_ret);
}

unsigned long long QLatin1String_ToULongLong2(const QLatin1String* self, bool* ok, int base) {
	qulonglong _ret = self->toULongLong(ok, static_cast<int>(base));
	return static_cast<unsigned long long>(_ret);
}

float QLatin1String_ToFloat1(const QLatin1String* self, bool* ok) {
	return self->toFloat(ok);
}

double QLatin1String_ToDouble1(const QLatin1String* self, bool* ok) {
	return self->toDouble(ok);
}

void QLatin1String_Delete(QLatin1String* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QLatin1String*>( self );
	} else {
		delete self;
	}
}

