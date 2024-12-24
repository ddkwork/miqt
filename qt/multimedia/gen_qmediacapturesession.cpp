// +build ignore

#include <QAudioBufferInput>
#include <QAudioInput>
#include <QAudioOutput>
#include <QCamera>
#include <QImageCapture>
#include <QMediaCaptureSession>
#include <QMediaRecorder>
#include <QMetaObject>
#include <QObject>
#include <QScreenCapture>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVideoFrameInput>
#include <QVideoSink>
#include <QWindowCapture>
#include <qmediacapturesession.h>
#include "gen_qmediacapturesession.h"
class MiqtVirtualQMediaCaptureSession : public virtual QMediaCaptureSession {
public:
MiqtVirtualQMediaCaptureSession(): QMediaCaptureSession() {};
MiqtVirtualQMediaCaptureSession(QObject* parent): QMediaCaptureSession(parent) {};

virtual ~MiqtVirtualQMediaCaptureSession() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QMediaCaptureSession::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QMediaCaptureSession_MetaObject(const_cast<MiqtVirtualQMediaCaptureSession*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QMediaCaptureSession::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QMediaCaptureSession::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QMediaCaptureSession_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QMediaCaptureSession::qt_metacast(param1);

	}
};
QMediaCaptureSession* QMediaCaptureSession_new() {
return new MiqtVirtualQMediaCaptureSession();
}
QMediaCaptureSession* QMediaCaptureSession_new2(QObject* parent) {
return new MiqtVirtualQMediaCaptureSession(parent);
}
void QMediaCaptureSession_virtbase(QMediaCaptureSession* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QMediaCaptureSession_MetaObject(const QMediaCaptureSession* self) {
	return (QMetaObject*) self->metaObject();
}
void* QMediaCaptureSession_Metacast(QMediaCaptureSession* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QMediaCaptureSession_Tr(const char* s) {
	QString _ret = QMediaCaptureSession::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QAudioInput* QMediaCaptureSession_AudioInput(const QMediaCaptureSession* self) {
	return self->audioInput();
}
void QMediaCaptureSession_SetAudioInput(QMediaCaptureSession* self, QAudioInput* input) {
	self->setAudioInput(input);
}
QAudioBufferInput* QMediaCaptureSession_AudioBufferInput(const QMediaCaptureSession* self) {
	return self->audioBufferInput();
}
void QMediaCaptureSession_SetAudioBufferInput(QMediaCaptureSession* self, QAudioBufferInput* input) {
	self->setAudioBufferInput(input);
}
QCamera* QMediaCaptureSession_Camera(const QMediaCaptureSession* self) {
	return self->camera();
}
void QMediaCaptureSession_SetCamera(QMediaCaptureSession* self, QCamera* camera) {
	self->setCamera(camera);
}
QImageCapture* QMediaCaptureSession_ImageCapture(QMediaCaptureSession* self) {
	return self->imageCapture();
}
void QMediaCaptureSession_SetImageCapture(QMediaCaptureSession* self, QImageCapture* imageCapture) {
	self->setImageCapture(imageCapture);
}
QScreenCapture* QMediaCaptureSession_ScreenCapture(QMediaCaptureSession* self) {
	return self->screenCapture();
}
void QMediaCaptureSession_SetScreenCapture(QMediaCaptureSession* self, QScreenCapture* screenCapture) {
	self->setScreenCapture(screenCapture);
}
QWindowCapture* QMediaCaptureSession_WindowCapture(QMediaCaptureSession* self) {
	return self->windowCapture();
}
void QMediaCaptureSession_SetWindowCapture(QMediaCaptureSession* self, QWindowCapture* windowCapture) {
	self->setWindowCapture(windowCapture);
}
QVideoFrameInput* QMediaCaptureSession_VideoFrameInput(const QMediaCaptureSession* self) {
	return self->videoFrameInput();
}
void QMediaCaptureSession_SetVideoFrameInput(QMediaCaptureSession* self, QVideoFrameInput* input) {
	self->setVideoFrameInput(input);
}
QMediaRecorder* QMediaCaptureSession_Recorder(QMediaCaptureSession* self) {
	return self->recorder();
}
void QMediaCaptureSession_SetRecorder(QMediaCaptureSession* self, QMediaRecorder* recorder) {
	self->setRecorder(recorder);
}
void QMediaCaptureSession_SetVideoOutput(QMediaCaptureSession* self, QObject* output) {
	self->setVideoOutput(output);
}
QObject* QMediaCaptureSession_VideoOutput(const QMediaCaptureSession* self) {
	return self->videoOutput();
}
void QMediaCaptureSession_SetVideoSink(QMediaCaptureSession* self, QVideoSink* sink) {
	self->setVideoSink(sink);
}
QVideoSink* QMediaCaptureSession_VideoSink(const QMediaCaptureSession* self) {
	return self->videoSink();
}
void QMediaCaptureSession_SetAudioOutput(QMediaCaptureSession* self, QAudioOutput* output) {
	self->setAudioOutput(output);
}
QAudioOutput* QMediaCaptureSession_AudioOutput(const QMediaCaptureSession* self) {
	return self->audioOutput();
}
void QMediaCaptureSession_AudioInputChanged(QMediaCaptureSession* self) {
	self->audioInputChanged();
}
void QMediaCaptureSession_connect_AudioInputChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::audioInputChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_AudioInputChanged(slot);
	});
}
void QMediaCaptureSession_AudioBufferInputChanged(QMediaCaptureSession* self) {
	self->audioBufferInputChanged();
}
void QMediaCaptureSession_connect_AudioBufferInputChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::audioBufferInputChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_AudioBufferInputChanged(slot);
	});
}
void QMediaCaptureSession_CameraChanged(QMediaCaptureSession* self) {
	self->cameraChanged();
}
void QMediaCaptureSession_connect_CameraChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::cameraChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_CameraChanged(slot);
	});
}
void QMediaCaptureSession_ScreenCaptureChanged(QMediaCaptureSession* self) {
	self->screenCaptureChanged();
}
void QMediaCaptureSession_connect_ScreenCaptureChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::screenCaptureChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_ScreenCaptureChanged(slot);
	});
}
void QMediaCaptureSession_WindowCaptureChanged(QMediaCaptureSession* self) {
	self->windowCaptureChanged();
}
void QMediaCaptureSession_connect_WindowCaptureChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::windowCaptureChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_WindowCaptureChanged(slot);
	});
}
void QMediaCaptureSession_VideoFrameInputChanged(QMediaCaptureSession* self) {
	self->videoFrameInputChanged();
}
void QMediaCaptureSession_connect_VideoFrameInputChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::videoFrameInputChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_VideoFrameInputChanged(slot);
	});
}
void QMediaCaptureSession_ImageCaptureChanged(QMediaCaptureSession* self) {
	self->imageCaptureChanged();
}
void QMediaCaptureSession_connect_ImageCaptureChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::imageCaptureChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_ImageCaptureChanged(slot);
	});
}
void QMediaCaptureSession_RecorderChanged(QMediaCaptureSession* self) {
	self->recorderChanged();
}
void QMediaCaptureSession_connect_RecorderChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::recorderChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_RecorderChanged(slot);
	});
}
void QMediaCaptureSession_VideoOutputChanged(QMediaCaptureSession* self) {
	self->videoOutputChanged();
}
void QMediaCaptureSession_connect_VideoOutputChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::videoOutputChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_VideoOutputChanged(slot);
	});
}
void QMediaCaptureSession_AudioOutputChanged(QMediaCaptureSession* self) {
	self->audioOutputChanged();
}
void QMediaCaptureSession_connect_AudioOutputChanged(QMediaCaptureSession* self, intptr_t slot) {
	MiqtVirtualQMediaCaptureSession::connect(self, static_cast<void (QMediaCaptureSession::*)()>(&QMediaCaptureSession::audioOutputChanged), self, [=]() {
		miqt_exec_callback_QMediaCaptureSession_AudioOutputChanged(slot);
	});
}
struct miqt_string QMediaCaptureSession_Tr2(const char* s, const char* c) {
	QString _ret = QMediaCaptureSession::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMediaCaptureSession_Tr3(const char* s, const char* c, int n) {
	QString _ret = QMediaCaptureSession::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMediaCaptureSession_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMediaCaptureSession*>( (QMediaCaptureSession*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QMediaCaptureSession_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQMediaCaptureSession*)(self) )->virtualbase_MetaObject();
}
void QMediaCaptureSession_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMediaCaptureSession*>( (QMediaCaptureSession*)(self) )->handle__Metacast = slot;
}
void* QMediaCaptureSession_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQMediaCaptureSession*)(self) )->virtualbase_Metacast(param1);
}
void QMediaCaptureSession_Delete(QMediaCaptureSession* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQMediaCaptureSession*>( self );
	} else {
		delete self;
	}
}
