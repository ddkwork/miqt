// +build ignore

#include <QGlyphRun>
#include <QList>
#include <QPointF>
#include <QRawFont>
#include <QRectF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qglyphrun.h>
#include "gen_qglyphrun.h"

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

QGlyphRun* QGlyphRun_new() {
	return new QGlyphRun();
}

QGlyphRun* QGlyphRun_new2(QGlyphRun* other) {
	return new QGlyphRun(*other);
}

void QGlyphRun_OperatorAssign(QGlyphRun* self, QGlyphRun* other) {
	self->operator=(*other);
}

void QGlyphRun_Swap(QGlyphRun* self, QGlyphRun* other) {
	self->swap(*other);
}

QRawFont* QGlyphRun_RawFont(const QGlyphRun* self) {
	return new QRawFont(self->rawFont());
}

void QGlyphRun_SetRawFont(QGlyphRun* self, QRawFont* rawFont) {
	self->setRawFont(*rawFont);
}

void QGlyphRun_SetRawData(QGlyphRun* self, const unsigned int* glyphIndexArray, QPointF* glyphPositionArray, int size) {
	self->setRawData(static_cast<const quint32*>(glyphIndexArray), glyphPositionArray, static_cast<int>(size));
}

struct miqt_array /* of unsigned int */  QGlyphRun_GlyphIndexes(const QGlyphRun* self) {
	QList<quint32> _ret = self->glyphIndexes();
	// Convert QList<> from C++ memory to manually-managed C memory
	unsigned int* _arr = static_cast<unsigned int*>(malloc(sizeof(unsigned int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QGlyphRun_SetGlyphIndexes(QGlyphRun* self, struct miqt_array /* of unsigned int */  glyphIndexes) {
	QList<quint32> glyphIndexes_QList;
	glyphIndexes_QList.reserve(glyphIndexes.len);
	unsigned int* glyphIndexes_arr = static_cast<unsigned int*>(glyphIndexes.data);
	for(size_t i = 0; i < glyphIndexes.len; ++i) {
		glyphIndexes_QList.push_back(static_cast<unsigned int>(glyphIndexes_arr[i]));
	}
	self->setGlyphIndexes(glyphIndexes_QList);
}

struct miqt_array /* of QPointF* */  QGlyphRun_Positions(const QGlyphRun* self) {
	QList<QPointF> _ret = self->positions();
	// Convert QList<> from C++ memory to manually-managed C memory
	QPointF** _arr = static_cast<QPointF**>(malloc(sizeof(QPointF*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QPointF(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QGlyphRun_SetPositions(QGlyphRun* self, struct miqt_array /* of QPointF* */  positions) {
	QList<QPointF> positions_QList;
	positions_QList.reserve(positions.len);
	QPointF** positions_arr = static_cast<QPointF**>(positions.data);
	for(size_t i = 0; i < positions.len; ++i) {
		positions_QList.push_back(*(positions_arr[i]));
	}
	self->setPositions(positions_QList);
}

void QGlyphRun_Clear(QGlyphRun* self) {
	self->clear();
}

bool QGlyphRun_OperatorEqual(const QGlyphRun* self, QGlyphRun* other) {
	return (*self == *other);
}

bool QGlyphRun_OperatorNotEqual(const QGlyphRun* self, QGlyphRun* other) {
	return (*self != *other);
}

void QGlyphRun_SetOverline(QGlyphRun* self, bool overline) {
	self->setOverline(overline);
}

bool QGlyphRun_Overline(const QGlyphRun* self) {
	return self->overline();
}

void QGlyphRun_SetUnderline(QGlyphRun* self, bool underline) {
	self->setUnderline(underline);
}

bool QGlyphRun_Underline(const QGlyphRun* self) {
	return self->underline();
}

void QGlyphRun_SetStrikeOut(QGlyphRun* self, bool strikeOut) {
	self->setStrikeOut(strikeOut);
}

bool QGlyphRun_StrikeOut(const QGlyphRun* self) {
	return self->strikeOut();
}

void QGlyphRun_SetRightToLeft(QGlyphRun* self, bool on) {
	self->setRightToLeft(on);
}

bool QGlyphRun_IsRightToLeft(const QGlyphRun* self) {
	return self->isRightToLeft();
}

void QGlyphRun_SetFlag(QGlyphRun* self, GlyphRunFlag flag) {
	self->setFlag(flag);
}

void QGlyphRun_SetFlags(QGlyphRun* self, GlyphRunFlags flags) {
	self->setFlags(flags);
}

GlyphRunFlags QGlyphRun_Flags(const QGlyphRun* self) {
	return self->flags();
}

void QGlyphRun_SetBoundingRect(QGlyphRun* self, QRectF* boundingRect) {
	self->setBoundingRect(*boundingRect);
}

QRectF* QGlyphRun_BoundingRect(const QGlyphRun* self) {
	return new QRectF(self->boundingRect());
}

struct miqt_array /* of ptrdiff_t */  QGlyphRun_StringIndexes(const QGlyphRun* self) {
	QList<qsizetype> _ret = self->stringIndexes();
	// Convert QList<> from C++ memory to manually-managed C memory
	ptrdiff_t* _arr = static_cast<ptrdiff_t*>(malloc(sizeof(ptrdiff_t) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QGlyphRun_SetStringIndexes(QGlyphRun* self, struct miqt_array /* of ptrdiff_t */  stringIndexes) {
	QList<qsizetype> stringIndexes_QList;
	stringIndexes_QList.reserve(stringIndexes.len);
	ptrdiff_t* stringIndexes_arr = static_cast<ptrdiff_t*>(stringIndexes.data);
	for(size_t i = 0; i < stringIndexes.len; ++i) {
		stringIndexes_QList.push_back(static_cast<QIntegerForSizeof<std::size_t>::Signed>(stringIndexes_arr[i]));
	}
	self->setStringIndexes(stringIndexes_QList);
}

void QGlyphRun_SetSourceString(QGlyphRun* self, struct miqt_string sourceString) {
	QString sourceString_QString = QString::fromUtf8(sourceString.data, sourceString.len);
	self->setSourceString(sourceString_QString);
}

struct miqt_string QGlyphRun_SourceString(const QGlyphRun* self) {
	QString _ret = self->sourceString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QGlyphRun_IsEmpty(const QGlyphRun* self) {
	return self->isEmpty();
}

void QGlyphRun_SetFlag2(QGlyphRun* self, GlyphRunFlag flag, bool enabled) {
	self->setFlag(flag, enabled);
}

void QGlyphRun_Delete(QGlyphRun* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QGlyphRun*>( self );
	} else {
		delete self;
	}
}

