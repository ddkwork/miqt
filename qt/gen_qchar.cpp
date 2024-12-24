// +build ignore

#include <QChar>
#include <QLatin1Char>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qchar.h>
#include "gen_qchar.h"

QLatin1Char* QLatin1Char_new(char c) {
	return new QLatin1Char(static_cast<char>(c));
}

QLatin1Char* QLatin1Char_new2(QLatin1Char* param1) {
	return new QLatin1Char(*param1);
}

char QLatin1Char_ToLatin1(const QLatin1Char* self) {
	return self->toLatin1();
}

void QLatin1Char_Delete(QLatin1Char* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QLatin1Char*>( self );
	} else {
		delete self;
	}
}

QChar* QChar_new() {
	return new QChar();
}

QChar* QChar_new2(unsigned char c, unsigned char r) {
	return new QChar(static_cast<uchar>(c), static_cast<uchar>(r));
}

QChar* QChar_new3(QChar* param1) {
	return new QChar(*param1);
}

Category QChar_Category(const QChar* self) {
	return self->category();
}

Direction QChar_Direction(const QChar* self) {
	return self->direction();
}

JoiningType QChar_JoiningType(const QChar* self) {
	return self->joiningType();
}

unsigned char QChar_CombiningClass(const QChar* self) {
	return self->combiningClass();
}

QChar* QChar_MirroredChar(const QChar* self) {
	return new QChar(self->mirroredChar());
}

bool QChar_HasMirrored(const QChar* self) {
	return self->hasMirrored();
}

struct miqt_string QChar_Decomposition(const QChar* self) {
	QString _ret = self->decomposition();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

Decomposition QChar_DecompositionTag(const QChar* self) {
	return self->decompositionTag();
}

int QChar_DigitValue(const QChar* self) {
	return self->digitValue();
}

QChar* QChar_ToLower(const QChar* self) {
	return new QChar(self->toLower());
}

QChar* QChar_ToUpper(const QChar* self) {
	return new QChar(self->toUpper());
}

QChar* QChar_ToTitleCase(const QChar* self) {
	return new QChar(self->toTitleCase());
}

QChar* QChar_ToCaseFolded(const QChar* self) {
	return new QChar(self->toCaseFolded());
}

Script QChar_Script(const QChar* self) {
	return self->script();
}

UnicodeVersion QChar_UnicodeVersion(const QChar* self) {
	return self->unicodeVersion();
}

char QChar_ToLatin1(const QChar* self) {
	return self->toLatin1();
}

QChar* QChar_FromLatin1(char c) {
	return new QChar(QChar::fromLatin1(static_cast<char>(c)));
}

bool QChar_IsNull(const QChar* self) {
	return self->isNull();
}

bool QChar_IsPrint(const QChar* self) {
	return self->isPrint();
}

bool QChar_IsSpace(const QChar* self) {
	return self->isSpace();
}

bool QChar_IsMark(const QChar* self) {
	return self->isMark();
}

bool QChar_IsPunct(const QChar* self) {
	return self->isPunct();
}

bool QChar_IsSymbol(const QChar* self) {
	return self->isSymbol();
}

bool QChar_IsLetter(const QChar* self) {
	return self->isLetter();
}

bool QChar_IsNumber(const QChar* self) {
	return self->isNumber();
}

bool QChar_IsLetterOrNumber(const QChar* self) {
	return self->isLetterOrNumber();
}

bool QChar_IsDigit(const QChar* self) {
	return self->isDigit();
}

bool QChar_IsLower(const QChar* self) {
	return self->isLower();
}

bool QChar_IsUpper(const QChar* self) {
	return self->isUpper();
}

bool QChar_IsTitleCase(const QChar* self) {
	return self->isTitleCase();
}

bool QChar_IsNonCharacter(const QChar* self) {
	return self->isNonCharacter();
}

bool QChar_IsHighSurrogate(const QChar* self) {
	return self->isHighSurrogate();
}

bool QChar_IsLowSurrogate(const QChar* self) {
	return self->isLowSurrogate();
}

bool QChar_IsSurrogate(const QChar* self) {
	return self->isSurrogate();
}

unsigned char QChar_Cell(const QChar* self) {
	uchar _ret = self->cell();
	return static_cast<unsigned char>(_ret);
}

unsigned char QChar_Row(const QChar* self) {
	uchar _ret = self->row();
	return static_cast<unsigned char>(_ret);
}

void QChar_SetCell(QChar* self, unsigned char acell) {
	self->setCell(static_cast<uchar>(acell));
}

void QChar_SetRow(QChar* self, unsigned char arow) {
	self->setRow(static_cast<uchar>(arow));
}

UnicodeVersion QChar_CurrentUnicodeVersion() {
	return QChar::currentUnicodeVersion();
}

void QChar_Delete(QChar* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QChar*>( self );
	} else {
		delete self;
	}
}

