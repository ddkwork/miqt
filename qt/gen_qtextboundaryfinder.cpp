// +build ignore

#include <QChar>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTextBoundaryFinder>
#include <qtextboundaryfinder.h>
#include "gen_qtextboundaryfinder.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

QTextBoundaryFinder* QTextBoundaryFinder_new() {
	return new QTextBoundaryFinder();
}

QTextBoundaryFinder* QTextBoundaryFinder_new2(QTextBoundaryFinder* other) {
	return new QTextBoundaryFinder(*other);
}

QTextBoundaryFinder* QTextBoundaryFinder_new3(BoundaryType typeVal, struct miqt_string stringVal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QTextBoundaryFinder(typeVal, stringVal_QString);
}

QTextBoundaryFinder* QTextBoundaryFinder_new4(BoundaryType typeVal, QChar* chars, ptrdiff_t length) {
	return new QTextBoundaryFinder(typeVal, chars, (qsizetype)(length));
}

QTextBoundaryFinder* QTextBoundaryFinder_new5(BoundaryType typeVal, QChar* chars, ptrdiff_t length, unsigned char* buffer) {
	return new QTextBoundaryFinder(typeVal, chars, (qsizetype)(length), static_cast<unsigned char*>(buffer));
}

QTextBoundaryFinder* QTextBoundaryFinder_new6(BoundaryType typeVal, QChar* chars, ptrdiff_t length, unsigned char* buffer, ptrdiff_t bufferSize) {
	return new QTextBoundaryFinder(typeVal, chars, (qsizetype)(length), static_cast<unsigned char*>(buffer), (qsizetype)(bufferSize));
}

void QTextBoundaryFinder_OperatorAssign(QTextBoundaryFinder* self, QTextBoundaryFinder* other) {
	self->operator=(*other);
}

bool QTextBoundaryFinder_IsValid(const QTextBoundaryFinder* self) {
	return self->isValid();
}

BoundaryType QTextBoundaryFinder_Type(const QTextBoundaryFinder* self) {
	return self->type();
}

struct miqt_string QTextBoundaryFinder_String(const QTextBoundaryFinder* self) {
	QString _ret = self->string();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QTextBoundaryFinder_ToStart(QTextBoundaryFinder* self) {
	self->toStart();
}

void QTextBoundaryFinder_ToEnd(QTextBoundaryFinder* self) {
	self->toEnd();
}

ptrdiff_t QTextBoundaryFinder_Position(const QTextBoundaryFinder* self) {
	qsizetype _ret = self->position();
	return static_cast<ptrdiff_t>(_ret);
}

void QTextBoundaryFinder_SetPosition(QTextBoundaryFinder* self, ptrdiff_t position) {
	self->setPosition((qsizetype)(position));
}

ptrdiff_t QTextBoundaryFinder_ToNextBoundary(QTextBoundaryFinder* self) {
	qsizetype _ret = self->toNextBoundary();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QTextBoundaryFinder_ToPreviousBoundary(QTextBoundaryFinder* self) {
	qsizetype _ret = self->toPreviousBoundary();
	return static_cast<ptrdiff_t>(_ret);
}

bool QTextBoundaryFinder_IsAtBoundary(const QTextBoundaryFinder* self) {
	return self->isAtBoundary();
}

BoundaryReasons QTextBoundaryFinder_BoundaryReasons(const QTextBoundaryFinder* self) {
	return self->boundaryReasons();
}

void QTextBoundaryFinder_Delete(QTextBoundaryFinder* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTextBoundaryFinder*>( self );
	} else {
		delete self;
	}
}

