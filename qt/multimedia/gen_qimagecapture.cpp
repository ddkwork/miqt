// +build ignore

#include <QImage>
#include <QImageCapture>
#include <QList>
#include <QMediaCaptureSession>
#include <QMediaMetaData>
#include <QMetaObject>
#include <QObject>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVideoFrame>
#include <qimagecapture.h>
#include "gen_qimagecapture.h"
class MiqtVirtualQImageCapture : public virtual QImageCapture {
public:
MiqtVirtualQImageCapture(): QImageCapture() {};
MiqtVirtualQImageCapture(QObject* parent): QImageCapture(parent) {};

virtual ~MiqtVirtualQImageCapture() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QImageCapture::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QImageCapture_MetaObject(const_cast<MiqtVirtualQImageCapture*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QImageCapture::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QImageCapture::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QImageCapture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QImageCapture::qt_metacast(param1);

	}
};
QImageCapture* QImageCapture_new() {
return new MiqtVirtualQImageCapture();
}
QImageCapture* QImageCapture_new2(QObject* parent) {
return new MiqtVirtualQImageCapture(parent);
}
void QImageCapture_virtbase(QImageCapture* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QImageCapture_MetaObject(const QImageCapture* self) {
	return (QMetaObject*) self->metaObject();
}
void* QImageCapture_Metacast(QImageCapture* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QImageCapture_Tr(const char* s) {
	QString _ret = QImageCapture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QImageCapture_IsAvailable(const QImageCapture* self) {
	return self->isAvailable();
}
QMediaCaptureSession* QImageCapture_CaptureSession(const QImageCapture* self) {
	return self->captureSession();
}
Error QImageCapture_Error(const QImageCapture* self) {
	return self->error();
}
struct miqt_string QImageCapture_ErrorString(const QImageCapture* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QImageCapture_IsReadyForCapture(const QImageCapture* self) {
	return self->isReadyForCapture();
}
FileFormat QImageCapture_FileFormat(const QImageCapture* self) {
	return self->fileFormat();
}
void QImageCapture_SetFileFormat(QImageCapture* self, FileFormat format) {
	self->setFileFormat(format);
}
struct miqt_array /* of FileFormat */  QImageCapture_SupportedFormats() {
	QList<FileFormat> _ret = QImageCapture::supportedFormats();
	// Convert QList<> from C++ memory to manually-managed C memory
	FileFormat* _arr = static_cast<FileFormat*>(malloc(sizeof(FileFormat) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_string QImageCapture_FileFormatName(FileFormat c) {
	QString _ret = QImageCapture::fileFormatName(c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QImageCapture_FileFormatDescription(FileFormat c) {
	QString _ret = QImageCapture::fileFormatDescription(c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QImageCapture_Resolution(const QImageCapture* self) {
	return new QSize(self->resolution());
}
void QImageCapture_SetResolution(QImageCapture* self, QSize* resolution) {
	self->setResolution(*resolution);
}
void QImageCapture_SetResolution2(QImageCapture* self, int width, int height) {
	self->setResolution(static_cast<int>(width), static_cast<int>(height));
}
Quality QImageCapture_Quality(const QImageCapture* self) {
	return self->quality();
}
void QImageCapture_SetQuality(QImageCapture* self, Quality quality) {
	self->setQuality(quality);
}
QMediaMetaData* QImageCapture_MetaData(const QImageCapture* self) {
	return new QMediaMetaData(self->metaData());
}
void QImageCapture_SetMetaData(QImageCapture* self, QMediaMetaData* metaData) {
	self->setMetaData(*metaData);
}
void QImageCapture_AddMetaData(QImageCapture* self, QMediaMetaData* metaData) {
	self->addMetaData(*metaData);
}
int QImageCapture_CaptureToFile(QImageCapture* self) {
	return self->captureToFile();
}
int QImageCapture_Capture(QImageCapture* self) {
	return self->capture();
}
void QImageCapture_ErrorChanged(QImageCapture* self) {
	self->errorChanged();
}
void QImageCapture_connect_ErrorChanged(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)()>(&QImageCapture::errorChanged), self, [=]() {
		miqt_exec_callback_QImageCapture_ErrorChanged(slot);
	});
}
void QImageCapture_ErrorOccurred(QImageCapture* self, int id, int error, struct miqt_string errorString) {
	QString errorString_QString = QString::fromUtf8(errorString.data, errorString.len);
	self->errorOccurred(static_cast<int>(id), static_cast<QImageCapture::Error>(error), errorString_QString);
}
void QImageCapture_connect_ErrorOccurred(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)(int, QImageCapture::Error, const QString&)>(&QImageCapture::errorOccurred), self, [=](int id, QImageCapture::Error error, const QString& errorString) {
		int sigval1 = id;
		QImageCapture::Error error_ret = error;
		int sigval2 = static_cast<int>(error_ret);
		const QString errorString_ret = errorString;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray errorString_b = errorString_ret.toUtf8();
		struct miqt_string errorString_ms;
		errorString_ms.len = errorString_b.length();
		errorString_ms.data = static_cast<char*>(malloc(errorString_ms.len));
		memcpy(errorString_ms.data, errorString_b.data(), errorString_ms.len);
		struct miqt_string sigval3 = errorString_ms;
		miqt_exec_callback_QImageCapture_ErrorOccurred(slot, sigval1, sigval2, sigval3);
	});
}
void QImageCapture_ReadyForCaptureChanged(QImageCapture* self, bool ready) {
	self->readyForCaptureChanged(ready);
}
void QImageCapture_connect_ReadyForCaptureChanged(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)(bool)>(&QImageCapture::readyForCaptureChanged), self, [=](bool ready) {
		bool sigval1 = ready;
		miqt_exec_callback_QImageCapture_ReadyForCaptureChanged(slot, sigval1);
	});
}
void QImageCapture_MetaDataChanged(QImageCapture* self) {
	self->metaDataChanged();
}
void QImageCapture_connect_MetaDataChanged(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)()>(&QImageCapture::metaDataChanged), self, [=]() {
		miqt_exec_callback_QImageCapture_MetaDataChanged(slot);
	});
}
void QImageCapture_FileFormatChanged(QImageCapture* self) {
	self->fileFormatChanged();
}
void QImageCapture_connect_FileFormatChanged(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)()>(&QImageCapture::fileFormatChanged), self, [=]() {
		miqt_exec_callback_QImageCapture_FileFormatChanged(slot);
	});
}
void QImageCapture_QualityChanged(QImageCapture* self) {
	self->qualityChanged();
}
void QImageCapture_connect_QualityChanged(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)()>(&QImageCapture::qualityChanged), self, [=]() {
		miqt_exec_callback_QImageCapture_QualityChanged(slot);
	});
}
void QImageCapture_ResolutionChanged(QImageCapture* self) {
	self->resolutionChanged();
}
void QImageCapture_connect_ResolutionChanged(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)()>(&QImageCapture::resolutionChanged), self, [=]() {
		miqt_exec_callback_QImageCapture_ResolutionChanged(slot);
	});
}
void QImageCapture_ImageExposed(QImageCapture* self, int id) {
	self->imageExposed(static_cast<int>(id));
}
void QImageCapture_connect_ImageExposed(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)(int)>(&QImageCapture::imageExposed), self, [=](int id) {
		int sigval1 = id;
		miqt_exec_callback_QImageCapture_ImageExposed(slot, sigval1);
	});
}
void QImageCapture_ImageCaptured(QImageCapture* self, int id, QImage* preview) {
	self->imageCaptured(static_cast<int>(id), *preview);
}
void QImageCapture_connect_ImageCaptured(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)(int, const QImage&)>(&QImageCapture::imageCaptured), self, [=](int id, const QImage& preview) {
		int sigval1 = id;
		const QImage& preview_ret = preview;
		// Cast returned reference into pointer
		QImage* sigval2 = const_cast<QImage*>(&preview_ret);
		miqt_exec_callback_QImageCapture_ImageCaptured(slot, sigval1, sigval2);
	});
}
void QImageCapture_ImageMetadataAvailable(QImageCapture* self, int id, QMediaMetaData* metaData) {
	self->imageMetadataAvailable(static_cast<int>(id), *metaData);
}
void QImageCapture_connect_ImageMetadataAvailable(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)(int, const QMediaMetaData&)>(&QImageCapture::imageMetadataAvailable), self, [=](int id, const QMediaMetaData& metaData) {
		int sigval1 = id;
		const QMediaMetaData& metaData_ret = metaData;
		// Cast returned reference into pointer
		QMediaMetaData* sigval2 = const_cast<QMediaMetaData*>(&metaData_ret);
		miqt_exec_callback_QImageCapture_ImageMetadataAvailable(slot, sigval1, sigval2);
	});
}
void QImageCapture_ImageAvailable(QImageCapture* self, int id, QVideoFrame* frame) {
	self->imageAvailable(static_cast<int>(id), *frame);
}
void QImageCapture_connect_ImageAvailable(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)(int, const QVideoFrame&)>(&QImageCapture::imageAvailable), self, [=](int id, const QVideoFrame& frame) {
		int sigval1 = id;
		const QVideoFrame& frame_ret = frame;
		// Cast returned reference into pointer
		QVideoFrame* sigval2 = const_cast<QVideoFrame*>(&frame_ret);
		miqt_exec_callback_QImageCapture_ImageAvailable(slot, sigval1, sigval2);
	});
}
void QImageCapture_ImageSaved(QImageCapture* self, int id, struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->imageSaved(static_cast<int>(id), fileName_QString);
}
void QImageCapture_connect_ImageSaved(QImageCapture* self, intptr_t slot) {
	MiqtVirtualQImageCapture::connect(self, static_cast<void (QImageCapture::*)(int, const QString&)>(&QImageCapture::imageSaved), self, [=](int id, const QString& fileName) {
		int sigval1 = id;
		const QString fileName_ret = fileName;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray fileName_b = fileName_ret.toUtf8();
		struct miqt_string fileName_ms;
		fileName_ms.len = fileName_b.length();
		fileName_ms.data = static_cast<char*>(malloc(fileName_ms.len));
		memcpy(fileName_ms.data, fileName_b.data(), fileName_ms.len);
		struct miqt_string sigval2 = fileName_ms;
		miqt_exec_callback_QImageCapture_ImageSaved(slot, sigval1, sigval2);
	});
}
struct miqt_string QImageCapture_Tr2(const char* s, const char* c) {
	QString _ret = QImageCapture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QImageCapture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QImageCapture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QImageCapture_CaptureToFile1(QImageCapture* self, struct miqt_string location) {
	QString location_QString = QString::fromUtf8(location.data, location.len);
	return self->captureToFile(location_QString);
}
void QImageCapture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQImageCapture*>( (QImageCapture*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QImageCapture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQImageCapture*)(self) )->virtualbase_MetaObject();
}
void QImageCapture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQImageCapture*>( (QImageCapture*)(self) )->handle__Metacast = slot;
}
void* QImageCapture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQImageCapture*)(self) )->virtualbase_Metacast(param1);
}
void QImageCapture_Delete(QImageCapture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQImageCapture*>( self );
	} else {
		delete self;
	}
}