// +build ignore

#include <QFont>
#include <QFontInfo>
#include <QFontVariableAxis>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qfontinfo.h>
#include "gen_qfontinfo.h"

#ifndef _Bool
#define _Bool bool
#endif

QFontInfo* QFontInfo_new(QFont* param1) {
	return new QFontInfo(*param1);
}

QFontInfo* QFontInfo_new2(QFontInfo* param1) {
	return new QFontInfo(*param1);
}

void QFontInfo_OperatorAssign(QFontInfo* self, QFontInfo* param1) {
	self->operator=(*param1);
}

void QFontInfo_Swap(QFontInfo* self, QFontInfo* other) {
	self->swap(*other);
}

struct miqt_string QFontInfo_Family(const QFontInfo* self) {
	QString _ret = self->family();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QFontInfo_StyleName(const QFontInfo* self) {
	QString _ret = self->styleName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QFontInfo_PixelSize(const QFontInfo* self) {
	return self->pixelSize();
}

int QFontInfo_PointSize(const QFontInfo* self) {
	return self->pointSize();
}

double QFontInfo_PointSizeF(const QFontInfo* self) {
	qreal _ret = self->pointSizeF();
	return static_cast<double>(_ret);
}

bool QFontInfo_Italic(const QFontInfo* self) {
	return self->italic();
}

int QFontInfo_Style(const QFontInfo* self) {
	QFont::Style _ret = self->style();
	return static_cast<int>(_ret);
}

int QFontInfo_Weight(const QFontInfo* self) {
	return self->weight();
}

bool QFontInfo_Bold(const QFontInfo* self) {
	return self->bold();
}

bool QFontInfo_Underline(const QFontInfo* self) {
	return self->underline();
}

bool QFontInfo_Overline(const QFontInfo* self) {
	return self->overline();
}

bool QFontInfo_StrikeOut(const QFontInfo* self) {
	return self->strikeOut();
}

bool QFontInfo_FixedPitch(const QFontInfo* self) {
	return self->fixedPitch();
}

int QFontInfo_StyleHint(const QFontInfo* self) {
	QFont::StyleHint _ret = self->styleHint();
	return static_cast<int>(_ret);
}

struct miqt_array /* of QFontVariableAxis* */  QFontInfo_VariableAxes(const QFontInfo* self) {
	QList<QFontVariableAxis> _ret = self->variableAxes();
	// Convert QList<> from C++ memory to manually-managed C memory
	QFontVariableAxis** _arr = static_cast<QFontVariableAxis**>(malloc(sizeof(QFontVariableAxis*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QFontVariableAxis(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

int QFontInfo_LegacyWeight(const QFontInfo* self) {
	return self->legacyWeight();
}

bool QFontInfo_ExactMatch(const QFontInfo* self) {
	return self->exactMatch();
}

void QFontInfo_Delete(QFontInfo* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFontInfo*>( self );
	} else {
		delete self;
	}
}

