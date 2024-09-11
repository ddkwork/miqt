#include <QByteArray>
#include <QIODevice>
#include <QImage>
#include <QImageIOHandler>
#include <QImageIOPlugin>
#include <QMetaObject>
#include <QRect>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include "qimageiohandler.h"

#include "gen_qimageiohandler.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

void QImageIOHandler_SetDevice(QImageIOHandler* self, QIODevice* device) {
	self->setDevice(device);
}

QIODevice* QImageIOHandler_Device(const QImageIOHandler* self) {
	return self->device();
}

void QImageIOHandler_SetFormat(QImageIOHandler* self, QByteArray* format) {
	self->setFormat(*format);
}

void QImageIOHandler_SetFormatWithFormat(const QImageIOHandler* self, QByteArray* format) {
	self->setFormat(*format);
}

QByteArray* QImageIOHandler_Format(const QImageIOHandler* self) {
	QByteArray ret = self->format();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(ret));
}

QByteArray* QImageIOHandler_Name(const QImageIOHandler* self) {
	QByteArray ret = self->name();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(ret));
}

bool QImageIOHandler_CanRead(const QImageIOHandler* self) {
	return self->canRead();
}

bool QImageIOHandler_Read(QImageIOHandler* self, QImage* image) {
	return self->read(image);
}

bool QImageIOHandler_Write(QImageIOHandler* self, QImage* image) {
	return self->write(*image);
}

QVariant* QImageIOHandler_Option(const QImageIOHandler* self, uintptr_t option) {
	QVariant ret = self->option(static_cast<QImageIOHandler::ImageOption>(option));
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QVariant*>(new QVariant(ret));
}

void QImageIOHandler_SetOption(QImageIOHandler* self, uintptr_t option, QVariant* value) {
	self->setOption(static_cast<QImageIOHandler::ImageOption>(option), *value);
}

bool QImageIOHandler_SupportsOption(const QImageIOHandler* self, uintptr_t option) {
	return self->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

bool QImageIOHandler_JumpToNextImage(QImageIOHandler* self) {
	return self->jumpToNextImage();
}

bool QImageIOHandler_JumpToImage(QImageIOHandler* self, int imageNumber) {
	return self->jumpToImage(static_cast<int>(imageNumber));
}

int QImageIOHandler_LoopCount(const QImageIOHandler* self) {
	return self->loopCount();
}

int QImageIOHandler_ImageCount(const QImageIOHandler* self) {
	return self->imageCount();
}

int QImageIOHandler_NextImageDelay(const QImageIOHandler* self) {
	return self->nextImageDelay();
}

int QImageIOHandler_CurrentImageNumber(const QImageIOHandler* self) {
	return self->currentImageNumber();
}

QRect* QImageIOHandler_CurrentImageRect(const QImageIOHandler* self) {
	QRect ret = self->currentImageRect();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(ret));
}

void QImageIOHandler_Delete(QImageIOHandler* self) {
	delete self;
}

QMetaObject* QImageIOPlugin_MetaObject(const QImageIOPlugin* self) {
	return (QMetaObject*) self->metaObject();
}

void QImageIOPlugin_Tr(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QImageIOPlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QImageIOPlugin_TrUtf8(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QImageIOPlugin::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

int QImageIOPlugin_Capabilities(const QImageIOPlugin* self, QIODevice* device, QByteArray* format) {
	QImageIOPlugin::Capabilities ret = self->capabilities(device, *format);
	return static_cast<int>(ret);
}

QImageIOHandler* QImageIOPlugin_Create(const QImageIOPlugin* self, QIODevice* device) {
	return self->create(device);
}

void QImageIOPlugin_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QImageIOPlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QImageIOPlugin_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QImageIOPlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QImageIOPlugin_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QImageIOPlugin::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QImageIOPlugin_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QImageIOPlugin::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

QImageIOHandler* QImageIOPlugin_Create2(const QImageIOPlugin* self, QIODevice* device, QByteArray* format) {
	return self->create(device, *format);
}

void QImageIOPlugin_Delete(QImageIOPlugin* self) {
	delete self;
}

