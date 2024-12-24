// +build ignore

#include <QPageSize>
#include <QRect>
#include <QRectF>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qpagesize.h>
#include "gen_qpagesize.h"

#ifndef _Bool
#define _Bool bool
#endif

QPageSize* QPageSize_new() {
	return new QPageSize();
}

QPageSize* QPageSize_new2(PageSizeId pageSizeId) {
	return new QPageSize(pageSizeId);
}

QPageSize* QPageSize_new3(QSize* pointSize) {
	return new QPageSize(*pointSize);
}

QPageSize* QPageSize_new4(QSizeF* size, Unit units) {
	return new QPageSize(*size, units);
}

QPageSize* QPageSize_new5(QPageSize* other) {
	return new QPageSize(*other);
}

QPageSize* QPageSize_new6(QSize* pointSize, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new QPageSize(*pointSize, name_QString);
}

QPageSize* QPageSize_new7(QSize* pointSize, struct miqt_string name, SizeMatchPolicy matchPolicy) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new QPageSize(*pointSize, name_QString, matchPolicy);
}

QPageSize* QPageSize_new8(QSizeF* size, Unit units, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new QPageSize(*size, units, name_QString);
}

QPageSize* QPageSize_new9(QSizeF* size, Unit units, struct miqt_string name, SizeMatchPolicy matchPolicy) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new QPageSize(*size, units, name_QString, matchPolicy);
}

void QPageSize_OperatorAssign(QPageSize* self, QPageSize* other) {
	self->operator=(*other);
}

void QPageSize_Swap(QPageSize* self, QPageSize* other) {
	self->swap(*other);
}

bool QPageSize_IsEquivalentTo(const QPageSize* self, QPageSize* other) {
	return self->isEquivalentTo(*other);
}

bool QPageSize_IsValid(const QPageSize* self) {
	return self->isValid();
}

struct miqt_string QPageSize_Key(const QPageSize* self) {
	QString _ret = self->key();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPageSize_Name(const QPageSize* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

PageSizeId QPageSize_Id(const QPageSize* self) {
	return self->id();
}

int QPageSize_WindowsId(const QPageSize* self) {
	return self->windowsId();
}

QSizeF* QPageSize_DefinitionSize(const QPageSize* self) {
	return new QSizeF(self->definitionSize());
}

Unit QPageSize_DefinitionUnits(const QPageSize* self) {
	return self->definitionUnits();
}

QSizeF* QPageSize_Size(const QPageSize* self, Unit units) {
	return new QSizeF(self->size(units));
}

QSize* QPageSize_SizePoints(const QPageSize* self) {
	return new QSize(self->sizePoints());
}

QSize* QPageSize_SizePixels(const QPageSize* self, int resolution) {
	return new QSize(self->sizePixels(static_cast<int>(resolution)));
}

QRectF* QPageSize_Rect(const QPageSize* self, Unit units) {
	return new QRectF(self->rect(units));
}

QRect* QPageSize_RectPoints(const QPageSize* self) {
	return new QRect(self->rectPoints());
}

QRect* QPageSize_RectPixels(const QPageSize* self, int resolution) {
	return new QRect(self->rectPixels(static_cast<int>(resolution)));
}

struct miqt_string QPageSize_KeyWithPageSizeId(PageSizeId pageSizeId) {
	QString _ret = QPageSize::key(pageSizeId);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPageSize_NameWithPageSizeId(PageSizeId pageSizeId) {
	QString _ret = QPageSize::name(pageSizeId);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

PageSizeId QPageSize_IdWithPointSize(QSize* pointSize) {
	return QPageSize::id(*pointSize);
}

PageSizeId QPageSize_Id2(QSizeF* size, Unit units) {
	return QPageSize::id(*size, units);
}

PageSizeId QPageSize_IdWithWindowsId(int windowsId) {
	return QPageSize::id(static_cast<int>(windowsId));
}

int QPageSize_WindowsIdWithPageSizeId(PageSizeId pageSizeId) {
	return QPageSize::windowsId(pageSizeId);
}

QSizeF* QPageSize_DefinitionSizeWithPageSizeId(PageSizeId pageSizeId) {
	return new QSizeF(QPageSize::definitionSize(pageSizeId));
}

Unit QPageSize_DefinitionUnitsWithPageSizeId(PageSizeId pageSizeId) {
	return QPageSize::definitionUnits(pageSizeId);
}

QSizeF* QPageSize_Size2(PageSizeId pageSizeId, Unit units) {
	return new QSizeF(QPageSize::size(pageSizeId, units));
}

QSize* QPageSize_SizePointsWithPageSizeId(PageSizeId pageSizeId) {
	return new QSize(QPageSize::sizePoints(pageSizeId));
}

QSize* QPageSize_SizePixels2(PageSizeId pageSizeId, int resolution) {
	return new QSize(QPageSize::sizePixels(pageSizeId, static_cast<int>(resolution)));
}

PageSizeId QPageSize_Id22(QSize* pointSize, SizeMatchPolicy matchPolicy) {
	return QPageSize::id(*pointSize, matchPolicy);
}

PageSizeId QPageSize_Id3(QSizeF* size, Unit units, SizeMatchPolicy matchPolicy) {
	return QPageSize::id(*size, units, matchPolicy);
}

void QPageSize_Delete(QPageSize* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPageSize*>( self );
	} else {
		delete self;
	}
}

