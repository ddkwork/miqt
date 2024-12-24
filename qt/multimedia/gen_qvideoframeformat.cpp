// +build ignore

#include <QRect>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVideoFrameFormat>
#include <qvideoframeformat.h>
#include "gen_qvideoframeformat.h"

QVideoFrameFormat* QVideoFrameFormat_new() {
	return new QVideoFrameFormat();
}

QVideoFrameFormat* QVideoFrameFormat_new2(QSize* size, PixelFormat pixelFormat) {
	return new QVideoFrameFormat(*size, pixelFormat);
}

QVideoFrameFormat* QVideoFrameFormat_new3(QVideoFrameFormat* format) {
	return new QVideoFrameFormat(*format);
}

void QVideoFrameFormat_Swap(QVideoFrameFormat* self, QVideoFrameFormat* other) {
	self->swap(*other);
}

void QVideoFrameFormat_Detach(QVideoFrameFormat* self) {
	self->detach();
}

void QVideoFrameFormat_OperatorAssign(QVideoFrameFormat* self, QVideoFrameFormat* format) {
	self->operator=(*format);
}

bool QVideoFrameFormat_OperatorEqual(const QVideoFrameFormat* self, QVideoFrameFormat* format) {
	return (*self == *format);
}

bool QVideoFrameFormat_OperatorNotEqual(const QVideoFrameFormat* self, QVideoFrameFormat* format) {
	return (*self != *format);
}

bool QVideoFrameFormat_IsValid(const QVideoFrameFormat* self) {
	return self->isValid();
}

int QVideoFrameFormat_PixelFormat(const QVideoFrameFormat* self) {
	QVideoFrameFormat::PixelFormat _ret = self->pixelFormat();
	return static_cast<int>(_ret);
}

QSize* QVideoFrameFormat_FrameSize(const QVideoFrameFormat* self) {
	return new QSize(self->frameSize());
}

void QVideoFrameFormat_SetFrameSize(QVideoFrameFormat* self, QSize* size) {
	self->setFrameSize(*size);
}

void QVideoFrameFormat_SetFrameSize2(QVideoFrameFormat* self, int width, int height) {
	self->setFrameSize(static_cast<int>(width), static_cast<int>(height));
}

int QVideoFrameFormat_FrameWidth(const QVideoFrameFormat* self) {
	return self->frameWidth();
}

int QVideoFrameFormat_FrameHeight(const QVideoFrameFormat* self) {
	return self->frameHeight();
}

int QVideoFrameFormat_PlaneCount(const QVideoFrameFormat* self) {
	return self->planeCount();
}

QRect* QVideoFrameFormat_Viewport(const QVideoFrameFormat* self) {
	return new QRect(self->viewport());
}

void QVideoFrameFormat_SetViewport(QVideoFrameFormat* self, QRect* viewport) {
	self->setViewport(*viewport);
}

Direction QVideoFrameFormat_ScanLineDirection(const QVideoFrameFormat* self) {
	return self->scanLineDirection();
}

void QVideoFrameFormat_SetScanLineDirection(QVideoFrameFormat* self, Direction direction) {
	self->setScanLineDirection(direction);
}

double QVideoFrameFormat_FrameRate(const QVideoFrameFormat* self) {
	qreal _ret = self->frameRate();
	return static_cast<double>(_ret);
}

void QVideoFrameFormat_SetFrameRate(QVideoFrameFormat* self, double rate) {
	self->setFrameRate(static_cast<qreal>(rate));
}

double QVideoFrameFormat_StreamFrameRate(const QVideoFrameFormat* self) {
	qreal _ret = self->streamFrameRate();
	return static_cast<double>(_ret);
}

void QVideoFrameFormat_SetStreamFrameRate(QVideoFrameFormat* self, double rate) {
	self->setStreamFrameRate(static_cast<qreal>(rate));
}

YCbCrColorSpace QVideoFrameFormat_YCbCrColorSpace(const QVideoFrameFormat* self) {
	return self->yCbCrColorSpace();
}

void QVideoFrameFormat_SetYCbCrColorSpace(QVideoFrameFormat* self, YCbCrColorSpace colorSpace) {
	self->setYCbCrColorSpace(colorSpace);
}

ColorSpace QVideoFrameFormat_ColorSpace(const QVideoFrameFormat* self) {
	return self->colorSpace();
}

void QVideoFrameFormat_SetColorSpace(QVideoFrameFormat* self, ColorSpace colorSpace) {
	self->setColorSpace(colorSpace);
}

ColorTransfer QVideoFrameFormat_ColorTransfer(const QVideoFrameFormat* self) {
	return self->colorTransfer();
}

void QVideoFrameFormat_SetColorTransfer(QVideoFrameFormat* self, ColorTransfer colorTransfer) {
	self->setColorTransfer(colorTransfer);
}

ColorRange QVideoFrameFormat_ColorRange(const QVideoFrameFormat* self) {
	return self->colorRange();
}

void QVideoFrameFormat_SetColorRange(QVideoFrameFormat* self, ColorRange rangeVal) {
	self->setColorRange(rangeVal);
}

bool QVideoFrameFormat_IsMirrored(const QVideoFrameFormat* self) {
	return self->isMirrored();
}

void QVideoFrameFormat_SetMirrored(QVideoFrameFormat* self, bool mirrored) {
	self->setMirrored(mirrored);
}

int QVideoFrameFormat_Rotation(const QVideoFrameFormat* self) {
	QtVideo::Rotation _ret = self->rotation();
	return static_cast<int>(_ret);
}

void QVideoFrameFormat_SetRotation(QVideoFrameFormat* self, int rotation) {
	self->setRotation(static_cast<QtVideo::Rotation>(rotation));
}

struct miqt_string QVideoFrameFormat_VertexShaderFileName(const QVideoFrameFormat* self) {
	QString _ret = self->vertexShaderFileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QVideoFrameFormat_FragmentShaderFileName(const QVideoFrameFormat* self) {
	QString _ret = self->fragmentShaderFileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

float QVideoFrameFormat_MaxLuminance(const QVideoFrameFormat* self) {
	return self->maxLuminance();
}

void QVideoFrameFormat_SetMaxLuminance(QVideoFrameFormat* self, float lum) {
	self->setMaxLuminance(static_cast<float>(lum));
}

PixelFormat QVideoFrameFormat_PixelFormatFromImageFormat(int format) {
	return QVideoFrameFormat::pixelFormatFromImageFormat(static_cast<QImage::Format>(format));
}

int QVideoFrameFormat_ImageFormatFromPixelFormat(PixelFormat format) {
	QImage::Format _ret = QVideoFrameFormat::imageFormatFromPixelFormat(format);
	return static_cast<int>(_ret);
}

struct miqt_string QVideoFrameFormat_PixelFormatToString(int pixelFormat) {
	QString _ret = QVideoFrameFormat::pixelFormatToString(static_cast<QVideoFrameFormat::PixelFormat>(pixelFormat));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QVideoFrameFormat_Delete(QVideoFrameFormat* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QVideoFrameFormat*>( self );
	} else {
		delete self;
	}
}

