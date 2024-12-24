// +build ignore

#include <QByteArray>
#include <QIODevice>
#include <QImage>
#include <QImageIOHandler>
#include <QImageIOPlugin>
#include <QMetaObject>
#include <QObject>
#include <QRect>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qimageiohandler.h>
#include "gen_qimageiohandler.h"
QImageIOHandler* QImageIOHandler_new() {
return new QImageIOHandler();
}
void QImageIOHandler_SetDevice(QImageIOHandler* self, QIODevice* device) {
	self->setDevice(device);
}
QIODevice* QImageIOHandler_Device(const QImageIOHandler* self) {
	return self->device();
}
void QImageIOHandler_SetFormat(QImageIOHandler* self, struct miqt_string format) {
	QByteArray format_QByteArray(format.data, format.len);
	self->setFormat(format_QByteArray);
}
void QImageIOHandler_SetFormatWithFormat(const QImageIOHandler* self, struct miqt_string format) {
	QByteArray format_QByteArray(format.data, format.len);
	self->setFormat(format_QByteArray);
}
struct miqt_string QImageIOHandler_Format(const QImageIOHandler* self) {
	QByteArray _qb = self->format();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
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
QVariant* QImageIOHandler_Option(const QImageIOHandler* self, ImageOption option) {
	return new QVariant(self->option(option));
}
void QImageIOHandler_SetOption(QImageIOHandler* self, ImageOption option, QVariant* value) {
	self->setOption(option, *value);
}
bool QImageIOHandler_SupportsOption(const QImageIOHandler* self, ImageOption option) {
	return self->supportsOption(option);
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
	return new QRect(self->currentImageRect());
}
bool QImageIOHandler_AllocateImage(QSize* size, int format, QImage* image) {
	return QImageIOHandler::allocateImage(*size, static_cast<QImage::Format>(format), image);
}
void QImageIOHandler_Delete(QImageIOHandler* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QImageIOHandler*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQImageIOPlugin : public virtual QImageIOPlugin {
public:
MiqtVirtualQImageIOPlugin(): QImageIOPlugin() {};
MiqtVirtualQImageIOPlugin(QObject* parent): QImageIOPlugin(parent) {};

virtual ~MiqtVirtualQImageIOPlugin() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QImageIOPlugin::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QImageIOPlugin_MetaObject(const_cast<MiqtVirtualQImageIOPlugin*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QImageIOPlugin::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QImageIOPlugin::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QImageIOPlugin_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QImageIOPlugin::qt_metacast(param1);

	}
};
QImageIOPlugin* QImageIOPlugin_new() {
return new MiqtVirtualQImageIOPlugin();
}
QImageIOPlugin* QImageIOPlugin_new2(QObject* parent) {
return new MiqtVirtualQImageIOPlugin(parent);
}
void QImageIOPlugin_virtbase(QImageIOPlugin* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QImageIOPlugin_MetaObject(const QImageIOPlugin* self) {
	return (QMetaObject*) self->metaObject();
}
void* QImageIOPlugin_Metacast(QImageIOPlugin* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QImageIOPlugin_Tr(const char* s) {
	QString _ret = QImageIOPlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
Capabilities QImageIOPlugin_Capabilities(const QImageIOPlugin* self, QIODevice* device, struct miqt_string format) {
	QByteArray format_QByteArray(format.data, format.len);
	return self->capabilities(device, format_QByteArray);
}
QImageIOHandler* QImageIOPlugin_Create(const QImageIOPlugin* self, QIODevice* device, struct miqt_string format) {
	QByteArray format_QByteArray(format.data, format.len);
	return self->create(device, format_QByteArray);
}
struct miqt_string QImageIOPlugin_Tr2(const char* s, const char* c) {
	QString _ret = QImageIOPlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QImageIOPlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QImageIOPlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QImageIOPlugin_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQImageIOPlugin*>( (QImageIOPlugin*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QImageIOPlugin_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQImageIOPlugin*)(self) )->virtualbase_MetaObject();
}
void QImageIOPlugin_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQImageIOPlugin*>( (QImageIOPlugin*)(self) )->handle__Metacast = slot;
}
void* QImageIOPlugin_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQImageIOPlugin*)(self) )->virtualbase_Metacast(param1);
}
void QImageIOPlugin_Delete(QImageIOPlugin* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQImageIOPlugin*>( self );
	} else {
		delete self;
	}
}
