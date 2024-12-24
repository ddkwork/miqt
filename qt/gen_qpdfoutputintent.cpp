// +build ignore

#include <QColorSpace>
#include <QPdfOutputIntent>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <qpdfoutputintent.h>
#include "gen_qpdfoutputintent.h"

#ifndef _Bool
#define _Bool bool
#endif

QPdfOutputIntent* QPdfOutputIntent_new() {
	return new QPdfOutputIntent();
}

QPdfOutputIntent* QPdfOutputIntent_new2(QPdfOutputIntent* other) {
	return new QPdfOutputIntent(*other);
}

void QPdfOutputIntent_OperatorAssign(QPdfOutputIntent* self, QPdfOutputIntent* other) {
	self->operator=(*other);
}

void QPdfOutputIntent_Swap(QPdfOutputIntent* self, QPdfOutputIntent* other) {
	self->swap(*other);
}

struct miqt_string QPdfOutputIntent_OutputConditionIdentifier(const QPdfOutputIntent* self) {
	QString _ret = self->outputConditionIdentifier();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPdfOutputIntent_SetOutputConditionIdentifier(QPdfOutputIntent* self, struct miqt_string identifier) {
	QString identifier_QString = QString::fromUtf8(identifier.data, identifier.len);
	self->setOutputConditionIdentifier(identifier_QString);
}

struct miqt_string QPdfOutputIntent_OutputCondition(const QPdfOutputIntent* self) {
	QString _ret = self->outputCondition();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPdfOutputIntent_SetOutputCondition(QPdfOutputIntent* self, struct miqt_string condition) {
	QString condition_QString = QString::fromUtf8(condition.data, condition.len);
	self->setOutputCondition(condition_QString);
}

QUrl* QPdfOutputIntent_RegistryName(const QPdfOutputIntent* self) {
	return new QUrl(self->registryName());
}

void QPdfOutputIntent_SetRegistryName(QPdfOutputIntent* self, QUrl* name) {
	self->setRegistryName(*name);
}

QColorSpace* QPdfOutputIntent_OutputProfile(const QPdfOutputIntent* self) {
	return new QColorSpace(self->outputProfile());
}

void QPdfOutputIntent_SetOutputProfile(QPdfOutputIntent* self, QColorSpace* profile) {
	self->setOutputProfile(*profile);
}

void QPdfOutputIntent_Delete(QPdfOutputIntent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPdfOutputIntent*>( self );
	} else {
		delete self;
	}
}

