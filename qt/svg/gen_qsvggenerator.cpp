// +build ignore

#include <QIODevice>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QRect>
#include <QRectF>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSvgGenerator>
#include <qsvggenerator.h>
#include "gen_qsvggenerator.h"
QSvgGenerator* QSvgGenerator_new() {
return new QSvgGenerator();
}
QSvgGenerator* QSvgGenerator_new2(SvgVersion version) {
return new QSvgGenerator(version);
}
void QSvgGenerator_virtbase(QSvgGenerator* src
, QPaintDevice** outptr_QPaintDevice
) {
*outptr_QPaintDevice = static_cast<QPaintDevice*>(src);
}
struct miqt_string QSvgGenerator_Title(const QSvgGenerator* self) {
	QString _ret = self->title();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSvgGenerator_SetTitle(QSvgGenerator* self, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	self->setTitle(title_QString);
}
struct miqt_string QSvgGenerator_Description(const QSvgGenerator* self) {
	QString _ret = self->description();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSvgGenerator_SetDescription(QSvgGenerator* self, struct miqt_string description) {
	QString description_QString = QString::fromUtf8(description.data, description.len);
	self->setDescription(description_QString);
}
QSize* QSvgGenerator_Size(const QSvgGenerator* self) {
	return new QSize(self->size());
}
void QSvgGenerator_SetSize(QSvgGenerator* self, QSize* size) {
	self->setSize(*size);
}
QRect* QSvgGenerator_ViewBox(const QSvgGenerator* self) {
	return new QRect(self->viewBox());
}
QRectF* QSvgGenerator_ViewBoxF(const QSvgGenerator* self) {
	return new QRectF(self->viewBoxF());
}
void QSvgGenerator_SetViewBox(QSvgGenerator* self, QRect* viewBox) {
	self->setViewBox(*viewBox);
}
void QSvgGenerator_SetViewBoxWithViewBox(QSvgGenerator* self, QRectF* viewBox) {
	self->setViewBox(*viewBox);
}
struct miqt_string QSvgGenerator_FileName(const QSvgGenerator* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSvgGenerator_SetFileName(QSvgGenerator* self, struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setFileName(fileName_QString);
}
QIODevice* QSvgGenerator_OutputDevice(const QSvgGenerator* self) {
	return self->outputDevice();
}
void QSvgGenerator_SetOutputDevice(QSvgGenerator* self, QIODevice* outputDevice) {
	self->setOutputDevice(outputDevice);
}
void QSvgGenerator_SetResolution(QSvgGenerator* self, int dpi) {
	self->setResolution(static_cast<int>(dpi));
}
int QSvgGenerator_Resolution(const QSvgGenerator* self) {
	return self->resolution();
}
SvgVersion QSvgGenerator_SvgVersion(const QSvgGenerator* self) {
	return self->svgVersion();
}
void QSvgGenerator_Delete(QSvgGenerator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QSvgGenerator*>( self );
	} else {
		delete self;
	}
}
