// +build ignore

#define WORKAROUND_INNER_CLASS_DEFINITION_QFont__Tag
#include <QFontVariableAxis>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qfontvariableaxis.h>
#include "gen_qfontvariableaxis.h"

#ifndef _Bool
#define _Bool bool
#endif

QFontVariableAxis* QFontVariableAxis_new() {
	return new QFontVariableAxis();
}

QFontVariableAxis* QFontVariableAxis_new2(QFontVariableAxis* axis) {
	return new QFontVariableAxis(*axis);
}

void QFontVariableAxis_Swap(QFontVariableAxis* self, QFontVariableAxis* other) {
	self->swap(*other);
}

void QFontVariableAxis_OperatorAssign(QFontVariableAxis* self, QFontVariableAxis* axis) {
	self->operator=(*axis);
}

QFont__Tag* QFontVariableAxis_Tag(const QFontVariableAxis* self) {
	return new QFont::Tag(self->tag());
}

void QFontVariableAxis_SetTag(QFontVariableAxis* self, QFont__Tag* tag) {
	self->setTag(*tag);
}

struct miqt_string QFontVariableAxis_Name(const QFontVariableAxis* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QFontVariableAxis_SetName(QFontVariableAxis* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setName(name_QString);
}

double QFontVariableAxis_MinimumValue(const QFontVariableAxis* self) {
	qreal _ret = self->minimumValue();
	return static_cast<double>(_ret);
}

void QFontVariableAxis_SetMinimumValue(QFontVariableAxis* self, double minimumValue) {
	self->setMinimumValue(static_cast<qreal>(minimumValue));
}

double QFontVariableAxis_MaximumValue(const QFontVariableAxis* self) {
	qreal _ret = self->maximumValue();
	return static_cast<double>(_ret);
}

void QFontVariableAxis_SetMaximumValue(QFontVariableAxis* self, double maximumValue) {
	self->setMaximumValue(static_cast<qreal>(maximumValue));
}

double QFontVariableAxis_DefaultValue(const QFontVariableAxis* self) {
	qreal _ret = self->defaultValue();
	return static_cast<double>(_ret);
}

void QFontVariableAxis_SetDefaultValue(QFontVariableAxis* self, double defaultValue) {
	self->setDefaultValue(static_cast<qreal>(defaultValue));
}

void QFontVariableAxis_Delete(QFontVariableAxis* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFontVariableAxis*>( self );
	} else {
		delete self;
	}
}

